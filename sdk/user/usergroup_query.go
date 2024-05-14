package user

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk/common"
	"net/http"
)

type userGroupQueryApi struct {
	builder common.CtyunRequestBuilder
	client  *common.CtyunSender
}

func NewUserGroupQueryApi(client *common.CtyunSender) common.ApiHandler[UserGroupQueryRequest, UserGroupQueryResponse] {
	return &userGroupQueryApi{
		builder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v1/userGroup/getGroups",
		},
		client: client,
	}
}

type userGroupQueryRealRequest struct {
	GroupName string `json:"groupName"`
	PageNum   int64  `json:"pageNum"`
	PageSize  int64  `json:"pageSize"`
}

type userGroupQueryRealResponse struct {
	PageNum       int64           `json:"pageNum"`
	PageSize      int64           `json:"pageSize"`
	StartRow      int64           `json:"startRow"`
	EndRow        int64           `json:"endRow"`
	Total         int64           `json:"total"`
	Pages         int64           `json:"pages"`
	NavigatePages int64           `json:"navigatePages"`
	Result        []groupRealInfo `json:"result"`
}

type groupRealInfo struct {
	Id         string `json:"id"`
	GroupName  string `json:"groupName"`
	AccountId  string `json:"accountId"`
	GroupIntro string `json:"groupIntro"`
	IsRoot     int64  `json:"isRoot"`
	IsValid    int64  `json:"isValid"`
	UserCount  int64  `json:"userCount"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

type UserGroupQueryRequest struct {
	GroupName string
	PageNum   int64
	PageSize  int64
}

type UserGroupQueryResponse struct {
	PageNum       int64
	PageSize      int64
	StartRow      int64
	EndRow        int64
	Total         int64
	Pages         int64
	NavigatePages int64
	Result        []GroupInfo
}

type GroupInfo struct {
	Id         string
	GroupName  string
	AccountId  string
	GroupIntro string
	IsRoot     int64
	IsValid    int64
	UserCount  int64
	CreateTime int64
	UpdateTime int64
}

func (receiver *userGroupQueryApi) Do(ctx context.Context, credential *common.Credential, req *UserGroupQueryRequest) (*UserGroupQueryResponse, common.CtyunRequestError) {
	// 构建请求
	builder := receiver.builder.WithCredential(credential)
	builder, err := builder.WriteJson(&userGroupQueryRealRequest{
		GroupName: req.GroupName,
		PageNum:   req.PageNum,
		PageSize:  req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	send, err := receiver.client.SendCtiam(ctx, builder)
	if err != nil {
		return nil, err
	}
	resp := &userGroupQueryRealResponse{}
	err = send.ParseWithCheck(resp)
	if err != nil {
		return nil, err
	}
	var infos []GroupInfo
	for _, info := range resp.Result {
		infos = append(infos, GroupInfo{
			Id:         info.Id,
			GroupName:  info.GroupName,
			AccountId:  info.AccountId,
			GroupIntro: info.GroupIntro,
			IsRoot:     info.IsRoot,
			IsValid:    info.IsValid,
			UserCount:  info.UserCount,
			CreateTime: info.CreateTime,
			UpdateTime: info.UpdateTime,
		})
	}

	return &UserGroupQueryResponse{
		PageNum:       resp.PageNum,
		PageSize:      resp.PageSize,
		StartRow:      resp.StartRow,
		EndRow:        resp.EndRow,
		Total:         resp.Total,
		Pages:         resp.Pages,
		NavigatePages: resp.NavigatePages,
		Result:        infos,
	}, nil
}
