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

package v20200724

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DescribeFlakeTimerLogsRequest struct {
	*tchttp.BaseRequest

	// 入参

	Params *TimerLogParams `json:"Params,omitempty" name:"Params"`
}

func (r *DescribeFlakeTimerLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlakeTimerLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateFlakeTimerRequest struct {
	*tchttp.BaseRequest

	// 入参

	Params *TimerParams `json:"Params,omitempty" name:"Params"`
}

func (r *UpdateFlakeTimerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateFlakeTimerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProdParams struct {
	// 数据集名字

	CollName *string `json:"CollName,omitempty" name:"CollName"`
	// 数据集主键类型字段

	KeyFields []*string `json:"KeyFields,omitempty" name:"KeyFields"`
	// 数据集时间类型字段

	TimeFields []*string `json:"TimeFields,omitempty" name:"TimeFields"`
	// 数据集描述

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 数据保留天数

	MaintainDays *uint64 `json:"MaintainDays,omitempty" name:"MaintainDays"`
	// 数据来源miner/timer

	Source *string `json:"Source,omitempty" name:"Source"`
	// 标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 分页参数Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 按照创建人名字查询条件

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 按照修改人名字查询条件

	Editor *string `json:"Editor,omitempty" name:"Editor"`
}

type DeleteFlakeViewRequest struct {
	*tchttp.BaseRequest

	// 入参

	Params *ViewParams `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteFlakeViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFlakeViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlakeTimersRequest struct {
	*tchttp.BaseRequest

	// 入参

	Params *TimerParams `json:"Params,omitempty" name:"Params"`
}

func (r *DescribeFlakeTimersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlakeTimersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TimerOutput struct {
	// 触发器列表

	Data []*Timer `json:"Data,omitempty" name:"Data"`
	// 总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type DeleteFlakeViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFlakeViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFlakeViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Timer struct {
	// 触发器名字

	TimerName *string `json:"TimerName,omitempty" name:"TimerName"`
	// 创建时间（戳）

	CTime *uint64 `json:"CTime,omitempty" name:"CTime"`
	// cron表达式

	Cron *string `json:"Cron,omitempty" name:"Cron"`
	// 下一轮执行时间（戳）

	NextRunAt *uint64 `json:"NextRunAt,omitempty" name:"NextRunAt"`
	// 上一次执行时间（戳）

	LastRunAt *uint64 `json:"LastRunAt,omitempty" name:"LastRunAt"`
	// 是否启用

	Activated *bool `json:"Activated,omitempty" name:"Activated"`
	// fql语句

	Fql *string `json:"Fql,omitempty" name:"Fql"`
	// 数据集

	OutTo *Coll `json:"OutTo,omitempty" name:"OutTo"`
	// 附加信息

	Ext *string `json:"Ext,omitempty" name:"Ext"`
}

type OspVisionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OspVisionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OspVisionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Prod struct {
	// 数据集名字

	CollName *string `json:"CollName,omitempty" name:"CollName"`
	// 描述

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 来源miner/timer

	Source *string `json:"Source,omitempty" name:"Source"`
	// cron表达式

	Cron *string `json:"Cron,omitempty" name:"Cron"`
	// 标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 是否是时序表

	IsTimeSeries *bool `json:"IsTimeSeries,omitempty" name:"IsTimeSeries"`
	// 数据保留天数

	MaintainDays *uint64 `json:"MaintainDays,omitempty" name:"MaintainDays"`
	// 创建时间（戳）

	CTime *uint64 `json:"CTime,omitempty" name:"CTime"`
	// 修改时间（戳）

	Uptime *uint64 `json:"Uptime,omitempty" name:"Uptime"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 修改人

	Editor *string `json:"Editor,omitempty" name:"Editor"`
	// 上次执行时间（戳）

	LastScheTime *uint64 `json:"LastScheTime,omitempty" name:"LastScheTime"`
	// 数据量大小（字节）

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 实际存储量（字节）

	StorageSize *uint64 `json:"StorageSize,omitempty" name:"StorageSize"`
	// 索引数量

	IndexCnt *uint64 `json:"IndexCnt,omitempty" name:"IndexCnt"`
	// 索引键

	KeyFields []*string `json:"KeyFields,omitempty" name:"KeyFields"`
	// 时序键

	TimeFields []*string `json:"TimeFields,omitempty" name:"TimeFields"`
	// 创建人uin

	CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`
	// 修改人uin

	EditorUin *string `json:"EditorUin,omitempty" name:"EditorUin"`
}

type CreateFlakeTimerRequest struct {
	*tchttp.BaseRequest

	// 入参

	Params *TimerParams `json:"Params,omitempty" name:"Params"`
}

func (r *CreateFlakeTimerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFlakeTimerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFlakeViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFlakeViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFlakeViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProdCollRequest struct {
	*tchttp.BaseRequest

	// 参数

	Params *ProdParams `json:"Params,omitempty" name:"Params"`
}

func (r *CreateProdCollRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProdCollRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProdCollRequest struct {
	*tchttp.BaseRequest

	// 入参

	Params *ProdParams `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteProdCollRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProdCollRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CallFlakeViewRequest struct {
	*tchttp.BaseRequest

	// 入参

	Params *ViewCallParam `json:"Params,omitempty" name:"Params"`
}

func (r *CallFlakeViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CallFlakeViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProdCollResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProdCollResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProdCollResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateFlakeTimerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateFlakeTimerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateFlakeTimerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OspVisionRequest struct {
	*tchttp.BaseRequest

	// input

	Params *OspVisionInput `json:"Params,omitempty" name:"Params"`
}

func (r *OspVisionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OspVisionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateFlakeViewRequest struct {
	*tchttp.BaseRequest

	// 入参

	Params *ViewParams `json:"Params,omitempty" name:"Params"`
}

func (r *UpdateFlakeViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateFlakeViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ViewCallVar struct {
	// Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateProdCollResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProdCollResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProdCollResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OspVisionInput struct {
	// 取值: "osp", "cvm", "cbs"

	VisionType *string `json:"VisionType,omitempty" name:"VisionType"`
}

type DeleteFlakeTimerRequest struct {
	*tchttp.BaseRequest

	// 入参

	Params *TimerParams `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteFlakeTimerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFlakeTimerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProdCollsRequest struct {
	*tchttp.BaseRequest

	// 入参

	Params *ProdParams `json:"Params,omitempty" name:"Params"`
}

func (r *DescribeProdCollsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProdCollsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFlakeTimerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFlakeTimerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFlakeTimerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProdCollResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProdCollResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProdCollResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TimerParams struct {
	// 分页参数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 触发器名字

	TimerName *string `json:"TimerName,omitempty" name:"TimerName"`
	// cron表达式

	Cron *string `json:"Cron,omitempty" name:"Cron"`
	// fql语句

	Fql *string `json:"Fql,omitempty" name:"Fql"`
	// 是否启用

	Activated *bool `json:"Activated,omitempty" name:"Activated"`
	// 数据集

	OutTo *Coll `json:"OutTo,omitempty" name:"OutTo"`
	// 附加信息

	Ext *string `json:"Ext,omitempty" name:"Ext"`
}

type ViewOutput struct {
	// 视图列表

	Data []*View `json:"Data,omitempty" name:"Data"`
	// 总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type DescribeFlakeTimerLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 出参

		Data *TimerLogOutput `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlakeTimerLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlakeTimerLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProdCollsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 出参

		Data *ProdOutput `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProdCollsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProdCollsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ViewCallParam struct {
	// 调用的视图名字

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 调用视图的参数

	Vars []*ViewCallVar `json:"Vars,omitempty" name:"Vars"`
}

type ViewParams struct {
	// 视图名

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// fql语句

	Fql *string `json:"Fql,omitempty" name:"Fql"`
	// 说明

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 分页参数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 按照创建人名字查询条件

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 按照修改人名字查询

	Editor *string `json:"Editor,omitempty" name:"Editor"`
}

type CreateFlakeViewRequest struct {
	*tchttp.BaseRequest

	// 入参

	Params *ViewParams `json:"Params,omitempty" name:"Params"`
}

func (r *CreateFlakeViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFlakeViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFlakeTimerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFlakeTimerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFlakeTimerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateFlakeViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateFlakeViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateFlakeViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TimerLog struct {
	// 触发器名字

	TimerName *string `json:"TimerName,omitempty" name:"TimerName"`
	// 影响行数

	AffectedRows *uint64 `json:"AffectedRows,omitempty" name:"AffectedRows"`
	// 执行时间（戳）

	CTime *uint64 `json:"CTime,omitempty" name:"CTime"`
	// fql语句

	Fql *string `json:"Fql,omitempty" name:"Fql"`
	// 附加信息

	Ext *string `json:"Ext,omitempty" name:"Ext"`
}

type CallFlakeViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CallFlakeViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CallFlakeViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProdCollRequest struct {
	*tchttp.BaseRequest

	// Params

	Params *ProdParams `json:"Params,omitempty" name:"Params"`
}

func (r *UpdateProdCollRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProdCollRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VisionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VisionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VisionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type View struct {
	// 视图名

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// fql语句

	Fql *string `json:"Fql,omitempty" name:"Fql"`
	// 创建时间（戳）

	CTime *uint64 `json:"CTime,omitempty" name:"CTime"`
	// 说明

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 创建人uin

	CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`
	// 修改人

	Editor *string `json:"Editor,omitempty" name:"Editor"`
	// 修改人uin

	EditorUin *string `json:"EditorUin,omitempty" name:"EditorUin"`
}

type DescribeFlakeViewsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 出参

		Data *ViewOutput `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlakeViewsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlakeViewsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VisionRequest struct {
	*tchttp.BaseRequest
}

func (r *VisionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VisionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Coll struct {
	// 数据集名字

	CollName *string `json:"CollName,omitempty" name:"CollName"`
	// 主键字段

	KeyFields []*string `json:"KeyFields,omitempty" name:"KeyFields"`
	// 时间字段

	TimeFields []*string `json:"TimeFields,omitempty" name:"TimeFields"`
}

type TimerLogOutput struct {
	// 日志列表

	Data []*TimerLog `json:"Data,omitempty" name:"Data"`
	// 总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type TimerLogParams struct {
	// 触发器名字

	TimerName *string `json:"TimerName,omitempty" name:"TimerName"`
	// 分页参数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 开始时间(时间戳)

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间(时间戳)

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeFlakeTimersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 出参

		Data *TimerOutput `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlakeTimersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlakeTimersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlakeViewsRequest struct {
	*tchttp.BaseRequest

	// 入参

	Params *ViewParams `json:"Params,omitempty" name:"Params"`
}

func (r *DescribeFlakeViewsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlakeViewsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProdOutput struct {
	// 数据集列表

	Data []*Prod `json:"Data,omitempty" name:"Data"`
	// 总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}
