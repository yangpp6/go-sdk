package ecs

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// ecsQueryAsyncResultApi  关闭一台云主机
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=5543&data=87
type ecsQueryAsyncResultApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEcsQueryAsyncResultApi(client *common.CtyunSender) common.ApiHandler[EcsQueryAsyncResultRequest, EcsQueryAsyncResultResponse] {
	return &ecsQueryAsyncResultApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/query-async-result",
		},
	}
}

func (v *ecsQueryAsyncResultApi) Do(ctx context.Context, credential *common.Credential, req *EcsQueryAsyncResultRequest) (*EcsQueryAsyncResultResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&ecsQueryAsyncResultRealRequest{
		RegionID: req.RegionId,
		JobID:    req.JobId,
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
	var realResponse EcsQueryAsyncResultRealResponse
	err = response.ParseWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	return &EcsQueryAsyncResultResponse{
		JobStatus: realResponse.JobStatus,
	}, nil
}

type ecsQueryAsyncResultRealRequest struct {
	RegionID string `json:"regionID"`
	JobID    string `json:"jobID"`
}

type EcsQueryAsyncResultRequest struct {
	RegionId string
	JobId    string
}

type EcsQueryAsyncResultRealResponse struct {
	JobStatus int `json:"jobStatus"`
}

type EcsQueryAsyncResultResponse struct {
	JobStatus int
}
