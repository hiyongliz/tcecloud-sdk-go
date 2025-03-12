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

package v20191101

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DescribeModuleHostsRequest struct {
	*tchttp.BaseRequest

	// 模块的 bk_tx_unique_code

	UniqueCode *string `json:"UniqueCode,omitempty" name:"UniqueCode"`
	// 模块的 level，如 "5|16"

	Level *string `json:"Level,omitempty" name:"Level"`
}

func (r *DescribeModuleHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModuleHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePatrolCasesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 巡检用例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 巡检用例列表

		Items []*PatrolCase `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePatrolCasesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePatrolCasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePatrolCaseRequest struct {
	*tchttp.BaseRequest

	// 巡检用例名称，支持中文等字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 巡检类型，有效值：Alarm, Metric, Availability, General

	Type *string `json:"Type,omitempty" name:"Type"`
	// 该巡检用例的创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 该巡检用例是否是内置的用例

	BuiltIn *bool `json:"BuiltIn,omitempty" name:"BuiltIn"`
	// 业务模块树上的节点路径，由节点 ID 连接组成，例如 5>15>21

	NodePath *string `json:"NodePath,omitempty" name:"NodePath"`
	// 该巡检用例的描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 巡检用例的全局参数，在巡检任务中使用

	Params []*Param `json:"Params,omitempty" name:"Params"`
	// 巡检任务

	Jobs []*Job `json:"Jobs,omitempty" name:"Jobs"`
	// 定时执行该巡检用例的 Cron 表达式，如： 0 0 * * *

	Cron *string `json:"Cron,omitempty" name:"Cron"`
	// 是否启用该巡检用例

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 业务模块树上的节点路径，由名称构成

	NodeNamePath *string `json:"NodeNamePath,omitempty" name:"NodeNamePath"`
}

func (r *CreatePatrolCaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePatrolCaseRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowDetailRequest struct {
	*tchttp.BaseRequest

	// 作业编排 ID

	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DescribeFlowDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePatrolCaseRequest struct {
	*tchttp.BaseRequest

	// 巡检用例的 ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeletePatrolCaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePatrolCaseRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PatrolRecord struct {
	// 巡检记录 ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 巡检用例 ID

	CaseId *string `json:"CaseId,omitempty" name:"CaseId"`
	// 巡检用例名称

	CaseName *string `json:"CaseName,omitempty" name:"CaseName"`
	// 巡检类型，有效值：Alarm, Metric, Availability, General

	CaseType *string `json:"CaseType,omitempty" name:"CaseType"`
	// 业务模块树上的节点路径，由节点 ID 连接组成，例如 5>15>21

	NodePath *string `json:"NodePath,omitempty" name:"NodePath"`
	// 巡检任务列表

	Jobs []*Job `json:"Jobs,omitempty" name:"Jobs"`
	// 巡检用例整体执行结果，有效值：Healthy，Running，Unhealthy

	Result *string `json:"Result,omitempty" name:"Result"`
	// 该记录的创建时间，也即该记录的执行时间，为 Unix 时间

	CreationTime *int64 `json:"CreationTime,omitempty" name:"CreationTime"`
	// 该记录的最近更新时间，为 Unix 时间

	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// 业务模块树上的节点路径，由名称构成

	NodeNamePath *string `json:"NodeNamePath,omitempty" name:"NodeNamePath"`
}

type CreatePatrolCaseResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建的巡检用例 ID

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePatrolCaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePatrolCaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePatrolCaseResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePatrolCaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePatrolCaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// JSON 序列化后的作业编排详情

		Detail *string `json:"Detail,omitempty" name:"Detail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPatrolCaseRequest struct {
	*tchttp.BaseRequest

	// 巡检用例的 ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 巡检用例名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 该巡检用例是否是内置的

	BuiltIn *bool `json:"BuiltIn,omitempty" name:"BuiltIn"`
	// 业务模块树上的节点路径，由节点 ID 连接组成，例如 5>15>21

	NodePath *string `json:"NodePath,omitempty" name:"NodePath"`
	// 巡检用例的描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 巡检用例参数列表

	Params []*Param `json:"Params,omitempty" name:"Params"`
	// 巡检任务

	Jobs []*Job `json:"Jobs,omitempty" name:"Jobs"`
	// 巡检用例定时执行的 Cron 表达式

	Cron *string `json:"Cron,omitempty" name:"Cron"`
	// 该巡检用例是否开启

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 业务模块树上的节点路径，由名称构成

	NodeNamePath *string `json:"NodeNamePath,omitempty" name:"NodeNamePath"`
}

