package security_group_rule

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// securityGroupRuleIngressModifyApi 修改安全组入向规则
type securityGroupRuleIngressModifyApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewSecurityGroupRuleIngressModifyApi(client *common.CtyunSender) common.ApiHandler[SecurityGroupRuleIngressModifyRequest, SecurityGroupRuleIngressModifyResponse] {
	return &securityGroupRuleIngressModifyApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/vpc/modify-security-group-ingress",
		},
	}
}

func (v *securityGroupRuleIngressModifyApi) Do(ctx context.Context, credential *common.Credential, req *SecurityGroupRuleIngressModifyRequest) (*SecurityGroupRuleIngressModifyResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&securityGroupRuleIngressModifyRealRequest{
		RegionID:            req.RegionId,
		SecurityGroupID:     req.SecurityGroupId,
		SecurityGroupRuleID: req.SecurityGroupRuleId,
		ClientToken:         req.ClientToken,
		Description:         req.Description,
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
	return &SecurityGroupRuleIngressModifyResponse{}, nil
}

type securityGroupRuleIngressModifyRealRequest struct {
	RegionID            string `json:"regionID"`
	SecurityGroupID     string `json:"securityGroupID"`
	SecurityGroupRuleID string `json:"securityGroupRuleID"`
	ClientToken         string `json:"clientToken"`
	Description         string `json:"description"`
}

type SecurityGroupRuleIngressModifyRequest struct {
	RegionId            string // 资源池ID，请根据查询资源池列表接口返回值进行传参，获取“regionId”参数
	SecurityGroupId     string // 安全组ID
	SecurityGroupRuleId string // 安全规则ID
	ClientToken         string // 客户端存根
	Description         string // 描述
}

type SecurityGroupRuleIngressModifyResponse struct {
}
