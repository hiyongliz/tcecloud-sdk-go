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

package v20210809

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type ModifyInstanceReleaseTimeRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如： tdsql-xxxxxxxx 或 tdsqlshard-xxxxxxxx 。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 释放时间 ，格式为 2021-01-01 12:12:12

	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`
}

func (r *ModifyInstanceReleaseTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceReleaseTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSpecConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置插入后的Id

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSpecConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSpecConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBindExclusiveGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源池总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 租户资源池绑定列表

		BindExclusiveGroupSet []*BindExclusiveGroup `json:"BindExclusiveGroupSet,omitempty" name:"BindExclusiveGroupSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBindExclusiveGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBindExclusiveGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindExclusiveGroupRequest struct {
	*tchttp.BaseRequest

	// 待绑定的独享资源池Id。

	FenceId *string `json:"FenceId,omitempty" name:"FenceId"`
	// 独享资源池绑定的租户 AppId。

	TenantAppIds []*int64 `json:"TenantAppIds,omitempty" name:"TenantAppIds"`
	// 独享资源池绑定的租户 AppId。兼容老版本

	TenantAppId *int64 `json:"TenantAppId,omitempty" name:"TenantAppId"`
}

func (r *BindExclusiveGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindExclusiveGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTenantDBInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例详细信息列表。

		InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTenantDBInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTenantDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelFlowRequest struct {
	*tchttp.BaseRequest

	// 流程Id

	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *CancelFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExclusiveGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 独享资源池总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 独享资源池详情列表。

		ExclusiveGroupSet []*ExclusiveGroup `json:"ExclusiveGroupSet,omitempty" name:"ExclusiveGroupSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExclusiveGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExclusiveGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExclusiveGroupsRequest struct {
	*tchttp.BaseRequest

	// 参数过滤查询

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页参数。页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页参数。页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeExclusiveGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExclusiveGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWorkFlowsRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个流程ID查询。每次请求的流程的上限为100。参数不支持同时指定 `Ids` 和 `Filters` 。

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
	// 创建时间范围起点，精确到天

	CreateTimeBegin *string `json:"CreateTimeBegin,omitempty" name:"CreateTimeBegin"`
	// 创建时间范围终点，精确到天

	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" name:"CreateTimeEnd"`
	// 偏移，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件<br/>支持的key：app-id、flow-name、status、process-instance-id<br/>每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定FlowIds和Filters。其中process-instance-id为taskcenter侧的标识。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序依据，比如ID，FlowName，CreateTime。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序规则，升序还是降序，DESC为降序，ASC为升序。

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeWorkFlowsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWorkFlowsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如： `tdsql-xxxxxxxx` 或 `tdsqlshard-xxxxxxxx`

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DestroyInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBEnginesRequest struct {
	*tchttp.BaseRequest

	// 根据引擎版本查询

	Versions []*string `json:"Versions,omitempty" name:"Versions"`
	// 返回数量，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询偏移，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDBEnginesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBEnginesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActivateInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ActivateInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActivateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindExclusiveGroup struct {
	// 独享资源池Id。

	FenceId *string `json:"FenceId,omitempty" name:"FenceId"`
	// 资源池内的设备数量

	DeviceCount *int64 `json:"DeviceCount,omitempty" name:"DeviceCount"`
	// 地域ID

	Region *string `json:"Region,omitempty" name:"Region"`
	// 已绑定的租户Id

	BondAppId *int64 `json:"BondAppId,omitempty" name:"BondAppId"`
	// 绑定时间

	BindTime *string `json:"BindTime,omitempty" name:"BindTime"`
	// 资源池名称

	FenceName *string `json:"FenceName,omitempty" name:"FenceName"`
	// 租户账号名称

	TenantAccount *string `json:"TenantAccount,omitempty" name:"TenantAccount"`
	// 租户账号昵称

	TenantNickname *string `json:"TenantNickname,omitempty" name:"TenantNickname"`
}

type BindExclusiveGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindExclusiveGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindExclusiveGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSpecConfigRequest struct {
	*tchttp.BaseRequest

	// 待删除的规格ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteSpecConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSpecConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceRecycleNetRecycleTimeRequest struct {
	*tchttp.BaseRequest

	// 网络资源ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 回收时间，格式为 2021-01-01 12:12:12

	RecycleTime *string `json:"RecycleTime,omitempty" name:"RecycleTime"`
}

func (r *ModifyInstanceRecycleNetRecycleTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceRecycleNetRecycleTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流程ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流程ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigrateInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSaleZoneConfigRequest struct {
	*tchttp.BaseRequest

	// 主可用区ID

	MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`
	// 从可用区ID

	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`
	// 集中式实例售卖开关，默认为False，不开启售卖

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 分布式式实例售卖开关，默认为False，不开启售卖

	ShardEnable *bool `json:"ShardEnable,omitempty" name:"ShardEnable"`
}

func (r *AddSaleZoneConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSaleZoneConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTenantDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个实例ID查询。实例ID形如： tdsql-xxxxxxxx 或 tdsqlshard-xxxxxxxx 。 每次请求的实例的上限为100。参数不支持同时指定 InstanceIds 和 Filters 。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 偏移，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件 支持的key：instance-id、vip、app-id、serial-id、instance-name 、status、fence-id、fence-name。 每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定InstanceIds和Filters。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 实例名称，进行模糊查询

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 排序列名，支持CreateTime

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序方式，DESC、ASC， 默认ASC升序

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeTenantDBInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTenantDBInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTenantDBInstanceDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例名称

		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
		// 枚举值，0 - 创建中，1 - 流程处理中，2 - 运行中，3 - 未初始化, 4 - 初始化中，5 - 实例删除中，-1 - 已隔离

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 状态描述

		StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
		// 所属租户appId

		AppId *int64 `json:"AppId,omitempty" name:"AppId"`
		// 数字的 VpcId

		NumericVpcId *int64 `json:"NumericVpcId,omitempty" name:"NumericVpcId"`
		// 数字的 SubnetId

		NumericSubnetId *int64 `json:"NumericSubnetId,omitempty" name:"NumericSubnetId"`
		// 实例访问 vip 地址

		Vip *string `json:"Vip,omitempty" name:"Vip"`
		// 实例访问 vport

		Vport *int64 `json:"Vport,omitempty" name:"Vport"`
		// 后台的 setId 或 groupId

		SerialId *string `json:"SerialId,omitempty" name:"SerialId"`
		// 付费模式，取值范围：`prepaid`：表示预付费，即包年包月（TCE暂不支持）；`postpaid`：表示后付费，即按量计费

		PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
		// 引擎类型，mysql/mariadb

		DBEngine *string `json:"DBEngine,omitempty" name:"DBEngine"`
		// 引擎版本

		DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`
		// 实例cpu核数（分布式实例为除创建中和待删除分片外的其他分片总和）

		Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
		// 实例内存，单位GB（分布式实例为除创建中和待删除分片外的其他分片总和）

		Memory *int64 `json:"Memory,omitempty" name:"Memory"`
		// 实例数据盘，单位GB（分布式实例为除创建中和待删除分片外的其他分片总和）

		Storage *int64 `json:"Storage,omitempty" name:"Storage"`
		// 是否分布式实例，1为分布式实例，0为集中式实例

		Shard *int64 `json:"Shard,omitempty" name:"Shard"`
		// 实例分片数，仅分布式实例有效

		ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`
		// 实例主从节点数

		NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
		// 实例 Cpu 架构，取值范围： ARM/X86

		CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
		// 实例创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 实例所在地域

		Region *string `json:"Region,omitempty" name:"Region"`
		// 是否临时实例，取值范围：0 : 常规实例，1: 临时实例(回档中)，2: 临时实例(回档成功)，3:临时实例(回档失败)，4: 延时删除的扩容旧实例，5: 等待删除的切换后的临时实例

		IsTmp *int64 `json:"IsTmp,omitempty" name:"IsTmp"`
		// 独享资源池id

		ExClusterId *string `json:"ExClusterId,omitempty" name:"ExClusterId"`
		// 当前流程信息，可能返回 null，表示无相关流程

		CurrentFlow *WorkFlowBrief `json:"CurrentFlow,omitempty" name:"CurrentFlow"`
		// 可以执行的操作。取值范围 IsolateInstance, ActivateInstance, DestroyInstance和MigrateInstance

		AllowedActions []*string `json:"AllowedActions,omitempty" name:"AllowedActions"`
		// 临时实例释放时间

		ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`
		// 临时实例类型 isolateInstance隔离实例 expandInstance扩容实例 rollbackInstance归档实例

		TempInstanceType *string `json:"TempInstanceType,omitempty" name:"TempInstanceType"`
		// 分片信息。

		Shards []*ShardInfo `json:"Shards,omitempty" name:"Shards"`
		// 存储空间使用率。

		StorageUsage *string `json:"StorageUsage,omitempty" name:"StorageUsage"`
		// 实例ID

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 隔离类型。

		IsolateType *int64 `json:"IsolateType,omitempty" name:"IsolateType"`
		// 隔离时间。

		Isolatetime *string `json:"Isolatetime,omitempty" name:"Isolatetime"`
		// 资源池名称

		ExClusterName *string `json:"ExClusterName,omitempty" name:"ExClusterName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTenantDBInstanceDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTenantDBInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSpecConfigRequest struct {
	*tchttp.BaseRequest

	// Cpu分配值，单位为1/100核

	RequestCpu *int64 `json:"RequestCpu,omitempty" name:"RequestCpu"`
	// Cpu展示值，单位为核（注意若开启 cpu 限制开关，将会把实例可用的cpu限制到这个规格）

	LimitCpu *int64 `json:"LimitCpu,omitempty" name:"LimitCpu"`
	// 内存大小，单位GB

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 可选数据盘下限，单位GB

	MinStorage *int64 `json:"MinStorage,omitempty" name:"MinStorage"`
	// 可选数据盘上限，单位GB

	MaxStorage *int64 `json:"MaxStorage,omitempty" name:"MaxStorage"`
	// 集中式实例售卖开关，默认为False，不开启售卖

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 分布式式实例售卖开关，默认为False，不开启售卖

	ShardEnable *bool `json:"ShardEnable,omitempty" name:"ShardEnable"`
}

func (r *AddSpecConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSpecConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如： `tdsql-xxxxxxxx` 或 `tdsqlshard-xxxxxxxx` 。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 切换起始时间，格式为 2021-01-01 12:12:12

	SwitchStartTime *string `json:"SwitchStartTime,omitempty" name:"SwitchStartTime"`
	// 切换结束时间，格式为 2021-01-01 12:12:12

	SwitchEndTime *string `json:"SwitchEndTime,omitempty" name:"SwitchEndTime"`
	// 迁移到指定的目标机器上（通过DescribeDevices查询或在运营端设备资源页面查看），机器cpu架构需要与实例原始架构相同，IP个数需要与实例节点数相同

	TargetDevices []*string `json:"TargetDevices,omitempty" name:"TargetDevices"`
	// 分片ID，形如： shard-xxx

	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
}

func (r *MigrateInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrateInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WorkFlow struct {
	// 流程Id，每个地域内唯一。

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 流程所属租户 AppId。

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 流程名称。

	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`
	// taskcenter 侧的流程调度ID。

	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" name:"ProcessInstanceId"`
	// 流程执行状态。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 开始调度时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 下次调度时间。

	NextRunTime *string `json:"NextRunTime,omitempty" name:"NextRunTime"`
	// 可以执行的操作。取值范围`RetryFlow`、`CancelFlow`

	AllowedActions []*string `json:"AllowedActions,omitempty" name:"AllowedActions"`
	// 流程输入参数信息

	Input *string `json:"Input,omitempty" name:"Input"`
	// 流程调度控制参数和临时参数

	Data *string `json:"Data,omitempty" name:"Data"`
	// 地域ID

	Region *string `json:"Region,omitempty" name:"Region"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ModifyExclusiveGroupVisibilityRequest struct {
	*tchttp.BaseRequest

	// 资源池ID。默认资源池ID为空

	FenceId *string `json:"FenceId,omitempty" name:"FenceId"`
	// 该资源池是否所有人可见。0-不可见；1-可见

	Visibility *int64 `json:"Visibility,omitempty" name:"Visibility"`
}

func (r *ModifyExclusiveGroupVisibilityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyExclusiveGroupVisibilityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShardInfoBrief struct {
	// 分片ID，数据库中的行数

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 分片所属的实例的id

	ShardInstanceId *int64 `json:"ShardInstanceId,omitempty" name:"ShardInstanceId"`
	// 分片序列号

	ShardSerialId *string `json:"ShardSerialId,omitempty" name:"ShardSerialId"`
	// 分片uuid，分片唯一标识

	ShardResourceId *string `json:"ShardResourceId,omitempty" name:"ShardResourceId"`
	// 分片状态，0为创建中，2为运行中，1为流程中，-1为以隔离，-2为已删除

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 分片创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeInstanceRecycleNetInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网络资源详情列表。

		NetInfoSet []*RecycleNetInfo `json:"NetInfoSet,omitempty" name:"NetInfoSet"`
		// 独享资源池总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceRecycleNetInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceRecycleNetInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyExclusiveGroupNameRequest struct {
	*tchttp.BaseRequest

	// 资源池ID。缺省资源池ID为空

	FenceId *string `json:"FenceId,omitempty" name:"FenceId"`
	// 资源池名称。不能为空

	FenceName *string `json:"FenceName,omitempty" name:"FenceName"`
}

func (r *ModifyExclusiveGroupNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyExclusiveGroupNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBindExclusiveGroupsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件 支持的key：app-id、fence-id每次请求的Filters的上限为10，Filter.Values的上限为5。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 以什么某一项进行排序，比如绑定时间BindTime

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序方向，升序还是降序，DESC降序，ASC升序

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeBindExclusiveGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBindExclusiveGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWorkFlowsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的流程总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 流程详细信息列表。

		WorkFlowSet []*WorkFlow `json:"WorkFlowSet,omitempty" name:"WorkFlowSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWorkFlowsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWorkFlowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySpecConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySpecConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySpecConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AvailableResourceRange struct {
	// 资源下限

	Min *int64 `json:"Min,omitempty" name:"Min"`
	// 资源上限

	Max *int64 `json:"Max,omitempty" name:"Max"`
}

type ModifySaleZoneConfigStatusRequest struct {
	*tchttp.BaseRequest

	// 待修改的规则ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 集中式实例售卖开关

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 分布式式实例售卖开关

	ShardEnable *bool `json:"ShardEnable,omitempty" name:"ShardEnable"`
}

func (r *ModifySaleZoneConfigStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySaleZoneConfigStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaleZoneConfig struct {
	// 售卖规则Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 地域ID

	Region *string `json:"Region,omitempty" name:"Region"`
	// 主可用区ID

	MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`
	// 从可用区ID

	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`
	// 集中式实例售卖开关

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 分布式实例售卖开关

	ShardEnable *bool `json:"ShardEnable,omitempty" name:"ShardEnable"`
}

type DescribeSpecConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的售卖规格总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 售卖规格详细信息列表。

		SpecConfigSet []*SpecConfig `json:"SpecConfigSet,omitempty" name:"SpecConfigSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSpecConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpecConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySpecConfigRequest struct {
	*tchttp.BaseRequest

	// 待修改的规格ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// Cpu分配值，单位为1/100核，不可大于 LimitCpu 核数

	RequestCpu *int64 `json:"RequestCpu,omitempty" name:"RequestCpu"`
	// 可选数据盘下限，单位GB

	MinStorage *int64 `json:"MinStorage,omitempty" name:"MinStorage"`
	// 可选数据盘上限，单位GB

	MaxStorage *int64 `json:"MaxStorage,omitempty" name:"MaxStorage"`
	// 集中式实例售卖开关

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 分布式式实例售卖开关

	ShardEnable *bool `json:"ShardEnable,omitempty" name:"ShardEnable"`
}

func (r *ModifySpecConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySpecConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Database struct {
	// 是否存活

	IsAlive *int64 `json:"IsAlive,omitempty" name:"IsAlive"`
	// 延迟时间

	DelaySecond *int64 `json:"DelaySecond,omitempty" name:"DelaySecond"`
	// IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 是否为主机

	IsMaster *int64 `json:"IsMaster,omitempty" name:"IsMaster"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type DescribeSaleZoneConfigsRequest struct {
	*tchttp.BaseRequest

	// 根据 ID 查询售卖规则

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
	// 返回数量，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询偏移，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSaleZoneConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSaleZoneConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBEngine struct {
	// 引擎版本条目Id。

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 引擎类型，取值范围：MariaDB、Percona、MySQL

	Type *string `json:"Type,omitempty" name:"Type"`
	// 内部标识用的引擎版本。

	Version *string `json:"Version,omitempty" name:"Version"`
	// 显示引擎版本。

	DisplayVersion *string `json:"DisplayVersion,omitempty" name:"DisplayVersion"`
	// 显示名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述信息。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 集中式实例售卖开关

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 分布式式实例售卖开关

	ShardEnable *bool `json:"ShardEnable,omitempty" name:"ShardEnable"`
	// 地域ID

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeAvailableResourceRangeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可配置到规格表中的CPU范围，单位：核

		Cpu *AvailableResourceRange `json:"Cpu,omitempty" name:"Cpu"`
		// 可配置到规格表中的内存范围，单位：GB

		Memory *AvailableResourceRange `json:"Memory,omitempty" name:"Memory"`
		// 可配置到规格表中的数据盘范围，单位：GB

		Storage *AvailableResourceRange `json:"Storage,omitempty" name:"Storage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableResourceRangeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableResourceRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAvailableResourceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如： tdsql-xxxxxxxx 或 tdsqlshard-xxxxxxxx

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 分片ID，形如： shard-xxx

	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`
}

func (r *DescribeInstanceAvailableResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceAvailableResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSaleZoneConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSaleZoneConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSaleZoneConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActivateInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如： `tdsql-xxxxxxxx` 或 `tdsqlshard-xxxxxxxx`

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ActivateInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActivateInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindExclusiveGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindExclusiveGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindExclusiveGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWorkFlowTaskInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流程任务详细信息列表。

		TaskInfoSet []*WorkFlowTaskInfo `json:"TaskInfoSet,omitempty" name:"TaskInfoSet"`
		// 流程总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWorkFlowTaskInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWorkFlowTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Machine struct {
	// ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// IDC

	Idc *string `json:"Idc,omitempty" name:"Idc"`
	// 机架信息，老版本该值可能为空

	Frame *string `json:"Frame,omitempty" name:"Frame"`
	// 资源池

	FenceId *string `json:"FenceId,omitempty" name:"FenceId"`
	// cpu架构

	Arch *string `json:"Arch,omitempty" name:"Arch"`
}

type ShardInfo struct {
	// 分片ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 分片ShardSerialId

	ShardSerialId *string `json:"ShardSerialId,omitempty" name:"ShardSerialId"`
	// 分片ShardResourceId

	ShardResourceId *string `json:"ShardResourceId,omitempty" name:"ShardResourceId"`
	// 分片对应的实例ID

	ShardInstanceId *int64 `json:"ShardInstanceId,omitempty" name:"ShardInstanceId"`
	// 分片状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 分片范围

	Range *string `json:"Range,omitempty" name:"Range"`
	// cpu

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 内存大小

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 磁盘大小

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 主可用区

	MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`
	// 备可用区

	SlaveZones []*string `json:"SlaveZones,omitempty" name:"SlaveZones"`
	// db信息

	DataBases []*Database `json:"DataBases,omitempty" name:"DataBases"`
	// proxy信息

	Proxies []*Proxy `json:"Proxies,omitempty" name:"Proxies"`
}

type RecycleNetInfo struct {
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// port

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// VpcId

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// SubnetId

	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 资源状态，0是待回收，1是回收中

	Status *string `json:"Status,omitempty" name:"Status"`
	// 回收时间

	RecycleTime *string `json:"RecycleTime,omitempty" name:"RecycleTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 网络类型 ipv4/ipv6

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 允许的操作,包含`ReleaseResource`

	AllowedActions []*string `json:"AllowedActions,omitempty" name:"AllowedActions"`
}

type DescribeDBEnginesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的引擎总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 售卖规格详细信息列表。

		DBEngineSet []*DBEngine `json:"DBEngineSet,omitempty" name:"DBEngineSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBEnginesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBEnginesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTenantDBInstanceDetailRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如： `tdsql-xxxxxxxx` 或 `tdsqlshard-xxxxxxxx`

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeTenantDBInstanceDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTenantDBInstanceDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WorkFlowBrief struct {
	// 流程ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 流程名称

	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`
}

type DescribeInstanceRecycleNetInfosRequest struct {
	*tchttp.BaseRequest

	// 过滤条件 支持的key：instance-id、vip、 vpc-id每次请求的Filters的上限为10，Filter.Values的上限为5。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// ip，模糊匹配

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 偏移，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序依据，以什么排序，比如InstanceId，RecycleTime,CreateTime。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序规则，升序还是降序，DESC降序，ASC升序

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeInstanceRecycleNetInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceRecycleNetInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExclusiveGroup struct {
	// 独享资源池Id。

	FenceId *string `json:"FenceId,omitempty" name:"FenceId"`
	// 资源池内的设备数量

	DeviceCount *int64 `json:"DeviceCount,omitempty" name:"DeviceCount"`
	// 地域ID

	Region *string `json:"Region,omitempty" name:"Region"`
	// 兼容老接口，后期可能废弃

	BondAppIds []*int64 `json:"BondAppIds,omitempty" name:"BondAppIds"`
	// 资源池名称

	FenceName *string `json:"FenceName,omitempty" name:"FenceName"`
	// 资源池是否所有人可见。0-不可见；1-可见

	Visibility *int64 `json:"Visibility,omitempty" name:"Visibility"`
	// 该资源池允许的接口操作列表

	AllowedActions []*string `json:"AllowedActions,omitempty" name:"AllowedActions"`
}

type Instance struct {
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 枚举值，**待补充**

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 状态描述

	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
	// 所属租户appId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 数字的 VpcId

	NumericVpcId *int64 `json:"NumericVpcId,omitempty" name:"NumericVpcId"`
	// 数字的 SubnetId

	NumericSubnetId *int64 `json:"NumericSubnetId,omitempty" name:"NumericSubnetId"`
	// 实例访问 vip 地址

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 实例访问 vport

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 后台的 setId 或 groupId

	SerialId *string `json:"SerialId,omitempty" name:"SerialId"`
	// 付费模式，取值范围：`prepaid`：表示预付费，即包年包月（TCE暂不支持）；`postpaid`：表示后付费，即按量计费

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 引擎类型，mysql/mariadb

	DBEngine *string `json:"DBEngine,omitempty" name:"DBEngine"`
	// 引擎版本

	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`
	// 实例cpu核数（分布式实例为除创建中和待删除分片外的其他分片总和）

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 实例内存，单位GB（分布式实例为除创建中和待删除分片外的其他分片总和）

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 实例数据盘，单位GB（分布式实例为除创建中和待删除分片外的其他分片总和）

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
	// 是否分布式实例，1为分布式实例，0为集中式实例

	Shard *int64 `json:"Shard,omitempty" name:"Shard"`
	// 实例分片数，仅分布式实例有效

	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`
	// 实例主从节点数

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 实例 Cpu 架构，取值范围： ARM/X86

	CpuArch *string `json:"CpuArch,omitempty" name:"CpuArch"`
	// 实例创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 实例所在地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 是否临时实例，取值范围：0 : 常规实例，1: 临时实例(回档中)，2: 临时实例(回档成功)，3:临时实例(回档失败)，4: 延时删除的扩容旧实例，5: 等待删除的切换后的临时实例

	IsTmp *int64 `json:"IsTmp,omitempty" name:"IsTmp"`
	// 独享资源池id

	ExClusterId *string `json:"ExClusterId,omitempty" name:"ExClusterId"`
	// 当前流程信息，可能返回 null，表示无相关流程

	CurrentFlow *WorkFlowBrief `json:"CurrentFlow,omitempty" name:"CurrentFlow"`
	// 可以执行的操作。取值范围 IsolateInstance, ActivateInstance, DestroyInstance和MigrateInstance

	AllowedActions []*string `json:"AllowedActions,omitempty" name:"AllowedActions"`
	// 临时实例释放时间

	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`
	// 临时实例类型 isolateInstance隔离实例 expandInstance扩容实例 rollbackInstance归档实例

	TempInstanceType *string `json:"TempInstanceType,omitempty" name:"TempInstanceType"`
	// 分片简要信息，不全部展示

	Shards []*ShardInfoBrief `json:"Shards,omitempty" name:"Shards"`
	// 实例ID，uuid

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 隔离类型

	IsolateType *int64 `json:"IsolateType,omitempty" name:"IsolateType"`
	// 隔离时间

	Isolatetime *string `json:"Isolatetime,omitempty" name:"Isolatetime"`
	// ipv6是否开启标志，1为已开启ipv6，0为未开启

	Ipv6Flag *int64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
	// ipv6地址

	Vip6 *string `json:"Vip6,omitempty" name:"Vip6"`
	// 资源池名称

	ExClusterName *string `json:"ExClusterName,omitempty" name:"ExClusterName"`
	// 租户账号名称

	TenantAccount *string `json:"TenantAccount,omitempty" name:"TenantAccount"`
	// 租户账号昵称

	TenantNickname *string `json:"TenantNickname,omitempty" name:"TenantNickname"`
}

type DeleteSpecConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSpecConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSpecConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableResourceRangeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAvailableResourceRangeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableResourceRangeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryFlowRequest struct {
	*tchttp.BaseRequest

	// 流程Id

	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *RetryFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindExclusiveGroupRequest struct {
	*tchttp.BaseRequest

	// 待解绑的独享资源池Id。

	FenceId *string `json:"FenceId,omitempty" name:"FenceId"`
	// 与哪个租户解绑。

	TenantAppId *int64 `json:"TenantAppId,omitempty" name:"TenantAppId"`
}

func (r *UnbindExclusiveGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindExclusiveGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSaleZoneConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该条数据的id

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSaleZoneConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSaleZoneConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpecConfigsRequest struct {
	*tchttp.BaseRequest

	// 根据 ID 查询售卖规格

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
	// 返回数量，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询偏移，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSpecConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpecConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceRecycleNetRecycleTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceRecycleNetRecycleTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceRecycleNetRecycleTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpecConfig struct {
	// 售卖规格Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// Cpu分配值，单位为1/100核

	RequestCpu *int64 `json:"RequestCpu,omitempty" name:"RequestCpu"`
	// Cpu展示值，单位为核（注意若开启 cpu 限制开关，将会把实例可用的cpu限制到这个规格）

	LimitCpu *int64 `json:"LimitCpu,omitempty" name:"LimitCpu"`
	// 内存大小，单位GB

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 可选数据盘下限，单位GB

	MinStorage *int64 `json:"MinStorage,omitempty" name:"MinStorage"`
	// 可选数据盘上限，单位GB

	MaxStorage *int64 `json:"MaxStorage,omitempty" name:"MaxStorage"`
	// 集中式实例售卖开关

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 地域ID

	Region *string `json:"Region,omitempty" name:"Region"`
	// 分布式实例售卖开关

	ShardEnable *bool `json:"ShardEnable,omitempty" name:"ShardEnable"`
}

type DeleteSaleZoneConfigRequest struct {
	*tchttp.BaseRequest

	// 待删除的规则ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteSaleZoneConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSaleZoneConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 需要过滤的字段。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段的过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DescribeSaleZoneConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的售卖规则总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 可用区售卖规则详细信息列表。

		SaleZoneConfigSet []*SaleZoneConfig `json:"SaleZoneConfigSet,omitempty" name:"SaleZoneConfigSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSaleZoneConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSaleZoneConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBEngineStatusRequest struct {
	*tchttp.BaseRequest

	// 待修改的ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 集中式实例售卖开关

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 分布式式实例售卖开关

	ShardEnable *bool `json:"ShardEnable,omitempty" name:"ShardEnable"`
}

func (r *ModifyDBEngineStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBEngineStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyExclusiveGroupNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyExclusiveGroupNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyExclusiveGroupNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAvailableResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源信息

		ResourceSet []*Resource `json:"ResourceSet,omitempty" name:"ResourceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceAvailableResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceAvailableResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceReleaseTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceReleaseTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceReleaseTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyExclusiveGroupVisibilityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyExclusiveGroupVisibilityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyExclusiveGroupVisibilityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如： `tdsql-xxxxxxxx` 或 `tdsqlshard-xxxxxxxx`

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *IsolateInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolateInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySaleZoneConfigStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySaleZoneConfigStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySaleZoneConfigStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Resource struct {
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 机器资源

	Machines []*Machine `json:"Machines,omitempty" name:"Machines"`
}

type DescribeWorkFlowTaskInfoRequest struct {
	*tchttp.BaseRequest

	// 流程Id

	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
	// 偏移，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeWorkFlowTaskInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWorkFlowTaskInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WorkFlowTaskInfo struct {
	// Id。

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 流程Id。

	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
	// 流程名称。

	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`
	// 任务名称。

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务首次开始执行时间。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 任务最后一次执行结束时间。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 任务最后一次执行状态。

	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
	// 任务执行次数。

	TaskRetryCount *int64 `json:"TaskRetryCount,omitempty" name:"TaskRetryCount"`
	// 任务描述信息。

	TaskDescription *string `json:"TaskDescription,omitempty" name:"TaskDescription"`
	// 任务ID。

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type IsolateInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流程ID

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBEngineStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBEngineStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBEngineStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Proxy struct {
	// ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 是否为主机

	IsMaster *int64 `json:"IsMaster,omitempty" name:"IsMaster"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
}
