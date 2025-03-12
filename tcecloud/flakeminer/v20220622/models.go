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

package v20220622

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type UpdateTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 参数名

	Name *string `json:"Name,omitempty" name:"Name"`
	// "equal", "like" 操作符

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 参数

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type MinerCondition struct {
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type MinerStage struct {
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 可用参数列表

	Parameters []*TemplateParameter `json:"Parameters,omitempty" name:"Parameters"`
	// 示例数据

	Example *string `json:"Example,omitempty" name:"Example"`
}

type Template struct {
	// 模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 模板输入

	Inputs *TemplateInputs `json:"Inputs,omitempty" name:"Inputs"`
	// 模板输出

	Outputs *TemplateOutputs `json:"Outputs,omitempty" name:"Outputs"`
	// dag定义

	DAG *DAGTemplate `json:"DAG,omitempty" name:"DAG"`
	// yunapi定义

	YunAPI *YunAPITemplate `json:"YunAPI,omitempty" name:"YunAPI"`
	// udf定义

	UDF *UDFTemplate `json:"UDF,omitempty" name:"UDF"`
	// remap定义

	Remap *RemapTemplate `json:"Remap,omitempty" name:"Remap"`
	// unwind定义

	Unwind *UnwindTemplate `json:"Unwind,omitempty" name:"Unwind"`
}

type DeleteTaskRequest struct {
	*tchttp.BaseRequest

	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MinerRunSpec struct {
	// 模板列表

	Templates []*Template `json:"Templates,omitempty" name:"Templates"`
	// 任务入口（第一个执行的任务）

	Entrypoint *string `json:"Entrypoint,omitempty" name:"Entrypoint"`
	// 模板参数

	Arguments *TemplateArguments `json:"Arguments,omitempty" name:"Arguments"`
	// flake表名

	FlakeTableName *string `json:"FlakeTableName,omitempty" name:"FlakeTableName"`
	// flake时间字段

	FlakeTimeFields []*string `json:"FlakeTimeFields,omitempty" name:"FlakeTimeFields"`
	// flake索引字段

	FlakeKeyFields []*string `json:"FlakeKeyFields,omitempty" name:"FlakeKeyFields"`
	// flake导入策略。可用值为 [Upsert, ReplaceAll]

	FlakeImportPolicy *string `json:"FlakeImportPolicy,omitempty" name:"FlakeImportPolicy"`
}

type RemapTemplate struct {
	// 映射列表

	Schema []*RemapSchema `json:"Schema,omitempty" name:"Schema"`
	// 循环对象表达式（循环执行的数据对象）

	WithParam *string `json:"WithParam,omitempty" name:"WithParam"`
	// 循环参数表达式（具体执行次数）

	WithSequence *Sequence `json:"WithSequence,omitempty" name:"WithSequence"`
	// 保持参数

	Holder *bool `json:"Holder,omitempty" name:"Holder"`
}

type DescribeTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务详情

		Task *MinerTask `json:"Task,omitempty" name:"Task"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DAGTemplate struct {
	// 子任务数组

	Tasks []*DAGTask `json:"Tasks,omitempty" name:"Tasks"`
}

type ListStagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// stage列表

		StageSet []*MinerStage `json:"StageSet,omitempty" name:"StageSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListStagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListStagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskRunRequest struct {
	*tchttp.BaseRequest

	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeTaskRunRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskRunRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTaskRunsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 执行实例列表

		RunSet []*MinerTaskRun `json:"RunSet,omitempty" name:"RunSet"`
		// 命名空间列表

		NamespaceSet []*string `json:"NamespaceSet,omitempty" name:"NamespaceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListTaskRunsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTaskRunsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DAGTask struct {
	// dag子任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 引用模板名称

	Template *string `json:"Template,omitempty" name:"Template"`
	// 引用的stage名称

	StageRef *MinerStageRef `json:"StageRef,omitempty" name:"StageRef"`
	// 传入参数列表

	Arguments *TemplateArguments `json:"Arguments,omitempty" name:"Arguments"`
	// 依赖的子任务名称

	Dependencies []*string `json:"Dependencies,omitempty" name:"Dependencies"`
	// 循环执行列表的表达式（循环的数组对象）

	WithParam *string `json:"WithParam,omitempty" name:"WithParam"`
	// 循环执行参数的表达式（具体循环次数）

	WithSequence *Sequence `json:"WithSequence,omitempty" name:"WithSequence"`
	// 条件语句

	When *string `json:"When,omitempty" name:"When"`
}

type TemplateOutputs struct {
	// 参数列表

	Parameters []*TemplateParameter `json:"Parameters,omitempty" name:"Parameters"`
	// 独立赋值表达式，输出由这个表达式定义

	Value *string `json:"Value,omitempty" name:"Value"`
}

type UnwindTemplate struct {
	// unwind的键

	Keys []*string `json:"Keys,omitempty" name:"Keys"`
	// 保持unwind字段，规避yunapibug

	Holder *bool `json:"Holder,omitempty" name:"Holder"`
}

type Sequence struct {
	// 表达式，循环的次数

	Count *string `json:"Count,omitempty" name:"Count"`
	// 循环起始值

	Start *int64 `json:"Start,omitempty" name:"Start"`
	// 循环结束值

	End *int64 `json:"End,omitempty" name:"End"`
}

type ExecuteTaskRequest struct {
	*tchttp.BaseRequest

	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 任务定义

	Spec *MinerRunSpec `json:"Spec,omitempty" name:"Spec"`
}

func (r *ExecuteTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExecuteTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTaskRunsRequest struct {
	*tchttp.BaseRequest

	// 任务的命名空间

	TaskNamespace *string `json:"TaskNamespace,omitempty" name:"TaskNamespace"`
	// 任务名

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 是否降序排列。默认否

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 排序依据

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 创建时间起始值，unix时间戳

	CreatedTimeStart *int64 `json:"CreatedTimeStart,omitempty" name:"CreatedTimeStart"`
	// 创建时间结束值，unix时间戳。

	CreatedTimeEnd *int64 `json:"CreatedTimeEnd,omitempty" name:"CreatedTimeEnd"`
}

func (r *ListTaskRunsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTaskRunsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest

	// 采集任务定义

	Task *MinerTask `json:"Task,omitempty" name:"Task"`
	// flake时序表主键

	FlakeTimeFields []*string `json:"FlakeTimeFields,omitempty" name:"FlakeTimeFields"`
	// flake非时序表主键

	FlakeKeyFields []*string `json:"FlakeKeyFields,omitempty" name:"FlakeKeyFields"`
}

func (r *CreateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MinerNodeStatus struct {
	// 节点名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 开始时间

	StartedAt *string `json:"StartedAt,omitempty" name:"StartedAt"`
	// 结束时间

	FinishedAt *string `json:"FinishedAt,omitempty" name:"FinishedAt"`
	// 执行信息，报错

	Message *string `json:"Message,omitempty" name:"Message"`
	// 状态

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 节点输出示例值

	OutputSample *string `json:"OutputSample,omitempty" name:"OutputSample"`
	// 节点自身输出示例值

	SelfOutputSample *string `json:"SelfOutputSample,omitempty" name:"SelfOutputSample"`
	// 节点输入示例值

	InputSample *string `json:"InputSample,omitempty" name:"InputSample"`
	// 节点级循环变量值

	TaskItemSample *string `json:"TaskItemSample,omitempty" name:"TaskItemSample"`
	// remap循环变量

	RemapItemSample *string `json:"RemapItemSample,omitempty" name:"RemapItemSample"`
	// 可选的unwind key

	AvailableUnwindKeys *string `json:"AvailableUnwindKeys,omitempty" name:"AvailableUnwindKeys"`
}

type UpdateTaskRequest struct {
	*tchttp.BaseRequest

	// 采集任务定义1

	Task *MinerTask `json:"Task,omitempty" name:"Task"`
}

func (r *UpdateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateInputs struct {
	// 参数列表

	Parameters []*TemplateParameter `json:"Parameters,omitempty" name:"Parameters"`
}

type UDFTemplate struct {
	// 脚本内容

	Source *string `json:"Source,omitempty" name:"Source"`
	// 保持字段

	Holder *bool `json:"Holder,omitempty" name:"Holder"`
}

type DeleteTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListStagesRequest struct {
	*tchttp.BaseRequest

	// 长度限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 是否降序排列

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 根据什么排列

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *ListStagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListStagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采集任务详情

		Task *MinerTask `json:"Task,omitempty" name:"Task"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskRequest struct {
	*tchttp.BaseRequest

	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskRunResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采集任务实例详情

		Run *MinerTaskRun `json:"Run,omitempty" name:"Run"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskRunResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTasksRequest struct {
	*tchttp.BaseRequest

	// 长度限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 是否降序排列

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 根据什么字段排序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *ListTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MinerStageRef struct {
	// stage名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type MinerTask struct {
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 任务调度crontab表达式

	Schedule *string `json:"Schedule,omitempty" name:"Schedule"`
	// 是否为草稿

	Draft *bool `json:"Draft,omitempty" name:"Draft"`
	// 任务创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 任务是否暂停执行

	Suspend *bool `json:"Suspend,omitempty" name:"Suspend"`
	// 历史成功任务保留数量

	SuccessfulTasksHistoryLimit *int64 `json:"SuccessfulTasksHistoryLimit,omitempty" name:"SuccessfulTasksHistoryLimit"`
	// 失败任务保留数量

	FailedTasksHistoryLimit *int64 `json:"FailedTasksHistoryLimit,omitempty" name:"FailedTasksHistoryLimit"`
	// 最后一次成功调度时间

	LastScheduleTime *string `json:"LastScheduleTime,omitempty" name:"LastScheduleTime"`
	// 采集任务定义

	Spec *MinerRunSpec `json:"Spec,omitempty" name:"Spec"`
}

type TemplateParameter struct {
	// 参数名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数的值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ListTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 任务列表

		TaskSet []*MinerTask `json:"TaskSet,omitempty" name:"TaskSet"`
		// 命名空间集合

		NamespaceSet []*string `json:"NamespaceSet,omitempty" name:"NamespaceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemapSchema struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 值表达式

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ExecuteTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 整体任务执行结果状态

		Phase *string `json:"Phase,omitempty" name:"Phase"`
		// 节点执行的结果

		NodeStatuses []*MinerNodeStatus `json:"NodeStatuses,omitempty" name:"NodeStatuses"`
		// 任务输出信息（报错）

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExecuteTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExecuteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MinerTaskRun struct {
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 任务定义

	Spec *MinerRunSpec `json:"Spec,omitempty" name:"Spec"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 执行状态

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 启动时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 完成时间

	CompletionTime *string `json:"CompletionTime,omitempty" name:"CompletionTime"`
	// 事件数组

	Conditions []*MinerCondition `json:"Conditions,omitempty" name:"Conditions"`
	// 节点执行状态数组

	NodeStatuses []*MinerNodeStatus `json:"NodeStatuses,omitempty" name:"NodeStatuses"`
}

type TemplateArguments struct {
	// 参数列表

	Parameters []*TemplateParameter `json:"Parameters,omitempty" name:"Parameters"`
}

type YunAPITemplate struct {
	// yunapi 方法

	Action *string `json:"Action,omitempty" name:"Action"`
	// yunapi版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// yunapi服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// region

	Region *string `json:"Region,omitempty" name:"Region"`
	// 超时时间。可用10s,10m,2h等

	Timeout *string `json:"Timeout,omitempty" name:"Timeout"`
	// 分页大小表达式

	PageSize *string `json:"PageSize,omitempty" name:"PageSize"`
	// 请求body，json格式

	Body *string `json:"Body,omitempty" name:"Body"`
}
