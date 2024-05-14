package test

import (
	"context"
	"fmt"
	"github.com/yangpp6/go-sdk/sdk/ecs"
	"io"
	"testing"
)

func TestEcsJoinSecurityGroupApi_Do(t *testing.T) {
	client := BuildProdClient()
	_, err := client.Apis.EcsJoinSecurityGroupApi.Do(context.Background(), ProdCredential, &ecs.EcsJoinSecurityGroupRequest{
		RegionId:        ProdHuaDong1RegionId,
		SecurityGroupId: "sg-3ol398ey9w",
		InstanceId:      "917af019-96bd-0437-2bba-c11ae45d243a",
		Action:          "joinSecurityGroup",
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
}
