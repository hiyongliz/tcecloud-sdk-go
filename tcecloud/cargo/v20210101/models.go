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

package v20210101

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type GetAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户名

		User *string `json:"User,omitempty" name:"User"`
		// 访问仓库的 Token

		Token *string `json:"Token,omitempty" name:"Token"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListImageBuildingTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的总量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 构建任务数据集

		ImageBuildingTaskSet []*ImageBuildingTask `json:"ImageBuildingTaskSet,omitempty" name:"ImageBuildingTaskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListImageBuildingTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListImageBuildingTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Schedule struct {
	// GC类型，Manual或者定期时间

	Type *string `json:"Type,omitempty" name:"Type"`
	// cron表达式

	Cron *string `json:"Cron,omitempty" name:"Cron"`
}

type BlockProjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BlockProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Settings struct {
	// cron时间周期

	Cron *string `json:"Cron,omitempty" name:"Cron"`
}

type Scope struct {
	// 关联资源

	Level *string `json:"Level,omitempty" name:"Level"`
	// 关联资源ID

	Ref *uint64 `json:"Ref,omitempty" name:"Ref"`
}

type ScopeSelectors struct {
	// repo匹配规则

	Repository []*ScopeSelectorsRepository `json:"Repository,omitempty" name:"Repository"`
}

type DescribeTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配的 Tag 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回的 Tag 列表

		Tags []*ImageTag `json:"Tags,omitempty" name:"Tags"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGCHistoryRequest struct {
	*tchttp.BaseRequest

	// gc id

	GcID *uint64 `json:"GcID,omitempty" name:"GcID"`
}

func (r *GetGCHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGCHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageScanningStatusRequest struct {
	*tchttp.BaseRequest

	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像的摘要 ID

	Digest *string `json:"Digest,omitempty" name:"Digest"`
}

func (r *ImageScanningStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImageScanningStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGCScheduleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGCScheduleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateGCScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGCScheduleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGCScheduleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGCScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageBuildingTaskRequest struct {
	*tchttp.BaseRequest

	// 待查询的任务, 需指定Name

	ImageBuildingTask *ImageBuildingTask `json:"ImageBuildingTask,omitempty" name:"ImageBuildingTask"`
}

func (r *DescribeImageBuildingTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageBuildingTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulnerabilitiesRequest struct {
	*tchttp.BaseRequest

	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像的摘要

	Digest *string `json:"Digest,omitempty" name:"Digest"`
}

func (r *DescribeVulnerabilitiesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulnerabilitiesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListImageBuildingTasksRequest struct {
	*tchttp.BaseRequest

	// 查询过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 倒序排序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 查询数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListImageBuildingTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListImageBuildingTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询项目成员总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 项目成员信息

		Members []*Member `json:"Members,omitempty" name:"Members"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProjectMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRetentionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRetentionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRetentionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRetentionRequest struct {
	*tchttp.BaseRequest

	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 策略ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 算法

	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
	// 规则

	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`
	// 触发条件

	Trigger *Trigger `json:"Trigger,omitempty" name:"Trigger"`
	// 相关资源信息

	Scope *Scope `json:"Scope,omitempty" name:"Scope"`
}

func (r *UpdateRetentionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRetentionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageBuildingTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageBuildingTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageBuildingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageBuildingTaskRequest struct {
	*tchttp.BaseRequest

	// 待删除的任务, 需要设置Name

	ImageBuildingTask *ImageBuildingTask `json:"ImageBuildingTask,omitempty" name:"ImageBuildingTask"`
}

func (r *DeleteImageBuildingTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageBuildingTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CleanImagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行的ID

		ExecutionID *uint64 `json:"ExecutionID,omitempty" name:"ExecutionID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CleanImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CleanImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyImagesRequest struct {
	*tchttp.BaseRequest

	// 镜像的源项目

	SourceProject *string `json:"SourceProject,omitempty" name:"SourceProject"`
	// 镜像复制的目标项目

	TargetProject *string `json:"TargetProject,omitempty" name:"TargetProject"`
	// 复制的镜像名称

	Image *string `json:"Image,omitempty" name:"Image"`
	// 复制的镜像 Tag 列表

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

func (r *CopyImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyImagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Parameters struct {
	// 是否清理无tag镜像

	DeleteUntagged *bool `json:"DeleteUntagged,omitempty" name:"DeleteUntagged"`
	// 是否模拟运行

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

type GetRetentionExecutionLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志信息

		Log *string `json:"Log,omitempty" name:"Log"`
		// 日志总量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRetentionExecutionLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRetentionExecutionLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetTokenRequest struct {
	*tchttp.BaseRequest

	// 用户名

	User *string `json:"User,omitempty" name:"User"`
}

func (r *ResetTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchImagesRequest struct {
	*tchttp.BaseRequest

	// 镜像名称，用于模糊匹配。匹配 projectName/imageName

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
}

func (r *SearchImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchImagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTagRequest struct {
	*tchttp.BaseRequest

	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// Tag 名称

	TagName *string `json:"TagName,omitempty" name:"TagName"`
}

func (r *DeleteTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetGCStatusRequest struct {
	*tchttp.BaseRequest

	// 打开/关闭GC

	GcOn *bool `json:"GcOn,omitempty" name:"GcOn"`
}

func (r *SetGCStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetGCStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartImageBuildingTaskRequest struct {
	*tchttp.BaseRequest

	// 任务名, 需要设置Name

	ImageBuildingTask *ImageBuildingTask `json:"ImageBuildingTask,omitempty" name:"ImageBuildingTask"`
}

func (r *StartImageBuildingTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartImageBuildingTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Project struct {
	// 项目 ID

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 项目名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 项目别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 项目的描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 是否是公共项目

	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`
	// 项目下的镜像数

	RepoCount *uint64 `json:"RepoCount,omitempty" name:"RepoCount"`
	// 是否开启了自动漏洞扫描

	AutoScan *bool `json:"AutoScan,omitempty" name:"AutoScan"`
	// 项目的创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
	// 项目的更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 项目是否锁定

	IsBlocked *bool `json:"IsBlocked,omitempty" name:"IsBlocked"`
	// 项目配额

	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`
	// 镜像清理策略ID

	RetentionID *uint64 `json:"RetentionID,omitempty" name:"RetentionID"`
	// 项目使用量

	StorageUsed *int64 `json:"StorageUsed,omitempty" name:"StorageUsed"`
	// 权限ID；1管理，3只读

	RoleID *int64 `json:"RoleID,omitempty" name:"RoleID"`
}

type DescribeVulnerabilitiesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞扫描的时间

		GeneratedAt *string `json:"GeneratedAt,omitempty" name:"GeneratedAt"`
		// 漏洞扫描器

		Scanner *Scanner `json:"Scanner,omitempty" name:"Scanner"`
		// 漏洞的严重程度

		Severity *string `json:"Severity,omitempty" name:"Severity"`
		// 漏洞列表

		Vulnerabilities []*Vulnerability `json:"Vulnerabilities,omitempty" name:"Vulnerabilities"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulnerabilitiesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulnerabilitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRetentionExecutionLogRequest struct {
	*tchttp.BaseRequest

	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 策略ID

	RetentionID *uint64 `json:"RetentionID,omitempty" name:"RetentionID"`
	// 策略执行ID

	ExecutionID *uint64 `json:"ExecutionID,omitempty" name:"ExecutionID"`
	// 获取第几页

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页记录数

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *GetRetentionExecutionLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRetentionExecutionLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegistryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegistryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegistryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGCScheduleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 时间表配置

		Schedule *Schedule `json:"Schedule,omitempty" name:"Schedule"`
		// 参数配置

		JobParameters []*Parameters `json:"JobParameters,omitempty" name:"JobParameters"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGCScheduleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGCScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegistryScanner struct {
	// DatabaseAutoUpdate

	DatabaseAutoUpdate *bool `json:"DatabaseAutoUpdate,omitempty" name:"DatabaseAutoUpdate"`
	// DatabaseUpdatedTime

	DatabaseUpdatedTime *string `json:"DatabaseUpdatedTime,omitempty" name:"DatabaseUpdatedTime"`
	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateRetentionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRetentionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRetentionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyImagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRetentionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 策略算法

		Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
		// 策略规则

		Rules []*Rule `json:"Rules,omitempty" name:"Rules"`
		// 策略触发条件

		Trigger *Trigger `json:"Trigger,omitempty" name:"Trigger"`
		// 策略关联资源

		Scope *Scope `json:"Scope,omitempty" name:"Scope"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRetentionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRetentionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Trigger struct {
	// 类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 设置

	Settings *Settings `json:"Settings,omitempty" name:"Settings"`
}

type BlockProjectRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 是否锁定

	Block *bool `json:"Block,omitempty" name:"Block"`
}

func (r *BlockProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BlockProjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GCHistory struct {
	// GcID

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 执行类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 回收大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeImageBuildingTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务

		ImageBuildingTask *ImageBuildingTask `json:"ImageBuildingTask,omitempty" name:"ImageBuildingTask"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageBuildingTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageBuildingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配的镜像总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回的镜像列表

		Images []*Image `json:"Images,omitempty" name:"Images"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Image struct {
	// 镜像的 ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 镜像的名字，注意它包含项目信息，如 library/golang

	Name *string `json:"Name,omitempty" name:"Name"`
	// 项目的 ID

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 该镜像包含多少 Artifact

	ArtifactCount *uint64 `json:"ArtifactCount,omitempty" name:"ArtifactCount"`
	// 该镜像被拉取了多少次

	PullCount *uint64 `json:"PullCount,omitempty" name:"PullCount"`
	// 该镜像的创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
	// 该镜像的修改时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeRegistryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共项目数

		PublicProjects *uint64 `json:"PublicProjects,omitempty" name:"PublicProjects"`
		// 私有项目数

		PrivateProjects *uint64 `json:"PrivateProjects,omitempty" name:"PrivateProjects"`
		// 公共镜像数

		PublicImages *uint64 `json:"PublicImages,omitempty" name:"PublicImages"`
		// 私有镜像数

		PrivateImages *uint64 `json:"PrivateImages,omitempty" name:"PrivateImages"`
		// 仓库的域名

		RegistryHost *string `json:"RegistryHost,omitempty" name:"RegistryHost"`
		// 仓库的扫描器

		Scanner *RegistryScanner `json:"Scanner,omitempty" name:"Scanner"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegistryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegistryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest

	// 项目的名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 镜像名称的模糊匹配

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 分页中指定第几页，从 1 开始，默认 1

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 分页中的页大小，默认 10

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageScanningStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像扫描的状态信息

		ScanOverview *ScanOverview `json:"ScanOverview,omitempty" name:"ScanOverview"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImageScanningStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImageScanningStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRetentionRequest struct {
	*tchttp.BaseRequest

	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 算法

	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
	// 规则

	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`
	// 相关资源信息

	Scope *Scope `json:"Scope,omitempty" name:"Scope"`
}

func (r *CreateRetentionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRetentionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGCScheduleRequest struct {
	*tchttp.BaseRequest

	// 仓库容量回收时间配置

	Schedule *Schedule `json:"Schedule,omitempty" name:"Schedule"`
	// 仓库容量回收参数配置

	Parameters *Parameters `json:"Parameters,omitempty" name:"Parameters"`
}

func (r *UpdateGCScheduleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateGCScheduleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGCScheduleRequest struct {
	*tchttp.BaseRequest

	// 仓库回收时间配置

	Schedule *Schedule `json:"Schedule,omitempty" name:"Schedule"`
	// 仓库回收参数

	Parameters *Parameters `json:"Parameters,omitempty" name:"Parameters"`
}

func (r *CreateGCScheduleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGCScheduleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageBuildingTaskRequest struct {
	*tchttp.BaseRequest

	// 构建任务

	ImageBuildingTask *ImageBuildingTask `json:"ImageBuildingTask,omitempty" name:"ImageBuildingTask"`
}

func (r *CreateImageBuildingTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageBuildingTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsRequest struct {
	*tchttp.BaseRequest

	// 项目名称的模糊匹配

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否只列举公共项目，true 只列举公共项目，false 只列举私有项目，如果要列举所有项目，请求里去掉该字段。

	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`
	// 分页，从第一页开始，默认为 1。

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 分页大小，默认为 10。

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 是否计算仓库下镜像下载总次数

	SumPullCount *bool `json:"SumPullCount,omitempty" name:"SumPullCount"`
	// 是否精确匹配

	IsExactMatch *bool `json:"IsExactMatch,omitempty" name:"IsExactMatch"`
}

func (r *DescribeProjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRetentionExecutionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// exec

		Execution *Executions `json:"Execution,omitempty" name:"Execution"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRetentionExecutionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRetentionExecutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 参数名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤值

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 操作符, 不传默认like, 可选equal/gte/lte

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type CleanImagesRequest struct {
	*tchttp.BaseRequest

	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 策略ID

	RetentionID *uint64 `json:"RetentionID,omitempty" name:"RetentionID"`
	// 是否仅计算不执行

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *CleanImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CleanImagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectMembersRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 项目成员名称

	EntityName *string `json:"EntityName,omitempty" name:"EntityName"`
}

func (r *ListProjectMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest

	// 项目的别名，支持中文等特殊字符，使用者可以只指定 Alias 而不指定 Name，如果没有指定 Name 会自动从 Alias 转换一个符合规范的 Name。

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 项目的名称，需要满足格式规范，如果不指定会从 Alias 中生成一个。Alias 和 Name 至少要指定一个。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 项目的说明。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 项目是否是公共项目，true 表示公共项目。

	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`
	// 是否开启自动漏洞扫描。

	AutoScan *bool `json:"AutoScan,omitempty" name:"AutoScan"`
	// 创建改项目的用户名称，和 SubAccountUin 对应

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 项目成员

	Members []*Member `json:"Members,omitempty" name:"Members"`
	// 项目配额

	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`
}

func (r *CreateProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageBuildingTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageBuildingTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageBuildingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartImageBuildingTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartImageBuildingTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartImageBuildingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAllProjectMembersRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 项目成员

	Members []*Member `json:"Members,omitempty" name:"Members"`
}

func (r *UpdateAllProjectMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAllProjectMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRetentionExecutionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// exec

		Executions []*Executions `json:"Executions,omitempty" name:"Executions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRetentionExecutionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRetentionExecutionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGCScheduleRequest struct {
	*tchttp.BaseRequest
}

func (r *GetGCScheduleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGCScheduleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest

	// 项目的 ID

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 项目的别名，设置该值以修改别名。

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 项目的描述，设置该值以修改描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 项目自动扫描

	AutoScan *bool `json:"AutoScan,omitempty" name:"AutoScan"`
	// 项目配额

	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`
}

func (r *ModifyProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Params struct {
	// 最近推送/拉取的K个

	LatestActiveK *uint64 `json:"LatestActiveK,omitempty" name:"LatestActiveK"`
	// 最近n天推/拉过的

	LastXDays *uint64 `json:"LastXDays,omitempty" name:"LastXDays"`
	// 最近n天拉过的

	NDaysSinceLastPull *uint64 `json:"NDaysSinceLastPull,omitempty" name:"NDaysSinceLastPull"`
	// 最近n天推过的

	NDaysSinceLastPush *uint64 `json:"NDaysSinceLastPush,omitempty" name:"NDaysSinceLastPush"`
	// 最近推拉取的N个

	LatestPulledN *uint64 `json:"LatestPulledN,omitempty" name:"LatestPulledN"`
	// 最近推送的K个

	LatestPushedK *uint64 `json:"LatestPushedK,omitempty" name:"LatestPushedK"`
}

type Vulnerability struct {
	// 镜像的摘要

	ArtifactDigest *string `json:"ArtifactDigest,omitempty" name:"ArtifactDigest"`
	// 漏洞的描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 漏洞的 ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 发现漏洞的包

	Package *string `json:"Package,omitempty" name:"Package"`
	// 发现漏洞的包的版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 漏洞修复的版本

	FixVersion *string `json:"FixVersion,omitempty" name:"FixVersion"`
	// 漏洞的严重程度

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 漏洞详情的链接地址

	Links []*string `json:"Links,omitempty" name:"Links"`
}

type ListProjectsRequest struct {
	*tchttp.BaseRequest

	// 项目名称的模糊匹配

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否只列举公共项目，true 只列举公共项目，false 只列举私有项目，如果要列举所有项目，请求里去掉该字段。

	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`
	// 分页，从第一页开始，默认为 1。

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 分页大小，默认为 10。

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListProjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageRequest struct {
	*tchttp.BaseRequest

	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
}

func (r *DeleteImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGCHistoryLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的日志

		Log *string `json:"Log,omitempty" name:"Log"`
		// 日志总量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGCHistoryLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGCHistoryLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Scanner struct {
	// 扫描器名称，如 Trivy

	Name *string `json:"Name,omitempty" name:"Name"`
	// 扫描器提供者

	Vendor *string `json:"Vendor,omitempty" name:"Vendor"`
	// 扫描器版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DeleteImageBuildingTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageBuildingTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageBuildingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配的项目总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回的项目列表

		Projects []*Project `json:"Projects,omitempty" name:"Projects"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Rule struct {
	// 是否启用

	Disabled *bool `json:"Disabled,omitempty" name:"Disabled"`
	// 动作

	Action *string `json:"Action,omitempty" name:"Action"`
	// 规则

	ScopeSelectors *ScopeSelectors `json:"ScopeSelectors,omitempty" name:"ScopeSelectors"`
	// tag规则

	TagSelectors []*TagSelectors `json:"TagSelectors,omitempty" name:"TagSelectors"`
	// *

	Template *string `json:"Template,omitempty" name:"Template"`
	// 参数

	Params *Params `json:"Params,omitempty" name:"Params"`
	// 优先级

	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}

type DescribeGCHistoryLogRequest struct {
	*tchttp.BaseRequest

	// GC ID

	GcID *uint64 `json:"GcID,omitempty" name:"GcID"`
	// 获取第几页

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页记录数

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeGCHistoryLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGCHistoryLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetGCStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetGCStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetGCStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRetentionRequest struct {
	*tchttp.BaseRequest

	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 策略ID

	RetentionID *uint64 `json:"RetentionID,omitempty" name:"RetentionID"`
}

func (r *DescribeRetentionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRetentionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageBuildingTaskRequest struct {
	*tchttp.BaseRequest

	// 镜像构建任务.

	ImageBuildingTask *ImageBuildingTask `json:"ImageBuildingTask,omitempty" name:"ImageBuildingTask"`
}

func (r *ModifyImageBuildingTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageBuildingTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanImageRequest struct {
	*tchttp.BaseRequest

	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像摘要名称

	Digest *string `json:"Digest,omitempty" name:"Digest"`
}

func (r *ScanImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchImagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 搜索镜像结果

		Images []*SearchImage `json:"Images,omitempty" name:"Images"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Executions struct {
	// execID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// retentionID

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 触发方式

	Trigger *string `json:"Trigger,omitempty" name:"Trigger"`
	// 是否仅计算大小不执行

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 清理个数

	Cleaned *uint64 `json:"Cleaned,omitempty" name:"Cleaned"`
	// 总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type TagSelectors struct {
	// 类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 描述

	Decoration *string `json:"Decoration,omitempty" name:"Decoration"`
	// 正则

	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`
	// 其它

	Extras *string `json:"Extras,omitempty" name:"Extras"`
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 项目的名称。

		Name *string `json:"Name,omitempty" name:"Name"`
		// 项目的 ID。

		ID *uint64 `json:"ID,omitempty" name:"ID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGCHistoriesRequest struct {
	*tchttp.BaseRequest

	// 分页，从第一页开始，默认为 1。

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 分页大小，默认为 10。

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeGCHistoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGCHistoriesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRetentionExecutionsRequest struct {
	*tchttp.BaseRequest

	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 策略ID

	RetentionID *uint64 `json:"RetentionID,omitempty" name:"RetentionID"`
	// 页码

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeRetentionExecutionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRetentionExecutionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGCStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *GetGCStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGCStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageBuildingTask struct {
	// 任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 输出的镜像仓库

	Repository *string `json:"Repository,omitempty" name:"Repository"`
	// 输出的镜像TAG

	Tag *string `json:"Tag,omitempty" name:"Tag"`
	// Dockerfile

	Dockerfile *string `json:"Dockerfile,omitempty" name:"Dockerfile"`
	// 启动次数

	Runtimes *int64 `json:"Runtimes,omitempty" name:"Runtimes"`
	// 当前状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 上次完成时间

	LastCompletionTimestamp *string `json:"LastCompletionTimestamp,omitempty" name:"LastCompletionTimestamp"`
	// 运行任务所在集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 运行任务所在Pod

	Pod *string `json:"Pod,omitempty" name:"Pod"`
	// 运行任务所在Container

	Container *string `json:"Container,omitempty" name:"Container"`
	// 运行任务所在名字空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 运行中的任务状态

	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
}

type Member struct {
	// 成员uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 成员名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeTagsRequest struct {
	*tchttp.BaseRequest

	// 项目的名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 镜像的名称，不包含项目信息

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 分页中指定第几页，从 1 开始，默认 1

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 分页中页的大小，默认 10

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配的项目总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回的项目列表

		Projects []*Project `json:"Projects,omitempty" name:"Projects"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProjectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGCHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// gc信息

		GCHistory *GCHistory `json:"GCHistory,omitempty" name:"GCHistory"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGCHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGCHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGCStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// GC功能是否开启

		IsGcOn *bool `json:"IsGcOn,omitempty" name:"IsGcOn"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGCStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGCStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageTag struct {
	// 镜像对应的摘要

	Digest *string `json:"Digest,omitempty" name:"Digest"`
	// 镜像的 Tag 名

	Tag *string `json:"Tag,omitempty" name:"Tag"`
	// 镜像大小，单位为字节

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 镜像的架构，如 amd64

	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
	// 镜像是否为可变的，不可变的 tag 不允许删除或覆盖

	Immutable *bool `json:"Immutable,omitempty" name:"Immutable"`
	// 该 Tag 推送的时间

	PushTime *string `json:"PushTime,omitempty" name:"PushTime"`
	// 操作系统

	OS *string `json:"OS,omitempty" name:"OS"`
	// 漏洞扫描状态

	ScanOverview *ScanOverview `json:"ScanOverview,omitempty" name:"ScanOverview"`
	// 拉取时间

	PullTime *string `json:"PullTime,omitempty" name:"PullTime"`
}

type ScopeSelectorsRepository struct {
	// 类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 解释

	Decoration *string `json:"Decoration,omitempty" name:"Decoration"`
	// 正则表达式

	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`
}

type DeleteTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAccountRequest struct {
	*tchttp.BaseRequest

	// 用户名

	User *string `json:"User,omitempty" name:"User"`
}

func (r *GetAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest

	// 项目 ID

	ID *uint64 `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanOverview struct {
	// 扫描完成百分比

	CompletePercent *uint64 `json:"CompletePercent,omitempty" name:"CompletePercent"`
	// 扫描时间，单位秒

	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
	// 漏洞扫描的状态，如 Success

	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`
	// 镜像的漏洞严重程度

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 漏洞扫描开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 漏洞扫描结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type ModifyProjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAllProjectMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAllProjectMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAllProjectMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGCHistoriesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// gc总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回的gc历史列表

		GCHistory []*GCHistory `json:"GCHistory,omitempty" name:"GCHistory"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGCHistoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGCHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchImage struct {
	// 镜像总量

	ArtifactCount *int64 `json:"ArtifactCount,omitempty" name:"ArtifactCount"`
	// 项目ID

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 是否公共项目

	ProjectPublic *bool `json:"ProjectPublic,omitempty" name:"ProjectPublic"`
	// 镜像拉取次数

	PullCount *int64 `json:"PullCount,omitempty" name:"PullCount"`
	// 镜像名称

	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`
}

type GetRetentionExecutionRequest struct {
	*tchttp.BaseRequest

	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 策略ID

	RetentionID *uint64 `json:"RetentionID,omitempty" name:"RetentionID"`
	// 执行记录ID

	ExecutionID *uint64 `json:"ExecutionID,omitempty" name:"ExecutionID"`
}

func (r *GetRetentionExecutionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRetentionExecutionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScanImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScanImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
