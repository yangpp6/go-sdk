package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/yangpp6/go-sdk/sdk/common"
	"github.com/yangpp6/go-sdk/sdk/vpc"
	"testing"
)

func TestVpcDeleteApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("cb65661223244d2fb9f99baac12478a3", "181fcdb898ae4ac9ac7af3563874d091")
	// do, _ := of.Apis.VpcListApi.Do(context.Background(), credential, &vpc.VpcListRequest{
	// 	RegionId:   "bb9fdb42056f11eda1610242ac110002",
	// 	VpcIds:     []string{"vpc-qjfduzwab9"},
	// 	PageNumber: 1,
	// 	PageSize:   1,
	// })
	do, _ := BuildDevClient().Apis.VpcCreateApi.Do(context.Background(), credential, &vpc.VpcCreateRequest{
		RegionId:    TestNeiMeng8RegionId,
		ClientToken: uuid.NewString(),
		Name:        "sdsdfdfss222",
		Cidr:        "10.0.0.0/8",
		Description: "asdfasdf",
		EnableIpv6:  false,
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}
