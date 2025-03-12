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

package v20220121

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2022-01-21"

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

func NewGetSelfApiKeyRequest() (request *GetSelfApiKeyRequest) {
	request = &GetSelfApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetSelfApiKey")
	return
}

func NewGetSelfApiKeyResponse() (response *GetSelfApiKeyResponse) {
	response = &GetSelfApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取自己的密钥
func (c *Client) GetSelfApiKey(request *GetSelfApiKeyRequest) (response *GetSelfApiKeyResponse, err error) {
	if request == nil {
		request = NewGetSelfApiKeyRequest()
	}
	response = NewGetSelfApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewQuerySelfApiKeyRequest() (request *QuerySelfApiKeyRequest) {
	request = &QuerySelfApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "QuerySelfApiKey")
	return
}

func NewQuerySelfApiKeyResponse() (response *QuerySelfApiKeyResponse) {
	response = &QuerySelfApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取自己的API密钥列表
func (c *Client) QuerySelfApiKey(request *QuerySelfApiKeyRequest) (response *QuerySelfApiKeyResponse, err error) {
	if request == nil {
		request = NewQuerySelfApiKeyRequest()
	}
	response = NewQuerySelfApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSelfApiKeyRequest() (request *DeleteSelfApiKeyRequest) {
	request = &DeleteSelfApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "DeleteSelfApiKey")
	return
}

func NewDeleteSelfApiKeyResponse() (response *DeleteSelfApiKeyResponse) {
	response = &DeleteSelfApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除自己的密钥
func (c *Client) DeleteSelfApiKey(request *DeleteSelfApiKeyRequest) (response *DeleteSelfApiKeyResponse, err error) {
	if request == nil {
		request = NewDeleteSelfApiKeyRequest()
	}
	response = NewDeleteSelfApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewListSelfAttachedUserAllPoliciesRequest() (request *ListSelfAttachedUserAllPoliciesRequest) {
	request = &ListSelfAttachedUserAllPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "ListSelfAttachedUserAllPolicies")
	return
}

func NewListSelfAttachedUserAllPoliciesResponse() (response *ListSelfAttachedUserAllPoliciesResponse) {
	response = &ListSelfAttachedUserAllPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出自己关联的策略（包括随组关联）
func (c *Client) ListSelfAttachedUserAllPolicies(request *ListSelfAttachedUserAllPoliciesRequest) (response *ListSelfAttachedUserAllPoliciesResponse, err error) {
	if request == nil {
		request = NewListSelfAttachedUserAllPoliciesRequest()
	}
	response = NewListSelfAttachedUserAllPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewModifySelfApiKeyRequest() (request *ModifySelfApiKeyRequest) {
	request = &ModifySelfApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "ModifySelfApiKey")
	return
}

func NewModifySelfApiKeyResponse() (response *ModifySelfApiKeyResponse) {
	response = &ModifySelfApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改自己的密钥
func (c *Client) ModifySelfApiKey(request *ModifySelfApiKeyRequest) (response *ModifySelfApiKeyResponse, err error) {
	if request == nil {
		request = NewModifySelfApiKeyRequest()
	}
	response = NewModifySelfApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewEnableSelfApiKeyRequest() (request *EnableSelfApiKeyRequest) {
	request = &EnableSelfApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "EnableSelfApiKey")
	return
}

func NewEnableSelfApiKeyResponse() (response *EnableSelfApiKeyResponse) {
	response = &EnableSelfApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用自己的密钥
func (c *Client) EnableSelfApiKey(request *EnableSelfApiKeyRequest) (response *EnableSelfApiKeyResponse, err error) {
	if request == nil {
		request = NewEnableSelfApiKeyRequest()
	}
	response = NewEnableSelfApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDisableSelfApiKeyRequest() (request *DisableSelfApiKeyRequest) {
	request = &DisableSelfApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "DisableSelfApiKey")
	return
}

func NewDisableSelfApiKeyResponse() (response *DisableSelfApiKeyResponse) {
	response = &DisableSelfApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 禁用自己的密钥
func (c *Client) DisableSelfApiKey(request *DisableSelfApiKeyRequest) (response *DisableSelfApiKeyResponse, err error) {
	if request == nil {
		request = NewDisableSelfApiKeyRequest()
	}
	response = NewDisableSelfApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSelfApiKeyRequest() (request *CreateSelfApiKeyRequest) {
	request = &CreateSelfApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "CreateSelfApiKey")
	return
}

func NewCreateSelfApiKeyResponse() (response *CreateSelfApiKeyResponse) {
	response = &CreateSelfApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建自己的密钥
func (c *Client) CreateSelfApiKey(request *CreateSelfApiKeyRequest) (response *CreateSelfApiKeyResponse, err error) {
	if request == nil {
		request = NewCreateSelfApiKeyRequest()
	}
	response = NewCreateSelfApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApiKeyRequest() (request *ModifyApiKeyRequest) {
	request = &ModifyApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "ModifyApiKey")
	return
}

func NewModifyApiKeyResponse() (response *ModifyApiKeyResponse) {
	response = &ModifyApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改密钥
func (c *Client) ModifyApiKey(request *ModifyApiKeyRequest) (response *ModifyApiKeyResponse, err error) {
	if request == nil {
		request = NewModifyApiKeyRequest()
	}
	response = NewModifyApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewListSelfGroupsForUserRequest() (request *ListSelfGroupsForUserRequest) {
	request = &ListSelfGroupsForUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "ListSelfGroupsForUser")
	return
}

func NewListSelfGroupsForUserResponse() (response *ListSelfGroupsForUserResponse) {
	response = &ListSelfGroupsForUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询子账户所属用户组列表,无鉴权
func (c *Client) ListSelfGroupsForUser(request *ListSelfGroupsForUserRequest) (response *ListSelfGroupsForUserResponse, err error) {
	if request == nil {
		request = NewListSelfGroupsForUserRequest()
	}
	response = NewListSelfGroupsForUserResponse()
	err = c.Send(request, response)
	return
}
