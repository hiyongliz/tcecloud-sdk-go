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

package v20200613

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type CreateLoginPolicyRequest struct {
	*tchttp.BaseRequest

	// 是否关闭

	DisablePolicy *int64 `json:"DisablePolicy,omitempty" name:"DisablePolicy"`
	// IP策略

	IpPolices []*IpPolices `json:"IpPolices,omitempty" name:"IpPolices"`
}

func (r *CreateLoginPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLoginPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLoginPolicyRequest struct {
	*tchttp.BaseRequest
}

func (r *GetLoginPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLoginPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPasswordConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetPasswordConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPasswordConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePasswordConfigRequest struct {
	*tchttp.BaseRequest

	// 配置id

	ConfigId *int64 `json:"ConfigId,omitempty" name:"ConfigId"`
	// 第一次登录是否修改

	IsFirstTimeModify *int64 `json:"IsFirstTimeModify,omitempty" name:"IsFirstTimeModify"`
	// 是否包含大写

	IsContainUpper *int64 `json:"IsContainUpper,omitempty" name:"IsContainUpper"`
	// 是否包含小写

	IsContainLower *int64 `json:"IsContainLower,omitempty" name:"IsContainLower"`
	// 是否包含数字

	IsContainDigital *int64 `json:"IsContainDigital,omitempty" name:"IsContainDigital"`
	// 是否包含特殊字符

	IsContainSpecialCharacter *int64 `json:"IsContainSpecialCharacter,omitempty" name:"IsContainSpecialCharacter"`
	// 最短长度

	ShortestLength *int64 `json:"ShortestLength,omitempty" name:"ShortestLength"`
	// 重复次数

	RepeatNum *int64 `json:"RepeatNum,omitempty" name:"RepeatNum"`
	// 密码过期时间

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

func (r *UpdatePasswordConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePasswordConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePasswordConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePasswordConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePasswordConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLoginPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLoginPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLoginPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPasswordConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPasswordConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPasswordConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateLoginPolicyRequest struct {
	*tchttp.BaseRequest

	// IP

	IP *string `json:"IP,omitempty" name:"IP"`
}

func (r *ValidateLoginPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateLoginPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateLoginPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ValidateLoginPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateLoginPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IpPolices struct {
	// IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// Effect

	Effect *string `json:"Effect,omitempty" name:"Effect"`
}

type CreateLoginPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLoginPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLoginPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
