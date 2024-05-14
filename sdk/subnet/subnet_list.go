package subnet

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
	"strconv"
	"strings"
)

// subnetListApi 查询子网
type subnetListApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewSubnetListApi(client *common.CtyunSender) common.ApiHandler[SubnetListRequest, SubnetListResponse] {
	return &subnetListApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/vpc/new-list-subnet",
		},
	}
}

func (v *subnetListApi) Do(ctx context.Context, credential *common.Credential, req *SubnetListRequest) (*SubnetListResponse, common.CtyunRequestError) {
	// 构建请求
	realRequest := v.
		WithCredential(credential).
		AddParam("regionID", req.RegionId).
		AddParam("vpcID", req.VpcId).
		AddParam("subnetID", strings.Join(req.SubnetIds, ",")).
		AddParam("pageNumber", strconv.Itoa(req.PageNumber)).
		AddParam("pageSize", strconv.Itoa(req.PageSize))
	// 发起请求
	response, err := v.client.SendCtVpc(ctx, realRequest)
	if err != nil {
		return nil, err
	}
	// 解析返回
	realResponse := &subnetListRealResponse{}
	err = response.ParseWithCheck(realResponse)
	if err != nil {
		return nil, err
	}
	var subnets []SubnetListSubnetsResponse
	for _, s := range realResponse.Subnets {
		subnets = append(subnets, SubnetListSubnetsResponse{
			SubnetId:          s.SubnetID,
			Name:              s.Name,
			Description:       s.Description,
			VpcId:             s.VpcID,
			Cidr:              s.CIDR,
			AvailableIpCount:  s.AvailableIPCount,
			GatewayIp:         s.GatewayIP,
			AvailabilityZones: s.AvailabilityZones,
			RouteTableId:      s.RouteTableID,
			NetworkAclId:      s.NetworkAclID,
			Start:             s.Start,
			End:               s.End,
			Ipv6Enabled:       s.Ipv6Enabled,
			Ipv6Cidr:          s.Ipv6CIDR,
			Ipv6Start:         s.Ipv6Start,
			Ipv6End:           s.Ipv6End,
			Ipv6GatewayIp:     s.Ipv6GatewayIP,
			DnsList:           s.DnsList,
			NtpList:           s.NtpList,
			Type:              s.Type,
			CreateAt:          s.CreateAt,
			UpdateAt:          s.UpdateAt,
		})
	}
	return &SubnetListResponse{
		Subnets:      subnets,
		CurrentCount: realResponse.CurrentCount,
		TotalPage:    realResponse.TotalPage,
	}, nil
}

type subnetListRealResponse struct {
	Subnets []struct {
		SubnetID          string   `json:"subnetID"`
		Name              string   `json:"name"`
		Description       string   `json:"description"`
		VpcID             string   `json:"vpcID"`
		CIDR              string   `json:"CIDR"`
		AvailableIPCount  int      `json:"availableIPCount"`
		GatewayIP         string   `json:"gatewayIP"`
		AvailabilityZones []string `json:"availabilityZones"`
		RouteTableID      string   `json:"routeTableID"`
		NetworkAclID      string   `json:"networkAclID"`
		Start             string   `json:"start"`
		End               string   `json:"end"`
		Ipv6Enabled       int      `json:"ipv6Enabled"`
		EnableIpv6        bool     `json:"enableIpv6"`
		Ipv6CIDR          string   `json:"ipv6CIDR"`
		Ipv6Start         string   `json:"ipv6Start"`
		Ipv6End           string   `json:"ipv6End"`
		Ipv6GatewayIP     string   `json:"ipv6GatewayIP"`
		DnsList           []string `json:"dnsList"`
		NtpList           []string `json:"ntpList"`
		Type              int      `json:"type"`
		CreateAt          string   `json:"createAt"`
		UpdateAt          string   `json:"updateAt"`
	} `json:"subnets"`
	CurrentCount int `json:"currentCount"`
	TotalCount   int `json:"totalCount"`
	TotalPage    int `json:"totalPage"`
}

type SubnetListRequest struct {
	RegionId   string   // 资源池id
	VpcId      string   // 查询的vpcId
	SubnetIds  []string // 查询的vpcSubnetId
	PageNumber int      // 列表的页码，默认值为 1。
	PageSize   int      // 分页查询时每页的行数，最大值为 50，默认值为 10。
}

type SubnetListSubnetsResponse struct {
	SubnetId          string
	Name              string
	Description       string
	VpcId             string
	Cidr              string
	AvailableIpCount  int
	GatewayIp         string
	AvailabilityZones []string
	RouteTableId      string
	NetworkAclId      string
	Start             string
	End               string
	Ipv6Enabled       int
	EnableIpv6        bool
	Ipv6Cidr          string
	Ipv6Start         string
	Ipv6End           string
	Ipv6GatewayIp     string
	DnsList           []string
	NtpList           []string
	Type              int
	CreateAt          string
	UpdateAt          string
}

type SubnetListResponse struct {
	Subnets      []SubnetListSubnetsResponse
	CurrentCount int
	TotalCount   int
	TotalPage    int
}
