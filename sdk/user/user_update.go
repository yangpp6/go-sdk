package user

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

type userUpdateApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewUserUpdateApi(client *common.CtyunSender) common.ApiHandler[UserUpdateRequest, UserUpdateResponse] {
	return &userUpdateApi{
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v1/user/updateUser",
		},
		client: client,
	}
}

type UserUpdateRequest struct {
	UserId      string `json:"userId"`
	Remark      string `json:"remark"`
	LoginEmail  string `json:"loginEmail"`
	MobilePhone string `json:"mobilePhone"`
	UserName    string `json:"userName"`
	Prohibit    int    `json:"prohibit"`
}

type UserUpdateResponse struct {
	LoginEmail  string `json:"loginEmail"`
	AccountId   string `json:"accountId"`
	MobilePhone string `json:"mobilePhone"`
	Remark      string `json:"remark"`
	UserName    string `json:"userName"`
}

func (receiver userUpdateApi) Do(ctx context.Context, credential *common.Credential, req *UserUpdateRequest) (*UserUpdateResponse, common.CtyunRequestError) {
	builder := receiver.CtyunRequestBuilder.WithCredential(credential)
	builder, err := builder.WriteJson(req)
	if err != nil {
		return nil, err
	}
	send, err := receiver.client.SendCtiam(ctx, builder)
	if err != nil {
		return nil, err
	}
	resp := &UserUpdateResponse{}
	err = send.ParseWithCheck(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
