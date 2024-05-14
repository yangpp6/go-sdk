package user

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

type userInvalidApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewUserInvalidApi(client *common.CtyunSender) common.ApiHandler[UserInvalidRequest, UserInvalidResponse] {
	return &userInvalidApi{
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v1/user/invalidUser",
		},
		client: client,
	}
}

type UserInvalidRequest struct {
	UserId string `json:"userId"`
}

type UserInvalidResponse struct {
}

func (i *userInvalidApi) Do(ctx context.Context, credential *common.Credential, t *UserInvalidRequest) (*UserInvalidResponse, common.CtyunRequestError) {
	builder := i.CtyunRequestBuilder.WithCredential(credential)
	json, err := builder.WriteJson(t)
	if err != nil {
		return nil, err
	}
	ctiam, err := i.client.SendCtiam(ctx, json)
	if err != nil {
		return nil, err
	}
	resp := &UserInvalidResponse{}
	err = ctiam.ParseWithCheck(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
