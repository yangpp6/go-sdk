package idp

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

type idpDeleteApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewIdpDeleteApi(client *common.CtyunSender) common.ApiHandler[IdpDeleteRequest, IdpDeleteResponse] {
	return &idpDeleteApi{
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v1/identityProvider/deleteIdP",
		},
		client: client,
	}

}

type IdpDeleteRequest struct {
	Id int64 `json:"id"`
}

type IdpDeleteResponse struct {
}

func (receiver idpDeleteApi) Do(ctx context.Context, credential *common.Credential, req *IdpDeleteRequest) (*IdpDeleteResponse, common.CtyunRequestError) {
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
	return &IdpDeleteResponse{}, err
}
