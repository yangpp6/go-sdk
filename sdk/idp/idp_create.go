package idp

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

type idpCreateApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewIdpCreateApi(client *common.CtyunSender) common.ApiHandler[IdpCreateRequest, IdpCreateResponse] {
	return &idpCreateApi{
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v1/identityProvider/createIdP",
		},
		client: client,
	}

}

type IdpCreateRequest struct {
	Name     string `json:"name"`
	Type     int64  `json:"type"`
	Protocol int64  `json:"protocol"`
	Remark   string `json:"remark"`
	FileName string `json:"fileName"`
	File     []byte `json:"file"`
}

type IdpCreateResponse struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Type       int    `json:"type"`
	Protocol   int    `json:"protocol"`
	AccountId  string `json:"accountId"`
	Remark     string `json:"remark"`
	CreateTime int64  `json:"createTime"`
	Uuid       string `json:"uuid"`
}

func (receiver idpCreateApi) Do(ctx context.Context, credential *common.Credential, req *IdpCreateRequest) (*IdpCreateResponse, common.CtyunRequestError) {
	builder := receiver.CtyunRequestBuilder.WithCredential(credential)
	builder, err := builder.WriteJson(req)
	if err != nil {
		return nil, err
	}
	send, err := receiver.client.SendCtiam(ctx, builder)
	if err != nil {
		return nil, err
	}
	response := &IdpCreateResponse{}
	err = send.ParseWithCheck(response)
	return response, err
}
