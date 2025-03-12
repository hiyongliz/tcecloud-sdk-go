// All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20170312

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2017-03-12"

var _ = tchttp.POST

type Client struct {
	common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
	cpf := profile.NewClientProfile()
	client = &Client{}
	client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
	return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewInquiryResourceResetInstancesTypeRequest() (request *InquiryResourceResetInstancesTypeRequest) {
	request = &InquiryResourceResetInstancesTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "InquiryResourceResetInstancesType")
	return
}

func NewInquiryResourceResetInstancesTypeResponse() (response *InquiryResourceResetInstancesTypeResponse) {
	response = &InquiryResourceResetInstancesTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InquiryResourceResetInstancesType）用于调整实例的资源查询
func (c *Client) InquiryResourceResetInstancesType(request *InquiryResourceResetInstancesTypeRequest) (response *InquiryResourceResetInstancesTypeResponse, err error) {
	if request == nil {
		request = NewInquiryResourceResetInstancesTypeRequest()
	}
	response = NewInquiryResourceResetInstancesTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceConfigInfosRequest() (request *DescribeInstanceConfigInfosRequest) {
	request = &DescribeInstanceConfigInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceConfigInfos")
	return
}

func NewDescribeInstanceConfigInfosResponse() (response *DescribeInstanceConfigInfosResponse) {
	response = &DescribeInstanceConfigInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeInstanceConfigInfos) 获取实例静态配置信息，包含CPU核数、CPU型号、内存大小和带宽信息等。
func (c *Client) DescribeInstanceConfigInfos(request *DescribeInstanceConfigInfosRequest) (response *DescribeInstanceConfigInfosResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceConfigInfosRequest()
	}
	response = NewDescribeInstanceConfigInfosResponse()
	err = c.Send(request, response)
	return
}

func NewResetInstancesTypeRequest() (request *ResetInstancesTypeRequest) {
	request = &ResetInstancesTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ResetInstancesType")
	return
}

func NewResetInstancesTypeResponse() (response *ResetInstancesTypeResponse) {
	response = &ResetInstancesTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ResetInstancesType) 用于调整实例的机型。
// * 目前只支持[系统盘类型](../数据结构#systemdisk)是`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`类型的实例使用该接口进行机型调整。
// * 目前不支持[]实例使用该接口调整机型。* 目前不支持跨机型系统来调整机型，即使用该接口时指定的`InstanceType`和实例原来的机型需要属于同一系列。* 对于包年包月实例，使用该接口会涉及扣费，请确保账户余额充足。
func (c *Client) ResetInstancesType(request *ResetInstancesTypeRequest) (response *ResetInstancesTypeResponse, err error) {
	if request == nil {
		request = NewResetInstancesTypeRequest()
	}
	response = NewResetInstancesTypeResponse()
	err = c.Send(request, response)
	return
}

func NewResizeInstanceDisksRequest() (request *ResizeInstanceDisksRequest) {
	request = &ResizeInstanceDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "ResizeInstanceDisks")
	return
}

func NewResizeInstanceDisksResponse() (response *ResizeInstanceDisksResponse) {
	response = &ResizeInstanceDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ResizeInstanceDisks) 用于扩容实例的数据盘。
//
// * 目前只支持扩容随实例购买的数据盘，且[数据盘类型](../数据结构#datadisk)为：`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`。* 目前不支持[CDH]实例使用该接口扩容数据盘。
// * 对于包年包月实例，使用该接口会涉及扣费，请确保账户余额充足。
// * 目前只支持扩容一块数据盘。
func (c *Client) ResizeInstanceDisks(request *ResizeInstanceDisksRequest) (response *ResizeInstanceDisksResponse, err error) {
	if request == nil {
		request = NewResizeInstanceDisksRequest()
	}
	response = NewResizeInstanceDisksResponse()
	err = c.Send(request, response)
	return
}

func NewQueryInstancesRequest() (request *QueryInstancesRequest) {
	request = &QueryInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cvm", APIVersion, "QueryInstances")
	return
}

func NewQueryInstancesResponse() (response *QueryInstancesResponse) {
	response = &QueryInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询实例列表
func (c *Client) QueryInstances(request *QueryInstancesRequest) (response *QueryInstancesResponse, err error) {
	if request == nil {
		request = NewQueryInstancesRequest()
	}
	response = NewQueryInstancesResponse()
	err = c.Send(request, response)
	return
}
