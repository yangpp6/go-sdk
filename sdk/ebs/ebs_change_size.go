package ebs

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// ebsCreateApi 云硬盘修改名称
type ebsChangeSizeApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEbsChangeSizeApi(client *common.CtyunSender) common.ApiHandler[EbsChangeSizeRequest, EbsChangeSizeResponse] {
	return &ebsChangeSizeApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ebs/resize-ebs",
		},
	}
}

func (v *ebsChangeSizeApi) Do(ctx context.Context, credential *common.Credential, req *EbsChangeSizeRequest) (*EbsChangeSizeResponse, common.CtyunRequestError) {
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&ebsChangeSizeRealRequest{
		RegionID:    req.RegionId,
		DiskID:      req.DiskId,
		DiskSize:    req.DiskSize,
		ClientToken: req.ClientToken,
	})
	if err != nil {
		return nil, err
	}

	resp, err := v.client.SendCtEbs(ctx, builder)
	if err != nil {
		return nil, err
	}

	response := &ebsChangeSizeRealResponse{}
	err = resp.ParseWithCheck(response)
	if err != nil {
		return nil, err
	}
	return &EbsChangeSizeResponse{
		MasterOrderId: response.MasterOrderID,
		MasterOrderNo: response.MasterOrderNO,
	}, nil
}

type ebsChangeSizeRealRequest struct {
	RegionID    string `json:"regionID"`
	DiskID      string `json:"diskID"`
	DiskSize    int64  `json:"diskSize"`
	ClientToken string `json:"clientToken"`
}

type ebsChangeSizeRealResponse struct {
	MasterOrderID string `json:"masterOrderID"`
	MasterOrderNO string `json:"masterOrderNO"`
}

type EbsChangeSizeRequest struct {
	RegionId    string
	DiskId      string
	DiskSize    int64
	ClientToken string
}

type EbsChangeSizeResponse struct {
	MasterOrderId string
	MasterOrderNo string
}
