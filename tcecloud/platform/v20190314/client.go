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

package v20190314

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-03-14"

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

func NewQueryStrategyParentInfoRequest() (request *QueryStrategyParentInfoRequest) {
	request = &QueryStrategyParentInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "QueryStrategyParentInfo")
	return
}

func NewQueryStrategyParentInfoResponse() (response *QueryStrategyParentInfoResponse) {
	response = &QueryStrategyParentInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 白名单管理——查询分组信息
func (c *Client) QueryStrategyParentInfo(request *QueryStrategyParentInfoRequest) (response *QueryStrategyParentInfoResponse, err error) {
	if request == nil {
		request = NewQueryStrategyParentInfoRequest()
	}
	response = NewQueryStrategyParentInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateStrategyWhitelistRequest() (request *CreateStrategyWhitelistRequest) {
	request = &CreateStrategyWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "CreateStrategyWhitelist")
	return
}

func NewCreateStrategyWhitelistResponse() (response *CreateStrategyWhitelistResponse) {
	response = &CreateStrategyWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 白名单管理——新建白名单
func (c *Client) CreateStrategyWhitelist(request *CreateStrategyWhitelistRequest) (response *CreateStrategyWhitelistResponse, err error) {
	if request == nil {
		request = NewCreateStrategyWhitelistRequest()
	}
	response = NewCreateStrategyWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePasswordsRequest() (request *DescribePasswordsRequest) {
	request = &DescribePasswordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "DescribePasswords")
	return
}

func NewDescribePasswordsResponse() (response *DescribePasswordsResponse) {
	response = &DescribePasswordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取密码库设备密码
func (c *Client) DescribePasswords(request *DescribePasswordsRequest) (response *DescribePasswordsResponse, err error) {
	if request == nil {
		request = NewDescribePasswordsRequest()
	}
	response = NewDescribePasswordsResponse()
	err = c.Send(request, response)
	return
}

func NewPasswordsRequest() (request *PasswordsRequest) {
	request = &PasswordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "Passwords")
	return
}

func NewPasswordsResponse() (response *PasswordsResponse) {
	response = &PasswordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取密码
func (c *Client) Passwords(request *PasswordsRequest) (response *PasswordsResponse, err error) {
	if request == nil {
		request = NewPasswordsRequest()
	}
	response = NewPasswordsResponse()
	err = c.Send(request, response)
	return
}

func NewQueryStrategyWhitelistRequest() (request *QueryStrategyWhitelistRequest) {
	request = &QueryStrategyWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "QueryStrategyWhitelist")
	return
}

func NewQueryStrategyWhitelistResponse() (response *QueryStrategyWhitelistResponse) {
	response = &QueryStrategyWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 白名单管理——查询白名单信息
func (c *Client) QueryStrategyWhitelist(request *QueryStrategyWhitelistRequest) (response *QueryStrategyWhitelistResponse, err error) {
	if request == nil {
		request = NewQueryStrategyWhitelistRequest()
	}
	response = NewQueryStrategyWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateIdpConfigRequest() (request *UpdateIdpConfigRequest) {
	request = &UpdateIdpConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "UpdateIdpConfig")
	return
}

func NewUpdateIdpConfigResponse() (response *UpdateIdpConfigResponse) {
	response = &UpdateIdpConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新idp配置
func (c *Client) UpdateIdpConfig(request *UpdateIdpConfigRequest) (response *UpdateIdpConfigResponse, err error) {
	if request == nil {
		request = NewUpdateIdpConfigRequest()
	}
	response = NewUpdateIdpConfigResponse()
	err = c.Send(request, response)
	return
}

func NewSendActiveMailRequest() (request *SendActiveMailRequest) {
	request = &SendActiveMailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "SendActiveMail")
	return
}

func NewSendActiveMailResponse() (response *SendActiveMailResponse) {
	response = &SendActiveMailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客户账号管理——发送激活邮件
func (c *Client) SendActiveMail(request *SendActiveMailRequest) (response *SendActiveMailResponse, err error) {
	if request == nil {
		request = NewSendActiveMailRequest()
	}
	response = NewSendActiveMailResponse()
	err = c.Send(request, response)
	return
}

func NewAddOauthConfigRequest() (request *AddOauthConfigRequest) {
	request = &AddOauthConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "AddOauthConfig")
	return
}

func NewAddOauthConfigResponse() (response *AddOauthConfigResponse) {
	response = &AddOauthConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增oauth配置
func (c *Client) AddOauthConfig(request *AddOauthConfigRequest) (response *AddOauthConfigResponse, err error) {
	if request == nil {
		request = NewAddOauthConfigRequest()
	}
	response = NewAddOauthConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetIdpConfigRequest() (request *GetIdpConfigRequest) {
	request = &GetIdpConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "GetIdpConfig")
	return
}

func NewGetIdpConfigResponse() (response *GetIdpConfigResponse) {
	response = &GetIdpConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取idp配置
func (c *Client) GetIdpConfig(request *GetIdpConfigRequest) (response *GetIdpConfigResponse, err error) {
	if request == nil {
		request = NewGetIdpConfigRequest()
	}
	response = NewGetIdpConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteStrategyWhitelistRequest() (request *DeleteStrategyWhitelistRequest) {
	request = &DeleteStrategyWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "DeleteStrategyWhitelist")
	return
}

func NewDeleteStrategyWhitelistResponse() (response *DeleteStrategyWhitelistResponse) {
	response = &DeleteStrategyWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 白名单管理——删除白名单
func (c *Client) DeleteStrategyWhitelist(request *DeleteStrategyWhitelistRequest) (response *DeleteStrategyWhitelistResponse, err error) {
	if request == nil {
		request = NewDeleteStrategyWhitelistRequest()
	}
	response = NewDeleteStrategyWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewSetPasswordTypesRequest() (request *SetPasswordTypesRequest) {
	request = &SetPasswordTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "SetPasswordTypes")
	return
}

func NewSetPasswordTypesResponse() (response *SetPasswordTypesResponse) {
	response = &SetPasswordTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重新设置密码类型(长期或临时密码)并修改密码
func (c *Client) SetPasswordTypes(request *SetPasswordTypesRequest) (response *SetPasswordTypesResponse, err error) {
	if request == nil {
		request = NewSetPasswordTypesRequest()
	}
	response = NewSetPasswordTypesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAccountRequest() (request *CreateAccountRequest) {
	request = &CreateAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "CreateAccount")
	return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
	response = &CreateAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户账号管理——创建/导入账号
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
	if request == nil {
		request = NewCreateAccountRequest()
	}
	response = NewCreateAccountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPasswordRequest() (request *ModifyPasswordRequest) {
	request = &ModifyPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "ModifyPassword")
	return
}

func NewModifyPasswordResponse() (response *ModifyPasswordResponse) {
	response = &ModifyPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改设备密码
func (c *Client) ModifyPassword(request *ModifyPasswordRequest) (response *ModifyPasswordResponse, err error) {
	if request == nil {
		request = NewModifyPasswordRequest()
	}
	response = NewModifyPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewModifyStrategyWhitelistRequest() (request *ModifyStrategyWhitelistRequest) {
	request = &ModifyStrategyWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "ModifyStrategyWhitelist")
	return
}

func NewModifyStrategyWhitelistResponse() (response *ModifyStrategyWhitelistResponse) {
	response = &ModifyStrategyWhitelistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 白名单管理——修改白名单
func (c *Client) ModifyStrategyWhitelist(request *ModifyStrategyWhitelistRequest) (response *ModifyStrategyWhitelistResponse, err error) {
	if request == nil {
		request = NewModifyStrategyWhitelistRequest()
	}
	response = NewModifyStrategyWhitelistResponse()
	err = c.Send(request, response)
	return
}

func NewGetCustomSubAccountRequest() (request *GetCustomSubAccountRequest) {
	request = &GetCustomSubAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "GetCustomSubAccount")
	return
}

func NewGetCustomSubAccountResponse() (response *GetCustomSubAccountResponse) {
	response = &GetCustomSubAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过主账号拉取子账号
func (c *Client) GetCustomSubAccount(request *GetCustomSubAccountRequest) (response *GetCustomSubAccountResponse, err error) {
	if request == nil {
		request = NewGetCustomSubAccountRequest()
	}
	response = NewGetCustomSubAccountResponse()
	err = c.Send(request, response)
	return
}

func NewEnabledIdpConfigRequest() (request *EnabledIdpConfigRequest) {
	request = &EnabledIdpConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "EnabledIdpConfig")
	return
}

func NewEnabledIdpConfigResponse() (response *EnabledIdpConfigResponse) {
	response = &EnabledIdpConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用idp登录
func (c *Client) EnabledIdpConfig(request *EnabledIdpConfigRequest) (response *EnabledIdpConfigResponse, err error) {
	if request == nil {
		request = NewEnabledIdpConfigRequest()
	}
	response = NewEnabledIdpConfigResponse()
	err = c.Send(request, response)
	return
}

func NewResetPasswordRequest() (request *ResetPasswordRequest) {
	request = &ResetPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "ResetPassword")
	return
}

func NewResetPasswordResponse() (response *ResetPasswordResponse) {
	response = &ResetPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置密码库设备密码
func (c *Client) ResetPassword(request *ResetPasswordRequest) (response *ResetPasswordResponse, err error) {
	if request == nil {
		request = NewResetPasswordRequest()
	}
	response = NewResetPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewDisabledIdpConfigRequest() (request *DisabledIdpConfigRequest) {
	request = &DisabledIdpConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "DisabledIdpConfig")
	return
}

func NewDisabledIdpConfigResponse() (response *DisabledIdpConfigResponse) {
	response = &DisabledIdpConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭idp登录
func (c *Client) DisabledIdpConfig(request *DisabledIdpConfigRequest) (response *DisabledIdpConfigResponse, err error) {
	if request == nil {
		request = NewDisabledIdpConfigRequest()
	}
	response = NewDisabledIdpConfigResponse()
	err = c.Send(request, response)
	return
}

func NewQueryPasswordsRequest() (request *QueryPasswordsRequest) {
	request = &QueryPasswordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "QueryPasswords")
	return
}

func NewQueryPasswordsResponse() (response *QueryPasswordsResponse) {
	response = &QueryPasswordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询密码
func (c *Client) QueryPasswords(request *QueryPasswordsRequest) (response *QueryPasswordsResponse, err error) {
	if request == nil {
		request = NewQueryPasswordsRequest()
	}
	response = NewQueryPasswordsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOauthConfigRequest() (request *UpdateOauthConfigRequest) {
	request = &UpdateOauthConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "UpdateOauthConfig")
	return
}

func NewUpdateOauthConfigResponse() (response *UpdateOauthConfigResponse) {
	response = &UpdateOauthConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新oauth配置
func (c *Client) UpdateOauthConfig(request *UpdateOauthConfigRequest) (response *UpdateOauthConfigResponse, err error) {
	if request == nil {
		request = NewUpdateOauthConfigRequest()
	}
	response = NewUpdateOauthConfigResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCustomAccountRequest() (request *QueryCustomAccountRequest) {
	request = &QueryCustomAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "QueryCustomAccount")
	return
}

func NewQueryCustomAccountResponse() (response *QueryCustomAccountResponse) {
	response = &QueryCustomAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端客户管理查询租户端账号
func (c *Client) QueryCustomAccount(request *QueryCustomAccountRequest) (response *QueryCustomAccountResponse, err error) {
	if request == nil {
		request = NewQueryCustomAccountRequest()
	}
	response = NewQueryCustomAccountResponse()
	err = c.Send(request, response)
	return
}

func NewAddIdpConfigRequest() (request *AddIdpConfigRequest) {
	request = &AddIdpConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "AddIdpConfig")
	return
}

func NewAddIdpConfigResponse() (response *AddIdpConfigResponse) {
	response = &AddIdpConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加idp配置
func (c *Client) AddIdpConfig(request *AddIdpConfigRequest) (response *AddIdpConfigResponse, err error) {
	if request == nil {
		request = NewAddIdpConfigRequest()
	}
	response = NewAddIdpConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBPasswordRequest() (request *DescribeDBPasswordRequest) {
	request = &DescribeDBPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "DescribeDBPassword")
	return
}

func NewDescribeDBPasswordResponse() (response *DescribeDBPasswordResponse) {
	response = &DescribeDBPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 数据库密码查询接口
func (c *Client) DescribeDBPassword(request *DescribeDBPasswordRequest) (response *DescribeDBPasswordResponse, err error) {
	if request == nil {
		request = NewDescribeDBPasswordRequest()
	}
	response = NewDescribeDBPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewGetCustomAccountRequest() (request *GetCustomAccountRequest) {
	request = &GetCustomAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform", APIVersion, "GetCustomAccount")
	return
}

func NewGetCustomAccountResponse() (response *GetCustomAccountResponse) {
	response = &GetCustomAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取全部租户端主账号
func (c *Client) GetCustomAccount(request *GetCustomAccountRequest) (response *GetCustomAccountResponse, err error) {
	if request == nil {
		request = NewGetCustomAccountRequest()
	}
	response = NewGetCustomAccountResponse()
	err = c.Send(request, response)
	return
}
