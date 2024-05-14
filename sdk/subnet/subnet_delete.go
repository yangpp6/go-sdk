package subnet

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// subnetDeleteApi 删除子网
type subnetDeleteApi struct {
	builder common.CtyunRequestBuilder
	client  *common.CtyunSender
}

func NewSubnetDeleteApi(client *common.CtyunSender) common.ApiHandler[SubnetDeleteRequest, SubnetDeleteResponse] {
	return &subnetDeleteApi{
		client: client,
		builder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/vpc/delete-subnet",
		},
	}
}

func (v *subnetDeleteApi) Do(ctx context.Context, credential *common.Credential, req *SubnetDeleteRequest) (*SubnetDeleteResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.builder.WithCredential(credential)
	builder, err := builder.WriteJson(subnetDeleteRealRequest{
		RegionId: req.RegionId,
		SubnetId: req.SubnetId,
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
	return &SubnetDeleteResponse{}, nil
}

type subnetDeleteRealRequest struct {
	RegionId string `json:"regionID"`
	SubnetId string `json:"subnetID"`
}

type SubnetDeleteRequest struct {
	RegionId string
	SubnetId string
}

type SubnetDeleteResponse struct {
}
