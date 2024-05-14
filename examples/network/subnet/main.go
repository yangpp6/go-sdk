package main

import (
	"context"
	"flag"
	"github.com/yangpp6/go-sdk/sdk/common"
	"github.com/yangpp6/go-sdk/sdk/subnet"
	"log"
	"net/http"
)

func listSubnets(ak, sk, regionID string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtVpcEndPoint: "ctvpc-global.ctapi.ctyun.cn",
	}

	handler := subnet.NewSubnetListApi(client)
	credential, _ := common.NewCredential(ak, sk)
	res, err := handler.Do(context.TODO(), credential, &subnet.SubnetListRequest{RegionId: regionID})
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", res)
}

func subnetOperation(ak, sk, regionID, vpcID string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtVpcEndPoint: "ctvpc-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)

	createHandler := subnet.NewSubnetCreateApi(client)
	ctx := context.TODO()
	createRes, err := createHandler.Do(ctx, credential, &subnet.SubnetCreateRequest{
		RegionId:    regionID,
		VpcId:       vpcID,
		ClientToken: "xyz",
		Name:        "test",
		Cidr:        "192.168.1.0/24",
		Description: "test",
	})

	if err != nil {
		panic(err)
	}

	updateHandler := subnet.NewSubnetUpdateApi(client)
	_, err = updateHandler.Do(ctx, credential, &subnet.SubnetUpdateRequest{
		SubnetId: createRes.SubnetId,
		RegionId: regionID,
		Name:     "test-test",
	})

	if err != nil {
		panic(err)
	}

	deleteHandler := subnet.NewSubnetDeleteApi(client)
	_, err = deleteHandler.Do(ctx, credential, &subnet.SubnetDeleteRequest{
		SubnetId: createRes.SubnetId,
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
	var vpcID string
	flag.StringVar(&action, "action", "list", "example action: list / operation")
	flag.StringVar(&ak, "ak", "", "access key")
	flag.StringVar(&sk, "sk", "", "secret key")
	flag.StringVar(&regionID, "region-id", "", "region id")
	flag.StringVar(&vpcID, "vpc-id", "", "vpc id")
	flag.Parse()

	if len(ak) == 0 || len(sk) == 0 || len(regionID) == 0 {
		log.Print("ak or sk or region-id is required")
		return
	}

	if action == "list" {
		listSubnets(ak, sk, regionID)
	} else if action == "operation" {
		if len(vpcID) == 0 {
			log.Print("vpc-id field required")
		} else {
			subnetOperation(ak, sk, regionID, vpcID)
		}
	} else {
		log.Print("unknown action")
	}
}
