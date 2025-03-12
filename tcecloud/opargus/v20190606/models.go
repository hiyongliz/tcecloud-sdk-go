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

package v20190606

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type ArgusReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回描述信息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ArgusReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ArgusReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescFavoriteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配到的收藏夹

		Favorite *FavoriteOut `json:"Favorite,omitempty" name:"Favorite"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescFavoriteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescFavoriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateViewRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 视图Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 视图名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 视图描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlertHistory struct {
	// 告警状态。未恢复、已恢复、已失效

	AlarmStatus *string `json:"AlarmStatus,omitempty" name:"AlarmStatus"`
	// 告警内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 告警发生时间

	OccurTime *uint64 `json:"OccurTime,omitempty" name:"OccurTime"`
	// 告警发送状态。已收敛、已屏蔽、已发送、无订阅

	Status *string `json:"Status,omitempty" name:"Status"`
	// 所属告警策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 过滤条件

	FilterCondition *string `json:"FilterCondition,omitempty" name:"FilterCondition"`
	// 所属namespace名称

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 告警持续时间，单位秒

	AlarmDuration *uint64 `json:"AlarmDuration,omitempty" name:"AlarmDuration"`
	// 告警接收渠道，多个逗号分隔

	ReceiveChannels *string `json:"ReceiveChannels,omitempty" name:"ReceiveChannels"`
}

type MetricIn struct {
	// 英文名

	EnName *string `json:"EnName,omitempty" name:"EnName"`
	// 中文名

	CnName *string `json:"CnName,omitempty" name:"CnName"`
	// 聚合方式

	AggType *string `json:"AggType,omitempty" name:"AggType"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 分类

	Category *string `json:"Category,omitempty" name:"Category"`
	// 操作类型。支持add,delete,update,noChange

	OpType *string `json:"OpType,omitempty" name:"OpType"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type AddFavoriteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddFavoriteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddFavoriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescNamespaceRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// namespace的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// namespace的Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescNamespaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateFavoriteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateFavoriteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateFavoriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 老的告警策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 新的告警策略名称

	NewName *string `json:"NewName,omitempty" name:"NewName"`
	// 告警策略ID

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 告警策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// NamespaceId

	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 由于云网关限制，每个条件需要单独进行JSON编码。多个条件之间是与的关系，用二级json格式来描述。比如：[[name,=,zhangsan],[city,in,[shenzhen,beijing]]]。当前支持比较符有“=”、“!=”、“>”、“>=”、“<”、“<=”、“in”

	FilterConditions []*string `json:"FilterConditions,omitempty" name:"FilterConditions"`
	// 分组字段

	GroupBys []*string `json:"GroupBys,omitempty" name:"GroupBys"`
	// 告警接收用户的ID

	ReceiveUsers []*uint64 `json:"ReceiveUsers,omitempty" name:"ReceiveUsers"`
	// 告警接收用户组的ID

	ReceiveGroups []*uint64 `json:"ReceiveGroups,omitempty" name:"ReceiveGroups"`
	// 告警接收渠道

	ReceiveChannels []*string `json:"ReceiveChannels,omitempty" name:"ReceiveChannels"`
	// 触发条件

	PolicyTriggers []*PolicyTriggerIn `json:"PolicyTriggers,omitempty" name:"PolicyTriggers"`
	// 字段名为中文名，其他与FilterConditions相同

	FilterConditionCns []*string `json:"FilterConditionCns,omitempty" name:"FilterConditionCns"`
}

func (r *UpdatePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFavoriteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFavoriteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFavoriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFavoritesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配到的收藏夹列表

		Favorites []*FavoriteOut `json:"Favorites,omitempty" name:"Favorites"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFavoritesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFavoritesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddViewRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 所属的收藏夹Id列表

	FavoriteIds []*uint64 `json:"FavoriteIds,omitempty" name:"FavoriteIds"`
	// 所属的NamespaceId

	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 视图名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 视图参数

	Param *string `json:"Param,omitempty" name:"Param"`
	// 视图描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 指标字段

	Metric *string `json:"Metric,omitempty" name:"Metric"`
}

