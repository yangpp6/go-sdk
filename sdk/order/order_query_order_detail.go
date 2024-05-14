package order

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// orderQueryOrderDetailApi 通用订购接口
type orderQueryOrderDetailApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewOrderQueryOrderDetailApi(client *common.CtyunSender) common.ApiHandler[OrderQueryOrderDetailRequest, OrderQueryOrderDetailResponse] {
	return &orderQueryOrderDetailApi{
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v3/order/queryOrderDetail",
		},
		client: client,
	}
}

func (o *orderQueryOrderDetailApi) Do(ctx context.Context, credential *common.Credential, req *OrderQueryOrderDetailRequest) (*OrderQueryOrderDetailResponse, common.CtyunRequestError) {
	// 构建请求
	builder := o.WithCredential(credential)
	builder.AddHeader("masterOrderId", req.MasterOrderId)

	// 发起请求
	response, err := o.client.SendApiProxy(ctx, builder)
	if err != nil {
		return nil, err
	}

	// 解析返回
	result := &OrderQueryOrderDetailResponse{}
	err = response.ParseWithCheck(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type OrderQueryOrderDetailRequest struct {
	MasterOrderId string
}

type OrderQueryOrderDetailResponse struct {
	MasterOrderId      string `json:"masterOrderId"`
	MasterOrderNo      string `json:"masterOrderNo"`
	MasterOrderOrigId  string `json:"masterOrderOrigId"`
	MasterOrderType    int    `json:"masterOrderType"`
	AccountId          string `json:"accountId"`
	UserId             string `json:"userId"`
	TotalPrice         string `json:"totalPrice"`
	PayType            string `json:"payType"`
	CreateDate         int64  `json:"createDate"`
	UpdateDate         int64  `json:"updateDate"`
	PackageType        int    `json:"packageType"`
	IsVirtualOrder     string `json:"isVirtualOrder"`
	IsTrialOrder       int    `json:"isTrialOrder"`
	IsAgencyOrder      string `json:"isAgencyOrder"`
	ChargingStatus     int    `json:"chargingStatus"`
	IsOnDemand         bool   `json:"isOnDemand"`
	HasContract        bool   `json:"hasContract"`
	Cancelable         bool   `json:"cancelable"`
	Payable            bool   `json:"payable"`
	CanApplyTrial      bool   `json:"canApplyTrial"`
	CanForceToComplete bool   `json:"canForceToComplete"`
	PaymentSite        int    `json:"paymentSite"`
	AbandonStatus      int    `json:"abandonStatus"`
	Status             int    `json:"status"`
	OrderEventType     int    `json:"orderEventType"`
	Orders             []struct {
		OrderId     string `json:"orderId"`
		OrderNo     string `json:"orderNo"`
		OrderOrigId string `json:"orderOrigId"`
		InstanceCnt int    `json:"instanceCnt"`
		CycleCnt    int    `json:"cycleCnt"`
		CycleType   int    `json:"cycleType"`
		OrderConfig struct {
		} `json:"orderConfig"`
		TotalPrice string `json:"totalPrice"`
		ServiceTag string `json:"serviceTag"`
		OrderItems []struct {
			ItemId          string  `json:"itemId"`
			ItemNo          string  `json:"itemNo"`
			ItemValue       float64 `json:"itemValue"`
			OrderItemConfig struct {
				ZoneType         string `json:"zoneType"`
				AvailabilityZone string `json:"availability_zone"`
				Value            string `json:"value"`
				Number           string `json:"number"`
				VolumeType       string `json:"volumeType"`
				ZoneId           string `json:"zoneId"`
				RegionId         string `json:"regionId"`
				Size             string `json:"size"`
				VolumeName       string `json:"volumeName"`
			} `json:"orderItemConfig"`
			ServiceTag   string `json:"serviceTag"`
			Cost         string `json:"cost"`
			IsMaster     bool   `json:"isMaster"`
			ResourceType string `json:"resourceType"`
		} `json:"orderItems"`
	} `json:"orders"`
	OrderChannel int `json:"orderChannel"`
}
