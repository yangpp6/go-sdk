package bandwidth

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// bandwidthShowApi
type bandwidthShowApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewBandwidthShowApi(client *common.CtyunSender) common.ApiHandler[BandwidthShowRequest, BandwidthShowResponse] {
	return &bandwidthShowApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/bandwidth/describe",
		},
	}
}

func (v *bandwidthShowApi) Do(ctx context.Context, credential *common.Credential, req *BandwidthShowRequest) (*BandwidthShowResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	builder.AddParam("regionID", req.RegionId)
	builder.AddParam("projectID", req.ProjectID)
	builder.AddParam("bandwidthID", req.BandwidthID)
	// 发起请求
	response, err := v.client.SendCtVpc(ctx, builder)
	if err != nil {
		return nil, err
	}

	// 解析返回
	result := &bandwidthShowRealResponse{}
	err = response.ParseWithCheck(result)
	if err != nil {
		return nil, err
	}
	var eips []Eip
	for _, each := range result.Eips {
		eips = append(eips, Eip{
			Ip:    each.Ip,
			EipID: each.EipID,
		})
	}
	return &BandwidthShowResponse{
		Id:        result.Id,
		Status:    result.Status,
		Bandwidth: result.Bandwidth,
		Name:      result.Name,
		Eips:      eips,
	}, nil
}

type bandwidthShowRealResponse struct {
	Id        string `json:"id"`
	Status    string `json:"status"`
	Bandwidth int    `json:"bandwidth"`
	Name      string `json:"name"`
	Eips      []eip  `json:"eips"`
}

type eip struct {
	Ip    string `json:"ip"`
	EipID string `json:"eipID"`
}

type BandwidthShowRequest struct {
	RegionId    string
	ProjectID   string
	BandwidthID string
}

type BandwidthShowResponse struct {
	Id        string
	Status    string
	Bandwidth int
	Name      string
	Eips      []Eip
}

type Eip struct {
	Ip    string `json:"ip"`
	EipID string `json:"eipID"`
}

var Empty = BandwidthShowResponse{}
