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

func TestNewEcsResetPasswordApi_Do(t *testing.T) {
	var ak = "<YOUR AK>"
	var sk = "<YOUR SK>"
	client := test.BuildProdClient()
	credential, _ := common.NewCredential(ak, sk)
	response, err := client.Apis.EcsResetPasswordApi.Do(context.Background(), credential, &ecs.EcsResetPasswordRequest{
		RegionId:    "bb9fdb42056f11eda1610242ac110002",
		InstanceId:  "77493826-d038-2a9c-f684-e2f6adabeba3",
		NewPassword: "test-test-test-960",
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}