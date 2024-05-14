package user

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

type userAssociationGroupApi struct {
	builder common.CtyunRequestBuilder
	client  *common.CtyunSender
}

func NewUserAssociationGroupApi(client *common.CtyunSender) common.ApiHandler[UserAssociationGroupRequest, UserAssociationGroupResponse] {
	return &userAssociationGroupApi{
		builder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v1/userGroup/userToGroup",
		},
		client: client,
	}
}

type userId struct {
	id string `json:"id"`
}

type userAssociationGroupRealRequest struct {
	GroupId string   `json:"groupId"`
	UserIds []userId `json:"userIds"`
}

type userAssociationGroupRealResponse struct {
}

type UserAssociationGroupRequest struct {
	GroupId string
	UserId  string
}

type UserAssociationGroupResponse struct {
}

func (receiver *userAssociationGroupApi) Do(ctx context.Context, credential *common.Credential, req *UserAssociationGroupRequest) (*UserAssociationGroupResponse, common.CtyunRequestError) {
	builder := receiver.builder.WithCredential(credential)
	builder, err := builder.WriteJson(&userAssociationGroupRealRequest{
		GroupId: req.GroupId,
		UserIds: []userId{{id: req.UserId}},
	})
	if err != nil {
		return nil, err
	}
	send, err := receiver.client.SendCtiam(ctx, builder)
	if err != nil {
		return nil, err
	}
	resp := &userAssociationGroupRealResponse{}
	err = send.ParseWithCheck(resp)
	if err != nil {
		return nil, err
	}
	return &UserAssociationGroupResponse{}, nil
}
