package eip

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// eipDisassociateApi 解绑弹性IP
type eipDisassociateApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEipDisassociateApi(client *common.CtyunSender) common.ApiHandler[EipDisassociateRequest, EipDisassociateResponse] {
	return &eipDisassociateApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/eip/disassociate",
		},
	}
}

func (v *eipDisassociateApi) Do(ctx context.Context, credential *common.Credential, req *EipDisassociateRequest) (*EipDisassociateResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	requestContent := eipDisassociateRealRequest{
		RegionID:    req.RegionId,
		ClientToken: req.ClientToken,
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
	err = response.ParseWithCheck(nil)
	if err != nil {
		return nil, err
	}

	return &EipDisassociateResponse{}, nil
}

type eipDisassociateRealRequest struct {
	RegionID    string `json:"regionID"`
	ClientToken string `json:"clientToken"`
	EipID       string `json:"eipID"`
}

type EipDisassociateRequest struct {
	RegionId    string
	ClientToken string
	EipId       string
}

type EipDisassociateResponse struct {
}
