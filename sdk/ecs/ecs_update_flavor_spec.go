package ecs

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// ecsUpdateFlavorSpecApi  云主机修改规格
// https://www.ctyun.cn/document/10026730/10106612
type ecsUpdateFlavorSpecApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEcsUpdateFlavorSpecApi(client *common.CtyunSender) common.ApiHandler[EcsUpdateFlavorSpecRequest, EcsUpdateFlavorSpecResponse] {
	return &ecsUpdateFlavorSpecApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/update-flavor-spec",
		},
	}
}

func (v *ecsUpdateFlavorSpecApi) Do(ctx context.Context, credential *common.Credential, req *EcsUpdateFlavorSpecRequest) (*EcsUpdateFlavorSpecResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&ecsUpdateFlavorSpecRealRequest{
		RegionID:        req.RegionId,
		InstanceID:      req.InstanceId,
		FlavorID:        req.FlavorId,
		ClientToken:     req.ClientToken,
		PayVoucherPrice: req.PayVoucherPrice,
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
	var realResponse ecsUpdateFlavorSpecRealResponse
	err = response.ParseWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	return &EcsUpdateFlavorSpecResponse{
		MasterOrderNo: realResponse.MasterOrderNO,
		RegionId:      realResponse.RegionID,
		MasterOrderId: realResponse.MasterOrderID,
	}, nil
}

type ecsUpdateFlavorSpecRealRequest struct {
	RegionID        string `json:"regionID"`
	InstanceID      string `json:"instanceID"`
	FlavorID        string `json:"flavorID"`
	ClientToken     string `json:"clientToken"`
	PayVoucherPrice string `json:"payVoucherPrice"`
}

type EcsUpdateFlavorSpecRequest struct {
	RegionId        string
	InstanceId      string
	FlavorId        string
	ClientToken     string
	PayVoucherPrice string
}

type ecsUpdateFlavorSpecRealResponse struct {
	MasterOrderNO string `json:"masterOrderNO"`
	RegionID      string `json:"regionID"`
	MasterOrderID string `json:"masterOrderID"`
}

type EcsUpdateFlavorSpecResponse struct {
	MasterOrderNo string
	RegionId      string
	MasterOrderId string
}
