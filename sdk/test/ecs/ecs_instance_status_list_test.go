package test

import (
	"encoding/json"
	"fmt"
	"github.com/yangpp6/go-sdk/sdk/common"
	"github.com/yangpp6/go-sdk/sdk/ecs"
	"github.com/yangpp6/go-sdk/sdk/test"
	"golang.org/x/net/context"
	"io"
	"testing"
)

func TestNewEcsInstanceStatusListApi_Do(t *testing.T) {
	var ak = "<YOUR AK>"
	var sk = "<YOUR SK>"
	client := test.BuildProdClient()
	credential, _ := common.NewCredential(ak, sk)
	response, err := client.Apis.EcsInstanceStatusListApi.Do(context.Background(), credential, &ecs.EcsInstanceStatusListRequest{
		RegionId:       "bb9fdb42056f11eda1610242ac110002",
		AzName:         "cn-huadong1-jsnj1A-public-ctcloud",
		InstanceIdList: "77493826-d038-2a9c-f684-e2f6adabeba3,76d0c1cb-553f-be99-6e2e-723bcb096303",
		PageNo:         1,
		PageSize:       10,
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
