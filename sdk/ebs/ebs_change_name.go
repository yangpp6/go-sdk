package ebs

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// ebsCreateApi 云硬盘修改名称
type ebsChangeNameApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEbsChangeNameApi(client *common.CtyunSender) common.ApiHandler[EbsChangeNameRequest, EbsChangeNameResponse] {
	return &ebsChangeNameApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ebs/update-attr-ebs",
		},
	}
}

func (v *ebsChangeNameApi) Do(ctx context.Context, credential *common.Credential, req *EbsChangeNameRequest) (*EbsChangeNameResponse, common.CtyunRequestError) {
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&ebsChangeNameRealRequest{
		RegionID: req.RegionId,
		DiskID:   req.DiskId,
		DiskName: req.DiskName,
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
	return nil, nil
}

type ebsChangeNameRealRequest struct {
	RegionID string `json:"regionID"`
	DiskID   string `json:"diskID"`
	DiskName string `json:"diskName"`
}

type ebsChangeNameRealResponse struct {
}

type EbsChangeNameRequest struct {
	RegionId string
	DiskId   string
	DiskName string
}

type EbsChangeNameResponse struct {
}
