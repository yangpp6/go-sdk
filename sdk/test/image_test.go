package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"openapi-sdk-go/sdk/common"
	"openapi-sdk-go/sdk/ecs"
	"openapi-sdk-go/sdk/image"
	"testing"
)

func TestImageDetailApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("", "")
	do, _ := BuildProdClient().Apis.ImageDetailApi.Do(context.Background(), credential, &image.ImageDetailRequest{
		RegionId: ProdHuaDong1RegionId,
		ImageId:  "aa",
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}

func TestImageDeleteApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("", "")
	do, _ := BuildProdClient().Apis.ImageDeleteApi.Do(context.Background(), credential, &image.ImageDeleteRequest{
		RegionId: ProdHuaDong1RegionId,
		ImageId:  "aa",
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}
func TestImageList1Api_Do(t *testing.T) {
	client := BuildProdClient()
	credential, _ := common.NewCredential("", "")
	response, err := client.Apis.ImageListApi.Do(context.Background(), credential, &image.ImageListRequest{
		RegionId:   ProdHuaDong1RegionId,
		AzName:     ProdHuaDong1Az1AzName,
		Visibility: 1,
		PageNo:     1,
		PageSize:   1,
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	responseJsonStr, _ := json.Marshal(response)
	fmt.Printf("成功响应为：%s", string(responseJsonStr))
}
func TestImageUpdateApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("", "")
	do, _ := BuildProdClient().Apis.ImageUpdateApi.Do(context.Background(), credential, &image.ImageUpdateRequest{
		RegionId:   ProdHuaDong1RegionId,
		ImageId:    "aa",
		ImageName:  "test",
		MaximumRam: 16,
		MinimumRam: 4,
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}

func TestNewEcsRedeployApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("", "")
	do, _ := BuildProdClient().Apis.EcsRedeployApi.Do(context.Background(), credential, &ecs.EcsRedeployRequest{
		RegionId:   "200000001852",
		InstanceId: "9416fb88-18a3-87c5-dd91-59f1d2e2ac3b",
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}