func (r *ModifyPatrolCaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPatrolCaseRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JobStatus struct {
	// 任务的执行状态，有效值：Pending， Running，Succeeded，Failed

	State *string `json:"State,omitempty" name:"State"`
	// 巡检任务结果，有效值: Healthy, Unhealthy

	Result *string `json:"Result,omitempty" name:"Result"`
	// 巡检任务执行时产生的日志

	Logs *string `json:"Logs,omitempty" name:"Logs"`
	// 巡检任务的数据输出，例如告警记录列表

	Outputs *string `json:"Outputs,omitempty" name:"Outputs"`
	// 日志是否是被截断的，我们最大只保留 64K 日志

	LogTruncated *bool `json:"LogTruncated,omitempty" name:"LogTruncated"`
}

type Param struct {
	// 参数名称，例如 url

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 参数的描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DescribeOperationsRequest struct {
	*tchttp.BaseRequest

	// 巡检用例 ID

	CaseId *string `json:"CaseId,omitempty" name:"CaseId"`
	// 巡检用例名称

	CaseName *string `json:"CaseName,omitempty" name:"CaseName"`
	// 巡检类型，有效值：Alarm, Metric, Availability, General

	CaseType *string `json:"CaseType,omitempty" name:"CaseType"`
	// 业务模块树上的节点路径，由节点 ID 连接组成，例如 5>15>21

	NodePath *string `json:"NodePath,omitempty" name:"NodePath"`
	// 操作类型，有效值：Create, Modify, Delete, Execute

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 操作者

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 操作的起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 操作的终止时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 用于分页的 Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回的最大结果数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOperationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePatrolCasesRequest struct {
	*tchttp.BaseRequest

	// 巡检用例 ID，如果指定了该值，直接返回该巡检用例

	Id *string `json:"Id,omitempty" name:"Id"`
	// 业务模块树上的节点路径，由节点 ID 连接组成，例如 5>15>21

	NodePath *string `json:"NodePath,omitempty" name:"NodePath"`
	// 偏移量，用于分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回的最大巡检用例数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePatrolCasesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePatrolCasesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSuggestionsRequest struct {
	*tchttp.BaseRequest

	// 推荐的内容类型，CaseId，CaseName，ArgusNamespace， ToolNames

	Type *string `json:"Type,omitempty" name:"Type"`
	// 查询字符串，将进行模糊匹配

	Query *string `json:"Query,omitempty" name:"Query"`
	// 最多返回推荐条目的个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetSuggestionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSuggestionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSuggestionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配的总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 推荐的数据列表

		Items []*string `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSuggestionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSuggestionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePatrolRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配的巡检记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 巡检记录列表

		Items []*PatrolRecord `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePatrolRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePatrolRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PatrolCase struct {
	// 巡检用例 ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 巡检用例名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 巡检类型巡检用例类型，有效值：Alarm, Metric, Availability, General

	Type *string `json:"Type,omitempty" name:"Type"`
	// 巡检用例创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 该巡检用例是否是内置的

	BuiltIn *bool `json:"BuiltIn,omitempty" name:"BuiltIn"`
	// 业务模块树上的节点路径，由节点 ID 连接组成，例如 5>15>21

	NodePath *string `json:"NodePath,omitempty" name:"NodePath"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 巡检用例全局参数

	Params []*Param `json:"Params,omitempty" name:"Params"`
	// 巡检任务

	Jobs []*Job `json:"Jobs,omitempty" name:"Jobs"`
	// 定时执行该用例的 Cron 表达式

	Cron *string `json:"Cron,omitempty" name:"Cron"`
	// 是否启用该巡检用例

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 该巡检用例最近一次执行的记录 ID

	LatestRecordId *string `json:"LatestRecordId,omitempty" name:"LatestRecordId"`
	// 业务模块树上的节点路径，由名称构成

	NodeNamePath *string `json:"NodeNamePath,omitempty" name:"NodeNamePath"`
}

type RunPatrolCaseRequest struct {
	*tchttp.BaseRequest

	// 巡检用例 ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *RunPatrolCaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunPatrolCaseRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePatrolRecordsRequest struct {
	*tchttp.BaseRequest

	// 巡检用例 ID

	CaseId *string `json:"CaseId,omitempty" name:"CaseId"`
	// 巡检用例名称

	CaseName *string `json:"CaseName,omitempty" name:"CaseName"`
	// 巡检类型，有效值：Alarm, Metric, Availability, General

	CaseType *string `json:"CaseType,omitempty" name:"CaseType"`
	// 巡检记录 ID

	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`
	// 业务模块树上的节点路径，由节点 ID 连接组成，例如 5>15>21

	NodePath *string `json:"NodePath,omitempty" name:"NodePath"`
	// 巡检记录的执行结果，有效值：Healthy, Unhealthy

	Result *string `json:"Result,omitempty" name:"Result"`
	// 是否只返回每个巡检用例最近的一次记录

	OnlyLatest *bool `json:"OnlyLatest,omitempty" name:"OnlyLatest"`
	// 是否只返回每个记录精简的信息

	Briefly *bool `json:"Briefly,omitempty" name:"Briefly"`
	// 列举巡检记录的起始执行时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 列举巡检记录的终止执行时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 偏移量，用于分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回的最大记录数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePatrolRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePatrolRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunPatrolCaseResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 巡检记录 ID

		RecordId *string `json:"RecordId,omitempty" name:"RecordId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunPatrolCaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunPatrolCaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeToolDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// JSON 序列化后的作业详情

		Detail *string `json:"Detail,omitempty" name:"Detail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeToolDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeToolDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPatrolCaseResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPatrolCaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPatrolCaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Job struct {
	// 巡检任务的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 巡检任务的详细描述，以 JSON 字符串形式保存

	Spec *string `json:"Spec,omitempty" name:"Spec"`
	// 巡检任务的执行状态

	Status *JobStatus `json:"Status,omitempty" name:"Status"`
}

type DescribeOperationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配的操作记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 操作记录列表

		Items []*AuditLog `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeToolDetailRequest struct {
	*tchttp.BaseRequest

	// 工具的 ID

	ToolId *int64 `json:"ToolId,omitempty" name:"ToolId"`
}

func (r *DescribeToolDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeToolDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Host struct {
	// Host 的资产编号

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 内部 IP

	InnerIp *string `json:"InnerIp,omitempty" name:"InnerIp"`
	// 外部 IP

	OuterIp *string `json:"OuterIp,omitempty" name:"OuterIp"`
}

type DescribeModuleHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Hosts 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Host 列表

		Items []*Host `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeModuleHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModuleHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditLog struct {
	// 巡检用例 ID

	CaseId *string `json:"CaseId,omitempty" name:"CaseId"`
	// 巡检用例名称

	CaseName *string `json:"CaseName,omitempty" name:"CaseName"`
	// 业务模块树上的节点路径，由节点 ID 连接组成，例如 5>15>21

	NodePath *string `json:"NodePath,omitempty" name:"NodePath"`
	// 操作类型，有效值：Create, Modify, Delete, Execute

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 操作者

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 操作时间

	OpTime *string `json:"OpTime,omitempty" name:"OpTime"`
	// 业务模块树上的节点路径，由名称构成

	NodeNamePath *string `json:"NodeNamePath,omitempty" name:"NodeNamePath"`
}
