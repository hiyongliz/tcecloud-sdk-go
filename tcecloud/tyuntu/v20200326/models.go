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

package v20200326

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type UpdateCategoryRequest struct {
	*tchttp.BaseRequest

	// 更新官方文档分类名参数

	DataParam *UpdateCategoryParam `json:"DataParam,omitempty" name:"DataParam"`
	// DataParam的API信息的module

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *UpdateCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddComplexObjectParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 复杂对象名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 复杂对象描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 复杂对象下的成员

	Members *string `json:"Members,omitempty" name:"Members"`
	// 出参入参类型

	UsedFor *uint64 `json:"UsedFor,omitempty" name:"UsedFor"`
}

type FilterQueryCategoryParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 官方文档名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 更新人

	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`
	// 更新起始时间

	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`
	// 更新截止时间

	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// asc或desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

type GetVersionsRequest struct {
	*tchttp.BaseRequest

	// 指定模块

	DataParam *GetVersionParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateComplexObjectRequest struct {
	*tchttp.BaseRequest

	// 更新复杂对象的参数

	DataParam *UpdateComplexObjectParam `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *UpdateComplexObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateComplexObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddComplexObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddComplexObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddComplexObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRegionsParam struct {
	// id

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 地区

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地区名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地区状态

	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`
	// 修改时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 修改人

	UpdateUsername *string `json:"UpdateUsername,omitempty" name:"UpdateUsername"`
}

type GetEnvironmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的是环境判断

		Data *bool `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetEnvironmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddVersionParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 接口版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type SystemParam struct {
	// 系统公共参数名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数说明

	Description *string `json:"Description,omitempty" name:"Description"`
	// 参数说明国际化

	DescriptionIntlEnUs *string `json:"DescriptionIntlEnUs,omitempty" name:"DescriptionIntlEnUs"`
	// 设置参数传递位置

	Position *Position `json:"Position,omitempty" name:"Position"`
}

type GetModuleErrorCodeRelateApiResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的API接口名

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetModuleErrorCodeRelateApiResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetModuleErrorCodeRelateApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateComplexObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateComplexObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateComplexObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApiInfoParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 接口名

	Action *string `json:"Action,omitempty" name:"Action"`
}

type QueryComplexObjectFilterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的列表复杂对象数据

		Data []*ComplexObject `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryComplexObjectFilterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryComplexObjectFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryModuleErrorCodeFilterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的错误码数据

		Data []*ModuleErrorcode `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryModuleErrorCodeFilterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryModuleErrorCodeFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCategoryIDParam struct {
	// 官方文档ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
}

type UpdateComplexObjectParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 复杂对象id

	ObjectID *uint64 `json:"ObjectID,omitempty" name:"ObjectID"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 成员变量

	Members *string `json:"Members,omitempty" name:"Members"`
	// 出参入参类型

	UsedFor *int64 `json:"UsedFor,omitempty" name:"UsedFor"`
	// 复杂对象名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type GetComplexObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 复杂对象

		Data []*ComplexObject `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComplexObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComplexObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// API Explorer 调用其他接口的返回

		Data *RunActionData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SdkDemoSetInfo struct {
	// 代码模板

	DemoJsonCode *string `json:"DemoJsonCode,omitempty" name:"DemoJsonCode"`
	// 类型映射

	TypeMapSet []*TypeMap `json:"TypeMapSet,omitempty" name:"TypeMapSet"`
	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
}

type AddModuleErrorCodeRequest struct {
	*tchttp.BaseRequest

	// 新增错误码参数

	DataParam *AddModuleErrorCodeParam `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *AddModuleErrorCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddModuleErrorCodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFirstLevelErrorCodeRequest struct {
	*tchttp.BaseRequest
}

func (r *GetFirstLevelErrorCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFirstLevelErrorCodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetObjectHistoryIdsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data []*ObjectHistoryIdParam `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetObjectHistoryIdsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetObjectHistoryIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryApiInfoFilterRequest struct {
	*tchttp.BaseRequest

	// 过滤分页查询字段

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
	// 分页每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页当前页号

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *QueryApiInfoFilterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryApiInfoFilterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetModuleUrlRequest struct {
	*tchttp.BaseRequest

	// 获取单个地址详情

	DataParam *GetModuleUrlParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetModuleUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetModuleUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplexObject struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 模块版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 复杂对象名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 复杂对象id

	ObjectID *int64 `json:"ObjectID,omitempty" name:"ObjectID"`
	// 复杂对象的成员变量

	Members *string `json:"Members,omitempty" name:"Members"`
	// 出参入参

	UsedFor *int64 `json:"UsedFor,omitempty" name:"UsedFor"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建人

	CreateName *string `json:"CreateName,omitempty" name:"CreateName"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 更新人

	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`
	// id

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 是否内置 0:TCE产品; 1:第三方产品

	OwnerTce *int64 `json:"OwnerTce,omitempty" name:"OwnerTce"`
}

type GetModuleNamesRequest struct {
	*tchttp.BaseRequest
}

func (r *GetModuleNamesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetModuleNamesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetModuleUrlsRequest struct {
	*tchttp.BaseRequest

	// 模块

	DataParam *GetModuleUrsParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetModuleUrlsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetModuleUrlsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryApiInfoFilterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的API数据

		Data []*APIDetailParam `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryApiInfoFilterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryApiInfoFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteModuleErrorCodeParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 错误码ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 二级错误码Code

	Code *string `json:"Code,omitempty" name:"Code"`
}

type FilterApiInfoParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 接口名

	Action *string `json:"Action,omitempty" name:"Action"`
	// 接口中文名

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 修改人名

	OperatorName *string `json:"OperatorName,omitempty" name:"OperatorName"`
	// 线上后端地址

	UrlName *string `json:"UrlName,omitempty" name:"UrlName"`
	// 更新起始时间

	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`
	// 更新截止时间

	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`
	// 按某个字段排序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// desc或asc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

type ObjectHistoryIdParam struct {
	// id

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 修改时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 操作类型

	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

type DeleteApiInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApiInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApiInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回

		Data *DescribeActionData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncDocumentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncDocumentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateModuleUrlToActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateModuleUrlToActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateModuleUrlToActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadSdkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// sdk下载地址

		Url *string `json:"Url,omitempty" name:"Url"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadSdkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadSdkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApiInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data *APIDetailParam `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApiInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApiInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddModuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddModuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddModuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteComplexObjectRequest struct {
	*tchttp.BaseRequest

	// 删除复杂对象参数

	DataParam *DeleteComplexObjectParam `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *DeleteComplexObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteComplexObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetObjectHistoryIdsRequest struct {
	*tchttp.BaseRequest

	// 各类型的变更记录ID

	DataParam *GetObjectHistoryIdsParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetObjectHistoryIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetObjectHistoryIdsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunActionData struct {
	// response body

	Body *string `json:"Body,omitempty" name:"Body"`
	// request headers

	Headers *string `json:"Headers,omitempty" name:"Headers"`
	// request curl command

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
}

type GetApiInfoRequest struct {
	*tchttp.BaseRequest

	// 获取API详情参数

	DataParam *GetApiInfoParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetApiInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApiInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetModuleUrlRelateApiResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetModuleUrlRelateApiResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetModuleUrlRelateApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApiInfoRequest struct {
	*tchttp.BaseRequest

	// 更新后的接口信息

	DataParam *ApiInfo `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *UpdateApiInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApiInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Category struct {
	// 官方文档ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 官方文档名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建人

	CreateName *string `json:"CreateName,omitempty" name:"CreateName"`
	// 更新人

	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 是否内置 0:TCE产品; 1:第三方产品

	OwnerTce *int64 `json:"OwnerTce,omitempty" name:"OwnerTce"`
	// 名称国际化

	NameIntlEn *string `json:"NameIntlEn,omitempty" name:"NameIntlEn"`
	// 权重

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type DeleteModuleErrorCodeRequest struct {
	*tchttp.BaseRequest

	// 删除错误码的参数

	DataParam *DeleteModuleErrorCodeParam `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *DeleteModuleErrorCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteModuleErrorCodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCategoryByIDRequest struct {
	*tchttp.BaseRequest

	// 官方文档ID参数

	DataParam *GetCategoryIDParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetCategoryByIDRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCategoryByIDRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSystemCommonParamsRequest struct {
	*tchttp.BaseRequest
}

func (r *GetSystemCommonParamsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSystemCommonParamsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的接口版本

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryModuleUrlFilterRequest struct {
	*tchttp.BaseRequest

	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页号

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryModuleUrlFilterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryModuleUrlFilterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteModuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteModuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteModuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteModuleRequest struct {
	*tchttp.BaseRequest

	// 删除产品模块名

	DataParam *DeleteModuleParam `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *DeleteModuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteModuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteModuleUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteModuleUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteModuleUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateModuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateModuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateModuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddModuleErrorCodeParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 描述信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 一级错误码

	FirstLevelErr *string `json:"FirstLevelErr,omitempty" name:"FirstLevelErr"`
	// 二级错误码

	ModuleErr *string `json:"ModuleErr,omitempty" name:"ModuleErr"`
}

type DescribeActionData struct {
	// 接口文档，markdown格式

	Document *string `json:"Document,omitempty" name:"Document"`
}

type FilterQueryModuleParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 产品描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 产品负责人

	SuperName *string `json:"SuperName,omitempty" name:"SuperName"`
	// 更新起始时间

	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`
	// 更新截止时间

	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// desc或asc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

type AddModuleParam struct {
	// 产品模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 产品描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 产品负责人

	SuperUser *string `json:"SuperUser,omitempty" name:"SuperUser"`
	// 属于tcloud还是ocloud

	Belongs *int64 `json:"Belongs,omitempty" name:"Belongs"`
	// 所属产品

	ProductShortCode *string `json:"ProductShortCode,omitempty" name:"ProductShortCode"`
	// 文档中心关联产品

	DocProductCode *string `json:"DocProductCode,omitempty" name:"DocProductCode"`
}

type DeleteVersionParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 接口版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type GetModuleErrorCodesParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
}

type GetVersionParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
}

type DownloadSdkRequest struct {
	*tchttp.BaseRequest

	// 编程语言，取值go,python,nodejs,php,java

	CodeLanguage *string `json:"CodeLanguage,omitempty" name:"CodeLanguage"`
}

func (r *DownloadSdkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadSdkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComplexObjectByObjIDResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data []*ComplexObject `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComplexObjectByObjIDResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComplexObjectByObjIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddModuleUrlParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 单地址多地址类型

	UrlType *int64 `json:"UrlType,omitempty" name:"UrlType"`
	// 单地址url

	SvrUrl *string `json:"SvrUrl,omitempty" name:"SvrUrl"`
	// 多地址url

	MultiSvrUrl *string `json:"MultiSvrUrl,omitempty" name:"MultiSvrUrl"`
	// 后端地址名

	UrlName *string `json:"UrlName,omitempty" name:"UrlName"`
}

type RunActionRequest struct {
	*tchttp.BaseRequest

	// TargetModule

	TargetModule *string `json:"TargetModule,omitempty" name:"TargetModule"`
	// TargetAction

	TargetAction *string `json:"TargetAction,omitempty" name:"TargetAction"`
	// TargetActionParams

	TargetActionParams *string `json:"TargetActionParams,omitempty" name:"TargetActionParams"`
	// TargetRegion

	TargetRegion *string `json:"TargetRegion,omitempty" name:"TargetRegion"`
	// TargetSecretId

	TargetSecretId *string `json:"TargetSecretId,omitempty" name:"TargetSecretId"`
	// TargetSecretKey

	TargetSecretKey *string `json:"TargetSecretKey,omitempty" name:"TargetSecretKey"`
	// TargetVersion

	TargetVersion *string `json:"TargetVersion,omitempty" name:"TargetVersion"`
}

func (r *RunActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComplexObjectByObjIDRequest struct {
	*tchttp.BaseRequest

	// 复杂对象ID

	DataParam *GetComplexObjectByObjIDParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetComplexObjectByObjIDRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComplexObjectByObjIDRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryModuleFilterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表模块数据

		Data []*ModuleModule `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryModuleFilterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryModuleFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Versions struct {
	// 自增id

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 模块id

	ModuleID *int64 `json:"ModuleID,omitempty" name:"ModuleID"`
	// 接口版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建人

	CreateName *string `json:"CreateName,omitempty" name:"CreateName"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 更新人

	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`
	// 模块

	Module *string `json:"Module,omitempty" name:"Module"`
	// 所属标识 0:TCE产品; 1:第三方产品;

	OwnerTce *int64 `json:"OwnerTce,omitempty" name:"OwnerTce"`
}

type UpdateApiInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApiInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApiInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetModuleUrlParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 后端地址名

	UrlName *string `json:"UrlName,omitempty" name:"UrlName"`
}

type UpdateModuleUrlParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 单、多地址标志

	UrlType *int64 `json:"UrlType,omitempty" name:"UrlType"`
	// 单地址url

	SvrUrl *string `json:"SvrUrl,omitempty" name:"SvrUrl"`
	// 多地址url

	MultiSvrUrl *string `json:"MultiSvrUrl,omitempty" name:"MultiSvrUrl"`
	// 原来的后端地址名

	UrlNameOrigin *string `json:"UrlNameOrigin,omitempty" name:"UrlNameOrigin"`
	// 新后端地址名

	UrlName *string `json:"UrlName,omitempty" name:"UrlName"`
}

type DeleteComplexObjectParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 复杂对象ID

	ObjectID *int64 `json:"ObjectID,omitempty" name:"ObjectID"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 复杂对象名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type UpdateModuleRequest struct {
	*tchttp.BaseRequest

	// 更新的信息

	DataParam *UpdateModuleParam `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *UpdateModuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateModuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FirstLevelErrorcodeParam struct {
	// id

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 一级错误码

	FirstLevelErrorcode *string `json:"FirstLevelErrorcode,omitempty" name:"FirstLevelErrorcode"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建人

	CreateName *string `json:"CreateName,omitempty" name:"CreateName"`
	// 修改时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 修改人

	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`
}

type AddVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDataTypesRequest struct {
	*tchttp.BaseRequest

	// 指定模块名和版本和参数标志

	DataParam *GetDataTypesParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetDataTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDataTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetModuleErrorCodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回模块下的错误码code和message

		Data []*DataGetModuleErrorCode `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetModuleErrorCodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetModuleErrorCodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncDocumentRequest struct {
	*tchttp.BaseRequest

	// 同步指定服务（module）的文档

	Modules []*string `json:"Modules,omitempty" name:"Modules"`
	// 同步指定产品的文档

	DocProductCodes []*string `json:"DocProductCodes,omitempty" name:"DocProductCodes"`
}

func (r *SyncDocumentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncDocumentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComplexObjectRelateApiParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 复杂对象ID

	ObjectId *int64 `json:"ObjectId,omitempty" name:"ObjectId"`
	// 复杂对象名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type GetModuleErrorCodesRequest struct {
	*tchttp.BaseRequest

	// 指定模块名

	DataParam *GetModuleErrorCodesParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetModuleErrorCodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetModuleErrorCodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSystemCommonParamsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 系统公共参数列表

		SystemParam []*SystemParam `json:"SystemParam,omitempty" name:"SystemParam"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSystemCommonParamsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSystemCommonParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetModuleErrorCodeRelateApiRequest struct {
	*tchttp.BaseRequest

	// 错误码

	DataParam *GetModuleErrorCodeRelateApiParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetModuleErrorCodeRelateApiRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetModuleErrorCodeRelateApiRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataGetCategory struct {
	// 文档分类ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 文档分类名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DeleteModuleParam struct {
	// 模块产品名

	Module *string `json:"Module,omitempty" name:"Module"`
}

type AddComplexObjectRequest struct {
	*tchttp.BaseRequest

	// 新增复杂对象参数

	DataParam *AddComplexObjectParam `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *AddComplexObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddComplexObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeActionRequest struct {
	*tchttp.BaseRequest

	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 接口名称

	ProductAction *string `json:"ProductAction,omitempty" name:"ProductAction"`
	// 接口版本

	ProductVersion *string `json:"ProductVersion,omitempty" name:"ProductVersion"`
}

func (r *DescribeActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetObjectHistoryIdsParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 类型名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 类型

	ObjectType *string `json:"ObjectType,omitempty" name:"ObjectType"`
}

type ObjectHistoryParam struct {
	// 模块

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// id

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 修改人

	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`
	// 修改时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 操作类型

	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

type DeleteApiInfoParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 接口名

	Action *string `json:"Action,omitempty" name:"Action"`
}

type AddApiInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddApiInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddApiInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data []*GetRegionsParam `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCategoryParam struct {
	// 官方文档ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 官方文档名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type UpdateModuleErrorCodeParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 新的描述

	Message *string `json:"Message,omitempty" name:"Message"`
	// 一级错误码

	FirstLevelErr *string `json:"FirstLevelErr,omitempty" name:"FirstLevelErr"`
	// 旧的二级错误码

	ModuleErrOgigin *string `json:"ModuleErrOgigin,omitempty" name:"ModuleErrOgigin"`
	// 新的二级错误码

	ModuleErr *string `json:"ModuleErr,omitempty" name:"ModuleErr"`
}

type QueryComplexObjectFilterRequest struct {
	*tchttp.BaseRequest

	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页号

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤字段

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryComplexObjectFilterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryComplexObjectFilterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiInfo struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 接口名

	Action *string `json:"Action,omitempty" name:"Action"`
	// 白名单策略

	WhiteList *string `json:"WhiteList,omitempty" name:"WhiteList"`
	// timeout超时时间

	TimeoutInMillisec *int64 `json:"TimeoutInMillisec,omitempty" name:"TimeoutInMillisec"`
	// 入参

	InputInfo *string `json:"InputInfo,omitempty" name:"InputInfo"`
	// 系统公共参数

	SystemParam []*SystemParam `json:"SystemParam,omitempty" name:"SystemParam"`
	// 标志

	Flag *int64 `json:"Flag,omitempty" name:"Flag"`
	// 单地址URL

	SvrUrl *string `json:"SvrUrl,omitempty" name:"SvrUrl"`
	// 接口中文名

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 接口英文名称

	ActionNameIntlEnUs *string `json:"ActionNameIntlEnUs,omitempty" name:"ActionNameIntlEnUs"`
	// 接口英文描述

	DescriptionIntlEnUs *string `json:"DescriptionIntlEnUs,omitempty" name:"DescriptionIntlEnUs"`
	// CAM资源

	CamResource *string `json:"CamResource,omitempty" name:"CamResource"`
	// 接口对外是否可见

	CliVisibility *int64 `json:"CliVisibility,omitempty" name:"CliVisibility"`
	// CAM条件

	CamCondition *string `json:"CamCondition,omitempty" name:"CamCondition"`
	// 限频

	SecRate *int64 `json:"SecRate,omitempty" name:"SecRate"`
	// 限频白名单

	RateLev *string `json:"RateLev,omitempty" name:"RateLev"`
	// 签名和鉴权

	IsAuth *int64 `json:"IsAuth,omitempty" name:"IsAuth"`
	// 后端地址名

	UrlName *string `json:"UrlName,omitempty" name:"UrlName"`
	// 全程票据校验

	WholeProcessToken *int64 `json:"WholeProcessToken,omitempty" name:"WholeProcessToken"`
	// 输出参数

	OutputInfo *string `json:"OutputInfo,omitempty" name:"OutputInfo"`
	// 错误码

	Errorcodes *string `json:"Errorcodes,omitempty" name:"Errorcodes"`
	// 接口描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 接口文档分类

	Category *string `json:"Category,omitempty" name:"Category"`
	// 多地址类型

	MultiSvrUrl *string `json:"MultiSvrUrl,omitempty" name:"MultiSvrUrl"`
	// 通过 HTTP Header 传递的业务参数列表

	HeaderParam []*string `json:"HeaderParam,omitempty" name:"HeaderParam"`
	// 请求体大小, 单位KB. ( 默认: 1M, 最大:10MB)

	ReqSize *int64 `json:"ReqSize,omitempty" name:"ReqSize"`
	// 业务后端交互协议

	BackendProtocolSupported *string `json:"BackendProtocolSupported,omitempty" name:"BackendProtocolSupported"`
}

type AddModuleRequest struct {
	*tchttp.BaseRequest

	// 模块信息

	DataParam *AddModuleParam `json:"DataParam,omitempty" name:"DataParam"`
	// cam资源鉴权

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *AddModuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddModuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回模块下文档分类ID和名

		Data []*DataGetCategory `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddVersionRequest struct {
	*tchttp.BaseRequest

	// 接口版本参数

	DataParam *AddVersionParam `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *AddVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateModuleUrlRequest struct {
	*tchttp.BaseRequest

	// 更新后端地址

	DataParam *UpdateModuleUrlParam `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *UpdateModuleUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateModuleUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterQueryModuleErrorCodeParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 创建人

	CreateName *string `json:"CreateName,omitempty" name:"CreateName"`
	// 二级错误码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 更新起始时间

	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`
	// 更新截止时间

	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// desc或asc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

type ResApiInfo struct {
	// 接口详细info

	ActionInfo *ActionInfo `json:"ActionInfo,omitempty" name:"ActionInfo"`
	// 接口精简描述信息

	ActionDesc *ActionDescSample `json:"ActionDesc,omitempty" name:"ActionDesc"`
	// 白名单

	WhiteList *string `json:"WhiteList,omitempty" name:"WhiteList"`
	// 接口超时时间

	TimeoutInMillisec *int64 `json:"TimeoutInMillisec,omitempty" name:"TimeoutInMillisec"`
}

type DeleteComplexObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteComplexObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteComplexObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddApiInfoRequest struct {
	*tchttp.BaseRequest

	// 接口信息

	DataParam *ApiInfo `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *AddApiInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddApiInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetEnvironmentRequest struct {
	*tchttp.BaseRequest
}

func (r *GetEnvironmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetEnvironmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComplexObjectRequest struct {
	*tchttp.BaseRequest

	// 指定模块名和版本和复杂对象名

	DataParam *GetComplexObjectParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetComplexObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComplexObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCategoryFilterRequest struct {
	*tchttp.BaseRequest

	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页号

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤字段参数

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryCategoryFilterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCategoryFilterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFirstLevelErrorCodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 一级错误码

		Data []*FirstLevelErrorcodeParam `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFirstLevelErrorCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFirstLevelErrorCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComplexObjectRelateApiRequest struct {
	*tchttp.BaseRequest

	// 复杂对象

	DataParam *GetComplexObjectRelateApiParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetComplexObjectRelateApiRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComplexObjectRelateApiRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionInfoSample struct {
	// 入参

	InputInfo *string `json:"InputInfo,omitempty" name:"InputInfo"`
	// 系统公共参数

	SystemParam []*SystemParam `json:"SystemParam,omitempty" name:"SystemParam"`
	// 标志

	Flag *int64 `json:"Flag,omitempty" name:"Flag"`
	// 单地址URL

	SvrUrl *string `json:"SvrUrl,omitempty" name:"SvrUrl"`
	// 接口中文名

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// CAM资源

	CamResource *string `json:"CamResource,omitempty" name:"CamResource"`
	// 接口对外是否可见

	CliVisibility *int64 `json:"CliVisibility,omitempty" name:"CliVisibility"`
	// CAM条件

	CamCondition *string `json:"CamCondition,omitempty" name:"CamCondition"`
	// 限频

	SecRate *int64 `json:"SecRate,omitempty" name:"SecRate"`
	// 限频白名单

	RateLev *string `json:"RateLev,omitempty" name:"RateLev"`
	// 签名和鉴权

	IsAuth *int64 `json:"IsAuth,omitempty" name:"IsAuth"`
	// 后端地址名

	UrlName *string `json:"UrlName,omitempty" name:"UrlName"`
	// 全程票据校验

	WholeProcessToken *int64 `json:"WholeProcessToken,omitempty" name:"WholeProcessToken"`
}

type ActionDescSample struct {
	// 输出参数

	OutputInfo *string `json:"OutputInfo,omitempty" name:"OutputInfo"`
	// 错误码

	Errorcodes *string `json:"Errorcodes,omitempty" name:"Errorcodes"`
	// 接口描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 接口文档分类

	Category *string `json:"Category,omitempty" name:"Category"`
}

type DescribeSdkDemosData struct {
	// 代码模板列表

	SdkDemoSet []*SdkDemoSetInfo `json:"SdkDemoSet,omitempty" name:"SdkDemoSet"`
}

type DeleteVersionRequest struct {
	*tchttp.BaseRequest

	// 指定模块名和版本

	DataParam *DeleteVersionParam `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *DeleteVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteModuleUrlRequest struct {
	*tchttp.BaseRequest

	// 后端地址

	DataParam *DeleteModuleUrlParam `json:"DataParam,omitempty" name:"DataParam"`
	// 模块名

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *DeleteModuleUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteModuleUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCategoryByIDResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的官方文档详细信息

		Data *Category `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCategoryByIDResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCategoryByIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApiInfoRequest struct {
	*tchttp.BaseRequest

	// 删除接口信息

	DataParam *DeleteApiInfoParam `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *DeleteApiInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApiInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteModuleErrorCodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteModuleErrorCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteModuleErrorCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApiInfoHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data *GetApiInfoHistoryParam `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApiInfoHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApiInfoHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCategoryRequest struct {
	*tchttp.BaseRequest

	// 指定文档分类的模块和版本

	DataParam *GetCategoryParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCategoryRequest struct {
	*tchttp.BaseRequest

	// 删除官方文档分类参数

	DataParam *DeleteCategoryParam `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *DeleteCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteModuleUrlParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 后端地址名

	UrlName *string `json:"UrlName,omitempty" name:"UrlName"`
}

type GetDataTypesParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 入参出参

	UsedFor *int64 `json:"UsedFor,omitempty" name:"UsedFor"`
}

type GetApiInfoHistoryRequest struct {
	*tchttp.BaseRequest

	// 获取接口历史

	DataParam *GetApiInfoHistoryParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetApiInfoHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApiInfoHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetModuleUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的后端地址

		Data []*ModuleUrl `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetModuleUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetModuleUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateModuleUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateModuleUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateModuleUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryModuleErrorCodeFilterRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页号

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *QueryModuleErrorCodeFilterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryModuleErrorCodeFilterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryObjectHistoryFilterRequest struct {
	*tchttp.BaseRequest

	// 页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤字段

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
	// 类型说明

	ObjectType *string `json:"ObjectType,omitempty" name:"ObjectType"`
}

func (r *QueryObjectHistoryFilterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryObjectHistoryFilterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type APIDetailParam struct {
	// 白名单

	WhiteList *string `json:"WhiteList,omitempty" name:"WhiteList"`
	// 超时毫秒数

	TimeoutInMillisec *int64 `json:"TimeoutInMillisec,omitempty" name:"TimeoutInMillisec"`
	// 模块

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 接口

	Action *string `json:"Action,omitempty" name:"Action"`
	// 入参

	InputInfo *string `json:"InputInfo,omitempty" name:"InputInfo"`
	// 系统公共参数

	SystemParam []*SystemParam `json:"SystemParam,omitempty" name:"SystemParam"`
	// 标志

	Flag *int64 `json:"Flag,omitempty" name:"Flag"`
	// 单地址URL

	SvrUrl *string `json:"SvrUrl,omitempty" name:"SvrUrl"`
	// 多地址类型

	MultiSvrUrl *string `json:"MultiSvrUrl,omitempty" name:"MultiSvrUrl"`
	// 接口名

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 接口地址

	ApiUrl *string `json:"ApiUrl,omitempty" name:"ApiUrl"`
	// 状态

	RecordStatus *int64 `json:"RecordStatus,omitempty" name:"RecordStatus"`
	// CAM资源

	CamResource *string `json:"CamResource,omitempty" name:"CamResource"`
	// 接口标签

	ActionLabel *string `json:"ActionLabel,omitempty" name:"ActionLabel"`
	// 接口对外是否可见

	CliVisibility *int64 `json:"CliVisibility,omitempty" name:"CliVisibility"`
	// CAM条件

	CamCondition *string `json:"CamCondition,omitempty" name:"CamCondition"`
	// 限频

	SecRate *int64 `json:"SecRate,omitempty" name:"SecRate"`
	// 限频白名单

	RateLev *string `json:"RateLev,omitempty" name:"RateLev"`
	// 签名和鉴权

	IsAuth *int64 `json:"IsAuth,omitempty" name:"IsAuth"`
	// 后端地址名

	UrlName *string `json:"UrlName,omitempty" name:"UrlName"`
	// 全程票据校验

	WholeProcessToken *int64 `json:"WholeProcessToken,omitempty" name:"WholeProcessToken"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建人

	CreateName *string `json:"CreateName,omitempty" name:"CreateName"`
	// 修改时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 修改人

	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`
	// 业务后端交互协议

	BackendProtocolSupported *string `json:"BackendProtocolSupported,omitempty" name:"BackendProtocolSupported"`
	// 通过 HTTP Header 传递的业务参数列表

	HeaderParam []*string `json:"HeaderParam,omitempty" name:"HeaderParam"`
	// 请求体大小, 单位KB. ( 默认: 1M, 最大:10MB)

	ReqSize *int64 `json:"ReqSize,omitempty" name:"ReqSize"`
	// 是否内置 0:TCE产品; 1:第三方产品.

	OwnerTce *int64 `json:"OwnerTce,omitempty" name:"OwnerTce"`
	// 所属产品

	ProductShortCode *string `json:"ProductShortCode,omitempty" name:"ProductShortCode"`
	// 输出参数

	OutputInfo *string `json:"OutputInfo,omitempty" name:"OutputInfo"`
	// 错误码

	Errorcodes *string `json:"Errorcodes,omitempty" name:"Errorcodes"`
	// 接口描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 接口文档分类

	Category *string `json:"Category,omitempty" name:"Category"`
	// 接口文档分类名称

	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`
}

type AddModuleErrorCodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddModuleErrorCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddModuleErrorCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComplexObjectRelateApiResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的接口

		Data *InOutPutInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComplexObjectRelateApiResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComplexObjectRelateApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetModuleUrlsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data []*ModuleUrl `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetModuleUrlsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetModuleUrlsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateModuleErrorCodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateModuleErrorCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateModuleErrorCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSdkDemosRequest struct {
	*tchttp.BaseRequest

	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 接口名称

	ProductAction *string `json:"ProductAction,omitempty" name:"ProductAction"`
	// 接口版本

	ProductVersion *string `json:"ProductVersion,omitempty" name:"ProductVersion"`
}

func (r *DescribeSdkDemosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSdkDemosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryModuleUrlFilterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的列表后端地址

		Data []*ModuleUrl `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryModuleUrlFilterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryModuleUrlFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetModuleUrlRelateApiRequest struct {
	*tchttp.BaseRequest

	// 后端地址

	DataParam *GetModuleUrlParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetModuleUrlRelateApiRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetModuleUrlRelateApiRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComplexObjectByObjIDParam struct {
	// 数组ID

	ObjectIDs []*int64 `json:"ObjectIDs,omitempty" name:"ObjectIDs"`
}

type QueryCategoryFilterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的官方文档数据

		Data []*Category `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCategoryFilterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCategoryFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryVersionFilterRequest struct {
	*tchttp.BaseRequest

	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页号

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤字段

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryVersionFilterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryVersionFilterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModuleUrl struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 单多地址标志

	UrlType *int64 `json:"UrlType,omitempty" name:"UrlType"`
	// 单地址url

	SvrUrl *string `json:"SvrUrl,omitempty" name:"SvrUrl"`
	// 多地址url

	MultiSvrUrl *string `json:"MultiSvrUrl,omitempty" name:"MultiSvrUrl"`
	// 后端地址名

	UrlName *string `json:"UrlName,omitempty" name:"UrlName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建人

	CreateName *string `json:"CreateName,omitempty" name:"CreateName"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 更新人

	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`
	// 是否编码

	IsCode *int64 `json:"IsCode,omitempty" name:"IsCode"`
	// 是否删除

	IsDeleted *int64 `json:"IsDeleted,omitempty" name:"IsDeleted"`
	// 是否内置 0:TCE产品; 1:第三方产品

	OwnerTce *int64 `json:"OwnerTce,omitempty" name:"OwnerTce"`
}

type AddModuleUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddModuleUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddModuleUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSdkDemosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回

		Data *DescribeSdkDemosData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSdkDemosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSdkDemosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterString struct {
	// key过滤字段

	Name *string `json:"Name,omitempty" name:"Name"`
	// value过滤字段的值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type UpdateCategoryParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 原来文档分类名

	NameOrigin *string `json:"NameOrigin,omitempty" name:"NameOrigin"`
	// 新文档分类名

	NameNew *string `json:"NameNew,omitempty" name:"NameNew"`
}

type GetCategoryParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type GetComplexObjectParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 复杂对象名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type GetModuleUrsParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
}

type GetDataTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回模块下数据类型数组

		Data *DataType `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDataTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDataTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterComplexObjectParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 复杂对象名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 复杂对象描述名

	Description *string `json:"Description,omitempty" name:"Description"`
	// 更新人

	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`
	// 复杂对象版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 更新起始时间

	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`
	// 更新截止时间

	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// desc或asc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

type TestTest struct {
	// gdsfgd

	DataInfo *int64 `json:"DataInfo,omitempty" name:"DataInfo"`
}

type GetCategorysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 类别数据

		Data []*DataGetCategory `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCategorysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCategorysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryObjectHistoryFilterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据

		Data []*ObjectHistoryParam `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryObjectHistoryFilterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryObjectHistoryFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModuleErrorcode struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 错误码描述

	Message *string `json:"Message,omitempty" name:"Message"`
	// 创建人

	CreateName *string `json:"CreateName,omitempty" name:"CreateName"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新人

	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`
	// 二级错误码

	Code *string `json:"Code,omitempty" name:"Code"`
	// id

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 是否内置 0:TCE产品; 1:第三方产品;

	OwnerTce *int64 `json:"OwnerTce,omitempty" name:"OwnerTce"`
}

type DataGetModuleErrorCode struct {
	// 二级错误码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 错误码描述

	Message *string `json:"Message,omitempty" name:"Message"`
	// ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
}

type GetModuleErrorCodeRelateApiParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 错误码ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
}

type UpdateModuleParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 产品描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 产品负责人

	SuperUser *string `json:"SuperUser,omitempty" name:"SuperUser"`
	// 所属产品

	ProductShortCode *string `json:"ProductShortCode,omitempty" name:"ProductShortCode"`
	// 文档中心关联产品

	DocProductCode *string `json:"DocProductCode,omitempty" name:"DocProductCode"`
}

type AddCategoryRequest struct {
	*tchttp.BaseRequest

	// 新增官方文档分类参数

	DataParam *AddCategoryParam `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *AddCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddModuleUrlRequest struct {
	*tchttp.BaseRequest

	// 新增后端地址参数

	DataParam *AddModuleUrlParam `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *AddModuleUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddModuleUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionInfo struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 接口版本号

	Version *string `json:"Version,omitempty" name:"Version"`
	// 接口名

	Action *string `json:"Action,omitempty" name:"Action"`
	// 入参

	InputInfo *string `json:"InputInfo,omitempty" name:"InputInfo"`
	// 系统参数

	SystemParam []*SystemParam `json:"SystemParam,omitempty" name:"SystemParam"`
	// attribute_flag

	Flag *int64 `json:"Flag,omitempty" name:"Flag"`
	// 后端地址单地址

	SvrUrl *string `json:"SvrUrl,omitempty" name:"SvrUrl"`
	// 后端地址多地址

	MultiSvrUrl *string `json:"MultiSvrUrl,omitempty" name:"MultiSvrUrl"`
	// 接口中文名

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 发布状态

	RecordStatus *int64 `json:"RecordStatus,omitempty" name:"RecordStatus"`
	// CAM资源

	CamResource *string `json:"CamResource,omitempty" name:"CamResource"`
	// 接口对外是否可见

	CliVisibility *int64 `json:"CliVisibility,omitempty" name:"CliVisibility"`
	// CAM条件

	CamCondition *string `json:"CamCondition,omitempty" name:"CamCondition"`
	// 限频

	SecRate *int64 `json:"SecRate,omitempty" name:"SecRate"`
	// 限频白名单

	RateLev *string `json:"RateLev,omitempty" name:"RateLev"`
	// 签名和鉴权

	IsAuth *int64 `json:"IsAuth,omitempty" name:"IsAuth"`
	// 后端地址

	UrlName *string `json:"UrlName,omitempty" name:"UrlName"`
	// 全程票据校验

	WholeProcessToken *int64 `json:"WholeProcessToken,omitempty" name:"WholeProcessToken"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建人

	CreateName *string `json:"CreateName,omitempty" name:"CreateName"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 更新人

	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`
	// 更新时间

	UpdateAt *string `json:"UpdateAt,omitempty" name:"UpdateAt"`
}

type UpdateModuleErrorCodeRequest struct {
	*tchttp.BaseRequest

	// 更新错误码的参数

	DataParam *UpdateModuleErrorCodeParam `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *UpdateModuleErrorCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateModuleErrorCodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApiInfoHistoryParam struct {
	// ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 复杂对象名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 变更历史数据

	Data *YunapiAPIInfoHistory `json:"Data,omitempty" name:"Data"`
}

type Position struct {
	// 通过 HTTP JSON Body 传递

	Body *bool `json:"Body,omitempty" name:"Body"`
	// 通过 HTTP Header 传递

	Header *bool `json:"Header,omitempty" name:"Header"`
}

type UpdateModuleUrlToActionParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 接口名

	Action *string `json:"Action,omitempty" name:"Action"`
	// 单地址URL

	SvrUrl *string `json:"SvrUrl,omitempty" name:"SvrUrl"`
	// 多地址URL

	MultiSvrUrl *string `json:"MultiSvrUrl,omitempty" name:"MultiSvrUrl"`
	// 后端地址名

	UrlName *string `json:"UrlName,omitempty" name:"UrlName"`
}

type GetModuleNamesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模块数据

		Data []*ModuleModule `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetModuleNamesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetModuleNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TypeMap struct {
	// 基本类型

	OriginalType *string `json:"OriginalType,omitempty" name:"OriginalType"`
	// 映射类型

	DestinationType *string `json:"DestinationType,omitempty" name:"DestinationType"`
	// sdk中类型表示方式

	SdkType *string `json:"SdkType,omitempty" name:"SdkType"`
	// 后缀

	Postfix *string `json:"Postfix,omitempty" name:"Postfix"`
}

type QueryVersionFilterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的接口版本数据

		Data []*Versions `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryVersionFilterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryVersionFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModuleModule struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 权限人

	PermissionUser *string `json:"PermissionUser,omitempty" name:"PermissionUser"`
	// 负责人

	SuperUser *string `json:"SuperUser,omitempty" name:"SuperUser"`
	// url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 行业分类

	CategoryID *int64 `json:"CategoryID,omitempty" name:"CategoryID"`
	// 所属产品

	ProductShortCode *string `json:"ProductShortCode,omitempty" name:"ProductShortCode"`
	// ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 旧文档产品code

	OldDocProductCode *string `json:"OldDocProductCode,omitempty" name:"OldDocProductCode"`
	// 所属标识 0:TCE产品; 1:第三方产品;

	OwnerTce *int64 `json:"OwnerTce,omitempty" name:"OwnerTce"`
	// cam策略id

	CamPolicyID *int64 `json:"CamPolicyID,omitempty" name:"CamPolicyID"`
	// cam策略名称

	CamPolicyName *string `json:"CamPolicyName,omitempty" name:"CamPolicyName"`
	// 文档产品code

	DocProductCode *string `json:"DocProductCode,omitempty" name:"DocProductCode"`
}

type GetRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *GetRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddCategoryParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 分类名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type InOutPutInfo struct {
	// 入参接口

	ActionsInput []*string `json:"ActionsInput,omitempty" name:"ActionsInput"`
	// 出参接口

	ActionsOutput []*string `json:"ActionsOutput,omitempty" name:"ActionsOutput"`
}

type DataType struct {
	// 数据类型名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 数据类型说明

	Message *string `json:"Message,omitempty" name:"Message"`
	// 复杂对象

	ComplexObjectTypes []*ComplexObject `json:"ComplexObjectTypes,omitempty" name:"ComplexObjectTypes"`
	// 基础类型

	BasicTypes []*BasicTypeParam `json:"BasicTypes,omitempty" name:"BasicTypes"`
}

type QueryModuleFilterRequest struct {
	*tchttp.BaseRequest

	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页号

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤字段

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryModuleFilterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryModuleFilterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateModuleUrlToActionRequest struct {
	*tchttp.BaseRequest

	// 指定接口和后端地址

	DataParam *UpdateModuleUrlToActionParam `json:"DataParam,omitempty" name:"DataParam"`
	// cam鉴权资源条件

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *UpdateModuleUrlToActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateModuleUrlToActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BasicTypeParam struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 对象类型

	ObjectType *string `json:"ObjectType,omitempty" name:"ObjectType"`
}

type YunapiAPIInfoHistory struct {
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 模块

	Module *string `json:"Module,omitempty" name:"Module"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 接口

	Action *string `json:"Action,omitempty" name:"Action"`
	// 入参

	InputInfo *string `json:"InputInfo,omitempty" name:"InputInfo"`
	// 系统参数

	SystemParam *string `json:"SystemParam,omitempty" name:"SystemParam"`
	// 属性标志

	AttributeFlag *uint64 `json:"AttributeFlag,omitempty" name:"AttributeFlag"`
	// 单地址URL

	SvrUrl *string `json:"SvrUrl,omitempty" name:"SvrUrl"`
	// 多地址类型

	MultiSvrUrl *string `json:"MultiSvrUrl,omitempty" name:"MultiSvrUrl"`
	// 接口名

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 接口地址

	ApiUrl *string `json:"ApiUrl,omitempty" name:"ApiUrl"`
	// 变更时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 操作人

	OperatorName *string `json:"OperatorName,omitempty" name:"OperatorName"`
	// 状态

	RecordStatus *int64 `json:"RecordStatus,omitempty" name:"RecordStatus"`
	// CAM资源

	CamResource *string `json:"CamResource,omitempty" name:"CamResource"`
	// 接口标志

	ActionLabel *string `json:"ActionLabel,omitempty" name:"ActionLabel"`
	// 修改人

	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`
	// 接口对外是否可见

	CliVisibility *int64 `json:"CliVisibility,omitempty" name:"CliVisibility"`
	// CAM条件

	CamCondition *string `json:"CamCondition,omitempty" name:"CamCondition"`
	// 限频

	SecRate *int64 `json:"SecRate,omitempty" name:"SecRate"`
	// 限频白名单

	RateLev *string `json:"RateLev,omitempty" name:"RateLev"`
	// 签名和鉴权

	IsAuth *int64 `json:"IsAuth,omitempty" name:"IsAuth"`
	// 后端地址名

	UrlName *string `json:"UrlName,omitempty" name:"UrlName"`
	// 修改时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 业务后端交互协议

	BackendProtocolSupported *string `json:"BackendProtocolSupported,omitempty" name:"BackendProtocolSupported"`
	// 全程票据校验

	WholeProcessToken *int64 `json:"WholeProcessToken,omitempty" name:"WholeProcessToken"`
	// 出参

	OutputInfo *string `json:"OutputInfo,omitempty" name:"OutputInfo"`
	// 错误码

	Errorcodes *string `json:"Errorcodes,omitempty" name:"Errorcodes"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 类别

	Category *string `json:"Category,omitempty" name:"Category"`
	// 白名单

	WhiteList *string `json:"WhiteList,omitempty" name:"WhiteList"`
	// 超时毫秒数

	TimeoutInMillisec *int64 `json:"TimeoutInMillisec,omitempty" name:"TimeoutInMillisec"`
	// 操作类型

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 系统参数详情

	SystemParamDetail *string `json:"SystemParamDetail,omitempty" name:"SystemParamDetail"`
	// 通过 HTTP Header 传递的业务参数列表

	HeaderParam *string `json:"HeaderParam,omitempty" name:"HeaderParam"`
	// 请求体大小, 单位KB. ( 默认: 1M, 最大:10MB)

	ReqSize *int64 `json:"ReqSize,omitempty" name:"ReqSize"`
}

type GetCategorysRequest struct {
	*tchttp.BaseRequest

	// 参数

	DataParam *GetCategoryParam `json:"DataParam,omitempty" name:"DataParam"`
}

func (r *GetCategorysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCategorysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterQueryModuleUrlParam struct {
	// 模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 更新人

	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`
	// 创建人

	CreateName *string `json:"CreateName,omitempty" name:"CreateName"`
	// 后端地址名

	UrlName *string `json:"UrlName,omitempty" name:"UrlName"`
	// 更新起始时间

	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`
	// 更新截止时间

	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// asc或desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}
