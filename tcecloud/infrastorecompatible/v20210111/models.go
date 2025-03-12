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

package v20210111

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DescribeProductEventListRequest struct {
	*tchttp.BaseRequest

	// 接口模块名，固定值"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 产品类型过滤，比如"cvm"表示云服务器

	ProductName []*string `json:"ProductName,omitempty" name:"ProductName"`
	// 事件名称过滤，比如"guest_reboot"表示机器重启

	EventName []*string `json:"EventName,omitempty" name:"EventName"`
	// 影响对象，比如ins-19708ino

	InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 维度过滤，比如外网IP:10.0.0.1

	Dimensions []*DescribeProductEventListDimensions `json:"Dimensions,omitempty" name:"Dimensions"`
	// 地域过滤，比如gz

	RegionList []*string `json:"RegionList,omitempty" name:"RegionList"`
	// 事件类型过滤，取值范围["status_change","abnormal"]，分别表示状态变更、异常事件

	Type []*string `json:"Type,omitempty" name:"Type"`
	// 事件状态过滤，取值范围["recover","alarm","-"]，分别表示已恢复、未恢复、无状态

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 项目ID过滤

	Project []*string `json:"Project,omitempty" name:"Project"`
	// 告警状态配置过滤，1表示已配置，0表示未配置

	IsAlarmConfig *int64 `json:"IsAlarmConfig,omitempty" name:"IsAlarmConfig"`
	// 按更新时间排序，ASC表示升序，DESC表示降序，默认DESC

	TimeOrder *string `json:"TimeOrder,omitempty" name:"TimeOrder"`
	// 起始时间，默认一天前的时间戳

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，默认当前时间戳

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 页偏移量，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页返回的数量，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeProductEventListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductEventListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductEventListEvents struct {
	// 事件ID

	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
	// 事件中文名

	EventCName *string `json:"EventCName,omitempty" name:"EventCName"`
	// 事件英文名

	EventEName *string `json:"EventEName,omitempty" name:"EventEName"`
	// 事件简称

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 产品中文名

	ProductCName *string `json:"ProductCName,omitempty" name:"ProductCName"`
	// 产品英文名

	ProductEName *string `json:"ProductEName,omitempty" name:"ProductEName"`
	// 产品简称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 是否支持告警

	SupportAlarm *int64 `json:"SupportAlarm,omitempty" name:"SupportAlarm"`
	// 事件类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 更新时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 实例对象信息

	Dimensions []*DescribeProductEventListEventsDimensions `json:"Dimensions,omitempty" name:"Dimensions"`
	// 实例对象附加信息

	AdditionMsg []*DescribeProductEventListEventsDimensions `json:"AdditionMsg,omitempty" name:"AdditionMsg"`
	// 是否配置告警

	IsAlarmConfig *int64 `json:"IsAlarmConfig,omitempty" name:"IsAlarmConfig"`
	// 策略信息

	GroupInfo []*DescribeProductEventListEventsGroupInfo `json:"GroupInfo,omitempty" name:"GroupInfo"`
}

type DataPoint struct {
	// 实例对象维度组合

	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions"`
	// 时间戳数组，表示那些时间点有数据，缺失的时间戳，没有数据点，可以理解为掉点了

	Timestamps []*float64 `json:"Timestamps,omitempty" name:"Timestamps"`
	// 监控值数组，该数组和Timestamps一一对应

	Values []*float64 `json:"Values,omitempty" name:"Values"`
}

type DescribeProductEventListDimensions struct {
	// 维度名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 维度值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeProductEventListEventsGroupInfo struct {
	// 策略ID

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 策略名

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type Instance struct {
	// 实例的维度组合

	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions"`
}

type DescribeProductEventListEventsDimensions struct {
	// 维度名（英文）

	Key *string `json:"Key,omitempty" name:"Key"`
	// 维度名（中文）

	Name *string `json:"Name,omitempty" name:"Name"`
	// 维度值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type GetDataBaradMetricRequest struct {
	*tchttp.BaseRequest

	// 命名空间，每个云产品会有一个命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 监控统计周期。默认为取值为300，单位为s

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 起始时间，如 2018-01-01 00:00:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，默认为当前时间。 endTime不能小于startTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 实例对象的维度组合

	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 视图名称

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
}

func (r *GetDataBaradMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDataBaradMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMonitorDataV2018Request struct {
	*tchttp.BaseRequest

	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 周期

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 起始时间，如2018-09-22T19:51:23+08:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，如2018-09-22T20:51:23+08:00，默认为当前时间。 EndTime不能小于StartTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 实例对象的维度组合，格式为key-value键值对形式的集合。

	Instances []*Instance `json:"Instances,omitempty" name:"Instances"`
}

func (r *GetMonitorDataV2018Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMonitorDataV2018Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductEventListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件列表

		Events []*DescribeProductEventListEvents `json:"Events,omitempty" name:"Events"`
		// 事件统计

		OverView *DescribeProductEventListOverView `json:"OverView,omitempty" name:"OverView"`
		// 事件总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductEventListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDataBaradMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDataBaradMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDataBaradMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMonitorDataV2018Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 统计周期

		Period *uint64 `json:"Period,omitempty" name:"Period"`
		// 指标名

		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
		// 数据点数组

		DataPoints []*DataPoint `json:"DataPoints,omitempty" name:"DataPoints"`
		// 开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMonitorDataV2018Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMonitorDataV2018Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductEventListOverView struct {
	// 状态变更的事件数量

	StatusChangeAmount *int64 `json:"StatusChangeAmount,omitempty" name:"StatusChangeAmount"`
	// 告警状态未配置的事件数量

	UnConfigAlarmAmount *int64 `json:"UnConfigAlarmAmount,omitempty" name:"UnConfigAlarmAmount"`
	// 异常事件数量

	UnNormalEventAmount *int64 `json:"UnNormalEventAmount,omitempty" name:"UnNormalEventAmount"`
	// 未恢复的事件数量

	UnRecoverAmount *int64 `json:"UnRecoverAmount,omitempty" name:"UnRecoverAmount"`
}

type Dimension struct {
	// 实例维度名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实例维度值

	Value *string `json:"Value,omitempty" name:"Value"`
}
