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

package v20201202

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-12-02"

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

func NewGetWorkWeixinOpenAppConfigRequest() (request *GetWorkWeixinOpenAppConfigRequest) {
	request = &GetWorkWeixinOpenAppConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("open", APIVersion, "GetWorkWeixinOpenAppConfig")
	return
}

func NewGetWorkWeixinOpenAppConfigResponse() (response *GetWorkWeixinOpenAppConfigResponse) {
	response = &GetWorkWeixinOpenAppConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取企业微信内部应用配置
func (c *Client) GetWorkWeixinOpenAppConfig(request *GetWorkWeixinOpenAppConfigRequest) (response *GetWorkWeixinOpenAppConfigResponse, err error) {
	if request == nil {
		request = NewGetWorkWeixinOpenAppConfigRequest()
	}
	response = NewGetWorkWeixinOpenAppConfigResponse()
	err = c.Send(request, response)
	return
}

func NewListIdentityProviderRequest() (request *ListIdentityProviderRequest) {
	request = &ListIdentityProviderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("open", APIVersion, "ListIdentityProvider")
	return
}

func NewListIdentityProviderResponse() (response *ListIdentityProviderResponse) {
	response = &ListIdentityProviderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取企业认证源列表。
func (c *Client) ListIdentityProvider(request *ListIdentityProviderRequest) (response *ListIdentityProviderResponse, err error) {
	if request == nil {
		request = NewListIdentityProviderRequest()
	}
	response = NewListIdentityProviderResponse()
	err = c.Send(request, response)
	return
}

func NewBatchBindWorkWeixinAccountRequest() (request *BatchBindWorkWeixinAccountRequest) {
	request = &BatchBindWorkWeixinAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("open", APIVersion, "BatchBindWorkWeixinAccount")
	return
}

func NewBatchBindWorkWeixinAccountResponse() (response *BatchBindWorkWeixinAccountResponse) {
	response = &BatchBindWorkWeixinAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// BatchBindWorkWeixinAccount
func (c *Client) BatchBindWorkWeixinAccount(request *BatchBindWorkWeixinAccountRequest) (response *BatchBindWorkWeixinAccountResponse, err error) {
	if request == nil {
		request = NewBatchBindWorkWeixinAccountRequest()
	}
	response = NewBatchBindWorkWeixinAccountResponse()
	err = c.Send(request, response)
	return
}

func NewCreateWorkWeixinOpenAppConfigRequest() (request *CreateWorkWeixinOpenAppConfigRequest) {
	request = &CreateWorkWeixinOpenAppConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("open", APIVersion, "CreateWorkWeixinOpenAppConfig")
	return
}

func NewCreateWorkWeixinOpenAppConfigResponse() (response *CreateWorkWeixinOpenAppConfigResponse) {
	response = &CreateWorkWeixinOpenAppConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建企业微信内部应用配置
func (c *Client) CreateWorkWeixinOpenAppConfig(request *CreateWorkWeixinOpenAppConfigRequest) (response *CreateWorkWeixinOpenAppConfigResponse, err error) {
	if request == nil {
		request = NewCreateWorkWeixinOpenAppConfigRequest()
	}
	response = NewCreateWorkWeixinOpenAppConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetWorkWeixinAuthorizationScopeRequest() (request *GetWorkWeixinAuthorizationScopeRequest) {
	request = &GetWorkWeixinAuthorizationScopeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("open", APIVersion, "GetWorkWeixinAuthorizationScope")
	return
}

func NewGetWorkWeixinAuthorizationScopeResponse() (response *GetWorkWeixinAuthorizationScopeResponse) {
	response = &GetWorkWeixinAuthorizationScopeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetWorkWeixinAuthorizationScope
func (c *Client) GetWorkWeixinAuthorizationScope(request *GetWorkWeixinAuthorizationScopeRequest) (response *GetWorkWeixinAuthorizationScopeResponse, err error) {
	if request == nil {
		request = NewGetWorkWeixinAuthorizationScopeRequest()
	}
	response = NewGetWorkWeixinAuthorizationScopeResponse()
	err = c.Send(request, response)
	return
}

func NewTestLdapRequest() (request *TestLdapRequest) {
	request = &TestLdapRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("open", APIVersion, "TestLdap")
	return
}

func NewTestLdapResponse() (response *TestLdapResponse) {
	response = &TestLdapResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 测试ldap认证源。
func (c *Client) TestLdap(request *TestLdapRequest) (response *TestLdapResponse, err error) {
	if request == nil {
		request = NewTestLdapRequest()
	}
	response = NewTestLdapResponse()
	err = c.Send(request, response)
	return
}

func NewGetLdapIdpConfigRequest() (request *GetLdapIdpConfigRequest) {
	request = &GetLdapIdpConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("open", APIVersion, "GetLdapIdpConfig")
	return
}

func NewGetLdapIdpConfigResponse() (response *GetLdapIdpConfigResponse) {
	response = &GetLdapIdpConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ldap配置
func (c *Client) GetLdapIdpConfig(request *GetLdapIdpConfigRequest) (response *GetLdapIdpConfigResponse, err error) {
	if request == nil {
		request = NewGetLdapIdpConfigRequest()
	}
	response = NewGetLdapIdpConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetWorkWeixinOpenAppMemberRequest() (request *GetWorkWeixinOpenAppMemberRequest) {
	request = &GetWorkWeixinOpenAppMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("open", APIVersion, "GetWorkWeixinOpenAppMember")
	return
}

func NewGetWorkWeixinOpenAppMemberResponse() (response *GetWorkWeixinOpenAppMemberResponse) {
	response = &GetWorkWeixinOpenAppMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取企业微信内部应用成员信息
func (c *Client) GetWorkWeixinOpenAppMember(request *GetWorkWeixinOpenAppMemberRequest) (response *GetWorkWeixinOpenAppMemberResponse, err error) {
	if request == nil {
		request = NewGetWorkWeixinOpenAppMemberRequest()
	}
	response = NewGetWorkWeixinOpenAppMemberResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateLdapIdpRequest() (request *UpdateLdapIdpRequest) {
	request = &UpdateLdapIdpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("open", APIVersion, "UpdateLdapIdp")
	return
}

func NewUpdateLdapIdpResponse() (response *UpdateLdapIdpResponse) {
	response = &UpdateLdapIdpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新ldap认证源。
func (c *Client) UpdateLdapIdp(request *UpdateLdapIdpRequest) (response *UpdateLdapIdpResponse, err error) {
	if request == nil {
		request = NewUpdateLdapIdpRequest()
	}
	response = NewUpdateLdapIdpResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLdapIdpRequest() (request *CreateLdapIdpRequest) {
	request = &CreateLdapIdpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("open", APIVersion, "CreateLdapIdp")
	return
}

func NewCreateLdapIdpResponse() (response *CreateLdapIdpResponse) {
	response = &CreateLdapIdpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建ldap认证源
func (c *Client) CreateLdapIdp(request *CreateLdapIdpRequest) (response *CreateLdapIdpResponse, err error) {
	if request == nil {
		request = NewCreateLdapIdpRequest()
	}
	response = NewCreateLdapIdpResponse()
	err = c.Send(request, response)
	return
}

func NewBindWorkWeixinAccountRequest() (request *BindWorkWeixinAccountRequest) {
	request = &BindWorkWeixinAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("open", APIVersion, "BindWorkWeixinAccount")
	return
}

func NewBindWorkWeixinAccountResponse() (response *BindWorkWeixinAccountResponse) {
	response = &BindWorkWeixinAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 子账号绑定企业微信
func (c *Client) BindWorkWeixinAccount(request *BindWorkWeixinAccountRequest) (response *BindWorkWeixinAccountResponse, err error) {
	if request == nil {
		request = NewBindWorkWeixinAccountRequest()
	}
	response = NewBindWorkWeixinAccountResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateWorkWeixinConfigRequest() (request *UpdateWorkWeixinConfigRequest) {
	request = &UpdateWorkWeixinConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("open", APIVersion, "UpdateWorkWeixinConfig")
	return
}

func NewUpdateWorkWeixinConfigResponse() (response *UpdateWorkWeixinConfigResponse) {
	response = &UpdateWorkWeixinConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// UpdateWorkWeixinConfig
func (c *Client) UpdateWorkWeixinConfig(request *UpdateWorkWeixinConfigRequest) (response *UpdateWorkWeixinConfigResponse, err error) {
	if request == nil {
		request = NewUpdateWorkWeixinConfigRequest()
	}
	response = NewUpdateWorkWeixinConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetWorkWeixinConfigRequest() (request *GetWorkWeixinConfigRequest) {
	request = &GetWorkWeixinConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("open", APIVersion, "GetWorkWeixinConfig")
	return
}

func NewGetWorkWeixinConfigResponse() (response *GetWorkWeixinConfigResponse) {
	response = &GetWorkWeixinConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetWorkWeixinConfig
func (c *Client) GetWorkWeixinConfig(request *GetWorkWeixinConfigRequest) (response *GetWorkWeixinConfigResponse, err error) {
	if request == nil {
		request = NewGetWorkWeixinConfigRequest()
	}
	response = NewGetWorkWeixinConfigResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateWorkWeixinOpenAppConfigRequest() (request *UpdateWorkWeixinOpenAppConfigRequest) {
	request = &UpdateWorkWeixinOpenAppConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("open", APIVersion, "UpdateWorkWeixinOpenAppConfig")
	return
}

func NewUpdateWorkWeixinOpenAppConfigResponse() (response *UpdateWorkWeixinOpenAppConfigResponse) {
	response = &UpdateWorkWeixinOpenAppConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新企业微信内部应用配置
func (c *Client) UpdateWorkWeixinOpenAppConfig(request *UpdateWorkWeixinOpenAppConfigRequest) (response *UpdateWorkWeixinOpenAppConfigResponse, err error) {
	if request == nil {
		request = NewUpdateWorkWeixinOpenAppConfigRequest()
	}
	response = NewUpdateWorkWeixinOpenAppConfigResponse()
	err = c.Send(request, response)
	return
}
