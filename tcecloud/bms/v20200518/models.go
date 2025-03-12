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

package v20200518

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DescribeBmsNodesExRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 可用区域id

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区域名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 是否已售卖的标识

	Allocated *bool `json:"Allocated,omitempty" name:"Allocated"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// SN号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 机型类型：0=标准机型，1=自定义机型。

	ModelType *uint64 `json:"ModelType,omitempty" name:"ModelType"`
	// 机型名称。

	BoundType *string `json:"BoundType,omitempty" name:"BoundType"`
	// 固资号

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// CPU型号。

	CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
}

func (r *DescribeBmsNodesExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBmsNodesExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeObmsPoolNodesToAddExRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 可用区域ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 地域名

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 可用区域名

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 服务器序列号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 服务器内网IP

	UnderlayIp *string `json:"UnderlayIp,omitempty" name:"UnderlayIp"`
	// 服务器机型

	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
	// CPU型号：X86/ARM

	CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
	// 机器的网卡数筛选参数，用于4网口过滤

	NetworkPorts *int64 `json:"NetworkPorts,omitempty" name:"NetworkPorts"`
}

func (r *DescribeObmsPoolNodesToAddExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeObmsPoolNodesToAddExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDefaultQuotaRequest struct {
	*tchttp.BaseRequest

	// 需要更新的配额资源类型，默认为`Bare Metal Server`

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 默认配额值

	DefaultQuota *uint64 `json:"DefaultQuota,omitempty" name:"DefaultQuota"`
	// 默认的最大配额值

	MaxQuota *uint64 `json:"MaxQuota,omitempty" name:"MaxQuota"`
}

func (r *UpdateDefaultQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDefaultQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddNodePlacement struct {
	// Region

	Region *int64 `json:"Region,omitempty" name:"Region"`
	// Zone

	Zone *int64 `json:"Zone,omitempty" name:"Zone"`
}

type Os struct {
	// Linux系统

	Linux []*string `json:"Linux,omitempty" name:"Linux"`
	// Windows系统

	Windows []*string `json:"Windows,omitempty" name:"Windows"`
	// 其他操作系统

	Other []*string `json:"Other,omitempty" name:"Other"`
}

type AddFlavorExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddFlavorExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddFlavorExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeObmsNodeFilterExRequest struct {
	*tchttp.BaseRequest

	// 可用区域ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 是否售卖的标识

	Allocated *bool `json:"Allocated,omitempty" name:"Allocated"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *DescribeObmsNodeFilterExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeObmsNodeFilterExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskExRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 创建BMS用户的AppId

	InnerAppId *string `json:"InnerAppId,omitempty" name:"InnerAppId"`
}

