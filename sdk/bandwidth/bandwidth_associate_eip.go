package bandwidth

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// bandwidthAssociationEipApi 绑定弹性IP
type bandwidthAssociationEipApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewBandwidthAssociationEipApi(client *common.CtyunSender) common.ApiHandler[BandwidthAssociationEipRequest, BandwidthAssociationEipResponse] {
	return &bandwidthAssociationEipApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/bandwidth/associate-eip",
		},
	}
}

func (v *bandwidthAssociationEipApi) Do(ctx context.Context, credential *common.Credential, req *BandwidthAssociationEipRequest) (*BandwidthAssociationEipResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	requestContent := bandwidthAssociationEipRealRequest{
		RegionID:    req.RegionID,
		ClientToken: req.ClientToken,
		BandwidthID: req.BandwidthID,
		EipIDs:      req.EipIDs,
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

	return &BandwidthAssociationEipResponse{}, nil
}

type bandwidthAssociationEipRealRequest struct {
	RegionID    string   `json:"regionID"`
	ClientToken string   `json:"clientToken"`
	BandwidthID string   `json:"bandwidthID"`
	EipIDs      []string `json:"eipIDs"`
}

type BandwidthAssociationEipRequest struct {
	RegionID    string
	ClientToken string
	BandwidthID string
	EipIDs      []string
}

type BandwidthAssociationEipResponse struct {
}
