package ecs

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// ecsLeaveSecurityGroupApi  绑定安全组
// https://www.ctyun.cn/document/10026730/10040193
type ecsLeaveSecurityGroupApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEcsLeaveSecurityGroupApi(client *common.CtyunSender) common.ApiHandler[EcsLeaveSecurityGroupRequest, EcsLeaveSecurityGroupResponse] {
	return &ecsLeaveSecurityGroupApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/vpc/leave-security-group",
		},
	}
}

func (v *ecsLeaveSecurityGroupApi) Do(ctx context.Context, credential *common.Credential, req *EcsLeaveSecurityGroupRequest) (*EcsLeaveSecurityGroupResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&ecsLeaveSecurityGroupRealRequest{
		RegionID:        req.RegionId,
		SecurityGroupID: req.SecurityGroupId,
		InstanceID:      req.InstanceId,
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
	return &EcsLeaveSecurityGroupResponse{}, nil
}

type ecsLeaveSecurityGroupRealRequest struct {
	RegionID        string `json:"regionID,omitempty"`
	SecurityGroupID string `json:"securityGroupID,omitempty"`
	InstanceID      string `json:"instanceID,omitempty"`
}

type EcsLeaveSecurityGroupRequest struct {
	RegionId        string
	SecurityGroupId string
	InstanceId      string
}

type EcsLeaveSecurityGroupResponse struct {
}
