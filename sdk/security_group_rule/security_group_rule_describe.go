package security_group_rule

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// securityGroupRuleDescribeApi 创建安全组出向规则
type securityGroupRuleDescribeApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewSecurityGroupRuleDescribeApi(client *common.CtyunSender) common.ApiHandler[SecurityGroupRuleDescribeRequest, SecurityGroupRuleDescribeResponse] {
	return &securityGroupRuleDescribeApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/vpc/describe-security-group-rule",
		},
	}
}

func (v *securityGroupRuleDescribeApi) Do(ctx context.Context, credential *common.Credential, req *SecurityGroupRuleDescribeRequest) (*SecurityGroupRuleDescribeResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	builder.AddParam("regionID", req.RegionId).
		AddParam("securityGroupID", req.SecurityGroupId).
		AddParam("securityGroupRuleID", req.SecurityGroupRuleId)

	// 发起请求
	response, err := v.client.SendCtVpc(ctx, builder)
	if err != nil {
		return nil, err
	}

	// 解析返回
	var realResponse securityGroupRuleDescribeRealResponse
	err = response.ParseWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	return &SecurityGroupRuleDescribeResponse{
		Direction:       realResponse.Direction,
		Priority:        realResponse.Priority,
		Ethertype:       realResponse.Ethertype,
		Protocol:        realResponse.Protocol,
		Range:           realResponse.Range,
		DestCidrIp:      realResponse.DestCidrIp,
		Description:     realResponse.Description,
		Id:              realResponse.Id,
		SecurityGroupId: realResponse.SecurityGroupID,
		Action:          realResponse.Action,
	}, nil
}

type securityGroupRuleDescribeRealResponse struct {
	Direction       string `json:"direction"`
	Priority        int    `json:"priority"`
	Ethertype       string `json:"ethertype"`
	Protocol        string `json:"protocol"`
	Range           string `json:"range"`
	DestCidrIp      string `json:"destCidrIp"`
	Description     string `json:"description"`
	Id              string `json:"id"`
	SecurityGroupID string `json:"securityGroupID"`
	Action          string `json:"action"`
}

type SecurityGroupRuleDescribeRequest struct {
	RegionId            string // 区域id
	SecurityGroupId     string // 安全组ID
	SecurityGroupRuleId string // 安全组规则id
}

type SecurityGroupRuleDescribeResponse struct {
	Direction       string // 规则方向，出方向则填写egress
	Priority        int    // 优先级，取值范围：[1, 100]，取值越小优先级越大
	Ethertype       string // IP类型，取值范围：IPv4、IPv6
	Protocol        string // 网络协议，取值范围：ANY（任意）、TCP、UDP、ICMP(v4)
	Range           string // 安全组开放的传输层协议相关的源端端口范围
	DestCidrIp      string // 远端地址:0.0.0.0/0
	Description     string // 安全组规则描述信息，满足以下规则： ① 长度0-128字符， ② 支持拉丁字母、中文、数字, 特殊字符 ！@#￥%……&*（）——-+={}《》？：“”【】、；‘'，。、 不能以      http: / https: 开头
	Id              string // id
	SecurityGroupId string // 安全组id
	Action          string // 授权策略，取值范围：accept（允许），drop（拒绝）。
}
