package ebs

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// ebsAssociateApi 绑定弹性IP
type ebsAssociateApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEbsAssociateApi(client *common.CtyunSender) common.ApiHandler[EbsAssociateRequest, EbsAssociateResponse] {
	return &ebsAssociateApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ebs/attach-ebs",
		},
	}
}

func (v *ebsAssociateApi) Do(ctx context.Context, credential *common.Credential, req *EbsAssociateRequest) (*EbsAssociateResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	requestContent := ebsAssociateRealRequest{
		RegionID:   req.RegionId,
		DiskID:     req.DiskId,
		InstanceID: req.InstanceID,
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
	resp := &EbsAssociateRealResponse{}
	err = response.ParseWithCheck(resp)
	if err != nil {
		return nil, err
	}

	return &EbsAssociateResponse{
		DiskJobID: resp.DiskJobID,
	}, nil
}

type ebsAssociateRealRequest struct {
	RegionID   string `json:"regionID"`
	DiskID     string `json:"diskID"`
	InstanceID string `json:"instanceID"`
}

type EbsAssociateRealResponse struct {
	DiskJobID string `json:"diskJobID"`
}

type EbsAssociateRequest struct {
	RegionId   string `json:"regionID"`
	DiskId     string `json:"diskID"`
	InstanceID string `json:"instanceID"`
}

type EbsAssociateResponse struct {
	DiskJobID string `json:"diskJobID"`
}
