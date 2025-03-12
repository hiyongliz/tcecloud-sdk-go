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

package v20200103

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type RetRsList struct {
	// IP

	RsIp *string `json:"RsIp,omitempty" name:"RsIp"`
	// 端口

	RsPort *int64 `json:"RsPort,omitempty" name:"RsPort"`
	// 权重

	RsWeight *int64 `json:"RsWeight,omitempty" name:"RsWeight"`
	// V6Id

	V6Id *string `json:"V6Id,omitempty" name:"V6Id"`
	// VpcId

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// HostIp

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// ProbePort

	ProbePort *uint64 `json:"ProbePort,omitempty" name:"ProbePort"`
}

type DeleteBgRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBgRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBgRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdVipGroupCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 使用详情

		UsageList []*UsageList `json:"UsageList,omitempty" name:"UsageList"`
		// 组详情

		GroupTagList []*string `json:"GroupTagList,omitempty" name:"GroupTagList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApdVipGroupCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdVipGroupCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsageList struct {
	// 所属规则类型

	Belong *string `json:"Belong,omitempty" name:"Belong"`
	// 规则类型

	Category *string `json:"Category,omitempty" name:"Category"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type BgRsList struct {
	// RS ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// IP地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// Vpc ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 权重

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// 母机IP

	Host *string `json:"Host,omitempty" name:"Host"`
	// 母机IP

	Rshost *string `json:"Rshost,omitempty" name:"Rshost"`
	// V6Id

	V6Id *string `json:"V6Id,omitempty" name:"V6Id"`
}

type DescribeApdRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否全部显示

		AllAppShown *bool `json:"AllAppShown,omitempty" name:"AllAppShown"`
		// 规则列表

		RuleList []*RspRuleList `json:"RuleList,omitempty" name:"RuleList"`
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApdRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RspRuleList struct {
	// 心跳周期

	AliveInterval *int64 `json:"AliveInterval,omitempty" name:"AliveInterval"`
	// 业务名称

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 业务管理人

	BizAdmin *string `json:"BizAdmin,omitempty" name:"BizAdmin"`
	// 业务ID

	BizId *string `json:"BizId,omitempty" name:"BizId"`
	// 业务类型

	BizType *string `json:"BizType,omitempty" name:"BizType"`
	// 城市

	City *string `json:"City,omitempty" name:"City"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 转发协议

	FwdMode *string `json:"FwdMode,omitempty" name:"FwdMode"`
	// IDC名称

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// 运营商名称

	IspName *string `json:"IspName,omitempty" name:"IspName"`
	// 持续时间

	KeepTime *int64 `json:"KeepTime,omitempty" name:"KeepTime"`
	// 踢出间隔

	KickInterval *int64 `json:"KickInterval,omitempty" name:"KickInterval"`
	// 最大传输单元

	Mtu *int64 `json:"Mtu,omitempty" name:"Mtu"`
	// 探测间隔

	ProbeInterval *int64 `json:"ProbeInterval,omitempty" name:"ProbeInterval"`
	// 探测超时阈值

	ProbeTimeout *uint64 `json:"ProbeTimeout,omitempty" name:"ProbeTimeout"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 快速启用

	QuicEnable *int64 `json:"QuicEnable,omitempty" name:"QuicEnable"`
	// RS数量

	RsNum *uint64 `json:"RsNum,omitempty" name:"RsNum"`
	// RS列表

	RsList []*RetRsList `json:"RsList,omitempty" name:"RsList"`
	// 规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 时间类型

	ScheduleType *string `json:"ScheduleType,omitempty" name:"ScheduleType"`
	// Set名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// Tsv地址

	TsvVip *string `json:"TsvVip,omitempty" name:"TsvVip"`
	// 用途

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// IP地址

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// VPC ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 端口

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// 集群ID

	SetId *uint64 `json:"SetId,omitempty" name:"SetId"`
	// 是否为七层衍生规则

	FromL7 *bool `json:"FromL7,omitempty" name:"FromL7"`
	// 是否为TCS规则

	FromTcs *int64 `json:"FromTcs,omitempty" name:"FromTcs"`
	// 是否为租户端创建的规则

	FromClient *bool `json:"FromClient,omitempty" name:"FromClient"`
	// AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// Range

	Range *int64 `json:"Range,omitempty" name:"Range"`
	// 自定义健康检查发送内容

	SendContext *string `json:"SendContext,omitempty" name:"SendContext"`
	// 自定义健康检查接收内容

	RecvContext *string `json:"RecvContext,omitempty" name:"RecvContext"`
	// 运营商别名

	IspAlias *string `json:"IspAlias,omitempty" name:"IspAlias"`
}

type VipList struct {
	// 运营商名称

	Isp *string `json:"Isp,omitempty" name:"Isp"`
	// Set集群名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// VPC网络ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// IDC名称

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 规则个数

	RuleCount *uint64 `json:"RuleCount,omitempty" name:"RuleCount"`
	// 网段

	Segment *string `json:"Segment,omitempty" name:"Segment"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// IP地址

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// Tvs

	Tvs *string `json:"Tvs,omitempty" name:"Tvs"`
	// Tsv

	Tsv *string `json:"Tsv,omitempty" name:"Tsv"`
	// 运营商别名

	IspAlias *string `json:"IspAlias,omitempty" name:"IspAlias"`
}

type DescribeIspResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询运营商返回列表

		IspList []*IspList `json:"IspList,omitempty" name:"IspList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIspResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIspResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateBgCertificateRequest struct {
	*tchttp.BaseRequest

	// 证书名称

	CertName *string `json:"CertName,omitempty" name:"CertName"`
	// 业务名称

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 证书类型

	CertType *string `json:"CertType,omitempty" name:"CertType"`
	// 业务类型

	BizType *string `json:"BizType,omitempty" name:"BizType"`
	// 国际证书私钥或者国密证书证书私钥

	CertPublicCert *string `json:"CertPublicCert,omitempty" name:"CertPublicCert"`
	// 国际证书公钥或者国密证书证书公钥

	CertPrivateKey *string `json:"CertPrivateKey,omitempty" name:"CertPrivateKey"`
	// 业务ID

	BizId *string `json:"BizId,omitempty" name:"BizId"`
	// 旧证书属性

	OldCert *OldCert `json:"OldCert,omitempty" name:"OldCert"`
	// 登录账户

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 国密证书签名公钥

	SignCertPublicCert *string `json:"SignCertPublicCert,omitempty" name:"SignCertPublicCert"`
	// 国密证书签名私钥

	SignCertPrivateKey *string `json:"SignCertPrivateKey,omitempty" name:"SignCertPrivateKey"`
}

func (r *UpdateBgCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateBgCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateBgCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新证书相关信息

		NewCert *NewCertInfo `json:"NewCert,omitempty" name:"NewCert"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateBgCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateBgCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type L4SetList struct {
	// 容灾值

	DisasterRecovery *int64 `json:"DisasterRecovery,omitempty" name:"DisasterRecovery"`
	// 转发类型

	Fwd *string `json:"Fwd,omitempty" name:"Fwd"`
	// 所属IDC

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 运营商名称

	Isp *string `json:"Isp,omitempty" name:"Isp"`
	// LD列表

	LdList []*string `json:"LdList,omitempty" name:"LdList"`
	// Mater属性

	Master *Master `json:"Master,omitempty" name:"Master"`
	// Master名称

	MasterName *string `json:"MasterName,omitempty" name:"MasterName"`
	// Module名称

	Module *string `json:"Module,omitempty" name:"Module"`
	// 无用的TVS

	NouseTvs *int64 `json:"NouseTvs,omitempty" name:"NouseTvs"`
	// 系统类型

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 模式

	RouteMode *string `json:"RouteMode,omitempty" name:"RouteMode"`
	// 规则个数

	RuleCount *int64 `json:"RuleCount,omitempty" name:"RuleCount"`
	// 专区

	Section *string `json:"Section,omitempty" name:"Section"`
	// 专区列表

	SectionList []*string `json:"SectionList,omitempty" name:"SectionList"`
	// Set名称

	Set *string `json:"Set,omitempty" name:"Set"`
	// Set ID

	SetId *int64 `json:"SetId,omitempty" name:"SetId"`
	// 特殊Set 类型

	SpecialSetType *string `json:"SpecialSetType,omitempty" name:"SpecialSetType"`
	// Tsv IP数

	TsvVipCount *int64 `json:"TsvVipCount,omitempty" name:"TsvVipCount"`
	// Tvs IP数

	TvsVipCount *int64 `json:"TvsVipCount,omitempty" name:"TvsVipCount"`
	// IP类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 运营商别名

	IspAlias *string `json:"IspAlias,omitempty" name:"IspAlias"`
}

type SessionPersistence struct {
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 超时阈值

	TimeOut *int64 `json:"TimeOut,omitempty" name:"TimeOut"`
}

type CreateProductFailureMigrateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行产品故障转移返回结果

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

type CreateProductFailureMigrateRequest struct {
	*tchttp.BaseRequest

	// 产品集群主节点当前所在AZ

	CurrentAz *string `json:"CurrentAz,omitempty" name:"CurrentAz"`
	// 产品集群主节点当前所在区域

	CurrentRegion *string `json:"CurrentRegion,omitempty" name:"CurrentRegion"`
	// 产品集群主节点目标迁移AZ

	TargetAz *string `json:"TargetAz,omitempty" name:"TargetAz"`
	// 产品集群主节点目标迁移区域

	TargetRegion *string `json:"TargetRegion,omitempty" name:"TargetRegion"`
	// 故障场景

	FailureScenario *string `json:"FailureScenario,omitempty" name:"FailureScenario"`
	// 操作类型，两种类型：1、切走；2、切回；

	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
	// 产品名称

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集群名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 调用方唯一凭据，避免重复调用

	CallerToken *string `json:"CallerToken,omitempty" name:"CallerToken"`
}

func (r *CreateProductFailureMigrateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductFailureMigrateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdLdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回LD列表

		LdList []*RspLdList `json:"LdList,omitempty" name:"LdList"`
		// LD总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 服务器列表

		SvrArr []*string `json:"SvrArr,omitempty" name:"SvrArr"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApdLdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdLdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportApdRulePubCertResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 证书信息

		RulePubCertInfo *string `json:"RulePubCertInfo,omitempty" name:"RulePubCertInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportApdRulePubCertResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportApdRulePubCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActiveHealthCheck struct {
	// 下发周期

	DownTimes *int64 `json:"DownTimes,omitempty" name:"DownTimes"`
	// 期望在线时间

	ExpectAlive *uint64 `json:"ExpectAlive,omitempty" name:"ExpectAlive"`
	// 发送路径

	HttpSend *string `json:"HttpSend,omitempty" name:"HttpSend"`
	// 协议值

	Interval *int64 `json:"Interval,omitempty" name:"Interval"`
	// 方法

	Method *uint64 `json:"Method,omitempty" name:"Method"`
	// 协议

	Protocol *int64 `json:"Protocol,omitempty" name:"Protocol"`
	// 服务名称

	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`
	// 上传周期

	UpTimes *uint64 `json:"UpTimes,omitempty" name:"UpTimes"`
	// 响应超时时间

	RspTimeout *int64 `json:"RspTimeout,omitempty" name:"RspTimeout"`
	// HcIsSnat

	HcIsSnat *int64 `json:"HcIsSnat,omitempty" name:"HcIsSnat"`
}

type ProductFailureMigrateTaskRsp struct {
	// 切换任务UUID

	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 进度

	Progress *int64 `json:"Progress,omitempty" name:"Progress"`
	// 执行信息或错误信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 详情

	Detail []*string `json:"Detail,omitempty" name:"Detail"`
	// 补充数据

	ExtendData []*string `json:"ExtendData,omitempty" name:"ExtendData"`
}

type VirtualService struct {
	// app名称

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 运营名称

	BgName *string `json:"BgName,omitempty" name:"BgName"`
	// 业务编号

	BizId *string `json:"BizId,omitempty" name:"BizId"`
	// 业务类型

	BizType *string `json:"BizType,omitempty" name:"BizType"`
	// 是否允许自定义开关

	BlockedCustomizedField *string `json:"BlockedCustomizedField,omitempty" name:"BlockedCustomizedField"`
	// 是否阻塞开关

	BlockedEnable *int64 `json:"BlockedEnable,omitempty" name:"BlockedEnable"`
	// cc_guard开关

	CcGuardEnable *int64 `json:"CcGuardEnable,omitempty" name:"CcGuardEnable"`
	// 证书ID

	CertIds []*uint64 `json:"CertIds,omitempty" name:"CertIds"`
	// 城市

	City *string `json:"City,omitempty" name:"City"`
	// 默认服务

	DefaultServer *int64 `json:"DefaultServer,omitempty" name:"DefaultServer"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 转发协议

	ForwardProtocol *int64 `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`
	// 转发模式

	FwdMode *string `json:"FwdMode,omitempty" name:"FwdMode"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// IDC名称

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// LD端口

	LdPort *uint64 `json:"LdPort,omitempty" name:"LdPort"`
	// 集群名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 证书ID

	SngCertId *string `json:"SngCertId,omitempty" name:"SngCertId"`
	// 客户端证书ID

	SngCaCertId *string `json:"SngCaCertId,omitempty" name:"SngCaCertId"`
	// vpc网络ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 监听端口

	Vports []*uint64 `json:"Vports,omitempty" name:"Vports"`
	// Vip列表

	Vips []*CertVips `json:"Vips,omitempty" name:"Vips"`
	// 是否租户端创建

	FromClient *bool `json:"FromClient,omitempty" name:"FromClient"`
	// 是否由TCS创建

	FromTcs *uint64 `json:"FromTcs,omitempty" name:"FromTcs"`
	// 证书名称

	SngCertName *string `json:"SngCertName,omitempty" name:"SngCertName"`
	// 客户端证书名称

	SngCaCertName *string `json:"SngCaCertName,omitempty" name:"SngCaCertName"`
	// 限速阈值

	VipLimitRate *int64 `json:"VipLimitRate,omitempty" name:"VipLimitRate"`
	// 限速状态值

	VipLimitStatusCode *int64 `json:"VipLimitStatusCode,omitempty" name:"VipLimitStatusCode"`
	// 用户AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 用户名

	BizAdmin *string `json:"BizAdmin,omitempty" name:"BizAdmin"`
	// CdcId

	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
	// 类型

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// SslVerifyClientType

	SslVerifyClientType *int64 `json:"SslVerifyClientType,omitempty" name:"SslVerifyClientType"`
	// 七层个性化配置

	CustomziedConf *string `json:"CustomziedConf,omitempty" name:"CustomziedConf"`
}

