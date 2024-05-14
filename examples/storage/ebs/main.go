package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/yangpp6/go-sdk/sdk/common"
	"github.com/yangpp6/go-sdk/sdk/ebs"
	"net/http"
)

func ebsShow(ak, sk, regionID, diskID string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEbsEndPoint: "ebs-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)

	createHandler := ebs.NewEbsShowApi(client)
	ctx := context.TODO()
	createRes, err := createHandler.Do(ctx, credential, &ebs.EbsShowRequest{
		RegionId: regionID,
		DiskId:   diskID,
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(createRes)
	fmt.Printf(string(marshal))
}

func ebsChangeName(ak, sk, regionID, diskID, diskName string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEbsEndPoint: "ebs-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)

	createHandler := ebs.NewEbsChangeNameApi(client)
	ctx := context.TODO()
	createRes, err := createHandler.Do(ctx, credential, &ebs.EbsChangeNameRequest{
		RegionId: regionID,
		DiskId:   diskID,
		DiskName: diskName,
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(createRes)
	fmt.Printf(string(marshal))
}

func ebsAssociate(ak, sk, regionID, diskID, instanceID string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEbsEndPoint: "ebs-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)

	createHandler := ebs.NewEbsAssociateApi(client)
	ctx := context.TODO()
	createRes, err := createHandler.Do(ctx, credential, &ebs.EbsAssociateRequest{
		RegionId:   regionID,
		DiskId:     diskID,
		InstanceID: instanceID,
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(createRes)
	fmt.Printf(string(marshal))
}

func ebsDisassociate(ak, sk, regionID, diskID string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEbsEndPoint: "ebs-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)

	createHandler := ebs.NewEbsDisassociateApi(client)
	ctx := context.TODO()
	res, err := createHandler.Do(ctx, credential, &ebs.EbsDisassociateRequest{
		RegionId: regionID,
		DiskId:   diskID,
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func ebsCreate(ak, sk, clientToken, regionID, azName string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEbsEndPoint: "ebs-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)

	createHandler := ebs.NewEbsCreateApi(client)
	ctx := context.TODO()
	res, err := createHandler.Do(ctx, credential, &ebs.EbsCreateRequest{
		ClientToken: clientToken,
		DiskName:    "sdktest",
		DiskMode:    "VBD",
		DiskType:    "SATA",
		DiskSize:    20,
		RegionID:    regionID,
		AzName:      azName,
		OnDemand:    true,
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func ebsChangeSize(ak, sk, clientToken, regionID, diskID string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEbsEndPoint: "ebs-global.ctapi.ctyun.cn",
	}
	credential, _ := common.NewCredential(ak, sk)
	createHandler := ebs.NewEbsChangeSizeApi(client)
	ctx := context.TODO()
	res, err := createHandler.Do(ctx, credential, &ebs.EbsChangeSizeRequest{
		ClientToken: clientToken,
		DiskId:      diskID,
		DiskSize:    30,
		RegionId:    regionID,
	})
	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func ebsDelete(ak, sk, clientToken, regionID, diskID string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEbsEndPoint: "ebs-global.ctapi.ctyun.cn",
	}
	credential, _ := common.NewCredential(ak, sk)
	createHandler := ebs.NewEbsDeleteApi(client)
	ctx := context.TODO()
	res, err := createHandler.Do(ctx, credential, &ebs.EbsDeleteRequest{
		ClientToken: clientToken,
		DiskID:      diskID,
		RegionID:    regionID,
	})
	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func main() {
	var ak string
	var sk string
	var regionID string
	var azName string
	var clientToken string
	var diskID string
	var instanceID string
	flag.StringVar(&ak, "ak", "", "access key")
	flag.StringVar(&sk, "sk", "", "secret key")
	flag.StringVar(&regionID, "region-id", "", "region id")
	flag.StringVar(&clientToken, "clientToken", "", "client token")
	flag.StringVar(&diskID, "diskID", "", "disk id")
	flag.StringVar(&azName, "azName", "", "az name")
	flag.StringVar(&instanceID, "instanceID", "", "instance id")
	flag.Parse()
	ebsCreate(ak, sk, clientToken, regionID, azName)
	ebsShow(ak, sk, regionID, diskID)
	ebsChangeName(ak, sk, regionID, diskID, "test")
	ebsChangeSize(ak, sk, clientToken, regionID, diskID)
	ebsAssociate(ak, sk, regionID, diskID, instanceID)
	ebsDisassociate(ak, sk, regionID, diskID)
	ebsDelete(ak, sk, clientToken, regionID, diskID)
}
