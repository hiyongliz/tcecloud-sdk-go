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

package v20200921

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DescribeApiDispatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json字符串

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiDispatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiDispatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateUrlRequest struct {
	*tchttp.BaseRequest

	// 修改系统升级URL的Md5版本号

	Md5sum *string `json:"Md5sum,omitempty" name:"Md5sum"`
}

func (r *UpdateUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiDispatchRequest struct {
	*tchttp.BaseRequest

	// json字符串，必须包含一个cmd

	ReqParams *string `json:"ReqParams,omitempty" name:"ReqParams"`
}

func (r *DescribeApiDispatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiDispatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReleaseVersionRequest struct {
	*tchttp.BaseRequest

	// 模块版本

	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`
	// 模块信息

	ModuleValue *string `json:"ModuleValue,omitempty" name:"ModuleValue"`
}

func (r *ModifyReleaseVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReleaseVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRollbackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误标识

		Code *bool `json:"Code,omitempty" name:"Code"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRollbackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetCvmConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误标识

		Code *bool `json:"Code,omitempty" name:"Code"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetCvmConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetCvmConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdataConfigRequest struct {
	*tchttp.BaseRequest

	// 产品设置-修改cvm log 配置

	ConfigString *string `json:"ConfigString,omitempty" name:"ConfigString"`
	// 地区

	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *UpdataConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdataConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddUrlRequest struct {
	*tchttp.BaseRequest

	// 新增系统升级URL

	Url *string `json:"Url,omitempty" name:"Url"`
	// 新增系统升级URL版本

	UrlVersion *string `json:"UrlVersion,omitempty" name:"UrlVersion"`
}

func (r *AddUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateModuleConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功标识

		Code *bool `json:"Code,omitempty" name:"Code"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateModuleConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateModuleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddKeepIdRequest struct {
	*tchttp.BaseRequest

	// 新增保留网段信息

	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *AddKeepIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddKeepIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCvmConfigRequest struct {
	*tchttp.BaseRequest

	// 版本信息

	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`
	// 地域相关信息

	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DeleteCvmConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCvmConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回信息

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功提示Code

		Code *bool `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddModuleConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功标识

		Code *bool `json:"Code,omitempty" name:"Code"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddModuleConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddModuleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// DEVLOPMENT: 测试环境， PRODUCTION: 正式环境

		Env *string `json:"Env,omitempty" name:"Env"`
		// 1:有权限访问, 0:无权限访问

		HasAuthority *uint64 `json:"HasAuthority,omitempty" name:"HasAuthority"`
		// 是否为白名单用户

		IsWhiteListUser *uint64 `json:"IsWhiteListUser,omitempty" name:"IsWhiteListUser"`
		// 用户配置信息

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRollbackRequest struct {
	*tchttp.BaseRequest

	// 模块版本

	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`
}

func (r *ModifyRollbackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRollbackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdataConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功标识

		Code *bool `json:"Code,omitempty" name:"Code"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdataConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdataConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddModuleConfigRequest struct {
	*tchttp.BaseRequest

	// 模块版本

	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`
	// 地域信息

	Area *string `json:"Area,omitempty" name:"Area"`
	// 模块信息

	ModuleValue *string `json:"ModuleValue,omitempty" name:"ModuleValue"`
}

func (r *AddModuleConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddModuleConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateModuleConfigRequest struct {
	*tchttp.BaseRequest

	// 模块版本

	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`
	// 地域信息

	Area *string `json:"Area,omitempty" name:"Area"`
	// 模块信息

	ModuleValue *string `json:"ModuleValue,omitempty" name:"ModuleValue"`
}

func (r *UpdateModuleConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateModuleConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReleaseVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误标识

		Code *bool `json:"Code,omitempty" name:"Code"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyReleaseVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReleaseVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateKeepIdRequest struct {
	*tchttp.BaseRequest

	// 设置功能修改网段

	ConfigString *string `json:"ConfigString,omitempty" name:"ConfigString"`
}

func (r *UpdateKeepIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateKeepIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddKeepIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功提示Code

		Code *string `json:"Code,omitempty" name:"Code"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddKeepIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddKeepIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetCvmConfigRequest struct {
	*tchttp.BaseRequest

	// 版本信息

	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`
	// 地域相关信息

	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *ResetCvmConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetCvmConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateKeepIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功标识

		Code *bool `json:"Code,omitempty" name:"Code"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateKeepIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateKeepIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCvmConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误标识

		Code *bool `json:"Code,omitempty" name:"Code"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCvmConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCvmConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
