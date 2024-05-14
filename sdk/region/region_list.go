package region

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// RegionListApi 查询弹性IP详情
type RegionListApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewRegionListApi(client *common.CtyunSender) common.ApiHandler[RegionListRequest, RegionListResponse] {
	return &RegionListApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/region/list-regions",
		},
	}
}

func (v *RegionListApi) Do(ctx context.Context, credential *common.Credential, req *RegionListRequest) (*RegionListResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	builder.AddParam("regionName", req.RegionName)
	// 发起请求
	response, err := v.client.SendCtEcs(ctx, builder)
	if err != nil {
		return nil, err
	}

	// 解析返回
	result := &RegionListRealResponse{}
	err = response.ParseWithCheck(result)
	if err != nil {
		return nil, err
	}
	var regionList []RegionListRegionListResponse
	for _, region := range result.RegionList {
		regionList = append(regionList, RegionListRegionListResponse{
			RegionParent: region.RegionParent,
			RegionId:     region.RegionID,
			RegionType:   region.RegionType,
			RegionName:   region.RegionName,
			IsMultiZones: region.IsMultiZones,
			ZoneList:     region.ZoneList,
		})
	}
	return &RegionListResponse{
		RegionList: regionList,
	}, nil
}

type RegionListRealResponse struct {
	RegionList []struct {
		RegionParent string   `json:"regionParent"`
		RegionID     string   `json:"regionID"`
		RegionType   string   `json:"regionType"`
		RegionName   string   `json:"regionName"`
		IsMultiZones bool     `json:"isMultiZones"`
		ZoneList     []string `json:"zoneList"`
	} `json:"regionList"`
}

type RegionListRequest struct {
	RegionName string
}

type RegionListRegionListResponse struct {
	RegionParent string
	RegionId     string
	RegionType   string
	RegionName   string
	IsMultiZones bool
	ZoneList     []string
}

type RegionListResponse struct {
	RegionList []RegionListRegionListResponse
}
