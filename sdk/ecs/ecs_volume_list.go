package ecs

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// ecsVolumeListApi  查询云主机的云硬盘列表
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=8290&data=87
type ecsVolumeListApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEcsVolumeListApi(client *common.CtyunSender) common.ApiHandler[EcsVolumeListRequest, EcsVolumeListResponse] {
	return &ecsVolumeListApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/volume/list",
		},
	}
}

func (v *ecsVolumeListApi) Do(ctx context.Context, credential *common.Credential, req *EcsVolumeListRequest) (*EcsVolumeListResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&ecsVolumeListRealRequest{
		RegionID:   req.RegionId,
		InstanceID: req.InstanceId,
		PageNo:     req.PageNo,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	// 发起请求
	response, err := v.client.SendCtEcs(ctx, builder)
	if err != nil {
		return nil, err
	}

	// 解析返回
	var realResponse ecsVolumeListRealResponse
	err = response.ParseWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	var results []EcsVolumeListResultsResponse
	for _, result := range realResponse.Results {
		results = append(results, EcsVolumeListResultsResponse{
			DiskType:     result.DiskType,
			IsEncrypt:    result.IsEncrypt,
			DiskSize:     result.DiskSize,
			DiskMode:     result.DiskMode,
			DiskId:       result.DiskID,
			DiskDataType: result.DiskDataType,
		})
	}
	return &EcsVolumeListResponse{
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		Results:      results,
	}, nil
}

type ecsVolumeListRealRequest struct {
	RegionID   string `json:"regionID,omitempty"`
	InstanceID string `json:"instanceID,omitempty"`
	PageNo     int    `json:"securityGroupID,omitempty"`
	PageSize   int    `json:"networkInterfaceID,omitempty"`
}

type EcsVolumeListRequest struct {
	RegionId   string
	InstanceId string
	PageNo     int
	PageSize   int
}

type ecsVolumeListRealResponse struct {
	CurrentCount int `json:"currentCount"`
	TotalCount   int `json:"totalCount"`
	TotalPage    int `json:"totalPage"`
	Results      []struct {
		DiskType     string `json:"diskType"`
		IsEncrypt    bool   `json:"isEncrypt"`
		DiskSize     int    `json:"diskSize"`
		DiskMode     string `json:"diskMode"`
		DiskID       string `json:"diskID"`
		DiskDataType string `json:"diskDataType"`
	} `json:"results"`
}

type EcsVolumeListResultsResponse struct {
	DiskType     string
	IsEncrypt    bool
	DiskSize     int
	DiskMode     string
	DiskId       string
	DiskDataType string
}

type EcsVolumeListResponse struct {
	CurrentCount int
	TotalCount   int
	TotalPage    int
	Results      []EcsVolumeListResultsResponse
}
