package main

import (
	"context"
	"flag"
	"github.com/yangpp6/go-sdk/sdk/common"
	"github.com/yangpp6/go-sdk/sdk/vpc"
	"log"
	"net/http"
)

func listVpcs(ak, sk, regionID string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtVpcEndPoint: "ctvpc-global.ctapi.ctyun.cn",
	}

	handler := vpc.NewVpcListApi(client)
	credential, _ := common.NewCredential(ak, sk)
	res, err := handler.Do(context.TODO(), credential, &vpc.VpcListRequest{RegionId: regionID})
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", res)
}

func vpcOperation(ak, sk, regionID string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtVpcEndPoint: "ctvpc-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)

	createHandler := vpc.NewVpcCreateApi(client)
	ctx := context.TODO()
	createRes, err := createHandler.Do(ctx, credential, &vpc.VpcCreateRequest{
		RegionId:    regionID,
		ClientToken: "xyz",
		Name:        "test",
		Cidr:        "192.168.1.0/24",
		Description: "test",
		EnableIpv6:  false,
	})

	if err != nil {
		panic(err)
	}

	updateHandler := vpc.NewVpcUpdateApi(client)
	_, err = updateHandler.Do(ctx, credential, &vpc.VpcUpdateRequest{
		VpcId:       createRes.VpcId,
		ClientToken: "xyz",
		RegionId:    regionID,
		Name:        "test-test",
	})

	if err != nil {
		panic(err)
	}

	deleteHandler := vpc.NewVpcDeleteApi(client)
	_, err = deleteHandler.Do(ctx, credential, &vpc.VpcDeleteRequest{
		VpcId:    createRes.VpcId,
		RegionId: regionID,
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
	flag.StringVar(&action, "action", "list", "example action: list / operation")
	flag.StringVar(&ak, "ak", "", "access key")
	flag.StringVar(&sk, "sk", "", "secret key")
	flag.StringVar(&regionID, "region-id", "", "region id")
	flag.Parse()

	if len(ak) == 0 || len(sk) == 0 || len(regionID) == 0 {
		log.Print("ak or sk or region-id is required")
		return
	}

	if action == "list" {
		listVpcs(ak, sk, regionID)
	} else if action == "operation" {
		vpcOperation(ak, sk, regionID)
	} else {
		log.Print("unknown action")
	}
}
