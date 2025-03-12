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

package v20200727

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DeleteBackupRecordRequest struct {
	*tchttp.BaseRequest

	// 实例名字

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例所在命名空间

	InstanceNamespace *string `json:"InstanceNamespace,omitempty" name:"InstanceNamespace"`
	// 备份记录名字

	RecordName *string `json:"RecordName,omitempty" name:"RecordName"`
	// 查询的cluster名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *DeleteBackupRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBackupRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAllServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAllServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAllServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBindingRequest struct {
	*tchttp.BaseRequest

	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例的命名空间

	InstanceNamespace *string `json:"InstanceNamespace,omitempty" name:"InstanceNamespace"`
	// 绑定的名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 绑定的命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群的名字

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *DeleteBindingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBindingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListBindingsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListBindingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListBindingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBindingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBindingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBindingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListInstancesRequest struct {
	*tchttp.BaseRequest

	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 规格名称

	PlanName *string `json:"PlanName,omitempty" name:"PlanName"`
	// 实例名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 查询的cluster名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *ListInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBindingRequest struct {
	*tchttp.BaseRequest

	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例的命名空间

	InstanceNamespace *string `json:"InstanceNamespace,omitempty" name:"InstanceNamespace"`
	// 绑定的名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 绑定的命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群的名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 参数JSON

	ParametersStr *string `json:"ParametersStr,omitempty" name:"ParametersStr"`
}

func (r *CreateBindingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBindingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实例所在命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 查询的cluster名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateBackupRequest struct {
	*tchttp.BaseRequest

	// 实例名字

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例所在命名空间

	InstanceNamespace *string `json:"InstanceNamespace,omitempty" name:"InstanceNamespace"`
	// 备份类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 备份名字前缀

	NamePrefix *string `json:"NamePrefix,omitempty" name:"NamePrefix"`
	// 实例所属服务名字

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 备份保留天数

	RetainDays *int64 `json:"RetainDays,omitempty" name:"RetainDays"`
	// 定时备份的备份策略

	Schedule *string `json:"Schedule,omitempty" name:"Schedule"`
	// 是否开启定时备份任务

	EnableBackup *bool `json:"EnableBackup,omitempty" name:"EnableBackup"`
	// 查询的cluster名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *UpdateBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest

	// 服务名字

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务版本

	ServiceVersion *string `json:"ServiceVersion,omitempty" name:"ServiceVersion"`
	// 规格名字

	PlanName *string `json:"PlanName,omitempty" name:"PlanName"`
	// 实例名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实例创建参数

	ParamsStr *string `json:"ParamsStr,omitempty" name:"ParamsStr"`
	// 创建者名字

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 是否开启定时备份

	EnableBackup *bool `json:"EnableBackup,omitempty" name:"EnableBackup"`
	// 备份保留时间

	RetainDays *int64 `json:"RetainDays,omitempty" name:"RetainDays"`
	// 备份策略

	Schedule *string `json:"Schedule,omitempty" name:"Schedule"`
	// 实例需要添加的labels

	LabelsStr *string `json:"LabelsStr,omitempty" name:"LabelsStr"`
	// 查询的cluster

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBindingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBindingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBindingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlansRequest struct {
	*tchttp.BaseRequest

	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 服务名字

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 规格名字，支持模糊搜索

	PlanName *string `json:"PlanName,omitempty" name:"PlanName"`
	// 查询的cluster名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *ListPlansRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlansRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlansResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPlansResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBackupRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBackupRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBackupRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListBackupRecordsRequest struct {
	*tchttp.BaseRequest

	// 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 实例名字

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例所在命名空间

	InstanceNamespace *string `json:"InstanceNamespace,omitempty" name:"InstanceNamespace"`
	// 查询的cluster名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *ListBackupRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListBackupRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInstanceRequest struct {
	*tchttp.BaseRequest

	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 实例名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实例所在命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 实例对应的 service 名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 查询的cluster名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *GetInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAllServicesRequest struct {
	*tchttp.BaseRequest

	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询的cluster名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *ListAllServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAllServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListBackupRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListBackupRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListBackupRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest

	// 实例名字

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例所在命名空间

	InstanceNamespace *string `json:"InstanceNamespace,omitempty" name:"InstanceNamespace"`
	// 备份类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 备份名字前缀

	NamePrefix *string `json:"NamePrefix,omitempty" name:"NamePrefix"`
	// 实例所属服务名字

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 备份保留天数

	RetainDays *int64 `json:"RetainDays,omitempty" name:"RetainDays"`
	// 定时备份的备份策略

	Schedule *string `json:"Schedule,omitempty" name:"Schedule"`
	// 是否开启备份

	EnableBackup *bool `json:"EnableBackup,omitempty" name:"EnableBackup"`
	// 查询的cluster名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *CreateBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanRequest struct {
	*tchttp.BaseRequest

	// 服务名字

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 创建者名字

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 规格创建参数

	PlanStr *string `json:"PlanStr,omitempty" name:"PlanStr"`
	// 查询的cluster名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *CreatePlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanRequest struct {
	*tchttp.BaseRequest

	// 服务名字

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 规格更新参数

	PlanStr *string `json:"PlanStr,omitempty" name:"PlanStr"`
	// 查询的cluster名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *UpdatePlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanRequest struct {
	*tchttp.BaseRequest

	// 服务名字

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务规格名字

	PlanName *string `json:"PlanName,omitempty" name:"PlanName"`
	// 查询的cluster名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *DeletePlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListBindingsRequest struct {
	*tchttp.BaseRequest

	// 实例的名字

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例的命名空间

	InstanceNamespace *string `json:"InstanceNamespace,omitempty" name:"InstanceNamespace"`
	// 查询的数量限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询的偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 集群的名字

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *ListBindingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListBindingsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
