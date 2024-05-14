package ebs

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// ebsCreateApi 创建云硬盘
type ebsCreateApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewEbsCreateApi(client *common.CtyunSender) common.ApiHandler[EbsCreateRequest, EbsCreateResponse] {
	return &ebsCreateApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ebs/new-ebs",
		},
	}
}

func (v *ebsCreateApi) Do(ctx context.Context, credential *common.Credential, req *EbsCreateRequest) (*EbsCreateResponse, common.CtyunRequestError) {
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(ebsCreateRealRequest{
		ClientToken: req.ClientToken,
		DiskName:    req.DiskName,
		DiskMode:    req.DiskMode,
		DiskType:    req.DiskType,
		DiskSize:    req.DiskSize,
		RegionID:    req.RegionID,
		AzName:      req.AzName,
		OnDemand:    req.OnDemand,
		CycleType:   req.CycleType,
		CycleCount:  req.CycleCount,
	})
	if err != nil {
		return nil, err
	}

	resp, err := v.client.SendCtEbs(ctx, builder)
	if err != nil {
		return nil, err
	}

	response := &ebsCreateRealResponse{}
	err = resp.ParseWithCheck(response)
	if err != nil {
		return nil, err
	}

	resources := []EbsCreateResource{}
	for _, resource := range response.Resources {
		resources = append(resources, EbsCreateResource{
			OrderID:          resource.OrderID,
			Status:           resource.Status,
			IsMaster:         resource.IsMaster,
			DiskName:         resource.DiskID,
			ResourceType:     resource.ResourceType,
			MasterOrderID:    resource.MasterOrderID,
			UpdateTime:       resource.UpdateTime,
			MasterResourceID: resource.MasterResourceID,
			ItemValue:        resource.ItemValue,
			StartTime:        resource.StartTime,
			CreateTime:       resource.CreateTime,
			DiskID:           resource.DiskID,
		})
	}

	return &EbsCreateResponse{
		MasterResourceStatus: response.MasterResourceStatus,
		RegionID:             response.RegionID,
		MasterOrderID:        response.MasterOrderID,
		MasterResourceID:     response.MasterResourceID,
		MasterOrderNO:        response.MasterOrderNO,
		Resources:            resources,
	}, nil
}

type ebsCreateRealRequest struct {
	ClientToken string `json:"clientToken"`
	DiskName    string `json:"diskName"`
	DiskMode    string `json:"diskMode"`
	DiskType    string `json:"diskType"`
	DiskSize    int64  `json:"diskSize"`
	RegionID    string `json:"regionID"`
	AzName      string `json:"azName"`
	OnDemand    bool   `json:"onDemand"`
	CycleType   string `json:"cycleType"`
	CycleCount  int64  `json:"cycleCount"`
}

type ebsCreateRealResponse struct {
	MasterResourceStatus string                  `json:"masterResourceStatus"`
	RegionID             string                  `json:"regionID"`
	MasterOrderID        string                  `json:"masterOrderID"`
	MasterResourceID     string                  `json:"masterResourceID"`
	MasterOrderNO        string                  `json:"masterOrderNO"`
	Resources            []ebsCreateResourceReal `json:"resources"`
}

type ebsCreateResourceReal struct {
	OrderID          string `json:"orderID"`
	Status           int64  `json:"status"`
	IsMaster         bool   `json:"isMaster"`
	DiskName         string `json:"diskName"`
	ResourceType     string `json:"resourceType"`
	MasterOrderID    string `json:"masterOrderID"`
	UpdateTime       int64  `json:"updateTime"`
	MasterResourceID string `json:"masterResourceID"`
	ItemValue        int64  `json:"itemValue"`
	StartTime        int64  `json:"startTime"`
	CreateTime       int64  `json:"createTime"`
	DiskID           string `json:"diskID"`
}

type EbsCreateRequest struct {
	ClientToken string
	DiskName    string
	DiskMode    string
	DiskType    string
	DiskSize    int64
	RegionID    string
	AzName      string
	OnDemand    bool
	CycleType   string
	CycleCount  int64
}

type EbsCreateResponse struct {
	MasterResourceStatus string
	RegionID             string
	MasterOrderID        string
	MasterResourceID     string
	MasterOrderNO        string
	Resources            []EbsCreateResource
}

type EbsCreateResource struct {
	OrderID          string
	Status           int64
	IsMaster         bool
	DiskName         string
	ResourceType     string
	MasterOrderID    string
	UpdateTime       int64
	MasterResourceID string
	ItemValue        int64
	StartTime        int64
	CreateTime       int64
	DiskID           string
}
