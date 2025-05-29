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

package v20170312

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DescribeCustomerCapacityTopChangesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomerCapacityTopChangesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerCapacityTopChangesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskStorageDeviceTypesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDiskStorageDeviceTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskStorageDeviceTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSetHostSideIOStatRequest struct {
	*tchttp.BaseRequest

	// 开始时间，例："2021-03-17 00:00:00"

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 结束时间，例："2021-03-17 00:10:00"

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// 指定从秒级监控系统或分钟级监控系统查询数据，MINUTE表示从分钟级监控系统查询数据；SECOND表示从秒级监控系统查询数据

	Type *string `json:"Type,omitempty" name:"Type"`
	// 监控数据的周期，单位秒。取值范围：查询秒级监控时可传入10或60，表示查询周期为10s或60s的监控数据；查询分钟级监控数据，只能传入60

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 过滤条件。<br><li>set-uuid - Array of String - 是否必填：否 -（过滤条件）指定setUuid查询监控数据

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeSetHostSideIOStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSetHostSideIOStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallDiskStorageDepotCellRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储集群ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// CELL物理服务器IP，至少需要三台物理服务器作为一组CELL，要求所有机器必须为同一种机型，且ip数必须为3的倍数

	DeviceIps []*string `json:"DeviceIps,omitempty" name:"DeviceIps"`
	// 是否以dry run模式运行；dry run模式下只校验参数、机型是否合法，不执行实际的安装操作

	DryRun *uint64 `json:"DryRun,omitempty" name:"DryRun"`
	// 1:格式化硬盘,0:不格式化硬盘，建议取值为1

	CleanSign *uint64 `json:"CleanSign,omitempty" name:"CleanSign"`
	// 同一组件使用机器的容灾级别要求，传入此参数，上架前，会进行容灾级别的检验，不满足不会进行上架，取值范围：<br><li>any: 无要求<br><li>motherMachine: 跨母机，仅针对虚拟机生效<br><li>rack: 跨机架<br><li>switch: 跨交换机

	DisasterTolerance *string `json:"DisasterTolerance,omitempty" name:"DisasterTolerance"`
}

func (r *InstallDiskStorageDepotCellRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallDiskStorageDepotCellRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindUserFromDiskStoragePoolGroupRequest struct {
	*tchttp.BaseRequest

	// 属性ID

	DepotAttributeId *uint64 `json:"DepotAttributeId,omitempty" name:"DepotAttributeId"`
	// 用户AppId列表

	AppIds []*uint64 `json:"AppIds,omitempty" name:"AppIds"`
	// 可用区类别，可以只解绑部分可用区；默认解绑全部可用区

	ZoneIds []*uint64 `json:"ZoneIds,omitempty" name:"ZoneIds"`
	// 要解除绑定的云硬盘资源池(pool-xxxx)，该参数与DepotAttributeId参数只能有一个进行传递。

	PoolGroup *string `json:"PoolGroup,omitempty" name:"PoolGroup"`
}

func (r *UnbindUserFromDiskStoragePoolGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindUserFromDiskStoragePoolGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallDiskStorageDepotCellResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务ID，可在运维任务页面查看任务执行结果。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UninstallDiskStorageDepotCellResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallDiskStorageDepotCellResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Price struct {
	// 预支费用的原价，单位：元。

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
	// 预支费用的折扣价，单位：元。

	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
	// 后付费云盘折扣单价，单位：元。

	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitempty" name:"UnitPriceDiscount"`
	// 后付费云盘原单价，单位：元。

	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`
	// 高精度后付费云盘原单价, 单位：元

	UnitPriceHigh *string `json:"UnitPriceHigh,omitempty" name:"UnitPriceHigh"`
	// 高精度预付费云盘预支费用的原价, 单位：元

	OriginalPriceHigh *string `json:"OriginalPriceHigh,omitempty" name:"OriginalPriceHigh"`
	// 高精度预付费云盘预支费用的折扣价, 单位：元

	DiscountPriceHigh *string `json:"DiscountPriceHigh,omitempty" name:"DiscountPriceHigh"`
	// 高精度后付费云盘折扣单价, 单位：元

	UnitPriceDiscountHigh *string `json:"UnitPriceDiscountHigh,omitempty" name:"UnitPriceDiscountHigh"`
	// 后付费云盘的计价单元，取值范围：<br><li>HOUR：表示后付费云盘的计价单元是按小时计算。

	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`
}

type DescribeCbsAlarmEventsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>region - Array of String - 是否必填：否 -（过滤条件）按region过滤。<br><li>alarm-object - Array of String - 是否必填：否 -（过滤条件）按告警对象过滤。（DISK:云硬盘）。<br><li>event-type - Array of String - 是否必填：否 -（过滤条件）按事件类型过滤。<br><li>alarm-state - Array of String - 是否必填：否 -（过滤条件）按告警状态过滤。（ALARMING:告警中 | RECOVER: 已恢复 | SHIELD: 已屏蔽）

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCbsAlarmEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCbsAlarmEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDepotTransferTaskOverviewRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储池ID列表

	DepotIds []*string `json:"DepotIds,omitempty" name:"DepotIds"`
}

func (r *DescribeDepotTransferTaskOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepotTransferTaskOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorEsIndexStorageTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存储云盘监控数据的ES索引保留时间

		EsIndexStorageTimeSet []*EsIndexStorageTime `json:"EsIndexStorageTimeSet,omitempty" name:"EsIndexStorageTimeSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMonitorEsIndexStorageTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorEsIndexStorageTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmEvent struct {
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 告警对象

	AlarmObject *string `json:"AlarmObject,omitempty" name:"AlarmObject"`
	// 告警状态

	AlarmState *string `json:"AlarmState,omitempty" name:"AlarmState"`
	// 事件类型

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 云盘类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云盘UUID

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
	// 存储仓库ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 存储仓库名称

	DepotName *string `json:"DepotName,omitempty" name:"DepotName"`
	// 持续时长

	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
	// 宿主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// AppID

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 告警详情

	AlarmDetail *string `json:"AlarmDetail,omitempty" name:"AlarmDetail"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DiskStoragePoolGroupBound struct {
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 云硬盘类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 账号AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 绑定类型

	BoundType *string `json:"BoundType,omitempty" name:"BoundType"`
	// 资源池组名称

	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
	// 是否可见。

	Visible *uint64 `json:"Visible,omitempty" name:"Visible"`
}

type CreateSnapshotBoxUpgradeTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前该快照集群所有升级任务的事务ID，用于唯一标识一次快照集群的升级任务。

		TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSnapshotBoxUpgradeTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSnapshotBoxUpgradeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoCloseAndOpenBlockRateRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储池ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
}

func (r *DescribeAutoCloseAndOpenBlockRateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoCloseAndOpenBlockRateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSetBlockUsageDayStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 时间列表，按照时间升序排列

		DateSet []*string `json:"DateSet,omitempty" name:"DateSet"`
		// block利用率列表，和时间列表中的数据一一对应

		BlockUsageSet []*float64 `json:"BlockUsageSet,omitempty" name:"BlockUsageSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSetBlockUsageDayStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSetBlockUsageDayStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoCloseAndOpenBlockRateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoCloseAndOpenBlockRateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoCloseAndOpenBlockRateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SimpleInstance struct {
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例UUID

	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
	// 实例别名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type ModifyDiskAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDiskAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiskAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskStorageDepotAttribute struct {
	// 属性ID

	AttrId *uint64 `json:"AttrId,omitempty" name:"AttrId"`
	// 属性名称

	AttrName *string `json:"AttrName,omitempty" name:"AttrName"`
	// 属性配置

	Attributes []*string `json:"Attributes,omitempty" name:"Attributes"`
	// 属性创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 属性创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 存储池类型名称

	DepotAttributeName *string `json:"DepotAttributeName,omitempty" name:"DepotAttributeName"`
	// 存储池属性列表

	DepotAttributes []*string `json:"DepotAttributes,omitempty" name:"DepotAttributes"`
	// 存储池数量

	DepotCount *uint64 `json:"DepotCount,omitempty" name:"DepotCount"`
	// 存储池已使用大小

	DepotSizeCreated *uint64 `json:"DepotSizeCreated,omitempty" name:"DepotSizeCreated"`
	// 存储池超售率

	DepotSizeOverSoldTotal *uint64 `json:"DepotSizeOverSoldTotal,omitempty" name:"DepotSizeOverSoldTotal"`
	// 存储池总大小

	DepotSizeTotal *uint64 `json:"DepotSizeTotal,omitempty" name:"DepotSizeTotal"`
	// 存储池已使用大小

	DepotSizeUsed *uint64 `json:"DepotSizeUsed,omitempty" name:"DepotSizeUsed"`
	// 存储池售卖状态列表

	DepotSoldStateSet []*string `json:"DepotSoldStateSet,omitempty" name:"DepotSoldStateSet"`
	// 存储池状态列表

	DepotStateSet []*string `json:"DepotStateSet,omitempty" name:"DepotStateSet"`
	// 资源池

	DepotStoragePoolGroup *string `json:"DepotStoragePoolGroup,omitempty" name:"DepotStoragePoolGroup"`
	// 资源池名称

	DepotStoragePoolGroupName *string `json:"DepotStoragePoolGroupName,omitempty" name:"DepotStoragePoolGroupName"`
	// 资源池

	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
	// 存储池类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 最大可创建的云盘大小

	MaxDiskCapacity *uint64 `json:"MaxDiskCapacity,omitempty" name:"MaxDiskCapacity"`
	// 最小可创建的云盘大小

	MinDiskCapacity *uint64 `json:"MinDiskCapacity,omitempty" name:"MinDiskCapacity"`
	// 修改者

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 修改人

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 备注信息

	Note *string `json:"Note,omitempty" name:"Note"`
	// 存储池类型

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 存储池类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 资源池绑定的云盘数量

	UserBoundCount *uint64 `json:"UserBoundCount,omitempty" name:"UserBoundCount"`
	// 存储池所在的可用区列表

	ZoneSet []*uint64 `json:"ZoneSet,omitempty" name:"ZoneSet"`
}

type ProductInfo struct {
	// 属性名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 属性值。

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeBalanceDepotTasksRequest struct {
	*tchttp.BaseRequest

	// 存储池ID列表。不传该参数时，默认查询当前地域的所有存储池

	DepotIds []*uint64 `json:"DepotIds,omitempty" name:"DepotIds"`
}

func (r *DescribeBalanceDepotTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBalanceDepotTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDepotCellDeviceConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDepotCellDeviceConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDepotCellDeviceConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceConfigUseInfo struct {
	// 该机型正在使用的region。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 该机型正在使用的可用区列表。

	ZoneList []*string `json:"ZoneList,omitempty" name:"ZoneList"`
}

type Filters struct {
	// app-id uin nickname

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DescribeDiskResourceUsageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日期列表

		Dates []*string `json:"Dates,omitempty" name:"Dates"`
		// 云硬盘大小列表

		TotalSizes []*int64 `json:"TotalSizes,omitempty" name:"TotalSizes"`
		// 云硬盘数量列表

		TotalCounts []*int64 `json:"TotalCounts,omitempty" name:"TotalCounts"`
		// 统计资源类型

		ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
		// 统计周期

		Period *string `json:"Period,omitempty" name:"Period"`
		// 统计的起始时间

		StartDate *string `json:"StartDate,omitempty" name:"StartDate"`
		// 统计类型

		StatisticType *string `json:"StatisticType,omitempty" name:"StatisticType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskResourceUsageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskResourceUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceDiskStorageDepotCellResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务ID。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 任务描述。

		TaskDescription *uint64 `json:"TaskDescription,omitempty" name:"TaskDescription"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceDiskStorageDepotCellResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceDiskStorageDepotCellResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TabletPairStatistic struct {
	// 总小表对数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// normal小表对数

	NormalCount *uint64 `json:"NormalCount,omitempty" name:"NormalCount"`
	// free小表对数

	FreeCount *uint64 `json:"FreeCount,omitempty" name:"FreeCount"`
	// used小表对数

	UsedCount *uint64 `json:"UsedCount,omitempty" name:"UsedCount"`
	// abnormal小表对数

	AbnormalCount *uint64 `json:"AbnormalCount,omitempty" name:"AbnormalCount"`
	// onedead小表对数

	OneDeadCount *uint64 `json:"OneDeadCount,omitempty" name:"OneDeadCount"`
	// twodead小表对数

	TwoDeadCount *uint64 `json:"TwoDeadCount,omitempty" name:"TwoDeadCount"`
	// born小表对数

	BornCount *uint64 `json:"BornCount,omitempty" name:"BornCount"`
	// errfree小表对数

	ErrFreeCount *uint64 `json:"ErrFreeCount,omitempty" name:"ErrFreeCount"`
	// errborn小表对数

	ErrBornCount *uint64 `json:"ErrBornCount,omitempty" name:"ErrBornCount"`
}

type ModifyAutoSnapshotPolicyAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoSnapshotPolicyAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoSnapshotPolicyAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DepotConfigItem struct {
	// 存储池配置项名称。

	ConfigItem *string `json:"ConfigItem,omitempty" name:"ConfigItem"`
	// 配置项所属的组件名称。

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
	// 配置项所处的配置文件。

	ConfigFile *string `json:"ConfigFile,omitempty" name:"ConfigFile"`
	// 配置项在配置文件中的路径。

	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`
	// 配置值。

	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`
	// 配置项取值范围。

	Range *string `json:"Range,omitempty" name:"Range"`
	// 配置项取值的单位。

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 配置项取值的类型。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 配置项作用的IP列表。

	DeviceIps []*string `json:"DeviceIps,omitempty" name:"DeviceIps"`
}

type DescribeDiskStorageDepotsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>depot-id - Array of Integer - 是否必填：否 -（过滤条件）按存储池ID过滤。<br><li>depot-name - Array of String - 是否必填：否 -（过滤条件）按存储池名称过滤。<br><li>zone-id - Array of Integer - 是否必填：否 -（过滤条件）按可用区ID过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 控制台搜索用

	InnerSearch *string `json:"InnerSearch,omitempty" name:"InnerSearch"`
}

func (r *DescribeDiskStorageDepotsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskStorageDepotsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiskStoragePoolGroupRequest struct {
	*tchttp.BaseRequest

	// 存储池属性ID，该参数已废弃。

	DepotAttributeId *uint64 `json:"DepotAttributeId,omitempty" name:"DepotAttributeId"`
	// 存储池组名称

	DiskStoragePoolGroupName *string `json:"DiskStoragePoolGroupName,omitempty" name:"DiskStoragePoolGroupName"`
	// 要修改的云硬盘资源池(pool-xxxx)。

	PoolGroup *string `json:"PoolGroup,omitempty" name:"PoolGroup"`
}

func (r *ModifyDiskStoragePoolGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiskStoragePoolGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorData struct {
	// 存储池UUID

	DepotUuid *string `json:"DepotUuid,omitempty" name:"DepotUuid"`
	// 存储池ID

	DepotId *int64 `json:"DepotId,omitempty" name:"DepotId"`
	// 存储池名称

	DepotName *string `json:"DepotName,omitempty" name:"DepotName"`
	// 存储池类型

	DepotDiskType *string `json:"DepotDiskType,omitempty" name:"DepotDiskType"`
	// 存储池属性名

	DepotAttributeName *string `json:"DepotAttributeName,omitempty" name:"DepotAttributeName"`
	// 存储池属性列表

	DepotAttributes []*string `json:"DepotAttributes,omitempty" name:"DepotAttributes"`
	// 可用区ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// IO性能详情

	Metric *Metric `json:"Metric,omitempty" name:"Metric"`
	// 存储池cell节点

	CellIp *string `json:"CellIp,omitempty" name:"CellIp"`
	// 存储池cell节点组uuid

	CellGroupUuid *string `json:"CellGroupUuid,omitempty" name:"CellGroupUuid"`
	// 云盘或存储池物理盘名称。

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 硬盘组ID

	DiskPairId *string `json:"DiskPairId,omitempty" name:"DiskPairId"`
	// 硬盘smcd

	Smcd *int64 `json:"Smcd,omitempty" name:"Smcd"`
	// 硬盘SN号

	DiskSN *string `json:"DiskSN,omitempty" name:"DiskSN"`
	// 云盘实例ID

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 云盘UUID

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
	// 云盘挂载的实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 云盘挂载的实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 云盘挂载的实例UUID

	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
	// app id

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
}

type ResOpLog struct {
	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 操作执行时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 操作结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// AppID

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 任务描述

	TaskDescription *string `json:"TaskDescription,omitempty" name:"TaskDescription"`
	// 父任务ID

	ParentTaskId *uint64 `json:"ParentTaskId,omitempty" name:"ParentTaskId"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 云硬盘UUID

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
	// 云服务器UUID

	InsUuid *string `json:"InsUuid,omitempty" name:"InsUuid"`
	// 快照UUID

	SnapshotUuid *string `json:"SnapshotUuid,omitempty" name:"SnapshotUuid"`
	// 定期快照策略ID

	AspId *string `json:"AspId,omitempty" name:"AspId"`
	// 任务输入

	Input *string `json:"Input,omitempty" name:"Input"`
	// 任务输出

	Output *string `json:"Output,omitempty" name:"Output"`
	// 任务错误码

	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 云硬盘ID

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 云硬盘名称

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 快照名称

	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
	// 快照大小

	SnapshotSize *uint64 `json:"SnapshotSize,omitempty" name:"SnapshotSize"`
	// 快照ID

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 任务执行情况

	SubTaskStatistic *SubTaskStatistic `json:"SubTaskStatistic,omitempty" name:"SubTaskStatistic"`
}

type SharePermission struct {
	// 快照分享的时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 分享的账号Id

	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`
}

type InstallDiskStorageDepotMasterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务ID。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InstallDiskStorageDepotMasterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallDiskStorageDepotMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallBoxDeviceRequest struct {
	*tchttp.BaseRequest

	// 快照集群ID

	BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
	// 设备IP

	DeviceIp *string `json:"DeviceIp,omitempty" name:"DeviceIp"`
	// 机器类型，取值范围：manager, scheduler, transfer

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
	// 0-正常卸载 1-强制卸载，强制卸载进程，不保证卸载成功都会删除其运营系统中的数据，在裁撤此box时使用

	ForceUninstall *uint64 `json:"ForceUninstall,omitempty" name:"ForceUninstall"`
}

func (r *UninstallBoxDeviceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallBoxDeviceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageRecycleRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageRecycleRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageRecycleRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PhysicalDisk struct {
	// 物理盘所在节点IP

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 物理盘状态。<br><li>NORMAL: 正常<br><li>DOWN：隔离<br><li>DEAD：剔除

	State *string `json:"State,omitempty" name:"State"`
	// 磁盘已使用容量，单位MB

	JournalSize *uint64 `json:"JournalSize,omitempty" name:"JournalSize"`
	// 磁盘快照使用容量，单位MB

	SnapshotSize *uint64 `json:"SnapshotSize,omitempty" name:"SnapshotSize"`
	// 磁盘SN号

	PhysicalDiskSN *string `json:"PhysicalDiskSN,omitempty" name:"PhysicalDiskSN"`
	// 磁盘总容量，单位MB

	StorageDiskSize *uint64 `json:"StorageDiskSize,omitempty" name:"StorageDiskSize"`
	// 磁盘替换状态。。<br><li>NO_NEED_REPLACE: 正常<br><li>REPLACE_MAIN_DISK：主存替换<br><li>REPLACE_CACHE_DISK：缓存替换<br><li>REPLACING：替换中

	DiskReplaceState *string `json:"DiskReplaceState,omitempty" name:"DiskReplaceState"`
	// 当盘的状态为剔除时，本字段标识具体的事件ID，可通过接口DescribeDepotRepotEvents查询事件详情；默认值-1，表示没有剔除事件

	DiskMasterEventId *int64 `json:"DiskMasterEventId,omitempty" name:"DiskMasterEventId"`
	// 用于换盘时，标识换盘操作是否支持热插拔磁盘；如果取值为true，表示支持热插拔，换盘时无需重启机器；取值为false，表示不支持热插拔，换盘时需要重启机器

	SupportDiskHotSwap *bool `json:"SupportDiskHotSwap,omitempty" name:"SupportDiskHotSwap"`
	// 盘符

	OSDiskWwn *string `json:"OSDiskWwn,omitempty" name:"OSDiskWwn"`
	// 盘名

	OSDiskName *string `json:"OSDiskName,omitempty" name:"OSDiskName"`
	// SN号

	OSDiskSn *string `json:"OSDiskSn,omitempty" name:"OSDiskSn"`
}

type AttachDetail struct {
	// 实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例已挂载数据盘的数量。

	AttachedDiskCount *uint64 `json:"AttachedDiskCount,omitempty" name:"AttachedDiskCount"`
	// 实例最大可挂载数据盘的数量。

	MaxAttachCount *uint64 `json:"MaxAttachCount,omitempty" name:"MaxAttachCount"`
}

type DescribeDiskStorageDepotsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的云硬盘存储仓库数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 云硬盘存储仓库列表

		DiskStorageDepotSet []*DiskStorageDepot `json:"DiskStorageDepotSet,omitempty" name:"DiskStorageDepotSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskStorageDepotsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskStorageDepotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMonitorEsIndexStorageTimeRequest struct {
	*tchttp.BaseRequest

	// 指定修改具体索引的保留时间设置；取值范围：<br><li>qemu_disk_iostat_ten_sec: 云盘10秒级监控数据<br><li>qemu_disk_iostat_min: 云盘1分钟级监控数据<br><li>qemu_disk_iostat_five_min: 云盘5分钟级监控数据<br><li>qemu_disk_info_ten_sec: 存储池10秒级监控数据<br><li>qemu_disk_info_min: 存储池分钟级监控数据<br><li>qemu_disk_info_five_min: 存储斌5分钟级监控数据

	IndexNames []*string `json:"IndexNames,omitempty" name:"IndexNames"`
	// 关闭索引的时间，即设置把几天前的索引关闭，变为冷数据，减轻ES集群的负载

	CloseTime *uint64 `json:"CloseTime,omitempty" name:"CloseTime"`
	// 删除索引的时间，即设置把几天前的索引删除，节省ES的存储空间

	DeleteTime *uint64 `json:"DeleteTime,omitempty" name:"DeleteTime"`
}

func (r *ModifyMonitorEsIndexStorageTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMonitorEsIndexStorageTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunClusterUpgradeTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行本次升级任务的异步des任务ID。如果DryRun为1，不会生成des异步任务，此时TaskId取值为空

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 发起任务的描述

		TaskDescription *string `json:"TaskDescription,omitempty" name:"TaskDescription"`
		// 环境检查的结果

		CheckState *string `json:"CheckState,omitempty" name:"CheckState"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunClusterUpgradeTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunClusterUpgradeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceDiskStorageDepotDiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 换盘的任务ID。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceDiskStorageDepotDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceDiskStorageDepotDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceDiskStorageZKResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务ID。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 任务描述。

		TaskDescription *string `json:"TaskDescription,omitempty" name:"TaskDescription"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceDiskStorageZKResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceDiskStorageZKResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlowControlAction struct {
	// 本次流控的具体指标，取值范围：ReadThrough，ReadIops，WriteThrough，WriteIops

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 本次流控的类型，取值范围：<br<li>Recover: 恢复流控<br><li>TopSuppress: 整体流控<br><li>BurstSuppress: 突发流量流控

	Action *string `json:"Action,omitempty" name:"Action"`
}

type DescribeSnapshotBoxOverViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 快照集群资源概况

		SnapshotBoxOverView *SnapshotBoxResourceOverView `json:"SnapshotBoxOverView,omitempty" name:"SnapshotBoxOverView"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotBoxOverViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotBoxOverViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecoverDiskRequest struct {
	*tchttp.BaseRequest

	// 需恢复云盘的id

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

func (r *RecoverDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecoverDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSetHostSideIOStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 时间序列

		Times []*string `json:"Times,omitempty" name:"Times"`
		// 存储池IO监控详情

		SetHostSideIOStatSet []*SetIOStat `json:"SetHostSideIOStatSet,omitempty" name:"SetHostSideIOStatSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSetHostSideIOStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSetHostSideIOStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSnapshotBoxConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSnapshotBoxConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSnapshotBoxConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnapshotGroup struct {
	// 快照组ID

	SnapshotGroupId *string `json:"SnapshotGroupId,omitempty" name:"SnapshotGroupId"`
	// 快照组类型。NORMAL: 普通快照组，非一致性快照

	SnapshotGroupType *string `json:"SnapshotGroupType,omitempty" name:"SnapshotGroupType"`
	// 快照组是否包含系统盘快照

	ContainRootSnapshot *bool `json:"ContainRootSnapshot,omitempty" name:"ContainRootSnapshot"`
	// 快照组包含的快照ID列表

	SnapshotIdSet []*string `json:"SnapshotIdSet,omitempty" name:"SnapshotIdSet"`
	// 快照组状态。<br><li>NORMAL: 正常<br><li>CREATING:创建中<br><li>DELETED:已删除<br><li>FAILED:创建失败<br><li>DISMISS:已解散<br><li>ROLLBACKING:回滚中

	SnapshotGroupState *string `json:"SnapshotGroupState,omitempty" name:"SnapshotGroupState"`
	// 快照组创建进度

	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最新修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 快照组关联的镜像列表

	Images []*string `json:"Images,omitempty" name:"Images"`
	// 快照组名称

	SnapshotGroupName *string `json:"SnapshotGroupName,omitempty" name:"SnapshotGroupName"`
	// 快照组关联的镜像数量

	ImageCount *uint64 `json:"ImageCount,omitempty" name:"ImageCount"`
	// 最后一次执行的异步任务执行结果

	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`
	// 最后一次执行的异步任务对应的API RequestId

	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitempty" name:"LatestOperationRequestId"`
	// 最后一次执行的异步任务操作

	LatestOperation *string `json:"LatestOperation,omitempty" name:"LatestOperation"`
	// 最后一次异步任务操作失败的资源列表

	OperationFailedResourceIdSet []*string `json:"OperationFailedResourceIdSet,omitempty" name:"OperationFailedResourceIdSet"`
	// 最后一次异步操作成功的资源列表

	OperationSuccessResourceIdSet []*string `json:"OperationSuccessResourceIdSet,omitempty" name:"OperationSuccessResourceIdSet"`
	// 最后一次异步操作结果的提示码

	ErrorPrompt *string `json:"ErrorPrompt,omitempty" name:"ErrorPrompt"`
	// 用户appId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type SnapshotSizeInfo struct {
	// 总快照大小，单位MB

	SnapAllSize *float64 `json:"SnapAllSize,omitempty" name:"SnapAllSize"`
	// 已标记为删除，cos中暂未清理的快照大小，单位MB

	SnapDeletingSize *float64 `json:"SnapDeletingSize,omitempty" name:"SnapDeletingSize"`
	// 状态正常的快照大小，单位MB

	SnapNormalSize *float64 `json:"SnapNormalSize,omitempty" name:"SnapNormalSize"`
	// 正在创建中的快照大小，单位MB

	SnapRunningSize *float64 `json:"SnapRunningSize,omitempty" name:"SnapRunningSize"`
	// 已删除待回收的快照容量。

	DeletedSnapshotSize *float64 `json:"DeletedSnapshotSize,omitempty" name:"DeletedSnapshotSize"`
	// 待merge快照容量。

	WaitMergeSnapshotSize *float64 `json:"WaitMergeSnapshotSize,omitempty" name:"WaitMergeSnapshotSize"`
	// 已merge待回收快照容量。

	MergedSnapshotSize *float64 `json:"MergedSnapshotSize,omitempty" name:"MergedSnapshotSize"`
	// 处于已merge待回收状态超过3天的快照容量。

	Merged3DaysSnapshotSize *float64 `json:"Merged3DaysSnapshotSize,omitempty" name:"Merged3DaysSnapshotSize"`
}

type StorageDepot struct {
	// 存储池ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 存储池名称

	DepotName *string `json:"DepotName,omitempty" name:"DepotName"`
	// 存储池所在的cbs zone id

	CbsZoneId *uint64 `json:"CbsZoneId,omitempty" name:"CbsZoneId"`
	// 存储池所在的zone id

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储池属性ID

	DepotAttributeId *uint64 `json:"DepotAttributeId,omitempty" name:"DepotAttributeId"`
	// 存储池属性名称

	DepotAttributeName *string `json:"DepotAttributeName,omitempty" name:"DepotAttributeName"`
	// 存储池属性列表，形如：[ "ssd" ]

	DepotAttributes []*string `json:"DepotAttributes,omitempty" name:"DepotAttributes"`
	// 存储池状态。NORMAL：正常；FAULT：异常

	DepotState *string `json:"DepotState,omitempty" name:"DepotState"`
	// 存储池售卖状态。NORMAL:已开打；DISABLE：已关闭；PREPARE：准备中；FAULT：异常；DESTROYING：待删除

	DepotSoldState *string `json:"DepotSoldState,omitempty" name:"DepotSoldState"`
}

type RecycleDisksRequest struct {
	*tchttp.BaseRequest

	// 需要回收的云盘uuid列表

	DiskUuids []*string `json:"DiskUuids,omitempty" name:"DiskUuids"`
}

func (r *RecycleDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonModule struct {
	// common组件的首次安装时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// common组件所在的节点ip

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// common最新安装时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// common组件的名称

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// cmmon组件的版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type ModifyReplaceDiskTaskRequest struct {
	*tchttp.BaseRequest

	// 换盘任务ID。

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 换盘任务状态。取值范围，<br><li>check_finish：完成确认<br><li>stop：任务终止

	JobStatus *string `json:"JobStatus,omitempty" name:"JobStatus"`
}

func (r *ModifyReplaceDiskTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReplaceDiskTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDepotConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDepotConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDepotConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DepotOverviewByDiskTypeSetInAZ struct {
	// 可用区的ZoneId；

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储仓库概览数据；

	DepotOverviewByDiskTypeSetInZone *DepotOverview `json:"DepotOverviewByDiskTypeSetInZone,omitempty" name:"DepotOverviewByDiskTypeSetInZone"`
}

type CreateSnapshotBoxResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建的快照集群ID

		BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSnapshotBoxResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSnapshotBoxResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskDetailsRequest struct {
	*tchttp.BaseRequest

	// 云盘UUID列表

	DiskUuids []*string `json:"DiskUuids,omitempty" name:"DiskUuids"`
}

func (r *DescribeDiskDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskStorageDepotAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存储仓库属性配置列表

		DiskStorageDepotAttributeSet []*DiskStorageDepotAttribute `json:"DiskStorageDepotAttributeSet,omitempty" name:"DiskStorageDepotAttributeSet"`
		// 有效的存储仓库属性配置数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskStorageDepotAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskStorageDepotAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDepotCellDeviceConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 更新的存储机型配置详情

		DepotCellDeviceConfig *DepotCellDeviceConfig `json:"DepotCellDeviceConfig,omitempty" name:"DepotCellDeviceConfig"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDepotCellDeviceConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDepotCellDeviceConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBoxAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBoxAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBoxAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskTask struct {
	// 用户appId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 任务创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 云盘实例ID

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 任务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务入参详情

	TaskInput *string `json:"TaskInput,omitempty" name:"TaskInput"`
	// 任务输出详情

	TaskOutput *string `json:"TaskOutput,omitempty" name:"TaskOutput"`
	// 任务状态。（CREATED：等待执行 | DISCARDED：执行失败 | EXECUTING：执行中 | DONE：已完成 | FAILED：执行失败）

	TaskState *string `json:"TaskState,omitempty" name:"TaskState"`
	// 任务类型

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
	// zone id

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 集群ID。

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 任务属于运营端还是租户端。

	TaskSource *string `json:"TaskSource,omitempty" name:"TaskSource"`
}

type ReplaceDiskStorageDepotMasterRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储集群ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 出现故障，待替换slave节点IP

	OldDeviceIp *string `json:"OldDeviceIp,omitempty" name:"OldDeviceIp"`
	// 新slave节点IP

	NewDeviceIp *string `json:"NewDeviceIp,omitempty" name:"NewDeviceIp"`
	// 是否执行强制替换。0：需要成功卸载故障slave节点才进行后续的替换；1：尝试卸载，不计成败

	ForceReplace *uint64 `json:"ForceReplace,omitempty" name:"ForceReplace"`
	// 是否以dry run模式运行；0: 非dry run模式； 1： dry run模式，默认非dry run模式。dry run模式下只校验参数、机型是否合法，不执行实际的替换操作

	DryRun *uint64 `json:"DryRun,omitempty" name:"DryRun"`
	// 同一组件使用机器的容灾级别要求，传入此参数，上架前，会进行容灾级别的检验，不满足不会进行上架，取值范围：<br><li>any: 无要求<br><li>motherMachine: 跨母机，仅针对虚拟机生效<br><li>rack: 跨机架<br><li>switch: 跨交换机

	DisasterTolerance *string `json:"DisasterTolerance,omitempty" name:"DisasterTolerance"`
}

func (r *ReplaceDiskStorageDepotMasterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceDiskStorageDepotMasterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComponentPerformanceData struct {
	// 监控数据的时刻列表

	TimestampSet []*string `json:"TimestampSet,omitempty" name:"TimestampSet"`
	// 具体监控指标值列表

	MetricValueSet []*MetricValue `json:"MetricValueSet,omitempty" name:"MetricValueSet"`
}

type DiskZkCluster struct {
	// cbs底层zone id

	CbsZoneId *int64 `json:"CbsZoneId,omitempty" name:"CbsZoneId"`
	// zone id

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 已接入的存储池数量

	DepotCnt *uint64 `json:"DepotCnt,omitempty" name:"DepotCnt"`
	// 最大可接入存储池数量

	MaxDepotNum *uint64 `json:"MaxDepotNum,omitempty" name:"MaxDepotNum"`
	// 最后修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 最后修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// zk名称

	ZkName *string `json:"ZkName,omitempty" name:"ZkName"`
	// zk类型。master: 存储池zk；snapshoht:快照zk

	ZkType *string `json:"ZkType,omitempty" name:"ZkType"`
	// zk集群ID

	ZKClusterId *uint64 `json:"ZKClusterId,omitempty" name:"ZKClusterId"`
	// ZK集群状态。<br><li>NORMAL: 正常<br><li>FAULT：异常<br><li>INSTALLING：上架中<br><li>UNINSTALLING：下架中

	ZkState *string `json:"ZkState,omitempty" name:"ZkState"`
	// ZK节点详情

	ZKNodeDetailSet []*ZKNodeDetail `json:"ZKNodeDetailSet,omitempty" name:"ZKNodeDetailSet"`
	// 已接入当前zk的存储池列表

	StorageDepotSet []*StorageDepot `json:"StorageDepotSet,omitempty" name:"StorageDepotSet"`
	// zk机器的机型

	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
	// 已接入当前zk的快照集群列表

	SnapshotBoxInfoSet []*SnapshotBoxInfo `json:"SnapshotBoxInfoSet,omitempty" name:"SnapshotBoxInfoSet"`
	// zk集群的vip，仅针对“跨可用区级zk集群”才有vip，cbs独立上架的zk集群该值为null

	ZKVip *string `json:"ZKVip,omitempty" name:"ZKVip"`
}

type SnapshotBoxInfo struct {
	// 快照集群ID

	BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
	// 快照集群名称

	BoxName *string `json:"BoxName,omitempty" name:"BoxName"`
	// 快照集群状态。<br><li>READ_WRITE: 可读写<br><li>READ_ONLY：只读<br><li>CLOSED：已关闭<br><li>UNINSTALLING：上架中<br><li>INSTALLING：上架中

	BoxState *string `json:"BoxState,omitempty" name:"BoxState"`
}

type DescribeDepotTransferTaskOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存储池小表迁移概况详情

		TransferTaskOverviewSet []*TransferTaskOverview `json:"TransferTaskOverviewSet,omitempty" name:"TransferTaskOverviewSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDepotTransferTaskOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepotTransferTaskOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotResourceUsageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 快照资源详情

		SnapshotResourceSet []*SnapshotResource `json:"SnapshotResourceSet,omitempty" name:"SnapshotResourceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotResourceUsageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotResourceUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDiskStoragePoolGroupRequest struct {
	*tchttp.BaseRequest

	// 是否是CFS专属云硬盘资源池组

	IsCFS *bool `json:"IsCFS,omitempty" name:"IsCFS"`
	// 云硬盘类型，CLOUD_BASIC/CLOUD_PREMIUM/CLOUD_SSD

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云硬盘资源池组名称

	DiskStoragePoolGroupName *string `json:"DiskStoragePoolGroupName,omitempty" name:"DiskStoragePoolGroupName"`
	// 云硬盘资源池组属性名称，只能使用字母和数字，以及“_”符合

	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
	// 云硬盘存储仓库ID列表；创建云硬盘资源池组的同时会修改对应的云硬盘存储仓库的属性

	DepotIds []*uint64 `json:"DepotIds,omitempty" name:"DepotIds"`
	// 云硬盘类型列表，CLOUD_BASIC/CLOUD_PREMIUM/CLOUD_SSD

	DiskTypes []*string `json:"DiskTypes,omitempty" name:"DiskTypes"`
}

func (r *CreateDiskStoragePoolGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDiskStoragePoolGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDiskStorageDepotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDiskStorageDepotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDiskStorageDepotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskPair struct {
	// 物理盘详情

	PhysicalDisks []*PhysicalDisk `json:"PhysicalDisks,omitempty" name:"PhysicalDisks"`
	// 磁盘小表对ID

	DiskPairId *uint64 `json:"DiskPairId,omitempty" name:"DiskPairId"`
	// 磁盘对应的smcd号

	Smcd *uint64 `json:"Smcd,omitempty" name:"Smcd"`
	// 小表状态详情

	DiskTpStateInfo *DiskTpStateInfo `json:"DiskTpStateInfo,omitempty" name:"DiskTpStateInfo"`
}

type DiskStorageDepot struct {
	// 存储系统可用区ID

	CbsZoneId *uint64 `json:"CbsZoneId,omitempty" name:"CbsZoneId"`
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// ZooKeeper集群ID

	ZKClusterId *int64 `json:"ZKClusterId,omitempty" name:"ZKClusterId"`
	// 存储集群创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 存储集群创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 存储集群属性配置ID

	DepotAttributeId *uint64 `json:"DepotAttributeId,omitempty" name:"DepotAttributeId"`
	// 存储集群属性名称

	DepotAttributeName *string `json:"DepotAttributeName,omitempty" name:"DepotAttributeName"`
	// 存储集群ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 存储集群名称

	DepotName *string `json:"DepotName,omitempty" name:"DepotName"`
	// 存储集群已创建的云硬盘大小

	DepotSizeCreated *uint64 `json:"DepotSizeCreated,omitempty" name:"DepotSizeCreated"`
	// 存储集群实际的总存储大小

	DepotSizeTotal *uint64 `json:"DepotSizeTotal,omitempty" name:"DepotSizeTotal"`
	// 存储集群实际上已使用的总存储大小

	DepotSizeUsed *uint64 `json:"DepotSizeUsed,omitempty" name:"DepotSizeUsed"`
	// 存储集群售卖状态，取值范围：<br><li>NORMAL: 已打开<br><li>DISABLE: 已关闭<br><li>PREPARE:上架中<br><li>DESTROYING: 下架中

	DepotSoldState *string `json:"DepotSoldState,omitempty" name:"DepotSoldState"`
	// 存储集群状态，取值范围：<br><li>NORMAL: 正常<br><li>FAULT:异常

	DepotState *string `json:"DepotState,omitempty" name:"DepotState"`
	// 存储集群UUID

	DepotUuid *string `json:"DepotUuid,omitempty" name:"DepotUuid"`
	// 已创建的云硬盘数量

	DiskNumCreated *uint64 `json:"DiskNumCreated,omitempty" name:"DiskNumCreated"`
	// 可创建的云硬盘总数量

	DiskNumTotal *uint64 `json:"DiskNumTotal,omitempty" name:"DiskNumTotal"`
	// 超卖比（%）

	OverSold *uint64 `json:"OverSold,omitempty" name:"OverSold"`
	// 快照集群ID

	SnapshotBoxId *int64 `json:"SnapshotBoxId,omitempty" name:"SnapshotBoxId"`
	// 存储池当前版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 当前存储池支持升级到哪些版本

	SupportUpgradeVersions []*string `json:"SupportUpgradeVersions,omitempty" name:"SupportUpgradeVersions"`
	// 当前存储池的升级状态详情

	UpgradeDetail *UpgradeDetail `json:"UpgradeDetail,omitempty" name:"UpgradeDetail"`
	// 存储池云盘类型，取值范围: <br><li>CLOUD_SATA: 普通云盘<br><li>CLOUD_PREMIUM: 高性能云盘<br><li>CLOUD_SSD: SSD云盘

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 协议

	DepotProtocol *string `json:"DepotProtocol,omitempty" name:"DepotProtocol"`
	// 资源池名称

	DepotStoragePoolGroupName *string `json:"DepotStoragePoolGroupName,omitempty" name:"DepotStoragePoolGroupName"`
	// 存储池属性列表

	DepotAttributes []*string `json:"DepotAttributes,omitempty" name:"DepotAttributes"`
	// 存储池类型

	DepotDiskType *string `json:"DepotDiskType,omitempty" name:"DepotDiskType"`
	// 所在资源池

	DepotStoragePoolGroup *string `json:"DepotStoragePoolGroup,omitempty" name:"DepotStoragePoolGroup"`
	// 存储池快照占用的容量

	DepotSizeSnapshotUsed *uint64 `json:"DepotSizeSnapshotUsed,omitempty" name:"DepotSizeSnapshotUsed"`
	// 存储池小表对详情

	TabletPairStatistic *TabletPairStatistic `json:"TabletPairStatistic,omitempty" name:"TabletPairStatistic"`
	// 组件机器数量详情

	DeviceSet *DeviceTypeCount `json:"DeviceSet,omitempty" name:"DeviceSet"`
	// 存储池小表详情

	TabletStatistic *TabletStatistic `json:"TabletStatistic,omitempty" name:"TabletStatistic"`
	// 存储池的平均block利用率

	BlockUsageRate *float64 `json:"BlockUsageRate,omitempty" name:"BlockUsageRate"`
	// 存储池大最block利用率

	MaxBlockUsageRate *float64 `json:"MaxBlockUsageRate,omitempty" name:"MaxBlockUsageRate"`
}

type SnapshotBoxResourceOverView struct {
	// 快照集群ID

	BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
	// 已做快照的云盘数量

	SnapDiskNum *uint64 `json:"SnapDiskNum,omitempty" name:"SnapDiskNum"`
	// 各状态快照大小

	SnapshotSizeInfo *SnapshotSizeInfo `json:"SnapshotSizeInfo,omitempty" name:"SnapshotSizeInfo"`
	// 各状态快照数量

	SnapshotStateInfo *SnapshotStateInfo `json:"SnapshotStateInfo,omitempty" name:"SnapshotStateInfo"`
}

type DescribeDepotRepotEventsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>regoin - Array of String - 是否必填：否 -（过滤条件）按region过滤。<br><li>event-type - Array of String - 是否必填：否 -（过滤条件）按事件类型过滤。<br><li>cell-ip - Array of String - 是否必填：否 -（过滤条件）按cell ip过滤。<br><li>depot-id - Array of Integer - 是否必填：否 -（过滤条件）按存储池ID过滤。<br><li>depot-name - Array of String - 是否必填：否 -（过滤条件）按存储池名称过滤。<br><li>occur-time-range - Array of String - 是否必填：否 -（过滤条件）按事件发生的时间范围过滤，格式[start_time, end_time], 例如：[2020-01-16 19:55:23, 2020-01-17 19:55:23]。<br><li>replace-disk-type - Array of String - 是否必填：否 -（过滤条件）按事件的换盘类型过滤。（NO_NEED_REPLACE：无需换盘 | REPLACE_MAIN_DISK：主存换盘 | REPLACE_CACHE_DISK:缓存换盘 | NEED_MANUAL_CHECK:人工确认）<br><li>event-id - Array of Integer - 是否必填：否 -（过滤条件）按事件ID过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDepotRepotEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepotRepotEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskIOStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 时间序列

		Times []*string `json:"Times,omitempty" name:"Times"`
		// 监控数据详情

		DiskIOStatSet []*DiskIOStat `json:"DiskIOStatSet,omitempty" name:"DiskIOStatSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskIOStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskIOStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskConfigForSale struct {
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 云硬盘类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 当前售卖配置是否生效

	Available *uint64 `json:"Available,omitempty" name:"Available"`
	// 用户可见白名单：不在该白名单中的用户无法看到该种售卖类型的云硬盘

	VisibleWhiteKey *string `json:"VisibleWhiteKey,omitempty" name:"VisibleWhiteKey"`
	// 售卖开关，后台自动更新维护。`SOLD_OUT`：资源售罄；`ALL_USER_VALID`：所有用户都可购买；`VIP_USER_VALID`：仅VIP用户可购买

	AutoSoldFlag *string `json:"AutoSoldFlag,omitempty" name:"AutoSoldFlag"`
	// `FORCE_SOLD_OUT`：强制售罄，所有用户都不可购买；`AUTO_SOLD_OUT`：由系统资源余量决定；`WHITE_SOLD_ONLY`：仅白名单用户可购买；`VIP_SOLD_ONLY`：仅VIP用户可购买

	ForceSoldFlag *string `json:"ForceSoldFlag,omitempty" name:"ForceSoldFlag"`
	// 可售卖资源低于此阈值时，全部用户显示售罄

	NormalSoldOutLowerLimit *uint64 `json:"NormalSoldOutLowerLimit,omitempty" name:"NormalSoldOutLowerLimit"`
	// `WHITE_SOLD_ONLY`模式下有效，用于仅对指定用户售卖

	ForceSoldWhiteKey *string `json:"ForceSoldWhiteKey,omitempty" name:"ForceSoldWhiteKey"`
	// 可售卖资源低于此阈值时，仅向VIP用户售卖

	VipSoldOutLowerLimit *uint64 `json:"VipSoldOutLowerLimit,omitempty" name:"VipSoldOutLowerLimit"`
	// 可售卖云硬盘容量，后台自动更新维护

	AvailCapacity *uint64 `json:"AvailCapacity,omitempty" name:"AvailCapacity"`
	// 可用存储集群数量；后台自动更新维护

	AvailSetNumber *uint64 `json:"AvailSetNumber,omitempty" name:"AvailSetNumber"`
	// 可售卖云硬盘数量；后台自动更新维护

	AvailDiskNumber *uint64 `json:"AvailDiskNumber,omitempty" name:"AvailDiskNumber"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 主建id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 最后告警时间

	LastAlarmTime *string `json:"LastAlarmTime,omitempty" name:"LastAlarmTime"`
	// 规格字义详情

	DiskModelConfigForSaleSet []*DiskModelConfigForSale `json:"DiskModelConfigForSaleSet,omitempty" name:"DiskModelConfigForSaleSet"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DiskMaintenanceTask struct {
	// 当前步骤详情具体描述

	StepInfo *string `json:"StepInfo,omitempty" name:"StepInfo"`
	// 任务最后执行时间

	LastRunningTime *string `json:"LastRunningTime,omitempty" name:"LastRunningTime"`
	// 创建人

	Author *string `json:"Author,omitempty" name:"Author"`
	// 任务描述信息

	JobDesc *string `json:"JobDesc,omitempty" name:"JobDesc"`
	// 任务完成时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 任务错误码

	ExitCode *string `json:"ExitCode,omitempty" name:"ExitCode"`
	// 任务状态, running 正在运行， pause 运行中止/异常， normal_finish 任务完成

	Status *string `json:"Status,omitempty" name:"Status"`
	// 当前步骤描述

	StepDesc *string `json:"StepDesc,omitempty" name:"StepDesc"`
	// 任务创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 任务启动时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 出错信息

	Err *string `json:"Err,omitempty" name:"Err"`
	// 任务步骤总数

	StepAll *string `json:"StepAll,omitempty" name:"StepAll"`
	// 当前所处步骤

	StepNow *string `json:"StepNow,omitempty" name:"StepNow"`
	// 任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 任务类型

	JobType *string `json:"JobType,omitempty" name:"JobType"`
}

type DeleteSnapshotsRequest struct {
	*tchttp.BaseRequest

	// 要删除的快照ID列表，可通过[DescribeSnapshots](/document/product/362/15647)查询。

	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
	// 需要删除后台快照时，传入该参数值为true，表示要删除租户的后台快照

	DeleteBackendSnap *bool `json:"DeleteBackendSnap,omitempty" name:"DeleteBackendSnap"`
}

func (r *DeleteSnapshotsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSnapshotsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskConfigForSaleRequest struct {
	*tchttp.BaseRequest

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 云硬盘付费模式，PREPAID与POSTPAID_BY_HOUR

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 云硬盘类型，CLOUD_BASIC、CLOUD_PREMIUM、CLOUD_SSD

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// SYSTEM_DISK、DATA_DISK

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 云硬盘规格类型，base、large、xlarge

	CFS_DATA_DISK *string `json:"CFS_DATA_DISK,omitempty" name:"CFS_DATA_DISK"`
}

func (r *DescribeDiskConfigForSaleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskConfigForSaleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskMigrateTasksRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>disk-id - Array of String - 是否必填：否 -（过滤条件）按云盘实例ID过滤。<br><li>creator - Array of String - 是否必填：否 -（过滤条件）按任务创建人过滤。<br><li>task-id - Array of Integer - 是否必填：否 -（过滤条件）按任务ID过滤。<br><li>status - Array of String - 是否必填：否 -（过滤条件）按任务状态过滤。<br><li>app-id - Array of Integer - 是否必填：否 -（过滤条件）按用户appId过滤。<br><li>source-depot-id - Array of String - 是否必填：否 -（过滤条件）按源存储池ID过滤。<br><li>destination-depot-id - Array of Integer - 是否必填：否 -（过滤条件）按目标存储池ID过滤。<br><li>zone-id - Array of Integer - 是否必填：否 -（过滤条件）按可用区过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为20，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDiskMigrateTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskMigrateTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotBoxConfigsRequest struct {
	*tchttp.BaseRequest

	// 快照集群ID

	BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
	// 需要查询的配置列表，可传入时，会默认返回所以支持查询的配置

	ConfigItems []*string `json:"ConfigItems,omitempty" name:"ConfigItems"`
	// 配置所在的组件类型，取值范围：manager, scheduler, tranfser, trsf_ctrl, trsf_proxy

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
}

func (r *DescribeSnapshotBoxConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotBoxConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GoodsDetail struct {
	// 购买或续费云盘的时长。

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 产品ID。

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 云盘类型，系统盘或数据盘。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 产品信息描述。

	ProductInfo []*ProductInfo `json:"ProductInfo,omitempty" name:"ProductInfo"`
	// 云盘介质类型。

	MediumType *string `json:"MediumType,omitempty" name:"MediumType"`
	// 时长“TimeSpan”的单位。

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 云盘大小。

	CbsSize *uint64 `json:"CbsSize,omitempty" name:"CbsSize"`
}

type DescribeDepotTransferTasksRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>task-id - Array of Integer - 是否必填：否 -（过滤条件）按taskId过滤。<br><li>task-state - Array of String - 是否必填：否 -（过滤条件）按任务状态过滤，当前只有RUNNING取值。<br><li>depot-id - Array of Integer - 是否必填：否 -（过滤条件）按存储池ID过滤。<br><li>depot-name - Array of String - 是否必填：否 -（过滤条件）按存储池名称过滤。<br><li>trsf-type - Array of String - 是否必填：否 -（过滤条件）按任务类型过滤，（MOVE: 迁移 | RECOVERY: 恢复）。<br><li>src-cell-ip - Array of String - 是否必填：否 -（过滤条件）按源cell ip过滤。<br><li>src-smcd - Array of Integer - 是否必填：否 -（过滤条件）按源smcd过滤。<br><li>dst-cell-ip - Array of String - 是否必填：否 -（过滤条件）按目的cell ip过滤。<br><li>dst-smcd - Array of Integer - 是否必填：否 -（过滤条件）按目的smcd过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDepotTransferTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepotTransferTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotBoxsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>box-id - Array of Integer - 是否必填：否 -（过滤条件）按快照集群ID过滤。<br><li>box-name - Array of String - 是否必填：否 -（过滤条件）按快照集群名称过滤。<br><li>box-state - Array of Integer - 是否必填：否 -（过滤条件）按快照集群状态过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeSnapshotBoxsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotBoxsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeDetail struct {
	// 集群升级状态。取值范围：<br><li>NOT_START: 不在升级中<br><li>WAIT_CHECK: 等待检查<br><li>CHECK_FINISH: 检查完成<br><li>RUNNING: 升级中<br><li>SUCCESS: 升级成功

	UpgradeState *string `json:"UpgradeState,omitempty" name:"UpgradeState"`
	// 集群升级的事务ID，用于DescribeClusterUpgradeTasks接口，能够查询到当前集群所有组件的升级任务。如果集群不在升级中，取值为空

	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`
	// 如果集群在升级中，该字段表示集群升级的目标版本。如果不在升级，取值为空

	NewVersion *string `json:"NewVersion,omitempty" name:"NewVersion"`
}

type CreateDiskStorageDepotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新创建的云硬盘存储集群ID

		DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDiskStorageDepotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDiskStorageDepotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDiskStoragePoolGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存储池属性列表

		DepotAttributeIds []*uint64 `json:"DepotAttributeIds,omitempty" name:"DepotAttributeIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDiskStoragePoolGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDiskStoragePoolGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSnapshotBoxRequest struct {
	*tchttp.BaseRequest

	// 快照ZooKeeper集群ID

	ZKClusterId *uint64 `json:"ZKClusterId,omitempty" name:"ZKClusterId"`
	// 快照集群名称

	BoxName *string `json:"BoxName,omitempty" name:"BoxName"`
	// 存储快照数据的COS存储桶名称

	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`
	// COS接入层网关地址或网关前面的负载均衡地址, 可以是ip形式如192.168.1.x 这样， 也可以是域名如api.cos.x.com 这样

	CosHostName *string `json:"CosHostName,omitempty" name:"CosHostName"`
	// COS secret ID

	CosSecretId *string `json:"CosSecretId,omitempty" name:"CosSecretId"`
	// COS secret key

	CosSecretKey *string `json:"CosSecretKey,omitempty" name:"CosSecretKey"`
	// 存储快照元数据的COS存储桶名称

	CosBucketNameMeta *string `json:"CosBucketNameMeta,omitempty" name:"CosBucketNameMeta"`
}

func (r *CreateSnapshotBoxRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSnapshotBoxRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceDiskStorageDepotCellRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储仓库ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 出现故障，待替换cell机器的原IP

	OldDeviceIp *string `json:"OldDeviceIp,omitempty" name:"OldDeviceIp"`
	// 新cell机器IP

	NewDeviceIp *string `json:"NewDeviceIp,omitempty" name:"NewDeviceIp"`
	// 是否执行强制替换。0：需要成功卸载故障cell才进行后续的替换；1：尝试卸载，不计成败

	ForceReplace *string `json:"ForceReplace,omitempty" name:"ForceReplace"`
	// 是否格式化机器。0：强制格式化；1：不格式化

	CleanSign *uint64 `json:"CleanSign,omitempty" name:"CleanSign"`
	// 是否以dry run模式运行；0: 非dry run模式； 1： dry run模式，默认非dry run模式。dry run模式下只校验参数、机型是否合法，不执行实际的替换操作

	DryRun *uint64 `json:"DryRun,omitempty" name:"DryRun"`
	// 同一组件使用机器的容灾级别要求，传入此参数，上架前，会进行容灾级别的检验，不满足不会进行上架，取值范围：<br><li>any: 无要求<br><li>motherMachine: 跨母机，仅针对虚拟机生效<br><li>rack: 跨机架<br><li>switch: 跨交换机

	DisasterTolerance *string `json:"DisasterTolerance,omitempty" name:"DisasterTolerance"`
}

func (r *ReplaceDiskStorageDepotCellRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceDiskStorageDepotCellRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindUserFromDiskStoragePoolGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindUserFromDiskStoragePoolGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindUserFromDiskStoragePoolGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnapshotStateInfo struct {
	// 总快照数量

	SnapAllNum *uint64 `json:"SnapAllNum,omitempty" name:"SnapAllNum"`
	// 已标记为删除，cos中暂未清理的快照数量

	SnapDeletingNum *uint64 `json:"SnapDeletingNum,omitempty" name:"SnapDeletingNum"`
	// 状态正常的快照数量

	SnapNormalNum *uint64 `json:"SnapNormalNum,omitempty" name:"SnapNormalNum"`
	// 创建中的快照数量

	SnapRunningNum *uint64 `json:"SnapRunningNum,omitempty" name:"SnapRunningNum"`
	// 已删除待回收的快照数量。

	DeletedSnapshotNumber *uint64 `json:"DeletedSnapshotNumber,omitempty" name:"DeletedSnapshotNumber"`
	// 待merge快照数量。

	WaitMergeSnapshotNumber *uint64 `json:"WaitMergeSnapshotNumber,omitempty" name:"WaitMergeSnapshotNumber"`
	// 已merge待回收快照数量。

	MergedSnapshotNumber *uint64 `json:"MergedSnapshotNumber,omitempty" name:"MergedSnapshotNumber"`
	// 处于已merge待回收状态超过3天的快照数量。

	Merged3DaysSnapshotNumber *uint64 `json:"Merged3DaysSnapshotNumber,omitempty" name:"Merged3DaysSnapshotNumber"`
}

type CreateDepotUpgradeTasksRequest struct {
	*tchttp.BaseRequest

	// 存储池所在的可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储池ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 存储池当前版本，可通过接口DescribeDiskStorageDepots查询，见Version字段

	CurVersion *string `json:"CurVersion,omitempty" name:"CurVersion"`
	// 升级的目标版本。当前存储池支持升级的目标版本列表可通过接口DescribeDiskStorageDepots查询，见SupportUpgradeVersions字段

	NewVersion *string `json:"NewVersion,omitempty" name:"NewVersion"`
}

func (r *CreateDepotUpgradeTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDepotUpgradeTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterUpgradeTasksRequest struct {
	*tchttp.BaseRequest

	// 当集群正在升级时才需要使用该参数，表示集群升级的事务ID，通过该参数可以查到一次集群升级的所有任务集合。存储池可通过接口DescribeDiskStorageDepots查询，见UpgradeDetail字段，快照集群可通过DescribeSnapshotBoxs接口查询。如果TransactionId为空，说明当前集群不在升级中

	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`
	// 存储池ID，查询存储池升级历史时可传入该参数

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 快照集群ID，查询快照集群升级历史时可传入该参数

	BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
	// 过滤条件。<br><li>device-ip - Array of String - 是否必填：否 -（过滤条件）按机器IP过滤。<br><li>module-name - Array of String - 是否必填：否 -（过滤条件）按模块名称过滤。<br><li>task-id - Array of Integer - 是否必填：否 -（过滤条件）按任务ID过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 传入该参数，表示要查询集群的升级历史，此时，只会返回集群升级成功的任务，且一种组件只会返回一条记录；取值只能为1

	QueryUpgradeHistory *uint64 `json:"QueryUpgradeHistory,omitempty" name:"QueryUpgradeHistory"`
}

func (r *DescribeClusterUpgradeTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterUpgradeTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnapshotCopyResult struct {
	// 跨地复制的目标地域。

	DestinationRegion *string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
	// 复制到目标地域的新快照ID。

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

type DescribeCustomerAccountsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，从0开始

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCustomerAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerAccountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type A struct {
	// 1

	Aa []*int64 `json:"Aa,omitempty" name:"Aa"`
}

type DiskConfigSet struct {
	// 系统盘或数据盘

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// CLOUD_BASIC、CLOUD_PREMIUM、CLOUD_SSD

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云盘大小

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 快照ID

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

type ResourceOperationLog struct {
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 子任务名称

	SubTaskName *string `json:"SubTaskName,omitempty" name:"SubTaskName"`
	// 账户AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 任务类型

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
	// 任务输入

	TaskInput *string `json:"TaskInput,omitempty" name:"TaskInput"`
	// 任务输出

	TaskOutput *string `json:"TaskOutput,omitempty" name:"TaskOutput"`
	// 任务发生时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 云盘类型。

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 云盘计费类型。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 操作任务名称。

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 操任务英文名称。

	OperationEN *string `json:"OperationEN,omitempty" name:"OperationEN"`
	// 操作人uin。

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 操作人子账号uin。

	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
	// 云盘类型。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 云盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 任务详情。

	Message *string `json:"Message,omitempty" name:"Message"`
	// 云盘uuid。

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
	// 操作结果状态码。

	Code *int64 `json:"Code,omitempty" name:"Code"`
}

type CreateDepotUpgradeTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前该存储池所有升级任务的事务ID，用于唯一标识一次存储池的升级任务

		TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDepotUpgradeTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDepotUpgradeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReplaceDiskTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 换盘的任务详情。

		ReplaceDiskTaskSet []*ReplaceDiskTask `json:"ReplaceDiskTaskSet,omitempty" name:"ReplaceDiskTaskSet"`
		// 符合条件的任务总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReplaceDiskTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReplaceDiskTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReviveDiskStorageDepotCellRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储集群ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 需要恢复的cell节点IP

	DeviceIp *string `json:"DeviceIp,omitempty" name:"DeviceIp"`
	// 恢复的对象。<br><li>DEVICE：整机恢复；<br><li>DISK：单盘恢复，此时，必需传入参数DiskPairId，指定恢复的盘号。<br><li> DISK_PAIR：指定小表恢复，此时需要传入TabletPairId或TabletIndex

	ObjectType *string `json:"ObjectType,omitempty" name:"ObjectType"`
	// 单盘恢复时传入，表示盘对应的smcd值

	Smcd *uint64 `json:"Smcd,omitempty" name:"Smcd"`
	// 指定小表恢复时传入

	TabletPairId *uint64 `json:"TabletPairId,omitempty" name:"TabletPairId"`
	// 指定小表恢复时传入

	TabletIndex *uint64 `json:"TabletIndex,omitempty" name:"TabletIndex"`
}

func (r *ReviveDiskStorageDepotCellRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReviveDiskStorageDepotCellRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDiskStorageDepotRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储仓库Cell机型，根据机型确定存储仓库可用的云盘介质类型。比如TS60/TSC10对应普通云盘，Y0-TSH10-00对应高效云盘，TS80/Y0-TS80-10对应SSD云盘

	CellDeviceType *string `json:"CellDeviceType,omitempty" name:"CellDeviceType"`
	// 存储仓库属性，比如`cnas`表示该存储仓库用于CFS存储

	DepotAttribute []*string `json:"DepotAttribute,omitempty" name:"DepotAttribute"`
	// 存储仓库使用的ZooKeeper集群ID

	ZKClusterId *uint64 `json:"ZKClusterId,omitempty" name:"ZKClusterId"`
	// 存储仓库属性ID

	DepotAttributeId *uint64 `json:"DepotAttributeId,omitempty" name:"DepotAttributeId"`
	// 存储仓库超卖百分比，默认100

	OverSold *uint64 `json:"OverSold,omitempty" name:"OverSold"`
	// 快照集群ID

	SnapshotBoxId *uint64 `json:"SnapshotBoxId,omitempty" name:"SnapshotBoxId"`
	// 存储仓库别名

	DepotName *string `json:"DepotName,omitempty" name:"DepotName"`
}

func (r *CreateDiskStorageDepotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDiskStorageDepotRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallCapacityServerRequest struct {
	*tchttp.BaseRequest

	// 容量server所属快照box id

	BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
	// 要卸载的容量server机器ip

	DeviceIp *string `json:"DeviceIp,omitempty" name:"DeviceIp"`
	// 是否强制卸载，0表示非强制卸载；1表示强制卸载

	ForceUninstall *uint64 `json:"ForceUninstall,omitempty" name:"ForceUninstall"`
}

func (r *UninstallCapacityServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallCapacityServerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BalanceDiskStorageDepotRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储池ID

	DepotId *int64 `json:"DepotId,omitempty" name:"DepotId"`
	// `START_BALANCE`：开始小表均衡；`CANCEL_BALANCE`：取消小表均衡

	BalanceType *string `json:"BalanceType,omitempty" name:"BalanceType"`
}

func (r *BalanceDiskStorageDepotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BalanceDiskStorageDepotRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskTasksRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>app-id - Array of Integer - 是否必填：否 -（过滤条件）按用户appId过滤。<br><li>task-id - Array of Integer - 是否必填：否 -（过滤条件）按任务ID过滤。<br><li>disk-id - Array of String - 是否必填：否 -（过滤条件）按云盘实例ID过滤。<br><li> - Array of Integer - 是否必填：否 -（过滤条件）按zoneId过滤。<br><li>task-state - Array of String - 是否必填：否 -（过滤条件）按任务状态过滤。（CREATED：等待执行 | DISCARDED：执行失败 | EXECUTING：执行中 | DONE：已完成 | FAILED：执行失败）<br><li>task-type - Array of String - 是否必填：否 -（过滤条件）按任务类型过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDiskTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DepotCellDeviceConfig struct {
	// 存储机型，如Y0-SH11-25G

	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
	// 存储机型主存盘数量

	DiskNum *uint64 `json:"DiskNum,omitempty" name:"DiskNum"`
	// 存储机型主存盘大小，单位TB

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 存储池类型，取值范围：<br><li>sata: 普通存储池<br><li>fusion: 高性能存储池<br><li>ssd: SSD存储池

	DepotMedia *string `json:"DepotMedia,omitempty" name:"DepotMedia"`
	// 主存盘接口类型，取值范围: <br><li>sata: sata接口<br><li>nvme: nvme接口

	PrimaryInfType *string `json:"PrimaryInfType,omitempty" name:"PrimaryInfType"`
	// 缓存盘接口类型，取值范围: <br><li>nvme: nvme接口<br><li>none: 当前机型没有缓存盘时取值为null

	CacheInfType *string `json:"CacheInfType,omitempty" name:"CacheInfType"`
	// 机型添加时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 机型记录更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 当前机型是否可用。0表示不可用，1表示可用

	Available *uint64 `json:"Available,omitempty" name:"Available"`
	// 当前机型是否支持磁盘热插拔。true表示支持热插拔盘；false表示不支持热插拔盘

	SupportDiskHotSwap *bool `json:"SupportDiskHotSwap,omitempty" name:"SupportDiskHotSwap"`
	// 存储机型缓存盘数量

	CacheDiskNum *uint64 `json:"CacheDiskNum,omitempty" name:"CacheDiskNum"`
	// 存储机型缓存盘大小，可能为小数，所以类型为字符串类型，单位TB

	CacheDiskSize *string `json:"CacheDiskSize,omitempty" name:"CacheDiskSize"`
	// 存储机型的网络类型，取值范围：10Gb, 25Gb

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 存储机型的使用情况。“InUse”为1表示现网正在使用；InUse为0表示没有使用。

	InUse *int64 `json:"InUse,omitempty" name:"InUse"`
	// 存储机型的使用信息。

	UseInfo *DeviceConfigUseInfo `json:"UseInfo,omitempty" name:"UseInfo"`
	// 机型支持上架存储池的副本策略，取值范围，THREE_COPIES：三副本存储，EC(9+3): 9+3 EC存储

	CopyStrategy []*string `json:"CopyStrategy,omitempty" name:"CopyStrategy"`
}

type DiskAttachedInstanceInfo struct {
	// 云盘挂载时间。

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 实例UUID。

	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
}

type DiskDepotComponentDetailTransferPair struct {
	// 上架时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 上架人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 设备类型

	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
	// ip地址

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 最后升级时间

	LastUpgradeTime *string `json:"LastUpgradeTime,omitempty" name:"LastUpgradeTime"`
	// trsf_pair id

	PairId *int64 `json:"PairId,omitempty" name:"PairId"`
	// 端口号

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 机架

	RackName *string `json:"RackName,omitempty" name:"RackName"`
	// 组件状态

	State *string `json:"State,omitempty" name:"State"`
	// 上联交换

	SwitchName *string `json:"SwitchName,omitempty" name:"SwitchName"`
	// 升级状态

	UpgradeState *string `json:"UpgradeState,omitempty" name:"UpgradeState"`
	// 组件版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 可用区id

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 组件类型

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 机器的cpu型号

	DeviceCpuModel *string `json:"DeviceCpuModel,omitempty" name:"DeviceCpuModel"`
}

type ModifyDisksMonitorAlarmRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDisksMonitorAlarmRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDisksMonitorAlarmRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceDiskStorageDepotDiskRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储池ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 更换物理盘所在的存储池节点IP

	CellIp *string `json:"CellIp,omitempty" name:"CellIp"`
	// 更换物理盘对应的Smcd

	Smcd *uint64 `json:"Smcd,omitempty" name:"Smcd"`
	// 换盘类型。<br><li>REPLACE_MAIN_DISK：主存换盘；REPLACE_CACHE_DISK：缓存换盘

	ReplaceDiskType *string `json:"ReplaceDiskType,omitempty" name:"ReplaceDiskType"`
	// 换盘是否触发迁移。NEED_TRANSFER：触发迁移；SKIP_TRANSFER：不触发迁移

	AutoTransferFlag *string `json:"AutoTransferFlag,omitempty" name:"AutoTransferFlag"`
	// 标志是否DryRun，1位DryRun，0为非DryRun。

	DryRun *uint64 `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *ReplaceDiskStorageDepotDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceDiskStorageDepotDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindDiskStorageDepotFromDiskStoragePoolGroupRequest struct {
	*tchttp.BaseRequest

	// 存储池仓库ID列表

	DepotIds []*uint64 `json:"DepotIds,omitempty" name:"DepotIds"`
	// 存储池属性ID，指定此参数时会修改为此属性ID

	DepotAttributeId *uint64 `json:"DepotAttributeId,omitempty" name:"DepotAttributeId"`
}

func (r *UnbindDiskStorageDepotFromDiskStoragePoolGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindDiskStorageDepotFromDiskStoragePoolGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EsIndexStorageTime struct {
	// ES索引的名称

	IndexName *string `json:"IndexName,omitempty" name:"IndexName"`
	// 设置几天后把索引关闭

	CloseTime *uint64 `json:"CloseTime,omitempty" name:"CloseTime"`
	// 设置几天后把索引删除

	DeleteTime *uint64 `json:"DeleteTime,omitempty" name:"DeleteTime"`
	// 配置的最新更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type CreateAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest

	// 定期快照的执行策略。

	Policy []*Policy `json:"Policy,omitempty" name:"Policy"`
	// 要创建的定期快照策略名。不传则默认为“未命名”。最大长度不能超60个字节。

	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" name:"AutoSnapshotPolicyName"`
	// 是否激活定期快照策略，FALSE表示未激活，TRUE表示激活，默认为TRUE。

	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`
	// 通过该定期快照策略创建的快照是否永久保留。FALSE表示非永久保留，TRUE表示永久保留，默认为FALSE。

	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`
	// 通过该定期快照策略创建的快照保留天数，默认保留7天，该参数不可与`IsPermanent`参数冲突，即若定期快照策略设置为永久保留，本参数应置0。

	RetentionDays *uint64 `json:"RetentionDays,omitempty" name:"RetentionDays"`
	// 是否创建定期快照的执行策略。TRUE表示只需获取首次开始备份的时间，不实际创建定期快照策略，FALSE表示创建，默认为FALSE。

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *CreateAutoSnapshotPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitializeDiskStorageDepotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 初始化存储集群的任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitializeDiskStorageDepotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitializeDiskStorageDepotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceDiskStorageZKRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// ZooKeeper集群ID

	ZKClusterId *uint64 `json:"ZKClusterId,omitempty" name:"ZKClusterId"`
	// 出现故障，待替换的flower节点IP

	OldDeviceIp *string `json:"OldDeviceIp,omitempty" name:"OldDeviceIp"`
	// 新的ZK节点IP

	NewDeviceIp *string `json:"NewDeviceIp,omitempty" name:"NewDeviceIp"`
	// 是否执行强制替换。0：需要成功卸载故障zk才进行后续的替换；1：尝试卸载，不计成败

	ForceReplace *uint64 `json:"ForceReplace,omitempty" name:"ForceReplace"`
	// 是否以dry run模式运行；0: 非dry run模式； 1： dry run模式，默认非dry run模式。dry run模式下只校验参数、机型是否合法，不执行实际的替换操作

	DryRun *uint64 `json:"DryRun,omitempty" name:"DryRun"`
	// 同一组件使用机器的容灾级别要求，传入此参数，上架前，会进行容灾级别的检验，不满足不会进行上架，取值范围：<br><li>any: 无要求<br><li>motherMachine: 跨母机，仅针对虚拟机生效<br><li>rack: 跨机架<br><li>switch: 跨交换机

	DisasterTolerance *string `json:"DisasterTolerance,omitempty" name:"DisasterTolerance"`
}

func (r *ReplaceDiskStorageZKRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceDiskStorageZKRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallDiskZKClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务ID，可在运维任务页面查看任务执行结果。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UninstallDiskZKClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallDiskZKClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterUpgradeTask struct {
	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务添加时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// CBS底层可用区ID

	CbsZoneId *uint64 `json:"CbsZoneId,omitempty" name:"CbsZoneId"`
	// 集群ID

	ClusterId *uint64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群升级状态。取值范围：<br><li>WAIT_CHECK: 等待检查<br><li>CHECK_FINISH: 检查完成<br><li>RUNNING: 升级中<br><li>SUCCESS: 升级成功

	ClusterState *string `json:"ClusterState,omitempty" name:"ClusterState"`
	// 集群类型。取值范围：<br><li>product-cbs-depot: 存储池集群<br><li>product-cbs-snap: 快照集群

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 任务创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 当前步骤详情

	CurDetail *string `json:"CurDetail,omitempty" name:"CurDetail"`
	// 当前步骤

	CurStep *uint64 `json:"CurStep,omitempty" name:"CurStep"`
	// 升级机器的IP

	DeviceIp *string `json:"DeviceIp,omitempty" name:"DeviceIp"`
	// 升级机器的机型

	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
	// 任务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 任务记录修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 升级的模块名。取值范围：<br><li>cell: cell模块<br><li>dbtrsf: dbtrsf模块<br><li>master: master模块<br><li>loadctrl: loadctrl模块<br><li>manager: 快照manager模块<br><li>scheduler: 快照scheduler模块<br><li>transfer: 快照transfer模块<br><li>trsf_ctrl: 快照trsf_ctrl模块<br><li>trsf_proxy: 快照trsf_proxy模块

	Module *string `json:"Module,omitempty" name:"Module"`
	// 当前模块的升级状态。取值范围：<br><li>WAIT_CHECK: 等待检查<br><li>CHECK_FINISH: 检查完成<br><li>CHECK_FAILED: 检查失败<br><li>RUNNING: 升级中<br><li>SUCCESS: 升级成功<br><li>FAILED: 升级失败<br><li>ROLLBACKED: 已回滚

	ModuleState *string `json:"ModuleState,omitempty" name:"ModuleState"`
	// 升级的目标集群版本

	NewClusterVersion *string `json:"NewClusterVersion,omitempty" name:"NewClusterVersion"`
	// 升级的目标模块版本

	NewModuleVersion *string `json:"NewModuleVersion,omitempty" name:"NewModuleVersion"`
	// 当前升级任务用到的升级包

	NewPkgName *string `json:"NewPkgName,omitempty" name:"NewPkgName"`
	// 集群的当前版本

	OldClusterVersion *string `json:"OldClusterVersion,omitempty" name:"OldClusterVersion"`
	// 模块的当前版本

	OldModuleVersion *string `json:"OldModuleVersion,omitempty" name:"OldModuleVersion"`
	// 集群所在的可用区

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 任务开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 任务的总步骤数

	TotalStep *uint64 `json:"TotalStep,omitempty" name:"TotalStep"`
	// 集群升级的事务ID

	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`
	// 组件的当前状态。取值范围：<br><li>NORMAL: 正常<br><li>FAULT: 异常

	State *string `json:"State,omitempty" name:"State"`
	// 是否为主节点。1表示为主节点，0表示备节点，空表示当前模块不是主备模式

	IsMaster *uint64 `json:"IsMaster,omitempty" name:"IsMaster"`
	// cell组件ID。仅当Module为cell时有意义，表示当前cell节点所在分组ID

	CellGroupId *uint64 `json:"CellGroupId,omitempty" name:"CellGroupId"`
	// 上次升级时间

	LastUpgradeTime *string `json:"LastUpgradeTime,omitempty" name:"LastUpgradeTime"`
	// 仅在Module为cell时有效，表示当前cell是属于第几副本，取值范围：0, 1, 2，分别表示cell三副本中的第一副本、第二副本、第三副本

	CellReplicateIndex *uint64 `json:"CellReplicateIndex,omitempty" name:"CellReplicateIndex"`
}

type SnapOverview struct {
	// 快照数量

	TotalSnapshotCount *uint64 `json:"TotalSnapshotCount,omitempty" name:"TotalSnapshotCount"`
	// 快照大小

	TotalSnapshotSize *uint64 `json:"TotalSnapshotSize,omitempty" name:"TotalSnapshotSize"`
	// 定期快照策略数量

	TotalAutoSnapshotPolicyCount *uint64 `json:"TotalAutoSnapshotPolicyCount,omitempty" name:"TotalAutoSnapshotPolicyCount"`
	// 绑定了定期快照策略的云硬盘数量

	TotalDiskBoundWithAutoSnapshotPolicyCount *uint64 `json:"TotalDiskBoundWithAutoSnapshotPolicyCount,omitempty" name:"TotalDiskBoundWithAutoSnapshotPolicyCount"`
}

type Policy struct {
	// 选定周一到周日中需要创建快照的日期，取值范围：[0, 6]。0表示周一触发，依此类推。

	DayOfWeek []*uint64 `json:"DayOfWeek,omitempty" name:"DayOfWeek"`
	// 指定定期快照策略的触发时间。单位为小时，取值范围：[0, 23]。00:00 ~ 23:00 共 24 个时间点可选，1表示 01:00，依此类推。

	Hour []*uint64 `json:"Hour,omitempty" name:"Hour"`
}

type DeleteSnapshotBoxResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSnapshotBoxResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSnapshotBoxResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerCapacityTopTrendsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomerCapacityTopTrendsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerCapacityTopTrendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerCapacityTopTrendsRequest struct {
	*tchttp.BaseRequest

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 过滤

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 开始日期

	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`
	// 截止日期

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeCustomerCapacityTopTrendsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerCapacityTopTrendsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSetCellSideIOStatRequest struct {
	*tchttp.BaseRequest

	// 开始时间，例："2021-03-17 00:00:00"

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 结束时间，例："2021-03-17 00:10:00"

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// 指定从秒级监控系统或分钟级监控系统查询数据，MINUTE表示从分钟级监控系统查询数据；SECOND表示从秒级监控系统查询数据；本接口固定传MINUTE

	Type *string `json:"Type,omitempty" name:"Type"`
	// 监控数据的周期，单位秒。取值范围：查询秒级监控时可传入10或60，表示查询周期为10s或60s的监控数据；查询分钟级监控数据，只能传入60，本接口固定传60

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 过滤条件。<br><li>set-uuid - Array of String - 是否必填：否 -（过滤条件）指定setUuid查询监控数据

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeSetCellSideIOStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSetCellSideIOStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 快照的数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 快照的详情列表。

		SnapshotSet []*Snapshot `json:"SnapshotSet,omitempty" name:"SnapshotSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallTransferPairsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID列表

		TaskIdList []*uint64 `json:"TaskIdList,omitempty" name:"TaskIdList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InstallTransferPairsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallTransferPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskDetail struct {
	// 云硬盘大小

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 云硬盘实际使用大小（GB）

	DiskSizeUsed *uint64 `json:"DiskSizeUsed,omitempty" name:"DiskSizeUsed"`
	// 快照已使用大小（GB） TODO：这里的快照大小指的是在存储仓库中占有的大小，当快照创建完成一段时间后会自动归为0

	SnapshotSizeUsed *uint64 `json:"SnapshotSizeUsed,omitempty" name:"SnapshotSizeUsed"`
	// 云硬盘状态

	DiskState *string `json:"DiskState,omitempty" name:"DiskState"`
	// 已挂载宿主机列表

	AttachedHostSet []*AttachedHost `json:"AttachedHostSet,omitempty" name:"AttachedHostSet"`
	// 用户appId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 云盘UUID

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
	// 云盘所在zone

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 云盘是否支持快照功能

	SnapshotAbility *int64 `json:"SnapshotAbility,omitempty" name:"SnapshotAbility"`
	// 云盘已挂载到子机，且子机与云盘都是包年包月。<br><li>true：子机设置了自动续费标识，但云盘未设置<br><li>false：云盘自动续费标识正常。

	AutoRenewFlagError *bool `json:"AutoRenewFlagError,omitempty" name:"AutoRenewFlagError"`
	// 在云盘已挂载到实例，且实例与云盘都是包年包月的条件下，此字段才有意义。<br><li>true:云盘到期时间早于实例。<br><li>false：云盘到期时间晚于实例。

	DeadlineError *bool `json:"DeadlineError,omitempty" name:"DeadlineError"`
	// 云盘是否已挂载

	Attached *bool `json:"Attached,omitempty" name:"Attached"`
	// 云盘创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 云盘到期时间

	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`
	// 云硬盘类型。取值范围：<br><li>SYSTEM_DISK：系统盘<br><li>DATA_DISK：数据盘。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 付费模式。取值范围：<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：后付费，即按量计费。

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 是否为弹性云盘，false表示非弹性云盘，true表示弹性云盘。

	// Portable *bool `json:"Portable,omitempty" name:"Portable"`
	Portable *interface{} `json:"Portable,omitempty" name:"Portable"`
	// 云硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 云盘名称

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 云盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：SSD表示SSD云硬盘。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云硬盘挂载的云主机ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 云硬盘挂载的云主机名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 云硬盘挂载的云主机UUID。

	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
	// 云盘是否在回滚中

	Rollbacking *int64 `json:"Rollbacking,omitempty" name:"Rollbacking"`
	// 云盘的回滚进度

	RollbackPercent *uint64 `json:"RollbackPercent,omitempty" name:"RollbackPercent"`
	// 云盘是否为加密盘。取值范围：<br><li>false:表示非加密盘<br><li>true:表示加密盘。

	Encrypt *string `json:"Encrypt,omitempty" name:"Encrypt"`
	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 云盘所在的存储池ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 是否为CFS云盘

	IsCfsDisk *uint64 `json:"IsCfsDisk,omitempty" name:"IsCfsDisk"`
	// 云盘挂载的母机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 云盘所在存储池的名称

	DepotName *string `json:"DepotName,omitempty" name:"DepotName"`
	// 云盘是否在迁移中

	Migrating *bool `json:"Migrating,omitempty" name:"Migrating"`
	// 云盘的迁移进度

	MigratePercent *uint64 `json:"MigratePercent,omitempty" name:"MigratePercent"`
	// 云盘流控状态

	FlowControlState *string `json:"FlowControlState,omitempty" name:"FlowControlState"`
	// 云盘IO性能

	FlowControlConfig *FlowControlConfig `json:"FlowControlConfig,omitempty" name:"FlowControlConfig"`
	// 版本。

	Version *uint64 `json:"Version,omitempty" name:"Version"`
	// 是否随子机销毁

	DeleteWithInstance *uint64 `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
	// 母机client的版本

	DriverVersion *string `json:"DriverVersion,omitempty" name:"DriverVersion"`
	// 是否为共享盘

	Shareable *uint64 `json:"Shareable,omitempty" name:"Shareable"`
	// 当前与盘到期相差的天数

	DifferDaysOfDeadline *int64 `json:"DifferDaysOfDeadline,omitempty" name:"DifferDaysOfDeadline"`
	// 所属资源池

	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
	// 云盘关联的定期快照ID

	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitempty" name:"AutoSnapshotPolicyIds"`
	// 云盘底层运维状态。

	DiskRealState *string `json:"DiskRealState,omitempty" name:"DiskRealState"`
	// 购买云盘订单名。

	DealName *string `json:"DealName,omitempty" name:"DealName"`
	// 绑定定期快照策略的数量。

	AutoSnapshotPolicyBoundCount *uint64 `json:"AutoSnapshotPolicyBoundCount,omitempty" name:"AutoSnapshotPolicyBoundCount"`
	// 云盘信息修改时间。

	ModifyTimeStamp *string `json:"ModifyTimeStamp,omitempty" name:"ModifyTimeStamp"`
	// 共享盘挂载的云服务器实例信息。

	ShareAttachedInstanceSet []*SimpleInstance `json:"ShareAttachedInstanceSet,omitempty" name:"ShareAttachedInstanceSet"`
	// 云盘挂载实例信息。

	Instance *DiskAttachedInstanceInfo `json:"Instance,omitempty" name:"Instance"`
}

type DiskMigrateTask struct {
	// 迁移任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 云硬盘ID

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 云硬盘UUID

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
	// 迁移任务发起人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 迁移任务创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 迁移任务启动时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 迁移任务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 源存储集群ID

	SourceDepotId *uint64 `json:"SourceDepotId,omitempty" name:"SourceDepotId"`
	// 目标存储集群ID

	DestinationDepotId *uint64 `json:"DestinationDepotId,omitempty" name:"DestinationDepotId"`
	// 源存储集群名称

	SourceDepotName *string `json:"SourceDepotName,omitempty" name:"SourceDepotName"`
	// 目标存储集群名称

	DestinationDepotName *string `json:"DestinationDepotName,omitempty" name:"DestinationDepotName"`
	// AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 迁移任务进度

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 迁移任务状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 迁移任务说明

	Message *string `json:"Message,omitempty" name:"Message"`
	// 迁移大小

	MigrateSize *uint64 `json:"MigrateSize,omitempty" name:"MigrateSize"`
	// 迁移任务类型。`SINGLE_WRITE`：单写迁移，`DOUBLE_WRITE`：双写迁移

	MigrateTaskType *string `json:"MigrateTaskType,omitempty" name:"MigrateTaskType"`
	// 云盘大小

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 目标存储池属性列表

	DestinationDepotAttributeName *string `json:"DestinationDepotAttributeName,omitempty" name:"DestinationDepotAttributeName"`
	// 云盘名称

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
}

type TransferPair struct {
	// transfer controller 组件ip。

	ControllerIP *string `json:"ControllerIP,omitempty" name:"ControllerIP"`
	// transfer proxy 组件ip。

	ProxyIP *string `json:"ProxyIP,omitempty" name:"ProxyIP"`
}

type DepotOverview struct {
	// 云硬盘类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 总容量（超卖后）

	TotalDepotSizeOverSold *uint64 `json:"TotalDepotSizeOverSold,omitempty" name:"TotalDepotSizeOverSold"`
	// 已创建的云盘容量

	TotalDepotSizeCreated *uint64 `json:"TotalDepotSizeCreated,omitempty" name:"TotalDepotSizeCreated"`
	// 总容量（实际）

	TotalDepotSize *uint64 `json:"TotalDepotSize,omitempty" name:"TotalDepotSize"`
	// 实际已写入容量

	TotalDepotSizeUsed *uint64 `json:"TotalDepotSizeUsed,omitempty" name:"TotalDepotSizeUsed"`
	// 快照占用容量

	TotalDepotSizeSnapshotUsed *uint64 `json:"TotalDepotSizeSnapshotUsed,omitempty" name:"TotalDepotSizeSnapshotUsed"`
	// 存储集群数量

	TotalDepotCount *uint64 `json:"TotalDepotCount,omitempty" name:"TotalDepotCount"`
	// 已创建的云硬盘数量

	TotalDepotDiskCountCreated *uint64 `json:"TotalDepotDiskCountCreated,omitempty" name:"TotalDepotDiskCountCreated"`
	// 总共可创建的云硬盘数量

	TotalDepotDiskCount *uint64 `json:"TotalDepotDiskCount,omitempty" name:"TotalDepotDiskCount"`
	// 存储仓库状态统计

	DepotStateSet []*DepotStateCount `json:"DepotStateSet,omitempty" name:"DepotStateSet"`
	// 存储仓库售卖状态统计

	DepotSoldStateSet []*DepotSoldStateSet `json:"DepotSoldStateSet,omitempty" name:"DepotSoldStateSet"`
	// 仓库可售卖容量。

	TotalDepotSizeSoldable *uint64 `json:"TotalDepotSizeSoldable,omitempty" name:"TotalDepotSizeSoldable"`
}

type DeleteDiskStoragePoolGroupRequest struct {
	*tchttp.BaseRequest

	// 云硬盘存储池属性ID

	DepotAttributeId *uint64 `json:"DepotAttributeId,omitempty" name:"DepotAttributeId"`
	// 要删除的云硬盘资源池(pool-xxxx)，该参数与DepotAttributeId参数只能有一个进行传递。

	PoolGroup *string `json:"PoolGroup,omitempty" name:"PoolGroup"`
}

func (r *DeleteDiskStoragePoolGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDiskStoragePoolGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResizeDiskRequest struct {
	*tchttp.BaseRequest

	// 云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 云硬盘扩容后的大小，单位为GB，必须大于当前云硬盘大小。取值范围： 普通云硬盘:10GB ~ 16000G；高性能云硬盘:50GB ~ 16000GB；SSD云硬盘:100GB ~ 16000GB，步长均为10GB。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 从运营端发起云盘扩容时，本参数必传，取值为TRUE。表示从运营端直接发起云盘扩容。

	InternalDirectResize *bool `json:"InternalDirectResize,omitempty" name:"InternalDirectResize"`
}

func (r *ResizeDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResizeDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DepotStateCount struct {
	// 存储仓库状态

	DepotState *string `json:"DepotState,omitempty" name:"DepotState"`
	// 数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type DescribeDiskDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云硬盘详情

		DiskDetailsSet []*DiskDetail `json:"DiskDetailsSet,omitempty" name:"DiskDetailsSet"`
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DepotFlowControlRecord struct {
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储仓库I

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 读带宽流控阈值

	ReadThroughLimit *int64 `json:"ReadThroughLimit,omitempty" name:"ReadThroughLimit"`
	// 读IOPS流控阈值

	ReadIopsLimit *int64 `json:"ReadIopsLimit,omitempty" name:"ReadIopsLimit"`
	// 写带宽流控阈值

	WriteThrouhLimit *int64 `json:"WriteThrouhLimit,omitempty" name:"WriteThrouhLimit"`
	// 写IOPS流控阈值

	WriteIopsLimit *int64 `json:"WriteIopsLimit,omitempty" name:"WriteIopsLimit"`
	// 流程操作时间

	OperateTime *string `json:"OperateTime,omitempty" name:"OperateTime"`
	// 触发这次仓库流控的原因 TODO

	ReadThroughRole *int64 `json:"ReadThroughRole,omitempty" name:"ReadThroughRole"`
	// 触发这次仓库流控的原因 TODO

	ReadIopsRole *int64 `json:"ReadIopsRole,omitempty" name:"ReadIopsRole"`
	// 触发这次仓库流控的原因 TODO

	WriteThroughRole *int64 `json:"WriteThroughRole,omitempty" name:"WriteThroughRole"`
	// 触发这次仓库流控的原因 TODO

	WriteIopsRole *int64 `json:"WriteIopsRole,omitempty" name:"WriteIopsRole"`
	// 存储池当前读带宽

	ReadThrough *int64 `json:"ReadThrough,omitempty" name:"ReadThrough"`
	// 存储池当前读IOPS

	ReadIops *int64 `json:"ReadIops,omitempty" name:"ReadIops"`
	// 存储池当前写带宽

	WriteThrough *int64 `json:"WriteThrough,omitempty" name:"WriteThrough"`
	// 存储池当前写IOPS

	WriteIops *int64 `json:"WriteIops,omitempty" name:"WriteIops"`
	// 流控的具体指标及类型

	Actions []*FlowControlAction `json:"Actions,omitempty" name:"Actions"`
}

type CacheDisk struct {
	// 缓存盘sn号

	CacheDiskSn *string `json:"CacheDiskSn,omitempty" name:"CacheDiskSn"`
	// 缓存盘所在的cell节点ip

	CellIp *string `json:"CellIp,omitempty" name:"CellIp"`
	// 缓存盘状态。<br><li>NORMAL: 正常<br><li>DEAD: 剔除<br><li>REPLACING: 替换中

	State *string `json:"State,omitempty" name:"State"`
	// 缓存盘关联的主存盘列表

	MainDisks []*MainDisk `json:"MainDisks,omitempty" name:"MainDisks"`
	// 缓存盘容量，单位MB

	CacheDiskSize *uint64 `json:"CacheDiskSize,omitempty" name:"CacheDiskSize"`
	// 是否支持热插拔

	SupportDiskHotSwap *bool `json:"SupportDiskHotSwap,omitempty" name:"SupportDiskHotSwap"`
	// 盘位号

	DiskSlot *string `json:"DiskSlot,omitempty" name:"DiskSlot"`
}

type SnapshotSImpleDiskInfo struct {
	// 盘关联cvm的类型，取值有CVM, RAW_CVM等

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 是否为共享盘

	Shareable *uint64 `json:"Shareable,omitempty" name:"Shareable"`
	// 云盘名称

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 云盘大小，单位GB

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 云盘id

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 是否为弹性云盘

	// Portable *bool `json:"Portable,omitempty" name:"Portable"`
	Portable *interface{} `json:"Portable,omitempty" name:"Portable"`
	// 云盘类型。CLOUD_SATA: 普通云硬盘；CLOUD_PREMIUM: 高性能云硬盘；CLOUD_SSD： SSD云硬盘

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 是否为加密盘；取值为encrypt表示加密盘

	Encrypt *string `json:"Encrypt,omitempty" name:"Encrypt"`
	// 云盘续费标记，仅对预付费云盘生效

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 云盘是否随挂载的cvm实例一起销毁

	DeleteWithInstance *uint64 `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
	// 云盘所在的资源池

	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
	// 云盘的状态

	DiskState *string `json:"DiskState,omitempty" name:"DiskState"`
	// 云盘的最近属性修改时间

	ModifyTimeStamp *string `json:"ModifyTimeStamp,omitempty" name:"ModifyTimeStamp"`
	// 云盘到期时间，仅对预付费云盘生效

	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`
	// 云盘的付费类型

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 云盘挂载的cvm实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 云盘用途类型。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 加密算法类型。SM4:国密；AES256:非国密

	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" name:"EncryptAlgorithm"`
}

type DepotDriver struct {
	// 存储池ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 母机IP

	DriverIp *string `json:"DriverIp,omitempty" name:"DriverIp"`
	// 母机上总挂载的云盘数

	DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`
	// driver的版本号

	DriverVersion *string `json:"DriverVersion,omitempty" name:"DriverVersion"`
	// driver的状态。NORMAL表示正常，FAULT表示异常

	DriverState *string `json:"DriverState,omitempty" name:"DriverState"`
}

type DescribeDepotDiskDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存储池资源详情

		DepotDiskDetails *DepotDiskDetails `json:"DepotDiskDetails,omitempty" name:"DepotDiskDetails"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDepotDiskDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepotDiskDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecycleRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 回收规则详情

		RecycleRuleSet []*RecycleRule `json:"RecycleRuleSet,omitempty" name:"RecycleRuleSet"`
		// TotalCount

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecycleRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecycleRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDepotCellDeviceConfigRequest struct {
	*tchttp.BaseRequest

	// 机型型号

	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
	// 存储池类型，取值范围：<br><li>sata: 普通存储池<br><li>fusion: 高性能存储池<br><li>ssd: SSD存储池

	DepotMedia *string `json:"DepotMedia,omitempty" name:"DepotMedia"`
	// 存储机型主存盘数量

	DiskNum *uint64 `json:"DiskNum,omitempty" name:"DiskNum"`
	// 存储机型主存盘容量，单位TB

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 主存盘接口类型，取值范围: <br><li>sata: sata接口<br><li>nvme: nvme接口

	PrimaryInfType *string `json:"PrimaryInfType,omitempty" name:"PrimaryInfType"`
	// 当前机型是否支持磁盘热插拔。true表示支持热插拔盘；false表示不支持热插拔盘

	SupportDiskHotSwap *bool `json:"SupportDiskHotSwap,omitempty" name:"SupportDiskHotSwap"`
	// 缓存盘接口类型，取值范围: <br><li>nvme: nvme接口<br><li>仅在有缓存盘时，传入该值

	CacheInfType *string `json:"CacheInfType,omitempty" name:"CacheInfType"`
	// 存储机型缓存盘数量，仅在有缓存盘时传入该值

	CacheDiskNum *uint64 `json:"CacheDiskNum,omitempty" name:"CacheDiskNum"`
	// 存储机型缓存盘大小，可能为小数，所以类型为字符串类型，单位TB。仅在有缓存盘时传入该值

	CacheDiskSize *string `json:"CacheDiskSize,omitempty" name:"CacheDiskSize"`
	// 存储机型的网络类型，取值范围：10Gb, 25Gb

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 新的存储机型型号。

	NewDeviceType *string `json:"NewDeviceType,omitempty" name:"NewDeviceType"`
	// 机型支持上架存储池的副本策略，取值范围，THREE_COPIES：三副本存储，EC(9+3): 9+3 EC存储

	CopyStrategies []*string `json:"CopyStrategies,omitempty" name:"CopyStrategies"`
}

func (r *UpdateDepotCellDeviceConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDepotCellDeviceConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 过滤键的名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DescribeSnapshotBoxsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的快照box数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 快照集群详情列表。

		SnapshotBoxSet []*SnapshotBox `json:"SnapshotBoxSet,omitempty" name:"SnapshotBoxSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotBoxsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotBoxsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecycleDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorInstance struct {
	// 监控实例维度描述

	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions"`
}

type TabletStatistic struct {
	// 总小表数量

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// dead小表数量

	DeadCount *uint64 `json:"DeadCount,omitempty" name:"DeadCount"`
	// normal小表数量

	NormalCount *uint64 `json:"NormalCount,omitempty" name:"NormalCount"`
	// err小表数量

	ErrCount *uint64 `json:"ErrCount,omitempty" name:"ErrCount"`
	// down小表数量

	DownCount *uint64 `json:"DownCount,omitempty" name:"DownCount"`
	// revocer小表数量

	RecoverCount *uint64 `json:"RecoverCount,omitempty" name:"RecoverCount"`
}

type InquiryPriceModifyDiskAttributesRequest struct {
	*tchttp.BaseRequest

	// 需迁移的云盘实例ID列表，只能对租户端盘进行类型升级询价。运营端的盘类型升级，不需要询价。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 云盘变更的目标类型，取值范围：<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：表示SSD云硬盘。<br>**当前不支持类型降级**

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
}

func (r *InquiryPriceModifyDiskAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceModifyDiskAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallDiskZKClusterRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// ZK集群类型（master、snapshot）

	ZKClusterType *string `json:"ZKClusterType,omitempty" name:"ZKClusterType"`
	// Master支撑机IP，至少需要两个作为主备

	DeviceIps []*string `json:"DeviceIps,omitempty" name:"DeviceIps"`
	// ZK集群名称

	ZKClusterName *string `json:"ZKClusterName,omitempty" name:"ZKClusterName"`
	// ZK集群最多可接入的存储池集群数量。当ZKClusterType为master时，该参数才生效，建议取值为6，最大取值为10，默认取值为1

	MaxDepotNum *uint64 `json:"MaxDepotNum,omitempty" name:"MaxDepotNum"`
	// 同一组件使用机器的容灾级别要求，传入此参数，上架前，会进行容灾级别的检验，不满足不会进行上架，取值范围：<br><li>any: 无要求<br><li>motherMachine: 跨母机，仅针对虚拟机生效<br><li>rack: 跨机架<br><li>switch: 跨交换机

	DisasterTolerance *string `json:"DisasterTolerance,omitempty" name:"DisasterTolerance"`
}

func (r *InstallDiskZKClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallDiskZKClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BalanceDiskStorageDepotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BalanceDiskStorageDepotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BalanceDiskStorageDepotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 快照组列表详情

		SnapshotGroupSet []*SnapshotGroup `json:"SnapshotGroupSet,omitempty" name:"SnapshotGroupSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryDiskTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务类型，必须与待重试的任务任务类型一致

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
	// 指定重试租户后台任务（CUSTOMER）或者运维任务（MAINTENANCE）；默认运维任务

	TaskCategory *string `json:"TaskCategory,omitempty" name:"TaskCategory"`
	// 用于设置重试（RETRY）任务或者禁止任务重试（BLOCK）

	TaskState *string `json:"TaskState,omitempty" name:"TaskState"`
	// 是否只做新盘检查, true表示只做新盘检查,默认为false。

	CheckReplaceDisk *bool `json:"CheckReplaceDisk,omitempty" name:"CheckReplaceDisk"`
}

func (r *RetryDiskTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryDiskTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDiskUsageRequest struct {
	*tchttp.BaseRequest

	// 过滤设置，支持按`app-id`，`uin`过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 统计日期

	AddDate *string `json:"AddDate,omitempty" name:"AddDate"`
}

func (r *DescribeUserDiskUsageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserDiskUsageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskDepotComponentDetailCell struct {
	// 上架时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 上架人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// cell机型

	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
	// cell ip

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 已使用大小

	JournalSize *uint64 `json:"JournalSize,omitempty" name:"JournalSize"`
	// 最近升级时间

	LastUpgradeTime *string `json:"LastUpgradeTime,omitempty" name:"LastUpgradeTime"`
	// 修改者

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 机架名

	RackName *string `json:"RackName,omitempty" name:"RackName"`
	// 状态

	State *string `json:"State,omitempty" name:"State"`
	// 物理盘大小

	StorageDiskSize *uint64 `json:"StorageDiskSize,omitempty" name:"StorageDiskSize"`
	// 上联交换机

	SwitchName *string `json:"SwitchName,omitempty" name:"SwitchName"`
	// 升级状态

	UpgradeState *string `json:"UpgradeState,omitempty" name:"UpgradeState"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 组件类型

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 机器的cpu型号

	DeviceCpuModel *string `json:"DeviceCpuModel,omitempty" name:"DeviceCpuModel"`
}

type ModifyReplaceDiskTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyReplaceDiskTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReplaceDiskTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSnapshotBoxConfigRequest struct {
	*tchttp.BaseRequest

	// 快照集群ID

	BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
	// 配置名称

	ConfigItem *string `json:"ConfigItem,omitempty" name:"ConfigItem"`
	// 具体配置值

	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`
	// 配置所在的组件类型，取值范围：manager, scheduler, tranfser, trsf_ctrl, trsf_proxy

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
}

func (r *UpdateSnapshotBoxConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSnapshotBoxConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAutoSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest

	// 要删除的定期快照策略ID列表。

	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitempty" name:"AutoSnapshotPolicyIds"`
}

func (r *DeleteAutoSnapshotPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAutoSnapshotPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskStorageDepotDetailsRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储集群ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDiskStorageDepotDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskStorageDepotDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskChargePrepaid struct {
	// 购买云盘的时长，默认单位为月，此时，取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>默认取值：NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 需要将云盘的到期时间与挂载的子机对齐时，可传入该参数。该参数表示子机当前的到期时间，此时Period如果传入，则表示子机需要续费的时长，云盘会自动按对齐到子机续费后的到期时间续费。

	CurInstanceDeadline *string `json:"CurInstanceDeadline,omitempty" name:"CurInstanceDeadline"`
}

type DiskConfig struct {
	// 配置是否可用。

	Available *bool `json:"Available,omitempty" name:"Available"`
	// 云盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬盘<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：SSD表示SSD云硬盘。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云盘类型。取值范围：<br><li>SYSTEM_DISK：表示系统盘<br><li>DATA_DISK：表示数据盘。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 付费模式。取值范围：<br><li>PREPAID：表示预付费，即包年包月<br><li>POSTPAID_BY_HOUR：表示后付费，即按量计费。

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 最大可配置云盘大小。

	MaxDiskSize *uint64 `json:"MaxDiskSize,omitempty" name:"MaxDiskSize"`
	// 最小可配置云盘大小。

	MinDiskSize *uint64 `json:"MinDiskSize,omitempty" name:"MinDiskSize"`
	// 所在[可用区](/document/api/213/9452#zone)。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例机型。

	DeviceClass *string `json:"DeviceClass,omitempty" name:"DeviceClass"`
	// 实例机型系列。详见[实例类型](https://cloud.tencent.com/document/product/213/11518)

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 步长

	StepSize *uint64 `json:"StepSize,omitempty" name:"StepSize"`
	// 额外的性能区间。

	ExtraPerformanceRange []*int64 `json:"ExtraPerformanceRange,omitempty" name:"ExtraPerformanceRange"`
}

type CancelMigrateDiskTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 取消成功的任务、云盘列表。

		CancelSuccessSet []*MigrateDiskTask `json:"CancelSuccessSet,omitempty" name:"CancelSuccessSet"`
		// 取消失败的任务、云盘列表。

		CancelFailedSet []*MigrateDiskTask `json:"CancelFailedSet,omitempty" name:"CancelFailedSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelMigrateDiskTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelMigrateDiskTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSnapshotsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDepotRepotEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的事件总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 存储池事件列表

		DepotEventSet []*DepotEvent `json:"DepotEventSet,omitempty" name:"DepotEventSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDepotRepotEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepotRepotEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDisksMonitorDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监控数据起始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 监控数据结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 监控周期

		Period *uint64 `json:"Period,omitempty" name:"Period"`
		// 监控指标

		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
		// 监控数据

		DataPoints []*DataPoint `json:"DataPoints,omitempty" name:"DataPoints"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDisksMonitorDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisksMonitorDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallTransferPairsRequest struct {
	*tchttp.BaseRequest

	// 快照集群ID

	BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
	// 待安装的transfer pair组件列表。

	TransferPairs []*TransferPair `json:"TransferPairs,omitempty" name:"TransferPairs"`
	// 可选，默认为0，仅仅安装，跳过向manager注册、激活的过程

	OnlyInstall *uint64 `json:"OnlyInstall,omitempty" name:"OnlyInstall"`
	// 可选，在OnlyInstall为0的情况下，选择是否要跳过激活步骤，为0表示需要激活，为1表示不激活。默认取值为0。

	SkipActive *uint64 `json:"SkipActive,omitempty" name:"SkipActive"`
	// 可选，是否使用DryRun模式，默认为0；

	DryRun *uint64 `json:"DryRun,omitempty" name:"DryRun"`
	// 同一组件使用机器的容灾级别要求，传入此参数，上架前，会进行容灾级别的检验，不满足不会进行上架，取值范围：<br><li>any: 无要求<br><li>motherMachine: 跨母机，仅针对虚拟机生效<br><li>rack: 跨机架<br><li>switch: 跨交换机

	DisasterTolerance *string `json:"DisasterTolerance,omitempty" name:"DisasterTolerance"`
}

func (r *InstallTransferPairsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallTransferPairsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiskStoragePoolGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDiskStoragePoolGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiskStoragePoolGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDepotDriverDetailsRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储池ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDepotDriverDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepotDriverDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnapshotResource struct {
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 表示当前统计的是哪天的快照容量

	Date *string `json:"Date,omitempty" name:"Date"`
	// 执行快照统计的时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 快照总数量

	TotalSnapCount *uint64 `json:"TotalSnapCount,omitempty" name:"TotalSnapCount"`
	// 快照总实际容量，单位MB

	TotalSnapSize *uint64 `json:"TotalSnapSize,omitempty" name:"TotalSnapSize"`
	// 快照云盘数量

	TotalSnapDiskCount *uint64 `json:"TotalSnapDiskCount,omitempty" name:"TotalSnapDiskCount"`
	// 快照关联盘的总容量，单位GB

	TotalSnapDiskSize *uint64 `json:"TotalSnapDiskSize,omitempty" name:"TotalSnapDiskSize"`
	// 定期快照总数量

	AutoSnapCount *uint64 `json:"AutoSnapCount,omitempty" name:"AutoSnapCount"`
	// 定期快照总容量，单位MB

	AutoSnapSize *uint64 `json:"AutoSnapSize,omitempty" name:"AutoSnapSize"`
	// 定期快照云盘数量

	AutoSnapDiskCount *uint64 `json:"AutoSnapDiskCount,omitempty" name:"AutoSnapDiskCount"`
	// 定期快照关联云盘总容量，单位GB

	AutoSnapDiskSize *uint64 `json:"AutoSnapDiskSize,omitempty" name:"AutoSnapDiskSize"`
	// 后台快照数量。

	InnerSnapCount *int64 `json:"InnerSnapCount,omitempty" name:"InnerSnapCount"`
	// 私有镜像大小。

	PrivateImageSize *int64 `json:"PrivateImageSize,omitempty" name:"PrivateImageSize"`
	// 快照用户数量。

	CostSnapAppIdCount *int64 `json:"CostSnapAppIdCount,omitempty" name:"CostSnapAppIdCount"`
	// 后台快照大小。

	InnerSnapSize *int64 `json:"InnerSnapSize,omitempty" name:"InnerSnapSize"`
	// 公有镜像大小。

	PublicImageSize *int64 `json:"PublicImageSize,omitempty" name:"PublicImageSize"`
	// 私有镜像大小。

	PrivateImageCount *int64 `json:"PrivateImageCount,omitempty" name:"PrivateImageCount"`
	// 快照用户数量。

	TotalSnapAppIdCount *int64 `json:"TotalSnapAppIdCount,omitempty" name:"TotalSnapAppIdCount"`
	// 后台快照云盘数量。

	InnerSnapDiskCount *int64 `json:"InnerSnapDiskCount,omitempty" name:"InnerSnapDiskCount"`
	// 私有镜像数量。

	PrivateCostImageCount *int64 `json:"PrivateCostImageCount,omitempty" name:"PrivateCostImageCount"`
	// 免费快照用户数量。

	FreeSnapAppIdCount *int64 `json:"FreeSnapAppIdCount,omitempty" name:"FreeSnapAppIdCount"`
	// cos存储容量。

	StorageCosSize *int64 `json:"StorageCosSize,omitempty" name:"StorageCosSize"`
	// 底层快照容量。

	OssSnapSize *int64 `json:"OssSnapSize,omitempty" name:"OssSnapSize"`
	// 已删除后台快照大小。

	DeleteBackupSize *int64 `json:"DeleteBackupSize,omitempty" name:"DeleteBackupSize"`
	// 已删除后台快照数量。

	DeleteBackupCount *int64 `json:"DeleteBackupCount,omitempty" name:"DeleteBackupCount"`
	// 后台快照云盘大小。

	InnerSnapDiskSize *int64 `json:"InnerSnapDiskSize,omitempty" name:"InnerSnapDiskSize"`
	// 免费快照大小。

	FreeSnapSize *int64 `json:"FreeSnapSize,omitempty" name:"FreeSnapSize"`
	// 私有镜像大小。

	PrivateCostImageSize *int64 `json:"PrivateCostImageSize,omitempty" name:"PrivateCostImageSize"`
	// 底层快照数量。

	OssSnapCount *int64 `json:"OssSnapCount,omitempty" name:"OssSnapCount"`
	// 公有镜像数量。

	PublicImageCount *int64 `json:"PublicImageCount,omitempty" name:"PublicImageCount"`
	// 全region快照用户数量。

	AllRegionAppIdCount *int64 `json:"AllRegionAppIdCount,omitempty" name:"AllRegionAppIdCount"`
}

type DescribeDepotDriverDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前存储池云盘挂载的总母机数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 存储池云盘概况

		DepotDriverSet []*DepotDriver `json:"DepotDriverSet,omitempty" name:"DepotDriverSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDepotDriverDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepotDriverDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSetBlockUsageDayStatisticsRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储池ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
}

func (r *DescribeSetBlockUsageDayStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSetBlockUsageDayStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallCapacityServerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 发起安装任务成功后，生产的异步des任务ID

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 如果一次安装多台机器，会生产多个异步任务

		TaskIdList []*uint64 `json:"TaskIdList,omitempty" name:"TaskIdList"`
		// 任务描述

		TaskDescription *string `json:"TaskDescription,omitempty" name:"TaskDescription"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InstallCapacityServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallCapacityServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectDiskStorageDepotCellRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储仓库ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 剔除的对象。<br><li>DEVICE：整机剔除；<br><li>DISK：单盘剔除，此时，必需传入参数Smcd，指定恢复的盘号。

	ObjectType *string `json:"ObjectType,omitempty" name:"ObjectType"`
	// 单盘剔除时传入，表示盘的smcd号。

	Smcd *uint64 `json:"Smcd,omitempty" name:"Smcd"`
	// 是否触发数据迁移。<br><li>NEED_TRANSFER：剔除盘后，需要触发数据迁移<br><li>SKIP_TRANSFER：剔除盘后无需数据迁移

	AutoTransferFlag *string `json:"AutoTransferFlag,omitempty" name:"AutoTransferFlag"`
	// 待剔除的存储池cell节点IP

	DeviceIp *string `json:"DeviceIp,omitempty" name:"DeviceIp"`
}

func (r *RejectDiskStorageDepotCellRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectDiskStorageDepotCellRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDiskStorageDepotDetails struct {
	// CELL分组信息列表

	CellGroups []*CellGroup `json:"CellGroups,omitempty" name:"CellGroups"`
	// LoadCtrl组件列表

	LoadCtrls *int64 `json:"LoadCtrls,omitempty" name:"LoadCtrls"`
	// 设备类型

	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
	// FlowControlConfig TODO

	FlowControlConfig *int64 `json:"FlowControlConfig,omitempty" name:"FlowControlConfig"`
	// Master组件列表

	Masters *int64 `json:"Masters,omitempty" name:"Masters"`
	// 小表对数量

	TotalTpCount *int64 `json:"TotalTpCount,omitempty" name:"TotalTpCount"`
}

type Tag struct {
	// 标签健。

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值。

	Value *string `json:"Value,omitempty" name:"Value"`
}

type Placement struct {
	// 实例所属的[可用区](/document/api/213/9452#zone)ID。该参数也可以通过调用  [DescribeZones](/document/product/213/15707) 的返回值中的Zone字段来获取。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例所属项目ID。该参数可以通过调用 [DescribeProject](/document/api/378/4400) 的返回值中的 projectId 字段来获取。不填为默认项目。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type SnapshotBox struct {
	// 快照box绑定的存储池ID列表。

	BindDepotIds []*uint64 `json:"BindDepotIds,omitempty" name:"BindDepotIds"`
	// 快照box id

	BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
	// 快照box名称

	BoxName *string `json:"BoxName,omitempty" name:"BoxName"`
	// 快照集群状态，`READ_WRITE`：可读写；`READ_ONLY`：只读；`CLOSED`：暂停使用；`INSTALLING`: 上架中；`UNINSTALLING`: 下架中

	BoxState *string `json:"BoxState,omitempty" name:"BoxState"`
	// 快照元数据存储池cos中，CosBucketName4Meta表示存储在cos的桶名称

	CosBucketName4Meta *string `json:"CosBucketName4Meta,omitempty" name:"CosBucketName4Meta"`
	// 快照数据存储池cos的桶名称

	CosBucketNameGaia *string `json:"CosBucketNameGaia,omitempty" name:"CosBucketNameGaia"`
	// 存储快照数据所使用cos账户的secretId

	CosSecretId *string `json:"CosSecretId,omitempty" name:"CosSecretId"`
	// 存储快照数据所使用cos账户的secretKey

	CosSecretKey *string `json:"CosSecretKey,omitempty" name:"CosSecretKey"`
	// 存储快照元数据所使用cos账户的secretId，一般与CosSecretId相同即可。

	CosSecretId4Meta *string `json:"CosSecretId4Meta,omitempty" name:"CosSecretId4Meta"`
	// 存储快照元数据所使用cos账户的secretKey，一般与CosSecretKey相同即可。

	CosSecretKey4Meta *string `json:"CosSecretKey4Meta,omitempty" name:"CosSecretKey4Meta"`
	// 快照集群创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 快照集群创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 快照使用cos的url

	CosHostName *string `json:"CosHostName,omitempty" name:"CosHostName"`
	// "1": ceph, "0": TFS

	CosType *string `json:"CosType,omitempty" name:"CosType"`
	// 如果CosHostName为ip，CosUriStyle应该设置为1,如果为域名，则设置为0

	CosUriStyle *uint64 `json:"CosUriStyle,omitempty" name:"CosUriStyle"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 最后修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 快照region

	Region *string `json:"Region,omitempty" name:"Region"`
	// 当前快照集群内正常快照的数量

	SnapshotNormalNumber *uint64 `json:"SnapshotNormalNumber,omitempty" name:"SnapshotNormalNumber"`
	// 当前快照集群内正常快照的容量，单位MB

	SnapshotNormalSize *float64 `json:"SnapshotNormalSize,omitempty" name:"SnapshotNormalSize"`
	// 快照集群状态。NORMAL：正常；FAULT：异常

	Status *string `json:"Status,omitempty" name:"Status"`
	// 快照集群版本，形如：v1.0.1

	Version *string `json:"Version,omitempty" name:"Version"`
	// 当前快照集群使用的ZK集群ID

	ZKClusterId *uint64 `json:"ZKClusterId,omitempty" name:"ZKClusterId"`
	// 当前快照集群支持升级到哪些版本

	SupportUpgradeVersions []*string `json:"SupportUpgradeVersions,omitempty" name:"SupportUpgradeVersions"`
	// 当前快照集群的升级状态详情

	UpgradeDetail *UpgradeDetail `json:"UpgradeDetail,omitempty" name:"UpgradeDetail"`
	// scheduler组件详情

	DiskSnapshotBoxSchedulers []*DiskDepotComponentDetail `json:"DiskSnapshotBoxSchedulers,omitempty" name:"DiskSnapshotBoxSchedulers"`
	// manager组件详情

	DiskSnapshotBoxManagers []*DiskDepotComponentDetail `json:"DiskSnapshotBoxManagers,omitempty" name:"DiskSnapshotBoxManagers"`
	// transfer组件详情

	DiskSnapshotBoxTransfers []*DiskDepotComponentDetail `json:"DiskSnapshotBoxTransfers,omitempty" name:"DiskSnapshotBoxTransfers"`
	// trsf_pair组件详情

	DiskSnapshotBoxTransferPairs []*DiskDepotComponentDetailTransferPair `json:"DiskSnapshotBoxTransferPairs,omitempty" name:"DiskSnapshotBoxTransferPairs"`
	// cap_server组件详情

	DiskSnapshotBoxCapacityServers []*DiskDepotComponentDetail `json:"DiskSnapshotBoxCapacityServers,omitempty" name:"DiskSnapshotBoxCapacityServers"`
	// common组件详情

	CommonModules []*CommonModule `json:"CommonModules,omitempty" name:"CommonModules"`
}

type InstallBoxDevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InstallBoxDevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallBoxDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceDiskStorageDepotMasterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceDiskStorageDepotMasterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceDiskStorageDepotMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskConfigQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云盘配置列表。

		DiskConfigSet []*DiskConfig `json:"DiskConfigSet,omitempty" name:"DiskConfigSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskConfigQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskConfigQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskMaintenanceTasksRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>job-id - Array of Integer - 是否必填：否 -（过滤条件）按任务ID过滤。<br><li>job-type - Array of String - 是否必填：否 -（过滤条件）按任务类型过滤。<br><li>status - Array of String - 是否必填：否 -（过滤条件）按任务状态过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDiskMaintenanceTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskMaintenanceTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotSharePermissionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 快照的分享信息的集合

		SharePermissionSet []*SharePermission `json:"SharePermissionSet,omitempty" name:"SharePermissionSet"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 最大可共享的数量

		MaxShareAccount *uint64 `json:"MaxShareAccount,omitempty" name:"MaxShareAccount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotSharePermissionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotSharePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBoxAttributesRequest struct {
	*tchttp.BaseRequest

	// 快照集群ID

	BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
	// 快照集群状态。快照集群状态，`READ_WRITE`：可读写；`READ_ONLY`：只读；`CLOSED`：暂停使用；`INSTALLING`: 上架中；`UNINSTALLING`: 下架中。该参数与SwitchMasterComponent至少填入一项。

	BoxState *string `json:"BoxState,omitempty" name:"BoxState"`
	// 指定需要切主的快照组件，取值范围：<br><li>manager: manager组件切主；<br><li>cap_server: cap_server组件切主。快照集群状态。快照集群状态，`READ_WRITE`：可读写；`READ_ONLY`：只读；`CLOSED`：暂停使用；`INSTALLING`: 上架中；`UNINSTALLING`: 下架中。该参数与BoxState至少填入一项。

	SwitchMasterComponent *string `json:"SwitchMasterComponent,omitempty" name:"SwitchMasterComponent"`
}

func (r *ModifyBoxAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBoxAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDisksMonitorAlarmRulesRequest struct {
	*tchttp.BaseRequest

	// 告警类型，取值范围：iohang, svctm_high, avg_svctm_high

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 云盘类型，取值范围：<br><li>CLOUD_BASIC: 普通云盘；<br><li>CLOUD_SSD,CLOUD_PREMIUM: 高性能云盘和SSD云盘

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 告警对象，取值范围：<br><li>disk: 云盘<br><li>set:存储池

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
	// 需要修改的告警规则详情

	DiskAlarmRuleDetails []*DiskAlarmRuleDetail `json:"DiskAlarmRuleDetails,omitempty" name:"DiskAlarmRuleDetails"`
}

func (r *ModifyDisksMonitorAlarmRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDisksMonitorAlarmRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecoverDiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 针对需恢复云盘创建的快照，可通过DescribeSnapshots接口查询快照的详情

		SnapInstanceId *string `json:"SnapInstanceId,omitempty" name:"SnapInstanceId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecoverDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecoverDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TransferDiskStorageDepotCellDiskPairRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储集群ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 源存储集群节点IP

	SourceCellIp *string `json:"SourceCellIp,omitempty" name:"SourceCellIp"`
	// 目的存储集群节点IP

	DestinationCellIp *string `json:"DestinationCellIp,omitempty" name:"DestinationCellIp"`
	// 源Smcd

	SourceSmcd *uint64 `json:"SourceSmcd,omitempty" name:"SourceSmcd"`
	// 目的Smcd

	DestinationSmcd *uint64 `json:"DestinationSmcd,omitempty" name:"DestinationSmcd"`
	// 迁移对象。<br><li>DEVICE：整组cell数据迁移；<br><li>DISK：单组磁盘数据迁移，此时，必需传入参数SourceSmcd，指定要迁移盘对应的smcd值

	ObjectType *string `json:"ObjectType,omitempty" name:"ObjectType"`
	// 是否自动选择目的cell或磁盘。0：手动迁移；1：自动迁移。当取消为1时，不能传入DestinationCellIp或DestinationSmcd。

	AutoSelectDestination *uint64 `json:"AutoSelectDestination,omitempty" name:"AutoSelectDestination"`
	// 指定迁移小表的比例，比如50，表示仅迁移源端50%的used小表数；仅在ObjectType=DISK且AutoSelectDestination=0，即手动迁移单盘时才支持传入此参数

	TransferTpRatio *uint64 `json:"TransferTpRatio,omitempty" name:"TransferTpRatio"`
	// 是否执行实际迁移操作。1表示不执行实际迁移操作，只做检查；0表示会做实际迁移操作；默认取值为0

	DryRun *uint64 `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *TransferDiskStorageDepotCellDiskPairRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TransferDiskStorageDepotCellDiskPairRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallTransferPairsRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 快照集群ID

	BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
	// trasf controller节点ip

	TrasfControllerIP *string `json:"TrasfControllerIP,omitempty" name:"TrasfControllerIP"`
	// trasf proxy节点ip

	TrasfProxyIP *string `json:"TrasfProxyIP,omitempty" name:"TrasfProxyIP"`
	// trasfPair的ID，可通过接口DescribeSnapshotBoxs查询。 如果传入-1，则卸载过程中不会执行unactive和unregister操作。

	PairId *int64 `json:"PairId,omitempty" name:"PairId"`
	// 可选，默认为0，是否强制卸载

	Force *uint64 `json:"Force,omitempty" name:"Force"`
	// 可选，默认为0；DryRun模式

	DryRun *uint64 `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *UninstallTransferPairsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallTransferPairsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskOverview struct {
	// 云硬盘类型，如果是多种类型的云硬盘汇总数据则为空

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云硬盘数量

	TotalDiskCount *uint64 `json:"TotalDiskCount,omitempty" name:"TotalDiskCount"`
	// 云硬盘大小

	TotalDiskSize *uint64 `json:"TotalDiskSize,omitempty" name:"TotalDiskSize"`
}

type Snapshot struct {
	// 快照ID。

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 快照所在的位置。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 创建此快照的云硬盘类型。取值范围：<br><li>SYSTEM_DISK：系统盘<br><li>DATA_DISK：数据盘。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 创建此快照的云硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 创建此快照的云硬盘大小。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 快照的状态。取值范围：<br><li>NORMAL：正常<br><li>CREATING：创建中<br><li>ROLLBACKING：回滚中<br><li>COPYING_FROM_REMOTE：跨地域复制快照拷贝中。

	SnapshotState *string `json:"SnapshotState,omitempty" name:"SnapshotState"`
	// 快照名称，用户自定义的快照别名。调用[ModifySnapshotAttribute](/document/product/362/15650)可修改此字段。

	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
	// 快照创建进度百分比，快照创建成功后此字段恒为100。

	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`
	// 快照的创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 快照到期时间。如果快照为永久保留，此字段为空。

	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`
	// 是否为加密盘创建的快照。取值范围：<br><li>true：该快照为加密盘创建的<br><li>false:非加密盘创建的快照。

	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`
	// 是否为永久快照。取值范围：<br><li>true：永久快照<br><li>false：非永久快照。

	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`
	// 快照正在跨地域复制的目的地域，默认取值为[]。

	CopyingToRegions []*string `json:"CopyingToRegions,omitempty" name:"CopyingToRegions"`
	// 是否为跨地域复制的快照。取值范围：<br><li>true：表示为跨地域复制的快照。<br><li>false:本地域的快照。

	CopyFromRemote *bool `json:"CopyFromRemote,omitempty" name:"CopyFromRemote"`
	// 快照当前被共享数

	ShareReference *uint64 `json:"ShareReference,omitempty" name:"ShareReference"`
	// 快照uuid

	SnapUuid *string `json:"SnapUuid,omitempty" name:"SnapUuid"`
	// 快照大小

	SnapshotSize *uint64 `json:"SnapshotSize,omitempty" name:"SnapshotSize"`
	// 快照删除时间

	DeletedTime *string `json:"DeletedTime,omitempty" name:"DeletedTime"`
	// 快照记录最新修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 快照所属appId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 快照关联的镜像数量

	ImageCount *uint64 `json:"ImageCount,omitempty" name:"ImageCount"`
	// 标志快照是否为后台快照，1为后台快照，0为用户快照。

	BackendSnap *uint64 `json:"BackendSnap,omitempty" name:"BackendSnap"`
	// 快照所属云盘的uuid。

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
	// 快照的额外信息。

	// ExtraInfo *string `json:"ExtraInfo,omitempty" name:"ExtraInfo"`
	ExtraInfo *interface{} `json:"ExtraInfo,omitempty" name:"ExtraInfo"`
	// 快照所属云盘的挂载实例信息。

	Instance *DiskAttachedInstanceInfo `json:"Instance,omitempty" name:"Instance"`
	// 标志是否为cfs快照。

	IsCfsSnapshot *uint64 `json:"IsCfsSnapshot,omitempty" name:"IsCfsSnapshot"`
	// 父快照的uuid

	ParentSnapshotUuid *string `json:"ParentSnapshotUuid,omitempty" name:"ParentSnapshotUuid"`
	// 快照实际大小，单位MB

	SnapshotRealSize *uint64 `json:"SnapshotRealSize,omitempty" name:"SnapshotRealSize"`
	// 快照在底层的状态。取值范围：<br><li>NORMAL：正常<br><li>DELETED：已删除

	SnapshotRealState *string `json:"SnapshotRealState,omitempty" name:"SnapshotRealState"`
	// 快照类型。FULL: 全量快照；INC: 增量快照

	SnapshotType *string `json:"SnapshotType,omitempty" name:"SnapshotType"`
	// 快照关联的镜像列表

	Images []*Image `json:"Images,omitempty" name:"Images"`
	// 快照关联盘的简要信息

	Disk *SnapshotSImpleDiskInfo `json:"Disk,omitempty" name:"Disk"`
	// 加密算法类型。SM4:国密；AES256:非国密

	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" name:"EncryptAlgorithm"`
}

type UseInfo struct {
	// 标识存储机型是否正在使用，0表示没有使用，1表示正在使用。

	InUse *int64 `json:"InUse,omitempty" name:"InUse"`
}

type DescribeDiskStoragePoolGroupBoundResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定云硬盘存储资源池的用户数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 云硬盘存储池组绑定信息列表

		DiskStoragePoolGroupBoundSet []*DiskStoragePoolGroupBound `json:"DiskStoragePoolGroupBoundSet,omitempty" name:"DiskStoragePoolGroupBoundSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskStoragePoolGroupBoundResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskStoragePoolGroupBoundResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDepotMasterConfigRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储池ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 支持设置的master配置项，取值范围：<br><li>MAX_TRSF_TASK_NUM: 小表迁移并发数<br><li>TRSF_SWITCH：数据自动重建路径

	ConfigItem *string `json:"ConfigItem,omitempty" name:"ConfigItem"`
	// 配置项需要设置的值

	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`
}

func (r *ModifyDepotMasterConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDepotMasterConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindAutoSnapshotPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceOverview struct {
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 存储集群资源使用概览

	DepotOverview *DepotOverview `json:"DepotOverview,omitempty" name:"DepotOverview"`
	// 分类型的存储集群资源使用概览

	DepotOverviewByDiskTypeSet []*DepotOverview `json:"DepotOverviewByDiskTypeSet,omitempty" name:"DepotOverviewByDiskTypeSet"`
	// 云硬盘资源使用概览

	DiskOverview *DiskOverview `json:"DiskOverview,omitempty" name:"DiskOverview"`
	// 分类型的云硬盘资源使用概览

	DiskOverviewByDiskTypeSet []*DiskOverview `json:"DiskOverviewByDiskTypeSet,omitempty" name:"DiskOverviewByDiskTypeSet"`
	// 快照资源使用概览

	SnapOverview *SnapOverview `json:"SnapOverview,omitempty" name:"SnapOverview"`
	// zk集群概览

	ZKClusterOverview []*ZKClusterOverview `json:"ZKClusterOverview,omitempty" name:"ZKClusterOverview"`
	// 各状态组件概览

	MachineStateOverview []*MachineStateOverview `json:"MachineStateOverview,omitempty" name:"MachineStateOverview"`
	// 存储池事件概览

	TopDepotReportEvent []*TopDepotReportEvent `json:"TopDepotReportEvent,omitempty" name:"TopDepotReportEvent"`
	// 快照集群概览

	SnapshotBoxOverview []*SnapshotBoxOverview `json:"SnapshotBoxOverview,omitempty" name:"SnapshotBoxOverview"`
	// 可用区级别的存储仓库使用概览；

	DepotOverviewInAZ []*DepotOverviewInAZ `json:"DepotOverviewInAZ,omitempty" name:"DepotOverviewInAZ"`
	// 可用区级别的不同云盘类型的使用概览；

	DepotOverviewByDiskTypeSetInAZ []*DepotOverviewByDiskTypeSetInAZ `json:"DepotOverviewByDiskTypeSetInAZ,omitempty" name:"DepotOverviewByDiskTypeSetInAZ"`
}

type DescribeDiskTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的任务总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 任务详情列表

		DiskTaskSet []*DiskTask `json:"DiskTaskSet,omitempty" name:"DiskTaskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CellGroup struct {
	// CELL设备信息列表

	Cells []*DiskDepotComponentDetailCell `json:"Cells,omitempty" name:"Cells"`
	// CELL分组UUID

	CellGroupUuid *string `json:"CellGroupUuid,omitempty" name:"CellGroupUuid"`
	// DBTRSF设备信息列表

	Dbtrsfs []*DiskDepotComponentDetail `json:"Dbtrsfs,omitempty" name:"Dbtrsfs"`
	// CELL上物理盘及上表详情

	DiskPairs []*DiskPair `json:"DiskPairs,omitempty" name:"DiskPairs"`
	// CELL小表状态详情

	CellTpStateInfo *CellTpStateInfo `json:"CellTpStateInfo,omitempty" name:"CellTpStateInfo"`
	// 当前cell分组上的缓存盘信息。只有高性能存储池才有缓存盘

	CacheDisks []*CacheDisk `json:"CacheDisks,omitempty" name:"CacheDisks"`
	// CELL端口号。

	Port *uint64 `json:"Port,omitempty" name:"Port"`
}

type AttachedHost struct {
	// 宿主机IP

	HostIpAddress *string `json:"HostIpAddress,omitempty" name:"HostIpAddress"`
	// 写带宽 TODO：限制？还是啥？

	ReadThroughLimit *int64 `json:"ReadThroughLimit,omitempty" name:"ReadThroughLimit"`
	// 读带宽 TODO：限制？还是啥？

	WriteThroughLimit *int64 `json:"WriteThroughLimit,omitempty" name:"WriteThroughLimit"`
	// 读次数  TODO：限制？还是啥？

	ReadIopsLimit *int64 `json:"ReadIopsLimit,omitempty" name:"ReadIopsLimit"`
	// 写次数 TODO：限制？还是啥？

	WriteIopsLimit *int64 `json:"WriteIopsLimit,omitempty" name:"WriteIopsLimit"`
}

type DeleteSnapshotBoxRequest struct {
	*tchttp.BaseRequest

	// 快照集群ID

	BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
	// 0: 正常删除，要在此存储里所有模块都卸载后才能调用；1: 强制删除，只是把快照集群信息从从存储系统中删除，不卸载快照集群里的模块。默认取值为0

	ForceDelete *int64 `json:"ForceDelete,omitempty" name:"ForceDelete"`
}

func (r *DeleteSnapshotBoxRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSnapshotBoxRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerAccountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomerAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorEsIndexStorageTimeRequest struct {
	*tchttp.BaseRequest

	// 指定查询具体索引的保留时间设置；取值范围：<br><li>qemu_disk_iostat_ten_sec: 云盘10秒级监控数据<br><li>qemu_disk_iostat_min: 云盘1分钟级监控数据<br><li>qemu_disk_iostat_five_min: 云盘5分钟级监控数据<br><li>qemu_disk_info_ten_sec: 存储池10秒级监控数据<br><li>qemu_disk_info_min: 存储池分钟级监控数据<br><li>qemu_disk_info_five_min: 存储斌5分钟级监控数据

	IndexNames []*string `json:"IndexNames,omitempty" name:"IndexNames"`
}

func (r *DescribeMonitorEsIndexStorageTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorEsIndexStorageTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskAssociatedAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest

	// 要查询的云硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

func (r *DescribeDiskAssociatedAutoSnapshotPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskAssociatedAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelClusterUpgradeTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelClusterUpgradeTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelClusterUpgradeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新创建的定期快照策略ID。

		AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
		// 首次开始备份的时间。

		NextTriggerTime *string `json:"NextTriggerTime,omitempty" name:"NextTriggerTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAutoSnapshotPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DepotMasterConfig struct {
	// 配置项名称

	ConfigItem *string `json:"ConfigItem,omitempty" name:"ConfigItem"`
	// 配置项值。为了统一处理，如果配置项值是int的，返回结果里也是string型的数字。

	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`
}

type BindUserToDiskStoragePoolGroupRequest struct {
	*tchttp.BaseRequest

	// 可用区列表

	ZoneIds []*uint64 `json:"ZoneIds,omitempty" name:"ZoneIds"`
	// 云硬盘类型，不填写此参数时会绑定全部可用的云硬盘类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 用户账号列表

	AppIds []*uint64 `json:"AppIds,omitempty" name:"AppIds"`
	// 要绑定的云硬盘资源池属性ID

	DepotAttributeId *uint64 `json:"DepotAttributeId,omitempty" name:"DepotAttributeId"`
	// 要绑定的云硬盘资源池(pool-xxxx)，该参数与DepotAttributeId参数只能有一个进行传递。

	PoolGroup *string `json:"PoolGroup,omitempty" name:"PoolGroup"`
}

func (r *BindUserToDiskStoragePoolGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindUserToDiskStoragePoolGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySnapshotAttributeRequest struct {
	*tchttp.BaseRequest

	// 快照ID。

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 修改快照的状态。当前仅用于将后台快照转为用户快照，传入值normal。

	LifeState *string `json:"LifeState,omitempty" name:"LifeState"`
}

func (r *ModifySnapshotAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySnapshotAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindDiskStorageDepotFromDiskStoragePoolGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindDiskStorageDepotFromDiskStoragePoolGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindDiskStorageDepotFromDiskStoragePoolGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallCapacityServerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 卸载容量server生成的底层运维任务

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UninstallCapacityServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallCapacityServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerCapacityTopChangeStartDateRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCustomerCapacityTopChangeStartDateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerCapacityTopChangeStartDateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDisksMonitorAlarmRulesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>metric - Array of String - 是否必填：否 -（过滤条件）按告警类型过滤。 (iohang, svctm_high, avg_svctm_high)<br><li>disk-type - Array of String - 是否必填：否 -（过滤条件）按照云硬盘类型过滤。 (CLOUD_BACIS: 表示普通云盘 | CLOUD_SSD,CLOUD_PREMIUM：表示高性能云盘和SSD云盘)

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDisksMonitorAlarmRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisksMonitorAlarmRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest

	// 要查询的定期快照策略ID列表。参数不支持同时指定`AutoSnapshotPolicyIds`和`Filters`。

	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitempty" name:"AutoSnapshotPolicyIds"`
	// 过滤条件。参数不支持同时指定`AutoSnapshotPolicyIds`和`Filters`。<br><li>auto-snapshot-policy-id - Array of String - 是否必填：否 -（过滤条件）按定期快照策略ID进行过滤。定期快照策略ID形如：`asp-11112222`。<br><li>auto-snapshot-policy-state - Array of String - 是否必填：否 -（过滤条件）按定期快照策略的状态进行过滤。定期快照策略ID形如：`asp-11112222`。(NORMAL：正常 | ISOLATED：已隔离。)<br><li>auto-snapshot-policy-name - Array of String - 是否必填：否 -（过滤条件）按定期快照策略名称进行过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/362/13158)中的相关小节。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/362/13158)中的相关小节。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 输出定期快照列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 定期快照列表排序的依据字段。取值范围：<br><li>CREATETIME：依据定期快照的创建时间排序<br>默认按创建时间排序。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeAutoSnapshotPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoSnapshotPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotBoxPerformanceDataRequest struct {
	*tchttp.BaseRequest

	// 快照集群ID

	BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
	// 组件类型，取值范围：scheduler, transfer, trsf_ctrl

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
	// 指标列表，支持一次查询多个指标

	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 指定要查询的性能监控数据的机器ip。当前scheduler上的指标是需要聚合多台scheduler上的值，此时不能传入DeviceIp。其他组件的指标需要传入DeviceIp

	DeviceIp *string `json:"DeviceIp,omitempty" name:"DeviceIp"`
	// 监控数据的周期，单位秒。取值范围：60, 300。

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
}

func (r *DescribeSnapshotBoxPerformanceDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotBoxPerformanceDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiskAttributesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的云硬盘ID。如果传入多个云盘ID，仅支持所有云盘修改为同一属性。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 新的云硬盘项目ID，只支持修改弹性云盘的项目ID。通过[DescribeProject](/document/api/378/4400)接口查询可用项目及其ID。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 新的云硬盘名称。

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 设置流控信息

	FlowControlConfig *FlowControlConfig `json:"FlowControlConfig,omitempty" name:"FlowControlConfig"`
	// 云盘UUID列表

	DiskUuids []*string `json:"DiskUuids,omitempty" name:"DiskUuids"`
	// 变更云盘类型时，可传入该参数，表示变更的目标类型，取值范围：<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：表示SSD云硬盘。；<br>变更云盘类型时不支持同时变更其他属性；支持对运营端盘和租户端盘变更类型，如果变更的盘是租户端盘，会经过计量或计费系统（取决于当前环境是用的计量还是计费）

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 仅在变更云盘类型时需要用到，表示是否需要经过计量、计费系统（取决于当前环境是用的计量还是计费）；<br><li>如果变更的盘为运营端盘，需取值为true，表示不需要经过计量、计费系统；<br><li>如果变更的盘为租户端盘，需取值为false，表示需要经过计量、计费系统；默认取值为false

	InternalDirectUpgrade *bool `json:"InternalDirectUpgrade,omitempty" name:"InternalDirectUpgrade"`
}

func (r *ModifyDiskAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiskAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResizeDiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResizeDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResizeDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryDiskTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryDiskTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryDiskTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSnapshotBoxUpgradeTasksRequest struct {
	*tchttp.BaseRequest

	// 快照集群ID

	BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
	// 存储池当前版本，可通过接口DescribeSnapshotBoxs查询，见Version字段

	CurVersion *string `json:"CurVersion,omitempty" name:"CurVersion"`
	// 升级的目标版本。当前存储池支持升级的目标版本列表可通过接口DescribeSnapshotBoxs查询，见SupportUpgradeVersions字段

	NewVersion *string `json:"NewVersion,omitempty" name:"NewVersion"`
}

func (r *CreateSnapshotBoxUpgradeTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSnapshotBoxUpgradeTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiskStorageDepotAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDiskStorageDepotAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiskStorageDepotAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DepotEvent struct {
	// 事件ID。

	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
	// CBS底层zone id。

	CbsZoneId *uint64 `json:"CbsZoneId,omitempty" name:"CbsZoneId"`
	// cell节点ip。

	CellIp *string `json:"CellIp,omitempty" name:"CellIp"`
	// cell pair id。

	CellPairId *uint64 `json:"CellPairId,omitempty" name:"CellPairId"`
	// 事件版本。

	CellStateChgVer *string `json:"CellStateChgVer,omitempty" name:"CellStateChgVer"`
	// 存储池ID。

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 存储池名称。

	DepotName *string `json:"DepotName,omitempty" name:"DepotName"`
	// 当前上报的事件是否需要换盘处理。取值范围：<br><li>NO_NEED_REPLACE：无需换盘<br><li>REPLACE_MAIN_DISK：主存换盘<br><li>REPLACE_CACHE_DISK:缓存换盘<br><li>NEED_MANUAL_CHECK:人工确认

	ReplaceDiskType *string `json:"ReplaceDiskType,omitempty" name:"ReplaceDiskType"`
	// 事件简述。

	EventInfo *string `json:"EventInfo,omitempty" name:"EventInfo"`
	// 事件类型整数枚举值。

	EventTypeEnum *uint64 `json:"EventTypeEnum,omitempty" name:"EventTypeEnum"`
	// 事件类型字符串枚举值。

	EventTypeStr *string `json:"EventTypeStr,omitempty" name:"EventTypeStr"`
	// 事件uuid

	EventUuid *string `json:"EventUuid,omitempty" name:"EventUuid"`
	// oss是否捕获该事件。

	HasOssFetched *uint64 `json:"HasOssFetched,omitempty" name:"HasOssFetched"`
	// 事件发生时间。

	OccurTime *string `json:"OccurTime,omitempty" name:"OccurTime"`
	// zone id。

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储池所在地域。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 如果当前事件有换盘任务在处理，ReplaceTaskId表示对应的换盘任务ID。

	ReplaceTaskId *uint64 `json:"ReplaceTaskId,omitempty" name:"ReplaceTaskId"`
	// smcd

	Smcd *uint64 `json:"Smcd,omitempty" name:"Smcd"`
	// 盘的sn号。

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 事件维度。cell, master。

	EventObject *string `json:"EventObject,omitempty" name:"EventObject"`
	// 存储池属性

	DepotAttributeName *string `json:"DepotAttributeName,omitempty" name:"DepotAttributeName"`
	// 硬盘组号

	DiskPairId *uint64 `json:"DiskPairId,omitempty" name:"DiskPairId"`
	// 是否可用于换盘事件

	UseForReplace *uint64 `json:"UseForReplace,omitempty" name:"UseForReplace"`
}

type DiskFlowControlRecord struct {
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储池ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 云盘UUID

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
	// 云盘读带宽上限，单位KB

	ReadThroughLimit *uint64 `json:"ReadThroughLimit,omitempty" name:"ReadThroughLimit"`
	// 云盘读IOPS上限

	ReadIopsLimit *uint64 `json:"ReadIopsLimit,omitempty" name:"ReadIopsLimit"`
	// 云盘写带宽上限，单位KB

	WriteThrouthLimit *uint64 `json:"WriteThrouthLimit,omitempty" name:"WriteThrouthLimit"`
	// 云盘写IOPS上限

	WriteIopsLimit *uint64 `json:"WriteIopsLimit,omitempty" name:"WriteIopsLimit"`
	// 云盘当前读带宽

	ReadThrough *uint64 `json:"ReadThrough,omitempty" name:"ReadThrough"`
	// 云盘当前读IOPS

	ReadIops *uint64 `json:"ReadIops,omitempty" name:"ReadIops"`
	// 云盘当前写带宽

	WriteThrough *uint64 `json:"WriteThrough,omitempty" name:"WriteThrough"`
	// 云盘当前写IOPS

	WriteIops *uint64 `json:"WriteIops,omitempty" name:"WriteIops"`
	// 操作开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 操作类型。<br><li>RECOVER: 恢复<br><li>SUPRESS:打压<br><li>SET_BLACK_LIST:黑名单流控

	Action *string `json:"Action,omitempty" name:"Action"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DiskStorageDeviceConfig struct {
	// 存储设备类型

	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
	// 存储设备物理硬盘数量

	PhysicalDiskCount *uint64 `json:"PhysicalDiskCount,omitempty" name:"PhysicalDiskCount"`
	// 存储设备物理硬盘大小（TB）

	PhysicalDiskSize *uint64 `json:"PhysicalDiskSize,omitempty" name:"PhysicalDiskSize"`
	// 存储设备缓存盘类型

	StorageCacheInterfaceType *string `json:"StorageCacheInterfaceType,omitempty" name:"StorageCacheInterfaceType"`
	// 存储设备物理盘类型

	StorageMediaType *string `json:"StorageMediaType,omitempty" name:"StorageMediaType"`
	// 存储设备物理盘接口类型

	StoragePrimaryInterfaceType *string `json:"StoragePrimaryInterfaceType,omitempty" name:"StoragePrimaryInterfaceType"`
}

type DescribeDepotTransferTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的迁移任务总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 迁移任务详情。

		DepotTransferTaskSet []*DepotTransferTask `json:"DepotTransferTaskSet,omitempty" name:"DepotTransferTaskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDepotTransferTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepotTransferTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskStorageDepotDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云硬盘存储仓库详情

		DiskStorageDepotDetails *DiskStorageDepotDetails `json:"DiskStorageDepotDetails,omitempty" name:"DiskStorageDepotDetails"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskStorageDepotDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskStorageDepotDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAutoSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAutoSnapshotPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAutoSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskStoragePoolGroup struct {
	// 资源池名称

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 资源池类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 资源池名称

	Pool *string `json:"Pool,omitempty" name:"Pool"`
	// 是否对外可见

	Visible *int64 `json:"Visible,omitempty" name:"Visible"`
}

type SubTaskStatistic struct {
	// 成功数量

	SucceedCount *uint64 `json:"SucceedCount,omitempty" name:"SucceedCount"`
	// 失败数量

	FailedCount *uint64 `json:"FailedCount,omitempty" name:"FailedCount"`
}

type DescribeDepotConfigsRequest struct {
	*tchttp.BaseRequest

	// 存储池ID，存储池的唯一标识。

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 配置项所在的模块名，取值范围：master, cell, dbtrsf.

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
	// 需要查询的配置项列表，每一个配置项由section和item组成，通过冒号“:”分割，如:["log:log_level", "trsf_conf: trsf_main_switch"]

	ConfigItems []*string `json:"ConfigItems,omitempty" name:"ConfigItems"`
}

func (r *DescribeDepotConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepotConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskIOStatRequest struct {
	*tchttp.BaseRequest

	// 开始时间，例："2021-03-17 00:00:00"

	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`
	// 结束时间，例："2021-03-17 00:10:00"

	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`
	// 查询分钟级或秒级监控数据，MINUTE表示查询分钟级监控数据；SECOND表示查询秒级监控数据

	Type *string `json:"Type,omitempty" name:"Type"`
	// 监控数据的周期，单位秒。取值范围：查询秒级监控时可传入10或60，表示查询周期为10s或60s的监控数据；查询分钟级监控数据，只能传入60

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 过滤条件。<br><li>disk-uuid - Array of String - 是否必填：否 -（过滤条件）指定diskUuid查询监控数据

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDiskIOStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskIOStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallBoxDeviceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UninstallBoxDeviceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallBoxDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DepotTransferTask struct {
	// 任务ID。

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// cbs底层zone id。

	CbsZoneId *uint64 `json:"CbsZoneId,omitempty" name:"CbsZoneId"`
	// 存储池ID。

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 存储池名称。

	DepotName *string `json:"DepotName,omitempty" name:"DepotName"`
	// 目的cell节点IP。

	DstCellIp *string `json:"DstCellIp,omitempty" name:"DstCellIp"`
	// 目的盘smcd。

	DstSmcd *uint64 `json:"DstSmcd,omitempty" name:"DstSmcd"`
	// 目的小表ID。

	DstTpId *string `json:"DstTpId,omitempty" name:"DstTpId"`
	// 目的小表版本号。

	DstTpVer *string `json:"DstTpVer,omitempty" name:"DstTpVer"`
	// 任务结束时间。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 本次迁移任务如果是由存储池踢盘事件引起的，则EventUuid表示对应的踢盘事件；否则，EventUuid取值为"N/A。

	EventUuid *string `json:"EventUuid,omitempty" name:"EventUuid"`
	// zone id。

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 源cell节点ip。

	SrcCellIp *string `json:"SrcCellIp,omitempty" name:"SrcCellIp"`
	// 源盘smcd。

	SrcSmcd *uint64 `json:"SrcSmcd,omitempty" name:"SrcSmcd"`
	// 源小表号。

	SrcTpId *string `json:"SrcTpId,omitempty" name:"SrcTpId"`
	// 迁移任务开始时间。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 小表都是三副本的，TabletIndex指示当前迁移的是哪个副本，取值范围：0，1，2。

	TabletIndex *uint64 `json:"TabletIndex,omitempty" name:"TabletIndex"`
	// 迁移任务状态，取值范围：RUNNING 运行中；

	TaskState *string `json:"TaskState,omitempty" name:"TaskState"`
	// 串联路由版本号。

	TrsfRouteVer *string `json:"TrsfRouteVer,omitempty" name:"TrsfRouteVer"`
	// 迁移任务类型，取消范围：<br><li>MOVE: 迁移<br><li>RECOVERY: 恢复

	TrsfType *string `json:"TrsfType,omitempty" name:"TrsfType"`
	// 迁移任务uuid。

	TrsfUuid *string `json:"TrsfUuid,omitempty" name:"TrsfUuid"`
	// 任务耗时，单位秒

	CostTime *uint64 `json:"CostTime,omitempty" name:"CostTime"`
}

type DescribeCustomerCapacityTopChangeStartDateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomerCapacityTopChangeStartDateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerCapacityTopChangeStartDateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskFlowControlRecordsRequest struct {
	*tchttp.BaseRequest

	// 云盘UUID

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDiskFlowControlRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskFlowControlRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotResourceUsageRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>region - Array of String - 是否必填：否 -（过滤条件）按地域过滤。<br><li>start-date - Array of String - 是否必填：否 -（过滤条件）开始日期，如：2020-01-01。<br><li>end-date - Array of String - 是否必填：否 -（过滤条件） 结束日期，如：2020-02-01。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 支持按DAY/WEEK/MONTH统计；默认按天统计

	Period *string `json:"Period,omitempty" name:"Period"`
}

func (r *DescribeSnapshotResourceUsageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotResourceUsageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDiskUsageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 用户资源统计

		UserDiskUsageSet []*UserDiskUsage `json:"UserDiskUsageSet,omitempty" name:"UserDiskUsageSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserDiskUsageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserDiskUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotBoxOverViewRequest struct {
	*tchttp.BaseRequest

	// 快照集群ID

	BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
}

func (r *DescribeSnapshotBoxOverViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotBoxOverViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallDiskStorageDepotMasterRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储集群ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// Master机器IP，需要两台机器，作为主备

	DeviceIps []*string `json:"DeviceIps,omitempty" name:"DeviceIps"`
	// 同一组件使用机器的容灾级别要求，传入此参数，上架前，会进行容灾级别的检验，不满足不会进行上架，取值范围：<br><li>any: 无要求<br><li>motherMachine: 跨母机，仅针对虚拟机生效<br><li>rack: 跨机架<br><li>switch: 跨交换机

	DisasterTolerance *string `json:"DisasterTolerance,omitempty" name:"DisasterTolerance"`
}

func (r *InstallDiskStorageDepotMasterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallDiskStorageDepotMasterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DepotBalanceTask struct {
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储集群ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 当前剩余未均衡的disk pair数量

	RemainCpNum *uint64 `json:"RemainCpNum,omitempty" name:"RemainCpNum"`
	// 自动均衡任务的开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

type MachineStateOverview struct {
	// 组件类型

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 机器状态详情

	StatusSet []*MachineStateStatus `json:"StatusSet,omitempty" name:"StatusSet"`
}

type SetIOStat struct {
	// 存储池集群uuid

	SetUuid *string `json:"SetUuid,omitempty" name:"SetUuid"`
	// 存储池io util，单位百分比

	Util []*float64 `json:"Util,omitempty" name:"Util"`
	// 存储池io svctm，单位毫秒

	Svctm []*float64 `json:"Svctm,omitempty" name:"Svctm"`
	// 存储池平均写io次数，单位个/秒

	Wcnt []*float64 `json:"Wcnt,omitempty" name:"Wcnt"`
	// 存储池平均io队列深度

	Avgqu []*float64 `json:"Avgqu,omitempty" name:"Avgqu"`
	// 存储池平均写带宽，单位MB/秒

	Wband []*float64 `json:"Wband,omitempty" name:"Wband"`
	// 存储池io await，单位毫秒

	Await []*float64 `json:"Await,omitempty" name:"Await"`
	// 存储池io平均数据大小，单位扇区数/秒

	Avgrq []*float64 `json:"Avgrq,omitempty" name:"Avgrq"`
	// 存储池平均读带宽，单位MB/秒

	Rband []*float64 `json:"Rband,omitempty" name:"Rband"`
	// 存储池平均读io次数，单位个/秒

	Rcnt []*float64 `json:"Rcnt,omitempty" name:"Rcnt"`
	// io写await。

	Wawait []*float64 `json:"Wawait,omitempty" name:"Wawait"`
	// io读await。

	Rawait []*float64 `json:"Rawait,omitempty" name:"Rawait"`
}

type DescribeCbsAlarmEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的云硬盘数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 云硬盘告警详情列表

		AlarmEventSet []*AlarmEvent `json:"AlarmEventSet,omitempty" name:"AlarmEventSet"`
		// 返回数量

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 偏移量

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCbsAlarmEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCbsAlarmEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskModelConfigForSale struct {
	// 当前售卖配置是否生效

	Available *uint64 `json:"Available,omitempty" name:"Available"`
	// 用户可见白名单：不在该白名单中的用户无法看到该种售卖类型的云硬盘

	VisibleWhiteKey *string `json:"VisibleWhiteKey,omitempty" name:"VisibleWhiteKey"`
	// `FORCE_SOLD_OUT`：强制售罄，所有用户都不可购买；`AUTO_SOLD_OUT`：由系统资源余量决定；`WHITE_SOLD_ONLY`：仅白名单用户可购买；`VIP_SOLD_ONLY`：仅VIP用户可购买

	ForceSoldFlag *string `json:"ForceSoldFlag,omitempty" name:"ForceSoldFlag"`
	// `SYSTEM_DISK`：系统盘；`DATA_DISK`：数据盘

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 该规格允许的云硬盘最小大小

	MiniDiskSize *uint64 `json:"MiniDiskSize,omitempty" name:"MiniDiskSize"`
	// 该规格允许的云硬盘最大大

	MaxiDiskSize *uint64 `json:"MaxiDiskSize,omitempty" name:"MaxiDiskSize"`
	// 售罄标记

	AutoSoldFlag *string `json:"AutoSoldFlag,omitempty" name:"AutoSoldFlag"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 计费类型

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 售罄白名单

	ForceSoldWhiteKey *string `json:"ForceSoldWhiteKey,omitempty" name:"ForceSoldWhiteKey"`
	// 主键id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 型号

	ModelType *string `json:"ModelType,omitempty" name:"ModelType"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 可用区id

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeAutoSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 有效的定期快照策略数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 定期快照策略列表。

		AutoSnapshotPolicySet []*AutoSnapshotPolicy `json:"AutoSnapshotPolicySet,omitempty" name:"AutoSnapshotPolicySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoSnapshotPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskOperationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务列表

		ResOpLogSet []*ResOpLog `json:"ResOpLogSet,omitempty" name:"ResOpLogSet"`
		// 符合条件的任务数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskOperationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskOperationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDepotConfigRequest struct {
	*tchttp.BaseRequest

	// 存储池ID，存储池的唯一标识。

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 配置项所在的模块名，取值范围：master, cell, dbtrsf.

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
	// 需要查询的配置项，由section和item组成，通过冒号“:”分割，如:"log:log_level"。

	ConfigItem *string `json:"ConfigItem,omitempty" name:"ConfigItem"`
	// 新的配置值。

	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`
}

func (r *UpdateDepotConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDepotConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDepotDiskDetailsRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储仓库ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
}

func (r *DescribeDepotDiskDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepotDiskDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDisksMonitorDataOverviewRequest struct {
	*tchttp.BaseRequest

	// 取固定值`CBS`

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 监控数据类型，可选值：`DEPOT`，`DEPOT_CELL`， `DEPOT_PHYSICAL_DISK`， `CLOUD_DISK`

	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
	// 起始时间，如2018-09-22T19:51:23+08:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，默认为当前时间。 EndTime不能小于EtartTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 已废弃，用OrderField代替

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 取值范围：MINUTE，指定是从分钟级监控系统查询数据；SECOND，从秒级监控系统查询数据；<br><li>默认取值为MINUTE<br><li>只有MonitorType取值为DEPOT或CLOUD_DISK，才支持MonitorSystemType传入SECOND

	MonitorSystemType *string `json:"MonitorSystemType,omitempty" name:"MonitorSystemType"`
	// 监控数据的粒度，取值范围与入参MonitorSystemType相关：<br><li>当MonitorSystemType="MINUTE"，只能取值为60；<br><li>当MonitorSystemType="SECOND"，可以取值为10， 60， 300；

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 排序字段，默认使用`wband`写带宽排序，支持"await", "rcnt", "wcnt", "rband", "wband", "util", "svctm"排序

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeDisksMonitorDataOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisksMonitorDataOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataPoint struct {
	// 监控维度

	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions"`
	// 监控指标时间戳

	Timestamps []*string `json:"Timestamps,omitempty" name:"Timestamps"`
	// 监控值

	Values []*float64 `json:"Values,omitempty" name:"Values"`
}

type UserDiskUsage struct {
	// 云硬盘总大小

	TotalDiskSize *float64 `json:"TotalDiskSize,omitempty" name:"TotalDiskSize"`
	// 云硬盘总数量

	TotalDiskCount *float64 `json:"TotalDiskCount,omitempty" name:"TotalDiskCount"`
	// 快照总大小

	TotalSnapSize *float64 `json:"TotalSnapSize,omitempty" name:"TotalSnapSize"`
	// 快照总数量

	TotalSnapCount *float64 `json:"TotalSnapCount,omitempty" name:"TotalSnapCount"`
	// 可用的云硬盘资源池列表

	DiskStoragePoolGroupSet []*DiskStoragePoolGroup `json:"DiskStoragePoolGroupSet,omitempty" name:"DiskStoragePoolGroupSet"`
	// 用户AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 用户UIN

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// SSD容量

	Cbsssd *int64 `json:"Cbsssd,omitempty" name:"Cbsssd"`
	// HSSD容量

	CbsHSSD *int64 `json:"CbsHSSD,omitempty" name:"CbsHSSD"`
	// 高性能容量

	CbsPremium *int64 `json:"CbsPremium,omitempty" name:"CbsPremium"`
}

type CancelDepotTransferTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelDepotTransferTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelDepotTransferTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskStoragePoolGroupBoundRequest struct {
	*tchttp.BaseRequest

	// 当前支持 zone-id/app-id/disk-type/depot-attribute-id 过滤

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDiskStoragePoolGroupBoundRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskStoragePoolGroupBoundRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySnapshotAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySnapshotAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySnapshotAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplySnapshotRequest struct {
	*tchttp.BaseRequest

	// 快照ID, 可通过[DescribeSnapshots](/document/product/362/15647)查询。

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 快照原云硬盘ID，可通过[DescribeDisks](/document/product/362/16315)接口查询。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

func (r *ApplySnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplySnapshotRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSnapshotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新创建的快照ID。

		SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBalanceDepotTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 正在进行自动均衡的存储集群进度详情

		DepotBalanceTaskSet []*DepotBalanceTask `json:"DepotBalanceTaskSet,omitempty" name:"DepotBalanceTaskSet"`
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBalanceDepotTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBalanceDepotTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDisksMonitorAlarmRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 告警规则详情

		DiskAlarmRuleSet []*DiskAlarmRule `json:"DiskAlarmRuleSet,omitempty" name:"DiskAlarmRuleSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDisksMonitorAlarmRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisksMonitorAlarmRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeCommonComponentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务ID

		DesTaskId []*uint64 `json:"DesTaskId,omitempty" name:"DesTaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeCommonComponentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeCommonComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDepotCellDeviceConfigRequest struct {
	*tchttp.BaseRequest

	// 机型型号

	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
}

func (r *DeleteDepotCellDeviceConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDepotCellDeviceConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterUpgradeTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 升级任务列表

		ClusterUpgradeTaskSet []*ClusterUpgradeTask `json:"ClusterUpgradeTaskSet,omitempty" name:"ClusterUpgradeTaskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterUpgradeTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterUpgradeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoSnapshotPolicy struct {
	// 定期快照策略ID。

	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
	// 定期快照策略名称。

	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" name:"AutoSnapshotPolicyName"`
	// 定期快照策略的状态。取值范围：<br><li>NORMAL：正常<br><li>ISOLATED：已隔离。

	AutoSnapshotPolicyState *string `json:"AutoSnapshotPolicyState,omitempty" name:"AutoSnapshotPolicyState"`
	// 定期快照策略是否激活。

	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`
	// 使用该定期快照策略创建出来的快照是否永久保留。

	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`
	// 使用该定期快照策略创建出来的快照保留天数。

	RetentionDays *uint64 `json:"RetentionDays,omitempty" name:"RetentionDays"`
	// 定期快照策略的创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 定期快照下次触发的时间。

	NextTriggerTime *string `json:"NextTriggerTime,omitempty" name:"NextTriggerTime"`
	// 定期快照的执行策略。

	Policy []*Policy `json:"Policy,omitempty" name:"Policy"`
	// 已绑定当前定期快照策略的云盘ID列表。

	DiskIdSet []*string `json:"DiskIdSet,omitempty" name:"DiskIdSet"`
	// 用户appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type ReplaceDiskTask struct {
	// 底层zoneId。

	CbsZoneId *uint64 `json:"CbsZoneId,omitempty" name:"CbsZoneId"`
	// 存储节点ip。

	CellIp *string `json:"CellIp,omitempty" name:"CellIp"`
	// cell pair id。

	CellPairId *uint64 `json:"CellPairId,omitempty" name:"CellPairId"`
	// 任务创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 存储池ID。

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 任务结束时间。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 任务详情描述。

	JobDesc *string `json:"JobDesc,omitempty" name:"JobDesc"`
	// 任务ID。

	JobId *uint64 `json:"JobId,omitempty" name:"JobId"`
	// 任务状态。<br><li>finish: 任务完成<br><li>replacing：更换中<br><li>wait_check：等待确认<br><li>check_finish：确认完成<br><li>：取消

	JobStatus *string `json:"JobStatus,omitempty" name:"JobStatus"`
	// 任务类型。<br><li>REPLACE_MAIN_DISK: 主存替换<br><li>REPLACE_CACHE_DISK：缓存替换<br><li>NO_NEED_REPLACE：无需换盘<br><li>NEED_MANUAL_CHECK：人工确认

	JobType *string `json:"JobType,omitempty" name:"JobType"`
	// 新盘的盘符。

	NewDev *string `json:"NewDev,omitempty" name:"NewDev"`
	// 新盘sn号。

	NewSn *string `json:"NewSn,omitempty" name:"NewSn"`
	// 新盘wwn。

	NewWwn *string `json:"NewWwn,omitempty" name:"NewWwn"`
	// 原盘盘符。

	OldDev *string `json:"OldDev,omitempty" name:"OldDev"`
	// 原盘sn号。

	OldSn *string `json:"OldSn,omitempty" name:"OldSn"`
	// 原盘wwn。

	OldWwn *string `json:"OldWwn,omitempty" name:"OldWwn"`
	// zone id。

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// smcd。

	Smcd *uint64 `json:"Smcd,omitempty" name:"Smcd"`
	// 任务开始时间。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 任务的换盘状态。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 任务总步骤数。

	StepAll *uint64 `json:"StepAll,omitempty" name:"StepAll"`
	// 当前步骤信息。

	StepInfo *string `json:"StepInfo,omitempty" name:"StepInfo"`
	// 当前步骤序号。

	StepNow *uint64 `json:"StepNow,omitempty" name:"StepNow"`
	// 换盘任务的类型。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 任务执行人。

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 更换盘所在的槽位号。

	DiskSlot *string `json:"DiskSlot,omitempty" name:"DiskSlot"`
	// 本次换盘任务是否支持热插拔。true表示支持热插拔，即不用关机换盘；false表示需要关机换盘

	SupportDiskHotSwap *bool `json:"SupportDiskHotSwap,omitempty" name:"SupportDiskHotSwap"`
	// 存储池名称。

	DepotName *string `json:"DepotName,omitempty" name:"DepotName"`
	// 本次换盘事件ID。

	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`
}

type CreateSnapshotRequest struct {
	*tchttp.BaseRequest

	// 需要创建快照的云硬盘ID，可通过[DescribeDisks](/document/product/362/16315)接口查询。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 快照名称，不传则新快照名称默认为“未命名”。

	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
	// 从运营端发起快照创建时，本参数必传，取值为TRUE。表示从运营端直接发起快照创建。

	InternalDirectCreate *bool `json:"InternalDirectCreate,omitempty" name:"InternalDirectCreate"`
	// 快照的项目ID，不传则默认与云盘项目ID保持一致。

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 快照绑定的标签列表，不传则默认与云盘标签保持一致。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// TCE环境的项目ID，与公有云的ProjectId不相关，不传则与云盘项目保持一致。

	PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
}

func (r *CreateSnapshotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSnapshotRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskResourceOverviewRequest struct {
	*tchttp.BaseRequest

	// 过滤参数，仅支持根据“region”可用区进行过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDiskResourceOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskResourceOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDepotMasterConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDepotMasterConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDepotMasterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotsRequest struct {
	*tchttp.BaseRequest

	// 要查询快照的ID列表。参数不支持同时指定`SnapshotIds`和`Filters`。

	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
	// 过滤条件。参数不支持同时指定`SnapshotIds`和`Filters`。<br><li>snapshot-id - Array of String - 是否必填：否 -（过滤条件）按照快照的ID过滤。快照ID形如：`snap-11112222`。<br><li>snapshot-name - Array of String - 是否必填：否 -（过滤条件）按照快照名称过滤。<br><li>snapshot-state - Array of String - 是否必填：否 -（过滤条件）按照快照状态过滤。 (NORMAL：正常 | CREATING：创建中 | ROLLBACKING：回滚中。)<br><li>disk-usage - Array of String - 是否必填：否 -（过滤条件）按创建快照的云盘类型过滤。 (SYSTEM_DISK：代表系统盘 | DATA_DISK：代表数据盘。)<br><li>project-id  - Array of String - 是否必填：否 -（过滤条件）按云硬盘所属项目ID过滤。<br><li>disk-id  - Array of String - 是否必填：否 -（过滤条件）按照创建快照的云硬盘ID过滤。<br><li>zone - Array of String - 是否必填：否 -（过滤条件）按照[可用区](/document/api/213/9452#zone)过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 输出云盘列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 快照列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据快照的创建时间排序<br>默认按创建时间排序。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeSnapshotsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotSharePermissionRequest struct {
	*tchttp.BaseRequest

	// 要查询快照的ID。可通过DescribeSnapshots查询获取。

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

func (r *DescribeSnapshotSharePermissionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotSharePermissionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DriverStateInfo struct {
	// 挂载到存储池的母机总数

	DriverAllNum *uint64 `json:"DriverAllNum,omitempty" name:"DriverAllNum"`
	// 挂载到当前存储池，且cbs driver异常的母机数理

	DriverDownNum *uint64 `json:"DriverDownNum,omitempty" name:"DriverDownNum"`
	// 挂载到当前存储池，且cbs driver正常的母机数理

	DriverNormalNum *uint64 `json:"DriverNormalNum,omitempty" name:"DriverNormalNum"`
}

type DescribeDiskFlowControlRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的云硬盘流控记录数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 云硬盘流控记录详情

		DiskFlowControlRecordSet []*DiskFlowControlRecord `json:"DiskFlowControlRecordSet,omitempty" name:"DiskFlowControlRecordSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskFlowControlRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskFlowControlRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskIOStat struct {
	// 云盘uuid

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
	// 云盘io util，单位百分比

	Util []*float64 `json:"Util,omitempty" name:"Util"`
	// 云盘io svctm，单位毫秒

	Svctm []*float64 `json:"Svctm,omitempty" name:"Svctm"`
	// 云盘平均写io次数，单位个/秒

	Wcnt []*float64 `json:"Wcnt,omitempty" name:"Wcnt"`
	// 云盘平均io队列深度

	Avgqu []*float64 `json:"Avgqu,omitempty" name:"Avgqu"`
	// 云盘平均写带宽，单位MB/秒

	Wband []*float64 `json:"Wband,omitempty" name:"Wband"`
	// 云盘io await，单位毫秒

	Await []*float64 `json:"Await,omitempty" name:"Await"`
	// 云盘io平均数据大小，单位扇区数/秒

	Avgrq []*float64 `json:"Avgrq,omitempty" name:"Avgrq"`
	// 云盘平均读带宽，单位MB/秒

	Rband []*float64 `json:"Rband,omitempty" name:"Rband"`
	// 云盘平均读io次数，单位个/秒

	Rcnt []*float64 `json:"Rcnt,omitempty" name:"Rcnt"`
	// 读await。

	Rawait []*float64 `json:"Rawait,omitempty" name:"Rawait"`
	// 写await。

	Wawait []*float64 `json:"Wawait,omitempty" name:"Wawait"`
}

type CancelClusterUpgradeTasksRequest struct {
	*tchttp.BaseRequest

	// 标识一次集群升级的事务ID

	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`
}

func (r *CancelClusterUpgradeTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelClusterUpgradeTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskResourceOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存储仓库资源使用概览

		DepotOverview *DepotOverview `json:"DepotOverview,omitempty" name:"DepotOverview"`
		// 按类型的存储仓库资源使用概览

		DepotOverviewByDiskTypeSet []*DepotOverview `json:"DepotOverviewByDiskTypeSet,omitempty" name:"DepotOverviewByDiskTypeSet"`
		// 云硬盘资源使用概览

		DiskOverview *DiskOverview `json:"DiskOverview,omitempty" name:"DiskOverview"`
		// 按类型的云硬盘资源使用概览

		DiskOverviewByDiskTypeSet []*DiskOverview `json:"DiskOverviewByDiskTypeSet,omitempty" name:"DiskOverviewByDiskTypeSet"`
		// 快照资源使用概览

		SnapOverview *SnapOverview `json:"SnapOverview,omitempty" name:"SnapOverview"`
		// 分地域的资源使用概览

		ResourcesOverviewSet []*ResourceOverview `json:"ResourcesOverviewSet,omitempty" name:"ResourcesOverviewSet"`
		// 各状态zk集群概览

		ZKClusterOverview []*ZKClusterOverview `json:"ZKClusterOverview,omitempty" name:"ZKClusterOverview"`
		// 各组件机器状态详情

		MachineStateOverview []*MachineStateOverview `json:"MachineStateOverview,omitempty" name:"MachineStateOverview"`
		// 事件详情

		TopDepotReportEvent []*TopDepotReportEvent `json:"TopDepotReportEvent,omitempty" name:"TopDepotReportEvent"`
		// 快照集群概览

		SnapshotBoxOverview []*SnapshotBoxOverview `json:"SnapshotBoxOverview,omitempty" name:"SnapshotBoxOverview"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskResourceOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskResourceOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecycleRulesRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储池ID，查询存储池回收规则时传入

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 云盘uuid，查询云盘回收规则时传入

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
}

func (r *DescribeRecycleRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecycleRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MachineStateStatus struct {
	// 机器数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 机器ip列表

	DeviceSet []*string `json:"DeviceSet,omitempty" name:"DeviceSet"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DescribeDepotCellDeviceConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存储机型配置详情

		DepotCellDeviceConfigSet []*DepotCellDeviceConfig `json:"DepotCellDeviceConfigSet,omitempty" name:"DepotCellDeviceConfigSet"`
		// 符合条件的总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDepotCellDeviceConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepotCellDeviceConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskAssociatedAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云盘绑定的定期快照数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 云盘绑定的定期快照列表。

		AutoSnapshotPolicySet []*AutoSnapshotPolicy `json:"AutoSnapshotPolicySet,omitempty" name:"AutoSnapshotPolicySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskAssociatedAutoSnapshotPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskAssociatedAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MainDisk struct {
	// 主存盘sn号

	MainDiskSn *string `json:"MainDiskSn,omitempty" name:"MainDiskSn"`
	// 主存盘smcd

	Smcd *uint64 `json:"Smcd,omitempty" name:"Smcd"`
	// 主存盘的替换状态。<br><li>NO_NEED_REPLACE: 正常<br><li>REPLACE_MAIN_DISK：主存替换<br><li>REPLACE_CACHE_DISK：缓存替换<br><li>REPLACING：替换中

	DiskReplaceState *string `json:"DiskReplaceState,omitempty" name:"DiskReplaceState"`
}

type InstallCapacityServerRequest struct {
	*tchttp.BaseRequest

	// 要安装容量server的快照box id

	BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
	// 要安装容量server的机器IP列表

	DeviceIps []*string `json:"DeviceIps,omitempty" name:"DeviceIps"`
	// 是否真正执行安装操作，1表示不支持安装操作，仅做检查；0表示会执行实际安装，默认取值为0

	DryRun *uint64 `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *InstallCapacityServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallCapacityServerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskAlarmRule struct {
	// 告警类型，取值范围：iohang, svctm_high, avg_svctm_high

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 云盘类型，取值范围：<br><li>CLOUD_BASIC: 普通云盘；<br><li>CLOUD_SSD,CLOUD_PREMIUM: 高性能云盘和SSD云盘

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 告警对象，取值范围：<br><li>disk: 云盘<br><li>set:存储池

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
	// 告警规则详情

	DiskAlarmRuleDetailSet []*DiskAlarmRuleDetail `json:"DiskAlarmRuleDetailSet,omitempty" name:"DiskAlarmRuleDetailSet"`
}

type DescribeDiskOperationsV2Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作记录条目数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 操作记录详情

		ResourceOperationLogSet []*ResourceOperationLog `json:"ResourceOperationLogSet,omitempty" name:"ResourceOperationLogSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskOperationsV2Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskOperationsV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotGroupsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>snapshot-group-id - Array of String - 是否必填：否 -（过滤条件）按快照组ID过滤 <br><li>snapshot-group-state - Array of String - 是否必填：否 -（过滤条件）按快照组状态过滤。(NORMAL: 正常 | CREATING:创建中 | ROLLBACKING:回滚中) <br><li>snapshot-group-name - Array of String - 是否必填：否 -（过滤条件）按快照组名称过滤 <br><li>snapshot-id - Array of String - 是否必填：否 -（过滤条件）按快照组内的快照ID过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSnapshotGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallDiskStorageDepotCellRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储集群ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// CELL物理服务器IP，一次需要下架一组CELL（三台物理服务器

	DeviceIps []*string `json:"DeviceIps,omitempty" name:"DeviceIps"`
	// 是否强制卸载。0：不强制卸载；1：强制卸载

	ForceUninstall *uint64 `json:"ForceUninstall,omitempty" name:"ForceUninstall"`
}

func (r *UninstallDiskStorageDepotCellRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallDiskStorageDepotCellRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DepotDiskDetails struct {
	// 存储仓库ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// vol版本号

	DiskVer *uint64 `json:"DiskVer,omitempty" name:"DiskVer"`
	// Master IP

	MasterIp *string `json:"MasterIp,omitempty" name:"MasterIp"`
	// Master版本

	MasterVer *string `json:"MasterVer,omitempty" name:"MasterVer"`
	// 存储仓库最大可创建云盘数量

	MaxDiskNum *uint64 `json:"MaxDiskNum,omitempty" name:"MaxDiskNum"`
	// 超卖比

	OverSoldRatio *float64 `json:"OverSoldRatio,omitempty" name:"OverSoldRatio"`
	// 小表路由版本号

	RouteVer *uint64 `json:"RouteVer,omitempty" name:"RouteVer"`
	// 存储扇区大小

	SectorSize *uint64 `json:"SectorSize,omitempty" name:"SectorSize"`
	// 从Master IP

	SlaveIp *string `json:"SlaveIp,omitempty" name:"SlaveIp"`
	// 仓库总容量（GB）；total_capacity * over_sold_ratio - vol_statesize_info.vol_all_size

	TotalCapacity *uint64 `json:"TotalCapacity,omitempty" name:"TotalCapacity"`
	// 挂载到当前存储池的母机数量详情

	DriverStateInfo *DriverStateInfo `json:"DriverStateInfo,omitempty" name:"DriverStateInfo"`
	// 存储池各状态云盘数量详情

	DiskStateInfo *DiskStateInfo `json:"DiskStateInfo,omitempty" name:"DiskStateInfo"`
	// 存储池各状态云盘容量详情

	DiskStatesizeInfo *DiskStatesizeInfo `json:"DiskStatesizeInfo,omitempty" name:"DiskStatesizeInfo"`
	// 标识当前存储池小表是否均衡，取值：1表示已均衡，0表示不均衡

	TpBalanceFlag *uint64 `json:"TpBalanceFlag,omitempty" name:"TpBalanceFlag"`
}

type DepotSoldStateSet struct {
	// 售卖状态

	DepotSoldState *string `json:"DepotSoldState,omitempty" name:"DepotSoldState"`
	// 数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type SnapshotBoxOverview struct {
	// 快照集群数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 快照状态

	SnapBoxState *string `json:"SnapBoxState,omitempty" name:"SnapBoxState"`
}

type UninstallTransferPairsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务ID，可在运维任务页面查看任务执行结果。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 任务描述。

		TaskDescription *string `json:"TaskDescription,omitempty" name:"TaskDescription"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UninstallTransferPairsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallTransferPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskStatesizeInfo struct {
	// 存储池总云盘容量

	DiskAllSize *uint64 `json:"DiskAllSize,omitempty" name:"DiskAllSize"`
	// 存储池已挂载云盘容量

	DiskAttachedSize *uint64 `json:"DiskAttachedSize,omitempty" name:"DiskAttachedSize"`
	// 存储池挂载中云盘容量

	DiskAttachingSize *uint64 `json:"DiskAttachingSize,omitempty" name:"DiskAttachingSize"`
	// 存储池已删除云盘容量

	DiskDeletedSize *uint64 `json:"DiskDeletedSize,omitempty" name:"DiskDeletedSize"`
	// 存储池已解挂云盘容量

	DiskDettachedSize *uint64 `json:"DiskDettachedSize,omitempty" name:"DiskDettachedSize"`
	// 存储池解挂中云盘容量

	DiskDettachingSize *uint64 `json:"DiskDettachingSize,omitempty" name:"DiskDettachingSize"`
	// 存储池扩容中云盘容量

	DiskEditingSize *uint64 `json:"DiskEditingSize,omitempty" name:"DiskEditingSize"`
	// 存储池已回收云盘容量

	DiskRecycledSize *uint64 `json:"DiskRecycledSize,omitempty" name:"DiskRecycledSize"`
	// 存储池回收中云盘容量

	DiskRecyclingSize *uint64 `json:"DiskRecyclingSize,omitempty" name:"DiskRecyclingSize"`
	// 存储池迁移中云盘容量

	DiskTrsfingSize *uint64 `json:"DiskTrsfingSize,omitempty" name:"DiskTrsfingSize"`
}

type DescribeDiskMigrateTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的迁移任务总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 迁移任务列表

		DiskMigrateTaskSet []*DiskMigrateTask `json:"DiskMigrateTaskSet,omitempty" name:"DiskMigrateTaskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskMigrateTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskMigrateTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallDiskStorageDepotCellResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID列表

		TaskIdList []*uint64 `json:"TaskIdList,omitempty" name:"TaskIdList"`
		// 任务ID列表的第一个元素。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 任务描述。

		TaskDescription *string `json:"TaskDescription,omitempty" name:"TaskDescription"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InstallDiskStorageDepotCellResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallDiskStorageDepotCellResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopDiskStorageDepotCellResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务ID，可在运维任务页面查看任务执行结果。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopDiskStorageDepotCellResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopDiskStorageDepotCellResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotBoxConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置详情

		SnapshotBoxConfigSet []*ClusterConfigDetail `json:"SnapshotBoxConfigSet,omitempty" name:"SnapshotBoxConfigSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotBoxConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotBoxConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CellTpStateInfo struct {
	// 总小表数

	TpAllNum *uint64 `json:"TpAllNum,omitempty" name:"TpAllNum"`
	// 已使用小表数

	TpUsedNum *uint64 `json:"TpUsedNum,omitempty" name:"TpUsedNum"`
	// free小表数

	TpFreeNum *uint64 `json:"TpFreeNum,omitempty" name:"TpFreeNum"`
}

type BindDiskStorageDepotToDiskStoragePoolGroupRequest struct {
	*tchttp.BaseRequest

	// ss

	DepotIds []*uint64 `json:"DepotIds,omitempty" name:"DepotIds"`
	// aa

	DepotAttributeId *uint64 `json:"DepotAttributeId,omitempty" name:"DepotAttributeId"`
}

func (r *BindDiskStorageDepotToDiskStoragePoolGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindDiskStorageDepotToDiskStoragePoolGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSetCellSideIOStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 时间序列

		Times []*string `json:"Times,omitempty" name:"Times"`
		// 存储池IO监控详情

		SetCellSideIOStatSet []*CellSideSetIOStat `json:"SetCellSideIOStatSet,omitempty" name:"SetCellSideIOStatSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSetCellSideIOStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSetCellSideIOStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachDisksRequest struct {
	*tchttp.BaseRequest

	// 将要解挂的云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询，单次请求最多可解挂20块弹性云盘。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 强制解挂，当前仅用于从物理机卸载云盘

	ForceDetach *bool `json:"ForceDetach,omitempty" name:"ForceDetach"`
}

func (r *DetachDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyDiskAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 描述了变更云盘类型的价格。

		DiskPrice *Price `json:"DiskPrice,omitempty" name:"DiskPrice"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceModifyDiskAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceModifyDiskAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskDepotComponentDetail struct {
	// dbtrsf组件所在机器IP

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// dbtrsf组件使用的端口号

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 上架人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 最后修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 最后修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 上架时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// dbtrsf组件状态。

	State *string `json:"State,omitempty" name:"State"`
	// dbtrsf组件所在机器的机型

	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
	// dbtrsf组件的版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// dbtrsf组件所在机器所处的机架

	RackName *string `json:"RackName,omitempty" name:"RackName"`
	// 当前模块的升级状态。取值范围：<br><li>WAIT_CHECK: 等待检查<br><li>CHECK_FINISH: 检查完成<br><li>CHECK_FAILED: 检查失败<br><li>RUNNING: 升级中<br><li>SUCCESS: 升级成功<br><li>FAILED: 升级失败<br><li>ROLLBACKED: 已回滚

	UpgradeState *string `json:"UpgradeState,omitempty" name:"UpgradeState"`
	// 组件的上次升级时间

	LastUpgradeTime *string `json:"LastUpgradeTime,omitempty" name:"LastUpgradeTime"`
	// 机器所在上联交换机设备名称

	SwitchName *string `json:"SwitchName,omitempty" name:"SwitchName"`
	// 标识主备，取值为1表示为主，取值为0表示备；如果为null，表示该组件不分主备

	IsMaster *uint64 `json:"IsMaster,omitempty" name:"IsMaster"`
	// 可用区id

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 组件类型

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
	// 机器的cpu型号

	DeviceCpuModel *string `json:"DeviceCpuModel,omitempty" name:"DeviceCpuModel"`
}

type ApplySnapshotResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplySnapshotResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplySnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskOperationsV2Request struct {
	*tchttp.BaseRequest

	// 过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询截止时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDiskOperationsV2Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskOperationsV2Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskAlarmRuleDetail struct {
	// 最小值

	MinValue *float64 `json:"MinValue,omitempty" name:"MinValue"`
	// 最大值

	MaxValue *float64 `json:"MaxValue,omitempty" name:"MaxValue"`
	// 告警判断的指标，取值范围：rcnt,wcnt,rband,wband,util,svctm,avgrq

	Field *string `json:"Field,omitempty" name:"Field"`
	// 告警判断的持续周期

	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
	// 告警规则添加时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 告警规则更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 在告警判断的持续周期内，多少个周期异常会导致告警

	AbnormalNum *float64 `json:"AbnormalNum,omitempty" name:"AbnormalNum"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
}

type MigrateDiskTask struct {
	// 云盘迁移任务ID。

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 云盘UUID。

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
}

type TopDepotReportEvent struct {
	// 事件数量

	EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
	// 事件类型

	EventTypeStr *string `json:"EventTypeStr,omitempty" name:"EventTypeStr"`
}

type BindDiskStorageDepotToDiskStoragePoolGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindDiskStorageDepotToDiskStoragePoolGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindDiskStorageDepotToDiskStoragePoolGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoCloseAndOpenBlockRateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自动打开售卖的block利用率

		OpenBlockRate *uint64 `json:"OpenBlockRate,omitempty" name:"OpenBlockRate"`
		// 自动关闭售卖的block利用率

		CloseBlockRate *uint64 `json:"CloseBlockRate,omitempty" name:"CloseBlockRate"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoCloseAndOpenBlockRateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAutoCloseAndOpenBlockRateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateDisksRequest struct {
	*tchttp.BaseRequest

	// 云硬盘ID列表

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 云硬盘UUID列表

	DiskUuids []*string `json:"DiskUuids,omitempty" name:"DiskUuids"`
	// 目标存储集群ID

	DestinationDepotId *uint64 `json:"DestinationDepotId,omitempty" name:"DestinationDepotId"`
	// cbs与cfs存储池的盘能否互迁，取值范围：0，不能互迁；1：可以互迁。默认取值0

	AllowMixedMigration *uint64 `json:"AllowMixedMigration,omitempty" name:"AllowMixedMigration"`
	// 是否使用双写迁移；未挂载云硬盘只能使用单写迁移，只有上架了trsf_pair组件时才支持双写迁移。默认使用单写迁移

	DoubleWrite *uint64 `json:"DoubleWrite,omitempty" name:"DoubleWrite"`
}

func (r *MigrateDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateDisksRequest struct {
	*tchttp.BaseRequest

	// 需退还的云盘ID列表。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

func (r *TerminateDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallDiskZKClusterRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// ZooKeeper集群ID

	ZKClusterId *uint64 `json:"ZKClusterId,omitempty" name:"ZKClusterId"`
	// ZooKeeper集群节点IP列表

	DeviceIps []*string `json:"DeviceIps,omitempty" name:"DeviceIps"`
	// 是否强制卸载。0：不强制卸载；1：强制卸载

	ForceUninstall *uint64 `json:"ForceUninstall,omitempty" name:"ForceUninstall"`
}

func (r *UninstallDiskZKClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallDiskZKClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelMigrateDiskTasksRequest struct {
	*tchttp.BaseRequest

	// 需取消的迁移任务ID列表。

	TaskIds []*uint64 `json:"TaskIds,omitempty" name:"TaskIds"`
	// 需取消的云盘UUID列表。

	DiskUuids []*string `json:"DiskUuids,omitempty" name:"DiskUuids"`
}

func (r *CancelMigrateDiskTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelMigrateDiskTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoCloseAndOpenBlockRateRequest struct {
	*tchttp.BaseRequest

	// 自动关闭售卖的block利用率阈值，范围在[1,95]，必须比自动打开售卖的block利用率阈值高

	CloseBlockRate *int64 `json:"CloseBlockRate,omitempty" name:"CloseBlockRate"`
	// 自动打开售卖的block利用率阈值，范围在[1,90]

	OpenBlockRate *int64 `json:"OpenBlockRate,omitempty" name:"OpenBlockRate"`
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储池ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
}

func (r *ModifyAutoCloseAndOpenBlockRateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoCloseAndOpenBlockRateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMonitorEsIndexStorageTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMonitorEsIndexStorageTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMonitorEsIndexStorageTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallDiskStorageDepotMasterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务ID，可在运维任务页面查看任务执行结果。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UninstallDiskStorageDepotMasterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallDiskStorageDepotMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Metric struct {
	// 写流量，单位MB/s

	Wband *float64 `json:"Wband,omitempty" name:"Wband"`
	// 读流量，单位MB/s

	Rband *float64 `json:"Rband,omitempty" name:"Rband"`
	// 写IOPS

	Wcnt *uint64 `json:"Wcnt,omitempty" name:"Wcnt"`
	// 读IOPS

	Rcnt *uint64 `json:"Rcnt,omitempty" name:"Rcnt"`
	// IO wait

	Await *float64 `json:"Await,omitempty" name:"Await"`
	// IO svctm

	Svctm *float64 `json:"Svctm,omitempty" name:"Svctm"`
	// IO util

	Util *float64 `json:"Util,omitempty" name:"Util"`
	// 脏页百分比

	Dirtypage *int64 `json:"Dirtypage,omitempty" name:"Dirtypage"`
}

type MetricValue struct {
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标值列表

	ValueSet []*float64 `json:"ValueSet,omitempty" name:"ValueSet"`
}

type ZKNodeDetail struct {
	// zk节点ip

	ZKNodeIp *string `json:"ZKNodeIp,omitempty" name:"ZKNodeIp"`
	// zk节点状态。<br><li>NORMAL: 正常<br><li>FAULT：异常

	ZKNodeState *string `json:"ZKNodeState,omitempty" name:"ZKNodeState"`
	// zk节点所在机架

	RackName *string `json:"RackName,omitempty" name:"RackName"`
	// 当前zk节点是否为master节点

	IsMaster *bool `json:"IsMaster,omitempty" name:"IsMaster"`
	// zk节点所在可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 上联交换机

	SwitchName *string `json:"SwitchName,omitempty" name:"SwitchName"`
	// 机器的cpu型号

	DeviceCpuModel *string `json:"DeviceCpuModel,omitempty" name:"DeviceCpuModel"`
	// 机器架构类型，取值x86, arm

	DeviceArchType *string `json:"DeviceArchType,omitempty" name:"DeviceArchType"`
}

type CreateDepotCellDeviceConfigRequest struct {
	*tchttp.BaseRequest

	// 机型型号

	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
	// 存储池类型，取值范围：<br><li>hdd: 普通存储池<br><li>fusion: 高性能存储池<br><li>ssd: SSD存储池

	DepotMedia *string `json:"DepotMedia,omitempty" name:"DepotMedia"`
	// 存储机型主存盘数量

	DiskNum *uint64 `json:"DiskNum,omitempty" name:"DiskNum"`
	// 存储机型主存盘容量，单位TB

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 主存盘接口类型，取值范围: <br><li>sata: sata接口<br><li>nvme: nvme接口<br><li>sas: SAS接口

	PrimaryInfType *string `json:"PrimaryInfType,omitempty" name:"PrimaryInfType"`
	// 当前机型是否支持磁盘热插拔。true表示支持热插拔盘；false表示不支持热插拔盘

	SupportDiskHotSwap *bool `json:"SupportDiskHotSwap,omitempty" name:"SupportDiskHotSwap"`
	// 缓存盘接口类型，取值范围: <br><li>nvme: nvme接口<br><li>仅在有缓存盘时，传入该值

	CacheInfType *string `json:"CacheInfType,omitempty" name:"CacheInfType"`
	// 存储机型缓存盘数量，仅在有缓存盘时传入该值

	CacheDiskNum *uint64 `json:"CacheDiskNum,omitempty" name:"CacheDiskNum"`
	// 存储机型缓存盘大小，可能为小数，所以类型为字符串类型，单位TB。仅在有缓存盘时传入该值

	CacheDiskSize *string `json:"CacheDiskSize,omitempty" name:"CacheDiskSize"`
	// 存储机型的网络类型，取值范围：10Gb, 25Gb

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 机型支持上架存储池的副本策略，取值范围，THREE_COPIES：三副本存储，EC(9+3): 9+3 EC存储

	CopyStrategies []*string `json:"CopyStrategies,omitempty" name:"CopyStrategies"`
}

func (r *CreateDepotCellDeviceConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDepotCellDeviceConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindUserToDiskStoragePoolGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindUserToDiskStoragePoolGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindUserToDiskStoragePoolGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 迁移任务ID列表

		MigrateTaskId []*uint64 `json:"MigrateTaskId,omitempty" name:"MigrateTaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigrateDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Disk struct {
	// 云硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 云硬盘类型。取值范围：<br><li>SYSTEM_DISK：系统盘<br><li>DATA_DISK：数据盘。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 付费模式。取值范围：<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：后付费，即按量计费。

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 是否为弹性云盘，false表示非弹性云盘，true表示弹性云盘。

	// Portable *bool `json:"Portable,omitempty" name:"Portable"`
	Portable *interface{} `json:"Portable,omitempty" name:"Portable"`
	// 云硬盘所在的位置。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 云盘是否具备创建快照的能力。取值范围：<br><li>false表示不具备<br><li>true表示具备。

	SnapshotAbility *bool `json:"SnapshotAbility,omitempty" name:"SnapshotAbility"`
	// 云硬盘名称。

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 云硬盘大小。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 云盘状态。取值范围：<br><li>UNATTACHED：未挂载<br><li>ATTACHING：挂载中<br><li>ATTACHED：已挂载<br><li>DETACHING：解挂中<br><li>EXPANDING：扩容中<br><li>ROLLBACKING：回滚中。

	DiskState *string `json:"DiskState,omitempty" name:"DiskState"`
	// 云盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：SSD表示SSD云硬盘。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云盘是否挂载到云主机上。取值范围：<br><li>false:表示未挂载<br><li>true:表示已挂载。

	Attached *bool `json:"Attached,omitempty" name:"Attached"`
	// 云硬盘挂载的云主机ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 云硬盘的创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 云硬盘的到期时间。

	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`
	// 云盘是否处于快照回滚状态。取值范围：<br><li>0:表示不处于快照回滚状态<br><li>1:表示处于快照回滚状态。

	Rollbacking *uint64 `json:"Rollbacking,omitempty" name:"Rollbacking"`
	// 云盘快照回滚的进度。

	RollbackPercent *uint64 `json:"RollbackPercent,omitempty" name:"RollbackPercent"`
	// 云盘是否为加密盘。取值范围：<br><li>false:表示非加密盘<br><li>true:表示加密盘。

	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`
	// 云盘已挂载到子机，且子机与云盘都是包年包月。<br><li>true：子机设置了自动续费标识，但云盘未设置<br><li>false：云盘自动续费标识正常。

	AutoRenewFlagError *bool `json:"AutoRenewFlagError,omitempty" name:"AutoRenewFlagError"`
	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 在云盘已挂载到实例，且实例与云盘都是包年包月的条件下，此字段才有意义。<br><li>true:云盘到期时间早于实例。<br><li>false：云盘到期时间晚于实例。

	DeadlineError *bool `json:"DeadlineError,omitempty" name:"DeadlineError"`
	// 云盘关联的定期快照ID。只有在调用DescribeDisks接口时，入参ReturnBindAutoSnapshotPolicy取值为TRUE才会返回该参数。

	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitempty" name:"AutoSnapshotPolicyIds"`
	// 是否随实例销毁

	DeleteWithInstance *uint64 `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
	// 是否为共享盘

	Shareable *uint64 `json:"Shareable,omitempty" name:"Shareable"`
	// 用户appId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 盘挂载的母机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 是否为CFS存储池的云盘

	IsCfsDisk *uint64 `json:"IsCfsDisk,omitempty" name:"IsCfsDisk"`
	// 云盘所在存储池ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 云盘所在存储池名称

	DepotName *string `json:"DepotName,omitempty" name:"DepotName"`
	// 云盘的流控模式。<br><li>IN_BLACK_LIST:黑名单流控<br><li>IN_WHITE_LIST:白名单流控<br><li>AUTO：自动流控<br><li>IN_BLACK_WHITE_LIST: 同时处于黑名单、白名单流控中

	FlowControlState *string `json:"FlowControlState,omitempty" name:"FlowControlState"`
	// 云盘是否在迁移中

	Migrating *bool `json:"Migrating,omitempty" name:"Migrating"`
	// 云盘迁移的进度，只有当Migrating为true时才有意义

	MigratePercent *uint64 `json:"MigratePercent,omitempty" name:"MigratePercent"`
	// 云盘挂载实例的uuid

	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
	// 云盘挂载实例的名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 云盘挂载母机的driver版本，如果云盘未挂载，或查询不到版本号，则显示为“”

	DriverVersion *string `json:"DriverVersion,omitempty" name:"DriverVersion"`
	// 云盘uuid

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
	// 云盘实际已使用容量，单位GB

	DiskSizeUsed *int64 `json:"DiskSizeUsed,omitempty" name:"DiskSizeUsed"`
	// 当前时间与云盘到期时间相差的天数，仅去预付费盘有效

	DifferDaysOfDeadline *int64 `json:"DifferDaysOfDeadline,omitempty" name:"DifferDaysOfDeadline"`
	// 云盘所属资源池

	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
	// 共享盘挂载的云服务器实例信息

	ShareAttachedInstanceSet []*SimpleInstance `json:"ShareAttachedInstanceSet,omitempty" name:"ShareAttachedInstanceSet"`
	// 绑定定期快照策略的数量

	AutoSnapshotPolicyBoundCount *uint64 `json:"AutoSnapshotPolicyBoundCount,omitempty" name:"AutoSnapshotPolicyBoundCount"`
	// 购买云盘的订单名。

	DealName *string `json:"DealName,omitempty" name:"DealName"`
	// 云盘的底层运维状态。

	DiskRealState *string `json:"DiskRealState,omitempty" name:"DiskRealState"`
	// 云盘信息修改时间。

	ModifyTimeStamp *string `json:"ModifyTimeStamp,omitempty" name:"ModifyTimeStamp"`
	// 云盘挂载实例信息。

	Instance *DiskAttachedInstanceInfo `json:"Instance,omitempty" name:"Instance"`
	// 加密算法类型。SM4:国密；AES256:非国密

	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" name:"EncryptAlgorithm"`
}

type DescribeDiskOperationsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>disk-id - Array of String - 是否必填：否 -（过滤条件）按云盘实例ID过滤。<br><li>task-name - Array of String - 是否必填：否 -（过滤条件）按任务名称过滤。<br><li>task-id - Array of Integer - 是否必填：否 -（过滤条件）按任务ID过滤。<br><li>snapshot-id - Array of String - 是否必填：否 -（过滤条件）按快照ID过滤。<br><li>app-id - Array of Integer - 是否必填：否 -（过滤条件）按用户appId过滤。<br><li>cbs-uuid - Array of String - 是否必填：否 -（过滤条件）按云盘UUID过滤。<br><li>snap-uuid - Array of Integer - 是否必填：否 -（过滤条件）按快照UUID过滤。<br><li>ins-uuid - Array of String - 是否必填：否 -（过滤条件）按云服务实例ID过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为20，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDiskOperationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskOperationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeniedAction struct {
	// 不能操作接口名，比如查询云盘列表"DescribeDisks"，对于多用途的接口，会以"接口.入参"的形式返回。比如"ModifyDiskAttributes.DiskType"

	Action *string `json:"Action,omitempty" name:"Action"`
	// 接口不能操作的原因

	Message *string `json:"Message,omitempty" name:"Message"`
	// 接口不能操作对应提示的错误码

	Code *string `json:"Code,omitempty" name:"Code"`
}

type DescribeDepotFlowControlRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流控记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 流控记录详情列表

		DepotFlowControlRecordSet []*DepotFlowControlRecord `json:"DepotFlowControlRecordSet,omitempty" name:"DepotFlowControlRecordSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDepotFlowControlRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepotFlowControlRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskMaintenanceTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 有效的运维任务数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 运维任务列表

		DiskMaintenanceTaskSet []*DiskMaintenanceTask `json:"DiskMaintenanceTaskSet,omitempty" name:"DiskMaintenanceTaskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskMaintenanceTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskMaintenanceTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallBoxDevicesRequest struct {
	*tchttp.BaseRequest

	// 快照集群ID

	BoxId *uint64 `json:"BoxId,omitempty" name:"BoxId"`
	// 安装的机器IP列表

	DeviceIps []*string `json:"DeviceIps,omitempty" name:"DeviceIps"`
	// 安装的机器类型, 取值范围manager, scheduler, transfer

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
	// 取值为0时会执行实际上架；取值为1时不执行实际上架操作，仅检查操作合法性

	DryRun *uint64 `json:"DryRun,omitempty" name:"DryRun"`
	// 同一组件使用机器的容灾级别要求，传入此参数，上架前，会进行容灾级别的检验，不满足不会进行上架，取值范围：<br><li>any: 无要求<br><li>motherMachine: 跨母机，仅针对虚拟机生效<br><li>rack: 跨机架<br><li>switch: 跨交换机

	DisasterTolerance *string `json:"DisasterTolerance,omitempty" name:"DisasterTolerance"`
}

func (r *InstallBoxDevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallBoxDevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TransferDiskStorageDepotCellDiskPairResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 迁移任务发起结果描述

		TaskDescription *string `json:"TaskDescription,omitempty" name:"TaskDescription"`
		// 在单盘手动迁移的场景下，此值才有意义，表示最大可迁移的小表百分比

		MaxTransferTpRatio *uint64 `json:"MaxTransferTpRatio,omitempty" name:"MaxTransferTpRatio"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TransferDiskStorageDepotCellDiskPairResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TransferDiskStorageDepotCellDiskPairResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectDiskStorageDepotCellResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectDiskStorageDepotCellResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectDiskStorageDepotCellResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceTypeCount struct {
	// master机器数量

	Master *DeviceCount `json:"Master,omitempty" name:"Master"`
	// cell机器数量

	Cell *DeviceCount `json:"Cell,omitempty" name:"Cell"`
}

type TransferTaskOverview struct {
	// 迁移dst端的并发度

	DstCellConcurrent *uint64 `json:"DstCellConcurrent,omitempty" name:"DstCellConcurrent"`
	// 迁移dst端单组cp并发度控制开关 , 0关,1开

	DstCpConcurrentSw *uint64 `json:"DstCpConcurrentSw,omitempty" name:"DstCpConcurrentSw"`
	// 迁移dst端单组cp并发度

	DstCpMaxConcurrent *uint64 `json:"DstCpMaxConcurrent,omitempty" name:"DstCpMaxConcurrent"`
	// 迁移完成比例

	FinishTrsfRatio *float64 `json:"FinishTrsfRatio,omitempty" name:"FinishTrsfRatio"`
	// mov dst端cell预留容量百分比

	MovDstRemainCapPercent *float64 `json:"MovDstRemainCapPercent,omitempty" name:"MovDstRemainCapPercent"`
	// 迁移开关, 0关,1开

	MoveSw *uint64 `json:"MoveSw,omitempty" name:"MoveSw"`
	// 恢复开关, 0关,1开

	RecoverySw *uint64 `json:"RecoverySw,omitempty" name:"RecoverySw"`
	// recovery dst cell预留容量百分比

	RecvDstRemainCapPercent *float64 `json:"RecvDstRemainCapPercent,omitempty" name:"RecvDstRemainCapPercent"`
	// 迁移src端的并发度

	SrcCellConcurrent *uint64 `json:"SrcCellConcurrent,omitempty" name:"SrcCellConcurrent"`
	// 取消迁移任务超时时间

	TaskCancelTimeout *uint64 `json:"TaskCancelTimeout,omitempty" name:"TaskCancelTimeout"`
	// 任务超时时间

	TaskTimeout *uint64 `json:"TaskTimeout,omitempty" name:"TaskTimeout"`
	// 总的要迁移的tp数目

	TotalTrsfTp *uint64 `json:"TotalTrsfTp,omitempty" name:"TotalTrsfTp"`
	// 迁移失败的tp数目

	TrsfFailTpNum *uint64 `json:"TrsfFailTpNum,omitempty" name:"TrsfFailTpNum"`
	// 迁移总开关,含自动和手动迁移, 0关,1开

	TrsfMainSw *uint64 `json:"TrsfMainSw,omitempty" name:"TrsfMainSw"`
	// 迁移等待的tp数目

	WaitTrsfTp *uint64 `json:"WaitTrsfTp,omitempty" name:"WaitTrsfTp"`
	// 底层cbs zoneId

	CbsZoneId *uint64 `json:"CbsZoneId,omitempty" name:"CbsZoneId"`
	// 存储池ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// zone id

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type CreateDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建的云硬盘ID列表。

		DiskIdSet []*string `json:"DiskIdSet,omitempty" name:"DiskIdSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskZKClustersRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>zk-cluster-id - Array of Integer - 是否必填：否 -（过滤条件）按ZK集群ID过滤。<br><li>zk-type - Array of String - 是否必填：否 -（过滤条件）按ZK集群类型过滤。（master: 存储池zk | snapshot: 快照zk）。<br><li>zone - Array of String - 是否必填：否 -（过滤条件）按ZK集群所有zone过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDiskZKClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskZKClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest

	// 要解绑定期快照策略的云盘ID列表。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 要解绑的定期快照策略ID。

	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
}

func (r *UnbindAutoSnapshotPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DepotOverviewInAZ struct {
	// 可用区的ZoneId；

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区下的仓库概览信息；

	DepotOverviewInZone *DepotOverview `json:"DepotOverviewInZone,omitempty" name:"DepotOverviewInZone"`
}

type DiskStateInfo struct {
	// 存储池云盘总数

	DiskAllNum *uint64 `json:"DiskAllNum,omitempty" name:"DiskAllNum"`
	// 存储池已挂载云盘数量

	DiskAttachedNum *uint64 `json:"DiskAttachedNum,omitempty" name:"DiskAttachedNum"`
	// 存储池挂载中云盘数量

	DiskAttachingNum *uint64 `json:"DiskAttachingNum,omitempty" name:"DiskAttachingNum"`
	// 存储池已删除云盘数量

	DiskDeletedNum *uint64 `json:"DiskDeletedNum,omitempty" name:"DiskDeletedNum"`
	// 存储池已解挂云盘数量

	DiskDettachedNum *uint64 `json:"DiskDettachedNum,omitempty" name:"DiskDettachedNum"`
	// 存储池解挂中云盘数量

	DiskDettachingNum *uint64 `json:"DiskDettachingNum,omitempty" name:"DiskDettachingNum"`
	// 存储池扩容中云盘数量

	DiskEditingNum *uint64 `json:"DiskEditingNum,omitempty" name:"DiskEditingNum"`
	// 存储池已回收云盘数量

	DiskRecycledNum *uint64 `json:"DiskRecycledNum,omitempty" name:"DiskRecycledNum"`
	// 存储池回收中云盘数量

	DiskRecyclingNum *uint64 `json:"DiskRecyclingNum,omitempty" name:"DiskRecyclingNum"`
	// 存储池迁移中云盘数量

	DiskTrsfingNum *uint64 `json:"DiskTrsfingNum,omitempty" name:"DiskTrsfingNum"`
}

type Image struct {
	// 镜像实例ID。

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像名称。

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
}

type DescribeDisksMonitorDataRequest struct {
	*tchttp.BaseRequest

	// 取固定值`CBS`

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 监控指标，"await", "rcnt", "wcnt", "rband", "wband", "util", "svctm"

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 监控统计周期，只能为`60`

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 起始时间，如2018-09-22T19:51:23+08:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，默认为当前时间。 EndTime不能小于EtartTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 实例列表

	Instances []*MonitorInstance `json:"Instances,omitempty" name:"Instances"`
}

func (r *DescribeDisksMonitorDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisksMonitorDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoSnapshotPolicyAttributeRequest struct {
	*tchttp.BaseRequest

	// 定期快照策略ID。

	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
	// 定期快照的执行策略。

	Policy []*Policy `json:"Policy,omitempty" name:"Policy"`
	// 要创建的定期快照策略名。不传则默认为“未命名”。最大长度不能超60个字节。

	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" name:"AutoSnapshotPolicyName"`
	// 是否激活定期快照策略，FALSE表示未激活，TRUE表示激活，默认为TRUE。

	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`
	// 通过该定期快照策略创建的快照是否永久保留。FALSE表示非永久保留，TRUE表示永久保留，默认为FALSE。

	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`
	// 通过该定期快照策略创建的快照保留天数，该参数不可与`IsPermanent`参数冲突，即若定期快照策略设置为永久保留，`RetentionDays`应置0。

	RetentionDays *uint64 `json:"RetentionDays,omitempty" name:"RetentionDays"`
}

func (r *ModifyAutoSnapshotPolicyAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoSnapshotPolicyAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZKClusterOverview struct {
	// 快照集群数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 快照集群状态

	ZKClusterState *string `json:"ZKClusterState,omitempty" name:"ZKClusterState"`
}

type DescribeCustomerCapacityTopChangesRequest struct {
	*tchttp.BaseRequest

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 开始日期

	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`
	// 截止日期

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
	// 取值： Increase Decrease  （创建量，退量）

	StatMode *string `json:"StatMode,omitempty" name:"StatMode"`
	// 类型，all, inner, outer, ''

	UserType *string `json:"UserType,omitempty" name:"UserType"`
	// 返回个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCustomerCapacityTopChangesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerCapacityTopChangesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskConfigQuotaRequest struct {
	*tchttp.BaseRequest

	// 查询类别，取值范围。<br><li>INQUIRY_CBS_CONFIG：查询云盘配置列表<br><li>INQUIRY_CVM_CONFIG：查询云盘与实例搭配的配置列表。

	InquiryType *string `json:"InquiryType,omitempty" name:"InquiryType"`
	// 查询一个或多个[可用区](/document/api/213/9452#zone)下的配置。

	Zones []*string `json:"Zones,omitempty" name:"Zones"`
	// 付费模式。取值范围：<br><li>PREPAID：预付费<br><li>POSTPAID_BY_HOUR：后付费。

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 硬盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬盘<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：表示SSD云硬盘。

	DiskTypes []*string `json:"DiskTypes,omitempty" name:"DiskTypes"`
	// 系统盘或数据盘。取值范围：<br><li>SYSTEM_DISK：表示系统盘<br><li>DATA_DISK：表示数据盘。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 实例机型。

	DeviceClasses []*string `json:"DeviceClasses,omitempty" name:"DeviceClasses"`
	// 按照实例机型系列过滤。实例机型系列形如：S1、I1、M1等。详见[实例类型](https://cloud.tencent.com/document/product/213/11518)

	InstanceFamilies []*string `json:"InstanceFamilies,omitempty" name:"InstanceFamilies"`
	// 实例CPU核数。

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 实例内存大小。

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// INQUIRY_RESIZE、INQUIRY_CREATE

	InnerInquiryType *string `json:"InnerInquiryType,omitempty" name:"InnerInquiryType"`
}

func (r *DescribeDiskConfigQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskConfigQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskStorageDepotDetails struct {
	// 存储池CELL节点的机型

	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
	// master组件的设备列表

	Masters []*DiskDepotComponentDetail `json:"Masters,omitempty" name:"Masters"`
	// loadtrsf组件的设备列表

	LoadCtrls []*DiskDepotComponentDetail `json:"LoadCtrls,omitempty" name:"LoadCtrls"`
	// CELL节点详情

	CellGroups []*CellGroup `json:"CellGroups,omitempty" name:"CellGroups"`
	// 存储池总小表数

	TotalTpCount *uint64 `json:"TotalTpCount,omitempty" name:"TotalTpCount"`
	// 存储池IO性能

	FlowControlConfig *FlowControlConfig `json:"FlowControlConfig,omitempty" name:"FlowControlConfig"`
	// 存储池版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// common组件详情

	CommonModules []*CommonModule `json:"CommonModules,omitempty" name:"CommonModules"`
	// EC存储池的EC副本信息。

	EcDiskPairs []*DiskPair `json:"EcDiskPairs,omitempty" name:"EcDiskPairs"`
}

type ModifyDiskConfigForSaleRequest struct {
	*tchttp.BaseRequest

	// 云硬盘售卖配置

	DiskModelConfigForSale []*DiskConfigForSale `json:"DiskModelConfigForSale,omitempty" name:"DiskModelConfigForSale"`
}

func (r *ModifyDiskConfigForSaleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiskConfigForSaleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Dimension struct {
	// 监控指标名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 监控指标值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeDiskConfigForSaleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的云硬盘配置数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 云盘规格详情

		DiskConfigForSaleSet []*DiskConfigForSale `json:"DiskConfigForSaleSet,omitempty" name:"DiskConfigForSaleSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskConfigForSaleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskConfigForSaleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiskConfigForSaleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDiskConfigForSaleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiskConfigForSaleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeCommonComponentsRequest struct {
	*tchttp.BaseRequest

	// 待安装组件所属的集群类型，取值范围：<br><li>product-cbs-depot: 存储池集群<br><li>product-cbs-snap: 快照集群

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 待安装组件所属的集群ID

	ClusterId *uint64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 待安装common组件的节点属于集群的组件类型，取值范围: master, cell, manager, scheduler, transfer, cap_server, trsf_ctrl, trsf_proxy

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
	// 待安装common组件的节点ip列表

	DeviceIps []*string `json:"DeviceIps,omitempty" name:"DeviceIps"`
	// 需要安装的common组件列表，取值范围: send_alarm, collect_agent, ops_monitor, filebeat

	ModuleNames []*string `json:"ModuleNames,omitempty" name:"ModuleNames"`
}

func (r *UpgradeCommonComponentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeCommonComponentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CellSideSetIOStat struct {
	// 存储池集群uuid

	SetUuid *string `json:"SetUuid,omitempty" name:"SetUuid"`
	// 存储池io util，单位百分比

	Util []*float64 `json:"Util,omitempty" name:"Util"`
	// 存储池io svctm，单位毫秒

	Svctm []*float64 `json:"Svctm,omitempty" name:"Svctm"`
	// 存储池平均写io次数，单位个/秒

	Wcnt []*float64 `json:"Wcnt,omitempty" name:"Wcnt"`
	// 存储池平均写带宽，单位MB/秒

	Wband []*float64 `json:"Wband,omitempty" name:"Wband"`
	// 存储池io await，单位毫秒

	Await []*float64 `json:"Await,omitempty" name:"Await"`
	// 存储池平均读带宽，单位MB/秒

	Rband []*float64 `json:"Rband,omitempty" name:"Rband"`
	// 存储池平均读io次数，单位个/秒

	Rcnt []*float64 `json:"Rcnt,omitempty" name:"Rcnt"`
	// 网卡0入流量

	Eth0InkbpsTotal []*float64 `json:"Eth0InkbpsTotal,omitempty" name:"Eth0InkbpsTotal"`
	// 网卡0出流量

	Eth0OutkbpsTotal []*float64 `json:"Eth0OutkbpsTotal,omitempty" name:"Eth0OutkbpsTotal"`
	// 网卡1入流量

	Eth1InkbpsTotal *float64 `json:"Eth1InkbpsTotal,omitempty" name:"Eth1InkbpsTotal"`
	// 网卡1出流量

	Eth1OutkbpsTotal []*float64 `json:"Eth1OutkbpsTotal,omitempty" name:"Eth1OutkbpsTotal"`
	// Cell进程CPU利用率最大值

	McdCpuutilMax []*float64 `json:"McdCpuutilMax,omitempty" name:"McdCpuutilMax"`
	// 单台cell剩余可用内存最小值

	MemAvailable []*float64 `json:"MemAvailable,omitempty" name:"MemAvailable"`
}

type DeleteDiskStoragePoolGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDiskStoragePoolGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDiskStoragePoolGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskStorageDepotAttributesRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 是否展示关联仓库信息，默认否

	ShowDepotUsage *bool `json:"ShowDepotUsage,omitempty" name:"ShowDepotUsage"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDiskStorageDepotAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskStorageDepotAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReplaceDiskTasksRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>job-id - Array of Integer - 是否必填：否 -（过滤条件）按任务id过滤。<br><li>job-status - Array of String - 是否必填：否 -（过滤条件）按照任务状态过滤。取值范围：<br><li>running：运行中<br><li>pause：异常<br><li>normal_finish：任务完成<br><li>wait_check：等待确认<br><li>stop：任务取消

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeReplaceDiskTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReplaceDiskTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitializeDiskStorageDepotRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储集群ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
}

func (r *InitializeDiskStorageDepotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitializeDiskStorageDepotRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReviveDiskStorageDepotCellResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReviveDiskStorageDepotCellResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReviveDiskStorageDepotCellResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDepotConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询的配置项内容。

		DepotConfigSet []*DepotConfigItem `json:"DepotConfigSet,omitempty" name:"DepotConfigSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDepotConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepotConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallDiskZKClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务ID

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InstallDiskZKClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallDiskZKClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskFeature struct {
	// 云硬盘使用的计费系统。取值范围：<br><li>bmppro：计费系统<br><li>bmp：计量系统

	TradeType *string `json:"TradeType,omitempty" name:"TradeType"`
}

type DiskOrder struct {
	// 产品类别。

	GoodsCategoryId *uint64 `json:"GoodsCategoryId,omitempty" name:"GoodsCategoryId"`
	// 产品数量。

	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 项目ID。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 产品详情。

	GoodsDetail *GoodsDetail `json:"GoodsDetail,omitempty" name:"GoodsDetail"`
	// 地域ID。

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 可用区ID。

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 付费模式。

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 云盘类别。

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DeleteDiskStorageDepotRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储集群ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
}

func (r *DeleteDiskStorageDepotRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDiskStorageDepotRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDepotCellDeviceConfigsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。<br><li>device-type:按机型过滤<br><li>depot-media:按存储池类型过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDepotCellDeviceConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepotCellDeviceConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDisksRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个云硬盘ID查询。云硬盘ID形如：`disk-11112222`，此参数的具体格式可参考API[简介](/document/product/362/15633)的ids.N一节）。参数不支持同时指定`DiskIds`和`Filters`。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 过滤条件。参数不支持同时指定`DiskIds`和`Filters`。<br><li>disk-usage - Array of String - 是否必填：否 -（过滤条件）按云盘类型过滤。 (SYSTEM_DISK：表示系统盘 | DATA_DISK：表示数据盘)<br><li>disk-charge-type - Array of String - 是否必填：否 -（过滤条件）按照云硬盘计费模式过滤。 (PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费。)<br><li>portable - Array of Bool - 是否必填：否 -（过滤条件）按是否为弹性云盘过滤。 (TRUE：表示弹性云盘 | FALSE：表示非弹性云盘。)<br><li>project-id - Array of Integer - 是否必填：否 -（过滤条件）按云硬盘所属项目ID过滤。<br><li>disk-id - Array of String - 是否必填：否 -（过滤条件）按照云硬盘ID过滤。云盘ID形如：`disk-11112222`。<br><li>disk-name - Array of String - 是否必填：否 -（过滤条件）按照云盘名称过滤。<br><li>disk-type - Array of String - 是否必填：否 -（过滤条件）按照云盘介质类型过滤。(CLOUD_BASIC：表示普通云硬盘 | CLOUD_PREMIUM：表示高性能云硬盘。| CLOUD_SSD：SSD表示SSD云硬盘。)<br><li>disk-state - Array of String - 是否必填：否 -（过滤条件）按照云盘状态过滤。(UNATTACHED：未挂载 | ATTACHING：挂载中 | ATTACHED：已挂载 | DETACHING：解挂中 | EXPANDING：扩容中 | ROLLBACKING：回滚中 | TORECYCLE：待回收。)<br><li>instance-id - Array of String - 是否必填：否 -（过滤条件）按照云盘挂载的云主机实例ID过滤。可根据此参数查询挂载在指定云主机下的云硬盘。<br><li>zone - Array of String - 是否必填：否 -（过滤条件）按照[可用区](/document/api/213/9452#zone)过滤。<br><li>instance-ip-address - Array of String - 是否必填：否 -（过滤条件）按云盘所挂载云主机的内网或外网IP过滤。<br><li>instance-name - Array of String - 是否必填：否 -（过滤条件）按云盘所挂载的实例名称过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 输出云盘列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 云盘列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据云盘的创建时间排序<br><li>DEADLINE：依据云盘的到期时间排序<br>默认按云盘创建时间排序。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 云盘详情中是否需要返回云盘绑定的定期快照策略ID，TRUE表示需要返回，FALSE表示不返回。

	ReturnBindAutoSnapshotPolicy *bool `json:"ReturnBindAutoSnapshotPolicy,omitempty" name:"ReturnBindAutoSnapshotPolicy"`
	// 内部参数，用于支持搜索框搜索。

	InnerSearch *string `json:"InnerSearch,omitempty" name:"InnerSearch"`
}

func (r *DescribeDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiskStorageDepotAttributesRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 云硬盘存储仓库ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 存储仓库可创建的最大云盘数量

	DiskNumTotal *uint64 `json:"DiskNumTotal,omitempty" name:"DiskNumTotal"`
	// 存储仓库的超卖比（%）

	OverSold *uint64 `json:"OverSold,omitempty" name:"OverSold"`
	// 快照集群ID。**注意快照集群绑定后无法修改**

	SnapshotBoxId *uint64 `json:"SnapshotBoxId,omitempty" name:"SnapshotBoxId"`
	// 是否打开售卖。打开售卖时需要确保已完成存储Cell上架操作（NORMAL/DISABLE/PREPARE)

	DepotSoldState *string `json:"DepotSoldState,omitempty" name:"DepotSoldState"`
	// 仓库属性ID。

	DepotAttributeId *string `json:"DepotAttributeId,omitempty" name:"DepotAttributeId"`
	// 存储池的最大容量。

	StorageDepotSize *uint64 `json:"StorageDepotSize,omitempty" name:"StorageDepotSize"`
	// 切换存储池无数据管理节点的主备，传入1表示进行主备切换；传入SwitchMaster参数时，不支持同时修改存储池其他属性。

	SwitchMaster *uint64 `json:"SwitchMaster,omitempty" name:"SwitchMaster"`
}

func (r *ModifyDiskStorageDepotAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiskStorageDepotAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunClusterUpgradeTasksRequest struct {
	*tchttp.BaseRequest

	// 升级任务ID列表。当前仅支持传入一个升级任务，不支持批量

	TaskIds []*uint64 `json:"TaskIds,omitempty" name:"TaskIds"`
	// 是否执行实际升级操作。取值为0，表示需要实际发起升级操作；取值为1，表示仅执行环境检查，检查成功后，任务状态会变为CHECK_FINISH

	DryRun *uint64 `json:"DryRun,omitempty" name:"DryRun"`
	// 集群升级的事务ID，通过该参数可以查到一次集群升级的所有任务集合。执行cell单边并发升级时要传入该参数，同时传入CellReplicateIndex参数，表示指定升级cell的某一副本。传入本参数时，不能再传入TaskIds。

	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`
	// 执行cell单边并发升级时要传入该参数，与TransactionId一起来指定升级当前存储池的某一副本。

	CellReplicateIndex *uint64 `json:"CellReplicateIndex,omitempty" name:"CellReplicateIndex"`
}

func (r *RunClusterUpgradeTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunClusterUpgradeTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlowControlConfig struct {
	// 设置云盘流控的动作。<br><li>0: 加入黑名单<br><li>1: 解除黑名单<br><li>2: 解除软白名单<br><li>3:加入软白名单

	AutoAdjust *uint64 `json:"AutoAdjust,omitempty" name:"AutoAdjust"`
	// 每秒读带宽限制，最低512KB/s

	ReadThroughLimit *uint64 `json:"ReadThroughLimit,omitempty" name:"ReadThroughLimit"`
	// 每秒写带宽限制，最低512KB/s

	WriteThroughLimit *uint64 `json:"WriteThroughLimit,omitempty" name:"WriteThroughLimit"`
	// 每秒读次数限制，最低2次/秒

	ReadIopsLimit *uint64 `json:"ReadIopsLimit,omitempty" name:"ReadIopsLimit"`
	// 每秒写次数限制，最低2次/秒

	WriteIopsLimit *uint64 `json:"WriteIopsLimit,omitempty" name:"WriteIopsLimit"`
}

type DiskTpStateInfo struct {
	// 两副本dead小表数量

	TwoDeadNum *uint64 `json:"TwoDeadNum,omitempty" name:"TwoDeadNum"`
	// born小表数量

	TpBornNum *uint64 `json:"TpBornNum,omitempty" name:"TpBornNum"`
	// normal小表数量

	TpNormalNum *uint64 `json:"TpNormalNum,omitempty" name:"TpNormalNum"`
	// free小表数量

	TpFreeNum *uint64 `json:"TpFreeNum,omitempty" name:"TpFreeNum"`
	// errfree小表数量

	TpErrfreeNum *uint64 `json:"TpErrfreeNum,omitempty" name:"TpErrfreeNum"`
	// errborn小表数量

	TpErrbornNum *uint64 `json:"TpErrbornNum,omitempty" name:"TpErrbornNum"`
	// 一份dead小表数量

	OneDeadNum *uint64 `json:"OneDeadNum,omitempty" name:"OneDeadNum"`
	// 总小表数

	TpAllNum *uint64 `json:"TpAllNum,omitempty" name:"TpAllNum"`
	// abnormal小表数量

	TpAbnormalNum *uint64 `json:"TpAbnormalNum,omitempty" name:"TpAbnormalNum"`
	// 已使用小表数量

	TpUsedNum *uint64 `json:"TpUsedNum,omitempty" name:"TpUsedNum"`
	// 3副本dead小表数，只在ec副本中有意义。

	ThreeDeadNum *uint64 `json:"ThreeDeadNum,omitempty" name:"ThreeDeadNum"`
}

type BindAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest

	// 要绑定的定期快照策略ID。

	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
	// 要绑定的云硬盘ID列表，一次请求最多绑定80块云盘。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

func (r *BindAutoSnapshotPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的云硬盘数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 云硬盘的详细信息列表。

		DiskSet []*Disk `json:"DiskSet,omitempty" name:"DiskSet"`
		// 偏移量

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 返回数量

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceCount struct {
	// 设备数量

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type CancelDepotTransferTasksRequest struct {
	*tchttp.BaseRequest

	// 任务ID列表

	TaskIds []*uint64 `json:"TaskIds,omitempty" name:"TaskIds"`
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储池ID，要取消整个存储池的迁移任务时传入

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
}

func (r *CancelDepotTransferTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelDepotTransferTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDepotCellDeviceConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新增的存储机型配置详情

		DepotCellDeviceConfig *DepotCellDeviceConfig `json:"DepotCellDeviceConfig,omitempty" name:"DepotCellDeviceConfig"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDepotCellDeviceConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDepotCellDeviceConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskZKClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的ZK集群总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// ZK集群说情

		DiskZkClusterSet []*DiskZkCluster `json:"DiskZkClusterSet,omitempty" name:"DiskZkClusterSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskZKClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskZKClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopDiskStorageDepotCellRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储集群ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 要停止的CELL节点IP列表

	DeviceIps []*string `json:"DeviceIps,omitempty" name:"DeviceIps"`
	// 正常停进程时，如果有在使用的小表，则不能停进程；强制停进程时，在裁撤仓库时使用，则忽略小表状态

	ForceUninstall *uint64 `json:"ForceUninstall,omitempty" name:"ForceUninstall"`
}

func (r *StopDiskStorageDepotCellRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopDiskStorageDepotCellRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDisksRequest struct {
	*tchttp.BaseRequest

	// 硬盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬盘<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：表示SSD云硬盘。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云盘显示名称。不传则默认为“未命名”。最大长度不能超60个字节。

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 创建云硬盘数量，不传则默认为1。单次请求最多可创建的云盘数有限制，具体参见[云硬盘使用限制](https://cloud.tencent.com/doc/product/362/5145)。

	DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`
	// 云硬盘计费类型。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br>目前只支持后付费模式。

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月云盘的购买时长、是否设置自动续费等属性，创建预付费云盘该参数必传。

	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitempty" name:"DiskChargePrepaid"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目。若不指定项目，将在默认项目下进行创建。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 云硬盘大小，单位为GB。<br><li>如果传入`SnapshotId`则可不传`DiskSize`，此时新建云盘的大小为快照大小<br><li>如果传入`SnapshotId`同时传入`DiskSize`，则云盘大小必须大于或等于快照大小<br><li>云盘大小取值范围： 普通云硬盘:10GB ~ 4000G；高性能云硬盘:50GB ~ 4000GB；SSD云硬盘:100GB ~ 4000GB。步长均为10GB

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 快照ID，如果传入则根据此快照创建云硬盘，快照类型必须为数据盘快照，可通过[DescribeSnapshots](/document/product/362/15647)接口查询快照，见输出参数DiskUsage解释。

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。

	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
	// 传入该参数用于创建加密云盘，取值固定为ENCRYPT。

	Encrypt *string `json:"Encrypt,omitempty" name:"Encrypt"`
	// 云盘绑定的标签列表。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 从运营端发起云盘创建时，本参数必传，取值为TRUE。表示从运营端直接发起云盘创建。

	InternalDirectCreate *bool `json:"InternalDirectCreate,omitempty" name:"InternalDirectCreate"`
	// 可选参数。使用此参数可给云硬盘购买额外的性能。<br>当前仅支持极速型云盘（CLOUD_TSSD）和增强型SSD云硬盘（CLOUD_HSSD）

	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
}

func (r *CreateDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskStorageDeviceTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存储设备的配置信息列表

		DiskStorageDeviceConfigSet []*DiskStorageDeviceConfig `json:"DiskStorageDeviceConfigSet,omitempty" name:"DiskStorageDeviceConfigSet"`
		// 支持的存储设备数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskStorageDeviceTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskStorageDeviceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleRule struct {
	// 回收规则设置的对象，depot: 存储池；disk: 云盘

	RecycleType *string `json:"RecycleType,omitempty" name:"RecycleType"`
	// 回收间隔，单位秒

	RecycleDiff *uint64 `json:"RecycleDiff,omitempty" name:"RecycleDiff"`
	// 所在可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储池池，仅在RecycleType为depot时有效

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 云盘uuid，仅在RecycleType为disk时有效

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
	// 回收规则最后修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 回收规则添加时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 回收规则创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
}

type AttachDisksRequest struct {
	*tchttp.BaseRequest

	// 将要被挂载的弹性云盘ID。通过[DescribeDisks](/document/product/362/16315)接口查询。单次最多可挂载10块弹性云盘。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 云服务器实例ID。云盘将被挂载到此云服务器上，通过[DescribeInstances](/document/product/213/15728)接口查询。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 可选参数，不传该参数则仅执行挂载操作。<br>传入该参数，会在挂载时将云盘的生命周期与待挂载主机对齐。取值范围：<br><li>AUTO_RENEW：云盘未设置自动续费时，可传该值，将云盘设置为自动续费。<br><li> DEADLINE_ALIGN：当云盘的到期时间早于待挂载主机，可传该值，将云盘的到期时间与主机对齐。

	AlignType *string `json:"AlignType,omitempty" name:"AlignType"`
	// 云盘挂载对象的类别。取值范围：<br><li>CVM: 租户端CVM<br><li>RAW_CVM：运营端创建的CVM<br><li>HOST: 直接挂载到宿主机上

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 可选参数，不传该参数则仅执行挂载操作。传入`True`时，会在挂载成功后将云硬盘设置为随云主机销毁模式，仅对按量计费云硬盘有效。

	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
	// 可选参数，用于控制台批量挂载共享盘到多个CVM时指定实例ID。如果传入多个InstanceId，那么DisksId中指定的云盘必须全部为共享云盘，否则返回错误。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *AttachDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDisksMonitorDataOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IO 性能详情

		MonitorDataSet []*MonitorData `json:"MonitorDataSet,omitempty" name:"MonitorDataSet"`
		// 符合条件的记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDisksMonitorDataOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisksMonitorDataOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotBoxPerformanceDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件指标监控数据详情

		ComponentPerformanceData *ComponentPerformanceData `json:"ComponentPerformanceData,omitempty" name:"ComponentPerformanceData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotBoxPerformanceDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotBoxPerformanceDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterConfigDetail struct {
	// 配置名称

	ConfigItem *string `json:"ConfigItem,omitempty" name:"ConfigItem"`
	// 配置所在的组件

	ConfigModule *string `json:"ConfigModule,omitempty" name:"ConfigModule"`
	// 配置所在的配置文件名称

	ConfigFile *string `json:"ConfigFile,omitempty" name:"ConfigFile"`
	// 配置在配置文件中的位置

	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`
	// 配置值

	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`
	// 配置允许的最小值

	MinValue *string `json:"MinValue,omitempty" name:"MinValue"`
	// 配置允许的最大值

	MaxValue *string `json:"MaxValue,omitempty" name:"MaxValue"`
	// 配置的单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 配置是从哪些机器上读取的

	DeviceIps []*string `json:"DeviceIps,omitempty" name:"DeviceIps"`
}

type DiskDeniedAction struct {
	// 云硬盘ID，比如"disk-7rf24jvb"

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 具体的云硬盘的操作掩码信息

	DeniedActions []*DeniedAction `json:"DeniedActions,omitempty" name:"DeniedActions"`
}

type BindAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindAutoSnapshotPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDepotFlowControlRecordsRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储仓库ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDepotFlowControlRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDepotFlowControlRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskResourceUsageRequest struct {
	*tchttp.BaseRequest

	// 支持"disk-type"、“zone-id"过滤，只能传入一个值

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 支持按DAY/WEEK/MONTH统计；默认按天统计

	Period *string `json:"Period,omitempty" name:"Period"`
	// 按云硬盘（DISK）或者快照（SNAP）统计

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 按总量（TOTAL）或者净增量（INCREASE）统计

	StatisticType *string `json:"StatisticType,omitempty" name:"StatisticType"`
	// 统计起始时间，默认31天前

	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`
	// 截止日期

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeDiskResourceUsageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskResourceUsageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManageRecycleRulesRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储池ID列表，在调用时该参数与DiskUuids必传且只能传一个。

	DepotIds []*uint64 `json:"DepotIds,omitempty" name:"DepotIds"`
	// 云盘uuid列表，在调用时该参数与DepotIds必传且只能传一个。

	DiskUuids []*string `json:"DiskUuids,omitempty" name:"DiskUuids"`
	// 操作类型。<br><li>ADD: 添加回收规则<br><li>EDIT: 修改回收规则<br><li>DELETE:删除回收规则

	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`
	// 回收间隔，单位秒

	RecycleDiff *uint64 `json:"RecycleDiff,omitempty" name:"RecycleDiff"`
}

func (r *ManageRecycleRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ManageRecycleRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallDiskStorageDepotMasterRequest struct {
	*tchttp.BaseRequest

	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 存储集群ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// Master节点IP，至少需要一台Master正常工作。

	DeviceIp *string `json:"DeviceIp,omitempty" name:"DeviceIp"`
	// 是否强制卸载。0：不强制卸载；1：强制卸载

	ForceUninstall *uint64 `json:"ForceUninstall,omitempty" name:"ForceUninstall"`
}

func (r *UninstallDiskStorageDepotMasterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallDiskStorageDepotMasterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
