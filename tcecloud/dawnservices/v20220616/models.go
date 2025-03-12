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

package v20220616

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type PlanApplicationInstance struct {
	// 应用分支ID

	ApplicationBranchID *string `json:"ApplicationBranchID,omitempty" name:"ApplicationBranchID"`
	// 应用分支名称

	ApplicationBranchName *string `json:"ApplicationBranchName,omitempty" name:"ApplicationBranchName"`
	// 应用英文代号

	ApplicationCode *string `json:"ApplicationCode,omitempty" name:"ApplicationCode"`
	// 应用ID

	ApplicationID *string `json:"ApplicationID,omitempty" name:"ApplicationID"`
	// 应用实例ID

	ApplicationInstanceID *string `json:"ApplicationInstanceID,omitempty" name:"ApplicationInstanceID"`
	// 应用实例名称

	ApplicationInstanceName *string `json:"ApplicationInstanceName,omitempty" name:"ApplicationInstanceName"`
	// 应用实例UUID

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
	// 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// 应用配置,规划包中的values.yaml格式

	Config *string `json:"Config,omitempty" name:"Config"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 删除时间

	DeletedAt *int64 `json:"DeletedAt,omitempty" name:"DeletedAt"`
	// ( primary key ) 应用自增ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 应用部署位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 关联的规划ID

	PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
	// 产品实例ID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// 子系统实例UUID

	SubsystemInstanceUUID *string `json:"SubsystemInstanceUUID,omitempty" name:"SubsystemInstanceUUID"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type PlanSite struct {
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 删除时间

	DeletedAt *int64 `json:"DeletedAt,omitempty" name:"DeletedAt"`
	// 规划UUID

	PlanUUID *string `json:"PlanUUID,omitempty" name:"PlanUUID"`
	// 局点中生效的规划版本

	PlanVersion *string `json:"PlanVersion,omitempty" name:"PlanVersion"`
	// ( primary key ) 局点的UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type DeleteSolutionTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSolutionTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSolutionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCompareChartFileListRequest struct {
	*tchttp.BaseRequest

	// 应用制品uuid

	ApplicationArtifactUUID *string `json:"ApplicationArtifactUUID,omitempty" name:"ApplicationArtifactUUID"`
	// 需要对比的chart包uuid列表

	ChartUUIDList []*string `json:"ChartUUIDList,omitempty" name:"ChartUUIDList"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *GetCompareChartFileListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCompareChartFileListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSiteApplicationsRequest struct {
	*tchttp.BaseRequest

	// 局点uuid

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// 筛选信息

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
}

func (r *ListSiteApplicationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSiteApplicationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProgressDetail struct {
	// 子任务Id

	SubTaskId *int64 `json:"SubTaskId,omitempty" name:"SubTaskId"`
	// 物料类型

	Model *string `json:"Model,omitempty" name:"Model"`
	// 上传状态

	ModelStatus *string `json:"ModelStatus,omitempty" name:"ModelStatus"`
	// 提示信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品版本UUID

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	// 产品版本

	ProductVersion *string `json:"ProductVersion,omitempty" name:"ProductVersion"`
	// 应用数

	TotalAppCount *int64 `json:"TotalAppCount,omitempty" name:"TotalAppCount"`
	// 状态详情

	ProductDetail *ProductDetail `json:"ProductDetail,omitempty" name:"ProductDetail"`
}

type DescribeSheetTemplateRequest struct {
	*tchttp.BaseRequest

	// SheetID 运维单ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// Usage 模板用途

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// FilterApp 过滤模板

	FilterApp *bool `json:"FilterApp,omitempty" name:"FilterApp"`
}

func (r *DescribeSheetTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSheetTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLatestMaterialPathResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 物料目录

		MaterialPath []*string `json:"MaterialPath,omitempty" name:"MaterialPath"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLatestMaterialPathResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLatestMaterialPathResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSiteOperationSheetRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
}

func (r *DescribeSiteOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSiteOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLatestPackageUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Url package下载Url

		Url *string `json:"Url,omitempty" name:"Url"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLatestPackageUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLatestPackageUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCommonOperationSheetTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCommonOperationSheetTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCommonOperationSheetTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrchestratingItemsSheet struct {
	// 子变更单ID

	ChildSheetID *string `json:"ChildSheetID,omitempty" name:"ChildSheetID"`
	// 子变更单类别，变更单或产品变更单

	ChildSheetKind *string `json:"ChildSheetKind,omitempty" name:"ChildSheetKind"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// ( primary key ) 自增主键

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 产品实例运维单UUID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type ImportDeploymentPhaseApplicationArtifactResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		PackageManagerURL *string `json:"PackageManagerURL,omitempty" name:"PackageManagerURL"`
		//

		Version *string `json:"Version,omitempty" name:"Version"`
		//

		Values *string `json:"Values,omitempty" name:"Values"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportDeploymentPhaseApplicationArtifactResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportDeploymentPhaseApplicationArtifactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubsystemDetail struct {
	// 子系统 ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 子系统 UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 子系统 Tag

	Tag *string `json:"Tag,omitempty" name:"Tag"`
	// 子系统名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DeployEnvResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// EnvID 环境 ID

		EnvID *int64 `json:"EnvID,omitempty" name:"EnvID"`
		// EnvUUID 环境 UUID

		EnvUUID *string `json:"EnvUUID,omitempty" name:"EnvUUID"`
		// SheetID 运维单 UUID, SheetID + EnvUUID 可用于拼接产品市场部署页面链接

		SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
		// JobID 任务 UUID

		JobID *string `json:"JobID,omitempty" name:"JobID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeployEnvResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeployEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationSheetRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
}

func (r *DeleteOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartApplicationDeployResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SheetID 运维单 ID

		SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
		// JobID 部署任务 ID

		JobID *string `json:"JobID,omitempty" name:"JobID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartApplicationDeployResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartApplicationDeployResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProductInstanceOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProductInstanceOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProductInstanceOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProductPlanRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// CloudID 云UUID

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// ProductInstanceUUID 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// Config 修改后的产品配置信息

	Config *string `json:"Config,omitempty" name:"Config"`
	// Modifier 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
}

func (r *ModifyProductPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProductPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSheetTemplateRequest struct {
	*tchttp.BaseRequest

	// SheetID 运维单ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// TemplateUsage 模板用途

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// Template 模板内容

	Template *string `json:"Template,omitempty" name:"Template"`
	// Parameters 模板参数

	Parameters *string `json:"Parameters,omitempty" name:"Parameters"`
}

func (r *UpdateSheetTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSheetTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationArtifactChartFileRequest struct {
	*tchttp.BaseRequest

	// 应用制品UUID

	ApplicationArtifactUUID *string `json:"ApplicationArtifactUUID,omitempty" name:"ApplicationArtifactUUID"`
	// chart包UUID

	ChartUUID *string `json:"ChartUUID,omitempty" name:"ChartUUID"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *GetApplicationArtifactChartFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationArtifactChartFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KVPair struct {
	//

	Key *string `json:"Key,omitempty" name:"Key"`
	//

	Value *string `json:"Value,omitempty" name:"Value"`
	//

	Options *Options `json:"Options,omitempty" name:"Options"`
}

type SiteProduct struct {
	// 产品UUID

	ProductUUID *string `json:"ProductUUID,omitempty" name:"ProductUUID"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品下应用信息

	Applications []*SiteProductApplication `json:"Applications,omitempty" name:"Applications"`
}

type AppRuntimeInfo struct {
	// AppName 应用名称

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// Phase 应用状态

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// Ready 组件就绪数量

	Ready *int64 `json:"Ready,omitempty" name:"Ready"`
	// Modulues 组件总数量

	Modulues *int64 `json:"Modulues,omitempty" name:"Modulues"`
	// Health 应用健康状态

	Health *bool `json:"Health,omitempty" name:"Health"`
	// Age 应用运行时长

	Age *string `json:"Age,omitempty" name:"Age"`
	// Version 版本信息

	Version *string `json:"Version,omitempty" name:"Version"`
}

type GetRunningApplicationNameRequest struct {
	*tchttp.BaseRequest

	//

	Applications []*ApplicationUUIDAndName `json:"Applications,omitempty" name:"Applications"`
}

func (r *GetRunningApplicationNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRunningApplicationNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishCommonOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishCommonOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishCommonOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployPlanProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SheetID 运维单 UUID

		SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
		// JobID 任务 UUID

		JobID *string `json:"JobID,omitempty" name:"JobID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeployPlanProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeployPlanProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductLocation struct {
	//

	RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
	//

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	//

	RegionNameCn *string `json:"RegionNameCn,omitempty" name:"RegionNameCn"`
}

type DescribeProductInstanceInfoRequest struct {
	*tchttp.BaseRequest

	// CloudID 云ID

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// RegionID 区域ID

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
}

func (r *DescribeProductInstanceInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductInstanceInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartProductInstanceOperationSheetRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	//

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	//

	PipelineID *string `json:"PipelineID,omitempty" name:"PipelineID"`
	//

	Pipeline *string `json:"Pipeline,omitempty" name:"Pipeline"`
	//

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *StartProductInstanceOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartProductInstanceOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationOverview struct {
	// Total 总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// Healthy 健康数量

	Healthy *int64 `json:"Healthy,omitempty" name:"Healthy"`
	// Unhealthy 异常数量

	Unhealthy *int64 `json:"Unhealthy,omitempty" name:"Unhealthy"`
}

type DescribeApplicationYamlRequest struct {
	*tchttp.BaseRequest

	// 应用包UUID

	ApplicationPackageID *string `json:"ApplicationPackageID,omitempty" name:"ApplicationPackageID"`
}

func (r *DescribeApplicationYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApplicationYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationTemplateIDRequest struct {
	*tchttp.BaseRequest

	// ApplicationStandard 应用标准

	ApplicationStandard *string `json:"ApplicationStandard,omitempty" name:"ApplicationStandard"`
	// ApplicationName 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// ApplicationPackageUUID 应用包名称

	ApplicationPackageUUID *string `json:"ApplicationPackageUUID,omitempty" name:"ApplicationPackageUUID"`
	// Type 描述模版类别,取值为product application等

	Type *string `json:"Type,omitempty" name:"Type"`
	// Kind 模板类型 取值为 deploy update upgrade

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// Labels 标签

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
}

func (r *GetApplicationTemplateIDRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationTemplateIDRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishProductInstanceOperationSheetRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
}

func (r *PublishProductInstanceOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishProductInstanceOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialProductDetailsRequest struct {
	*tchttp.BaseRequest

	// MaterialUUID 物料 UUID

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
}

func (r *ListMaterialProductDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialProductDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductCommonEnumRequest struct {
	*tchttp.BaseRequest
}

func (r *ListProductCommonEnumRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductCommonEnumRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductInstanceOperationSheetRequest struct {
	*tchttp.BaseRequest

	// SheetName 运维单名称

	SheetName *string `json:"SheetName,omitempty" name:"SheetName"`
	// SheetType 运维单类型，Deploy 部署， Update 变更， Upgrade 升级

	SheetType *string `json:"SheetType,omitempty" name:"SheetType"`
	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// RegionUUID 区域UUID

	RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
	//

	RegionNameCN *string `json:"RegionNameCN,omitempty" name:"RegionNameCN"`
	//

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 变更影响

	Influence *string `json:"Influence,omitempty" name:"Influence"`
	// 关联产品和应用实例id

	ProductInstanceAppUUIDList []*ProductInstanceAppUUID `json:"ProductInstanceAppUUIDList,omitempty" name:"ProductInstanceAppUUIDList"`
	// Description 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// Creator 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// ParentSheetID 父变更单ID , 用于给父单创建变更项的时候会传，如果不传则会自动创建父变更单

	ParentSheetID *string `json:"ParentSheetID,omitempty" name:"ParentSheetID"`
	// StageName 分组名称

	StageName *string `json:"StageName,omitempty" name:"StageName"`
	// Risk 风险项

	Risk *string `json:"Risk,omitempty" name:"Risk"`
	// IsInternal 是否内置运维单

	IsInternal *bool `json:"IsInternal,omitempty" name:"IsInternal"`
}

func (r *CreateProductInstanceOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductInstanceOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductVersionBrief struct {
	// 产品版本 ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 产品版本 UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 产品版本数据 Tag

	Tag *string `json:"Tag,omitempty" name:"Tag"`
	// 产品版本名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品版本发布状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 产品版本描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 应用全局配置

	Config *string `json:"Config,omitempty" name:"Config"`
}

type DescribeAppRollbackConfigRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	//

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
}

func (r *DescribeAppRollbackConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppRollbackConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportSolutionTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		SolutionTemplateList []*OperationTemplateKey `json:"SolutionTemplateList,omitempty" name:"SolutionTemplateList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportSolutionTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportSolutionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProjectSiteInfo struct {
	// 认证信息

	Authorization *string `json:"Authorization,omitempty" name:"Authorization"`
	// 云拓扑的ID

	CloudID *int64 `json:"CloudID,omitempty" name:"CloudID"`
	// 连接状态

	ConnectStatus *string `json:"ConnectStatus,omitempty" name:"ConnectStatus"`
	// 连接的URL

	ConnectURL *string `json:"ConnectURL,omitempty" name:"ConnectURL"`
	//

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 删除时间

	DeletedAt *int64 `json:"DeletedAt,omitempty" name:"DeletedAt"`
	// 底座部署时使用的运维单ID

	DeploySheetID *string `json:"DeploySheetID,omitempty" name:"DeploySheetID"`
	// 部署状态

	DeployStatus *string `json:"DeployStatus,omitempty" name:"DeployStatus"`
	// 局点描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 局点关联环境UUID

	EnvID *string `json:"EnvID,omitempty" name:"EnvID"`
	// ( primary key )

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 是否生效的局点

	IsActive *bool `json:"IsActive,omitempty" name:"IsActive"`
	// 局点元数据

	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`
	// 局点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 局点的网络类型

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	//

	PlanSite *PlanSite `json:"PlanSite,omitempty" name:"PlanSite"`
	//

	ProjectTopology *ProjectTopology `json:"ProjectTopology,omitempty" name:"ProjectTopology"`
	// 服务器的通用访问信息

	ServerLoginInfo *string `json:"ServerLoginInfo,omitempty" name:"ServerLoginInfo"`
	// 局点来源

	Source *int64 `json:"Source,omitempty" name:"Source"`
	// 局点的UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	//

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type QueryUploadMaterialTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskID *int64 `json:"TaskID,omitempty" name:"TaskID"`
}

func (r *QueryUploadMaterialTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryUploadMaterialTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImportMaterialTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		TaskID *int64 `json:"TaskID,omitempty" name:"TaskID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImportMaterialTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImportMaterialTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanMiddlewareInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// MiddlewareSet 中间件列表

		MiddlewareSet []*MiddlewareInfo `json:"MiddlewareSet,omitempty" name:"MiddlewareSet"`
		// TotalCount 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPlanMiddlewareInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanMiddlewareInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCompareChartFileListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Files []*CompareChartFile `json:"Files,omitempty" name:"Files"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCompareChartFileListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCompareChartFileListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationDataSchemaRequest struct {
	*tchttp.BaseRequest
}

func (r *GetOperationDataSchemaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationDataSchemaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProjectTopology struct {
	//

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 局点描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// ( primary key )

	ID *int64 `json:"ID,omitempty" name:"ID"`
	//

	ProjectRegion []*ProjectRegion `json:"ProjectRegion,omitempty" name:"ProjectRegion"`
	// 云拓扑的UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	//

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type PauseNodeRequest struct {
	*tchttp.BaseRequest

	//

	DagJobUUID *string `json:"DagJobUUID,omitempty" name:"DagJobUUID"`
	//

	DagNodeUUID *string `json:"DagNodeUUID,omitempty" name:"DagNodeUUID"`
}

func (r *PauseNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PauseNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActiveSiteInfo struct {
	//

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	//

	SiteName *string `json:"SiteName,omitempty" name:"SiteName"`
}

type CreateSolutionTemplateRequest struct {
	*tchttp.BaseRequest

	//

	TemplateInfo *SolutionTemplate `json:"TemplateInfo,omitempty" name:"TemplateInfo"`
}

func (r *CreateSolutionTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSolutionTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudItem struct {
	//

	CloudUUID *string `json:"CloudUUID,omitempty" name:"CloudUUID"`
	//

	CloudName *string `json:"CloudName,omitempty" name:"CloudName"`
	//

	Regions []*RegionItem `json:"Regions,omitempty" name:"Regions"`
}

type GeneratePipelineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		PipelineID *string `json:"PipelineID,omitempty" name:"PipelineID"`
		//

		Pipeline *string `json:"Pipeline,omitempty" name:"Pipeline"`
		//

		DAG *DAG `json:"DAG,omitempty" name:"DAG"`
		//

		WarnMessage *string `json:"WarnMessage,omitempty" name:"WarnMessage"`
		//

		MissApplicationList []*string `json:"MissApplicationList,omitempty" name:"MissApplicationList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneratePipelineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GeneratePipelineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationInstancesByApplicationNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ApplicationInstances  应用实例列表

		ApplicationInstances []*PlanApplicationInstance `json:"ApplicationInstances,omitempty" name:"ApplicationInstances"`
		// SiteUUID 局点 UUID

		SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
		// PlanID 规划 ID

		PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
		// CloudUUID 云 UUID

		CloudUUID *string `json:"CloudUUID,omitempty" name:"CloudUUID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationInstancesByApplicationNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationInstancesByApplicationNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCommonOperationSheetAppInstancesRequest struct {
	*tchttp.BaseRequest

	// SheetID 运维单ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// Usage 模板用途

	Usage *string `json:"Usage,omitempty" name:"Usage"`
}

func (r *ListCommonOperationSheetAppInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCommonOperationSheetAppInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrchestratingItemsVar struct {
	// SheetKind为common时，可以编排此项

	Sheets []*OrchestratingItemsSheet `json:"Sheets,omitempty" name:"Sheets"`
	// SheetKind为product时，此项只包含一个产品

	Products []*OrchestratingItemsProduct `json:"Products,omitempty" name:"Products"`
}

type ListCloudClustersRequest struct {
	*tchttp.BaseRequest
}

func (r *ListCloudClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CompareChartFile struct {
	// 文件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 文件路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 文件列表

	Contents []*ChartContent `json:"Contents,omitempty" name:"Contents"`
}

type ListProjectSiteInfoDetailsRequest struct {
	*tchttp.BaseRequest

	// Filter 筛选信息

	_Filter []*FilterType `json:"_Filter,omitempty" name:"_Filter"`
	// Offset 偏移值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit 返回数据量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListProjectSiteInfoDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectSiteInfoDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplySiteMaterialRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点 UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// MaterialUUID 使用的物料 UUID

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
	// CreateSite 是否创建局点

	CreateSite *int64 `json:"CreateSite,omitempty" name:"CreateSite"`
	// SiteName 局点名称

	SiteName *string `json:"SiteName,omitempty" name:"SiteName"`
	// Description 云描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// ConnectUrl 云连接信息

	ConnectUrl *string `json:"ConnectUrl,omitempty" name:"ConnectUrl"`
	// EnvID 环境ID

	EnvID *string `json:"EnvID,omitempty" name:"EnvID"`
	// Metadata 局点元数据

	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`
}

func (r *ApplySiteMaterialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplySiteMaterialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportedProductInstance struct {
	// ProductUUID 产品UUID

	ProductUUID *string `json:"ProductUUID,omitempty" name:"ProductUUID"`
	// ProductVersion 产品版本UUID

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	// ProductInstanceIDSet 产品实例对应的ID集合

	ProductInstanceIDSet []*int64 `json:"ProductInstanceIDSet,omitempty" name:"ProductInstanceIDSet"`
}

type ListOperationSheetInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Total *int64 `json:"Total,omitempty" name:"Total"`
		//

		Sheets []*CommonOperationSheet `json:"Sheets,omitempty" name:"Sheets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOperationSheetInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationSheetInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJobOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJobOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJobOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductInstanceAppsRequest struct {
	*tchttp.BaseRequest

	//

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	//

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	//

	ApplicationUUIDList []*string `json:"ApplicationUUIDList,omitempty" name:"ApplicationUUIDList"`
}

func (r *ListProductInstanceAppsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductInstanceAppsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationSheetInstancesRequest struct {
	*tchttp.BaseRequest

	//

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	//

	SheetId *string `json:"SheetId,omitempty" name:"SheetId"`
	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	//

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListOperationSheetInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationSheetInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationInfoWithProduct struct {
	// 产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// ProductInstanceName 产品实例名称

	ProductInstanceName *string `json:"ProductInstanceName,omitempty" name:"ProductInstanceName"`
	// SubsystemCode 子系统Code

	SubsystemCode *string `json:"SubsystemCode,omitempty" name:"SubsystemCode"`
	// SubsystemName 子系统名字

	SubsystemName *string `json:"SubsystemName,omitempty" name:"SubsystemName"`
	// ApplicationName 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// ApplicationID 应用ID

	ApplicationID *string `json:"ApplicationID,omitempty" name:"ApplicationID"`
	// ApplicationCode 应用Code

	ApplicationCode *string `json:"ApplicationCode,omitempty" name:"ApplicationCode"`
	//

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	//

	ApplicationVersion *string `json:"ApplicationVersion,omitempty" name:"ApplicationVersion"`
}

type DiffApplicationInstanceConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		ApplicationInstances []*AppDiff `json:"ApplicationInstances,omitempty" name:"ApplicationInstances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DiffApplicationInstanceConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DiffApplicationInstanceConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationPackageDetailRequest struct {
	*tchttp.BaseRequest

	// ApplicationPackageID 应用包ID

	ApplicationPackageID *int64 `json:"ApplicationPackageID,omitempty" name:"ApplicationPackageID"`
}

func (r *GetApplicationPackageDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationPackageDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectSiteInfoDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		ProjectSiteInfo *ProjectSiteInfo `json:"ProjectSiteInfo,omitempty" name:"ProjectSiteInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProjectSiteInfoDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectSiteInfoDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAppRuntimeNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SuccessSet 修改成功的应用列表

		SuccessSet []*AppRuntimeNameInfo `json:"SuccessSet,omitempty" name:"SuccessSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAppRuntimeNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAppRuntimeNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAtomicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Atomic 原子操作Json结构

		Atomic *JsonAtomic `json:"Atomic,omitempty" name:"Atomic"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAtomicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAtomicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Application 产品下应用概览信息

		Application *ApplicationOverview `json:"Application,omitempty" name:"Application"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationMetadata struct {
	// Standard 应用的类型，用于应用的模版匹配

	Standard *string `json:"Standard,omitempty" name:"Standard"`
	// Type 应用的类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type OperationSheetRelation struct {
	// 子变更单ID

	ChildSheetID *string `json:"ChildSheetID,omitempty" name:"ChildSheetID"`
	// 子变更单类别，变更单或产品变更单

	ChildSheetKind *string `json:"ChildSheetKind,omitempty" name:"ChildSheetKind"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// ( primary key ) 自增主键

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 产品实例运维单UUID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type SubmitOperationSheetApprovalRequest struct {
	*tchttp.BaseRequest

	// SheetID 运维单 ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// Approver 审批人

	Approver *string `json:"Approver,omitempty" name:"Approver"`
	// ApprovalComment 审批意见

	ApprovalComment *string `json:"ApprovalComment,omitempty" name:"ApprovalComment"`
	// ApprovalStatus 审批状态

	ApprovalStatus *string `json:"ApprovalStatus,omitempty" name:"ApprovalStatus"`
}

func (r *SubmitOperationSheetApprovalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubmitOperationSheetApprovalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCommonOperationSheetRequest struct {
	*tchttp.BaseRequest

	//

	ParentSheetID *string `json:"ParentSheetID,omitempty" name:"ParentSheetID"`
	// 额外信息

	Annotations *string `json:"Annotations,omitempty" name:"Annotations"`
	// 云ID, 产品部署运维单依赖该字段查询

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// ( primary key ) 自增主键

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 变更影响

	Influence *string `json:"Influence,omitempty" name:"Influence"`
	// 是否进行展示

	IsDisplay *int64 `json:"IsDisplay,omitempty" name:"IsDisplay"`
	// 是否是预置变更单

	IsInternal *int64 `json:"IsInternal,omitempty" name:"IsInternal"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 附件信息

	OperationSheetAttachment []*OperationSheetAttachment `json:"OperationSheetAttachment,omitempty" name:"OperationSheetAttachment"`
	// 父子单关联关系，仅当此单为父单时有效

	OperationSheetRelation []*OperationSheetRelation `json:"OperationSheetRelation,omitempty" name:"OperationSheetRelation"`
	// 变更单具体的编排信息, 一个变更单默认只有一个编排模板，预留model_array用于扩展

	OperationSheetTemplate []*OperationSheetTemplate `json:"OperationSheetTemplate,omitempty" name:"OperationSheetTemplate"`
	// 待编排项

	OrchestratingItems *string `json:"OrchestratingItems,omitempty" name:"OrchestratingItems"`
	// 负责人

	Owners *string `json:"Owners,omitempty" name:"Owners"`
	// 产品实例ID, 产品部署运维单依赖该字段查询

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// 区域ID, 产品部署运维单依赖该字段查询

	RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
	// 风险项

	Risk *string `json:"Risk,omitempty" name:"Risk"`
	// 场景

	Scene *string `json:"Scene,omitempty" name:"Scene"`
	// 运维单ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// 运维单类型（common,product,site,productInstance）

	SheetKind *string `json:"SheetKind,omitempty" name:"SheetKind"`
	// 运维单名称

	SheetName *string `json:"SheetName,omitempty" name:"SheetName"`
	// 运维模板类型(deploy,update,upgrade)

	SheetType *string `json:"SheetType,omitempty" name:"SheetType"`
	// 运维单来源:bug单, 交付单, 导入，人工创建

	Source *string `json:"Source,omitempty" name:"Source"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// AnnotationsVar是Annotations的结构体化展示

	AnnotationsVar *AnnotationsVar `json:"AnnotationsVar,omitempty" name:"AnnotationsVar"`
	// OrchestratingItemsVar是OrchestratingItems的结构体化展示

	OrchestratingItemsVar *OrchestratingItemsVar `json:"OrchestratingItemsVar,omitempty" name:"OrchestratingItemsVar"`
	// Children 是子变更单基本信息

	Children []*CommonOperationSheet `json:"Children,omitempty" name:"Children"`
	// Templates 是变更单编排相关的信息

	Templates []*OperationSheetTemplate `json:"Templates,omitempty" name:"Templates"`
}

func (r *UpdateCommonOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCommonOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportAtomicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// AtomicList 导入的到的原子操作查询key列表

		AtomicList []*AtomicKey `json:"AtomicList,omitempty" name:"AtomicList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportAtomicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportAtomicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationInstancesByApplicationNameRequest struct {
	*tchttp.BaseRequest

	// ApplicationName 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// SiteUUID is site uuid

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
}

func (r *ListApplicationInstancesByApplicationNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationInstancesByApplicationNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectSiteInfoDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ProjectSiteInfos 云列表

		ProjectSiteInfos []*ProjectSiteInfo `json:"ProjectSiteInfos,omitempty" name:"ProjectSiteInfos"`
		// TotalCount 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProjectSiteInfoDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectSiteInfoDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyChartFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// chart包生成uuid

		ChartUUID *string `json:"ChartUUID,omitempty" name:"ChartUUID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyChartFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyChartFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMachineSSHConfigRequest struct {
	*tchttp.BaseRequest

	// 机器ip

	IP *string `json:"IP,omitempty" name:"IP"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *GetMachineSSHConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMachineSSHConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParseMaterialRequest struct {
	*tchttp.BaseRequest

	// 物料所在服务器

	MaterialServer *string `json:"MaterialServer,omitempty" name:"MaterialServer"`
	// 物料父目录

	MaterialPath *string `json:"MaterialPath,omitempty" name:"MaterialPath"`
	// 机器用户

	MaterialSSHUser *string `json:"MaterialSSHUser,omitempty" name:"MaterialSSHUser"`
	// 机器密码

	MaterialSSHPassword *string `json:"MaterialSSHPassword,omitempty" name:"MaterialSSHPassword"`
	// 机器端口

	MaterialSSHPort *int64 `json:"MaterialSSHPort,omitempty" name:"MaterialSSHPort"`
}

func (r *ParseMaterialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ParseMaterialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSheetAppInstancesRequest struct {
	*tchttp.BaseRequest

	// SheetID 运维单ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// Usage 模板用途

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// ApplicationInstanceList 产品实例列表

	ApplicationInstanceList []*SheetApplicationInstance `json:"ApplicationInstanceList,omitempty" name:"ApplicationInstanceList"`
}

func (r *UpdateSheetAppInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSheetAppInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductInstanceOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SheetID 变更单ID

		SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductInstanceOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductInstanceOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPublicMaterialSolutionVersionsRequest struct {
	*tchttp.BaseRequest
}

func (r *ListPublicMaterialSolutionVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPublicMaterialSolutionVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOtaCheckConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOtaCheckConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOtaCheckConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckSiteOperationSheetMaterialRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
}

func (r *CheckSiteOperationSheetMaterialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSiteOperationSheetMaterialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportApplicationPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		ApplicationInstanceList []*ApplicationInstanceForImportApplicationPlan `json:"ApplicationInstanceList,omitempty" name:"ApplicationInstanceList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportApplicationPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportApplicationPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialOperationTemplateRequest struct {
	*tchttp.BaseRequest

	// MaterialUUID 物料UUID

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
	// ProductVersionUUID 产品版本uuid

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	// Kind 运维模板类型  deploy-部署模板，update-变更模板，expand-扩容模板

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	//

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListMaterialOperationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialOperationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialSyncTaskRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// MaterialList 物料信息

	MaterialList []*MaterialSync `json:"MaterialList,omitempty" name:"MaterialList"`
}

func (r *CreateMaterialSyncTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialSyncTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportSolutionTemplateRequest struct {
	*tchttp.BaseRequest

	//

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
	//

	LocalFilePath *string `json:"LocalFilePath,omitempty" name:"LocalFilePath"`
	//

	Package *string `json:"Package,omitempty" name:"Package"`
}

func (r *ImportSolutionTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportSolutionTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudTopologyRequest struct {
	*tchttp.BaseRequest

	// Clouds 局点/云 ID列表

	Clouds []*string `json:"Clouds,omitempty" name:"Clouds"`
}

func (r *DescribeCloudTopologyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudTopologyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyProductInstanceOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SheetID 复制后的变更单ID

		SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyProductInstanceOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyProductInstanceOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCommonOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCommonOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCommonOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChartHistory struct {
	// chart包uuid

	ChartUUID *string `json:"ChartUUID,omitempty" name:"ChartUUID"`
	// 制品类型： chart

	Type *string `json:"Type,omitempty" name:"Type"`
	// 架构

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	// 表示导入还是修改  Import:0 or Modify：1

	Source *int64 `json:"Source,omitempty" name:"Source"`
	// 表示是否生效  未生效:0 or 生效：1

	Effect *int64 `json:"Effect,omitempty" name:"Effect"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 修改时间

	CreateAt *string `json:"CreateAt,omitempty" name:"CreateAt"`
	// 下载链接

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
}

type StartSiteOperationSheetRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
}

func (r *StartSiteOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartSiteOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Values *string `json:"Values,omitempty" name:"Values"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckSheetTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CheckOK 检查是否通过

		CheckOK *bool `json:"CheckOK,omitempty" name:"CheckOK"`
		// Template 模板

		Template *string `json:"Template,omitempty" name:"Template"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckSheetTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSheetTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetActivePlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规划协议版本

		APIVersion *string `json:"APIVersion,omitempty" name:"APIVersion"`
		// 云拓扑的ID

		CloudID *int64 `json:"CloudID,omitempty" name:"CloudID"`
		// 规划commit的UUID

		CommitUUID *string `json:"CommitUUID,omitempty" name:"CommitUUID"`
		// 创建时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 创建人

		Creator *string `json:"Creator,omitempty" name:"Creator"`
		// 删除时间

		DeletedAt *int64 `json:"DeletedAt,omitempty" name:"DeletedAt"`
		// ( primary key ) 规划ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// 修改人

		Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
		// 规划名称

		Name *string `json:"Name,omitempty" name:"Name"`
		//

		PlanProductInstance []*PlanProductInstance `json:"PlanProductInstance,omitempty" name:"PlanProductInstance"`
		//

		PlanServerInfo []*PlanServerInfo `json:"PlanServerInfo,omitempty" name:"PlanServerInfo"`
		// 做规划的时间

		PlanTime *string `json:"PlanTime,omitempty" name:"PlanTime"`
		// 规划包的类型:new/expand/upgrade

		PlanType *string `json:"PlanType,omitempty" name:"PlanType"`
		// 局点名称

		SiteName *string `json:"SiteName,omitempty" name:"SiteName"`
		// 局点UUID

		SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
		// 规划使用的解决方案Code

		SolutionCode *string `json:"SolutionCode,omitempty" name:"SolutionCode"`
		// 规划使用的解决方案名称

		SolutionName *string `json:"SolutionName,omitempty" name:"SolutionName"`
		// 规划时使用的解决方案版本数据Tag

		SolutionVersionDataTag *string `json:"SolutionVersionDataTag,omitempty" name:"SolutionVersionDataTag"`
		// 规划使用的解决方案版本ID

		SolutionVersionID *string `json:"SolutionVersionID,omitempty" name:"SolutionVersionID"`
		// 规划使用的解决方案版本名称

		SolutionVersionName *string `json:"SolutionVersionName,omitempty" name:"SolutionVersionName"`
		// 规划包中的云拓扑信息

		Topology *string `json:"Topology,omitempty" name:"Topology"`
		// 规划的类型:standard/site

		Type *string `json:"Type,omitempty" name:"Type"`
		// 规划UUID

		UUID *string `json:"UUID,omitempty" name:"UUID"`
		// 更新时间

		UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
		// 规划版本

		Version *string `json:"Version,omitempty" name:"Version"`
		// 解决方案版本数据Tag创建时间

		VersionTagCreationTime *string `json:"VersionTagCreationTime,omitempty" name:"VersionTagCreationTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetActivePlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetActivePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppBriefInfo struct {
	// ID 应用自增ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// ApplicationInstanceID 应用实例ID

	ApplicationInstanceID *string `json:"ApplicationInstanceID,omitempty" name:"ApplicationInstanceID"`
	// ApplicationInstanceUUID 应用实例UUID

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
	// ApplicationID 应用ID

	ApplicationID *string `json:"ApplicationID,omitempty" name:"ApplicationID"`
	// ApplicationName 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// ApplicationInstanceName 应用实例名称

	ApplicationInstanceName *string `json:"ApplicationInstanceName,omitempty" name:"ApplicationInstanceName"`
	// ApplicationValues 应用values.yaml数据

	ApplicationValues *string `json:"ApplicationValues,omitempty" name:"ApplicationValues"`
	// AppComponentSet 应用模块信息列表

	AppComponentSet []*Component `json:"AppComponentSet,omitempty" name:"AppComponentSet"`
	// AppServerSet 应用服务器信息列表

	AppServerSet []*Server `json:"AppServerSet,omitempty" name:"AppServerSet"`
	// Location 应用部署位置

	Location *string `json:"Location,omitempty" name:"Location"`
}

type GetRuntimeSecretResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Content secret 信息，yaml格式

		Content *string `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRuntimeSecretResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRuntimeSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJobDAGRequest struct {
	*tchttp.BaseRequest

	//

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	//

	DAGNodeUUID *string `json:"DAGNodeUUID,omitempty" name:"DAGNodeUUID"`
	//

	Deep *int64 `json:"Deep,omitempty" name:"Deep"`
	//

	NodePath *string `json:"NodePath,omitempty" name:"NodePath"`
	//

	NodeRunColumns []*string `json:"NodeRunColumns,omitempty" name:"NodeRunColumns"`
}

func (r *DescribeJobDAGRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJobDAGRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TemplateInfo 运维模版详情信息

		TemplateInfo *Template `json:"TemplateInfo,omitempty" name:"TemplateInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductInfo struct {
	// ProductName 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// ProductCode 产品英文名

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// ProductID 产品ID

	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`
	// ProductInstanceSet 产品实例集合

	ProductInstanceSet []*ProductInstance `json:"ProductInstanceSet,omitempty" name:"ProductInstanceSet"`
}

type ListSiteProductInstancesRequest struct {
	*tchttp.BaseRequest

	//

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	//

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	//

	Orders []*OrderType `json:"Orders,omitempty" name:"Orders"`
	//

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListSiteProductInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSiteProductInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParseProductTemplateRequest struct {
	*tchttp.BaseRequest

	// 产品模板

	Template *string `json:"Template,omitempty" name:"Template"`
}

func (r *ParseProductTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ParseProductTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PlanDiff struct {
	// Type 配置类型：全局配置、产品配置、应用配置

	Type *string `json:"Type,omitempty" name:"Type"`
	// ProductInstanceUUID 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// ApplicationInstanceUUID 应用实例UUID

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
	// ServerUUID 服务器UUID

	ServerUUID *string `json:"ServerUUID,omitempty" name:"ServerUUID"`
	// ConfigUpdate 配置更新

	ConfigUpdate *bool `json:"ConfigUpdate,omitempty" name:"ConfigUpdate"`
}

type ProductInstanceOperationSheet struct {
	// 额外信息

	Annotations *string `json:"Annotations,omitempty" name:"Annotations"`
	// 应用列表

	Applications *string `json:"Applications,omitempty" name:"Applications"`
	// 云ID

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 运维单描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// ( primary key ) 自增主键

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 是否内置运维单

	IsInternal *bool `json:"IsInternal,omitempty" name:"IsInternal"`
	// 修改者

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 父子变更单关联关系，仅当此变更单为子变更单时有效

	OperationSheetRelation []*OperationSheetRelation `json:"OperationSheetRelation,omitempty" name:"OperationSheetRelation"`
	// 变更单具体的编排信息, 一个变更单默认只有一个编排模板，预留model_array用于扩展

	OperationSheetTemplate []*OperationSheetTemplate `json:"OperationSheetTemplate,omitempty" name:"OperationSheetTemplate"`
	// 规划ID

	PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
	// 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品运维单UUID

	ProductSheetID *string `json:"ProductSheetID,omitempty" name:"ProductSheetID"`
	// 产品UUID

	ProductUUID *string `json:"ProductUUID,omitempty" name:"ProductUUID"`
	// 产品版本制品Tag

	ProductVersionArtifactTag *string `json:"ProductVersionArtifactTag,omitempty" name:"ProductVersionArtifactTag"`
	// 产品版本名称

	ProductVersionName *string `json:"ProductVersionName,omitempty" name:"ProductVersionName"`
	// 产品版本UUID

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	// 区域UUID

	RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
	// 风险信息

	Risk *string `json:"Risk,omitempty" name:"Risk"`
	// 场景

	Scene *string `json:"Scene,omitempty" name:"Scene"`
	// 产品实例运维单UUID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// 运维单名称

	SheetName *string `json:"SheetName,omitempty" name:"SheetName"`
	// 运维单类型

	SheetType *string `json:"SheetType,omitempty" name:"SheetType"`
	// 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// 运维单来源:导入，人工创建

	Source *string `json:"Source,omitempty" name:"Source"`
	// 运维单状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type GetMaterialCleanupSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// AutoCleanup 自动清理

		AutoCleanup *bool `json:"AutoCleanup,omitempty" name:"AutoCleanup"`
		// StartTime 自动清理开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// DefaultReservedInstall 默认保留安装包数量

		DefaultReservedInstall *int64 `json:"DefaultReservedInstall,omitempty" name:"DefaultReservedInstall"`
		// DefaultReservedFix 默认保留补丁包数量

		DefaultReservedFix *int64 `json:"DefaultReservedFix,omitempty" name:"DefaultReservedFix"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMaterialCleanupSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialCleanupSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAtomicRenderDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Params 原子操作渲染后参数

		Params []*KVPair `json:"Params,omitempty" name:"Params"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAtomicRenderDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAtomicRenderDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationPackageDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		ApplicationPackage *ApplicationPackage `json:"ApplicationPackage,omitempty" name:"ApplicationPackage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationPackageDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationPackageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGlobalConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CloudID 云ID

		CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
		// PLanID 规划ID

		PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGlobalConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGlobalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SheetAttachment struct {
	//

	Name *string `json:"Name,omitempty" name:"Name"`
	//

	Format *string `json:"Format,omitempty" name:"Format"`
	//

	Description *string `json:"Description,omitempty" name:"Description"`
	//

	Size *int64 `json:"Size,omitempty" name:"Size"`
	//

	DownloadURL *string `json:"DownloadURL,omitempty" name:"DownloadURL"`
}

type DescribeProductRuntimeConfigRequest struct {
	*tchttp.BaseRequest

	//

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	//

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	//

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	//

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
}

func (r *DescribeProductRuntimeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductRuntimeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportAtomicRequest struct {
	*tchttp.BaseRequest

	// AtomicPackageDownloadUrl 原子操作包远程下载地址，可单独使用，前端界面导入使用（第一优先级）

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
	// AtomicPackageFilePath 原子操作包本地目录，可单独使用，后段接口导入使用（第二优先级）

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// AtomicPackageData 原子操作包二进制包内容，base64字符串格式，可单独使用，备用（优先级最低）

	Package *string `json:"Package,omitempty" name:"Package"`
}

func (r *ImportAtomicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportAtomicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAppPlanRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// CloudID 云UUID

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// ApplicationInstanceID 应用实例ID

	ApplicationInstanceID *string `json:"ApplicationInstanceID,omitempty" name:"ApplicationInstanceID"`
	// ApplicationInstanceUUID 应用实例UUID

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
	// Config 修改后的应用配置

	Config *string `json:"Config,omitempty" name:"Config"`
	// Modifier 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
}

func (r *ModifyAppPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAppPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationSheetRequest struct {
	*tchttp.BaseRequest

	//

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	//

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	//

	RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
	//

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	//

	RegionNameCN *string `json:"RegionNameCN,omitempty" name:"RegionNameCN"`
	//

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	//

	Template *string `json:"Template,omitempty" name:"Template"`
	//

	Type *string `json:"Type,omitempty" name:"Type"`
	// TypeDetail 单子类型细分，例如是应用部署or产品部署or工具组件部署or原子操作部署

	TypeDetail *string `json:"TypeDetail,omitempty" name:"TypeDetail"`
	//

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	//

	Creator *string `json:"Creator,omitempty" name:"Creator"`
}

func (r *CreateOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// PlanID 规划ID

		PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
		// CloudID 云ID

		CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
		// ProductSet 产品应用实例集合

		ProductSet []*ProductInfo `json:"ProductSet,omitempty" name:"ProductSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnnotationsVar struct {
	//

	ProductUUID *string `json:"ProductUUID,omitempty" name:"ProductUUID"`
	// sheet from jiguang: ProductCode = "", ProductName is actually ProductCode
	// sheet created from local: ProductCode = ProductCode(like 'cbs'), ProductName is ProductName(like '云硬盘（CBS）')

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	//

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	//

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	//

	ProductVersionName *string `json:"ProductVersionName,omitempty" name:"ProductVersionName"`
	//

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	//

	ProductInstanceName *string `json:"ProductInstanceName,omitempty" name:"ProductInstanceName"`
	//

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	//

	ProductVersionArtifactTag *string `json:"ProductVersionArtifactTag,omitempty" name:"ProductVersionArtifactTag"`
	//

	Applications []*ApplicationInfo `json:"Applications,omitempty" name:"Applications"`
	// 缺陷单ID

	IssueID *string `json:"IssueID,omitempty" name:"IssueID"`
	// 缺陷单名称

	Title *string `json:"Title,omitempty" name:"Title"`
	// 缺陷严重程度

	SeverityLevel *string `json:"SeverityLevel,omitempty" name:"SeverityLevel"`
	// 发现缺陷的产品版本

	ProductVersion *string `json:"ProductVersion,omitempty" name:"ProductVersion"`
	// 发现缺陷的解决方案版本

	SolutionVersion *string `json:"SolutionVersion,omitempty" name:"SolutionVersion"`
	// 解决问题单涉及的解决方案&架构列表，即创建Bugfix变更单所需要应用的范围

	Solutions []*IssueSolution `json:"Solutions,omitempty" name:"Solutions"`
	// 问题评估涉及的产品以及应用信息

	Products []*ProductDetail `json:"Products,omitempty" name:"Products"`
	//

	PageSolutions *string `json:"PageSolutions,omitempty" name:"PageSolutions"`
	//

	PageProducts *string `json:"PageProducts,omitempty" name:"PageProducts"`
	//

	PageArchs *string `json:"PageArchs,omitempty" name:"PageArchs"`
	// 问题单基本信息

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	//

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
	//

	TaskID *int64 `json:"TaskID,omitempty" name:"TaskID"`
	//

	RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
	//

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	//

	RegionNameCN *string `json:"RegionNameCN,omitempty" name:"RegionNameCN"`
	//

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	//

	PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
	//

	SourceSolutionVersion *string `json:"SourceSolutionVersion,omitempty" name:"SourceSolutionVersion"`
	//

	TargetSolutionVersion *string `json:"TargetSolutionVersion,omitempty" name:"TargetSolutionVersion"`
	//

	TargetSolutionDataTag *string `json:"TargetSolutionDataTag,omitempty" name:"TargetSolutionDataTag"`
	//

	AssociatedSiteUUID *string `json:"AssociatedSiteUUID,omitempty" name:"AssociatedSiteUUID"`
	//

	AssociatedSiteName *string `json:"AssociatedSiteName,omitempty" name:"AssociatedSiteName"`
	//

	AssociatedClientUUID *string `json:"AssociatedClientUUID,omitempty" name:"AssociatedClientUUID"`
	//

	AssociatedClientName *string `json:"AssociatedClientName,omitempty" name:"AssociatedClientName"`
	//

	PatchID *string `json:"PatchID,omitempty" name:"PatchID"`
	//

	Orchestrated *string `json:"Orchestrated,omitempty" name:"Orchestrated"`
}

type SiteProductInstance struct {
	//

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	//

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	//

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	//

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	//

	Status *string `json:"Status,omitempty" name:"Status"`
	// ProductID 产品ID

	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`
	// ProductName 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// ProductCode 产品英文名

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// ProductVersionID 产品版本ID

	ProductVersionID *string `json:"ProductVersionID,omitempty" name:"ProductVersionID"`
	// ProductVersionName 产品版本名称

	ProductVersionName *string `json:"ProductVersionName,omitempty" name:"ProductVersionName"`
	// ProductInstanceUUID 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// ProductInstanceName 产品实例名称

	ProductInstanceName *string `json:"ProductInstanceName,omitempty" name:"ProductInstanceName"`
	// ProductVersionDataTag 产品版本的数据Tag

	ProductVersionDataTag *string `json:"ProductVersionDataTag,omitempty" name:"ProductVersionDataTag"`
	// ProductCategory 产品分类

	ProductCategory *string `json:"ProductCategory,omitempty" name:"ProductCategory"`
	// Description 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// Location 位置信息

	Location *ProductLocation `json:"Location,omitempty" name:"Location"`
}

type ConvertPipelineFormatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Pipeline 编排内容

		Pipeline *PipelineContent `json:"Pipeline,omitempty" name:"Pipeline"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConvertPipelineFormatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConvertPipelineFormatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SupportedAppInfo struct {
	// ApplicationInstanceUUID 应用实例UUID

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
	// ApplicationInstanceName 应用实例名称

	ApplicationInstanceName *string `json:"ApplicationInstanceName,omitempty" name:"ApplicationInstanceName"`
	// ProductInstanceUUID 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
}

type DescribeAppRuntimeInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// AppRuntimeInfos 应用运行时信息

		AppRuntimeInfos []*AppRuntimeInfo `json:"AppRuntimeInfos,omitempty" name:"AppRuntimeInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppRuntimeInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppRuntimeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAtomicRequest struct {
	*tchttp.BaseRequest

	// AtomicList 导出的原子操作key列表

	AtomicList []*AtomicKey `json:"AtomicList,omitempty" name:"AtomicList"`
}

func (r *ExportAtomicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAtomicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProductTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProductTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProductTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppComponentInfo struct {
	// ApplicationInstanceUUID 应用实例UUID

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
	// ApplicationInstanceName 应用实例名称

	ApplicationInstanceName *string `json:"ApplicationInstanceName,omitempty" name:"ApplicationInstanceName"`
	// ApplicationName 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// ComponentSet 模块列表

	ComponentSet []*ComponentInfo `json:"ComponentSet,omitempty" name:"ComponentSet"`
}

type ListLatestApplicationPackagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		ApplicationPackages []*ApplicationPackage `json:"ApplicationPackages,omitempty" name:"ApplicationPackages"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListLatestApplicationPackagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLatestApplicationPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCommonOperationSheetRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
}

func (r *DescribeCommonOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCommonOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOtaCheckConfigRequest struct {
	*tchttp.BaseRequest

	// 自动检查配置 1：开启 0：关闭

	CheckConfig *string `json:"CheckConfig,omitempty" name:"CheckConfig"`
	// 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
}

func (r *UpdateOtaCheckConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOtaCheckConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Material struct {
	// 物料名称

	MaterialName *string `json:"MaterialName,omitempty" name:"MaterialName"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品uuid

	ProductUUID *string `json:"ProductUUID,omitempty" name:"ProductUUID"`
	// 产品版本uuid

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	// 产品版本名称

	ProductVersionName *string `json:"ProductVersionName,omitempty" name:"ProductVersionName"`
	// 解决方案名称

	SolutionName *string `json:"SolutionName,omitempty" name:"SolutionName"`
	// 解决方案uuid

	SolutionUUID *string `json:"SolutionUUID,omitempty" name:"SolutionUUID"`
	// 解决方案版本UUID

	SolutionVersionUUID *string `json:"SolutionVersionUUID,omitempty" name:"SolutionVersionUUID"`
	// 解决方案版本名称

	SolutionVersionName *string `json:"SolutionVersionName,omitempty" name:"SolutionVersionName"`
	// 物料版本UUID

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
	// 物料类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 物料版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 物料架构

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	// 物料来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 物料标签

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
	// 发布日志

	ReleaseNote *string `json:"ReleaseNote,omitempty" name:"ReleaseNote"`
	//

	Status *string `json:"Status,omitempty" name:"Status"`
}

type PlanVersionState struct {
	// PlanID 规划ID

	PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
	// PlanState 规划状态

	PlanState *string `json:"PlanState,omitempty" name:"PlanState"`
	// Tag 规划快照

	Tag *int64 `json:"Tag,omitempty" name:"Tag"`
}

type AtomicJsonToYamlRequest struct {
	*tchttp.BaseRequest

	//

	Atomic *JsonAtomic `json:"Atomic,omitempty" name:"Atomic"`
}

func (r *AtomicJsonToYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AtomicJsonToYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckSheetTemplateRequest struct {
	*tchttp.BaseRequest

	// SheetID 运维单ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// Usage 模板用途

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// FilterApp 过滤模板

	FilterApp *bool `json:"FilterApp,omitempty" name:"FilterApp"`
}

func (r *CheckSheetTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSheetTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 创建人

		Creator *string `json:"Creator,omitempty" name:"Creator"`
		// 删除时间

		DeletedAt *int64 `json:"DeletedAt,omitempty" name:"DeletedAt"`
		// ( primary key ) 自增ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// 最大的修改历史ID

		MaxPlanHistoryID *int64 `json:"MaxPlanHistoryID,omitempty" name:"MaxPlanHistoryID"`
		// 规划ID

		PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
		// 局点UUID

		SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
		// 规划tag号

		TagNum *int64 `json:"TagNum,omitempty" name:"TagNum"`
		// 更新时间

		UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePlanTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSubOperationSheetsRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
}

func (r *ListSubOperationSheetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSubOperationSheetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AtomicYamlToJsonRequest struct {
	*tchttp.BaseRequest

	//

	AtomicYaml *string `json:"AtomicYaml,omitempty" name:"AtomicYaml"`
}

func (r *AtomicYamlToJsonRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AtomicYamlToJsonRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationTemplateRequest struct {
	*tchttp.BaseRequest

	// TemplateInfo 要新建的运维模版信息

	TemplateInfo *Template `json:"TemplateInfo,omitempty" name:"TemplateInfo"`
}

func (r *CreateOperationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationSheetTemplate struct {
	// 选择的应用实例列表

	ApplicationInstanceList *string `json:"ApplicationInstanceList,omitempty" name:"ApplicationInstanceList"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// Tcs集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// ( primary key ) 自增主键

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 变更单的待编排项

	OrchestratingItems *string `json:"OrchestratingItems,omitempty" name:"OrchestratingItems"`
	// 模板自定义参数

	Parameters *string `json:"Parameters,omitempty" name:"Parameters"`
	// 产品实例运维单UUID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// 模板内容

	Template *string `json:"Template,omitempty" name:"Template"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 运维单模板用途（部署，升级，变更，回滚）

	Usage *string `json:"Usage,omitempty" name:"Usage"`
}

type CreateCommonOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCommonOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCommonOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDeployedApplicationDetailsRequest struct {
	*tchttp.BaseRequest

	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// SubsystemID 子系统名称

	SubsystemInstanceName *string `json:"SubsystemInstanceName,omitempty" name:"SubsystemInstanceName"`
	//

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	//

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	//

	Orders []*OrderType `json:"Orders,omitempty" name:"Orders"`
	//

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListDeployedApplicationDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDeployedApplicationDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductMaterialsRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// ProductUUID 产品UUID

	ProductUUID *string `json:"ProductUUID,omitempty" name:"ProductUUID"`
	// Filter 物料包的类型

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
	// Offset 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit 分页限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListProductMaterialsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductMaterialsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySystemConfigRequest struct {
	*tchttp.BaseRequest

	// SystemConfigs 系统配置

	SystemConfigs []*SystemConfig `json:"SystemConfigs,omitempty" name:"SystemConfigs"`
}

func (r *ModifySystemConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySystemConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SkipNodeRequest struct {
	*tchttp.BaseRequest

	//

	DagJobUUID *string `json:"DagJobUUID,omitempty" name:"DagJobUUID"`
	//

	DagNodeUUID *string `json:"DagNodeUUID,omitempty" name:"DagNodeUUID"`
}

func (r *SkipNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SkipNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		JobID *string `json:"JobID,omitempty" name:"JobID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportOperationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		OperationTemplateList []*OperationTemplateKey `json:"OperationTemplateList,omitempty" name:"OperationTemplateList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportOperationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportOperationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProjectZone struct {
	// 中文名称

	ChineseName *string `json:"ChineseName,omitempty" name:"ChineseName"`
	//

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// ( primary key )

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 关联云拓扑地域的ID

	RegionID *int64 `json:"RegionID,omitempty" name:"RegionID"`
	// 关联的真实可用区

	RelatedZoneUUID *string `json:"RelatedZoneUUID,omitempty" name:"RelatedZoneUUID"`
	// 角色

	RoleType *string `json:"RoleType,omitempty" name:"RoleType"`
	// 安灯局点状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 类型, 0表示真实可用区、1表示虚拟可用区

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	//

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 地区ID

	ZoneID *string `json:"ZoneID,omitempty" name:"ZoneID"`
	// 地区名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type GeneratePipelineRunResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		PipelineRun *string `json:"PipelineRun,omitempty" name:"PipelineRun"`
		//

		DAG *DAG `json:"DAG,omitempty" name:"DAG"`
		//

		WarnMessage *string `json:"WarnMessage,omitempty" name:"WarnMessage"`
		//

		MissApplicationList []*string `json:"MissApplicationList,omitempty" name:"MissApplicationList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneratePipelineRunResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GeneratePipelineRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAtomicRequest struct {
	*tchttp.BaseRequest

	// Kind 原子操作类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// Solution 原子操作适用的解决方案，存在标签中

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// OperationScenes 原子操作适用的运维长江，存在标签中

	OperationScenes *string `json:"OperationScenes,omitempty" name:"OperationScenes"`
	// Labels 原子操作标签

	Label []*KVPair `json:"Label,omitempty" name:"Label"`
	// Offset 分页用

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit 分页用

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Filter 类似数据引擎的构造filter, 用于原子操作筛选

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
}

func (r *ListAtomicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAtomicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductTemplateIDResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// DefaultID 模版 ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// TemplateList 可能适配的运维模版信息

		TemplateList []*Template `json:"TemplateList,omitempty" name:"TemplateList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductTemplateIDResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductTemplateIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperationTemplateRequest struct {
	*tchttp.BaseRequest

	// TemplateInfo 要修改的运维模版信息

	TemplateInfo *Template `json:"TemplateInfo,omitempty" name:"TemplateInfo"`
}

func (r *ModifyOperationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductBaseInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ProductCode 产品Code

		ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
		// Status 产品运行状态，

		Healthy *string `json:"Healthy,omitempty" name:"Healthy"`
		// Dictionaries 产品标签列表

		Dictionaries []*Dictionary `json:"Dictionaries,omitempty" name:"Dictionaries"`
		// Version 部署版本

		Version *string `json:"Version,omitempty" name:"Version"`
		// Cluster 部署集群

		Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductBaseInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductBaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductStatusCount struct {
	//

	Status *string `json:"Status,omitempty" name:"Status"`
	//

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type SubsystemInstanceInfo struct {
	//

	SubsystemInstanceID *string `json:"SubsystemInstanceID,omitempty" name:"SubsystemInstanceID"`
	//

	SubSystemInstanceName *string `json:"SubSystemInstanceName,omitempty" name:"SubSystemInstanceName"`
	//

	ApplicationInstanceSet []*ApplicationInstanceInfo `json:"ApplicationInstanceSet,omitempty" name:"ApplicationInstanceSet"`
}

type Application struct {
	// 应用名且全局唯一，默认新接入的应用需要以workspace的名字作为前缀开头，workspace无法修改，应用名也无法修改，

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// 和模板管理确认，模板那边使用模板名作为主键，全局唯一，这里模板id我们创建的时候，就直接调用模板管理的api，以应用名创建一个模板，保存在这里

	ApplicationTemplateID *string `json:"ApplicationTemplateID,omitempty" name:"ApplicationTemplateID"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 删除时间

	DeletedAt *int64 `json:"DeletedAt,omitempty" name:"DeletedAt"`
	// ( primary key ) 自增主键

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// Applicaton级别的流水线配置

	PipelineConfigUUID *string `json:"PipelineConfigUUID,omitempty" name:"PipelineConfigUUID"`
	// 应用的状态 1启用 0废弃，其他状态后续扩展

	State *int64 `json:"State,omitempty" name:"State"`
	// tenant和application是1对多的关系

	TenantUUID *string `json:"TenantUUID,omitempty" name:"TenantUUID"`
	// uuid键，用户后续和其他数据的关联

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type ModifyGlobalConfigRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// CloudID 云UUID

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// Config 全局配置信息

	Config *string `json:"Config,omitempty" name:"Config"`
	// Modifier 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
}

func (r *ModifyGlobalConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGlobalConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckMaterialException struct {
	// 产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// ProductInstanceName 产品实例名称

	ProductInstanceName *string `json:"ProductInstanceName,omitempty" name:"ProductInstanceName"`
	// SubsystemCode 子系统Code

	SubsystemCode *string `json:"SubsystemCode,omitempty" name:"SubsystemCode"`
	// SubsystemName 子系统名字

	SubsystemName *string `json:"SubsystemName,omitempty" name:"SubsystemName"`
	// ApplicationName 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// ApplicationID 应用ID

	ApplicationID *string `json:"ApplicationID,omitempty" name:"ApplicationID"`
	// ApplicationCode 应用Code

	ApplicationCode *string `json:"ApplicationCode,omitempty" name:"ApplicationCode"`
	//

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	//

	ApplicationVersion *string `json:"ApplicationVersion,omitempty" name:"ApplicationVersion"`
	// CurrentVersion 当前版本, 如果有多个实例版本不同, 逗号分隔

	CurrentVersion *string `json:"CurrentVersion,omitempty" name:"CurrentVersion"`
	// TargetVersion 目标版本

	TargetVersion *string `json:"TargetVersion,omitempty" name:"TargetVersion"`
	// ExceptionCode 异常代码, MaterialNotFound(物料不存在)/MaterialVersionTooOld(目标版本小于环境内版本)

	ExceptionCode *string `json:"ExceptionCode,omitempty" name:"ExceptionCode"`
}

type DescribeSheetAttachmentsRequest struct {
	*tchttp.BaseRequest

	// SheetID 运维单 ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
}

func (r *DescribeSheetAttachmentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSheetAttachmentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Options struct {
	//

	ShouldNotRender *bool `json:"ShouldNotRender,omitempty" name:"ShouldNotRender"`
	//

	BehaviorOnRenderError *string `json:"BehaviorOnRenderError,omitempty" name:"BehaviorOnRenderError"`
	//

	DefaultValueOnRenderError *string `json:"DefaultValueOnRenderError,omitempty" name:"DefaultValueOnRenderError"`
	// Format indicates the value's format, like json/yaml/text

	Format *string `json:"Format,omitempty" name:"Format"`
}

type DescribePlanApplicationTreeRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// ProductInstanceUUID 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// ApplicationUUIDList 应用UUID列表

	ApplicationUUIDList []*string `json:"ApplicationUUIDList,omitempty" name:"ApplicationUUIDList"`
	// ApplicationInstanceUUIDList 应用实例UUID列表

	ApplicationInstanceUUIDList []*string `json:"ApplicationInstanceUUIDList,omitempty" name:"ApplicationInstanceUUIDList"`
	// ShowLevel 展示层级，支持product/application(默认)

	ShowLevel *string `json:"ShowLevel,omitempty" name:"ShowLevel"`
	// 指定Region/Zone查询

	Location *LocationInfo `json:"Location,omitempty" name:"Location"`
	// PlanID 规划ID

	PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
	// PlanState 规划状态

	PlanState *string `json:"PlanState,omitempty" name:"PlanState"`
	// Tag 规划快照

	Tag *int64 `json:"Tag,omitempty" name:"Tag"`
}

func (r *DescribePlanApplicationTreeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePlanApplicationTreeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialProductDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Products 产品列表

		ProductID []*MaterialProduct `json:"ProductID,omitempty" name:"ProductID"`
		// SolutionVersionName 解决方案版本名称

		SolutionVersionName *string `json:"SolutionVersionName,omitempty" name:"SolutionVersionName"`
		// SolutionVersionUUID 解决方案版本 UUID

		SolutionVersionUUID *string `json:"SolutionVersionUUID,omitempty" name:"SolutionVersionUUID"`
		// MaterialName 物料名称

		MaterialName *string `json:"MaterialName,omitempty" name:"MaterialName"`
		// Arch 物料架构

		Arch *string `json:"Arch,omitempty" name:"Arch"`
		// Version 物料版本

		Version *string `json:"Version,omitempty" name:"Version"`
		// Label 物料标签

		Label *string `json:"Label,omitempty" name:"Label"`
		// Source 物料来源

		Source *string `json:"Source,omitempty" name:"Source"`
		// MaterialUUID 物料 UUID

		MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
		// MaterialID 物料 ID

		MaterialID *int64 `json:"MaterialID,omitempty" name:"MaterialID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialProductDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialProductDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JsonAtomic struct {
	// UUID 原子操作uuid

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// Name 原子操作名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Version 原子操作版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// Description 原子操作描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// Labels 原子操作标签

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// Annotations 原子操作Annotations

	Annotations []*KVPair `json:"Annotations,omitempty" name:"Annotations"`
	// Params 原子操作入参申明列表

	Params []*AtomicParam `json:"Params,omitempty" name:"Params"`
	// Outputs 原子操作出参列表

	Outputs []*KVPair `json:"Outputs,omitempty" name:"Outputs"`
	// Steps 原子操作步骤列表

	Steps []*AtomicStep `json:"Steps,omitempty" name:"Steps"`
	// Files 原子操作关联文件列表

	Files []*AtomicFile `json:"Files,omitempty" name:"Files"`
	// CreateTime 原子操作创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// CreateBy 原子操作创建者

	CreateBy *string `json:"CreateBy,omitempty" name:"CreateBy"`
	// UpdateTime 原子操作修改时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// UpdateBy 原子操作修改着

	UpdateBy *string `json:"UpdateBy,omitempty" name:"UpdateBy"`
	// DeleteTime 原子操作删除时间

	DeleteTime *string `json:"DeleteTime,omitempty" name:"DeleteTime"`
}

type ModifyPlanServerInfoRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// CloudID 云UUID

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// ServerInfoSet 修改的服务器信息列表

	ServerInfoSet []*ServerInfo `json:"ServerInfoSet,omitempty" name:"ServerInfoSet"`
	// Modifier 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
}

func (r *ModifyPlanServerInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPlanServerInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartDevelopmentPhaseApplicationDeployRequest struct {
	*tchttp.BaseRequest

	// ApplicationType 应用的类型

	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`
	// ApplicationName 应用的名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// ApplicationUUID 应用的 UUID

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	// ApplicationVersion 应用的版本

	ApplicationVersion *string `json:"ApplicationVersion,omitempty" name:"ApplicationVersion"`
	// RegionID 应用的 RegionID

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
}

func (r *StartDevelopmentPhaseApplicationDeployRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartDevelopmentPhaseApplicationDeployRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppRuntimeNameInfo struct {
	// ApplicationInstanceUUID 应用实例UUID

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
	// ApplicationInstanceName 规划的应用实例名称

	ApplicationInstanceName *string `json:"ApplicationInstanceName,omitempty" name:"ApplicationInstanceName"`
	// ApplicationRuntimeName 运行态实例名称

	ApplicationRuntimeName *string `json:"ApplicationRuntimeName,omitempty" name:"ApplicationRuntimeName"`
	// ApplicationName 应用名称，用于模糊查询

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
}

type CopyProductInstanceOperationSheetRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点ID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// SheetID 变更单ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// RegionID 区域ID

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	// SheetName 变更单名称

	SheetName *string `json:"SheetName,omitempty" name:"SheetName"`
	// Description 变更单描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// Operator 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *CopyProductInstanceOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyProductInstanceOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportOperationSheetPackageApplication struct {
	//

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	//

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	//

	ProductVersionArtifactTag *int64 `json:"ProductVersionArtifactTag,omitempty" name:"ProductVersionArtifactTag"`
	//

	PackageUUID *string `json:"PackageUUID,omitempty" name:"PackageUUID"`
	//

	ApplicationVersion *string `json:"ApplicationVersion,omitempty" name:"ApplicationVersion"`
}

type ModifyProductInstanceOperationSheetRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点ID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// SheetID 变更单ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// SheetName 变更单名称

	SheetName *string `json:"SheetName,omitempty" name:"SheetName"`
	// Description 运维单描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// ApplicationUUIDList 运维单关联应用UUID列表

	ApplicationUUIDList []*string `json:"ApplicationUUIDList,omitempty" name:"ApplicationUUIDList"`
	// Operator 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// Status 变更单状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyProductInstanceOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProductInstanceOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckProductInstanceOperationSheetTemplateAndAtomicRequest struct {
	*tchttp.BaseRequest

	// 变更模板的yaml

	Template *string `json:"Template,omitempty" name:"Template"`
	// TemplateBase64 通过直接传输变更YAML的base64数据

	TemplateBase64 *string `json:"TemplateBase64,omitempty" name:"TemplateBase64"`
}

func (r *CheckProductInstanceOperationSheetTemplateAndAtomicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckProductInstanceOperationSheetTemplateAndAtomicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanTagRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// Creator 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
}

func (r *CreatePlanTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Status 导入状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Message 导入消息

		Message *string `json:"Message,omitempty" name:"Message"`
		// Result 导入结果

		Result *ImportPlanResult `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialOperationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Total *int64 `json:"Total,omitempty" name:"Total"`
		//

		OperationTemplateList []*Template `json:"OperationTemplateList,omitempty" name:"OperationTemplateList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialOperationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialOperationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PauseNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PauseNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PauseNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParseMaterialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 物料信息

		Material *MaterialData `json:"Material,omitempty" name:"Material"`
		// 物料解析提示信息

		ImportMessage *string `json:"ImportMessage,omitempty" name:"ImportMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ParseMaterialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ParseMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductBaseInfoRequest struct {
	*tchttp.BaseRequest

	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	//

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
}

func (r *DescribeProductBaseInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductBaseInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationValuesRequest struct {
	*tchttp.BaseRequest

	// 制品包UUID

	ApplicationPackagesUUID *string `json:"ApplicationPackagesUUID,omitempty" name:"ApplicationPackagesUUID"`
}

func (r *GetApplicationValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJobDAGResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Status *string `json:"Status,omitempty" name:"Status"`
		//

		ProgressCount *ProgressCount `json:"ProgressCount,omitempty" name:"ProgressCount"`
		//

		DAG *DAG `json:"DAG,omitempty" name:"DAG"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJobDAGResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJobDAGResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialProductApplicationInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Total *int64 `json:"Total,omitempty" name:"Total"`
		//

		ApplicationInstanceList []*ApplicationInstance `json:"ApplicationInstanceList,omitempty" name:"ApplicationInstanceList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialProductApplicationInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialProductApplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationTemplatePair struct {
	//

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	//

	TemplateInfo *Template `json:"TemplateInfo,omitempty" name:"TemplateInfo"`
}

type AtomicJsonToYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		AtomicYaml *string `json:"AtomicYaml,omitempty" name:"AtomicYaml"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AtomicJsonToYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AtomicJsonToYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyCommonOperationSheetRequest struct {
	*tchttp.BaseRequest

	// 被拷贝的变更单SheetID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// 拷贝场景

	Scene *string `json:"Scene,omitempty" name:"Scene"`
	// 拷贝的业务方

	Source *string `json:"Source,omitempty" name:"Source"`
	// 拷贝时的标识名称

	CopyID *string `json:"CopyID,omitempty" name:"CopyID"`
}

func (r *CopyCommonOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyCommonOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSiteOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Sheet 变更单

		Sheet *CommonOperationSheet `json:"Sheet,omitempty" name:"Sheet"`
		// Template 编排内容

		Template *string `json:"Template,omitempty" name:"Template"`
		// Templates 编排内容列表

		Templates []*OperationSheetTemplate `json:"Templates,omitempty" name:"Templates"`
		// Children 子变更变单

		Children []*CommonOperationSheet `json:"Children,omitempty" name:"Children"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSiteOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSiteOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSiteInfoDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SiteID 局点 ID

		SiteID *int64 `json:"SiteID,omitempty" name:"SiteID"`
		// SiteUUID 局点 UUID

		SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
		// SiteName 局点名称

		SiteName *string `json:"SiteName,omitempty" name:"SiteName"`
		// ConnectURL 局点连接地址

		ConnectURL *string `json:"ConnectURL,omitempty" name:"ConnectURL"`
		// Description 局点描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// CloudID 云 ID

		CloudID *int64 `json:"CloudID,omitempty" name:"CloudID"`
		// CloudUUID 云 UUID

		CloudUUID *string `json:"CloudUUID,omitempty" name:"CloudUUID"`
		// MaterialID 物料 ID

		MaterialID *int64 `json:"MaterialID,omitempty" name:"MaterialID"`
		// MaterialName 物料名称

		MaterialName *string `json:"MaterialName,omitempty" name:"MaterialName"`
		// MaterialUUID 物料 UUID

		MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
		// SolutionVersionName 解决方案版本名称

		SolutionVersionName *string `json:"SolutionVersionName,omitempty" name:"SolutionVersionName"`
		// SolutionVersionUUID 解决方案版本 UUID

		SolutionVersionUUID *string `json:"SolutionVersionUUID,omitempty" name:"SolutionVersionUUID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectSiteInfoDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectSiteInfoDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProgressCount struct {
	//

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	//

	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
}

type SubsystemInfo struct {
	// SubsystemID 子系统ID

	SubsystemID *string `json:"SubsystemID,omitempty" name:"SubsystemID"`
	// SubsystemName 子系统名称

	SubsystemName *string `json:"SubsystemName,omitempty" name:"SubsystemName"`
	// SubsystemCode 子系统Code

	SubsystemCode *string `json:"SubsystemCode,omitempty" name:"SubsystemCode"`
	// ApplicationSet 应用列表

	ApplicationSet []*ApplicationInfo `json:"ApplicationSet,omitempty" name:"ApplicationSet"`
}

type DescribeJobOverviewRequest struct {
	*tchttp.BaseRequest

	//

	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeJobOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJobOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportProductInstanceOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportProductInstanceOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportProductInstanceOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartOPMControllerRequest struct {
	*tchttp.BaseRequest
}

func (r *RestartOPMControllerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartOPMControllerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceTreeProduct struct {
	// ProductCode 产品英文名

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// ProductName 产品中文名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// ProductInstanceUUID 产品实例 UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// ProductInstanceID 产品实例 ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// ProductInstanceName 产品实例名称

	ProductInstanceName *string `json:"ProductInstanceName,omitempty" name:"ProductInstanceName"`
	// SubsystemInstanceSet 子系统实例列表

	SubsystemInstanceSet []*InstanceTreeSubsystem `json:"SubsystemInstanceSet,omitempty" name:"SubsystemInstanceSet"`
}

type ImportPlanRequest struct {
	*tchttp.BaseRequest

	// TaskID 任务 ID

	TaskID *int64 `json:"TaskID,omitempty" name:"TaskID"`
	// Data 规划包

	Data *string `json:"Data,omitempty" name:"Data"`
	// Creator

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// SiteUUID 生效局点

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
}

func (r *ImportPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParseProductDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品数据

		ProductData *MaterialData `json:"ProductData,omitempty" name:"ProductData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ParseProductDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ParseProductDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartSiteOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		JobID *string `json:"JobID,omitempty" name:"JobID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartSiteOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartSiteOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PlanProductInstance struct {
	// 规划中产品实例配置,yaml格式

	Config *string `json:"Config,omitempty" name:"Config"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 删除时间

	DeletedAt *int64 `json:"DeletedAt,omitempty" name:"DeletedAt"`
	// ( primary key ) 产品实例自增ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 产品部署位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 关联的规划ID

	PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
	// 产品实例下的子系统实例

	PlanSubsystemInstance []*PlanSubsystemInstance `json:"PlanSubsystemInstance,omitempty" name:"PlanSubsystemInstance"`
	// 产品英文代号

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品ID

	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`
	// 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// 产品实例名称

	ProductInstanceName *string `json:"ProductInstanceName,omitempty" name:"ProductInstanceName"`
	// 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品版本的数据tag

	ProductVersionDataTag *string `json:"ProductVersionDataTag,omitempty" name:"ProductVersionDataTag"`
	// 产品版本的ID

	ProductVersionID *string `json:"ProductVersionID,omitempty" name:"ProductVersionID"`
	// 产品版本名称

	ProductVersionName *string `json:"ProductVersionName,omitempty" name:"ProductVersionName"`
	// 产品实例的类型:standard/site

	Type *string `json:"Type,omitempty" name:"Type"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type ModifyAtomicRenderDataRequest struct {
	*tchttp.BaseRequest

	// CloudID 云ID

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// RegionUUID 地域UUID

	RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
	// 产品信息如果不提供，则无法渲染产品级别的配置
	// ProductInstanceID或者ProductInstanceUUID 提供其中一个即可。
	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// ProductInstanceUUID 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// ApplicationInstanceID或者ApplicationName 提供其中一个即可
	// ApplicationInstanceID 应用实例ID

	ApplicationInstanceID *string `json:"ApplicationInstanceID,omitempty" name:"ApplicationInstanceID"`
	// ApplicationName 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// Params 原子操作入参

	Params []*KVPair `json:"Params,omitempty" name:"Params"`
}

func (r *ModifyAtomicRenderDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAtomicRenderDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartProductInstanceOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
		//

		JobID *string `json:"JobID,omitempty" name:"JobID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartProductInstanceOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartProductInstanceOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品 ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// 产品 UUID

		UUID *string `json:"UUID,omitempty" name:"UUID"`
		// 产品的中文名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 产品 Code

		Code *string `json:"Code,omitempty" name:"Code"`
		// 产品标签

		Label *string `json:"Label,omitempty" name:"Label"`
		// 产品描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// 产品来源

		BuiltIn *int64 `json:"BuiltIn,omitempty" name:"BuiltIn"`
		// 最新版本

		LatestVersion *string `json:"LatestVersion,omitempty" name:"LatestVersion"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAtomicRequest struct {
	*tchttp.BaseRequest

	// UUID 原子操作uuid，可单独使用，第一优先级查询

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// AtomicName 原子操作名称，和version一起使用，第二优先级查询

	AtomicName *string `json:"AtomicName,omitempty" name:"AtomicName"`
	// Version 原子操作版本，和name一起使用，第二优先级查询

	AtomicVersion *string `json:"AtomicVersion,omitempty" name:"AtomicVersion"`
	// AtomicKey 需要修改的原子操作Json结构体

	Atomic *JsonAtomic `json:"Atomic,omitempty" name:"Atomic"`
}

func (r *UpdateAtomicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAtomicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationTemplateRequest struct {
	*tchttp.BaseRequest

	// ID 唯一ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// UUID 运维模版UUID标识

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// Name 运维模版名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Version 运维模板版本，和运维模板名称一起使用

	Version *string `json:"Version,omitempty" name:"Version"`
	// Type 运维模版类别: 支持product, application

	Type *string `json:"Type,omitempty" name:"Type"`
	// ProductVersionUUID 产品版本UUID

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	// ApplicationStandard 应用标准

	ApplicationStandard *string `json:"ApplicationStandard,omitempty" name:"ApplicationStandard"`
	// CloudID、ProductInstanceID、ProductVersionID 三者同时存在，且Type为product时，按标签精确匹配运维模板
	// CloudID 云ID

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
}

func (r *DescribeOperationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductTemplateRequest struct {
	*tchttp.BaseRequest

	//

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	//

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	//

	ProductVersionID *string `json:"ProductVersionID,omitempty" name:"ProductVersionID"`
	//

	Type *string `json:"Type,omitempty" name:"Type"`
	//

	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *DescribeProductTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationInstanceDetail struct {
	// CollectID 采集到的ApplicationID，因为存在多命名空间部署的情况，所以ID需要注入命名空间信息

	CollectID *string `json:"CollectID,omitempty" name:"CollectID"`
	// Status 运行状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// Version 应用版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// Namespace 应用命名空间

	NameSpace *string `json:"NameSpace,omitempty" name:"NameSpace"`
	// CreateAt 应用首次部署成功时间

	CreateAt *string `json:"CreateAt,omitempty" name:"CreateAt"`
	// UpdateAt 应用最近一次变更时间

	UpdateAt *string `json:"UpdateAt,omitempty" name:"UpdateAt"`
	// Message 应用的Condition信息，能看到应用异常信息详情

	Message *string `json:"Message,omitempty" name:"Message"`
}

type DeploySingleAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeploySingleAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeploySingleAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAtomicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// AtomicPackageDownloadUrl 导出原子操作生成的物料包的下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAtomicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAtomicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryImportMaterialRequest struct {
	*tchttp.BaseRequest

	// 主任务id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 子任务id

	SubTaskId *int64 `json:"SubTaskId,omitempty" name:"SubTaskId"`
	// 物料类型

	Model *string `json:"Model,omitempty" name:"Model"`
}

func (r *RetryImportMaterialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryImportMaterialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetApplicationArtifactChartEffectRequest struct {
	*tchttp.BaseRequest

	// 应用制品uuid

	ApplicationArtifactUUID *string `json:"ApplicationArtifactUUID,omitempty" name:"ApplicationArtifactUUID"`
	// chart包uuid

	ChartUUID *string `json:"ChartUUID,omitempty" name:"ChartUUID"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *SetApplicationArtifactChartEffectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetApplicationArtifactChartEffectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductMaterialInfo struct {
	//

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	//

	ProductVersionName *string `json:"ProductVersionName,omitempty" name:"ProductVersionName"`
	//

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	//

	MaterialName *string `json:"MaterialName,omitempty" name:"MaterialName"`
	//

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
	//

	Type *string `json:"Type,omitempty" name:"Type"`
	//

	ImportTime *string `json:"ImportTime,omitempty" name:"ImportTime"`
	// 是否导入成功

	Imported *bool `json:"Imported,omitempty" name:"Imported"`
	// 导入任务id

	ImportTaskId *int64 `json:"ImportTaskId,omitempty" name:"ImportTaskId"`
}

type ExportServerResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ServerExcelPkgUrl 包管理下载链接

		ServerExcelPkgUrl *string `json:"ServerExcelPkgUrl,omitempty" name:"ServerExcelPkgUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportServerResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportServerResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProgressImportTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskID *int64 `json:"TaskID,omitempty" name:"TaskID"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 已完成数

		FinishedCount *int64 `json:"FinishedCount,omitempty" name:"FinishedCount"`
		// 物料所在服务器

		MaterialServer *string `json:"MaterialServer,omitempty" name:"MaterialServer"`
		// 物料目录

		MaterialPath *string `json:"MaterialPath,omitempty" name:"MaterialPath"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProgressImportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProgressImportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResumeNodeRequest struct {
	*tchttp.BaseRequest

	//

	DagJobUUID *string `json:"DagJobUUID,omitempty" name:"DagJobUUID"`
	//

	DagNodeUUID *string `json:"DagNodeUUID,omitempty" name:"DagNodeUUID"`
}

func (r *ResumeNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResumeNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SiteID 局点 ID

		SiteID *int64 `json:"SiteID,omitempty" name:"SiteID"`
		// SiteUUID 局点 UUID

		SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
		// SiteName 局点名称

		SiteName *string `json:"SiteName,omitempty" name:"SiteName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProjectSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportCommonOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 变更单包下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 变更单包MD5信息

		MD5 *string `json:"MD5,omitempty" name:"MD5"`
		// 依赖的原子操作由运维模板进行管理, 不进行二次遍历，仅遍历当前编排中的内容

		AtomicKeyList []*AtomicKey `json:"AtomicKeyList,omitempty" name:"AtomicKeyList"`
		// 依赖的运维模板，一般是应用运维模板

		OperationTemplateKeyList []*OperationTemplateKey `json:"OperationTemplateKeyList,omitempty" name:"OperationTemplateKeyList"`
		// 应用名称列表

		ApplicationNameList []*string `json:"ApplicationNameList,omitempty" name:"ApplicationNameList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportCommonOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportCommonOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationInstanceInfo struct {
	// ApplicationName 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// ApplicationID 应用ID

	ApplicationID *string `json:"ApplicationID,omitempty" name:"ApplicationID"`
	// ApplicationCode 应用Code

	ApplicationCode *string `json:"ApplicationCode,omitempty" name:"ApplicationCode"`
	//

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	//

	ApplicationVersion *string `json:"ApplicationVersion,omitempty" name:"ApplicationVersion"`
	//

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
	//

	ApplicationInstanceName *string `json:"ApplicationInstanceName,omitempty" name:"ApplicationInstanceName"`
	//

	ApplicationInstanceID *string `json:"ApplicationInstanceID,omitempty" name:"ApplicationInstanceID"`
	//

	CurrentApplicationVersion *string `json:"CurrentApplicationVersion,omitempty" name:"CurrentApplicationVersion"`
	//

	TargetApplicationVersion *string `json:"TargetApplicationVersion,omitempty" name:"TargetApplicationVersion"`
	//

	Selected *bool `json:"Selected,omitempty" name:"Selected"`
	//

	PlanTag *int64 `json:"PlanTag,omitempty" name:"PlanTag"`
	//

	LastUpdatePlanTag *int64 `json:"LastUpdatePlanTag,omitempty" name:"LastUpdatePlanTag"`
	//

	DeployVersionID *int64 `json:"DeployVersionID,omitempty" name:"DeployVersionID"`
}

type ListChildOperationSheetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Total *int64 `json:"Total,omitempty" name:"Total"`
		//

		Sheets []*CommonOperationSheet `json:"Sheets,omitempty" name:"Sheets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListChildOperationSheetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListChildOperationSheetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationCcDeclareRequest struct {
	*tchttp.BaseRequest

	// ApplicationInfoList 应用信息列表

	ApplicationInfoList []*ApplicationInfo `json:"ApplicationInfoList,omitempty" name:"ApplicationInfoList"`
}

func (r *GetApplicationCcDeclareRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationCcDeclareRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationArtifactChartFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// chart包内文件信息

		Files []*ChartArtifact `json:"Files,omitempty" name:"Files"`
		// chart包UUID

		ChartUUID *string `json:"ChartUUID,omitempty" name:"ChartUUID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationArtifactChartFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationArtifactChartFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialCleanupSetRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// AutoCleanup 自动清理

	AutoCleanup *bool `json:"AutoCleanup,omitempty" name:"AutoCleanup"`
	// StartTime 自动清理开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// DefaultReservedInstall 默认保留安装包数量

	DefaultReservedInstall *int64 `json:"DefaultReservedInstall,omitempty" name:"DefaultReservedInstall"`
	// DefaultReservedFix 默认保留补丁包数量

	DefaultReservedFix *int64 `json:"DefaultReservedFix,omitempty" name:"DefaultReservedFix"`
}

func (r *GetMaterialCleanupSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialCleanupSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudTopologyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Clouds 局点/云 信息列表

		Cloud []*CloudItem `json:"Cloud,omitempty" name:"Cloud"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudTopologyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudTopologyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MaterialInfo struct {
	// 物料架构

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 删除时间

	DeletedAt *int64 `json:"DeletedAt,omitempty" name:"DeletedAt"`
	// 原始物料下载地址

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
	// 物料中的 files

	Files *string `json:"Files,omitempty" name:"Files"`
	// ( primary key ) 自增主键

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 是否为标准物料（true-表示标准无聊，false-局点物料）

	IsStandard *bool `json:"IsStandard,omitempty" name:"IsStandard"`
	// 物料标签（多个用，分隔）

	Label *string `json:"Label,omitempty" name:"Label"`
	// 物料元数据, 可用于后续产品化扩展

	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`
	// 物料名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 物料发布说明

	ReleaseNotes *string `json:"ReleaseNotes,omitempty" name:"ReleaseNotes"`
	// 物料来源（Import, OTA）

	Source *string `json:"Source,omitempty" name:"Source"`
	// 物料状态（deployed 已部署，not-deployed 未部署，fail-deployed 部署失败，history 历史版本）

	Status *string `json:"Status,omitempty" name:"Status"`
	// 物料分类（install,fix,upgrade）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 物料uuid，全局唯一标识

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 物料版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 流水线实例ID

	WorkflowInstanceID *int64 `json:"WorkflowInstanceID,omitempty" name:"WorkflowInstanceID"`
}

type ProductInstanceTreeInfo struct {
	// ProductCode 产品英文名

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// ProductName 产品中文名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// ProductVersionName 产品版本名称

	ProductVersionName *string `json:"ProductVersionName,omitempty" name:"ProductVersionName"`
	// ProductInstanceUUID 产品实例 UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// ProductInstanceID 产品实例 ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	//

	ProductInstanceName *string `json:"ProductInstanceName,omitempty" name:"ProductInstanceName"`
	//

	SubsystemInstanceSet []*SubsystemInstanceTreeInfo `json:"SubsystemInstanceSet,omitempty" name:"SubsystemInstanceSet"`
}

type UpdateOperationSheetRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	//

	Template *string `json:"Template,omitempty" name:"Template"`
	//

	Status *string `json:"Status,omitempty" name:"Status"`
	//

	ApplicationInstanceList []*OperationSheetApplicationInstance `json:"ApplicationInstanceList,omitempty" name:"ApplicationInstanceList"`
	//

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *UpdateOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		TaskList []*Task `json:"TaskList,omitempty" name:"TaskList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RerunNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RerunNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RerunNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否为官方产品

		BuiltIn *int64 `json:"BuiltIn,omitempty" name:"BuiltIn"`
		// 产品版本列表

		ProductVersionList []*ProductVersionBrief `json:"ProductVersionList,omitempty" name:"ProductVersionList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubsystemInstanceTreeInfo struct {
	//

	SubsystemInstanceID *string `json:"SubsystemInstanceID,omitempty" name:"SubsystemInstanceID"`
	//

	SubSystemInstanceName *string `json:"SubSystemInstanceName,omitempty" name:"SubSystemInstanceName"`
	//

	ApplicationInstanceSet []*ApplicationInstanceInfo `json:"ApplicationInstanceSet,omitempty" name:"ApplicationInstanceSet"`
}

type DescribeProductInstanceInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		CloudName *string `json:"CloudName,omitempty" name:"CloudName"`
		//

		ProductID *string `json:"ProductID,omitempty" name:"ProductID"`
		//

		ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
		//

		ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
		//

		ProductInstanceName *string `json:"ProductInstanceName,omitempty" name:"ProductInstanceName"`
		//

		AppInstanceCount *int64 `json:"AppInstanceCount,omitempty" name:"AppInstanceCount"`
		//

		AppCount *int64 `json:"AppCount,omitempty" name:"AppCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductInstanceInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductInstanceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductRollbackConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		ProductRollbackConfig *string `json:"ProductRollbackConfig,omitempty" name:"ProductRollbackConfig"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductRollbackConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductRollbackConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Sheet *OperationSheet `json:"Sheet,omitempty" name:"Sheet"`
		//

		Status *string `json:"Status,omitempty" name:"Status"`
		//

		ProgressCount *ProgressCount `json:"ProgressCount,omitempty" name:"ProgressCount"`
		//

		DAG *DAG `json:"DAG,omitempty" name:"DAG"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationTemplateIDResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// DefaultID 模版 ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// TemplateList 可能适配的运维模版信息

		TemplateList []*Template `json:"TemplateList,omitempty" name:"TemplateList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationTemplateIDResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationTemplateIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPublicMaterialInfosRequest struct {
	*tchttp.BaseRequest

	// SolutionVersionUUID 解决方案版本 UUID

	SolutionVersionUUID *string `json:"SolutionVersionUUID,omitempty" name:"SolutionVersionUUID"`
	// ProductVersionUUID 产品版本 UUID

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	// Type 物料包的类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// Arch 物料的架构

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	// Offset 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit 分页限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListPublicMaterialInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPublicMaterialInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaterialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemConfigRequest struct {
	*tchttp.BaseRequest

	// Keys 配置项

	Keys []*string `json:"Keys,omitempty" name:"Keys"`
}

func (r *DescribeSystemConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPublicMaterialInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Total 物料总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// MaterialList 物料列表

		MaterialList []*MaterialInfo `json:"MaterialList,omitempty" name:"MaterialList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPublicMaterialInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPublicMaterialInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationDetail struct {
	// ID 应用实例ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// UUID 应用实例UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// AppID 应用ID

	AppID *string `json:"AppID,omitempty" name:"AppID"`
	// ApplicationName

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// Name 应用名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Type 类型 application/tool/atomic

	Type *string `json:"Type,omitempty" name:"Type"`
	// SheetID 最近一次和应用相关的运维单ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// JobID 最近一次编排引擎的任务ID

	JobID *string `json:"JobID,omitempty" name:"JobID"`
	// OperationStatus 运维状态

	OperationStatus *string `json:"OperationStatus,omitempty" name:"OperationStatus"`
	// OperationVersion 运维记录中的版本

	OperationVersion *string `json:"OperationVersion,omitempty" name:"OperationVersion"`
	// LatestPackageVersion 当前应用物料的最新版本

	LatestPackageVersion *string `json:"LatestPackageVersion,omitempty" name:"LatestPackageVersion"`
	// SupportOperation 当前实例支持的操作

	SupportOperation []*string `json:"SupportOperation,omitempty" name:"SupportOperation"`
	// StartAt App/Tool/Atomic最近一次部署启动时间

	LatestStartAt *string `json:"LatestStartAt,omitempty" name:"LatestStartAt"`
	// Instances 应用实例列表，一个应用可能在多个命名空间下运行

	Instances []*ApplicationInstanceDetail `json:"Instances,omitempty" name:"Instances"`
	// Weights 展示权重，权重大的优先展示

	Weights *int64 `json:"Weights,omitempty" name:"Weights"`
}

type ConvertPipelineFormatRequest struct {
	*tchttp.BaseRequest

	// SourceFormat 源编排格式，yaml, tree, dag

	SourceFormat *string `json:"SourceFormat,omitempty" name:"SourceFormat"`
	// Pipeline 编排内容

	Pipeline *PipelineContent `json:"Pipeline,omitempty" name:"Pipeline"`
}

func (r *ConvertPipelineFormatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConvertPipelineFormatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUploadMaterialTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskID *int64 `json:"TaskID,omitempty" name:"TaskID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUploadMaterialTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUploadMaterialTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ClusterList 集群列表

		ClusterList []*ClusterInfo `json:"ClusterList,omitempty" name:"ClusterList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCloudClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationInfoData struct {
	//

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	//

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	//

	ApplicationInstanceID *string `json:"ApplicationInstanceID,omitempty" name:"ApplicationInstanceID"`
	//

	ApplicationInstanceName *string `json:"ApplicationInstanceName,omitempty" name:"ApplicationInstanceName"`
	//

	ApplicationValues *string `json:"ApplicationValues,omitempty" name:"ApplicationValues"`
	//

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	//

	CloudUUID *string `json:"CloudUUID,omitempty" name:"CloudUUID"`
	//

	RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
	//

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	//

	ApplicationPackageManagerURL *string `json:"ApplicationPackageManagerURL,omitempty" name:"ApplicationPackageManagerURL"`
	//

	ApplicationVersion *string `json:"ApplicationVersion,omitempty" name:"ApplicationVersion"`
	//

	ApplicationMetadata *ApplicationMetadata `json:"ApplicationMetadata,omitempty" name:"ApplicationMetadata"`
}

type GetRuntimePlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Content 运行态规划信息，yaml格式

		Content *string `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRuntimePlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRuntimePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppRuntimeNamesRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点 UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// ProductInstanceUUID 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// Limit 数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Offset 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListAppRuntimeNamesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppRuntimeNamesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MaterialCleanupSetRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// AutoCleanup 自动清理

	AutoCleanup *bool `json:"AutoCleanup,omitempty" name:"AutoCleanup"`
	// StartTime 自动清理开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// DefaultReservedInstall 默认保留安装包数量

	DefaultReservedInstall *int64 `json:"DefaultReservedInstall,omitempty" name:"DefaultReservedInstall"`
	// DefaultReservedFix 默认保留补丁包数量

	DefaultReservedFix *int64 `json:"DefaultReservedFix,omitempty" name:"DefaultReservedFix"`
}

func (r *MaterialCleanupSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MaterialCleanupSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSiteInfoDetailRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点 UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
}

func (r *DescribeProjectSiteInfoDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectSiteInfoDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationInstancePlan struct {
	// ApplicationName 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// ApplicationID 应用ID

	ApplicationID *string `json:"ApplicationID,omitempty" name:"ApplicationID"`
	// ApplicationCode 应用Code

	ApplicationCode *string `json:"ApplicationCode,omitempty" name:"ApplicationCode"`
	// ApplicationInstanceID 应用实例ID

	ApplicationInstanceID *string `json:"ApplicationInstanceID,omitempty" name:"ApplicationInstanceID"`
	// ApplicationInstanceUUID 应用实例UUID

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
	// ApplicationInstanceName 应用实例名称

	ApplicationInstanceName *string `json:"ApplicationInstanceName,omitempty" name:"ApplicationInstanceName"`
	// ApplicationBranchName 应用分支名称

	ApplicationBranchName *string `json:"ApplicationBranchName,omitempty" name:"ApplicationBranchName"`
	// ApplicationBranchID 应用分支ID

	ApplicationBranchID *string `json:"ApplicationBranchID,omitempty" name:"ApplicationBranchID"`
	// Location 示例部署位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// Config 应用配置:values.yaml

	Config *string `json:"Config,omitempty" name:"Config"`
	// AppComponentSet 应用模块信息列表

	AppComponentSet []*Component `json:"AppComponentSet,omitempty" name:"AppComponentSet"`
	// AppServerSet 应用服务器信息列表

	AppServerSet []*Server `json:"AppServerSet,omitempty" name:"AppServerSet"`
}

type PhysicalServerInfo struct {
	// UUID 服务器UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// ServerID 服务器ID

	ServerID *string `json:"ServerID,omitempty" name:"ServerID"`
	// IP 服务器IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// HostUUID 宿主机UUID

	HostUUID *string `json:"HostUUID,omitempty" name:"HostUUID"`
	// HostIP 服务器HostIP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// Type 服务器类型：物理机/虚拟机

	Type *string `json:"Type,omitempty" name:"Type"`
	// Location 位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// DiskInfo 磁盘信息

	DiskInfo *string `json:"DiskInfo,omitempty" name:"DiskInfo"`
	// OsInfo os信息

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// Metadata 服务器属性

	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`
	// ServerType 服务器型号

	ServerType *string `json:"ServerType,omitempty" name:"ServerType"`
	// ExtraAttr 服务器其他属性

	ExtraAttr *string `json:"ExtraAttr,omitempty" name:"ExtraAttr"`
	// CVMUUID 虚拟机在CVM的UUID

	CVMUUID *string `json:"CVMUUID,omitempty" name:"CVMUUID"`
	// OsName 操作系统名称

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// LogicArea 逻辑区域

	LogicArea *string `json:"LogicArea,omitempty" name:"LogicArea"`
	// // RackID 机架 从location解析
	// RackID string
	// // RackPosID 机位号 从location解析
	// RackPosID int
	// Bonding 网卡bonding 从metadata获取

	Bonding *string `json:"Bonding,omitempty" name:"Bonding"`
	// NUMACount NUMA 个数 从metadata获取

	NUMACount *string `json:"NUMACount,omitempty" name:"NUMACount"`
	// HyperThreading 超线程  从metadata获取

	HyperThreading *string `json:"HyperThreading,omitempty" name:"HyperThreading"`
	// CPUArch CPU架构  从metadata获取

	CPUArch *string `json:"CPUArch,omitempty" name:"CPUArch"`
	// DescribeInfo 服务器配置  从metadata获取

	DescribeInfo *string `json:"DescribeInfo,omitempty" name:"DescribeInfo"`
	// SN 编码：从ExtraAttr 获取

	SN *string `json:"SN,omitempty" name:"SN"`
	// CPUNum cpu数量 	从metadata获取

	CPUNum *int64 `json:"CPUNum,omitempty" name:"CPUNum"`
	// AppComponentSet 服务器上的应用模块信息

	AppComponentSet []*AppComponentInfo `json:"AppComponentSet,omitempty" name:"AppComponentSet"`
	// 服务器上部署的产品

	ProductSet []*string `json:"ProductSet,omitempty" name:"ProductSet"`
}

type ConfirmCommonOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConfirmCommonOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfirmCommonOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTaskRequest struct {
	*tchttp.BaseRequest
}

func (r *ListTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MaterialCleanupSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// AutoCleanup 自动清理

		AutoCleanup *bool `json:"AutoCleanup,omitempty" name:"AutoCleanup"`
		// StartTime 自动清理开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// DefaultReservedInstall 默认保留安装包数量

		DefaultReservedInstall *int64 `json:"DefaultReservedInstall,omitempty" name:"DefaultReservedInstall"`
		// DefaultReservedFix 默认保留补丁包数量

		DefaultReservedFix *int64 `json:"DefaultReservedFix,omitempty" name:"DefaultReservedFix"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MaterialCleanupSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MaterialCleanupSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialAtomicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Total *int64 `json:"Total,omitempty" name:"Total"`
		//

		AtomicList []*JsonAtomic `json:"AtomicList,omitempty" name:"AtomicList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialAtomicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialAtomicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductInstanceAppUUID struct {
	// ProductInstanceUUID 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// ApplicationUUIDList 运维单关联应用UUID列表

	ApplicationIDList []*string `json:"ApplicationIDList,omitempty" name:"ApplicationIDList"`
}

type ListProductInstancePlansRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// CloudID 云UUID

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// ProductID 产品ID

	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`
	// ProductCode 产品名称

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// ProductInstanceUUID 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// DetailLevel 详情级别: 可选值:product, subsystem, application, default:product

	DetailLevel *string `json:"DetailLevel,omitempty" name:"DetailLevel"`
	// Location 产品位置信息

	Location *LocationInfo `json:"Location,omitempty" name:"Location"`
	// ApplicationUUIDList 应用UUID列表

	ApplicationUUIDList []*string `json:"ApplicationUUIDList,omitempty" name:"ApplicationUUIDList"`
	// ApplicationName 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// IgnoreDeployType

	IgnoreDeployType *bool `json:"IgnoreDeployType,omitempty" name:"IgnoreDeployType"`
	// PlanID 规划ID

	PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
	// PlanState 规划状态

	PlanState *string `json:"PlanState,omitempty" name:"PlanState"`
	// Tag 规划快照

	Tag *int64 `json:"Tag,omitempty" name:"Tag"`
}

func (r *ListProductInstancePlansRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductInstancePlansRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSiteApplicationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Products []*SiteProduct `json:"Products,omitempty" name:"Products"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSiteApplicationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSiteApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnvInfoData struct {
	//

	TCSMasterIP *string `json:"TCSMasterIP,omitempty" name:"TCSMasterIP"`
}

type DescribeApprovalHistoryRequest struct {
	*tchttp.BaseRequest

	// SheetID 运维单 ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
}

func (r *DescribeApprovalHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApprovalHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportOtaCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CertificateUUID 证书UUID，从证书Content提取

		CertificateUUID *string `json:"CertificateUUID,omitempty" name:"CertificateUUID"`
		// IssueDate 签发时间

		IssueDate *string `json:"IssueDate,omitempty" name:"IssueDate"`
		// Content 证书内容

		Content *string `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportOtaCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportOtaCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAtomicRenderDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAtomicRenderDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAtomicRenderDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOtaDownloadConfigRequest struct {
	*tchttp.BaseRequest

	// 自动下载配置 1：开启 0：关闭

	DownloadConfig *string `json:"DownloadConfig,omitempty" name:"DownloadConfig"`
	// 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
}

func (r *UpdateOtaDownloadConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOtaDownloadConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PipelineContent struct {
	//

	DAG *DAG `json:"DAG,omitempty" name:"DAG"`
	//

	JSON *string `json:"JSON,omitempty" name:"JSON"`
	//

	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
}

type ListSiteOperationSheetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Total *int64 `json:"Total,omitempty" name:"Total"`
		//

		Sheets []*CommonOperationSheet `json:"Sheets,omitempty" name:"Sheets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSiteOperationSheetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSiteOperationSheetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductDetail struct {
	//

	ProductUUID *string `json:"ProductUUID,omitempty" name:"ProductUUID"`
	// sheet from jiguang: ProductCode = "", ProductName is actually ProductCode
	// sheet created from local: ProductCode = ProductCode(like 'cbs'), ProductName is ProductName(like '云硬盘（CBS）')

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	//

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	//

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	//

	ProductVersionName *string `json:"ProductVersionName,omitempty" name:"ProductVersionName"`
	//

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	//

	ProductInstanceName *string `json:"ProductInstanceName,omitempty" name:"ProductInstanceName"`
	//

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	//

	ProductVersionArtifactTag *string `json:"ProductVersionArtifactTag,omitempty" name:"ProductVersionArtifactTag"`
	//

	Applications []*ApplicationInfo `json:"Applications,omitempty" name:"Applications"`
}

type DescribeAppRuntimeYAMLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// AppYAML 应用YAML

		AppYAML *string `json:"AppYAML,omitempty" name:"AppYAML"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppRuntimeYAMLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppRuntimeYAMLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyOperationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ID 被复制运维模版的主键 ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyOperationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyOperationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSolutionTemplatesRequest struct {
	*tchttp.BaseRequest

	//

	IsActive *bool `json:"IsActive,omitempty" name:"IsActive"`
	//

	SolutionCode *string `json:"SolutionCode,omitempty" name:"SolutionCode"`
	//

	SolutionName *string `json:"SolutionName,omitempty" name:"SolutionName"`
	//

	SolutionVersion *string `json:"SolutionVersion,omitempty" name:"SolutionVersion"`
	// 模版名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 模版版本

	Version *string `json:"Version,omitempty" name:"Version"`
	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	//

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Filter 类似数据引擎的构造filter, 用于原子操作筛选

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
}

func (r *ListSolutionTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSolutionTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductInstanceOperationSheetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalCount 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Sheets 变更单列表

		Sheets []*ProductInstanceOperationSheet `json:"Sheets,omitempty" name:"Sheets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductInstanceOperationSheetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductInstanceOperationSheetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplySiteMaterialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ID 云 ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// UUID 云 UUID

		UUID *string `json:"UUID,omitempty" name:"UUID"`
		// Name 云名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// PlanID 规划 ID

		PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
		// PlanName 规划名称

		PlanName *string `json:"PlanName,omitempty" name:"PlanName"`
		// PlanUUID 规划 UUID

		PlanUUID *string `json:"PlanUUID,omitempty" name:"PlanUUID"`
		// PlanVersion 规划版本

		PlanVersion *string `json:"PlanVersion,omitempty" name:"PlanVersion"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplySiteMaterialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplySiteMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOtaDownloadConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOtaDownloadConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOtaDownloadConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Strategy struct {
	//

	RetryLimit *int64 `json:"RetryLimit,omitempty" name:"RetryLimit"`
	//

	RetryInterval *int64 `json:"RetryInterval,omitempty" name:"RetryInterval"`
	//

	TimeoutInSeconds *int64 `json:"TimeoutInSeconds,omitempty" name:"TimeoutInSeconds"`
}

type DeleteMaterialRequest struct {
	*tchttp.BaseRequest

	// 导入任务id

	ImportTaskId *int64 `json:"ImportTaskId,omitempty" name:"ImportTaskId"`
	// 物料所在机器

	MachineIP *string `json:"MachineIP,omitempty" name:"MachineIP"`
	// 物料在机器上地址

	MaterialPath *string `json:"MaterialPath,omitempty" name:"MaterialPath"`
	// 删除操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DeleteMaterialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportProductDataRequest struct {
	*tchttp.BaseRequest

	// package 包

	Data *string `json:"Data,omitempty" name:"Data"`
	// 任务ID

	TaskID *int64 `json:"TaskID,omitempty" name:"TaskID"`
	// 操作者

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ImportProductDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportProductDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationPackagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		ApplicationPackages []*ApplicationPackage `json:"ApplicationPackages,omitempty" name:"ApplicationPackages"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationPackagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckSiteOperationSheetPlanExistenceRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
}

func (r *CheckSiteOperationSheetPlanExistenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSiteOperationSheetPlanExistenceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEnvRequest struct {
	*tchttp.BaseRequest

	// EnvName 环境名称

	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`
	// Description 环境描述描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// ConnectUrl 环境连接信息

	ConnectUrl *string `json:"ConnectUrl,omitempty" name:"ConnectUrl"`
	// AuroraEnvUUID 环境管理 UUID

	AuroraEnvUUID *string `json:"AuroraEnvUUID,omitempty" name:"AuroraEnvUUID"`
	// InstallConfig 环境部署配置

	InstallConfig *string `json:"InstallConfig,omitempty" name:"InstallConfig"`
	// MaterialUUID 标准物料 UUID

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
	// MaterialURL 物料下载地址

	MaterialURL *string `json:"MaterialURL,omitempty" name:"MaterialURL"`
}

func (r *CreateEnvRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEnvRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartApplicationDeployRequest struct {
	*tchttp.BaseRequest

	// ApplicationType 应用的类型

	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`
	// ApplicationName 应用的名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// ApplicationArtifactURL 应用的地址

	ApplicationArtifactURL *string `json:"ApplicationArtifactURL,omitempty" name:"ApplicationArtifactURL"`
	// ApplicationValue 应用的 values

	ApplicationValue *string `json:"ApplicationValue,omitempty" name:"ApplicationValue"`
	// MaterIP 部署机的 IP

	MaterIP *string `json:"MaterIP,omitempty" name:"MaterIP"`
	// Namespace 应用的 namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// ApplicationMetadata 应用的元数据信息

	ApplicationMetadata *ApplicationMetadata `json:"ApplicationMetadata,omitempty" name:"ApplicationMetadata"`
	// ApplicationArch 应用的架构

	ApplicationArch *string `json:"ApplicationArch,omitempty" name:"ApplicationArch"`
	// PackageType 制品包类型

	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`
}

func (r *StartApplicationDeployRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartApplicationDeployRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialProductApplicationInstanceRequest struct {
	*tchttp.BaseRequest

	// SubsystemInstanceID 产品子系统实例ID

	SubsystemInstanceID *int64 `json:"SubsystemInstanceID,omitempty" name:"SubsystemInstanceID"`
	// Name 产品子系统名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Limit 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Offset 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListMaterialProductApplicationInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialProductApplicationInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationTemplateRequest struct {
	*tchttp.BaseRequest

	// Kind 模板类型 取值为deploy update upgrade

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// Type 描述模版类别,取值为product application等

	Type *string `json:"Type,omitempty" name:"Type"`
	// IsActive 描述运维模板是否是生效版本，0的时候不生效，1的时候生效

	IsActive *bool `json:"IsActive,omitempty" name:"IsActive"`
	// Labels 运维模版标签

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// Offset 分页使用

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit 分页使用

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Filter 过滤条件

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
}

func (r *ListOperationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductApplicationsRequest struct {
	*tchttp.BaseRequest

	// 产品版本 UUID

	ProductVersionID *string `json:"ProductVersionID,omitempty" name:"ProductVersionID"`
}

func (r *ListProductApplicationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductApplicationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// PlanID 规划ID

		PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
		// ProductInstanceID 产品实例ID

		ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
		// ProductInstanceUUID 产品实例UUID

		ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
		// ProductInstanceName 产品实例名称

		ProductInstanceName *string `json:"ProductInstanceName,omitempty" name:"ProductInstanceName"`
		// ProductName 产品名称

		ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
		// ProductID 产品ID

		ProductID *string `json:"ProductID,omitempty" name:"ProductID"`
		// ProductCode 产品Code

		ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
		// ProductVersionID 产品版本ID

		ProductVersionID *string `json:"ProductVersionID,omitempty" name:"ProductVersionID"`
		// ProductVersionName 产品版本名称

		ProductVersionName *string `json:"ProductVersionName,omitempty" name:"ProductVersionName"`
		// ProductVersionDataTag 产品版本的数据Tag

		ProductVersionDataTag *string `json:"ProductVersionDataTag,omitempty" name:"ProductVersionDataTag"`
		// ProductCategory 产品分类

		ProductCategory *string `json:"ProductCategory,omitempty" name:"ProductCategory"`
		// Location 产品部署位置

		Location *string `json:"Location,omitempty" name:"Location"`
		// AppInstanceCount 当前产品实例的应用实例数量

		AppInstanceCount *int64 `json:"AppInstanceCount,omitempty" name:"AppInstanceCount"`
		// AppInstanceList 应用实例列表

		AppInstanceList []*AppBriefInfo `json:"AppInstanceList,omitempty" name:"AppInstanceList"`
		// Config 产品配置信息

		Config *string `json:"Config,omitempty" name:"Config"`
		// ConfigSchema 产品配置用于前端展示的schema

		ConfigSchema *string `json:"ConfigSchema,omitempty" name:"ConfigSchema"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MaterialData struct {
	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// 解决方案Code

	SolutionCode *string `json:"SolutionCode,omitempty" name:"SolutionCode"`
	// 解决方案名称

	SolutionName *string `json:"SolutionName,omitempty" name:"SolutionName"`
	// 解决方案版本列表

	SolutionVersions []*SolutionVersion `json:"SolutionVersions,omitempty" name:"SolutionVersions"`
	// 通用物料包

	Packages []*PackageData `json:"Packages,omitempty" name:"Packages"`
	// 缺失的应用包信息

	MissingApplicationPackage []*MissingApplicationPackage `json:"MissingApplicationPackage,omitempty" name:"MissingApplicationPackage"`
}

type ApplicationInstance struct {
	//

	ApplicationInstanceID *int64 `json:"ApplicationInstanceID,omitempty" name:"ApplicationInstanceID"`
	//

	ApplicationInstanceName *string `json:"ApplicationInstanceName,omitempty" name:"ApplicationInstanceName"`
	//

	Config *string `json:"Config,omitempty" name:"Config"`
}

type ListMaterialProductSubsystemInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		ProductInstanceList []*ProductInstance `json:"ProductInstanceList,omitempty" name:"ProductInstanceList"`
		// 解决方案名称

		SolutionName *string `json:"SolutionName,omitempty" name:"SolutionName"`
		// 解决方案版本名称

		SolutionVersionName *string `json:"SolutionVersionName,omitempty" name:"SolutionVersionName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialProductSubsystemInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialProductSubsystemInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportProductDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导入状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 响应详情

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportProductDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportProductDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DAGNode struct {
	//

	NodeUUID *string `json:"NodeUUID,omitempty" name:"NodeUUID"`
	//

	Status *string `json:"Status,omitempty" name:"Status"`
	//

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	//

	Name *string `json:"Name,omitempty" name:"Name"`
	//

	Description *string `json:"Description,omitempty" name:"Description"`
	//

	DescriptionExtended *string `json:"DescriptionExtended,omitempty" name:"DescriptionExtended"`
	//

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	//

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	//

	RunType *string `json:"RunType,omitempty" name:"RunType"`
	//

	Parallel *bool `json:"Parallel,omitempty" name:"Parallel"`
	//

	RetryLimit *int64 `json:"RetryLimit,omitempty" name:"RetryLimit"`
	//

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	//

	Condition *string `json:"Condition,omitempty" name:"Condition"`
	//

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	//

	Annotations *string `json:"Annotations,omitempty" name:"Annotations"`
	//

	Inputs *string `json:"Inputs,omitempty" name:"Inputs"`
	//

	Outputs *string `json:"Outputs,omitempty" name:"Outputs"`
	//

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
	//

	PipelineName *string `json:"PipelineName,omitempty" name:"PipelineName"`
	//

	PipelineDisplayName *string `json:"PipelineDisplayName,omitempty" name:"PipelineDisplayName"`
	//

	SubDagNodes []*string `json:"SubDagNodes,omitempty" name:"SubDagNodes"`
	//

	IsStepNode *bool `json:"IsStepNode,omitempty" name:"IsStepNode"`
	//

	Params []*Param `json:"Params,omitempty" name:"Params"`
	//

	NodePath *string `json:"NodePath,omitempty" name:"NodePath"`
}

type OperationSheetApplicationInstance struct {
	// ApplicationName 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// ApplicationInstanceUUID 应用实例的 UUID

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
}

type ImportCommonOperationSheetRequest struct {
	*tchttp.BaseRequest

	// DownloadUrl 通过HttpGET下载变更单包

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
	// DownloadNode 通过SSH登录远端节点下载变更单包

	DownloadNode *ImportOperationSheetPackageDownloadNode `json:"DownloadNode,omitempty" name:"DownloadNode"`
	// DownloadBase64 通过直接传输变更单包的base64数据

	DownloadBase64 *string `json:"DownloadBase64,omitempty" name:"DownloadBase64"`
	// 变更单包的MD5信息，如果没有则不做校验

	MD5 *string `json:"MD5,omitempty" name:"MD5"`
	// 物料导入的制品，制品版本信息

	Applications []*ImportOperationSheetPackageApplication `json:"Applications,omitempty" name:"Applications"`
	// 问题单基本信息

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	//

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
	//

	TaskID *int64 `json:"TaskID,omitempty" name:"TaskID"`
	//

	DoInstantiation *bool `json:"DoInstantiation,omitempty" name:"DoInstantiation"`
}

func (r *ImportCommonOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportCommonOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckOperationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		WarnMessage *string `json:"WarnMessage,omitempty" name:"WarnMessage"`
		//

		MissApplicationList []*string `json:"MissApplicationList,omitempty" name:"MissApplicationList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckOperationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckOperationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrderType struct {
	//

	Key *string `json:"Key,omitempty" name:"Key"`
	//

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

type DiffPlanPackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Diffs 差异资源列表

		Diffs []*PlanDiff `json:"Diffs,omitempty" name:"Diffs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DiffPlanPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DiffPlanPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DAG struct {
	//

	DAGNode *DAGNode `json:"DAGNode,omitempty" name:"DAGNode"`
	//

	DAGNodeList []*DAGNode `json:"DAGNodeList,omitempty" name:"DAGNodeList"`
}

type DescribeAppRollbackConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		ApplicationRollbackConfig *string `json:"ApplicationRollbackConfig,omitempty" name:"ApplicationRollbackConfig"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppRollbackConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppRollbackConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOtaCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CertificateUUID 证书UUID，从证书Content提取

		CertificateUUID *string `json:"CertificateUUID,omitempty" name:"CertificateUUID"`
		// IssueDate 签发时间

		IssueDate *string `json:"IssueDate,omitempty" name:"IssueDate"`
		// Content 证书内容

		Content *string `json:"Content,omitempty" name:"Content"`
		// CheckConfig 更新检查配置 0-手动检查 1-自动检查

		CheckConfig *string `json:"CheckConfig,omitempty" name:"CheckConfig"`
		// DownloadConfig 物料下载配置 0-手动下载 1-自动下载

		DownloadConfig *string `json:"DownloadConfig,omitempty" name:"DownloadConfig"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOtaCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOtaCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAtomicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// UUID 原子操作uuid，可单独使用，第一优先级查询

		UUID *string `json:"UUID,omitempty" name:"UUID"`
		// AtomicName 原子操作名称，和version一起使用，第二优先级查询

		AtomicName *string `json:"AtomicName,omitempty" name:"AtomicName"`
		// Version 原子操作版本，和name一起使用，第二优先级查询

		AtomicVersion *string `json:"AtomicVersion,omitempty" name:"AtomicVersion"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAtomicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAtomicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUploadMaterialTaskRequest struct {
	*tchttp.BaseRequest

	// 物料所在服务器

	MaterialServer *string `json:"MaterialServer,omitempty" name:"MaterialServer"`
	// 物料父目录

	MaterialPath *string `json:"MaterialPath,omitempty" name:"MaterialPath"`
	// 操作者

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 物料来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 物料 UUID, 如果存在 UUID 表示当前物料包为标准物料

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
}

func (r *CreateUploadMaterialTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUploadMaterialTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AtomicParam struct {
	// Name 参数名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Description 参数描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// Type 参数类型，比如string，int，object，list

	Type *string `json:"Type,omitempty" name:"Type"`
	// Default 参数默认值

	Default *string `json:"Default,omitempty" name:"Default"`
	// Required 参数是否必须，默认false

	Required *bool `json:"Required,omitempty" name:"Required"`
	// ValueFrom 参数来源，业务原子操作使用

	ValueFrom *string `json:"ValueFrom,omitempty" name:"ValueFrom"`
	// ValueRendered 参数渲染后的值

	ValueRendered *string `json:"ValueRendered,omitempty" name:"ValueRendered"`
}

type ListApplicationPackagesRequest struct {
	*tchttp.BaseRequest

	// 产品版本UUID

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	// 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// 应用UUID

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	// 筛选信息

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
	//

	Order []*OrderType `json:"Order,omitempty" name:"Order"`
	//

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListApplicationPackagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationPackagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCommonOperationSheetAppInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品实例列表

		ProductInstanceList []*ProductInstanceTreeInfo `json:"ProductInstanceList,omitempty" name:"ProductInstanceList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCommonOperationSheetAppInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCommonOperationSheetAppInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PauseJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PauseJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PauseJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportOperationSheetPackageDownloadNode struct {
	//

	Node *string `json:"Node,omitempty" name:"Node"`
	//

	Port *int64 `json:"Port,omitempty" name:"Port"`
	//

	User *string `json:"User,omitempty" name:"User"`
	//

	Password *string `json:"Password,omitempty" name:"Password"`
	//

	Path *string `json:"Path,omitempty" name:"Path"`
}

type CreateOperationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ID 返回创建的运维模版ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSiteOperationSheetRequest struct {
	*tchttp.BaseRequest

	//

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	//

	SheetName *string `json:"SheetName,omitempty" name:"SheetName"`
	//

	SheetType *string `json:"SheetType,omitempty" name:"SheetType"`
	//

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	//

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateSiteOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSiteOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePlanApplicationTreeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品实例列表

		ProductInstanceList []*InstanceTreeProduct `json:"ProductInstanceList,omitempty" name:"ProductInstanceList"`
		// 解决方案名称

		SolutionName *string `json:"SolutionName,omitempty" name:"SolutionName"`
		// 解决方案版本名称

		SolutionVersionName *string `json:"SolutionVersionName,omitempty" name:"SolutionVersionName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePlanApplicationTreeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePlanApplicationTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryImportMaterialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主任务id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryImportMaterialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryImportMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KubeVmCapacity struct {
	//

	Limits *VmSpec `json:"Limits,omitempty" name:"Limits"`
	//

	Requests *VmSpec `json:"Requests,omitempty" name:"Requests"`
}

type ProductInstancePlan struct {
	// ProductName 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// ProductID 产品ID

	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`
	// ProductCode 产品Code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// ProductInstanceUUID 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// ProductInstanceName 产品实例名称

	ProductInstanceName *string `json:"ProductInstanceName,omitempty" name:"ProductInstanceName"`
	// Location 产品部署位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// Config 产品配置

	Config *string `json:"Config,omitempty" name:"Config"`
	// ProductVersionID 产品版本ID

	ProductVersionID *string `json:"ProductVersionID,omitempty" name:"ProductVersionID"`
	// ProductVersionName 产品版本名称

	ProductVersionName *string `json:"ProductVersionName,omitempty" name:"ProductVersionName"`
	// ProductVersionDataTag 产品版本的数据Tag

	ProductVersionDataTag *string `json:"ProductVersionDataTag,omitempty" name:"ProductVersionDataTag"`
	// ProductCategory 产品分类

	ProductCategory *string `json:"ProductCategory,omitempty" name:"ProductCategory"`
	// SubsystemInstanceSet 子系统实例规划列表

	SubsystemInstanceSet []*SubsystemInstancePlan `json:"SubsystemInstanceSet,omitempty" name:"SubsystemInstanceSet"`
	// ProductLocation 产品位置信息

	ProductLocation *ProductLocation `json:"ProductLocation,omitempty" name:"ProductLocation"`
}

type GetProgressImportTaskRequest struct {
	*tchttp.BaseRequest
}

func (r *GetProgressImportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProgressImportTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationSheetRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	//

	Deep *int64 `json:"Deep,omitempty" name:"Deep"`
	//

	NodePath *string `json:"NodePath,omitempty" name:"NodePath"`
}

func (r *DescribeOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductOverviewRequest struct {
	*tchttp.BaseRequest

	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	//

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
}

func (r *DescribeProductOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialProductSubsystemInstanceRequest struct {
	*tchttp.BaseRequest

	// MaterialUUID 物料UUID

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
	// ProductVersionUUID 产品版本uuid

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
}

func (r *ListMaterialProductSubsystemInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialProductSubsystemInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GeneratePipelineRunRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	//

	Deep *int64 `json:"Deep,omitempty" name:"Deep"`
}

func (r *GeneratePipelineRunRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GeneratePipelineRunRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPublicMaterialSolutionVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SolutionVersions 解决方案版本列表

		SolutionVersions []*SolutionVersion `json:"SolutionVersions,omitempty" name:"SolutionVersions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPublicMaterialSolutionVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPublicMaterialSolutionVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOtaCertificateRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
}

func (r *DescribeOtaCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOtaCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrchestratingItemsProduct struct {
	//

	ProductUUID *string `json:"ProductUUID,omitempty" name:"ProductUUID"`
	// sheet from jiguang: ProductCode = "", ProductName is actually ProductCode
	// sheet created from local: ProductCode = ProductCode(like 'cbs'), ProductName is ProductName(like '云硬盘（CBS）')

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	//

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	//

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	//

	ProductVersionName *string `json:"ProductVersionName,omitempty" name:"ProductVersionName"`
	//

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	//

	ProductInstanceName *string `json:"ProductInstanceName,omitempty" name:"ProductInstanceName"`
	//

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	//

	ProductVersionArtifactTag *string `json:"ProductVersionArtifactTag,omitempty" name:"ProductVersionArtifactTag"`
	//

	Applications []*ApplicationInfo `json:"Applications,omitempty" name:"Applications"`
}

type ServerInfo struct {
	// UUID 服务器UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// IP 服务器IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// HostUUID 宿主机UUID

	HostUUID *string `json:"HostUUID,omitempty" name:"HostUUID"`
	// HostIP 宿主机IP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// Type 服务器类型：物理机/虚拟机

	Type *string `json:"Type,omitempty" name:"Type"`
	// Location 位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// DiskInfo 磁盘信息

	DiskInfo *string `json:"DiskInfo,omitempty" name:"DiskInfo"`
	// OsInfo os信息

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// OsName string

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// Metadata 服务器属性

	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`
	// ServerType 服务器型号

	ServerType *string `json:"ServerType,omitempty" name:"ServerType"`
	// ExtraAttr 服务器其他属性

	ExtraAttr *string `json:"ExtraAttr,omitempty" name:"ExtraAttr"`
	// CVMUUID 虚拟机在CVM的UUID

	CVMUUID *string `json:"CVMUUID,omitempty" name:"CVMUUID"`
	// Capacity vm capacity

	Capacity *KubeVmCapacity `json:"Capacity,omitempty" name:"Capacity"`
}

type GetPackageManagerHostRequest struct {
	*tchttp.BaseRequest
}

func (r *GetPackageManagerHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPackageManagerHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetActivePlanRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// PlanVersion 规划版本

	PlanVersion *string `json:"PlanVersion,omitempty" name:"PlanVersion"`
}

func (r *SetActivePlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetActivePlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MaterialProduct struct {
	// ProductUUID 产品 UUID

	ProductUUID *string `json:"ProductUUID,omitempty" name:"ProductUUID"`
	// ProductName 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// ProductVersionUUID 产品版本 UUID

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	// ProductVersionName 产品版本名称

	ProductVersionName *string `json:"ProductVersionName,omitempty" name:"ProductVersionName"`
	// ProductCategory 产品分类

	ProductCategory *string `json:"ProductCategory,omitempty" name:"ProductCategory"`
	// ProductCategoryName 产品分类名称

	ProductCategoryName *string `json:"ProductCategoryName,omitempty" name:"ProductCategoryName"`
}

type Server struct {
	// UUID 服务器UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// ServerID 服务器ID

	ServerID *string `json:"ServerID,omitempty" name:"ServerID"`
	// IP 服务器IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// HostUUID 宿主机UUID

	HostUUID *string `json:"HostUUID,omitempty" name:"HostUUID"`
	// HostIP 服务器HostIP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// Type 服务器类型：物理机/虚拟机

	Type *string `json:"Type,omitempty" name:"Type"`
	// Location 位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// DiskInfo 磁盘信息

	DiskInfo *string `json:"DiskInfo,omitempty" name:"DiskInfo"`
	// OsInfo os信息

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// Metadata 服务器属性

	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`
	// ServerType 服务器型号

	ServerType *string `json:"ServerType,omitempty" name:"ServerType"`
	// ExtraAttr 服务器其他属性

	ExtraAttr *string `json:"ExtraAttr,omitempty" name:"ExtraAttr"`
	// CVMUUID 虚拟机在CVM的UUID

	CVMUUID *string `json:"CVMUUID,omitempty" name:"CVMUUID"`
}

type CheckSiteOperationSheetMaterialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Exceptions []*CheckMaterialException `json:"Exceptions,omitempty" name:"Exceptions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckSiteOperationSheetMaterialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSiteOperationSheetMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerLocation struct {
	// RegionID 地域 ID

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	// ZoneID 可用区 ID

	ZoneID *string `json:"ZoneID,omitempty" name:"ZoneID"`
}

type SiteMaterial struct {
	// 物料名称

	MaterialName *string `json:"MaterialName,omitempty" name:"MaterialName"`
	// 物料导入id

	MaterialImportTaskID *int64 `json:"MaterialImportTaskID,omitempty" name:"MaterialImportTaskID"`
	// 物料uuid

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
	// 物料状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 物料导入中状态时的百分比

	Progress *string `json:"Progress,omitempty" name:"Progress"`
	// 物料类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 物料版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 物料架构

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	// 物料来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 物料标签

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
	// 物料发布日志

	ReleaseNote *string `json:"ReleaseNote,omitempty" name:"ReleaseNote"`
	// 是否允许删除

	DeleteAllow *bool `json:"DeleteAllow,omitempty" name:"DeleteAllow"`
	// 物料所在机器

	MachineIP *string `json:"MachineIP,omitempty" name:"MachineIP"`
	// 物料在机器上地址

	MaterialPath *string `json:"MaterialPath,omitempty" name:"MaterialPath"`
	// 物料导入时间

	ImportTime *string `json:"ImportTime,omitempty" name:"ImportTime"`
}

type CreateSolutionTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ID 返回创建的运维模版ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSolutionTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSolutionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductInstanceOperationSheetsRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// RegionUUID 区域UUID

	RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
	// Limit 限制返回数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Offset 偏移量（用于分页）

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Orders 排序条件

	Orders []*OrderType `json:"Orders,omitempty" name:"Orders"`
	// Filters 过滤条件

	Filters []*FilterType `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListProductInstanceOperationSheetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductInstanceOperationSheetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductInstancePlansResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CloudID 云ID

		CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
		// PlanID 生效的规划ID

		PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
		// ProductInstanceSet 产品实例规划列表

		ProductInstanceSet []*ProductInstancePlan `json:"ProductInstanceSet,omitempty" name:"ProductInstanceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductInstancePlansResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductInstancePlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetInstantiateAnnotationForOperationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Template *string `json:"Template,omitempty" name:"Template"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetInstantiateAnnotationForOperationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetInstantiateAnnotationForOperationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSiteMaterialInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Products 产品列表

		ProductID []*MaterialProduct `json:"ProductID,omitempty" name:"ProductID"`
		// SolutionVersionName 解决方案版本名称

		SolutionVersionName *string `json:"SolutionVersionName,omitempty" name:"SolutionVersionName"`
		// SolutionVersionUUID 解决方案版本 UUID

		SolutionVersionUUID *string `json:"SolutionVersionUUID,omitempty" name:"SolutionVersionUUID"`
		// MaterialName 物料名称

		MaterialName *string `json:"MaterialName,omitempty" name:"MaterialName"`
		// Arch 物料架构

		Arch *string `json:"Arch,omitempty" name:"Arch"`
		// Version 物料版本

		Version *string `json:"Version,omitempty" name:"Version"`
		// Label 物料标签

		Label *string `json:"Label,omitempty" name:"Label"`
		// Source 物料来源

		Source *string `json:"Source,omitempty" name:"Source"`
		// MaterialUUID 物料 UUID

		MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
		// MaterialID 物料 ID

		MaterialID *int64 `json:"MaterialID,omitempty" name:"MaterialID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSiteMaterialInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSiteMaterialInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Template *string `json:"Template,omitempty" name:"Template"`
		//

		DAG *DAG `json:"DAG,omitempty" name:"DAG"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubsystemInstancePlan struct {
	// SubsystemInstanceUUID 子系统实例UUID

	SubsystemInstanceUUID *string `json:"SubsystemInstanceUUID,omitempty" name:"SubsystemInstanceUUID"`
	// SubsystemInstanceID 子系统实例ID

	SubsystemInstanceID *string `json:"SubsystemInstanceID,omitempty" name:"SubsystemInstanceID"`
	// SubsystemInstanceName 子系统实例名称

	SubsystemInstanceName *string `json:"SubsystemInstanceName,omitempty" name:"SubsystemInstanceName"`
	// Location 子系统部署位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// SubsystemID 子系统ID

	SubsystemID *string `json:"SubsystemID,omitempty" name:"SubsystemID"`
	// SubsystemName 子系统名称

	SubsystemName *string `json:"SubsystemName,omitempty" name:"SubsystemName"`
	// SubsystemCode 子系统Code

	SubsystemCode *string `json:"SubsystemCode,omitempty" name:"SubsystemCode"`
	// ApplicationInstanceSet 应用实例规划列表

	ApplicationInstanceSet []*ApplicationInstancePlan `json:"ApplicationInstanceSet,omitempty" name:"ApplicationInstanceSet"`
}

type DeleteCommonOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCommonOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCommonOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartDevelopmentPhaseApplicationDeployResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SheetID 运维单 ID

		SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
		// JobID 部署任务 ID

		JobID *string `json:"JobID,omitempty" name:"JobID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartDevelopmentPhaseApplicationDeployResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartDevelopmentPhaseApplicationDeployResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInitDeploySheetOperationTemplateRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
}

func (r *GetInitDeploySheetOperationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInitDeploySheetOperationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PlanSubsystemInstance struct {
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 删除时间

	DeletedAt *int64 `json:"DeletedAt,omitempty" name:"DeletedAt"`
	// ( primary key ) 子系统自增ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 子系统级别

	Level *string `json:"Level,omitempty" name:"Level"`
	// 子系统实例部署位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 子系统实例下的应用实例

	PlanApplicationInstance []*PlanApplicationInstance `json:"PlanApplicationInstance,omitempty" name:"PlanApplicationInstance"`
	// 关联的规划ID

	PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
	// 关联的产品实例的UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// 子系统英文代号

	SubsystemCode *string `json:"SubsystemCode,omitempty" name:"SubsystemCode"`
	// 子系统UUID

	SubsystemID *string `json:"SubsystemID,omitempty" name:"SubsystemID"`
	// 子系统实例ID

	SubsystemInstanceID *string `json:"SubsystemInstanceID,omitempty" name:"SubsystemInstanceID"`
	// 子系统实例名称

	SubsystemInstanceName *string `json:"SubsystemInstanceName,omitempty" name:"SubsystemInstanceName"`
	// 子系统实例UUID

	SubsystemInstanceUUID *string `json:"SubsystemInstanceUUID,omitempty" name:"SubsystemInstanceUUID"`
	// 子系统名称

	SubsystemName *string `json:"SubsystemName,omitempty" name:"SubsystemName"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type InstantiateCommonOperationSheetRequest struct {
	*tchttp.BaseRequest

	//

	SheetIDs []*string `json:"SheetIDs,omitempty" name:"SheetIDs"`
	//

	PlanID *string `json:"PlanID,omitempty" name:"PlanID"`
	//

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
}

func (r *InstantiateCommonOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstantiateCommonOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Dictionary struct {
	//

	Code *string `json:"Code,omitempty" name:"Code"`
	//

	Name *string `json:"Name,omitempty" name:"Name"`
}

type ModifyAppPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAppPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAppPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopySiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ID 云 ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// UUID 云 UUID

		UUID *string `json:"UUID,omitempty" name:"UUID"`
		// Name 云名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopySiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopySiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemConfig struct {
	// Key 配置项

	Key *string `json:"Key,omitempty" name:"Key"`
	// Value 配置值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ApplicationBranch struct {
	// 关联的应用

	Application *Application `json:"Application,omitempty" name:"Application"`
	// 分支对应应用的UUID

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	// 应用分支名，只能小写

	BranchName *string `json:"BranchName,omitempty" name:"BranchName"`
	// 应用分支类型，只有 dev/release

	BranchType *string `json:"BranchType,omitempty" name:"BranchType"`
	// 应用分支级别的的代码分支配置

	CodeBranch *string `json:"CodeBranch,omitempty" name:"CodeBranch"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 删除时间

	DeletedAt *int64 `json:"DeletedAt,omitempty" name:"DeletedAt"`
	// ( primary key ) 自增主键

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 应用分支级别的流水线配置，如果不配置，集成应用级别

	PipelineConfigUUID *string `json:"PipelineConfigUUID,omitempty" name:"PipelineConfigUUID"`
	// 应用的状态 1启用 0废弃，其他状态后续扩展

	State *int64 `json:"State,omitempty" name:"State"`
	// uuid键

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type QueryUploadMaterialTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 上传状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 总体进度

		TotalProgress *float64 `json:"TotalProgress,omitempty" name:"TotalProgress"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 详细处理进度信息

		ProgressDetail []*ProgressDetail `json:"ProgressDetail,omitempty" name:"ProgressDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryUploadMaterialTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryUploadMaterialTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationSheetApprovalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Total *int64 `json:"Total,omitempty" name:"Total"`
		//

		OperationSheetApprovalInfos []*OperationSheetApprovalInfo `json:"OperationSheetApprovalInfos,omitempty" name:"OperationSheetApprovalInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOperationSheetApprovalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationSheetApprovalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChartArtifact struct {
	// chart包内文件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// chart包内文件相对路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// chart包内文件内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

type CheckProductInstanceOperationSheetMaterialRequest struct {
	*tchttp.BaseRequest

	//

	ApplicationList []*ApplicationBase `json:"ApplicationList,omitempty" name:"ApplicationList"`
}

func (r *CheckProductInstanceOperationSheetMaterialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckProductInstanceOperationSheetMaterialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRuntimePlanRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// Type 运行态规划类型 global、product、application

	Type *string `json:"Type,omitempty" name:"Type"`
	// ApplicationInstanceUUID 应用实例UUID，查询应用实例规划需要传

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
}

func (r *GetRuntimePlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRuntimePlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Total *int64 `json:"Total,omitempty" name:"Total"`
		//

		OperationTemplateList []*Template `json:"OperationTemplateList,omitempty" name:"OperationTemplateList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOperationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishCommonOperationSheetRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	//

	IsPublish *bool `json:"IsPublish,omitempty" name:"IsPublish"`
	//

	IsRecursively *bool `json:"IsRecursively,omitempty" name:"IsRecursively"`
}

func (r *PublishCommonOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishCommonOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListChildOperationSheetsRequest struct {
	*tchttp.BaseRequest

	// 父单SheetID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	//

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
	// 分页查询Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页查询Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListChildOperationSheetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListChildOperationSheetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSiteMaterialsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		MaterialList []*SiteMaterial `json:"MaterialList,omitempty" name:"MaterialList"`
		//

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSiteMaterialsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSiteMaterialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AtomicFile struct {
	// DownloadUrl 文件对应的包管理下载地址

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
	// LocalFilePath 文件对应的本地存储地址

	LocalFilePath *string `json:"LocalFilePath,omitempty" name:"LocalFilePath"`
}

type OperationSheetApprovalInfo struct {
	// SheetID 运维单 ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// Approver 审批人

	Approver *string `json:"Approver,omitempty" name:"Approver"`
	// ApprovalComment 审批意见

	ApprovalComment *string `json:"ApprovalComment,omitempty" name:"ApprovalComment"`
	// ApprovalTime 审批时间

	ApprovalTime *string `json:"ApprovalTime,omitempty" name:"ApprovalTime"`
	// ApprovalStatus 审批状态

	ApprovalStatus *string `json:"ApprovalStatus,omitempty" name:"ApprovalStatus"`
	// Sheet 变更单

	Sheet *CommonOperationSheet `json:"Sheet,omitempty" name:"Sheet"`
}

type ApplicationInfo struct {
	// ApplicationUUID 应用UUID

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	// ApplicationName 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// ApplicationVersion 应用版本

	ApplicationVersion *string `json:"ApplicationVersion,omitempty" name:"ApplicationVersion"`
}

type ImportOtaCertificateRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// Content 证书内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *ImportOtaCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportOtaCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MatchOperationTemplateByApplicationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		ApplicationTemplatePairs []*ApplicationTemplatePair `json:"ApplicationTemplatePairs,omitempty" name:"ApplicationTemplatePairs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MatchOperationTemplateByApplicationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MatchOperationTemplateByApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionItem struct {
	//

	RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
	//

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	//

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	//

	RegionNameCh *string `json:"RegionNameCh,omitempty" name:"RegionNameCh"`
	//

	Zones []*ZoneItem `json:"Zones,omitempty" name:"Zones"`
	//

	RoleType *string `json:"RoleType,omitempty" name:"RoleType"`
}

type GetApplicationArtifactChartHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		ChartList []*ChartHistory `json:"ChartList,omitempty" name:"ChartList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationArtifactChartHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationArtifactChartHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSheetTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSheetTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSheetTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanHistoriesRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// Filter 筛选条件：支持SiteUUID, CreatedAt, Creator, ID

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
	// Offset 偏移值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit 返回记录数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListPlanHistoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanHistoriesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryNodeRequest struct {
	*tchttp.BaseRequest

	//

	DagJobUUID *string `json:"DagJobUUID,omitempty" name:"DagJobUUID"`
	//

	DagNodeUUID *string `json:"DagNodeUUID,omitempty" name:"DagNodeUUID"`
}

func (r *RetryNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationDeployInfo struct {
	// CloudUUID 云 UUID

	CloudUUID *string `json:"CloudUUID,omitempty" name:"CloudUUID"`
	// ProductVersionID 产品版本

	ProductVersionID *string `json:"ProductVersionID,omitempty" name:"ProductVersionID"`
	// RegionUUID 地域 UUID

	RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
	// RegionID 地域 ID

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	// RegionNameCN 地域的中文名称

	RegionNameCN *string `json:"RegionNameCN,omitempty" name:"RegionNameCN"`
	// RegionName 地域的名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

type ResumeNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResumeNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectSiteRequest struct {
	*tchttp.BaseRequest

	// SiteID 局点 ID

	SiteID *int64 `json:"SiteID,omitempty" name:"SiteID"`
}

func (r *DeleteProjectSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportProductInstanceOperationSheetRequest struct {
	*tchttp.BaseRequest

	// SheetID 待导出的产品实例变更单ID

	SheetIDs []*string `json:"SheetIDs,omitempty" name:"SheetIDs"`
}

func (r *ExportProductInstanceOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportProductInstanceOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MiddlewareInfo struct {
	// ServiceInstanceName 支撑实例名称

	ServiceInstanceName *string `json:"ServiceInstanceName,omitempty" name:"ServiceInstanceName"`
	// ServiceInstanceAttr 支撑实例属性

	ServiceInstanceAttr *string `json:"ServiceInstanceAttr,omitempty" name:"ServiceInstanceAttr"`
	// ServiceBindingName 服务绑定名称

	ServiceBindingName *string `json:"ServiceBindingName,omitempty" name:"ServiceBindingName"`
	// ServiceBindingAttr 服务绑定属性

	ServiceBindingAttr *string `json:"ServiceBindingAttr,omitempty" name:"ServiceBindingAttr"`
	// SBSecretInfo 服务绑定连接信息

	SBSecretInfo *string `json:"SBSecretInfo,omitempty" name:"SBSecretInfo"`
	// SupportAppInstanceName 支撑应用实例名称

	SupportAppInstanceName *string `json:"SupportAppInstanceName,omitempty" name:"SupportAppInstanceName"`
	// SupportAppInstanceUUID 支撑应用实例UUID

	SupportAppInstanceUUID *string `json:"SupportAppInstanceUUID,omitempty" name:"SupportAppInstanceUUID"`
	// SuppportAppName 支撑应用名称

	SupportAppName *string `json:"SupportAppName,omitempty" name:"SupportAppName"`
	// RegionName 支撑实例所在region

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// ApplicationSet 依赖当前支撑的应用列表

	ApplicationSet []*SupportedAppInfo `json:"ApplicationSet,omitempty" name:"ApplicationSet"`
}

type SheetApplication struct {
	//

	ApplicationID *string `json:"ApplicationID,omitempty" name:"ApplicationID"`
	//

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	//

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
}

type DescribePlanRequest struct {
	*tchttp.BaseRequest

	// 局点UUID,仅指定局点时，默认返回当前生效的规划版本信息

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// Filter 筛选条件：支持ID

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribePlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitOperationSheetApprovalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitOperationSheetApprovalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubmitOperationSheetApprovalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductInstance struct {
	//

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	//

	ProductInstanceName *string `json:"ProductInstanceName,omitempty" name:"ProductInstanceName"`
	//

	SubsystemInstanceSet []*SubsystemInstance `json:"SubsystemInstanceSet,omitempty" name:"SubsystemInstanceSet"`
}

type CopyCommonOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyCommonOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyCommonOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanHistoriesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// PlanHistoryList 规划修改历史列表

		PlanHistoryList []*DescribePlanHistoryResponse `json:"PlanHistoryList,omitempty" name:"PlanHistoryList"`
		// TotalCount 历史记录总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPlanHistoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInitDeploySheetOperationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Template *string `json:"Template,omitempty" name:"Template"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetInitDeploySheetOperationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInitDeploySheetOperationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportDeploymentPhaseApplicationArtifactRequest struct {
	*tchttp.BaseRequest

	//

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	//

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	//

	ApplicationPackageUrl *string `json:"ApplicationPackageUrl,omitempty" name:"ApplicationPackageUrl"`
	//

	ArtifactUUID *string `json:"ArtifactUUID,omitempty" name:"ArtifactUUID"`
	//

	Type *string `json:"Type,omitempty" name:"Type"`
	//

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	//

	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`
}

func (r *ImportDeploymentPhaseApplicationArtifactRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportDeploymentPhaseApplicationArtifactRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductMaterialsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		MaterialList []*ProductMaterialInfo `json:"MaterialList,omitempty" name:"MaterialList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductMaterialsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductMaterialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlansRequest struct {
	*tchttp.BaseRequest

	// 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// Offset 偏移值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit 返回记录数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	//

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
}

func (r *ListPlansRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlansRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationSheet struct {
	// 是否处于激活状态

	Active *bool `json:"Active,omitempty" name:"Active"`
	// 指定的应用实例列表

	ApplicationInstanceList *string `json:"ApplicationInstanceList,omitempty" name:"ApplicationInstanceList"`
	// 变更应用列表

	ApplicationList *string `json:"ApplicationList,omitempty" name:"ApplicationList"`
	// 云ID

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// Tcs集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 结束时间

	EndedAt *string `json:"EndedAt,omitempty" name:"EndedAt"`
	// 执行时间, 单位为秒

	ExecutionTime *int64 `json:"ExecutionTime,omitempty" name:"ExecutionTime"`
	// ( primary key ) 自增主键

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 任务ID

	JobID *string `json:"JobID,omitempty" name:"JobID"`
	// 整个job的完整描述, 包含每个节点的输入输出等信息

	JobRawData *string `json:"JobRawData,omitempty" name:"JobRawData"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 运维任务快照

	PipelineRun *string `json:"PipelineRun,omitempty" name:"PipelineRun"`
	// 规划ID

	PlanID *string `json:"PlanID,omitempty" name:"PlanID"`
	// 产品ID

	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`
	// 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// 产品版本

	ProductVersionID *string `json:"ProductVersionID,omitempty" name:"ProductVersionID"`
	// 区域ID

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	// 运维单ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// 运维单名称

	SheetName *string `json:"SheetName,omitempty" name:"SheetName"`
	// 当前状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 模板快照

	Template *string `json:"Template,omitempty" name:"Template"`
	// 运维类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 运维详细类型（应用部署(applicaiton)、工具组件部署(tool)、原子操作部署(atomic)等）

	TypeDetail *string `json:"TypeDetail,omitempty" name:"TypeDetail"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type TaskSpec struct {
	//

	Description *string `json:"Description,omitempty" name:"Description"`
	//

	YamlDemo *string `json:"YamlDemo,omitempty" name:"YamlDemo"`
	//

	Action []*AtomicParam `json:"Action,omitempty" name:"Action"`
	//

	Result *AtomicParam `json:"Result,omitempty" name:"Result"`
}

type DeleteAtomicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAtomicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAtomicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSheetAttachmentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SheetID 运维单 ID

		SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
		//

		Attachments []*SheetAttachment `json:"Attachments,omitempty" name:"Attachments"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSheetAttachmentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSheetAttachmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMachineSSHConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 机器用户

		User *string `json:"User,omitempty" name:"User"`
		// 机器密码

		Password *string `json:"Password,omitempty" name:"Password"`
		// 机器端口

		Port *int64 `json:"Port,omitempty" name:"Port"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMachineSSHConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMachineSSHConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDeployedApplicationDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Total 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Items app信息列表

		Items []*ApplicationDetail `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListDeployedApplicationDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDeployedApplicationDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyOperationTemplateRequest struct {
	*tchttp.BaseRequest

	// ID 被复制运维模版的主键 ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// Name 新运维模版的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Operator 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *CopyOperationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyOperationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSolutionTemplateRequest struct {
	*tchttp.BaseRequest

	// ID 返回创建的运维模版ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// uuid

	UUID *string `json:"UUID,omitempty" name:"UUID"`
}

func (r *DeleteSolutionTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSolutionTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAtomicRenderDataRequest struct {
	*tchttp.BaseRequest

	// CloudID 云ID

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// RegionUUID 地域UUID

	RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
	// 产品信息如果不提供，则无法渲染产品级别的配置
	// ProductInstanceID或者ProductInstanceUUID 提供其中一个即可。
	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// ProductInstanceUUID 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// ApplicationInstanceID或者ApplicationName 提供其中一个即可
	// ApplicationInstanceID 应用实例ID

	ApplicationInstanceID *string `json:"ApplicationInstanceID,omitempty" name:"ApplicationInstanceID"`
	// ApplicationName 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// AtomicName 原子操作名称

	AtomicName *string `json:"AtomicName,omitempty" name:"AtomicName"`
	// Params 原子操作入参

	Params []*KVPair `json:"Params,omitempty" name:"Params"`
}

func (r *DescribeAtomicRenderDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAtomicRenderDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AtomicKey struct {
	// UUID 原子操作uuid，可单独使用，第一优先级查询

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// AtomicName 原子操作名称，和version一起使用，第二优先级查询

	AtomicName *string `json:"AtomicName,omitempty" name:"AtomicName"`
	// Version 原子操作版本，和name一起使用，第二优先级查询

	AtomicVersion *string `json:"AtomicVersion,omitempty" name:"AtomicVersion"`
}

type AtomicYamlToJsonResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Atomic *JsonAtomic `json:"Atomic,omitempty" name:"Atomic"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AtomicYamlToJsonResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AtomicYamlToJsonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCommonOperationSheetTemplateRequest struct {
	*tchttp.BaseRequest

	// 选择的应用实例列表

	ApplicationInstanceList *string `json:"ApplicationInstanceList,omitempty" name:"ApplicationInstanceList"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// Tcs集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// ( primary key ) 自增主键

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 变更单的待编排项

	OrchestratingItems *string `json:"OrchestratingItems,omitempty" name:"OrchestratingItems"`
	// 模板自定义参数

	Parameters *string `json:"Parameters,omitempty" name:"Parameters"`
	// 产品实例运维单UUID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// 模板内容

	Template *string `json:"Template,omitempty" name:"Template"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 运维单模板用途（部署，升级，变更，回滚）

	Usage *string `json:"Usage,omitempty" name:"Usage"`
}

func (r *UpdateCommonOperationSheetTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCommonOperationSheetTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationTemplateKey struct {
	//

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	//

	ApplicationTemplateName *string `json:"ApplicationTemplateName,omitempty" name:"ApplicationTemplateName"`
}

type CheckMaterialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		MaterialList []*Material `json:"MaterialList,omitempty" name:"MaterialList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckMaterialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePlanHistoryRequest struct {
	*tchttp.BaseRequest

	// 规划历史ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
}

func (r *DescribePlanHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePlanHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSolutionTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ID 返回创建的运维模版ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSolutionTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSolutionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRunningApplicationNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Applications []*ApplicationRunningName `json:"Applications,omitempty" name:"Applications"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRunningApplicationNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRunningApplicationNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationCcDeclareResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ApplicationPackageCcDeclareList 应用包CcDeclare列表

		ApplicationCcDeclareList []*ApplicationCcDeclare `json:"ApplicationCcDeclareList,omitempty" name:"ApplicationCcDeclareList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationCcDeclareResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationCcDeclareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSolutionTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Total *int64 `json:"Total,omitempty" name:"Total"`
		//

		SolutionTemplateList []*SolutionTemplate `json:"SolutionTemplateList,omitempty" name:"SolutionTemplateList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSolutionTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSolutionTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckSiteOperationSheetPlanExistenceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		MissingPlanApplicationList []*ApplicationInfoWithProduct `json:"MissingPlanApplicationList,omitempty" name:"MissingPlanApplicationList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckSiteOperationSheetPlanExistenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSiteOperationSheetPlanExistenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPlanServerInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SuccessSet 成功修改的服务器

		SuccessSet []*ServerDiff `json:"SuccessSet,omitempty" name:"SuccessSet"`
		// 没有修改的服务器信息

		FailedSet []*ServerDiff `json:"FailedSet,omitempty" name:"FailedSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPlanServerInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPlanServerInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSheetAppInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ProductCode 产品英文名

		ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
		// ProductName 产品中文名

		ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
		// ProductVersionName 产品版本名称

		ProductVersionName *string `json:"ProductVersionName,omitempty" name:"ProductVersionName"`
		// ProductInstanceUUID 产品实例 UUID

		ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
		// ProductInstanceID 产品实例 ID

		ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
		//

		ProductInstanceName *string `json:"ProductInstanceName,omitempty" name:"ProductInstanceName"`
		//

		SubsystemInstanceSet []*SubsystemInstanceInfo `json:"SubsystemInstanceSet,omitempty" name:"SubsystemInstanceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSheetAppInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSheetAppInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportPlanResult struct {
	// PlanProductInstanceIDSet 导入的产品实例的ID列表

	PlanProductInstanceSet []*ImportedProductInstance `json:"PlanProductInstanceSet,omitempty" name:"PlanProductInstanceSet"`
	// PlanID 导入的规划ID

	PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
}

type ListCommonOperationSheetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Total *int64 `json:"Total,omitempty" name:"Total"`
		//

		Sheets []*CommonOperationSheet `json:"Sheets,omitempty" name:"Sheets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCommonOperationSheetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCommonOperationSheetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenerateOperationSheetInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenerateOperationSheetInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateOperationSheetInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAppRuntimeNameRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点 UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	//

	ApplicationInstanceSet []*AppRuntimeNameInfo `json:"ApplicationInstanceSet,omitempty" name:"ApplicationInstanceSet"`
	// Modifier  修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
}

func (r *ModifyAppRuntimeNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAppRuntimeNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProductTemplateRequest struct {
	*tchttp.BaseRequest

	//

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	//

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	//

	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`
	//

	Template *string `json:"Template,omitempty" name:"Template"`
}

func (r *ModifyProductTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProductTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialAtomicRequest struct {
	*tchttp.BaseRequest

	//

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
	// ProductVersionUUID 产品版本uuid

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	//

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	//

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	//

	OperationScenes *string `json:"OperationScenes,omitempty" name:"OperationScenes"`
	//

	Label []*KVPair `json:"Label,omitempty" name:"Label"`
	//

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	//

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
}

func (r *ListMaterialAtomicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialAtomicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationPlanUUIDResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationPlanUUIDResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationPlanUUIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppComponentTreeRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// ProductInstanceUUID 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	//

	Location *LocationInfo `json:"Location,omitempty" name:"Location"`
}

func (r *DescribeAppComponentTreeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppComponentTreeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppRuntimeInfoRequest struct {
	*tchttp.BaseRequest

	// Namespace 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// AppNames 应用名称列表

	AppNames []*string `json:"AppNames,omitempty" name:"AppNames"`
}

func (r *DescribeAppRuntimeInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppRuntimeInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportOperationTemplateRequest struct {
	*tchttp.BaseRequest

	//

	OperationTemplateList []*OperationTemplateKey `json:"OperationTemplateList,omitempty" name:"OperationTemplateList"`
}

func (r *ExportOperationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportOperationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RerunNodeRequest struct {
	*tchttp.BaseRequest

	//

	DagJobUUID *string `json:"DagJobUUID,omitempty" name:"DagJobUUID"`
	//

	DagNodeUUID *string `json:"DagNodeUUID,omitempty" name:"DagNodeUUID"`
}

func (r *RerunNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RerunNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanServerInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// PhysicalServerSet 物理机信息列表

		PhysicalServerSet []*PhysicalServerInfo `json:"PhysicalServerSet,omitempty" name:"PhysicalServerSet"`
		// VirtualServerSet 虚拟机信息列表

		VirtualServerSet []*VirtualServerInfo `json:"VirtualServerSet,omitempty" name:"VirtualServerSet"`
		// PmTotalCount 物理机总数

		PmTotalCount *int64 `json:"PmTotalCount,omitempty" name:"PmTotalCount"`
		// VmTotalCount 虚拟机总数

		VmTotalCount *int64 `json:"VmTotalCount,omitempty" name:"VmTotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPlanServerInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanServerInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfirmCommonOperationSheetRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
}

func (r *ConfirmCommonOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfirmCommonOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationSheetApprovalRequest struct {
	*tchttp.BaseRequest

	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	//

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	//

	Filters []*FilterType `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListOperationSheetApprovalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationSheetApprovalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAtomicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAtomicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAtomicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SolutionVersion struct {
	// UUID 解决方案版本的 UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// Code 解决方案版本的 Code

	Code *string `json:"Code,omitempty" name:"Code"`
}

type CheckProductInstanceOperationSheetTemplateAndAtomicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Code *int64 `json:"Code,omitempty" name:"Code"`
		//

		CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckProductInstanceOperationSheetTemplateAndAtomicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckProductInstanceOperationSheetTemplateAndAtomicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLocalTCSCoreIPListRequest struct {
	*tchttp.BaseRequest
}

func (r *GetLocalTCSCoreIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLocalTCSCoreIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MatchOperationTemplateByApplicationRequest struct {
	*tchttp.BaseRequest

	//

	ApplicationNames []*string `json:"ApplicationNames,omitempty" name:"ApplicationNames"`
}

func (r *MatchOperationTemplateByApplicationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MatchOperationTemplateByApplicationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSolutionTemplateRequest struct {
	*tchttp.BaseRequest

	//

	TemplateInfo *SolutionTemplate `json:"TemplateInfo,omitempty" name:"TemplateInfo"`
}

func (r *UpdateSolutionTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSolutionTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComponentInfo struct {
	// ComponentName 模块名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// ComponentInstanceID 模块ID

	ComponentInstanceID *string `json:"ComponentInstanceID,omitempty" name:"ComponentInstanceID"`
}

type ProductCategoryCount struct {
	//

	Category *string `json:"Category,omitempty" name:"Category"`
	//

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type PublishProductInstanceOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishProductInstanceOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishProductInstanceOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationBrief struct {
	// 应用 ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 应用 UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 应用名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type SetApplicationArtifactChartEffectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// chart包生成uuid

		ChartUUID *string `json:"ChartUUID,omitempty" name:"ChartUUID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetApplicationArtifactChartEffectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetApplicationArtifactChartEffectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAtomicRequest struct {
	*tchttp.BaseRequest

	// UUID 原子操作uuid，可单独使用，第一优先级查询

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// AtomicName 原子操作名称，和version一起使用，第二优先级查询

	AtomicName *string `json:"AtomicName,omitempty" name:"AtomicName"`
	// Version 原子操作版本，和name一起使用，第二优先级查询

	AtomicVersion *string `json:"AtomicVersion,omitempty" name:"AtomicVersion"`
}

func (r *DeleteAtomicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAtomicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListClusterRequest struct {
	*tchttp.BaseRequest
}

func (r *ListClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SkipNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SkipNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SkipNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeInfo struct {
	// NodeType 服务器类型,物理机/虚拟机

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// NodeUUID 服务器UUID

	NodeUUID *string `json:"NodeUUID,omitempty" name:"NodeUUID"`
	// 模块实例ID

	ComponentInstanceID *string `json:"ComponentInstanceID,omitempty" name:"ComponentInstanceID"`
}

type ImportCommonOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导入的变更单最顶层父单ID列表

		SheetIDs []*string `json:"SheetIDs,omitempty" name:"SheetIDs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportCommonOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportCommonOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AtomicStep struct {
	// Name 步骤名称

	Name *string `json:"Name,omitempty" name:"Name"`
	//

	Strategy *Strategy `json:"Strategy,omitempty" name:"Strategy"`
	// Description 步骤描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// Executor  步骤执行动作，和AtomicRef，Workflow不可共存

	Executor *string `json:"Executor,omitempty" name:"Executor"`
	// AtomicRef 引用的原子操作，和Executor，Workflow不可共存

	AtomicRef *string `json:"AtomicRef,omitempty" name:"AtomicRef"`
	// Workflow 人工组件标识，和AtomicRef，Executor不可共存

	Workflow *string `json:"Workflow,omitempty" name:"Workflow"`
	// Version 执行动作版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// Inputs 步骤输入

	Inputs []*KVPair `json:"Inputs,omitempty" name:"Inputs"`
	// Result 步骤输出申明

	Result []*AtomicParam `json:"Result,omitempty" name:"Result"`
	// Loop 遍历控制，决策当前原子操作步骤会根据传参遍历

	Loop *LoopStrategy `json:"Loop,omitempty" name:"Loop"`
	// Condition 条件控制，决策当前原子操作步骤是否执行

	Condition *string `json:"Condition,omitempty" name:"Condition"`
}

type DescribeApprovalHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Total *int64 `json:"Total,omitempty" name:"Total"`
		//

		OperationSheetApprovalInfos []*OperationSheetApprovalInfo `json:"OperationSheetApprovalInfos,omitempty" name:"OperationSheetApprovalInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApprovalHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApprovalHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSheetAppInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSheetAppInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSheetAppInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChartContent struct {
	// 文件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 文件所属chart包uuid

	ChartUUID *string `json:"ChartUUID,omitempty" name:"ChartUUID"`
	// 文件内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

type CreateMaterialSyncTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TaskId 同步任务id

		TaskId []*int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMaterialSyncTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialSyncTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ID 云 ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// UUID 云 UUID

		UUID *string `json:"UUID,omitempty" name:"UUID"`
		// Name 云名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppRuntimeYAMLRequest struct {
	*tchttp.BaseRequest

	// Namespace 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// AppName 应用名称

	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

func (r *DescribeAppRuntimeYAMLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppRuntimeYAMLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PauseJobRequest struct {
	*tchttp.BaseRequest

	//

	DagJobUUID *string `json:"DagJobUUID,omitempty" name:"DagJobUUID"`
	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
}

func (r *PauseJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PauseJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationPlanUUIDRequest struct {
	*tchttp.BaseRequest

	// ApplicationName 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// SiteUUID is site uuid

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
}

func (r *GetApplicationPlanUUIDRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationPlanUUIDRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySystemConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySystemConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySystemConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCommonOperationSheetsRequest struct {
	*tchttp.BaseRequest

	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	//

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Filter查询的Key，分为两种：一种是变更单结构中的一级字段,例如Status、SheetKind等；还有是通过AnnotationJson结构中的字段，例如BugID，
	// Key则按照Annotations.BugID的方式去过滤查找

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
	// Orders 排序条件

	Orders []*OrderType `json:"Orders,omitempty" name:"Orders"`
}

func (r *ListCommonOperationSheetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCommonOperationSheetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LatestApplicationPackage struct {
	//

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	//

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	//

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	//

	ApplicationVersion *string `json:"ApplicationVersion,omitempty" name:"ApplicationVersion"`
	//

	SolutionVersionUUID *string `json:"SolutionVersionUUID,omitempty" name:"SolutionVersionUUID"`
	//

	SolutionVersionArtifactTag *int64 `json:"SolutionVersionArtifactTag,omitempty" name:"SolutionVersionArtifactTag"`
}

type PackageData struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 数据包URL

	URL *string `json:"URL,omitempty" name:"URL"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 拓展信息

	Extensions *string `json:"Extensions,omitempty" name:"Extensions"`
}

type ListAppRuntimeNamesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ProductInstanceSet 产品实例列表

		ApplicationSet []*AppRuntimeNameInfo `json:"ApplicationSet,omitempty" name:"ApplicationSet"`
		// TotalCount 应用总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAppRuntimeNamesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppRuntimeNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductTemplateIDRequest struct {
	*tchttp.BaseRequest

	// ProductVersionUUID 产品版本 UUID

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	// Kind 模板类型 取值为 deploy update upgrade

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// Type 描述模版类别,取值为product application等

	Type *string `json:"Type,omitempty" name:"Type"`
	// Labels 产品版本 UUID

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
}

func (r *GetProductTemplateIDRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductTemplateIDRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeploySingleAppRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点的 UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// ProductInstanceID 产品实例的 ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// Name 部署对象实例名称，应用/工具组件/原子操作名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// ApplicationName 部署对象名称，应用/工具组件/原子操作名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// Type 部署对象类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// ApplicationExtraInfo 应用的详细信息

	ApplicationExtraInfo *ApplicationDeployInfo `json:"ApplicationExtraInfo,omitempty" name:"ApplicationExtraInfo"`
}

func (r *DeploySingleAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeploySingleAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSiteMaterialInfoRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点 UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
}

func (r *GetSiteMaterialInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSiteMaterialInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSubOperationSheetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Sheets []*ProductInstanceOperationSheet `json:"Sheets,omitempty" name:"Sheets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSubOperationSheetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSubOperationSheetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyChartFileRequest struct {
	*tchttp.BaseRequest

	// 应用制品uuid

	ApplicationArtifactUUID *string `json:"ApplicationArtifactUUID,omitempty" name:"ApplicationArtifactUUID"`
	// chart包UUID

	ChartUUID *string `json:"ChartUUID,omitempty" name:"ChartUUID"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 修改文件内容

	Files []*ChartArtifact `json:"Files,omitempty" name:"Files"`
}

func (r *ModifyChartFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyChartFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCommonOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 额外信息

		Annotations *string `json:"Annotations,omitempty" name:"Annotations"`
		// 云ID, 产品部署运维单依赖该字段查询

		CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
		// 创建时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 创建人

		Creator *string `json:"Creator,omitempty" name:"Creator"`
		// 描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// ( primary key ) 自增主键

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// 变更影响

		Influence *string `json:"Influence,omitempty" name:"Influence"`
		// 是否进行展示

		IsDisplay *int64 `json:"IsDisplay,omitempty" name:"IsDisplay"`
		// 是否是预置变更单

		IsInternal *int64 `json:"IsInternal,omitempty" name:"IsInternal"`
		// 修改人

		Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
		// 附件信息

		OperationSheetAttachment []*OperationSheetAttachment `json:"OperationSheetAttachment,omitempty" name:"OperationSheetAttachment"`
		// 父子单关联关系，仅当此单为父单时有效

		OperationSheetRelation []*OperationSheetRelation `json:"OperationSheetRelation,omitempty" name:"OperationSheetRelation"`
		// 变更单具体的编排信息, 一个变更单默认只有一个编排模板，预留model_array用于扩展

		OperationSheetTemplate []*OperationSheetTemplate `json:"OperationSheetTemplate,omitempty" name:"OperationSheetTemplate"`
		// 待编排项

		OrchestratingItems *string `json:"OrchestratingItems,omitempty" name:"OrchestratingItems"`
		// 负责人

		Owners *string `json:"Owners,omitempty" name:"Owners"`
		// 产品实例ID, 产品部署运维单依赖该字段查询

		ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
		// 区域ID, 产品部署运维单依赖该字段查询

		RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
		// 风险项

		Risk *string `json:"Risk,omitempty" name:"Risk"`
		// 场景

		Scene *string `json:"Scene,omitempty" name:"Scene"`
		// 运维单ID

		SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
		// 运维单类型（common,product,site,productInstance）

		SheetKind *string `json:"SheetKind,omitempty" name:"SheetKind"`
		// 运维单名称

		SheetName *string `json:"SheetName,omitempty" name:"SheetName"`
		// 运维模板类型(deploy,update,upgrade)

		SheetType *string `json:"SheetType,omitempty" name:"SheetType"`
		// 运维单来源:bug单, 交付单, 导入，人工创建

		Source *string `json:"Source,omitempty" name:"Source"`
		// 状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 更新时间

		UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
		// AnnotationsVar是Annotations的结构体化展示

		AnnotationsVar *AnnotationsVar `json:"AnnotationsVar,omitempty" name:"AnnotationsVar"`
		// OrchestratingItemsVar是OrchestratingItems的结构体化展示

		OrchestratingItemsVar *OrchestratingItemsVar `json:"OrchestratingItemsVar,omitempty" name:"OrchestratingItemsVar"`
		// Children 是子变更单基本信息

		Children []*CommonOperationSheet `json:"Children,omitempty" name:"Children"`
		// Templates 是变更单编排相关的信息

		Templates []*OperationSheetTemplate `json:"Templates,omitempty" name:"Templates"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCommonOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCommonOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationDataSchemaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Schema 数据格式定义

		Schema *string `json:"Schema,omitempty" name:"Schema"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOperationDataSchemaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationDataSchemaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAtomicRequest struct {
	*tchttp.BaseRequest

	// UUID 原子操作uuid，可单独使用，第一优先级查询

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// AtomicName 原子操作名称，和version一起使用，第二优先级查询

	AtomicName *string `json:"AtomicName,omitempty" name:"AtomicName"`
	// Version 原子操作版本，和name一起使用，第二优先级查询

	AtomicVersion *string `json:"AtomicVersion,omitempty" name:"AtomicVersion"`
}

func (r *DescribeAtomicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAtomicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAtomicLabelRequest struct {
	*tchttp.BaseRequest
}

func (r *ListAtomicLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAtomicLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCommonOperationSheetRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// 是否要递归删除关联的子单，前提是子单没有被别的父单引用

	IsRecursively *bool `json:"IsRecursively,omitempty" name:"IsRecursively"`
}

func (r *DeleteCommonOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCommonOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiffPlanPackageRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// ProductInstanceUUID 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// FirstPlan 被对比的规划版本信息

	FirstPlan *PlanVersionState `json:"FirstPlan,omitempty" name:"FirstPlan"`
	// SecondPlan 对比的规划版本

	SecondPlan *PlanVersionState `json:"SecondPlan,omitempty" name:"SecondPlan"`
}

func (r *DiffPlanPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DiffPlanPackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SiteProductApplication struct {
	// 应用UUID

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	// 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
}

type GetLocalTCSCoreIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		MasterIPList []*string `json:"MasterIPList,omitempty" name:"MasterIPList"`
		//

		NodeIPList []*string `json:"NodeIPList,omitempty" name:"NodeIPList"`
		//

		MaterialNodeIP *string `json:"MaterialNodeIP,omitempty" name:"MaterialNodeIP"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLocalTCSCoreIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLocalTCSCoreIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductApplicationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用列表

		ApplicationList []*ApplicationBrief `json:"ApplicationList,omitempty" name:"ApplicationList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductApplicationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductRuntimeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		ProductInstanceConfig *string `json:"ProductInstanceConfig,omitempty" name:"ProductInstanceConfig"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductRuntimeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductRuntimeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelateCommonOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RelateCommonOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelateCommonOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationBase struct {
	//

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	//

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	//

	ApplicationVersion *string `json:"ApplicationVersion,omitempty" name:"ApplicationVersion"`
}

type SheetApplicationInstance struct {
	//

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	//

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
	//

	TargetApplicationVersion *string `json:"TargetApplicationVersion,omitempty" name:"TargetApplicationVersion"`
	//

	PlanTag *int64 `json:"PlanTag,omitempty" name:"PlanTag"`
	//

	DeployVersionID *int64 `json:"DeployVersionID,omitempty" name:"DeployVersionID"`
	//

	DeployVersionIDBeforeUpdate *int64 `json:"DeployVersionIDBeforeUpdate,omitempty" name:"DeployVersionIDBeforeUpdate"`
}

type ListProductsRequest struct {
	*tchttp.BaseRequest

	//

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	//

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	//

	MultiCloud *int64 `json:"MultiCloud,omitempty" name:"MultiCloud"`
	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	//

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	//

	Orders []*OrderType `json:"Orders,omitempty" name:"Orders"`
	//

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// ExcludeProductList 排除的产品列表

	ExcludeProductList []*string `json:"ExcludeProductList,omitempty" name:"ExcludeProductList"`
}

func (r *ListProductsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Task struct {
	//

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	//

	Version *string `json:"Version,omitempty" name:"Version"`
	//

	Spec *TaskSpec `json:"Spec,omitempty" name:"Spec"`
}

type GenerateOperationSheetInstanceRequest struct {
	*tchttp.BaseRequest

	//

	SheetId *string `json:"SheetId,omitempty" name:"SheetId"`
}

func (r *GenerateOperationSheetInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateOperationSheetInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GeneratePipelineRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	//

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	//

	Template *string `json:"Template,omitempty" name:"Template"`
	//

	Deep *int64 `json:"Deep,omitempty" name:"Deep"`
}

func (r *GeneratePipelineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GeneratePipelineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialApplicationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Total *int64 `json:"Total,omitempty" name:"Total"`
		//

		ApplicationList []*ApplicationInfo `json:"ApplicationList,omitempty" name:"ApplicationList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialApplicationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Total *int64 `json:"Total,omitempty" name:"Total"`
		//

		Sheets []*CommonOperationSheet `json:"Sheets,omitempty" name:"Sheets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployEnvRequest struct {
	*tchttp.BaseRequest

	// EnvName 环境名称

	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`
	// Description 环境描述描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// ConnectUrl 环境连接信息

	ConnectUrl *string `json:"ConnectUrl,omitempty" name:"ConnectUrl"`
	// AuroraEnvUUID 环境管理 UUID

	AuroraEnvUUID *string `json:"AuroraEnvUUID,omitempty" name:"AuroraEnvUUID"`
	// InstallConfig 环境部署配置

	InstallConfig *string `json:"InstallConfig,omitempty" name:"InstallConfig"`
	// MaterialUUID 标准物料 UUID

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
	// MaterialURL 物料下载地址

	MaterialURL *string `json:"MaterialURL,omitempty" name:"MaterialURL"`
	// ImportMaterial 是否需要导入物料, 0 不需要导入, 1 需要导入

	ImportMaterial *int64 `json:"ImportMaterial,omitempty" name:"ImportMaterial"`
}

func (r *DeployEnvRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeployEnvRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationTemplateRequest struct {
	*tchttp.BaseRequest

	// ID 被删除运维模版的主键 ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// UUID 被删除运维模版的名称

	UUID *string `json:"UUID,omitempty" name:"UUID"`
}

func (r *DeleteOperationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductInstanceOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 额外信息

		Annotations *string `json:"Annotations,omitempty" name:"Annotations"`
		// 应用列表

		Applications *string `json:"Applications,omitempty" name:"Applications"`
		// 云ID

		CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
		// 创建时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 创建者

		Creator *string `json:"Creator,omitempty" name:"Creator"`
		// 运维单描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// ( primary key ) 自增主键

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// 是否内置运维单

		IsInternal *bool `json:"IsInternal,omitempty" name:"IsInternal"`
		// 修改者

		Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
		// 父子变更单关联关系，仅当此变更单为子变更单时有效

		OperationSheetRelation []*OperationSheetRelation `json:"OperationSheetRelation,omitempty" name:"OperationSheetRelation"`
		// 变更单具体的编排信息, 一个变更单默认只有一个编排模板，预留model_array用于扩展

		OperationSheetTemplate []*OperationSheetTemplate `json:"OperationSheetTemplate,omitempty" name:"OperationSheetTemplate"`
		// 规划ID

		PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
		// 产品实例ID

		ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
		// 产品实例UUID

		ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
		// 产品名称

		ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
		// 产品运维单UUID

		ProductSheetID *string `json:"ProductSheetID,omitempty" name:"ProductSheetID"`
		// 产品UUID

		ProductUUID *string `json:"ProductUUID,omitempty" name:"ProductUUID"`
		// 产品版本制品Tag

		ProductVersionArtifactTag *string `json:"ProductVersionArtifactTag,omitempty" name:"ProductVersionArtifactTag"`
		// 产品版本名称

		ProductVersionName *string `json:"ProductVersionName,omitempty" name:"ProductVersionName"`
		// 产品版本UUID

		ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
		// 区域UUID

		RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
		// 风险信息

		Risk *string `json:"Risk,omitempty" name:"Risk"`
		// 场景

		Scene *string `json:"Scene,omitempty" name:"Scene"`
		// 产品实例运维单UUID

		SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
		// 运维单名称

		SheetName *string `json:"SheetName,omitempty" name:"SheetName"`
		// 运维单类型

		SheetType *string `json:"SheetType,omitempty" name:"SheetType"`
		// 局点UUID

		SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
		// 运维单来源:导入，人工创建

		Source *string `json:"Source,omitempty" name:"Source"`
		// 运维单状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 更新时间

		UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
		//

		ApplicationList []*SheetApplication `json:"ApplicationList,omitempty" name:"ApplicationList"`
		//

		Templates []*OperationSheetTemplate `json:"Templates,omitempty" name:"Templates"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductInstanceOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductInstanceOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubsystemInstance struct {
	//

	SubsystemInstanceID *int64 `json:"SubsystemInstanceID,omitempty" name:"SubsystemInstanceID"`
	//

	SubsystemInstanceName *string `json:"SubsystemInstanceName,omitempty" name:"SubsystemInstanceName"`
}

type ModifyProductPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProductPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProductPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiffApplicationInstanceConfigRequest struct {
	*tchttp.BaseRequest

	//

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	//

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	//

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	//

	ApplicationUUIDList []*string `json:"ApplicationUUIDList,omitempty" name:"ApplicationUUIDList"`
	//

	ApplicationInstanceUUIDList []*string `json:"ApplicationInstanceUUIDList,omitempty" name:"ApplicationInstanceUUIDList"`
}

func (r *DiffApplicationInstanceConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DiffApplicationInstanceConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParseProductDataRequest struct {
	*tchttp.BaseRequest

	// package 包

	Package *string `json:"Package,omitempty" name:"Package"`
	// MD5 md5值

	MD5 *string `json:"MD5,omitempty" name:"MD5"`
}

func (r *ParseProductDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ParseProductDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationRunningName struct {
	// ApplicationInstanceUUID 应用实例UUID

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
	// ApplicationName 应用实例可能名称

	PossibleApplicationInstanceNames []*string `json:"PossibleApplicationInstanceNames,omitempty" name:"PossibleApplicationInstanceNames"`
	// ApplicationInstanceName 应用实例确定名称

	ApplicationInstanceName *string `json:"ApplicationInstanceName,omitempty" name:"ApplicationInstanceName"`
}

type ExportServerResourcesRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// FileName 导出文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

func (r *ExportServerResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportServerResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MissingApplicationPackage struct {
	//

	OperationHistoryID *int64 `json:"OperationHistoryID,omitempty" name:"OperationHistoryID"`
	//

	ApplicationPackageList []*LatestApplicationPackage `json:"ApplicationPackageList,omitempty" name:"ApplicationPackageList"`
}

type CreateImportMaterialTaskRequest struct {
	*tchttp.BaseRequest

	// 物料链接

	MaterialURL *string `json:"MaterialURL,omitempty" name:"MaterialURL"`
	// 物料所在服务器

	MaterialServer *string `json:"MaterialServer,omitempty" name:"MaterialServer"`
	// 物料父目录

	MaterialPath *string `json:"MaterialPath,omitempty" name:"MaterialPath"`
	// 操作者

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *CreateImportMaterialTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImportMaterialTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppRuntimeConfigRequest struct {
	*tchttp.BaseRequest

	//

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	//

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	//

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	//

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
}

func (r *DescribeAppRuntimeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppRuntimeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckProductInstanceOperationSheetMaterialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Code *int64 `json:"Code,omitempty" name:"Code"`
		//

		CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckProductInstanceOperationSheetMaterialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckProductInstanceOperationSheetMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOperationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOperationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOperationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSheetTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Template *string `json:"Template,omitempty" name:"Template"`
		//

		Usage *string `json:"Usage,omitempty" name:"Usage"`
		//

		Parameters *string `json:"Parameters,omitempty" name:"Parameters"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSheetTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSheetTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZoneItem struct {
	//

	ZoneUUID *string `json:"ZoneUUID,omitempty" name:"ZoneUUID"`
	//

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	//

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	//

	ZoneNameCh *string `json:"ZoneNameCh,omitempty" name:"ZoneNameCh"`
	//

	RoleType *string `json:"RoleType,omitempty" name:"RoleType"`
}

type ListMaterialApplicationRequest struct {
	*tchttp.BaseRequest

	// MaterialUUID 物料UUID

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
	// ProductVersionUUID 产品版本uuid

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	// Limit 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Offset 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Filter 筛选

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
}

func (r *ListMaterialApplicationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialApplicationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportSolutionTemplateRequest struct {
	*tchttp.BaseRequest

	//

	SolutionTemplateList []*OperationTemplateKey `json:"SolutionTemplateList,omitempty" name:"SolutionTemplateList"`
}

func (r *ExportSolutionTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportSolutionTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportStandardMaterialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TaskID 任务ID

		TaskID *int64 `json:"TaskID,omitempty" name:"TaskID"`
		// ImportSuccessRecord 是否需要导入, 0 表示当前没有导入成功记录, 1 表示当前产品市场有成功导入记录

		ImportSuccessRecord *int64 `json:"ImportSuccessRecord,omitempty" name:"ImportSuccessRecord"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportStandardMaterialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportStandardMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationCcDeclare struct {
	// ApplicationUUID 应用UUID

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	// ApplicationName 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// ApplicationVersion 应用版本

	ApplicationVersion *string `json:"ApplicationVersion,omitempty" name:"ApplicationVersion"`
	// Content ccDeclare文件内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

type ClusterInfo struct {
	// Name 集群名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// ProductCount 集群部署的产品数量

	ProductCount *int64 `json:"ProductCount,omitempty" name:"ProductCount"`
}

type ExportSolutionTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportSolutionTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportSolutionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationArtifactChartHistoryRequest struct {
	*tchttp.BaseRequest

	// 应用制品uuid

	ApplicationArtifactUUID *string `json:"ApplicationArtifactUUID,omitempty" name:"ApplicationArtifactUUID"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *GetApplicationArtifactChartHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationArtifactChartHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectSiteInfoDetailRequest struct {
	*tchttp.BaseRequest
}

func (r *GetProjectSiteInfoDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectSiteInfoDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationInstanceForImportApplicationPlan struct {
	//

	CloudUUID *string `json:"CloudUUID,omitempty" name:"CloudUUID"`
	//

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	//

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	//

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	//

	PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
	// ProductName 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// ProductCode 产品英文名

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// ProductID 产品ID

	ProductUUID *string `json:"ProductUUID,omitempty" name:"ProductUUID"`
	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// ProductInstanceName 产品实例名称

	ProductInstanceName *string `json:"ProductInstanceName,omitempty" name:"ProductInstanceName"`
	// ProductVersionID 产品版本ID

	ProductVersionID *string `json:"ProductVersionID,omitempty" name:"ProductVersionID"`
	// ProductVersionName 产品版本名称

	ProductVersionName *string `json:"ProductVersionName,omitempty" name:"ProductVersionName"`
	// ProductVersionDataTag 产品版本的数据Tag

	ProductVersionDataTag *string `json:"ProductVersionDataTag,omitempty" name:"ProductVersionDataTag"`
	// ProductCategory 产品分类

	ProductCategory *string `json:"ProductCategory,omitempty" name:"ProductCategory"`
	// ApplicationInstanceUUID 应用实例UUID

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
	// ApplicationInstanceID 应用实例ID

	ApplicationInstanceID *string `json:"ApplicationInstanceID,omitempty" name:"ApplicationInstanceID"`
	// ApplicationInstanceName 应用实例名称

	ApplicationInstanceName *string `json:"ApplicationInstanceName,omitempty" name:"ApplicationInstanceName"`
	// ApplicationBranchName 应用分支名称

	ApplicationBranchName *string `json:"ApplicationBranchName,omitempty" name:"ApplicationBranchName"`
	// ApplicationBranchID 应用分支ID

	ApplicationBranchID *string `json:"ApplicationBranchID,omitempty" name:"ApplicationBranchID"`
	// Location 实例部署位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// Config 应用配置:values.yaml

	Config *string `json:"Config,omitempty" name:"Config"`
	// AppComponentSet 应用模块信息列表

	AppComponentSet []*Component `json:"AppComponentSet,omitempty" name:"AppComponentSet"`
	// AppServerSet 应用服务器信息列表

	AppServerSet []*Server `json:"AppServerSet,omitempty" name:"AppServerSet"`
	// RegionUUID region id

	RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
	// ZoneUUID zone id

	ZoneUUID *string `json:"ZoneUUID,omitempty" name:"ZoneUUID"`
}

type InstanceTreeSubsystem struct {
	//

	SubsystemCode *string `json:"SubsystemCode,omitempty" name:"SubsystemCode"`
	//

	SubsystemName *string `json:"SubsystemName,omitempty" name:"SubsystemName"`
	//

	SubsystemInstanceUUID *string `json:"SubsystemInstanceUUID,omitempty" name:"SubsystemInstanceUUID"`
	//

	SubsystemInstanceID *string `json:"SubsystemInstanceID,omitempty" name:"SubsystemInstanceID"`
	//

	SubsystemInstanceName *string `json:"SubsystemInstanceName,omitempty" name:"SubsystemInstanceName"`
	//

	ApplicationInstanceSet []*InstanceTreeApplication `json:"ApplicationInstanceSet,omitempty" name:"ApplicationInstanceSet"`
}

type ExportProductInstanceOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 变更单包下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 变更单包MD5信息

		MD5 *string `json:"MD5,omitempty" name:"MD5"`
		// 依赖的原子操作由运维模板进行管理, 不进行二次遍历，仅遍历当前编排中的内容

		AtomicKeyList []*AtomicKey `json:"AtomicKeyList,omitempty" name:"AtomicKeyList"`
		// 依赖的运维模板，一般是应用运维模板

		OperationTemplateKeyList []*OperationTemplateKey `json:"OperationTemplateKeyList,omitempty" name:"OperationTemplateKeyList"`
		// 应用名称列表

		ApplicationNameList []*string `json:"ApplicationNameList,omitempty" name:"ApplicationNameList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportProductInstanceOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportProductInstanceOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartOperationSheetRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	//

	PipelineRun *string `json:"PipelineRun,omitempty" name:"PipelineRun"`
}

func (r *StartOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoopStrategy struct {
	// Items is a json format string containing a list of object/string/int,
	// e.g. "[{\"Zone\":\"ABC\"}, {\"Zone\":\"DEF\"}]"
	// Each object/string/int will cause a sub-node. In node.spec.input, user can use `item` to
	// reference each item in the `Items` list. e.g: {"Key":"Zone", "Value": "<no value>"}

	Items *string `json:"Items,omitempty" name:"Items"`
	// Parallel decides sub-nodes dependency

	Parallel *bool `json:"Parallel,omitempty" name:"Parallel"`
}

type ListSiteOperationSheetsRequest struct {
	*tchttp.BaseRequest

	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	//

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Filter查询的Key，分为两种：一种是变更单结构中的一级字段,例如Status、SheetKind等；还有是通过AnnotationJson结构中的字段，例如BugID，
	// Key则按照Annotations.BugID的方式去过滤查找

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
	// Orders 排序条件

	Orders []*OrderType `json:"Orders,omitempty" name:"Orders"`
}

func (r *ListSiteOperationSheetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSiteOperationSheetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductRollbackConfigRequest struct {
	*tchttp.BaseRequest

	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	//

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
}

func (r *DescribeProductRollbackConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductRollbackConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PlanBriefInfo struct {
	// 规划协议版本

	APIVersion *string `json:"APIVersion,omitempty" name:"APIVersion"`
	// 云拓扑的ID

	CloudID *int64 `json:"CloudID,omitempty" name:"CloudID"`
	// 规划commit的UUID

	CommitUUID *string `json:"CommitUUID,omitempty" name:"CommitUUID"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 删除时间

	DeletedAt *int64 `json:"DeletedAt,omitempty" name:"DeletedAt"`
	// ( primary key ) 规划ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 规划名称

	Name *string `json:"Name,omitempty" name:"Name"`
	//

	PlanProductInstance []*PlanProductInstance `json:"PlanProductInstance,omitempty" name:"PlanProductInstance"`
	//

	PlanServerInfo []*PlanServerInfo `json:"PlanServerInfo,omitempty" name:"PlanServerInfo"`
	// 做规划的时间

	PlanTime *string `json:"PlanTime,omitempty" name:"PlanTime"`
	// 规划包的类型:new/expand/upgrade

	PlanType *string `json:"PlanType,omitempty" name:"PlanType"`
	// 局点名称

	SiteName *string `json:"SiteName,omitempty" name:"SiteName"`
	// 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// 规划使用的解决方案Code

	SolutionCode *string `json:"SolutionCode,omitempty" name:"SolutionCode"`
	// 规划使用的解决方案名称

	SolutionName *string `json:"SolutionName,omitempty" name:"SolutionName"`
	// 规划时使用的解决方案版本数据Tag

	SolutionVersionDataTag *string `json:"SolutionVersionDataTag,omitempty" name:"SolutionVersionDataTag"`
	// 规划使用的解决方案版本ID

	SolutionVersionID *string `json:"SolutionVersionID,omitempty" name:"SolutionVersionID"`
	// 规划使用的解决方案版本名称

	SolutionVersionName *string `json:"SolutionVersionName,omitempty" name:"SolutionVersionName"`
	// 规划包中的云拓扑信息

	Topology *string `json:"Topology,omitempty" name:"Topology"`
	// 规划的类型:standard/site

	Type *string `json:"Type,omitempty" name:"Type"`
	// 规划UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 规划版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 解决方案版本数据Tag创建时间

	VersionTagCreationTime *string `json:"VersionTagCreationTime,omitempty" name:"VersionTagCreationTime"`
	// 规划包名称

	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`
	// MaterialUUID 所在物料包的物料UUID

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
	// 所在物料包的名称

	MaterialName *string `json:"MaterialName,omitempty" name:"MaterialName"`
	// IsActivePlan 是否为当前生效版本

	IsActivePlan *bool `json:"IsActivePlan,omitempty" name:"IsActivePlan"`
}

type DescribeAppRuntimeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		ApplicationInstanceConfig *string `json:"ApplicationInstanceConfig,omitempty" name:"ApplicationInstanceConfig"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppRuntimeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppRuntimeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePlanHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 操作人

		Creator *string `json:"Creator,omitempty" name:"Creator"`
		// ( primary key ) 规划修改历史ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// 操作类型,可选：导入、修改

		OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
		// 规划ID

		PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
		// 局点UUID

		SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
		// PlanPkgInfo 导入规划包信息

		PlanPkgInfo *PlanBriefInfo `json:"PlanPkgInfo,omitempty" name:"PlanPkgInfo"`
		// Modifications 修改前后数据对比

		Modifications []*PlanModification `json:"Modifications,omitempty" name:"Modifications"`
		// OperationObjects 操作对象

		OperationObjects []*string `json:"OperationObjects,omitempty" name:"OperationObjects"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePlanHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePlanHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationUUIDAndName struct {
	// ApplicationInstanceUUID 应用实例UUID

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
	// ApplicationName 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
}

type PlanServerInfo struct {
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 删除时间

	DeletedAt *int64 `json:"DeletedAt,omitempty" name:"DeletedAt"`
	// 服务器磁盘信息

	DiskInfo *string `json:"DiskInfo,omitempty" name:"DiskInfo"`
	// 额外服务器属性

	ExtraAttribute *string `json:"ExtraAttribute,omitempty" name:"ExtraAttribute"`
	// 宿主机ID

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 宿主机ID

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 宿主机UUID

	HostUUID *string `json:"HostUUID,omitempty" name:"HostUUID"`
	// ( primary key ) ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 服务器IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 应用实例部署位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// 服务器配置信息

	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 服务器操作系统信息

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 规划ID

	PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
	// 服务器ID

	ServerID *string `json:"ServerID,omitempty" name:"ServerID"`
	// 服务器型号

	ServerType *string `json:"ServerType,omitempty" name:"ServerType"`
	// 服务器类型(虚拟机或物理机)

	Type *string `json:"Type,omitempty" name:"Type"`
	// 服务器UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type ListAtomicToPipelineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		AtomicPipelineYaml *string `json:"AtomicPipelineYaml,omitempty" name:"AtomicPipelineYaml"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAtomicToPipelineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAtomicToPipelineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLatestApplicationPackagesRequest struct {
	*tchttp.BaseRequest

	// ApplicationUUIDList 应用的 UUID 列表

	ApplicationUUIDList []*string `json:"ApplicationUUIDList,omitempty" name:"ApplicationUUIDList"`
	// WithArtifact 是否包含制品信息

	WithArtifact *bool `json:"WithArtifact,omitempty" name:"WithArtifact"`
}

func (r *ListLatestApplicationPackagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLatestApplicationPackagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规划协议版本

		APIVersion *string `json:"APIVersion,omitempty" name:"APIVersion"`
		// 云拓扑的ID

		CloudID *int64 `json:"CloudID,omitempty" name:"CloudID"`
		// 规划commit的UUID

		CommitUUID *string `json:"CommitUUID,omitempty" name:"CommitUUID"`
		// 创建时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 创建人

		Creator *string `json:"Creator,omitempty" name:"Creator"`
		// 删除时间

		DeletedAt *int64 `json:"DeletedAt,omitempty" name:"DeletedAt"`
		// ( primary key ) 规划ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// 修改人

		Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
		// 规划名称

		Name *string `json:"Name,omitempty" name:"Name"`
		//

		PlanProductInstance []*PlanProductInstance `json:"PlanProductInstance,omitempty" name:"PlanProductInstance"`
		//

		PlanServerInfo []*PlanServerInfo `json:"PlanServerInfo,omitempty" name:"PlanServerInfo"`
		// 做规划的时间

		PlanTime *string `json:"PlanTime,omitempty" name:"PlanTime"`
		// 规划包的类型:new/expand/upgrade

		PlanType *string `json:"PlanType,omitempty" name:"PlanType"`
		// 局点名称

		SiteName *string `json:"SiteName,omitempty" name:"SiteName"`
		// 局点UUID

		SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
		// 规划使用的解决方案Code

		SolutionCode *string `json:"SolutionCode,omitempty" name:"SolutionCode"`
		// 规划使用的解决方案名称

		SolutionName *string `json:"SolutionName,omitempty" name:"SolutionName"`
		// 规划时使用的解决方案版本数据Tag

		SolutionVersionDataTag *string `json:"SolutionVersionDataTag,omitempty" name:"SolutionVersionDataTag"`
		// 规划使用的解决方案版本ID

		SolutionVersionID *string `json:"SolutionVersionID,omitempty" name:"SolutionVersionID"`
		// 规划使用的解决方案版本名称

		SolutionVersionName *string `json:"SolutionVersionName,omitempty" name:"SolutionVersionName"`
		// 规划包中的云拓扑信息

		Topology *string `json:"Topology,omitempty" name:"Topology"`
		// 规划的类型:standard/site

		Type *string `json:"Type,omitempty" name:"Type"`
		// 规划UUID

		UUID *string `json:"UUID,omitempty" name:"UUID"`
		// 更新时间

		UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
		// 规划版本

		Version *string `json:"Version,omitempty" name:"Version"`
		// 解决方案版本数据Tag创建时间

		VersionTagCreationTime *string `json:"VersionTagCreationTime,omitempty" name:"VersionTagCreationTime"`
		// 规划包名称

		PackageName *string `json:"PackageName,omitempty" name:"PackageName"`
		// MaterialUUID 所在物料包的物料UUID

		MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
		// 所在物料包的名称

		MaterialName *string `json:"MaterialName,omitempty" name:"MaterialName"`
		// IsActivePlan 是否为当前生效版本

		IsActivePlan *bool `json:"IsActivePlan,omitempty" name:"IsActivePlan"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductInfoRequest struct {
	*tchttp.BaseRequest

	// 产品 UUID

	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`
}

func (r *GetProductInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppDiff struct {
	//

	ApplicationCode *string `json:"ApplicationCode,omitempty" name:"ApplicationCode"`
	//

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
	//

	ApplicationInstanceID *string `json:"ApplicationInstanceID,omitempty" name:"ApplicationInstanceID"`
	//

	ConfigUpdate *bool `json:"ConfigUpdate,omitempty" name:"ConfigUpdate"`
}

type ApplicationPackage struct {
	//

	Application *Application `json:"Application,omitempty" name:"Application"`
	//

	ApplicationArtifact []*ApplicationArtifact `json:"ApplicationArtifact,omitempty" name:"ApplicationArtifact"`
	// 应用的branch注意这里是极光应用管理的概念，不是代码仓库的branch

	ApplicationBranch *string `json:"ApplicationBranch,omitempty" name:"ApplicationBranch"`
	//

	ApplicationBranchDetail *ApplicationBranch `json:"ApplicationBranchDetail,omitempty" name:"ApplicationBranchDetail"`
	// 应用管理里的分支的UUID，ApplicationBranch是分支在构建的时候用的名字

	ApplicationBranchUUID *string `json:"ApplicationBranchUUID,omitempty" name:"ApplicationBranchUUID"`
	// 1代表从applicationbranch的codebranch构建的包(代码release分支) 2从其他分支用这个applicationbranch构建的包(代码的开发分支)

	ApplicationPackageType *int64 `json:"ApplicationPackageType,omitempty" name:"ApplicationPackageType"`
	// 应用标准，tad, ted-product, ted-dbsql, ted-preset, ted-image, ted-yunapi, ted-tool

	ApplicationStandard *string `json:"ApplicationStandard,omitempty" name:"ApplicationStandard"`
	// 打包该应用包使用的应用的模板的id

	ApplicationTemplateID *string `json:"ApplicationTemplateID,omitempty" name:"ApplicationTemplateID"`
	// 打包该应用包使用的应用的模板的revision，如果是latest也要获取当前的那个revision

	ApplicationTemplateRevision *string `json:"ApplicationTemplateRevision,omitempty" name:"ApplicationTemplateRevision"`
	// 打包该应用包使用的应用的模板的version

	ApplicationTemplateVersion *string `json:"ApplicationTemplateVersion,omitempty" name:"ApplicationTemplateVersion"`
	// 关联应用管理里的application uuid

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	// 应用包的版本从Chart.yaml里来，这里也要保存一下revision

	ApplicationVersion *string `json:"ApplicationVersion,omitempty" name:"ApplicationVersion"`
	// 架构，直接显示Arch表的值

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	// 构建的时间戳 Y%m%d-%H%M%S

	BuildTime *string `json:"BuildTime,omitempty" name:"BuildTime"`
	// 构建应用包的代码分支

	CodeBranch *string `json:"CodeBranch,omitempty" name:"CodeBranch"`
	// 构建应用包的commit的提交者

	CodeCommitter *string `json:"CodeCommitter,omitempty" name:"CodeCommitter"`
	// 代码仓库的地址

	CodeRepo *string `json:"CodeRepo,omitempty" name:"CodeRepo"`
	// 打包应用包的代码仓库的commit完成hash

	CommitID *string `json:"CommitID,omitempty" name:"CommitID"`
	// 包的依赖信息，为后续公有云的组件版本依赖提供基础

	DependencyInfo *string `json:"DependencyInfo,omitempty" name:"DependencyInfo"`
	// ( primary key ) 自增主键

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 1代表最新· 2代表不是最新，最新只适用于ApplicationPackageType release和dev可能分别都有一个1

	LatestFlag *int64 `json:"LatestFlag,omitempty" name:"LatestFlag"`
	// 元数据，json格式

	MetaData *string `json:"MetaData,omitempty" name:"MetaData"`
	// 触发流水线的人

	PipelineTriggerUser *string `json:"PipelineTriggerUser,omitempty" name:"PipelineTriggerUser"`
	// 应用包的状态 1启用 0废弃，其他状态后续扩展

	State *int64 `json:"State,omitempty" name:"State"`
	// uuid键

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 渲染出来的values文本, 长度需要确认

	ValuesData *string `json:"ValuesData,omitempty" name:"ValuesData"`
}

type ImportProductInstanceOperationSheetRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点ID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// SheetID 变更单ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// SheetName 变更单名称

	SheetName *string `json:"SheetName,omitempty" name:"SheetName"`
	// Description 运维单描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// ApplicationUUIDList 运维单关联应用UUID列表

	ApplicationUUIDList []*string `json:"ApplicationUUIDList,omitempty" name:"ApplicationUUIDList"`
	// Operator 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// Status 变更单状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ImportProductInstanceOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportProductInstanceOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerDiff struct {
	// UUID 服务器ID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// Diff 修改数据

	Diff *string `json:"Diff,omitempty" name:"Diff"`
}

type CopySiteRequest struct {
	*tchttp.BaseRequest

	// CopySiteUUID 复制的云 UUID

	CopySiteUUID *string `json:"CopySiteUUID,omitempty" name:"CopySiteUUID"`
	// SiteName 云名称

	SiteName *string `json:"SiteName,omitempty" name:"SiteName"`
	// Description 云描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// ConnectUrl 云连接信息

	ConnectUrl *string `json:"ConnectUrl,omitempty" name:"ConnectUrl"`
}

func (r *CopySiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopySiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployPlanProductRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点 UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// RegionUUID 地域 UUID

	RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
	// ProductInstanceID 产品实例 ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
}

func (r *DeployPlanProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeployPlanProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportStandardMaterialRequest struct {
	*tchttp.BaseRequest

	// MaterialUUID 标准物料 UUID

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
	// MaterialURL 标准物料下载链接

	MaterialURL *string `json:"MaterialURL,omitempty" name:"MaterialURL"`
	// Operator 操作者

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *ImportStandardMaterialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportStandardMaterialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlansResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// PlanList 规划版本信息列表

		PlanList []*PlanBriefInfo `json:"PlanList,omitempty" name:"PlanList"`
		// TotalCount 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPlansResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppPlanRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// CloudID 云UUID

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// ApplicationInstanceID 应用实例ID

	ApplicationInstanceID *string `json:"ApplicationInstanceID,omitempty" name:"ApplicationInstanceID"`
	// ApplicationInstanceUUID 应用实例UUID

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
	// PlanID 规划ID

	PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
	// PlanState 规划状态

	PlanState *string `json:"PlanState,omitempty" name:"PlanState"`
	// Tag 规划快照

	Tag *int64 `json:"Tag,omitempty" name:"Tag"`
}

func (r *DescribeAppPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAtomicRequest struct {
	*tchttp.BaseRequest

	// Atomic 原子操作Json结构体

	Atomic *JsonAtomic `json:"Atomic,omitempty" name:"Atomic"`
}

func (r *CreateAtomicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAtomicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductOperationTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductOperationTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductOperationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckMaterialRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// SolutionVersionUUID 解决方案版本UUID

	SolutionVersionUUID *string `json:"SolutionVersionUUID,omitempty" name:"SolutionVersionUUID"`
}

func (r *CheckMaterialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckMaterialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParseProductTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ApplicationNameList 应用名称列表

		ApplicationNameList []*string `json:"ApplicationNameList,omitempty" name:"ApplicationNameList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ParseProductTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ParseProductTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationArtifact struct {
	// 使用arch表里的值

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	// ( primary key ) 自增主键

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 制品导入状态，1导入，0未导入

	ImportState *int64 `json:"ImportState,omitempty" name:"ImportState"`
	// 关联application_package的UUID

	PackageUUID *string `json:"PackageUUID,omitempty" name:"PackageUUID"`
	// 存储制品的仓库地址，可能是docker registry或者包存储

	Repo *string `json:"Repo,omitempty" name:"Repo"`
	// 制品物料的sha256值

	Sha256 *string `json:"Sha256,omitempty" name:"Sha256"`
	// 制品物料的大小

	Size *int64 `json:"Size,omitempty" name:"Size"`
	// 制品包的状态 1启用 0废弃，其他状态后续扩展

	State *int64 `json:"State,omitempty" name:"State"`
	// 制品类型： chart/image/rpm/deb/raw/ted

	Type *string `json:"Type,omitempty" name:"Type"`
	// 制品物料的存储地址

	URL *string `json:"URL,omitempty" name:"URL"`
	// uuid键

	UUID *string `json:"UUID,omitempty" name:"UUID"`
}

type LocationInfo struct {
	// RegionUUID region 唯一ID

	RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
	// RegionID regionID

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	// ZoneUUID zone 唯一ID

	ZoneUUID *string `json:"ZoneUUID,omitempty" name:"ZoneUUID"`
}

type ListProductsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ProductInstances 产品实例列表

		ProductInstances []*ProductInstance `json:"ProductInstances,omitempty" name:"ProductInstances"`
		// TotalCount 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// CategorySet 标签列表

		CategorySet []*ProductCategoryCount `json:"CategorySet,omitempty" name:"CategorySet"`
		// StatusSet 状态列表

		StatusSet []*ProductStatusCount `json:"StatusSet,omitempty" name:"StatusSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductInstanceTreeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ProductCode 产品英文名

		ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
		// ProductName 产品中文名

		ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
		// ProductVersionName 产品版本名称

		ProductVersionName *string `json:"ProductVersionName,omitempty" name:"ProductVersionName"`
		// ProductInstanceUUID 产品实例 UUID

		ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
		// ProductInstanceID 产品实例 ID

		ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
		//

		ProductInstanceName *string `json:"ProductInstanceName,omitempty" name:"ProductInstanceName"`
		//

		SubsystemInstanceSet []*InstanceTreeSubsystem `json:"SubsystemInstanceSet,omitempty" name:"SubsystemInstanceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductInstanceTreeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductInstanceTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckActiveSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		IsActive *bool `json:"IsActive,omitempty" name:"IsActive"`
		//

		ActiveSite *ActiveSiteInfo `json:"ActiveSite,omitempty" name:"ActiveSite"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckActiveSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckActiveSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		ClusterList []*TcsCluster `json:"ClusterList,omitempty" name:"ClusterList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	//

	Key *string `json:"Key,omitempty" name:"Key"`
	//

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	//

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type ResumeJobRequest struct {
	*tchttp.BaseRequest

	//

	DagJobUUID *string `json:"DagJobUUID,omitempty" name:"DagJobUUID"`
	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
}

func (r *ResumeJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResumeJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductEnum struct {
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 取值

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type SolutionTemplate struct {
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 删除时间

	DeletedAt *int64 `json:"DeletedAt,omitempty" name:"DeletedAt"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// ( primary key ) 主键

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 描述运维模板是否是生效版本，0的时候不生效，1的时候生效

	IsActive *bool `json:"IsActive,omitempty" name:"IsActive"`
	// 描述运维模板是否是公共版本，0的时候是自定义，1的时候是公共版本

	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`
	// 修改者

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 运维模版名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作者

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 描述模版使用场景：一键部署，变更，大版本升级，扩容

	Scene *string `json:"Scene,omitempty" name:"Scene"`
	// 解决方案Code

	SolutionCode *string `json:"SolutionCode,omitempty" name:"SolutionCode"`
	// 解决方案名

	SolutionName *string `json:"SolutionName,omitempty" name:"SolutionName"`
	// 解决方案版本

	SolutionVersion *string `json:"SolutionVersion,omitempty" name:"SolutionVersion"`
	// 模板内容-json格式

	Template *string `json:"Template,omitempty" name:"Template"`
	// 描述模版类别,默认取值为solution

	Type *string `json:"Type,omitempty" name:"Type"`
	// uuid

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 运维模版版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type VmSpec struct {
	//

	CPU *string `json:"CPU,omitempty" name:"CPU"`
	//

	Memory *string `json:"Memory,omitempty" name:"Memory"`
	//

	Storage *string `json:"Storage,omitempty" name:"Storage"`
}

type ListPlanMiddlewareInfosRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// Filter 筛选器，支持key:ProductInstanceUUID, ApplicationInstanceUUID, ServiceInstanceName
	// ServiceBindingName, SupportAppInstanceName, SupportAppName

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
	// Offset 偏移值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit 返回记录数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListPlanMiddlewareInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanMiddlewareInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductCommonEnumResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 枚举类型列表

		ProductEnumList []*ProductEnum `json:"ProductEnumList,omitempty" name:"ProductEnumList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductCommonEnumResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductCommonEnumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCommonOperationSheetRequest struct {
	*tchttp.BaseRequest

	// 父单的SheetID

	ParentSheetID *string `json:"ParentSheetID,omitempty" name:"ParentSheetID"`
	// 父单编排的分组名

	StageName *string `json:"StageName,omitempty" name:"StageName"`
	// 额外信息

	Annotations *string `json:"Annotations,omitempty" name:"Annotations"`
	// 云ID, 产品部署运维单依赖该字段查询

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// ( primary key ) 自增主键

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 变更影响

	Influence *string `json:"Influence,omitempty" name:"Influence"`
	// 是否进行展示

	IsDisplay *int64 `json:"IsDisplay,omitempty" name:"IsDisplay"`
	// 是否是预置变更单

	IsInternal *int64 `json:"IsInternal,omitempty" name:"IsInternal"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 附件信息

	OperationSheetAttachment []*OperationSheetAttachment `json:"OperationSheetAttachment,omitempty" name:"OperationSheetAttachment"`
	// 父子单关联关系，仅当此单为父单时有效

	OperationSheetRelation []*OperationSheetRelation `json:"OperationSheetRelation,omitempty" name:"OperationSheetRelation"`
	// 变更单具体的编排信息, 一个变更单默认只有一个编排模板，预留model_array用于扩展

	OperationSheetTemplate []*OperationSheetTemplate `json:"OperationSheetTemplate,omitempty" name:"OperationSheetTemplate"`
	// 待编排项

	OrchestratingItems *string `json:"OrchestratingItems,omitempty" name:"OrchestratingItems"`
	// 负责人

	Owners *string `json:"Owners,omitempty" name:"Owners"`
	// 产品实例ID, 产品部署运维单依赖该字段查询

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// 区域ID, 产品部署运维单依赖该字段查询

	RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
	// 风险项

	Risk *string `json:"Risk,omitempty" name:"Risk"`
	// 场景

	Scene *string `json:"Scene,omitempty" name:"Scene"`
	// 运维单ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// 运维单类型（common,product,site,productInstance）

	SheetKind *string `json:"SheetKind,omitempty" name:"SheetKind"`
	// 运维单名称

	SheetName *string `json:"SheetName,omitempty" name:"SheetName"`
	// 运维模板类型(deploy,update,upgrade)

	SheetType *string `json:"SheetType,omitempty" name:"SheetType"`
	// 运维单来源:bug单, 交付单, 导入，人工创建

	Source *string `json:"Source,omitempty" name:"Source"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// AnnotationsVar是Annotations的结构体化展示

	AnnotationsVar *AnnotationsVar `json:"AnnotationsVar,omitempty" name:"AnnotationsVar"`
	// OrchestratingItemsVar是OrchestratingItems的结构体化展示

	OrchestratingItemsVar *OrchestratingItemsVar `json:"OrchestratingItemsVar,omitempty" name:"OrchestratingItemsVar"`
	// Children 是子变更单基本信息

	Children []*CommonOperationSheet `json:"Children,omitempty" name:"Children"`
	// Templates 是变更单编排相关的信息

	Templates []*OperationSheetTemplate `json:"Templates,omitempty" name:"Templates"`
}

func (r *CreateCommonOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCommonOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductInstanceOperationSheetRequest struct {
	*tchttp.BaseRequest

	// SheetID 变更单ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
}

func (r *DescribeProductInstanceOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductInstanceOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Param struct {
	//

	Name *string `json:"Name,omitempty" name:"Name"`
	//

	Description *string `json:"Description,omitempty" name:"Description"`
	//

	Type *string `json:"Type,omitempty" name:"Type"`
}

type GenerateApplicationPipelineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Template *string `json:"Template,omitempty" name:"Template"`
		//

		Pipeline *string `json:"Pipeline,omitempty" name:"Pipeline"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenerateApplicationPipelineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateApplicationPipelineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLatestPackageUrlRequest struct {
	*tchttp.BaseRequest

	// PackageType 物料类型(plan,atomic,operationtemplate,ccdeclare,customize...)

	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`
}

func (r *GetLatestPackageUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLatestPackageUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterType struct {
	//

	Key *string `json:"Key,omitempty" name:"Key"`
	//

	Op *string `json:"Op,omitempty" name:"Op"`
	//

	SQL *string `json:"SQL,omitempty" name:"SQL"`
	//

	Value *string `json:"Value,omitempty" name:"Value"`
	//

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type ListSiteProductMaterialRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// ProductVersionUUID 产品版本UUID

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	// MaterialType 物料类型，install-安装包，fix-补丁包，upgrade-升级包

	MaterialType *string `json:"MaterialType,omitempty" name:"MaterialType"`
	// 筛选项

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
	// Limit 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Offset 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListSiteProductMaterialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSiteProductMaterialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSiteProductMaterialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Total *int64 `json:"Total,omitempty" name:"Total"`
		//

		MaterialList []*MaterialInfo `json:"MaterialList,omitempty" name:"MaterialList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSiteProductMaterialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSiteProductMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSiteProductInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		//

		ProductInstances []*SiteProductInstance `json:"ProductInstances,omitempty" name:"ProductInstances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSiteProductInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSiteProductInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportCommonOperationSheetRequest struct {
	*tchttp.BaseRequest

	//

	SheetIDs []*string `json:"SheetIDs,omitempty" name:"SheetIDs"`
}

func (r *ExportCommonOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportCommonOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductVersionsRequest struct {
	*tchttp.BaseRequest

	// 产品 UUID

	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`
}

func (r *ListProductVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartOPMControllerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartOPMControllerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartOPMControllerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAtomicToPipelineRequest struct {
	*tchttp.BaseRequest

	//

	AtomicList []*AtomicKey `json:"AtomicList,omitempty" name:"AtomicList"`
}

func (r *ListAtomicToPipelineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAtomicToPipelineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PlanModification struct {
	// Type 配置类型：全局配置、产品配置、应用配置

	Type *string `json:"Type,omitempty" name:"Type"`
	// ProductInstanceUUID 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// ApplicationInstanceUUID 应用实例UUID

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
	// ServerUUID 服务器UUID

	ServerUUID *string `json:"ServerUUID,omitempty" name:"ServerUUID"`
	// ConfigUpdate 配置更新

	ConfigUpdate *bool `json:"ConfigUpdate,omitempty" name:"ConfigUpdate"`
	// ProductInstanceName 产品实例名称

	ProductInstanceName *string `json:"ProductInstanceName,omitempty" name:"ProductInstanceName"`
	// ApplicationInstanceName 应用实例名称

	ApplicationInstanceName *string `json:"ApplicationInstanceName,omitempty" name:"ApplicationInstanceName"`
	// ServerID 服务器ID

	ServerID *string `json:"ServerID,omitempty" name:"ServerID"`
	// BeforeConfig 修改之前的配置值

	BeforeConfig *string `json:"BeforeConfig,omitempty" name:"BeforeConfig"`
	// CurrentConfig 修改后的配置值

	CurrentConfig *string `json:"CurrentConfig,omitempty" name:"CurrentConfig"`
}

type ProjectRegion struct {
	// 区域

	Area *string `json:"Area,omitempty" name:"Area"`
	// 中文名称

	ChineseName *string `json:"ChineseName,omitempty" name:"ChineseName"`
	// 城市

	City *string `json:"City,omitempty" name:"City"`
	// 云拓扑的ID

	CloudID *int64 `json:"CloudID,omitempty" name:"CloudID"`
	//

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// ( primary key )

	ID *int64 `json:"ID,omitempty" name:"ID"`
	//

	ProjectZone []*ProjectZone `json:"ProjectZone,omitempty" name:"ProjectZone"`
	// 区域ID

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	// 区域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 关联的真实地域

	RelatedRegionUUID *string `json:"RelatedRegionUUID,omitempty" name:"RelatedRegionUUID"`
	// 角色

	RoleType *string `json:"RoleType,omitempty" name:"RoleType"`
	// 类型, 0表示真实地域、1表示虚拟地域

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	//

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type TcsCluster struct {
	//

	Name *string `json:"Name,omitempty" name:"Name"`
	//

	IsMainCluster *bool `json:"IsMainCluster,omitempty" name:"IsMainCluster"`
}

type DescribeGlobalConfigRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// CloudID 云UUID

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// PlanID 规划ID

	PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
	// PlanState 规划状态

	PlanState *string `json:"PlanState,omitempty" name:"PlanState"`
	// Tag 规划快照

	Tag *int64 `json:"Tag,omitempty" name:"Tag"`
}

func (r *DescribeGlobalConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGlobalConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportApplicationPlanRequest struct {
	*tchttp.BaseRequest

	// 应用名称

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	// 制品包中应用的values

	ApplicationValues *string `json:"ApplicationValues,omitempty" name:"ApplicationValues"`
	// 指定mock数据时使用的application uuid(plan_application_instance.ApplicationID)

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	// 应用部署时外部传入的values，需要和制品中应用的values进行合并

	ParamValues *string `json:"ParamValues,omitempty" name:"ParamValues"`
	// SiteUUID is site uuid

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
}

func (r *ImportApplicationPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportApplicationPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSheetAppInstancesRequest struct {
	*tchttp.BaseRequest

	// SheetID 运维单ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// Usage 模板用途

	Usage *string `json:"Usage,omitempty" name:"Usage"`
}

func (r *ListSheetAppInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSheetAppInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IssueSolution struct {
	// IssueSolutionRel.SolutionVersion

	SolutionVersion *string `json:"SolutionVersion,omitempty" name:"SolutionVersion"`
	// IssueSolutionRel.Arch

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	// IssueSolutionRel.TapdUrl

	TapdUrl *string `json:"TapdUrl,omitempty" name:"TapdUrl"`
	// IssueSolutionRel.UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
}

type Template struct {
	// 应用运维模板适配的应用标准，已废弃，由标签替代，预计2023年5月后移除

	ApplicationStandard *string `json:"ApplicationStandard,omitempty" name:"ApplicationStandard"`
	// [Auto Create] branch field

	Branch *string `json:"Branch,omitempty" name:"Branch"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 删除时间

	DeletedAt *int64 `json:"DeletedAt,omitempty" name:"DeletedAt"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// [Auto Create] etag field

	Etag *string `json:"Etag,omitempty" name:"Etag"`
	// 起始版本，已废弃

	FromProductVersions *string `json:"FromProductVersions,omitempty" name:"FromProductVersions"`
	// ( primary key ) 主键

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 描述运维模板是否是生效版本，0的时候不生效，1的时候生效

	IsActive *bool `json:"IsActive,omitempty" name:"IsActive"`
	// 描述运维模板是否是公共版本，0的时候是自定义，1的时候是公共版本

	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`
	// 模板类型 取值为 deploy update upgrade

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 修改者

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 运维模版名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作者

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 产品版本的 UUID，兼容历史数据，预计2023年5月后移除

	ProductVersionID *string `json:"ProductVersionID,omitempty" name:"ProductVersionID"`
	// 产品版本的 UUID， 已废弃, 由标签替代，预计2023年5月后移除

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	// 局点的UUID，已废弃

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// 来源取值predefine or selfdefine，已废弃，由is_public替代，预计2023年5月后移除

	Source *string `json:"Source,omitempty" name:"Source"`
	// [Auto Create] tag field

	Tag *string `json:"Tag,omitempty" name:"Tag"`
	// 模板内容

	Template *string `json:"Template,omitempty" name:"Template"`
	// 描述模版类别,取值为product application等

	Type *string `json:"Type,omitempty" name:"Type"`
	// uuid

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 模版用途，只有当Kind为update时有效，update表示变更用途，rollback表示回滚用途

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// 运维模版版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// ProductCode 运维模板关联的产品版本名称

	ProductVersionName *string `json:"ProductVersionName,omitempty" name:"ProductVersionName"`
	// ProductName 运维模板关联的产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// ProductUUID 运维模板关联的产品UUID

	ProductUUID *string `json:"ProductUUID,omitempty" name:"ProductUUID"`
	// IsProductApplicationChanged 运维模板关联的产品UUID

	IsProductApplicationChanged *bool `json:"IsProductApplicationChanged,omitempty" name:"IsProductApplicationChanged"`
	// MissApplicationNameList 产品运维模版中缺少的对应的产品版本的应用列表，方便在列表界面做应用列表展示

	MissApplicationNameList []*string `json:"MissApplicationNameList,omitempty" name:"MissApplicationNameList"`
	// Labels 运维模板相关标签信息

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
}

type ListProductInstanceAppsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		SubsystemSet []*SubsystemInfo `json:"SubsystemSet,omitempty" name:"SubsystemSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductInstanceAppsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductInstanceAppsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductInstanceSubsystemsRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
}

func (r *ListProductInstanceSubsystemsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductInstanceSubsystemsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductPlanRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// CloudID 云UUID

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// ProductID 产品ID

	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`
	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// ProductInstanceUUID 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// Location 产品位置信息

	Location *LocationInfo `json:"Location,omitempty" name:"Location"`
	// ProductCode 产品名称

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

func (r *ListProductPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ID 应用自增ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// PlanID 规划ID

		PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
		// SiteUUID 局点UUID

		SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
		// ApplicationInstanceID 应用实例ID

		ApplicationInstanceID *string `json:"ApplicationInstanceID,omitempty" name:"ApplicationInstanceID"`
		// ApplicationInstanceName 应用实例名称

		ApplicationInstanceName *string `json:"ApplicationInstanceName,omitempty" name:"ApplicationInstanceName"`
		// ApplicationInstanceUUID 应用实例UUID

		ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
		// ApplicationName 应用名称

		ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
		// ApplicationID 应用ID

		ApplicationID *string `json:"ApplicationID,omitempty" name:"ApplicationID"`
		// ProductInstanceUUID 产品实例ID

		ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
		// SubsystemInstanceUUID 子系统实例ID

		SubsystemInstanceUUID *string `json:"SubsystemInstanceUUID,omitempty" name:"SubsystemInstanceUUID"`
		// Location 示例部署位置

		Location *string `json:"Location,omitempty" name:"Location"`
		// Config 应用配置:values.yaml

		Config *string `json:"Config,omitempty" name:"Config"`
		// IntegratedConfig 集成后的应用配置:values.yaml

		IntegratedConfig *string `json:"IntegratedConfig,omitempty" name:"IntegratedConfig"`
		// IntegratedServerConfig 集成服务器信息后的应用配置:values.yaml

		IntegratedServerConfig *string `json:"IntegratedServerConfig,omitempty" name:"IntegratedServerConfig"`
		// ProductConfig 产品配置信息

		ProductConfig *string `json:"ProductConfig,omitempty" name:"ProductConfig"`
		// AppComponentSet 应用模块信息列表

		AppComponentSet []*Component `json:"AppComponentSet,omitempty" name:"AppComponentSet"`
		// AppServerSet 应用服务器信息列表

		AppServerSet []*Server `json:"AppServerSet,omitempty" name:"AppServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLatestMaterialPathRequest struct {
	*tchttp.BaseRequest

	// 物料所在服务器

	MaterialServer *string `json:"MaterialServer,omitempty" name:"MaterialServer"`
	// 机器用户

	MaterialSSHUser *string `json:"MaterialSSHUser,omitempty" name:"MaterialSSHUser"`
	// 机器密码

	MaterialSSHPassword *string `json:"MaterialSSHPassword,omitempty" name:"MaterialSSHPassword"`
	// 机器端口

	MaterialSSHPort *int64 `json:"MaterialSSHPort,omitempty" name:"MaterialSSHPort"`
}

func (r *GetLatestMaterialPathRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLatestMaterialPathRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanServerInfosRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// ProductInstanceUUID 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// MachineType 服务器类型

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// Filter 过滤条件:支持ProductInstanceUUID,ApplicationInstanceUUID,ComponentName,
	// ZoneID, serverType, OsName, ProductCode

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
	// Limit 数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Offset 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Location 服务器位置信息

	Location *ServerLocation `json:"Location,omitempty" name:"Location"`
}

func (r *ListPlanServerInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanServerInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VirtualServerInfo struct {
	// UUID 服务器UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// ServerID 服务器ID

	ServerID *string `json:"ServerID,omitempty" name:"ServerID"`
	// IP 服务器IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// HostUUID 宿主机UUID

	HostUUID *string `json:"HostUUID,omitempty" name:"HostUUID"`
	// HostIP 服务器HostIP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// Type 服务器类型：物理机/虚拟机

	Type *string `json:"Type,omitempty" name:"Type"`
	// Location 位置

	Location *string `json:"Location,omitempty" name:"Location"`
	// DiskInfo 磁盘信息

	DiskInfo *string `json:"DiskInfo,omitempty" name:"DiskInfo"`
	// OsInfo os信息

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// Metadata 服务器属性

	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`
	// ServerType 服务器型号

	ServerType *string `json:"ServerType,omitempty" name:"ServerType"`
	// ExtraAttr 服务器其他属性

	ExtraAttr *string `json:"ExtraAttr,omitempty" name:"ExtraAttr"`
	// CVMUUID 虚拟机在CVM的UUID

	CVMUUID *string `json:"CVMUUID,omitempty" name:"CVMUUID"`
	// OsName 操作系统名称

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// LogicArea 逻辑区域

	LogicArea *string `json:"LogicArea,omitempty" name:"LogicArea"`
	// ProductSet 部署产品列表

	ProductSet []*string `json:"ProductSet,omitempty" name:"ProductSet"`
	// AppComponentSet 服务器上的应用模块信息

	AppComponentSet []*AppComponentInfo `json:"AppComponentSet,omitempty" name:"AppComponentSet"`
	// VmType 虚拟机类型：kubeVM, undelayVM

	VmType *string `json:"VmType,omitempty" name:"VmType"`
	// Capacity 虚拟机规格

	Capacity *KubeVmCapacity `json:"Capacity,omitempty" name:"Capacity"`
	// NewHostIP 虚拟机IP

	NewHostIP *string `json:"NewHostIP,omitempty" name:"NewHostIP"`
}

type DescribeProductInstanceTreeRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点的 UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// ProductInstanceID 产品实例的 ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// ApplicationInstanceName 应用实例名称，如果存在应用名称，则只拿一个应用的信息，用于单应用部署场景

	ApplicationInstanceName *string `json:"ApplicationInstanceName,omitempty" name:"ApplicationInstanceName"`
	// RegionID 地域的 UUID

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
}

func (r *DescribeProductInstanceTreeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductInstanceTreeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetInstantiateAnnotationForOperationTemplateRequest struct {
	*tchttp.BaseRequest

	//

	Template *string `json:"Template,omitempty" name:"Template"`
	//

	Usage *string `json:"Usage,omitempty" name:"Usage"`
}

func (r *SetInstantiateAnnotationForOperationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetInstantiateAnnotationForOperationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAtomicLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		LabelList []*KVPair `json:"LabelList,omitempty" name:"LabelList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAtomicLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAtomicLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用包UUID

		ApplicationPackageID *string `json:"ApplicationPackageID,omitempty" name:"ApplicationPackageID"`
		// 内容

		Content *string `json:"Content,omitempty" name:"Content"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApplicationYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MaterialSync struct {
	// 产品版本UUID

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	// 解决方案UUID

	SolutionVersionUUID *string `json:"SolutionVersionUUID,omitempty" name:"SolutionVersionUUID"`
	// 物料UUID

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
}

type OperationSheetAttachment struct {
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 附件描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 附件格式（pdf, yaml, xlsx, doc, docx）

	Format *string `json:"Format,omitempty" name:"Format"`
	// ( primary key ) 自增主键

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 附件类型（TestReport,Common）

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 变更包里的相对路径

	LocalFilePath *string `json:"LocalFilePath,omitempty" name:"LocalFilePath"`
	// 附件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 存储的相对路径

	RemotePath *string `json:"RemotePath,omitempty" name:"RemotePath"`
	// 变更单ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// 附件大小

	Size *int64 `json:"Size,omitempty" name:"Size"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type ListProductInstanceSubsystemsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品版本 ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// 产品版本 UUID

		UUID *string `json:"UUID,omitempty" name:"UUID"`
		// 产品版本数据 Tag

		Tag *string `json:"Tag,omitempty" name:"Tag"`
		// 产品版本名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 产品版本配置

		Config *string `json:"Config,omitempty" name:"Config"`
		// 子系统列表

		SubsystemList []*SubsystemDetail `json:"SubsystemList,omitempty" name:"SubsystemList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductInstanceSubsystemsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductInstanceSubsystemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckOperationTemplateRequest struct {
	*tchttp.BaseRequest

	// TemplateInfo 需要检查的运维模版信息

	TemplateInfo *Template `json:"TemplateInfo,omitempty" name:"TemplateInfo"`
}

func (r *CheckOperationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckOperationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductPlanRequest struct {
	*tchttp.BaseRequest

	// SiteUUID 局点UUID

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// CloudID 云UUID

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// ProductInstanceID 产品实例ID

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// ProductInstanceUUID 产品实例UUID

	ProductInstanceUUID *string `json:"ProductInstanceUUID,omitempty" name:"ProductInstanceUUID"`
	// Location 位置信息

	Location *LocationInfo `json:"Location,omitempty" name:"Location"`
	// PlanID 规划ID

	PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
	// PlanState 规划状态

	PlanState *string `json:"PlanState,omitempty" name:"PlanState"`
	// Tag 规划快照

	Tag *int64 `json:"Tag,omitempty" name:"Tag"`
	// AppFilters 应用过滤器

	AppFilters []*FilterType `json:"AppFilters,omitempty" name:"AppFilters"`
	// ExcludeApp 排除应用信息，当为true则不返回任何应用信息

	ExcludeApp *bool `json:"ExcludeApp,omitempty" name:"ExcludeApp"`
}

func (r *DescribeProductPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRuntimeSecretRequest struct {
	*tchttp.BaseRequest

	// RegionName Region名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// Namespace Secret命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// ServiceBindingName sb名称

	ServiceBindingName *string `json:"ServiceBindingName,omitempty" name:"ServiceBindingName"`
}

func (r *GetRuntimeSecretRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRuntimeSecretRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductOperationTaskRequest struct {
	*tchttp.BaseRequest

	//

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	//

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	//

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	//

	ProductVersionID *string `json:"ProductVersionID,omitempty" name:"ProductVersionID"`
}

func (r *CreateProductOperationTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductOperationTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSiteOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSiteOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSiteOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SystemConfigs 系统配置

		SystemConfigs []*SystemConfig `json:"SystemConfigs,omitempty" name:"SystemConfigs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSystemConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceTreeApplication struct {
	//

	ApplicationUUID *string `json:"ApplicationUUID,omitempty" name:"ApplicationUUID"`
	//

	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
	//

	ApplicationInstanceUUID *string `json:"ApplicationInstanceUUID,omitempty" name:"ApplicationInstanceUUID"`
	//

	ApplicationInstanceName *string `json:"ApplicationInstanceName,omitempty" name:"ApplicationInstanceName"`
	//

	ApplicationInstanceID *string `json:"ApplicationInstanceID,omitempty" name:"ApplicationInstanceID"`
	// ApplicationVersion 应用的目标版本

	ApplicationVersion *string `json:"ApplicationVersion,omitempty" name:"ApplicationVersion"`
	// CurrentApplicationVersion 应用当前运行的版本

	CurrentApplicationVersion *string `json:"CurrentApplicationVersion,omitempty" name:"CurrentApplicationVersion"`
}

type DescribeGlobalConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ID 自增ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// UUID 云唯一ID

		UUID *string `json:"UUID,omitempty" name:"UUID"`
		// CloudName 云名称

		CloudName *string `json:"CloudName,omitempty" name:"CloudName"`
		// PlanID 关联的规划ID

		PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
		// PlanUUID 关联的规划的UUID

		PlanUUID *string `json:"PlanUUID,omitempty" name:"PlanUUID"`
		// PlanVersion 关联的规划版本

		PlanVersion *string `json:"PlanVersion,omitempty" name:"PlanVersion"`
		// PlanName 关联的规划名称

		PlanName *string `json:"PlanName,omitempty" name:"PlanName"`
		// SolutionVersionID 使用的解决方案版本ID

		SolutionVersionID *string `json:"SolutionVersionID,omitempty" name:"SolutionVersionID"`
		// SolutionVersionName 使用的解决方案版本名称

		SolutionVersionName *string `json:"SolutionVersionName,omitempty" name:"SolutionVersionName"`
		// SolutionName 解决方案名称

		SolutionName *string `json:"SolutionName,omitempty" name:"SolutionName"`
		// SolutionCode 解决方案Code

		SolutionCode *string `json:"SolutionCode,omitempty" name:"SolutionCode"`
		// SolutionVersionDataTag 解决方案版本的数据tag

		SolutionVersionDataTag *string `json:"SolutionVersionDataTag,omitempty" name:"SolutionVersionDataTag"`
		// Config 全局配置信息

		Config *string `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGlobalConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGlobalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPackageManagerHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		PackageManagerHost *string `json:"PackageManagerHost,omitempty" name:"PackageManagerHost"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPackageManagerHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPackageManagerHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppComponentTreeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品实例列表

		ProductInstanceList []*InstanceTreeProduct `json:"ProductInstanceList,omitempty" name:"ProductInstanceList"`
		// 解决方案名称

		SolutionName *string `json:"SolutionName,omitempty" name:"SolutionName"`
		// 解决方案版本名称

		SolutionVersionName *string `json:"SolutionVersionName,omitempty" name:"SolutionVersionName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppComponentTreeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppComponentTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSiteMaterialsRequest struct {
	*tchttp.BaseRequest

	// 局点uuid

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
	// 筛选信息

	Filter []*FilterType `json:"Filter,omitempty" name:"Filter"`
	//

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListSiteMaterialsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSiteMaterialsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateJobRequest struct {
	*tchttp.BaseRequest

	//

	DagJobUUID *string `json:"DagJobUUID,omitempty" name:"DagJobUUID"`
	//

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
}

func (r *TerminateJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAtomicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Total 符合条件的原子操作数据和，用于分页

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// AtomicList 分页之后原子操作基础信息列表

		AtomicList []*JsonAtomic `json:"AtomicList,omitempty" name:"AtomicList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAtomicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAtomicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelateCommonOperationSheetRequest struct {
	*tchttp.BaseRequest

	// 父单SheetID

	ParentSheetID *string `json:"ParentSheetID,omitempty" name:"ParentSheetID"`
	// 子单SheetID

	ChildSheetID *string `json:"ChildSheetID,omitempty" name:"ChildSheetID"`
	// true是关联，false是解除关联

	IsRelate *bool `json:"IsRelate,omitempty" name:"IsRelate"`
}

func (r *RelateCommonOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelateCommonOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEnvResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// EnvID 环境 ID

		EnvID *int64 `json:"EnvID,omitempty" name:"EnvID"`
		// EnvUUID 环境 UUID

		EnvUUID *string `json:"EnvUUID,omitempty" name:"EnvUUID"`
		// EnvName 环境名称

		EnvName *string `json:"EnvName,omitempty" name:"EnvName"`
		// PlanID 规划 ID

		PlanID *int64 `json:"PlanID,omitempty" name:"PlanID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEnvResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSiteRequest struct {
	*tchttp.BaseRequest

	// SiteName 局点名称

	SiteName *string `json:"SiteName,omitempty" name:"SiteName"`
	// Description 云描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// ConnectUrl 云连接信息

	ConnectUrl *string `json:"ConnectUrl,omitempty" name:"ConnectUrl"`
}

func (r *CreateSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportOperationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		//

		AtomicKeyList []*AtomicKey `json:"AtomicKeyList,omitempty" name:"AtomicKeyList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportOperationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportOperationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckActiveSiteRequest struct {
	*tchttp.BaseRequest

	//

	SiteUUID *string `json:"SiteUUID,omitempty" name:"SiteUUID"`
}

func (r *CheckActiveSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckActiveSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialOperationSheetRequest struct {
	*tchttp.BaseRequest

	// MaterialUUID 物料UUID

	MaterialUUID *string `json:"MaterialUUID,omitempty" name:"MaterialUUID"`
	// ProductVersionUUID 产品版本uuid

	ProductVersionUUID *string `json:"ProductVersionUUID,omitempty" name:"ProductVersionUUID"`
	//

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	//

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListMaterialOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonOperationSheet struct {
	// 额外信息

	Annotations *string `json:"Annotations,omitempty" name:"Annotations"`
	// 云ID, 产品部署运维单依赖该字段查询

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// ( primary key ) 自增主键

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 变更影响

	Influence *string `json:"Influence,omitempty" name:"Influence"`
	// 是否进行展示

	IsDisplay *int64 `json:"IsDisplay,omitempty" name:"IsDisplay"`
	// 是否是预置变更单

	IsInternal *int64 `json:"IsInternal,omitempty" name:"IsInternal"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 附件信息

	OperationSheetAttachment []*OperationSheetAttachment `json:"OperationSheetAttachment,omitempty" name:"OperationSheetAttachment"`
	// 父子单关联关系，仅当此单为父单时有效

	OperationSheetRelation []*OperationSheetRelation `json:"OperationSheetRelation,omitempty" name:"OperationSheetRelation"`
	// 变更单具体的编排信息, 一个变更单默认只有一个编排模板，预留model_array用于扩展

	OperationSheetTemplate []*OperationSheetTemplate `json:"OperationSheetTemplate,omitempty" name:"OperationSheetTemplate"`
	// 待编排项

	OrchestratingItems *string `json:"OrchestratingItems,omitempty" name:"OrchestratingItems"`
	// 负责人

	Owners *string `json:"Owners,omitempty" name:"Owners"`
	// 产品实例ID, 产品部署运维单依赖该字段查询

	ProductInstanceID *string `json:"ProductInstanceID,omitempty" name:"ProductInstanceID"`
	// 区域ID, 产品部署运维单依赖该字段查询

	RegionUUID *string `json:"RegionUUID,omitempty" name:"RegionUUID"`
	// 风险项

	Risk *string `json:"Risk,omitempty" name:"Risk"`
	// 场景

	Scene *string `json:"Scene,omitempty" name:"Scene"`
	// 运维单ID

	SheetID *string `json:"SheetID,omitempty" name:"SheetID"`
	// 运维单类型（common,product,site,productInstance）

	SheetKind *string `json:"SheetKind,omitempty" name:"SheetKind"`
	// 运维单名称

	SheetName *string `json:"SheetName,omitempty" name:"SheetName"`
	// 运维模板类型(deploy,update,upgrade)

	SheetType *string `json:"SheetType,omitempty" name:"SheetType"`
	// 运维单来源:bug单, 交付单, 导入，人工创建

	Source *string `json:"Source,omitempty" name:"Source"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type Component struct {
	// ComponentName 模块名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// Type 模块类型：noncontainer,deployment等

	Type *string `json:"Type,omitempty" name:"Type"`
	// NodeSet 模块部署信息

	NodeSet []*NodeInfo `json:"NodeSet,omitempty" name:"NodeSet"`
}

type GenerateApplicationPipelineRequest struct {
	*tchttp.BaseRequest

	//

	ApplicationInfo *ApplicationInfoData `json:"ApplicationInfo,omitempty" name:"ApplicationInfo"`
	//

	EnvInfo *EnvInfoData `json:"EnvInfo,omitempty" name:"EnvInfo"`
}

func (r *GenerateApplicationPipelineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateApplicationPipelineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstantiateCommonOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Succeeded []*string `json:"Succeeded,omitempty" name:"Succeeded"`
		//

		Failure []*string `json:"Failure,omitempty" name:"Failure"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InstantiateCommonOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstantiateCommonOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportOperationTemplateRequest struct {
	*tchttp.BaseRequest

	//

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
	//

	LocalFilePath *string `json:"LocalFilePath,omitempty" name:"LocalFilePath"`
	//

	Package *string `json:"Package,omitempty" name:"Package"`
}

func (r *ImportOperationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportOperationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResumeJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResumeJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
