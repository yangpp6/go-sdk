package common

import (
	"context"
	"net/http"
)

type CtyunSender struct {
	ApiProxyEndPoint string
	CtiamEndPoint    string
	CtVpcEndPoint    string
	CtEcsEndPoint    string
	CtImageEndPoint  string
	CtEbsEndPoint    string

	Client    *http.Client
	HttpHooks []HttpHook
}

// SendApiProxy 对apiProxy发起请求
func (client CtyunSender) SendApiProxy(ctx context.Context, request *CtyunRequest) (*CtyunResponse, CtyunRequestError) {
	return client.requestToEndpoint(ctx, request, client.ApiProxyEndPoint)
}

// SendCtVpc 对ctvpc发起请求
func (client CtyunSender) SendCtVpc(ctx context.Context, request *CtyunRequest) (*CtyunResponse, CtyunRequestError) {
	return client.requestToEndpoint(ctx, request, client.CtVpcEndPoint)
}

// SendCtEcs 对ctecs发起请求
func (client CtyunSender) SendCtEcs(ctx context.Context, request *CtyunRequest) (*CtyunResponse, CtyunRequestError) {
	return client.requestToEndpoint(ctx, request, client.CtEcsEndPoint)
}

// SendCtEcs 对ctebs发起请求
func (client CtyunSender) SendCtEbs(ctx context.Context, request *CtyunRequest) (*CtyunResponse, CtyunRequestError) {
	req, err := request.buildRequest(client.CtEbsEndPoint)
	if err != nil {
		return nil, err
	}
	return client.send(ctx, req)
}

// SendCtVpc 对ctimage发起请求
func (client CtyunSender) SendCtImage(ctx context.Context, request *CtyunRequest) (*CtyunResponse, CtyunRequestError) {
	return client.requestToEndpoint(ctx, request, client.CtImageEndPoint)
}

// SendCtiam 对ctiam发起请求
func (client CtyunSender) SendCtiam(ctx context.Context, request *CtyunRequest) (*CtyunResponse, CtyunRequestError) {
	return client.requestToEndpoint(ctx, request, client.CtiamEndPoint)
}

// 向端点发送请求
func (client CtyunSender) requestToEndpoint(ctx context.Context, request *CtyunRequest, endpoint string) (*CtyunResponse, CtyunRequestError) {
	req, err := request.buildRequest(endpoint)
	if err != nil {
		return nil, ErrorBeforeResponse(err, req)
	}
	return client.send(ctx, req)
}

// send 发送请求
func (client CtyunSender) send(ctx context.Context, req *http.Request) (*CtyunResponse, CtyunRequestError) {
	for _, hook := range client.HttpHooks {
		hook.BeforeRequest(ctx, req)
	}
	resp, err := client.Client.Do(req)
	for _, hook := range client.HttpHooks {
		hook.AfterResponse(ctx, resp)
	}
	if err != nil {
		return nil, ErrorAfterResponse(err, req, resp)
	}
	return &CtyunResponse{Request: req, Response: resp}, nil
}
