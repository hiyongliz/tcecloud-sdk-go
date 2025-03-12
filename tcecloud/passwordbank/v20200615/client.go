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

package v20200615

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-06-15"

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

func NewOperationLogsRequest() (request *OperationLogsRequest) {
	request = &OperationLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("passwordbank", APIVersion, "OperationLogs")
	return
}

func NewOperationLogsResponse() (response *OperationLogsResponse) {
	response = &OperationLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取服务器密码操作记录
func (c *Client) OperationLogs(request *OperationLogsRequest) (response *OperationLogsResponse, err error) {
	if request == nil {
		request = NewOperationLogsRequest()
	}
	response = NewOperationLogsResponse()
	err = c.Send(request, response)
	return
}

func NewGetPasswordRuleRequest() (request *GetPasswordRuleRequest) {
	request = &GetPasswordRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("passwordbank", APIVersion, "GetPasswordRule")
	return
}

func NewGetPasswordRuleResponse() (response *GetPasswordRuleResponse) {
	response = &GetPasswordRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询密码规则设置
func (c *Client) GetPasswordRule(request *GetPasswordRuleRequest) (response *GetPasswordRuleResponse, err error) {
	if request == nil {
		request = NewGetPasswordRuleRequest()
	}
	response = NewGetPasswordRuleResponse()
	err = c.Send(request, response)
	return
}

func NewQuitFromBanksRequest() (request *QuitFromBanksRequest) {
	request = &QuitFromBanksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("passwordbank", APIVersion, "QuitFromBanks")
	return
}

func NewQuitFromBanksResponse() (response *QuitFromBanksResponse) {
	response = &QuitFromBanksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设备退库
func (c *Client) QuitFromBanks(request *QuitFromBanksRequest) (response *QuitFromBanksResponse, err error) {
	if request == nil {
		request = NewQuitFromBanksRequest()
	}
	response = NewQuitFromBanksResponse()
	err = c.Send(request, response)
	return
}

func NewResetPasswordRequest() (request *ResetPasswordRequest) {
	request = &ResetPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("passwordbank", APIVersion, "ResetPassword")
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

func NewDescribePasswordsRequest() (request *DescribePasswordsRequest) {
	request = &DescribePasswordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("passwordbank", APIVersion, "DescribePasswords")
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

func NewLogDetailRequest() (request *LogDetailRequest) {
	request = &LogDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("passwordbank", APIVersion, "LogDetail")
	return
}

func NewLogDetailResponse() (response *LogDetailResponse) {
	response = &LogDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询修改记录细节
func (c *Client) LogDetail(request *LogDetailRequest) (response *LogDetailResponse, err error) {
	if request == nil {
		request = NewLogDetailRequest()
	}
	response = NewLogDetailResponse()
	err = c.Send(request, response)
	return
}

func NewSetPasswordTypesRequest() (request *SetPasswordTypesRequest) {
	request = &SetPasswordTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("passwordbank", APIVersion, "SetPasswordTypes")
	return
}

func NewSetPasswordTypesResponse() (response *SetPasswordTypesResponse) {
	response = &SetPasswordTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改密码类型
func (c *Client) SetPasswordTypes(request *SetPasswordTypesRequest) (response *SetPasswordTypesResponse, err error) {
	if request == nil {
		request = NewSetPasswordTypesRequest()
	}
	response = NewSetPasswordTypesResponse()
	err = c.Send(request, response)
	return
}

func NewJoinToBanksRequest() (request *JoinToBanksRequest) {
	request = &JoinToBanksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("passwordbank", APIVersion, "JoinToBanks")
	return
}

func NewJoinToBanksResponse() (response *JoinToBanksResponse) {
	response = &JoinToBanksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设备入库
func (c *Client) JoinToBanks(request *JoinToBanksRequest) (response *JoinToBanksResponse, err error) {
	if request == nil {
		request = NewJoinToBanksRequest()
	}
	response = NewJoinToBanksResponse()
	err = c.Send(request, response)
	return
}

func NewJoinToBanksFromExternalRequest() (request *JoinToBanksFromExternalRequest) {
	request = &JoinToBanksFromExternalRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("passwordbank", APIVersion, "JoinToBanksFromExternal")
	return
}

func NewJoinToBanksFromExternalResponse() (response *JoinToBanksFromExternalResponse) {
	response = &JoinToBanksFromExternalResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 外部设备入库
func (c *Client) JoinToBanksFromExternal(request *JoinToBanksFromExternalRequest) (response *JoinToBanksFromExternalResponse, err error) {
	if request == nil {
		request = NewJoinToBanksFromExternalRequest()
	}
	response = NewJoinToBanksFromExternalResponse()
	err = c.Send(request, response)
	return
}

func NewPasswordResetTimeRequest() (request *PasswordResetTimeRequest) {
	request = &PasswordResetTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("passwordbank", APIVersion, "PasswordResetTime")
	return
}

func NewPasswordResetTimeResponse() (response *PasswordResetTimeResponse) {
	response = &PasswordResetTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询密码周期设置
func (c *Client) PasswordResetTime(request *PasswordResetTimeRequest) (response *PasswordResetTimeResponse, err error) {
	if request == nil {
		request = NewPasswordResetTimeRequest()
	}
	response = NewPasswordResetTimeResponse()
	err = c.Send(request, response)
	return
}

func NewPasswordsRequest() (request *PasswordsRequest) {
	request = &PasswordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("passwordbank", APIVersion, "Passwords")
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

func NewSetPasswordResetTimeRequest() (request *SetPasswordResetTimeRequest) {
	request = &SetPasswordResetTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("passwordbank", APIVersion, "SetPasswordResetTime")
	return
}

func NewSetPasswordResetTimeResponse() (response *SetPasswordResetTimeResponse) {
	response = &SetPasswordResetTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置密码周期
func (c *Client) SetPasswordResetTime(request *SetPasswordResetTimeRequest) (response *SetPasswordResetTimeResponse, err error) {
	if request == nil {
		request = NewSetPasswordResetTimeRequest()
	}
	response = NewSetPasswordResetTimeResponse()
	err = c.Send(request, response)
	return
}

func NewSetPasswordRuleRequest() (request *SetPasswordRuleRequest) {
	request = &SetPasswordRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("passwordbank", APIVersion, "SetPasswordRule")
	return
}

func NewSetPasswordRuleResponse() (response *SetPasswordRuleResponse) {
	response = &SetPasswordRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置密码规则
func (c *Client) SetPasswordRule(request *SetPasswordRuleRequest) (response *SetPasswordRuleResponse, err error) {
	if request == nil {
		request = NewSetPasswordRuleRequest()
	}
	response = NewSetPasswordRuleResponse()
	err = c.Send(request, response)
	return
}

func NewServerListsRequest() (request *ServerListsRequest) {
	request = &ServerListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("passwordbank", APIVersion, "ServerLists")
	return
}

func NewServerListsResponse() (response *ServerListsResponse) {
	response = &ServerListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取服务器信息列表
func (c *Client) ServerLists(request *ServerListsRequest) (response *ServerListsResponse, err error) {
	if request == nil {
		request = NewServerListsRequest()
	}
	response = NewServerListsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPasswordRequest() (request *ModifyPasswordRequest) {
	request = &ModifyPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("passwordbank", APIVersion, "ModifyPassword")
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
