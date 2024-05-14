package eip

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// eipAssociateApi 绑定弹性IP
type eipAssociateApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEipAssociateApi(client *common.CtyunSender) common.ApiHandler[EipAssociateRequest, EipAssociateResponse] {
	return &eipAssociateApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/eip/associate",
		},
	}
}

func (v *eipAssociateApi) Do(ctx context.Context, credential *common.Credential, req *EipAssociateRequest) (*EipAssociateResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	requestContent := eipAssociateRealRequest{
		RegionID:        req.RegionId,
		ClientToken:     req.ClientToken,
		EipID:           req.EipId,
		AssociationID:   req.AssociationId,
		AssociationType: req.AssociationType,
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
	err = response.ParseWithCheck(nil)
	if err != nil {
		return nil, err
	}

	return &EipAssociateResponse{}, nil
}

type eipAssociateRealRequest struct {
	RegionID        string `json:"regionID"`
	ClientToken     string `json:"clientToken"`
	EipID           string `json:"eipID"`
	AssociationID   string `json:"associationID"`
	AssociationType int    `json:"associationType"`
}

type EipAssociateRequest struct {
	RegionId        string
	ClientToken     string
	EipId           string
	AssociationId   string
	AssociationType int
}

type EipAssociateResponse struct {
}
