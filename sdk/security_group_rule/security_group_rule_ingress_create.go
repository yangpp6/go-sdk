package security_group_rule

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// securityGroupRuleIngressCreateApi 创建安全组入向规则
type securityGroupRuleIngressCreateApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewSecurityGroupRuleIngressCreateApi(client *common.CtyunSender) common.ApiHandler[SecurityGroupRuleIngressCreateRequest, SecurityGroupRuleIngressCreateResponse] {
	return &securityGroupRuleIngressCreateApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/vpc/create-security-group-ingress",
		},
	}
}

func (v *securityGroupRuleIngressCreateApi) Do(ctx context.Context, credential *common.Credential, req *SecurityGroupRuleIngressCreateRequest) (*SecurityGroupRuleIngressCreateResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	var securityGroupRules []securityGroupRuleIngressCreateSecurityGroupRulesRealRequest
	for _, rule := range req.SecurityGroupRules {
		securityGroupRules = append(securityGroupRules, securityGroupRuleIngressCreateSecurityGroupRulesRealRequest{
			Direction:   rule.Direction,
			Action:      rule.Action,
			Priority:    rule.Priority,
			Protocol:    rule.Protocol,
			Ethertype:   rule.Ethertype,
			DestCidrIp:  rule.DestCidrIp,
			Description: rule.Description,
			Range:       rule.Range,
		})
	}
	_, err := builder.WriteJson(&securityGroupRuleIngressCreateRealRequest{
		RegionID:           req.RegionId,
		SecurityGroupID:    req.SecurityGroupId,
		ClientToken:        req.ClientToken,
		SecurityGroupRules: securityGroupRules,
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
	var realResponse securityGroupRuleIngressCreateRealResponse
	err = response.ParseWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	return &SecurityGroupRuleIngressCreateResponse{
		SgRuleIds: realResponse.SgRuleIDs,
	}, nil
}

type securityGroupRuleIngressCreateSecurityGroupRulesRealRequest struct {
	Direction   string `json:"direction"`
	Action      string `json:"action"`
	Priority    int    `json:"priority"`
	Protocol    string `json:"protocol"`
	Ethertype   string `json:"ethertype"`
	DestCidrIp  string `json:"destCidrIp"`
	Description string `json:"description"`
	Range       string `json:"range"`
}

type securityGroupRuleIngressCreateRealRequest struct {
	RegionID           string                                                        `json:"regionID"`
	SecurityGroupID    string                                                        `json:"securityGroupID"`
	ClientToken        string                                                        `json:"clientToken"`
	SecurityGroupRules []securityGroupRuleIngressCreateSecurityGroupRulesRealRequest `json:"securityGroupRules"`
}

type securityGroupRuleIngressCreateRealResponse struct {
	SgRuleIDs []string `json:"sgRuleIDs"`
}

type SecurityGroupRuleIngressCreateSecurityGroupRulesRequest struct {
	Direction   string // 规则方向，出方向则填写ingress
	Action      string // 授权策略，取值范围：accept（允许），drop（拒绝）。
	Priority    int    // 优先级，取值范围：[1, 100]，取值越小优先级越大
	Protocol    string // 网络协议，取值范围：ANY（任意）、TCP、UDP、ICMP(v4)
	Ethertype   string // IP类型，取值范围：IPv4、IPv6
	DestCidrIp  string // 远端地址:0.0.0.0/0
	Description string // 安全组规则描述信息，满足以下规则： ① 长度0-128字符， ② 支持拉丁字母、中文、数字, 特殊字符 ！@#￥%……&*（）——-+={}《》？：“”【】、；‘'，。、 不能以      http: / https: 开头
	Range       string // 安全组开放的传输层协议相关的源端端口范围
}

type SecurityGroupRuleIngressCreateRequest struct {
	RegionId           string                                                    // 区域id
	SecurityGroupId    string                                                    // 安全组ID。
	ClientToken        string                                                    // 客户端存根
	SecurityGroupRules []SecurityGroupRuleIngressCreateSecurityGroupRulesRequest // 规则信息
}

type SecurityGroupRuleIngressCreateResponse struct {
	SgRuleIds []string
}
