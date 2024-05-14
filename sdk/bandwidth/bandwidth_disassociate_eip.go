package bandwidth

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// bandwidthDisassociationEipApi 解绑弹性IP
type bandwidthDisassociationEipApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewBandwidthDisassociationEipApi(client *common.CtyunSender) common.ApiHandler[BandwidthDisassociationEipRequest, BandwidthDisassociationEipResponse] {
	return &bandwidthDisassociationEipApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/bandwidth/disassociate-eip",
		},
	}
}

func (v *bandwidthDisassociationEipApi) Do(ctx context.Context, credential *common.Credential, req *BandwidthDisassociationEipRequest) (*BandwidthDisassociationEipResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	requestContent := bandwidthDisassociationEipRealRequest{
		RegionID:    req.RegionId,
		ClientToken: req.ClientToken,
		EipIds:      req.EipIds,
		BandwidthID: req.BandwidthID,
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

	return &BandwidthDisassociationEipResponse{}, nil
}

type bandwidthDisassociationEipRealRequest struct {
	RegionID    string   `json:"regionID"`
	ClientToken string   `json:"clientToken"`
	BandwidthID string   `json:"bandwidthID"`
	EipIds      []string `json:"eipIDs"`
}

type BandwidthDisassociationEipRequest struct {
	RegionId    string
	ClientToken string
	EipIds      []string
	BandwidthID string
}

type BandwidthDisassociationEipResponse struct {
}
