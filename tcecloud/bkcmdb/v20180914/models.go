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

package v20180914

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type HostDisplayColumn struct {
	// 属性ID

	BkPropertyId *string `json:"BkPropertyId,omitempty" name:"BkPropertyId"`
	// 属性名称

	BkPropertyName *string `json:"BkPropertyName,omitempty" name:"BkPropertyName"`
	// 实例ID

	BkObjId *string `json:"BkObjId,omitempty" name:"BkObjId"`
}

type HostInfoPro struct {
	// ID

	BkAssetId *string `json:"BkAssetId,omitempty" name:"BkAssetId"`
	// CPU

	BkCpuModule *string `json:"BkCpuModule,omitempty" name:"BkCpuModule"`
	// BkDeviceDescript

	BkDeviceDescript *string `json:"BkDeviceDescript,omitempty" name:"BkDeviceDescript"`
	// BkDeviceHeight

	BkDeviceHeight *string `json:"BkDeviceHeight,omitempty" name:"BkDeviceHeight"`
	// 驱动类型

	BkDeviceType *string `json:"BkDeviceType,omitempty" name:"BkDeviceType"`
	// 磁盘信息

	BkDisk *string `json:"BkDisk,omitempty" name:"BkDisk"`
	// 主机内网IP

	BkHostInnerip *string `json:"BkHostInnerip,omitempty" name:"BkHostInnerip"`
	// 主机名称

	BkHostName *string `json:"BkHostName,omitempty" name:"BkHostName"`
	// 主机外网IP

	BkHostOuterip *string `json:"BkHostOuterip,omitempty" name:"BkHostOuterip"`
	// 主机状态

	BkHostStatus *string `json:"BkHostStatus,omitempty" name:"BkHostStatus"`
	// 主机类型

	BkHostType *string `json:"BkHostType,omitempty" name:"BkHostType"`
	// IDC ID

	BkIdcId *string `json:"BkIdcId,omitempty" name:"BkIdcId"`
	// IDC名称

	BkIdcName *string `json:"BkIdcName,omitempty" name:"BkIdcName"`
	// 网关

	BkLanGateway *string `json:"BkLanGateway,omitempty" name:"BkLanGateway"`
	// 子网掩码

	BkLanMask *string `json:"BkLanMask,omitempty" name:"BkLanMask"`
	// 逻辑区域

	BkLogicArea *string `json:"BkLogicArea,omitempty" name:"BkLogicArea"`
	// 主负责人

	BkMaintainer *string `json:"BkMaintainer,omitempty" name:"BkMaintainer"`
	// 内存

	BkMem *string `json:"BkMem,omitempty" name:"BkMem"`
	// 操作系统名称

	BkOsName *string `json:"BkOsName,omitempty" name:"BkOsName"`
	// 操作系统类型

	BkOsType *string `json:"BkOsType,omitempty" name:"BkOsType"`
	// BkPosId

	BkPosId *string `json:"BkPosId,omitempty" name:"BkPosId"`
	// BkProducer

	BkProducer *string `json:"BkProducer,omitempty" name:"BkProducer"`
	// BkRackName

	BkRackName *string `json:"BkRackName,omitempty" name:"BkRackName"`
	// region名称

	BkRegionName *string `json:"BkRegionName,omitempty" name:"BkRegionName"`
	// region名称缩写

	BkRegionNameShort *string `json:"BkRegionNameShort,omitempty" name:"BkRegionNameShort"`
	// BkDeviceHeight

	BkSn *string `json:"BkSn,omitempty" name:"BkSn"`
	// BkDeviceHeight

	BkSwitch *string `json:"BkSwitch,omitempty" name:"BkSwitch"`
	// 网关

	BkWanGateway *string `json:"BkWanGateway,omitempty" name:"BkWanGateway"`
	// 子网掩码

	BkWanMask *string `json:"BkWanMask,omitempty" name:"BkWanMask"`
	// zone名称

	BkZoneName *string `json:"BkZoneName,omitempty" name:"BkZoneName"`
	// 主机IP

	BkHostIp *string `json:"BkHostIp,omitempty" name:"BkHostIp"`
	// 虚拟机状态

	BkVmStatus *string `json:"BkVmStatus,omitempty" name:"BkVmStatus"`
	// 云平台ID

	BkCloudId *string `json:"BkCloudId,omitempty" name:"BkCloudId"`
	// 数据来源

	ImportFrom *string `json:"ImportFrom,omitempty" name:"ImportFrom"`
	// 备运维人员

	BkBakMaintainer *string `json:"BkBakMaintainer,omitempty" name:"BkBakMaintainer"`
	// raid级别

	BkRaidType *string `json:"BkRaidType,omitempty" name:"BkRaidType"`
}

