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

package v20180410

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type Filter struct {
	// 需要过滤的字段。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段的过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type TcapVif struct {
	// tcap通道id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// tcap通道实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// tcap通道名称

	VifName *string `json:"VifName,omitempty" name:"VifName"`
	// tcap通道带宽

	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 用户侧互联IP

	ConnLocalIp *string `json:"ConnLocalIp,omitempty" name:"ConnLocalIp"`
	// 腾讯侧互联IP

	ConnPeerIp *string `json:"ConnPeerIp,omitempty" name:"ConnPeerIp"`
	// 路由同步协议，默认bgp

	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`
	// bgp 秘钥

	BgpKey *string `json:"BgpKey,omitempty" name:"BgpKey"`
	// bgp Asn 号

	BgpAsn *string `json:"BgpAsn,omitempty" name:"BgpAsn"`
	// 接入点ID

	AccessPointId *uint64 `json:"AccessPointId,omitempty" name:"AccessPointId"`
	// tcap通道状态; Enum[0, 1, 2, 5] 分别表示[ready, applying, deploying, deleted]

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 申请时间

	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`
	// 是否计费

	ChargeOrNot *bool `json:"ChargeOrNot,omitempty" name:"ChargeOrNot"`
	// 就绪时间

	ReadyTime *string `json:"ReadyTime,omitempty" name:"ReadyTime"`
	// 物理接口类型

	PhysicalPortType *string `json:"PhysicalPortType,omitempty" name:"PhysicalPortType"`
	// 接入点名称：$city-$idc

	AccessPointName *string `json:"AccessPointName,omitempty" name:"AccessPointName"`
}

type DescribeDirectConnectTunnelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 专用通道列表

		DirectConnectTunnelSet []*DirectConnectTunnel `json:"DirectConnectTunnelSet,omitempty" name:"DirectConnectTunnelSet"`
		// 通道总数目，并非经条件过滤后的通道数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectTunnelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectTunnelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudAttachGatewayInfo struct {
	// 高速上云服务网关实例信息
	// CreateCloudAttachServiceGateway：入参不填
	// DeleteCloudAttachServiceGateway：入参必填
	// DescribeCloudAttachServiceGateways：入参选填
	// UpdateCloudAttachServiceGateway：入参必填

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 高速上云服务实例ID
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参选填
	// UpdateCloudAttachServiceGateway：入参不填

	UniqCasId *string `json:"UniqCasId,omitempty" name:"UniqCasId"`
	// 高速上云服务网关名称
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	Name *string `json:"Name,omitempty" name:"Name"`
	// 合作伙伴的AppId。只作为出参使用

	IapId *string `json:"IapId,omitempty" name:"IapId"`
	// 专线ID
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参选填
	// UpdateCloudAttachServiceGateway：入参不填

	DcId *string `json:"DcId,omitempty" name:"DcId"`
	// vlan ID
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参选填
	// UpdateCloudAttachServiceGateway：入参不填

	VlanId *uint64 `json:"VlanId,omitempty" name:"VlanId"`
	// 带宽信息，单位为MB
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参选填
	// UpdateCloudAttachServiceGateway：入参选填

	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 腾讯侧互联IP
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	ConnLocalIp *string `json:"ConnLocalIp,omitempty" name:"ConnLocalIp"`
	// 用户侧互联IP
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	ConnPeerIp *string `json:"ConnPeerIp,omitempty" name:"ConnPeerIp"`
	// 互联IP掩码
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	ConnIpMask *string `json:"ConnIpMask,omitempty" name:"ConnIpMask"`
	// 电路编码
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	CircuitNumber *string `json:"CircuitNumber,omitempty" name:"CircuitNumber"`
	// 网关路由类型，静态路由或者BGP路由
	// 取值范围：static，bgp
	// CreateCloudAttachServiceGateway：入参不填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`
	// 用户侧IP段信息
	// CreateCloudAttachServiceGateway：入参不填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	PeerIp *string `json:"PeerIp,omitempty" name:"PeerIp"`
	// BGP路由的ASN信息
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	BgpAsn *string `json:"BgpAsn,omitempty" name:"BgpAsn"`
	// BGP路由的KEY信息
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参选填

	BgpKey *string `json:"BgpKey,omitempty" name:"BgpKey"`
	// 高速上云专线通道实例ID。只作为出参使用

	UniqVifId *string `json:"UniqVifId,omitempty" name:"UniqVifId"`
	// 高速上云专线网关实例ID。只作为出参使用

	UniqVpgId *string `json:"UniqVpgId,omitempty" name:"UniqVpgId"`
	// 高速上云网关所属地域。
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参不填

	Region *string `json:"Region,omitempty" name:"Region"`
	// BGP状态。只作为出参使用

	BgpStatus *string `json:"BgpStatus,omitempty" name:"BgpStatus"`
	// 网关类型：
	// cas：高速上云网关类型
	// dcg：普通网关类型（VXLAN迁移用）
	// CreateCloudAttachServiceGateway：入参选填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参不填

	Type *string `json:"Type,omitempty" name:"Type"`
	// 网关所在region的zone名称
	// CreateCloudAttachServiceGateway：入参必填
	// DeleteCloudAttachServiceGateway：入参不填
	// DescribeCloudAttachServiceGateways：入参不填
	// UpdateCloudAttachServiceGateway：入参不填

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 用户AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

