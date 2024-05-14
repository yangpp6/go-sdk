package ecs

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// ecsStartInstanceApi  开启一台云主机
// https://www.ctyun.cn/document/10026730/10106397
type ecsStartInstanceApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEcsStartInstanceApi(client *common.CtyunSender) common.ApiHandler[EcsStartInstanceRequest, EcsStartInstanceResponse] {
	return &ecsStartInstanceApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/start-instance",
		},
	}
}

func (v *ecsStartInstanceApi) Do(ctx context.Context, credential *common.Credential, req *EcsStartInstanceRequest) (*EcsStartInstanceResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&ecsStartInstanceRealRequest{
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
	var realResponse EcsStartInstanceRealResponse
	err = response.ParseWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	return &EcsStartInstanceResponse{
		JobId: realResponse.JobID,
	}, nil
}

type ecsStartInstanceRealRequest struct {
	RegionID   string `json:"regionID"`
	InstanceID string `json:"instanceID"`
}

type EcsStartInstanceRequest struct {
	RegionId   string
	InstanceId string
}

type EcsStartInstanceRealResponse struct {
	JobID string `json:"jobID"`
}

type EcsStartInstanceResponse struct {
	JobId string
}
