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

package v20201116

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DescribeDisksRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个云硬盘ID查询。云硬盘ID形如：`disk-11112222`，此参数的具体格式可参考API[简介](/document/product/362/15633)的ids.N一节）。参数不支持同时指定`DiskIds`和`Filters`。

	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 输出云盘列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 云盘列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据云盘的创建时间排序<br><li>DEADLINE：依据云盘的到期时间排序<br>默认按云盘创建时间排序。

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 云盘详情中是否需要返回云盘绑定的定期快照策略ID，TRUE表示需要返回，FALSE表示不返回。

	ReturnBindAutoSnapshotPolicy *bool `json:"ReturnBindAutoSnapshotPolicy,omitempty" name:"ReturnBindAutoSnapshotPolicy"`
	// 内部参数，用于支持搜索框搜索。

	InnerSearch *string `json:"InnerSearch,omitempty" name:"InnerSearch"`
	// 过滤条件。参数不支持同时指定`DiskIds`和`Filters`。<br><li>disk-usage - Array of String - 是否必填：否 -（过滤条件）按云盘类型过滤。 (SYSTEM_DISK：表示系统盘 | DATA_DISK：表示数据盘)<br><li>disk-charge-type - Array of String - 是否必填：否 -（过滤条件）按照云硬盘计费模式过滤。 (PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费。)<br><li>portable - Array of Bool - 是否必填：否 -（过滤条件）按是否为弹性云盘过滤。 (TRUE：表示弹性云盘 | FALSE：表示非弹性云盘。)<br><li>project-id - Array of Integer - 是否必填：否 -（过滤条件）按云硬盘所属项目ID过滤。<br><li>disk-id - Array of String - 是否必填：否 -（过滤条件）按照云硬盘ID过滤。云盘ID形如：`disk-11112222`。<br><li>disk-name - Array of String - 是否必填：否 -（过滤条件）按照云盘名称过滤。<br><li>disk-type - Array of String - 是否必填：否 -（过滤条件）按照云盘介质类型过滤。(CLOUD_BASIC：表示普通云硬盘 | CLOUD_PREMIUM：表示高性能云硬盘。| CLOUD_SSD：SSD表示SSD云硬盘。)<br><li>disk-state - Array of String - 是否必填：否 -（过滤条件）按照云盘状态过滤。(UNATTACHED：未挂载 | ATTACHING：挂载中 | ATTACHED：已挂载 | DETACHING：解挂中 | EXPANDING：扩容中 | ROLLBACKING：回滚中 | TORECYCLE：待回收。)<br><li>instance-id - Array of String - 是否必填：否 -（过滤条件）按照云盘挂载的云主机实例ID过滤。可根据此参数查询挂载在指定云主机下的云硬盘。<br><li>zone - Array of String - 是否必填：否 -（过滤条件）按照[可用区](/document/api/213/9452#zone)过滤。<br><li>instance-ip-address - Array of String - 是否必填：否 -（过滤条件）按云盘所挂载云主机的内网或外网IP过滤。<br><li>instance-name - Array of String - 是否必填：否 -（过滤条件）按云盘所挂载的实例名称过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Snapshot struct {
	// 快照创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 快照uuid

	SnapUuID *string `json:"SnapUuID,omitempty" name:"SnapUuID"`
	// 快照类型。FULL: 全量快照；INC: 增量快照

	SnapshotType *string `json:"SnapshotType,omitempty" name:"SnapshotType"`
	// 快照实际大小，单位MB

	SnapshotRealSize *uint64 `json:"SnapshotRealSize,omitempty" name:"SnapshotRealSize"`
	// 快照属性最新修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 快照关联云盘的uuid

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
	// 父快照的uuid

	ParentSnapshotUuid *string `json:"ParentSnapshotUuid,omitempty" name:"ParentSnapshotUuid"`
	// 快照所在的zone信息

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 快照关联云盘的大小，单位GB

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 快照的实例ID

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 是否为加密快照。encrypt表示为加密快照

	Encrypt *string `json:"Encrypt,omitempty" name:"Encrypt"`
	// 快照共享的次数

	ShareReference *uint64 `json:"ShareReference,omitempty" name:"ShareReference"`
	// 是否为跨地域复制的快照

	CopyFromRemote *uint64 `json:"CopyFromRemote,omitempty" name:"CopyFromRemote"`
	// 快照关联盘的类型。DISK_SYSTEM表示系统盘；DISK_DATA表示数据盘

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 跨地域复制的快照，表示当前正在复制的地域

	CopyingToRegions []*string `json:"CopyingToRegions,omitempty" name:"CopyingToRegions"`
	// 针对系统盘快照，记录子机的镜像信息

	ExtraInfo *string `json:"ExtraInfo,omitempty" name:"ExtraInfo"`
	// 快照名称

	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
	// 快照删除时间

	DeletedTime *string `json:"DeletedTime,omitempty" name:"DeletedTime"`
	// 快照创建的进度，100表示创建完成

	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`
	// 快照关联云盘的实例ID

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 快照的生命状态

	LifeState *string `json:"LifeState,omitempty" name:"LifeState"`
	// 快照的状态。取值范围：<br><li>NORMAL：正常<br><li>CREATING：创建中<br><li>ROLLBACKING：回滚中<br><li>COPYING_FROM_REMOTE：跨地域复制快照拷贝中。

	SnapshotState *string `json:"SnapshotState,omitempty" name:"SnapshotState"`
	// 快照的到期时间

	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`
	// 是否为永久快照

	IsPermanent *uint64 `json:"IsPermanent,omitempty" name:"IsPermanent"`
	// 快照所属的用户appId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 快照关联盘的简要信息

	Disk *SimpleDisk `json:"Disk,omitempty" name:"Disk"`
	// 快照关联云盘挂载的cvm实例信息

	Instance *SimpleInstance `json:"Instance,omitempty" name:"Instance"`
	// 快照关联的镜像列表

	Images []*Image `json:"Images,omitempty" name:"Images"`
	// 快照关联的镜像个数

	ImageCount *uint64 `json:"ImageCount,omitempty" name:"ImageCount"`
	// 快照在底层的状态。取值范围：<br><li>NORMAL：正常<br><li>DELETED：已删除

	SnapshotRealState *string `json:"SnapshotRealState,omitempty" name:"SnapshotRealState"`
	// 标志快照是否为后台快照，1为后台快照，0为用户快照。

	BackendSnap *uint64 `json:"BackendSnap,omitempty" name:"BackendSnap"`
	// 快照大小，单位为MiB。

	SnapshotSize *uint64 `json:"SnapshotSize,omitempty" name:"SnapshotSize"`
	// 标志是否为cfs快照。

	IsCfsSnapshot *uint64 `json:"IsCfsSnapshot,omitempty" name:"IsCfsSnapshot"`
}

