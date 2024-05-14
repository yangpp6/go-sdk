package user

import (
	"context"
	"errors"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

type userCreateApi struct {
	builder common.CtyunRequestBuilder
	client  *common.CtyunSender
}

func NewUserCreateApi(client *common.CtyunSender) common.ApiHandler[UserCreateRequest, UserCreateResponse] {
	return &userCreateApi{
		builder: common.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v1/user/createUser",
		},
		client: client,
	}
}

type UserCreateRequest struct {
	LoginEmail         string      `json:"loginEmail"`
	MobilePhone        string      `json:"mobilePhone"`
	Password           string      `json:"password"`
	UserName           string      `json:"userName"`
	Remark             string      `json:"remark"`
	Groups             []UserGroup `json:"groups"`
	GeneratePassword   bool        `json:"generatePassword"`
	LoginResetPassword bool        `json:"loginResetPassword"`
	SourcePassword     string      `json:"sourcePassword"`
}

type UserGroup struct {
	Id string `json:"id"`
}

type UserCreateResponse struct {
	AccountId        string      `json:"accountId"`
	Groups           []UserGroup `json:"groups"`
	IsVirtualAccount string      `json:"isVirtualAccount"`
	LoginEmail       string      `json:"loginEmail"`
	MobilePhone      string      `json:"mobilePhone"`
	Remark           string      `json:"remark"`
	UserId           string      `json:"userId"`
	UserName         string      `json:"userName"`
}

func (receiver *userCreateApi) Do(ctx context.Context, credential *common.Credential, req *UserCreateRequest) (*UserCreateResponse, common.CtyunRequestError) {
	e := receiver.valid(req)
	if e != nil {
		return nil, common.ErrorBeforeRequest(e)
	}
	// 构建请求
	builder := receiver.builder.WithCredential(credential)
	builder, err := builder.WriteJson(req)
	if err != nil {
		return nil, err
	}
	send, err := receiver.client.SendCtiam(ctx, builder)
	if err != nil {
		return nil, err
	}
	resp := &UserCreateResponse{}
	err = send.ParseWithCheck(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 校验
func (receiver userCreateApi) valid(req *UserCreateRequest) error {
	if req.LoginEmail == "" {
		return errors.New("loginEmail必填")
	}
	if req.MobilePhone == "" {
		return errors.New("mobilePhone必填")
	}
	if req.UserName == "" {
		return errors.New("userName必填")
	}
	if !req.GeneratePassword {
		if req.Password == "" && req.SourcePassword == "" {
			return errors.New("generatePassword字段为false时，password和sourcePassword，必填其中之一")
		}
		if req.Password != "" && req.SourcePassword != "" {
			return errors.New("generatePassword字段为false时，password和sourcePassword，二选一填写")
		}
	}
	return nil
}
