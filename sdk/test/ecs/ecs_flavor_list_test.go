package test

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"openapi-sdk-go/sdk/common"
	"openapi-sdk-go/sdk/ecs"
	"openapi-sdk-go/sdk/test"
	"testing"
)

func TestEcsInstanceStatusListApi_Do(t *testing.T) {
	var ak = "<YOUR AK>"
	var sk = "<YOUR SK>"
	client := test.BuildProdClient()
	credential, _ := common.NewCredential(ak, sk)
	response, err := client.Apis.EcsFlavorListApi.Do(context.Background(), credential, &ecs.EcsFlavorListRequest{
		RegionId:     "bb9fdb42056f11eda1610242ac110002",
		AzName:       "cn-huadong1-jsnj1A-public-ctcloud",
		FlavorType:   "CPU_KS1",
		FlavorName:   "ks1.medium.2",
		FlavorCpu:    1,
		FlavorArch:   "arm",
		FlavorSeries: "ks",
		FlavorId:     "b6779240-5649-803b-4a4c-8fc59d310ecf",
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