type Filter struct {
	// 过滤键的名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values *string `json:"Values,omitempty" name:"Values"`
}

type DescribeSnapshotsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 快照详情列表

		SnapshotSet []*Snapshot `json:"SnapshotSet,omitempty" name:"SnapshotSet"`
		// 符合条件的快照总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的云硬盘数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 云硬盘的详细信息列表。

		DiskSet []*Disk `json:"DiskSet,omitempty" name:"DiskSet"`
		// 偏移量

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 返回数量

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SimpleDisk struct {
	// 盘关联cvm的类型，取值有CVM, RAW_CVM等

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 是否为共享盘

	Shareable *uint64 `json:"Shareable,omitempty" name:"Shareable"`
	// 云盘名称

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 云盘大小，单位GB

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 云盘id

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 是否为弹性云盘

	Portable *bool `json:"Portable,omitempty" name:"Portable"`
	// 云盘类型。CLOUD_SATA: 普通云硬盘；CLOUD_PREMIUM: 高性能云硬盘；CLOUD_SSD： SSD云硬盘

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 是否为加密盘；取值为encrypt表示加密盘

	Encrypt *string `json:"Encrypt,omitempty" name:"Encrypt"`
	// 云盘续费标记，仅对预付费云盘生效

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 云盘是否随挂载的cvm实例一起销毁

	DeleteWithInstance *uint64 `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
	// 云盘所在的资源池

	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
	// 云盘的状态

	DiskState *string `json:"DiskState,omitempty" name:"DiskState"`
	// 云盘的最近属性修改时间

	ModifyTimeStamp *string `json:"ModifyTimeStamp,omitempty" name:"ModifyTimeStamp"`
	// 云盘到期时间，仅对预付费云盘生效

	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`
	// 云盘的付费类型

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 云盘挂载的cvm实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 云盘用途类型。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
}

type DescribeSnapshotsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 输出快照列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列

	Order *string `json:"Order,omitempty" name:"Order"`
	// 快照列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据快照的创建时间排序<br><li>默认按创建时间排序

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeSnapshotsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSnapshotsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskAssociatedSnapshotsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云盘关联的总快照数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 云盘关联的总快照大小，单位MB

		TotalSize *uint64 `json:"TotalSize,omitempty" name:"TotalSize"`
		// 快照列表详情

		SnapshotSet []*Snapshot `json:"SnapshotSet,omitempty" name:"SnapshotSet"`
		// 本次查询使用的偏移量。

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 本次查询返回数量。

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskAssociatedSnapshotsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskAssociatedSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SimpleInstance struct {
	// 实例UUID

	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
	// 实例所在的宿主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
}

type Placement struct {
	// 实例所属的[可用区](/document/api/213/9452#zone)ID。该参数也可以通过调用  [DescribeZones](/document/product/213/15707) 的返回值中的Zone字段来获取。

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 实例所属项目ID。该参数可以通过调用 [DescribeProject](/document/api/378/4400) 的返回值中的 projectId 字段来获取。不填为默认项目。

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeDiskAssociatedSnapshotsRequest struct {
	*tchttp.BaseRequest

	// 云盘uuid

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
	// 返回数量，默认为20，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件。<br><li>snapshot-id - Array of String - 是否必填：否 -（过滤条件）按照快照的ID过滤。快照ID形如：`snap-11112222`。<br><li>snapshot-name - Array of String - 是否必填：否 -（过滤条件）按照快照名称过滤。<br><li>snapshot-state - Array of String - 是否必填：否 -（过滤条件）按照快照状态过滤。 (NORMAL：正常 | CREATING：创建中 | ROLLBACKING：回滚中。)<br><li>disk-usage - Array of String - 是否必填：否 -（过滤条件）按创建快照的云盘类型过滤。 (SYSTEM_DISK：代表系统盘 | DATA_DISK：代表数据盘。)<br><li>project-id  - Array of String - 是否必填：否 -（过滤条件）按云硬盘所属项目ID过滤。<br><li>disk-id  - Array of String - 是否必填：否 -（过滤条件）按照创建快照的云硬盘ID过滤。<br><li>zone - Array of String - 是否必填：否 -（过滤条件）按照可用区过滤。<br><li>snapshot-type - Array of String - 是否必填：否 -（过滤条件）按快照的类型过滤。 (FULL：代表全量快照 | INC：代表增量快照。)

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDiskAssociatedSnapshotsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskAssociatedSnapshotsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Image struct {
	// 镜像实例ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
}

type Disk struct {
	// 云硬盘ID。

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 云硬盘类型。取值范围：<br><li>SYSTEM_DISK：系统盘<br><li>DATA_DISK：数据盘。

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 付费模式。取值范围：<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：后付费，即按量计费。

	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`
	// 是否为弹性云盘，false表示非弹性云盘，true表示弹性云盘。

	Portable *bool `json:"Portable,omitempty" name:"Portable"`
	// 云硬盘所在的位置。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 云盘是否具备创建快照的能力。取值范围：<br><li>false表示不具备<br><li>true表示具备。

	SnapshotAbility *bool `json:"SnapshotAbility,omitempty" name:"SnapshotAbility"`
	// 云硬盘名称。

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 云硬盘大小。

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 云盘状态。取值范围：<br><li>UNATTACHED：未挂载<br><li>ATTACHING：挂载中<br><li>ATTACHED：已挂载<br><li>DETACHING：解挂中<br><li>EXPANDING：扩容中<br><li>ROLLBACKING：回滚中。

	DiskState *string `json:"DiskState,omitempty" name:"DiskState"`
	// 云盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：SSD表示SSD云硬盘。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 云盘是否挂载到云主机上。取值范围：<br><li>false:表示未挂载<br><li>true:表示已挂载。

	Attached *bool `json:"Attached,omitempty" name:"Attached"`
	// 云硬盘挂载的云主机ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 云硬盘的创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 云硬盘的到期时间。

	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`
	// 云盘是否处于快照回滚状态。取值范围：<br><li>0:表示不处于快照回滚状态<br><li>1:表示处于快照回滚状态。

	Rollbacking *uint64 `json:"Rollbacking,omitempty" name:"Rollbacking"`
	// 云盘快照回滚的进度。

	RollbackPercent *uint64 `json:"RollbackPercent,omitempty" name:"RollbackPercent"`
	// 云盘是否为加密盘。取值范围：<br><li>false:表示非加密盘<br><li>true:表示加密盘。

	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`
	// 云盘已挂载到子机，且子机与云盘都是包年包月。<br><li>true：子机设置了自动续费标识，但云盘未设置<br><li>false：云盘自动续费标识正常。

	AutoRenewFlagError *bool `json:"AutoRenewFlagError,omitempty" name:"AutoRenewFlagError"`
	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费。

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 在云盘已挂载到实例，且实例与云盘都是包年包月的条件下，此字段才有意义。<br><li>true:云盘到期时间早于实例。<br><li>false：云盘到期时间晚于实例。

	DeadlineError *bool `json:"DeadlineError,omitempty" name:"DeadlineError"`
	// 云盘关联的定期快照ID。只有在调用DescribeDisks接口时，入参ReturnBindAutoSnapshotPolicy取值为TRUE才会返回该参数。

	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitempty" name:"AutoSnapshotPolicyIds"`
	// 是否随实例销毁

	DeleteWithInstance *uint64 `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`
	// 是否为共享盘

	Shareable *uint64 `json:"Shareable,omitempty" name:"Shareable"`
	// 用户appId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 盘挂载的母机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 是否为CFS存储池的云盘

	IsCfsDisk *uint64 `json:"IsCfsDisk,omitempty" name:"IsCfsDisk"`
	// 云盘所在存储池ID

	DepotId *uint64 `json:"DepotId,omitempty" name:"DepotId"`
	// 云盘所在存储池名称

	DepotName *string `json:"DepotName,omitempty" name:"DepotName"`
	// 云盘的流控模式。<br><li>IN_BLACK_LIST:黑名单流控<br><li>IN_WHITE_LIST:白名单流控<br><li>AUTO：自动流控<br><li>IN_BLACK_WHITE_LIST: 同时处于黑名单、白名单流控中

	FlowControlState *string `json:"FlowControlState,omitempty" name:"FlowControlState"`
	// 云盘是否在迁移中

	Migrating *bool `json:"Migrating,omitempty" name:"Migrating"`
	// 云盘迁移的进度，只有当Migrating为true时才有意义

	MigratePercent *uint64 `json:"MigratePercent,omitempty" name:"MigratePercent"`
	// 云盘挂载实例的uuid

	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
	// 云盘挂载实例的名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 云盘挂载母机的driver版本，如果云盘未挂载，或查询不到版本号，则显示为“”

	DriverVersion *string `json:"DriverVersion,omitempty" name:"DriverVersion"`
	// 云盘uuid

	DiskUuid *string `json:"DiskUuid,omitempty" name:"DiskUuid"`
	// 云盘实际已使用容量，单位GB

	DiskSizeUsed *uint64 `json:"DiskSizeUsed,omitempty" name:"DiskSizeUsed"`
	// 当前时间与云盘到期时间相差的天数，仅去预付费盘有效

	DifferDaysOfDeadline *int64 `json:"DifferDaysOfDeadline,omitempty" name:"DifferDaysOfDeadline"`
	// 云盘所属资源池

	DiskStoragePoolGroup *string `json:"DiskStoragePoolGroup,omitempty" name:"DiskStoragePoolGroup"`
	// 共享盘挂载的云服务器实例信息

	ShareAttachedInstanceSet []*SimpleInstance `json:"ShareAttachedInstanceSet,omitempty" name:"ShareAttachedInstanceSet"`
}
