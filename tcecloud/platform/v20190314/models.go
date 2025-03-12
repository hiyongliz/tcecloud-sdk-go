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
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type CreateAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomSubAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户列表

		SubAccountList []*AccountList `json:"SubAccountList,omitempty" name:"SubAccountList"`
		// 用户数量

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCustomSubAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomSubAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryStrategyParentInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryStrategyParentInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryStrategyParentInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddIdpConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddIdpConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddIdpConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomSubAccountRequest struct {
	*tchttp.BaseRequest

	// 主账号uin

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 是否获取加密数据

	GetSecretField *bool `json:"GetSecretField,omitempty" name:"GetSecretField"`
}

func (r *GetCustomSubAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomSubAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据库标识

		UID *string `json:"UID,omitempty" name:"UID"`
		// 数据库密码

		Password *string `json:"Password,omitempty" name:"Password"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPasswordsRequest struct {
	*tchttp.BaseRequest

	// 查询条件

	Filter []*Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *QueryPasswordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPasswordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountList struct {
	// 是否激活

	Active *int64 `json:"Active,omitempty" name:"Active"`
	// 新增日期

	AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`
	// APPID

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 绑定状态

	BindStatus *int64 `json:"BindStatus,omitempty" name:"BindStatus"`
	// 上次登录时间

	LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`
	// 邮件

	Mail *string `json:"Mail,omitempty" name:"Mail"`
	// 修改日期

	ModTimestamp *string `json:"ModTimestamp,omitempty" name:"ModTimestamp"`
	// 昵称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 输主uin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 手机号码

	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
	// 标记

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 子账号数量

	SubCount *int64 `json:"SubCount,omitempty" name:"SubCount"`
	// uid

	Uid *int64 `json:"Uid,omitempty" name:"Uid"`
	// uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// uintype

	UinType *int64 `json:"UinType,omitempty" name:"UinType"`
	// username

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 秘钥

	IdKeys []*IdKeys `json:"IdKeys,omitempty" name:"IdKeys"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

type GetIdpConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetIdpConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetIdpConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PasswordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PasswordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PasswordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetPasswordTypesRequest struct {
	*tchttp.BaseRequest

	// 设备信息列表

	Servers []*ServerDevice `json:"Servers,omitempty" name:"Servers"`
	// 是否自动生成密码

	AutoPwd *uint64 `json:"AutoPwd,omitempty" name:"AutoPwd"`
	// 密码类型

	PasswordType *uint64 `json:"PasswordType,omitempty" name:"PasswordType"`
}

func (r *SetPasswordTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetPasswordTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IdKeys struct {
	// createTime

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 秘钥id

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// 秘钥key

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
	// source

	Source *int64 `json:"Source,omitempty" name:"Source"`
	// 启用状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type UpdateOauthConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOauthConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOauthConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetPasswordTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetPasswordTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetPasswordTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryStrategyWhitelistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryStrategyWhitelistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryStrategyWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateIdpConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateIdpConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateIdpConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccountParams struct {
	// 主账号uin

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// mail

	Mail *string `json:"Mail,omitempty" name:"Mail"`
	// pwd

	Password *string `json:"Password,omitempty" name:"Password"`
	// 手机号

	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 需要重置密码

	NeedResetPassword *int64 `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`
	// IsVerifyCertify

	IsVerifyCertify *int64 `json:"IsVerifyCertify,omitempty" name:"IsVerifyCertify"`
}

type DeleteStrategyWhitelistRequest struct {
	*tchttp.BaseRequest

	// input

	Params *DeleteStrategyWhitelistParams `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteStrategyWhitelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteStrategyWhitelistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePasswordsRequest struct {
	*tchttp.BaseRequest

	// 设备ip列表，多个用;号分隔

	Ips *string `json:"Ips,omitempty" name:"Ips"`
}

func (r *DescribePasswordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePasswordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IdpConfig struct {
	// access token 地址

	AccessTokenUri *string `json:"AccessTokenUri,omitempty" name:"AccessTokenUri"`
	// 授权地址

	AuthorizeUri *string `json:"AuthorizeUri,omitempty" name:"AuthorizeUri"`
	// 秘钥 id

	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
	//  秘钥 key

	ClientSecret *string `json:"ClientSecret,omitempty" name:"ClientSecret"`
	// 创建时间

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建者 id

	CreatorUin *uint64 `json:"CreatorUin,omitempty" name:"CreatorUin"`
	// 自定义邮箱字段

	EmailField *string `json:"EmailField,omitempty" name:"EmailField"`
	// 获取账号地址

	GetUserInfoUri *string `json:"GetUserInfoUri,omitempty" name:"GetUserInfoUri"`
	// 配置 id

	IdpId *uint64 `json:"IdpId,omitempty" name:"IdpId"`
	// 配置名称

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// 是否开启认证配置

	IsEnabled *uint64 `json:"IsEnabled,omitempty" name:"IsEnabled"`
	// 是否自动同步账号信息

	IsSyncIdpUser *uint64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
	// 自定义登录字段

	LoginAccountField *string `json:"LoginAccountField,omitempty" name:"LoginAccountField"`
	// 自定义昵称字段

	NickNameField *string `json:"NickNameField,omitempty" name:"NickNameField"`
	// 自定义手机号码字段

	PhoneNumField *string `json:"PhoneNumField,omitempty" name:"PhoneNumField"`
	// 协议名称

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 更新时间

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type QueryCustomAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户数量

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 用户列表

		AccountList []*AccountList `json:"AccountList,omitempty" name:"AccountList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCustomAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCustomAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateIdpConfigRequest struct {
	*tchttp.BaseRequest

	// 主键id

	IdpId *int64 `json:"IdpId,omitempty" name:"IdpId"`
	// idp名称

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// idp登录uri

	IdpCasLoginUri *string `json:"IdpCasLoginUri,omitempty" name:"IdpCasLoginUri"`
	// idp退出uri

	IdpCasLogoutUri *string `json:"IdpCasLogoutUri,omitempty" name:"IdpCasLogoutUri"`
	// idp检验uri

	IdpCasValidateUri *string `json:"IdpCasValidateUri,omitempty" name:"IdpCasValidateUri"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是否同步 idp 用户

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
	// idp 用户自定义账号字段

	LoginAccountField *string `json:"LoginAccountField,omitempty" name:"LoginAccountField"`
	// idp 用户自定义昵称字段

	NickNameField *string `json:"NickNameField,omitempty" name:"NickNameField"`
	// idp 用户自定义邮箱字段

	EmailField *string `json:"EmailField,omitempty" name:"EmailField"`
	// idp 用户自定义手机字段

	PhoneNumField *string `json:"PhoneNumField,omitempty" name:"PhoneNumField"`
	// idp 用户自定义区域号码字段

	CountryCodeField *string `json:"CountryCodeField,omitempty" name:"CountryCodeField"`
}

func (r *UpdateIdpConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateIdpConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddIdpConfigRequest struct {
	*tchttp.BaseRequest

	// idp名称

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// idp登录url

	IdpCasLoginUri *string `json:"IdpCasLoginUri,omitempty" name:"IdpCasLoginUri"`
	// idp退出uri

	IdpCasLogoutUri *string `json:"IdpCasLogoutUri,omitempty" name:"IdpCasLogoutUri"`
	// idp检验uri

	IdpCasValidateUri *string `json:"IdpCasValidateUri,omitempty" name:"IdpCasValidateUri"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是否同步 idp user

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
	// idp用户映射登录账号字段

	LoginAccountField *string `json:"LoginAccountField,omitempty" name:"LoginAccountField"`
	// idp用户映射昵称字段

	NickNameField *string `json:"NickNameField,omitempty" name:"NickNameField"`
	// idp用户映射邮箱字段

	EmailField *string `json:"EmailField,omitempty" name:"EmailField"`
	// idp用户映射手机字段

	PhoneNumField *string `json:"PhoneNumField,omitempty" name:"PhoneNumField"`
	// idp用户映射区域字段

	CountryCodeField *string `json:"CountryCodeField,omitempty" name:"CountryCodeField"`
}

func (r *AddIdpConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddIdpConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordRequest struct {
	*tchttp.BaseRequest

	// 设备名称和密码

	Filter []*Filter `json:"Filter,omitempty" name:"Filter"`
	// 自动生成密码

	AutoPwd *uint64 `json:"AutoPwd,omitempty" name:"AutoPwd"`
}

func (r *ModifyPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyStrategyWhitelistRequest struct {
	*tchttp.BaseRequest

	// input

	Params *ModifyStrategyWhitelistParams `json:"Params,omitempty" name:"Params"`
}

func (r *ModifyStrategyWhitelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStrategyWhitelistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PasswordsRequest struct {
	*tchttp.BaseRequest

	// 查询条件

	Filter []*Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *PasswordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PasswordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendActiveMailRequest struct {
	*tchttp.BaseRequest

	// input

	Params *SendActiveMailParams `json:"Params,omitempty" name:"Params"`
}

func (r *SendActiveMailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendActiveMailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryStrategyWhitelistParams struct {
	// 白名单名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建者

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 白名单类型

	Status *string `json:"Status,omitempty" name:"Status"`
	// 白名单key

	Type *string `json:"Type,omitempty" name:"Type"`
	// 起始页

	Pageid *int64 `json:"Pageid,omitempty" name:"Pageid"`
	// 每页最大记录数

	Pagesize *int64 `json:"Pagesize,omitempty" name:"Pagesize"`
}

type GetCustomAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户列表

		AccountList []*AccountList `json:"AccountList,omitempty" name:"AccountList"`
		// 用户数量

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCustomAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCustomAccountRequest struct {
	*tchttp.BaseRequest

	// 租户端登录名

	Account *string `json:"Account,omitempty" name:"Account"`
	// 租户手机号码

	PhoneNumber *int64 `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
	// 租户昵称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 租户端appid

	CustomAppId *int64 `json:"CustomAppId,omitempty" name:"CustomAppId"`
	// 分页

	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`
	// 分页大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 账号uin

	CustomUin *int64 `json:"CustomUin,omitempty" name:"CustomUin"`
	// 排序key

	SortKey *string `json:"SortKey,omitempty" name:"SortKey"`
	// 顺序或者逆序，0：顺序，1：逆序

	SortTurn *int64 `json:"SortTurn,omitempty" name:"SortTurn"`
	// 开始时间

	StartLastLoginTime *int64 `json:"StartLastLoginTime,omitempty" name:"StartLastLoginTime"`
	// 结束时间

	EndLastLoginTime *int64 `json:"EndLastLoginTime,omitempty" name:"EndLastLoginTime"`
	// 是否获取

	GetCertificateField *bool `json:"GetCertificateField,omitempty" name:"GetCertificateField"`
	// 主体名称

	PrimaryName *string `json:"PrimaryName,omitempty" name:"PrimaryName"`
}

func (r *QueryCustomAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCustomAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetPasswordRequest struct {
	*tchttp.BaseRequest

	// ip列表，多个ip用;号分隔

	Ips *string `json:"Ips,omitempty" name:"Ips"`
}

func (r *ResetPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateStrategyWhitelistParams struct {
	// 创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 分组ID

	Parentid *int64 `json:"Parentid,omitempty" name:"Parentid"`
	// 白名单key

	Type *string `json:"Type,omitempty" name:"Type"`
	// 白名单名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 白名单类型

	Status *string `json:"Status,omitempty" name:"Status"`
	// 白名单列表

	Whitelist *string `json:"Whitelist,omitempty" name:"Whitelist"`
	// 备注说明

	Info *string `json:"Info,omitempty" name:"Info"`
}

type AddOauthConfigRequest struct {
	*tchttp.BaseRequest

	// idp名称

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// 协议名称

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 注册应用的id

	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
	// 注册应用的密钥

	ClientSecret *string `json:"ClientSecret,omitempty" name:"ClientSecret"`
	// oauth验证授权信息url

	AuthorizeUri *string `json:"AuthorizeUri,omitempty" name:"AuthorizeUri"`
	// 获取access_token url

	AccessTokenUri *string `json:"AccessTokenUri,omitempty" name:"AccessTokenUri"`
	// 获取用户信息url

	GetUserInfoUri *string `json:"GetUserInfoUri,omitempty" name:"GetUserInfoUri"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 登录账号映射字段

	LoginAccountField *string `json:"LoginAccountField,omitempty" name:"LoginAccountField"`
	// 昵称映射字段

	NickNameField *string `json:"NickNameField,omitempty" name:"NickNameField"`
	// 邮箱映射字段

	EmailField *string `json:"EmailField,omitempty" name:"EmailField"`
	// 手机号映射字段

	PhoneNumField *string `json:"PhoneNumField,omitempty" name:"PhoneNumField"`
	// 是否同步 idp 账号

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
	// LogoutUrl

	LogoutUrl *string `json:"LogoutUrl,omitempty" name:"LogoutUrl"`
}

func (r *AddOauthConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOauthConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateStrategyWhitelistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateStrategyWhitelistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateStrategyWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisabledIdpConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisabledIdpConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisabledIdpConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryStrategyWhitelistRequest struct {
	*tchttp.BaseRequest

	// input

	Params *QueryStrategyWhitelistParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryStrategyWhitelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryStrategyWhitelistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendActiveMailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendActiveMailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendActiveMailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnabledIdpConfigRequest struct {
	*tchttp.BaseRequest

	// 主键id

	IdpId *int64 `json:"IdpId,omitempty" name:"IdpId"`
	// 用户登录名

	TuserId *string `json:"TuserId,omitempty" name:"TuserId"`
}

func (r *EnabledIdpConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnabledIdpConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPasswordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryPasswordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPasswordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteStrategyWhitelistParams struct {
	// 白名单key

	Type *string `json:"Type,omitempty" name:"Type"`
	// 白名单名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type ModifyStrategyWhitelistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyStrategyWhitelistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStrategyWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyStrategyWhitelistParams struct {
	// 白名单ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 修改后的白名单分组ID

	Parentid *int64 `json:"Parentid,omitempty" name:"Parentid"`
	// 修改后的白名单key

	Type *string `json:"Type,omitempty" name:"Type"`
	// 修改后的白名单名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 修改后的白名单类型

	Status *string `json:"Status,omitempty" name:"Status"`
	// 修改后的白名单列表

	Whitelist *string `json:"Whitelist,omitempty" name:"Whitelist"`
	// 修改后的备注说明

	Info *string `json:"Info,omitempty" name:"Info"`
	// 修改后的创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
}

type SendActiveMailParams struct {
	// email

	Mail *string `json:"Mail,omitempty" name:"Mail"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest

	// input

	Params []*CreateAccountParams `json:"Params,omitempty" name:"Params"`
}

func (r *CreateAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnabledIdpConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnabledIdpConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnabledIdpConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBPasswordRequest struct {
	*tchttp.BaseRequest

	// 数据库标识

	UID *string `json:"UID,omitempty" name:"UID"`
}

func (r *DescribeDBPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomAccountRequest struct {
	*tchttp.BaseRequest

	// 分页

	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`
	// 页面数量

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 是否获取加密key数据

	GetSecretField *bool `json:"GetSecretField,omitempty" name:"GetSecretField"`
	// 排序key

	SortKey *string `json:"SortKey,omitempty" name:"SortKey"`
	// 排序顺序

	SortTurn *int64 `json:"SortTurn,omitempty" name:"SortTurn"`
	// 是否获取

	GetCertificateField *bool `json:"GetCertificateField,omitempty" name:"GetCertificateField"`
}

func (r *GetCustomAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateStrategyWhitelistRequest struct {
	*tchttp.BaseRequest

	// 无

	Params *CreateStrategyWhitelistParams `json:"Params,omitempty" name:"Params"`
}

func (r *CreateStrategyWhitelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateStrategyWhitelistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerDevice struct {
	// 设备id

	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 设备ip

	IntranetIp *string `json:"IntranetIp,omitempty" name:"IntranetIp"`
}

type UpdateOauthConfigRequest struct {
	*tchttp.BaseRequest

	// idp名称

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 注册应用的id

	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
	// 注册应用的密钥

	ClientSecret *string `json:"ClientSecret,omitempty" name:"ClientSecret"`
	// oauth验证授权信息url

	AuthorizeUri *string `json:"AuthorizeUri,omitempty" name:"AuthorizeUri"`
	// 获取access_token url

	AccessTokenUri *string `json:"AccessTokenUri,omitempty" name:"AccessTokenUri"`
	// 获取用户信息url

	GetUserInfoUri *string `json:"GetUserInfoUri,omitempty" name:"GetUserInfoUri"`
	// idp主键id

	IdpId *uint64 `json:"IdpId,omitempty" name:"IdpId"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 登录账号映射字段

	LoginAccountField *string `json:"LoginAccountField,omitempty" name:"LoginAccountField"`
	// 昵称映射字段

	NickNameField *string `json:"NickNameField,omitempty" name:"NickNameField"`
	// 邮箱映射字段

	EmailField *string `json:"EmailField,omitempty" name:"EmailField"`
	// 手机号映射字段

	PhoneNumField *string `json:"PhoneNumField,omitempty" name:"PhoneNumField"`
	// 是否同步 idp 用户

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
	// LogoutUrl

	LogoutUrl *string `json:"LogoutUrl,omitempty" name:"LogoutUrl"`
}

func (r *UpdateOauthConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOauthConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteStrategyWhitelistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteStrategyWhitelistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteStrategyWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePasswordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePasswordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePasswordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetIdpConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetIdpConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetIdpConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryStrategyParentInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryStrategyParentInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryStrategyParentInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOauthConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// idp 配置

		IdpConfig *IdpConfig `json:"IdpConfig,omitempty" name:"IdpConfig"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddOauthConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOauthConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisabledIdpConfigRequest struct {
	*tchttp.BaseRequest

	// 主键id

	IdpId *int64 `json:"IdpId,omitempty" name:"IdpId"`
}

func (r *DisabledIdpConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisabledIdpConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}
