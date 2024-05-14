package eip

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// eipCreateApi 创建弹性IP
type eipCreateApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEipCreateApi(client *common.CtyunSender) common.ApiHandler[EipCreateRequest, EipCreateResponse] {
	return &eipCreateApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/eip/create",
		},
	}
}

func (v *eipCreateApi) Do(ctx context.Context, credential *common.Credential, req *EipCreateRequest) (*EipCreateResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	requestContent := eipCreateRealRequest{
		ClientToken:       req.ClientToken,
		RegionID:          req.RegionId,
		CycleType:         req.CycleType,
		CycleCount:        req.CycleCount,
		Name:              req.Name,
		Bandwidth:         req.Bandwidth,
		BandwidthID:       req.BandwidthId,
		DemandBillingType: req.DemandBillingType,
	}
	_, err := builder.WriteJson(requestContent)
	if err != nil {
		return nil, err
	}

	// 发起请求
	response, err := v.client.SendCtVpc(ctx, builder)
	if err != nil {
		return nil, err
	}

	// 解析返回
	result := &eipCreateRealResponse{}
	err = response.ParseWithCheck(result)
	if err != nil {
		return nil, err
	}

	return &EipCreateResponse{
		MasterOrderId:        result.MasterOrderID,
		MasterOrderNo:        result.MasterOrderNO,
		MasterResourceId:     result.MasterResourceID,
		MasterResourceStatus: result.MasterResourceStatus,
		RegionId:             result.RegionID,
		EipId:                result.EipID,
	}, nil
}

type eipCreateRealRequest struct {
	ClientToken       string `json:"clientToken"`
	RegionID          string `json:"regionID"`
	CycleType         string `json:"cycleType"`
	CycleCount        int    `json:"cycleCount,omitempty"`
	Name              string `json:"name"`
	Bandwidth         int    `json:"bandwidth"`
	BandwidthID       string `json:"bandwidthID"`
	DemandBillingType string `json:"demandBillingType"`
}

type eipCreateRealResponse struct {
	MasterOrderID        string `json:"masterOrderID"`
	MasterOrderNO        string `json:"masterOrderNO"`
	MasterResourceID     string `json:"masterResourceID"`
	MasterResourceStatus string `json:"masterResourceStatus"`
	RegionID             string `json:"regionID"`
	EipID                string `json:"eipID"`
}

type EipCreateRequest struct {
	ClientToken       string
	RegionId          string
	CycleType         string
	CycleCount        int
	Name              string
	Bandwidth         int
	BandwidthId       string
	DemandBillingType string
}

type EipCreateResponse struct {
	MasterOrderId        string
	MasterOrderNo        string
	MasterResourceId     string
	MasterResourceStatus string
	RegionId             string
	EipId                string
}
