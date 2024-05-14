package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/yangpp6/go-sdk/sdk/common"
	"github.com/yangpp6/go-sdk/sdk/image"
	"net/http"
)

func imageDetail(ak, sk, regionID, imageID string) {
	client := &common.CtyunSender{
		Client:          &http.Client{},
		CtImageEndPoint: "ctimage-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)

	handler := image.NewImageDetailApi(client)
	ctx := context.TODO()
	createRes, err := handler.Do(ctx, credential, &image.ImageDetailRequest{
		RegionId: regionID,
		ImageId:  imageID,
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(createRes)
	fmt.Printf(string(marshal))
}

func imageDelete(ak, sk, regionID, imageID string) {
	client := &common.CtyunSender{
		Client:          &http.Client{},
		CtImageEndPoint: "ctimage-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)

	handler := image.NewImageDeleteApi(client)
	ctx := context.TODO()
	res, err := handler.Do(ctx, credential, &image.ImageDeleteRequest{
		RegionId: regionID,
		ImageId:  imageID,
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func imageList(ak, sk, regionID, azName string) {
	client := &common.CtyunSender{
		Client:          &http.Client{},
		CtImageEndPoint: "ctimage-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)

	handler := image.NewImageListApi(client)
	ctx := context.TODO()
	res, err := handler.Do(ctx, credential, &image.ImageListRequest{
		RegionId:   regionID,
		AzName:     azName,
		Visibility: 1,
		PageNo:     1,
		PageSize:   10,
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func imageUpdate(ak string, sk string, regionID string, imageID string, imageName string, maximumRam int, minimumRam int) {
	client := &common.CtyunSender{
		Client:          &http.Client{},
		CtImageEndPoint: "ctimage-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)
	handler := image.NewImageUpdateApi(client)
	ctx := context.TODO()
	res, err := handler.Do(ctx, credential, &image.ImageUpdateRequest{
		RegionId:   regionID,
		ImageId:    imageID,
		ImageName:  imageName,
		MaximumRam: maximumRam,
		MinimumRam: minimumRam,
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
	var imageID string
	var imageName string
	var azName string
	maximumRam := 16
	minimumRam := 4
	flag.StringVar(&ak, "ak", "", "access key")
	flag.StringVar(&sk, "sk", "", "secret key")
	flag.StringVar(&regionID, "region-id", "", "region id")
	flag.StringVar(&imageID, "image-id", "", "image id")
	flag.StringVar(&imageName, "image-name", "", "image name")
	flag.StringVar(&azName, "azname", "", "az name")
	flag.Parse()
	imageList(ak, sk, regionID, azName)
	imageDetail(ak, sk, regionID, imageID)
	imageUpdate(ak, sk, regionID, imageID, imageName, maximumRam, minimumRam)
	imageDelete(ak, sk, regionID, imageID)
}
