package bandwidth

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// bandwidthCreateApi
type bandwidthCreateApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewBandwidthCreateApi(client *common.CtyunSender) common.ApiHandler[BandwidthCreateRequest, BandwidthCreateResponse] {
	return &bandwidthCreateApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/bandwidth/create",
		},
	}
}

func (v *bandwidthCreateApi) Do(ctx context.Context, credential *common.Credential, req *BandwidthCreateRequest) (*BandwidthCreateResponse, common.CtyunRequestError) {
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(bandwidthCreateRealRequest{
		RegionID:    req.RegionID,
		ClientToken: req.ClientToken,
		CycleType:   req.CycleType,
		Bandwidth:   req.Bandwidth,
		CycleCount:  req.CycleCount,
		Name:        req.Name,
	})
	if err != nil {
		return nil, err
	}

	resp, err := v.client.SendCtVpc(ctx, builder)
	if err != nil {
		return nil, err
	}

	response := &bandwidthCreateRealResponse{}
	err = resp.ParseWithCheck(response)
	if err != nil {
		return nil, err
	}

	return &BandwidthCreateResponse{
		MasterOrderID:        response.MasterOrderID,
		MasterOrderNO:        response.MasterOrderNO,
		MasterResourceID:     response.MasterResourceID,
		MasterResourceStatus: response.MasterResourceStatus,
		RegionID:             response.RegionID,
		BandwidthId:          response.BandwidthId,
	}, nil
}

type bandwidthCreateRealRequest struct {
	RegionID    string `json:"regionID"`
	ClientToken string `json:"clientToken"`
	CycleType   string `json:"cycleType"`
	Bandwidth   int64  `json:"bandwidth"`
	CycleCount  int64  `json:"cycleCount"`
	Name        string `json:"name"`
}

type bandwidthCreateRealResponse struct {
	MasterOrderID        string `json:"masterOrderID"`
	MasterOrderNO        string `json:"masterOrderNO"`
	MasterResourceID     string `json:"masterResourceID"`
	MasterResourceStatus string `json:"masterResourceStatus"`
	RegionID             string `json:"regionID"`
	BandwidthId          string `json:"bandwidthId"`
}

type BandwidthCreateRequest struct {
	RegionID    string
	ClientToken string
	CycleType   string
	Bandwidth   int64
	CycleCount  int64
	Name        string
}

type BandwidthCreateResponse struct {
	MasterOrderID        string
	MasterOrderNO        string
	MasterResourceID     string
	MasterResourceStatus string
	RegionID             string
	BandwidthId          string
}
