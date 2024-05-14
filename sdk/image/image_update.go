package image

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// imageUpdateApi 删除私有镜像
type imageUpdateApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewImageUpdateApi(client *common.CtyunSender) common.ApiHandler[ImageUpdateRequest, ImageUpdateResponse] {
	return &imageUpdateApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/image/update",
		},
	}
}

func (v *imageUpdateApi) Do(ctx context.Context, credential *common.Credential, req *ImageUpdateRequest) (*ImageUpdateResponse, common.CtyunRequestError) {
	// 构建请求
	realRequest := v.WithCredential(credential)
	_, err := realRequest.WriteJson(&imageUpdateRealRequest{
		RegionID:    req.RegionId,
		ImageID:     req.ImageId,
		BootMode:    req.BootMode,
		Description: req.Description,
		ImageName:   req.ImageName,
		MaximumRAM:  req.MaximumRam,
		MinimumRAM:  req.MinimumRam,
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
	return &ImageUpdateResponse{}, nil
}

type imageUpdateRealRequest struct {
	ImageID     string `json:"imageID"`
	RegionID    string `json:"regionID"`
	BootMode    string `json:"bootMode,omitempty"`
	Description string `json:"description,omitempty"`
	ImageName   string `json:"imageName,omitempty"`
	MaximumRAM  int    `json:"maximumRAM,omitempty"`
	MinimumRAM  int    `json:"minimumRAM,omitempty"`
}

type ImageUpdateRequest struct {
	ImageId     string
	RegionId    string
	BootMode    string
	Description string
	ImageName   string
	MaximumRam  int
	MinimumRam  int
}

type ImageUpdateResponse struct {
}
