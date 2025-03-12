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

type EnableCustomerApiKeyRequest struct {
	*tchttp.BaseRequest

	// 租户账号OwnerUin，取值为主账号Uin

	CustomerOwnerUin *int64 `json:"CustomerOwnerUin,omitempty" name:"CustomerOwnerUin"`
	// 租户账号Uin，主账号时取值为OwnerUin，子账号时取值为子账号Uin

	CustomerUin *int64 `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 租户账号appid

	CustomerAppId *int64 `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
	// 租户持久秘钥 secreId

	CustomerSecretId *string `json:"CustomerSecretId,omitempty" name:"CustomerSecretId"`
}

func (r *EnableCustomerApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableCustomerApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAccountBaseInfoRequest struct {
	*tchttp.BaseRequest

	// 查询起始页

	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`
	// 每页记录数

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 根据名称查询

	Account *string `json:"Account,omitempty" name:"Account"`
	// Uin列表过滤查询

	Uins []*int64 `json:"Uins,omitempty" name:"Uins"`
}

func (r *GetAccountBaseInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAccountBaseInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetCustomerPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetCustomerPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetCustomerPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCustomAccountsIdentifyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCustomAccountsIdentifyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCustomAccountsIdentifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCustomerApiKeyIdKeys struct {
	// CreateTime

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 秘钥id

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// source

	Source *int64 `json:"Source,omitempty" name:"Source"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type BoundCustomCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BoundCustomCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BoundCustomCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomAccountAttributesRequest struct {
	*tchttp.BaseRequest
}

func (r *GetCustomAccountAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomAccountAttributesRequest) FromJsonString(s string) error {
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

type UpdateCustomSubAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCustomSubAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCustomSubAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAccountsLockStatusParams struct {
	// 锁定解锁的Uin列表

	Uins []*int64 `json:"Uins,omitempty" name:"Uins"`
	// 锁定解锁状态，0：未锁定；1：临时锁定；2：持久锁定

	LockedStatus *int64 `json:"LockedStatus,omitempty" name:"LockedStatus"`
}

type UpdateCustomAccountAttributeValueRequest struct {
	*tchttp.BaseRequest

	// id

	ValueId *uint64 `json:"ValueId,omitempty" name:"ValueId"`
	// 属性值

	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
}

func (r *UpdateCustomAccountAttributeValueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCustomAccountAttributeValueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type DeleteCustomCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCustomCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableCustomerApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableCustomerApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableCustomerApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type IpLimitParams struct {
	// 是否开启ip限制

	IsOpenLimit *uint64 `json:"IsOpenLimit,omitempty" name:"IsOpenLimit"`
	// 限制类型，0:黑名单，1：白名单，两者互斥

	LimitType *uint64 `json:"LimitType,omitempty" name:"LimitType"`
	// 限制ip列表，最多为10个

	IpList []*string `json:"IpList,omitempty" name:"IpList"`
	// 客户主账号uin

	CustomUin *uint64 `json:"CustomUin,omitempty" name:"CustomUin"`
	// 客户子账号uin

	SubCustomUin *uint64 `json:"SubCustomUin,omitempty" name:"SubCustomUin"`
}

type SendActiveMailParams struct {
	// email

	Mail *string `json:"Mail,omitempty" name:"Mail"`
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

type SendLinkPCloudAccountNoticeEmailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendLinkPCloudAccountNoticeEmailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendLinkPCloudAccountNoticeEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCustomOwnerAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCustomOwnerAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCustomOwnerAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckCustomAccountRequest struct {
	*tchttp.BaseRequest

	// 登录账号数组

	AccountList []*string `json:"AccountList,omitempty" name:"AccountList"`
}

func (r *CheckCustomAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckCustomAccountRequest) FromJsonString(s string) error {
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
	// 是否过滤已激活账号

	IsActive *bool `json:"IsActive,omitempty" name:"IsActive"`
}

func (r *GetCustomAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomerApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 持久密钥列表

		IdKeys []*IdKeys `json:"IdKeys,omitempty" name:"IdKeys"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCustomerApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomerApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateResult struct {
	// 应用id，创建子账号时为0（创建成功时有此值）

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 账号唯一uin信息（创建成功时有此值）

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 失败错误码(创建失败，fails里面才会有此值)

	ReturnValue *int64 `json:"ReturnValue,omitempty" name:"ReturnValue"`
	// 创建失败错误信息创建失败，fails里面才会有此值)

	ReturnMessage *string `json:"ReturnMessage,omitempty" name:"ReturnMessage"`
}

type ServerDevice struct {
	// 设备id

	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 设备ip

	IntranetIp *string `json:"IntranetIp,omitempty" name:"IntranetIp"`
}

type BoundCustomCategoryRequest struct {
	*tchttp.BaseRequest

	// 类别id

	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 类别键

	CategoryKey *string `json:"CategoryKey,omitempty" name:"CategoryKey"`
	// 类别值

	CategoryValue *string `json:"CategoryValue,omitempty" name:"CategoryValue"`
	// 租户账号uin列表

	CustomUinList []*uint64 `json:"CustomUinList,omitempty" name:"CustomUinList"`
}

func (r *BoundCustomCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BoundCustomCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type GetFullCustomCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分类列表

		CategoryList []*CategoryAttr `json:"CategoryList,omitempty" name:"CategoryList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFullCustomCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFullCustomCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryStrategyParentInfoRequest struct {
	*tchttp.BaseRequest

	// 根据指定分组Type查询白名单分组信息，当前支持的取值：CVM/CBS/COS/REGION/other

	Types []*string `json:"Types,omitempty" name:"Types"`
}

func (r *QueryStrategyParentInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryStrategyParentInfoRequest) FromJsonString(s string) error {
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

type GetCustomerUrgentCodeRequest struct {
	*tchttp.BaseRequest

	// 租户主账号uin

	CustomerOwnerUin *string `json:"CustomerOwnerUin,omitempty" name:"CustomerOwnerUin"`
	// 当前操作账号uin，暂只支持主账号uin

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
}

func (r *GetCustomerUrgentCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomerUrgentCodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCustomCategoryRequest struct {
	*tchttp.BaseRequest

	// 页码

	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`
	// 每页显示条数

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 类别键数组

	CategoryKey []*string `json:"CategoryKey,omitempty" name:"CategoryKey"`
}

func (r *QueryCustomCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCustomCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCustomOwnerAccountRequest struct {
	*tchttp.BaseRequest

	// 租户主账号uin

	CustomerOwnerUin *uint64 `json:"CustomerOwnerUin,omitempty" name:"CustomerOwnerUin"`
	// 主账号更新的详细参数

	Details *OwnerAccountInfo `json:"Details,omitempty" name:"Details"`
}

func (r *UpdateCustomOwnerAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCustomOwnerAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Strategy struct {
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 白名单key

	Type *string `json:"Type,omitempty" name:"Type"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分组id

	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 创建者

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 类型状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 附加信息

	Info *string `json:"Info,omitempty" name:"Info"`
	// 名单列表

	Whitelist []*string `json:"Whitelist,omitempty" name:"Whitelist"`
}

type GetCustomAccountsByUinListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账号信息列表

		AccountList []*AppAccountAttr `json:"AccountList,omitempty" name:"AccountList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCustomAccountsByUinListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomAccountsByUinListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustomerApiKeyRequest struct {
	*tchttp.BaseRequest

	// 客户账号uin

	CustomerUin *uint64 `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 客户主账号uin

	CustomerOwnerUin *uint64 `json:"CustomerOwnerUin,omitempty" name:"CustomerOwnerUin"`
	// 客户账号对应appId

	CustomerAppId *uint64 `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
}

func (r *CreateCustomerApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomerApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttributeAndValues struct {
	// 属性key

	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`
	// 属性值

	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
	// 属性名称

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 属性值id

	ValueId *int64 `json:"ValueId,omitempty" name:"ValueId"`
	// 属性值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 关联的用户uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
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
	// 白名单所在的分组id

	Parentid *int64 `json:"Parentid,omitempty" name:"Parentid"`
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

type ResetCustomerPasswordRequest struct {
	*tchttp.BaseRequest

	// 租户主账号uin

	CustomerOwnerUin *string `json:"CustomerOwnerUin,omitempty" name:"CustomerOwnerUin"`
	// 租户Uin

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ResetCustomerPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetCustomerPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomAccountsByAppIdsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账号列表

		AccountList []*AppAccountAttr `json:"AccountList,omitempty" name:"AccountList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCustomAccountsByAppIdsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomAccountsByAppIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SeedOIDCLoginTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 临时token

		Token *string `json:"Token,omitempty" name:"Token"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SeedOIDCLoginTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SeedOIDCLoginTokenResponse) FromJsonString(s string) error {
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

	Whitelist []*string `json:"Whitelist,omitempty" name:"Whitelist"`
	// 修改后的备注说明

	Info *string `json:"Info,omitempty" name:"Info"`
	// 修改后的创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
}

type AddCustomAccountAttributeValueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddCustomAccountAttributeValueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddCustomAccountAttributeValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomCategoryRequest struct {
	*tchttp.BaseRequest

	// 租户账号Uin

	CustomUin *uint64 `json:"CustomUin,omitempty" name:"CustomUin"`
	// 类别信息列表

	CategoryParams []*CustomCategoryParams `json:"CategoryParams,omitempty" name:"CategoryParams"`
}

func (r *ModifyCustomCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomSubAccountRequest struct {
	*tchttp.BaseRequest

	// 主账号uin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
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

type OwnerAccountInfo struct {
	// 主账号备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
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

type InsertOrUpdateCustomAccountAttributeValuesRequest struct {
	*tchttp.BaseRequest

	// Uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 扩展属性

	AttributeAndValues []*AttributeAndValues `json:"AttributeAndValues,omitempty" name:"AttributeAndValues"`
}

func (r *InsertOrUpdateCustomAccountAttributeValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InsertOrUpdateCustomAccountAttributeValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomIpLimitRequest struct {
	*tchttp.BaseRequest

	// ip白名单黑名单限制参数

	Params []*IpLimitParams `json:"Params,omitempty" name:"Params"`
}

func (r *ModifyCustomIpLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomIpLimitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryStrategyWhitelistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 白名单列表

		StrategyList []*Strategy `json:"StrategyList,omitempty" name:"StrategyList"`
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

type CreateAccountRespParam struct {
	// 安全邮箱

	Mail *string `json:"Mail,omitempty" name:"Mail"`
	// 手机号码

	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
	// 国家代码

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 联系邮箱

	ContactMail *string `json:"ContactMail,omitempty" name:"ContactMail"`
	// 身份类型

	IdentifyType *int64 `json:"IdentifyType,omitempty" name:"IdentifyType"`
	// 登录账号

	LoginAccount *string `json:"LoginAccount,omitempty" name:"LoginAccount"`
	// 是否需要重置密码,0：不需要，1：需要

	NeedResetPassword *int64 `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`
	// 主账号uIn（创建子账号时才会返回）

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 创建结果

	Result *CreateResult `json:"Result,omitempty" name:"Result"`
	// 是否进行资质审核认证

	IsVerifyCertify *int64 `json:"IsVerifyCertify,omitempty" name:"IsVerifyCertify"`
	// 创建账号时返回的持久密钥信息

	IdKey *IdKey `json:"IdKey,omitempty" name:"IdKey"`
}

type StrategyParent struct {
	// 分组id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 分组key

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DisableCustomerApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableCustomerApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableCustomerApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomAccountsByUinListRequest struct {
	*tchttp.BaseRequest

	// uin列表

	UinList []*uint64 `json:"UinList,omitempty" name:"UinList"`
}

func (r *GetCustomAccountsByUinListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomAccountsByUinListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomIpLimitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomIpLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomIpLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryStrategyParentInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 分组列表

		StrategyParentList []*StrategyParent `json:"StrategyParentList,omitempty" name:"StrategyParentList"`
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

type IdKey struct {
	// 创建失败时的错误信息

	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
	// 持久密钥id

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// 持久密钥key

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
	// 创建来源

	Source *int64 `json:"Source,omitempty" name:"Source"`
	// 密钥状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
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

type UpdateCustomSubAccountRequest struct {
	*tchttp.BaseRequest

	// 租户主账号Uin

	CustomerOwnerUin *uint64 `json:"CustomerOwnerUin,omitempty" name:"CustomerOwnerUin"`
	// 租户子账号用户名(登录账号名)

	CustomUserName *string `json:"CustomUserName,omitempty" name:"CustomUserName"`
	// 租户子账号备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *UpdateCustomSubAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCustomSubAccountRequest) FromJsonString(s string) error {
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

	Whitelist []*string `json:"Whitelist,omitempty" name:"Whitelist"`
	// 备注说明

	Info *string `json:"Info,omitempty" name:"Info"`
}

type QueryCustomerApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 租户秘钥列表

		IdKeys []*QueryCustomerApiKeyIdKeys `json:"IdKeys,omitempty" name:"IdKeys"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCustomerApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCustomerApiKeyResponse) FromJsonString(s string) error {
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

	UinType *string `json:"UinType,omitempty" name:"UinType"`
	// username

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 秘钥

	IdKeys []*IdKeys `json:"IdKeys,omitempty" name:"IdKeys"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 账号来源类型

	SystemType *string `json:"SystemType,omitempty" name:"SystemType"`
	// 企业微信唯一ID

	QywxUserId *string `json:"QywxUserId,omitempty" name:"QywxUserId"`
	// 账号名

	Account *string `json:"Account,omitempty" name:"Account"`
	// 关联的省份信息

	AddrProvince *string `json:"AddrProvince,omitempty" name:"AddrProvince"`
	// 联系邮箱

	ContactEmail *string `json:"ContactEmail,omitempty" name:"ContactEmail"`
	// 个人用户名或企业名称

	CertificateNumber *string `json:"CertificateNumber,omitempty" name:"CertificateNumber"`
	// 登录状态,0:正常, 1:临时锁定,2:持久锁定

	LoginStatus *int64 `json:"LoginStatus,omitempty" name:"LoginStatus"`
	// 是否允许挂失mfa

	AllowReportLoss *int64 `json:"AllowReportLoss,omitempty" name:"AllowReportLoss"`
	// 认证类型，0：个人认证，1: 企业认证

	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`
	// 账号主体名称

	PrimaryName *string `json:"PrimaryName,omitempty" name:"PrimaryName"`
	// 所属城市

	AddrCity *string `json:"AddrCity,omitempty" name:"AddrCity"`
	// 详细地址信息

	AddrDetail *string `json:"AddrDetail,omitempty" name:"AddrDetail"`
	// 所在地

	Address *string `json:"Address,omitempty" name:"Address"`
	// 国家代码

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 身份类型

	IdentifyType *uint64 `json:"IdentifyType,omitempty" name:"IdentifyType"`
	// 关联的用户属性

	UserAttributeAndValues []*AttributeAndValues `json:"UserAttributeAndValues,omitempty" name:"UserAttributeAndValues"`
	// 关联的用户类别

	UserCategory []*CategoryAttr `json:"UserCategory,omitempty" name:"UserCategory"`
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

type CategoryAttr struct {
	// 类别ID

	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 类别key

	CategoryKey *string `json:"CategoryKey,omitempty" name:"CategoryKey"`
	// 类别value

	CategoryValue *string `json:"CategoryValue,omitempty" name:"CategoryValue"`
	// 绑定账号数量

	BindCustomNum *int64 `json:"BindCustomNum,omitempty" name:"BindCustomNum"`
	// 关联账号uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建结果

		Data *CreateRespData `json:"Data,omitempty" name:"Data"`
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

type QueryCustomerApiKeyRequest struct {
	*tchttp.BaseRequest

	// 租户账号OwnerUin，取值为主账号Uin

	CustomerOwnerUin *int64 `json:"CustomerOwnerUin,omitempty" name:"CustomerOwnerUin"`
	// 租户账号Uin，主账号时取值为OwnerUin，子账号时取值为子账号Uin

	CustomerUin *int64 `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 租户账号appid

	CustomerAppId *int64 `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
}

func (r *QueryCustomerApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCustomerApiKeyRequest) FromJsonString(s string) error {
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

type CreateRespData struct {
	// 失败账号列表

	Fail *CreateAccountRespParam `json:"Fail,omitempty" name:"Fail"`
	// 成功账号列表

	Success *CreateAccountRespParam `json:"Success,omitempty" name:"Success"`
}

type AddCustomAccountAttributeValueRequest struct {
	*tchttp.BaseRequest

	// 属性id

	AttributeId *int64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 属性值

	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
	// 租户端账号uin

	CustomUin *int64 `json:"CustomUin,omitempty" name:"CustomUin"`
}

func (r *AddCustomAccountAttributeValueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddCustomAccountAttributeValueRequest) FromJsonString(s string) error {
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
	// 属性key

	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`
	// 属性值

	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
	// 类别属性值

	Category *string `json:"Category,omitempty" name:"Category"`
	// 是否为主账号查询

	IsOwner *uint64 `json:"IsOwner,omitempty" name:"IsOwner"`
	// 是否过滤已激活账号

	IsActive *bool `json:"IsActive,omitempty" name:"IsActive"`
	// 指定需要包含的扩展字段

	IncludeAttr []*string `json:"IncludeAttr,omitempty" name:"IncludeAttr"`
}

func (r *QueryCustomAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCustomAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SeedOIDCLoginTokenRequest struct {
	*tchttp.BaseRequest

	// 客户子账号uin

	CustomerUin *int64 `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 客户主账号Uin

	CustomerOwnerUin *int64 `json:"CustomerOwnerUin,omitempty" name:"CustomerOwnerUin"`
}

func (r *SeedOIDCLoginTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SeedOIDCLoginTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustomCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCustomCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomCategoryRequest struct {
	*tchttp.BaseRequest

	// 客户类别Id数组

	CategoryInfo []*CustomCategoryParams `json:"CategoryInfo,omitempty" name:"CategoryInfo"`
}

func (r *DeleteCustomCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomIpLimitRequest struct {
	*tchttp.BaseRequest

	// 客户主账号uin

	CustomUin *uint64 `json:"CustomUin,omitempty" name:"CustomUin"`
}

func (r *GetCustomIpLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomIpLimitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAccountsLockStatusRequest struct {
	*tchttp.BaseRequest

	// 账户登录名

	Accounts []*string `json:"Accounts,omitempty" name:"Accounts"`
	// 锁定解锁状态，0：未锁定；1：临时锁定；2：持久锁定

	LoginStatus *int64 `json:"LoginStatus,omitempty" name:"LoginStatus"`
}

func (r *UpdateAccountsLockStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAccountsLockStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomCategoryParams struct {
	// 类别id

	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 类别键

	CategoryKey *string `json:"CategoryKey,omitempty" name:"CategoryKey"`
	// 类别值

	CategoryValue *string `json:"CategoryValue,omitempty" name:"CategoryValue"`
}

type GetPasswordRuleByCustomUinRequest struct {
	*tchttp.BaseRequest

	// 客户主账号uin

	CustomUin *uint64 `json:"CustomUin,omitempty" name:"CustomUin"`
}

func (r *GetPasswordRuleByCustomUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPasswordRuleByCustomUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCustomAccountAttributeValueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCustomAccountAttributeValueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCustomAccountAttributeValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableCustomerApiKeyRequest struct {
	*tchttp.BaseRequest

	// 租户账号OwnerUin，取值为主账号Uin

	CustomerOwnerUin *int64 `json:"CustomerOwnerUin,omitempty" name:"CustomerOwnerUin"`
	// 租户账号Uin，主账号时取值为OwnerUin，子账号时取值为子账号Uin

	CustomerUin *int64 `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 租户账号appid

	CustomerAppId *int64 `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
	// 租户持久秘钥 secreId

	CustomerSecretId *string `json:"CustomerSecretId,omitempty" name:"CustomerSecretId"`
}

func (r *DisableCustomerApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableCustomerApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomerSafeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCustomerSafeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomerSafeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAccountsLockStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAccountsLockStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAccountsLockStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type GetCustomAccountAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 属性列表

		Items []*Attribute `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCustomAccountAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomAccountAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InsertOrUpdateCustomAccountAttributeValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InsertOrUpdateCustomAccountAttributeValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InsertOrUpdateCustomAccountAttributeValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryStrategyWhitelistRequest struct {
	*tchttp.BaseRequest

	// 白名单名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建者

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 白名单状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 白名单key

	Type *string `json:"Type,omitempty" name:"Type"`
	// 白名单所在分组id

	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`
	// 起始页

	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`
	// 每页最大记录数

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 白名单记录UIN或APPID

	Whitelist *string `json:"Whitelist,omitempty" name:"Whitelist"`
}

func (r *QueryStrategyWhitelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryStrategyWhitelistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppAccountAttr struct {
	// 账号AppID

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 账号uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 账号名称

	Account *string `json:"Account,omitempty" name:"Account"`
	// 昵称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
}

type BaseAccountInfo struct {
	// 用户名称

	Account *string `json:"Account,omitempty" name:"Account"`
	// 用户uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
}

type QueryCustomCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 客户类别列表

		CategoryList []*CategoryAttr `json:"CategoryList,omitempty" name:"CategoryList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCustomCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCustomCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Attribute struct {
	// 属性

	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`
	// 属性id

	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 属性名称

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
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

type GetCustomAccountsByAppIdsRequest struct {
	*tchttp.BaseRequest

	// appid列表

	AppIds []*int64 `json:"AppIds,omitempty" name:"AppIds"`
}

func (r *GetCustomAccountsByAppIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomAccountsByAppIdsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomerApiKeyRequest struct {
	*tchttp.BaseRequest

	// 客户账号uin

	CustomerUin *uint64 `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 客户账号ownerUin

	CustomerOwnerUin *uint64 `json:"CustomerOwnerUin,omitempty" name:"CustomerOwnerUin"`
	// 客户账号对应appId

	CustomerAppId *uint64 `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
}

func (r *GetCustomerApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomerApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomIpLimitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCustomIpLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomIpLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFullCustomCategoryRequest struct {
	*tchttp.BaseRequest
}

func (r *GetFullCustomCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFullCustomCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteStrategyWhitelistParams struct {
	// 白名单key

	Type *string `json:"Type,omitempty" name:"Type"`
	// 白名单名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateCustomCategoryRequest struct {
	*tchttp.BaseRequest

	// 新建客户类别参数

	CategoryList []*CustomCategoryParams `json:"CategoryList,omitempty" name:"CategoryList"`
}

func (r *CreateCustomCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAccountBaseInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账号列表

		AccountList []*BaseAccountInfo `json:"AccountList,omitempty" name:"AccountList"`
		// 总数

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAccountBaseInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAccountBaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomerUrgentCodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCustomerUrgentCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomerUrgentCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPasswordRuleByCustomUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPasswordRuleByCustomUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPasswordRuleByCustomUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendLinkPCloudAccountNoticeEmailRequest struct {
	*tchttp.BaseRequest

	// Uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 公有云Uin

	PCloudUin *string `json:"PCloudUin,omitempty" name:"PCloudUin"`
}

func (r *SendLinkPCloudAccountNoticeEmailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendLinkPCloudAccountNoticeEmailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCustomAccountsIdentifyRequest struct {
	*tchttp.BaseRequest

	// uin列表

	Uins []*uint64 `json:"Uins,omitempty" name:"Uins"`
	// 身份类型

	IdentifyType *int64 `json:"IdentifyType,omitempty" name:"IdentifyType"`
}

func (r *UpdateCustomAccountsIdentifyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCustomAccountsIdentifyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomerSafeConfigRequest struct {
	*tchttp.BaseRequest

	// 租户主账号

	CustomerOwnerUin *string `json:"CustomerOwnerUin,omitempty" name:"CustomerOwnerUin"`
	// 当前操作账号uin

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
}

func (r *GetCustomerSafeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomerSafeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustomerApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 密钥列表

		IdKeys []*IdKeys `json:"IdKeys,omitempty" name:"IdKeys"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCustomerApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomerApiKeyResponse) FromJsonString(s string) error {
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
	// 手机号对应国家码

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 属性列表

	AttributeAndValues []*AttributeAndValues `json:"AttributeAndValues,omitempty" name:"AttributeAndValues"`
	// 账号身份类型

	IdentifyType *int64 `json:"IdentifyType,omitempty" name:"IdentifyType"`
	// 登录账号

	LoginAccount *string `json:"LoginAccount,omitempty" name:"LoginAccount"`
	// 联系邮箱

	ContactEmail *string `json:"ContactEmail,omitempty" name:"ContactEmail"`
	// 需要发送新建成功提醒邮件

	NeedSendActivedNoticeMail *int64 `json:"NeedSendActivedNoticeMail,omitempty" name:"NeedSendActivedNoticeMail"`
	// 是否需要创建ak,sk，true-需要, false-不需要

	IsNeedApiKey *bool `json:"IsNeedApiKey,omitempty" name:"IsNeedApiKey"`
}

type CheckCustomAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckCustomAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckCustomAccountResponse) FromJsonString(s string) error {
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

	Status *int64 `json:"Status,omitempty" name:"Status"`
}
