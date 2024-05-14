package keypair

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// keypairDetachApi 解绑密钥对
type keypairDetachApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewKeypairDetachApi(client *common.CtyunSender) common.ApiHandler[KeypairDetachRequest, KeypairDetachResponse] {
	return &keypairDetachApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/keypair/detach-instance",
		},
	}
}

func (v *keypairDetachApi) Do(ctx context.Context, credential *common.Credential, req *KeypairDetachRequest) (*KeypairDetachResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&keypairDetachRealRequest{
		RegionID:    req.RegionId,
		KeyPairName: req.KeyPairName,
		InstanceID:  req.InstanceId,
	})
	if err != nil {
		return nil, err
	}

	// 发起请求
	response, err := v.client.SendCtEcs(ctx, builder)
	if err != nil {
		return nil, err
	}

	// 解析返回
	err = response.ParseWithCheck(nil)
	if err != nil {
		return nil, err
	}
	return &KeypairDetachResponse{}, nil
}

type keypairDetachRealRequest struct {
	RegionID    string `json:"regionID"`
	KeyPairName string `json:"keyPairName"`
	InstanceID  string `json:"instanceID"`
}

type KeypairDetachRequest struct {
	RegionId    string // 区域id
	KeyPairName string // 密钥对名称
	InstanceId  string // 云主机ID
}

type KeypairDetachResponse struct {
}
