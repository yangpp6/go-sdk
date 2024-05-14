package idp

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
	"strconv"
)

type idpListApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewIdpListApi(client *common.CtyunSender) common.ApiHandler[IdpListRequest, IdpListResponse] {
	return &idpListApi{
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v1/identityProvider/getIdPDetails",
		},
		client: client,
	}
}

type IdpListRequest struct {
	Id int64
}

func (i *idpListApi) Do(ctx context.Context, credential *common.Credential, r *IdpListRequest) (*IdpListResponse, common.CtyunRequestError) {
	request := i.WithCredential(credential).AddParam("id", strconv.FormatInt(r.Id, 10))
	ctiam, err := i.client.SendCtiam(ctx, request)
	if err != nil {
		return nil, err
	}
	resp := &IdpListResponse{}
	err = ctiam.ParseWithCheck(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type IdpListResponse struct {
	Id               int64  `json:"id"`
	Protocol         int64  `json:"protocol"`
	AccountId        string `json:"accountId"`
	Remark           string `json:"remark"`
	Name             string `json:"name"`
	Type             int64  `json:"type"`
	Status           int64  `json:"status"`
	CreateTime       int64  `json:"createTime"`
	UpdateTime       int64  `json:"updateTime"`
	MetadataDocument string `json:"metadataDocument"`
	FileName         string `json:"fileName"`
}
