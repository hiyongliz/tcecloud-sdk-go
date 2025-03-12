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

package v20200509

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DescribeIdcExportLineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 专线出口列表

		IdcExportLineResultSet []*IdcExportLineResult `json:"IdcExportLineResultSet,omitempty" name:"IdcExportLineResultSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdcExportLineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcExportLineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrapAlarmRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// <li>Title - String数组  - （过滤条件）TRAP告警主题，形如：`IF::linkupdown`。</li>
	// <li>DeviceIp - String数组  - （过滤条件）TRAP告警源设备IP，形如：`10.21.128.252`。</li>
	// <li>StartTime - String数组  - （过滤条件）TRAP告警开始时间，形如：`2020-07-14 00:00:00`。</li>
	// <li>EndTime - String数组  - （过滤条件）TRAP告警结束时间，形如：`2020-07-15 00:00:00`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTrapAlarmRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrapAlarmRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDetailPingDataRequest struct {
	*tchttp.BaseRequest

	// 探测类型：gw_ping, lan_ping, wan_ping

	PingType *string `json:"PingType,omitempty" name:"PingType"`
	// 查询目标

	TargetList []*string `json:"TargetList,omitempty" name:"TargetList"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *GetDetailPingDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDetailPingDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTrapAlarmRequest struct {
	*tchttp.BaseRequest

	// SNMP TRAP告警主题

	Title *string `json:"Title,omitempty" name:"Title"`
}

func (r *DeleteTrapAlarmRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTrapAlarmRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDeviceConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络设备配置信息

		NetDeviceConfigs *NetDeviceConfigSet `json:"NetDeviceConfigs,omitempty" name:"NetDeviceConfigs"`
		// 数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetDeviceConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDeviceConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WanStatus struct {
	// Region中文名

	RegionCN *string `json:"RegionCN,omitempty" name:"RegionCN"`
	// Region英文名

	Region *string `json:"Region,omitempty" name:"Region"`
	// 可用区中文名

	AzCN *string `json:"AzCN,omitempty" name:"AzCN"`
	// 可用区英文名

	Az *string `json:"Az,omitempty" name:"Az"`
	// 出口名

	IdcExitName *string `json:"IdcExitName,omitempty" name:"IdcExitName"`
	// 出口运营商

	IdcExitIsp *string `json:"IdcExitIsp,omitempty" name:"IdcExitIsp"`
	// 探测机器IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 当前状态

	CurrentStatus *int64 `json:"CurrentStatus,omitempty" name:"CurrentStatus"`
	// 历史状态

	PastStatus *int64 `json:"PastStatus,omitempty" name:"PastStatus"`
}

type DescribeIdcOverviewRequest struct {
	*tchttp.BaseRequest

	// 返回记录最大限制，Limit和Offset同时传入时生效

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 返回记录相对偏移，Limit和Offset同时传入时生效

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// idc名称

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 可用区

	Az *string `json:"Az,omitempty" name:"Az"`
	// 查询过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 查询类型

	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`
}

func (r *DescribeIdcOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMemoryUsageTop5Request struct {
	*tchttp.BaseRequest

	// 根据过滤条件 "Filters":[
	//         {
	//             "Name":"Idc",
	//             "Values":["YF-1"]
	//         }
	//     ]返回指定idc内存利用率top5,如果Filters为空表示返回整个region 内存利用率top5。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeMemoryUsageTop5Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMemoryUsageTop5Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpecialIdcLineMetricRequest struct {
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
	// 可用区名字

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// IDC专线名字

	SpeLineName *string `json:"SpeLineName,omitempty" name:"SpeLineName"`
}

func (r *DescribeSpecialIdcLineMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpecialIdcLineMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRackExtendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRackExtendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRackExtendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetDeviceConfigForNmsUpgradeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetDeviceConfigForNmsUpgradeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetDeviceConfigForNmsUpgradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIPCIDRSegmentsRequest struct {
	*tchttp.BaseRequest

	// IPCIDR网段信息

	IPCIDRSegmentSet []*IPCIDRSegment `json:"IPCIDRSegmentSet,omitempty" name:"IPCIDRSegmentSet"`
}

func (r *DeleteIPCIDRSegmentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIPCIDRSegmentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetServersTopo struct {
	// 本地网络设备系统名字

	LocalName *string `json:"LocalName,omitempty" name:"LocalName"`
	// 本地网络设备ip地址

	LocalIp *string `json:"LocalIp,omitempty" name:"LocalIp"`
	// 本地网络设备拓扑层次角色名字

	LocalRole *string `json:"LocalRole,omitempty" name:"LocalRole"`
	// 本地网络设备所属地域名字

	LocalRegion *string `json:"LocalRegion,omitempty" name:"LocalRegion"`
	// 本地网络设备所属AZ名字

	LocalZone *string `json:"LocalZone,omitempty" name:"LocalZone"`
	// 本地网络设备互联端口名字

	LocalPort *string `json:"LocalPort,omitempty" name:"LocalPort"`
	// 远端服务器名字

	RemoteName *string `json:"RemoteName,omitempty" name:"RemoteName"`
	// 远端服务器资产编号

	RemoteAssetId *string `json:"RemoteAssetId,omitempty" name:"RemoteAssetId"`
	// 远端服务器ip地址

	RemoteIp *string `json:"RemoteIp,omitempty" name:"RemoteIp"`
	// 远端服务器所属地域名字

	RemoteRegion *string `json:"RemoteRegion,omitempty" name:"RemoteRegion"`
	// 远端服务器所属AZ名字

	RemoteZone *string `json:"RemoteZone,omitempty" name:"RemoteZone"`
}

type UpdateNetDeviceSyslogKeywordRequest struct {
	*tchttp.BaseRequest

	// 网络设备系统日志关键字

	SyslogKeyword *SyslogKeywords `json:"SyslogKeyword,omitempty" name:"SyslogKeyword"`
}

func (r *UpdateNetDeviceSyslogKeywordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNetDeviceSyslogKeywordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateIPSegment struct {
	// 网段名称

	Segment *string `json:"Segment,omitempty" name:"Segment"`
	// 逻辑区域

	LogicalArea *string `json:"LogicalArea,omitempty" name:"LogicalArea"`
	// 网段用途

	SegmentUsage *string `json:"SegmentUsage,omitempty" name:"SegmentUsage"`
	// 网段所属产品

	SegmentProduct *string `json:"SegmentProduct,omitempty" name:"SegmentProduct"`
	// vlan id

	VlanId *uint64 `json:"VlanId,omitempty" name:"VlanId"`
	// 可用区

	Az *string `json:"Az,omitempty" name:"Az"`
	// IDC机房

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 交换机

	Switch *string `json:"Switch,omitempty" name:"Switch"`
	// 掩码

	Mask *string `json:"Mask,omitempty" name:"Mask"`
	// 网关地址

	Gateway *string `json:"Gateway,omitempty" name:"Gateway"`
	// 网络设备名称

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type GetWanConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 省份列表

		ProvinceList []*string `json:"ProvinceList,omitempty" name:"ProvinceList"`
		// 运营商列表

		IspList []*string `json:"IspList,omitempty" name:"IspList"`
		// 探测配置列表

		ProbeList []*ProbeList `json:"ProbeList,omitempty" name:"ProbeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetWanConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWanConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkDevicesRequest struct {
	*tchttp.BaseRequest

	// 设备固资编号

	AssetIDs []*string `json:"AssetIDs,omitempty" name:"AssetIDs"`
}

func (r *DeleteNetworkDevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkDevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpecialIdcLineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IDC专线对象数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// IDC专线对象数组

		SpecialIdcLineResultSet []*SpecialIdcLineResult `json:"SpecialIdcLineResultSet,omitempty" name:"SpecialIdcLineResultSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSpecialIdcLineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpecialIdcLineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Configs struct {
	// 网络设备资产信息

	NetDeviceAsset *string `json:"NetDeviceAsset,omitempty" name:"NetDeviceAsset"`
	// ConfigType类型："Baseline"/"Daily"/"Current"，只有"Daily"类型时需要传递ConfigTime.

	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
	// 获取对应日期的配置

	ConfigTime *string `json:"ConfigTime,omitempty" name:"ConfigTime"`
	// 展示编码

	Encoding *string `json:"Encoding,omitempty" name:"Encoding"`
}

type CreateSpecialIdcLineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// idc专线信息

		SpecialIdcLineResult *SpecialIdcLineResult `json:"SpecialIdcLineResult,omitempty" name:"SpecialIdcLineResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSpecialIdcLineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSpecialIdcLineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDevicePortChannelListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络设备端口光模块通道列表对象数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 网络设备端口光模块通道列表对象数组

		NetDevicePortChannelListSet []*DeviceChannel `json:"NetDevicePortChannelListSet,omitempty" name:"NetDevicePortChannelListSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetDevicePortChannelListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDevicePortChannelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIPSegmentsRequest struct {
	*tchttp.BaseRequest

	// IP网段名称列表

	IPSegments []*DeleteIPSegment `json:"IPSegments,omitempty" name:"IPSegments"`
}

func (r *DeleteIPSegmentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIPSegmentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IdcExportLine struct {
	// ExportLineId

	ExportLineId *string `json:"ExportLineId,omitempty" name:"ExportLineId"`
	// ExportLineName

	ExportLineName *string `json:"ExportLineName,omitempty" name:"ExportLineName"`
	// ExportLineStatus

	ExportLineStatus *int64 `json:"ExportLineStatus,omitempty" name:"ExportLineStatus"`
	// LineTechAvailableSpeed

	LineTechAvailableSpeed *string `json:"LineTechAvailableSpeed,omitempty" name:"LineTechAvailableSpeed"`
	// LineBusiAvailableSpeed

	LineBusiAvailableSpeed *string `json:"LineBusiAvailableSpeed,omitempty" name:"LineBusiAvailableSpeed"`
	// ExportDeviceIp

	ExportDeviceIp *string `json:"ExportDeviceIp,omitempty" name:"ExportDeviceIp"`
	// RemoteDeviceIp

	RemoteDeviceIp *string `json:"RemoteDeviceIp,omitempty" name:"RemoteDeviceIp"`
	// LineOperator

	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`
	// RemoteOperatorOwner

	RemoteOperatorOwner *string `json:"RemoteOperatorOwner,omitempty" name:"RemoteOperatorOwner"`
	// RemoteOperatorOwnerPhone

	RemoteOperatorOwnerPhone *string `json:"RemoteOperatorOwnerPhone,omitempty" name:"RemoteOperatorOwnerPhone"`
	// RemoteOperatorOwnerWarningPhone

	RemoteOperatorOwnerWarningPhone *string `json:"RemoteOperatorOwnerWarningPhone,omitempty" name:"RemoteOperatorOwnerWarningPhone"`
	// SlaNum

	SlaNum *string `json:"SlaNum,omitempty" name:"SlaNum"`
	// Comment

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// ExportDevice

	ExportDevice *string `json:"ExportDevice,omitempty" name:"ExportDevice"`
	// ExportDevicePort

	ExportDevicePort []*string `json:"ExportDevicePort,omitempty" name:"ExportDevicePort"`
	// ExportDeviceAndPort

	ExportDeviceAndPort []*string `json:"ExportDeviceAndPort,omitempty" name:"ExportDeviceAndPort"`
}

type NetDevicePortDesc struct {
	// 网络设备ip

	NetdeviceIp *string `json:"NetdeviceIp,omitempty" name:"NetdeviceIp"`
	// 网络设备名字

	NetdeviceName *string `json:"NetdeviceName,omitempty" name:"NetdeviceName"`
	// 网络设备端口描述

	PortDesc *string `json:"PortDesc,omitempty" name:"PortDesc"`
	// 网络设备端口ip

	PortIp *string `json:"PortIp,omitempty" name:"PortIp"`
	// 网络设备端口mac

	PortMac *string `json:"PortMac,omitempty" name:"PortMac"`
	// 网络设备端口名字

	PortName *string `json:"PortName,omitempty" name:"PortName"`
	// 网络设备端口索引

	PortSnmpIndex *string `json:"PortSnmpIndex,omitempty" name:"PortSnmpIndex"`
	// 网络设备端口速率

	PortSpeed *string `json:"PortSpeed,omitempty" name:"PortSpeed"`
	// 网络设备端口管理状态

	PortStateAdmin *int64 `json:"PortStateAdmin,omitempty" name:"PortStateAdmin"`
	// 网络设备端口操作状态

	PortStateOper *int64 `json:"PortStateOper,omitempty" name:"PortStateOper"`
	// 网络设备端口更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 网络设备端口创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 网络设备端口通道

	PortChannel *string `json:"PortChannel,omitempty" name:"PortChannel"`
	// 网络设备端口序号

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type DescribeCPUUsageTop5Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// CPU利用率统计信息

		CPUUsageStatistics []*UsageStatistic `json:"CPUUsageStatistics,omitempty" name:"CPUUsageStatistics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCPUUsageTop5Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCPUUsageTop5Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRacksRequest struct {
	*tchttp.BaseRequest

	// 创建机架信息

	CreateRacks []*RackInfo `json:"CreateRacks,omitempty" name:"CreateRacks"`
}

func (r *CreateRacksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRacksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkPortsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回网络设备端口信息

		NetPortSet []*NetworkPort `json:"NetPortSet,omitempty" name:"NetPortSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkPortsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkPortsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateIPAddress struct {
	// IP地址

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// IP地址所有者

	OwnerName *string `json:"OwnerName,omitempty" name:"OwnerName"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 是否已经分配，取值0（未分配），1（已分配）

	IsAllocated *int64 `json:"IsAllocated,omitempty" name:"IsAllocated"`
	// IP类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 对象ID

	ObjectId *string `json:"ObjectId,omitempty" name:"ObjectId"`
}

type DeleteRackPositionsRequest struct {
	*tchttp.BaseRequest

	// 删除机位信息

	DeleteRackPositions []*RackPositionInfo `json:"DeleteRackPositions,omitempty" name:"DeleteRackPositions"`
}

func (r *DeleteRackPositionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRackPositionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceAlarmResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警设备列表

		AlarmDeviceList []*string `json:"AlarmDeviceList,omitempty" name:"AlarmDeviceList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceAlarmResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddGateway struct {
	// 网关集群名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 网关资产ID

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 内部IP

	InnerIp *string `json:"InnerIp,omitempty" name:"InnerIp"`
	// 外部IP

	OuterIp *string `json:"OuterIp,omitempty" name:"OuterIp"`
}

