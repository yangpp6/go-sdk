package test

import (
	"context"
	"encoding/json"
	"fmt"
	"openapi-sdk-go/sdk/common"
	"openapi-sdk-go/sdk/ebs"
	"testing"

	"github.com/google/uuid"
)

func TestEbsShowApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("", "")
	do, _ := BuildProdClient().Apis.EbsShowApi.Do(context.Background(), credential, &ebs.EbsShowRequest{
		RegionId: ProdHuaDong1RegionId,
		DiskId:   "b85dee4c-eb7a-4345-aab3-6920b5ba04d2",
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}

func TestEbsChangeNameApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("", "")
	do, _ := BuildProdClient().Apis.EbsChangeNameApi.Do(context.Background(), credential, &ebs.EbsChangeNameRequest{
		RegionId: ProdHuaDong1RegionId,
		DiskId:   "b85dee4c-eb7a-4345-aab3-6920b5ba04d2",
		DiskName: "leontest",
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}

func TestEbsAssociateApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("", "")
	do, _ := BuildProdClient().Apis.EbsAssociateApi.Do(context.Background(), credential, &ebs.EbsAssociateRequest{
		RegionId:   ProdHuaDong1RegionId,
		DiskId:     "b85dee4c-eb7a-4345-aab3-6920b5ba04d2",
		InstanceID: "leontest",
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}

func TestEbsDisassociateApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("", "")
	do, _ := BuildProdClient().Apis.EbsDisassociateApi.Do(context.Background(), credential, &ebs.EbsDisassociateRequest{
		RegionId: ProdHuaDong1RegionId,
		DiskId:   "b85dee4c-eb7a-4345-aab3-6920b5ba04d2",
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}

func TestEbsCreateApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("", "")
	do, _ := BuildProdClient().Apis.EbsCreateApi.Do(context.Background(), credential, &ebs.EbsCreateRequest{
		ClientToken: uuid.NewString(),
		DiskName:    "sdktest",
		DiskMode:    "VBD",
		DiskType:    "SATA",
		DiskSize:    20,
		RegionID:    ProdHuaDong1RegionId,
		AzName:      ProdHuaDong1Az1AzName,
		OnDemand:    true,
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}

func TestEbsChangeSizeApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("", "")
	do, _ := BuildProdClient().Apis.EbsChangeSizeApi.Do(context.Background(), credential, &ebs.EbsChangeSizeRequest{
		ClientToken: uuid.NewString(),
		DiskId:      "d904f2ab-dcc8-472f-b294-787572b96826",
		DiskSize:    30,
		RegionId:    ProdHuaDong1RegionId,
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}

func TestEbsDeleteApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("", "")
	do, _ := BuildProdClient().Apis.EbsDeleteApi.Do(context.Background(), credential, &ebs.EbsDeleteRequest{
		ClientToken: uuid.NewString(),
		RegionID:    ProdHuaDong1RegionId,
		DiskID:      "d904f2ab-dcc8-472f-b294-787572b96826",
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}
