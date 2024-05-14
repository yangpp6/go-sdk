package security_group

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// securityGroupCreateApi 创建安全组
type securityGroupCreateApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewSecurityGroupCreateApi(client *common.CtyunSender) common.ApiHandler[SecurityGroupCreateRequest, SecurityGroupCreateResponse] {
	return &securityGroupCreateApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/vpc/create-security-group",
		},
	}
}

func (v *securityGroupCreateApi) Do(ctx context.Context, credential *common.Credential, req *SecurityGroupCreateRequest) (*SecurityGroupCreateResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(securityGroupCreateRealRequest{
		ClientToken: req.ClientToken,
		RegionID:    req.RegionId,
		ProjectID:   req.ProjectId,
		VpcID:       req.VpcId,
		Name:        req.Name,
		Description: req.Description,
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
	result := &securityGroupCreateRealResponse{}
	err = response.ParseWithCheck(result)
	if err != nil {
		return nil, err
	}
	return &SecurityGroupCreateResponse{
		SecurityGroupId: result.SecurityGroupId,
	}, nil
}

type securityGroupCreateRealRequest struct {
	ClientToken string `json:"clientToken"`
	RegionID    string `json:"regionID"`
	ProjectID   string `json:"projectID"`
	VpcID       string `json:"vpcID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type securityGroupCreateRealResponse struct {
	SecurityGroupId string `json:"securityGroupID"`
}

type SecurityGroupCreateRequest struct {
	RegionId    string // 资源池id
	VpcId       string // vpcId
	Name        string // 子网名称 取值范围：1-64，支持数字、字母、中文、_(下划线)、-（中划线）
	Description string // 支持拉丁字母、中文、数字, 特殊字符：~!@#$%^&*()_-+= <>?:{},./;'[]·~！@#￥%……&*（） —— -+={},
	ClientToken string // 客户端存根，用于保证订单幂等性, 长度 1 - 64
	ProjectId   string // 企业项目 ID，默认为0
}

type SecurityGroupCreateResponse struct {
	SecurityGroupId string
}
