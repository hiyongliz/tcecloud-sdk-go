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

package v20210412

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type CreateSceneDrillLibraryRequest struct {
	*tchttp.BaseRequest

	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 演练编排节点列表

	DrillNodeList []*SceneNodeContent `json:"DrillNodeList,omitempty" name:"DrillNodeList"`
	// biz名称

	BizName *string `json:"BizName,omitempty" name:"BizName"`
}

func (r *CreateSceneDrillLibraryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSceneDrillLibraryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditExperienceLibraryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditExperienceLibraryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditExperienceLibraryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditSceneDrillLibraryRequest struct {
	*tchttp.BaseRequest

	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// DrillNodeList

	DrillNodeList []*SceneNodeContent `json:"DrillNodeList,omitempty" name:"DrillNodeList"`
}

func (r *EditSceneDrillLibraryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditSceneDrillLibraryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTargetTypeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTargetTypeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTargetTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExecuteSceneDrillLibraryRequest struct {
	*tchttp.BaseRequest

	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *ExecuteSceneDrillLibraryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExecuteSceneDrillLibraryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComponentDrillLibraryNode struct {
	// 工作流模板

	WorkflowTemplateUuid *string `json:"WorkflowTemplateUuid,omitempty" name:"WorkflowTemplateUuid"`
	// 监控指标

	MonitorMetrics []*MonitorMetric `json:"MonitorMetrics,omitempty" name:"MonitorMetrics"`
	// 等待更新

	TerminateCondition *TerminateCondition `json:"TerminateCondition,omitempty" name:"TerminateCondition"`
	// 监控目标

	MonitorMetricsObject *DrillObject `json:"MonitorMetricsObject,omitempty" name:"MonitorMetricsObject"`
	// 终止目标

	TerminateConditionObject *DrillObject `json:"TerminateConditionObject,omitempty" name:"TerminateConditionObject"`
	// 演练目标

	DrillObject *DrillObject `json:"DrillObject,omitempty" name:"DrillObject"`
	// 演练目标类型

	DrillObjectType *string `json:"DrillObjectType,omitempty" name:"DrillObjectType"`
	// 节点列表

	DrillNodeList []*DrillNode `json:"DrillNodeList,omitempty" name:"DrillNodeList"`
}

type MonitorMetric struct {
	// 监控指标类型

	TcsType *string `json:"TcsType,omitempty" name:"TcsType"`
	// 监控指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 监控指标中文名称

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 监控指标产品

	TcsProduct *string `json:"TcsProduct,omitempty" name:"TcsProduct"`
	// 可选标签

	Option []*OptionKV `json:"Option,omitempty" name:"Option"`
	// 目标标签key

	TargetKey []*string `json:"TargetKey,omitempty" name:"TargetKey"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 监控查询参数

	Query *string `json:"Query,omitempty" name:"Query"`
}

type DescribeSceneDrillLibraryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSceneDrillLibraryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSceneDrillLibraryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetExperienceLibraryCountRequest struct {
	*tchttp.BaseRequest
}

func (r *GetExperienceLibraryCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetExperienceLibraryCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBaseInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *GetBaseInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBaseInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExperienceLibraryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExperienceLibraryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExperienceLibraryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetExperienceLibraryCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetExperienceLibraryCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetExperienceLibraryCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSubTypeListRequest struct {
	*tchttp.BaseRequest

	// MainTypeName

	MainTypeName *string `json:"MainTypeName,omitempty" name:"MainTypeName"`
	// 目标类型

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
}

func (r *GetSubTypeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubTypeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentDrillLibraryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComponentDrillLibraryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentDrillLibraryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComponentDrillLibraryListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComponentDrillLibraryListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComponentDrillLibraryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSceneDrillLibraryRequest struct {
	*tchttp.BaseRequest

	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeSceneDrillLibraryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSceneDrillLibraryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReExecuteTaskRequest struct {
	*tchttp.BaseRequest

	// uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ReExecuteTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReExecuteTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DrillConfig struct {
	// Param

	Param []*ExperienceLibraryParam `json:"Param,omitempty" name:"Param"`
	// Common

	Common *DrillConfigCommon `json:"Common,omitempty" name:"Common"`
}

type SceneNodeContent struct {
	// 节点ID

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 入度列表

	Incoming []*string `json:"Incoming,omitempty" name:"Incoming"`
	// 出度列表

	Outgoing []*string `json:"Outgoing,omitempty" name:"Outgoing"`
	// 节点状态

	NodeStatus *string `json:"NodeStatus,omitempty" name:"NodeStatus"`
	// 引用的组件演练库ID

	ReferencedCptId *int64 `json:"ReferencedCptId,omitempty" name:"ReferencedCptId"`
	// 公共配置

	CommonConfig *CommonConfig `json:"CommonConfig,omitempty" name:"CommonConfig"`
	// 控制配置

	DrillConfig *DrillConfig `json:"DrillConfig,omitempty" name:"DrillConfig"`
	// 待更新类型

	ComponentDrillLibraryNode *ComponentDrillLibraryNode `json:"ComponentDrillLibraryNode,omitempty" name:"ComponentDrillLibraryNode"`
	// 深拷贝后保存最新配置的演练库ID

	ActualCptId *int64 `json:"ActualCptId,omitempty" name:"ActualCptId"`
	// 应用的组件演练库名称

	ReferencedCptName *string `json:"ReferencedCptName,omitempty" name:"ReferencedCptName"`
}

type TimeFilter struct {
	// 筛选时间类型

	TimeType *string `json:"TimeType,omitempty" name:"TimeType"`
	// 筛选起始区间，按当前时区，必须和TimeType搭配

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 筛选结束区间，按当前时区，必须和TimeType搭配

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 排序字段，取值范围：StartTime，EndTime

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 排序顺序，取值范围：Descending, Ascending。分别表示最新在前，最新在后

	Order *string `json:"Order,omitempty" name:"Order"`
}

type CreateSceneDrillLibraryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSceneDrillLibraryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSceneDrillLibraryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSubTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSubTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReserveComponentDrillResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReserveComponentDrillResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReserveComponentDrillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonConfigCommon struct {
	// Continue表示该经验库执行失败并且恢复后，继续执行同队列中下一个经验库；
	//
	// Terminate表示该经验库执行失败并且恢复后，不继续执行同队列中下一个经验库

	ExceptionHandler *string `json:"ExceptionHandler,omitempty" name:"ExceptionHandler"`
}

type GetTaskInfoRequest struct {
	*tchttp.BaseRequest

	// uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

func (r *GetTaskInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTaskInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTaskListRequest struct {
	*tchttp.BaseRequest

	// 任务状态

	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 演练库名称

	DrillName *string `json:"DrillName,omitempty" name:"DrillName"`
	// 演练库类型

	DrillType *string `json:"DrillType,omitempty" name:"DrillType"`
	// 时间筛选参数

	TimeFilter *TimeFilter `json:"TimeFilter,omitempty" name:"TimeFilter"`
	// 分页起始位

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页记录数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 父任务uuid

	ParentUuid *string `json:"ParentUuid,omitempty" name:"ParentUuid"`
	// 可见级别

	VisibleLevel *string `json:"VisibleLevel,omitempty" name:"VisibleLevel"`
}

func (r *GetTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReserveComponentDrillRequest struct {
	*tchttp.BaseRequest

	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 下次演练的预约时间

	DrillReservation *string `json:"DrillReservation,omitempty" name:"DrillReservation"`
}

func (r *ReserveComponentDrillRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReserveComponentDrillRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExperienceLibraryParam struct {
	// 参数键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 参数单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 参数名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 变量类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是否为必填参数

	Required *bool `json:"Required,omitempty" name:"Required"`
}

type GetMainTypeListRequest struct {
	*tchttp.BaseRequest

	// 目标类型

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
}

func (r *GetMainTypeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMainTypeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditExperienceLibraryRequest struct {
	*tchttp.BaseRequest

	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// MainType

	MainType *string `json:"MainType,omitempty" name:"MainType"`
	// SubType

	SubType *string `json:"SubType,omitempty" name:"SubType"`
	// Remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 演练配置

	DrillConfig *DrillConfig `json:"DrillConfig,omitempty" name:"DrillConfig"`
	// 恢复配置

	RecoverConfig *RecoverConfig `json:"RecoverConfig,omitempty" name:"RecoverConfig"`
	// 公共配置

	CommonConfig *CommonConfig `json:"CommonConfig,omitempty" name:"CommonConfig"`
}

func (r *EditExperienceLibraryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditExperienceLibraryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOverviewInfoRequest struct {
	*tchttp.BaseRequest

	// StartTime

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *GetOverviewInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOverviewInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSceneDrillLibraryListRequest struct {
	*tchttp.BaseRequest

	// 演练库名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 时间筛选参数

	TimeFilter *TimeFilter `json:"TimeFilter,omitempty" name:"TimeFilter"`
	// 分页起始位

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页记录数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// biz名称

	BizName *string `json:"BizName,omitempty" name:"BizName"`
}

func (r *GetSceneDrillLibraryListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSceneDrillLibraryListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteComponentDrillLibraryRequest struct {
	*tchttp.BaseRequest

	// 经验库名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 演练库id，和演练库名称二选一，同时使用以id为准，可选

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteComponentDrillLibraryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteComponentDrillLibraryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOpRecordListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOpRecordListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOpRecordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReserveSceneDrillResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReserveSceneDrillResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReserveSceneDrillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecoverConfig struct {
	// Param

	Param []*ExperienceLibraryParam `json:"Param,omitempty" name:"Param"`
	// Common

	Common *RecoverConfigCommon `json:"Common,omitempty" name:"Common"`
}

type ExecuteSceneDrillLibraryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExecuteSceneDrillLibraryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExecuteSceneDrillLibraryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditComponentDrillLibraryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditComponentDrillLibraryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditComponentDrillLibraryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOverviewInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOverviewInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOverviewInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTaskInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTaskInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentDrillLibraryRequest struct {
	*tchttp.BaseRequest

	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeComponentDrillLibraryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentDrillLibraryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSubTaskListRequest struct {
	*tchttp.BaseRequest

	// 所属任务uuid

	TaskUuid *string `json:"TaskUuid,omitempty" name:"TaskUuid"`
	// 节点id

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// 容器模块名

	Module *string `json:"Module,omitempty" name:"Module"`
	// 机器sn

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 机器内网ip

	InnerIp *string `json:"InnerIp,omitempty" name:"InnerIp"`
	// 机器名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 演练状态

	DrillStatus *string `json:"DrillStatus,omitempty" name:"DrillStatus"`
	// 演练阶段

	Stage *string `json:"Stage,omitempty" name:"Stage"`
	// 时间筛选参数

	TimeFilter *TimeFilter `json:"TimeFilter,omitempty" name:"TimeFilter"`
	// 分页起始

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 记录数量限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetSubTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSceneDrillLibraryBizNamesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSceneDrillLibraryBizNamesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSceneDrillLibraryBizNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateComponentDrillLibraryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateComponentDrillLibraryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateComponentDrillLibraryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSceneDrillLibraryBizNamesRequest struct {
	*tchttp.BaseRequest
}

func (r *GetSceneDrillLibraryBizNamesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSceneDrillLibraryBizNamesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DrillObjectModule struct {
	// 业务

	Biz *string `json:"Biz,omitempty" name:"Biz"`
	// 集群

	Set *string `json:"Set,omitempty" name:"Set"`
	// 模块

	Module *string `json:"Module,omitempty" name:"Module"`
	// Region

	Region *string `json:"Region,omitempty" name:"Region"`
	// Zone

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 模块名称

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

type ExecuteComponentDrillLibraryRequest struct {
	*tchttp.BaseRequest

	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *ExecuteComponentDrillLibraryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExecuteComponentDrillLibraryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DrillNode struct {
	// 节点id

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// Incoming

	Incoming []*string `json:"Incoming,omitempty" name:"Incoming"`
	// Outgoing

	Outgoing []*string `json:"Outgoing,omitempty" name:"Outgoing"`
	// ExperienceLibraryId

	ExperienceLibraryId *uint64 `json:"ExperienceLibraryId,omitempty" name:"ExperienceLibraryId"`
	// DrillConfig

	DrillConfig *DrillConfig `json:"DrillConfig,omitempty" name:"DrillConfig"`
	// RecoverConfig

	RecoverConfig *RecoverConfig `json:"RecoverConfig,omitempty" name:"RecoverConfig"`
	// CommonConfig

	CommonConfig *CommonConfig `json:"CommonConfig,omitempty" name:"CommonConfig"`
	// 故障类型

	FaultType *string `json:"FaultType,omitempty" name:"FaultType"`
}

type OptionKV struct {
	// 标签英文名

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 标签中文名

	Text *string `json:"Text,omitempty" name:"Text"`
}

type GetComponentDrillLibraryListRequest struct {
	*tchttp.BaseRequest

	// 演练库名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 时间筛选参数

	TimeFilter *TimeFilter `json:"TimeFilter,omitempty" name:"TimeFilter"`
	// 分页起始位

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页记录数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// biz名称

	BizName *string `json:"BizName,omitempty" name:"BizName"`
}

func (r *GetComponentDrillLibraryListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComponentDrillLibraryListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReExecuteTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReExecuteTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReExecuteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CmdbNodes struct {
	// 节点id

	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`
	// uuid

	NodeUuid *string `json:"NodeUuid,omitempty" name:"NodeUuid"`
	// 节点name

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 关联model

	RelationModel *string `json:"RelationModel,omitempty" name:"RelationModel"`
	// cmdb树路径

	FrontNodeUuidList *string `json:"FrontNodeUuidList,omitempty" name:"FrontNodeUuidList"`
	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
}

type CommonConfig struct {
	// Common

	Common *CommonConfigCommon `json:"Common,omitempty" name:"Common"`
}

type DeleteSceneDrillLibraryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSceneDrillLibraryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSceneDrillLibraryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecoverConfigCommon struct {
	// 执行后等待时间

	WaitBeforeRun *uint64 `json:"WaitBeforeRun,omitempty" name:"WaitBeforeRun"`
	// 执行后等待时间单位

	WaitBeforeRunUnit *string `json:"WaitBeforeRunUnit,omitempty" name:"WaitBeforeRunUnit"`
	// 执行后等待时间

	WaitAfterRun *uint64 `json:"WaitAfterRun,omitempty" name:"WaitAfterRun"`
	// 执行后等待时间单位

	WaitAfterRunUnit *string `json:"WaitAfterRunUnit,omitempty" name:"WaitAfterRunUnit"`
	// 最小资源数要求

	RequiredResourceNum *uint64 `json:"RequiredResourceNum,omitempty" name:"RequiredResourceNum"`
}

type GetSubTypeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSubTypeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DrillObjectHost struct {
	// 业务

	Biz *string `json:"Biz,omitempty" name:"Biz"`
	// 集群

	Set *string `json:"Set,omitempty" name:"Set"`
	// 模块

	Module *string `json:"Module,omitempty" name:"Module"`
	// IP

	InnerIp *string `json:"InnerIp,omitempty" name:"InnerIp"`
	// SN

	Sn *string `json:"Sn,omitempty" name:"Sn"`
}

type TerminateConditionMetric struct {
	// 指标名称

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 操作符

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type GetMainTypeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMainTypeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMainTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReserveSceneDrillRequest struct {
	*tchttp.BaseRequest

	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 下次演练的预约时间

	DrillReservation *string `json:"DrillReservation,omitempty" name:"DrillReservation"`
}

func (r *ReserveSceneDrillRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReserveSceneDrillRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateCondition struct {
	// and或者or

	Logic *string `json:"Logic,omitempty" name:"Logic"`
	// 指标列表

	MetricList []*TerminateConditionMetric `json:"MetricList,omitempty" name:"MetricList"`
}

type GetTargetTypeListRequest struct {
	*tchttp.BaseRequest
}

func (r *GetTargetTypeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTargetTypeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetExperienceLibraryListRequest struct {
	*tchttp.BaseRequest

	// 主类型

	MainType *string `json:"MainType,omitempty" name:"MainType"`
	// 子类型

	SubType *string `json:"SubType,omitempty" name:"SubType"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 时间筛选条件

	TimeFilter *TimeFilter `json:"TimeFilter,omitempty" name:"TimeFilter"`
	// 经验库名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页起始位

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页记录数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 目标类型

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 演练类型

	DrillType *string `json:"DrillType,omitempty" name:"DrillType"`
}

func (r *GetExperienceLibraryListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetExperienceLibraryListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSceneDrillLibraryListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSceneDrillLibraryListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSceneDrillLibraryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateComponentDrillLibraryRequest struct {
	*tchttp.BaseRequest

	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 演练编排节点列表

	DrillNodeList []*DrillNode `json:"DrillNodeList,omitempty" name:"DrillNodeList"`
	// 监控项

	MonitorMetrics []*MonitorMetric `json:"MonitorMetrics,omitempty" name:"MonitorMetrics"`
	// 监控对象

	MonitorMetricsObject *DrillObject `json:"MonitorMetricsObject,omitempty" name:"MonitorMetricsObject"`
	// 演练对象

	DrillObject *DrillObject `json:"DrillObject,omitempty" name:"DrillObject"`
	// biz名称

	BizName *string `json:"BizName,omitempty" name:"BizName"`
	// 终止条件

	TerminateCondition *TerminateCondition `json:"TerminateCondition,omitempty" name:"TerminateCondition"`
	// 终止目标

	TerminateConditionObject *DrillObject `json:"TerminateConditionObject,omitempty" name:"TerminateConditionObject"`
}

func (r *CreateComponentDrillLibraryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateComponentDrillLibraryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExperienceLibraryRequest struct {
	*tchttp.BaseRequest

	// 经验库名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 经验库Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeExperienceLibraryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExperienceLibraryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditComponentDrillLibraryRequest struct {
	*tchttp.BaseRequest

	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// DrillNodeList

	DrillNodeList []*DrillNode `json:"DrillNodeList,omitempty" name:"DrillNodeList"`
	// MonitorMetrics

	MonitorMetrics []*MonitorMetric `json:"MonitorMetrics,omitempty" name:"MonitorMetrics"`
	// MonitorMetricsObject

	MonitorMetricsObject *DrillObject `json:"MonitorMetricsObject,omitempty" name:"MonitorMetricsObject"`
	// DrillObject

	DrillObject *DrillObject `json:"DrillObject,omitempty" name:"DrillObject"`
	// 终止条件

	TerminateCondition *TerminateCondition `json:"TerminateCondition,omitempty" name:"TerminateCondition"`
	// 终止条件目标

	TerminateConditionObject *DrillObject `json:"TerminateConditionObject,omitempty" name:"TerminateConditionObject"`
	// biz名称

	BizName *string `json:"BizName,omitempty" name:"BizName"`
}

func (r *EditComponentDrillLibraryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditComponentDrillLibraryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditSceneDrillLibraryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditSceneDrillLibraryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditSceneDrillLibraryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComponentDrillLibraryBizNamesRequest struct {
	*tchttp.BaseRequest
}

func (r *GetComponentDrillLibraryBizNamesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComponentDrillLibraryBizNamesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComponentDrillLibraryBizNamesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComponentDrillLibraryBizNamesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComponentDrillLibraryBizNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PodSelector struct {
	// 命名空间

	NameSpace *string `json:"NameSpace,omitempty" name:"NameSpace"`
	// pod label

	Labels []*OptionKV `json:"Labels,omitempty" name:"Labels"`
}

type TerminateTaskRequest struct {
	*tchttp.BaseRequest

	// uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *TerminateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteComponentDrillLibraryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteComponentDrillLibraryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteComponentDrillLibraryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DrillConfigCommon struct {
	// 执行前等待时间

	WaitBeforeRun *uint64 `json:"WaitBeforeRun,omitempty" name:"WaitBeforeRun"`
	// 执行前等待时间单位

	WaitBeforeRunUnit *string `json:"WaitBeforeRunUnit,omitempty" name:"WaitBeforeRunUnit"`
	// 执行后等待时间

	WaitAfterRun *uint64 `json:"WaitAfterRun,omitempty" name:"WaitAfterRun"`
	// 执行后等待时间单位

	WaitAfterRunUnit *string `json:"WaitAfterRunUnit,omitempty" name:"WaitAfterRunUnit"`
	// 演练持续时间

	DurationValue *uint64 `json:"DurationValue,omitempty" name:"DurationValue"`
	// 演练持续时间单位

	DurationUnit *string `json:"DurationUnit,omitempty" name:"DurationUnit"`
	// 最小资源数要求

	RequiredResourceNum *uint64 `json:"RequiredResourceNum,omitempty" name:"RequiredResourceNum"`
}

type DrillObject struct {
	// HostList

	HostList []*DrillObjectHost `json:"HostList,omitempty" name:"HostList"`
	// 演练比例

	FilterPercent *uint64 `json:"FilterPercent,omitempty" name:"FilterPercent"`
	// 是否继承演练对象

	MonitorOption *bool `json:"MonitorOption,omitempty" name:"MonitorOption"`
	// 目标类型

	DrillObjectType *string `json:"DrillObjectType,omitempty" name:"DrillObjectType"`
	// cmdb节点列表

	CmdbNodes []*CmdbNodes `json:"CmdbNodes,omitempty" name:"CmdbNodes"`
	// pod选择器

	PodSelector *PodSelector `json:"PodSelector,omitempty" name:"PodSelector"`
	// 模块列表，已经废弃

	ModuleList []*DrillObjectModule `json:"ModuleList,omitempty" name:"ModuleList"`
	// region

	Region *string `json:"Region,omitempty" name:"Region"`
}

type CancelReserveSceneDrillResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelReserveSceneDrillResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelReserveSceneDrillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExecuteComponentDrillLibraryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExecuteComponentDrillLibraryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExecuteComponentDrillLibraryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelReserveSceneDrillRequest struct {
	*tchttp.BaseRequest

	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *CancelReserveSceneDrillRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelReserveSceneDrillRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBaseInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBaseInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetExperienceLibraryListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetExperienceLibraryListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetExperienceLibraryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOpRecordListRequest struct {
	*tchttp.BaseRequest

	// “ExperienceLibrary”，“ComponentDrill”，“Task”,"All"分别表示“经验库”，“组件演练库”，“任务”，“所有”

	ObjectType *string `json:"ObjectType,omitempty" name:"ObjectType"`
	// 操作人

	OpUser *string `json:"OpUser,omitempty" name:"OpUser"`
	// “Edit”，“Clone”，“Terminate”，“Create”，“Delete”,"All"分别表示“编辑”，“克隆”，“终止”，“创建”，“删除”,"所有"

	OpType *string `json:"OpType,omitempty" name:"OpType"`
	// 对象名称

	ObjectName *string `json:"ObjectName,omitempty" name:"ObjectName"`
	// 时间筛选参数

	TimeFilter *TimeFilter `json:"TimeFilter,omitempty" name:"TimeFilter"`
	// 分页起始位

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页记录数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetOpRecordListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOpRecordListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSceneDrillLibraryRequest struct {
	*tchttp.BaseRequest

	// 经验库名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 演练库id，和演练库名称二选一，同时使用以id为准，可选

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 删除策略，前端传递Foreground

	Policy *string `json:"Policy,omitempty" name:"Policy"`
}

func (r *DeleteSceneDrillLibraryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSceneDrillLibraryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
