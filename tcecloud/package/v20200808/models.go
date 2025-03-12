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

package v20200808

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DeletePackageParamPackage struct {
	// 包UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 包名称

	Name []*string `json:"Name,omitempty" name:"Name"`
}

type DeletePackageInstancesRequest struct {
	*tchttp.BaseRequest

	// 软件包实例信息列表

	Instances []*DeletePackageInstancesParamInstance `json:"Instances,omitempty" name:"Instances"`
}

func (r *DeletePackageInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePackageInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCSPUploadURLRequest struct {
	*tchttp.BaseRequest

	// 文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件版本

	FileVersion *string `json:"FileVersion,omitempty" name:"FileVersion"`
	// 包名

	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`
	// 包CPU架构，可选项：all、x86、arm，不填默认为all

	PkgArch *string `json:"PkgArch,omitempty" name:"PkgArch"`
}

func (r *CreateCSPUploadURLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCSPUploadURLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePackageInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 包实例信息

		Instances []*DescribePackageInstancesResponseInstance `json:"Instances,omitempty" name:"Instances"`
		// 包实例个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePackageInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePackageInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SorterParams struct {
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒排

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
}

type DescribePackageVersionsResponseHealthCheck struct {
	// 端口信息

	Port []*DescribePackageVersionsResponsePort `json:"Port,omitempty" name:"Port"`
	// 进程信息

	Process []*DescribePackageVersionsResponseProcess `json:"Process,omitempty" name:"Process"`
	// 脚本信息

	Script []*string `json:"Script,omitempty" name:"Script"`
	// HTTP信息

	HTTP []*DescribePackageVersionsResponseHttp `json:"HTTP,omitempty" name:"HTTP"`
	// 模块信息

	Module []*string `json:"Module,omitempty" name:"Module"`
}

type DescribePackageVersionsResponseStorageInfo struct {
	// 存储空间大小

	Size *string `json:"Size,omitempty" name:"Size"`
	// 存储路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 存储驱动

	Driver *string `json:"Driver,omitempty" name:"Driver"`
}

type UpdatePackageParamUpdateData struct {
	// 包管理人员列表

	Manager []*string `json:"Manager,omitempty" name:"Manager"`
	// 包所属产品列表

	Product []*string `json:"Product,omitempty" name:"Product"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 包类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type RegisterPackageInstanceRequest struct {
	*tchttp.BaseRequest
}

func (r *RegisterPackageInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterPackageInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePackageVersionsResponseCheckRestart struct {
	// 检查间隔时间

	CheckInteval *string `json:"CheckInteval,omitempty" name:"CheckInteval"`
	// 数量限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Grace

	Grace *string `json:"Grace,omitempty" name:"Grace"`
	// 是否忽略警告

	IgnoreWarnings *bool `json:"IgnoreWarnings,omitempty" name:"IgnoreWarnings"`
	// 重启延迟

	RestartDelay *string `json:"RestartDelay,omitempty" name:"RestartDelay"`
}

type DescribePackageVersionsRequest struct {
	*tchttp.BaseRequest

	// 限制返回个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件信息

	Filters *DescribeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序

	Sorters []*SorterParams `json:"Sorters,omitempty" name:"Sorters"`
}

func (r *DescribePackageVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePackageVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePackageVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 包版本信息

		PackageVersions []*DescribePackageVersionsPackageVersion `json:"PackageVersions,omitempty" name:"PackageVersions"`
		// 版本数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePackageVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePackageVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePackagesParamPackage struct {
	// 软件包UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 软件包版本列表

	Versions []*string `json:"Versions,omitempty" name:"Versions"`
}

type DescribeFilter struct {
	// 字段名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 匹配类型，例如：term、fuzzy、range等，不区分大小写

	Type *string `json:"Type,omitempty" name:"Type"`
	// 字段值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeFilters struct {
	// 同SQL中的'and'

	Must []*DescribeFilter `json:"Must,omitempty" name:"Must"`
	// 同SQL中的'or'

	Should []*DescribeFilter `json:"Should,omitempty" name:"Should"`
	// 同SQL中的'!='

	MustNot []*DescribeFilter `json:"MustNot,omitempty" name:"MustNot"`
}

type CreatePackageVersionRequest struct {
	*tchttp.BaseRequest

	// 包版本信息实体

	PackageVersion *CreatePackageVersionParam `json:"PackageVersion,omitempty" name:"PackageVersion"`
}

func (r *CreatePackageVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePackageVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 影响行数

		RowsAffected *int64 `json:"RowsAffected,omitempty" name:"RowsAffected"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 影响行数

		RowsAffected *int64 `json:"RowsAffected,omitempty" name:"RowsAffected"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePackageInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 影响的数据行

		RowsAffected *int64 `json:"RowsAffected,omitempty" name:"RowsAffected"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePackageInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePackageInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePackageVersionsResponseScript struct {
	// 安装脚本路径

	Install *string `json:"Install,omitempty" name:"Install"`
	// 启动脚本路径

	Start *string `json:"Start,omitempty" name:"Start"`
	// 停止脚本路径

	Stop *string `json:"Stop,omitempty" name:"Stop"`
	// 卸载脚本路径

	Uninstall *string `json:"Uninstall,omitempty" name:"Uninstall"`
}

type CreateCSPUploadURLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 上传文件URL

		Url *string `json:"Url,omitempty" name:"Url"`
		// 保存至对象存储的Key

		Key *string `json:"Key,omitempty" name:"Key"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCSPUploadURLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCSPUploadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePackageParam struct {
	// 包类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 包名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 包创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 包管理者

	Manager []*string `json:"Manager,omitempty" name:"Manager"`
	// 包所属产品

	Product []*string `json:"Product,omitempty" name:"Product"`
	// 包描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreatePackageVersionParam struct {
	// UUID

	PkgUUID *string `json:"PkgUUID,omitempty" name:"PkgUUID"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 部署路径

	Destination *string `json:"Destination,omitempty" name:"Destination"`
	// 包启动用户

	ExecUser *string `json:"ExecUser,omitempty" name:"ExecUser"`
	// 软件控制脚本

	Script *DescribePackageVersionsResponseScript `json:"Script,omitempty" name:"Script"`
	// 软件健康检查配置

	HealthCheck *DescribePackageVersionsResponseHealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`
	// 软件包的存储信息

	StorageInfo *DescribePackageVersionsResponseStorageInfo `json:"StorageInfo,omitempty" name:"StorageInfo"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 包配置模版

	Template []*DescribePackageVersionsResponseTemplate `json:"Template,omitempty" name:"Template"`
	// 是否需要解析包

	NotParse *bool `json:"NotParse,omitempty" name:"NotParse"`
	// Md5值

	Hash *string `json:"Hash,omitempty" name:"Hash"`
	// CPU架构

	Arch *string `json:"Arch,omitempty" name:"Arch"`
}

type DeletedFilter struct {
	// 需要过滤的字段

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段的过滤值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeEventsResponseEvent struct {
	// 事件UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 事件类型

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 是否成功

	Success *bool `json:"Success,omitempty" name:"Success"`
	// 包名称

	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`
	// 包版本

	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`
	// 包类型

	PkgType *string `json:"PkgType,omitempty" name:"PkgType"`
	// 包CPU架构

	PkgArch *string `json:"PkgArch,omitempty" name:"PkgArch"`
	// 事件创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 发生时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
}

type DescribePackageVersionsResponsePort struct {
	// 端口号

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type PackageVersionDeleteQuery struct {
	// 包UUID

	PkgUUID *string `json:"PkgUUID,omitempty" name:"PkgUUID"`
	// 版本列表

	Version *string `json:"Version,omitempty" name:"Version"`
	// CPU架构

	Arch *string `json:"Arch,omitempty" name:"Arch"`
}

type CreateCSPDownloadURLRequest struct {
	*tchttp.BaseRequest

	// 文件的Key

	Key *string `json:"Key,omitempty" name:"Key"`
}

func (r *CreateCSPDownloadURLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCSPDownloadURLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePackageRequest struct {
	*tchttp.BaseRequest

	// 创建软件包信息列表

	Package *CreatePackageParam `json:"Package,omitempty" name:"Package"`
}

func (r *CreatePackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePackageRequest struct {
	*tchttp.BaseRequest

	// 包删除时定位参数

	Package *DeletePackageParamPackage `json:"Package,omitempty" name:"Package"`
	// 包删除时更新数据

	UpdateData *UpdatePackageParamUpdateData `json:"UpdateData,omitempty" name:"UpdateData"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
}

func (r *UpdatePackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePackageVersionsResponseProcess struct {
	// 进程名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 最小值

	Min *int64 `json:"Min,omitempty" name:"Min"`
	// 最大值

	Max *int64 `json:"Max,omitempty" name:"Max"`
}

type DescribePackageInstancesRequest struct {
	*tchttp.BaseRequest

	// 限制返回个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件信息

	Filters *DescribeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序

	Sorters []*SorterParams `json:"Sorters,omitempty" name:"Sorters"`
}

func (r *DescribePackageInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePackageInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePackagesRequest struct {
	*tchttp.BaseRequest

	// 限制返回个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件信息

	Filters *DescribeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序

	Sorters []*SorterParams `json:"Sorters,omitempty" name:"Sorters"`
}

func (r *DescribePackagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePackagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePackagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 软件包列表信息

		Packages []*DescribePackagesResponsePackage `json:"Packages,omitempty" name:"Packages"`
		// 包列表数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePackagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventsRequest struct {
	*tchttp.BaseRequest

	// 限制返回个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件信息

	Filters *DescribeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序

	Sorters []*SorterParams `json:"Sorters,omitempty" name:"Sorters"`
}

func (r *DescribeEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePackageInstancesParamInstance struct {
	// 包实例ID

	InstanceID *int64 `json:"InstanceID,omitempty" name:"InstanceID"`
	// 软件包UUID

	PkgUUID *string `json:"PkgUUID,omitempty" name:"PkgUUID"`
	// 软件包版本

	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`
	// 机器IP

	SvrIP *string `json:"SvrIP,omitempty" name:"SvrIP"`
}

type DescribePackagesResponsePackage struct {
	// 软件包UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 软件包类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 软件包名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 包创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 包管理者

	Manager []*string `json:"Manager,omitempty" name:"Manager"`
	// 包所属产品

	Product []*string `json:"Product,omitempty" name:"Product"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 包描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 修改者

	UpdatedBy *string `json:"UpdatedBy,omitempty" name:"UpdatedBy"`
	// 版本数量

	Versions *int64 `json:"Versions,omitempty" name:"Versions"`
}

type CreateCSPDownloadURLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 上传文件URL

		Url *string `json:"Url,omitempty" name:"Url"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCSPDownloadURLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCSPDownloadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePackageVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务UUID

		TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePackageVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePackageVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePackageVersionsPackageVersion struct {
	// 软件包UUID

	PkgUUID *string `json:"PkgUUID,omitempty" name:"PkgUUID"`
	// 软件包版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 目标地址

	Destination *string `json:"Destination,omitempty" name:"Destination"`
	// 执行用户

	ExecUser *string `json:"ExecUser,omitempty" name:"ExecUser"`
	// 脚本路径信息

	Script *DescribePackageVersionsResponseScript `json:"Script,omitempty" name:"Script"`
	// 健康检查信息

	HealthCheck *DescribePackageVersionsResponseHealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`
	// 存储信息

	StorageInfo *DescribePackageVersionsResponseStorageInfo `json:"StorageInfo,omitempty" name:"StorageInfo"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 包配置模版

	Template []*DescribePackageVersionsResponseTemplate `json:"Template,omitempty" name:"Template"`
	// 最后拉取时间

	PulledAt *string `json:"PulledAt,omitempty" name:"PulledAt"`
	// 实例数量

	Instances *int64 `json:"Instances,omitempty" name:"Instances"`
}

type DescribePackageVersionsResponseHttp struct {
	// 端口号

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// 请求头部

	Headers *string `json:"Headers,omitempty" name:"Headers"`
}

type RegisterPackageInstanceParamInstance struct {
	// 软件包UUID

	PkgUUID *string `json:"PkgUUID,omitempty" name:"PkgUUID"`
	// 包版本

	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`
	// 机器IP

	SvrIP *string `json:"SvrIP,omitempty" name:"SvrIP"`
	// 流程引擎任务ID

	JobID *string `json:"JobID,omitempty" name:"JobID"`
}

type CreatePackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 软件包UUID

		UUID *string `json:"UUID,omitempty" name:"UUID"`
		// 影响行数

		RowsAffected *int64 `json:"RowsAffected,omitempty" name:"RowsAffected"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePackageInstancesResponseInstance struct {
	// 包实例ID

	InstanceID *int64 `json:"InstanceID,omitempty" name:"InstanceID"`
	// 软件包UUID

	PkgUUID *string `json:"PkgUUID,omitempty" name:"PkgUUID"`
	// 包版本

	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`
	// 机器IP

	SvrIP *string `json:"SvrIP,omitempty" name:"SvrIP"`
	// 包状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 包安装时间

	InstalledAt *string `json:"InstalledAt,omitempty" name:"InstalledAt"`
	// 流程引擎任务ID

	JobID *string `json:"JobID,omitempty" name:"JobID"`
}

type DescribeEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件列表

		Events []*DescribeEventsResponseEvent `json:"Events,omitempty" name:"Events"`
		// 符合条件的事件总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePackageVersionsResponseTemplate struct {
	// 源文件

	Src *string `json:"Src,omitempty" name:"Src"`
	// 渲染后目标文件

	Dest *string `json:"Dest,omitempty" name:"Dest"`
}

type DeletePackageVersionsRequest struct {
	*tchttp.BaseRequest

	// 版本列表

	PackageVersions []*PackageVersionDeleteQuery `json:"PackageVersions,omitempty" name:"PackageVersions"`
	// 删除操作提交人

	DeletedBy *string `json:"DeletedBy,omitempty" name:"DeletedBy"`
}

func (r *DeletePackageVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePackageVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePackageVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePackageVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePackageVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterPackageInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 影响行数

		RowsAffected *int64 `json:"RowsAffected,omitempty" name:"RowsAffected"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RegisterPackageInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterPackageInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePackageRequest struct {
	*tchttp.BaseRequest

	// 包删除参数

	Package *DeletePackageParamPackage `json:"Package,omitempty" name:"Package"`
	// 删除操作者

	DeletedBy *string `json:"DeletedBy,omitempty" name:"DeletedBy"`
}

func (r *DeletePackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