type SwitchPort struct {
	// 端口数字ID

	MemberPorts []*string `json:"MemberPorts,omitempty" name:"MemberPorts"`
	// 端口名称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 端口状态，AVAILABLE表示可用

	State *string `json:"State,omitempty" name:"State"`
	// 交换机ID

	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`
	// 链路聚合口名称

	TrunkName *string `json:"TrunkName,omitempty" name:"TrunkName"`
}

type RejectDirectConnectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectDirectConnectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectDirectConnectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RouteFilterPrefix struct {
	// 用户侧网段地址

	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
}

type Tag struct {
	// 标签键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ModifySwitchRequest struct {
	*tchttp.BaseRequest

	// 专线接入交换机ID

	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`
	// 专线接入点ID

	AccessPointId *uint64 `json:"AccessPointId,omitempty" name:"AccessPointId"`
	// 交换机名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 交换机型号

	Model *string `json:"Model,omitempty" name:"Model"`
	// 交换机厂商，比如为 cisco, h3c, huawei 之一，其他字符串不支持

	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`
	// 云控制器使用的账号名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 云控制器使用的密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 云控制器使用的管理IP

	ManagementIp *string `json:"ManagementIp,omitempty" name:"ManagementIp"`
	// 云控制器使用的管理端口

	ManagementPort *uint64 `json:"ManagementPort,omitempty" name:"ManagementPort"`
	// 专线接入交换机vxlan隧道的VTEP ip

	VTEPIP *string `json:"VTEPIP,omitempty" name:"VTEPIP"`
}

func (r *ModifySwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImplementDirectConnectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImplementDirectConnectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImplementDirectConnectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectDirectConnectRequest struct {
	*tchttp.BaseRequest

	// 专线ID

	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
}

func (r *RejectDirectConnectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectDirectConnectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessPointsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接入点信息。

		AccessPointSet []*AccessPoint `json:"AccessPointSet,omitempty" name:"AccessPointSet"`
		// 接入点数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessPointsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectTunnelsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件:
	// 参数不支持同时指定DirectConnectTunnelIds和Filters。
	// <li> DirectConnectTunnelId, 专线通道实例ID，如dcx-abcdefgh。</li>
	// <li> DirectConnectId, 物理专线实例ID，如，dc-abcdefgh。</li>
	// <li> State, 专线通道状态。</li>
	// <li> AppId。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 专用通道 ID数组

	DirectConnectTunnelIds []*string `json:"DirectConnectTunnelIds,omitempty" name:"DirectConnectTunnelIds"`
	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDirectConnectTunnelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectTunnelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchVendorsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 厂商列表

		VendorSet []*SwitchVendor `json:"VendorSet,omitempty" name:"VendorSet"`
		// 厂商总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSwitchVendorsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchVendorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TcapPartPriceDetail struct {
	// 折扣

	Discount *float64 `json:"Discount,omitempty" name:"Discount"`
	// 折后价

	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
	// 原价

	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
}

type TcapUnchargeWhiteList struct {
	// 白名单ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 白名单创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DirectConnect struct {
	// 物理专线ID。

	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
	// 物理专线的名称。

	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`
	// 物理专线的接入点ID。

	AccessPointId *string `json:"AccessPointId,omitempty" name:"AccessPointId"`
	// 物理专线的状态。
	// 待审核：PENDING
	// 驳回：REJECTED
	// 实施中：CONSTRUCTING
	// 运行中：AVAILABLE
	// 已删除：DELETED

	State *string `json:"State,omitempty" name:"State"`
	// 物理专线创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 物理专线的开通时间。

	EnabledTime *string `json:"EnabledTime,omitempty" name:"EnabledTime"`
	// 提供接入物理专线的运营商。ChinaTelecom：中国电信， ChinaMobile：中国移动，ChinaUnicom：中国联通， In-houseWiring：楼内线，ChinaOther：中国其他， InternationalOperator：境外其他。

	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`
	// 本地数据中心的地理位置。

	Location *string `json:"Location,omitempty" name:"Location"`
	// 物理专线接入接口带宽，单位为Mbps。

	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 用户侧物理专线接入端口类型,取值：100Base-T：百兆电口,1000Base-T（默认值）：千兆电口,1000Base-LX：千兆单模光口（10千米）,10GBase-T：万兆电口10GBase-LR：万兆单模光口（10千米），默认值，千兆单模光口（10千米）

	IdcPortType *string `json:"IdcPortType,omitempty" name:"IdcPortType"`
	// 运营商或者服务商为物理专线提供的电路编码。

	CircuitCode *string `json:"CircuitCode,omitempty" name:"CircuitCode"`
	// 冗余物理专线的ID。

	RedundantDirectConnectId *string `json:"RedundantDirectConnectId,omitempty" name:"RedundantDirectConnectId"`
	// 物理专线申请者姓名。默认从账户体系获取。

	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`
	// 物理专线申请者联系邮箱。默认从账户体系获取。

	CustomerContactMail *string `json:"CustomerContactMail,omitempty" name:"CustomerContactMail"`
	// 物理专线申请者联系号码。默认从账户体系获取。

	CustomerContactNumber *string `json:"CustomerContactNumber,omitempty" name:"CustomerContactNumber"`
	// 报障联系人。

	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitempty" name:"FaultReportContactPerson"`
	// 报障联系电话。

	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitempty" name:"FaultReportContactNumber"`
	// IDC所在城市

	IdcCity *string `json:"IdcCity,omitempty" name:"IdcCity"`
	// 云侧端口类型

	CloudPortType *string `json:"CloudPortType,omitempty" name:"CloudPortType"`
	// 接入点名称

	AccessPointName *string `json:"AccessPointName,omitempty" name:"AccessPointName"`
	// 申请ID

	ApplyId *uint64 `json:"ApplyId,omitempty" name:"ApplyId"`
	// 交换机名称

	SwitchName *string `json:"SwitchName,omitempty" name:"SwitchName"`
	// 专线所有者AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 交换机ID

	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`
	// 交换机端口名称

	SwitchPort *string `json:"SwitchPort,omitempty" name:"SwitchPort"`
	// 专线下就绪的通道数目

	TunnelCount *uint64 `json:"TunnelCount,omitempty" name:"TunnelCount"`
}

type QueryProductStateInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品集群列表和集群内节点列表

		Data *ProductStateInfo `json:"Data,omitempty" name:"Data"`
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

type ProductFailureMigrateTaskDetail struct {
	// 故障切换任务UUID

	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
	// 状态：'initial'、'doing'、'done'、'failed'

	Status *string `json:"Status,omitempty" name:"Status"`
	// 进度：范围：0～100

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// 执行信息/错误信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 详情信息补充

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 补充数据

	ExtendData *string `json:"ExtendData,omitempty" name:"ExtendData"`
}

type CreateAccessPointRequest struct {
	*tchttp.BaseRequest

	// 接入点名称

	AccessPointName *string `json:"AccessPointName,omitempty" name:"AccessPointName"`
	// 接入点的状态。AVAILABLE, UNAVAILABLE

	State *string `json:"State,omitempty" name:"State"`
	// 接入点所在机房地址

	Location *string `json:"Location,omitempty" name:"Location"`
	// 接入点所在城市

	City *string `json:"City,omitempty" name:"City"`
	// 接入点支持的运营商列表，"中国电信", u"中国移动", "中国联通", "中国其他", "境外其他"

	LineOperator []*string `json:"LineOperator,omitempty" name:"LineOperator"`
	// 接入点可用带宽百分比

	AvailableBandwidthRatio *uint64 `json:"AvailableBandwidthRatio,omitempty" name:"AvailableBandwidthRatio"`
	// 带宽告警比例

	NotifyBandwidthRatio *uint64 `json:"NotifyBandwidthRatio,omitempty" name:"NotifyBandwidthRatio"`
	// 接入点总带宽, 单位 MBytes

	TotalBandwidth *uint64 `json:"TotalBandwidth,omitempty" name:"TotalBandwidth"`
}

func (r *CreateAccessPointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessPointRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductHealthState struct {
	// 产品名称

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集群名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 集群状态，normal或abnormal

	Status *string `json:"Status,omitempty" name:"Status"`
	// 错误信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 健康状态指标列表

	Metrics []*Metrics `json:"Metrics,omitempty" name:"Metrics"`
}

type DescribeDirectConnectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 物理专线列表。

		DirectConnectSet []*DirectConnect `json:"DirectConnectSet,omitempty" name:"DirectConnectSet"`
		// 符合物理专线列表数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudAttachServiceGatewayPingTaskInfo struct {
	// 高速上云服务网关实例ID

	UniqCasGwId *string `json:"UniqCasGwId,omitempty" name:"UniqCasGwId"`
	// Ping探测IP协议类型

	AddressFamily *string `json:"AddressFamily,omitempty" name:"AddressFamily"`
	// 探测的对端IP地址

	ConnPeerIp *string `json:"ConnPeerIp,omitempty" name:"ConnPeerIp"`
	// 探测次数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 探测包大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 探测间隔，单位为毫秒

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 探测状态
	// running：任务正在进行
	// done：任务执行完毕
	// error：任务执行错误

	Status *string `json:"Status,omitempty" name:"Status"`
	// 探测任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// Ping探测成功率

	SuccessRatio *uint64 `json:"SuccessRatio,omitempty" name:"SuccessRatio"`
	// 最小延时

	LatencyMin *string `json:"LatencyMin,omitempty" name:"LatencyMin"`
	// 平均延时

	LatencyAvg *string `json:"LatencyAvg,omitempty" name:"LatencyAvg"`
	// 最大延时

	LatencyMax *string `json:"LatencyMax,omitempty" name:"LatencyMax"`
	// 任务开始时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 任务结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type AccessPoint struct {
	// 接入点的名称。

	AccessPointName *string `json:"AccessPointName,omitempty" name:"AccessPointName"`
	// 接入点唯一ID。

	AccessPointId *int64 `json:"AccessPointId,omitempty" name:"AccessPointId"`
	// 接入点的状态。AVAILABLE, UNAVAILABLE

	State *string `json:"State,omitempty" name:"State"`
	// 接入点所在机房地址

	Location *string `json:"Location,omitempty" name:"Location"`
	// 接入点支持的运营商列表。

	LineOperator []*string `json:"LineOperator,omitempty" name:"LineOperator"`
	// 接入点可用带宽百分比

	AvailableBandwidthRatio *uint64 `json:"AvailableBandwidthRatio,omitempty" name:"AvailableBandwidthRatio"`
	// 当前已使用带宽, 单位MBytes

	CurrentBandwidth *uint64 `json:"CurrentBandwidth,omitempty" name:"CurrentBandwidth"`
	// 带宽告警比例

	NotifyBandwidthRatio *uint64 `json:"NotifyBandwidthRatio,omitempty" name:"NotifyBandwidthRatio"`
	// 接入点总带宽, 单位 MBytes

	TotalBandwidth *uint64 `json:"TotalBandwidth,omitempty" name:"TotalBandwidth"`
	// IDC所在城市

	City *string `json:"City,omitempty" name:"City"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 物理专线用户侧接入端口类型,取值：100Base-T：百兆电口,1000Base-T（默认值）：千兆电口,1000Base-LX：千兆单模光口（10千米）,10GBase-T：万兆电口10GBase-LR：万兆单模光口（10千米），默认值，千兆单模光口（10千米）。

	CloudPortType []*string `json:"CloudPortType,omitempty" name:"CloudPortType"`
}

