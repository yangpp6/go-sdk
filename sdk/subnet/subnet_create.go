package subnet

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// subnetCreateApi 创建子网
type subnetCreateApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewSubnetCreateApi(client *common.CtyunSender) common.ApiHandler[SubnetCreateRequest, SubnetCreateResponse] {
	return &subnetCreateApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/vpc/create-subnet",
		},
	}
}

func (v *subnetCreateApi) Do(ctx context.Context, credential *common.Credential, req *SubnetCreateRequest) (*SubnetCreateResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	requestContent := subnetCreateRealRequest{
		RegionID:        req.RegionId,
		ClientToken:     req.ClientToken,
		Name:            req.Name,
		VpcID:           req.VpcId,
		CIDR:            req.Cidr,
		Description:     req.Description,
		EnableIpv6:      req.EnableIpv6,
		DnsList:         req.DnsList,
		SubnetGatewayIP: req.SubnetGatewayIp,
		SubnetType:      req.SubnetType,
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
	result := &subnetCreateRealResponse{}
	err = response.ParseWithCheck(result)
	if err != nil {
		return nil, err
	}

	return &SubnetCreateResponse{
		SubnetId: result.SubnetID,
	}, nil
}

type subnetCreateRealRequest struct {
	RegionID        string   `json:"regionID"`
	ClientToken     string   `json:"clientToken"`
	Name            string   `json:"name"`
	VpcID           string   `json:"vpcID"`
	CIDR            string   `json:"CIDR"`
	Description     string   `json:"description"`
	EnableIpv6      bool     `json:"enableIpv6"`
	DnsList         []string `json:"dnsList"`
	SubnetGatewayIP string   `json:"subnetGatewayIP,omitempty"`
	SubnetType      string   `json:"subnetType"`
}

type subnetCreateRealResponse struct {
	SubnetID string `json:"subnetID"`
}

type SubnetCreateRequest struct {
	RegionId        string
	ClientToken     string
	Name            string
	VpcId           string
	Cidr            string
	Description     string
	EnableIpv6      bool
	DnsList         []string
	SubnetGatewayIp string
	SubnetType      string
}

type SubnetCreateResponse struct {
	SubnetId string
}
