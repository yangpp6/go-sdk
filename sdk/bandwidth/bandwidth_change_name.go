package bandwidth

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// bandwidthChangeName
type bandwidthChangeNameApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewBandwidthChangeNameApi(client *common.CtyunSender) common.ApiHandler[BandwidthChangeNameRequest, BandwidthChangeNameResponse] {
	return &bandwidthChangeNameApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/bandwidth/modify-attribute",
		},
	}
}

func (v *bandwidthChangeNameApi) Do(ctx context.Context, credential *common.Credential, req *BandwidthChangeNameRequest) (*BandwidthChangeNameResponse, common.CtyunRequestError) {
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&bandwidthChangeNameRealRequest{
		RegionID:    req.RegionID,
		ClientToken: req.ClientToken,
		BandwidthID: req.BandwidthID,
		Name:        req.Name,
	})
	if err != nil {
		return nil, err
	}

	resp, err := v.client.SendCtVpc(ctx, builder)
	if err != nil {
		return nil, err
	}

	response := &bandwidthChangeNameRealResponse{}
	err = resp.ParseWithCheck(response)
	if err != nil {
		return nil, err
	}
	return &BandwidthChangeNameResponse{}, nil
}

type bandwidthChangeNameRealRequest struct {
	RegionID    string `json:"regionID"`
	ClientToken string `json:"clientToken"`
	BandwidthID string `json:"bandwidthID"`
	Name        string `json:"name"`
}

type bandwidthChangeNameRealResponse struct {
}

type BandwidthChangeNameRequest struct {
	RegionID    string
	ClientToken string
	BandwidthID string
	Name        string
}

type BandwidthChangeNameResponse struct {
}
