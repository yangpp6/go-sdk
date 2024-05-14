package security_group

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// securityGroupModifyAttributionApi 修改安全组
type securityGroupModifyAttributionApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewSecurityGroupModifyAttributionApi(client *common.CtyunSender) common.ApiHandler[SecurityGroupModifyAttributionRequest, SecurityGroupModifyAttributionResponse] {
	return &securityGroupModifyAttributionApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/vpc/modify-security-group-attribute",
		},
	}
}

func (v *securityGroupModifyAttributionApi) Do(ctx context.Context, credential *common.Credential, req *SecurityGroupModifyAttributionRequest) (*SecurityGroupModifyAttributionResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(securityGroupModifyAttributionRealRequest{
		ClientToken:     req.ClientToken,
		RegionID:        req.RegionId,
		ProjectID:       req.ProjectId,
		Name:            req.Name,
		Description:     req.Description,
		Enabled:         req.Enabled,
		SecurityGroupID: req.SecurityGroupId,
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
	return &SecurityGroupModifyAttributionResponse{}, nil
}

type securityGroupModifyAttributionRealRequest struct {
	ClientToken     string `json:"clientToken"`
	RegionID        string `json:"regionID"`
	ProjectID       string `json:"projectID"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Enabled         bool   `json:"enabled"`
	SecurityGroupID string `json:"securityGroupID"`
}

type SecurityGroupModifyAttributionRequest struct {
	ClientToken     string // 客户端存根，用于保证订单幂等性, 长度 1 - 64
	RegionId        string // 资源池id
	ProjectId       string // 企业项目 ID，默认为0
	Name            string // 子网名称 取值范围：1-64，支持数字、字母、中文、_(下划线)、-（中划线）
	Description     string // 支持拉丁字母、中文、数字, 特殊字符：~!@#$%^&*()_-+= <>?:{},./;'[]·~！@#￥%……&*（） —— -+={},
	Enabled         bool   // 开启安全组 / 关闭安全组
	SecurityGroupId string // 安全组ID
}

type SecurityGroupModifyAttributionResponse struct {
}
