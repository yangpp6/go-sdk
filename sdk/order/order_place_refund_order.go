package order

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"openapi-sdk-go/sdk/common"
	"strconv"
)

// orderPlaceRefundOrderApi 通用退订接口
type orderPlaceRefundOrderApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewOrderPlaceRefundOrderApi(client *common.CtyunSender) common.ApiHandler[OrderPlaceRefundOrderRequest, OrderPlaceRefundOrderResponse] {
	return &orderPlaceRefundOrderApi{
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v3/order/placeRefundOrder",
		},
		client: client,
	}
}

func (o *orderPlaceRefundOrderApi) Do(ctx context.Context, credential *common.Credential, req *OrderPlaceRefundOrderRequest) (*OrderPlaceRefundOrderResponse, common.CtyunRequestError) {
	// 构建请求
	var resources []orderPlaceRefundOrderResourceRealRequest
	for _, resource := range req.Resources {
		resources = append(resources, orderPlaceRefundOrderResourceRealRequest{
			ResourceIds:   resource.ResourceIds,
			RefundingCash: resource.RefundingCash,
		})
	}
	bytes, e := json.Marshal(orderPlaceRefundOrderRealRequest{
		Resources:     resources,
		RefundReason:  req.RefundReason,
		RefundingCash: req.RefundingCash,
		AutoApproval:  req.AutoApproval,
	})
	if e != nil {
		return nil, common.ErrorBeforeRequest(e)
	}
	builder := o.WithCredential(credential)
	values := url.Values{}
	values.Set("resourceDetailJson", string(bytes))
	values.Set("type", strconv.Itoa(2))
	builder = builder.WriteXWwwFormUrlencoded(values)

	// 发起请求
	response, err := o.client.SendApiProxy(ctx, builder)
	if err != nil {
		return nil, err
	}

	// 解析返回
	result := &orderPlaceRefundOrderRealResponse{}
	err = response.ParseWithCheck(result)
	if err != nil {
		return nil, err
	}

	var batchOrderPlacementResults []OrderPlaceRefundOrderBatchOrderPlacementResultResponse
	for _, placementResult := range result.BatchOrderPlacementResults {
		var orderPlacedEvents []OrderPlaceRefundOrderBatchOrderPlacementResultOrderPlacedEventResponse
		for _, event := range placementResult.OrderPlacedEvents {
			orderPlacedEvents = append(orderPlacedEvents, OrderPlaceRefundOrderBatchOrderPlacementResultOrderPlacedEventResponse{
				ErrorMessage: event.ErrorMessage,
				Submitted:    event.Submitted,
				NewOrderId:   event.NewOrderId,
				NewOrderNo:   event.NewOrderNo,
				TotalPrice:   event.TotalPrice,
			})
		}
		batchOrderPlacementResults = append(batchOrderPlacementResults, OrderPlaceRefundOrderBatchOrderPlacementResultResponse{
			ErrorMessage:      placementResult.ErrorMessage,
			Submitted:         placementResult.Submitted,
			OrderPlacedEvents: orderPlacedEvents,
		})
	}
	return &OrderPlaceRefundOrderResponse{
		ErrorMessage:               result.ErrorMessage,
		BatchOrderPlacementResults: batchOrderPlacementResults,
	}, nil
}

type orderPlaceRefundOrderResourceRealRequest struct {
	ResourceIds   []string `json:"resourceIds"`
	RefundingCash string   `json:"refundingCash"`
}

type orderPlaceRefundOrderRealRequest struct {
	Resources     []orderPlaceRefundOrderResourceRealRequest `json:"resources"`
	ResourceIds   []string                                   `json:"resourceIds"`
	RefundReason  string                                     `json:"refundReason"`
	RefundingCash string                                     `json:"refundingCash"`
	AutoApproval  bool                                       `json:"autoApproval"`
}

type orderPlaceRefundOrderRealResponse struct {
	ErrorMessage               string `json:"errorMessage"`
	BatchOrderPlacementResults []struct {
		ErrorMessage      string `json:"errorMessage"`
		Submitted         bool   `json:"submitted"`
		OrderPlacedEvents []struct {
			ErrorMessage string  `json:"errorMessage"`
			Submitted    bool    `json:"submitted"`
			NewOrderId   string  `json:"newOrderId"`
			NewOrderNo   string  `json:"newOrderNo"`
			TotalPrice   float64 `json:"totalPrice"`
		} `json:"orderPlacedEvents"`
	} `json:"batchOrderPlacementResults"`
}

type OrderPlaceRefundOrderBatchOrderPlacementResultOrderPlacedEventResponse struct {
	ErrorMessage string
	Submitted    bool
	NewOrderId   string
	NewOrderNo   string
	TotalPrice   float64
}

type OrderPlaceRefundOrderBatchOrderPlacementResultResponse struct {
	ErrorMessage      string
	Submitted         bool
	OrderPlacedEvents []OrderPlaceRefundOrderBatchOrderPlacementResultOrderPlacedEventResponse
}

type OrderPlaceRefundOrderResponse struct {
	ErrorMessage               string
	BatchOrderPlacementResults []OrderPlaceRefundOrderBatchOrderPlacementResultResponse
}

type OrderPlaceRefundOrderResourceRequest struct {
	ResourceIds   []string
	RefundingCash string
}

type OrderPlaceRefundOrderRequest struct {
	Resources     []OrderPlaceRefundOrderResourceRequest
	ResourceIds   []string
	RefundReason  string
	RefundingCash string
	AutoApproval  bool
}
