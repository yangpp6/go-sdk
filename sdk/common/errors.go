package common

import (
	"net/http"
)

const (
	OccurOccasionBeforeRequest  = 0
	OccurOccasionBeforeResponse = 1
	OccurOccasionAfterResponse  = 2
)

type OccurOccasion int

type CtyunRequestError interface {
	error                         // 错误信息
	OccurOccasion() OccurOccasion // 发生场合
	Request() *http.Request       // 请求
	Response() *http.Response     // 返回
	ErrorCode() string            // 返回码
}

type ctyunRequestError struct {
	errorMessage  string
	occurOccasion OccurOccasion // 异常发生场合
	request       *http.Request
	response      *http.Response
	errorCode     string
}

func (c ctyunRequestError) Error() string {
	return c.errorMessage
}

func (c ctyunRequestError) OccurOccasion() OccurOccasion {
	return c.occurOccasion
}

func (c ctyunRequestError) Request() *http.Request {
	return c.request
}

func (c ctyunRequestError) Response() *http.Response {
	return c.response
}

func (c ctyunRequestError) ErrorCode() string {
	return c.errorCode
}

// WrapError 包裹异常
func WrapError(err error, resp *CtyunResponse) CtyunRequestError {
	occurOccasion := 0
	if resp.Request != nil {
		occurOccasion++
	}
	if resp.Response != nil {
		occurOccasion++
	}
	return ctyunRequestError{
		errorMessage:  err.Error(),
		occurOccasion: OccurOccasion(occurOccasion),
		request:       resp.Request,
		response:      resp.Response,
	}
}

func ErrorBeforeRequest(err error) CtyunRequestError {
	return ctyunRequestError{
		errorMessage:  err.Error(),
		occurOccasion: OccurOccasionBeforeRequest,
	}
}

func ErrorBeforeResponse(err error, request *http.Request) CtyunRequestError {
	return ctyunRequestError{
		errorMessage:  err.Error(),
		occurOccasion: OccurOccasionBeforeResponse,
		request:       request,
	}
}

func ErrorAfterResponse(err error, request *http.Request, response *http.Response) CtyunRequestError {
	return ctyunRequestError{
		errorMessage:  err.Error(),
		occurOccasion: OccurOccasionAfterResponse,
		request:       request,
		response:      response,
	}
}

// WrapWithErrorCode 包裹异常
func WrapWithErrorCode(err error, errorCode string, resp *CtyunResponse) CtyunRequestError {
	wrapError := WrapError(err, resp)
	requestError := any(wrapError).(ctyunRequestError)
	requestError.errorCode = errorCode
	return requestError
}
