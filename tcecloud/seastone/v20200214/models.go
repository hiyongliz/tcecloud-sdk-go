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

package v20200214

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest

	// 节点ID列表

	DeviceIds []*string `json:"DeviceIds,omitempty" name:"DeviceIds"`
	// 集群ID

	SetId *string `json:"SetId,omitempty" name:"SetId"`
	// 所在VPC网络ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 网口IP

	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 单次条目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 升序or降序

	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AclEntry struct {
	// 子网

	Subnet *string `json:"Subnet,omitempty" name:"Subnet"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 协议, 取值: TCP,UDP, ICMP, ALL。

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// LowerPort

	LowerPort *uint64 `json:"LowerPort,omitempty" name:"LowerPort"`
	// UpperPort

	UpperPort *uint64 `json:"UpperPort,omitempty" name:"UpperPort"`
	// IntMask

	IntMask *uint64 `json:"IntMask,omitempty" name:"IntMask"`
	// 0-入站规则 1-出站规则

	Dir *uint64 `json:"Dir,omitempty" name:"Dir"`
	// 0-允许 1-拒绝

	Op *uint64 `json:"Op,omitempty" name:"Op"`
}

type ControllerSet struct {
	// 唯一ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否独占

	IsShared *bool `json:"IsShared,omitempty" name:"IsShared"`
	// 部署位置

	WorkSite *string `json:"WorkSite,omitempty" name:"WorkSite"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 策略唯一ID

	WhiteListId *string `json:"WhiteListId,omitempty" name:"WhiteListId"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 包含节点数

	NodeCnt *uint64 `json:"NodeCnt,omitempty" name:"NodeCnt"`
	// 规则数

	RuleCnt *uint64 `json:"RuleCnt,omitempty" name:"RuleCnt"`
}

type DeleteDetectInstallRecord struct {
	// EventId

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// Ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type Ipv6AddressSet struct {
	// Primary

	Primary *bool `json:"Primary,omitempty" name:"Primary"`
	// Address

	Address *string `json:"Address,omitempty" name:"Address"`
	// AddressId

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// IsWanIpBlocked

	IsWanIpBlocked *bool `json:"IsWanIpBlocked,omitempty" name:"IsWanIpBlocked"`
	// State

	State *string `json:"State,omitempty" name:"State"`
	// Description

	Description *string `json:"Description,omitempty" name:"Description"`
}

type JnsGatewayServiceP struct {
	// DpdkSnatip

	DpdkSnatip *string `json:"DpdkSnatip,omitempty" name:"DpdkSnatip"`
	// Min

	Min *uint64 `json:"Min,omitempty" name:"Min"`
	// UniqVgwIndex

	UniqVgwIndex *string `json:"UniqVgwIndex,omitempty" name:"UniqVgwIndex"`
	// VgwIndex

	VgwIndex *string `json:"VgwIndex,omitempty" name:"VgwIndex"`
	// PipSetId

	PipSetId *uint64 `json:"PipSetId,omitempty" name:"PipSetId"`
	// PoolId

	PoolId *uint64 `json:"PoolId,omitempty" name:"PoolId"`
	// Vgw

	Vgw *string `json:"Vgw,omitempty" name:"Vgw"`
	// VoteGwRandomFactor

	VoteGwRandomFactor *string `json:"VoteGwRandomFactor,omitempty" name:"VoteGwRandomFactor"`
	// VpcId

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// Business

	Business *string `json:"Business,omitempty" name:"Business"`
	// Max

	Max *uint64 `json:"Max,omitempty" name:"Max"`
	// Proto

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// PipSubnet

	PipSubnet *string `json:"PipSubnet,omitempty" name:"PipSubnet"`
	// SubnetId

	SubnetId *uint64 `json:"SubnetId,omitempty" name:"SubnetId"`
	// Count

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// PipMask

	PipMask *string `json:"PipMask,omitempty" name:"PipMask"`
	// Pport

	Pport *uint64 `json:"Pport,omitempty" name:"Pport"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// BusinessOwner

	BusinessOwner *string `json:"BusinessOwner,omitempty" name:"BusinessOwner"`
	// UniqVpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// VgwType

	VgwType *uint64 `json:"VgwType,omitempty" name:"VgwType"`
	// Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// UniqSubnetId

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
	// Pip

	Pip *string `json:"Pip,omitempty" name:"Pip"`
	// GwIp

	GwIp *string `json:"GwIp,omitempty" name:"GwIp"`
	// Vport

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Vm

	Vm []*JnsGatewayServicePVm `json:"Vm,omitempty" name:"Vm"`
}

type L7RuleInfo struct {
	// location list

	LocationSet []*LocationInfo `json:"LocationSet,omitempty" name:"LocationSet"`
	// virtual service

	VirtualService *VirtualServiceInfo `json:"VirtualService,omitempty" name:"VirtualService"`
}

