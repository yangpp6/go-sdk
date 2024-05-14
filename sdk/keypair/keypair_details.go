package keypair

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// keypairDetailsApi 创建一对SSH密钥对
type keypairDetailApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewKeypairDetailApi(client *common.CtyunSender) common.ApiHandler[KeypairDetailRequest, KeypairDetailResponse] {
	return &keypairDetailApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/keypair/details",
		},
	}
}

func (v *keypairDetailApi) Do(ctx context.Context, credential *common.Credential, req *KeypairDetailRequest) (*KeypairDetailResponse, common.CtyunRequestError) {
	// 构建请求
	builder := v.WithCredential(credential)
	_, err := builder.WriteJson(keypairDetailRealRequest{
		RegionID:     req.RegionId,
		ProjectID:    req.ProjectID,
		KeypairName:  req.KeyPairName,
		QueryContent: req.QueryContent,
		PageNo:       req.PageNo,
		PageSize:     req.PageSize,
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
	result := &KeypairDetailRealResponse{}
	err = response.ParseWithCheck(&result)

	var keypairDetail []KeypairDetailResults
	for _, detail := range result.Results {
		keypairDetail = append(keypairDetail, KeypairDetailResults{
			PublicKey:   detail.PublicKey,
			KeyPairName: detail.KeyPairName,
			FingerPrint: detail.FingerPrint,
			KeyPairID:   detail.KeyPairID,
			ProjectID:   detail.ProjectID,
		})
	}
	return &KeypairDetailResponse{
		CurrentCount: result.CurrentCount,
		TotalCount:   result.TotalCount,
		Results:      keypairDetail,
	}, err
}

type keypairDetailRealRequest struct {
	RegionID     string `json:"regionID"`
	ProjectID    string `json:"projectID"`
	KeypairName  string `json:"keyPairName"`
	QueryContent string `json:"queryContent"`
	PageNo       int    `json:"pageNo"`
	PageSize     int    `json:"pageSize"`
}

type KeypairDetailRealResponse struct {
	CurrentCount int `json:"currentCount"`
	TotalCount   int `json:"totalCount"`
	Results      []struct {
		PublicKey   string `json:"publicKey"`
		KeyPairName string `json:"keyPairName"`
		FingerPrint string `json:"fingerPrint"`
		KeyPairID   string `json:"keyPairID"`
		ProjectID   string `json:"projectID"`
	} `json:"results"`
}

type KeypairDetailResults struct {
	PublicKey   string
	KeyPairName string
	FingerPrint string
	KeyPairID   string
	ProjectID   string
}

type KeypairDetailRequest struct {
	RegionId     string // 资源池ID
	ProjectID    string // 企业项目ID
	KeyPairName  string // 密钥对名称。只能由数字、字母、-组成,不能以数字和-开头、以-结尾,且长度为2-63字符
	QueryContent string
	PageNo       int
	PageSize     int
}

type KeypairDetailResponse struct {
	CurrentCount int // 当前页记录数目
	TotalCount   int // 总记录数
	Results      []KeypairDetailResults
}
