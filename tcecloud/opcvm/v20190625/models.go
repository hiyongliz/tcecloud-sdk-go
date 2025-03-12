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

package v20190625

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type RecordList struct {
	// 原镜像ID

	SrcImageId *string `json:"SrcImageId,omitempty" name:"SrcImageId"`
	// 转化后的镜像ID

	DstImageId *string `json:"DstImageId,omitempty" name:"DstImageId"`
	// 操作系统名称

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 操作记录

	OperationRecords []*OperationRecords `json:"OperationRecords,omitempty" name:"OperationRecords"`
	// 具体操作

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 状态

	ImageState *string `json:"ImageState,omitempty" name:"ImageState"`
	// 目的地域

	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`
	// 操作时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 错误码

	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 操作记录

	Para *string `json:"Para,omitempty" name:"Para"`
}

type CreateImageRequest struct {
	*tchttp.BaseRequest

	// 要制作自定义镜像的uuid

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 自定义镜像的名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 自定义镜像的备注说明

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Agent的配置列表

		AgentConfigs []*AgentConfig `json:"AgentConfigs,omitempty" name:"AgentConfigs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FaultEventInfo struct {
	// 故障时间id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 触发时间

	TriggerTime *uint64 `json:"TriggerTime,omitempty" name:"TriggerTime"`
	// 固资号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 宿主机ip

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
	// 机型

	SvrDeviceType *string `json:"SvrDeviceType,omitempty" name:"SvrDeviceType"`
	// 设备类型，0. underlay; 1. overlay

	HostType *int64 `json:"HostType,omitempty" name:"HostType"`
	// 可用区

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// 故障原因

	FaultDetail *string `json:"FaultDetail,omitempty" name:"FaultDetail"`
	// 宿主机所在逻辑区

	SvrLogicArea *string `json:"SvrLogicArea,omitempty" name:"SvrLogicArea"`
	// zoneid

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 交换机信息

	InnerSwitch *string `json:"InnerSwitch,omitempty" name:"InnerSwitch"`
	// 0 故障正常发起,迁移中 1 同交换机5分钟内有>=3台母机上报异常，第三台不触发实际迁移，仅记录故障信息 2 母机上子机全部故障迁移完成 # 3 母机故障迁移完成但存在子机失败 4 同一时间针对母机发起了多次故障迁移，故障冲突

	Detail *int64 `json:"Detail,omitempty" name:"Detail"`
}

type TerminateInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateImageVisibleRangeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateImageVisibleRangeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateImageVisibleRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VirtualPrivateCloud struct {
	// 私有网络ID，形如`vpc-xxx`。有效的VpcId可通过登录[控制台](//console.{{conf.main_domain}}/vpc)查询；也可以调用接口 [DescribeVpcEx](/ocloud/api/NetWork/VPC/APIs/云网络运营端api%20v3（opvpc）/版本（2020-02-14）/云网络运营端api%20v3接口/DescribeVpcEx) ，从接口返回中的`unVpcId`字段获取。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络子网ID，形如`subnet-xxx`。有效的私有网络子网ID可通过登录[控制台](//console.{{conf.main_domain}}/vpc/subnet?rid=1)查询；也可以调用接口  [DescribeSubnetEx](/ocloud/api/NetWork/VPC/APIs/云网络运营端api%20v3（opvpc）/版本（2020-02-14）/云网络运营端api%20v3接口/DescribeSubnetEx) ，从接口返回中的`unSubnetId`字段获取。

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 是否用作公网网关。公网网关只有在实例拥有公网IP以及处于私有网络下时才能正常使用。取值范围：<br><li>TRUE：表示用作公网网关<br><li>FALSE：表示不用作公网网关<br><br>默认取值：FALSE。

	AsVpcGateway *bool `json:"AsVpcGateway,omitempty" name:"AsVpcGateway"`
	// 私有网络子网 IP 数组，在创建实例、修改实例vpc属性操作中可使用此参数。当前仅批量创建多台实例时支持传入相同子网的多个 IP。

	PrivateIpAddresses *string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 为弹性网卡指定随机生成的 IPv6 地址数量。

	IpvAddressCount *int64 `json:"IpvAddressCount,omitempty" name:"IpvAddressCount"`
}

type DescribeZoneImportInstanceTypeConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可导入的子机详细信息

		InstanceTypeConfigImportSet []*InstanceTypeConfigImport `json:"InstanceTypeConfigImportSet,omitempty" name:"InstanceTypeConfigImportSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneImportInstanceTypeConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneImportInstanceTypeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRecycleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例在回收站中保留的时间

		RecycleTime *string `json:"RecycleTime,omitempty" name:"RecycleTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRecycleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRecycleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryFaultMigrateRequest struct {
	*tchttp.BaseRequest

	// 重试虚拟机的uuid

	UuidList []*string `json:"UuidList,omitempty" name:"UuidList"`
	// 选择迁移的宿主机ip

	HostIpList []*string `json:"HostIpList,omitempty" name:"HostIpList"`
}

