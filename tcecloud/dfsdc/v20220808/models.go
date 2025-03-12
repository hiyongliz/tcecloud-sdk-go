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

package v20220808

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DescribeAccountingBusinessesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccountingBusinessesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountingBusinessesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountingResultsOfLastPeriodResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 对账任务列表

		AccountingResults []*AccountingResult `json:"AccountingResults,omitempty" name:"AccountingResults"`
		// 本轮周期的起始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 本轮周期的结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountingResultsOfLastPeriodResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountingResultsOfLastPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountingResultsRequest struct {
	*tchttp.BaseRequest

	// 查询对账范围的起始时间；例：2019-07-05 00:00:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询对账范围的截止时间；例：2019-07-05 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 业务类型，当前只支持DC。

	Business *string `json:"Business,omitempty" name:"Business"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。暂未用到，先预留。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAccountingResultsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountingResultsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountingResult struct {
	// 对账业务类型

	Business *string `json:"Business,omitempty" name:"Business"`
	// 对账业务子类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 任务Id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 任务状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 对账结果详情

	ResultDetail *string `json:"ResultDetail,omitempty" name:"ResultDetail"`
	// 对账资源Id

	UniqResId *string `json:"UniqResId,omitempty" name:"UniqResId"`
	// 对账步骤

	Steps *string `json:"Steps,omitempty" name:"Steps"`
	// 当前对账步骤

	Step *string `json:"Step,omitempty" name:"Step"`
	// 当前对账步骤状态

	StepStatus *string `json:"StepStatus,omitempty" name:"StepStatus"`
	// 修复规则

	FixRule *string `json:"FixRule,omitempty" name:"FixRule"`
	// Region

	Region *string `json:"Region,omitempty" name:"Region"`
	// AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 对账类型，数据库、ZK、交换机配置等。

	AccountingType *string `json:"AccountingType,omitempty" name:"AccountingType"`
	// 对账源目的说明

	SrcDstDesc *string `json:"SrcDstDesc,omitempty" name:"SrcDstDesc"`
	// IDLE:等待对账     PROCESSING:正在对账     FAILURE：对账失败     SUCCESS: 对账完成无异常     FIX_SUCCESS：修复成功     FIX_NONE：未修复     FIX_FAILURE:修复失败     ROLLBACK_SUCCESS:回滚成功     ROLLBACK_FAILURE:回滚失败

	AccountingStatus *string `json:"AccountingStatus,omitempty" name:"AccountingStatus"`
}

type AccountingTask struct {
	// 业务类型

	Business *string `json:"Business,omitempty" name:"Business"`
	// 业务子类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// Region

	Region *string `json:"Region,omitempty" name:"Region"`
	// APPId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 对账资源Id

	UniqResId *string `json:"UniqResId,omitempty" name:"UniqResId"`
	// 任务Id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 是否自动修复，当前页面点击手工对账都传True。

	AutoFix *bool `json:"AutoFix,omitempty" name:"AutoFix"`
}

type UpdateAccountingTasksRequest struct {
	*tchttp.BaseRequest

	// 业务类型，当前只支持DC。

	Business *string `json:"Business,omitempty" name:"Business"`
	// 对账任务列表

	Tasks []*AccountingTask `json:"Tasks,omitempty" name:"Tasks"`
}

func (r *UpdateAccountingTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAccountingTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountingBusinessesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 支持综合对账的业务类型

		Businesses []*string `json:"Businesses,omitempty" name:"Businesses"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountingBusinessesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountingBusinessesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountingResultsByTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 对账任务列表

		AccountingResults []*AccountingResult `json:"AccountingResults,omitempty" name:"AccountingResults"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountingResultsByTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountingResultsByTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccountingTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 对账任务列表，比入参多了TaskId返回

		Tasks []*AccountingTask `json:"Tasks,omitempty" name:"Tasks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccountingTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccountingTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountingResultsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 对账任务列表

		AccountingResults []*AccountingResult `json:"AccountingResults,omitempty" name:"AccountingResults"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountingResultsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountingResultsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountingResultsOfLastPeriodRequest struct {
	*tchttp.BaseRequest

	// 业务类型，当前只支持DC。

	Business *string `json:"Business,omitempty" name:"Business"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAccountingResultsOfLastPeriodRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountingResultsOfLastPeriodRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type CreateAccountingTasksRequest struct {
	*tchttp.BaseRequest

	// 业务类型，当前只支持DC。

	Business *string `json:"Business,omitempty" name:"Business"`
	// 对账任务列表

	Tasks []*AccountingTask `json:"Tasks,omitempty" name:"Tasks"`
}

func (r *CreateAccountingTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccountingTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountingResultsByTaskRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 业务类型，当前只支持DC。

	Business *string `json:"Business,omitempty" name:"Business"`
	// 任务信息。

	Tasks []*AccountingTask `json:"Tasks,omitempty" name:"Tasks"`
}

func (r *DescribeAccountingResultsByTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountingResultsByTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAccountingTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 对账任务列表，比入参多了TaskId返回

		Tasks []*AccountingTask `json:"Tasks,omitempty" name:"Tasks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAccountingTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAccountingTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
