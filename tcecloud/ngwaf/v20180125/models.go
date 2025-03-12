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

package v20180125

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type WafGetExportAttackLogJobsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetExportAttackLogJobsResData `json:"Data,omitempty" name:"Data"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 错误码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 错误信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetExportAttackLogJobsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetExportAttackLogJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetAttackDetailResRows struct {
	// 用户id

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 攻击内容

	AttackContent *string `json:"AttackContent,omitempty" name:"AttackContent"`
	// 攻击ip

	AttackIp *string `json:"AttackIp,omitempty" name:"AttackIp"`
	// 攻击时间

	AttackTime *string `json:"AttackTime,omitempty" name:"AttackTime"`
	// 攻击类型

	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// 硬盘级别

	RiskLevel *int64 `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 地址

	Uri *string `json:"Uri,omitempty" name:"Uri"`
	// 用户签名

	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`
	// uuid

	UuId *string `json:"UuId,omitempty" name:"UuId"`
	// 动作

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 攻击类型英文

	AttackTypeEn *string `json:"AttackTypeEn,omitempty" name:"AttackTypeEn"`
}

type WafGetUpgredeProgressResData struct {
	// 返回码

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 消息

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 数据

	Rows *WafGetUpgredeProgressResRows `json:"Rows,omitempty" name:"Rows"`
}

type WafGetNgWafHostsRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
	// 时间查询

	Date *WafQueryDate `json:"Date,omitempty" name:"Date"`
	// 页面输入框查询条件

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 页面输入框查询条件

	Search *string `json:"Search,omitempty" name:"Search"`
	// 排序

	Sort *int64 `json:"Sort,omitempty" name:"Sort"`
	// 过滤

	Filter *WafQueryFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *WafGetNgWafHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetNgWafHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTCERuleRequest struct {
	*tchttp.BaseRequest

	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTCERuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTCERuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostStatusRequest struct {
	*tchttp.BaseRequest

	// clb防护域名开关切换数组

	HostsStatus []*HostsStatus `json:"HostsStatus,omitempty" name:"HostsStatus"`
}

func (r *ModifyHostStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySpartaProtectionModeRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 防护状态：10：规则观察&&AI关闭模式，11：规则观察&&AI观察模式，12：规则观察&&AI拦截模式20：规则拦截&&AI关闭模式，21：规则拦截&&AI观察模式，22：规则拦截&&AI拦截模式

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 运营端appid

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
}

func (r *ModifySpartaProtectionModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySpartaProtectionModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetIspsRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
}

func (r *WafGetIspsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetIspsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetEditHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafSetAppToSuperResponse `json:"Data,omitempty" name:"Data"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafSetEditHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetEditHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetUpgredeRuleByManuallyRequest struct {
	*tchttp.BaseRequest
}

func (r *WafSetUpgredeRuleByManuallyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetUpgredeRuleByManuallyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetRuleUpgradeConfRequest struct {
	*tchttp.BaseRequest

	// 分页

	EnabledStatus *string `json:"EnabledStatus,omitempty" name:"EnabledStatus"`
	// 升级时间

	UpgradeTime *string `json:"UpgradeTime,omitempty" name:"UpgradeTime"`
	// 升级url

	UpgradeUrl *string `json:"UpgradeUrl,omitempty" name:"UpgradeUrl"`
	// 升级周期

	UpgradeCycle *int64 `json:"UpgradeCycle,omitempty" name:"UpgradeCycle"`
}

func (r *WafSetRuleUpgradeConfRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetRuleUpgradeConfRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetSwitchCustomRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
	// id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *WafSetSwitchCustomRuleStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetSwitchCustomRuleStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfServiceInstanceQueryListData struct {
	// 服务实例列表

	ServiceInstanceList []*ConfServiceInstanceQueryListServiceInstanceList `json:"ServiceInstanceList,omitempty" name:"ServiceInstanceList"`
}

type GlobalWhiteCond struct {
	// 匹配目标，HTTP-Method,Host,URI,FULL-URL,Parameter,Cookie,HTTP-Header,JSON-Elements

	Target *string `json:"Target,omitempty" name:"Target"`
	// 匹配操作，String-Match 字符串匹配，支持通配符，例如/floder1/*  /floder1/*/index.htm；Regular-Expression-Match 正则表达式匹配， Include 包含（HTTP-Method），Exclude 不包含（HTTP-Method）

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 方法列表，"GET", "POST", "HEAD"

	HttpMethodList []*string `json:"HttpMethodList,omitempty" name:"HttpMethodList"`
	// 变量名称, 适用于Parameter,Cookie,HTTP-Header,JSON-Elements

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否检查变量值，0:disable, 1: enable。 适用于Parameter,Cookie,HTTP-Header,JSON-Elements

	CheckValue *uint64 `json:"CheckValue,omitempty" name:"CheckValue"`
	// 变量值，适用于Host,URI,FULL-URL,Parameter,Cookie,HTTP-Header,JSON-Elements

	Value *string `json:"Value,omitempty" name:"Value"`
	// 条件a之间的链接方式，可以为AND/OR。AND 表示与上一个条件与的关系，OR表示与上一个条件是或的关系。多个条件之间的链接关系如下： C1.AND C2.AND C3.OR C4.OR C5.AND C6.AND, 则组合的逻辑关系是： c1 and (c2 or c3 or c4) and c5 and c6

	Concatenate *string `json:"Concatenate,omitempty" name:"Concatenate"`
}

type DeleteGlobalWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGlobalWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGlobalWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetBwIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetBwIPResData `json:"Data,omitempty" name:"Data"`
		// 返回信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetBwIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetBwIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetCacheRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetCacheRulesResData `json:"Data,omitempty" name:"Data"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetCacheRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetCacheRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsRuleUpgradeRecordQueryListVersion struct {
	// 版本ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 指向Objects中的升级记录

	LogInst *MsRuleUpgradeRecordQueryListLogInst `json:"LogInst,omitempty" name:"LogInst"`
	// InsertTime

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// Operator

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type WafGetExportAttackLogData struct {
	// 返回码

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 导出结果

	ExportCsvResult *string `json:"ExportCsvResult,omitempty" name:"ExportCsvResult"`
	// 结果信息

	Msg *string `json:"Msg,omitempty" name:"Msg"`
}

type ConfNsServiceInstanceMappingDeleteMappingList struct {
	// 服务实例Id

	ServiceInstanceId *int64 `json:"ServiceInstanceId,omitempty" name:"ServiceInstanceId"`
	// 服务配置Id

	ServiceConfId *int64 `json:"ServiceConfId,omitempty" name:"ServiceConfId"`
	// 命名空间Id

	ConfNamespaceId *int64 `json:"ConfNamespaceId,omitempty" name:"ConfNamespaceId"`
}

type WafGetCosCredentialsResSecCredentials struct {
	// token

	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
	// 临时SecretId

	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`
	// 临时SecretKey

	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
}

type WafGetCustomPayloadsResponseResult struct {
	// AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// Category

	Category *string `json:"Category,omitempty" name:"Category"`
	// Domain

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// Flag

	Flag *string `json:"Flag,omitempty" name:"Flag"`
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// LearnStat

	LearnStat *string `json:"LearnStat,omitempty" name:"LearnStat"`
	// ObjectId

	ObjectId *string `json:"ObjectId,omitempty" name:"ObjectId"`
	// Remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// Source

	Source *string `json:"Source,omitempty" name:"Source"`
	// Timestamp

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// Uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type WafGetDomainEngineTypeRequest struct {
	*tchttp.BaseRequest

	// 域名数组

	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

func (r *WafGetDomainEngineTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetDomainEngineTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetQueryHostResRows struct {
	// cls状态

	ClsStatus *int64 `json:"ClsStatus,omitempty" name:"ClsStatus"`
	// 监听器设置

	LoadBalancerSet []*LoadBalancerSetResData `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`
	// 用户id

	Appid *int64 `json:"Appid,omitempty" name:"Appid"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 引擎

	Engine *int64 `json:"Engine,omitempty" name:"Engine"`
	// 流量模式

	FlowMode *int64 `json:"FlowMode,omitempty" name:"FlowMode"`
	// 是否cdn

	IsCdn *int64 `json:"IsCdn,omitempty" name:"IsCdn"`
	// 主要域名

	MainDomain *string `json:"MainDomain,omitempty" name:"MainDomain"`
	// 模式

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 状态

	State *int64 `json:"State,omitempty" name:"State"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
}

type WafListResponseClientData struct {
	// 数据列

	Rows []*WafListResponseDataClientRows `json:"Rows,omitempty" name:"Rows"`
	// 运营端用户id

	SuperApp *int64 `json:"SuperApp,omitempty" name:"SuperApp"`
	// 查询总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type WafGetExportAttackLogJobsResRows struct {
	// 用户id

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 删除表示

	DeleteFlag *int64 `json:"DeleteFlag,omitempty" name:"DeleteFlag"`
	// 失效天

	ExpiredDays *int64 `json:"ExpiredDays,omitempty" name:"ExpiredDays"`
	// 失效时间

	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// 文件下载地址

	FileDownloadUrl *string `json:"FileDownloadUrl,omitempty" name:"FileDownloadUrl"`
	// 文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 日志数

	LogsCount *int64 `json:"LogsCount,omitempty" name:"LogsCount"`
	// 查询条件

	QueryParams *string `json:"QueryParams,omitempty" name:"QueryParams"`
	// task状态

	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`
	// task类型

	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`
	// uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 上传数量

	UploadCount *int64 `json:"UploadCount,omitempty" name:"UploadCount"`
	// 上传进度

	UploadProgress *string `json:"UploadProgress,omitempty" name:"UploadProgress"`
}

type GetTCERuleCosKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Cos临时secret-id

		SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
		// Cos临时secret-key

		SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
		// Cos临时token

		SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
		// 上传Cos的bucket

		Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
		// 上传Cos的region

		Region *string `json:"Region,omitempty" name:"Region"`
		// ExpiredTime

		ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
		// Expiration

		Expiration *string `json:"Expiration,omitempty" name:"Expiration"`
		// 将文件上传到Cos的Key，可在请求参数中指定，留空由后台随机生成

		UploadKey *string `json:"UploadKey,omitempty" name:"UploadKey"`
		// CosUrl

		CosUrl *string `json:"CosUrl,omitempty" name:"CosUrl"`
		// Domain

		Domain *string `json:"Domain,omitempty" name:"Domain"`
		// Schema

		Schema *string `json:"Schema,omitempty" name:"Schema"`
		// StartTime

		StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTCERuleCosKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTCERuleCosKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TCERuleDeleteRequest struct {
	*tchttp.BaseRequest

	// 操作者

	Name *string `json:"Name,omitempty" name:"Name"`
	// 删除的规则库id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *TCERuleDeleteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TCERuleDeleteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetAddCustomRuleRequest struct {
	*tchttp.BaseRequest

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 旁路

	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`
	// 排序id

	SortId *string `json:"SortId,omitempty" name:"SortId"`
	// 如果动作是重定向，则表示重定向的地址；其他情况可以为空

	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`
	// 过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 策略

	Strategies []*WafQueryStrategies `json:"Strategies,omitempty" name:"Strategies"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 请求类型

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

func (r *WafSetAddCustomRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetAddCustomRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfNamespaceQueryListNamespaceList struct {
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 命名空间Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 是否默认

	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`
	// 是否不可用

	IsDisabled *int64 `json:"IsDisabled,omitempty" name:"IsDisabled"`
	// 命名空间名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Note *string `json:"Note,omitempty" name:"Note"`
	// 所属服务配置Id

	ServiceConfId *int64 `json:"ServiceConfId,omitempty" name:"ServiceConfId"`
	// 所属服务配置名称

	ServiceConfName *string `json:"ServiceConfName,omitempty" name:"ServiceConfName"`
	// 所属服务Id

	ServiceId *int64 `json:"ServiceId,omitempty" name:"ServiceId"`
	// 所属服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type WafGetUpgredeProgressRequest struct {
	*tchttp.BaseRequest
}

func (r *WafGetUpgredeProgressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetUpgredeProgressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TigaRuleStatusItem struct {
	// 规则库id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 规则库版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 是否使用中，1代表使用中，0代表没使用

	InUse *int64 `json:"InUse,omitempty" name:"InUse"`
	// 规则库来源，0代表cos

	Source *int64 `json:"Source,omitempty" name:"Source"`
	// 回滚标记，1代表已经下发回滚，0代表没有下发回滚

	Rollback *int64 `json:"Rollback,omitempty" name:"Rollback"`
	// 规则库状态，0代表新建，1代表下发升级，2代表升级中，3代表已完成升级，4代表已删除

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 升级进度

	Progress *int64 `json:"Progress,omitempty" name:"Progress"`
	// 该规则库的错误码

	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 完成升级时间

	FinishedTime *string `json:"FinishedTime,omitempty" name:"FinishedTime"`
	// 规则库详细

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 规则库上次操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 规则库文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

type ModifyOpSignatureRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOpSignatureRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOpSignatureRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WhiteHostQueryListData struct {
	// 域名策略白名单列表

	Objects []*WhiteHostQueryListObjects `json:"Objects,omitempty" name:"Objects"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 页码

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 页大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeOpUserSignaturePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 用户列表

		OpUserSigPolicy []*OpUserSignaturePolicy `json:"OpUserSigPolicy,omitempty" name:"OpUserSigPolicy"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpUserSignaturePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpUserSignaturePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafBanIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *string `json:"Data,omitempty" name:"Data"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafBanIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafBanIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafListResponseDataRuleRows struct {
	// 请求类型

	Action *string `json:"Action,omitempty" name:"Action"`
	// 用户id

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规则

	Rule *string `json:"Rule,omitempty" name:"Rule"`
	// 规则id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type ConfMachineQueryListSvrMachineList struct {
	// City

	City *string `json:"City,omitempty" name:"City"`
	// CpuProcessorNum

	CpuProcessorNum *int64 `json:"CpuProcessorNum,omitempty" name:"CpuProcessorNum"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 服务器Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// Idc

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 服务器IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 备注

	Note *string `json:"Note,omitempty" name:"Note"`
	// 资源组Id

	ResourceGroupId *int64 `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`
	// 服务器状态，0在线 1下线

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// VidcIp

	VidcIp *string `json:"VidcIp,omitempty" name:"VidcIp"`
}

type WafGetBwIPResRows struct {
	// 请求名称

	Action *string `json:"Action,omitempty" name:"Action"`
	// 用户id

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 笔记

	Note *string `json:"Note,omitempty" name:"Note"`
	// 来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 有效时间

	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`
}

type WafGetHostsResRowData struct {
	// appid

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名类型

	DomainType *string `json:"DomainType,omitempty" name:"DomainType"`
}

type WafBanIPRequest struct {
	*tchttp.BaseRequest

	// 分页

	OnOff *string `json:"OnOff,omitempty" name:"OnOff"`
	// 端口

	Ports *string `json:"Ports,omitempty" name:"Ports"`
	// 门槛

	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`
	// 超时时间

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	// 级别

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 到达地址

	Dst *int64 `json:"Dst,omitempty" name:"Dst"`
}

func (r *WafBanIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafBanIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetNgWafHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetNgWafHostsResData `json:"Data,omitempty" name:"Data"`
		// 错误码备注

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetNgWafHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetNgWafHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetRuleUpgradeConfResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafSetRuleUpgradeConfResData `json:"Data,omitempty" name:"Data"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafSetRuleUpgradeConfResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetRuleUpgradeConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudDomainModeSwitchData struct {
	// 更新结果

	Update *bool `json:"Update,omitempty" name:"Update"`
}

type WafSetSwitchCustomRuleStatusResData struct {
	// 消息

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type WafGetExportAttackLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetExportAttackLogData `json:"Data,omitempty" name:"Data"`
		// 返回码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetExportAttackLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetExportAttackLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetLicenseListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		Data []*WafGetLicenseResponse `json:"Data,omitempty" name:"Data"`
		// code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回码描述

		Message *string `json:"Message,omitempty" name:"Message"`
		// Message

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetLicenseListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetLicenseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetCustomConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafSetCustomConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetCustomConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetLicenseResponse struct {
	// license列表详细数据

	Infos *LicenseModel `json:"Infos,omitempty" name:"Infos"`
	// 系统部署时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 0：否 1：是

	NoLicenseOver *int64 `json:"NoLicenseOver,omitempty" name:"NoLicenseOver"`
}

type DescribeOpAttackWhiteRuleNewRequest struct {
	*tchttp.BaseRequest

	// 排序字段，支持 user_id, signature_id, modify_time

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选字段支持 AppId，Domain

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeOpAttackWhiteRuleNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpAttackWhiteRuleNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetFreqRulesResRows struct {
	// 动作

	Action *string `json:"Action,omitempty" name:"Action"`
	// 模式

	Advance *string `json:"Advance,omitempty" name:"Advance"`
	// 用户id

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 间隔

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 匹配方式

	MatchFunc *string `json:"MatchFunc,omitempty" name:"MatchFunc"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 优先事项

	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 地址

	Url *string `json:"Url,omitempty" name:"Url"`
	// 有效时间

	ValidTime *uint64 `json:"ValidTime,omitempty" name:"ValidTime"`
	// 动作英文

	ActionEn *string `json:"ActionEn,omitempty" name:"ActionEn"`
	// 模式英文

	AdvanceEn *string `json:"AdvanceEn,omitempty" name:"AdvanceEn"`
	// 匹配方式英文

	MatchFuncEn *string `json:"MatchFuncEn,omitempty" name:"MatchFuncEn"`
	// 状态英文

	StatusEn *string `json:"StatusEn,omitempty" name:"StatusEn"`
	// 用户id

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type WafSetCustomPayloadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafSetCustomPayloadResData `json:"Data,omitempty" name:"Data"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafSetCustomPayloadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetCustomPayloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostsStatus struct {
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// WAF的开关，1：开，0：关

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 实例id 非必填

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 域名所属的的appid

	DomainAppid *int64 `json:"DomainAppid,omitempty" name:"DomainAppid"`
}

type WafCLBPodInfo struct {
	// Pod ID

	PodId *string `json:"PodId,omitempty" name:"PodId"`
	// 内网IP

	InnerIp *string `json:"InnerIp,omitempty" name:"InnerIp"`
}

type WafGetCustomPayloadsWhere struct {
	// key

	Key *string `json:"Key,omitempty" name:"Key"`
	// value

	Value *string `json:"Value,omitempty" name:"Value"`
	// Result

	Result []*WafGetCustomPayloadsResponseResult `json:"Result,omitempty" name:"Result"`
	// Count

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// Status

	Status *string `json:"Status,omitempty" name:"Status"`
}

type TCERuleUploadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据库中的id

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TCERuleUploadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TCERuleUploadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetCustomConfigRequest struct {
	*tchttp.BaseRequest

	// xff开关

	XffStatus *int64 `json:"XffStatus,omitempty" name:"XffStatus"`
	// post参数个数限制开关

	UlimitArgs *int64 `json:"UlimitArgs,omitempty" name:"UlimitArgs"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 租户id

	Appid *int64 `json:"Appid,omitempty" name:"Appid"`
}

func (r *WafSetCustomConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetCustomConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetCustomPayloadsRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 类别

	Category *string `json:"Category,omitempty" name:"Category"`
	// 排序

	Sort *WafQuerySort `json:"Sort,omitempty" name:"Sort"`
	// WHERE

	Where *WafGetCustomPayloadsWhere `json:"Where,omitempty" name:"Where"`
}

func (r *WafGetCustomPayloadsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetCustomPayloadsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetExportAttackLogJobsRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
}

func (r *WafGetExportAttackLogJobsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetExportAttackLogJobsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetNgWafHostsResRows struct {
	// 客户端id

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 模式

	FlowMode *string `json:"FlowMode,omitempty" name:"FlowMode"`
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// lb信息

	Lb *WafGetNgWafHostsResLb `json:"Lb,omitempty" name:"Lb"`
	// 拦截模式

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// waf开启状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// engine

	Engine *string `json:"Engine,omitempty" name:"Engine"`
	// 是否开启代理

	IsCdn *int64 `json:"IsCdn,omitempty" name:"IsCdn"`
	// xff开关状态

	XffStatus *int64 `json:"XffStatus,omitempty" name:"XffStatus"`
	// post参数个数开关状态

	UlimitArgs *int64 `json:"UlimitArgs,omitempty" name:"UlimitArgs"`
}

type DescribeGlobalWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 规则列表

		List []*GlobalWhiteRule `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGlobalWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGlobalWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostModeRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名ID

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 防护状态：10：规则观察&&AI关闭模式，11：规则观察&&AI观察模式，12：规则观察&&AI拦截模式20：规则拦截&&AI关闭模式，21：规则拦截&&AI观察模式，22：规则拦截&&AI拦截模式

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
	// 运营端Appid

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
}

func (r *ModifyHostModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySpartaProtectionModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySpartaProtectionModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySpartaProtectionModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetLatestRuleVersionResData struct {
	// 返回码

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 返回信息

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 规则版本

	RuleVersion *string `json:"RuleVersion,omitempty" name:"RuleVersion"`
}

type ModifyGlobalWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 规则序号

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 特征序号

	SignatureId *string `json:"SignatureId,omitempty" name:"SignatureId"`
	// 规则状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 匹配规则项列表

	Rules []*GlobalWhiteCond `json:"Rules,omitempty" name:"Rules"`
}

func (r *ModifyGlobalWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGlobalWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonResponseUpdate struct {
	// 更新结果

	Update *bool `json:"Update,omitempty" name:"Update"`
}

type WafSetAddCustomRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafSetSwitchCustomRuleStatusResData `json:"Data,omitempty" name:"Data"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafSetAddCustomRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetAddCustomRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafListResponseDataRows struct {
	// 用户id

	Appid *string `json:"Appid,omitempty" name:"Appid"`
	// cmdb的id

	CmdbId *string `json:"CmdbId,omitempty" name:"CmdbId"`
	// cpu限制

	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`
	// CPU使用率

	CpuUsage *string `json:"CpuUsage,omitempty" name:"CpuUsage"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 版本

	Editon *string `json:"Editon,omitempty" name:"Editon"`
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// L5

	L5 *string `json:"L5,omitempty" name:"L5"`
	// 级别

	Level *string `json:"Level,omitempty" name:"Level"`
	// 内存限制

	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`
	// 内存使用率

	MemUsage *int64 `json:"MemUsage,omitempty" name:"MemUsage"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 评论

	Remarks *string `json:"Remarks,omitempty" name:"Remarks"`
	// 已用

	Used *string `json:"Used,omitempty" name:"Used"`
	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 地区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type WafListResponseDataRuleData struct {
	// 列表数据

	Rows []*WafListResponseDataRuleRows `json:"Rows,omitempty" name:"Rows"`
	// 查询总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type DescribeOpSignatureRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 规则列表

		Rules []*UserSignatureRule `json:"Rules,omitempty" name:"Rules"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpSignatureRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpSignatureRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTCERuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则库总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 规则库详细信息

		Objects []*TigaRuleStatusItem `json:"Objects,omitempty" name:"Objects"`
		// 正在使用中的规则库

		InUseRule *TigaRuleStatusItem `json:"InUseRule,omitempty" name:"InUseRule"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTCERuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTCERuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNgWafHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *string `json:"Data,omitempty" name:"Data"`
		// CodeDesc

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNgWafHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNgWafHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetRuleUpgradeConfResData struct {
	// 返回码

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 返回信息

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 生效数

	Rows *int64 `json:"Rows,omitempty" name:"Rows"`
}

type ConfNsServiceInstanceMappingQueryListData struct {
	// 映射列表数据集合

	ServiceInstanceMappingList []*ConfNsServiceInstanceMappingQueryListServiceInstanceMappingList `json:"ServiceInstanceMappingList,omitempty" name:"ServiceInstanceMappingList"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 页码

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 页大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

type MsRuleUpgradeRecordQueryListData struct {
	// 升级导入记录列表

	Objects []*MsRuleUpgradeRecordQueryListObjects `json:"Objects,omitempty" name:"Objects"`
	// 当前版本信息

	Version *MsRuleUpgradeRecordQueryListVersion `json:"Version,omitempty" name:"Version"`
	// 页码

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 页大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type WafUpsertSysLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回码

		Code *string `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafUpsertSysLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafUpsertSysLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafQueryDate struct {
	// 起始时间

	From *string `json:"From,omitempty" name:"From"`
	// 结束时间

	To *string `json:"To,omitempty" name:"To"`
	// 长度

	Length *int64 `json:"Length,omitempty" name:"Length"`
}

type WafDelCustomPayloadsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafDelCustomPayloadsResData `json:"Data,omitempty" name:"Data"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafDelCustomPayloadsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafDelCustomPayloadsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetProductUsagesRequest struct {
	*tchttp.BaseRequest
}

func (r *WafGetProductUsagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetProductUsagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetNgWafVipRemarkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回码

		Code *string `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafSetNgWafVipRemarkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetNgWafVipRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOpSignatureRuleRequest struct {
	*tchttp.BaseRequest

	// 要修改的规则id

	SignatureIDs []*string `json:"SignatureIDs,omitempty" name:"SignatureIDs"`
	// 要修改的状态 只能为0/1/2三种状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyOpSignatureRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOpSignatureRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetNgWafVipConfRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
}

func (r *WafGetNgWafVipConfRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetNgWafVipConfRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetUpgredeRuleByManuallyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafSetUpgredeRuleByManuallyResData `json:"Data,omitempty" name:"Data"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafSetUpgredeRuleByManuallyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetUpgredeRuleByManuallyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetDomainEngineTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 引擎数据

		Data []*WafGetDomainEngineTypeData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetDomainEngineTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetDomainEngineTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsRuleQueryListData struct {
	// 策略列表

	Objects []*MsRuleQueryListObjects `json:"Objects,omitempty" name:"Objects"`
	// 页码

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 页大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type WafGetBwIPResData struct {
	// 表格数据

	Rows []*WafGetBwIPResRows `json:"Rows,omitempty" name:"Rows"`
	// 数据总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type WafDeleteInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *WafDeleteInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafDeleteInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafDeleteSpartaProtectionRequest struct {
	*tchttp.BaseRequest

	// data数组

	Data []*WafQueryDelHostData `json:"Data,omitempty" name:"Data"`
}

func (r *WafDeleteSpartaProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafDeleteSpartaProtectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetQpsResData struct {
	// qps

	Qps *int64 `json:"Qps,omitempty" name:"Qps"`
	// 时间戳

	Time *int64 `json:"Time,omitempty" name:"Time"`
}

type WafGetAttackDetailRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	From *string `json:"From,omitempty" name:"From"`
	// 结束时间

	To *string `json:"To,omitempty" name:"To"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 内容游标

	Context *string `json:"Context,omitempty" name:"Context"`
	// 数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 攻击类型

	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`
	// 攻击ip

	AttackIp *string `json:"AttackIp,omitempty" name:"AttackIp"`
	// 规则id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *WafGetAttackDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetAttackDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetRuleUpgradeRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetRuleUpgradeRecordResData `json:"Data,omitempty" name:"Data"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetRuleUpgradeRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetRuleUpgradeRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetCosCredentialsRequest struct {
	*tchttp.BaseRequest

	// 链接

	Origin *string `json:"Origin,omitempty" name:"Origin"`
}

func (r *WafGetCosCredentialsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetCosCredentialsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetCacheRulesResData struct {
	// 表格数据

	Rows []*WafGetCacheRulesResRows `json:"Rows,omitempty" name:"Rows"`
	// 数据总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type MsRuleQueryListTypeInfo struct {
	// 策略所属类型名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 策略所属类型描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 策略所属类型英文名称

	EnName *string `json:"EnName,omitempty" name:"EnName"`
	// 策略所属类型英文描述

	EnDesc *string `json:"EnDesc,omitempty" name:"EnDesc"`
}

type DescribeOpUserSignaturePolicyRequest struct {
	*tchttp.BaseRequest

	// 排序字段，仅支持 level

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选条件，支持 AppId，Domain

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeOpUserSignaturePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpUserSignaturePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetClientsRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
	// 时间查询条件

	Date *WafQueryDate `json:"Date,omitempty" name:"Date"`
	// 过滤条件

	Filter *WafQueryFilter `json:"Filter,omitempty" name:"Filter"`
	// 精确查询条件

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 精确查询条件

	Search *string `json:"Search,omitempty" name:"Search"`
	// 精确查询条件

	Sort *int64 `json:"Sort,omitempty" name:"Sort"`
}

func (r *WafGetClientsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetClientsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetNgWafVipConfResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafSetNgWafVipConfResData `json:"Data,omitempty" name:"Data"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafSetNgWafVipConfResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetNgWafVipConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetNgWafVipRemarkRequest struct {
	*tchttp.BaseRequest

	// id标识

	Id *string `json:"Id,omitempty" name:"Id"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *WafSetNgWafVipRemarkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetNgWafVipRemarkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetWafHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetHostsResData `json:"Data,omitempty" name:"Data"`
		// 错误码备注

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetWafHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetWafHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LicenseModel struct {
	// 计费项

	ProductIdDesc *string `json:"ProductIdDesc,omitempty" name:"ProductIdDesc"`
	// 已购买数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 支持业务峰值/QPS

	QPS *string `json:"QPS,omitempty" name:"QPS"`
	// 已使用数量

	UsedCount *uint64 `json:"UsedCount,omitempty" name:"UsedCount"`
	// 使用率 按%返回 精确到小数点两位

	UsedRate *string `json:"UsedRate,omitempty" name:"UsedRate"`
	// 过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 使用状态： Normal表示资源状态正常，Overused 表示超用， Unauthorized表示未授权，Overdue表示超过有效期，此字段只会显示一个状态， 按 Unauthorized > Overdue > Overused > Normal优先级

	Status *string `json:"Status,omitempty" name:"Status"`
	// #采集状态 Normal 正常，否则异常

	CollectStatus *string `json:"CollectStatus,omitempty" name:"CollectStatus"`
	// 生效时间

	EffectiveTime *string `json:"EffectiveTime,omitempty" name:"EffectiveTime"`
}

type WafGetExportAttackLogResData struct {
	// ExportCsvResult

	ExportCsvResult *string `json:"ExportCsvResult,omitempty" name:"ExportCsvResult"`
	// 返回码

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 返回信息

	Msg *string `json:"Msg,omitempty" name:"Msg"`
}

type WafQueryDelHostData struct {
	// 用户id

	ClientAppid *int64 `json:"ClientAppid,omitempty" name:"ClientAppid"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 规则域名

	RuleDomain *string `json:"RuleDomain,omitempty" name:"RuleDomain"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 流量模式

	FlowMode *int64 `json:"FlowMode,omitempty" name:"FlowMode"`
	// 状态

	State *int64 `json:"State,omitempty" name:"State"`
	// 是否cdn

	IsCdn *int64 `json:"IsCdn,omitempty" name:"IsCdn"`
	// 引擎

	Engine *int64 `json:"Engine,omitempty" name:"Engine"`
	// 模式

	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
}

type WafGetQueryHostRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
	// 用户id

	ClientAppid *int64 `json:"ClientAppid,omitempty" name:"ClientAppid"`
	// 查询条件

	Search *string `json:"Search,omitempty" name:"Search"`
}

func (r *WafGetQueryHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetQueryHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetDomainEngineTypeData struct {
	// 引擎类型

	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type DescribeOpUserSignatureRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 规则详细

		Rules []*UserSignatureRule `json:"Rules,omitempty" name:"Rules"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpUserSignatureRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpUserSignatureRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetLatestRuleVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetLatestRuleVersionResData `json:"Data,omitempty" name:"Data"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetLatestRuleVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetLatestRuleVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafVipUpdateRsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误码备注

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 数据

		Data *string `json:"Data,omitempty" name:"Data"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafVipUpdateRsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafVipUpdateRsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleUpdateLog struct {
	// id值

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 详细修改动态

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// en/cn 描述信息语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 版本信息

	LogVersion *string `json:"LogVersion,omitempty" name:"LogVersion"`
}

type WafDelNgWafVipConfData struct {
	// 返回码

	Code *int64 `json:"Code,omitempty" name:"Code"`
}

type WafGetCacheRulesResRows struct {
	// 用户id

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 附件

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 规则id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// Uri

	Uri *string `json:"Uri,omitempty" name:"Uri"`
}

type WafDelNgWafVipConfResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafDelNgWafVipConfData `json:"Data,omitempty" name:"Data"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafDelNgWafVipConfResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafDelNgWafVipConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetExportAttackLogRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	From *string `json:"From,omitempty" name:"From"`
	// 结束时间

	To *string `json:"To,omitempty" name:"To"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 攻击ip

	AttackIp *string `json:"AttackIp,omitempty" name:"AttackIp"`
	// 攻击AttackType

	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`
	// 规则id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 0-默认中文，1-英文

	Intl *int64 `json:"Intl,omitempty" name:"Intl"`
}

func (r *WafGetExportAttackLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetExportAttackLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetDomainEngineRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 引擎类型，tiga或menshen

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *WafSetDomainEngineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetDomainEngineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSysLogConfig struct {
	// 开关状态 1-开启 0-关闭（默认）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 日志类型 1-攻击日志  2-访问日志  目前仅支持攻击日志

	LogType *int64 `json:"LogType,omitempty" name:"LogType"`
	// 网络协议 1-TCP 2-UDP 3-TLS

	Protocol *int64 `json:"Protocol,omitempty" name:"Protocol"`
	// 消息协议 1-RFC3164 2-RFC5424

	MsgProtocol *int64 `json:"MsgProtocol,omitempty" name:"MsgProtocol"`
	// syslog服务端地址

	LogHost *string `json:"LogHost,omitempty" name:"LogHost"`
	// syslog服务端端口

	LogPort *int64 `json:"LogPort,omitempty" name:"LogPort"`
}

type AppIdAddData struct {
	// 调用接口成功情况下 Update 表示 Mode 是否更新

	Update *bool `json:"Update,omitempty" name:"Update"`
}

type WafGetRSRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
}

func (r *WafGetRSRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetRSRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetEnableSysLogRequest struct {
	*tchttp.BaseRequest

	// ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 开关状态 1-开启 0-关闭（默认）

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *WafSetEnableSysLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetEnableSysLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfNsServiceInstanceMappingQueryListServiceInstanceMappingList struct {
	// 命名空间Id

	ConfNamespaceId *int64 `json:"ConfNamespaceId,omitempty" name:"ConfNamespaceId"`
	// 命名空间名称

	ConfNamespaceName *string `json:"ConfNamespaceName,omitempty" name:"ConfNamespaceName"`
	// 服务Id

	ServiceId *int64 `json:"ServiceId,omitempty" name:"ServiceId"`
	// 服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务配置Id

	ServiceConfId *int64 `json:"ServiceConfId,omitempty" name:"ServiceConfId"`
	// 服务配置名称

	ServiceConfName *string `json:"ServiceConfName,omitempty" name:"ServiceConfName"`
	// 服务器Id

	SvrMachineId *int64 `json:"SvrMachineId,omitempty" name:"SvrMachineId"`
	// 服务器Ip

	SvrMachineIp *string `json:"SvrMachineIp,omitempty" name:"SvrMachineIp"`
	// 服务实例

	ServiceInstanceId *int64 `json:"ServiceInstanceId,omitempty" name:"ServiceInstanceId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最后更新时间

	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`
	// 笨猪

	Note *string `json:"Note,omitempty" name:"Note"`
}

type MsRuleUpgradeRecordQueryListObjects struct {
	// 记录ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 升级版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 升级包名称

	Filename *string `json:"Filename,omitempty" name:"Filename"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 导入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 操作人（默认为root）

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 升级的基础防护策略ID列表

	RuleList []*int64 `json:"RuleList,omitempty" name:"RuleList"`
	// StrategyGroupId

	StrategyGroupId *int64 `json:"StrategyGroupId,omitempty" name:"StrategyGroupId"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// Env

	Env *string `json:"Env,omitempty" name:"Env"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
}

type AddGlobalWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddGlobalWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddGlobalWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetBanIPRequest struct {
	*tchttp.BaseRequest
}

func (r *WafGetBanIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetBanIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetConfigSwitchInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data []*WafGetConfigSwitchInfoResponseDataInfo `json:"Data,omitempty" name:"Data"`
		// 返回码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetConfigSwitchInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetConfigSwitchInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetNonstandardPortRequest struct {
	*tchttp.BaseRequest

	// 客户appid

	CustomAppid *string `json:"CustomAppid,omitempty" name:"CustomAppid"`
}

func (r *WafGetNonstandardPortRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetNonstandardPortRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafUpsertNonstandardPortResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回码

		Code *string `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafUpsertNonstandardPortResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafUpsertNonstandardPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetCheckDomainRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *WafGetCheckDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetCheckDomainRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafUpdateConfigSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafUpdateConfigSwitchResData `json:"Data,omitempty" name:"Data"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafUpdateConfigSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafUpdateConfigSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetQpsRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 起始时间

	From *string `json:"From,omitempty" name:"From"`
	// 结束时间

	To *string `json:"To,omitempty" name:"To"`
}

func (r *WafGetQpsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetQpsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfServiceQueryListServiceList struct {
	// 服务ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 服务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 服务备注

	Note *string `json:"Note,omitempty" name:"Note"`
	// 创建人

	CreateName *string `json:"CreateName,omitempty" name:"CreateName"`
	// 更新人

	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type GlobalWhiteRule struct {
	// 序号

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 规则开关

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 特征序号

	SignatureId *string `json:"SignatureId,omitempty" name:"SignatureId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 详细规则

	Rules []*GlobalWhiteCond `json:"Rules,omitempty" name:"Rules"`
}

type MainClassItem struct {
	// 主类id

	MainClassID *string `json:"MainClassID,omitempty" name:"MainClassID"`
	// 主类名字

	MainClassName *string `json:"MainClassName,omitempty" name:"MainClassName"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 子类

	SubClass []*SubClassItem `json:"SubClass,omitempty" name:"SubClass"`
	// 当前主类的规则个数

	RuleCount *uint64 `json:"RuleCount,omitempty" name:"RuleCount"`
}

type EditOpRuleUpdateLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditOpRuleUpdateLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditOpRuleUpdateLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TCERuleUpgradeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 启动升级的规则库id

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TCERuleUpgradeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TCERuleUpgradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafDeleteSpartaProtectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafSetDeleteHostResData `json:"Data,omitempty" name:"Data"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafDeleteSpartaProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafDeleteSpartaProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudDomainQueryListObjects struct {
	// AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 插入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 模式

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 备注

	Note *string `json:"Note,omitempty" name:"Note"`
}

type WafGetWafDomainLoadbalancerResData struct {
	// 数据集合

	Rows []*WafGetWafDomainLoadbalancerResRows `json:"Rows,omitempty" name:"Rows"`
	// 数据总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type WafInstanceData struct {
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 客户id

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 实例类型 clb-waf：负载均衡型实例，sparta-waf：saas型实例

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 主域名数

	MainDomainLimit *string `json:"MainDomainLimit,omitempty" name:"MainDomainLimit"`
	// 二级域名数

	SubDomainLimit *string `json:"SubDomainLimit,omitempty" name:"SubDomainLimit"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 资源id

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type WafListResponseDataHostData struct {
	// 表格列表

	Rows []*WafListResponseDataHostRows `json:"Rows,omitempty" name:"Rows"`
	// 总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type ModifyHostStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetCustomPayloadsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetCustomPayloadsWhere `json:"Data,omitempty" name:"Data"`
		// 返回码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetCustomPayloadsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetCustomPayloadsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetAppToSuperRequest struct {
	*tchttp.BaseRequest

	// 租户Id

	ClientAppId *uint64 `json:"ClientAppId,omitempty" name:"ClientAppId"`
}

func (r *WafSetAppToSuperRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetAppToSuperRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafQueryLoadBalancerSet struct {
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 监听器名称

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// vip端口

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// 地区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 约定

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// biz的id

	BizId *string `json:"BizId,omitempty" name:"BizId"`
}

type ConfMachineQueryListData struct {
	// 服务器列表数据集合

	SvrMachineList []*ConfMachineQueryListSvrMachineList `json:"SvrMachineList,omitempty" name:"SvrMachineList"`
}

type WafDelCustomPayloadsResData struct {
	// 返回结果

	Results []*string `json:"Results,omitempty" name:"Results"`
	// 返回状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type WafGetQueryHostResData struct {
	// 数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 数据集合

	Rows []*WafGetQueryHostResRows `json:"Rows,omitempty" name:"Rows"`
	// 信息

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// Data

	Data *WafGetQueryHostResData `json:"Data,omitempty" name:"Data"`
}

type WafGetRuleUpgradeConfResData struct {
	// 可用状态

	EnabledStatus *int64 `json:"EnabledStatus,omitempty" name:"EnabledStatus"`
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 升级循环

	UpgradeCycle *int64 `json:"UpgradeCycle,omitempty" name:"UpgradeCycle"`
	// 升级时间

	UpgradeTime *string `json:"UpgradeTime,omitempty" name:"UpgradeTime"`
	// 升级地址

	UpgradeUrl *string `json:"UpgradeUrl,omitempty" name:"UpgradeUrl"`
}

type WafAddRsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *string `json:"Data,omitempty" name:"Data"`
		// CodeDesc

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafAddRsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafAddRsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetCustomPayloadRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 客户id

	ClientAppid *string `json:"ClientAppid,omitempty" name:"ClientAppid"`
	// 类别

	Category *string `json:"Category,omitempty" name:"Category"`
	// 来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 评论

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 标志

	Flag *string `json:"Flag,omitempty" name:"Flag"`
}

func (r *WafSetCustomPayloadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetCustomPayloadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafNonstandardPort struct {
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 客户appid

	CustomAppid *string `json:"CustomAppid,omitempty" name:"CustomAppid"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// http非标端口

	HttpPort *string `json:"HttpPort,omitempty" name:"HttpPort"`
	// https非标端口

	HttpsPort *string `json:"HttpsPort,omitempty" name:"HttpsPort"`
}

type WafQueryStrategies struct {
	// 地域id

	Field *string `json:"Field,omitempty" name:"Field"`
	// 。

	Arg *string `json:"Arg,omitempty" name:"Arg"`
	// 。

	CompareFunc *string `json:"CompareFunc,omitempty" name:"CompareFunc"`
	// 内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

type CheckGlobalWhiteRuleNameRequest struct {
	*tchttp.BaseRequest

	// 要查询的名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CheckGlobalWhiteRuleNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckGlobalWhiteRuleNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubClassItem struct {
	// 子类id

	SubClassID *string `json:"SubClassID,omitempty" name:"SubClassID"`
	// 子类名字

	SubClassName *string `json:"SubClassName,omitempty" name:"SubClassName"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type WafGetQpsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data []*WafGetQpsResData `json:"Data,omitempty" name:"Data"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetQpsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetQpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetConfigSwitchInfoResponseDataInfo struct {
	// 配置内容

	ConfigContent *string `json:"ConfigContent,omitempty" name:"ConfigContent"`
	// 配置键

	ConfigKey *string `json:"ConfigKey,omitempty" name:"ConfigKey"`
	// 配置名称

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
	// 配置类型

	ConfigType *uint64 `json:"ConfigType,omitempty" name:"ConfigType"`
	// 配置值

	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 配置内容英文

	ConfigContentEn *string `json:"ConfigContentEn,omitempty" name:"ConfigContentEn"`
	// 配置名称英文

	ConfigNameEn *string `json:"ConfigNameEn,omitempty" name:"ConfigNameEn"`
}

type DeleteGlobalWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 规则序号

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteGlobalWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGlobalWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditOpRuleUpdateLogStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditOpRuleUpdateLogStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditOpRuleUpdateLogStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafAddVipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *string `json:"Data,omitempty" name:"Data"`
		// CodeDesc

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafAddVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafAddVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetCacheRulesRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
	// 过滤

	Filter *WafQueryFilter `json:"Filter,omitempty" name:"Filter"`
	// 筛选条件

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *WafGetCacheRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetCacheRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpRuleUpdateLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则日志总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 规则日志列表

		List []*RuleUpdateLog `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpRuleUpdateLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpRuleUpdateLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetFreqRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetFreqRulesResData `json:"Data,omitempty" name:"Data"`
		// 返回信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetFreqRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetFreqRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetUpgredeProgressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetUpgredeProgressResData `json:"Data,omitempty" name:"Data"`
		// 返回内容

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetUpgredeProgressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetUpgredeProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafLicenseProductModel struct {
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 版本 如3.8.0

	Version *string `json:"Version,omitempty" name:"Version"`
	// 返回计费项数组

	Data []*ProductUsages `json:"Data,omitempty" name:"Data"`
}

type CheckGlobalWhiteRuleNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1=有重复 0=无重复

		Duplicate *int64 `json:"Duplicate,omitempty" name:"Duplicate"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckGlobalWhiteRuleNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckGlobalWhiteRuleNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpMainClassRequest struct {
	*tchttp.BaseRequest

	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选条件，支持 MainClassID

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
	// cn/en

	OpLanguage *string `json:"OpLanguage,omitempty" name:"OpLanguage"`
}

func (r *DescribeOpMainClassRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpMainClassRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetSwitchCustomRuleStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafSetSwitchCustomRuleStatusResData `json:"Data,omitempty" name:"Data"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafSetSwitchCustomRuleStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetSwitchCustomRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsRuleCosKeyData struct {
	// Cos临时secret-id

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// Cos临时secret-key

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
	// Cos临时token

	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
	// 上传Cos的bucket

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 上传Cos的region

	Region *string `json:"Region,omitempty" name:"Region"`
	// ExpiredTime

	ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// Expiration

	Expiration *string `json:"Expiration,omitempty" name:"Expiration"`
	// 将文件上到到Cos的Key，可在请求参数中指定，留空由后台随机生成

	UploadKey *string `json:"UploadKey,omitempty" name:"UploadKey"`
	// CosUrl

	CosUrl *string `json:"CosUrl,omitempty" name:"CosUrl"`
	// Domain

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// Schema

	Schema *string `json:"Schema,omitempty" name:"Schema"`
	// StartTime

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
}

type WafQueryHost struct {
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则域名

	RuleDomain *string `json:"RuleDomain,omitempty" name:"RuleDomain"`
	// 监听器

	LoadBalancerSet []*WafQueryLoadBalancerSet `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`
	// 状态

	State *int64 `json:"State,omitempty" name:"State"`
	// 是否cdn

	IsCdn *int64 `json:"IsCdn,omitempty" name:"IsCdn"`
	// 域名id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
}

type WafDeleteInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回码

		Code *string `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafDeleteInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafDeleteInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetHostsRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
	// 查询时间

	Date *WafQueryDate `json:"Date,omitempty" name:"Date"`
	// 查询条件

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 查询条件

	Search *string `json:"Search,omitempty" name:"Search"`
	// 排序

	Sort *int64 `json:"Sort,omitempty" name:"Sort"`
	// 筛选条件

	Filter *WafQueryFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *WafGetHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetRSResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误码备注

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 数据

		Data *string `json:"Data,omitempty" name:"Data"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetRSResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetRSResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditLogQueryListObjects struct {
	// 内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 插入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 子类型

	SubType *string `json:"SubType,omitempty" name:"SubType"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// Keyword1

	Keyword1 *string `json:"Keyword1,omitempty" name:"Keyword1"`
	// Keyword2

	Keyword2 *string `json:"Keyword2,omitempty" name:"Keyword2"`
}

type WafGetWafHostsRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
}

func (r *WafGetWafHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetWafHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafUpsertNonstandardPortRequest struct {
	*tchttp.BaseRequest

	// 客户appid

	CustomAppid *string `json:"CustomAppid,omitempty" name:"CustomAppid"`
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// http非标端口，多个用英文逗号隔开

	HttpPort *string `json:"HttpPort,omitempty" name:"HttpPort"`
	// https非标端口，多个用英文逗号隔开

	HttpsPort *string `json:"HttpsPort,omitempty" name:"HttpsPort"`
}

func (r *WafUpsertNonstandardPortRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafUpsertNonstandardPortRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetClientsResData struct {
	// dst

	Dst *string `json:"Dst,omitempty" name:"Dst"`
	// 级别

	Level *string `json:"Level,omitempty" name:"Level"`
	// 开关

	OnOff *string `json:"OnOff,omitempty" name:"OnOff"`
	// 端口

	Ports *string `json:"Ports,omitempty" name:"Ports"`
	// 门槛

	Threshold *string `json:"Threshold,omitempty" name:"Threshold"`
	// 超时时间

	Timeout *string `json:"Timeout,omitempty" name:"Timeout"`
}

type DescribeCustomWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *string `json:"Data,omitempty" name:"Data"`
		// CodeDesc

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetSysLogRequest struct {
	*tchttp.BaseRequest
}

func (r *WafGetSysLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetSysLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetRuleUpgradeRecordResData struct {
	// 返回码

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 返回信息

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 表单数据

	Rows []*WafGetRuleUpgradeRecordResRows `json:"Rows,omitempty" name:"Rows"`
	// 表单总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type AuditLogQueryListData struct {
	// 日志列表

	Objects []*AuditLogQueryListObjects `json:"Objects,omitempty" name:"Objects"`
	// 页码

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 页大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type ConfServiceConfQueryListServiceConfList struct {
	// 创建人

	CreateName *string `json:"CreateName,omitempty" name:"CreateName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 服务配置文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 服务配置ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// IsDisabled默认为0，忽略此字段

	IsDisabled *int64 `json:"IsDisabled,omitempty" name:"IsDisabled"`
	// IsWithStrategyGroup默认为0，忽略此字段

	IsWithStrategyGroup *int64 `json:"IsWithStrategyGroup,omitempty" name:"IsWithStrategyGroup"`
	// 服务配置名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Note *string `json:"Note,omitempty" name:"Note"`
	// 所属服务ID

	ServiceId *int64 `json:"ServiceId,omitempty" name:"ServiceId"`
	// 所属服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 配置刷新脚本名称

	TaskScriptName *string `json:"TaskScriptName,omitempty" name:"TaskScriptName"`
	// 服务配置文件类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 更新人

	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type MsRuleUpgradeRecordQueryListLogInst struct {
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// Filename

	Filename *string `json:"Filename,omitempty" name:"Filename"`
}

type WafGetClientQpsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data []*WafGetQpsResData `json:"Data,omitempty" name:"Data"`
		// 错误码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 错误码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 错误信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetClientQpsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetClientQpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetProductUsagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品使用量数组

		Data *WafLicenseProductModel `json:"Data,omitempty" name:"Data"`
		// code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetProductUsagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetProductUsagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpUserWhiteRuleNew struct {
	// appid

	OpAppId *string `json:"OpAppId,omitempty" name:"OpAppId"`
	// 域名

	OpDomain *string `json:"OpDomain,omitempty" name:"OpDomain"`
	// id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 规则id

	SignatureId *string `json:"SignatureId,omitempty" name:"SignatureId"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 规则名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 匹配url

	MatchContent *string `json:"MatchContent,omitempty" name:"MatchContent"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// "eq"

	MatchMethod *string `json:"MatchMethod,omitempty" name:"MatchMethod"`
	// "URL"

	MatchField *string `json:"MatchField,omitempty" name:"MatchField"`
}

type WafSetAddIpv6HostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafSetAppToSuperResponse `json:"Data,omitempty" name:"Data"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafSetAddIpv6HostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetAddIpv6HostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetAppToSuperResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 消息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafSetAppToSuperResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetAppToSuperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoadBalancerSetResData struct {
	// 监听id

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 加载

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 附件

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 端口

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 地区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// bizid

	BizId *int64 `json:"BizId,omitempty" name:"BizId"`
	// 监听名

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
	// 负载均衡名

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
}

type ModifyGlobalWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则序号

		RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGlobalWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGlobalWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TCERuleRollbackRequest struct {
	*tchttp.BaseRequest

	// 操作者

	Name *string `json:"Name,omitempty" name:"Name"`
	// 回滚的规则库id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *TCERuleRollbackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TCERuleRollbackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafAddTestClientRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
}

func (r *WafAddTestClientRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafAddTestClientRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafDeleteNonstandardPortResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回码

		Code *string `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafDeleteNonstandardPortResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafDeleteNonstandardPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsRuleQueryListObjects struct {
	// 策略id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 漏洞编号

	Cve *string `json:"Cve,omitempty" name:"Cve"`
	// 中文描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 英文描述

	EnDesc *string `json:"EnDesc,omitempty" name:"EnDesc"`
	// 规则等级
	// 100 宽松
	// 200 中等
	// 300 严格
	// 400 超严格

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 策略标签

	Tag *string `json:"Tag,omitempty" name:"Tag"`
	// 策略类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 策略所属类型

	TypeInfo *MsRuleQueryListTypeInfo `json:"TypeInfo,omitempty" name:"TypeInfo"`
	// 插入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type WhiteHostQueryListObjects struct {
	// 策略Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 域名

	Host *string `json:"Host,omitempty" name:"Host"`
	// 加白策略列表

	RuleIdList []*int64 `json:"RuleIdList,omitempty" name:"RuleIdList"`
	// 增量加白列表

	AppendRuleIdList []*int64 `json:"AppendRuleIdList,omitempty" name:"AppendRuleIdList"`
	// 策略版本

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 插入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// cgi

	Cgi *string `json:"Cgi,omitempty" name:"Cgi"`
	// CgiPattern cgi匹配模式
	// eq 完全匹配
	// prefix 前缀匹配
	// suffix 后缀匹配

	CgiPattern *string `json:"CgiPattern,omitempty" name:"CgiPattern"`
}

type WafSetCustomPayloadResData struct {
	// InsertedId

	InsertedId *string `json:"InsertedId,omitempty" name:"InsertedId"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type WafGetNgWafVipConfResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetNgWafVipConfResData `json:"Data,omitempty" name:"Data"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetNgWafVipConfResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetNgWafVipConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafUpdateConfigSwitchRequest struct {
	*tchttp.BaseRequest

	// ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 配置值

	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`
	// 配置类型

	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`
	// 配置键

	ConfigKey *string `json:"ConfigKey,omitempty" name:"ConfigKey"`
}

func (r *WafUpdateConfigSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafUpdateConfigSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetCosCredentialsResData struct {
	// 存储桶

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 返回码

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 返回规则相关信息

	Credentials *WafGetCosCredentialsResCredentials `json:"Credentials,omitempty" name:"Credentials"`
	// 文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 域名

	Host *string `json:"Host,omitempty" name:"Host"`
	// 饭回信息

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 地址

	Url *string `json:"Url,omitempty" name:"Url"`
}

type DeleteOpRuleUpdateLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOpRuleUpdateLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOpRuleUpdateLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpRuleUpdateLogRequest struct {
	*tchttp.BaseRequest

	// 排序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 语言cn/en

	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

func (r *DescribeOpRuleUpdateLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpRuleUpdateLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetAddIpv6HostRequest struct {
	*tchttp.BaseRequest

	// 区域

	ClientRegion *string `json:"ClientRegion,omitempty" name:"ClientRegion"`
	// 状态

	State *int64 `json:"State,omitempty" name:"State"`
	// 是否cdn

	IsCdn *int64 `json:"IsCdn,omitempty" name:"IsCdn"`
	// host

	Host *WafQueryHost `json:"Host,omitempty" name:"Host"`
}

func (r *WafSetAddIpv6HostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetAddIpv6HostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafListResponseDataHostRows struct {
	// 用户id

	Appid *string `json:"Appid,omitempty" name:"Appid"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 引擎

	Engine *string `json:"Engine,omitempty" name:"Engine"`
	// 到期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 状态

	State *string `json:"State,omitempty" name:"State"`
	// 开启状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 模式

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 地址

	Src *string `json:"Src,omitempty" name:"Src"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

type WafListResponseDataClientRows struct {
	// 用户id

	Appid *string `json:"Appid,omitempty" name:"Appid"`
	// WAF类型

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 主域名数

	MainDomainLimit *string `json:"MainDomainLimit,omitempty" name:"MainDomainLimit"`
	// 子域名数

	SubDomainLimit *string `json:"SubDomainLimit,omitempty" name:"SubDomainLimit"`
	// 已用主域名数

	UsedMainDomain *int64 `json:"UsedMainDomain,omitempty" name:"UsedMainDomain"`
	// 已用子域名数

	UsedSubDomain *int64 `json:"UsedSubDomain,omitempty" name:"UsedSubDomain"`
	// 过期时间

	ValidTime *string `json:"ValidTime,omitempty" name:"ValidTime"`
	// AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 客户名称

	ClientName *string `json:"ClientName,omitempty" name:"ClientName"`
	// 客户用户名（登陆用户名）

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type ModifyHostModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TCERuleDeleteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TCERuleDeleteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TCERuleDeleteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetNgWafVipConfRequest struct {
	*tchttp.BaseRequest

	// VIP配置

	VipConf *WafQueryVipConf `json:"VipConf,omitempty" name:"VipConf"`
}

func (r *WafSetNgWafVipConfRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetNgWafVipConfRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafUpsertSysLogRequest struct {
	*tchttp.BaseRequest

	// ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 开关状态 1-开启 0-关闭（默认）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 日志类型 1-攻击日志  2-访问日志  目前仅支持攻击日志

	LogType *int64 `json:"LogType,omitempty" name:"LogType"`
	// 网络协议 1-TCP 2-UDP 3-TLS

	Protocol *int64 `json:"Protocol,omitempty" name:"Protocol"`
	// 消息协议 1-RFC3164 2-RFC5424

	MsgProtocol *int64 `json:"MsgProtocol,omitempty" name:"MsgProtocol"`
	// syslog服务端地址

	LogHost *string `json:"LogHost,omitempty" name:"LogHost"`
	// syslog服务端端口

	LogPort *int64 `json:"LogPort,omitempty" name:"LogPort"`
}

func (r *WafUpsertSysLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafUpsertSysLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserSignatureRule struct {
	// 特征ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 规则开关

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 主类ID

	MainClassID *string `json:"MainClassID,omitempty" name:"MainClassID"`
	// 子类ID

	SubClassID *string `json:"SubClassID,omitempty" name:"SubClassID"`
	// CveID

	CveID *string `json:"CveID,omitempty" name:"CveID"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 主类名字，根据Language字段输出中文/英文

	MainClassName *string `json:"MainClassName,omitempty" name:"MainClassName"`
	// 子类名字，根据Language字段输出中文/英文，若子类id为00000000，此字段为空

	SubClassName *string `json:"SubClassName,omitempty" name:"SubClassName"`
	// 规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type WafGetAttackDetailResData struct {
	// 游标

	Context *string `json:"Context,omitempty" name:"Context"`
	// Count

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 表单数据

	Rows []*WafGetAttackDetailResRows `json:"Rows,omitempty" name:"Rows"`
	// 表单总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type WafGetNgWafHostsResLb struct {
	// lb创建时间

	LbCreateTime *string `json:"LbCreateTime,omitempty" name:"LbCreateTime"`
	// lb域名id

	LbDomainId *string `json:"LbDomainId,omitempty" name:"LbDomainId"`
	// lb监听id

	LbListenerId *string `json:"LbListenerId,omitempty" name:"LbListenerId"`
	// lb监听名称

	LbListenerName *string `json:"LbListenerName,omitempty" name:"LbListenerName"`
	// lb负载均衡名称

	LbLoadbalancerName *string `json:"LbLoadbalancerName,omitempty" name:"LbLoadbalancerName"`
	// lb修改时间

	LbModifyTime *string `json:"LbModifyTime,omitempty" name:"LbModifyTime"`
	// lb附件

	LbProtocol *string `json:"LbProtocol,omitempty" name:"LbProtocol"`
	// 区域

	LbRegion *string `json:"LbRegion,omitempty" name:"LbRegion"`
	// 状态

	LbStatus *string `json:"LbStatus,omitempty" name:"LbStatus"`
	// vip

	LbVip *string `json:"LbVip,omitempty" name:"LbVip"`
	// 端口

	LbVport *string `json:"LbVport,omitempty" name:"LbVport"`
	// 空间

	LbZone *string `json:"LbZone,omitempty" name:"LbZone"`
}

type WafGetWafDomainLoadbalancerRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
	// 负载均衡

	Loadbalancer *string `json:"Loadbalancer,omitempty" name:"Loadbalancer"`
	// vip域名

	VipDomain *string `json:"VipDomain,omitempty" name:"VipDomain"`
	// 监听

	Listener *string `json:"Listener,omitempty" name:"Listener"`
	// 地域

	ClientRegion *string `json:"ClientRegion,omitempty" name:"ClientRegion"`
}

func (r *WafGetWafDomainLoadbalancerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetWafDomainLoadbalancerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetDeleteHostRequest struct {
	*tchttp.BaseRequest

	// data数组

	Data []*WafQueryDelHostData `json:"Data,omitempty" name:"Data"`
}

func (r *WafSetDeleteHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetDeleteHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafQuerySort struct {
	// 建

	Key *string `json:"Key,omitempty" name:"Key"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type WafGetLicenseListRequest struct {
	*tchttp.BaseRequest
}

func (r *WafGetLicenseListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetLicenseListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetQueryHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetQueryHostResData `json:"Data,omitempty" name:"Data"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetQueryHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetQueryHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpUserSignaturePolicy struct {
	// appid

	OpAppId *uint64 `json:"OpAppId,omitempty" name:"OpAppId"`
	// 域名

	OpDomain *string `json:"OpDomain,omitempty" name:"OpDomain"`
	// 防护等级 300=标准 400=扩展

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 防护状态 0=关闭，1=开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type WafGetWafDomainLoadbalancerResRows struct {
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 监听器id

	LoadbalancerId *string `json:"LoadbalancerId,omitempty" name:"LoadbalancerId"`
	// 监听器名称

	LoadbalancerName *string `json:"LoadbalancerName,omitempty" name:"LoadbalancerName"`
	// 礼仪

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 端口号

	Vport *string `json:"Vport,omitempty" name:"Vport"`
	// 监听id

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 监听名称

	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 域名id

	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
}

type WafAddRsRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
}

func (r *WafAddRsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafAddRsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetBwIPRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// ip地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 用户id

	ClientAppid *int64 `json:"ClientAppid,omitempty" name:"ClientAppid"`
	// 请求名

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
}

func (r *WafGetBwIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetBwIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetClientQpsRequest struct {
	*tchttp.BaseRequest

	// 用户id

	ClientAppid *string `json:"ClientAppid,omitempty" name:"ClientAppid"`
	// 起始时间

	From *string `json:"From,omitempty" name:"From"`
	// 结束时间

	To *string `json:"To,omitempty" name:"To"`
}

func (r *WafGetClientQpsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetClientQpsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafQueryFilter struct {
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type WafSetNgWafVipConfResData struct {
	// 返回code

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 行数

	Rows *int64 `json:"Rows,omitempty" name:"Rows"`
}

type WafGetRulesRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
	// 查询过滤条件

	Filter *WafQueryFilter `json:"Filter,omitempty" name:"Filter"`
	// 查询条件

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 查询条件

	Search *string `json:"Search,omitempty" name:"Search"`
	// 查询条件

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *WafGetRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppIdQueryListObjects struct {
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// AppIdName

	Name *string `json:"Name,omitempty" name:"Name"`
	// AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// Token

	Token *string `json:"Token,omitempty" name:"Token"`
	// 插入时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type WafGetExportAttackLogJobsResData struct {
	// 返回码

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 饭回信息

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 信息集合

	Rows []*WafGetExportAttackLogJobsResRows `json:"Rows,omitempty" name:"Rows"`
	// 信息总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type WafSetUpgredeRuleByManuallyResData struct {
	// 信息

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DeleteOpRuleUpdateLogRequest struct {
	*tchttp.BaseRequest

	// 要删除的日志id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteOpRuleUpdateLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOpRuleUpdateLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TCERuleUpgradeRequest struct {
	*tchttp.BaseRequest

	// 操作者

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规则库id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *TCERuleUpgradeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TCERuleUpgradeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误码备注

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 数据

		Data *WafListResponseDataHostData `json:"Data,omitempty" name:"Data"`
		// 返回码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafAddTestClientResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *string `json:"Data,omitempty" name:"Data"`
		// CodeDesc

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafAddTestClientResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafAddTestClientResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetSystemVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetSystemVersionResData `json:"Data,omitempty" name:"Data"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetSystemVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetSystemVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetEditHostRequest struct {
	*tchttp.BaseRequest

	// 区域

	ClientRegion *string `json:"ClientRegion,omitempty" name:"ClientRegion"`
	// 状态

	State *int64 `json:"State,omitempty" name:"State"`
	// 是否cdn

	IsCdn *int64 `json:"IsCdn,omitempty" name:"IsCdn"`
	// 域名地址

	Host *WafQueryHost `json:"Host,omitempty" name:"Host"`
}

func (r *WafSetEditHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetEditHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNgWafHostsRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
}

func (r *GetNgWafHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNgWafHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TCERuleUploadRequest struct {
	*tchttp.BaseRequest

	// 操作者

	Name *string `json:"Name,omitempty" name:"Name"`
	// 上传特征库文件名

	RuleFileName *string `json:"RuleFileName,omitempty" name:"RuleFileName"`
}

func (r *TCERuleUploadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TCERuleUploadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditOpRuleUpdateLogStatusRequest struct {
	*tchttp.BaseRequest

	// 要编辑的日志id

	UpdateLogIDs []*uint64 `json:"UpdateLogIDs,omitempty" name:"UpdateLogIDs"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *EditOpRuleUpdateLogStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditOpRuleUpdateLogStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfServiceQueryListData struct {
	// 服务列表

	ServiceList []*ConfServiceQueryListServiceList `json:"ServiceList,omitempty" name:"ServiceList"`
}

type WafGetRuleUpgradeRecordResRows struct {
	// 登陆用户id

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 规则数

	Rules *string `json:"Rules,omitempty" name:"Rules"`
	// 规则版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 描述英文

	DescriptionEn *string `json:"DescriptionEn,omitempty" name:"DescriptionEn"`
}

type WafQueryVipConf struct {
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// ld的ip

	IdIps *string `json:"IdIps,omitempty" name:"IdIps"`
	// 评论

	Remarks *string `json:"Remarks,omitempty" name:"Remarks"`
}

type TCERuleRollbackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 回滚的规则库id

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TCERuleRollbackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TCERuleRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetNgWafVipConfResData struct {
	// 查询数据

	Rows []*WafGetNgWafVipConfResRows `json:"Rows,omitempty" name:"Rows"`
	// 查询总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type WafGetFreqRulesResData struct {
	// 表单数据

	Rows []*WafGetFreqRulesResRows `json:"Rows,omitempty" name:"Rows"`
	// 表单数据总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type WafGetVipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误码备注

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 数据

		Data *string `json:"Data,omitempty" name:"Data"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfServiceConfQueryListData struct {
	// 服务配置列表

	ServiceConfList []*ConfServiceConfQueryListServiceConfList `json:"ServiceConfList,omitempty" name:"ServiceConfList"`
}

type WafGetRuleUpgradeConfRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
	// 日期筛选

	Date *WafQueryDate `json:"Date,omitempty" name:"Date"`
	// 过滤

	Filter *WafQueryFilter `json:"Filter,omitempty" name:"Filter"`
	// 查询条件

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 查询条件

	Search *string `json:"Search,omitempty" name:"Search"`
	// 排序

	Sort *int64 `json:"Sort,omitempty" name:"Sort"`
}

func (r *WafGetRuleUpgradeConfRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetRuleUpgradeConfRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfNamespaceQueryListData struct {
	// 服务配置namespace数据集合

	NamespaceList []*ConfNamespaceQueryListNamespaceList `json:"NamespaceList,omitempty" name:"NamespaceList"`
}

type WafUpdateConfigSwitchResData struct {
	// 返回码

	Code *int64 `json:"Code,omitempty" name:"Code"`
}

type DescribeOpSignatureRuleRequest struct {
	*tchttp.BaseRequest

	// 排序字段，支持 signature_id, modify_time

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式，asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选条件，支持 MainClassID，SubClassID , Status, CveID, ID, Status; Status为必填字段 status=all，查询全部，status=0/1/2，查询对应status的规则

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
	// cn/en

	OpLanguage *string `json:"OpLanguage,omitempty" name:"OpLanguage"`
}

func (r *DescribeOpSignatureRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpSignatureRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpUserWhiteRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 规则列表

		List []*OpUserWhiteRuleNew `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpUserWhiteRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpUserWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetLatestRuleVersionRequest struct {
	*tchttp.BaseRequest
}

func (r *WafGetLatestRuleVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetLatestRuleVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetNonstandardPortResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafNonstandardPort `json:"Data,omitempty" name:"Data"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回码

		Code *string `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetNonstandardPortResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetNonstandardPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddGlobalWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 规则序号

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规则Id

	SignatureId *string `json:"SignatureId,omitempty" name:"SignatureId"`
	// 规则状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 匹配规则项列表

	Rules []*GlobalWhiteCond `json:"Rules,omitempty" name:"Rules"`
}

func (r *AddGlobalWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddGlobalWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafDelCustomPayloadsRequest struct {
	*tchttp.BaseRequest

	// 选择行id

	Items []*string `json:"Items,omitempty" name:"Items"`
}

func (r *WafDelCustomPayloadsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafDelCustomPayloadsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafDeleteNonstandardPortRequest struct {
	*tchttp.BaseRequest

	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *WafDeleteNonstandardPortRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafDeleteNonstandardPortRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetDeleteHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafSetDeleteHostResData `json:"Data,omitempty" name:"Data"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafSetDeleteHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetDeleteHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafSetEnableSysLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回码

		Code *string `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafSetEnableSysLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetEnableSysLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudDomainQueryListData struct {
	// 数据集合

	Objects []*CloudDomainQueryListObjects `json:"Objects,omitempty" name:"Objects"`
	// 页码

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 单页数据量

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type WafSetDeleteHostResData struct {
	// 返回码

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 返回信息

	Msg *string `json:"Msg,omitempty" name:"Msg"`
}

type WafSetDomainEngineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafSetDomainEngineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafSetDomainEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FiltersItemNew struct {
	// 字段名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤值

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 是否精确查找

	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type WafGetClientsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafListResponseClientData `json:"Data,omitempty" name:"Data"`
		// 返回码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetClientsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetClientsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetConfigSwitchInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *WafGetConfigSwitchInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetConfigSwitchInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetBanIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetClientsResData `json:"Data,omitempty" name:"Data"`
		// 返回信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetBanIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetBanIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetNgWafVipConfResRows struct {
	// 用户id

	Appid *string `json:"Appid,omitempty" name:"Appid"`
	// cmdbid

	CmdbId *string `json:"CmdbId,omitempty" name:"CmdbId"`
	// cpu限制

	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`
	// cpu可用

	CpuUsage *int64 `json:"CpuUsage,omitempty" name:"CpuUsage"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// l5

	L5 *string `json:"L5,omitempty" name:"L5"`
	// 等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 内存限制

	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`
	// 可用内存

	MemUsage *int64 `json:"MemUsage,omitempty" name:"MemUsage"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 评论

	Remarks *string `json:"Remarks,omitempty" name:"Remarks"`
	// 已用

	Used *string `json:"Used,omitempty" name:"Used"`
	// Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 地区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// pod信息

	PodInfos []*WafCLBPodInfo `json:"PodInfos,omitempty" name:"PodInfos"`
}

type WafGetInstancesRequest struct {
	*tchttp.BaseRequest

	// 分页参数

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
	// 模糊查询条件

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *WafGetInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetSystemVersionRequest struct {
	*tchttp.BaseRequest
}

func (r *WafGetSystemVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetSystemVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 分页起始

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页结束

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页面筛选用的appid

	CustomAppid *string `json:"CustomAppid,omitempty" name:"CustomAppid"`
}

func (r *DescribeCustomWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpUserWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 排序字段，支持 id，signature_id

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选字段支持 AppId，Domain

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeOpUserWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpUserWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetUpgredeProgressResRows struct {
	// 整体状态

	CompleteStatus *int64 `json:"CompleteStatus,omitempty" name:"CompleteStatus"`
	// 整体时间

	CompleteTime *string `json:"CompleteTime,omitempty" name:"CompleteTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 当前状态

	CurrentStatus *int64 `json:"CurrentStatus,omitempty" name:"CurrentStatus"`
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 操作描述

	OperationDesc *string `json:"OperationDesc,omitempty" name:"OperationDesc"`
	// 操作状态

	OperationType *int64 `json:"OperationType,omitempty" name:"OperationType"`
	// 行id

	RowId *int64 `json:"RowId,omitempty" name:"RowId"`
	// 规则版本

	RuleVersion *string `json:"RuleVersion,omitempty" name:"RuleVersion"`
}

type WafListResponseData struct {
	// 数组

	Rows *WafListResponseDataRows `json:"Rows,omitempty" name:"Rows"`
}

type WafGetHostsResData struct {
	// 试试

	Rows []*WafGetHostsResRowData `json:"Rows,omitempty" name:"Rows"`
}

type DescribeOpMainClassResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主类的总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 主类的详细信息

		MainClass []*MainClassItem `json:"MainClass,omitempty" name:"MainClass"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpMainClassResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpMainClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafDelNgWafVipConfRequest struct {
	*tchttp.BaseRequest

	// vip配置

	VipConf *WafQueryVipConf `json:"VipConf,omitempty" name:"VipConf"`
}

func (r *WafDelNgWafVipConfRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafDelNgWafVipConfRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetCosCredentialsResCredentials struct {
	// 证书;

	Credentials *WafGetCosCredentialsResSecCredentials `json:"Credentials,omitempty" name:"Credentials"`
	// 截止时间

	Expiration *string `json:"Expiration,omitempty" name:"Expiration"`
	// 截止时间戳

	ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// 开始时间戳

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
}

type WafGetAttackDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetAttackDetailResData `json:"Data,omitempty" name:"Data"`
		// 错误码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetAttackDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetAttackDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetRuleUpgradeConfResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetRuleUpgradeConfResData `json:"Data,omitempty" name:"Data"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetRuleUpgradeConfResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetRuleUpgradeConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Paging struct {
	// 起始页

	PageIndex *int64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 单页数量

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

type WafGetFreqRulesRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 规则名

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 用户id

	ClientAppid *string `json:"ClientAppid,omitempty" name:"ClientAppid"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 请求名

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
}

func (r *WafGetFreqRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetFreqRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafListResponseDataRuleData `json:"Data,omitempty" name:"Data"`
		// 返回信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetVipRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
}

func (r *WafGetVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetVipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpUserSignatureRuleRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 地域信息

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序条件，支持signature_id, modify_time

	By *string `json:"By,omitempty" name:"By"`
	// 分片偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 页面显示个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选条件，支持MainClassID, SubClassID, CveID，Status，ID；ID为规则id, Status = 0/1

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
	// appid

	OpAppId *uint64 `json:"OpAppId,omitempty" name:"OpAppId"`
	// en/cn

	OpLanguage *string `json:"OpLanguage,omitempty" name:"OpLanguage"`
}

func (r *DescribeOpUserSignatureRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpUserSignatureRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfServiceInstanceQueryListServiceInstanceList struct {
	// 创建人

	CreateName *string `json:"CreateName,omitempty" name:"CreateName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 服务实例Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 服务器IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 备注

	Note *string `json:"Note,omitempty" name:"Note"`
	// 服务Id

	ServiceId *int64 `json:"ServiceId,omitempty" name:"ServiceId"`
	// 服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务器状态 0在线 1下线

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 服务器Id

	SvrMachineId *int64 `json:"SvrMachineId,omitempty" name:"SvrMachineId"`
	// Tcp端口

	TcpPort *string `json:"TcpPort,omitempty" name:"TcpPort"`
	// Udp端口

	UdpPort *string `json:"UdpPort,omitempty" name:"UdpPort"`
	// 更新人

	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type WafGetRuleUpgradeRecordRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
	// 日期查询

	Date *WafQueryDate `json:"Date,omitempty" name:"Date"`
	// 查询条件

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 查询条件

	Search *string `json:"Search,omitempty" name:"Search"`
	// 筛选条件

	Filter *WafQueryFilter `json:"Filter,omitempty" name:"Filter"`
	// 排序

	Sort *int64 `json:"Sort,omitempty" name:"Sort"`
}

func (r *WafGetRuleUpgradeRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetRuleUpgradeRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetSysLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafSysLogConfig `json:"Data,omitempty" name:"Data"`
		// 返回码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetSysLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetSysLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetNgWafHostsResData struct {
	// 数据集合

	Rows *WafGetNgWafHostsResRows `json:"Rows,omitempty" name:"Rows"`
	// 数据总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type WafGetSystemVersionResData struct {
	// 系统版本

	SystemVersion *string `json:"SystemVersion,omitempty" name:"SystemVersion"`
}

type GetTCERuleCosKeyRequest struct {
	*tchttp.BaseRequest

	// 可选参数，用于指定将文件上传到Cos的Key，留空则由后台随机生成

	Key *string `json:"Key,omitempty" name:"Key"`
	// 前端页面url地址，用于cos跨域上传，传空的话后台尝试从http header的Origin字段获取

	Origin *string `json:"Origin,omitempty" name:"Origin"`
	// 登录人

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *GetTCERuleCosKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTCERuleCosKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafAddVipRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
}

func (r *WafAddVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafAddVipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetCheckDomainResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafSetSwitchCustomRuleStatusResData `json:"Data,omitempty" name:"Data"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetCheckDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetCheckDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetWafDomainLoadbalancerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetWafDomainLoadbalancerResData `json:"Data,omitempty" name:"Data"`
		// 返回码描述

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetWafDomainLoadbalancerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetWafDomainLoadbalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafVipUpdateRsRequest struct {
	*tchttp.BaseRequest

	// 分页

	Paging *Paging `json:"Paging,omitempty" name:"Paging"`
}

func (r *WafVipUpdateRsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafVipUpdateRsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGlobalWhiteRuleRequest struct {
	*tchttp.BaseRequest

	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 分页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页容量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序参数

	By *string `json:"By,omitempty" name:"By"`
	// 支持 Status，SignatureId（规则id），Target, TargetValue (Target 支持字段为Host，URI，FullUri，Parameter，Cookie，TargetValue为对应的筛选值

	Filters []*FiltersItemNew `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeGlobalWhiteRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGlobalWhiteRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpAttackWhiteRuleNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 规则列表

		List []*OpUserWhiteRuleNew `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpAttackWhiteRuleNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOpAttackWhiteRuleNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditOpRuleUpdateLogRequest struct {
	*tchttp.BaseRequest

	// 要编辑的日志id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// version

	LogVersion *string `json:"LogVersion,omitempty" name:"LogVersion"`
}

func (r *EditOpRuleUpdateLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditOpRuleUpdateLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetIspsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误码备注

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 数据

		Data *string `json:"Data,omitempty" name:"Data"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetIspsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetIspsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppIdQueryListData struct {
	// AppId列表

	Objects []*AppIdQueryListObjects `json:"Objects,omitempty" name:"Objects"`
	// 页码

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 页大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type CommonResponseData struct {
	// 返回Id，可选，用于公共无返回

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// ServiceId

	ServiceId *int64 `json:"ServiceId,omitempty" name:"ServiceId"`
}

type ProductUsages struct {
	// 计费项ID（磐石产品ID），全局唯一ID

	ProductId *uint64 `json:"ProductId,omitempty" name:"ProductId"`
	// 计费项名称

	ProductIdDescribe *string `json:"ProductIdDescribe,omitempty" name:"ProductIdDescribe"`
	// 产品用量单位

	UsageUnit *string `json:"UsageUnit,omitempty" name:"UsageUnit"`
	// 产品用量

	UsageValue *uint64 `json:"UsageValue,omitempty" name:"UsageValue"`
}

type WafGetCosCredentialsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafGetCosCredentialsResData `json:"Data,omitempty" name:"Data"`
		// 返回信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetCosCredentialsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetCosCredentialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WafGetInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公共出参

		Data *WafInstanceData `json:"Data,omitempty" name:"Data"`
		// 返回码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 返回码信息

		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WafGetInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WafGetInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
