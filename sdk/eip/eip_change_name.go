package eip

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// eipChangeNameApi 修改弹性IP名称
type eipChangeNameApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEipChangeNameApi(client *common.CtyunSender) common.ApiHandler[EipChangeNameRequest, EipChangeNameResponse] {
	return &eipChangeNameApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/eip/change-name",
		},
	}
}

func (v *eipChangeNameApi) Do(ctx context.Context, credential *common.Credential, req *EipChangeNameRequest) (*EipChangeNameResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	requestContent := eipChangeNameRealRequest{
		ClientToken: req.ClientToken,
		RegionID:    req.RegionId,
		EipID:       req.EipId,
		Name:        req.Name,
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

	return &EipChangeNameResponse{}, nil
}

type eipChangeNameRealRequest struct {
	ClientToken string `json:"clientToken"`
	RegionID    string `json:"regionID"`
	EipID       string `json:"eipID"`
	Name        string `json:"name"`
}

type EipChangeNameRequest struct {
	ClientToken string
	RegionId    string
	EipId       string
	Name        string
}

type EipChangeNameResponse struct {
}
