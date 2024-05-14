# 公有云 Iaas SDK - GO
- 文档中 SDK 关于 API 的示例代码仅供参考，各 API 的完整使用步骤与说明请参见[OpenApi文档](https://www.ctyun.cn/document/10026730/10106056 "链接标题")。

## 环境要求

  Go 环境版本必须不低于 1.18

## 发布地址

https://github.com/yangpp6/go-sdk

## 源码仓库地址

https://github.com/yangpp6/go-sdk
## 安装方式

```
go get github.com/yangpp6/go-sdk
```

## 示例说明

以下代码介绍了SDK 的使用步骤，仅作步骤示范。示例展示了如何调用 NewEcsLeaveSecurityGroupApi解绑安全组请求。
说明：您需要设置账号的ak，sk以及对应资源的信息。

## 示例代码

```
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"github.com/yangpp6/go-sdk/sdk/common"
	"github.com/yangpp6/go-sdk/sdk/ecs")

func instanceLeaveSecurityGroup(ak, sk string) {
  NewEcsLeaveSecurityGroupApi	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "ctecs-global.ctapi.ctyun.cn",
	}
	credential, _ := common.NewCredential(ak, sk)
	handler := ecs.NewEcsLeaveSecurityGroupApi(client)

	response, err := handler.Do(context.Background(), credential, &ecs.EcsLeaveSecurityGroupRequest{
		RegionId:        "region_id",
		SecurityGroupId: "security_group_id",
		InstanceId:      "instance_id",
	})
	if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return解绑安全组	jsonstr, _ := json.Marshal(response)

  NewEcsLeaveSecurityGroupApi	fmt.Println(string(jsonstr))
}

func main() {
	var ak string
	var sk string

	flag.StringVar(&ak, "ak", "", "access key")
	flag.StringVar(&sk, "sk", "", "secret key")
	flag.Parse()

	instanceLeaveSecurityGroup(ak, sk)
}
```

## 步骤介绍

1.解绑安全组 您需要在代码中引入依赖包：
```
import (
  "github.com/yangpp6/go-sdk/sdk/common"
  "github.com/yangpp6/go-sdk/sdk/ecs"
)

```

2.初始化client对象，client对象存放endpoint信息，示例中为ctecs-global.ctapi.ctyun.cn

```
  client := &common.CtyunSender{
		Client:        &http.Client{},
		CtEcsEndPoint: "ctecs-global.ctapi.ctyun.cn",
	}
  handler := ecs.NewEcsLeaveSecurityGroupApi(client)
```

3.初始化credential， 需要使用用户的ak，sk

```
  credential, _ := common.NewCredential(ak, sk)
```

4.创建对应 API 的 Request,初始化请求request必要的参数，并请求调用。

```
	response, err := handler.Do(context.Background(), credential, &ecs.EcsLeaveSecurityGroupRequest{
		RegionId:        "region_id",
		SecurityGroupId: "security_group_id",
		InstanceId:      "instance_id",
	})
```

5.通过err获取错误信息

```
  if err != nil {
		errorContent, _ := io.ReadAll(err.Response().Body)
		fmt.Printf("错误信息为：%s", string(errorContent))
		return
	}
```

6.通过response获取请求返回

```
  jsonstr, _ := json.Marshal(response)
  fmt.Println(string(jsonstr))
```
