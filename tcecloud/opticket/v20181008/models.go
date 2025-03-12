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

package v20181008

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type StaffDismissRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 客服ID

	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`
}

func (r *StaffDismissRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffDismissRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportDatas struct {
	// 状态码

	Code *uint64 `json:"Code,omitempty" name:"Code"`
	// 提示信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 导出的文件链接

	Data *ExportUrl `json:"Data,omitempty" name:"Data"`
}

type GetCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品列表

		Data *DescribeCategorysData `json:"Data,omitempty" name:"Data"`
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateStaffResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码

		Data *Code `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateStaffResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateStaffResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTicketResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 工单ID

		Data *Code `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTicketResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTicketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateScheduleRequest struct {
	*tchttp.BaseRequest

	// 客服ID

	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`
	// 排班日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 当班开始时间

	TimeStart *string `json:"TimeStart,omitempty" name:"TimeStart"`
	// 当班结束时间

	TimeEnd *string `json:"TimeEnd,omitempty" name:"TimeEnd"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *UpdateScheduleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateScheduleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScheduleStaff struct {
	// 客服ID

	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`
	// 客服数目

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
	// 是否在线

	IsOnline *uint64 `json:"IsOnline,omitempty" name:"IsOnline"`
}

type CheckBeforeDeleteStaffStatusRequest struct {
	*tchttp.BaseRequest

	// 客服ID

	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *CheckBeforeDeleteStaffStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckBeforeDeleteStaffStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTicketsData struct {
	// 工单列表

	Data []*BasicTicket `json:"Data,omitempty" name:"Data"`
	// 工单总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type ObjectId struct {
	// 对象Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type UpdateCategoryRequest struct {
	*tchttp.BaseRequest

	// 产品级别分类

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 产品名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品展示顺序权重

	Weights *uint64 `json:"Weights,omitempty" name:"Weights"`
	// 问题描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 产品图标

	ImgUrl *string `json:"ImgUrl,omitempty" name:"ImgUrl"`
	// 产品ID

	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 父节点ID

	ParentId *uint64 `json:"ParentId,omitempty" name:"ParentId"`
	// 是否常用

	Common *uint64 `json:"Common,omitempty" name:"Common"`
}

func (r *UpdateCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest

	// 客服组ID数组

	GroupId []*uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportScheduleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出的数据结构

		Data *ExportDatas `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportScheduleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGroupInfoRequest struct {
	*tchttp.BaseRequest

	// 客服组ID

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 显示客服分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 显示客服每页记录数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *GetGroupInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGroupInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateFieldRequest struct {
	*tchttp.BaseRequest

	// 产品ID

	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 字段名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段类型

	Types *string `json:"Types,omitempty" name:"Types"`
	// 显示顺序权重

	Weights *uint64 `json:"Weights,omitempty" name:"Weights"`
	// 字段键值

	FieldKv *string `json:"FieldKv,omitempty" name:"FieldKv"`
	// 文案

	Copywriter *string `json:"Copywriter,omitempty" name:"Copywriter"`
	// 是否必填

	Required *uint64 `json:"Required,omitempty" name:"Required"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 字段id

	FieldId *string `json:"FieldId,omitempty" name:"FieldId"`
}

func (r *UpdateFieldRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateFieldRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BasicFirstCategory struct {
	// 一级分类Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 一级分类信息

	Cell *FirstCategoryCell `json:"Cell,omitempty" name:"Cell"`
	// 子节点分类数组

	Children []*ChildrenCategory `json:"Children,omitempty" name:"Children"`
}

type DescribeTicketRequest struct {
	*tchttp.BaseRequest

	// 工单ID

	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *DescribeTicketRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTicketRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTicketCategoryRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *GetTicketCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTicketCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTicketCategoryByLevelRequest struct {
	*tchttp.BaseRequest

	// 分类ID

	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *GetTicketCategoryByLevelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTicketCategoryByLevelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGroupsRequest struct {
	*tchttp.BaseRequest

	// 客服组过滤条件

	Filters []*GroupFilter `json:"Filters,omitempty" name:"Filters"`
	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页记录数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *GetGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest

	// 客服组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *CreateGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFieldRequest struct {
	*tchttp.BaseRequest

	// 产品ID

	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *GetFieldRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFieldRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SucCode struct {
	// 状态码

	Code *uint64 `json:"Code,omitempty" name:"Code"`
	// 数据

	Data []*int64 `json:"Data,omitempty" name:"Data"`
	// 提示信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type GetCosParamsByBucketResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// COS密钥

		Data *CosParams `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCosParamsByBucketResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCosParamsByBucketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMessagesData struct {
	// 消息列表

	Data []*Message `json:"Data,omitempty" name:"Data"`
	// 总记录数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type GetStaffStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 在线状态客服信息查询结果

		Data *DescribeOnlineStaffsData `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStaffStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStaffStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecondCategoryCell struct {
	// 分类名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分类描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 分类是否常用

	Common *uint64 `json:"Common,omitempty" name:"Common"`
	// 分类显示权重

	Weights *uint64 `json:"Weights,omitempty" name:"Weights"`
	// 分类图标

	ImgUrl *string `json:"ImgUrl,omitempty" name:"ImgUrl"`
}

type GetScheduleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 排班查询结果

		Data *DescribeSchedulesData `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetScheduleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStaffStatusRequest struct {
	*tchttp.BaseRequest

	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 排班客服

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
	// 是否在线

	IsOnline *uint64 `json:"IsOnline,omitempty" name:"IsOnline"`
	// 是否当班

	IsOnDuty *uint64 `json:"IsOnDuty,omitempty" name:"IsOnDuty"`
}

func (r *GetStaffStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStaffStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码

		Data *Code `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStaffsRequest struct {
	*tchttp.BaseRequest

	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页记录数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 请求参数

	Filters []*GetStaffParams `json:"Filters,omitempty" name:"Filters"`
}

func (r *GetStaffsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStaffsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffSignOutRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *StaffSignOutRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffSignOutRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功返回后的分类ID

		Data []*CategoryIdItem `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCosBucketResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// cos数据

		Data *CosBucketData `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCosBucketResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCosBucketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStaffsData struct {
	// 客服总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
	// 客服数据

	Data []*BasicStaff `json:"Data,omitempty" name:"Data"`
}

type BasicThirdLevel struct {
	// 三级分类ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 三级分类信息

	Cell *ThirdCategoryCell `json:"Cell,omitempty" name:"Cell"`
}

type CreateCategoryRequest struct {
	*tchttp.BaseRequest

	// 产品级别分类

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 产品名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品展示顺序权重

	Weights *uint64 `json:"Weights,omitempty" name:"Weights"`
	// 父节点ID

	ParentId *uint64 `json:"ParentId,omitempty" name:"ParentId"`
	// 问题描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 产品图标

	ImgUrl *string `json:"ImgUrl,omitempty" name:"ImgUrl"`
	// 是否常用

	Common *uint64 `json:"Common,omitempty" name:"Common"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *CreateCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTicketResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 工单详情

		Data *Ticket `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTicketResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTicketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetScheduleParams struct {
	// 查询开始日期

	Start *string `json:"Start,omitempty" name:"Start"`
	// 查询截至日期

	End *string `json:"End,omitempty" name:"End"`
	// 客服数组

	Users []*string `json:"Users,omitempty" name:"Users"`
}

type GetServerTimeRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *GetServerTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServerTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffReceiveTicketResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Data *Code `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StaffReceiveTicketResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffReceiveTicketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateFieldResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功后的返回json格式

		Data *SucCode `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateFieldResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateFieldResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddGroupStaffRequest struct {
	*tchttp.BaseRequest

	// 客服组ID

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 客服ID

	StaffId []*string `json:"StaffId,omitempty" name:"StaffId"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *AddGroupStaffRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddGroupStaffRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportUrl struct {
	// 导出的cos文件链接

	Url *string `json:"Url,omitempty" name:"Url"`
}

type GetServerTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServerTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServerTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffAssignTicketRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 工单ID

	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
	// 客服ID

	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`
}

func (r *StaffAssignTicketRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffAssignTicketRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Assignment struct {
	// 派单记录ID

	AssignmentId *uint64 `json:"AssignmentId,omitempty" name:"AssignmentId"`
	// 分派工单的人

	Assigner *string `json:"Assigner,omitempty" name:"Assigner"`
	// 接收工单的人

	Assignee *string `json:"Assignee,omitempty" name:"Assignee"`
	// 工单ID

	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
	// 派单时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 派单备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type BasicTicket struct {
	// 工单ID

	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
	// 工单一级分类

	TypeLevel1 *uint64 `json:"TypeLevel1,omitempty" name:"TypeLevel1"`
	// 工单二级分类

	TypeLevel2 *uint64 `json:"TypeLevel2,omitempty" name:"TypeLevel2"`
	// 工单三级分类

	TypeLevel3 *uint64 `json:"TypeLevel3,omitempty" name:"TypeLevel3"`
	// 租户端工单状态

	AskerStatus *uint64 `json:"AskerStatus,omitempty" name:"AskerStatus"`
	// 运营端工单状态

	AnswererStatus *uint64 `json:"AnswererStatus,omitempty" name:"AnswererStatus"`
	// 优先级

	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
	// 工单内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 机密信息

	SecretContent *string `json:"SecretContent,omitempty" name:"SecretContent"`
	// 客户ID

	CustomerId *string `json:"CustomerId,omitempty" name:"CustomerId"`
	// 工单当前处理客服

	CurrentStaff *string `json:"CurrentStaff,omitempty" name:"CurrentStaff"`
	// 结单人

	Finisher *string `json:"Finisher,omitempty" name:"Finisher"`
	// 工单创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 结单时间

	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
	// 撤单时间

	CancelTime *string `json:"CancelTime,omitempty" name:"CancelTime"`
}

type StaffSignInRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *StaffSignInRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffSignInRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BasicSchedule struct {
	// 排班Id

	ScheduleId *uint64 `json:"ScheduleId,omitempty" name:"ScheduleId"`
	// 排班日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 客服ID

	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`
	// 当班开始时间

	TimeStart *string `json:"TimeStart,omitempty" name:"TimeStart"`
	// 当班结束时间

	TimeEnd *string `json:"TimeEnd,omitempty" name:"TimeEnd"`
	// 当班时长

	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
}

type CategoryIdItem struct {
	// 分类ID

	CategoryId *string `json:"CategoryId,omitempty" name:"CategoryId"`
}

type PersonSchedule struct {
	// 客服ID

	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`
	// 客服中文名

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
	// 是否在线

	IsOnline *uint64 `json:"IsOnline,omitempty" name:"IsOnline"`
	// 客服的所有排班

	Schedules []*BasicSchedule `json:"Schedules,omitempty" name:"Schedules"`
}

type GetStaffsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 客服列表

		Data *DescribeStaffsData `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStaffsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStaffsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功返回的json格式

		Data *SucCode `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MessageFilter struct {
	// 消息类型

	Type []*uint64 `json:"Type,omitempty" name:"Type"`
	// 工单ID

	TicketId []*uint64 `json:"TicketId,omitempty" name:"TicketId"`
	// 接收人

	Receiver []*uint64 `json:"Receiver,omitempty" name:"Receiver"`
	// 消息触发人

	Trigger []*uint64 `json:"Trigger,omitempty" name:"Trigger"`
	// 是否已读

	IsRead *uint64 `json:"IsRead,omitempty" name:"IsRead"`
	// 按创建时间排序 asc或desc

	OrderByCreateTime *string `json:"OrderByCreateTime,omitempty" name:"OrderByCreateTime"`
	// 按是否已读排序 asc或desc

	OrderByIsRead *string `json:"OrderByIsRead,omitempty" name:"OrderByIsRead"`
	// 创建时间起点

	CreateTimeStart *string `json:"CreateTimeStart,omitempty" name:"CreateTimeStart"`
	// 创建时间终点

	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" name:"CreateTimeEnd"`
}

type GetWorkbenchDataRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *GetWorkbenchDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkbenchDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BasicField struct {
	// 表单字段Id

	FieldId *uint64 `json:"FieldId,omitempty" name:"FieldId"`
	// (三级)问题分类节点

	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 字段名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段类型

	Types *string `json:"Types,omitempty" name:"Types"`
	// 版权

	Copywriter *string `json:"Copywriter,omitempty" name:"Copywriter"`
	// 显示顺序权重

	Weights *uint64 `json:"Weights,omitempty" name:"Weights"`
	// 是否必填, 0|否, 1|是

	Required *uint64 `json:"Required,omitempty" name:"Required"`
	// 字段键值

	FieldKv *string `json:"FieldKv,omitempty" name:"FieldKv"`
}

type GetScheduleStaffRequest struct {
	*tchttp.BaseRequest

	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页页数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 客服中文名

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
}

func (r *GetScheduleStaffRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetScheduleStaffRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 工单一级分类

	TypeLevel1 []*uint64 `json:"TypeLevel1,omitempty" name:"TypeLevel1"`
	// 工单二级分类

	TypeLevel2 []*uint64 `json:"TypeLevel2,omitempty" name:"TypeLevel2"`
	// 工单三级分类

	TypeLevel3 []*uint64 `json:"TypeLevel3,omitempty" name:"TypeLevel3"`
	// 工单优先级

	Priority []*uint64 `json:"Priority,omitempty" name:"Priority"`
	// 客户ID

	CustomerId *string `json:"CustomerId,omitempty" name:"CustomerId"`
	// 当前处理客服

	CurrentStaff *string `json:"CurrentStaff,omitempty" name:"CurrentStaff"`
	// 结单人

	Finisher *string `json:"Finisher,omitempty" name:"Finisher"`
	// 创建时间（起点）

	CreateTimeStart *string `json:"CreateTimeStart,omitempty" name:"CreateTimeStart"`
	// 结单时间（起点）

	FinishTimeStart *string `json:"FinishTimeStart,omitempty" name:"FinishTimeStart"`
	// 撤单时间（起点）

	CancelTimeStart *string `json:"CancelTimeStart,omitempty" name:"CancelTimeStart"`
	// 创建时间（终点）

	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" name:"CreateTimeEnd"`
	// 结单时间（终点）

	FinishTimeEnd *string `json:"FinishTimeEnd,omitempty" name:"FinishTimeEnd"`
	// 撤单时间（终点）

	CancelTimeEnd *string `json:"CancelTimeEnd,omitempty" name:"CancelTimeEnd"`
	// 按创建时间排序,asc或desc

	OrderByCreateTime *string `json:"OrderByCreateTime,omitempty" name:"OrderByCreateTime"`
	// 按结单时间排序,asc或desc

	OrderByFinishTime *string `json:"OrderByFinishTime,omitempty" name:"OrderByFinishTime"`
	// 按撤单时间排序,asc或desc

	OrderByCancelTime *string `json:"OrderByCancelTime,omitempty" name:"OrderByCancelTime"`
	// 工单ID

	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
	// 工单状态（兼容租户端和运营端）

	Status []*uint64 `json:"Status,omitempty" name:"Status"`
	// 按更新时间排序,asc或desc

	OrderByUpdateTime *string `json:"OrderByUpdateTime,omitempty" name:"OrderByUpdateTime"`
}

type StaffUrgeRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 工单ID

	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
}

func (r *StaffUrgeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffUrgeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffEnquireCouldFinishTicketResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StaffEnquireCouldFinishTicketResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffEnquireCouldFinishTicketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffSignOutResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *Code `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StaffSignOutResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffSignOutResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGroupInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 客服组信息

		Data *BasicStaff `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGroupInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffEnterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码

		Data *Code `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StaffEnterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffEnterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFieldRequest struct {
	*tchttp.BaseRequest

	// 表单ID

	FieldId *uint64 `json:"FieldId,omitempty" name:"FieldId"`
	// 求情来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *DeleteFieldRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFieldRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Message struct {
	// 消息ID

	MessageId *uint64 `json:"MessageId,omitempty" name:"MessageId"`
	// 消息类型

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 工单ID

	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
	// 消息接收人

	Receiver *string `json:"Receiver,omitempty" name:"Receiver"`
	// 消息触发人

	Trigger *string `json:"Trigger,omitempty" name:"Trigger"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是否已读

	IsRead *uint64 `json:"IsRead,omitempty" name:"IsRead"`
	// 创建时间

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功返回的json格式

		Data *SucCode `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportTicketsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出的数据结构

		Data *ExportDatas `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportTicketsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportTicketsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFieldResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 表单字段列表

		Data *DescribeFieldsData `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFieldResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFieldResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetScheduleStaffResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询排班客服的数据结构

		Data *ScheduleStaffsData `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetScheduleStaffResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetScheduleStaffResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffRespondResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StaffRespondResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffRespondResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSecretContentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSecretContentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSecretContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupStaffResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功返回的json格式

		Data *SucCode `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGroupStaffResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGroupStaffResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupStaffRequest struct {
	*tchttp.BaseRequest

	// 客服组ID

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 客服ID

	StaffId []*string `json:"StaffId,omitempty" name:"StaffId"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *DeleteGroupStaffRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGroupStaffRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BasicCategory struct {
	// 分类ID

	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 父节点ID

	ParentId *uint64 `json:"ParentId,omitempty" name:"ParentId"`
	// 产品名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 节点描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 图标cdn存储地址

	ImgUrl *string `json:"ImgUrl,omitempty" name:"ImgUrl"`
	// 显示顺序权重

	Weights *uint64 `json:"Weights,omitempty" name:"Weights"`
	// 是否常用

	Common *uint64 `json:"Common,omitempty" name:"Common"`
}

type ThirdCategoryCell struct {
	// 三级分类名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 显示排序权重

	Weights *uint64 `json:"Weights,omitempty" name:"Weights"`
	// 表单字段数组

	CustomFields []*CustomField `json:"CustomFields,omitempty" name:"CustomFields"`
}

type GetCosParamsByBucketRequest struct {
	*tchttp.BaseRequest

	// 桶类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *GetCosParamsByBucketRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCosParamsByBucketRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupFilter struct {
	// 客服组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 按创建时间排序,asc或desc

	OrderByCreateTime *string `json:"OrderByCreateTime,omitempty" name:"OrderByCreateTime"`
	// 按更新时间排序,asc或desc

	OrderByUpdateTime *string `json:"OrderByUpdateTime,omitempty" name:"OrderByUpdateTime"`
}

type GetStaffByUserIdRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 客服登录名

	Uid *string `json:"Uid,omitempty" name:"Uid"`
}

func (r *GetStaffByUserIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStaffByUserIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffDismissResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码

		Data *Code `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StaffDismissResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffDismissResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffPromptToSupplyRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 回复内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 机密信息

	SecretContent *string `json:"SecretContent,omitempty" name:"SecretContent"`
	// 工单ID

	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
	// 设置工单优先级

	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
}

func (r *StaffPromptToSupplyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffPromptToSupplyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffReadMessageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StaffReadMessageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffReadMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Urge struct {
	// 催单记录ID

	UrgeId *uint64 `json:"UrgeId,omitempty" name:"UrgeId"`
	// 催单人

	Urger *string `json:"Urger,omitempty" name:"Urger"`
	// 催单人角色

	UrgerRole *string `json:"UrgerRole,omitempty" name:"UrgerRole"`
	// 工单ID

	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
	// 工单当前处理客服

	CurrentStaff *string `json:"CurrentStaff,omitempty" name:"CurrentStaff"`
	// 催单时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type CheckBeforeDeleteStaffStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功返回后的json

		Data []*SucCode `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckBeforeDeleteStaffStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckBeforeDeleteStaffStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCategoryRequest struct {
	*tchttp.BaseRequest

	// 父节点ID

	ParentId *uint64 `json:"ParentId,omitempty" name:"ParentId"`
	// 节点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *GetCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FirstCategoryCell struct {
	// 一级分类名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 显示排序权重

	Weights *uint64 `json:"Weights,omitempty" name:"Weights"`
}

type OnlineStatusStaff struct {
	// 客服Id

	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`
	// 客服名称

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
	// 是否在线

	IsOnline *uint64 `json:"IsOnline,omitempty" name:"IsOnline"`
	// 是否当班

	IsOnDuty *uint64 `json:"IsOnDuty,omitempty" name:"IsOnDuty"`
	// 工单数目

	Load *uint64 `json:"Load,omitempty" name:"Load"`
}

type DeleteFieldResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功返回的json格式

		Data *SucCode `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFieldResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFieldResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CosBucketData struct {
	// Cos桶参数

	Public *DescribeCosBuckets `json:"Public,omitempty" name:"Public"`
	// 私有桶参数

	Private []*DescribeCosBuckets `json:"Private,omitempty" name:"Private"`
	// COS域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type UpdateScheduleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功返回的json格式

		Data *SucCode `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateScheduleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomField struct {
	// 字段ID

	Fields *uint64 `json:"Fields,omitempty" name:"Fields"`
	// 归属的三级分类ID

	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 表单字段名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 表单字段类型

	Types *string `json:"Types,omitempty" name:"Types"`
	// 可选的键值

	FieldKv *string `json:"FieldKv,omitempty" name:"FieldKv"`
	// 字段文案

	Copywriter *string `json:"Copywriter,omitempty" name:"Copywriter"`
	// 排序权重

	Weights *uint64 `json:"Weights,omitempty" name:"Weights"`
	// 是否必填

	Required *uint64 `json:"Required,omitempty" name:"Required"`
}

type DescribeFirstCategorysData struct {
	// 状态码

	Code *uint64 `json:"Code,omitempty" name:"Code"`
	// 提示信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 一级分类信息数组

	Data []*BasicFirstCategory `json:"Data,omitempty" name:"Data"`
}

type GetSchedule struct {
	// 查询开始日期

	Start *string `json:"Start,omitempty" name:"Start"`
	// 查询截至日期

	End *string `json:"End,omitempty" name:"End"`
	// 客服数组

	Users *string `json:"Users,omitempty" name:"Users"`
}

type Code struct {
	// 状态码

	Code *uint64 `json:"Code,omitempty" name:"Code"`
}

type StaffSignInResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 响应

		Data *Code `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StaffSignInResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffSignInResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFieldRequest struct {
	*tchttp.BaseRequest

	// 产品ID

	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 字段名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段类型

	Types *string `json:"Types,omitempty" name:"Types"`
	// 显示顺序权重

	Weights *uint64 `json:"Weights,omitempty" name:"Weights"`
	// 字段键值

	FieldKv *string `json:"FieldKv,omitempty" name:"FieldKv"`
	// 文案

	Copywriter *string `json:"Copywriter,omitempty" name:"Copywriter"`
	// 是否必填

	Required *uint64 `json:"Required,omitempty" name:"Required"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *CreateFieldRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFieldRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OnlineStatusStaffs struct {
	// 客服数组

	Data []*OnlineStatusStaff `json:"Data,omitempty" name:"Data"`
	// 总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type DeleteCategoryRequest struct {
	*tchttp.BaseRequest

	// 分类ID

	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *DeleteCategoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCategoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffSupplyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StaffSupplyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffSupplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffPromptToSupplyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StaffPromptToSupplyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffPromptToSupplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffReceiveTicketRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 工单ID

	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
}

func (r *StaffReceiveTicketRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffReceiveTicketRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回状态码

		Data *Code `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMessagesRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 查询条件

	Filters []*MessageFilter `json:"Filters,omitempty" name:"Filters"`
	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页记录数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMessagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMessagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStaffSelfInfoRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *GetStaffSelfInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStaffSelfInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 客服组列表

		Data *DescribeStaffsData `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BasicUserInfoForStaff struct {
	// 用户登录名，即StaffName

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
	// 用户登录UIN

	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`
}

type GetStaffByUserIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 搜索到的用户结果

		Data *BasicUserInfoForStaff `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStaffByUserIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStaffByUserIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTicketCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询一级分类返回的数据结构

		Data *DescribeFirstCategorysData `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTicketCategoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTicketCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFieldResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功返回后的json格式

		Data *SucCode `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFieldResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFieldResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteStaffStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功返回后的json

		Data *SucCode `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteStaffStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteStaffStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BasicStaff struct {
	// 客服ID

	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`
	// 客服名称

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
	// 是否在线

	IsOnline *uint64 `json:"IsOnline,omitempty" name:"IsOnline"`
	// 是否在职

	IsDismissed *uint64 `json:"IsDismissed,omitempty" name:"IsDismissed"`
	// 客服负荷,指当前客服名下(即当前处理人)

	Load *uint64 `json:"Load,omitempty" name:"Load"`
}

type DescribeSchedulesData struct {
	// 状态码

	Code *uint64 `json:"Code,omitempty" name:"Code"`
	// 提示信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 所有客服的排班数据

	Data []*PersonSchedule `json:"Data,omitempty" name:"Data"`
}

type DescribeCategorysData struct {
	// 分类列表

	Data []*BasicCategory `json:"Data,omitempty" name:"Data"`
	// 状态码

	Code *uint64 `json:"Code,omitempty" name:"Code"`
	// 提示信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type GetCosBucketRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *GetCosBucketRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCosBucketRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChildrenCategory struct {
	// 分类Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 二级分类

	Cell *SecondCategoryCell `json:"Cell,omitempty" name:"Cell"`
}

type UpdateGroupRequest struct {
	*tchttp.BaseRequest

	// 客服组ID

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 客服组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *UpdateGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteStaffStatusRequest struct {
	*tchttp.BaseRequest

	// 客服ID

	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *DeleteStaffStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteStaffStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeThirdCategorysData struct {
	// 状态码

	Code *uint64 `json:"Code,omitempty" name:"Code"`
	// 提示信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 三级分类信息数组

	Data []*BasicThirdLevel `json:"Data,omitempty" name:"Data"`
}

type StaffReadMessageRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 消息ID

	MessageId *uint64 `json:"MessageId,omitempty" name:"MessageId"`
}

func (r *StaffReadMessageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffReadMessageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffSupplyRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 回复内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 机密信息

	SecretContent *string `json:"SecretContent,omitempty" name:"SecretContent"`
	// 工单ID

	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
	// 转单目标客服

	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`
	// 转单目标客服组

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 设置工单优先级

	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
}

func (r *StaffSupplyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffSupplyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTicketRequest struct {
	*tchttp.BaseRequest

	// 客户账号uin

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 工单一级分类

	TypeLevel1 *uint64 `json:"TypeLevel1,omitempty" name:"TypeLevel1"`
	// 工单二级分类

	TypeLevel2 *uint64 `json:"TypeLevel2,omitempty" name:"TypeLevel2"`
	// 工单三级分类

	TypeLevel3 *uint64 `json:"TypeLevel3,omitempty" name:"TypeLevel3"`
	// 工单内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 工单机密信息

	SecretContent *string `json:"SecretContent,omitempty" name:"SecretContent"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 自定义字段

	CustomFields *string `json:"CustomFields,omitempty" name:"CustomFields"`
}

func (r *CreateTicketRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTicketRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Conversation struct {
	// 回复记录ID

	ConversationId *uint64 `json:"ConversationId,omitempty" name:"ConversationId"`
	// 回复内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 机密信息

	SecretContent *string `json:"SecretContent,omitempty" name:"SecretContent"`
	// 工单ID

	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
	// 回复的人

	Speaker *string `json:"Speaker,omitempty" name:"Speaker"`
	// 回复人的角色

	SpeakerRole *string `json:"SpeakerRole,omitempty" name:"SpeakerRole"`
	// 接收回复的人

	Listener *string `json:"Listener,omitempty" name:"Listener"`
	// 接收回复的人的角色

	ListenerRole *string `json:"ListenerRole,omitempty" name:"ListenerRole"`
	// 回复时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeScheduleStaffsData struct {
	// 状态码

	Code *uint64 `json:"Code,omitempty" name:"Code"`
	// 提示信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 排班客服数据

	Data *ScheduleStaffsData `json:"Data,omitempty" name:"Data"`
}

type CreateStaffRequest struct {
	*tchttp.BaseRequest

	// 客服ID

	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`
	// 客服名称

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *CreateStaffRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateStaffRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffEnterRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 客服ID

	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`
}

func (r *StaffEnterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffEnterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffRespondRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 回复内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 机密信息

	SecretContent *string `json:"SecretContent,omitempty" name:"SecretContent"`
	// 工单ID

	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
	// 设置工单优先级

	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
}

func (r *StaffRespondRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffRespondRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOnlineStaffsData struct {
	// 状态码

	Code *uint64 `json:"Code,omitempty" name:"Code"`
	// 提示信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 客服信息

	Data *OnlineStatusStaffs `json:"Data,omitempty" name:"Data"`
}

type DeleteCategoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功返回的json格式

		Data *SucCode `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
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

type GetTicketCategoryByLevelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 三级分类和分类对应的表单字段

		Data *DescribeThirdCategorysData `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTicketCategoryByLevelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTicketCategoryByLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportTicketsRequest struct {
	*tchttp.BaseRequest

	// 工单过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页记录数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *ExportTicketsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportTicketsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetScheduleRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 开始日期

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束日期

	End *string `json:"End,omitempty" name:"End"`
	// 客服人员

	Users []*string `json:"Users,omitempty" name:"Users"`
}

func (r *GetScheduleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetScheduleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStaffSelfInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户信息

		Data *BasicStaff `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStaffSelfInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStaffSelfInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFieldsData struct {
	// 表单字段列表

	Data []*BasicField `json:"Data,omitempty" name:"Data"`
	// 状态码

	Code *uint64 `json:"Code,omitempty" name:"Code"`
	// 提示信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type WorkbenchData struct {
	// 待处理

	WaitingProcessCount *uint64 `json:"WaitingProcessCount,omitempty" name:"WaitingProcessCount"`
	// 处理中

	ProcessingCount *uint64 `json:"ProcessingCount,omitempty" name:"ProcessingCount"`
	// 优先级一般

	PriorityNormalCount *uint64 `json:"PriorityNormalCount,omitempty" name:"PriorityNormalCount"`
	// 优先级紧急

	PriorityUrgentCount *uint64 `json:"PriorityUrgentCount,omitempty" name:"PriorityUrgentCount"`
	// 优先级非常紧急

	PriorityVeryUrgentCount *uint64 `json:"PriorityVeryUrgentCount,omitempty" name:"PriorityVeryUrgentCount"`
	// 4小时未结单

	NotFinished4HoursCount *uint64 `json:"NotFinished4HoursCount,omitempty" name:"NotFinished4HoursCount"`
	// 8小时未结单

	NotFinished8HoursCount *uint64 `json:"NotFinished8HoursCount,omitempty" name:"NotFinished8HoursCount"`
	// 24小时未结单

	NotFinished24HoursCount *uint64 `json:"NotFinished24HoursCount,omitempty" name:"NotFinished24HoursCount"`
	// 48小时未结单

	NotFinished48HoursCount *uint64 `json:"NotFinished48HoursCount,omitempty" name:"NotFinished48HoursCount"`
	// 待客户补充

	WaitingCustomerSupplyCount *uint64 `json:"WaitingCustomerSupplyCount,omitempty" name:"WaitingCustomerSupplyCount"`
	// 待客户结单

	WaitingCustomerFinishCount *uint64 `json:"WaitingCustomerFinishCount,omitempty" name:"WaitingCustomerFinishCount"`
}

type DescribeTicketsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 工单列表

		Data *DescribeTicketsData `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTicketsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTicketsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStaffParams struct {
	// 客服ID

	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`
	// 客服名称

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
	// 是否离职

	IsDismissed *string `json:"IsDismissed,omitempty" name:"IsDismissed"`
	// 是否在线

	IsOnline *string `json:"IsOnline,omitempty" name:"IsOnline"`
}

type ScheduleStaffsData struct {
	// 客服排班数组

	Data []*ScheduleStaff `json:"Data,omitempty" name:"Data"`
	// 客服数目

	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type AddGroupStaffResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功返回的json格式

		Data *SucCode `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddGroupStaffResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddGroupStaffResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffAssignTicketResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Data *Code `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StaffAssignTicketResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffAssignTicketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStaffData struct {
	// 客服总数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
	// 客服数据

	Data []*BasicStaff `json:"Data,omitempty" name:"Data"`
}

type GetOnlineStaffParams struct {
	// 客服中文名

	StaffName *string `json:"StaffName,omitempty" name:"StaffName"`
	// 是否在线

	IsOnline *uint64 `json:"IsOnline,omitempty" name:"IsOnline"`
	// 是否当班

	IsOnDuty *uint64 `json:"IsOnDuty,omitempty" name:"IsOnDuty"`
}

type Ticket struct {
	// 工单ID

	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
	// 工单一级分类

	TypeLevel1 *uint64 `json:"TypeLevel1,omitempty" name:"TypeLevel1"`
	// 工单二级分类

	TypeLevel2 *uint64 `json:"TypeLevel2,omitempty" name:"TypeLevel2"`
	// 工单三级分类

	TypeLevel3 *uint64 `json:"TypeLevel3,omitempty" name:"TypeLevel3"`
	// 租户端工单状态

	AskerStatus *uint64 `json:"AskerStatus,omitempty" name:"AskerStatus"`
	// 运营端工单状态

	AnswererStatus *uint64 `json:"AnswererStatus,omitempty" name:"AnswererStatus"`
	// 优先级

	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
	// 工单内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 机密信息

	SecretContent *string `json:"SecretContent,omitempty" name:"SecretContent"`
	// 客户ID

	CustomerId *string `json:"CustomerId,omitempty" name:"CustomerId"`
	// 工单当前处理客服

	CurrentStaff *string `json:"CurrentStaff,omitempty" name:"CurrentStaff"`
	// 结单人

	Finisher *string `json:"Finisher,omitempty" name:"Finisher"`
	// 工单创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 回复记录

	Conversations []*Conversation `json:"Conversations,omitempty" name:"Conversations"`
	// 派单记录

	Assignments []*Assignment `json:"Assignments,omitempty" name:"Assignments"`
	// 结单时间

	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
	// 撤单时间

	CancelTime *string `json:"CancelTime,omitempty" name:"CancelTime"`
	// 催单记录

	Urges []*Urge `json:"Urges,omitempty" name:"Urges"`
}

type CosParams struct {
	// SecretId

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// SecretKey

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
}

type StaffUrgeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Data *Code `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StaffUrgeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffUrgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMessagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 消息列表

		Data []*DescribeMessagesData `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMessagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMessagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCosBuckets struct {
	// appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 桶名称

	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeTicketsRequest struct {
	*tchttp.BaseRequest

	// 工单过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页记录数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
}

func (r *DescribeTicketsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTicketsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportScheduleRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 开始日期

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束日期

	End *string `json:"End,omitempty" name:"End"`
}

func (r *ExportScheduleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportScheduleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSecretContentRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 工单ID

	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
	// 回复记录ID（指定时获取的是某条回复记录的机密信息）

	ConversationId *uint64 `json:"ConversationId,omitempty" name:"ConversationId"`
}

func (r *GetSecretContentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSecretContentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaffEnquireCouldFinishTicketRequest struct {
	*tchttp.BaseRequest

	// 请求来源

	Requester *string `json:"Requester,omitempty" name:"Requester"`
	// 工单ID

	TicketId *uint64 `json:"TicketId,omitempty" name:"TicketId"`
	// 回复内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 机密信息

	SecretContent *string `json:"SecretContent,omitempty" name:"SecretContent"`
	// 设置工单优先级

	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
}

func (r *StaffEnquireCouldFinishTicketRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StaffEnquireCouldFinishTicketRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkbenchDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运营端数据

		Data *WorkbenchData `json:"Data,omitempty" name:"Data"`
		// 状态码

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// 提示信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetWorkbenchDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetWorkbenchDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
