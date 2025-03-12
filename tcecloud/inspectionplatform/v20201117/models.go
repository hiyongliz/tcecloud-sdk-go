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

package v20201117

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type CronJob struct {
	// Receiver 报告接收者，后端不做校验

	Receivers []*Receiver `json:"Receivers,omitempty" name:"Receivers"`
	// Name 计划名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Remark 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// Items 巡检项

	Items []*string `json:"Items,omitempty" name:"Items"`
	// ReportTypes 报告格式，支持PDF、HTML、CSV

	ReportTypes []*string `json:"ReportTypes,omitempty" name:"ReportTypes"`
	// Notification 告警设置

	Notification *Notification `json:"Notification,omitempty" name:"Notification"`
	// Schedule Cron表达式

	Schedule *string `json:"Schedule,omitempty" name:"Schedule"`
	// Suspend 是否停用

	Suspend *bool `json:"Suspend,omitempty" name:"Suspend"`
	// Datetime 日期, 用于前端展示无业务逻辑

	Datetime *string `json:"Datetime,omitempty" name:"Datetime"`
	// MetadataName 巡检计划主键

	MetadataName *string `json:"MetadataName,omitempty" name:"MetadataName"`
	// ID 更新的主键ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// Creator

	Creator *string `json:"Creator,omitempty" name:"Creator"`
}

type Rule struct {
	// Mode 规则类型

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// Key 规则Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// Value 规则Value

	Value *string `json:"Value,omitempty" name:"Value"`
	// Operator 规则运算符

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// Remark 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// External 扩展字段

	External *External `json:"External,omitempty" name:"External"`
}

type ListCronJobRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *QueryRequest `json:"Body,omitempty" name:"Body"`
}

