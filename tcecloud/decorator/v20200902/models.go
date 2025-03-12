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

package v20200902

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type CommitPreviewRequest struct {
	*tchttp.BaseRequest

	// 路径信息

	Path *string `json:"Path,omitempty" name:"Path"`
}

func (r *CommitPreviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CommitPreviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelPreviewRequest struct {
	*tchttp.BaseRequest

	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
}

func (r *DelPreviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelPreviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PutComponentRequest struct {
	*tchttp.BaseRequest

	// input

	Params *ComponentTempInput `json:"Params,omitempty" name:"Params"`
}

func (r *PutComponentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PutComponentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelComponentRequest struct {
	*tchttp.BaseRequest

	// 模板名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DelComponentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelComponentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryComponentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryComponentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComponentTempInput struct {
	// 模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 模板描述

	Note *string `json:"Note,omitempty" name:"Note"`
	// 此次编辑人

	Editor *string `json:"Editor,omitempty" name:"Editor"`
	// 模板内容

	Template *string `json:"Template,omitempty" name:"Template"`
}

type QueryPagesRequest struct {
	*tchttp.BaseRequest

	// input

	Params *PageTempInput `json:"Params,omitempty" name:"Params"`
}

func (r *QueryPagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPreviewRequest struct {
	*tchttp.BaseRequest

	// input

	Params *PageTempInput `json:"Params,omitempty" name:"Params"`
}

func (r *QueryPreviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPreviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPreviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryPreviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PutComponentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutComponentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PutComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelPreviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DelPreviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryComponentsRequest struct {
	*tchttp.BaseRequest

	// input

	Params *ComponentTempInput `json:"Params,omitempty" name:"Params"`
}

func (r *QueryComponentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryComponentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommitPreviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CommitPreviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CommitPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelPageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DelPageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelPageRequest struct {
	*tchttp.BaseRequest

	// 路径信息

	Path *string `json:"Path,omitempty" name:"Path"`
}

func (r *DelPageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelPageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PutPreviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutPreviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PutPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryPagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelComponentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DelComponentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PutPreviewRequest struct {
	*tchttp.BaseRequest

	// input

	Params *PageTempInput `json:"Params,omitempty" name:"Params"`
}

func (r *PutPreviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PutPreviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PageTempInput struct {
	// 模板绑定的路径信息

	Path *string `json:"Path,omitempty" name:"Path"`
	// 模板内容

	Template *string `json:"Template,omitempty" name:"Template"`
	// 此模板的编辑人

	Editor *string `json:"Editor,omitempty" name:"Editor"`
}
