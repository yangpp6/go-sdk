package test

import (
	"openapi-sdk-go/sdk"
	"openapi-sdk-go/sdk/common"
)

var TestCredential, _ = common.NewCredential("", "")
var ProdCredential, _ = common.NewCredential("", "")

const (
	TestNeiMeng8RegionId  = "81f7728662dd11ec810800155d307d5b"
	TestNeiMeng8Az1AzName = ""

	ProdHuaDong1RegionId  = "bb9fdb42056f11eda1610242ac110002"
	ProdHuaDong1Az1AzName = "cn-huadong1-jsnj1A-public-ctcloud"
)

var DevClient = BuildDevClient()

func BuildDevClient() *sdk.CtyunClient {
	var DevClientConfig = sdk.CtyunClientConfig{
		ApiHooks: []common.ApiHook{
			common.ConsoleLogApiHook,
		},
		HttpHooks: []common.HttpHook{
			common.PrintLogHttpHook{},
			common.AddUserAgentHttpHook{},
		},
	}
	return sdk.EnvOf(sdk.EnvironmentDev, DevClientConfig)
}

func BuildProdClient() *sdk.CtyunClient {
	var ProdClientConfig = sdk.CtyunClientConfig{
		ApiHooks: []common.ApiHook{
			common.ConsoleLogApiHook,
		},
		HttpHooks: []common.HttpHook{
			common.PrintLogHttpHook{},
			common.AddUserAgentHttpHook{},
		},
	}
	return sdk.EnvOf(sdk.EnvironmentProd, ProdClientConfig)
}

func BuildDevClientWithStub(stub common.StubApiHook) *sdk.CtyunClient {
	var DevClientConfig = sdk.CtyunClientConfig{
		ApiHooks: []common.ApiHook{
			common.ConsoleLogApiHook,
			stub.Stub,
		},
		HttpHooks: []common.HttpHook{
			common.PrintLogHttpHook{},
			common.AddUserAgentHttpHook{},
		},
	}
	return sdk.EnvOf(sdk.EnvironmentDev, DevClientConfig)
}
