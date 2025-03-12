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

package v20200901

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type GetUserByLoginUinRequest struct {
	*tchttp.BaseRequest

	// 登录用户Uin

	LoginUin *uint64 `json:"LoginUin,omitempty" name:"LoginUin"`
}

func (r *GetUserByLoginUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserByLoginUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该用户有权限的所有项目

		ProjectInfos []*ProjectInfo `json:"ProjectInfos,omitempty" name:"ProjectInfos"`
		// 符合条件的总的数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProjectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProjectInfo struct {
	// 产品原型code，唯一id，不允许更改

	Name *string `json:"Name,omitempty" name:"Name"`
	// 允许更改，显示名字

	ShowName *string `json:"ShowName,omitempty" name:"ShowName"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// CPU最大核数

	CPULimit *int64 `json:"CPULimit,omitempty" name:"CPULimit"`
	// 最小内核数

	CPURequest *int64 `json:"CPURequest,omitempty" name:"CPURequest"`
	// 最大内存量，单位MB

	MemLimit *int64 `json:"MemLimit,omitempty" name:"MemLimit"`
	// 最小内存量， 单位MB

	MemRequest *int64 `json:"MemRequest,omitempty" name:"MemRequest"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 创建用户uin

	CreationUin *int64 `json:"CreationUin,omitempty" name:"CreationUin"`
	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 是否需要审批

	AppApproved *bool `json:"AppApproved,omitempty" name:"AppApproved"`
	// 集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 项目下pod数量

	PodsCount *int64 `json:"PodsCount,omitempty" name:"PodsCount"`
	// 项目下应用数量

	AppsCount *int64 `json:"AppsCount,omitempty" name:"AppsCount"`
	// CPUlimit已使用值，单位核数

	UsedCPULimit *int64 `json:"UsedCPULimit,omitempty" name:"UsedCPULimit"`
	// cpuRequest 已使用值，单位核数

	UsedCPURequest *int64 `json:"UsedCPURequest,omitempty" name:"UsedCPURequest"`
	// memRequest 已使用值，单位MB

	UsedMemRequest *int64 `json:"UsedMemRequest,omitempty" name:"UsedMemRequest"`
	// memLimit 已使用值 单位MB

	UsedMemLimit *int64 `json:"UsedMemLimit,omitempty" name:"UsedMemLimit"`
	// 项目关联的Namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type GetUserByLoginUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户邮箱

		Email *string `json:"Email,omitempty" name:"Email"`
		// 用户名

		Name *string `json:"Name,omitempty" name:"Name"`
		// 用户电话

		PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
		// 用户Uin

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 用户角色列表

		RoleIds []*string `json:"RoleIds,omitempty" name:"RoleIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserByLoginUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserByLoginUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectsRequest struct {
	*tchttp.BaseRequest

	// 列出clusterId= 所有项目，不填列出所有

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 非必填，合端后兼容原有

	Client *string `json:"Client,omitempty" name:"Client"`
	// 待搜索参数列表，支持Name, ClusterId

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 按某个字段排序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否降序排序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
}

func (r *ListProjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 参数名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤值

	Values []*string `json:"Values,omitempty" name:"Values"`
}
