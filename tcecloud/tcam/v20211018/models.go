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

package v20211018

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type GetServicePermListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误码

		Code *uint64 `json:"Code,omitempty" name:"Code"`
		// 错误信息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 接口数组

		Data []*ServicePermItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServicePermListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServicePermListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStrategyDetailRequest struct {
	*tchttp.BaseRequest

	// 策略id

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *GetStrategyDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStrategyDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPoliciesRequest struct {
	*tchttp.BaseRequest

	// 每页数量，默认值是 20，必须大于 0 且小于或等于 200

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
	// 页码，默认值是 1，从 1开始，不能大于 200

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 可取值 'All'、'QCS' 和 'Local'，'All' 获取所有策略，'QCS' 只获取预设策略，'Local' 只获取自定义策略，默认取 'All'

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 按策略名匹配

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 按Uin匹配

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 按组Id匹配

	TargetGroupId *uint64 `json:"TargetGroupId,omitempty" name:"TargetGroupId"`
	// 按角色Id匹配

	TargetRoleId *uint64 `json:"TargetRoleId,omitempty" name:"TargetRoleId"`
	// 按产品Id匹配，如cvm

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 调用来源，控制台为"console"

	Client *string `json:"Client,omitempty" name:"Client"`
	// 按Uin标记关联

	FlagUin *uint64 `json:"FlagUin,omitempty" name:"FlagUin"`
	// 按GroupId标记关联

	FlagGroupId *uint64 `json:"FlagGroupId,omitempty" name:"FlagGroupId"`
	// 按角色Id标记关联

	FlagRoleId *uint64 `json:"FlagRoleId,omitempty" name:"FlagRoleId"`
	// 创建类型：1.按产品功能或项目权限创建；2.按策略语法创建；3.按策略生成器创建；4.按标签授权创建；5.按权限边界规则创建

	CreateMode *uint64 `json:"CreateMode,omitempty" name:"CreateMode"`
}

func (r *ListPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceRequest struct {
	*tchttp.BaseRequest

	// 操作者

	Writer *string `json:"Writer,omitempty" name:"Writer"`
	// 服务

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 操作，update、delete、create

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 资源列表 json字符串

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 服务英文名

	ServiceEnName *string `json:"ServiceEnName,omitempty" name:"ServiceEnName"`
}

func (r *UpdateServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RolePolicy struct {
	// 英文描述

	EnRemark *string `json:"EnRemark,omitempty" name:"EnRemark"`
	// 中文描述

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 英文名称

	StrategyEnName *string `json:"StrategyEnName,omitempty" name:"StrategyEnName"`
	// 策略Id

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
	//  策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
}

type ListPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略数组，数组每个成员包括 policyId、policyName、addTime、type、description、 createMode 字段。其中： policyId：策略 id policyName：策略名addTime：策略创建时间type：1 表示自定义策略，2 表示预设策略 description：策略描述 createMode：1 表示按业务权限创建的策略，其他值表示可以查看策略语法和通过策略语法更新策略Attachments: 关联的用户数ServiceType: 策略关联的产品IsAttached: 当需要查询标记实体是否已经关联策略时不为null。0表示未关联策略，1表示已关联策略

		List []*StrategyInfo `json:"List,omitempty" name:"List"`
		// 策略总数

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePresetStrategyRequest struct {
	*tchttp.BaseRequest

	// 操作者

	Writer *string `json:"Writer,omitempty" name:"Writer"`
	// 服务

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 策略名

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 策略英文名

	StrategyEnName *string `json:"StrategyEnName,omitempty" name:"StrategyEnName"`
	// 策略描述

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 策略英文描述

	EnRemark *string `json:"EnRemark,omitempty" name:"EnRemark"`
	// 策略内容

	StrategyInfo *string `json:"StrategyInfo,omitempty" name:"StrategyInfo"`
	// 策略ID

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 操作，update、delete、create

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// 是否可见 0可见

	IsDelete *uint64 `json:"IsDelete,omitempty" name:"IsDelete"`
	// 是否可见 默认1可见

	IsSeen *string `json:"IsSeen,omitempty" name:"IsSeen"`
}

func (r *UpdatePresetStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePresetStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPresetStrategyListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Code *uint64 `json:"Code,omitempty" name:"Code"`
		// 消息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 预设策略列表

		Data []*PresetStrategy `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPresetStrategyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPresetStrategyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceRoleInfoRequest struct {
	*tchttp.BaseRequest

	// 操作者

	Writer *string `json:"Writer,omitempty" name:"Writer"`
	// 服务

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 操作，update、delete、create

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// 角色名

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 关联策略名

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 角色描述

	RoleDesc *string `json:"RoleDesc,omitempty" name:"RoleDesc"`
	// 角色英文描述

	RoleDescEn *string `json:"RoleDescEn,omitempty" name:"RoleDescEn"`
	// 服务角色id 更新必填

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *UpdateServiceRoleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateServiceRoleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceTypeItem struct {
	// 英文名

	ResourceEnName *string `json:"ResourceEnName,omitempty" name:"ResourceEnName"`
	// 中文名

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

type ServiceItem struct {
	// 创建时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// ArnDocument

	ArnDocument *string `json:"ArnDocument,omitempty" name:"ArnDocument"`
	// ColConf

	ColConf *string `json:"ColConf,omitempty" name:"ColConf"`
	// DefAddr

	DefAddr *string `json:"DefAddr,omitempty" name:"DefAddr"`
	// 默认策略列表

	DefaultStrategyList *string `json:"DefaultStrategyList,omitempty" name:"DefaultStrategyList"`
	// IsAllowDefProj

	IsAllowDefProj *string `json:"IsAllowDefProj,omitempty" name:"IsAllowDefProj"`
	// IsDisProject

	IsDisProject *string `json:"IsDisProject,omitempty" name:"IsDisProject"`
	// IsDisZone

	IsDisZone *string `json:"IsDisZone,omitempty" name:"IsDisZone"`
	// 是否可见

	IsSeen *string `json:"IsSeen,omitempty" name:"IsSeen"`
	// Online

	Online *string `json:"Online,omitempty" name:"Online"`
	// QueryAddr

	QueryAddr *string `json:"QueryAddr,omitempty" name:"QueryAddr"`
	// QueryInterface

	QueryInterface *string `json:"QueryInterface,omitempty" name:"QueryInterface"`
	// 服务英文名

	ServiceEnName *string `json:"ServiceEnName,omitempty" name:"ServiceEnName"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// SynInterface

	SynInterface *string `json:"SynInterface,omitempty" name:"SynInterface"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// Weight

	Weight *string `json:"Weight,omitempty" name:"Weight"`
	// WhiteKey

	WhiteKey *string `json:"WhiteKey,omitempty" name:"WhiteKey"`
	// 创建人

	Writter *string `json:"Writter,omitempty" name:"Writter"`
	// 资源类型数组

	ResourceTypeList []*ResourceTypeItem `json:"ResourceTypeList,omitempty" name:"ResourceTypeList"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type StrategyInfo struct {
	// 策略ID。

	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名称。

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 策略创建时间。

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 策略类型。1 表示自定义策略，2 表示预设策略。

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 策略描述。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建来源，1 通过控制台创建, 2 通过策略语法创建。

	CreateMode *uint64 `json:"CreateMode,omitempty" name:"CreateMode"`
	// 关联的用户数

	Attachments *uint64 `json:"Attachments,omitempty" name:"Attachments"`
	// 策略关联的产品

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 当需要查询标记实体是否已经关联策略时不为null。0表示未关联策略，1表示已关联策略

	IsAttached *uint64 `json:"IsAttached,omitempty" name:"IsAttached"`
	// 是否已下线

	Deactived *uint64 `json:"Deactived,omitempty" name:"Deactived"`
	// 已下线产品列表

	DeactivedDetail *string `json:"DeactivedDetail,omitempty" name:"DeactivedDetail"`
}

type GetPresetStrategyListRequest struct {
	*tchttp.BaseRequest

	// 服务

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 默认不传 传1为只展示可见策略

	IsSeen *uint64 `json:"IsSeen,omitempty" name:"IsSeen"`
	// 排序字段 默认updateTime

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段desc/asc 默认desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *GetPresetStrategyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPresetStrategyListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误码

		Code *uint64 `json:"Code,omitempty" name:"Code"`
		// 错误信息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 服务数组

		Data []*ServiceItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServicePermListRequest struct {
	*tchttp.BaseRequest

	// 业务标识 如：cvm

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 接口名

	InterfaceEnName *string `json:"InterfaceEnName,omitempty" name:"InterfaceEnName"`
	// 排序字段 默认updateTime

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段desc/asc  默认desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *GetServicePermListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServicePermListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PresetStrategy struct {
	// 添加时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 创建模式

	CreateMode *string `json:"CreateMode,omitempty" name:"CreateMode"`
	// 创建人

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// 英文描述

	EnRemark *string `json:"EnRemark,omitempty" name:"EnRemark"`
	// 是否检查

	IsCheck *string `json:"IsCheck,omitempty" name:"IsCheck"`
	// 是否删除

	IsDelete *string `json:"IsDelete,omitempty" name:"IsDelete"`
	// 是否资源权限

	IsOnResource *string `json:"IsOnResource,omitempty" name:"IsOnResource"`
	// 是否可见

	IsSeen *string `json:"IsSeen,omitempty" name:"IsSeen"`
	// 主账号Uin

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 描述

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 角色Id

	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
	// 密钥

	Secret *string `json:"Secret,omitempty" name:"Secret"`
	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 策略英文名称

	StrategyEnName *string `json:"StrategyEnName,omitempty" name:"StrategyEnName"`
	// 策略Id

	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
	// 策略详情

	StrategyInfo *string `json:"StrategyInfo,omitempty" name:"StrategyInfo"`
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 策略类型

	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 权重

	Weight *string `json:"Weight,omitempty" name:"Weight"`
	// 录入人

	Writter *string `json:"Writter,omitempty" name:"Writter"`
}

type ServicePermItem struct {
	// 创建时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// ApiAddr

	ApiAddr *string `json:"ApiAddr,omitempty" name:"ApiAddr"`
	// 接口中文名

	ApiZhName *string `json:"ApiZhName,omitempty" name:"ApiZhName"`
	// 鉴权调用的方法名，0: SigAndAuth、1: CheckResource

	AuthFunction *string `json:"AuthFunction,omitempty" name:"AuthFunction"`
	// CWildcardName

	CWildcardName *string `json:"CWildcardName,omitempty" name:"CWildcardName"`
	// 接口英文名

	InterfaceEnName *string `json:"InterfaceEnName,omitempty" name:"InterfaceEnName"`
	// 鉴权粒度，0:接口级别、1:资源级别

	InterfaceLevel *string `json:"InterfaceLevel,omitempty" name:"InterfaceLevel"`
	// 鉴权方式，0:由云API转发鉴权、1:业务自行调用鉴权接口

	IsAuthBusiness *string `json:"IsAuthBusiness,omitempty" name:"IsAuthBusiness"`
	// IsNeedObject

	IsNeedObject *string `json:"IsNeedObject,omitempty" name:"IsNeedObject"`
	// IsSeen

	IsSeen *string `json:"IsSeen,omitempty" name:"IsSeen"`
	// 是否可见

	IsSeenAtGenerator *string `json:"IsSeenAtGenerator,omitempty" name:"IsSeenAtGenerator"`
	// IsSpResource

	IsSpResource *string `json:"IsSpResource,omitempty" name:"IsSpResource"`
	// IsUserSet

	IsUserSet *string `json:"IsUserSet,omitempty" name:"IsUserSet"`
	// PermId

	PermId *string `json:"PermId,omitempty" name:"PermId"`
	// 接口类别

	ReadWriteDetail *string `json:"ReadWriteDetail,omitempty" name:"ReadWriteDetail"`
	// 资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// Weight

	Weight *string `json:"Weight,omitempty" name:"Weight"`
	// 操作人

	Writter *string `json:"Writter,omitempty" name:"Writter"`
	// ServiceName

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// ProductShortCode

	ProductShortCode *string `json:"ProductShortCode,omitempty" name:"ProductShortCode"`
	// ProductShortName

	ProductShortName *string `json:"ProductShortName,omitempty" name:"ProductShortName"`
}

type UpdateServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateServicePermResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateServicePermResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateServicePermResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceRoleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateServiceRoleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateServiceRoleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceListRequest struct {
	*tchttp.BaseRequest

	// 服务名

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 排序字段 默认updateTime

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段desc/asc 默认desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 调用方

	Caller *string `json:"Caller,omitempty" name:"Caller"`
}

func (r *GetServiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceRoleInfoListRequest struct {
	*tchttp.BaseRequest

	// 服务

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 排序字段 默认updateTime

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段desc/asc 默认desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *GetServiceRoleInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceRoleInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateServicePermRequest struct {
	*tchttp.BaseRequest

	// 操作者

	Writer *string `json:"Writer,omitempty" name:"Writer"`
	// 服务

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 操作，update、delete、create

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// API 描述-中

	ApiZhName *string `json:"ApiZhName,omitempty" name:"ApiZhName"`
	// API名

	InterfaceEnName *string `json:"InterfaceEnName,omitempty" name:"InterfaceEnName"`
	// 是否可见 0不可见 1可见

	IsSeenAtGenerator *uint64 `json:"IsSeenAtGenerator,omitempty" name:"IsSeenAtGenerator"`
	// 0云api鉴权 1非云api鉴权 2不鉴权

	IsAuthBusiness *uint64 `json:"IsAuthBusiness,omitempty" name:"IsAuthBusiness"`
	// 0、接口级别 1、资源级别

	InterfaceLevel *uint64 `json:"InterfaceLevel,omitempty" name:"InterfaceLevel"`
	// 0、SigAndAuth 1、CheckResource

	AuthFunction *uint64 `json:"AuthFunction,omitempty" name:"AuthFunction"`
	// 0、新增类型 1、修改类型 2、删除类型 4、查询类型 5、其他类型

	ReadWriteDetail *uint64 `json:"ReadWriteDetail,omitempty" name:"ReadWriteDetail"`
	// 资源前缀 传json字符串 如：['instance','test']

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 接口id

	PermId *uint64 `json:"PermId,omitempty" name:"PermId"`
}

func (r *UpdateServicePermRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateServicePermRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceRole struct {
	// Caction

	Caction *string `json:"Caction,omitempty" name:"Caction"`
	// 有效期

	Duration *string `json:"Duration,omitempty" name:"Duration"`
	// 英文描述

	EnRemark *string `json:"EnRemark,omitempty" name:"EnRemark"`
	// 策略名称

	HiddenPolicyName *string `json:"HiddenPolicyName,omitempty" name:"HiddenPolicyName"`
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// Principal

	Principal *string `json:"Principal,omitempty" name:"Principal"`
	// 描述

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 角色描述

	RoleDesc *string `json:"RoleDesc,omitempty" name:"RoleDesc"`
	// 角色英文描述

	RoleDescEn *string `json:"RoleDescEn,omitempty" name:"RoleDescEn"`
	// 角色名称

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 角色策略

	RolePolicies []*RolePolicy `json:"RolePolicies,omitempty" name:"RolePolicies"`
	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 英文服务类型

	ServiceTypeEn *string `json:"ServiceTypeEn,omitempty" name:"ServiceTypeEn"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// UserPolicyName

	UserPolicyName *string `json:"UserPolicyName,omitempty" name:"UserPolicyName"`
	// 操作人

	Writter *string `json:"Writter,omitempty" name:"Writter"`
}

type UpdatePresetStrategyStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePresetStrategyStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePresetStrategyStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiEnName struct {
	// IsNeedObject

	IsNeedObject *uint64 `json:"IsNeedObject,omitempty" name:"IsNeedObject"`
	// IsSeen

	IsSeen *uint64 `json:"IsSeen,omitempty" name:"IsSeen"`
	// 接口名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type GetStrategyDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Code *uint64 `json:"Code,omitempty" name:"Code"`
		// 消息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 预设策略

		Data *PresetStrategy `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStrategyDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStrategyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePresetStrategyStateRequest struct {
	*tchttp.BaseRequest

	// 操作者

	Writer *string `json:"Writer,omitempty" name:"Writer"`
	// 策略id列表

	StrategyIdList []*uint64 `json:"StrategyIdList,omitempty" name:"StrategyIdList"`
	// 0 、不可见  1、可见

	IsSeen *uint64 `json:"IsSeen,omitempty" name:"IsSeen"`
}

func (r *UpdatePresetStrategyStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePresetStrategyStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceRoleInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Code *uint64 `json:"Code,omitempty" name:"Code"`
		// 消息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 服务角色

		Data []*ServiceRole `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServiceRoleInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceRoleInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePresetStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePresetStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePresetStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
