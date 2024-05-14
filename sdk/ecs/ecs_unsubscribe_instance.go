package ecs

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// ecsUnsubscribeInstanceApi  退订云主机
// https://www.ctyun.cn/document/10026730/10106596
type ecsUnsubscribeInstanceApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEcsUnsubscribeInstanceApi(client *common.CtyunSender) common.ApiHandler[EcsUnsubscribeInstanceRequest, EcsUnsubscribeInstanceResponse] {
	return &ecsUnsubscribeInstanceApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/unsubscribe-instance",
		},
	}
}

func (v *ecsUnsubscribeInstanceApi) Do(ctx context.Context, credential *common.Credential, req *EcsUnsubscribeInstanceRequest) (*EcsUnsubscribeInstanceResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&ecsUnsubscribeInstanceRealRequest{
		ClientToken: req.ClientToken,
		RegionID:    req.RegionId,
		InstanceID:  req.InstanceId,
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
	var realResponse ecsUnsubscribeInstanceRealResponse
	err = response.ParseWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	return &EcsUnsubscribeInstanceResponse{
		MasterOrderNo: realResponse.MasterOrderNO,
		RegionId:      realResponse.RegionID,
		MasterOrderId: realResponse.MasterOrderID,
	}, nil
}

type ecsUnsubscribeInstanceRealRequest struct {
	ClientToken string `json:"clientToken,omitempty"`
	RegionID    string `json:"regionID,omitempty"`
	InstanceID  string `json:"instanceID,omitempty"`
}

type EcsUnsubscribeInstanceRequest struct {
	ClientToken string
	RegionId    string
	InstanceId  string
}

type ecsUnsubscribeInstanceRealResponse struct {
	MasterOrderNO string `json:"masterOrderNO"`
	RegionID      string `json:"regionID"`
	MasterOrderID string `json:"masterOrderID"`
}

type EcsUnsubscribeInstanceResponse struct {
	MasterOrderNo string
	RegionId      string
	MasterOrderId string
}
