package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/yangpp6/go-sdk/sdk/image"
	"io"
	"testing"
)

func TestImageListApi_Do(t *testing.T) {
	client := BuildDevClient()
	response, err := client.Apis.ImageListApi.Do(context.Background(), TestCredential, &image.ImageListRequest{
		RegionId:   TestNeiMeng8RegionId,
		AzName:     TestNeiMeng8Az1AzName,
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
