package eip

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// eipModifySpecApi 修改弹性IP规格
type eipModifySpecApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEipModifySpecApi(client *common.CtyunSender) common.ApiHandler[EipModifySpecRequest, EipModifySpecResponse] {
	return &eipModifySpecApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/eip/modify-spec",
		},
	}
}

func (v *eipModifySpecApi) Do(ctx context.Context, credential *common.Credential, req *EipModifySpecRequest) (*EipModifySpecResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	requestContent := eipModifySpecRealRequest{
		ClientToken: req.ClientToken,
		RegionID:    req.RegionId,
		EipID:       req.EipId,
		Bandwidth:   req.Bandwidth,
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
	result := &eipModifySpecRealResponse{}
	err = response.ParseWithCheck(result)
	if err != nil {
		return nil, err
	}

	return &EipModifySpecResponse{
		MasterOrderId: result.MasterOrderID,
		MasterOrderNo: result.MasterOrderNO,
		RegionId:      result.RegionID,
	}, nil
}

type eipModifySpecRealRequest struct {
	ClientToken string `json:"clientToken"`
	RegionID    string `json:"regionID"`
	EipID       string `json:"eipID"`
	Bandwidth   int    `json:"bandwidth"`
}

type eipModifySpecRealResponse struct {
	MasterOrderID string `json:"masterOrderID"`
	MasterOrderNO string `json:"masterOrderNO"`
	RegionID      string `json:"regionID"`
}

type EipModifySpecRequest struct {
	ClientToken string
	RegionId    string
	EipId       string
	Bandwidth   int
}

type EipModifySpecResponse struct {
	MasterOrderId string
	MasterOrderNo string
	RegionId      string
}
