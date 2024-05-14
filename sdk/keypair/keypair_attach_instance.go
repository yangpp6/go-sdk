package keypair

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// keypairAttachApi 绑定密钥对
type keypairAttachApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewKeypairAttachApi(client *common.CtyunSender) common.ApiHandler[KeypairAttachRequest, KeypairAttachResponse] {
	return &keypairAttachApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/keypair/attach-instance",
		},
	}
}

func (v *keypairAttachApi) Do(ctx context.Context, credential *common.Credential, req *KeypairAttachRequest) (*KeypairAttachResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(&keypairAttachRealRequest{
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
	return &KeypairAttachResponse{}, nil
}

type keypairAttachRealRequest struct {
	RegionID    string `json:"regionID"`
	KeyPairName string `json:"keyPairName"`
	InstanceID  string `json:"instanceID"`
}

type KeypairAttachRequest struct {
	RegionId    string // 区域id
	KeyPairName string // 密钥对名称
	InstanceId  string // 云主机ID
}

type KeypairAttachResponse struct {
}
