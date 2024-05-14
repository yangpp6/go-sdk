package vpc

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// vpcDeleteApi 删除vpc
type vpcDeleteApi struct {
	builder common.CtyunRequestBuilder
	client  *common.CtyunSender
}

func NewVpcDeleteApi(client *common.CtyunSender) common.ApiHandler[VpcDeleteRequest, VpcDeleteResponse] {
	return &vpcDeleteApi{
		client: client,
		builder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/vpc/delete",
		},
	}
}

func (v *vpcDeleteApi) Do(ctx context.Context, credential *common.Credential, req *VpcDeleteRequest) (*VpcDeleteResponse, common.CtyunRequestError) {
	// 构建请求
	realRequest, err := v.builder.WithCredential(credential).WriteJson(vpcDeleteRealRequest{
		RegionID: req.RegionId,
		VpcID:    req.VpcId,
	})

	// 发起请求
	response, err := v.client.SendCtVpc(ctx, realRequest)
	if err != nil {
		return nil, err
	}

	// 解析返回
	result := &VpcDeleteResponse{}
	err = response.ParseWithCheck(result)
	return result, err
}

type vpcDeleteRealRequest struct {
	RegionID string `json:"regionID"`
	VpcID    string `json:"vpcID"`
}

type VpcDeleteRequest struct {
	RegionId string // 资源池ID
	VpcId    string // VPC的ID
}

type VpcDeleteResponse struct {
}
