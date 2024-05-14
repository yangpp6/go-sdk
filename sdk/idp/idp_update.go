package idp

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

type idpUpdateApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewIdpUpdateApi(client *common.CtyunSender) common.ApiHandler[IdpUpdateRequest, IdpUpdateResponse] {
	return &idpUpdateApi{
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v1/identityProvider/updateIdP",
		},
		client: client,
	}

}

type IdpUpdateRequest struct {
	Id       int64  `json:"id"`
	Remark   string `json:"remark"`
	FileName string `json:"fileName"`
	File     []byte `json:"file"`
}

type IdpUpdateResponse struct {
}

func (receiver idpUpdateApi) Do(ctx context.Context, credential *common.Credential, req *IdpUpdateRequest) (*IdpUpdateResponse, common.CtyunRequestError) {
	builder := receiver.CtyunRequestBuilder.WithCredential(credential)
	builder, err := builder.WriteJson(req)
	if err != nil {
		return nil, err
	}
	send, err := receiver.client.SendCtiam(ctx, builder)
	if err != nil {
		return nil, err
	}
	err = send.ParseWithCheck(nil)
	return &IdpUpdateResponse{}, err
}
