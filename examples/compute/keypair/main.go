package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"openapi-sdk-go/sdk/common"
	"openapi-sdk-go/sdk/keypair"
)

func keypairDetail(ak, sk, regionID, keyPairName string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "https://ctecs-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)
	handler := keypair.NewKeypairCreateApi(client)
	ctx := context.TODO()
	res, err := handler.Do(ctx, credential, &keypair.KeypairCreateRequest{
		RegionId:    regionID,
		KeyPairName: keyPairName,
		ProjectId:   "0",
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func keypairDelete(ak, sk, regionID, keyPairName string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "https://ctecs-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)

	handler := keypair.NewKeypairDeleteApi(client)
	ctx := context.TODO()
	res, err := handler.Do(ctx, credential, &keypair.KeypairDeleteRequest{
		RegionId:    regionID,
		KeyPairName: keyPairName,
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func keypairImport(ak, sk, regionID, keyPairName, publicKey string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "https://ctecs-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)
	handler := keypair.NewKeypairImportApi(client)
	ctx := context.TODO()
	res, err := handler.Do(ctx, credential, &keypair.KeypairImportRequest{
		RegionId:    regionID,
		KeyPairName: keyPairName,
		PublicKey:   publicKey,
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func keypairAttach(ak, sk, regionID, keyPairName, instanceID string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "https://ctecs-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)

	handler := keypair.NewKeypairAttachApi(client)
	ctx := context.TODO()
	res, err := handler.Do(ctx, credential, &keypair.KeypairAttachRequest{
		RegionId:    regionID,
		KeyPairName: keyPairName,
		InstanceId:  instanceID,
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func keypairDetach(ak, sk, regionID, keyPairName, instanceID string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "https://ctecs-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)
	handler := keypair.NewKeypairDetachApi(client)
	ctx := context.TODO()
	res, err := handler.Do(ctx, credential, &keypair.KeypairDetachRequest{
		RegionId:    regionID,
		KeyPairName: keyPairName,
		InstanceId:  instanceID,
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func main() {
	var ak string
	var sk string
	var regionID string
	var keyPairName string
	var publicKey string
	var instanceID string
	flag.StringVar(&ak, "ak", "", "access key")
	flag.StringVar(&sk, "sk", "", "secret key")
	flag.StringVar(&regionID, "region-id", "", "region id")
	flag.StringVar(&keyPairName, "keyPairName", "", "keypair name")
	flag.StringVar(&publicKey, "publicKey", "", "public key")
	flag.StringVar(&instanceID, "instanceID", "", "instance id")
	flag.Parse()
	keypairImport(ak, sk, regionID, keyPairName, publicKey)
	keypairDetail(ak, sk, regionID, keyPairName)
	keypairAttach(ak, sk, regionID, keyPairName, instanceID)
	keypairDetach(ak, sk, regionID, keyPairName, instanceID)
	keypairDelete(ak, sk, regionID, keyPairName)
}