func (r *ListCronJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCronJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOperationRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InspectionItem struct {
	// Name 巡检项名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Remark 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// Type 巡检项类型归属

	Type *Type `json:"Type,omitempty" name:"Type"`
	// Alert 规则详情

	Alerts *Alerts `json:"Alerts,omitempty" name:"Alerts"`
	// RuleMode 规则数据源获取方式

	RuleMode *string `json:"RuleMode,omitempty" name:"RuleMode"`
	// Steps 流程step

	Steps []*Step `json:"Steps,omitempty" name:"Steps"`
	// MetadataName 巡检项主键

	MetadataName *string `json:"MetadataName,omitempty" name:"MetadataName"`
	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// Creator

	Creator *string `json:"Creator,omitempty" name:"Creator"`
}

type Type struct {
	// Product 产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// InspectionType 巡检类型

	InspectionType *string `json:"InspectionType,omitempty" name:"InspectionType"`
	// Internal 是否内置巡检项

	Internal *bool `json:"Internal,omitempty" name:"Internal"`
}

type body struct {
	// ID 主键ID

	ID *string `json:"ID,omitempty" name:"ID"`
}

type GeneralRequest struct {
	// ID

	ID *string `json:"ID,omitempty" name:"ID"`
	// 扩展字段

	External *External `json:"External,omitempty" name:"External"`
	// Type 类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type CheckImageURLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckImageURLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckImageURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCSPDownloadURLRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *GeneralRequest `json:"Body,omitempty" name:"Body"`
}

func (r *CreateCSPDownloadURLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCSPDownloadURLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateInspectionItemRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *InspectionItem `json:"Body,omitempty" name:"Body"`
}

func (r *UpdateInspectionItemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateInspectionItemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInspectionItemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetInspectionItemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInspectionItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckImageURLRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *GeneralRequest `json:"Body,omitempty" name:"Body"`
}

func (r *CheckImageURLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckImageURLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInspectionItemRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *InspectionItem `json:"Body,omitempty" name:"Body"`
}

func (r *CreateInspectionItemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInspectionItemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInspectionItemRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *GeneralRequest `json:"Body,omitempty" name:"Body"`
}

func (r *DeleteInspectionItemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInspectionItemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListInspectionItemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListInspectionItemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInspectionItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Step struct {
	// Remark 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// Timeout 超时时间

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	// ScriptType 脚本类型

	ScriptType *string `json:"ScriptType,omitempty" name:"ScriptType"`
	// Input 输入方式

	Input *Input `json:"Input,omitempty" name:"Input"`
	// Command 执行命令

	Command *string `json:"Command,omitempty" name:"Command"`
	// InspectionType 巡检类型

	InspectionType *string `json:"InspectionType,omitempty" name:"InspectionType"`
}

type GetReportHistoryRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *QueryRequest `json:"Body,omitempty" name:"Body"`
}

func (r *GetReportHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetReportHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCMDBProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCMDBProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCMDBProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopCronJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopCronJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopCronJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCronJobRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *GeneralRequest `json:"Body,omitempty" name:"Body"`
}

func (r *DeleteCronJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCronJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCronJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCronJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCronJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartCronJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartCronJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartCronJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListJobHistoryRequest struct {
	*tchttp.BaseRequest

	// 请求参数，Limit、Offset为预留参数，暂不起效

	Body *QueryRequest `json:"Body,omitempty" name:"Body"`
}

func (r *ListJobHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListJobHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCronJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCronJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCronJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type External struct {
	// 脚本内容

	ScriptContent *string `json:"ScriptContent,omitempty" name:"ScriptContent"`
	// 请求方法

	HTTPMethod *string `json:"HTTPMethod,omitempty" name:"HTTPMethod"`
	// 请求地址

	HTTPURL *string `json:"HTTPURL,omitempty" name:"HTTPURL"`
	// 请求体

	HTTPBody *string `json:"HTTPBody,omitempty" name:"HTTPBody"`
	// HTTPCookie

	HTTPCookie *string `json:"HTTPCookie,omitempty" name:"HTTPCookie"`
	// 目标IP列表

	IPIPList []*string `json:"IPIPList,omitempty" name:"IPIPList"`
	// 端口列表

	IPPortList []*int64 `json:"IPPortList,omitempty" name:"IPPortList"`
	// 巡检协议

	IPProtocol *string `json:"IPProtocol,omitempty" name:"IPProtocol"`
	// 脚本语言

	ScriptType *string `json:"ScriptType,omitempty" name:"ScriptType"`
	// 默认值

	DefaultValue *int64 `json:"DefaultValue,omitempty" name:"DefaultValue"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// HTTPHeader，JSON串

	HTTPHeader []*HTTPHeader `json:"HTTPHeader,omitempty" name:"HTTPHeader"`
	// CMDBIPIPList CMDB的IP列表

	CMDBIPIPList []*string `json:"CMDBIPIPList,omitempty" name:"CMDBIPIPList"`
}

type DeleteCronJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCronJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCronJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCMDBProductRequest struct {
	*tchttp.BaseRequest

	// 请求参数，Limit、Offset为预留参数，暂不起效

	Body *QueryRequest `json:"Body,omitempty" name:"Body"`
}

func (r *ListCMDBProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCMDBProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCronJobRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *CronJob `json:"Body,omitempty" name:"Body"`
}

func (r *UpdateCronJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCronJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateInspectionItemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateInspectionItemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateInspectionItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListJobRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *QueryRequest `json:"Body,omitempty" name:"Body"`
}

func (r *ListJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationRecordRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *QueryRequest `json:"Body,omitempty" name:"Body"`
}

func (r *ListOperationRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetKeepTimeRequest struct {
	*tchttp.BaseRequest

	// Body 响应体内容

	Body *body `json:"Body,omitempty" name:"Body"`
}

func (r *GetKeepTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetKeepTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunCronJobRequest struct {
	*tchttp.BaseRequest

	// 请求体

	Body *GeneralRequest `json:"Body,omitempty" name:"Body"`
}

func (r *RunCronJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunCronJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartCronJobRequest struct {
	*tchttp.BaseRequest

	// 请求体

	Body *GeneralRequest `json:"Body,omitempty" name:"Body"`
}

func (r *StartCronJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartCronJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Alert struct {
	// Suggest 建议

	Suggest *string `json:"Suggest,omitempty" name:"Suggest"`
	// Rules 规则列表

	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`
	// Operators 规则运算符

	Operators []*string `json:"Operators,omitempty" name:"Operators"`
	// 扩展字段

	External *External `json:"External,omitempty" name:"External"`
}

type ListCustomProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCustomProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCustomProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HTTPHeader struct {
	// Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// Value

	Value []*string `json:"Value,omitempty" name:"Value"`
}

type Notification struct {
	// Enable 是否开启

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// Level 告警级别，支持Info、Warn、Critical

	Level *string `json:"Level,omitempty" name:"Level"`
	// Methods 通知方式，支持Message（手机短信）、Mail（邮件）、Wechat（微信）、Wecom（企业微信）

	Methods []*string `json:"Methods,omitempty" name:"Methods"`
}

type CreateCronJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCronJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCronJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInspectionItemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInspectionItemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInspectionItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopJobRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *GeneralRequest `json:"Body,omitempty" name:"Body"`
}

func (r *StopJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCronJobHistoryRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *QueryRequest `json:"Body,omitempty" name:"Body"`
}

func (r *ListCronJobHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCronJobHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopCronJobRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *GeneralRequest `json:"Body,omitempty" name:"Body"`
}

func (r *StopCronJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopCronJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeepTime struct {
	// 天数

	Time *int64 `json:"Time,omitempty" name:"Time"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type CheckResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *string `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *GeneralRequest `json:"Body,omitempty" name:"Body"`
}

func (r *CheckRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigureKeepTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConfigureKeepTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfigureKeepTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCronJobRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *GeneralRequest `json:"Body,omitempty" name:"Body"`
}

func (r *GetCronJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCronJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigureKeepTimeRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *KeepTime `json:"Body,omitempty" name:"Body"`
}

func (r *ConfigureKeepTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfigureKeepTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCronJobRequest struct {
	*tchttp.BaseRequest

	// CreateCronJob 请求体

	Body *CronJob `json:"Body,omitempty" name:"Body"`
}

func (r *CreateCronJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCronJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Alerts struct {
	// Info

	Info *Alert `json:"Info,omitempty" name:"Info"`
	// Cratical

	Critical *Alert `json:"Critical,omitempty" name:"Critical"`
	// Warn

	Warn *Alert `json:"Warn,omitempty" name:"Warn"`
}

type GetKeepTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetKeepTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetKeepTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetReportRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *GeneralRequest `json:"Body,omitempty" name:"Body"`
}

func (r *GetReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetReportHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetReportHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetReportHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCronJobHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCronJobHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCronJobHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCronJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCronJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCronJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInspectionItemRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *GeneralRequest `json:"Body,omitempty" name:"Body"`
}

func (r *GetInspectionItemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInspectionItemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCustomProductRequest struct {
	*tchttp.BaseRequest

	// 请求参数，Limit、Offset为预留参数，暂不起效

	Body *QueryRequest `json:"Body,omitempty" name:"Body"`
}

func (r *ListCustomProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCustomProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInspectionItemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInspectionItemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInspectionItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunCronJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunCronJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunCronJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryRequest struct {
	// Limit 获取数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Offset 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Sort defines how to sort the service instances. It's in the format of '[+-]<field>', '+' means ascending order,'-' means descending order. Here '+' is optional. For example, '+instanceID', '-instanceID'

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// Query string to query resources. Supported query patterns are "exact match(k=v)", "fuzzy match(k=\~v)", "range(k=[min\~max])", "list with union releationship(k={v1 v2 v3})". All of these query patterns should be put in the query string "q=xxx" and splitted by ",". e.g. q=k1=v1,k2=\~v2,k3=[min\~max]

	Query *string `json:"Query,omitempty" name:"Query"`
}

type ListJobHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListJobHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListJobHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Receiver struct {
	// Username

	Username *string `json:"Username,omitempty" name:"Username"`
}

type CreateCSPDownloadURLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Body 响应体内容

		Body *body `json:"Body,omitempty" name:"Body"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCSPDownloadURLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCSPDownloadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListInspectionItemRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *QueryRequest `json:"Body,omitempty" name:"Body"`
}

func (r *ListInspectionItemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInspectionItemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListReportRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Body *QueryRequest `json:"Body,omitempty" name:"Body"`
}

func (r *ListReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Input struct {
	// Type 方式

	Type *string `json:"Type,omitempty" name:"Type"`
	// PackageName 包名称

	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`
	// PackageVersion 包版本

	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`
	// ImageURL 镜像URL

	ImageURL *string `json:"ImageURL,omitempty" name:"ImageURL"`
	// 扩展字段

	External *External `json:"External,omitempty" name:"External"`
}
