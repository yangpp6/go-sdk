package bandwidth

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// bandwidthChangeSpecApi
type bandwidthChangeSpecApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewBandwidthChangeSpecApi(client *common.CtyunSender) common.ApiHandler[BandwidthChangeSpecRequest, BandwidthChangeSpecResponse] {
	return &bandwidthChangeSpecApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/bandwidth/modify-spec",
		},
	}
}

func (v *bandwidthChangeSpecApi) Do(ctx context.Context, credential *common.Credential, req *BandwidthChangeSpecRequest) (*BandwidthChangeSpecResponse, common.CtyunRequestError) {
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&bandwidthChangeSpecRealRequest{
		RegionID:    req.RegionID,
		ClientToken: req.ClientToken,
		BandwidthID: req.BandwidthID,
		Bandwidth:   req.Bandwidth,
	})
	if err != nil {
		return nil, err
	}

	resp, err := v.client.SendCtVpc(ctx, builder)
	if err != nil {
		return nil, err
	}

	response := &bandwidthChangeSpecRealResponse{}
	err = resp.ParseWithCheck(response)
	if err != nil {
		return nil, err
	}
	return &BandwidthChangeSpecResponse{
		MasterOrderID: response.MasterOrderID,
		MasterOrderNO: response.MasterOrderNO,
		RegionID:      response.RegionID,
	}, nil
}

type bandwidthChangeSpecRealRequest struct {
	RegionID    string `json:"regionID"`
	ClientToken string `json:"clientToken"`
	BandwidthID string `json:"bandwidthID"`
	Bandwidth   int    `json:"bandwidth"`
}

type bandwidthChangeSpecRealResponse struct {
	MasterOrderID string `json:"masterOrderID"`
	MasterOrderNO string `json:"masterOrderNO"`
	RegionID      string `json:"regionID"`
}

type BandwidthChangeSpecRequest struct {
	RegionID    string
	ClientToken string
	BandwidthID string
	Bandwidth   int
}

type BandwidthChangeSpecResponse struct {
	MasterOrderID string
	MasterOrderNO string
	RegionID      string
}
