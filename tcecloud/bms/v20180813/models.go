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

package v20180813

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type BmsQuota struct {
	// 资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源配额

	ResourceQuota *uint64 `json:"ResourceQuota,omitempty" name:"ResourceQuota"`
	// 实际占用数量

	Realused *uint64 `json:"Realused,omitempty" name:"Realused"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 配额所属的AppId值

	Appid *string `json:"Appid,omitempty" name:"Appid"`
}

type Disk struct {
	// 硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 硬盘所在的位置。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 硬盘名称。

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 硬盘大小。

	DiskSize *string `json:"DiskSize,omitempty" name:"DiskSize"`
	// 硬盘挂载的云主机ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 云盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：SSD表示SSD云硬盘。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 硬盘的创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
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
	// 网口速率

	Netspeed *int64 `json:"Netspeed,omitempty" name:"Netspeed"`
	// 区ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 支持的RAID

	Raid *string `json:"Raid,omitempty" name:"Raid"`
}

type LoginSettings struct {
	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>Linux实例密码必须8到16位，至少包括两项[a-z，A-Z]、[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。<br><li>Windows实例密码必须12到16位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = { } [ ] : ; ' , . ? /]中的特殊符号。<br><br>若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。

	Password *string `json:"Password,omitempty" name:"Password"`
	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：<br><li>TRUE：表示保持镜像的登录设置<br><li>FALSE：表示不保持镜像的登录设置<br><br>默认取值：FALSE。

	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

type TagSpecification struct {
	// 标签绑定的资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 标签对列表

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type DescribeBmsDeployTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 部署实例的任务列表。

		DeployTasks []*DeployTasks `json:"DeployTasks,omitempty" name:"DeployTasks"`
		// 实例部署任务个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// igniter部署实例的任务ID

		DeployTaskId *string `json:"DeployTaskId,omitempty" name:"DeployTaskId"`
		// 部署实例所在机器的Sn

		Sn *string `json:"Sn,omitempty" name:"Sn"`
		// 部署实例所在机器的IDC

		Idc *uint64 `json:"Idc,omitempty" name:"Idc"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBmsDeployTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBmsDeployTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionTimer struct {
	// 定时器

	TimerAction *string `json:"TimerAction,omitempty" name:"TimerAction"`
	// 执行时间

	ActionTime *string `json:"ActionTime,omitempty" name:"ActionTime"`
	// 扩展数据

	Externals *Externals `json:"Externals,omitempty" name:"Externals"`
}

type Reply struct {
	// Code

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// Msg

	Msg *string `json:"Msg,omitempty" name:"Msg"`
}

type DeployTasks struct {
	// 状态码：
	// DCOS SubTask状态码：
	//  2      带外关机
	//  19     开机pxe
	//  10002  等待pxe上报，10min没上报超时，任务失败 （TPC agent）
	//  10003  RAID采集，采集成功即需求raid和机器现在raid一致，就不执行10004 直接执行10005 MINIOS安装，如果partition_flag=1，则不会有10003步.
	//  10004  RAID修改，如果partition_flag=1，必然会执行。
	//  10005  MINIOS安装
	//  10006  等待minios状态上报，10min没上报超时，任务失败
	//  10009  os安装第一步
	//  10010  os安装第二步
	//  10011  os安装第三步
	//  10016  os安装好后等待系统初始化重启，10min
	//  10013  ping检查
	//  10014  ssh检查
	//  10015  密码修改
	//  10017  安装插件
	//  10021  执行后置脚本
	//
	// BMS SubTask状态码：
	//  90001  设置网络Overlay配置
	//  90002  清理网络Overlay配置

	StepCode *int64 `json:"StepCode,omitempty" name:"StepCode"`
	// 状态信息：0=未执行，1=执行成功，2=执行失败，3=执行中。

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 任务执行的详细输出

	Details *string `json:"Details,omitempty" name:"Details"`
	// 任务执行消耗的时间

	ExecutionTime *int64 `json:"ExecutionTime,omitempty" name:"ExecutionTime"`
}

type Flavor struct {
	// 套餐ID。

	FlavorId *string `json:"FlavorId,omitempty" name:"FlavorId"`
	// 套餐名称。

	FlavorName *string `json:"FlavorName,omitempty" name:"FlavorName"`
	// 实例所在的位置。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 支持的Raid类型。

	RaidType []*string `json:"RaidType,omitempty" name:"RaidType"`
	// 支持的系统列表。

	OperatingSystem *OperatingSystem `json:"OperatingSystem,omitempty" name:"OperatingSystem"`
	// cpu信息。

	Cpu *string `json:"Cpu,omitempty" name:"Cpu"`
	// 内存信息。

	Memory *string `json:"Memory,omitempty" name:"Memory"`
	// 硬盘信息。

	SystemDisk *string `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 网卡信息

	NetSpeed *string `json:"NetSpeed,omitempty" name:"NetSpeed"`
	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 计费类型

	FlavorType *string `json:"FlavorType,omitempty" name:"FlavorType"`
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

type Filter struct {
	// 需要过滤的字段。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段的过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type InternetAccessible struct {
	// 公网出带宽上限，单位：Mbps。默认值：0Mbps。不同机型带宽上限范围不一致，具体限制详见[购买网络带宽]。

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 是否分配公网IP。取值范围：<br><li>TRUE：表示分配公网IP<br><li>FALSE：表示不分配公网IP<br><br>当公网带宽大于0Mbps时，可自由选择开通与否，默认开通公网IP；当公网带宽为0，则不允许分配公网IP。

	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`
}

type AddTime struct {
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 截止时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type Externals struct{}

type Instance struct {
	// 实例所在的位置。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 实例`ID`。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// Raid类型。

	RaidType *string `json:"RaidType,omitempty" name:"RaidType"`
	// 操作系统类型。

	OperatingSystemType *string `json:"OperatingSystemType,omitempty" name:"OperatingSystemType"`
	// 操作系统发行版本

	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`
	// 实例主网卡的内网`IP`列表。

	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 实例所属虚拟私有网络信息。

	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
	// 套餐信息。

	FlavorId *string `json:"FlavorId,omitempty" name:"FlavorId"`
	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 实例状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type Placement struct {
	// 实例所属的可用区(zone)ID。该参数也可以通过调用  [DescribeZones]的返回值中的Zone字段来获取。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例所属项目ID。该参数可以通过调用 [DescribeProject]的返回值中的 projectId 字段来获取。不填为默认项目。

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type EnhancedService struct {
	// 开启云安全服务。若不指定该参数，则默认开启

	SecurityService *bool `json:"SecurityService,omitempty" name:"SecurityService"`
	// 开启云安全服务。若不指定该参数，则默认开启云监控服务。

	MonitorService *bool `json:"MonitorService,omitempty" name:"MonitorService"`
}

type DescribeBmsDeployTasksRequest struct {
	*tchttp.BaseRequest

	// BMS的实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeBmsDeployTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBmsDeployTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSwitchInfo struct {
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

type OperatingSystem struct {
	// 支持的linux系统列表

	Linux []*string `json:"Linux,omitempty" name:"Linux"`
	// 支持的windows系统列表

	Windows []*string `json:"Windows,omitempty" name:"Windows"`
}

type Tag struct {
	// 标签键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type VirtualPrivateCloud struct {
	// 私有网络ID，形如`vpc-xxx`。有效的VpcId可通过登录[控制台]查询；也可以调用接口 [DescribeVpcEx] ，从接口返回中的`unVpcId`字段获取。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络子网ID，形如`subnet-xxx`。有效的私有网络子网ID可通过登录[控制台]查询；也可以调用接口  [DescribeSubnetEx]，从接口返回中的`unSubnetId`字段获取。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 是否用作公网网关。公网网关只有在实例拥有公网IP以及处于私有网络下时才能正常使用。取值范围：<br><li>TRUE：表示用作公网网关<br><li>FALSE：表示不用作公网网关<br><br>默认取值：FALSE。

	AsVpcGateway *bool `json:"AsVpcGateway,omitempty" name:"AsVpcGateway"`
	// 私有网络子网 IP 数组，在创建实例、修改实例vpc属性操作中可使用此参数。当前仅批量创建多台实例时支持传入相同子网的多个 IP。

	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
}
