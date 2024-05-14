package common

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"
)

const (
	StatusCodeSuccess = "800"
)

type CtyunRequest struct {
	method      string
	urlPath     string
	endpoint    string
	credential  Credential
	headers     http.Header
	params      url.Values
	body        []byte
	contentType string
}

type CtyunRequestBuilder struct {
	Method  string
	UrlPath string
}

type CtyunResponse struct {
	Request  *http.Request
	Response *http.Response
}

type CtyunResponseModel struct {
	StatusCode string `json:"-"` // 注意这里的类型值是有问题的，成功的时候返回int类型的800，错误的时候返回一个string值，因此不能做json的反序列化操作，wtf
	Message    string `json:"message"`
	ReturnObj  any    `json:"returnObj,omitempty"`
}

// WithCredential 增加请求credential
func (c CtyunRequestBuilder) WithCredential(credential *Credential) *CtyunRequest {
	result := &CtyunRequest{
		method:  c.Method,
		urlPath: c.UrlPath,
		headers: make(http.Header),
		params:  make(url.Values),
	}
	if credential != nil {
		result.credential = *credential
	}
	return result
}

// AddHeader 增加请求头
func (request *CtyunRequest) AddHeader(key, value string) *CtyunRequest {
	request.headers.Add(key, value)
	return request
}

// AddParam 增加参数
func (request *CtyunRequest) AddParam(key, value string) *CtyunRequest {
	request.params.Add(key, value)
	return request
}

// WriteXWwwFormUrlencoded 以x-www-form-urlencoded方式写入
func (request *CtyunRequest) WriteXWwwFormUrlencoded(data url.Values) *CtyunRequest {
	encode := data.Encode()
	request.body = []byte(encode)
	request.contentType = "application/x-www-form-urlencoded"
	return request
}

// WriteJson 以application/json方式写入
func (request *CtyunRequest) WriteJson(data any) (*CtyunRequest, CtyunRequestError) {
	marshal, err := json.Marshal(data)
	if err != nil {
		return nil, ErrorBeforeRequest(err)
	}
	request.body = marshal
	request.contentType = "application/json"
	return request, nil
}

// buildRequest 构造请求
func (request CtyunRequest) buildRequest(endPoint string) (*http.Request, CtyunRequestError) {
	// 构造url
	u := url.URL{}
	u.Scheme = "https"
	u.Host = endPoint
	u.Path = request.urlPath
	u.RawQuery = request.params.Encode()

	// 构造请求头
	tim := time.Now()
	eopDate := tim.Format("20060102T150405Z")
	id := uuid.NewString()
	sign := GetSign(u.RawQuery, request.body, tim, id, request.credential)
	headers := request.headers.Clone()
	headers.Add("ctyun-eop-request-id", id)
	headers.Add("Eop-Authorization", sign)
	headers.Add("Eop-date", eopDate)
	if request.contentType != "" {
		headers.Add("Content-Type", request.contentType)
	}
	if request.body != nil {
		headers.Add("Content-Length", strconv.Itoa(len(request.body)))
	}

	// 构造实际请求
	req, err := http.NewRequest(request.method, u.String(), bytes.NewReader(request.body))
	if err != nil {
		return nil, ErrorBeforeRequest(err)
	}
	req.Header = headers
	return req, nil
}

// GetSign 加签
func GetSign(
	query string,
	body []byte,
	tim time.Time,
	uuid string,
	credential Credential,
) string {
	hash := sha256.New()
	hash.Write(body)
	sum := hash.Sum(nil)
	calculateContentHash := hex.EncodeToString(sum)
	date := tim.Format("20060102T150405Z")
	sigture := fmt.Sprintf("ctyun-eop-request-id:%s\neop-date:%s\n\n%s\n%s", uuid, date, query, calculateContentHash)
	singerDd := tim.Format("20060102")
	kAk := hmacSHA256(credential.Ak(), string(hmacSHA256(date, credential.Sk())))
	kdate := hmacSHA256(singerDd, string(kAk))
	signaSha256 := hmacSHA256(sigture, string(kdate))
	Signature := base64.StdEncoding.EncodeToString(signaSha256)
	signHeader := credential.Ak() + " Headers=ctyun-eop-request-id;eop-date Signature=" + Signature
	return signHeader
}

// hmacSHA256 HmacSHA256加密
func hmacSHA256(signature, key string) []byte {
	s := []byte(signature)
	k := []byte(key)
	m := hmac.New(sha256.New, k)
	m.Write(s)
	sum := m.Sum(nil)
	return sum
}

// ParseWithCheck 解析并且判断
func (c CtyunResponse) ParseWithCheck(obj any) CtyunRequestError {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(c.Response.Body)
	respBody, err := io.ReadAll(c.Response.Body)
	c.Response.Body = io.NopCloser(bytes.NewBuffer(respBody))
	if err != nil {
		return WrapError(err, &c)
	}

	var content CtyunResponseModel
	bodyContent := string(respBody)
	parse := gjson.Parse(bodyContent)
	statusCode := parse.Get("statusCode")
	if !IsSuccess(statusCode.String()) {
		description := parse.Get("description")
		message := parse.Get("message")

		code := parse.Get("errorCode")
		err := errors.New(description.String() + "，" + message.String())
		wrapError := WrapWithErrorCode(err, code.String(), &c)
		return wrapError
	}
	content.ReturnObj = obj
	err = json.Unmarshal(respBody, &content)
	if err != nil {
		return WrapError(err, &c)
	}
	content.StatusCode = statusCode.String()
	return nil
}

// IsSuccess 判断响应是否成功
func IsSuccess(result string) bool {
	return result == StatusCodeSuccess
}

// Parse 解析为目标对象
func (c CtyunResponse) Parse(target any) error {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(c.Response.Body)
	respBody, err := io.ReadAll(c.Response.Body)
	c.Response.Body = io.NopCloser(bytes.NewBuffer(respBody))
	if err != nil {
		return err
	}
	return json.Unmarshal(respBody, target)
}

type HttpHook interface {
	BeforeRequest(context.Context, *http.Request)
	AfterResponse(context.Context, *http.Response)
}

type PrintLogHttpHook struct {
}

func (d PrintLogHttpHook) BeforeRequest(ctx context.Context, request *http.Request) {
	dumpRequest, err := httputil.DumpRequest(request, true)
	if err != nil {
		return
	}
	requestContent := string(dumpRequest)
	fmt.Printf("实际请求内容：\n%s\n", requestContent)
}

func (d PrintLogHttpHook) AfterResponse(ctx context.Context, response *http.Response) {
	dumpResponse, err := httputil.DumpResponse(response, true)
	if err != nil {
		return
	}
	responseContent := string(dumpResponse)
	fmt.Printf("实际请求返回：\n%s\n", responseContent)
}

type AddUserAgentHttpHook struct {
}

func (h AddUserAgentHttpHook) BeforeRequest(ctx context.Context, request *http.Request) {
	// 不添加请求头会出现被风控的现象
	if request.Header.Get("User-Agent") == "" {
		request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	}
}

func (h AddUserAgentHttpHook) AfterResponse(ctx context.Context, response *http.Response) {
}
