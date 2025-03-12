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

package v20221013

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DeletePvgwGwResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePvgwGwResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePvgwGwResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEipRequest struct {
	*tchttp.BaseRequest

	// 过滤器，具体如下：Owner：用户的APPID Eip：EIP（支持批量） IpType：IP类型，枚举值：EIP, EIP6,WanIP, CalcIP, AntiDDoSEIP Isp：线路类型，枚举值如下： 常规BGP：BGP 中国移动：CMCC 中国联通：CUCC 中国电信：CTCC Zone：可用区，枚举值从 DescribeEipZones 获取（https://capi.woa.com/apidoc?product=dfseip&version=2022-07-01&action=DescribeEipZones） BusinessType：枚举值如下： 中心出口一：DC 中心出口二：SOC PayMode：计费模式 UniqEipId：EIP唯一ID（支持批量） UniqEniId：弹性网卡ID UniqInstanceId：绑定实例唯一ID Uuid：实例UUID

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeEipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcEndPointServiceIpv6Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// EndPointServiceIpv6

		EndPointServiceIpv6 *EndPointServiceIpv6 `json:"EndPointServiceIpv6,omitempty" name:"EndPointServiceIpv6"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcEndPointServiceIpv6Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcEndPointServiceIpv6Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePvgwGwRequest struct {
	*tchttp.BaseRequest

	// 唯一id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// rip

	Rips *string `json:"Rips,omitempty" name:"Rips"`
	// 网关类型

	Module *string `json:"Module,omitempty" name:"Module"`
	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

func (r *UpdatePvgwGwRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePvgwGwRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePvgwGwResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePvgwGwResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePvgwGwResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EndPointIpv6 struct {
	// EndPointIPv6Id

	EndPointIPv6Id *string `json:"EndPointIPv6Id,omitempty" name:"EndPointIPv6Id"`
	// VpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// SubnetId

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// EndPointOwner

	EndPointOwner *string `json:"EndPointOwner,omitempty" name:"EndPointOwner"`
	// EndPointName

	EndPointName *string `json:"EndPointName,omitempty" name:"EndPointName"`
	// ServiceVpcId

	ServiceVpcId *string `json:"ServiceVpcId,omitempty" name:"ServiceVpcId"`
	// ServiceVip

	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`
	// EndPointServiceIPv6Id

	EndPointServiceIPv6Id *string `json:"EndPointServiceIPv6Id,omitempty" name:"EndPointServiceIPv6Id"`
	// EndPointVip

	EndPointVip *string `json:"EndPointVip,omitempty" name:"EndPointVip"`
	// State

	State *string `json:"State,omitempty" name:"State"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// GroupSet

	GroupSet []*GroupSet `json:"GroupSet,omitempty" name:"GroupSet"`
}

type DescribeNatSnatRequest struct {
	*tchttp.BaseRequest

	// 唯一NatId

	UniqNatId *string `json:"UniqNatId,omitempty" name:"UniqNatId"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 页个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤器，支持通过UniqResourceId、Eip过滤

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNatSnatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatSnatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupSet struct {
	// UsgId

	UsgId *string `json:"UsgId,omitempty" name:"UsgId"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeNatDnatRequest struct {
	*tchttp.BaseRequest

	// 唯一NatId

	UniqNatId *string `json:"UniqNatId,omitempty" name:"UniqNatId"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 页个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤器，支持通过Eip、Proto、Pip过滤

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNatDnatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatDnatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateNatGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 迁移任务ID

		TaskId []*string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigrateNatGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EndPointServiceIpv6 struct {
	// EndPointServiceIPv6Id

	EndPointServiceIPv6Id *string `json:"EndPointServiceIPv6Id,omitempty" name:"EndPointServiceIPv6Id"`
	// VpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// ServiceOwner

	ServiceOwner *string `json:"ServiceOwner,omitempty" name:"ServiceOwner"`
	// ServiceName

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// ServiceVip

	ServiceVip *string `json:"ServiceVip,omitempty" name:"ServiceVip"`
	// ServiceInstanceId

	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" name:"ServiceInstanceId"`
	// AutoAcceptFlag

	AutoAcceptFlag *bool `json:"AutoAcceptFlag,omitempty" name:"AutoAcceptFlag"`
	// EndPointCount

	EndPointCount *int64 `json:"EndPointCount,omitempty" name:"EndPointCount"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// WhiteListCount

	WhiteListCount *int64 `json:"WhiteListCount,omitempty" name:"WhiteListCount"`
	// EndPointSet

	EndPointSet []*EndPointIpv6 `json:"EndPointSet,omitempty" name:"EndPointSet"`
	// 整形终端节点服务id

	IntEndPointServiceIPv6Id *int64 `json:"IntEndPointServiceIPv6Id,omitempty" name:"IntEndPointServiceIPv6Id"`
}

type DescribeNatDnatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// dnat规则信息

		DnatSet []*DnatSet `json:"DnatSet,omitempty" name:"DnatSet"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatDnatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatDnatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatSnatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Snat规则信息

		SnatSet []*SnatSet `json:"SnatSet,omitempty" name:"SnatSet"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatSnatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatSnatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePvgwGwRequest struct {
	*tchttp.BaseRequest

	// 页个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// rip

	Rip *string `json:"Rip,omitempty" name:"Rip"`
	// 网关名称

	Search *string `json:"Search,omitempty" name:"Search"`
}

func (r *DescribePvgwGwRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePvgwGwRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcEndPointIpv6Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// EndPointIpv6

		EndPointIpv6 []*EndPointIpv6 `json:"EndPointIpv6,omitempty" name:"EndPointIpv6"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcEndPointIpv6Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcEndPointIpv6Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EipInfo struct {
	// 唯一eipid

	UniqEipId *string `json:"UniqEipId,omitempty" name:"UniqEipId"`
	// appid

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 用户的uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 弹性公网ip

	Eip *string `json:"Eip,omitempty" name:"Eip"`
	// 唯一实例id

	UniqInstanceId *string `json:"UniqInstanceId,omitempty" name:"UniqInstanceId"`
	// eip类型

	IpType *string `json:"IpType,omitempty" name:"IpType"`
	// anycast的ftp

	Ftp *bool `json:"Ftp,omitempty" name:"Ftp"`
	// anycast的sip

	Sip *bool `json:"Sip,omitempty" name:"Sip"`
	// 是否直通

	IsThrough *bool `json:"IsThrough,omitempty" name:"IsThrough"`
	// Eip状态，CREATING创建中，BINDING绑定中，BIND绑定，UNBINDING解绑中，UNBIND未绑定，OFFLINING离线，CREATE_FAILED创建失败，BIND_ENI悬空eip，PRIVATEP：IPV6内网。

	Status *string `json:"Status,omitempty" name:"Status"`
	// rs的ip

	RsIp *string `json:"RsIp,omitempty" name:"RsIp"`
	// 计费模式：POSTPAID_BY_MONTH(带宽按月后付费)，PREPAID_BY_MONTH(带宽按月预付费)，TRAFFIC_BY_HOUR(流量按小时后付费)，POSTPAID_BY_HOUR(带宽按小时后付费)，BANDWIDTH_PACKAGE(带宽包计费)。按流量小时后付费的才具有后付费详情

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 带宽上限

	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 是否欠费

	Arrears *bool `json:"Arrears,omitempty" name:"Arrears"`
	// 隔离时间

	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`
	// eip的id

	EipId *string `json:"EipId,omitempty" name:"EipId"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否自动续费

	AutoRenewFlag *bool `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// 包月购买带宽

	PurchaseBandwidth *int64 `json:"PurchaseBandwidth,omitempty" name:"PurchaseBandwidth"`
	// idc名称

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// 可用区名称

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区id

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// vpc的id

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 项目id

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 实例uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 唯一弹性网卡id

	UniqEniId *string `json:"UniqEniId,omitempty" name:"UniqEniId"`
	// 最近解绑时间

	FreeTime *string `json:"FreeTime,omitempty" name:"FreeTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 是否eip

	IsEip *bool `json:"IsEip,omitempty" name:"IsEip"`
	// 到期时间

	DeadLine *string `json:"DeadLine,omitempty" name:"DeadLine"`
	// 用户类型

	UserType *bool `json:"UserType,omitempty" name:"UserType"`
	// 购买源

	PurchaseSource *string `json:"PurchaseSource,omitempty" name:"PurchaseSource"`
	// 是否WANIP转换

	FromWanIP *bool `json:"FromWanIP,omitempty" name:"FromWanIP"`
	// regionEip

	IsNewGeneration *bool `json:"IsNewGeneration,omitempty" name:"IsNewGeneration"`
	// 是否主动外联

	IsActiveAccessEip *bool `json:"IsActiveAccessEip,omitempty" name:"IsActiveAccessEip"`
	// 网络组

	NetworkGroup *string `json:"NetworkGroup,omitempty" name:"NetworkGroup"`
	// CDC集群id

	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
	// isp

	Isp *string `json:"Isp,omitempty" name:"Isp"`
	// anycastd的city

	AnycastCity *string `json:"AnycastCity,omitempty" name:"AnycastCity"`
	// anycast的zone

	AnycastZone *string `json:"AnycastZone,omitempty" name:"AnycastZone"`
	// 地域

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 是否隔离

	Isolate *bool `json:"Isolate,omitempty" name:"Isolate"`
	// TGW组

	TgwGroup *string `json:"TgwGroup,omitempty" name:"TgwGroup"`
	// 地域ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 网络出口

	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`
}

type DescribeEipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// EIP实例详情数组。

		Eip []*EipInfo `json:"Eip,omitempty" name:"Eip"`
		// 符合条件的总数量。

		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GatewaySet struct {
	// 模块，用于区分pvgw1.0和2.0

	Module *string `json:"Module,omitempty" name:"Module"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// rip

	Rip *string `json:"Rip,omitempty" name:"Rip"`
	// 唯一id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// cdc集群id

	UniqCdcId *string `json:"UniqCdcId,omitempty" name:"UniqCdcId"`
	// 集群版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
}

type CreatePvgwGwRequest struct {
	*tchttp.BaseRequest

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// rip

	Rips *string `json:"Rips,omitempty" name:"Rips"`
	// 网关类型

	Module *string `json:"Module,omitempty" name:"Module"`
	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// cdc集群id

	UniqCdcId *string `json:"UniqCdcId,omitempty" name:"UniqCdcId"`
}

func (r *CreatePvgwGwRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePvgwGwRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateNatGatewayRequest struct {
	*tchttp.BaseRequest

	// 集群VIP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 新的集群VIP

	NewGwIp *string `json:"NewGwIp,omitempty" name:"NewGwIp"`
	// 唯一natid

	UniqNatId *string `json:"UniqNatId,omitempty" name:"UniqNatId"`
	// 冷热迁移

	Live *uint64 `json:"Live,omitempty" name:"Live"`
}

func (r *MigrateNatGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateNatGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePvgwGwResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePvgwGwResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePvgwGwResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePvgwGwResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网关列表

		GatewaySet []*GatewaySet `json:"GatewaySet,omitempty" name:"GatewaySet"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePvgwGwResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePvgwGwResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcEndPointIpv6Request struct {
	*tchttp.BaseRequest

	// 过滤器。通过AppId、VpcId、EndPointIPv6Id、EndPointName、EndPointServiceIPv6Id过滤

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 页个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeVpcEndPointIpv6Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcEndPointIpv6Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnatSet struct {
	// 弹性公网IP

	Eip []*string `json:"Eip,omitempty" name:"Eip"`
	// 唯一NatSantId

	UniqNatSnatId *string `json:"UniqNatSnatId,omitempty" name:"UniqNatSnatId"`
	// VpcId

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 唯一VpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// 资源Id

	UniqResourceId *string `json:"UniqResourceId,omitempty" name:"UniqResourceId"`
	// 资源类型

	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`
	// Appid

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// NatId

	NatId *int64 `json:"NatId,omitempty" name:"NatId"`
	// 源IP/源网段

	Pip *string `json:"Pip,omitempty" name:"Pip"`
	// 唯一NatId

	UniqNatId *string `json:"UniqNatId,omitempty" name:"UniqNatId"`
	// NatSnatId

	NatSnatId *int64 `json:"NatSnatId,omitempty" name:"NatSnatId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DeletePvgwGwRequest struct {
	*tchttp.BaseRequest

	// 唯一ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 网关类型

	Module *string `json:"Module,omitempty" name:"Module"`
	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

func (r *DeletePvgwGwRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePvgwGwRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcEndPointServiceIpv6Request struct {
	*tchttp.BaseRequest

	// 过滤器。支持AppId、EndPointServiceIPv6Id、ServiceName、ServiceInstanceId、VpcId过滤

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 页个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeVpcEndPointServiceIpv6Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcEndPointServiceIpv6Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DnatSet struct {
	// 外部IP

	Eip *string `json:"Eip,omitempty" name:"Eip"`
	// NatId

	NatId *int64 `json:"NatId,omitempty" name:"NatId"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 唯一VpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// 协议

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// 内部端口

	Pport *int64 `json:"Pport,omitempty" name:"Pport"`
	// 外部端口

	Eport *int64 `json:"Eport,omitempty" name:"Eport"`
	// Appid

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// VpcId

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 协议类型

	PipType *int64 `json:"PipType,omitempty" name:"PipType"`
	// 内部IP

	Pip *string `json:"Pip,omitempty" name:"Pip"`
	// 唯一NatId

	UniqNatId *string `json:"UniqNatId,omitempty" name:"UniqNatId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type Filters struct {
	// key

	Name *string `json:"Name,omitempty" name:"Name"`
	// values

	Values []*string `json:"Values,omitempty" name:"Values"`
}
