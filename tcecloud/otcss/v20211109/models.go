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

package v20211109

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DescribeRiskSyscallWhiteListsExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载链接

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskSyscallWhiteListsExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallWhiteListsExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyEventListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>EventType- String - 是否必填：否 -事件类型 EVENT_RISK:风险事件类型，EVENT_PRIVILEGE:特权。</li>
	// <li>DealBehavior- String - 是否必填：否 - 执行动作,BEHAVIOR_ALERT:告警，BEHAVIOR_HOLDUP_SUCCESSED:拦截。</li>
	// <li>RuleName- string - 是否必填：否 - 规则名称。</li>
	// <li>NodeName- string - 是否必填：否 - 节点名称。</li>
	// <li>NodeIP- string - 是否必填：否 - 内外IP。</li>
	// <li>PublicIP- string - 是否必填：否 - 外网IP。</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	// <li>TimeRange- String -是否必填: 否 -  时间范围，第一个值表示开始时间，第二个值表示结束时间 </li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：告警数量：EventCount，最近生成时间：LatestFoundTime

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeImageDenyEventListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyEventListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSupportRegionsRequest struct {
	*tchttp.BaseRequest

	// 1表示获取用户资产相关的地域列表。
	// 2表示获取所有支持的地域列表。

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSupportRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSupportRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulIgnoreRegistryImageListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 过滤条件:AppId 租户id

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVulIgnoreRegistryImageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulIgnoreRegistryImageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterWorkloadWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检测项名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// workload的白名单列表

		WhiteListWorkloads []*WhiteListWorkload `json:"WhiteListWorkloads,omitempty" name:"WhiteListWorkloads"`
		// workload类型白名单总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterWorkloadWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterWorkloadWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterAssetSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeClusterAssetSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterAssetSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallWhiteListDetailRequest struct {
	*tchttp.BaseRequest

	// 白名单ID

	WhiteListId *string `json:"WhiteListId,omitempty" name:"WhiteListId"`
	// 任务ID

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeRiskSyscallWhiteListDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallWhiteListDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogStorageStatisticRequest struct {
	*tchttp.BaseRequest

	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeLogStorageStatisticRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogStorageStatisticRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDetailRequest struct {
	*tchttp.BaseRequest

	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 当前appID

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeVulDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterCheckItem struct {
	// 唯一的检测项的ID

	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`
	// 风险项的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 检测项详细描述。

	ItemDetail *string `json:"ItemDetail,omitempty" name:"ItemDetail"`
	// 威胁等级。严重Serious,高危High,中危Middle,提示Hint

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 检查对象、风险对象.Runc,Kubelet,Containerd,Pods

	RiskTarget *string `json:"RiskTarget,omitempty" name:"RiskTarget"`
	// 风险类别,漏洞风险CVERisk,配置风险ConfigRisk

	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`
	// 检测项所属的风险类型,权限提升:PrivilegePromotion,拒绝服务:RefuseService,目录穿越:DirectoryEscape,未授权访问:UnauthorizedAccess,权限许可和访问控制问题:PrivilegeAndAccessControl,敏感信息泄露:SensitiveInfoLeak

	RiskAttribute *string `json:"RiskAttribute,omitempty" name:"RiskAttribute"`
	// 风险特征,Tag.存在EXP:ExistEXP,存在POD:ExistPOC,无需重启:NoNeedReboot, 服务重启:ServerRestart,远程信息泄露:RemoteInfoLeak,远程拒绝服务:RemoteRefuseService,远程利用:RemoteExploit,远程执行:RemoteExecute

	RiskProperty *string `json:"RiskProperty,omitempty" name:"RiskProperty"`
	// CVE编号

	CVENumber *string `json:"CVENumber,omitempty" name:"CVENumber"`
	// 披露时间

	DiscoverTime *string `json:"DiscoverTime,omitempty" name:"DiscoverTime"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// CVSS信息,用于画图

	CVSS *string `json:"CVSS,omitempty" name:"CVSS"`
	// CVSS分数

	CVSSScore *string `json:"CVSSScore,omitempty" name:"CVSSScore"`
	// 参考连接

	RelateLink *string `json:"RelateLink,omitempty" name:"RelateLink"`
	// 影响类型，为Node或者Workload

	AffectedType *string `json:"AffectedType,omitempty" name:"AffectedType"`
	// 受影响的版本信息

	AffectedVersion *string `json:"AffectedVersion,omitempty" name:"AffectedVersion"`
	// 忽略的资产数量

	IgnoredAssetNum *int64 `json:"IgnoredAssetNum,omitempty" name:"IgnoredAssetNum"`
	// 是否忽略该检测项

	IsIgnored *bool `json:"IsIgnored,omitempty" name:"IsIgnored"`
	// 受影响评估

	RiskAssessment *string `json:"RiskAssessment,omitempty" name:"RiskAssessment"`
}

type ClusterNamespaceInfo struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 标签

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
}

type DescribeAssetImageVirusListExportRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Name- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>RiskLevel - String - 是否必填：否 - 风险等级  1,2,3,4</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 列表支持字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
}

func (r *DescribeAssetImageVirusListExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageVirusListExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIngressForwardConfigRequest struct {
	*tchttp.BaseRequest

	// Ingress名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeIngressForwardConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIngressForwardConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageRisk struct {
	// 高危行为

	Behavior *uint64 `json:"Behavior,omitempty" name:"Behavior"`
	// 种类

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 解决方案

	InstructionContent *string `json:"InstructionContent,omitempty" name:"InstructionContent"`
}

type K8sApiAbnormalRuleScopeInfo struct {
	// 范围
	// 系统事件:
	// ANONYMOUS_ACCESS: 匿名访问
	// ABNORMAL_UA_REQ: 异常UA请求
	// ANONYMOUS_ABNORMAL_PERMISSION: 匿名用户权限异动
	// GET_CREDENTIALS: 凭据信息获取
	// MOUNT_SENSITIVE_PATH: 敏感路径挂载
	// COMMAND_RUN: 命令执行
	// PRIVILEGE_CONTAINER: 特权容器
	// EXCEPTION_CRONTAB_TASK: 异常定时任务
	// STATICS_POD: 静态pod创建
	// ABNORMAL_CREATE_POD: 异常pod创建
	// USER_DEFINED: 用户自定义

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 动作(RULE_MODE_ALERT: 告警 RULE_MODE_RELEASE:放行)

	Action *string `json:"Action,omitempty" name:"Action"`
	// 威胁等级 HIGH:高级 MIDDLE: 中级 LOW:低级 NOTICE:提示

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 开关状态(true:开 false:关) 适用于系统规则

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 是否被删除 适用于自定义规则入参

	IsDelete *bool `json:"IsDelete,omitempty" name:"IsDelete"`
}

type MaliciousConnectionRuleInfo struct {
	// 枚举：
	// IP: 表示ipv4或者ipv6
	// DOMAIN: 表示域名

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 自定义黑白名单的域名/IP

	Address *string `json:"Address,omitempty" name:"Address"`
	// 创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 规则ID

	RuleID *uint64 `json:"RuleID,omitempty" name:"RuleID"`
	// 当前appID

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type VulAffectedRegistryImageInfo struct {
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像版本

	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
	// 镜像命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 镜像地址

	ImageRepoAddress *string `json:"ImageRepoAddress,omitempty" name:"ImageRepoAddress"`
	// 组件列表

	ComponentList []*VulAffectedImageComponentInfo `json:"ComponentList,omitempty" name:"ComponentList"`
	// 是否为镜像的最新版本

	IsLatestImage *bool `json:"IsLatestImage,omitempty" name:"IsLatestImage"`
	// 内部镜像资产ID

	ImageAssetId *int64 `json:"ImageAssetId,omitempty" name:"ImageAssetId"`
}

type DescribeTaskResultSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 严重风险影响的节点数量,返回7天数据

		SeriousRiskNodeCount []*uint64 `json:"SeriousRiskNodeCount,omitempty" name:"SeriousRiskNodeCount"`
		// 高风险影响的节点的数量,返回7天数据

		HighRiskNodeCount []*uint64 `json:"HighRiskNodeCount,omitempty" name:"HighRiskNodeCount"`
		// 中风险检查项的节点数量,返回7天数据

		MiddleRiskNodeCount []*uint64 `json:"MiddleRiskNodeCount,omitempty" name:"MiddleRiskNodeCount"`
		// 提示风险检查项的节点数量,返回7天数据

		HintRiskNodeCount []*uint64 `json:"HintRiskNodeCount,omitempty" name:"HintRiskNodeCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskResultSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskResultSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessRulesExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// execle下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalProcessRulesExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessRulesExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤字段
	// IsAuthorized是否授权，取值全部all，未授权0，已授权1

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式，asc，desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 是否仅展示各repository最新的镜像, 默认为false

	OnlyShowLatest *bool `json:"OnlyShowLatest,omitempty" name:"OnlyShowLatest"`
}

func (r *DescribeAssetImageRegistryListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessRulesRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAbnormalProcessRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RiskDnsEventInfo struct {
	// 事件ID

	EventID *uint64 `json:"EventID,omitempty" name:"EventID"`
	// 事件类型，恶意域名请求：DOMAIN，恶意IP请求：IP

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 恶意请求域名/IP

	Address *string `json:"Address,omitempty" name:"Address"`
	// 容器ID

	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 隔离状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING

	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// "NODE_DESTROYED"      //节点已销毁
	// "CONTAINER_EXITED"    //容器已退出
	// "CONTAINER_DESTROYED" //容器已销毁
	// "SHARED_HOST"         // 容器与主机共享网络
	// "RESOURCE_LIMIT"      //隔离操作资源超限
	// "UNKNOW"              // 原因未知

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离操作来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 首次发现时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 事件状态
	// EVENT_UNDEAL： 待处理
	// EVENT_DEALED：已处理
	// EVENT_IGNORE： 已忽略
	// EVENT_ADD_WHITE：已加白

	EventStatus *string `json:"EventStatus,omitempty" name:"EventStatus"`
	// 恶意请求次数

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 事件描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 恶意IP所属城市

	City *string `json:"City,omitempty" name:"City"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机ID

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 内网IP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 外网IP

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 当前appID

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type DescribeAssetImageRegistryDetailRequest struct {
	*tchttp.BaseRequest

	// 仓库列表id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 过滤条件

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetImageRegistryDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlRuleDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运行时策略详细信息

		RuleDetail *AccessControlRuleInfo `json:"RuleDetail,omitempty" name:"RuleDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessControlRuleDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserPurchaseInfoExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserPurchaseInfoExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserPurchaseInfoExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulIgnoreRegistryImageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 镜像列表

		List []*VulIgnoreRegistryImage `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulIgnoreRegistryImageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulIgnoreRegistryImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 待处理事件个数

		UnhandleEventCount *uint64 `json:"UnhandleEventCount,omitempty" name:"UnhandleEventCount"`
		// 待处理高危事件个数

		UnhandleHighLevelEventCount *uint64 `json:"UnhandleHighLevelEventCount,omitempty" name:"UnhandleHighLevelEventCount"`
		// 待处理中危事件个数

		UnhandleMediumLevelEventCount *uint64 `json:"UnhandleMediumLevelEventCount,omitempty" name:"UnhandleMediumLevelEventCount"`
		// 待处理低危事件个数

		UnhandleLowLevelEventCount *uint64 `json:"UnhandleLowLevelEventCount,omitempty" name:"UnhandleLowLevelEventCount"`
		// 待处理提示级别事件个数

		UnhandleNoticeLevelEventCount *uint64 `json:"UnhandleNoticeLevelEventCount,omitempty" name:"UnhandleNoticeLevelEventCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeK8sApiAbnormalSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersNamespacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命名空间列表

		ClusterNamespaceList []*ClusterNamespaceInfo `json:"ClusterNamespaceList,omitempty" name:"ClusterNamespaceList"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClustersNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClustersNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type K8sApiAbnormalTendencyItem struct {
	// 匿名访问事件数

	AnonymousAccessCount *uint64 `json:"AnonymousAccessCount,omitempty" name:"AnonymousAccessCount"`
	// 异常UA请求事件数

	ExceptionUARequestCount *uint64 `json:"ExceptionUARequestCount,omitempty" name:"ExceptionUARequestCount"`
	// 匿名用户权限事件数

	AnonymousUserRightCount *uint64 `json:"AnonymousUserRightCount,omitempty" name:"AnonymousUserRightCount"`
	// 凭据信息获取事件数

	CredentialInformationObtainCount *uint64 `json:"CredentialInformationObtainCount,omitempty" name:"CredentialInformationObtainCount"`
	// 敏感数据挂载事件数

	SensitiveDataMountCount *uint64 `json:"SensitiveDataMountCount,omitempty" name:"SensitiveDataMountCount"`
	// 命令执行事件数

	CmdExecCount *uint64 `json:"CmdExecCount,omitempty" name:"CmdExecCount"`
	// 异常定时任务事件数

	AbnormalScheduledTaskCount *uint64 `json:"AbnormalScheduledTaskCount,omitempty" name:"AbnormalScheduledTaskCount"`
	// 静态Pod创建数

	StaticsPodCreateCount *uint64 `json:"StaticsPodCreateCount,omitempty" name:"StaticsPodCreateCount"`
	// 可疑容器创建数

	DoubtfulContainerCreateCount *uint64 `json:"DoubtfulContainerCreateCount,omitempty" name:"DoubtfulContainerCreateCount"`
	// 自定义规则事件数

	UserDefinedRuleCount *uint64 `json:"UserDefinedRuleCount,omitempty" name:"UserDefinedRuleCount"`
	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 特权容器事件数

	PrivilegeContainerCount *uint64 `json:"PrivilegeContainerCount,omitempty" name:"PrivilegeContainerCount"`
}

type DescribeVirusSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 最近的一次扫描任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 木马影响容器个数

		RiskContainerCnt *uint64 `json:"RiskContainerCnt,omitempty" name:"RiskContainerCnt"`
		// 待处理风险个数

		RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`
		// 病毒库更新时间

		VirusDataBaseModifyTime *string `json:"VirusDataBaseModifyTime,omitempty" name:"VirusDataBaseModifyTime"`
		// 木马影响容器个数较昨日增长

		RiskContainerIncrease *int64 `json:"RiskContainerIncrease,omitempty" name:"RiskContainerIncrease"`
		// 待处理风险个数较昨日增长

		RiskIncrease *int64 `json:"RiskIncrease,omitempty" name:"RiskIncrease"`
		// 隔离事件个数较昨日新增

		IsolateIncrease *int64 `json:"IsolateIncrease,omitempty" name:"IsolateIncrease"`
		// 隔离事件总数

		IsolateCnt *int64 `json:"IsolateCnt,omitempty" name:"IsolateCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogCleanSettingInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 触发清理的储量底线

		ReservesLimit *uint64 `json:"ReservesLimit,omitempty" name:"ReservesLimit"`
		// 清理停止时的储量截至线

		ReservesDeadline *uint64 `json:"ReservesDeadline,omitempty" name:"ReservesDeadline"`
		// 触发清理的存储天数

		DayLimit *uint64 `json:"DayLimit,omitempty" name:"DayLimit"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecLogCleanSettingInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogCleanSettingInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 风险详情数组

		ClusterRiskItems []*ClusterRiskItem `json:"ClusterRiskItems,omitempty" name:"ClusterRiskItems"`
		// 风险项的总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterNodeWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检测项名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 忽略的节点信息列表

		WhiteListNodes []*WhiteListNodeInfo `json:"WhiteListNodes,omitempty" name:"WhiteListNodes"`
		// 节点类型白名单总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterNodeWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterNodeWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWebVulExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为50000，最大值为50000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>AppId - String - 是否必填：是 - 租户id筛选</li>
	// <li>OnlyAffectedContainer- string - 是否必填：否 - 仅展示影响容器的漏洞true,false</li>
	// <li>OnlyAffectedNewestImage-string - 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CategoryType- string - 是否必填：否 - 漏洞子类型</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	// <li>FocusOnType - string - 是否必填：否 -关注紧急度类型 。ALL :全部，SERIOUS_LEVEL： 严重和高危 ，IS_SUGGEST： 重点关注，POC_EXP 有Poc或Exp ，NETWORK_EXP: 远程Exp</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateWebVulExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWebVulExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageSimpleListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Keywords- String - 是否必填：否 - 镜像名、镜像id 称筛选，</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetImageSimpleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageSimpleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EscapeRule struct {
	// 规则类型
	// ESCAPE_HOST_ACESS_FILE:宿主机文件访问逃逸
	//    ESCAPE_MOUNT_NAMESPACE:MountNamespace逃逸
	//    ESCAPE_PRIVILEDGE:程序提权逃逸
	//    ESCAPE_PRIVILEDGE_CONTAINER_START:特权容器启动逃逸
	//    ESCAPE_MOUNT_SENSITIVE_PTAH:敏感路径挂载
	// ESCAPE_SYSCALL:Syscall逃逸

	Type *string `json:"Type,omitempty" name:"Type"`
	// 规则名称
	// 宿主机文件访问逃逸、
	// Syscall逃逸、
	// MountNamespace逃逸、
	// 程序提权逃逸、
	// 特权容器启动逃逸、
	// 敏感路径挂载

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否打开：false否 ，true是

	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
	// 规则组别。RISK_CONTAINER：风险容器，PROCESS_PRIVILEGE：程序特权，CONTAINER_ESCAPE：容器逃逸

	Group *string `json:"Group,omitempty" name:"Group"`
	// 当前appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type CreateVulContainerExportJobRequest struct {
	*tchttp.BaseRequest

	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 需要返回的数量，默认为50000，最大值为50000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID</li>
	// <li>ContainerName- String -是否必填: 否 - 容器名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *CreateVulContainerExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulContainerExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageListExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageListExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecLogCleanSettingInfoRequest struct {
	*tchttp.BaseRequest

	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeSecLogCleanSettingInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecLogCleanSettingInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像列表

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 总数量

		List []*ImagesInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetPortListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 端口列表

		List []*PortInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetPortListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPortListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenTcssTrialRequest struct {
	*tchttp.BaseRequest

	// 用户uin

	CurrentUin *string `json:"CurrentUin,omitempty" name:"CurrentUin"`
	// 用户AppId

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
	// 使用结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 赠送的镜像授权数

	GivenAuthorizedCnt *uint64 `json:"GivenAuthorizedCnt,omitempty" name:"GivenAuthorizedCnt"`
}

func (r *OpenTcssTrialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenTcssTrialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>ImageName- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>ScanStatus - String - 是否必填：否 - 镜像扫描状态notScan，scanning，scanned，scanErr</li>
	// <li>ImageID- String - 是否必填：否 - 镜像ID筛选，</li>
	// <li>SecurityRisk- String - 是否必填：否 - 安全风险，VulCnt 、VirusCnt、RiskCnt、IsTrustImage</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateImageExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSearchTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 模板列表

		List []*SearchTemplate `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSearchTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSearchTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像风险列表

		List []*ImageRiskInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRiskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulIgnoreRegistryImage struct {
	// 记录ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 仓库名称

	RegistryName *string `json:"RegistryName,omitempty" name:"RegistryName"`
	// 镜像版本

	ImageVersion *string `json:"ImageVersion,omitempty" name:"ImageVersion"`
	// 仓库地址

	RegistryPath *string `json:"RegistryPath,omitempty" name:"RegistryPath"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 租户appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type DescribeSecEventsTendencyRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 当前AppId

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeSecEventsTendencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecEventsTendencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulImageExportJobRequest struct {
	*tchttp.BaseRequest

	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 需要返回的数量，默认为50000，最大值为50000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>AppId - String - 是否必填：是 - 租户id筛选</li>
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ClientIP- string -是否必填: 否 - 内网IP</li>
	// <li>PublicIP- string -是否必填: 否 - 外网IP</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>HostName- string -是否必填: 否 - 主机名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateVulImageExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulImageExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusEventTendencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 趋势列表

		List []*VirusTendencyInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusEventTendencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusEventTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterServiceListRequest struct {
	*tchttp.BaseRequest

	// Pod的名字,不填表示查询所有

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ServiceName, ExternalIp等

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 命名空间,不填表示查询所有

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群Id,不填表示查询所有

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterServiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterServiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryVirusListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像漏洞列表

		List []*ImageVirus `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryVirusListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryVirusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetHostListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机列表

		List []*HostInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetHostListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEmergencyVulExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为50000，最大值为50000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>AppId - String - 是否必填：是 - 租户id筛选</li>
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateEmergencyVulExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEmergencyVulExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespaceListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：Namespace

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNamespaceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNamespaceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellWhiteListsRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeReverseShellWhiteListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellWhiteListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserPodListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Pod列表详细信息

		PodList []*ClusterPodInfo `json:"PodList,omitempty" name:"PodList"`
		// Pod列表总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserPodListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserPodListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryVirusListExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像信息

	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`
	// 排序字段（Level）

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 + -

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 镜像标识Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryVirusListExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryVirusListExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAffectedNodeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 受影响的节点总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 受影响的节点列表

		AffectedNodeList []*AffectedNodeItem `json:"AffectedNodeList,omitempty" name:"AffectedNodeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAffectedNodeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAffectedNodeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespaceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命名空间列表

		Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNamespaceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNamespaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeEventTypeSummaryRequest struct {
	*tchttp.BaseRequest

	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeEscapeEventTypeSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventTypeSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlRulesExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// execle下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessControlRulesExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlRulesExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EscapeWhiteListInfo struct {
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 白名单记录ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 关联主机数量

	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
	// 关联容器数量

	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
	// 加白事件类型

	EventType []*string `json:"EventType,omitempty" name:"EventType"`
	// 创建时间

	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 镜像大小

	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`
	// 当前appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type DescribeVulDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞详情信息

		VulInfo *VulDetailInfo `json:"VulInfo,omitempty" name:"VulInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterServiceYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// yaml格式字符串,base64编码

		Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterServiceYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterServiceYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportJobResultRequest struct {
	*tchttp.BaseRequest

	// CreateExportComplianceStatusListJob返回的JobId字段的值

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 支持按appid过滤

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeExportJobResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportJobResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryRiskInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像高危列表

		List []*ImageRisk `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryRiskInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryRiskInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterRiskItem struct {
	// 检测项相关信息

	CheckItem *ClusterCheckItem `json:"CheckItem,omitempty" name:"CheckItem"`
	// 验证信息

	VerifyInfo *string `json:"VerifyInfo,omitempty" name:"VerifyInfo"`
	// 事件描述,检查的错误信息

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
	// 受影响的集群数量

	AffectedClusterCount *uint64 `json:"AffectedClusterCount,omitempty" name:"AffectedClusterCount"`
	// 受影响的节点数量

	AffectedNodeCount *uint64 `json:"AffectedNodeCount,omitempty" name:"AffectedNodeCount"`
}

type DescribeAffectedWorkloadListRequest struct {
	*tchttp.BaseRequest

	// 唯一的检测项的ID

	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：WorkloadType,ClusterId

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAffectedWorkloadListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAffectedWorkloadListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetContainerDetailRequest struct {
	*tchttp.BaseRequest

	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 租户id

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeAssetContainerDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetContainerDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceFilters struct {
	// 过滤键的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 是否模糊查询。默认为是。

	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type DescribeRiskTargetCountRequest struct {
	*tchttp.BaseRequest

	// 根据appid过滤

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRiskTargetCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskTargetCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageHostListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件 支持ImageID,HostID，增加AppId

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetImageHostListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageHostListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageVulListExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageVulListExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageVulListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessControlRuleInfo struct {
	// 策略id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 开关,true:开启，false:禁用

	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
	// 生效惊现id，空数组代表全部镜像

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
	// 子策略数组

	ChildRules []*AccessControlChildRuleInfo `json:"ChildRules,omitempty" name:"ChildRules"`
	// 策略名字

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// []

	SystemChildRules []*AccessControlSystemChildRuleInfo `json:"SystemChildRules,omitempty" name:"SystemChildRules"`
	// 是否是系统默认策略

	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

type ReverseShellEventInfo struct {
	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 容器名

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 生成时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 事件解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件详细描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 状态，EVENT_UNDEAL:事件未处理
	//     EVENT_DEALED:事件已经处理
	//     EVENT_INGNORE：事件已经忽略
	//     EVENT_ADD_WHITE：时间已经加白

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 父进程名

	PProcessName *string `json:"PProcessName,omitempty" name:"PProcessName"`
	// 事件数量

	EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// AppId值

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 目标地址

	DstAddress *string `json:"DstAddress,omitempty" name:"DstAddress"`
	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// 	"NODE_DESTROYED"      //节点已销毁
	// 	"CONTAINER_EXITED"    //容器已退出
	// 	"CONTAINER_DESTROYED" //容器已销毁
	// 	"SHARED_HOST"         // 容器与主机共享网络
	// 	"RESOURCE_LIMIT"      //隔离操作资源超限
	// 	"UNKNOW"              // 原因未知

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离操作来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING

	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
}

type DescribeAddImageDenyRuleInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAddImageDenyRuleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddImageDenyRuleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateK8sApiAbnormalRuleExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>RuleType - string  - 是否必填: 否 -规则类型</li>
	// <li>Status - string  - 是否必填: 否 -状态</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By []*string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateK8sApiAbnormalRuleExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateK8sApiAbnormalRuleExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemPolicyRequest struct {
	*tchttp.BaseRequest

	// RUNTIME_ESCAPE:  容器逃逸；RUNTIME_PROCESS  :  异常进程；RUNTIME_FILE  :  文件篡改；RUNTIME_SYSCALL  :  高危系统调用；RUNTIME_K8S_API : K8s API异常请求

	SystemType *string `json:"SystemType,omitempty" name:"SystemType"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤字段

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeSystemPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerAssetSummaryRequest struct {
	*tchttp.BaseRequest

	// 租户id，当前查询的租户Id

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeContainerAssetSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerAssetSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetClusterListItem struct {
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群状态
	// CSR_RUNNING: 运行中
	// CSR_EXCEPTION:异常
	// CSR_DEL:已经删除

	Status *string `json:"Status,omitempty" name:"Status"`
	// 绑定规则名称

	BindRuleName *string `json:"BindRuleName,omitempty" name:"BindRuleName"`
	// 集群类型:
	// CT_TKE: TKE集群
	// CT_USER_CREATE: 用户自建集群

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
}

type ClusterServiceInfo struct {
	// service名称.

	Name *string `json:"Name,omitempty" name:"Name"`
	// service类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// Selector信息

	Selectors *string `json:"Selectors,omitempty" name:"Selectors"`
	// Selector数量

	SelectorCount *uint64 `json:"SelectorCount,omitempty" name:"SelectorCount"`
	// 负载均衡Ip

	ExternalIp *string `json:"ExternalIp,omitempty" name:"ExternalIp"`
	// 服务Ip

	ClusterIp *string `json:"ClusterIp,omitempty" name:"ClusterIp"`
	// string类型,以,分隔

	Port *string `json:"Port,omitempty" name:"Port"`
	// 所属集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 所属集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 关联Pod名字

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 关联pod数量

	PodCount *uint64 `json:"PodCount,omitempty" name:"PodCount"`
	// 创建时间

	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 标签

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 集群状态

	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
	// 集群类型

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// service唯一标识符

	ServiceUniqueID *string `json:"ServiceUniqueID,omitempty" name:"ServiceUniqueID"`
}

type DescribeClusterAssetSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户的pod总量

		PodTotalCount *uint64 `json:"PodTotalCount,omitempty" name:"PodTotalCount"`
		// 用户运行状态的Pod数量

		RunningPodCount *uint64 `json:"RunningPodCount,omitempty" name:"RunningPodCount"`
		// Pending状态的Pod数量

		PendingPodCount *uint64 `json:"PendingPodCount,omitempty" name:"PendingPodCount"`
		// 用户的集群Service总量

		ServiceTotalCount *uint64 `json:"ServiceTotalCount,omitempty" name:"ServiceTotalCount"`
		// ClusterIp类型的Service数量

		ClusterIpCount *uint64 `json:"ClusterIpCount,omitempty" name:"ClusterIpCount"`
		// NodePort类型的Service数量

		NodePortCount *uint64 `json:"NodePortCount,omitempty" name:"NodePortCount"`
		// 用户的集群Ingress总量

		IngressTotalCount *uint64 `json:"IngressTotalCount,omitempty" name:"IngressTotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterAssetSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterAssetSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHostExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHostExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHostExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterServiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群Service列表信息

		ClusterServiceList []*ClusterServiceInfo `json:"ClusterServiceList,omitempty" name:"ClusterServiceList"`
		// 集群Service列表总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterServiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CompliancePolicyItemSummary struct {
	// 为客户分配的唯一的检测项的ID。

	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`
	// 检测项的原始ID。

	BasePolicyItemId *uint64 `json:"BasePolicyItemId,omitempty" name:"BasePolicyItemId"`
	// 检测项的名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 检测项所属的类型，枚举字符串。

	Category *string `json:"Category,omitempty" name:"Category"`
	// 所属的合规标准

	BenchmarkStandardName *string `json:"BenchmarkStandardName,omitempty" name:"BenchmarkStandardName"`
	// 威胁等级。RISK_CRITICAL, RISK_HIGH, RISK_MEDIUM, RISK_LOW, RISK_NOTICE。

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 检测项所属的资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 最近检测的时间

	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`
	// 检测状态
	//
	// CHECK_INIT, 待检测
	//
	// CHECK_RUNNING, 检测中
	//
	// CHECK_FINISHED, 检测完成
	//
	// CHECK_FAILED, 检测失败

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 检测结果。RESULT_PASSED: 通过

	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`
	// 通过检测的资产的数目

	PassedAssetCount *uint64 `json:"PassedAssetCount,omitempty" name:"PassedAssetCount"`
	// 未通过检测的资产的数目

	FailedAssetCount *uint64 `json:"FailedAssetCount,omitempty" name:"FailedAssetCount"`
	// 检测项对应的白名单项的ID。如果存在且非0，表示检测项被用户忽略。

	WhitelistId *uint64 `json:"WhitelistId,omitempty" name:"WhitelistId"`
	// 处理建议。

	FixSuggestion *string `json:"FixSuggestion,omitempty" name:"FixSuggestion"`
	// 所属的合规标准的ID

	BenchmarkStandardId *uint64 `json:"BenchmarkStandardId,omitempty" name:"BenchmarkStandardId"`
	// 应用版本

	ApplicableVersion *string `json:"ApplicableVersion,omitempty" name:"ApplicableVersion"`
}

type DescribeEscapeEventTypeSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 容器逃逸事件数

		ContainerEscapeEventCount *int64 `json:"ContainerEscapeEventCount,omitempty" name:"ContainerEscapeEventCount"`
		// 程序提权事件数

		ProcessPrivilegeEventCount *int64 `json:"ProcessPrivilegeEventCount,omitempty" name:"ProcessPrivilegeEventCount"`
		// 风险容器事件数

		RiskContainerEventCount *int64 `json:"RiskContainerEventCount,omitempty" name:"RiskContainerEventCount"`
		// 逃逸事件待处理数

		PendingEscapeEventCount *int64 `json:"PendingEscapeEventCount,omitempty" name:"PendingEscapeEventCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEscapeEventTypeSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventTypeSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceTaskPolicyItemSummaryListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回最近一次合规检查任务的ID。这个任务为本次所展示数据的来源。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 返回检测项的总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回各检测项对应的汇总信息的列表。

		PolicyItemSummaryList []*CompliancePolicyItemSummary `json:"PolicyItemSummaryList,omitempty" name:"PolicyItemSummaryList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComplianceTaskPolicyItemSummaryListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceTaskPolicyItemSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsEventDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件ID

		EventID *uint64 `json:"EventID,omitempty" name:"EventID"`
		// 事件类型，恶意域名请求：DOMAIN，恶意IP请求：IP

		EventType *string `json:"EventType,omitempty" name:"EventType"`
		// 恶意请求次数

		EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
		// 首次发现时间

		FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
		// 最近生成时间

		LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
		// 容器ID

		ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`
		// 容器名称

		ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
		// 隔离状态
		// 未隔离  	NORMAL
		// 已隔离		ISOLATED
		// 隔离中		ISOLATING
		// 隔离失败	ISOLATE_FAILED
		// 解除隔离中  RESTORING
		// 解除隔离失败 RESTORE_FAILED

		ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
		// 容器状态
		// 正在运行: RUNNING
		// 暂停: PAUSED
		// 停止: STOPPED
		// 已经创建: CREATED
		// 已经销毁: DESTROYED
		// 正在重启中: RESTARTING
		// 迁移中: REMOVING

		ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
		// 容器子状态
		// "AGENT_OFFLINE"       //Agent离线
		// "NODE_DESTROYED"      //节点已销毁
		// "CONTAINER_EXITED"    //容器已退出
		// "CONTAINER_DESTROYED" //容器已销毁
		// "SHARED_HOST"         // 容器与主机共享网络
		// "RESOURCE_LIMIT"      //隔离操作资源超限
		// "UNKNOW"              // 原因未知

		ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
		// 容器隔离操作来源

		ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
		// 镜像ID

		ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
		// 镜像名称

		ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
		// 主机名称

		HostName *string `json:"HostName,omitempty" name:"HostName"`
		// 内网IP

		HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
		// 外网IP

		PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
		// 节点名称

		PodName *string `json:"PodName,omitempty" name:"PodName"`
		// 事件描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// 解决方案

		Solution *string `json:"Solution,omitempty" name:"Solution"`
		// 参考链接

		Reference []*string `json:"Reference,omitempty" name:"Reference"`
		// 恶意域名或IP

		Address *string `json:"Address,omitempty" name:"Address"`
		// 恶意IP所属城市

		City *string `json:"City,omitempty" name:"City"`
		// 命中规则类型
		// SYSTEM：系统规则
		//  USER：用户自定义

		MatchRuleType *string `json:"MatchRuleType,omitempty" name:"MatchRuleType"`
		// 标签特征

		FeatureLabel *string `json:"FeatureLabel,omitempty" name:"FeatureLabel"`
		// 进程权限

		ProcessAuthority *string `json:"ProcessAuthority,omitempty" name:"ProcessAuthority"`
		// 进程md5

		ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`
		// 进程启动用户

		ProcessStartUser *string `json:"ProcessStartUser,omitempty" name:"ProcessStartUser"`
		// 进程用户组

		ProcessUserGroup *string `json:"ProcessUserGroup,omitempty" name:"ProcessUserGroup"`
		// 进程路径

		ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
		// 进程树

		ProcessTree *string `json:"ProcessTree,omitempty" name:"ProcessTree"`
		// 进程命令行参数

		ProcessParam *string `json:"ProcessParam,omitempty" name:"ProcessParam"`
		// 父进程启动用户

		ParentProcessStartUser *string `json:"ParentProcessStartUser,omitempty" name:"ParentProcessStartUser"`
		// 父进程用户组

		ParentProcessUserGroup *string `json:"ParentProcessUserGroup,omitempty" name:"ParentProcessUserGroup"`
		// 父进程路径

		ParentProcessPath *string `json:"ParentProcessPath,omitempty" name:"ParentProcessPath"`
		// 父进程命令行参数

		ParentProcessParam *string `json:"ParentProcessParam,omitempty" name:"ParentProcessParam"`
		// 祖先进程启动用户

		AncestorProcessStartUser *string `json:"AncestorProcessStartUser,omitempty" name:"AncestorProcessStartUser"`
		// 祖先进程用户组

		AncestorProcessUserGroup *string `json:"AncestorProcessUserGroup,omitempty" name:"AncestorProcessUserGroup"`
		// 祖先进程路径

		AncestorProcessPath *string `json:"AncestorProcessPath,omitempty" name:"AncestorProcessPath"`
		// 祖先进程命令行参数

		AncestorProcessParam *string `json:"AncestorProcessParam,omitempty" name:"AncestorProcessParam"`
		// 主机ID

		HostID *string `json:"HostID,omitempty" name:"HostID"`
		// 事件状态
		// EVENT_UNDEAL： 待处理
		// EVENT_DEALED：已处理
		// EVENT_IGNORE： 已忽略
		// EVENT_ADD_WHITE：已加白

		EventStatus *string `json:"EventStatus,omitempty" name:"EventStatus"`
		// 操作时间

		OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
		// 备注

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskDnsEventDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AbnormalProcessEventTendencyInfo struct {
	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 待处理代理软件事件数

	ProxyToolEventCount *int64 `json:"ProxyToolEventCount,omitempty" name:"ProxyToolEventCount"`
	// 待处理横向参透事件数

	TransferControlEventCount *int64 `json:"TransferControlEventCount,omitempty" name:"TransferControlEventCount"`
	// 待处理恶意命令事件数

	AttackCmdEventCount *int64 `json:"AttackCmdEventCount,omitempty" name:"AttackCmdEventCount"`
	// 待处理反弹shell事件数

	ReverseShellEventCount *int64 `json:"ReverseShellEventCount,omitempty" name:"ReverseShellEventCount"`
	// 待处理无文件程序执行事件数

	FilelessEventCount *int64 `json:"FilelessEventCount,omitempty" name:"FilelessEventCount"`
	// 待处理高危命令事件数

	RiskCmdEventCount *int64 `json:"RiskCmdEventCount,omitempty" name:"RiskCmdEventCount"`
	// 待处理敏感服务异常子进程启动事件数

	AbnormalChildProcessEventCount *int64 `json:"AbnormalChildProcessEventCount,omitempty" name:"AbnormalChildProcessEventCount"`
	// 待处理自定义规则事件数

	UserDefinedRuleEventCount *int64 `json:"UserDefinedRuleEventCount,omitempty" name:"UserDefinedRuleEventCount"`
}

type DescribeReverseShellDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件基本信息

		EventBaseInfo *RunTimeEventBaseInfo `json:"EventBaseInfo,omitempty" name:"EventBaseInfo"`
		// 进程信息

		ProcessInfo *ProcessDetailInfo `json:"ProcessInfo,omitempty" name:"ProcessInfo"`
		// 父进程信息

		ParentProcessInfo *ProcessDetailBaseInfo `json:"ParentProcessInfo,omitempty" name:"ParentProcessInfo"`
		// 事件描述

		EventDetail *ReverseShellEventDescription `json:"EventDetail,omitempty" name:"EventDetail"`
		// 祖先进程信息

		AncestorProcessInfo *ProcessBaseInfo `json:"AncestorProcessInfo,omitempty" name:"AncestorProcessInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulContainerExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVulContainerExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulContainerExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeEventsExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeEscapeEventsExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventsExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageHostListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像列表

		List []*ImageHost `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageHostListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageVulListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>AppId - String - 是否必填：是 - 租户id筛选</li>
	// <li>Name- String - 是否必填：否 - 漏洞名称名称筛选，</li>
	// <li>Level - String - 是否必填：否 - 风险等级  1,2,3,4</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 排序字段（Level）

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 + -

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetImageVulListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageVulListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryRiskListExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryRiskListExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryRiskListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VirusTendencyInfo struct {
	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 待处理事件总数

	PendingEventCount *uint64 `json:"PendingEventCount,omitempty" name:"PendingEventCount"`
	// 风险容器总数

	RiskContainerCount *uint64 `json:"RiskContainerCount,omitempty" name:"RiskContainerCount"`
	// 事件总数

	EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
	// 隔离事件总数

	IsolateEventCount *uint64 `json:"IsolateEventCount,omitempty" name:"IsolateEventCount"`
}

type DescribeAbnormalProcessEventsExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	// <li>AppId - String - 是否必填：否 - 租户id筛选</li>
	// <li>FilterTimeRange - String[] - 是否必填：否 - 筛选时间，第一个值为开始时间，第二个值为结束时间["2020-1-1 12:00:00", "2020-1-1 12:00:00"]</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeAbnormalProcessEventsExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessEventsExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellEventsRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeReverseShellEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WhiteListWorkload struct {
	// 加入白名单的时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 白名单的workloadId

	WorkloadId *uint64 `json:"WorkloadId,omitempty" name:"WorkloadId"`
	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 工作负载名称

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
	// 工作负载类型

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeReverseShellWhiteListsExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载链接

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellWhiteListsExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellWhiteListsExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallEventsExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Excel下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 任务ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskSyscallEventsExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallEventsExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessControlEventDescription struct {
	// 事件规则

	Description *string `json:"Description,omitempty" name:"Description"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 命中规则详细信息

	MatchRule *AccessControlChildRuleInfo `json:"MatchRule,omitempty" name:"MatchRule"`
	// 命中规则名字

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 命中规则id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 操作时间

	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
}

type ModifyEventEscapeImageStatusRequest struct {
	*tchttp.BaseRequest

	// 标记事件的状态:
	// EVENT_DEALED:事件处理
	// EVENT_IGNORE"：事件忽略
	// EVENT_DEL:事件删除
	// EVENT_ADD_WHITE:事件加白

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 唯一值

	UniqueKeys []*string `json:"UniqueKeys,omitempty" name:"UniqueKeys"`
	// 加白镜像id

	ImageIdSet []*string `json:"ImageIdSet,omitempty" name:"ImageIdSet"`
	// 加白事件类型

	EventType []*string `json:"EventType,omitempty" name:"EventType"`
}

func (r *ModifyEventEscapeImageStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEventEscapeImageStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetClusterListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群列表

		List []*AssetClusterListItem `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetClusterListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetClusterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventEscapeImageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 风险容器镜像列表

		List []*EventEscapeImageInfo `json:"List,omitempty" name:"List"`
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEventEscapeImageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventEscapeImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRiskListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 风险级别 1,2,3,4，</li>
	// <li>Behavior - String - 是否必填：否 - 风险行为 1,2,3,4</li>
	// <li>Type - String - 是否必填：否 - 风险类型  1,2,</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetImageRiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePodContainersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// pod的容器列表

		PodContainerList []*PodContainerInfo `json:"PodContainerList,omitempty" name:"PodContainerList"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePodContainersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePodContainersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnfinishRefreshTaskRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUnfinishRefreshTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnfinishRefreshTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 异常进程数组

		EventSet []*AbnormalProcessEventInfo `json:"EventSet,omitempty" name:"EventSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalProcessEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeESHitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ES查询结果JSON

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeESHitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeESHitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRiskRequest struct {
	*tchttp.BaseRequest

	// 漏洞风险CVERisk,配置风险ConfigRisk

	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`
	// 支持按照appid过滤

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeClusterRiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterRiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusEventTendencyRequest struct {
	*tchttp.BaseRequest

	// 趋势周期(默认为7天)

	TendencyPeriod *uint64 `json:"TendencyPeriod,omitempty" name:"TendencyPeriod"`
	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeVirusEventTendencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusEventTendencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateComponentExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10000，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>ComponentName- String - 是否必填：否 - 镜像组件名称</li><li>ComponentVersion- String - 是否必填：否 - 镜像组件版本</li><li>ComponentType- String - 是否必填：否 - 镜像组件类型</li><li>VulLevel- String - 是否必填：否 - 漏洞威胁等级</li><li>HasVul- String - 是否必填：否 -是否有漏洞；true：是，false，否；不传或ALL ：全部</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式desc ，asc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateComponentExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateComponentExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalRuleScopeListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 过滤条件。
	// <li>Action - string -是否必填: 否 - 执行动作</li>
	// <li>RiskLevel - string  - 是否必填: 否 -威胁等级筛选</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeK8sApiAbnormalRuleScopeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalRuleScopeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalRuleScopeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 列表

		List []*K8sApiAbnormalRuleScopeInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeK8sApiAbnormalRuleScopeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalRuleScopeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousConnectionWhiteListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>RequestType- string - 是否必填：是 - 请求类型，全部请求类型：ALL；域名：DOMAIN；IP: IP</li>
	// <li>WhiteDomain- string - 是否必填：否 - 自定义白域名</li>
	// <li>WhiteIP- string - 是否必填：否 - 自定义白名单IP</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeMaliciousConnectionWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaliciousConnectionWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSystemVulExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为50000，最大值为50000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>AppId - String - 是否必填：是 - 租户id筛选</li>
	// <li>OnlyAffectedContainer- string - 是否必填：否 - 仅展示影响容器的漏洞true,false</li>
	// <li>OnlyAffectedNewestImage-string - 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CategoryType- string - 是否必填：否 - 漏洞子类型</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	// <li>FocusOnType - string - 是否必填：否 -关注紧急度类型 。ALL :全部，SERIOUS_LEVEL： 严重和高危 ，IS_SUGGEST： 重点关注，POC_EXP 有Poc或Exp ，NETWORK_EXP: 远程Exp</li><li>Level - String - 是否必填：否 - 风险等级  1,2,3,4</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateSystemVulExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSystemVulExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessEventsRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	// <li>AppId - String - 是否必填：否 - 租户id筛选</li>
	// <li>FilterTimeRange - String[] - 是否必填：否 - 筛选时间，第一个值为开始时间，第二个值为结束时间["2020-1-1 12:00:00", "2020-1-1 12:00:00"]</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAbnormalProcessEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 风险总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 严重的风险数量

		SeriousRiskCount *uint64 `json:"SeriousRiskCount,omitempty" name:"SeriousRiskCount"`
		// 高危的风险数量

		HighRiskCount *uint64 `json:"HighRiskCount,omitempty" name:"HighRiskCount"`
		// 中危的风险数量

		MiddleRiskCount *uint64 `json:"MiddleRiskCount,omitempty" name:"MiddleRiskCount"`
		// 提示的风险数量

		HintRiskCount *uint64 `json:"HintRiskCount,omitempty" name:"HintRiskCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterRiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAffectedClusterCountRequest struct {
	*tchttp.BaseRequest

	// 支持按appid过滤

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAffectedClusterCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAffectedClusterCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventEscapeImageInfo struct {
	// 镜像id，用于跳转

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 唯一值

	UniqueKey *string `json:"UniqueKey,omitempty" name:"UniqueKey"`
	// 事件类型
	//    ESCAPE_CGROUPS：利用cgroup机制逃逸
	//    ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸
	//    ESCAPE_DOCKER_API：访问Docker API接口逃逸
	//    ESCAPE_VUL_OCCURRED：逃逸漏洞利用
	//    MOUNT_SENSITIVE_PTAH：敏感路径挂载
	//    PRIVILEGE_CONTAINER_START：特权容器
	//    PRIVILEGE：程序提权逃逸

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 原始事件类型

	OriginEventType *string `json:"OriginEventType,omitempty" name:"OriginEventType"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 容器数量

	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
	// 生成时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 状态，EVENT_UNDEAL:未处理，EVENT_DEALED:已处理，EVENT_INGNORE:忽略

	Status *string `json:"Status,omitempty" name:"Status"`
	// 风险描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// AppId值

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type ReverseShellWhiteListBaseInfo struct {
	// 白名单id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 镜像数量

	ImageCount *int64 `json:"ImageCount,omitempty" name:"ImageCount"`
	// 连接进程名字

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 目标地址ip

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 目标端口

	DstPort *string `json:"DstPort,omitempty" name:"DstPort"`
	// 是否是全局白名单，true全局

	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 镜像id数组，为空代表全部

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
	// AppId值

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type DescribeEscapeWhiteListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>EventType- String - 是否必填：否 - 加白事件类型，ESCAPE_CGROUPS：利用cgroup机制逃逸，ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸， ESCAPE_DOCKER_API：访问Docker API接口逃逸，ESCAPE_VUL_OCCURRED：逃逸漏洞利用，MOUNT_SENSITIVE_PTAH：敏感路径挂载，PRIVILEGE_CONTAINER_START：特权容器， PRIVILEGE：程序提权逃逸</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：主机数量：HostCount，容器数量：ContainerCount，更新时间：UpdateTime

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeEscapeWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlRulesExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	// <li>AppId - String - 是否必填：否 - 租户id筛选</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeAccessControlRulesExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlRulesExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusScanTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeVirusScanTaskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusScanTaskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceAssetPolicyItemListRequest struct {
	*tchttp.BaseRequest

	// 客户资产的ID。

	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`
	// 起始偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 要获取的数据量，默认为10，最大为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤器列表。Name字段支持
	// RiskLevel

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComplianceAssetPolicyItemListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceAssetPolicyItemListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 木马列表

		List []*VirusInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEscapeWhiteListExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEscapeWhiteListExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEscapeWhiteListExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoStartSecurityServiceStateRequest struct {
	*tchttp.BaseRequest

	// 用户appID数组

	AppIDs []*uint64 `json:"AppIDs,omitempty" name:"AppIDs"`
	// 新增机器是否自动开通容器安全服务true:是，false：否

	AutoStartSecurityService *bool `json:"AutoStartSecurityService,omitempty" name:"AutoStartSecurityService"`
}

func (r *ModifyAutoStartSecurityServiceStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoStartSecurityServiceStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像ID

		ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
		// 镜像名称

		ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
		// 创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 镜像大小

		Size *uint64 `json:"Size,omitempty" name:"Size"`
		// 关联主机个数

		HostCnt *uint64 `json:"HostCnt,omitempty" name:"HostCnt"`
		// 关联容器个数

		ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
		// 最近扫描时间

		ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
		// 漏洞个数

		VulCnt *uint64 `json:"VulCnt,omitempty" name:"VulCnt"`
		// 风险行为数

		RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`
		// 敏感信息数

		SensitiveInfoCnt *uint64 `json:"SensitiveInfoCnt,omitempty" name:"SensitiveInfoCnt"`
		// 是否信任镜像

		IsTrustImage *bool `json:"IsTrustImage,omitempty" name:"IsTrustImage"`
		// 镜像系统

		OsName *string `json:"OsName,omitempty" name:"OsName"`
		// agent镜像扫描错误

		AgentError *string `json:"AgentError,omitempty" name:"AgentError"`
		// 后端镜像扫描错误

		ScanError *string `json:"ScanError,omitempty" name:"ScanError"`
		// 系统架构

		Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
		// 作者

		Author *string `json:"Author,omitempty" name:"Author"`
		// 构建历史

		BuildHistory *string `json:"BuildHistory,omitempty" name:"BuildHistory"`
		// 木马扫描进度

		ScanVirusProgress *uint64 `json:"ScanVirusProgress,omitempty" name:"ScanVirusProgress"`
		// 漏洞扫进度

		ScanVulProgress *uint64 `json:"ScanVulProgress,omitempty" name:"ScanVulProgress"`
		// 敏感信息扫描进度

		ScanRiskProgress *uint64 `json:"ScanRiskProgress,omitempty" name:"ScanRiskProgress"`
		// 木马扫描错误

		ScanVirusError *string `json:"ScanVirusError,omitempty" name:"ScanVirusError"`
		// 漏洞扫描错误

		ScanVulError *string `json:"ScanVulError,omitempty" name:"ScanVulError"`
		// 敏感信息错误

		ScanRiskError *string `json:"ScanRiskError,omitempty" name:"ScanRiskError"`
		// 镜像扫描状态

		ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`
		// 木马病毒数

		VirusCnt *uint64 `json:"VirusCnt,omitempty" name:"VirusCnt"`
		// 镜像扫描状态

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// 剩余扫描时间

		RemainScanTime *uint64 `json:"RemainScanTime,omitempty" name:"RemainScanTime"`
		// 授权为：1，未授权为：0

		IsAuthorized *int64 `json:"IsAuthorized,omitempty" name:"IsAuthorized"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAffectedClusterCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 严重风险的集群数量

		SeriousRiskClusterCount *uint64 `json:"SeriousRiskClusterCount,omitempty" name:"SeriousRiskClusterCount"`
		// 高危风险的集群数量

		HighRiskClusterCount *uint64 `json:"HighRiskClusterCount,omitempty" name:"HighRiskClusterCount"`
		// 中危风险的集群数量

		MiddleRiskClusterCount *uint64 `json:"MiddleRiskClusterCount,omitempty" name:"MiddleRiskClusterCount"`
		// 低危风险的集群数量

		HintRiskClusterCount *uint64 `json:"HintRiskClusterCount,omitempty" name:"HintRiskClusterCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAffectedClusterCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAffectedClusterCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type K8sApiAbnormalEventListItem struct {
	// 事件ID

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 命中规则类型

	MatchRuleType *string `json:"MatchRuleType,omitempty" name:"MatchRuleType"`
	// 威胁等级

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群运行状态

	ClusterRunningStatus *string `json:"ClusterRunningStatus,omitempty" name:"ClusterRunningStatus"`
	// 初次生成时间

	FirstCreateTime *string `json:"FirstCreateTime,omitempty" name:"FirstCreateTime"`
	// 最近一次生成时间

	LastCreateTime *string `json:"LastCreateTime,omitempty" name:"LastCreateTime"`
	// 告警数量

	AlarmCount *uint64 `json:"AlarmCount,omitempty" name:"AlarmCount"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 规则类型

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 描述信息

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 解决方案

	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 命中规则

	MatchRule *K8sApiAbnormalRuleScopeInfo `json:"MatchRule,omitempty" name:"MatchRule"`
	// 当前appID

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type DescribeTaskResultSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTaskResultSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskResultSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>AppId - String - 是否必填：否 - 租户id筛选</li>
	// <li>ImageName- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>ScanStatus - String - 是否必填：否 - 镜像扫描状态notScan，scanning，scanned，scanErr</li>
	// <li>ImageID- String - 是否必填：否 - 镜像ID筛选，</li>
	// <li>SecurityRisk- String - 是否必填：否 - 安全风险，VulCnt 、VirusCnt、RiskCnt、IsTrustImage</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetImageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulIgnoreLocalImageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 镜像列表

		List []*VulIgnoreLocalImage `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulIgnoreLocalImageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulIgnoreLocalImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCompliancePolicyItemAffectedSummaryRequest struct {
	*tchttp.BaseRequest

	// DescribeComplianceTaskPolicyItemSummaryList返回的CustomerPolicyItemId，表示检测项的ID。

	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`
	// 运营后台过滤Appid

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCompliancePolicyItemAffectedSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCompliancePolicyItemAffectedSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateK8sApiAbnormalEventExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>TimeRange - string -是否必填: 否 - 时间范围筛选 ["2022-03-31 16:55:00", "2022-03-31 17:00:00"]</li>
	// <li>MatchRules - string  - 是否必填: 否 -命中规则筛选</li>
	// <li>RiskLevel - string  - 是否必填: 否 -威胁等级筛选</li>
	// <li>Status - string  - 是否必填: 否 -事件状态筛选</li>
	// <li>MatchRuleType - string  - 是否必填: 否 -命中规则类型筛选</li>
	// <li>ClusterRunningStatus - string  - 是否必填: 否 -集群运行状态</li>
	// <li>ClusterName - string  - 是否必填: 否 -集群名称</li>
	// <li>ClusterID - string  - 是否必填: 否 -集群ID</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateK8sApiAbnormalEventExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateK8sApiAbnormalEventExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostSecurityServiceStateRequest struct {
	*tchttp.BaseRequest

	// 主机id(已废弃，使用Hosts代替)

	HostIDs []*string `json:"HostIDs,omitempty" name:"HostIDs"`
	// true：开启，false：关闭

	StartTheService *bool `json:"StartTheService,omitempty" name:"StartTheService"`
	// 主机列表

	Hosts []*ModifyHostSecurityServiceStateRequestItems `json:"Hosts,omitempty" name:"Hosts"`
}

func (r *ModifyHostSecurityServiceStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostSecurityServiceStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetProcessListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>AppId - String - 是否必填：否 - 租户id筛选</li>
	// <li>RunAs - String - 是否必填：否 - 运行用户筛选，</li>
	// <li>ContainerID - String - 是否必填：否 - 容器id</li>
	// <li>HostID- String - 是否必填：是 - 主机id</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>ProcessName- string - 是否必填：否 - 进程名搜索</li>
	// <li>Pid- string - 是否必填：否 - 进程id搜索(关联进程)</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetProcessListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetProcessListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRiskDnsEventExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>EventStatus- String - 是否必填：否 - 事件状态，待处理：EVENT_UNDEAL，EVENT_DEALED：已处理，已忽略：EVENT_IGNORE， EVENT_ADD_WHITE：已加白</li>
	// <li>ContainerStatus- String - 是否必填：否 - 容器运行状态筛选，已创建：CREATED,正常运行：RUNNING, 暂定运行：PAUSED, 停止运行：	STOPPED，重启中：RESTARTING, 迁移中：REMOVING, 销毁：DESTROYED </li>
	// <li>ContainerNetStatus- String -是否必填: 否 -  容器网络状态筛选 未隔离：NORMAL，已隔离：ISOLATED，隔离失败：ISOLATE_FAILED，解除隔离失败：RESTORE_FAILED，解除隔离中：RESTORING，隔离中：ISOLATING</li>
	// <li>EventType - String -是否必填: 否 -  事件类型，恶意域名请求：DOMAIN，恶意IP请求：IP</li>
	// <li>TimeRange- String -是否必填: 否 -  时间范围，第一个值表示开始时间，第二个值表示结束时间 </li>
	// <li>RiskDns- string - 是否必填：否 - 恶意域名。</li>
	// <li>RiskIP- string - 是否必填：否 - 恶意IP。</li>
	// <li>ContainerName- string - 是否必填：否 - 容器名称。</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID。</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	// <li>HostName- string - 是否必填：否 - 主机名称。</li>
	// <li>HostIP- string - 是否必填：否 - 内网IP。</li>
	// <li>PublicIP- string - 是否必填：否 - 外网IP。</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，最大值为100000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：事件数量：EventCount

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateRiskDnsEventExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRiskDnsEventExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RiskSyscallWhiteListInfo struct {
	// 系统调用名称，通过DescribeRiskSyscallNames接口获取枚举列表

	SyscallNames []*string `json:"SyscallNames,omitempty" name:"SyscallNames"`
	// 目标进程(已废弃)

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 镜像id数组，为空代表全部

	ImageId []*string `json:"ImageId,omitempty" name:"ImageId"`
	// 白名单id，如果新建则id为空

	Id *string `json:"Id,omitempty" name:"Id"`
	// 目标进程

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 镜像id数组，为空代表全部

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
}

type DescribeReverseShellWhiteListDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件基本信息

		WhiteListDetailInfo *ReverseShellWhiteListInfo `json:"WhiteListDetailInfo,omitempty" name:"WhiteListDetailInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellWhiteListDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellWhiteListDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecEventsTendencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运行时安全事件趋势信息列表

		EventTendencySet []*SecTendencyEventInfo `json:"EventTendencySet,omitempty" name:"EventTendencySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecEventsTendencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecEventsTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterIngressInfo struct {
	// Ingress名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 预留字段，当前不生效.

	Type *string `json:"Type,omitempty" name:"Type"`
	// VIP地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// Ingress对应的后端服务，多个以,隔开

	BackendInfo *string `json:"BackendInfo,omitempty" name:"BackendInfo"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 后端服务数量

	BackendCount *uint64 `json:"BackendCount,omitempty" name:"BackendCount"`
	// 集群状态

	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
	// 集群类型

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 标签

	Labels *string `json:"Labels,omitempty" name:"Labels"`
}

type ImageVirus struct {
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 风险等级

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 分类

	Category *string `json:"Category,omitempty" name:"Category"`
	// 病毒名称

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
	// 标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 文件类型

	FileType *string `json:"FileType,omitempty" name:"FileType"`
	// 文件路径

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件md5

	FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`
	// 大小

	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`
	// 首次发现时间

	FirstScanTime *string `json:"FirstScanTime,omitempty" name:"FirstScanTime"`
	// 最近扫描时间

	LatestScanTime *string `json:"LatestScanTime,omitempty" name:"LatestScanTime"`
}

type DescribeScanIgnoreVulListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>VulName- string - 是否必填：否 - 漏洞名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式:DESC,ACS

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段 UpdateTime

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeScanIgnoreVulListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanIgnoreVulListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventEscapeImageListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数:
	// EventType: 事件类型(MOUNT_SENSITIVE_PTAH:敏感路径挂载 PRIVILEGE_CONTAINER_START:特权容器)
	// Status: 事件状态(EVENT_UNDEAL:未处理，EVENT_DEALED:已处理，EVENT_INGNORE:忽略)
	// ImageID: 镜像id
	// ImageName:镜像名称

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeEventEscapeImageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventEscapeImageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSearchTemplatesRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeSearchTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSearchTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecTendencyEventInfo struct {
	// 趋势列表

	EventSet []*RunTimeTendencyInfo `json:"EventSet,omitempty" name:"EventSet"`
	// 事件类型：
	// ET_ESCAPE : 容器逃逸
	// ET_REVERSE_SHELL: 反弹shell
	// ET_RISK_SYSCALL:高危系统调用
	// ET_ABNORMAL_PROCESS: 异常进程
	// ET_ACCESS_CONTROL 文件篡改

	EventType *string `json:"EventType,omitempty" name:"EventType"`
}

type DescribeAssetImageRegistryRiskListExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像信息

	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`
	// 排序字段（Level）

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 + -

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 镜像标识Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryRiskListExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryRiskListExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AbnormalProcessEventDescription struct {
	// 事件规则

	Description *string `json:"Description,omitempty" name:"Description"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 命中规则详细信息

	MatchRule *AbnormalProcessChildRuleInfo `json:"MatchRule,omitempty" name:"MatchRule"`
	// 命中规则名字

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 命中规则的id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 事件最后一次处理的时间

	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
	// 命中策略名称：SYSTEM_DEFINED_RULE （系统策略）或  用户自定义的策略名字

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type DescribeImageDenyRuleSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImageDenyRuleSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyRuleSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceAssetPolicyItemListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回检测项的总数。如果用户未启用基线检查，此处返回0。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回某个资产下的检测项的列表。

		AssetPolicyItemList []*ComplianceAssetPolicyItem `json:"AssetPolicyItemList,omitempty" name:"AssetPolicyItemList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComplianceAssetPolicyItemListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceAssetPolicyItemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellDetailRequest struct {
	*tchttp.BaseRequest

	// 事件唯一id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 租户id

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeReverseShellDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetAppServiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// db服务列表

		List []*ServiceInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetAppServiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetAppServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageVulListExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Name- String - 是否必填：否 - 漏洞名称名称筛选，</li>
	// <li>Level - String - 是否必填：否 - 风险等级  1,2,3,4</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
}

func (r *DescribeAssetImageVulListExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageVulListExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceImageDetailInfo struct {
	// 镜像在主机上的ID。

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像的名称。

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像的Tag。

	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
	// 镜像所在远程仓库的路径。

	Repository *string `json:"Repository,omitempty" name:"Repository"`
}

type DescribeComplianceAssetDetailInfoRequest struct {
	*tchttp.BaseRequest

	// 客户资产ID。

	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`
	// 用户Appid过滤

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComplianceAssetDetailInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceAssetDetailInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageComponentListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 镜像组件列表

		List []*ImageComponent `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageComponentListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageComponentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusAutoIsolateSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自动隔离开关(true:开 false:关)

		AutoIsolateSwitch *bool `json:"AutoIsolateSwitch,omitempty" name:"AutoIsolateSwitch"`
		// 是否中断隔离文件关联的进程(true:是 false:否)

		IsKillProgress *bool `json:"IsKillProgress,omitempty" name:"IsKillProgress"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusAutoIsolateSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImagesVul struct {
	// 漏洞id

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// 漏洞名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 分类

	Category *string `json:"Category,omitempty" name:"Category"`
	// 分类2

	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`
	// 风险等级

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 描述

	Des *string `json:"Des,omitempty" name:"Des"`
	// 解决方案

	OfficialSolution *string `json:"OfficialSolution,omitempty" name:"OfficialSolution"`
	// 引用

	Reference *string `json:"Reference,omitempty" name:"Reference"`
	// 防御方案

	DefenseSolution *string `json:"DefenseSolution,omitempty" name:"DefenseSolution"`
	// 提交时间

	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`
	// CVSS V3分数

	CVSSV3Score *float64 `json:"CVSSV3Score,omitempty" name:"CVSSV3Score"`
	// CVSS V3描述

	CVSSV3Desc *string `json:"CVSSV3Desc,omitempty" name:"CVSSV3Desc"`
	// 是否是重点关注：true：是，false：不是

	IsSuggest *bool `json:"IsSuggest,omitempty" name:"IsSuggest"`
	// 修复版本号

	FixedVersions *string `json:"FixedVersions,omitempty" name:"FixedVersions"`
	// 漏洞标签:"CanBeFixed","DynamicLevelPoc","DynamicLevelExp"

	Tag []*string `json:"Tag,omitempty" name:"Tag"`
}

type UserPurchaseInfo struct {
	// 租户id

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 客户名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 开始使用时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束使用时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 专业版购买状态

	State *int64 `json:"State,omitempty" name:"State"`
	// 购买的核心数

	PurchaseCores *uint64 `json:"PurchaseCores,omitempty" name:"PurchaseCores"`
	// 购买的镜像授权数

	PurchasedAuthorizedCnt *uint64 `json:"PurchasedAuthorizedCnt,omitempty" name:"PurchasedAuthorizedCnt"`
	// 赠送的镜像授权数

	GivenAuthorizedCnt *uint64 `json:"GivenAuthorizedCnt,omitempty" name:"GivenAuthorizedCnt"`
	// 新增主机是否开启容器安全服务

	AutoStartSecurityService *bool `json:"AutoStartSecurityService,omitempty" name:"AutoStartSecurityService"`
	// 弹性计费开关

	FlexBillingSwitch *bool `json:"FlexBillingSwitch,omitempty" name:"FlexBillingSwitch"`
	// 用户总核数

	CoresTotalCnt *int64 `json:"CoresTotalCnt,omitempty" name:"CoresTotalCnt"`
}

type DescribeReverseShellWhiteListDetailRequest struct {
	*tchttp.BaseRequest

	// 白名单id

	WhiteListId *string `json:"WhiteListId,omitempty" name:"WhiteListId"`
	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeReverseShellWhiteListDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellWhiteListDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellEventsExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeReverseShellEventsExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellEventsExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEmergencyVulListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 漏洞列表

		List []*EmergencyVulInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEmergencyVulListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEmergencyVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunTimeEventBaseInfo struct {
	// 事件唯一ID

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 事件发现时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// Pod名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 状态， “EVENT_UNDEAL”:事件未处理
	//     "EVENT_DEALED":事件已经处理
	//     "EVENT_INGNORE"：事件已经忽略

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件名称：
	// 宿主机文件访问逃逸、
	// Syscall逃逸、
	// MountNamespace逃逸、
	// 程序提权逃逸、
	// 特权容器启动逃逸、
	// 敏感路径挂载
	// 恶意进程启动
	// 文件篡改

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 事件类型
	//    ESCAPE_HOST_ACESS_FILE:宿主机文件访问逃逸
	//    ESCAPE_MOUNT_NAMESPACE:MountNamespace逃逸
	//    ESCAPE_PRIVILEDGE:程序提权逃逸
	//    ESCAPE_PRIVILEDGE_CONTAINER_START:特权容器启动逃逸
	//    ESCAPE_MOUNT_SENSITIVE_PTAH:敏感路径挂载
	//    ESCAPE_SYSCALL:Syscall逃逸

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 事件数量

	EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 容器隔离状态

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 容器隔离子状态

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// 内网ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 外网ip

	ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`
	// 节点ID

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// 节点类型:NORMAL:普通节点;SUPER:超级节点

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 	节点子网ID

	NodeSubNetID *string `json:"NodeSubNetID,omitempty" name:"NodeSubNetID"`
	// 节点子网名称

	NodeSubNetName *string `json:"NodeSubNetName,omitempty" name:"NodeSubNetName"`
	// 节点子网网段

	NodeSubNetCIDR *string `json:"NodeSubNetCIDR,omitempty" name:"NodeSubNetCIDR"`
	// podIP

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// pod状态

	PodStatus *string `json:"PodStatus,omitempty" name:"PodStatus"`
	// 集群id

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// uuid

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// Namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// WorkloadType

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
}

type RunTimeRiskInfo struct {
	// 数量

	Cnt *uint64 `json:"Cnt,omitempty" name:"Cnt"`
	// 风险等级：
	// CRITICAL: 严重
	// HIGH: 高
	// MEDIUM：中
	// LOW: 低

	Level *string `json:"Level,omitempty" name:"Level"`
}

type DescribeMaliciousConnectionBlackListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 恶意请求白名单总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 恶意请求白名单列表

		List []*MaliciousConnectionRuleInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMaliciousConnectionBlackListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaliciousConnectionBlackListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndexListRequest struct {
	*tchttp.BaseRequest

	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeIndexListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIndexListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRefreshTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 新任务ID

	NewTaskID *string `json:"NewTaskID,omitempty" name:"NewTaskID"`
}

func (r *DescribeRefreshTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRefreshTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagInfo struct {
	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type DescribeVulContainerListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 容器列表

		List []*VulAffectedContainerInfo `json:"List,omitempty" name:"List"`
		// 容器总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulContainerListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulContainerListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusAutoIsolateSettingListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 列表

		List []*VirusAutoIsolateSettingInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusAutoIsolateSettingListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSettingListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlRuleDetailRequest struct {
	*tchttp.BaseRequest

	// 策略唯一id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 镜像id, 仅仅在事件加白的时候使用

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 租户id

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeAccessControlRuleDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlRuleDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageRepoInfo struct {
	// 镜像Digest

	ImageDigest *string `json:"ImageDigest,omitempty" name:"ImageDigest"`
	// 镜像仓库地址

	ImageRepoAddress *string `json:"ImageRepoAddress,omitempty" name:"ImageRepoAddress"`
	// 仓库类型

	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像版本

	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
	// 镜像大小

	ImageSize *uint64 `json:"ImageSize,omitempty" name:"ImageSize"`
	// 最近扫描时间

	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
	// 扫描状态

	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`
	// 安全漏洞数

	VulCnt *uint64 `json:"VulCnt,omitempty" name:"VulCnt"`
	// 木马病毒数

	VirusCnt *uint64 `json:"VirusCnt,omitempty" name:"VirusCnt"`
	// 风险行为数

	RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`
	// 敏感信息数

	SentiveInfoCnt *uint64 `json:"SentiveInfoCnt,omitempty" name:"SentiveInfoCnt"`
	// 是否可信镜像

	IsTrustImage *bool `json:"IsTrustImage,omitempty" name:"IsTrustImage"`
	// 镜像系统

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 木马扫描错误

	ScanVirusError *string `json:"ScanVirusError,omitempty" name:"ScanVirusError"`
	// 漏洞扫描错误

	ScanVulError *string `json:"ScanVulError,omitempty" name:"ScanVulError"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 高危扫描错误

	ScanRiskError *string `json:"ScanRiskError,omitempty" name:"ScanRiskError"`
	// 木马扫描进度

	ScanVirusProgress *uint64 `json:"ScanVirusProgress,omitempty" name:"ScanVirusProgress"`
	// 漏洞扫描进度

	ScanVulProgress *uint64 `json:"ScanVulProgress,omitempty" name:"ScanVulProgress"`
	// 高危扫描进度

	ScanRiskProgress *uint64 `json:"ScanRiskProgress,omitempty" name:"ScanRiskProgress"`
	// 剩余扫描时间秒

	ScanRemainTime *uint64 `json:"ScanRemainTime,omitempty" name:"ScanRemainTime"`
	// cve扫描状态

	CveStatus *string `json:"CveStatus,omitempty" name:"CveStatus"`
	// 高危扫描状态

	RiskStatus *string `json:"RiskStatus,omitempty" name:"RiskStatus"`
	// 木马扫描状态

	VirusStatus *string `json:"VirusStatus,omitempty" name:"VirusStatus"`
	// 总进度

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// 授权状态

	IsAuthorized *uint64 `json:"IsAuthorized,omitempty" name:"IsAuthorized"`
	// 仓库区域

	RegistryRegion *string `json:"RegistryRegion,omitempty" name:"RegistryRegion"`
	// 列表id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 镜像Id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像所属用户的AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 敏感信息个数

	SensitiveInfoCnt *uint64 `json:"SensitiveInfoCnt,omitempty" name:"SensitiveInfoCnt"`
}

type VirusTaskInfo struct {
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像Id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 扫描状态：
	// WAIT: 等待扫描
	// FAILED: 失败
	// SCANNING: 扫描中
	// FINISHED: 结束
	// CANCELING: 取消中
	// CANCELED: 已取消
	// CANCEL_FAILED： 取消失败

	Status *string `json:"Status,omitempty" name:"Status"`
	// 检测开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 检测结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 风险个数

	RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`
	// 事件id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 错误原因:
	// SEND_SUCCESSED: 下发成功
	// SCAN_WAIT: agent排队扫描等待中
	// OFFLINE: 离线
	// SEND_FAILED:下发失败
	// TIMEOUT: 超时
	// LOW_AGENT_VERSION: 客户端版本过低
	// AGENT_NOT_FOUND: 镜像所属客户端版不存在
	// TOO_MANY: 任务过多
	// VALIDATION: 参数非法
	// INTERNAL: 服务内部错误
	// MISC: 其他错误
	// UNAUTH: 所在镜像未授权
	// SEND_CANCEL_SUCCESSED:下发成功

	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
	// 当前appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type DescribeImageDenyEventDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件ID

		EventID *int64 `json:"EventID,omitempty" name:"EventID"`
		// 事件类型 EVENT_RISK:风险事件类型，EVENT_PRIVILEGE:特权

		EventType *string `json:"EventType,omitempty" name:"EventType"`
		// 规则名称

		RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
		// 规则RuleID

		RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
		// 规则类型

		RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
		// 规则启用状态 0:开启，1:关闭

		RuleStatus *int64 `json:"RuleStatus,omitempty" name:"RuleStatus"`
		// 规则策略状态 IN_THE_TEST ：观察中，IN_EFFECT：生效中

		RuleEffectStatus *string `json:"RuleEffectStatus,omitempty" name:"RuleEffectStatus"`
		// 规则内容

		RuleInfo []*string `json:"RuleInfo,omitempty" name:"RuleInfo"`
		// 规则描述

		RuleDescription *string `json:"RuleDescription,omitempty" name:"RuleDescription"`
		// 镜像ID

		ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
		// 镜像名称

		ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
		// 节点名称

		NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
		// 内网IP

		NodeIP *string `json:"NodeIP,omitempty" name:"NodeIP"`
		// 外网IP

		PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
		// 主机Quuid

		QUUID *string `json:"QUUID,omitempty" name:"QUUID"`
		// 首次生成时间

		FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
		// 最近生成时间

		LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
		// 事件数量

		EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
		// 执行动作:
		// BEHAVIOR_ALERT:告警，
		// BEHAVIOR_HOLDUP_SUCCESSED:拦截

		DealBehavior *string `json:"DealBehavior,omitempty" name:"DealBehavior"`
		// Pod名称

		PodName *string `json:"PodName,omitempty" name:"PodName"`
		// 规则开始拦截时间

		RuleEffectTime *string `json:"RuleEffectTime,omitempty" name:"RuleEffectTime"`
		// 事件描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// 镜像启动参数

		StartParam *string `json:"StartParam,omitempty" name:"StartParam"`
		// 解决方案

		Solution *string `json:"Solution,omitempty" name:"Solution"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageDenyEventDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusAutoIsolateSampleDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件Md5值

		MD5 *string `json:"MD5,omitempty" name:"MD5"`
		// 文件大小(B)

		Size *uint64 `json:"Size,omitempty" name:"Size"`
		// 病毒名

		VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
		// 风险等级 RISK_CRITICAL, RISK_HIGH, RISK_MEDIUM, RISK_LOW, RISK_NOTICE。

		RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
		// 查杀引擎

		KillEngine []*string `json:"KillEngine,omitempty" name:"KillEngine"`
		// 标签

		Tags []*string `json:"Tags,omitempty" name:"Tags"`
		// 事件描述

		HarmDescribe *string `json:"HarmDescribe,omitempty" name:"HarmDescribe"`
		// 建议方案

		SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`
		// 参考链接

		ReferenceLink *string `json:"ReferenceLink,omitempty" name:"ReferenceLink"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusAutoIsolateSampleDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSampleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterWorkloadsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群workload信息

		ClusterWorkloadList []*ClusterWorkloadInfo `json:"ClusterWorkloadList,omitempty" name:"ClusterWorkloadList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterWorkloadsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterWorkloadsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageVulListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像漏洞列表

		List []*ImagesVul `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageVulListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExportComplianceStatusListJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的导出任务的ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateExportComplianceStatusListJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExportComplianceStatusListJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RiskSyscallWhiteListBaseInfo struct {
	// 白名单Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 镜像数量

	ImageCount *uint64 `json:"ImageCount,omitempty" name:"ImageCount"`
	// 连接进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 系统调用名称列表

	SyscallNames []*string `json:"SyscallNames,omitempty" name:"SyscallNames"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 是否是全局白名单,true全局

	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 镜像id数组

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
	// AppId值

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type CreateVulImageExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVulImageExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulImageExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessDetailRequest struct {
	*tchttp.BaseRequest

	// 事件唯一id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 租户id

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeAbnormalProcessDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetHostListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// 	<li>AppId - String - 是否必填：否 - 租户id筛选</li>
	// 	<li>Status - String - 是否必填：否 - 主机运行状态筛选，0："offline",1："online", 2："paused"</li>
	// 	<li>HostName - String - 是否必填：否 - 主机名筛选</li>
	// 	<li>Group- String - 是否必填：否 - 主机群组搜索</li>
	// 	<li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// 	<li>HostID- string - 是否必填：否 - 主机id搜索</li>
	// 	<li>DockerVersion- string - 是否必填：否 - docker版本搜索</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetHostListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetHostListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAbnormalProcessRulesExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAbnormalProcessRulesExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAbnormalProcessRulesExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageInfo struct {
	// 镜像id

	ImageDigest *string `json:"ImageDigest,omitempty" name:"ImageDigest"`
	// 仓库类型

	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
	// 镜像仓库地址

	ImageRepoAddress *string `json:"ImageRepoAddress,omitempty" name:"ImageRepoAddress"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像tag

	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
	// 强制扫描

	Force *string `json:"Force,omitempty" name:"Force"`
}

type CreateVulExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVulExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageDetailRequest struct {
	*tchttp.BaseRequest

	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 租户id, 运营端

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeAssetImageDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportJobResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出的状态。取值为, SUCCESS:成功、FAILURE:失败，RUNNING: 进行中。

		ExportStatus *string `json:"ExportStatus,omitempty" name:"ExportStatus"`
		// 返回下载URL

		DownloadURL *string `json:"DownloadURL,omitempty" name:"DownloadURL"`
		// 当ExportStatus为RUNNING时，返回导出进度。0~100范围的浮点数。

		ExportProgress *float64 `json:"ExportProgress,omitempty" name:"ExportProgress"`
		// 失败原因

		FailureMsg *string `json:"FailureMsg,omitempty" name:"FailureMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportJobResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportJobResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterSummaryRequest struct {
	*tchttp.BaseRequest

	// 支持appid过滤

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeClusterSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRiskListExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一的任务ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRiskListExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRiskListExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeEventDetailRequest struct {
	*tchttp.BaseRequest

	// 事件唯一id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 租户id

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
	// 事件类型

	EventType *string `json:"EventType,omitempty" name:"EventType"`
}

func (r *DescribeEscapeEventDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEventEscapeImageStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEventEscapeImageStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEventEscapeImageStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulImageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 受影响的镜像列表

		List []*VulAffectedImageInfo `json:"List,omitempty" name:"List"`
		// 镜像总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulImageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlEventsRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	// <li>AppId - String - 是否必填：否 - 租户id筛选</li>
	// <li>FilterTimeRange - String[] - 是否必填：否 - 筛选时间，第一个值为开始时间，第二个值为结束时间["2020-1-1 12:00:00", "2020-1-1 12:00:00"]</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAccessControlEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndexListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ES 索引信息

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIndexListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIndexListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterDetailExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要导出的集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 支持appid过滤

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *CreateClusterDetailExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterDetailExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyRuleListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>RuleType- String - 是否必填：否 -规则类型 RULE_RISK：风险， RULE_PRIVILEGE：特权</li>
	// <li>EffectStatus- String - 是否必填：否 - 生效状态 IN_THE_TEST ：观察中，IN_EFFECT：生效中。</li>
	// <li>RuleName- string - 是否必填：否 - 规则名称。</li>
	// <li>Status- string - 是否必填：否 - 开启状态 0：开启，1：关闭。</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：生效时间：EffectTime，更新时间：UpdateTime

	By *string `json:"By,omitempty" name:"By"`
	// 置顶已开启规则 true：是 ，否：false

	TopTurnOn *bool `json:"TopTurnOn,omitempty" name:"TopTurnOn"`
}

func (r *DescribeImageDenyRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetDBServiceListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>AppId - String - 是否必填：否 - 租户id筛选</li>
	// <li>Keywords- String - 是否必填：否 - 模糊查询可选字段</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetDBServiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDBServiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenTcssTrialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenTcssTrialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenTcssTrialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCompliancePolicyItemAffectedSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回各检测项影响的资产的汇总信息。

		PolicyItemSummary *CompliancePolicyItemSummary `json:"PolicyItemSummary,omitempty" name:"PolicyItemSummary"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCompliancePolicyItemAffectedSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCompliancePolicyItemAffectedSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageDenyEventTendency struct {
	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
}

type DescribeAssetImageVirusListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Name- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>RiskLevel - String - 是否必填：否 - 风险等级  1,2,3,4</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 排序 asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAssetImageVirusListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageVirusListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageBindRuleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 镜像绑定规则列表

		ImageBindRuleSet []*ImagesBindRuleInfo `json:"ImageBindRuleSet,omitempty" name:"ImageBindRuleSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageBindRuleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageBindRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceK8SDetailInfo struct {
	// K8S集群的名称。

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
}

type DescribeRiskListRequest struct {
	*tchttp.BaseRequest

	// 要查询的集群ID，如果不指定，则查询用户所有的风险项

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：RiskLevel风险等级, RiskTarget检查对象，风险对象,RiskType风险类别,RiskAttribute检测项所属的风险类型,Name

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeRiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlRulesRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	// <li>AppId - String - 是否必填：否 - 租户id筛选</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAccessControlRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSystemVulExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSystemVulExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSystemVulExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusScanSettingRequest struct {
	*tchttp.BaseRequest

	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeVirusScanSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusScanSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceScanFailedAssetListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回检测失败的资产的总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回各类检测失败的资产的汇总信息的列表。

		ScanFailedAssetList []*ComplianceScanFailedAsset `json:"ScanFailedAssetList,omitempty" name:"ScanFailedAssetList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComplianceScanFailedAssetListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceScanFailedAssetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImagesBindRuleInfo struct {
	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 关联容器数量

	ContainerCnt *int64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
	// 绑定规则id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名字

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 镜像大小

	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`
	// 最近扫描时间

	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
}

type DescribeK8sApiAbnormalEventListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>TimeRange - string -是否必填: 否 - 时间范围筛选 ["2022-03-31 16:55:00", "2022-03-31 17:00:00"]</li>
	// <li>MatchRules - string  - 是否必填: 否 -命中规则筛选</li>
	// <li>RiskLevel - string  - 是否必填: 否 -威胁等级筛选</li>
	// <li>Status - string  - 是否必填: 否 -事件状态筛选</li>
	// <li>MatchRuleType - string  - 是否必填: 否 -命中规则类型筛选</li>
	// <li>ClusterRunningStatus - string  - 是否必填: 否 -集群运行状态</li>
	// <li>ClusterName - string  - 是否必填: 否 -集群名称</li>
	// <li>ClusterID - string  - 是否必填: 否 -集群ID</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段
	// LatestFoundTime: 最近生成时间
	// AlarmCount: 告警数量

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeK8sApiAbnormalEventListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalEventListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryVulListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像信息

	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`
	// 排序字段（Level）

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 + -

	Order *string `json:"Order,omitempty" name:"Order"`
	// 镜像标识Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryVulListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryVulListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEventEscapeImageExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数:
	// EventType: 事件类型(MOUNT_SENSITIVE_PTAH:敏感路径挂载 PRIVILEGE_CONTAINER_START:特权容器)
	// ImageID: 镜像id
	// ImageName:镜像名称

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：latest_found_time

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateEventEscapeImageExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEventEscapeImageExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageRiskTendencyInfo struct {
	// 趋势列表

	ImageRiskSet []*RunTimeTendencyInfo `json:"ImageRiskSet,omitempty" name:"ImageRiskSet"`
	// 风险类型：
	// IRT_VULNERABILITY : 安全漏洞
	// IRT_MALWARE_VIRUS: 木马病毒
	// IRT_RISK:敏感信息

	ImageRiskType *string `json:"ImageRiskType,omitempty" name:"ImageRiskType"`
}

type DescribeEscapeWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 逃逸白名单列表

		List []*EscapeWhiteListInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEscapeWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 有风险的集群数量

		RiskClusterCount *uint64 `json:"RiskClusterCount,omitempty" name:"RiskClusterCount"`
		// 未检查的集群数量

		UncheckClusterCount *uint64 `json:"UncheckClusterCount,omitempty" name:"UncheckClusterCount"`
		// 托管集群数量

		ManagedClusterCount *uint64 `json:"ManagedClusterCount,omitempty" name:"ManagedClusterCount"`
		// 独立集群数量

		IndependentClusterCount *uint64 `json:"IndependentClusterCount,omitempty" name:"IndependentClusterCount"`
		// 无风险的集群数量

		NoRiskClusterCount *uint64 `json:"NoRiskClusterCount,omitempty" name:"NoRiskClusterCount"`
		// 已经检查集群数

		CheckedClusterCount *uint64 `json:"CheckedClusterCount,omitempty" name:"CheckedClusterCount"`
		// 自动检查集群数

		AutoCheckClusterCount *uint64 `json:"AutoCheckClusterCount,omitempty" name:"AutoCheckClusterCount"`
		// 手动检查集群数

		ManualCheckClusterCount *uint64 `json:"ManualCheckClusterCount,omitempty" name:"ManualCheckClusterCount"`
		// 检查失败集群数

		FailedClusterCount *uint64 `json:"FailedClusterCount,omitempty" name:"FailedClusterCount"`
		// 未导入的集群数量

		NotImportedClusterCount *uint64 `json:"NotImportedClusterCount,omitempty" name:"NotImportedClusterCount"`
		// eks集群数量

		ServerlessClusterCount *uint64 `json:"ServerlessClusterCount,omitempty" name:"ServerlessClusterCount"`
		// TKE集群数量

		TkeClusterCount *uint64 `json:"TkeClusterCount,omitempty" name:"TkeClusterCount"`
		// 用户自建腾讯云集群数量

		UserCreateTencentClusterCount *uint64 `json:"UserCreateTencentClusterCount,omitempty" name:"UserCreateTencentClusterCount"`
		// 用户自建集群混合云数量

		UserCreateHybridClusterCount *uint64 `json:"UserCreateHybridClusterCount,omitempty" name:"UserCreateHybridClusterCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EscapeEventTendencyInfo struct {
	// 待处理风险容器事件总数

	RiskContainerEventCount *int64 `json:"RiskContainerEventCount,omitempty" name:"RiskContainerEventCount"`
	// 待处理程序特权事件总数

	ProcessPrivilegeEventCount *int64 `json:"ProcessPrivilegeEventCount,omitempty" name:"ProcessPrivilegeEventCount"`
	// 待处理容器逃逸事件总数

	ContainerEscapeEventCount *int64 `json:"ContainerEscapeEventCount,omitempty" name:"ContainerEscapeEventCount"`
	// 日期

	Date *string `json:"Date,omitempty" name:"Date"`
}

type DescribeEscapeEventInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 逃逸事件数组

		EventSet []*EscapeEventInfo `json:"EventSet,omitempty" name:"EventSet"`
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEscapeEventInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterIngressYamlRequest struct {
	*tchttp.BaseRequest

	// ingress名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群地域

	ClusterRegion *string `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
}

func (r *DescribeClusterIngressYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterIngressYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceTaskAssetSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回用户的状态，
		//
		// USER_UNINIT: 用户未初始化。
		// USER_INITIALIZING，表示用户正在初始化环境。
		// USER_NORMAL: 正常状态。

		Status *string `json:"Status,omitempty" name:"Status"`
		// 返回各类资产的汇总信息的列表。

		AssetSummaryList []*ComplianceAssetSummary `json:"AssetSummaryList,omitempty" name:"AssetSummaryList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComplianceTaskAssetSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceTaskAssetSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetAppServiceListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>AppId - String - 是否必填：否 - 租户id筛选</li>
	// <li>Keywords- String - 是否必填：否 - 模糊查询可选字段</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetAppServiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetAppServiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusScanTimeoutSettingRequest struct {
	*tchttp.BaseRequest

	// 设置类型0一键检测，1定时检测

	ScanType *uint64 `json:"ScanType,omitempty" name:"ScanType"`
	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeVirusScanTimeoutSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusScanTimeoutSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRiskListExportJobRequest struct {
	*tchttp.BaseRequest

	// 风险项Id列表，为空则导出所有风险项

	CheckItemIdList []*uint64 `json:"CheckItemIdList,omitempty" name:"CheckItemIdList"`
	// 根据appid过滤

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *CreateRiskListExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRiskListExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RiskSyscallEventDescription struct {
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 系统调用名称

	SyscallName *string `json:"SyscallName,omitempty" name:"SyscallName"`
	// 事件最后一次处理的时间

	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
}

type ClusterInfoItem struct {
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 集群操作系统

	ClusterOs *string `json:"ClusterOs,omitempty" name:"ClusterOs"`
	// 集群类型

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 集群节点数

	ClusterNodeNum *uint64 `json:"ClusterNodeNum,omitempty" name:"ClusterNodeNum"`
	// 集群区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 监控组件的状态，为Defender_Uninstall、Defender_Normal、Defender_Error、Defender_Installing

	DefenderStatus *string `json:"DefenderStatus,omitempty" name:"DefenderStatus"`
	// 集群状态

	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
	// 集群的检测模式，为Cluster_Normal或者Cluster_Actived.

	ClusterCheckMode *string `json:"ClusterCheckMode,omitempty" name:"ClusterCheckMode"`
	// 是否自动定期检测

	ClusterAutoCheck *bool `json:"ClusterAutoCheck,omitempty" name:"ClusterAutoCheck"`
	// 防护容器部署失败原因，为UserDaemonSetNotReady时,和UnreadyNodeNum转成"N个节点防御容器为就绪"，其他错误直接展示

	DefenderErrorReason *string `json:"DefenderErrorReason,omitempty" name:"DefenderErrorReason"`
	// 防御容器没有ready状态的节点数量

	UnreadyNodeNum *uint64 `json:"UnreadyNodeNum,omitempty" name:"UnreadyNodeNum"`
	// 严重风险检查项的数量

	SeriousRiskCount *int64 `json:"SeriousRiskCount,omitempty" name:"SeriousRiskCount"`
	// 高风险检查项的数量

	HighRiskCount *int64 `json:"HighRiskCount,omitempty" name:"HighRiskCount"`
	// 中风险检查项的数量

	MiddleRiskCount *int64 `json:"MiddleRiskCount,omitempty" name:"MiddleRiskCount"`
	// 提示风险检查项的数量

	HintRiskCount *int64 `json:"HintRiskCount,omitempty" name:"HintRiskCount"`
	// 检查失败原因

	CheckFailReason *string `json:"CheckFailReason,omitempty" name:"CheckFailReason"`
	// 检查状态,为Task_Running, NoRisk, HasRisk, Uncheck, Task_Error

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 任务创建时间,检查时间

	TaskCreateTime *string `json:"TaskCreateTime,omitempty" name:"TaskCreateTime"`
	// 用户的AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 接入状态

	AccessedStatus *string `json:"AccessedStatus,omitempty" name:"AccessedStatus"`
	// 接入失败原因

	AccessedSubStatus *string `json:"AccessedSubStatus,omitempty" name:"AccessedSubStatus"`
	// 节点总数

	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 离线节点数

	OffLineNodeCount *uint64 `json:"OffLineNodeCount,omitempty" name:"OffLineNodeCount"`
	// 未安装agent节点数

	UnInstallAgentNodeCount *uint64 `json:"UnInstallAgentNodeCount,omitempty" name:"UnInstallAgentNodeCount"`
	// 计费核数

	ChargeCoresCnt *uint64 `json:"ChargeCoresCnt,omitempty" name:"ChargeCoresCnt"`
}

type ProcessInfo struct {
	// 进程启动时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 运行用户

	RunAs *string `json:"RunAs,omitempty" name:"RunAs"`
	// 命令行参数

	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`
	// Exe路径

	Exe *string `json:"Exe,omitempty" name:"Exe"`
	// 主机PID

	PID *uint64 `json:"PID,omitempty" name:"PID"`
	// 容器内pid

	ContainerPID *uint64 `json:"ContainerPID,omitempty" name:"ContainerPID"`
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 主机id

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 主机ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 租户id,运营端使用

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 外网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 节点id

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// podip

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// pod名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 超级节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
}

type VulAffectedImageComponentInfo struct {
	// 组件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 组件修复版本

	FixedVersion *string `json:"FixedVersion,omitempty" name:"FixedVersion"`
	// 组件路径

	Path *string `json:"Path,omitempty" name:"Path"`
}

type DescribeK8sApiAbnormalRuleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则详情

		Info *K8sApiAbnormalRuleInfo `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeK8sApiAbnormalRuleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceAssetListRequest struct {
	*tchttp.BaseRequest

	// 资产类型列表。
	// ASSET_CONTAINER, 容器
	// ASSET_IMAGE, 镜像
	// ASSET_HOST, 主机
	// ASSET_K8S, K8S资产

	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`
	// 起始偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回的数据量，默认为10，最大为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 查询过滤器

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComplianceAssetListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceAssetListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessRuleDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异常进程策略详细信息

		RuleDetail *AbnormalProcessRuleInfo `json:"RuleDetail,omitempty" name:"RuleDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalProcessRuleDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterIngressYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Ingress格式字符串,base64编码

		Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterIngressYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterIngressYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterWorkloadsRequest struct {
	*tchttp.BaseRequest

	// 集群Id,不输入表示查询所有

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：WorkloadName工作负载名字, WorkloadType工作负载类型等

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeClusterWorkloadsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterWorkloadsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterDetailRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 过滤字段 支持appid

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeClusterDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像Digest

		ImageDigest *string `json:"ImageDigest,omitempty" name:"ImageDigest"`
		// 镜像地址

		ImageRepoAddress *string `json:"ImageRepoAddress,omitempty" name:"ImageRepoAddress"`
		// 镜像类型

		RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
		// 仓库名称

		ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
		// 镜像版本

		ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
		// 扫描时间

		ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
		// 扫描状态

		ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`
		// 安全漏洞数

		VulCnt *uint64 `json:"VulCnt,omitempty" name:"VulCnt"`
		// 木马病毒数

		VirusCnt *uint64 `json:"VirusCnt,omitempty" name:"VirusCnt"`
		// 风险行为数

		RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`
		// 敏感信息数

		SentiveInfoCnt *uint64 `json:"SentiveInfoCnt,omitempty" name:"SentiveInfoCnt"`
		// 是否可信镜像

		IsTrustImage *bool `json:"IsTrustImage,omitempty" name:"IsTrustImage"`
		// 镜像系统

		OsName *string `json:"OsName,omitempty" name:"OsName"`
		// 木马扫描错误

		ScanVirusError *string `json:"ScanVirusError,omitempty" name:"ScanVirusError"`
		// 漏洞扫描错误

		ScanVulError *string `json:"ScanVulError,omitempty" name:"ScanVulError"`
		// 层文件信息

		LayerInfo *string `json:"LayerInfo,omitempty" name:"LayerInfo"`
		// 实例id

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 实例名称

		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
		// 命名空间

		Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
		// 高危扫描错误

		ScanRiskError *string `json:"ScanRiskError,omitempty" name:"ScanRiskError"`
		// 木马信息扫描进度

		ScanVirusProgress *uint64 `json:"ScanVirusProgress,omitempty" name:"ScanVirusProgress"`
		// 漏洞扫描进度

		ScanVulProgress *uint64 `json:"ScanVulProgress,omitempty" name:"ScanVulProgress"`
		// 高危扫描进度

		ScanRiskProgress *uint64 `json:"ScanRiskProgress,omitempty" name:"ScanRiskProgress"`
		// 剩余扫描时间秒

		ScanRemainTime *uint64 `json:"ScanRemainTime,omitempty" name:"ScanRemainTime"`
		// cve扫描状态

		CveStatus *string `json:"CveStatus,omitempty" name:"CveStatus"`
		// 高危扫描状态

		RiskStatus *string `json:"RiskStatus,omitempty" name:"RiskStatus"`
		// 木马扫描状态

		VirusStatus *string `json:"VirusStatus,omitempty" name:"VirusStatus"`
		// 总进度

		Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
		// 授权状态

		IsAuthorized *uint64 `json:"IsAuthorized,omitempty" name:"IsAuthorized"`
		// 镜像大小

		ImageSize *uint64 `json:"ImageSize,omitempty" name:"ImageSize"`
		// 镜像Id

		ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
		// 镜像区域

		RegistryRegion *string `json:"RegistryRegion,omitempty" name:"RegistryRegion"`
		// 镜像创建的时间

		ImageCreateTime *string `json:"ImageCreateTime,omitempty" name:"ImageCreateTime"`
		// 敏感信息数

		SensitiveInfoCnt *uint64 `json:"SensitiveInfoCnt,omitempty" name:"SensitiveInfoCnt"`
		// Id

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyEventTendencyRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeImageDenyEventTendencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyEventTendencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusAutoIsolateSampleDetailRequest struct {
	*tchttp.BaseRequest

	// 文件MD5值

	MD5 *string `json:"MD5,omitempty" name:"MD5"`
	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeVirusAutoIsolateSampleDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSampleDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccessControlsRuleExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>RuleType - string  - 是否必填: 否 -规则类型</li>
	// <li>Status - string  - 是否必填: 否 -状态</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateAccessControlsRuleExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessControlsRuleExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeEventInfoRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeEscapeEventInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemVulListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>AppId - String - 是否必填：是 - 租户id筛选</li>
	// <li>OnlyAffectedContainer- string - 是否必填：否 - 仅展示影响容器的漏洞true,false</li>
	// <li>OnlyAffectedNewestImage-string - 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>AppId - String - 是否必填：是 - 租户id筛选</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	// <li>FocusOnType - string - 是否必填：否 -关注紧急度类型 。ALL :全部，SERIOUS_LEVEL： 严重和高危 ，IS_SUGGEST： 重点关注，POC_EXP 有Poc或Exp ，NETWORK_EXP: 远程Exp</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeSystemVulListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemVulListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellEventsExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// execle下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 任务ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellEventsExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellEventsExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserWorkloadKindsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserWorkloadKindsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserWorkloadKindsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageSimpleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像列表

		List []*AssetSimpleImageInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageSimpleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageSimpleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalRuleListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>RuleType - string  - 是否必填: 否 -规则类型</li>
	// <li>Status - string  - 是否必填: 否 -状态</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段。
	// <li>UpdateTime - string  - 是否必填: 否 -最后更新时间</li>
	// <li>EffectClusterCount - string  - 是否必填: 否 -影响集群数</li>

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeK8sApiAbnormalRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulRegistryImageListRequest struct {
	*tchttp.BaseRequest

	// 漏洞ID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// OnlyAffectedNewestImage bool 是否影响最新镜像
	// ImageDigest 镜像Digest，支持模糊查询
	// ImageId 镜像ID，支持模糊查询
	// Namespace 命名空间，支持模糊查询
	// ImageTag 镜像版本，支持模糊查询
	// InstanceName 实例名称，支持模糊查询
	// ImageName 镜像名，支持模糊查询
	// ImageRepoAddress 镜像地址，支持模糊查询

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeVulRegistryImageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulRegistryImageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostStartTheServiceRequest struct {
	*tchttp.BaseRequest

	// 主机id

	HostId *string `json:"HostId,omitempty" name:"HostId"`
	// 租户端id

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
	// true开启，false关闭

	StartTheService *bool `json:"StartTheService,omitempty" name:"StartTheService"`
}

func (r *ModifyHostStartTheServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostStartTheServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyEventListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像拦截列表

		List []*ImageDenyEvent `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageDenyEventListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetComponentListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 容器id

	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`
}

func (r *DescribeAssetComponentListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetComponentListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessControlEventInfo struct {
	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 命中规则名称

	MatchRuleName *string `json:"MatchRuleName,omitempty" name:"MatchRuleName"`
	// 生成时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 容器名

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 动作执行结果，   BEHAVIOR_NONE: 无
	//     BEHAVIOR_ALERT: 告警
	//     BEHAVIOR_RELEASE：放行
	//     BEHAVIOR_HOLDUP_FAILED:拦截失败
	//     BEHAVIOR_HOLDUP_SUCCESSED：拦截失败

	Behavior *string `json:"Behavior,omitempty" name:"Behavior"`
	// 状态0:未处理  “EVENT_UNDEAL”:事件未处理
	//     "EVENT_DEALED":事件已经处理
	//     "EVENT_INGNORE"：事件已经忽略

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件记录的唯一id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 事件类型， FILE_ABNORMAL_READ:文件异常读取

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 镜像id, 用于跳转

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 容器id, 用于跳转

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 事件解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件详细描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 命中策略id

	MatchRuleId *string `json:"MatchRuleId,omitempty" name:"MatchRuleId"`
	// 命中规则行为：
	// RULE_MODE_RELEASE 放行
	// RULE_MODE_ALERT  告警
	// RULE_MODE_HOLDUP 拦截

	MatchAction *string `json:"MatchAction,omitempty" name:"MatchAction"`
	// 命中规则进程信息

	MatchProcessPath *string `json:"MatchProcessPath,omitempty" name:"MatchProcessPath"`
	// 命中规则文件信息

	MatchFilePath *string `json:"MatchFilePath,omitempty" name:"MatchFilePath"`
	// 文件路径，包含名字

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 租户id

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 命中策略是否存在

	RuleExist *bool `json:"RuleExist,omitempty" name:"RuleExist"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 规则组id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 容器隔离状态

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 容器隔离子状态

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// 容器状态

	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
}

type ComplianceAssetPolicyItem struct {
	// 为客户分配的唯一的检测项的ID。

	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`
	// 检测项的原始ID

	BasePolicyItemId *uint64 `json:"BasePolicyItemId,omitempty" name:"BasePolicyItemId"`
	// 检测项的名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 检测项所属的类型的名称

	Category *string `json:"Category,omitempty" name:"Category"`
	// 所属的合规标准的ID

	BenchmarkStandardId *uint64 `json:"BenchmarkStandardId,omitempty" name:"BenchmarkStandardId"`
	// 所属的合规标准的名称

	BenchmarkStandardName *string `json:"BenchmarkStandardName,omitempty" name:"BenchmarkStandardName"`
	// 威胁等级

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 检测状态
	// CHECK_INIT, 待检测
	// CHECK_RUNNING, 检测中
	// CHECK_FINISHED, 检测完成
	// CHECK_FAILED, 检测失败

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 检测结果
	// RESULT_PASSED: 通过
	// RESULT_FAILED: 未通过

	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`
	// 检测项对应的白名单项的ID。如果存在且非0，表示检测项被用户忽略。

	WhitelistId *uint64 `json:"WhitelistId,omitempty" name:"WhitelistId"`
	// 处理建议。

	FixSuggestion *string `json:"FixSuggestion,omitempty" name:"FixSuggestion"`
	// 最近检测的时间。

	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`
}

type VulAffectedImageInfo struct {
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 关联的主机数

	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
	// 关联的容器数

	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
	// 组件列表

	ComponentList []*VulAffectedImageComponentInfo `json:"ComponentList,omitempty" name:"ComponentList"`
	// 租户AppId，运营端专用

	AppID *uint64 `json:"AppID,omitempty" name:"AppID"`
}

type DescribeIngressForwardConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Ingress的转发配置信息列表

		ConfigList []*IngressForwardConfig `json:"ConfigList,omitempty" name:"ConfigList"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIngressForwardConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIngressForwardConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterNameListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群名称列表

		ClusterNames []*string `json:"ClusterNames,omitempty" name:"ClusterNames"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterNameListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterNameListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateComponentExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateComponentExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateComponentExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceScanFailedAssetListRequest struct {
	*tchttp.BaseRequest

	// 资产类型列表。
	// ASSET_CONTAINER, 容器
	// ASSET_IMAGE, 镜像
	// ASSET_HOST, 主机
	// ASSET_K8S, K8S资产

	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`
	// 起始偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回的数据量，默认为10，最大为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 查询过滤器
	// AppId

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComplianceScanFailedAssetListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceScanFailedAssetListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEscapeEventsExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,Status：EVENT_UNDEAL:未处理，EVENT_DEALED:已处理，EVENT_INGNORE:忽略

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：latest_found_time

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateEscapeEventsExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEscapeEventsExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件查杀列表

		List []*VirusTaskInfo `json:"List,omitempty" name:"List"`
		// 总数量(容器任务数量)

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalRuleInfoRequest struct {
	*tchttp.BaseRequest

	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 当前appID

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeK8sApiAbnormalRuleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalRuleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessEventsExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// execle下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalProcessEventsExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessEventsExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsEventDetailRequest struct {
	*tchttp.BaseRequest

	// 事件ID

	EventID *uint64 `json:"EventID,omitempty" name:"EventID"`
	// 当前appID

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeRiskDnsEventDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsEventDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceContainerDetailInfo struct {
	// 容器在主机上的ID。

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 容器所属的Pod的名称。

	PodName *string `json:"PodName,omitempty" name:"PodName"`
}

type ComponentInfo struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeRiskDnsSummaryRequest struct {
	*tchttp.BaseRequest

	// 当前appID

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeRiskDnsSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCompliancePolicyItemAffectedAssetListRequest struct {
	*tchttp.BaseRequest

	// DescribeComplianceTaskPolicyItemSummaryList返回的CustomerPolicyItemId，表示检测项的ID。

	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`
	// 起始偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 需要返回的数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// Name - String
	// Name 可取值：NodeName, CheckResult,AppId

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCompliancePolicyItemAffectedAssetListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCompliancePolicyItemAffectedAssetListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProcessEventsExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProcessEventsExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProcessEventsExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件基本信息

		EventBaseInfo *RunTimeEventBaseInfo `json:"EventBaseInfo,omitempty" name:"EventBaseInfo"`
		// 进程信息

		ProcessInfo *ProcessDetailInfo `json:"ProcessInfo,omitempty" name:"ProcessInfo"`
		// 被篡改信息

		TamperedFileInfo *FileAttributeInfo `json:"TamperedFileInfo,omitempty" name:"TamperedFileInfo"`
		// 事件描述

		EventDetail *AccessControlEventDescription `json:"EventDetail,omitempty" name:"EventDetail"`
		// 父进程信息

		ParentProcessInfo *ProcessBaseInfo `json:"ParentProcessInfo,omitempty" name:"ParentProcessInfo"`
		// 祖先进程信息

		AncestorProcessInfo *ProcessBaseInfo `json:"AncestorProcessInfo,omitempty" name:"AncestorProcessInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessControlDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerSecEventSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 未处理逃逸事件

		UnhandledEscapeCnt *uint64 `json:"UnhandledEscapeCnt,omitempty" name:"UnhandledEscapeCnt"`
		// 未处理反弹shell事件

		UnhandledReverseShellCnt *uint64 `json:"UnhandledReverseShellCnt,omitempty" name:"UnhandledReverseShellCnt"`
		// 未处理高危系统调用

		UnhandledRiskSyscallCnt *uint64 `json:"UnhandledRiskSyscallCnt,omitempty" name:"UnhandledRiskSyscallCnt"`
		// 未处理异常进程

		UnhandledAbnormalProcessCnt *uint64 `json:"UnhandledAbnormalProcessCnt,omitempty" name:"UnhandledAbnormalProcessCnt"`
		// 未处理文件篡改

		UnhandledFileCnt *uint64 `json:"UnhandledFileCnt,omitempty" name:"UnhandledFileCnt"`
		// 未处理木马事件

		UnhandledVirusEventCnt *uint64 `json:"UnhandledVirusEventCnt,omitempty" name:"UnhandledVirusEventCnt"`
		// 未处理恶意外连事件

		UnhandledMaliciousConnectionEventCnt *uint64 `json:"UnhandledMaliciousConnectionEventCnt,omitempty" name:"UnhandledMaliciousConnectionEventCnt"`
		// 未处理k8sApi事件

		UnhandledK8sApiEventCnt *uint64 `json:"UnhandledK8sApiEventCnt,omitempty" name:"UnhandledK8sApiEventCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContainerSecEventSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerSecEventSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessRulesExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	// <li>AppId - String - 是否必填：否 - 租户id筛选</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeAbnormalProcessRulesExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessRulesExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetComponentListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件列表

		List []*ComponentInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetComponentListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetComponentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulAffectedComponentInfo struct {
	// 组件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件版本

	Version []*string `json:"Version,omitempty" name:"Version"`
	// 组件修复版本

	FixedVersion []*string `json:"FixedVersion,omitempty" name:"FixedVersion"`
}

type DescribeClusterNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 节点列表

		ClusterNodeList []*ClusterNodeInfo `json:"ClusterNodeList,omitempty" name:"ClusterNodeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageDenyEvent struct {
	// 事件类型 EVENT_RISK:风险事件类型，EVENT_PRIVILEGE:特权

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 规则RuleID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 规则类型

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 规则启用状态 0:开启，1:关闭

	RuleStatus *int64 `json:"RuleStatus,omitempty" name:"RuleStatus"`
	// 规则策略状态 IN_THE_TEST ：观察中，IN_EFFECT：生效中

	RuleEffectStatus *string `json:"RuleEffectStatus,omitempty" name:"RuleEffectStatus"`
	// 规则内容

	RuleInfo []*string `json:"RuleInfo,omitempty" name:"RuleInfo"`
	// 规则描述

	RuleDescription *string `json:"RuleDescription,omitempty" name:"RuleDescription"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 内网IP

	NodeIP *string `json:"NodeIP,omitempty" name:"NodeIP"`
	// 主机Quuid

	QUUID *string `json:"QUUID,omitempty" name:"QUUID"`
	// 首次生成时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 执行动作:
	// BEHAVIOR_ALERT:告警，
	// BEHAVIOR_HOLDUP_SUCCESSED:拦截

	DealBehavior *string `json:"DealBehavior,omitempty" name:"DealBehavior"`
	// 事件ID

	EventID *int64 `json:"EventID,omitempty" name:"EventID"`
	// 外网IP

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 租户appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type AccessControlChildRuleInfo struct {
	// 子策略id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 策略模式,  RULE_MODE_RELEASE: 放行
	//    RULE_MODE_ALERT: 告警
	//    RULE_MODE_HOLDUP:拦截

	RuleMode *string `json:"RuleMode,omitempty" name:"RuleMode"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 被访问文件路径，仅仅在访问控制生效

	TargetFilePath *string `json:"TargetFilePath,omitempty" name:"TargetFilePath"`
}

type DescribeAssetWebServiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// web服务列表

		List []*ServiceInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebServiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryVirusListExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryVirusListExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryVirusListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetImageRiskExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为10000

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
}

func (r *CreateAssetImageRiskExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetImageRiskExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostSecurityServiceStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostSecurityServiceStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostSecurityServiceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceAssetDetailInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 某资产的详情。

		AssetDetailInfo *ComplianceAssetDetailInfo `json:"AssetDetailInfo,omitempty" name:"AssetDetailInfo"`
		// 当资产为容器时，返回此字段。

		ContainerDetailInfo *ComplianceContainerDetailInfo `json:"ContainerDetailInfo,omitempty" name:"ContainerDetailInfo"`
		// 当资产为镜像时，返回此字段。

		ImageDetailInfo *ComplianceImageDetailInfo `json:"ImageDetailInfo,omitempty" name:"ImageDetailInfo"`
		// 当资产为主机时，返回此字段。

		HostDetailInfo *ComplianceHostDetailInfo `json:"HostDetailInfo,omitempty" name:"HostDetailInfo"`
		// 当资产为K8S时，返回此字段。

		K8SDetailInfo *ComplianceK8SDetailInfo `json:"K8SDetailInfo,omitempty" name:"K8SDetailInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComplianceAssetDetailInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceAssetDetailInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 高危系统调用事件信息

		EventSet []*RiskSyscallEventInfo `json:"EventSet,omitempty" name:"EventSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskSyscallEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCompliancePolicyItemAffectedAssetListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回各检测项所影响的资产的列表。

		AffectedAssetList []*ComplianceAffectedAsset `json:"AffectedAssetList,omitempty" name:"AffectedAssetList"`
		// 检测项影响的资产的总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCompliancePolicyItemAffectedAssetListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCompliancePolicyItemAffectedAssetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySystemPolicyRequest struct {
	*tchttp.BaseRequest

	// RUNTIME_ESCAPE:  容器逃逸；RUNTIME_PROCESS  :  异常进程；RUNTIME_FILE  :  文件篡改；RUNTIME_SYSCALL  :  高危系统调用

	SystemType *string `json:"SystemType,omitempty" name:"SystemType"`
	// 编辑的策略数组

	RuleSet []*ModifySystemRule `json:"RuleSet,omitempty" name:"RuleSet"`
}

func (r *ModifySystemPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySystemPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeEventTendencyRequest struct {
	*tchttp.BaseRequest

	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 当前appid

	CurrentAppId *string `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeEscapeEventTendencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventTendencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyRuleDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// 规则名称

		RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
		// 规则类型 RULE_RISK：风险， RULE_PRIVILEGE：特权

		RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
		// 生效的镜像数量

		EffectImageCount *int64 `json:"EffectImageCount,omitempty" name:"EffectImageCount"`
		// 是否对全部扫描镜像生效。0:全选镜像，1:自选镜像

		IsEffectAllImage *int64 `json:"IsEffectAllImage,omitempty" name:"IsEffectAllImage"`
		// 规则开始生效时间

		EffectTime *string `json:"EffectTime,omitempty" name:"EffectTime"`
		// 更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 操作用户

		OperationUin *string `json:"OperationUin,omitempty" name:"OperationUin"`
		// 生效状态 IN_THE_TEST ：观察中，IN_EFFECT：生效中

		EffectStatus *string `json:"EffectStatus,omitempty" name:"EffectStatus"`
		// 规则描述

		RuleDescription *string `json:"RuleDescription,omitempty" name:"RuleDescription"`
		// 启用状态 0:开启，1:关闭

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 漏洞，0:未选中，1:选中

		Vul *int64 `json:"Vul,omitempty" name:"Vul"`
		// cve编号

		CVEIDSet []*string `json:"CVEIDSet,omitempty" name:"CVEIDSet"`
		// 组件编号

		ComponentSet []*string `json:"ComponentSet,omitempty" name:"ComponentSet"`
		// 漏洞分类

		VulClassSet []*string `json:"VulClassSet,omitempty" name:"VulClassSet"`
		// 漏洞等级

		VulLevelSet []*string `json:"VulLevelSet,omitempty" name:"VulLevelSet"`
		// 漏洞标签

		VulLabelSet []*string `json:"VulLabelSet,omitempty" name:"VulLabelSet"`
		// 木马，0:未选中，1:选中

		Virus *int64 `json:"Virus,omitempty" name:"Virus"`
		// 木马md5列表

		VirusMD5Set []*string `json:"VirusMD5Set,omitempty" name:"VirusMD5Set"`
		// 木马等级

		VirusLevelSet []*string `json:"VirusLevelSet,omitempty" name:"VirusLevelSet"`
		// 病毒名

		VirusName []*string `json:"VirusName,omitempty" name:"VirusName"`
		// 敏感信息，0:未选中，1:选中

		Risk *int64 `json:"Risk,omitempty" name:"Risk"`
		// 敏感等级

		RiskLevelSet []*string `json:"RiskLevelSet,omitempty" name:"RiskLevelSet"`
		// 敏感信息分类

		RiskType []*string `json:"RiskType,omitempty" name:"RiskType"`
		// 特权启动 0:不允许，1:允许

		PrivilegeRun *int64 `json:"PrivilegeRun,omitempty" name:"PrivilegeRun"`
		// 特权类型,

		PrivilegeDetail []*string `json:"PrivilegeDetail,omitempty" name:"PrivilegeDetail"`
		// 镜像ID列表

		EffectImageSet []*string `json:"EffectImageSet,omitempty" name:"EffectImageSet"`
		// 多少天后生效

		EffectDay *uint64 `json:"EffectDay,omitempty" name:"EffectDay"`
		// 规则RuelD

		RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageDenyRuleDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusScanTimeoutSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 超时时长单位小时

		Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusScanTimeoutSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusScanTimeoutSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusAutoIsolateSampleListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>MD5- String - 是否必填：否 - md5 </li>
	// <li>AutoIsolateSwitch- String - 是否必填：否 - 自动隔离开关 </li>
	// <li>VirusName- String - 是否必填：否 - 病毒名 </li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeVirusAutoIsolateSampleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSampleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserWorkloadKindsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户工作负载类型列表

		WorkloadKinds []*string `json:"WorkloadKinds,omitempty" name:"WorkloadKinds"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserWorkloadKindsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserWorkloadKindsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceScanFailedAsset struct {
	// 客户资产的ID。

	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`
	// 资产类别。

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 检测状态
	// CHECK_INIT, 待检测
	// CHECK_RUNNING, 检测中
	// CHECK_FINISHED, 检测完成
	// CHECK_FAILED, 检测失败

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 资产的名称。

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产检测失败的原因。

	FailureReason *string `json:"FailureReason,omitempty" name:"FailureReason"`
	// 检测失败的处理建议。

	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`
	// 检测的时间。

	CheckTime *string `json:"CheckTime,omitempty" name:"CheckTime"`
	// 用户的AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type VulDetailInfo struct {
	// CVE编号

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// 漏洞名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 漏洞标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 漏洞类型

	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`
	// 漏洞威胁等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 漏洞披露时间

	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`
	// 漏洞描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// CVSS V3描述

	CVSSV3Desc *string `json:"CVSSV3Desc,omitempty" name:"CVSSV3Desc"`
	// 漏洞修复建议

	OfficialSolution *string `json:"OfficialSolution,omitempty" name:"OfficialSolution"`
	// 缓解措施

	DefenseSolution *string `json:"DefenseSolution,omitempty" name:"DefenseSolution"`
	// 参考链接

	Reference []*string `json:"Reference,omitempty" name:"Reference"`
	// CVSS V3分数

	CVSSV3Score *float64 `json:"CVSSV3Score,omitempty" name:"CVSSV3Score"`
	// 受漏洞影响的组件列表

	ComponentList []*VulAffectedComponentInfo `json:"ComponentList,omitempty" name:"ComponentList"`
	// 影响本地镜像数

	LocalImageCount *int64 `json:"LocalImageCount,omitempty" name:"LocalImageCount"`
	// 影响容器数

	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
	// 影响仓库镜像数

	RegistryImageCount *int64 `json:"RegistryImageCount,omitempty" name:"RegistryImageCount"`
	// 漏洞子类型

	Category *string `json:"Category,omitempty" name:"Category"`
	// 影响最新本地镜像数

	LocalNewestImageCount *int64 `json:"LocalNewestImageCount,omitempty" name:"LocalNewestImageCount"`
	// 影响最新仓库镜像数

	RegistryNewestImageCount *int64 `json:"RegistryNewestImageCount,omitempty" name:"RegistryNewestImageCount"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 防御状态，NO_DEFENDED:未防御，DEFENDED:已防御

	DefenceStatus *string `json:"DefenceStatus,omitempty" name:"DefenceStatus"`
	// 漏洞防御主机范围: MANUAL:自选主机节点，ALL:全部

	DefenceScope *string `json:"DefenceScope,omitempty" name:"DefenceScope"`
	// 漏洞防御主机数量

	DefenceHostCount *int64 `json:"DefenceHostCount,omitempty" name:"DefenceHostCount"`
	// 已防御攻击次数

	DefendedCount *int64 `json:"DefendedCount,omitempty" name:"DefendedCount"`
	// 是否已扫描，NOT_SCAN:未扫描,SCANNED:已扫描

	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`
}

type DescribeImageDenyEventDetailRequest struct {
	*tchttp.BaseRequest

	// 事件ID

	EventID *int64 `json:"EventID,omitempty" name:"EventID"`
	// 当前AppID

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeImageDenyEventDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyEventDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeESHitsRequest struct {
	*tchttp.BaseRequest

	// ES查询条件JSON

	Query *string `json:"Query,omitempty" name:"Query"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeESHitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeESHitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceTaskPolicyItemSummaryListRequest struct {
	*tchttp.BaseRequest

	// 资产类型。仅查询与指定资产类型相关的检测项。
	//
	// ASSET_CONTAINER, 容器
	//
	// ASSET_IMAGE, 镜像
	//
	// ASSET_HOST, 主机
	//
	// ASSET_K8S, K8S资产

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 起始偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 需要返回的数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件。
	// Name - String
	// Name 可取值：ItemType, StandardId,  RiskLevel。
	// 当为K8S资产时，还可取ClusterName。

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComplianceTaskPolicyItemSummaryListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceTaskPolicyItemSummaryListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusAutoIsolateSettingRequest struct {
	*tchttp.BaseRequest

	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeVirusAutoIsolateSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件类型

		SystemRuleSet []*RuntimeSystemRuleInfo `json:"SystemRuleSet,omitempty" name:"SystemRuleSet"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 剩余未下发策略的租户数量

		NotifyAppIdCount *uint64 `json:"NotifyAppIdCount,omitempty" name:"NotifyAppIdCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSystemPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusAutoIsolateSampleDownloadURLRequest struct {
	*tchttp.BaseRequest

	// 样本Md5值

	MD5 *string `json:"MD5,omitempty" name:"MD5"`
	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeVirusAutoIsolateSampleDownloadURLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSampleDownloadURLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerSecEventSummaryRequest struct {
	*tchttp.BaseRequest

	// 当前appId

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeContainerSecEventSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerSecEventSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulInfo struct {
	// 漏洞名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 漏洞标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// CVSS V3分数

	CVSSV3Score *float64 `json:"CVSSV3Score,omitempty" name:"CVSSV3Score"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// CVE编号

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// 漏洞子类型

	Category *string `json:"Category,omitempty" name:"Category"`
	// 首次发现时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 最近发现时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 漏洞ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 影响本地镜像数

	LocalImageCount *int64 `json:"LocalImageCount,omitempty" name:"LocalImageCount"`
	// 影响容器数

	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
	// 影响仓库镜像数

	RegistryImageCount *int64 `json:"RegistryImageCount,omitempty" name:"RegistryImageCount"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 防御状态，NO_DEFENDED:未防御，DEFENDED:已防御

	DefenceStatus *string `json:"DefenceStatus,omitempty" name:"DefenceStatus"`
	// 漏洞防御主机范围: MANUAL:自选主机节点，ALL:全部

	DefenceScope *string `json:"DefenceScope,omitempty" name:"DefenceScope"`
	// 漏洞防御主机数量

	DefenceHostCount *int64 `json:"DefenceHostCount,omitempty" name:"DefenceHostCount"`
	// 已防御攻击次数

	DefendedCount *int64 `json:"DefendedCount,omitempty" name:"DefendedCount"`
	// 租户appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type DescribeVirusSampleDownloadUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 样本下载地址

		FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusSampleDownloadUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusSampleDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEscapeWhiteListExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>EventType- String - 是否必填：否 - 加白事件类型，ESCAPE_CGROUPS：利用cgroup机制逃逸，ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸， ESCAPE_DOCKER_API：访问Docker API接口逃逸，ESCAPE_VUL_OCCURRED：逃逸漏洞利用，MOUNT_SENSITIVE_PTAH：敏感路径挂载，PRIVILEGE_CONTAINER_START：特权容器， PRIVILEGE：程序提权逃逸</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10000，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：主机数量：HostCount，容器数量：ContainerCount，更新时间：UpdateTime

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateEscapeWhiteListExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEscapeWhiteListExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetImageVirusExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Name- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>RiskLevel - String - 是否必填：否 - 风险等级  1,2,3,4</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 列表支持字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 需要返回的数量，默认为10，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *CreateAssetImageVirusExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetImageVirusExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellWhiteListsExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeReverseShellWhiteListsExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellWhiteListsExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulRegistryImageExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVulRegistryImageExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulRegistryImageExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EscapeEventDescription struct {
	// 事件规则

	Description *string `json:"Description,omitempty" name:"Description"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 事件最后一次处理的时间

	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
}

type VirusAutoIsolateSampleInfo struct {
	// 文件MD5值

	MD5 *string `json:"MD5,omitempty" name:"MD5"`
	// 病毒名

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
	// 最近编辑时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 自动隔离开关(true:开 false:关)

	AutoIsolateSwitch *bool `json:"AutoIsolateSwitch,omitempty" name:"AutoIsolateSwitch"`
	// 当前appID

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type DescribeAssetPortListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>AppId - String - 是否必填：否 - 租户id筛选</li>
	// <li>All - String - 是否必填：否 - 模糊查询可选字段</li>
	// <li>RunAs - String - 是否必填：否 - 运行用户筛选，</li>
	// <li>ContainerID - String - 是否必填：否 - 容器id</li>
	// <li>HostID- String - 是否必填：是 - 主机id</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>ProcessName- string - 是否必填：否 - 进程名搜索</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetPortListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPortListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusAutoIsolateSampleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 木马自动隔离样本列表

		List []*VirusAutoIsolateSampleInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusAutoIsolateSampleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSampleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetDBServiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// db服务列表

		List []*ServiceInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetDBServiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDBServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallWhiteListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 高危系统调用白名单列表

		WhiteListSet []*RiskSyscallWhiteListBaseInfo `json:"WhiteListSet,omitempty" name:"WhiteListSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskSyscallWhiteListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallWhiteListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComponentsInfo struct {
	// 组件名称

	Component *string `json:"Component,omitempty" name:"Component"`
	// 组件版本信息

	Version *string `json:"Version,omitempty" name:"Version"`
	// 组件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 可修复版本

	FixedVersion *string `json:"FixedVersion,omitempty" name:"FixedVersion"`
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 组件类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type RuleBaseInfo struct {
	// true: 默认策略，false:自定义策略

	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
	// 策略生效镜像数量

	EffectImageCount *uint64 `json:"EffectImageCount,omitempty" name:"EffectImageCount"`
	// 策略Id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 策略更新时间, 存在为空的情况

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 策略名字

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 编辑用户名称

	EditUserName *string `json:"EditUserName,omitempty" name:"EditUserName"`
	// true: 策略启用，false：策略禁用

	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
	// 租户端id

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type DescribeCheckItemListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检查项详情数组

		ClusterCheckItems []*ClusterCheckItem `json:"ClusterCheckItems,omitempty" name:"ClusterCheckItems"`
		// 检查项总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCheckItemListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCheckItemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageDenyEventExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>EventType- String - 是否必填：否 -事件类型 EVENT_RISK:风险事件类型，EVENT_PRIVILEGE:特权。</li>
	// <li>DealBehavior- String - 是否必填：否 - 执行动作,BEHAVIOR_ALERT:告警，BEHAVIOR_HOLDUP_SUCCESSED:拦截。</li>
	// <li>RuleName- string - 是否必填：否 - 规则名称。</li>
	// <li>NodeName- string - 是否必填：否 - 节点名称。</li>
	// <li>NodeIP- string - 是否必填：否 - 内外IP。</li>
	// <li>PublicIP- string - 是否必填：否 - 外网IP。</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	// <li>TimeRange- String -是否必填: 否 -  时间范围，第一个值表示开始时间，第二个值表示结束时间 </li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10000，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：告警数量：EventCount，最近生成时间：LatestFoundTime

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateImageDenyEventExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageDenyEventExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallNamesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 系统调用名称列表

		SyscallNames []*string `json:"SyscallNames,omitempty" name:"SyscallNames"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskSyscallNamesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessEventTendencyRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeAbnormalProcessEventTendencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessEventTendencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterDetailExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务Id

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterDetailExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterDetailExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostSecurityServiceStateRequestItems struct {
	// appid

	AppID *uint64 `json:"AppID,omitempty" name:"AppID"`
	// 主机ID

	HostID *string `json:"HostID,omitempty" name:"HostID"`
}

type DescribeReverseShellWhiteListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 白名单信息列表

		WhiteListSet []*ReverseShellWhiteListBaseInfo `json:"WhiteListSet,omitempty" name:"WhiteListSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellWhiteListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellWhiteListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterNameListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ClusterName

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeClusterNameListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterNameListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusManualScanEstimateTimeoutRequest struct {
	*tchttp.BaseRequest

	// 扫描范围0容器1主机节点

	ScanRangeType *uint64 `json:"ScanRangeType,omitempty" name:"ScanRangeType"`
	// true 全选，false 自选

	ScanRangeAll *bool `json:"ScanRangeAll,omitempty" name:"ScanRangeAll"`
	// 自选扫描范围的容器id或者主机id 根据ScanRangeType决定

	ScanIds []*string `json:"ScanIds,omitempty" name:"ScanIds"`
	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeVirusManualScanEstimateTimeoutRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusManualScanEstimateTimeoutRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AbnormalProcessRuleInfo struct {
	// 策略id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// true:策略启用，false:策略禁用

	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
	// 生效惊现id，空数组代表全部镜像

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
	// 子策略数组

	ChildRules []*AbnormalProcessChildRuleInfo `json:"ChildRules,omitempty" name:"ChildRules"`
	// 策略名字

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 系统策略的子策略数组

	SystemChildRules []*AbnormalProcessSystemChildRuleInfo `json:"SystemChildRules,omitempty" name:"SystemChildRules"`
	// 是否是系统默认策略

	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

type RunTimeTendencyInfo struct {
	// 时间

	CurTime *string `json:"CurTime,omitempty" name:"CurTime"`
	// 当前数量

	Cnt *uint64 `json:"Cnt,omitempty" name:"Cnt"`
}

type DescribePodContainersRequest struct {
	*tchttp.BaseRequest

	// Pod名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// Name 可取值：ClusterId集群id,Namespace命名空间等

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribePodContainersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePodContainersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnfinishRefreshTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回最近的一次任务ID

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 任务状态，为Task_New,Task_Running,Task_Finish,Task_Error,Task_NoExist.Task_New,Task_Running表示有任务存在，不允许新下发

		TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
		// 新任务ID

		NewTaskID *string `json:"NewTaskID,omitempty" name:"NewTaskID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUnfinishRefreshTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnfinishRefreshTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddImageDenyRuleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否存在特权拦截规则

		ExistPrivilegeDenyRule *bool `json:"ExistPrivilegeDenyRule,omitempty" name:"ExistPrivilegeDenyRule"`
		// 已扫描的镜像数量

		ScannedImageCount *int64 `json:"ScannedImageCount,omitempty" name:"ScannedImageCount"`
		// 特权拦截规则RuleID

		PrivilegeDenyRuleID *string `json:"PrivilegeDenyRuleID,omitempty" name:"PrivilegeDenyRuleID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddImageDenyRuleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddImageDenyRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHostExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - agent状态筛选，"ALL":"全部"(或不传该字段),"UNINSTALL"："未安装","OFFLINE"："离线", "ONLINE"："防护中"</li>
	// <li>HostName - String - 是否必填：否 - 主机名筛选</li>
	// <li>Group- String - 是否必填：否 - 主机群组搜索</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>HostID- string - 是否必填：否 - 主机id搜索</li>
	// <li>DockerVersion- string - 是否必填：否 - docker版本搜索</li>
	// <li>MachineType- string - 是否必填：否 - 主机来源MachineType搜索，"ALL":"全部"(或不传该字段),主机来源：["CVM", "ECM", "LH", "BM"]  中的之一为腾讯云服务器；["Other"]之一非腾讯云服务器；</li>
	// <li>DockerStatus- string - 是否必填：否 - docker安装状态，"ALL":"全部"(或不传该字段),"INSTALL":"已安装","UNINSTALL":"未安装"</li>
	// <li>ProjectID- string - 是否必填：否 - 所属项目id搜索</li>
	// <li>Tag:xxx(tag:key)- string- 是否必填：否 - 标签键值搜索 示例Filters":[{"Name":"tag:tke-kind","Values":["service"]}]</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 需要返回的数量，默认为10，最大值为10000

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 租户id，当前查询的租户Id

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *CreateHostExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHostExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebServiceListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>AppId - String - 是否必填：否 - 租户id筛选</li>
	// <li>Keywords- String - 是否必填：否 - 模糊查询可选字段</li>
	// <li>Type- String - 是否必填：否 - 主机运行状态筛选，"Apache"
	// "Jboss"
	// "lighttpd"
	// "Nginx"
	// "Tomcat"</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetWebServiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebServiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AbnormalProcessEventInfo struct {
	// 进程目录

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 事件类型，MALICE_PROCESS_START:恶意进程启动

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 命中规则

	MatchRuleName *string `json:"MatchRuleName,omitempty" name:"MatchRuleName"`
	// 生成时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 容器名

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 动作执行结果，    BEHAVIOR_NONE: 无
	//     BEHAVIOR_ALERT: 告警
	//     BEHAVIOR_RELEASE：放行
	//     BEHAVIOR_HOLDUP_FAILED:拦截失败
	//     BEHAVIOR_HOLDUP_SUCCESSED：拦截失败

	Behavior *string `json:"Behavior,omitempty" name:"Behavior"`
	// 状态，EVENT_UNDEAL:事件未处理
	//     EVENT_DEALED:事件已经处理
	//     EVENT_INGNORE：事件已经忽略

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件记录的唯一id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 镜像id，用于跳转

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 容器id，用于跳转

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 事件解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件详细描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 命中策略id

	MatchRuleId *string `json:"MatchRuleId,omitempty" name:"MatchRuleId"`
	// 命中规则行为：
	// RULE_MODE_RELEASE 放行
	// RULE_MODE_ALERT  告警
	// RULE_MODE_HOLDUP 拦截

	MatchAction *string `json:"MatchAction,omitempty" name:"MatchAction"`
	// 命中规则进程信息

	MatchProcessPath *string `json:"MatchProcessPath,omitempty" name:"MatchProcessPath"`
	// 租户id

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 命中策略是否存在

	RuleExist *bool `json:"RuleExist,omitempty" name:"RuleExist"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 规则组Id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 命中策略名称：SYSTEM_DEFINED_RULE （系统策略）或  用户自定义的策略名字

	MatchGroupName *string `json:"MatchGroupName,omitempty" name:"MatchGroupName"`
	// 命中规则等级，HIGH：高危，MIDDLE：中危，LOW：低危。

	MatchRuleLevel *string `json:"MatchRuleLevel,omitempty" name:"MatchRuleLevel"`
	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// "NODE_DESTROYED"      //节点已销毁
	// "CONTAINER_EXITED"    //容器已退出
	// "CONTAINER_DESTROYED" //容器已销毁
	// "SHARED_HOST"         // 容器与主机共享网络
	// "RESOURCE_LIMIT"      //隔离操作资源超限
	// "UNKNOW"              // 原因未知

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离操作来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING

	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
}

type ContainerInfo struct {
	// 容器id

	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 容器运行状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 运行用户

	RunAs *string `json:"RunAs,omitempty" name:"RunAs"`
	// 命令行

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// cpu 使用率 *1000

	CPUUsage *uint64 `json:"CPUUsage,omitempty" name:"CPUUsage"`
	// 内存使用 kb

	RamUsage *uint64 `json:"RamUsage,omitempty" name:"RamUsage"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像id

	POD *string `json:"POD,omitempty" name:"POD"`
	// 主机id

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 主机ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 租户AppId，运营端专用

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 外网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 网络状态 未隔离 NORMAL 已隔离 ISOLATED 隔离中 ISOLATING 隔离失败 ISOLATE_FAILED 解除隔离中 RESTORING 解除隔离失败 RESTORE_FAILED

	NetStatus *string `json:"NetStatus,omitempty" name:"NetStatus"`
	// 网络子状态

	NetSubStatus *string `json:"NetSubStatus,omitempty" name:"NetSubStatus"`
	// 隔离来源

	IsolateSource *string `json:"IsolateSource,omitempty" name:"IsolateSource"`
	// 隔离时间

	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`
	// 超级节点id

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// podip

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// podip

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 节点类型:节点类型：NORMAL普通节点、SUPER超级节点

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 超级节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 所属Pod的CPU

	PodCpu *int64 `json:"PodCpu,omitempty" name:"PodCpu"`
	// 所属Pod的内存

	PodMem *int64 `json:"PodMem,omitempty" name:"PodMem"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// pod uid

	PodUid *string `json:"PodUid,omitempty" name:"PodUid"`
}

type SearchTemplate struct {
	// 规则ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 检索名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 检索索引类型

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 检索语句

	Condition *string `json:"Condition,omitempty" name:"Condition"`
	// 时间范围

	TimeRange *string `json:"TimeRange,omitempty" name:"TimeRange"`
	// 转换的检索语句内容

	Query *string `json:"Query,omitempty" name:"Query"`
	// 检索方式。输入框检索：standard,过滤，检索：simple

	Flag *string `json:"Flag,omitempty" name:"Flag"`
	// 展示数据

	DisplayData *string `json:"DisplayData,omitempty" name:"DisplayData"`
}

type DescribeImageDenyRuleDetailRequest struct {
	*tchttp.BaseRequest

	// 规则RuleID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 规则所属用户的appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeImageDenyRuleDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyRuleDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalEventInfoRequest struct {
	*tchttp.BaseRequest

	// 事件ID

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 当前APPID

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeK8sApiAbnormalEventInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalEventInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRiskSyscallWhiteListsExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRiskSyscallWhiteListsExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRiskSyscallWhiteListsExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceAffectedAsset struct {
	// 为客户分配的唯一的资产项的ID。

	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`
	// 资产项的名称。

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产项的类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 检测状态
	//
	// CHECK_INIT, 待检测
	//
	// CHECK_RUNNING, 检测中
	//
	// CHECK_FINISHED, 检测完成
	//
	// CHECK_FAILED, 检测失败

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 节点名称。

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 上次检测的时间，格式为”YYYY-MM-DD HH:m::SS“。

	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`
	// 检测结果。取值为：
	//
	// RESULT_FAILED: 未通过
	//
	// RESULT_PASSED: 通过

	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`
	// 主机IP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 镜像的tag

	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
	// 用户的AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 验证消息

	VerifyInfo *string `json:"VerifyInfo,omitempty" name:"VerifyInfo"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ProcessBaseInfo struct {
	// 进程启动用户

	ProcessStartUser *string `json:"ProcessStartUser,omitempty" name:"ProcessStartUser"`
	// 进程用户组

	ProcessUserGroup *string `json:"ProcessUserGroup,omitempty" name:"ProcessUserGroup"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 进程命令行参数

	ProcessParam *string `json:"ProcessParam,omitempty" name:"ProcessParam"`
}

type DescribeVirusManualScanEstimateTimeoutResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 预估超时时间(h)

		Timeout *float64 `json:"Timeout,omitempty" name:"Timeout"`
		// 单台主机并行扫描容器数

		ContainerScanConcurrencyCount *uint64 `json:"ContainerScanConcurrencyCount,omitempty" name:"ContainerScanConcurrencyCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusManualScanEstimateTimeoutResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusManualScanEstimateTimeoutResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterRegionItem struct {
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域缩写

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
	// 地域Id

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域中文名

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域英文名

	RegionNameEn *string `json:"RegionNameEn,omitempty" name:"RegionNameEn"`
}

type VirusAutoIsolateSettingInfo struct {
	// 自动隔离开关(true:开 false:关)

	AutoIsolateSwitch *bool `json:"AutoIsolateSwitch,omitempty" name:"AutoIsolateSwitch"`
	// 是否中断隔离文件关联的进程(true:是 false:否)

	IsKillProgress *bool `json:"IsKillProgress,omitempty" name:"IsKillProgress"`
	// 当前AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type DescribeMaliciousConnectionRuleSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 黑名单内容数量

		BlackListCount *uint64 `json:"BlackListCount,omitempty" name:"BlackListCount"`
		// 黑名单域名数量

		BlackListDomainCount *uint64 `json:"BlackListDomainCount,omitempty" name:"BlackListDomainCount"`
		// 黑名单IP数量

		BlackListIPCount *uint64 `json:"BlackListIPCount,omitempty" name:"BlackListIPCount"`
		// 白名单内容数量

		WhiteListCount *uint64 `json:"WhiteListCount,omitempty" name:"WhiteListCount"`
		// 白单域名数量

		WhiteListDomainCount *uint64 `json:"WhiteListDomainCount,omitempty" name:"WhiteListDomainCount"`
		// 白名单IP数量

		WhiteListIPCount *uint64 `json:"WhiteListIPCount,omitempty" name:"WhiteListIPCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMaliciousConnectionRuleSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaliciousConnectionRuleSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserClusterExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务唯一ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserClusterExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserClusterExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistrySummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像个数

		ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`
		// 有风险的镜像个数

		ImageHasRiskInfoCnt *uint64 `json:"ImageHasRiskInfoCnt,omitempty" name:"ImageHasRiskInfoCnt"`
		// 有病毒的镜像个数

		ImageHasVirusCnt *uint64 `json:"ImageHasVirusCnt,omitempty" name:"ImageHasVirusCnt"`
		// 有漏洞的镜像个数

		ImageHasVulsCnt *uint64 `json:"ImageHasVulsCnt,omitempty" name:"ImageHasVulsCnt"`
		// 不受信任的镜像个数

		ImageUntrustCnt *uint64 `json:"ImageUntrustCnt,omitempty" name:"ImageUntrustCnt"`
		// 最近镜像扫描时间

		LatestImageScanTime *string `json:"LatestImageScanTime,omitempty" name:"LatestImageScanTime"`
		// 风险镜像个数

		ImageUnsafeCnt *uint64 `json:"ImageUnsafeCnt,omitempty" name:"ImageUnsafeCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistrySummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistrySummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccessControlsRuleExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccessControlsRuleExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAccessControlsRuleExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusScanSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否开启定期扫描

		EnableScan *bool `json:"EnableScan,omitempty" name:"EnableScan"`
		// 检测周期每隔多少天

		Cycle *uint64 `json:"Cycle,omitempty" name:"Cycle"`
		// 扫描开始时间

		BeginScanAt *string `json:"BeginScanAt,omitempty" name:"BeginScanAt"`
		// 扫描全部路径

		ScanPathAll *bool `json:"ScanPathAll,omitempty" name:"ScanPathAll"`
		// 当ScanPathAll为true 生效 0扫描以下路径 1、扫描除以下路径

		ScanPathType *uint64 `json:"ScanPathType,omitempty" name:"ScanPathType"`
		// 超时时长，单位小时

		Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
		// 扫描范围0容器1主机节点

		ScanRangeType *uint64 `json:"ScanRangeType,omitempty" name:"ScanRangeType"`
		// true 全选，false 自选

		ScanRangeAll *bool `json:"ScanRangeAll,omitempty" name:"ScanRangeAll"`
		// 自选扫描范围的容器id或者主机id 根据ScanRangeType决定

		ScanIds []*string `json:"ScanIds,omitempty" name:"ScanIds"`
		// 自选排除或扫描的地址

		ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
		// 一键检测的超时设置

		ClickTimeout *uint64 `json:"ClickTimeout,omitempty" name:"ClickTimeout"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusScanSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusScanSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageBindRuleInfoRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"EventType","Values":[""]}]
	// EventType取值：
	// "FILE_ABNORMAL_READ" 访问控制
	// "MALICE_PROCESS_START" 恶意进程启动
	// <li>AppId - String - 是否必填：是 - 租户id筛选</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAssetImageBindRuleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageBindRuleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageComponent struct {
	// 组件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 组件路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 组件类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 组件漏洞数量

	VulCount *uint64 `json:"VulCount,omitempty" name:"VulCount"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 用户appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type DescribeVirusListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>FileName - String - 是否必填：否 - 文件名称</li>
	// <li>FilePath - String - 是否必填：否 - 文件路径</li>
	// <li>VirusName - String - 是否必填：否 - 病毒名称</li>
	// <li>ContainerName- String - 是否必填：是 - 容器名称</li>
	// <li>ContainerId- string - 是否必填：否 - 容器id</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称</li>
	// <li>ImageId- string - 是否必填：否 - 镜像id</li>
	// <li>IsRealTime- int - 是否必填：否 - 过滤是否实时监控数据</li>
	// <li>TaskId- string - 是否必填：否 - 任务ID</li>
	// <li>ContainerNetStatus - String -是否必填: 否 -  容器网络状态筛选 NORMAL ISOLATED ISOLATING RESTORING RESTORE_FAILED</li>
	// <li>TimeRange - string -是否必填: 否 - 时间范围筛选 ["2022-03-31 16:55:00", "2022-03-31 17:00:00"]</li>
	// <li>ContainerStatus - string -是否必填: 否 - 容器状态 RUNNING PAUSED STOPPED CREATED DESTROYED RESTARTING REMOVING</li>
	// <li>AutoIsolateMode - string -是否必填: 否 - 隔离方式 MANUAL AUTO</li>
	// <li>MD5 - string -是否必填: 否 - md5 </li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeVirusListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞总数量

		VulTotalCount *int64 `json:"VulTotalCount,omitempty" name:"VulTotalCount"`
		// 严重及高危漏洞数量

		SeriousVulCount *int64 `json:"SeriousVulCount,omitempty" name:"SeriousVulCount"`
		// 重点关注漏洞数量

		SuggestVulCount *int64 `json:"SuggestVulCount,omitempty" name:"SuggestVulCount"`
		// 有Poc或者Exp的漏洞数量

		PocExpLevelVulCount *int64 `json:"PocExpLevelVulCount,omitempty" name:"PocExpLevelVulCount"`
		// 有远程Exp的漏洞数量

		RemoteExpLevelVulCount *int64 `json:"RemoteExpLevelVulCount,omitempty" name:"RemoteExpLevelVulCount"`
		// 受严重或高危漏洞影响的最新版本镜像数

		SeriousVulNewestImageCount *int64 `json:"SeriousVulNewestImageCount,omitempty" name:"SeriousVulNewestImageCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAffectedWorkloadListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 受影响的workload列表数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 受影响的workload列表

		AffectedWorkloadList []*AffectedWorkloadItem `json:"AffectedWorkloadList,omitempty" name:"AffectedWorkloadList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAffectedWorkloadListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAffectedWorkloadListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlDetailRequest struct {
	*tchttp.BaseRequest

	// 事件唯一id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 租户id

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeAccessControlDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessEventTendencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 待处理异常进程事件趋势

		List []*AbnormalProcessEventTendencyInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalProcessEventTendencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessEventTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PodContainerInfo struct {
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 运行状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// CPU使用率

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 内存使用 KB

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 重启次数

	RestartCount *uint64 `json:"RestartCount,omitempty" name:"RestartCount"`
}

type RiskSyscallEventInfo struct {
	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 容器名

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 首次生成时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 事件解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件详细描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 系统调用名称

	SyscallName *string `json:"SyscallName,omitempty" name:"SyscallName"`
	// 状态，EVENT_UNDEAL;事件未处理

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 实例的名字

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 系统监控名称是否存在

	RuleExist *bool `json:"RuleExist,omitempty" name:"RuleExist"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// AppId值

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// "NODE_DESTROYED"      //节点已销毁
	// "CONTAINER_EXITED"    //容器已退出
	// "CONTAINER_DESTROYED" //容器已销毁
	// "SHARED_HOST"         // 容器与主机共享网络
	// "RESOURCE_LIMIT"      //隔离操作资源超限
	// "UNKNOW"              // 原因未知

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离操作来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING

	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
}

type DescribeEscapeEventsExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// execle下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEscapeEventsExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventsExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallWhiteListDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 白名单基本信息

		WhiteListDetailInfo *RiskSyscallWhiteListInfo `json:"WhiteListDetailInfo,omitempty" name:"WhiteListDetailInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskSyscallWhiteListDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallWhiteListDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusSampleDownloadUrlRequest struct {
	*tchttp.BaseRequest

	// 木马id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeVirusSampleDownloadUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusSampleDownloadUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 待处理的事件总数

		UnhandledEventCount *int64 `json:"UnhandledEventCount,omitempty" name:"UnhandledEventCount"`
		// 待处理的恶意域名请求数

		RiskDnsEventCount *int64 `json:"RiskDnsEventCount,omitempty" name:"RiskDnsEventCount"`
		// 待处理的恶意IP请求数

		RiskIPEventCount *int64 `json:"RiskIPEventCount,omitempty" name:"RiskIPEventCount"`
		// 影响的容器总数

		EffectContainerCount *int64 `json:"EffectContainerCount,omitempty" name:"EffectContainerCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskDnsSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVirusExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>FileName - String - 是否必填：否 - 文件名称</li>
	// <li>FilePath - String - 是否必填：否 - 文件路径</li>
	// <li>VirusName - String - 是否必填：否 - 病毒名称</li>
	// <li>ContainerName- String - 是否必填：是 - 容器名称</li>
	// <li>ContainerId- string - 是否必填：否 - 容器id</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称</li>
	// <li>ImageId- string - 是否必填：否 - 镜像id</li>
	// <li>IsRealTime- int - 是否必填：否 - 过滤是否实时监控数据</li>
	// <li>TaskId- string - 是否必填：否 - 任务ID</li>
	// <li>TimeRange - string -是否必填: 否 - 时间范围筛选 ["2022-03-31 16:55:00", "2022-03-31 17:00:00"]</li>
	// <li>ContainerNetStatus - String -是否必填: 否 -  容器网络状态筛选 NORMAL ISOLATED ISOLATING RESTORING RESTORE_FAILED</li>
	// <li>ContainerStatus - string -是否必填: 否 - 容器状态 RUNNING PAUSED STOPPED CREATED DESTROYED RESTARTING REMOVING</li>
	// <li>AutoIsolateMode - string -是否必填: 否 - 隔离方式 MANUAL AUTO</li>
	// <li>MD5 - string -是否必填: 否 - md5 </li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateVirusExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVirusExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRiskTendencyRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 当前AppId

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeImageRiskTendencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRiskTendencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceAssetInfo struct {
	// 客户资产的ID。

	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`
	// 资产类别。

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 资产的名称。

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 当资产为镜像时，这个字段为镜像Tag。

	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`
	// 资产所在的主机IP。

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 资产所属的节点的名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 检测状态
	//
	// CHECK_INIT, 待检测
	//
	// CHECK_RUNNING, 检测中
	//
	// CHECK_FINISHED, 检测完成
	//
	// CHECK_FAILED, 检测失败

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 此类资产通过的检测项的数目。

	PassedPolicyItemCount *uint64 `json:"PassedPolicyItemCount,omitempty" name:"PassedPolicyItemCount"`
	// 此类资产未通过的检测的数目。

	FailedPolicyItemCount *uint64 `json:"FailedPolicyItemCount,omitempty" name:"FailedPolicyItemCount"`
	// 上次检测的时间。

	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`
	// 检测结果：
	// RESULT_FAILED: 未通过。
	// RESULT_PASSED: 通过。

	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`
	// 用户的AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type ScanIgnoreVul struct {
	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
	// 漏洞CVEID

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 忽略的仓库镜像数

	RegistryImageCount *int64 `json:"RegistryImageCount,omitempty" name:"RegistryImageCount"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 是否忽略所有镜像：0：否/1：是

	IsIgnoreAll *int64 `json:"IsIgnoreAll,omitempty" name:"IsIgnoreAll"`
	// 忽略的本地镜像数

	LocalImageCount *int64 `json:"LocalImageCount,omitempty" name:"LocalImageCount"`
	// 用户appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type DescribeAssetImageVirusListExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageVirusListExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageVirusListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySystemPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySystemPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySystemPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusSummaryRequest struct {
	*tchttp.BaseRequest

	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeVirusSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>EventStatus- String - 是否必填：否 - 事件状态，待处理：EVENT_UNDEAL，EVENT_DEALED：已处理，已忽略：EVENT_IGNORE， EVENT_ADD_WHITE：已加白</li>
	// <li>ContainerStatus- String - 是否必填：否 - 容器运行状态筛选，已创建：CREATED,正常运行：RUNNING, 暂定运行：PAUSED, 停止运行：	STOPPED，重启中：RESTARTING, 迁移中：REMOVING, 销毁：DESTROYED </li>
	// <li>ContainerNetStatus- String -是否必填: 否 -  容器网络状态筛选 未隔离：NORMAL，已隔离：ISOLATED，隔离失败：ISOLATE_FAILED，解除隔离失败：RESTORE_FAILED，解除隔离中：RESTORING，隔离中：ISOLATING</li>
	// <li>EventType - String -是否必填: 否 -  事件类型，恶意域名请求：DOMAIN，恶意IP请求：IP</li>
	// <li>TimeRange- String -是否必填: 否 -  时间范围，第一个值表示开始时间，第二个值表示结束时间 </li>
	// <li>RiskDns- string - 是否必填：否 - 恶意域名。</li>
	// <li>RiskIP- string - 是否必填：否 - 恶意IP。</li>
	// <li>ContainerName- string - 是否必填：否 - 容器名称。</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID。</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	// <li>HostName- string - 是否必填：否 - 主机名称。</li>
	// <li>HostIP- string - 是否必填：否 - 内网IP。</li>
	// <li>PublicIP- string - 是否必填：否 - 外网IP。</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：告警数量：EventCount，最近生成时间：LatestFoundTime

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeRiskDnsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallWhiteListsExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数，"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeRiskSyscallWhiteListsExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallWhiteListsExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateK8sApiAbnormalRuleExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateK8sApiAbnormalRuleExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateK8sApiAbnormalRuleExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageVirusListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像病毒列表

		List []*ImageVirusInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 病毒扫描状态
		// 0:未扫描
		// 1:扫描中
		// 2:扫描完成
		// 3:扫描出错
		// 4:扫描取消

		VirusScanStatus *uint64 `json:"VirusScanStatus,omitempty" name:"VirusScanStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageVirusListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageVirusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistrySummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetImageRegistrySummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistrySummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProjectInfo struct {
	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 项目ID

	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`
}

type DescribeClusterDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群id

		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
		// 集群名字

		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
		// 当前集群扫描任务的进度，100表示扫描完成.

		ScanTaskProgress *int64 `json:"ScanTaskProgress,omitempty" name:"ScanTaskProgress"`
		// 集群版本

		ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
		// 运行时组件

		ContainerRuntime *string `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`
		// 集群节点数

		ClusterNodeNum *uint64 `json:"ClusterNodeNum,omitempty" name:"ClusterNodeNum"`
		// 集群状态 (Running 运行中 Creating 创建中 Abnormal 异常 )

		ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
		// 集群类型：为托管集群MANAGED_CLUSTER、独立集群INDEPENDENT_CLUSTER

		ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
		// 集群区域

		Region *string `json:"Region,omitempty" name:"Region"`
		// 严重风险检查项的数量

		SeriousRiskCount *uint64 `json:"SeriousRiskCount,omitempty" name:"SeriousRiskCount"`
		// 高风险检查项的数量

		HighRiskCount *uint64 `json:"HighRiskCount,omitempty" name:"HighRiskCount"`
		// 中风险检查项的数量

		MiddleRiskCount *uint64 `json:"MiddleRiskCount,omitempty" name:"MiddleRiskCount"`
		// 提示风险检查项的数量

		HintRiskCount *uint64 `json:"HintRiskCount,omitempty" name:"HintRiskCount"`
		// 检查任务的状态

		CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
		// 防御容器状态

		DefenderStatus *string `json:"DefenderStatus,omitempty" name:"DefenderStatus"`
		// 扫描任务创建时间

		TaskCreateTime *string `json:"TaskCreateTime,omitempty" name:"TaskCreateTime"`
		// 网络类型.PublicNetwork为公网类型,VPCNetwork为VPC网络

		NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
		// API Server地址

		ApiServerAddress *string `json:"ApiServerAddress,omitempty" name:"ApiServerAddress"`
		// 节点数

		NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`
		// 命名空间数

		NamespaceCount *uint64 `json:"NamespaceCount,omitempty" name:"NamespaceCount"`
		// 工作负载数

		WorkloadCount *uint64 `json:"WorkloadCount,omitempty" name:"WorkloadCount"`
		// Pod数量

		PodCount *uint64 `json:"PodCount,omitempty" name:"PodCount"`
		// Service数量

		ServiceCount *uint64 `json:"ServiceCount,omitempty" name:"ServiceCount"`
		// Ingress数量

		IngressCount *uint64 `json:"IngressCount,omitempty" name:"IngressCount"`
		// 主节点的ip列表

		MasterIps *string `json:"MasterIps,omitempty" name:"MasterIps"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EscapeEventInfo struct {
	// 事件类型
	//    ESCAPE_HOST_ACESS_FILE:宿主机文件访问逃逸
	//    ESCAPE_MOUNT_NAMESPACE:MountNamespace逃逸
	//    ESCAPE_PRIVILEDGE:程序提权逃逸
	//    ESCAPE_PRIVILEDGE_CONTAINER_START:特权容器启动逃逸
	//    ESCAPE_MOUNT_SENSITIVE_PTAH:敏感路径挂载
	//    ESCAPE_SYSCALL:Syscall逃逸

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 容器名

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 状态
	//      EVENT_UNDEAL:事件未处理
	//      EVENT_DEALED:事件已经处理
	//      EVENT_INGNORE：事件忽略

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件记录的唯一id

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// pod(实例)的名字

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 生成时间

	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
	// 事件名字，
	// 宿主机文件访问逃逸、
	// Syscall逃逸、
	// MountNamespace逃逸、
	// 程序提权逃逸、
	// 特权容器启动逃逸、
	// 敏感路径挂载

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 镜像id，用于跳转

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 容器id，用于跳转

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 事件解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// AppId值

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 事件数量

	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
	// 最近生成时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 节点IP

	NodeIP *string `json:"NodeIP,omitempty" name:"NodeIP"`
	// 主机IP

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// "NODE_DESTROYED"      //节点已销毁
	// "CONTAINER_EXITED"    //容器已退出
	// "CONTAINER_DESTROYED" //容器已销毁
	// "SHARED_HOST"         // 容器与主机共享网络
	// "RESOURCE_LIMIT"      //隔离操作资源超限
	// "UNKNOW"              // 原因未知

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离操作来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING

	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
	// 节点所属集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 节点类型：NORMAL普通节点、SUPER超级节点

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// pod ip

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// 节点唯一id

	NodeUniqueID *string `json:"NodeUniqueID,omitempty" name:"NodeUniqueID"`
	// 节点公网ip

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 节点id

	NodeID *string `json:"NodeID,omitempty" name:"NodeID"`
	// 节点内网ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type DescribeAssetImageRegistryVulListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像漏洞列表

		List []*ImageVul `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryVulListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousConnectionRuleSummaryRequest struct {
	*tchttp.BaseRequest

	// 当前appID

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeMaliciousConnectionRuleSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaliciousConnectionRuleSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoStartSecurityServiceStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoStartSecurityServiceStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAutoStartSecurityServiceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSearchLogsRequest struct {
	*tchttp.BaseRequest

	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeSearchLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSearchLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContainerNetwork struct {
	// endpoint id

	EndpointID *string `json:"EndpointID,omitempty" name:"EndpointID"`
	// 模式:bridge

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 网络名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 网络ID

	NetworkID *string `json:"NetworkID,omitempty" name:"NetworkID"`
	// 网关

	Gateway *string `json:"Gateway,omitempty" name:"Gateway"`
	// IPV4地址

	Ipv4 *string `json:"Ipv4,omitempty" name:"Ipv4"`
	// IPV6地址

	Ipv6 *string `json:"Ipv6,omitempty" name:"Ipv6"`
	// MAC 地址

	MAC *string `json:"MAC,omitempty" name:"MAC"`
}

type ReverseShellEventDescription struct {
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 解决方案

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 事件备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 目标地址

	DstAddress *string `json:"DstAddress,omitempty" name:"DstAddress"`
	// 事件最后一次处理的时间

	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
}

type DescribeK8sApiAbnormalTendencyRequest struct {
	*tchttp.BaseRequest

	// 趋势周期(默认为7天)

	TendencyPeriod *uint64 `json:"TendencyPeriod,omitempty" name:"TendencyPeriod"`
	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeK8sApiAbnormalTendencyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalTendencyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageComponentListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>ComponentName- String - 是否必填：否 - 镜像组件名称</li><li>ComponentVersion- String - 是否必填：否 - 镜像组件版本</li><li>ComponentType- String - 是否必填：否 - 镜像组件类型</li><li>VulLevel- String - 是否必填：否 - 漏洞威胁等级</li><li>HasVul- String - 是否必填：否 -是否有漏洞；true：是，false，否；不传或ALL ：全部</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式desc ，asc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeImageComponentListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageComponentListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersNamespacesRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤字段

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeClustersNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClustersNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAffectedNodeListRequest struct {
	*tchttp.BaseRequest

	// 唯一的检测项的ID

	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ClusterName, ClusterId,InstanceId,PrivateIpAddresses

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAffectedNodeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAffectedNodeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileAttributeInfo struct {
	// 文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件类型

	FileType *string `json:"FileType,omitempty" name:"FileType"`
	// 文件大小(字节)

	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`
	// 文件路径

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 文件创建时间

	FileCreateTime *string `json:"FileCreateTime,omitempty" name:"FileCreateTime"`
	// 最近被篡改文件创建时间

	LatestTamperedFileMTime *string `json:"LatestTamperedFileMTime,omitempty" name:"LatestTamperedFileMTime"`
	// 新文件内容

	NewFile *string `json:"NewFile,omitempty" name:"NewFile"`
	// 差异化内容

	FileDiff *string `json:"FileDiff,omitempty" name:"FileDiff"`
}

type VulAffectedContainerInfo struct {
	// 内网IP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 容器ID

	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// Pod名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// PodIP值

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机ID

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 外网IP

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 租户appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type CreateVulExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10000，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>ComponentName- String - 是否必填：否 - 镜像组件名称</li><li>ComponentVersion- String - 是否必填：否 - 镜像组件版本</li><li>ComponentType- String - 是否必填：否 - 镜像组件类型</li><li>VulLevel- String - 是否必填：否 - 漏洞威胁等级</li><li>HasVul- String - 是否必填：否 -是否有漏洞；true：是，false，否；不传或ALL ：全部</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式desc ，asc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateVulExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFlexBillingStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFlexBillingStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFlexBillingStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterRiskTarget struct {
	// 检查对象名称

	RiskTarget *string `json:"RiskTarget,omitempty" name:"RiskTarget"`
	// 严重的风险数量

	SeriousRiskCount *uint64 `json:"SeriousRiskCount,omitempty" name:"SeriousRiskCount"`
	// 高危的风险数量

	HighRiskCount *uint64 `json:"HighRiskCount,omitempty" name:"HighRiskCount"`
	// 中危的风险数量

	MiddleRiskCount *uint64 `json:"MiddleRiskCount,omitempty" name:"MiddleRiskCount"`
	// 提示的风险数量

	HintRiskCount *uint64 `json:"HintRiskCount,omitempty" name:"HintRiskCount"`
}

type EmergencyVulInfo struct {
	// 漏洞名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 漏洞标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// CVSS V3分数

	CVSSV3Score *float64 `json:"CVSSV3Score,omitempty" name:"CVSSV3Score"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// CVE编号

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// 漏洞类型

	Category *string `json:"Category,omitempty" name:"Category"`
	// 漏洞披露时间

	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`
	// 最近发现时间

	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
	// 应急漏洞风险情况：NOT_SCAN：未扫描，SCANNING：扫描中，SCANNED_NOT_RISK：已扫描，暂未风险 ，SCANNED_RISK：已扫描存在风险

	Status *string `json:"Status,omitempty" name:"Status"`
	// 漏洞ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 防御状态，NO_DEFENDED:未防御，DEFENDED:已防御

	DefenceStatus *string `json:"DefenceStatus,omitempty" name:"DefenceStatus"`
	// 漏洞防御主机范围: MANUAL:自选主机节点，ALL:全部

	DefenceScope *string `json:"DefenceScope,omitempty" name:"DefenceScope"`
	// 漏洞防御主机数量

	DefenceHostCount *int64 `json:"DefenceHostCount,omitempty" name:"DefenceHostCount"`
	// 已防御攻击次数

	DefendedCount *int64 `json:"DefendedCount,omitempty" name:"DefendedCount"`
	// 租户appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type ImageVirusInfo struct {
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 风险等级

	RiskLevel *uint64 `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 病毒名称

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
	// 标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 修复建议

	Solution *string `json:"Solution,omitempty" name:"Solution"`
	// 大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 首次发现时间

	FirstScanTime *string `json:"FirstScanTime,omitempty" name:"FirstScanTime"`
	// 最近扫描时间

	LatestScanTime *string `json:"LatestScanTime,omitempty" name:"LatestScanTime"`
	// 文件md5

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 检测平台
	// 1: 云查杀引擎
	// 2: tav
	// 3: binaryAi
	// 4: 异常行为
	// 5: 威胁情报

	CheckPlatform []*string `json:"CheckPlatform,omitempty" name:"CheckPlatform"`
}

type DescribeUserPurchaseInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 购买容器安全服务列表

		UserPurchaseInfoSet []*UserPurchaseInfo `json:"UserPurchaseInfoSet,omitempty" name:"UserPurchaseInfoSet"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserPurchaseInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserPurchaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVirusExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVirusExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVirusExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRiskSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全漏洞

		VulnerabilityCnt []*RunTimeRiskInfo `json:"VulnerabilityCnt,omitempty" name:"VulnerabilityCnt"`
		// 木马病毒

		MalwareVirusCnt []*RunTimeRiskInfo `json:"MalwareVirusCnt,omitempty" name:"MalwareVirusCnt"`
		// 敏感信息

		RiskCnt []*RunTimeRiskInfo `json:"RiskCnt,omitempty" name:"RiskCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageRiskSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRiskSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetProcessListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 端口列表

		List []*ProcessInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetProcessListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetProcessListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 异常进程策略信息列表

		RuleSet []*RuleBaseInfo `json:"RuleSet,omitempty" name:"RuleSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalProcessRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterIngressListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群Ingress列表

		ClusterIngressList []*ClusterIngressInfo `json:"ClusterIngressList,omitempty" name:"ClusterIngressList"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterIngressListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterIngressListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则列表

		List []*ImageDenyRule `json:"List,omitempty" name:"List"`
		// 规则总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageDenyRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlEventsExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// execle下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 任务ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessControlEventsExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlEventsExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostStartTheServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostStartTheServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostStartTheServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceAssetDetailInfo struct {
	// 客户资产的ID。

	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`
	// 资产类别。

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 资产的名称。

	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`
	// 资产所属的节点的名称。

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 资产所在的主机的名称。

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 资产所在的主机的IP。

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 检测状态
	// CHECK_INIT, 待检测
	// CHECK_RUNNING, 检测中
	// CHECK_FINISHED, 检测完成
	// CHECK_FAILED, 检测失败

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 此类资产通过的检测项的数目。

	PassedPolicyItemCount *uint64 `json:"PassedPolicyItemCount,omitempty" name:"PassedPolicyItemCount"`
	// 此类资产未通过的检测的数目。

	FailedPolicyItemCount *uint64 `json:"FailedPolicyItemCount,omitempty" name:"FailedPolicyItemCount"`
	// 上次检测的时间。

	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`
	// 检测结果：
	// RESULT_FAILED: 未通过。
	// RESULT_PASSED: 通过。

	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`
	// 资产的运行状态。
	// ASSET_NORMAL: 正常运行，
	// ASSET_PAUSED: 暂停运行，
	// ASSET_STOPPED: 停止运行，
	// ASSET_ABNORMAL: 异常

	AssetStatus *string `json:"AssetStatus,omitempty" name:"AssetStatus"`
	// 创建资产的时间。

	AssetCreateTime *string `json:"AssetCreateTime,omitempty" name:"AssetCreateTime"`
}

type DescribeRiskSyscallNamesRequest struct {
	*tchttp.BaseRequest

	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeRiskSyscallNamesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallNamesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallEventsRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数，"Filters":[{"Name":"Status"}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeRiskSyscallEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusAutoIsolateSampleDownloadURLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 样本下载链接

		FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusAutoIsolateSampleDownloadURLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSampleDownloadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulRegistryImageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 仓库镜像列表

		List []*VulAffectedRegistryImageInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulRegistryImageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulRegistryImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExportComplianceStatusListJobRequest struct {
	*tchttp.BaseRequest

	// 要导出信息的资产类型

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 按照检测项导出，还是按照资产导出。true: 按照资产导出；false: 按照检测项导出。

	ExportByAsset *bool `json:"ExportByAsset,omitempty" name:"ExportByAsset"`
	// 要导出的资产ID列表或检测项ID列表，由ExportByAsset的取值决定。

	IdList []*uint64 `json:"IdList,omitempty" name:"IdList"`
	// true, 全部导出；false, 根据IdList来导出数据。

	ExportAll *bool `json:"ExportAll,omitempty" name:"ExportAll"`
	// AppId

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *CreateExportComplianceStatusListJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExportComplianceStatusListJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeEventTendencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 待处理逃逸事件趋势

		List []*EscapeEventTendencyInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEscapeEventTendencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContainerMount struct {
	// 挂载类型 bind

	Type *string `json:"Type,omitempty" name:"Type"`
	// 宿主机路径

	Source *string `json:"Source,omitempty" name:"Source"`
	// 容器内路径

	Destination *string `json:"Destination,omitempty" name:"Destination"`
	// 模式

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 读写权限

	RW *bool `json:"RW,omitempty" name:"RW"`
	// 传播类型

	Propagation *string `json:"Propagation,omitempty" name:"Propagation"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 驱动

	Driver *string `json:"Driver,omitempty" name:"Driver"`
}

type PortInfo struct {
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 对外ip

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 主机端口

	PublicPort *uint64 `json:"PublicPort,omitempty" name:"PublicPort"`
	// 容器端口

	ContainerPort *uint64 `json:"ContainerPort,omitempty" name:"ContainerPort"`
	// 容器Pid

	ContainerPID *uint64 `json:"ContainerPID,omitempty" name:"ContainerPID"`
	// 容器名

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 主机id

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 主机ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 容器内监听地址

	ListenContainer *string `json:"ListenContainer,omitempty" name:"ListenContainer"`
	// 容器外监听地址

	ListenHost *string `json:"ListenHost,omitempty" name:"ListenHost"`
	// 运行账号

	RunAs *string `json:"RunAs,omitempty" name:"RunAs"`
	// 租户id，运营端使用

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type DescribeVirusDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像ID

		ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
		// 镜像名称

		ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
		// 创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 木马文件大小

		Size *uint64 `json:"Size,omitempty" name:"Size"`
		// 木马文件路径

		FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
		// 最近生成时间

		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
		// 病毒名称

		VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
		// 风险等级 RISK_CRITICAL, RISK_HIGH, RISK_MEDIUM, RISK_LOW, RISK_NOTICE。

		RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
		// 容器名称

		ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
		// 容器id

		ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
		// 主机名称

		HostName *string `json:"HostName,omitempty" name:"HostName"`
		// 主机id

		HostId *string `json:"HostId,omitempty" name:"HostId"`
		// 进程名称

		ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
		// 进程路径

		ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
		// 进程md5

		ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`
		// 进程id

		ProcessId *uint64 `json:"ProcessId,omitempty" name:"ProcessId"`
		// 进程参数

		ProcessArgv *string `json:"ProcessArgv,omitempty" name:"ProcessArgv"`
		// 进程链

		ProcessChan *string `json:"ProcessChan,omitempty" name:"ProcessChan"`
		// 进程组

		ProcessAccountGroup *string `json:"ProcessAccountGroup,omitempty" name:"ProcessAccountGroup"`
		// 进程启动用户

		ProcessStartAccount *string `json:"ProcessStartAccount,omitempty" name:"ProcessStartAccount"`
		// 进程文件权限

		ProcessFileAuthority *string `json:"ProcessFileAuthority,omitempty" name:"ProcessFileAuthority"`
		// 来源：0：一键扫描， 1：定时扫描 2：实时监控

		SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`
		// 集群名称

		PodName *string `json:"PodName,omitempty" name:"PodName"`
		// 标签

		Tags []*string `json:"Tags,omitempty" name:"Tags"`
		// 事件描述

		HarmDescribe *string `json:"HarmDescribe,omitempty" name:"HarmDescribe"`
		// 建议方案

		SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`
		// 备注

		Mark *string `json:"Mark,omitempty" name:"Mark"`
		// 风险文件名称

		FileName *string `json:"FileName,omitempty" name:"FileName"`
		// 文件MD5

		FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`
		// 事件类型

		EventType *string `json:"EventType,omitempty" name:"EventType"`
		// DEAL_NONE:文件待处理
		// DEAL_IGNORE:已经忽略
		// DEAL_ADD_WHITELIST:加白
		// DEAL_DEL:文件已经删除
		// DEAL_ISOLATE:已经隔离
		// DEAL_ISOLATING:隔离中
		// DEAL_ISOLATE_FAILED:隔离失败
		// DEAL_RECOVERING:恢复中
		// DEAL_RECOVER_FAILED: 恢复失败

		Status *string `json:"Status,omitempty" name:"Status"`
		// 失败子状态:
		// FILE_NOT_FOUND:文件不存在
		// FILE_ABNORMAL:文件异常
		// FILE_ABNORMAL_DEAL_RECOVER:恢复文件时，文件异常
		// BACKUP_FILE_NOT_FOUND:备份文件不存在
		// CONTAINER_NOT_FOUND_DEAL_ISOLATE:隔离时，容器不存在
		// CONTAINER_NOT_FOUND_DEAL_RECOVER:恢复时，容器不存在

		SubStatus *string `json:"SubStatus,omitempty" name:"SubStatus"`
		// 内网ip

		HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
		// 外网ip

		ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`
		// 父进程启动用户

		PProcessStartUser *string `json:"PProcessStartUser,omitempty" name:"PProcessStartUser"`
		// 父进程用户组

		PProcessUserGroup *string `json:"PProcessUserGroup,omitempty" name:"PProcessUserGroup"`
		// 父进程路径

		PProcessPath *string `json:"PProcessPath,omitempty" name:"PProcessPath"`
		// 父进程命令行参数

		PProcessParam *string `json:"PProcessParam,omitempty" name:"PProcessParam"`
		// 祖先进程启动用户

		AncestorProcessStartUser *string `json:"AncestorProcessStartUser,omitempty" name:"AncestorProcessStartUser"`
		// 祖先进程用户组

		AncestorProcessUserGroup *string `json:"AncestorProcessUserGroup,omitempty" name:"AncestorProcessUserGroup"`
		// 祖先进程路径

		AncestorProcessPath *string `json:"AncestorProcessPath,omitempty" name:"AncestorProcessPath"`
		// 祖先进程命令行参数

		AncestorProcessParam *string `json:"AncestorProcessParam,omitempty" name:"AncestorProcessParam"`
		// 事件最后一次处理的时间

		OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
		// 容器隔离状态

		ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
		// 容器隔离子状态

		ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
		// 容器隔离操作来源

		ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
		// 检测平台
		// 1: 云查杀引擎
		// 2: tav
		// 3: binaryAi
		// 4: 异常行为
		// 5: 威胁情报

		CheckPlatform []*string `json:"CheckPlatform,omitempty" name:"CheckPlatform"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRegistryNamespaceListRequest struct {
	*tchttp.BaseRequest

	// 本次查询的起始偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 本次查询的数据量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 查询的过滤条件。Name字段可取值"Namespace"。

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeImageRegistryNamespaceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRegistryNamespaceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulContainerListRequest struct {
	*tchttp.BaseRequest

	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID</li>
	// <li>ContainerName- String -是否必填: 否 - 容器名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVulContainerListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulContainerListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterPodInfo struct {
	// Pod名称.

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// Pod状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// Pod IP

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// 节点内网Ip

	NodeLanIP *string `json:"NodeLanIP,omitempty" name:"NodeLanIP"`
	// 所属的工作负载名字

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
	// 所属工作负载类型

	WorkloadKind *string `json:"WorkloadKind,omitempty" name:"WorkloadKind"`
	// 所属集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 所属集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 所属命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 所属地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 运行时间

	Age *string `json:"Age,omitempty" name:"Age"`
	// 创建时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 重启次数

	Restarts *uint64 `json:"Restarts,omitempty" name:"Restarts"`
	// 关联的service名字

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 关联的service数量

	ServiceCount *uint64 `json:"ServiceCount,omitempty" name:"ServiceCount"`
	// 关联的容器名字

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 关联的容器数量

	ContainerCount *uint64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
	// CPU占用率

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 内存占用量

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// Pod标签

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 集群状态

	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
	// 工作负载标签

	WorkloadLabels *string `json:"WorkloadLabels,omitempty" name:"WorkloadLabels"`
	// 容器Id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机Id

	HostId *string `json:"HostId,omitempty" name:"HostId"`
	// 集群类型

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
}

type ImageRiskInfo struct {
	// 行为

	Behavior *uint64 `json:"Behavior,omitempty" name:"Behavior"`
	// 类型

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 级别

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 详情

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 解决方案

	InstructionContent *string `json:"InstructionContent,omitempty" name:"InstructionContent"`
}

type ProcessDetailBaseInfo struct {
	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 进程pid

	ProcessId *uint64 `json:"ProcessId,omitempty" name:"ProcessId"`
	// 进程启动用户

	ProcessStartUser *string `json:"ProcessStartUser,omitempty" name:"ProcessStartUser"`
	// 进程用户组

	ProcessUserGroup *string `json:"ProcessUserGroup,omitempty" name:"ProcessUserGroup"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 进程命令行参数

	ProcessParam *string `json:"ProcessParam,omitempty" name:"ProcessParam"`
}

type DescribeAssetImageRegistryRiskInfoListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像id

	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`
	// 排序字段（Level）

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 + -

	Order *string `json:"Order,omitempty" name:"Order"`
	// 镜像标识Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryRiskInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryRiskInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRefreshTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 刷新任务状态，可能为：Task_New,Task_Running,Task_Finish,Task_Error,Task_NoExist

		TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRefreshTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRefreshTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterWorkloadInfo struct {
	// 工作负载名称

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
	// 工作负载类型

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 工作负载标签

	WorkloadLabels *string `json:"WorkloadLabels,omitempty" name:"WorkloadLabels"`
	// 关联的Pod数量

	PodCount *uint64 `json:"PodCount,omitempty" name:"PodCount"`
	// 工作负载Selector

	WorkloadSelectors *string `json:"WorkloadSelectors,omitempty" name:"WorkloadSelectors"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type CreateRiskSyscallWhiteListsExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数，"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateRiskSyscallWhiteListsExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRiskSyscallWhiteListsExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryVulListExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryVulListExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryVulListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserClusterRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeUserClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AffectedNodeItem struct {
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 内网ip地址

	PrivateIpAddresses *string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 节点的角色，Master、Work等

	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
	// k8s版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 运行时组件,docker或者containerd

	ContainerRuntime *string `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 检查结果的验证信息

	VerifyInfo *string `json:"VerifyInfo,omitempty" name:"VerifyInfo"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
}

type DescribeEscapeEventDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件基本信息

		EventBaseInfo *RunTimeEventBaseInfo `json:"EventBaseInfo,omitempty" name:"EventBaseInfo"`
		// 进程信息

		ProcessInfo *ProcessDetailInfo `json:"ProcessInfo,omitempty" name:"ProcessInfo"`
		// 事件描述

		EventDetail *EscapeEventDescription `json:"EventDetail,omitempty" name:"EventDetail"`
		// 父进程信息

		ParentProcessInfo *ProcessBaseInfo `json:"ParentProcessInfo,omitempty" name:"ParentProcessInfo"`
		// 祖先进程信息

		AncestorProcessInfo *ProcessBaseInfo `json:"AncestorProcessInfo,omitempty" name:"AncestorProcessInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEscapeEventDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetClusterListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>ClusterID - string  - 是否必填: 否 -集群ID</li>
	// <li>ClusterName - string  - 是否必填: 否 -集群名称</li>
	// <li>Status - string  - 是否必填: 否 -集群状态</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段。

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAssetClusterListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetClusterListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetContainerListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 容器列表

		List []*ContainerInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetContainerListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetContainerListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusScanTaskStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查杀容器个数

		ContainerTotal *uint64 `json:"ContainerTotal,omitempty" name:"ContainerTotal"`
		// 风险容器个数

		RiskContainerCnt *uint64 `json:"RiskContainerCnt,omitempty" name:"RiskContainerCnt"`
		// 扫描状态 任务状态:
		// SCAN_NONE:无，
		// SCAN_SCANNING:正在扫描中，
		// SCAN_FINISH：扫描完成，
		// SCAN_TIMEOUT：扫描超时
		// SCAN_CANCELING: 取消中
		// SCAN_CANCELED:已取消

		Status *string `json:"Status,omitempty" name:"Status"`
		// 扫描进度 I

		Schedule *uint64 `json:"Schedule,omitempty" name:"Schedule"`
		// 已经扫描了的容器个数

		ContainerScanCnt *uint64 `json:"ContainerScanCnt,omitempty" name:"ContainerScanCnt"`
		// 风险个数

		RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`
		// 剩余扫描时间

		LeftSeconds *uint64 `json:"LeftSeconds,omitempty" name:"LeftSeconds"`
		// 扫描开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 扫描结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 扫描类型，"CYCLE"：周期扫描， "MANUAL"：手动扫描

		ScanType *string `json:"ScanType,omitempty" name:"ScanType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusScanTaskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusScanTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogStorageStatisticResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总容量（单位：B）

		TotalSize *uint64 `json:"TotalSize,omitempty" name:"TotalSize"`
		// 已使用容量（单位：B）

		UsedSize *uint64 `json:"UsedSize,omitempty" name:"UsedSize"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogStorageStatisticResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogStorageStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySystemRule struct {
	// 事件类型

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 状态；true 开；false 关

	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
	// 策略优先级;true系统策略优先，否则用户设置优先；逃逸需要字段；

	Priority *bool `json:"Priority,omitempty" name:"Priority"`
}

type DescribeRiskSyscallDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件基本信息

		EventBaseInfo *RunTimeEventBaseInfo `json:"EventBaseInfo,omitempty" name:"EventBaseInfo"`
		// 进程信息

		ProcessInfo *ProcessDetailInfo `json:"ProcessInfo,omitempty" name:"ProcessInfo"`
		// 父进程信息

		ParentProcessInfo *ProcessDetailBaseInfo `json:"ParentProcessInfo,omitempty" name:"ParentProcessInfo"`
		// 事件描述

		EventDetail *RiskSyscallEventDescription `json:"EventDetail,omitempty" name:"EventDetail"`
		// 祖先进程信息

		AncestorProcessInfo *ProcessBaseInfo `json:"AncestorProcessInfo,omitempty" name:"AncestorProcessInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskSyscallDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousConnectionWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 恶意请求白名单总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 恶意请求白名单列表

		List []*MaliciousConnectionRuleInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMaliciousConnectionWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaliciousConnectionWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventEscapeImageDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一值

		UniqueKey *string `json:"UniqueKey,omitempty" name:"UniqueKey"`
		// 事件类型 MOUNT_SENSITIVE_PTAH : 敏感路径挂载
		// PRIVILEGE_CONTAINER_START:特权容器

		EventType *string `json:"EventType,omitempty" name:"EventType"`
		// 备注

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 操作时间

		OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
		// 镜像ID

		ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
		// 镜像名称

		ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
		// 影响容器数量

		ContainerCount *uint64 `json:"ContainerCount,omitempty" name:"ContainerCount"`
		// 告警数量

		EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
		// 首次生成时间

		FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`
		// 最近生成时间

		LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`
		// 状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// 解决方案

		Solution *string `json:"Solution,omitempty" name:"Solution"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEventEscapeImageDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventEscapeImageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterNodeWhiteListRequest struct {
	*tchttp.BaseRequest

	// 检测项ID

	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：InstanceId实例id,InstanceRole节点的角色，Master、Work等

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeClusterNodeWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterNodeWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetContainerListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// 	<li>AppId - String - 是否必填：否 - 租户Id匹配</li>
	// 	<li>ContainerName - String - 是否必填：否 - 容器名称模糊搜索</li>
	// 	<li>ContainerId - String - 是否必填：否 - 容器id</li>
	// 	<li>ImageId - String - 是否必填：否 - 镜像Id</li>
	// 	<li>ImageName- String - 是否必填：否 - 镜像名称搜索</li>
	// 	<li>HostName- String - 是否必填：否 - 节点名称</li>
	// 	<li>Status - String - 是否必填：否 - 容器运行状态筛选，0："created",1："running", 2："paused", 3："restarting", 4："removing", 5："exited", 6："dead" </li>
	// 	<li>ContainerStatus - String - 是否必填：否 - 容器运行状态</li>
	// 	<li>ContainerNetStatus - String - 是否必填：否 - 容器隔离状态</li>
	// 	<li>HostIP- string - 是否必填：否 - 节点内网IP</li>
	// 	<li>PublicIp- string - 是否必填：否 - 节点外网IP</li>
	// 	<li>OrderBy - String 是否必填：否 -排序字段，支持：cpu_usage, mem_usage的动态排序 ["cpu_usage","+"]  '+'升序、'-'降序</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetContainerListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetContainerListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterIngressListRequest struct {
	*tchttp.BaseRequest

	// 集群Id，不填表示查询用户所有ingress

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：Name, Address等

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeClusterIngressListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterIngressListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessRuleDetailRequest struct {
	*tchttp.BaseRequest

	// 策略唯一id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 镜像id, 在添加白名单的时候使用

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 租户端id

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAbnormalProcessRuleDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessRuleDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSupportRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域列表

		RegionList []*ClusterRegionItem `json:"RegionList,omitempty" name:"RegionList"`
		// 数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSupportRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSupportRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyEventTendencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像拦截成功事件趋势

		DenyList []*ImageDenyEventTendency `json:"DenyList,omitempty" name:"DenyList"`
		// 镜像拦截告警事件趋势

		AlarmList []*ImageDenyEventTendency `json:"AlarmList,omitempty" name:"AlarmList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageDenyEventTendencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyEventTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetHostDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云镜uuid

		UUID *string `json:"UUID,omitempty" name:"UUID"`
		// 更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 主机名

		HostName *string `json:"HostName,omitempty" name:"HostName"`
		// 主机分组

		Group *string `json:"Group,omitempty" name:"Group"`
		// 主机IP

		HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
		// 操作系统

		OsName *string `json:"OsName,omitempty" name:"OsName"`
		// agent版本

		AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
		// 内核版本

		KernelVersion *string `json:"KernelVersion,omitempty" name:"KernelVersion"`
		// docker版本

		DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`
		// docker api版本

		DockerAPIVersion *string `json:"DockerAPIVersion,omitempty" name:"DockerAPIVersion"`
		// docker go 版本

		DockerGoVersion *string `json:"DockerGoVersion,omitempty" name:"DockerGoVersion"`
		// docker 文件系统类型

		DockerFileSystemDriver *string `json:"DockerFileSystemDriver,omitempty" name:"DockerFileSystemDriver"`
		// docker root 目录

		DockerRootDir *string `json:"DockerRootDir,omitempty" name:"DockerRootDir"`
		// 镜像数

		ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`
		// 容器数

		ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
		// k8s ip

		K8sMasterIP *string `json:"K8sMasterIP,omitempty" name:"K8sMasterIP"`
		// k8s version

		K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`
		// kube proxy

		KubeProxyVersion *string `json:"KubeProxyVersion,omitempty" name:"KubeProxyVersion"`
		// 主机运行状态 offline,online,pause

		Status *string `json:"Status,omitempty" name:"Status"`
		// 外网ip

		PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
		// 主机Quuid

		Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
		// 是否Containerd

		IsContainerd *bool `json:"IsContainerd,omitempty" name:"IsContainerd"`
		// 主机来源;"TENCENTCLOUD":"腾讯云服务器","OTHERCLOUD":"非腾讯云服务器"

		MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
		// 主机实例ID

		InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
		// 地域ID

		RegionID *int64 `json:"RegionID,omitempty" name:"RegionID"`
		// 所属项目

		Project *ProjectInfo `json:"Project,omitempty" name:"Project"`
		// 标签

		Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`
		// 集群ID

		ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
		// 集群名称

		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
		// 集群接入状态

		ClusterAccessedStatus *string `json:"ClusterAccessedStatus,omitempty" name:"ClusterAccessedStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetHostDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetHostDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CompliancePeriodTaskRule struct {
	// 执行的频率（几天一次），取值为：1,3,7。

	Frequency *uint64 `json:"Frequency,omitempty" name:"Frequency"`
	// 在这天的什么时间执行，格式为：HH:mm:SS。

	ExecutionTime *string `json:"ExecutionTime,omitempty" name:"ExecutionTime"`
	// 是否开启

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

type CreateAssetImageVirusExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssetImageVirusExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetImageVirusExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessControlSystemChildRuleInfo struct {
	// 子策略Id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 策略模式,  RULE_MODE_RELEASE: 放行
	//    RULE_MODE_ALERT: 告警
	//    RULE_MODE_HOLDUP:拦截

	RuleMode *string `json:"RuleMode,omitempty" name:"RuleMode"`
	// 子策略状态，true为开启，false为关闭

	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
	// 子策略检测的入侵行为类型
	// CHANGE_CRONTAB：篡改计划任务
	// CHANGE_SYS_BIN：篡改系统程序
	// CHANGE_USRCFG：篡改用户配置

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
}

type CreateK8sApiAbnormalEventExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateK8sApiAbnormalEventExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateK8sApiAbnormalEventExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 恶意请求事件列表

		List []*RiskDnsEventInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskDnsListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskDnsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryListExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤字段

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式，asc，desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 是否仅展示repository版本最新的镜像，默认为false

	OnlyShowLatest *bool `json:"OnlyShowLatest,omitempty" name:"OnlyShowLatest"`
}

func (r *DescribeAssetImageRegistryListExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryListExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterNodesRequest struct {
	*tchttp.BaseRequest

	// 集群Id,不输入表示查询所有

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：InstanceState节点状态, InstanceRole节点类型等

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeClusterNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeRuleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则信息

		RuleSet []*EscapeRule `json:"RuleSet,omitempty" name:"RuleSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEscapeRuleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallWhiteListsRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数，"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeRiskSyscallWhiteListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallWhiteListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetHostDetailRequest struct {
	*tchttp.BaseRequest

	// 主机id

	HostId *string `json:"HostId,omitempty" name:"HostId"`
	// 租户id，当前查询的租户Id

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeAssetHostDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetHostDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRiskTendencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本地镜像新增风险趋势信息列表

		ImageRiskTendencySet []*ImageRiskTendencyInfo `json:"ImageRiskTendencySet,omitempty" name:"ImageRiskTendencySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageRiskTendencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRiskTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalTendencyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 趋势列表

		List []*K8sApiAbnormalTendencyItem `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeK8sApiAbnormalTendencyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateReverseShellWhiteListsExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateReverseShellWhiteListsExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateReverseShellWhiteListsExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousConnectionBlackListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>RequestType- string - 是否必填：否 - 请求类型，全部请求类型：ALL；域名：DOMAIN；IP: IP</li>
	// <li>BlackDomain- string - 是否必填：否 - 自定义黑域名</li>
	// <li>BlackIP- string - 是否必填：否 - 自定义黑IP</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeMaliciousConnectionBlackListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMaliciousConnectionBlackListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageListExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>AppId - String - 是否必填：否 - 租户id筛选</li>
	// <li>ImageName- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>ScanStatus - String - 是否必填：否 - 镜像扫描状态notScan，scanning，scanned，scanErr</li>
	// <li>ImageID- String - 是否必填：否 - 镜像ID筛选，</li>
	// <li>SecurityRisk- String - 是否必填：否 - 安全风险，VulCnt 、VirusCnt、RiskCnt、IsTrustImage</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeAssetImageListExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageListExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type K8sApiAbnormalEventInfo struct {
	// 命中规则名称

	MatchRuleName *string `json:"MatchRuleName,omitempty" name:"MatchRuleName"`
	// 命中规则类型

	MatchRuleType *string `json:"MatchRuleType,omitempty" name:"MatchRuleType"`
	// 告警等级

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群运行状态

	ClusterRunningStatus *string `json:"ClusterRunningStatus,omitempty" name:"ClusterRunningStatus"`
	// 初次生成时间

	FirstCreateTime *string `json:"FirstCreateTime,omitempty" name:"FirstCreateTime"`
	// 最近一次生成时间

	LastCreateTime *string `json:"LastCreateTime,omitempty" name:"LastCreateTime"`
	// 告警数量

	AlarmCount *uint64 `json:"AlarmCount,omitempty" name:"AlarmCount"`
	// 状态
	// "EVENT_UNDEAL":未处理
	// "EVENT_DEALED": 已处理
	// "EVENT_IGNORE": 忽略
	// "EVENT_DEL": 删除
	// "EVENT_ADD_WHITE": 加白

	Status *string `json:"Status,omitempty" name:"Status"`
	// 集群masterIP

	ClusterMasterIP *string `json:"ClusterMasterIP,omitempty" name:"ClusterMasterIP"`
	// k8s版本

	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`
	// 运行时组件

	RunningComponent []*string `json:"RunningComponent,omitempty" name:"RunningComponent"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 建议

	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`
	// 请求信息

	Info *string `json:"Info,omitempty" name:"Info"`
	// 规则ID

	MatchRuleID *string `json:"MatchRuleID,omitempty" name:"MatchRuleID"`
	// 高亮字段数组

	HighLightFields []*string `json:"HighLightFields,omitempty" name:"HighLightFields"`
	// 命中规则

	MatchRule *K8sApiAbnormalRuleScopeInfo `json:"MatchRule,omitempty" name:"MatchRule"`
}

type RunTimeFilters struct {
	// 过滤键的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 是否模糊查询

	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type DescribeAssetImageRiskListExportRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 风险级别 1,2,3,4，</li>
	// <li>Behavior - String - 是否必填：否 - 风险行为 1,2,3,4</li>
	// <li>Type - String - 是否必填：否 - 风险类型  1,2,</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
}

func (r *DescribeAssetImageRiskListExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRiskListExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 访问控制事件数组

		EventSet []*AccessControlEventInfo `json:"EventSet,omitempty" name:"EventSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessControlEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则列表

		List []*K8sApiAbnormalRuleListItem `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeK8sApiAbnormalRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceTaskAssetSummaryRequest struct {
	*tchttp.BaseRequest

	// 资产类型列表。
	// ASSET_CONTAINER, 容器
	// ASSET_IMAGE, 镜像
	// ASSET_HOST, 主机
	// ASSET_K8S, K8S资产

	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`
	// 筛选条件 AppId

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComplianceTaskAssetSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceTaskAssetSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalEventListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件列表

		List []*K8sApiAbnormalEventListItem `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeK8sApiAbnormalEventListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalEventInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件详情

		Info *K8sApiAbnormalEventInfo `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeK8sApiAbnormalEventInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalEventInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterNodeInfo struct {
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 内网ip地址

	PrivateIpAddresses *string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 节点的角色，Master、Work等

	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
	// 实例的状态（running 运行中，initializing 初始化中，failed 异常）

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// agent安装状态

	AgentStatus *string `json:"AgentStatus,omitempty" name:"AgentStatus"`
	// 公网ip

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 节点ID

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 主机类型(普通节点情况)

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 节点类型( NORMAL: 普通节点 SUPER:超级节点 )

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// uuid

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 计费核数

	ChargeCoresCnt *uint64 `json:"ChargeCoresCnt,omitempty" name:"ChargeCoresCnt"`
	// 防护状态: 已防护: Defended 未防护: UnDefended

	DefendStatus *string `json:"DefendStatus,omitempty" name:"DefendStatus"`
}

type DescribeRiskSyscallEventsExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeRiskSyscallEventsExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallEventsExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IngressForwardConfig struct {
	// 协议.

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 监听端口

	ListeningPort *uint64 `json:"ListeningPort,omitempty" name:"ListeningPort"`
	// 域名

	Host *string `json:"Host,omitempty" name:"Host"`
	// URL路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 后端服务

	BackendServiceName *string `json:"BackendServiceName,omitempty" name:"BackendServiceName"`
	// 服务端口

	ServicePort *uint64 `json:"ServicePort,omitempty" name:"ServicePort"`
	// 重定向目标地址

	RewriteAddress *string `json:"RewriteAddress,omitempty" name:"RewriteAddress"`
	// 1 ForwardConfig:转发配置,2 RewriteForwardConfig:重定向转发配置

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type DescribeVirusDetailRequest struct {
	*tchttp.BaseRequest

	// 木马文件id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeVirusDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulImageListRequest struct {
	*tchttp.BaseRequest

	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>AppId - String - 是否必填：是 - 租户id筛选</li>
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>HostIP- string -是否必填: 否 - 内网IP</li>
	// <li>PublicIP- string -是否必填: 否 - 外网IP</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>HostName- string -是否必填: 否 - 主机名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeVulImageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulImageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRiskListExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRiskListExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRiskListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageVulListExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageVulListExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageVulListExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterWorkloadWhiteListRequest struct {
	*tchttp.BaseRequest

	// 检测项ID

	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name - String
	// Name 可取值：ClusterId集群Id,ClusterName集群名字,WorkloadName工作负载名称

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeClusterWorkloadWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterWorkloadWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateReverseShellWhiteListsExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateReverseShellWhiteListsExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateReverseShellWhiteListsExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAssetImageRiskExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssetImageRiskExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAssetImageRiskExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulRegistryImageExportJobRequest struct {
	*tchttp.BaseRequest

	// 漏洞ID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 需要返回的数量，默认为50000，最大值为50000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// OnlyAffectedNewestImage bool 是否影响最新镜像
	// ImageDigest 镜像Digest，支持模糊查询
	// ImageId 镜像ID，支持模糊查询
	// Namespace 命名空间，支持模糊查询
	// ImageTag 镜像版本，支持模糊查询
	// InstanceName 实例名称，支持模糊查询
	// ImageName 镜像名，支持模糊查询
	// ImageRepoAddress 镜像地址，支持模糊查询

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateVulRegistryImageExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulRegistryImageExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventEscapeImageDetailRequest struct {
	*tchttp.BaseRequest

	// 唯一值

	UniqueKey *string `json:"UniqueKey,omitempty" name:"UniqueKey"`
}

func (r *DescribeEventEscapeImageDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventEscapeImageDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sApiAbnormalSummaryRequest struct {
	*tchttp.BaseRequest

	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeK8sApiAbnormalSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sApiAbnormalSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageDenyEventExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageDenyEventExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageDenyEventExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEmergencyVulListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>AppId - String - 是否必填：是 - 租户id筛选</li><li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li><li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li><li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li><li>CVEID- string - 是否必填：否 - CVE编号</li><li>ImageID- string - 是否必填：否 - 镜像ID</li><li>ImageName- String -是否必填: 否 - 镜像名称</li><li>ContainerID- string -是否必填: 否 - 容器ID</li><li>ContainerName- string -是否必填: 否 - 容器名称</li><li>ComponentName- string -是否必填: 否 - 组件名称</li><li>ComponentVersion- string -是否必填: 否 - 组件版本</li><li>Name- string -是否必填: 否 - 漏洞名称</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeEmergencyVulListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEmergencyVulListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessLevelSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异常进程高危待处理事件数

		HighLevelEventCount *int64 `json:"HighLevelEventCount,omitempty" name:"HighLevelEventCount"`
		// 异常进程中危待处理事件数

		MediumLevelEventCount *int64 `json:"MediumLevelEventCount,omitempty" name:"MediumLevelEventCount"`
		// 异常进程低危待处理事件数

		LowLevelEventCount *int64 `json:"LowLevelEventCount,omitempty" name:"LowLevelEventCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalProcessLevelSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessLevelSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageDenyRuleExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>RuleType- String - 是否必填：否 -规则类型 RULE_RISK：风险， RULE_PRIVILEGE：特权</li>
	// <li>EffectStatus- String - 是否必填：否 - 生效状态 IN_THE_TEST ：观察中，IN_EFFECT：生效中。</li>
	// <li>RuleName- string - 是否必填：否 - 规则名称。</li>
	// <li>Status- string - 是否必填：否 - 开启状态 0：开启，1：关闭。</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10000，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式：asc/desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：生效时间：EffectTime，更新时间：UpdateTime，Status：启动状态

	By *string `json:"By,omitempty" name:"By"`
	// 置顶已开启规则 true：是 ，否：false

	TopTurnOn *bool `json:"TopTurnOn,omitempty" name:"TopTurnOn"`
}

func (r *CreateImageDenyRuleExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageDenyRuleExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusMonitorSettingRequest struct {
	*tchttp.BaseRequest

	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeVirusMonitorSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusMonitorSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetContainerDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机id

		HostID *string `json:"HostID,omitempty" name:"HostID"`
		// 主机ip

		HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
		// 容器名称

		ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
		// 运行状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 运行账户

		RunAs *string `json:"RunAs,omitempty" name:"RunAs"`
		// 命令行

		Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
		// CPU使用率 * 1000

		CPUUsage *uint64 `json:"CPUUsage,omitempty" name:"CPUUsage"`
		// 内存使用 KB

		RamUsage *uint64 `json:"RamUsage,omitempty" name:"RamUsage"`
		// 镜像名

		ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
		// 镜像ID

		ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
		// 归属POD

		POD *string `json:"POD,omitempty" name:"POD"`
		// k8s 主节点

		K8sMaster *string `json:"K8sMaster,omitempty" name:"K8sMaster"`
		// 容器内进程数

		ProcessCnt *uint64 `json:"ProcessCnt,omitempty" name:"ProcessCnt"`
		// 容器内端口数

		PortCnt *uint64 `json:"PortCnt,omitempty" name:"PortCnt"`
		// 组件数

		ComponentCnt *uint64 `json:"ComponentCnt,omitempty" name:"ComponentCnt"`
		// app数

		AppCnt *uint64 `json:"AppCnt,omitempty" name:"AppCnt"`
		// websvc数

		WebServiceCnt *uint64 `json:"WebServiceCnt,omitempty" name:"WebServiceCnt"`
		// 挂载

		Mounts []*ContainerMount `json:"Mounts,omitempty" name:"Mounts"`
		// 容器网络信息

		Network *ContainerNetwork `json:"Network,omitempty" name:"Network"`
		// 创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 镜像创建时间

		ImageCreateTime *string `json:"ImageCreateTime,omitempty" name:"ImageCreateTime"`
		// 镜像大小

		ImageSize *uint64 `json:"ImageSize,omitempty" name:"ImageSize"`
		// 主机状态 offline,online,pause

		HostStatus *string `json:"HostStatus,omitempty" name:"HostStatus"`
		// 网络子状态

		NetSubStatus *string `json:"NetSubStatus,omitempty" name:"NetSubStatus"`
		// 隔离来源

		IsolateSource *string `json:"IsolateSource,omitempty" name:"IsolateSource"`
		// 隔离时间

		IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`
		// 网络状态 未隔离 NORMAL 已隔离 ISOLATED 隔离中 ISOLATING 隔离失败 ISOLATE_FAILED 解除隔离中 RESTORING 解除隔离失败 RESTORE_FAILED

		NetStatus *string `json:"NetStatus,omitempty" name:"NetStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetContainerDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetContainerDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemVulListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 漏洞列表

		List []*VulInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSystemVulListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProcessEventsExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,Status：EVENT_UNDEAL:未处理，EVENT_DEALED:已处理，EVENT_INGNORE:忽略

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段：latest_found_time

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateProcessEventsExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProcessEventsExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRiskSummaryRequest struct {
	*tchttp.BaseRequest

	// 当前AppId

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeImageRiskSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRiskSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AbnormalProcessChildRuleInfo struct {
	// 子策略id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 策略模式，   RULE_MODE_RELEASE: 放行
	//    RULE_MODE_ALERT: 告警
	//    RULE_MODE_HOLDUP:拦截

	RuleMode *string `json:"RuleMode,omitempty" name:"RuleMode"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 威胁等级，HIGH:高，MIDDLE:中，LOW:低

	RuleLevel *string `json:"RuleLevel,omitempty" name:"RuleLevel"`
}

type AssetFilters struct {
	// 过滤键的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 是否模糊查询

	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type DescribeAbnormalProcessDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件基本信息

		EventBaseInfo *RunTimeEventBaseInfo `json:"EventBaseInfo,omitempty" name:"EventBaseInfo"`
		// 进程信息

		ProcessInfo *ProcessDetailInfo `json:"ProcessInfo,omitempty" name:"ProcessInfo"`
		// 父进程信息

		ParentProcessInfo *ProcessDetailBaseInfo `json:"ParentProcessInfo,omitempty" name:"ParentProcessInfo"`
		// 事件描述

		EventDetail *AbnormalProcessEventDescription `json:"EventDetail,omitempty" name:"EventDetail"`
		// 祖先进程信息

		AncestorProcessInfo *ProcessBaseInfo `json:"AncestorProcessInfo,omitempty" name:"AncestorProcessInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalProcessDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulIgnoreLocalImageListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 排序方式:DESC,ACS

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段 ImageSize

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件:AppId 租户id

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVulIgnoreLocalImageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulIgnoreLocalImageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusMonitorSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否开启实时监控

		EnableScan *bool `json:"EnableScan,omitempty" name:"EnableScan"`
		// 扫描全部路径

		ScanPathAll *bool `json:"ScanPathAll,omitempty" name:"ScanPathAll"`
		// 当ScanPathAll为true 生效 0扫描以下路径 1、扫描除以下路径

		ScanPathType *uint64 `json:"ScanPathType,omitempty" name:"ScanPathType"`
		// 自选排除或扫描的地址

		ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirusMonitorSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusMonitorSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRiskDnsEventExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRiskDnsEventExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRiskDnsEventExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeESAggregationsRequest struct {
	*tchttp.BaseRequest

	// ES聚合条件JSON

	Query *string `json:"Query,omitempty" name:"Query"`
	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeESAggregationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeESAggregationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageVul struct {
	// 漏洞id

	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
	// 观点验证程序id

	POCID *string `json:"POCID,omitempty" name:"POCID"`
	// 漏洞名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 涉及组件信息

	Components []*ComponentsInfo `json:"Components,omitempty" name:"Components"`
	// 分类

	Category *string `json:"Category,omitempty" name:"Category"`
	// 分类2

	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`
	// 风险等级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 描述

	Des *string `json:"Des,omitempty" name:"Des"`
	// 解决方案

	OfficialSolution *string `json:"OfficialSolution,omitempty" name:"OfficialSolution"`
	// 引用

	Reference *string `json:"Reference,omitempty" name:"Reference"`
	// 防御方案

	DefenseSolution *string `json:"DefenseSolution,omitempty" name:"DefenseSolution"`
	// 提交时间

	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`
	// Cvss分数

	CvssScore *string `json:"CvssScore,omitempty" name:"CvssScore"`
	// Cvss信息

	CvssVector *string `json:"CvssVector,omitempty" name:"CvssVector"`
	// 是否建议修复

	IsSuggest *string `json:"IsSuggest,omitempty" name:"IsSuggest"`
	// 修复版本号

	FixedVersions *string `json:"FixedVersions,omitempty" name:"FixedVersions"`
	// 漏洞标签:"CanBeFixed","DynamicLevelPoc","DynamicLevelExp"

	Tag []*string `json:"Tag,omitempty" name:"Tag"`
	// 组件名

	Component *string `json:"Component,omitempty" name:"Component"`
	// 组件版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type RuntimeSystemRuleInfo struct {
	// 执行动作

	RuleMode *string `json:"RuleMode,omitempty" name:"RuleMode"`
	// 事件类型

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 状态

	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
	// 策略优先级;true系统策略优先，否则用户设置优先；逃逸需要字段；

	Priority *bool `json:"Priority,omitempty" name:"Priority"`
}

type DescribeRiskTargetCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 检查对象风险数量

		ClusterRiskTargets []*ClusterRiskTarget `json:"ClusterRiskTargets,omitempty" name:"ClusterRiskTargets"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskTargetCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskTargetCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryListExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryListExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageDenyRuleSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像拦截规则总数(含关闭的和开启的)

		RuleTotalCount *uint64 `json:"RuleTotalCount,omitempty" name:"RuleTotalCount"`
		// 开启的镜像拦截规则数

		EnabledRuleCount *uint64 `json:"EnabledRuleCount,omitempty" name:"EnabledRuleCount"`
		// 观察期中的镜像拦截规则数

		ObservedRuleCount *uint64 `json:"ObservedRuleCount,omitempty" name:"ObservedRuleCount"`
		// 已生效的镜像拦截规则数

		EffectiveRuleCount *uint64 `json:"EffectiveRuleCount,omitempty" name:"EffectiveRuleCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageDenyRuleSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageDenyRuleSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryVulListExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像信息

	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`
	// 排序字段（Level）

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 + -

	Order *string `json:"Order,omitempty" name:"Order"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 镜像标识Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryVulListExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryVulListExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallDetailRequest struct {
	*tchttp.BaseRequest

	// 事件唯一ID

	EventId *string `json:"EventId,omitempty" name:"EventId"`
	// 租户ID

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeRiskSyscallDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRiskSyscallDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterServiceYamlRequest struct {
	*tchttp.BaseRequest

	// Service名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// Service的命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群地域

	ClusterRegion *string `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
}

func (r *DescribeClusterServiceYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterServiceYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebVulListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>AppId - String - 是否必填：是 - 租户id筛选</li>
	// <li>OnlyAffectedContainer- string - 是否必填：否 - 仅展示影响容器的漏洞true,false</li>
	// <li>OnlyAffectedNewestImage-string - 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>AppId - String - 是否必填：是 - 租户id筛选</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	// <li>FocusOnType - string - 是否必填：否 -关注紧急度类型 。ALL :全部，SERIOUS_LEVEL： 严重和高危 ，IS_SUGGEST： 重点关注，POC_EXP 有Poc或Exp ，NETWORK_EXP: 远程Exp</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeWebVulListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebVulListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFlexBillingStateRequest struct {
	*tchttp.BaseRequest

	// appid数组

	AppIDs []*uint64 `json:"AppIDs,omitempty" name:"AppIDs"`
	// 状态(true:开 false:关)

	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyFlexBillingStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFlexBillingStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRegistryNamespaceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可返回的项目空间的总量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回的项目空间列表

		NamespaceList []*string `json:"NamespaceList,omitempty" name:"NamespaceList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageRegistryNamespaceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRegistryNamespaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserPodListRequest struct {
	*tchttp.BaseRequest

	// 集群Id,不填表示获取用户所有pod

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name 可取值：ClusterId集群id,Namespace命名空间等

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// Service名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeUserPodListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserPodListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageDenyRuleExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageDenyRuleExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageDenyRuleExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusTaskListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>ContainerName - String - 是否必填：否 - 容器名称</li>
	// <li>ContainerId - String - 是否必填：否 - 容器id</li>
	// <li>Hostname - String - 是否必填：否 - 主机名称</li>
	// <li>HostIp- String - 是否必填：否 - 主机IP</li>
	// <li>ImageId- String - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String - 是否必填：否 - 镜像名称</li>
	// <li>Status- String - 是否必填：否 - 状态</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeVirusTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEscapeRuleInfoRequest struct {
	*tchttp.BaseRequest

	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeEscapeRuleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEscapeRuleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSearchLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 历史搜索记录 保留最新的10条

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSearchLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSearchLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEventEscapeImageExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEventEscapeImageExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEventEscapeImageExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type K8sApiAbnormalRuleListItem struct {
	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 规则类型

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 受影响集群总数

	EffectClusterCount *uint64 `json:"EffectClusterCount,omitempty" name:"EffectClusterCount"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 编辑账号

	OprUin *string `json:"OprUin,omitempty" name:"OprUin"`
	// 状态

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 当前appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type ServiceInfo struct {
	// 服务id

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
	// 主机id

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 主机ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 容器名

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 服务名 例如nginx/redis

	Type *string `json:"Type,omitempty" name:"Type"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 账号

	RunAs *string `json:"RunAs,omitempty" name:"RunAs"`
	// 监听端口

	Listen []*string `json:"Listen,omitempty" name:"Listen"`
	// 配置

	Config *string `json:"Config,omitempty" name:"Config"`
	// 关联进程数

	ProcessCnt *uint64 `json:"ProcessCnt,omitempty" name:"ProcessCnt"`
	// 访问日志

	AccessLog *string `json:"AccessLog,omitempty" name:"AccessLog"`
	// 错误日志

	ErrorLog *string `json:"ErrorLog,omitempty" name:"ErrorLog"`
	// 数据目录

	DataPath *string `json:"DataPath,omitempty" name:"DataPath"`
	// web目录

	WebRoot *string `json:"WebRoot,omitempty" name:"WebRoot"`
	// 关联的进程id

	Pids []*uint64 `json:"Pids,omitempty" name:"Pids"`
	// 服务类型 app,web,db

	MainType *string `json:"MainType,omitempty" name:"MainType"`
	// 执行文件

	Exe *string `json:"Exe,omitempty" name:"Exe"`
	// 服务命令行参数

	Parameter *string `json:"Parameter,omitempty" name:"Parameter"`
	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 租户id，运营端专用

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 外网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
}

type DescribeVulSummaryRequest struct {
	*tchttp.BaseRequest

	// <li>CategoryType- string - 是否必填：否 - 漏洞分类: SYSTEM:系统漏洞 WEB:web应用漏洞 ALL:全部漏洞</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVulSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像仓库列表

		List []*ImageRepoInfo `json:"List,omitempty" name:"List"`
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetImageRegistryListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostInfo struct {
	// 主机id

	HostID *string `json:"HostID,omitempty" name:"HostID"`
	// 主机ip

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 业务组

	Group *string `json:"Group,omitempty" name:"Group"`
	// docker 版本

	DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`
	// docker 文件系统类型

	DockerFileSystemDriver *string `json:"DockerFileSystemDriver,omitempty" name:"DockerFileSystemDriver"`
	// 镜像个数

	ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`
	// 容器个数

	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
	// 主机运行状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 租户id,运营端使用

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 是否是Containerd

	IsContainerd *bool `json:"IsContainerd,omitempty" name:"IsContainerd"`
	// 主机来源：["CVM", "ECM", "LH", "BM"] 中的之一为腾讯云服务器；["Other"]之一非腾讯云服务器；

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 外网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 主机实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 地域ID

	RegionID *int64 `json:"RegionID,omitempty" name:"RegionID"`
	// 所属项目

	Project *ProjectInfo `json:"Project,omitempty" name:"Project"`
	// 标签

	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`
	// 是否开启服务

	StartTheService *bool `json:"StartTheService,omitempty" name:"StartTheService"`
	// 集群id

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群接入状态

	ClusterAccessedStatus *string `json:"ClusterAccessedStatus,omitempty" name:"ClusterAccessedStatus"`
	// 计费核数

	ChargeCoresCnt *uint64 `json:"ChargeCoresCnt,omitempty" name:"ChargeCoresCnt"`
	// 防护状态: 已防护: Defended 未防护: UnDefended

	DefendStatus *string `json:"DefendStatus,omitempty" name:"DefendStatus"`
}

type DescribeUserPurchaseInfoExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为10000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeUserPurchaseInfoExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserPurchaseInfoExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceAssetSummary struct {
	// 资产类别。

	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`
	// 是否为客户的首次检测。与CheckStatus配合使用。

	IsCustomerFirstCheck *bool `json:"IsCustomerFirstCheck,omitempty" name:"IsCustomerFirstCheck"`
	// 检测状态

	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 此类别的检测进度，为 0~100 的数。若未在检测中，无此字段。

	CheckProgress *float64 `json:"CheckProgress,omitempty" name:"CheckProgress"`
	// 此类资产通过的检测项的数目。

	PassedPolicyItemCount *uint64 `json:"PassedPolicyItemCount,omitempty" name:"PassedPolicyItemCount"`
	// 此类资产未通过的检测的数目。

	FailedPolicyItemCount *uint64 `json:"FailedPolicyItemCount,omitempty" name:"FailedPolicyItemCount"`
	// 此类资产下未通过的严重级别的检测项的数目。

	FailedCriticalPolicyItemCount *uint64 `json:"FailedCriticalPolicyItemCount,omitempty" name:"FailedCriticalPolicyItemCount"`
	// 此类资产下未通过的高危检测项的数目。

	FailedHighRiskPolicyItemCount *uint64 `json:"FailedHighRiskPolicyItemCount,omitempty" name:"FailedHighRiskPolicyItemCount"`
	// 此类资产下未通过的中危检测项的数目。

	FailedMediumRiskPolicyItemCount *uint64 `json:"FailedMediumRiskPolicyItemCount,omitempty" name:"FailedMediumRiskPolicyItemCount"`
	// 此类资产下未通过的低危检测项的数目。

	FailedLowRiskPolicyItemCount *uint64 `json:"FailedLowRiskPolicyItemCount,omitempty" name:"FailedLowRiskPolicyItemCount"`
	// 此类资产下提示级别的检测项的数目。

	NoticePolicyItemCount *uint64 `json:"NoticePolicyItemCount,omitempty" name:"NoticePolicyItemCount"`
	// 通过检测的资产的数目。

	PassedAssetCount *uint64 `json:"PassedAssetCount,omitempty" name:"PassedAssetCount"`
	// 未通过检测的资产的数目。

	FailedAssetCount *uint64 `json:"FailedAssetCount,omitempty" name:"FailedAssetCount"`
	// 此类资产的合规率，0~100的数。

	AssetPassedRate *float64 `json:"AssetPassedRate,omitempty" name:"AssetPassedRate"`
	// 检测失败的资产的数目。

	ScanFailedAssetCount *uint64 `json:"ScanFailedAssetCount,omitempty" name:"ScanFailedAssetCount"`
	// 上次检测的耗时，单位为秒。

	CheckCostTime *float64 `json:"CheckCostTime,omitempty" name:"CheckCostTime"`
	// 上次检测的时间。

	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`
	// 定时检测规则。

	PeriodRule *CompliancePeriodTaskRule `json:"PeriodRule,omitempty" name:"PeriodRule"`
	// 开启的规则项数量

	OpenPolicyItemCount *uint64 `json:"OpenPolicyItemCount,omitempty" name:"OpenPolicyItemCount"`
	// 忽略的规则项数量

	IgnoredPolicyItemCount *uint64 `json:"IgnoredPolicyItemCount,omitempty" name:"IgnoredPolicyItemCount"`
}

type WhiteListNodeInfo struct {
	// 加入白名单的时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 白名单节点Id

	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 节点的角色，Master、Work等

	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
	// 内网ip地址

	PrivateIpAddresses *string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
}

type DescribeReverseShellEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 反弹shell事件数组

		EventSet []*ReverseShellEventInfo `json:"EventSet,omitempty" name:"EventSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComplianceAssetListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回资产的总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回各类资产的列表。

		AssetInfoList []*ComplianceAssetInfo `json:"AssetInfoList,omitempty" name:"AssetInfoList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComplianceAssetListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComplianceAssetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAbnormalProcessRulesExportJobRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>RuleType - string  - 是否必填: 否 -规则类型</li>
	// <li>Status - string  - 是否必填: 否 -状态</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateAbnormalProcessRulesExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAbnormalProcessRulesExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeESAggregationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ES聚合结果JSON

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeESAggregationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeESAggregationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceHostDetailInfo struct {
	// 主机上的Docker版本。

	DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`
	// 主机上的K8S的版本。

	K8SVersion *string `json:"K8SVersion,omitempty" name:"K8SVersion"`
}

type ImageHost struct {
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 主机id

	HostID *string `json:"HostID,omitempty" name:"HostID"`
}

type VulIgnoreLocalImage struct {
	// 记录ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像大小

	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`
	// 漏洞PocID

	PocID *string `json:"PocID,omitempty" name:"PocID"`
	// 租户appid

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type CreateImageExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// excel文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerAssetSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 容器总数

		ContainerTotalCnt *uint64 `json:"ContainerTotalCnt,omitempty" name:"ContainerTotalCnt"`
		// 正在运行容器数量

		ContainerRunningCnt *uint64 `json:"ContainerRunningCnt,omitempty" name:"ContainerRunningCnt"`
		// 暂停运行容器数量

		ContainerPauseCnt *uint64 `json:"ContainerPauseCnt,omitempty" name:"ContainerPauseCnt"`
		// 停止运行容器数量

		ContainerStopped *uint64 `json:"ContainerStopped,omitempty" name:"ContainerStopped"`
		// 本地镜像数量

		ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`
		// 主机节点数量

		HostCnt *uint64 `json:"HostCnt,omitempty" name:"HostCnt"`
		// 主机正在运行节点数量

		HostRunningCnt *uint64 `json:"HostRunningCnt,omitempty" name:"HostRunningCnt"`
		// 主机离线节点数量

		HostOfflineCnt *uint64 `json:"HostOfflineCnt,omitempty" name:"HostOfflineCnt"`
		// 镜像仓库数量

		ImageRegistryCnt *uint64 `json:"ImageRegistryCnt,omitempty" name:"ImageRegistryCnt"`
		// 镜像总数

		ImageTotalCnt *uint64 `json:"ImageTotalCnt,omitempty" name:"ImageTotalCnt"`
		// 主机未安装agent数量

		HostUnInstallCnt *uint64 `json:"HostUnInstallCnt,omitempty" name:"HostUnInstallCnt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContainerAssetSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContainerAssetSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群的详细信息

		ClusterInfoList []*ClusterInfoItem `json:"ClusterInfoList,omitempty" name:"ClusterInfoList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserClusterExportJobRequest struct {
	*tchttp.BaseRequest

	// 集群ID列表，为空表示导出全部

	ClusterIdList []*string `json:"ClusterIdList,omitempty" name:"ClusterIdList"`
	// 支持按appid过滤

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *CreateUserClusterExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserClusterExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AffectedWorkloadItem struct {
	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 工作负载名称

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
	// 工作负载类型

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 检测结果的验证信息

	VerifyInfo *string `json:"VerifyInfo,omitempty" name:"VerifyInfo"`
}

type DescribeUserPurchaseInfoRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeUserPurchaseInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserPurchaseInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 访问控制策略信息列表

		RuleSet []*RuleBaseInfo `json:"RuleSet,omitempty" name:"RuleSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessControlRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type K8sApiAbnormalRuleInfo struct {
	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 状态

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 规则信息列表

	RuleInfoList []*K8sApiAbnormalRuleScopeInfo `json:"RuleInfoList,omitempty" name:"RuleInfoList"`
	// 生效集群IDSet

	EffectClusterIDSet []*string `json:"EffectClusterIDSet,omitempty" name:"EffectClusterIDSet"`
	// 规则类型
	// RT_SYSTEM 系统规则
	// RT_USER 用户自定义

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 是否所有集群生效

	EffectAllCluster *bool `json:"EffectAllCluster,omitempty" name:"EffectAllCluster"`
}

type ReverseShellWhiteListInfo struct {
	// 目标IP

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 目标端口

	DstPort *string `json:"DstPort,omitempty" name:"DstPort"`
	// 目标进程

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 镜像id数组，为空代表全部

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
	// 白名单id，如果新建则id为空

	Id *string `json:"Id,omitempty" name:"Id"`
}

type AssetSimpleImageInfo struct {
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 关联容器个数

	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
	// 最后扫描时间

	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
	// 镜像大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
}

type ImageDenyRule struct {
	// 规则RuleID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 规则类型 RULE_RISK：风险， RULE_PRIVILEGE：特权

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 生效的镜像数量

	EffectImageCount *int64 `json:"EffectImageCount,omitempty" name:"EffectImageCount"`
	// 是否对全部扫描镜像生效。0:全选镜像，1:自选镜像

	IsEffectAllImage *int64 `json:"IsEffectAllImage,omitempty" name:"IsEffectAllImage"`
	// 规则开始生效时间

	EffectTime *string `json:"EffectTime,omitempty" name:"EffectTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 操作用户

	OperationUin *string `json:"OperationUin,omitempty" name:"OperationUin"`
	// 启用状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 生效状态 IN_THE_TEST ：观察中，IN_EFFECT：生效中

	EffectStatus *string `json:"EffectStatus,omitempty" name:"EffectStatus"`
	// 规则ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 用户id

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type CreateImageVulListExportJobRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Name- String - 是否必填：否 - 漏洞名称名称筛选，</li>
	// <li>Level - String - 是否必填：否 - 风险等级  1,2,3,4</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
}

func (r *CreateImageVulListExportJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageVulListExportJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEscapeEventsExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEscapeEventsExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEscapeEventsExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWebVulExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWebVulExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWebVulExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlEventsExportRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	// <li>AppId - String - 是否必填：否 - 租户id筛选</li>
	// <li>FilterTimeRange - String[] - 是否必填：否 - 筛选时间，第一个值为开始时间，第二个值为结束时间["2020-1-1 12:00:00", "2020-1-1 12:00:00"]</li>

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 升序降序,asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 导出字段

	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeAccessControlEventsExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccessControlEventsExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImagesInfo struct {
	// 镜像id

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 镜像大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 主机个数

	HostCnt *uint64 `json:"HostCnt,omitempty" name:"HostCnt"`
	// 容器个数

	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
	// 扫描时间

	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
	// 漏洞个数

	VulCnt *uint64 `json:"VulCnt,omitempty" name:"VulCnt"`
	// 病毒个数

	VirusCnt *uint64 `json:"VirusCnt,omitempty" name:"VirusCnt"`
	// 敏感信息个数

	RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`
	// 是否信任镜像

	IsTrustImage *bool `json:"IsTrustImage,omitempty" name:"IsTrustImage"`
	// 镜像系统

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// agent镜像扫描错误

	AgentError *string `json:"AgentError,omitempty" name:"AgentError"`
	// 后端镜像扫描错误

	ScanError *string `json:"ScanError,omitempty" name:"ScanError"`
	// 扫描状态

	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`
	// 木马扫描错误信息

	ScanVirusError *string `json:"ScanVirusError,omitempty" name:"ScanVirusError"`
	// 漏洞扫描错误信息

	ScanVulError *string `json:"ScanVulError,omitempty" name:"ScanVulError"`
	// 风险扫描错误信息

	ScanRiskError *string `json:"ScanRiskError,omitempty" name:"ScanRiskError"`
	// 租户ID, 运营端使用

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 是否是重点关注镜像，为0不是，非0是

	IsSuggest *uint64 `json:"IsSuggest,omitempty" name:"IsSuggest"`
	// 是否授权，1是0否

	IsAuthorized *uint64 `json:"IsAuthorized,omitempty" name:"IsAuthorized"`
	// 组件个数

	ComponentCnt *uint64 `json:"ComponentCnt,omitempty" name:"ComponentCnt"`
}

type DescribeCheckItemListRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询的最大记录数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Name 可取值：risk_level风险等级, risk_target检查对象，风险对象,risk_type风险类别,risk_attri检测项所属的风险类型

	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCheckItemListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCheckItemListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetImageRegistryVirusListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>

	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
	// 镜像信息

	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`
	// 排序字段（Level）

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式 + -

	Order *string `json:"Order,omitempty" name:"Order"`
	// 镜像标识Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryVirusListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetImageRegistryVirusListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEmergencyVulExportJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID，前端拿着任务ID查询任务进度

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEmergencyVulExportJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEmergencyVulExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirusAutoIsolateSettingListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。

	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	By *string `json:"By,omitempty" name:"By"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeVirusAutoIsolateSettingListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirusAutoIsolateSettingListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalProcessLevelSummaryRequest struct {
	*tchttp.BaseRequest

	// 当前appid

	CurrentAppId *uint64 `json:"CurrentAppId,omitempty" name:"CurrentAppId"`
}

func (r *DescribeAbnormalProcessLevelSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalProcessLevelSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanIgnoreVulListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总是

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 漏洞列表

		List []*ScanIgnoreVul `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanIgnoreVulListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScanIgnoreVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProcessDetailInfo struct {
	// 进程名称

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 进程权限

	ProcessAuthority *string `json:"ProcessAuthority,omitempty" name:"ProcessAuthority"`
	// 进程pid

	ProcessId *uint64 `json:"ProcessId,omitempty" name:"ProcessId"`
	// 进程启动用户

	ProcessStartUser *string `json:"ProcessStartUser,omitempty" name:"ProcessStartUser"`
	// 进程用户组

	ProcessUserGroup *string `json:"ProcessUserGroup,omitempty" name:"ProcessUserGroup"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 进程树

	ProcessTree *string `json:"ProcessTree,omitempty" name:"ProcessTree"`
	// 进程md5

	ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`
	// 进程命令行参数

	ProcessParam *string `json:"ProcessParam,omitempty" name:"ProcessParam"`
}

type VirusInfo struct {
	// 文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件路径

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 病毒名称

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 容器id

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING

	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// DEAL_NONE:文件待处理
	// DEAL_IGNORE:已经忽略
	// DEAL_ADD_WHITELIST:加白
	// DEAL_DEL:文件已经删除
	// DEAL_ISOLATE:已经隔离
	// DEAL_ISOLATING:隔离中
	// DEAL_ISOLATE_FAILED:隔离失败
	// DEAL_RECOVERING:恢复中
	// DEAL_RECOVER_FAILED: 恢复失败

	Status *string `json:"Status,omitempty" name:"Status"`
	// 事件id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 事件描述

	HarmDescribe *string `json:"HarmDescribe,omitempty" name:"HarmDescribe"`
	// 建议方案

	SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`
	// 失败子状态:
	// FILE_NOT_FOUND:文件不存在
	// FILE_ABNORMAL:文件异常
	// FILE_ABNORMAL_DEAL_RECOVER:恢复文件时，文件异常
	// BACKUP_FILE_NOT_FOUND:备份文件不存在
	// CONTAINER_NOT_FOUND_DEAL_ISOLATE:隔离时，容器不存在
	// CONTAINER_NOT_FOUND_DEAL_RECOVER:恢复时，容器不存在
	// TIMEOUT: 超时
	// TOO_MANY: 任务过多
	// OFFLINE: 离线
	// INTERNAL: 服务内部错误
	// VALIDATION: 参数非法

	SubStatus *string `json:"SubStatus,omitempty" name:"SubStatus"`
	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED

	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`
	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// 	"NODE_DESTROYED"      //节点已销毁
	// 	"CONTAINER_EXITED"    //容器已退出
	// 	"CONTAINER_DESTROYED" //容器已销毁
	// 	"SHARED_HOST"         // 容器与主机共享网络
	// 	"RESOURCE_LIMIT"      //隔离操作资源超限
	// 	"UNKNOW"              // 原因未知

	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`
	// 容器隔离操作来源

	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
	// md5值

	MD5 *string `json:"MD5,omitempty" name:"MD5"`
	// 风险等级 RISK_CRITICAL, RISK_HIGH, RISK_MEDIUM, RISK_LOW, RISK_NOTICE。

	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 检测平台
	// 1: 云查杀引擎
	// 2: tav
	// 3: binaryAi
	// 4: 异常行为
	// 5: 威胁情报

	CheckPlatform []*string `json:"CheckPlatform,omitempty" name:"CheckPlatform"`
	// AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
}

type DescribeWebVulListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 漏洞列表

		List []*VulInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebVulListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AbnormalProcessSystemChildRuleInfo struct {
	// 子策略Id

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 子策略状态，true为开启，false为关闭

	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
	// 策略模式,  RULE_MODE_RELEASE: 放行
	//    RULE_MODE_ALERT: 告警
	//    RULE_MODE_HOLDUP:拦截

	RuleMode *string `json:"RuleMode,omitempty" name:"RuleMode"`
	// 子策略检测的行为类型
	// PROXY_TOOL： 代理软件
	// TRANSFER_CONTROL：横向渗透
	// ATTACK_CMD： 恶意命令
	// REVERSE_SHELL：反弹shell
	// FILELESS：无文件程序执行
	// RISK_CMD：高危命令
	// ABNORMAL_CHILD_PROC: 敏感服务异常子进程启动

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 威胁等级，HIGH:高，MIDDLE:中，LOW:低

	RuleLevel *string `json:"RuleLevel,omitempty" name:"RuleLevel"`
}
