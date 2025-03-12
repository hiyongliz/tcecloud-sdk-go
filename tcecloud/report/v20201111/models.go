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

package v20201111

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type ListCatetoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 报表分类列表

		Categories []*Category `json:"Categories,omitempty" name:"Categories"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCatetoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCatetoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTemplateRequest struct {
	*tchttp.BaseRequest

	// 模版元数据

	Template *TemplateExt `json:"Template,omitempty" name:"Template"`
}

func (r *UpdateTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Category struct {
	// id

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 报表分类(用户自定义)

	Category *string `json:"Category,omitempty" name:"Category"`
}

type DWF struct {
	// 采集接口返回的:字段名称

	SourceName *string `json:"SourceName,omitempty" name:"SourceName"`
	// 数据仓库:可读字段名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 前端展示使用

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// metric 扩展前的名称

	Base *string `json:"Base,omitempty" name:"Base"`
	// 数仓字段类型, 用于业务特殊逻辑. 例如: base,instance,label,metric

	Type *string `json:"Type,omitempty" name:"Type"`
	// 数据仓库:字段名称

	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`
	// 数据仓库:字段类型

	FieldType *string `json:"FieldType,omitempty" name:"FieldType"`
	// 数据仓库:字段操作(仅 metric 类型字段使用)

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 监控查询:query(每小时)

	Query *string `json:"Query,omitempty" name:"Query"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type UpdateSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// 模版编号

		TemplateID *int64 `json:"TemplateID,omitempty" name:"TemplateID"`
		// 邮件标题

		Title *string `json:"Title,omitempty" name:"Title"`
		// 接收者, 包含: 用户/用户组/邮箱列表

		Receiver *string `json:"Receiver,omitempty" name:"Receiver"`
		// 抄送, 包含: 用户/用户组/邮箱列表

		Ccreceiver *string `json:"Ccreceiver,omitempty" name:"Ccreceiver"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTemplateRequest struct {
	*tchttp.BaseRequest

	// 偏移位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 本次返回个数限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	Order *SortField `json:"Order,omitempty" name:"Order"`
}

