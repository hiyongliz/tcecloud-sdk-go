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

package v20200421

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type CmdbCondition struct {
	// 网络设备id

	NetdevId *string `json:"NetdevId,omitempty" name:"NetdevId"`
}

type OutbandStrategies struct {
	// 删除/修改条件

	Condition *OutbandStrategiesCondition `json:"Condition,omitempty" name:"Condition"`
	// 修改内容

	Modify *AddOutbandStrategies `json:"Modify,omitempty" name:"Modify"`
}

type RenameServerCondition struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// Sn

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 内网IP

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
}

type ServerProcessInfo struct {
	// "Detail":"ok"

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// "Result":0

	Result *string `json:"Result,omitempty" name:"Result"`
	// 服务器进程结果信息

	ResultInfo *ServerProcessResultInfo `json:"ResultInfo,omitempty" name:"ResultInfo"`
}

type ModifyLabelsRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID，system_id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key，system_key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"modify"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 修改服务器参数

	LabelServerSet []*ModifyLabel `json:"LabelServerSet,omitempty" name:"LabelServerSet"`
}

func (r *ModifyLabelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLabelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOutbandStrategies struct {
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 是否包含数字（默认值1）

	HaveNumber *uint64 `json:"HaveNumber,omitempty" name:"HaveNumber"`
	// 是否包含小写字母（默认值1）

	HaveLowerLetter *uint64 `json:"HaveLowerLetter,omitempty" name:"HaveLowerLetter"`
	// 是否包含大写字母（默认值1）

	HaveUpperLetter *uint64 `json:"HaveUpperLetter,omitempty" name:"HaveUpperLetter"`
	// 特殊字符

	SpecialCharacters *string `json:"SpecialCharacters,omitempty" name:"SpecialCharacters"`
	// 密码长度，默认15

	PasswdLength *uint64 `json:"PasswdLength,omitempty" name:"PasswdLength"`
	// 是否启用该策略 （默认禁用，即0）

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
}

type ReinstallOsInfo struct {
	// SvrIdcName

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// 机器类型

	SvrDeviceType *string `json:"SvrDeviceType,omitempty" name:"SvrDeviceType"`
	// SN号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// Lan信息

	Lan *Lan `json:"Lan,omitempty" name:"Lan"`
	// Wan信息

	Wan *Wan `json:"Wan,omitempty" name:"Wan"`
	// 操作系统

	Os *string `json:"Os,omitempty" name:"Os"`
	// Raid信息

	Raid *string `json:"Raid,omitempty" name:"Raid"`
	// 固资编号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 1：Bonding 0：非Bonding

	BondingFlag *int64 `json:"BondingFlag,omitempty" name:"BondingFlag"`
	// UserShellPath

	UserShellPath *string `json:"UserShellPath,omitempty" name:"UserShellPath"`
	// SvrIdcId

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// 是否绑定raid，0:不绑定;1:绑定

	RaidFlag *int64 `json:"RaidFlag,omitempty" name:"RaidFlag"`
	// 硬件描述信息

	SvrDeviceDescript *string `json:"SvrDeviceDescript,omitempty" name:"SvrDeviceDescript"`
	// 架构

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	// 分组

	Group *string `json:"Group,omitempty" name:"Group"`
	// 是否支持ipv6 yes or no

	EnableIpv6 *string `json:"EnableIpv6,omitempty" name:"EnableIpv6"`
	// 设置后，只装机智能网卡，否则安装智能网卡+黑石主机。

	OnlySoc *bool `json:"OnlySoc,omitempty" name:"OnlySoc"`
	// 服务器类型（3 黑石主机 or 4 黑石智能网卡）

	SvrType *string `json:"SvrType,omitempty" name:"SvrType"`
	// 关联智能网卡和黑石主机。

	SvrRelatedSn *string `json:"SvrRelatedSn,omitempty" name:"SvrRelatedSn"`
	// 是否清理数据盘（0：不清理，8：清理）

	PartitionFlag *int64 `json:"PartitionFlag,omitempty" name:"PartitionFlag"`
}

type DeleteAlarmStrategy struct {
	// 策略MD5，删除策略时给定

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 0->服务器，1->网络设备，2->网络专线，3->网络出口，4->自定义告警项，5->网络质量

	Type *string `json:"Type,omitempty" name:"Type"`
	// CI项属性，比如"机房=深圳"、"负责人=ABC"，支持多个维度组合，维度之间用逗号隔开。如"机房=深圳,机型=DELL"。type=4时为空

	CiAttr *int64 `json:"CiAttr,omitempty" name:"CiAttr"`
	// 命令空间，比如"tgw_custom"(TGW定制化)、"nas_group"(NAS集群)。type=4时有效

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 告警指标类型，0->数值型，1->字符型

	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`
	// 告警指标名。type=4时无效

	AlarmName *int64 `json:"AlarmName,omitempty" name:"AlarmName"`
	// 字符型告警取值为-1。数值判断方法，-1->无效，0->大于，1->小于，2->大于等于，3->小于等于，4->and组合，5->or组合

	JudgeMethod *int64 `json:"JudgeMethod,omitempty" name:"JudgeMethod"`
	// 字符型告警取值为空。-1 -> 无效。若判断方法为0-3，则该字段为阈值(整型)；若为4-5，则填充规则如下：type1-val1，type2-val2，其中type取值同判断类型，val为具体值

	JudgeValue *int64 `json:"JudgeValue,omitempty" name:"JudgeValue"`
}

type AddServerLabelInfo struct {
	// '设备固资编号'

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// '服务器SN(虚拟机此处为UUID,该值由云平提供和更新)'

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// '新增服务器标签'

	SvrLabel *string `json:"SvrLabel,omitempty" name:"SvrLabel"`
}

type CustomScripts struct {
	// 自定义脚本名

	CustomScriptName *string `json:"CustomScriptName,omitempty" name:"CustomScriptName"`
	// 自定义脚本加入的脚本集

	CustomScriptSetName *string `json:"CustomScriptSetName,omitempty" name:"CustomScriptSetName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type Overview struct {
	// 镜像类型

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 物理服务器状态

	SvrCurrentStatus *SvrCurrentStatus `json:"SvrCurrentStatus,omitempty" name:"SvrCurrentStatus"`
	// 物理服务器类型

	SvrType *string `json:"SvrType,omitempty" name:"SvrType"`
	// 资源容量监控

	CapacityAgent *string `json:"CapacityAgent,omitempty" name:"CapacityAgent"`
}

type TemplateActionParams struct {
	// StartId

	StartId *int64 `json:"StartId,omitempty" name:"StartId"`
	// Count

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// NeedTotal

	NeedTotal *int64 `json:"NeedTotal,omitempty" name:"NeedTotal"`
}

type ModifyRAIDResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改RAID出参信息

		DataSet *string `json:"DataSet,omitempty" name:"DataSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRAIDResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRAIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateServerWanIP struct {
	// 服务器资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器需要分配IP

	SvrIpWanted *string `json:"SvrIpWanted,omitempty" name:"SvrIpWanted"`
}

type RelocateServerCancelRequest struct {
	*tchttp.BaseRequest

	// Scheme

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"relocate_cancel"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 取消搬迁信息

	RelocateCancelServers []*RelocateCancelServer `json:"RelocateCancelServers,omitempty" name:"RelocateCancelServers"`
}

func (r *RelocateServerCancelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelocateServerCancelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerAllocateLanIPExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 请求动作

	Op *string `json:"Op,omitempty" name:"Op"`
	// 物理服务器分配内网IP

	ServersAllocateLanIp []*ServersAllocateLanIp `json:"ServersAllocateLanIp,omitempty" name:"ServersAllocateLanIp"`
}

func (r *ServerAllocateLanIPExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerAllocateLanIPExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Condition struct {
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 固资号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// Lanip

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
}

type ModifyImageSet struct {
	// 返回详情

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 返回信息

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 返回结果； 0:成功/1:失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
}

type ServerRecycleLanIPInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// SvrLanIp

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
}

type ServerRecyVirtualIPSet struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// SvrVirtualLanIp

	SvrVirtualLanIp *string `json:"SvrVirtualLanIp,omitempty" name:"SvrVirtualLanIp"`
	// SvrVirtualWanIp

	SvrVirtualWanIp *string `json:"SvrVirtualWanIp,omitempty" name:"SvrVirtualWanIp"`
}

type DescribeServerHarddisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数量

		ReturnNum *int64 `json:"ReturnNum,omitempty" name:"ReturnNum"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回硬盘列表

		ServerHardDiskSet []*ServerHardDisk `json:"ServerHardDiskSet,omitempty" name:"ServerHardDiskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServerHarddisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServerHarddisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServersSpecialRecycleLanIp struct {
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// SvrLanIp

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
}

type RenameServer struct {
	// 修改内容

	RenameServerModify *RenameServerModify `json:"RenameServerModify,omitempty" name:"RenameServerModify"`
	// 修改条件

	RenameServerCondition *RenameServerCondition `json:"RenameServerCondition,omitempty" name:"RenameServerCondition"`
}

type ServerSpecialAllocateLanIPInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type ModifyCustomScriptSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomScriptSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomScriptSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerAllocVirtualIPSet struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// SvrAllocateGateway

	SvrAllocateGateway *string `json:"SvrAllocateGateway,omitempty" name:"SvrAllocateGateway"`
	// SvrAllocateMask

	SvrAllocateMask *string `json:"SvrAllocateMask,omitempty" name:"SvrAllocateMask"`
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// SvrVirtualLanIp

	SvrVirtualLanIp *string `json:"SvrVirtualLanIp,omitempty" name:"SvrVirtualLanIp"`
}

type DeleteImageRequest struct {
	*tchttp.BaseRequest

	// "Op": "delete_image"

	Op *string `json:"Op,omitempty" name:"Op"`
	// "DeleteImageParams":[{"ImageId": "ins-123456"},{"ImageId": "ins-56789"}

	DeleteImageParams *DeleteImageParams `json:"DeleteImageParams,omitempty" name:"DeleteImageParams"`
}

func (r *DeleteImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageRequest struct {
	*tchttp.BaseRequest

	// "Op": "modify_image"

	Op *string `json:"Op,omitempty" name:"Op"`
	// "ModifyImageParams":[{"Condition":{"ImageId": "ins-123456"},"Modify":{"ImageName": "image_modify_test","ImageDescribe": "镜像修改","SystemArch":"arm64","OsType":"windows"}}]

	ModifyImageParams *ModifyImageParams `json:"ModifyImageParams,omitempty" name:"ModifyImageParams"`
}

func (r *ModifyImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOutbandStrategyData struct {
	// 自定义带外策略具体内容

	Strategies []*AddOutbandStrategies `json:"Strategies,omitempty" name:"Strategies"`
}

type OSDictionarySet struct {
	// 镜像名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 镜像来源：pxe/igniter

	Source *string `json:"Source,omitempty" name:"Source"`
}

type OutbandInfo struct {
	// DhcpIp

	DhcpIp *string `json:"DhcpIp,omitempty" name:"DhcpIp"`
	// Password

	Password *string `json:"Password,omitempty" name:"Password"`
	// Sn

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// User

	User *string `json:"User,omitempty" name:"User"`
}

type AllocateServerSpecialWanIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分配服务器特殊外网IP结果信息

		SvrSpecAllocWanIPSet []*SvrSpecAllocWanIP `json:"SvrSpecAllocWanIPSet,omitempty" name:"SvrSpecAllocWanIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateServerSpecialWanIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateServerSpecialWanIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImagePathResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像地址列表

		ImagePath *string `json:"ImagePath,omitempty" name:"ImagePath"`
		// 返回结果;0 成功；1失败

		Result *int64 `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImagePathResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagePathResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVMListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除虚拟机出参

		VMDeleteSet []*ServerAllocateLanIPInfo `json:"VMDeleteSet,omitempty" name:"VMDeleteSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVMListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVMListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleServerVirtualIPRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 虚拟内网："Op":"recycle_virtual_lan_ip"     虚拟外网： "Op":"recycle_virtual_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 回收物理服务器内网外网信息集合

	RecycleServerVirtualIPSet []*VirtualLanOrWanIP `json:"RecycleServerVirtualIPSet,omitempty" name:"RecycleServerVirtualIPSet"`
}

func (r *RecycleServerVirtualIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleServerVirtualIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddIdcDatas struct {
	// 数据中心ID

	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
	// 序列ID

	IndexId *int64 `json:"IndexId,omitempty" name:"IndexId"`
	// IDC名称

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// 设备名称

	DeviceRackName *string `json:"DeviceRackName,omitempty" name:"DeviceRackName"`
	// 设备ID

	DevicePosId *string `json:"DevicePosId,omitempty" name:"DevicePosId"`
}

type ConfigServerDefSet struct {
	// 服务器SN

	SN *string `json:"SN,omitempty" name:"SN"`
	// 服务器架构

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	// 内网IP信息

	Lan []*ConfigServerDefSetLan `json:"Lan,omitempty" name:"Lan"`
	// ZoneName

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// IdcName

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// 是否强制初始化

	Force *bool `json:"Force,omitempty" name:"Force"`
	// 配置初始化后的脚本集名称

	PreInstallScript *string `json:"PreInstallScript,omitempty" name:"PreInstallScript"`
}

type AddLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回串中的DataSet节点，规范协议请求

		DataSet *AddLabelDataSet `json:"DataSet,omitempty" name:"DataSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLabelRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"label"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 操作类型："Op":"delete"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 修改的标签内容

	DeleteLabelSet []*LabelConditionInfo `json:"DeleteLabelSet,omitempty" name:"DeleteLabelSet"`
}

func (r *DeleteLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LineXflowInfo struct {
	// 管理IP

	AdminIp *string `json:"AdminIp,omitempty" name:"AdminIp"`
	// 接口索引

	Ifindex *int64 `json:"Ifindex,omitempty" name:"Ifindex"`
	// 专线ID

	LineId *int64 `json:"LineId,omitempty" name:"LineId"`
	// 其他IP

	OtherIp *string `json:"OtherIp,omitempty" name:"OtherIp"`
}

type ModifyImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// "ImageSet":[{"Detail": "ok","Result": 0}]

		ImageSet *ModifyImageSet `json:"ImageSet,omitempty" name:"ImageSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLabelRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"label"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 操作类型："Op":"modify"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 修改的标签内容

	ModifyLibelSet *ModifyLabelInfo `json:"ModifyLibelSet,omitempty" name:"ModifyLibelSet"`
}

