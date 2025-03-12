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

package v20210808

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type GetCustomOpenPlatformSecretInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 密钥配置信息

		SecretInfo *SecretInfo `json:"SecretInfo,omitempty" name:"SecretInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCustomOpenPlatformSecretInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomOpenPlatformSecretInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCustomOpenPlatformResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数.

		Cnt *int64 `json:"Cnt,omitempty" name:"Cnt"`
		// 开放平台数据列表

		List []*AppAttr `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCustomOpenPlatformResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCustomOpenPlatformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomOpenPlatformRequest struct {
	*tchttp.BaseRequest

	// 开放平台id

	OpenId *uint64 `json:"OpenId,omitempty" name:"OpenId"`
	// 回调域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 开放平台名称

	OpenName *string `json:"OpenName,omitempty" name:"OpenName"`
	// 备注信息

	Memo *string `json:"Memo,omitempty" name:"Memo"`
}

func (r *ModifyCustomOpenPlatformRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomOpenPlatformRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomOpenPlatformResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomOpenPlatformResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomOpenPlatformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCustomOpenPlatformRequest struct {
	*tchttp.BaseRequest

	// 页码

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 页码

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤平台列表，0-正常，1-冻结

	State *string `json:"State,omitempty" name:"State"`
	// 应用名

	FuzzyName *string `json:"FuzzyName,omitempty" name:"FuzzyName"`
}

func (r *ListCustomOpenPlatformRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCustomOpenPlatformRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomOpenPlatformSecretInfoRequest struct {
	*tchttp.BaseRequest

	// 应用id

	OpenId *uint64 `json:"OpenId,omitempty" name:"OpenId"`
}

func (r *GetCustomOpenPlatformSecretInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomOpenPlatformSecretInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppAttr struct {
	// 应用ID

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 回调域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 备注

	Memo *string `json:"Memo,omitempty" name:"Memo"`
	// 修改用户

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 主域名

	OpenHome *string `json:"OpenHome,omitempty" name:"OpenHome"`
	// 开放平台ID

	OpenId *int64 `json:"OpenId,omitempty" name:"OpenId"`
	// 名称

	OpenName *string `json:"OpenName,omitempty" name:"OpenName"`
	// 使用状态

	State *int64 `json:"State,omitempty" name:"State"`
	// 应用状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type SecretInfo struct {
	// 应用ID

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 应用ak

	AppSecretId *string `json:"AppSecretId,omitempty" name:"AppSecretId"`
	// 应用sk

	AppSecretKey *string `json:"AppSecretKey,omitempty" name:"AppSecretKey"`
	// 开放平台id

	OpenId *int64 `json:"OpenId,omitempty" name:"OpenId"`
}

type ApplyCustomOpenPlatformRequest struct {
	*tchttp.BaseRequest

	// 回调域名（去掉http[s]的部分，保证回调地址在这个域名之下，不能只是顶级域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 平台名称

	OpenName *string `json:"OpenName,omitempty" name:"OpenName"`
	// 备注信息

	Memo *string `json:"Memo,omitempty" name:"Memo"`
}

func (r *ApplyCustomOpenPlatformRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyCustomOpenPlatformRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageCustomOpenPlatformRequest struct {
	*tchttp.BaseRequest

	// 应用id

	OpenId *uint64 `json:"OpenId,omitempty" name:"OpenId"`
	// 0为冻结，1为恢复

	ManageType *string `json:"ManageType,omitempty" name:"ManageType"`
}

func (r *ManageCustomOpenPlatformRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageCustomOpenPlatformRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageCustomOpenPlatformResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageCustomOpenPlatformResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageCustomOpenPlatformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomOpenPlatformBasicInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 基础信息

		BasicInfo *BasicInfoAttr `json:"BasicInfo,omitempty" name:"BasicInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCustomOpenPlatformBasicInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomOpenPlatformBasicInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyCustomOpenPlatformResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用Id

		OpenId *int64 `json:"OpenId,omitempty" name:"OpenId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyCustomOpenPlatformResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyCustomOpenPlatformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCustomOpenPlatformBasicInfoRequest struct {
	*tchttp.BaseRequest

	// 申请时返回的open_id

	OpenId *uint64 `json:"OpenId,omitempty" name:"OpenId"`
}

func (r *GetCustomOpenPlatformBasicInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCustomOpenPlatformBasicInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BasicInfoAttr struct {
	// 开放平台ID

	OpenId *int64 `json:"OpenId,omitempty" name:"OpenId"`
	// 应用ID

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 开放平台名称

	OpenName *string `json:"OpenName,omitempty" name:"OpenName"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 使用状态,0：正常，1：冻结

	State *int64 `json:"State,omitempty" name:"State"`
	// 修改者

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 授权url

	TicketUrl *string `json:"TicketUrl,omitempty" name:"TicketUrl"`
	// 授权列表

	BakGrantList *string `json:"BakGrantList,omitempty" name:"BakGrantList"`
	// logo url地址

	OpenLogo *string `json:"OpenLogo,omitempty" name:"OpenLogo"`
	// silent

	Silent *int64 `json:"Silent,omitempty" name:"Silent"`
	// 开放类型

	OpenType *int64 `json:"OpenType,omitempty" name:"OpenType"`
	// 备注

	Memo *string `json:"Memo,omitempty" name:"Memo"`
	// client对应key

	ClientKey *string `json:"ClientKey,omitempty" name:"ClientKey"`
	// 应用主域名

	OpenHome *string `json:"OpenHome,omitempty" name:"OpenHome"`
}
