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

package v20190325

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-03-25"

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

func NewDescribeTagResourceListRequest() (request *DescribeTagResourceListRequest) {
	request = &DescribeTagResourceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tag", APIVersion, "DescribeTagResourceList")
	return
}

func NewDescribeTagResourceListResponse() (response *DescribeTagResourceListResponse) {
	response = &DescribeTagResourceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取已接入标签业务列表详情
func (c *Client) DescribeTagResourceList(request *DescribeTagResourceListRequest) (response *DescribeTagResourceListResponse, err error) {
	if request == nil {
		request = NewDescribeTagResourceListRequest()
	}
	response = NewDescribeTagResourceListResponse()
	err = c.Send(request, response)
	return
}

func NewCheckTagResourceInfoRequest() (request *CheckTagResourceInfoRequest) {
	request = &CheckTagResourceInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tag", APIVersion, "CheckTagResourceInfo")
	return
}

func NewCheckTagResourceInfoResponse() (response *CheckTagResourceInfoResponse) {
	response = &CheckTagResourceInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验标签配置
func (c *Client) CheckTagResourceInfo(request *CheckTagResourceInfoRequest) (response *CheckTagResourceInfoResponse, err error) {
	if request == nil {
		request = NewCheckTagResourceInfoRequest()
	}
	response = NewCheckTagResourceInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTagServiceListRequest() (request *DescribeTagServiceListRequest) {
	request = &DescribeTagServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tag", APIVersion, "DescribeTagServiceList")
	return
}

func NewDescribeTagServiceListResponse() (response *DescribeTagServiceListResponse) {
	response = &DescribeTagServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取已接入标签业务模块
func (c *Client) DescribeTagServiceList(request *DescribeTagServiceListRequest) (response *DescribeTagServiceListResponse, err error) {
	if request == nil {
		request = NewDescribeTagServiceListRequest()
	}
	response = NewDescribeTagServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewSaveTagServiceConfigRequest() (request *SaveTagServiceConfigRequest) {
	request = &SaveTagServiceConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tag", APIVersion, "SaveTagServiceConfig")
	return
}

func NewSaveTagServiceConfigResponse() (response *SaveTagServiceConfigResponse) {
	response = &SaveTagServiceConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新产品接入标签
func (c *Client) SaveTagServiceConfig(request *SaveTagServiceConfigRequest) (response *SaveTagServiceConfigResponse, err error) {
	if request == nil {
		request = NewSaveTagServiceConfigRequest()
	}
	response = NewSaveTagServiceConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetServiceApiListRequest() (request *GetServiceApiListRequest) {
	request = &GetServiceApiListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tag", APIVersion, "GetServiceApiList")
	return
}

func NewGetServiceApiListResponse() (response *GetServiceApiListResponse) {
	response = &GetServiceApiListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取服务及其API列表
func (c *Client) GetServiceApiList(request *GetServiceApiListRequest) (response *GetServiceApiListResponse, err error) {
	if request == nil {
		request = NewGetServiceApiListRequest()
	}
	response = NewGetServiceApiListResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateServiceConfigRequest() (request *UpdateServiceConfigRequest) {
	request = &UpdateServiceConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tag", APIVersion, "UpdateServiceConfig")
	return
}

func NewUpdateServiceConfigResponse() (response *UpdateServiceConfigResponse) {
	response = &UpdateServiceConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步yunapi配置
func (c *Client) UpdateServiceConfig(request *UpdateServiceConfigRequest) (response *UpdateServiceConfigResponse, err error) {
	if request == nil {
		request = NewUpdateServiceConfigRequest()
	}
	response = NewUpdateServiceConfigResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateTagStatusConfigRequest() (request *UpdateTagStatusConfigRequest) {
	request = &UpdateTagStatusConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tag", APIVersion, "UpdateTagStatusConfig")
	return
}

func NewUpdateTagStatusConfigResponse() (response *UpdateTagStatusConfigResponse) {
	response = &UpdateTagStatusConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改标签配置状态
func (c *Client) UpdateTagStatusConfig(request *UpdateTagStatusConfigRequest) (response *UpdateTagStatusConfigResponse, err error) {
	if request == nil {
		request = NewUpdateTagStatusConfigRequest()
	}
	response = NewUpdateTagStatusConfigResponse()
	err = c.Send(request, response)
	return
}
