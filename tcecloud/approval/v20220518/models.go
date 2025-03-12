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

package v20220518

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type CreateFlowRequest struct {
	*tchttp.BaseRequest

	// 审批流名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// schema

	SchemaProps *string `json:"SchemaProps,omitempty" name:"SchemaProps"`
	// 是否启用审批流

	Activated *bool `json:"Activated,omitempty" name:"Activated"`
	// 业务id列表

	ActionIDs []*uint64 `json:"ActionIDs,omitempty" name:"ActionIDs"`
	// 节点列表

	Stages []*StageParam `json:"Stages,omitempty" name:"Stages"`
	// 作用范围

	Scopes []*Scope `json:"Scopes,omitempty" name:"Scopes"`
}

func (r *CreateFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApprovalActionRequest struct {
	*tchttp.BaseRequest

	// 业务参数

	Params *ActionParams `json:"Params,omitempty" name:"Params"`
}

func (r *UpdateApprovalActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApprovalActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApprovalActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApprovalActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApprovalActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionRelatedInfo struct {
	// 审批业务ID

	ActionID *int64 `json:"ActionID,omitempty" name:"ActionID"`
	// 关联的审批流信息

	Flow []*FlowAttr `json:"Flow,omitempty" name:"Flow"`
	// 关联的用户信息，-1代表全体租户

	Users []*UserScope `json:"Users,omitempty" name:"Users"`
}

type DescribeFlowSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回参数

		Data *FlowSet `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PerformApprovalRequest struct {
	*tchttp.BaseRequest

	// 当前审批节点序号

	StageSerialNum *uint64 `json:"StageSerialNum,omitempty" name:"StageSerialNum"`
	// 审批单id

	PaperID *uint64 `json:"PaperID,omitempty" name:"PaperID"`
	// 审批意见 12: 拒绝，14: 同意

	Operate *uint64 `json:"Operate,omitempty" name:"Operate"`
	// 审批意见

	Opinion *string `json:"Opinion,omitempty" name:"Opinion"`
	// 审批修改后的参数

	Request *string `json:"Request,omitempty" name:"Request"`
}

func (r *PerformApprovalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PerformApprovalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApprovalPaperSet struct {
	// 审批单列表

	PaperSet []*ApprovalPaperAttr `json:"PaperSet,omitempty" name:"PaperSet"`
	// 总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type GetApprovedSetRequest struct {
	*tchttp.BaseRequest

	// 业务名称

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 申请人uin

	ApplicantUin *string `json:"ApplicantUin,omitempty" name:"ApplicantUin"`
	// 申请人

	Applicant *string `json:"Applicant,omitempty" name:"Applicant"`
	// 所属主账号

	OwnerAccount *string `json:"OwnerAccount,omitempty" name:"OwnerAccount"`
	// 申请理由

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 页数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 审批单id

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 审批单状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 排序字段

	Sort *Order `json:"Sort,omitempty" name:"Sort"`
}

func (r *GetApprovedSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApprovedSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryApprovalDocDetailRequest struct {
	*tchttp.BaseRequest

	// 审批单id

	ID *uint64 `json:"ID,omitempty" name:"ID"`
}

func (r *QueryApprovalDocDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryApprovalDocDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlowAttr struct {
	// 审批流是否已开启

	Activated *bool `json:"Activated,omitempty" name:"Activated"`
	// 关联业务id列表

	ActionIDs []*int64 `json:"ActionIDs,omitempty" name:"ActionIDs"`
	// 创建时间

	CTime *string `json:"CTime,omitempty" name:"CTime"`
	// 创建人uin

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// 创建用户

	CreateUser *string `json:"CreateUser,omitempty" name:"CreateUser"`
	// 审批流描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 审批流id

	FlowID *int64 `json:"FlowID,omitempty" name:"FlowID"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建用户主账号uin

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 创建端，ocloud:运营端, tcloud:租户端

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 显示字段描述信息，此处无作用

	SchemaProps *string `json:"SchemaProps,omitempty" name:"SchemaProps"`
	// 审批流更新时间

	UpTime *string `json:"UpTime,omitempty" name:"UpTime"`
	// 版本

	Version *int64 `json:"Version,omitempty" name:"Version"`
}

type DeleteApprovalActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApprovalActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApprovalActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionOutput struct {
	// 审批业务参数集合

	Actions []*ActionParams `json:"Actions,omitempty" name:"Actions"`
	// 总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type FilterAttr struct {
	// 参数名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 操作符，支持equal, like

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type StageParam struct {
	// 节点id

	StageID *uint64 `json:"StageID,omitempty" name:"StageID"`
	// 节点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// false会签，true或签

	SingleSeal *bool `json:"SingleSeal,omitempty" name:"SingleSeal"`
	// 审批人uin列表

	Approvers []*string `json:"Approvers,omitempty" name:"Approvers"`
	// 节点顺序序号

	SerialNumber *uint64 `json:"SerialNumber,omitempty" name:"SerialNumber"`
	// 此节点审批信息

	Seals []*SealAttr `json:"Seals,omitempty" name:"Seals"`
	// 此节点审批状态, 0：未进行到此节点，14：此节点已审批通过，12：此节点审批不通过，15：此节点为自动同意

	StageStatus *int64 `json:"StageStatus,omitempty" name:"StageStatus"`
	// 审批人信息

	ApproverInfo []*ApproverAttr `json:"ApproverInfo,omitempty" name:"ApproverInfo"`
}

type DeleteApprovalFlowRequest struct {
	*tchttp.BaseRequest

	// 审批流id

	FlowID *uint64 `json:"FlowID,omitempty" name:"FlowID"`
}

func (r *DeleteApprovalFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApprovalFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApprovalFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApprovalFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApprovalFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryApprovalDocDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 审批单详情

		Data []*ApprovalPaperAttr `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryApprovalDocDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryApprovalDocDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApproverAttr struct {
	// 审批人状态,Normal：正常,Unknown：异常

	ApproverStatus *string `json:"ApproverStatus,omitempty" name:"ApproverStatus"`
	// 审批人uin

	ApproverUin *string `json:"ApproverUin,omitempty" name:"ApproverUin"`
	// 审批人用户名

	ApproverUserName *string `json:"ApproverUserName,omitempty" name:"ApproverUserName"`
}

type DeleteApprovalActionRequest struct {
	*tchttp.BaseRequest

	// 审批业务ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteApprovalActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApprovalActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PerformApprovalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PerformApprovalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PerformApprovalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryActionParams struct {
	// id列表

	ActionIDs []*uint64 `json:"ActionIDs,omitempty" name:"ActionIDs"`
	// 接口名

	Action *string `json:"Action,omitempty" name:"Action"`
	// 接口对应模块

	Module *string `json:"Module,omitempty" name:"Module"`
	// 每页显示条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 业务名称

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 接口对应模块(中文名)

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 接口对应产品(中文名)

	YunProductName *string `json:"YunProductName,omitempty" name:"YunProductName"`
	// 接口分类(tcloud, ocloud)

	Category *string `json:"Category,omitempty" name:"Category"`
	// 排序字段

	Sort *Order `json:"Sort,omitempty" name:"Sort"`
	// 过滤条件

	Filters []*FilterAttr `json:"Filters,omitempty" name:"Filters"`
}

type CreateFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserScope struct {
	// 用户uin，-1代表全体租户

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 是否为子用户

	IsSubAccount *bool `json:"IsSubAccount,omitempty" name:"IsSubAccount"`
	// UserName

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type GetFlowDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 审批流详情出参

		Data *FlowOutputAttr `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFlowDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFlowDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryApprovalDocResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *ApprovalPaperSet `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryApprovalDocResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryApprovalDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Scope struct {
	// 业务id

	ActionID *uint64 `json:"ActionID,omitempty" name:"ActionID"`
	// 作用用户范围

	Users []*UserScope `json:"Users,omitempty" name:"Users"`
}

type FlowSet struct {
	// 审批流集合

	Flows []*FlowOutputAttr `json:"Flows,omitempty" name:"Flows"`
	// 总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type DescribeFlowSetRequest struct {
	*tchttp.BaseRequest

	// 审批流名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 页数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 是否启用审批流, 0-否，1-是

	IsActive *int64 `json:"IsActive,omitempty" name:"IsActive"`
	// 排序字段信息

	Sort *Order `json:"Sort,omitempty" name:"Sort"`
	// 过滤参数列表

	Filter []*FilterAttr `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeFlowSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApprovalFlowRequest struct {
	*tchttp.BaseRequest

	// 审批流名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// schema

	SchemaProps *string `json:"SchemaProps,omitempty" name:"SchemaProps"`
	// 是否启用审批流

	Activated *bool `json:"Activated,omitempty" name:"Activated"`
	// 业务id列表

	ActionIDs []*uint64 `json:"ActionIDs,omitempty" name:"ActionIDs"`
	// 节点列表

	Stages []*StageParam `json:"Stages,omitempty" name:"Stages"`
	// 作用范围

	Scopes []*Scope `json:"Scopes,omitempty" name:"Scopes"`
	// 审批流ID

	FlowID *uint64 `json:"FlowID,omitempty" name:"FlowID"`
}

func (r *ModifyApprovalFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApprovalFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCurrApprovalDetailRequest struct {
	*tchttp.BaseRequest

	// 审批单 id

	ID *uint64 `json:"ID,omitempty" name:"ID"`
}

func (r *QueryCurrApprovalDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCurrApprovalDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryUsersUnderActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 业务已关联的审批流以及用户数据

		Data []*ActionRelatedInfo `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryUsersUnderActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryUsersUnderActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionParams struct {
	// 业务id

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 审批业务名称

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 所属服务名称(中文)

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 产品名称

	YunProductName *string `json:"YunProductName,omitempty" name:"YunProductName"`
	// 审批业务(云Api)

	Action *string `json:"Action,omitempty" name:"Action"`
	// 所属服务(英文)

	Module *string `json:"Module,omitempty" name:"Module"`
	// 接口版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 业务所属端，tcloud:租户端，ocloud:运营端

	Category *string `json:"Category,omitempty" name:"Category"`
	// 显示内容字段schema(jsonpath)

	Schema *string `json:"Schema,omitempty" name:"Schema"`
	// 审批业务描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 业务Id，冗余

	ActionID *int64 `json:"ActionID,omitempty" name:"ActionID"`
	// 审批业务创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建者uin

	CreateUserUin *string `json:"CreateUserUin,omitempty" name:"CreateUserUin"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// jsonpath

	JsonPath *string `json:"JsonPath,omitempty" name:"JsonPath"`
	// 业务来源, 1: 预设，2: 用户自助接入

	Origin *int64 `json:"Origin,omitempty" name:"Origin"`
	// 业务类型，1：申请，2: 创建，3:删除，4:修改

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 更新用户

	UpdateUser *string `json:"UpdateUser,omitempty" name:"UpdateUser"`
	// 更新用户uin

	UpdateUserUin *string `json:"UpdateUserUin,omitempty" name:"UpdateUserUin"`
}

type CreateApprovalActionRequest struct {
	*tchttp.BaseRequest

	// 业务参数

	Params *ActionParams `json:"Params,omitempty" name:"Params"`
}

func (r *CreateApprovalActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApprovalActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApprovedSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *ApprovalPaperSet `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApprovedSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApprovedSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Order struct {
	// 排序字段英文名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 是否降序, true-降序，false-升序

	IsDesc *bool `json:"IsDesc,omitempty" name:"IsDesc"`
}

type CreateApprovalActionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApprovalActionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApprovalActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryActionSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 审批业务出参

		Data *ActionOutput `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryActionSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryActionSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryApprovalDocRequest struct {
	*tchttp.BaseRequest

	// 业务名称

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 申请人uin

	ApplicantUin *string `json:"ApplicantUin,omitempty" name:"ApplicantUin"`
	// 申请人

	Applicant *string `json:"Applicant,omitempty" name:"Applicant"`
	// 所属主账号

	OwnerAccount *string `json:"OwnerAccount,omitempty" name:"OwnerAccount"`
	// 申请理由

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 页数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 审批单状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 审批单id

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 排序字段信息

	Sort *Order `json:"Sort,omitempty" name:"Sort"`
}

func (r *QueryApprovalDocRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryApprovalDocRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCurrApprovalDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// data

		Data *ApprovalPaperAttr `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCurrApprovalDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCurrApprovalDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCurrApprovalDocResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *ApprovalPaperSet `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCurrApprovalDocResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCurrApprovalDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApprovalFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApprovalFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApprovalFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFlowDetailRequest struct {
	*tchttp.BaseRequest

	// 审批流id

	FlowID *uint64 `json:"FlowID,omitempty" name:"FlowID"`
}

func (r *GetFlowDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFlowDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateFlowStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OperateFlowStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OperateFlowStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryActionSetRequest struct {
	*tchttp.BaseRequest

	// 请求参数

	Params *QueryActionParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryActionSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryActionSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateFlowStatusRequest struct {
	*tchttp.BaseRequest

	// 审批流id

	FlowID *uint64 `json:"FlowID,omitempty" name:"FlowID"`
	// true启用，false停用

	Activated *bool `json:"Activated,omitempty" name:"Activated"`
}

func (r *OperateFlowStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OperateFlowStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApprovalPaperAttr struct {
	// 对应云api接口

	Action *string `json:"Action,omitempty" name:"Action"`
	// 审批业务描述信息

	ActionDescription *string `json:"ActionDescription,omitempty" name:"ActionDescription"`
	// 审批业务ID

	ActionID *int64 `json:"ActionID,omitempty" name:"ActionID"`
	// 审批业务中文名

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 申请人

	Applicant *string `json:"Applicant,omitempty" name:"Applicant"`
	// 申请人uin

	ApplicantUin *string `json:"ApplicantUin,omitempty" name:"ApplicantUin"`
	// 申请时间

	CTime *string `json:"CTime,omitempty" name:"CTime"`
	// 审批通过后，回调状态. 0：未开始, 100：回调成功，101：回调失败

	CallbackStatus *int64 `json:"CallbackStatus,omitempty" name:"CallbackStatus"`
	// 当前审批节点

	CurrStageNum *int64 `json:"CurrStageNum,omitempty" name:"CurrStageNum"`
	// 关联审批流描述信息

	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`
	// 审批流ID

	FlowID *int64 `json:"FlowID,omitempty" name:"FlowID"`
	// 审批业务对应云API模块中文名

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 主账号名称

	OwnerAccount *string `json:"OwnerAccount,omitempty" name:"OwnerAccount"`
	// 主账号uIn

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 审批单ID

	PaperID *int64 `json:"PaperID,omitempty" name:"PaperID"`
	// 上层审批单ID

	ParentID *int64 `json:"ParentID,omitempty" name:"ParentID"`
	// 云产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 审批理由

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// schema

	Schema *string `json:"Schema,omitempty" name:"Schema"`
	// 请求参数，此处无法按云API3规范

	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`
	// 已审批信息

	Seals *SealAttr `json:"Seals,omitempty" name:"Seals"`
	// 审批单状态，0：审批单初始化，1：审批进行中，11:申请撤销，12:审批拒绝, 14: 审批通过，15：自动同意

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 对应审批节点信息

	Stages []*StageParam `json:"Stages,omitempty" name:"Stages"`
}

type QueryCurrApprovalDocRequest struct {
	*tchttp.BaseRequest

	// 业务名称

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 申请人uin

	ApplicantUin *string `json:"ApplicantUin,omitempty" name:"ApplicantUin"`
	// 申请人

	Applicant *string `json:"Applicant,omitempty" name:"Applicant"`
	// 所属主账号

	OwnerAccount *string `json:"OwnerAccount,omitempty" name:"OwnerAccount"`
	// 申请理由

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 页数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 审批单id

	ID *uint64 `json:"ID,omitempty" name:"ID"`
	// 审批单状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 排序字段信息

	Sort *Order `json:"Sort,omitempty" name:"Sort"`
}

func (r *QueryCurrApprovalDocRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCurrApprovalDocRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryUsersUnderActionRequest struct {
	*tchttp.BaseRequest

	// action id 列表

	ActionIDs []*uint64 `json:"ActionIDs,omitempty" name:"ActionIDs"`
}

func (r *QueryUsersUnderActionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryUsersUnderActionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlowOutputAttr struct {
	// 关联ID列表

	ActionIDs []*int64 `json:"ActionIDs,omitempty" name:"ActionIDs"`
	// 审批流启用状态，true：启用，false：关闭

	Activated *bool `json:"Activated,omitempty" name:"Activated"`
	// 创建时间

	CTime *string `json:"CTime,omitempty" name:"CTime"`
	// 创建者uin

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// 创建用户名称

	CreateUser *string `json:"CreateUser,omitempty" name:"CreateUser"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 唯一id

	FlowID *int64 `json:"FlowID,omitempty" name:"FlowID"`
	// 审批流名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建者主账号uin

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 创建端，ocloud/tcloud

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 申请内容显示字段

	SchemaProps *string `json:"SchemaProps,omitempty" name:"SchemaProps"`
	// 更新时间

	UpTime *string `json:"UpTime,omitempty" name:"UpTime"`
	// 版本号

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 审批节点信息

	Stages []*StageParam `json:"Stages,omitempty" name:"Stages"`
	// 审批作用范围

	Scopes []*Scope `json:"Scopes,omitempty" name:"Scopes"`
}

type SealAttr struct {
	// 审批处理ID

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 对应审批单ID

	PaperID *int64 `json:"PaperID,omitempty" name:"PaperID"`
	// 关联审批步骤序号

	StageSerialNum *int64 `json:"StageSerialNum,omitempty" name:"StageSerialNum"`
	// 审批人

	OpUin *string `json:"OpUin,omitempty" name:"OpUin"`
	// 审批时间

	ApproveTime *string `json:"ApproveTime,omitempty" name:"ApproveTime"`
	// 审批操作，12：decline, 14：approval

	Operate *int64 `json:"Operate,omitempty" name:"Operate"`
	// 审批意见

	Opinion *string `json:"Opinion,omitempty" name:"Opinion"`
	// 审批人修改的request

	Request *string `json:"Request,omitempty" name:"Request"`
}
