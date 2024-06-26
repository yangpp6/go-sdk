package security_group

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// securityGroupDeleteApi 查询用户安全组详情
type securityGroupDescribeAttributeApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewSecurityGroupDescribeAttributeApi(client *common.CtyunSender) common.ApiHandler[SecurityGroupDescribeAttributeRequest, SecurityGroupDescribeAttributeResponse] {
	return &securityGroupDescribeAttributeApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/vpc/describe-security-group-attribute",
		},
	}
}

func (v *securityGroupDescribeAttributeApi) Do(ctx context.Context, credential *common.Credential, req *SecurityGroupDescribeAttributeRequest) (*SecurityGroupDescribeAttributeResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	builder.
		AddParam("regionID", req.RegionId).
		AddParam("securityGroupID", req.SecurityGroupId).
		AddParam("projectID", req.ProjectId).
		AddParam("direction", req.Direction)

	// 发起请求
	response, err := v.client.SendCtVpc(ctx, builder)
	if err != nil {
		return nil, err
	}

	// 解析返回
	var realResponse securityGroupDescribeAttributeRealResponse
	err = response.ParseWithCheck(&realResponse)

	var sgrl []SecurityGroupDescribeAttributeSecurityGroupRuleListResponse
	for _, sg := range realResponse.SecurityGroupRuleList {
		sgrl = append(sgrl, SecurityGroupDescribeAttributeSecurityGroupRuleListResponse{
			Direction:       sg.Direction,
			Priority:        sg.Priority,
			Ethertype:       sg.Ethertype,
			Protocol:        sg.Protocol,
			Range:           sg.Range,
			DestCidrIp:      sg.DestCidrIp,
			Description:     sg.Description,
			Origin:          sg.Origin,
			CreateTime:      sg.CreateTime,
			Id:              sg.Id,
			Action:          sg.Action,
			SecurityGroupId: sg.SecurityGroupId,
		})
	}
	return &SecurityGroupDescribeAttributeResponse{
		SecurityGroupName:     realResponse.SecurityGroupName,
		Id:                    realResponse.Id,
		VmNum:                 realResponse.VmNum,
		Origin:                realResponse.Origin,
		VpcName:               realResponse.VpcName,
		VpcId:                 realResponse.VpcId,
		CreationTime:          realResponse.CreationTime,
		Description:           realResponse.Description,
		SecurityGroupRuleList: sgrl,
	}, err
}

type securityGroupDescribeAttributeRealResponse struct {
	SecurityGroupName     string `json:"securityGroupName"`
	Id                    string `json:"id"`
	VmNum                 int    `json:"vmNum"`
	Origin                string `json:"origin"`
	VpcName               string `json:"vpcName"`
	VpcId                 string `json:"vpcID"`
	CreationTime          string `json:"creationTime"`
	Description           string `json:"description"`
	SecurityGroupRuleList []struct {
		Direction       string `json:"direction"`
		Priority        int    `json:"priority"`
		Ethertype       string `json:"ethertype"`
		Protocol        string `json:"protocol"`
		Range           string `json:"range"`
		DestCidrIp      string `json:"destCidrIp"`
		Description     string `json:"description"`
		Origin          string `json:"origin"`
		CreateTime      string `json:"createTime"`
		Id              string `json:"id"`
		Action          string `json:"action"`
		SecurityGroupId string `json:"securityGroupID"`
	} `json:"securityGroupRuleList"`
}

type SecurityGroupDescribeAttributeSecurityGroupRuleListResponse struct {
	Direction       string
	Priority        int
	Ethertype       string
	Protocol        string
	Range           string
	DestCidrIp      string
	Description     string
	Origin          string
	CreateTime      string
	Id              string
	Action          string
	SecurityGroupId string
}

type SecurityGroupDescribeAttributeRequest struct {
	RegionId        string // 区域id
	SecurityGroupId string // 安全组ID
	ProjectId       string // 企业项目 ID，默认为0
	Direction       string // 安全组规则授权方向： egress：安全组出方向 ingress：安全组入方向 all：不区分方向 默认是：all
}

type SecurityGroupDescribeAttributeResponse struct {
	SecurityGroupName     string
	Id                    string
	VmNum                 int
	Origin                string
	VpcName               string
	VpcId                 string
	CreationTime          string
	Description           string
	SecurityGroupRuleList []SecurityGroupDescribeAttributeSecurityGroupRuleListResponse
}
