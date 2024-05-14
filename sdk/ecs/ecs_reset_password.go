package ecs

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// ecsResetPasswordApi  更新云主机密码
// https://www.ctyun.cn/document/10026730/10106390
type ecsResetPasswordApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEcsResetPasswordApi(client *common.CtyunSender) common.ApiHandler[EcsResetPasswordRequest, EcsResetPasswordResponse] {
	return &ecsResetPasswordApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/reset-password",
		},
	}
}

func (v *ecsResetPasswordApi) Do(ctx context.Context, credential *common.Credential, req *EcsResetPasswordRequest) (*EcsResetPasswordResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&ecsResetPasswordRealRequest{
		RegionID:    req.RegionId,
		InstanceID:  req.InstanceId,
		NewPassword: req.NewPassword,
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
	var realResponse EcsResetPasswordRealResponse
	err = response.ParseWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	return &EcsResetPasswordResponse{
		InstanceId: realResponse.InstanceID,
	}, nil
}

type ecsResetPasswordRealRequest struct {
	RegionID    string `json:"regionID"`
	InstanceID  string `json:"instanceID"`
	NewPassword string `json:"newPassword"`
}

type EcsResetPasswordRealResponse struct {
	InstanceID string `json:"instanceID"`
}

type EcsResetPasswordRequest struct {
	RegionId    string
	InstanceId  string
	NewPassword string
}

type EcsResetPasswordResponse struct {
	InstanceId string
}