type DcNumberOfAR struct {
	// 接入点数字ID

	AccessPointId *uint64 `json:"AccessPointId,omitempty" name:"AccessPointId"`
	// 接入点下的专线数目

	NumberOfDirectConnects *uint64 `json:"NumberOfDirectConnects,omitempty" name:"NumberOfDirectConnects"`
}

type SwitchVendor struct {
	// 厂商名

	Manufacture *string `json:"Manufacture,omitempty" name:"Manufacture"`
}

type DescribeAvailableSwitchesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 专线接入交换机列表

		SwitchSet []*SwitchNetdevice `json:"SwitchSet,omitempty" name:"SwitchSet"`
		// 专线接入交换机数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableSwitchesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableSwitchesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchesRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 支持的过滤条件：AccessPointId、SwitchId

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeSwitchesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudAttachInfo struct {
	// 高速上云实例id
	// CreateCloudAttachService：入参不填
	// DeleteCloudAttachService：入参必填
	// CancelCloudAttachService：入参必填
	// DescribeCloudAttachServices：入参选填
	// UpdateCloudAttachService：入参必填
	// ConfirmCloudAttachService：入参必填
	// DeliveryCloudAttachService：入参必填

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 高速上云名称
	// CreateCloudAttachService：入参必填
	// DeleteCloudAttachService：入参选填
	// CancelCloudAttachService：入参不填
	// DescribeCloudAttachServices：入参选填
	// UpdateCloudAttachService：入参选填
	// ConfirmCloudAttachService：入参不填
	// DeliveryCloudAttachService：入参不填

	Name *string `json:"Name,omitempty" name:"Name"`
	// 合作伙伴的AppId
	// CreateCloudAttachService：入参不填
	// DeleteCloudAttachService：入参不填
	// CancelCloudAttachService：入参不填
	// DescribeCloudAttachServices：入参不填
	// UpdateCloudAttachService：入参不填
	// ConfirmCloudAttachService：入参不填
	// DeliveryCloudAttachService：入参不填

	IapId *string `json:"IapId,omitempty" name:"IapId"`
	// 需要接入高速上云的IDC的地址
	// CreateCloudAttachService：入参必填
	// DeleteCloudAttachService：入参不填
	// CancelCloudAttachService：入参不填
	// DescribeCloudAttachServices：入参不填
	// UpdateCloudAttachService：入参选填
	// ConfirmCloudAttachService：入参不填
	// DeliveryCloudAttachService：入参不填

	IdcAddress *string `json:"IdcAddress,omitempty" name:"IdcAddress"`
	// 需要接入高速上云的IDC的互联网服务提供商类型
	// CreateCloudAttachService：入参必填
	// DeleteCloudAttachService：入参不填
	// CancelCloudAttachService：入参不填
	// DescribeCloudAttachServices：入参不填
	// UpdateCloudAttachService：入参选填
	// ConfirmCloudAttachService：入参不填
	// DeliveryCloudAttachService：入参不填

	IdcType *string `json:"IdcType,omitempty" name:"IdcType"`
	// 高速上云的带宽，单位为MB
	// CreateCloudAttachService：入参必填
	// DeleteCloudAttachService：入参不填
	// CancelCloudAttachService：入参不填
	// DescribeCloudAttachServices：入参不填
	// UpdateCloudAttachService：入参不填
	// ConfirmCloudAttachService：入参不填
	// DeliveryCloudAttachService：入参不填

	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 联系电话
	// CreateCloudAttachService：入参必填
	// DeleteCloudAttachService：入参不填
	// CancelCloudAttachService：入参不填
	// DescribeCloudAttachServices：入参不填
	// UpdateCloudAttachService：入参不填
	// ConfirmCloudAttachService：入参不填
	// DeliveryCloudAttachService：入参不填

	Telephone *string `json:"Telephone,omitempty" name:"Telephone"`
	// 高速上云的状态
	// CreateCloudAttachService：入参不填
	// DeleteCloudAttachService：入参不填
	// CancelCloudAttachService：入参不填
	// DescribeCloudAttachServices：入参选填
	// UpdateCloudAttachService：入参不填
	// ConfirmCloudAttachService：入参必填
	// DeliveryCloudAttachService：入参不填
	// 出参状态：
	// available：就绪状态
	// applying：申请，待审核状态
	// pendingpay：代付款状态
	// building：建设中状态
	// confirming：待确认状态
	// isolate: 隔离状态
	// stoped：终止状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 高速上云申请的时间。只作为出参

	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`
	// 高速上云建设完成的时间。只作为出参

	ReadyTime *string `json:"ReadyTime,omitempty" name:"ReadyTime"`
	// 高速上云过期时间。只作为出参

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 备注信息
	// CreateCloudAttachService：入参必填
	// DeleteCloudAttachService：入参不填
	// CancelCloudAttachService：入参不填
	// DescribeCloudAttachServices：入参不填
	// UpdateCloudAttachService：入参选填
	// ConfirmCloudAttachService：入参不填
	// DeliveryCloudAttachService：入参不填

	Remarks *string `json:"Remarks,omitempty" name:"Remarks"`
	// 高速上云的地域状态。只作为出参
	// 出参状态：
	// same-region：同地域
	// cross-region：跨地域

	RegionStatus *string `json:"RegionStatus,omitempty" name:"RegionStatus"`
	// 用户的AppId。只作为出参

	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

type SwitchNetdevice struct {
	// 云控制器使用的管理IP

	ManagementIp *string `json:"ManagementIp,omitempty" name:"ManagementIp"`
	// 云控制器使用的管理端口，为空时，业务使用默认值

	ManagementPort *string `json:"ManagementPort,omitempty" name:"ManagementPort"`
	// 交换机名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 交换机厂商，比如为 cisco, h3c, huawei 之一，其他字符串不支持

	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`
	// 交换机型号

	Model *string `json:"Model,omitempty" name:"Model"`
	// 交换机序列号

	NetdevSn *string `json:"NetdevSn,omitempty" name:"NetdevSn"`
	// 网络设备固资号

	NetdevAssetId *string `json:"NetdevAssetId,omitempty" name:"NetdevAssetId"`
	// 网络设备用途

	NetdevFunc *string `json:"NetdevFunc,omitempty" name:"NetdevFunc"`
	// 网络设备机架名

	NetdevRackName *string `json:"NetdevRackName,omitempty" name:"NetdevRackName"`
	// 网络设备类型

	NetdevType *string `json:"NetdevType,omitempty" name:"NetdevType"`
	// 网络设备状态

	NetdevCurrentStatus *string `json:"NetdevCurrentStatus,omitempty" name:"NetdevCurrentStatus"`
	// 网络设备Id

	NetdevId *string `json:"NetdevId,omitempty" name:"NetdevId"`
	// 网络设备Idc

	NetdevIdc *string `json:"NetdevIdc,omitempty" name:"NetdevIdc"`
	// 网络设备IdcId

	NetdevIdcId *string `json:"NetdevIdcId,omitempty" name:"NetdevIdcId"`
}

type DescribeSwitchesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 专线接入交换机列表

		SwitchSet []*DcSwitch `json:"SwitchSet,omitempty" name:"SwitchSet"`
		// 专线接入交换机数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSwitchesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LocalIpTranslationNatRule struct {
	// 原始IP

	OriginalIp *string `json:"OriginalIp,omitempty" name:"OriginalIp"`
	// 映射IP

	TranslationIp *string `json:"TranslationIp,omitempty" name:"TranslationIp"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type TcapIpCertificate struct {
	// 证书ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 证书存储路径

	StoragePath *string `json:"StoragePath,omitempty" name:"StoragePath"`
	// 证书上的ipv4地址

	Ipv4Addresses *string `json:"Ipv4Addresses,omitempty" name:"Ipv4Addresses"`
	// 证书上的ipv6地址

	Ipv6Addresses *string `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`
	// 证书上的asn号

	Asn *string `json:"Asn,omitempty" name:"Asn"`
	// 证书上的ID

	CertificateNumber *string `json:"CertificateNumber,omitempty" name:"CertificateNumber"`
	// 证书生效时间

	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`
	// 证书过期时间

	ExpireDate *string `json:"ExpireDate,omitempty" name:"ExpireDate"`
	// 证书是否有效

	Valid *bool `json:"Valid,omitempty" name:"Valid"`
}

type DescribeSwitchPortsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 端口列表

		SwitchPortSet []*SwitchPort `json:"SwitchPortSet,omitempty" name:"SwitchPortSet"`
		// 端口数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSwitchPortsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchPortsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudAttachIapInfo struct {
	// 高速上云服务提供商的AppId信息

	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

type DirectConnectTunnel struct {
	// 专线通道ID

	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
	// 物理专线ID

	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
	// 专线通道状态
	// AVAILABLE 就绪
	// APPLYING 申请中
	// ALLOCATING 配置中
	// ALLOCATEFAIL 配置失败
	// ALLOCATED 配置完成
	// ALTERING 修改中
	// ALTERFAIL 修改失败
	// DELETING 删除中
	// DELETEFAIL 删除失败
	// DELETED 已删除
	// PENDING 待接受
	// REJECTED 拒绝

	State *string `json:"State,omitempty" name:"State"`
	// 专线的拥有者，开发商账号 ID

	OwnerAccount *int64 `json:"OwnerAccount,omitempty" name:"OwnerAccount"`
	// 网络类型，分别为VPC、BMVPC、CCN
	//  VPC：私有网络 ，BMVPC：黑石网络，CCN：云联网

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// 私有网络ID 或者黑石网络ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 专线网关 ID

	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	// BGP ：BGP路由 STATIC：静态

	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`
	// 用户侧BGP，Asn，AuthKey

	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`
	// 用户侧网段地址

	IdcRoutes []*string `json:"IdcRoutes,omitempty" name:"IdcRoutes"`
	// 专线通道的Vlan

	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`
	// 云侧互联IP

	CloudAddress *string `json:"CloudAddress,omitempty" name:"CloudAddress"`
	// 用户侧互联 IP

	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`
	// 专线通道名称

	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`
	// 专线通道创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 专线通道带宽值

	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 关联的网络自定义探测ID

	NetDetectId *string `json:"NetDetectId,omitempty" name:"NetDetectId"`
	// BGP community开关

	EnableBGPCommunity *bool `json:"EnableBGPCommunity,omitempty" name:"EnableBGPCommunity"`
	// 是否为Nat通道

	NatType *bool `json:"NatType,omitempty" name:"NatType"`
	// 通道连接的VPC所在地域

	VpcRegion *string `json:"VpcRegion,omitempty" name:"VpcRegion"`
	// 是否开启BFD

	BfdState *string `json:"BfdState,omitempty" name:"BfdState"`
	// 专线网关名称

	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitempty" name:"DirectConnectGatewayName"`
	// VPC名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// True: 开启, False: 不开启

	EnableMulticast *bool `json:"EnableMulticast,omitempty" name:"EnableMulticast"`
	// 通道下支持的组播组

	MulticastGroups []*string `json:"MulticastGroups,omitempty" name:"MulticastGroups"`
	// 专线通道的拥有者，开发上账号ID

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 关联通道ID

	RelatedDirectConnectTunnelId *string `json:"RelatedDirectConnectTunnelId,omitempty" name:"RelatedDirectConnectTunnelId"`
	// 通道负载均衡模式：
	// 0：None 1: LoadBalance 2: MasterSlave

	LoadMode *string `json:"LoadMode,omitempty" name:"LoadMode"`
	// 关联通道名称

	RelatedDirectConnectTunnelName *string `json:"RelatedDirectConnectTunnelName,omitempty" name:"RelatedDirectConnectTunnelName"`
	// 在主备模式下，通道是否为主

	MasterStatus *bool `json:"MasterStatus,omitempty" name:"MasterStatus"`
	// 私有网络统一 ID 或者黑石网络统一 ID

	VpcInstanceId *string `json:"VpcInstanceId,omitempty" name:"VpcInstanceId"`
	// 通道IP协议类型

	IpType *string `json:"IpType,omitempty" name:"IpType"`
}

