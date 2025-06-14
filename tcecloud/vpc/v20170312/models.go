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

type DescribeVpnGatewaysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// VPN网关实例详细信息列表。

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

type Ip6Rule struct {
	// IPV6转换规则唯一ID，形如rule6-xxxxxxxx

	Ip6RuleId *string `json:"Ip6RuleId,omitempty" name:"Ip6RuleId"`
	// IPV6转换规则名称

	Ip6RuleName *string `json:"Ip6RuleName,omitempty" name:"Ip6RuleName"`
	// IPV6地址

	Vip6 *string `json:"Vip6,omitempty" name:"Vip6"`
	// IPV6端口号

	Vport6 *int64 `json:"Vport6,omitempty" name:"Vport6"`
	// 协议类型，支持TCP/UDP

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// IPV4地址

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// IPV4端口号

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 转换规则状态，限于CREATING,RUNNING,DELETING,MODIFYING

	RuleStatus *string `json:"RuleStatus,omitempty" name:"RuleStatus"`
	// 转换规则创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type DescribeFlowLogsRequest struct {
	*tchttp.BaseRequest

	// 私用网络ID或者统一ID，建议使用统一ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 流日志唯一ID

	FlowLogId *string `json:"FlowLogId,omitempty" name:"FlowLogId"`
	// 流日志实例名字

	FlowLogName *string `json:"FlowLogName,omitempty" name:"FlowLogName"`
	// 流日志所属资源类型，VPC|SUBNET|NETWORKINTERFACE

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源唯一ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 流日志采集类型，ACCEPT|REJECT|ALL

	TrafficType *string `json:"TrafficType,omitempty" name:"TrafficType"`
	// 流日志存储ID

	CloudLogId *string `json:"CloudLogId,omitempty" name:"CloudLogId"`
	// 流日志存储ID状态

	CloudLogState *string `json:"CloudLogState,omitempty" name:"CloudLogState"`
	// 按某个字段排序,支持字段：flowLogName,createTime，默认按createTime

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 升序（asc）还是降序（desc）,默认：desc

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页行数，默认为10

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeFlowLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressAssociationQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressAssociationQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressAssociationQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBandwidthAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 带宽属性数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 带宽属性详细信息

		BandwidthAttribute []*BandwidthAttribute `json:"BandwidthAttribute,omitempty" name:"BandwidthAttribute"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBandwidthAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBandwidthAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// 专线网关ID，形如：`dcg-prpqlmg1`。

	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	// 云联网路由学习类型，可选值：
	// <li>`BGP` - 自动学习。</li>
	// <li>`STATIC` - 静态，即用户配置，默认值。</li>

	CcnRouteType *string `json:"CcnRouteType,omitempty" name:"CcnRouteType"`
	// 偏移量。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDirectConnectGatewayCcnRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectGatewayCcnRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpOnlineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IP地址库信息列表

		AddressInfo *AddressLibInfo `json:"AddressInfo,omitempty" name:"AddressInfo"`
		// IP地址库信息个数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpOnlineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpOnlineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateNatGatewayAddressRequest struct {
	*tchttp.BaseRequest

	// NAT网关的ID，形如：`nat-df45454`。

	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`
	// 绑定NAT网关的弹性IP数组。

	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`
}

func (r *DisassociateNatGatewayAddressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateNatGatewayAddressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetDetectRequest struct {
	*tchttp.BaseRequest

	// `VPC`实例`ID`。形如：`vpc-12345678`

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网实例ID。形如：subnet-12345678。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 网络探测名称，最大长度不能超过60个字节。

	NetDetectName *string `json:"NetDetectName,omitempty" name:"NetDetectName"`
	// 探测目的IPv4地址数组。最多两个。

	DetectDestinationIp []*string `json:"DetectDestinationIp,omitempty" name:"DetectDestinationIp"`
	// 下一跳类型，目前我们支持的类型有：
	// VPN：VPN网关；
	// DIRECTCONNECT：专线网关；
	// PEERCONNECTION：对等连接；
	// NAT：NAT网关；
	// NORMAL_CVM：普通云主机；

	NextHopType *string `json:"NextHopType,omitempty" name:"NextHopType"`
	// 下一跳目的网关，取值与“下一跳类型”相关：
	// 下一跳类型为VPN，取值VPN网关ID，形如：vpngw-12345678；
	// 下一跳类型为DIRECTCONNECT，取值专线网关ID，形如：dcg-12345678；
	// 下一跳类型为PEERCONNECTION，取值对等连接ID，形如：pcx-12345678；
	// 下一跳类型为NAT，取值Nat网关，形如：nat-12345678；
	// 下一跳类型为NORMAL_CVM，取值云主机IPv4地址，形如：10.0.0.12；

	NextHopDestination *string `json:"NextHopDestination,omitempty" name:"NextHopDestination"`
	// 网络探测描述。

	NetDetectDescription *string `json:"NetDetectDescription,omitempty" name:"NetDetectDescription"`
}

func (r *CreateNetDetectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetDetectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTrafficPackagesRequest struct {
	*tchttp.BaseRequest

	// 流量包规格。

	TrafficAmount *uint64 `json:"TrafficAmount,omitempty" name:"TrafficAmount"`
	// 流量包数量。

	TrafficPackageCount *uint64 `json:"TrafficPackageCount,omitempty" name:"TrafficPackageCount"`
}

func (r *CreateTrafficPackagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTrafficPackagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SingleIspZone struct {
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 三网运营商，包括"CMCC", "CTCC", "CUCC"

	Singleisp []*string `json:"Singleisp,omitempty" name:"Singleisp"`
}

type CheckNetDetectStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络探测验证结果对象数组。

		NetDetectIpStateSet []*NetDetectIpState `json:"NetDetectIpStateSet,omitempty" name:"NetDetectIpStateSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckNetDetectStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckNetDetectStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BandwidthAttribute struct {
	// 网络计费模式列表，取值包括
	// 'BANDWIDTH_POSTPAID_BY_MONTH',
	// 'BANDWIDTH_PREPAID_BY_MONTH',
	// 'TRAFFIC_POSTPAID_BY_HOUR',
	// 'BANDWIDTH_POSTPAID_BY_HOUR',
	// 'BANDWIDTH_PACKAGE'

	NetworkPayMode []*string `json:"NetworkPayMode,omitempty" name:"NetworkPayMode"`
	// cvm实例计费模式列表，取值包括"PREPAID","POSTPAID_BY_HOUR","CDHPAID"。

	CvmPayMode []*string `json:"CvmPayMode,omitempty" name:"CvmPayMode"`
	// cpu核数范围，前闭后闭区间。如[9,23]表示9 <= cpu核心数 <= 23。空区间表示任意cpu核心数都支持。

	Cpu []*int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 可用区ID列表。空列表表示任意可用区Id都支持。

	ZoneId []*int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 带宽最大值

	BandwidthUpper *int64 `json:"BandwidthUpper,omitempty" name:"BandwidthUpper"`
	// 带宽是否必须有上限。1表示有上限，0表示无上限

	BandwidthLimit *int64 `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`
	// 资源唯一ID列表，可以是云服务器InstanceId，或者弹性公网IP唯一ID

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
}

type TrafficMirrorFilter struct {
	// 过滤规则的源网段

	SrcNet *string `json:"SrcNet,omitempty" name:"SrcNet"`
	// 过滤规则的目的网段

	DstNet *string `json:"DstNet,omitempty" name:"DstNet"`
	// 过滤规则的源端口，默认值1-65535

	SrcPort *string `json:"SrcPort,omitempty" name:"SrcPort"`
	// 过滤规则的目的端口，默认值1-65535

	DstPort *string `json:"DstPort,omitempty" name:"DstPort"`
	// 过滤规则的协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DescribeNetworkInterfacesExtraResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例详细信息列表。

		NetworkInterfaceSet []*NetworkInterface `json:"NetworkInterfaceSet,omitempty" name:"NetworkInterfaceSet"`
		// 符合条件的实例数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkInterfacesExtraResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkInterfacesExtraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpnConnectionRequest struct {
	*tchttp.BaseRequest

	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// VPN网关实例ID。

	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	// 对端网关ID，例如：cgw-2wqq41m9，可通过DescribeCustomerGateways接口查询对端网关。

	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`
	// 通道名称，可任意命名，但不得超过60个字符。

	VpnConnectionName *string `json:"VpnConnectionName,omitempty" name:"VpnConnectionName"`
	// 预共享密钥。

	PreShareKey *string `json:"PreShareKey,omitempty" name:"PreShareKey"`
	// SPD策略组，例如：{"10.0.0.5/24":["172.123.10.5/16"]}，10.0.0.5/24是vpc内网段172.123.10.5/16是IDC网段。用户指定VPC内哪些网段可以和您IDC中哪些网段通信。

	SecurityPolicyDatabases []*SecurityPolicyDatabase `json:"SecurityPolicyDatabases,omitempty" name:"SecurityPolicyDatabases"`
	// IKE配置（Internet Key Exchange，因特网密钥交换），IKE具有一套自我保护机制，用户配置网络安全协议

	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitempty" name:"IKEOptionsSpecification"`
	// IPSec配置，腾讯云提供IPSec安全会话设置

	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitempty" name:"IPSECOptionsSpecification"`
}