type BindAppIdLabelRequest struct {
	*tchttp.BaseRequest

	// 集群标签

	Label *string `json:"Label,omitempty" name:"Label"`
	// 用户

	Owners []*string `json:"Owners,omitempty" name:"Owners"`
}

func (r *BindAppIdLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindAppIdLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApdVipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApdVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApdVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetApdIdcMasterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetApdIdcMasterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetApdIdcMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelVipList struct {
	// Tvs 地址

	Tvs *string `json:"Tvs,omitempty" name:"Tvs"`
	// Tsv 地址

	Tsv *string `json:"Tsv,omitempty" name:"Tsv"`
	// Vpc ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// Set名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 运营商名称

	Isp *string `json:"Isp,omitempty" name:"Isp"`
}

type SetApdLdStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetApdLdStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetApdLdStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JobStatusArr struct {
	// 任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 任务状态

	JobStatus *string `json:"JobStatus,omitempty" name:"JobStatus"`
	// 任务信息

	JobMsg *string `json:"JobMsg,omitempty" name:"JobMsg"`
}

type BindSetL4SetAndL7Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 七层set列表

		L7SetList []*BindL7RspSetList `json:"L7SetList,omitempty" name:"L7SetList"`
		// Set名称

		SetName *string `json:"SetName,omitempty" name:"SetName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindSetL4SetAndL7Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindSetL4SetAndL7Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdCertificateRequest struct {
	*tchttp.BaseRequest

	// 查询页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 查询页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 查询证书名称

	CertName *string `json:"CertName,omitempty" name:"CertName"`
	// 证书ID

	SngCertId *string `json:"SngCertId,omitempty" name:"SngCertId"`
	// 证书类型：CA/SVR

	CertType *string `json:"CertType,omitempty" name:"CertType"`
	// App名称

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 业务ID

	BizType *string `json:"BizType,omitempty" name:"BizType"`
}

func (r *DescribeApdCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIspRequest struct {
	*tchttp.BaseRequest

	// IPV4运营商可见性筛选，取值NO-拉取不可见的，YES-只拉取可见的，ALL-拉取所有的

	Visible *string `json:"Visible,omitempty" name:"Visible"`
	// IPV6运营商可见性筛选，取值NO-拉取不可见的，YES-只拉取可见的，ALL-拉取所有的

	VisibleIpv6 *string `json:"VisibleIpv6,omitempty" name:"VisibleIpv6"`
}

func (r *DescribeIspRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIspRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApdMasterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApdMasterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApdMasterResponse) FromJsonString(s string) error {
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

type DescribeApdVipGroupCategoryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeApdVipGroupCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdVipGroupCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgTargetHealthRequest struct {
	*tchttp.BaseRequest

	// clb实例唯一ID

	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

func (r *DescribeBgTargetHealthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgTargetHealthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSetLabelRequest struct {
	*tchttp.BaseRequest

	// 标签名

	Label *string `json:"Label,omitempty" name:"Label"`
	// 集群名

	SetNames []*string `json:"SetNames,omitempty" name:"SetNames"`
	// 集群类型，支持L4_LAN_CLB, L7_CLB, L7_LAN_CLB, L7_WAN_CLB

	SetTypes []*string `json:"SetTypes,omitempty" name:"SetTypes"`
}

func (r *DescribeSetLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSetLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpApdIdcResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpApdIdcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpApdIdcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDefaultIspListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDefaultIspListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDefaultIspListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApdMasterRequest struct {
	*tchttp.BaseRequest

	// IDC所在城市

	City *string `json:"City,omitempty" name:"City"`
	// 添加 Master 的类型，当前有四层、七层、NAT64三种类型

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 具有区分其他Master的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 主Master Ip，装了TGW MASTER模块的机器内网IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 备用Master Ip

	SlaveIp *string `json:"SlaveIp,omitempty" name:"SlaveIp"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
}

