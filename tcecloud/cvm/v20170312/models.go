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

type HostOrder struct {
	// 订单所属帐号的应用id

	// 创建订单的帐号uin

	// 订单所属帐号的所有者uin

	// 订单发货的资源信息列表
}

type HostResource struct {
	// cdh实例总cpu核数

	CpuTotal *uint64 `json:"CpuTotal,omitempty" name:"CpuTotal"`
	// cdh实例可用cpu核数

	CpuAvailable *uint64 `json:"CpuAvailable,omitempty" name:"CpuAvailable"`
	// cdh实例总内存大小（单位为:GiB）

	MemTotal *float64 `json:"MemTotal,omitempty" name:"MemTotal"`
	// cdh实例可用内存大小（单位为:GiB）

	MemAvailable *float64 `json:"MemAvailable,omitempty" name:"MemAvailable"`
	// cdh实例总磁盘大小（单位为:GiB）

	DiskTotal *uint64 `json:"DiskTotal,omitempty" name:"DiskTotal"`
	// cdh实例可用磁盘大小（单位为:GiB）

	DiskAvailable *uint64 `json:"DiskAvailable,omitempty" name:"DiskAvailable"`
}

type InstanceTypeNameConfig struct {
	// 是否显示到实例类型列表

	ShowInMenu *bool `json:"ShowInMenu,omitempty" name:"ShowInMenu"`
	// 实例类型中文名

	InstanceFamilyName *string `json:"InstanceFamilyName,omitempty" name:"InstanceFamilyName"`
	// 实例类型

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
}

type PriceForHostTypeQuota struct {
	// 不打折价格。

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
}

type ZoneInfo struct {
	// 可用区名称，例如，ap-guangzhou-3

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区描述，例如，广州三区

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 可用区ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区状态

	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`
}

type ResetInstancesTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetInstancesTypeResponse) FromJsonString(s string) error {
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

type HostPrice struct {
	// 描述了cdh实例相关的价格信息

	HostPrice *ItemPrice `json:"HostPrice,omitempty" name:"HostPrice"`
}

type InternetBandwidthConfig struct {
	// 开始时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 实例带宽信息。

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
}

type DescribeInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例静态配置信息列表。

		InstanceConfigInfos []*InstanceConfigInfoItem `json:"InstanceConfigInfos,omitempty" name:"InstanceConfigInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceConfigInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceConfigInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnhancedService struct {
	// 开启云安全服务。若不指定该参数，则默认开启云安全服务。

	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`
	// 开启云安全服务。若不指定该参数，则默认开启云监控服务。

	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`
}

type Filter struct {
	// 需要过滤的字段。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段的过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type KeyPairInstancesinternetaccessible struct {
	// 密钥对的`ID`，是密钥对的唯一标识。

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 密钥对关联的实例`ID`列表。

	AssociatedInstanceIdSet []*string `json:"AssociatedInstanceIdSet,omitempty" name:"AssociatedInstanceIdSet"`
}

type QueryPlacement struct {
	// 可用区名称

	Zone *int64 `json:"Zone,omitempty" name:"Zone"`
	// VpcId

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 子网Id

	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`
	// 子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// Vpc名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// UnVpcId

	UnVpcId *string `json:"UnVpcId,omitempty" name:"UnVpcId"`
	// 交换机

	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type SystemDisk struct {
	// 系统盘类型。系统盘类型。取值范围：<br><li>LOCAL_BASIC：本地硬盘<br><li>LOCAL_SSD：本地SSD硬盘<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><br>默认取值：LOCAL_BASIC。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 系统盘ID。LOCAL_BASIC 和 LOCAL_SSD 类型没有ID。暂时不支持该参数。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 系统盘大小，单位：GB。默认值为 50

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type Placement struct {
	// 实例所属的[可用区](#zoneinfo)ID。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例所属项目ID。不填为默认项目。

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 实例所属的专用宿主机ID列表。如果您有购买专用宿主机并且指定了该参数，则您购买的实例就会随机的部署在这些专用宿主机上。当前暂不支持。

	HostIds []*string `json:"HostIds,omitempty" name:"HostIds"`
}

type DescribeInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest

	// 1

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeInstanceConfigInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceConfigInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Image struct {
	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像操作系统

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 镜像类型

	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`
	// 镜像创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像描述

	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`
	// 镜像大小

	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`
	// 镜像架构

	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
	// 镜像状态

	ImageState *string `json:"ImageState,omitempty" name:"ImageState"`
	// 镜像来源平台

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 镜像创建者

	ImageCreator *string `json:"ImageCreator,omitempty" name:"ImageCreator"`
	// 镜像来源

	ImageSource *string `json:"ImageSource,omitempty" name:"ImageSource"`
}