func (r *DescribeTaskExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Disk struct {
	// 硬盘名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 硬盘容量

	Size *string `json:"Size,omitempty" name:"Size"`
	// 硬盘sn值

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 硬盘类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type UpdateNodeInfo struct {
	// BMS服务器名

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 电源状态

	PowerState *string `json:"PowerState,omitempty" name:"PowerState"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 售卖状态

	SellState *string `json:"SellState,omitempty" name:"SellState"`
	// 区域

	Region *int64 `json:"Region,omitempty" name:"Region"`
	// 类型

	IloIp *string `json:"IloIp,omitempty" name:"IloIp"`
	// overlay ip地址

	FlavorId *string `json:"FlavorId,omitempty" name:"FlavorId"`
	// SN编码

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 可用区ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// BGP接口名称

	NetInfo *string `json:"NetInfo,omitempty" name:"NetInfo"`
	// UnderlayMask

	UnderlayMask *string `json:"UnderlayMask,omitempty" name:"UnderlayMask"`
	// Mac地址

	Mac *string `json:"Mac,omitempty" name:"Mac"`
	// UnderlayGateway

	UnderlayGateway *string `json:"UnderlayGateway,omitempty" name:"UnderlayGateway"`
	// 上联交换机名称

	SwitchPortPhyName *string `json:"SwitchPortPhyName,omitempty" name:"SwitchPortPhyName"`
	// 上联交换机端口连接的端口号

	SwitchPortName *string `json:"SwitchPortName,omitempty" name:"SwitchPortName"`
	// 上联交换机Sn号

	SwitchSn *string `json:"SwitchSn,omitempty" name:"SwitchSn"`
}

type DescribeBmsNodesExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的裸金属服务器总个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 裸金属服务器信息

		BmsNode []*BmsNode `json:"BmsNode,omitempty" name:"BmsNode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBmsNodesExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBmsNodesExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建任务详细

		Output *Output `json:"Output,omitempty" name:"Output"`
		// 任务状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Placement struct {
	// 可用区名称

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type UpdateSwitchInfo struct {
	// 网络自治系统数量

	AsNum *string `json:"AsNum,omitempty" name:"AsNum"`
	// 网关边界协议类型

	BgpIfType *string `json:"BgpIfType,omitempty" name:"BgpIfType"`
	// 网关边界协议名称

	BgpIfName *string `json:"BgpIfName,omitempty" name:"BgpIfName"`
	// 可用区ID

	Zone *uint64 `json:"Zone,omitempty" name:"Zone"`
	// 固资号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 交换机业务IP

	VtepIp *string `json:"VtepIp,omitempty" name:"VtepIp"`
	// 网关端点名称

	Role *string `json:"Role,omitempty" name:"Role"`
	// 可用地域ID

	Region *uint64 `json:"Region,omitempty" name:"Region"`
	// 产品型号

	Product *string `json:"Product,omitempty" name:"Product"`
	// 设备名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 厂商

	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`
	// 主IP

	AdminIp *string `json:"AdminIp,omitempty" name:"AdminIp"`
	// 交换机管理端口号

	AdminPort *string `json:"AdminPort,omitempty" name:"AdminPort"`
}

type DescribeNodeExRequest struct {
	*tchttp.BaseRequest

	// 固资号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 机型ID

	FlavorId *string `json:"FlavorId,omitempty" name:"FlavorId"`
	// 分页起始位置，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页返回的条目数，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNodeExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSellstateExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json格式的SellstateInfo

		SellstateInfo *string `json:"SellstateInfo,omitempty" name:"SellstateInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSellstateExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSellstateExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNodeInfo struct {
	// 可用地域id

	Region *uint64 `json:"Region,omitempty" name:"Region"`
	// 可用区域id

	Zone *uint64 `json:"Zone,omitempty" name:"Zone"`
	// 固资号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
}

type EditSwitchInfo struct {
	// 网络自治系统数量

	AsNum *string `json:"AsNum,omitempty" name:"AsNum"`
	// 网关边界协议类型

	BgpIfType *string `json:"BgpIfType,omitempty" name:"BgpIfType"`
	// 网关边界协议名称

	BgpIfName *string `json:"BgpIfName,omitempty" name:"BgpIfName"`
	// 可用区ID

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 固资号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 交换机业务IP

	VtepIp *string `json:"VtepIp,omitempty" name:"VtepIp"`
	// 网关端点名称

	Role *string `json:"Role,omitempty" name:"Role"`
	// 可用地域ID

	Region *string `json:"Region,omitempty" name:"Region"`
	// 产品型号

	Product *string `json:"Product,omitempty" name:"Product"`
	// 设备名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 厂商

	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`
	// 主IP

	AdminIp *string `json:"AdminIp,omitempty" name:"AdminIp"`
}

type ReturnNodeToObmsPoolInfo struct {
	// 初始字段

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统键值

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 执行动作

	Op *string `json:"Op,omitempty" name:"Op"`
	// 机器信息

	ServersInfo []*ServersInfo `json:"ServersInfo,omitempty" name:"ServersInfo"`
}

type FlavorSet struct {
	// 磁盘阵列

	Raid []*string `json:"Raid,omitempty" name:"Raid"`
	// 机型名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否用户自定义机型（1：是，0：否）

	UserDefined *uint64 `json:"UserDefined,omitempty" name:"UserDefined"`
	// 网卡速度

	NetSpeed *string `json:"NetSpeed,omitempty" name:"NetSpeed"`
	// 处理器

	Cpu *string `json:"Cpu,omitempty" name:"Cpu"`
	// 是否售罄

	SoldOutMode *string `json:"SoldOutMode,omitempty" name:"SoldOutMode"`
	// 租户端是否用户可见

	IsShow *uint64 `json:"IsShow,omitempty" name:"IsShow"`
	// 此机型绑定的白名单

	WhiteType *string `json:"WhiteType,omitempty" name:"WhiteType"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 内存大小

	Memory *string `json:"Memory,omitempty" name:"Memory"`
	// 地域

	Region *uint64 `json:"Region,omitempty" name:"Region"`
	// 硬盘大小

	Disk *string `json:"Disk,omitempty" name:"Disk"`
	// 操作系统

	Os *Os `json:"Os,omitempty" name:"Os"`
	// 机型ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 可用区

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 是否可用

	Available *uint64 `json:"Available,omitempty" name:"Available"`
	// 已售卖个数

	Sold *uint64 `json:"Sold,omitempty" name:"Sold"`
	// 绑定此机型的BMS服务器个数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
	// 机型类型

	FlavorType *string `json:"FlavorType,omitempty" name:"FlavorType"`
	// CPU架构：X86/ARM

	CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
	// 机型网口数

	NetworkPorts *int64 `json:"NetworkPorts,omitempty" name:"NetworkPorts"`
	// dcos后置脚本

	DcosShell *string `json:"DcosShell,omitempty" name:"DcosShell"`
}

type AddNodeToFlavor struct {
	// 端口

	AggregationPort *string `json:"AggregationPort,omitempty" name:"AggregationPort"`
	// bonding标志

	BondingFlag *string `json:"BondingFlag,omitempty" name:"BondingFlag"`
	// 硬盘

	Disk []*Disk `json:"Disk,omitempty" name:"Disk"`
	// 是否是用户定义标识

	IsUserDefineModel *string `json:"IsUserDefineModel,omitempty" name:"IsUserDefineModel"`
	// mac地址

	Mac *string `json:"Mac,omitempty" name:"Mac"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 网关端口名称

	PosInnerSwitchPortName *string `json:"PosInnerSwitchPortName,omitempty" name:"PosInnerSwitchPortName"`
	// 可用地域

	Region *int64 `json:"Region,omitempty" name:"Region"`
	// 可用地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 固资号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 服务器主网关id

	SvrAdminSwitchAssetid *string `json:"SvrAdminSwitchAssetid,omitempty" name:"SvrAdminSwitchAssetid"`
	// 服务器网关版本

	SvrAgentVersion *string `json:"SvrAgentVersion,omitempty" name:"SvrAgentVersion"`
	// 服务器产品id

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器belond产品

	SvrBelongingProducts *string `json:"SvrBelongingProducts,omitempty" name:"SvrBelongingProducts"`
	// 服务器信息

	SvrBussiness *string `json:"SvrBussiness,omitempty" name:"SvrBussiness"`
	// 服务器商用id

	SvrBussinessId *string `json:"SvrBussinessId,omitempty" name:"SvrBussinessId"`
	// 服务器当前状态

	SvrCurrentStatus *string `json:"SvrCurrentStatus,omitempty" name:"SvrCurrentStatus"`
	// 服务器设备描述

	SvrDeviceDescript *string `json:"SvrDeviceDescript,omitempty" name:"SvrDeviceDescript"`
	// 服务器设备信息

	SvrDeviceHeight *string `json:"SvrDeviceHeight,omitempty" name:"SvrDeviceHeight"`
	// 服务器设备名称

	SvrDeviceName *string `json:"SvrDeviceName,omitempty" name:"SvrDeviceName"`
	// 服务器设备类型

	SvrDeviceType *string `json:"SvrDeviceType,omitempty" name:"SvrDeviceType"`
	// 服务器首次使用时间

	SvrFirstUseTime *string `json:"SvrFirstUseTime,omitempty" name:"SvrFirstUseTime"`
	// 服务器id

	SvrId *string `json:"SvrId,omitempty" name:"SvrId"`
	// 服务器专线id

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// 服务器专线名称

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// 服务器插入时间

	SvrInsertTime *string `json:"SvrInsertTime,omitempty" name:"SvrInsertTime"`
	// 服务器lanip

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
	// 服务器lan网关assetid

	SvrLanSwitchAssetid *string `json:"SvrLanSwitchAssetid,omitempty" name:"SvrLanSwitchAssetid"`
	// 服务器登录区

	SvrLogicArea *string `json:"SvrLogicArea,omitempty" name:"SvrLogicArea"`
	// 服务器系统名称

	SvrOsName *string `json:"SvrOsName,omitempty" name:"SvrOsName"`
	// 服务器posid

	SvrPosId *string `json:"SvrPosId,omitempty" name:"SvrPosId"`
	// 服务器产品

	SvrProducer *string `json:"SvrProducer,omitempty" name:"SvrProducer"`
	// 服务器rack名称

	SvrRackName *string `json:"SvrRackName,omitempty" name:"SvrRackName"`
	// 服务器raid类型

	SvrRaidType *string `json:"SvrRaidType,omitempty" name:"SvrRaidType"`
	// 服务器固资号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 服务器类型

	SvrType *string `json:"SvrType,omitempty" name:"SvrType"`
	// 服务器虚拟序列

	SvrVmIndex *string `json:"SvrVmIndex,omitempty" name:"SvrVmIndex"`
	// 服务器waneip

	SvrWanEip *string `json:"SvrWanEip,omitempty" name:"SvrWanEip"`
	// 服务器wanIP

	SvrWanIp *string `json:"SvrWanIp,omitempty" name:"SvrWanIp"`
	// 服务器网关assetid

	SvrWanSwitchAssetid *string `json:"SvrWanSwitchAssetid,omitempty" name:"SvrWanSwitchAssetid"`
	// 网关

	Switch []*Switch `json:"Switch,omitempty" name:"Switch"`
	// 网关端口名称

	SwitchPortName *string `json:"SwitchPortName,omitempty" name:"SwitchPortName"`
	// 网关端口物理名称

	SwitchPortPhyName *string `json:"SwitchPortPhyName,omitempty" name:"SwitchPortPhyName"`
	// 网关端口类型

	SwitchPortType *string `json:"SwitchPortType,omitempty" name:"SwitchPortType"`
	// 网关sn

	SwitchSn []*string `json:"SwitchSn,omitempty" name:"SwitchSn"`
	// underlay网关

	UnderlayGateway *string `json:"UnderlayGateway,omitempty" name:"UnderlayGateway"`
	// underlayip

	UnderlayIp *string `json:"UnderlayIp,omitempty" name:"UnderlayIp"`
	// underlay掩码

	UnderlayMask *string `json:"UnderlayMask,omitempty" name:"UnderlayMask"`
	// 可用区域id

	Zone *int64 `json:"Zone,omitempty" name:"Zone"`
	// 可用区名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 服务器id

	Id *string `json:"Id,omitempty" name:"Id"`
}

type MiddleInput struct {
	// 模糊匹配输入的aapid或者bms-id字段

	Input []*string `json:"Input,omitempty" name:"Input"`
}

type Reply struct {
	// Code

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// Msg

	Msg *string `json:"Msg,omitempty" name:"Msg"`
}

type ServersInfo struct {
	// 机器可用环境信息

	Condition *Condition `json:"Condition,omitempty" name:"Condition"`
	// 服务器修改信息

	Modify *Modify `json:"Modify,omitempty" name:"Modify"`
}

type UpdateDefaultQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDefaultQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDefaultQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditSwitchUserInfo struct {
	// 可用区ID

	Zone *uint64 `json:"Zone,omitempty" name:"Zone"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 可用地域ID

	Region *uint64 `json:"Region,omitempty" name:"Region"`
	// 密码

	PassWord *string `json:"PassWord,omitempty" name:"PassWord"`
}

type Tags struct {
	// 标签主键

	Tag *string `json:"Tag,omitempty" name:"Tag"`
	// 标准值

	Vlaue *string `json:"Vlaue,omitempty" name:"Vlaue"`
}

type DeleteNodeExRequest struct {
	*tchttp.BaseRequest

	// 裸金属服务器序列号

	Sn []*string `json:"Sn,omitempty" name:"Sn"`
}

func (r *DeleteNodeExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNodeExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchUserExRequest struct {
	*tchttp.BaseRequest

	// 可用地域Id

	Zone *uint64 `json:"Zone,omitempty" name:"Zone"`
	// 可用区Id

	RegionNum *uint64 `json:"RegionNum,omitempty" name:"RegionNum"`
}

func (r *DescribeSwitchUserExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchUserExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchtypeExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 交换机厂商类型

		SwitchType []*string `json:"SwitchType,omitempty" name:"SwitchType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSwitchtypeExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchtypeExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Switch struct {
	// 主IP

	AdminIp *string `json:"AdminIp,omitempty" name:"AdminIp"`
	// 主端口

	AdminPort *string `json:"AdminPort,omitempty" name:"AdminPort"`
	// 网络自治系统数量

	AsNum *string `json:"AsNum,omitempty" name:"AsNum"`
	// 网关边界协议名称

	BgpIfName *string `json:"BgpIfName,omitempty" name:"BgpIfName"`
	// 网关边界协议类型

	BgpIfType *string `json:"BgpIfType,omitempty" name:"BgpIfType"`
	// 厂商

	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 地域ID

	Region *uint64 `json:"Region,omitempty" name:"Region"`
	// 网关端点名称

	Role *string `json:"Role,omitempty" name:"Role"`
	// 固资号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 可用区域ID

	Zone *uint64 `json:"Zone,omitempty" name:"Zone"`
	// VtepIp

	VtepIp *string `json:"VtepIp,omitempty" name:"VtepIp"`
	// 可用区域ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// AZ名

	Az *string `json:"Az,omitempty" name:"Az"`
	// 固资号

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 删除状态

	DeleteStatus *string `json:"DeleteStatus,omitempty" name:"DeleteStatus"`
	// 机器类型

	Engine *string `json:"Engine,omitempty" name:"Engine"`
	// IDC信息

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 管理IP

	ManagementIp *string `json:"ManagementIp,omitempty" name:"ManagementIp"`
	// 机型

	Model *string `json:"Model,omitempty" name:"Model"`
	// 网络设备号

	NetworkDeviceId *int64 `json:"NetworkDeviceId,omitempty" name:"NetworkDeviceId"`
	// 操作系统

	Os *string `json:"Os,omitempty" name:"Os"`
	// 机架号

	Rack *string `json:"Rack,omitempty" name:"Rack"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 机器类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type SwitchUserSet struct {
	// 用户名称

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 可用区域ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用地域ID

	Region *int64 `json:"Region,omitempty" name:"Region"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

type DescribeSellstateExRequest struct {
	*tchttp.BaseRequest

	// 可用区Id

	Zone *uint64 `json:"Zone,omitempty" name:"Zone"`
	// "FlavorId":["flavor-0eff6916","flavor-288014ac","flavor-2473629e","flavor-49405f99","flavor-53264765","flavor-43a3d65a","flavor-02309427","flavor-5381534c","flavor-58642aaa","flavor-56674ce5"]

	FlavorId []*string `json:"FlavorId,omitempty" name:"FlavorId"`
	// 可用区Id

	RegionNum *uint64 `json:"RegionNum,omitempty" name:"RegionNum"`
}

func (r *DescribeSellstateExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSellstateExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchtypeExRequest struct {
	*tchttp.BaseRequest

	// 可用地域Id

	Zone *uint64 `json:"Zone,omitempty" name:"Zone"`
	// 可用区Id

	RegionNum *uint64 `json:"RegionNum,omitempty" name:"RegionNum"`
}

func (r *DescribeSwitchtypeExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchtypeExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSwitchExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSwitchExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSwitchExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBmsTasksExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的返回条目数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// BMS创建任务返回信息

		BmsTaskSet []*BmsTaskSet `json:"BmsTaskSet,omitempty" name:"BmsTaskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBmsTasksExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBmsTasksExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFlavorInfo struct {
	// 可用地域ID

	Region *uint64 `json:"Region,omitempty" name:"Region"`
	// 可用区ID

	Zone *uint64 `json:"Zone,omitempty" name:"Zone"`
	// 机型ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

type AddNodeToFlavorExRequest struct {
	*tchttp.BaseRequest

	// 添加的裸金属资源信息

	BmsNode []*BmsNode `json:"BmsNode,omitempty" name:"BmsNode"`
	// 机型ID

	FlavorId *string `json:"FlavorId,omitempty" name:"FlavorId"`
}

func (r *AddNodeToFlavorExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNodeToFlavorExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Netinfo struct {
	// Bond

	Bond *string `json:"Bond,omitempty" name:"Bond"`
	// Mac

	Mac *string `json:"Mac,omitempty" name:"Mac"`
	// Nicid

	Nicid *string `json:"Nicid,omitempty" name:"Nicid"`
	// Overlaygateway

	Overlaygateway *string `json:"Overlaygateway,omitempty" name:"Overlaygateway"`
	// Overlayip

	Overlayip *string `json:"Overlayip,omitempty" name:"Overlayip"`
	// Overlaymask

	Overlaymask *string `json:"Overlaymask,omitempty" name:"Overlaymask"`
	// Overlaysubnet

	Overlaysubnet *string `json:"Overlaysubnet,omitempty" name:"Overlaysubnet"`
	// Underlaygateway

	Underlaygateway *string `json:"Underlaygateway,omitempty" name:"Underlaygateway"`
	// Underlayip

	Underlayip *string `json:"Underlayip,omitempty" name:"Underlayip"`
	// Underlaymask

	Underlaymask *string `json:"Underlaymask,omitempty" name:"Underlaymask"`
}

type AddSwitchInfo struct {
	// 网络自治系统数量

	AsNum *string `json:"AsNum,omitempty" name:"AsNum"`
	// 网关边界协议类型

	BgpIfType *string `json:"BgpIfType,omitempty" name:"BgpIfType"`
	// 网关边界协议名称

	BgpIfName *string `json:"BgpIfName,omitempty" name:"BgpIfName"`
	// 可用区ID

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 固资号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 交换机业务IP

	VtepIp *string `json:"VtepIp,omitempty" name:"VtepIp"`
	// 网关端点名称

	Role *string `json:"Role,omitempty" name:"Role"`
	// 可用地域ID

	Region *string `json:"Region,omitempty" name:"Region"`
	// 产品型号

	Product *string `json:"Product,omitempty" name:"Product"`
	// 主IP

	AdminIp *string `json:"AdminIp,omitempty" name:"AdminIp"`
	// 厂商

	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`
	// 设备名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type Condition struct {
	// 可用地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 可用环境名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 机器序列号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
}

type DefaultQuotaInfo struct {
	// 资源配型

	Resourcetype *string `json:"Resourcetype,omitempty" name:"Resourcetype"`
	// 最大配额

	Maxquota *uint64 `json:"Maxquota,omitempty" name:"Maxquota"`
	// 默认配额

	Defaultquota *uint64 `json:"Defaultquota,omitempty" name:"Defaultquota"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
}

type Instanceset struct {
	// Appid

	Appid *int64 `json:"Appid,omitempty" name:"Appid"`
	// Dcostaskid

	Dcostaskid *int64 `json:"Dcostaskid,omitempty" name:"Dcostaskid"`
	// Flavorid

	Flavorid *string `json:"Flavorid,omitempty" name:"Flavorid"`
	// Instancename

	Instancename *string `json:"Instancename,omitempty" name:"Instancename"`
	// Internetaccessible

	Internetaccessible *Internetaccessible `json:"Internetaccessible,omitempty" name:"Internetaccessible"`
	// Nodesn

	Nodesn *string `json:"Nodesn,omitempty" name:"Nodesn"`
	// Operatingsystem

	Operatingsystem *string `json:"Operatingsystem,omitempty" name:"Operatingsystem"`
	// Operatingsystemtype

	Operatingsystemtype *string `json:"Operatingsystemtype,omitempty" name:"Operatingsystemtype"`
	// AddNodePlacement

	AddNodePlacement *AddNodePlacement `json:"AddNodePlacement,omitempty" name:"AddNodePlacement"`
	// Raidtype

	Raidtype *string `json:"Raidtype,omitempty" name:"Raidtype"`
	// Subaccountuin

	Subaccountuin *string `json:"Subaccountuin,omitempty" name:"Subaccountuin"`
	// Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// Virtualprivatecloud

	Virtualprivatecloud *VirtualPrivateCloud `json:"Virtualprivatecloud,omitempty" name:"Virtualprivatecloud"`
	// Bprocessfinal

	Bprocessfinal *int64 `json:"Bprocessfinal,omitempty" name:"Bprocessfinal"`
	// Bmsid

	Bmsid *string `json:"Bmsid,omitempty" name:"Bmsid"`
	// Netinfo

	Netinfo *Netinfo `json:"Netinfo,omitempty" name:"Netinfo"`
	// Peerconnect

	Peerconnect *Peerconnect `json:"Peerconnect,omitempty" name:"Peerconnect"`
	// Reply

	Reply *Reply `json:"Reply,omitempty" name:"Reply"`
	// Routermac

	Routermac *string `json:"Routermac,omitempty" name:"Routermac"`
	// Sn

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// SvrIdcId

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// Sxgwip

	Sxgwip *string `json:"Sxgwip,omitempty" name:"Sxgwip"`
	// Userdefined

	Userdefined *int64 `json:"Userdefined,omitempty" name:"Userdefined"`
	// Vtepip

	Vtepip *string `json:"Vtepip,omitempty" name:"Vtepip"`
	// Xgwip

	Xgwip *string `json:"Xgwip,omitempty" name:"Xgwip"`
}

type DeleteFlavorExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFlavorExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFlavorExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeObmsNodeFilterExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// BMS服务器已绑定的机型名称列表

		BoundType []*string `json:"BoundType,omitempty" name:"BoundType"`
		// BMS服务器已绑定的机型类型列表

		ModelType []*int64 `json:"ModelType,omitempty" name:"ModelType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeObmsNodeFilterExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeObmsNodeFilterExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceObj struct {
	// available

	Available *uint64 `json:"Available,omitempty" name:"Available"`
	// sold

	Sold *uint64 `json:"Sold,omitempty" name:"Sold"`
	// total

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type DescribeBmsQuotaRequest struct {
	*tchttp.BaseRequest

	// 需要获取配额信息的AppId列表

	AppList []*string `json:"AppList,omitempty" name:"AppList"`
	// 需要获取的资源类型，默认为’Bare Metal Server‘

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 每页返回的条目数，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页起始位置，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBmsQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBmsQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNodeExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateNodeExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNodeExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Between struct {
	// 过滤条件

	AddTime []*string `json:"AddTime,omitempty" name:"AddTime"`
}

type BmsQuota struct {
	// 资源类型

	Resourcetype *string `json:"Resourcetype,omitempty" name:"Resourcetype"`
	// 资源配额

	Resourcequota *uint64 `json:"Resourcequota,omitempty" name:"Resourcequota"`
	// 实际占用数量

	Realused *uint64 `json:"Realused,omitempty" name:"Realused"`
	// 描述信息

	Description *uint64 `json:"Description,omitempty" name:"Description"`
	// 配额所属的AppId值

	Appid *string `json:"Appid,omitempty" name:"Appid"`
	// 最大配额

	Maxquota *uint64 `json:"Maxquota,omitempty" name:"Maxquota"`
}

type AddBmsTypeBillExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddBmsTypeBillExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddBmsTypeBillExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFlavorExRequest struct {
	*tchttp.BaseRequest

	// "DeleteFlavorInfo":[{"Region":50000005,"Zone":50050002,"Id":"flavor-35198b6f"}]

	DeleteFlavorInfo []*DeleteFlavorInfo `json:"DeleteFlavorInfo,omitempty" name:"DeleteFlavorInfo"`
}

func (r *DeleteFlavorExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFlavorExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlavorExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的机型数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 机型套餐详细信息列表

		FlavorSet []*FlavorSet `json:"FlavorSet,omitempty" name:"FlavorSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlavorExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlavorExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSwitchExRequest struct {
	*tchttp.BaseRequest

	// 更新交换机的详细信息

	UpdateSwitchInfo []*UpdateSwitchInfo `json:"UpdateSwitchInfo,omitempty" name:"UpdateSwitchInfo"`
}

func (r *UpdateSwitchExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSwitchExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BmsNode struct {
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名字

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区名字

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 资产ID

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// 序列号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 机器在CMDB的运行状态

	HostStatus *string `json:"HostStatus,omitempty" name:"HostStatus"`
	// 机器MAC地址

	Mac *string `json:"Mac,omitempty" name:"Mac"`
	// 聚合端口

	AggregationPort *string `json:"AggregationPort,omitempty" name:"AggregationPort"`
	// 机器磁盘信息

	Disk []*Disk `json:"Disk,omitempty" name:"Disk"`
	// 机器名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 网卡Bond标志

	BondingFlag *string `json:"BondingFlag,omitempty" name:"BondingFlag"`
	// 物理IP地址

	UnderlayIp *string `json:"UnderlayIp,omitempty" name:"UnderlayIp"`
	// 物理IP地址掩码

	UnderlayMask *string `json:"UnderlayMask,omitempty" name:"UnderlayMask"`
	// 物理IP地址网关

	UnderlayGateway *string `json:"UnderlayGateway,omitempty" name:"UnderlayGateway"`
	// 交换机端口类型

	SwitchPortType *string `json:"SwitchPortType,omitempty" name:"SwitchPortType"`
	// 交换机端口名字

	SwitchPortName *string `json:"SwitchPortName,omitempty" name:"SwitchPortName"`
	// 交换机物理端口名字

	SwitchPortPhyName *string `json:"SwitchPortPhyName,omitempty" name:"SwitchPortPhyName"`
	// 交换机资产信息

	SwitchSn []*string `json:"SwitchSn,omitempty" name:"SwitchSn"`
	// 交换机详细信息

	Switch []*Switch `json:"Switch,omitempty" name:"Switch"`
	// BMS实例运行状态

	TaskState *string `json:"TaskState,omitempty" name:"TaskState"`
	// 绑定机型的详细信息

	FlavorInfo *FlavorSet `json:"FlavorInfo,omitempty" name:"FlavorInfo"`
	// 自定义机型标志

	UserDefined *uint64 `json:"UserDefined,omitempty" name:"UserDefined"`
	// 用户ID

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 虚拟私有网络ID

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// 虚拟网络IP地址

	OverlayIp *string `json:"OverlayIp,omitempty" name:"OverlayIp"`
	// 虚拟IP地址掩码

	OverlayMask *string `json:"OverlayMask,omitempty" name:"OverlayMask"`
	// 售卖状态

	SellState *string `json:"SellState,omitempty" name:"SellState"`
	// 裸金属服务器ID

	BmsId *string `json:"BmsId,omitempty" name:"BmsId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 机型ID

	FlavorId *string `json:"FlavorId,omitempty" name:"FlavorId"`
	// 操作系统版本

	OsVersion *string `json:"OsVersion,omitempty" name:"OsVersion"`
	// 虚拟IP地址网关

	OverlayGateway *string `json:"OverlayGateway,omitempty" name:"OverlayGateway"`
	// 子网ID

	SubnetId *uint64 `json:"SubnetId,omitempty" name:"SubnetId"`
	// 数据中心ID

	IdcId *int64 `json:"IdcId,omitempty" name:"IdcId"`
	// 数据中心名字

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// 服务器机型

	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
	// CPU型号，X86/ARM

	CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
	// 系统类型

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 电源情况

	PowerState *string `json:"PowerState,omitempty" name:"PowerState"`
	// 主机名

	Hostname *string `json:"Hostname,omitempty" name:"Hostname"`
	// Raid信息

	Raid *string `json:"Raid,omitempty" name:"Raid"`
	// 下次分配权重

	ScheduleWeight *int64 `json:"ScheduleWeight,omitempty" name:"ScheduleWeight"`
	// nic信息

	NicId *string `json:"NicId,omitempty" name:"NicId"`
	// 当前使用者的Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 当前上联sxgw网关的VIP

	SxgwIp *string `json:"SxgwIp,omitempty" name:"SxgwIp"`
	// 当前上联xgw网关的VIP

	XgwIp *string `json:"XgwIp,omitempty" name:"XgwIp"`
	// 绑定标签信息

	Tag []*Tags `json:"Tag,omitempty" name:"Tag"`
	// IloIp信息

	IloIp *string `json:"IloIp,omitempty" name:"IloIp"`
	// 申请的主机名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 地域IP

	Region *uint64 `json:"Region,omitempty" name:"Region"`
	// 机架名称

	Rack *string `json:"Rack,omitempty" name:"Rack"`
	// 心跳交换机端口

	HeartPort *string `json:"HeartPort,omitempty" name:"HeartPort"`
	// 心跳交换机端口名

	HeartPortName *string `json:"HeartPortName,omitempty" name:"HeartPortName"`
	// 备用IP

	BackupIPs *string `json:"BackupIPs,omitempty" name:"BackupIPs"`
}

type Modify struct {
	// 服务器设备名称

	SvrDeviceName *string `json:"SvrDeviceName,omitempty" name:"SvrDeviceName"`
	// 服务器产品标识

	SvrBelongingProducts *string `json:"SvrBelongingProducts,omitempty" name:"SvrBelongingProducts"`
}

type SwitchSet struct {
	// 产品名称

	Product *string `json:"Product,omitempty" name:"Product"`
	// 网关IP

	Vtepip *string `json:"Vtepip,omitempty" name:"Vtepip"`
	// BGP的IP

	Bgpip *string `json:"Bgpip,omitempty" name:"Bgpip"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 可用地域

	Region *int64 `json:"Region,omitempty" name:"Region"`
	// 主IP

	Adminip *string `json:"Adminip,omitempty" name:"Adminip"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 状态

	State *string `json:"State,omitempty" name:"State"`
	// 密码状态

	Passwordstate *int64 `json:"Passwordstate,omitempty" name:"Passwordstate"`
	// 规则

	Role *string `json:"Role,omitempty" name:"Role"`
	// 序列号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 路由mac

	Routermac *string `json:"Routermac,omitempty" name:"Routermac"`
	// 主端口

	Adminport *string `json:"Adminport,omitempty" name:"Adminport"`
	// 可用区域ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 配置状态

	Configstate *int64 `json:"Configstate,omitempty" name:"Configstate"`
	// BGP的AS号

	Bgpas *int64 `json:"Bgpas,omitempty" name:"Bgpas"`
	// 厂商名称

	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`
}

type AddSwitchUserExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSwitchUserExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSwitchUserExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSwitchExRequest struct {
	*tchttp.BaseRequest

	// "AddSwitchInfo":[{"AdminIp":"10.10.208.16","Manufacturer":"H3C","Name":"CQ-YX","Product":"LS-6800","Region":50000005,"Role":"ROLE_RR","VtepIp":"3.3.3.3","Sn":"TYN11111","Zone":50050002,"BgpIfName":"0","BgpIfType":"LoopBack","As_num":""}]

	AddSwitchInfo []*Switch `json:"AddSwitchInfo,omitempty" name:"AddSwitchInfo"`
}

func (r *AddSwitchExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSwitchExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Output struct {
	// App ID

	Appid *int64 `json:"Appid,omitempty" name:"Appid"`
	// 子账号Uni

	Subaccountuin *string `json:"Subaccountuin,omitempty" name:"Subaccountuin"`
	// 主账号Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 程序结果

	Bprocessfinal *int64 `json:"Bprocessfinal,omitempty" name:"Bprocessfinal"`
	// 错误码

	Errorcode *int64 `json:"Errorcode,omitempty" name:"Errorcode"`
	// 错误信息

	Errormsg *string `json:"Errormsg,omitempty" name:"Errormsg"`
	// 创建BMS的信息

	Instanceset []*Instanceset `json:"Instanceset,omitempty" name:"Instanceset"`
	// 创建任务返回信息

	Reply *Reply `json:"Reply,omitempty" name:"Reply"`
	// 创建任务的ID

	Taskid *string `json:"Taskid,omitempty" name:"Taskid"`
	// 任务Ws名

	Wsname *string `json:"Wsname,omitempty" name:"Wsname"`
}

type AddBmsTypeBillExRequest struct {
	*tchttp.BaseRequest

	// 新增机型自定义类型

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 地域id

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

func (r *AddBmsTypeBillExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddBmsTypeBillExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Peerconnect struct {
	// Localip

	Localip *string `json:"Localip,omitempty" name:"Localip"`
	// Localmac

	Localmac *string `json:"Localmac,omitempty" name:"Localmac"`
	// Netmask

	Netmask *string `json:"Netmask,omitempty" name:"Netmask"`
	// Peerip

	Peerip *string `json:"Peerip,omitempty" name:"Peerip"`
	// Peermac

	Peermac *string `json:"Peermac,omitempty" name:"Peermac"`
}

type DescribeFlavorExRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 可用地域id

	Zone *uint64 `json:"Zone,omitempty" name:"Zone"`
	// 可用区域id

	RegionNum *uint64 `json:"RegionNum,omitempty" name:"RegionNum"`
	// 排序参数，只能填server_num，sold_num，available_num这三个参数中的某一个

	SortItem *string `json:"SortItem,omitempty" name:"SortItem"`
	// asc, desc 两种排序，默认是desc

	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
	// 机型ID

	Id []*string `json:"Id,omitempty" name:"Id"`
	// 机型名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否用户自定义机型

	UserDefined *string `json:"UserDefined,omitempty" name:"UserDefined"`
	// CPU类型：X86/ARM

	CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
	// 机型网口数

	NetworkPorts *int64 `json:"NetworkPorts,omitempty" name:"NetworkPorts"`
}

func (r *DescribeFlavorExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlavorExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateFlavorExRequest struct {
	*tchttp.BaseRequest

	// 更改后机型类型详细信息

	UpdateFlavorInfo []*UpdateFlavorInfo `json:"UpdateFlavorInfo,omitempty" name:"UpdateFlavorInfo"`
}

func (r *UpdateFlavorExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateFlavorExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchExRequest struct {
	*tchttp.BaseRequest

	// 每页返回的条目数，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页起始位置，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 可用区Id

	Zone *uint64 `json:"Zone,omitempty" name:"Zone"`
	// 交换机类型

	Role []*string `json:"Role,omitempty" name:"Role"`
	// 可用区Id

	RegionNum *uint64 `json:"RegionNum,omitempty" name:"RegionNum"`
	// 交换机Sn号

	Sn []*string `json:"Sn,omitempty" name:"Sn"`
	// 交换机管理IP

	AdminIp *string `json:"AdminIp,omitempty" name:"AdminIp"`
}

func (r *DescribeSwitchExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BmsTaskSet struct {
	// 任务ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 功能名称

	Functionname *string `json:"Functionname,omitempty" name:"Functionname"`
	// APP ID

	Appid *int64 `json:"Appid,omitempty" name:"Appid"`
	// Bms ID

	Bmsid *string `json:"Bmsid,omitempty" name:"Bmsid"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 返回消息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 添加时间

	Addtime *string `json:"Addtime,omitempty" name:"Addtime"`
}

type VirtualPrivateCloud struct {
	// 私有网络ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// IP地址

	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 是否分配IPv6地址

	Ipv6Address *bool `json:"Ipv6Address,omitempty" name:"Ipv6Address"`
}

type DescribeBmsDefaultQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 默认配额信息列表

		QuotaSet []*DefaultQuotaInfo `json:"QuotaSet,omitempty" name:"QuotaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBmsDefaultQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBmsDefaultQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReturnNodeSet struct {
	// 结果信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 可用地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 返回结果标识

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 服务器固资号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 可用区域名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type UpdateSwitchUserExRequest struct {
	*tchttp.BaseRequest

	// 更新交换机的的用户信息

	UpdateSwitchUserInfo []*EditSwitchUserInfo `json:"UpdateSwitchUserInfo,omitempty" name:"UpdateSwitchUserInfo"`
}

func (r *UpdateSwitchUserExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSwitchUserExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperatingSystem struct {
	// 支持的linux系统列表

	Linux []*string `json:"Linux,omitempty" name:"Linux"`
	// 支持的其他系统列表

	Other []*string `json:"Other,omitempty" name:"Other"`
	// 支持的Windows系统列表

	Windows []*string `json:"Windows,omitempty" name:"Windows"`
}

type SwitchDel struct {
	// 固资号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
}

type AddSwitchUserExRequest struct {
	*tchttp.BaseRequest

	// "AddSwitchUserInfo":[{"Region":50000005,"Zone":50050002,"UserName":"tencent","PassWord":"tencentpoc"}]

	AddSwitchUserInfo []*EditSwitchUserInfo `json:"AddSwitchUserInfo,omitempty" name:"AddSwitchUserInfo"`
}

func (r *AddSwitchUserExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSwitchUserExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateBmsQuotaRequest struct {
	*tchttp.BaseRequest

	// 需要更新的配额值

	ResourceQuota *uint64 `json:"ResourceQuota,omitempty" name:"ResourceQuota"`
	// 需要更新的资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 需要更新配额的AppId值

	OwnerId *string `json:"OwnerId,omitempty" name:"OwnerId"`
}

func (r *UpdateBmsQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateBmsQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSwitchUserExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSwitchUserExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSwitchUserExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Internetaccessible struct {
	// Internetmaxbandwidthout

	Internetmaxbandwidthout *int64 `json:"Internetmaxbandwidthout,omitempty" name:"Internetmaxbandwidthout"`
	// Publicipassigned

	Publicipassigned *bool `json:"Publicipassigned,omitempty" name:"Publicipassigned"`
}

type DescribeBmsDefaultQuotaRequest struct {
	*tchttp.BaseRequest

	// 配额资源类型，默认值为“BMS Metal Server”

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 每页返回的条目数，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页起始位置，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBmsDefaultQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBmsDefaultQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNodeExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNodeExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNodeExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddNodeToVPCExRequest struct {
	*tchttp.BaseRequest

	// 平台信息

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 机型id

	FlavorId *string `json:"FlavorId,omitempty" name:"FlavorId"`
	// 操作系统类型

	OperatingSystemType *string `json:"OperatingSystemType,omitempty" name:"OperatingSystemType"`
	// 操作系统

	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`
	// Raid类型

	RaidType *string `json:"RaidType,omitempty" name:"RaidType"`
	// 虚拟私有云信息

	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
	// 实例数

	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// 资源列表

	NodeList []*string `json:"NodeList,omitempty" name:"NodeList"`
	// 主账号Uin

	InnerUin *string `json:"InnerUin,omitempty" name:"InnerUin"`
	// 子账号Uin

	InnerSubAccountUin *string `json:"InnerSubAccountUin,omitempty" name:"InnerSubAccountUin"`
	// APP ID

	InnerAppId *string `json:"InnerAppId,omitempty" name:"InnerAppId"`
}

func (r *AddNodeToVPCExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNodeToVPCExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReturnNodeToObmsPoolExRequest struct {
	*tchttp.BaseRequest

	// 裸金属服务器序列号

	Sn []*string `json:"Sn,omitempty" name:"Sn"`
}

func (r *ReturnNodeToObmsPoolExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReturnNodeToObmsPoolExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNodeExRequest struct {
	*tchttp.BaseRequest

	// 修改BMS服务器的详细信息

	UpdateNodeInfo []*UpdateNodeInfo `json:"UpdateNodeInfo,omitempty" name:"UpdateNodeInfo"`
}

func (r *UpdateNodeExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNodeExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSwitchExRequest struct {
	*tchttp.BaseRequest

	// "DeleteSwitchInfo":[{"Sn":"TYN11111"}]

	DeleteSwitchInfo []*SwitchDel `json:"DeleteSwitchInfo,omitempty" name:"DeleteSwitchInfo"`
}

func (r *DeleteSwitchExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSwitchExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// BMS资源池的服务器总个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// BMS服务器信息

		BmsNode []*BmsNode `json:"BmsNode,omitempty" name:"BmsNode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNodeExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeObmsPoolNodesToAddExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的服务器个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 裸金属服务器节点信息

		BmsNode []*BmsNode `json:"BmsNode,omitempty" name:"BmsNode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeObmsPoolNodesToAddExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeObmsPoolNodesToAddExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的交换机条目数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 交换机详细信息

		SwitchSet []*SwitchSet `json:"SwitchSet,omitempty" name:"SwitchSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSwitchExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateFlavorExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateFlavorExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateFlavorExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddTime struct {
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 截止时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type AddNodeToVPCExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddNodeToVPCExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNodeToVPCExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddNodeToFlavorExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddNodeToFlavorExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNodeToFlavorExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSwitchExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSwitchExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSwitchExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBmsTasksExRequest struct {
	*tchttp.BaseRequest

	// 分页起始位置，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页返回的条目数，

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 可用区域id

	Zone *uint64 `json:"Zone,omitempty" name:"Zone"`
	// 排序类型

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 可用区域id

	RegionNum *uint64 `json:"RegionNum,omitempty" name:"RegionNum"`
	// 过滤条件

	Between *Between `json:"Between,omitempty" name:"Between"`
	// 模糊匹配参数

	MiddleLike *MiddleInput `json:"MiddleLike,omitempty" name:"MiddleLike"`
	// task-id列表

	BmsTaskId []*string `json:"BmsTaskId,omitempty" name:"BmsTaskId"`
	// 查询状态

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 查询操作类型（生产、退还、重启等）字符串

	FunctionName []*string `json:"FunctionName,omitempty" name:"FunctionName"`
}

func (r *DescribeBmsTasksExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBmsTasksExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchUserExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SwitchUserSet

		SwitchUserSet []*SwitchUserSet `json:"SwitchUserSet,omitempty" name:"SwitchUserSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSwitchUserExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchUserExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReturnNodeToObmsPoolExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReturnNodeToObmsPoolExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReturnNodeToObmsPoolExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddFlavorExRequest struct {
	*tchttp.BaseRequest

	// "AddFlavorInfo":[{"Os":{"Linux":["suse12-1"]},"Name":"Test-M12","Cpu":"E5-2670v3","Memory":"128GB","Disk":"1","Netspeed":"1","Zone":50050002,"Raid":["{7RAID5+1HOT}+{7RAID5+1HOT+7RAID5+1HOT}"]

	AddFlavorInfo []*AddFlavorInfo `json:"AddFlavorInfo,omitempty" name:"AddFlavorInfo"`
}

func (r *AddFlavorExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddFlavorExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSwitchExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSwitchExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSwitchExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddFlavorInfo struct {
	// 支持的操作系统列表

	Os *OperatingSystem `json:"Os,omitempty" name:"Os"`
	// 机型名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// CPU

	Cpu *string `json:"Cpu,omitempty" name:"Cpu"`
	// 内存

	Memory *string `json:"Memory,omitempty" name:"Memory"`
	// 磁盘

	Disk *string `json:"Disk,omitempty" name:"Disk"`
	// 网盘速率

	NetSpeed *string `json:"NetSpeed,omitempty" name:"NetSpeed"`
	// 可用区ID

	Zone *uint64 `json:"Zone,omitempty" name:"Zone"`
	// 支持的RAID

	Raid []*string `json:"Raid,omitempty" name:"Raid"`
	// 上架标识
	// 0:未上架
	// 1:上架

	IsShow *uint64 `json:"IsShow,omitempty" name:"IsShow"`
	// 可用地域ID

	Region *uint64 `json:"Region,omitempty" name:"Region"`
	// 用户ID

	UserDefined *uint64 `json:"UserDefined,omitempty" name:"UserDefined"`
	// 计费机型

	FlavorType *string `json:"FlavorType,omitempty" name:"FlavorType"`
	// CPU类型：X86/ARM

	CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
	// 上联网口数，2表示双网卡，4表示有心跳交换机的四网口

	NetworkPorts *uint64 `json:"NetworkPorts,omitempty" name:"NetworkPorts"`
	// dcos后置脚本

	DcosShell *string `json:"DcosShell,omitempty" name:"DcosShell"`
}

type DescribeBmsQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 租户配额列表

		QuotaSet []*BmsQuota `json:"QuotaSet,omitempty" name:"QuotaSet"`
		// 列表大小

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBmsQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBmsQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateBmsQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateBmsQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateBmsQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlavorInfo struct {
	// Raid

	Raid []*string `json:"Raid,omitempty" name:"Raid"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Userdefined

	Userdefined *int64 `json:"Userdefined,omitempty" name:"Userdefined"`
	// Netspeed

	Netspeed *string `json:"Netspeed,omitempty" name:"Netspeed"`
	// Cpu

	Cpu *string `json:"Cpu,omitempty" name:"Cpu"`
	// Soldoutmode

	Soldoutmode *string `json:"Soldoutmode,omitempty" name:"Soldoutmode"`
	// Isshow

	Isshow *int64 `json:"Isshow,omitempty" name:"Isshow"`
	// Whitetype

	Whitetype *string `json:"Whitetype,omitempty" name:"Whitetype"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Memory

	Memory *string `json:"Memory,omitempty" name:"Memory"`
	// Region

	Region *int64 `json:"Region,omitempty" name:"Region"`
	// Disk

	Disk *string `json:"Disk,omitempty" name:"Disk"`
	// Os

	Os *Os `json:"Os,omitempty" name:"Os"`
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// ZoneId

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type UpdateFlavorInfo struct {
	// 可用地域ID

	Region *uint64 `json:"Region,omitempty" name:"Region"`
	// 可用区ID

	Zone *uint64 `json:"Zone,omitempty" name:"Zone"`
	// 上架标识
	// 0:未上架
	// 1:上架

	IsShow *uint64 `json:"IsShow,omitempty" name:"IsShow"`
	// 机型ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 处理器

	Cpu *string `json:"Cpu,omitempty" name:"Cpu"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 硬盘大小

	Disk *string `json:"Disk,omitempty" name:"Disk"`
	// 内存大小

	Memory *string `json:"Memory,omitempty" name:"Memory"`
	// 机型名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 网卡速度

	NetSpeed *string `json:"NetSpeed,omitempty" name:"NetSpeed"`
	// 操作系统

	Os *Os `json:"Os,omitempty" name:"Os"`
	// 磁盘阵列

	Raid []*string `json:"Raid,omitempty" name:"Raid"`
	// 是否售罄

	SoldOutMode *string `json:"SoldOutMode,omitempty" name:"SoldOutMode"`
	// 是否用户自定义机型

	UserDefined *uint64 `json:"UserDefined,omitempty" name:"UserDefined"`
	// 机器类型

	DeviceObj *DeviceObj `json:"DeviceObj,omitempty" name:"DeviceObj"`
	// 机型绑定白名单

	WhiteType *string `json:"WhiteType,omitempty" name:"WhiteType"`
	// 计费机型

	FlavorType *string `json:"FlavorType,omitempty" name:"FlavorType"`
	// CPU型号，X86/ARM

	CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
	// 机型网口数

	NetworkPorts *uint64 `json:"NetworkPorts,omitempty" name:"NetworkPorts"`
	// dcos后置脚本

	DcosShell *string `json:"DcosShell,omitempty" name:"DcosShell"`
}
