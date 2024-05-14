package security_group_rule

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// securityGroupRuleEgressRevokeApi 删除安全组出向规则
type securityGroupRuleEgressRevokeApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewSecurityGroupRuleEgressRevokeApi(client *common.CtyunSender) common.ApiHandler[SecurityGroupRuleEgressRevokeRequest, SecurityGroupRuleEgressRevokeResponse] {
	return &securityGroupRuleEgressRevokeApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/vpc/revoke-security-group-egress",
		},
	}
}

func (v *securityGroupRuleEgressRevokeApi) Do(ctx context.Context, credential *common.Credential, req *SecurityGroupRuleEgressRevokeRequest) (*SecurityGroupRuleEgressRevokeResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&securityGroupRuleEgressRevokeRealRequest{
		RegionID:            req.RegionId,
		SecurityGroupID:     req.SecurityGroupId,
		SecurityGroupRuleID: req.SecurityGroupRuleId,
		ClientToken:         req.ClientToken,
	})
	if err != nil {
		return nil, err
	}

	// 发起请求
	response, err := v.client.SendCtVpc(ctx, builder)
	if err != nil {
		return nil, err
	}

	// 解析返回
	err = response.ParseWithCheck(nil)
	if err != nil {
		return nil, err
	}
	return &SecurityGroupRuleEgressRevokeResponse{}, nil
}

type securityGroupRuleEgressRevokeRealRequest struct {
	RegionID            string `json:"regionID"`
	SecurityGroupID     string `json:"securityGroupID"`
	SecurityGroupRuleID string `json:"securityGroupRuleID"`
	ClientToken         string `json:"clientToken"`
}

type SecurityGroupRuleEgressRevokeRequest struct {
	RegionId            string // 资源池ID，请根据查询资源池列表接口返回值进行传参，获取“regionId”参数
	SecurityGroupId     string // 安全组ID
	SecurityGroupRuleId string // 安全规则ID
	ClientToken         string // 客户端存根
}

type SecurityGroupRuleEgressRevokeResponse struct {
}
