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

package v20220508

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type CreateOrganizationMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成员账户uin

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrganizationMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrganizationMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrganizationMemberRequest struct {
	*tchttp.BaseRequest

	// 集团组织管理员uin

	HostUin *uint64 `json:"HostUin,omitempty" name:"HostUin"`
	// tce账户用户名， 不填则使Name

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 成员账户名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 集团组织部门ID, 默认放到根节点下

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateOrganizationMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrganizationMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOrganizationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationRequest struct {
	*tchttp.BaseRequest

	// 管理员uin

	HostUin *uint64 `json:"HostUin,omitempty" name:"HostUin"`
}

func (r *DeleteOrganizationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationRequest struct {
	*tchttp.BaseRequest

	// 集团组织管理员uin

	HostUin *uint64 `json:"HostUin,omitempty" name:"HostUin"`
	// 集团组织名称

	OrgName *string `json:"OrgName,omitempty" name:"OrgName"`
}

func (r *UpdateOrganizationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOrganizationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrganizationBasicInfo struct {
	// 组织id

	OrgId *uint64 `json:"OrgId,omitempty" name:"OrgId"`
	// 管理员uin

	HostUin *uint64 `json:"HostUin,omitempty" name:"HostUin"`
	// 成员数量

	MemberNum *int64 `json:"MemberNum,omitempty" name:"MemberNum"`
	// 集团组织名称

	OrgName *string `json:"OrgName,omitempty" name:"OrgName"`
	// 集团组织类型

	OrgType *string `json:"OrgType,omitempty" name:"OrgType"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OrganizationMember struct {
	// 成员账号uin

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 成员账号名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 成员类型

	MemberType *string `json:"MemberType,omitempty" name:"MemberType"`
	// 集团策略类型

	OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`
	// 集团策略名称

	OrgPolicyName *string `json:"OrgPolicyName,omitempty" name:"OrgPolicyName"`
	// 代付uin

	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`
	// 代付用户名称

	PayName *string `json:"PayName,omitempty" name:"PayName"`
	// 节点id

	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是否允许退出

	IsAllowQuit *string `json:"IsAllowQuit,omitempty" name:"IsAllowQuit"`
	// 权限状态

	PermissionStatus *string `json:"PermissionStatus,omitempty" name:"PermissionStatus"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 权限列表

	OrgPermission []*Permission `json:"OrgPermission,omitempty" name:"OrgPermission"`
	// 身份列表

	OrgIdentity []*Identity `json:"OrgIdentity,omitempty" name:"OrgIdentity"`
}

type Permission struct {
	// id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateOrganizationRequest struct {
	*tchttp.BaseRequest

	// 管理员uin

	HostUin *uint64 `json:"HostUin,omitempty" name:"HostUin"`
	// 组织名称

	OrgName *string `json:"OrgName,omitempty" name:"OrgName"`
	// 创建类型，运营端填CMGT

	CreateType *string `json:"CreateType,omitempty" name:"CreateType"`
}

func (r *CreateOrganizationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrganizationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOrganizationMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationsRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询数量大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 时间起点 2006-01-02 15:04:05格式

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 时间终点 2006-01-02 15:04:05格式

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 检索关键字， 支持OrgName

	KeywordInfo []*KeywordInfo `json:"KeywordInfo,omitempty" name:"KeywordInfo"`
}

func (r *DescribeOrganizationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOrganizationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 全部数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 组织列表

		Items []*OrganizationBasicInfo `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeywordInfo struct {
	// 关键字类型

	KeywordType *string `json:"KeywordType,omitempty" name:"KeywordType"`
	// 关键字内容

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type CreateOrganizationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组织id

		OrgId *uint64 `json:"OrgId,omitempty" name:"OrgId"`
		// 组织别名

		NickName *string `json:"NickName,omitempty" name:"NickName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrganizationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationMemberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOrganizationMemberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOrganizationMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationMembersRequest struct {
	*tchttp.BaseRequest

	// 成员账号uin列表

	MemberUin []*uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 管理员uin

	HostUin *uint64 `json:"HostUin,omitempty" name:"HostUin"`
}

func (r *DeleteOrganizationMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrganizationMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMembersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 全部数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 成员账户列表

		Items []*OrganizationMember `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationMembersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrganizationMemberRequest struct {
	*tchttp.BaseRequest

	// 管理员uin

	HostUin *uint64 `json:"HostUin,omitempty" name:"HostUin"`
	// 成员账号uin

	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
	// 成员名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *UpdateOrganizationMemberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOrganizationMemberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Identity struct {
	// id

	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`
	// 别名

	IdentityAliasName *string `json:"IdentityAliasName,omitempty" name:"IdentityAliasName"`
}

type DescribeOrganizationMembersRequest struct {
	*tchttp.BaseRequest

	// 集团管理员uin

	HostUin *uint64 `json:"HostUin,omitempty" name:"HostUin"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小，[0,50]

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索关键字， 支持名称和uin

	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeOrganizationMembersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrganizationMembersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
