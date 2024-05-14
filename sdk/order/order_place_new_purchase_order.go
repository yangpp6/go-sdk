package order

import (
	"context"
	"encoding/json"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
	"net/url"
)

// orderPlaceNewPurchaseOrderApi 通用订购接口
type orderPlaceNewPurchaseOrderApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewOrderOrderApi(client *common.CtyunSender) common.ApiHandler[OrderPlaceNewPurchaseOrderRequest, OrderPlaceNewPurchaseOrderResponse] {
	return &orderPlaceNewPurchaseOrderApi{
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v3/order/placeNewPurchaseOrder",
		},
		client: client,
	}
}

func (o *orderPlaceNewPurchaseOrderApi) Do(ctx context.Context, credential *common.Credential, req *OrderPlaceNewPurchaseOrderRequest) (*OrderPlaceNewPurchaseOrderResponse, common.CtyunRequestError) {
	// 构建请求
	var orders []orderPlaceNewPurchaseOrderOrderRealRequest
	for _, order := range req.Orders {
		var items []orderPlaceNewPurchaseOrderOrderItemRealRequest
		for _, item := range order.Items {
			items = append(items, orderPlaceNewPurchaseOrderOrderItemRealRequest{
				ResourceType: item.ResourceType,
				ServiceTag:   item.ServiceTag,
				Master:       item.Master,
				ItemValue:    item.ItemValue,
				ItemConfig:   item.ItemConfig,
				// IsSystemVolume: item.IsSystemVolume,
			})
		}
		orders = append(orders, orderPlaceNewPurchaseOrderOrderRealRequest{
			Items:       items,
			CycleType:   order.CycleType,
			InstanceCnt: order.InstanceCnt,
			CycleCnt:    order.CycleCnt,
		})
	}
	bytes, e := json.Marshal(&orderPlaceNewPurchaseOrderRealRequest{
		AutoPay: req.AutoPay,
		Orders:  orders,
	})
	if e != nil {
		return nil, common.ErrorBeforeRequest(e)
	}
	builder := o.WithCredential(credential)
	values := url.Values{}
	values.Set("orderDetailJson", string(bytes))
	builder = builder.WriteXWwwFormUrlencoded(values)

	// 发起请求
	response, err := o.client.SendApiProxy(ctx, builder)
	if err != nil {
		return nil, err
	}

	// 解析返回
	result := &orderPlaceNewPurchaseOrderRealResponse{}
	err = response.ParseWithCheck(result)
	if err != nil {
		return nil, err
	}
	return &OrderPlaceNewPurchaseOrderResponse{
		ErrorMessage: result.ErrorMessage,
		Submitted:    result.Submitted,
		NewOrderId:   result.NewOrderId,
		NewOrderNo:   result.NewOrderNo,
		TotalPrice:   result.TotalPrice,
	}, nil
}

type orderPlaceNewPurchaseOrderOrderItemRealRequest struct {
	ResourceType string  `json:"resourceType"`
	ServiceTag   string  `json:"serviceTag"`
	Master       bool    `json:"master"`
	ItemValue    float64 `json:"itemValue"`
	// IsSystemVolume bool    `json:"isSystemVolume"`
	ItemConfig any `json:"itemConfig"`
}

type orderPlaceNewPurchaseOrderOrderRealRequest struct {
	Items       []orderPlaceNewPurchaseOrderOrderItemRealRequest `json:"items"`
	CycleType   string                                           `json:"cycleType"`
	InstanceCnt int                                              `json:"instanceCnt"`
	CycleCnt    int                                              `json:"cycleCnt"`
}

type orderPlaceNewPurchaseOrderRealRequest struct {
	AutoPay bool                                         `json:"autoPay"`
	Orders  []orderPlaceNewPurchaseOrderOrderRealRequest `json:"orders"`
}

type orderPlaceNewPurchaseOrderRealResponse struct {
	ErrorMessage string  `json:"errorMessage"`
	Submitted    bool    `json:"submitted"`
	NewOrderId   string  `json:"newOrderId"`
	NewOrderNo   string  `json:"newOrderNo"`
	TotalPrice   float64 `json:"totalPrice"`
}

type OrderPlaceNewPurchaseOrderOrderItemRequest struct {
	ResourceType   string
	ServiceTag     string
	Master         bool
	ItemValue      float64
	ItemConfig     any
	IsSystemVolume bool
}

type OrderPlaceNewPurchaseOrderOrderRequest struct {
	Items       []OrderPlaceNewPurchaseOrderOrderItemRequest
	CycleType   string
	InstanceCnt int
	CycleCnt    int
}

type OrderPlaceNewPurchaseOrderRequest struct {
	AutoPay bool
	Orders  []OrderPlaceNewPurchaseOrderOrderRequest
}

type OrderPlaceNewPurchaseOrderResponse struct {
	ErrorMessage string
	Submitted    bool
	NewOrderId   string
	NewOrderNo   string
	TotalPrice   float64
}
