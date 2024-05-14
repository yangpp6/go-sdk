package user

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

type userGroupCreateApi struct {
	builder common.CtyunRequestBuilder
	client  *common.CtyunSender
}

func NewUserGroupCreateApi(client *common.CtyunSender) common.ApiHandler[UserGroupCreateRequest, UserGroupCreateResponse] {
	return &userGroupCreateApi{
		builder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v1/userGroup/createGroup",
		},
		client: client,
	}
}

type userGroupCreateRealRequest struct {
	GroupName  string `json:"groupName"`
	GroupIntro string `json:"groupIntro"`
}

type userGroupRealResponse struct {
	Id string `json:"id"`
}

type UserGroupCreateRequest struct {
	GroupName  string
	GroupIntro string
}

type UserGroupCreateResponse struct {
	Id string
}

func (receiver *userGroupCreateApi) Do(ctx context.Context, credential *common.Credential, req *UserGroupCreateRequest) (*UserGroupCreateResponse, common.CtyunRequestError) {
	builder := receiver.builder.WithCredential(credential)
	builder, err := builder.WriteJson(&userGroupCreateRealRequest{
		GroupName:  req.GroupName,
		GroupIntro: req.GroupIntro,
	})
	if err != nil {
		return nil, err
	}
	send, err := receiver.client.SendCtiam(ctx, builder)
	if err != nil {
		return nil, err
	}
	resp := &userGroupRealResponse{}
	err = send.ParseWithCheck(resp)
	if err != nil {
		return nil, err
	}
	return &UserGroupCreateResponse{
		Id: resp.Id,
	}, nil
}
