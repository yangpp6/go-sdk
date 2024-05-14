package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/yangpp6/go-sdk/sdk/region"
	"io"
	"testing"
)

func TestRegionList_Do(t *testing.T) {
	client := BuildProdClient()
	response, err := client.Apis.RegionListApi.Do(context.Background(), ProdCredential, &region.RegionListRequest{})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	jsonstr, _ := json.Marshal(response.RegionList)
	fmt.Println(string(jsonstr))
}
