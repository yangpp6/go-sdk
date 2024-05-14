package subnet

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// subnetUpdateApi 创建子网
type subnetUpdateApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewSubnetUpdateApi(client *common.CtyunSender) common.ApiHandler[SubnetUpdateRequest, SubnetUpdateResponse] {
	return &subnetUpdateApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/vpc/update-subnet",
		},
	}
}

func (v *subnetUpdateApi) Do(ctx context.Context, credential *common.Credential, req *SubnetUpdateRequest) (*SubnetUpdateResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	requestContent := subnetUpdateRealRequest{
		RegionID:    req.RegionId,
		SubnetID:    req.SubnetId,
		Name:        req.Name,
		Description: req.Description,
		DnsList:     req.DnsList,
	}
	_, err := builder.WriteJson(requestContent)
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
	return &SubnetUpdateResponse{}, nil
}

type subnetUpdateRealRequest struct {
	RegionID    string   `json:"regionID"`
	SubnetID    string   `json:"subnetID"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	DnsList     []string `json:"dnsList"`
}

type SubnetUpdateRequest struct {
	RegionId    string
	SubnetId    string
	Name        string
	Description string
	DnsList     []string
}

type SubnetUpdateResponse struct {
}