type Option struct {
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
	// 无

	Type *string `json:"Type,omitempty" name:"Type"`
}

type RequestElementKeyValue struct {
	// 键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeTagsFilter struct {
	// 标签key

	Name *string `json:"Name,omitempty" name:"Name"`
	// 标签value

	Value []*string `json:"Value,omitempty" name:"Value"`
}

type Ext struct{}

type FilterValue struct {
	// 标签key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签value

	Value []*string `json:"Value,omitempty" name:"Value"`
}

type Position struct{}

type SearchHostRequest struct {
	*tchttp.BaseRequest

	// 主机ip列表

	Ip *Ip `json:"Ip,omitempty" name:"Ip"`
	// 组合条件

	Condition []*Condition `json:"Condition,omitempty" name:"Condition"`
	// 查询条件

	Page *Page `json:"Page,omitempty" name:"Page"`
	// 按表达式搜索

	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`
	// 业务ID

	BkBizId *int64 `json:"BkBizId,omitempty" name:"BkBizId"`
	// 业务状态

	BkBusinessStatus *string `json:"BkBusinessStatus,omitempty" name:"BkBusinessStatus"`
}

func (r *SearchHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Condition struct {
	// bk_data_status

	BkDataStatus *BkDataStatus `json:"BkDataStatus,omitempty" name:"BkDataStatus"`
	// 业务名

	BkBizName *string `json:"BkBizName,omitempty" name:"BkBizName"`
	// 运维人员

	BkBizMaintainer *string `json:"BkBizMaintainer,omitempty" name:"BkBizMaintainer"`
	// 产品人员

	BkBizProductor *string `json:"BkBizProductor,omitempty" name:"BkBizProductor"`
	// 开发人员

	BkBizDeveloper *string `json:"BkBizDeveloper,omitempty" name:"BkBizDeveloper"`
	// 测试人员

	BkBizTester *string `json:"BkBizTester,omitempty" name:"BkBizTester"`
	// 生命周期

	LifeCycle *string `json:"LifeCycle,omitempty" name:"LifeCycle"`
	// 操作人员

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 时区

	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 对象名,可以为biz,set,module,host,object

	BkObjId *string `json:"BkObjId,omitempty" name:"BkObjId"`
	// 查询输出字段

	Fields []*string `json:"Fields,omitempty" name:"Fields"`
	// 查询条件

	Condition []*InnerCondition `json:"Condition,omitempty" name:"Condition"`
	// 此处仅为示例数据，需要被设置为模型的标识符，在页面上配置的英文名

	BkWeblogic []*InnerCondition `json:"BkWeblogic,omitempty" name:"BkWeblogic"`
	// 父实例节点的ID，当前实例节点的上一级实例节点，在拓扑结构中对于Module一般指的是Set的bk_set_id

	BkParentId *int64 `json:"BkParentId,omitempty" name:"BkParentId"`
	// 集群id

	BkSetId *int64 `json:"BkSetId,omitempty" name:"BkSetId"`
	// 集群名称

	BkSetName *string `json:"BkSetName,omitempty" name:"BkSetName"`
	// 设计容量

	BkCapacity *int64 `json:"BkCapacity,omitempty" name:"BkCapacity"`
	// 备注

	Description *string `json:"Description,omitempty" name:"Description"`
	// 服务状态:1/2(1:开放,2:关闭)

	BkServiceStatus *uint64 `json:"BkServiceStatus,omitempty" name:"BkServiceStatus"`
	// 环境类型：1/2/3(1:测试,2:体验,3:正式)

	BkSetEnv *uint64 `json:"BkSetEnv,omitempty" name:"BkSetEnv"`
	// 集群描述

	BkSetDesc *string `json:"BkSetDesc,omitempty" name:"BkSetDesc"`
	// 内置模块:0-普通模块,1-空闲机,2-故障机

	Default *int64 `json:"Default,omitempty" name:"Default"`
	// 模块标识

	BkModuleId *int64 `json:"BkModuleId,omitempty" name:"BkModuleId"`
	// 模块名

	BkModuleName *string `json:"BkModuleName,omitempty" name:"BkModuleName"`
	// 开发商账号

	BkSupplierAccount *string `json:"BkSupplierAccount,omitempty" name:"BkSupplierAccount"`
	// 模块类型：1/2 (1:普通, 2:数据库)

	BkModuleType *uint64 `json:"BkModuleType,omitempty" name:"BkModuleType"`
	// 备份维护人

	BkBakOperator *string `json:"BkBakOperator,omitempty" name:"BkBakOperator"`
	// 业务ID

	BkBizId *int64 `json:"BkBizId,omitempty" name:"BkBizId"`
	// 操作对象，可以为biz host process set module object

	OpTarget *string `json:"OpTarget,omitempty" name:"OpTarget"`
	// 操作类型， add delete update

	OpType *string `json:"OpType,omitempty" name:"OpType"`
	// 没有条件，为空, 开始和结束时间成对出现

	OpTime []*string `json:"OpTime,omitempty" name:"OpTime"`
	// ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// BkFunc

	BkFunc *string `json:"BkFunc,omitempty" name:"BkFunc"`
	// BkBizStatus

	BkBizStatus *string `json:"BkBizStatus,omitempty" name:"BkBizStatus"`
	// 操作者

	BkOperator *string `json:"BkOperator,omitempty" name:"BkOperator"`
	// 详情

	BkDetail *string `json:"BkDetail,omitempty" name:"BkDetail"`
	// 模型

	BkModel *string `json:"BkModel,omitempty" name:"BkModel"`
	// 操作系统详情

	BkOsDetail *string `json:"BkOsDetail,omitempty" name:"BkOsDetail"`
	// 物资号

	BkAssetId *string `json:"BkAssetId,omitempty" name:"BkAssetId"`
	// 管理员IP

	BkAdminIp *string `json:"BkAdminIp,omitempty" name:"BkAdminIp"`
	// Connect

	Connect *string `json:"Connect,omitempty" name:"Connect"`
	// BkSn

	BkSn *string `json:"BkSn,omitempty" name:"BkSn"`
	// BkVendor

	BkVendor *string `json:"BkVendor,omitempty" name:"BkVendor"`
	// 实例名称

	BkInstName *string `json:"BkInstName,omitempty" name:"BkInstName"`
	// 工作路径

	WorkPath *string `json:"WorkPath,omitempty" name:"WorkPath"`
	// AutoTimeGap

	AutoTimeGap *string `json:"AutoTimeGap,omitempty" name:"AutoTimeGap"`
	// 优先级

	Priority *string `json:"Priority,omitempty" name:"Priority"`
	// 重启命令

	RestartCmd *string `json:"RestartCmd,omitempty" name:"RestartCmd"`
	// 函数名称

	BkFuncName *string `json:"BkFuncName,omitempty" name:"BkFuncName"`
	// 启动命令

	StartCmd *string `json:"StartCmd,omitempty" name:"StartCmd"`
	// 进程数

	ProcNum *int64 `json:"ProcNum,omitempty" name:"ProcNum"`
	// 进程名称

	BkProcessName *string `json:"BkProcessName,omitempty" name:"BkProcessName"`
	// 用户名

	User *string `json:"User,omitempty" name:"User"`
	// 停止命令

	StopCmd *string `json:"StopCmd,omitempty" name:"StopCmd"`
	// 绑定IP

	BindIp *string `json:"BindIp,omitempty" name:"BindIp"`
	// 重载命令

	ReloadCmd *string `json:"ReloadCmd,omitempty" name:"ReloadCmd"`
	// 函数ID

	BkFuncId *int64 `json:"BkFuncId,omitempty" name:"BkFuncId"`
	// 时限

	Timeout *string `json:"Timeout,omitempty" name:"Timeout"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// FaceStopCmd

	FaceStopCmd *string `json:"FaceStopCmd,omitempty" name:"FaceStopCmd"`
	// 是否自动启动

	AutoStart *bool `json:"AutoStart,omitempty" name:"AutoStart"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// PID文件

	PidFile *string `json:"PidFile,omitempty" name:"PidFile"`
	// 实例ID

	InstId *int64 `json:"InstId,omitempty" name:"InstId"`
	// IsDefault

	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// ExtKey

	ExtKey *ExtKey `json:"ExtKey,omitempty" name:"ExtKey"`
	// 无

	BkPropertyId *string `json:"BkPropertyId,omitempty" name:"BkPropertyId"`
	// 无

	SurplusLanguage *string `json:"SurplusLanguage,omitempty" name:"SurplusLanguage"`
	// 无

	BkBizFollowUser *string `json:"BkBizFollowUser,omitempty" name:"BkBizFollowUser"`
	// 无

	BkBizBakMaintainer *string `json:"BkBizBakMaintainer,omitempty" name:"BkBizBakMaintainer"`
	// 主运维人员

	BkMaintainer *string `json:"BkMaintainer,omitempty" name:"BkMaintainer"`
	// 主运维人员

	BkBakMaintainer *string `json:"BkBakMaintainer,omitempty" name:"BkBakMaintainer"`
	// 产品人员

	BkProductor *string `json:"BkProductor,omitempty" name:"BkProductor"`
	// 研发人员

	BkDeveloper *string `json:"BkDeveloper,omitempty" name:"BkDeveloper"`
	// 测试人员

	BkTester *string `json:"BkTester,omitempty" name:"BkTester"`
	// 关注人

	BkFollowUser *string `json:"BkFollowUser,omitempty" name:"BkFollowUser"`
}

type FilterString struct {
	// key

	Name *string `json:"Name,omitempty" name:"Name"`
	// value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type InnerCondition struct {
	// 对象的字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// 操作符, $eq为相等，$neq为不等，$in为属于，$nin为不属于

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 字段对应的值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type Resource struct {
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

type Page struct {
	// 记录开始位置

	Start *int64 `json:"Start,omitempty" name:"Start"`
	// 每页限制条数,最大200

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

type BatchDeleteSet struct {
	// 集群ID

	InstIds []*int64 `json:"InstIds,omitempty" name:"InstIds"`
}

type Biz struct {
	// 业务ID

	BkBizId *int64 `json:"BkBizId,omitempty" name:"BkBizId"`
}

type BkDataStatus struct {
	// 等于

	TCEEq *string `json:"TCEEq,omitempty" name:"TCEEq"`
	// 不等于

	TCENe *string `json:"TCENe,omitempty" name:"TCENe"`
	// 属于

	TCEIn *string `json:"TCEIn,omitempty" name:"TCEIn"`
	// 不属于

	TCENin *string `json:"TCENin,omitempty" name:"TCENin"`
}

type ConditionPro struct {
	// 对象名,可以为biz,set,module,host,object

	BkObjId *string `json:"BkObjId,omitempty" name:"BkObjId"`
	// 查询输出字段

	Fields []*string `json:"Fields,omitempty" name:"Fields"`
	// 查询条件

	Condition []*InnerConditionPro `json:"Condition,omitempty" name:"Condition"`
}

type Delete struct {
	// 实例ID集合

	InstIds []*int64 `json:"InstIds,omitempty" name:"InstIds"`
}

type Filter struct {
	// 筛选的key

	Name *string `json:"Name,omitempty" name:"Name"`
	// 筛选的value

	Value *FilterValue `json:"Value,omitempty" name:"Value"`
}

type InnerConditionPro struct {
	// 对象的字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// 操作符,$in为属于，$nin为不属于

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 字段对应的值

	Value []*int64 `json:"Value,omitempty" name:"Value"`
}

type Data struct {
	// 模型分类的图标,取值可参考，取值可参考(classIcon.json)

	BkClassificationIcon *string `json:"BkClassificationIcon,omitempty" name:"BkClassificationIcon"`
	// 分类ID，英文描述用于系统内部使用

	BkClassificationId *string `json:"BkClassificationId,omitempty" name:"BkClassificationId"`
	// 分类名

	BkClassificationName *string `json:"BkClassificationName,omitempty" name:"BkClassificationName"`
	// 用于对分类进行分类（如：inner代码为内置分类，空字符串为自定义分类）

	BkClassificationType *string `json:"BkClassificationType,omitempty" name:"BkClassificationType"`
	// 数据记录ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// BkGroupName

	BkGroupName *string `json:"BkGroupName,omitempty" name:"BkGroupName"`
	// 无

	BkGroupIndex *int64 `json:"BkGroupIndex,omitempty" name:"BkGroupIndex"`
	// 无

	BkPropertyGroup *string `json:"BkPropertyGroup,omitempty" name:"BkPropertyGroup"`
	// 无

	BkPropertyIndex *int64 `json:"BkPropertyIndex,omitempty" name:"BkPropertyIndex"`
}

type ExtKey struct {
	// $in

	TCEIn []*string `json:"TCEIn,omitempty" name:"TCEIn"`
}

type Info struct {
	// 业务信息查询条件

	Biz *Biz `json:"Biz,omitempty" name:"Biz"`
	// 是否精确查询

	Exact *int64 `json:"Exact,omitempty" name:"Exact"`
	// true 或者false

	BkHostInnerip *bool `json:"BkHostInnerip,omitempty" name:"BkHostInnerip"`
	// true 或者alse

	BkHostOuterip *bool `json:"BkHostOuterip,omitempty" name:"BkHostOuterip"`
}

type Ip struct {
	// 是否根据ip精确搜索

	Exact *int64 `json:"Exact,omitempty" name:"Exact"`
	// bk_host_innerip只匹配内网ip,bk_host_outerip只匹配外网ip, bk_host_innerip,bk_host_outerip同时匹配

	Flag *string `json:"Flag,omitempty" name:"Flag"`
	// ip list for search

	Data []*string `json:"Data,omitempty" name:"Data"`
}

type UrlPara struct {
	// api版本号

	Version *string `json:"Version,omitempty" name:"Version"`
	// 业务id

	BkBizId *string `json:"BkBizId,omitempty" name:"BkBizId"`
	// 开发商账号

	BkSupplierAccount *string `json:"BkSupplierAccount,omitempty" name:"BkSupplierAccount"`
	// 模型ID

	BkObjId *string `json:"BkObjId,omitempty" name:"BkObjId"`
	// 主机ID

	BkHostId *int64 `json:"BkHostId,omitempty" name:"BkHostId"`
	// 记录开始位置

	Start *int64 `json:"Start,omitempty" name:"Start"`
	// 每页限制条数,最大200

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 收藏的主键

	Id *string `json:"Id,omitempty" name:"Id"`
	// 图形范围类型,可选global,biz,cls(当前只有global)

	ScopeType *string `json:"ScopeType,omitempty" name:"ScopeType"`
	// 图形范围类型下的ID,如果为global,则填0

	ScopeId *string `json:"ScopeId,omitempty" name:"ScopeId"`
	// 更新方法,可选update,override

	Action *string `json:"Action,omitempty" name:"Action"`
	// 实例ID

	BkInstId *string `json:"BkInstId,omitempty" name:"BkInstId"`
	// 分组ID

	BkGroupId *string `json:"BkGroupId,omitempty" name:"BkGroupId"`
	// 属性ID

	BkPropertyId *string `json:"BkPropertyId,omitempty" name:"BkPropertyId"`
	// 拓扑的层级索引，索引取值从0开始，当设置为 -1 的时候会读取完整的业务实例拓扑

	Level *string `json:"Level,omitempty" name:"Level"`
	// 集群id

	BkSetId *int64 `json:"BkSetId,omitempty" name:"BkSetId"`
	// 模块id

	BkModuleId *int64 `json:"BkModuleId,omitempty" name:"BkModuleId"`
	// 进程 id

	BkProcessId *int64 `json:"BkProcessId,omitempty" name:"BkProcessId"`
	// 模块名称

	BkModuleName *string `json:"BkModuleName,omitempty" name:"BkModuleName"`
	// 分组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 订阅ID

	SubscriptionId *int64 `json:"SubscriptionId,omitempty" name:"SubscriptionId"`
	// 收藏ID

	FavoriteId *string `json:"FavoriteId,omitempty" name:"FavoriteId"`
	// 无

	Flag *string `json:"Flag,omitempty" name:"Flag"`
	// 无

	InstId *int64 `json:"InstId,omitempty" name:"InstId"`
	// 无

	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`
	// 无

	BizId *int64 `json:"BizId,omitempty" name:"BizId"`
}

type SearchHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误码

		BkErrorCode *string `json:"BkErrorCode,omitempty" name:"BkErrorCode"`
		// 错误信息

		BkErrorMsg *string `json:"BkErrorMsg,omitempty" name:"BkErrorMsg"`
		// 数据

		Data *string `json:"Data,omitempty" name:"Data"`
		// 请求是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostDeviceTypeConfig struct {
	// 机型

	HostDeviceType *string `json:"HostDeviceType,omitempty" name:"HostDeviceType"`
	// CPU

	HostCpu *string `json:"HostCpu,omitempty" name:"HostCpu"`
	// 内存

	HostMem *string `json:"HostMem,omitempty" name:"HostMem"`
	// 磁盘

	HostDisk *string `json:"HostDisk,omitempty" name:"HostDisk"`
	// 网络

	HostNetwork *string `json:"HostNetwork,omitempty" name:"HostNetwork"`
	// PCE-E板卡

	HostPciE *string `json:"HostPciE,omitempty" name:"HostPciE"`
	// 是否启用
	// true:启动
	// false:停用

	IsUsed *bool `json:"IsUsed,omitempty" name:"IsUsed"`
}

type HostInfo struct {
	// 主机内网ip

	BkHostInnerip *string `json:"BkHostInnerip,omitempty" name:"BkHostInnerip"`
	// 云区域ID

	BkCloudId *int64 `json:"BkCloudId,omitempty" name:"BkCloudId"`
}

