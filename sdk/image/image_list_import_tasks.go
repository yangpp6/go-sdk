package image

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
	"strconv"
)

// imageListImportTasksApi 查询镜像
type imageListImportTasksApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewImageListImportTasksApi(client *common.CtyunSender) common.ApiHandler[ImageListImportTasksRequest, ImageListImportTasksResponse] {
	return &imageListImportTasksApi{
		client: client,
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/image/list-import-tasks",
		},
	}
}

func (v *imageListImportTasksApi) Do(ctx context.Context, credential *common.Credential, req *ImageListImportTasksRequest) (*ImageListImportTasksResponse, common.CtyunRequestError) {
	// 构建请求
	request := v.WithCredential(credential)
	request.AddParam("regionID", req.RegionId)
	request.AddParam("pageNo", strconv.Itoa(req.PageNo))
	request.AddParam("pageSize", strconv.Itoa(req.PageSize))

	// 发起请求
	response, err := v.client.SendCtImage(ctx, request)
	if err != nil {
		return nil, err
	}

	// 解析返回
	var resp imageListImportTasksRealResponse
	err = response.ParseWithCheck(&resp)
	if err != nil {
		return nil, err
	}
	return &ImageListImportTasksResponse{
		PageNo:       resp.PageNo,
		CurrentPage:  resp.CurrentPage,
		PageSize:     resp.PageSize,
		CurrentCount: resp.CurrentCount,
		TotalCount:   resp.TotalCount,
		ImageImportTasks: ImageListImportTasksImageImportTasksResponse{
			ImageName:  resp.ImageImportTasks.ImageName,
			OsType:     resp.ImageImportTasks.OsType,
			TaskId:     resp.ImageImportTasks.TaskID,
			TaskStatus: resp.ImageImportTasks.TaskStatus,
		},
	}, nil
}

type imageListImportTasksImageImportTasksRealResponse struct {
	ImageName  string `json:"imageName"`
	OsType     string `json:"osType"`
	TaskID     string `json:"taskID"`
	TaskStatus string `json:"taskStatus"`
}

type imageListImportTasksRealResponse struct {
	PageNo           int                                              `json:"pageNo"`
	CurrentPage      int                                              `json:"currentPage"`
	PageSize         int                                              `json:"pageSize"`
	CurrentCount     int                                              `json:"currentCount"`
	TotalCount       int                                              `json:"totalCount"`
	ImageImportTasks imageListImportTasksImageImportTasksRealResponse `json:"imageImportTasks"`
}

type ImageListImportTasksRequest struct {
	RegionId string // 资源池id
	PageNo   int    // 页码，取值范围：最小 1（默认值）
	PageSize int    // 每页记录数目，取值范围：最小 1，最大 50，默认值 10
}

type ImageListImportTasksImageImportTasksResponse struct {
	ImageName  string
	OsType     string
	TaskId     string
	TaskStatus string
}

type ImageListImportTasksResponse struct {
	PageNo           int
	CurrentPage      int
	PageSize         int
	CurrentCount     int
	TotalCount       int
	ImageImportTasks ImageListImportTasksImageImportTasksResponse
}