func (r *ListTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ID

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// 模版编号

		TemplateID *int64 `json:"TemplateID,omitempty" name:"TemplateID"`
		// 邮件标题

		Title *string `json:"Title,omitempty" name:"Title"`
		// 接收者, 包含: 用户/用户组/邮箱列表

		Receiver *Receiver `json:"Receiver,omitempty" name:"Receiver"`
		// 抄送, 包含: 用户/用户组/邮箱列表

		Ccreceiver *Receiver `json:"Ccreceiver,omitempty" name:"Ccreceiver"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenderTemplateRequest struct {
	*tchttp.BaseRequest

	// 模版元数据

	Template *TemplateExt `json:"Template,omitempty" name:"Template"`
}

func (r *RenderTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenderTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Column struct {
	// 字段名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段类型. 例如: VARCHAR(20)

	Type *string `json:"Type,omitempty" name:"Type"`
	// 是否为空

	Nullable *bool `json:"Nullable,omitempty" name:"Nullable"`
	// 字段默认值. 例如: DEFUALT '' 或者 DEFAULT 1.0

	Default *string `json:"Default,omitempty" name:"Default"`
	// 字段备注信息. 例如: COMMENT 'xxx'

	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type UserTuple struct {
	// UIN

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// cam 用户数据库自增ID

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 用户登陆名称

	UserID *string `json:"UserID,omitempty" name:"UserID"`
}

type QueryTemplateRequest struct {
	*tchttp.BaseRequest

	// 模版编号

	ID *uint64 `json:"ID,omitempty" name:"ID"`
}

func (r *QueryTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCategoryRequest struct {
	*tchttp.BaseRequest

	// 分类id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 模板列表

		Templates []*TemplateExt `json:"Templates,omitempty" name:"Templates"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenameCatetoryRequest struct {
	*tchttp.BaseRequest

	// 分类编号

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 分类名称

	Category *string `json:"Category,omitempty" name:"Category"`
}

func (r *RenameCatetoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameCatetoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChartAttribute struct {
	// 保留字段

	PlaceHolder *string `json:"PlaceHolder,omitempty" name:"PlaceHolder"`
}

type TemplateTableExt struct {
	// 表格编号

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 表格名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 表格展示样式. 固定为: table

	ShowType *string `json:"ShowType,omitempty" name:"ShowType"`
	// 模版编号

	TemplateID *int64 `json:"TemplateID,omitempty" name:"TemplateID"`
	// 表格在模版中的排序

	SeqInTemplate *int64 `json:"SeqInTemplate,omitempty" name:"SeqInTemplate"`
	// 数据仓库:表格名称

	DwTableName *string `json:"DwTableName,omitempty" name:"DwTableName"`
	// 字段列表

	Fields []*TemplateFieldExt `json:"Fields,omitempty" name:"Fields"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 图表组名称

	ChartGroupName *string `json:"ChartGroupName,omitempty" name:"ChartGroupName"`
	// 图表属性

	ChartAttribute *ChartAttribute `json:"ChartAttribute,omitempty" name:"ChartAttribute"`
	// 图表属性-后续删除

	DwChartAttribute *string `json:"DwChartAttribute,omitempty" name:"DwChartAttribute"`
	// 分类

	Category *MetaCategory `json:"Category,omitempty" name:"Category"`
}

type MetaCategory struct {
	// 产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type CreateSubscribeRequest struct {
	*tchttp.BaseRequest

	// 模版编号

	TemplateID *uint64 `json:"TemplateID,omitempty" name:"TemplateID"`
	// 邮件标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 邮件接收者

	Receiver *Receiver `json:"Receiver,omitempty" name:"Receiver"`
	// 邮件抄送

	Ccreceiver *Receiver `json:"Ccreceiver,omitempty" name:"Ccreceiver"`
}

func (r *CreateSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTemplateRequest struct {
	*tchttp.BaseRequest

	// 模版编号

	ID *uint64 `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMetricSourceRequest struct {
	*tchttp.BaseRequest

	// 产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 分类

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *ListMetricSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMetricSourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySubscribeRequest struct {
	*tchttp.BaseRequest

	// 模版编号

	TemplateID *uint64 `json:"TemplateID,omitempty" name:"TemplateID"`
}

func (r *QuerySubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenderTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板

		Template *TemplateExt `json:"Template,omitempty" name:"Template"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenderTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenderTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SortField struct {
	// 排序对应字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// 排序方式，ASC: 升序， DESC: 降序

	SortWay *string `json:"SortWay,omitempty" name:"SortWay"`
}

type TemplateFieldExt struct {
	// 数据库ID:字段编号

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 自定义字段名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 模版编号

	TemplateID *int64 `json:"TemplateID,omitempty" name:"TemplateID"`
	// 表格编号

	TableID *int64 `json:"TableID,omitempty" name:"TableID"`
	// 字段在表格中的编号

	SeqInTable *int64 `json:"SeqInTable,omitempty" name:"SeqInTable"`
	// 字段过滤选项:(存储使用)

	DwOperation *string `json:"DwOperation,omitempty" name:"DwOperation"`
	// 数据仓库:表格名称

	DwTableName *string `json:"DwTableName,omitempty" name:"DwTableName"`
	// 数据仓库:字段名称

	DwFieldName *string `json:"DwFieldName,omitempty" name:"DwFieldName"`
	// 字段过滤选项

	Operation *Operation `json:"Operation,omitempty" name:"Operation"`
	// 字段单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type UnitConversion struct {
	// 计算方式，0: 空，1：乘，2：除

	CalculationMethod *string `json:"CalculationMethod,omitempty" name:"CalculationMethod"`
	// 具体运算数值

	Scale *float64 `json:"Scale,omitempty" name:"Scale"`
	// 目标单位

	TargetUnit *string `json:"TargetUnit,omitempty" name:"TargetUnit"`
	// 数据精度，1~5

	DataPrecision *uint64 `json:"DataPrecision,omitempty" name:"DataPrecision"`
}

type AddCatetoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分类

		Category *Category `json:"Category,omitempty" name:"Category"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddCatetoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddCatetoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListInstanceCategoryRequest struct {
	*tchttp.BaseRequest
}

func (r *ListInstanceCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInstanceCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListInstanceCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 报表元数据分类

		Categories []*MetaCategory `json:"Categories,omitempty" name:"Categories"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListInstanceCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInstanceCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMetricSourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 字段

		Fields []*Field `json:"Fields,omitempty" name:"Fields"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMetricSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMetricSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Field struct {
	// 数据仓库:表格信息

	DWT *DWT `json:"DWT,omitempty" name:"DWT"`
	// 数据仓库:字段信息

	DWF *DWF `json:"DWF,omitempty" name:"DWF"`
}

type Filter struct {
	// 字段名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段值列表

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DeleteCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetReportChartUrlRequest struct {
	*tchttp.BaseRequest

	// 报表id

	ID *int64 `json:"ID,omitempty" name:"ID"`
}

func (r *GetReportChartUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetReportChartUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenameCatetoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分类列表

		Category *Category `json:"Category,omitempty" name:"Category"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenameCatetoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameCatetoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Operation struct {
	// 筛选: 前N条

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选: 区间 (开始/结束)

	Between []*float64 `json:"Between,omitempty" name:"Between"`
	// 排序: asc / desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 单对象运算: min / max / avg

	SingleObject *string `json:"SingleObject,omitempty" name:"SingleObject"`
	// 分组运算: min / max / avg / sum / count

	Group *string `json:"Group,omitempty" name:"Group"`
	// 分组标签列表

	GroupLabel []*string `json:"GroupLabel,omitempty" name:"GroupLabel"`
	// 标签筛选操作

	ObjectSelector []*ObjectOperation `json:"ObjectSelector,omitempty" name:"ObjectSelector"`
	// 对象标识列表

	ObjectIdent []*string `json:"ObjectIdent,omitempty" name:"ObjectIdent"`
	// 是否展示环比，1：是, 2：否

	ShowPOPR *uint64 `json:"ShowPOPR,omitempty" name:"ShowPOPR"`
	// 单位转换

	UnitConversion *UnitConversion `json:"UnitConversion,omitempty" name:"UnitConversion"`
}

type TTemplateSubscribe struct {
	// id

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 模版编号

	TemplateID *string `json:"TemplateID,omitempty" name:"TemplateID"`
	// 邮件标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 接收者, 包含: 用户/用户组/邮箱列表

	Receiver *string `json:"Receiver,omitempty" name:"Receiver"`
	// 抄送, 包含: 用户/用户组/邮箱列表

	Ccreceiver *string `json:"Ccreceiver,omitempty" name:"Ccreceiver"`
}

type QueryTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板

		Template *TemplateExt `json:"Template,omitempty" name:"Template"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSubscribeRequest struct {
	*tchttp.BaseRequest

	// 订阅关系数据库内编号

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 模版编号

	TemplateID *uint64 `json:"TemplateID,omitempty" name:"TemplateID"`
	// 邮件标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 邮件接受人

	Receiver *Receiver `json:"Receiver,omitempty" name:"Receiver"`
	// 邮件抄送

	Ccreceiver *Receiver `json:"Ccreceiver,omitempty" name:"Ccreceiver"`
}

func (r *UpdateSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ObjectOperation struct {
	// 标签对象名称

	ObjectName *string `json:"ObjectName,omitempty" name:"ObjectName"`
	// 操作符，目前固定为等于： =

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 标签对象对应的值

	Value []*string `json:"Value,omitempty" name:"Value"`
}

type Receiver struct {
	// 运营端登陆用户名

	UserList []*UserTuple `json:"UserList,omitempty" name:"UserList"`
	// 用户组编号

	GroupList []*GroupTuple `json:"GroupList,omitempty" name:"GroupList"`
	// 邮箱地址

	EmailList []*string `json:"EmailList,omitempty" name:"EmailList"`
}

type AddCatetoryRequest struct {
	*tchttp.BaseRequest

	// 分类名称

	Category *string `json:"Category,omitempty" name:"Category"`
}

func (r *AddCatetoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddCatetoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTemplateRequest struct {
	*tchttp.BaseRequest

	// 模版元数据

	Template *TemplateExt `json:"Template,omitempty" name:"Template"`
}

func (r *CreateTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListInstanceSourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 字段

		Fields []*Field `json:"Fields,omitempty" name:"Fields"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListInstanceSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInstanceSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMetricCategoryRequest struct {
	*tchttp.BaseRequest
}

func (r *ListMetricCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMetricCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCatetoryRequest struct {
	*tchttp.BaseRequest

	// 参数过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListCatetoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCatetoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListReportRequest struct {
	*tchttp.BaseRequest

	// 偏移位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 本次返回个数限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetReportChartUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 临时下载url

		TmpDownLoadUrl *string `json:"TmpDownLoadUrl,omitempty" name:"TmpDownLoadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetReportChartUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetReportChartUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListInstanceSourceRequest struct {
	*tchttp.BaseRequest

	// 产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 分类

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *ListInstanceSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInstanceSourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryReportRequest struct {
	*tchttp.BaseRequest

	// 模版编号

	ID *uint64 `json:"ID,omitempty" name:"ID"`
}

func (r *QueryReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DWT struct {
	// 数据仓库:表格名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type GroupTuple struct {
	// 用户组编号

	GID *uint64 `json:"GID,omitempty" name:"GID"`
	// 用户组名称

	GName *string `json:"GName,omitempty" name:"GName"`
}

type CreateSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// id

		ID *int64 `json:"ID,omitempty" name:"ID"`
		// 模版编号

		TemplateID *int64 `json:"TemplateID,omitempty" name:"TemplateID"`
		// 邮件标题

		Title *string `json:"Title,omitempty" name:"Title"`
		// 接收者, 包含: 用户/用户组/邮箱列表

		Receiver *string `json:"Receiver,omitempty" name:"Receiver"`
		// 抄送, 包含: 用户/用户组/邮箱列表

		Ccreceiver *string `json:"Ccreceiver,omitempty" name:"Ccreceiver"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMetricCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 报表元数据分类

		Categories []*MetaCategory `json:"Categories,omitempty" name:"Categories"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMetricCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMetricCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 历史报表列表

		Reports []*TReport `json:"Reports,omitempty" name:"Reports"`
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

type QueryReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 报表内容

		Report *TReport `json:"Report,omitempty" name:"Report"`
		// 模板

		Template *TemplateExt `json:"Template,omitempty" name:"Template"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateExt struct {
	// 模版编号

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 模版名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 模版类型: 日报:daily, 周报:weekly, 月报:monthly

	Type *string `json:"Type,omitempty" name:"Type"`
	// 模版分类(用户自定义)

	CategoryID *uint64 `json:"CategoryID,omitempty" name:"CategoryID"`
	// 模版负责人

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 模版创建时间

	Createtime *string `json:"Createtime,omitempty" name:"Createtime"`
	// 模版更新时间

	Updatetime *string `json:"Updatetime,omitempty" name:"Updatetime"`
	// 报表触发时间

	Triggertime *string `json:"Triggertime,omitempty" name:"Triggertime"`
	// 表格

	Tables []*TemplateTableExt `json:"Tables,omitempty" name:"Tables"`
	// 触发周期

	TriggerPeriod *int64 `json:"TriggerPeriod,omitempty" name:"TriggerPeriod"`
	// 版本

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 更新用户

	UpdateUser *string `json:"UpdateUser,omitempty" name:"UpdateUser"`
	// 创建用户

	CreateUser *string `json:"CreateUser,omitempty" name:"CreateUser"`
}

type TReport struct {
	// id

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 报表名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 报表类型. 日报:daily, 周报:weekly, 月报:monthly, 季度报:quarter

	Type *string `json:"Type,omitempty" name:"Type"`
	// 报表负责人

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 报表生成时间

	Createtime *string `json:"Createtime,omitempty" name:"Createtime"`
	// 模版编号

	TemplateID *int64 `json:"TemplateID,omitempty" name:"TemplateID"`
	// 模版分类(用户自定义)

	CategoryID *int64 `json:"CategoryID,omitempty" name:"CategoryID"`
	// 数据时间跨度,与 type 字段值绑定

	Timerangebegin *string `json:"Timerangebegin,omitempty" name:"Timerangebegin"`
	// 数据时间跨度,与 type 字段值绑定

	Timerangeend *string `json:"Timerangeend,omitempty" name:"Timerangeend"`
	// 模版快照:版本

	Snapshotversion *uint64 `json:"Snapshotversion,omitempty" name:"Snapshotversion"`
	// 模版快照:内容

	Snapshotcontent *string `json:"Snapshotcontent,omitempty" name:"Snapshotcontent"`
	// 邮件标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 接收者, 包含: 用户/用户组/邮箱列表

	Receiver *string `json:"Receiver,omitempty" name:"Receiver"`
	// 抄送, 包含: 用户/用户组/邮箱列表

	Ccreceiver *string `json:"Ccreceiver,omitempty" name:"Ccreceiver"`
	// 可视化图表存储KEY

	ChartStorageKey *string `json:"ChartStorageKey,omitempty" name:"ChartStorageKey"`
}
