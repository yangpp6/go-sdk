package vpc

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// vpcUpdateApi 创建vpc
type vpcUpdateApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewVpcUpdateApi(client *common.CtyunSender) common.ApiHandler[VpcUpdateRequest, VpcUpdateResponse] {
	return &vpcUpdateApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/vpc/update",
		},
	}
}

func (v *vpcUpdateApi) Do(ctx context.Context, credential *common.Credential, req *VpcUpdateRequest) (*VpcUpdateResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	realRequest := vpcUpdateRealRequest{
		VpcId:       req.VpcId,
		ClientToken: req.ClientToken,
		RegionID:    req.RegionId,
		Name:        req.Name,
		Description: req.Description,
	}
	_, err := builder.WriteJson(realRequest)
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
	return &VpcUpdateResponse{}, nil
}

type vpcUpdateRealRequest struct {
	VpcId       string `json:"vpcID"`
	ClientToken string `json:"clientToken"`
	RegionID    string `json:"regionID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type VpcUpdateRequest struct {
	VpcId       string // 更新的vpcId
	ClientToken string // 客户端存根，用于保证订单幂等性, 长度 1 - 64
	RegionId    string // 资源池id
	Name        string // 虚拟私有云名称
	Description string // 描述
}

type VpcUpdateResponse struct {
}
