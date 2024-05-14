package image

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

// imageDetailApi 查询镜像详细信息
type imageDetailApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewImageDetailApi(client *common.CtyunSender) common.ApiHandler[ImageDetailRequest, ImageDetailResponse] {
	return &imageDetailApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/image/detail",
		},
	}
}

func (v *imageDetailApi) Do(ctx context.Context, credential *common.Credential, req *ImageDetailRequest) (*ImageDetailResponse, common.CtyunRequestError) {
	// 构建请求
	realRequest := v.
		WithCredential(credential).
		AddParam("regionID", req.RegionId).
		AddParam("imageID", req.ImageId)

	// 发起请求
	response, err := v.client.SendCtImage(ctx, realRequest)
	if err != nil {
		return nil, err
	}

	// 解析返回
	realResponse := &imageDetailRealResponse{}
	err = response.ParseWithCheck(realResponse)
	if err != nil {
		return nil, err
	}
	var images []ImageDetailImagesResponse
	for _, img := range realResponse.Images {
		images = append(images, ImageDetailImagesResponse{
			Architecture:     img.Architecture,
			AzName:           img.AzName,
			BootMode:         img.BootMode,
			ContainerFormat:  img.ContainerFormat,
			CreatedTime:      img.CreatedTime,
			Description:      img.Description,
			DestinationUser:  img.DestinationUser,
			DiskFormat:       img.DiskFormat,
			DiskId:           img.DiskID,
			DiskSize:         img.DiskSize,
			ImageClass:       img.ImageClass,
			ImageId:          img.ImageID,
			ImageName:        img.ImageName,
			ImageType:        img.ImageType,
			MaximumRam:       img.MaximumRAM,
			MinimumRam:       img.MinimumRAM,
			OsDistro:         img.OsDistro,
			OsType:           img.OsType,
			OsVersion:        img.OsVersion,
			ProjectId:        img.ProjectID,
			SharedListLength: img.SharedListLength,
			Size:             img.Size,
			SourceServerId:   img.SourceServerID,
			SourceUser:       img.SourceUser,
			Status:           img.Status,
			Tags:             img.Tags,
			UpdatedTime:      img.UpdatedTime,
			Visibility:       img.Visibility,
		})
	}
	return &ImageDetailResponse{
		Images: images,
	}, nil
}

type imageDetailRealResponse struct {
	Images []imageDetailImagesRealResponse
}

type imageDetailImagesRealResponse struct {
	Architecture     string `json:"architecture"`
	AzName           string `json:"azName"`
	BootMode         string `json:"bootMode"`
	ContainerFormat  string `json:"containerFormat"`
	CreatedTime      int    `json:"createdTime"`
	Description      string `json:"description"`
	DestinationUser  string `json:"destinationUser"`
	DiskFormat       string `json:"diskFormat"`
	DiskID           string `json:"diskID"`
	DiskSize         int    `json:"diskSize"`
	ImageClass       string `json:"imageClass"`
	ImageID          string `json:"imageID"`
	ImageName        string `json:"imageName"`
	ImageType        string `json:"imageType"`
	MaximumRAM       int    `json:"maximumRAM"`
	MinimumRAM       int    `json:"minimumRAM"`
	OsDistro         string `json:"osDistro"`
	OsType           string `json:"osType"`
	OsVersion        string `json:"osVersion"`
	ProjectID        string `json:"projectID"`
	SharedListLength int    `json:"sharedListLength"`
	Size             int64  `json:"size"`
	SourceServerID   string `json:"sourceServerID"`
	SourceUser       string `json:"sourceUser"`
	Status           string `json:"status"`
	Tags             string `json:"tags"`
	UpdatedTime      int    `json:"updatedTime"`
	Visibility       string `json:"visibility"`
}

type ImageDetailRequest struct {
	RegionId string // 资源池id
	ImageId  string // 镜像id
}

type ImageDetailImagesResponse struct {
	Architecture     string
	AzName           string
	BootMode         string
	ContainerFormat  string
	CreatedTime      int
	Description      string
	DestinationUser  string
	DiskFormat       string
	DiskId           string
	DiskSize         int
	ImageClass       string
	ImageId          string
	ImageName        string
	ImageType        string
	MaximumRam       int
	MinimumRam       int
	OsDistro         string
	OsType           string
	OsVersion        string
	ProjectId        string
	SharedListLength int
	Size             int64
	SourceServerId   string
	SourceUser       string
	Status           string
	Tags             string
	UpdatedTime      int
	Visibility       string
}

type ImageDetailResponse struct {
	Images []ImageDetailImagesResponse
}