type RackPositionOverview struct {
	// region

	Region *string `json:"Region,omitempty" name:"Region"`
	// 可用区

	Az *string `json:"Az,omitempty" name:"Az"`
	// idc

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 机架编号

	RackName *string `json:"RackName,omitempty" name:"RackName"`
	// 机架类型

	RackType *string `json:"RackType,omitempty" name:"RackType"`
	// 机架通电状态

	PowerStatus *int64 `json:"PowerStatus,omitempty" name:"PowerStatus"`
	// 机架上的设备列表

	Devices []*RackPositionDevice `json:"Devices,omitempty" name:"Devices"`
	// 设备数量

	DevicesOnNum *int64 `json:"DevicesOnNum,omitempty" name:"DevicesOnNum"`
	// 空闲机位数

	IdlePosNum *int64 `json:"IdlePosNum,omitempty" name:"IdlePosNum"`
	// 机位数

	RackPosNum *int64 `json:"RackPosNum,omitempty" name:"RackPosNum"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeDeviceAlarmRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDeviceAlarmRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceAlarmRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDevicesTopoExRequest struct {
	*tchttp.BaseRequest

	// ZoneName

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

func (r *DescribeNetDevicesTopoExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDevicesTopoExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkPortsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetworkPortsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetworkPortsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDevicesTopoRequest struct {
	*tchttp.BaseRequest

	// 可用区名字

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为100，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// <li>LocalIp - String数组  - （过滤条件）本地网络设备ip地址，形如：`10.21.128.252`。</li>
	// <li>RemoteIp - String数组  - （过滤条件）远端网络设备ip地址，形如：`10.21.128.253`。</li>
	// <li>LocalPort - String数组  - （过滤条件）本地网络设备物理连接端口，形如：`HundredGigE1/1/1`。</li>
	// <li>RemotePort - String数组  - （过滤条件）远端网络设备物理连接端口，形如：`FortyGigE1/0/49`。</li>
	// <li>LocalRole - String数组 - （过滤条件）本地网络设备拓扑层次名字，形如：['LC','LA']。</li>
	// <li>RemoteRole - String数组 - （过滤条件）远端网络设备拓扑层次名字，形如：['WC','']。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNetDevicesTopoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDevicesTopoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServersTopoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络设备和服务器拓扑连接对象数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 网络设备和服务器拓扑连接对象数组

		NetServersTopoSet *NetServersTopo `json:"NetServersTopoSet,omitempty" name:"NetServersTopoSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServersTopoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServersTopoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGatewayClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGatewayClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGatewayClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewaysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewaysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWanDetailDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 探测数据明细列表

		DetailData []*DetailPingDataItem `json:"DetailData,omitempty" name:"DetailData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetWanDetailDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWanDetailDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetDeviceSyslogKeywordRequest struct {
	*tchttp.BaseRequest

	// 关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 网络设备厂商

	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`
	// 网络设备类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 网络设备角色

	Role *string `json:"Role,omitempty" name:"Role"`
	// 网络设备型号

	Model *string `json:"Model,omitempty" name:"Model"`
	// 唯一ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteNetDeviceSyslogKeywordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetDeviceSyslogKeywordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkDevicesRequest struct {
	*tchttp.BaseRequest

	// 修改网络设备信息

	ModifyNetDevices []*NetworkDevice `json:"ModifyNetDevices,omitempty" name:"ModifyNetDevices"`
}

