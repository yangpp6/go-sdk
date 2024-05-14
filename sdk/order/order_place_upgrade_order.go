package order

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"openapi-sdk-go/sdk/common"
)

// orderPlaceUpgradeOrderApi 通用升配接口
type orderPlaceUpgradeOrderApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewOrderPlaceUpgradeOrderApi(client *common.CtyunSender) common.ApiHandler[OrderPlaceUpgradeOrderRequest, OrderPlaceUpgradeOrderResponse] {
	return &orderPlaceUpgradeOrderApi{
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v3/order/placeUpgradeOrder",
		},
		client: client,
	}
}

func (o *orderPlaceUpgradeOrderApi) Do(ctx context.Context, credential *common.Credential, req *OrderPlaceUpgradeOrderRequest) (*OrderPlaceUpgradeOrderResponse, common.CtyunRequestError) {
	// 构建请求
	var orders []orderPlaceUpgradeOrderOrderRealRequest
	for _, order := range req.Orders {
		var items []orderPlaceUpgradeOrderOrderItemRealRequest
		for _, item := range order.Items {
			items = append(items, orderPlaceUpgradeOrderOrderItemRealRequest{
				ResourceType:   item.ResourceType,
				ServiceTag:     item.ServiceTag,
				Master:         item.Master,
				ItemValue:      item.ItemValue,
				IsSystemVolume: item.IsSystemVolume,
				ItemConfig:     item.ItemConfig,
			})
		}
		orders = append(orders, orderPlaceUpgradeOrderOrderRealRequest{
			Items: items,
		})
	}
	bytes, e := json.Marshal(&orderPlaceUpgradeOrderRealRequest{
		Orders: orders,
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
	result := &orderPlaceUpgradeOrderRealResponse{}
	err = response.ParseWithCheck(result)
	if err != nil {
		return nil, err
	}

	return &OrderPlaceUpgradeOrderResponse{
		ErrorMessage: result.ErrorMessage,
		Submitted:    result.Submitted,
		NewOrderId:   result.NewOrderId,
		NewOrderNo:   result.NewOrderNo,
		TotalPrice:   result.TotalPrice,
	}, nil
}

type orderPlaceUpgradeOrderOrderItemRealRequest struct {
	ResourceType   string `json:"resourceType"`
	ServiceTag     string `json:"serviceTag"`
	Master         bool   `json:"master"`
	ItemValue      int    `json:"itemValue"`
	IsSystemVolume bool   `json:"isSystemVolume"`
	ItemConfig     any    `json:"itemConfig"`
}

type orderPlaceUpgradeOrderOrderRealRequest struct {
	Items []orderPlaceUpgradeOrderOrderItemRealRequest `json:"items"`
}

type orderPlaceUpgradeOrderRealRequest struct {
	Orders []orderPlaceUpgradeOrderOrderRealRequest `json:"orders"`
}

type orderPlaceUpgradeOrderRealResponse struct {
	ErrorMessage string  `json:"errorMessage"`
	Submitted    bool    `json:"submitted"`
	NewOrderId   string  `json:"newOrderId"`
	NewOrderNo   string  `json:"newOrderNo"`
	TotalPrice   float64 `json:"totalPrice"`
}

type OrderPlaceUpgradeOrderOrderItemRequest struct {
	ResourceType   string
	ServiceTag     string
	Master         bool
	ItemValue      int
	IsSystemVolume bool
	ItemConfig     any
}

type OrderPlaceUpgradeOrderOrderRequest struct {
	Items []OrderPlaceUpgradeOrderOrderItemRequest
}

type OrderPlaceUpgradeOrderRequest struct {
	Orders []OrderPlaceUpgradeOrderOrderRequest
}

type OrderPlaceUpgradeOrderResponse struct {
	ErrorMessage string
	Submitted    bool
	NewOrderId   string
	NewOrderNo   string
	TotalPrice   float64
}
