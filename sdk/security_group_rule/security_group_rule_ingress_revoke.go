package security_group_rule

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// securityGroupRuleIngressRevokeApi 删除安全组入向规则
type securityGroupRuleIngressRevokeApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewSecurityGroupRuleIngressRevokeApi(client *common.CtyunSender) common.ApiHandler[SecurityGroupRuleIngressRevokeRequest, SecurityGroupRuleIngressRevokeResponse] {
	return &securityGroupRuleIngressRevokeApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/vpc/revoke-security-group-ingress",
		},
	}
}

func (v *securityGroupRuleIngressRevokeApi) Do(ctx context.Context, credential *common.Credential, req *SecurityGroupRuleIngressRevokeRequest) (*SecurityGroupRuleIngressRevokeResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&securityGroupRuleIngressRevokeRealRequest{
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
	return &SecurityGroupRuleIngressRevokeResponse{}, nil
}

type securityGroupRuleIngressRevokeRealRequest struct {
	RegionID            string `json:"regionID"`
	SecurityGroupID     string `json:"securityGroupID"`
	SecurityGroupRuleID string `json:"securityGroupRuleID"`
	ClientToken         string `json:"clientToken"`
}

type SecurityGroupRuleIngressRevokeRequest struct {
	RegionId            string // 资源池ID，请根据查询资源池列表接口返回值进行传参，获取“regionId”参数
	SecurityGroupId     string // 安全组ID
	SecurityGroupRuleId string // 安全规则ID
	ClientToken         string // 客户端存根
}

type SecurityGroupRuleIngressRevokeResponse struct {
}
