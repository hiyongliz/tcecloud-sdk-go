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

package v20200102

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type PreLoginResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PreLoginResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreLoginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UinInfo struct {
	// uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户uid

	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type VerifyMenuRequest struct {
	*tchttp.BaseRequest

	// 菜单

	Menu *string `json:"Menu,omitempty" name:"Menu"`
	// 子菜单

	Submenu *string `json:"Submenu,omitempty" name:"Submenu"`
}

func (r *VerifyMenuRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyMenuRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLoginActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 需要进行的操作。如果不需要则等于continue

		Action *string `json:"Action,omitempty" name:"Action"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLoginActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLoginActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyMenuResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyMenuResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyMenuResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupUserInfo struct {
	// 接收者id

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 账户唯一id

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 手机号

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 国家编码

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 手机号标识

	PhoneFlag *string `json:"PhoneFlag,omitempty" name:"PhoneFlag"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 邮箱标识

	EmailFlag *string `json:"EmailFlag,omitempty" name:"EmailFlag"`
	// 用户类型

	UserType *int64 `json:"UserType,omitempty" name:"UserType"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 是否是主账户

	IsReceiverOwner *int64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`
	// 账户类型

	SystemType *string `json:"SystemType,omitempty" name:"SystemType"`
}

type CheckLoginDetail struct {
	// 登陆类型

	LoginType *string `json:"LoginType,omitempty" name:"LoginType"`
	// 登陆到到地址

	LoginTo *string `json:"LoginTo,omitempty" name:"LoginTo"`
	// 登陆平台

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 客户端类型

	ClientType *string `json:"ClientType,omitempty" name:"ClientType"`
	// 是否使用多因子认证

	MFAUsed *string `json:"MFAUsed,omitempty" name:"MFAUsed"`
}

type CheckLoginRequest struct {
	*tchttp.BaseRequest

	// 详细请求参数

	Detail []*CheckLoginDetail `json:"Detail,omitempty" name:"Detail"`
	// 临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
	// 客户端代理信息

	ClientUA *string `json:"ClientUA,omitempty" name:"ClientUA"`
	// 登陆平台

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 是否作为登陆行为进行上报

	ReportLoginInfo *bool `json:"ReportLoginInfo,omitempty" name:"ReportLoginInfo"`
	// 用户标识ID

	Tinyid *string `json:"Tinyid,omitempty" name:"Tinyid"`
	// 是否进行单点登录剔除，1-是

	Single_login *int64 `json:"Single_login,omitempty" name:"Single_login"`
}

func (r *CheckLoginRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckLoginRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckLoginResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckLoginResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckLoginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelLoginForbidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DelLoginForbidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelLoginForbidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUidByUinArrRequest struct {
	*tchttp.BaseRequest

	// uin列表

	UinList []*uint64 `json:"UinList,omitempty" name:"UinList"`
}

func (r *GetUidByUinArrRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUidByUinArrRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUidByUinArrResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// uin对应的Uid信息

		UinInfo []*UinInfo `json:"UinInfo,omitempty" name:"UinInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUidByUinArrResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUidByUinArrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginData struct {
	// 用户名

	Account *string `json:"Account,omitempty" name:"Account"`
	// 登陆密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

type LoginDetail struct {
	// 登陆账户类型

	LoginType *string `json:"LoginType,omitempty" name:"LoginType"`
	// 登陆的平台地址

	LoginTo *string `json:"LoginTo,omitempty" name:"LoginTo"`
	// 登陆平台

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 客户端类型

	ClientType *string `json:"ClientType,omitempty" name:"ClientType"`
}

type GetLoginActionRequest struct {
	*tchttp.BaseRequest

	// 登陆临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
	// 客户端用户代理

	ClientUA *string `json:"ClientUA,omitempty" name:"ClientUA"`
}

func (r *GetLoginActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLoginActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelLoginForbidRequest struct {
	*tchttp.BaseRequest

	// 用户名

	Account *string `json:"Account,omitempty" name:"Account"`
}

func (r *DelLoginForbidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelLoginForbidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreLoginRequest struct {
	*tchttp.BaseRequest

	// 客户端代理信息

	ClientUA *string `json:"ClientUA,omitempty" name:"ClientUA"`
	// 登陆平台

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 是否作为登陆事件进行上报

	ReportLoginInfo *bool `json:"ReportLoginInfo,omitempty" name:"ReportLoginInfo"`
	// 登陆详细信息

	Detail *LoginDetail `json:"Detail,omitempty" name:"Detail"`
	// 登陆类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 区域类型，1-中国

	AreaType *string `json:"AreaType,omitempty" name:"AreaType"`
	// 登陆账户信息

	Data []*LoginData `json:"Data,omitempty" name:"Data"`
}

func (r *PreLoginRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreLoginRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
