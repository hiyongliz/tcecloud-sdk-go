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

package v20210808

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2021-08-08"

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

func NewListCustomOpenPlatformRequest() (request *ListCustomOpenPlatformRequest) {
	request = &ListCustomOpenPlatformRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("application", APIVersion, "ListCustomOpenPlatform")
	return
}

func NewListCustomOpenPlatformResponse() (response *ListCustomOpenPlatformResponse) {
	response = &ListCustomOpenPlatformResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户端开放平台列表
func (c *Client) ListCustomOpenPlatform(request *ListCustomOpenPlatformRequest) (response *ListCustomOpenPlatformResponse, err error) {
	if request == nil {
		request = NewListCustomOpenPlatformRequest()
	}
	response = NewListCustomOpenPlatformResponse()
	err = c.Send(request, response)
	return
}

func NewGetCustomOpenPlatformBasicInfoRequest() (request *GetCustomOpenPlatformBasicInfoRequest) {
	request = &GetCustomOpenPlatformBasicInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("application", APIVersion, "GetCustomOpenPlatformBasicInfo")
	return
}

func NewGetCustomOpenPlatformBasicInfoResponse() (response *GetCustomOpenPlatformBasicInfoResponse) {
	response = &GetCustomOpenPlatformBasicInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取oauth应用基础信息
func (c *Client) GetCustomOpenPlatformBasicInfo(request *GetCustomOpenPlatformBasicInfoRequest) (response *GetCustomOpenPlatformBasicInfoResponse, err error) {
	if request == nil {
		request = NewGetCustomOpenPlatformBasicInfoRequest()
	}
	response = NewGetCustomOpenPlatformBasicInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomOpenPlatformRequest() (request *ModifyCustomOpenPlatformRequest) {
	request = &ModifyCustomOpenPlatformRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("application", APIVersion, "ModifyCustomOpenPlatform")
	return
}

func NewModifyCustomOpenPlatformResponse() (response *ModifyCustomOpenPlatformResponse) {
	response = &ModifyCustomOpenPlatformResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改开放平台配置信息
func (c *Client) ModifyCustomOpenPlatform(request *ModifyCustomOpenPlatformRequest) (response *ModifyCustomOpenPlatformResponse, err error) {
	if request == nil {
		request = NewModifyCustomOpenPlatformRequest()
	}
	response = NewModifyCustomOpenPlatformResponse()
	err = c.Send(request, response)
	return
}

func NewGetCustomOpenPlatformSecretInfoRequest() (request *GetCustomOpenPlatformSecretInfoRequest) {
	request = &GetCustomOpenPlatformSecretInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("application", APIVersion, "GetCustomOpenPlatformSecretInfo")
	return
}

func NewGetCustomOpenPlatformSecretInfoResponse() (response *GetCustomOpenPlatformSecretInfoResponse) {
	response = &GetCustomOpenPlatformSecretInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取客户oauth 应用秘钥信息
func (c *Client) GetCustomOpenPlatformSecretInfo(request *GetCustomOpenPlatformSecretInfoRequest) (response *GetCustomOpenPlatformSecretInfoResponse, err error) {
	if request == nil {
		request = NewGetCustomOpenPlatformSecretInfoRequest()
	}
	response = NewGetCustomOpenPlatformSecretInfoResponse()
	err = c.Send(request, response)
	return
}

func NewManageCustomOpenPlatformRequest() (request *ManageCustomOpenPlatformRequest) {
	request = &ManageCustomOpenPlatformRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("application", APIVersion, "ManageCustomOpenPlatform")
	return
}

func NewManageCustomOpenPlatformResponse() (response *ManageCustomOpenPlatformResponse) {
	response = &ManageCustomOpenPlatformResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理客户应用相关
func (c *Client) ManageCustomOpenPlatform(request *ManageCustomOpenPlatformRequest) (response *ManageCustomOpenPlatformResponse, err error) {
	if request == nil {
		request = NewManageCustomOpenPlatformRequest()
	}
	response = NewManageCustomOpenPlatformResponse()
	err = c.Send(request, response)
	return
}

func NewApplyCustomOpenPlatformRequest() (request *ApplyCustomOpenPlatformRequest) {
	request = &ApplyCustomOpenPlatformRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("application", APIVersion, "ApplyCustomOpenPlatform")
	return
}

func NewApplyCustomOpenPlatformResponse() (response *ApplyCustomOpenPlatformResponse) {
	response = &ApplyCustomOpenPlatformResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 申请租户端开放平台
func (c *Client) ApplyCustomOpenPlatform(request *ApplyCustomOpenPlatformRequest) (response *ApplyCustomOpenPlatformResponse, err error) {
	if request == nil {
		request = NewApplyCustomOpenPlatformRequest()
	}
	response = NewApplyCustomOpenPlatformResponse()
	err = c.Send(request, response)
	return
}
