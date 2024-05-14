package vpc

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
	"strconv"
	"strings"
)

// vpcListApi 查询vpc
type vpcListApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewVpcListApi(client *common.CtyunSender) common.ApiHandler[VpcListRequest, VpcListResponse] {
	return &vpcListApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/vpc/new-list",
		},
	}
}

func (v *vpcListApi) Do(ctx context.Context, credential *common.Credential, req *VpcListRequest) (*VpcListResponse, common.CtyunRequestError) {
	// 构建请求
	realRequest := v.
		WithCredential(credential).
		AddParam("regionID", req.RegionId).
		AddParam("vpcID", strings.Join(req.VpcIds, ",")).
		AddParam("pageNumber", strconv.Itoa(req.PageNumber)).
		AddParam("pageSize", strconv.Itoa(req.PageSize))

	// 发起请求
	response, err := v.client.SendCtVpc(ctx, realRequest)
	if err != nil {
		return nil, err
	}

	// 解析返回
	realResponse := &vpcListRealResponse{}
	err = response.ParseWithCheck(realResponse)
	if err != nil {
		return nil, err
	}
	var vpcs []VpcListVpcsResponse
	for _, vpc := range realResponse.Vpcs {
		vpcs = append(vpcs, VpcListVpcsResponse{
			VpcId:          vpc.VpcID,
			Name:           vpc.Name,
			Description:    vpc.Description,
			Cidr:           vpc.CIDR,
			Ipv6Enabled:    vpc.Ipv6Enabled,
			EnableIpv6:     vpc.EnableIpv6,
			Ipv6Cidrs:      vpc.Ipv6CIDRS,
			SecondaryCidrs: vpc.SecondaryCIDRs,
			SubnetIDs:      vpc.SubnetIDs,
			NatGatewayIDs:  vpc.NatGatewayIDs,
		})
	}
	return &VpcListResponse{
		Vpcs:         vpcs,
		CurrentCount: realResponse.CurrentCount,
		TotalPage:    realResponse.TotalPage,
	}, nil
}

type vpcListRealResponse struct {
	Vpcs []struct {
		VpcID          string   `json:"vpcID"`
		Name           string   `json:"name"`
		Description    string   `json:"description"`
		CIDR           string   `json:"CIDR"`
		Ipv6Enabled    bool     `json:"ipv6Enabled"`
		EnableIpv6     bool     `json:"enableIpv6"`
		Ipv6CIDRS      []string `json:"ipv6CIDRS"`
		SecondaryCIDRs []string `json:"secondaryCIDRs"`
		SubnetIDs      []string `json:"subnetIDs"`
		NatGatewayIDs  []string `json:"natGatewayIDs"`
	} `json:"vpcs"`
	CurrentCount int `json:"currentCount"`
	TotalCount   int `json:"totalCount"`
	TotalPage    int `json:"totalPage"`
}

type VpcListRequest struct {
	RegionId   string   // 资源池id
	VpcIds     []string // 查询的vpcId
	VpcName    string   // vpc名称
	PageNumber int      // 列表的页码，默认值为 1。
	PageSize   int      // 分页查询时每页的行数，最大值为 50，默认值为 10。
}

type VpcListVpcsResponse struct {
	VpcId          string
	Name           string
	Description    string
	Cidr           string
	Ipv6Enabled    bool
	EnableIpv6     bool
	Ipv6Cidrs      []string
	SecondaryCidrs []string
	SubnetIDs      []string
	NatGatewayIDs  []string
}

type VpcListResponse struct {
	Vpcs         []VpcListVpcsResponse
	CurrentCount int
	TotalCount   int
	TotalPage    int
}
