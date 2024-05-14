package user

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

type userListApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewUserListApi(client *common.CtyunSender) common.ApiHandler[UserListRequest, UserListResponse] {
	return &userListApi{
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v1/user/getUser",
		},
		client: client,
	}
}

type UserListRequest struct {
	UserId string
}

func (i *userListApi) Do(ctx context.Context, credential *common.Credential, r *UserListRequest) (*UserListResponse, common.CtyunRequestError) {
	request := i.WithCredential(credential).AddParam("userId", r.UserId)
	ctiam, err := i.client.SendCtiam(ctx, request)
	if err != nil {
		return nil, err
	}
	resp := &UserListResponse{}
	err = ctiam.ParseWithCheck(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type UserListResponse struct {
	LoginEmail  string      `json:"loginEmail"`
	AccountId   string      `json:"accountId"`
	MobilePhone string      `json:"mobilePhone"`
	Groups      []UserGroup `json:"groups"`
	Remark      string      `json:"remark"`
	UserName    string      `json:"userName"`
}