func (r *ModifyLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelsResultInfo struct {
	// 标签ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 标签键

	LabelKey *string `json:"LabelKey,omitempty" name:"LabelKey"`
	// 标签值

	LabelValue *string `json:"LabelValue,omitempty" name:"LabelValue"`
}

type AddRAIDRequest struct {
	*tchttp.BaseRequest

	// "Op": "add_raid"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 参考"AddRAIDParams": {"Name":"Linux"}

	AddRAIDParams *AddRAIDParam `json:"AddRAIDParams,omitempty" name:"AddRAIDParams"`
}

func (r *AddRAIDRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddRAIDRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImagePathRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImagePathRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagePathRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SocPerformServerTaskExRequest struct {
	*tchttp.BaseRequest

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// SvrIdcName

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 操作码：1（开机），2（关机），3（重启）

	OpCode *string `json:"OpCode,omitempty" name:"OpCode"`
	// 操作的服务器信息

	Servers []*SocPerformServerTaskEsServers `json:"Servers,omitempty" name:"Servers"`
	// SvrIdcId

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
}

func (r *SocPerformServerTaskExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SocPerformServerTaskExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLabelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLabelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IPDetailResultInfo struct {
	// AssetId

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// Ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// IpStatusId

	IpStatusId *string `json:"IpStatusId,omitempty" name:"IpStatusId"`
	// NetsegmentName

	NetsegmentName *string `json:"NetsegmentName,omitempty" name:"NetsegmentName"`
}

type RecycleVMVirtualIPRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server_vm"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 虚拟内网："Op":"recycle_virtual_lan_ip"     虚拟外网： "Op":"recycle_virtual_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 回收VM虚拟内网外网信息集合

	RecycleVMVirtualIPSet []*VirtualLanOrWanIP `json:"RecycleVMVirtualIPSet,omitempty" name:"RecycleVMVirtualIPSet"`
}

func (r *RecycleVMVirtualIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleVMVirtualIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddRAIDParam struct {
	// RAID名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type CmdbModify struct {
	// 网络AssetId

	NetdevAssetId *string `json:"NetdevAssetId,omitempty" name:"NetdevAssetId"`
	// 网络设备固资号

	NetdevSn *string `json:"NetdevSn,omitempty" name:"NetdevSn"`
	// NetdevEngine

	NetdevEngine *string `json:"NetdevEngine,omitempty" name:"NetdevEngine"`
	// 网络设备名称

	NetdevName *string `json:"NetdevName,omitempty" name:"NetdevName"`
	// 网络设备类型

	NetdevType *string `json:"NetdevType,omitempty" name:"NetdevType"`
	// 网络设备厂商

	NetdevFunc *string `json:"NetdevFunc,omitempty" name:"NetdevFunc"`
	// NetdevPro

	NetdevPro *string `json:"NetdevPro,omitempty" name:"NetdevPro"`
	// 网络设备模型

	NetdevModel *string `json:"NetdevModel,omitempty" name:"NetdevModel"`
	// 网络设备系统

	NetdevOs *string `json:"NetdevOs,omitempty" name:"NetdevOs"`
	// 网络设备专线

	NetdevIdc *string `json:"NetdevIdc,omitempty" name:"NetdevIdc"`
	// 网络设备idcid

	NetdevIdcId *string `json:"NetdevIdcId,omitempty" name:"NetdevIdcId"`
	// 网络设备rack名称

	NetdevRackName *string `json:"NetdevRackName,omitempty" name:"NetdevRackName"`
	// 网络设备当前状态

	NetdevCurrentStatus *string `json:"NetdevCurrentStatus,omitempty" name:"NetdevCurrentStatus"`
	// 网络设备主ip

	NetdevAdminIp *string `json:"NetdevAdminIp,omitempty" name:"NetdevAdminIp"`
	// 网络设备其他ip

	NetdevOtherIp *string `json:"NetdevOtherIp,omitempty" name:"NetdevOtherIp"`
	// 可用地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 可用区域名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 网络设备升级时间

	NetdevUpdateTime *string `json:"NetdevUpdateTime,omitempty" name:"NetdevUpdateTime"`
}

type DeleteRAIDRequest struct {
	*tchttp.BaseRequest

	// "Op": "del_raid"

	Op *string `json:"Op,omitempty" name:"Op"`
	// "DeleteRAIDParams": {"Id":3}

	DeleteRAIDParams *DeleteRAIDParam `json:"DeleteRAIDParams,omitempty" name:"DeleteRAIDParams"`
}

func (r *DeleteRAIDRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRAIDRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePhysvrsOverviewRequest struct {
	*tchttp.BaseRequest

	// 筛选AZ

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePhysvrsOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePhysvrsOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleVMLanIP struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器序列号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 服务器内网IP

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
}

type DeleteSyslogKeywordCond struct {
	// 厂商

	NetdevVendor *string `json:"NetdevVendor,omitempty" name:"NetdevVendor"`
	// 设备类型

	NetdevType *string `json:"NetdevType,omitempty" name:"NetdevType"`
	// 设备功能

	NetdevFun *string `json:"NetdevFun,omitempty" name:"NetdevFun"`
	// 型号

	NetdevModel *string `json:"NetdevModel,omitempty" name:"NetdevModel"`
	// 关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type DhcpConfSet struct {
	// TaskId

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

type ImageSet struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// ImageInfo

	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`
}

type ModifyLabelModifyInfo struct {
	// 修改后的标签键键

	LabelKey *string `json:"LabelKey,omitempty" name:"LabelKey"`
	// 修改后的标签键键

	LabelValue *string `json:"LabelValue,omitempty" name:"LabelValue"`
}

type DescribeCustomScriptsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li> CustomScriptName - String List - 是否必填：否 -（过滤条件）按照脚本名过滤。</li>
	// <li> CustomScriptSetName - String List - 是否必填：否 - （过滤条件）通过脚本集名字过滤，如果需要没有加入脚本集的脚本，Values增加""，例如["xxx", ""]。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCustomScriptsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomScriptsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleServerVirtualIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 回收物理服务器虚拟内外网ip出参

		RecycleServerVirtualIP []*ServerRecyVirtualIPSet `json:"RecycleServerVirtualIP,omitempty" name:"RecycleServerVirtualIP"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecycleServerVirtualIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleServerVirtualIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigServerDefData struct {
	// SvrCurrentStatus

	SvrCurrentStatus *string `json:"SvrCurrentStatus,omitempty" name:"SvrCurrentStatus"`
	// SvrIdcName

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// SvrSn

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// TaskId

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type ModifyServer struct {
	// 修改条件

	Condition *ModifyServerCondition `json:"Condition,omitempty" name:"Condition"`
	// 修改具体字段内容

	Modify *ModifyServerPara `json:"Modify,omitempty" name:"Modify"`
}

type RelocateFinishServer struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// IDC名称

	SvrRelocateIdcName *string `json:"SvrRelocateIdcName,omitempty" name:"SvrRelocateIdcName"`
	// 机架名称

	SvrRelocateRackName *string `json:"SvrRelocateRackName,omitempty" name:"SvrRelocateRackName"`
	// 机架位置

	SvrRelocatePosId *string `json:"SvrRelocatePosId,omitempty" name:"SvrRelocatePosId"`
}

type RecycleVMLanIPListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server_vm"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"recycle_lan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 回收VM内网IP

	RecycleVMLanIPSet []*RecycleVMLanIP `json:"RecycleVMLanIPSet,omitempty" name:"RecycleVMLanIPSet"`
}

func (r *RecycleVMLanIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleVMLanIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServersRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server or server_vm"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 操作类型："Op":"dispose"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 删除服务器参数集合

	DeleteServerSet []*DeleteServer `json:"DeleteServerSet,omitempty" name:"DeleteServerSet"`
}

func (r *DeleteServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReinstallOsServers struct {
	// 服务器信息

	Servers []*ReinstallOsInfo `json:"Servers,omitempty" name:"Servers"`
}

type ServerProcessModifyCondition struct {
	// 服务器资产Id

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type ModifyOutbandStrategyRequest struct {
	*tchttp.BaseRequest

	// 请求串类型说明。
	// 固定值： "Scheme":"outband_passwd_strategy"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方systemid	此ID为请求方向cmdb注册的权限ID，由cmdb系统管理员提供和管理。
	// 默认值："SystemId":"sid_test"

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方systemkey	和上面systemid类似，由cmdb系统管理员提供和管理，cmdb侧会根据id和key以及调用方ip三者共同确定请求方的合法性。
	// 默认值："SystemKey":"sid_test"

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 修改自定义带外密码策略详细内容

	Data *OutbandStrategyData `json:"Data,omitempty" name:"Data"`
}

func (r *ModifyOutbandStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOutbandStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PortModify struct {
	// 服务器端口

	SvrMonPort *string `json:"SvrMonPort,omitempty" name:"SvrMonPort"`
}

type ImportCmdbServersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导入物理服务器接口返回信息

		ImportServersResult []*ImportServersResult `json:"ImportServersResult,omitempty" name:"ImportServersResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportCmdbServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportCmdbServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PerformServerTaskExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// PerformServerTaskSet

		PerformServerTaskSet []*string `json:"PerformServerTaskSet,omitempty" name:"PerformServerTaskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PerformServerTaskExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PerformServerTaskExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageFieldsEnumSet struct {
	// 操作系统枚举值

	OsType []*string `json:"OsType,omitempty" name:"OsType"`
	// 系统架构枚举值

	SystemArch []*string `json:"SystemArch,omitempty" name:"SystemArch"`
	// 系统平台枚举值

	SystemPlatform []*string `json:"SystemPlatform,omitempty" name:"SystemPlatform"`
	// 镜像格式枚举值

	ImageType []*string `json:"ImageType,omitempty" name:"ImageType"`
	// 可用区枚举值

	ZoneName []*string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 镜像格式;IOS/sqfs/等

	ImageFormat []*string `json:"ImageFormat,omitempty" name:"ImageFormat"`
}

type ServersAllocateLanIp struct {
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// SvrIpWanted

	SvrIpWanted *string `json:"SvrIpWanted,omitempty" name:"SvrIpWanted"`
	// Vlan

	Vlan *string `json:"Vlan,omitempty" name:"Vlan"`
}

type DeleteRAIDResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除RAID出参信息

		DataSet *string `json:"DataSet,omitempty" name:"DataSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRAIDResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRAIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePhysvrsListExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求动作

	Op *string `json:"Op,omitempty" name:"Op"`
	// 查询条件

	FieldFilters *string `json:"FieldFilters,omitempty" name:"FieldFilters"`
	// RoughItem

	RoughItem *string `json:"RoughItem,omitempty" name:"RoughItem"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePhysvrsListExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePhysvrsListExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomScriptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomScriptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddIPResourceInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// IndexId

	IndexId *int64 `json:"IndexId,omitempty" name:"IndexId"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
}

type DescribeSysLogKeywordCond struct {
	// 描述，可选

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 厂商,可选

	NetdevVendor *string `json:"NetdevVendor,omitempty" name:"NetdevVendor"`
	// 类型,可选

	NetdevType *string `json:"NetdevType,omitempty" name:"NetdevType"`
	// 功能,可选

	NetdevFun *string `json:"NetdevFun,omitempty" name:"NetdevFun"`
	// 型号,可选

	NetdevModel *string `json:"NetdevModel,omitempty" name:"NetdevModel"`
	// 告警类型，可选

	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`
	// 关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type OsParam struct {
	// 编号

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 系统名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type WithdrawServerSet struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *string `json:"Result,omitempty" name:"Result"`
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type DescribeAllOSListRequest struct {
	*tchttp.BaseRequest

	// "Op":"query_os"

	Op *string `json:"Op,omitempty" name:"Op"`
	// [{"Name":"Name", "Values":["TencentOS Server 2.2 (TK2-0053-1)"]}

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAllOSListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllOSListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LineXflowParams struct {
	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 显示条目数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// "Filters": [{"Name":"LineId","Values":["1"]},{"Name":"IfIndex","Values": ["45"]},{"Name":"ip","Values":["10.10.127.234"]}]

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type ModifyImageModify struct {
	// 修改后的镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 修改后的系统架构

	SystemArch *string `json:"SystemArch,omitempty" name:"SystemArch"`
	// 修改后的操作系统类型

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 修改后的镜像描述

	ImageDescribe *string `json:"ImageDescribe,omitempty" name:"ImageDescribe"`
	// 修改后的系统平台

	SystemPlatform *string `json:"SystemPlatform,omitempty" name:"SystemPlatform"`
	// 修改后的系统版本

	SystemVersions *string `json:"SystemVersions,omitempty" name:"SystemVersions"`
	// 镜像使用服务器设备类型(0:通用服务器;1:黑石服务器)

	AvailableModel *uint64 `json:"AvailableModel,omitempty" name:"AvailableModel"`
}

type RecycleServerSpecialWanIP struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器IP

	SvrIp *string `json:"SvrIp,omitempty" name:"SvrIp"`
}

type IdcInfoItems struct {
	// IdcCondition

	IdcCondition *IdcCondition `json:"IdcCondition,omitempty" name:"IdcCondition"`
	// IdcModify

	IdcModify *IdcModify `json:"IdcModify,omitempty" name:"IdcModify"`
}

type WanData struct {
	// 网卡

	Nic *string `json:"Nic,omitempty" name:"Nic"`
	// IP地址

	Ipaddress *string `json:"Ipaddress,omitempty" name:"Ipaddress"`
}

type DescribeServerRelatedNetdevicesRequest struct {
	*tchttp.BaseRequest

	// 服务器机架名

	SvrRackName *string `json:"SvrRackName,omitempty" name:"SvrRackName"`
	// 服务器机架位置编号

	SvrPosId *string `json:"SvrPosId,omitempty" name:"SvrPosId"`
}

func (r *DescribeServerRelatedNetdevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServerRelatedNetdevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSyslogKeyword struct {
	// 关键字描述

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 厂商

	NetdevVendor *string `json:"NetdevVendor,omitempty" name:"NetdevVendor"`
	// 类型

	NetdevType *string `json:"NetdevType,omitempty" name:"NetdevType"`
	// 功能

	NetdevFun *string `json:"NetdevFun,omitempty" name:"NetdevFun"`
	// 型号

	NetdevModel *string `json:"NetdevModel,omitempty" name:"NetdevModel"`
	// 关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 告警类型

	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`
}

type NetDeviceDatas struct {
	// IndexId

	IndexId *int64 `json:"IndexId,omitempty" name:"IndexId"`
	// NetdevAssetId

	NetdevAssetId *string `json:"NetdevAssetId,omitempty" name:"NetdevAssetId"`
	// NetdevSn

	NetdevSn *string `json:"NetdevSn,omitempty" name:"NetdevSn"`
	// NetdevEngine

	NetdevEngine *string `json:"NetdevEngine,omitempty" name:"NetdevEngine"`
	// NetdevName

	NetdevName *string `json:"NetdevName,omitempty" name:"NetdevName"`
	// NetdevType

	NetdevType *string `json:"NetdevType,omitempty" name:"NetdevType"`
	// NetdevFunc

	NetdevFunc *string `json:"NetdevFunc,omitempty" name:"NetdevFunc"`
	// NetdevPro

	NetdevPro *string `json:"NetdevPro,omitempty" name:"NetdevPro"`
	// NetdevModel

	NetdevModel *string `json:"NetdevModel,omitempty" name:"NetdevModel"`
	// NetdevOs

	NetdevOs *string `json:"NetdevOs,omitempty" name:"NetdevOs"`
	// NetdevIdc

	NetdevIdc *string `json:"NetdevIdc,omitempty" name:"NetdevIdc"`
	// NetdevIdcId

	NetdevIdcId *string `json:"NetdevIdcId,omitempty" name:"NetdevIdcId"`
	// NetdevRackName

	NetdevRackName *string `json:"NetdevRackName,omitempty" name:"NetdevRackName"`
	// NetdevCurrentStatus

	NetdevCurrentStatus *string `json:"NetdevCurrentStatus,omitempty" name:"NetdevCurrentStatus"`
	// NetdevAdminIp

	NetdevAdminIp *string `json:"NetdevAdminIp,omitempty" name:"NetdevAdminIp"`
	// NetdevOtherIp

	NetdevOtherIp *string `json:"NetdevOtherIp,omitempty" name:"NetdevOtherIp"`
}

type RAIDInfo struct {
	// 编号

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// RAID名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type RelocateServerCancelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 物理服务器取消搬迁接口出参

		CancelRelocateServerSet []*ServerAllocateLanIPInfo `json:"CancelRelocateServerSet,omitempty" name:"CancelRelocateServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RelocateServerCancelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelocateServerCancelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SocPerformServerTaskExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SocPerformServerTaskSet

		SocPerformServerTaskSet []*string `json:"SocPerformServerTaskSet,omitempty" name:"SocPerformServerTaskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SocPerformServerTaskExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SocPerformServerTaskExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmStrategyMd5 struct {
	// 需要删除的MD5列表

	StrategyMd5 []*string `json:"StrategyMd5,omitempty" name:"StrategyMd5"`
}

type DescribePhysvrsOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 物理服务器资源概览出参

		Overview *string `json:"Overview,omitempty" name:"Overview"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePhysvrsOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePhysvrsOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IdcCondition struct {
	// SpeLineId

	SpeLineId *string `json:"SpeLineId,omitempty" name:"SpeLineId"`
}

type DeleteCustomScriptsRequest struct {
	*tchttp.BaseRequest

	// 自定义脚本的名称

	CustomScriptName []*string `json:"CustomScriptName,omitempty" name:"CustomScriptName"`
}

func (r *DeleteCustomScriptsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomScriptsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageListRequest struct {
	*tchttp.BaseRequest

	// "Op": "query_image"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 过滤条件信息

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制返回个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeImageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Laninfo struct {
	// Nic

	Nic *string `json:"Nic,omitempty" name:"Nic"`
	// Ipaddress

	Ipaddress *string `json:"Ipaddress,omitempty" name:"Ipaddress"`
	// Netmask

	Netmask *string `json:"Netmask,omitempty" name:"Netmask"`
	// Gateway

	Gateway *string `json:"Gateway,omitempty" name:"Gateway"`
}

type NetdeviceSyslogParams struct {
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件

	Filters *Filter `json:"Filters,omitempty" name:"Filters"`
}

type Strategies struct {
	// 返回结果详细说明

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	//  运营端 总览 产品运营 平台运营 平台运维 帮助与文档 语言 平台管理   v_fwei  云API管理 icon 租户端 icon 运营端 接口管理 服务管理 接口版本 接口分类 接口错误码 接口参数 接口后端地址 变更历史 接口日志 接口统计 API_Explorer SDK下载 云API规范校验 收起 接口编辑(运营端) dcos 2020-04-21 DescribeLabel 请求参数的"Action" 查询标签 标签管理--查询标签  json 参数名称	是否必填	类型	数组类型	是否容许NULL	描述	操作 Scheme  string 字符串型 "Scheme":"label" SystemId  string 字符串型 系统ID SystemKey  string 字符串型 系统Key Op  string 字符串型 操作类型："Op":"query" Limit  int64 有符号整型 返回数量 Offset  int64 有符号整型 偏移 Filters  Filter 查询条件 ResultItem  string 字符串型 返回参数  适用于低频高敏感可容忍高时延的接口，选项为是将增加安全性，但会大幅增加接口时延，同时业务后端需要改造配合该功能。  参数名称	是否必填	类型	数组类型	是否容许NULL	描述	操作 DataSet  Labels 返回串中的DataSet节点，规范协议请求 ReturnNum  int64 有符号整型 本次查询返回的数据条数 TotalCount  int64 有符号整型 满足本次查询条件的数据总条数  请选择 请选择 错误码是由“一级错误码”和“二级错误码”组合而成，例如"InvalidParameter.ImageId"。一级错误码由系统统一规定，二级错误码业务可自定义  既验证签名又鉴权 签名 ，是用于判断请求是否来自合法的云帐号，通过比对客户端和服务端对请求串结合SecretKey的hash串是否一致实现。鉴权，即CAM鉴权，在签名通过证明其合法身份后，再进一步判断该账户(uin)，是否能操作(Action）指定的资源(Resource)。API框架一定会验签名，但不一定会向CAM请求鉴权，复杂的业务可以在各自的后端服务向CAM服务请求鉴权 SigAndAuth 请选择 查询 * 形式如 qcs::${ApiModule}:${Region}:uin/:resourceName/${ResourceName}, * 表示接口级别鉴权  ocloud-dcos-cgw-apiv3-dcos 多地址 地域	后端地址	操作  重庆 http://ocloud-dcos-cgw-apiv3-dcos.chongqing.yfm18.tcepoc.fsphere.cn:80/dcos-cgw-apiv3/dcos  无  20 次/每秒 (账号自身uin+action，账号自身Uin的含义是：SecretId对应的Uin。即如果SecretId对应的Uin是主账号的Uin，则填主账号的Uin，如果SecretId对应的Uin是子账号的Uin，则填子账号的Uin） 形式如 uin1:多少次每秒;uin2:多少次每秒 举例 2345346:30;345349056:150 注意不是倍数！！！  公开接口 如可公开给租户或管理员进行调用，则设置为“公开接口”；如租户或管理员无需使用，则设置为“内部接口”。  DCOS服务  5000 1024 请求包大小请设置在1024kB到10240KB之间,即1M到10M,如果不填默认是1M 编辑接口参数 dcos 2020-04-21 Labels 查询标签返回参数 参数名称	是否必填	类型	数组类型	是否容许NULL	描述	操作 Detail  string 字符串型 返回结果详细说明 Result  int64 有符号整型 返回结果值，0成功， 1失败 ResultInfo  LabelsResultInfo 标签的详细信息

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 自定义带外密码策略详细信息

	ResultInfo *StrategiesResultInfo `json:"ResultInfo,omitempty" name:"ResultInfo"`
}

type SvrSpecAllocWanIP struct {
	// 结果详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 资产Id

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type AddImageRequest struct {
	*tchttp.BaseRequest

	// "Op": "add_image"

	Op *string `json:"Op,omitempty" name:"Op"`
	// "AddImageParams"{"ImageUrl":os_url,"RegionName": "region_name","ImageName": "test_os","OsType": "Linux","SystemArch": "X86","SystemPlatform" : "CentOS",,"SystemVersions": 7,"ImageDescribe": "示例文字"}

	AddImageParams *AddImageParams `json:"AddImageParams,omitempty" name:"AddImageParams"`
}

func (r *AddImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Nodesvr struct {
	// IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// NeedRestartDhcp

	NeedRestartDhcp *string `json:"NeedRestartDhcp,omitempty" name:"NeedRestartDhcp"`
}

type SocPerformServerTaskEsServers struct {
	// Sn

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// idcid

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// 服务器ip

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
	// idcname

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
}

type SpecialCondition struct {
	// SpeLineId

	SpeLineId *string `json:"SpeLineId,omitempty" name:"SpeLineId"`
}

type TargetIp struct {
	// NQ网关

	NqAgent *string `json:"NqAgent,omitempty" name:"NqAgent"`
	// IP列表

	Ip []*string `json:"Ip,omitempty" name:"Ip"`
	// IP类型

	IpType *string `json:"IpType,omitempty" name:"IpType"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

type DescribeOSListRequest struct {
	*tchttp.BaseRequest

	// "Op": "query_os"

	Op *string `json:"Op,omitempty" name:"Op"`
	// "Filters": [{"Name": "Arch", "Values": ["x86"]}]

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeOSListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOSListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerSpecialRecycleLanIPExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 请求动作

	Op *string `json:"Op,omitempty" name:"Op"`
	// 用户

	User *string `json:"User,omitempty" name:"User"`
	// 物理服务器批量回收相应虚拟内网段IP接口

	ServersSpecialRecycleLanIp []*ServersSpecialRecycleLanIp `json:"ServersSpecialRecycleLanIp,omitempty" name:"ServersSpecialRecycleLanIp"`
}

func (r *ServerSpecialRecycleLanIPExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerSpecialRecycleLanIPExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProcessModify struct {
	// 进程名

	SvrMonProc *string `json:"SvrMonProc,omitempty" name:"SvrMonProc"`
}

type RAIDParams struct {
	// 编号

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// RAID名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type AddOutbandStrategyRequest struct {
	*tchttp.BaseRequest

	// 请求方systemid	此ID为请求方向cmdb注册的权限ID，由cmdb系统管理员提供和管理。
	// 默认值："SystemId":"sid_test"

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方systemkey	和上面systemid类似，由cmdb系统管理员提供和管理，cmdb侧会根据id和key以及调用方ip三者共同确定请求方的合法性。
	// 默认值："SystemKey":"sid_test"

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 新增自定义带外密码策略详细内容

	Data *AddOutbandStrategyData `json:"Data,omitempty" name:"Data"`
	// 请求串类型说明。
	// 固定值： "Scheme":"outband_passwd_strategy"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
}

func (r *AddOutbandStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOutbandStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIdcLineInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// SpeLineId

	SpeLineId *string `json:"SpeLineId,omitempty" name:"SpeLineId"`
}

type DescribeOutbandInfoExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// OutbandInfo

		OutbandInfo []*OutbandInfo `json:"OutbandInfo,omitempty" name:"OutbandInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOutbandInfoExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOutbandInfoExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddImageParams struct {
	// 镜像路径

	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
	// 镜像名字

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 操作系统类型

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 系统架构

	SystemArch *string `json:"SystemArch,omitempty" name:"SystemArch"`
	// 系统平台

	SystemPlatform *string `json:"SystemPlatform,omitempty" name:"SystemPlatform"`
	// 系统版本

	SystemVersions *string `json:"SystemVersions,omitempty" name:"SystemVersions"`
	// 镜像描述

	ImageDescribe *string `json:"ImageDescribe,omitempty" name:"ImageDescribe"`
	// ZoneName

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 镜像格式(iso/sqfs)

	ImageFormat *string `json:"ImageFormat,omitempty" name:"ImageFormat"`
	// 镜像使用服务器设备类型(0:通用服务器;1:黑石服务器)

	AvailableModel *int64 `json:"AvailableModel,omitempty" name:"AvailableModel"`
	// IdcName

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
}

type NetDeviceData struct {
	// Netdevices

	NetdevicesInfo []*NetdevicesInfo `json:"NetdevicesInfo,omitempty" name:"NetdevicesInfo"`
}

type SvrSpecRecyWanIP struct {
	// 结果详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 结果错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器IP

	SvrIp *string `json:"SvrIp,omitempty" name:"SvrIp"`
}

type DescribeServersRequest struct {
	*tchttp.BaseRequest

	// "scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// "system_id":"sid_test"

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// "system_key":"sid_test"

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"query"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// //(根据需要填充需要返回的字段名，不同字段名使用","隔开)

	ResultItem *string `json:"ResultItem,omitempty" name:"ResultItem"`
	// //(根据需要填充需要模糊查询的字段名，不同字段名使用","隔开，对应的字段名一定要在Filters里面体现)

	RoughItem *string `json:"RoughItem,omitempty" name:"RoughItem"`
	// 过滤条件："Filters":[{"Name":"SvrLanIp","Value":"10.21.12.139"},{"Name":"SvrAssetId","Value":""},{"Name":"SvrIdcName","Value":""},{"Name":"SvrLogicArea","Value":""},{"Name":"SvrCurrentStatus","Value":""},{"Name":"SvrBussiness","Value":""}]}

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomScriptSetRequest struct {
	*tchttp.BaseRequest

	// 自定义脚本集名

	CustomScriptSetName *string `json:"CustomScriptSetName,omitempty" name:"CustomScriptSetName"`
	// 重新命名的自定义脚本集名

	CustomScriptSetNameNew *string `json:"CustomScriptSetNameNew,omitempty" name:"CustomScriptSetNameNew"`
	// 加入脚本集中的自定义脚本名列表

	CustomScriptName []*string `json:"CustomScriptName,omitempty" name:"CustomScriptName"`
	// 脚本集类型(0:初始化脚本,1:后置脚本)

	CustomScriptSetType *int64 `json:"CustomScriptSetType,omitempty" name:"CustomScriptSetType"`
}

func (r *ModifyCustomScriptSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomScriptSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyServerCondition struct {
	// 设备固资编号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type AllocateServerWanIPListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"allocate_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 分配服务器外网IP

	AllocateServerWanIPSet []*AllocateServerWanIP `json:"AllocateServerWanIPSet,omitempty" name:"AllocateServerWanIPSet"`
}

func (r *AllocateServerWanIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateServerWanIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeploySubtaskStepsExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 部署过程各步骤日志信息

		DeploySubtaskInfo *string `json:"DeploySubtaskInfo,omitempty" name:"DeploySubtaskInfo"`
		// 装机工具

		DeployTool *string `json:"DeployTool,omitempty" name:"DeployTool"`
		// 查询的是否成功与前端匹配。0成功

		Errno *int64 `json:"Errno,omitempty" name:"Errno"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeploySubtaskStepsExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeploySubtaskStepsExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOutbandStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回串中的DataSet节点，规范协议请求

		DataSet *DeleteOutbandStrategyDataSet `json:"DataSet,omitempty" name:"DataSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOutbandStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOutbandStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServerRelatedNetdevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务器关联网络设备

		Netdeivces []*Netdevice `json:"Netdeivces,omitempty" name:"Netdeivces"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServerRelatedNetdevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServerRelatedNetdevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Netdeviceportlist struct {
	// Netdevice

	Netdevice *string `json:"Netdevice,omitempty" name:"Netdevice"`
	// Portname

	Portname *string `json:"Portname,omitempty" name:"Portname"`
}

type Rackinfo struct {
	// IdcDevices

	IdcDevices []*IdcDevices `json:"IdcDevices,omitempty" name:"IdcDevices"`
	// ZoneName

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// RackName

	RackName *string `json:"RackName,omitempty" name:"RackName"`
	// DevicesOnNum

	DevicesOnNum *int64 `json:"DevicesOnNum,omitempty" name:"DevicesOnNum"`
	// IdlePosNum

	IdlePosNum *int64 `json:"IdlePosNum,omitempty" name:"IdlePosNum"`
	// RackPosNum

	RackPosNum *int64 `json:"RackPosNum,omitempty" name:"RackPosNum"`
}

type DescribeOsDeployDebugInfoExRequest struct {
	*tchttp.BaseRequest

	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// os部署调试信息

	OsDeployDebugInfo []*OsDeployDebugInfo `json:"OsDeployDebugInfo,omitempty" name:"OsDeployDebugInfo"`
}

func (r *DescribeOsDeployDebugInfoExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOsDeployDebugInfoExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailedInfo struct {
	// 导入失败的物理服务器详细信息

	Servers []*ImportServerList `json:"Servers,omitempty" name:"Servers"`
	// 导入失败的总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type HarddiskInfo struct {
	// 容量

	HdCapacity *string `json:"HdCapacity,omitempty" name:"HdCapacity"`
	// 固件版本

	HdFw *string `json:"HdFw,omitempty" name:"HdFw"`
	// 接口类型

	HdInfType *string `json:"HdInfType,omitempty" name:"HdInfType"`
	// 介质

	HdMedium *string `json:"HdMedium,omitempty" name:"HdMedium"`
	// HdRp

	HdRp *string `json:"HdRp,omitempty" name:"HdRp"`
	// 槽位

	HdSlot *string `json:"HdSlot,omitempty" name:"HdSlot"`
	// SN编号

	HdSn *string `json:"HdSn,omitempty" name:"HdSn"`
	// 类型

	HdType *string `json:"HdType,omitempty" name:"HdType"`
	// 厂家

	HdVendor *string `json:"HdVendor,omitempty" name:"HdVendor"`
	// 编号

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 关联服务器

	RelatedSvr *string `json:"RelatedSvr,omitempty" name:"RelatedSvr"`
	// 地域

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 可用区

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 数据中心名称

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// 数据中心ID

	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
	// 服务器资产编号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type ServerHardDisk struct {
	// 返回详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 返回错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 硬盘信息

	HarddiskInfo *HarddiskInfo `json:"HarddiskInfo,omitempty" name:"HarddiskInfo"`
}

type AddLabelInfo struct {
	// 标签建

	LabelKey *string `json:"LabelKey,omitempty" name:"LabelKey"`
	// 标签值

	LabelValue *string `json:"LabelValue,omitempty" name:"LabelValue"`
}

type IdcExporResultInfo struct {
	// Comment

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// ExportDevice

	ExportDevice *string `json:"ExportDevice,omitempty" name:"ExportDevice"`
	// ExportDeviceIp

	ExportDeviceIp *string `json:"ExportDeviceIp,omitempty" name:"ExportDeviceIp"`
	// ExportDevicePort

	ExportDevicePort *string `json:"ExportDevicePort,omitempty" name:"ExportDevicePort"`
	// ExportLineId

	ExportLineId *string `json:"ExportLineId,omitempty" name:"ExportLineId"`
	// ExportLineName

	ExportLineName *string `json:"ExportLineName,omitempty" name:"ExportLineName"`
	// ExportLineStatus

	ExportLineStatus *string `json:"ExportLineStatus,omitempty" name:"ExportLineStatus"`
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// LineBusiAvailableSpeed

	LineBusiAvailableSpeed *string `json:"LineBusiAvailableSpeed,omitempty" name:"LineBusiAvailableSpeed"`
	// LineEndTime

	LineEndTime *string `json:"LineEndTime,omitempty" name:"LineEndTime"`
	// LineOperator

	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`
	// LineStartTime

	LineStartTime *string `json:"LineStartTime,omitempty" name:"LineStartTime"`
	// LineTechAvailableSpeed

	LineTechAvailableSpeed *string `json:"LineTechAvailableSpeed,omitempty" name:"LineTechAvailableSpeed"`
	// RemoteDeviceIp

	RemoteDeviceIp *string `json:"RemoteDeviceIp,omitempty" name:"RemoteDeviceIp"`
	// RemoteOperatorOwner

	RemoteOperatorOwner *string `json:"RemoteOperatorOwner,omitempty" name:"RemoteOperatorOwner"`
	// RemoteOperatorOwnerPhone

	RemoteOperatorOwnerPhone *string `json:"RemoteOperatorOwnerPhone,omitempty" name:"RemoteOperatorOwnerPhone"`
	// RemoteOperatorOwnerWarningPhone

	RemoteOperatorOwnerWarningPhone *string `json:"RemoteOperatorOwnerWarningPhone,omitempty" name:"RemoteOperatorOwnerWarningPhone"`
	// SlaNum

	SlaNum *string `json:"SlaNum,omitempty" name:"SlaNum"`
}

type AddServersRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 操作类型："Op":"add"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 增加服务器信息，具体参数在复杂类型中定义

	AddServerSet []*AddServer `json:"AddServerSet,omitempty" name:"AddServerSet"`
}

func (r *AddServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddServersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateSet struct {
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// TempId

	TempId *int64 `json:"TempId,omitempty" name:"TempId"`
	// TempName

	TempName *string `json:"TempName,omitempty" name:"TempName"`
}

type ReinstallOsExRequest struct {
	*tchttp.BaseRequest

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// idcid

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// 请求动作

	Command *string `json:"Command,omitempty" name:"Command"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 用户

	User *string `json:"User,omitempty" name:"User"`
	// 操作类型

	OperateType *int64 `json:"OperateType,omitempty" name:"OperateType"`
	// OS部署信息

	Data *ReinstallOsServers `json:"Data,omitempty" name:"Data"`
	// 选择装机工具：dcos或igniter

	DeployTool *string `json:"DeployTool,omitempty" name:"DeployTool"`
	// true:只装机智能网卡，false:安装智能网卡+黑石主机

	OnlySoc *bool `json:"OnlySoc,omitempty" name:"OnlySoc"`
	// 黑石主机和网卡关联的SN

	SvrRelatedSn *string `json:"SvrRelatedSn,omitempty" name:"SvrRelatedSn"`
	// 3：黑石主机，4：智能网卡

	SvrType *int64 `json:"SvrType,omitempty" name:"SvrType"`
}

func (r *ReinstallOsExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReinstallOsExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateVMWanIP struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 设备SN

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 设备外网IP

	SvrWanIp *string `json:"SvrWanIp,omitempty" name:"SvrWanIp"`
}

type DescribeRAIDListRequest struct {
	*tchttp.BaseRequest

	// "Op": "query_raid"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 模糊匹配

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制返回个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRAIDListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRAIDListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBackendComponentsExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果；0：成功；1：失败

		Result *int64 `json:"Result,omitempty" name:"Result"`
		// 返回详情

		Detail *string `json:"Detail,omitempty" name:"Detail"`
		// 后端组件；dcos / tcsplatform

		BackendComponent *string `json:"BackendComponent,omitempty" name:"BackendComponent"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBackendComponentsExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBackendComponentsExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerSpecialAllocateLanIPExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ServerSpecialAllocateLanIPInfo

		ServerSpecialAllocateLanIPInfo []*ServerSpecialAllocateLanIPInfo `json:"ServerSpecialAllocateLanIPInfo,omitempty" name:"ServerSpecialAllocateLanIPInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ServerSpecialAllocateLanIPExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerSpecialAllocateLanIPExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PerformServerTaskSet struct {
	// SvrCurrentStatus

	SvrCurrentStatus *string `json:"SvrCurrentStatus,omitempty" name:"SvrCurrentStatus"`
	// SvrIdcId

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// SvrSn

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// TaskId

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeOsDeployDebugInfoExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询os部署调试信息出参

		DebugInfo *string `json:"DebugInfo,omitempty" name:"DebugInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOsDeployDebugInfoExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOsDeployDebugInfoExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOSParam struct {
	// 系统名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type VMVirtualIPSet struct {
	// 设备固资编号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 设备SN

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 设备虚拟外网ip

	SvrVirtualWanIp *string `json:"SvrVirtualWanIp,omitempty" name:"SvrVirtualWanIp"`
	// 0成功 1失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 成功或者失败原因

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 物理服务器ip

	SvrAllocateIp *string `json:"SvrAllocateIp,omitempty" name:"SvrAllocateIp"`
	// 掩码

	SvrAllocateMask *string `json:"SvrAllocateMask,omitempty" name:"SvrAllocateMask"`
	// 网关

	SvrAllocateGateway *string `json:"SvrAllocateGateway,omitempty" name:"SvrAllocateGateway"`
	// 设备虚拟内网ip

	SvrVirtualLanIp *string `json:"SvrVirtualLanIp,omitempty" name:"SvrVirtualLanIp"`
}

type AllocateServerSpecialWanIPListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"special_allocate_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 分配服务器特殊外网IP

	AllocateServerSpecialWanIPSet []*AllocateServerSpecialWanIP `json:"AllocateServerSpecialWanIPSet,omitempty" name:"AllocateServerSpecialWanIPSet"`
}

func (r *AllocateServerSpecialWanIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateServerSpecialWanIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOutbandStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回串中的DataSet节点，规范协议请求

		DataSet *ItemsDataSet `json:"DataSet,omitempty" name:"DataSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOutbandStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOutbandStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateVMWanIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟机分配外网IP信息

		VMAllocWanIPSet []*VMAllocWanIP `json:"VMAllocWanIPSet,omitempty" name:"VMAllocWanIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateVMWanIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateVMWanIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustomScriptSetRequest struct {
	*tchttp.BaseRequest

	// 需要创建的自定义脚本集名

	CustomScriptSetName *string `json:"CustomScriptSetName,omitempty" name:"CustomScriptSetName"`
	// 需要加入的自定义脚本名列表

	CustomScriptName []*string `json:"CustomScriptName,omitempty" name:"CustomScriptName"`
	// 脚本集类型(0:初始化脚本,1:后置脚本)

	CustomScriptSetType *int64 `json:"CustomScriptSetType,omitempty" name:"CustomScriptSetType"`
}

func (r *CreateCustomScriptSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomScriptSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateVMLanIP struct {
	// 服务器资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器序列号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 服务器内网IP

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
}

type ConfigServerDefSetLan struct {
	// 内网IP

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 掩码

	Netmask *string `json:"Netmask,omitempty" name:"Netmask"`
	// 网关

	Gateway *string `json:"Gateway,omitempty" name:"Gateway"`
}

type Items struct {
	// 返回结果详细说明

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 返回结果值，0成功， 1失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 修改成功的策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
}

type RecycleServerWanIP struct {
	// 服务器资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器外网IP

	SvrWanIp *string `json:"SvrWanIp,omitempty" name:"SvrWanIp"`
}

type DescribeAllVlanIdsExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 所有网段vlanid

		AllVlanIds []*string `json:"AllVlanIds,omitempty" name:"AllVlanIds"`
		// 返回总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllVlanIdsExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllVlanIdsExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOSParam struct {
	// 序号

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type ModifySpecialIdcLineItems struct {
	// SpecialCondition

	SpecialCondition *SpecialCondition `json:"SpecialCondition,omitempty" name:"SpecialCondition"`
	// SpecialModify

	SpecialModify *SpecialModify `json:"SpecialModify,omitempty" name:"SpecialModify"`
}

type DescribeLabelRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"label"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 操作类型："Op":"query"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 返回数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回参数

	ResultItem *string `json:"ResultItem,omitempty" name:"ResultItem"`
}

func (r *DescribeLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySyslogKeywordParams struct {
	// 修改系统日志关键字参数

	ModifySyslogKeywordConds []*ModifySyslogKeywordCond `json:"ModifySyslogKeywordConds,omitempty" name:"ModifySyslogKeywordConds"`
}

type SuccessInfo struct {
	// 导入成功的固资号

	SvrAssetIds []*string `json:"SvrAssetIds,omitempty" name:"SvrAssetIds"`
	// 导入成功的机器数量

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type DeleteLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// DeleteLabel返回参数

		DataSet *Labels `json:"DataSet,omitempty" name:"DataSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomScriptsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自定义脚本详细列表

		CustomScripts []*CustomScripts `json:"CustomScripts,omitempty" name:"CustomScripts"`
		// 自定义脚本的总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomScriptsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomScriptsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryServerStatusRequest struct {
	*tchttp.BaseRequest

	// Zone

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// IdcName

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// TaskId

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *QueryServerStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryServerStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelocateServerStartRequest struct {
	*tchttp.BaseRequest

	// Scheme

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"relocate_start"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 发起搬迁信息

	RelocateStartServers []*RelocateStartServer `json:"RelocateStartServers,omitempty" name:"RelocateStartServers"`
}

func (r *RelocateServerStartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelocateServerStartRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomScriptSets struct {
	// 自定义脚本集名

	CustomScriptSetName *string `json:"CustomScriptSetName,omitempty" name:"CustomScriptSetName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 加入自定义脚本集的脚本

	CustomScriptName []*string `json:"CustomScriptName,omitempty" name:"CustomScriptName"`
	// 脚本集类型(0:初始化脚本,1:后置脚本)

	CustomScriptSetType *int64 `json:"CustomScriptSetType,omitempty" name:"CustomScriptSetType"`
}

type ModifyImageCondition struct {
	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

type AllocateVMLanIPListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server_vm"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"allocate_lan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 分配VM内网信息集合

	AllocateVMLanIPSet []*AllocateVMLanIP `json:"AllocateVMLanIPSet,omitempty" name:"AllocateVMLanIPSet"`
}

func (r *AllocateVMLanIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateVMLanIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLabelInfo struct {
	// 要修改的标签id

	Condition *LabelConditionInfo `json:"Condition,omitempty" name:"Condition"`
	// 修改后的标签内容

	Modify *ModifyLabelModifyInfo `json:"Modify,omitempty" name:"Modify"`
}

type SpecialModify struct {
	// SpeLine名称

	SpeLineName *string `json:"SpeLineName,omitempty" name:"SpeLineName"`
	// 类型

	SpeLineType *int64 `json:"SpeLineType,omitempty" name:"SpeLineType"`
	// 状态

	SpeLineStatus *int64 `json:"SpeLineStatus,omitempty" name:"SpeLineStatus"`
	// Func

	SpeLineFunc *int64 `json:"SpeLineFunc,omitempty" name:"SpeLineFunc"`
	// Speed

	SpeLineSpeed *string `json:"SpeLineSpeed,omitempty" name:"SpeLineSpeed"`
	// 本地设备

	LocalDevice *string `json:"LocalDevice,omitempty" name:"LocalDevice"`
	// 本地设备端口

	LocalDevicePort *string `json:"LocalDevicePort,omitempty" name:"LocalDevicePort"`
	// 本地设备ip

	LocalDeviceIp *string `json:"LocalDeviceIp,omitempty" name:"LocalDeviceIp"`
	// 移除本地设备ip

	RemoteDeviceIp *string `json:"RemoteDeviceIp,omitempty" name:"RemoteDeviceIp"`
	// 专线操作

	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`
	// 操作owner

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

type ServerAllocateLanIPInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type DeleteImageParams struct {
	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 自增ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type ReinstallOsSet struct {
	// SvrSn

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// TaskId

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// SvrIdcName

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// SvrIdcId

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
}

type DeleteAlarmStrategyMd5 struct {
	// /策略MD5，新增策略时给定

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
}

type ModifyServerPara struct {
	// '服务器所属业务'

	SvrBussiness *string `json:"SvrBussiness,omitempty" name:"SvrBussiness"`
	// 服务器SN

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 服务器当前状态(0未开电 1已开电 2运营中 3已关机 4已下架 5待部署 6搬迁中 20运行中(VM) 21已销毁(VM))

	SvrCurrentStatus *string `json:"SvrCurrentStatus,omitempty" name:"SvrCurrentStatus"`
	// 服务器厂商

	SvrProducer *string `json:"SvrProducer,omitempty" name:"SvrProducer"`
	// 设备固资编号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 机架名称

	SvrRackName *string `json:"SvrRackName,omitempty" name:"SvrRackName"`
	// 设备型号

	SvrDeviceType *string `json:"SvrDeviceType,omitempty" name:"SvrDeviceType"`
	// 服务器硬件描述

	SvrDeviceDescript *string `json:"SvrDeviceDescript,omitempty" name:"SvrDeviceDescript"`
	// 存放机位ID

	SvrPosId *string `json:"SvrPosId,omitempty" name:"SvrPosId"`
	// SvrGpu

	SvrGpu *string `json:"SvrGpu,omitempty" name:"SvrGpu"`
}

type AllocateServerSpecialWanIP struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// Vlan

	Vlan *string `json:"Vlan,omitempty" name:"Vlan"`
	// IP数

	IpNum *string `json:"IpNum,omitempty" name:"IpNum"`
}

type ModifyLabel struct {
	// 修改条件

	Condition *LabelCondition `json:"Condition,omitempty" name:"Condition"`
	// 修改具体字段内容

	Modify *LabelModify `json:"Modify,omitempty" name:"Modify"`
}

type ModifySyslogKeywordCond struct {
	// 新的关键字

	NewKeyword *string `json:"NewKeyword,omitempty" name:"NewKeyword"`
	// 厂商

	NetdevVendor *string `json:"NetdevVendor,omitempty" name:"NetdevVendor"`
	// 类型

	NetdevType *string `json:"NetdevType,omitempty" name:"NetdevType"`
	// 功能

	NetdevFun *string `json:"NetdevFun,omitempty" name:"NetdevFun"`
	// 型号

	NetdevModel *string `json:"NetdevModel,omitempty" name:"NetdevModel"`
	// 旧的关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 告警类型

	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`
}

type AllocateVMWanIPListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server_vm"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"allocate_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 回收VM外网IP信息

	AllocateVMWanIPSet []*AllocateVMWanIP `json:"AllocateVMWanIPSet,omitempty" name:"AllocateVMWanIPSet"`
}

func (r *AllocateVMWanIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateVMWanIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetDeviceRelatedDatas struct {
	// IndexId

	IndexId *int64 `json:"IndexId,omitempty" name:"IndexId"`
	// NetdevAssetId

	NetdevAssetId *string `json:"NetdevAssetId,omitempty" name:"NetdevAssetId"`
	// PortName

	PortName *string `json:"PortName,omitempty" name:"PortName"`
	// DeviceRackName

	DeviceRackName *string `json:"DeviceRackName,omitempty" name:"DeviceRackName"`
	// ConnectionNetdevAssetid

	ConnectionNetdevAssetid *string `json:"ConnectionNetdevAssetid,omitempty" name:"ConnectionNetdevAssetid"`
	// ConnectionNetdevPort

	ConnectionNetdevPort *string `json:"ConnectionNetdevPort,omitempty" name:"ConnectionNetdevPort"`
}

type FastReinstallOSData struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 外网参数

	WanDatas []*WanData `json:"WanDatas,omitempty" name:"WanDatas"`
	// BondingFlag

	BondingFlag *string `json:"BondingFlag,omitempty" name:"BondingFlag"`
	// Sn

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 操作系统

	Os *string `json:"Os,omitempty" name:"Os"`
	// 用户Shell路径

	UserShellPath *string `json:"UserShellPath,omitempty" name:"UserShellPath"`
	// 内网数据

	LanDatas []*LanData `json:"LanDatas,omitempty" name:"LanDatas"`
	// Raid

	Raid *string `json:"Raid,omitempty" name:"Raid"`
	// 设备类型

	SvrDeviceType *string `json:"SvrDeviceType,omitempty" name:"SvrDeviceType"`
	// IDC名称

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

type ExportModify struct {
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
}

type DeleteOutbandStrategyRequest struct {
	*tchttp.BaseRequest

	// 请求串类型说明。
	// 固定值： "Scheme":"outband_passwd_strategy"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方systemid	此ID为请求方向cmdb注册的权限ID，由cmdb系统管理员提供和管理。
	// 默认值："SystemId":"sid_test"

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方systemkey	和上面systemid类似，由cmdb系统管理员提供和管理，cmdb侧会根据id和key以及调用方ip三者共同确定请求方的合法性。
	// 默认值："SystemKey":"sid_test"

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 删除自定义带外密码策略详细内容

	Data *OutbandStrategyData `json:"Data,omitempty" name:"Data"`
}

func (r *DeleteOutbandStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOutbandStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFieldsEnumExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// FieldToEnum

		FieldToEnum *Field2enum `json:"FieldToEnum,omitempty" name:"FieldToEnum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFieldsEnumExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFieldsEnumExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddLabelDataSetServers struct {
	// 返回结果详细说明

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 返回结果值，0成功， 1失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 标签的键

	LabelKey *string `json:"LabelKey,omitempty" name:"LabelKey"`
}

type AllocateServerWanIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分配物理服务器外网IP出参

		SvrAllocWanIPSet []*ServerAllocateLanIPInfo `json:"SvrAllocWanIPSet,omitempty" name:"SvrAllocWanIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateServerWanIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateServerWanIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelocateStartServer struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// IDC名

	SvrRelocateIdcName *string `json:"SvrRelocateIdcName,omitempty" name:"SvrRelocateIdcName"`
	// 机架名

	SvrRelocateRackName *string `json:"SvrRelocateRackName,omitempty" name:"SvrRelocateRackName"`
	// 机架位置

	SvrRelocatePosId *string `json:"SvrRelocatePosId,omitempty" name:"SvrRelocatePosId"`
	// 逻辑区

	SvrRelocateLogicArea *string `json:"SvrRelocateLogicArea,omitempty" name:"SvrRelocateLogicArea"`
}

type RelocateServerQueryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前返回数

		ReturnNum *int64 `json:"ReturnNum,omitempty" name:"ReturnNum"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 搬迁服务器状态查询信息

		RelocateServerSet *string `json:"RelocateServerSet,omitempty" name:"RelocateServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RelocateServerQueryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelocateServerQueryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOutbandStrategyDateSet struct {
	// 增加自定义带外密码策略接口出参DataSet内层参数

	Strategies []*AddStrategiesRes `json:"Strategies,omitempty" name:"Strategies"`
}

type PortServer struct {
	// 示例：{"SvrAssetId":"TYSV190521X1"}

	Condition *PortModifyCondition `json:"Condition,omitempty" name:"Condition"`
	// 示例：{"SvrMonPort":"523"}

	Modify *PortModify `json:"Modify,omitempty" name:"Modify"`
}

type RAIDDictionaryInfo struct {
	// 序号

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeRelocateServerHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前返回数

		ReturnNum *int64 `json:"ReturnNum,omitempty" name:"ReturnNum"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 历史搬迁服务器查询结果

		HistoryRelocateServerSet *string `json:"HistoryRelocateServerSet,omitempty" name:"HistoryRelocateServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRelocateServerHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRelocateServerHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetImageFieldsEnumRequest struct {
	*tchttp.BaseRequest

	// "Op": "get_image_enum"

	Op *string `json:"Op,omitempty" name:"Op"`
}

func (r *GetImageFieldsEnumRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetImageFieldsEnumRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LanData struct {
	// 网卡

	Nic *string `json:"Nic,omitempty" name:"Nic"`
	// 子网掩码

	Netmask *string `json:"Netmask,omitempty" name:"Netmask"`
	// IP地址

	Ipaddress *string `json:"Ipaddress,omitempty" name:"Ipaddress"`
	// 网关

	Gateway *string `json:"Gateway,omitempty" name:"Gateway"`
}

type DescribeIgniterImagesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeIgniterImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIgniterImagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOSParam struct {
	// 序号

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// OS名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeRelocateInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 所有逻辑区域集合

		SvrLogicArea []*string `json:"SvrLogicArea,omitempty" name:"SvrLogicArea"`
		// idc name 对应的空余机架机位信息

		Svridcname2deviceracknameDevposid *string `json:"Svridcname2deviceracknameDevposid,omitempty" name:"Svridcname2deviceracknameDevposid"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRelocateInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRelocateInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelCondition struct {
	// 物理服务器固资编号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 物理服务器SN

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
}

type ConfigServerDefResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回串中的data节点。"data":{"task_id":xxx}

		Data []*ConfigServerDefData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConfigServerDefResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfigServerDefResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRAIDParam struct {
	// RAID名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type RelocateCancelServer struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// IDC名称

	SvrRelocateIdcName *string `json:"SvrRelocateIdcName,omitempty" name:"SvrRelocateIdcName"`
	// 机架名称

	SvrRelocateRackName *string `json:"SvrRelocateRackName,omitempty" name:"SvrRelocateRackName"`
	// 机架位置

	SvrRelocatePosId *string `json:"SvrRelocatePosId,omitempty" name:"SvrRelocatePosId"`
}

type DescribeAllVlanIdsExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 条数限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 请求动作

	Op *string `json:"Op,omitempty" name:"Op"`
	// 查询条件

	FieldFilters *string `json:"FieldFilters,omitempty" name:"FieldFilters"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAllVlanIdsExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllVlanIdsExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServerInfo struct {
	// 详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 服务器查询信息

	ResultInfo *ServerResultInfo `json:"ResultInfo,omitempty" name:"ResultInfo"`
}

type ModifyAlarmStrategyMd5 struct {
	// 策略MD5，新增策略时给定."Strategy":"md5"

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 生频率，值的格式如 num1->num2，表示num1分钟内发生num2次

	OccurrenceFrequency *string `json:"OccurrenceFrequency,omitempty" name:"OccurrenceFrequency"`
	// 抑制类型。-2->永久抑制，-1->恢复前永久抑制，0->不抑制，>0->抑制x次后重新发送告警

	InhibitType *int64 `json:"InhibitType,omitempty" name:"InhibitType"`
	// 选填。告警通知人

	NoticePeople *string `json:"NoticePeople,omitempty" name:"NoticePeople"`
	// 选填。告警级别

	AlarmLevel *string `json:"AlarmLevel,omitempty" name:"AlarmLevel"`
	// 选填。

	ToEcc *int64 `json:"ToEcc,omitempty" name:"ToEcc"`
	// 告警恢复方式，0->主动恢复，1->被动恢复

	RecoverType *int64 `json:"RecoverType,omitempty" name:"RecoverType"`
	// 仅在recove_type=1时有效，表示连续X次正常恢复告警（数值型），连续X次没有告警上报则恢复告警（字符型）

	RecoverTimes *int64 `json:"RecoverTimes,omitempty" name:"RecoverTimes"`
}

type NetdevicesSet struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// NetdeviceRelatedResultInfo

	NetdeviceRelatedResultInfo []*NetdeviceRelatedResultInfo `json:"NetdeviceRelatedResultInfo,omitempty" name:"NetdeviceRelatedResultInfo"`
}

type PortModifyCondition struct {
	// 资产Id

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type QueryOutbandStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回串中的DataSet节点，规范协议请求

		DataSet *StrategiesDataSet `json:"DataSet,omitempty" name:"DataSet"`
		// 本次查询返回的数据条数

		ReturnNum *int64 `json:"ReturnNum,omitempty" name:"ReturnNum"`
		// 满足本次查询条件的数据总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryOutbandStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryOutbandStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PhysvrsResultInfo struct {
	// RegionName

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// SvrAdminSwitchAssetid

	SvrAdminSwitchAssetid *string `json:"SvrAdminSwitchAssetid,omitempty" name:"SvrAdminSwitchAssetid"`
	// SvrAgentVersion

	SvrAgentVersion *string `json:"SvrAgentVersion,omitempty" name:"SvrAgentVersion"`
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// SvrBelongingProducts

	SvrBelongingProducts *string `json:"SvrBelongingProducts,omitempty" name:"SvrBelongingProducts"`
	// SvrBussiness

	SvrBussiness *string `json:"SvrBussiness,omitempty" name:"SvrBussiness"`
	// SvrBussinessId

	SvrBussinessId *string `json:"SvrBussinessId,omitempty" name:"SvrBussinessId"`
	// SvrCurrentStatus

	SvrCurrentStatus *string `json:"SvrCurrentStatus,omitempty" name:"SvrCurrentStatus"`
	// SvrDeviceDescript

	SvrDeviceDescript *string `json:"SvrDeviceDescript,omitempty" name:"SvrDeviceDescript"`
	// SvrDeviceHeight

	SvrDeviceHeight *string `json:"SvrDeviceHeight,omitempty" name:"SvrDeviceHeight"`
	// SvrDeviceName

	SvrDeviceName *string `json:"SvrDeviceName,omitempty" name:"SvrDeviceName"`
	// SvrDeviceType

	SvrDeviceType *string `json:"SvrDeviceType,omitempty" name:"SvrDeviceType"`
	// SvrFirstUseTime

	SvrFirstUseTime *string `json:"SvrFirstUseTime,omitempty" name:"SvrFirstUseTime"`
	// SvrId

	SvrId *string `json:"SvrId,omitempty" name:"SvrId"`
	// SvrIdcId

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// SvrIdcName

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// SvrInsertTime

	SvrInsertTime *string `json:"SvrInsertTime,omitempty" name:"SvrInsertTime"`
	// SvrLanIp

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
	// SvrLanSwitchAssetid

	SvrLanSwitchAssetid *string `json:"SvrLanSwitchAssetid,omitempty" name:"SvrLanSwitchAssetid"`
	// SvrLogicArea

	SvrLogicArea *string `json:"SvrLogicArea,omitempty" name:"SvrLogicArea"`
	// SvrOsName

	SvrOsName *string `json:"SvrOsName,omitempty" name:"SvrOsName"`
	// SvrPosId

	SvrPosId *string `json:"SvrPosId,omitempty" name:"SvrPosId"`
	// SvrProducer

	SvrProducer *string `json:"SvrProducer,omitempty" name:"SvrProducer"`
	// SvrRackName

	SvrRackName *string `json:"SvrRackName,omitempty" name:"SvrRackName"`
	// SvrRaidType

	SvrRaidType *string `json:"SvrRaidType,omitempty" name:"SvrRaidType"`
	// SvrSn

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// SvrType

	SvrType *string `json:"SvrType,omitempty" name:"SvrType"`
	// SvrVmIndex

	SvrVmIndex *string `json:"SvrVmIndex,omitempty" name:"SvrVmIndex"`
	// SvrWanEip

	SvrWanEip *string `json:"SvrWanEip,omitempty" name:"SvrWanEip"`
	// SvrWanIp

	SvrWanIp *string `json:"SvrWanIp,omitempty" name:"SvrWanIp"`
	// SvrWanSwitchAssetid

	SvrWanSwitchAssetid *string `json:"SvrWanSwitchAssetid,omitempty" name:"SvrWanSwitchAssetid"`
	// ZoneName

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// TaskId

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// SvrLanGateway

	SvrLanGateway *string `json:"SvrLanGateway,omitempty" name:"SvrLanGateway"`
	// SvrLanMask

	SvrLanMask *string `json:"SvrLanMask,omitempty" name:"SvrLanMask"`
	// SvrRelatedSn

	SvrRelatedSn *string `json:"SvrRelatedSn,omitempty" name:"SvrRelatedSn"`
	// SvrFaultStatus

	SvrFaultStatus []*int64 `json:"SvrFaultStatus,omitempty" name:"SvrFaultStatus"`
	// SvrVirtualLanIp

	SvrVirtualLanIp *string `json:"SvrVirtualLanIp,omitempty" name:"SvrVirtualLanIp"`
	// SvrArch

	SvrArch *string `json:"SvrArch,omitempty" name:"SvrArch"`
	// SvrVirtualWanIp

	SvrVirtualWanIp *string `json:"SvrVirtualWanIp,omitempty" name:"SvrVirtualWanIp"`
	// SvrLabels

	SvrLabels *string `json:"SvrLabels,omitempty" name:"SvrLabels"`
}

type DeleteVMListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server_vm"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"delete"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 删除VM信息

	DeleteVMSet []*VM `json:"DeleteVMSet,omitempty" name:"DeleteVMSet"`
}

func (r *DeleteVMListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVMListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRelocateData struct {
	// 服务器IDC

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
}

type DeleteRAIDParam struct {
	// 编号

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type ExportIdcLineInfo struct {
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
}

type IpResourceDatas struct {
	// indexid

	IndexId *int64 `json:"IndexId,omitempty" name:"IndexId"`
	// NetdevAssetId

	NetdevAssetId *string `json:"NetdevAssetId,omitempty" name:"NetdevAssetId"`
	// NetsegmentName

	NetsegmentName *string `json:"NetsegmentName,omitempty" name:"NetsegmentName"`
	// IdcName

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// NetsegmentMask

	NetsegmentMask *string `json:"NetsegmentMask,omitempty" name:"NetsegmentMask"`
	// NetsegmentGateway

	NetsegmentGateway *string `json:"NetsegmentGateway,omitempty" name:"NetsegmentGateway"`
	// NetsegmentIpPool

	NetsegmentIpPool *string `json:"NetsegmentIpPool,omitempty" name:"NetsegmentIpPool"`
	// NetsegmentType

	NetsegmentType *string `json:"NetsegmentType,omitempty" name:"NetsegmentType"`
	// idcid

	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
	// RelatedNetdevices

	RelatedNetdevices *string `json:"RelatedNetdevices,omitempty" name:"RelatedNetdevices"`
}

type NetdeviceRelatedResultInfo struct {
	// ConnectionNetdevAssetid

	ConnectionNetdevAssetid *string `json:"ConnectionNetdevAssetid,omitempty" name:"ConnectionNetdevAssetid"`
	// ConnectionNetdevPort

	ConnectionNetdevPort *string `json:"ConnectionNetdevPort,omitempty" name:"ConnectionNetdevPort"`
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// NetdevAssetId

	NetdevAssetId *string `json:"NetdevAssetId,omitempty" name:"NetdevAssetId"`
	// PortName

	PortName *string `json:"PortName,omitempty" name:"PortName"`
}

type VMAllocLanIP struct {
	// 结果详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 服务器资产Id

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器序列号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
}

type AddOutbandStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回串中的DataSet节点，规范协议请求

		DataSet *AddOutbandStrategyDateSet `json:"DataSet,omitempty" name:"DataSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddOutbandStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOutbandStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OutbandStrategiesCondition struct {
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 策略ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

type AggregationXflowParams struct {
	// 开始时间

	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`
	// 结束时间

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
	// 专线信息

	Lines []*Line `json:"Lines,omitempty" name:"Lines"`
}

type NetDeviceSet struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// IndexId

	IndexId *int64 `json:"IndexId,omitempty" name:"IndexId"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
}

type RenameServersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 重命名服务器接口出参

		RenameServerSet []*ServerAllocateLanIPInfo `json:"RenameServerSet,omitempty" name:"RenameServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenameServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelConditionInfo struct {
	// 要修改标签的id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type DescribeAlarmRecord struct {
	// {"Name":"ItemType","Value": 0},{"Name":"AlarmName","Value":""},{"Name":"Ip","Value":""},{"Name":"AssetId","Value":""},{"Name":"NetdeviceName","Value":""},{"Name":"NoticePeople","Value":""},{"Name":"AlarmLevel","Value":""},{"Name":"HasRecovery","Value":""}

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type IdcDevices struct {
	// SvrPosId

	SvrPosId *string `json:"SvrPosId,omitempty" name:"SvrPosId"`
	// SvrDeviceHeight

	SvrDeviceHeight *string `json:"SvrDeviceHeight,omitempty" name:"SvrDeviceHeight"`
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// PosInnerSwitchPortName

	PosInnerSwitchPortName *string `json:"PosInnerSwitchPortName,omitempty" name:"PosInnerSwitchPortName"`
	// PosOuterSwitchPortName

	PosOuterSwitchPortName *string `json:"PosOuterSwitchPortName,omitempty" name:"PosOuterSwitchPortName"`
	// PosAdminSwitchPortName

	PosAdminSwitchPortName *string `json:"PosAdminSwitchPortName,omitempty" name:"PosAdminSwitchPortName"`
	// DeviceHeight

	DeviceHeight *string `json:"DeviceHeight,omitempty" name:"DeviceHeight"`
	// PosLogicArea

	PosLogicArea *string `json:"PosLogicArea,omitempty" name:"PosLogicArea"`
}

type OutbandStrategyData struct {
	// 删除\修改自定义带外策略Strategies

	Strategies []*OutbandStrategies `json:"Strategies,omitempty" name:"Strategies"`
}

type StrategiesDataSet struct {
	// 查询带外密码规则返回参数集合

	Strategies []*Strategies `json:"Strategies,omitempty" name:"Strategies"`
}

type DeleteSyslogKeywordParams struct {
	// 删除系统关键字条件

	DeleteSyslogKeywordConds []*DeleteSyslogKeywordCond `json:"DeleteSyslogKeywordConds,omitempty" name:"DeleteSyslogKeywordConds"`
}

type ServerSpecialAllocateLanIPExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 请求动作

	Op *string `json:"Op,omitempty" name:"Op"`
	// 物理服务器批量分配相应虚拟内网段IP接口

	ServersSpecialAllocateLanIp []*ServersSpecialAllocateLanIp `json:"ServersSpecialAllocateLanIp,omitempty" name:"ServersSpecialAllocateLanIp"`
}

func (r *ServerSpecialAllocateLanIPExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerSpecialAllocateLanIPExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OSInfo struct {
	// 编号

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 操作系统名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DeleteImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// "ImageSet": [{"Detail": "ok","Result": 0,"ImageId": "ins-123456","ImageName": "test_os"}

		ImageSet *DeleteImageSet `json:"ImageSet,omitempty" name:"ImageSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenameServersRequest struct {
	*tchttp.BaseRequest

	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// "Scheme": "server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 操作类型：modify

	Op *string `json:"Op,omitempty" name:"Op"`
	// 重命名服务器集合

	RenameServerSet []*RenameServer `json:"RenameServerSet,omitempty" name:"RenameServerSet"`
}

func (r *RenameServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameServersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IdcExportLineInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// IdcExporResultInfo

	IdcExporResultInfo []*IdcExporResultInfo `json:"IdcExporResultInfo,omitempty" name:"IdcExporResultInfo"`
}

type IPSesource struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// IPResultInfo

	IPResultInfo *IPResultInfo `json:"IPResultInfo,omitempty" name:"IPResultInfo"`
}

type ProcessServerConditionModify struct {
	// 服务器进程修改条件

	Condition *ProcessCondition `json:"Condition,omitempty" name:"Condition"`
	// 服务器进程修改字段

	Modify *ProcessModify `json:"Modify,omitempty" name:"Modify"`
}

type CreateCustomScriptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCustomScriptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePhysvrsListExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalCount

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// PhysvrsListSet

		PhysvrsListSet []*PhysvrsListSet `json:"PhysvrsListSet,omitempty" name:"PhysvrsListSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePhysvrsListExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePhysvrsListExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VirtualLanOrWanIP struct {
	// 设备固资编号（实体服务器必须传）

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 虚拟机设备UUID

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 虚拟内网ip

	SvrVirtualLanIp *string `json:"SvrVirtualLanIp,omitempty" name:"SvrVirtualLanIp"`
	// 虚拟外网ip

	SvrVirtualWanIp *string `json:"SvrVirtualWanIp,omitempty" name:"SvrVirtualWanIp"`
}

type AddIdcInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// IndexId

	IndexId *int64 `json:"IndexId,omitempty" name:"IndexId"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
}

type PerformServerInfo struct {
	// 固资号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// idcid

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// Lanip

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
	// idcName

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
}

type RenameServerModify struct {
	// 设备名称

	SvrDeviceName *string `json:"SvrDeviceName,omitempty" name:"SvrDeviceName"`
}

type Wan struct {
	// OS部署

	Data []*WanInfo `json:"Data,omitempty" name:"Data"`
}

type AllocateServerVirtualIPRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 虚拟内网："Op":"allocate_vitrual_lan_ip"     虚拟外网： "Op":"allocate_vitrual_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 分配物理服务器虚拟内网外网信息集合

	AllocateServerVirtualIPSet []*VirtualLanOrWanIP `json:"AllocateServerVirtualIPSet,omitempty" name:"AllocateServerVirtualIPSet"`
}

func (r *AllocateServerVirtualIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateServerVirtualIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSysLogKeywordParams struct {
	// 查询系统日志关键字条件

	DescribeSysLogKeywordConds []*DescribeSysLogKeywordCond `json:"DescribeSysLogKeywordConds,omitempty" name:"DescribeSysLogKeywordConds"`
}

type Field2enum struct {
	// SvrBelongingProducts

	SvrBelongingProducts []*string `json:"SvrBelongingProducts,omitempty" name:"SvrBelongingProducts"`
	// SvrCurrentStatus

	SvrCurrentStatus []*string `json:"SvrCurrentStatus,omitempty" name:"SvrCurrentStatus"`
	// SvrDeviceType

	SvrDeviceType []*string `json:"SvrDeviceType,omitempty" name:"SvrDeviceType"`
	// SvrLogicArea

	SvrLogicArea []*string `json:"SvrLogicArea,omitempty" name:"SvrLogicArea"`
	// SvrRackName

	SvrRackName []*string `json:"SvrRackName,omitempty" name:"SvrRackName"`
	// CmdbStatusToModify

	CmdbStatusToModify []*string `json:"CmdbStatusToModify,omitempty" name:"CmdbStatusToModify"`
}

type AddServersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 增加服务器信息返回参数，示例：{"ServerSet":[{"Detail":"ok","IndexId":0,"Result":0}}

		ServerSet []*ServerSetInfo `json:"ServerSet,omitempty" name:"ServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomScriptRequest struct {
	*tchttp.BaseRequest

	// 自定义脚本名

	CustomScriptName *string `json:"CustomScriptName,omitempty" name:"CustomScriptName"`
	// 重命名的自定义脚本名

	CustomScriptNameNew *string `json:"CustomScriptNameNew,omitempty" name:"CustomScriptNameNew"`
	// 自定义脚本的内容，base64编码

	CustomScriptContext *string `json:"CustomScriptContext,omitempty" name:"CustomScriptContext"`
}

func (r *ModifyCustomScriptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomScriptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmStrategy struct {
	// 必填。0->服务器，1->网络设备，2->网络专线，3->网络出口，4->自定义告警项，5->网络质量

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 必填。CI项属性，比如"机房=深圳"、"负责人=ABC"，支持多个维度组合，维度之间用逗号隔开。如"机房=深圳,机型=DELL"。type=4时为空

	CiAttr *string `json:"CiAttr,omitempty" name:"CiAttr"`
	// //必填。命令空间，比如"tgw_custom"(TGW定制化)、"nas_group"(NAS集群)。type=4时有效

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 告警指标类型，0->数值型，1->字符型

	AlarmType *int64 `json:"AlarmType,omitempty" name:"AlarmType"`
	// 告警指标名

	AlarmName *string `json:"AlarmName,omitempty" name:"AlarmName"`
	// 字符型告警取值为-1。数值判断方法，-1->无效，0->大于，1->小于，2->大于等于，3->小于等于，4->and组合，5->or组合

	JudgeMethod *int64 `json:"JudgeMethod,omitempty" name:"JudgeMethod"`
	// 字符型告警取值为空。-1 -> 无效。若判断方法为0-3，则该字段为阈值(整型)；若为4-5，则填充规则如下：type1-val1，type2-val2，其中type取值同判断类型，val为具体值

	JudgeValue *string `json:"JudgeValue,omitempty" name:"JudgeValue"`
	// 发生频率，值的格式如 num1->num2，表示num1分钟内发生num2次。默认num1=1，num2=1

	OccurrenceFrequency *string `json:"OccurrenceFrequency,omitempty" name:"OccurrenceFrequency"`
	// 抑制类型。-2->永久抑制，-1->恢复前永久抑制，0->不抑制，>0->抑制x次后重新发送告警。默认-1

	InhibitType *int64 `json:"InhibitType,omitempty" name:"InhibitType"`
	// 告警通知人

	NoticePeople *string `json:"NoticePeople,omitempty" name:"NoticePeople"`
	// 告警级别

	AlarmLevel *string `json:"AlarmLevel,omitempty" name:"AlarmLevel"`
	// 选填。

	ToEcc *int64 `json:"ToEcc,omitempty" name:"ToEcc"`
	// 必填。告警恢复方式，0->主动恢复，1->被动恢复。

	RecoverType *int64 `json:"RecoverType,omitempty" name:"RecoverType"`
	// 仅在recovery_type=1时有效，表示连续X次正常恢复告警（数值型），连续X分钟没有告警上报则恢复告警（字符型）

	RecoverTimes *int64 `json:"RecoverTimes,omitempty" name:"RecoverTimes"`
	// 更新告警策略配置-MD5版使用该参数时，对应Type，CiAttr，Namespace，AlarmType，AlarmName，JudgeMethod，JudgeValue可以不填

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
}

type ModifyExportIdcLineInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// ExportLineId

	ExportLineId *string `json:"ExportLineId,omitempty" name:"ExportLineId"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
}

type ImportServersResult struct {
	// 导入失败的信息

	Failed *FailedInfo `json:"Failed,omitempty" name:"Failed"`
	// 导入成功的信息

	Success *SuccessInfo `json:"Success,omitempty" name:"Success"`
	// 导入的总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type ServerRecycleLanIPExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ServerRecycleLanIPInfo

		ServerRecycleLanIPInfo []*ServerRecycleLanIPInfo `json:"ServerRecycleLanIPInfo,omitempty" name:"ServerRecycleLanIPInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ServerRecycleLanIPExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerRecycleLanIPExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelModify struct {
	// 物理服务器标签

	SvrLabels *string `json:"SvrLabels,omitempty" name:"SvrLabels"`
}

type QueryServerStatus struct {
	// state

	State *string `json:"State,omitempty" name:"State"`
	// Status

	Status *string `json:"Status,omitempty" name:"Status"`
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Error

	Error *string `json:"Error,omitempty" name:"Error"`
	// TaskType

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
}

type AddVMListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新增虚拟服务器返回出参

		VMAddSet []*VMAddSet `json:"VMAddSet,omitempty" name:"VMAddSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddVMListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVMListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSyslogKeywordParams struct {
	// 增加系统日志关键字列表

	AddSyslogKeywords []*AddSyslogKeyword `json:"AddSyslogKeywords,omitempty" name:"AddSyslogKeywords"`
}

type RelocateServerQueryRequest struct {
	*tchttp.BaseRequest

	// Scheme

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// relocate_query

	Op *string `json:"Op,omitempty" name:"Op"`
	// svr_asset_id,svr_idc_name_relocate_before,svr_rack_name_relocate_before,svr_pos_id_relocate_before,svr_lan_ip_relocate_before,svr_idc_name_relocate_after,svr_rack_name_relocate_after,svr_pos_id_relocate_after,svr_logic_area_relocate_after,svr_lan_ip_relocate_after,svr_lan_gateway_relocate_after,svr_lan_mask_relocate_after

	ResultItem *string `json:"ResultItem,omitempty" name:"ResultItem"`
	// "Filters":[{"Name":"SvrIdcId","Values":["2001"]}

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *RelocateServerQueryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelocateServerQueryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSysLogParams struct {
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// StartTime

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制返回数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 需要返回数量

	NeedTotal *int64 `json:"NeedTotal,omitempty" name:"NeedTotal"`
	// "Filters":[
	//             {
	//                 "Name":"Ip",
	//                 "Values":[
	//                     "10.10.144.214",
	//                     "10.10.177.17"
	//                 ]
	//             }
	//         ]

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type RecycleVMWanIP struct {
	// 服务器资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器SN

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 服务器内网IP

	SvrWanIp *string `json:"SvrWanIp,omitempty" name:"SvrWanIp"`
}

type ServerRecycleWanIPInfo struct {
	// 返回结果详细说明

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 返回结果值，0成功， 1失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 固资号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 外网IP

	SvrWanIp *string `json:"SvrWanIp,omitempty" name:"SvrWanIp"`
}

type VM struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// Vlan

	Vlan *string `json:"Vlan,omitempty" name:"Vlan"`
	// 默认内网IP

	DefaultLanIP *string `json:"DefaultLanIP,omitempty" name:"DefaultLanIP"`
}

type DescribeAllOSListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询所有镜像列表出参集合

		OSDictionarySet []*OSDictionarySet `json:"OSDictionarySet,omitempty" name:"OSDictionarySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllOSListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllOSListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelocateServerStartResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务器搬迁结果信息

		StartRelocateServerSet []*StartRelocateServerInfo `json:"StartRelocateServerSet,omitempty" name:"StartRelocateServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RelocateServerStartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelocateServerStartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRelocateServerHistoryRequest struct {
	*tchttp.BaseRequest

	// Scheme

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"relocate_history_query"

	Op *string `json:"Op,omitempty" name:"Op"`
	// "ResultItem":"svr_asset_id"

	ResultItem *string `json:"ResultItem,omitempty" name:"ResultItem"`
	// "Filters":[{"Name":"SvrSn","Value":""},{"Name":"SvrAssetId","Value":"TYSV190521X1"}]}

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRelocateServerHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRelocateServerHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetdevPortState struct {
	// PortName

	PortName *string `json:"PortName,omitempty" name:"PortName"`
	// PortSnmpIndex

	PortSnmpIndex *string `json:"PortSnmpIndex,omitempty" name:"PortSnmpIndex"`
	// PortDesc

	PortDesc *string `json:"PortDesc,omitempty" name:"PortDesc"`
	// PortSpeed

	PortSpeed *string `json:"PortSpeed,omitempty" name:"PortSpeed"`
	// PortStateAdmin

	PortStateAdmin *int64 `json:"PortStateAdmin,omitempty" name:"PortStateAdmin"`
	// PortStateOper

	PortStateOper *int64 `json:"PortStateOper,omitempty" name:"PortStateOper"`
	// PortIp

	PortIp *string `json:"PortIp,omitempty" name:"PortIp"`
	// PortMac

	PortMac *string `json:"PortMac,omitempty" name:"PortMac"`
	// UpdateTime

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type TopnXflowParams struct {
	// TopN

	TopN *int64 `json:"TopN,omitempty" name:"TopN"`
	// Line信息

	Line *Line `json:"Line,omitempty" name:"Line"`
	// "Condition":["SrcIP","DstIP","SrcPort","DstPort","Protocol","Direction"]

	Condition []*string `json:"Condition,omitempty" name:"Condition"`
	// 开始时间

	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`
	// 结束时间

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type AddServer struct {
	// '设备固资编号'

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器SN(虚拟机此处为UUID,该值由云平提供和更新)

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// '设备型号',

	SvrDeviceType *string `json:"SvrDeviceType,omitempty" name:"SvrDeviceType"`
	// 'IDC名称',

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// '机架名称',

	SvrRackName *string `json:"SvrRackName,omitempty" name:"SvrRackName"`
	// '存放机位ID',

	SvrPosId *string `json:"SvrPosId,omitempty" name:"SvrPosId"`
	// '服务器厂商',

	SvrProducer *string `json:"SvrProducer,omitempty" name:"SvrProducer"`
	// 'IDC id',

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
}

type LabelsDataSet struct {
	// 查询标签返回参数集合

	Labels []*Labels `json:"Labels,omitempty" name:"Labels"`
}

type NetDeviceInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// IndexId

	IndexId *int64 `json:"IndexId,omitempty" name:"IndexId"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
}

type AllocateServerVirtualIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 物理服务器分配虚拟ip返回参数

		ServerAllocVirtualIPSet []*ServerAllocVirtualIPSet `json:"ServerAllocVirtualIPSet,omitempty" name:"ServerAllocVirtualIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateServerVirtualIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateServerVirtualIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBackendComponentsExRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryBackendComponentsExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBackendComponentsExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateVMVirtualIPRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server_vm"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 虚拟内网："Op":"allocate_vitrual_lan_ip"     虚拟外网： "Op":"allocate_vitrual_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 分配VM虚拟内网外网信息集合

	AllocateVMVirtualIPSet []*VirtualLanOrWanIP `json:"AllocateVMVirtualIPSet,omitempty" name:"AllocateVMVirtualIPSet"`
}

func (r *AllocateVMVirtualIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateVMVirtualIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomScriptSetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自定义脚本集的详细信息列表

		CustomScriptSets []*CustomScriptSets `json:"CustomScriptSets,omitempty" name:"CustomScriptSets"`
		// 自定义脚本集的总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomScriptSetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomScriptSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIdcInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// SpeLineId

	SpeLineId *string `json:"SpeLineId,omitempty" name:"SpeLineId"`
}

type DescribeNetdeviceRelatedServersExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 条数限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 网络设备资产号

	NetdevAssetId *string `json:"NetdevAssetId,omitempty" name:"NetdevAssetId"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNetdeviceRelatedServersExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetdeviceRelatedServersExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTaskExRequest struct {
	*tchttp.BaseRequest

	// 服务器SN

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *QueryTaskExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTaskExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WithdrawServerExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 请求动作

	Op *string `json:"Op,omitempty" name:"Op"`
	// 物理服务器信息

	WithdrawServers []*WithdrawServers `json:"WithdrawServers,omitempty" name:"WithdrawServers"`
}

func (r *WithdrawServerExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WithdrawServerExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpecialLineInfo struct {
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

type VMAddSet struct {
	// 返回结果详细说明

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 返回结果值，0成功， 1失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 固资号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 网关

	SvrIpGateway *string `json:"SvrIpGateway,omitempty" name:"SvrIpGateway"`
	// 掩码

	SvrIpMask *string `json:"SvrIpMask,omitempty" name:"SvrIpMask"`
	// 虚拟机固资号

	SvrVmAssetId *string `json:"SvrVmAssetId,omitempty" name:"SvrVmAssetId"`
	// 虚拟机ip

	SvrVmLanIp *string `json:"SvrVmLanIp,omitempty" name:"SvrVmLanIp"`
}

type DeleteCustomScriptSetsRequest struct {
	*tchttp.BaseRequest

	// 需要删除的自定义脚本集

	CustomScriptSetName []*string `json:"CustomScriptSetName,omitempty" name:"CustomScriptSetName"`
}

func (r *DeleteCustomScriptSetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomScriptSetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OriginalXflowParams struct {
	// 开始时间

	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`
	// 结束时间

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
	// 专线信息

	Lines []*Line `json:"Lines,omitempty" name:"Lines"`
}

type IdcInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// IdcQueryResultInfo

	IdcQueryResultInfo *IdcQueryResultInfo `json:"IdcQueryResultInfo,omitempty" name:"IdcQueryResultInfo"`
}

type ServerAllocateLanIPExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ServerAllocateLanIPInfo

		ServerAllocateLanIPInfo []*ServerAllocateLanIPInfo `json:"ServerAllocateLanIPInfo,omitempty" name:"ServerAllocateLanIPInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ServerAllocateLanIPExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerAllocateLanIPExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddExportIdcLineInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// IndexId

	IndexId *int64 `json:"IndexId,omitempty" name:"IndexId"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
}

type ModifyImageParams struct {
	// 要修改目标

	Condition *ModifyImageCondition `json:"Condition,omitempty" name:"Condition"`
	// 修改的内容

	Modify *ModifyImageModify `json:"Modify,omitempty" name:"Modify"`
}

type ServerResultInfo struct {
	// 服务器资产ID

	SvrId *string `json:"SvrId,omitempty" name:"SvrId"`
}

type StrategiesResultInfo struct {
	// 自增id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 是否包含小写字母（默认值1）

	HaveLowerLetter *string `json:"HaveLowerLetter,omitempty" name:"HaveLowerLetter"`
	// 是否包含小写字母（默认值1）

	HaveNumber *string `json:"HaveNumber,omitempty" name:"HaveNumber"`
	// 是否包含大写字母（默认值1）

	HaveUpperLetter *string `json:"HaveUpperLetter,omitempty" name:"HaveUpperLetter"`
	// 是否启用该策略 （默认禁用，即0）

	Enable *string `json:"Enable,omitempty" name:"Enable"`
	// 密码长度，默认15

	PasswdLength *string `json:"PasswdLength,omitempty" name:"PasswdLength"`
	// 特殊字符

	SpecialCharacters *string `json:"SpecialCharacters,omitempty" name:"SpecialCharacters"`
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type QueryTaskExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryTaskExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTaskExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleVMLanIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟机回收内网IP结果

		VmRecyLanIPSet []*VMRecyLanIP `json:"VmRecyLanIPSet,omitempty" name:"VmRecyLanIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecycleVMLanIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleVMLanIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageSet struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// ImageId

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
}

type ReinstallOsExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ReinstallOsSet

		ReinstallOsSet []*ReinstallOsSet `json:"ReinstallOsSet,omitempty" name:"ReinstallOsSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReinstallOsExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReinstallOsExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleServerSpecialWanIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 回收服务器特殊外网IP信息

		SvrSpecRecyWanIPSet []*SvrSpecRecyWanIP `json:"SvrSpecRecyWanIPSet,omitempty" name:"SvrSpecRecyWanIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecycleServerSpecialWanIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleServerSpecialWanIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageInfo struct {
	// 镜像ID

	ImageId *int64 `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像状态

	ImageStatus *int64 `json:"ImageStatus,omitempty" name:"ImageStatus"`
	// 镜像大小

	ImageSize *string `json:"ImageSize,omitempty" name:"ImageSize"`
	// 系统类型

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 系统架构

	SystemArch *string `json:"SystemArch,omitempty" name:"SystemArch"`
	// 系统平台

	SystemPlatform *string `json:"SystemPlatform,omitempty" name:"SystemPlatform"`
	// 系统版本

	SystemVersions *string `json:"SystemVersions,omitempty" name:"SystemVersions"`
	// 镜像描述

	ImageDescribe *string `json:"ImageDescribe,omitempty" name:"ImageDescribe"`
	// 镜像类型

	ImageType *int64 `json:"ImageType,omitempty" name:"ImageType"`
	// 镜像url

	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
	// 镜像使用服务器设备类型(0:通用服务器;1:黑石服务器)

	AvailableModel *int64 `json:"AvailableModel,omitempty" name:"AvailableModel"`
	// 镜像格式(iso/sqfs)

	ImageFormat *string `json:"ImageFormat,omitempty" name:"ImageFormat"`
	// ZoneName

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type Modify struct {
	// 服务器当前状态

	SvrCurrentStatus *string `json:"SvrCurrentStatus,omitempty" name:"SvrCurrentStatus"`
}

type SpecialIdcResultInfo struct {
	// BackupStatus

	BackupStatus *string `json:"BackupStatus,omitempty" name:"BackupStatus"`
	// Comment

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// ConstructUnit

	ConstructUnit *string `json:"ConstructUnit,omitempty" name:"ConstructUnit"`
	// CooperateUnit

	CooperateUnit *string `json:"CooperateUnit,omitempty" name:"CooperateUnit"`
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// IpSla

	IpSla *string `json:"IpSla,omitempty" name:"IpSla"`
	// LineBusiness

	LineBusiness *string `json:"LineBusiness,omitempty" name:"LineBusiness"`
	// LineDistance

	LineDistance *string `json:"LineDistance,omitempty" name:"LineDistance"`
	// LineEndTime

	LineEndTime *string `json:"LineEndTime,omitempty" name:"LineEndTime"`
	// LineOperator

	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`
	// LineStartTime

	LineStartTime *string `json:"LineStartTime,omitempty" name:"LineStartTime"`
	// LineOwner

	LineOwner *string `json:"LineOwner,omitempty" name:"LineOwner"`
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
	// LocalSecondDevice

	LocalSecondDevice *string `json:"LocalSecondDevice,omitempty" name:"LocalSecondDevice"`
	// LocalSecondDevicePort

	LocalSecondDevicePort *string `json:"LocalSecondDevicePort,omitempty" name:"LocalSecondDevicePort"`
	// RemoteDeviceIp

	RemoteDeviceIp *string `json:"RemoteDeviceIp,omitempty" name:"RemoteDeviceIp"`
	// RemoteOperatorOwner

	RemoteOperatorOwner *string `json:"RemoteOperatorOwner,omitempty" name:"RemoteOperatorOwner"`
	// RemoteOperatorOwnerPhone

	RemoteOperatorOwnerPhone *string `json:"RemoteOperatorOwnerPhone,omitempty" name:"RemoteOperatorOwnerPhone"`
	// RemoteOperatorOwnerWarningPhone

	RemoteOperatorOwnerWarningPhone *string `json:"RemoteOperatorOwnerWarningPhone,omitempty" name:"RemoteOperatorOwnerWarningPhone"`
	// SlaNum

	SlaNum *string `json:"SlaNum,omitempty" name:"SlaNum"`
	// SpeLineFunc

	SpeLineFunc *string `json:"SpeLineFunc,omitempty" name:"SpeLineFunc"`
	// SpeLineId

	SpeLineId *string `json:"SpeLineId,omitempty" name:"SpeLineId"`
	// SpeLineName

	SpeLineName *string `json:"SpeLineName,omitempty" name:"SpeLineName"`
	// SpeLineSpeed

	SpeLineSpeed *string `json:"SpeLineSpeed,omitempty" name:"SpeLineSpeed"`
	// SpeLineStatus

	SpeLineStatus *string `json:"SpeLineStatus,omitempty" name:"SpeLineStatus"`
	// SpeLineType

	SpeLineType *string `json:"SpeLineType,omitempty" name:"SpeLineType"`
	// SwitchAbility

	SwitchAbility *string `json:"SwitchAbility,omitempty" name:"SwitchAbility"`
	// VlanInfo

	VlanInfo *string `json:"VlanInfo,omitempty" name:"VlanInfo"`
}

type WithdrawServers struct {
	// 服务器信息

	Condition *Condition `json:"Condition,omitempty" name:"Condition"`
	// 服务器当前状态

	Modify *Modify `json:"Modify,omitempty" name:"Modify"`
}

type AllocateVMVirtualIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分配VM虚拟内网IP出参

		VMAllocVirtualIPSet []*VMVirtualIPSet `json:"VMAllocVirtualIPSet,omitempty" name:"VMAllocVirtualIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateVMVirtualIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateVMVirtualIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// "ImageSet": {"task_id": "xxx","Result": 0}

		ImageSet []*string `json:"ImageSet,omitempty" name:"ImageSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportIdcLineItems struct {
	// ExportCondition

	ExportCondition *ExportCondition `json:"ExportCondition,omitempty" name:"ExportCondition"`
	// ExportModify

	ExportModify *ExportModify `json:"ExportModify,omitempty" name:"ExportModify"`
}

type VMRecyWanIP struct {
	// 结果详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 服务器内网IP

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
	// 服务器编号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
}

type CreateCustomScriptRequest struct {
	*tchttp.BaseRequest

	// 自定义脚本的名称

	CustomScriptName *string `json:"CustomScriptName,omitempty" name:"CustomScriptName"`
	// 自定义脚本的内容，base64编码

	CustomScriptContext *string `json:"CustomScriptContext,omitempty" name:"CustomScriptContext"`
}

func (r *CreateCustomScriptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomScriptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOutbandInfoExRequest struct {
	*tchttp.BaseRequest

	// 可用区域ID

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// idcid

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// 请求动作

	Command *string `json:"Command,omitempty" name:"Command"`
	// 用户

	User *string `json:"User,omitempty" name:"User"`
	// 数据

	Data []*string `json:"Data,omitempty" name:"Data"`
}

func (r *DescribeOutbandInfoExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOutbandInfoExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerPortInfo struct {
	// 服务器资产编号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器ID

	SvrId *string `json:"SvrId,omitempty" name:"SvrId"`
	// 端口号

	SvrMonPort []*int64 `json:"SvrMonPort,omitempty" name:"SvrMonPort"`
	// 服务器SN

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 服务器端口数

	SvrMonPortLen *int64 `json:"SvrMonPortLen,omitempty" name:"SvrMonPortLen"`
}

type QueryServerStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// QueryServerStatusList

		QueryServerStatusList *string `json:"QueryServerStatusList,omitempty" name:"QueryServerStatusList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryServerStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryServerStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServersSpecialAllocateLanIp struct {
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// SvrIpWanted

	SvrIpWanted *string `json:"SvrIpWanted,omitempty" name:"SvrIpWanted"`
	// Vlan

	Vlan *string `json:"Vlan,omitempty" name:"Vlan"`
}

type AddLabelDataSet struct {
	// AddLabelDataSetServers

	Servers []*AddLabelDataSetServers `json:"Servers,omitempty" name:"Servers"`
}

type DescribeOSListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// {                "Id":270,                "Name":"Linux",                "CreateTime":"2020-05-08 16:30:39"            }

		OsDictionarySet []*OsParam `json:"OsDictionarySet,omitempty" name:"OsDictionarySet"`
		// 总量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOSListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOSListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IPResultInfo struct {
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// IdcId

	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
	// IdcName

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// NetsegmentGateway

	NetsegmentGateway *string `json:"NetsegmentGateway,omitempty" name:"NetsegmentGateway"`
	// NetsegmentMask

	NetsegmentMask *string `json:"NetsegmentMask,omitempty" name:"NetsegmentMask"`
	// NetsegmentName

	NetsegmentName *string `json:"NetsegmentName,omitempty" name:"NetsegmentName"`
	// NetsegmentType

	NetsegmentType *string `json:"NetsegmentType,omitempty" name:"NetsegmentType"`
	// RelatedNetdevices

	RelatedNetdevices *string `json:"RelatedNetdevices,omitempty" name:"RelatedNetdevices"`
	// RegionName

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// VlanId

	VlanId *string `json:"VlanId,omitempty" name:"VlanId"`
	// ZoneName

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type Netdevice struct {
	// 网络设备名称

	NetdevName *string `json:"NetdevName,omitempty" name:"NetdevName"`
	// 网络设备类型

	NetdevFunc *string `json:"NetdevFunc,omitempty" name:"NetdevFunc"`
	// 网络设备资产号

	NetdevAssetId *string `json:"NetdevAssetId,omitempty" name:"NetdevAssetId"`
}

type RecycleVMWanIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟机回收IP信息

		VmRecyWanIPSet []*VMRecyWanIP `json:"VmRecyWanIPSet,omitempty" name:"VmRecyWanIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecycleVMWanIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleVMWanIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleServerSpecialWanIPListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"special_recycle_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 回收服务器特殊外网IP

	RecycleServerSpecialWanIPSet []*RecycleServerSpecialWanIP `json:"RecycleServerSpecialWanIPSet,omitempty" name:"RecycleServerSpecialWanIPSet"`
}

func (r *RecycleServerSpecialWanIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleServerSpecialWanIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOSParam struct {
	// OS名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type IdcModify struct {
	// SvrBussiness

	SvrBussiness *string `json:"SvrBussiness,omitempty" name:"SvrBussiness"`
}

type ModifyRAIDParam struct {
	// 编号

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type ProcessCondition struct {
	// 资产id

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type ServerProcessResultInfo struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器ID

	SvrId *string `json:"SvrId,omitempty" name:"SvrId"`
	// 监控进程

	SvrMonProc *string `json:"SvrMonProc,omitempty" name:"SvrMonProc"`
	// 服务器编号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 进程数

	SvrMonProcLen *int64 `json:"SvrMonProcLen,omitempty" name:"SvrMonProcLen"`
}

type DescribeNetdeviceRelatedServersExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// RelatedServersSet

		RelatedServersSet []*Netdevice `json:"RelatedServersSet,omitempty" name:"RelatedServersSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetdeviceRelatedServersExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetdeviceRelatedServersExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerSpecialRecycleLanIPExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ServerSpecialRecycleLanIPInfo

		ServerSpecialRecycleLanIPInfo []*ServerSpecialAllocateLanIPInfo `json:"ServerSpecialRecycleLanIPInfo,omitempty" name:"ServerSpecialRecycleLanIPInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ServerSpecialRecycleLanIPExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerSpecialRecycleLanIPExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回串中的DataSet节点，规范协议请求

		DataSet *LabelsDataSet `json:"DataSet,omitempty" name:"DataSet"`
		// 本次查询返回的数据条数

		ReturnNum *int64 `json:"ReturnNum,omitempty" name:"ReturnNum"`
		// 满足本次查询条件的数据总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetDeviceResultInfo struct {
	// NetdevAdminIp

	NetdevAdminIp *string `json:"NetdevAdminIp,omitempty" name:"NetdevAdminIp"`
	// NetdevAssetId

	NetdevAssetId *string `json:"NetdevAssetId,omitempty" name:"NetdevAssetId"`
	// NetdevCurrentStatus

	NetdevCurrentStatus *string `json:"NetdevCurrentStatus,omitempty" name:"NetdevCurrentStatus"`
	// NetdevEngine

	NetdevEngine *string `json:"NetdevEngine,omitempty" name:"NetdevEngine"`
	// NetdevFunc

	NetdevFunc *string `json:"NetdevFunc,omitempty" name:"NetdevFunc"`
	// NetdevId

	NetdevId *string `json:"NetdevId,omitempty" name:"NetdevId"`
	// NetdevIdc

	NetdevIdc *string `json:"NetdevIdc,omitempty" name:"NetdevIdc"`
	// NetdevIdcId

	NetdevIdcId *string `json:"NetdevIdcId,omitempty" name:"NetdevIdcId"`
	// NetdevModel

	NetdevModel *string `json:"NetdevModel,omitempty" name:"NetdevModel"`
	// NetdevName

	NetdevName *string `json:"NetdevName,omitempty" name:"NetdevName"`
	// NetdevOs

	NetdevOs *string `json:"NetdevOs,omitempty" name:"NetdevOs"`
	// NetdevOtherIp

	NetdevOtherIp *string `json:"NetdevOtherIp,omitempty" name:"NetdevOtherIp"`
	// NetdevPro

	NetdevPro *string `json:"NetdevPro,omitempty" name:"NetdevPro"`
	// NetdevRackName

	NetdevRackName *string `json:"NetdevRackName,omitempty" name:"NetdevRackName"`
	// NetdevSn

	NetdevSn *string `json:"NetdevSn,omitempty" name:"NetdevSn"`
	// NetdevType

	NetdevType *string `json:"NetdevType,omitempty" name:"NetdevType"`
	// RegionName

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// ZoneName

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type ServerStatus struct {
	// SN

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// mac地址

	Mac *string `json:"Mac,omitempty" name:"Mac"`
	// 机器状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DescribeImageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本次查询返回的数据条数，对应下面data节点的数组个数

		ReturnNum *int64 `json:"ReturnNum,omitempty" name:"ReturnNum"`
		// 满足本次查询条件的数据总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回结果详细说明

		Detail *string `json:"Detail,omitempty" name:"Detail"`
		// 返回结果值，0成功， 1失败

		Result *int64 `json:"Result,omitempty" name:"Result"`
		// 返回串中的ImageSet节点，规范协议请求，方便后续节点增加。

		ImageSet []*ImageInfo `json:"ImageSet,omitempty" name:"ImageSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VMAllocWanIP struct {
	// 结果详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 服务器资产Id

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器序列号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
}

type Filter struct {
	// 需要过滤的字段。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段的过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type IdcQueryResultInfo struct {
	// DeviceHeight

	DeviceHeight *string `json:"DeviceHeight,omitempty" name:"DeviceHeight"`
	// IdcName

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
}

type RelocateServerFinishResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 搬迁完成接口出参

		FinishRelocateServerSet []*ServerAllocateLanIPInfo `json:"FinishRelocateServerSet,omitempty" name:"FinishRelocateServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RelocateServerFinishResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelocateServerFinishResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddStrategiesRes struct {
	// 返回结果详细说明

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 返回结果值，0成功， 1失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 新增的策略索引id

	IndexId *int64 `json:"IndexId,omitempty" name:"IndexId"`
}

type ConditionModify struct {
	// 修改条件

	Condition *ServerProcessModifyCondition `json:"Condition,omitempty" name:"Condition"`
	// 修改内容，字符串

	Modify *ServerProcessModify `json:"Modify,omitempty" name:"Modify"`
}

type IgniterImages struct {
	// 通用服务器镜像集合

	Normal []*string `json:"Normal,omitempty" name:"Normal"`
	// 黑石服务器镜像集合

	Soc []*string `json:"Soc,omitempty" name:"Soc"`
}

type ImageFilters struct {
	// 过滤键名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type Netdevices struct {
	// CmdbCondition

	CmdbCondition *CmdbCondition `json:"CmdbCondition,omitempty" name:"CmdbCondition"`
	// CmdbModify

	CmdbModify *CmdbModify `json:"CmdbModify,omitempty" name:"CmdbModify"`
}

type DescribeRelocateInfoRequest struct {
	*tchttp.BaseRequest

	// 查询搬迁信息

	DescribeRelocateData *DescribeRelocateData `json:"DescribeRelocateData,omitempty" name:"DescribeRelocateData"`
}

func (r *DescribeRelocateInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRelocateInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleVMWanIPListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server_vm"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"recycle_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 回收VM外网IP信息

	RecycleVMWanIPSet []*RecycleVMWanIP `json:"RecycleVMWanIPSet,omitempty" name:"RecycleVMWanIPSet"`
}

func (r *RecycleVMWanIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleVMWanIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServerHarddisksRequest struct {
	*tchttp.BaseRequest

	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeServerHarddisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServerHarddisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartRelocateServerInfo struct {
	// 结果详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 结果错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 服务器资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 网关

	SvrRelocateLanGateway *string `json:"SvrRelocateLanGateway,omitempty" name:"SvrRelocateLanGateway"`
	// 内网ip

	SvrRelocateLanIp *string `json:"SvrRelocateLanIp,omitempty" name:"SvrRelocateLanIp"`
	// 掩码

	SvrRelocateLanMask *string `json:"SvrRelocateLanMask,omitempty" name:"SvrRelocateLanMask"`
}

type AllocateVMLanIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟机分配内网IP信息

		VMAllocMLanIPSet []*VMAllocLanIP `json:"VMAllocMLanIPSet,omitempty" name:"VMAllocMLanIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateVMLanIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateVMLanIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFieldsEnumExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 显示条目数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统键值

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 执行动作

	Op *string `json:"Op,omitempty" name:"Op"`
	// 过滤条件

	FieldFilters *string `json:"FieldFilters,omitempty" name:"FieldFilters"`
	// 用于条件过滤查询。例如过滤ID、名称、状态等

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
}

func (r *DescribeFieldsEnumExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFieldsEnumExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeploySubtaskStepsExRequest struct {
	*tchttp.BaseRequest

	// SN号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// idcid

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// TaskId

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeDeploySubtaskStepsExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeploySubtaskStepsExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRAIDRequest struct {
	*tchttp.BaseRequest

	// "Op": "change_raid"

	Op *string `json:"Op,omitempty" name:"Op"`
	// "ModifyRAIDParams": {"Id":3,"Name":"Linux"}

	ModifyRAIDParams *ModifyRAIDParam `json:"ModifyRAIDParams,omitempty" name:"ModifyRAIDParams"`
}

func (r *ModifyRAIDRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRAIDRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOutbandStrategyDataSet struct {
	// 删除自定义带外密码策略接返回参数集合

	Strategies []*Items `json:"Strategies,omitempty" name:"Strategies"`
}

type AddRAIDResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新增RAID出参信息

		DataSet *string `json:"DataSet,omitempty" name:"DataSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddRAIDResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddRAIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerSetInfo struct {
	// 返回结果详细说明

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 导入服务器的顺序ID

	IndexId *int64 `json:"IndexId,omitempty" name:"IndexId"`
	// 返回结果值，0成功， 1失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
}

type RecycleServerWanIPListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统Id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"recycle_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 回收服务器外网IP

	RecycleServerWanIPSet []*RecycleServerWanIP `json:"RecycleServerWanIPSet,omitempty" name:"RecycleServerWanIPSet"`
}

func (r *RecycleServerWanIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleServerWanIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportCondition struct {
	// ExportLineId

	ExportLineId *string `json:"ExportLineId,omitempty" name:"ExportLineId"`
}

type ItemsDataSet struct {
	// 修改自定义带外密码策略接返回参数集合

	Items []*Items `json:"Items,omitempty" name:"Items"`
}

type Labels struct {
	// 返回结果详细说明

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 返回结果值，0成功， 1失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 标签的详细信息

	ResultInfo *LabelsResultInfo `json:"ResultInfo,omitempty" name:"ResultInfo"`
}

type PhysvrsListSet struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// PhysvrsResultInfo

	PhysvrsResultInfo *PhysvrsResultInfo `json:"PhysvrsResultInfo,omitempty" name:"PhysvrsResultInfo"`
}

type SpecialIdcLineInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// SpecialIdcResultInfo

	SpecialIdcResultInfo *SpecialIdcResultInfo `json:"SpecialIdcResultInfo,omitempty" name:"SpecialIdcResultInfo"`
}

type ConfigServerDefRequest struct {
	*tchttp.BaseRequest

	// 请求动作。"Command": "config_server_def"

	Command *string `json:"Command,omitempty" name:"Command"`
	// 此ID为请求方向cmdb注册的权限ID，由cmdb系统管理员提供和管理

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 和上面systemid类似，由cmdb系统管理员提供和管理，cmdb侧会根据id和key以及调用方ip三者共同确定请求方的合法性

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 本次初始化涉及的物理机信息

	ConfigServerDefSet []*ConfigServerDefSet `json:"ConfigServerDefSet,omitempty" name:"ConfigServerDefSet"`
	// Scheme: "server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
}

func (r *ConfigServerDefRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfigServerDefRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IpsDetailSet struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// IPDetailResultInfo

	IPDetailResultInfo *IPDetailResultInfo `json:"IPDetailResultInfo,omitempty" name:"IPDetailResultInfo"`
}

type PerformServerTaskExRequest struct {
	*tchttp.BaseRequest

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// SvrIdcId

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 用户

	User *string `json:"User,omitempty" name:"User"`
	// 操作类型

	OperateType *int64 `json:"OperateType,omitempty" name:"OperateType"`
	// 请求动作

	Command *string `json:"Command,omitempty" name:"Command"`
	// 服务方信息

	Servers []*PerformServerInfo `json:"Servers,omitempty" name:"Servers"`
}

func (r *PerformServerTaskExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PerformServerTaskExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerRecycleLanIPExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 请求动作

	Op *string `json:"Op,omitempty" name:"Op"`
	// 用户

	User *string `json:"User,omitempty" name:"User"`
	// 物理服务器回收内网IP接口

	ServersRecycleLanIp []*ServersRecycleLanIp `json:"ServersRecycleLanIp,omitempty" name:"ServersRecycleLanIp"`
}

func (r *ServerRecycleLanIPExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerRecycleLanIPExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetdevicesInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *string `json:"Result,omitempty" name:"Result"`
	// NetDeviceResultInfo

	NetDeviceResultInfo *NetDeviceResultInfo `json:"NetDeviceResultInfo,omitempty" name:"NetDeviceResultInfo"`
}

type OsDeployDebugInfo struct {
	// IdcId

	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
	// 物理服务器Sn

	Sns []*string `json:"Sns,omitempty" name:"Sns"`
}

type SvrCurrentStatus struct {
	// 总量

	Total *string `json:"Total,omitempty" name:"Total"`
	// 运行中

	Running *string `json:"Running,omitempty" name:"Running"`
	// 部署中

	OnInstall *string `json:"OnInstall,omitempty" name:"OnInstall"`
	// 待部署

	PreInstall *string `json:"PreInstall,omitempty" name:"PreInstall"`
	// 部署失败

	InstallFailed *string `json:"InstallFailed,omitempty" name:"InstallFailed"`
}

type GetImageFieldsEnumResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ImageFieldsEnumSet

		ImageFieldsEnumSet *ImageFieldsEnumSet `json:"ImageFieldsEnumSet,omitempty" name:"ImageFieldsEnumSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetImageFieldsEnumResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetImageFieldsEnumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WithdrawServerExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// WithdrawServerSet

		WithdrawServerSet []*WithdrawServerSet `json:"WithdrawServerSet,omitempty" name:"WithdrawServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WithdrawServerExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WithdrawServerExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportCmdbServersRequest struct {
	*tchttp.BaseRequest

	// 前端解析导入文件后生成的列表

	ImportServerList []*ImportServerList `json:"ImportServerList,omitempty" name:"ImportServerList"`
}

func (r *ImportCmdbServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportCmdbServersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalNetdeviceSyslogParams struct {
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// NeedTotal

	NeedTotal *int64 `json:"NeedTotal,omitempty" name:"NeedTotal"`
	// Filters

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// StartTime

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// EndTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type Line struct {
	// 专线ID，必选

	LineId *int64 `json:"LineId,omitempty" name:"LineId"`
	// 源服务器IP，可选

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 目标服务器IP，可选

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 源端口，可选

	SrcPort *int64 `json:"SrcPort,omitempty" name:"SrcPort"`
	// 目标端口，可选

	DstPort *int64 `json:"DstPort,omitempty" name:"DstPort"`
	// 协议，可选，1->ICMP，6->TCP，17->UDP

	Protocol *int64 `json:"Protocol,omitempty" name:"Protocol"`
	// 方向，可选，1->入流量，2->出流量

	Direction *int64 `json:"Direction,omitempty" name:"Direction"`
}

type AddLabelRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"label"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 操作类型："Op":"add"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 新增标签信息

	AddLabelSet []*AddLabelInfo `json:"AddLabelSet,omitempty" name:"AddLabelSet"`
}

func (r *AddLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前返回数

		ReturnNum *int64 `json:"ReturnNum,omitempty" name:"ReturnNum"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 查询服务器结果信息

		ServerSet []*DescribeServerInfo `json:"ServerSet,omitempty" name:"ServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIgniterImagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// igniter镜像接口出参

		Images *IgniterImages `json:"Images,omitempty" name:"Images"`
		// 返回结果值，0成功

		Result *int64 `json:"Result,omitempty" name:"Result"`
		// 返回总数

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIgniterImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIgniterImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServer struct {
	// 物理服务器固资号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type PeriodAggregationXflowParams struct {
	// 开始时间

	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`
	// 结束时间

	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
	// 专线信息

	Lines []*Line `json:"Lines,omitempty" name:"Lines"`
}

type Syslogkeywordset struct {
	// 告警类别

	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`
	// 说明

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 设备类别

	NetdevFunc *string `json:"NetdevFunc,omitempty" name:"NetdevFunc"`
	// 设备型号

	NetdevModel *string `json:"NetdevModel,omitempty" name:"NetdevModel"`
	// 设备类型

	NetdevType *string `json:"NetdevType,omitempty" name:"NetdevType"`
	// 设备厂商

	NetdevVendor *string `json:"NetdevVendor,omitempty" name:"NetdevVendor"`
}

type RelocateServerFinishRequest struct {
	*tchttp.BaseRequest

	// Scheme

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"relocate_finish"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 搬迁完成信息

	RelocateFinishServers []*RelocateFinishServer `json:"RelocateFinishServers,omitempty" name:"RelocateFinishServers"`
}

func (r *RelocateServerFinishRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelocateServerFinishRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleServerWanIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 回收服务器外网IP出参

		SvrRecyWanIPSet []*ServerRecycleWanIPInfo `json:"SvrRecyWanIPSet,omitempty" name:"SvrRecyWanIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecycleServerWanIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleServerWanIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOutbandStrategyRequest struct {
	*tchttp.BaseRequest

	// 请求串类型说明。
	// 固定值： "Scheme":"outband_passwd_strategy"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方systemid	此ID为请求方向cmdb注册的权限ID，由cmdb系统管理员提供和管理。
	// 默认值："SystemId":"sid_test"

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方systemkey	和上面systemid类似，由cmdb系统管理员提供和管理，cmdb侧会根据id和key以及调用方ip三者共同确定请求方的合法性。
	// 默认值："SystemKey":"sid_test"

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 本次请求动作的起始索引地址	不填该字段，系统默认值为0，适合前台翻页查询等场景

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 本次请求动作的数据返回值数量。

	Limit *string `json:"Limit,omitempty" name:"Limit"`
	// 查询条件。
	// 默认："result_item":"id, strategy_name, have_number,have_lower_letter,have_upper_letter,special_characters,passwd_length,enable"

	FieldFilters *string `json:"FieldFilters,omitempty" name:"FieldFilters"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryOutbandStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryOutbandStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WanInfo struct {
	// Nic

	Nic *string `json:"Nic,omitempty" name:"Nic"`
	// Ipaddress

	Ipaddress *string `json:"Ipaddress,omitempty" name:"Ipaddress"`
}

type ServersRecycleLanIp struct {
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// SvrLanIp

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
}

type CreateCustomScriptSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCustomScriptSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomScriptSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddVMListRequest struct {
	*tchttp.BaseRequest

	// server_vm

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// "SystemId":"sid_test"

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// "SystemKey":"sid_test"

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// add

	Op *string `json:"Op,omitempty" name:"Op"`
	// 增加VM信息

	AddVMSet []*VM `json:"AddVMSet,omitempty" name:"AddVMSet"`
}

func (r *AddVMListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVMListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyServersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 更新服务器信息出参

		ServerSet []*ServerAllocateLanIPInfo `json:"ServerSet,omitempty" name:"ServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerProcessModify struct {
	// 服务器进程名

	SvrMonProc *string `json:"SvrMonProc,omitempty" name:"SvrMonProc"`
}

type DeleteCustomScriptSetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCustomScriptSetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomScriptSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除服务器出参

		ServerSet []*ServerAllocateLanIPInfo `json:"ServerSet,omitempty" name:"ServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRAIDListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// RAID查询结果

		RaidDictionarySet []*RAIDDictionaryInfo `json:"RaidDictionarySet,omitempty" name:"RaidDictionarySet"`
		// 返回总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRAIDListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRAIDListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Lan struct {
	// OS部署

	Data []*Laninfo `json:"Data,omitempty" name:"Data"`
}

type DeleteCustomScriptsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCustomScriptsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomScriptsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomScriptSetsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li> CustomScriptSetName - String List - 是否必填：否 - （过滤条件）通过脚本集名字过滤。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCustomScriptSetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomScriptSetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportServerList struct {
	// 设备固资编号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器SN

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 设备型号

	SvrDeviceType *string `json:"SvrDeviceType,omitempty" name:"SvrDeviceType"`
	// IDC名称

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// 机架名称

	SvrRackName *string `json:"SvrRackName,omitempty" name:"SvrRackName"`
	// 存放机位ID

	SvrPosId *string `json:"SvrPosId,omitempty" name:"SvrPosId"`
	// 设备高度

	SvrDeviceHeight *string `json:"SvrDeviceHeight,omitempty" name:"SvrDeviceHeight"`
	// 服务器厂商

	SvrProducer *string `json:"SvrProducer,omitempty" name:"SvrProducer"`
	// 逻辑区域

	SvrLogicArea *string `json:"SvrLogicArea,omitempty" name:"SvrLogicArea"`
	// 服务器硬件描述

	SvrDeviceDescript *string `json:"SvrDeviceDescript,omitempty" name:"SvrDeviceDescript"`
	// region name

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// zone name

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// IDC id

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// 服务器架构

	SvrArch *string `json:"SvrArch,omitempty" name:"SvrArch"`
	// 服务器类型（0 物理机 or 3 黑石服务器 or 4 智能网卡）

	SvrType *string `json:"SvrType,omitempty" name:"SvrType"`
	// 关联SN(黑石服务器sn和智能网卡sn互相关联)

	SvrRelatedSn *string `json:"SvrRelatedSn,omitempty" name:"SvrRelatedSn"`
	// 导入失败的原因

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 节点label

	SvrLabels *string `json:"SvrLabels,omitempty" name:"SvrLabels"`
}

type ModifyLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleVMVirtualIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 回收VM虚拟内外网IP出参

		VMRecyVirtualIPSet []*VMVirtualIPSet `json:"VMRecyVirtualIPSet,omitempty" name:"VMRecyVirtualIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecycleVMVirtualIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleVMVirtualIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VMRecyLanIP struct {
	// 结果详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 服务器内网IP

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
	// 服务器序列号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
}

type ModifyServersRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID，system_id 和 system_key为双方约定，主要是进行请求方鉴权

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key，system_id 和 system_key为双方约定，主要是进行请求方鉴权

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"modify"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 修改服务器参数

	ModifyServerSet []*ModifyServer `json:"ModifyServerSet,omitempty" name:"ModifyServerSet"`
}

func (r *ModifyServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddIdcLineInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// IndexId

	IndexId *int64 `json:"IndexId,omitempty" name:"IndexId"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
}
