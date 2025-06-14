# 简介

欢迎使用开发者工具套件（SDK）3.0，SDK3.0 是云 API3.0 平台的配套工具。后续所有的云服务产品都会接入进来。新版 SDK实 现了统一化，具有各个语言版本的 SDK 使用方法相同，接口调用方式相同，统一的错误码和返回包格式这些优点。

为方便 GO 开发者调试和接入产品 API，这里向您介绍适用于 GO 的开发工具包，并提供首次使用开发工具包的简单示例。让您快速获取GO SDK 并开始调用。

# 依赖环境

1. Go 1.9 版本及以上，并设置好 GOPATH 等必须的环境变量。
2. 使用相关产品前需要在控制台已开通相应产品。
3. 在控制台[访问管理](http://cmgt.cloud.sunhongs.com/cam/capi)页面获取 SecretID 和 SecretKey 。

# 使用

使用 Go SDK 前，先获取安全凭证。在第一次使用云 API 之前，用户首先需要在控制台上申请安全凭证，安全凭证包括 SecretID 和 SecretKey, SecretID 是用于标识 API 调用者的身份，SecretKey 是用于加密签名字符串和服务器端验证签名字符串的密钥。SecretKey 必须严格保管，避免泄露。


# 示例

每个接口都有一个对应的 Request 结构和一个 Response 结构。例如云服务器的查询实例列表接口 DescribeInstances 有对应的请求结构体 DescribeInstancesRequest 和返回结构体 DescribeInstancesResponse 。

下面以云服务器查询实例列表接口为例，介绍 SDK 的基础用法。出于演示的目的，有一些非必要的内容也加上去了，以尽量展示 SDK 常用的功能，但也显得臃肿。在实际编写代码使用 SDK 的时候，应尽量简化。
```go
package main

import (
        "fmt"
        "os"

        "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
        "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/errors"
        "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
        "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/regions"
        cvm "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/cvm/v20170312"

)

func main() {
        // 必要步骤：
        // 实例化一个认证对象，入参需要传入账户密钥对secretId，secretKey。
        // 这里采用的是从环境变量读取的方式，需要在环境变量中先设置这两个值。
        // 你也可以直接在代码中写死密钥对，但是小心不要将代码复制、上传或者分享给他人，
        // 以免泄露密钥对危及你的财产安全。
        credential := common.NewCredential(
                os.Getenv("SECRET_ID"),
                os.Getenv("SECRET_KEY"),
        )

        // 非必要步骤
        // 实例化一个客户端配置对象，可以指定超时时间等配置
        cpf := profile.NewClientProfile()
        // SDK默认使用POST方法。
        // 如果你一定要使用GET方法，可以在这里设置。GET方法无法处理一些较大的请求。
        cpf.HttpProfile.ReqMethod = "GET"
        // SDK有默认的超时时间，非必要请不要进行调整。
        // 如有需要请在代码中查阅以获取最新的默认值。
        cpf.HttpProfile.ReqTimeout = 10
        // SDK会自动指定域名。通常是不需要特地指定域名的。
        // cvm 是接入的产品名，yunapi3是调用版本，cloud.sunhongs.com是主域名。
        cpf.HttpProfile.Endpoint = "cvm.yunapi3.cloud.sunhongs.com"
        // SDK默认用HmacSHA256进行签名，它更安全但是会轻微降低性能。
        // 非必要请不要修改这个字段。
        cpf.SignMethod = "HmacSHA1"

        // 实例化要请求产品(以cvm为例)的client对象
        // 第二个参数是地域信息，可以直接填写字符串，如shanghai，或者引用预设的常量
        // 您可以通过 API location 接口 DescribeRegionZone 查询地域列表 查看完整的地域列表，并选择其中的一个地域发起请求。
        client, _ := cvm.NewClient(credential, "shanghai", cpf)
        // 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
        // 你可以直接查询SDK源码确定DescribeInstancesRequest有哪些属性可以设置，
        // 属性可能是基本类型，也可能引用了另一个数据结构。
        // 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明。
        request := cvm.NewDescribeInstancesRequest()

        // 基本类型的设置。
        // 此接口允许设置返回的实例数量。此处指定为只返回一个。
        // SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
        // SDK提供对基本类型的指针引用封装函数
        request.Limit = common.Int64Ptr(1)

        // 数组类型的设置。
        // 此接口允许指定实例 ID 进行过滤，但是由于和接下来要演示的 Filter 参数冲突，先注释掉。
        // request.InstanceIds = common.StringPtrs([]string{"ins-r8hr2upy"})

        // 复杂对象的设置。
        // 在这个接口中，Filters是数组，数组的元素是复杂对象Filter，Filter的成员Values是string数组。
        request.Filters = []*cvm.Filter{
            &cvm.Filter{
                Name: common.StringPtr("zone"),
                Values: common.StringPtrs([]string{"shanghai-1"}),
            },
        }

        // 使用json字符串设置一个request，注意这里实际是更新request，即Limit=1将会被保留，
        // 而过滤条件的zone将会变为shanghai-2。
        // 如果需要一个全新的request，则需要用cvm.NewDescribeInstancesRequest()创建。
        err := request.FromJsonString(`{"Filters":[{"Name":"zone","Values":["shanghai-2"]}]}`)
        if err != nil {
                panic(err)
        }
        // 通过client对象调用想要访问的接口，需要传入请求对象
        response, err := client.DescribeInstances(request)
        // 处理异常
        if _, ok := err.(*errors.CloudSDKError); ok {
                fmt.Printf("An API error has returned: %s", err)
                return
        }
        // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
        if err != nil {
                panic(err)
        }
        // 打印返回的json字符串
        fmt.Printf("%s", response.ToJsonString())
}
```

更多示例参见examples目录。对于复杂接口的 Request 初始化例子，可以参考 examples/cvm/v20170312/run_instances.go 。对于使用json字符串初始化 Request 的例子，可以参考 examples/cvm/v20170312/describe_instances.go 。