type Route struct {
	// 目的端cidr

	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
	// 下一跳类型，0 子机网关 1 VPNGW 3 专线网关 4 跨地域互通 1025 vpcgw

	NextType *uint64 `json:"NextType,omitempty" name:"NextType"`
	// 下一跳地址

	NextHub *string `json:"NextHub,omitempty" name:"NextHub"`
	// 下一跳地址唯一 ID

	UnNextHub *string `json:"UnNextHub,omitempty" name:"UnNextHub"`
	// 路由描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type ServiceRouteRoute struct {
	// Pip

	Pip *string `json:"Pip,omitempty" name:"Pip"`
	// Pport

	Pport *uint64 `json:"Pport,omitempty" name:"Pport"`
	// Description

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DescribeWhiteListsRequest struct {
	*tchttp.BaseRequest

	// 实例id列表

	WhiteListIds []*string `json:"WhiteListIds,omitempty" name:"WhiteListIds"`
	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 单次条目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 升序or降序

	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeWhiteListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWhiteListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteControllerSetsRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表

	ControllerSetIds []*string `json:"ControllerSetIds,omitempty" name:"ControllerSetIds"`
}

func (r *DeleteControllerSetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteControllerSetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnatgwWip struct {
	// CarrierId

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// WanIp

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// UnatgwId

	UnatgwId *uint64 `json:"UnatgwId,omitempty" name:"UnatgwId"`
}

type DeleteGatewaySetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGatewaySetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGatewaySetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlgorithmsRequest struct {
	*tchttp.BaseRequest

	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 单次条目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 升序or降序

	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAlgorithmsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlgorithmsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateLimitRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigrateLimitRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateLimitRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AclSubNetInfo struct {
	// Subnet

	Subnet *string `json:"Subnet,omitempty" name:"Subnet"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// BDefault

	BDefault *string `json:"BDefault,omitempty" name:"BDefault"`
	// Mask

	Mask *string `json:"Mask,omitempty" name:"Mask"`
	// ZoneId

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// IntMask

	IntMask *string `json:"IntMask,omitempty" name:"IntMask"`
	// UniqSubnetId

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
	// SubnetId

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// DhcpFlag

	DhcpFlag *string `json:"DhcpFlag,omitempty" name:"DhcpFlag"`
}

type SgVmCount struct {
	// general

	General *string `json:"General,omitempty" name:"General"`
	// cvm

	Cvm *string `json:"Cvm,omitempty" name:"Cvm"`
}

type VpcLimit struct {
	// ProdName

	ProdName *string `json:"ProdName,omitempty" name:"ProdName"`
	// Val

	Val *uint64 `json:"Val,omitempty" name:"Val"`
	// QuotaDesc

	QuotaDesc *string `json:"QuotaDesc,omitempty" name:"QuotaDesc"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// MaxQuota

	MaxQuota *uint64 `json:"MaxQuota,omitempty" name:"MaxQuota"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// DefaultQuota

	DefaultQuota *uint64 `json:"DefaultQuota,omitempty" name:"DefaultQuota"`
	// Revisability

	Revisability *uint64 `json:"Revisability,omitempty" name:"Revisability"`
}

type DeleteGatewaySetsRequest struct {
	*tchttp.BaseRequest

	// 实例Id列表

	GatewaySetIds []*string `json:"GatewaySetIds,omitempty" name:"GatewaySetIds"`
}

func (r *DeleteGatewaySetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGatewaySetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeControllerSetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		InstanceSet []*ControllerSet `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 实例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeControllerSetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeControllerSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetectInstallState struct {
	// EventId

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// FinishTime

	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
	// Status

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// Version

	Version *string `json:"Version,omitempty" name:"Version"`
	// AzName

	AzName *string `json:"AzName,omitempty" name:"AzName"`
	// Gwname

	Gwname *string `json:"Gwname,omitempty" name:"Gwname"`
	// Msg

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// Ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Groupid

	Groupid *string `json:"Groupid,omitempty" name:"Groupid"`
}

type UNatGwWanIp struct {
	// WanIp

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// CarrierId

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
	// UnatgwId

	UnatgwId *uint64 `json:"UnatgwId,omitempty" name:"UnatgwId"`
	// UniqUnatgwId

	UniqUnatgwId *string `json:"UniqUnatgwId,omitempty" name:"UniqUnatgwId"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type JnsGatewayService struct {
	// VpcId

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// DpdkSnatip

	DpdkSnatip *string `json:"DpdkSnatip,omitempty" name:"DpdkSnatip"`
	// Business

	Business *string `json:"Business,omitempty" name:"Business"`
	// Business

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// UniqVpcId

	UniqVpcId *uint64 `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// VgwType

	VgwType *uint64 `json:"VgwType,omitempty" name:"VgwType"`
	// UniqVgwIndex

	UniqVgwIndex *string `json:"UniqVgwIndex,omitempty" name:"UniqVgwIndex"`
	// Pport

	Pport *uint64 `json:"Pport,omitempty" name:"Pport"`
	// Pip

	Pip *string `json:"Pip,omitempty" name:"Pip"`
	// Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// UniqSubnetId

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
	// VgwIndex

	VgwIndex *string `json:"VgwIndex,omitempty" name:"VgwIndex"`
	// PoolId

	PoolId *uint64 `json:"PoolId,omitempty" name:"PoolId"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// GwIp

	GwIp *string `json:"GwIp,omitempty" name:"GwIp"`
	// BusinessOwner

	BusinessOwner *string `json:"BusinessOwner,omitempty" name:"BusinessOwner"`
	// SubnetId

	SubnetId *uint64 `json:"SubnetId,omitempty" name:"SubnetId"`
	// Vport

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// Snatip

	Snatip *int64 `json:"Snatip,omitempty" name:"Snatip"`
}

type LimitItem struct {
	// 细项类型，如：INBAND

	Type *string `json:"Type,omitempty" name:"Type"`
	// 限速配额，单位Mbps

	Quota *uint64 `json:"Quota,omitempty" name:"Quota"`
	// 采用算法，默认为AutoAdjust

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 单位进阶，如：1000或1024，默认为1000

	Unit *uint64 `json:"Unit,omitempty" name:"Unit"`
}

type VirtualServiceInfo struct {
	// app name

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// bg name

	BgName *string `json:"BgName,omitempty" name:"BgName"`
	// biz id

	BizId *string `json:"BizId,omitempty" name:"BizId"`
	// biz type

	BizType *string `json:"BizType,omitempty" name:"BizType"`
	// blocaked enable

	BlockedEnable *string `json:"BlockedEnable,omitempty" name:"BlockedEnable"`
	// cc guard enable

	CcGuardEnable *string `json:"CcGuardEnable,omitempty" name:"CcGuardEnable"`
	// city

	City *string `json:"City,omitempty" name:"City"`
	// default server

	DefaultServer *string `json:"DefaultServer,omitempty" name:"DefaultServer"`
	// domain

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// forward protocol

	ForwardProtocol *string `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`
	// fwd mode

	FwdMode *string `json:"FwdMode,omitempty" name:"FwdMode"`
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// ld port

	LdPort *string `json:"LdPort,omitempty" name:"LdPort"`
	// 机房名字

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// set name

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// vip数组

	Vips []*VipInfo `json:"Vips,omitempty" name:"Vips"`
	// vpc id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 端口列表

	Vports []*int64 `json:"Vports,omitempty" name:"Vports"`
}

type DeleteDevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLimitRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		InstanceSet []*LimitRule `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 实例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLimitRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLimitRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EniSet struct {
	// PipWanipDesc

	PipWanipDesc []*PipWanDesc `json:"PipWanipDesc,omitempty" name:"PipWanipDesc"`
	// Ipv6AddressSet

	Ipv6AddressSet []*Ipv6AddressSet `json:"Ipv6AddressSet,omitempty" name:"Ipv6AddressSet"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// VpcId

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// InstanceId

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// UniqEniId

	UniqEniId *string `json:"UniqEniId,omitempty" name:"UniqEniId"`
	// IfnName

	IfnName *string `json:"IfnName,omitempty" name:"IfnName"`
	// WanIp

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// HostIp

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Pip

	Pip *string `json:"Pip,omitempty" name:"Pip"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// Description

	Description *string `json:"Description,omitempty" name:"Description"`
}

type Ipv6Set struct {
	// UniqEniId

	UniqEniId *string `json:"UniqEniId,omitempty" name:"UniqEniId"`
	// Info

	Info []*VmIpv6 `json:"Info,omitempty" name:"Info"`
}

type SubnetSubMask struct {
	// IntMask

	IntMask *uint64 `json:"IntMask,omitempty" name:"IntMask"`
	// Mask

	Mask *string `json:"Mask,omitempty" name:"Mask"`
	// Subnet

	Subnet *string `json:"Subnet,omitempty" name:"Subnet"`
}

type UsgPolicyInfo struct {
	// ErrorCode

	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// InBound

	InBound []*InOutBoundInfo `json:"InBound,omitempty" name:"InBound"`
	// ErrorInfo

	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
	// OutBound

	OutBound []*InOutBoundInfo `json:"OutBound,omitempty" name:"OutBound"`
	// Version

	Version *string `json:"Version,omitempty" name:"Version"`
}

type CreateLimitRuleRequest struct {
	*tchttp.BaseRequest

	// 集群唯一键值

	ControllerSetId *string `json:"ControllerSetId,omitempty" name:"ControllerSetId"`
	// 产品类型

	ProductType *string `json:"ProductType,omitempty" name:"ProductType"`
	// 规则string类型唯一键

	StrKey *string `json:"StrKey,omitempty" name:"StrKey"`
	// 规则int类型唯一键

	IntKey []*uint64 `json:"IntKey,omitempty" name:"IntKey"`
	// 限速细项

	LimitItems []*LimitItem `json:"LimitItems,omitempty" name:"LimitItems"`
	// 规则范围，以集群id为基本信息

	LimitScope []*string `json:"LimitScope,omitempty" name:"LimitScope"`
	// 规则范围，以集群VIP为基本信息，与LimitScope二选一

	LimitVipScope []*string `json:"LimitVipScope,omitempty" name:"LimitVipScope"`
	// 所在虚拟网络id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 调试开关

	Debug *uint64 `json:"Debug,omitempty" name:"Debug"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *CreateLimitRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLimitRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		InstanceSet []*WhiteList `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 实例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWhiteListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWhiteListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActiveHealthCheckInfo struct {
	// DownTimes

	DownTimes *string `json:"DownTimes,omitempty" name:"DownTimes"`
	// expect alive

	ExpectAlive *string `json:"ExpectAlive,omitempty" name:"ExpectAlive"`
	// http send

	HttpSend *string `json:"HttpSend,omitempty" name:"HttpSend"`
	// interval

	Interval *string `json:"Interval,omitempty" name:"Interval"`
	// method

	Method *string `json:"Method,omitempty" name:"Method"`
	// protocol

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// server name

	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`
	// up times

	UpTimes *string `json:"UpTimes,omitempty" name:"UpTimes"`
}

type UserGateway struct {
	// UsrgwId

	UsrgwId *uint64 `json:"UsrgwId,omitempty" name:"UsrgwId"`
	// UniqUsrgwId

	UniqUsrgwId *string `json:"UniqUsrgwId,omitempty" name:"UniqUsrgwId"`
	// Ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateGatewaySetRequest struct {
	*tchttp.BaseRequest

	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 网关类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 网关模式（主备/负载）

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 集群所在vpcid

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vip列表

	VipList []*string `json:"VipList,omitempty" name:"VipList"`
	// 限速开关

	LimitSwitch *bool `json:"LimitSwitch,omitempty" name:"LimitSwitch"`
	// 描述符

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *CreateGatewaySetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGatewaySetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGatewaySetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		InstanceSet []*GatewaySet `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 实例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGatewaySetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGatewaySetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWhiteListRequest struct {
	*tchttp.BaseRequest

	// 命名，要求必须命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 匹配细项

	MatchField []*MatchField `json:"MatchField,omitempty" name:"MatchField"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *CreateWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateControllersRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表

	ControllerIds []*string `json:"ControllerIds,omitempty" name:"ControllerIds"`
	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *UpdateControllersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateControllersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Algorithm struct {
	// 唯一键值

	Id *string `json:"Id,omitempty" name:"Id"`
	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 新建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DirectConnectGateway struct {
	// VPC实例ID。如66379

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// VPC唯一ID。如uniqVpcId-45678

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// VPC名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// VPG实例ID。如204

	VpgId *uint64 `json:"VpgId,omitempty" name:"VpgId"`
	// VPG唯一ID。如dcg-pl9mmvu1

	UniqVpgId *string `json:"UniqVpgId,omitempty" name:"UniqVpgId"`
	// AppId

	Owner *uint64 `json:"Owner,omitempty" name:"Owner"`
	// 专线网关Ip

	VpgIp *string `json:"VpgIp,omitempty" name:"VpgIp"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 专线网关名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// vpc掩码。如255.255.0.0

	VpcMask *string `json:"VpcMask,omitempty" name:"VpcMask"`
}

type DescribeControllersRequest struct {
	*tchttp.BaseRequest

	// 节点id列表

	ControllerIds []*string `json:"ControllerIds,omitempty" name:"ControllerIds"`
	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 集群id

	ControllerSetId *string `json:"ControllerSetId,omitempty" name:"ControllerSetId"`
	// 物理ip

	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`
	// 归属VPC

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 单次条目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 升序or降序

	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeControllersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeControllersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LimitRule struct {
	// 规则唯一键

	Id *string `json:"Id,omitempty" name:"Id"`
	// 控制器集群唯一键

	ControllerSetId *string `json:"ControllerSetId,omitempty" name:"ControllerSetId"`
	// 规则类型

	ProductType *string `json:"ProductType,omitempty" name:"ProductType"`
	// 规则str唯一键值

	StrKey *string `json:"StrKey,omitempty" name:"StrKey"`
	// 规则int唯一键值

	IntKey []*uint64 `json:"IntKey,omitempty" name:"IntKey"`
	// zk上的区域值

	ZkZone *uint64 `json:"ZkZone,omitempty" name:"ZkZone"`
	// 限速范围

	LimitScope []*string `json:"LimitScope,omitempty" name:"LimitScope"`
	// 限速细项

	LimitItems []*LimitItem `json:"LimitItems,omitempty" name:"LimitItems"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// debug开关

	Debug *uint64 `json:"Debug,omitempty" name:"Debug"`
	// 新建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 限速vip范围

	LimitVipScope []*string `json:"LimitVipScope,omitempty" name:"LimitVipScope"`
}

type NatInfo struct {
	// WanInLimit

	WanInLimit *string `json:"WanInLimit,omitempty" name:"WanInLimit"`
	// VpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// UniqVpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// Connections

	Connections *string `json:"Connections,omitempty" name:"Connections"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// UniqNatId

	UniqNatId *string `json:"UniqNatId,omitempty" name:"UniqNatId"`
	// NatId

	NatId *string `json:"NatId,omitempty" name:"NatId"`
	// Zone

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// ZoneId

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// State

	State *string `json:"State,omitempty" name:"State"`
	// BDeleted

	BDeleted *string `json:"BDeleted,omitempty" name:"BDeleted"`
	// OwedWanOutLimit

	OwedWanOutLimit *string `json:"OwedWanOutLimit,omitempty" name:"OwedWanOutLimit"`
	// GwId

	GwId *string `json:"GwId,omitempty" name:"GwId"`
	// HealthStatus

	HealthStatus *string `json:"HealthStatus,omitempty" name:"HealthStatus"`
	// DetectGwIp

	DetectGwIp *string `json:"DetectGwIp,omitempty" name:"DetectGwIp"`
	// GwIp

	GwIp *string `json:"GwIp,omitempty" name:"GwIp"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateConnectionsTime

	LastUpdateConnectionsTime *string `json:"LastUpdateConnectionsTime,omitempty" name:"LastUpdateConnectionsTime"`
	// OwedFlag

	OwedFlag *string `json:"OwedFlag,omitempty" name:"OwedFlag"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// LastMaxConnections

	LastMaxConnections *string `json:"LastMaxConnections,omitempty" name:"LastMaxConnections"`
	// WanOutLimit

	WanOutLimit *string `json:"WanOutLimit,omitempty" name:"WanOutLimit"`
	// Ip

	Ip []*NatIpInfo `json:"Ip,omitempty" name:"Ip"`
}

type VpcInfo struct {
	// VpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// UniqVpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// VpcName

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// CidrBlock

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// RtbNum

	RtbNum *string `json:"RtbNum,omitempty" name:"RtbNum"`
	// SubnetNum

	SubnetNum *string `json:"SubnetNum,omitempty" name:"SubnetNum"`
	// VpcPeerNum

	VpcPeerNum *string `json:"VpcPeerNum,omitempty" name:"VpcPeerNum"`
	// VpgNum

	VpgNum *string `json:"VpgNum,omitempty" name:"VpgNum"`
	// VpngwNum

	VpngwNum *string `json:"VpngwNum,omitempty" name:"VpngwNum"`
	// VmNum

	VmNum *string `json:"VmNum,omitempty" name:"VmNum"`
	// NatNum

	NatNum *string `json:"NatNum,omitempty" name:"NatNum"`
	// AclNum

	AclNum *string `json:"AclNum,omitempty" name:"AclNum"`
}

type DeleteControllerSetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteControllerSetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteControllerSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWhiteListsRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表

	WhiteListIds []*string `json:"WhiteListIds,omitempty" name:"WhiteListIds"`
}

func (r *DeleteWhiteListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWhiteListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGatewaySetsRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表

	GatewaySetIds []*string `json:"GatewaySetIds,omitempty" name:"GatewaySetIds"`
	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *UpdateGatewaySetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateGatewaySetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnetSubMask struct {
	// 16

	IntMask *uint64 `json:"IntMask,omitempty" name:"IntMask"`
	// 255.255.0.0

	Mask *string `json:"Mask,omitempty" name:"Mask"`
	// 10.0.0.0

	Subnet *string `json:"Subnet,omitempty" name:"Subnet"`
}

type VpnGateway struct {
	// InternalHotHaIp

	InternalHotHaIp *string `json:"InternalHotHaIp,omitempty" name:"InternalHotHaIp"`
	// Zone

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// State

	State *uint64 `json:"State,omitempty" name:"State"`
	// SslvpnPort

	SslvpnPort *uint64 `json:"SslvpnPort,omitempty" name:"SslvpnPort"`
	// BBlocked

	BBlocked *uint64 `json:"BBlocked,omitempty" name:"BBlocked"`
	// VpcId

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// BAutoRenewals

	BAutoRenewals *uint64 `json:"BAutoRenewals,omitempty" name:"BAutoRenewals"`
	// InternalHotHaMask

	InternalHotHaMask *string `json:"InternalHotHaMask,omitempty" name:"InternalHotHaMask"`
	// Quota

	Quota *string `json:"Quota,omitempty" name:"Quota"`
	// ImageId

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// UniqVpngwId

	UniqVpngwId *string `json:"UniqVpngwId,omitempty" name:"UniqVpngwId"`
	// HotHaSubnet

	HotHaSubnet *string `json:"HotHaSubnet,omitempty" name:"HotHaSubnet"`
	// SlaLocalIp

	SlaLocalIp *string `json:"SlaLocalIp,omitempty" name:"SlaLocalIp"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// BNewAfc

	BNewAfc *uint64 `json:"BNewAfc,omitempty" name:"BNewAfc"`
	// BHotHa

	BHotHa *uint64 `json:"BHotHa,omitempty" name:"BHotHa"`
	// ZoneId

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// WanIp

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// HotHaIp

	HotHaIp *string `json:"HotHaIp,omitempty" name:"HotHaIp"`
	// InternalHotHaSubnet

	InternalHotHaSubnet *string `json:"InternalHotHaSubnet,omitempty" name:"InternalHotHaSubnet"`
	// UniqVpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// VpnType

	VpnType *uint64 `json:"VpnType,omitempty" name:"VpnType"`
	// HotHaMac

	HotHaMac *string `json:"HotHaMac,omitempty" name:"HotHaMac"`
	// DealId

	DealId *string `json:"DealId,omitempty" name:"DealId"`
	// HotHaMask

	HotHaMask *string `json:"HotHaMask,omitempty" name:"HotHaMask"`
	// InternalHotHaMac

	InternalHotHaMac *string `json:"InternalHotHaMac,omitempty" name:"InternalHotHaMac"`
	// VpngwId

	VpngwId *uint64 `json:"VpngwId,omitempty" name:"VpngwId"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type Filter struct {
	// 字段名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type InOutBoundInfo struct {
	// Proto

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// Ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// Action

	Action *string `json:"Action,omitempty" name:"Action"`
	// Port

	Port *string `json:"Port,omitempty" name:"Port"`
	// Desc

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type SgInfo struct {
	// usg id

	UsgId *string `json:"UsgId,omitempty" name:"UsgId"`
	// name

	Name *string `json:"Name,omitempty" name:"Name"`
	// owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// counts

	Counts []*SgVmCount `json:"Counts,omitempty" name:"Counts"`
}

type UNatGwWanInfo struct {
	// CarrierId

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
	// UnatgwId

	UnatgwId *uint64 `json:"UnatgwId,omitempty" name:"UnatgwId"`
	// UniqUnatgwId

	UniqUnatgwId *string `json:"UniqUnatgwId,omitempty" name:"UniqUnatgwId"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
}

type RtbInfo struct {
	// VPC实例ID。如65967

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// VPC唯一ID。如vpc-4cukfjyf

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// VPC名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// appid

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// nickname

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// vpc的cidr，只能为10.0.0.0/16，172.16.0.0/16，192.168.0.0/16这三个内网网段内。

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 路由表实例ID

	RtbId *uint64 `json:"RtbId,omitempty" name:"RtbId"`
	// 路由表名称。

	RtbName *string `json:"RtbName,omitempty" name:"RtbName"`
	// 路由表唯一ID。如rtb-65r1ypuo

	UniqRtbId *string `json:"UniqRtbId,omitempty" name:"UniqRtbId"`
	// 路由表类型

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 子网个数

	SubnetNum *uint64 `json:"SubnetNum,omitempty" name:"SubnetNum"`
	// 路由策略信息数组

	Rt []*Route `json:"Rt,omitempty" name:"Rt"`
}

type UpdateDevicesRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表

	DeviceIds []*string `json:"DeviceIds,omitempty" name:"DeviceIds"`
	// 限速开关

	LimitSwitch *uint64 `json:"LimitSwitch,omitempty" name:"LimitSwitch"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *UpdateDevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZkRecoveryRequest struct {
	*tchttp.BaseRequest

	// 网关集群实例ID列表

	SetIds []*string `json:"SetIds,omitempty" name:"SetIds"`
	// 控制器集群实例ID列表

	ControllerSetIds []*string `json:"ControllerSetIds,omitempty" name:"ControllerSetIds"`
	// 规则实例ID列表

	LimitRuleIds []*string `json:"LimitRuleIds,omitempty" name:"LimitRuleIds"`
	// 是否同步执行恢复操作

	Execute *uint64 `json:"Execute,omitempty" name:"Execute"`
	// zk数路径和值对应列表

	PathList []*ZkPathInfo `json:"PathList,omitempty" name:"PathList"`
	// 产品类型

	ProductType *string `json:"ProductType,omitempty" name:"ProductType"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 单页限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ZkRecoveryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ZkRecoveryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkInterface struct {
	// AppId

	Owner *uint64 `json:"Owner,omitempty" name:"Owner"`
	// VPC实例ID。如66379

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// 弹性网卡实例ID。如ins-awf82u1r

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 弹性网卡唯一ID。如eni-agfsccmj

	UniqEniId *string `json:"UniqEniId,omitempty" name:"UniqEniId"`
	// 网卡设备名称。如veth_64DF2Ce8

	IfnName *string `json:"IfnName,omitempty" name:"IfnName"`
	// 外网IP

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// 宿主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 子机IP

	Pip *string `json:"Pip,omitempty" name:"Pip"`
	// 网卡类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// Description

	Description *string `json:"Description,omitempty" name:"Description"`
	// PipWanipDesc

	PipWanipDesc []*string `json:"PipWanipDesc,omitempty" name:"PipWanipDesc"`
	// UniqSubnetId

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
	// Mac

	Mac *string `json:"Mac,omitempty" name:"Mac"`
}

type UNatGwRuleWip struct {
	// WanIp

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// UnatgwRuleId

	UnatgwRuleId *uint64 `json:"UnatgwRuleId,omitempty" name:"UnatgwRuleId"`
	// CarrierId

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
	// UniqUnatgwRuleId

	UniqUnatgwRuleId *string `json:"UniqUnatgwRuleId,omitempty" name:"UniqUnatgwRuleId"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type UNatGwWanInfoAndWanIp struct {
	// CarrierId

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// UnatgwId

	UnatgwId *uint64 `json:"UnatgwId,omitempty" name:"UnatgwId"`
	// UnatgwInfo

	UnatgwInfo *UnatgwInfo `json:"UnatgwInfo,omitempty" name:"UnatgwInfo"`
	// UniqUnatgwId

	UniqUnatgwId *string `json:"UniqUnatgwId,omitempty" name:"UniqUnatgwId"`
	// WanIp

	WanIp []*string `json:"WanIp,omitempty" name:"WanIp"`
}

type Vm struct {
	// AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// VPC唯一ID。如uniqVpcId-45678

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// VPC实例ID。如66379

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// 子机内网IP

	VmIp *string `json:"VmIp,omitempty" name:"VmIp"`
	// vpc子网

	Subnet *string `json:"Subnet,omitempty" name:"Subnet"`
	// 掩码

	Mask *string `json:"Mask,omitempty" name:"Mask"`
	// 宿主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// Mac

	Mac *string `json:"Mac,omitempty" name:"Mac"`
	// GreTunnelName

	GreTunnelName *string `json:"GreTunnelName,omitempty" name:"GreTunnelName"`
	// SlaveHostIp

	SlaveHostIp *string `json:"SlaveHostIp,omitempty" name:"SlaveHostIp"`
	// VmType

	VmType *uint64 `json:"VmType,omitempty" name:"VmType"`
	// GroupId

	GroupId []*uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// WanIp

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// EvpnUnderlayIp

	EvpnUnderlayIp *string `json:"EvpnUnderlayIp,omitempty" name:"EvpnUnderlayIp"`
	// BrName

	BrName *string `json:"BrName,omitempty" name:"BrName"`
	// VpcGwIp

	VpcGwIp *string `json:"VpcGwIp,omitempty" name:"VpcGwIp"`
	// PeerIp

	PeerIp *string `json:"PeerIp,omitempty" name:"PeerIp"`
	// GwIp

	GwIp *string `json:"GwIp,omitempty" name:"GwIp"`
	// TsvIp

	TsvIp *string `json:"TsvIp,omitempty" name:"TsvIp"`
	// ModuleId

	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`
	// DhcpFlag

	DhcpFlag *uint64 `json:"DhcpFlag,omitempty" name:"DhcpFlag"`
	// ForwardHostIp

	ForwardHostIp *string `json:"ForwardHostIp,omitempty" name:"ForwardHostIp"`
	// BMonitorVm

	BMonitorVm *int64 `json:"BMonitorVm,omitempty" name:"BMonitorVm"`
	// IfnName

	IfnName *string `json:"IfnName,omitempty" name:"IfnName"`
	// SetId

	SetId *uint64 `json:"SetId,omitempty" name:"SetId"`
	// WanInLimit

	WanInLimit *float64 `json:"WanInLimit,omitempty" name:"WanInLimit"`
	// WanOutLimit

	WanOutLimit *float64 `json:"WanOutLimit,omitempty" name:"WanOutLimit"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type WhiteList struct {
	// 唯一键值

	Id *string `json:"Id,omitempty" name:"Id"`
	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 匹配范围定义

	MatchField []*MatchField `json:"MatchField,omitempty" name:"MatchField"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 新建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type CreateDeviceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		InstanceSet []*Device `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 实例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDeviceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GatewayGroup struct {
	// 网关集群组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 网关类型(0-pvgw 1-vpcgw 2-jnsgw 3-natgw 4-mcgw 5-dcgw 6-pcgw 7-xgw 8-sxgw 9-underly-natgw)

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 本组内网关概要信息列表

	Vip []*GatewayBrief `json:"Vip,omitempty" name:"Vip"`
	// 网关集群组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 权重

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type GatewayGroupBrief struct {
	// 网关名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 权重

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
	// vpcgw的唯一ID。如vpcgw-cihvb3j1

	UniqVpcgwId *string `json:"UniqVpcgwId,omitempty" name:"UniqVpcgwId"`
	// vpcgw的实例ID。如7

	VpcgwId *uint64 `json:"VpcgwId,omitempty" name:"VpcgwId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type Service struct {
	// Proto

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// UniqClassicVpcgwId

	UniqClassicVpcgwId *string `json:"UniqClassicVpcgwId,omitempty" name:"UniqClassicVpcgwId"`
	// SetId

	SetId *uint64 `json:"SetId,omitempty" name:"SetId"`
	// Type

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// SubnetRouteAlgorithm

	SubnetRouteAlgorithm *uint64 `json:"SubnetRouteAlgorithm,omitempty" name:"SubnetRouteAlgorithm"`
	// VpcId

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// VipSetId

	VipSetId *uint64 `json:"VipSetId,omitempty" name:"VipSetId"`
	// ToaHostIpFlag

	ToaHostIpFlag *uint64 `json:"ToaHostIpFlag,omitempty" name:"ToaHostIpFlag"`
	// LbType

	LbType *uint64 `json:"LbType,omitempty" name:"LbType"`
	// VpcName

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// StickyFlag

	StickyFlag *uint64 `json:"StickyFlag,omitempty" name:"StickyFlag"`
	// UniqVpcGwIpVpcgwId

	UniqVpcGwIpVpcgwId *string `json:"UniqVpcGwIpVpcgwId,omitempty" name:"UniqVpcGwIpVpcgwId"`
	// BypassFlag

	BypassFlag *uint64 `json:"BypassFlag,omitempty" name:"BypassFlag"`
	// RstFlag

	RstFlag *uint64 `json:"RstFlag,omitempty" name:"RstFlag"`
	// RemoteVpcId

	RemoteVpcId *uint64 `json:"RemoteVpcId,omitempty" name:"RemoteVpcId"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// HealthFlag

	HealthFlag *uint64 `json:"HealthFlag,omitempty" name:"HealthFlag"`
	// StickyTimeout

	StickyTimeout *uint64 `json:"StickyTimeout,omitempty" name:"StickyTimeout"`
	// UniqVpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// VpcGwIp

	VpcGwIp *string `json:"VpcGwIp,omitempty" name:"VpcGwIp"`
	// UniqVpcgwId

	UniqVpcgwId *string `json:"UniqVpcgwId,omitempty" name:"UniqVpcgwId"`
	// NoSnatFlag

	NoSnatFlag *uint64 `json:"NoSnatFlag,omitempty" name:"NoSnatFlag"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// ServiceId

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// StickyMaxCount

	StickyMaxCount *uint64 `json:"StickyMaxCount,omitempty" name:"StickyMaxCount"`
	// Vport

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// Route

	Route []*ServiceRoute `json:"Route,omitempty" name:"Route"`
	// GroupId

	GroupId []*uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

type SgVmInfo struct {
	// network interface id

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	// instance id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// alias

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// vpc id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc name

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// eni type

	EniType *string `json:"EniType,omitempty" name:"EniType"`
	// eni name

	EniName *string `json:"EniName,omitempty" name:"EniName"`
}

type SubNetInfo struct {
	// VpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// ZoneId

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// UniqVpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// VpcName

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// SubnetId

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// UniqSubnetId

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
	// SubnetName

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// RtbId

	RtbId *uint64 `json:"RtbId,omitempty" name:"RtbId"`
	// RtbName

	RtbName *string `json:"RtbName,omitempty" name:"RtbName"`
	// UniqRtbId

	UniqRtbId *string `json:"UniqRtbId,omitempty" name:"UniqRtbId"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// CidrBlock

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// AvailableIpNum

	AvailableIpNum *string `json:"AvailableIpNum,omitempty" name:"AvailableIpNum"`
	// UsedIpNum

	UsedIpNum *string `json:"UsedIpNum,omitempty" name:"UsedIpNum"`
	// TotalIpNum

	TotalIpNum *string `json:"TotalIpNum,omitempty" name:"TotalIpNum"`
	// Mask

	Mask *string `json:"Mask,omitempty" name:"Mask"`
}

type VpcLimitNamesToType struct {
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Types

	Types []*uint64 `json:"Types,omitempty" name:"Types"`
}

type CreateControllerRequest struct {
	*tchttp.BaseRequest

	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 节点物理ip

	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`
	// 关联控制器集群

	ControllerSetId *string `json:"ControllerSetId,omitempty" name:"ControllerSetId"`
	// 若节点在overlay网络需要传入正确vpcid，默认值为-1

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 可用区信息

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *CreateControllerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateControllerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		InstanceSet []*WhiteList `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 实例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteControllersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteControllersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteControllersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetectPkgVersion struct {
	// Weight

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
	// Ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// Gwname

	Gwname *string `json:"Gwname,omitempty" name:"Gwname"`
	// Version

	Version *string `json:"Version,omitempty" name:"Version"`
	// AzName

	AzName *string `json:"AzName,omitempty" name:"AzName"`
	// Groupid

	Groupid *string `json:"Groupid,omitempty" name:"Groupid"`
}

type VmIpv6 struct {
	// Description

	Description *string `json:"Description,omitempty" name:"Description"`
	// BPublicIp

	BPublicIp *string `json:"BPublicIp,omitempty" name:"BPublicIp"`
	// State

	State *string `json:"State,omitempty" name:"State"`
	// WanIp

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// BHavip

	BHavip *string `json:"BHavip,omitempty" name:"BHavip"`
	// Pip

	Pip *string `json:"Pip,omitempty" name:"Pip"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// BBlocked

	BBlocked *string `json:"BBlocked,omitempty" name:"BBlocked"`
}

type CreateDeviceRequest struct {
	*tchttp.BaseRequest

	// 物理机系列号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 节点网口地址

	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`
	// 集群唯一键

	SetId *string `json:"SetId,omitempty" name:"SetId"`
	// VPCID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 限速开关

	LimitSwitch *bool `json:"LimitSwitch,omitempty" name:"LimitSwitch"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *CreateDeviceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeviceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteControllersRequest struct {
	*tchttp.BaseRequest

	// 实例id列表

	ControllerIds []*string `json:"ControllerIds,omitempty" name:"ControllerIds"`
}

func (r *DeleteControllersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteControllersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewaySetsRequest struct {
	*tchttp.BaseRequest

	// 集群ID列表

	GatewaySetIds []*string `json:"GatewaySetIds,omitempty" name:"GatewaySetIds"`
	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 集群产品类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 集群模式

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 集群vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 单次条目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 升序or降序

	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeGatewaySetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewaySetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubnetDel struct {
	// AppId

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// VPC实例ID。如66379

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// 子网实例ID。如7623

	SubnetId *uint64 `json:"SubnetId,omitempty" name:"SubnetId"`
}

type UNatGwRuleRip struct {
	// UnatgwRuleId

	UnatgwRuleId *uint64 `json:"UnatgwRuleId,omitempty" name:"UnatgwRuleId"`
	// UniqUnatgwRuleId

	UniqUnatgwRuleId *string `json:"UniqUnatgwRuleId,omitempty" name:"UniqUnatgwRuleId"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Rip

	Rip *string `json:"Rip,omitempty" name:"Rip"`
	// CarrierId

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
}

type VpnQuota struct {
	// HaFlag

	HaFlag *uint64 `json:"HaFlag,omitempty" name:"HaFlag"`
	// UniqVpngwQuotaId

	UniqVpngwQuotaId *string `json:"UniqVpngwQuotaId,omitempty" name:"UniqVpngwQuotaId"`
	// DataDiskType

	DataDiskType *uint64 `json:"DataDiskType,omitempty" name:"DataDiskType"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Cname

	Cname *string `json:"Cname,omitempty" name:"Cname"`
	// Cpu

	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
	// RootDisk

	RootDisk *uint64 `json:"RootDisk,omitempty" name:"RootDisk"`
	// RootDiskType

	RootDiskType *uint64 `json:"RootDiskType,omitempty" name:"RootDiskType"`
	// Mem

	Mem *uint64 `json:"Mem,omitempty" name:"Mem"`
	// Bandwidth

	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// DataDisk

	DataDisk *uint64 `json:"DataDisk,omitempty" name:"DataDisk"`
}

type DeleteWhiteListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWhiteListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWhiteListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateControllerSetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateControllerSetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateControllerSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MatchField struct {
	// 匹配项键值，如：AppId

	Key *string `json:"Key,omitempty" name:"Key"`
	// 匹配值列表

	Value []*string `json:"Value,omitempty" name:"Value"`
}

type Rs struct {
	// rs ip

	RsIp *string `json:"RsIp,omitempty" name:"RsIp"`
	// rs port

	RsPort *int64 `json:"RsPort,omitempty" name:"RsPort"`
	// rs 权重

	RsWeight *int64 `json:"RsWeight,omitempty" name:"RsWeight"`
	// rs vpc id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type Vsr struct {
	// UniqVsrId

	UniqVsrId *string `json:"UniqVsrId,omitempty" name:"UniqVsrId"`
	// UniqVpngwId

	UniqVpngwId *string `json:"UniqVpngwId,omitempty" name:"UniqVpngwId"`
	// Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// UniqVpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// Ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// IfnName

	IfnName *string `json:"IfnName,omitempty" name:"IfnName"`
	// Zone

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// VpcId

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// VpngwId

	VpngwId *uint64 `json:"VpngwId,omitempty" name:"VpngwId"`
	// HostIp

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// HotHaFlag

	HotHaFlag *uint64 `json:"HotHaFlag,omitempty" name:"HotHaFlag"`
	// Type

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// AdminIp

	AdminIp *string `json:"AdminIp,omitempty" name:"AdminIp"`
	// ZoneId

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeAlgorithmsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		InstanceSet []*Algorithm `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 实例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlgorithmsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlgorithmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		InstanceSet []*Device `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 实例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BasicSubnetInfo struct {
	// vpc0的cidr

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 可用`IP`数。

	Available *uint64 `json:"Available,omitempty" name:"Available"`
	// `IP`总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
	// 已使用`IP`总数

	Used *uint64 `json:"Used,omitempty" name:"Used"`
	// vpc0网段名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type PipWanDesc struct {
	// Pip

	Pip *string `json:"Pip,omitempty" name:"Pip"`
	// WanIp

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// Description

	Description *string `json:"Description,omitempty" name:"Description"`
}

type SessionPersistenceStrategyInfo struct {
	// cookie type

	CookieKey *string `json:"CookieKey,omitempty" name:"CookieKey"`
	// timeout

	Timeout *string `json:"Timeout,omitempty" name:"Timeout"`
	// type

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DeleteDevicesRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表

	DeviceIds []*string `json:"DeviceIds,omitempty" name:"DeviceIds"`
}

func (r *DeleteDevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HealthCheckInfo struct {
	// ActiveHealthCheck

	ActiveHealthCheck *ActiveHealthCheckInfo `json:"ActiveHealthCheck,omitempty" name:"ActiveHealthCheck"`
	// PassiveHealthCheck

	PassiveHealthCheck *PassiveHealthCheckInfo `json:"PassiveHealthCheck,omitempty" name:"PassiveHealthCheck"`
	// type

	Type *string `json:"Type,omitempty" name:"Type"`
}

type SubnetBaseInfo struct {
	// 子网，如10.23.0.0

	Subnet *string `json:"Subnet,omitempty" name:"Subnet"`
	// VPC实例ID。如65967

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// 子网名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// VPC唯一ID。如vpc-4cukfjyf

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// vpc网络类型（1：cvm；2：黑石1.0；3：黑石2.0）

	VpcType *uint64 `json:"VpcType,omitempty" name:"VpcType"`
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 掩码长度

	IntMask *uint64 `json:"IntMask,omitempty" name:"IntMask"`
	// 子网唯一ID。如subnet-jdbjmn4i

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
	// appid

	Owner *uint64 `json:"Owner,omitempty" name:"Owner"`
	// 子网实例ID。如9963

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 路由表类型

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type SubnetOutput struct {
	// 子网名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 接收方VPC唯一ID。如uniqVpcId-45678

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// VPC实例ID。如45678

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// 子网唯一ID。如subnet-f3mw1qf0

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
	// 子网实例ID。如10133

	SubnetId *uint64 `json:"SubnetId,omitempty" name:"SubnetId"`
	// 子网，如172.16.36.0

	Subnet *string `json:"Subnet,omitempty" name:"Subnet"`
	// IntMask，如24

	IntMask *uint64 `json:"IntMask,omitempty" name:"IntMask"`
	// VpcSubnet，如172.16.0.0

	VpcSubnet *string `json:"VpcSubnet,omitempty" name:"VpcSubnet"`
	// VpcIntMask，如16

	VpcIntMask *uint64 `json:"VpcIntMask,omitempty" name:"VpcIntMask"`
	// AppId

	Owner *uint64 `json:"Owner,omitempty" name:"Owner"`
	// vpc网络类型（1：cvm；2：黑石1.0；3：黑石2.0）

	VpcType *int64 `json:"VpcType,omitempty" name:"VpcType"`
	// BVpcDefault

	BVpcDefault *int64 `json:"BVpcDefault,omitempty" name:"BVpcDefault"`
	// 子网的 `IPv4` `CIDR`。

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
	// Type

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type UnatgwRuleRipR struct {
	// CarrierId

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
	// Rip

	Rip *string `json:"Rip,omitempty" name:"Rip"`
	// UnatgwRuleId

	UnatgwRuleId *uint64 `json:"UnatgwRuleId,omitempty" name:"UnatgwRuleId"`
}

type CreateControllerSetRequest struct {
	*tchttp.BaseRequest

	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 部署位置

	WorkSite *string `json:"WorkSite,omitempty" name:"WorkSite"`
	// 是否独占

	IsShared *bool `json:"IsShared,omitempty" name:"IsShared"`
	// 策略ID

	WhiteListId *string `json:"WhiteListId,omitempty" name:"WhiteListId"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *CreateControllerSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateControllerSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateControllerSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		InstanceSet []*ControllerSet `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 实例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateControllerSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateControllerSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GatewayBrief struct {
	// VpcgwId

	VpcgwId *uint64 `json:"VpcgwId,omitempty" name:"VpcgwId"`
	// Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// Rip1

	Rip1 *string `json:"Rip1,omitempty" name:"Rip1"`
	// Rip2

	Rip2 *string `json:"Rip2,omitempty" name:"Rip2"`
}

type Rule struct {
	// app name

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// biz admin

	BizAdmin *string `json:"BizAdmin,omitempty" name:"BizAdmin"`
	// bizid

	BizId *string `json:"BizId,omitempty" name:"BizId"`
	// biz type

	BizType *string `json:"BizType,omitempty" name:"BizType"`
	// 城市

	City *string `json:"City,omitempty" name:"City"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 转发模式

	FwdMode *string `json:"FwdMode,omitempty" name:"FwdMode"`
	// 机房名字

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// 运营商名字

	IspName *string `json:"IspName,omitempty" name:"IspName"`
	// keep time

	KeepTime *string `json:"KeepTime,omitempty" name:"KeepTime"`
	// mtu

	Mtu *string `json:"Mtu,omitempty" name:"Mtu"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 允许quic协议

	QuicEnable *string `json:"QuicEnable,omitempty" name:"QuicEnable"`
	// rule id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 轮询类型

	ScheduleType *string `json:"ScheduleType,omitempty" name:"ScheduleType"`
	// tsv vip

	TsvVip *string `json:"TsvVip,omitempty" name:"TsvVip"`
	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// vport

	Vport *string `json:"Vport,omitempty" name:"Vport"`
	// vpc id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// usage

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// set name

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// rs 数量

	RsNum *string `json:"RsNum,omitempty" name:"RsNum"`
	// rs 列表

	RsSet []*Rs `json:"RsSet,omitempty" name:"RsSet"`
	// alive interval

	AliveInterval *string `json:"AliveInterval,omitempty" name:"AliveInterval"`
	// kick interval

	KickInterval *string `json:"KickInterval,omitempty" name:"KickInterval"`
	// probe interval

	ProbeInterval *string `json:"ProbeInterval,omitempty" name:"ProbeInterval"`
	// probe timeout

	ProbeTimeout *string `json:"ProbeTimeout,omitempty" name:"ProbeTimeout"`
}

type SubnetNameID struct {
	// 子网名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 子网实例ID。如3706

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 子网唯一ID。如subnet-3cxo1o9e

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
}

type VmSgInfoOuter struct {
	// 错误码

	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 错误信息

	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
	// vm 管理的sg信息

	UsgInfo []*VmSgInfo `json:"UsgInfo,omitempty" name:"UsgInfo"`
}

type Controller struct {
	// 唯一ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 集群唯一ID

	ControllerSetId *string `json:"ControllerSetId,omitempty" name:"ControllerSetId"`
	// 物理ip

	LanIp *float64 `json:"LanIp,omitempty" name:"LanIp"`
	// vpcid

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// region信息

	Region *string `json:"Region,omitempty" name:"Region"`
	// zone信息

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 描述符

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 集群命名

	SetName *string `json:"SetName,omitempty" name:"SetName"`
}

type ForwardStrategyInfo struct {
	// detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// type

	Type *string `json:"Type,omitempty" name:"Type"`
}

type RtbRoute struct {
	// 目的端cidr列表

	DestinationCidrBlockSet []*string `json:"DestinationCidrBlockSet,omitempty" name:"DestinationCidrBlockSet"`
	// 路由表实例ID

	RtbId *uint64 `json:"RtbId,omitempty" name:"RtbId"`
	// 路由表唯一ID。如rtb-65r1ypuo

	UniqRtbId *uint64 `json:"UniqRtbId,omitempty" name:"UniqRtbId"`
	// 路由表名称

	RtbName *string `json:"RtbName,omitempty" name:"RtbName"`
	// 子网名称ID列表

	SubnetSet []*SubnetNameID `json:"SubnetSet,omitempty" name:"SubnetSet"`
}

type UNatGwRuleWanIp struct {
	// CarrierId

	CarrierId *string `json:"CarrierId,omitempty" name:"CarrierId"`
	// Ips

	Ips []*string `json:"Ips,omitempty" name:"Ips"`
}

type LocationInfo struct {
	// CustomizedConf

	CustomizedConf *string `json:"CustomizedConf,omitempty" name:"CustomizedConf"`
	// ForwardStrategy

	ForwardStrategy *ForwardStrategyInfo `json:"ForwardStrategy,omitempty" name:"ForwardStrategy"`
	// HealthCheck

	HealthCheck *HealthCheckInfo `json:"HealthCheck,omitempty" name:"HealthCheck"`
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// KeepaliveConnNum

	KeepaliveConnNum *string `json:"KeepaliveConnNum,omitempty" name:"KeepaliveConnNum"`
	// KeepaliveEnable

	KeepaliveEnable *string `json:"KeepaliveEnable,omitempty" name:"KeepaliveEnable"`
	// RsSet

	RsSet []*Rs `json:"RsSet,omitempty" name:"RsSet"`
	// SessionPersistenceStrategy

	SessionPersistenceStrategy *SessionPersistenceStrategyInfo `json:"SessionPersistenceStrategy,omitempty" name:"SessionPersistenceStrategy"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// Url

	Url *string `json:"Url,omitempty" name:"Url"`
}

type SubnetInput struct {
	// AppId

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// VPC实例ID。如66379

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// 子网的 `IPv4` `CIDR`。

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type UnatgwInfo struct {
	// UnatgwId

	UnatgwId *uint64 `json:"UnatgwId,omitempty" name:"UnatgwId"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// GwIp

	GwIp *string `json:"GwIp,omitempty" name:"GwIp"`
	// UniqUnatgwId

	UniqUnatgwId *string `json:"UniqUnatgwId,omitempty" name:"UniqUnatgwId"`
	// IdcName

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Rip

	Rip *string `json:"Rip,omitempty" name:"Rip"`
	// ZoneId

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type VpnConn struct {
	// SlaRemoteIp

	SlaRemoteIp *string `json:"SlaRemoteIp,omitempty" name:"SlaRemoteIp"`
	// UsrgwId

	UsrgwId *uint64 `json:"UsrgwId,omitempty" name:"UsrgwId"`
	// LocalAddress

	LocalAddress *string `json:"LocalAddress,omitempty" name:"LocalAddress"`
	// PropoAuthenAlgorithm

	PropoAuthenAlgorithm *string `json:"PropoAuthenAlgorithm,omitempty" name:"PropoAuthenAlgorithm"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// IntegrityAlgorithm

	IntegrityAlgorithm *string `json:"IntegrityAlgorithm,omitempty" name:"IntegrityAlgorithm"`
	// PreSharedKey

	PreSharedKey *string `json:"PreSharedKey,omitempty" name:"PreSharedKey"`
	// SecurityProto

	SecurityProto *string `json:"SecurityProto,omitempty" name:"SecurityProto"`
	// EncryptAlgorithm

	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" name:"EncryptAlgorithm"`
	// EncryptProto

	EncryptProto *string `json:"EncryptProto,omitempty" name:"EncryptProto"`
	// RemoteIdentity

	RemoteIdentity *string `json:"RemoteIdentity,omitempty" name:"RemoteIdentity"`
	// BEnableSla

	BEnableSla *uint64 `json:"BEnableSla,omitempty" name:"BEnableSla"`
	// ExchangeMode

	ExchangeMode *string `json:"ExchangeMode,omitempty" name:"ExchangeMode"`
	// AfcTaskId

	AfcTaskId *string `json:"AfcTaskId,omitempty" name:"AfcTaskId"`
	// IkeSaLifetimeSeconds

	IkeSaLifetimeSeconds *string `json:"IkeSaLifetimeSeconds,omitempty" name:"IkeSaLifetimeSeconds"`
	// EncapMode

	EncapMode *string `json:"EncapMode,omitempty" name:"EncapMode"`
	// UniqUsrgwId

	UniqUsrgwId *string `json:"UniqUsrgwId,omitempty" name:"UniqUsrgwId"`
	// PfsDhGroup

	PfsDhGroup *string `json:"PfsDhGroup,omitempty" name:"PfsDhGroup"`
	// IKEVersion

	IKEVersion *string `json:"IKEVersion,omitempty" name:"IKEVersion"`
	// VpnProto

	VpnProto *string `json:"VpnProto,omitempty" name:"VpnProto"`
	// VpcId

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// NatType

	NatType *uint64 `json:"NatType,omitempty" name:"NatType"`
	// UniqVpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// PropoEncryAlgorithm

	PropoEncryAlgorithm *string `json:"PropoEncryAlgorithm,omitempty" name:"PropoEncryAlgorithm"`
	// LocalIdentity

	LocalIdentity *string `json:"LocalIdentity,omitempty" name:"LocalIdentity"`
	// RemoteAddress

	RemoteAddress *string `json:"RemoteAddress,omitempty" name:"RemoteAddress"`
	// State

	State *uint64 `json:"State,omitempty" name:"State"`
	// IpsecSaLifetimeTraffic

	IpsecSaLifetimeTraffic *string `json:"IpsecSaLifetimeTraffic,omitempty" name:"IpsecSaLifetimeTraffic"`
	// UniqVpnconnId

	UniqVpnconnId *string `json:"UniqVpnconnId,omitempty" name:"UniqVpnconnId"`
	// VpnconnId

	VpnconnId *uint64 `json:"VpnconnId,omitempty" name:"VpnconnId"`
	// SlaInterval

	SlaInterval *uint64 `json:"SlaInterval,omitempty" name:"SlaInterval"`
	// VpngwId

	VpngwId *uint64 `json:"VpngwId,omitempty" name:"VpngwId"`
	// PropoAuthenMethod

	PropoAuthenMethod *string `json:"PropoAuthenMethod,omitempty" name:"PropoAuthenMethod"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// UniqVpngwId

	UniqVpngwId *string `json:"UniqVpngwId,omitempty" name:"UniqVpngwId"`
	// IpsecSaLifetimeSeconds

	IpsecSaLifetimeSeconds *string `json:"IpsecSaLifetimeSeconds,omitempty" name:"IpsecSaLifetimeSeconds"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// RouteType

	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`
	// DhGroupName

	DhGroupName *string `json:"DhGroupName,omitempty" name:"DhGroupName"`
	// Acl

	Acl *string `json:"Acl,omitempty" name:"Acl"`
	// Snet

	Snet []*SnetSubMask `json:"Snet,omitempty" name:"Snet"`
}

type DescribeControllersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		InstanceSet []*Controller `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 实例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeControllersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeControllersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLimitRuleRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	LimitRuleId *string `json:"LimitRuleId,omitempty" name:"LimitRuleId"`
	// 限速限速

	LimitItems []*LimitItem `json:"LimitItems,omitempty" name:"LimitItems"`
	// 调试开关

	Debug *uint64 `json:"Debug,omitempty" name:"Debug"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *UpdateLimitRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLimitRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZkRecoveryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ZkRecoveryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ZkRecoveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LbLimit struct {
	// app id

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 配额类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 当前配额

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 最大配额

	MaxLimit *int64 `json:"MaxLimit,omitempty" name:"MaxLimit"`
}

type PassiveHealthCheckInfo struct {
	// 超时

	FailTimeout *string `json:"FailTimeout,omitempty" name:"FailTimeout"`
	// 最大失败次数

	MaxFails *string `json:"MaxFails,omitempty" name:"MaxFails"`
}

type CreateLimitRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		InstanceSet []*LimitRule `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 实例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLimitRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLimitRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeprecateUNatGwWanInfoAndWanIp struct {
	// CarrierId

	CarrierId []*uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// UnatgwId

	UnatgwId *uint64 `json:"UnatgwId,omitempty" name:"UnatgwId"`
	// UnatgwInfo

	UnatgwInfo *UnatgwInfo `json:"UnatgwInfo,omitempty" name:"UnatgwInfo"`
	// UniqUnatgwId

	UniqUnatgwId *string `json:"UniqUnatgwId,omitempty" name:"UniqUnatgwId"`
	// WanIp

	WanIp []*string `json:"WanIp,omitempty" name:"WanIp"`
}

type GatewaySet struct {
	// 唯一键值

	Id *string `json:"Id,omitempty" name:"Id"`
	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 集群类似

	Type *string `json:"Type,omitempty" name:"Type"`
	// 集群模式

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// vip列表

	VipList []*string `json:"VipList,omitempty" name:"VipList"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 限速开关

	LimitSwitch *bool `json:"LimitSwitch,omitempty" name:"LimitSwitch"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 新建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 关联节点数

	NodeCnt *uint64 `json:"NodeCnt,omitempty" name:"NodeCnt"`
	// 关联规则数

	RuleCnt *float64 `json:"RuleCnt,omitempty" name:"RuleCnt"`
}

type ServiceRoute struct {
	// ServiceId

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// VpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// Routes

	Routes []*ServiceRouteRoute `json:"Routes,omitempty" name:"Routes"`
}

type UpdateLimitRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateLimitRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLimitRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatIpInfo struct {
	// Eip

	Eip *string `json:"Eip,omitempty" name:"Eip"`
	// LanIp

	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`
	// EipId

	EipId *string `json:"EipId,omitempty" name:"EipId"`
	// TsvIp

	TsvIp *string `json:"TsvIp,omitempty" name:"TsvIp"`
	// BBlocked

	BBlocked *string `json:"BBlocked,omitempty" name:"BBlocked"`
	// BAutoApply

	BAutoApply *string `json:"BAutoApply,omitempty" name:"BAutoApply"`
}

type VipInfo struct {
	// 城市

	City *string `json:"City,omitempty" name:"City"`
	// 运营商

	IspName *string `json:"IspName,omitempty" name:"IspName"`
	// 地址

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// vpc id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeGatewaySetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		InstanceSet []*GatewaySet `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 实例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewaySetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewaySetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLimitRulesRequest struct {
	*tchttp.BaseRequest

	// 规则ID列表

	LimitRuleIds []*string `json:"LimitRuleIds,omitempty" name:"LimitRuleIds"`
	// 网关集群列表

	LimitScope []*string `json:"LimitScope,omitempty" name:"LimitScope"`
	// 控制器集群id列表

	ControllerSetIds []*string `json:"ControllerSetIds,omitempty" name:"ControllerSetIds"`
	// 产品类型

	ProductType []*string `json:"ProductType,omitempty" name:"ProductType"`
	// 规则string类型key

	StrKey *string `json:"StrKey,omitempty" name:"StrKey"`
	// 调试开关

	Debug []*uint64 `json:"Debug,omitempty" name:"Debug"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 单次条目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 升序or降序

	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeLimitRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLimitRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateControllersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateControllersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateControllersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GatewayInstallState struct {
	// EventId

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// FinishTime

	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
	// Sip

	Sip *string `json:"Sip,omitempty" name:"Sip"`
	// Version

	Version *string `json:"Version,omitempty" name:"Version"`
	// Msg

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// Status

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// VrouterId

	VrouterId *uint64 `json:"VrouterId,omitempty" name:"VrouterId"`
	// Ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// Gwname

	Gwname *string `json:"Gwname,omitempty" name:"Gwname"`
	// Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// SvrIdcName

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type VmSgInfo struct {
	// Sys

	Sys *string `json:"Sys,omitempty" name:"Sys"`
	// UsgId

	UsgId *string `json:"UsgId,omitempty" name:"UsgId"`
	// UsgName

	UsgName *string `json:"UsgName,omitempty" name:"UsgName"`
	// ProjectId

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// Os

	Os *string `json:"Os,omitempty" name:"Os"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// UsgRemark

	UsgRemark *string `json:"UsgRemark,omitempty" name:"UsgRemark"`
}

type ZkPathInfo struct {
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeControllerSetsRequest struct {
	*tchttp.BaseRequest

	// 集群id列表

	ControllerSetIds []*string `json:"ControllerSetIds,omitempty" name:"ControllerSetIds"`
	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 部署位置

	WorkSite *string `json:"WorkSite,omitempty" name:"WorkSite"`
	// 独占标识

	IsShared *bool `json:"IsShared,omitempty" name:"IsShared"`
	// 关联策略ID

	WhiteListId *string `json:"WhiteListId,omitempty" name:"WhiteListId"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 单次条目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 升序or降序

	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeControllerSetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeControllerSetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGatewaySetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGatewaySetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateGatewaySetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AclInfo struct {
	// VpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// AclId

	AclId *string `json:"AclId,omitempty" name:"AclId"`
	// UniqVpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// SubnetNum

	SubnetNum *string `json:"SubnetNum,omitempty" name:"SubnetNum"`
	// VpcSubnet

	VpcSubnet *string `json:"VpcSubnet,omitempty" name:"VpcSubnet"`
	// VpcMask

	VpcMask *string `json:"VpcMask,omitempty" name:"VpcMask"`
	// VpcName

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// VpcBDefault

	VpcBDefault *string `json:"VpcBDefault,omitempty" name:"VpcBDefault"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// UniqAclId

	UniqAclId *string `json:"UniqAclId,omitempty" name:"UniqAclId"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// SubNetSet

	SubNetSet []*AclSubNetInfo `json:"SubNetSet,omitempty" name:"SubNetSet"`
}

type Device struct {
	// 唯一键值

	Id *string `json:"Id,omitempty" name:"Id"`
	// 物理机序列号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 集群唯一键

	SetId *string `json:"SetId,omitempty" name:"SetId"`
	// 集群名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// vip列表

	VipList []*string `json:"VipList,omitempty" name:"VipList"`
	// 物理ip

	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`
	// 限速开关

	LimitSwitch *bool `json:"LimitSwitch,omitempty" name:"LimitSwitch"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type JnsGatewayServicePVm struct {
	// HostIp

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// UniqVgwIndex

	UniqVgwIndex *string `json:"UniqVgwIndex,omitempty" name:"UniqVgwIndex"`
	// VgwIndex

	VgwIndex *string `json:"VgwIndex,omitempty" name:"VgwIndex"`
}

type VpcPeer struct {
	// 接收方VPC实例ID。如45678

	DVpcId *uint64 `json:"DVpcId,omitempty" name:"DVpcId"`
	// 接收方VPC唯一ID。如uniqVpcId-45678

	UniqDVpcId *string `json:"UniqDVpcId,omitempty" name:"UniqDVpcId"`
	// 接收方AppId

	DOwner *string `json:"DOwner,omitempty" name:"DOwner"`
	// 发起方VPC实例ID。如66379

	SVpcId *uint64 `json:"SVpcId,omitempty" name:"SVpcId"`
	// 发起方VPC唯一ID。如vpc-oazo8ujl

	UniqSVpcId *string `json:"UniqSVpcId,omitempty" name:"UniqSVpcId"`
	// 发起方AppId

	SOwner *string `json:"SOwner,omitempty" name:"SOwner"`
	// 带宽，单位Mb

	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 本地带宽，单位Mb

	LocalBandwidth *uint64 `json:"LocalBandwidth,omitempty" name:"LocalBandwidth"`
	// 对等连接状态，0-连接失败, 1-连接成功

	State *int64 `json:"State,omitempty" name:"State"`
	// 对等连接实例ID。如23

	VpcPeerId *uint64 `json:"VpcPeerId,omitempty" name:"VpcPeerId"`
	// 发起方地域(0 广州 1 上海金桥 2 北美 3 香港 4 上海金融 5 北京）

	SRegion *uint64 `json:"SRegion,omitempty" name:"SRegion"`
	// 接收方地域(0 广州 1 上海金桥 2 北美 3 香港 4 上海金融 5 北京）

	DRegion *uint64 `json:"DRegion,omitempty" name:"DRegion"`
	// 是否欠费，0不欠费，1欠费

	OwedFlag *int64 `json:"OwedFlag,omitempty" name:"OwedFlag"`
	// 对等连接名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 0或者1

	BNfv *int64 `json:"BNfv,omitempty" name:"BNfv"`
	// 专线网关ip

	VsgIp *string `json:"VsgIp,omitempty" name:"VsgIp"`
	// 0或者1

	BNewAfc *int64 `json:"BNewAfc,omitempty" name:"BNewAfc"`
	// 发起方过期时间

	SExpireTime *string `json:"SExpireTime,omitempty" name:"SExpireTime"`
	// 接收方过期时间

	DExpireTime *string `json:"DExpireTime,omitempty" name:"DExpireTime"`
}

type CreateControllerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例列表

		InstanceSet []*Controller `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateControllerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateControllerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLimitRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLimitRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLimitRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateLimitRuleRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表

	LimitRuleIds []*string `json:"LimitRuleIds,omitempty" name:"LimitRuleIds"`
	// 控制器实例ID

	ControllerSetId *string `json:"ControllerSetId,omitempty" name:"ControllerSetId"`
	// 规则限速范围信息

	LimitScope []*string `json:"LimitScope,omitempty" name:"LimitScope"`
}

func (r *MigrateLimitRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateLimitRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateControllerSetsRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表

	ControllerSetIds []*string `json:"ControllerSetIds,omitempty" name:"ControllerSetIds"`
	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 集群关联策略ID

	WhiteListId *string `json:"WhiteListId,omitempty" name:"WhiteListId"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *UpdateControllerSetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateControllerSetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWhiteListRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	WhiteListId *string `json:"WhiteListId,omitempty" name:"WhiteListId"`
	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 匹配细项

	MatchField []*MatchField `json:"MatchField,omitempty" name:"MatchField"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *UpdateWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLimitRulesRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表

	LimitRuleIds []*string `json:"LimitRuleIds,omitempty" name:"LimitRuleIds"`
}

func (r *DeleteLimitRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLimitRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VmSet struct {
	// AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// VpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// UniqVpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// VmIp

	VmIp *string `json:"VmIp,omitempty" name:"VmIp"`
	// Subnet

	Subnet *string `json:"Subnet,omitempty" name:"Subnet"`
	// Mask

	Mask *string `json:"Mask,omitempty" name:"Mask"`
	// HostIp

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Ipv6

	Ipv6 []*Ipv6Set `json:"Ipv6,omitempty" name:"Ipv6"`
}
