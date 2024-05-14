package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/yangpp6/go-sdk/sdk/common"
	"github.com/yangpp6/go-sdk/sdk/ecs"
	"io"
	"net/http"
)

func createInstance(ak, sk string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "ctecs-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)
	handler := ecs.NewEcsCreateInstanceApi(client)

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
	response, err := handler.Do(context.Background(), credential, &ecs.EcsCreateInstanceRequest{
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

func describeInstances(ak, sk string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "ctecs-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)
	handler := ecs.NewEcsDescribeInstancesApi(client)
	response, err := handler.Do(context.Background(), credential, &ecs.EcsDescribeInstancesRequest{
		RegionId:        "bb9fdb42056f11eda1610242ac110002",
		AzName:          "cn-huadong1-jsnj1A-public-ctcloud",
		ProjectId:       "0",
		PageNo:          1,
		PageSize:        10,
		State:           "active",
		Keyword:         "ecm-57fd",
		InstanceName:    "ecm-57fd",
		InstanceIdList:  "77493826-d038-2a9c-f684-e2f6adabeba3",
		SecurityGroupId: "sg-tdzefke02r",
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func listFlavors(ak, sk string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "ctecs-global.ctapi.ctyun.cn",
	}
	credential, _ := common.NewCredential(ak, sk)
	handler := ecs.NewEcsFlavorListApi(client)

	response, err := handler.Do(context.Background(), credential, &ecs.EcsFlavorListRequest{
		RegionId:     "bb9fdb42056f11eda1610242ac110002",
		AzName:       "cn-huadong1-jsnj1A-public-ctcloud",
		FlavorType:   "CPU_KS1",
		FlavorName:   "ks1.medium.2",
		FlavorCpu:    1,
		FlavorArch:   "arm",
		FlavorSeries: "ks",
		FlavorId:     "b6779240-5649-803b-4a4c-8fc59d310ecf",
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func instanceDetail(ak, sk string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "ctecs-global.ctapi.ctyun.cn",
	}
	credential, _ := common.NewCredential(ak, sk)
	handler := ecs.NewEcsInstanceDetailsApi(client)

	response, err := handler.Do(context.Background(), credential, &ecs.EcsInstanceDetailsRequest{
		RegionId:   "bb9fdb42056f11eda1610242ac110002",
		InstanceId: "77493826-d038-2a9c-f684-e2f6adabeba3",
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func instanceJoinSecurityGroup(ak, sk string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "ctecs-global.ctapi.ctyun.cn",
	}
	credential, _ := common.NewCredential(ak, sk)

	handler := ecs.NewEcsJoinSecurityGroupApi(client)

	response, err := handler.Do(context.Background(), credential, &ecs.EcsJoinSecurityGroupRequest{
		RegionId:           "bb9fdb42056f11eda1610242ac110002",
		SecurityGroupId:    "sg-tdzefke02r",
		InstanceId:         "77493826-d038-2a9c-f684-e2f6adabeba3",
		NetworkInterfaceId: "port-pja7l0zfvk",
		Action:             "joinSecurityGroup",
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func instanceLeaveSecurityGroup(ak, sk string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "ctecs-global.ctapi.ctyun.cn",
	}
	credential, _ := common.NewCredential(ak, sk)
	handler := ecs.NewEcsLeaveSecurityGroupApi(client)

	response, err := handler.Do(context.Background(), credential, &ecs.EcsLeaveSecurityGroupRequest{
		RegionId:        "bb9fdb42056f11eda1610242ac110002",
		SecurityGroupId: "sg-tdzefke02r",
		InstanceId:      "77493826-d038-2a9c-f684-e2f6adabeba3",
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func instanceQueryAsyncResult(ak, sk string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "ctecs-global.ctapi.ctyun.cn",
	}
	credential, _ := common.NewCredential(ak, sk)
	handler := ecs.NewEcsQueryAsyncResultApi(client)
	response, err := handler.Do(context.Background(), credential, &ecs.EcsQueryAsyncResultRequest{
		RegionId: "bb9fdb42056f11eda1610242ac110002",
		JobId:    "",
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func resetInstancePassword(ak, sk string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "ctecs-global.ctapi.ctyun.cn",
	}
	credential, _ := common.NewCredential(ak, sk)
	handler := ecs.NewEcsResetPasswordApi(client)
	response, err := handler.Do(context.Background(), credential, &ecs.EcsResetPasswordRequest{
		RegionId:    "bb9fdb42056f11eda1610242ac110002",
		InstanceId:  "77493826-d038-2a9c-f684-e2f6adabeba3",
		NewPassword: "test-test-test-960",
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func startInstance(ak, sk string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "ctecs-global.ctapi.ctyun.cn",
	}
	credential, _ := common.NewCredential(ak, sk)
	handler := ecs.NewEcsStartInstanceApi(client)
	response, err := handler.Do(context.Background(), credential, &ecs.EcsStartInstanceRequest{
		RegionId:   "bb9fdb42056f11eda1610242ac110002",
		InstanceId: "77493826-d038-2a9c-f684-e2f6adabeba3",
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func stopInstance(ak, sk string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "ctecs-global.ctapi.ctyun.cn",
	}
	credential, _ := common.NewCredential(ak, sk)
	handler := ecs.NewEcsStopInstanceApi(client)
	response, err := handler.Do(context.Background(), credential, &ecs.EcsStopInstanceRequest{
		RegionId:   "bb9fdb42056f11eda1610242ac110002",
		InstanceId: "77493826-d038-2a9c-f684-e2f6adabeba3",
		Force:      false,
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func UnsubscribeInstance(ak, sk string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "ctecs-global.ctapi.ctyun.cn",
	}
	credential, _ := common.NewCredential(ak, sk)
	handler := ecs.NewEcsUnsubscribeInstanceApi(client)
	response, err := handler.Do(context.Background(), credential, &ecs.EcsUnsubscribeInstanceRequest{
		ClientToken: "unsubscribe-instance-test",
		RegionId:    "bb9fdb42056f11eda1610242ac110002",
		InstanceId:  "77493826-d038-2a9c-f684-e2f6adabeba3",
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func UpdateFlavorSpec(ak, sk string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "ctecs-global.ctapi.ctyun.cn",
	}
	credential, _ := common.NewCredential(ak, sk)
	handler := ecs.NewEcsUpdateFlavorSpecApi(client)
	response, err := handler.Do(context.Background(), credential, &ecs.EcsUpdateFlavorSpecRequest{
		RegionId:    "bb9fdb42056f11eda1610242ac110002",
		InstanceId:  "77493826-d038-2a9c-f684-e2f6adabeba3",
		FlavorId:    "b6779240-5649-803b-4a4c-8fc59d310ecf",
		ClientToken: "update-flavor-spec-test",
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func listInstanceVolumes(ak, sk string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "ctecs-global.ctapi.ctyun.cn",
	}
	credential, _ := common.NewCredential(ak, sk)
	handler := ecs.NewEcsVolumeListApi(client)
	response, err := handler.Do(context.Background(), credential, &ecs.EcsVolumeListRequest{
		RegionId:   "bb9fdb42056f11eda1610242ac110002",
		InstanceId: "77493826-d038-2a9c-f684-e2f6adabeba3",
		PageNo:     1,
		PageSize:   10,
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func main() {
	var ak = "184fc5f6be184d3794e00e81fe202e4c"
	var sk = "7d8c1186994147d9bc3352f967f8ca1a"

	flag.StringVar(&ak, "ak", "", "access key")
	flag.StringVar(&sk, "sk", "", "secret key")
	flag.Parse()

	listFlavors(ak, sk)
	createInstance(ak, sk)
	describeInstances(ak, sk)
	instanceDetail(ak, sk)
	instanceJoinSecurityGroup(ak, sk)
	instanceLeaveSecurityGroup(ak, sk)
	instanceQueryAsyncResult(ak, sk)
	resetInstancePassword(ak, sk)
	startInstance(ak, sk)
	stopInstance(ak, sk)
	UnsubscribeInstance(ak, sk)
	UpdateFlavorSpec(ak, sk)
	listInstanceVolumes(ak, sk)
}
