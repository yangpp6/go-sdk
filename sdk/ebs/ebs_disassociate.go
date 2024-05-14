package ebs

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// ebsDisassociateApi 解绑弹性IP
type ebsDisassociateApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEbsDisassociateApi(client *common.CtyunSender) common.ApiHandler[EbsDisassociateRequest, EbsDisassociateResponse] {
	return &ebsDisassociateApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ebs/detach-ebs",
		},
	}
}

func (v *ebsDisassociateApi) Do(ctx context.Context, credential *common.Credential, req *EbsDisassociateRequest) (*EbsDisassociateResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	requestContent := ebsDisassociateRealRequest{
		DiskID:   req.DiskId,
		RegionID: req.RegionId,
	}
	_, err := builder.WriteJson(requestContent)
	if err != nil {
		return nil, err
	}

	// 发起请求
	response, err := v.client.SendCtEbs(ctx, builder)
	if err != nil {
		return nil, err
	}

	// 解析返回
	resp := &ebsDisassociateRealResponse{}
	err = response.ParseWithCheck(resp)
	if err != nil {
		return nil, err
	}

	return &EbsDisassociateResponse{
		DiskJobID: resp.DiskJobID,
	}, nil
}

type ebsDisassociateRealRequest struct {
	DiskID   string `json:"diskID"`
	RegionID string `json:"regionID"`
}

type ebsDisassociateRealResponse struct {
	DiskJobID string `json:"diskJobID"`
}

type EbsDisassociateRequest struct {
	DiskId   string
	RegionId string
}

type EbsDisassociateResponse struct {
	DiskJobID string
}
