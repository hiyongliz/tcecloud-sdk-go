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

type DeleteUnderlayNatGatewayWanInfoRequest struct {
	*tchttp.BaseRequest

	// 运营商

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
	// underlay natgateway的ID

	UnatgwId *uint64 `json:"UnatgwId,omitempty" name:"UnatgwId"`
}

func (r *DeleteUnderlayNatGatewayWanInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUnderlayNatGatewayWanInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNfvSwitchGroupRequest struct {
	*tchttp.BaseRequest

	// 网关类型

	GwType *string `json:"GwType,omitempty" name:"GwType"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 内网/外网交换机
	// inner
	// outer

	Type *string `json:"Type,omitempty" name:"Type"`
	// tgw外网口vlanid信息

	WanVlanId *string `json:"WanVlanId,omitempty" name:"WanVlanId"`
	// 子机虚拟化类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

func (r *GetNfvSwitchGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNfvSwitchGroupRequest) FromJsonString(s string) error {
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

type Ipv6Set struct {
	// UniqEniId

	UniqEniId *string `json:"UniqEniId,omitempty" name:"UniqEniId"`
	// Info

	Info []*VmIpv6 `json:"Info,omitempty" name:"Info"`
}

type DescribeProtectPolicyIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网关过载保护配置信息集合。

		ProtectPolicyIpSet []*ProtectPolicyIp `json:"ProtectPolicyIpSet,omitempty" name:"ProtectPolicyIpSet"`
		// 符合条件的总条数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProtectPolicyIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProtectPolicyIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRtbSubnetsRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 条数限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤器列表

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方式

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeRtbSubnetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRtbSubnetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcLimitsRequest struct {
	*tchttp.BaseRequest

	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// Type

	Types []*uint64 `json:"Types,omitempty" name:"Types"`
	// Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeVpcLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 任务状态，0-成功 1-失败 2-运行中

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVmRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方式

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 过滤器列表

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// vpc子网

	Subnet *string `json:"Subnet,omitempty" name:"Subnet"`
	// VPC实例ID。如66379

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// 宿主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 子机IP

	VmIp *string `json:"VmIp,omitempty" name:"VmIp"`
}

func (r *DescribeVmRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVmRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVmIpv6InfoRequest struct {
	*tchttp.BaseRequest

	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序列表

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件，参数不支持同时指定VpcIds和Filters。
	// <li>AppId - String - （过滤条件）用户的AppId</li>
	// <li>UniqVpcId - String - （过滤条件）用户的vpc唯一id。</li>
	// <li>VmIp- String - （过滤条件）cvm ip。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVmIpv6InfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVmIpv6InfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDetectNodeInfoRequest struct {
	*tchttp.BaseRequest

	// GatewayName

	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`
	// HostIp

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// GroupId

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetDetectNodeInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDetectNodeInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductStateInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品状态信息查询接口返回数据

		Data *ProductStateInfoRsp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryProductStateInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductStateInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type DeleteUnderlayNatGatewayRuleRequest struct {
	*tchttp.BaseRequest

	// underlay natgateway规则的ID

	UnatgwRuleId *uint64 `json:"UnatgwRuleId,omitempty" name:"UnatgwRuleId"`
}

func (r *DeleteUnderlayNatGatewayRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUnderlayNatGatewayRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEniListRequest struct {
	*tchttp.BaseRequest

	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序列表

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件，参数不支持同时指定VpcIds和Filters。
	// <li>AppId - String - （过滤条件）用户的AppId</li>
	// <li>UniqVpcId - String - （过滤条件）用户的vpc唯一id。</li>
	// <li>EniId - String - （过滤条件）弹性网卡唯一id。</li>
	// <li>InstanceId - String - （过滤条件）实例id。</li>
	// <li>Pip - String - （过滤条件）pip。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeEniListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEniListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SessionPersistenceStrategyInfo struct {
	// cookie type

	CookieKey *string `json:"CookieKey,omitempty" name:"CookieKey"`
	// timeout

	Timeout *string `json:"Timeout,omitempty" name:"Timeout"`
	// type

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeGatewayTypeNamesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeGatewayTypeNamesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayTypeNamesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSwitchGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 交换机组信息

		SwitchGroup []*string `json:"SwitchGroup,omitempty" name:"SwitchGroup"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSwitchGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSwitchGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExtendCidr struct {
	// 外部用扩展cidr。

	ExtendCidrBlock *string `json:"ExtendCidrBlock,omitempty" name:"ExtendCidrBlock"`
}

type GlobalExtendCidr struct {
	// 无

	Subnet *string `json:"Subnet,omitempty" name:"Subnet"`
	// 无

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 无

	Mask *string `json:"Mask,omitempty" name:"Mask"`
	// 无

	IntMask *uint64 `json:"IntMask,omitempty" name:"IntMask"`
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type CreateSubnetsExRequest struct {
	*tchttp.BaseRequest

	// 创建的子网对象输入信息列表

	Subnets []*SubnetInput `json:"Subnets,omitempty" name:"Subnets"`
}

func (r *CreateSubnetsExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubnetsExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcPeersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// vpc对等连接列表

		VpcPeerSet []*VpcPeer `json:"VpcPeerSet,omitempty" name:"VpcPeerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcPeersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcPeersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceRoute struct {
	// ServiceId

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// VpcId

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// Routes

	Routes []*ServiceRouteRoute `json:"Routes,omitempty" name:"Routes"`
}

type UNatGwRuleWanIp struct {
	// CarrierId

	CarrierId *string `json:"CarrierId,omitempty" name:"CarrierId"`
	// Ips

	Ips []*string `json:"Ips,omitempty" name:"Ips"`
}

type DescribeVpcGlobalExtendCidrResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalCount

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// GlobalExtendCidrSet

		GlobalExtendCidrSet []*GlobalExtendCidr `json:"GlobalExtendCidrSet,omitempty" name:"GlobalExtendCidrSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcGlobalExtendCidrResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcGlobalExtendCidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSubnetsRequest struct {
	*tchttp.BaseRequest

	// 请求删除的subnet信息列表

	SubnetDels []*SubnetDel `json:"SubnetDels,omitempty" name:"SubnetDels"`
}

func (r *DeleteSubnetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSubnetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProtectPolicyIpRequest struct {
	*tchttp.BaseRequest

	// Rip.

	Rip *string `json:"Rip,omitempty" name:"Rip"`
}

func (r *DescribeProtectPolicyIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProtectPolicyIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcLimitRequest struct {
	*tchttp.BaseRequest

	// 配额类型，Type取值范围：
	// * 0：每个开发商每个地域可创建的VPC数；
	// * 1：每个VPC可创建的子网数；
	// * 2：每个VPC可创建的路由表数；
	// * 3：每个路由表可添加的策略数；
	// * 4：每个VPC可创建的VPN网关数；
	// *  5：每个开发商可创建的对端网关数；
	// *  6：每个开发商可创建的VPN通道数；
	// *  7：每个对端网关可创建的VPN通道数；
	// * 8：每个VPNGW可以创建的通道数；
	// * 15：每个VPC可创建的网络ACL数；
	// * 17：每个网络ACL可添加的入站规则数；
	// * 18：每个网络ACL可添加的出站规则数；
	// * 19：每个VPC可创建的对等连接数；
	// * 20：每个VPC可创建的有效对等连接数；
	// * 22：每个VPC可创建的基础网络云主机与VPC互通数；
	// * 32：每个专线网关可创建的SNAT数；
	// * 34：每个专线网关可创建的DNAT数；
	// *38：每个专线网关可创建的SNAPT数；
	// * 36：每个专线网关可创建的DNAPT数；
	// *55：每个VPC可创建的NAT网关数；
	// * 56：每个NAT可以购买的外网IP数量；
	// * 57：每个VPC可创建弹性网卡数；
	// * 89：每个VPC可创建HAVIP数；
	// * 61：每个ENI可以绑定的内网IP数（ENI未绑定子机）；
	// * 74：每个NAT网关可创建的DNAPT数；
	// * 102：每个VPC可分配的IPv6地址数；
	// * 105：每个ENI可分配的IPv6地址数；
	// * 107：每个VPC可分配的辅助CIDR数；

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 私有网络配额值。

	Val *uint64 `json:"Val,omitempty" name:"Val"`
	// 用户APPID。

	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

func (r *ModifyVpcLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcLimitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductFailureMigrateTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 切换任务ID

	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
}

func (r *QueryProductFailureMigrateTaskDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductFailureMigrateTaskDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductHealthStateRequest struct {
	*tchttp.BaseRequest

	// 产品名

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 特殊参数

	Params []*string `json:"Params,omitempty" name:"Params"`
}

func (r *QueryProductHealthStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductHealthStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVmListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// vm信息

		VmSet []*VmSet `json:"VmSet,omitempty" name:"VmSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVmListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVmListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcPeer struct {
	// 接收方VPC实例ID。如45678

	DVpcId *uint64 `json:"DVpcId,omitempty" name:"DVpcId"`
	// 类型。如0

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// WanOutLimit

	WanOutLimit *uint64 `json:"WanOutLimit,omitempty" name:"WanOutLimit"`
	// ImageId

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// BPublishDockerCidr

	BPublishDockerCidr *uint64 `json:"BPublishDockerCidr,omitempty" name:"BPublishDockerCidr"`
	// DVpcSubnet

	DVpcSubnet *string `json:"DVpcSubnet,omitempty" name:"DVpcSubnet"`
	// AfcVpcPeerGlobalId

	AfcVpcPeerGlobalId *uint64 `json:"AfcVpcPeerGlobalId,omitempty" name:"AfcVpcPeerGlobalId"`
	// SVpcIntMask

	SVpcIntMask *uint64 `json:"SVpcIntMask,omitempty" name:"SVpcIntMask"`
	// SViewFlag

	SViewFlag *uint64 `json:"SViewFlag,omitempty" name:"SViewFlag"`
	// DVpcIntMask

	DVpcIntMask *uint64 `json:"DVpcIntMask,omitempty" name:"DVpcIntMask"`
	// DViewFlag

	DViewFlag *uint64 `json:"DViewFlag,omitempty" name:"DViewFlag"`
	// BZoneVpcPeer

	BZoneVpcPeer *uint64 `json:"BZoneVpcPeer,omitempty" name:"BZoneVpcPeer"`
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
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// UniqVpcPeerId

	UniqVpcPeerId *string `json:"UniqVpcPeerId,omitempty" name:"UniqVpcPeerId"`
	// SVpcSubnet

	SVpcSubnet *string `json:"SVpcSubnet,omitempty" name:"SVpcSubnet"`
	// LastState

	LastState *uint64 `json:"LastState,omitempty" name:"LastState"`
}

type AddUnderlayNatGatewayWanIpRequest struct {
	*tchttp.BaseRequest

	// underlay natgateway详细信息

	UnatgwWips []*UnatgwWip `json:"UnatgwWips,omitempty" name:"UnatgwWips"`
}

func (r *AddUnderlayNatGatewayWanIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUnderlayNatGatewayWanIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectGatewaysExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方式

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 过滤器列表

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDirectConnectGatewaysExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectGatewaysExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnderlayNatGatewayCarrierInfoRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeUnderlayNatGatewayCarrierInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnderlayNatGatewayCarrierInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVsrsRequest struct {
	*tchttp.BaseRequest

	// 无

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// 无

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	VpngwId *uint64 `json:"VpngwId,omitempty" name:"VpngwId"`
}

func (r *DescribeVsrsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVsrsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVsrsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 无

		VsrSet []*Vsr `json:"VsrSet,omitempty" name:"VsrSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVsrsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVsrsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupForCvmRequest struct {
	*tchttp.BaseRequest

	// 虚拟机id列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 虚拟机id

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeSecurityGroupForCvmRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupForCvmRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGatewayAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGatewayAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGatewayAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JnsGatewayService struct {
	// VpcId

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// ZoneId

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// VpcgwType

	VpcgwType *uint64 `json:"VpcgwType,omitempty" name:"VpcgwType"`
	// Bandwidth

	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// DpdkSnatip

	DpdkSnatip *string `json:"DpdkSnatip,omitempty" name:"DpdkSnatip"`
	// Business

	Business *string `json:"Business,omitempty" name:"Business"`
	// Business

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// UniqVpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
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

	Snatip []*Snatip `json:"Snatip,omitempty" name:"Snatip"`
}

type ModifyIpv6PrefixRequest struct {
	*tchttp.BaseRequest

	// 掩码范围；目前支持的范围56到64。

	SplitIntPrefix *int64 `json:"SplitIntPrefix,omitempty" name:"SplitIntPrefix"`
}

func (r *ModifyIpv6PrefixRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIpv6PrefixRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
	// 网关集群类型。如0

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 用户级别。如-1

	OwnerLevel *int64 `json:"OwnerLevel,omitempty" name:"OwnerLevel"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ArrayNfvGwRequest struct {
	*tchttp.BaseRequest

	// 网关名称

	GwName *string `json:"GwName,omitempty" name:"GwName"`
	// 网关类型

	GwType *string `json:"GwType,omitempty" name:"GwType"`
	// 节点规格

	Format *string `json:"Format,omitempty" name:"Format"`
	// 节点数量

	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`
	// 网关版本

	GwVersion *string `json:"GwVersion,omitempty" name:"GwVersion"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 交换机组

	SwitchGroup *string `json:"SwitchGroup,omitempty" name:"SwitchGroup"`
	// 外网交换机组

	OuterSwitchGroup *string `json:"OuterSwitchGroup,omitempty" name:"OuterSwitchGroup"`
	// tgw 部署vlan id信息

	WanVlanId *uint64 `json:"WanVlanId,omitempty" name:"WanVlanId"`
	// tgw外网vip段

	WanSegment *string `json:"WanSegment,omitempty" name:"WanSegment"`
	// tgw 可用区类型

	MultiAz *uint64 `json:"MultiAz,omitempty" name:"MultiAz"`
	// tgw 外网多运营商类型

	WanMultiIsp *uint64 `json:"WanMultiIsp,omitempty" name:"WanMultiIsp"`
	// tgw ospf 超时时间

	OspfDeadInterval *uint64 `json:"OspfDeadInterval,omitempty" name:"OspfDeadInterval"`
	// ipv6接口地址，多个用英文分号相连

	Ipv6NodeNetwork *string `json:"Ipv6NodeNetwork,omitempty" name:"Ipv6NodeNetwork"`
	// 生产节点虚拟化类型，如VSELF等

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// kni1_ip地址

	Kni1Ip *string `json:"Kni1Ip,omitempty" name:"Kni1Ip"`
	// nat64心跳ip和loalip网段信息

	TgwExtendSegment *string `json:"TgwExtendSegment,omitempty" name:"TgwExtendSegment"`
}

func (r *ArrayNfvGwRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ArrayNfvGwRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDetectHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDetectHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDetectHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNfvGwNodeRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
}

func (r *DeleteNfvGwNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNfvGwNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEniListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 弹性网卡信息

		EniSet []*EniSet `json:"EniSet,omitempty" name:"EniSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEniListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEniListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcEndPointServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合查询条件的个数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 终端节点服务对象数组。

		EndPointServiceSet []*EndPointService `json:"EndPointServiceSet,omitempty" name:"EndPointServiceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcEndPointServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcEndPointServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNfvGwLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志步骤数组

		Logs []*Logs `json:"Logs,omitempty" name:"Logs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNfvGwLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNfvGwLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GwNodes struct {
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 资源id

	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
	// master的vm

	MasterVm *string `json:"MasterVm,omitempty" name:"MasterVm"`
	// 格式

	Format *string `json:"Format,omitempty" name:"Format"`
	// 最后一跳开关

	LasthopSwitch *string `json:"LasthopSwitch,omitempty" name:"LasthopSwitch"`
	// node节点id

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// Rip

	Rip []*string `json:"Rip,omitempty" name:"Rip"`
	// SnatIp

	SnatIp *string `json:"SnatIp,omitempty" name:"SnatIp"`
	// 网关版本

	GwVersion *string `json:"GwVersion,omitempty" name:"GwVersion"`
	// 虚拟ip

	Vips []*Vips `json:"Vips,omitempty" name:"Vips"`
	// 交换机组信息

	SwitchGroup *string `json:"SwitchGroup,omitempty" name:"SwitchGroup"`
	// 可用区信息

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// OuterSwitchGroup

	OuterSwitchGroup *string `json:"OuterSwitchGroup,omitempty" name:"OuterSwitchGroup"`
	// NetworkInterfaces

	NetworkInterfaces *NetworkInterface `json:"NetworkInterfaces,omitempty" name:"NetworkInterfaces"`
	// ReplaceTaskId

	ReplaceTaskId *int64 `json:"ReplaceTaskId,omitempty" name:"ReplaceTaskId"`
	// InstanceType

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// ReplaceLog

	ReplaceLog []*string `json:"ReplaceLog,omitempty" name:"ReplaceLog"`
	// GroupInfo

	GroupInfo *GatewayGroup `json:"GroupInfo,omitempty" name:"GroupInfo"`
	// VpcgwGroupId

	VpcgwGroupId *int64 `json:"VpcgwGroupId,omitempty" name:"VpcgwGroupId"`
	// Weight

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// VpcgwType

	VpcgwType *int64 `json:"VpcgwType,omitempty" name:"VpcgwType"`
}

type UNatGwRuleWanIps struct {
	// Ips

	Ips []*string `json:"Ips,omitempty" name:"Ips"`
	// CarrierId

	CarrierId *string `json:"CarrierId,omitempty" name:"CarrierId"`
}

type UNatGwWanInfo struct {
	// 运营商

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
	// underlay natgateway 的ID

	UnatgwId *uint64 `json:"UnatgwId,omitempty" name:"UnatgwId"`
	// underlay natgateway网关唯一ID

	UniqUnatgwId *string `json:"UniqUnatgwId,omitempty" name:"UniqUnatgwId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeGatewayIpsRequest struct {
	*tchttp.BaseRequest

	// GatewayName

	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`
}

func (r *DescribeGatewayIpsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayIpsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateServiceRouteRequest struct {
	*tchttp.BaseRequest

	// VpcId数字ID。

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// Type

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// Proto

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// Vport

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 集群组数字ID。

	VpcgwId *int64 `json:"VpcgwId,omitempty" name:"VpcgwId"`
	// 规则数组。

	Routes []*RouteRule `json:"Routes,omitempty" name:"Routes"`
}

func (r *MigrateServiceRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateServiceRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeNfvGwNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误信息

		Error *Error `json:"Error,omitempty" name:"Error"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeNfvGwNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeNfvGwNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterState struct {
	// VPCGW大集群名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 集群状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 主节点所在AZ

	MainNodeAzList []*string `json:"MainNodeAzList,omitempty" name:"MainNodeAzList"`
	// 节点列表

	NodeList []*VPCGWNode `json:"NodeList,omitempty" name:"NodeList"`
}

type UNatGwRuleSet struct {
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实际的设备内网IP

	Rip []*string `json:"Rip,omitempty" name:"Rip"`
	// underlay natgateway的ID

	UnatgwId *uint64 `json:"UnatgwId,omitempty" name:"UnatgwId"`
	// underlay natgateway规则ID

	UnatgwRuleId *uint64 `json:"UnatgwRuleId,omitempty" name:"UnatgwRuleId"`
	// underlay natgateway网关唯一ID

	UniqUnatgwId *string `json:"UniqUnatgwId,omitempty" name:"UniqUnatgwId"`
	// underlay natgateway规则唯一ID

	UniqUnatgwRuleId *string `json:"UniqUnatgwRuleId,omitempty" name:"UniqUnatgwRuleId"`
	// 外网IP

	WanIp []*UNatGwRuleWanIps `json:"WanIp,omitempty" name:"WanIp"`
}

type ModifyVpcGlobalExtendCidrResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcGlobalExtendCidrResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcGlobalExtendCidrResponse) FromJsonString(s string) error {
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

type RouteService struct {
	// Alive

	Alive *uint64 `json:"Alive,omitempty" name:"Alive"`
	// BAutoActualMachine

	BAutoActualMachine *uint64 `json:"BAutoActualMachine,omitempty" name:"BAutoActualMachine"`
	// Description

	Description *string `json:"Description,omitempty" name:"Description"`
	// ForwardingMode

	ForwardingMode *uint64 `json:"ForwardingMode,omitempty" name:"ForwardingMode"`
	// HostIp

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// Pip

	Pip *string `json:"Pip,omitempty" name:"Pip"`
	// Pport

	Pport *uint64 `json:"Pport,omitempty" name:"Pport"`
	// Proto

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// UdpOption

	UdpOption *uint64 `json:"UdpOption,omitempty" name:"UdpOption"`
	// UniqVpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// VpcId

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// Vport

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// Weight

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
	// ZoneId

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type UnatgwRuleRipR struct {
	// 运营商

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
	// 网关实际内网IP

	Rip *string `json:"Rip,omitempty" name:"Rip"`
	// underlay natgateway规则的ID

	UnatgwRuleId *uint64 `json:"UnatgwRuleId,omitempty" name:"UnatgwRuleId"`
}

type VPCGWNode struct {
	// 内网IP

	PrivateIP *string `json:"PrivateIP,omitempty" name:"PrivateIP"`
	// 公网IP

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 机器名

	Host *string `json:"Host,omitempty" name:"Host"`
	// 机架

	Rack *string `json:"Rack,omitempty" name:"Rack"`
	// 可用区

	Az *string `json:"Az,omitempty" name:"Az"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 节点存活状态，可选项：'alive'、'dead'

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DeleteDetectInstallRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDetectInstallRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDetectInstallRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGatewayInstallRecordsRequest struct {
	*tchttp.BaseRequest

	// EventIds

	EventIds []*string `json:"EventIds,omitempty" name:"EventIds"`
}

func (r *DeleteGatewayInstallRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGatewayInstallRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type Filters struct {
	// 虚拟机内存大小

	VmMemorySize *int64 `json:"VmMemorySize,omitempty" name:"VmMemorySize"`
	// 虚拟机cpu数量

	VmCpuCount *int64 `json:"VmCpuCount,omitempty" name:"VmCpuCount"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 系统盘大小

	SystemDiskSize *int64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 数据磁盘大小

	DataDiskSize *int64 `json:"DataDiskSize,omitempty" name:"DataDiskSize"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeMulticastGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组播组对象数组。

		MulticastGroupSet []*MulticastGroup `json:"MulticastGroupSet,omitempty" name:"MulticastGroupSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMulticastGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMulticastGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNfvGwDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网关名称

		GwName *string `json:"GwName,omitempty" name:"GwName"`
		// 网关id

		ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
		// 网关类型

		GwType *string `json:"GwType,omitempty" name:"GwType"`
		// 网关模式

		GwMode *int64 `json:"GwMode,omitempty" name:"GwMode"`
		// 可用区

		Zone *string `json:"Zone,omitempty" name:"Zone"`
		// 节点数目

		NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`
		// 创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 虚拟ip

		Vips []*Vips `json:"Vips,omitempty" name:"Vips"`
		// 网关状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNfvGwDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNfvGwDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProtectPolicyIpRequest struct {
	*tchttp.BaseRequest

	// Rip。

	Rip *string `json:"Rip,omitempty" name:"Rip"`
	// PPS限制值（个/秒）。

	PktLimit *uint64 `json:"PktLimit,omitempty" name:"PktLimit"`
	// 宽带限制值（Mbps）。

	TcLimit *uint64 `json:"TcLimit,omitempty" name:"TcLimit"`
	// 并发连接数限制值（个）。

	ConnLimit *uint64 `json:"ConnLimit,omitempty" name:"ConnLimit"`
	// 新增连接限制值（个）。

	IncconnLimit *uint64 `json:"IncconnLimit,omitempty" name:"IncconnLimit"`
	// 是否开启过载保护，0不开启，1开启，默认开启。

	Switch *uint64 `json:"Switch,omitempty" name:"Switch"`
}

func (r *ModifyProtectPolicyIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProtectPolicyIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchUnderlayNatGatewayHaStatusRequest struct {
	*tchttp.BaseRequest

	// NatgwMasterRIp

	NatgwMasterRIp *string `json:"NatgwMasterRIp,omitempty" name:"NatgwMasterRIp"`
}

func (r *SwitchUnderlayNatGatewayHaStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchUnderlayNatGatewayHaStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDetectInstallationRequest struct {
	*tchttp.BaseRequest

	// GatewayName

	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`
	// Ips

	Ips []*string `json:"Ips,omitempty" name:"Ips"`
	// AzName

	AzName *string `json:"AzName,omitempty" name:"AzName"`
	// GroupId

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// Ver

	Ver *string `json:"Ver,omitempty" name:"Ver"`
}

func (r *CreateDetectInstallationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDetectInstallationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIpV6SubnetsRequest struct {
	*tchttp.BaseRequest

	// 网段。

	Address *string `json:"Address,omitempty" name:"Address"`
	// 掩码。

	IntPrefix *uint64 `json:"IntPrefix,omitempty" name:"IntPrefix"`
	// 分解掩码，目前固定取值56.

	SplitIntPrefix *uint64 `json:"SplitIntPrefix,omitempty" name:"SplitIntPrefix"`
}

func (r *DeleteIpV6SubnetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIpV6SubnetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUnderlayNatGatewayWanInfoAndWanIpAttributeRequest struct {
	*tchttp.BaseRequest

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// IP地址

	Ip []*string `json:"Ip,omitempty" name:"Ip"`
	// 运营商

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
	// underlay natgateway的ID

	UnatgwId *uint64 `json:"UnatgwId,omitempty" name:"UnatgwId"`
}

func (r *ModifyUnderlayNatGatewayWanInfoAndWanIpAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUnderlayNatGatewayWanInfoAndWanIpAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGatewayInstallationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// EventId

		EventId *string `json:"EventId,omitempty" name:"EventId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGatewayInstallationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGatewayInstallationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInterfacesExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段,可选值有：createTime, owner。
	// OrderField与OrderDirection一起传入排序才能生效。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方式,asc：升序， desc：降序，默认升序。

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 过滤条件,可选字段有：
	// <li>UniqVpcId - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>UniqSubnetId - String - （过滤条件）所属子网实例ID，形如：subnet-f49l6u0z。</li>
	// <li>UniqEniId - String - （过滤条件）弹性网卡实例ID，形如：eni-5k56k7k7。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNetworkInterfacesExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkInterfacesExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnderlayNatGatewayRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// underlay natgateway规则详细信息

		UNatGwRuleSet []*UNatGwRuleSet `json:"UNatGwRuleSet,omitempty" name:"UNatGwRuleSet"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUnderlayNatGatewayRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnderlayNatGatewayRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGatewayRequest struct {
	*tchttp.BaseRequest

	// 网关集群类型(0-pvgw 1-vpcgw 2-jnsgw 3-natgw 4-mcgw 5-dcgw 6-pcgw 7-xgw 8-sxgw 9-underly-natgw)

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// EnableNat

	EnableNat *string `json:"EnableNat,omitempty" name:"EnableNat"`
	// 网关集群名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 网关节点IP列表，不同的网关集群类型支持的IP列表大小不同，pcgw dcgw mcgw支持两个和两个以上网关节点IP，xgw sxgw为单个网关节点IP，其他类型网关只能为两个网关节点IP，多个IP用分号隔开，如10.120.66.135;10.120.66.136

	Rips *string `json:"Rips,omitempty" name:"Rips"`
	// 权重，如20

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
	// vpcgw网关集群组ID，只有Type为vpcgw网关集群时才需要传递这个参数

	VpcgwGroupId *uint64 `json:"VpcgwGroupId,omitempty" name:"VpcgwGroupId"`
	// 地域ID，Type为xgw和sxgw集群网关时，此参数为必填

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 可用区ID，Type为sxgw, natgw集群网关时，此参数为必填

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// Zone，Type为natgw集群网关时，此参数为必填

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 网关类型，0 :ko,1:dpdk

	VpcgwType *int64 `json:"VpcgwType,omitempty" name:"VpcgwType"`
	// 可用区名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

func (r *CreateGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGatewayRequest struct {
	*tchttp.BaseRequest

	// 过滤参数EnableNat

	EnableNat *string `json:"EnableNat,omitempty" name:"EnableNat"`
	// 网关集群类型(0-pvgw 1-vpcgw 2-jnsgw 3-natgw 4-mcgw 5-dcgw 6-pcgw 7-xgw 8-sxgw 9-underly-natgw)

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 网关集群Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 网关集群类型为dcgw时为必传参数，网关集群类型为pcgw为可选参数，其他类型网关不用传递该参数

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// Rips，如10.120.66.135;10.120.66.136，网关集群类型为pcgw时为可选参数

	Rips *string `json:"Rips,omitempty" name:"Rips"`
}

func (r *DeleteGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpv6PrefixResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SplitIntPrefix

		SplitIntPrefix *uint64 `json:"SplitIntPrefix,omitempty" name:"SplitIntPrefix"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpv6PrefixResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpv6PrefixResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResultRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GwFormats struct {
	// 虚拟机内存大小

	VmMemorySize *int64 `json:"VmMemorySize,omitempty" name:"VmMemorySize"`
	// 虚拟机cpu数量

	VmCpuCount *int64 `json:"VmCpuCount,omitempty" name:"VmCpuCount"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 系统盘大小

	SystemDiskSize *int64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`
	// 值

	Value *int64 `json:"Value,omitempty" name:"Value"`
	// 数据磁盘大小

	DataDiskSize *int64 `json:"DataDiskSize,omitempty" name:"DataDiskSize"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeGatewayNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalCount

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// GatewayNodeSet

		GatewayNodeSet []*string `json:"GatewayNodeSet,omitempty" name:"GatewayNodeSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Metric struct {
	// 指标名称，例如：InstanceHealthState

	Name *string `json:"Name,omitempty" name:"Name"`
	// 健康状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 健康信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type CreateUnderlayNatGatewayWanInfoAndWanIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败信息

		Error *string `json:"Error,omitempty" name:"Error"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUnderlayNatGatewayWanInfoAndWanIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUnderlayNatGatewayWanInfoAndWanIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnConnsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		VpnConnSet []*VpnConn `json:"VpnConnSet,omitempty" name:"VpnConnSet"`
		// 无

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpnConnsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnConnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNfvCvmInstanceTypeRequest struct {
	*tchttp.BaseRequest

	// 获取数据的起始位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 要获取的数据条目数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetNfvCvmInstanceTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNfvCvmInstanceTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClbGatewayTypeNamesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalCount

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// GateWayTypeNameSet

		GateWayTypeNameSet []*string `json:"GateWayTypeNameSet,omitempty" name:"GateWayTypeNameSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClbGatewayTypeNamesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClbGatewayTypeNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpv6PrefixRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeIpv6PrefixRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpv6PrefixRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpV6SubnetsRequest struct {
	*tchttp.BaseRequest

	// 过滤器

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 条数限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeIpV6SubnetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpV6SubnetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRtbsExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 条数限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤器列表

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方式

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeRtbsExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRtbsExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnderlayNatGatewayCarriorInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalCount

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// CarrierInfoSet

		CarrierInfoSet []*CarrierInfo `json:"CarrierInfoSet,omitempty" name:"CarrierInfoSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUnderlayNatGatewayCarriorInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnderlayNatGatewayCarriorInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIpv6PrefixResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIpv6PrefixResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIpv6PrefixResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 条数限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序类型

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 过滤器列表

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJnsGatewayServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalCount

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// JnsGatewayServiceSet

		JnsGatewayServiceSet []*JnsGatewayService `json:"JnsGatewayServiceSet,omitempty" name:"JnsGatewayServiceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJnsGatewayServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJnsGatewayServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnGatewayDetailRequest struct {
	*tchttp.BaseRequest

	// VpcId

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// VpngwId

	VpngwId *uint64 `json:"VpngwId,omitempty" name:"VpngwId"`
	// Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

func (r *DescribeVpnGatewayDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnGatewayDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNfvSwitchGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 交换机组信息

		SwitchGroup []*string `json:"SwitchGroup,omitempty" name:"SwitchGroup"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNfvSwitchGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNfvSwitchGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	EniType *uint64 `json:"EniType,omitempty" name:"EniType"`
	// eni name

	EniName *string `json:"EniName,omitempty" name:"EniName"`
}

type CreateGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误信息

		Error *Error `json:"Error,omitempty" name:"Error"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNfvGwNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNfvGwNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNfvGwNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 子网类型

	Type []*uint64 `json:"Type,omitempty" name:"Type"`
	// 条数限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件，参数不支持同时指定VpcIds和Filters。
	// <li>AppId - String - （过滤条件）用户的AppId</li>
	// <li>UniqVpcId - String - （过滤条件）用户的vpc唯一id。</li>
	// <li>UniqSubnetId- String - （过滤条件）用户子网的唯一id。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方式

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 子网唯一ID。

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
}

func (r *DescribeSubnetExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubnetExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnderlayNatGatewayWanInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUnderlayNatGatewayWanInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnderlayNatGatewayWanInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsgPolicyInfo struct {
	// ErrorCode

	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 安全组唯一id

	UsgId *string `json:"UsgId,omitempty" name:"UsgId"`
	// InBound

	InBound []*InOutBoundInfo `json:"InBound,omitempty" name:"InBound"`
	// ErrorInfo

	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
	// OutBound

	OutBound []*InOutBoundInfo `json:"OutBound,omitempty" name:"OutBound"`
	// Version

	Version *uint64 `json:"Version,omitempty" name:"Version"`
}

type DeleteServiceRequest struct {
	*tchttp.BaseRequest

	// 无

	VpcServiceRouleSet []*VpcServiceRoule `json:"VpcServiceRouleSet,omitempty" name:"VpcServiceRouleSet"`
}

func (r *DeleteServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHaVipListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，参数不支持同时指定`HaVipIds`和`Filters`。
	// <li>AppId - String - 用户AppId。</li>
	// <li>UniqVpcId- String - `HAVIP`所在私有网络`ID`。</li>
	// <li>Subnet - String - `HAVIP`所在子网`ID`。</li>
	// <li>UniqHavipId- String - `HAVIP`唯一`ID`。</li>
	// <li>HaVipName - String - `HAVIP`名字。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方向，`desc`降序，`asc`升序

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 状态，支持`UNBIND`,`AVAILABLE`两种状态的查询。

	State []*string `json:"State,omitempty" name:"State"`
}

func (r *DescribeHaVipListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHaVipListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnderlayNatGatewayCarrierInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运营商信息条目数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 运营商详细信息

		CarrierInfoSet []*CarrierInfo `json:"CarrierInfoSet,omitempty" name:"CarrierInfoSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUnderlayNatGatewayCarrierInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnderlayNatGatewayCarrierInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVmIpv6InfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 子机信息

		VmSet []*VmSet `json:"VmSet,omitempty" name:"VmSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVmIpv6InfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVmIpv6InfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PassiveHealthCheckInfo struct {
	// 超时

	FailTimeout *string `json:"FailTimeout,omitempty" name:"FailTimeout"`
	// 最大失败次数

	MaxFails *string `json:"MaxFails,omitempty" name:"MaxFails"`
}

type GatewayInstallStateSet struct {
	// 无

	Gwname *string `json:"Gwname,omitempty" name:"Gwname"`
	// 无

	LanSegment *string `json:"LanSegment,omitempty" name:"LanSegment"`
	// 无

	WanSegment *string `json:"WanSegment,omitempty" name:"WanSegment"`
	// 无

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
	// 无

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// 无

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 无

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 无

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 无

	OuterIp *string `json:"OuterIp,omitempty" name:"OuterIp"`
	// 无

	Aztype *string `json:"Aztype,omitempty" name:"Aztype"`
	// 无

	MultiispType *string `json:"MultiispType,omitempty" name:"MultiispType"`
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
}

type DescribeRtbSubnetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 子网列表

		SubnetSet []*SubNetInfo `json:"SubnetSet,omitempty" name:"SubnetSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRtbSubnetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRtbSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDetectInstallationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// EventId

		EventId *string `json:"EventId,omitempty" name:"EventId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDetectInstallationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDetectInstallationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNfvGwTypesRequest struct {
	*tchttp.BaseRequest
}

func (r *GetNfvGwTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNfvGwTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AdjustNfvGwNodeRequest struct {
	*tchttp.BaseRequest

	// 操作类型：隔离节点/启用节点

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 实例id

	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
}

func (r *AdjustNfvGwNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AdjustNfvGwNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayVersionsRequest struct {
	*tchttp.BaseRequest

	// Ips

	Ips []*string `json:"Ips,omitempty" name:"Ips"`
	// GatewayName

	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`
}

func (r *DescribeGatewayVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnatgwInfo struct {
	// underlay natgateway的ID

	UnatgwId *uint64 `json:"UnatgwId,omitempty" name:"UnatgwId"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 网关VIP

	GwIp *string `json:"GwIp,omitempty" name:"GwIp"`
	// underlay natgateway网关唯一ID

	UniqUnatgwId *string `json:"UniqUnatgwId,omitempty" name:"UniqUnatgwId"`
	// IDC名称

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 网关RIP

	Rip *string `json:"Rip,omitempty" name:"Rip"`
	// 可用区

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type CreateProductFailureMigrateRsp struct {
	// 迁移任务UUID

	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
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

type DescribeUnderlayNatGatewayWanInfoAndWanIpRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUnderlayNatGatewayWanInfoAndWanIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnderlayNatGatewayWanInfoAndWanIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnderlayNatGatewayWanIpsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// underlay natgateway信息条目数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// underlay natgateway详细信息

		UNatGwWanIpSet []*UNatGwWanIp `json:"UNatGwWanIpSet,omitempty" name:"UNatGwWanIpSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUnderlayNatGatewayWanIpsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnderlayNatGatewayWanIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNfvGwFormatsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网关所有规格-数组

		GwFormats []*GwFormats `json:"GwFormats,omitempty" name:"GwFormats"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNfvGwFormatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNfvGwFormatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGatewayWeightRequest struct {
	*tchttp.BaseRequest

	// 网关集群类型(0-pvgw 1-vpcgw 2-jnsgw 3-natgw 4-mcgw 5-dcgw 6-pcgw  9-underly-natgw)

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 网关集群Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 新的权重值，如90

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
}

func (r *ModifyGatewayWeightRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGatewayWeightRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubNetInfo struct {
	// VpcId

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// ZoneId

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// UniqVpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// Ipv6CidrBlock

	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`
	// VpcName

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// SubnetId

	SubnetId *uint64 `json:"SubnetId,omitempty" name:"SubnetId"`
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

	AvailableIpNum *uint64 `json:"AvailableIpNum,omitempty" name:"AvailableIpNum"`
	// UsedIpNum

	UsedIpNum *uint64 `json:"UsedIpNum,omitempty" name:"UsedIpNum"`
	// TotalIpNum

	TotalIpNum *uint64 `json:"TotalIpNum,omitempty" name:"TotalIpNum"`
	// Mask

	Mask *string `json:"Mask,omitempty" name:"Mask"`
	// cdc集群id

	UniqCdcId *string `json:"UniqCdcId,omitempty" name:"UniqCdcId"`
}

type GetNfvGwListRequest struct {
	*tchttp.BaseRequest

	// 索引

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤选项

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *GetNfvGwListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNfvGwListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchUnderlayNatGatewayHaStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchUnderlayNatGatewayHaStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchUnderlayNatGatewayHaStatusResponse) FromJsonString(s string) error {
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

type ChangeNfvGwNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ChangeNfvGwNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeNfvGwNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGatewayCLBInstallationRequest struct {
	*tchttp.BaseRequest

	// Ips

	Ips []*string `json:"Ips,omitempty" name:"Ips"`
	// Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// Ver

	Ver *string `json:"Ver,omitempty" name:"Ver"`
	// GatewayName

	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`
	// IdcName

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// WanSegment

	WanSegment *string `json:"WanSegment,omitempty" name:"WanSegment"`
	// LanSegment

	LanSegment *string `json:"LanSegment,omitempty" name:"LanSegment"`
	// IntranetType

	IntranetType *string `json:"IntranetType,omitempty" name:"IntranetType"`
	// Az

	Az *uint64 `json:"Az,omitempty" name:"Az"`
	// MultiIsp

	MultiIsp *uint64 `json:"MultiIsp,omitempty" name:"MultiIsp"`
	// wan_if

	WanIf []*WanIf `json:"WanIf,omitempty" name:"WanIf"`
	// Kni1Ip

	Kni1Ip *string `json:"Kni1Ip,omitempty" name:"Kni1Ip"`
	// V6Segment

	V6Segment *string `json:"V6Segment,omitempty" name:"V6Segment"`
}

func (r *CreateGatewayCLBInstallationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGatewayCLBInstallationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGatewayGroupRequest struct {
	*tchttp.BaseRequest

	// 网关集群类型(0-pvgw 1-vpcgw 2-jnsgw 3-natgw 4-mcgw 5-dcgw 6-pcgw 7-xgw 8-sxgw 9-underly-natgw)

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 请求删除的网关集群组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteGatewayGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGatewayGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIpV6SubnetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIpV6SubnetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIpV6SubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClbGatewayVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalCount

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// GateWayVersionSet

		GateWayVersionSet []*VpnGateway `json:"GateWayVersionSet,omitempty" name:"GateWayVersionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClbGatewayVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClbGatewayVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatDetectIpsRequest struct {
	*tchttp.BaseRequest

	// 过滤字段，仅支持 `Dip` 查询。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNatDetectIpsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatDetectIpsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type RegionPair struct {
	// 本端地域

	SrcRegion *string `json:"SrcRegion,omitempty" name:"SrcRegion"`
	// 对端地域

	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type CreateBasicNetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新建的vpc0网段子网列表

		BasicNetSet []*SubnetBaseInfo `json:"BasicNetSet,omitempty" name:"BasicNetSet"`
		// 错误信息

		Error *Error `json:"Error,omitempty" name:"Error"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBasicNetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBasicNetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 网关集群组

		GatewayGroupSet []*GatewayGroup `json:"GatewayGroupSet,omitempty" name:"GatewayGroupSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAclAndItemRequest struct {
	*tchttp.BaseRequest

	// AppId

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// VPC实例ID。如66379

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// ACl规则实例ID。如243

	AclId *uint64 `json:"AclId,omitempty" name:"AclId"`
	// 0-入站规则 1-出站规则

	Dir *uint64 `json:"Dir,omitempty" name:"Dir"`
}

func (r *GetAclAndItemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAclAndItemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNfvGwLogsRequest struct {
	*tchttp.BaseRequest

	// 网关id

	GwId *int64 `json:"GwId,omitempty" name:"GwId"`
}

func (r *GetNfvGwLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNfvGwLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDetectRouteStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDetectRouteStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDetectRouteStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type CreateSubnetsExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建的子网对象输出信息列表

		SubnetSet []*SubnetOutput `json:"SubnetSet,omitempty" name:"SubnetSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubnetsExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubnetsExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollbackNfvGwLogStepResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RollbackNfvGwLogStepResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RollbackNfvGwLogStepResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 字段名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type JnsGatewayServicePVm struct {
	// HostIp

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// UniqVgwIndex

	UniqVgwIndex *string `json:"UniqVgwIndex,omitempty" name:"UniqVgwIndex"`
	// VgwIndex

	VgwIndex *string `json:"VgwIndex,omitempty" name:"VgwIndex"`
}

type ModifyGatewayGroupAttributeRequest struct {
	*tchttp.BaseRequest

	// 网关类型(0-pvgw 1-vpcgw 2-jnsgw 3-natgw 4-mcgw 5-dcgw 6-pcgw 7-xgw 8-sxgw 9-underly-natgw)

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 网关集群组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 网关集群组修改后的新名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 网关集群组修改后的新权重

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
	// 网关集群组修改后的新备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// appid

	Owners *string `json:"Owners,omitempty" name:"Owners"`
}

func (r *ModifyGatewayGroupAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGatewayGroupAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcLimitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddUnderlayNatGatewayRuleWipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUnderlayNatGatewayRuleWipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUnderlayNatGatewayRuleWipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Snatip struct {
	// Snatip

	Snatip *string `json:"Snatip,omitempty" name:"Snatip"`
}

type DescribeClbGatewayTypeNamesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeClbGatewayTypeNamesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClbGatewayTypeNamesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCvmForSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 条数限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 安全组id

	SgId *string `json:"SgId,omitempty" name:"SgId"`
	// 业务id

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// Filters

	Filter []*Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeCvmForSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCvmForSecurityGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayClbInstallStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// GatewayInstallStateSet

		GatewayInstallStateSet []*GatewayInstallStateSet `json:"GatewayInstallStateSet,omitempty" name:"GatewayInstallStateSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayClbInstallStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayClbInstallStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GwTypes struct {
	// 网关模块

	GwMode *int64 `json:"GwMode,omitempty" name:"GwMode"`
	// 网关Vision

	GwVision []*string `json:"GwVision,omitempty" name:"GwVision"`
	// 网关类型

	GwType *string `json:"GwType,omitempty" name:"GwType"`
	// 允许操作

	AllowOperate []*string `json:"AllowOperate,omitempty" name:"AllowOperate"`
	// 是否为外网网关类型

	GwOuterType *uint64 `json:"GwOuterType,omitempty" name:"GwOuterType"`
	// 是否支持跨az部署

	IsSupportMultiAz *uint64 `json:"IsSupportMultiAz,omitempty" name:"IsSupportMultiAz"`
	// 跨az 相同的vip字段

	MultiAzSameVip []*string `json:"MultiAzSameVip,omitempty" name:"MultiAzSameVip"`
	// IsGlobalAzKeepVipSame

	IsGlobalAzKeepVipSame *uint64 `json:"IsGlobalAzKeepVipSame,omitempty" name:"IsGlobalAzKeepVipSame"`
}

type SubnetSubMask struct {
	// IntMask

	IntMask *uint64 `json:"IntMask,omitempty" name:"IntMask"`
	// Mask

	Mask *string `json:"Mask,omitempty" name:"Mask"`
	// Subnet

	Subnet *string `json:"Subnet,omitempty" name:"Subnet"`
}

type AddUnderlayNatGatewayRuleWipRequest struct {
	*tchttp.BaseRequest

	// underlay natgateway规则详细信息

	UnatgwRuleWips []*UNatGwRuleWip `json:"UnatgwRuleWips,omitempty" name:"UnatgwRuleWips"`
}

func (r *AddUnderlayNatGatewayRuleWipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUnderlayNatGatewayRuleWipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGatewayGroupRequest struct {
	*tchttp.BaseRequest

	// 网关集群组名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 权重，如20

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
	// 网关集群类型(0-pvgw 1-vpcgw 2-jnsgw 3-natgw 4-mcgw 5-dcgw 6-pcgw 7-xgw 8-sxgw 9-underly-natgw)

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// appid

	Owners *string `json:"Owners,omitempty" name:"Owners"`
	// cdc集群id

	UniqCdcId *string `json:"UniqCdcId,omitempty" name:"UniqCdcId"`
}

func (r *CreateGatewayGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGatewayGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNatDetectIpRequest struct {
	*tchttp.BaseRequest

	// IP。

	Dip *string `json:"Dip,omitempty" name:"Dip"`
}

func (r *DeleteNatDetectIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNatDetectIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRtbsExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 路由信息列表

		RtbSet []*RtbInfo `json:"RtbSet,omitempty" name:"RtbSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRtbsExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRtbsExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SgVmCount struct {
	// general

	General *uint64 `json:"General,omitempty" name:"General"`
	// cvm

	Cvm *uint64 `json:"Cvm,omitempty" name:"Cvm"`
}

type DeleteSubnetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSubnetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcGlobalExtendCidrRequest struct {
	*tchttp.BaseRequest

	// AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// BAdd

	BAdd *bool `json:"BAdd,omitempty" name:"BAdd"`
	// CidrBlock

	CidrBlock []*string `json:"CidrBlock,omitempty" name:"CidrBlock"`
}

func (r *ModifyVpcGlobalExtendCidrRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcGlobalExtendCidrRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceRouteRoute struct {
	// Pip

	Pip *string `json:"Pip,omitempty" name:"Pip"`
	// Pport

	Pport *uint64 `json:"Pport,omitempty" name:"Pport"`
	// Description

	Description *string `json:"Description,omitempty" name:"Description"`
	// ZoneId

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type AddUnderlayNatGatewayWanIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUnderlayNatGatewayWanIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUnderlayNatGatewayWanIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TaskId

		TaskId []*uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 错误信息

		Error *Error `json:"Error,omitempty" name:"Error"`
		// 网关迁移返回码

		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigrateGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebootNfvGwLogStepRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 步骤名

	Step *string `json:"Step,omitempty" name:"Step"`
}

func (r *RebootNfvGwLogStepRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootNfvGwLogStepRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicNetsRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 条数限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方式

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeBasicNetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBasicNetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InOutBoundInfo struct {
	// Proto

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// Modifytime

	Modifytime *string `json:"Modifytime,omitempty" name:"Modifytime"`
	// Ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// Ip

	Ip6 *string `json:"Ip6,omitempty" name:"Ip6"`
	// Action

	Action *string `json:"Action,omitempty" name:"Action"`
	// Port

	Port *string `json:"Port,omitempty" name:"Port"`
	// Desc

	Desc *string `json:"Desc,omitempty" name:"Desc"`
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

type RtbInfo struct {
	// VPC实例ID。如65967

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// VPC唯一ID。如vpc-4cukfjyf

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// VPC名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// appid

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// nickname

	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
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

type ArrayNfvGwResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ArrayNfvGwResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ArrayNfvGwResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNfvGwNameRequest struct {
	*tchttp.BaseRequest

	// 网关id

	GwId *int64 `json:"GwId,omitempty" name:"GwId"`
	// 新修改的网关名称

	GwName *string `json:"GwName,omitempty" name:"GwName"`
}

func (r *UpdateNfvGwNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNfvGwNameRequest) FromJsonString(s string) error {
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
	// AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// BindResourceId

	BindResourceId *int64 `json:"BindResourceId,omitempty" name:"BindResourceId"`
	// BindResourceType

	BindResourceType *string `json:"BindResourceType,omitempty" name:"BindResourceType"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Creator

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// InstanceUuid

	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
	// IpAddress

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// IpGateway

	IpGateway *string `json:"IpGateway,omitempty" name:"IpGateway"`
	// IpMask

	IpMask *string `json:"IpMask,omitempty" name:"IpMask"`
	// IpV6Address

	IpV6Address *string `json:"IpV6Address,omitempty" name:"IpV6Address"`
	// IpV6Gateway

	IpV6Gateway *string `json:"IpV6Gateway,omitempty" name:"IpV6Gateway"`
	// NetworkDevice

	NetworkDevice *string `json:"NetworkDevice,omitempty" name:"NetworkDevice"`
	// NetworkType

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// UpdateTime

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// Vlan

	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`
}

type DescribeDetectPkgVersionsRequest struct {
	*tchttp.BaseRequest

	// GatewayName

	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`
}

func (r *DescribeDetectPkgVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDetectPkgVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatDetectIpsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// IP对象。

		IpSet []*DetectIp `json:"IpSet,omitempty" name:"IpSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatDetectIpsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatDetectIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDetectHostWeightResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDetectHostWeightResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDetectHostWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcInfo struct {
	// VpcId

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
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

	RtbNum *uint64 `json:"RtbNum,omitempty" name:"RtbNum"`
	// SubnetNum

	SubnetNum *uint64 `json:"SubnetNum,omitempty" name:"SubnetNum"`
	// Ipv6CidrBlock

	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`
	// EnableMulticast

	EnableMulticast *bool `json:"EnableMulticast,omitempty" name:"EnableMulticast"`
	// VpcPeerNum

	VpcPeerNum *uint64 `json:"VpcPeerNum,omitempty" name:"VpcPeerNum"`
	// VpgNum

	VpgNum *uint64 `json:"VpgNum,omitempty" name:"VpgNum"`
	// VpngwNum

	VpngwNum *uint64 `json:"VpngwNum,omitempty" name:"VpngwNum"`
	// VmNum

	VmNum *uint64 `json:"VmNum,omitempty" name:"VmNum"`
	// NatNum

	NatNum *uint64 `json:"NatNum,omitempty" name:"NatNum"`
	// AclNum

	AclNum *uint64 `json:"AclNum,omitempty" name:"AclNum"`
	// 外部用扩展cidr

	ExtendCidrSet []*ExtendCidr `json:"ExtendCidrSet,omitempty" name:"ExtendCidrSet"`
}

type ProceedNfvGwLogStepRequest struct {
	*tchttp.BaseRequest

	// 任务类型

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 步骤类型

	Step *string `json:"Step,omitempty" name:"Step"`
}

func (r *ProceedNfvGwLogStepRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProceedNfvGwLogStepRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayClbInstallStateRequest struct {
	*tchttp.BaseRequest

	// InnerIp

	InnerIp *string `json:"InnerIp,omitempty" name:"InnerIp"`
	// GateWayName

	GateWayName *string `json:"GateWayName,omitempty" name:"GateWayName"`
	// EventId

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeGatewayClbInstallStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayClbInstallStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type ChangeNfvGwNodeRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
}

func (r *ChangeNfvGwNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeNfvGwNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Detail struct {
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 步骤

	MiniStep *string `json:"MiniStep,omitempty" name:"MiniStep"`
}

type GatewayInfo struct {
	// 唯一键值

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 集群名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 多个rip，分号隔开

	Rip *string `json:"Rip,omitempty" name:"Rip"`
	// 地域

	Region *uint64 `json:"Region,omitempty" name:"Region"`
	// rip和zoneId对应关系的json

	ZoneMap *string `json:"ZoneMap,omitempty" name:"ZoneMap"`
	// ZoneId

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// ZoneName

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// Module

	Module *string `json:"Module,omitempty" name:"Module"`
	// MasterStatus

	MasterStatus *string `json:"MasterStatus,omitempty" name:"MasterStatus"`
	// VpcgwType

	VpcgwType *int64 `json:"VpcgwType,omitempty" name:"VpcgwType"`
	// VpcgwGroupId

	VpcgwGroupId *int64 `json:"VpcgwGroupId,omitempty" name:"VpcgwGroupId"`
	// GroupInfo

	GroupInfo *GatewayGroup `json:"GroupInfo,omitempty" name:"GroupInfo"`
	// Status

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// Rip1

	Rip1 *string `json:"Rip1,omitempty" name:"Rip1"`
	// Rip2

	Rip2 *string `json:"Rip2,omitempty" name:"Rip2"`
	// Zone

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// MigrateTaskId

	MigrateTaskId *string `json:"MigrateTaskId,omitempty" name:"MigrateTaskId"`
	// Weight

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type DescribeSubnetExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 子网列表

		SubNetSet []*SubNetInfo `json:"SubNetSet,omitempty" name:"SubNetSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubnetExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubnetExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 无

		VpcLimitSet []*VpcLimit `json:"VpcLimitSet,omitempty" name:"VpcLimitSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AclSubNetInfo struct {
	// Subnet

	Subnet *string `json:"Subnet,omitempty" name:"Subnet"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// BDefault

	BDefault *int64 `json:"BDefault,omitempty" name:"BDefault"`
	// Mask

	Mask *string `json:"Mask,omitempty" name:"Mask"`
	// ZoneId

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// IntMask

	IntMask *int64 `json:"IntMask,omitempty" name:"IntMask"`
	// UniqSubnetId

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
	// SubnetId

	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`
	// DhcpFlag

	DhcpFlag *int64 `json:"DhcpFlag,omitempty" name:"DhcpFlag"`
	// Bcdc

	BCdc *uint64 `json:"BCdc,omitempty" name:"BCdc"`
	// UniqCbcId

	UniqCdcId *string `json:"UniqCdcId,omitempty" name:"UniqCdcId"`
}

type DescribeNatGatewaysRequest struct {
	*tchttp.BaseRequest

	// 排序类型

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 条数限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤器列表

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNatGatewaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatGatewaysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupPolicyRequest struct {
	*tchttp.BaseRequest

	// 安全组id列表

	UsgIdSet []*string `json:"UsgIdSet,omitempty" name:"UsgIdSet"`
}

func (r *DescribeSecurityGroupPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNfvGwTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网关类型的数组

		GwTypes []*GwTypes `json:"GwTypes,omitempty" name:"GwTypes"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNfvGwTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNfvGwTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGatewayWeightResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGatewayWeightResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGatewayWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AzProductFailureMigrateTaskDetailResponse struct {
	// 切换任务ID

	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 进度

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// 执行信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 进度详情

	Detail []*AzMigrateDetailDO `json:"Detail,omitempty" name:"Detail"`
	// 补充数据

	ExtendData []*AzExtendDataDO `json:"ExtendData,omitempty" name:"ExtendData"`
}

type CreateUnderlayNatGatewayWanInfoAndWanIpRequest struct {
	*tchttp.BaseRequest

	// 外网IP

	WanIps []*string `json:"WanIps,omitempty" name:"WanIps"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 运营商

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
	// underlay natgateway的ID

	UnatgwId *uint64 `json:"UnatgwId,omitempty" name:"UnatgwId"`
}

func (r *CreateUnderlayNatGatewayWanInfoAndWanIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUnderlayNatGatewayWanInfoAndWanIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcLimitNamesToTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalCount

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// VpcLimitNamesToTypeSet

		VpcLimitNamesToTypeSet []*VpcLimitNamesToType `json:"VpcLimitNamesToTypeSet,omitempty" name:"VpcLimitNamesToTypeSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcLimitNamesToTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcLimitNamesToTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AdjustNfvGwRequest struct {
	*tchttp.BaseRequest

	// 网关id

	GwId *int64 `json:"GwId,omitempty" name:"GwId"`
	// 操作类型：调整网关规格/增加节点

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 节点规格

	Format *string `json:"Format,omitempty" name:"Format"`
	// 节点数量

	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 内网交换机组

	SwitchGroup *string `json:"SwitchGroup,omitempty" name:"SwitchGroup"`
	// 外网交换机组

	OuterSwitchGroup *string `json:"OuterSwitchGroup,omitempty" name:"OuterSwitchGroup"`
	// ipv6节点地址

	Ipv6NodeNetwork *string `json:"Ipv6NodeNetwork,omitempty" name:"Ipv6NodeNetwork"`
}

func (r *AdjustNfvGwRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AdjustNfvGwRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNatDetectIpRequest struct {
	*tchttp.BaseRequest

	// natgw探测IP。

	IpSet []*DetectIp `json:"IpSet,omitempty" name:"IpSet"`
}

func (r *CreateNatDetectIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNatDetectIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UNatGwWanIp struct {
	// 外网IP

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// 运营商

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
	// underlay natgateway的ID

	UnatgwId *uint64 `json:"UnatgwId,omitempty" name:"UnatgwId"`
	// underlay natgateway网关唯一ID

	UniqUnatgwId *string `json:"UniqUnatgwId,omitempty" name:"UniqUnatgwId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type Error struct {
	// 错误码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 错误内容

	Message *string `json:"Message,omitempty" name:"Message"`
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

type DescribeClbGatewayVersionsRequest struct {
	*tchttp.BaseRequest

	// GatewayName

	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`
}

func (r *DescribeClbGatewayVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClbGatewayVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnderlayNatGatewayRuleWanIpsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUnderlayNatGatewayRuleWanIpsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnderlayNatGatewayRuleWanIpsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNfvGwListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 网关详细信息数组

		GwLists []*GwLists `json:"GwLists,omitempty" name:"GwLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNfvGwListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNfvGwListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcEndPointResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 终端节点对象。

		EndPointSet []*EndPoint `json:"EndPointSet,omitempty" name:"EndPointSet"`
		// 符合查询条件的终端节点个数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcEndPointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcEndPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcPeerCrossRegionRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数据总量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeVpcPeerCrossRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcPeerCrossRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClbGatewayIfcfgRequest struct {
	*tchttp.BaseRequest

	// Ip

	Ip []*string `json:"Ip,omitempty" name:"Ip"`
}

func (r *GetClbGatewayIfcfgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClbGatewayIfcfgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateGatewayRequest struct {
	*tchttp.BaseRequest

	// Type

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// GatewayIp

	GatewayIp *string `json:"GatewayIp,omitempty" name:"GatewayIp"`
	// NewGatewayIp

	NewGatewayIp *string `json:"NewGatewayIp,omitempty" name:"NewGatewayIp"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *MigrateGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIpV6SubnetsRequest struct {
	*tchttp.BaseRequest

	// 网段类型，0非自研网段，1自研网段，目前只取0。

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// ipv6网段。

	Address *string `json:"Address,omitempty" name:"Address"`
	// 掩码。

	IntPrefix *uint64 `json:"IntPrefix,omitempty" name:"IntPrefix"`
	// 是否启用，0不启用，1启用，目前只取1。

	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
	// 裂解掩码，目前取值56。

	SplitIntPrefix *uint64 `json:"SplitIntPrefix,omitempty" name:"SplitIntPrefix"`
	// 地址类型，0：GUA， 1： ULA， 默认0。

	AddressType *uint64 `json:"AddressType,omitempty" name:"AddressType"`
	// 备注。

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateIpV6SubnetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIpV6SubnetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CarrierInfo struct {
	// 运营商ID

	CarrierId *string `json:"CarrierId,omitempty" name:"CarrierId"`
	// 运营商名称

	CarrierName *string `json:"CarrierName,omitempty" name:"CarrierName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type AddUnderlayNatGatewayRuleRipRequest struct {
	*tchttp.BaseRequest

	// underlay natgateway规则对应rip信息

	UnatgwRuleRips []*UnatgwRuleRipR `json:"UnatgwRuleRips,omitempty" name:"UnatgwRuleRips"`
}

func (r *AddUnderlayNatGatewayRuleRipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUnderlayNatGatewayRuleRipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcPeerCrossRegionRequest struct {
	*tchttp.BaseRequest

	// 本端地域

	SrcRegion *string `json:"SrcRegion,omitempty" name:"SrcRegion"`
	// 对端地域

	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`
}

func (r *CreateVpcPeerCrossRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcPeerCrossRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnGatewayDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		InternalHotHaIp *string `json:"InternalHotHaIp,omitempty" name:"InternalHotHaIp"`
		// 无

		UniqCdcId *string `json:"UniqCdcId,omitempty" name:"UniqCdcId"`
		// 无

		IspId *uint64 `json:"IspId,omitempty" name:"IspId"`
		// 无

		VpngwDomain []*string `json:"VpngwDomain,omitempty" name:"VpngwDomain"`
		// 无

		VpcSubnet *string `json:"VpcSubnet,omitempty" name:"VpcSubnet"`
		// 无

		Zone *string `json:"Zone,omitempty" name:"Zone"`
		// 无

		InternalHotHaMac *string `json:"InternalHotHaMac,omitempty" name:"InternalHotHaMac"`
		// 无

		VpcIntMask *uint64 `json:"VpcIntMask,omitempty" name:"VpcIntMask"`
		// 无

		SslvpnPort *uint64 `json:"SslvpnPort,omitempty" name:"SslvpnPort"`
		// 无

		BBlocked *uint64 `json:"BBlocked,omitempty" name:"BBlocked"`
		// 无

		VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
		// 无

		BAutoRenewals *uint64 `json:"BAutoRenewals,omitempty" name:"BAutoRenewals"`
		// 无

		InternalHotHaMask *string `json:"InternalHotHaMask,omitempty" name:"InternalHotHaMask"`
		// 无

		Quota *VpnQuota `json:"Quota,omitempty" name:"Quota"`
		// 无

		ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
		// 无

		VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
		// 无

		UniqVpngwId *string `json:"UniqVpngwId,omitempty" name:"UniqVpngwId"`
		// 无

		HotHaSubnet *string `json:"HotHaSubnet,omitempty" name:"HotHaSubnet"`
		// 无

		SlaLocalIp *string `json:"SlaLocalIp,omitempty" name:"SlaLocalIp"`
		// 无

		Name *string `json:"Name,omitempty" name:"Name"`
		// 无

		BNewAfc *uint64 `json:"BNewAfc,omitempty" name:"BNewAfc"`
		// 无

		BHotHa *uint64 `json:"BHotHa,omitempty" name:"BHotHa"`
		// 无

		VpcMask *string `json:"VpcMask,omitempty" name:"VpcMask"`
		// 无

		ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
		// 无

		WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
		// 无

		HotHaIp *string `json:"HotHaIp,omitempty" name:"HotHaIp"`
		// 无

		InternalHotHaSubnet *string `json:"InternalHotHaSubnet,omitempty" name:"InternalHotHaSubnet"`
		// 无

		UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
		// 无

		VpnType *uint64 `json:"VpnType,omitempty" name:"VpnType"`
		// 无

		HotHaMac *string `json:"HotHaMac,omitempty" name:"HotHaMac"`
		// 无

		VpcBDefault *uint64 `json:"VpcBDefault,omitempty" name:"VpcBDefault"`
		// 无

		State *uint64 `json:"State,omitempty" name:"State"`
		// 无

		DealId *string `json:"DealId,omitempty" name:"DealId"`
		// 无

		ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
		// 无

		HotHaMask *string `json:"HotHaMask,omitempty" name:"HotHaMask"`
		// 无

		Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
		// 无

		VpngwId *uint64 `json:"VpngwId,omitempty" name:"VpngwId"`
		// 无

		Owner *string `json:"Owner,omitempty" name:"Owner"`
		// 无

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpnGatewayDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnGatewayDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GwLists struct {
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 虚拟ip

	Vips []*Vips `json:"Vips,omitempty" name:"Vips"`
	// 可用区信息

	ZoneInfo []*string `json:"ZoneInfo,omitempty" name:"ZoneInfo"`
	// 网关模式

	GwMode *int64 `json:"GwMode,omitempty" name:"GwMode"`
	// 网关统一id

	UniqGwId *string `json:"UniqGwId,omitempty" name:"UniqGwId"`
	// 网关id

	GwId *int64 `json:"GwId,omitempty" name:"GwId"`
	// 网关命名

	GwName *string `json:"GwName,omitempty" name:"GwName"`
	// 节点数量

	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`
	// 网管类型

	GwType *string `json:"GwType,omitempty" name:"GwType"`
	// 操作状态

	OperateStatus *string `json:"OperateStatus,omitempty" name:"OperateStatus"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeSecurityGroupForCvmResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全组列表

		UsgSet []*VmSgInfoOuter `json:"UsgSet,omitempty" name:"UsgSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupForCvmResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupForCvmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductFailureMigrateTaskRsp struct {
	// 切换任务UUID

	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
	// 任务状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 任务进度

	Progress *int64 `json:"Progress,omitempty" name:"Progress"`
	// 执行信息或错误信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 详情

	Detail []*string `json:"Detail,omitempty" name:"Detail"`
	// 补充数据

	ExtendData []*string `json:"ExtendData,omitempty" name:"ExtendData"`
}

type WanIf struct {
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// eth0

	Eth0 *string `json:"Eth0,omitempty" name:"Eth0"`
	// gateway

	Gateway *string `json:"Gateway,omitempty" name:"Gateway"`
	// vipseg

	VipSeg []*string `json:"VipSeg,omitempty" name:"VipSeg"`
	// vlanid

	VlanId *string `json:"VlanId,omitempty" name:"VlanId"`
	// ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type DeleteServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductHealthStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// VPC集群健康状态详情

		Data *ProductHealthStateRsp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryProductHealthStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductHealthStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
	// VipBusiness

	VipBusiness *string `json:"VipBusiness,omitempty" name:"VipBusiness"`
	// ZoneId

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
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

	Vm *JnsGatewayServicePVm `json:"Vm,omitempty" name:"Vm"`
}

type VpcEndPointServiceUser struct {
	// AppId。

	Owner *uint64 `json:"Owner,omitempty" name:"Owner"`
	// Uin。

	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`
	// 描述信息。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 终端节点服务ID。

	EndPointServiceId *string `json:"EndPointServiceId,omitempty" name:"EndPointServiceId"`
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
	// UniqVpcGwIpVpcgwId

	UniqVpcGwIpVpcgwId *string `json:"UniqVpcGwIpVpcgwId,omitempty" name:"UniqVpcGwIpVpcgwId"`
	// cdc集群id

	UniqCdcId *string `json:"UniqCdcId,omitempty" name:"UniqCdcId"`
}

type Ifcfg struct {
	// IPADDR

	IPADDR *string `json:"IPADDR,omitempty" name:"IPADDR"`
	// GATEWAY

	GATEWAY *string `json:"GATEWAY,omitempty" name:"GATEWAY"`
	// NETMASK

	NETMASK *string `json:"NETMASK,omitempty" name:"NETMASK"`
	// IFNAME

	IFNAME *string `json:"IFNAME,omitempty" name:"IFNAME"`
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

	Route []*RouteService `json:"Route,omitempty" name:"Route"`
	// GroupId

	GroupId []*uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// AliveFlag

	AliveFlag *uint64 `json:"AliveFlag,omitempty" name:"AliveFlag"`
	// BVpcgwAttrFlag

	BVpcgwAttrFlag *uint64 `json:"BVpcgwAttrFlag,omitempty" name:"BVpcgwAttrFlag"`
	// Business

	Business *string `json:"Business,omitempty" name:"Business"`
	// GroupName

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// IpType

	IpType *uint64 `json:"IpType,omitempty" name:"IpType"`
	// RouteGroupFlag

	RouteGroupFlag *uint64 `json:"RouteGroupFlag,omitempty" name:"RouteGroupFlag"`
	// RouteGroupId

	RouteGroupId *uint64 `json:"RouteGroupId,omitempty" name:"RouteGroupId"`
	// RouteGroupVpcId

	RouteGroupVpcId *uint64 `json:"RouteGroupVpcId,omitempty" name:"RouteGroupVpcId"`
	// RsZoneVpcgwId

	RsZoneVpcgwId *uint64 `json:"RsZoneVpcgwId,omitempty" name:"RsZoneVpcgwId"`
	// SubnetId

	SubnetId *uint64 `json:"SubnetId,omitempty" name:"SubnetId"`
	// SubnetName

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// UniqRouteGroupId

	UniqRouteGroupId *string `json:"UniqRouteGroupId,omitempty" name:"UniqRouteGroupId"`
	// UniqSubnetId

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
	// VpcgwZoneId

	VpcgwZoneId *uint64 `json:"VpcgwZoneId,omitempty" name:"VpcgwZoneId"`
}

type AdjustNfvGwNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AdjustNfvGwNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AdjustNfvGwNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUnderlayNatGatewayRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误信息

		Error *string `json:"Error,omitempty" name:"Error"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUnderlayNatGatewayRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUnderlayNatGatewayRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceRoutesRequest struct {
	*tchttp.BaseRequest

	// 无

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// 无

	Pport *uint64 `json:"Pport,omitempty" name:"Pport"`
	// 无

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 无

	Pip *string `json:"Pip,omitempty" name:"Pip"`
}

func (r *DeleteServiceRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServiceRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VmSgInfoOuter struct {
	// 错误码

	ErrorCode *uint64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 错误信息

	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
	// vm 管理的sg信息

	UsgInfo []*VmSgInfo `json:"UsgInfo,omitempty" name:"UsgInfo"`
}

type CreateGatewayInstallationRequest struct {
	*tchttp.BaseRequest

	// GatewayName

	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`
	// Ips

	Ips []*string `json:"Ips,omitempty" name:"Ips"`
	// Vip

	Vip []*string `json:"Vip,omitempty" name:"Vip"`
	// Ver

	Ver *string `json:"Ver,omitempty" name:"Ver"`
	// IdcName

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// snat_ip

	SnatIpRange []*string `json:"SnatIpRange,omitempty" name:"SnatIpRange"`
	// JnsgwVip

	JnsgwVip []*string `json:"JnsgwVip,omitempty" name:"JnsgwVip"`
	// PvgwVip

	PvgwVip []*string `json:"PvgwVip,omitempty" name:"PvgwVip"`
	// OuterIp

	OuterIp []*string `json:"OuterIp,omitempty" name:"OuterIp"`
	// OuterVip

	OuterVip []*string `json:"OuterVip,omitempty" name:"OuterVip"`
	// DcgwVip

	DcgwVip *string `json:"DcgwVip,omitempty" name:"DcgwVip"`
	// SxgwVip

	SxgwVip *string `json:"SxgwVip,omitempty" name:"SxgwVip"`
}

func (r *CreateGatewayInstallationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGatewayInstallationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GwDataList struct {
	// 网关模板列表

	GwData []*GwData `json:"GwData,omitempty" name:"GwData"`
}

type DeleteUnderlayNatGatewayWanIpRequest struct {
	*tchttp.BaseRequest

	// 运营商IP地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 运营商

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
	// underlay natgateway的ID

	UnatgwId *uint64 `json:"UnatgwId,omitempty" name:"UnatgwId"`
}

func (r *DeleteUnderlayNatGatewayWanIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUnderlayNatGatewayWanIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnderlayNatGatewayWanIpsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUnderlayNatGatewayWanIpsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnderlayNatGatewayWanIpsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubnetInput struct {
	// AppId

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// Type

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// VPC实例ID。如66379

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// 子网的 `IPv4` `CIDR`。

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type CreateNatDetectIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNatDetectIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNatDetectIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNfvGwDetailRequest struct {
	*tchttp.BaseRequest

	// 网关id

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetNfvGwDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNfvGwDetailRequest) FromJsonString(s string) error {
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

type Vips struct {
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 命名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 值

	Value []*string `json:"Value,omitempty" name:"Value"`
}

type CreateServiceRequest struct {
	*tchttp.BaseRequest

	// 无

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// 无

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// 无

	SubnetId *uint64 `json:"SubnetId,omitempty" name:"SubnetId"`
	// 无

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// 无

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 无

	Routes []*ServiceRouteRoute `json:"Routes,omitempty" name:"Routes"`
}

func (r *CreateServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayIpsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalCount

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// GatewayVipSet

		GatewayVipSet []*string `json:"GatewayVipSet,omitempty" name:"GatewayVipSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayIpsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGatewayGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建的网关组对象的简要信息

		GatewayGroup *GatewayGroupBrief `json:"GatewayGroup,omitempty" name:"GatewayGroup"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGatewayGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGatewayGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDetectHostStateRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDetectHostStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDetectHostStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDetectTypeNamesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalCount

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// DetectTypeNameSet

		DetectTypeNameSet []*string `json:"DetectTypeNameSet,omitempty" name:"DetectTypeNameSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDetectTypeNamesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDetectTypeNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductFailureMigrateTaskDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务数据

		Data *AzProductFailureMigrateTaskDetailResponse `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryProductFailureMigrateTaskDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductFailureMigrateTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceAttributeRequest struct {
	*tchttp.BaseRequest

	// 无

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// 无

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 无

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// 无

	StickyFlag *uint64 `json:"StickyFlag,omitempty" name:"StickyFlag"`
}

func (r *ModifyServiceAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServiceAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MulticastGroup struct {
	// VPC唯一ID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 组播组地址。

	GroupAddr *string `json:"GroupAddr,omitempty" name:"GroupAddr"`
	// 过滤组播源地址。

	FilterAddr *string `json:"FilterAddr,omitempty" name:"FilterAddr"`
	// 过滤组播类型。

	FilterType *int64 `json:"FilterType,omitempty" name:"FilterType"`
	// 组播组限速pps值。

	Pps *int64 `json:"Pps,omitempty" name:"Pps"`
	// 组播组限速bps值。

	Bps *int64 `json:"Bps,omitempty" name:"Bps"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
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

type DeleteServiceRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误信息

		Error *Error `json:"Error,omitempty" name:"Error"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServiceRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRtbRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 路由表路由列表

		RtbRouteSet []*RtbRoute `json:"RtbRouteSet,omitempty" name:"RtbRouteSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRtbRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRtbRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAclAndItemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// acl规则信息列表

		AclAndItemSet []*AclEntry `json:"AclAndItemSet,omitempty" name:"AclAndItemSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAclAndItemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAclAndItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PipWanDesc struct {
	// Pip

	Pip *string `json:"Pip,omitempty" name:"Pip"`
	// WanIp

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// Description

	Description *string `json:"Description,omitempty" name:"Description"`
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

type DeleteUnderlayNatGatewayRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误信息

		Error *string `json:"Error,omitempty" name:"Error"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUnderlayNatGatewayRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUnderlayNatGatewayRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnderlayNatGatewayRuleRequest struct {
	*tchttp.BaseRequest

	// 分页起始位置，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页返回的条目数，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 查询的字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 查询说明

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeUnderlayNatGatewayRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnderlayNatGatewayRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGatewayCLBInstallationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// EventId

		EventId *string `json:"EventId,omitempty" name:"EventId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGatewayCLBInstallationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGatewayCLBInstallationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUnderlayNatGatewayRuleRipRequest struct {
	*tchttp.BaseRequest

	// underlay natgateway规则详细信息

	UnatgwRuleRips []*UNatGwRuleRip `json:"UnatgwRuleRips,omitempty" name:"UnatgwRuleRips"`
}

func (r *DeleteUnderlayNatGatewayRuleRipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUnderlayNatGatewayRuleRipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServicesRequest struct {
	*tchttp.BaseRequest

	// 无

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 无

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 无

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// 无

	UniqVpcgwId *string `json:"UniqVpcgwId,omitempty" name:"UniqVpcgwId"`
	// 无

	VpcGwIp *string `json:"VpcGwIp,omitempty" name:"VpcGwIp"`
	// 无

	Pip *string `json:"Pip,omitempty" name:"Pip"`
	// 无

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 无

	IsFetchRule *bool `json:"IsFetchRule,omitempty" name:"IsFetchRule"`
}

func (r *DescribeServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VPCIPv6Section struct {
	// 地域名称

	Region *string `json:"Region,omitempty" name:"Region"`
	// 网段

	Address *string `json:"Address,omitempty" name:"Address"`
	// 总数量

	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
	// 剩余总数量

	LastTotalNum *uint64 `json:"LastTotalNum,omitempty" name:"LastTotalNum"`
	// 备注

	Description *string `json:"Description,omitempty" name:"Description"`
	// 掩码

	IntPrefix *uint64 `json:"IntPrefix,omitempty" name:"IntPrefix"`
	// 地址类型，0: GUA; 1: ULA; 默认0

	AddressType *uint64 `json:"AddressType,omitempty" name:"AddressType"`
}

type GetSwitchGroupRequest struct {
	*tchttp.BaseRequest

	// 网关类型

	GwType *string `json:"GwType,omitempty" name:"GwType"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *GetSwitchGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSwitchGroupRequest) FromJsonString(s string) error {
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

type VpcServiceRoule struct {
	// VpcId

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// VpcId

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// Proto

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// Vport

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

type DescribeGatewayInstallStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalCount

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// GatewayInstallStateSet

		GatewayInstallStateSet []*GatewayInstallState `json:"GatewayInstallStateSet,omitempty" name:"GatewayInstallStateSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayInstallStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayInstallStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRtbRoutesRequest struct {
	*tchttp.BaseRequest

	// 过滤器列表

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// VPC实例ID。如65967

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方式

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeRtbRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRtbRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnGatewaysRequest struct {
	*tchttp.BaseRequest

	// 无

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	State []*uint64 `json:"State,omitempty" name:"State"`
	// 无

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// 无

	UniqVpngwId *string `json:"UniqVpngwId,omitempty" name:"UniqVpngwId"`
	// 无

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 无

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

func (r *DescribeVpnGatewaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnGatewaysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProtectPolicyIp struct {
	// Rip。

	Rip *string `json:"Rip,omitempty" name:"Rip"`
	// PPS限制值（个/秒）。

	PktLimit *uint64 `json:"PktLimit,omitempty" name:"PktLimit"`
	// 宽带限制值（Mbps）

	TcLimit *uint64 `json:"TcLimit,omitempty" name:"TcLimit"`
	// 并发连接数限制值（个）。

	ConnLimit *uint64 `json:"ConnLimit,omitempty" name:"ConnLimit"`
	// 新增连接限制值（个）。

	InconnLimit *uint64 `json:"InconnLimit,omitempty" name:"InconnLimit"`
	// 是否开启过载保护，0不开启，1开启，默认开启。

	Switch *uint64 `json:"Switch,omitempty" name:"Switch"`
}

type DescribeVpcEndPointServiceRequest struct {
	*tchttp.BaseRequest

	// 单页返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>AppId - String - （过滤条件）用户的AppId</li>
	// <li>VpcId - String - （过滤条件）用户的vpc唯一id。</li>
	// <li>ServiceUniqInstanceId - String - （过滤条件）后端服务的唯一id。</li>
	// <li>ServiceName - String - （过滤条件）终端节点服务名称。</li>
	// <li>EndPointServiceId - String - （过滤条件）终端节点服务唯一id。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVpcEndPointServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcEndPointServiceRequest) FromJsonString(s string) error {
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

	VrouterId *int64 `json:"VrouterId,omitempty" name:"VrouterId"`
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

type DescribeUnderlayNatGatewayRuleWanIpsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 信息条目数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// underlay natgateway规则详细信息

		UNatGwRuleWipSet []*UNatGwRuleWip `json:"UNatGwRuleWipSet,omitempty" name:"UNatGwRuleWipSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUnderlayNatGatewayRuleWanIpsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnderlayNatGatewayRuleWanIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcEndPointRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 单页返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// <li> AppId- String - （过滤条件）用户APPID。</li>
	// <li> EndPointId- String - （过滤条件）终端节点ID。</li>
	// <li> EndPointName- String - （过滤条件）终端节点名称。</li>
	// <li> VpcId- String - （过滤条件）VPC唯一ID。</li>
	// <li> EndPointServiceId- String - （过滤条件）终端节点服务ID。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVpcEndPointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcEndPointRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type RouteRule struct {
	// VpcId数字ID。

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// Vip。

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// Type。

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// Proto。

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// Vport。

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 集群组数字ID。

	VpcgwId *int64 `json:"VpcgwId,omitempty" name:"VpcgwId"`
}

type DescribeGatewayInstallConfRequest struct {
	*tchttp.BaseRequest

	// 环境信息

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// GwType

	GwType *string `json:"GwType,omitempty" name:"GwType"`
}

func (r *DescribeGatewayInstallConfRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayInstallConfRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnderlayNatGatewayWanInfoAndWanIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// underlay natgateway信息条目数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// underlay natgateway详细信息

		UNatGwWanInfoAndWanIpSet []*UNatGwWanInfoAndWanIp `json:"UNatGwWanInfoAndWanIpSet,omitempty" name:"UNatGwWanInfoAndWanIpSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUnderlayNatGatewayWanInfoAndWanIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnderlayNatGatewayWanInfoAndWanIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnatgwWip struct {
	// 运营商

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 外网IP

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// underlay natgateway的ID

	UnatgwId *uint64 `json:"UnatgwId,omitempty" name:"UnatgwId"`
}

type DeleteClbGatewayInstallRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClbGatewayInstallRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClbGatewayInstallRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDetectHostRequest struct {
	*tchttp.BaseRequest

	// Ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *DeleteDetectHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDetectHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductStateInfoRequest struct {
	*tchttp.BaseRequest

	// 产品名称

	Product *string `json:"Product,omitempty" name:"Product"`
	// 特殊参数

	Params []*string `json:"Params,omitempty" name:"Params"`
}

func (r *QueryProductStateInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductStateInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeNfvGwNodeRequest struct {
	*tchttp.BaseRequest

	// 软件版本

	GwVersion *string `json:"GwVersion,omitempty" name:"GwVersion"`
	// 替换节点id

	InstanceUuids []*string `json:"InstanceUuids,omitempty" name:"InstanceUuids"`
	// 集群id

	ClusterId *uint64 `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *UpgradeNfvGwNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeNfvGwNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubnetNameID struct {
	// 子网名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 子网实例ID。如3706

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 子网唯一ID。如subnet-3cxo1o9e

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
}

type CreateJnsGatewayServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// JnsGatewayService

		JnsGatewayService *JnsGatewayServiceP `json:"JnsGatewayService,omitempty" name:"JnsGatewayService"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateJnsGatewayServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateJnsGatewayServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAclListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// acl总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// acl 列表

		AclSet []*AclInfo `json:"AclSet,omitempty" name:"AclSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAclListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAclListResponse) FromJsonString(s string) error {
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

type Vm struct {
	// AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// EipType

	EipType *uint64 `json:"EipType,omitempty" name:"EipType"`
	// SxgwIp

	SxgwIp *string `json:"SxgwIp,omitempty" name:"SxgwIp"`
	// Eip

	Eip *string `json:"Eip,omitempty" name:"Eip"`
	// LanOutLimit

	LanOutLimit *float64 `json:"LanOutLimit,omitempty" name:"LanOutLimit"`
	// EipOutLimit

	EipOutLimit *float64 `json:"EipOutLimit,omitempty" name:"EipOutLimit"`
	// VmEip

	VmEip *string `json:"VmEip,omitempty" name:"VmEip"`
	// EvpnOverlayMac

	EvpnOverlayMac *string `json:"EvpnOverlayMac,omitempty" name:"EvpnOverlayMac"`
	// BAlgSip

	BAlgSip *int64 `json:"BAlgSip,omitempty" name:"BAlgSip"`
	// AliasVpcId

	AliasVpcId *uint64 `json:"AliasVpcId,omitempty" name:"AliasVpcId"`
	// UniqCdcId

	UniqCdcId *string `json:"UniqCdcId,omitempty" name:"UniqCdcId"`
	// BAlgFtp

	BAlgFtp *uint64 `json:"BAlgFtp,omitempty" name:"BAlgFtp"`
	// LanInLimit

	LanInLimit *float64 `json:"LanInLimit,omitempty" name:"LanInLimit"`
	// BCdc

	BCdc *uint64 `json:"BCdc,omitempty" name:"BCdc"`
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

type DeleteUnderlayNatGatewayWanIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误信息

		Error *string `json:"Error,omitempty" name:"Error"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUnderlayNatGatewayWanIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUnderlayNatGatewayWanIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 无

		ServiceSet []*Service `json:"ServiceSet,omitempty" name:"ServiceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AclInfo struct {
	// VpcId

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// AclId

	AclId *uint64 `json:"AclId,omitempty" name:"AclId"`
	// UniqVpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// SubnetNum

	SubnetNum *uint64 `json:"SubnetNum,omitempty" name:"SubnetNum"`
	// VpcSubnet

	VpcSubnet *string `json:"VpcSubnet,omitempty" name:"VpcSubnet"`
	// VpcMask

	VpcMask *string `json:"VpcMask,omitempty" name:"VpcMask"`
	// VpcName

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// VpcBDefault

	VpcBDefault *uint64 `json:"VpcBDefault,omitempty" name:"VpcBDefault"`
	// VpcIntMask

	VpcIntMask *uint64 `json:"VpcIntMask,omitempty" name:"VpcIntMask"`
	// Subnet

	Subnet []*AclSubNetInfo `json:"Subnet,omitempty" name:"Subnet"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// UniqAclId

	UniqAclId *string `json:"UniqAclId,omitempty" name:"UniqAclId"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeVmListRequest struct {
	*tchttp.BaseRequest

	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序列表

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件，参数不支持同时指定VpcIds和Filters。
	// <li>AppId - String - （过滤条件）用户的AppId</li>
	// <li>UniqVpcId - String - （过滤条件）用户的vpc唯一id。</li>
	// <li>Subnet - String - （过滤条件）cvm子网。</li>
	// <li>VmIp- String - （过滤条件）cvm ip。</li>
	// <li>HostIp - String - （过滤条件）cvm宿主机ip。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 是否获取ipv6信息

	BExtendIpv6Info *bool `json:"BExtendIpv6Info,omitempty" name:"BExtendIpv6Info"`
}

func (r *DescribeVmListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVmListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatInfo struct {
	// WanInLimit

	WanInLimit *uint64 `json:"WanInLimit,omitempty" name:"WanInLimit"`
	// NatType

	NatType *string `json:"NatType,omitempty" name:"NatType"`
	// BLearnTsvIp

	BLearnTsvIp *uint64 `json:"BLearnTsvIp,omitempty" name:"BLearnTsvIp"`
	// BLearnTsvIp

	BExclusive *uint64 `json:"BExclusive,omitempty" name:"BExclusive"`
	// BLearnHostIp

	BLearnHostIp *uint64 `json:"BLearnHostIp,omitempty" name:"BLearnHostIp"`
	// GwType

	GwType *uint64 `json:"GwType,omitempty" name:"GwType"`
	// VpcId

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// UniqVpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// Connections

	Connections *uint64 `json:"Connections,omitempty" name:"Connections"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// UniqNatId

	UniqNatId *string `json:"UniqNatId,omitempty" name:"UniqNatId"`
	// NatId

	NatId *uint64 `json:"NatId,omitempty" name:"NatId"`
	// Zone

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// ZoneId

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// State

	State *uint64 `json:"State,omitempty" name:"State"`
	// BDeleted

	BDeleted *uint64 `json:"BDeleted,omitempty" name:"BDeleted"`
	// OwedWanOutLimit

	OwedWanOutLimit *uint64 `json:"OwedWanOutLimit,omitempty" name:"OwedWanOutLimit"`
	// GwId

	GwId *uint64 `json:"GwId,omitempty" name:"GwId"`
	// HealthStatus

	HealthStatus *uint64 `json:"HealthStatus,omitempty" name:"HealthStatus"`
	// DetectGwIp

	DetectGwIp *string `json:"DetectGwIp,omitempty" name:"DetectGwIp"`
	// GwIp

	GwIp *string `json:"GwIp,omitempty" name:"GwIp"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// LastUpdateConnectionsTime

	LastUpdateConnectionsTime *string `json:"LastUpdateConnectionsTime,omitempty" name:"LastUpdateConnectionsTime"`
	// OwedFlag

	OwedFlag *uint64 `json:"OwedFlag,omitempty" name:"OwedFlag"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// LastMaxConnections

	LastMaxConnections *uint64 `json:"LastMaxConnections,omitempty" name:"LastMaxConnections"`
	// WanOutLimit

	WanOutLimit *uint64 `json:"WanOutLimit,omitempty" name:"WanOutLimit"`
	// Ip

	Ip []*NatIpInfo `json:"Ip,omitempty" name:"Ip"`
}

type AdjustNfvGwResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误信息

		Error *Error `json:"Error,omitempty" name:"Error"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AdjustNfvGwResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AdjustNfvGwResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcLimitNamesToTypeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVpcLimitNamesToTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcLimitNamesToTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMulticastGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMulticastGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMulticastGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcLimitNamesToType struct {
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Types

	Types []*uint64 `json:"Types,omitempty" name:"Types"`
}

type DeleteUnderlayNatGatewayRuleRipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUnderlayNatGatewayRuleRipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUnderlayNatGatewayRuleRipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDetectRouteStateRequest struct {
	*tchttp.BaseRequest

	// 探测机内网IP。

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 重置操作时传-1即可。

	State *int64 `json:"State,omitempty" name:"State"`
}

func (r *ModifyDetectRouteStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDetectRouteStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDetectInstallRecord struct {
	// EventId

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// Ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type L7RuleInfo struct {
	// location list

	LocationSet []*LocationInfo `json:"LocationSet,omitempty" name:"LocationSet"`
	// virtual service

	VirtualService *VirtualServiceInfo `json:"VirtualService,omitempty" name:"VirtualService"`
}

type ProductStateInfoRsp struct {
	// 产品名称，取值VPCGW

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集群信息列表

	ClusterList []*ClusterState `json:"ClusterList,omitempty" name:"ClusterList"`
	// 状态最后上报时间（更新时间）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type DescribeVpnConnsRequest struct {
	*tchttp.BaseRequest

	// 无

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	UsrgwId *uint64 `json:"UsrgwId,omitempty" name:"UsrgwId"`
	// 无

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// 无

	UniqVpngwId *string `json:"UniqVpngwId,omitempty" name:"UniqVpngwId"`
	// 无

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 无

	UniqVpnconnId *string `json:"UniqVpnconnId,omitempty" name:"UniqVpnconnId"`
}

func (r *DescribeVpnConnsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnConnsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EndPointService struct {
	// 终端节点服务ID。

	EndPointServiceId *string `json:"EndPointServiceId,omitempty" name:"EndPointServiceId"`
	// VPCID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// APPID。

	ServiceOwner *string `json:"ServiceOwner,omitempty" name:"ServiceOwner"`
	// 终端节点服务名称。

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 后端服务的VIP。

	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`
	// 后端服务的ID，比如lb-xxx。

	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" name:"ServiceInstanceId"`
	// 是否自动接受。

	AutoAcceptFlag *bool `json:"AutoAcceptFlag,omitempty" name:"AutoAcceptFlag"`
	// 关联的终端节点个数。

	EndPointCount *uint64 `json:"EndPointCount,omitempty" name:"EndPointCount"`
	// 终端节点对象数组。

	EndPointSet []*EndPoint `json:"EndPointSet,omitempty" name:"EndPointSet"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 终端节点服务白名单个数。

	WhiteListCount *uint64 `json:"WhiteListCount,omitempty" name:"WhiteListCount"`
	// 整形终端节点服务id

	IntEndPointServiceIPv6Id *int64 `json:"IntEndPointServiceIPv6Id,omitempty" name:"IntEndPointServiceIPv6Id"`
}

type DeleteJnsGatewayServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteJnsGatewayServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteJnsGatewayServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMulticastGroupsRequest struct {
	*tchttp.BaseRequest

	// VPC数字ID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 组播组地址。

	GroupAddr *string `json:"GroupAddr,omitempty" name:"GroupAddr"`
}

func (r *DescribeMulticastGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMulticastGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNfvGwFormatsRequest struct {
	*tchttp.BaseRequest

	// 网关类型

	GwRole *string `json:"GwRole,omitempty" name:"GwRole"`
	// 实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

func (r *GetNfvGwFormatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNfvGwFormatsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebootNfvGwLogStepResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RebootNfvGwLogStepResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootNfvGwLogStepResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// sg列表

		SgSet []*SgInfo `json:"SgSet,omitempty" name:"SgSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchGatewayRequest struct {
	*tchttp.BaseRequest

	// 无

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 无

	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

func (r *SwitchGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnGatewaysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 无

		VpnGatewaySet []*VpnGateway `json:"VpnGatewaySet,omitempty" name:"VpnGatewaySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpnGatewaysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnGatewaysResponse) FromJsonString(s string) error {
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
	// 0:无；24：铜；25：银；26：金

	EniQos *int64 `json:"EniQos,omitempty" name:"EniQos"`
}

type AddServiceRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddServiceRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddServiceRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUnderlayNatGatewayRuleRequest struct {
	*tchttp.BaseRequest

	// underlay natgateway网关的ID

	UnatgwId *uint64 `json:"UnatgwId,omitempty" name:"UnatgwId"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateUnderlayNatGatewayRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUnderlayNatGatewayRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayInstallStateRequest struct {
	*tchttp.BaseRequest

	// Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// EventId

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// GateWayName

	GateWayName *string `json:"GateWayName,omitempty" name:"GateWayName"`
	// InnerIp

	InnerIp *string `json:"InnerIp,omitempty" name:"InnerIp"`
}

func (r *DescribeGatewayInstallStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayInstallStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnderlayNatGatewayRuleRipsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 详细信息条目数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// underlay natgateway规则详细信息

		UNatGwRuleRipSet []*UNatGwRuleRip `json:"UNatGwRuleRipSet,omitempty" name:"UNatGwRuleRipSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUnderlayNatGatewayRuleRipsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnderlayNatGatewayRuleRipsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDetectHostWeightRequest struct {
	*tchttp.BaseRequest

	// Ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// Weight

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
}

func (r *ModifyDetectHostWeightRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDetectHostWeightRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AzMigrateDetailDO struct {
	// 切换步骤名称

	StepName *string `json:"StepName,omitempty" name:"StepName"`
	// 产品状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 切换信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 开始时间

	StartAt *string `json:"StartAt,omitempty" name:"StartAt"`
	// 结束时间

	EndAt *string `json:"EndAt,omitempty" name:"EndAt"`
}

type DescribeGatewaysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 网关信息集合

		GatewaySet []*GatewayInfo `json:"GatewaySet,omitempty" name:"GatewaySet"`
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

type GetNfvCvmInstanceTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// NFV母机类型列表

		NfvCvmInstanceTypeSet []*NfvCvmInstanceType `json:"NfvCvmInstanceTypeSet,omitempty" name:"NfvCvmInstanceTypeSet"`
		// NFV母机类型数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNfvCvmInstanceTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNfvCvmInstanceTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNfvGwRequest struct {
	*tchttp.BaseRequest

	// 网关id

	GwId *int64 `json:"GwId,omitempty" name:"GwId"`
}

func (r *DeleteNfvGwRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNfvGwRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNfvGwNodeDetailRequest struct {
	*tchttp.BaseRequest

	// 网关id

	GwId *int64 `json:"GwId,omitempty" name:"GwId"`
}

func (r *GetNfvGwNodeDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNfvGwNodeDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AzExtendDataDO struct {
	// 数据复制成功

	DataReplicationSuccess *int64 `json:"DataReplicationSuccess,omitempty" name:"DataReplicationSuccess"`
	// 数据复制失败

	DataReplicationFailed *int64 `json:"DataReplicationFailed,omitempty" name:"DataReplicationFailed"`
}

type DescribeCvmForSecurityGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// vm 列表

		VmSet []*SgVmInfo `json:"VmSet,omitempty" name:"VmSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCvmForSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCvmForSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcEndPointServiceWhiteListRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 单页返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// <li> AppId String - （过滤条件）用户APPID。</li>
	// <li> EndPointServiceId String - （过滤条件）终端节点服务唯一ID。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVpcEndPointServiceWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcEndPointServiceWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNfvGwNodeDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点详细信息数组

		GwNodes []*GwNodes `json:"GwNodes,omitempty" name:"GwNodes"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNfvGwNodeDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNfvGwNodeDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddUnderlayNatGatewayRuleRipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Error

		Error *string `json:"Error,omitempty" name:"Error"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUnderlayNatGatewayRuleRipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUnderlayNatGatewayRuleRipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条目数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// vpc列表

		VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EndPoint struct {
	// 终端节点ID。

	EndPointId *string `json:"EndPointId,omitempty" name:"EndPointId"`
	// VPCID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// APPID。

	EndPointOwner *string `json:"EndPointOwner,omitempty" name:"EndPointOwner"`
	// 终端节点名称。

	EndPointName *string `json:"EndPointName,omitempty" name:"EndPointName"`
	// 终端节点服务的VPCID。

	ServiceVpcId *string `json:"ServiceVpcId,omitempty" name:"ServiceVpcId"`
	// 终端节点服务的VIP。

	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`
	// 终端节点服务的ID。

	EndPointServiceId *string `json:"EndPointServiceId,omitempty" name:"EndPointServiceId"`
	// 终端节点的VIP。

	EndPointVip *string `json:"EndPointVip,omitempty" name:"EndPointVip"`
	// 终端节点状态，ACTIVE：可用，PENDING：待接受，ACCEPTING：接受中，REJECTED：已拒绝，FAILED：失败。

	State *string `json:"State,omitempty" name:"State"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 终端节点绑定的安全组实例列表。

	GroupSet []*SgInfo `json:"GroupSet,omitempty" name:"GroupSet"`
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

type UNatGwWanInfoAndWanIp struct {
	// 运营商

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// underlay natgateway的ID

	UnatgwId *uint64 `json:"UnatgwId,omitempty" name:"UnatgwId"`
	// underlay natgateway网关详细信息

	UnatgwInfo *UnatgwInfo `json:"UnatgwInfo,omitempty" name:"UnatgwInfo"`
	// underlay natgateway网关唯一ID

	UniqUnatgwId *string `json:"UniqUnatgwId,omitempty" name:"UniqUnatgwId"`
	// 外网IP

	WanIp []*string `json:"WanIp,omitempty" name:"WanIp"`
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
	// UniqCdcId

	UniqCdcId *string `json:"UniqCdcId,omitempty" name:"UniqCdcId"`
	// IspId

	IspId *uint64 `json:"IspId,omitempty" name:"IspId"`
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

type DescribeDetectPkgVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalCount

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// DetectPkgVersionSet

		DetectPkgVersionSet *string `json:"DetectPkgVersionSet,omitempty" name:"DetectPkgVersionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDetectPkgVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDetectPkgVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInterfacesExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 弹性网卡列表

		NetworkInterfaceSet []*NetworkInterface `json:"NetworkInterfaceSet,omitempty" name:"NetworkInterfaceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkInterfacesExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkInterfacesExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateDirectConnectGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误码，0为成功，非0为失败

		ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`
		// 异步任务taskid

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigrateDirectConnectGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateDirectConnectGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	Counts *SgVmCount `json:"Counts,omitempty" name:"Counts"`
}

type ModifyProtectPolicyIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProtectPolicyIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProtectPolicyIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HaVip struct {
	// 用户AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// `HAVIP`的`ID`，是`HAVIP`的唯一标识。

	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`
	// `HAVIP`名称。

	HaVipName *string `json:"HaVipName,omitempty" name:"HaVipName"`
	// 虚拟IP地址。

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// `HAVIP`所在私有网络`ID`。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// `HAVIP`所在子网`ID`。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// `HAVIP`关联弹性网卡`ID`。

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	// 被绑定的实例`ID`。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 绑定`EIP`。

	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`
	// 状态：
	// <li>`AVAILABLE`：运行中</li>
	// <li>`UNBIND`：未绑定</li>

	State *string `json:"State,omitempty" name:"State"`
	// 创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 使用havip的业务标识。

	Business *string `json:"Business,omitempty" name:"Business"`
}

type VmSet struct {
	// AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// VpcId

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
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

type DescribeBasicNetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// VPC0网段列表

		BasicNetSet []*BasicSubnetInfo `json:"BasicNetSet,omitempty" name:"BasicNetSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBasicNetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBasicNetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVmResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// vpc子机列表

		VmSet []*Vm `json:"VmSet,omitempty" name:"VmSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVmResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDetectInstallStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalCount

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// DetectInstallStateSet

		DetectInstallStateSet []*DetectInstallState `json:"DetectInstallStateSet,omitempty" name:"DetectInstallStateSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDetectInstallStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDetectInstallStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type RollbackNfvGwLogStepRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 步骤名

	Step *string `json:"Step,omitempty" name:"Step"`
}

func (r *RollbackNfvGwLogStepRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RollbackNfvGwLogStepRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayInstallConfResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网关部署配置列表

		GwData []*GwDataList `json:"GwData,omitempty" name:"GwData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayInstallConfResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayInstallConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpV6SubnetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 网段详情

		VPCIPV6SectionSet []*VPCIPv6Section `json:"VPCIPV6SectionSet,omitempty" name:"VPCIPV6SectionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpV6SubnetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpV6SubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnderlayNatGatewayRuleRipsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUnderlayNatGatewayRuleRipsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnderlayNatGatewayRuleRipsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVmRequest struct {
	*tchttp.BaseRequest

	// VmInfo

	VmSet []*VmConf `json:"VmSet,omitempty" name:"VmSet"`
}

func (r *ModifyVmRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVmRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProceedNfvGwLogStepResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProceedNfvGwLogStepResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProceedNfvGwLogStepResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcPeerCrossRegionRequest struct {
	*tchttp.BaseRequest

	// 本端地域

	SrcRegion *string `json:"SrcRegion,omitempty" name:"SrcRegion"`
	// 对端地域

	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`
}

func (r *DeleteVpcPeerCrossRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcPeerCrossRegionRequest) FromJsonString(s string) error {
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

type DeleteUnderlayNatGatewayWanInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误信息

		Error *Error `json:"Error,omitempty" name:"Error"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUnderlayNatGatewayWanInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUnderlayNatGatewayWanInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAclListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序类型

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 过滤器列表

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAclListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAclListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectGatewaysExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 专线网关列表

		DirectConnectGatewaySet []*DirectConnectGateway `json:"DirectConnectGatewaySet,omitempty" name:"DirectConnectGatewaySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectGatewaysExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectGatewaysExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcPeerCrossRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域对总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 有效的地域对。

		RegionPairSet []*RegionPair `json:"RegionPairSet,omitempty" name:"RegionPairSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcPeerCrossRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcPeerCrossRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateDirectConnectGatewayRequest struct {
	*tchttp.BaseRequest

	// 要迁移的专线网关id

	UniqVpgId *string `json:"UniqVpgId,omitempty" name:"UniqVpgId"`
	// 目标集群vip

	GwIp *string `json:"GwIp,omitempty" name:"GwIp"`
}

func (r *MigrateDirectConnectGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateDirectConnectGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcPeersRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方式

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 过滤器列表

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVpcPeersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcPeersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGatewayGroupAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGatewayGroupAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGatewayGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServiceAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Logs struct {
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 步骤

	Step *string `json:"Step,omitempty" name:"Step"`
	// 步骤名

	StepName *string `json:"StepName,omitempty" name:"StepName"`
	// 详情

	Detail []*Detail `json:"Detail,omitempty" name:"Detail"`
	// 任务id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务类型

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
}

type MigrateServiceRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigrateServiceRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateServiceRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UNatGwRuleWip struct {
	// 外网IP

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// underlay natgateway规则的ID

	UnatgwRuleId *uint64 `json:"UnatgwRuleId,omitempty" name:"UnatgwRuleId"`
	// 运营商

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
	// underlay natgateway规则唯一ID

	UniqUnatgwRuleId *string `json:"UniqUnatgwRuleId,omitempty" name:"UniqUnatgwRuleId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeDetectHostStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// DetectPkgVersionSet

		DetectPkgVersionSet []*DetectPkgVersion `json:"DetectPkgVersionSet,omitempty" name:"DetectPkgVersionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDetectHostStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDetectHostStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGatewaysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// nat 网关数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// nat 网关列表

		NatSet []*NatInfo `json:"NatSet,omitempty" name:"NatSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatGatewaysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatGatewaysResponse) FromJsonString(s string) error {
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

	BBlocked *uint64 `json:"BBlocked,omitempty" name:"BBlocked"`
	// BAutoApply

	BAutoApply *uint64 `json:"BAutoApply,omitempty" name:"BAutoApply"`
}

type NfvCvmInstanceType struct {
	// 母机类型名称

	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
	// 母机类型描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DescribeDetectInstallStateRequest struct {
	*tchttp.BaseRequest

	// Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// EventId

	EventId []*string `json:"EventId,omitempty" name:"EventId"`
	// GroupId

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// AzName

	AzName *string `json:"AzName,omitempty" name:"AzName"`
	// Status

	Status []*uint64 `json:"Status,omitempty" name:"Status"`
	// HostIp

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
}

func (r *DescribeDetectInstallStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDetectInstallStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayTypeNamesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalCount

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// GateWayTypeNameSet

		GateWayTypeNameSet []*string `json:"GateWayTypeNameSet,omitempty" name:"GateWayTypeNameSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayTypeNamesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayTypeNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVmResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVmResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVmResponse) FromJsonString(s string) error {
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

type DeleteClbGatewayInstallRecordRequest struct {
	*tchttp.BaseRequest

	// EventIds

	EventIds []*string `json:"EventIds,omitempty" name:"EventIds"`
}

func (r *DeleteClbGatewayInstallRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClbGatewayInstallRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDetectInstallRecordsRequest struct {
	*tchttp.BaseRequest

	// Records

	Records []*DeleteDetectInstallRecord `json:"Records,omitempty" name:"Records"`
}

func (r *DeleteDetectInstallRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDetectInstallRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUnderlayNatGatewayWanInfoAndWanIpAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUnderlayNatGatewayWanInfoAndWanIpAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUnderlayNatGatewayWanInfoAndWanIpAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIpV6SubnetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIpV6SubnetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIpV6SubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayNodesRequest struct {
	*tchttp.BaseRequest

	// IdcName

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// GatewayType

	GatewayType *string `json:"GatewayType,omitempty" name:"GatewayType"`
}

func (r *DescribeGatewayNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClbGatewayIfcfgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Ifcfg

		Ifcfg []*Ifcfg `json:"Ifcfg,omitempty" name:"Ifcfg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClbGatewayIfcfgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClbGatewayIfcfgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type DeleteJnsGatewayServiceRequest struct {
	*tchttp.BaseRequest

	// UniqVpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// Proto

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// Pip

	Pip *string `json:"Pip,omitempty" name:"Pip"`
	// Pport

	Pport *uint64 `json:"Pport,omitempty" name:"Pport"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

func (r *DeleteJnsGatewayServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteJnsGatewayServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnderlayNatGatewayWanInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// underlay natgateway信息条目数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// underlay natgateway详细信息

		UNatGwWanInfoSet []*UNatGwWanInfo `json:"UNatGwWanInfoSet,omitempty" name:"UNatGwWanInfoSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUnderlayNatGatewayWanInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnderlayNatGatewayWanInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserGatewaysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 无

		UserGatewaySet []*UserGateway `json:"UserGatewaySet,omitempty" name:"UserGatewaySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserGatewaysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type CreateBasicNetsRequest struct {
	*tchttp.BaseRequest

	// vpc0网段新建信息列表

	BasicNets []*BasicSubnetInfo `json:"BasicNets,omitempty" name:"BasicNets"`
}

func (r *CreateBasicNetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBasicNetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGatewayInstallRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGatewayInstallRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGatewayInstallRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcPeerCrossRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcPeerCrossRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcPeerCrossRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGatewayAttributeRequest struct {
	*tchttp.BaseRequest

	// 过滤参数EnableNat

	EnableNat *string `json:"EnableNat,omitempty" name:"EnableNat"`
	// 网关集群类型(0-pvgw 1-vpcgw 2-jnsgw 3-natgw 4-mcgw 5-dcgw 6-pcgw 7-xgw 8-sxgw 9-underly-natgw)

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 网关集群Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 新的网关节点IP列表

	Rips *string `json:"Rips,omitempty" name:"Rips"`
	// 新的网关集群名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 当网关集群类型为pvgw和mcgw时，可以指定网关集群Vip作为查询条件

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 新的权重，目前只有natgw支持该字段的修改

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
	// 可用区ID。

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区名称。

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 网关类型，0:ko;1:dpdk。

	VpcgwType *int64 `json:"VpcgwType,omitempty" name:"VpcgwType"`
}

func (r *ModifyGatewayAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGatewayAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcPeerCrossRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcPeerCrossRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcPeerCrossRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalCount

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// GateWayVersionSet

		GateWayVersionSet []*string `json:"GateWayVersionSet,omitempty" name:"GateWayVersionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductHealthStateRsp struct {
	// 产品名

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 错误信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 健康状态指标列表

	Metrics []*Metric `json:"Metrics,omitempty" name:"Metrics"`
}

type ForwardStrategyInfo struct {
	// detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// type

	Type *string `json:"Type,omitempty" name:"Type"`
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

type AddServiceRoutesRequest struct {
	*tchttp.BaseRequest

	// 无

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// 无

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// 无

	ServiceRoutes []*ServiceRoute `json:"ServiceRoutes,omitempty" name:"ServiceRoutes"`
	// 无

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 无

	SubnetId *uint64 `json:"SubnetId,omitempty" name:"SubnetId"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Proto *string `json:"Proto,omitempty" name:"Proto"`
}

func (r *AddServiceRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddServiceRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewaysRequest struct {
	*tchttp.BaseRequest

	// 网关集群类型(0-pvgw 1-vpcgw 2-jnsgw 3-natgw 4-mcgw 5-dcgw 6-pcgw 7-xgw 8-sxgw 9-underly-natgw)

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 翻页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页展示的条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤参数rip

	Rip *string `json:"Rip,omitempty" name:"Rip"`
	// 过滤参数VpcgwGroupId

	VpcgwGroupId *string `json:"VpcgwGroupId,omitempty" name:"VpcgwGroupId"`
	// 过滤参数vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 其他过滤字段

	Search *string `json:"Search,omitempty" name:"Search"`
	// 过滤参数EnableNat

	EnableNat *string `json:"EnableNat,omitempty" name:"EnableNat"`
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeGatewaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewaysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 简略返回

	Brief *string `json:"Brief,omitempty" name:"Brief"`
	// 组播

	EnableMulticast *int64 `json:"EnableMulticast,omitempty" name:"EnableMulticast"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方式

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 过滤条件，参数不支持同时指定VpcIds和Filters。
	// <li>AppId - String - （过滤条件）用户的AppId</li>
	// <li>UniqVpcId - String - （过滤条件）用户的vpc唯一id。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVpcExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMulticastGroupRequest struct {
	*tchttp.BaseRequest

	// VPC数字ID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 组播组地址。

	GroupAddr *string `json:"GroupAddr,omitempty" name:"GroupAddr"`
	// 组播组限速pps值，Pps和Bps至少有一个不为空。

	Pps *int64 `json:"Pps,omitempty" name:"Pps"`
	// 组播组限速bps值，Pps和Bps至少有一个不为空。

	Bps *int64 `json:"Bps,omitempty" name:"Bps"`
}

func (r *ModifyMulticastGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMulticastGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNatDetectIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNatDetectIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNatDetectIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUnderlayNatGatewayRuleWipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误信息

		Error *Error `json:"Error,omitempty" name:"Error"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUnderlayNatGatewayRuleWipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUnderlayNatGatewayRuleWipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJnsGatewayServicesRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 条数限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段,可选值有：createTime, owner。
	// OrderField与OrderDirection一起传入排序才能生效。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方式,asc：升序， desc：降序，默认升序。

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 过滤器,支持的字段：AppId,UniqVpcId

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeJnsGatewayServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJnsGatewayServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcEndPointServiceWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 白名单对象数组。

		VpcEndpointServiceUserSet []*VpcEndPointServiceUser `json:"VpcEndpointServiceUserSet,omitempty" name:"VpcEndpointServiceUserSet"`
		// 符合条件的白名单个数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcEndPointServiceWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcEndPointServiceWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHaVipListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// `HAVIP`对象数组。

		HaVipSet []*HaVip `json:"HaVipSet,omitempty" name:"HaVipSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHaVipListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHaVipListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetectIp struct {
	// natgw探测目的IP。

	Dip *string `json:"Dip,omitempty" name:"Dip"`
	// dip的可用状态，0:可用， 1:不可用，默认为0。

	Status *int64 `json:"Status,omitempty" name:"Status"`
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

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 专线网关Ip

	VpgIp *string `json:"VpgIp,omitempty" name:"VpgIp"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 专线网关名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// vpc掩码。如255.255.0.0

	VpcMask *string `json:"VpcMask,omitempty" name:"VpcMask"`
	// 专线网关通道数

	VpcIdConnectTunnelNum *uint64 `json:"VpcIdConnectTunnelNum,omitempty" name:"VpcIdConnectTunnelNum"`
	// 网关类型

	GwType *string `json:"GwType,omitempty" name:"GwType"`
	// 专线网关所在的集群名称

	DcgwClusterName *string `json:"DcgwClusterName,omitempty" name:"DcgwClusterName"`
	// BBgp

	BBgp *uint64 `json:"BBgp,omitempty" name:"BBgp"`
	// BBgpV2

	BBgpV2 *uint64 `json:"BBgpV2,omitempty" name:"BBgpV2"`
	// BBgpV2Sync

	BBgpV2Sync *uint64 `json:"BBgpV2Sync,omitempty" name:"BBgpV2Sync"`
	// BCas

	BCas *uint64 `json:"BCas,omitempty" name:"BCas"`
	// BEcmpVpg

	BEcmpVpg *uint64 `json:"BEcmpVpg,omitempty" name:"BEcmpVpg"`
	// BFlowDetails

	BFlowDetails *uint64 `json:"BFlowDetails,omitempty" name:"BFlowDetails"`
	// BLocalZoneVpg

	BLocalZoneVpg *uint64 `json:"BLocalZoneVpg,omitempty" name:"BLocalZoneVpg"`
	// BNewAfc

	BNewAfc *uint64 `json:"BNewAfc,omitempty" name:"BNewAfc"`
	// BPublishDockerCidr

	BPublishDockerCidr *uint64 `json:"BPublishDockerCidr,omitempty" name:"BPublishDockerCidr"`
	// BPublishPublicServiceSubnet

	BPublishPublicServiceSubnet *uint64 `json:"BPublishPublicServiceSubnet,omitempty" name:"BPublishPublicServiceSubnet"`
	// BPublishVbcRouteInNoneIdcSubnet

	BPublishVbcRouteInNoneIdcSubnet *uint64 `json:"BPublishVbcRouteInNoneIdcSubnet,omitempty" name:"BPublishVbcRouteInNoneIdcSubnet"`
	// BPublishVbcVpcCidr

	BPublishVbcVpcCidr *uint64 `json:"BPublishVbcVpcCidr,omitempty" name:"BPublishVbcVpcCidr"`
	// BPublishVbcVpcCommunity

	BPublishVbcVpcCommunity *uint64 `json:"BPublishVbcVpcCommunity,omitempty" name:"BPublishVbcVpcCommunity"`
	// BVbc

	BVbc *uint64 `json:"BVbc,omitempty" name:"BVbc"`
	// BVbcDc

	BVbcDc *uint64 `json:"BVbcDc,omitempty" name:"BVbcDc"`
	// BZoneVpg

	BZoneVpg *uint64 `json:"BZoneVpg,omitempty" name:"BZoneVpg"`
	// Bandwidth

	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// BandwidthRuleId

	BandwidthRuleId *string `json:"BandwidthRuleId,omitempty" name:"BandwidthRuleId"`
	// FlowDetailsUpdateTime

	FlowDetailsUpdateTime *string `json:"FlowDetailsUpdateTime,omitempty" name:"FlowDetailsUpdateTime"`
	// NatType

	NatType *uint64 `json:"NatType,omitempty" name:"NatType"`
	// UniqVbcId

	UniqVbcId *string `json:"UniqVbcId,omitempty" name:"UniqVbcId"`
	// VbcId

	VbcId *uint64 `json:"VbcId,omitempty" name:"VbcId"`
	// VbcRtbType

	VbcRtbType *uint64 `json:"VbcRtbType,omitempty" name:"VbcRtbType"`
	// 可用区Id。

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DeleteNfvGwResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNfvGwResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNfvGwResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDetectTypeNamesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDetectTypeNamesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDetectTypeNamesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UNatGwRuleRip struct {
	// underlay natgateway规则的ID

	UnatgwRuleId *uint64 `json:"UnatgwRuleId,omitempty" name:"UnatgwRuleId"`
	// underlay natgateway规则唯一ID

	UniqUnatgwRuleId *string `json:"UniqUnatgwRuleId,omitempty" name:"UniqUnatgwRuleId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 网关实际内网IP

	Rip *string `json:"Rip,omitempty" name:"Rip"`
	// 运营商

	CarrierId *uint64 `json:"CarrierId,omitempty" name:"CarrierId"`
}

type CreateJnsGatewayServiceRequest struct {
	*tchttp.BaseRequest

	// ZoneId

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// UniqVgwIndex

	UniqVgwIndex *string `json:"UniqVgwIndex,omitempty" name:"UniqVgwIndex"`
	// UniqVgwIndex

	VgwIndex *string `json:"VgwIndex,omitempty" name:"VgwIndex"`
	// Count

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// Business

	Business *string `json:"Business,omitempty" name:"Business"`
	// UniqVpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// VgwType

	VgwType *uint64 `json:"VgwType,omitempty" name:"VgwType"`
	// Pip

	Pip *string `json:"Pip,omitempty" name:"Pip"`
	// Proto

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// Pport

	Pport *uint64 `json:"Pport,omitempty" name:"Pport"`
	// Owner

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// UniqSubnetId

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
	// BusinessOwner

	BusinessOwner *string `json:"BusinessOwner,omitempty" name:"BusinessOwner"`
}

func (r *CreateJnsGatewayServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateJnsGatewayServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGatewayGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGatewayGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGatewayGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcGlobalExtendCidrRequest struct {
	*tchttp.BaseRequest

	// 无

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVpcGlobalExtendCidrRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcGlobalExtendCidrRequest) FromJsonString(s string) error {
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

type VmConf struct {
	// 无

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
	// 无

	VmIp *string `json:"VmIp,omitempty" name:"VmIp"`
	// 无

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 无

	LanInLimit *uint64 `json:"LanInLimit,omitempty" name:"LanInLimit"`
	// 无

	LanOutLimit *uint64 `json:"LanOutLimit,omitempty" name:"LanOutLimit"`
	// 无

	WanInLimit *uint64 `json:"WanInLimit,omitempty" name:"WanInLimit"`
	// 无

	WanOutLimit *uint64 `json:"WanOutLimit,omitempty" name:"WanOutLimit"`
}

type DeleteUnderlayNatGatewayRuleWipRequest struct {
	*tchttp.BaseRequest

	// underlay natgateway规则详细信息

	UnatgwRuleWips []*UNatGwRuleWip `json:"UnatgwRuleWips,omitempty" name:"UnatgwRuleWips"`
}

func (r *DeleteUnderlayNatGatewayRuleWipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUnderlayNatGatewayRuleWipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnderlayNatGatewayCarriorInfoRequest struct {
	*tchttp.BaseRequest

	// Filters

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeUnderlayNatGatewayCarriorInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnderlayNatGatewayCarriorInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserGatewaysRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 条数限制

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段,可选值有：createTime, owner。
	// OrderField与OrderDirection一起传入排序才能生效。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方式,asc：升序， desc：降序，默认升序。

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 用户APPID。

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 网关ID。

	UsrgwId *uint64 `json:"UsrgwId,omitempty" name:"UsrgwId"`
	// UniqUsrgwId

	UniqUsrgwId *string `json:"UniqUsrgwId,omitempty" name:"UniqUsrgwId"`
}

func (r *DescribeUserGatewaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserGatewaysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNfvGwNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateNfvGwNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNfvGwNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayGroupsRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方式

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 网关类型(0-pvgw 1-vpcgw 2-jnsgw 3-natgw 4-mcgw 5-dcgw 6-pcgw 7-xgw 8-sxgw 9-underly-natgw)

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 网关集群组ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 搜索字段，可为网关集群组名字或者ID，支持模糊查询

	Search *string `json:"Search,omitempty" name:"Search"`
}

func (r *DescribeGatewayGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全组策略列表

		UsgSet []*UsgPolicyInfo `json:"UsgSet,omitempty" name:"UsgSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDetectNodeInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ReturnCode

		AzName *uint64 `json:"AzName,omitempty" name:"AzName"`
		// ErrorInfo

		SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
		// SvrBelongingProducts

		SvrBelongingProducts *string `json:"SvrBelongingProducts,omitempty" name:"SvrBelongingProducts"`
		// SvrIdcName

		SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
		// SvrRackName

		SvrRackName *string `json:"SvrRackName,omitempty" name:"SvrRackName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDetectNodeInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDetectNodeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GwData struct {
	// 网关名称

	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`
	// 环境信息

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// rip列表

	Ips []*string `json:"Ips,omitempty" name:"Ips"`
	// 网关版本

	Ver *string `json:"Ver,omitempty" name:"Ver"`
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
