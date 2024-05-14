package ecs

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// ecsJoinSecurityGroupApi  绑定安全组
// https://www.ctyun.cn/document/10026730/10040193
type ecsJoinSecurityGroupApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEcsJoinSecurityGroupApi(client *common.CtyunSender) common.ApiHandler[EcsJoinSecurityGroupRequest, EcsJoinSecurityGroupResponse] {
	return &ecsJoinSecurityGroupApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/vpc/join-security-group",
		},
	}
}

func (v *ecsJoinSecurityGroupApi) Do(ctx context.Context, credential *common.Credential, req *EcsJoinSecurityGroupRequest) (*EcsJoinSecurityGroupResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&ecsJoinSecurityGroupRealRequest{
		RegionID:           req.RegionId,
		SecurityGroupID:    req.SecurityGroupId,
		InstanceID:         req.InstanceId,
		NetworkInterfaceID: req.NetworkInterfaceId,
		Action:             req.Action,
	})
	if err != nil {
		return nil, err
	}

	// 发起请求
	response, err := v.client.SendCtEcs(ctx, builder)
	if err != nil {
		return nil, err
	}

	// 解析返回
	err = response.ParseWithCheck(nil)
	if err != nil {
		return nil, err
	}
	return &EcsJoinSecurityGroupResponse{}, nil
}

type ecsJoinSecurityGroupRealRequest struct {
	RegionID           string `json:"regionID,omitempty"`
	SecurityGroupID    string `json:"securityGroupID,omitempty"`
	InstanceID         string `json:"instanceID,omitempty"`
	NetworkInterfaceID string `json:"networkInterfaceID,omitempty"`
	Action             string `json:"action,omitempty"`
}

type EcsJoinSecurityGroupRequest struct {
	RegionId           string
	SecurityGroupId    string
	InstanceId         string
	NetworkInterfaceId string
	Action             string
}

type EcsJoinSecurityGroupResponse struct {
}
