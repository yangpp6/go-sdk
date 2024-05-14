package keypair

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

// keypairCreateApi 创建一对SSH密钥对
type keypairCreateApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewKeypairCreateApi(client *common.CtyunSender) common.ApiHandler[KeypairCreateRequest, KeypairCreateResponse] {
	return &keypairCreateApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/keypair/create-keypair",
		},
	}
}

func (v *keypairCreateApi) Do(ctx context.Context, credential *common.Credential, req *KeypairCreateRequest) (*KeypairCreateResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(keypairCreateRealRequest{
		RegionID:    req.RegionId,
		KeypairName: req.KeyPairName,
		ProjectID:   req.ProjectId,
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
	result := &KeypairCreateRealResponse{}
	err = response.ParseWithCheck(result)
	if err != nil {
		return nil, err
	}
	return &KeypairCreateResponse{
		PublicKey:   result.PublicKey,
		PrivateKey:  result.PrivateKey,
		KeyPairName: result.KeyPairName,
		FingerPrint: result.FingerPrint,
		KeyPairID:   result.KeyPairID,
	}, nil
}

type keypairCreateRealRequest struct {
	RegionID    string `json:"regionID"`
	KeypairName string `json:"keyPairName"`
	ProjectID   string `json:"projectID"`
}

type KeypairCreateRealResponse struct {
	PublicKey   string `json:"publicKey"`
	PrivateKey  string `json:"privateKey"`
	KeyPairName string `json:"keyPairName"`
	FingerPrint string `json:"fingerPrint"`
	KeyPairID   string `json:"keyPairID"`
}

type KeypairCreateRequest struct {
	RegionId    string //资源池ID
	KeyPairName string //密钥对名称。只能由数字、字母、-组成,不能以数字和-开头、以-结尾,且长度为2-63字符
	ProjectId   string //企业项目ID
}

type KeypairCreateResponse struct {
	PublicKey   string //密钥对的公钥
	PrivateKey  string //密钥对的私钥
	KeyPairName string //密钥对名称
	FingerPrint string //密钥对的指纹，采用MD5信息摘要算法
	KeyPairID   string //密钥对的ID
}