type QueryParams struct {
	// BkBizId

	BkBizId *int64 `json:"BkBizId,omitempty" name:"BkBizId"`
	// BkObjId

	BkObjId *int64 `json:"BkObjId,omitempty" name:"BkObjId"`
	// Field

	Field *string `json:"Field,omitempty" name:"Field"`
	// Operator

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type Update struct {
	// 实例ID

	InstId *int64 `json:"InstId,omitempty" name:"InstId"`
	// Datas

	Datas *Condition `json:"Datas,omitempty" name:"Datas"`
}

type BkWeblogic struct {
	// 操作者

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeleteTag struct {
	// 标签key

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
}

type Process struct {
	// 进程名称

	BkProcessName *string `json:"BkProcessName,omitempty" name:"BkProcessName"`
	// 进程启动命令

	StartCmd *string `json:"StartCmd,omitempty" name:"StartCmd"`
	// 是否正则
	// true:正则
	// false:非正则

	IsRegex *bool `json:"IsRegex,omitempty" name:"IsRegex"`
	// 进程监听的端口

	Port []*int64 `json:"Port,omitempty" name:"Port"`
	// 是否启用
	// true:启动
	// false:停用

	IsUsed *bool `json:"IsUsed,omitempty" name:"IsUsed"`
}

type PropertyGroupInfo struct {
	// 无

	Condition *Condition `json:"Condition,omitempty" name:"Condition"`
	// 无

	Data *Data `json:"Data,omitempty" name:"Data"`
}

type ReplaceTag struct {
	// 标签key

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签value

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}
