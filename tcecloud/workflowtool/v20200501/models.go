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

package v20200501

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type CreateWorkflowInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWorkflowInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWorkflowInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMultiTemplateRequest struct {
	*tchttp.BaseRequest

	// 批量创建的作业工具或者作业编排

	Templates []*Template `json:"Templates,omitempty" name:"Templates"`
}

func (r *CreateMultiTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMultiTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAgentListRequest struct {
	*tchttp.BaseRequest

	// 通用过滤器

	Filter []*TemplateFilter `json:"Filter,omitempty" name:"Filter"`
	// 分页查询的起始查询偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页查询的查询的数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryAgentListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryAgentListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBizListRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryBizListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBizListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartPeriodicTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartPeriodicTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartPeriodicTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopWorkflowInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopWorkflowInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopWorkflowInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePeriodicTaskRequest struct {
	*tchttp.BaseRequest

	// 定时任务详情

	PeriodicTask *TemplatePeriodicTask `json:"PeriodicTask,omitempty" name:"PeriodicTask"`
}

func (r *CreatePeriodicTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePeriodicTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopPeriodicTaskRequest struct {
	*tchttp.BaseRequest

	// 所属组件（前端可硬编码为workflow-tool）

	Component *string `json:"Component,omitempty" name:"Component"`
	// 定时任务UUID

	TimingUUID *string `json:"TimingUUID,omitempty" name:"TimingUUID"`
}

func (r *StopPeriodicTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopPeriodicTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateWorkflowDetail struct {
	// 变量信息

	Constants []*TemplateConstant `json:"Constants,omitempty" name:"Constants"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 实例UUID

	InstanceUUID *string `json:"InstanceUUID,omitempty" name:"InstanceUUID"`
	// 节点坐标信息

	Location []*TemplateNodeLocation `json:"Location,omitempty" name:"Location"`
	// 流程名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 节点列表

	Nodes []*TemplateNode `json:"Nodes,omitempty" name:"Nodes"`
	// 实例执行过程中的输出

	Output *string `json:"Output,omitempty" name:"Output"`
	// 父流程实例UUID

	ParentInstanceUUID *string `json:"ParentInstanceUUID,omitempty" name:"ParentInstanceUUID"`
	// 启动节点

	StartNode *string `json:"StartNode,omitempty" name:"StartNode"`
	// 启动时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 流程状态

	State *string `json:"State,omitempty" name:"State"`
}

type CreateWorkflowInstanceRequest struct {
	*tchttp.BaseRequest

	// 启动的流程实例对象

	Instance *TemplateInstance `json:"Instance,omitempty" name:"Instance"`
}

func (r *CreateWorkflowInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWorkflowInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPeriodicTaskResponseData struct {
	// 下一次调度时间

	NextScheduleTime []*string `json:"NextScheduleTime,omitempty" name:"NextScheduleTime"`
	// 定时任务列表

	PeriodicTasks []*TemplatePeriodicTask `json:"PeriodicTasks,omitempty" name:"PeriodicTasks"`
	// 查询结果的条数

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 查询的条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询的起始条数

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type TemplatePluginSpec struct {
	// 插件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 插件UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
}

type QueryPluginOutputSpecRepsonseData struct {
	// 插件出参规范列表

	Output []*TemplatePluginArgSpecInOut `json:"Output,omitempty" name:"Output"`
}

type QueryWorkflowTemplateData struct {
	// 获取数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 获取起始偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 获取的总条数

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 返回的模板列表

	Templates []*Template `json:"Templates,omitempty" name:"Templates"`
}

type GetAllPluginWithCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求的返回体内容

		Data *GetAllPluginWithCategoryResponseData `json:"Data,omitempty" name:"Data"`
		// 请求的响应内容

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAllPluginWithCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAllPluginWithCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAgentListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回体

		Data *QueryAgentListResponseData `json:"Data,omitempty" name:"Data"`
		// 请求返回消息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAgentListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryAgentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryWorkflowTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板查询返回体

		Data *QueryWorkflowTemplateData `json:"Data,omitempty" name:"Data"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryWorkflowTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryWorkflowTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopPeriodicTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopPeriodicTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopPeriodicTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyAgentListResponseData struct {
	// 存在于CMDB中的主机列表

	LegalHosts []*TemplateCMDBHost `json:"LegalHosts,omitempty" name:"LegalHosts"`
	// 不存在与CMDB中的主机IP列表

	IllegalHosts []*string `json:"IllegalHosts,omitempty" name:"IllegalHosts"`
}

type CreateMultiTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMultiTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMultiTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateNodeResult struct {
	// 节点执行过程中的信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 成功下发的总数（成功下发不等于下发的任务成功执行）

	Success *int64 `json:"Success,omitempty" name:"Success"`
	// 下发的机器总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type TemplateNodeRetry struct {
	// 是否开启重试功能

	CanRetry *bool `json:"CanRetry,omitempty" name:"CanRetry"`
	// 自动重试的时间间隔

	RetryInterval *int64 `json:"RetryInterval,omitempty" name:"RetryInterval"`
	// 自动重试的最大次数

	RetryTime *int64 `json:"RetryTime,omitempty" name:"RetryTime"`
}

type QueryWorkflowInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回体

		Data *QueryWorkflowInstanceResponseData `json:"Data,omitempty" name:"Data"`
		// 请求返回消息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryWorkflowInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryWorkflowInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryWorkflowInstanceDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回体

		Data *QueryWorkflowInstanceDetailResponseData `json:"Data,omitempty" name:"Data"`
		// 请求返回消息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryWorkflowInstanceDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryWorkflowInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePeriodicTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePeriodicTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePeriodicTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPluginOutputSpecResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回体

		Data *QueryPluginOutputSpecRepsonseData `json:"Data,omitempty" name:"Data"`
		// 请求返回消息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryPluginOutputSpecResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPluginOutputSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryWorkflowInstanceDetailRequest struct {
	*tchttp.BaseRequest

	// 所属组件（前端可硬编码为workflow-tool）

	Component *string `json:"Component,omitempty" name:"Component"`
	// 目标流程实例UUID

	InstanceUUID *string `json:"InstanceUUID,omitempty" name:"InstanceUUID"`
}

func (r *QueryWorkflowInstanceDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryWorkflowInstanceDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopWorkflowInstanceRequest struct {
	*tchttp.BaseRequest

	// 所属组件（前端可编码为workflow-tool）

	Component *string `json:"Component,omitempty" name:"Component"`
	// 通用过滤器

	Filter []*TemplateFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *StopWorkflowInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopWorkflowInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWorkflowTemplateRequest struct {
	*tchttp.BaseRequest

	// 被更新的作业工具/编排（需含TemplateUUID)

	Template *Template `json:"Template,omitempty" name:"Template"`
}

func (r *UpdateWorkflowTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateWorkflowTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateCMDBHost struct {
	// 内网IP

	InnerIP *string `json:"InnerIP,omitempty" name:"InnerIP"`
	// 主机状态

	HostStatus *string `json:"HostStatus,omitempty" name:"HostStatus"`
	// 主机IP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 主机所在地域

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 主机的序列号

	SN *string `json:"SN,omitempty" name:"SN"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机Agent状态

	AgentStatus *string `json:"AgentStatus,omitempty" name:"AgentStatus"`
}

type TemplatePeriodicTask struct {
	// 所属组件（前端可编码为workflow-tool）

	Component *string `json:"Component,omitempty" name:"Component"`
	// 用户ID

	UID *string `json:"UID,omitempty" name:"UID"`
	// 定时任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 流程模板UUID

	TemplateUUID *string `json:"TemplateUUID,omitempty" name:"TemplateUUID"`
	// 此定时任务启动时需提供的全局变量（即为其关联的目标的全局变量）

	Constants []*TemplateConstant `json:"Constants,omitempty" name:"Constants"`
	// 定时任务执行表达式

	Cron *string `json:"Cron,omitempty" name:"Cron"`
	// 是否在创建时启用定时任务

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 定时任务所关联的模板是否为作业工具

	WorkflowTool *bool `json:"WorkflowTool,omitempty" name:"WorkflowTool"`
	// 定时任务描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 定时任务UUID

	TimingUUID *string `json:"TimingUUID,omitempty" name:"TimingUUID"`
	// 状态

	State *string `json:"State,omitempty" name:"State"`
	// 上一次调度时间

	LastScheduleTime *string `json:"LastScheduleTime,omitempty" name:"LastScheduleTime"`
	// 上一次调度结果

	LastInstanceStatus *string `json:"LastInstanceStatus,omitempty" name:"LastInstanceStatus"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 定时任务ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 周期性任务执行总次数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 定时任务cron表达式格式

	Format *string `json:"Format,omitempty" name:"Format"`
}

type TemplatePluginCall struct {
	// 插件调用的创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 插件调用的目标IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 插件调用所属的实例UUID

	InstanceUUID *string `json:"InstanceUUID,omitempty" name:"InstanceUUID"`
	// 任务ID

	JobID *string `json:"JobID,omitempty" name:"JobID"`
	// 节点UUID

	NodeUUID *string `json:"NodeUUID,omitempty" name:"NodeUUID"`
	// 插件调用序号

	Number *int64 `json:"Number,omitempty" name:"Number"`
	// 插件执行的输出参数

	Output []*TemplatePluginArgInOut `json:"Output,omitempty" name:"Output"`
	// 插件调用的参数

	Param *string `json:"Param,omitempty" name:"Param"`
	// 插件UUID

	PluginUUID *string `json:"PluginUUID,omitempty" name:"PluginUUID"`
	// 插件调用响应

	Response *string `json:"Response,omitempty" name:"Response"`
	// 插件调用结果

	Result *TemplatePluginResult `json:"Result,omitempty" name:"Result"`
	// 插件重试策略

	Retry *TemplateNodeRetry `json:"Retry,omitempty" name:"Retry"`
	// 插件调用状态（Enum: [Running Success Failed Timeout]）

	State *string `json:"State,omitempty" name:"State"`
	// 插件调用的超时时间

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
}

type CheckWorkflowTemplateRequest struct {
	*tchttp.BaseRequest

	// 被校验的作业工具或者编排uuid数组

	TemplateUUIDs []*string `json:"TemplateUUIDs,omitempty" name:"TemplateUUIDs"`
	// 访问组件（前端可硬编码为 workflow-tool）

	Component *string `json:"Component,omitempty" name:"Component"`
}

func (r *CheckWorkflowTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckWorkflowTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryRunWorkflowNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryRunWorkflowNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryRunWorkflowNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPluginInputSpecRepsonseData struct {
	// 插件入参规范列表

	Input []*TemplatePluginArgSpecInOut `json:"Input,omitempty" name:"Input"`
}

type QueryWorkflowInstanceDetailResponseData struct {
	// 节点列表

	Nodes []*TemplateNode `json:"Nodes,omitempty" name:"Nodes"`
	// 节点坐标

	Location []*TemplateNodeLocation `json:"Location,omitempty" name:"Location"`
	// 变量信息

	Constants []*TemplateConstant `json:"Constants,omitempty" name:"Constants"`
	// 起始节点

	StartNode *string `json:"StartNode,omitempty" name:"StartNode"`
	// 终止时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 实例UUID

	InstanceUUID *string `json:"InstanceUUID,omitempty" name:"InstanceUUID"`
	// 实例名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 实例状态

	State *string `json:"State,omitempty" name:"State"`
	// 父流程实例UUID

	ParentInstanceUUID *string `json:"ParentInstanceUUID,omitempty" name:"ParentInstanceUUID"`
	// 实例执行输出

	Output *string `json:"Output,omitempty" name:"Output"`
}

type TemplateCMDBBizTree struct {
	// 业务唯一ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 业务ID

	BizID *string `json:"BizID,omitempty" name:"BizID"`
	// 业务名称

	BizName *string `json:"BizName,omitempty" name:"BizName"`
	// 业务UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 是否为资源池

	Default *bool `json:"Default,omitempty" name:"Default"`
	// 是否预设

	Preset *bool `json:"Preset,omitempty" name:"Preset"`
	// 是否展示

	Display *bool `json:"Display,omitempty" name:"Display"`
	// 当前为哪一个层级

	ModelUUID *string `json:"ModelUUID,omitempty" name:"ModelUUID"`
	// 业务所包含的集群列表

	SetList []*TemplateSetWithModule `json:"SetList,omitempty" name:"SetList"`
}

type TemplateBiz struct {
	// 业务唯一ID

	BizId *int64 `json:"BizId,omitempty" name:"BizId"`
	// 业务名称

	BizName *string `json:"BizName,omitempty" name:"BizName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 是否为资源池

	Default *bool `json:"Default,omitempty" name:"Default"`
	// 开发人员

	Developer *string `json:"Developer,omitempty" name:"Developer"`
	// 是否展示

	Display *bool `json:"Display,omitempty" name:"Display"`
	// 关注人员

	Follower *string `json:"Follower,omitempty" name:"Follower"`
	// 声明周期

	LifeCycle *string `json:"LifeCycle,omitempty" name:"LifeCycle"`
	// 维护人员

	Maintainer *string `json:"Maintainer,omitempty" name:"Maintainer"`
	// 更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 是否为预设

	Preset *bool `json:"Preset,omitempty" name:"Preset"`
	// 产品经理

	ProductManager *string `json:"ProductManager,omitempty" name:"ProductManager"`
	// 测试人员

	Tester *string `json:"Tester,omitempty" name:"Tester"`
	// 业务UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
}

type CreatePeriodicTasksRequest struct {
	*tchttp.BaseRequest

	// 定时任务详情数组

	PeriodicTasks []*TemplatePeriodicTask `json:"PeriodicTasks,omitempty" name:"PeriodicTasks"`
}

func (r *CreatePeriodicTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePeriodicTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAllPluginWithCategoryResponseData struct {
	// 插件规范数据及其类别

	Result []*TemplatePluginDataWithCategory `json:"Result,omitempty" name:"Result"`
}

type TemplateInstance struct {
	// 流程实例名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 所属组件（前端可硬编码为workflow-tool)

	Component *string `json:"Component,omitempty" name:"Component"`
	// 模板中定义的变量（需含value）

	Constants []*TemplateConstant `json:"Constants,omitempty" name:"Constants"`
	// 用户ID

	UID *string `json:"UID,omitempty" name:"UID"`
	// 模板UUID

	TemplateUUID *string `json:"TemplateUUID,omitempty" name:"TemplateUUID"`
	// 流程实例运行结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 流程实例UUID

	InstanceUUID *string `json:"InstanceUUID,omitempty" name:"InstanceUUID"`
	// 流程实例的父流程实例UUID

	ParentUUID *string `json:"ParentUUID,omitempty" name:"ParentUUID"`
	// 流程实例运行开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 流程实例的状态

	State *string `json:"State,omitempty" name:"State"`
	// 流程实例所隶属的定时任务UUID（如果有关联）

	TimingUUID *string `json:"TimingUUID,omitempty" name:"TimingUUID"`
	// 是否为作业工具任务

	WorkflowTool *bool `json:"WorkflowTool,omitempty" name:"WorkflowTool"`
	// 关联的定时任务名称

	TimingName *string `json:"TimingName,omitempty" name:"TimingName"`
	// 该流程实例关联的完整模板快照。

	TplSnapshot *Template `json:"TplSnapshot,omitempty" name:"TplSnapshot"`
}

type QueryBizListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回体

		Data *QueryBizListResponseData `json:"Data,omitempty" name:"Data"`
		// 请求返回消息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBizListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBizListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryWorkflowInstanceRequest struct {
	*tchttp.BaseRequest

	// 所属组件（前端可硬编码为workflow-tool）

	Component *string `json:"Component,omitempty" name:"Component"`
	// 分页查询的起始位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页查询的条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 通用过滤器

	Filter []*TemplateFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *QueryWorkflowInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryWorkflowInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyAgentListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求的返回内容

		Data *VerifyAgentListResponseData `json:"Data,omitempty" name:"Data"`
		// 请求返回消息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyAgentListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyAgentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBizTreeResponseData struct {
	// 业务模块树

	BizList *TemplateCMDBBizTree `json:"BizList,omitempty" name:"BizList"`
}

type CheckWorkflowTemplateResponseData struct {
	// 入参中对应的作业工具及编排是否存在

	Exists []*bool `json:"Exists,omitempty" name:"Exists"`
}

type CheckWorkflowTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回体

		Data *CheckWorkflowTemplateResponseData `json:"Data,omitempty" name:"Data"`
		// 请求返消息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckWorkflowTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckWorkflowTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWorkflowTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWorkflowTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWorkflowTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPeriodicTaskRequest struct {
	*tchttp.BaseRequest

	// 所属组件（前端可编码为workflow-tool）

	Component *string `json:"Component,omitempty" name:"Component"`
	// 通用过滤器

	Filter []*TemplateFilter `json:"Filter,omitempty" name:"Filter"`
	// 分页查询的条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页查询的起始页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *QueryPeriodicTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPeriodicTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryWorkflowNodeDetailResponseData struct {
	// 是否忽略节点

	Ignore *bool `json:"Ignore,omitempty" name:"Ignore"`
	// 节点UUID

	NodeUUID *string `json:"NodeUUID,omitempty" name:"NodeUUID"`
	// 节点类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 重试设置

	Retry *TemplateNodeRetry `json:"Retry,omitempty" name:"Retry"`
	// 节点插件

	Plugin *TemplatePlugin `json:"Plugin,omitempty" name:"Plugin"`
	// 超时时间

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	// 节点入度

	Incoming []*string `json:"Incoming,omitempty" name:"Incoming"`
	// 节点出度

	Outgoing []*string `json:"Outgoing,omitempty" name:"Outgoing"`
	// 变量信息

	Constants []*TemplateConstant `json:"Constants,omitempty" name:"Constants"`
	// 节点是否可失败手动跳过

	IsSkipped *bool `json:"IsSkipped,omitempty" name:"IsSkipped"`
	// 节点是否可失败手动重试

	IsRetried *bool `json:"IsRetried,omitempty" name:"IsRetried"`
	// 步骤名称

	StageName *string `json:"StageName,omitempty" name:"StageName"`
	// 条件信息

	Conditions []*TemplateCondition `json:"Conditions,omitempty" name:"Conditions"`
	// 模板UUID

	TemplateUUID *string `json:"TemplateUUID,omitempty" name:"TemplateUUID"`
	// 实例UUIID

	InstanceUUID *string `json:"InstanceUUID,omitempty" name:"InstanceUUID"`
	// 是否出错自动忽略

	ErrorIgnorable *bool `json:"ErrorIgnorable,omitempty" name:"ErrorIgnorable"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 插件调用情况

	Result *TemplateNodeResult `json:"Result,omitempty" name:"Result"`
	// 节点状态

	State *string `json:"State,omitempty" name:"State"`
	// 分支数量

	BranchNumber *int64 `json:"BranchNumber,omitempty" name:"BranchNumber"`
	// 子流程实例UUID

	SubInstanceUUID *string `json:"SubInstanceUUID,omitempty" name:"SubInstanceUUID"`
	// 通知设置

	Notify *TemplateNotify `json:"Notify,omitempty" name:"Notify"`
	// 节点及执行记录

	Nodes []*TemplatePluginCall `json:"Nodes,omitempty" name:"Nodes"`
	// 节点对应的输入参数

	Input []*TemplatePluginArgInOut `json:"Input,omitempty" name:"Input"`
	// 节点执行总次数

	ExecutionCnt *int64 `json:"ExecutionCnt,omitempty" name:"ExecutionCnt"`
}

type Template struct {
	//  分类

	Category *string `json:"Category,omitempty" name:"Category"`
	// 所属组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 是否删除

	Deleted *bool `json:"Deleted,omitempty" name:"Deleted"`
	// 模板描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 管理人员

	Managers []*string `json:"Managers,omitempty" name:"Managers"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否私有

	Private *bool `json:"Private,omitempty" name:"Private"`
	// 是否为系统预设模板

	System *bool `json:"System,omitempty" name:"System"`
	// 模板UUID

	TemplateUUID *string `json:"TemplateUUID,omitempty" name:"TemplateUUID"`
	// 用户UID

	UID *string `json:"UID,omitempty" name:"UID"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 流程详情

	WorkflowDetail *TemplateWorkflowDetail `json:"WorkflowDetail,omitempty" name:"WorkflowDetail"`
	// 是否为作业工具

	WorkflowTool *bool `json:"WorkflowTool,omitempty" name:"WorkflowTool"`
	// 作业编排通知配置

	Notify *TemplateNotify `json:"Notify,omitempty" name:"Notify"`
}

type TemplateConstant struct {
	// 变量描述信息

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 是否允许变量输出

	Exposed *bool `json:"Exposed,omitempty" name:"Exposed"`
	// 变量产生的主机 IP

	Host *string `json:"Host,omitempty" name:"Host"`
	// 变量关联的实例UUID

	InstanceUUID *string `json:"InstanceUUID,omitempty" name:"InstanceUUID"`
	// 变量的 Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 变量是否应该在流程启动时进行初始化

	MustInit *bool `json:"MustInit,omitempty" name:"MustInit"`
	// 该插件出入参数所对应的节点 UUIDNodeUUID

	NodeUUID *string `json:"NodeUUID,omitempty" name:"NodeUUID"`
	// 变量被引用的次数

	ReferenceCnt *uint64 `json:"ReferenceCnt,omitempty" name:"ReferenceCnt"`
	// 变量属于全局变量还是常量（Enum: [global local]）

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 变量的来源

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 变量值的类型（Enum: [string integer boolean complex host]）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 变量值。若该变量（比如插件的入参）引用其他变量，则依据其引用的变量的类型来决定其占位符的形式。非字符串类型需要在双括号内侧加中划线。若该变量为全局变量，则其Value省略。

	Value *string `json:"Value,omitempty" name:"Value"`
	// 变量名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type TemplatePlugin struct {
	// 插件用到的数据（json数据格式）。其中，若入参为变量，则依据其本身的类型或者引用的变量的类型来决定其占位符的形式。非字符串类型需要在双括号内侧加中划线。。

	Data *string `json:"Data,omitempty" name:"Data"`
	// 插件UUID

	PluginUUID *string `json:"PluginUUID,omitempty" name:"PluginUUID"`
}

type TemplateNotify struct {
	// 通知内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 是否开启通知

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 通知标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 消息订阅ID

	TopicID *int64 `json:"TopicID,omitempty" name:"TopicID"`
	// 通知的事件集合（枚举：SMS、Mail 和 RTX）

	Channels []*string `json:"Channels,omitempty" name:"Channels"`
	// 通知接收者集合（username数组）

	Receivers []*string `json:"Receivers,omitempty" name:"Receivers"`
	// 通知的条件（枚举：OnFailure、OnSuccess、Always 和 Never）

	Cond *string `json:"Cond,omitempty" name:"Cond"`
}

type CreateWorkflowTemplatesRequest struct {
	*tchttp.BaseRequest

	// 需要被创建的作业工具或者编排的数组

	Templates []*Template `json:"Templates,omitempty" name:"Templates"`
}

func (r *CreateWorkflowTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWorkflowTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWorkflowTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWorkflowTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWorkflowTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBizTreeRequest struct {
	*tchttp.BaseRequest

	// 业务UUID

	BizUUID *string `json:"BizUUID,omitempty" name:"BizUUID"`
}

func (r *QueryBizTreeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBizTreeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPluginInputSpecRequest struct {
	*tchttp.BaseRequest

	// 插件UUID

	PluginUUID *string `json:"PluginUUID,omitempty" name:"PluginUUID"`
}

func (r *QueryPluginInputSpecRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPluginInputSpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPluginInputSpecResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回体

		Data *QueryPluginInputSpecRepsonseData `json:"Data,omitempty" name:"Data"`
		// 请求返回内容

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryPluginInputSpecResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPluginInputSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartPeriodicTaskRequest struct {
	*tchttp.BaseRequest

	// 所属组件（前端可编码为workflow-tool）

	Component *string `json:"Component,omitempty" name:"Component"`
	// 定时任务UUID

	TimingUUID *string `json:"TimingUUID,omitempty" name:"TimingUUID"`
}

func (r *StartPeriodicTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartPeriodicTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplatePluginArgSpecInOut struct {
	// 插件参数名称

	Key *string `json:"Key,omitempty" name:"Key"`
	// 参数描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 参数类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 参数是否在前端展示

	Display *bool `json:"Display,omitempty" name:"Display"`
}

type QueryBizListResponseData struct {
	// 总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 业务列表

	BizList []*TemplateBiz `json:"BizList,omitempty" name:"BizList"`
}

type TemplateFilter struct {
	// 过滤项名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤项对应的值（json格式）

	Value []*string `json:"Value,omitempty" name:"Value"`
}

type QueryWorkflowTemplateRequest struct {
	*tchttp.BaseRequest

	// 所属组件（前端可直接填 workflow-tool）

	Component *string `json:"Component,omitempty" name:"Component"`
	// 查询过滤器的数组

	Filter []*TemplateFilter `json:"Filter,omitempty" name:"Filter"`
	// 分页查询获取的数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页查询获取的起始偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *QueryWorkflowTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryWorkflowTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWorkflowTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateWorkflowTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateWorkflowTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWorkflowTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWorkflowTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWorkflowTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePeriodicTaskRequest struct {
	*tchttp.BaseRequest

	// 需更新的定时任务

	PeriodicTask *TemplatePeriodicTask `json:"PeriodicTask,omitempty" name:"PeriodicTask"`
}

func (r *UpdatePeriodicTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePeriodicTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePeriodicTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePeriodicTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePeriodicTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAllPluginWithCategoryRequest struct {
	*tchttp.BaseRequest

	// 当前用户

	UID *string `json:"UID,omitempty" name:"UID"`
}

func (r *GetAllPluginWithCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAllPluginWithCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryWorkflowNodeDetailRequest struct {
	*tchttp.BaseRequest

	// 所属组件（前端可编码为workflow-tool）

	Component *string `json:"Component,omitempty" name:"Component"`
	// 目标流程实例UUID

	InstanceUUID *string `json:"InstanceUUID,omitempty" name:"InstanceUUID"`
	// 目标节点UUID

	NodeUUID *string `json:"NodeUUID,omitempty" name:"NodeUUID"`
	// 查询节点第几次的执行记录（从1开始）

	Number *uint64 `json:"Number,omitempty" name:"Number"`
}

func (r *QueryWorkflowNodeDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryWorkflowNodeDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAgentListResponseData struct {
	// 主机列表大小

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 主机列表

	AgentList []*TemplateCMDBHost `json:"AgentList,omitempty" name:"AgentList"`
}

type TemplateModuleWithModelUUId struct {
	// 模块唯一ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 模块ID

	ModuleID *string `json:"ModuleID,omitempty" name:"ModuleID"`
	// 模块名称

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 模块UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 模块所属集群UUID

	SetUUID *string `json:"SetUUID,omitempty" name:"SetUUID"`
	// 模块所属业务UUID

	BizUUID *string `json:"BizUUID,omitempty" name:"BizUUID"`
	// 是否为资源池

	Default *bool `json:"Default,omitempty" name:"Default"`
	// 是否预设

	Preset *bool `json:"Preset,omitempty" name:"Preset"`
	// 模块等级

	ModuleLevel *string `json:"ModuleLevel,omitempty" name:"ModuleLevel"`
	// 模块类别

	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`
	// 服务等级

	ServiceLevel *string `json:"ServiceLevel,omitempty" name:"ServiceLevel"`
	// 当前为哪一个等级i

	ModelUUID *string `json:"ModelUUID,omitempty" name:"ModelUUID"`
}

type CreateWorkflowTemplateRequest struct {
	*tchttp.BaseRequest

	// 需要创建的作业工具或作业编排

	Template *Template `json:"Template,omitempty" name:"Template"`
}

func (r *CreateWorkflowTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWorkflowTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBizTreeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回消息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 请求返回的请求体内容

		Data *QueryBizTreeResponseData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBizTreeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBizTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateNodeLocation struct {
	// 节点类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 节点 UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 节点横坐标

	X1 *float64 `json:"X1,omitempty" name:"X1"`
	// 节点纵坐标

	Y1 *float64 `json:"Y1,omitempty" name:"Y1"`
}

type TemplatePluginArgInOut struct {
	// 参数的key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 参数的value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeletePeriodicTaskRequest struct {
	*tchttp.BaseRequest

	// 所属组件（前端可编码为workflow-tool）

	Component *string `json:"Component,omitempty" name:"Component"`
	// 定时任务UUID

	TimingUUID *string `json:"TimingUUID,omitempty" name:"TimingUUID"`
}

func (r *DeletePeriodicTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePeriodicTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SkipWorkflowNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SkipWorkflowNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SkipWorkflowNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyAgentListRequest struct {
	*tchttp.BaseRequest

	// 需校验的用户手动输入的IP列表

	IPList []*string `json:"IPList,omitempty" name:"IPList"`
}

func (r *VerifyAgentListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyAgentListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateCondition struct {
	// 该表达式关联的分支的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 该表达式关联的分支的下一个节点 UUID

	NodeUUID *string `json:"NodeUUID,omitempty" name:"NodeUUID"`
	// 条件表达式

	Evaluate *string `json:"Evaluate,omitempty" name:"Evaluate"`
}

type TemplatePluginDataWithCategory struct {
	// 插件规范数据

	PluginData []*TemplatePluginSpec `json:"PluginData,omitempty" name:"PluginData"`
	// 插件所属类别

	Category *string `json:"Category,omitempty" name:"Category"`
}

type TemplatePluginResult struct {
	// 任务状态码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 扩展字段

	External *string `json:"External,omitempty" name:"External"`
	// 消息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 标准错误

	StdErr *string `json:"StdErr,omitempty" name:"StdErr"`
	// 标准输出

	StdOut *string `json:"StdOut,omitempty" name:"StdOut"`
}

type TemplateSetWithModule struct {
	// 集群唯一ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 集群ID

	SetID *string `json:"SetID,omitempty" name:"SetID"`
	// 集群名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 集群UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 集群所属业务UUID

	BizUUID *string `json:"BizUUID,omitempty" name:"BizUUID"`
	// 是否为资源池

	Default *bool `json:"Default,omitempty" name:"Default"`
	// 是否预设

	Preset *bool `json:"Preset,omitempty" name:"Preset"`
	// 集群类别

	SetType *string `json:"SetType,omitempty" name:"SetType"`
	// 服务等级

	ServiceLevel *string `json:"ServiceLevel,omitempty" name:"ServiceLevel"`
	// 当前为哪一个层级

	ModelUUID *string `json:"ModelUUID,omitempty" name:"ModelUUID"`
	// 当前集群所包含的模块集合

	ModuleList []*TemplateModuleWithModelUUId `json:"ModuleList,omitempty" name:"ModuleList"`
}

type QueryPeriodicTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回体

		Data *QueryPeriodicTaskResponseData `json:"Data,omitempty" name:"Data"`
		// 请求返回消息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryPeriodicTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPeriodicTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPluginOutputSpecRequest struct {
	*tchttp.BaseRequest

	// 插件UUID

	PluginUUID *string `json:"PluginUUID,omitempty" name:"PluginUUID"`
}

func (r *QueryPluginOutputSpecRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPluginOutputSpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePeriodicTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePeriodicTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePeriodicTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryWorkflowInstanceResponseData struct {
	// 查询数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询的总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 实例列表

	Instances []*TemplateInstance `json:"Instances,omitempty" name:"Instances"`
}

type CreatePeriodicTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePeriodicTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePeriodicTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWorkflowTemplateRequest struct {
	*tchttp.BaseRequest

	// 访问的组件（前端可硬编码为workflow-tool）

	Component *string `json:"Component,omitempty" name:"Component"`
	// 通用过滤条件

	Filter []*TemplateFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DeleteWorkflowTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWorkflowTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryWorkflowNodeDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求的返回体

		Data *QueryWorkflowNodeDetailResponseData `json:"Data,omitempty" name:"Data"`
		// 请求返回消息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryWorkflowNodeDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryWorkflowNodeDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryRunWorkflowNodeRequest struct {
	*tchttp.BaseRequest

	// 所属组件（前端可编码为workflow-tool）

	Component *string `json:"Component,omitempty" name:"Component"`
	// 实例UUID

	InstanceUUID *string `json:"InstanceUUID,omitempty" name:"InstanceUUID"`
	// 节点UUID

	NodeUUID *string `json:"NodeUUID,omitempty" name:"NodeUUID"`
}

func (r *RetryRunWorkflowNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryRunWorkflowNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SkipWorkflowNodeRequest struct {
	*tchttp.BaseRequest

	// 所属组件（前端可编码为workflow-tool）

	Component *string `json:"Component,omitempty" name:"Component"`
	// 实例UUID

	InstanceUUID *string `json:"InstanceUUID,omitempty" name:"InstanceUUID"`
	// 节点UUID

	NodeUUID *string `json:"NodeUUID,omitempty" name:"NodeUUID"`
}

func (r *SkipWorkflowNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SkipWorkflowNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateNode struct {
	// 分支数量

	BranchNumber *int64 `json:"BranchNumber,omitempty" name:"BranchNumber"`
	// 变量信息

	Constants []*TemplateConstant `json:"Constants,omitempty" name:"Constants"`
	// 插件调用详情

	Detail []*TemplatePluginCall `json:"Detail,omitempty" name:"Detail"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 是否失败自动忽略

	ErrorIgnorable *bool `json:"ErrorIgnorable,omitempty" name:"ErrorIgnorable"`
	// 忽略节点（无用）

	Ignore *bool `json:"Ignore,omitempty" name:"Ignore"`
	// 节点入度

	Incoming []*string `json:"Incoming,omitempty" name:"Incoming"`
	// 实例UUID

	InstanceUUID *string `json:"InstanceUUID,omitempty" name:"InstanceUUID"`
	// 节点是否可以失败手动重试

	IsRetried *bool `json:"IsRetried,omitempty" name:"IsRetried"`
	// 节点是否失败手动跳过

	IsSkipped *bool `json:"IsSkipped,omitempty" name:"IsSkipped"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 节点UUID

	NodeUUID *string `json:"NodeUUID,omitempty" name:"NodeUUID"`
	// 通知设置

	Notify *TemplateNotify `json:"Notify,omitempty" name:"Notify"`
	// 节点出度

	Outgoing []*string `json:"Outgoing,omitempty" name:"Outgoing"`
	// 插件

	Plugin *TemplatePlugin `json:"Plugin,omitempty" name:"Plugin"`
	// 插件调用情况

	Result *TemplateNodeResult `json:"Result,omitempty" name:"Result"`
	// 节点重试策略

	Retry *TemplateNodeRetry `json:"Retry,omitempty" name:"Retry"`
	// 步骤名称

	StageName *string `json:"StageName,omitempty" name:"StageName"`
	// 节点开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 节点状态

	State *string `json:"State,omitempty" name:"State"`
	// 子流程节点关联的实例UUID

	SubInstanceUUID *string `json:"SubInstanceUUID,omitempty" name:"SubInstanceUUID"`
	// 模板UUID

	TemplateUUID *string `json:"TemplateUUID,omitempty" name:"TemplateUUID"`
	// 节点超时时间

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	// 节点类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 条件信息

	Conditions []*TemplateCondition `json:"Conditions,omitempty" name:"Conditions"`
}