func (r *CreateVpnConnectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpnConnectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FirewallRuleSet struct {
	// 出站规则集合

	Egress []*Rule `json:"Egress,omitempty" name:"Egress"`
	// 入站规则集合

	Ingress []*Rule `json:"Ingress,omitempty" name:"Ingress"`
}

type Tag struct {
	// 标签键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeEipStatisticsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeEipStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEipStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInterfacesRequest struct {
	*tchttp.BaseRequest

	// 弹性网卡实例ID查询。形如：eni-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定NetworkInterfaceIds和Filters。

	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" name:"NetworkInterfaceIds"`
	// 过滤条件，参数不支持同时指定NetworkInterfaceIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>subnet-id - String - （过滤条件）所属子网实例ID，形如：subnet-f49l6u0z。</li>
	// <li>network-interface-id - String - （过滤条件）弹性网卡实例ID，形如：eni-5k56k7k7。</li>
	// <li>attachment.instance-id - String - （过滤条件）绑定的云服务器实例ID，形如：ins-3nqpdn3i。</li>
	// <li>groups.security-group-id - String - （过滤条件）绑定的安全组实例ID，例如：sg-f9ekbxeq。</li>
	// <li>network-interface-name - String - （过滤条件）网卡实例名称。</li>
	// <li>network-interface-description - String - （过滤条件）网卡实例描述。</li>
	// <li>address-ip - String - （过滤条件）内网IPv4地址。</li>
	// <li>tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。使用请参考示例2</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例3。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNetworkInterfacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkInterfacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseIp6AddressesBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseIp6AddressesBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseIp6AddressesBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetVpnConnectionRequest struct {
	*tchttp.BaseRequest

	// VPN网关实例ID。

	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	// VPN通道实例ID。形如：vpnx-f49l6u0z。

	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`
}

func (r *ResetVpnConnectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetVpnConnectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddressBasic struct {
	// 外网IP地址

	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`
	// `EIP`的`ID`，是`EIP`的唯一标识。

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// eip资源类型，包括"CalcIP","WanIP","EIP","AnycastEIP"。其中"CalcIP"表示设备ip，“WanIP”表示普通公网ip，“EIP”表示弹性公网ip，“AnycastEip”表示加速EIP

	AddressType *string `json:"AddressType,omitempty" name:"AddressType"`
	// EIP运营商信息，包括"CMCC","CTCC","CUCC","BGP"，其中"CMCC"表示移动运营商，"CTCC"表示电信运营商，"CUCC"表示联通运营商，"BGP"表示BGP对接运营商

	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`
	// EIP使用状态，"used"表示该ip地址在当前地域已被占用，"unused"表示该ip地址在当前地域未被占用

	Status *string `json:"Status,omitempty" name:"Status"`
	// Anycast 发布域是加速 IP 地址发布的地域，即 Anycast CLB 的 VIP 所发布的 POP 接入点，客户端接入最近的 POP 接入点。目前 Anycast CLB 支持以下发布域：
	// 发布域 A(ANYCAST_ZONE_A)：主要是中国和欧美地区，VIP 将同时在在北京、上海、广州、中国香港、多伦多、硅谷、法兰克福、弗吉尼亚和莫斯科发布。
	// 发布域 B(ANYCAST_ZONE_B)：主要是中国和东南亚地区，VIP 将同时在北京、上海、广州、中国香港、新加坡、首尔、孟买、曼谷和东京发布。

	AnycastZone *string `json:"AnycastZone,omitempty" name:"AnycastZone"`
}

type CreateBandwidthPackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 带宽包Id

		BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
		// 带宽包Ids(申请数量大于1时有效)

		BandwidthPackageIds []*string `json:"BandwidthPackageIds,omitempty" name:"BandwidthPackageIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBandwidthPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBandwidthPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBandwidthPackageBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBandwidthPackageBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBandwidthPackageBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayFlowMonitorDetailRequest struct {
	*tchttp.BaseRequest

	// 时间点。表示要查询这分钟内的明细。如：`2019-02-28 18:15:20`，将查询 `18:15` 这一分钟内的明细。

	TimePoint *string `json:"TimePoint,omitempty" name:"TimePoint"`
	// VPN网关实例ID，形如：`vpn-ltjahce6`。

	VpnId *string `json:"VpnId,omitempty" name:"VpnId"`
	// 专线网关实例ID，形如：`dcg-ltjahce6`。

	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	// 对等连接实例ID，形如：`pcx-ltjahce6`。

	PeeringConnectionId *string `json:"PeeringConnectionId,omitempty" name:"PeeringConnectionId"`
	// NAT网关实例ID，形如：`nat-ltjahce6`。

	NatId *string `json:"NatId,omitempty" name:"NatId"`
	// 偏移量。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段。支持 `InPkg` `OutPkg` `InTraffic` `OutTraffic`。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方法。顺序：`ASC`，倒序：`DESC`。

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeGatewayFlowMonitorDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayFlowMonitorDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全组规则集合。

		SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnAttributeRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// CCN名称，最大长度不能超过60个字节。

	CcnName *string `json:"CcnName,omitempty" name:"CcnName"`
	// CCN描述信息，最大长度不能超过100个字节。

	CcnDescription *string `json:"CcnDescription,omitempty" name:"CcnDescription"`
}

func (r *ModifyCcnAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkInterfaceRequest struct {
	*tchttp.BaseRequest

	// 弹性网卡实例ID，例如：eni-m6dyj72l。

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
}

func (r *DeleteNetworkInterfaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkInterfaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *DescribeCcnRegionBandwidthLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteConflictsRequest struct {
	*tchttp.BaseRequest

	// 路由表实例ID，例如：rtb-azd4dt1c。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 要检查的与之冲突的目的端列表

	DestinationCidrBlocks []*string `json:"DestinationCidrBlocks,omitempty" name:"DestinationCidrBlocks"`
}

func (r *DescribeRouteConflictsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteConflictsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceDirectConnectGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// 专线网关ID，形如：dcg-prpqlmg1

	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	// 需要连通的IDC网段列表

	Routes []*DirectConnectGatewayCcnRoute `json:"Routes,omitempty" name:"Routes"`
}

func (r *ReplaceDirectConnectGatewayCcnRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceDirectConnectGatewayCcnRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBandwidthPackageQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBandwidthPackageQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBandwidthPackageQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNatGatewayDestinationIpPortTranslationNatRuleRequest struct {
	*tchttp.BaseRequest

	// NAT网关的ID，形如：`nat-df45454`。

	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`
	// NAT网关的端口转换规则。

	DestinationIpPortTranslationNatRules []*DestinationIpPortTranslationNatRule `json:"DestinationIpPortTranslationNatRules,omitempty" name:"DestinationIpPortTranslationNatRules"`
}

func (r *CreateNatGatewayDestinationIpPortTranslationNatRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNatGatewayDestinationIpPortTranslationNatRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnassignIpv6AddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnassignIpv6AddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnassignIpv6AddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcAttributeRequest struct {
	*tchttp.BaseRequest

	// VPC实例ID。形如：vpc-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpcIds和Filters。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络名称，可任意命名，但不得超过60个字符。

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 是否开启组播。true: 开启, false: 关闭。

	EnableMulticast *string `json:"EnableMulticast,omitempty" name:"EnableMulticast"`
	// DNS地址，最多支持4个，第1个默认为主，其余为备

	DnsServers []*string `json:"DnsServers,omitempty" name:"DnsServers"`
	// 域名

	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *ModifyVpcAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全组对象。

		SecurityGroupSet []*SecurityGroup `json:"SecurityGroupSet,omitempty" name:"SecurityGroupSet"`
		// 符合条件的实例数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateAddressesRequest struct {
	*tchttp.BaseRequest

	// 弹性公网IP当前所属的账号AppId。公共请求参数中的AppId是迁移目标账号AppId。

	OldAppId *uint64 `json:"OldAppId,omitempty" name:"OldAppId"`
	// 弹性公网IP当前所属的账号Uin。公共请求参数中的Uin是迁移目标账号Uin

	OldUin *string `json:"OldUin,omitempty" name:"OldUin"`
	// 用于检测弹性公网IP是否满足迁移条件，不会执行实际迁移。

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 待迁移的弹性公网IP唯一ID列表

	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
}

func (r *MigrateAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIp6AddressesBandwidthRequest struct {
	*tchttp.BaseRequest

	// IPV6地址。Ip6Addresses和Ip6AddressId必须且只能传一个

	Ip6Addresses []*string `json:"Ip6Addresses,omitempty" name:"Ip6Addresses"`
	// IPV6地址对应的唯一ID，形如eip-xxxxxxxx。Ip6Addresses和Ip6AddressId必须且只能传一个

	Ip6AddressIds []*string `json:"Ip6AddressIds,omitempty" name:"Ip6AddressIds"`
	// 修改的目标带宽，单位Mbps

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

func (r *ModifyIp6AddressesBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIp6AddressesBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Ipv6SubnetCidrBlock struct {
	// 子网实例`ID`。形如：`subnet-pxir56ns`。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// `IPv6`子网段。形如：`3402:4e00:20:1001::/64`

	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`
}

type DescribeCustomerGatewaysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 对端网关对象列表

		CustomerGatewaySet []*CustomerGateway `json:"CustomerGatewaySet,omitempty" name:"CustomerGatewaySet"`
		// 符合条件的实例数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomerGatewaysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTrafficMirrorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流量镜像实例

		TrafficMirror *TrafficMirror `json:"TrafficMirror,omitempty" name:"TrafficMirror"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTrafficMirrorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTrafficMirrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnConnectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpnConnectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpnConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateAddressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务TaskId。可以使用[DescribeTaskResult](https://cloud.tencent.com/document/api/215/36271)接口查询任务状态。

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateAddressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGatewaysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// NAT网关对象数组。

		NatGatewaySet []*NatGateway `json:"NatGatewaySet,omitempty" name:"NatGatewaySet"`
		// 符合条件的NAT网关对象个数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
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

type DescribeRouteTablesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 路由表对象。

		RouteTableSet []*RouteTable `json:"RouteTableSet,omitempty" name:"RouteTableSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteTablesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteTableAttributeRequest struct {
	*tchttp.BaseRequest

	// 路由表实例ID，例如：rtb-azd4dt1c。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 路由表名称。

	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
}

func (r *ModifyRouteTableAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRouteTableAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetTrafficMirrorTargetRequest struct {
	*tchttp.BaseRequest

	// 流量镜像实例ID

	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`
	// 流量镜像的接收目的信息

	CollectorTarget *TrafficMirrorTarget `json:"CollectorTarget,omitempty" name:"CollectorTarget"`
}

func (r *ResetTrafficMirrorTargetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetTrafficMirrorTargetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAddressTemplateRequest struct {
	*tchttp.BaseRequest

	// IP地址模版名称

	AddressTemplateName *string `json:"AddressTemplateName,omitempty" name:"AddressTemplateName"`
	// 地址信息，支持 IP、CIDR、IP 范围。

	Addresses []*string `json:"Addresses,omitempty" name:"Addresses"`
}

func (r *CreateAddressTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAddressTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest

	// VPN网关实例ID。

	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100；单位：Mbps。

	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteTableAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRouteTableAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRouteTableAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicLinkInstancesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>vpc-id - String - （过滤条件）VPC实例ID。</li>
	// <li>vm-ip - String - （过滤条件）基础网络云主机IP。</li>

	Filters []*FilterObject `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeClassicLinkInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClassicLinkInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrafficPackagesRequest struct {
	*tchttp.BaseRequest

	// 共享流量包ID，支持批量

	TrafficPackageIds []*string `json:"TrafficPackageIds,omitempty" name:"TrafficPackageIds"`
	// 每次请求的`Filters`的上限为10。参数不支持同时指定`TrafficPackageIds`和`Filters`。详细的过滤条件如下：
	// <li> traffic-package_id - String - 是否必填：否 - （过滤条件）按照共享流量包的唯一标识ID过滤。</li>
	// <li> traffic-package-name - String - 是否必填：否 - （过滤条件）按照共享流量包名称过滤。不支持模糊过滤。</li>
	// <li> status - String - 是否必填：否 - （过滤条件）按照共享流量包状态过滤。可选状态：[AVAILABLE|EXPIRED|EXHAUSTED]</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 查询共享流量包偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询共享流量包数量限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTrafficPackagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrafficPackagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTrafficMirrorAllFilterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateTrafficMirrorAllFilterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTrafficMirrorAllFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckNetDetectStateRequest struct {
	*tchttp.BaseRequest

	// 网络探测实例ID。形如：netd-12345678。

	NetDetectId *string `json:"NetDetectId,omitempty" name:"NetDetectId"`
	// `VPC`实例`ID`。形如：`vpc-12345678`

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网实例ID。形如：subnet-12345678。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 网络探测名称，最大长度不能超过60个字节。

	NetDetectName *string `json:"NetDetectName,omitempty" name:"NetDetectName"`
	// 探测目的IPv4地址数组，最多两个。

	DetectDestinationIp []*string `json:"DetectDestinationIp,omitempty" name:"DetectDestinationIp"`
	// 下一跳类型，目前我们支持的类型有：
	// VPN：VPN网关；
	// DIRECTCONNECT：专线网关；
	// PEERCONNECTION：对等连接；
	// NAT：NAT网关；
	// NORMAL_CVM：普通云主机；

	NextHopType *string `json:"NextHopType,omitempty" name:"NextHopType"`
	// 下一跳目的网关，取值与“下一跳类型”相关：
	// 下一跳类型为VPN，取值VPN网关ID，形如：vpngw-12345678；
	// 下一跳类型为DIRECTCONNECT，取值专线网关ID，形如：dcg-12345678；
	// 下一跳类型为PEERCONNECTION，取值对等连接ID，形如：pcx-12345678；
	// 下一跳类型为NAT，取值Nat网关，形如：nat-12345678；
	// 下一跳类型为NORMAL_CVM，取值云主机IPv4地址，形如：10.0.0.12；

	NextHopDestination *string `json:"NextHopDestination,omitempty" name:"NextHopDestination"`
}

func (r *CheckNetDetectStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckNetDetectStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnAttachedInstance struct {
	// 云联网实例ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 关联实例类型：
	// <li>`VPC`：私有网络</li>
	// <li>`DIRECTCONNECT`：专线网关</li>
	// <li>`BMVPC`：黑石私有网络</li>

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 关联实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 关联实例名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 关联实例所属大区，例如：ap-guangzhou。

	InstanceRegion *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`
	// 关联实例所属UIN（根账号）。

	InstanceUin *string `json:"InstanceUin,omitempty" name:"InstanceUin"`
	// 关联实例CIDR。

	CidrBlock []*string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 关联实例状态：
	// <li>`PENDING`：申请中</li>
	// <li>`ACTIVE`：已连接</li>
	// <li>`EXPIRED`：已过期</li>
	// <li>`REJECTED`：已拒绝</li>
	// <li>`DELETED`：已删除</li>
	// <li>`FAILED`：失败的（2小时后将异步强制解关联）</li>
	// <li>`ATTACHING`：关联中</li>
	// <li>`DETACHING`：解关联中</li>
	// <li>`DETACHFAILED`：解关联失败（2小时后将异步强制解关联）</li>

	State *string `json:"State,omitempty" name:"State"`
	// 关联时间。

	AttachedTime *string `json:"AttachedTime,omitempty" name:"AttachedTime"`
	// 云联网所属UIN（根账号）。

	CcnUin *string `json:"CcnUin,omitempty" name:"CcnUin"`
}

type GetDealStatusByNameRequest struct {
	*tchttp.BaseRequest

	// 订单号

	DealName *string `json:"DealName,omitempty" name:"DealName"`
}

func (r *GetDealStatusByNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDealStatusByNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewCcnBandwidthRequest struct {
	*tchttp.BaseRequest

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 流量配置ID。

	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`
}

func (r *RenewCcnBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewCcnBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetVpnGatewayInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest

	// VPN网关实例ID。

	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100；单位：Mbps。

	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

func (r *ResetVpnGatewayInternetMaxBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetVpnGatewayInternetMaxBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HaVip struct {
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
}

type GetRenewCcnBandwidthDealRequest struct {
	*tchttp.BaseRequest

	// 是否自动续费标识；NOTIFY_AND_AUTO_RENEW：自动续费，NOTIFY_AND_MANUAL_RENEW：手动续费。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 流量配置ID。

	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`
}

func (r *GetRenewCcnBandwidthDealRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRenewCcnBandwidthDealRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountAttribute struct {
	// 属性名

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 属性值

	AttributeValues []*string `json:"AttributeValues,omitempty" name:"AttributeValues"`
}

type HaVipDisassociateAddressIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *HaVipDisassociateAddressIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *HaVipDisassociateAddressIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupPolicyCheckInfoSet struct {
	// 安全组策略检测结果集合

	SecurityGroupPolicyCheckInfoSet []*SecurityGroupPolicyCheckInfo `json:"SecurityGroupPolicyCheckInfoSet,omitempty" name:"SecurityGroupPolicyCheckInfoSet"`
}

type DescribeSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest

	// 安全组实例ID，例如：sg-33ocnj9n，可通过DescribeSecurityGroups获取。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *DescribeSecurityGroupPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBandwidthPackageBandwidthRequest struct {
	*tchttp.BaseRequest

	// [已废弃]共享带宽包ID列表

	BandwidthPackageIds []*string `json:"BandwidthPackageIds,omitempty" name:"BandwidthPackageIds"`
	// 带宽包限速大小。单位：Mbps，-1表示不限速。

	InternetMaxBandwidth *int64 `json:"InternetMaxBandwidth,omitempty" name:"InternetMaxBandwidth"`
	// 共享带宽包ID

	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
}

func (r *ModifyBandwidthPackageBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBandwidthPackageBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Subnet struct {
	// `VPC`实例`ID`。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网实例`ID`，例如：subnet-bthucmmy。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 子网名称。

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 子网的 `IPv4` `CIDR`。

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 是否默认子网。

	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
	// 是否开启广播。

	EnableBroadcast *bool `json:"EnableBroadcast,omitempty" name:"EnableBroadcast"`
	// 可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 路由表实例ID，例如：rtb-l2h8d7c2。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 可用`IP`数。

	AvailableIpAddressCount *uint64 `json:"AvailableIpAddressCount,omitempty" name:"AvailableIpAddressCount"`
	// 子网的 `IPv6` `CIDR`。

	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`
	// 关联`ACL`ID

	NetworkAclId *string `json:"NetworkAclId,omitempty" name:"NetworkAclId"`
	// 是否为 `SNAT` 地址池子网。

	IsRemoteVpcSnat *bool `json:"IsRemoteVpcSnat,omitempty" name:"IsRemoteVpcSnat"`
}

type DeleteAddressTemplateGroupRequest struct {
	*tchttp.BaseRequest

	// IP地址模板集合实例ID，例如：ipmg-90cex8mq。

	AddressTemplateGroupId *string `json:"AddressTemplateGroupId,omitempty" name:"AddressTemplateGroupId"`
}

func (r *DeleteAddressTemplateGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAddressTemplateGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDetectStatesRequest struct {
	*tchttp.BaseRequest

	// 网络探测实例`ID`数组。形如：[`netd-12345678`]

	NetDetectIds []*string `json:"NetDetectIds,omitempty" name:"NetDetectIds"`
	// 过滤条件，参数不支持同时指定NetDetectIds和Filters。
	// <li>net-detect-id - String - （过滤条件）网络探测实例ID，形如：netd-12345678</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNetDetectStatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDetectStatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcResourceDashboardRequest struct {
	*tchttp.BaseRequest

	// Vpc实例ID，例如：vpc-f1xjkw1b。

	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`
}

func (r *DescribeVpcResourceDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcResourceDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCcnAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcceptAttachCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 接受关联实例列表。

	Instances []*CcnInstance `json:"Instances,omitempty" name:"Instances"`
}

func (r *AcceptAttachCcnInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcceptAttachCcnInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubnetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 子网对象。

		Subnet *Subnet `json:"Subnet,omitempty" name:"Subnet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubnetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateIp6AddressesBandwidthRequest struct {
	*tchttp.BaseRequest

	// 需要开通公网访问能力的IPV6地址

	Ip6Addresses []*string `json:"Ip6Addresses,omitempty" name:"Ip6Addresses"`
	// 带宽，单位Mbps。默认是1Mbps

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 网络计费模式。IPV6当前支持"TRAFFIC_POSTPAID_BY_HOUR"，默认是"TRAFFIC_POSTPAID_BY_HOUR"。

	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
}

func (r *AllocateIp6AddressesBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateIp6AddressesBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectAttachCcnInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectAttachCcnInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectAttachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGatewaysRequest struct {
	*tchttp.BaseRequest

	// NAT网关统一 ID，形如：`nat-123xx454`。

	NatGatewayIds []*string `json:"NatGatewayIds,omitempty" name:"NatGatewayIds"`
	// 过滤条件，参数不支持同时指定NatGatewayIds和Filters。
	// <li>nat-gateway-id - String - （过滤条件）协议端口模板实例ID，形如：`nat-123xx454`。</li>
	// <li>vpc-id - String - （过滤条件）私有网络 唯一ID，形如：`vpc-123xx454`。</li>
	// <li>nat-gateway-name - String - （过滤条件）协议端口模板实例ID，形如：`test_nat`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNatGatewaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatGatewaysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcsRequest struct {
	*tchttp.BaseRequest

	// VPC实例ID。形如：vpc-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpcIds和Filters。

	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`
	// 过滤条件，参数不支持同时指定VpcIds和Filters。
	// <li>vpc-name - String - （过滤条件）VPC实例名称。</li>
	// <li>is-default - String - （过滤条件）是否默认VPC。</li>
	// <li>vpc-id - String - （过滤条件）VPC实例ID形如：vpc-f49l6u0z。</li>
	// <li>cidr-block - String - （过滤条件）vpc的cidr。</li>
	// <li>tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例2。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeVpcsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 过滤条件。
	// <li>sregion - String - （过滤条件）源地域，形如：ap-guangzhou。</li>
	// <li>dregion - String - （过滤条件）目的地域，形如：ap-shanghai-bm</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序条件，目前支持带宽（BandwidthLimit）和过期时间（ExpireTime）

	SortedBy *string `json:"SortedBy,omitempty" name:"SortedBy"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序方式，'ASC':升序,'DESC':降序。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *GetCcnRegionBandwidthLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIpv6AddressesAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIpv6AddressesAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIpv6AddressesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCcnResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCcnResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCcnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// 项目id，默认0。可在qcloud控制台项目管理页面查询到。

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 安全组名称，可任意命名，但不得超过60个字符。

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 安全组备注，最多100个字符。

	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`
}

func (r *CreateSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSubnetRequest struct {
	*tchttp.BaseRequest

	// 子网实例ID。可通过DescribeSubnets接口返回值中的SubnetId获取。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *DeleteSubnetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSubnetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBandwidthPackageAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBandwidthPackageAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBandwidthPackageAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTrafficMirrorDirectionRequest struct {
	*tchttp.BaseRequest

	// 流量镜像实例ID

	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`
	// 流量镜像采集方向

	Direction *string `json:"Direction,omitempty" name:"Direction"`
}

func (r *UpdateTrafficMirrorDirectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTrafficMirrorDirectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadCustomerGatewayConfigurationRequest struct {
	*tchttp.BaseRequest

	// VPN网关实例ID。

	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	// VPN通道实例ID。形如：vpnx-f49l6u0z。

	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`
	// 对端网关厂商信息对象，可通过DescribeCustomerGatewayVendors获取。

	CustomerGatewayVendor *CustomerGatewayVendor `json:"CustomerGatewayVendor,omitempty" name:"CustomerGatewayVendor"`
	// 通道接入设备物理接口名称。

	InterfaceName *string `json:"InterfaceName,omitempty" name:"InterfaceName"`
}

func (r *DownloadCustomerGatewayConfigurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadCustomerGatewayConfigurationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceSecurityGroupPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceSecurityGroupPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceSecurityGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressesHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该账户在当前地域指定eip找回的配额限制

		RecoverLimit *int64 `json:"RecoverLimit,omitempty" name:"RecoverLimit"`
		// 该账户在当前地域当月使用指定eip找回的次数

		Recovered *int64 `json:"Recovered,omitempty" name:"Recovered"`
		// 查询EIP的基本信息

		AddressSet *AddressBasic `json:"AddressSet,omitempty" name:"AddressSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressesHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressesHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcRequest struct {
	*tchttp.BaseRequest

	// vpc名称，最大长度不能超过60个字节。

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// vpc的cidr，只能为10.0.0.0/16，172.16.0.0/16，192.168.0.0/16这三个内网网段内。

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 是否开启组播。true: 开启, false: 不开启。

	EnableMulticast *string `json:"EnableMulticast,omitempty" name:"EnableMulticast"`
	// DNS地址，最多支持4个

	DnsServers []*string `json:"DnsServers,omitempty" name:"DnsServers"`
	// 域名

	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *CreateVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIp6IdcInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIp6IdcInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIp6IdcInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAddressTemplateGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IP地址模板集合对象。

		AddressTemplateGroup *AddressTemplateGroup `json:"AddressTemplateGroup,omitempty" name:"AddressTemplateGroup"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAddressTemplateGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAddressTemplateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenBgRsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenBgRsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenBgRsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectGatewayRequest struct {
	*tchttp.BaseRequest

	// 专线网关唯一`ID`，形如：`dcg-9o233uri`。

	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
}

func (r *DeleteDirectConnectGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirectConnectGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewVpnGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewVpnGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroup struct {
	// 项目id，默认0。可在qcloud控制台项目管理页面查询到。

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 安全组实例ID，例如：sg-ohuuioma。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 安全组名称，可任意命名，但不得超过60个字符。

	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	// 安全组备注，最多100个字符。

	SecurityGroupDesc *string `json:"SecurityGroupDesc,omitempty" name:"SecurityGroupDesc"`
	// 是否是默认安全组，默认安全组不支持删除。

	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
	// 安全组创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type CreateDefaultVpcRequest struct {
	*tchttp.BaseRequest

	// 子网所在的可用区ID，不指定将随机选择可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 是否强制返回默认VPC

	Force *bool `json:"Force,omitempty" name:"Force"`
}

func (r *CreateDefaultVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDefaultVpcRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAddressTemplateRequest struct {
	*tchttp.BaseRequest

	// IP地址模板实例ID，例如：ipm-09o5m8kc。

	AddressTemplateId *string `json:"AddressTemplateId,omitempty" name:"AddressTemplateId"`
}

func (r *DeleteAddressTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAddressTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcPrivateIpAddressesRequest struct {
	*tchttp.BaseRequest

	// `VPC`实例`ID`，形如：`vpc-f49l6u0z`。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 内网`IP`地址列表，批量查询单次请求最多支持`10`个。

	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
}

func (r *DescribeVpcPrivateIpAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcPrivateIpAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnRegionBandwidthLimitsTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCcnRegionBandwidthLimitsTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnRegionBandwidthLimitsTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupAttributeRequest struct {
	*tchttp.BaseRequest

	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 安全组名称，可任意命名，但不得超过60个字符。

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 安全组备注，最多100个字符。

	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`
}

func (r *ModifySecurityGroupAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteTablesRequest struct {
	*tchttp.BaseRequest

	// 路由表实例ID，例如：rtb-azd4dt1c。

	RouteTableIds []*string `json:"RouteTableIds,omitempty" name:"RouteTableIds"`
	// 过滤条件，参数不支持同时指定RouteTableIds和Filters。
	// <li>route-table-id - String - （过滤条件）路由表实例ID。</li>
	// <li>route-table-name - String - （过滤条件）路由表名称。</li>
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>association.main - String - （过滤条件）是否主路由表。</li>
	// <li>tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例2。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量。

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 请求对象个数。

	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRouteTablesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteTablesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkInterfaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNetworkInterfaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceTemplateGroup struct {
	// 协议端口模板集合实例ID，例如：ppmg-2klmrefu。

	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId,omitempty" name:"ServiceTemplateGroupId"`
	// 协议端口模板集合名称。

	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName,omitempty" name:"ServiceTemplateGroupName"`
	// 协议端口模板实例ID。

	ServiceTemplateIdSet []*string `json:"ServiceTemplateIdSet,omitempty" name:"ServiceTemplateIdSet"`
	// 创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type CreateSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest

	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 安全组规则集合。

	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *CreateSecurityGroupPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressesHistoryRequest struct {
	*tchttp.BaseRequest

	// 待查询EIP的唯一ID数组，形如eip-xxxxxxx

	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
	// 待查询IP地址列表数组

	AddressIps []*string `json:"AddressIps,omitempty" name:"AddressIps"`
	// EIP类型过滤，包括"WanIP","EIP","AnycastEIP"三种类型。不填该参数默认查询所有EIP类型

	AddressType []*string `json:"AddressType,omitempty" name:"AddressType"`
	// EIP运营商过滤，包括"CMCC","CUCC","CTCC","BGP"四种运营商。不填该参数默认查询所有运营商类型

	InternetServiceProvider []*string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`
}

func (r *DescribeAddressesHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressesHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceVacancyAddressesRequest struct {
	*tchttp.BaseRequest
}

func (r *InquiryPriceVacancyAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceVacancyAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IPSECOptionsSpecification struct {
	// 加密算法，可选值：'3DES-CBC', 'AES-CBC-128', 'AES-CBC-192', 'AES-CBC-256', 'DES-CBC', 'NULL'， 默认为AES-CBC-128

	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" name:"EncryptAlgorithm"`
	// 认证算法：可选值：'MD5', 'SHA1'，默认为

	IntegrityAlgorith *string `json:"IntegrityAlgorith,omitempty" name:"IntegrityAlgorith"`
	// IPsec SA lifetime(s)：单位秒，取值范围：180-604800

	IPSECSaLifetimeSeconds *uint64 `json:"IPSECSaLifetimeSeconds,omitempty" name:"IPSECSaLifetimeSeconds"`
	// PFS：可选值：'NULL', 'DH-GROUP1', 'DH-GROUP2', 'DH-GROUP5', 'DH-GROUP14', 'DH-GROUP24'，默认为NULL

	PfsDhGroup *string `json:"PfsDhGroup,omitempty" name:"PfsDhGroup"`
	// IPsec SA lifetime(KB)：单位KB，取值范围：2560-604800

	IPSECSaLifetimeTraffic *uint64 `json:"IPSECSaLifetimeTraffic,omitempty" name:"IPSECSaLifetimeTraffic"`
}

type DeleteNetDetectRequest struct {
	*tchttp.BaseRequest

	// 网络探测实例`ID`。形如：`netd-12345678`

	NetDetectId *string `json:"NetDetectId,omitempty" name:"NetDetectId"`
}

func (r *DeleteNetDetectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetDetectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateTrafficPackagesRequest struct {
	*tchttp.BaseRequest

	// 无

	TrafficAmount *uint64 `json:"TrafficAmount,omitempty" name:"TrafficAmount"`
	// 无

	TrafficPackageCount *uint64 `json:"TrafficPackageCount,omitempty" name:"TrafficPackageCount"`
}

func (r *InquiryPriceCreateTrafficPackagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceCreateTrafficPackagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressTemplateGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// IP地址模板。

		AddressTemplateGroupSet []*AddressTemplateGroup `json:"AddressTemplateGroupSet,omitempty" name:"AddressTemplateGroupSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressTemplateGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressTemplateGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckBandwidthPackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckBandwidthPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckBandwidthPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户账号属性对象

		AccountAttributeSet []*AccountAttribute `json:"AccountAttributeSet,omitempty" name:"AccountAttributeSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableZoneRequest struct {
	*tchttp.BaseRequest

	// EIP运营商信息，比如"BGP","CMCC","CUCC","CTCC"

	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`
	// EIP可用区，比如"ap-guangzhou-1"

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeAvailableZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetsRequest struct {
	*tchttp.BaseRequest

	// 子网实例ID查询。形如：subnet-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定SubnetIds和Filters。

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// 过滤条件，参数不支持同时指定SubnetIds和Filters。
	// <li>subnet-id - String - （过滤条件）Subnet实例名称。</li>
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>cidr-block - String - （过滤条件）子网网段，形如: 192.168.1.0 。</li>
	// <li>is-default - Boolean - （过滤条件）是否是默认子网。</li>
	// <li>is-remote-vpc-snat - Boolean - （过滤条件）是否为VPC SNAT地址池子网。</li>
	// <li>subnet-name - String - （过滤条件）子网名称。</li>
	// <li>zone - String - （过滤条件）可用区。</li>
	// <li>tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例2。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSubnetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubnetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetVpnConnectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetVpnConnectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetVpnConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubnetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新创建的子网列表。

		SubnetSet []*Subnet `json:"SubnetSet,omitempty" name:"SubnetSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubnetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestinationIpPortTranslationNatRule struct {
	// 网络协议，可选值：`TCP`、`UDP`。

	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`
	// 弹性IP。

	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`
	// 公网端口。

	PublicPort *uint64 `json:"PublicPort,omitempty" name:"PublicPort"`
	// 内网地址。

	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	// 内网端口。

	PrivatePort *uint64 `json:"PrivatePort,omitempty" name:"PrivatePort"`
	// NAT网关转发规则描述。

	Description *string `json:"Description,omitempty" name:"Description"`
}

type InquiryPriceAllocateAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceAllocateAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceAllocateAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssistantCidrRequest struct {
	*tchttp.BaseRequest

	// `VPC`实例`ID`。形如：`vpc-6v2ht8q5`

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 待添加的负载CIDR。CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]

	NewCidrBlocks []*string `json:"NewCidrBlocks,omitempty" name:"NewCidrBlocks"`
	// 待删除的负载CIDR。CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]

	OldCidrBlocks []*string `json:"OldCidrBlocks,omitempty" name:"OldCidrBlocks"`
}

func (r *ModifyAssistantCidrRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssistantCidrRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupPolicy struct {
	// 安全组规则索引号。

	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`
	// 协议, 取值: TCP,UDP, ICMP。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 端口(all, 离散port,  range)。

	Port *string `json:"Port,omitempty" name:"Port"`
	// 协议端口ID或者协议端口组ID。ServiceTemplate和Protocol+Port互斥。

	ServiceTemplate *ServiceTemplateSpecification `json:"ServiceTemplate,omitempty" name:"ServiceTemplate"`
	// 网段或IP(互斥)。

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 网段或IPv6(互斥)。

	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`
	// 安全组实例ID，例如：sg-ohuuioma。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// IP地址ID或者ID地址组ID。

	AddressTemplate *AddressTemplateSpecification `json:"AddressTemplate,omitempty" name:"AddressTemplate"`
	// ACCEPT 或 DROP。

	Action *string `json:"Action,omitempty" name:"Action"`
	// 安全组规则描述。

	PolicyDescription *string `json:"PolicyDescription,omitempty" name:"PolicyDescription"`
	// 安全组最近修改时间。

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type DescribeAddressBandwidthConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressBandwidthConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressBandwidthConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDetectsRequest struct {
	*tchttp.BaseRequest

	// 网络探测实例`ID`数组。形如：[`netd-12345678`]

	NetDetectIds []*string `json:"NetDetectIds,omitempty" name:"NetDetectIds"`
	// 过滤条件，参数不支持同时指定NetDetectIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-12345678</li>
	// <li>net-detect-id - String - （过滤条件）网络探测实例ID，形如：netd-12345678</li>
	// <li>subnet-id - String - （过滤条件）子网实例ID，形如：subnet-12345678</li>
	// <li>net-detect-name - String - （过滤条件）网络探测名称</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNetDetectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDetectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrafficMirrorsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流量镜像实例信息

		TrafficMirrorSet []*TrafficMirror `json:"TrafficMirrorSet,omitempty" name:"TrafficMirrorSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrafficMirrorsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrafficMirrorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressBandwidthConfigsRequest struct {
	*tchttp.BaseRequest

	// 1

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeAddressBandwidthConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressBandwidthConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTemplateGroupsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>service-template-group-name - String - （过滤条件）协议端口模板集合名称。</li>
	// <li>service-template-group-id - String - （过滤条件）协议端口模板集合实例ID，例如：ppmg-e6dy460g。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeServiceTemplateGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceTemplateGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCcnRegionBandwidthLimitsTypeRequest struct {
	*tchttp.BaseRequest

	// 云联网实例ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 云联网限速类型，INTER_REGION_LIMIT：地域间限速，OUTER_REGION_LIMIT：地域出口限速。

	BandwidthLimitType *string `json:"BandwidthLimitType,omitempty" name:"BandwidthLimitType"`
}

func (r *ModifyCcnRegionBandwidthLimitsTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCcnRegionBandwidthLimitsTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecurityGroupAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceRouteTableAssociationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceRouteTableAssociationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceRouteTableAssociationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcceptAttachCcnInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AcceptAttachCcnInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcceptAttachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RouteTable struct {
	// VPC实例ID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 路由表实例ID，例如：rtb-azd4dt1c。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 路由表名称。

	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
	// 路由表关联关系。

	AssociationSet []*RouteTableAssociation `json:"AssociationSet,omitempty" name:"AssociationSet"`
	// 路由表策略集合。

	RouteSet []*Route `json:"RouteSet,omitempty" name:"RouteSet"`
	// 是否默认路由表。

	Main *bool `json:"Main,omitempty" name:"Main"`
	// 创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type ModifyAddressAttributeRequest struct {
	*tchttp.BaseRequest

	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：`eip-11112222`。

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// 修改后的 EIP 名称。长度上限为20个字符。

	AddressName *string `json:"AddressName,omitempty" name:"AddressName"`
	// 设定EIP是否直通，"TRUE"表示直通，"FALSE"表示非直通。注意该参数仅对EIP直通功能可见的用户可以设定。

	EipDirectConnection *string `json:"EipDirectConnection,omitempty" name:"EipDirectConnection"`
	// 1

	CascadeRelease *bool `json:"CascadeRelease,omitempty" name:"CascadeRelease"`
}

func (r *ModifyAddressAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAddressAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BandwidthPackage struct {
	// 带宽包唯一标识Id

	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
	// 带宽包类型，包括'BGP','SINGLEISP','ANYCAST'

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// 带宽包计费类型，包括'TOP5_POSTPAID_BY_MONTH'和'PERCENT95_POSTPAID_BY_MONTH'

	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
	// 带宽包名称

	BandwidthPackageName *string `json:"BandwidthPackageName,omitempty" name:"BandwidthPackageName"`
	// 带宽包创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 带宽包状态，包括'CREATING','CREATED','DELETING','DELETED'

	Status *string `json:"Status,omitempty" name:"Status"`
	// 带宽包资源信息

	ResourceSet []*Resource `json:"ResourceSet,omitempty" name:"ResourceSet"`
	// 带宽包限速大小。单位：Mbps，-1表示不限速。

	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
}

type DescribeIp6TranslatorsRequest struct {
	*tchttp.BaseRequest

	// IPV6转换实例唯一ID数组，形如ip6-xxxxxxxx

	Ip6TranslatorIds []*string `json:"Ip6TranslatorIds,omitempty" name:"Ip6TranslatorIds"`
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`Ip6TranslatorIds`和`Filters`。详细的过滤条件如下：
	// <li> ip6-translator-id - String - 是否必填：否 - （过滤条件）按照IPV6转换实例的唯一ID过滤,形如ip6-xxxxxxx。</li>
	// <li> ip6-translator-vip6 - String - 是否必填：否 - （过滤条件）按照IPV6地址过滤。不支持模糊过滤。</li>
	// <li> ip6-translator-name - String - 是否必填：否 - （过滤条件）按照IPV6转换实例名称过滤。不支持模糊过滤。</li>
	// <li> ip6-translator-status - String - 是否必填：否 - （过滤条件）按照IPV6转换实例的状态过滤。状态取值范围为"CREATING","RUNNING","DELETING","MODIFYING"

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeIp6TranslatorsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIp6TranslatorsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachCcnInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachCcnInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Address struct {
	// `EIP`的`ID`，是`EIP`的唯一标识。

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// `EIP`名称。

	AddressName *string `json:"AddressName,omitempty" name:"AddressName"`
	// `EIP`状态，包含'CREATING'(创建中),'BINDING'(绑定中),'BIND'(已绑定),'UNBINDING'(解绑中),'UNBIND'(已解绑),'OFFLINING'(释放中),'BIND_ENI'(绑定悬空弹性网卡)

	AddressStatus *string `json:"AddressStatus,omitempty" name:"AddressStatus"`
	// 外网IP地址

	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`
	// 绑定的资源实例`ID`。可能是一个`CVM`，`NAT`。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 绑定的弹性网卡ID

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	// 绑定的资源内网ip

	PrivateAddressIp *string `json:"PrivateAddressIp,omitempty" name:"PrivateAddressIp"`
	// 资源隔离状态。true表示eip处于隔离状态，false表示资源处于未隔离状态

	IsArrears *bool `json:"IsArrears,omitempty" name:"IsArrears"`
	// 资源封堵状态。true表示eip处于封堵状态，false表示eip处于未封堵状态

	IsBlocked *bool `json:"IsBlocked,omitempty" name:"IsBlocked"`
	// eip是否支持直通模式。true表示eip支持直通模式，false表示资源不支持直通模式

	IsEipDirectConnection *bool `json:"IsEipDirectConnection,omitempty" name:"IsEipDirectConnection"`
	// eip资源类型，包括"CalcIP","WanIP","EIP","AnycastEIP"。其中"CalcIP"表示设备ip，“WanIP”表示普通公网ip，“EIP”表示弹性公网ip，“AnycastEip”表示加速EIP

	AddressType *string `json:"AddressType,omitempty" name:"AddressType"`
	// eip是否在解绑后自动释放。true表示eip将会在解绑后自动释放，false表示eip在解绑后不会自动释放

	CascadeRelease *bool `json:"CascadeRelease,omitempty" name:"CascadeRelease"`
}

type DefaultVpcSubnet struct {
	// 默认VpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 默认SubnetId

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type CreateTrafficMirrorRequest struct {
	*tchttp.BaseRequest

	// VPC实例ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 流量镜像名字

	TrafficMirrorName *string `json:"TrafficMirrorName,omitempty" name:"TrafficMirrorName"`
	// 流量镜像描述

	TrafficMirrorDescribe *string `json:"TrafficMirrorDescribe,omitempty" name:"TrafficMirrorDescribe"`
	// 流量镜像状态, 支持RUNNING/STOPED

	State *string `json:"State,omitempty" name:"State"`
	// 流量镜像采集方向，支持EGRESS/INGRESS/ALL

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 流量镜像的采集对象，支持vpc_xxxx, subnet_xxxx, eni_xxxx

	CollectorSrcs []*string `json:"CollectorSrcs,omitempty" name:"CollectorSrcs"`
	// 流量镜像过滤的natgw实例

	NatId *string `json:"NatId,omitempty" name:"NatId"`
	// 需要过滤的五元组规则

	CollectorNormalFilters []*TrafficMirrorFilter `json:"CollectorNormalFilters,omitempty" name:"CollectorNormalFilters"`
	// 流量镜像的目的地址

	CollectorTarget *TrafficMirrorTarget `json:"CollectorTarget,omitempty" name:"CollectorTarget"`
}

func (r *CreateTrafficMirrorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTrafficMirrorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZoneInfo struct {
	// 可用区名称，例如，ap-guangzhou-1

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区描述，例如，广州一区

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 可用区ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区状态，目前只有两种状态，分别是UNAVAILABLE和AVAILABLE

	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`
}

type DescribeBandwidthPackagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的带宽包数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 描述带宽包详细信息

		BandwidthPackageSet []*BandwidthPackage `json:"BandwidthPackageSet,omitempty" name:"BandwidthPackageSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBandwidthPackagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBandwidthPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateTrafficPackagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateTrafficPackagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceCreateTrafficPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewCcnBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 商品价格。

		Price *ItemPrice `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewCcnBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewCcnBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBandwidthPackageRequest struct {
	*tchttp.BaseRequest

	// 带宽包类型，包括'BGP'，'SINGLEISP'，'ANYCAST'

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// 带宽包计费类型，包括‘TOP5_POSTPAID_BY_MONTH’，‘PERCENT95_POSTPAID_BY_MONTH’

	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
	// 带宽包名字

	BandwidthPackageName *string `json:"BandwidthPackageName,omitempty" name:"BandwidthPackageName"`
	// 带宽包数量(非上移账户只能填1)

	BandwidthPackageCount *uint64 `json:"BandwidthPackageCount,omitempty" name:"BandwidthPackageCount"`
	// 带宽包限速大小。单位：Mbps，-1表示不限速。

	InternetMaxBandwidth *int64 `json:"InternetMaxBandwidth,omitempty" name:"InternetMaxBandwidth"`
	// 需要关联的标签列表。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 带宽包协议类型。当前支持'ipv4'和'ipv6'协议带宽包，默认值是'ipv4'。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *CreateBandwidthPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBandwidthPackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnGatewayQuotaRequest struct {
	*tchttp.BaseRequest

	// VPC实例ID。形如：vpc-f49l6u0z。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// VPN网关实例ID。形如：vpngw-f49l6u0z。

	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
}

func (r *DescribeVpnGatewayQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnGatewayQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetDetect struct {
	// `VPC`实例`ID`。形如：`vpc-12345678`

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// `VPC`实例名称。

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 子网实例ID。形如：subnet-12345678。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 子网实例名称。

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 网络探测实例ID。形如：netd-12345678。

	NetDetectId *string `json:"NetDetectId,omitempty" name:"NetDetectId"`
	// 网络探测名称，最大长度不能超过60个字节。

	NetDetectName *string `json:"NetDetectName,omitempty" name:"NetDetectName"`
	// 探测目的IPv4地址数组，最多两个。

	DetectDestinationIp []*string `json:"DetectDestinationIp,omitempty" name:"DetectDestinationIp"`
	// 系统自动分配的探测源IPv4数组。长度为2。

	DetectSourceIp []*string `json:"DetectSourceIp,omitempty" name:"DetectSourceIp"`
	// 下一跳类型，目前我们支持的类型有：
	// VPN：VPN网关；
	// DIRECTCONNECT：专线网关；
	// PEERCONNECTION：对等连接；
	// NAT：NAT网关；
	// NORMAL_CVM：普通云主机；

	NextHopType *string `json:"NextHopType,omitempty" name:"NextHopType"`
	// 下一跳目的网关，取值与“下一跳类型”相关：
	// 下一跳类型为VPN，取值VPN网关ID，形如：vpngw-12345678；
	// 下一跳类型为DIRECTCONNECT，取值专线网关ID，形如：dcg-12345678；
	// 下一跳类型为PEERCONNECTION，取值对等连接ID，形如：pcx-12345678；
	// 下一跳类型为NAT，取值Nat网关，形如：nat-12345678；
	// 下一跳类型为NORMAL_CVM，取值云主机IPv4地址，形如：10.0.0.12；

	NextHopDestination *string `json:"NextHopDestination,omitempty" name:"NextHopDestination"`
	// 下一跳网关名称。

	NextHopName *string `json:"NextHopName,omitempty" name:"NextHopName"`
	// 网络探测描述。

	NetDetectDescription *string `json:"NetDetectDescription,omitempty" name:"NetDetectDescription"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type SecurityPolicyDatabase struct {
	// 本端网段

	LocalCidrBlock *string `json:"LocalCidrBlock,omitempty" name:"LocalCidrBlock"`
	// 对端网段

	RemoteCidrBlock []*string `json:"RemoteCidrBlock,omitempty" name:"RemoteCidrBlock"`
}

type DeleteCustomerGatewayRequest struct {
	*tchttp.BaseRequest

	// 对端网关ID，例如：cgw-2wqq41m9，可通过DescribeCustomerGateways接口查询对端网关。

	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`
}

func (r *DeleteCustomerGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomerGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Message struct {
	// 错误编码

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 错误提示信息

	Error *string `json:"Error,omitempty" name:"Error"`
}

type Rule struct {
	// 防火墙规则优先级

	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 源地址网段

	SourceCidr *string `json:"SourceCidr,omitempty" name:"SourceCidr"`
	// 源端口范围

	SourcePort *string `json:"SourcePort,omitempty" name:"SourcePort"`
	// 目的地址网段

	DestCidr *string `json:"DestCidr,omitempty" name:"DestCidr"`
	// 目的端口范围

	DestPort *string `json:"DestPort,omitempty" name:"DestPort"`
	// 策略

	Action *string `json:"Action,omitempty" name:"Action"`
	// 规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type BandwidthSpecification struct {
	// Anycast入口地域名称。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 带宽值。单位 Mbps。

	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
}

type HaVipAssociateAddressIpRequest struct {
	*tchttp.BaseRequest

	// `HAVIP`唯一`ID`，形如：`havip-9o233uri`。必须是没有绑定`EIP`的`HAVIP`

	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`
	// 弹性公网`IP`。必须是没有绑定`HAVIP`的`EIP`

	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`
}

func (r *HaVipAssociateAddressIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *HaVipAssociateAddressIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIp6IdcInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeIp6IdcInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIp6IdcInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTrafficMirrorAttributeRequest struct {
	*tchttp.BaseRequest

	// 流量镜像实例ID

	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`
	// 流量镜像实例名称

	TrafficMirrorName *string `json:"TrafficMirrorName,omitempty" name:"TrafficMirrorName"`
	// 流量镜像实例描述信息

	TrafficMirrorDescription *string `json:"TrafficMirrorDescription,omitempty" name:"TrafficMirrorDescription"`
}

func (r *ModifyTrafficMirrorAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTrafficMirrorAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnConnectionRequest struct {
	*tchttp.BaseRequest

	// VPN网关实例ID。

	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	// VPN通道实例ID。形如：vpnx-f49l6u0z。

	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`
}

func (r *DeleteVpnConnectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpnConnectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 子网对象。

		SubnetSet []*Subnet `json:"SubnetSet,omitempty" name:"SubnetSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubnetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectGatewayRequest struct {
	*tchttp.BaseRequest

	// 专线网关名称

	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitempty" name:"DirectConnectGatewayName"`
	// 关联网络类型，可选值：
	// <li>VPC - 私有网络</li>
	// <li>CCN - 云联网</li>

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// <li>NetworkType 为 VPC 时，这里传值为私有网络实例ID</li>
	// <li>NetworkType 为 CCN 时，这里传值为云联网实例ID</li>

	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" name:"NetworkInstanceId"`
	// 网关类型，可选值：
	// <li>NORMAL - （默认）标准型，注：云联网只支持标准型</li>
	// <li>NAT - NAT型</li>NAT类型支持网络地址转换配置，类型确定后不能修改；一个私有网络可以创建一个NAT类型的专线网关和一个非NAT类型的专线网关

	GatewayType *string `json:"GatewayType,omitempty" name:"GatewayType"`
}

func (r *CreateDirectConnectGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirectConnectGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyBandwidthPackageBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceModifyBandwidthPackageBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceModifyBandwidthPackageBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupLimitSet struct {
	// 每个项目每个地域可创建安全组数

	SecurityGroupLimit *uint64 `json:"SecurityGroupLimit,omitempty" name:"SecurityGroupLimit"`
	// 安全组下的最大规则数

	SecurityGroupPolicyLimit *uint64 `json:"SecurityGroupPolicyLimit,omitempty" name:"SecurityGroupPolicyLimit"`
	// 安全组展开后的规则数

	SecurityGroupExtendedPolicyLimit *uint64 `json:"SecurityGroupExtendedPolicyLimit,omitempty" name:"SecurityGroupExtendedPolicyLimit"`
	// 安全组下嵌套安全组规则数

	ReferedSecurityGroupLimit *uint64 `json:"ReferedSecurityGroupLimit,omitempty" name:"ReferedSecurityGroupLimit"`
	// 单安全组关联实例数

	SecurityGroupInstanceLimit *uint64 `json:"SecurityGroupInstanceLimit,omitempty" name:"SecurityGroupInstanceLimit"`
	// 实例关联安全组数

	InstanceSecurityGroupLimit *uint64 `json:"InstanceSecurityGroupLimit,omitempty" name:"InstanceSecurityGroupLimit"`
	// 安全组下嵌套安全组关联实例数-云主机或弹性网卡

	SecurityGroupReferedCvmAndEniLimit *uint64 `json:"SecurityGroupReferedCvmAndEniLimit,omitempty" name:"SecurityGroupReferedCvmAndEniLimit"`
	// 安全组下嵌套安全组关联实例数-数据库、LB等

	SecurityGroupReferedSvcLimit *uint64 `json:"SecurityGroupReferedSvcLimit,omitempty" name:"SecurityGroupReferedSvcLimit"`
}

type DescribeAddressQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAddressQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectAttachCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 拒绝关联实例列表。

	Instances []*CcnInstance `json:"Instances,omitempty" name:"Instances"`
}

func (r *RejectAttachCcnInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectAttachCcnInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnRoute struct {
	// 路由策略ID

	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`
	// 目的端

	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
	// 下一跳类型（关联实例类型），所有类型：VPC、DIRECTCONNECT

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 下一跳（关联实例）

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 下一跳名称（关联实例名称）

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 下一跳所属地域（关联实例所属地域）

	InstanceRegion *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 路由是否启用

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 关联实例所属UIN（根账号）

	InstanceUin *string `json:"InstanceUin,omitempty" name:"InstanceUin"`
}

type DescribeAddressAvailabilityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressAvailabilityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressAvailabilityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateServiceTemplateRequest struct {
	*tchttp.BaseRequest

	// 协议端口模板名称

	ServiceTemplateName *string `json:"ServiceTemplateName,omitempty" name:"ServiceTemplateName"`
	// 支持单个端口、多个端口、连续端口及所有端口，协议支持：TCP、UDP、ICMP、GRE 协议。

	Services []*string `json:"Services,omitempty" name:"Services"`
}

func (r *CreateServiceTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateServiceTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupAssociationStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全组关联实例统计。

		SecurityGroupAssociationStatisticsSet []*SecurityGroupAssociationStatistics `json:"SecurityGroupAssociationStatisticsSet,omitempty" name:"SecurityGroupAssociationStatisticsSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupAssociationStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupAssociationStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewBandwidthPackageRequest struct {
	*tchttp.BaseRequest

	// 带宽包ID

	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
	// 续费时长(月)

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
}

func (r *InquiryPriceRenewBandwidthPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewBandwidthPackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnLimitsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：
	// type- Int -（过滤条件）云联网配额类型。
	// 1：每个开发商可创建的云联网数，2：每个云联网绑定的实例个数，3：单个云联网支持的路由条目。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCcnLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayFlowMonitorDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 网关流量监控明细。

		GatewayFlowMonitorDetailSet []*GatewayFlowMonitorDetail `json:"GatewayFlowMonitorDetailSet,omitempty" name:"GatewayFlowMonitorDetailSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGatewayFlowMonitorDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGatewayFlowMonitorDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIp6AddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的 IPV6 数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// IPV6 详细信息列表。

		AddressSet []*Address `json:"AddressSet,omitempty" name:"AddressSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIp6AddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIp6AddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnGatewayQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// VPN网关配额对象。

		VpnGatewayQuotaSet []*VpnGatewayQuota `json:"VpnGatewayQuotaSet,omitempty" name:"VpnGatewayQuotaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpnGatewayQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnGatewayQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomerGatewayAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomerGatewayAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomerGatewayAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddIp6RulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IPV6转换规则唯一ID数组，形如rule6-xxxxxxxx

		Ip6RuleSet []*string `json:"Ip6RuleSet,omitempty" name:"Ip6RuleSet"`
		// TaskId，运维使用，无实际业务含义

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddIp6RulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddIp6RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest

	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 安全组规则集合。 SecurityGroupPolicySet对象必须同时指定新的出（Egress）入（Ingress）站规则。 SecurityGroupPolicy对象不支持自定义索引（PolicyIndex）。

	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
	// 排序安全组标识。值为True时，支持安全组排序；SortPolicys不存在或SortPolicys为False时，为修改安全组规则。

	SortPolicys *bool `json:"SortPolicys,omitempty" name:"SortPolicys"`
}

func (r *ModifySecurityGroupPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIp6TranslatorsRequest struct {
	*tchttp.BaseRequest

	// 转换实例名称

	Ip6TranslatorName *string `json:"Ip6TranslatorName,omitempty" name:"Ip6TranslatorName"`
	// 创建转换实例数量，默认是1个

	Ip6TranslatorCount *int64 `json:"Ip6TranslatorCount,omitempty" name:"Ip6TranslatorCount"`
	// 转换实例运营商属性，可取"CMCC","CTCC","CUCC","BGP"

	Ip6InternetServiceProvider *string `json:"Ip6InternetServiceProvider,omitempty" name:"Ip6InternetServiceProvider"`
	// 创建转换实例所属IDC

	IdcList []*string `json:"IdcList,omitempty" name:"IdcList"`
}

func (r *CreateIp6TranslatorsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIp6TranslatorsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest

	// CVM主机的服务器ID

	InstanceUUid *string `json:"InstanceUUid,omitempty" name:"InstanceUUid"`
}

func (r *CheckSecurityGroupPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSecurityGroupPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetTrafficMirrorFilterRequest struct {
	*tchttp.BaseRequest

	// 流量镜像实例ID

	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`
	// 流量镜像需要过滤的natgw实例ID

	NatId *string `json:"NatId,omitempty" name:"NatId"`
	// 流量镜像需要过滤的五元组规则

	CollectorNormalFilters []*TrafficMirrorFilter `json:"CollectorNormalFilters,omitempty" name:"CollectorNormalFilters"`
}

func (r *ResetTrafficMirrorFilterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetTrafficMirrorFilterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CCN struct {
	// 云联网唯一ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 云联网名称

	CcnName *string `json:"CcnName,omitempty" name:"CcnName"`
	// 云联网描述信息

	CcnDescription *string `json:"CcnDescription,omitempty" name:"CcnDescription"`
	// 关联实例数量

	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 实例状态， 'ISOLATED': 隔离中（欠费停服），'AVAILABLE'：运行中。

	State *string `json:"State,omitempty" name:"State"`
	// 实例服务质量，’PT’：白金，'AU'：金，'AG'：银。

	QosLevel *string `json:"QosLevel,omitempty" name:"QosLevel"`
	// 付费类型，PREPAID为预付费，POSTPAID为后付费。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 限速类型，INTER_REGION_LIMIT为地域间限速；OUTER_REGION_LIMIT为地域出口限速。

	BandwidthLimitType *string `json:"BandwidthLimitType,omitempty" name:"BandwidthLimitType"`
}

type CreateCcnRequest struct {
	*tchttp.BaseRequest

	// CCN名称，最大长度不能超过60个字节。

	CcnName *string `json:"CcnName,omitempty" name:"CcnName"`
	// CCN描述信息，最大长度不能超过100个字节。

	CcnDescription *string `json:"CcnDescription,omitempty" name:"CcnDescription"`
	// CCN服务质量，'PT'：白金，'AU'：金，'AG'：银，默认为‘AU’。

	QosLevel *string `json:"QosLevel,omitempty" name:"QosLevel"`
	// 计费模式，PREPAID：表示预付费，即包年包月，POSTPAID：表示后付费，即按量计费。默认：POSTPAID。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 限速类型，OUTER_REGION_LIMIT表示地域出口限速，INTER_REGION_LIMIT为地域间限速，默认为OUTER_REGION_LIMIT

	BandwidthLimitType *string `json:"BandwidthLimitType,omitempty" name:"BandwidthLimitType"`
}

func (r *CreateCcnRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCcnRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerGatewayVendorsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCustomerGatewayVendorsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerGatewayVendorsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流日志信息

		FlowLog []*FlowLog `json:"FlowLog,omitempty" name:"FlowLog"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInterfaceLimitRequest struct {
	*tchttp.BaseRequest

	// 要查询的CVM实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeNetworkInterfaceLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkInterfaceLimitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnGatewayRequest struct {
	*tchttp.BaseRequest

	// VPN网关实例ID。

	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
}

func (r *DeleteVpnGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpnGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpnGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PrivateIp6AddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PrivateIp6AddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PrivateIp6AddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddressChargePrepaid struct {
	// 购买实例的时长

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 自动续费标志

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type TrafficMirror struct {
	// VPC实例ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 流量镜像实例

	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`
	// TrafficMirrorName

	TrafficMirrorName *string `json:"TrafficMirrorName,omitempty" name:"TrafficMirrorName"`
	// 流量镜像描述

	TrafficMirrorDescribe *string `json:"TrafficMirrorDescribe,omitempty" name:"TrafficMirrorDescribe"`
	// 流量镜像状态

	State *string `json:"State,omitempty" name:"State"`
	// 流量镜像采集方向

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 流量镜像采集对象

	CollectorSrcs []*string `json:"CollectorSrcs,omitempty" name:"CollectorSrcs"`
	// 流量镜像过滤的nat网关实例ID

	NatId *string `json:"NatId,omitempty" name:"NatId"`
	// 流量镜像过滤的五元组规则

	CollectorNormalFilters []*TrafficMirrorFilter `json:"CollectorNormalFilters,omitempty" name:"CollectorNormalFilters"`
	// 流量镜接收目标

	CollectorTarget *TrafficMirrorTarget `json:"CollectorTarget,omitempty" name:"CollectorTarget"`
	// 流量镜像创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type TrafficMirrorTarget struct {
	// 流量镜像的接收IP

	TargetIps []*string `json:"TargetIps,omitempty" name:"TargetIps"`
	// 流量镜像接收IP组，均衡规则，支持ENI/FIVE_TUPLE_FLOW

	AlgHash *string `json:"AlgHash,omitempty" name:"AlgHash"`
}

type CreateServiceTemplateGroupRequest struct {
	*tchttp.BaseRequest

	// 协议端口模板集合名称

	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName,omitempty" name:"ServiceTemplateGroupName"`
	// 协议端口模板实例ID，例如：ppm-4dw6agho。

	ServiceTemplateIds []*string `json:"ServiceTemplateIds,omitempty" name:"ServiceTemplateIds"`
}

func (r *CreateServiceTemplateGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateServiceTemplateGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateBandwidthPackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateBandwidthPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceCreateBandwidthPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressesBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务TaskId。可以使用[DescribeTaskResult](https://cloud.tencent.com/document/api/215/36271)接口查询任务状态。

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressesBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAddressesBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetVpnGatewayInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetVpnGatewayInternetMaxBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetVpnGatewayInternetMaxBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36。

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 自动续费标识。取值范围： NOTIFY_AND_AUTO_RENEW：通知过期且自动续费， NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费。默认：NOTIFY_AND_MANUAL_RENEW

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type ModifyServiceTemplateAttributeRequest struct {
	*tchttp.BaseRequest

	// 协议端口模板实例ID，例如：ppm-529nwwj8。

	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`
	// 协议端口模板名称。

	ServiceTemplateName *string `json:"ServiceTemplateName,omitempty" name:"ServiceTemplateName"`
	// 支持单个端口、多个端口、连续端口及所有端口，协议支持：TCP、UDP、ICMP、GRE 协议。

	Services []*string `json:"Services,omitempty" name:"Services"`
}

func (r *ModifyServiceTemplateAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServiceTemplateAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnConnectionAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpnConnectionAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpnConnectionAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableCcnRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableCcnRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIpv6AddressesAttributeRequest struct {
	*tchttp.BaseRequest

	// 弹性网卡实例`ID`，形如：`eni-m6dyj72l`。

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	// 指定的内网IPv6`地址信息。

	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`
}

func (r *ModifyIpv6AddressesAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIpv6AddressesAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveBandwidthPackageResourcesRequest struct {
	*tchttp.BaseRequest

	// 带宽包唯一标识ID，形如'bwp-xxxx'

	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
	// 资源类型，包括‘Address’, ‘LoadBalance’

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源Id，形如'eip-xxxx', 'lb-xxxx'

	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
}

func (r *RemoveBandwidthPackageResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveBandwidthPackageResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterObject struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DeleteDirectConnectGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDirectConnectGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirectConnectGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServiceTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressTemplateAttributeRequest struct {
	*tchttp.BaseRequest

	// IP地址模板实例ID，例如：ipm-mdunqeb6。

	AddressTemplateId *string `json:"AddressTemplateId,omitempty" name:"AddressTemplateId"`
	// IP地址模板名称。

	AddressTemplateName *string `json:"AddressTemplateName,omitempty" name:"AddressTemplateName"`
	// 地址信息，支持 IP、CIDR、IP 范围。

	Addresses []*string `json:"Addresses,omitempty" name:"Addresses"`
}

func (r *ModifyAddressTemplateAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAddressTemplateAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetDetectRequest struct {
	*tchttp.BaseRequest

	// 网络探测实例`ID`。形如：`netd-12345678`

	NetDetectId *string `json:"NetDetectId,omitempty" name:"NetDetectId"`
	// 网络探测名称，最大长度不能超过60个字节。

	NetDetectName *string `json:"NetDetectName,omitempty" name:"NetDetectName"`
	// 探测目的IPv4地址数组，最多两个。

	DetectDestinationIp []*string `json:"DetectDestinationIp,omitempty" name:"DetectDestinationIp"`
	// 下一跳类型，目前我们支持的类型有：
	// VPN：VPN网关；
	// DIRECTCONNECT：专线网关；
	// PEERCONNECTION：对等连接；
	// NAT：NAT网关；
	// NORMAL_CVM：普通云主机；

	NextHopType *string `json:"NextHopType,omitempty" name:"NextHopType"`
	// 下一跳目的网关，取值与“下一跳类型”相关：
	// 下一跳类型为VPN，取值VPN网关ID，形如：vpngw-12345678；
	// 下一跳类型为DIRECTCONNECT，取值专线网关ID，形如：dcg-12345678；
	// 下一跳类型为PEERCONNECTION，取值对等连接ID，形如：pcx-12345678；
	// 下一跳类型为NAT，取值Nat网关，形如：nat-12345678；
	// 下一跳类型为NORMAL_CVM，取值云主机IPv4地址，形如：10.0.0.12；

	NextHopDestination *string `json:"NextHopDestination,omitempty" name:"NextHopDestination"`
	// 网络探测描述。

	NetDetectDescription *string `json:"NetDetectDescription,omitempty" name:"NetDetectDescription"`
}

func (r *ModifyNetDetectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetDetectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCcnBandwidthRenewFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCcnBandwidthRenewFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCcnBandwidthRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RsList struct {
	// RsWeight

	RsWeight *int64 `json:"RsWeight,omitempty" name:"RsWeight"`
	// VpcId

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// RsPort

	RsPort *int64 `json:"RsPort,omitempty" name:"RsPort"`
	// RsIp

	RsIp *string `json:"RsIp,omitempty" name:"RsIp"`
	// Ipv6Id

	Ipv6Id *int64 `json:"Ipv6Id,omitempty" name:"Ipv6Id"`
	// OldRsIp

	OldRsIp *string `json:"OldRsIp,omitempty" name:"OldRsIp"`
	// OldRsPort

	OldRsPort *int64 `json:"OldRsPort,omitempty" name:"OldRsPort"`
	// HostIp

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// OldHostIp

	OldHostIp *string `json:"OldHostIp,omitempty" name:"OldHostIp"`
}

type CreateHaVipRequest struct {
	*tchttp.BaseRequest

	// `HAVIP`所在私有网络`ID`。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// `HAVIP`所在子网`ID`。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// `HAVIP`名称。

	HaVipName *string `json:"HaVipName,omitempty" name:"HaVipName"`
	// 指定虚拟IP地址，必须在`VPC`网段内且未被占用。不指定则自动分配。

	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

func (r *CreateHaVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHaVipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnGatewayAttributeRequest struct {
	*tchttp.BaseRequest

	// VPN网关实例ID。

	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	// VPN网关名称，最大长度不能超过60个字节。

	VpnGatewayName *string `json:"VpnGatewayName,omitempty" name:"VpnGatewayName"`
	// VPN网关计费模式，目前只支持预付费（即包年包月）到后付费（即按量计费）的转换。即参数只支持：POSTPAID_BY_HOUR。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

func (r *ModifyVpnGatewayAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpnGatewayAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCcnRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *DeleteCcnRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCcnRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCreateCcnBandwidthDealResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单详情。

		Deal *string `json:"Deal,omitempty" name:"Deal"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCreateCcnBandwidthDealResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCreateCcnBandwidthDealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateAddressesRequest struct {
	*tchttp.BaseRequest

	// EIP数量。默认值：1。

	AddressCount *int64 `json:"AddressCount,omitempty" name:"AddressCount"`
	// EIP线路类型。默认值：BGP。
	// <ul style="margin:0"><li>已开通静态单线IP白名单的用户，可选值：<ul><li>CMCC：中国移动</li>
	// <li>CTCC：中国电信</li>
	// <li>CUCC：中国联通</li></ul>注意：仅部分地域支持静态单线IP。</li></ul>

	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`
	// 内部参数，用于指定集群申请EIP。

	TgwGroup *string `json:"TgwGroup,omitempty" name:"TgwGroup"`
	// 内部参数。

	ApplySilence *int64 `json:"ApplySilence,omitempty" name:"ApplySilence"`
	// EIP计费方式。
	// <ul style="margin:0"><li>已开通带宽上移白名单的用户，可选值：<ul><li>BANDWIDTH_PACKAGE：[共享带宽包](https://cloud.tencent.com/document/product/684/15255)付费（需额外开通共享带宽包白名单）</li>
	// <li>BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费</li>
	// <li>TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费</li></ul>默认值：TRAFFIC_POSTPAID_BY_HOUR。</li>
	// <li>未开通带宽上移白名单的用户，EIP计费方式与其绑定的实例的计费方式一致，无需传递此参数。</li></ul>

	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
	// EIP出带宽上限，单位：Mbps。
	// <ul style="margin:0"><li>已开通带宽上移白名单的用户，可选值范围取决于EIP计费方式：<ul><li>BANDWIDTH_PACKAGE：1 Mbps 至 1000 Mbps</li>
	// <li>BANDWIDTH_POSTPAID_BY_HOUR：1 Mbps 至 100 Mbps</li>
	// <li>TRAFFIC_POSTPAID_BY_HOUR：1 Mbps 至 100 Mbps</li></ul>默认值：1 Mbps。</li>
	// <li>未开通带宽上移白名单的用户，EIP出带宽上限取决于与其绑定的实例的公网出带宽上限，无需传递此参数。</li></ul>

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 计费周期，预付费需要传递。

	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitempty" name:"AddressChargePrepaid"`
	// 订单ID，内部参数。

	DealId *string `json:"DealId,omitempty" name:"DealId"`
	// EIP类型。默认值：EIP。
	// <ul style="margin:0"><li>已开通Anycast公网加速白名单的用户，可选值：<ul><li>AnycastEIP：加速IP，可参见 [Anycast 公网加速](https://cloud.tencent.com/document/product/644)</li></ul>注意：仅部分地域支持加速IP。</li></ul>

	AddressType *string `json:"AddressType,omitempty" name:"AddressType"`
	// Anycast发布域。
	// <ul style="margin:0"><li>已开通Anycast公网加速白名单的用户，可选值：<ul><li>ANYCAST_ZONE_GLOBAL：全球发布域（需要额外开通Anycast全球加速白名单）</li><li>ANYCAST_ZONE_OVERSEAS：境外发布域</li></ul>默认值：ANYCAST_ZONE_OVERSEAS。</li></ul>

	AnycastZone *string `json:"AnycastZone,omitempty" name:"AnycastZone"`
	// EIP可用区，用于指定可用区申请EIP。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 指定IP地址申请EIP，每个账户每个月只有三次配额

	VipCluster []*string `json:"VipCluster,omitempty" name:"VipCluster"`
	// [Deprecated] 该字段无实际功能，废弃

	VipSet []*string `json:"VipSet,omitempty" name:"VipSet"`
	// AnycastEIP是否用于绑定负载均衡。
	// <ul style="margin:0"><li>已开通Anycast公网加速白名单的用户，可选值：<ul><li>TRUE：AnycastEIP可绑定对象为负载均衡</li>
	// <li>FALSE：AnycastEIP可绑定对象为云服务器、NAT网关、高可用虚拟IP等</li></ul>默认值：FALSE。</li></ul>

	ApplicableForCLB *bool `json:"ApplicableForCLB,omitempty" name:"ApplicableForCLB"`
	// 需要关联的标签列表。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// TGW集群ID，用于指定集群ID申请弹性公网IP。注意，anycast EIP不支持指定集群ID申请，静态单线EIP指定集群ID申请时，InternetServiceProvider参数必须和集群ID的运营商属性一致。

	SetId *int64 `json:"SetId,omitempty" name:"SetId"`
	// 是否是本地带宽IP。默认值：否

	Restrictive *bool `json:"Restrictive,omitempty" name:"Restrictive"`
}

func (r *AllocateAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNatGatewayRequest struct {
	*tchttp.BaseRequest

	// NAT网关的ID，形如：`nat-df45454`。

	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`
}

func (r *DeleteNatGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNatGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Firewall struct {
	// 防火墙唯一ID

	FirewallId *string `json:"FirewallId,omitempty" name:"FirewallId"`
	// 防火墙名称

	FirewallName *string `json:"FirewallName,omitempty" name:"FirewallName"`
	// 防火墙描述信息

	FirewallDescription *string `json:"FirewallDescription,omitempty" name:"FirewallDescription"`
	// 防火墙类型

	FirewallType *string `json:"FirewallType,omitempty" name:"FirewallType"`
	// 防火墙关联的Edge设备数量

	FirewallAttachedCount *int64 `json:"FirewallAttachedCount,omitempty" name:"FirewallAttachedCount"`
}

type CreateServiceTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 协议端口模板对象。

		ServiceTemplate *ServiceTemplate `json:"ServiceTemplate,omitempty" name:"ServiceTemplate"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateServiceTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcIpv6AddressesRequest struct {
	*tchttp.BaseRequest

	// `VPC`实例`ID`，形如：`vpc-f49l6u0z`。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// `IP`地址列表，批量查询单次请求最多支持`10`个。

	Ipv6Addresses []*string `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`
	// 偏移量。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeVpcIpv6AddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcIpv6AddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGatewayQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// NAT网关配额对象。

		NatGatewayQuota *NatGatewayQuota `json:"NatGatewayQuota,omitempty" name:"NatGatewayQuota"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatGatewayQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatGatewayQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceAllocateIp6AddressesBandwidthRequest struct {
	*tchttp.BaseRequest

	// 需要开通internet访问能力的IPV6地址

	Ip6Addresses []*string `json:"Ip6Addresses,omitempty" name:"Ip6Addresses"`
	// 带宽，单位Mbps。默认是1Mbps

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 网络计费模式。IPV6当前支持"TRAFFIC_POSTPAID_BY_HOUR"，默认是TRAFFIC_POSTPAID_BY_HOUR"。

	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
}

func (r *InquiryPriceAllocateIp6AddressesBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceAllocateIp6AddressesBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInterfacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例详细信息列表。

		NetworkInterfaceSet []*NetworkInterface `json:"NetworkInterfaceSet,omitempty" name:"NetworkInterfaceSet"`
		// 符合条件的实例数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkInterfacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkInterfacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcResourceDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源对象列表。

		ResourceDashboardSet []*ResourceDashboard `json:"ResourceDashboardSet,omitempty" name:"ResourceDashboardSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcResourceDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcResourceDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssignIpv6AddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分配给弹性网卡的`IPv6`地址列表。

		Ipv6AddressSet []*Ipv6Address `json:"Ipv6AddressSet,omitempty" name:"Ipv6AddressSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssignIpv6AddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssignIpv6AddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateIpAddressesAttributeRequest struct {
	*tchttp.BaseRequest

	// 弹性网卡实例ID，例如：eni-m6dyj72l。

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	// 指定的内网IP信息。

	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
}

func (r *ModifyPrivateIpAddressesAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPrivateIpAddressesAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest

	// 弹性网卡实例ID，例如：eni-m6dyj72l。

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	// CVM实例ID。形如：ins-r8hr2upy。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *AttachNetworkInterfaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachNetworkInterfaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublicIp6AddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 申请到的 IPV6 的唯一 ID 列表。

		AddressSet []*string `json:"AddressSet,omitempty" name:"AddressSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublicIp6AddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublicIp6AddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveIp6RulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveIp6RulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveIp6RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnassignPrivateIpAddressesRequest struct {
	*tchttp.BaseRequest

	// 弹性网卡实例ID，例如：eni-m6dyj72l。

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	// 指定的内网IP信息，单次最多指定10个。

	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
}

func (r *UnassignPrivateIpAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnassignPrivateIpAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 申请到的 EIP 的唯一 ID 列表。

		AddressSet []*string `json:"AddressSet,omitempty" name:"AddressSet"`
		// 异步任务TaskId。可以使用[DescribeTaskResult](https://cloud.tencent.com/document/api/215/36271)接口查询任务状态。

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcLimit struct {
	// 私有网络配额描述

	VpcLimitName *string `json:"VpcLimitName,omitempty" name:"VpcLimitName"`
	// 私有网络配额值

	VpcLimitValue *string `json:"VpcLimitValue,omitempty" name:"VpcLimitValue"`
}

type CreateAssistantCidrRequest struct {
	*tchttp.BaseRequest

	// `VPC`实例`ID`。形如：`vpc-6v2ht8q5`

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]

	CidrBlocks []*string `json:"CidrBlocks,omitempty" name:"CidrBlocks"`
}

func (r *CreateAssistantCidrRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssistantCidrRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTrafficPackagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTrafficPackagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTrafficPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableCcnRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableCcnRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceTemplate struct {
	// 协议端口实例ID，例如：ppm-f5n1f8da。

	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`
	// 模板名称。

	ServiceTemplateName *string `json:"ServiceTemplateName,omitempty" name:"ServiceTemplateName"`
	// 协议端口信息。

	ServiceSet []*string `json:"ServiceSet,omitempty" name:"ServiceSet"`
	// 创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type AllocateIp6AddressesBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 弹性公网 IPV6 的唯一 ID 列表。

		AddressSet []*string `json:"AddressSet,omitempty" name:"AddressSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateIp6AddressesBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateIp6AddressesBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAssistantCidrRequest struct {
	*tchttp.BaseRequest

	// `VPC`实例`ID`。形如：`vpc-6v2ht8q5`

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]

	CidrBlocks []*string `json:"CidrBlocks,omitempty" name:"CidrBlocks"`
}

func (r *DeleteAssistantCidrRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAssistantCidrRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatGatewayAddress struct {
	// 弹性公网IP（EIP）的唯一 ID，形如：`eip-11112222`。

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// 外网IP地址，形如：`123.121.34.33`。

	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`
	// 资源封堵状态。true表示弹性ip处于封堵状态，false表示弹性ip处于未封堵状态。

	IsBlocked *bool `json:"IsBlocked,omitempty" name:"IsBlocked"`
}

type CreateSecurityGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全组对象。

		SecurityGroup *SecurityGroup `json:"SecurityGroup,omitempty" name:"SecurityGroup"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustomerGatewayRequest struct {
	*tchttp.BaseRequest

	// 对端网关名称，可任意命名，但不得超过60个字符。

	CustomerGatewayName *string `json:"CustomerGatewayName,omitempty" name:"CustomerGatewayName"`
	// 对端网关公网IP。

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
}

func (r *CreateCustomerGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomerGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EipStatistics struct {
	// 地域英文标识

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域中文名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域EIP总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type QueryTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务执行结果

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckBandwidthPackageRequest struct {
	*tchttp.BaseRequest
}

func (r *CheckBandwidthPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckBandwidthPackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIp6TranslatorsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIp6TranslatorsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIp6TranslatorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceTemplateGroupRequest struct {
	*tchttp.BaseRequest

	// 协议端口模板集合实例ID，例如：ppmg-n17uxvve。

	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId,omitempty" name:"ServiceTemplateGroupId"`
}

func (r *DeleteServiceTemplateGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServiceTemplateGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 协议端口模板对象。

		ServiceTemplateSet []*ServiceTemplate `json:"ServiceTemplateSet,omitempty" name:"ServiceTemplateSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDirectConnectGatewayCcnRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirectConnectGatewayCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HaVipDisassociateAddressIpRequest struct {
	*tchttp.BaseRequest

	// `HAVIP`唯一`ID`，形如：`havip-9o233uri`。必须是已绑定`EIP`的`HAVIP`。

	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`
}

func (r *HaVipDisassociateAddressIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *HaVipDisassociateAddressIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyBandwidthPackageBandwidthRequest struct {
	*tchttp.BaseRequest

	// 带宽包ID

	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
	// 修改后的带宽

	InternetMaxBandwidth *uint64 `json:"InternetMaxBandwidth,omitempty" name:"InternetMaxBandwidth"`
}

func (r *InquiryPriceModifyBandwidthPackageBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceModifyBandwidthPackageBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIp6TranslatorRequest struct {
	*tchttp.BaseRequest

	// IPV6转换实例唯一ID，形如ip6-xxxxxxxxx

	Ip6TranslatorId *string `json:"Ip6TranslatorId,omitempty" name:"Ip6TranslatorId"`
	// IPV6转换实例修改名称

	Ip6TranslatorName *string `json:"Ip6TranslatorName,omitempty" name:"Ip6TranslatorName"`
}

func (r *ModifyIp6TranslatorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIp6TranslatorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcPrivateIpAddress struct {
	// `VPC`内网`IP`。

	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	// 所属子网`CIDR`。

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 内网`IP`类型。

	PrivateIpAddressType *string `json:"PrivateIpAddressType,omitempty" name:"PrivateIpAddressType"`
	// `IP`申请时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type AttachClassicLinkVpcRequest struct {
	*tchttp.BaseRequest

	// VPC实例ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// CVM实例ID

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *AttachClassicLinkVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachClassicLinkVpcRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressTemplateAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressTemplateAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAddressTemplateAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateNatGatewayAddressRequest struct {
	*tchttp.BaseRequest

	// NAT网关的ID，形如：`nat-df45454`。

	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`
	// 需要申请的弹性IP个数，系统会按您的要求生产N个弹性IP, 其中AddressCount和PublicAddresses至少传递一个。

	AddressCount *uint64 `json:"AddressCount,omitempty" name:"AddressCount"`
	// 绑定NAT网关的弹性IP数组，其中AddressCount和PublicAddresses至少传递一个。。

	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`
	// 弹性IP可以区，自动分配弹性IP时传递。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *AssociateNatGatewayAddressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateNatGatewayAddressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTrafficPackageAttributeRequest struct {
	*tchttp.BaseRequest

	// 共享流量包唯一标识ID

	TrafficPackageId *string `json:"TrafficPackageId,omitempty" name:"TrafficPackageId"`
	// 共享流量包名称

	TrafficPackageName *string `json:"TrafficPackageName,omitempty" name:"TrafficPackageName"`
}

func (r *ModifyTrafficPackageAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTrafficPackageAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcIpv6AddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// `IPv6`地址列表。

		Ipv6AddressSet []*VpcIpv6Address `json:"Ipv6AddressSet,omitempty" name:"Ipv6AddressSet"`
		// `IPv6`地址总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcIpv6AddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcIpv6AddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachClassicLinkVpcResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachClassicLinkVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachClassicLinkVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateAddressRequest struct {
	*tchttp.BaseRequest

	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：`eip-11112222`。

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// 表示解绑时是EIP从ENI上解绑，还是ENI+EIP从实例上解绑

	KeepAddressIdBindWithEniPip *bool `json:"KeepAddressIdBindWithEniPip,omitempty" name:"KeepAddressIdBindWithEniPip"`
	// 表示解绑 EIP 之后是否分配普通公网 IP。取值范围：<br><li>TRUE：表示解绑 EIP 之后分配普通公网 IP。<br><li>FALSE：表示解绑 EIP 之后不分配普通公网 IP。<br>默认取值：FALSE。<br><br>只有满足以下条件时才能指定该参数：<br><li> 只有在解绑主网卡的主内网 IP 上的 EIP 时才能指定该参数。<br><li>解绑 EIP 后重新分配普通公网 IP 操作一个账号每天最多操作 10 次；详情可通过 [DescribeAddressQuota](https://cloud.tencent.com/document/api/213/1378) 接口获取。

	ReallocateNormalPublicIp *bool `json:"ReallocateNormalPublicIp,omitempty" name:"ReallocateNormalPublicIp"`
	// [Deprecated] 表示解绑后是否自动释放EIP

	CascadeRelease *bool `json:"CascadeRelease,omitempty" name:"CascadeRelease"`
}

func (r *DisassociateAddressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateAddressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TransformAddressRequest struct {
	*tchttp.BaseRequest

	// 待操作有普通公网 IP 的实例 ID。实例 ID 形如：`ins-11112222`。可通过登录[控制台](https://console.cloud.tencent.com/cvm)查询，也可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/9389) 接口返回值中的`InstanceId`获取。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *TransformAddressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TransformAddressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupWithPoliciesRequest struct {
	*tchttp.BaseRequest

	// 项目id，默认0。可在qcloud控制台项目管理页面查询到。

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 安全组名称，可任意命名，但不得超过60个字符。

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 安全组备注，最多100个字符。

	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`
	// 安全组规则集合。

	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *CreateSecurityGroupWithPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupWithPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAttachCcnInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetAttachCcnInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetAttachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubnetRequest struct {
	*tchttp.BaseRequest

	// 待操作的VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网名称，最大长度不能超过60个字节。

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 子网网段，子网网段必须在VPC网段内，相同VPC内子网网段不能重叠。

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 子网所在的可用区ID，不同子网选择不同可用区可以做跨可用区灾备。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *CreateSubnetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubnetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateCcnBandwidthRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 云联网（CCN）地域带宽详情。

	CcnRegionBandwidthLimits []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimits,omitempty" name:"CcnRegionBandwidthLimits"`
}

func (r *InquiryPriceCreateCcnBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceCreateCcnBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressTemplateGroupAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressTemplateGroupAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAddressTemplateGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRouteTableRequest struct {
	*tchttp.BaseRequest

	// 待操作的VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 路由表名称，最大长度不能超过60个字节。

	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
}

func (r *CreateRouteTableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRouteTableRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetDetectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNetDetectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateVpnGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 商品价格。

		Price *Price `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateVpnGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceCreateVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAssistantCidrResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 辅助CIDR数组。

		AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitempty" name:"AssistantCidrSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAssistantCidrResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAssistantCidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConflictSource struct {
	// 冲突资源ID

	ConflictSourceId *string `json:"ConflictSourceId,omitempty" name:"ConflictSourceId"`
	// 冲突资源

	SourceItem *string `json:"SourceItem,omitempty" name:"SourceItem"`
	// 冲突资源条目信息

	ConflictItemSet []*ConflictItem `json:"ConflictItemSet,omitempty" name:"ConflictItemSet"`
}

type BandwidthLimitForCcnAlarmOnlyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网（CCN）各地域出带宽上限

		CcnRegionBandwidthLimitSet []*BandwidthLimitForCcnAlarmOnly `json:"CcnRegionBandwidthLimitSet,omitempty" name:"CcnRegionBandwidthLimitSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BandwidthLimitForCcnAlarmOnlyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BandwidthLimitForCcnAlarmOnlyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachClassicLinkVpcRequest struct {
	*tchttp.BaseRequest

	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// CVM实例ID查询。形如：ins-r8hr2upy。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DetachClassicLinkVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachClassicLinkVpcRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetDetectStatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的网络探测验证结果对象数组。

		NetDetectStateSet []*NetDetectState `json:"NetDetectStateSet,omitempty" name:"NetDetectStateSet"`
		// 符合条件的网络探测验证结果对象数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetDetectStatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDetectStatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyAddressesBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceModifyAddressesBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceModifyAddressesBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFlowLogAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFlowLogAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFlowLogAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceDirectConnectGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceDirectConnectGatewayCcnRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceDirectConnectGatewayCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateNatGatewayAddressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateNatGatewayAddressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateNatGatewayAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomerGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCustomerGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomerGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableRoutesRequest struct {
	*tchttp.BaseRequest

	// 路由表唯一ID。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 路由策略唯一ID。

	RouteIds []*uint64 `json:"RouteIds,omitempty" name:"RouteIds"`
}

func (r *EnableRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIp6AddressesBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIp6AddressesBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIp6AddressesBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ItemPrice struct {
	// 按量计费后付费单价，单位：元。

	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`
	// 按量计费后付费计价单元，可取值范围： HOUR：表示计价单元是按每小时来计算。当前涉及该计价单元的场景有：实例按小时后付费（POSTPAID_BY_HOUR）、带宽按小时后付费（BANDWIDTH_POSTPAID_BY_HOUR）： GB：表示计价单元是按每GB来计算。当前涉及该计价单元的场景有：流量按小时后付费（TRAFFIC_POSTPAID_BY_HOUR）。

	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`
	// 预付费商品的原价，单位：元。

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
	// 预付费商品的折扣价，单位：元。

	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
}

type AddBandwidthPackageResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddBandwidthPackageResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddBandwidthPackageResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteTableResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRouteTableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcPrivateIpAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 内网`IP`地址信息列表。

		VpcPrivateIpAddressSet []*VpcPrivateIpAddress `json:"VpcPrivateIpAddressSet,omitempty" name:"VpcPrivateIpAddressSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcPrivateIpAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcPrivateIpAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetTrafficMirrorSrcsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetTrafficMirrorSrcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetTrafficMirrorSrcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkInterface struct {
	// 弹性网卡实例ID，例如：eni-f1xjkw1b。

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	// 弹性网卡名称。

	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`
	// 弹性网卡描述。

	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitempty" name:"NetworkInterfaceDescription"`
	// 子网实例ID。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// VPC实例ID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 绑定的安全组。

	GroupSet []*string `json:"GroupSet,omitempty" name:"GroupSet"`
	// 是否是主网卡。

	Primary *bool `json:"Primary,omitempty" name:"Primary"`
	// MAC地址。

	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`
	// 弹性网卡状态：
	// <li>`PENDING`：创建中</li>
	// <li>`AVAILABLE`：可用的</li>
	// <li>`ATTACHING`：绑定中</li>
	// <li>`DETACHING`：解绑中</li>
	// <li>`DELETING`：删除中</li>

	State *string `json:"State,omitempty" name:"State"`
	// 内网IP信息。

	PrivateIpAddressSet []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet,omitempty" name:"PrivateIpAddressSet"`
	// 绑定的云服务器对象。

	Attachment *NetworkInterfaceAttachment `json:"Attachment,omitempty" name:"Attachment"`
	// 可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// `IPv6`地址列表。

	Ipv6AddressSet []*Ipv6Address `json:"Ipv6AddressSet,omitempty" name:"Ipv6AddressSet"`
	// 标签键值对。

	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`
	// 网卡类型。0 - 弹性网卡；1 - evm弹性网卡。

	EniType *uint64 `json:"EniType,omitempty" name:"EniType"`
}

type CreateServiceTemplateGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 协议端口模板集合对象。

		ServiceTemplateGroup *ServiceTemplateGroup `json:"ServiceTemplateGroup,omitempty" name:"ServiceTemplateGroup"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceTemplateGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateServiceTemplateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID，形如：ccn-gree226l。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// CCN路由策略唯一ID。形如：ccnr-f49l6u0z。

	RouteIds []*string `json:"RouteIds,omitempty" name:"RouteIds"`
	// 过滤条件，参数不支持同时指定RouteIds和Filters。
	// <li>route-id - String -（过滤条件）路由策略ID。</li>
	// <li>cidr-block - String -（过滤条件）目的端。</li>
	// <li>instance-type - String -（过滤条件）下一跳类型。</li>
	// <li>instance-region - String -（过滤条件）下一跳所属地域。</li>
	// <li>instance-id - String -（过滤条件）下一跳实例ID。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCcnRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrafficPackageStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrafficPackageStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrafficPackageStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomerGatewayAttributeRequest struct {
	*tchttp.BaseRequest

	// 对端网关ID，例如：cgw-2wqq41m9，可通过DescribeCustomerGateways接口查询对端网关。

	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`
	// 对端网关名称，可任意命名，但不得超过60个字符。

	CustomerGatewayName *string `json:"CustomerGatewayName,omitempty" name:"CustomerGatewayName"`
}

func (r *ModifyCustomerGatewayAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomerGatewayAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupPolicySet struct {
	// 安全组规则当前版本。用户每次更新安全规则版本会自动加1，防止更新的路由规则已过期，不填不考虑冲突。

	Version *string `json:"Version,omitempty" name:"Version"`
	// 出站规则。

	Egress []*SecurityGroupPolicy `json:"Egress,omitempty" name:"Egress"`
	// 入站规则。

	Ingress []*SecurityGroupPolicy `json:"Ingress,omitempty" name:"Ingress"`
}

type Route struct {
	// 路由策略ID。

	RouteId *uint64 `json:"RouteId,omitempty" name:"RouteId"`
	// 目的网段，取值不能在私有网络网段内，例如：112.20.51.0/24。

	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
	// 下一跳类型，目前我们支持的类型有：
	// CVM：公网网关类型的云主机；
	// VPN：VPN网关；
	// DIRECTCONNECT：专线网关；
	// PEERCONNECTION：对等连接；
	// SSLVPN：sslvpn网关；
	// NAT：NAT网关;
	// NORMAL_CVM：普通云主机；
	// EIP：云主机的公网IP；
	// CCN：云联网。

	GatewayType *string `json:"GatewayType,omitempty" name:"GatewayType"`
	// 下一跳地址，这里只需要指定不同下一跳类型的网关ID，系统会自动匹配到下一跳地址。
	// 特别注意：当 GatewayType 为 EIP 时，GatewayId 固定值 '0'

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 路由策略描述。

	RouteDescription *string `json:"RouteDescription,omitempty" name:"RouteDescription"`
	// 是否启用

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 路由类型，目前我们支持的类型有：
	// USER：用户路由；
	// NETD：网络探测路由，创建网络探测实例时，系统默认下发，不可编辑与删除；
	// CCN：云联网路由，系统默认下发，不可编辑与删除。
	// 用户只能添加和操作 USER 类型的路由。

	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`
}

type AddressLibInfo struct {
	// 国家

	Country *string `json:"Country,omitempty" name:"Country"`
	// 省

	Province *string `json:"Province,omitempty" name:"Province"`
	// 城市

	City *string `json:"City,omitempty" name:"City"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 运营商信息

	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`
	// IP地址

	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`
}

type DescribeAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的 EIP 数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// EIP 详细信息列表。

		AddressSet []*Address `json:"AddressSet,omitempty" name:"AddressSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupAssociationStatisticsRequest struct {
	*tchttp.BaseRequest

	// 安全实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

func (r *DescribeSecurityGroupAssociationStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupAssociationStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 执行结果，包括"SUCCESS", "FAILED", "RUNNING"

		Result *string `json:"Result,omitempty" name:"Result"`
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

type CreateTrafficPackagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTrafficPackagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTrafficPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCcnBandwidthRequest struct {
	*tchttp.BaseRequest

	// 是否自动续费标识；NOTIFY_AND_AUTO_RENEW：自动续费，NOTIFY_AND_MANUAL_RENEW：手动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 流量配置ID。

	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`
	// 地域间设置带宽，单位：Mbps。

	MaxBandwidthLimit *int64 `json:"MaxBandwidthLimit,omitempty" name:"MaxBandwidthLimit"`
}

func (r *UpdateCcnBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCcnBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSingleIspRegionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSingleIspRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSingleIspRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceVacancyAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceVacancyAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceVacancyAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDirectConnectGatewayCcnRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirectConnectGatewayCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPricePublicIp6AddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPricePublicIp6AddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPricePublicIp6AddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceSecurityGroupPolicyRequest struct {
	*tchttp.BaseRequest

	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 安全组规则集合对象。

	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *ReplaceSecurityGroupPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceSecurityGroupPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddIp6RulesRequest struct {
	*tchttp.BaseRequest

	// IPV6转换实例唯一ID，形如ip6-xxxxxxxx

	Ip6TranslatorId *string `json:"Ip6TranslatorId,omitempty" name:"Ip6TranslatorId"`
	// IPV6转换规则名称

	Ip6RuleName *string `json:"Ip6RuleName,omitempty" name:"Ip6RuleName"`
	// IPV6转换规则信息

	Ip6RuleInfos []*Ip6RuleInfo `json:"Ip6RuleInfos,omitempty" name:"Ip6RuleInfos"`
}

func (r *AddIp6RulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddIp6RulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回可用区信息数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 可用区详细信息

		ZoneSet []*ZoneInfo `json:"ZoneSet,omitempty" name:"ZoneSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateVpnGatewayRequest struct {
	*tchttp.BaseRequest

	// VPN网关计费模式，PREPAID：表示预付费，即包年包月，POSTPAID_BY_HOUR：表示后付费，即按量计费。默认：POSTPAID_BY_HOUR，如果指定预付费模式，参数InstanceChargePrepaid必填。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100；单位：Mbps。

	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

func (r *InquiryPriceCreateVpnGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceCreateVpnGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetNatGatewayConnectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetNatGatewayConnectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetNatGatewayConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAssistantCidrRequest struct {
	*tchttp.BaseRequest

	// `VPC`实例`ID`。形如：`vpc-6v2ht8q5`

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 待添加的负载CIDR。CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]

	NewCidrBlocks []*string `json:"NewCidrBlocks,omitempty" name:"NewCidrBlocks"`
	// 待删除的负载CIDR。CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]

	OldCidrBlocks []*string `json:"OldCidrBlocks,omitempty" name:"OldCidrBlocks"`
}

func (r *CheckAssistantCidrRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAssistantCidrRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnassignIpv6SubnetCidrBlockResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnassignIpv6SubnetCidrBlockResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnassignIpv6SubnetCidrBlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrafficPackageQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrafficPackageQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrafficPackageQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityGroupPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpnGatewayRequest struct {
	*tchttp.BaseRequest

	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// VPN网关名称，最大长度不能超过60个字节。

	VpnGatewayName *string `json:"VpnGatewayName,omitempty" name:"VpnGatewayName"`
	// VPN网关计费模式，PREPAID：表示预付费，即包年包月，POSTPAID_BY_HOUR：表示后付费，即按量计费。默认：POSTPAID_BY_HOUR，如果指定预付费模式，参数InstanceChargePrepaid必填。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100；单位：Mbps

	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 可用区，如：ap-guangzhou-2。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *CreateVpnGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpnGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetDetectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetDetectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BandwidthLimitForCcnAlarmOnly struct {
	// 地域，例如：ap-guangzhou。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 出带宽上限，单位：Mbps。

	BandwidthLimit *uint64 `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`
	// 整型ID。

	VbcId *uint64 `json:"VbcId,omitempty" name:"VbcId"`
	// 实例ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

type ModifyAddressesAttributeRequest struct {
	*tchttp.BaseRequest

	// 标识 EIP 的唯一 ID 列表。EIP 唯一 ID 形如：`eip-11112222`。

	AddressIds []*Address `json:"AddressIds,omitempty" name:"AddressIds"`
	// 1

	CascadeRelease *bool `json:"CascadeRelease,omitempty" name:"CascadeRelease"`
}

func (r *ModifyAddressesAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAddressesAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnassignIpv6CidrBlockRequest struct {
	*tchttp.BaseRequest

	// `VPC`实例`ID`，形如：`vpc-f49l6u0z`。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// `IPv6`网段。形如：`3402:4e00:20:1000::/56`

	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`
}

func (r *UnassignIpv6CidrBlockRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnassignIpv6CidrBlockRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRouteTableResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 路由表对象。

		RouteTable *RouteTable `json:"RouteTable,omitempty" name:"RouteTable"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRouteTableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupAssociationStatistics struct {
	// 安全组实例ID。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 云主机实例数。

	CVM *uint64 `json:"CVM,omitempty" name:"CVM"`
	// 数据库实例数。

	CDB *uint64 `json:"CDB,omitempty" name:"CDB"`
	// 弹性网卡实例数。

	ENI *uint64 `json:"ENI,omitempty" name:"ENI"`
	// 被安全组引用数。

	SG *uint64 `json:"SG,omitempty" name:"SG"`
	// 负载均衡实例数。

	CLB *uint64 `json:"CLB,omitempty" name:"CLB"`
	// 全量实例的绑定统计。

	InstanceStatistics []*InstanceStatistic `json:"InstanceStatistics,omitempty" name:"InstanceStatistics"`
}

type DirectConnectGateway struct {
	// 专线网关`ID`。

	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	// 专线网关名称。

	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitempty" name:"DirectConnectGatewayName"`
	// 专线网关关联`VPC`实例`ID`。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 关联网络类型：
	// <li>`VPC` - 私有网络</li>
	// <li>`CCN` - 云联网</li>

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// 关联网络实例`ID`：
	// <li>`NetworkType`为`VPC`时，这里为私有网络实例`ID`</li>
	// <li>`NetworkType`为`CCN`时，这里为云联网实例`ID`</li>

	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" name:"NetworkInstanceId"`
	// 网关类型：
	// <li>NORMAL - 标准型，注：云联网只支持标准型</li>
	// <li>NAT - NAT型</li>
	// NAT类型支持网络地址转换配置，类型确定后不能修改；一个私有网络可以创建一个NAT类型的专线网关和一个非NAT类型的专线网关

	GatewayType *string `json:"GatewayType,omitempty" name:"GatewayType"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 专线网关IP。

	DirectConnectGatewayIp *string `json:"DirectConnectGatewayIp,omitempty" name:"DirectConnectGatewayIp"`
	// 专线网关关联`CCN`实例`ID`。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 云联网路由学习类型：
	// <li>`BGP` - 自动学习。</li>
	// <li>`STATIC` - 静态，即用户配置。</li>

	CcnRouteType *string `json:"CcnRouteType,omitempty" name:"CcnRouteType"`
	// 是否启用BGP。

	EnableBGP *bool `json:"EnableBGP,omitempty" name:"EnableBGP"`
	// 开启和关闭BGP的community属性。

	EnableBGPCommunity *bool `json:"EnableBGPCommunity,omitempty" name:"EnableBGPCommunity"`
}

type PrivateIpAddressSpecification struct {
	// 内网IP地址。

	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	// 是否是主IP。

	Primary *bool `json:"Primary,omitempty" name:"Primary"`
	// 公网IP地址。

	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`
	// EIP实例ID，例如：eip-11112222。

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// 内网IP描述信息。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 公网IP是否被封堵。

	IsWanIpBlocked *bool `json:"IsWanIpBlocked,omitempty" name:"IsWanIpBlocked"`
	// IP状态：
	// PENDING：生产中
	// MIGRATING：迁移中
	// DELETING：删除中
	// AVAILABLE：可用的

	State *string `json:"State,omitempty" name:"State"`
}

type InquiryPricePublicIp6AddressesRequest struct {
	*tchttp.BaseRequest

	// 需要开通internet访问能力的IPV6地址

	Ip6Addresses []*string `json:"Ip6Addresses,omitempty" name:"Ip6Addresses"`
	// 带宽，单位Mbps。默认是1Mbps

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 网络计费模式。IPV6当前支持"TRAFFIC_POSTPAID_BY_HOUR"，默认是TRAFFIC_POSTPAID_BY_HOUR"。

	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
}

func (r *InquiryPricePublicIp6AddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPricePublicIp6AddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewAddressesRequest struct {
	*tchttp.BaseRequest

	// 1

	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
	// 1

	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitempty" name:"AddressChargePrepaid"`
	// 1

	DealId *string `json:"DealId,omitempty" name:"DealId"`
	// 1

	CurrentDeadline *string `json:"CurrentDeadline,omitempty" name:"CurrentDeadline"`
}

func (r *InquiryPriceRenewAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAttachCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// CCN所属UIN（根账号）。

	CcnUin *string `json:"CcnUin,omitempty" name:"CcnUin"`
	// 重新申请关联网络实例列表。

	Instances []*CcnInstance `json:"Instances,omitempty" name:"Instances"`
}

func (r *ResetAttachCcnInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetAttachCcnInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddressBase struct {
	// 运营商信息，包括移动，联通，电信，腾讯CAP

	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`
	// 弹性公网IP集群名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 弹性公网IP集群唯一Id

	SetId *int64 `json:"SetId,omitempty" name:"SetId"`
	// 弹性公网IP地址

	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`
	// 弹性公网IP是否支持DPDK

	IsDpdk *bool `json:"IsDpdk,omitempty" name:"IsDpdk"`
	// 弹性公网IP是否可用

	Status *string `json:"Status,omitempty" name:"Status"`
}

type InquiryPriceAllocateIp6AddressesBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceAllocateIp6AddressesBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceAllocateIp6AddressesBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest struct {
	*tchttp.BaseRequest

	// NAT网关ID。

	NatGatewayIds []*string `json:"NatGatewayIds,omitempty" name:"NatGatewayIds"`
	// 过滤条件:
	// 参数不支持同时指定NatGatewayIds和Filters。
	// <li> nat-gateway-id，NAT网关的ID，如`nat-0yi4hekt`</li>
	// <li> vpc-id，私有网络VPC的ID，如`vpc-0yi4hekt`</li>
	// <li> public-ip-address， 弹性IP，如`139.199.232.238`。</li>
	// <li>public-port， 公网端口。</li>
	// <li>private-ip-address， 内网IP，如`10.0.0.1`。</li>
	// <li>private-port， 内网端口。</li>
	// <li>description，规则描述。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest

	// 弹性网卡实例ID，例如：eni-m6dyj72l。

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	// CVM实例ID。形如：ins-r8hr2upy。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DetachNetworkInterfaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachNetworkInterfaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomerGateway struct {
	// 用户网关唯一ID

	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`
	// 网关名称

	CustomerGatewayName *string `json:"CustomerGatewayName,omitempty" name:"CustomerGatewayName"`
	// 公网地址

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type GatewayFlowMonitorDetail struct {
	// 来源`IP`。

	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	// 入包量。

	InPkg *uint64 `json:"InPkg,omitempty" name:"InPkg"`
	// 出包量。

	OutPkg *uint64 `json:"OutPkg,omitempty" name:"OutPkg"`
	// 入带宽，单位：`Byte`。

	InTraffic *uint64 `json:"InTraffic,omitempty" name:"InTraffic"`
	// 出带宽，单位：`Byte`。

	OutTraffic *uint64 `json:"OutTraffic,omitempty" name:"OutTraffic"`
}

type DescribeBandwidthPackageResourcesRequest struct {
	*tchttp.BaseRequest

	// 标识 共享带宽包 的唯一 ID 列表。共享带宽包 唯一 ID 形如：`bwp-11112222`。

	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`AddressIds`和`Filters`。详细的过滤条件如下：
	// <li> resource-id - String - 是否必填：否 - （过滤条件）按照 共享带宽包内资源 的唯一 ID 过滤。共享带宽包内资源 唯一 ID 形如：eip-11112222。</li>
	// <li> resource-type - String - 是否必填：否 - （过滤条件）按照 共享带宽包内资源 类型过滤，目前仅支持 弹性IP 和 负载均衡 两种类型，可选值为 Address 和 LoadBalance。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBandwidthPackageResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBandwidthPackageResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCreateCcnBandwidthDealRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 云联网（CCN）各地域带宽上限。

	CcnRegionBandwidthLimits []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimits,omitempty" name:"CcnRegionBandwidthLimits"`
}

func (r *GetCreateCcnBandwidthDealRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCreateCcnBandwidthDealRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceRouteTableAssociationRequest struct {
	*tchttp.BaseRequest

	// 子网实例ID，例如：subnet-3x5lf5q0。可通过DescribeSubnets接口查询。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 路由表实例ID，例如：rtb-azd4dt1c。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

func (r *ReplaceRouteTableAssociationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceRouteTableAssociationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopTrafficMirrorRequest struct {
	*tchttp.BaseRequest

	// 流量镜像实例ID

	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`
}

func (r *StopTrafficMirrorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopTrafficMirrorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceDashboard struct {
	// Vpc实例ID，例如：vpc-f1xjkw1b。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网实例ID，例如：subnet-bthucmmy。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 基础网络互通。

	Classiclink *uint64 `json:"Classiclink,omitempty" name:"Classiclink"`
	// 专线网关。

	Dcg *uint64 `json:"Dcg,omitempty" name:"Dcg"`
	// 对等连接。

	Pcx *uint64 `json:"Pcx,omitempty" name:"Pcx"`
	// 当前已使用的IP总数。

	Ip *uint64 `json:"Ip,omitempty" name:"Ip"`
	// NAT网关。

	Nat *uint64 `json:"Nat,omitempty" name:"Nat"`
	// VPN网关。

	Vpngw *uint64 `json:"Vpngw,omitempty" name:"Vpngw"`
	// 流日志。

	FlowLog *uint64 `json:"FlowLog,omitempty" name:"FlowLog"`
	// 网络探测。

	NetworkDetect *uint64 `json:"NetworkDetect,omitempty" name:"NetworkDetect"`
	// 网络ACL。

	NetworkACL *uint64 `json:"NetworkACL,omitempty" name:"NetworkACL"`
	// 云主机。

	CVM *uint64 `json:"CVM,omitempty" name:"CVM"`
	// 负载均衡。

	LB *uint64 `json:"LB,omitempty" name:"LB"`
	// 关系型数据库。

	CDB *uint64 `json:"CDB,omitempty" name:"CDB"`
	// 云数据库 TencentDB for Memcached。

	Cmem *uint64 `json:"Cmem,omitempty" name:"Cmem"`
	// 时序数据库。

	CTSDB *uint64 `json:"CTSDB,omitempty" name:"CTSDB"`
	// 数据库 TencentDB for MariaDB（TDSQL）。

	MariaDB *uint64 `json:"MariaDB,omitempty" name:"MariaDB"`
	// 数据库 TencentDB for SQL Server。

	SQLServer *uint64 `json:"SQLServer,omitempty" name:"SQLServer"`
	// 云数据库 TencentDB for PostgreSQL。

	Postgres *uint64 `json:"Postgres,omitempty" name:"Postgres"`
	// 网络附加存储。

	NAS *uint64 `json:"NAS,omitempty" name:"NAS"`
	// Snova云数据仓库。

	Greenplumn *uint64 `json:"Greenplumn,omitempty" name:"Greenplumn"`
	// 消息队列 CKAFKA。

	Ckafka *uint64 `json:"Ckafka,omitempty" name:"Ckafka"`
	// Grocery。

	Grocery *uint64 `json:"Grocery,omitempty" name:"Grocery"`
	// 数据加密服务。

	HSM *uint64 `json:"HSM,omitempty" name:"HSM"`
	// 游戏存储 Tcaplus。

	Tcaplus *uint64 `json:"Tcaplus,omitempty" name:"Tcaplus"`
	// Cnas。

	Cnas *uint64 `json:"Cnas,omitempty" name:"Cnas"`
	// HTAP 数据库 TiDB。

	TiDB *uint64 `json:"TiDB,omitempty" name:"TiDB"`
	// EMR 集群。

	Emr *uint64 `json:"Emr,omitempty" name:"Emr"`
	// SEAL。

	SEAL *uint64 `json:"SEAL,omitempty" name:"SEAL"`
	// 文件存储 CFS。

	CFS *uint64 `json:"CFS,omitempty" name:"CFS"`
	// Oracle。

	Oracle *uint64 `json:"Oracle,omitempty" name:"Oracle"`
	// ElasticSearch服务。

	ElasticSearch *uint64 `json:"ElasticSearch,omitempty" name:"ElasticSearch"`
	// 区块链服务。

	TBaaS *uint64 `json:"TBaaS,omitempty" name:"TBaaS"`
	// Itop。

	Itop *uint64 `json:"Itop,omitempty" name:"Itop"`
	// 云数据库审计。

	DBAudit *uint64 `json:"DBAudit,omitempty" name:"DBAudit"`
	// 企业级云数据库 CynosDB for Postgres。

	CynosDBPostgres *uint64 `json:"CynosDBPostgres,omitempty" name:"CynosDBPostgres"`
	// 数据库 TencentDB for Redis。

	Redis *uint64 `json:"Redis,omitempty" name:"Redis"`
	// 数据库 TencentDB for MongoDB。

	MongoDB *uint64 `json:"MongoDB,omitempty" name:"MongoDB"`
	// 分布式数据库 TencentDB for TDSQL。

	DCDB *uint64 `json:"DCDB,omitempty" name:"DCDB"`
	// 企业级云数据库 CynosDB for MySQL。

	CynosDBMySQL *uint64 `json:"CynosDBMySQL,omitempty" name:"CynosDBMySQL"`
}

type NatGatewayDestinationIpPortTranslationNatRule struct {
	// 网络协议，可选值：`TCP`、`UDP`。

	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`
	// 弹性IP。

	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`
	// 公网端口。

	PublicPort *uint64 `json:"PublicPort,omitempty" name:"PublicPort"`
	// 内网地址。

	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	// 内网端口。

	PrivatePort *uint64 `json:"PrivatePort,omitempty" name:"PrivatePort"`
	// NAT网关转发规则描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// NAT网关的ID。

	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`
	// 私有网络VPC的ID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// NAT网关转发规则创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFlowLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFlowLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressAvailabilityRequest struct {
	*tchttp.BaseRequest

	// 申请 EIP 数量，默认值为1。

	AddressCount *uint64 `json:"AddressCount,omitempty" name:"AddressCount"`
	// EIP类型，EIP|AnycastEIP，默认EIP。

	AddressType *string `json:"AddressType,omitempty" name:"AddressType"`
	// 申请适用于CLB的EIP，默认值为false。

	ApplicableForCLB *bool `json:"ApplicableForCLB,omitempty" name:"ApplicableForCLB"`
	// 运营商名称，可选值[BGP|CTCC|CMCC]，默认BGP。

	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`
	// EIP可用区，用于指定可用区申请EIP。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 购买渠道。

	LaunchSource *bool `json:"LaunchSource,omitempty" name:"LaunchSource"`
	// 全部可选参数组合。

	AllPossibleCombinations *string `json:"AllPossibleCombinations,omitempty" name:"AllPossibleCombinations"`
}

func (r *DescribeAddressAvailabilityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressAvailabilityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssistantCidrRequest struct {
	*tchttp.BaseRequest

	// `VPC`实例`ID`数组。形如：[`vpc-6v2ht8q5`]

	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`
	// 过滤条件，参数不支持同时指定NetworkInterfaceIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAssistantCidrRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssistantCidrRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnsRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定CcnIds和Filters。

	CcnIds []*string `json:"CcnIds,omitempty" name:"CcnIds"`
	// 过滤条件，参数不支持同时指定CcnIds和Filters。
	// <li>ccn-id - String - （过滤条件）CCN唯一ID，形如：vpc-f49l6u0z。</li>
	// <li>ccn-name - String - （过滤条件）CCN名称。</li>
	// <li>ccn-description - String - （过滤条件）CCN描述。</li>
	// <li>state - String - （过滤条件）实例状态， 'ISOLATED': 隔离中（欠费停服），'AVAILABLE'：运行中。</li>
	// <li>tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例：查询绑定了标签的CCN列表。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段。支持：`CcnId` `CcnName` `CreateTime` `State` `QosLevel`

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方法。顺序：`ASC`，倒序：`DESC`。

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeCcnsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIpOnlineRequest struct {
	*tchttp.BaseRequest

	// IP地址列表，支持IPv4和IPv6

	AddressIps []*string `json:"AddressIps,omitempty" name:"AddressIps"`
}

func (r *DescribeIpOnlineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIpOnlineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyAddressesBandwidthRequest struct {
	*tchttp.BaseRequest

	// EIP唯一ID

	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
	// 新带宽值

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 订单ID

	DealId *string `json:"DealId,omitempty" name:"DealId"`
	// 预约贷款表记录ID

	BandwidthId *int64 `json:"BandwidthId,omitempty" name:"BandwidthId"`
}

func (r *InquiryPriceModifyAddressesBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceModifyAddressesBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFlowLogAttributeRequest struct {
	*tchttp.BaseRequest

	// 私用网络ID或者统一ID，建议使用统一ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 流日志唯一ID

	FlowLogId *string `json:"FlowLogId,omitempty" name:"FlowLogId"`
	// 流日志实例名字

	FlowLogName *string `json:"FlowLogName,omitempty" name:"FlowLogName"`
	// 流日志实例描述

	FlowLogDescription *string `json:"FlowLogDescription,omitempty" name:"FlowLogDescription"`
}

func (r *ModifyFlowLogAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFlowLogAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHaVipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// `HAVIP`对象。

		HaVip *HaVip `json:"HaVip,omitempty" name:"HaVip"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHaVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHaVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlowLog struct {
	// 私用网络ID或者统一ID，建议使用统一ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 流日志唯一ID

	FlowLogId *string `json:"FlowLogId,omitempty" name:"FlowLogId"`
	// 流日志实例名字

	FlowLogName *string `json:"FlowLogName,omitempty" name:"FlowLogName"`
	// 流日志所属资源类型，VPC|SUBNET|NETWORKINTERFACE

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源唯一ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 流日志采集类型，ACCEPT|REJECT|ALL

	TrafficType *string `json:"TrafficType,omitempty" name:"TrafficType"`
	// 流日志存储ID

	CloudLogId *string `json:"CloudLogId,omitempty" name:"CloudLogId"`
	// 流日志存储ID状态

	CloudLogState *string `json:"CloudLogState,omitempty" name:"CloudLogState"`
	// 流日志描述信息

	FlowLogDescription *string `json:"FlowLogDescription,omitempty" name:"FlowLogDescription"`
	// 流日志创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type DescribeTrafficMirrorsRequest struct {
	*tchttp.BaseRequest

	// 流量镜像实例ID集合

	TrafficMirrorIds []*string `json:"TrafficMirrorIds,omitempty" name:"TrafficMirrorIds"`
	// 流量镜像查询过滤调节

	Filters *Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTrafficMirrorsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrafficMirrorsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetDetectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络探测（NetDetect）对象。

		NetDetect *NetDetect `json:"NetDetect,omitempty" name:"NetDetect"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetDetectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectGatewaysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 专线网关对象数组。

		DirectConnectGatewaySet []*DirectConnectGateway `json:"DirectConnectGatewaySet,omitempty" name:"DirectConnectGatewaySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectGatewaysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnGatewaysRequest struct {
	*tchttp.BaseRequest

	// VPN网关实例ID。形如：vpngw-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpnGatewayIds和Filters。

	VpnGatewayIds []*string `json:"VpnGatewayIds,omitempty" name:"VpnGatewayIds"`
	// 过滤条件，参数不支持同时指定VpnGatewayIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID形如：vpc-f49l6u0z。</li>
	// <li>vpn-gateway-id - String - （过滤条件）VPN实例ID形如：vpngw-5aluhh9t。</li>
	// <li>vpn-gateway-name - String - （过滤条件）VPN实例名称。</li>
	// <li>type - String - （过滤条件）VPN网关类型：'IPSEC', 'SSL'。</li>
	// <li>public-ip-address- String - （过滤条件）公网IP。</li>
	// <li>renew-flag - String - （过滤条件）网关续费类型，手动续费：'NOTIFY_AND_MANUAL_RENEW'、自动续费：'NOTIFY_AND_AUTO_RENEW'。</li>
	// <li>zone - String - （过滤条件）VPN所在可用区，形如：ap-guangzhou-2。</li>

	Filters []*FilterObject `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 请求对象个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeVpnGatewaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnGatewaysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IKEOptionsSpecification struct {
	// 加密算法，可选值：'3DES-CBC', 'AES-CBC-128', 'AES-CBS-192', 'AES-CBC-256', 'DES-CBC'，默认为3DES-CBC

	PropoEncryAlgorithm *string `json:"PropoEncryAlgorithm,omitempty" name:"PropoEncryAlgorithm"`
	// 认证算法：可选值：'MD5', 'SHA1'，默认为MD5

	PropoAuthenAlgorithm *string `json:"PropoAuthenAlgorithm,omitempty" name:"PropoAuthenAlgorithm"`
	// 协商模式：可选值：'AGGRESSIVE', 'MAIN'，默认为MAIN

	ExchangeMode *string `json:"ExchangeMode,omitempty" name:"ExchangeMode"`
	// 本端标识类型：可选值：'ADDRESS', 'FQDN'，默认为ADDRESS

	LocalIdentity *string `json:"LocalIdentity,omitempty" name:"LocalIdentity"`
	// 对端标识类型：可选值：'ADDRESS', 'FQDN'，默认为ADDRESS

	RemoteIdentity *string `json:"RemoteIdentity,omitempty" name:"RemoteIdentity"`
	// 本端标识，当LocalIdentity选为ADDRESS时，LocalAddress必填。localAddress默认为vpn网关公网IP

	LocalAddress *string `json:"LocalAddress,omitempty" name:"LocalAddress"`
	// 对端标识，当RemoteIdentity选为ADDRESS时，RemoteAddress必填

	RemoteAddress *string `json:"RemoteAddress,omitempty" name:"RemoteAddress"`
	// 本端标识，当LocalIdentity选为FQDN时，LocalFqdnName必填

	LocalFqdnName *string `json:"LocalFqdnName,omitempty" name:"LocalFqdnName"`
	// 对端标识，当remoteIdentity选为FQDN时，RemoteFqdnName必填

	RemoteFqdnName *string `json:"RemoteFqdnName,omitempty" name:"RemoteFqdnName"`
	// DH group，指定IKE交换密钥时使用的DH组，可选值：'GROUP1', 'GROUP2', 'GROUP5', 'GROUP14', 'GROUP24'，

	DhGroupName *string `json:"DhGroupName,omitempty" name:"DhGroupName"`
	// IKE SA Lifetime，单位：秒，设置IKE SA的生存周期，取值范围：60-604800

	IKESaLifetimeSeconds *uint64 `json:"IKESaLifetimeSeconds,omitempty" name:"IKESaLifetimeSeconds"`
	// IKE版本

	IKEVersion *string `json:"IKEVersion,omitempty" name:"IKEVersion"`
}

type AssignIpv6AddressesRequest struct {
	*tchttp.BaseRequest

	// 弹性网卡实例`ID`，形如：`eni-m6dyj72l`。

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	// 指定的`IPv6`地址列表，单次最多指定10个。与入参`Ipv6AddressCount`合并计算配额。

	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`
	// 自动分配`IPv6`地址个数，内网IP地址个数总和不能超过配数。与入参`Ipv6Addresses`合并计算配额。

	Ipv6AddressCount *uint64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`
}

func (r *AssignIpv6AddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssignIpv6AddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBandwidthPackagesRequest struct {
	*tchttp.BaseRequest

	// 带宽包Id，支持批量

	BandwidthPackageIds []*string `json:"BandwidthPackageIds,omitempty" name:"BandwidthPackageIds"`
	// 每次请求的`Filters`的上限为10。参数不支持同时指定`BandwidthPackageIds`和`Filters`。详细的过滤条件如下：
	// <li> bandwidth-package_id - String - 是否必填：否 - （过滤条件）按照带宽包的唯一标识ID过滤。</li>
	// <li> bandwidth-package-name - String - 是否必填：否 - （过滤条件）按照 带宽包名称过滤。不支持模糊过滤。</li>
	// <li> network-type - String - 是否必填：否 - （过滤条件）按照带宽包的类型过滤。类型包括'BGP','SINGLEISP'和'ANYCAST'。</li>
	// <li> charge-type - String - 是否必填：否 - （过滤条件）按照带宽包的计费类型过滤。计费类型包括'TOP5_POSTPAID_BY_MONTH'和'PERCENT95_POSTPAID_BY_MONTH'</li>
	// <li> resource.resource-type - String - 是否必填：否 - （过滤条件）按照带宽包资源类型过滤。资源类型包括'Address'和'LoadBalance'</li>
	// <li> resource.resource-id - String - 是否必填：否 - （过滤条件）按照带宽包资源Id过滤。资源Id形如'eip-xxxx','lb-xxxx'</li>
	// <li> resource.address-ip - String - 是否必填：否 - （过滤条件）按照带宽包资源Ip过滤。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 查询带宽包偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询带宽包数量限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 是否展示带宽包内部资源，默认展示。True表示展示带宽包内详细资源，False表示只展示带宽包而不展示内部详细资源。

	QueryResources *bool `json:"QueryResources,omitempty" name:"QueryResources"`
}

func (r *DescribeBandwidthPackagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBandwidthPackagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateIpAddressesAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPrivateIpAddressesAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPrivateIpAddressesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceTemplateGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceTemplateGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServiceTemplateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHaVipsResponse struct {
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

func (r *DescribeHaVipsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHaVipsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressesAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressesAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAddressesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReturnNormalAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReturnNormalAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReturnNormalAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnassignIpv6AddressesRequest struct {
	*tchttp.BaseRequest

	// 弹性网卡实例`ID`，形如：`eni-m6dyj72l`。

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	// 指定的`IPv6`地址列表，单次最多指定10个。

	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`
}

func (r *UnassignIpv6AddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnassignIpv6AddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkInterfaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 弹性网卡实例。

		NetworkInterface *NetworkInterface `json:"NetworkInterface,omitempty" name:"NetworkInterface"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetworkInterfaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressAssociationQuotaRequest struct {
	*tchttp.BaseRequest

	// [已废弃]待绑定的云主机实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 需要查询的云主机实例ID列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// [已废弃]待绑定的弹性网卡ID

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	// 需要查询的弹性网卡ID列表

	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" name:"NetworkInterfaceIds"`
	// 待绑定的云主机实例ID

	TargetInstanceId *string `json:"TargetInstanceId,omitempty" name:"TargetInstanceId"`
	// 待绑定的弹性网卡ID

	TargetNetworkInterfaceId *string `json:"TargetNetworkInterfaceId,omitempty" name:"TargetNetworkInterfaceId"`
	// 指定CPU核数查询云主机绑定配额

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 指定内存大小查询云主机绑定配额

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
}

func (r *DescribeAddressAssociationQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressAssociationQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流日志实例集合

		FlowLog []*FlowLog `json:"FlowLog,omitempty" name:"FlowLog"`
		// 流日志总数目

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetTrafficMirrorTargetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetTrafficMirrorTargetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetTrafficMirrorTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTrafficMirrorRequest struct {
	*tchttp.BaseRequest

	// 流量镜像实例ID

	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`
}

func (r *DeleteTrafficMirrorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTrafficMirrorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSingleIspRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 三网运营商和可用区的映射关系

		SingleIspZone []*SingleIspZone `json:"SingleIspZone,omitempty" name:"SingleIspZone"`
		// 支持三网运营商的可用区数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSingleIspRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSingleIspRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIp6RuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIp6RuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIp6RuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpnGatewayQuota struct {
	// 带宽配额

	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 配额中文名称

	Cname *string `json:"Cname,omitempty" name:"Cname"`
	// 配额英文名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeNetDetectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的网络探测对象数组。

		NetDetectSet []*NetDetect `json:"NetDetectSet,omitempty" name:"NetDetectSet"`
		// 符合条件的网络探测对象数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetDetectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetDetectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetTrafficMirrorSrcsRequest struct {
	*tchttp.BaseRequest

	// 流量镜像实例ID

	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`
	// 流量镜像采集对象

	CollectorSrcs []*string `json:"CollectorSrcs,omitempty" name:"CollectorSrcs"`
}

func (r *ResetTrafficMirrorSrcsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetTrafficMirrorSrcsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubnetInput struct {
	// 子网的`CIDR`。

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 子网名称。

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 可用区。形如：`ap-guangzhou-2`。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 指定关联路由表，形如：`rtb-3ryrwzuu`。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

type DescribeAddressQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户 EIP 配额信息。

		QuotaSet []*Quota `json:"QuotaSet,omitempty" name:"QuotaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountAttributesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccountAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatGatewayAttributeRequest struct {
	*tchttp.BaseRequest

	// NAT网关的ID，形如：`nat-df45454`。

	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`
	// NAT网关的名称，形如：`test_nat`。

	NatGatewayName *string `json:"NatGatewayName,omitempty" name:"NatGatewayName"`
	// NAT网关最大外网出带宽(单位:Mbps)。

	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

func (r *ModifyNatGatewayAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatGatewayAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnGatewayAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpnGatewayAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpnGatewayAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssignIpv6CidrBlockResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分配的 `IPv6` 网段。形如：`3402:4e00:20:1000::/56`

		Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssignIpv6CidrBlockResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssignIpv6CidrBlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCcnBandwidthRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 流量配置ID。

	RegionFlowControlIds []*string `json:"RegionFlowControlIds,omitempty" name:"RegionFlowControlIds"`
	// 是否自动续费标识，1：自动续费；2：不自动续费

	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// 是否自动续费标识，NOTIFY_AND_AUTO_RENEW：自动续费；NOTIFY_AND_MANUAL_RENEW： 不自动续费

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *SetCcnBandwidthRenewFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCcnBandwidthRenewFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceStatistic struct {
	// 实例的类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例的个数

	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
}

type CreateAddressTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IP地址模板对象。

		AddressTemplate *AddressTemplate `json:"AddressTemplate,omitempty" name:"AddressTemplate"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAddressTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAddressTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// CCN路由策略唯一ID。形如：ccnr-f49l6u0z。

	RouteIds []*string `json:"RouteIds,omitempty" name:"RouteIds"`
}

func (r *EnableCcnRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableCcnRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateAddressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务TaskId。可以使用[DescribeTaskResult](https://cloud.tencent.com/document/api/215/36271)接口查询任务状态。

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateAddressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpnConnectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 通道实例对象。

		VpnConnection *VpnConnection `json:"VpnConnection,omitempty" name:"VpnConnection"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpnConnectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpnConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnAttachedInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 关联实例列表。

		InstanceSet []*CcnAttachedInstance `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnAttachedInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnAttachedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecurityGroupPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceTemplateGroupAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceTemplateGroupAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServiceTemplateGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachNetworkInterfaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachNetworkInterfaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssignIpv6SubnetCidrBlockResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分配 `IPv6` 子网段列表。

		Ipv6SubnetCidrBlockSet []*Ipv6SubnetCidrBlock `json:"Ipv6SubnetCidrBlockSet,omitempty" name:"Ipv6SubnetCidrBlockSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssignIpv6SubnetCidrBlockResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssignIpv6SubnetCidrBlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubnetsRequest struct {
	*tchttp.BaseRequest

	// `VPC`实例`ID`。形如：`vpc-6v2ht8q5`

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网对象列表。

	Subnets []*SubnetInput `json:"Subnets,omitempty" name:"Subnets"`
}

func (r *CreateSubnetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubnetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerGatewaysRequest struct {
	*tchttp.BaseRequest

	// 对端网关ID，例如：cgw-2wqq41m9。每次请求的实例的上限为100。参数不支持同时指定CustomerGatewayIds和Filters。

	CustomerGatewayIds []*string `json:"CustomerGatewayIds,omitempty" name:"CustomerGatewayIds"`
	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定CustomerGatewayIds和Filters。
	// <li>customer-gateway-id - String - （过滤条件）用户网关唯一ID形如：`cgw-mgp33pll`。</li>
	// <li>customer-gateway-name - String - （过滤条件）用户网关名称形如：`test-cgw`。</li>
	// <li>ip-address - String - （过滤条件）公网地址形如：`58.211.1.12`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCustomerGatewaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerGatewaysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddressTemplateItem struct {
	// 起始地址。

	From *string `json:"From,omitempty" name:"From"`
	// 结束地址。

	To *string `json:"To,omitempty" name:"To"`
}

type CreateFlowLogRequest struct {
	*tchttp.BaseRequest

	// 私用网络ID或者统一ID，建议使用统一ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 流日志实例名字

	FlowLogName *string `json:"FlowLogName,omitempty" name:"FlowLogName"`
	// 流日志实例描述

	FlowLogDescription *string `json:"FlowLogDescription,omitempty" name:"FlowLogDescription"`
	// 流日志所属资源类型，VPC|SUBNET|NETWORKINTERFACE

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源唯一ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 流日志采集类型，ACCEPT|REJECT|ALL

	TrafficType *string `json:"TrafficType,omitempty" name:"TrafficType"`
	// 流日志存储ID

	CloudLogId *string `json:"CloudLogId,omitempty" name:"CloudLogId"`
}

func (r *CreateFlowLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFlowLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBandwidthPackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBandwidthPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBandwidthPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网配额列表

		CcnLimitSet []*CcnLimit `json:"CcnLimitSet,omitempty" name:"CcnLimitSet"`
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTrafficMirrorAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTrafficMirrorAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTrafficMirrorAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DirectConnectGatewayCcnRoute struct {
	// 路由ID。

	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`
	// IDC网段。

	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
	// `BGP`的`AS-Path`属性。

	ASPath []*string `json:"ASPath,omitempty" name:"ASPath"`
}

type InquiryPriceRenewBandwidthPackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewBandwidthPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewBandwidthPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTrafficPackageAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTrafficPackageAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTrafficPackageAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateNatGatewayAddressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateNatGatewayAddressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateNatGatewayAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// NAT网关端口转发规则对象数组。

		NatGatewayDestinationIpPortTranslationNatRuleSet []*NatGatewayDestinationIpPortTranslationNatRule `json:"NatGatewayDestinationIpPortTranslationNatRuleSet,omitempty" name:"NatGatewayDestinationIpPortTranslationNatRuleSet"`
		// 符合条件的NAT网关端口转发规则对象数目。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 专线网关对象。

		DirectConnectGateway *DirectConnectGateway `json:"DirectConnectGateway,omitempty" name:"DirectConnectGateway"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDirectConnectGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirectConnectGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressTemplateGroupAttributeRequest struct {
	*tchttp.BaseRequest

	// IP地址模板集合实例ID，例如：ipmg-2uw6ujo6。

	AddressTemplateGroupId *string `json:"AddressTemplateGroupId,omitempty" name:"AddressTemplateGroupId"`
	// IP地址模板集合名称。

	AddressTemplateGroupName *string `json:"AddressTemplateGroupName,omitempty" name:"AddressTemplateGroupName"`
	// IP地址模板实例ID， 例如：ipm-mdunqeb6。

	AddressTemplateIds []*string `json:"AddressTemplateIds,omitempty" name:"AddressTemplateIds"`
}

func (r *ModifyAddressTemplateGroupAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAddressTemplateGroupAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrafficPackageQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTrafficPackageQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrafficPackageQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 商品价格。

		Price *Price `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReturnNormalAddressesRequest struct {
	*tchttp.BaseRequest

	// 1

	AddressIps []*string `json:"AddressIps,omitempty" name:"AddressIps"`
}

func (r *ReturnNormalAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReturnNormalAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupPolicyCheckInfo struct {
	// 数据流方向，取值：inbound(入站)，outbound(出站)。

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 协议, 取值: TCP,UDP, ICMP。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 端口号，对于入站为目标端口，对于出站为源端口。

	Port *string `json:"Port,omitempty" name:"Port"`
	// 数据流是否方通，取值：ACCEPT(方通)，DROP(未方通)。

	Action *string `json:"Action,omitempty" name:"Action"`
}

type CreateDirectConnectGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// 专线网关ID，形如：dcg-prpqlmg1

	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	// 需要连通的IDC网段列表

	Routes []*DirectConnectGatewayCcnRoute `json:"Routes,omitempty" name:"Routes"`
}

func (r *CreateDirectConnectGatewayCcnRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirectConnectGatewayCcnRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerGatewayVendorsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 对端网关厂商信息对象。

		CustomerGatewayVendorSet []*CustomerGatewayVendor `json:"CustomerGatewayVendorSet,omitempty" name:"CustomerGatewayVendorSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomerGatewayVendorsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerGatewayVendorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIp6TranslatorQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户在指定地域的IPV6转换实例及规则配额信息
		// QUOTAID属性是TOTAL_TRANSLATOR_QUOTA，表示账户在指定地域的IPV6转换实例配额信息；QUOTAID属性是IPV6转转换实例唯一ID（形如ip6-xxxxxxxx），表示账户在该转换实例允许创建的转换规则配额

		QuotaSet []*Quota `json:"QuotaSet,omitempty" name:"QuotaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIp6TranslatorQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIp6TranslatorQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHaVipAttributeRequest struct {
	*tchttp.BaseRequest

	// `HAVIP`唯一`ID`，形如：`havip-9o233uri`。

	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`
	// `HAVIP`名称，可任意命名，但不得超过60个字符。

	HaVipName *string `json:"HaVipName,omitempty" name:"HaVipName"`
}

func (r *ModifyHaVipAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHaVipAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClassicLinkInstance struct {
	// VPC实例ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 云服务器实例唯一ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type AttachCcnInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachCcnInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomerGatewayVendor struct {
	// 平台。

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 软件版本。

	SoftwareVersion *string `json:"SoftwareVersion,omitempty" name:"SoftwareVersion"`
	// 供应商名称。

	VendorName *string `json:"VendorName,omitempty" name:"VendorName"`
}

type NetDetectState struct {
	// 网络探测实例ID。形如：netd-12345678。

	NetDetectId *string `json:"NetDetectId,omitempty" name:"NetDetectId"`
	// 网络探测目的IP验证结果对象数组。

	NetDetectIpStateSet []*NetDetectIpState `json:"NetDetectIpStateSet,omitempty" name:"NetDetectIpStateSet"`
}

type CreateIp6TranslatorsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 转换实例的唯一ID数组，形如"ip6-xxxxxxxx"

		Ip6TranslatorSet []*string `json:"Ip6TranslatorSet,omitempty" name:"Ip6TranslatorSet"`
		// TaskId，运维使用，无实际业务含义

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIp6TranslatorsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIp6TranslatorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNatGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNatGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowLogRequest struct {
	*tchttp.BaseRequest

	// 私用网络ID或者统一ID，建议使用统一ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 流日志唯一ID

	FlowLogId *string `json:"FlowLogId,omitempty" name:"FlowLogId"`
}

func (r *DescribeFlowLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpdateCcnBandwidthRequest struct {
	*tchttp.BaseRequest

	// 流量配置ID。

	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`
	// 云联网（CCN）地域带宽上限。

	MaxBandwidthLimit *int64 `json:"MaxBandwidthLimit,omitempty" name:"MaxBandwidthLimit"`
}

func (r *InquiryPriceUpdateCcnBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceUpdateCcnBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnInstance struct {
	// 关联实例类型，可选值：
	// <li>`VPC`：私有网络</li>
	// <li>`DIRECTCONNECT`：专线网关</li>
	// <li>`BMVPC`：黑石私有网络</li>

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 关联实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 关联实例ID所属大区，例如：ap-guangzhou。

	InstanceRegion *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`
}

type CreateAddressTemplateGroupRequest struct {
	*tchttp.BaseRequest

	// IP地址模版集合名称。

	AddressTemplateGroupName *string `json:"AddressTemplateGroupName,omitempty" name:"AddressTemplateGroupName"`
	// IP地址模版实例ID，例如：ipm-mdunqeb6。

	AddressTemplateIds []*string `json:"AddressTemplateIds,omitempty" name:"AddressTemplateIds"`
}

func (r *CreateAddressTemplateGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAddressTemplateGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpdateCcnBandwidthDealRequest struct {
	*tchttp.BaseRequest

	// 是否自动续费标识；NOTIFY_AND_AUTO_RENEW：自动续费，NOTIFY_AND_MANUAL_RENEW：手动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 流量配置ID。

	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`
	// 地域间设置带宽，单位：Mbps。

	MaxBandwidthLimit *int64 `json:"MaxBandwidthLimit,omitempty" name:"MaxBandwidthLimit"`
}

func (r *GetUpdateCcnBandwidthDealRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpdateCcnBandwidthDealRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssistantCidrResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的辅助CIDR数组。

		AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitempty" name:"AssistantCidrSet"`
		// 符合条件的实例数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssistantCidrResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssistantCidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnRegionBandwidthLimit struct {
	// 地域，例如：ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
	// 出带宽上限，单位：Mbps

	BandwidthLimit *uint64 `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`
	// 是否黑石地域，默认`false`。

	IsBm *bool `json:"IsBm,omitempty" name:"IsBm"`
	// 目的地域，例如：ap-shanghai

	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`
	// 目的地域是否为黑石地域，默认`false`。

	DstIsBm *bool `json:"DstIsBm,omitempty" name:"DstIsBm"`
}

type DescribeDirectConnectGatewaysRequest struct {
	*tchttp.BaseRequest

	// 专线网关唯一`ID`，形如：`dcg-9o233uri`。

	DirectConnectGatewayIds []*string `json:"DirectConnectGatewayIds,omitempty" name:"DirectConnectGatewayIds"`
	// 过滤条件，参数不支持同时指定`DirectConnectGatewayIds`和`Filters`。
	// <li>direct-connect-gateway-id - String - 专线网关唯一`ID`，形如：`dcg-9o233uri`。</li>
	// <li>direct-connect-gateway-name - String - 专线网关名称，默认模糊查询。</li>
	// <li>direct-connect-gateway-ip - String - 专线网关`IP`。</li>
	// <li>gateway-type - String - 网关类型，可选值：`NORMAL`（普通型）、`NAT`（NAT型）。</li>
	// <li>network-type- String - 网络类型，可选值：`VPC`（私有网络类型）、`CCN`（云联网类型）。</li>
	// <li>ccn-id - String - 专线网关所在云联网`ID`。</li>
	// <li>vpc-id - String - 专线网关所在私有网络`ID`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDirectConnectGatewaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectGatewaysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Ip6Translator struct {
	// IPV6转换实例唯一ID，形如ip6-xxxxxxxx

	Ip6TranslatorId *string `json:"Ip6TranslatorId,omitempty" name:"Ip6TranslatorId"`
	// IPV6转换实例名称

	Ip6TranslatorName *string `json:"Ip6TranslatorName,omitempty" name:"Ip6TranslatorName"`
	// IPV6地址

	Vip6 *string `json:"Vip6,omitempty" name:"Vip6"`
	// IPV6转换地址所属运营商

	IspName *string `json:"IspName,omitempty" name:"IspName"`
	// 转换实例状态，限于CREATING,RUNNING,DELETING,MODIFYING

	TranslatorStatus *string `json:"TranslatorStatus,omitempty" name:"TranslatorStatus"`
	// IPV6转换实例创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 绑定的IPV6转换规则数量

	Ip6RuleCount *int64 `json:"Ip6RuleCount,omitempty" name:"Ip6RuleCount"`
	// IPV6转换规则信息

	IP6RuleSet []*Ip6Rule `json:"IP6RuleSet,omitempty" name:"IP6RuleSet"`
	// IPV6转换实例创建本地时间

	CreatedTimeLocal *string `json:"CreatedTimeLocal,omitempty" name:"CreatedTimeLocal"`
}

type DescribeAddressesRequest struct {
	*tchttp.BaseRequest

	// 标识 EIP 的唯一 ID 列表。EIP 唯一 ID 形如：`eip-11112222`。参数不支持同时指定`AddressIds`和`Filters`。

	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`AddressIds`和`Filters`。详细的过滤条件如下：
	// <li> address-id - String - 是否必填：否 - （过滤条件）按照 EIP 的唯一 ID 过滤。EIP 唯一 ID 形如：eip-11112222。</li>
	// <li> address-name - String - 是否必填：否 - （过滤条件）按照 EIP 名称过滤。不支持模糊过滤。</li>
	// <li> address-ip - String - 是否必填：否 - （过滤条件）按照 EIP 的 IP 地址过滤。</li>
	// <li> address-status - String - 是否必填：否 - （过滤条件）按照 EIP 的状态过滤。状态包含：'CREATING'，'BINDING'，'BIND'，'UNBINDING'，'UNBIND'，'OFFLINING'，'BIND_ENI'。</li>
	// <li> instance-id - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的实例 ID 过滤。实例 ID 形如：ins-11112222。</li>
	// <li> private-ip-address - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的内网 IP 过滤。</li>
	// <li> network-interface-id - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的弹性网卡 ID 过滤。弹性网卡 ID 形如：eni-11112222。</li>
	// <li> is-arrears - String - 是否必填：否 - （过滤条件）按照 EIP 是否欠费进行过滤。（TRUE：EIP 处于欠费状态|FALSE：EIP 费用状态正常）</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// CCN路由策略对象。

		RouteSet []*CcnRoute `json:"RouteSet,omitempty" name:"RouteSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachClassicLinkVpcResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachClassicLinkVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachClassicLinkVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewVpnGatewayRequest struct {
	*tchttp.BaseRequest

	// VPN网关实例ID。

	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	// 预付费计费模式。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

func (r *RenewVpnGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewVpnGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceNatGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 商品价格。

		Price *Price `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceNatGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBandwidthPackageAttributeRequest struct {
	*tchttp.BaseRequest

	// 带宽包唯一标识ID

	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
	// 带宽包名称

	BandwidthPackageName *string `json:"BandwidthPackageName,omitempty" name:"BandwidthPackageName"`
}

func (r *ModifyBandwidthPackageAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBandwidthPackageAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteConflictsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 路由策略冲突列表

		RouteConflictSet []*RouteConflict `json:"RouteConflictSet,omitempty" name:"RouteConflictSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteConflictsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteConflictsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCcnBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单号。

		DealName *string `json:"DealName,omitempty" name:"DealName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCcnBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCcnBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupWithPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全组对象。

		SecurityGroup *SecurityGroup `json:"SecurityGroup,omitempty" name:"SecurityGroup"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityGroupWithPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupWithPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Vpc struct {
	// `VPC`名称。

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// `VPC`实例`ID`，例如：vpc-azd4dt1c。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// `VPC`的`IPv4` `CIDR`。

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 是否默认`VPC`。

	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
	// 是否开启组播。

	EnableMulticast *bool `json:"EnableMulticast,omitempty" name:"EnableMulticast"`
	// 创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// `DNS`列表。

	DnsServerSet []*string `json:"DnsServerSet,omitempty" name:"DnsServerSet"`
	// `DHCP`域名选项值。

	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
	// `DHCP`选项集`ID`。

	DhcpOptionsId *string `json:"DhcpOptionsId,omitempty" name:"DhcpOptionsId"`
	// 是否开启`DHCP`。

	EnableDhcp *bool `json:"EnableDhcp,omitempty" name:"EnableDhcp"`
	// `VPC`的`IPv6` `CIDR`。

	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`
	// 标签键值对

	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`
	// 辅助CIDR

	AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitempty" name:"AssistantCidrSet"`
}

type CreateCcnBandwidthRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 云联网（CCN）各地域带宽上限。

	CcnRegionBandwidthLimits []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimits,omitempty" name:"CcnRegionBandwidthLimits"`
}

func (r *CreateCcnBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCcnBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressSetRequest struct {
	*tchttp.BaseRequest

	// 弹性公网IP集群的标签

	TgwGroup *string `json:"TgwGroup,omitempty" name:"TgwGroup"`
	// 弹性公网IP集群的唯一ID

	SetId *int64 `json:"SetId,omitempty" name:"SetId"`
	// 弹性公网IP地址列表

	AddressIps []*string `json:"AddressIps,omitempty" name:"AddressIps"`
}

func (r *DescribeAddressSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceNatGatewayRequest struct {
	*tchttp.BaseRequest

	// NAT网关并发连接上限，默认值`1000000`

	MaxConcurrentConnection *uint64 `json:"MaxConcurrentConnection,omitempty" name:"MaxConcurrentConnection"`
}

func (r *InquiryPriceNatGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceNatGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetRoutesRequest struct {
	*tchttp.BaseRequest

	// 路由表实例ID，例如：rtb-azd4dt1c。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 路由表名称，最大长度不能超过60个字节。

	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
	// 路由策略。

	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

func (r *ResetRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyIp6AddressesBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceModifyIp6AddressesBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceModifyIp6AddressesBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 安全组实例ID，例如：sg-33ocnj9n，可通过DescribeSecurityGroups获取。每次请求的实例的上限为100。参数不支持同时指定SecurityGroupIds和Filters。

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 过滤条件，参数不支持同时指定SecurityGroupIds和Filters。
	// <li>security-group-id - String - （过滤条件）安全组ID。</li>
	// <li>project-id - Integer - （过滤条件）项目id。</li>
	// <li>security-group-name - String - （过滤条件）安全组名称。</li>
	// <li>tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。使用请参考示例2。</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例3。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量。

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 返回数量。

	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRenewCcnBandwidthDealResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单详情。

		Deal *string `json:"Deal,omitempty" name:"Deal"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRenewCcnBandwidthDealResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRenewCcnBandwidthDealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceAllocateAddressesRequest struct {
	*tchttp.BaseRequest

	// 1

	AddressCount *int64 `json:"AddressCount,omitempty" name:"AddressCount"`
	// 1

	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`
	// 1

	VipSet *string `json:"VipSet,omitempty" name:"VipSet"`
	// 1

	TgwGroup *string `json:"TgwGroup,omitempty" name:"TgwGroup"`
	// 1

	ApplySilence *int64 `json:"ApplySilence,omitempty" name:"ApplySilence"`
	// 1

	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
	// 1

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 1

	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitempty" name:"AddressChargePrepaid"`
	// 1

	DealId *string `json:"DealId,omitempty" name:"DealId"`
	// 1

	AddressType *string `json:"AddressType,omitempty" name:"AddressType"`
	// Anycast发布域，ANYCAST_ZONE_A|ANYCAST_ZONE_B，默认为当前地域可选的任一发布域。

	AnycastZone *string `json:"AnycastZone,omitempty" name:"AnycastZone"`
	// EIP可用区，用于指定可用区申请EIP。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 申请可绑定CLB的EIP，默认值为否。仅适用于AnycastEIP。

	ApplicableForCLB *bool `json:"ApplicableForCLB,omitempty" name:"ApplicableForCLB"`
	// 1

	Restrictive *bool `json:"Restrictive,omitempty" name:"Restrictive"`
}

func (r *InquiryPriceAllocateAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceAllocateAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddBandwidthPackageResourcesRequest struct {
	*tchttp.BaseRequest

	// 带宽包唯一标识ID，形如'bwp-xxxx'

	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
	// 带宽包类型，当前支持'BGP'类型，表示内部资源是BGP IP。

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// 资源类型，包括'Address', 'LoadBalance'

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源Id，形如'eip-xxxx', 'lb-xxxx'

	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 带宽包协议类型。当前支持'ipv4'和'ipv6'协议类型。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *AddBandwidthPackageResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddBandwidthPackageResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceTemplateAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceTemplateAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServiceTemplateAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleGroup struct {
	// Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// VpcId

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// Protocol

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// VPort

	VPort *int64 `json:"VPort,omitempty" name:"VPort"`
	// Domain

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// SetId

	SetId *int64 `json:"SetId,omitempty" name:"SetId"`
}

type DescribeIp6TranslatorQuotaRequest struct {
	*tchttp.BaseRequest

	// 待查询IPV6转换实例的唯一ID列表，形如ip6-xxxxxxxx

	Ip6TranslatorIds []*string `json:"Ip6TranslatorIds,omitempty" name:"Ip6TranslatorIds"`
}

func (r *DescribeIp6TranslatorQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIp6TranslatorQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 要解关联网络实例列表

	Instances []*CcnInstance `json:"Instances,omitempty" name:"Instances"`
}

func (r *DetachCcnInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachCcnInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIp6TranslatorsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合过滤条件的IPV6转换实例数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 符合过滤条件的IPV6转换实例详细信息

		Ip6TranslatorSet []*Ip6Translator `json:"Ip6TranslatorSet,omitempty" name:"Ip6TranslatorSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIp6TranslatorsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIp6TranslatorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网（CCN）各地域出带宽带宽详情。

		CcnBandwidthSet []*CcnBandwidthInfo `json:"CcnBandwidthSet,omitempty" name:"CcnBandwidthSet"`
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCcnRegionBandwidthLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigratePrivateIpAddressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigratePrivateIpAddressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigratePrivateIpAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnassignPrivateIpAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnassignPrivateIpAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnassignPrivateIpAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AvailableZone struct {
	// 可用区ID，如：ap-guangzhou-2。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区数字ID。

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区名称，如：广州二区。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 可用网络产品列表。

	Products []*string `json:"Products,omitempty" name:"Products"`
	// 白名单`key`列表。

	WhiteListKey []*string `json:"WhiteListKey,omitempty" name:"WhiteListKey"`
}

type EnableRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectGatewayAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDirectConnectGatewayAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDirectConnectGatewayAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartTrafficMirrorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartTrafficMirrorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartTrafficMirrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Resource struct {
	// 带宽包资源类型，包括'Address'和'LoadBalance'

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 带宽包资源Id，形如'eip-xxxx', 'lb-xxxx'

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 带宽包资源Ip

	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`
}

type CreateFlowLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建的流日志信息

		FlowLog []*FlowLog `json:"FlowLog,omitempty" name:"FlowLog"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFlowLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatGateway struct {
	// NAT网关的ID。

	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`
	// NAT网关的名称。

	NatGatewayName *string `json:"NatGatewayName,omitempty" name:"NatGatewayName"`
	// NAT网关创建的时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// NAT网关的状态。
	//  'PENDING'：生产中，'DELETING'：删除中，'AVAILABLE'：运行中，'UPDATING'：升级中，
	// ‘FAILED’：失败。

	State *string `json:"State,omitempty" name:"State"`
	// 网关最大外网出带宽(单位:Mbps)。

	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 网关并发连接上限。

	MaxConcurrentConnection *uint64 `json:"MaxConcurrentConnection,omitempty" name:"MaxConcurrentConnection"`
	// 绑定NAT网关的公网IP对象数组。

	PublicIpAddressSet []*NatGatewayAddress `json:"PublicIpAddressSet,omitempty" name:"PublicIpAddressSet"`
	// NAT网关网络状态。“AVAILABLE”:运行中, “UNAVAILABLE”:不可用, “INSUFFICIENT”:欠费停服。

	NetworkState *string `json:"NetworkState,omitempty" name:"NetworkState"`
	// NAT网关的端口转发规则。

	DestinationIpPortTranslationNatRuleSet []*DestinationIpPortTranslationNatRule `json:"DestinationIpPortTranslationNatRuleSet,omitempty" name:"DestinationIpPortTranslationNatRuleSet"`
	// VPC实例ID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// NAT网关所在的可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DeleteFlowLogRequest struct {
	*tchttp.BaseRequest

	// 私用网络ID或者统一ID，建议使用统一ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 流日志唯一ID

	FlowLogId *string `json:"FlowLogId,omitempty" name:"FlowLogId"`
}

func (r *DeleteFlowLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFlowLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpnGateway struct {
	// 网关实例ID。

	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	// VPC实例ID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 网关实例名称。

	VpnGatewayName *string `json:"VpnGatewayName,omitempty" name:"VpnGatewayName"`
	// 网关实例类型：'IPSEC', 'SSL'。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 网关实例状态， 'PENDING'：生产中，'DELETING'：删除中，'AVAILABLE'：运行中。

	State *string `json:"State,omitempty" name:"State"`
	// 网关公网IP。

	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`
	// 网关续费类型：'NOTIFY_AND_MANUAL_RENEW'：手动续费，'NOTIFY_AND_AUTO_RENEW'：自动续费，'NOT_NOTIFY_AND_NOT_RENEW'：到期不续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 网关付费类型：POSTPAID_BY_HOUR：按小时后付费，PREPAID：包年包月预付费，

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 网关出带宽。

	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 预付费网关过期时间。

	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// 公网IP是否被封堵。

	IsAddressBlocked *bool `json:"IsAddressBlocked,omitempty" name:"IsAddressBlocked"`
	// 计费模式变更，PREPAID_TO_POSTPAID：包年包月预付费到期转按小时后付费。

	NewPurchasePlan *string `json:"NewPurchasePlan,omitempty" name:"NewPurchasePlan"`
	// 网关计费装，PROTECTIVELY_ISOLATED：被安全隔离的实例，NORMAL：正常。

	RestrictState *string `json:"RestrictState,omitempty" name:"RestrictState"`
	// 可用区，如：ap-guangzhou-2

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 网关带宽配额信息

	VpnGatewayQuotaSet []*VpnGatewayQuota `json:"VpnGatewayQuotaSet,omitempty" name:"VpnGatewayQuotaSet"`
}

type DescribeDirectConnectGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 云联网路由（IDC网段）列表。

		RouteSet []*DirectConnectGatewayCcnRoute `json:"RouteSet,omitempty" name:"RouteSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectGatewayCcnRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectGatewayCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNatGatewayDestinationIpPortTranslationNatRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNatGatewayDestinationIpPortTranslationNatRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNatGatewayDestinationIpPortTranslationNatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpdateCcnBandwidthDealResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单详情。

		Deal *string `json:"Deal,omitempty" name:"Deal"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUpdateCcnBandwidthDealResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpdateCcnBandwidthDealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateNetworkInterfaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigrateNetworkInterfaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnConnectionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// VPN通道实例。

		VpnConnectionSet []*VpnConnection `json:"VpnConnectionSet,omitempty" name:"VpnConnectionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpnConnectionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCcnRegionBandwidthLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAssistantCidrResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAssistantCidrResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAssistantCidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTemplateGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 协议端口模板集合。

		ServiceTemplateGroupSet []*ServiceTemplateGroup `json:"ServiceTemplateGroupSet,omitempty" name:"ServiceTemplateGroupSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceTemplateGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceTemplateGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewCcnBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单号。

		DealName *string `json:"DealName,omitempty" name:"DealName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewCcnBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewCcnBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetTrafficMirrorFilterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetTrafficMirrorFilterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetTrafficMirrorFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnassignIpv6SubnetCidrBlockRequest struct {
	*tchttp.BaseRequest

	// 子网所在私有网络`ID`。形如：`vpc-f49l6u0z`。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// `IPv6` 子网段列表。

	Ipv6SubnetCidrBlocks []*Ipv6SubnetCidrBlock `json:"Ipv6SubnetCidrBlocks,omitempty" name:"Ipv6SubnetCidrBlocks"`
}

func (r *UnassignIpv6SubnetCidrBlockRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnassignIpv6SubnetCidrBlockRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartTrafficMirrorRequest struct {
	*tchttp.BaseRequest

	// 流量镜像实例ID

	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`
}

func (r *StartTrafficMirrorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartTrafficMirrorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TransformAddressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TransformAddressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TransformAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressTemplatesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>address-template-name - String - （过滤条件）IP地址模板名称。</li>
	// <li>address-template-id - String - （过滤条件）IP地址模板实例ID，例如：ipm-mdunqeb6。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAddressTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnBandwidthInfo struct {
	// 带宽所属的云联网ID。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 实例的创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 实例的过期时间

	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// 带宽实例的唯一ID。

	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`
	// 带宽是否自动续费的标记。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 描述带宽的地域和限速上限信息。

	CcnRegionBandwidthLimit *CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimit,omitempty" name:"CcnRegionBandwidthLimit"`
}

type DeleteSubnetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSubnetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrafficPackageStatisticsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTrafficPackageStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrafficPackageStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceModifyIp6AddressesBandwidthRequest struct {
	*tchttp.BaseRequest

	// IPV6地址。Ip6Addresses和Ip6AddressId必须且只能传一个

	Ip6Addresses []*string `json:"Ip6Addresses,omitempty" name:"Ip6Addresses"`
	// IPV6地址对应的唯一ID，形如eip-xxxxxxxx。Ip6Addresses和Ip6AddressId必须且只能传一个

	Ip6AddressIds []*string `json:"Ip6AddressIds,omitempty" name:"Ip6AddressIds"`
	// 修改的目标带宽，单位Mbps

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

func (r *InquiryPriceModifyIp6AddressesBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceModifyIp6AddressesBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RouteConflict struct {
	// 路由表实例ID，例如：rtb-azd4dt1c。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 要检查的与之冲突的目的端

	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
	// 冲突的路由策略列表

	ConflictSet []*Route `json:"ConflictSet,omitempty" name:"ConflictSet"`
}

type CreateRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新增的实例个数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 路由表对象。

		RouteTableSet []*RouteTable `json:"RouteTableSet,omitempty" name:"RouteTableSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResultRequest struct {
	*tchttp.BaseRequest

	// 异步任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 计费订单号

	DealName *string `json:"DealName,omitempty" name:"DealName"`
}

func (r *DescribeTaskResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveBandwidthPackageResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveBandwidthPackageResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveBandwidthPackageResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCcnRegionBandwidthLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateAddressRequest struct {
	*tchttp.BaseRequest

	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：`eip-11112222`。

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// 要绑定的实例 ID。实例 ID 形如：`ins-11112222`。可通过登录[控制台](https://console.cloud.tencent.com/cvm)查询，也可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 要绑定的弹性网卡 ID。 弹性网卡 ID 形如：`eni-11112222`。`NetworkInterfaceId` 与 `InstanceId` 不可同时指定。弹性网卡 ID 可通过登录[控制台](https://console.cloud.tencent.com/vpc/eni)查询，也可通过[DescribeNetworkInterfaces](https://cloud.tencent.com/document/api/215/15817)接口返回值中的`networkInterfaceId`获取。

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	// 要绑定的内网 IP。如果指定了 `NetworkInterfaceId` 则也必须指定 `PrivateIpAddress` ，表示将 EIP 绑定到指定弹性网卡的指定内网 IP 上。同时要确保指定的 `PrivateIpAddress` 是指定的 `NetworkInterfaceId` 上的一个内网 IP。指定弹性网卡的内网 IP 可通过登录[控制台](https://console.cloud.tencent.com/vpc/eni)查询，也可通过[DescribeNetworkInterfaces](https://cloud.tencent.com/document/api/215/15817)接口返回值中的`privateIpAddress`获取。

	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

func (r *AssociateAddressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateAddressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest

	// 限速实例的ID。

	RegionFlowControlIds []*string `json:"RegionFlowControlIds,omitempty" name:"RegionFlowControlIds"`
}

func (r *DeleteCcnRegionBandwidthLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// 专线网关ID，形如：dcg-prpqlmg1

	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	// 路由ID。形如：ccnr-f49l6u0z。

	RouteIds []*string `json:"RouteIds,omitempty" name:"RouteIds"`
}

func (r *DeleteDirectConnectGatewayCcnRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirectConnectGatewayCcnRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHaVipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteHaVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHaVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIp6RuleRequest struct {
	*tchttp.BaseRequest

	// IPV6转换实例唯一ID，形如ip6-xxxxxxxx

	Ip6TranslatorId *string `json:"Ip6TranslatorId,omitempty" name:"Ip6TranslatorId"`
	// IPV6转换规则唯一ID，形如rule6-xxxxxxxx

	Ip6RuleId *string `json:"Ip6RuleId,omitempty" name:"Ip6RuleId"`
	// IPV6转换规则修改后的名称

	Ip6RuleName *string `json:"Ip6RuleName,omitempty" name:"Ip6RuleName"`
	// IPV6转换规则修改后的IPV4地址

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// IPV6转换规则修改后的IPV4端口号

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
}

func (r *ModifyIp6RuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIp6RuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIp6TranslatorsRequest struct {
	*tchttp.BaseRequest

	// 待释放的IPV6转换实例的唯一ID，形如‘ip6-xxxxxxxx’

	Ip6TranslatorIds []*string `json:"Ip6TranslatorIds,omitempty" name:"Ip6TranslatorIds"`
}

func (r *DeleteIp6TranslatorsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIp6TranslatorsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// CCN对象。

		CcnSet []*CCN `json:"CcnSet,omitempty" name:"CcnSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkInterfaceAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetworkInterfaceAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetworkInterfaceAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Vpc对象。

		Vpc *Vpc `json:"Vpc,omitempty" name:"Vpc"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest struct {
	*tchttp.BaseRequest

	// NAT网关的ID，形如：`nat-df45454`。

	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`
	// NAT网关的端口转换规则。

	DestinationIpPortTranslationNatRules []*DestinationIpPortTranslationNatRule `json:"DestinationIpPortTranslationNatRules,omitempty" name:"DestinationIpPortTranslationNatRules"`
}

func (r *DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// 地域ID，如：ap-guangzhou。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域数字ID。

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域简称，如：gz。

	ShortName *string `json:"ShortName,omitempty" name:"ShortName"`
	// 地域名称，如：广州。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否中国大陆地域。

	IsChinaMainland *bool `json:"IsChinaMainland,omitempty" name:"IsChinaMainland"`
	// 是否金融地域。

	IsFinance *bool `json:"IsFinance,omitempty" name:"IsFinance"`
	// 可用区列表。

	AvailableZoneSet []*AvailableZone `json:"AvailableZoneSet,omitempty" name:"AvailableZoneSet"`
	// 白名单`key`列表。

	WhiteListKey []*string `json:"WhiteListKey,omitempty" name:"WhiteListKey"`
}

type AssignPrivateIpAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 内网IP详细信息。

		PrivateIpAddressSet []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet,omitempty" name:"PrivateIpAddressSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssignPrivateIpAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssignPrivateIpAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAddressAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddressTemplateSpecification struct {
	// IP地址ID，例如：ipm-2uw6ujo6。

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// IP地址组ID，例如：ipmg-2uw6ujo6。

	AddressGroupId *string `json:"AddressGroupId,omitempty" name:"AddressGroupId"`
}

type UpdateTrafficMirrorAllFilterRequest struct {
	*tchttp.BaseRequest

	// 流量镜像实例ID

	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`
	// 流量镜像采集对象

	CollectorSrcs []*string `json:"CollectorSrcs,omitempty" name:"CollectorSrcs"`
	// 流量镜像采集方向

	Direction *string `json:"Direction,omitempty" name:"Direction"`
	// 流量镜像需要过滤的natgw实例

	NatId *string `json:"NatId,omitempty" name:"NatId"`
	// 流量镜像需要过滤的五元组规则

	CollectorNormalFilters []*TrafficMirrorFilter `json:"CollectorNormalFilters,omitempty" name:"CollectorNormalFilters"`
}

func (r *UpdateTrafficMirrorAllFilterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTrafficMirrorAllFilterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustomerGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 对端网关对象

		CustomerGateway *CustomerGateway `json:"CustomerGateway,omitempty" name:"CustomerGateway"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCustomerGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomerGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PrivateIp6AddressesRequest struct {
	*tchttp.BaseRequest

	// IPV6地址。Ip6Addresses和Ip6AddressIds必须且只能传一个

	Ip6Addresses []*string `json:"Ip6Addresses,omitempty" name:"Ip6Addresses"`
	// IPV6地址对应的唯一ID，形如eip-xxxxxxxx。Ip6Addresses和Ip6AddressIds必须且只能传一个。

	Ip6AddressIds []*string `json:"Ip6AddressIds,omitempty" name:"Ip6AddressIds"`
}

func (r *PrivateIp6AddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PrivateIp6AddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssignPrivateIpAddressesRequest struct {
	*tchttp.BaseRequest

	// 弹性网卡实例ID，例如：eni-m6dyj72l。

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	// 指定的内网IP信息，单次最多指定10个。

	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 新申请的内网IP地址个数，内网IP地址个数总和不能超过配额数，详见<a href="/document/product/576/18527">弹性网卡使用限制</a>。

	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitempty" name:"SecondaryPrivateIpAddressCount"`
}

func (r *AssignPrivateIpAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssignPrivateIpAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicLinkInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 私有网络和基础网络互通设备。

		ClassicLinkInstanceSet []*ClassicLinkInstance `json:"ClassicLinkInstanceSet,omitempty" name:"ClassicLinkInstanceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClassicLinkInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClassicLinkInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcRequest struct {
	*tchttp.BaseRequest

	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DeleteVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// CCN路由策略唯一ID。形如：ccnr-f49l6u0z。

	RouteIds []*string `json:"RouteIds,omitempty" name:"RouteIds"`
}

func (r *DisableCcnRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableCcnRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest struct {
	*tchttp.BaseRequest

	// NAT网关的ID，形如：`nat-df45454`。

	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`
	// 源NAT网关的端口转换规则。

	SourceNatRule *DestinationIpPortTranslationNatRule `json:"SourceNatRule,omitempty" name:"SourceNatRule"`
	// 目的NAT网关的端口转换规则。

	DestinationNatRule *DestinationIpPortTranslationNatRule `json:"DestinationNatRule,omitempty" name:"DestinationNatRule"`
}

func (r *ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecurityGroupPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBandwidthAttributeRequest struct {
	*tchttp.BaseRequest

	// 当前地域的可用区ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 网络计费模式列表。可取
	// 'BANDWIDTH_POSTPAID_BY_MONTH',
	// 'BANDWIDTH_PREPAID_BY_MONTH',
	// 'TRAFFIC_POSTPAID_BY_HOUR',
	// 'BANDWIDTH_POSTPAID_BY_HOUR',
	// 'BANDWIDTH_PACKAGE'

	NetworkPayMode []*int64 `json:"NetworkPayMode,omitempty" name:"NetworkPayMode"`
	// 云服务器InstanceId，形如ins-xxx, 批量最大限制100个

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 弹性公网IP唯一ID，形如eip-xxx,批量最大限制100个

	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
}

func (r *DescribeBandwidthAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBandwidthAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全组策略检测结果集合

		SecurityGroupPolicyCheckInfoSet []*SecurityGroupPolicyCheckInfo `json:"SecurityGroupPolicyCheckInfoSet,omitempty" name:"SecurityGroupPolicyCheckInfoSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckSecurityGroupPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatGatewayQuota struct {
	// NAT网关带宽配额

	BandwidthQuota *uint64 `json:"BandwidthQuota,omitempty" name:"BandwidthQuota"`
}

type DescribeBandwidthPackageQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 带宽包配额数据结构

		QuotaSet []*Quota `json:"QuotaSet,omitempty" name:"QuotaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBandwidthPackageQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBandwidthPackageQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressesBandwidthRequest struct {
	*tchttp.BaseRequest

	// EIP唯一标识id，形如'eip-xxxx'

	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
	// 调整带宽目标值

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 包月带宽起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 包月带宽结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ModifyAddressesBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAddressesBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConflictItem struct {
	// 冲突资源的ID

	ConfilctId *string `json:"ConfilctId,omitempty" name:"ConfilctId"`
	// 冲突目的资源

	DestinationItem *string `json:"DestinationItem,omitempty" name:"DestinationItem"`
}

type Quota struct {
	// 配额名称，取值范围：<br><li>`TOTAL_EIP_QUOTA`：用户当前地域下EIP的配额数；<br><li>`DAILY_EIP_APPLY`：用户当前地域下今日申购次数；<br><li>`DAILY_PUBLIC_IP_ASSIGN`：用户当前地域下，重新分配公网 IP次数。

	QuotaId *string `json:"QuotaId,omitempty" name:"QuotaId"`
	// 当前数量

	QuotaCurrent *int64 `json:"QuotaCurrent,omitempty" name:"QuotaCurrent"`
	// 配额数量

	QuotaLimit *int64 `json:"QuotaLimit,omitempty" name:"QuotaLimit"`
}

type DeleteRoutesRequest struct {
	*tchttp.BaseRequest

	// 路由表实例ID。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 路由策略对象。

	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

func (r *DeleteRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGatewayQuotaRequest struct {
	*tchttp.BaseRequest

	// 可用区信息

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// NAT网关统一 ID，形如：“nat-abceef"

	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`
}

func (r *DescribeNatGatewayQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatGatewayQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的对象数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// VPC对象。

		VpcSet []*Vpc `json:"VpcSet,omitempty" name:"VpcSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDealStatusByNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单号对应的订单状态，UNPAID(未支付)，PAID(已支付)，DELIVERING(发货中)，DELIVER_SUCC(发货成功)，DELIVER_FAIL(发货失败)，REFUNDED(已退款)，CANCELED(已取消)，DENY_PAIDBYOTHERS(代付拒绝)，UNKOWN(未知状态)

		DealStatus *string `json:"DealStatus,omitempty" name:"DealStatus"`
		// 订单号

		DealName *string `json:"DealName,omitempty" name:"DealName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDealStatusByNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDealStatusByNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BandwidthLimitForCcnAlarmOnlyRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *BandwidthLimitForCcnAlarmOnlyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BandwidthLimitForCcnAlarmOnlyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNatGatewayRequest struct {
	*tchttp.BaseRequest

	// NAT网关名称

	NatGatewayName *string `json:"NatGatewayName,omitempty" name:"NatGatewayName"`
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// NAT网关最大外网出带宽(单位:Mbps)，支持的参数值：`20, 50, 100, 200, 500, 1000, 2000, 5000`，默认: `100Mbps`。

	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// NAT网关并发连接上限，支持参数值：`1000000、3000000、10000000`，默认值为`100000`。

	MaxConcurrentConnection *uint64 `json:"MaxConcurrentConnection,omitempty" name:"MaxConcurrentConnection"`
	// 需要申请的弹性IP个数，系统会按您的要求生产N个弹性IP，其中AddressCount和PublicAddresses至少传递一个。

	AddressCount *uint64 `json:"AddressCount,omitempty" name:"AddressCount"`
	// 绑定NAT网关的弹性IP数组，其中AddressCount和PublicAddresses至少传递一个。

	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`
	// 可用区，形如：`ap-guangzhou-1`。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *CreateNatGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNatGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInterfacesExtraRequest struct {
	*tchttp.BaseRequest

	// 弹性网卡实例ID查询。形如：eni-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定NetworkInterfaceIds和Filters。

	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" name:"NetworkInterfaceIds"`
	// 过滤条件，参数不支持同时指定NetworkInterfaceIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>network-interface-id - String - （过滤条件）弹性网卡实例ID，形如：eni-5k56k7k7。</li>
	// <li>address-ip - String - （过滤条件）弹性网卡内网IP。单IP模糊查询，多IP精确查询。</li>
	// <li>wan-ip - String - （过滤条件）弹性网卡公网IP。单IP模糊查询，多IP精确查询。</li>
	// <li>instance-id - String - （过滤条件）子机实例ID。</li>
	// <li>subnet-id - String - （过滤条件）子网ID。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNetworkInterfacesExtraRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkInterfacesExtraRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type InquiryPriceRenewVpnGatewayRequest struct {
	*tchttp.BaseRequest

	// VPN网关实例ID。

	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

func (r *InquiryPriceRenewVpnGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewVpnGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnassignIpv6CidrBlockResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnassignIpv6CidrBlockResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnassignIpv6CidrBlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Price struct {
	// 实例价格。

	InstancePrice *ItemPrice `json:"InstancePrice,omitempty" name:"InstancePrice"`
	// 网络价格。

	BandwidthPrice *ItemPrice `json:"BandwidthPrice,omitempty" name:"BandwidthPrice"`
}

type InquiryPriceRenewCcnBandwidthRequest struct {
	*tchttp.BaseRequest

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 流量配置ID。

	RegionFlowControlId *string `json:"RegionFlowControlId,omitempty" name:"RegionFlowControlId"`
}

func (r *InquiryPriceRenewCcnBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewCcnBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBandwidthPackageResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的 共享带宽包内资源 数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 共享带宽包内资源 详细信息列表。

		ResourceSet []*Resource `json:"ResourceSet,omitempty" name:"ResourceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBandwidthPackageResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBandwidthPackageResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GatewayFlowMonitor struct {
	// 私有网络ID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 网关类型。

	GatewayType *string `json:"GatewayType,omitempty" name:"GatewayType"`
	// 网关ID。

	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
	// 网关流量监控ID。

	GatewayFlowId *uint64 `json:"GatewayFlowId,omitempty" name:"GatewayFlowId"`
	// 网关流量监控开启时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type Ip6RuleInfo struct {
	// IPV6端口号，可在0~65535范围取值

	Vport6 *int64 `json:"Vport6,omitempty" name:"Vport6"`
	// 协议类型，支持TCP/UDP

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// IPV4地址

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// IPV4端口号，可在0~65535范围取值

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
}

type DeleteTrafficPackagesRequest struct {
	*tchttp.BaseRequest

	// 待删除的流量包唯一ID数组

	TrafficPackageIds []*string `json:"TrafficPackageIds,omitempty" name:"TrafficPackageIds"`
}

func (r *DeleteTrafficPackagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTrafficPackagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachNetworkInterfaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachNetworkInterfaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublicIp6AddressesRequest struct {
	*tchttp.BaseRequest

	// 需要开通internet访问能力的IPV6地址

	Ip6Addresses []*string `json:"Ip6Addresses,omitempty" name:"Ip6Addresses"`
	// 带宽，单位Mbps。默认是1Mbps

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
	// 网络计费模式。IPV6当前支持"TRAFFIC_POSTPAID_BY_HOUR"，默认是"TRAFFIC_POSTPAID_BY_HOUR"。

	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
}

func (r *PublicIp6AddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublicIp6AddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigratePrivateIpAddressRequest struct {
	*tchttp.BaseRequest

	// 当内网IP绑定的弹性网卡实例ID，例如：eni-m6dyj72l。

	SourceNetworkInterfaceId *string `json:"SourceNetworkInterfaceId,omitempty" name:"SourceNetworkInterfaceId"`
	// 待迁移的目的弹性网卡实例ID。

	DestinationNetworkInterfaceId *string `json:"DestinationNetworkInterfaceId,omitempty" name:"DestinationNetworkInterfaceId"`
	// 迁移的内网IP地址，例如：10.0.0.6。

	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

func (r *MigratePrivateIpAddressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigratePrivateIpAddressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIp6TranslatorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIp6TranslatorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIp6TranslatorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HaVipAssociateAddressIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *HaVipAssociateAddressIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *HaVipAssociateAddressIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectGatewayAttributeRequest struct {
	*tchttp.BaseRequest

	// 专线网关唯一`ID`，形如：`dcg-9o233uri`。

	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	// 专线网关名称，可任意命名，但不得超过60个字符。

	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitempty" name:"DirectConnectGatewayName"`
	// 云联网路由学习类型，可选值：`BGP`（自动学习）、`STATIC`（静态，即用户配置）。只有云联网类型专线网关且开启了BGP功能才支持修改`CcnRouteType`。

	CcnRouteType *string `json:"CcnRouteType,omitempty" name:"CcnRouteType"`
	// 专线网关开启BGP community属性。

	EnableBGPCommunity *bool `json:"EnableBGPCommunity,omitempty" name:"EnableBGPCommunity"`
}

func (r *ModifyDirectConnectGatewayAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDirectConnectGatewayAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddressTemplate struct {
	// IP地址模板名称。

	AddressTemplateName *string `json:"AddressTemplateName,omitempty" name:"AddressTemplateName"`
	// IP地址模板实例唯一ID。

	AddressTemplateId *string `json:"AddressTemplateId,omitempty" name:"AddressTemplateId"`
	// IP地址信息。

	AddressSet []*string `json:"AddressSet,omitempty" name:"AddressSet"`
	// 创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type AssignIpv6CidrBlockRequest struct {
	*tchttp.BaseRequest

	// `VPC`实例`ID`，形如：`vpc-f49l6u0z`。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *AssignIpv6CidrBlockRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssignIpv6CidrBlockRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest

	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 安全组规则集合。一个请求中只能删除单个方向的一条或多条规则。支持指定索引（PolicyIndex） 匹配删除和安全组规则匹配删除两种方式，一个请求中只能使用一种匹配方式。

	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *DeleteSecurityGroupPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// CCN所属UIN（根账号），默认当前账号所属UIN

	CcnUin *string `json:"CcnUin,omitempty" name:"CcnUin"`
	// 关联网络实例列表

	Instances []*CcnInstance `json:"Instances,omitempty" name:"Instances"`
}

func (r *AttachCcnInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachCcnInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTrafficMirrorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTrafficMirrorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTrafficMirrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatGatewayAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNatGatewayAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatGatewayAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnConnectionAttributeRequest struct {
	*tchttp.BaseRequest

	// VPN通道实例ID。形如：vpnx-f49l6u0z。

	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`
	// VPN通道名称，可任意命名，但不得超过60个字符。

	VpnConnectionName *string `json:"VpnConnectionName,omitempty" name:"VpnConnectionName"`
	// 预共享密钥。

	PreShareKey *string `json:"PreShareKey,omitempty" name:"PreShareKey"`
	// SPD策略组，例如：{"10.0.0.5/24":["172.123.10.5/16"]}，10.0.0.5/24是vpc内网段172.123.10.5/16是IDC网段。用户指定VPC内哪些网段可以和您IDC中哪些网段通信。

	SecurityPolicyDatabases []*SecurityPolicyDatabase `json:"SecurityPolicyDatabases,omitempty" name:"SecurityPolicyDatabases"`
	// IKE配置（Internet Key Exchange，因特网密钥交换），IKE具有一套自我保护机制，用户配置网络安全协议。

	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitempty" name:"IKEOptionsSpecification"`
	// IPSec配置，腾讯云提供IPSec安全会话设置。

	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitempty" name:"IPSECOptionsSpecification"`
}

func (r *ModifyVpnConnectionAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpnConnectionAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTemplatesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>service-template-name - String - （过滤条件）协议端口模板名称。</li>
	// <li>service-template-id - String - （过滤条件）协议端口模板实例ID，例如：ppm-e6dy460g。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeServiceTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkInterfaceAttachment struct {
	// 云主机实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 网卡在云主机实例内的序号。

	DeviceIndex *uint64 `json:"DeviceIndex,omitempty" name:"DeviceIndex"`
	// 云主机所有者账户信息。

	InstanceAccountId *string `json:"InstanceAccountId,omitempty" name:"InstanceAccountId"`
	// 绑定时间。

	AttachTime *string `json:"AttachTime,omitempty" name:"AttachTime"`
}

type DescribeEipStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 全地域EIP总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 各地域EIP详细统计

		EipStatisticsSet []*EipStatistics `json:"EipStatisticsSet,omitempty" name:"EipStatisticsSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEipStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEipStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnLimit struct {
	// 云联网配额类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 云联网配额数值

	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type DeleteAddressTemplateGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAddressTemplateGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAddressTemplateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网（CCN）各地域出带宽上限

		CcnRegionBandwidthLimitSet []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimitSet,omitempty" name:"CcnRegionBandwidthLimitSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnRegionBandwidthLimitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewVpnGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 商品价格。

		Price *Price `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewVpnGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenBgRsRequest struct {
	*tchttp.BaseRequest

	// 无

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 默认False

	OssTestUndo *bool `json:"OssTestUndo,omitempty" name:"OssTestUndo"`
	// 默认False

	NeedSetupRs *bool `json:"NeedSetupRs,omitempty" name:"NeedSetupRs"`
	// 无

	ClientType *string `json:"ClientType,omitempty" name:"ClientType"`
	// 无

	Sync *bool `json:"Sync,omitempty" name:"Sync"`
	// 默认False

	NeedNat *bool `json:"NeedNat,omitempty" name:"NeedNat"`
	// 默认False

	ClearRule *bool `json:"ClearRule,omitempty" name:"ClearRule"`
	// 默认False

	UntieGslb *bool `json:"UntieGslb,omitempty" name:"UntieGslb"`
	// 默认“”

	NatIsp *string `json:"NatIsp,omitempty" name:"NatIsp"`
	// 必传add

	OpType *string `json:"OpType,omitempty" name:"OpType"`
	// 必传

	RsType *string `json:"RsType,omitempty" name:"RsType"`
	// 默认False

	ForceSet *bool `json:"ForceSet,omitempty" name:"ForceSet"`
	// 默认ipv4

	IpType *string `json:"IpType,omitempty" name:"IpType"`
	// RuleList

	RuleList []*RuleList `json:"RuleList,omitempty" name:"RuleList"`
}

func (r *OpenBgRsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenBgRsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssistantCidr struct {
	// `VPC`实例`ID`。形如：`vpc-6v2ht8q5`

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 辅助CIDR。形如：`172.16.0.0/16`

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// 辅助CIDR类型（0：普通辅助CIDR，1：容器辅助CIDR），默认都是0。

	AssistantType *int64 `json:"AssistantType,omitempty" name:"AssistantType"`
	// 辅助CIDR拆分的子网。

	SubnetSet []*Subnet `json:"SubnetSet,omitempty" name:"SubnetSet"`
}

type ServiceTemplateSpecification struct {
	// 协议端口ID，例如：ppm-f5n1f8da。

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 协议端口组ID，例如：ppmg-f5n1f8da。

	ServiceGroupId *string `json:"ServiceGroupId,omitempty" name:"ServiceGroupId"`
}

type DescribeVpnConnectionsRequest struct {
	*tchttp.BaseRequest

	// VPN通道实例ID。形如：vpnx-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpnConnectionIds和Filters。

	VpnConnectionIds []*string `json:"VpnConnectionIds,omitempty" name:"VpnConnectionIds"`
	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定VpnConnectionIds和Filters。
	// <li>vpc-id - String - VPC实例ID，形如：`vpc-0a36uwkr`。</li>
	// <li>vpn-gateway-id - String - VPN网关实例ID，形如：`vpngw-p4lmqawn`。</li>
	// <li>customer-gateway-id - String - 对端网关实例ID，形如：`cgw-l4rblw63`。</li>
	// <li>vpn-connection-name - String - 通道名称，形如：`test-vpn`。</li>
	// <li>vpn-connection-id - String - 通道实例ID，形如：`vpnx-5p7vkch8"`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeVpnConnectionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnConnectionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigrateAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkInterfaceRequest struct {
	*tchttp.BaseRequest

	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 弹性网卡名称，最大长度不能超过60个字节。

	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`
	// 弹性网卡描述，可任意命名，但不得超过60个字符。

	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitempty" name:"NetworkInterfaceDescription"`
	// 弹性网卡所在的子网实例ID，例如：subnet-0ap8nwca。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 新申请的内网IP地址个数，内网IP地址个数总和不能超过配数。

	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitempty" name:"SecondaryPrivateIpAddressCount"`
	// 指定绑定的安全组，例如：['sg-1dd51d']。

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 指定的内网IP信息，单次最多指定10个。

	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateNetworkInterfaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkInterfaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpnGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// VPN网关对象

		VpnGateway *VpnGateway `json:"VpnGateway,omitempty" name:"VpnGateway"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpnGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseAddressesRequest struct {
	*tchttp.BaseRequest

	// 标识 EIP 的唯一 ID 列表。EIP 唯一 ID 形如：`eip-11112222`。

	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
}

func (r *ReleaseAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceRoutesRequest struct {
	*tchttp.BaseRequest

	// 路由表实例ID，例如：rtb-azd4dt1c。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 路由策略对象。需要指定路由策略ID（RouteId）。

	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

func (r *ReplaceRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReplaceRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteTableRequest struct {
	*tchttp.BaseRequest

	// 路由表实例ID，例如：rtb-azd4dt1c。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

func (r *DeleteRouteTableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRouteTableRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 弹性公网IP总数量

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 弹性公网IP使用数量

		UsedNum *int64 `json:"UsedNum,omitempty" name:"UsedNum"`
		// 弹性公网IP可用数量

		AvailableNum *int64 `json:"AvailableNum,omitempty" name:"AvailableNum"`
		// 弹性公网IP详细信息

		AddressList []*AddressBase `json:"AddressList,omitempty" name:"AddressList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcIpv6Address struct {
	// `VPC`内`IPv6`地址。

	Ipv6Address *string `json:"Ipv6Address,omitempty" name:"Ipv6Address"`
	// 所属子网 `IPv6` `CIDR`。

	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	// `IPv6`类型。

	Ipv6AddressType *string `json:"Ipv6AddressType,omitempty" name:"Ipv6AddressType"`
	// `IPv6`申请时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type CheckAssistantCidrResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 冲突资源信息数组。

		ConflictSourceSet []*ConflictSource `json:"ConflictSourceSet,omitempty" name:"ConflictSourceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAssistantCidrResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAssistantCidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkInterfaceAttributeRequest struct {
	*tchttp.BaseRequest

	// 弹性网卡实例ID，例如：eni-pxir56ns。

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	// 弹性网卡名称，最大长度不能超过60个字节。

	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`
	// 弹性网卡描述，可任意命名，但不得超过60个字符。

	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitempty" name:"NetworkInterfaceDescription"`
	// 指定绑定的安全组，例如:['sg-1dd51d']。

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

func (r *ModifyNetworkInterfaceAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetworkInterfaceAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnAttachedInstancesRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件：
	// <li>ccn-id - String -（过滤条件）CCN实例ID。</li>
	// <li>instance-type - String -（过滤条件）关联实例类型。</li>
	// <li>instance-region - String -（过滤条件）关联实例所属地域。</li>
	// <li>instance-id - String -（过滤条件）关联实例实例ID。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 云联网实例ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 排序字段。支持：`CcnId` `InstanceType` `InstanceId` `InstanceName` `InstanceRegion` `AttachedTime` `State`。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 排序方法。顺序：`ASC`，倒序：`DESC`。

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeCcnAttachedInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnAttachedInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInterfaceLimitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 弹性网卡配额

		EniQuantity *int64 `json:"EniQuantity,omitempty" name:"EniQuantity"`
		// 每个弹性网卡可以分配的ip配额

		EniPrivateIpAddressQuantity *int64 `json:"EniPrivateIpAddressQuantity,omitempty" name:"EniPrivateIpAddressQuantity"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkInterfaceLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkInterfaceLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHaVipAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHaVipAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHaVipAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpnConnection struct {
	// 通道实例ID。

	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`
	// 通道名称。

	VpnConnectionName *string `json:"VpnConnectionName,omitempty" name:"VpnConnectionName"`
	// VPC实例ID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// VPN网关实例ID。

	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	// 对端网关实例ID。

	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`
	// 预共享密钥。

	PreShareKey *string `json:"PreShareKey,omitempty" name:"PreShareKey"`
	// 通道传输协议。

	VpnProto *string `json:"VpnProto,omitempty" name:"VpnProto"`
	// 通道加密协议。

	EncryptProto *string `json:"EncryptProto,omitempty" name:"EncryptProto"`
	// 路由类型。

	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`
	// 创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 通道的生产状态，PENDING：生产中，AVAILABLE：运行中，DELETING：删除中。

	State *string `json:"State,omitempty" name:"State"`
	// 通道连接状态，AVAILABLE：已连接。

	NetStatus *string `json:"NetStatus,omitempty" name:"NetStatus"`
	// SPD。

	SecurityPolicyDatabaseSet []*SecurityPolicyDatabase `json:"SecurityPolicyDatabaseSet,omitempty" name:"SecurityPolicyDatabaseSet"`
	// IKE选项。

	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitempty" name:"IKEOptionsSpecification"`
	// IPSEC选择。

	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitempty" name:"IPSECOptionsSpecification"`
}

type DeleteAddressTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAddressTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAddressTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveIp6RulesRequest struct {
	*tchttp.BaseRequest

	// IPV6转换规则所属的转换实例唯一ID，形如ip6-xxxxxxxx

	Ip6TranslatorId *string `json:"Ip6TranslatorId,omitempty" name:"Ip6TranslatorId"`
	// 待删除IPV6转换规则，形如rule6-xxxxxxxx

	Ip6RuleIds []*string `json:"Ip6RuleIds,omitempty" name:"Ip6RuleIds"`
}

func (r *RemoveIp6RulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveIp6RulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDefaultVpcResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 默认VPC和子网ID

		Vpc *DefaultVpcSubnet `json:"Vpc,omitempty" name:"Vpc"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDefaultVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDefaultVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceRenewAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCcnResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网（CCN）对象。

		Ccn *CCN `json:"Ccn,omitempty" name:"Ccn"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCcnResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCcnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateCcnBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 商品价格。

		Price *ItemPrice `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateCcnBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceCreateCcnBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceTemplateRequest struct {
	*tchttp.BaseRequest

	// 协议端口模板实例ID，例如：ppm-e6dy460g。

	ServiceTemplateId *string `json:"ServiceTemplateId,omitempty" name:"ServiceTemplateId"`
}

func (r *DeleteServiceTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServiceTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressTemplateGroupsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>address-template-group-name - String - （过滤条件）IP地址模板集合名称。</li>
	// <li>address-template-group-id - String - （过滤条件）IP地址模板实集合例ID，例如：ipmg-mdunqeb6。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *string `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAddressTemplateGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressTemplateGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIp6AddressesRequest struct {
	*tchttp.BaseRequest

	// 标识 IPV6 的唯一 ID 列表。IPV6 唯一 ID 形如：`eip-11112222`。参数不支持同时指定`Ip6AddressIds`和`Filters`。

	Ip6AddressIds []*string `json:"Ip6AddressIds,omitempty" name:"Ip6AddressIds"`
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`AddressIds`和`Filters`。详细的过滤条件如下：
	// <li> address-ip - String - 是否必填：否 - （过滤条件）按照 EIP 的 IP 地址过滤。</li>
	// <li> network-interface-id - String - 是否必填：否 - （过滤条件）按照弹性网卡的唯一ID过滤。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeIp6AddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIp6AddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest

	// CCN实例ID。形如：ccn-f49l6u0z。

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 云联网（CCN）各地域出带宽上限。

	CcnRegionBandwidthLimits []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimits,omitempty" name:"CcnRegionBandwidthLimits"`
}

func (r *SetCcnRegionBandwidthLimitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrafficPackagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrafficPackagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrafficPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateBandwidthPackageRequest struct {
	*tchttp.BaseRequest

	// 无

	BandwidthPackageName *string `json:"BandwidthPackageName,omitempty" name:"BandwidthPackageName"`
	// 无

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// 无

	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
	// 无

	BandwidthPackageCount *uint64 `json:"BandwidthPackageCount,omitempty" name:"BandwidthPackageCount"`
	// 无

	InternetMaxBandwidth *int64 `json:"InternetMaxBandwidth,omitempty" name:"InternetMaxBandwidth"`
	// 无

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
}

func (r *InquiryPriceCreateBandwidthPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceCreateBandwidthPackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopTrafficMirrorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopTrafficMirrorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopTrafficMirrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleList struct {
	// RsList

	RsList []*RsList `json:"RsList,omitempty" name:"RsList"`
	// RuleGroup

	RuleGroup []*RuleGroup `json:"RuleGroup,omitempty" name:"RuleGroup"`
	// RsOsType

	RsOsType *string `json:"RsOsType,omitempty" name:"RsOsType"`
}

type CreateAssistantCidrResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 辅助CIDR数组。

		AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitempty" name:"AssistantCidrSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssistantCidrResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssistantCidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EdgeVport struct {
	// 逻辑接口名称

	VportName *string `json:"VportName,omitempty" name:"VportName"`
	// 逻辑端口ID

	VportId *string `json:"VportId,omitempty" name:"VportId"`
	// 逻辑端口的vlan

	VportVlanId *uint64 `json:"VportVlanId,omitempty" name:"VportVlanId"`
	// 逻辑端口绑定的物理端口名

	AttachedPorts []*string `json:"AttachedPorts,omitempty" name:"AttachedPorts"`
	// 逻辑端口的DHCP类型

	DhcpMode *string `json:"DhcpMode,omitempty" name:"DhcpMode"`
	// 逻辑端口的状态

	State *string `json:"State,omitempty" name:"State"`
	// 逻辑端口的DHCP地址池

	DhcpCidr *string `json:"DhcpCidr,omitempty" name:"DhcpCidr"`
	// 逻辑接口的IP地址（CIDR格式）。

	VportIpAddr *string `json:"VportIpAddr,omitempty" name:"VportIpAddr"`
	// Edge设备ID。

	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
}

type CreateRoutesRequest struct {
	*tchttp.BaseRequest

	// 路由表实例ID。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 路由策略对象。

	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

func (r *CreateRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadCustomerGatewayConfigurationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// XML格式配置信息。

		CustomerGatewayConfiguration *string `json:"CustomerGatewayConfiguration,omitempty" name:"CustomerGatewayConfiguration"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadCustomerGatewayConfigurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadCustomerGatewayConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceTemplateGroupAttributeRequest struct {
	*tchttp.BaseRequest

	// 协议端口模板集合实例ID，例如：ppmg-ei8hfd9a。

	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId,omitempty" name:"ServiceTemplateGroupId"`
	// 协议端口模板集合名称。

	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName,omitempty" name:"ServiceTemplateGroupName"`
	// 协议端口模板实例ID，例如：ppm-4dw6agho。

	ServiceTemplateIds []*string `json:"ServiceTemplateIds,omitempty" name:"ServiceTemplateIds"`
}

func (r *ModifyServiceTemplateGroupAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServiceTemplateGroupAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddressTemplateGroup struct {
	// IP地址模板集合名称。

	AddressTemplateGroupName *string `json:"AddressTemplateGroupName,omitempty" name:"AddressTemplateGroupName"`
	// IP地址模板集合实例ID，例如：ipmg-dih8xdbq。

	AddressTemplateGroupId *string `json:"AddressTemplateGroupId,omitempty" name:"AddressTemplateGroupId"`
	// IP地址模板ID。

	AddressTemplateIdSet []*string `json:"AddressTemplateIdSet,omitempty" name:"AddressTemplateIdSet"`
	// 创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type DeleteSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *DeleteSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNatGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// NAT网关对象数组。

		NatGatewaySet []*NatGateway `json:"NatGatewaySet,omitempty" name:"NatGatewaySet"`
		// 符合条件的 NAT网关对象数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNatGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 已开服的所有大区列表。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 地域信息列表。

		RegionSet []*RegionInfo `json:"RegionSet,omitempty" name:"RegionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpdateCcnBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 商品价格。

		Price *ItemPrice `json:"Price,omitempty" name:"Price"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceUpdateCcnBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquiryPriceUpdateCcnBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseAddressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务TaskId。可以使用[DescribeTaskResult](https://cloud.tencent.com/document/api/215/36271)接口查询任务状态。

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubnetAttributeRequest struct {
	*tchttp.BaseRequest

	// 子网实例ID。形如：subnet-pxir56ns。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 子网名称，最大长度不能超过60个字节。

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 子网是否开启广播。

	EnableBroadcast *string `json:"EnableBroadcast,omitempty" name:"EnableBroadcast"`
}

func (r *ModifySubnetAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubnetAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCcnBandwidthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网（CNN）地域带宽详情。

		CcnBandwidth *CcnBandwidthInfo `json:"CcnBandwidth,omitempty" name:"CcnBandwidth"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCcnBandwidthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCcnBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTaskRequest struct {
	*tchttp.BaseRequest

	// 异步任务请求返回的RequestId。

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *QueryTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// IP地址模版。

		AddressTemplateSet []*AddressTemplate `json:"AddressTemplateSet,omitempty" name:"AddressTemplateSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddressTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableRoutesRequest struct {
	*tchttp.BaseRequest

	// 路由表唯一ID。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	// 路由策略唯一ID。

	RouteIds []*uint64 `json:"RouteIds,omitempty" name:"RouteIds"`
}

func (r *DisableRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHaVipRequest struct {
	*tchttp.BaseRequest

	// `HAVIP`唯一`ID`，形如：`havip-9o233uri`。

	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`
}

func (r *DeleteHaVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHaVipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBandwidthPackageRequest struct {
	*tchttp.BaseRequest

	// 待删除带宽包bwpId

	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
}

func (r *DeleteBandwidthPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBandwidthPackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTrafficMirrorDirectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateTrafficMirrorDirectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTrafficMirrorDirectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Ipv6Address struct {
	// `IPv6`地址，形如：`3402:4e00:20:100:0:8cd9:2a67:71f3`

	Address *string `json:"Address,omitempty" name:"Address"`
	// 是否是主`IP`。

	Primary *bool `json:"Primary,omitempty" name:"Primary"`
	// `EIP`实例`ID`，形如：`eip-hxlqja90`。

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// 描述信息。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 公网IP是否被封堵。

	IsWanIpBlocked *bool `json:"IsWanIpBlocked,omitempty" name:"IsWanIpBlocked"`
	// `IPv6`地址状态：
	// <li>`PENDING`：生产中</li>
	// <li>`MIGRATING`：迁移中</li>
	// <li>`DELETING`：删除中</li>
	// <li>`AVAILABLE`：可用的</li>

	State *string `json:"State,omitempty" name:"State"`
}

type RenewAddressesRequest struct {
	*tchttp.BaseRequest

	// 1

	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
	// 1

	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitempty" name:"AddressChargePrepaid"`
	// 1

	DealId *string `json:"DealId,omitempty" name:"DealId"`
	// 1

	CurrentDeadline *string `json:"CurrentDeadline,omitempty" name:"CurrentDeadline"`
}

func (r *RenewAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewAddressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubnetAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubnetAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubnetAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetNatGatewayConnectionRequest struct {
	*tchttp.BaseRequest

	// NAT网关ID。

	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`
	// NAT网关并发连接上限，形如：1000000、3000000、10000000。

	MaxConcurrentConnection *uint64 `json:"MaxConcurrentConnection,omitempty" name:"MaxConcurrentConnection"`
}

func (r *ResetNatGatewayConnectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetNatGatewayConnectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetDetectIpState struct {
	// 探测目的IPv4地址。

	DetectDestinationIp *string `json:"DetectDestinationIp,omitempty" name:"DetectDestinationIp"`
	// 探测结果。
	// 0：成功；
	// -1：查询不到路由丢包；
	// -2：外出ACL丢包；
	// -3：IN ACL丢包；
	// -4：其他错误；

	State *int64 `json:"State,omitempty" name:"State"`
	// 时延，单位毫秒

	Delay *uint64 `json:"Delay,omitempty" name:"Delay"`
	// 丢包率

	PacketLossRate *uint64 `json:"PacketLossRate,omitempty" name:"PacketLossRate"`
}

type DescribeHaVipsRequest struct {
	*tchttp.BaseRequest

	// `HAVIP`唯一`ID`，形如：`havip-9o233uri`。

	HaVipIds []*string `json:"HaVipIds,omitempty" name:"HaVipIds"`
	// 过滤条件，参数不支持同时指定`HaVipIds`和`Filters`。
	// <li>havip-id - String - `HAVIP`唯一`ID`，形如：`havip-9o233uri`。</li>
	// <li>havip-name - String - `HAVIP`名称。</li>
	// <li>vpc-id - String - `HAVIP`所在私有网络`ID`。</li>
	// <li>subnet-id - String - `HAVIP`所在子网`ID`。</li>
	// <li>address-ip - String - `HAVIP`绑定的弹性公网`IP`。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeHaVipsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHaVipsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RouteTableAssociation struct {
	// 子网实例ID。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 路由表实例ID。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

type MigrateNetworkInterfaceRequest struct {
	*tchttp.BaseRequest

	// 弹性网卡实例ID，例如：eni-m6dyj72l。

	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	// 弹性网卡当前绑定的CVM实例ID。形如：ins-r8hr2upy。

	SourceInstanceId *string `json:"SourceInstanceId,omitempty" name:"SourceInstanceId"`
	// 待迁移的目的CVM实例ID。

	DestinationInstanceId *string `json:"DestinationInstanceId,omitempty" name:"DestinationInstanceId"`
}

func (r *MigrateNetworkInterfaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateNetworkInterfaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseIp6AddressesBandwidthRequest struct {
	*tchttp.BaseRequest

	// IPV6地址。Ip6Addresses和Ip6AddressIds必须且只能传一个

	Ip6Addresses []*string `json:"Ip6Addresses,omitempty" name:"Ip6Addresses"`
	// IPV6地址对应的唯一ID，形如eip-xxxxxxxx。Ip6Addresses和Ip6AddressIds必须且只能传一个。

	Ip6AddressIds []*string `json:"Ip6AddressIds,omitempty" name:"Ip6AddressIds"`
}

func (r *ReleaseIp6AddressesBandwidthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseIp6AddressesBandwidthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssignIpv6SubnetCidrBlockRequest struct {
	*tchttp.BaseRequest

	// 子网所在私有网络`ID`。形如：`vpc-f49l6u0z`。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 分配 `IPv6` 子网段列表。

	Ipv6SubnetCidrBlocks []*Ipv6SubnetCidrBlock `json:"Ipv6SubnetCidrBlocks,omitempty" name:"Ipv6SubnetCidrBlocks"`
}

func (r *AssignIpv6SubnetCidrBlockRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssignIpv6SubnetCidrBlockRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
