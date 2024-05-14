package bandwidth

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// bandwidthDeleteApi
type bandwidthDeleteApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewBandwidthDeleteApi(client *common.CtyunSender) common.ApiHandler[BandwidthDeleteRequest, BandwidthDeleteResponse] {
	return &bandwidthDeleteApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/bandwidth/delete",
		},
	}
}

func (v *bandwidthDeleteApi) Do(ctx context.Context, credential *common.Credential, req *BandwidthDeleteRequest) (*BandwidthDeleteResponse, common.CtyunRequestError) {
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&bandwidthDeleteRealRequest{
		RegionID:    req.RegionID,
		BandwidthID: req.BandwidthID,
		ClientToken: req.ClientToken,
	})
	if err != nil {
		return nil, err
	}

	resp, err := v.client.SendCtVpc(ctx, builder)
	if err != nil {
		return nil, err
	}

	response := &bandwidthDeleteRealResponse{}
	err = resp.ParseWithCheck(response)
	if err != nil {
		return nil, err
	}
	return &BandwidthDeleteResponse{
		MasterOrderID: response.MasterOrderID,
		MasterOrderNO: response.MasterOrderNO,
		RegionID:      response.RegionID,
	}, nil
}

type bandwidthDeleteRealRequest struct {
	RegionID    string `json:"regionID"`
	ClientToken string `json:"clientToken"`
	BandwidthID string `json:"bandwidthID"`
}

type bandwidthDeleteRealResponse struct {
	MasterOrderID string `json:"masterOrderID"`
	MasterOrderNO string `json:"masterOrderNO"`
	RegionID      string `json:"regionID"`
}

type BandwidthDeleteRequest struct {
	RegionID    string
	ClientToken string
	BandwidthID string
}

type BandwidthDeleteResponse struct {
	MasterOrderID string
	MasterOrderNO string
	RegionID      string
}
