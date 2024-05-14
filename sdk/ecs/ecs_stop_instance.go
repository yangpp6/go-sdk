package ecs

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// ecsStopInstanceApi  关闭一台云主机
// https://www.ctyun.cn/document/10026730/10106393
type ecsStopInstanceApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEcsStopInstanceApi(client *common.CtyunSender) common.ApiHandler[EcsStopInstanceRequest, EcsStopInstanceResponse] {
	return &ecsStopInstanceApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/stop-instance",
		},
	}
}

func (v *ecsStopInstanceApi) Do(ctx context.Context, credential *common.Credential, req *EcsStopInstanceRequest) (*EcsStopInstanceResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&ecsStopInstanceRealRequest{
		RegionID:   req.RegionId,
		InstanceID: req.InstanceId,
		Force:      req.Force,
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
	var realResponse EcsStopInstanceRealResponse
	err = response.ParseWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	return &EcsStopInstanceResponse{
		JobId: realResponse.JobID,
	}, nil
}

type ecsStopInstanceRealRequest struct {
	RegionID   string `json:"regionID"`
	InstanceID string `json:"instanceID"`
	Force      bool   `json:"force"`
}

type EcsStopInstanceRequest struct {
	RegionId   string
	InstanceId string
	Force      bool
}

type EcsStopInstanceRealResponse struct {
	JobID string `json:"jobID"`
}

type EcsStopInstanceResponse struct {
	JobId string
}
