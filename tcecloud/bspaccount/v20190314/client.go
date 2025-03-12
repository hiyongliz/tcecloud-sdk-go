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

func NewGetCustomAccountsByAppIdsRequest() (request *GetCustomAccountsByAppIdsRequest) {
	request = &GetCustomAccountsByAppIdsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "GetCustomAccountsByAppIds")
	return
}

func NewGetCustomAccountsByAppIdsResponse() (response *GetCustomAccountsByAppIdsResponse) {
	response = &GetCustomAccountsByAppIdsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 使用appid列表获取账号信息
func (c *Client) GetCustomAccountsByAppIds(request *GetCustomAccountsByAppIdsRequest) (response *GetCustomAccountsByAppIdsResponse, err error) {
	if request == nil {
		request = NewGetCustomAccountsByAppIdsRequest()
	}
	response = NewGetCustomAccountsByAppIdsResponse()
	err = c.Send(request, response)
	return
}

func NewCheckCustomAccountRequest() (request *CheckCustomAccountRequest) {
	request = &CheckCustomAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "CheckCustomAccount")
	return
}

func NewCheckCustomAccountResponse() (response *CheckCustomAccountResponse) {
	response = &CheckCustomAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查客户账号是否已存在
func (c *Client) CheckCustomAccount(request *CheckCustomAccountRequest) (response *CheckCustomAccountResponse, err error) {
	if request == nil {
		request = NewCheckCustomAccountRequest()
	}
	response = NewCheckCustomAccountResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCustomCategoryRequest() (request *CreateCustomCategoryRequest) {
	request = &CreateCustomCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "CreateCustomCategory")
	return
}

func NewCreateCustomCategoryResponse() (response *CreateCustomCategoryResponse) {
	response = &CreateCustomCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建客户类别
func (c *Client) CreateCustomCategory(request *CreateCustomCategoryRequest) (response *CreateCustomCategoryResponse, err error) {
	if request == nil {
		request = NewCreateCustomCategoryRequest()
	}
	response = NewCreateCustomCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewGetAccountBaseInfoRequest() (request *GetAccountBaseInfoRequest) {
	request = &GetAccountBaseInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "GetAccountBaseInfo")
	return
}

func NewGetAccountBaseInfoResponse() (response *GetAccountBaseInfoResponse) {
	response = &GetAccountBaseInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账号基本信息
func (c *Client) GetAccountBaseInfo(request *GetAccountBaseInfoRequest) (response *GetAccountBaseInfoResponse, err error) {
	if request == nil {
		request = NewGetAccountBaseInfoRequest()
	}
	response = NewGetAccountBaseInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomCategoryRequest() (request *ModifyCustomCategoryRequest) {
	request = &ModifyCustomCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "ModifyCustomCategory")
	return
}

func NewModifyCustomCategoryResponse() (response *ModifyCustomCategoryResponse) {
	response = &ModifyCustomCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改租户绑定类别信息
func (c *Client) ModifyCustomCategory(request *ModifyCustomCategoryRequest) (response *ModifyCustomCategoryResponse, err error) {
	if request == nil {
		request = NewModifyCustomCategoryRequest()
	}
	response = NewModifyCustomCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCustomAccountsIdentifyRequest() (request *UpdateCustomAccountsIdentifyRequest) {
	request = &UpdateCustomAccountsIdentifyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "UpdateCustomAccountsIdentify")
	return
}

func NewUpdateCustomAccountsIdentifyResponse() (response *UpdateCustomAccountsIdentifyResponse) {
	response = &UpdateCustomAccountsIdentifyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新租户端账号身份类型
func (c *Client) UpdateCustomAccountsIdentify(request *UpdateCustomAccountsIdentifyRequest) (response *UpdateCustomAccountsIdentifyResponse, err error) {
	if request == nil {
		request = NewUpdateCustomAccountsIdentifyRequest()
	}
	response = NewUpdateCustomAccountsIdentifyResponse()
	err = c.Send(request, response)
	return
}

func NewSendLinkPCloudAccountNoticeEmailRequest() (request *SendLinkPCloudAccountNoticeEmailRequest) {
	request = &SendLinkPCloudAccountNoticeEmailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "SendLinkPCloudAccountNoticeEmail")
	return
}

func NewSendLinkPCloudAccountNoticeEmailResponse() (response *SendLinkPCloudAccountNoticeEmailResponse) {
	response = &SendLinkPCloudAccountNoticeEmailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// SendLinkPCloudAccountNoticeEmail
func (c *Client) SendLinkPCloudAccountNoticeEmail(request *SendLinkPCloudAccountNoticeEmailRequest) (response *SendLinkPCloudAccountNoticeEmailResponse, err error) {
	if request == nil {
		request = NewSendLinkPCloudAccountNoticeEmailRequest()
	}
	response = NewSendLinkPCloudAccountNoticeEmailResponse()
	err = c.Send(request, response)
	return
}

func NewSendActiveMailRequest() (request *SendActiveMailRequest) {
	request = &SendActiveMailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "SendActiveMail")
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

func NewDisableCustomerApiKeyRequest() (request *DisableCustomerApiKeyRequest) {
	request = &DisableCustomerApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "DisableCustomerApiKey")
	return
}

func NewDisableCustomerApiKeyResponse() (response *DisableCustomerApiKeyResponse) {
	response = &DisableCustomerApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 禁用持久秘钥
func (c *Client) DisableCustomerApiKey(request *DisableCustomerApiKeyRequest) (response *DisableCustomerApiKeyResponse, err error) {
	if request == nil {
		request = NewDisableCustomerApiKeyRequest()
	}
	response = NewDisableCustomerApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewQueryStrategyParentInfoRequest() (request *QueryStrategyParentInfoRequest) {
	request = &QueryStrategyParentInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "QueryStrategyParentInfo")
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

func NewQueryCustomCategoryRequest() (request *QueryCustomCategoryRequest) {
	request = &QueryCustomCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "QueryCustomCategory")
	return
}

func NewQueryCustomCategoryResponse() (response *QueryCustomCategoryResponse) {
	response = &QueryCustomCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询客户类别信息
func (c *Client) QueryCustomCategory(request *QueryCustomCategoryRequest) (response *QueryCustomCategoryResponse, err error) {
	if request == nil {
		request = NewQueryCustomCategoryRequest()
	}
	response = NewQueryCustomCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewInsertOrUpdateCustomAccountAttributeValuesRequest() (request *InsertOrUpdateCustomAccountAttributeValuesRequest) {
	request = &InsertOrUpdateCustomAccountAttributeValuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "InsertOrUpdateCustomAccountAttributeValues")
	return
}

func NewInsertOrUpdateCustomAccountAttributeValuesResponse() (response *InsertOrUpdateCustomAccountAttributeValuesResponse) {
	response = &InsertOrUpdateCustomAccountAttributeValuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// InsertOrUpdateCustomAccountAttributeValues
func (c *Client) InsertOrUpdateCustomAccountAttributeValues(request *InsertOrUpdateCustomAccountAttributeValuesRequest) (response *InsertOrUpdateCustomAccountAttributeValuesResponse, err error) {
	if request == nil {
		request = NewInsertOrUpdateCustomAccountAttributeValuesRequest()
	}
	response = NewInsertOrUpdateCustomAccountAttributeValuesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCustomSubAccountRequest() (request *UpdateCustomSubAccountRequest) {
	request = &UpdateCustomSubAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "UpdateCustomSubAccount")
	return
}

func NewUpdateCustomSubAccountResponse() (response *UpdateCustomSubAccountResponse) {
	response = &UpdateCustomSubAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新租户子账号备注信息
func (c *Client) UpdateCustomSubAccount(request *UpdateCustomSubAccountRequest) (response *UpdateCustomSubAccountResponse, err error) {
	if request == nil {
		request = NewUpdateCustomSubAccountRequest()
	}
	response = NewUpdateCustomSubAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCustomCategoryRequest() (request *DeleteCustomCategoryRequest) {
	request = &DeleteCustomCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "DeleteCustomCategory")
	return
}

func NewDeleteCustomCategoryResponse() (response *DeleteCustomCategoryResponse) {
	response = &DeleteCustomCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除客户类别
func (c *Client) DeleteCustomCategory(request *DeleteCustomCategoryRequest) (response *DeleteCustomCategoryResponse, err error) {
	if request == nil {
		request = NewDeleteCustomCategoryRequest()
	}
	response = NewDeleteCustomCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCustomAccountRequest() (request *QueryCustomAccountRequest) {
	request = &QueryCustomAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "QueryCustomAccount")
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

func NewGetCustomAccountsByUinListRequest() (request *GetCustomAccountsByUinListRequest) {
	request = &GetCustomAccountsByUinListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "GetCustomAccountsByUinList")
	return
}

func NewGetCustomAccountsByUinListResponse() (response *GetCustomAccountsByUinListResponse) {
	response = &GetCustomAccountsByUinListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据uin列表批量获取租户账号信息
func (c *Client) GetCustomAccountsByUinList(request *GetCustomAccountsByUinListRequest) (response *GetCustomAccountsByUinListResponse, err error) {
	if request == nil {
		request = NewGetCustomAccountsByUinListRequest()
	}
	response = NewGetCustomAccountsByUinListResponse()
	err = c.Send(request, response)
	return
}

func NewEnableCustomerApiKeyRequest() (request *EnableCustomerApiKeyRequest) {
	request = &EnableCustomerApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "EnableCustomerApiKey")
	return
}

func NewEnableCustomerApiKeyResponse() (response *EnableCustomerApiKeyResponse) {
	response = &EnableCustomerApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用持久秘钥
func (c *Client) EnableCustomerApiKey(request *EnableCustomerApiKeyRequest) (response *EnableCustomerApiKeyResponse, err error) {
	if request == nil {
		request = NewEnableCustomerApiKeyRequest()
	}
	response = NewEnableCustomerApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewGetCustomerUrgentCodeRequest() (request *GetCustomerUrgentCodeRequest) {
	request = &GetCustomerUrgentCodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "GetCustomerUrgentCode")
	return
}

func NewGetCustomerUrgentCodeResponse() (response *GetCustomerUrgentCodeResponse) {
	response = &GetCustomerUrgentCodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户mfa紧急口令
func (c *Client) GetCustomerUrgentCode(request *GetCustomerUrgentCodeRequest) (response *GetCustomerUrgentCodeResponse, err error) {
	if request == nil {
		request = NewGetCustomerUrgentCodeRequest()
	}
	response = NewGetCustomerUrgentCodeResponse()
	err = c.Send(request, response)
	return
}

func NewGetFullCustomCategoryRequest() (request *GetFullCustomCategoryRequest) {
	request = &GetFullCustomCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "GetFullCustomCategory")
	return
}

func NewGetFullCustomCategoryResponse() (response *GetFullCustomCategoryResponse) {
	response = &GetFullCustomCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取全量客户类别信息
func (c *Client) GetFullCustomCategory(request *GetFullCustomCategoryRequest) (response *GetFullCustomCategoryResponse, err error) {
	if request == nil {
		request = NewGetFullCustomCategoryRequest()
	}
	response = NewGetFullCustomCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewGetCustomAccountRequest() (request *GetCustomAccountRequest) {
	request = &GetCustomAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "GetCustomAccount")
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

func NewModifyStrategyWhitelistRequest() (request *ModifyStrategyWhitelistRequest) {
	request = &ModifyStrategyWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "ModifyStrategyWhitelist")
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
	request.Init().WithApiInfo("bspaccount", APIVersion, "GetCustomSubAccount")
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

func NewGetCustomerApiKeyRequest() (request *GetCustomerApiKeyRequest) {
	request = &GetCustomerApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "GetCustomerApiKey")
	return
}

func NewGetCustomerApiKeyResponse() (response *GetCustomerApiKeyResponse) {
	response = &GetCustomerApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户账号持久秘钥
func (c *Client) GetCustomerApiKey(request *GetCustomerApiKeyRequest) (response *GetCustomerApiKeyResponse, err error) {
	if request == nil {
		request = NewGetCustomerApiKeyRequest()
	}
	response = NewGetCustomerApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewQueryStrategyWhitelistRequest() (request *QueryStrategyWhitelistRequest) {
	request = &QueryStrategyWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "QueryStrategyWhitelist")
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

func NewQueryCustomerApiKeyRequest() (request *QueryCustomerApiKeyRequest) {
	request = &QueryCustomerApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "QueryCustomerApiKey")
	return
}

func NewQueryCustomerApiKeyResponse() (response *QueryCustomerApiKeyResponse) {
	response = &QueryCustomerApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// QueryCustomerApiKey
func (c *Client) QueryCustomerApiKey(request *QueryCustomerApiKeyRequest) (response *QueryCustomerApiKeyResponse, err error) {
	if request == nil {
		request = NewQueryCustomerApiKeyRequest()
	}
	response = NewQueryCustomerApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewSeedOIDCLoginTokenRequest() (request *SeedOIDCLoginTokenRequest) {
	request = &SeedOIDCLoginTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "SeedOIDCLoginToken")
	return
}

func NewSeedOIDCLoginTokenResponse() (response *SeedOIDCLoginTokenResponse) {
	response = &SeedOIDCLoginTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 生成oidc临时token
func (c *Client) SeedOIDCLoginToken(request *SeedOIDCLoginTokenRequest) (response *SeedOIDCLoginTokenResponse, err error) {
	if request == nil {
		request = NewSeedOIDCLoginTokenRequest()
	}
	response = NewSeedOIDCLoginTokenResponse()
	err = c.Send(request, response)
	return
}

func NewCreateStrategyWhitelistRequest() (request *CreateStrategyWhitelistRequest) {
	request = &CreateStrategyWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "CreateStrategyWhitelist")
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

func NewCreateCustomerApiKeyRequest() (request *CreateCustomerApiKeyRequest) {
	request = &CreateCustomerApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "CreateCustomerApiKey")
	return
}

func NewCreateCustomerApiKeyResponse() (response *CreateCustomerApiKeyResponse) {
	response = &CreateCustomerApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建租户账号持久秘钥
func (c *Client) CreateCustomerApiKey(request *CreateCustomerApiKeyRequest) (response *CreateCustomerApiKeyResponse, err error) {
	if request == nil {
		request = NewCreateCustomerApiKeyRequest()
	}
	response = NewCreateCustomerApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCustomOwnerAccountRequest() (request *UpdateCustomOwnerAccountRequest) {
	request = &UpdateCustomOwnerAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "UpdateCustomOwnerAccount")
	return
}

func NewUpdateCustomOwnerAccountResponse() (response *UpdateCustomOwnerAccountResponse) {
	response = &UpdateCustomOwnerAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新租户主账号信息
func (c *Client) UpdateCustomOwnerAccount(request *UpdateCustomOwnerAccountRequest) (response *UpdateCustomOwnerAccountResponse, err error) {
	if request == nil {
		request = NewUpdateCustomOwnerAccountRequest()
	}
	response = NewUpdateCustomOwnerAccountResponse()
	err = c.Send(request, response)
	return
}

func NewAddCustomAccountAttributeValueRequest() (request *AddCustomAccountAttributeValueRequest) {
	request = &AddCustomAccountAttributeValueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "AddCustomAccountAttributeValue")
	return
}

func NewAddCustomAccountAttributeValueResponse() (response *AddCustomAccountAttributeValueResponse) {
	response = &AddCustomAccountAttributeValueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加租户端账号拓展属性值
func (c *Client) AddCustomAccountAttributeValue(request *AddCustomAccountAttributeValueRequest) (response *AddCustomAccountAttributeValueResponse, err error) {
	if request == nil {
		request = NewAddCustomAccountAttributeValueRequest()
	}
	response = NewAddCustomAccountAttributeValueResponse()
	err = c.Send(request, response)
	return
}

func NewGetCustomerSafeConfigRequest() (request *GetCustomerSafeConfigRequest) {
	request = &GetCustomerSafeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "GetCustomerSafeConfig")
	return
}

func NewGetCustomerSafeConfigResponse() (response *GetCustomerSafeConfigResponse) {
	response = &GetCustomerSafeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户安全认证配置
func (c *Client) GetCustomerSafeConfig(request *GetCustomerSafeConfigRequest) (response *GetCustomerSafeConfigResponse, err error) {
	if request == nil {
		request = NewGetCustomerSafeConfigRequest()
	}
	response = NewGetCustomerSafeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetPasswordRuleByCustomUinRequest() (request *GetPasswordRuleByCustomUinRequest) {
	request = &GetPasswordRuleByCustomUinRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "GetPasswordRuleByCustomUin")
	return
}

func NewGetPasswordRuleByCustomUinResponse() (response *GetPasswordRuleByCustomUinResponse) {
	response = &GetPasswordRuleByCustomUinResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据主账号uin获取对应的密码规则
func (c *Client) GetPasswordRuleByCustomUin(request *GetPasswordRuleByCustomUinRequest) (response *GetPasswordRuleByCustomUinResponse, err error) {
	if request == nil {
		request = NewGetPasswordRuleByCustomUinRequest()
	}
	response = NewGetPasswordRuleByCustomUinResponse()
	err = c.Send(request, response)
	return
}

func NewGetCustomAccountAttributesRequest() (request *GetCustomAccountAttributesRequest) {
	request = &GetCustomAccountAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "GetCustomAccountAttributes")
	return
}

func NewGetCustomAccountAttributesResponse() (response *GetCustomAccountAttributesResponse) {
	response = &GetCustomAccountAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetCustomAccountAttributes
func (c *Client) GetCustomAccountAttributes(request *GetCustomAccountAttributesRequest) (response *GetCustomAccountAttributesResponse, err error) {
	if request == nil {
		request = NewGetCustomAccountAttributesRequest()
	}
	response = NewGetCustomAccountAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCustomAccountAttributeValueRequest() (request *UpdateCustomAccountAttributeValueRequest) {
	request = &UpdateCustomAccountAttributeValueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "UpdateCustomAccountAttributeValue")
	return
}

func NewUpdateCustomAccountAttributeValueResponse() (response *UpdateCustomAccountAttributeValueResponse) {
	response = &UpdateCustomAccountAttributeValueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// UpdateCustomAccountAttributeValue
func (c *Client) UpdateCustomAccountAttributeValue(request *UpdateCustomAccountAttributeValueRequest) (response *UpdateCustomAccountAttributeValueResponse, err error) {
	if request == nil {
		request = NewUpdateCustomAccountAttributeValueRequest()
	}
	response = NewUpdateCustomAccountAttributeValueResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomIpLimitRequest() (request *ModifyCustomIpLimitRequest) {
	request = &ModifyCustomIpLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "ModifyCustomIpLimit")
	return
}

func NewModifyCustomIpLimitResponse() (response *ModifyCustomIpLimitResponse) {
	response = &ModifyCustomIpLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改客户账号Ip白名单/黑名单限制
func (c *Client) ModifyCustomIpLimit(request *ModifyCustomIpLimitRequest) (response *ModifyCustomIpLimitResponse, err error) {
	if request == nil {
		request = NewModifyCustomIpLimitRequest()
	}
	response = NewModifyCustomIpLimitResponse()
	err = c.Send(request, response)
	return
}

func NewGetCustomIpLimitRequest() (request *GetCustomIpLimitRequest) {
	request = &GetCustomIpLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "GetCustomIpLimit")
	return
}

func NewGetCustomIpLimitResponse() (response *GetCustomIpLimitResponse) {
	response = &GetCustomIpLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取客户限制IP策略
func (c *Client) GetCustomIpLimit(request *GetCustomIpLimitRequest) (response *GetCustomIpLimitResponse, err error) {
	if request == nil {
		request = NewGetCustomIpLimitRequest()
	}
	response = NewGetCustomIpLimitResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAccountRequest() (request *CreateAccountRequest) {
	request = &CreateAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "CreateAccount")
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

func NewUpdateAccountsLockStatusRequest() (request *UpdateAccountsLockStatusRequest) {
	request = &UpdateAccountsLockStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "UpdateAccountsLockStatus")
	return
}

func NewUpdateAccountsLockStatusResponse() (response *UpdateAccountsLockStatusResponse) {
	response = &UpdateAccountsLockStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 锁定解锁租户账号
func (c *Client) UpdateAccountsLockStatus(request *UpdateAccountsLockStatusRequest) (response *UpdateAccountsLockStatusResponse, err error) {
	if request == nil {
		request = NewUpdateAccountsLockStatusRequest()
	}
	response = NewUpdateAccountsLockStatusResponse()
	err = c.Send(request, response)
	return
}

func NewBoundCustomCategoryRequest() (request *BoundCustomCategoryRequest) {
	request = &BoundCustomCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "BoundCustomCategory")
	return
}

func NewBoundCustomCategoryResponse() (response *BoundCustomCategoryResponse) {
	response = &BoundCustomCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客户类别批量绑定租户账号
func (c *Client) BoundCustomCategory(request *BoundCustomCategoryRequest) (response *BoundCustomCategoryResponse, err error) {
	if request == nil {
		request = NewBoundCustomCategoryRequest()
	}
	response = NewBoundCustomCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteStrategyWhitelistRequest() (request *DeleteStrategyWhitelistRequest) {
	request = &DeleteStrategyWhitelistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "DeleteStrategyWhitelist")
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

func NewResetCustomerPasswordRequest() (request *ResetCustomerPasswordRequest) {
	request = &ResetCustomerPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bspaccount", APIVersion, "ResetCustomerPassword")
	return
}

func NewResetCustomerPasswordResponse() (response *ResetCustomerPasswordResponse) {
	response = &ResetCustomerPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置租户主账号密码
func (c *Client) ResetCustomerPassword(request *ResetCustomerPasswordRequest) (response *ResetCustomerPasswordResponse, err error) {
	if request == nil {
		request = NewResetCustomerPasswordRequest()
	}
	response = NewResetCustomerPasswordResponse()
	err = c.Send(request, response)
	return
}
