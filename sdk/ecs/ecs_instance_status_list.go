package ecs

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// ecsInstanceStatusListApi  获取多台云主机状态
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=8308&data=87
type ecsInstanceStatusListApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEcsInstanceStatusListApi(client *common.CtyunSender) common.ApiHandler[EcsInstanceStatusListRequest, EcsInstanceStatusListResponse] {
	return &ecsInstanceStatusListApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/instance-status-list",
		},
	}
}

func (v *ecsInstanceStatusListApi) Do(ctx context.Context, credential *common.Credential, req *EcsInstanceStatusListRequest) (*EcsInstanceStatusListResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&ecsInstanceStatusListRealRequest{
		RegionId:       req.RegionId,
		AzName:         req.AzName,
		InstanceIDList: req.InstanceIdList,
		PageNo:         req.PageNo,
		PageSize:       req.PageSize,
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
	var realResponse ecsInstanceStatusListRealResponse
	err = response.ParseWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	var ecsInstanceStatusListStatusListResponse []EcsInstanceStatusListStatusListResponse
	for _, listRealResponse := range realResponse.StatusList {
		ecsInstanceStatusListStatusListResponse = append(ecsInstanceStatusListStatusListResponse, EcsInstanceStatusListStatusListResponse{
			InstanceId:     listRealResponse.InstanceID,
			InstanceStatus: listRealResponse.InstanceStatus,
		})
	}
	return &EcsInstanceStatusListResponse{
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		StatusList:   ecsInstanceStatusListStatusListResponse,
	}, nil
}

type ecsInstanceStatusListRealRequest struct {
	RegionId       string `json:"regionID"`
	AzName         string `json:"azName"`
	InstanceIDList string `json:"instanceIDList"`
	PageNo         int    `json:"pageNo"`
	PageSize       int    `json:"pageSize"`
}

type ecsInstanceStatusListRealResponse struct {
	CurrentCount int `json:"currentCount"`
	TotalCount   int `json:"totalCount"`
	TotalPage    int `json:"totalPage"`
	StatusList   []struct {
		InstanceID     string `json:"instanceID"`
		InstanceStatus string `json:"instanceStatus"`
	} `json:"statusList"`
}

type EcsInstanceStatusListRequest struct {
	RegionId       string
	AzName         string
	InstanceIdList string
	PageNo         int
	PageSize       int
}

type EcsInstanceStatusListResponse struct {
	CurrentCount int
	TotalCount   int
	TotalPage    int
	StatusList   []EcsInstanceStatusListStatusListResponse
}

type EcsInstanceStatusListStatusListResponse struct {
	InstanceId     string
	InstanceStatus string
}
