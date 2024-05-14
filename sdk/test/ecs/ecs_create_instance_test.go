package test

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"openapi-sdk-go/sdk/common"
	"openapi-sdk-go/sdk/ecs"
	"openapi-sdk-go/sdk/test"
	"testing"
)

func TestNewEcsCreateInstanceApi_Do(t *testing.T) {
	var ak = "<YOUR AK>"
	var sk = "<YOUR SK>"
	client := test.BuildProdClient()
	credential, _ := common.NewCredential(ak, sk)
	networkCardList := make([]ecs.EcsCreateInstanceNetworkCardListRequest, 0)
	network_card := ecs.EcsCreateInstanceNetworkCardListRequest{
		NicName:  "nic-test",
		IsMaster: true,
		SubnetId: "subnet-4c4333pc67",
	}
	networkCardList = append(networkCardList, network_card)
	dataDiskList := make([]ecs.EcsCreateInstanceDataDiskListRequest, 0)
	data_disk := ecs.EcsCreateInstanceDataDiskListRequest{
		DiskMode: "VBD",
		DiskName: "data-disk-test",
		DiskType: "SATA",
		DiskSize: 20,
	}
	dataDiskList = append(dataDiskList, data_disk)
	labelList := make([]ecs.EcsCreateInstanceLabelListRequest, 0)
	label := ecs.EcsCreateInstanceLabelListRequest{
		LabelKey:   "label-key-test",
		LabelValue: "label-value-test",
	}
	labelList = append(labelList, label)
	response, err := client.Apis.EcsCreateInstanceApi.Do(context.Background(), credential, &ecs.EcsCreateInstanceRequest{
		ClientToken:     "ecs-create-instance-test-02",
		RegionId:        "bb9fdb42056f11eda1610242ac110002",
		AzName:          "cn-huadong1-jsnj1A-public-ctcloud",
		InstanceName:    "ecm-go-test",
		DisplayName:     "ecm-go-test",
		FlavorId:        "b6779240-5649-803b-4a4c-8fc59d310ecf",
		ImageType:       1,
		ImageId:         "939c131f-a986-420f-a3b2-57feb9995e47",
		BootDiskType:    "SATA",
		BootDiskSize:    40,
		VpcId:           "vpc-chz0ilszsp",
		OnDemand:        false,
		NetworkCardList: networkCardList,
		ExtIp:           "1",
		ProjectID:       "0",
		SecGroupList:    []string{"sg-bqv0t629h6", "sg-bqv0t629h6"},
		DataDiskList:    dataDiskList,
		IpVersion:       "ipv4",
		Bandwidth:       50,
		UserPassword:    "qyo84!*ymd",
		CycleCount:      1,
		CycleType:       "MONTH",
		AutoRenewStatus: 0,
		UserData:        "YmF0Y2hDcmVhdGVUZXN0MDgwMw==",
		PayVoucherPrice: 1819.99,
		LabelList:       labelList,
		MonitorService:  true,
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))

}
