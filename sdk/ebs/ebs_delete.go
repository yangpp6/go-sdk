package ebs

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// ebsCreateApi 删除云硬盘
type ebsDeleteApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEbsDeleteApi(client *common.CtyunSender) common.ApiHandler[EbsDeleteRequest, EbsDeleteResponse] {
	return &ebsDeleteApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ebs/refund-ebs",
		},
	}
}

func (v *ebsDeleteApi) Do(ctx context.Context, credential *common.Credential, req *EbsDeleteRequest) (*EbsDeleteResponse, common.CtyunRequestError) {
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&ebsDeleteRealRequest{
		RegionID:    req.RegionID,
		DiskID:      req.DiskID,
		ClientToken: req.ClientToken,
	})
	if err != nil {
		return nil, err
	}

	resp, err := v.client.SendCtEbs(ctx, builder)
	if err != nil {
		return nil, err
	}

	response := &ebsDeleteRealResponse{}
	err = resp.ParseWithCheck(response)
	if err != nil {
		return nil, err
	}
	return &EbsDeleteResponse{
		MasterOrderID: response.MasterOrderID,
		MasterOrderNO: response.MasterOrderNO,
	}, nil
}

type ebsDeleteRealRequest struct {
	RegionID    string `json:"regionID"`
	DiskID      string `json:"diskID"`
	ClientToken string `json:"clientToken"`
}

type ebsDeleteRealResponse struct {
	MasterOrderID string `json:"masterOrderID"`
	MasterOrderNO string `json:"masterOrderNO"`
}

type EbsDeleteRequest struct {
	RegionID    string
	DiskID      string
	ClientToken string
}

type EbsDeleteResponse struct {
	MasterOrderID string
	MasterOrderNO string
}
