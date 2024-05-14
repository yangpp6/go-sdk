package test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"openapi-sdk-go/sdk/common"
	"openapi-sdk-go/sdk/keypair"
)

func TestKeypaireCreateApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("", "")
	do, _ := BuildProdClient().Apis.KeypairCreateApi.Do(context.Background(), credential, &keypair.KeypairCreateRequest{
		RegionId:    ProdHuaDong1RegionId,
		KeyPairName: "sdsdfdfss222",
		ProjectId:   "0",
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}

func TestKeypaireDeleteApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("e7b8b85b08734185b7694692e93b3a84", "c20ac65747fb439693cf6d3a97de751c")
	do, _ := BuildProdClient().Apis.KeypairDeleteApi.Do(context.Background(), credential, &keypair.KeypairDeleteRequest{
		RegionId:    ProdHuaDong1RegionId,
		KeyPairName: "sdsdfdfss222",
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}

func TestKeypaireDetailApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("e7b8b85b08734185b7694692e93b3a84", "c20ac65747fb439693cf6d3a97de751c")
	do, _ := BuildProdClient().Apis.KeypairDetailApi.Do(context.Background(), credential, &keypair.KeypairDetailRequest{
		RegionId:    ProdHuaDong1RegionId,
		KeyPairName: "sdsdfdfss222",
		PageNo:      1,
		PageSize:    10,
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}

func TestKeypaireImportApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("e7b8b85b08734185b7694692e93b3a84", "c20ac65747fb439693cf6d3a97de751c")
	do, _ := BuildProdClient().Apis.KeypairImportApi.Do(context.Background(), credential, &keypair.KeypairImportRequest{
		RegionId:    ProdHuaDong1RegionId,
		KeyPairName: "sdsdfdfss222",
		PublicKey:   "afaaagagg",
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}

func TestKeypairAttachApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("e7b8b85b08734185b7694692e93b3a84", "c20ac65747fb439693cf6d3a97de751c")
	do, _ := BuildProdClient().Apis.KeypairAttachApi.Do(context.Background(), credential, &keypair.KeypairAttachRequest{
		RegionId:    ProdHuaDong1RegionId,
		KeyPairName: "sdsdfdfss222",
		InstanceId:  "afaaagagg",
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}

func TestKeypairDetachApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("e7b8b85b08734185b7694692e93b3a84", "c20ac65747fb439693cf6d3a97de751c")
	do, _ := BuildProdClient().Apis.KeypairDetachApi.Do(context.Background(), credential, &keypair.KeypairDetachRequest{
		RegionId:    ProdHuaDong1RegionId,
		KeyPairName: "sdsdfdfss222",
		InstanceId:  "afaaagagg",
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}