func (r *CreateApdMasterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApdMasterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgRuleRequest struct {
	*tchttp.BaseRequest

	// 查询页

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 查询页大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 查询RS详情过滤条件

	RuleList []*FRuleList `json:"RuleList,omitempty" name:"RuleList"`
	// 过滤参数，包括SearchAppName, SearchVip, SearchRs, BizId, SetName

	Filters []*FilterStr `json:"Filters,omitempty" name:"Filters"`
	// rs信息列表

	RsInfoList []*RsList `json:"RsInfoList,omitempty" name:"RsInfoList"`
}

func (r *DescribeBgRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpApdSetRequest struct {
	*tchttp.BaseRequest

	// 操作类型

	OpType *string `json:"OpType,omitempty" name:"OpType"`
	// set集群列表

	SetList []*SetList `json:"SetList,omitempty" name:"SetList"`
}

func (r *OpApdSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpApdSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpSectionList struct {
	// 专区名称

	SectionName *string `json:"SectionName,omitempty" name:"SectionName"`
	// 所属Module

	Module *string `json:"Module,omitempty" name:"Module"`
	// 单运营商Eth0描述

	Eth0 *string `json:"Eth0,omitempty" name:"Eth0"`
	// 运营商描述

	Isp *string `json:"Isp,omitempty" name:"Isp"`
	// Eth1地址

	Eth1 *string `json:"Eth1,omitempty" name:"Eth1"`
	// 多运营商Eth0描述

	MultiEth0 []*MultiEth0 `json:"MultiEth0,omitempty" name:"MultiEth0"`
}

type DeleteApdVipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApdVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApdVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否展示所有

		AllAppShown *bool `json:"AllAppShown,omitempty" name:"AllAppShown"`
		// Rule列表

		RuleList []*RspRuleList `json:"RuleList,omitempty" name:"RuleList"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBgRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApdVipGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApdVipGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApdVipGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindL7RspSetList struct {
	// 转发协议

	Fwd *string `json:"Fwd,omitempty" name:"Fwd"`
	// 是否绑定

	IsBind *bool `json:"IsBind,omitempty" name:"IsBind"`
	// LD列表

	LdList []*string `json:"LdList,omitempty" name:"LdList"`
	// SET ID

	SetId *int64 `json:"SetId,omitempty" name:"SetId"`
	// SET 名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
}

type SetL7LocationList struct {
	// Url地址

	Url *string `json:"Url,omitempty" name:"Url"`
	// 转发策略

	ForwardStrategy *ForwardStrategy `json:"ForwardStrategy,omitempty" name:"ForwardStrategy"`
	// 会话策略

	SessionPersistence *SessionPersistence `json:"SessionPersistence,omitempty" name:"SessionPersistence"`
	// Location Id

	LocationId *int64 `json:"LocationId,omitempty" name:"LocationId"`
	// 健康检查属性

	HealthCheck *SetL7HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`
	// 重定向信息

	ReWrite *ReWrite `json:"ReWrite,omitempty" name:"ReWrite"`
	// 修改权重rs_weight时传递

	RsList []*RsList `json:"RsList,omitempty" name:"RsList"`
}

type DeleteApdVipGroupRequest struct {
	*tchttp.BaseRequest

	// 被删除的VIP组列表

	VipGroupList []*DelVipGroupList `json:"VipGroupList,omitempty" name:"VipGroupList"`
}

func (r *DeleteApdVipGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApdVipGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductFailureMigrateTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 切换任务UUID

	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
}

func (r *QueryProductFailureMigrateTaskDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductFailureMigrateTaskDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ForwardStrategy struct {
	// 转发细节

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 转发类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type ReWriteV2Arg struct {
	// http返回码

	Code *uint64 `json:"Code,omitempty" name:"Code"`
	// 协议类型，http/https

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 端口

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// url

	Url *string `json:"Url,omitempty" name:"Url"`
	// TakeUrl

	TakeUrl *bool `json:"TakeUrl,omitempty" name:"TakeUrl"`
	// TakeArgs

	TakeArgs *bool `json:"TakeArgs,omitempty" name:"TakeArgs"`
}

type ServiceList struct {
	// Biz 编号

	BizId *string `json:"BizId,omitempty" name:"BizId"`
	// 主域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// vpc网络ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 端口

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// IP列表

	VipList []*string `json:"VipList,omitempty" name:"VipList"`
	// 服务位置列表

	LocationList []*LocationList `json:"LocationList,omitempty" name:"LocationList"`
}

type CreateApdVipRequest struct {
	*tchttp.BaseRequest

	// Set集群名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 网络运营商

	Isp *string `json:"Isp,omitempty" name:"Isp"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 网段信息

	NetmaskSeg []*string `json:"NetmaskSeg,omitempty" name:"NetmaskSeg"`
	// 操作类型,该接口固定为add

	OpType *string `json:"OpType,omitempty" name:"OpType"`
	// 组信息

	GroupTag *string `json:"GroupTag,omitempty" name:"GroupTag"`
	// 用途

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// IPv6 Vip类型，取值OIPV6(Underlay IPV6)、TIPV6(Overlay IPV6)

	VipType *string `json:"VipType,omitempty" name:"VipType"`
	// TvsIP列表

	TvsVipList []*string `json:"TvsVipList,omitempty" name:"TvsVipList"`
	// TsvIP列表

	TsvVipList []*string `json:"TsvVipList,omitempty" name:"TsvVipList"`
	// 网络掩码

	Netmask *string `json:"Netmask,omitempty" name:"Netmask"`
	// Idc信息

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// tvs网段信息

	TvsSegment *string `json:"TvsSegment,omitempty" name:"TvsSegment"`
}

func (r *CreateApdVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApdVipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LdList struct {
	// 所属Idc

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 内网IP

	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`
	// 机架号

	Rack *string `json:"Rack,omitempty" name:"Rack"`
	// 专区

	Section *string `json:"Section,omitempty" name:"Section"`
	// 机器型号

	SvrType *string `json:"SvrType,omitempty" name:"SvrType"`
	// 资产编号

	SvrId *string `json:"SvrId,omitempty" name:"SvrId"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 外网IP

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// 是否IPV6, 0或1

	IsIpv6 *uint64 `json:"IsIpv6,omitempty" name:"IsIpv6"`
}

type LocationSet struct {
	// 自定义配置

	CustomizedConf *string `json:"CustomizedConf,omitempty" name:"CustomizedConf"`
	// 转发策略

	ForwardStrategy *ForwardStrategyOut `json:"ForwardStrategy,omitempty" name:"ForwardStrategy"`
	// 健康检查属性

	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`
	// ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// Keepalive链接数量

	KeepaliveConnNum *uint64 `json:"KeepaliveConnNum,omitempty" name:"KeepaliveConnNum"`
	// Keepalive开关

	KeepaliveEnable *uint64 `json:"KeepaliveEnable,omitempty" name:"KeepaliveEnable"`
	// RS集群

	RsSet []*RsSet `json:"RsSet,omitempty" name:"RsSet"`
	// 会话策略

	SessionPersistenceStrategy *SessionPersistenceStrategy `json:"SessionPersistenceStrategy,omitempty" name:"SessionPersistenceStrategy"`
	// 类型

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// URL地址

	Url *string `json:"Url,omitempty" name:"Url"`
	// 转发类型

	ForwardType *int64 `json:"ForwardType,omitempty" name:"ForwardType"`
	// 绑定的rs列表

	RsList []*RsList `json:"RsList,omitempty" name:"RsList"`
	// 规则类型

	LocType *int64 `json:"LocType,omitempty" name:"LocType"`
	// 重定向信息

	Rewrite *ReWrite `json:"Rewrite,omitempty" name:"Rewrite"`
}

type RspRsList struct {
	// Server IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// Server 类型

	ServerType *string `json:"ServerType,omitempty" name:"ServerType"`
}

type SectionList struct {
	// Section 名称

	SectionName *string `json:"SectionName,omitempty" name:"SectionName"`
	// Eth0地址列表

	Eth0 []*MultiEth0 `json:"Eth0,omitempty" name:"Eth0"`
	// Eth1地址

	Eth1 *string `json:"Eth1,omitempty" name:"Eth1"`
	// 网络运营商

	Isp *string `json:"Isp,omitempty" name:"Isp"`
	// Set列表

	SetList []*RspSetList `json:"SetList,omitempty" name:"SetList"`
	// Tvs数量

	TvsCount *uint64 `json:"TvsCount,omitempty" name:"TvsCount"`
	// Tsv数量

	TsvCount *uint64 `json:"TsvCount,omitempty" name:"TsvCount"`
	// Ld数量

	LdCount *uint64 `json:"LdCount,omitempty" name:"LdCount"`
}

type DescribeApdAppNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询返回的申请名称

		AppNames []*string `json:"AppNames,omitempty" name:"AppNames"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApdAppNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdAppNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgJobRequest struct {
	*tchttp.BaseRequest

	// 任务ID数组

	JobId []*string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeBgJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterStr struct {
	// 过滤参数名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤参数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type LbLimitSet struct {
	// 管理员

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 配额

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 最大配额

	MaxLimit *int64 `json:"MaxLimit,omitempty" name:"MaxLimit"`
	// 最后修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 分配给子账号的总配额

	SubLimitSum *int64 `json:"SubLimitSum,omitempty" name:"SubLimitSum"`
	// 主账号已使用配额

	CurCount *int64 `json:"CurCount,omitempty" name:"CurCount"`
}

type DeleteApdLdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApdLdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApdLdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeL7RuleExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 本次返回的规则总数

		RulesCount *int64 `json:"RulesCount,omitempty" name:"RulesCount"`
		// 规则详情

		Rules []*FRuleList `json:"Rules,omitempty" name:"Rules"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeL7RuleExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeL7RuleExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSetLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群标签名和集群对应关系

		LabelSetNames []*LabelSetName `json:"LabelSetNames,omitempty" name:"LabelSetNames"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSetLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSetLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpApdModuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpApdModuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpApdModuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductStateInfoRequest struct {
	*tchttp.BaseRequest

	// 产品名称,取CLB

	Product *string `json:"Product,omitempty" name:"Product"`
	// 特殊参数,预留无实际意义

	Params []*string `json:"Params,omitempty" name:"Params"`
}

func (r *QueryProductStateInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductStateInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CertList struct {
	// 业务编号

	BizId *string `json:"BizId,omitempty" name:"BizId"`
	// 证书可选域域名

	CertAlternativeName *string `json:"CertAlternativeName,omitempty" name:"CertAlternativeName"`
	// 证书申请ID

	CertApplyId *string `json:"CertApplyId,omitempty" name:"CertApplyId"`
	// 证书运营名称

	CertBgName *string `json:"CertBgName,omitempty" name:"CertBgName"`
	// 证书主域名

	CertCommonName *string `json:"CertCommonName,omitempty" name:"CertCommonName"`
	// 证书启用时间

	CertEnableTime *string `json:"CertEnableTime,omitempty" name:"CertEnableTime"`
	// 证书过期时间

	CertExpiredTime *string `json:"CertExpiredTime,omitempty" name:"CertExpiredTime"`
	// 国际证书公钥或国密证书证书公钥

	CertPubkeyInfo *string `json:"CertPubkeyInfo,omitempty" name:"CertPubkeyInfo"`
	// 证书内容

	CertPublicCert *string `json:"CertPublicCert,omitempty" name:"CertPublicCert"`
	// 证书备注

	CertRemark *string `json:"CertRemark,omitempty" name:"CertRemark"`
	// 证书签名算法

	CertSignAlgorithm *string `json:"CertSignAlgorithm,omitempty" name:"CertSignAlgorithm"`
	// 证书种类

	CertSpecies *string `json:"CertSpecies,omitempty" name:"CertSpecies"`
	// 证书上传时间

	CertUploadTime *string `json:"CertUploadTime,omitempty" name:"CertUploadTime"`
	// 证书上传着

	CertUploader *string `json:"CertUploader,omitempty" name:"CertUploader"`
	// ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 证书ID

	SngCertId *string `json:"SngCertId,omitempty" name:"SngCertId"`
	// 业务类型

	BizType *string `json:"BizType,omitempty" name:"BizType"`
	// app名称

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 证书名称

	CertName *string `json:"CertName,omitempty" name:"CertName"`
	// 绑定规则详情

	Rules []*Rules `json:"Rules,omitempty" name:"Rules"`
	// 国密证书签名公钥

	SignCertPublicCert *string `json:"SignCertPublicCert,omitempty" name:"SignCertPublicCert"`
	// SignId

	SignId *uint64 `json:"SignId,omitempty" name:"SignId"`
}

type PassiveHealthCheck struct {
	// 失败时间阈值

	FailTimeout *uint64 `json:"FailTimeout,omitempty" name:"FailTimeout"`
	// 最大允许失败次数

	MaxFails *uint64 `json:"MaxFails,omitempty" name:"MaxFails"`
}

type DescribeBgCertificateRequest struct {
	*tchttp.BaseRequest

	// 查询页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 查询页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// App名称

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 证书名称

	CertName *string `json:"CertName,omitempty" name:"CertName"`
	// 证书ID

	SngCertId *string `json:"SngCertId,omitempty" name:"SngCertId"`
	// 证书类型

	CertType *string `json:"CertType,omitempty" name:"CertType"`
	// 业务ID

	BizType *string `json:"BizType,omitempty" name:"BizType"`
}

func (r *DescribeBgCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpApdStgwSetRequest struct {
	*tchttp.BaseRequest

	// 操作类型

	OpType *string `json:"OpType,omitempty" name:"OpType"`
	// Set集群列表

	SetList []*SetList7 `json:"SetList,omitempty" name:"SetList"`
}

func (r *OpApdStgwSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpApdStgwSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 过滤器的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤器的值数组

	Value *string `json:"Value,omitempty" name:"Value"`
}

type Filters struct {
	// value 为String 的过滤条件

	FilterStr *FilterStr `json:"FilterStr,omitempty" name:"FilterStr"`
	// value 为StringArray 的过滤条件

	FilterStrArray *FilterStrArray `json:"FilterStrArray,omitempty" name:"FilterStrArray"`
	// value 为FilterRuleList 的过滤条件

	FilterRuleList []*FRuleList `json:"FilterRuleList,omitempty" name:"FilterRuleList"`
}

type RsSet struct {
	// RS 地址

	RsIp *string `json:"RsIp,omitempty" name:"RsIp"`
	// RS端口

	RsPort *uint64 `json:"RsPort,omitempty" name:"RsPort"`
	// RS权重

	RsWeight *int64 `json:"RsWeight,omitempty" name:"RsWeight"`
	// 主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// v6 ID

	V6Id *string `json:"V6Id,omitempty" name:"V6Id"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// Vpc网络ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeBgLocationL7Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则列表

		Rules []*Rules `json:"Rules,omitempty" name:"Rules"`
		// 规则数量

		RulesCount *uint64 `json:"RulesCount,omitempty" name:"RulesCount"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Requestid

		Requestid *string `json:"Requestid,omitempty" name:"Requestid"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBgLocationL7Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgLocationL7Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgRsByRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// RS端口可操作标识

		OpAllRsPort *bool `json:"OpAllRsPort,omitempty" name:"OpAllRsPort"`
		// 查询的RS 详情

		RsList []*RuleOutRsList `json:"RsList,omitempty" name:"RsList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBgRsByRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgRsByRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OldCert struct {
	// 证书ID

	SngCertId *string `json:"SngCertId,omitempty" name:"SngCertId"`
	// 旧的证书名称

	OldCertName *string `json:"OldCertName,omitempty" name:"OldCertName"`
}

type Master struct {
	// 城市

	City *string `json:"City,omitempty" name:"City"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// IDC列表

	IdcList []*string `json:"IdcList,omitempty" name:"IdcList"`
	// IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 模式

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 备用Ip

	SlaveIp *string `json:"SlaveIp,omitempty" name:"SlaveIp"`
	// 备用端口

	SlavePort *uint64 `json:"SlavePort,omitempty" name:"SlavePort"`
	// 集群列表

	SetList []*string `json:"SetList,omitempty" name:"SetList"`
	// 7层服务器属性

	L7Slave []*L7Slave `json:"L7Slave,omitempty" name:"L7Slave"`
}

type DeleteBgCertificateRequest struct {
	*tchttp.BaseRequest

	// 证书名称

	CertName *string `json:"CertName,omitempty" name:"CertName"`
	// 证书ID

	SngCertId *string `json:"SngCertId,omitempty" name:"SngCertId"`
}

func (r *DeleteBgCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBgCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgLoadBalancersRequest struct {
	*tchttp.BaseRequest

	// CLB的VIP列表

	Vips []*string `json:"Vips,omitempty" name:"Vips"`
	// 租户的appId列表

	AppIds []*uint64 `json:"AppIds,omitempty" name:"AppIds"`
	// CLB的uLBId列表

	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
	// CLB的名称列表

	LoadBalancerNames []*string `json:"LoadBalancerNames,omitempty" name:"LoadBalancerNames"`
	// VpcId列表

	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`
	// 页码

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 页内个数

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// CLB实例状态

	Status []*int64 `json:"Status,omitempty" name:"Status"`
	// CLB实例类型；internal：内网，open：外网

	LoadBalancerType []*string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
	// VIP的运营商ID

	IspId []*int64 `json:"IspId,omitempty" name:"IspId"`
	// 实例规格；0：共享型，1：独占型

	Specs []*int64 `json:"Specs,omitempty" name:"Specs"`
	// 地址版本：ipv4、ipv6

	AddressIPVersion []*string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`
}

func (r *DescribeBgLoadBalancersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgLoadBalancersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeL7RuleExRequest struct {
	*tchttp.BaseRequest

	// 页偏移

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 页大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 通用查询参数，其值可以是:AppName, RsIp, Vip的值

	SearchCommon *string `json:"SearchCommon,omitempty" name:"SearchCommon"`
	// 服务ID

	VsIds *string `json:"VsIds,omitempty" name:"VsIds"`
	// 用户标识

	BizId *string `json:"BizId,omitempty" name:"BizId"`
}

func (r *DescribeL7RuleExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeL7RuleExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLbLimitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询配额返回描述

		LbLimitSet []*LbLimitSet `json:"LbLimitSet,omitempty" name:"LbLimitSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLbLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLbLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBgL7RuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetBgL7RuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBgL7RuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HbIpBusiIp struct {
	// 健康检查ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// ld物理ip

	LdIp *string `json:"LdIp,omitempty" name:"LdIp"`
	// 网段信息

	NetSegment *string `json:"NetSegment,omitempty" name:"NetSegment"`
	// 集群名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 掩码

	SubMask *string `json:"SubMask,omitempty" name:"SubMask"`
	// ip类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type Nat64SetList struct {
	// Busi Ip个数

	BusiLipCount *int64 `json:"BusiLipCount,omitempty" name:"BusiLipCount"`
	// Busi Ip 列表

	BusiLipList []*BusiLipList `json:"BusiLipList,omitempty" name:"BusiLipList"`
	// 容灾值

	DisasterRecovery *string `json:"DisasterRecovery,omitempty" name:"DisasterRecovery"`
	// 转发类型

	Fwd *string `json:"Fwd,omitempty" name:"Fwd"`
	// Hb Ip 个数

	HbLipCount *int64 `json:"HbLipCount,omitempty" name:"HbLipCount"`
	// 所属IDC

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 网络运营商

	Isp *string `json:"Isp,omitempty" name:"Isp"`
	// LD 列表

	LdList []*string `json:"LdList,omitempty" name:"LdList"`
	// Master 属性

	Master *Master `json:"Master,omitempty" name:"Master"`
	// Master 名称

	MasterName *string `json:"MasterName,omitempty" name:"MasterName"`
	// Module 名称

	Module *string `json:"Module,omitempty" name:"Module"`
	// 无用的TVS

	NouseTvs *int64 `json:"NouseTvs,omitempty" name:"NouseTvs"`
	// 系统类型

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 模式

	RouteMode *string `json:"RouteMode,omitempty" name:"RouteMode"`
	// 规则个数

	RuleCount *int64 `json:"RuleCount,omitempty" name:"RuleCount"`
	// 专区

	Section *string `json:"Section,omitempty" name:"Section"`
	// 专区列表

	SectionList []*string `json:"SectionList,omitempty" name:"SectionList"`
	// Set 名称

	Set *string `json:"Set,omitempty" name:"Set"`
	// Set ID

	SetId *int64 `json:"SetId,omitempty" name:"SetId"`
	// 特殊Set 类型

	SpecialSetType *string `json:"SpecialSetType,omitempty" name:"SpecialSetType"`
	// Tsv IP数

	TsvVipCount *int64 `json:"TsvVipCount,omitempty" name:"TsvVipCount"`
	// Tvs IP数

	TvsVipCount *int64 `json:"TvsVipCount,omitempty" name:"TvsVipCount"`
	// IP类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// Hb Ip 列表

	HbLipList []*HbLipList `json:"HbLipList,omitempty" name:"HbLipList"`
}

type ProcessApply struct {
	// 申请ID

	ApplyId *string `json:"ApplyId,omitempty" name:"ApplyId"`
	// 实施意见

	ProposeRemarks *string `json:"ProposeRemarks,omitempty" name:"ProposeRemarks"`
}

type ExportApdRuleL7Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回7层规则详情

		RulesL7Info *string `json:"RulesL7Info,omitempty" name:"RulesL7Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportApdRuleL7Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportApdRuleL7Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetApdIdcMasterRequest struct {
	*tchttp.BaseRequest

	// IDC名称

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// Master名称

	Master *string `json:"Master,omitempty" name:"Master"`
}

func (r *SetApdIdcMasterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetApdIdcMasterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBgL7RuleRequest struct {
	*tchttp.BaseRequest

	// 同步标识

	Sync *bool `json:"Sync,omitempty" name:"Sync"`
	// Set列表

	SetList []*string `json:"SetList,omitempty" name:"SetList"`
	// Service List

	ServiceList []*SetL7ServiceList `json:"ServiceList,omitempty" name:"ServiceList"`
}

func (r *SetBgL7RuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBgL7RuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QSetIdcList struct {
	// 地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// 城市

	City *string `json:"City,omitempty" name:"City"`
	// Master描述

	Master *Master `json:"Master,omitempty" name:"Master"`
	// Module列表

	ModuleList []*RspModuleList `json:"ModuleList,omitempty" name:"ModuleList"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 安全网关Master

	StgwMaster *StgwMaster `json:"StgwMaster,omitempty" name:"StgwMaster"`
	// 安全网关Set集群

	StgwSetList []*L7SetList `json:"StgwSetList,omitempty" name:"StgwSetList"`
}

type IspList struct {
	// 运营商ID

	IspId *int64 `json:"IspId,omitempty" name:"IspId"`
	// 运营商名称

	IspName *string `json:"IspName,omitempty" name:"IspName"`
	// 是否IPV4可见运营商

	Visible *uint64 `json:"Visible,omitempty" name:"Visible"`
	// 是否IPV6可见运营商

	VisibleIpv6 *uint64 `json:"VisibleIpv6,omitempty" name:"VisibleIpv6"`
	// 是否是默认ISP

	DefaultIsp *uint64 `json:"DefaultIsp,omitempty" name:"DefaultIsp"`
	// 是否是IPV6的默认ISP

	DefaultIspIpv6 *uint64 `json:"DefaultIspIpv6,omitempty" name:"DefaultIspIpv6"`
	// IPV4运营商别名

	V4Alias *string `json:"V4Alias,omitempty" name:"V4Alias"`
	// IPV6运营商别名

	V6Alias *string `json:"V6Alias,omitempty" name:"V6Alias"`
}

type DescribeApdRuleL7Request struct {
	*tchttp.BaseRequest

	// 查询页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 查询页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// service id

	VsIds *string `json:"VsIds,omitempty" name:"VsIds"`
	// 过滤参数，包括：AppName, BizId, Domain, Idc, SetName, Type, SearchCommon, SearchRs, Admin

	Filters []*FilterStr `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApdRuleL7Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdRuleL7Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgLocationL7Request struct {
	*tchttp.BaseRequest

	// 查询页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 查询页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// VS ID

	VsIds *string `json:"VsIds,omitempty" name:"VsIds"`
	// URL地址

	Url *string `json:"Url,omitempty" name:"Url"`
	// RS IP

	RsIp *string `json:"RsIp,omitempty" name:"RsIp"`
}

func (r *DescribeBgLocationL7Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgLocationL7Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgTargetHealthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例健康状态详情

		LoadBalancers []*LoadBalancerTargetInfo `json:"LoadBalancers,omitempty" name:"LoadBalancers"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBgTargetHealthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgTargetHealthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetireBgApplyRequest struct {
	*tchttp.BaseRequest

	// 申请ID列表

	ApplyId []*string `json:"ApplyId,omitempty" name:"ApplyId"`
}

func (r *RetireBgApplyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetireBgApplyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetireBgApplyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetireBgApplyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetireBgApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllNat64SetList struct {
	// 业务ip个数

	BusiLipCount *uint64 `json:"BusiLipCount,omitempty" name:"BusiLipCount"`
	// 业务ip列表

	BusiLipList []*BusiLipList `json:"BusiLipList,omitempty" name:"BusiLipList"`
	// 是否容灾

	DisasterRecovery *uint64 `json:"DisasterRecovery,omitempty" name:"DisasterRecovery"`
	// 转发类型

	Fwd *string `json:"Fwd,omitempty" name:"Fwd"`
	// 健康检查ip个数

	HbLipCount *uint64 `json:"HbLipCount,omitempty" name:"HbLipCount"`
	// 健康检查ip列表

	HbLipList []*HbLipList `json:"HbLipList,omitempty" name:"HbLipList"`
	// 园区

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 运营商

	Isp *string `json:"Isp,omitempty" name:"Isp"`
	// ld信息

	LdList []*string `json:"LdList,omitempty" name:"LdList"`
	// master信息

	Master *Master `json:"Master,omitempty" name:"Master"`
	// Master Name

	MasterName *string `json:"MasterName,omitempty" name:"MasterName"`
	// Module

	Module *string `json:"Module,omitempty" name:"Module"`
	// 未使用的tvsip数

	NouseTvs *uint64 `json:"NouseTvs,omitempty" name:"NouseTvs"`
	// 系统类型

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 部署模式

	RouteMode *string `json:"RouteMode,omitempty" name:"RouteMode"`
	// 规则数量

	RuleCount *uint64 `json:"RuleCount,omitempty" name:"RuleCount"`
	// Section

	Section *string `json:"Section,omitempty" name:"Section"`
	// SectionList

	SectionList []*string `json:"SectionList,omitempty" name:"SectionList"`
	// Set

	Set *string `json:"Set,omitempty" name:"Set"`
	// SetId

	SetId *uint64 `json:"SetId,omitempty" name:"SetId"`
	// SpecialSetType

	SpecialSetType *string `json:"SpecialSetType,omitempty" name:"SpecialSetType"`
	// TsvVipCount

	TsvVipCount *uint64 `json:"TsvVipCount,omitempty" name:"TsvVipCount"`
	// TvsVipCount

	TvsVipCount *uint64 `json:"TvsVipCount,omitempty" name:"TvsVipCount"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// 运营商别名

	IspAlias *string `json:"IspAlias,omitempty" name:"IspAlias"`
}

type SessionPersistenceStrategy struct {
	// cookie key

	CookieKey *string `json:"CookieKey,omitempty" name:"CookieKey"`
	// 超时阈值

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	// 类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type CreateLdLocalIpRequest struct {
	*tchttp.BaseRequest

	// Set名称描述

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// IP 列表

	LipList []*LipList `json:"LipList,omitempty" name:"LipList"`
	// 操作类型

	Op *string `json:"Op,omitempty" name:"Op"`
	// 网关IP

	LdIp *string `json:"LdIp,omitempty" name:"LdIp"`
}

func (r *CreateLdLocalIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLdLocalIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleList struct {
	// 规则Id

	BizId *string `json:"BizId,omitempty" name:"BizId"`
	// 规则名称

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 协义类型

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// Tsv 地址

	Tsv *string `json:"Tsv,omitempty" name:"Tsv"`
	// 服务器Ip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// vpc网络ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 监听端口

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
}

type DescribeApdMasterListRequest struct {
	*tchttp.BaseRequest

	// 查询页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 查询页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeApdMasterListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdMasterListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpApdModuleRequest struct {
	*tchttp.BaseRequest

	// 需要操作的Module列表

	ModuleList []*ModuleList `json:"ModuleList,omitempty" name:"ModuleList"`
	// 操作类型

	OpType *string `json:"OpType,omitempty" name:"OpType"`
}

func (r *OpApdModuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpApdModuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HealthCheck struct {
	// 主动检查属性

	ActiveHealthCheck *ActiveHealthCheck `json:"ActiveHealthCheck,omitempty" name:"ActiveHealthCheck"`
	// 被动检查属性

	PassiveHealthCheck *PassiveHealthCheck `json:"PassiveHealthCheck,omitempty" name:"PassiveHealthCheck"`
	// 检查类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type CreateApdLdL7Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApdLdL7Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApdLdL7Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdRuleRequest struct {
	*tchttp.BaseRequest

	// 查询页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 查询页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// RS IP列表

	RsIpList []*string `json:"RsIpList,omitempty" name:"RsIpList"`
	// 规则接口

	RuleList []*FRuleList `json:"RuleList,omitempty" name:"RuleList"`
	// 过滤参数，包括:BizId, Idc, AppName, SetName

	Filters []*FilterStr `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApdRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitIspListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitIspListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitIspListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpBgL7RsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Requestid

		Requestid *string `json:"Requestid,omitempty" name:"Requestid"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpBgL7RsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpBgL7RsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBgRuleRequest struct {
	*tchttp.BaseRequest

	// 同步标识

	Sync *bool `json:"Sync,omitempty" name:"Sync"`
	// Set列表

	SetList []*string `json:"SetList,omitempty" name:"SetList"`
	// 规则列表

	RuleList []*OpRuleList `json:"RuleList,omitempty" name:"RuleList"`
}

func (r *SetBgRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBgRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdIdcSetRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeApdIdcSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdIdcSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLdLocalIpRequest struct {
	*tchttp.BaseRequest

	// Set集群名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// LD地址

	LdIp *string `json:"LdIp,omitempty" name:"LdIp"`
	// Local IP

	LocalIp *string `json:"LocalIp,omitempty" name:"LocalIp"`
	// LD状态列表

	Status []*string `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeLdLocalIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLdLocalIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportApdRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出规则信息

		RulesInfo *string `json:"RulesInfo,omitempty" name:"RulesInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportApdRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportApdRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportApdRulePubCertRequest struct {
	*tchttp.BaseRequest

	// 证书ID

	SngCertId *string `json:"SngCertId,omitempty" name:"SngCertId"`
}

func (r *ExportApdRulePubCertRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportApdRulePubCertRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpApdSectionRequest struct {
	*tchttp.BaseRequest

	// Section列表

	SectionList []*OpSectionList `json:"SectionList,omitempty" name:"SectionList"`
	// 操作类型

	OpType *string `json:"OpType,omitempty" name:"OpType"`
}

func (r *OpApdSectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpApdSectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpApdStgwSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpApdStgwSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpApdStgwSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssignApdLdToStgwSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssignApdLdToStgwSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssignApdLdToStgwSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApdMasterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApdMasterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApdMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdAppNameRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeApdAppNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdAppNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回证书个数

		CertAllCount *uint64 `json:"CertAllCount,omitempty" name:"CertAllCount"`
		// 证书列表

		CertList []*CertList `json:"CertList,omitempty" name:"CertList"`
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Requestid

		Requestid *string `json:"Requestid,omitempty" name:"Requestid"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApdCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppIdLabelRequest struct {
	*tchttp.BaseRequest

	// 集群标签

	Label *string `json:"Label,omitempty" name:"Label"`
	// AppId/Uin

	Owners []*string `json:"Owners,omitempty" name:"Owners"`
}

func (r *DescribeAppIdLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppIdLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApdLdL7Request struct {
	*tchttp.BaseRequest

	// LD详情列表

	LdList []*LdList `json:"LdList,omitempty" name:"LdList"`
}

func (r *CreateApdLdL7Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApdLdL7Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgApplyRequest struct {
	*tchttp.BaseRequest

	// 申请时间范围起点

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 申请时间范围终点

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询页，起点1

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 查询页大小，默认20

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 登录账户

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 过滤条件：ApplyIdAppName, Flag

	Filters []*FilterStr `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeBgApplyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgApplyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpApdSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpApdSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpApdSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdIdcSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的Idc列表

		IdcList []*IdcList `json:"IdcList,omitempty" name:"IdcList"`
		// 所有四层集群

		AllL4SetList []*L4SetList `json:"AllL4SetList,omitempty" name:"AllL4SetList"`
		// 所有七层集群

		AllL7SetList []*L7SetList `json:"AllL7SetList,omitempty" name:"AllL7SetList"`
		// 集群标签信息

		GroupTagList []*string `json:"GroupTagList,omitempty" name:"GroupTagList"`
		// section与idc的对应关系

		Section2Idc []*Section2Idc `json:"Section2Idc,omitempty" name:"Section2Idc"`
		// 集群列表

		SetList []*L4SetList `json:"SetList,omitempty" name:"SetList"`
		// 所有nat64集群信息

		AllNat64SetList []*AllNat64SetList `json:"AllNat64SetList,omitempty" name:"AllNat64SetList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApdIdcSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdIdcSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelSetName struct {
	// 标签名

	Label *string `json:"Label,omitempty" name:"Label"`
	// 集群名

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 集群类型，支持L4_LAN_CLB, L7_CLB, L7_LAN_CLB, L7_WAN_CLB

	SetType *string `json:"SetType,omitempty" name:"SetType"`
}

type ReqList struct {
	// Eip 地址

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// vpc网络编号

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 服务器IP

	RsIp *string `json:"RsIp,omitempty" name:"RsIp"`
	// 主机Ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
}

type CreateApdLdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApdLdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApdLdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApdVipGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApdVipGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApdVipGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdLdL7Request struct {
	*tchttp.BaseRequest

	// 页偏移

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// section列表

	SectionList []*string `json:"SectionList,omitempty" name:"SectionList"`
	// 过滤参数，包括:Idc, Module,SetName, SvrId, LanIp, WanIp, Status, Type

	Filters []*FilterStr `json:"Filters,omitempty" name:"Filters"`
	// 转发类型列表

	FwdMode []*string `json:"FwdMode,omitempty" name:"FwdMode"`
}

func (r *DescribeApdLdL7Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdLdL7Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LipList struct {
	// 本地IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 子网掩码

	SubMask *string `json:"SubMask,omitempty" name:"SubMask"`
	// 网段划分

	NetSegment *string `json:"NetSegment,omitempty" name:"NetSegment"`
}

type RsList struct {
	// 服务端 IP

	RsIp *string `json:"RsIp,omitempty" name:"RsIp"`
	// 服务端端口

	RsPort *int64 `json:"RsPort,omitempty" name:"RsPort"`
	// 服务端负载权重,op_type为add 时需要传该参数

	RsWeight *int64 `json:"RsWeight,omitempty" name:"RsWeight"`
	// vpc 编号

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 服务器类型

	ServerType *string `json:"ServerType,omitempty" name:"ServerType"`
	// 主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// ipv6 rs需要传此参数，支撑不需要

	V6Id *string `json:"V6Id,omitempty" name:"V6Id"`
	// set时需传该参数

	OldRsIp *string `json:"OldRsIp,omitempty" name:"OldRsIp"`
	// set时需传该参数

	OldRsPort *int64 `json:"OldRsPort,omitempty" name:"OldRsPort"`
	// set时需传该参数

	OldVpcid *int64 `json:"OldVpcid,omitempty" name:"OldVpcid"`
	// set时需传该参数

	OldHostIp *string `json:"OldHostIp,omitempty" name:"OldHostIp"`
	// rs绑定的vip

	Ipv4Vip *string `json:"Ipv4Vip,omitempty" name:"Ipv4Vip"`
	// rs母机ip

	Rshost *string `json:"Rshost,omitempty" name:"Rshost"`
	// 虚拟网络id

	Rsvpcid *int64 `json:"Rsvpcid,omitempty" name:"Rsvpcid"`
}

type ZoneInfo struct {
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type OwnerLabel struct {
	// AppId/Uin

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 集群标签

	Label *string `json:"Label,omitempty" name:"Label"`
}

type AssignApdLdToSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssignApdLdToSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssignApdLdToSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindSetL4SetAndL7Request struct {
	*tchttp.BaseRequest

	// Set集群名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 操作类型

	Op *string `json:"Op,omitempty" name:"Op"`
	// 七层SetID列表

	BindL7SetId []*string `json:"BindL7SetId,omitempty" name:"BindL7SetId"`
}

func (r *BindSetL4SetAndL7Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindSetL4SetAndL7Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BusiLipList struct {
	// IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// LD IP

	LdIp *string `json:"LdIp,omitempty" name:"LdIp"`
	// 网段划分

	NetSegment *string `json:"NetSegment,omitempty" name:"NetSegment"`
	// Set名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 子网掩码

	SubMask *string `json:"SubMask,omitempty" name:"SubMask"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type L7SetList struct {
	// 容灾值

	DisasterRecovery *int64 `json:"DisasterRecovery,omitempty" name:"DisasterRecovery"`
	// 转发类型

	Fwd *string `json:"Fwd,omitempty" name:"Fwd"`
	// 所属IDC

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 运营商名称

	Isp *string `json:"Isp,omitempty" name:"Isp"`
	// LD列表

	LdList []*string `json:"LdList,omitempty" name:"LdList"`
	// Module 名称

	Module *string `json:"Module,omitempty" name:"Module"`
	// 无用的TVS

	NouseTvs *int64 `json:"NouseTvs,omitempty" name:"NouseTvs"`
	// 系统类型

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 模式

	RouteMode *string `json:"RouteMode,omitempty" name:"RouteMode"`
	// 规则个数

	RuleCount *int64 `json:"RuleCount,omitempty" name:"RuleCount"`
	// 专区

	Section *string `json:"Section,omitempty" name:"Section"`
	// 专区列表

	SectionList []*string `json:"SectionList,omitempty" name:"SectionList"`
	// Set名称

	Set *string `json:"Set,omitempty" name:"Set"`
	// Set ID

	SetId *int64 `json:"SetId,omitempty" name:"SetId"`
	// 特殊Set 类型

	SpecialSetType *string `json:"SpecialSetType,omitempty" name:"SpecialSetType"`
	// Tsv IP数

	TsvVipCount *int64 `json:"TsvVipCount,omitempty" name:"TsvVipCount"`
	// Tvs IP数

	TvsVipCount *int64 `json:"TvsVipCount,omitempty" name:"TvsVipCount"`
	// IP类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 运营商别名

	IspAlias *string `json:"IspAlias,omitempty" name:"IspAlias"`
}

type MasterList struct {
	// Master环境地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 环境名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 服务端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 模式

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 主备模式

	MasterStatus *string `json:"MasterStatus,omitempty" name:"MasterStatus"`
}

type ClusterState struct {
	// CLB产品下的网关集群名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 集群

	Status *string `json:"Status,omitempty" name:"Status"`
	// 主节点所在AZ

	MainNodeAzList []*string `json:"MainNodeAzList,omitempty" name:"MainNodeAzList"`
	// 节点列表

	NodeList []*TgwNode `json:"NodeList,omitempty" name:"NodeList"`
}

type CreateApdLdRequest struct {
	*tchttp.BaseRequest

	// Ld列表

	LdList []*LdList `json:"LdList,omitempty" name:"LdList"`
}

func (r *CreateApdLdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApdLdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApdVipSegmentRequest struct {
	*tchttp.BaseRequest

	// 操作类型

	OpType *string `json:"OpType,omitempty" name:"OpType"`
	// 网段信息

	Segments []*Segments `json:"Segments,omitempty" name:"Segments"`
	// Set集群

	Set *string `json:"Set,omitempty" name:"Set"`
}

func (r *DeleteApdVipSegmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApdVipSegmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdLdRequest struct {
	*tchttp.BaseRequest

	// 查询页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 查询页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// Section数组过滤

	SectionList []*string `json:"SectionList,omitempty" name:"SectionList"`
	// 过滤参数，包括:Idc, Module,SetName, SvrId, LanIp, WanIp, Status, Type

	Filters []*FilterStr `json:"Filters,omitempty" name:"Filters"`
	// 转发类型列表

	FwdMode []*string `json:"FwdMode,omitempty" name:"FwdMode"`
}

func (r *DescribeApdLdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdLdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgRsByRuleRequest struct {
	*tchttp.BaseRequest

	// 查询页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 查询页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 规则列表

	RuleList []*FRuleList `json:"RuleList,omitempty" name:"RuleList"`
}

func (r *DescribeBgRsByRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgRsByRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgServersRequest struct {
	*tchttp.BaseRequest

	// 查询页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 查询页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 查询IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *DescribeBgServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgServersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ForwardStrategyOut struct {
	// 转发策略类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 转发策略详情

	Detail *string `json:"Detail,omitempty" name:"Detail"`
}

type SetList7 struct {
	// Set集群名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 网络类型

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 是否运营专属集群

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 集群模式

	RouteMode *string `json:"RouteMode,omitempty" name:"RouteMode"`
	// 所属IDC

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
}

type UploadBgCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 证书名称

		CertName *string `json:"CertName,omitempty" name:"CertName"`
		// 证书ID

		SngCertId *string `json:"SngCertId,omitempty" name:"SngCertId"`
		// Requestid

		Requestid *string `json:"Requestid,omitempty" name:"Requestid"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadBgCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadBgCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindSetLabelRequest struct {
	*tchttp.BaseRequest

	// 集群标签名

	Label *string `json:"Label,omitempty" name:"Label"`
	// 集群名

	SetNames []*string `json:"SetNames,omitempty" name:"SetNames"`
	// 修改前标签名，如果存在且有值，说明是修改标签名

	OldLabel *string `json:"OldLabel,omitempty" name:"OldLabel"`
	// 集群类型，支持L4_LAN_CLB, L7_CLB

	SetType *string `json:"SetType,omitempty" name:"SetType"`
}

func (r *BindSetLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindSetLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApdLdFromStgwSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApdLdFromStgwSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApdLdFromStgwSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllIspListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAllIspListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllIspListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelLdList struct {
	// Ld IP

	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`
	// 服务器

	SvrId *string `json:"SvrId,omitempty" name:"SvrId"`
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

type OpApdIdcRequest struct {
	*tchttp.BaseRequest

	// 城市

	City *string `json:"City,omitempty" name:"City"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// 区域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 二级区域

	SubZoneName *string `json:"SubZoneName,omitempty" name:"SubZoneName"`
	// 操作类型

	OpType *string `json:"OpType,omitempty" name:"OpType"`
}

func (r *OpApdIdcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpApdIdcRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetApdIdcStgwMasterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetApdIdcStgwMasterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetApdIdcStgwMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModuleList struct {
	// Module 名称

	Module *string `json:"Module,omitempty" name:"Module"`
	// 所属Idc

	Idc *string `json:"Idc,omitempty" name:"Idc"`
}

type MultiEth0 struct {
	// Ip地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 运营商描述

	Isp *string `json:"Isp,omitempty" name:"Isp"`
}

type VipGroupList struct {
	// 用户标识

	BizId *string `json:"BizId,omitempty" name:"BizId"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 转发类型

	Fwd *string `json:"Fwd,omitempty" name:"Fwd"`
	// 标签

	GroupTag *string `json:"GroupTag,omitempty" name:"GroupTag"`
	// 所属idc

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 规则数

	RuleCount *uint64 `json:"RuleCount,omitempty" name:"RuleCount"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 用途

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// 使用状态

	UseStatus *string `json:"UseStatus,omitempty" name:"UseStatus"`
	// 组id

	VipGroupId *string `json:"VipGroupId,omitempty" name:"VipGroupId"`
	// 集群列表

	SetNameList []*string `json:"SetNameList,omitempty" name:"SetNameList"`
	// Vip信息

	VipList []*VipList `json:"VipList,omitempty" name:"VipList"`
	// tvs列表

	TvsList []*string `json:"TvsList,omitempty" name:"TvsList"`
	// tsv列表

	TsvList []*string `json:"TsvList,omitempty" name:"TsvList"`
}

type SetApdLdStatusRequest struct {
	*tchttp.BaseRequest

	// 启用的LD地址列表

	LdList []*string `json:"LdList,omitempty" name:"LdList"`
	// 启用的LD名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 操作动作

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// ip个数

	LocalIpNum *int64 `json:"LocalIpNum,omitempty" name:"LocalIpNum"`
}

func (r *SetApdLdStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetApdLdStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FRuleList struct {
	// IP地址

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 端口

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// 协议类型

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// VpcID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
}

type LoadBalancerInfo struct {
	// uLBId

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 名称

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 所属用户appId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 网络类型

	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
	// 所属vpc

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 网络属性

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 运营商id

	VipIspId *int64 `json:"VipIspId,omitempty" name:"VipIspId"`
	// vip所属版本

	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`
	// 主可用区

	MasterZone *ZoneInfo `json:"MasterZone,omitempty" name:"MasterZone"`
	// 备可用区

	BackupZoneSet []*ZoneInfo `json:"BackupZoneSet,omitempty" name:"BackupZoneSet"`
	// 安全组信息

	SgIdList []*string `json:"SgIdList,omitempty" name:"SgIdList"`
	// 四层集群标签

	TgwSetLabel *string `json:"TgwSetLabel,omitempty" name:"TgwSetLabel"`
	// 七层集群标签

	StgwSetLabel *string `json:"StgwSetLabel,omitempty" name:"StgwSetLabel"`
	// 绑定的个性化配置ID

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 绑定的个性化配置名称

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
	// 实例创建时间

	AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`
}

type TgwNode struct {
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

type DescribeLdLocalIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// IpList

		IpList []*IpList `json:"IpList,omitempty" name:"IpList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLdLocalIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLdLocalIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LbLimit struct {
	// 租户id

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 4代表公网实例个数，7代表私有网实例个数，8代表单个实例监听的个数，9代表单个实例后端服务器的个数

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 当前配额数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 最大配额数量，必须大于等于Limit

	MaxLimit *int64 `json:"MaxLimit,omitempty" name:"MaxLimit"`
}

type IpList struct {
	// Ip地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// LD Ip

	LdIp *string `json:"LdIp,omitempty" name:"LdIp"`
	// 网段信息

	NetSegment *string `json:"NetSegment,omitempty" name:"NetSegment"`
	// 名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 子网掩码

	SubMask *string `json:"SubMask,omitempty" name:"SubMask"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type LoadBalancerTargetInfo struct {
	// 租户CLB实例ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 租户CLB实例名称

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 监听器

	Listeners []*LBListenerTargetInfo `json:"Listeners,omitempty" name:"Listeners"`
	// CLB实例类型，internal 内网，open 公网

	LBType *string `json:"LBType,omitempty" name:"LBType"`
	// CLB实例VIP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// Vpcid

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
}

type ApplyBgRuleRequest struct {
	*tchttp.BaseRequest

	// 业务名称

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 账户类型

	BizType *string `json:"BizType,omitempty" name:"BizType"`
	// 规则类型

	FwdMode *string `json:"FwdMode,omitempty" name:"FwdMode"`
	// 运营商名称

	IspName *string `json:"IspName,omitempty" name:"IspName"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 是否容灾,取值0， 1

	DisasterRecovery *int64 `json:"DisasterRecovery,omitempty" name:"DisasterRecovery"`
	// vpc id

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 子网id

	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`
	// IDC

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 用途

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// 端口号

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 标记

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是否启用ssl, 取值0，1

	SslVerifyClientEnable *int64 `json:"SslVerifyClientEnable,omitempty" name:"SslVerifyClientEnable"`
	// 城市

	City *string `json:"City,omitempty" name:"City"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 证书ID

	SngCertId *string `json:"SngCertId,omitempty" name:"SngCertId"`
	// 证书名称

	SngCertName *string `json:"SngCertName,omitempty" name:"SngCertName"`
	// CA证书ID

	SngCaCertId *string `json:"SngCaCertId,omitempty" name:"SngCaCertId"`
	// CA证书名称

	SngCaCertName *string `json:"SngCaCertName,omitempty" name:"SngCaCertName"`
	// 四层集群标签列表

	SetLabels []*string `json:"SetLabels,omitempty" name:"SetLabels"`
	// 七层集群标签列表

	StgwSetLabels []*string `json:"StgwSetLabels,omitempty" name:"StgwSetLabels"`
	// 申请人名称

	Uid *string `json:"Uid,omitempty" name:"Uid"`
}

func (r *ApplyBgRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyBgRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Segments struct {
	// vip类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 运营商名称

	IspName *string `json:"IspName,omitempty" name:"IspName"`
	// IP地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 网段数量

	BitCount *uint64 `json:"BitCount,omitempty" name:"BitCount"`
}

type DescribeApdVipGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// VIP组详情

		VipGroupList []*VipGroupList `json:"VipGroupList,omitempty" name:"VipGroupList"`
		// VIP组总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApdVipGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdVipGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgApplyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 申请记录

		Apply []*Apply `json:"Apply,omitempty" name:"Apply"`
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBgApplyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLbLimitRequest struct {
	*tchttp.BaseRequest

	// 修改配额的类型，4代表公网实例个数，7代表私有网实例个数，8代表单个实例监听的个数，9代表单个实例后端服务器的个数

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 当前配额限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 最大配额

	MaxLimit *int64 `json:"MaxLimit,omitempty" name:"MaxLimit"`
	// 管理员

	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

func (r *ModifyLbLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLbLimitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductHealthStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CLB集群健康状态详情

		Data []*ProductHealthStateRsp `json:"Data,omitempty" name:"Data"`
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

type LBListenerTargetInfo struct {
	// 监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 监听器名称

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
	// 监听器协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 转发规则

	Rules []*LBLocationTargetInfo `json:"Rules,omitempty" name:"Rules"`
}

type TargetHealthInfo struct {
	// 实例ID

	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`
	// 实例IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 健康状态，-1未知，0不健康，1健康

	HealthStatus *int64 `json:"HealthStatus,omitempty" name:"HealthStatus"`
}

type DescribeBgLoadBalancersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 当前个数

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// CLB实例列表

		LoadBalancerSet []*LoadBalancerInfo `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBgLoadBalancersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Vips struct {
	// 申请编号

	Applyid *string `json:"Applyid,omitempty" name:"Applyid"`
	// 所属IDC

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 所属SET集群

	Set *string `json:"Set,omitempty" name:"Set"`
	// TVS地址

	Tvs *string `json:"Tvs,omitempty" name:"Tvs"`
	// 运营商名称

	IspName *string `json:"IspName,omitempty" name:"IspName"`
	// TSV地址

	Tsv *string `json:"Tsv,omitempty" name:"Tsv"`
	// 可用值

	Avaliable *int64 `json:"Avaliable,omitempty" name:"Avaliable"`
}

type DelVipGroupList struct {
	// 被删除的VIP列表

	VipList []*DelVipList `json:"VipList,omitempty" name:"VipList"`
}

type DeleteApdVipRequest struct {
	*tchttp.BaseRequest

	// VIP类型

	VipType *string `json:"VipType,omitempty" name:"VipType"`
	// 需删除的VIP列表

	VipList []*string `json:"VipList,omitempty" name:"VipList"`
	// 私有网络ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DeleteApdVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApdVipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdVipRequest struct {
	*tchttp.BaseRequest

	// 查询页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 查询页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 区分是否NAT64

	VipType *string `json:"VipType,omitempty" name:"VipType"`
	// 过滤器,支持：IpType - string IP类型； Idc -string所属IDC； Status - string 状态 use, nouse； Isp - string 网络运营商；IpList -  string 数组 IP列表；SetName-string, 集群名称

	Filters []*FilterStr `json:"Filters,omitempty" name:"Filters"`
	// IP列表

	IpList []*string `json:"IpList,omitempty" name:"IpList"`
	// 是否容灾集群

	DisasterRecovery *bool `json:"DisasterRecovery,omitempty" name:"DisasterRecovery"`
}

func (r *DescribeApdVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdVipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgRsL7Request struct {
	*tchttp.BaseRequest

	// 查询页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 查询页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 主域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// Url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 端口

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// Vpc ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// VipLlist过滤条件

	VipList []*string `json:"VipList,omitempty" name:"VipList"`
}

func (r *DescribeBgRsL7Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgRsL7Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLbLimitRequest struct {
	*tchttp.BaseRequest

	// 管理着

	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

func (r *DescribeLbLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLbLimitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBgCertListRequest struct {
	*tchttp.BaseRequest

	// 导出页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 导出页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ExportBgCertListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBgCertListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApdLdFromSetRequest struct {
	*tchttp.BaseRequest

	// Set集群名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// LD 列表

	LdList []*LdListToSet `json:"LdList,omitempty" name:"LdList"`
}

func (r *DeleteApdLdFromSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApdLdFromSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpBgRsRequest struct {
	*tchttp.BaseRequest

	// 用途

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// 默认为true，true表示同步返回errno为0则表示变更完成； false表示异步操作，返回成功不表示变更成功，需要进一步检查任务状态

	Sync *bool `json:"Sync,omitempty" name:"Sync"`
	// TGW实现方式,默认为空，由TGW来自动选择，有如下方式： linux_tunl：在RS（物理机）上安装隧道、下发NAT配置（取决于apply_rule时的need_nat选项） ,最常见的方式，缺点依赖RS环境     piaoyi_ko：qcloud ip漂移专用，在母机上安装IP漂移的KO模块，在母机下发NAT配置     fuzai_ko：虚拟机， 不需要安装隧道，在母机上安装负载均衡的KO（母机需要提前安装），在母机上下发NAT配置（取决于apply_rule时的need_nat选项），SNG开放平台windows虚拟子机常用，linux子机也在规划中    win_tunl：在RS（windows物理机）上安装隧道、下发NAT配置（取决于apply_rule时的need_nat选项），目前只支持IEG     cdb ： CDB专用，不会去母机、子机做任何事情（由cdb自己下发）*/

	RsType *string `json:"RsType,omitempty" name:"RsType"`
	// 操作类型,add:绑定RS，del:解绑rs, set:更换rs（set只支持overlay rs）

	OpType *string `json:"OpType,omitempty" name:"OpType"`
	// 规则详情

	RuleList []*OpRuleList `json:"RuleList,omitempty" name:"RuleList"`
	// 可选，用于CAE使用CDB set 下的vip上线RS，其他请勿填写该字段。默认false不安装

	NeedSetupRs *bool `json:"NeedSetupRs,omitempty" name:"NeedSetupRs"`
	// ipv4|ipv6，默认为ipv4，CLB如果要申请ipv6 vip需明确指定为ipv6，支撑环境ipv6不需要传

	AfType *string `json:"AfType,omitempty" name:"AfType"`
	// 1:sdn_rs，不传默认0

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
	// 可选，默认false，op_type==del时有效；为true时,当规则下全部rs被删除后，删除该规则，否则保留规则；

	ClearRule *bool `json:"ClearRule,omitempty" name:"ClearRule"`
	// 可选，默认false，op_type==del时有效；为true时,当规则下全部rs被删除后，且clear_rule为true时，将7层规则域名指向的vip修改为1.1.1.1

	DebindGslb *bool `json:"DebindGslb,omitempty" name:"DebindGslb"`
}

func (r *OpBgRsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpBgRsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type L7Slave struct {
	// 城市

	City *string `json:"City,omitempty" name:"City"`
	// 连续时间

	ContinuousTimes *string `json:"ContinuousTimes,omitempty" name:"ContinuousTimes"`
	// 心跳

	Heartbeat *string `json:"Heartbeat,omitempty" name:"Heartbeat"`
	// IDC名称

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 端口

	MasterPort *string `json:"MasterPort,omitempty" name:"MasterPort"`
}

type VipOutByApply struct {
	// tvs vip

	Tvs *string `json:"Tvs,omitempty" name:"Tvs"`
	// tsv vip

	Tsv *string `json:"Tsv,omitempty" name:"Tsv"`
	// vpcid

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
}

type SetApdIdcStgwMasterRequest struct {
	*tchttp.BaseRequest

	// IDC名称

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// Master名称

	Master *string `json:"Master,omitempty" name:"Master"`
}

func (r *SetApdIdcStgwMasterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetApdIdcStgwMasterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApdVipGroupRequest struct {
	*tchttp.BaseRequest

	// Vip组名称

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 用途

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// VIP组配置详情

	VipGroupList *VipGroupList `json:"VipGroupList,omitempty" name:"VipGroupList"`
}

func (r *CreateApdVipGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApdVipGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdApplyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 申请列表

		Apply []*Apply `json:"Apply,omitempty" name:"Apply"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApdApplyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdMasterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Master列表

		MasterList []*Master `json:"MasterList,omitempty" name:"MasterList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApdMasterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgRuleL7Request struct {
	*tchttp.BaseRequest

	// 查询页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 查询页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// service id 数组

	VsIds *string `json:"VsIds,omitempty" name:"VsIds"`
	// 过滤参数，包括AppName,SearchCommon, SearchRs, Idc,Admin,  BizId, Type

	Filters []*FilterStr `json:"Filters,omitempty" name:"Filters"`
	// 规则列表

	RuleList []*FRuleList `json:"RuleList,omitempty" name:"RuleList"`
}

func (r *DescribeBgRuleL7Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgRuleL7Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeL4RuleExRequest struct {
	*tchttp.BaseRequest

	// 查询页码

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 每页的大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 详细的过滤条件如下：<li> SearchVip- String - 是否必填：否 - （过滤条件）clb vip。</li><li> BizId - String - 是否必填 ：否 - （过滤条件）用户ID。</li><li> SearchRs - String - 是否必填：否 - （过滤条件）绑定RSIP。</li><li> SearchAppName - String - 是否必填：否 - （过滤条 件）业务名称。</li>

	Filters []*FilterStr `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeL4RuleExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeL4RuleExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpBgL7RsRequest struct {
	*tchttp.BaseRequest

	// 操作类型

	OpType *string `json:"OpType,omitempty" name:"OpType"`
	// 同步标识

	Sync *bool `json:"Sync,omitempty" name:"Sync"`
	// Service列表

	ServiceList []*ServiceList `json:"ServiceList,omitempty" name:"ServiceList"`
}

func (r *OpBgL7RsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpBgL7RsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RspLdList struct {
	// 城市

	City *string `json:"City,omitempty" name:"City"`
	// IDC名称

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 内网地址

	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`
	// Module名称

	Module *string `json:"Module,omitempty" name:"Module"`
	// 机架

	Rack *string `json:"Rack,omitempty" name:"Rack"`
	// Section名称

	Section *string `json:"Section,omitempty" name:"Section"`
	// 所属集群

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 服务器ID

	SvrId *string `json:"SvrId,omitempty" name:"SvrId"`
	// 服务器类型

	SvrType *string `json:"SvrType,omitempty" name:"SvrType"`
	// 绑定规则类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 外网地址

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// TvsVip数量

	TvsVipCount *uint64 `json:"TvsVipCount,omitempty" name:"TvsVipCount"`
	// 规则数量

	RuleCount *uint64 `json:"RuleCount,omitempty" name:"RuleCount"`
	// 健康检查ip个数

	HbLipCount *uint64 `json:"HbLipCount,omitempty" name:"HbLipCount"`
	// 健康检查ip列表

	HbLipList []*HbLipList `json:"HbLipList,omitempty" name:"HbLipList"`
	// 业务ip个数

	BusiLipCount *uint64 `json:"BusiLipCount,omitempty" name:"BusiLipCount"`
	// 业务ip列表

	BusiLipList []*BusiLipList `json:"BusiLipList,omitempty" name:"BusiLipList"`
}

type DescribeBgRuleL7Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回规则列表

		Rules []*Rules `json:"Rules,omitempty" name:"Rules"`
		// 规则数量

		RulesCount *uint64 `json:"RulesCount,omitempty" name:"RulesCount"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Requestid

		Requestid *string `json:"Requestid,omitempty" name:"Requestid"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBgRuleL7Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgRuleL7Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReWrite struct {
	// 操作类型

	Op *string `json:"Op,omitempty" name:"Op"`
	// 协议类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// Url

	Url *string `json:"Url,omitempty" name:"Url"`
	// V2Arg

	V2Arg *ReWriteV2Arg `json:"V2Arg,omitempty" name:"V2Arg"`
}

type Error struct {
	// 错误编码

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 错误描述

	Message *string `json:"Message,omitempty" name:"Message"`
}

type DescribeApdVipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// VIP列表

		VipList []*VipList `json:"VipList,omitempty" name:"VipList"`
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApdVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 证书总数

		CertAllCount *uint64 `json:"CertAllCount,omitempty" name:"CertAllCount"`
		// 证书列表

		CertList []*CertList `json:"CertList,omitempty" name:"CertList"`
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBgCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProcessApdApplyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实施规则任务ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 实施的规则对应的VIP信息

		Vip *VipOutByApply `json:"Vip,omitempty" name:"Vip"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProcessApdApplyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProcessApdApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApdMasterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApdMasterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApdMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdApplyRequest struct {
	*tchttp.BaseRequest

	// 查询页号

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 查询页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 过滤条件：AppName, Applyer, ApplyId, BeginTime, EndTime, Flag, Idc, Protocol,

	Filters []*FilterStr `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApdApplyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdApplyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppIdLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// AppId和集群标签列表

		OwnerLabelSet []*OwnerLabel `json:"OwnerLabelSet,omitempty" name:"OwnerLabelSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppIdLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppIdLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeL4RuleExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则列表

		RuleList []*FRuleList `json:"RuleList,omitempty" name:"RuleList"`
		// 规则总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 是否所有业务

		AllAppShown *bool `json:"AllAppShown,omitempty" name:"AllAppShown"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeL4RuleExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeL4RuleExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApdMasterRequest struct {
	*tchttp.BaseRequest

	// 区域

	City *string `json:"City,omitempty" name:"City"`
	// 模式

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 备用IP

	SlaveIp *string `json:"SlaveIp,omitempty" name:"SlaveIp"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
}

func (r *ModifyApdMasterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApdMasterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewCertInfo struct {
	// 签名证书ID

	SngCertId *string `json:"SngCertId,omitempty" name:"SngCertId"`
}

type Rules struct {
	// 本地集群列表

	LocationSet []*LocationSet `json:"LocationSet,omitempty" name:"LocationSet"`
	// 服务器属性

	VirtualService *VirtualService `json:"VirtualService,omitempty" name:"VirtualService"`
	// 本地所有数量

	LocationAllTotal *uint64 `json:"LocationAllTotal,omitempty" name:"LocationAllTotal"`
	// 本地数量

	LocationTotal *uint64 `json:"LocationTotal,omitempty" name:"LocationTotal"`
}

type Section2Idc struct {
	// Section名称

	Section *string `json:"Section,omitempty" name:"Section"`
	// IDC名称

	Idc *string `json:"Idc,omitempty" name:"Idc"`
}

type SetL7ServiceList struct {
	// 业务ID

	BizId *string `json:"BizId,omitempty" name:"BizId"`
	// Vpc ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// VIP 列表

	VipList []*string `json:"VipList,omitempty" name:"VipList"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 端口

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// Location List

	LocationList []*SetL7LocationList `json:"LocationList,omitempty" name:"LocationList"`
}

type CreateApdVipGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApdVipGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApdVipGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadBgCertificateRequest struct {
	*tchttp.BaseRequest

	// 业务名称

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 证书名称

	CertName *string `json:"CertName,omitempty" name:"CertName"`
	// 国际证书私钥或者国密证书证书私钥

	CertPrivateKey *string `json:"CertPrivateKey,omitempty" name:"CertPrivateKey"`
	// 国际证书公钥或者国密证书证书公钥

	CertPublicCert *string `json:"CertPublicCert,omitempty" name:"CertPublicCert"`
	// 证书类型

	CertType *string `json:"CertType,omitempty" name:"CertType"`
	// 业务类型

	BizType *string `json:"BizType,omitempty" name:"BizType"`
	// 登录账户

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 国密证书签名公钥

	SignCertPublicCert *string `json:"SignCertPublicCert,omitempty" name:"SignCertPublicCert"`
	// 国密证书签名私钥

	SignCertPrivateKey *string `json:"SignCertPrivateKey,omitempty" name:"SignCertPrivateKey"`
}

func (r *UploadBgCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadBgCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApdMasterRequest struct {
	*tchttp.BaseRequest

	// 被删除的Mastter列表

	MasterList []*MasterList `json:"MasterList,omitempty" name:"MasterList"`
}

func (r *DeleteApdMasterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApdMasterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportApdRuleRequest struct {
	*tchttp.BaseRequest

	// 导出页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 导出页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ExportApdRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportApdRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Apply struct {
	// 申请ID

	ApplyId *string `json:"ApplyId,omitempty" name:"ApplyId"`
	// 申请时间

	ApplyTime *uint64 `json:"ApplyTime,omitempty" name:"ApplyTime"`
	// 申请人

	Applyer *string `json:"Applyer,omitempty" name:"Applyer"`
	// 申请人备注

	ApplyRemarks *string `json:"ApplyRemarks,omitempty" name:"ApplyRemarks"`
	// 备注时间

	ProposeTime *uint64 `json:"ProposeTime,omitempty" name:"ProposeTime"`
	// 提出人

	Proposer *string `json:"Proposer,omitempty" name:"Proposer"`
	// 备注

	ProposeRemarks *string `json:"ProposeRemarks,omitempty" name:"ProposeRemarks"`
	// 状态标识

	Flag *int64 `json:"Flag,omitempty" name:"Flag"`
	// app名称

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 业务类型

	BizType *string `json:"BizType,omitempty" name:"BizType"`
	// 所属城市

	City *string `json:"City,omitempty" name:"City"`
	// 主域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 协议模式

	FwdMode *string `json:"FwdMode,omitempty" name:"FwdMode"`
	// 协议类型

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 运营商名称

	IspName *string `json:"IspName,omitempty" name:"IspName"`
	// 监听端口

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// 容灾恢复

	DisasterRecovery *int64 `json:"DisasterRecovery,omitempty" name:"DisasterRecovery"`
	// vpc网络ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 任务ID

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 规则Id

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 绑定Vip列表

	Vips []*Vips `json:"Vips,omitempty" name:"Vips"`
	// 校验客户端

	VerifyClient *string `json:"VerifyClient,omitempty" name:"VerifyClient"`
	// http规则ID

	SngCertId *string `json:"SngCertId,omitempty" name:"SngCertId"`
	// http规则名称

	SngCertName *string `json:"SngCertName,omitempty" name:"SngCertName"`
	// https规则id

	SngCaCertId *string `json:"SngCaCertId,omitempty" name:"SngCaCertId"`
	// https规则名称

	SngCaCertName *string `json:"SngCaCertName,omitempty" name:"SngCaCertName"`
	// 所属IDC

	Idc *bool `json:"Idc,omitempty" name:"Idc"`
	// 探测间隔

	ProbeInterval *uint64 `json:"ProbeInterval,omitempty" name:"ProbeInterval"`
	// 探测超时时间

	ProbeTimeout *uint64 `json:"ProbeTimeout,omitempty" name:"ProbeTimeout"`
	// 踢掉时间

	KickInterval *uint64 `json:"KickInterval,omitempty" name:"KickInterval"`
	// 恢复时间

	AliveInterval *uint64 `json:"AliveInterval,omitempty" name:"AliveInterval"`
	// 四层集群标签

	SetLabels []*string `json:"SetLabels,omitempty" name:"SetLabels"`
	// 七层集群标签

	StgwSetLabels []*string `json:"StgwSetLabels,omitempty" name:"StgwSetLabels"`
	// 用途

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// 运营商别名

	IspAlias *string `json:"IspAlias,omitempty" name:"IspAlias"`
}

type RuleOutRsList struct {
	// RS IP

	RsIp *string `json:"RsIp,omitempty" name:"RsIp"`
	// RS 端口

	RsPort *uint64 `json:"RsPort,omitempty" name:"RsPort"`
	// 权重

	RsWeight *uint64 `json:"RsWeight,omitempty" name:"RsWeight"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 宿主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// V6Id，IPv6的rs选填

	V6Id *string `json:"V6Id,omitempty" name:"V6Id"`
	// VpcId

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// RS类型，主要分为linux_tunl、fuzai_ko

	RsType *string `json:"RsType,omitempty" name:"RsType"`
}

type OpBgRsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务编号

		JobId *uint64 `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpBgRsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpBgRsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LocationList struct {
	// 服务端URl

	Url *string `json:"Url,omitempty" name:"Url"`
	// 服务端列表

	RsList []*RsList `json:"RsList,omitempty" name:"RsList"`
}

type RspModuleList struct {
	// Module名称

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// Section列表

	SectionList []*SectionList `json:"SectionList,omitempty" name:"SectionList"`
}

type DeleteBgCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBgCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBgCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportApdRuleL7Request struct {
	*tchttp.BaseRequest

	// 导出页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 导出页大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ExportApdRuleL7Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportApdRuleL7Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductFailureMigrateTaskDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询产品故障转移结果描述

		Data *ProductFailureMigrateTaskRsp `json:"Data,omitempty" name:"Data"`
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

type LdListToSet struct {
	// LD 地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type SetL7HealthCheck struct {
	// 操作类型

	Op *string `json:"Op,omitempty" name:"Op"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 在线阈值

	AliveNum *int64 `json:"AliveNum,omitempty" name:"AliveNum"`
	// 踢出阈值

	KickNum *int64 `json:"KickNum,omitempty" name:"KickNum"`
	// 探测间隔

	ProbeInterval *int64 `json:"ProbeInterval,omitempty" name:"ProbeInterval"`
	// 在线code

	AliveCode *int64 `json:"AliveCode,omitempty" name:"AliveCode"`
	// 服务名称

	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`
	// 探测URL

	ProbeUrl *string `json:"ProbeUrl,omitempty" name:"ProbeUrl"`
}

type DeleteBgRuleRequest struct {
	*tchttp.BaseRequest

	// 清除规则标识

	ClearTunnel *bool `json:"ClearTunnel,omitempty" name:"ClearTunnel"`
	// 删除RS标识

	DelRs *bool `json:"DelRs,omitempty" name:"DelRs"`
	// 同步标识

	Sync *bool `json:"Sync,omitempty" name:"Sync"`
	// 删除的规则列表

	RuleList []*RuleList `json:"RuleList,omitempty" name:"RuleList"`
}

func (r *DeleteBgRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBgRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdRuleL7Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// Rules列表

		Rules []*Rules `json:"Rules,omitempty" name:"Rules"`
		// 返回规则个数

		RulesCount *uint64 `json:"RulesCount,omitempty" name:"RulesCount"`
		// 返回总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Requestid

		Requestid *string `json:"Requestid,omitempty" name:"Requestid"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApdRuleL7Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdRuleL7Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CertVips struct {
	// 所属城市

	City *string `json:"City,omitempty" name:"City"`
	// 运营商名称

	IspName *string `json:"IspName,omitempty" name:"IspName"`
	// 虚拟IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// vpc网络ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 运营商别名

	IspAlias *string `json:"IspAlias,omitempty" name:"IspAlias"`
}

type FilterStrArray struct {
	// 过滤参数名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤参数值

	Value []*string `json:"Value,omitempty" name:"Value"`
}

type StgwVipGroupList struct {
	// 所属IDC

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 提交的VIP列表

	VipList []*DelLdList `json:"VipList,omitempty" name:"VipList"`
}

type DeleteApdVipSegmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApdVipSegmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApdVipSegmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBgRuleL7Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// Requestid

		Requestid *string `json:"Requestid,omitempty" name:"Requestid"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBgRuleL7Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBgRuleL7Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务状态

		JobStatusArr []*JobStatusArr `json:"JobStatusArr,omitempty" name:"JobStatusArr"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBgJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpApdSectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpApdSectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpApdSectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HbLipList struct {
	// Ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// LD IP

	LdIp *string `json:"LdIp,omitempty" name:"LdIp"`
	// 网段划分

	NetSegment *string `json:"NetSegment,omitempty" name:"NetSegment"`
	// Set 名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 子网掩码

	SubMask *string `json:"SubMask,omitempty" name:"SubMask"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type SetBgRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetBgRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBgRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RspSetList struct {
	// 容灾

	DisasterRecovery *int64 `json:"DisasterRecovery,omitempty" name:"DisasterRecovery"`
	// 转发类型

	Fwd *string `json:"Fwd,omitempty" name:"Fwd"`
	// 所属Idc

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 网络运营商

	Isp *string `json:"Isp,omitempty" name:"Isp"`
	// LD列表

	LdList []*string `json:"LdList,omitempty" name:"LdList"`
	// 模式

	Module *string `json:"Module,omitempty" name:"Module"`
	// 系统类型

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 集群模式

	RouteMode *string `json:"RouteMode,omitempty" name:"RouteMode"`
	// 规则数量

	RuleCount *uint64 `json:"RuleCount,omitempty" name:"RuleCount"`
	// Section

	Section *string `json:"Section,omitempty" name:"Section"`
	// Section 列表

	SectionList []*string `json:"SectionList,omitempty" name:"SectionList"`
	// Set名称

	Set *string `json:"Set,omitempty" name:"Set"`
	// SetID

	SetId *uint64 `json:"SetId,omitempty" name:"SetId"`
	// 扩展类型

	SpecialSetType *string `json:"SpecialSetType,omitempty" name:"SpecialSetType"`
	// Tsv 地址数量

	TsvVipCount *uint64 `json:"TsvVipCount,omitempty" name:"TsvVipCount"`
	// Tvs 地址数量

	TvsVipCount *uint64 `json:"TvsVipCount,omitempty" name:"TvsVipCount"`
	// 废弃的Tvs数量

	NouseTvs *uint64 `json:"NouseTvs,omitempty" name:"NouseTvs"`
	// Set类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// Master名称

	MasterName *string `json:"MasterName,omitempty" name:"MasterName"`
	// Master属性

	Master *Master `json:"Master,omitempty" name:"Master"`
	// 业务ip个数

	BusiLipCount *uint64 `json:"BusiLipCount,omitempty" name:"BusiLipCount"`
	// 业务ip列表

	BusiLipList []*BusiLipList `json:"BusiLipList,omitempty" name:"BusiLipList"`
	// 健康检查ip个数

	HbLipCount *uint64 `json:"HbLipCount,omitempty" name:"HbLipCount"`
	// 健康检查ip列表

	HbLipList []*HbLipList `json:"HbLipList,omitempty" name:"HbLipList"`
}

type RuleGroup struct {
	// 规则所分配的IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 所属vpc网络Id

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 规则执行端口

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// "TCP|UDP", //传输层协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 持续时间

	KeepTime *int64 `json:"KeepTime,omitempty" name:"KeepTime"`
	// 探测周期

	ProbeInterval *int64 `json:"ProbeInterval,omitempty" name:"ProbeInterval"`
	// 探测超时阈值

	ProbeTimeOut *int64 `json:"ProbeTimeOut,omitempty" name:"ProbeTimeOut"`
	// 踢出周期

	KickInterval *int64 `json:"KickInterval,omitempty" name:"KickInterval"`
	// 心跳周期

	AliveInterval *int64 `json:"AliveInterval,omitempty" name:"AliveInterval"`
	// 可选,cdc集群规则

	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
}

type DescribeApdMasterListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询Master列表

		MasterList []*Master `json:"MasterList,omitempty" name:"MasterList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApdMasterListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdMasterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdVipGroupRequest struct {
	*tchttp.BaseRequest

	// 查询页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 查询页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 过滤参数，包括：UseStatus，Usage， SetName， Idc， BizId

	Filters []*FilterStr `json:"Filters,omitempty" name:"Filters"`
	// 是否容灾

	DisasterRecovery *bool `json:"DisasterRecovery,omitempty" name:"DisasterRecovery"`
	// 标签数组

	SetLabels []*string `json:"SetLabels,omitempty" name:"SetLabels"`
	// 运营商数组

	IspList []*string `json:"IspList,omitempty" name:"IspList"`
	// 是否STGW,当vport_list不为空时才生效，

	ForStgw *bool `json:"ForStgw,omitempty" name:"ForStgw"`
	// 将返回可以绑定列表中的vport端口的vip组

	VportList []*int64 `json:"VportList,omitempty" name:"VportList"`
	// 是否校验集群标签，默认不校验

	NeedCheckSetlabel *bool `json:"NeedCheckSetlabel,omitempty" name:"NeedCheckSetlabel"`
}

func (r *DescribeApdVipGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdVipGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgRsL7Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询Rs返回

		RsList []*BgRsList `json:"RsList,omitempty" name:"RsList"`
		// Rs数量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// Requestid

		Requestid *string `json:"Requestid,omitempty" name:"Requestid"`
		// virtual server列表

		VsList []*ServiceList `json:"VsList,omitempty" name:"VsList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBgRsL7Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgRsL7Response) FromJsonString(s string) error {
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

type ProcessVipGroupList struct {
	// 业务ID

	BizId *string `json:"BizId,omitempty" name:"BizId"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 转发类型

	Fwd *string `json:"Fwd,omitempty" name:"Fwd"`
	// 组标签

	GroupTag *string `json:"GroupTag,omitempty" name:"GroupTag"`
	// 所属IDC

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 规则数量

	RuleCount *int64 `json:"RuleCount,omitempty" name:"RuleCount"`
	// Set名称列表

	SetNameList []*string `json:"SetNameList,omitempty" name:"SetNameList"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 用途

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// 使用状态

	UseStatus *string `json:"UseStatus,omitempty" name:"UseStatus"`
	// Vip列表

	VipList []*DelVipList `json:"VipList,omitempty" name:"VipList"`
	// VIP组IP

	VipGroupId *string `json:"VipGroupId,omitempty" name:"VipGroupId"`
	// 是否禁用

	Disable *bool `json:"Disable,omitempty" name:"Disable"`
	// 数量

	Num *int64 `json:"Num,omitempty" name:"Num"`
	// 集群标签

	SetLabel *string `json:"SetLabel,omitempty" name:"SetLabel"`
}

type ExportBgCertListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 证书列表

		CertList *string `json:"CertList,omitempty" name:"CertList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBgCertListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBgCertListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IdcList struct {
	// IDC地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// 城市

	City *string `json:"City,omitempty" name:"City"`
	// Master属性

	Master *MasterList `json:"Master,omitempty" name:"Master"`
	// Moudle列表

	ModuleList []*RspModuleList `json:"ModuleList,omitempty" name:"ModuleList"`
	// 7层Master

	StgwMaster *StgwMaster `json:"StgwMaster,omitempty" name:"StgwMaster"`
	// 7层集群列表

	StgwSetList []*L7SetList `json:"StgwSetList,omitempty" name:"StgwSetList"`
	// idc名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type InternetAccessible struct {
	// TRAFFIC_POSTPAID_BY_HOUR 按流量按小时后计费 ; BANDWIDTH_POSTPAID_BY_HOUR 按带宽按小时后计费; BANDWIDTH_PACKAGE 按带宽包计费; 注意：此字段可能返回 null，表示取不到有效值。

	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`
	// 最大出带宽，单位Mbps

	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

type ProductStateInfoRsp struct {
	// 产品名称，取值CLB

	Product *string `json:"Product,omitempty" name:"Product"`
	// 产品状态信息查询接口

	ClusterList []*ClusterState `json:"ClusterList,omitempty" name:"ClusterList"`
	// 状态最后上报时间（更新时间）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type StgwMaster struct {
	// IP地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 模式

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
}

type ModifyLbLimitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改配额返回值

		ModifyRet *int64 `json:"ModifyRet,omitempty" name:"ModifyRet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLbLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLbLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProcessApdApplyRequest struct {
	*tchttp.BaseRequest

	// 实施申请的规则描述

	Apply []*ProcessApply `json:"Apply,omitempty" name:"Apply"`
	// 标识

	Flag *int64 `json:"Flag,omitempty" name:"Flag"`
	// Vip组列表

	VipGroupList []*ProcessVipGroupList `json:"VipGroupList,omitempty" name:"VipGroupList"`
	// 登录账户

	Uid *string `json:"Uid,omitempty" name:"Uid"`
}

func (r *ProcessApdApplyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProcessApdApplyRequest) FromJsonString(s string) error {
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

type BindSetLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindSetLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindSetLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssignApdLdToSetRequest struct {
	*tchttp.BaseRequest

	// Set 集群名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// LD 地址列表

	LdList []*LdListToSet `json:"LdList,omitempty" name:"LdList"`
}

func (r *AssignApdLdToSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssignApdLdToSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssignApdLdToStgwSetRequest struct {
	*tchttp.BaseRequest

	// Set集群名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// LD 地址列表

	LdList []*LdListToSet `json:"LdList,omitempty" name:"LdList"`
}

func (r *AssignApdLdToStgwSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssignApdLdToStgwSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindAppIdLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindAppIdLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindAppIdLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBgServersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回RS列表数量

		ReturnNum *uint64 `json:"ReturnNum,omitempty" name:"ReturnNum"`
		// 查询总数

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// Server列表

		RsList []*RspRsList `json:"RsList,omitempty" name:"RsList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBgServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBgServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductFailureMigrateRsp struct {
	// 切换任务UUID

	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
}

type UpdateApdVipGroupRequest struct {
	*tchttp.BaseRequest

	// 新值

	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`
	// 操作类型

	OpType *string `json:"OpType,omitempty" name:"OpType"`
	// Vip组列表

	VipGroupIdList []*string `json:"VipGroupIdList,omitempty" name:"VipGroupIdList"`
}

func (r *UpdateApdVipGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApdVipGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDefaultIspListRequest struct {
	*tchttp.BaseRequest

	// ipv4 or ipv6, 默认是ipv4

	Type *string `json:"Type,omitempty" name:"Type"`
	// 要设置成默认运营商的ID

	DefaultIspId *int64 `json:"DefaultIspId,omitempty" name:"DefaultIspId"`
}

func (r *UpdateDefaultIspListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDefaultIspListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpRuleList struct {
	// 规则列表

	RuleGroup []*RuleGroup `json:"RuleGroup,omitempty" name:"RuleGroup"`
	// RS系统类型

	RsOsType *string `json:"RsOsType,omitempty" name:"RsOsType"`
	// RS列表

	RsList []*RsList `json:"RsList,omitempty" name:"RsList"`
	// RS 新端口

	NewRsPort *int64 `json:"NewRsPort,omitempty" name:"NewRsPort"`
}

type SetList struct {
	// section名称

	SectionName *string `json:"SectionName,omitempty" name:"SectionName"`
	// Set集群名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 网络类型

	Isp *string `json:"Isp,omitempty" name:"Isp"`
	// 系统类型,"mixed"  or "linux"or "windows", 必填

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// HTTP or IP

	Fwd *string `json:"Fwd,omitempty" name:"Fwd"`
	// 模式,"HA"

	RouteMode *string `json:"RouteMode,omitempty" name:"RouteMode"`
	// Set类型, WAN or LAN

	SetType *string `json:"SetType,omitempty" name:"SetType"`
	//  特殊的set类型值可以为 'common','mass_rule','cdb','cos','vpc_inner', ‘fullnat’ 默认为common

	SpecialSetType *string `json:"SpecialSetType,omitempty" name:"SpecialSetType"`
	// iipv6

	AfType *string `json:"AfType,omitempty" name:"AfType"`
	// 关联Master名称

	MasterName *string `json:"MasterName,omitempty" name:"MasterName"`
	// 集群类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 是否容灾, 默认false

	DisasterRecovery *bool `json:"DisasterRecovery,omitempty" name:"DisasterRecovery"`
	// 备section名称

	DrSectionName *string `json:"DrSectionName,omitempty" name:"DrSectionName"`
	// 上报数据所需, 例如，EIP, 公网集群， 用于设置上报的namespace, url等,EIP集群 -> "EIP"; 内网LB -> "内网LB"； 公网LB -> "公网LB"

	Product *string `json:"Product,omitempty" name:"Product"`
}

type ApplyBgRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyBgRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyBgRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLdLocalIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLdLocalIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLdLocalIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApdLdRequest struct {
	*tchttp.BaseRequest

	// 被删除的LD 列表

	LdList []*DelLdList `json:"LdList,omitempty" name:"LdList"`
}

func (r *DeleteApdLdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApdLdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApdLdFromSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApdLdFromSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApdLdFromSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApdLdFromStgwSetRequest struct {
	*tchttp.BaseRequest

	// Set名称

	SetName *string `json:"SetName,omitempty" name:"SetName"`
	// 被删除的LD IP列表

	LdList []*LdListToSet `json:"LdList,omitempty" name:"LdList"`
}

func (r *DeleteApdLdFromStgwSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApdLdFromStgwSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdMasterRequest struct {
	*tchttp.BaseRequest

	// 查询页

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 查询页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 过滤参数，包括:Ip, Name, Mode,

	Filters []*FilterStr `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApdMasterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdMasterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBgRuleL7Request struct {
	*tchttp.BaseRequest

	// 服务列表

	ServiceList []*ServiceList `json:"ServiceList,omitempty" name:"ServiceList"`
}

func (r *DeleteBgRuleL7Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBgRuleL7Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllIspListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运营商信息

		IspList []*IspList `json:"IspList,omitempty" name:"IspList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllIspListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllIspListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApdLdL7Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询返回的LD列表

		LdList []*RspLdList `json:"LdList,omitempty" name:"LdList"`
		// 返回列表总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 服务器列表

		SvrArr []*string `json:"SvrArr,omitempty" name:"SvrArr"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApdLdL7Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApdLdL7Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitIspListRequest struct {
	*tchttp.BaseRequest

	// 设置运营商信息

	IspList []*IspList `json:"IspList,omitempty" name:"IspList"`
}

func (r *InitIspListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitIspListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LBLocationTargetInfo struct {
	// 转发规则ID

	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
	// 域名，四层规则为空

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// url，四层规则为空

	Url *string `json:"Url,omitempty" name:"Url"`
	// 后端RS

	Targets []*TargetHealthInfo `json:"Targets,omitempty" name:"Targets"`
}
