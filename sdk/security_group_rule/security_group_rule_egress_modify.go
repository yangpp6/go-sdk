package security_group_rule

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// securityGroupRuleEgressModifyApi 修改安全组出向规则
type securityGroupRuleEgressModifyApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewSecurityGroupRuleEgressModifyApi(client *common.CtyunSender) common.ApiHandler[SecurityGroupRuleEgressModifyRequest, SecurityGroupRuleEgressModifyResponse] {
	return &securityGroupRuleEgressModifyApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/vpc/modify-security-group-egress",
		},
	}
}

func (v *securityGroupRuleEgressModifyApi) Do(ctx context.Context, credential *common.Credential, req *SecurityGroupRuleEgressModifyRequest) (*SecurityGroupRuleEgressModifyResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&securityGroupRuleEgressModifyRealRequest{
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
	return &SecurityGroupRuleEgressModifyResponse{}, nil
}

type securityGroupRuleEgressModifyRealRequest struct {
	RegionID            string `json:"regionID"`
	SecurityGroupID     string `json:"securityGroupID"`
	SecurityGroupRuleID string `json:"securityGroupRuleID"`
	ClientToken         string `json:"clientToken"`
	Description         string `json:"description"`
}

type SecurityGroupRuleEgressModifyRequest struct {
	RegionId            string // 资源池ID，请根据查询资源池列表接口返回值进行传参，获取“regionId”参数
	SecurityGroupId     string // 安全组ID
	SecurityGroupRuleId string // 安全规则ID
	ClientToken         string // 客户端存根
	Description         string // 描述
}

type SecurityGroupRuleEgressModifyResponse struct {
}
