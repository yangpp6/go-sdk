package ecs

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// ecsResetPasswordApi  更新云主机密码
// https://www.ctyun.cn/document/10026730/10106390
type ecsRedeployApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEcsRedeployApi(client *common.CtyunSender) common.ApiHandler[EcsRedeployRequest, EcsRedeployResponse] {
	return &ecsRedeployApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/redeploy-instance",
		},
	}
}

func (v *ecsRedeployApi) Do(ctx context.Context, credential *common.Credential, req *EcsRedeployRequest) (*EcsRedeployResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&ecsRedeployRealRequest{
		RegionID:   req.RegionId,
		InstanceID: req.InstanceId,
	})
	if err != nil {
		return nil, err
	}

	// 发起请求
	response, err := v.client.SendCtEcs(ctx, builder)
	if err != nil {
		return nil, err
	}

	// 解析返回
	var realResponse EcsRedeployRealResponse
	err = response.ParseWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	return &EcsRedeployResponse{
		InstanceId: realResponse.InstanceID,
	}, nil
}

type ecsRedeployRealRequest struct {
	RegionID   string `json:"regionID"`
	InstanceID string `json:"instanceID"`
}

type EcsRedeployRealResponse struct {
	InstanceID string `json:"instanceID"`
}

type EcsRedeployRequest struct {
	RegionId    string
	InstanceId  string
	NewPassword string
}

type EcsRedeployResponse struct {
	InstanceId string
}
