package order

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// orderQueryResourceApi 查询资源详细信息
type orderQueryResourceApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewOrderQueryResourceApi(client *common.CtyunSender) common.ApiHandler[OrderQueryResourceRequest, OrderQueryResourceResponse] {
	return &orderQueryResourceApi{
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v3/order/queryResourceInfoByMasterOrderId",
		},
		client: client,
	}
}

func (o *orderQueryResourceApi) Do(ctx context.Context, credential *common.Credential, req *OrderQueryResourceRequest) (*OrderQueryResourceResponse, common.CtyunRequestError) {
	// 构建请求
	builder := o.WithCredential(credential)
	builder.AddHeader("masterOrderId", req.MasterOrderId)

	// 发起请求
	response, err := o.client.SendApiProxy(ctx, builder)
	if err != nil {
		return nil, err
	}

	// 解析返回
	var realResponse []orderQueryResourceRealResponse
	err = response.ParseWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	var resources []OrderQueryResourceResourceResponse
	for _, res := range realResponse {
		resources = append(resources, OrderQueryResourceResourceResponse{
			OrderItemId:        res.OrderItemId,
			InstanceId:         res.InstanceId,
			AccountId:          res.AccountId,
			UserId:             res.UserId,
			InnerOrderId:       res.InnerOrderId,
			InnerOrderItemId:   res.InnerOrderItemId,
			ProductId:          res.ProductId,
			MasterOrderId:      res.MasterOrderId,
			OrderId:            res.OrderId,
			MasterResourceId:   res.MasterResourceId,
			ResourceId:         res.ResourceId,
			ServiceTag:         res.ServiceTag,
			ResourceType:       res.ResourceType,
			ResourceInfo:       res.ResourceInfo,
			StartDate:          res.StartDate,
			ExpireDate:         res.ExpireDate,
			CreateDate:         res.CreateDate,
			UpdateDate:         res.UpdateDate,
			Status:             res.Status,
			WorkOrderId:        res.WorkOrderId,
			WorkOrderItemId:    res.WorkOrderItemId,
			SalesEntryId:       res.SalesEntryId,
			OrderStatus:        res.OrderStatus,
			ToOndemand:         res.ToOndemand,
			ItemValue:          res.ItemValue,
			DoubleItemValue:    res.DoubleItemValue,
			ChargingStatus:     res.ChargingStatus,
			ChargingDate:       res.ChargingDate,
			ResourceConfig:     res.ResourceConfig,
			AutoToOnDemand:     res.AutoToOnDemand,
			BuildingChannel:    res.BuildingChannel,
			IsPlatformSpecific: res.IsPlatformSpecific,
			BillingOwner:       res.BillingOwner,
			IsPackage:          res.IsPackage,
			CanRelease:         res.CanRelease,
			IsChargeOff:        res.IsChargeOff,
			IsPublicTest:       res.IsPublicTest,
			AutoRenewStatus:    res.AutoRenewStatus,
			Master:             res.Master,
			ResourceConfigMap:  res.ResourceConfigMap,
		})
	}
	return &OrderQueryResourceResponse{
		Resource: resources,
	}, nil
}

type OrderQueryResourceRequest struct {
	MasterOrderId string
}

type OrderQueryResourceResourceResponse struct {
	OrderItemId        string
	InstanceId         string
	AccountId          string
	UserId             string
	InnerOrderId       string
	InnerOrderItemId   string
	ProductId          string
	MasterOrderId      string
	OrderId            string
	MasterResourceId   string
	ResourceId         string
	ServiceTag         string
	ResourceType       string
	ResourceInfo       string
	StartDate          int64
	ExpireDate         int64
	CreateDate         int64
	UpdateDate         int64
	Status             int
	WorkOrderId        string
	WorkOrderItemId    string
	SalesEntryId       string
	OrderStatus        int
	ToOndemand         string
	ItemValue          string
	DoubleItemValue    float64
	ChargingStatus     int
	ChargingDate       int64
	ResourceConfig     string
	AutoToOnDemand     bool
	BuildingChannel    int
	IsPlatformSpecific bool
	BillingOwner       int
	IsPackage          bool
	CanRelease         bool
	IsChargeOff        bool
	IsPublicTest       int
	AutoRenewStatus    int
	Master             bool
	ResourceConfigMap  map[string]any
}

type OrderQueryResourceResponse struct {
	Resource []OrderQueryResourceResourceResponse
}

type orderQueryResourceRealResponse struct {
	OrderItemId        string         `json:"orderItemId"`
	InstanceId         string         `json:"instanceId"`
	AccountId          string         `json:"accountId"`
	UserId             string         `json:"userId"`
	InnerOrderId       string         `json:"innerOrderId"`
	InnerOrderItemId   string         `json:"innerOrderItemId"`
	ProductId          string         `json:"productId"`
	MasterOrderId      string         `json:"masterOrderId"`
	OrderId            string         `json:"orderId"`
	MasterResourceId   string         `json:"masterResourceId"`
	ResourceId         string         `json:"resourceId"`
	ServiceTag         string         `json:"serviceTag"`
	ResourceType       string         `json:"resourceType"`
	ResourceInfo       string         `json:"resourceInfo"`
	StartDate          int64          `json:"startDate"`
	ExpireDate         int64          `json:"expireDate"`
	CreateDate         int64          `json:"createDate"`
	UpdateDate         int64          `json:"updateDate"`
	Status             int            `json:"status"`
	WorkOrderId        string         `json:"workOrderId"`
	WorkOrderItemId    string         `json:"workOrderItemId"`
	SalesEntryId       string         `json:"salesEntryId"`
	OrderStatus        int            `json:"orderStatus"`
	ToOndemand         string         `json:"toOndemand"`
	ItemValue          string         `json:"itemValue"`
	DoubleItemValue    float64        `json:"doubleItemValue"`
	ChargingStatus     int            `json:"chargingStatus"`
	ChargingDate       int64          `json:"chargingDate"`
	ResourceConfig     string         `json:"resourceConfig"`
	AutoToOnDemand     bool           `json:"autoToOnDemand"`
	BuildingChannel    int            `json:"buildingChannel"`
	IsPlatformSpecific bool           `json:"isPlatformSpecific"`
	BillingOwner       int            `json:"billingOwner"`
	IsPackage          bool           `json:"isPackage"`
	CanRelease         bool           `json:"canRelease"`
	IsChargeOff        bool           `json:"isChargeOff"`
	IsPublicTest       int            `json:"isPublicTest"`
	AutoRenewStatus    int            `json:"autoRenewStatus"`
	Master             bool           `json:"master"`
	ResourceConfigMap  map[string]any `json:"resourceConfigMap"`
}
