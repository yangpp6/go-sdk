package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/yangpp6/go-sdk/sdk/common"
	"github.com/yangpp6/go-sdk/sdk/region"
	"testing"
)

func TestRegionListApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("e7b8b85b08734185b7694692e93b3a84", "c20ac65747fb439693cf6d3a97de751c")
	do, _ := BuildProdClient().Apis.RegionListApi.Do(context.Background(), credential, &region.RegionListRequest{
		RegionName: "华东1",
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}
