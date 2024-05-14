package keypair

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// keypairDeleteApi 删除一对SSH密钥对
type keypairDeleteApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewKeypairDeleteApi(client *common.CtyunSender) common.ApiHandler[KeypairDeleteRequest, KeypairDeleteResponse] {
	return &keypairDeleteApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/keypair/delete",
		},
	}
}

func (v *keypairDeleteApi) Do(ctx context.Context, credential *common.Credential, req *KeypairDeleteRequest) (*KeypairDeleteResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(keypairDeleteRealRequest{
		RegionID:    req.RegionId,
		KeyPairName: req.KeyPairName,
	})
	if err != nil {
		return nil, err
	}

	//发起请求
	response, err := v.client.SendCtEcs(ctx, builder)
	if err != nil {
		return nil, err
	}

	// 解析返回
	result := &KeypairDeleteRealResponse{}
	err = response.ParseWithCheck(result)
	if err != nil {
		return nil, err
	}
	return &KeypairDeleteResponse{
		KeyPairName: result.KeyPairName,
	}, nil
}

type keypairDeleteRealRequest struct {
	RegionID    string `json:"regionID"`
	KeyPairName string `json:"keyPairName"`
}

type KeypairDeleteRealResponse struct {
	KeyPairName string `json:"keyPairName"`
}

type KeypairDeleteRequest struct {
	RegionId    string //资源池ID
	KeyPairName string //密钥对名称。
}

type KeypairDeleteResponse struct {
	KeyPairName string //密钥对名称
}
