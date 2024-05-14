package eip

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
	"time"
)

// eipShowApi 查询弹性IP详情
type eipShowApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEipShowApi(client *common.CtyunSender) common.ApiHandler[EipShowRequest, EipShowResponse] {
	return &eipShowApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/eip/show",
		},
	}
}

func (v *eipShowApi) Do(ctx context.Context, credential *common.Credential, req *EipShowRequest) (*EipShowResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	builder.AddParam("regionID", req.RegionId)
	builder.AddParam("eipID", req.EipId)
	// 发起请求
	response, err := v.client.SendCtVpc(ctx, builder)
	if err != nil {
		return nil, err
	}

	// 解析返回
	result := &eipShowRealResponse{}
	err = response.ParseWithCheck(result)
	if err != nil {
		return nil, err
	}
	return &EipShowResponse{
		Id:               result.ID,
		Name:             result.Name,
		EipAddress:       result.EipAddress,
		AssociationId:    result.AssociationID,
		AssociationType:  result.AssociationType,
		PrivateIpAddress: result.PrivateIpAddress,
		Bandwidth:        result.Bandwidth,
		BandwidthId:      result.BandwidthID,
		BandwidthType:    result.BandwidthType,
		Status:           result.Status,
		Tags:             result.Tags,
		CreatedAt:        result.CreatedAt,
		UpdatedAt:        result.UpdatedAt,
		ExpiredAt:        result.ExpiredAt,
	}, nil
}

type eipShowRealResponse struct {
	ID               string    `json:"ID"`
	Name             string    `json:"name"`
	EipAddress       string    `json:"eipAddress"`
	AssociationID    string    `json:"associationID"`
	AssociationType  string    `json:"associationType"`
	PrivateIpAddress string    `json:"privateIpAddress"`
	Bandwidth        int       `json:"bandwidth"`
	BandwidthID      string    `json:"bandwidthID"`
	BandwidthType    string    `json:"bandwidthType"`
	Status           string    `json:"status"`
	Tags             string    `json:"tags"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	ExpiredAt        time.Time `json:"expiredAt"`
}

type EipShowRequest struct {
	RegionId string
	EipId    string
}

type EipShowResponse struct {
	Id               string
	Name             string
	EipAddress       string
	AssociationId    string
	AssociationType  string
	PrivateIpAddress string
	Bandwidth        int
	BandwidthId      string
	BandwidthType    string
	Status           string
	Tags             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	ExpiredAt        time.Time
}
