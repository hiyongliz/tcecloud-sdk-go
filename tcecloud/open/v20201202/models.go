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
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type GetLdapIdpConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLdapIdpConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLdapIdpConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWorkWeixinConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateWorkWeixinConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateWorkWeixinConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWorkWeixinOpenAppConfigRequest struct {
	*tchttp.BaseRequest

	// 内部应用名

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 内部应用id

	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
	// 内部应用secret

	Secret *string `json:"Secret,omitempty" name:"Secret"`
	// 企业id

	CorpId *string `json:"CorpId,omitempty" name:"CorpId"`
}

func (r *CreateWorkWeixinOpenAppConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWorkWeixinOpenAppConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWorkWeixinOpenAppConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWorkWeixinOpenAppConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWorkWeixinOpenAppConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkWeixinAuthorizationScopeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetWorkWeixinAuthorizationScopeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinAuthorizationScopeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWorkWeixinOpenAppConfigRequest struct {
	*tchttp.BaseRequest

	// 内部应用名

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 内部应用id

	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
	// 内部应用secret

	Secret *string `json:"Secret,omitempty" name:"Secret"`
	// 企业id

	CorpId *string `json:"CorpId,omitempty" name:"CorpId"`
	// 配置id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *UpdateWorkWeixinOpenAppConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateWorkWeixinOpenAppConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLdapIdpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateLdapIdpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLdapIdpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLdapIdpConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetLdapIdpConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLdapIdpConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListIdentityProviderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListIdentityProviderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListIdentityProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWorkWeixinOpenAppConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateWorkWeixinOpenAppConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateWorkWeixinOpenAppConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkWeixinOpenAppMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 企业微信部门成员信息

		DepartmentList []*DepartmentAttr `json:"DepartmentList,omitempty" name:"DepartmentList"`
		// 企业微信可见成员信息

		VisibleUserList []*WechatUserAttr `json:"VisibleUserList,omitempty" name:"VisibleUserList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetWorkWeixinOpenAppMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinOpenAppMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TestLdapResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TestLdapResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestLdapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindWorkWeixinAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindWorkWeixinAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindWorkWeixinAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TestLdapRequest struct {
	*tchttp.BaseRequest

	// 认证类型

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// idp名称

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// ldap类型

	LdapType *string `json:"LdapType,omitempty" name:"LdapType"`
	// ldap url

	LdapUrl *string `json:"LdapUrl,omitempty" name:"LdapUrl"`
	// 连接类型

	ConnectType *string `json:"ConnectType,omitempty" name:"ConnectType"`
	// 证书

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// base dn

	BaseDn *string `json:"BaseDn,omitempty" name:"BaseDn"`
	// 管理账号

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// 管理账号密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 过滤条件

	FilterCondition *string `json:"FilterCondition,omitempty" name:"FilterCondition"`
	// 用户账号字段

	UserAccountField *string `json:"UserAccountField,omitempty" name:"UserAccountField"`
	// 用户邮箱字段

	UserMailField *string `json:"UserMailField,omitempty" name:"UserMailField"`
	// 用户昵称字段

	UserNicknameField *string `json:"UserNicknameField,omitempty" name:"UserNicknameField"`
	// 用户电话字段

	UserPhoneField *string `json:"UserPhoneField,omitempty" name:"UserPhoneField"`
	// 测试账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 测试账号面膜

	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *TestLdapRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestLdapRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkWeixinOpenAppConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetWorkWeixinOpenAppConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinOpenAppConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Items struct {
	// Uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// OwnerUin

	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// WechatUserId

	WechatUserId *string `json:"WechatUserId,omitempty" name:"WechatUserId"`
}

type CreateLdapIdpRequest struct {
	*tchttp.BaseRequest

	// 认证类型

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// idp名称

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// ldap类型

	LdapType *string `json:"LdapType,omitempty" name:"LdapType"`
	// ldap url

	LdapUrl *string `json:"LdapUrl,omitempty" name:"LdapUrl"`
	// 连接类型

	ConnectType *string `json:"ConnectType,omitempty" name:"ConnectType"`
	// 证书

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// base dn

	BaseDn *string `json:"BaseDn,omitempty" name:"BaseDn"`
	// 管理账号

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// 管理账号密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 过滤条件

	FilterCondition *string `json:"FilterCondition,omitempty" name:"FilterCondition"`
	// 用户账号字段

	UserAccountField *string `json:"UserAccountField,omitempty" name:"UserAccountField"`
	// 用户邮箱字段

	UserMailField *string `json:"UserMailField,omitempty" name:"UserMailField"`
	// 用户昵称字段

	UserNicknameField *string `json:"UserNicknameField,omitempty" name:"UserNicknameField"`
	// 用户电话字段

	UserPhoneField *string `json:"UserPhoneField,omitempty" name:"UserPhoneField"`
	// 是否同步 idp 用户

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
}

func (r *CreateLdapIdpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLdapIdpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkWeixinConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetWorkWeixinConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkWeixinConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetWorkWeixinConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLdapIdpRequest struct {
	*tchttp.BaseRequest

	// 认证类型

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// idp名称

	IdpName *string `json:"IdpName,omitempty" name:"IdpName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// ldap类型

	LdapType *string `json:"LdapType,omitempty" name:"LdapType"`
	// ldap url

	LdapUrl *string `json:"LdapUrl,omitempty" name:"LdapUrl"`
	// 连接类型

	ConnectType *string `json:"ConnectType,omitempty" name:"ConnectType"`
	// 证书

	Cert *string `json:"Cert,omitempty" name:"Cert"`
	// base dn

	BaseDn *string `json:"BaseDn,omitempty" name:"BaseDn"`
	// 管理账号

	AdminAccount *string `json:"AdminAccount,omitempty" name:"AdminAccount"`
	// 管理账号密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 过滤条件

	FilterCondition *string `json:"FilterCondition,omitempty" name:"FilterCondition"`
	// 用户账号字段

	UserAccountField *string `json:"UserAccountField,omitempty" name:"UserAccountField"`
	// 用户邮箱字段

	UserMailField *string `json:"UserMailField,omitempty" name:"UserMailField"`
	// 用户昵称字段

	UserNicknameField *string `json:"UserNicknameField,omitempty" name:"UserNicknameField"`
	// 用户电话字段

	UserPhoneField *string `json:"UserPhoneField,omitempty" name:"UserPhoneField"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否同步 idp 用户

	IsSyncIdpUser *int64 `json:"IsSyncIdpUser,omitempty" name:"IsSyncIdpUser"`
}

func (r *UpdateLdapIdpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLdapIdpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DepartmentAttr struct {
	// 部门id

	DepartmentId *uint64 `json:"DepartmentId,omitempty" name:"DepartmentId"`
	// 部门名称

	DepartmentName *string `json:"DepartmentName,omitempty" name:"DepartmentName"`
	// 父部门id

	ParentDepartmentId *uint64 `json:"ParentDepartmentId,omitempty" name:"ParentDepartmentId"`
	// 部门成员列表

	UserList []*WechatUserAttr `json:"UserList,omitempty" name:"UserList"`
}

type GetWorkWeixinOpenAppMemberRequest struct {
	*tchttp.BaseRequest
}

func (r *GetWorkWeixinOpenAppMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinOpenAppMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListIdentityProviderRequest struct {
	*tchttp.BaseRequest
}

func (r *ListIdentityProviderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListIdentityProviderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchBindWorkWeixinAccountRequest struct {
	*tchttp.BaseRequest

	// Items

	Items []*Items `json:"Items,omitempty" name:"Items"`
}

func (r *BatchBindWorkWeixinAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchBindWorkWeixinAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchBindWorkWeixinAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchBindWorkWeixinAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchBindWorkWeixinAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWorkWeixinConfigRequest struct {
	*tchttp.BaseRequest

	// AppName

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// AgentId

	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
	// Secret

	Secret *string `json:"Secret,omitempty" name:"Secret"`
	// CorpId

	CorpId *string `json:"CorpId,omitempty" name:"CorpId"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// SyncFlag

	SyncFlag *uint64 `json:"SyncFlag,omitempty" name:"SyncFlag"`
	// Status

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *UpdateWorkWeixinConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateWorkWeixinConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WechatUserAttr struct {
	// 企业微信userId

	WechatUserId *string `json:"WechatUserId,omitempty" name:"WechatUserId"`
	// 企业微信用户名

	WechatUserName *string `json:"WechatUserName,omitempty" name:"WechatUserName"`
}

type BindWorkWeixinAccountRequest struct {
	*tchttp.BaseRequest

	// 子账号uin

	UserUin *uint64 `json:"UserUin,omitempty" name:"UserUin"`
	// 企业微信userId

	WechatUserId *string `json:"WechatUserId,omitempty" name:"WechatUserId"`
}

func (r *BindWorkWeixinAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindWorkWeixinAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLdapIdpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLdapIdpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLdapIdpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkWeixinAuthorizationScopeRequest struct {
	*tchttp.BaseRequest
}

func (r *GetWorkWeixinAuthorizationScopeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinAuthorizationScopeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkWeixinOpenAppConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 内部应用id

		AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
		// 内部应用名称

		AppName *string `json:"AppName,omitempty" name:"AppName"`
		// 企业id

		CorpId *string `json:"CorpId,omitempty" name:"CorpId"`
		// 关联时间

		CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
		// 应用存储id

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 内部应用秘钥

		Secret *string `json:"Secret,omitempty" name:"Secret"`
		// 更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetWorkWeixinOpenAppConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkWeixinOpenAppConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
