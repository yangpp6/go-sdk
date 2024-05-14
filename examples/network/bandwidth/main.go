package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"openapi-sdk-go/sdk/bandwidth"
	"openapi-sdk-go/sdk/common"
)

func bandwidthOperation(ak, sk, regionID, eipID string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtVpcEndPoint: "https://ctvpc-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)

	createHandler := bandwidth.NewBandwidthCreateApi(client)
	ctx := context.TODO()
	createRes, err := createHandler.Do(ctx, credential, &bandwidth.BandwidthCreateRequest{
		RegionID:    regionID,
		ClientToken: "abcedefghijklmn",
		Name:        "test",
		CycleType:   "on_demand",
		Bandwidth:   5,
	})

	if err != nil {
		panic(err)
	}

	for {
		if createRes.MasterResourceStatus == "in_progress" {
			createRes, err = createHandler.Do(ctx, credential, &bandwidth.BandwidthCreateRequest{
				RegionID:    regionID,
				ClientToken: "abcedefghijklmn",
				Name:        "test",
				CycleType:   "on_demand",
				Bandwidth:   5,
			})
			if err != nil {
				panic(err)
			}
		}

		if createRes.MasterResourceStatus != "in_progress" {
			break
		}
	}

	var bandwidthID string
	if createRes.MasterResourceStatus == "started" {
		bandwidthID = createRes.BandwidthId
	} else {
		panic(fmt.Errorf("order status %s", createRes.MasterResourceStatus))
	}

	changeNameHandler := bandwidth.NewBandwidthChangeNameApi(client)
	_, err = changeNameHandler.Do(ctx, credential, &bandwidth.BandwidthChangeNameRequest{
		ClientToken: "yyyyyy",
		RegionID:    regionID,
		BandwidthID: bandwidthID,
		Name:        "wwwww",
	})
	if err != nil {
		panic(err)
	}

	addEipHandler := bandwidth.NewBandwidthAssociationEipApi(client)
	_, err = addEipHandler.Do(ctx, credential, &bandwidth.BandwidthAssociationEipRequest{
		ClientToken: "yyyyyy",
		RegionID:    regionID,
		BandwidthID: bandwidthID,
		EipIDs:      []string{eipID},
	})
	if err != nil {
		panic(err)
	}

	removeEipHandler := bandwidth.NewBandwidthDisassociationEipApi(client)
	_, err = removeEipHandler.Do(ctx, credential, &bandwidth.BandwidthDisassociationEipRequest{
		ClientToken: "yyyyyy",
		RegionId:    regionID,
		BandwidthID: bandwidthID,
		EipIds:      []string{eipID},
	})
	if err != nil {
		panic(err)
	}

	showHandler := bandwidth.NewBandwidthShowApi(client)
	showRes, err := showHandler.Do(ctx, credential, &bandwidth.BandwidthShowRequest{
		BandwidthID: bandwidthID,
		RegionId:    regionID,
	})
	if err != nil {
		panic(err)
	}

	log.Printf("%+v\n", showRes)

	modifySpecHandler := bandwidth.NewBandwidthChangeSpecApi(client)
	_, err = modifySpecHandler.Do(ctx, credential, &bandwidth.BandwidthChangeSpecRequest{
		ClientToken: "tttttttttttttt",
		RegionID:    regionID,
		BandwidthID: bandwidthID,
		Bandwidth:   10,
	})
	if err != nil {
		panic(err)
	}

	deleteHandler := bandwidth.NewBandwidthDeleteApi(client)
	_, err = deleteHandler.Do(ctx, credential, &bandwidth.BandwidthDeleteRequest{
		ClientToken: "oooooooooooo",
		RegionID:    regionID,
		BandwidthID: bandwidthID,
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
	var eipID string
	flag.StringVar(&action, "action", "list", "example action: list / operation")
	flag.StringVar(&ak, "ak", "", "access key")
	flag.StringVar(&sk, "sk", "", "secret key")
	flag.StringVar(&regionID, "region-id", "", "region id")
	flag.StringVar(&eipID, "eip-id", "", "vm id")
	flag.Parse()

	if len(ak) == 0 || len(sk) == 0 || len(regionID) == 0 || len(eipID) == 0 {
		log.Print("ak or sk or region-id or eip-id is required")
		return
	}

	bandwidthOperation(ak, sk, regionID, eipID)
}
