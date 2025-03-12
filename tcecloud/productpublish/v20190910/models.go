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

package v20190910

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type InsertOperateHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InsertOperateHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InsertOperateHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryReleaseStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryReleaseStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryReleaseStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchReleaseDetailRequest struct {
	*tchttp.BaseRequest

	// 发布ID

	ReleaseId *int64 `json:"ReleaseId,omitempty" name:"ReleaseId"`
}

func (r *SearchReleaseDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchReleaseDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomDeployResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CustomDeployResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CustomDeployResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InsertOperateHistoryRequest struct {
	*tchttp.BaseRequest

	// 信息主题

	InfoTheme *string `json:"InfoTheme,omitempty" name:"InfoTheme"`
	// 所属平台

	ProductPlatform *string `json:"ProductPlatform,omitempty" name:"ProductPlatform"`
	// 所属标签

	TabTypeCode *string `json:"TabTypeCode,omitempty" name:"TabTypeCode"`
	// 产品简称

	ProductShortCode *string `json:"ProductShortCode,omitempty" name:"ProductShortCode"`
	// 操作类型：增/删/改

	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
	// 修改前的值

	PreValue *string `json:"PreValue,omitempty" name:"PreValue"`
	// 修改后的值

	AfterValue *string `json:"AfterValue,omitempty" name:"AfterValue"`
}

func (r *InsertOperateHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InsertOperateHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OuterCustomizeProductHandlerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OuterCustomizeProductHandlerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OuterCustomizeProductHandlerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseRequest struct {
	*tchttp.BaseRequest

	// 所属平台

	ProductPlatform *string `json:"ProductPlatform,omitempty" name:"ProductPlatform"`
}

func (r *ReleaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchOperateHistoryRequest struct {
	*tchttp.BaseRequest

	// 所属平台

	ProductPlatform *string `json:"ProductPlatform,omitempty" name:"ProductPlatform"`
	// 所属标签

	TabTypeCode *string `json:"TabTypeCode,omitempty" name:"TabTypeCode"`
}

func (r *SearchOperateHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchOperateHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OuterCustomizeProductHandlerRequest struct {
	*tchttp.BaseRequest

	// 原客制化系统的注册的服务名

	DispatchService *string `json:"DispatchService,omitempty" name:"DispatchService"`
	// 原客制化系统服务下的方法名

	DispatchAction *string `json:"DispatchAction,omitempty" name:"DispatchAction"`
	// 原客制化系统data参数序列化后的值

	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *OuterCustomizeProductHandlerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OuterCustomizeProductHandlerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchReleaseHistoryRequest struct {
	*tchttp.BaseRequest

	// 返回数据个数限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询条件

	Condition *string `json:"Condition,omitempty" name:"Condition"`
}

func (r *SearchReleaseHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchReleaseHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchReleaseHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchReleaseHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchReleaseHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryReleaseStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryReleaseStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryReleaseStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomDeployRequest struct {
	*tchttp.BaseRequest

	// 客制化包所在的路径

	CustomizeFilePath *string `json:"CustomizeFilePath,omitempty" name:"CustomizeFilePath"`
	// 发布ID

	ReleaseId *int64 `json:"ReleaseId,omitempty" name:"ReleaseId"`
	// xxx路径

	Xxx1Path *string `json:"Xxx1Path,omitempty" name:"Xxx1Path"`
}

func (r *CustomDeployRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CustomDeployRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchReleaseDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchReleaseDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchReleaseDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchOperateHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchOperateHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchOperateHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
