package sdk

import (
	"crypto/tls"
	"errors"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
	"time"
)

type Environment string

const (
	EnvironmentDev  = "dev"
	EnvironmentTest = "test"
	EnvironmentProd = "prod"
)

var Environments = []Environment{
	EnvironmentDev,
	EnvironmentTest,
	EnvironmentProd,
}

type CtyunClient struct {
	sender *common.CtyunSender
	Apis   *Apis
}

// NewCtyunClient 新建环境
func NewCtyunClient(apiProxyEndPoint string, ctiamEndPoint string, ctVpcEndPoint string, ctImageEndPoint string, ctEcsEndPoint string, ctEbsEndPoint string, client *http.Client, cfg CtyunClientConfig) (*CtyunClient, error) {
	err := CheckEndpoint(apiProxyEndPoint)
	if err != nil {
		return nil, err
	}
	err = CheckEndpoint(ctiamEndPoint)
	if err != nil {
		return nil, err
	}
	err = CheckEndpoint(ctVpcEndPoint)
	if err != nil {
		return nil, err
	}
	err = CheckEndpoint(ctImageEndPoint)
	if err != nil {
		return nil, err
	}
	err = CheckEndpoint(ctEcsEndPoint)
	if err != nil {
		return nil, err
	}
	err = CheckEndpoint(ctEbsEndPoint)
	if err != nil {
		return nil, err
	}
	if client == nil {
		return nil, errors.New("client不能为空")
	}
	c := &CtyunClient{
		sender: &common.CtyunSender{
			ApiProxyEndPoint: apiProxyEndPoint,
			CtiamEndPoint:    ctiamEndPoint,
			CtVpcEndPoint:    ctVpcEndPoint,
			CtImageEndPoint:  ctImageEndPoint,
			CtEcsEndPoint:    ctEcsEndPoint,
			CtEbsEndPoint:    ctEbsEndPoint,
			HttpHooks:        cfg.HttpHooks,
			Client:           client,
		},
	}
	c.Apis = NewApis(c.sender, cfg.ApiHooks...)
	return c, nil
}

type CtyunClientConfig struct {
	ApiHooks  []common.ApiHook
	HttpHooks []common.HttpHook
}

var DefaultCtyunClientConfig = CtyunClientConfig{
	HttpHooks: []common.HttpHook{common.AddUserAgentHttpHook{}},
}

// EnvOf 通过指定环境构建
func EnvOf(env Environment, cfg CtyunClientConfig) *CtyunClient {
	switch env {
	case EnvironmentDev:
		fallthrough
	case EnvironmentTest:
		client, _ := NewCtyunClient(
			"apiproxy-global.ctapi-test.ctyun.cn",
			"ctiam-global.ctapi-test.ctyun.cn",
			"ctvpc-global.ctapi-test.ctyun.cn",
			"ctimage-global.ctapi-test.ctyun.cn",
			"ctecs-global.ctapi-test.ctyun.cn",
			"ebs-global.ctapi-test.ctyun.cn",
			&http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
				Timeout: 30 * time.Second,
			},
			cfg,
		)
		return client
	case EnvironmentProd:
		fallthrough
	default:
		client, _ := NewCtyunClient(
			"apiproxy-global.ctapi.ctyun.cn",
			"ctiam-global.ctapi.ctyun.cn",
			"ctvpc-global.ctapi.ctyun.cn",
			"ctimage-global.ctapi.ctyun.cn",
			"ctecs-global.ctapi.ctyun.cn",
			"ebs-global.ctapi.ctyun.cn",
			&http.Client{},
			cfg,
		)
		return client
	}
}

// CheckEndpoint 校验endpoint
func CheckEndpoint(endPoint string) error {
	if endPoint == "" {
		return errors.New("endPoint不能为空")
	}
	return nil
}