type HostForSellZoneStatus struct {
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可售卖情况, (SELL: 正常售卖，SOLD_OUT: 售罄)。

	Status *string `json:"Status,omitempty" name:"Status"`
}

type QueryDataDisk struct {
	// 数据盘大小，单位：GB。最小调整步长为10G，不同数据盘类型取值范围不同

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 数据盘类型：LOCAL_BASIC、CLOUD_BASIC、LOCAL_SSD、CLOUD_SSD、CLOUD_PREMIUM、CLOUD_ENHANCEDSSD

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 数据盘ID。LOCAL_BASIC 和 LOCAL_SSD 类型没有ID，暂时不支持该参数。 该参数目前仅用于DescribeInstances等查询类接口的返回参数，不可用于RunInstances等写接口的入参。数据盘id

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 数据盘是否随子机销毁。取值范围： TRUE：子机销毁时，销毁数据盘，只支持按小时后付费云盘 FALSE：子机销毁时，保留数据盘  默认取值：TRUE  该参数目前仅用于 RunInstances 接口。 注意：此字段可能返回 null，表示取不到有效值。

	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
}

type ResizeInstanceDisksRequest struct {
	*tchttp.BaseRequest

	// 待操作的实例ID。可通过[`DescribeInstances`](/ocloud/api/Compute/CVM/APIs/cvm运营端（opcvm）/版本（2019-06-25）/cvm运营端/DescribeInstances)接口返回值中的`InstanceId`获取。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 待扩容的数据盘配置信息。只支持扩容随实例购买的数据盘，且[数据盘类型](../数据结构#datadisk)为：`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`。数据盘容量单位：GB。最小扩容步长：10G。关于数据盘类型的选择请参考硬盘产品简介。可选数据盘类型受到实例类型`InstanceType`限制。另外允许扩容的最大容量也因数据盘类型的不同而有所差异。

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机<br><li>FALSE：表示在正常关机失败后不进行强制关机<br><br>默认取值：FALSE。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。

	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *ResizeInstanceDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResizeInstanceDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>默认取值：NOTIFY_AND_AUTO_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type InquiryResourceResetInstancesTypeRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](/ocloud/api/Compute/CVM/APIs/cvm运营端（opcvm）/版本（2019-06-25）/cvm运营端/DescribeInstances)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为1。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 是否在线升级配置

	Online *bool `json:"Online,omitempty" name:"Online"`
}

func (r *InquiryResourceResetInstancesTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryResourceResetInstancesTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Externals struct {
	// 释放地址

	ReleaseAddress *bool `json:"ReleaseAddress,omitempty" name:"ReleaseAddress"`
}

type InstanceStatus struct {
	// 实例`ID`。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// [实例状态](#instancestatus)。

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
}

type TagSpecification struct {
	// 标签绑定的资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 标签对列表

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type UserZoneStatusItem struct {
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 计费类型

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 售卖状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type ImageSharedAccount struct {
	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 账户ID

	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`
}

type InstanceReturnable struct {
	// 实例`ID`。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例是否可退还。

	IsReturnable *bool `json:"IsReturnable,omitempty" name:"IsReturnable"`
	// 实例退还失败错误码。

	ReturnFailCode *int64 `json:"ReturnFailCode,omitempty" name:"ReturnFailCode"`
	// 实例退还失败错误信息。

	ReturnFailMessage *string `json:"ReturnFailMessage,omitempty" name:"ReturnFailMessage"`
}

type InstanceFamilyConfig struct {
	// 机型族名称的中文全称。

	InstanceFamilyName *string `json:"InstanceFamilyName,omitempty" name:"InstanceFamilyName"`
	// 机型族名称的英文简称。

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
}

type RunMonitorServiceEnabled struct {
	// 是否开启云监控服务。取值范围：<br><li>TRUE：表示开启云监控服务<br><li>FALSE：表示不开启云监控服务<br><br>默认取值：TRUE。

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type DisasterRecoverGroup struct {
	// 容灾组id

	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`
	// 容灾组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 容灾组类型:HOST(物理机)、SW(交换机)。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 容灾组内最大容纳实例数量。

	CvmQuotaTotal *int64 `json:"CvmQuotaTotal,omitempty" name:"CvmQuotaTotal"`
	// 当前容灾组内实例数量。

	CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`
}

type RegionInfo struct {
	// 地域名称，例如，ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域描述，例如，华南地区(广州)

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域是否可用状态

	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`
}

type Address struct {
	// `EIP`的`ID`，是`EIP`的唯一标识。

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// `EIP`名称。

	AddressName *string `json:"AddressName,omitempty" name:"AddressName"`
	// `EIP`状态。

	AddressState *string `json:"AddressState,omitempty" name:"AddressState"`
	// 弹性外网IP

	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`
	// 绑定的资源实例`ID`。可能是一个`CVM`，`NAT`，或是弹性网卡。

	BindedResourceId *string `json:"BindedResourceId,omitempty" name:"BindedResourceId"`
	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type AddressChargePrepaid struct {
	// 购买实例的时长

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 自动续费标志

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type InstanceConnectivity struct {
	// 实例`ID`

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 默认远程登录端口连通性状态

	DefaultLoginPortConnectivity *bool `json:"DefaultLoginPortConnectivity,omitempty" name:"DefaultLoginPortConnectivity"`
	// ping包是否可达

	ICMPConnectivity *bool `json:"ICMPConnectivity,omitempty" name:"ICMPConnectivity"`
}

type InternetAccessibleModifyChargeType struct {
	// 网络付费模式

	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
	// 外网出带宽值

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

type InternetChargeTypeConfig struct {
	// 网络计费模式。

	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
	// 网络计费模式描述信息。

	Description *string `json:"Description,omitempty" name:"Description"`
}

type ZoneCpuQuota struct {
	// 可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例计费模式。(PREPAID：表示预付费，即包年包月。POSTPAID_BY_HOUR：表示后付费，即按量计费。CDHPAID：表示CDH付费，即只对CDH计费，不对CDH上的实例计费。)

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 可用CPU配额。
}

type QueryInstancesRequest struct {
	*tchttp.BaseRequest

	// 可填列表:uuid：实例uuid过滤 类型 listprivate-ip：实例内网ip 类型 listvpc-id：实例vpc-id 类型 listpublic-ip：实例公网ip 类型 listappid：实例appId 类型 listinstance-id：实例id 类型 listhost-ip：实例所在宿主机ip 类型 listipv6-address：实例ipv6地址 类型 operator-uin：实例创建者uin 类型 listfilter长度限制为10个，value长度限制为5  instance-name: 实例名称 类型 list image-id: 镜像id 类型 list instance-family 实例族 类型 list

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *QueryInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KvType struct {
	// 键的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 键所对应的值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type InstanceChargeTypeConfig struct {
	// 实例计费类型

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 实例计费类型描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type QuerySystemDisk struct {
	// 系统盘大小

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 系统盘类型：LOCAL_BASIC、CLOUD_BASIC、LOCAL_SSD、CLOUD_SSD、CLOUD_PREMIUM、CLOUD_ENHANCEDSSD

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 系统盘Id

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

type InquiryResourceResetInstancesTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例类型的资源

		InquiryResourceResetInstancesType []*ResourceForInstanceType `json:"InquiryResourceResetInstancesType,omitempty" name:"InquiryResourceResetInstancesType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryResourceResetInstancesTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryResourceResetInstancesTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Price struct {
	// 描述了实例价格。

	InstancePrice *ItemPrice `json:"InstancePrice,omitempty" name:"InstancePrice"`
	// 描述了网络价格。

	BandwidthPrice *ItemPrice `json:"BandwidthPrice,omitempty" name:"BandwidthPrice"`
}

type SharePermission struct {
	// 镜像分享时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 镜像分享的账户ID

	Account *string `json:"Account,omitempty" name:"Account"`
}

type RunSecurityServiceEnabled struct {
	// 是否开启云安全服务。取值范围：<br><li>TRUE：表示开启云安全服务<br><li>FALSE：表示不开启云安全服务<br><br>默认取值：TRUE。

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type AvailabilityZone struct {
	// 地域ID。

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 可用区ID。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区名称。

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 可用区状态。

	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`
}

type InstancesRecentFailedOperationSet struct {
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 操作事件类型

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 操作事件发生时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ItemPrice struct {
	// 后续单价，单位：元。

	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`
	// 后续计价单元，可取值范围： <br><li>HOUR：表示计价单元是按每小时来计算。当前涉及该计价单元的场景有：实例按小时后付费（POSTPAID_BY_HOUR）、带宽按小时后付费（BANDWIDTH_POSTPAID_BY_HOUR）：<br><li>GB：表示计价单元是按每GB来计算。当前涉及该计价单元的场景有：流量按小时后付费（TRAFFIC_POSTPAID_BY_HOUR）。

	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`
	// 预支费用的原价，单位：元。

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
	// 预支费用的折扣价，单位：元。

	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
}

type ClientSysInfo struct {
	// 操作系统

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 操作系统版本

	OsVersion *string `json:"OsVersion,omitempty" name:"OsVersion"`
	// 需要导入的系统盘数据盘信息

	DiskSize []*uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 额外信息

	ExtraInfo *string `json:"ExtraInfo,omitempty" name:"ExtraInfo"`
}

type Quota struct {
	// 配额名称，取值范围：<br><li>`TOTAL_EIP_QUOTA`：用户当前地域下EIP的配额数；<br><li>`DAILY_EIP_APPLY`：用户当前地域下今日申购次数；<br><li>`DAILY_PUBLIC_IP_ASSIGN`：用户当前地域下，重新分配公网 IP次数。

	QuotaId *string `json:"QuotaId,omitempty" name:"QuotaId"`
	// 当前数量

	QuotaCurrent *int64 `json:"QuotaCurrent,omitempty" name:"QuotaCurrent"`
	// 配额数量

	QuotaLimit *int64 `json:"QuotaLimit,omitempty" name:"QuotaLimit"`
}

type ResourceForInstanceType struct {
	// 内存大小，单位GB。

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// CPU核数

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 实例机型规格，比如S1.SMALL1、S1.SMALL2等。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例规格状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DataDisk struct {
	// 数据盘类型。数据盘类型。取值范围：<br><li>LOCAL_BASIC：本地硬盘<br><li>LOCAL_SSD：本地SSD硬盘<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_PREMIUM：高性能云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><br>默认取值：LOCAL_BASIC。<br><br>该参数对`ResizeInstanceDisk`接口无效。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 系统盘ID。LOCAL_BASIC 和 LOCAL_SSD 类型没有ID。暂时不支持该参数。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 数据盘大小，单位：GB。最小调整步长为10G，不同数据盘类型取值范围不同。默认值为0，表示不购买数据盘。更多限制详见产品文档。

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type HostGoodsDetailItem struct {
	// 请求事务id

	// 操作名称

	// 自动续费标记

	// pid

	// 购买或续费时长

	// 时间单位

	// 资源id

	// 当前到期时间

	// 数字签名

	// 产品信息项列表
}

type Tag struct {
	// 标签键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type HostItem struct {
	// cdh实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// cdh实例id

	HostId *string `json:"HostId,omitempty" name:"HostId"`
	// cdh实例类型

	HostType *string `json:"HostType,omitempty" name:"HostType"`
	// cdh实例名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// cdh实例付费模式

	HostChargeType *string `json:"HostChargeType,omitempty" name:"HostChargeType"`
	// cdh实例自动续费标记

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// cdh实例创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// cdh实例过期时间

	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// cdh实例上已创建云子机的实例id列表

	InstanceIds *string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// cdh实例状态

	HostState *string `json:"HostState,omitempty" name:"HostState"`
	// cdh实例ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// cdh实例资源信息

	HostResource *HostResource `json:"HostResource,omitempty" name:"HostResource"`
}

type InstanceConfigInfoItem struct {
	// 实例规格。

	// 实例规格名称。

	// 优先级。

	// 实例族信息列表。
}

type InstanceTypeItem struct {
	// 实例类型。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// CPU核数。

	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
	// 内存大小。

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// GPU核数。

	Gpu *uint64 `json:"Gpu,omitempty" name:"Gpu"`
	// FPGA核数。

	Fpga *uint64 `json:"Fpga,omitempty" name:"Fpga"`
	// 存储块数。

	StorageBlock *uint64 `json:"StorageBlock,omitempty" name:"StorageBlock"`
	// 网卡数。

	NetworkCard *uint64 `json:"NetworkCard,omitempty" name:"NetworkCard"`
	// 最大带宽。

	MaxBandwidth *float64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// 主频。

	Frequency *string `json:"Frequency,omitempty" name:"Frequency"`
	// CPU型号名称。

	CpuModelName *string `json:"CpuModelName,omitempty" name:"CpuModelName"`
	// 包转发率。

	Pps *uint64 `json:"Pps,omitempty" name:"Pps"`
	// 外部信息。

	Externals *Externals `json:"Externals,omitempty" name:"Externals"`
	// 备注信息。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type HostGoodsItem struct {
	// goodsCategoryId

	// 实例付费模式

	// 发货的实例个数

	// 地域id

	// 发起发货的用户uin

	// 发起发货帐号的所有者uin

	// 发起发货帐号对应的appId

	// 项目id

	// 可用区id

	// cdh实例详细信息
}

type HostTypeQuotaSet struct {
	// CPU核数，单位：核。

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// CPU型号。

	CpuModelName *string `json:"CpuModelName,omitempty" name:"CpuModelName"`
	// 磁盘大小，单位：GB。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 磁盘类型。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 专用宿主机机型系列。

	HostFamily *string `json:"HostFamily,omitempty" name:"HostFamily"`
	// 专用宿主机类型。

	HostType *string `json:"HostType,omitempty" name:"HostType"`
	// 内存大小，单位：GB。

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 售价。

	Price *Price `json:"Price,omitempty" name:"Price"`
	// 售卖状态，“ SOLD_OUT”：售罄，“SELL”：售卖中

	Status *string `json:"Status,omitempty" name:"Status"`
	// 可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type InstanceFamilyItem struct {
	// 实例族。

	// 优先级。

	// 实例类型信息列表。
}

type LoginSettings struct {
	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>Linux实例密码必须8到16位，至少包括两项[a-z，A-Z]、[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。<br><li>Windows实例密码必须12到16位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = { } [ ] : ; ' , . ? /]中的特殊符号。<br><br>若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。

	Password *string `json:"Password,omitempty" name:"Password"`
	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：<br><li>TRUE：表示保持镜像的登录设置<br><li>FALSE：表示不保持镜像的登录设置<br><br>默认取值：FALSE。

	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

type ProductInfoItem struct {
	// 信息项名称

	// 信息项对应的值
}

type QueryInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 租户端云服务器实例总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 租户端云服务器实例列表

		Instances []*QueryInstance `json:"Instances,omitempty" name:"Instances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>默认取值：NOTIFY_AND_AUTO_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type InstanceTypeConfig struct {
	// 可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例机型。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例机型系列。

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// GPU核数，单位：核。

	GPU *int64 `json:"GPU,omitempty" name:"GPU"`
	// CPU核数，单位：核。

	CPU *int64 `json:"CPU,omitempty" name:"CPU"`
	// 内存容量，单位：`GB`。

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 是否支持云硬盘。取值范围：<br><li>`TRUE`：表示支持云硬盘；<br><li>`FALSE`：表示不支持云硬盘。

	CbsSupport *string `json:"CbsSupport,omitempty" name:"CbsSupport"`
	// 机型状态。取值范围：<br><li>`AVAILABLE`：表示机型可用；<br><li>`UNAVAILABLE`：表示机型不可用。

	InstanceTypeState *string `json:"InstanceTypeState,omitempty" name:"InstanceTypeState"`
}

type VirtualPrivateCloud struct {
	// 私有网络ID，形如`vpc-xxx`。有效的VpcId可通过登录[控制台](//console.{{conf.main_domain}}/vpc/vpc?rid=1)查询；也可以调用接口 [DescribeVpcEx](/ocloud/api/NetWork/VPC/APIs/云网络运营端api%20v3（opvpc）/版本（2020-02-14）/云网络运营端api%20v3接口/DescribeVpcEx) ，从接口返回中的`unVpcId`字段获取。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络子网ID，形如`subnet-xxx`。有效的私有网络子网ID可通过登录[控制台](//console.{{conf.main_domain}}/vpc/subnet?rid=1)查询；也可以调用接口  [DescribeSubnetEx](/ocloud/api/NetWork/VPC/APIs/云网络运营端api%20v3（opvpc）/版本（2020-02-14）/云网络运营端api%20v3接口/DescribeSubnetEx) ，从接口返回中的`unSubnetId`字段获取。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 是否用作公网网关。公网网关只有在实例拥有公网IP以及处于私有网络下时才能正常使用。取值范围：<br><li>TRUE：表示用作公网网关<br><li>FALSE：表示不用作公网网关<br><br>默认取值：FALSE。

	AsVpcGateway *bool `json:"AsVpcGateway,omitempty" name:"AsVpcGateway"`
	// 私有网络子网 IP 数组，在创建实例、修改实例vpc属性操作中可使用此参数。当前仅批量创建多台实例时支持传入相同子网的多个 IP。

	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
}

type ResizeInstanceDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResizeInstanceDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResizeInstanceDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostTypeQuota struct {
	// CPU核数，单位：核。

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// CPU型号。

	CpuModelName *string `json:"CpuModelName,omitempty" name:"CpuModelName"`
	// 磁盘大小，单位：GB。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 磁盘类型。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 专用宿主机机型系列。

	HostFamily *string `json:"HostFamily,omitempty" name:"HostFamily"`
	// 专用宿主机类型。

	HostType *string `json:"HostType,omitempty" name:"HostType"`
	// 内存大小，单位：GB。

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 售价。

	Price *PriceForHostTypeQuota `json:"Price,omitempty" name:"Price"`
	// 售卖状态，“ SOLD_OUT”：售罄，“SELL”：售卖中

	Status *string `json:"Status,omitempty" name:"Status"`
	// 可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type ResetInstancesTypeRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](/ocloud/api/Compute/CVM/APIs/cvm运营端（opcvm）/版本（2019-06-25）/cvm运营端/DescribeInstances)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为1。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 实例机型。不同实例机型指定了不同的资源规格，具体取值可参见附表实例资源规格对照表，也可以调用查询实例资源规格列表接口获得最新的规格表。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机<br><li>FALSE：表示在正常关机失败后不进行强制关机<br><br>默认取值：FALSE。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。

	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *ResetInstancesTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetInstancesTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceTypeQuota struct {
	// 可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例机型。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例机型配额。

	InstanceQuota *int64 `json:"InstanceQuota,omitempty" name:"InstanceQuota"`
	// 实例计费模式。取值范围： <br>*`PREPAID`：表示预付费，即包年包月 <br>* `POSTPAID_BY_HOUR`：表示后付费，即按量计费 * `CDHPAID`：CDH付费，即只对[CDH(/document/product/416)]计费，不对CDH上的实例计费。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 实例类型。

	DeviceClass *string `json:"DeviceClass,omitempty" name:"DeviceClass"`
	// 实例的CPU核数，单位：核。

	CPU *int64 `json:"CPU,omitempty" name:"CPU"`
	// 实例内存容量，单位：`GB`。

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 实例的GPU核数，单位：核。

	GPU *int64 `json:"GPU,omitempty" name:"GPU"`
}

type QueryInstance struct {
	// Placement结构

	Placement *QueryPlacement `json:"Placement,omitempty" name:"Placement"`
	// 内存大小

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// cpu核数

	CPU *int64 `json:"CPU,omitempty" name:"CPU"`
	// 创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例拥有者AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 内网Ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网Ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 所在宿主机ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// Ipv6地址

	IPv6Addresses []*string `json:"IPv6Addresses,omitempty" name:"IPv6Addresses"`
	// 当前状态

	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`
	// 当前状态

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 所有者ownerUin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 子机机型规格

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 子机机型类型

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// Quota使用情况

	NodeQuota []*int64 `json:"NodeQuota,omitempty" name:"NodeQuota"`
	// 系统盘详情

	SystemDisk *QuerySystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 数据盘详情

	DataDisks []*QueryDataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 系统名

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 镜像Id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 操作人Uin

	OperatorUin *string `json:"OperatorUin,omitempty" name:"OperatorUin"`
}

type InternetAccessible struct {
	// 网络计费类型。取值范围：<br><li>BANDWIDTH_PREPAID：预付费按带宽结算<br><li>TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费<br><li>BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费<br><li>BANDWIDTH_PACKAGE：带宽包用户<br>默认取值：TRAFFIC_POSTPAID_BY_HOUR。

	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
	// 公网出带宽上限，单位：Mbps。默认值：0Mbps。不同机型带宽上限范围不一致。

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 是否分配公网IP。取值范围：<br><li>TRUE：表示分配公网IP<br><li>FALSE：表示不分配公网IP<br><br>当公网带宽大于0Mbps时，可自由选择开通与否，默认开通公网IP；当公网带宽为0，则不允许分配公网IP。

	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`
}

type KeyPair struct {
	// 密钥对的`ID`，是密钥对的唯一标识。

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 密钥对名称。

	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
	// 密钥对所属的项目`ID`。

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 密钥对描述信息。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 密钥对的纯文本公钥。

	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
	// 密钥对的纯文本私钥。腾讯云不会保管私钥，请用户自行妥善保存。

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	// 密钥关联的实例`ID`列表。

	AssociatedInstanceIds []*string `json:"AssociatedInstanceIds,omitempty" name:"AssociatedInstanceIds"`
	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type InstanceOrder struct{}

type ImageAttribute struct {
	// unImgId

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// deviceImageId

	InnerImageId *string `json:"InnerImageId,omitempty" name:"InnerImageId"`
}

type Instance struct {
	// 实例所在的位置。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 实例`ID`。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例机型。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例的CPU核数，单位：核。

	CPU *int64 `json:"CPU,omitempty" name:"CPU"`
	// 实例内存容量，单位：`GB`。

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 实例业务状态。取值范围：<br><li>NORMAL：表示正常状态的实例<br><li>EXPIRED：表示过期的实例<br><li>PROTECTIVELY_ISOLATED：表示被安全隔离的实例。

	RestrictState *string `json:"RestrictState,omitempty" name:"RestrictState"`
	// 实例名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例计费模式。取值范围：<br><li>`PREPAID`：表示预付费，即包年包月<br><li>`POSTPAID_BY_HOUR`：表示后付费，即按量计费<br><li>`CDHPAID`：`CDH`付费，即只对`CDH`计费，不对`CDH`上的实例计费。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 实例系统盘信息。

	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 实例数据盘信息。只包含随实例购买的数据盘。

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 实例主网卡的内网`IP`列表。

	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 实例主网卡的公网`IP`列表。

	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`
	// 实例带宽信息。

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 实例所属虚拟私有网络信息。

	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
	// 生产实例所使用的镜像`ID`。

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 自动续费标识。取值范围：<br><li>`NOTIFY_AND_MANUAL_RENEW`：表示通知即将过期，但不自动续费<br><li>`NOTIFY_AND_AUTO_RENEW`：表示通知即将过期，而且自动续费<br><li>`DISABLE_NOTIFY_AND_MANUAL_RENEW`：表示不通知即将过期，也不自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 到期时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
}

type SharePermissionSet struct {
	// 镜像分享时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 镜像分享的账户ID

	Account *string `json:"Account,omitempty" name:"Account"`
}
