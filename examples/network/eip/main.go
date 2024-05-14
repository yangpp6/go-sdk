package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/yangpp6/go-sdk/sdk/common"
	"github.com/yangpp6/go-sdk/sdk/eip"
	"log"
	"net/http"
)

func eipOperation(ak, sk, regionID, vmID string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtVpcEndPoint: "ctvpc-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)

	createHandler := eip.NewEipCreateApi(client)
	ctx := context.TODO()
	createRes, err := createHandler.Do(ctx, credential, &eip.EipCreateRequest{
		RegionId:    regionID,
		ClientToken: "abcedefghij",
		Name:        "test",
		CycleType:   "on_demand",
	})

	if err != nil {
		panic(err)
	}

	for {
		if createRes.MasterResourceStatus == "in_progress" {
			createRes, err = createHandler.Do(ctx, credential, &eip.EipCreateRequest{
				RegionId:    regionID,
				ClientToken: "abcedefghij",
				Name:        "test",
				CycleType:   "on_demand",
			})
			if err != nil {
				panic(err)
			}
		}

		if createRes.MasterResourceStatus != "in_progress" {
			break
		}
	}

	var eipID string
	if createRes.MasterResourceStatus == "started" {
		eipID = createRes.EipId
	} else {
		panic(fmt.Errorf("order status %s", createRes.MasterResourceStatus))
	}

	changeNameHandler := eip.NewEipChangeNameApi(client)
	_, err = changeNameHandler.Do(ctx, credential, &eip.EipChangeNameRequest{
		ClientToken: "yyyyyy",
		RegionId:    regionID,
		EipId:       eipID,
		Name:        "kkkkk",
	})
	if err != nil {
		panic(err)
	}

	associateHandler := eip.NewEipAssociateApi(client)
	_, err = associateHandler.Do(ctx, credential, &eip.EipAssociateRequest{
		ClientToken:     "yyyyyy",
		RegionId:        regionID,
		EipId:           eipID,
		AssociationType: 1, // 1 vm; 2 bm; 3 vip
		AssociationId:   vmID,
	})
	if err != nil {
		panic(err)
	}

	disassociationHandler := eip.NewEipDisassociateApi(client)
	_, err = disassociationHandler.Do(ctx, credential, &eip.EipDisassociateRequest{
		ClientToken: "yyyyyy",
		RegionId:    regionID,
		EipId:       eipID,
	})
	if err != nil {
		panic(err)
	}

	changeSpecHandler := eip.NewEipModifySpecApi(client)
	_, err = changeSpecHandler.Do(ctx, credential, &eip.EipModifySpecRequest{
		ClientToken: "xxxxxxxxxxx",
		RegionId:    regionID,
		EipId:       eipID,
		Bandwidth:   2,
	})
	if err != nil {
		panic(err)
	}

	deleteHandler := eip.NewEipDeleteApi(client)
	_, err = deleteHandler.Do(ctx, credential, &eip.EipDeleteRequest{
		ClientToken: "zzzzzzzzzzzz",
		RegionId:    regionID,
		EipId:       eipID,
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	var action string
	var ak string
	var sk string
	var regionID string
	var vmID string
	flag.StringVar(&action, "action", "list", "example action: list / operation")
	flag.StringVar(&ak, "ak", "", "access key")
	flag.StringVar(&sk, "sk", "", "secret key")
	flag.StringVar(&regionID, "region-id", "", "region id")
	flag.StringVar(&vmID, "vm-id", "", "vm id")
	flag.Parse()

	if len(ak) == 0 || len(sk) == 0 || len(regionID) == 0 || len(vmID) == 0 {
		log.Print("ak or sk or region-id or vm-id is required")
		return
	}

	eipOperation(ak, sk, regionID, vmID)
}
