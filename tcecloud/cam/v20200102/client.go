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

package v20200102

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-01-02"

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

func NewGetUidByUinArrRequest() (request *GetUidByUinArrRequest) {
	request = &GetUidByUinArrRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetUidByUinArr")
	return
}

func NewGetUidByUinArrResponse() (response *GetUidByUinArrResponse) {
	response = &GetUidByUinArrResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据uin列表返回对应的uid信息
func (c *Client) GetUidByUinArr(request *GetUidByUinArrRequest) (response *GetUidByUinArrResponse, err error) {
	if request == nil {
		request = NewGetUidByUinArrRequest()
	}
	response = NewGetUidByUinArrResponse()
	err = c.Send(request, response)
	return
}

func NewVerifyMenuRequest() (request *VerifyMenuRequest) {
	request = &VerifyMenuRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "VerifyMenu")
	return
}

func NewVerifyMenuResponse() (response *VerifyMenuResponse) {
	response = &VerifyMenuResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 菜单鉴权
func (c *Client) VerifyMenu(request *VerifyMenuRequest) (response *VerifyMenuResponse, err error) {
	if request == nil {
		request = NewVerifyMenuRequest()
	}
	response = NewVerifyMenuResponse()
	err = c.Send(request, response)
	return
}

func NewGetLoginActionRequest() (request *GetLoginActionRequest) {
	request = &GetLoginActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetLoginAction")
	return
}

func NewGetLoginActionResponse() (response *GetLoginActionResponse) {
	response = &GetLoginActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取登陆过程事件
func (c *Client) GetLoginAction(request *GetLoginActionRequest) (response *GetLoginActionResponse, err error) {
	if request == nil {
		request = NewGetLoginActionRequest()
	}
	response = NewGetLoginActionResponse()
	err = c.Send(request, response)
	return
}

func NewPreLoginRequest() (request *PreLoginRequest) {
	request = &PreLoginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "PreLogin")
	return
}

func NewPreLoginResponse() (response *PreLoginResponse) {
	response = &PreLoginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 账户登陆a
func (c *Client) PreLogin(request *PreLoginRequest) (response *PreLoginResponse, err error) {
	if request == nil {
		request = NewPreLoginRequest()
	}
	response = NewPreLoginResponse()
	err = c.Send(request, response)
	return
}

func NewDelLoginForbidRequest() (request *DelLoginForbidRequest) {
	request = &DelLoginForbidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DelLoginForbid")
	return
}

func NewDelLoginForbidResponse() (response *DelLoginForbidResponse) {
	response = &DelLoginForbidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除禁止登陆标记
func (c *Client) DelLoginForbid(request *DelLoginForbidRequest) (response *DelLoginForbidResponse, err error) {
	if request == nil {
		request = NewDelLoginForbidRequest()
	}
	response = NewDelLoginForbidResponse()
	err = c.Send(request, response)
	return
}

func NewCheckLoginRequest() (request *CheckLoginRequest) {
	request = &CheckLoginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "CheckLogin")
	return
}

func NewCheckLoginResponse() (response *CheckLoginResponse) {
	response = &CheckLoginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查是否可以登陆
func (c *Client) CheckLogin(request *CheckLoginRequest) (response *CheckLoginResponse, err error) {
	if request == nil {
		request = NewCheckLoginRequest()
	}
	response = NewCheckLoginResponse()
	err = c.Send(request, response)
	return
}
