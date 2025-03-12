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

package v20230915

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DescribeTurboTaskInfoRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTurboTaskInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboTaskInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VirtualPrivateCloud struct {
	// VPC ID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网ID。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 私有网络IP。

	PrivateIpAddresses *string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
}

type DescribeTurboFsCvmInstancesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。 <br><li>component-type - Array of String - 是否必填：否 -（过滤条件）按组件类型过滤，取值范围：mon, mgs, mds, oss, nfsd。 <br><li>instance-id - Array of String - 是否必填：否 -（过滤条件）按实例ID过滤。 <br><li>private-ip-addresses - Array of String - 是否必填：否 -（过滤条件）按实例私有网络IP过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 文件系统ID。

	FsId *string `json:"FsId,omitempty" name:"FsId"`
	// 文件系统集群ID。

	SetId *string `json:"SetId,omitempty" name:"SetId"`
}

func (r *DescribeTurboFsCvmInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboFsCvmInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryTurboTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryTurboTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryTurboTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键。

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值。

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TurboTask struct {
	// 任务类型

	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
	// 任务状态

	StatusName *string `json:"StatusName,omitempty" name:"StatusName"`
	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 类型ID

	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
	// 拓展信息

	Param *string `json:"Param,omitempty" name:"Param"`
	// 状态ID

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 创建时间

	Ctime *string `json:"Ctime,omitempty" name:"Ctime"`
	// 修改时间

	Utime *string `json:"Utime,omitempty" name:"Utime"`
	// fsid

	TaskIndex *string `json:"TaskIndex,omitempty" name:"TaskIndex"`
	// 错误原因

	Result *string `json:"Result,omitempty" name:"Result"`
}

type DescribeTurboImageConfigRequest struct {
	*tchttp.BaseRequest

	// 存储类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
}

func (r *DescribeTurboImageConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboImageConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTurboCreateFailedResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateTurboCreateFailedResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTurboCreateFailedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TurboImageInfo struct {
	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像操作系统

	ImageOs *string `json:"ImageOs,omitempty" name:"ImageOs"`
	// server or client

	Type *string `json:"Type,omitempty" name:"Type"`
	// 存储类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 架构

	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
	// 是否在使用中

	Used *int64 `json:"Used,omitempty" name:"Used"`
}

type TurboOverview struct {
	// 文件系统存储量

	FsSize *uint64 `json:"FsSize,omitempty" name:"FsSize"`
	// 文件系统容量

	Capacity *uint64 `json:"Capacity,omitempty" name:"Capacity"`
	// 文件系统数量

	FsCount *uint64 `json:"FsCount,omitempty" name:"FsCount"`
	// 用户数量

	UserCount *uint64 `json:"UserCount,omitempty" name:"UserCount"`
	// zone信息

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DescribeTurboDiskConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点类型

		NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
		// 磁盘类型

		DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
		// 系统盘类型

		SysDiskType *string `json:"SysDiskType,omitempty" name:"SysDiskType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTurboDiskConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboDiskConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScaleUpFileSystemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// fsid

		FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
		// 扩容大小

		TargetCapacity *uint64 `json:"TargetCapacity,omitempty" name:"TargetCapacity"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScaleUpFileSystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScaleUpFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 过滤键的名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type CreateTurboVMConfigRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTurboVMConfigRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTurboVMConfigRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllTurboTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务列表

		TurboTaskSet []*TurboTask `json:"TurboTaskSet,omitempty" name:"TurboTaskSet"`
		// 任务数量

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllTurboTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllTurboTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTurboFsCvmInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例总数。

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 实例详情。

		InstanceSet *TurboFsInstance `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTurboFsCvmInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboFsCvmInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Cfs3Amount struct {
	// 挂载点ID。

	Mountid *string `json:"Mountid,omitempty" name:"Mountid"`
	// 文件系统ID。

	Fsid *string `json:"Fsid,omitempty" name:"Fsid"`
	// 网络类型。

	Nettype *string `json:"Nettype,omitempty" name:"Nettype"`
	// vpc id。

	Vpcid *string `json:"Vpcid,omitempty" name:"Vpcid"`
	// Unvpcid。

	Unvpcid *string `json:"Unvpcid,omitempty" name:"Unvpcid"`
	// 子网ID。

	Subnetid *string `json:"Subnetid,omitempty" name:"Subnetid"`
	// Unsubnetid。

	Unsubnetid *string `json:"Unsubnetid,omitempty" name:"Unsubnetid"`
	// 私有网络IP。

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 状态。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 创建时间。

	Ctime *string `json:"Ctime,omitempty" name:"Ctime"`
	// 更新时间。

	Utime *string `json:"Utime,omitempty" name:"Utime"`
}

type InstanceDisk struct {
	// 云硬盘类型。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 云硬盘大小，单位GB。

	DiskSize *string `json:"DiskSize,omitempty" name:"DiskSize"`
	// 云硬盘uuid。

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
}

type UpdateTurboCreateFailedRequest struct {
	*tchttp.BaseRequest

	// 文件系统ID

	FsId *string `json:"FsId,omitempty" name:"FsId"`
}

func (r *UpdateTurboCreateFailedRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTurboCreateFailedRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCfsFileSystemRequest struct {
	*tchttp.BaseRequest

	// 文件系统 ID

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DeleteCfsFileSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCfsFileSystemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllTurboTasksRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 起始参数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 允许条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAllTurboTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllTurboTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTurboClusterStatusRequest struct {
	*tchttp.BaseRequest

	// 文件系统ID

	FsId *string `json:"FsId,omitempty" name:"FsId"`
}

func (r *DescribeTurboClusterStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboClusterStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTurboFsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。 <br><li>zone-id - Array of Integer - 是否必填：否 -（过滤条件）按可用区ID过滤。 <br><li>fs-id - Array of String - 是否必填：否 -（过滤条件）按文件系统ID过滤。 <br><li>uin - Array of String - 是否必填：否 -（过滤条件）按用户uin过滤。 <br><li>app-id - Array of Integer - 是否必填：否 -（过滤条件）按用户appId过滤。 <br><li>status - Array of Integer - 是否必填：否 -（过滤条件）按文件系统状态过滤。（0：创建中；1：可使用；2：删除中；3：已删除；4：失败；10：扩容中） <br><li>account-name - Array of String - 是否必填：否 -（过滤条件）按用户名过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// ASC：升序；DESC：降序。

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeTurboFsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboFsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTurbofsAnalysisRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTurbofsAnalysisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurbofsAnalysisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTurboDiskConfigRequest struct {
	*tchttp.BaseRequest

	// 存储类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 磁盘类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 系统盘类型

	SysDiskType *string `json:"SysDiskType,omitempty" name:"SysDiskType"`
}

func (r *ModifyTurboDiskConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTurboDiskConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTurboImageConfigRequest struct {
	*tchttp.BaseRequest

	// 存储类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// turbo镜像类型（server/client）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像操作系统

	ImageOS *string `json:"ImageOS,omitempty" name:"ImageOS"`
}

func (r *ModifyTurboImageConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTurboImageConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTurboClientInfoRequest struct {
	*tchttp.BaseRequest

	// 文件系统ID。

	FsId *string `json:"FsId,omitempty" name:"FsId"`
}

func (r *DescribeTurboClientInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboClientInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryTurboTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *RetryTurboTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryTurboTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTurboTaskInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTurboTaskInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTurboVMConfigRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTurboVMConfigRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTurboVMConfigRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTurboTaskByTaskIdRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTurboTaskByTaskIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboTaskByTaskIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTurboImageConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTurboImageConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTurboImageConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfs3AmountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数。

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 挂载点列表。

		Cfs3AmountSet []*Cfs3Amount `json:"Cfs3AmountSet,omitempty" name:"Cfs3AmountSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCfs3AmountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfs3AmountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTurboFsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的总数。

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 文件系统详情。

		FsSet []*TurboFs `json:"FsSet,omitempty" name:"FsSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTurboFsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboFsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTurboVMConfigRequest struct {
	*tchttp.BaseRequest

	// 存储类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
}

func (r *DescribeTurboVMConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboVMConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScaleUpFileSystemRequest struct {
	*tchttp.BaseRequest

	// 扩容大小

	TargetCapacity *uint64 `json:"TargetCapacity,omitempty" name:"TargetCapacity"`
	// fsid

	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *ScaleUpFileSystemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScaleUpFileSystemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TurboVMInfo struct {
	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 虚拟机类型

	VMConfig *string `json:"VMConfig,omitempty" name:"VMConfig"`
	// 优先级

	Priority *string `json:"Priority,omitempty" name:"Priority"`
}

type DescribeTurboClusterStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误信息

		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
		// 状态信息

		CreateMsg *string `json:"CreateMsg,omitempty" name:"CreateMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTurboClusterStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboClusterStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTurboImageConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务端镜像列表

		ServerImages []*TurboImageInfo `json:"ServerImages,omitempty" name:"ServerImages"`
		// mon客户列表

		ClientImages []*TurboImageInfo `json:"ClientImages,omitempty" name:"ClientImages"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTurboImageConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboImageConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTurboTaskByTaskIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回任务信息

		TurboTask *TurboTask `json:"TurboTask,omitempty" name:"TurboTask"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTurboTaskByTaskIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboTaskByTaskIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTurboTaskStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTurboTaskStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTurboTaskStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClientInfo struct {
	// 客户端IP。

	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`
}

type Instance struct {
	// 实例uuid。

	VmUuid *string `json:"VmUuid,omitempty" name:"VmUuid"`
	// 用户appId。

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 实例所在的母机IP。

	Ip1 *string `json:"Ip1,omitempty" name:"Ip1"`
	// 实例所在region。

	RegionApCode *string `json:"RegionApCode,omitempty" name:"RegionApCode"`
	// 实例ID。

	DiskInst *string `json:"DiskInst,omitempty" name:"DiskInst"`
	// 实例ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例私有IP。

	PrivateIpAddresses *string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 实例状态。

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 实例机型。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例CPU核数。

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 实例内存。

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 实例系统盘信息。

	SystemDisk *InstanceDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 实例数据盘信息。

	DataDisks []*InstanceDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 实例私有网络信息。

	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
	// 实例镜像ID。

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 实例OS名称。

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 实例创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 实例绑定的标签。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 实例所有可用区。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例所属turbo组件类型。

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
}

type DeleteCfsFileSystemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCfsFileSystemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCfsFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTurbofsAnalysisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 纵览信息

		Overview *TurboOverview `json:"Overview,omitempty" name:"Overview"`
		// zone下统计信息

		OverviewByZone []*TurboOverview `json:"OverviewByZone,omitempty" name:"OverviewByZone"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTurbofsAnalysisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurbofsAnalysisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTurboDiskConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTurboDiskConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTurboDiskConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TurboFs struct {
	// 是否存在伸缩策略。

	ExistsScaleupPolicy *string `json:"ExistsScaleupPolicy,omitempty" name:"ExistsScaleupPolicy"`
	// 用户账号uin。

	CustomerAccountUin *string `json:"CustomerAccountUin,omitempty" name:"CustomerAccountUin"`
	// 用户账号名称。

	CustomerAccountName *string `json:"CustomerAccountName,omitempty" name:"CustomerAccountName"`
	// 用户昵称。

	CustomerAccountNickName *string `json:"CustomerAccountNickName,omitempty" name:"CustomerAccountNickName"`
	// 用户等级。

	CustomerAccountGrade *string `json:"CustomerAccountGrade,omitempty" name:"CustomerAccountGrade"`
	// 用户所属行业1。

	IncomeIndustry1Name *string `json:"IncomeIndustry1Name,omitempty" name:"IncomeIndustry1Name"`
	// 用户所属行业2。

	IncomeIndustry2Name *string `json:"IncomeIndustry2Name,omitempty" name:"IncomeIndustry2Name"`
	// CustomerSystemArchitect。

	CustomerSystemArchitect *string `json:"CustomerSystemArchitect,omitempty" name:"CustomerSystemArchitect"`
	// CustomerOperationManager。

	CustomerOperationManager *string `json:"CustomerOperationManager,omitempty" name:"CustomerOperationManager"`
	// 是否为主账号。

	UserIsOwner *uint64 `json:"UserIsOwner,omitempty" name:"UserIsOwner"`
	// UserIsOwnerDesc

	UserIsOwnerDesc *string `json:"UserIsOwnerDesc,omitempty" name:"UserIsOwnerDesc"`
	// UserIsAuthenticate

	UserIsAuthenticate *string `json:"UserIsAuthenticate,omitempty" name:"UserIsAuthenticate"`
	// UserIsAuthenticateDesc

	UserIsAuthenticateDesc *string `json:"UserIsAuthenticateDesc,omitempty" name:"UserIsAuthenticateDesc"`
	// 用户邮箱。

	UserMail *string `json:"UserMail,omitempty" name:"UserMail"`
	// 可能区名称。

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 文件系统状态。（0：创建中；1：可使用；2：删除中；3：已删除；4：失败；10：扩容中）

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 文件系统状态描述。

	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
	// 文件系统ID。

	FsId *string `json:"FsId,omitempty" name:"FsId"`
	// BackFsId。

	BackFsId *string `json:"BackFsId,omitempty" name:"BackFsId"`
	// 文件系统别名。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 固定为turbo。

	Version *string `json:"Version,omitempty" name:"Version"`
	// 文件系统类型。

	FsType *uint64 `json:"FsType,omitempty" name:"FsType"`
	// FsTypeDesc

	FsTypeDesc *string `json:"FsTypeDesc,omitempty" name:"FsTypeDesc"`
	// Protocal

	Protocal *uint64 `json:"Protocal,omitempty" name:"Protocal"`
	// ProtocalDesc

	ProtocalDesc *string `json:"ProtocalDesc,omitempty" name:"ProtocalDesc"`
	// Vid

	Vid *string `json:"Vid,omitempty" name:"Vid"`
	// SetId

	SetId *string `json:"SetId,omitempty" name:"SetId"`
	// 可用区ID。

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 读流量。

	Read *uint64 `json:"Read,omitempty" name:"Read"`
	// 写流量。

	Write *uint64 `json:"Write,omitempty" name:"Write"`
	// 读流量。

	ReadSingle *uint64 `json:"ReadSingle,omitempty" name:"ReadSingle"`
	// 写流量。

	WriteSingle *uint64 `json:"WriteSingle,omitempty" name:"WriteSingle"`
	// 容量。

	Capacity *float64 `json:"Capacity,omitempty" name:"Capacity"`
	// 文件系统大小。

	FsSize *float64 `json:"FsSize,omitempty" name:"FsSize"`
	// 已使用大小。

	CostSize *float64 `json:"CostSize,omitempty" name:"CostSize"`
	// GroupId

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 用户appId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 用户账号。

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 是否为内部用户。

	IsInnerCustomer *string `json:"IsInnerCustomer,omitempty" name:"IsInnerCustomer"`
	// 地域名。

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
}

type TurboFsInstance struct {
	// oss组件实例列表。

	OssComponents []*Instance `json:"OssComponents,omitempty" name:"OssComponents"`
	// mon组件实例列表。

	MonComponents []*Instance `json:"MonComponents,omitempty" name:"MonComponents"`
	// mds组件实例列表。

	MdsComponents []*Instance `json:"MdsComponents,omitempty" name:"MdsComponents"`
	// mgs组件实例列表。

	MgsComponents []*Instance `json:"MgsComponents,omitempty" name:"MgsComponents"`
}

type DeleteTurboVMConfigRuleRequest struct {
	*tchttp.BaseRequest

	// 存储类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 优先级

	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
}

func (r *DeleteTurboVMConfigRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTurboVMConfigRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTurboClientInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 客户端列表。

		ClientInfo []*ClientInfo `json:"ClientInfo,omitempty" name:"ClientInfo"`
		// 报错详情。

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTurboClientInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboClientInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTurboTaskStateRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ModifyTurboTaskStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTurboTaskStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTurboVMConfigRuleRequest struct {
	*tchttp.BaseRequest

	// 存储类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 虚拟机类型

	VMType *string `json:"VMType,omitempty" name:"VMType"`
	// 虚拟机规格

	Spec *string `json:"Spec,omitempty" name:"Spec"`
	// 优先级

	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
}

func (r *CreateTurboVMConfigRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTurboVMConfigRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfs3AmountsRequest struct {
	*tchttp.BaseRequest

	// 文件系统ID。

	FsId *string `json:"FsId,omitempty" name:"FsId"`
}

func (r *DescribeCfs3AmountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfs3AmountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTurboDiskConfigRequest struct {
	*tchttp.BaseRequest

	// 存储类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
}

func (r *DescribeTurboDiskConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboDiskConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTurboVMConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟机列表

		VMConfigList []*TurboVMInfo `json:"VMConfigList,omitempty" name:"VMConfigList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTurboVMConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTurboVMConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageInfo struct {
	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像操作系统

	ImageOS *string `json:"ImageOS,omitempty" name:"ImageOS"`
}
