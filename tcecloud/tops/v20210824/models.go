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

package v20210824

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type InfrastoreProduct struct {
	// 产品名

	Product *string `json:"Product,omitempty" name:"Product"`
	// 子产品名

	TcsType *string `json:"TcsType,omitempty" name:"TcsType"`
}

type RunningEnv struct {
	// 运行环境

	Env []*string `json:"Env,omitempty" name:"Env"`
	// 筛选执行Node

	NodeSelector *string `json:"NodeSelector,omitempty" name:"NodeSelector"`
	// 筛选执行Pod

	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`
}

type DescribeProductComponentsRequest struct {
	*tchttp.BaseRequest

	// 产品UUID, 如tcs

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤器, Name支持: ComponentType/ComponentArch/DeployLevel/CurrentVersion/Location/Healthy

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeProductComponentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductComponentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationInstructionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 变更单列表

		OperationInstructions []*OperationInstruction `json:"OperationInstructions,omitempty" name:"OperationInstructions"`
		// 符合条件的组件列表, 用于筛选

		AssociatedComponentsSet []*AssociatedComponentsSet `json:"AssociatedComponentsSet,omitempty" name:"AssociatedComponentsSet"`
		// 符合条件的大版本列表, 用于搜索

		AssociatedVersionSet []*string `json:"AssociatedVersionSet,omitempty" name:"AssociatedVersionSet"`
		// 列出大版本升级变更单时, 此字段表示云产品列表的值的集合

		ProductSet []*string `json:"ProductSet,omitempty" name:"ProductSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOperationInstructionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationInstructionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComponentPrepareTasks struct {
	// 组件名

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 步骤信息

	Steps []*PrepareTaskStep `json:"Steps,omitempty" name:"Steps"`
}

type TreeNodeMiddleware struct {
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 服务类型

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type OperationInstructionPlan struct {
	// 变更单名字

	OperationInstruction *string `json:"OperationInstruction,omitempty" name:"OperationInstruction"`
	// 组件总数

	TotalComponets *int64 `json:"TotalComponets,omitempty" name:"TotalComponets"`
	// 有规划的组件数量

	ComponentsWithPlan *int64 `json:"ComponentsWithPlan,omitempty" name:"ComponentsWithPlan"`
	// 所有实例的规划都已校正的组件数量

	MergeChecked *int64 `json:"MergeChecked,omitempty" name:"MergeChecked"`
}

type ProductTopologyNode struct {
	// 名称, 如某个region

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件列表

	Components []*TreeNode `json:"Components,omitempty" name:"Components"`
}

type ModifyToolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyToolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyToolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductComponentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件列表

		Components []*ComponentWithInstance `json:"Components,omitempty" name:"Components"`
		// 架构列表, 用于过滤

		ArchSet []*string `json:"ArchSet,omitempty" name:"ArchSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductComponentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Host struct {
	// Zone名字

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// ZoneId

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// VM状态

	VmStatus *string `json:"VmStatus,omitempty" name:"VmStatus"`
	// SN

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// region名字

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// Region ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 内网IP

	InnerIp *string `json:"InnerIp,omitempty" name:"InnerIp"`
	// 主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 架构, 例如x86_64

	CpuArchitecture *string `json:"CpuArchitecture,omitempty" name:"CpuArchitecture"`
}

type OperationInstructionWithUploadStatistics struct {
	// 变更单ID

	Name *string `json:"Name,omitempty" name:"Name"`
	// 变更单名字

	ShowName *string `json:"ShowName,omitempty" name:"ShowName"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 组件数量

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 所属云产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 上传进度

	UploadStatistics *UploadStatistics `json:"UploadStatistics,omitempty" name:"UploadStatistics"`
}

type BatchImportSheetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导入结果列表

		SheetResults []*SheetResult `json:"SheetResults,omitempty" name:"SheetResults"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchImportSheetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchImportSheetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateToolImportURLRequest struct {
	*tchttp.BaseRequest

	// 文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

func (r *CreateToolImportURLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateToolImportURLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductTopologyRequest struct {
	*tchttp.BaseRequest

	// 父节点名称

	ParentName *string `json:"ParentName,omitempty" name:"ParentName"`
	// 云产品名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeProductTopologyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductTopologyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParseTarFileDirTreeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件目录

		DirTree *DirTreeNode `json:"DirTree,omitempty" name:"DirTree"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ParseTarFileDirTreeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ParseTarFileDirTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowRequest struct {
	*tchttp.BaseRequest

	// Flow名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductBizTreeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云产品名

		Name *string `json:"Name,omitempty" name:"Name"`
		// 容器类型组件列表

		ContainerComponents []*TreeNodeContainerComponent `json:"ContainerComponents,omitempty" name:"ContainerComponents"`
		// 支持类型组件列表

		SupportComponents []*TreeNodeMiddleware `json:"SupportComponents,omitempty" name:"SupportComponents"`
		// 生产类型组件列表

		ProductComponents []*string `json:"ProductComponents,omitempty" name:"ProductComponents"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductBizTreeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductBizTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TreeNode struct {
	// 结点名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 状态, healthy/unhealthy

	Status *string `json:"Status,omitempty" name:"Status"`
	// 用于显示的结点名

	ShowName *string `json:"ShowName,omitempty" name:"ShowName"`
	// 组件数量

	ComponentsCount *int64 `json:"ComponentsCount,omitempty" name:"ComponentsCount"`
	// 告警数量

	AlertsCount *int64 `json:"AlertsCount,omitempty" name:"AlertsCount"`
}

type CreateOperationInstructionRequest struct {
	*tchttp.BaseRequest

	// 变更单YAML

	OperationInstruction *string `json:"OperationInstruction,omitempty" name:"OperationInstruction"`
	// 变更类型, upgrade(大版本升级)/bugifx

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *CreateOperationInstructionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationInstructionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExecutionRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 过滤器, 支持Node/StdOut/CodeStatus

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeExecutionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExecutionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationInstructionPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件规划

		Components []*ComponentWithInstance `json:"Components,omitempty" name:"Components"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationInstructionPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationInstructionPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductTopologyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各region拓扑结构

		Topology []*Topology `json:"Topology,omitempty" name:"Topology"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductTopologyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductTopologyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProjectRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 项目分组, 其中的变更单仅需传Name即可

	Partitions []*ProjectPartition `json:"Partitions,omitempty" name:"Partitions"`
	// 设置true表示一键上传所有变更单的物料

	UploadOperationInstructions *bool `json:"UploadOperationInstructions,omitempty" name:"UploadOperationInstructions"`
	// 单个单子内最多允许多少镜像同时上传. 设置为0表示全部并行

	MaxConcurrentUploadImage *int64 `json:"MaxConcurrentUploadImage,omitempty" name:"MaxConcurrentUploadImage"`
	// 物料机IP

	MaterialMachine *string `json:"MaterialMachine,omitempty" name:"MaterialMachine"`
}

func (r *UpdateProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateToolParseURLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 上传链接

		URL *string `json:"URL,omitempty" name:"URL"`
		// 上传ID

		UploadId *string `json:"UploadId,omitempty" name:"UploadId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateToolParseURLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateToolParseURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExecutionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 工具元数据

		Manifest *Manifest `json:"Manifest,omitempty" name:"Manifest"`
		// 执行命令

		Command *string `json:"Command,omitempty" name:"Command"`
		// 结点列表

		Nodes []*string `json:"Nodes,omitempty" name:"Nodes"`
		// 任务状态. Progressing/Pending/Failed/Error/Skipped/Completed/Stopped

		TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
		// 开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 操作人

		Operator *string `json:"Operator,omitempty" name:"Operator"`
		// 结点运行状态

		NodeRunningStatus []*NodeRunningStatus `json:"NodeRunningStatus,omitempty" name:"NodeRunningStatus"`
		// 符合筛选条件的结点总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回码列表, 供筛选

		CodeStatusSet []*string `json:"CodeStatusSet,omitempty" name:"CodeStatusSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExecutionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExecutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationStatsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 变更单总数

		OperationCount *int64 `json:"OperationCount,omitempty" name:"OperationCount"`
		// 成功的变更单数量

		SucceededOperationCount *int64 `json:"SucceededOperationCount,omitempty" name:"SucceededOperationCount"`
		// 失败的变更单数量

		FailedOperationCount *int64 `json:"FailedOperationCount,omitempty" name:"FailedOperationCount"`
		// 运行中的变更单数量

		RunningOperationCount *int64 `json:"RunningOperationCount,omitempty" name:"RunningOperationCount"`
		// 运维工具数量

		OperationToolsCount *int64 `json:"OperationToolsCount,omitempty" name:"OperationToolsCount"`
		// 按产品统计变更单数量

		ProductOperationCount []*KVPair `json:"ProductOperationCount,omitempty" name:"ProductOperationCount"`
		// 变更单历史趋势(按创建时间)

		OperationTrend []*OperationTrendPoint `json:"OperationTrend,omitempty" name:"OperationTrend"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationStatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationTaskLogRequest struct {
	*tchttp.BaseRequest

	// 变更单名字

	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`
	// 步骤名

	Step *string `json:"Step,omitempty" name:"Step"`
	// 节点名

	Node *string `json:"Node,omitempty" name:"Node"`
	// 变更单ID

	OperationInstruction *string `json:"OperationInstruction,omitempty" name:"OperationInstruction"`
}

func (r *GetOperationTaskLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationTaskLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportToolRequest struct {
	*tchttp.BaseRequest

	// 上传链接

	UploadId *string `json:"UploadId,omitempty" name:"UploadId"`
	// 上传文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

func (r *ImportToolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportToolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListFlowsRequest struct {
	*tchttp.BaseRequest

	// 指定标签筛选

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// 过滤器, 支持Name/Status/TargetVersion/SourceVersion

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListFlowsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListFlowsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileSource struct {
	// 来源于上传文件

	FromUpload []*string `json:"FromUpload,omitempty" name:"FromUpload"`
	// 来源于压缩包

	FromPackage []*string `json:"FromPackage,omitempty" name:"FromPackage"`
	// 来源于原工具包

	FromTool []*string `json:"FromTool,omitempty" name:"FromTool"`
}

type KVPair struct {
	// 监控指标

	Key *string `json:"Key,omitempty" name:"Key"`
	// 监控数据

	Value *string `json:"Value,omitempty" name:"Value"`
	// 数字类型

	ValueInt *int64 `json:"ValueInt,omitempty" name:"ValueInt"`
}

type StepLog struct {
	// 步骤名称

	StepName *string `json:"StepName,omitempty" name:"StepName"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 命令

	Command *string `json:"Command,omitempty" name:"Command"`
	// 标准输出

	StdOut *string `json:"StdOut,omitempty" name:"StdOut"`
	// 标准错误

	StdErr *string `json:"StdErr,omitempty" name:"StdErr"`
	// 命令状态码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 错误消息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type DescribeBizTreeRequest struct {
	*tchttp.BaseRequest

	// 指定云产品UUID查询业务树. 不填查询所有.

	BizUuid *string `json:"BizUuid,omitempty" name:"BizUuid"`
}

func (r *DescribeBizTreeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBizTreeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParseInstructionComponentsRequest struct {
	*tchttp.BaseRequest

	// 变更单YAML

	OperationInstruction *string `json:"OperationInstruction,omitempty" name:"OperationInstruction"`
	// 变更单名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ParseInstructionComponentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ParseInstructionComponentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeToolStatsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前运维工具数量

		OperationToolsCount *int64 `json:"OperationToolsCount,omitempty" name:"OperationToolsCount"`
		// 执行次数趋势

		ExecutionTrend []*ExecutionTrendPoint `json:"ExecutionTrend,omitempty" name:"ExecutionTrend"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeToolStatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeToolStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUploadPackageTaskRequest struct {
	*tchttp.BaseRequest

	// 变更单名字

	OperationInstruction *string `json:"OperationInstruction,omitempty" name:"OperationInstruction"`
}

func (r *DescribeUploadPackageTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUploadPackageTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListToolVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 工具版本列表

		ToolVersions []*ToolVersion `json:"ToolVersions,omitempty" name:"ToolVersions"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListToolVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListToolVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BizTree struct {
	// 云产品UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 云产品名字

	BizName *string `json:"BizName,omitempty" name:"BizName"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// Set列表

	SetList []*Set `json:"SetList,omitempty" name:"SetList"`
}

type DescribeToolDownloadURLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 临时下载链接

		URL *string `json:"URL,omitempty" name:"URL"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeToolDownloadURLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeToolDownloadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchImportSheetsRequest struct {
	*tchttp.BaseRequest

	// 压缩包CSP地址

	FileKey *string `json:"FileKey,omitempty" name:"FileKey"`
	// 项目ID

	ProjectID *string `json:"ProjectID,omitempty" name:"ProjectID"`
}

func (r *BatchImportSheetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchImportSheetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationInstructionDAGRequest struct {
	*tchttp.BaseRequest

	// 变更单ID

	Name *string `json:"Name,omitempty" name:"Name"`
	// 类型, update/rollback

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeOperationInstructionDAGRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationInstructionDAGRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeToolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 工具信息

		Manifest *Manifest `json:"Manifest,omitempty" name:"Manifest"`
		// 目录信息

		DirTree *DirTreeNode `json:"DirTree,omitempty" name:"DirTree"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeToolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeToolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ToolVersion struct {
	// 工具名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 工具版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 是否是当前版本

	Active *bool `json:"Active,omitempty" name:"Active"`
	// 更新者

	Updater *string `json:"Updater,omitempty" name:"Updater"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 工具ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

type CreateUploadPackageTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUploadPackageTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUploadPackageTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFindPackagePathTaskRequest struct {
	*tchttp.BaseRequest

	// 查找物料路径任务ID

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 变更单YAML

	OperationInstruction *string `json:"OperationInstruction,omitempty" name:"OperationInstruction"`
	// 组件版本信息

	Components []*ComponentInstruction `json:"Components,omitempty" name:"Components"`
}

func (r *DescribeFindPackagePathTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFindPackagePathTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFindPackagePathTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 错误信息

		ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFindPackagePathTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFindPackagePathTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportChangePackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 变更单ID

		ChangeId *string `json:"ChangeId,omitempty" name:"ChangeId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportChangePackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportChangePackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Partition struct {
	// 组件名称列表

	ComponentNames []*string `json:"ComponentNames,omitempty" name:"ComponentNames"`
}

type DescribeOperationInstructionLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志

		Logs []*StepLog `json:"Logs,omitempty" name:"Logs"`
		// 任务引擎UUID, 用于排查

		InstanceUUID *string `json:"InstanceUUID,omitempty" name:"InstanceUUID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationInstructionLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationInstructionLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductAlertsRequest struct {
	*tchttp.BaseRequest

	// 产品名称

	Product *string `json:"Product,omitempty" name:"Product"`
	// time range start

	StartsAt *string `json:"StartsAt,omitempty" name:"StartsAt"`
	// time range end

	EndsAt *string `json:"EndsAt,omitempty" name:"EndsAt"`
	// page offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// page limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 降序排列字段，逗号分隔多个字段

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 告警级别, 1/2/3/4

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 是否确认, "false"/"true"

	Confirmed *string `json:"Confirmed,omitempty" name:"Confirmed"`
	// 告警类型, metric/log, 分别表示指标/事件

	Class *string `json:"Class,omitempty" name:"Class"`
}

func (r *ListProductAlertsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductAlertsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetFlowStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetFlowStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetFlowStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeToolStatsRequest struct {
	*tchttp.BaseRequest

	// 查询区间开始时间, 如"2021-10-10T16:00:00"

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询区间结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeToolStatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeToolStatsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryStrategy struct {
	// 重试次数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 重试间隔

	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

type CreateToolParseURLRequest struct {
	*tchttp.BaseRequest

	// 文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

func (r *CreateToolParseURLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateToolParseURLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListToolLabelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 标签列表

		Labels []*string `json:"Labels,omitempty" name:"Labels"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListToolLabelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListToolLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Zone struct {
	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 中文名

	CnName *string `json:"CnName,omitempty" name:"CnName"`
	// region名字

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// region中文名

	RegionCnName *string `json:"RegionCnName,omitempty" name:"RegionCnName"`
}

type DescribeOperationInstructionPlanRequest struct {
	*tchttp.BaseRequest

	// 变更单名字

	OperationInstruction *string `json:"OperationInstruction,omitempty" name:"OperationInstruction"`
}

func (r *DescribeOperationInstructionPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationInstructionPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListZonesRequest struct {
	*tchttp.BaseRequest
}

func (r *ListZonesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListZonesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParseTarFileDirTreeRequest struct {
	*tchttp.BaseRequest

	// 上传ID

	UploadId *string `json:"UploadId,omitempty" name:"UploadId"`
}

func (r *ParseTarFileDirTreeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ParseTarFileDirTreeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Alert struct {
	// 告警名

	AlertName *string `json:"AlertName,omitempty" name:"AlertName"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 告警策略

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 归属产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 告警对象

	Object *string `json:"Object,omitempty" name:"Object"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 确认人

	ConfirmedBy *string `json:"ConfirmedBy,omitempty" name:"ConfirmedBy"`
	// 确认时间，为空表示未确认

	ConfirmedAt *string `json:"ConfirmedAt,omitempty" name:"ConfirmedAt"`
	// 恢复状态，firing未恢复，resolved已恢复

	Status *string `json:"Status,omitempty" name:"Status"`
	// 告警级别，1级严重，2级重要，3级次要，4级警告

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 告警ID

	NotificationId *int64 `json:"NotificationId,omitempty" name:"NotificationId"`
	// 恢复时间

	EndsAt *string `json:"EndsAt,omitempty" name:"EndsAt"`
}

type ComponentWithInstance struct {
	// 组件名字

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 组件版本

	ComponentVersion *string `json:"ComponentVersion,omitempty" name:"ComponentVersion"`
	// 组件类型

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
	// 组件架构

	ComponentArch *string `json:"ComponentArch,omitempty" name:"ComponentArch"`
	// 部署级别, global/region/zone

	DeployLevel *string `json:"DeployLevel,omitempty" name:"DeployLevel"`
	// 下属实例列表

	Instances []*ComponentInstance `json:"Instances,omitempty" name:"Instances"`
	// 最近变更时间

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
}

type ProductComponentStatus struct {
	// 健康的pod数量

	NormalCount *int64 `json:"NormalCount,omitempty" name:"NormalCount"`
	// 不健康的pod数量

	AbnormalCount *int64 `json:"AbnormalCount,omitempty" name:"AbnormalCount"`
	// 组件名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type ModifyToolRequest struct {
	*tchttp.BaseRequest

	// 工具信息

	Manifest *Manifest `json:"Manifest,omitempty" name:"Manifest"`
	// 上传ID

	UploadId *string `json:"UploadId,omitempty" name:"UploadId"`
	// 文件来源

	FileSource *FileSource `json:"FileSource,omitempty" name:"FileSource"`
	// 解压压缩包上传ID

	ParseUploadId *string `json:"ParseUploadId,omitempty" name:"ParseUploadId"`
}

func (r *ModifyToolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyToolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetTaskInstanceStatusRequest struct {
	*tchttp.BaseRequest

	// 步骤ID

	TaskInstance *string `json:"TaskInstance,omitempty" name:"TaskInstance"`
	// 状态, 可选retry/skip

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *SetTaskInstanceStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetTaskInstanceStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommandInput struct {
	// 参数名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 默认值

	Default *string `json:"Default,omitempty" name:"Default"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 是否可选

	Optional *bool `json:"Optional,omitempty" name:"Optional"`
	// 入参值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateUpdatePlanTaskRequest struct {
	*tchttp.BaseRequest

	// 当前规划版本

	SourcePlanVersion *string `json:"SourcePlanVersion,omitempty" name:"SourcePlanVersion"`
	// 目标规划版本

	PlanVersion *string `json:"PlanVersion,omitempty" name:"PlanVersion"`
	// 物料机IP

	MasterNode *string `json:"MasterNode,omitempty" name:"MasterNode"`
	// 变更类型, update/rollback

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *CreateUpdatePlanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUpdatePlanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// flow的DAG图

		DAG *DAG `json:"DAG,omitempty" name:"DAG"`
		// 标签列表

		Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
		// 注解列表

		Annotations []*KVPair `json:"Annotations,omitempty" name:"Annotations"`
		// flow运行状态, Completed/Progressing/Failed/Error

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationInstructionComponentFlowRequest struct {
	*tchttp.BaseRequest

	// 变更单ID

	Name *string `json:"Name,omitempty" name:"Name"`
	// 类型, update/rollback

	Type *string `json:"Type,omitempty" name:"Type"`
	// 组件名

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
}

func (r *DescribeOperationInstructionComponentFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationInstructionComponentFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateToolUploadURLRequest struct {
	*tchttp.BaseRequest

	// 工具名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 文件名列表

	FileNames []*string `json:"FileNames,omitempty" name:"FileNames"`
}

func (r *CreateToolUploadURLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateToolUploadURLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBizTreeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 业务树列表

		BizTreeList []*BizTree `json:"BizTreeList,omitempty" name:"BizTreeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBizTreeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBizTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationStatsRequest struct {
	*tchttp.BaseRequest

	// 开始时间, 格式如2021-10-01T00:00:00Z

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间, 格式如2021-10-01T00:00:00Z

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeOperationStatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationStatsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentInstanceYAMLRequest struct {
	*tchttp.BaseRequest

	// 实例名字

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 组件类型

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
}

func (r *DescribeComponentInstanceYAMLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentInstanceYAMLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaterialNodeListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMaterialNodeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaterialNodeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollbackToolRequest struct {
	*tchttp.BaseRequest

	// 工具ID

	ToolId *string `json:"ToolId,omitempty" name:"ToolId"`
	// 回滚目标版本

	TargetVersion *string `json:"TargetVersion,omitempty" name:"TargetVersion"`
}

func (r *RollbackToolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RollbackToolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationSheetDetail struct {
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 当前状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 变更单执行流程DAG图

	DAG *DAG `json:"DAG,omitempty" name:"DAG"`
	// 变更单YAML

	YAML *string `json:"YAML,omitempty" name:"YAML"`
	// 总步骤数量

	TotalSteps *int64 `json:"TotalSteps,omitempty" name:"TotalSteps"`
	// 已完成的步骤数量

	CompletedSteps *int64 `json:"CompletedSteps,omitempty" name:"CompletedSteps"`
	// 用于显示的名字, 不唯一

	ShowName *string `json:"ShowName,omitempty" name:"ShowName"`
	// 关联的TAPD单子

	TAPD *string `json:"TAPD,omitempty" name:"TAPD"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 产品名

	Product *string `json:"Product,omitempty" name:"Product"`
}

type CreatePackageUploadURLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件csp路径

		FileKey *string `json:"FileKey,omitempty" name:"FileKey"`
		// 链接

		URL *string `json:"URL,omitempty" name:"URL"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePackageUploadURLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePackageUploadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListToolsRequest struct {
	*tchttp.BaseRequest

	// 每页数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListToolsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListToolsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateArgument struct {
	// Flow名字

	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`
	// 步骤名字

	StepName *string `json:"StepName,omitempty" name:"StepName"`
	// 新参数

	Arguments *Arguments `json:"Arguments,omitempty" name:"Arguments"`
}

type ListProjectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的项目总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 项目列表

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

type OperationInstructionReadyToStartResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否可启动

		ReadyToStart *bool `json:"ReadyToStart,omitempty" name:"ReadyToStart"`
		// 处理中的组件及对应的变更单

		ProcessingComponents []*ComponentInstruction `json:"ProcessingComponents,omitempty" name:"ProcessingComponents"`
		// 枚举值. NotLatest表示不是最新启动的变更单, 无法回滚. Processing表示有组件正在别的变更单中处理, 无法启动

		NotReadyReason *string `json:"NotReadyReason,omitempty" name:"NotReadyReason"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OperationInstructionReadyToStartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OperationInstructionReadyToStartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PrepareTaskStep struct {
	// 已完成步骤数量

	CompletedSteps *int64 `json:"CompletedSteps,omitempty" name:"CompletedSteps"`
	// 总步骤数量

	TotalSteps *int64 `json:"TotalSteps,omitempty" name:"TotalSteps"`
	// 步骤描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 错误信息

	Error *string `json:"Error,omitempty" name:"Error"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type StepNode struct {
	// 主机名字

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 启动时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 日志

	Message *string `json:"Message,omitempty" name:"Message"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 变更流程名字

	TaskInstanceName *string `json:"TaskInstanceName,omitempty" name:"TaskInstanceName"`
}

type AssociatedComponentsSet struct {
	// 云产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 组件列表

	Components []*string `json:"Components,omitempty" name:"Components"`
}

type Commad struct {
	// 命令名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数列表

	Inputs []*CommandInput `json:"Inputs,omitempty" name:"Inputs"`
	// 命令内容

	Command *string `json:"Command,omitempty" name:"Command"`
	// 退出码

	ExitCodes []*ExitCode `json:"ExitCodes,omitempty" name:"ExitCodes"`
}

type CreateOperationInstructionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 变更单ID

		Name *string `json:"Name,omitempty" name:"Name"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperationInstructionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationInstructionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteToolRequest struct {
	*tchttp.BaseRequest

	// 工具ID

	ToolId *string `json:"ToolId,omitempty" name:"ToolId"`
}

func (r *DeleteToolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteToolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBizTreeHostsRequest struct {
	*tchttp.BaseRequest

	// 查询指定云产品的所有结点

	BizUuid *string `json:"BizUuid,omitempty" name:"BizUuid"`
	// 查询指定Set的所有节点

	SetUuid *string `json:"SetUuid,omitempty" name:"SetUuid"`
	// 查询指定Module的Uuid

	ModuleUuid *string `json:"ModuleUuid,omitempty" name:"ModuleUuid"`
}

func (r *DescribeBizTreeHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBizTreeHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// YAML内容

		YAML *string `json:"YAML,omitempty" name:"YAML"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComponentPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductComponentResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// pod列表

		Pods []*TreeNodePod `json:"Pods,omitempty" name:"Pods"`
		// 主机列表

		Hosts []*Host `json:"Hosts,omitempty" name:"Hosts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductComponentResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductComponentResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TreeNodeContainerComponent struct {
	// 该组件下的Pod列表

	Pods *TreeNodePod `json:"Pods,omitempty" name:"Pods"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeMaterialNodeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机列表

		Hosts []*Host `json:"Hosts,omitempty" name:"Hosts"`
		// DBSQL物料机

		DBSqlMaterialNodes []*Host `json:"DBSqlMaterialNodes,omitempty" name:"DBSqlMaterialNodes"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMaterialNodeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaterialNodeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollbackToolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RollbackToolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RollbackToolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExecutionTrendPoint struct {
	// 日期, 如2022-10-10

	Date *string `json:"Date,omitempty" name:"Date"`
	// 当天执行的工具数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type DescribeOperationInstructionLogRequest struct {
	*tchttp.BaseRequest

	// Flow名字

	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`
	// 步骤名字

	Step *string `json:"Step,omitempty" name:"Step"`
	// 结点名字

	Node *string `json:"Node,omitempty" name:"Node"`
	// 变更单ID

	OperationInstruction *string `json:"OperationInstruction,omitempty" name:"OperationInstruction"`
}

func (r *DescribeOperationInstructionLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationInstructionLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkflowCommandsRequest struct {
	*tchttp.BaseRequest
}

func (r *ListWorkflowCommandsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListWorkflowCommandsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Flow struct {
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 标签列表

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// 注解列表

	Annotations []*KVPair `json:"Annotations,omitempty" name:"Annotations"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 启动时间

	BeginTimestamp *string `json:"BeginTimestamp,omitempty" name:"BeginTimestamp"`
	// 完成时间

	CompletionTimestamp *string `json:"CompletionTimestamp,omitempty" name:"CompletionTimestamp"`
	// 当前状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 发起人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type OperationInstruction struct {
	// 变更单ID

	Name *string `json:"Name,omitempty" name:"Name"`
	// 变更单名字

	ShowName *string `json:"ShowName,omitempty" name:"ShowName"`
	// 当前状态, Ready/Running/Paused/Completed/Acked/Failed/Aborted/Rollbacked, 分别对应待运行/运行中/暂停/完成未确认/完成已确认/失败终止/人工终止/已回滚

	Status *string `json:"Status,omitempty" name:"Status"`
	// 已完成的组件数量

	Completed *int64 `json:"Completed,omitempty" name:"Completed"`
	// 组件总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 组件详情

	Components []*ComponentInstruction `json:"Components,omitempty" name:"Components"`
	// 涉及到的组件

	AssociatedComponents []*string `json:"AssociatedComponents,omitempty" name:"AssociatedComponents"`
	// 回滚涉及到的组件, 如为空, 表示不涉及回滚

	AssociatedRollbackComponents []*string `json:"AssociatedRollbackComponents,omitempty" name:"AssociatedRollbackComponents"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 启动时间

	StartTimestamp *string `json:"StartTimestamp,omitempty" name:"StartTimestamp"`
	// 结束时间

	FinishTimestamp *string `json:"FinishTimestamp,omitempty" name:"FinishTimestamp"`
	// 回滚启动时间

	RollbackStartTimestamp *string `json:"RollbackStartTimestamp,omitempty" name:"RollbackStartTimestamp"`
	// 回滚结束时间

	RollbackFinishTimestamp *string `json:"RollbackFinishTimestamp,omitempty" name:"RollbackFinishTimestamp"`
	// 上一次操作人, 默认为操作账号的UIN

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 可用区列表. 如果为空说明没有az级别的组件

	ZoneOrder []*string `json:"ZoneOrder,omitempty" name:"ZoneOrder"`
	// 回滚状态, 与Status字段类似

	RollbackStatus *string `json:"RollbackStatus,omitempty" name:"RollbackStatus"`
	// 是否含有回滚流程

	WithRollback *bool `json:"WithRollback,omitempty" name:"WithRollback"`
	// 是否可回滚

	ReadyToRollback *bool `json:"ReadyToRollback,omitempty" name:"ReadyToRollback"`
	// 大版本号

	TCEVersion *string `json:"TCEVersion,omitempty" name:"TCEVersion"`
	// 产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 组件执行策略

	RunningStrategy *RunningStrategy `json:"RunningStrategy,omitempty" name:"RunningStrategy"`
	// 是否资源已被回收

	Recycled *bool `json:"Recycled,omitempty" name:"Recycled"`
}

type CreatePackageUploadURLRequest struct {
	*tchttp.BaseRequest

	// 文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

func (r *CreatePackageUploadURLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePackageUploadURLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUploadPackageTaskRequest struct {
	*tchttp.BaseRequest

	// 变更单名字

	OperationInstruction *string `json:"OperationInstruction,omitempty" name:"OperationInstruction"`
	// 物料机IP

	MasterNode *string `json:"MasterNode,omitempty" name:"MasterNode"`
	// 组件信息

	Components []*ComponentInstruction `json:"Components,omitempty" name:"Components"`
	// 跳过适配规划流程

	SkipAdaptInstance *bool `json:"SkipAdaptInstance,omitempty" name:"SkipAdaptInstance"`
	// 跳过上传镜像

	SkipImageComponents *bool `json:"SkipImageComponents,omitempty" name:"SkipImageComponents"`
}

func (r *CreateUploadPackageTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUploadPackageTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationInstructionsRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤器, 支持Name/Status/CreationTimestamp

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 是否需要直接从master获取数据, 而不从缓存中获取. 默认false

	Direct *bool `json:"Direct,omitempty" name:"Direct"`
	// 变更类型, upgrade(大版本升级)/bugfix

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *ListOperationInstructionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationInstructionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReadyToRollbackPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否可回滚

		ReadyToRollback *bool `json:"ReadyToRollback,omitempty" name:"ReadyToRollback"`
		// 不可回滚的原因, 目前有NotStartYet(未启动)/NotLatestStarted(不是最新启动的任务)

		Reason *string `json:"Reason,omitempty" name:"Reason"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReadyToRollbackPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReadyToRollbackPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Topology struct {
	// Region级别信息, 主要是路由

	Region *ProductTopologyNode `json:"Region,omitempty" name:"Region"`
	// Zone级别信息, 主要是组件

	Zones []*ProductTopologyNode `json:"Zones,omitempty" name:"Zones"`
	// 支撑组件

	Middleware *ProductTopologyNode `json:"Middleware,omitempty" name:"Middleware"`
}

type DeleteFlowRequest struct {
	*tchttp.BaseRequest

	// flow名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductBizTreeRequest struct {
	*tchttp.BaseRequest

	// 父级云产品名

	ParentName *string `json:"ParentName,omitempty" name:"ParentName"`
	// 云产品名

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeProductBizTreeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductBizTreeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetOperationInstructionStatusRequest struct {
	*tchttp.BaseRequest

	// 变更单ID

	Name *string `json:"Name,omitempty" name:"Name"`
	// 变更单状态, 支持Rollbacked/Running/Paused/Acked/Aborted, 分别表示确认回滚/启动变更/暂停变更/确认完成/终止任务

	Status *string `json:"Status,omitempty" name:"Status"`
	// 如果是设置成已回滚状态, 需要填写回滚时间. 时间格式如1985-04-12T23:20:50Z

	RollbackTimestamp *string `json:"RollbackTimestamp,omitempty" name:"RollbackTimestamp"`
	// 操作用户, 不填默认为当前登录用户UIN

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 针对更新操作还是回滚操作设置状态. update/rollback

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *SetOperationInstructionStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetOperationInstructionStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetOperationInstructionStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetOperationInstructionStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetOperationInstructionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StepStrategy struct {
	// 重试策略

	Retry *RetryStrategy `json:"Retry,omitempty" name:"Retry"`
	// 超时时间, 例如10s/6m/1h

	Timeout *string `json:"Timeout,omitempty" name:"Timeout"`
	// 忽略错误继续执行

	IgnoreError *bool `json:"IgnoreError,omitempty" name:"IgnoreError"`
}

type ExecuteToolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务名称

		TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExecuteToolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExecuteToolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductComponentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品列表

		Products []*CloudProduct `json:"Products,omitempty" name:"Products"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductComponentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationInstructionRequest struct {
	*tchttp.BaseRequest

	// 修改升级单还是回滚单. update/rollback. 不填默认update

	Type *string `json:"Type,omitempty" name:"Type"`
	// 更新可用区顺序

	ZoneOrder []*string `json:"ZoneOrder,omitempty" name:"ZoneOrder"`
	// 更新参数

	Arguments []*UpdateArgument `json:"Arguments,omitempty" name:"Arguments"`
	// 变更单name

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *UpdateOperationInstructionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationInstructionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云产品详情

		Product *CloudProduct `json:"Product,omitempty" name:"Product"`
		// 变更数量

		OperationCount *int64 `json:"OperationCount,omitempty" name:"OperationCount"`
		// 成功的变更数量

		SucceededOperationCount *int64 `json:"SucceededOperationCount,omitempty" name:"SucceededOperationCount"`
		// 失败的变更数量

		FailedOperationCount *int64 `json:"FailedOperationCount,omitempty" name:"FailedOperationCount"`
		// 运行中的变更数量

		RunningOperationCount *int64 `json:"RunningOperationCount,omitempty" name:"RunningOperationCount"`
		// 按组件维度统计变更数量

		OperationCountByComponent []*KVPair `json:"OperationCountByComponent,omitempty" name:"OperationCountByComponent"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MapInfrastoreProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云哨中的产品概念

		Products []*InfrastoreProduct `json:"Products,omitempty" name:"Products"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MapInfrastoreProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MapInfrastoreProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Manifest struct {
	// 工具Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 工具名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 工具版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 使用文档

	Doc *ToolDoc `json:"Doc,omitempty" name:"Doc"`
	// 运行环境

	RunningEnv *RunningEnv `json:"RunningEnv,omitempty" name:"RunningEnv"`
	// 命令列表

	Commands []*Commad `json:"Commands,omitempty" name:"Commands"`
	// 分类标签

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
}

type DeleteToolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteToolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteToolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductAlertsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 告警列表

		Alerts []*Alert `json:"Alerts,omitempty" name:"Alerts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductAlertsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductAlertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationInstructionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 变更控制表详情

		OperationInstruction *OperationInstruction `json:"OperationInstruction,omitempty" name:"OperationInstruction"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationInstructionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationInstructionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Arguments struct {
	// 重试次数

	RetryTimes *int64 `json:"RetryTimes,omitempty" name:"RetryTimes"`
	// 并发策略, Concurrent/Serial

	ConcurrencyPolicy *string `json:"ConcurrencyPolicy,omitempty" name:"ConcurrencyPolicy"`
	// 自定义参数

	CustomParameters []*CustomParameter `json:"CustomParameters,omitempty" name:"CustomParameters"`
	// 组件参数

	Artifacts []*Artifact `json:"Artifacts,omitempty" name:"Artifacts"`
	// 多阶段命令参数

	NodeStages []*NodeStage `json:"NodeStages,omitempty" name:"NodeStages"`
}

type NodeStage struct {
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 步骤列表

	Step *NodeStageStep `json:"Step,omitempty" name:"Step"`
}

type ListToolLabelsRequest struct {
	*tchttp.BaseRequest

	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 搜索关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 按产品过滤

	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *ListToolLabelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListToolLabelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TreeNodePod struct {
	// Pod名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Pod下的容器列表

	Containers []*string `json:"Containers,omitempty" name:"Containers"`
	// pod所在集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// Pod所在namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// region名字

	Region *string `json:"Region,omitempty" name:"Region"`
	// az名字

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 当前状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type ZoneDAG struct {
	// Zone ID

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// TFlow名

	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`
	// 流程图

	DAG *DAG `json:"DAG,omitempty" name:"DAG"`
	// 流程当前状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// Region ID

	Region *string `json:"Region,omitempty" name:"Region"`
}

type AdaptInstructionWithPlanRequest struct {
	*tchttp.BaseRequest

	// 变更单名字

	OperationInstruction *string `json:"OperationInstruction,omitempty" name:"OperationInstruction"`
	// 忽略组件列表

	IgnoredComponents []*string `json:"IgnoredComponents,omitempty" name:"IgnoredComponents"`
}

func (r *AdaptInstructionWithPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AdaptInstructionWithPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListZonesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可用区列表

		Zones []*Zone `json:"Zones,omitempty" name:"Zones"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomParameter struct {
	// 参数名字

	Key *string `json:"Key,omitempty" name:"Key"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 参数是否是可选的

	Optional *bool `json:"Optional,omitempty" name:"Optional"`
	// 参数描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 是否可编辑, 暂时没有用到

	Editable *bool `json:"Editable,omitempty" name:"Editable"`
	// 参数类型, 更新时原样传给后端

	Type *string `json:"Type,omitempty" name:"Type"`
	// 用于展示的参数名

	ShowName *string `json:"ShowName,omitempty" name:"ShowName"`
}

type AdaptInstructionWithPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AdaptInstructionWithPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AdaptInstructionWithPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudProductRequest struct {
	*tchttp.BaseRequest

	// 云产品名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 变更单查询, 开始时间, 格式如2021-10-01T00:00:00Z

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 变更单查询, 结束时间, 格式如2021-10-01T00:00:00Z

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeCloudProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductComponentStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件状态列表

		Components []*ProductComponentStatus `json:"Components,omitempty" name:"Components"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductComponentStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductComponentStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExecuteToolRequest struct {
	*tchttp.BaseRequest

	// 工具ID

	ToolId *string `json:"ToolId,omitempty" name:"ToolId"`
	// 子命令名字

	CommandName *string `json:"CommandName,omitempty" name:"CommandName"`
	// 命令入参

	Input []*CommandInput `json:"Input,omitempty" name:"Input"`
	// 待运行主机IP列表

	Nodes []*string `json:"Nodes,omitempty" name:"Nodes"`
}

func (r *ExecuteToolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExecuteToolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationTaskLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志

		Log *string `json:"Log,omitempty" name:"Log"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOperationTaskLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationTaskLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileURL struct {
	// 文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件标识

	FileKey *string `json:"FileKey,omitempty" name:"FileKey"`
	// 上传链接

	URL *string `json:"URL,omitempty" name:"URL"`
}

type ComponentFlow struct {
	// 准备阶段步骤

	PrepareStage []*DAGStep `json:"PrepareStage,omitempty" name:"PrepareStage"`
	// 执行阶段步骤

	ImplementStage []*DAGStep `json:"ImplementStage,omitempty" name:"ImplementStage"`
	// 检查阶段步骤

	CheckStage []*DAGStep `json:"CheckStage,omitempty" name:"CheckStage"`
	// YAML

	YAML *DAGStep `json:"YAML,omitempty" name:"YAML"`
}

type DeleteOperationInstructionRequest struct {
	*tchttp.BaseRequest

	// 待删除变更单名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteOperationInstructionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationInstructionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListFlowsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Flow列表

		Flows []*Flow `json:"Flows,omitempty" name:"Flows"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListFlowsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListFlowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationInstructionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOperationInstructionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationInstructionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentCurrentVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件详情

		Components []*ComponentInstruction `json:"Components,omitempty" name:"Components"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComponentCurrentVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentCurrentVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunningStrategy struct {
	// 策略名称, 枚举值: [serial, manual, concurrent], 分别表示串行执行, 手动执行, 并行执行

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 并行执行时, 同时可执行的组件数量

	MaxConcurrent *int64 `json:"MaxConcurrent,omitempty" name:"MaxConcurrent"`
	// 并行分组列表, 分组内部并行执行, 分组之间串行执行

	Partitions []*Partition `json:"Partitions,omitempty" name:"Partitions"`
}

type DescribeProductComponentStatusRequest struct {
	*tchttp.BaseRequest

	// 父级云产品名

	ParentName *string `json:"ParentName,omitempty" name:"ParentName"`
	// 云产品名

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeProductComponentStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductComponentStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 项目详情

		Project *Project `json:"Project,omitempty" name:"Project"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductComponentsRequest struct {
	*tchttp.BaseRequest
}

func (r *ListProductComponentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductComponentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUploadPackagesTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 上传物料任务ID

		TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUploadPackagesTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUploadPackagesTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectPlanRequest struct {
	*tchttp.BaseRequest

	// 变更单ID

	Project *string `json:"Project,omitempty" name:"Project"`
}

func (r *DescribeProjectPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudProductsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品名列表. 每个云产品结点如果HasChildren为True, 需要再次调用此接口查询子节点.

		Products []*TreeNode `json:"Products,omitempty" name:"Products"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCloudProductsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectsRequest struct {
	*tchttp.BaseRequest

	// 过滤器. 支持CreationTimestamp, Name, Version, Creator

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每一页的数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListProjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudProduct struct {
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 告警是否已经配置

	AlertConfiged *bool `json:"AlertConfiged,omitempty" name:"AlertConfiged"`
	// 首次部署时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 最后变更时间, 可能为空

	LastUpdateTimestamp *string `json:"LastUpdateTimestamp,omitempty" name:"LastUpdateTimestamp"`
	// 用于显示的云产品名字

	ShowName *string `json:"ShowName,omitempty" name:"ShowName"`
	// 支撑组件数量

	MiddlewareComponentsCount *int64 `json:"MiddlewareComponentsCount,omitempty" name:"MiddlewareComponentsCount"`
	// 组件数量

	ComponentsCount *int64 `json:"ComponentsCount,omitempty" name:"ComponentsCount"`
	// cmdb中的结点ID

	CMDBInstanceId *int64 `json:"CMDBInstanceId,omitempty" name:"CMDBInstanceId"`
	// 组件列表

	Components []*ComponentWithInstance `json:"Components,omitempty" name:"Components"`
}

type DescribeCurrentPlanVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前规划版本

		CurrentPlanVersion *string `json:"CurrentPlanVersion,omitempty" name:"CurrentPlanVersion"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCurrentPlanVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCurrentPlanVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComponentInstruction struct {
	// 组件名, 如ocloud-tcenter-message-svr

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 组件版本

	ComponentVersion *string `json:"ComponentVersion,omitempty" name:"ComponentVersion"`
	// 组件架构

	ComponentArch *string `json:"ComponentArch,omitempty" name:"ComponentArch"`
	// 组件所属云产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 变更内容描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 组件部署级别, global/region/zone

	DeployLevel *string `json:"DeployLevel,omitempty" name:"DeployLevel"`
	// TFlow名

	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`
	// TFlow YAML, 组件DeployLevel!=zone时才返回

	YAML *string `json:"YAML,omitempty" name:"YAML"`
	// 状态,组件DeployLevel!=zone时才返回

	Status *string `json:"Status,omitempty" name:"Status"`
	// 准备阶段

	PrepareStage []*DAGStep `json:"PrepareStage,omitempty" name:"PrepareStage"`
	// 实施阶段

	ImplementStage []*DAGStep `json:"ImplementStage,omitempty" name:"ImplementStage"`
	// 检查阶段

	CheckStage []*DAGStep `json:"CheckStage,omitempty" name:"CheckStage"`
	// update/rollback

	Type *string `json:"Type,omitempty" name:"Type"`
	// 组件类型, image/product/preset/dbsql/yunapi

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
	// 如果组件DeployLevel为zone/region级别, 返回本字段

	DAGs []*ZoneDAG `json:"DAGs,omitempty" name:"DAGs"`
	// 变更单ID

	OperationInstruction *string `json:"OperationInstruction,omitempty" name:"OperationInstruction"`
	// 当前版本

	CurrentVersion *string `json:"CurrentVersion,omitempty" name:"CurrentVersion"`
	// 目标版本

	TargetVersion *string `json:"TargetVersion,omitempty" name:"TargetVersion"`
	// 实例列表

	Instances []*ComponentInstance `json:"Instances,omitempty" name:"Instances"`
}

type WorkflowCommand struct {
	// 原子命令名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 原子命令中文名

	ShowName *string `json:"ShowName,omitempty" name:"ShowName"`
	// 原子命令描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 支持的参数

	Parameters *CustomParameter `json:"Parameters,omitempty" name:"Parameters"`
	// 组件参数

	Artifacts []*Artifact `json:"Artifacts,omitempty" name:"Artifacts"`
}

type CreateToolRequest struct {
	*tchttp.BaseRequest

	// 工具信息

	Manifest *Manifest `json:"Manifest,omitempty" name:"Manifest"`
	// 上传ID

	UploadId *string `json:"UploadId,omitempty" name:"UploadId"`
	// 创建模式，package上传压缩包，manual在线编写

	CreateMod *string `json:"CreateMod,omitempty" name:"CreateMod"`
	// 解析压缩包的上传ID

	ParseUploadId *string `json:"ParseUploadId,omitempty" name:"ParseUploadId"`
	// 文件来源

	FileSource *FileSource `json:"FileSource,omitempty" name:"FileSource"`
}

func (r *CreateToolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateToolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUpdatePlanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务名字

		TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUpdatePlanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUpdatePlanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUploadPackagesTaskRequest struct {
	*tchttp.BaseRequest

	// 查询物料路径任务ID

	FindPackageTaskName *string `json:"FindPackageTaskName,omitempty" name:"FindPackageTaskName"`
	// 物料机IP

	MasterNode *string `json:"MasterNode,omitempty" name:"MasterNode"`
	// 变更单YAML

	OperationInstruction *string `json:"OperationInstruction,omitempty" name:"OperationInstruction"`
	// 组件版本信息

	Components []*ComponentInstruction `json:"Components,omitempty" name:"Components"`
}

func (r *CreateUploadPackagesTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUploadPackagesTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProjectPartition struct {
	// 变更单列表

	OperationInstructions []*OperationInstruction `json:"OperationInstructions,omitempty" name:"OperationInstructions"`
}

type DescribeOperationInstructionRequest struct {
	*tchttp.BaseRequest

	// 变更单ID

	Name *string `json:"Name,omitempty" name:"Name"`
	// 变更单类型, update/rollback分别表示查看升级变更单及回滚变更单

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeOperationInstructionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationInstructionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductAlertStatsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 严重告警数量

		Severity1 *int64 `json:"Severity1,omitempty" name:"Severity1"`
		// 重要告警数量

		Severity2 *int64 `json:"Severity2,omitempty" name:"Severity2"`
		// 次要告警数量

		Severity3 *int64 `json:"Severity3,omitempty" name:"Severity3"`
		// 警告告警数量

		Severity4 *int64 `json:"Severity4,omitempty" name:"Severity4"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductAlertStatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductAlertStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeToolRequest struct {
	*tchttp.BaseRequest

	// 工具ID

	ToolId *string `json:"ToolId,omitempty" name:"ToolId"`
	// 工具版本，不传默认当前版本

	ToolVersion *string `json:"ToolVersion,omitempty" name:"ToolVersion"`
}

func (r *DescribeToolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeToolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StageStepArgument struct {
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type OperationTrendPoint struct {
	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 当日变更单总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 当日成功变更单数量

	Succeed *int64 `json:"Succeed,omitempty" name:"Succeed"`
	// 当日失败变更单数量

	Failed *int64 `json:"Failed,omitempty" name:"Failed"`
	// 当日运行中变更单数量

	Running *int64 `json:"Running,omitempty" name:"Running"`
}

type Set struct {
	// uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// Model UUID

	ModelUuid *string `json:"ModelUuid,omitempty" name:"ModelUuid"`
	// 模块列表

	ModuleList []*Module `json:"ModuleList,omitempty" name:"ModuleList"`
	// 名字

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DescribeCurrentPlanVersionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCurrentPlanVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCurrentPlanVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MapInfrastoreProductRequest struct {
	*tchttp.BaseRequest

	// 产品名

	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *MapInfrastoreProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MapInfrastoreProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReadyToRollbackPlanRequest struct {
	*tchttp.BaseRequest

	// 规划变更任务名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ReadyToRollbackPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReadyToRollbackPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComponentInstance struct {
	// 实例名字

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// region名字

	Region *string `json:"Region,omitempty" name:"Region"`
	// zone名字

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 健康状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 生产组件序号

	ClusterIndex *string `json:"ClusterIndex,omitempty" name:"ClusterIndex"`
	// 原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// IDC ID

	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
	// 规划是否已校正

	MergeChecked *bool `json:"MergeChecked,omitempty" name:"MergeChecked"`
}

type DescribeProductAlertStatsRequest struct {
	*tchttp.BaseRequest

	// 云产品UUID

	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeProductAlertStatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductAlertStatsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportToolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导入任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportToolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportToolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListExecutionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合筛选条件的任务数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 任务列表

		ExecutionTasks []*ExecutionSummary `json:"ExecutionTasks,omitempty" name:"ExecutionTasks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListExecutionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListExecutionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Artifact struct {
	// 参数名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件名

	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 架构, arm/x86

	Arch *string `json:"Arch,omitempty" name:"Arch"`
}

type DAG struct {
	// 步骤列表

	Steps []*DAGStep `json:"Steps,omitempty" name:"Steps"`
	// 边列表

	Edges []*DAGStep `json:"Edges,omitempty" name:"Edges"`
}

type Tool struct {
	// 工具ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 工具名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 工具版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 分类标签

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
}

type UploadStatistics struct {
	// 当前状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 出错时显示错误信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 完成的组件数量

	Completed *int64 `json:"Completed,omitempty" name:"Completed"`
	// 运行中的组件数量

	Running *int64 `json:"Running,omitempty" name:"Running"`
	// 跳过的组件数量

	Skipped *int64 `json:"Skipped,omitempty" name:"Skipped"`
	// 总的组件数量

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 失败的组件数量

	Failed *int64 `json:"Failed,omitempty" name:"Failed"`
}

type DescribeProductComponentResourceRequest struct {
	*tchttp.BaseRequest

	// 云产品uuid

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 组件名字

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
}

func (r *DescribeProductComponentResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductComponentResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeToolDownloadURLRequest struct {
	*tchttp.BaseRequest

	// 工具ID

	ToolId *string `json:"ToolId,omitempty" name:"ToolId"`
	// 工具版本

	ToolVersion *string `json:"ToolVersion,omitempty" name:"ToolVersion"`
}

func (r *DescribeToolDownloadURLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeToolDownloadURLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportChangePackageRequest struct {
	*tchttp.BaseRequest

	// 文件csp路径

	FileKey *string `json:"FileKey,omitempty" name:"FileKey"`
	// 变更类型，默认bugfix(缺陷修复)，可填upgrade(大版本升级)

	ChangeType *string `json:"ChangeType,omitempty" name:"ChangeType"`
	// 项目ID

	ProjectID *string `json:"ProjectID,omitempty" name:"ProjectID"`
}

func (r *ImportChangePackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportChangePackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListToolVersionsRequest struct {
	*tchttp.BaseRequest

	// 每页总数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 搜索参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 工具ID

	ToolId *string `json:"ToolId,omitempty" name:"ToolId"`
}

func (r *ListToolVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListToolVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentCurrentVersionRequest struct {
	*tchttp.BaseRequest

	// 变更单ID

	OperationInstruction *string `json:"OperationInstruction,omitempty" name:"OperationInstruction"`
}

func (r *DescribeComponentCurrentVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentCurrentVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationInstructionComponentFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 变更详情

		ComponentFlow *ComponentFlow `json:"ComponentFlow,omitempty" name:"ComponentFlow"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationInstructionComponentFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationInstructionComponentFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateToolUploadURLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 上传链接列表

		URLs []*FileURL `json:"URLs,omitempty" name:"URLs"`
		// 上传ID

		UploadId *string `json:"UploadId,omitempty" name:"UploadId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateToolUploadURLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateToolUploadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListExecutionsRequest struct {
	*tchttp.BaseRequest

	// 过滤器, 支持搜索Name/ToolName/Status/Operator

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询区间开始时间, 例如"2020-10-10T16:00:00"

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询区间区间时间, 例如"2020-10-10T16:00:00"

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ListExecutionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListExecutionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExitCode struct {
	// 操作符（=,><）

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 退出码

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 退出码信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type OperationSheet struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 启动时间

	StartTimestamp *string `json:"StartTimestamp,omitempty" name:"StartTimestamp"`
	// 结束时间

	FinishTimestamp *string `json:"FinishTimestamp,omitempty" name:"FinishTimestamp"`
	// 当前状态, Ready/Running/Paused/Completed/Acked/Failed/Aborted/Rollbacked, 分别对应待运行/运行中/暂停/完成未确认/完成已确认/失败终止/人工终止/已回滚

	Status *string `json:"Status,omitempty" name:"Status"`
	// 上次操作人员uin

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 回滚时间

	RollbackTimestamp *string `json:"RollbackTimestamp,omitempty" name:"RollbackTimestamp"`
	// 变更单内容

	YAML *string `json:"YAML,omitempty" name:"YAML"`
	// 回滚步骤说明

	RollbackDescription *string `json:"RollbackDescription,omitempty" name:"RollbackDescription"`
	// 变更单步骤详情

	DAG *DAG `json:"DAG,omitempty" name:"DAG"`
}

type OperationInstructionReadyToStartRequest struct {
	*tchttp.BaseRequest

	// 变更单ID

	Name *string `json:"Name,omitempty" name:"Name"`
	// 升级类型还是回滚类型, 不填默认升级

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *OperationInstructionReadyToStartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OperationInstructionReadyToStartRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationSheetArgument struct {
	// 步骤名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数列表

	Arguments *Arguments `json:"Arguments,omitempty" name:"Arguments"`
}

type ListWorkflowCommandsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 原子命令列表

		Commands []*WorkflowCommand `json:"Commands,omitempty" name:"Commands"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListWorkflowCommandsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListWorkflowCommandsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Project struct {
	// 唯一ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 项目名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 目标版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 变更单数量

	OperationInstructionCount *int64 `json:"OperationInstructionCount,omitempty" name:"OperationInstructionCount"`
	// 组件数量

	ComponentCount *int64 `json:"ComponentCount,omitempty" name:"ComponentCount"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 创建日期

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 变更单分组

	Partitions []*ProjectPartition `json:"Partitions,omitempty" name:"Partitions"`
}

type DeleteOperationInstructionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperationInstructionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationInstructionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUploadPackageTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件级别上传详情, 只有当Stage为Upload时有数据

		Components *ComponentPrepareTasks `json:"Components,omitempty" name:"Components"`
		// 任务状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 当前阶段, 分为Init(初始化阶段)/FindPath(查找路径)/Upload(上传阶段)

		Stage *string `json:"Stage,omitempty" name:"Stage"`
		// 当Status为Failed或者Error, 包含错误信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 上传物料步骤是否被跳过

		UploadProcedureSkipped *bool `json:"UploadProcedureSkipped,omitempty" name:"UploadProcedureSkipped"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUploadPackageTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUploadPackageTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExecutionSummary struct {
	// 任务ID

	Name *string `json:"Name,omitempty" name:"Name"`
	// 工具名字

	ToolName *string `json:"ToolName,omitempty" name:"ToolName"`
	// 运行结点

	Nodes []*string `json:"Nodes,omitempty" name:"Nodes"`
	// 当前状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 开始运行时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 操作人员

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DescribeProjectPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 变更单规划详情

		OperationInstructions []*OperationInstructionPlan `json:"OperationInstructions,omitempty" name:"OperationInstructions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeRunningStatus struct {
	// 结点

	Node *string `json:"Node,omitempty" name:"Node"`
	// 运行结果状态码

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// normal/abnormal

	CodeStatus *string `json:"CodeStatus,omitempty" name:"CodeStatus"`
	// 标准输出

	StdOut *string `json:"StdOut,omitempty" name:"StdOut"`
	// 标准错误

	StdErr *string `json:"StdErr,omitempty" name:"StdErr"`
}

type DescribeBizTreeHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机信息

		Hosts []*Host `json:"Hosts,omitempty" name:"Hosts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBizTreeHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBizTreeHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ToolDoc struct {
	// 文档内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 文档路径

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
}

type CreateToolImportURLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 上传链接

		URL *string `json:"URL,omitempty" name:"URL"`
		// 上传ID

		UploadId *string `json:"UploadId,omitempty" name:"UploadId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateToolImportURLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateToolImportURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCloudProductsRequest struct {
	*tchttp.BaseRequest
}

func (r *ListCloudProductsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCloudProductsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectRequest struct {
	*tchttp.BaseRequest

	// 项目ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParseInstructionComponentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件列表

		Components []*ComponentInstruction `json:"Components,omitempty" name:"Components"`
		// 物料机IP

		MaterialNode *string `json:"MaterialNode,omitempty" name:"MaterialNode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ParseInstructionComponentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ParseInstructionComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SheetResult struct {
	// 变更单ID

	SheetId *string `json:"SheetId,omitempty" name:"SheetId"`
	// 是否成功

	Success *bool `json:"Success,omitempty" name:"Success"`
	// 错误信息

	Error *string `json:"Error,omitempty" name:"Error"`
}

type ListToolsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 工具列表

		Tools []*Tool `json:"Tools,omitempty" name:"Tools"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListToolsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListToolsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DAGStep struct {
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 状态, Running/Ready

	Status *string `json:"Status,omitempty" name:"Status"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 参数

	Arguments *Arguments `json:"Arguments,omitempty" name:"Arguments"`
	// 运行在哪些宿主机上及状态

	Nodes []*StepNode `json:"Nodes,omitempty" name:"Nodes"`
	// 模板名字

	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
	// 准备阶段/实施阶段/验证阶段

	Stage *string `json:"Stage,omitempty" name:"Stage"`
	// 启动时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type CreateToolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 工具ID

		ToolId *string `json:"ToolId,omitempty" name:"ToolId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateToolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateToolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentInstanceYAMLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// YAML

		YAML *string `json:"YAML,omitempty" name:"YAML"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComponentInstanceYAMLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentInstanceYAMLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentPlanRequest struct {
	*tchttp.BaseRequest

	// 组件名

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *DescribeComponentPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationInstructionDAGResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 变更单详情及DAG

		OperationInstruction *OperationInstruction `json:"OperationInstruction,omitempty" name:"OperationInstruction"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationInstructionDAGResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationInstructionDAGResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeStageStep struct {
	// 命令名字, 当前只有command/pause两种, 分别表示执行shell以及暂停

	Command *string `json:"Command,omitempty" name:"Command"`
	// 命令列表

	Commands []*string `json:"Commands,omitempty" name:"Commands"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// Pause ID

	PauseID *string `json:"PauseID,omitempty" name:"PauseID"`
	// 参数描述

	Arguments []*StageStepArgument `json:"Arguments,omitempty" name:"Arguments"`
	// 执行策略

	Strategy *StepStrategy `json:"Strategy,omitempty" name:"Strategy"`
}

type SetFlowStatusRequest struct {
	*tchttp.BaseRequest

	// Flow名字

	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`
	// Flow状态, 支持retry/resume/kill/stop/pause/skip

	Status *string `json:"Status,omitempty" name:"Status"`
	// 要重试的阶段名字列表

	Stages []*string `json:"Stages,omitempty" name:"Stages"`
}

func (r *SetFlowStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetFlowStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetTaskInstanceStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetTaskInstanceStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetTaskInstanceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DirTreeNode struct {
	// 文件名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 结点类型，dir目录，file文件

	Type *string `json:"Type,omitempty" name:"Type"`
	// 完整路径

	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`
	// 文件类型，shell脚本，python脚本

	FileType *string `json:"FileType,omitempty" name:"FileType"`
	// 文件内容

	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`
	// 子结点

	Children []*DirTreeNode `json:"Children,omitempty" name:"Children"`
}

type Filter struct {
	// 参数名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤值

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 操作符, 不传默认like, 可选equal/gte/lte

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type Module struct {
	// UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 名字

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// Model UUID

	ModelUuid *string `json:"ModelUuid,omitempty" name:"ModelUuid"`
}
