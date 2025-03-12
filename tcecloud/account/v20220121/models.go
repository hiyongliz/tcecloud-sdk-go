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
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type ListSelfGroupsForUserRequest struct {
	*tchttp.BaseRequest

	// 接收者用户id，uid和subUin传一个

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 单页数量

	Rp *int64 `json:"Rp,omitempty" name:"Rp"`
	// 分页数

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 被查询的子账号uin，uid和subUin传一个

	SubUin *uint64 `json:"SubUin,omitempty" name:"SubUin"`
}

func (r *ListSelfGroupsForUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSelfGroupsForUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableSelfApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableSelfApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableSelfApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSelfApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API密钥数据列表

		IdKeys []*ApiKeyDetail `json:"IdKeys,omitempty" name:"IdKeys"`
		// 账号Uin

		Uin *string `json:"Uin,omitempty" name:"Uin"`
		// 项目Id

		SecretProjectId *uint64 `json:"SecretProjectId,omitempty" name:"SecretProjectId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSelfApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSelfApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSelfApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
}

func (r *GetSelfApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSelfApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySelfApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API密钥数据列表

		IdKeys []*ApiKey `json:"IdKeys,omitempty" name:"IdKeys"`
		// 账号uin

		Uin *string `json:"Uin,omitempty" name:"Uin"`
		// SecretProjectId

		SecretProjectId *int64 `json:"SecretProjectId,omitempty" name:"SecretProjectId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySelfApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySelfApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupInfo struct {
	// 组id

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 息接收渠道 0:无 1: 短信 2：邮件 3：短信+邮件

	Channel *int64 `json:"Channel,omitempty" name:"Channel"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 用户组成员信息

	UserInfo []*GroupMemberInfo `json:"UserInfo,omitempty" name:"UserInfo"`
	// 用户组类型，0-自定义，1-预设

	GroupType *int64 `json:"GroupType,omitempty" name:"GroupType"`
}

type ListSelfAttachedUserAllPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略列表数据

		PolicyList []*AttachedUserPolicy `json:"PolicyList,omitempty" name:"PolicyList"`
		// 策略总数

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSelfAttachedUserAllPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSelfAttachedUserAllPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSelfApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSelfApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSelfApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySelfApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySelfApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySelfApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSelfAttachedUserAllPoliciesRequest struct {
	*tchttp.BaseRequest

	// 0:返回直接关联和随组关联策略，1:只返回直接关联策略，2:只返回随组关联策略

	AttachType *uint64 `json:"AttachType,omitempty" name:"AttachType"`
	// 策略类型

	StrategyType *uint64 `json:"StrategyType,omitempty" name:"StrategyType"`
	// 每页数量，默认值是 20，必须大于 0 且小于或等于 200

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
	// 页码，默认值是 1，从 1开始，不能大于 200

	Page *uint64 `json:"Page,omitempty" name:"Page"`
}

func (r *ListSelfAttachedUserAllPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSelfAttachedUserAllPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableSelfApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
}

func (r *EnableSelfApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableSelfApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
	// 用户Uin

	ApiUin *uint64 `json:"ApiUin,omitempty" name:"ApiUin"`
	// 备注

	ApiSecretIdRemark *string `json:"ApiSecretIdRemark,omitempty" name:"ApiSecretIdRemark"`
}

func (r *ModifyApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableSelfApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
}

func (r *DisableSelfApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableSelfApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupMemberInfo struct {
	// 子用户 Uid。

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 子用户 Uin。

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 子用户用户名。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 手机号。

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 手机区域代码。

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 是否已验证手机。

	PhoneFlag *int64 `json:"PhoneFlag,omitempty" name:"PhoneFlag"`
	// 邮箱地址。

	Email *string `json:"Email,omitempty" name:"Email"`
	// 是否已验证邮箱。

	EmailFlag *int64 `json:"EmailFlag,omitempty" name:"EmailFlag"`
	// 用户类型。

	UserType *int64 `json:"UserType,omitempty" name:"UserType"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 是否为主消息接收人。

	IsReceiverOwner *int64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`
}

type ApiKey struct {
	// 密钥ID

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// 创建时间(时间戳)

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 状态(2:有效, 3:禁用)

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 数据源

	Source *uint64 `json:"Source,omitempty" name:"Source"`
}

type DisableSelfApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableSelfApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableSelfApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySelfApiKeyRequest struct {
	*tchttp.BaseRequest
}

func (r *QuerySelfApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySelfApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSelfApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
}

func (r *DeleteSelfApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSelfApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSelfGroupsForUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总体数量

		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`
		// 用户组信息列表

		GroupInfo []*GroupInfo `json:"GroupInfo,omitempty" name:"GroupInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSelfGroupsForUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSelfGroupsForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySelfApiKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥Id

	ApiSecretId *string `json:"ApiSecretId,omitempty" name:"ApiSecretId"`
	// 备注

	ApiSecretIdRemark *string `json:"ApiSecretIdRemark,omitempty" name:"ApiSecretIdRemark"`
}

func (r *ModifySelfApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySelfApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSelfApiKeyRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateSelfApiKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSelfApiKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachedUserPolicyGroupInfo struct {
	// 分组ID

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type CreateSelfApiKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 持久密钥

		IdKeys []*ApiKey `json:"IdKeys,omitempty" name:"IdKeys"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSelfApiKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSelfApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiKeyDetail struct {
	// 密钥ID

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// 密钥Key

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
	// 创建时间(时间戳)

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 状态(2:有效, 3:禁用)

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 数据源

	Source *uint64 `json:"Source,omitempty" name:"Source"`
}

type AttachedUserPolicy struct {
	// 策略ID

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 策略类型(1表示自定义策略，2表示预设策略)

	StrategyType *uint64 `json:"StrategyType,omitempty" name:"StrategyType"`
	// 创建模式(1表示按产品或项目权限创建的策略，其他表示策略语法创建的策略)

	CreateMode *uint64 `json:"CreateMode,omitempty" name:"CreateMode"`
	// 随组关联信息

	Groups []*AttachedUserPolicyGroupInfo `json:"Groups,omitempty" name:"Groups"`
}
