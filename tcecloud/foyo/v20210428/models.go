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

package v20210428

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type SearchNamespacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSwitchListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchSwitchListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSwitchListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollbackToVersionParams struct {
	// 项id

	EntryID *int64 `json:"EntryID,omitempty" name:"EntryID"`
	// 版本号

	Version *int64 `json:"Version,omitempty" name:"Version"`
}

type QueryNamespacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果

		Data []*SearchNamespacesOutput `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchNamespacesOutput struct {
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type DrawSwitchesRequest struct {
	*tchttp.BaseRequest

	// input

	Params []*DrawSwitchesParams `json:"Params,omitempty" name:"Params"`
}

func (r *DrawSwitchesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DrawSwitchesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DrawSwitchesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DrawSwitchesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DrawSwitchesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchNamespaceParams struct {
	// 关键词

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type QueryNamespacesRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySwitchListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySwitchListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySwitchListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySwitchParams struct {
	// key

	Key *string `json:"Key,omitempty" name:"Key"`
	// ns

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type QuerySwitchRequest struct {
	*tchttp.BaseRequest

	// input

	Params *QuerySwitchParams `json:"Params,omitempty" name:"Params"`
}

func (r *QuerySwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollbackToVersionRequest struct {
	*tchttp.BaseRequest

	// input

	Params *RollbackToVersionParams `json:"Params,omitempty" name:"Params"`
}

func (r *RollbackToVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RollbackToVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchNamespacesRequest struct {
	*tchttp.BaseRequest

	// params

	Params *SearchNamespaceParams `json:"Params,omitempty" name:"Params"`
}

func (r *SearchNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSwitchListParams struct {
	// 关键词

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// ns

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type QuerySwitchVersionsParams struct {
	// key

	Key *string `json:"Key,omitempty" name:"Key"`
	// ns

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type QuerySwitchVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySwitchVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySwitchVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySwitchVersionsRequest struct {
	*tchttp.BaseRequest

	// PageSize

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// PageNum

	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`
	// input

	Params *QuerySwitchVersionsParams `json:"Params,omitempty" name:"Params"`
}

func (r *QuerySwitchVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySwitchVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSwitchListRequest struct {
	*tchttp.BaseRequest

	// PageNum

	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`
	// PageSize

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// params

	Params *SearchSwitchListParams `json:"Params,omitempty" name:"Params"`
}

func (r *SearchSwitchListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSwitchListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySwitchListParams struct {
	// key

	Key *string `json:"Key,omitempty" name:"Key"`
	// ns

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type QuerySwitchListRequest struct {
	*tchttp.BaseRequest

	// input

	Params *QuerySwitchListParams `json:"Params,omitempty" name:"Params"`
	// PageSize

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// PageNum

	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`
}

func (r *QuerySwitchListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySwitchListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollbackToVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RollbackToVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RollbackToVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DrawSwitchesParams struct {
	// key命名

	Key *string `json:"Key,omitempty" name:"Key"`
	// value

	Value *string `json:"Value,omitempty" name:"Value"`
	// ns

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
}
