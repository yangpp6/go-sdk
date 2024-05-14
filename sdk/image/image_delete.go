package image

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// imageDeleteApi 删除私有镜像
type imageDeleteApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewImageDeleteApi(client *common.CtyunSender) common.ApiHandler[ImageDeleteRequest, ImageDeleteResponse] {
	return &imageDeleteApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/image/delete",
		},
	}
}

func (v *imageDeleteApi) Do(ctx context.Context, credential *common.Credential, req *ImageDeleteRequest) (*ImageDeleteResponse, common.CtyunRequestError) {
	// 构建请求
	realRequest := v.WithCredential(credential)
	_, err := realRequest.WriteJson(&imageDeleteRealRequest{
		RegionID: req.RegionId,
		ImageID:  req.ImageId,
	})

	// 发起请求
	response, err := v.client.SendCtImage(ctx, realRequest)
	if err != nil {
		return nil, err
	}

	// 解析返回
	err = response.ParseWithCheck(nil)
	if err != nil {
		return nil, err
	}
	return &ImageDeleteResponse{}, nil
}

type imageDeleteRealRequest struct {
	RegionID string `json:"regionID"`
	ImageID  string `json:"imageID"`
}

type ImageDeleteRequest struct {
	RegionId string // 资源池id
	ImageId  string // 镜像id
}

type ImageDeleteResponse struct {
}
