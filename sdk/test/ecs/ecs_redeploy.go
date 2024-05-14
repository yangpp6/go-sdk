package test

import (
	"encoding/json"
	"fmt"
	"openapi-sdk-go/sdk/common"
	"openapi-sdk-go/sdk/ecs"
	"openapi-sdk-go/sdk/test"
	"testing"

	"golang.org/x/net/context"
)

func TestNewEcsRedeployApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("e7b8b85b08734185b7694692e93b3a84", "c20ac65747fb439693cf6d3a97de751c")
	do, _ := test.BuildProdClient().Apis.EcsRedeployApi.Do(context.Background(), credential, &ecs.EcsRedeployRequest{
		RegionId:   test.ProdHuaDong1RegionId,
		InstanceId: "77493826-d038-2a9c-f684-e2f6adabeba3",
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}