type ModifyDirectConnectAttributeRequest struct {
	*tchttp.BaseRequest

	// 物理专线的ID。

	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
	// 运营商或者服务商为物理专线提供的电路编码。

	CircuitCode *string `json:"CircuitCode,omitempty" name:"CircuitCode"`
	// 报障联系人。

	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitempty" name:"FaultReportContactPerson"`
	// 报障联系电话。

	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitempty" name:"FaultReportContactNumber"`
	// 链路聚合口名称

	TrunkName *string `json:"TrunkName,omitempty" name:"TrunkName"`
	// 交换机ID

	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`
	// 端口数组

	MemberPorts []*string `json:"MemberPorts,omitempty" name:"MemberPorts"`
}

func (r *ModifyDirectConnectAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDirectConnectAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 专线接入交换机列表

		SwitchSet []*DcSwitch `json:"SwitchSet,omitempty" name:"SwitchSet"`
		// 专线接入交换机数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TcapAccessPoint struct {
	// 接入点 ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 接入点所在国家

	Country *string `json:"Country,omitempty" name:"Country"`
	// 接入点所在城市

	City *string `json:"City,omitempty" name:"City"`
	// 接入点所在机房

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 接入点类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 接入点当前是否可用

	Valid *bool `json:"Valid,omitempty" name:"Valid"`
}

type AcceptDirectConnectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AcceptDirectConnectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcceptDirectConnectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessPointResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接入点信息。

		AccessPointSet []*AccessPoint `json:"AccessPointSet,omitempty" name:"AccessPointSet"`
		// 接入点数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessPointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DcSwitch struct {
	// 专线接入点ID，通过DescribeAccessPoints接口可获取接入点列表

	AccessPointId *uint64 `json:"AccessPointId,omitempty" name:"AccessPointId"`
	// 用户自定义交换机名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 交换机型号

	Model *string `json:"Model,omitempty" name:"Model"`
	// 交换机厂商，比如为 cisco, h3c, huawei 之一，其他字符串不支持

	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`
	// 云控制器使用的管理IP

	ManagementIp *string `json:"ManagementIp,omitempty" name:"ManagementIp"`
	// 云控制器使用的管理端口

	ManagementPort *string `json:"ManagementPort,omitempty" name:"ManagementPort"`
	// 专线接入交换机vxlan隧道的VTEP ip

	VTEPIP *string `json:"VTEPIP,omitempty" name:"VTEPIP"`
	// 专线接入交换机ID

	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 云控制器使用的账号名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 云控制器使用的密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

type CreateProductFailureMigrateRequest struct {
	*tchttp.BaseRequest

	// 故障可用区名称

	CurrentAz *string `json:"CurrentAz,omitempty" name:"CurrentAz"`
	// 故障可用区所在地域

	CurrentRegion *string `json:"CurrentRegion,omitempty" name:"CurrentRegion"`
	// 故障迁移目标区域

	TargetAz *string `json:"TargetAz,omitempty" name:"TargetAz"`
	// 故障迁移目标地域

	TargetRegion *string `json:"TargetRegion,omitempty" name:"TargetRegion"`
	// 故障场景，如：脑裂、掉电等，需要统一约定好字典

	FailureScenario *string `json:"FailureScenario,omitempty" name:"FailureScenario"`
	// 操作类型，两种类型：1、切走；2、切回；

	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
	// 产品名称

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集群唯一名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 调用唯一凭证，避免重复调用

	CallerToken *string `json:"CallerToken,omitempty" name:"CallerToken"`
}

func (r *CreateProductFailureMigrateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductFailureMigrateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductStateInfoRequest struct {
	*tchttp.BaseRequest

	// 产品名，实际未使用，传空值即可

	Product *string `json:"Product,omitempty" name:"Product"`
	// 特殊参数

	Params *string `json:"Params,omitempty" name:"Params"`
}

func (r *QueryProductStateInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductStateInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudAttachPriceInfo struct {
	// 折扣

	Discount *string `json:"Discount,omitempty" name:"Discount"`
	// 折后价

	DiscountPrice *string `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
	// 原价

	OriginalPrice *string `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
}

type ClusterNode struct {
	// 集群唯一标识

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 集群状态：Abnormal, Normal

	Status *string `json:"Status,omitempty" name:"Status"`
	// 主节点所在AZ

	MainNodeAzList []*string `json:"MainNodeAzList,omitempty" name:"MainNodeAzList"`
	// 集群内节点列表

	NodeList []*ClusterMemberNode `json:"NodeList,omitempty" name:"NodeList"`
}

type TcapIps struct {
	// ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 对应的tcap 通道 ID

	TcapId *uint64 `json:"TcapId,omitempty" name:"TcapId"`
	// 4（ipv4）或者6（ipv6）

	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`
	// ip网段地址，多个用','分隔

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 网段审核状态; 可能取值为: [auditing, valid, invalid]

	Status *string `json:"Status,omitempty" name:"Status"`
	// 创建时间

	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`
	// 包含此网段的用户证书ID

	IpCertificateId *int64 `json:"IpCertificateId,omitempty" name:"IpCertificateId"`
	// ip 审核时间

	AuditTime *string `json:"AuditTime,omitempty" name:"AuditTime"`
	// 对应的tcap 通道 instance id

	TcapInstanceId *string `json:"TcapInstanceId,omitempty" name:"TcapInstanceId"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DeleteSwitchRequest struct {
	*tchttp.BaseRequest

	// 交换机ID，通过DescribeSwitches可获取到交换机列表

	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`
	// 交换机管理IP

	ManagementIp *string `json:"ManagementIp,omitempty" name:"ManagementIp"`
}

func (r *DeleteSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchPortsRequest struct {
	*tchttp.BaseRequest

	// 交换机ID, 通过DescribeSwitches可获取交换机列表

	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`
}

func (r *DescribeSwitchPortsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchPortsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TcapPriceDetail struct {
	// 铜牌流量价格信息

	BwBronze *TcapPartPriceDetail `json:"BwBronze,omitempty" name:"BwBronze"`
	// 银牌流量价格信息

	BwSilver *TcapPartPriceDetail `json:"BwSilver,omitempty" name:"BwSilver"`
	// 金牌流量价格信息

	BwGolden *TcapPartPriceDetail `json:"BwGolden,omitempty" name:"BwGolden"`
}

type CreateProductFailureMigrateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 切换任务TaskId

		Data *CreateProductFailureMigrateRsp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductFailureMigrateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductFailureMigrateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDirectConnectAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDirectConnectAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Metrics struct {
	// 指标名称，比如"NetReachable"

	Name *string `json:"Name,omitempty" name:"Name"`
	// 健康状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 健康信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type ImplementDirectConnectRequest struct {
	*tchttp.BaseRequest

	// 物理专线的ID。

	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
	// 运营商或者服务商为物理专线提供的电路编码。

	CircuitCode *string `json:"CircuitCode,omitempty" name:"CircuitCode"`
	// 报障联系人。

	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitempty" name:"FaultReportContactPerson"`
	// 报障联系电话。

	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitempty" name:"FaultReportContactNumber"`
}

func (r *ImplementDirectConnectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImplementDirectConnectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductHealthStateRequest struct {
	*tchttp.BaseRequest

	// 产品名称，实际未使用，传空值即可

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集群唯一名称

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

type InquiryCloudAttach struct {
	// 高速上云实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 询价购买月数

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 询价类型：新购create，续费renew

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeAccessPointsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 接入点ID(AccessPointId)数组

	AccessPointIds []*uint64 `json:"AccessPointIds,omitempty" name:"AccessPointIds"`
}

func (r *DescribeAccessPointsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessPointsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessPointRequest struct {
	*tchttp.BaseRequest

	// 接入点ID

	AccessPointId *uint64 `json:"AccessPointId,omitempty" name:"AccessPointId"`
	// 接入点名称

	AccessPointName *string `json:"AccessPointName,omitempty" name:"AccessPointName"`
	// 接入点的状态。AVAILABLE, UNAVAILABLE

	State *string `json:"State,omitempty" name:"State"`
	// 接入点所在机房地址

	Location *string `json:"Location,omitempty" name:"Location"`
	// 接入点所在城市

	City *string `json:"City,omitempty" name:"City"`
	// 接入点支持的运营商列表，"中国电信", u"中国移动", "中国联通", "中国其他", "境外其他"

	LineOperator []*string `json:"LineOperator,omitempty" name:"LineOperator"`
	// 接入点可用带宽百分比

	AvailableBandwidthRatio *uint64 `json:"AvailableBandwidthRatio,omitempty" name:"AvailableBandwidthRatio"`
	// 带宽告警比例

	NotifyBandwidthRatio *uint64 `json:"NotifyBandwidthRatio,omitempty" name:"NotifyBandwidthRatio"`
	// 接入点总带宽, 单位 MBytes

	TotalBandwidth *uint64 `json:"TotalBandwidth,omitempty" name:"TotalBandwidth"`
}

func (r *ModifyAccessPointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccessPointRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BgpPeer struct {
	// 用户侧，BGP Asn

	Asn *int64 `json:"Asn,omitempty" name:"Asn"`
	// 用户侧BGP密钥

	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`
}

type CloudAttachGetPingTask struct {
	// 高速上云网关实例ID

	UniqCasGwId *string `json:"UniqCasGwId,omitempty" name:"UniqCasGwId"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type QueryProductHealthStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// DCGW集群健康状态详情

		Data []*ProductHealthState `json:"Data,omitempty" name:"Data"`
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

type CloudAttachServiceGatewayPingInfo struct {
	// 高速上云服务网关实例ID

	UniqCasGwId *string `json:"UniqCasGwId,omitempty" name:"UniqCasGwId"`
	// Ping探测次数

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// Ping探测包的大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// Ping探测间隔，单位毫秒

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 需要探测的对端IP地址

	ConnPeerIp *string `json:"ConnPeerIp,omitempty" name:"ConnPeerIp"`
}

type ClusterMemberNode struct {
	// 内网IP

	PrivateIP *string `json:"PrivateIP,omitempty" name:"PrivateIP"`
	// 公网IP

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 机器名称

	Host *string `json:"Host,omitempty" name:"Host"`
	// 机架名称

	Rack *string `json:"Rack,omitempty" name:"Rack"`
	// 设备所在可用区名称

	Az *string `json:"Az,omitempty" name:"Az"`
	// 设备所在地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 设备状态：alive, dead

	Status *string `json:"Status,omitempty" name:"Status"`
}

type AcceptDirectConnectRequest struct {
	*tchttp.BaseRequest

	// 链路聚合口名称

	TrunkName *string `json:"TrunkName,omitempty" name:"TrunkName"`
	// 专线ID

	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
	// 交换机ID

	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`
	// 端口数组

	MemberPorts []*string `json:"MemberPorts,omitempty" name:"MemberPorts"`
	// 物理专线接入端口类型

	CloudPortType *string `json:"CloudPortType,omitempty" name:"CloudPortType"`
}

func (r *AcceptDirectConnectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcceptDirectConnectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductFailureMigrateTaskDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 迁移结果

		Data *ProductFailureMigrateTaskDetail `json:"Data,omitempty" name:"Data"`
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

type ISP struct {
	// ISP数字ID

	IspId *uint64 `json:"IspId,omitempty" name:"IspId"`
	// ISP名称，国际站AppId返回英文

	IspName *string `json:"IspName,omitempty" name:"IspName"`
}

type CreateSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 专线接入交换机列表

		SwitchSet []*DcSwitch `json:"SwitchSet,omitempty" name:"SwitchSet"`
		// 专线接入交换机数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchVendorsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSwitchVendorsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSwitchVendorsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件:支持State过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 物理专线 ID数组

	DirectConnectIds []*string `json:"DirectConnectIds,omitempty" name:"DirectConnectIds"`
	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDirectConnectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductStateInfo struct {
	// 产品名称标识

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集群列表

	ClusterList []*ClusterNode `json:"ClusterList,omitempty" name:"ClusterList"`
	// 集群信息更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type AccessPointListEntry struct {
	// 接入点数字ID

	AccessPointId *uint64 `json:"AccessPointId,omitempty" name:"AccessPointId"`
	// 地域名称

	Region *string `json:"Region,omitempty" name:"Region"`
	// vpc所在地域代号，比如"sh"

	VpcRegion *string `json:"VpcRegion,omitempty" name:"VpcRegion"`
	// 是否对客户可见，接入点已下线时返回false

	Visible *bool `json:"Visible,omitempty" name:"Visible"`
	// 接入点名称，调用AppId在国际站时，返回英文

	AccessPointName *string `json:"AccessPointName,omitempty" name:"AccessPointName"`
	// 接入点支持的运营商列表

	Isp []*ISP `json:"Isp,omitempty" name:"Isp"`
	// 接入点云端支持的端口列表

	PhysicalPort []*SwitchPort `json:"PhysicalPort,omitempty" name:"PhysicalPort"`
	// 接入点云端支持的端口ID列表

	PhysicalPortSort []*uint64 `json:"PhysicalPortSort,omitempty" name:"PhysicalPortSort"`
	// 接入点用户IDC侧支持的端口列表

	IdcPhysicalPortSort []*uint64 `json:"IdcPhysicalPortSort,omitempty" name:"IdcPhysicalPortSort"`
	// 接入点用户IDC侧支持的端口ID列表

	IdcPhysicalPort []*SwitchPort `json:"IdcPhysicalPort,omitempty" name:"IdcPhysicalPort"`
}

type CreateAccessPointResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前接入点列表

		AccessPointSet []*AccessPoint `json:"AccessPointSet,omitempty" name:"AccessPointSet"`
		// 当前接入点数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccessPointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductFailureMigrateRsp struct {
	// 切换任务UUID

	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
}

type CreateSwitchRequest struct {
	*tchttp.BaseRequest

	// 专线接入点ID，通过DescribeAccessPoints接口可获取接入点列表

	AccessPointId *uint64 `json:"AccessPointId,omitempty" name:"AccessPointId"`
	// 用户自定义交换机名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 交换机型号

	Model *string `json:"Model,omitempty" name:"Model"`
	// 交换机厂商，比如为 cisco, h3c, huawei 之一，其他字符串不支持

	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`
	// 云控制器使用的账号名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 云控制器使用的密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 云控制器使用的管理IP

	ManagementIp *string `json:"ManagementIp,omitempty" name:"ManagementIp"`
	// 云控制器使用的管理端口

	ManagementPort *uint64 `json:"ManagementPort,omitempty" name:"ManagementPort"`
	// 专线接入交换机vxlan隧道的VTEP ip

	VTEPIP *string `json:"VTEPIP,omitempty" name:"VTEPIP"`
}

func (r *CreateSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableSwitchesRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAvailableSwitchesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableSwitchesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductFailureMigrateTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 故障迁移任务UUID

	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
}

func (r *QueryProductFailureMigrateTaskDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductFailureMigrateTaskDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
