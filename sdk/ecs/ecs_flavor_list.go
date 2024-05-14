package ecs

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// ecsFlavorListApi  查询一个或多个云主机规格资源
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=8327&data=87
type ecsFlavorListApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEcsFlavorListApi(client *common.CtyunSender) common.ApiHandler[EcsFlavorListRequest, EcsFlavorListResponse] {
	return &ecsFlavorListApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/flavor/list",
		},
	}
}

func (v *ecsFlavorListApi) Do(ctx context.Context, credential *common.Credential, req *EcsFlavorListRequest) (*EcsFlavorListResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&ecsFlavorListRealRequest{
		RegionID:     req.RegionId,
		AzName:       req.AzName,
		FlavorType:   req.FlavorType,
		FlavorName:   req.FlavorName,
		FlavorCPU:    req.FlavorCpu,
		FlavorRAM:    req.FlavorRam,
		FlavorArch:   req.FlavorArch,
		FlavorSeries: req.FlavorSeries,
		FlavorID:     req.FlavorId,
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
	var realResponse ecsFlavorListRealResponse
	err = response.ParseWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	var flavorList []EcsFlavorListFlavorListResponse
	for _, listRealResponse := range realResponse.FlavorList {
		flavorList = append(flavorList, EcsFlavorListFlavorListResponse{
			GpuVendor:        listRealResponse.GpuVendor,
			CpuInfo:          listRealResponse.CpuInfo,
			BaseBandwidth:    listRealResponse.BaseBandwidth,
			FlavorName:       listRealResponse.FlavorName,
			VideoMemSize:     listRealResponse.VideoMemSize,
			FlavorType:       listRealResponse.FlavorType,
			FlavorSeries:     listRealResponse.FlavorSeries,
			FlavorRam:        listRealResponse.FlavorRAM,
			NicMultiQueue:    listRealResponse.NicMultiQueue,
			Pps:              listRealResponse.Pps,
			FlavorCpu:        listRealResponse.FlavorCPU,
			Bandwidth:        listRealResponse.Bandwidth,
			GpuType:          listRealResponse.GpuType,
			FlavorId:         listRealResponse.FlavorID,
			GpuCount:         listRealResponse.GpuCount,
			Available:        listRealResponse.Available,
			AzList:           listRealResponse.AzList,
			FlavorSeriesName: listRealResponse.FlavorSeriesName,
		})
	}
	return &EcsFlavorListResponse{
		FlavorList: flavorList,
	}, nil
}

type ecsFlavorListRealRequest struct {
	RegionID     string `json:"regionID,omitempty"`
	AzName       string `json:"azName,omitempty"`
	FlavorType   string `json:"flavorType,omitempty"`
	FlavorName   string `json:"flavorName,omitempty"`
	FlavorCPU    int    `json:"flavorCPU,omitempty"`
	FlavorRAM    int    `json:"flavorRAM,omitempty"`
	FlavorArch   string `json:"flavorArch,omitempty"`
	FlavorSeries string `json:"flavorSeries,omitempty"`
	FlavorID     string `json:"flavorID,omitempty"`
}

type ecsFlavorListFlavorListRealResponse struct {
	GpuVendor        string   `json:"gpuVendor"`
	CpuInfo          string   `json:"cpuInfo"`
	BaseBandwidth    float64  `json:"baseBandwidth"`
	FlavorName       string   `json:"flavorName"`
	VideoMemSize     int      `json:"videoMemSize"`
	FlavorType       string   `json:"flavorType"`
	FlavorSeries     string   `json:"flavorSeries"`
	FlavorRAM        int      `json:"flavorRAM"`
	NicMultiQueue    int      `json:"nicMultiQueue"`
	Pps              int      `json:"pps"`
	FlavorCPU        int      `json:"flavorCPU"`
	Bandwidth        float64  `json:"bandwidth"`
	GpuType          string   `json:"gpuType"`
	FlavorID         string   `json:"flavorID"`
	GpuCount         int      `json:"gpuCount"`
	Available        bool     `json:"available"`
	AzList           []string `json:"azList"`
	FlavorSeriesName string   `json:"flavorSeriesName"`
}

type ecsFlavorListRealResponse struct {
	FlavorList []ecsFlavorListFlavorListRealResponse `json:"flavorList"`
}

type EcsFlavorListRequest struct {
	RegionId     string
	AzName       string
	FlavorType   string
	FlavorName   string
	FlavorCpu    int
	FlavorRam    int
	FlavorArch   string
	FlavorSeries string
	FlavorId     string
}

type EcsFlavorListFlavorListResponse struct {
	GpuVendor        string
	CpuInfo          string
	BaseBandwidth    float64
	FlavorName       string
	VideoMemSize     int
	FlavorType       string
	FlavorSeries     string
	FlavorRam        int
	NicMultiQueue    int
	Pps              int
	FlavorCpu        int
	Bandwidth        float64
	GpuType          string
	FlavorId         string
	GpuCount         int
	Available        bool
	AzList           []string
	FlavorSeriesName string
}

type EcsFlavorListResponse struct {
	FlavorList []EcsFlavorListFlavorListResponse
}
