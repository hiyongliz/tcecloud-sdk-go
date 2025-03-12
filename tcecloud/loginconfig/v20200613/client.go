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

package v20200613

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-06-13"

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

func NewUpdatePasswordConfigRequest() (request *UpdatePasswordConfigRequest) {
	request = &UpdatePasswordConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("loginconfig", APIVersion, "UpdatePasswordConfig")
	return
}

func NewUpdatePasswordConfigResponse() (response *UpdatePasswordConfigResponse) {
	response = &UpdatePasswordConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新密码配置
func (c *Client) UpdatePasswordConfig(request *UpdatePasswordConfigRequest) (response *UpdatePasswordConfigResponse, err error) {
	if request == nil {
		request = NewUpdatePasswordConfigRequest()
	}
	response = NewUpdatePasswordConfigResponse()
	err = c.Send(request, response)
	return
}

func NewValidateLoginPolicyRequest() (request *ValidateLoginPolicyRequest) {
	request = &ValidateLoginPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("loginconfig", APIVersion, "ValidateLoginPolicy")
	return
}

func NewValidateLoginPolicyResponse() (response *ValidateLoginPolicyResponse) {
	response = &ValidateLoginPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ValidateLoginPolicy
func (c *Client) ValidateLoginPolicy(request *ValidateLoginPolicyRequest) (response *ValidateLoginPolicyResponse, err error) {
	if request == nil {
		request = NewValidateLoginPolicyRequest()
	}
	response = NewValidateLoginPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewGetPasswordConfigRequest() (request *GetPasswordConfigRequest) {
	request = &GetPasswordConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("loginconfig", APIVersion, "GetPasswordConfig")
	return
}

func NewGetPasswordConfigResponse() (response *GetPasswordConfigResponse) {
	response = &GetPasswordConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取密码登录配置
func (c *Client) GetPasswordConfig(request *GetPasswordConfigRequest) (response *GetPasswordConfigResponse, err error) {
	if request == nil {
		request = NewGetPasswordConfigRequest()
	}
	response = NewGetPasswordConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLoginPolicyRequest() (request *CreateLoginPolicyRequest) {
	request = &CreateLoginPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("loginconfig", APIVersion, "CreateLoginPolicy")
	return
}

func NewCreateLoginPolicyResponse() (response *CreateLoginPolicyResponse) {
	response = &CreateLoginPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateLoginPolicy
func (c *Client) CreateLoginPolicy(request *CreateLoginPolicyRequest) (response *CreateLoginPolicyResponse, err error) {
	if request == nil {
		request = NewCreateLoginPolicyRequest()
	}
	response = NewCreateLoginPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewGetLoginPolicyRequest() (request *GetLoginPolicyRequest) {
	request = &GetLoginPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("loginconfig", APIVersion, "GetLoginPolicy")
	return
}

func NewGetLoginPolicyResponse() (response *GetLoginPolicyResponse) {
	response = &GetLoginPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取登录策略
func (c *Client) GetLoginPolicy(request *GetLoginPolicyRequest) (response *GetLoginPolicyResponse, err error) {
	if request == nil {
		request = NewGetLoginPolicyRequest()
	}
	response = NewGetLoginPolicyResponse()
	err = c.Send(request, response)
	return
}
