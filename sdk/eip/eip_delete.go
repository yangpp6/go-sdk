package eip

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// eipDeleteApi 删除弹性IP
type eipDeleteApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEipDeleteApi(client *common.CtyunSender) common.ApiHandler[EipDeleteRequest, EipDeleteResponse] {
	return &eipDeleteApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/eip/delete",
		},
	}
}

func (v *eipDeleteApi) Do(ctx context.Context, credential *common.Credential, req *EipDeleteRequest) (*EipDeleteResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	requestContent := eipDeleteRealRequest{
		ClientToken: req.ClientToken,
		RegionID:    req.RegionId,
		EipID:       req.EipId,
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
	result := &eipDeleteRealResponse{}
	err = response.ParseWithCheck(result)
	if err != nil {
		return nil, err
	}

	return &EipDeleteResponse{
		MasterOrderId: result.MasterOrderID,
		MasterOrderNo: result.MasterOrderNO,
		RegionId:      result.RegionID,
	}, nil
}

type eipDeleteRealRequest struct {
	ClientToken string `json:"clientToken"`
	RegionID    string `json:"regionID"`
	EipID       string `json:"eipID"`
}

type eipDeleteRealResponse struct {
	MasterOrderID string `json:"masterOrderID"`
	MasterOrderNO string `json:"masterOrderNO"`
	RegionID      string `json:"regionID"`
}

type EipDeleteRequest struct {
	ClientToken string
	RegionId    string
	EipId       string
}

type EipDeleteResponse struct {
	MasterOrderId string
	MasterOrderNo string
	RegionId      string
}
