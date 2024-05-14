package test

import (
	"context"
	"fmt"
	"io"
	"openapi-sdk-go/sdk/security_group"
	"testing"
)

func TestSecurityGroupDescribeAttributeApi_Do(t *testing.T) {
	client := BuildProdClient()
	_, err := client.Apis.SecurityGroupDescribeAttributeApi.Do(context.Background(), ProdCredential, &security_group.SecurityGroupDescribeAttributeRequest{
		RegionId:        ProdHuaDong1RegionId,
		SecurityGroupId: "ssss",
		Direction:       "all",
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
}