func (r *AddViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ViewOut struct {
	// 视图Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 所属收藏夹Id

	FavoriteId *uint64 `json:"FavoriteId,omitempty" name:"FavoriteId"`
	// 所属命名空间Id

	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 视图名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 视图参数

	Param *string `json:"Param,omitempty" name:"Param"`
	// AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 创建人

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// 创建的时间戳

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新人

	UpdateUin *string `json:"UpdateUin,omitempty" name:"UpdateUin"`
	// 更新的时间戳

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 视图描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 所属命名空间名称

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 指标字段

	Metric *string `json:"Metric,omitempty" name:"Metric"`
}

type PolicyOut struct {
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 告警策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 告警策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// NamespaceId

	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 过滤条件，逗号分隔

	FilterCondition *string `json:"FilterCondition,omitempty" name:"FilterCondition"`
	// 分组字段，逗号分隔

	GroupBy *string `json:"GroupBy,omitempty" name:"GroupBy"`
	// 告警接收用户，逗号分隔

	ReceiveUsers *string `json:"ReceiveUsers,omitempty" name:"ReceiveUsers"`
	// 告警接收用户组，逗号分隔

	ReceiveGroups *string `json:"ReceiveGroups,omitempty" name:"ReceiveGroups"`
	// 告警接收渠道，逗号分隔

	ReceiveChannels *string `json:"ReceiveChannels,omitempty" name:"ReceiveChannels"`
	// AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 创建人

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// 创建的时间戳

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新人

	UpdateUin *string `json:"UpdateUin,omitempty" name:"UpdateUin"`
	// 更新的时间戳

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// namespace名称

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 状态。正常ok，异常error

	Status *string `json:"Status,omitempty" name:"Status"`
	// 主账号Uin

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

type PolicyTriggerOut struct {
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 告警策略Id

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 检测类型。threshold表示阈值检测

	Type *string `json:"Type,omitempty" name:"Type"`
	// 指标名称

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 比较符

	CompareSymbol *string `json:"CompareSymbol,omitempty" name:"CompareSymbol"`
	// 阈值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 创建人

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// 创建的时间戳

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新人

	UpdateUin *string `json:"UpdateUin,omitempty" name:"UpdateUin"`
	// 更新的时间戳

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 连续触发阈值的周期次数

	ContinueCount *uint64 `json:"ContinueCount,omitempty" name:"ContinueCount"`
	// 告警频率，单位分钟

	AlertFrequency *uint64 `json:"AlertFrequency,omitempty" name:"AlertFrequency"`
	// AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 主账号Uin

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

type GetAlertHistoriesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配到的告警历史

		AlertHistories []*AlertHistory `json:"AlertHistories,omitempty" name:"AlertHistories"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAlertHistoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DimensionIn struct {
	// 英文名

	EnName *string `json:"EnName,omitempty" name:"EnName"`
	// 中文名

	CnName *string `json:"CnName,omitempty" name:"CnName"`
	// 操作类型。支持add,delete,update,noChange

	OpType *string `json:"OpType,omitempty" name:"OpType"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DescPolicyRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警策略的名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配到的视图

		View *ViewOut `json:"View,omitempty" name:"View"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetViewsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配到的收藏夹列表

		Views []*ViewOut `json:"Views,omitempty" name:"Views"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetViewsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetViewsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ArgusReportRequest struct {
	*tchttp.BaseRequest

	// 指定上报的数据格式，其中lp代表Line Protocol，json代表Json数据协议

	Type *string `json:"Type,omitempty" name:"Type"`
	// 当Type指定为lp时使用，数据格式说明请参考：http://docs.influxdata.com/influxdb/v1.7/write_protocols/line_protocol_tutorial/
	// 其中measurement部分填写上报的命名空间名称，时间戳精确到纳秒，支持单个请求上报多个命名空间数据

	LpData []*string `json:"LpData,omitempty" name:"LpData"`
	// 当Type指定为json时使用，支持单个请求上报多个命名空间数据

	JsonData []*ReportJson `json:"JsonData,omitempty" name:"JsonData"`
}

func (r *ArgusReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ArgusReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNamespaceRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// namespace的名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNamespaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReportJson struct {
	// 命名空间名称

	Ns *string `json:"Ns,omitempty" name:"Ns"`
	// 维度组合信息

	Dimensions []*ReportDimension `json:"Dimensions,omitempty" name:"Dimensions"`
	// 指标信息

	Metrics []*ReportMetric `json:"Metrics,omitempty" name:"Metrics"`
	// 时间戳（精确到毫秒）

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type GetOverviewRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// namespace的名称列表

	NamespaceNames []*string `json:"NamespaceNames,omitempty" name:"NamespaceNames"`
	// 返回namespace个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetParamsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标单位

		Units []*ParamOut `json:"Units,omitempty" name:"Units"`
		// 聚合方式

		AggTypes []*ParamOut `json:"AggTypes,omitempty" name:"AggTypes"`
		// 聚合周期

		AggPeriods []*ParamOut `json:"AggPeriods,omitempty" name:"AggPeriods"`
		// 告警接收渠道

		AlertChannels []*ParamOut `json:"AlertChannels,omitempty" name:"AlertChannels"`
		// 连续触发阈值的周期次数

		AlertContinues []*ParamOut `json:"AlertContinues,omitempty" name:"AlertContinues"`
		// 告警频率

		AlertFrequencies []*ParamOut `json:"AlertFrequencies,omitempty" name:"AlertFrequencies"`
		// 指标比较符

		CompareSymbol []*ParamOut `json:"CompareSymbol,omitempty" name:"CompareSymbol"`
		// namespace配置限制

		NamespaceRestrict []*ParamOut `json:"NamespaceRestrict,omitempty" name:"NamespaceRestrict"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetParamsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FavoriteOut struct {
	// 收藏夹Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 收藏夹名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 收藏夹下视图数量

	ViewNum *uint64 `json:"ViewNum,omitempty" name:"ViewNum"`
	// AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 创建人

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// 创建的时间戳

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新人

	UpdateUin *string `json:"UpdateUin,omitempty" name:"UpdateUin"`
	// 更新的时间戳

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 收藏夹描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type UpdateViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReportDimension struct {
	// 维度名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 维度值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type AddNamespaceRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// namespace名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// namespace描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 数据最大保存时间，单位天

	DataTimeLimit *uint64 `json:"DataTimeLimit,omitempty" name:"DataTimeLimit"`
	// 数据最大存储容量，单位GB

	DataDiskLimit *uint64 `json:"DataDiskLimit,omitempty" name:"DataDiskLimit"`
	// 聚合周期

	AggPeriod *uint64 `json:"AggPeriod,omitempty" name:"AggPeriod"`
	// 维度列表

	Dimensions []*DimensionIn `json:"Dimensions,omitempty" name:"Dimensions"`
	// 指标列表

	Metrics []*MetricIn `json:"Metrics,omitempty" name:"Metrics"`
}

func (r *AddNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNamespaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAlertHistoriesRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 开始时间，毫秒时间戳。默认值1小时前

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，毫秒时间戳。默认值当前时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 告警策略Id

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 告警接收用户组的ID

	ReceiveGroup *uint64 `json:"ReceiveGroup,omitempty" name:"ReceiveGroup"`
	// 查询关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 分页Offset，默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页Limit，默认10

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段，默认occurTime。另外还支持status、alarmStatus

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 升序用asc，降序用desc，默认desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 告警发送状态。已收敛、已屏蔽、已发送、无订阅

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 告警状态。未恢复、已恢复、已失效

	AlarmStatus []*string `json:"AlarmStatus,omitempty" name:"AlarmStatus"`
}

func (r *GetAlertHistoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertHistoriesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配到的告警策略列表

		Policies []*PolicyOut `json:"Policies,omitempty" name:"Policies"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNamespacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配到的namespace列表

		Namespaces []*NamespaceOut `json:"Namespaces,omitempty" name:"Namespaces"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NamespaceOut struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 数据最大保存时间，单位天

	DataTimeLimit *uint64 `json:"DataTimeLimit,omitempty" name:"DataTimeLimit"`
	// 数据最大存储容量，单位GB

	DataDiskLimit *uint64 `json:"DataDiskLimit,omitempty" name:"DataDiskLimit"`
	// 聚合周期

	AggPeriod *uint64 `json:"AggPeriod,omitempty" name:"AggPeriod"`
	// AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 创建人

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// 创建的时间戳

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新人

	UpdateUin *string `json:"UpdateUin,omitempty" name:"UpdateUin"`
	// 更新的时间戳

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否有数据

	HasData *bool `json:"HasData,omitempty" name:"HasData"`
	// 磁盘已使用容量，单位MB

	DataDiskUsage *uint64 `json:"DataDiskUsage,omitempty" name:"DataDiskUsage"`
}

type DeleteFavoriteRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 收藏夹Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteFavoriteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFavoriteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateFavoriteRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 收藏夹Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 收藏夹名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 收藏夹描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateFavoriteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateFavoriteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddFavoriteRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 收藏夹名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 收藏夹描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *AddFavoriteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddFavoriteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteViewRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 视图Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamOut struct {
	// 选项值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type AddNamespaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNamespaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetParamsRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *GetParamsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetParamsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNamespaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DimensionOut struct {
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// namespace的Id

	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 英文名

	EnName *string `json:"EnName,omitempty" name:"EnName"`
	// 中文名

	CnName *string `json:"CnName,omitempty" name:"CnName"`
	// 创建人

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// 创建的时间戳

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新人

	UpdateUin *string `json:"UpdateUin,omitempty" name:"UpdateUin"`
	// 更新的时间戳

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type DescFavoriteRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 收藏夹Id。Id和Name必填其一，优先匹配Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 收藏夹名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescFavoriteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescFavoriteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配到的告警策略

		Policy *PolicyOut `json:"Policy,omitempty" name:"Policy"`
		// 触发条件

		PolicyTriggers []*PolicyTriggerOut `json:"PolicyTriggers,omitempty" name:"PolicyTriggers"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 最近访问的namespace列表

		RecentNamespaces []*NamespaceOut `json:"RecentNamespaces,omitempty" name:"RecentNamespaces"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警策略的名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeletePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDimensionsRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// namespace的名称

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

func (r *GetDimensionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDimensionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNamespaceRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 老的namespace名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// namespace描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 数据最大保存时间，单位天

	DataTimeLimit *uint64 `json:"DataTimeLimit,omitempty" name:"DataTimeLimit"`
	// 数据最大存储容量，单位GB

	DataDiskLimit *uint64 `json:"DataDiskLimit,omitempty" name:"DataDiskLimit"`
	// 聚合周期

	AggPeriod *uint64 `json:"AggPeriod,omitempty" name:"AggPeriod"`
	// 维度列表

	Dimensions []*DimensionIn `json:"Dimensions,omitempty" name:"Dimensions"`
	// 指标列表

	Metrics []*MetricIn `json:"Metrics,omitempty" name:"Metrics"`
	// 新的namespace名称

	NewName *string `json:"NewName,omitempty" name:"NewName"`
}

func (r *UpdateNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNamespaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetricOut struct {
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// namespace的Id

	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 英文名

	EnName *string `json:"EnName,omitempty" name:"EnName"`
	// 中文名

	CnName *string `json:"CnName,omitempty" name:"CnName"`
	// 聚合方式

	AggType *string `json:"AggType,omitempty" name:"AggType"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 分类

	Category *string `json:"Category,omitempty" name:"Category"`
	// 创建人

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// 创建的时间戳

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新人

	UpdateUin *string `json:"UpdateUin,omitempty" name:"UpdateUin"`
	// 更新的时间戳

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type AddPolicyRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 告警策略描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// NamespaceId

	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 由于云网关限制，每个条件需要单独进行JSON编码。多个条件之间是与的关系，用二级json格式来描述。比如：[[name,=,zhangsan],[city,in,[shenzhen,beijing]]]。当前支持比较符有“=”、“!=”、“>”、“>=”、“<”、“<=”、“in”

	FilterConditions []*string `json:"FilterConditions,omitempty" name:"FilterConditions"`
	// 分组字段

	GroupBys []*string `json:"GroupBys,omitempty" name:"GroupBys"`
	// 告警接收用户的ID

	ReceiveUsers []*uint64 `json:"ReceiveUsers,omitempty" name:"ReceiveUsers"`
	// 告警接收用户组的ID

	ReceiveGroups []*uint64 `json:"ReceiveGroups,omitempty" name:"ReceiveGroups"`
	// 告警接收渠道

	ReceiveChannels []*string `json:"ReceiveChannels,omitempty" name:"ReceiveChannels"`
	// 触发条件

	PolicyTriggers []*PolicyTriggerIn `json:"PolicyTriggers,omitempty" name:"PolicyTriggers"`
	// 字段名为中文名，其他与FilterConditions相同

	FilterConditionCns []*string `json:"FilterConditionCns,omitempty" name:"FilterConditionCns"`
}

func (r *AddPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescNamespaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配到的namespace

		Namespace *NamespaceOut `json:"Namespace,omitempty" name:"Namespace"`
		// 维度列表

		Dimensions []*DimensionOut `json:"Dimensions,omitempty" name:"Dimensions"`
		// 指标列表

		Metrics []*MetricOut `json:"Metrics,omitempty" name:"Metrics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDataRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 开始时间，毫秒时间戳。默认值1小时前

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，毫秒时间戳。默认值当前时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// namespace的名称

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 如果同时指定了groupBy参数，那么就能使用聚合函数修饰字段，当前支持sum()、avg()、min()、max()、count()。如果不填，则忽略所有其他参数，返回指定table的10条明细数据（一般用于测试）

	Fields []*string `json:"Fields,omitempty" name:"Fields"`
	// 由于云网关限制，每个条件需要单独进行JSON编码。多个条件之间是与的关系，用二级json格式来描述。比如：[[name,=,zhangsan],[city,in,[shenzhen,beijing]]]。当前支持比较符有“=”、“!=”、“>”、“>=”、“<”、“<=”、“in” 。暂不支持对聚合后的结果进行过滤

	Conditions []*string `json:"Conditions,omitempty" name:"Conditions"`
	// 支持按时间粒度聚合，如：timestamp(5s)，粒度单位支持s，m，h，d。如果用了groupBy参数，则查询结果的字段名称为固定格式，比如sum(timeCost)返回的字段名称为timeCost_sum

	GroupBys []*string `json:"GroupBys,omitempty" name:"GroupBys"`
	// 排序字段，只支持一个字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 升序用asc，降序用desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 分页Offset，仅支持明细查询，不支持聚合查询。Limit+Offset不能超过65536

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页Limit，仅支持明细查询，不支持聚合查询。Limit+Offset不能超过65536

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// namespace的Id

	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *GetDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PolicyTriggerIn struct {
	// 检测类型。threshold表示阈值检测

	Type *string `json:"Type,omitempty" name:"Type"`
	// 指标名称

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 比较符

	CompareSymbol *string `json:"CompareSymbol,omitempty" name:"CompareSymbol"`
	// 阈值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 连续触发阈值的周期次数

	ContinueCount *uint64 `json:"ContinueCount,omitempty" name:"ContinueCount"`
	// 告警频率，单位分钟

	AlertFrequency *uint64 `json:"AlertFrequency,omitempty" name:"AlertFrequency"`
	// 触发条件的Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 操作类型。支持add,delete,update,noChange

	OpType *string `json:"OpType,omitempty" name:"OpType"`
}

type GetViewsRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 收藏夹Id

	FavoriteId *uint64 `json:"FavoriteId,omitempty" name:"FavoriteId"`
	// 视图名称，支持模糊搜索

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页Offset，默认值0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页Limit，默认值10

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段，只支持一个字段。Id, Name, FavoriteId, CreateTime, UpdateTime

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 升序用asc，降序用desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 收藏夹名称

	FavoriteName *string `json:"FavoriteName,omitempty" name:"FavoriteName"`
}

func (r *GetViewsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetViewsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DimensionValuesOut struct {
	// 英文名

	EnName *string `json:"EnName,omitempty" name:"EnName"`
	// 值列表

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 中文名

	CnName *string `json:"CnName,omitempty" name:"CnName"`
}

type GetFavoritesRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 收藏夹名称，支持模糊搜索

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页Offset，默认值0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页Limit，默认值10

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段，只支持一个字段。Id, Name, CreateTime, UpdateTime

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 升序用asc，降序用desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *GetFavoritesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFavoritesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReportMetric struct {
	// 指标名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 指标值

	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type GetDimensionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Dimension下拉框数据的列表

		Dimensions []*DimensionValuesOut `json:"Dimensions,omitempty" name:"Dimensions"`
		// 指标列表

		Metrics []*MetricOut `json:"Metrics,omitempty" name:"Metrics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDimensionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDimensionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNamespacesRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// namespace名称，支持模糊搜索

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段，只支持一个字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 升序用asc，降序用desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *GetNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPoliciesRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警策略名称，支持模糊搜索

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页Offset，默认值0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页Limit，默认值10

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段，只支持一个字段。Id, Name, NamespaceId, ContinueCount, AlertFrequency, CreateTime, UpdateTime

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 升序用asc，降序用desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 告警策略Id列表

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 创建人Uin

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// NamespaceId

	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *GetPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescViewRequest struct {
	*tchttp.BaseRequest

	// 模块路径

	Module *string `json:"Module,omitempty" name:"Module"`
	// 视图Id。Id和Name必填其一，优先匹配Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 视图名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 收藏夹Id

	FavoriteId *uint64 `json:"FavoriteId,omitempty" name:"FavoriteId"`
}

func (r *DescViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