func (r *RetryFaultMigrateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryFaultMigrateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceClassConfigsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstanceClassConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceClassConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePrivateImageStatusRequest struct {
	*tchttp.BaseRequest

	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 状态。上线(ONLINE)，下线(OFFLINE)

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *UpdatePrivateImageStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePrivateImageStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAllHostTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 上报成功列表

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 失败列表

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 执行总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAllHostTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAllHostTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncHostTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总共同步的宿主机类型个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 同步成功的宿主机类型列表

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 同步失败的宿主机类型列表

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncHostTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncHostTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ColdMigrateTenantInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务数组

		TaskSet []*TenantTask `json:"TaskSet,omitempty" name:"TaskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ColdMigrateTenantInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ColdMigrateTenantInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecycleTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// cvm配置

		Setting []*string `json:"Setting,omitempty" name:"Setting"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRecycleTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRecycleTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindSoldPoolOwnersRequest struct {
	*tchttp.BaseRequest

	// 待绑定的售卖池名称

	SoldPoolName *string `json:"SoldPoolName,omitempty" name:"SoldPoolName"`
	// 待绑定的账号列表

	OwnerList []*string `json:"OwnerList,omitempty" name:"OwnerList"`
}

func (r *BindSoldPoolOwnersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindSoldPoolOwnersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFaultEventRequest struct {
	*tchttp.BaseRequest

	// svr-lan-ip：宿主机ip； svr-asset-id：固资号

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回最大数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序key,default="triggerTime"

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 降序(DESC)和升序(ASC),default="DESC"

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeFaultEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFaultEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecycleTimeRequest struct {
	*tchttp.BaseRequest

	// 回收时间，以小时为单位。

	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *ModifyRecycleTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRecycleTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskReportItems struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 类型(流程总体、消息队列)

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 子模块任务总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 子模块任务等待中总数

	Ready *int64 `json:"Ready,omitempty" name:"Ready"`
	// 子模块平均等待时间

	Wait *float64 `json:"Wait,omitempty" name:"Wait"`
	// 子模块失败总数

	Error *int64 `json:"Error,omitempty" name:"Error"`
	// 是否有效 (1 有效)

	Current *int64 `json:"Current,omitempty" name:"Current"`
	// 检测日期

	Date *string `json:"Date,omitempty" name:"Date"`
	// 检测星期

	Week *string `json:"Week,omitempty" name:"Week"`
	// 检测时间

	Time *string `json:"Time,omitempty" name:"Time"`
	// 消费者总数

	Comsumer *int64 `json:"Comsumer,omitempty" name:"Comsumer"`
	// 活跃消费者

	Ativity *int64 `json:"Ativity,omitempty" name:"Ativity"`
	// 子模块任务处理中总数

	Unack *int64 `json:"Unack,omitempty" name:"Unack"`
}

type DescribeModifyRecycleTimeLogsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeModifyRecycleTimeLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModifyRecycleTimeLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostsResourceRequest struct {
	*tchttp.BaseRequest

	// ip-list: 要修改的宿主机的ip列表

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 待投放的超卖池

	SoldPoolName *string `json:"SoldPoolName,omitempty" name:"SoldPoolName"`
}

func (r *ModifyHostsResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostsResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindSoldPoolOwnersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总共解绑的账号数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 成功解绑的账号列表

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 未成功解绑的账号列表

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindSoldPoolOwnersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindSoldPoolOwnersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePreSchedulerRequest struct {
	*tchttp.BaseRequest

	// 地域Id

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 过滤条件，详见下表：实例过滤条件表。每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`InstanceIds`和`Filters`。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 用户的APPID

	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

func (r *DescribePreSchedulerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePreSchedulerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostsResourceExclusiveHostFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改的总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 修改成功的列表

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostsResourceExclusiveHostFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostsResourceExclusiveHostFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostResource struct {
	// 宿主机总cpu核数

	CpuTotal *uint64 `json:"CpuTotal,omitempty" name:"CpuTotal"`
	// 宿主机可用cpu核数

	CpuAvailable *uint64 `json:"CpuAvailable,omitempty" name:"CpuAvailable"`
	// 宿主机总内存大小

	MemTotal *float64 `json:"MemTotal,omitempty" name:"MemTotal"`
	// 宿主机可用内存大小

	MemAvailable *float64 `json:"MemAvailable,omitempty" name:"MemAvailable"`
	// 宿主机总硬盘大小

	DiskTotal *uint64 `json:"DiskTotal,omitempty" name:"DiskTotal"`
	// 宿主机可用硬盘大小

	DiskAvailable *uint64 `json:"DiskAvailable,omitempty" name:"DiskAvailable"`
	// 宿主机可用NUMA大小

	NodeQuotaAvailable []*int64 `json:"NodeQuotaAvailable,omitempty" name:"NodeQuotaAvailable"`
	// 宿主机总Numa大小

	NodeQuotaTotal []*uint64 `json:"NodeQuotaTotal,omitempty" name:"NodeQuotaTotal"`
	// 宿主机超分前总cpu核数

	CpuOrigin *uint64 `json:"CpuOrigin,omitempty" name:"CpuOrigin"`
	// 宿主机可用内存大小

	MemoryAvailable *float64 `json:"MemoryAvailable,omitempty" name:"MemoryAvailable"`
	// 宿主机总内存大小

	MemoryTotal *float64 `json:"MemoryTotal,omitempty" name:"MemoryTotal"`
}

type CreateZoneInstanceTypeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 尝试创建的配置个数

		TotalCountT *uint64 `json:"TotalCountT,omitempty" name:"TotalCountT"`
		// 创建失败的配置ID列表

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 创建成功的配置ID列表

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateZoneInstanceTypeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateZoneInstanceTypeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceFamilyRequest struct {
	*tchttp.BaseRequest

	// 实例族唯一Id标识

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 修改实例族对应的白名单

	WhiteListKey *string `json:"WhiteListKey,omitempty" name:"WhiteListKey"`
}

func (r *ModifyInstanceFamilyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceFamilyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDeviceStatusRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否强制刷新

	Force *bool `json:"Force,omitempty" name:"Force"`
}

func (r *ModifyDeviceStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDeviceStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceClassQuota struct {
	// 机型大类名称，比如“S”、“N”

	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`
	// 机型大类显示名称

	InstanceClassName *string `json:"InstanceClassName,omitempty" name:"InstanceClassName"`
	// 机型大类的优先级

	Order *int64 `json:"Order,omitempty" name:"Order"`
	// 实例机型族详细信息

	InstanceFamilyQuotaSet []*InstanceFamilyQuota `json:"InstanceFamilyQuotaSet,omitempty" name:"InstanceFamilyQuotaSet"`
}

type InstanceFamilyQuota struct {
	// 实例族名

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 优先级

	Order *uint64 `json:"Order,omitempty" name:"Order"`
	// 实例族显示名称

	InstanceFamilyName *string `json:"InstanceFamilyName,omitempty" name:"InstanceFamilyName"`
	// 每个机型的详情

	InstanceTypeQuotaSet []*InstanceTypeQuota `json:"InstanceTypeQuotaSet,omitempty" name:"InstanceTypeQuotaSet"`
}

type DescribeVsChildTasksRequest struct {
	*tchttp.BaseRequest

	// 查询日志范围的起始时间；例：2019-07-05 00:00:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询日志范围的截止时间；例：2019-07-05 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件，详见下表：实例过滤条件表。参数不支持同时指定`InstanceIds`和`Filters`。  3100opcvm相对于380opcvm，/InputInfo/4/desc发生变化，3100opcvm为: 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/api/213/11646#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)中的相关小节。 ,380opcvm为: 返回数量，默认为20。关于`Limit`的更进一步介绍请参考 API [简介](/document/api/213/11646#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)中的相关小节。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](/document/api/213/11646#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/api/213/11646#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 是否包含MQ详情，默认为False。

	IncludeDetails *bool `json:"IncludeDetails,omitempty" name:"IncludeDetails"`
}

func (r *DescribeVsChildTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVsChildTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationRecords struct {
	// 操作，包括自定义镜像转公共镜像和公共镜像跨地域复制等操作

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 状态：已转为公共镜像，转换失败，转化中，已复制到其它地域，复制中，复制失败等

	Status *string `json:"Status,omitempty" name:"Status"`
	// 操作时间

	OperateTime *string `json:"OperateTime,omitempty" name:"OperateTime"`
}

type TenantInstanceSet struct {
	// CAM角色名。

	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`
	// 创建时间。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 数据卷。

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 设备Id。

	DeviceId *int64 `json:"DeviceId,omitempty" name:"DeviceId"`
	// 分散置放群组ID。

	DisasterRecoverGroupId []*string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`
	// 错误码。

	ErrorCode *uint64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 错误key。

	ErrorKey *string `json:"ErrorKey,omitempty" name:"ErrorKey"`
	// 到期时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。

	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// GPU核数。

	GPU *int64 `json:"GPU,omitempty" name:"GPU"`
	// 虚拟机监控程序。

	Hypervisor *int64 `json:"Hypervisor,omitempty" name:"Hypervisor"`
	// 实例的ipv6地址。

	IPv6Addresses *string `json:"IPv6Addresses,omitempty" name:"IPv6Addresses"`
	// 生产实例所使用的镜像`ID`。

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像类型，例如`PUBLIC_IMAGE`。

	ImageType []*string `json:"ImageType,omitempty" name:"ImageType"`
	// Vpc id。

	InnerVpcId *uint64 `json:"InnerVpcId,omitempty" name:"InnerVpcId"`
	// 实例计费模式。取值范围：<br><li>`PREPAID`：表示预付费，即包年包月<br><li>`POSTPAID_BY_HOUR`：表示后付费，即按量计费<br><li>`CDHPAID`：`CDH`付费，即只对`CDH`计费，不对`CDH`上的实例计费。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 实例类，例如'S'。

	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`
	// 实例类型，例如 `S5l`。

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 实例`ID`。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例状态。取值范围：<br><li>PENDING：表示创建中<br></li><li>LAUNCH_FAILED：表示创建失败<br></li><li>RUNNING：表示运行中<br></li><li>STOPPED：表示关机<br></li><li>STARTING：表示开机中<br></li><li>STOPPING：表示关机中<br></li><li>REBOOTING：表示重启中<br></li><li>SHUTDOWN：表示停止待销毁<br></li><li>TERMINATING：表示销毁中。<br></li>

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 实例机型。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 是否安全隔离

	IsSafeIsolated *bool `json:"IsSafeIsolated,omitempty" name:"IsSafeIsolated"`
	// 隔离来源

	IsolatedSource *string `json:"IsolatedSource,omitempty" name:"IsolatedSource"`
	// 隔离时间

	IsolatedTime *string `json:"IsolatedTime,omitempty" name:"IsolatedTime"`
	// 密钥对id

	KeyPairIds *uint64 `json:"KeyPairIds,omitempty" name:"KeyPairIds"`
	// 实例的最新操作。例：StopInstances、ResetInstance。

	LatestOperation *string `json:"LatestOperation,omitempty" name:"LatestOperation"`
	// 实例的最新操作状态。取值范围：<br><li>SUCCESS：表示操作成功<br><li>OPERATING：表示操作执行中<br><li>FAILED：表示操作失败

	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`
	// 实例最新操作的唯一请求 ID。

	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitempty" name:"LatestOperationRequestId"`
	// 实例登录设置。目前只返回实例所关联的密钥。

	LoginSettings []*LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 实例内存容量，单位：`GB`。

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 是否为新创建。

	NewCreationIdentify *bool `json:"NewCreationIdentify,omitempty" name:"NewCreationIdentify"`
	// 身份验证。

	OperationMask *uint64 `json:"OperationMask,omitempty" name:"OperationMask"`
	// 实例的os名称

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 实例所在的位置。

	Placement *TenantInstancesPlacement `json:"Placement,omitempty" name:"Placement"`
	// 资源所属项目Id。

	PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
	// 实例主网卡的内网`IP`列表。

	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 实例主网卡的公网`IP`列表。

	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`
	// 自动续费标识。取值范围：<br><li>`NOTIFY_AND_MANUAL_RENEW`：表示通知即将过期，但不自动续费<br><li>`NOTIFY_AND_AUTO_RENEW`：表示通知即将过期，而且自动续费<br><li>`DISABLE_NOTIFY_AND_MANUAL_RENEW`：表示不通知即将过期，也不自动续费。

	RenewFlag []*string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 实例业务状态。取值范围：<br><li>NORMAL：表示正常状态的实例<br><li>EXPIRED：表示过期的实例<br><li>PROTECTIVELY_ISOLATED：表示被安全隔离的实例。

	RestrictState *string `json:"RestrictState,omitempty" name:"RestrictState"`
	// 设备运行状态。

	RunFlag *uint64 `json:"RunFlag,omitempty" name:"RunFlag"`
	// 安全隔离信息。

	SafeIsolatedInfo *string `json:"SafeIsolatedInfo,omitempty" name:"SafeIsolatedInfo"`
	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](/ocloud/api/NetWork/VPC/APIs/私有网络（vpc）/版本（2017-03-12）/安全组相关接口/DescribeSecurityGroups) 的返回值中的sgId字段来获取。

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 实例的关机计费模式。

	StopChargingMode *string `json:"StopChargingMode,omitempty" name:"StopChargingMode"`
	// 实例系统盘信息。

	SystemDisk []*SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 实例的uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 实例所属虚拟私有网络信息。

	VirtualPrivateCloud []*VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
	// 实例关联的标签列表。

	Tag *Tag `json:"Tag,omitempty" name:"Tag"`
}

type DescribeAllinoneTasksRequest struct {
	*tchttp.BaseRequest

	// 查询日志范围的起始时间；例：2019-07-05 00:00:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询日志范围的截止时间；例：2019-07-05 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件，详见下表：实例过滤条件表。每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`InstanceIds`和`Filters`。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](/document/api/213/11646#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/api/213/11646#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAllinoneTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllinoneTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目，专用宿主机（对于独享母机付费模式的子机创建）等属性。

	Placement *Placement `json:"Placement,omitempty" name:"Placement"`
	// 有效镜像Id

	ImageUuid *string `json:"ImageUuid,omitempty" name:"ImageUuid"`
	// 系统盘配置信息

	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 数据盘配置信息

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 购买数目

	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
	// 实例显示名字

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例登录设置

	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// CPU核数

	CPU *int64 `json:"CPU,omitempty" name:"CPU"`
	// 内存大小，单位为MB

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 母鸡类型，如“M10”

	HostType *string `json:"HostType,omitempty" name:"HostType"`
	// NFV机型跨numa参数，有该参数且不为NULL，可以跨numa，参数为NULL或没有该参数，则不允许跨numa

	NFVAcrossNodeFlag *bool `json:"NFVAcrossNodeFlag,omitempty" name:"NFVAcrossNodeFlag"`
	// 支持子机自定义IP，目前只针对NFV子机的场景，限制条件是一次只允许创建一台子机

	CustomIP *string `json:"CustomIP,omitempty" name:"CustomIP"`
}

func (r *RunInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceTypeDisasterGroupConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceTypeDisasterGroupConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceTypeDisasterGroupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesPasswordRequest struct {
	*tchttp.BaseRequest

	// 一个或多个实例

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 实例登录密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 待重置密码的实例操作系统用户名。不得超过64个字符。

	Username *string `json:"Username,omitempty" name:"Username"`
}

func (r *ResetInstancesPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetInstancesPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FaultRecordInfo struct {
	// 任务id

	Id *string `json:"Id,omitempty" name:"Id"`
	// uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// cpu核数，除100为真实值

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 内存，单位为M

	Mem *int64 `json:"Mem,omitempty" name:"Mem"`
	// 磁盘大小，单位为G

	VolumeSize *uint64 `json:"VolumeSize,omitempty" name:"VolumeSize"`
	// 磁盘类型

	VolumeType *string `json:"VolumeType,omitempty" name:"VolumeType"`
	// 子机ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 发生故障的宿主机ip

	OrigHost *string `json:"OrigHost,omitempty" name:"OrigHost"`
	// 宿主机类型

	HostType *int64 `json:"HostType,omitempty" name:"HostType"`
	// 迁移目标宿主机ip

	TargetHost *string `json:"TargetHost,omitempty" name:"TargetHost"`
	// 开始时间

	BeginTime *uint64 `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 耗时

	Elapse *uint64 `json:"Elapse,omitempty" name:"Elapse"`
	// 操作者

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 0待发起，1成功，2失败，3迁移中

	Detail *int64 `json:"Detail,omitempty" name:"Detail"`
	// appid

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
}

type CreateCvmTypeMapResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败列表

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 成功列表

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 执行总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCvmTypeMapResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCvmTypeMapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LiveMigrateInstancesRequest struct {
	*tchttp.BaseRequest

	// 要做热迁移的实例id列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 指定迁移的母鸡

	HostIps []*string `json:"HostIps,omitempty" name:"HostIps"`
	// 超时时间，单位为秒

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
	// 虚拟机热迁移pause阶段最大暂停时间，单位毫秒。
	// VStation默认100毫秒，建议30ms~30000ms

	PauseDuration *uint64 `json:"PauseDuration,omitempty" name:"PauseDuration"`
	// 指定带宽，单位为MiB/s

	MaxBandwidth *uint64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// 指定售卖池迁移

	SoldPool *string `json:"SoldPool,omitempty" name:"SoldPool"`
}

func (r *LiveMigrateInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LiveMigrateInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCvmTypeMapRequest struct {
	*tchttp.BaseRequest

	// 机型type

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// CPU架构

	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
}

func (r *ModifyCvmTypeMapRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCvmTypeMapRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志流水信息。

		FlowLogSet []*FlowLogSet `json:"FlowLogSet,omitempty" name:"FlowLogSet"`
		// 日志总条目数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LocalDiskType struct {
	// 本地磁盘类型：“LOCAL_BASIC”、“LOCAL_SSD”

	Type *string `json:"Type,omitempty" name:"Type"`
	// 本地磁盘属性。“ROOT”、“DATA”

	PartitionType *string `json:"PartitionType,omitempty" name:"PartitionType"`
	// 本地磁盘最小值

	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`
	// 本地磁盘最大值

	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

type DescribeInstanceTypeConfigsRequest struct {
	*tchttp.BaseRequest

	// 获取的可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 实例族名

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
}

func (r *DescribeInstanceTypeConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceTypeConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailoverMigrateInstancesRequest struct {
	*tchttp.BaseRequest

	// 要进行迁移的虚拟机

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 迁移到指定的母机ip列表

	HostIps []*string `json:"HostIps,omitempty" name:"HostIps"`
	// 如果虚拟机是本地盘，这个字段必传。1表示数据允许丢失

	Lossy *uint64 `json:"Lossy,omitempty" name:"Lossy"`
	// 如果传了lossy字段并且为1，那么该字段必传。
	// 用该镜像ID来创建虚拟机

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *FailoverMigrateInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FailoverMigrateInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCpuModelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 支持的CPU架构列表

		CPUModelSet []*CPUModelInfoArchitecture `json:"CPUModelSet,omitempty" name:"CPUModelSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCpuModelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCpuModelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LiveMigrateTenantInstancesRequest struct {
	*tchttp.BaseRequest

	// 待迁移实例uuid列表:["5182c761-90d7-4cd9-bbf8-4c19c4896dc8"]

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 指定母机迁移

	HostIps []*string `json:"HostIps,omitempty" name:"HostIps"`
	// 要迁入的售卖池， PLAIN, OVERSOLD等

	SoldPool *string `json:"SoldPool,omitempty" name:"SoldPool"`
	// 跨可用区迁移，要传入可用区表

	Zones []*string `json:"Zones,omitempty" name:"Zones"`
	// 最大迁移带宽

	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// 最大迁移超时

	MaxTimeout *int64 `json:"MaxTimeout,omitempty" name:"MaxTimeout"`
}

func (r *LiveMigrateTenantInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LiveMigrateTenantInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Status struct {
	// 1表示可用 3表示不可用
}

type DescribeLogStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Vstation日志不同错误码的统计

		VstationLogStatistics *VstationLog `json:"VstationLogStatistics,omitempty" name:"VstationLogStatistics"`
		// Allinone日志不同错误码的统计

		AllinoneLogStatistics *AllinoneLog `json:"AllinoneLogStatistics,omitempty" name:"AllinoneLogStatistics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceType struct {
	// 实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// CPU个数

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 内存大小

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// GPU核数

	Gpu *uint64 `json:"Gpu,omitempty" name:"Gpu"`
	// FPGA核数

	Fpga *uint64 `json:"Fpga,omitempty" name:"Fpga"`
	// 存储块数

	StorageBlock *uint64 `json:"StorageBlock,omitempty" name:"StorageBlock"`
	// 网卡数

	NetworkCard *uint64 `json:"NetworkCard,omitempty" name:"NetworkCard"`
	// 最大带宽

	MaxBandWidth *float64 `json:"MaxBandWidth,omitempty" name:"MaxBandWidth"`
	// 主频

	Frequency *string `json:"Frequency,omitempty" name:"Frequency"`
	// CPU型号名称

	CpuModelName *string `json:"CpuModelName,omitempty" name:"CpuModelName"`
	// 包转发率

	Pps *uint64 `json:"Pps,omitempty" name:"Pps"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DescribeAgentsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认999999

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAgentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFaultRecordRequest struct {
	*tchttp.BaseRequest

	// uuid:uuid; ip:ip id:故障事件id orig-host：源宿主机ip record-status：0待发起，1成功，2失败，3迁移中

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回最大数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序key(beginTime、endTime),default="beginTime",

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 降序(DESC)和升序(ASC),default="DESC"

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeFaultRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFaultRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceTypeConfigImport struct {
	// 机型类型名称，如“S3.SMALL!”

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// CPU个数

	CPU *int64 `json:"CPU,omitempty" name:"CPU"`
	// 内存大小

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 售卖状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 评论

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 最大网络带宽G/s

	MaxBandwidth *float64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// 本地盘类型

	DiskType *int64 `json:"DiskType,omitempty" name:"DiskType"`
}

type Filter struct {
	// 需要过滤的字段

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段的过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DescribeImageRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的记录数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 符合条件的记录列表

		ImageRecordSet []*RecordList `json:"ImageRecordSet,omitempty" name:"ImageRecordSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostsTagRequest struct {
	*tchttp.BaseRequest

	// 要修改的母机ip列表

	HostIps []*string `json:"HostIps,omitempty" name:"HostIps"`
	// 要增加的Tag:
	// 如HOST_FAULT

	Tag *string `json:"Tag,omitempty" name:"Tag"`
	// 操作原因。当删除标签的时候，需要使用这个信息。

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 本次的操作人。用于追溯信息使用。

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 标签的描述。详细的描述一下标签的说明。

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyHostsTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostsTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindSoldPoolOwnersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定成功的账号个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 绑定成功的账号列表

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 绑定失败的账号列表

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindSoldPoolOwnersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindSoldPoolOwnersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务详细信息

		TaskSet []*Task `json:"TaskSet,omitempty" name:"TaskSet"`
		// 任务数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentLog struct {
	// 日志ID

	LogId *uint64 `json:"LogId,omitempty" name:"LogId"`
	// 日志正文

	LogContent *string `json:"LogContent,omitempty" name:"LogContent"`
	// 从那一时刻开始

	LogUpdateTime *string `json:"LogUpdateTime,omitempty" name:"LogUpdateTime"`
}

type VsTaskInfo struct {
	// AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 错误码

	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 错误信息

	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
	// 物理机地址

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 实例ip

	InstanceIp *string `json:"InstanceIp,omitempty" name:"InstanceIp"`
	// 实例uuid

	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
	// 日志详情

	LogDetail *string `json:"LogDetail,omitempty" name:"LogDetail"`
	// 操作系统类型

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 父任务id

	ParentTaskId *int64 `json:"ParentTaskId,omitempty" name:"ParentTaskId"`
	// 任务状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 子任务id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DeleteImagesRequest struct {
	*tchttp.BaseRequest

	// 要操作的镜像数组

	ImageUuids []*string `json:"ImageUuids,omitempty" name:"ImageUuids"`
}

func (r *DeleteImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateImageVisibleRangeRequest struct {
	*tchttp.BaseRequest

	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 资源类型：T_CLOUD_RESOURCE表示增加镜像增加租户端的生效范围、O_CLOUD_RESOURCE表示镜像增加运营端的生效范围

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *UpdateImageVisibleRangeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateImageVisibleRangeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCvmTypeBillResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败列表

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 成功列表

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 执行总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCvmTypeBillResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCvmTypeBillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 是否强制退还

	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Placement struct {
	// 实例所属可用区的ID。该参数也可以通过调用的返回值中的Zone字段来获取。

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 所属项目Id

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 宿主机ip列表

	HostIps []*string `json:"HostIps,omitempty" name:"HostIps"`
}

type CreateHostTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总共导入的宿主机类型个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 成功的宿主机类型名列表

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 失败的宿主机类型名列表

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHostTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHostTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// 要查询的任务的id数组

	TaskIds []*int64 `json:"TaskIds,omitempty" name:"TaskIds"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVsTaskInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志总量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 日志详情

		TaskInfoSet []*TaskStep `json:"TaskInfoSet,omitempty" name:"TaskInfoSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVsTaskInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVsTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest

	// zone-id - Integer - 是否必填：是 -（过滤条件）按照可用区过滤。
	//
	// instance-family-list  是否必填：否 -（过滤条件）按照机型系列过滤。按照实例机型系列过滤。实例机型系列形如：S1、I1、M1等

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeZoneInstanceConfigInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneInstanceConfigInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostItem struct {
	// 宿主机内网IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 宿主机资源信息

	HostResource *HostResource `json:"HostResource,omitempty" name:"HostResource"`
	// 宿主机标记

	HostTag []*string `json:"HostTag,omitempty" name:"HostTag"`
	// 宿主机固资号

	HostAsset *string `json:"HostAsset,omitempty" name:"HostAsset"`
	// 母机状态：
	// FAULT、ONHOLD、NORMAL

	HostStatus *string `json:"HostStatus,omitempty" name:"HostStatus"`
	// VPC: VPC网络宿主机
	// CLASSIC:物理网络宿主机

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// 宿主机状态，装箱流程使用
	// MOADOWN: 无法获取moa信息
	// NONEXIST：无法获取该宿主机信息
	// TIMEOUT：获取的宿主机信息超时
	// NORMAL：状态正常
	// BADSTATE：宿主机状态不正常

	HeartBeat *string `json:"HeartBeat,omitempty" name:"HeartBeat"`
	// 宿主机所在可用区Id

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 宿主机类型，如：“M10”

	HostType *string `json:"HostType,omitempty" name:"HostType"`
	// 宿主机所在资源池

	Pool *string `json:"Pool,omitempty" name:"Pool"`
	// 宿主机是否已经上线

	IsOnline *bool `json:"IsOnline,omitempty" name:"IsOnline"`
	// 宿主机上联交换机

	NetworkDeviceId *string `json:"NetworkDeviceId,omitempty" name:"NetworkDeviceId"`
	// 宿主机当前所在的售卖池

	SoldPoolName *string `json:"SoldPoolName,omitempty" name:"SoldPoolName"`
	// 宿主机CPU架构

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	// 宿主机作为专用宿主机的购买用户

	ExclusiveHostOwner *string `json:"ExclusiveHostOwner,omitempty" name:"ExclusiveHostOwner"`
	// 宿主机虚拟化类型

	HostHypervisor *string `json:"HostHypervisor,omitempty" name:"HostHypervisor"`
	// 宿主机状态：ONLINE；

	HostState *string `json:"HostState,omitempty" name:"HostState"`
	// 宿主机是否作为专用宿主机

	IsExclusiveHost *bool `json:"IsExclusiveHost,omitempty" name:"IsExclusiveHost"`
	// 宿主机可用内存，单位G

	MemoryAvailable *int64 `json:"MemoryAvailable,omitempty" name:"MemoryAvailable"`
	// 宿主机总共内存，单位G

	MemoryTotal *int64 `json:"MemoryTotal,omitempty" name:"MemoryTotal"`
}

type ResetInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 虚拟机新的镜像ID

	ImageUuid *string `json:"ImageUuid,omitempty" name:"ImageUuid"`
	// 重装虚拟机的ROOT盘lv的大小。

	ResetSize *uint64 `json:"ResetSize,omitempty" name:"ResetSize"`
	// 重装系统后的登录信息

	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
}

func (r *ResetInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ColdMigrateTenantInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例的uuid数组，如["5182c761-90d7-4cd9-bbf8-4c19c4896dc8"]

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 目标母机数组

	HostIps []*string `json:"HostIps,omitempty" name:"HostIps"`
	// 要嵌入的售卖池，默认不变更

	SoldPool *string `json:"SoldPool,omitempty" name:"SoldPool"`
	// 需要跨可用迁移，可传入可用区列表

	Zones []*string `json:"Zones,omitempty" name:"Zones"`
	// 最大迁移带宽

	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// 最大迁移超时

	MaxTimeout *int64 `json:"MaxTimeout,omitempty" name:"MaxTimeout"`
	// 是否保留hostname

	HostNameReserved *bool `json:"HostNameReserved,omitempty" name:"HostNameReserved"`
	// 是否全部带宽

	AllBandwidth *bool `json:"AllBandwidth,omitempty" name:"AllBandwidth"`
}

func (r *ColdMigrateTenantInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ColdMigrateTenantInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateZoneInstanceTypeConfigRequest struct {
	*tchttp.BaseRequest

	// 要导入的可用区的zoneid

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 机型大类

	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`
	// 实例族

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 实例类型，可用区

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 机型核数

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 机型内存大小，单位为GB

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 机型的最大带宽

	MaxBandwidth *float64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// 本地盘类型

	LocalDiskTypeList []*LocalDiskType `json:"LocalDiskTypeList,omitempty" name:"LocalDiskTypeList"`
	// 售罄策略：
	// 0：资源决定
	// 1：强制售罄

	SoldOutStrategy *int64 `json:"SoldOutStrategy,omitempty" name:"SoldOutStrategy"`
	// 是否可售卖：
	// 1：开放售卖
	// 0：停止售卖

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *CreateZoneInstanceTypeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateZoneInstanceTypeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePublicImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePublicImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePublicImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSoldPoolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 成功的SoldPoolName

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 失败的SoldPoolName

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSoldPoolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSoldPoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLatestOperationStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLatestOperationStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLatestOperationStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceTypeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 尝试添加的总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 失败的config的id列表，可以用来排错

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 成功的config的id列表

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceTypeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceTypeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneTenantInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeZoneTenantInstanceConfigInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneTenantInstanceConfigInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LiveMigrateTenantInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务明细

		TaskSet []*TenantTask `json:"TaskSet,omitempty" name:"TaskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LiveMigrateTenantInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LiveMigrateTenantInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ColdMigrateInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ColdMigrateInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ColdMigrateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncInstanceFamilyConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 尝试同步的配置个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 同步失败的配置ID列表

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 同步成功的配置ID列表

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncInstanceFamilyConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncInstanceFamilyConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 升级或者降级到的CPU核数

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 升级或降级的内存大小

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 升级或者降级磁盘大小

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

func (r *UpgradeInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskInfo struct {
	// 任务操作的实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 任务状态: WAITING: 任务等待中。RUNNING：任务正在执行中。SUCCESS：任务已经完成。FAIL：任务失败。UNKNOWN：未知

	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
}

type ModifyHostsResourceExclusiveHostFlagRequest struct {
	*tchttp.BaseRequest

	// 支持host-ip

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 要设置为是/否 1是0否

	ExclusiveHostFlag *int64 `json:"ExclusiveHostFlag,omitempty" name:"ExclusiveHostFlag"`
}

func (r *ModifyHostsResourceExclusiveHostFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostsResourceExclusiveHostFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSoldPoolsOwnerRequest struct {
	*tchttp.BaseRequest

	// 待查询的售卖池列表

	SoldPoolNameList []*string `json:"SoldPoolNameList,omitempty" name:"SoldPoolNameList"`
}

func (r *DescribeSoldPoolsOwnerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSoldPoolsOwnerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Image struct {
	// 镜像id如：img-qy9cl0e3

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像操作系统

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 镜像大小

	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`
	// 镜像操作系统

	OsKey *string `json:"OsKey,omitempty" name:"OsKey"`
	// PUBLIC_IMAGE: 公共镜像 PRIVATE_IMAGE:私有镜像

	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`
	// 镜像创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 镜像当前状态,标识镜像是否下线，如果ResourceType为O_CLOUD_RESOURCE，标识运营端下线，如果ResourceType为T_CLOUD_RESOURCE,标识租户端镜像，如果ResourceType为T_O_CLOUD_RESOURCE，标识租户端或运营端下线

	ImageState *string `json:"ImageState,omitempty" name:"ImageState"`
	// EXTERNAL_IMPORT, CREATE_IMAGE, OFFICIAL

	ImageSource *string `json:"ImageSource,omitempty" name:"ImageSource"`
	// 镜像名称

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 操作系统平台

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 操作系统架构

	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
	// 镜像全局id如：img2018053006451178

	ImageUuid *string `json:"ImageUuid,omitempty" name:"ImageUuid"`
	// 镜像所有者

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 镜像描述

	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`
	// 镜像容量，单位G。

	ImgCapacitySize *uint64 `json:"ImgCapacitySize,omitempty" name:"ImgCapacitySize"`
	// 镜像实际大小，单位G。

	ImgFullSize *float64 `json:"ImgFullSize,omitempty" name:"ImgFullSize"`
	// 关联实例

	InstanceNumber *uint64 `json:"InstanceNumber,omitempty" name:"InstanceNumber"`
	// 镜像资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

type DescribeFaultRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 详细信息

		Detail []*FaultRecordInfo `json:"Detail,omitempty" name:"Detail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFaultRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFaultRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSchedulerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建虚拟机任务失败的原因结果列表

		SchedulerInfoSet []*SchedulerInfo `json:"SchedulerInfoSet,omitempty" name:"SchedulerInfoSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSchedulerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSchedulerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVsTaskInfoRequest struct {
	*tchttp.BaseRequest

	// Vs的子任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeVsTaskInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVsTaskInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySoldPoolRequest struct {
	*tchttp.BaseRequest

	// 要修改的售卖池名称

	SoldPoolName *string `json:"SoldPoolName,omitempty" name:"SoldPoolName"`
	// 售卖池的新描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 售卖池超分比

	SoldRate *float64 `json:"SoldRate,omitempty" name:"SoldRate"`
}

func (r *ModifySoldPoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySoldPoolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCDHExternalMachineTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCDHExternalMachineTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCDHExternalMachineTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetInstancesPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SoldPoolOwner struct {
	// 售卖池名称

	SoldPoolName *string `json:"SoldPoolName,omitempty" name:"SoldPoolName"`
	// 账户列表

	OwnerList []*string `json:"OwnerList,omitempty" name:"OwnerList"`
}

type OperatingInstanceTask struct {
	// 接口名

	Action *string `json:"Action,omitempty" name:"Action"`
	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 虚拟机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 是否可以刷新

	Prepare *bool `json:"Prepare,omitempty" name:"Prepare"`
}

type UpdatePrivateImageStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作结果

		Result *Result `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePrivateImageStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePrivateImageStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest

	// 接受三个数据参数：
	// instance-class-list,
	// instance-family-list,
	// instance-type-list

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeInstanceConfigInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceConfigInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误信息

		Error *Results `json:"Error,omitempty" name:"Error"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FailTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FailTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {
	// 数据盘类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 数据盘ID

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 数据盘大小

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 数据盘的序列号

	DiskSerial *string `json:"DiskSerial,omitempty" name:"DiskSerial"`
}

type CreateImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务Id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseAgentRequest struct {
	*tchttp.BaseRequest

	// Agent的名称

	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`
	// 要发布到的os类型：可选windows、linux、freebsd

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// Agent的URL，一般为COS链接

	AgentUrl *string `json:"AgentUrl,omitempty" name:"AgentUrl"`
	// Agent的URL，一般为COS链接

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// 网络

	Network *string `json:"Network,omitempty" name:"Network"`
}

func (r *ReleaseAgentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseAgentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVncResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// VNC ip地址

		VNCIP *string `json:"VNCIP,omitempty" name:"VNCIP"`
		// VNC 端口号

		VNCPort *uint64 `json:"VNCPort,omitempty" name:"VNCPort"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVncResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachUsbRequest struct {
	*tchttp.BaseRequest

	// 宿主机IP

	Host *string `json:"Host,omitempty" name:"Host"`
	// 虚拟机uuid

	VmUuid *string `json:"VmUuid,omitempty" name:"VmUuid"`
	// USB的VendorId

	VendorId *string `json:"VendorId,omitempty" name:"VendorId"`
	// USB的ProductId

	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
	// USB的DeviceNumber

	DeviceNumber *string `json:"DeviceNumber,omitempty" name:"DeviceNumber"`
	// USB的BusNumber

	BusNumber *string `json:"BusNumber,omitempty" name:"BusNumber"`
}

func (r *DetachUsbRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachUsbRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Results struct {
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 执行结果

	Result *string `json:"Result,omitempty" name:"Result"`
	// 反馈信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 错误码

	Code *string `json:"Code,omitempty" name:"Code"`
}

type CreatePublicImageRequest struct {
	*tchttp.BaseRequest

	// 操作系统后缀

	Postfix *string `json:"Postfix,omitempty" name:"Postfix"`
	// OS类别

	OsClass *string `json:"OsClass,omitempty" name:"OsClass"`
	// OS类型

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// OS架构

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	// 镜像描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 是否可见

	Visible *int64 `json:"Visible,omitempty" name:"Visible"`
	// 自定义镜像名称

	SourceImageId *string `json:"SourceImageId,omitempty" name:"SourceImageId"`
	// 是否reset base

	ResetBase *string `json:"ResetBase,omitempty" name:"ResetBase"`
	// 是否检查入参

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
	// 镜像默认登录账户

	DefaultUser *string `json:"DefaultUser,omitempty" name:"DefaultUser"`
	// 0：租户运营都可见，1:租户，2：运营

	ImageVisible *uint64 `json:"ImageVisible,omitempty" name:"ImageVisible"`
}

func (r *CreatePublicImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePublicImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TenantInstancesPlacement struct {
	// 项目id

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 专用宿主机id

	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

type DescribeInstanceUsbInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceUsbInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceUsbInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrantLiveMigrateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigrantLiveMigrateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrantLiveMigrateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserCpuTypeRequest struct {
	*tchttp.BaseRequest

	// 用户的AppId

	OwnerId *string `json:"OwnerId,omitempty" name:"OwnerId"`
	// 绑定类型。
	// Set:按Set绑定
	// Node: 按Numa节点绑定

	CpuBindType *string `json:"CpuBindType,omitempty" name:"CpuBindType"`
}

func (r *ModifyUserCpuTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserCpuTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceTypeConfigRequest struct {
	*tchttp.BaseRequest

	// 实例大类，比如“S”，“I”

	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`
	// 实例族“S3”，“S2”

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 实例名，如“S2.2XLARGE2”

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// CPU数量，单位个

	CPU *int64 `json:"CPU,omitempty" name:"CPU"`
	// 内存大小，单位GB

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 最大带宽大小，如1.5，单位MB/s

	MaxBandwidth *float64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// 是否售卖
	// 1 : 售卖
	// 0：暂停售卖

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// Gpu数量，GPU机型必填

	Gpu *int64 `json:"Gpu,omitempty" name:"Gpu"`
	// 直通盘数量，直通盘机型必填

	StorageBlock *int64 `json:"StorageBlock,omitempty" name:"StorageBlock"`
}

func (r *CreateInstanceTypeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceTypeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFaultEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 故障事件信息

		Detail []*FaultEventInfo `json:"Detail,omitempty" name:"Detail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFaultEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFaultEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MQQueueInfo struct {
	// 消费速率

	PublishRate *float64 `json:"PublishRate,omitempty" name:"PublishRate"`
	// 应答速率

	AckRate *float64 `json:"AckRate,omitempty" name:"AckRate"`
	// 传递速率

	DeliverGetRate *float64 `json:"DeliverGetRate,omitempty" name:"DeliverGetRate"`
	// 消息总数

	Messages *int64 `json:"Messages,omitempty" name:"Messages"`
	// 消息未应答数

	MessagesUnacknowledged *int64 `json:"MessagesUnacknowledged,omitempty" name:"MessagesUnacknowledged"`
	// 消息就绪数

	MessagesReady *int64 `json:"MessagesReady,omitempty" name:"MessagesReady"`
	// 虚拟主机

	Host *string `json:"Host,omitempty" name:"Host"`
	// 队列名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 节点

	Node *string `json:"Node,omitempty" name:"Node"`
	// 队列状态

	State *string `json:"State,omitempty" name:"State"`
	// 消费者利用率

	ConsumerUtilisation *int64 `json:"ConsumerUtilisation,omitempty" name:"ConsumerUtilisation"`
	// 消费者数量

	ConsumerNumber *int64 `json:"ConsumerNumber,omitempty" name:"ConsumerNumber"`
}

type DeleteImagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除任务Id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceUsbControllerRequest struct {
	*tchttp.BaseRequest

	// 宿主机IP

	Host *string `json:"Host,omitempty" name:"Host"`
	// 虚拟机uuid

	VmUuids []*string `json:"VmUuids,omitempty" name:"VmUuids"`
}

func (r *DescribeInstanceUsbControllerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceUsbControllerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCvmTypeBillRequest struct {
	*tchttp.BaseRequest

	// 新增机型自定义类型

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 新增机型自定义类型中文名

	InstanceFamilyChnName *string `json:"InstanceFamilyChnName,omitempty" name:"InstanceFamilyChnName"`
	// 新增机型自定义类型英文名

	InstanceFamilyEngName *string `json:"InstanceFamilyEngName,omitempty" name:"InstanceFamilyEngName"`
	// 内部代码

	DeviceClass *string `json:"DeviceClass,omitempty" name:"DeviceClass"`
}

func (r *CreateCvmTypeBillRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCvmTypeBillRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTenantInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例详细信息列表。

		InstanceSet []*TenantInstanceSet `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 数据总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTenantInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTenantInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneInstanceTypeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 尝试修改的个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 修改失败的配置的ID列表

		FailList *string `json:"FailList,omitempty" name:"FailList"`
		// 修改成功的配置的ID列表

		SuccessList *string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyZoneInstanceTypeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyZoneInstanceTypeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostType struct {
	// 类型名称

	HostTypeName *string `json:"HostTypeName,omitempty" name:"HostTypeName"`
	// 该宿主机可支持的虚拟子机类型

	InstanceFamilySet []*string `json:"InstanceFamilySet,omitempty" name:"InstanceFamilySet"`
	// 可售卖CPU核数，单位个

	CPUTotal *uint64 `json:"CPUTotal,omitempty" name:"CPUTotal"`
	// 可售卖内存大小，单位GB

	MemTotal *uint64 `json:"MemTotal,omitempty" name:"MemTotal"`
	// 可售卖磁盘大小，单位GB

	DiskTotal *uint64 `json:"DiskTotal,omitempty" name:"DiskTotal"`
	// 宿主机定制化参数，为json格式字符串

	HostExtraSpec *string `json:"HostExtraSpec,omitempty" name:"HostExtraSpec"`
	// 宿主机定制化参数，磁盘直通功能相关，为json格式字符串

	CheckInfo *string `json:"CheckInfo,omitempty" name:"CheckInfo"`
}

type CreateHostTypesRequest struct {
	*tchttp.BaseRequest

	// 宿主机类型名称，如："Y0-GI51-25G"

	HostTypeName *string `json:"HostTypeName,omitempty" name:"HostTypeName"`
	// 该宿主机类型可售卖的核数，单位个

	CPUTotal *uint64 `json:"CPUTotal,omitempty" name:"CPUTotal"`
	// 该宿主机类型可售卖的内存大小，单位GB

	MemTotal *uint64 `json:"MemTotal,omitempty" name:"MemTotal"`
	// 该宿主机类型可售卖的本地磁盘大小，单位GB

	DiskTotal *uint64 `json:"DiskTotal,omitempty" name:"DiskTotal"`
	// 该宿主机定制化参数，磁盘直通场景等需要填写

	CheckInfo *string `json:"CheckInfo,omitempty" name:"CheckInfo"`
	// 该宿主机定制化参数，和QOS服务相关

	HostExtraSpec *string `json:"HostExtraSpec,omitempty" name:"HostExtraSpec"`
	// 该宿主机类型可装箱的子机类型，如["S3", "S2"]

	InstanceFamilySet []*string `json:"InstanceFamilySet,omitempty" name:"InstanceFamilySet"`
	// 该宿主机支持的CPU Flag

	CPUFlags []*string `json:"CPUFlags,omitempty" name:"CPUFlags"`
	// 该宿主机的CPU架构，如Broadwell, Skylake

	CPUModel *string `json:"CPUModel,omitempty" name:"CPUModel"`
	// 该宿主机的CPU型号，如"Intel(R) Xeon(R) Gold 6146 CPU @ 3.20GHz"

	CPUModelName *string `json:"CPUModelName,omitempty" name:"CPUModelName"`
}

func (r *CreateHostTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHostTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyPairsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 密钥对信息

		KeyPairSet []*KeyPair `json:"KeyPairSet,omitempty" name:"KeyPairSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKeyPairsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseAgentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务信息

		TaskSet []*TaskSet `json:"TaskSet,omitempty" name:"TaskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseAgentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostsRequest struct {
	*tchttp.BaseRequest

	// zone-id-list:要查的可用区的id，net-version-list: 过滤要查询的宿主机的网络版本，host-asset-list: 固资号host-ip-list: 母机ipsold-pool-list: 待查询的售卖池列表; instance-family-list:过滤的机型规格  如 S3，host-hypervisor:虚拟化类型,可取值"kvm" "baremetal"

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 每页的数据

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 起始查找位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryHostInstanceMapResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// JSON格式的交叉索引返回数组

		SuccessList []*HostInstanceMap `json:"SuccessList,omitempty" name:"SuccessList"`
		// 失败列表

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 执行总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryHostInstanceMapResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryHostInstanceMapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTaskTranslogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务详情列表

		TaskTranslogItems []*TaskTranslogItems `json:"TaskTranslogItems,omitempty" name:"TaskTranslogItems"`
		// 返回数据的条数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryTaskTranslogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTaskTranslogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceFamilyConfig struct {
	// 实例族名称

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 优先级

	Order *uint64 `json:"Order,omitempty" name:"Order"`
	// 显示名称

	InstanceFamilyName *string `json:"InstanceFamilyName,omitempty" name:"InstanceFamilyName"`
	// 这个实例族最大可配置CPU数量（单位个）

	CPUTotal *int64 `json:"CPUTotal,omitempty" name:"CPUTotal"`
	// 这个实例族最大可配置内存大小（单位G）

	MemTotal *int64 `json:"MemTotal,omitempty" name:"MemTotal"`
	// 网卡数量（单位个）

	NetworkCardTotal *int64 `json:"NetworkCardTotal,omitempty" name:"NetworkCardTotal"`
	// 实例族的唯一标识

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 实例族当前的白名单

	WhiteListKey *string `json:"WhiteListKey,omitempty" name:"WhiteListKey"`
}

type DescribeCDHHostTypesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCDHHostTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCDHHostTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCvmApiLogStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CVM-API日志的统计结果

		CvmApiLogStatistics *string `json:"CvmApiLogStatistics,omitempty" name:"CvmApiLogStatistics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCvmApiLogStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCvmApiLogStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceFamilyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总共修改的数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 失败的实例族id列表

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 成功的实例族id列表

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceFamilyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceFamilyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceFamilyConfigArchitecture struct {
	// 实例族名称

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 优先级

	Order *uint64 `json:"Order,omitempty" name:"Order"`
	// 显示名称

	InstanceFamilyName *string `json:"InstanceFamilyName,omitempty" name:"InstanceFamilyName"`
	// 这个实例族最大可配置CPU数量（单位个）

	CPUTotal *int64 `json:"CPUTotal,omitempty" name:"CPUTotal"`
	// 这个实例族最大可配置内存大小（单位G）

	MemTotal *int64 `json:"MemTotal,omitempty" name:"MemTotal"`
	// 网卡数量（单位个）

	NetworkCardTotal *int64 `json:"NetworkCardTotal,omitempty" name:"NetworkCardTotal"`
	// 实例族的唯一标识

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 实例族当前的白名单

	WhiteListKey *string `json:"WhiteListKey,omitempty" name:"WhiteListKey"`
	// 当前实例族的CPU架构

	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
	// 显示名称（中文）

	InstanceFamilyNameCN *string `json:"InstanceFamilyNameCN,omitempty" name:"InstanceFamilyNameCN"`
	// 实例族磁盘类型

	DiskType *int64 `json:"DiskType,omitempty" name:"DiskType"`
}

type DescribeVsChildTaskIdByParentRequest struct {
	*tchttp.BaseRequest

	// CVM_VS 父任务ID列表

	VSParentTaskIds []*string `json:"VSParentTaskIds,omitempty" name:"VSParentTaskIds"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](/document/api/213/11646#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/api/213/11646#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeVsChildTaskIdByParentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVsChildTaskIdByParentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHostsTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteHostsTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHostsTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComputeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 母机日志文本

		ComputeLog *string `json:"ComputeLog,omitempty" name:"ComputeLog"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComputeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComputeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceFamilyConfigsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstanceFamilyConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceFamilyConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JobExtraSpecs struct {
	// 热迁移最大带宽

	MaxBandwidth *uint64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// 迁移最大超时时间

	MaxTimeout *uint64 `json:"MaxTimeout,omitempty" name:"MaxTimeout"`
}

type DescribeZoneImportInstanceTypeConfigsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：
	// zone-id:可用区ID
	// instance-family-list: 机型族列表

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeZoneImportInstanceTypeConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneImportInstanceTypeConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTaskReportErrorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务系统失败报表

		TaskReportErrorItems []*TaskReportErrorItems `json:"TaskReportErrorItems,omitempty" name:"TaskReportErrorItems"`
		// 返回数据的条数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryTaskReportErrorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTaskReportErrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebootInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RebootInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailoverMigrateInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FailoverMigrateInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FailoverMigrateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TypeList struct {
	// 宿主机类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 状态

	State *string `json:"State,omitempty" name:"State"`
}

type ModifyHostsResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总共操作的宿主机数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 成功修改的宿主机ip列表

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 失败的宿主机ip列表

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostsResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostsResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMQQueuesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMQQueuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMQQueuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCvmTypeMapWithHostInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例信息和对应宿主机信息列表

		CvmTypeAndHostInfoMapList []*CvmTypeAndHostInfo `json:"CvmTypeAndHostInfoMapList,omitempty" name:"CvmTypeAndHostInfoMapList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCvmTypeMapWithHostInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCvmTypeMapWithHostInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostUsbResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostUsbResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostUsbResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例相信信息列表

		InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskStep struct {
	// 任务游标

	Cursor *int64 `json:"Cursor,omitempty" name:"Cursor"`
	// 任务步骤对应模块

	Module *string `json:"Module,omitempty" name:"Module"`
	// 任务步骤对应命令

	Command *string `json:"Command,omitempty" name:"Command"`
	// 结果码

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 结果信息

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 开始时间

	STime *string `json:"STime,omitempty" name:"STime"`
	// 结束时间

	ETime *string `json:"ETime,omitempty" name:"ETime"`
	// 消费者进程id

	ConsumerId *string `json:"ConsumerId,omitempty" name:"ConsumerId"`
	// 消费者进程的ip

	ConsumerIp *string `json:"ConsumerIp,omitempty" name:"ConsumerIp"`
	// 日志文件名

	LogFileName *string `json:"LogFileName,omitempty" name:"LogFileName"`
}

type RollbackFailedTask struct {
	// 任务名

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 用户AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeVsChildTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志详情。

		TaskInfoSet []*VsTaskInfo `json:"TaskInfoSet,omitempty" name:"TaskInfoSet"`
		// 日志总量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVsChildTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVsChildTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务重试结果详情

		Results []*Results `json:"Results,omitempty" name:"Results"`
		// 错误信息

		Error *Results `json:"Error,omitempty" name:"Error"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CDHDataSet struct {
	// CDH机型名称

	ExternalMachineType *string `json:"ExternalMachineType,omitempty" name:"ExternalMachineType"`
	// 宿主机类型

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// cpu核数

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 内存大小

	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`
	// 硬盘大小

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// cdh机型关联宿主机个数

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type TaskInfoSet struct {
	// 接口名

	Action *string `json:"Action,omitempty" name:"Action"`
	// 添加时间；例：2019-07-05 00:00:00

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 账号ID

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 截止时间；例：2019-07-05 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// Input

	Input *string `json:"Input,omitempty" name:"Input"`
	// Output

	Output *string `json:"Output,omitempty" name:"Output"`
	// 状态

	State *string `json:"State,omitempty" name:"State"`
	// Vs任务ID

	VsTaskId *int64 `json:"VsTaskId,omitempty" name:"VsTaskId"`
	// 日志详细信息

	LogDetail *string `json:"LogDetail,omitempty" name:"LogDetail"`
	// ip地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 开始时间；例：2019-07-05 23:59:59

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 任务ID

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

type ColdMigrateInstancesRequest struct {
	*tchttp.BaseRequest

	// 要迁移的一个或者多个实例

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 迁移超时时间

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
	// 指定带宽

	MaxBandwidth *uint64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// 是否强制光机

	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
	// 指定要迁移的目标母鸡列表

	HostIps []*string `json:"HostIps,omitempty" name:"HostIps"`
	// 是否保留ip

	IpReserved *bool `json:"IpReserved,omitempty" name:"IpReserved"`
	// 指定售卖池冷迁移

	SoldPool *string `json:"SoldPool,omitempty" name:"SoldPool"`
	// 是否全部带宽

	AllBandwidth *bool `json:"AllBandwidth,omitempty" name:"AllBandwidth"`
}

func (r *ColdMigrateInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ColdMigrateInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentLogsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页限制，默认0

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAgentLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VstationLog struct {
	// Running任务个数

	Running *int64 `json:"Running,omitempty" name:"Running"`
	// Success任务个数

	Success *int64 `json:"Success,omitempty" name:"Success"`
	// Fail任务个数

	Fail *int64 `json:"Fail,omitempty" name:"Fail"`
	// RollbackFail任务个数

	RollbackFail *int64 `json:"RollbackFail,omitempty" name:"RollbackFail"`
}

type DescribeHostUsbRequest struct {
	*tchttp.BaseRequest

	// 宿主机IP

	Host *string `json:"Host,omitempty" name:"Host"`
}

func (r *DescribeHostUsbRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostUsbRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterHostReason struct {
	// 过滤器名

	FilterName *string `json:"FilterName,omitempty" name:"FilterName"`
	// 母机列表

	HostList *string `json:"HostList,omitempty" name:"HostList"`
}

type InstanceClassConfig struct {
	// 实例大类

	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`
	// 实例大类显示名称

	InstanceClassName *string `json:"InstanceClassName,omitempty" name:"InstanceClassName"`
	// 优先级

	Order *int64 `json:"Order,omitempty" name:"Order"`
	// 实例族信息

	InstanceFamilyConfigSet []*InstanceFamilyConfig `json:"InstanceFamilyConfigSet,omitempty" name:"InstanceFamilyConfigSet"`
}

type DescribeCDHExternalMachineTypeRequest struct {
	*tchttp.BaseRequest

	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCDHExternalMachineTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCDHExternalMachineTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailTaskRequest struct {
	*tchttp.BaseRequest

	// 任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 失败信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

func (r *FailTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FailTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentConfig struct {
	// Agent名称

	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`
	// OS类型

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 状态 online：在线  offline：离线

	Status *string `json:"Status,omitempty" name:"Status"`
	// 更新时间

	Update *string `json:"Update,omitempty" name:"Update"`
}

type DescribeSoldPoolsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 售卖池信息数组

		SoldPoolSet []*SoldPool `json:"SoldPoolSet,omitempty" name:"SoldPoolSet"`
		// 符合过滤条件的售卖池总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSoldPoolsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSoldPoolsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceTypeConfigRequest struct {
	*tchttp.BaseRequest

	// 1 - 售卖
	// 2 - 停止售卖

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 需要修改的机型配置名

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *ModifyInstanceTypeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceTypeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLatestOperationStateRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例最近操作的状态，主要有两种：SUCCESS和FAILED两种状态

	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`
	// 是否强制刷新

	Force *bool `json:"Force,omitempty" name:"Force"`
}

func (r *ModifyLatestOperationStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLatestOperationStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopInstancesRequest struct {
	*tchttp.BaseRequest

	// 关闭一个或者多个实例

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 是否在正常关闭失败后选择强制关闭实例。取值范围：<br><li>TRUE：表示在正常关闭失败后进行强制关闭<br><li>FALSE：表示在正常关闭失败后不进行强制关闭<br><br>默认取值：FALSE。

	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *StopInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TenantTask struct {
	// 租户端任务Id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 租户端资源操作的Instance id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 任务信息（报错）

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
}

type ModifyHostsTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostsTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostsTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCpuModelsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCpuModelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCpuModelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemDisk struct {
	// 系统盘类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 系统盘Id

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 系统盘大小

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 系统盘序列号

	DiskSerial *string `json:"DiskSerial,omitempty" name:"DiskSerial"`
	// 系统盘的镜像ID

	ImageUuid *string `json:"ImageUuid,omitempty" name:"ImageUuid"`
}

type DeleteSoldPoolRequest struct {
	*tchttp.BaseRequest

	// 要删除的售卖池的名称

	SoldPoolName *string `json:"SoldPoolName,omitempty" name:"SoldPoolName"`
}

func (r *DeleteSoldPoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSoldPoolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceUsbControllerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceUsbControllerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceUsbControllerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllinoneLog struct {
	// Success任务类型个数

	Success *int64 `json:"Success,omitempty" name:"Success"`
	// Failed任务类型个数

	Failed *int64 `json:"Failed,omitempty" name:"Failed"`
	// Pending任务类型个数

	Pending *int64 `json:"Pending,omitempty" name:"Pending"`
	// Running任务类型个数

	Running *int64 `json:"Running,omitempty" name:"Running"`
}

type Result struct {
	// 原状态

	// 错误码
}

type LiveMigrateInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LiveMigrateInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LiveMigrateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneInstanceTypeConfigRequest struct {
	*tchttp.BaseRequest

	// 可用区的zone id

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 要修改的实例机型config id

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 售卖状态：
	// 0：停止售卖
	// 1：开启售卖

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 售罄策略：
	// 0：按资源决定
	// 1：强制售罄

	SoldOutStrategy *int64 `json:"SoldOutStrategy,omitempty" name:"SoldOutStrategy"`
	// 本地盘类型

	LocalDiskTypeList []*LocalDiskType `json:"LocalDiskTypeList,omitempty" name:"LocalDiskTypeList"`
}

func (r *ModifyZoneInstanceTypeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyZoneInstanceTypeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCDHExternalMachineTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCDHExternalMachineTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCDHExternalMachineTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryFaultMigrateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件ID

		// 返回码

		// 组件名

		// 版本

		// 返回值

		// 时间戳

		// 弃用警告

		// 返回信息

		// 返回数据

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryFaultMigrateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryFaultMigrateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePublicImageStatusRequest struct {
	*tchttp.BaseRequest

	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 状态。上线(ONLINE)，下线(OFFLINE)两

	Status *string `json:"Status,omitempty" name:"Status"`
	// T_CLOUD_RESOURCE租户端、O_CLOUD_RESOURCE运营端

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *UpdatePublicImageStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePublicImageStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceInfoList struct {
	// 设备类型

	DeviceClass *string `json:"DeviceClass,omitempty" name:"DeviceClass"`
	// 过滤器名称

	FilterName *string `json:"FilterName,omitempty" name:"FilterName"`
	// 实例类型

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 类型名称

	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
	// 类型名称_en

	TypeNameEn *string `json:"TypeNameEn,omitempty" name:"TypeNameEn"`
}

type DescribeCDHExternalMachineTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CDH信息

		CDHTypeSet *CDHDataSet `json:"CDHTypeSet,omitempty" name:"CDHTypeSet"`
		// CDH信息总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCDHExternalMachineTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCDHExternalMachineTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSoldPoolsOwnerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 每个售卖池对应的用户列表

		SoldPoolOwnerSet []*SoldPoolOwner `json:"SoldPoolOwnerSet,omitempty" name:"SoldPoolOwnerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSoldPoolsOwnerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSoldPoolsOwnerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceSet struct{}

type ModifyInstanceTypeDisasterGroupConfigRequest struct {
	*tchttp.BaseRequest

	// 要管理的机型

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// False 添加到黑名单 , True 从黑名单移除

	SupportDisasterRecoverGroup *bool `json:"SupportDisasterRecoverGroup,omitempty" name:"SupportDisasterRecoverGroup"`
}

func (r *ModifyInstanceTypeDisasterGroupConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceTypeDisasterGroupConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Data struct {
	// 任务ID

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeSoldPoolsRequest struct {
	*tchttp.BaseRequest

	// 可选filter:
	// sold-pool-name-list: 要查询的售卖池名称

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 开始查询的位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 获取售卖池的个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSoldPoolsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSoldPoolsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTaskReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务系统报表

		TaskReportItems []*TaskReportItems `json:"TaskReportItems,omitempty" name:"TaskReportItems"`
		// 返回数据的条数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryTaskReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTaskReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCDHExternalMachineTypeRequest struct {
	*tchttp.BaseRequest

	// CDH机型名称

	ExternalMachineType *string `json:"ExternalMachineType,omitempty" name:"ExternalMachineType"`
	// 宿主机机型信息

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 地域id

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
}

func (r *CreateCDHExternalMachineTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCDHExternalMachineTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceTypeQuota struct {
	// 机型类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// CPU大小

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 内存大小

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// GPU大小

	GPU *uint64 `json:"GPU,omitempty" name:"GPU"`
	// FPGA核数

	FPGA *uint64 `json:"FPGA,omitempty" name:"FPGA"`
	// 存储快熟

	StorageBlock *uint64 `json:"StorageBlock,omitempty" name:"StorageBlock"`
	// 网卡数

	NetworkCard *uint64 `json:"NetworkCard,omitempty" name:"NetworkCard"`
	// 最大带宽

	MaxBandwidth *uint64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 配置Id

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 售卖状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 售罄规则。0：按资源决定，1：强制售罄

	SoldOutStrategy *int64 `json:"SoldOutStrategy,omitempty" name:"SoldOutStrategy"`
	// 本地盘信息

	LocalDiskTypeList []*LocalDiskType `json:"LocalDiskTypeList,omitempty" name:"LocalDiskTypeList"`
}

type JobInfo struct {
	// 是否全部带宽

	// 迁移目的ip

	// job传给vstation的参数
}

type CPUModelInfoArchitecture struct {
	// CPU架构名称，如"Broadwell"

	CPUModel *string `json:"CPUModel,omitempty" name:"CPUModel"`
	// CPU架构所支持的CPU Flags

	CPUFlags []*string `json:"CPUFlags,omitempty" name:"CPUFlags"`
	// CPU架构

	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
}

type HostFilterReason struct {
	// 母机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 过滤原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type DescribeZoneTenantInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 机型列表

		InstanceTypeQuotaSet []*InstanceTypeQuota `json:"InstanceTypeQuotaSet,omitempty" name:"InstanceTypeQuotaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneTenantInstanceConfigInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneTenantInstanceConfigInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePublicImageStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回信息

		Result *Result `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePublicImageStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePublicImageStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModifyRecycleTimeLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeModifyRecycleTimeLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModifyRecycleTimeLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTenantInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例的uuid数组，如["5182c761-90d7-4cd9-bbf8-4c19c4896dc8"]

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeTenantInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTenantInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVncRequest struct {
	*tchttp.BaseRequest

	// 要查看的VNC的实例的uuid

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeVncRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVncRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachUsbRequest struct {
	*tchttp.BaseRequest

	// 宿主机IP

	Host *string `json:"Host,omitempty" name:"Host"`
	// 虚拟机uuid

	VmUuid *string `json:"VmUuid,omitempty" name:"VmUuid"`
	// USB的VendorId

	VendorId *string `json:"VendorId,omitempty" name:"VendorId"`
	// USB的ProductId

	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
	// USB的DeviceNumber

	DeviceNumber *string `json:"DeviceNumber,omitempty" name:"DeviceNumber"`
	// USB的BusNumber

	BusNumber *string `json:"BusNumber,omitempty" name:"BusNumber"`
}

func (r *AttachUsbRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachUsbRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 宿主机详细信息列表

		HostTypeSet []*HostType `json:"HostTypeSet,omitempty" name:"HostTypeSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRecordRequest struct {
	*tchttp.BaseRequest

	// 查询起始值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// DescribeImageRecord

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 每一页个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 筛选任务创建时间的开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 筛选任务创建时间的结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 排序类型

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 排序方式

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeImageRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCvmApiLogStatisticsRequest struct {
	*tchttp.BaseRequest

	// 查询日志范围的起始时间；例：2019-07-05 00:00:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询日志范围的截止时间；例：2019-07-05 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeCvmApiLogStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCvmApiLogStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceUsbInfoRequest struct {
	*tchttp.BaseRequest

	// 虚拟机uuid

	VmUuids []*string `json:"VmUuids,omitempty" name:"VmUuids"`
}

func (r *DescribeInstanceUsbInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceUsbInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceClassConfigArchitecture struct {
	// 实例大类

	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`
	// 实例大类显示名称

	InstanceClassName *string `json:"InstanceClassName,omitempty" name:"InstanceClassName"`
	// 优先级

	Order *int64 `json:"Order,omitempty" name:"Order"`
	// 实例族信息

	InstanceFamilyConfigSet []*InstanceFamilyConfigArchitecture `json:"InstanceFamilyConfigSet,omitempty" name:"InstanceFamilyConfigSet"`
}

type Task struct {
	// 任务Id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务详细信息

	TaskInfo []*TaskInfo `json:"TaskInfo,omitempty" name:"TaskInfo"`
}

type DescribeComputeRequest struct {
	*tchttp.BaseRequest

	// 母机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 日志日期，对LogType为procedure,service,monitor类型时候必须传

	Date *string `json:"Date,omitempty" name:"Date"`
	// 日志类型：procedure,service,monitor,libvirt,qemu

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 实例的uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeComputeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComputeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDeviceStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDeviceStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDeviceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](DescribeInstances)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *StartInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，详见下表：实例过滤条件表。 instance-name: 实例名字筛选（模糊查找） host-ip-list: 根据母机ip筛选 network-device-id-list：根据上联网络设备号筛选 status-list：根据子机状态列表筛选 zone-id-list： 根据可用区筛选 net-version-list：根据网络类型筛选，VPC:租户网络，CLASSIC：物理网络，net-version-list目前只支持两者其一 operator-uin-list：根据实例创建者uin类型筛选

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，常常和Limit结合

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询的数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostsCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		Count *uint64 `json:"Count,omitempty" name:"Count"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHostsCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostsCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceTypeConfig struct {
	// 实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// CPU个数

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 内存大小

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// GPU核数

	GPU *uint64 `json:"GPU,omitempty" name:"GPU"`
	// FPGA核数

	FPGA *uint64 `json:"FPGA,omitempty" name:"FPGA"`
	// 存储块数

	StorageBlock *uint64 `json:"StorageBlock,omitempty" name:"StorageBlock"`
	// 网卡数

	NetworkCard *uint64 `json:"NetworkCard,omitempty" name:"NetworkCard"`
	// 最大带宽

	MaxBandwidth *float64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 配置Id

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 售卖状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 表示当前 cpu 架构

	CpuModelName *string `json:"CpuModelName,omitempty" name:"CpuModelName"`
	// 虚机网络转发性能

	Pps *uint64 `json:"Pps,omitempty" name:"Pps"`
	// 表示物理机核心最高频率

	Frequency *string `json:"Frequency,omitempty" name:"Frequency"`
	// 备注信息(英文)

	RemarkEn *string `json:"RemarkEn,omitempty" name:"RemarkEn"`
}

type CreateCvmTypeMapRequest struct {
	*tchttp.BaseRequest

	// 新增机型归属的机型簇

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 新增机型自定义类型

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 新增机型自定义类型中文名

	InstanceFamilyChnName *string `json:"InstanceFamilyChnName,omitempty" name:"InstanceFamilyChnName"`
	// 新增机型自定义类型英文名

	InstanceFamilyEngName *string `json:"InstanceFamilyEngName,omitempty" name:"InstanceFamilyEngName"`
	// 内部代码

	DeviceClass *string `json:"DeviceClass,omitempty" name:"DeviceClass"`
	// 新增机型的CPU架构

	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
}

func (r *CreateCvmTypeMapRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCvmTypeMapRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回当前创建任务的task_id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCDHExternalMachineTypeRequest struct {
	*tchttp.BaseRequest

	// CDH机型名称

	ExternalMachineType *string `json:"ExternalMachineType,omitempty" name:"ExternalMachineType"`
}

func (r *DeleteCDHExternalMachineTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCDHExternalMachineTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCvmTypeMapResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定结果

		UpdateResult *string `json:"UpdateResult,omitempty" name:"UpdateResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCvmTypeMapResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCvmTypeMapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrantLiveMigrateRequest struct {
	*tchttp.BaseRequest

	// 要做批量热迁移的实例uuid列表

	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
	// 用作角色权限判断的，TCE上不使用，所以UserName为空就行

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 传给Job组件的参数

	JobInfo *JobInfo `json:"JobInfo,omitempty" name:"JobInfo"`
}

func (r *MigrantLiveMigrateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MigrantLiveMigrateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总共的镜像个数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 查询的镜像的信息详情。

		ImageSet []*Image `json:"ImageSet,omitempty" name:"ImageSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceClassConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 所有的实例系列，如“S”，“D”，“I”

		InstanceClassConfigSet []*InstanceClassConfigArchitecture `json:"InstanceClassConfigSet,omitempty" name:"InstanceClassConfigSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceClassConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceClassConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVsChildTaskIdByParentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 父任务到子任务的映射列表。

		ParentToChildSet []*ParentToChild `json:"ParentToChildSet,omitempty" name:"ParentToChildSet"`
		// 总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVsChildTaskIdByParentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVsChildTaskIdByParentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindCpuCoresRequest struct {
	*tchttp.BaseRequest

	// 实例uuid

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 待绑定的cpu核心的数组

	CPUSet []*string `json:"CPUSet,omitempty" name:"CPUSet"`
}

func (r *BindCpuCoresRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindCpuCoresRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyPair struct {
	// 唯一id

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 密钥名称

	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
}

type UnbindSoldPoolOwnersRequest struct {
	*tchttp.BaseRequest

	// 售卖池名称

	SoldPoolName *string `json:"SoldPoolName,omitempty" name:"SoldPoolName"`
	// 从售卖池中解绑的账号列表

	OwnerList []*string `json:"OwnerList,omitempty" name:"OwnerList"`
}

func (r *UnbindSoldPoolOwnersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindSoldPoolOwnersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 宿主机资源信息

		HostItem []*HostItem `json:"HostItem,omitempty" name:"HostItem"`
		// 分页所用信息

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckCreatePublicImageNameRequest struct {
	*tchttp.BaseRequest

	// 源os name

	SourceOsName *string `json:"SourceOsName,omitempty" name:"SourceOsName"`
	// 后缀

	Postfix *string `json:"Postfix,omitempty" name:"Postfix"`
}

func (r *CheckCreatePublicImageNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckCreatePublicImageNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMQQueuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// MQ队列信息

		Queues []*MQQueueInfo `json:"Queues,omitempty" name:"Queues"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMQQueuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMQQueuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckCreatePublicImageNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckCreatePublicImageNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckCreatePublicImageNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSoldPoolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除操作的售卖池数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 成功删除的售卖池列表

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 删除失败的售卖池列表

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSoldPoolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSoldPoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest

	// image-uuid-list: 传入镜像的uuid列表，获取镜像的详细信息
	// image-id-list: 传入镜像的短id列表
	// appid: 传入镜像所有者的appId
	// exclude-appid-list: 不显示列表中的appid拥有的镜像

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// PUBLIC_IMAGE - 公共镜像
	// PRIVATE_IMAGE -私有镜像
	// 默认为公共镜像

	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`
	// 开始查找的位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 租户端运营端类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 是否展示下线镜像,如果不展示下线镜像(IsShowOffline为False时)需要同时表示镜像类型(租户端 or 运营端)

	IsShowOffline *bool `json:"IsShowOffline,omitempty" name:"IsShowOffline"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSchedulerRequest struct {
	*tchttp.BaseRequest

	// 任务Id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 过滤条件，详见下表：实例过滤条件表。每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`InstanceIds`和`Filters`。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeSchedulerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSchedulerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRecycleInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *GetRecycleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRecycleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SchedulerInfo struct {
	// cpu核数

	CPU *int64 `json:"CPU,omitempty" name:"CPU"`
	// 内存大小

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 磁盘大小

	Disk *int64 `json:"Disk,omitempty" name:"Disk"`
	// 用户AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 售卖池

	SoldPool *string `json:"SoldPool,omitempty" name:"SoldPool"`
	// 实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 地域Id

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// Cbs标识

	CbsFlag *string `json:"CbsFlag,omitempty" name:"CbsFlag"`
	// 母机类型

	PlainHostType *string `json:"PlainHostType,omitempty" name:"PlainHostType"`
	// VPC的id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 母机列表

	HostSet []*HostFilterReason `json:"HostSet,omitempty" name:"HostSet"`
	// 过滤列表

	FilterSet []*FilterHostReason `json:"FilterSet,omitempty" name:"FilterSet"`
}

type ResetInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowLogsRequest struct {
	*tchttp.BaseRequest

	// 查询日志范围的起始时间；例：2019-07-05 00:00:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询日志范围的截止时间；例：2019-07-05 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件，详见下表：实例过滤条件表。每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`InstanceIds`和`Filters`。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](/document/api/213/11646#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/api/213/11646#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeFlowLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BasicallyCompletedTask struct {
	// 任务名

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type QueryTaskTranslogRequest struct {
	*tchttp.BaseRequest

	// task-id:过滤要查询的任务id，instance-uuid: 过滤要查询的实例uuid，host-ip：过滤要查询的宿主机ip，task-name: 过滤要查询的任务类型，status: 过滤要查询的任务状态

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 起始查找位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页的数据

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询 VStation 任务详情列表的起始时间；例：2019-07-05 00:00:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询 VStation 任务详情列表的截止时间；例：2019-07-05 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *QueryTaskTranslogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTaskTranslogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAllHostTypesRequest struct {
	*tchttp.BaseRequest

	// overlay|underlay

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 投放机型归属的实例类型，overlay 多选，underlay 不能选（需要默认参数）

	InstanceFamilySet []*string `json:"InstanceFamilySet,omitempty" name:"InstanceFamilySet"`
	// CPU 名称

	CPUModelName *string `json:"CPUModelName,omitempty" name:"CPUModelName"`
	// CPU架构

	CPUModel *string `json:"CPUModel,omitempty" name:"CPUModel"`
	// CPU 指令集

	CPUFlags *string `json:"CPUFlags,omitempty" name:"CPUFlags"`
	// 物理机名称

	HostType *string `json:"HostType,omitempty" name:"HostType"`
	// 可售卖 CPU 总核心数

	CPUTotal *string `json:"CPUTotal,omitempty" name:"CPUTotal"`
	// 可售卖 MEM 总数（MB）

	MemTotal *string `json:"MemTotal,omitempty" name:"MemTotal"`
	// 可售卖 DISK 总数（MB）

	DiskTotal *string `json:"DiskTotal,omitempty" name:"DiskTotal"`
	// JSON格式预留字段，涉及 pcie 直通设备的时候使用

	HostExtraSpec *string `json:"HostExtraSpec,omitempty" name:"HostExtraSpec"`
	// 预留字段：涉及调度检查参数的时候使用

	CheckInfo *string `json:"CheckInfo,omitempty" name:"CheckInfo"`
}

func (r *CreateAllHostTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAllHostTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachUsbResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachUsbResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachUsbResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyPairsRequest struct {
	*tchttp.BaseRequest

	// 要查询的过滤器： uuid-list：要查询的密钥对所属的子机列表（现在只支持填一个）

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeKeyPairsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKeyPairsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySoldPoolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 成功列表

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 失败列表

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySoldPoolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySoldPoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceFamilyConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceFamilyConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceFamilyConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskSet struct {
	// 任务id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 请求状态

	Message *string `json:"Message,omitempty" name:"Message"`
}

type DescribeLogStatisticsRequest struct {
	*tchttp.BaseRequest

	// 统计日志的起始时间；例：2019-07-05 00:00:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 统计日志的截止时间；例：2019-07-05 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeLogStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest

	// 一个或者多个实例

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 实例显示名字

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyInstancesAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryTaskRequest struct {
	*tchttp.BaseRequest

	// 任务id数组

	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`
	// 0表示重试当前，1跳过一步，-1后退一步

	Pass *int64 `json:"Pass,omitempty" name:"Pass"`
}

func (r *RetryTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CPUModelInfo struct {
	// CPU架构名称，如"Broadwell"

	CPUModel *string `json:"CPUModel,omitempty" name:"CPUModel"`
	// CPU架构所支持的CPU Flags

	CPUFlags []*string `json:"CPUFlags,omitempty" name:"CPUFlags"`
}

type SyncImagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTaskReportErrorRequest struct {
	*tchttp.BaseRequest

	// 起始查找位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页的数据

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询 VStation 任务系统失败报表的起始时间；例：2019-07-05 00:00:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询 VStation 任务系统失败报表的截止时间；例：2019-07-05 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 是否需要关注 (1表示 需关注 0 表示忽略)

	Current []*string `json:"Current,omitempty" name:"Current"`
	// 异常分类， MSG_LOST 消息丢失。FATAL_ERROR 致命错误。BYPASS_ERROR  bypass错误 ROLLBACK_ERROR 回滚错误

	ErrorType []*string `json:"ErrorType,omitempty" name:"ErrorType"`
}

func (r *QueryTaskReportErrorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTaskReportErrorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostTypesRequest struct {
	*tchttp.BaseRequest

	// 可支持过滤key:
	// host-type-list：按照HostTypeName进行过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeHostTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DesRunningOvertimeTask struct {
	// 接口名

	Action *string `json:"Action,omitempty" name:"Action"`
	// 实例列表

	InstanceList *string `json:"InstanceList,omitempty" name:"InstanceList"`
	// Uuid列表

	UuidList *string `json:"UuidList,omitempty" name:"UuidList"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 订单号

	DealName *string `json:"DealName,omitempty" name:"DealName"`
	// Allinone任务Id

	RunningDesTask *string `json:"RunningDesTask,omitempty" name:"RunningDesTask"`
	// 订单详情

	DealInfo *string `json:"DealInfo,omitempty" name:"DealInfo"`
}

type InstanceConfigInfo struct {
	// 实例大类

	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`
	// 实例大类显示名称

	InstanceClassName *string `json:"InstanceClassName,omitempty" name:"InstanceClassName"`
	// 优先级

	Order *int64 `json:"Order,omitempty" name:"Order"`
	// 实例族信息

	InstanceFamilySet *int64 `json:"InstanceFamilySet,omitempty" name:"InstanceFamilySet"`
}

type TaskTranslogItems struct {
	// 错误码

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 任务的创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 母机 IP

	Host *string `json:"Host,omitempty" name:"Host"`
	// 虚拟机的虚拟化类型

	Hypervisor *string `json:"Hypervisor,omitempty" name:"Hypervisor"`
	// 虚拟机镜像 id

	ImgId *string `json:"ImgId,omitempty" name:"ImgId"`
	// 虚拟机 IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 错误信息

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 虚拟机操作系统

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 用戶

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 任务的状态(0:运行中、1 成功、2 失败、3 回滚失败、4 基本完成)

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// vs 的内部任务 id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 虚拟机 uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 任务的详细参数

	Value *string `json:"Value,omitempty" name:"Value"`
}

type FlowLogSet struct {
	// 接口名称

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 接口名称

	InterfaceName *string `json:"InterfaceName,omitempty" name:"InterfaceName"`
	// 日志源

	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`
	// uin

	OperateSubAccountUin *string `json:"OperateSubAccountUin,omitempty" name:"OperateSubAccountUin"`
	// Next Id

	NextId *string `json:"NextId,omitempty" name:"NextId"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 实例id

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 组间

	Component *string `json:"Component,omitempty" name:"Component"`
	// 错误码

	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 日志类型

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// appid

	OperateAppId *string `json:"OperateAppId,omitempty" name:"OperateAppId"`
	// 输出数据

	OutputData *string `json:"OutputData,omitempty" name:"OutputData"`
	// url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 调用者

	Caller *string `json:"Caller,omitempty" name:"Caller"`
	// des任务id

	DESTaskId *string `json:"DESTaskId,omitempty" name:"DESTaskId"`
	// uin

	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`
	// Callee

	Callee *string `json:"Callee,omitempty" name:"Callee"`
	// 新des任务id

	NewDESTaskId *string `json:"NewDESTaskId,omitempty" name:"NewDESTaskId"`
	// vs父任务id

	VSParentTaskId *string `json:"VSParentTaskId,omitempty" name:"VSParentTaskId"`
	// 错误信息

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 输入数据

	InputData *string `json:"InputData,omitempty" name:"InputData"`
}

type TroubleShootingTask struct {
	// 任务名

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务Id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 用户AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribePreSchedulerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例售罄的原因结果列表

		SchedulerInfoSet []*SchedulerInfo `json:"SchedulerInfoSet,omitempty" name:"SchedulerInfoSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePreSchedulerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePreSchedulerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParentToChild struct {
	// 父任务task id

	ParentTaskId *int64 `json:"ParentTaskId,omitempty" name:"ParentTaskId"`
	// 子任务task id

	ChildTaskId *int64 `json:"ChildTaskId,omitempty" name:"ChildTaskId"`
}

type ModifyInstanceTypeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 尝试修改的配置个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 修改失败的配置ID列表

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 修改成功的配置ID列表

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceTypeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceTypeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserCpuTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定结果

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserCpuTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserCpuTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCDHHostTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CDH宿主机信息

		TypeList *TypeList `json:"TypeList,omitempty" name:"TypeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCDHHostTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCDHHostTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncInstanceFamilyConfigsRequest struct {
	*tchttp.BaseRequest

	// 需要同步到的可用区

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 需要同步的实例类型

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
}

func (r *SyncInstanceFamilyConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncInstanceFamilyConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可用区内机型配置详细列表

		InstanceTypeQuotaSet []*InstanceTypeQuota `json:"InstanceTypeQuotaSet,omitempty" name:"InstanceTypeQuotaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneInstanceConfigInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneInstanceConfigInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostsCountRequest struct {
	*tchttp.BaseRequest

	// 网络类型

	NetWorkType *string `json:"NetWorkType,omitempty" name:"NetWorkType"`
}

func (r *GetHostsCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostsCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginSettings struct {
	// 实例登录密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。

	Username *string `json:"Username,omitempty" name:"Username"`
}

type DescribeInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 全局机器类型详细列表

		InstanceTypeConfigSet []*InstanceTypeConfig `json:"InstanceTypeConfigSet,omitempty" name:"InstanceTypeConfigSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceConfigInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceConfigInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryHostInstanceMapRequest struct {
	*tchttp.BaseRequest

	// 机型筛选参数

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

func (r *QueryHostInstanceMapRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryHostInstanceMapRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSoldPoolRequest struct {
	*tchttp.BaseRequest

	// 售卖池的名称，唯一标识

	SoldPoolName *string `json:"SoldPoolName,omitempty" name:"SoldPoolName"`
	// 售卖比例

	SoldRate *float64 `json:"SoldRate,omitempty" name:"SoldRate"`
	// 售卖池的描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateSoldPoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSoldPoolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachUsbResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachUsbResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachUsbResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllinoneTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志详情。

		TaskInfoSet []*TaskInfoSet `json:"TaskInfoSet,omitempty" name:"TaskInfoSet"`
		// 日志总量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllinoneTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllinoneTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncImagesRequest struct {
	*tchttp.BaseRequest

	// 镜像ID列表 ，镜像ID可以通过如下方式获取：<br><li>通过[DescribeImages](/ocloud/api/Compute/CVM/APIs/cvm运营端（opcvm）/版本（2019-06-25）/cvm运营端/DescribeImages)接口返回的`ImageId`获取。<br><li>通过[[镜像控制台](//console.{{conf.main_domain}}/cvm/image)](https://console.cloud.tencent.com/cvm/image)获取。<br>镜像ID必须满足限制：<br><li>镜像ID对应的镜像状态必须为`NORMAL`。<br><li>镜像大小小于50GB。<br>镜像状态请参考[镜像数据表](/ocloud/api/Compute/CVM/APIs/cvm运营端（opcvm）/版本（2019-06-25）/数据结构#image)。

	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
	// 目的同步地域列表；必须满足限制：<br><li>不能为源地域，<br><li>必须是一个合法的Region。<br><li>暂不支持部分地域同步。

	DestinationRegions []*string `json:"DestinationRegions,omitempty" name:"DestinationRegions"`
}

func (r *SyncImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncImagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总共操作的机器个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 成功的机器的uuid列表

		SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`
		// 失败的机器的uuid列表

		FailList []*string `json:"FailList,omitempty" name:"FailList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstancesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncHostTypesRequest struct {
	*tchttp.BaseRequest

	// 需要同步过来的地域名称，如“ap-guangzhou”

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

func (r *SyncHostTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncHostTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SoldPool struct {
	// 售卖池名称

	SoldPoolName *string `json:"SoldPoolName,omitempty" name:"SoldPoolName"`
	// 售卖池描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 售卖比例

	SoldRate *float64 `json:"SoldRate,omitempty" name:"SoldRate"`
	// 售卖池内的宿主机个数

	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
	// 售卖池内的账号数

	OwnerCount *int64 `json:"OwnerCount,omitempty" name:"OwnerCount"`
}

type DescribeTroubleshootingPanelRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，主要指定过滤返回的列表，Filter.Name为return-type,Filter.Values主要为rollback-failed，basically-completed，operating-instance，des-running-overtime，message-lost，fatal-error，query-hosts

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 查询故障列表的起始时间；例：2019-07-05 00:00:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询故障列表的截止时间；例：2019-07-05 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeTroubleshootingPanelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTroubleshootingPanelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTaskReportRequest struct {
	*tchttp.BaseRequest

	// 起始查找位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页的数据

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询 VStation 任务系统报表的起始时间；例：2019-07-05 00:00:00 只支持天数筛选,不支持小时、分钟和秒筛选

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询 VStation 任务系统报表的截止时间；例：2019-07-05 00:00:00 只支持天数筛选,不支持小时、分钟和秒筛选

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 是否有效状态 (1 有效, 0 无效  )

	Current []*string `json:"Current,omitempty" name:"Current"`
}

func (r *QueryTaskReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTaskReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest

	// 要中期的实例id

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 是否强制重启

	ForceReboot *bool `json:"ForceReboot,omitempty" name:"ForceReboot"`
}

func (r *RebootInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Logs

		Logs []*AgentLog `json:"Logs,omitempty" name:"Logs"`
		// Logs的总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Instance struct {
	// CPU核数

	CPU *int64 `json:"CPU,omitempty" name:"CPU"`
	// 内存大小

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 实例名称(alias)

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例运行状态(realStatus)

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 内网IP地址

	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	// 所在母机Ip地址

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 模块Id

	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`
	// 上联网络设备Id(netdev_asset_id)

	NetworkDeviceId *string `json:"NetworkDeviceId,omitempty" name:"NetworkDeviceId"`
	// 对应asset

	Asset *string `json:"Asset,omitempty" name:"Asset"`
	// 可用区Id

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 操作系统名字

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 原uuid

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 拥有者名字

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 完成时间

	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
	// vpcId

	VirtualPrivateCloudId *int64 `json:"VirtualPrivateCloudId,omitempty" name:"VirtualPrivateCloudId"`
	// 设备Id

	DeviceId *int64 `json:"DeviceId,omitempty" name:"DeviceId"`
	// 数据卷

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 系统卷

	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// NUMA情况

	NodeQuota []*uint64 `json:"NodeQuota,omitempty" name:"NodeQuota"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 操作者uin

	OperatorUin *string `json:"OperatorUin,omitempty" name:"OperatorUin"`
	// 镜像id

	UnImgId *string `json:"UnImgId,omitempty" name:"UnImgId"`
}

type TaskReportErrorItems struct {
	// 任务 id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 虚拟机 uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 所在的母机 IP

	Host *string `json:"Host,omitempty" name:"Host"`
	// 异常分类， MSG_LOST 消息丢失。FATAL_ERROR 致命错误。BYPASS_ERROR  bypass错误 ROLLBACK_ERROR 回滚错误

	ErrorType *string `json:"ErrorType,omitempty" name:"ErrorType"`
	// 任务开始时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 是否需要关注 (1表示 需关注 0 表示忽略)

	Current *int64 `json:"Current,omitempty" name:"Current"`
	// 上次检测异常时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
}

type DescribeCvmTypeMapWithHostInfosRequest struct {
	*tchttp.BaseRequest

	// 每页的数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCvmTypeMapWithHostInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCvmTypeMapWithHostInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CvmTypeAndHostInfo struct {
	// 机型中文名称

	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
	// 机型英文名称

	TypeNameEn *string `json:"TypeNameEn,omitempty" name:"TypeNameEn"`
	// 宿主机类型

	HostType *string `json:"HostType,omitempty" name:"HostType"`
	// 宿主机CPU型号

	HostCpuModelName *string `json:"HostCpuModelName,omitempty" name:"HostCpuModelName"`
	// 机型的type

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 机型实例族

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// CPU架构

	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
}

type DescribeInstanceTypeConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceTypeConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceTypeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTroubleshootingPanelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 根据过滤字段rollback-failed才会展示，表示回滚失败任务结果

		RollbackFailedTaskList []*RollbackFailedTask `json:"RollbackFailedTaskList,omitempty" name:"RollbackFailedTaskList"`
		// 根据过滤字段basically-completed才会展示，表示基本完成任务结果

		BasicallyCompletedTaskList []*BasicallyCompletedTask `json:"BasicallyCompletedTaskList,omitempty" name:"BasicallyCompletedTaskList"`
		// 根据过滤字段operating-instance才会展示，表示CCDB长期operatiing任务结果

		OperatingInstanceTaskList []*OperatingInstanceTask `json:"OperatingInstanceTaskList,omitempty" name:"OperatingInstanceTaskList"`
		// 根据过滤字段des-running-overtime才会展示，表示Allinone任务运行超过30分钟的任务结果

		DesRunningOvertimeTaskList []*DesRunningOvertimeTask `json:"DesRunningOvertimeTaskList,omitempty" name:"DesRunningOvertimeTaskList"`
		// 根据过滤字段message-lost才会展示，表示消息丢失的任务结果

		MessageLostTaskList []*TroubleShootingTask `json:"MessageLostTaskList,omitempty" name:"MessageLostTaskList"`
		// 根据过滤字段fatal-error才会展示，表示失败后没有触发回滚的任务结果

		FatalErrorTaskList []*TroubleShootingTask `json:"FatalErrorTaskList,omitempty" name:"FatalErrorTaskList"`
		// 根据过滤字段query-hosts才会展示，表示状态异常的宿主机列表

		QueryHostsTaskList []*HostItem `json:"QueryHostsTaskList,omitempty" name:"QueryHostsTaskList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTroubleshootingPanelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTroubleshootingPanelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostInstanceMap struct {
	// 主机类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// ip

	IP *string `json:"IP,omitempty" name:"IP"`
	// 系统名称

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 内核版本

	KernelVersion *string `json:"KernelVersion,omitempty" name:"KernelVersion"`
	// CPU型号名称

	CpuModelName *string `json:"CpuModelName,omitempty" name:"CpuModelName"`
	// CPU名称

	CpuModel *string `json:"CpuModel,omitempty" name:"CpuModel"`
	// 主机状态

	State *string `json:"State,omitempty" name:"State"`
	// 设备信息列表

	DeviceInfoList []*DeviceInfoList `json:"DeviceInfoList,omitempty" name:"DeviceInfoList"`
}

type BindCpuCoresResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定结果

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindCpuCoresResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindCpuCoresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHostsTagRequest struct {
	*tchttp.BaseRequest

	// 要处理的宿主机列表

	HostIps []*string `json:"HostIps,omitempty" name:"HostIps"`
	// 要删除的标签名，如：HOST_TEMP_FORBID_LAUNCH

	Tag *string `json:"Tag,omitempty" name:"Tag"`
	// 操作原因。填写增加时的原因。

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 本次的操作人。用于追溯信息使用。

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 标签的描述。详细的描述一下标签的说明。

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *DeleteHostsTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHostsTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
