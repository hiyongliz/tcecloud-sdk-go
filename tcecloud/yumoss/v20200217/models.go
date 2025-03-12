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

package v20200217

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DeleteYumFileRequest struct {
	*tchttp.BaseRequest

	// 文件名称（包含路径信息）

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 资源仓库名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 返回提示信息的语言格式。可选值：English、Chinese（不填时默认返回英文）

	MessageLanguage *string `json:"MessageLanguage,omitempty" name:"MessageLanguage"`
}

func (r *DeleteYumFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteYumFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteYumFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteYumFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteYumFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InsertSyncConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InsertSyncConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InsertSyncConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InsertNewFolderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InsertNewFolderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InsertNewFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySyncConfigRequest struct {
	*tchttp.BaseRequest

	// yum源仓库名称，如：centos

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// yum源仓库描述

	ResourceDescribe *string `json:"ResourceDescribe,omitempty" name:"ResourceDescribe"`
	// 同步类型  1: rsync    2:  自定义脚本  3: 人工上传

	SyncType *int64 `json:"SyncType,omitempty" name:"SyncType"`
	// yum源rsync同步的上级源（同步类型为rsync时有值）

	SuperiorSource *string `json:"SuperiorSource,omitempty" name:"SuperiorSource"`
	// rsync的参数（同步类型为“rsync”时有值）

	SyncParameter *string `json:"SyncParameter,omitempty" name:"SyncParameter"`
	// 自定义脚本（同步类型为“自定义脚本”时有值）

	Script *string `json:"Script,omitempty" name:"Script"`
	// 是否开启定时同步  1:是  0:否

	IsTimingSync *bool `json:"IsTimingSync,omitempty" name:"IsTimingSync"`
	// 同步周期

	SyncCycle *string `json:"SyncCycle,omitempty" name:"SyncCycle"`
	// 帮助文档（前端以markdown形式展示）

	HelpMessage *string `json:"HelpMessage,omitempty" name:"HelpMessage"`
	// 同步类型为“人工上传”时的后置脚本

	PostScript *string `json:"PostScript,omitempty" name:"PostScript"`
	// 返回提示信息的语言格式。可选值：English、Chinese（不填时默认返回英文）

	MessageLanguage *string `json:"MessageLanguage,omitempty" name:"MessageLanguage"`
}

func (r *ModifySyncConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySyncConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartSyncRequest struct {
	*tchttp.BaseRequest

	// yum源仓库名称（如centos）

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 返回提示信息的语言格式。可选值：English、Chinese（不填时默认返回英文）

	MessageLanguage *string `json:"MessageLanguage,omitempty" name:"MessageLanguage"`
}

func (r *StartSyncRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartSyncRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSyncConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSyncConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSyncConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySyncConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySyncConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySyncConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchYumCatalogRequest struct {
	*tchttp.BaseRequest

	// 要查询的详细路径，注意要以“/”开头

	Route *string `json:"Route,omitempty" name:"Route"`
	// 查询数量（用于分页）

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值（用于分页）

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *SearchYumCatalogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchYumCatalogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchYumCatalogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchYumCatalogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchYumCatalogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSyncHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchSyncHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSyncHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartSyncResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartSyncResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartSyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadYumFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadYumFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadYumFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InsertSyncConfigRequest struct {
	*tchttp.BaseRequest

	// yum源仓库名称，如：centos

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// yum源仓库描述

	ResourceDescribe *string `json:"ResourceDescribe,omitempty" name:"ResourceDescribe"`
	// 同步类型  1: rsync    2:  自定义脚本  3: 人工上传

	SyncType *int64 `json:"SyncType,omitempty" name:"SyncType"`
	// yum源rsync同步的上级源（同步类型为rsync时有值）

	SuperiorSource *string `json:"SuperiorSource,omitempty" name:"SuperiorSource"`
	// rsync的参数（同步类型为“rsync”时有值）

	SyncParameter *string `json:"SyncParameter,omitempty" name:"SyncParameter"`
	// 自定义脚本（同步类型为“自定义脚本”时有值）

	Script *string `json:"Script,omitempty" name:"Script"`
	// 是否开启定时同步  1:是  0:否

	IsTimingSync *bool `json:"IsTimingSync,omitempty" name:"IsTimingSync"`
	// 同步周期

	SyncCycle *string `json:"SyncCycle,omitempty" name:"SyncCycle"`
	// 帮助文档（前端以markdown形式展示）

	HelpMessage *string `json:"HelpMessage,omitempty" name:"HelpMessage"`
	// 同步类型为“人工上传”时的后置脚本

	PostScript *string `json:"PostScript,omitempty" name:"PostScript"`
}

func (r *InsertSyncConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InsertSyncConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSyncHistoryRequest struct {
	*tchttp.BaseRequest

	// 查询数量（用于分页）

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值（用于分页）

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 筛选条件：每个筛选条件为一个Name-Value组合，如果本list为空则代表查全量数据

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *SearchSyncHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSyncHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InsertNewFolderRequest struct {
	*tchttp.BaseRequest

	// 文件夹名（包含详细的路径）

	NewFolder *string `json:"NewFolder,omitempty" name:"NewFolder"`
	// 资源仓库名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 返回提示信息的语言格式。可选值：English、Chinese（不填时默认返回英文）

	MessageLanguage *string `json:"MessageLanguage,omitempty" name:"MessageLanguage"`
}

func (r *InsertNewFolderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InsertNewFolderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchOperationHistoryRequest struct {
	*tchttp.BaseRequest

	// 查询数量（用于分页）

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值（用于分页）

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 筛选条件：每个筛选条件为一个Name-Value组合，如果本list为空则代表查全量数据

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *SearchOperationHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchOperationHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSyncConfigRequest struct {
	*tchttp.BaseRequest

	// 查询数量（用于分页）

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移值（用于分页）

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 筛选条件：每个筛选条件为一个Name-Value组合，如果本list为空则代表查全量数据

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *SearchSyncConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSyncConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSyncConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchSyncConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSyncConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckRouteEmptyRequest struct {
	*tchttp.BaseRequest

	// yum源仓库名称（如centos）

	Route *string `json:"Route,omitempty" name:"Route"`
}

func (r *CheckRouteEmptyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckRouteEmptyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSyncConfigRequest struct {
	*tchttp.BaseRequest

	// yum源仓库名称（如centos）

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 返回提示信息的语言格式。可选值：English、Chinese（不填时默认返回英文）

	MessageLanguage *string `json:"MessageLanguage,omitempty" name:"MessageLanguage"`
}

func (r *DeleteSyncConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSyncConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenameYumFileRequest struct {
	*tchttp.BaseRequest

	// 原先的文件名称（包含路径信息）

	OldFileName *string `json:"OldFileName,omitempty" name:"OldFileName"`
	// 文件最终要放到的目标路径

	NewFileName *string `json:"NewFileName,omitempty" name:"NewFileName"`
	// 资源仓库名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 返回提示信息的语言格式。可选值：English、Chinese（不填时默认返回英文）

	MessageLanguage *string `json:"MessageLanguage,omitempty" name:"MessageLanguage"`
}

func (r *RenameYumFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameYumFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadYumFileRequest struct {
	*tchttp.BaseRequest

	// 上传的文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件最终要放到的目标路径

	DestRoute *string `json:"DestRoute,omitempty" name:"DestRoute"`
	// 资源仓库名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 前端调用fileserver返回的TaskId

	FileServerTaskId *int64 `json:"FileServerTaskId,omitempty" name:"FileServerTaskId"`
	// 返回提示信息的语言格式。可选值：English、Chinese（不填时默认返回英文）

	MessageLanguage *string `json:"MessageLanguage,omitempty" name:"MessageLanguage"`
}

func (r *UploadYumFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadYumFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckRouteEmptyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckRouteEmptyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckRouteEmptyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenameYumFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenameYumFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameYumFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchOperationHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchOperationHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchOperationHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 查询条件的key

	Name *string `json:"Name,omitempty" name:"Name"`
	// 查询条件的value

	Value *string `json:"Value,omitempty" name:"Value"`
}
