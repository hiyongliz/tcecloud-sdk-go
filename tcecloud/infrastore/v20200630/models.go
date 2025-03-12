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

package v20200630

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type GenTimeFormatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志时间格式

		TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenTimeFormatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenTimeFormatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProxyTypeInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProxyTypeInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProxyTypeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogRuleGroupTemplesRequest struct {
	*tchttp.BaseRequest

	// 搜索条件

	Params *SearchRuleGroupTemples `json:"Params,omitempty" name:"Params"`
}

func (r *SearchLogRuleGroupTemplesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogRuleGroupTemplesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchNotificationAlertsRequest struct {
	*tchttp.BaseRequest

	// 消息id

	MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
	// 参数对象

	Command *SearchNotificationAlerts `json:"Command,omitempty" name:"Command"`
}

func (r *SearchNotificationAlertsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchNotificationAlertsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddMetaResourceTypeRequest struct {
	*tchttp.BaseRequest

	// 资源组对象

	Command *MetaResourceTypeCreateCommand `json:"Command,omitempty" name:"Command"`
	// 资源归属信息，运营端使用

	Param *ResourceOwnerQueryParam `json:"Param,omitempty" name:"Param"`
}

func (r *AddMetaResourceTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMetaResourceTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchNotificationMsgsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchNotificationMsgsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchNotificationMsgsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeFilter struct {
	// > < >= <= = !=

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 过滤指标key

	FilterKey *string `json:"FilterKey,omitempty" name:"FilterKey"`
	// 过滤指标值

	FilterValue *uint64 `json:"FilterValue,omitempty" name:"FilterValue"`
}

type CreateOrUpdateLogDashboard struct {
	// dashboard Id，根据id是否>0区分create还是update

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 待复制源dashboardId

	SourceDashboardId *int64 `json:"SourceDashboardId,omitempty" name:"SourceDashboardId"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 是否收藏 true 已收藏，false 未收藏

	IsFavorite *string `json:"IsFavorite,omitempty" name:"IsFavorite"`
	// dashboard名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// dashboard数据，json对象

	Data *string `json:"Data,omitempty" name:"Data"`
	// 乐观锁版本, 创建传0

	Version *int64 `json:"Version,omitempty" name:"Version"`
}

type CreateOrUpdateInfraTemplateRequest struct {
	*tchttp.BaseRequest

	// 云哨模板创建或更新接口参数对象

	Command *CreateOrUpdateInfraTemplate `json:"Command,omitempty" name:"Command"`
}

func (r *CreateOrUpdateInfraTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateInfraTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRoutesRequest struct {
	*tchttp.BaseRequest

	// 批量删除，Ids  1,2,3

	Params *BatchDeleteIds `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMetaMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 指标列表

		Data []*MetaMetric `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMetaMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMetaMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardMeta struct {
	// 是否可管理

	CanAdmin *bool `json:"CanAdmin,omitempty" name:"CanAdmin"`
	// 是否可编辑

	CanEdit *bool `json:"CanEdit,omitempty" name:"CanEdit"`
	// 是否可保存

	CanSave *bool `json:"CanSave,omitempty" name:"CanSave"`
	// 是否可加星

	CanStar *bool `json:"CanStar,omitempty" name:"CanStar"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 创建用户

	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`
	// 过期时间

	Expires *string `json:"Expires,omitempty" name:"Expires"`
	// 所在目录ID

	FolderId *int64 `json:"FolderId,omitempty" name:"FolderId"`
	// 所在目录名称

	FolderTitle *string `json:"FolderTitle,omitempty" name:"FolderTitle"`
	// 是否启用ACL

	HasAcl *bool `json:"HasAcl,omitempty" name:"HasAcl"`
	// 更新时间

	Updated *string `json:"Updated,omitempty" name:"Updated"`
	// 更新用户

	UpdatedBy *string `json:"UpdatedBy,omitempty" name:"UpdatedBy"`
	// 版本号

	Version *int64 `json:"Version,omitempty" name:"Version"`
}

type UpdateDashAclCommand struct {
	// 所有要添加的权限列表

	Items []*DashAclItem `json:"Items,omitempty" name:"Items"`
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

type GetNotificationMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNotificationMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNotificationMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDashFolderByUidRequest struct {
	*tchttp.BaseRequest

	// 仪表盘目录唯一ID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 仪表盘目录更新内容

	Command *UpdateDashFolderCommand `json:"Command,omitempty" name:"Command"`
}

func (r *UpdateDashFolderByUidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDashFolderByUidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSilencesRequest struct {
	*tchttp.BaseRequest

	// 批量删除ids，"Ids":"1,2,3"

	Params *BatchDeleteIds `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteSilencesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSilencesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogDeliveryTargetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogDeliveryTargetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogDeliveryTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRelatedAlertNamesRequest struct {
	*tchttp.BaseRequest

	// 参数

	Params *GetRelatedAlertNames `json:"Params,omitempty" name:"Params"`
}

func (r *GetRelatedAlertNamesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRelatedAlertNamesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchMetaValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// metavalue

		MetaNames []*string `json:"MetaNames,omitempty" name:"MetaNames"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchMetaValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchMetaValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLogTransformTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateLogTransformTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogTransformTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDashFolderByUidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDashFolderByUidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashFolderByUidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogConfigResp struct {
	// CRD文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// yaml文件内容

	Data *string `json:"Data,omitempty" name:"Data"`
}

type SearchRuleGroups struct {
	// 策略id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// true or false

	Enable *string `json:"Enable,omitempty" name:"Enable"`
	// page offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// page limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 升序字段: field1[,field2 ...]

	Asc *string `json:"Asc,omitempty" name:"Asc"`
	// 降序字段: field1[,field2 ...]

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 归属产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 创建人

	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`
	// 是否启用正则匹配，可选true、false，默认false

	IsRegex *string `json:"IsRegex,omitempty" name:"IsRegex"`
	// 策略类型 1 自定义事件策略 0 告警策略

	GroupType *int64 `json:"GroupType,omitempty" name:"GroupType"`
	// alert record

	Type *string `json:"Type,omitempty" name:"Type"`
	// tcs type

	TcsType *string `json:"TcsType,omitempty" name:"TcsType"`
	// 扩展字段json

	Extension *string `json:"Extension,omitempty" name:"Extension"`
}

type DeleteMetaResourceTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMetaResourceTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMetaResourceTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateResourceObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateResourceObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateResourceObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRuleGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRuleGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAlertStatsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAlertStatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddMetaMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 指标信息

		Data *MetaMetric `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddMetaMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMetaMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashboardTagsRequest struct {
	*tchttp.BaseRequest
}

func (r *GetDashboardTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSilenceRequest struct {
	*tchttp.BaseRequest

	// 策略 id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 策略定义

	Command *UpdateSilence `json:"Command,omitempty" name:"Command"`
}

func (r *UpdateSilenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSilenceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VisualizeLogsRequest struct {
	*tchttp.BaseRequest

	// LogQL查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 分桶参数

	Buckets []*AnalyzeBucket `json:"Buckets,omitempty" name:"Buckets"`
	// 指标参数

	Metrics []*AnalyzeMetric `json:"Metrics,omitempty" name:"Metrics"`
	// 过滤参数

	Filters []*AnalyzeFilter `json:"Filters,omitempty" name:"Filters"`
	// 样例: 1m30s    , 评估时间间隔，选时间分组时，必填，用来控制评估时间

	Interval *string `json:"Interval,omitempty" name:"Interval"`
	// 有时间分组时，用来控制展示间隔， 单位为s

	Step *int64 `json:"Step,omitempty" name:"Step"`
	// 起始时间 纳秒时间戳 或 RFC3339Nano

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 截止时间 纳秒时间戳 或 RFC3339Nano

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *VisualizeLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VisualizeLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateLogPanel struct {
	// 仪表盘id

	DashboardId *int64 `json:"DashboardId,omitempty" name:"DashboardId"`
	// 面板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	//  面板数据 json对象

	Data *string `json:"Data,omitempty" name:"Data"`
	// 乐观锁版本号，大于0 为更新 默认为0 创建

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 面板ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type ApplyInfraTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyInfraTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyInfraTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Bucket struct {
	// 分桶key的格式化结果

	KeyAsString *string `json:"KeyAsString,omitempty" name:"KeyAsString"`
	// 分桶key 纳秒时间戳

	Key *uint64 `json:"Key,omitempty" name:"Key"`
	// 桶内数据计数

	DocCount *uint64 `json:"DocCount,omitempty" name:"DocCount"`
}

type CreateOrUpdateLogPanelRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Command *CreateOrUpdateLogPanel `json:"Command,omitempty" name:"Command"`
}

func (r *CreateOrUpdateLogPanelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateLogPanelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRuleGroupRequest struct {
	*tchttp.BaseRequest

	// 策略 id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetRuleGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRuleGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogMappingRequest struct {
	*tchttp.BaseRequest

	// 查询LogQL(只支持日志查询LogQL，不能包含count_over_time等指标计算)

	Query *string `json:"Query,omitempty" name:"Query"`
	// 是否只展示原始字段

	OnlyOriginalField *bool `json:"OnlyOriginalField,omitempty" name:"OnlyOriginalField"`
}

func (r *LogMappingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LogMappingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProxyAllDataSourceRequestByNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProxyAllDataSourceRequestByNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProxyAllDataSourceRequestByNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateResourceObjectRequest struct {
	*tchttp.BaseRequest

	// 资源归属信息

	Params *ResourceOwnerQueryParam `json:"Params,omitempty" name:"Params"`
	// 服务树节点标签

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 更新信息

	Command *UpdateResourceCommand `json:"Command,omitempty" name:"Command"`
}

func (r *UpdateResourceObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateResourceObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Meta struct {
	// meta key

	MetaKey *string `json:"MetaKey,omitempty" name:"MetaKey"`
	// meta value

	MetaValue *string `json:"MetaValue,omitempty" name:"MetaValue"`
	// 子 meta数组

	SubMetas []*SubMeta `json:"SubMetas,omitempty" name:"SubMetas"`
	// meta name

	MetaName *string `json:"MetaName,omitempty" name:"MetaName"`
	// 配置数量

	ConfigCount *uint64 `json:"ConfigCount,omitempty" name:"ConfigCount"`
}

type MetaMetricUpdateCommand struct {
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 说明

	Help *string `json:"Help,omitempty" name:"Help"`
}

type CreateOrUpdateLogRuleGroupTempleRequest struct {
	*tchttp.BaseRequest

	// 策略模板定义

	Command *CreateOrUpdateRuleGroupTemple `json:"Command,omitempty" name:"Command"`
}

func (r *CreateOrUpdateLogRuleGroupTempleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateLogRuleGroupTempleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateRuleGroup struct {
	// 策略名称，同一个租户不可重复

	Name *string `json:"Name,omitempty" name:"Name"`
	// 评估间隔，建议1m

	Interval *string `json:"Interval,omitempty" name:"Interval"`
	// 类型、产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 告警对象Labels

	Objects []*Matcher `json:"Objects,omitempty" name:"Objects"`
	// 规则数组

	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 是否启用，true启用，false暂不启用

	Enable *string `json:"Enable,omitempty" name:"Enable"`
	// 乐观锁版本, 创建传0

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 策略Id，根据id是否>0区分create还是update

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 更新时间

	Updated *string `json:"Updated,omitempty" name:"Updated"`
	// 策略类型，1为自定义事件生成策略，默认为0 ，告警策略

	GroupType *int64 `json:"GroupType,omitempty" name:"GroupType"`
	// 自定义事件名称，GroupType为1时必填

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 策略类型：record或alert

	Type *string `json:"Type,omitempty" name:"Type"`
	// 关联模板ID

	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
	// labels

	Labels []*LabelPair `json:"Labels,omitempty" name:"Labels"`
	// 关联ID（已废弃，后续跟后端一起删除）

	RefId *int64 `json:"RefId,omitempty" name:"RefId"`
	// 扩展字段json

	Extension *string `json:"Extension,omitempty" name:"Extension"`
}

type UpdateRoute struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品名称，支持多个

	Products []*string `json:"Products,omitempty" name:"Products"`
	// 告警等级，支持多个和all

	Severities []*string `json:"Severities,omitempty" name:"Severities"`
	// 是否发送告警

	SendFiring *bool `json:"SendFiring,omitempty" name:"SendFiring"`
	// 是否发送恢复

	SendResolved *bool `json:"SendResolved,omitempty" name:"SendResolved"`
	// 发送通道，支持微信、短信、企业微信、邮件和callback

	Channels []*LabelPair `json:"Channels,omitempty" name:"Channels"`
	// 用户id

	Users []*LabelPair `json:"Users,omitempty" name:"Users"`
	// 用户组id

	Groups []*LabelPair `json:"Groups,omitempty" name:"Groups"`
	// 乐观锁

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 静默时间，也即重发间隔

	RepeatInterval *string `json:"RepeatInterval,omitempty" name:"RepeatInterval"`
	// 策略，支持多个和all

	Strategies []*string `json:"Strategies,omitempty" name:"Strategies"`
	// 0,1->product+severity

	Flag *int64 `json:"Flag,omitempty" name:"Flag"`
}

type DeleteRuleGroupTemplesRequest struct {
	*tchttp.BaseRequest

	// 批量删除ids，"Ids":"1,2,3"

	Params *BatchDeleteIds `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteRuleGroupTemplesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRuleGroupTemplesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDataBaradMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监控指标

		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
		// 数据点起始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 数据点结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 数据统计周期

		Period *uint64 `json:"Period,omitempty" name:"Period"`
		// 监控数据列表

		DataPionts []*PointsObject `json:"DataPionts,omitempty" name:"DataPionts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDataBaradMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDataBaradMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogRuleGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchLogRuleGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogRuleGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMetaMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 指标信息

		Data *MetaMetric `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMetaMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMetaMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfirmAlertResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConfirmAlertResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfirmAlertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VisualizeLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VisualizeLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VisualizeLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardVersion struct {
	// 版本ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 仪表盘ID

	DashboardId *int64 `json:"DashboardId,omitempty" name:"DashboardId"`
	// 父版本号

	ParentVersion *int64 `json:"ParentVersion,omitempty" name:"ParentVersion"`
	// 恢复来源版本

	RestoredFrom *int64 `json:"RestoredFrom,omitempty" name:"RestoredFrom"`
	// 版本

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 创建人

	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`
	// 更新消息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type LabelPair struct {
	// Label Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Label Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateOrUpdateInfraTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrUpdateInfraTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateInfraTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogDeliveryTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogDeliveryTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogDeliveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogPanelRequest struct {
	*tchttp.BaseRequest

	// 查询参数

	Params *DescribeLogPanel `json:"Params,omitempty" name:"Params"`
}

func (r *DescribeLogPanelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogPanelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogConfigRequest struct {
	*tchttp.BaseRequest

	// meta信息

	Metas []*Meta `json:"Metas,omitempty" name:"Metas"`
	// 日志用途标识

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志路径，支持统配

	Path *string `json:"Path,omitempty" name:"Path"`
	// 当path包含通配符时，该字段用于排除不必要的日志文件

	PathExcluded *string `json:"PathExcluded,omitempty" name:"PathExcluded"`
	// 是否启用多行匹配 默认不启用

	MultipleLine *bool `json:"MultipleLine,omitempty" name:"MultipleLine"`
	// 指定系统已有的parse

	ParserRef *string `json:"ParserRef,omitempty" name:"ParserRef"`
	// 用于手工指定配置，可用配置有json、regex等

	ParserFormat *string `json:"ParserFormat,omitempty" name:"ParserFormat"`
	// time字段名称

	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`
	// time格式

	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`
	// 提取正则

	Regex *string `json:"Regex,omitempty" name:"Regex"`
	// 行首正则

	LineHeaderRegex *string `json:"LineHeaderRegex,omitempty" name:"LineHeaderRegex"`
	// 当ParserFormat为LTSV时需要指定

	Type *string `json:"Type,omitempty" name:"Type"`
	// 用于手工指定配置，主要用于对特定定段再次进行解码

	Decoders []*string `json:"Decoders,omitempty" name:"Decoders"`
	// 是否提取字段

	FieldExtraction *bool `json:"FieldExtraction,omitempty" name:"FieldExtraction"`
	// 是否使用系统时间

	UseSystemTime *bool `json:"UseSystemTime,omitempty" name:"UseSystemTime"`
	// 日志样例

	Example *string `json:"Example,omitempty" name:"Example"`
	// 字段信息，json  样例:[{"Field":"cost","Type":"int"},{"Field":"log_level","Type":"string"},{"Field":"old_price","Type":"double"}]

	FieldMapping *string `json:"FieldMapping,omitempty" name:"FieldMapping"`
	// 自定义标签

	CustomLabels []*CustomLabelItem `json:"CustomLabels,omitempty" name:"CustomLabels"`
}

func (r *CreateLogConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportLogAlertRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportLogAlertRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportLogAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashboardContentByUidRequest struct {
	*tchttp.BaseRequest

	// 仪表唯一ID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
}

func (r *GetDashboardContentByUidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardContentByUidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRuleGroupTemplesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchRuleGroupTemplesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRuleGroupTemplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRetention struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type UpdateRouteRequest struct {
	*tchttp.BaseRequest

	// 路由ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 路由信息struct

	Command *UpdateRoute `json:"Command,omitempty" name:"Command"`
}

func (r *UpdateRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateResourceCommand struct {
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type DeleteLogRuleGroupsRequest struct {
	*tchttp.BaseRequest

	// 批量删除ids，"Ids":"1,2,3"

	Params *BatchDeleteIds `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteLogRuleGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogRuleGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddLogDeliveryTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddLogDeliveryTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLogDeliveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMetaResourceTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMetaResourceTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMetaResourceTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaEventQueryParam struct {
	// AppID

	Organization *string `json:"Organization,omitempty" name:"Organization"`
}

type QueryRecord struct {
	// 创建时间 s级时间戳

	CreatedAt *uint64 `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 记录id

	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`
	// 用户id

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 筛选过滤查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 查询名称

	QueryName *string `json:"QueryName,omitempty" name:"QueryName"`
}

type SearchRuleGroupTemples struct {
	// 策略模板id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 策略模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 归属产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// page offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// page limit, default 1000

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 升序字段，file1,field2...

	Asc *string `json:"Asc,omitempty" name:"Asc"`
	// 降序字段，file1,field2...

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 对象类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 是否启用正则匹配，可选true、false，默认false

	IsRegex *string `json:"IsRegex,omitempty" name:"IsRegex"`
}

type ExportLogConfigCRResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出CRD文件名及文件内容

		LogConfigs []*LogConfigResp `json:"LogConfigs,omitempty" name:"LogConfigs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportLogConfigCRResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportLogConfigCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeSort struct {
	// 排序字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// 默认desc 倒序，asc正序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 字段值是否数值型

	IsNumber *bool `json:"IsNumber,omitempty" name:"IsNumber"`
}

type SilenceGroup struct {
	// 策略ID和Name

	Group *LabelPair `json:"Group,omitempty" name:"Group"`
	// 规则Id和Name数组

	Rules []*LabelPair `json:"Rules,omitempty" name:"Rules"`
	// 分类

	Class *string `json:"Class,omitempty" name:"Class"`
}

type AddResourceObjectRequest struct {
	*tchttp.BaseRequest

	// 资源归属信息

	Params *ResourceOwnerQueryParam `json:"Params,omitempty" name:"Params"`
	// 资源对象信息

	Command *CreateResourceCommand `json:"Command,omitempty" name:"Command"`
}

func (r *AddResourceObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddResourceObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDashFolderByUidRequest struct {
	*tchttp.BaseRequest

	// 仪表盘目录唯一ID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
}

func (r *DeleteDashFolderByUidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashFolderByUidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogTransformTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogTransformTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogTransformTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogMappingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志字段映射

		Mapping []*LogField `json:"Mapping,omitempty" name:"Mapping"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LogMappingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LogMappingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDashboardCommand struct {
	// 仪表盘内容，JSON字符串

	Dashboard *string `json:"Dashboard,omitempty" name:"Dashboard"`
	// 仪表盘变更消息说明

	Message *string `json:"Message,omitempty" name:"Message"`
	// 是否覆盖

	Overwrite *bool `json:"Overwrite,omitempty" name:"Overwrite"`
	// 目录ID

	FolderId *int64 `json:"FolderId,omitempty" name:"FolderId"`
}

type SearchRoutesByLabelSet struct {
	// 产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// tcs_type 产品模块

	Type *string `json:"Type,omitempty" name:"Type"`
	// 告警等级

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 告警策略名

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 告警策略类型 log metric

	AlertClass *string `json:"AlertClass,omitempty" name:"AlertClass"`
}

type DeleteLogAlertRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogAlertRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashFoldersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仪表盘目录检索结果

		Data []*DashFolderQueryResult `json:"Data,omitempty" name:"Data"`
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDashFoldersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashFoldersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceObject struct {
	// 服务树节点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 父节点ID

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 类型标签

	LabelName *string `json:"LabelName,omitempty" name:"LabelName"`
	// 服务树节点选择标签，格式为{a="b",c="d"}

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 资源归属，取用户的AppId

	Organization *string `json:"Organization,omitempty" name:"Organization"`
}

type DescribeLogDashboardRequest struct {
	*tchttp.BaseRequest

	// 查询参数

	Params *DescribeLogDashboard `json:"Params,omitempty" name:"Params"`
}

func (r *DescribeLogDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogPanels struct {
	// 待删除ids，逗号分隔

	Ids *string `json:"Ids,omitempty" name:"Ids"`
}

type MetaMetricCreateCommand struct {
	// 指标名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 指标Label列表

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 说明

	Help *string `json:"Help,omitempty" name:"Help"`
	// 类型，如counter，gauge，summary，histogram

	Type *string `json:"Type,omitempty" name:"Type"`
}

type SearchDashboardsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchDashboardsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchDashboardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDashFolderCommand struct {
	// Unique Id

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 仪表盘目录名称

	Title *string `json:"Title,omitempty" name:"Title"`
}

type DashFolderQueryParam struct {
	// 查询字符串

	Query *string `json:"Query,omitempty" name:"Query"`
	// 页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 第几页

	Page *int64 `json:"Page,omitempty" name:"Page"`
}

type SearchRoutesByLabelSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchRoutesByLabelSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRoutesByLabelSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaResourceTypeMetric struct {
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标维度

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
	// 指标别名，用于界面展示

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 指标单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 指标Hash

	Fingerprint *uint64 `json:"Fingerprint,omitempty" name:"Fingerprint"`
	// 资源对象类型ID

	ResourceTypeId *int64 `json:"ResourceTypeId,omitempty" name:"ResourceTypeId"`
}

type ListResourceTypesRequest struct {
	*tchttp.BaseRequest
}

func (r *ListResourceTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLogConfigStateRequest struct {
	*tchttp.BaseRequest

	// 待修改配置ids

	ConfigIds []*string `json:"ConfigIds,omitempty" name:"ConfigIds"`
	// 状态，0 禁用 1启用

	State *uint64 `json:"State,omitempty" name:"State"`
}

func (r *UpdateLogConfigStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogConfigStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Aggs struct {
	// 自定义聚合命名

	AggName *string `json:"AggName,omitempty" name:"AggName"`
	// 聚合行为

	AggOperator *AggOperator `json:"AggOperator,omitempty" name:"AggOperator"`
}

type ApplyInfraTemplates struct {
	// 待应用云哨模板对象数组

	Templates []*InfraTemplate `json:"Templates,omitempty" name:"Templates"`
}

type LogDeliveryTaskResponseItem struct {
	// 日志投递任务id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 投递任务关联日志

	Metas []*Meta `json:"Metas,omitempty" name:"Metas"`
	// 接收端类型

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 接收端id

	TargetId *uint64 `json:"TargetId,omitempty" name:"TargetId"`
	// 接收端名称

	TargetName *string `json:"TargetName,omitempty" name:"TargetName"`
	// 投递数据格式

	Format *string `json:"Format,omitempty" name:"Format"`
	// 投递任务状态，1:启用，0:停用

	State *uint64 `json:"State,omitempty" name:"State"`
	// 最近一次更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type SearchNotificationMsgsRequest struct {
	*tchttp.BaseRequest

	// 云哨消息查询参数

	Command *SearchNotificationMsgs `json:"Command,omitempty" name:"Command"`
}

func (r *SearchNotificationMsgsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchNotificationMsgsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchDeleteIds struct {
	// 逗号分隔，"Ids": "3,4,5"

	Ids *string `json:"Ids,omitempty" name:"Ids"`
	// 逗号分隔，”names“:"z,z"

	Names *string `json:"Names,omitempty" name:"Names"`
}

type CreateOrUpdateRetentionRequest struct {
	*tchttp.BaseRequest

	// Retention定义

	Command *CreateOrUpdateRetention `json:"Command,omitempty" name:"Command"`
}

func (r *CreateOrUpdateRetentionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateRetentionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardByUidRequest struct {
	*tchttp.BaseRequest

	// 仪表盘唯一ID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
}

func (r *DeleteDashboardByUidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardByUidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLogDeliveryTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志投递任务列表

		Task []*LogDeliveryTaskResponseItem `json:"Task,omitempty" name:"Task"`
		// 日志投递任务总量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListLogDeliveryTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLogDeliveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSilencesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchSilencesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSilencesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Rule struct {
	// 策略中的规则名

	Name *string `json:"Name,omitempty" name:"Name"`
	// record or alert

	Type *string `json:"Type,omitempty" name:"Type"`
	// 持续时间

	For *string `json:"For,omitempty" name:"For"`
	// 规则表达式

	Expr *string `json:"Expr,omitempty" name:"Expr"`
	// 自定义label

	Labels []*LabelPair `json:"Labels,omitempty" name:"Labels"`
	// 自定义告警描述

	Annotation []*LabelPair `json:"Annotation,omitempty" name:"Annotation"`
	// 规则Id，前端忽略

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 策略id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 更新时间

	Updated *string `json:"Updated,omitempty" name:"Updated"`
	// 以产品场景进行分类，根据不同场景返回进行不同的解析

	ExprArray []*Expr `json:"ExprArray,omitempty" name:"ExprArray"`
	// 规则描述，便于用户理解

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// > >= <= < = !=

	ExprAct *string `json:"ExprAct,omitempty" name:"ExprAct"`
	// 阈值，会有一个默认值，但允许用户改写

	ExprVal *string `json:"ExprVal,omitempty" name:"ExprVal"`
	// 告警等级

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 规则拆解后的值

	ExprArgs *LabelPair `json:"ExprArgs,omitempty" name:"ExprArgs"`
	// 云哨默认模板告警规则可编辑表达式

	EditableExpr *EditableExpr `json:"EditableExpr,omitempty" name:"EditableExpr"`
}

type UpdateSilence struct {
	// 屏蔽单名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 级别

	Level *string `json:"Level,omitempty" name:"Level"`
	// 屏蔽开始时间

	StartsAt *string `json:"StartsAt,omitempty" name:"StartsAt"`
	// 屏蔽结束时间

	EndsAt *string `json:"EndsAt,omitempty" name:"EndsAt"`
	// 产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 对象

	Objects []*Matcher `json:"Objects,omitempty" name:"Objects"`
	// 策略和规则信息

	Labels []*SilenceGroup `json:"Labels,omitempty" name:"Labels"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 乐观锁版本

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 自定义变更人员，默认是当前用户，多个用户间逗号分隔

	Updater *string `json:"Updater,omitempty" name:"Updater"`
	// 扩展字段json

	Extension *string `json:"Extension,omitempty" name:"Extension"`
}

type GetRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 服务树资源对象

		Data []*ResourceObject `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListResourceObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogConfigRequest struct {
	*tchttp.BaseRequest

	// 配置id

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 日志路径，支持通配

	Path *string `json:"Path,omitempty" name:"Path"`
	// 当path包含通配符时，该字段用于排除不必要的日志文件

	PathExcluded *string `json:"PathExcluded,omitempty" name:"PathExcluded"`
	// 是否启用多行匹配

	MultipleLine *bool `json:"MultipleLine,omitempty" name:"MultipleLine"`
	// 指定系统已有的parse

	ParserRef *string `json:"ParserRef,omitempty" name:"ParserRef"`
	// 用于手工指定配置，可用配置有json、regex等

	ParserFormat *string `json:"ParserFormat,omitempty" name:"ParserFormat"`
	// time字段名称

	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`
	// time格式

	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`
	// 提取正则

	Regex *string `json:"Regex,omitempty" name:"Regex"`
	// 行首正则

	LineHeaderRegex *string `json:"LineHeaderRegex,omitempty" name:"LineHeaderRegex"`
	// 当ParserFormat为LTSV时需要指定

	Type *string `json:"Type,omitempty" name:"Type"`
	// 用于手工指定配置，主要用于对特定定段再次进行解码

	Decoders []*string `json:"Decoders,omitempty" name:"Decoders"`
	// 是否提取字段

	FieldExtraction *bool `json:"FieldExtraction,omitempty" name:"FieldExtraction"`
	// 是否使用系统时间

	UseSystemTime *bool `json:"UseSystemTime,omitempty" name:"UseSystemTime"`
	// 日志名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志样例

	Example *string `json:"Example,omitempty" name:"Example"`
	// 字段信息，json  样例:[{"Field":"cost","Type":"int"},{"Field":"log_level","Type":"string"},{"Field":"old_price","Type":"double"}]

	FieldMapping *string `json:"FieldMapping,omitempty" name:"FieldMapping"`
	// 自定义标签

	CustomLabels []*CustomLabelItem `json:"CustomLabels,omitempty" name:"CustomLabels"`
}

func (r *ModifyLogConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AggOperator struct {
	// 日期直方聚合

	DateHistogram *DateHistogram `json:"DateHistogram,omitempty" name:"DateHistogram"`
}

type LogTransformTaskParamItem struct {
	// 任务id

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 处理任务关联日志

	Metas []*Meta `json:"Metas,omitempty" name:"Metas"`
	// 启用/停用投递任务

	Enabled *uint64 `json:"Enabled,omitempty" name:"Enabled"`
	// 处理函数

	Transforms []*LogTransformFunction `json:"Transforms,omitempty" name:"Transforms"`
}

type CreateQueryRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询记录

		QueryRecord *QueryRecord `json:"QueryRecord,omitempty" name:"QueryRecord"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateQueryRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateQueryRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogAlertRuleRequest struct {
	*tchttp.BaseRequest

	// 需要删除的报警ID

	RuleIDs []*uint64 `json:"RuleIDs,omitempty" name:"RuleIDs"`
}

func (r *DeleteLogAlertRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogAlertRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogTransformTaskRequest struct {
	*tchttp.BaseRequest

	// 删除任务id列表

	TaskIDs []*uint64 `json:"TaskIDs,omitempty" name:"TaskIDs"`
}

func (r *DeleteLogTransformTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogTransformTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRuleGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchRuleGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRuleGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLogDeliveryTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateLogDeliveryTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogDeliveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogConfigParam struct {
	// 日志配置Id

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

type DeleteMetaMetricRequest struct {
	*tchttp.BaseRequest

	// 资源归属信息

	Params *ResourceOwnerQueryParam `json:"Params,omitempty" name:"Params"`
	// 资源对象类型名称

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 指标指纹ID

	Fingerprint *string `json:"Fingerprint,omitempty" name:"Fingerprint"`
}

func (r *DeleteMetaMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMetaMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogDashboardsRequest struct {
	*tchttp.BaseRequest

	// 查询参数

	Params *SearchLogDashboards `json:"Params,omitempty" name:"Params"`
}

func (r *SearchLogDashboardsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogDashboardsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Expr struct {
	// metric name, label使用RuleGroup.Objects

	Name *string `json:"Name,omitempty" name:"Name"`
	// expr value

	Val *string `json:"Val,omitempty" name:"Val"`
	// == != > < >= <=  [ metric op val ]

	Op *string `json:"Op,omitempty" name:"Op"`
	// 聚合函数，count_over_time, rate

	AggrFunc *string `json:"AggrFunc,omitempty" name:"AggrFunc"`
	// sum,avg     [ func(metric) op val ]

	Func *string `json:"Func,omitempty" name:"Func"`
	// [1d,1w,1m]

	Range *string `json:"Range,omitempty" name:"Range"`
	// duraton 1d,1m,1w

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// and,or,*,/

	NextRel *string `json:"NextRel,omitempty" name:"NextRel"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// up,down

	Thresh *string `json:"Thresh,omitempty" name:"Thresh"`
	// 场景类型

	SceneType *string `json:"SceneType,omitempty" name:"SceneType"`
}

type GetDashboardVersionsByIdRequest struct {
	*tchttp.BaseRequest

	// 仪表盘ID

	DashId *string `json:"DashId,omitempty" name:"DashId"`
}

func (r *GetDashboardVersionsByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardVersionsByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLogTransformTaskEnabledRequest struct {
	*tchttp.BaseRequest

	// 任务id列表

	TaskIDs []*uint64 `json:"TaskIDs,omitempty" name:"TaskIDs"`
	// 开启状态

	Enabled *uint64 `json:"Enabled,omitempty" name:"Enabled"`
}

func (r *UpdateLogTransformTaskEnabledRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogTransformTaskEnabledRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleGroupsRequest struct {
	*tchttp.BaseRequest

	// 批量删除ids，"Ids":"1,2,3"

	Params *BatchDeleteIds `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteRuleGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRuleGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashboardVersionContentByIdRequest struct {
	*tchttp.BaseRequest

	// 仪表盘ID

	DashId *string `json:"DashId,omitempty" name:"DashId"`
	// 版本号

	DashVersion *string `json:"DashVersion,omitempty" name:"DashVersion"`
}

func (r *GetDashboardVersionContentByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardVersionContentByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchInfraTemplatesRefsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchInfraTemplatesRefsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchInfraTemplatesRefsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRetentionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchRetentionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRetentionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportLogAlertRuleRequest struct {
	*tchttp.BaseRequest

	// 日志告警策略ID

	ID []*uint64 `json:"ID,omitempty" name:"ID"`
}

func (r *ExportLogAlertRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportLogAlertRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAlertTrend struct {
	// 开始时间

	StartsAt *string `json:"StartsAt,omitempty" name:"StartsAt"`
	// 结束时间

	EndsAt *string `json:"EndsAt,omitempty" name:"EndsAt"`
}

type CreateOrUpdateRuleGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrUpdateRuleGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogRuleGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogRuleGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogRuleGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteQueryRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录id

		RecordId *string `json:"RecordId,omitempty" name:"RecordId"`
		// 删除是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteQueryRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteQueryRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置id

		ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSilencesRequest struct {
	*tchttp.BaseRequest

	// 搜索条件

	Params *SearchSilences `json:"Params,omitempty" name:"Params"`
}

func (r *SearchSilencesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSilencesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AggResults struct {
	// 多桶型聚合结果

	Buckets []*Bucket `json:"Buckets,omitempty" name:"Buckets"`
}

type AddMetaMetricRequest struct {
	*tchttp.BaseRequest

	// 资源归属信息，运营端使用

	Params *ResourceOwnerQueryParam `json:"Params,omitempty" name:"Params"`
	// 指标信息

	Command *MetaMetricCreateCommand `json:"Command,omitempty" name:"Command"`
}

func (r *AddMetaMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMetaMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSilencesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSilencesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSilencesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenLineHeaderRegexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 行首正则

		LineHeaderRegex *string `json:"LineHeaderRegex,omitempty" name:"LineHeaderRegex"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenLineHeaderRegexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenLineHeaderRegexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenRegexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自动生成的正则表达式

		Regex *string `json:"Regex,omitempty" name:"Regex"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenRegexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenRegexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashAclContent struct {
	// 仪表盘ID

	DashboardId *int64 `json:"DashboardId,omitempty" name:"DashboardId"`
	// 仪表盘目录ID

	FolderId *int64 `json:"FolderId,omitempty" name:"FolderId"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 更新时间

	Updated *string `json:"Updated,omitempty" name:"Updated"`
	// 用户ID

	UserId *int64 `json:"UserId,omitempty" name:"UserId"`
	// 用户名称

	UserLogin *string `json:"UserLogin,omitempty" name:"UserLogin"`
	// 用户邮件地址

	UserEmail *string `json:"UserEmail,omitempty" name:"UserEmail"`
	// 角色

	Role *string `json:"Role,omitempty" name:"Role"`
	// 权限

	Permission *string `json:"Permission,omitempty" name:"Permission"`
	// 权限名称

	PermissionName *string `json:"PermissionName,omitempty" name:"PermissionName"`
	// 仪表盘或仪表盘目录Unique ID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 仪表盘或仪表盘目录名称

	Title *string `json:"Title,omitempty" name:"Title"`
	// 是否继续来的权限

	Inherited *bool `json:"Inherited,omitempty" name:"Inherited"`
}

type UpdateLogAlertRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateLogAlertRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateInfraTemplate struct {
	// 模板所属产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 模板所属模块

	Module *string `json:"Module,omitempty" name:"Module"`
	// 模板所属业务（log/event/metric）

	BizType *string `json:"BizType,omitempty" name:"BizType"`
	// 模板类型(log_rule_group/metric_rule_group)

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
	// 模板来源(preset/preset_cr/custom)

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 模板内容（json）

	Content *string `json:"Content,omitempty" name:"Content"`
	// 模板版本

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 模板ID（更新时必传）

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 扩展字段json

	Extension *string `json:"Extension,omitempty" name:"Extension"`
	// 模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type GetRelatedAlertNames struct {
	// 待查询事件名列表

	EventNames []*string `json:"EventNames,omitempty" name:"EventNames"`
}

type MetaMetricQueryParam struct {
	// 租户ID，对应AppId，通常不需要传，默认取当前登录用户的AppId

	Organization *string `json:"Organization,omitempty" name:"Organization"`
}

type DummyContent struct {
	// 标题名称

	Title *string `json:"Title,omitempty" name:"Title"`
}

type ProxyAllDataSourceRequestByNameRequest struct {
	*tchttp.BaseRequest

	// 额外路径

	AdditionalPath *string `json:"AdditionalPath,omitempty" name:"AdditionalPath"`
	// 代理方法，当前只支持GET

	Method *string `json:"Method,omitempty" name:"Method"`
	// 数据类型，当前支持metrics和logs

	Name *string `json:"Name,omitempty" name:"Name"`
	// 数据查询参数

	Params *DataQueryParams `json:"Params,omitempty" name:"Params"`
}

func (r *ProxyAllDataSourceRequestByNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProxyAllDataSourceRequestByNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenTimeFormatRequest struct {
	*tchttp.BaseRequest

	// 日志时间值

	LogTime *string `json:"LogTime,omitempty" name:"LogTime"`
}

func (r *GenTimeFormatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenTimeFormatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMetaRequest struct {
	*tchttp.BaseRequest

	// Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤类型，默认0 不过滤， 1 只展示已接入日志配置的meta

	FilterType *uint64 `json:"FilterType,omitempty" name:"FilterType"`
	// 模糊查询MetaName

	MetaName *string `json:"MetaName,omitempty" name:"MetaName"`
}

func (r *GetMetaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMetaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchInfraTemplatesRefsRequest struct {
	*tchttp.BaseRequest

	// 云哨模板关联实例查询接口参数

	Command *SearchInfraTemplatesRefs `json:"Command,omitempty" name:"Command"`
}

func (r *SearchInfraTemplatesRefsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchInfraTemplatesRefsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSilence struct {
	// 规则名，工单名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 屏蔽级别

	Level *string `json:"Level,omitempty" name:"Level"`
	// 屏蔽开始时间

	StartsAt *string `json:"StartsAt,omitempty" name:"StartsAt"`
	// 屏蔽结束时间

	EndsAt *string `json:"EndsAt,omitempty" name:"EndsAt"`
	// 产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 对象信息

	Objects []*Matcher `json:"Objects,omitempty" name:"Objects"`
	// 策略和规则信息

	Labels []*SilenceGroup `json:"Labels,omitempty" name:"Labels"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 自定义变更人员，默认是当前用户，多个用户间逗号分隔

	Updater *string `json:"Updater,omitempty" name:"Updater"`
	// 扩展字段json

	Extension *string `json:"Extension,omitempty" name:"Extension"`
}

type LogTransformFunction struct {
	// 处理函数

	Function *string `json:"Function,omitempty" name:"Function"`
	// 处理字段名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 对象名称

	Object *string `json:"Object,omitempty" name:"Object"`
	// 函数参数

	Args *string `json:"Args,omitempty" name:"Args"`
	// 函数id

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 任务id

	TaskID *uint64 `json:"TaskID,omitempty" name:"TaskID"`
}

type CreateOrUpdateLogDashboardRequest struct {
	*tchttp.BaseRequest

	// LogDashboard

	Command *CreateOrUpdateLogDashboard `json:"Command,omitempty" name:"Command"`
}

func (r *CreateOrUpdateLogDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateLogDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateLogRuleGroupTempleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrUpdateLogRuleGroupTempleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateLogRuleGroupTempleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogConfigRequest struct {
	*tchttp.BaseRequest

	// 配置id

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeleteLogConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashboardVersionContentByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDashboardVersionContentByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardVersionContentByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogsRequest struct {
	*tchttp.BaseRequest

	// Meta信息

	Metas []*Meta `json:"Metas,omitempty" name:"Metas"`
	// 日志名，可以选多个

	Name []*string `json:"Name,omitempty" name:"Name"`
	// 时间范围起始，纳秒时间戳，默认当前时间(总数精确到秒级，如果需要准确查询时间范围内总数，时间起止改为 秒级时间戳*10^9)

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 时间范围结束，纳秒时间戳，默认一小时前总数精确到秒级，如果需要准确查询时间范围内总数，时间起止改为 秒级时间戳*10^9)

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 返回记录条数，默认值：500，最大值：1000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 聚集查询

	Aggs *Aggs `json:"Aggs,omitempty" name:"Aggs"`
	// 排序

	Sort []*SortInfo `json:"Sort,omitempty" name:"Sort"`
	// 查询方式，0: LogQL

	QueryType *uint64 `json:"QueryType,omitempty" name:"QueryType"`
	// 数据类型  默认 0：日志  1：事件

	DataType *uint64 `json:"DataType,omitempty" name:"DataType"`
	// 分步查询类型，为空时，所有数据都查。值为log_detail时，只查日志详情，值为date_histogram时，只查直方图和总数

	StepQuery *string `json:"StepQuery,omitempty" name:"StepQuery"`
}

func (r *SearchLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchNotificationMsgs struct {
	// 开始时间

	StartsAt *string `json:"StartsAt,omitempty" name:"StartsAt"`
	// 结束时间

	EndsAt *string `json:"EndsAt,omitempty" name:"EndsAt"`
	// 翻页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 单页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// label matcher

	Query []*Matcher `json:"Query,omitempty" name:"Query"`
	// 消息Id

	SendId *string `json:"SendId,omitempty" name:"SendId"`
}

type CopyLogDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyLogDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyLogDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateRuleGroupTempleRequest struct {
	*tchttp.BaseRequest

	// 策略模板定义

	Command *CreateOrUpdateRuleGroupTemple `json:"Command,omitempty" name:"Command"`
}

func (r *CreateOrUpdateRuleGroupTempleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateRuleGroupTempleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRelatedAlertNamesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Result []*EventRelatedAlertNames `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRelatedAlertNamesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRelatedAlertNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportLogConfigCRRequest struct {
	*tchttp.BaseRequest

	// 导出日志配置Id列表

	LogConfigs []*LogConfigParam `json:"LogConfigs,omitempty" name:"LogConfigs"`
}

func (r *ExportLogConfigCRRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportLogConfigCRRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchMetaValuesRequest struct {
	*tchttp.BaseRequest

	// Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 模糊查询的MetaName

	MetaName *string `json:"MetaName,omitempty" name:"MetaName"`
}

func (r *SearchMetaValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchMetaValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddLogDeliveryTaskRequest struct {
	*tchttp.BaseRequest

	// 日志投递任务列表

	Task []*LogDeliveryTaskParamItem `json:"Task,omitempty" name:"Task"`
}

func (r *AddLogDeliveryTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLogDeliveryTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogRuleGroupTemplesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchLogRuleGroupTemplesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogRuleGroupTemplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRoutesByLabelSetRequest struct {
	*tchttp.BaseRequest

	// 路由匹配查询参数

	Params *SearchRoutesByLabelSet `json:"Params,omitempty" name:"Params"`
}

func (r *SearchRoutesByLabelSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRoutesByLabelSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParsedField struct {
	// 字段名

	FieldKey *string `json:"FieldKey,omitempty" name:"FieldKey"`
	// 字段值

	FieldValue *string `json:"FieldValue,omitempty" name:"FieldValue"`
}

type SearchInfraTemplatesRefs struct {
	// 云哨模板meta信息参数对象

	TemplateMetas []*InfraTemplateMeta `json:"TemplateMetas,omitempty" name:"TemplateMetas"`
}

type DeleteLogDeliveryTargetRequest struct {
	*tchttp.BaseRequest

	// 日志投递下游源列表

	Target []*LogDeliveryTargetParamItem `json:"Target,omitempty" name:"Target"`
}

func (r *DeleteLogDeliveryTargetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogDeliveryTargetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PointsObject struct {
	// 监控实例的维度组合

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 监控数据点数组，每个点的时间跨度为一个Period值

	Points []*float64 `json:"Points,omitempty" name:"Points"`
}

type CreateRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRuleGroupTempleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRuleGroupTempleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRuleGroupTempleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLogTransformTaskRequest struct {
	*tchttp.BaseRequest

	// 搜索任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 任务id搜索

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 按照Enabled状态过滤

	Enabled *uint64 `json:"Enabled,omitempty" name:"Enabled"`
}

func (r *ListLogTransformTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLogTransformTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataQueryParams struct {
	// 查询表达式

	Query *string `json:"Query,omitempty" name:"Query"`
	// 时间，支持RFC3339和Unixj时间戳

	Time *string `json:"Time,omitempty" name:"Time"`
	// 超时时间

	Timeout *string `json:"Timeout,omitempty" name:"Timeout"`
	// 开始时间，支持RFC3339和Unixj时间戳

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束时间，支持RFC3339和Unixj时间戳

	End *string `json:"End,omitempty" name:"End"`
	// 查询配置条件

	Matches *string `json:"Matches,omitempty" name:"Matches"`
	// 步长

	Step *string `json:"Step,omitempty" name:"Step"`
	// 限制条数

	Limit *string `json:"Limit,omitempty" name:"Limit"`
	// BACKWARD倒序 FORWARD正序

	Direction *string `json:"Direction,omitempty" name:"Direction"`
}

type EventRelatedAlertNames struct {
	// 事件唯一标识，事件名

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 关联该事件的告警名

	RelatedAlertNames []*string `json:"RelatedAlertNames,omitempty" name:"RelatedAlertNames"`
}

type DeleteLogPanelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogPanelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogPanelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLogConfigStorageStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志投递状态列表

		ConfigState []*ConfigStateItem `json:"ConfigState,omitempty" name:"ConfigState"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateLogConfigStorageStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogConfigStorageStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardTag struct {
	// Tag名称

	Term *string `json:"Term,omitempty" name:"Term"`
	// 出现次数

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type CreateOrUpdateLogRuleGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrUpdateLogRuleGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateLogRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenLineHeaderRegexRequest struct {
	*tchttp.BaseRequest

	// 日志样例

	Example *string `json:"Example,omitempty" name:"Example"`
}

func (r *GenLineHeaderRegexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenLineHeaderRegexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAllUnitRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAllUnitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAllUnitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSilenceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSilenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSilenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaResourceTypeUpdateCommand struct {
	// 资源对象类型名称，在租户和资源组下唯一

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资源对象类型别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 生效资源对象，支持通配，如{tcs_product="*"}

	Selector *string `json:"Selector,omitempty" name:"Selector"`
}

type DeleteLogPanelsRequest struct {
	*tchttp.BaseRequest

	// 删除参数

	Params *DeleteLogPanels `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteLogPanelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogPanelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogRuleGroupTemplesRequest struct {
	*tchttp.BaseRequest

	// 批量删除ids，"Ids":"1,2,3"

	Params *BatchDeleteIds `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteLogRuleGroupTemplesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogRuleGroupTemplesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaResourceTypeQueryParam struct {
	// 租户ID，对应AppId，通常不需要传，默认取当前登录用户的AppId。

	Organization *string `json:"Organization,omitempty" name:"Organization"`
	// 要查询的对象在服务树上路径，格式为{a="b",c="d"}

	ResourcePath *string `json:"ResourcePath,omitempty" name:"ResourcePath"`
	// 查询类型，可选值为on/under，分别代表当前节点上的数据和当前节点下的数据两种含义。

	ListType *string `json:"ListType,omitempty" name:"ListType"`
	// 过滤类型，可选类型为metric和event，可不填，代表取所有类型

	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`
}

type UpdateMetaResourceTypeRequest struct {
	*tchttp.BaseRequest

	// 资源组归属信息，运营端使用

	Params *ResourceOwnerQueryParam `json:"Params,omitempty" name:"Params"`
	// 资源对象类型名称

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源对象类型信息，仅Alias和ResourcePath支持修改

	Command *MetaResourceTypeUpdateCommand `json:"Command,omitempty" name:"Command"`
}

func (r *UpdateMetaResourceTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMetaResourceTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardContent struct {
	// 仪表盘元数据

	Meta *DashboardMeta `json:"Meta,omitempty" name:"Meta"`
	// 仪表盘内容

	Dashboard *DummyContent `json:"Dashboard,omitempty" name:"Dashboard"`
}

type LogDeliveryTargetResponseItem struct {
	// 日志投递下游源id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 日志投递下游源名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志投递下游源类型

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 日志投递下游源配置

	TargetConfig *string `json:"TargetConfig,omitempty" name:"TargetConfig"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ResourceOwnerQueryParam struct {
	// 租户ID，对应AppId，运营端代理租户身份时使用

	Organization *string `json:"Organization,omitempty" name:"Organization"`
}

type ListMetaResourceTypeRequest struct {
	*tchttp.BaseRequest

	// 资源组查询参数

	Params *MetaResourceTypeQueryParam `json:"Params,omitempty" name:"Params"`
}

func (r *ListMetaResourceTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMetaResourceTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyLogConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true 校验成功 false 校验失败

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyLogConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSilenceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSilenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSilenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogDashboardsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogDashboardsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogDashboardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLogDeliveryTaskRequest struct {
	*tchttp.BaseRequest

	// 搜索任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 投递状态搜索

	State *uint64 `json:"State,omitempty" name:"State"`
	// 任务id搜索

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *ListLogDeliveryTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLogDeliveryTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAllUnitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAllUnitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAllUnitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLogAlertRuleRequest struct {
	*tchttp.BaseRequest

	// 报警策略ID

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 报警策略名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 报警级别

	AlertSeverity *string `json:"AlertSeverity,omitempty" name:"AlertSeverity"`
	// 根据Enabled过滤

	Enabled *uint64 `json:"Enabled,omitempty" name:"Enabled"`
	// 根据ReadOnly过滤

	ReadOnly *bool `json:"ReadOnly,omitempty" name:"ReadOnly"`
}

func (r *ListLogAlertRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLogAlertRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLogAlertRuleRequest struct {
	*tchttp.BaseRequest

	// 需要更新的日志策略ID

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 日志策略数据

	Rule *LogAlertRule `json:"Rule,omitempty" name:"Rule"`
}

func (r *UpdateLogAlertRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogAlertRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRoutes struct {
	// 路由名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品名称

	Product *string `json:"Product,omitempty" name:"Product"`
	// 创建人

	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`
	// 更新人

	UpdatedBy *string `json:"UpdatedBy,omitempty" name:"UpdatedBy"`
	// page offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// page limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 升序字段，filed1,filed2,...

	Asc *string `json:"Asc,omitempty" name:"Asc"`
	// 降序字段，field1,filed2,...

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 是否启用正则匹配，可选true、false，默认false

	IsRegex *string `json:"IsRegex,omitempty" name:"IsRegex"`
}

type CreateOrUpdateDashboardRequest struct {
	*tchttp.BaseRequest

	// 创建仪表盘内容

	Command *CreateDashboardCommand `json:"Command,omitempty" name:"Command"`
}

func (r *CreateOrUpdateDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFiringRulesRequest struct {
	*tchttp.BaseRequest

	// offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 升序

	Asc *string `json:"Asc,omitempty" name:"Asc"`
	// 规则名称

	AlertName *string `json:"AlertName,omitempty" name:"AlertName"`
	// 降序

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 策略名称

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 等级

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 类别

	Class *string `json:"Class,omitempty" name:"Class"`
}

func (r *GetFiringRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFiringRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRuleGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRuleGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListInfraTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListInfraTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInfraTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchNotification struct {
	// notification group label hash

	Fingerprint *string `json:"Fingerprint,omitempty" name:"Fingerprint"`
	// route name

	Name *string `json:"Name,omitempty" name:"Name"`
	// firing/resolved

	Status *string `json:"Status,omitempty" name:"Status"`
	// page size

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// page num

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 升序排列的字段，逗号分隔

	Asc *string `json:"Asc,omitempty" name:"Asc"`
	// 降序排列的字段，逗号分隔

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 是否启用正则匹配，可选true、false，默认false

	IsRegex *string `json:"IsRegex,omitempty" name:"IsRegex"`
}

type CreateQueryRecordRequest struct {
	*tchttp.BaseRequest

	// 用户id

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 查询名称

	QueryName *string `json:"QueryName,omitempty" name:"QueryName"`
}

func (r *CreateQueryRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateQueryRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogRuleGroupTemplesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogRuleGroupTemplesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogRuleGroupTemplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命中条件总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 查询耗时 单位 ms

		Took *float64 `json:"Took,omitempty" name:"Took"`
		// 命中数据信息

		Hits []*Hit `json:"Hits,omitempty" name:"Hits"`
		// 聚合信息

		Aggregations *Aggregations `json:"Aggregations,omitempty" name:"Aggregations"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigStateItem struct {
	// 日志配置id

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 是否关联投递任务

	Delivery *uint64 `json:"Delivery,omitempty" name:"Delivery"`
	// 是否投递到日志平台后端存储

	BackendStorage *uint64 `json:"BackendStorage,omitempty" name:"BackendStorage"`
}

type InfraTemplate struct {
	// 模板ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 模板版本

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 模板类型

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
	// 模板内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 模板来源

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
}

type SearchNotificationAlertsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchNotificationAlertsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchNotificationAlertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogInfo struct {
	// 纳秒级时间戳

	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
	// 原始日志

	OriginalLog *string `json:"OriginalLog,omitempty" name:"OriginalLog"`
}

type ListLogDeliveryTargetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志投递下游源总量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 日志投递下游源列表

		Target []*LogDeliveryTargetResponseItem `json:"Target,omitempty" name:"Target"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListLogDeliveryTargetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLogDeliveryTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyLogConfigRequest struct {
	*tchttp.BaseRequest

	// 待校验字段(TimeFormat,LineHeaderRegex,Regex)

	Field *string `json:"Field,omitempty" name:"Field"`
	// 待校验字段值

	FieldValue *string `json:"FieldValue,omitempty" name:"FieldValue"`
	// 待校验字段规则

	FieldRule *string `json:"FieldRule,omitempty" name:"FieldRule"`
}

func (r *VerifyLogConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyLogConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogFilter struct {
	// 表达式的key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 表达式操作符

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 表达式的值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateDashFolderRequest struct {
	*tchttp.BaseRequest

	// 创建仪表盘目录内容

	Command *CreateDashFolderCommand `json:"Command,omitempty" name:"Command"`
}

func (r *CreateDashFolderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDashFolderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRouteRequest struct {
	*tchttp.BaseRequest

	// 路由信息struct

	Command *CreateRoute `json:"Command,omitempty" name:"Command"`
}

func (r *CreateRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleGroupTemplesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRuleGroupTemplesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRuleGroupTemplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRuleGroupTempleRequest struct {
	*tchttp.BaseRequest

	// 策略模板 id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetRuleGroupTempleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRuleGroupTempleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置id

		ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenRegexRequest struct {
	*tchttp.BaseRequest

	// 划线值

	FieldContent *string `json:"FieldContent,omitempty" name:"FieldContent"`
	// 划线前缀

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
}

func (r *GenRegexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenRegexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDataBaradMetricRequest struct {
	*tchttp.BaseRequest

	// 命名空间，每个云产品会有一个命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 监控统计周期。默认为取值为300，单位为s

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 起始时间，如 2018-01-01 00:00:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，默认为当前时间。 endTime不能小于startTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 实例对象的维度组合

	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
}

func (r *GetDataBaradMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDataBaradMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMetaMetricRequest struct {
	*tchttp.BaseRequest

	// 资源对象类型归属信息

	Params *MetaMetricQueryParam `json:"Params,omitempty" name:"Params"`
	// 资源对象类型名称

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *ListMetaMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMetaMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddLogTransformTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddLogTransformTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLogTransformTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRoute struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品名称，支持多个

	Products []*string `json:"Products,omitempty" name:"Products"`
	// 告警等级，支持多个和all

	Severities []*string `json:"Severities,omitempty" name:"Severities"`
	// 是否发送告警

	SendFiring *bool `json:"SendFiring,omitempty" name:"SendFiring"`
	// 是否发送恢复

	SendResolved *bool `json:"SendResolved,omitempty" name:"SendResolved"`
	// 发送通道，支持微信、短信、企业微信、邮件，callback

	Channels []*LabelPair `json:"Channels,omitempty" name:"Channels"`
	// 用户id

	Users []*LabelPair `json:"Users,omitempty" name:"Users"`
	// 用户组id

	Groups []*LabelPair `json:"Groups,omitempty" name:"Groups"`
	// 静默时间，也即重发间隔

	RepeatInterval *string `json:"RepeatInterval,omitempty" name:"RepeatInterval"`
	// 策略，支持多个和all

	Strategies []*string `json:"Strategies,omitempty" name:"Strategies"`
	// 0,1: product+severity 2 :基于指标告警策略 3 基于事件告警策略

	Flag *int64 `json:"Flag,omitempty" name:"Flag"`
}

type Hit struct {
	// 标签信息

	Labels []*Label `json:"Labels,omitempty" name:"Labels"`
	// 日志信息

	LogInfos []*LogInfo `json:"LogInfos,omitempty" name:"LogInfos"`
}

type GetNotificationMsgRequest struct {
	*tchttp.BaseRequest

	// 消息id

	MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
}

func (r *GetNotificationMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNotificationMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLogDeliveryTaskRequest struct {
	*tchttp.BaseRequest

	// 日志投递任务列表

	Task []*LogDeliveryTaskParamItem `json:"Task,omitempty" name:"Task"`
}

func (r *UpdateLogDeliveryTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogDeliveryTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditableElement struct {
	// 元素类型(目前只有Val，后续会扩展)

	Type *string `json:"Type,omitempty" name:"Type"`
	// 元素值

	Val *string `json:"Val,omitempty" name:"Val"`
}

type DescribeLogPanel struct {
	// 面板id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 面板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 仪表盘id

	DashboardId *int64 `json:"DashboardId,omitempty" name:"DashboardId"`
}

type AddMetaResourceTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 资源对象类型对象

		Data *MetaResourceType `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddMetaResourceTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMetaResourceTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteQueryRecordRequest struct {
	*tchttp.BaseRequest

	// 用户id

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 记录id

	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`
}

func (r *DeleteQueryRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteQueryRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNotificationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNotificationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLogTransformTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志处理任务总量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 任务详情

		Tasks *LogTransformTaskParamItem `json:"Tasks,omitempty" name:"Tasks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListLogTransformTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLogTransformTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeMetric struct {
	// 指标类型。count: 全部查询总数；field_count: 包含字段计数；cardinality:字段值去重计数；sum: 字段值加和；avg: 字段值平均值；max:字段值最大值；min：字段值最小值

	Type *string `json:"Type,omitempty" name:"Type"`
	// 字段名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 指标前端展示名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type CreateResourceCommand struct {
	// 资源对象名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 资源归属，取租户AppId

	Organization *string `json:"Organization,omitempty" name:"Organization"`
	// 类型标签

	LabelName *string `json:"LabelName,omitempty" name:"LabelName"`
	// 服务树节点选择标签，格式为{a="b",c="d"}

	Labels *string `json:"Labels,omitempty" name:"Labels"`
}

type DeleteLogDashboards struct {
	// 待删除ids 逗号分隔

	Ids *string `json:"Ids,omitempty" name:"Ids"`
}

type LogAlertRule struct {
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品名

	Product *string `json:"Product,omitempty" name:"Product"`
	// 产品信息

	Metas []*Meta `json:"Metas,omitempty" name:"Metas"`
	// 过滤规则

	Filters []*LogFilter `json:"Filters,omitempty" name:"Filters"`
	// 查询的时间范围

	EvaluationInterval *uint64 `json:"EvaluationInterval,omitempty" name:"EvaluationInterval"`
	// 查询的执行时间间隔

	ScrapeInterval *uint64 `json:"ScrapeInterval,omitempty" name:"ScrapeInterval"`
	// 报警的条件

	AlertCondition []*LogAggregator `json:"AlertCondition,omitempty" name:"AlertCondition"`
	// 报警级别

	AlertSeverity *string `json:"AlertSeverity,omitempty" name:"AlertSeverity"`
	// 满足条件几次才报警

	AlertConsecutive *uint64 `json:"AlertConsecutive,omitempty" name:"AlertConsecutive"`
	// 是否开启

	Enabled *int64 `json:"Enabled,omitempty" name:"Enabled"`
	// 报警描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 扩展字段

	Extension *string `json:"Extension,omitempty" name:"Extension"`
	// 表示是否只读(预制)

	ReadOnly *bool `json:"ReadOnly,omitempty" name:"ReadOnly"`
}

type CreateOrUpdateRuleGroupTempleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrUpdateRuleGroupTempleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateRuleGroupTempleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogDashboardsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchLogDashboardsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogDashboardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInfraTemplates struct {
	// 模板Ids，逗号分隔

	Ids *string `json:"Ids,omitempty" name:"Ids"`
	// 模板类型

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
	// 模板Names，逗号分隔

	Names *string `json:"Names,omitempty" name:"Names"`
}

type SearchRuleGroupsRequest struct {
	*tchttp.BaseRequest

	// 搜索条件

	Params *SearchRuleGroups `json:"Params,omitempty" name:"Params"`
}

func (r *SearchRuleGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRuleGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashFolderQueryResult struct {
	// 数据ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// Unique ID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 目录名称

	Title *string `json:"Title,omitempty" name:"Title"`
}

type DeleteResourceObjectRequest struct {
	*tchttp.BaseRequest

	// 资源归属

	Params *ResourceOwnerQueryParam `json:"Params,omitempty" name:"Params"`
	// 服务树节点标签

	Labels *string `json:"Labels,omitempty" name:"Labels"`
}

func (r *DeleteResourceObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteResourceObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLogAlertRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志报警策略总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 日志策略内容

		Rules []*LogAlertRule `json:"Rules,omitempty" name:"Rules"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListLogAlertRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLogAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaResourceTypeEvent struct {
	// 事件名称

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 事件维度

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
	// 事件中文名称

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 事件hash

	Fingerprint *uint64 `json:"Fingerprint,omitempty" name:"Fingerprint"`
	// 资源对象类型ID

	ResourceTypeId *int64 `json:"ResourceTypeId,omitempty" name:"ResourceTypeId"`
}

type SearchResourceCommand struct {
	// 资源归属租户

	Organization *string `json:"Organization,omitempty" name:"Organization"`
	// 服务树节点选择器

	Selector *string `json:"Selector,omitempty" name:"Selector"`
}

type DeleteDashboardByUidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDashboardByUidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardByUidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDefaultGroupByResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDefaultGroupByResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDefaultGroupByResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProxyTypeInfoRequest struct {
	*tchttp.BaseRequest

	// 额外路径

	AdditionalPath *string `json:"AdditionalPath,omitempty" name:"AdditionalPath"`
	// 代理方法，当前只支持GET

	Method *string `json:"Method,omitempty" name:"Method"`
	// 数据类型，当前支持metrics和logs

	Name *string `json:"Name,omitempty" name:"Name"`
	// 数据查询参数

	Params *DataQueryParams `json:"Params,omitempty" name:"Params"`
}

func (r *ProxyTypeInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProxyTypeInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRuleGroupTemplesRequest struct {
	*tchttp.BaseRequest

	// 搜索条件

	Params *SearchRuleGroupTemples `json:"Params,omitempty" name:"Params"`
}

func (r *SearchRuleGroupTemplesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRuleGroupTemplesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateRuleGroupTemple struct {
	// 规则名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 评估间隔，建议1m

	Interval *string `json:"Interval,omitempty" name:"Interval"`
	// 归属产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 规则数组

	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`
	// 对象类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 乐观锁版本

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 策略Id，根据id是否>0区分create还是update

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 更新时间

	Updated *string `json:"Updated,omitempty" name:"Updated"`
	// labels

	Labels *LabelPair `json:"Labels,omitempty" name:"Labels"`
	// 被引用数

	RefNum *int64 `json:"RefNum,omitempty" name:"RefNum"`
}

type GetAlertTrendRequest struct {
	*tchttp.BaseRequest

	// 参数对象

	Params *GetAlertTrend `json:"Params,omitempty" name:"Params"`
}

func (r *GetAlertTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchAlertsRequest struct {
	*tchttp.BaseRequest

	// 搜索条件

	Params *SearchAlert `json:"Params,omitempty" name:"Params"`
}

func (r *SearchAlertsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchAlertsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Label struct {
	// 标签名

	LabelName *string `json:"LabelName,omitempty" name:"LabelName"`
	// 标签值

	LabelValue *string `json:"LabelValue,omitempty" name:"LabelValue"`
}

type LogDeliveryTargetParamItem struct {
	// 日志投递下游源id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 日志投递下游源名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志投递下游源类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 日志投递下游源配置

	Config *string `json:"Config,omitempty" name:"Config"`
}

type UpdateMetaMetricRequest struct {
	*tchttp.BaseRequest

	// 资源归属信息

	Params *ResourceOwnerQueryParam `json:"Params,omitempty" name:"Params"`
	// 资源对象类型名称

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 指标指纹ID

	Fingerprint *string `json:"Fingerprint,omitempty" name:"Fingerprint"`
	// 指标更新信息，包含Alias，Type，Help，Unit

	Command *MetaMetricUpdateCommand `json:"Command,omitempty" name:"Command"`
}

func (r *UpdateMetaMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMetaMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDashFolderByUidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// 仪表盘目录更新结果

		Data *DashFolderContent `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDashFolderByUidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDashFolderByUidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLogDeliveryTargetRequest struct {
	*tchttp.BaseRequest

	// 日志投递下游源列表

	Target []*LogDeliveryTargetParamItem `json:"Target,omitempty" name:"Target"`
}

func (r *UpdateLogDeliveryTargetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogDeliveryTargetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLogTransformTaskEnabledResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateLogTransformTaskEnabledResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogTransformTaskEnabledResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalyseLogsRequest struct {
	*tchttp.BaseRequest

	// LogQL查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 返回条数限制，默认10条

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分桶参数

	Buckets []*AnalyzeBucket `json:"Buckets,omitempty" name:"Buckets"`
	// 指标参数

	Metrics []*AnalyzeMetric `json:"Metrics,omitempty" name:"Metrics"`
	// 过滤参数

	Filters []*AnalyzeFilter `json:"Filters,omitempty" name:"Filters"`
	// 排序参数

	Sorts []*AnalyzeSort `json:"Sorts,omitempty" name:"Sorts"`
	// 起始时间 纳秒时间戳 或 RFC3339Nano

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 截止时间 纳秒时间戳 或 RFC3339Nano

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *AnalyseLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AnalyseLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashboardVersionsByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仪表盘更新历史记录

		Data []*DashboardVersion `json:"Data,omitempty" name:"Data"`
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDashboardVersionsByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardVersionsByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogDashboards struct {
	// 是否收藏 true 已收藏，false 未收藏

	IsFavorite *string `json:"IsFavorite,omitempty" name:"IsFavorite"`
	// page offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// page limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 升序字段: field1[,field2 ...]

	Asc *string `json:"Asc,omitempty" name:"Asc"`
	// 降序字段: field1[,field2 ...]

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 仪表盘名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否正则匹配 true false

	IsRegex *string `json:"IsRegex,omitempty" name:"IsRegex"`
}

type ExportLogsRequest struct {
	*tchttp.BaseRequest

	// Meta信息

	Metas []*Meta `json:"Metas,omitempty" name:"Metas"`
	// 日志名，可以选多个

	Name []*string `json:"Name,omitempty" name:"Name"`
	// 时间范围起始，纳秒时间戳，默认当前时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 时间范围结束，纳秒时间戳，默认一小时前

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 返回记录条数，默认值：500，最大值：1000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// LogQL

	Query *string `json:"Query,omitempty" name:"Query"`
	// 排序

	Sort []*SortInfo `json:"Sort,omitempty" name:"Sort"`
	// 查询方式，0: LogQL

	QueryType *uint64 `json:"QueryType,omitempty" name:"QueryType"`
	// 数据类型 默认 0：日志 1 事件

	DataType *uint64 `json:"DataType,omitempty" name:"DataType"`
}

func (r *ExportLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubSubMeta struct {
	// meta key

	MetaKey *string `json:"MetaKey,omitempty" name:"MetaKey"`
	// meta value

	MetaValue *string `json:"MetaValue,omitempty" name:"MetaValue"`
	// meta name

	MetaName *string `json:"MetaName,omitempty" name:"MetaName"`
	// name

	Names []*string `json:"Names,omitempty" name:"Names"`
	// 配置数量

	ConfigCount *uint64 `json:"ConfigCount,omitempty" name:"ConfigCount"`
}

type UpdateDashFolderCommand struct {
	// Unique Id

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 目录名称

	Title *string `json:"Title,omitempty" name:"Title"`
	// 目录版本

	Version *uint64 `json:"Version,omitempty" name:"Version"`
	// 是否覆写

	Overwrite *bool `json:"Overwrite,omitempty" name:"Overwrite"`
}

type DeleteLogDashboardsRequest struct {
	*tchttp.BaseRequest

	// 删除参数

	Params *DeleteLogDashboards `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteLogDashboardsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogDashboardsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogRuleGroupsRequest struct {
	*tchttp.BaseRequest

	// 搜索条件

	Params *SearchRuleGroups `json:"Params,omitempty" name:"Params"`
}

func (r *SearchLogRuleGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogRuleGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaResourceType struct {
	// 租户ID，对应AppId，当为*时，代表所有租户通用

	Organization *string `json:"Organization,omitempty" name:"Organization"`
	// 资源对象类型名称，在租户和资源组下唯一

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资源对象类型别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 生效资源对象，支持通配，如{tcs_product="*"}

	Selector *string `json:"Selector,omitempty" name:"Selector"`
}

type GetDefaultGroupByRequest struct {
	*tchttp.BaseRequest
}

func (r *GetDefaultGroupByRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDefaultGroupByRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRoutesRequest struct {
	*tchttp.BaseRequest

	// 路由搜索条件

	Params *SearchRoutes `json:"Params,omitempty" name:"Params"`
}

func (r *SearchRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDashFolderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建仪表盘目录结果内容

		Data *DashFolderContent `json:"Data,omitempty" name:"Data"`
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDashFolderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDashFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaMetric struct {
	// 指标名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 指标Label列表

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 说明

	Help *string `json:"Help,omitempty" name:"Help"`
	// 指纹ID

	Fingerprint *uint64 `json:"Fingerprint,omitempty" name:"Fingerprint"`
}

type CreateOrUpdateLogPanelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrUpdateLogPanelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateLogPanelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateRuleGroupRequest struct {
	*tchttp.BaseRequest

	// 策略定义

	Command *CreateOrUpdateRuleGroup `json:"Command,omitempty" name:"Command"`
}

func (r *CreateOrUpdateRuleGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateRuleGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRetentionsRequest struct {
	*tchttp.BaseRequest

	// 搜索Retention

	Command *SearchRetention `json:"Command,omitempty" name:"Command"`
}

func (r *SearchRetentionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRetentionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogAggregator struct {
	// 表达式函数

	Aggregator *string `json:"Aggregator,omitempty" name:"Aggregator"`
	// 表达式计算的label

	Label *string `json:"Label,omitempty" name:"Label"`
	// 表达式的操作符

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 表达式的值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type AddLogDeliveryTargetRequest struct {
	*tchttp.BaseRequest

	// 日志投递下游源列表

	Target []*LogDeliveryTargetParamItem `json:"Target,omitempty" name:"Target"`
}

func (r *AddLogDeliveryTargetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLogDeliveryTargetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchDashboardsRequest struct {
	*tchttp.BaseRequest

	// 仪表盘查询参数

	Params *DashboardQueryParam `json:"Params,omitempty" name:"Params"`
}

func (r *SearchDashboardsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchDashboardsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddLogAlertRuleRequest struct {
	*tchttp.BaseRequest

	// 报警规则名称

	Rule *LogAlertRule `json:"Rule,omitempty" name:"Rule"`
}

func (r *AddLogAlertRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLogAlertRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyInfraTemplatesRequest struct {
	*tchttp.BaseRequest

	// 云哨模板批量应用接口参数对象

	Command *ApplyInfraTemplates `json:"Command,omitempty" name:"Command"`
}

func (r *ApplyInfraTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyInfraTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchAlert struct {
	// 告警名称，支持正则匹配，depreated

	Alertname *string `json:"Alertname,omitempty" name:"Alertname"`
	// 当前状态，firing/resolved

	Status *string `json:"Status,omitempty" name:"Status"`
	// 产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 地区

	Region *string `json:"Region,omitempty" name:"Region"`
	// 告警级别， 1、2、3、4

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 确认状态，true/false

	Confirmed *string `json:"Confirmed,omitempty" name:"Confirmed"`
	// time range start

	StartsAt *string `json:"StartsAt,omitempty" name:"StartsAt"`
	// time range end

	EndsAt *string `json:"EndsAt,omitempty" name:"EndsAt"`
	// page offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// page limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 升序排列字段，逗号分隔多个字段

	Asc *string `json:"Asc,omitempty" name:"Asc"`
	// 降序排列字段，逗号分隔多个字段

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 对象查询字段，支持正则匹配

	Object *string `json:"Object,omitempty" name:"Object"`
	// 屏蔽规则id

	SilencedBy *string `json:"SilencedBy,omitempty" name:"SilencedBy"`
	// 告警策略名称

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 告警名称，支持正则匹配，替代Alertname

	AlertName *string `json:"AlertName,omitempty" name:"AlertName"`
	// 是否启用正则匹配，可选true、false，默认false

	IsRegex *string `json:"IsRegex,omitempty" name:"IsRegex"`
	// 产品类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 聚合id

	NotificationId *int64 `json:"NotificationId,omitempty" name:"NotificationId"`
	// 告警分类：metric、log

	Class *string `json:"Class,omitempty" name:"Class"`
}

type SearchNotificationAlerts struct {
	// offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// label matchers

	Query []*Matcher `json:"Query,omitempty" name:"Query"`
}

type GetDashFolderByUidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仪表盘目录内容

		Data *DashFolderContent `json:"Data,omitempty" name:"Data"`
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDashFolderByUidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashFolderByUidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNotificationRequest struct {
	*tchttp.BaseRequest

	// notification id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetNotificationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNotificationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLogDeliveryTargetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateLogDeliveryTargetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogDeliveryTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLogDeliveryTargetTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 投递下游源类型列表

		Type []*string `json:"Type,omitempty" name:"Type"`
		// 投递下游源类型数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListLogDeliveryTargetTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLogDeliveryTargetTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddLogTransformTaskRequest struct {
	*tchttp.BaseRequest

	// 任务详情

	Tasks []*LogTransformTaskParamItem `json:"Tasks,omitempty" name:"Tasks"`
}

func (r *AddLogTransformTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLogTransformTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogDeliveryTaskRequest struct {
	*tchttp.BaseRequest

	// 日志投递任务列表

	Task []*LogDeliveryTaskParamItem `json:"Task,omitempty" name:"Task"`
}

func (r *DeleteLogDeliveryTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogDeliveryTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigInfo struct {
	// meta信息

	Metas []*Meta `json:"Metas,omitempty" name:"Metas"`
	// 日志用途标识

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志路径，支持通配

	Path *string `json:"Path,omitempty" name:"Path"`
	// 当path包含通配符时，该字段用于排除不必要的日志文件

	PathExcluded *string `json:"PathExcluded,omitempty" name:"PathExcluded"`
	// 是否启用多行匹配

	MultipleLine *bool `json:"MultipleLine,omitempty" name:"MultipleLine"`
	// 指定系统已有的parse

	ParserRef *string `json:"ParserRef,omitempty" name:"ParserRef"`
	// 用于手工指定配置，可用配置有json、regex等

	ParserFormat *string `json:"ParserFormat,omitempty" name:"ParserFormat"`
	// time字段名称

	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`
	// time格式

	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`
	// 提取正则

	Regex *string `json:"Regex,omitempty" name:"Regex"`
	// 行首正则

	LineHeaderRegex *string `json:"LineHeaderRegex,omitempty" name:"LineHeaderRegex"`
	// 当ParserFormat为LTSV时需要指定

	Type *string `json:"Type,omitempty" name:"Type"`
	// 用于手工指定配置，主要用于对特定定段再次进行解码

	Decoders []*string `json:"Decoders,omitempty" name:"Decoders"`
	// 配置状态，1 开启 0关闭

	State *uint64 `json:"State,omitempty" name:"State"`
	// 配置id

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 配置创建时间 秒级时间戳

	CreatedAt *uint64 `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 是否提取字段

	FieldExtraction *bool `json:"FieldExtraction,omitempty" name:"FieldExtraction"`
	// 是否使用系统时间

	UseSystemTime *bool `json:"UseSystemTime,omitempty" name:"UseSystemTime"`
	// 配置最近更新时间 秒级时间戳

	UpdatedAt *uint64 `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 日志样例

	Example *string `json:"Example,omitempty" name:"Example"`
	// 字段信息，json  样例:[{"Field":"cost","Type":"int"},{"Field":"log_level","Type":"string"},{"Field":"old_price","Type":"double"}]

	FieldMapping *string `json:"FieldMapping,omitempty" name:"FieldMapping"`
}

type DateHistogram struct {
	// 聚合日期字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// 时间间隔  s:秒，m:分，h：小时，d:天，w:周，M:月，y:年

	FixedInterval *string `json:"FixedInterval,omitempty" name:"FixedInterval"`
	// 时区，默认 Asia/Shanghai

	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
	// 最小匹配数，小于最小匹配数的日期聚合数据不展示

	MinDocCount *uint64 `json:"MinDocCount,omitempty" name:"MinDocCount"`
}

type SubMeta struct {
	// meta key

	MetaKey *string `json:"MetaKey,omitempty" name:"MetaKey"`
	// meta value

	MetaValue *string `json:"MetaValue,omitempty" name:"MetaValue"`
	// meta name

	MetaName *string `json:"MetaName,omitempty" name:"MetaName"`
	// 子meta

	SubMetas []*SubSubMeta `json:"SubMetas,omitempty" name:"SubMetas"`
	// 配置数量

	ConfigCount *uint64 `json:"ConfigCount,omitempty" name:"ConfigCount"`
}

type CreateOrUpdateDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrUpdateDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceTypeEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 资源对象元数据-对象类型事件列表

		Data []*MetaResourceTypeEvent `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListResourceTypeEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceTypeEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashFolderContent struct {
	// 数据库ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// Unique Id

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 目录名称

	Title *string `json:"Title,omitempty" name:"Title"`
	// 创建用户

	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`
	// 最近更新用户

	UpdatedBy *string `json:"UpdatedBy,omitempty" name:"UpdatedBy"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 更新时间

	Updated *string `json:"Updated,omitempty" name:"Updated"`
	// 是否包含ACL

	HasAcl *bool `json:"HasAcl,omitempty" name:"HasAcl"`
	// 版本

	Version *uint64 `json:"Version,omitempty" name:"Version"`
	// 是否可保存

	CanSave *bool `json:"CanSave,omitempty" name:"CanSave"`
	// 是否可编辑

	CanEdit *bool `json:"CanEdit,omitempty" name:"CanEdit"`
	// 是否可管理

	CanAdmin *bool `json:"CanAdmin,omitempty" name:"CanAdmin"`
}

type CreateOrUpdateRetentionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrUpdateRetentionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateRetentionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMetaResourceTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 资源类型列表

		Data []*MetaResourceType `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMetaResourceTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMetaResourceTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Aggregations struct {
	// 聚合名称

	AggName *string `json:"AggName,omitempty" name:"AggName"`
	// 聚合结果

	AggResults *AggResults `json:"AggResults,omitempty" name:"AggResults"`
}

type AddLogDeliveryTargetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddLogDeliveryTargetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLogDeliveryTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogTransformLineItem struct {
	// 日志内容行

	Line *string `json:"Line,omitempty" name:"Line"`
	// 处理函数

	Transforms []*LogTransformFunction `json:"Transforms,omitempty" name:"Transforms"`
}

type GetAlertTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAlertTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLogDeliveryTargetRequest struct {
	*tchttp.BaseRequest

	// 投递方式筛选

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 任务名称筛选

	Name *string `json:"Name,omitempty" name:"Name"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 投递下游源id筛选

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *ListLogDeliveryTargetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLogDeliveryTargetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQueryRecordsRequest struct {
	*tchttp.BaseRequest

	// 用户id

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 分页偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回记录条数，默认值：10，最大值：100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeQueryRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQueryRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMetaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// meta数组

		Meta []*Meta `json:"Meta,omitempty" name:"Meta"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMetaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMetaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAlertRequest struct {
	*tchttp.BaseRequest

	// alert id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetAlertRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashboardTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仪表盘Tag统计

		Data []*DashboardTag `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDashboardTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMetaMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMetaMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMetaMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAlertResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAlertResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListInfraTemplatesRequest struct {
	*tchttp.BaseRequest

	// 云哨模板查询接口参数对象

	Params *ListInfraTemplates `json:"Params,omitempty" name:"Params"`
}

func (r *ListInfraTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInfraTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源类型列表

		Data []*ResourceType `json:"Data,omitempty" name:"Data"`
		// 状态

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListResourceTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchAlertsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchAlertsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchAlertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQueryRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询记录列表

		Records []*QueryRecord `json:"Records,omitempty" name:"Records"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQueryRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQueryRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SortInfo struct {
	// 待排序字段

	SortField *string `json:"SortField,omitempty" name:"SortField"`
	// 顺序，asc：正序，desc：倒序

	Order *string `json:"Order,omitempty" name:"Order"`
}

type PreviewLogTransformTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Results []*string `json:"Results,omitempty" name:"Results"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PreviewLogTransformTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreviewLogTransformTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaResourceTypeCreateCommand struct {
	// 资源对象类型名称，在租户和资源组下唯一

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资源对象类型别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 生效资源对象，支持通配，如{tcs_product="*"}

	Selector *string `json:"Selector,omitempty" name:"Selector"`
}

type DeleteInfraTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInfraTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInfraTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInfraTemplatesRequest struct {
	*tchttp.BaseRequest

	// 云哨模板批量删除接口参数对象

	Params *DeleteInfraTemplates `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteInfraTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInfraTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeBucket struct {
	// 分桶字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// 前端指标展示

	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type DashAclItem struct {
	// 用户ID

	UserId *int64 `json:"UserId,omitempty" name:"UserId"`
	// 角色名称，可先值为Viewer、Editor、Admin

	Role *string `json:"Role,omitempty" name:"Role"`
	// 权限，可选值为1,2,4，分别代表View,Edit,Admin

	Permission *int64 `json:"Permission,omitempty" name:"Permission"`
}

type InfraTemplateMeta struct {
	// 模板类型

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
	// 模板ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type DescribeLogDashboard struct {
	// dashboard id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// dashboard name

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeLogConfigsRequest struct {
	*tchttp.BaseRequest

	// 日志名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// meta信息

	Metas []*Meta `json:"Metas,omitempty" name:"Metas"`
	// 分页偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回记录条数，默认值：10，最大值：100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 配置id

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribeLogConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogField struct {
	// 字段类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 字段名

	Field *string `json:"Field,omitempty" name:"Field"`
}

type Matcher struct {
	// 类型0,1,2,3 -> =,!=,=~,!~

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// label name

	Name *string `json:"Name,omitempty" name:"Name"`
	// label value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type GetAlertStatsRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAlertStatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertStatsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogDeliveryTaskParamItem struct {
	// 日志投递任务id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 日志投递任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 投递任务关联日志

	Metas []*Meta `json:"Metas,omitempty" name:"Metas"`
	// 接收端类型

	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`
	// 接收端id

	TargetId *uint64 `json:"TargetId,omitempty" name:"TargetId"`
	// 投递数据格式

	Format *string `json:"Format,omitempty" name:"Format"`
	// 启用/停用投递任务

	State *uint64 `json:"State,omitempty" name:"State"`
}

type CustomLabelItem struct {
	// 自定义标签类型

	LabelType *uint64 `json:"LabelType,omitempty" name:"LabelType"`
	// 自定义标签key

	LabelKey *string `json:"LabelKey,omitempty" name:"LabelKey"`
	// 自定义标签value

	LabelValue *string `json:"LabelValue,omitempty" name:"LabelValue"`
}

type AddLogAlertRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddLogAlertRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLogAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceObjectRequest struct {
	*tchttp.BaseRequest

	// 查询资源服务对象参数

	Params *SearchResourceCommand `json:"Params,omitempty" name:"Params"`
}

func (r *ListResourceObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSilences struct {
	// 屏蔽规则id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 屏蔽规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 级别

	Level *string `json:"Level,omitempty" name:"Level"`
	// 对象类型、产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 创建者

	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`
	// 更新者

	UpdatedBy *string `json:"UpdatedBy,omitempty" name:"UpdatedBy"`
	// page offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// page limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 升序字段，field1,field2...

	Asc *string `json:"Asc,omitempty" name:"Asc"`
	// 降序字段，field1,field2...

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 屏蔽状态，active/pending/expired

	Effective *string `json:"Effective,omitempty" name:"Effective"`
	// 时间范围，开始时间

	StartsAt *string `json:"StartsAt,omitempty" name:"StartsAt"`
	// 时间范围，结束时间

	EndsAt *string `json:"EndsAt,omitempty" name:"EndsAt"`
	// 自定义变更人员，默认是当前用户，多个用户间逗号分隔

	Updater *string `json:"Updater,omitempty" name:"Updater"`
	// 是否启用正则匹配，可选true、false，默认false

	IsRegex *string `json:"IsRegex,omitempty" name:"IsRegex"`
}

type AnalyseLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AnalyseLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AnalyseLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfirmAlertRequest struct {
	*tchttp.BaseRequest

	// alert id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *ConfirmAlertRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfirmAlertRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteResourceObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteResourceObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteResourceObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashFolderByUidRequest struct {
	*tchttp.BaseRequest

	// 仪表盘目录唯一ID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
}

func (r *GetDashFolderByUidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashFolderByUidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSilenceRequest struct {
	*tchttp.BaseRequest

	// 屏蔽策略 id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetSilenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSilenceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLogConfigStorageStateRequest struct {
	*tchttp.BaseRequest

	// 日志配置id

	ConfigIds []*string `json:"ConfigIds,omitempty" name:"ConfigIds"`
	// 日志是否存储到日志平台后端 是：1，否：0

	BackendStorage *uint64 `json:"BackendStorage,omitempty" name:"BackendStorage"`
}

func (r *UpdateLogConfigStorageStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogConfigStorageStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志配置列表

		Configs []*ConfigInfo `json:"Configs,omitempty" name:"Configs"`
		// 日志总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashFoldersRequest struct {
	*tchttp.BaseRequest

	// 仪表盘目录查询参数

	Params *DashFolderQueryParam `json:"Params,omitempty" name:"Params"`
}

func (r *GetDashFoldersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashFoldersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLogDeliveryTargetTypeRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListLogDeliveryTargetTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLogDeliveryTargetTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardQueryParam struct {
	// 查询字符串

	Query *string `json:"Query,omitempty" name:"Query"`
	// 页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Page

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 是否加星

	Starred *string `json:"Starred,omitempty" name:"Starred"`
	// 标签

	Tag []*string `json:"Tag,omitempty" name:"Tag"`
	// 类型，可以是dashboard或dash-folder，默认两者都返回

	Type *string `json:"Type,omitempty" name:"Type"`
}

type ListInfraTemplates struct {
	// 模板所属产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 模板Ids，逗号分隔

	Ids *string `json:"Ids,omitempty" name:"Ids"`
	// 模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 模板所属模块

	Module *string `json:"Module,omitempty" name:"Module"`
	// 模板所属业务（log/event/metric）

	BizType *string `json:"BizType,omitempty" name:"BizType"`
	// 模板类型（log_rule_group/metric_rule_group）

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
	// 模板来源（preset/preset_cr/custom）

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 是否已发布(“true”/"false"")

	IsDeployed *string `json:"IsDeployed,omitempty" name:"IsDeployed"`
	// 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 单页限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 是否需要处理

	IsHandled *string `json:"IsHandled,omitempty" name:"IsHandled"`
}

type AddResourceObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 资源对象

		Data *ResourceObject `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddResourceObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddResourceObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMetaResourceTypeRequest struct {
	*tchttp.BaseRequest

	// 资源对象归属信息，运营端使用

	Params *ResourceOwnerQueryParam `json:"Params,omitempty" name:"Params"`
	// 资源对象类型名称

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *DeleteMetaResourceTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMetaResourceTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSilenceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSilenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSilenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreviewLogTransformTaskRequest struct {
	*tchttp.BaseRequest

	// 任务处理详情

	Tasks *LogTransformLineItem `json:"Tasks,omitempty" name:"Tasks"`
}

func (r *PreviewLogTransformTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreviewLogTransformTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLogTransformTaskRequest struct {
	*tchttp.BaseRequest

	// 任务处理详情

	Tasks *LogTransformTaskParamItem `json:"Tasks,omitempty" name:"Tasks"`
}

func (r *UpdateLogTransformTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogTransformTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditableExpr struct {
	// 已提取可编辑元素用占位符替代的表达式

	HandledExpr *string `json:"HandledExpr,omitempty" name:"HandledExpr"`
	// 可编辑元素

	Elements []*EditableElement `json:"Elements,omitempty" name:"Elements"`
}

type GetFiringRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFiringRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFiringRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogPanelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogPanelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogPanelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志配置

		ConfigInfo *ConfigInfo `json:"ConfigInfo,omitempty" name:"ConfigInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLogConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSilenceRequest struct {
	*tchttp.BaseRequest

	// 屏蔽struct

	Command *CreateSilence `json:"Command,omitempty" name:"Command"`
}

func (r *CreateSilenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSilenceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchNotificationsRequest struct {
	*tchttp.BaseRequest

	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// Fingerprint

	Fingerprint *string `json:"Fingerprint,omitempty" name:"Fingerprint"`
	// route name also receiver name

	Name *string `json:"Name,omitempty" name:"Name"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// asc

	Asc *string `json:"Asc,omitempty" name:"Asc"`
	// desc

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 模糊

	IsRegex *bool `json:"IsRegex,omitempty" name:"IsRegex"`
}

func (r *SearchNotificationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchNotificationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateLogDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrUpdateLogDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateLogDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLogConfigStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateLogConfigStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogConfigStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateRetention struct {
	// 主键，根据Id是否大于0来区分create or update

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 名称，唯一

	Name *string `json:"Name,omitempty" name:"Name"`
	// 运行间隔，默认1d

	Interval *string `json:"Interval,omitempty" name:"Interval"`
	// 数据保留时长

	Retention *string `json:"Retention,omitempty" name:"Retention"`
	// 辅助信息，比如{"max":"365d", "min":"90d"}

	Annotation []*LabelPair `json:"Annotation,omitempty" name:"Annotation"`
	// 乐观锁

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 更新时间

	Updated *string `json:"Updated,omitempty" name:"Updated"`
}

type CreateOrUpdateLogRuleGroupRequest struct {
	*tchttp.BaseRequest

	// 策略定义

	Command *CreateOrUpdateRuleGroup `json:"Command,omitempty" name:"Command"`
}

func (r *CreateOrUpdateLogRuleGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateLogRuleGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashboardContentByUidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 仪表盘内容

		Data *DashboardContent `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDashboardContentByUidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardContentByUidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchNotificationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchNotificationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchNotificationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyLogDashboardRequest struct {
	*tchttp.BaseRequest

	// LogDashboard

	Command *CreateOrUpdateLogDashboard `json:"Command,omitempty" name:"Command"`
}

func (r *CopyLogDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyLogDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRouteRequest struct {
	*tchttp.BaseRequest

	// 路由ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceTypeEventRequest struct {
	*tchttp.BaseRequest

	// 资源对象元数据-对象类型名称

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 事件元数据查询参数

	Params *MetaEventQueryParam `json:"Params,omitempty" name:"Params"`
}

func (r *ListResourceTypeEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceTypeEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceType struct {
	// 类型名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 类型名称别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 父类型

	Parent *string `json:"Parent,omitempty" name:"Parent"`
}

type RestoreDashboardCommand struct {
	// 仪表盘版本

	Version *int64 `json:"Version,omitempty" name:"Version"`
}
