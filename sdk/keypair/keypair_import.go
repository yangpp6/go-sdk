package keypair

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// keypairImportApi 创建一对SSH密钥对
type keypairImportApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewKeypairImportApi(client *common.CtyunSender) common.ApiHandler[KeypairImportRequest, KeypairImportResponse] {
	return &keypairImportApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/keypair/import-keypair",
		},
	}
}

func (v *keypairImportApi) Do(ctx context.Context, credential *common.Credential, req *KeypairImportRequest) (*KeypairImportResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(keypairImportRealRequest{
		RegionID:    req.RegionId,
		KeypairName: req.KeyPairName,
		PublicKey:   req.PublicKey,
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
	result := &KeypairImportRealResponse{}
	err = response.ParseWithCheck(result)
	if err != nil {
		return nil, err
	}
	return &KeypairImportResponse{
		PublicKey:   result.PublicKey,
		KeyPairName: result.KeyPairName,
		FingerPrint: result.FingerPrint,
		KeyPairId:   result.KeyPairID,
	}, nil
}

type keypairImportRealRequest struct {
	RegionID    string `json:"regionID"`
	KeypairName string `json:"keyPairName"`
	PublicKey   string `json:"publicKey"`
	ProjectID   string `json:"projectID,omitempty"`
}

type KeypairImportRealResponse struct {
	PublicKey   string `json:"publicKey"`
	KeyPairName string `json:"keyPairName"`
	FingerPrint string `json:"fingerPrint"`
	KeyPairID   string `json:"keyPairID"`
}

type KeypairImportRequest struct {
	RegionId    string // 资源池ID
	KeyPairName string // 密钥对名称。只能由数字、字母、-组成,不能以数字和-开头、以-结尾,且长度为2-63字符
	PublicKey   string // 公钥
}

type KeypairImportResponse struct {
	PublicKey   string // 密钥对的公钥
	KeyPairName string // 密钥对名称
	FingerPrint string // 密钥对的指纹，采用MD5信息摘要算法
	KeyPairId   string // 密钥对的ID
}
