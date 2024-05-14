package main

import (
	"context"
	"flag"
	"github.com/yangpp6/go-sdk/sdk/common"
	"github.com/yangpp6/go-sdk/sdk/security_group"
	"log"
	"net/http"
)

func sgOperation(ak, sk, regionID, vpcID string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtVpcEndPoint: "ctvpc-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)

	createHandler := security_group.NewSecurityGroupCreateApi(client)
	ctx := context.TODO()
	createRes, err := createHandler.Do(ctx, credential, &security_group.SecurityGroupCreateRequest{
		RegionId:    regionID,
		VpcId:       vpcID,
		ClientToken: "xyz",
		Name:        "test",
		Description: "test",
	})

	if err != nil {
		panic(err)
	}

	sgID := createRes.SecurityGroupId

	descHandler := security_group.NewSecurityGroupDescribeAttributeApi(client)
	sg, err := descHandler.Do(ctx, credential, &security_group.SecurityGroupDescribeAttributeRequest{
		RegionId:        regionID,
		SecurityGroupId: sgID,
	})

	if err != nil {
		panic(err)
	}

	log.Printf("%+v\n", sg)

	updateHandler := security_group.NewSecurityGroupModifyAttributionApi(client)
	_, err = updateHandler.Do(ctx, credential, &security_group.SecurityGroupModifyAttributionRequest{
		SecurityGroupId: sgID,
		RegionId:        regionID,
		Name:            "test-test",
	})

	if err != nil {
		panic(err)
	}

	deleteHandler := security_group.NewSecurityGroupDeleteApi(client)
	_, err = deleteHandler.Do(ctx, credential, &security_group.SecurityGroupDeleteRequest{
		SecurityGroupId: sgID,
		RegionId:        regionID,
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

	if len(ak) == 0 || len(sk) == 0 || len(regionID) == 0 || len(vpcID) == 0 {
		log.Print("ak or sk or region-id or vpc-id is required")
		return
	}

	sgOperation(ak, sk, regionID, vpcID)
}
