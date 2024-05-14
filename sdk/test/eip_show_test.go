package test

import (
	"context"
	"fmt"
	"io"
	"openapi-sdk-go/sdk/eip"
	"testing"
)

func TestEipShowApi_Do(t *testing.T) {
	client := BuildProdClient()
	_, err := client.Apis.EipShowApi.Do(context.Background(), ProdCredential, &eip.EipShowRequest{
		RegionId: ProdHuaDong1RegionId,
		EipId:    "eip-dmw8yb7uij",
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
}