func (r *ModifyNetworkDevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetworkDevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RackPositionInfo struct {
	// 机架

	Rack *string `json:"Rack,omitempty" name:"Rack"`
	// 机位

	Position *string `json:"Position,omitempty" name:"Position"`
	// 高度

	Height *string `json:"Height,omitempty" name:"Height"`
	// 逻辑区域

	LogicalArea *string `json:"LogicalArea,omitempty" name:"LogicalArea"`
	// 内网网络设备端口

	LanNetworkPort *string `json:"LanNetworkPort,omitempty" name:"LanNetworkPort"`
	// 外网网络设备端口

	WanNetworkPort *string `json:"WanNetworkPort,omitempty" name:"WanNetworkPort"`
	// 管理网络设备端口

	ManagementNetworkPort *string `json:"ManagementNetworkPort,omitempty" name:"ManagementNetworkPort"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 可用区

	Az *string `json:"Az,omitempty" name:"Az"`
	// 机房

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 机位状态（0:可用，1：预占用，2：占用，3：不可用）

	PositionStatus *int64 `json:"PositionStatus,omitempty" name:"PositionStatus"`
	// 内网网络设备名称-端口

	LanNetworkNamePort *string `json:"LanNetworkNamePort,omitempty" name:"LanNetworkNamePort"`
	// 外网网络设备名称-端口

	WanNetworkNamePort *string `json:"WanNetworkNamePort,omitempty" name:"WanNetworkNamePort"`
	// 管理网络设备名称-端口

	ManagementNetworkNamePort *string `json:"ManagementNetworkNamePort,omitempty" name:"ManagementNetworkNamePort"`
}

type TrapAlarm struct {
	// SNMP TRAP明细告警状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// SNMP TRAP明细告警主题

	Title *string `json:"Title,omitempty" name:"Title"`
	// SNMP TRAP明细告警创建时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// SNMP TRAP明细告警源ip地址

	DeviceIp *string `json:"DeviceIp,omitempty" name:"DeviceIp"`
	// SNMP TRAP明细告警源的物理端口

	DevicePort *string `json:"DevicePort,omitempty" name:"DevicePort"`
	// SNMP TRAP明细告警的详细描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// SNMP TRAP明细告警源名字

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type DescribeDeviceRoleProportionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 设备角色占比信息

		DevRolePropSet []*DevRoleProp `json:"DevRolePropSet,omitempty" name:"DevRolePropSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceRoleProportionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceRoleProportionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPSegmentFilterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 过滤值范围

		FilterMap *string `json:"FilterMap,omitempty" name:"FilterMap"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIPSegmentFilterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPSegmentFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpecialDciLineResult struct {
	// 专线id

	SpeLineId *string `json:"SpeLineId,omitempty" name:"SpeLineId"`
	// 专线名称

	SpeLineName *string `json:"SpeLineName,omitempty" name:"SpeLineName"`
}

type TopologyCount struct {
	// 服务器数目

	ServerCount *uint64 `json:"ServerCount,omitempty" name:"ServerCount"`
	// LA网络设备数目

	LaCount *uint64 `json:"LaCount,omitempty" name:"LaCount"`
	// LC网络设备数目

	LcCount *uint64 `json:"LcCount,omitempty" name:"LcCount"`
}

type CreateSpecialIdcLineRequest struct {
	*tchttp.BaseRequest

	// IDC专线信息

	SpecialIdcLine *SpecialIdcLine `json:"SpecialIdcLine,omitempty" name:"SpecialIdcLine"`
}

func (r *CreateSpecialIdcLineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSpecialIdcLineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportIdcExportLineFromDcosToNmsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 同步IDC专线出口数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportIdcExportLineFromDcosToNmsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportIdcExportLineFromDcosToNmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsageStatistic struct {
	// 统计指标维度

	Dimensions *Dimensions `json:"Dimensions,omitempty" name:"Dimensions"`
	// 统计数据

	Points *string `json:"Points,omitempty" name:"Points"`
}

type CreateRacksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRacksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDeviceSlotEnumRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// <li>SlotId - String数组  - （过滤条件）网络设备端口名字，形如：`1`。</li>
	// <li>AssetId - String数组  - （过滤条件）网络设备端口资产ID，形如：`TYNT190501G7`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNetDeviceSlotEnumRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDeviceSlotEnumRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportSpecialIdcLineFromDcosToNmsRequest struct {
	*tchttp.BaseRequest
}

func (r *ImportSpecialIdcLineFromDcosToNmsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportSpecialIdcLineFromDcosToNmsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRackPositionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRackPositionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRackPositionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServerRequest struct {
	*tchttp.BaseRequest

	// 可用区名字。

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 过滤条件。
	// <li>AssetId - String - （过滤条件）远端服务器设备资产编号，形如：`TYSV19060730`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为100，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportIPAddressesRequest struct {
	*tchttp.BaseRequest

	// IP信息列表

	IPs []*ImportIpAddresses `json:"IPs,omitempty" name:"IPs"`
}

func (r *ImportIPAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportIPAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRackPositionsRequest struct {
	*tchttp.BaseRequest

	// 创建机位信息

	CreateRackPositions []*RackPositionInfo `json:"CreateRackPositions,omitempty" name:"CreateRackPositions"`
}

func (r *CreateRackPositionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRackPositionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDeviceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络设备对象数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 网络设备对象数组

		NetdeviceSet []*Netdevice `json:"NetdeviceSet,omitempty" name:"NetdeviceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetDeviceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDeviceConfigsRequest struct {
	*tchttp.BaseRequest

	// 获取网络设备过滤条件

	Configs []*Configs `json:"Configs,omitempty" name:"Configs"`
}

func (r *DescribeNetDeviceConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDeviceConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLanDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 探测数据列表

		ChartData []*PingDataItem `json:"ChartData,omitempty" name:"ChartData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLanDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLanDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLanDetailDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 探测数据明细列表

		DetailData []*DetailPingDataItem `json:"DetailData,omitempty" name:"DetailData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLanDetailDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLanDetailDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRacksRequest struct {
	*tchttp.BaseRequest

	// 删除机架信息

	DeleteRacks []*RackInfo `json:"DeleteRacks,omitempty" name:"DeleteRacks"`
}

func (r *DeleteRacksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRacksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回机架机位信息

		RackSet []*RackPositionOverview `json:"RackSet,omitempty" name:"RackSet"`
		// 返回记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdcOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIPSegment struct {
	// 网段

	Segment *string `json:"Segment,omitempty" name:"Segment"`
}

type DeleteNetDeviceSyslogKeywordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNetDeviceSyslogKeywordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetDeviceSyslogKeywordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRacksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRacksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDevicesTopoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络设备拓扑连接关系数组

		NetDevicesTopoSet *NetDevicesTopo `json:"NetDevicesTopoSet,omitempty" name:"NetDevicesTopoSet"`
		// 网络设备拓扑连接对象数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetDevicesTopoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDevicesTopoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDeviceSyslogKeywordsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，Limit和Offset同时传入时生效

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 数量，Limit和Offset同时传入时生效

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件，支持Strategy/Keyword/AlarmType/Manufacturer/Type/Role/Model

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNetDeviceSyslogKeywordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDeviceSyslogKeywordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWanDataRequest struct {
	*tchttp.BaseRequest

	// 省份

	Province *string `json:"Province,omitempty" name:"Province"`
	// 运营商

	Isp *string `json:"Isp,omitempty" name:"Isp"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询指标："avg_delay", "lost_ip_count", "unreach_ip_count", "all"，“status”

	TargetTag *string `json:"TargetTag,omitempty" name:"TargetTag"`
	// 出口region

	ExitRegion *string `json:"ExitRegion,omitempty" name:"ExitRegion"`
	// 出口可用区

	ExitAz *string `json:"ExitAz,omitempty" name:"ExitAz"`
	// 出口名称

	ExitName *string `json:"ExitName,omitempty" name:"ExitName"`
}

func (r *GetWanDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWanDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkDevice struct {
	// 设备固资编号

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 网络设备SN

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 设备名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 设备类型(MA，WC-100G)

	Type *string `json:"Type,omitempty" name:"Type"`
	// 设备用途(内网接入交换机,外网接入交换机,核心交换机,防火墙)

	Role *string `json:"Role,omitempty" name:"Role"`
	// 设备厂家

	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`
	// 设备型号

	Model *string `json:"Model,omitempty" name:"Model"`
	// 机架名称

	Rack *string `json:"Rack,omitempty" name:"Rack"`
	// 管理ip

	ManagementIp *string `json:"ManagementIp,omitempty" name:"ManagementIp"`
	// 设备当前状态(0待运营 1运营中 2已下架)，更新接口必填

	Status *string `json:"Status,omitempty" name:"Status"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 可用区

	Az *string `json:"Az,omitempty" name:"Az"`
	// IDC机房

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 自增Id，更新接口必填

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DescribeNetDeviceSyslogKeywordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络设备日志关键字详情

		SyslogKeywords []*SyslogKeywords `json:"SyslogKeywords,omitempty" name:"SyslogKeywords"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetDeviceSyslogKeywordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDeviceSyslogKeywordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayClustersRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 条数限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeGatewayClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRacksRequest struct {
	*tchttp.BaseRequest

	// 修改机架信息

	ModifyRacks []*RackInfo `json:"ModifyRacks,omitempty" name:"ModifyRacks"`
}

func (r *ModifyRacksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRacksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRackPositionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRackPositionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRackPositionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGwDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 探测数据列表

		ChartData []*PingDataItem `json:"ChartData,omitempty" name:"ChartData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGwDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGwDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGatewaysRequest struct {
	*tchttp.BaseRequest

	// "Gateways":[{
	//         "AssetId":"TYSV120531X1",
	//         "Cluster":"VPCGW-CLUSTER1"
	//      }]

	Gateways []*DeleteGateway `json:"Gateways,omitempty" name:"Gateways"`
}

func (r *DeleteGatewaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGatewaysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpecialDciLineInfo struct {
	// 开始时间

	StTime *string `json:"StTime,omitempty" name:"StTime"`
	// 终止时间

	EdTime *string `json:"EdTime,omitempty" name:"EdTime"`
	// 专线ID

	SpeLineIds []*string `json:"SpeLineIds,omitempty" name:"SpeLineIds"`
}

type UpdateIPAddressesRequest struct {
	*tchttp.BaseRequest

	// 更新IP地址列表

	IPs []*UpdateIPAddress `json:"IPs,omitempty" name:"IPs"`
}

func (r *UpdateIPAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateIPAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DciSpeLineListRequest struct {
	*tchttp.BaseRequest

	// 起始az

	SpecialDciLinePar *SpecialDciLinePar `json:"SpecialDciLinePar,omitempty" name:"SpecialDciLinePar"`
}

func (r *DciSpeLineListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DciSpeLineListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLanDetailDataRequest struct {
	*tchttp.BaseRequest

	// 目标类型：la, lc

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 查询目标

	TargetName *string `json:"TargetName,omitempty" name:"TargetName"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *GetLanDetailDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLanDetailDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrapAlarmAbstractRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// <li>Title - String数组  - （过滤条件）TRAP告警主题，形如：`IF::linkupdown`。</li>
	// <li>DeviceIp - String数组  - （过滤条件）TRAP告警源设备IP，形如：`10.21.128.252`。</li>
	// <li>StartTime - String数组  - （过滤条件）TRAP告警开始时间，形如：`2020-07-14 00:00:00`。</li>
	// <li>EndTime - String数组  - （过滤条件）TRAP告警结束时间，形如：`2020-07-15 00:00:00`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTrapAlarmAbstractRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrapAlarmAbstractRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLanConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 探测配置项列表

		TargetList []*LanPingConfigItem `json:"TargetList,omitempty" name:"TargetList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLanConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLanConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTemplateConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTemplateConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTemplateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkRoleDictRequest struct {
	*tchttp.BaseRequest

	// 产品返回数量限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 返回记录偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询过滤条件， "Filters":[
	//         {
	//             "Name":"RoleEn",
	//             "Values":["BMS-LA-25G"]
	//         }
	//     ]。如果为空，查询返回全部记录。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNetworkRoleDictRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkRoleDictRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDciConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Dci探测线路列表

		LineLists []*DciLine `json:"LineLists,omitempty" name:"LineLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDciConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDciConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ip列表

		IPSet []*IPSet `json:"IPSet,omitempty" name:"IPSet"`
		// 记录总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIPAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetDeviceConfigBaseLineRequest struct {
	*tchttp.BaseRequest

	// 网络设备固定资产ID

	NetDeviceAsset *string `json:"NetDeviceAsset,omitempty" name:"NetDeviceAsset"`
	// 包含：Daily/Current，只有Daily时需要传递ConfigTime

	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
	// 日期

	ConfigTime *string `json:"ConfigTime,omitempty" name:"ConfigTime"`
}

func (r *ModifyNetDeviceConfigBaseLineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetDeviceConfigBaseLineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveTopoSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveTopoSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveTopoSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpecialIdcLine struct {
	// SpeLineId

	SpeLineId *string `json:"SpeLineId,omitempty" name:"SpeLineId"`
	// SpeLineName

	SpeLineName *string `json:"SpeLineName,omitempty" name:"SpeLineName"`
	// SpeLineType

	SpeLineType *int64 `json:"SpeLineType,omitempty" name:"SpeLineType"`
	// SpeLineStatus

	SpeLineStatus *int64 `json:"SpeLineStatus,omitempty" name:"SpeLineStatus"`
	// SpeLineFunc

	SpeLineFunc *int64 `json:"SpeLineFunc,omitempty" name:"SpeLineFunc"`
	// SpeLineSpeed

	SpeLineSpeed *string `json:"SpeLineSpeed,omitempty" name:"SpeLineSpeed"`
	// LocalDevice

	LocalDevice *string `json:"LocalDevice,omitempty" name:"LocalDevice"`
	// LocalDevicePort

	LocalDevicePort *string `json:"LocalDevicePort,omitempty" name:"LocalDevicePort"`
	// LocalDeviceIp

	LocalDeviceIp *string `json:"LocalDeviceIp,omitempty" name:"LocalDeviceIp"`
	// RemoteDevice

	RemoteDevice *string `json:"RemoteDevice,omitempty" name:"RemoteDevice"`
	// RemoteDevicePort

	RemoteDevicePort *string `json:"RemoteDevicePort,omitempty" name:"RemoteDevicePort"`
	// RemoteDeviceIp

	RemoteDeviceIp *string `json:"RemoteDeviceIp,omitempty" name:"RemoteDeviceIp"`
	// LineOperator

	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`
	// LocalOperatorOwner

	LocalOperatorOwner *string `json:"LocalOperatorOwner,omitempty" name:"LocalOperatorOwner"`
	// LocalOperatorOwnerPhone

	LocalOperatorOwnerPhone *string `json:"LocalOperatorOwnerPhone,omitempty" name:"LocalOperatorOwnerPhone"`
	// LocalOperatorWarningPhone

	LocalOperatorWarningPhone *string `json:"LocalOperatorWarningPhone,omitempty" name:"LocalOperatorWarningPhone"`
	// RemoteOperatorOwner

	RemoteOperatorOwner *string `json:"RemoteOperatorOwner,omitempty" name:"RemoteOperatorOwner"`
	// RemoteOperatorOwnerPhone

	RemoteOperatorOwnerPhone *string `json:"RemoteOperatorOwnerPhone,omitempty" name:"RemoteOperatorOwnerPhone"`
	// RemoteOperatorOwnerWarningPhone

	RemoteOperatorOwnerWarningPhone *string `json:"RemoteOperatorOwnerWarningPhone,omitempty" name:"RemoteOperatorOwnerWarningPhone"`
	// LineBusiness

	LineBusiness *string `json:"LineBusiness,omitempty" name:"LineBusiness"`
	// LineOwner

	LineOwner *string `json:"LineOwner,omitempty" name:"LineOwner"`
	// SlaNum

	SlaNum *string `json:"SlaNum,omitempty" name:"SlaNum"`
	// VlanInfo

	VlanInfo *string `json:"VlanInfo,omitempty" name:"VlanInfo"`
	// Comment

	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type DescribeGatewayClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRacksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回机架信息

		RackSet []*RackInfo `json:"RackSet,omitempty" name:"RackSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRacksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRacksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRacksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置模板对象数组

		TemplateResultSet *TemplateResult `json:"TemplateResultSet,omitempty" name:"TemplateResultSet"`
		// 配置模板对象数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescirbeNetDeviceSyslogsRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 支持的Filter包含：Keyword(string)、StartTime(datatime)、EndTime(datatime)。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 网络设备管理IP

	ManagementIp *string `json:"ManagementIp,omitempty" name:"ManagementIp"`
}

func (r *DescirbeNetDeviceSyslogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescirbeNetDeviceSyslogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkPortsRequest struct {
	*tchttp.BaseRequest

	// 修改网络设备端口信息

	ModifyNetPorts []*NetworkPort `json:"ModifyNetPorts,omitempty" name:"ModifyNetPorts"`
}

func (r *ModifyNetworkPortsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetworkPortsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GatewaySet struct {
	// 网关集群名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 资产ID

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 内部IP

	InnerIp *string `json:"InnerIp,omitempty" name:"InnerIp"`
	// 外部IP

	OuterIp *string `json:"OuterIp,omitempty" name:"OuterIp"`
}

type Netdevice struct {
	// 设备固资编号

	NetDeviceAssetId *string `json:"NetDeviceAssetId,omitempty" name:"NetDeviceAssetId"`
	// 网络设备SN

	NetDeviceSn *string `json:"NetDeviceSn,omitempty" name:"NetDeviceSn"`
	// 设备名称

	NetDeviceName *string `json:"NetDeviceName,omitempty" name:"NetDeviceName"`
	// 设备类型(交换机,路由器,防火墙)

	NetDeviceType *string `json:"NetDeviceType,omitempty" name:"NetDeviceType"`
	// 设备厂家

	NetDeviceProduct *string `json:"NetDeviceProduct,omitempty" name:"NetDeviceProduct"`
	// 设备型号

	NetDeviceModel *string `json:"NetDeviceModel,omitempty" name:"NetDeviceModel"`
	// 机架名称

	NetDeviceRackName *string `json:"NetDeviceRackName,omitempty" name:"NetDeviceRackName"`
	// 设备当前状态(0待运营 1运营中 2已下架)

	NetDeviceStatus *string `json:"NetDeviceStatus,omitempty" name:"NetDeviceStatus"`
	// 管理ip

	NetDeviceAdminIp *string `json:"NetDeviceAdminIp,omitempty" name:"NetDeviceAdminIp"`
	// 其他ip(多个)

	NetDeviceOtherIp *string `json:"NetDeviceOtherIp,omitempty" name:"NetDeviceOtherIp"`
	// 设备所属地域名字

	NetDeviceRegionName *string `json:"NetDeviceRegionName,omitempty" name:"NetDeviceRegionName"`
	// 设备所属可用区名字

	NetDeviceZoneName *string `json:"NetDeviceZoneName,omitempty" name:"NetDeviceZoneName"`
}

type ModifyIdcExportLineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IDC专线出口信息

		IdcExportLineResult []*IdcExportLineResult `json:"IdcExportLineResult,omitempty" name:"IdcExportLineResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIdcExportLineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIdcExportLineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDevicePortEnumResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络设备端口枚举对象数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 网络设备端口枚举对象数组

		NetDevicePortEnumSet []*NetDevicePortEnum `json:"NetDevicePortEnumSet,omitempty" name:"NetDevicePortEnumSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetDevicePortEnumResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDevicePortEnumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkDevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNetworkDevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpecialIdcLineTopoRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为100，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 可用区

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

func (r *DescribeSpecialIdcLineTopoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpecialIdcLineTopoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetDeviceConfigForNmsUpgradeRequest struct {
	*tchttp.BaseRequest

	// 过滤需要修改的网络设备

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ModifyNetDeviceConfigForNmsUpgradeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetDeviceConfigForNmsUpgradeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIPSegmentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 完成插入的网段数量

		TotalInsert *uint64 `json:"TotalInsert,omitempty" name:"TotalInsert"`
		// 和已有网段重复，无法插入的网段数量

		TotalRepeat *uint64 `json:"TotalRepeat,omitempty" name:"TotalRepeat"`
		// 重复网段列表

		RepeatSegment []*string `json:"RepeatSegment,omitempty" name:"RepeatSegment"`
		// 失败网段列表

		FailedSegment []*string `json:"FailedSegment,omitempty" name:"FailedSegment"`
		// 导入失败的条数

		TotalFailed *uint64 `json:"TotalFailed,omitempty" name:"TotalFailed"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIPSegmentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIPSegmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTrapAlarmResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTrapAlarmResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTrapAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportNetDeviceSyslogKeywordRequest struct {
	*tchttp.BaseRequest

	// 网络设备系统日志关键字

	SyslogKeywords []*SyslogKeywords `json:"SyslogKeywords,omitempty" name:"SyslogKeywords"`
}

func (r *ImportNetDeviceSyslogKeywordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportNetDeviceSyslogKeywordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceChannel struct {
	// 设备名字

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 端口通道

	PortChannel []*PortChannelList `json:"PortChannel,omitempty" name:"PortChannel"`
}

type CreateIdcExportLineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IDC专线出口

		IdcExportLineResult *IdcExportLineResult `json:"IdcExportLineResult,omitempty" name:"IdcExportLineResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIdcExportLineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIdcExportLineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWanDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 探测数据列表

		ChartData []*PingDataItem `json:"ChartData,omitempty" name:"ChartData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetWanDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWanDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Syslogs struct {
	// 日志上报时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 日志详情

	RawString *string `json:"RawString,omitempty" name:"RawString"`
}

type Zone struct {
	// 可用区名字

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type DescribeRacksRequest struct {
	*tchttp.BaseRequest

	// 查询过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回记录偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRacksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRacksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateConfigRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 条数限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTemplateConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplateConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLanConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetLanConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLanConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateIPSegmentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateIPSegmentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateIPSegmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WithdrawNetworkDevicesRequest struct {
	*tchttp.BaseRequest

	// 下架网络设备

	WithdrawNetDevices []*WithdrawNetDevice `json:"WithdrawNetDevices,omitempty" name:"WithdrawNetDevices"`
}

func (r *WithdrawNetworkDevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WithdrawNetworkDevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationLogRequest struct {
	*tchttp.BaseRequest

	// 条数限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeOperationLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopologyNetDeviceCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 拓扑网络设备数目

		TopologyCount *TopologyCount `json:"TopologyCount,omitempty" name:"TopologyCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopologyNetDeviceCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopologyNetDeviceCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkRoleDictRequest struct {
	*tchttp.BaseRequest

	// 修改网络设备角色中英文字典

	ModifyNetdevRoleDict []*NetdevRoleDict `json:"ModifyNetdevRoleDict,omitempty" name:"ModifyNetdevRoleDict"`
}

func (r *ModifyNetworkRoleDictRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetworkRoleDictRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetDevicePortList struct {
	// 网络设备名字

	NetdeviceName *string `json:"NetdeviceName,omitempty" name:"NetdeviceName"`
	// 网络设备端口列表

	PortList []*string `json:"PortList,omitempty" name:"PortList"`
}

type NetDevicesTopo struct {
	// 本地网络设备系统名字

	LocalName *string `json:"LocalName,omitempty" name:"LocalName"`
	// 本地网络设备ip地址

	LocalIp *string `json:"LocalIp,omitempty" name:"LocalIp"`
	// 本地网络设备拓扑层次角色名字

	LocalRole *string `json:"LocalRole,omitempty" name:"LocalRole"`
	// 本地网络设备所属地域名字

	LocalRegion *string `json:"LocalRegion,omitempty" name:"LocalRegion"`
	// 本地网络设备所属AZ名字

	LocalZone *string `json:"LocalZone,omitempty" name:"LocalZone"`
	// 本地网络设备互联端口速率

	LocalSpeed *string `json:"LocalSpeed,omitempty" name:"LocalSpeed"`
	// 本地网络设备互联端口名字

	LocalPort *string `json:"LocalPort,omitempty" name:"LocalPort"`
	// 远端网络设备系统名字

	RemoteName *string `json:"RemoteName,omitempty" name:"RemoteName"`
	// 远端网络设备ip地址

	RemoteIp *string `json:"RemoteIp,omitempty" name:"RemoteIp"`
	// 远端网络设备拓扑层次名字

	RemoteRole *string `json:"RemoteRole,omitempty" name:"RemoteRole"`
	// 远端网络设备所属地域名字

	RemoteRegion *string `json:"RemoteRegion,omitempty" name:"RemoteRegion"`
	// 远端网络设备所属AZ名字

	RemoteZone *string `json:"RemoteZone,omitempty" name:"RemoteZone"`
	// 远端网络设备互联端口速率

	RemoteSpeed *string `json:"RemoteSpeed,omitempty" name:"RemoteSpeed"`
	// 远端网络设备互联端口名字

	RemotePort *string `json:"RemotePort,omitempty" name:"RemotePort"`
}

type PingDataItem struct {
	// 时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 值

	Value *float64 `json:"Value,omitempty" name:"Value"`
	// 数据项名称

	Instance *string `json:"Instance,omitempty" name:"Instance"`
}

type DescribeIdcExportLineRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为100，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// <li>ExportLineId - String - （过滤条件）IDC专线出口id，形如：`test`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeIdcExportLineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcExportLineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为100，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportIdcExportLineFromDcosToNmsRequest struct {
	*tchttp.BaseRequest
}

func (r *ImportIdcExportLineFromDcosToNmsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportIdcExportLineFromDcosToNmsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可用区对象数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 可用区对象数组

		ZoneSet *Zone `json:"ZoneSet,omitempty" name:"ZoneSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOperationLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 获取到的操作日志对象信息

		Logs []*Operationlog `json:"Logs,omitempty" name:"Logs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRackPositionsRequest struct {
	*tchttp.BaseRequest

	// 修改机位信息

	ModifyRackPositions []*RackPositionInfo `json:"ModifyRackPositions,omitempty" name:"ModifyRackPositions"`
}

func (r *ModifyRackPositionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRackPositionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySpecialIdcLineRequest struct {
	*tchttp.BaseRequest

	// IDC专线信息

	SpecialIdcLine *SpecialIdcLine `json:"SpecialIdcLine,omitempty" name:"SpecialIdcLine"`
}

func (r *ModifySpecialIdcLineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySpecialIdcLineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RspIPAddress struct {
	// IP地址

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// IP所属网段

	Segment *string `json:"Segment,omitempty" name:"Segment"`
}

type DescribeNetDevicePortStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络设备端口对象数组

		NetDevicePortDescSet *NetDevicePortDesc `json:"NetDevicePortDescSet,omitempty" name:"NetDevicePortDescSet"`
		// 网络设备端口对象数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetDevicePortStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDevicePortStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkPortsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回记录偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNetworkPortsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkPortsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGatewaysRequest struct {
	*tchttp.BaseRequest

	// "Gateways": [
	//         {
	//             "Cluster": "xxx",   //需要扩容的Cluster名称
	//             "AssetId": "xxxx", //需要添加的网关资产id
	//             "InnerIp": "xxxx",
	//             "OuterIp": “xxx”
	//         }
	//     ]

	Gateways []*AddGateway `json:"Gateways,omitempty" name:"Gateways"`
}

func (r *CreateGatewaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGatewaysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateConfig struct {
	// 网络设备类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 网络设备厂商

	Vendor *string `json:"Vendor,omitempty" name:"Vendor"`
	// 网络设备型号

	Model *string `json:"Model,omitempty" name:"Model"`
	// 网络设备绑定模板名字

	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
}

type Server struct {
	// 设备固资编号

	ServerAssetId *string `json:"ServerAssetId,omitempty" name:"ServerAssetId"`
	// 服务器SN

	ServereSn *string `json:"ServereSn,omitempty" name:"ServereSn"`
	// 设备别名

	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`
	// 设备型号

	ServerType *string `json:"ServerType,omitempty" name:"ServerType"`
	// IDC名称

	ServerIdcName *string `json:"ServerIdcName,omitempty" name:"ServerIdcName"`
	// IDC序号

	ServerIdcId *string `json:"ServerIdcId,omitempty" name:"ServerIdcId"`
	// 机架名称

	ServerRackName *string `json:"ServerRackName,omitempty" name:"ServerRackName"`
	// 设备内网IP

	ServerLanIp *string `json:"ServerLanIp,omitempty" name:"ServerLanIp"`
	// 设备外网IP

	ServerWanIp *string `json:"ServerWanIp,omitempty" name:"ServerWanIp"`
	// 设备外网eip

	ServerWanEip *string `json:"ServerWanEip,omitempty" name:"ServerWanEip"`
	// 设备所属地区名字

	ServerRegionName *string `json:"ServerRegionName,omitempty" name:"ServerRegionName"`
	// 设备所属可用区名字

	ServerZoneName *string `json:"ServerZoneName,omitempty" name:"ServerZoneName"`
}

type DescribeNetDevicePortChannelListRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为500。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// <li>PortName - String数组  - （过滤条件）网络设备端口名字，形如：`GigabitEthernet1/0/10`。</li>
	// <li>AssetId - String数组  - （过滤条件）网络设备端口资产ID，形如：`TYNT190501G7`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNetDevicePortChannelListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDevicePortChannelListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGatewayClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGatewayClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateGatewayClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescirbeGlobalNetDeviceSyslogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 全局的网络设备日志信息详情

		GlobalSyslogs []*GlobalSyslogs `json:"GlobalSyslogs,omitempty" name:"GlobalSyslogs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescirbeGlobalNetDeviceSyslogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescirbeGlobalNetDeviceSyslogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServersTopoRequest struct {
	*tchttp.BaseRequest

	// 可用区名字。

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 过滤条件。
	// <li>LocalIp - String - （过滤条件）本地网络设备ip地址，形如：`10.21.128.252`。</li>
	// <li>RemoteIp - String - （过滤条件）远端服务器设备ip地址，形如：`10.21.128.253`。</li>
	// <li>AssetId - String - （过滤条件）远端服务器设备资产编号，形如：`TYSV19060730`。</li>
	// <li>Role - String - （过滤条件）本地网络设备拓扑层次，形如：`WC`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，Limit和Offset同时传入时生效

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，Limit和Offset同时传入时生效

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeServersTopoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServersTopoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpecialIdcLineRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为100，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// <li>SpeLineId - String - （过滤条件）IDC专线id，形如：`test`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeSpecialIdcLineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpecialIdcLineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpecialIdcLineMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标名称

		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
		// 监控统计周期

		Period *uint64 `json:"Period,omitempty" name:"Period"`
		// 起始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 监控数据列表

		DataPoints []*DataPoints `json:"DataPoints,omitempty" name:"DataPoints"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSpecialIdcLineMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpecialIdcLineMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetDeviceConfigBaseLineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetDeviceConfigBaseLineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetDeviceConfigBaseLineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SegmentSet struct {
	// 网段名称

	Segment *string `json:"Segment,omitempty" name:"Segment"`
	// 网段掩码

	Mask *string `json:"Mask,omitempty" name:"Mask"`
	// 网关

	Gateway *string `json:"Gateway,omitempty" name:"Gateway"`
	// 可分配IP

	Assignable *string `json:"Assignable,omitempty" name:"Assignable"`
	// 裂解规则

	CrackRule *int64 `json:"CrackRule,omitempty" name:"CrackRule"`
	// 逻辑区域

	LogicalArea *string `json:"LogicalArea,omitempty" name:"LogicalArea"`
	// 网段用途

	SegmentUsage *string `json:"SegmentUsage,omitempty" name:"SegmentUsage"`
	// 网段所属产品

	SegmentProduct *string `json:"SegmentProduct,omitempty" name:"SegmentProduct"`
	// IP版本，IPv4

	Version *string `json:"Version,omitempty" name:"Version"`
	// vlan id

	VlanId *int64 `json:"VlanId,omitempty" name:"VlanId"`
	// 交换机

	Switch *string `json:"Switch,omitempty" name:"Switch"`
	// IDC机房

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 可用区

	Az *string `json:"Az,omitempty" name:"Az"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 网段中总的IP个数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
	// 网段中可分配的IP地址数量

	AssignableCnt *uint64 `json:"AssignableCnt,omitempty" name:"AssignableCnt"`
	// 网段中已经使用的IP地址数量

	Used *uint64 `json:"Used,omitempty" name:"Used"`
}

type DescribeDeviceVendorProportionRequest struct {
	*tchttp.BaseRequest

	// 根据过滤条件，返回指定idc厂商占比
	// "Filters":[
	//         {
	//             “Name”："Idc"，
	//             “Values”:["YF-1"]
	//         }
	//     ]

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDeviceVendorProportionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceVendorProportionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDeviceSlotEnumResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络设备单板枚举对象数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 网络设备单板枚举对象数组

		NetDeviceSlotEnumSet []*Filter `json:"NetDeviceSlotEnumSet,omitempty" name:"NetDeviceSlotEnumSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetDeviceSlotEnumResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDeviceSlotEnumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGwConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群列表

		SetList []*string `json:"SetList,omitempty" name:"SetList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGwConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGwConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkRoleDictResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNetworkRoleDictResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkRoleDictResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveTopoSettingRequest struct {
	*tchttp.BaseRequest

	// 网络拓扑位置配置

	Setting *string `json:"Setting,omitempty" name:"Setting"`
	// 可用区

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

func (r *SaveTopoSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveTopoSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddGatewayCluster struct {
	// 网关集群名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 网关

	Gateways *string `json:"Gateways,omitempty" name:"Gateways"`
	// 产品名称

	Product *string `json:"Product,omitempty" name:"Product"`
	// VIP信息

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 网关扩展信息

	ClusterExpand *string `json:"ClusterExpand,omitempty" name:"ClusterExpand"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
}

type LanPingConfigItem struct {
	// 所属Region

	Region *string `json:"Region,omitempty" name:"Region"`
	// 内网交换机

	La *string `json:"La,omitempty" name:"La"`
	// 核心交换机

	Lc *string `json:"Lc,omitempty" name:"Lc"`
}

type ProbeList struct {
	// region中文名

	RegionCN *string `json:"RegionCN,omitempty" name:"RegionCN"`
	// region英文名A

	Region *string `json:"Region,omitempty" name:"Region"`
	// 可用区中文名

	AzCN *string `json:"AzCN,omitempty" name:"AzCN"`
	// 可用区英文名

	Az *string `json:"Az,omitempty" name:"Az"`
	// 出口名称

	IdcExitName *string `json:"IdcExitName,omitempty" name:"IdcExitName"`
	// 出口运营商

	IdcExitIsp *string `json:"IdcExitIsp,omitempty" name:"IdcExitIsp"`
	// 探测机IP

	IP *string `json:"IP,omitempty" name:"IP"`
}

type CreateNetworkRoleDictResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetworkRoleDictResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkRoleDictResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTemplateRequest struct {
	*tchttp.BaseRequest

	// 配置模板信息

	Template *string `json:"Template,omitempty" name:"Template"`
}

func (r *CreateTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置模板对象数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 配置模板对象数组

		TemplateResultSet *TemplateResult `json:"TemplateResultSet,omitempty" name:"TemplateResultSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkRoleDictRequest struct {
	*tchttp.BaseRequest

	// 增加网络设备角色中英文字典

	CreateNetdevRoleDict []*NetdevRoleDict `json:"CreateNetdevRoleDict,omitempty" name:"CreateNetdevRoleDict"`
}

func (r *CreateNetworkRoleDictRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkRoleDictRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDevicePortListRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为500。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// <li>PortName - String数组  - （过滤条件）网络设备端口名字，形如：`GigabitEthernet1/0/10`。</li>
	// <li>AssetId - String数组  - （过滤条件）网络设备资产ID，形如：`TYNT190501G7`。</li>
	// <li>DeviceName - String数组  - （过滤条件）网络设备名字，形如：`CQ-YF-201-D17-H5130-MA-01`。</li>
	// <li>NetdeviceRole - String数组  - （过滤条件）网络设备层次类型，形如：`LA`。</li>
	// <li>PortRole - String数组  - （过滤条件）网络设备端口类型，形如：`GigabitEthernet`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNetDevicePortListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDevicePortListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkDevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetworkDevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopoSettingRequest struct {
	*tchttp.BaseRequest

	// 可用区

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

func (r *DeleteTopoSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTopoSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDevicesTopoExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络设备拓扑连接角色关系数组

		NetDevicesTopoExSet []*NetDevicesTopoEx `json:"NetDevicesTopoExSet,omitempty" name:"NetDevicesTopoExSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetDevicesTopoExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDevicesTopoExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPSegmentsRequest struct {
	*tchttp.BaseRequest

	// 过滤器

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，Limit和Offset同时传入时生效

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 条数限制，Limit和Offset同时传入时生效

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeIPSegmentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPSegmentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWanExitStatusStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 出口状态列表

		StatusList []*WanStatus `json:"StatusList,omitempty" name:"StatusList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetWanExitStatusStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWanExitStatusStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGatewayClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGatewayClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGatewayClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopologyNetDeviceCountRequest struct {
	*tchttp.BaseRequest

	// 可用区名字

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

func (r *DescribeTopologyNetDeviceCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopologyNetDeviceCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterSet struct {
	// 网关集群名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 网关集群包含的网关列表

	Gateways []*Gateway `json:"Gateways,omitempty" name:"Gateways"`
	// 网关产品名称

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集群扩展信息

	ClusterExpand *string `json:"ClusterExpand,omitempty" name:"ClusterExpand"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// VIP信息

	Vip []*GatewayVip `json:"Vip,omitempty" name:"Vip"`
}

type CreateIPSegmentsRequest struct {
	*tchttp.BaseRequest

	// 增加IP网段类型，"IPSegments": [{
	//      "Segment":"10.0.1.0/24",
	//      "Mask":"255.255.255.0",
	//      "Gateway":"10.0.1.1",
	//      "CrackRule":32,
	//      "Assignable":"[{\"StartIP\":\"10.0.1.2\",\"EndIP\":\"10.0.1.65\",\"Type\":\"SE\"}]",
	//      "LogicalArea":"NFV",
	//      "SegmentUsage":"NFV-VM",
	//      "Version":"IPv4",
	//      "VlanId":3,
	//      "Switch":"TYNT190501EE,TYNT190501GF",
	//      "Idc":"YF-1",
	//      "Az":"YF",
	//      "Region":"chongqing"
	// }]

	IPSegments []*AddIPSegment `json:"IPSegments,omitempty" name:"IPSegments"`
}

func (r *CreateIPSegmentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIPSegmentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDeviceSlotListRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// <li>SlotId - String数组  - （过滤条件）网络设备端口名字，形如：`1`。</li>
	// <li>AssetId - String数组  - （过滤条件）网络设备端口资产ID，形如：`TYNT190501G7`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNetDeviceSlotListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDeviceSlotListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCommonServicePingDataRequest struct {
	*tchttp.BaseRequest

	// 服务类型：yum，ntp，dns

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询指标："avg_delay", "lost_ip_count", "unreach_ip_count", "all"

	TargetTag *string `json:"TargetTag,omitempty" name:"TargetTag"`
}

func (r *GetCommonServicePingDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCommonServicePingDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIdcExportLineRequest struct {
	*tchttp.BaseRequest

	// IDC专线出口信息

	IdcExportLine *IdcExportLine `json:"IdcExportLine,omitempty" name:"IdcExportLine"`
}

func (r *ModifyIdcExportLineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIdcExportLineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTrapAlarmLevelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTrapAlarmLevelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTrapAlarmLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDevicePortChannelEnumRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为500。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// <li>PortName - String数组  - （过滤条件）网络设备端口名字，形如：`GigabitEthernet1/0/10`。</li>
	// <li>AssetId - String数组  - （过滤条件）网络设备端口资产ID，形如：`TYNT190501G7`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNetDevicePortChannelEnumRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDevicePortChannelEnumRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Field2Enum struct {
	// key

	Key *string `json:"Key,omitempty" name:"Key"`
}

type DescribeBandWidthUsageTop5Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 贷款利用率统计信息

		BandWidthUsageStatistics []*UsageStatistic `json:"BandWidthUsageStatistics,omitempty" name:"BandWidthUsageStatistics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBandWidthUsageTop5Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBandWidthUsageTop5Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置模板绑定对象数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 配置模板绑定对象数组

		TemplateConfigResultSet []*TemplateConfigResult `json:"TemplateConfigResultSet,omitempty" name:"TemplateConfigResultSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTemplateConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateIPAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateIPAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateIPAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RolePair struct {
	// 本地网络设备角色名字

	LocalRole *string `json:"LocalRole,omitempty" name:"LocalRole"`
	// 远端网络设备角色名字

	RemoteRole *string `json:"RemoteRole,omitempty" name:"RemoteRole"`
}

type DeleteNetworkPortsRequest struct {
	*tchttp.BaseRequest

	// 删除网络设备端口信息

	DeleteNetPorts []*NetworkPort `json:"DeleteNetPorts,omitempty" name:"DeleteNetPorts"`
}

func (r *DeleteNetworkPortsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkPortsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIPCIDRSegmentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IPCIDR网段信息

		IPCIDRSegmentSet []*IPCIDRSegment `json:"IPCIDRSegmentSet,omitempty" name:"IPCIDRSegmentSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIPCIDRSegmentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIPCIDRSegmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGwDetailDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 探测数据明细列表

		DetailData []*DetailPingDataItem `json:"DetailData,omitempty" name:"DetailData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGwDetailDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGwDetailDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIdcExportLineRequest struct {
	*tchttp.BaseRequest

	// IDC专线出口信息

	IdcExportLine *IdcExportLine `json:"IdcExportLine,omitempty" name:"IdcExportLine"`
}

func (r *CreateIdcExportLineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIdcExportLineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddIpAddress struct {
	// IP地址列表

	IpAddresses []*SingleIPAddress `json:"IpAddresses,omitempty" name:"IpAddresses"`
	// 交换机

	Switch *string `json:"Switch,omitempty" name:"Switch"`
	// 网段用途

	SegmentUsage *string `json:"SegmentUsage,omitempty" name:"SegmentUsage"`
	// 网段所属产品

	SegmentProduct *string `json:"SegmentProduct,omitempty" name:"SegmentProduct"`
	// 申请数量

	AllocateCnt *uint64 `json:"AllocateCnt,omitempty" name:"AllocateCnt"`
	// 是否需要连续

	Continuously *bool `json:"Continuously,omitempty" name:"Continuously"`
	// 标识符

	ObjectId *string `json:"ObjectId,omitempty" name:"ObjectId"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// IP类型信息

	Type *string `json:"Type,omitempty" name:"Type"`
	// 申请人

	OwnerName *string `json:"OwnerName,omitempty" name:"OwnerName"`
	// 机架信息

	Rack *string `json:"Rack,omitempty" name:"Rack"`
	// 服务器资产号

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// vlan号

	VlanId *int64 `json:"VlanId,omitempty" name:"VlanId"`
}

type CreateNetDeviceSyslogKeywordRequest struct {
	*tchttp.BaseRequest

	// 策略名称

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 告警类型，包含："emergencies"/"alerts"/"critical"/"errors"/"warnings"/"notifications"/"informational"/"debugging"

	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`
	// 网络设备厂商

	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`
	// 网络设备类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 网络设备角色

	Role *string `json:"Role,omitempty" name:"Role"`
	// 网络设备型号

	Model *string `json:"Model,omitempty" name:"Model"`
	// 说明

	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *CreateNetDeviceSyslogKeywordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetDeviceSyslogKeywordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetDevicesTopoEx struct {
	// 网络设备角色

	NetDevicesRole []*string `json:"NetDevicesRole,omitempty" name:"NetDevicesRole"`
	// 网络设备拓扑连接关系数组

	NetDevicesTopoSet []*NetDevicesTopo `json:"NetDevicesTopoSet,omitempty" name:"NetDevicesTopoSet"`
}

type RackInfo struct {
	// Region名称

	Region *string `json:"Region,omitempty" name:"Region"`
	// 可用区

	Az *string `json:"Az,omitempty" name:"Az"`
	// 机房

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 机柜编号

	Rack *string `json:"Rack,omitempty" name:"Rack"`
	// 机柜类型（NETWORK：网络柜，SERVER：服务器柜）

	RackType *string `json:"RackType,omitempty" name:"RackType"`
	// 是否通电（0：未通电，1：通电）

	PowerStatus *int64 `json:"PowerStatus,omitempty" name:"PowerStatus"`
}

type DciCapacityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DciCapacityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DciCapacityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescirbeNetDeviceSyslogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络设备系统日志详情

		Syslogs []*Syslogs `json:"Syslogs,omitempty" name:"Syslogs"`
		// 网络设备系统日志数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescirbeNetDeviceSyslogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescirbeNetDeviceSyslogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDciConfigRequest struct {
	*tchttp.BaseRequest

	// 探测源地域

	SrcArea *string `json:"SrcArea,omitempty" name:"SrcArea"`
}

func (r *GetDciConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDciConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DevRoleProp struct {
	// 设备角色

	DevRole *string `json:"DevRole,omitempty" name:"DevRole"`
	// 角色百分比

	Proportion *float64 `json:"Proportion,omitempty" name:"Proportion"`
	// 数量

	DevNum *int64 `json:"DevNum,omitempty" name:"DevNum"`
}

type DeleteGateway struct {
	// 资产ID

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 集群名称，删除集群下的所有网关

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

type DescribeNetworkDevicesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，Limit和Offset同时传入时生效

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移，Limit和Offset同时传入时生效

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNetworkDevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkDevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TrapAlarmAbstract struct {
	// SNMP TRAP概要告警主题

	Title *string `json:"Title,omitempty" name:"Title"`
	// SNMP TRAP概要告警等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// SNMP TRAP概要告警更新时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// SNMP TRAP概要告警源IP地址

	DeviceIp *string `json:"DeviceIp,omitempty" name:"DeviceIp"`
	// SNMP TRAP概要告警源端口

	DevciePort *string `json:"DevciePort,omitempty" name:"DevciePort"`
	// SNMP TRAP概要告警发生次数

	Count *string `json:"Count,omitempty" name:"Count"`
	// SNMP TRAP概要告警描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// SNMP TRAP概要告警源IP名字

	DevcieName *string `json:"DevcieName,omitempty" name:"DevcieName"`
}

type DeleteIPAddressesRequest struct {
	*tchttp.BaseRequest

	// "IPs": [
	//       {
	//         "IpAddress":"xxx",
	//         "Segment":"xxxx" //只有在没有指定IpAddress地址的时候有效，说明释放此网段下的所有ip地址
	//       }

	IPs []*DeleteIPAddress `json:"IPs,omitempty" name:"IPs"`
}

func (r *DeleteIPAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIPAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBandWidthUsageTop5Request struct {
	*tchttp.BaseRequest

	// 根据过滤条件 "Filters":[
	//         {
	//             "Name":"Idc",
	//             "Values":["YF-1"]
	//         }
	//     ]返回指定idc的带宽利用率top5,如果Filters为空表示返回整个region的带宽利用率top5。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeBandWidthUsageTop5Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBandWidthUsageTop5Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcExportLineMetricRequest struct {
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
	// 可用区名字

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// IDC专线出口名字

	ExportLineName *string `json:"ExportLineName,omitempty" name:"ExportLineName"`
}

func (r *DescribeIdcExportLineMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcExportLineMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPAddressesRequest struct {
	*tchttp.BaseRequest

	// 条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeIPAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkDevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetworkDevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetworkDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type Gateway struct {
	// 资产ID

	AssertId *string `json:"AssertId,omitempty" name:"AssertId"`
	// 内部IP地址

	InnerIp *string `json:"InnerIp,omitempty" name:"InnerIp"`
	// 外部IP地址

	OuterIp *string `json:"OuterIp,omitempty" name:"OuterIp"`
}

type PortChannelList struct {
	// 端口名字

	PortName *string `json:"PortName,omitempty" name:"PortName"`
	// 端口下通道列表

	ChannelList []*string `json:"ChannelList,omitempty" name:"ChannelList"`
}

type RackInfoOld struct {
	// 机架

	Rack *string `json:"Rack,omitempty" name:"Rack"`
	// 机位

	Position *string `json:"Position,omitempty" name:"Position"`
	// 高度

	Height *string `json:"Height,omitempty" name:"Height"`
	// 逻辑区域

	LogicalArea *string `json:"LogicalArea,omitempty" name:"LogicalArea"`
	// 内网网络设备端口

	LanNetworkPort *string `json:"LanNetworkPort,omitempty" name:"LanNetworkPort"`
	// 外网网络设备端口

	WanNetworkPort *string `json:"WanNetworkPort,omitempty" name:"WanNetworkPort"`
	// 管理网络设备端口

	ManagementNetworkPort *string `json:"ManagementNetworkPort,omitempty" name:"ManagementNetworkPort"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 可用区

	Az *string `json:"Az,omitempty" name:"Az"`
	// 机房

	Idc *string `json:"Idc,omitempty" name:"Idc"`
}

type DescribeNetworkRoleDictResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回网络设备角色字典

		NetworkRoleDictSet []*NetdevRoleDict `json:"NetworkRoleDictSet,omitempty" name:"NetworkRoleDictSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkRoleDictResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkRoleDictResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkPortsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetworkPortsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkPortsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIPSegmentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除失败的网段数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 删除失败的网段列表

		FailedSet []*string `json:"FailedSet,omitempty" name:"FailedSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIPSegmentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIPSegmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddIPSegment struct {
	// 网段名称，例"Segment": "10.0.0.0/24"

	Segment *string `json:"Segment,omitempty" name:"Segment"`
	// 网段掩码，例"Mask": "255.255.255.0"

	Mask *string `json:"Mask,omitempty" name:"Mask"`
	// 网段网关，例"Gateway": "10.0.0.1"

	Gateway *string `json:"Gateway,omitempty" name:"Gateway"`
	// 可以分配的IP地址范围，例"Assignable": "[{'StartIP': '10.0.0.2', 'EndIP': '10.0.0.65', "Type":'SE'}]"

	Assignable *string `json:"Assignable,omitempty" name:"Assignable"`
	// 裂解规则，整数

	CrackRule *uint64 `json:"CrackRule,omitempty" name:"CrackRule"`
	// 逻辑区域

	LogicalArea *string `json:"LogicalArea,omitempty" name:"LogicalArea"`
	// 网段用途

	SegmentUsage *string `json:"SegmentUsage,omitempty" name:"SegmentUsage"`
	// 网段所属产品

	SegmentProduct *string `json:"SegmentProduct,omitempty" name:"SegmentProduct"`
	// 网段ip类型，取值为IPv4

	Version *string `json:"Version,omitempty" name:"Version"`
	// vlan id

	VlanId *int64 `json:"VlanId,omitempty" name:"VlanId"`
	// 交换机

	Switch *string `json:"Switch,omitempty" name:"Switch"`
	// IDC机房名称

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 可用区

	Az *string `json:"Az,omitempty" name:"Az"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 网络设备名称

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type TemplateConfigResult struct {
	// 网络设备模板配置信息ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 网络设备型号

	Model *string `json:"Model,omitempty" name:"Model"`
	// 网络设备配置模板名字

	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
	// 网络设备配置模板时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 网络设备类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 网络设备厂商

	Vendor *string `json:"Vendor,omitempty" name:"Vendor"`
}

type WithdrawNetDevice struct {
	// 资产号

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 机架

	Rack *string `json:"Rack,omitempty" name:"Rack"`
	// 管理IP

	ManagementIp *string `json:"ManagementIp,omitempty" name:"ManagementIp"`
}

type CreateIPCIDRSegmentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIPCIDRSegmentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIPCIDRSegmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWanExitStatusStatRequest struct {
	*tchttp.BaseRequest
}

func (r *GetWanExitStatusStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWanExitStatusStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySpecialIdcLineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IDC专线信息

		SpecialIdcLineResult []*SpecialIdcLineResult `json:"SpecialIdcLineResult,omitempty" name:"SpecialIdcLineResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySpecialIdcLineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySpecialIdcLineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络服务器对象数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 网络服务器对象数组

		ServerSet []*Server `json:"ServerSet,omitempty" name:"ServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpecialIdcLineTopo struct {
	// IDC专线本地设备系统名字

	LocalName *string `json:"LocalName,omitempty" name:"LocalName"`
	// IDC专线本地设备ip地址

	LocalIp *string `json:"LocalIp,omitempty" name:"LocalIp"`
	// IDC专线本地设备物理出口

	LocalPort *string `json:"LocalPort,omitempty" name:"LocalPort"`
	// IDC专线远端设备系统名称

	RemoteName *string `json:"RemoteName,omitempty" name:"RemoteName"`
	// IDC专线远端设备ip地址

	RemoteIp *string `json:"RemoteIp,omitempty" name:"RemoteIp"`
	// IDC专线远端设备物理出口

	RemotePort *string `json:"RemotePort,omitempty" name:"RemotePort"`
	// IDC专线端口速率

	Speed *string `json:"Speed,omitempty" name:"Speed"`
	// IDC专线本地可用区

	LocalZone *string `json:"LocalZone,omitempty" name:"LocalZone"`
	// IDC专线远端可用区

	RemoteZone *string `json:"RemoteZone,omitempty" name:"RemoteZone"`
	// IDC专线名字

	SpeLineName *string `json:"SpeLineName,omitempty" name:"SpeLineName"`
}

type DescribeNetdeviceFieldsEnumRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNetdeviceFieldsEnumRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetdeviceFieldsEnumRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkDevicesExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回网络设备信息

		NetDeviceSet []*NetworkDevice `json:"NetDeviceSet,omitempty" name:"NetDeviceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkDevicesExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkDevicesExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIPAddressesRequest struct {
	*tchttp.BaseRequest

	// "IPs": [
	//       {
	//         "Version": "IPv4",            // 必须填写的字段，目前就支持IPv4
	//         "IpAddresses": [{"Ip":"xxxx/32", "ObjectId":"xxx", "OwnerName":"xxx", "Description":"xxx"}], // 指定IP地址申请，可以指定多个，如果指定了这个，就不需要指定其他的后续参数
	//         "Switch": "xxx,xxx",        // 所属交换机
	//         "SegmentUsage": "xxx",      // 网段用途就能覆盖逻辑区信息了，所以不再传递LogicArea。
	//         "SegmentProduct": "xxx"     // 所属产品，主要应用于NFV-VM/NFVW-VM
	//         "AllocateCnt": 5,           // 申请IP个数，不指定就默认申请一个，此处参数生效是在IpAddresses参数没有传入的情况下
	//         "Continuously": True,       // 申请多个IP时有效，是否需要IP地址连续，默认连续IP地址只能在一个网段中申请
	//         "ObjectId": "xxxx",          // 资产ID，主要用于标示IP所属，方便后续查询。
	//         "Description": "xxx"               // 描述信息
	//       }
	//     ]

	IPs []*AddIpAddress `json:"IPs,omitempty" name:"IPs"`
}

func (r *CreateIPAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIPAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFieldsEnumResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回枚举结果

		Field2Enum *Field2Enum `json:"Field2Enum,omitempty" name:"Field2Enum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFieldsEnumResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFieldsEnumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDeviceSlotListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络设备单板列表对象数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 网络设备单板列表对象数组

		NetDeviceSlotListSet []*NetDeviceSlotList `json:"NetDeviceSlotListSet,omitempty" name:"NetDeviceSlotListSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetDeviceSlotListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDeviceSlotListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTrapAlarmLevelRequest struct {
	*tchttp.BaseRequest

	// SNMP TRAP告警主题

	Title *string `json:"Title,omitempty" name:"Title"`
	// SNMP TRAP告警等级

	Level *string `json:"Level,omitempty" name:"Level"`
}

func (r *ModifyTrapAlarmLevelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTrapAlarmLevelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRacksOldRequest struct {
	*tchttp.BaseRequest

	// 创建机架信息

	CreateRacks []*RackInfoOld `json:"CreateRacks,omitempty" name:"CreateRacks"`
}

func (r *CreateRacksOldRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRacksOldRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DciSpeLineListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// dci专线列表

		SpecialDciLineResult []*SpecialDciLineResult `json:"SpecialDciLineResult,omitempty" name:"SpecialDciLineResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DciSpeLineListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DciSpeLineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcExportLineTopoRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为100，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// zone名字，形如："yf-1"

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

func (r *DescribeIdcExportLineTopoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcExportLineTopoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcExportLineTopoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的专线拓扑连接对象数目。

		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`
		// 专线拓扑连接对象数组

		IdcExportLineTopoSet *IdcExportLineTopo `json:"IdcExportLineTopoSet,omitempty" name:"IdcExportLineTopoSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdcExportLineTopoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcExportLineTopoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRackPositionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRackPositionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRackPositionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataPoints struct {
	// 监控实例的维度组合

	Dimensions *Dimensions `json:"Dimensions,omitempty" name:"Dimensions"`
	// 监控数据点数组，每个点的时间跨度为一个Period值

	Points []*float64 `json:"Points,omitempty" name:"Points"`
}

type NetDevicePortChannelEnum struct {
	// 网络设备名字

	NetdeviceName *string `json:"NetdeviceName,omitempty" name:"NetdeviceName"`
	// 网络设备端口

	PortName *string `json:"PortName,omitempty" name:"PortName"`
	// 网络设备端口光模块通道

	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
}

type NetworkPort struct {
	// 设备固资编号

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 端口名称

	PortName *string `json:"PortName,omitempty" name:"PortName"`
	// 端口IP

	PortIp *string `json:"PortIp,omitempty" name:"PortIp"`
	// 对端设备固资编号

	RemoteAssetId *string `json:"RemoteAssetId,omitempty" name:"RemoteAssetId"`
	// 对端端口名称

	RemotePortName *string `json:"RemotePortName,omitempty" name:"RemotePortName"`
	// 对端端口IP

	RemotePortIp *string `json:"RemotePortIp,omitempty" name:"RemotePortIp"`
	// 自增Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 本端设备名称

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 对端设备名称

	RemoteDeviceName *string `json:"RemoteDeviceName,omitempty" name:"RemoteDeviceName"`
}

type GetCommonServicePingDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 探测数据列表

		ChartData []*PingDataItem `json:"ChartData,omitempty" name:"ChartData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCommonServicePingDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCommonServicePingDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetdevRoleDict struct {
	// 设备角色英文

	RoleEn *string `json:"RoleEn,omitempty" name:"RoleEn"`
	// 设备角色中文

	RoleCn *string `json:"RoleCn,omitempty" name:"RoleCn"`
	// 设备角色缩写

	RoleAbstract *string `json:"RoleAbstract,omitempty" name:"RoleAbstract"`
}

type CreateIPAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// "IPAddresses": [       //多个ip地址按照此种情况进行存储返回
		//           {
		//             "IPAddress": "10.0.0.3/32",
		//             "Segment": "10.0.0.0/24"
		//          },
		//      ]

		IPAddresses []*RspIPAddress `json:"IPAddresses,omitempty" name:"IPAddresses"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIPAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIPAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateIPSegmentsRequest struct {
	*tchttp.BaseRequest

	// 更新IP网段所需信息

	IPSegments []*UpdateIPSegment `json:"IPSegments,omitempty" name:"IPSegments"`
}

func (r *UpdateIPSegmentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateIPSegmentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetDevicePortEnumSet struct {
	// 本端设备名称

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 本端端口名称

	PortName *string `json:"PortName,omitempty" name:"PortName"`
	// 对端端口名称

	RemotePortName *string `json:"RemotePortName,omitempty" name:"RemotePortName"`
	// 对端设备名称

	RemoteDeviceName *string `json:"RemoteDeviceName,omitempty" name:"RemoteDeviceName"`
	// 对端设备固资号

	RemoteAssetId *string `json:"RemoteAssetId,omitempty" name:"RemoteAssetId"`
	// 对端设备类型

	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
}

type NetDeviceConfigSet struct {
	// 网络设备资产ID

	NetDeviceAsset *string `json:"NetDeviceAsset,omitempty" name:"NetDeviceAsset"`
	// 网络设备配置名称

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
	// 网络设备配置详细信息，base64编码

	ConfigContext *string `json:"ConfigContext,omitempty" name:"ConfigContext"`
}

type DescirbeGlobalNetDeviceSyslogsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，Limit和Offset同时传入时生效

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 数量，Limit和Offset同时传入时生效

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件，支持的过滤关键字：Manufacture/ManagementIp/DeviceName/Keyword/StartTime/EndTime

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescirbeGlobalNetDeviceSyslogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescirbeGlobalNetDeviceSyslogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceRoleProportionRequest struct {
	*tchttp.BaseRequest

	// 根据过滤条件，返回指定idc的设备角色占比信息。
	// 不填查询出来的是region级别的统计信息

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDeviceRoleProportionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceRoleProportionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLanDataRequest struct {
	*tchttp.BaseRequest

	// 目标类型：la, lc

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 目标列表

	TargetList []*string `json:"TargetList,omitempty" name:"TargetList"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询指标："avg_delay", "lost_ip_count", "unreach_ip_count", "all"

	TargetTag *string `json:"TargetTag,omitempty" name:"TargetTag"`
}

func (r *GetLanDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLanDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IPSet struct {
	// IP地址

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// IP网段

	Segment *string `json:"Segment,omitempty" name:"Segment"`
	// IP类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 是否已经分配

	IsAllocated *uint64 `json:"IsAllocated,omitempty" name:"IsAllocated"`
	// 拥有者

	OwnerName *string `json:"OwnerName,omitempty" name:"OwnerName"`
	// 所属信息标识符

	ObjectId *string `json:"ObjectId,omitempty" name:"ObjectId"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
}

type SpecialIdcLineResult struct {
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// Comment

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// LineBusiness

	LineBusiness *string `json:"LineBusiness,omitempty" name:"LineBusiness"`
	// LineEndTime

	LineEndTime *string `json:"LineEndTime,omitempty" name:"LineEndTime"`
	// LineOperator

	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`
	// LineOwner

	LineOwner *string `json:"LineOwner,omitempty" name:"LineOwner"`
	// LineStartTime

	LineStartTime *string `json:"LineStartTime,omitempty" name:"LineStartTime"`
	// LocalDevice

	LocalDevice *string `json:"LocalDevice,omitempty" name:"LocalDevice"`
	// LocalDeviceIp

	LocalDeviceIp *string `json:"LocalDeviceIp,omitempty" name:"LocalDeviceIp"`
	// LocalDevicePort

	LocalDevicePort *string `json:"LocalDevicePort,omitempty" name:"LocalDevicePort"`
	// LocalOperatorOwner

	LocalOperatorOwner *string `json:"LocalOperatorOwner,omitempty" name:"LocalOperatorOwner"`
	// LocalOperatorOwnerPhone

	LocalOperatorOwnerPhone *string `json:"LocalOperatorOwnerPhone,omitempty" name:"LocalOperatorOwnerPhone"`
	// LocalOperatorWarningPhone

	LocalOperatorWarningPhone *string `json:"LocalOperatorWarningPhone,omitempty" name:"LocalOperatorWarningPhone"`
	// RemoteDevice

	RemoteDevice *string `json:"RemoteDevice,omitempty" name:"RemoteDevice"`
	// RemoteDeviceIp

	RemoteDeviceIp *string `json:"RemoteDeviceIp,omitempty" name:"RemoteDeviceIp"`
	// RemoteDevicePort

	RemoteDevicePort *string `json:"RemoteDevicePort,omitempty" name:"RemoteDevicePort"`
	// RemoteOperatorOwner

	RemoteOperatorOwner *string `json:"RemoteOperatorOwner,omitempty" name:"RemoteOperatorOwner"`
	// RemoteOperatorOwnerPhone

	RemoteOperatorOwnerPhone *string `json:"RemoteOperatorOwnerPhone,omitempty" name:"RemoteOperatorOwnerPhone"`
	// RemoteOperatorOwnerWarningPhone

	RemoteOperatorOwnerWarningPhone *string `json:"RemoteOperatorOwnerWarningPhone,omitempty" name:"RemoteOperatorOwnerWarningPhone"`
	// SlaNum

	SlaNum *string `json:"SlaNum,omitempty" name:"SlaNum"`
	// SpeLineFunc

	SpeLineFunc *int64 `json:"SpeLineFunc,omitempty" name:"SpeLineFunc"`
	// SpeLineId

	SpeLineId *string `json:"SpeLineId,omitempty" name:"SpeLineId"`
	// SpeLineName

	SpeLineName *string `json:"SpeLineName,omitempty" name:"SpeLineName"`
	// SpeLineSpeed

	SpeLineSpeed *string `json:"SpeLineSpeed,omitempty" name:"SpeLineSpeed"`
	// SpeLineStatus

	SpeLineStatus *int64 `json:"SpeLineStatus,omitempty" name:"SpeLineStatus"`
	// SpeLineType

	SpeLineType *int64 `json:"SpeLineType,omitempty" name:"SpeLineType"`
	// VlanInfo

	VlanInfo *string `json:"VlanInfo,omitempty" name:"VlanInfo"`
}

type DeleteIPAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIPAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIPAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkRoleDictResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetworkRoleDictResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetworkRoleDictResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRackExtendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRackExtendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRackExtendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcDirectConnectTopoRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为100，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 可用区名字

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

func (r *DescribeIdcDirectConnectTopoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcDirectConnectTopoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkDevicesRequest struct {
	*tchttp.BaseRequest

	// 创建网络设备信息

	CreateNetDevices []*NetworkDevice `json:"CreateNetDevices,omitempty" name:"CreateNetDevices"`
}

func (r *CreateNetworkDevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkDevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMemoryUsageTop5Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 内存利用率统计信息

		MemoryUsageStatistics []*UsageStatistic `json:"MemoryUsageStatistics,omitempty" name:"MemoryUsageStatistics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMemoryUsageTop5Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMemoryUsageTop5Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGatewaysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGatewaysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportIPAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportIPAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportIPAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDevicePortEnumRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为500。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// <li>PortName - String数组  - （过滤条件）网络设备端口名字，形如：`GigabitEthernet1/0/10`。</li>
	// <li>AssetId - String数组  - （过滤条件）网络设备资产ID，形如：`TYNT190501G7`。</li>
	// <li>DeviceName - String数组  - （过滤条件）网络设备名字，形如：`CQ-YF-201-D17-H5130-MA-01`。</li>
	// <li>NetdeviceRole - String数组  - （过滤条件）网络设备层次类型，形如：`LA`。</li>
	// <li>PortRole - String数组  - （过滤条件）网络设备端口类型，形如：`GigabitEthernet`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNetDevicePortEnumRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDevicePortEnumRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpecialDciLinePar struct {
	// 起始az

	StAz *string `json:"StAz,omitempty" name:"StAz"`
	// 终止az

	EdAz *string `json:"EdAz,omitempty" name:"EdAz"`
}

type DeleteTemplateConfigRequest struct {
	*tchttp.BaseRequest

	// 配置模板绑定关系ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteTemplateConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTemplateConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateIPCidrSegmentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateIPCidrSegmentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateIPCidrSegmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetDevicePortEnum struct {
	// 网络设备名字

	NetdeviceName *string `json:"NetdeviceName,omitempty" name:"NetdeviceName"`
	// 网络设备端口名字

	PortName *string `json:"PortName,omitempty" name:"PortName"`
}

type DescribeIdcDirectConnectMetricRequest struct {
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
	// 可用区名字

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 物理专线名字

	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`
}

func (r *DescribeIdcDirectConnectMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcDirectConnectMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IdcDirectConnectTopo struct {
	// 物理专线带宽

	DirectConnectBandwidth *string `json:"DirectConnectBandwidth,omitempty" name:"DirectConnectBandwidth"`
	// 物理专线ID

	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
	// 物理专线名字

	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`
	// 物理专线运营商

	DirectConnectOperator *string `json:"DirectConnectOperator,omitempty" name:"DirectConnectOperator"`
	// 物理专线出口网络设备IP地址

	DirectConnectSwitchIp *string `json:"DirectConnectSwitchIp,omitempty" name:"DirectConnectSwitchIp"`
	// 物理专线出口网络设备名字

	DirectConnectSwitchName *string `json:"DirectConnectSwitchName,omitempty" name:"DirectConnectSwitchName"`
	// 物理专线出口网络设备出端口

	DirectConnectSwitchPort *string `json:"DirectConnectSwitchPort,omitempty" name:"DirectConnectSwitchPort"`
}

type IdcExportLineTopo struct {
	// 专线本地网络设备系统名字

	LocalName *string `json:"LocalName,omitempty" name:"LocalName"`
	// 专线本地网络设备ip地址

	LocalIp *string `json:"LocalIp,omitempty" name:"LocalIp"`
	// 专线本地网络设备出口名字

	LocalPort []*string `json:"LocalPort,omitempty" name:"LocalPort"`
	// 专线远端网络设备ip地址

	RemoteIp *string `json:"RemoteIp,omitempty" name:"RemoteIp"`
	// 专线名字

	ExportLineName *string `json:"ExportLineName,omitempty" name:"ExportLineName"`
}

type CreateNetDeviceSyslogKeywordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetDeviceSyslogKeywordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetDeviceSyslogKeywordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopoSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopoSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTopoSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GatewayVip struct {
	// IP地址

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// IP所属类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type CreateRacksOldResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRacksOldResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRacksOldResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DevVendorProp struct {
	// 设备厂家

	DevVendor *string `json:"DevVendor,omitempty" name:"DevVendor"`
	// 占比

	Proportion *float64 `json:"Proportion,omitempty" name:"Proportion"`
	// 数量

	DevNum *int64 `json:"DevNum,omitempty" name:"DevNum"`
}

type DescribeGatewaysRequest struct {
	*tchttp.BaseRequest

	// "Filters":[
	//      {
	//       "Name":"AssetId",
	//       "Values":["TYSV190531X1", "TYSV190521X1"]
	//      },
	//      {
	//      "Name":"Cluster",
	//      "Values":["VPCGW-CLUSTER1"]
	//      }
	//     ]

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 获取条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGatewaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewaysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportNetDeviceSyslogKeywordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportNetDeviceSyslogKeywordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportNetDeviceSyslogKeywordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelatedServerInfo struct {
	// 服务器资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器设备类型

	SvrDeviceType *string `json:"SvrDeviceType,omitempty" name:"SvrDeviceType"`
	// 服务器内网IP

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
	// 服务器用途

	SvrBussiness *string `json:"SvrBussiness,omitempty" name:"SvrBussiness"`
	// 服务器机位

	SvrPosId *string `json:"SvrPosId,omitempty" name:"SvrPosId"`
	// 服务器机架高度

	SvrDeviceHeight *string `json:"SvrDeviceHeight,omitempty" name:"SvrDeviceHeight"`
	// 内网交换机端口名称

	PosInnerSwitchPortName *string `json:"PosInnerSwitchPortName,omitempty" name:"PosInnerSwitchPortName"`
	// 外网交换机端口名称

	PosOuterSwitchPortName *string `json:"PosOuterSwitchPortName,omitempty" name:"PosOuterSwitchPortName"`
	// 管理交换机端口名称

	PosAdminSwitchPortName *string `json:"PosAdminSwitchPortName,omitempty" name:"PosAdminSwitchPortName"`
}

type DeleteGatewayClustersRequest struct {
	*tchttp.BaseRequest

	// "GatewayClusters": [
	//         {
	//             "Cluster": "xxx"
	//         }
	//     ]

	GatewayClusters []*DeleteGatewayCluster `json:"GatewayClusters,omitempty" name:"GatewayClusters"`
}

func (r *DeleteGatewayClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGatewayClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkPortSortRequest struct {
	*tchttp.BaseRequest

	// 筛选条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 指标

	Target *string `json:"Target,omitempty" name:"Target"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 终止时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 当前页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 当前页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNetworkPortSortRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkPortSortRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RackPositionDevice struct {
	// 机位高度

	Height *string `json:"Height,omitempty" name:"Height"`
	// 逻辑区

	LogicalArea *string `json:"LogicalArea,omitempty" name:"LogicalArea"`
	// 管理交换机端口

	PosAdminSwitchPortName *string `json:"PosAdminSwitchPortName,omitempty" name:"PosAdminSwitchPortName"`
	// 内网交换机端口

	PosInnerSwitchPortName *string `json:"PosInnerSwitchPortName,omitempty" name:"PosInnerSwitchPortName"`
	// 外网交换机端口

	PosOuterSwitchPortName *string `json:"PosOuterSwitchPortName,omitempty" name:"PosOuterSwitchPortName"`
	// 机位占用状态

	PositionStatus *int64 `json:"PositionStatus,omitempty" name:"PositionStatus"`
	// 服务器固资号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器所在业务

	SvrBussiness *string `json:"SvrBussiness,omitempty" name:"SvrBussiness"`
	// 服务器高度

	SvrDeviceHeight *string `json:"SvrDeviceHeight,omitempty" name:"SvrDeviceHeight"`
	// 服务器类型

	SvrDeviceType *string `json:"SvrDeviceType,omitempty" name:"SvrDeviceType"`
	// 服务器内网IP

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
	// 机位号

	SvrPosId *string `json:"SvrPosId,omitempty" name:"SvrPosId"`
	// 机位创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 机位更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type CreateTemplateConfigRequest struct {
	*tchttp.BaseRequest

	// 配置模板绑定信息

	TemplateConfig *TemplateConfig `json:"TemplateConfig,omitempty" name:"TemplateConfig"`
}

func (r *CreateTemplateConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTemplateConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkRoleDictRequest struct {
	*tchttp.BaseRequest

	// 根据过滤条件删除字典条目，"Filters":[
	//         {
	//             "Name":"RoleEn",
	//             "Values":["BMS-LA-25G"]
	//         }
	//     ]

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DeleteNetworkRoleDictRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkRoleDictRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRackPositionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回机架信息

		RackPositionSet []*RackPositionInfo `json:"RackPositionSet,omitempty" name:"RackPositionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRackPositionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRackPositionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRackExtendRequest struct {
	*tchttp.BaseRequest

	// az

	Az *string `json:"Az,omitempty" name:"Az"`
	// idc

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 机柜

	Rack *string `json:"Rack,omitempty" name:"Rack"`
}

func (r *DeleteRackExtendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRackExtendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRackPositionsRequest struct {
	*tchttp.BaseRequest

	// 查询过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回记录偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRackPositionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRackPositionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrapAlarmAbstractResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SNMP TRAP概要告警对象数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// SNMP TRAP概要告警对象数组

		TrapAlarmAbstractSet []*TrapAlarmAbstract `json:"TrapAlarmAbstractSet,omitempty" name:"TrapAlarmAbstractSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrapAlarmAbstractResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrapAlarmAbstractResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGatewayCluster struct {
	// 网关集群名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

type NetDeviceSlotList struct {
	// 网络设备名字

	NetdeviceName *string `json:"NetdeviceName,omitempty" name:"NetdeviceName"`
	// 网络设备单板列表

	SlotList []*string `json:"SlotList,omitempty" name:"SlotList"`
}

type DeleteNetworkPortsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNetworkPortsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkPortsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateIPCidrSegmentsRequest struct {
	*tchttp.BaseRequest

	// IPCIDR网段信息

	IPCIDRSegmentSet []*IPCIDRSegment `json:"IPCIDRSegmentSet,omitempty" name:"IPCIDRSegmentSet"`
}

func (r *UpdateIPCidrSegmentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateIPCidrSegmentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalSyslogs struct {
	// 网络设备名称列表

	DeviceName []*string `json:"DeviceName,omitempty" name:"DeviceName"`
	// 网络设备管理IP

	ManagementIp *string `json:"ManagementIp,omitempty" name:"ManagementIp"`
	// 最近发生的日志时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 匹配的日志总量

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 网络设备日志详细信息

	Syslogs []*Syslogs `json:"Syslogs,omitempty" name:"Syslogs"`
	// 网络设备资产Id列表

	AssetId []*string `json:"AssetId,omitempty" name:"AssetId"`
}

type DciCapacityRequest struct {
	*tchttp.BaseRequest

	// dci容量查询参数

	SpecialDciLineInfo *SpecialDciLineInfo `json:"SpecialDciLineInfo,omitempty" name:"SpecialDciLineInfo"`
}

func (r *DciCapacityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DciCapacityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPSegmentFilterRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeIPSegmentFilterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPSegmentFilterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrapAlarmResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SNMP TRAP明细告警对象数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// SNMP TRAP明细告警对象数目

		TrapAlarmSet []*TrapAlarm `json:"TrapAlarmSet,omitempty" name:"TrapAlarmSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrapAlarmResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrapAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGwConfigRequest struct {
	*tchttp.BaseRequest

	// 网关类型：vpcgw, tgw...

	GwType *string `json:"GwType,omitempty" name:"GwType"`
	// 探测类型，默认为ping, 可选值["ping", "l4", "l7", "lan_l4", "lan_l7", "wan_l4", "wan_l7"]

	ProbeType *string `json:"ProbeType,omitempty" name:"ProbeType"`
}

func (r *GetGwConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGwConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DciLine struct {
	// 探测源地域

	SrcArea *string `json:"SrcArea,omitempty" name:"SrcArea"`
	// 探测目标地域

	DstArea *string `json:"DstArea,omitempty" name:"DstArea"`
}

type CreateGatewayClustersRequest struct {
	*tchttp.BaseRequest

	// "GatewayClusters": [
	//         {
	//             "Cluster": "xxx",   //Cluster名称
	//             "Gateways": "xxxx", //Cluster中所包含的网关机器信息，json字符串
	//             "Product": "xxxx", //网关所属产品分类
	//             "Vip": “”, //已规划的虚拟IP信息，json字符串
	//             "ClusterExpand": "xxx", //集群扩展信息
	//             "Description": "xxx"
	//         }
	//     ]

	GatewayClusters []*AddGatewayCluster `json:"GatewayClusters,omitempty" name:"GatewayClusters"`
}

func (r *CreateGatewayClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGatewayClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkPortSortResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本端设备和对端设备信息

		NetDevicePortEnumSet []*NetDevicePortEnumSet `json:"NetDevicePortEnumSet,omitempty" name:"NetDevicePortEnumSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkPortSortResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkPortSortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGatewaysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGatewaysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkDevicesExRequest struct {
	*tchttp.BaseRequest

	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，Limit和Offset同时传入时生效

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移，Limit和Offset同时传入时生效

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNetworkDevicesExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkDevicesExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNetDeviceSyslogKeywordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateNetDeviceSyslogKeywordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNetDeviceSyslogKeywordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRackExtendRequest struct {
	*tchttp.BaseRequest

	// az

	Az *string `json:"Az,omitempty" name:"Az"`
	// idc

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 机柜

	Rack *string `json:"Rack,omitempty" name:"Rack"`
	// 是否通电（0：未通电，1：通电）

	PowerStatus *int64 `json:"PowerStatus,omitempty" name:"PowerStatus"`
	// 机柜类型（NETWORK：网络柜，SERVER：服务器柜）

	RackType *string `json:"RackType,omitempty" name:"RackType"`
}

func (r *UpdateRackExtendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRackExtendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIPAddress struct {
	// IP地址

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 根据网段删除

	Segment *string `json:"Segment,omitempty" name:"Segment"`
}

type DetailPingDataItem struct {
	// 探测时间

	PingTime *string `json:"PingTime,omitempty" name:"PingTime"`
	// 探测源IP

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 探测目标IP

	PingIp *string `json:"PingIp,omitempty" name:"PingIp"`
	// 时延

	PingAvg *float64 `json:"PingAvg,omitempty" name:"PingAvg"`
	// 丢包率

	PingLost *float64 `json:"PingLost,omitempty" name:"PingLost"`
	// 是否可达，1：可达，0：不可达

	Reachable *int64 `json:"Reachable,omitempty" name:"Reachable"`
}

type DescribeNetdeviceFieldsEnumResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回枚举结果

		Field2Enum *Field2Enum `json:"Field2Enum,omitempty" name:"Field2Enum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetdeviceFieldsEnumResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetdeviceFieldsEnumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcDirectConnectTopoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的物理专线拓扑集对象数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 符合条件的物理专线拓扑集对象数组

		IdcDirectConnectTopoSet []*IdcDirectConnectTopo `json:"IdcDirectConnectTopoSet,omitempty" name:"IdcDirectConnectTopoSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdcDirectConnectTopoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcDirectConnectTopoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDevicePortListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络设备端口列表对象数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 网络设备端口列表对象数组

		NetDevicePortListSet []*NetDevicePortList `json:"NetDevicePortListSet,omitempty" name:"NetDevicePortListSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetDevicePortListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDevicePortListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WithdrawNetworkDevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WithdrawNetworkDevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WithdrawNetworkDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Dimensions struct {
	// 设备名称

	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
	// slotid

	SlotId *string `json:"SlotId,omitempty" name:"SlotId"`
	// 可用区

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 出口名称

	ExportLineName *string `json:"ExportLineName,omitempty" name:"ExportLineName"`
	// 专线名称

	SpeLineName *string `json:"SpeLineName,omitempty" name:"SpeLineName"`
	// 物理专线名字

	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`
}

type CreateIPCIDRSegmentsRequest struct {
	*tchttp.BaseRequest

	// IPCIDR信息

	IPCIDRSegmentSet []*IPCIDRSegment `json:"IPCIDRSegmentSet,omitempty" name:"IPCIDRSegmentSet"`
}

func (r *CreateIPCIDRSegmentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIPCIDRSegmentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCPUUsageTop5Request struct {
	*tchttp.BaseRequest

	// 根据过滤条件 "Filters":[
	//         {
	//             "Name":"Idc",
	//             "Values":["YF-1"]
	//         }
	//     ]返回指定idc的cup 利用率top5数据。Filters为空表示返回整个region的cpu利用率top5。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCPUUsageTop5Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCPUUsageTop5Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFieldsEnumRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeFieldsEnumRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFieldsEnumRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcExportLineMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监控指标

		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
		// 数据点起始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 数据点结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 数据统计周期

		Period *uint64 `json:"Period,omitempty" name:"Period"`
		// 监控数据列表

		DataPoints []*DataPoints `json:"DataPoints,omitempty" name:"DataPoints"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdcExportLineMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcExportLineMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWanConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetWanConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWanConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTemplateConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置模板绑定对象数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 配置模板绑定对象数组

		TemplateConfigResultSet *TemplateConfigResult `json:"TemplateConfigResultSet,omitempty" name:"TemplateConfigResultSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTemplateConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTemplateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGwDetailDataRequest struct {
	*tchttp.BaseRequest

	// 网关类型：tgw, vpcgw

	GwType *string `json:"GwType,omitempty" name:"GwType"`
	// 查询目标

	TargetName *string `json:"TargetName,omitempty" name:"TargetName"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// ip类型：vip, rip, tsv_ip,ld_ip...

	IpType *string `json:"IpType,omitempty" name:"IpType"`
	// 探测类型，默认为ping, 可选值["ping", "l4", "l7", "lan_l4", "lan_l7", "wan_l4", "wan_l7"]

	ProbeType *string `json:"ProbeType,omitempty" name:"ProbeType"`
}

func (r *GetGwDetailDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGwDetailDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetDeviceSlotEnum struct {
	// 网络设备名字

	NetdeviceName *string `json:"NetdeviceName,omitempty" name:"NetdeviceName"`
	// 网络设备单板ID

	SlotId *string `json:"SlotId,omitempty" name:"SlotId"`
}

type Operationlog struct {
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 操作类型，增删改查

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 操作资源，如IP

	Resource *string `json:"Resource,omitempty" name:"Resource"`
	// 操作人员

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 操作的对象

	Object *string `json:"Object,omitempty" name:"Object"`
	// 操作结果

	Result *string `json:"Result,omitempty" name:"Result"`
	// 操作详细描述

	Detail *string `json:"Detail,omitempty" name:"Detail"`
}

type UpdateGatewayCluster struct {
	// 集群名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 集群扩展信息

	ClusterExpand *string `json:"ClusterExpand,omitempty" name:"ClusterExpand"`
	// 网关产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateNetworkPortsRequest struct {
	*tchttp.BaseRequest

	// 创建网络设备端口信息

	CreateNetPorts []*NetworkPort `json:"CreateNetPorts,omitempty" name:"CreateNetPorts"`
}

func (r *CreateNetworkPortsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkPortsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGatewayClustersRequest struct {
	*tchttp.BaseRequest

	// "GatewayClusters": [
	//         {
	//             "Cluster": "xxx",   //Cluster名称，用于标识更新哪个cluster
	//             "ClusterExpand": "xxx", //集群扩展信息
	//             "Product":"SXGW",
	//             "Description": "xxx"
	//         }
	//     ]

	GatewayClusters []*UpdateGatewayCluster `json:"GatewayClusters,omitempty" name:"GatewayClusters"`
}

func (r *UpdateGatewayClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateGatewayClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkDevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回网络设备信息

		NetDeviceSet []*NetworkDevice `json:"NetDeviceSet,omitempty" name:"NetDeviceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkDevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 条数限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDeviceRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为100，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 可用区名字。

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 过滤条件。
	// <li>AssetID - String - （过滤条件）设备固资号，形如：`TYNT190501T9`。</li>
	// <li>Role - String - （过滤条件）设备层次角色名字，形如：`PL`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNetDeviceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDeviceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGatewaysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGatewaysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IdcExportLineResult struct {
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// ExportLineId

	ExportLineId *string `json:"ExportLineId,omitempty" name:"ExportLineId"`
	// ExportLineName

	ExportLineName *string `json:"ExportLineName,omitempty" name:"ExportLineName"`
	// ExportLineStatus

	ExportLineStatus *int64 `json:"ExportLineStatus,omitempty" name:"ExportLineStatus"`
	// LineTechAvailableSpeed

	LineTechAvailableSpeed *string `json:"LineTechAvailableSpeed,omitempty" name:"LineTechAvailableSpeed"`
	// LineBusiAvailableSpeed

	LineBusiAvailableSpeed *string `json:"LineBusiAvailableSpeed,omitempty" name:"LineBusiAvailableSpeed"`
	// ExportDeviceIp

	ExportDeviceIp *string `json:"ExportDeviceIp,omitempty" name:"ExportDeviceIp"`
	// RemoteDeviceIp

	RemoteDeviceIp *string `json:"RemoteDeviceIp,omitempty" name:"RemoteDeviceIp"`
	// LineOperator

	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`
	// RemoteOperatorOwner

	RemoteOperatorOwner *string `json:"RemoteOperatorOwner,omitempty" name:"RemoteOperatorOwner"`
	// RemoteOperatorOwnerPhone

	RemoteOperatorOwnerPhone *string `json:"RemoteOperatorOwnerPhone,omitempty" name:"RemoteOperatorOwnerPhone"`
	// RemoteOperatorOwnerWarningPhone

	RemoteOperatorOwnerWarningPhone *string `json:"RemoteOperatorOwnerWarningPhone,omitempty" name:"RemoteOperatorOwnerWarningPhone"`
	// SlaNum

	SlaNum *string `json:"SlaNum,omitempty" name:"SlaNum"`
	// Comment

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// ExportDevice

	ExportDevice *string `json:"ExportDevice,omitempty" name:"ExportDevice"`
	// ExportDevicePort

	ExportDevicePort []*string `json:"ExportDevicePort,omitempty" name:"ExportDevicePort"`
	// LineStartTime

	LineStartTime *string `json:"LineStartTime,omitempty" name:"LineStartTime"`
	// LineEndTime

	LineEndTime *string `json:"LineEndTime,omitempty" name:"LineEndTime"`
	// ExportDeviceAndPort

	ExportDeviceAndPort []*string `json:"ExportDeviceAndPort,omitempty" name:"ExportDeviceAndPort"`
}

type DescribeIdcDirectConnectMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标名称

		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
		// 监控统计周期

		Period *uint64 `json:"Period,omitempty" name:"Period"`
		// 起始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 监控数据列表

		DataPoints []*DataPoints `json:"DataPoints,omitempty" name:"DataPoints"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdcDirectConnectMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcDirectConnectMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPCIDRSegmentsRequest struct {
	*tchttp.BaseRequest

	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件："Filters":[
	//         {
	//             "Name":"AddressType",
	//             "Values":["Overlay"]
	//         },
	//         {
	//             "Name":"Version",
	//             "Values":["IPv4"]
	//         },
	//         {
	//             "Name":"AddrSegment",
	//             "Values":["20.00.00.00/8"]
	//         },
	//         {
	//             "Name":"IsReserved",
	//             "Values":["True"]
	//         }
	//     ]

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeIPCIDRSegmentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPCIDRSegmentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWanDetailDataRequest struct {
	*tchttp.BaseRequest

	// 省份

	Province *string `json:"Province,omitempty" name:"Province"`
	// 运营商

	Isp *string `json:"Isp,omitempty" name:"Isp"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *GetWanDetailDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWanDetailDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IPCIDRSegment struct {
	// 地址类型：Overlay或者Underlay

	AddressType *string `json:"AddressType,omitempty" name:"AddressType"`
	// IP版本：IPv4 （后续可以扩展IPv6）

	Version *string `json:"Version,omitempty" name:"Version"`
	// 格式如：20.00.00.00/8

	AddrSegment *string `json:"AddrSegment,omitempty" name:"AddrSegment"`
	// 是否系统保留网段

	IsReserved *bool `json:"IsReserved,omitempty" name:"IsReserved"`
	// 白名单key

	WhiteKey *string `json:"WhiteKey,omitempty" name:"WhiteKey"`
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 网段用途

	UsageType *int64 `json:"UsageType,omitempty" name:"UsageType"`
}

type SyslogKeywords struct {
	// 策略

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 日志告警类型

	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`
	// 网络设备厂商

	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`
	// 网络设备类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 网络设备角色

	Role *string `json:"Role,omitempty" name:"Role"`
	// 网络设备型号

	Model *string `json:"Model,omitempty" name:"Model"`
	// 说明

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type TemplateResult struct {
	// 配置模板ID

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 配置模板描述

	TemplateDesc *string `json:"TemplateDesc,omitempty" name:"TemplateDesc"`
	// 配置模板名字

	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
	// 配置模板导入时间

	Time *string `json:"Time,omitempty" name:"Time"`
}

type DescribeAlarmStatisticsRequest struct {
	*tchttp.BaseRequest

	// 根据过滤条件 "Filters":[
	//         {
	//             "Name":"Idc",
	//             "Values":["YF-1"]
	//         }
	//     ]返回指定idc未处理告警总数。如果Filters为空表示返回整个region的未处理告警总数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAlarmStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDevicePortChannelEnumResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络设备端口光模块通道枚举对象数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 网络设备端口光模块通道枚举对象数目

		NetDevicePortChannelEnumSet []*NetDevicePortChannelEnum `json:"NetDevicePortChannelEnumSet,omitempty" name:"NetDevicePortChannelEnumSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetDevicePortChannelEnumResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDevicePortChannelEnumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetdeviceRelatedServersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 关联服务器信息

		RelatedServersSet []*RelatedServerInfo `json:"RelatedServersSet,omitempty" name:"RelatedServersSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetdeviceRelatedServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetdeviceRelatedServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDetailPingDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 探测数据明细列表

		DetailData []*DetailPingDataItem `json:"DetailData,omitempty" name:"DetailData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDetailPingDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDetailPingDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTemplateRequest struct {
	*tchttp.BaseRequest

	// 配置模板ID

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashboardInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportSpecialIdcLineFromDcosToNmsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 同步IDC专线数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportSpecialIdcLineFromDcosToNmsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportSpecialIdcLineFromDcosToNmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPCIDRSegmentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// IPCIDR网段信息

		IPCIDRSegmentSet []*IPCIDRSegment `json:"IPCIDRSegmentSet,omitempty" name:"IPCIDRSegmentSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIPCIDRSegmentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPCIDRSegmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPSegmentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取到的网段数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 获取到的网段列表

		SegmentSet []*SegmentSet `json:"SegmentSet,omitempty" name:"SegmentSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIPSegmentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIPSegmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGwDataRequest struct {
	*tchttp.BaseRequest

	// 网关类型：vpcgw, tgw...

	GwType *string `json:"GwType,omitempty" name:"GwType"`
	// IP类型：vip, ld_ip, tsv_ip....

	IpType *string `json:"IpType,omitempty" name:"IpType"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询指标："avg_delay", "lost_ip_count", "unreach_ip_count", "all"

	TargetTag *string `json:"TargetTag,omitempty" name:"TargetTag"`
	// 查询目标名称

	TargetList []*string `json:"TargetList,omitempty" name:"TargetList"`
	// 探测类型，默认为ping, 可选值["ping", "l4", "l7", "lan_l4", "lan_l7", "wan_l4", "wan_l7"]

	ProbeType *string `json:"ProbeType,omitempty" name:"ProbeType"`
}

func (r *GetGwDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGwDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Idc struct {
	// Idc名称或编号

	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
}

type ImportIpAddresses struct {
	// IP所属网段

	Segment *string `json:"Segment,omitempty" name:"Segment"`
	// IP地址

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 是否已经分配

	IsAllocated *int64 `json:"IsAllocated,omitempty" name:"IsAllocated"`
	// 使用者

	OwnerName *string `json:"OwnerName,omitempty" name:"OwnerName"`
	// 类型，SE,SE-Spical，PXE等

	Type *string `json:"Type,omitempty" name:"Type"`
	// 关联对象

	ObjectId *string `json:"ObjectId,omitempty" name:"ObjectId"`
	// 描述信息

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type DescribeDashboardInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDashboardInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IPSegmentFilter struct {
	// 可用区列表

	Az []*string `json:"Az,omitempty" name:"Az"`
	// 机房列表

	Idc []*string `json:"Idc,omitempty" name:"Idc"`
	// 逻辑分区

	LogicalArea []*string `json:"LogicalArea,omitempty" name:"LogicalArea"`
	// 网段分类

	SegmentUsage []*string `json:"SegmentUsage,omitempty" name:"SegmentUsage"`
}

type DescribeSpecialIdcLineTopoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 专线拓扑连接对象数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 专线拓扑连接关系数组

		SpecialIdcLineTopoSet []*SpecialIdcLineTopo `json:"SpecialIdcLineTopoSet,omitempty" name:"SpecialIdcLineTopoSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSpecialIdcLineTopoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpecialIdcLineTopoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopoSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络拓扑位置配置

		Setting *string `json:"Setting,omitempty" name:"Setting"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopoSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopoSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGatewaysRequest struct {
	*tchttp.BaseRequest

	// "Gateways": [
	//         {
	//             "AssetId": "xxxx", //需要添加的网关资产id
	//             "InnerIp": "xxxx",
	//             "OuterIp": “xxx”
	//         }
	//     ]

	Gateways []*AddGateway `json:"Gateways,omitempty" name:"Gateways"`
}

func (r *UpdateGatewaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateGatewaysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceVendorProportionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 设备厂家占比信息

		DevVendorPropSet []*DevVendorProp `json:"DevVendorPropSet,omitempty" name:"DevVendorPropSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceVendorProportionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceVendorProportionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDevicePortStateRequest struct {
	*tchttp.BaseRequest

	// 网络设备固资号

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为500，最大值为500。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// <li>PortName - String数组  - （过滤条件）网络设备端口名字，形如：`GigabitEthernet1/0/10`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNetDevicePortStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDevicePortStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetdeviceRelatedServersRequest struct {
	*tchttp.BaseRequest

	// 网络设备资产ID

	NetdevAssetId *string `json:"NetdevAssetId,omitempty" name:"NetdevAssetId"`
}

func (r *DescribeNetdeviceRelatedServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetdeviceRelatedServersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopoSettingRequest struct {
	*tchttp.BaseRequest

	// 可用区

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

func (r *DescribeTopoSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopoSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SingleIPAddress struct {
	// IP地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 所属对象，交换机或者其他标识

	ObjectId *string `json:"ObjectId,omitempty" name:"ObjectId"`
	// 申请人

	OwnerName *string `json:"OwnerName,omitempty" name:"OwnerName"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// IP类型信息

	Type *string `json:"Type,omitempty" name:"Type"`
}
