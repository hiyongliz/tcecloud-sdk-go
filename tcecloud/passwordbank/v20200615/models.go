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
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type LogDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LogDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LogDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PasswordResetTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PasswordResetTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PasswordResetTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetPasswordResetTimeRequest struct {
	*tchttp.BaseRequest

	// 重置周期

	Time *int64 `json:"Time,omitempty" name:"Time"`
	// 重置时间类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *SetPasswordResetTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetPasswordResetTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationLogsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件信息

	Filter []*Filter `json:"Filter,omitempty" name:"Filter"`
	// 限制返回个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *OperationLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OperationLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OperationLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OperationLogsResponse) FromJsonString(s string) error {
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

type ServerListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ServerListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerListsResponse) FromJsonString(s string) error {
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

type JoinToBanksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *JoinToBanksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *JoinToBanksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JoinToBanksFromExternalRequest struct {
	*tchttp.BaseRequest

	// 入库设备信息

	Servers []*Servers `json:"Servers,omitempty" name:"Servers"`
	// 密码类型

	PasswordType *int64 `json:"PasswordType,omitempty" name:"PasswordType"`
}

func (r *JoinToBanksFromExternalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *JoinToBanksFromExternalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordRequest struct {
	*tchttp.BaseRequest

	// 设备信息

	Servers []*Servers `json:"Servers,omitempty" name:"Servers"`
}

func (r *ModifyPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PasswordsRequest struct {
	*tchttp.BaseRequest

	// 设备信息列表

	Servers []*Servers `json:"Servers,omitempty" name:"Servers"`
}

func (r *PasswordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PasswordsRequest) FromJsonString(s string) error {
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

type GetPasswordRuleRequest struct {
	*tchttp.BaseRequest
}

func (r *GetPasswordRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPasswordRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JoinToBanksFromExternalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *JoinToBanksFromExternalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *JoinToBanksFromExternalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Rule struct {
	// 规则序号，可选值：1,2,3,4,-1,-2。其他规则序号的规则会被忽略

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 字符集

	Charset *string `json:"Charset,omitempty" name:"Charset"`
	// 该规则对应的字符最少数目

	Length *int64 `json:"Length,omitempty" name:"Length"`
	// 规则是否生效，值大于0为有效

	Valid *int64 `json:"Valid,omitempty" name:"Valid"`
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

type Servers struct {
	// 内网IP

	IntranetIp *string `json:"IntranetIp,omitempty" name:"IntranetIp"`
	// 设备ID

	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 长期密码

	FixedCipher *string `json:"FixedCipher,omitempty" name:"FixedCipher"`
	// 密码类型

	PasswordType *int64 `json:"PasswordType,omitempty" name:"PasswordType"`
}

type JoinToBanksRequest struct {
	*tchttp.BaseRequest

	// 入库设备信息

	Servers []*Servers `json:"Servers,omitempty" name:"Servers"`
	// 密码类型

	PasswordType *int64 `json:"PasswordType,omitempty" name:"PasswordType"`
}

func (r *JoinToBanksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *JoinToBanksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogDetailRequest struct {
	*tchttp.BaseRequest

	// 过滤条件信息

	Filter []*Filter `json:"Filter,omitempty" name:"Filter"`
	// 限制返回个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *LogDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LogDetailRequest) FromJsonString(s string) error {
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

type QuitFromBanksRequest struct {
	*tchttp.BaseRequest

	// 操作类型，7：强制退库 8：修改密码并退库(需要Password参数) 9：校验密码并退库

	OpType *int64 `json:"OpType,omitempty" name:"OpType"`
	// 退库机器信息

	Servers []*Servers `json:"Servers,omitempty" name:"Servers"`
}

func (r *QuitFromBanksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuitFromBanksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerListsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件信息

	Filter []*Filter `json:"Filter,omitempty" name:"Filter"`
	// 限制返回个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ServerListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetPasswordRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetPasswordRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetPasswordRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 需要过滤的字段

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段的过滤值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type QuitFromBanksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuitFromBanksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuitFromBanksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPasswordRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPasswordRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPasswordRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetPasswordResetTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetPasswordResetTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetPasswordResetTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetPasswordRuleRequest struct {
	*tchttp.BaseRequest

	// 密码长度

	Length *int64 `json:"Length,omitempty" name:"Length"`
	// 密码规则

	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`
}

func (r *SetPasswordRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetPasswordRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetPasswordTypesRequest struct {
	*tchttp.BaseRequest

	// 目标密码类型

	PasswordType *int64 `json:"PasswordType,omitempty" name:"PasswordType"`
	// 设备信息列表

	Servers []*Servers `json:"Servers,omitempty" name:"Servers"`
}

func (r *SetPasswordTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetPasswordTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PasswordResetTimeRequest struct {
	*tchttp.BaseRequest

	// 重置时间类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *PasswordResetTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PasswordResetTimeRequest) FromJsonString(s string) error {
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
