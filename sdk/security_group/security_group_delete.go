package security_group

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// securityGroupDeleteApi 删除安全组
type securityGroupDeleteApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewSecurityGroupDeleteApi(client *common.CtyunSender) common.ApiHandler[SecurityGroupDeleteRequest, SecurityGroupDeleteResponse] {
	return &securityGroupDeleteApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/vpc/delete-security-group",
		},
	}
}

func (v *securityGroupDeleteApi) Do(ctx context.Context, credential *common.Credential, req *SecurityGroupDeleteRequest) (*SecurityGroupDeleteResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(securityGroupDeleteRealRequest{
		ClientToken:     req.ClientToken,
		RegionId:        req.RegionId,
		ProjectId:       req.ProjectId,
		SecurityGroupId: req.SecurityGroupId,
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
	return &SecurityGroupDeleteResponse{}, err
}

type securityGroupDeleteRealRequest struct {
	ClientToken     string `json:"clientToken"`
	RegionId        string `json:"regionID"`
	ProjectId       string `json:"projectID"`
	SecurityGroupId string `json:"securityGroupID"`
}

type SecurityGroupDeleteRequest struct {
	ClientToken     string // 客户端存根，用于保证订单幂等性, 长度 1 - 64
	RegionId        string // 资源池id
	ProjectId       string // 企业项目 ID，默认为0
	SecurityGroupId string // 安全组ID
}

type SecurityGroupDeleteResponse struct {
}
