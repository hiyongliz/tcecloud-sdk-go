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

package v20200512

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type ProductFailureMigrateTask struct {
	// 切换任务UUID

	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 进度

	Progress *string `json:"Progress,omitempty" name:"Progress"`
	// 执行信息/错误信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 详情

	Detail []*Detail `json:"Detail,omitempty" name:"Detail"`
	// 补充数据

	ExtendData *string `json:"ExtendData,omitempty" name:"ExtendData"`
}

type CapacityStatisticSet struct {
	// 集群key

	Clusterkey *string `json:"Clusterkey,omitempty" name:"Clusterkey"`
	// 机器

	Machine *string `json:"Machine,omitempty" name:"Machine"`
	// CPU可用量

	CpuFree *int64 `json:"CpuFree,omitempty" name:"CpuFree"`
	// CPU总量

	CpuTotal *int64 `json:"CpuTotal,omitempty" name:"CpuTotal"`
	// 内存可用量

	MemoryFree *int64 `json:"MemoryFree,omitempty" name:"MemoryFree"`
	// 内存总量

	MemoryTotal *int64 `json:"MemoryTotal,omitempty" name:"MemoryTotal"`
	// 数据盘可用量

	DataDiskFree *int64 `json:"DataDiskFree,omitempty" name:"DataDiskFree"`
	// 数据盘总量

	DataDiskTotal *int64 `json:"DataDiskTotal,omitempty" name:"DataDiskTotal"`
	// 机器总数

	MachineTotal *int64 `json:"MachineTotal,omitempty" name:"MachineTotal"`
}

type ProductDataReplicationState struct {
	// 产品名

	Product *string `json:"Product,omitempty" name:"Product"`
	// 数据复制成功个数

	DataReplicationSuccess *int64 `json:"DataReplicationSuccess,omitempty" name:"DataReplicationSuccess"`
	// 数据复制失败个数

	DataReplicationFailed *int64 `json:"DataReplicationFailed,omitempty" name:"DataReplicationFailed"`
	// 实体列表

	EntityList []*Entity `json:"EntityList,omitempty" name:"EntityList"`
	// 状态最后上报时间（更新时间）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type QueryCreateProductFailureMigrateTaskDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品故障转移结果

		Data *ProductFailureMigrateTask `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCreateProductFailureMigrateTaskDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCreateProductFailureMigrateTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群id

		ClusterSet []*ClusterSet `json:"ClusterSet,omitempty" name:"ClusterSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductDataReplicationStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品实例数据同步状态信息

		Data *ProductDataReplicationState `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryProductDataReplicationStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductDataReplicationStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductDataReplicationStateRequest struct {
	*tchttp.BaseRequest

	// 产品名

	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *QueryProductDataReplicationStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductDataReplicationStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Entity struct {
	// 实体唯一标识，不同产品实体类型不同

	EntityID *string `json:"EntityID,omitempty" name:"EntityID"`
	// 数据同步状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 容灾模式

	FailToleranceModel *string `json:"FailToleranceModel,omitempty" name:"FailToleranceModel"`
	// 数据复制单位

	DataReplicationUnit *string `json:"DataReplicationUnit,omitempty" name:"DataReplicationUnit"`
	// 数据复制延迟

	DataReplicationDelay *string `json:"DataReplicationDelay,omitempty" name:"DataReplicationDelay"`
	// 数据复制模式

	DataReplicationModel *string `json:"DataReplicationModel,omitempty" name:"DataReplicationModel"`
	// 主副本所在AZ

	MasterCopyAzList []*string `json:"MasterCopyAzList,omitempty" name:"MasterCopyAzList"`
	// 备副本所在AZ

	BackupCopyAzList []*string `json:"BackupCopyAzList,omitempty" name:"BackupCopyAzList"`
}

type Specinfo struct {
	// 集群key

	ClusterKey *string `json:"ClusterKey,omitempty" name:"ClusterKey"`
	// CPU大小

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 内存大小

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 数据盘大小

	DataDisk *int64 `json:"DataDisk,omitempty" name:"DataDisk"`
	// 机器

	Machine *string `json:"Machine,omitempty" name:"Machine"`
	// 日志盘大小

	LogDisk *int64 `json:"LogDisk,omitempty" name:"LogDisk"`
}

type Detail struct {
	// 步骤

	StepName *string `json:"StepName,omitempty" name:"StepName"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 执行相关信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 开始时间

	StartAt *string `json:"StartAt,omitempty" name:"StartAt"`
	// 结束时间

	EndAt *string `json:"EndAt,omitempty" name:"EndAt"`
}

type CreateProductFailureMigrateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品实例数据同步状态信息

		Data []*TaskUUID `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductFailureMigrateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductFailureMigrateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Metrics struct {
	// 实例ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 实例状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 当前实例读写模式

	ReadOnly *string `json:"ReadOnly,omitempty" name:"ReadOnly"`
	// 当前实例复制模式

	DataReplicationModel *string `json:"DataReplicationModel,omitempty" name:"DataReplicationModel"`
	// 当前实例网关列表

	ProxyList *string `json:"ProxyList,omitempty" name:"ProxyList"`
}

type QueryProductHealthStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 一个或者多个维度的健康状态和总体健康状态

		Data []*Clusterhealth `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryProductHealthStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductHealthStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例信息

		InstanceSet []*InstanceSet `json:"InstanceSet,omitempty" name:"InstanceSet"`
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

type QueryProductStateInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品集群状态和节点的分布情况

		Data *Cluster `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryProductStateInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductStateInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskUUID struct {
	// 切换任务UUID

	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
}

type Cluster struct {
	// 产品集群中文名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 产品集群key

	ClusterKey *string `json:"ClusterKey,omitempty" name:"ClusterKey"`
	// 产品集群容灾状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 主节点所在AZ

	MainNodeAzList []*string `json:"MainNodeAzList,omitempty" name:"MainNodeAzList"`
	// 节点列表

	NodeList []*Node `json:"NodeList,omitempty" name:"NodeList"`
	// 状态最后上报时间（更新时间）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type InstanceSet struct {
	// 集群key

	ClusterKey *string `json:"ClusterKey,omitempty" name:"ClusterKey"`
	// 集群名称

	ClientName *string `json:"ClientName,omitempty" name:"ClientName"`
	// 实例名称+setId

	Mid *string `json:"Mid,omitempty" name:"Mid"`
	// 实例名称

	PMid *string `json:"PMid,omitempty" name:"PMid"`
	// 集群类型

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 机器ip

	IP *string `json:"IP,omitempty" name:"IP"`
	// 机器Port

	Port *string `json:"Port,omitempty" name:"Port"`
	// 实例setId

	ID *string `json:"ID,omitempty" name:"ID"`
	// porxy Ip:Port

	Proxy *string `json:"Proxy,omitempty" name:"Proxy"`
	// 实例核数

	CPU *string `json:"CPU,omitempty" name:"CPU"`
	// 实例内存数

	Memory *string `json:"Memory,omitempty" name:"Memory"`
	// 实例逻辑磁盘大小

	LogDisk *string `json:"LogDisk,omitempty" name:"LogDisk"`
	// 实例数据盘大小

	DataDisk *string `json:"DataDisk,omitempty" name:"DataDisk"`
	// 机器类型

	Machine *string `json:"Machine,omitempty" name:"Machine"`
	// 是否可退化

	DegradeFlag *string `json:"DegradeFlag,omitempty" name:"DegradeFlag"`
	// 实例状态，0

	Status *string `json:"Status,omitempty" name:"Status"`
	// 同步模式

	SqlAsyn *string `json:"SqlAsyn,omitempty" name:"SqlAsyn"`
	// range

	Range *string `json:"Range,omitempty" name:"Range"`
	// 节点数

	NodeNum *string `json:"NodeNum,omitempty" name:"NodeNum"`
}

type ClusterSet struct {
	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群key

	ClusterKey *string `json:"ClusterKey,omitempty" name:"ClusterKey"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群告警总数

	ClusterAlarmTotal *int64 `json:"ClusterAlarmTotal,omitempty" name:"ClusterAlarmTotal"`
	// 集群mysql告警总数

	ClusterMysqlAlarmTotal *int64 `json:"ClusterMysqlAlarmTotal,omitempty" name:"ClusterMysqlAlarmTotal"`
	// 集群proxy告警总数

	ClusterProxyAlarmTotal *int64 `json:"ClusterProxyAlarmTotal,omitempty" name:"ClusterProxyAlarmTotal"`
}

type DescribeDeviceSpecificationsRequest struct {
	*tchttp.BaseRequest

	// 集群key

	ClusterKey *string `json:"ClusterKey,omitempty" name:"ClusterKey"`
}

func (r *DescribeDeviceSpecificationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceSpecificationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceSpecificationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 设备规格列表

		DeviceSpecificationSet []*Specinfo `json:"DeviceSpecificationSet,omitempty" name:"DeviceSpecificationSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceSpecificationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeviceSpecificationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCapacityStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 容量统计列表

		CapacityStatisticSet []*CapacityStatisticSet `json:"CapacityStatisticSet,omitempty" name:"CapacityStatisticSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCapacityStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCapacityStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceSet struct {
	// 集群key

	ClusterKey *string `json:"ClusterKey,omitempty" name:"ClusterKey"`
	// Mid

	Mid *string `json:"Mid,omitempty" name:"Mid"`
	// PMid

	PMid *string `json:"PMid,omitempty" name:"PMid"`
	// OSS类型

	OssType *string `json:"OssType,omitempty" name:"OssType"`
	// OSS IDC

	OssIdc *string `json:"OssIdc,omitempty" name:"OssIdc"`
	// OSS可用区

	OssZone *string `json:"OssZone,omitempty" name:"OssZone"`
	// OSS机器

	OssMachine *string `json:"OssMachine,omitempty" name:"OssMachine"`
	// OSS状态

	OssStatus *string `json:"OssStatus,omitempty" name:"OssStatus"`
}

type ProductState struct {
	// 产品名

	Product *string `json:"Product,omitempty" name:"Product"`
	// 产品集群列表

	ClusterList []*uint64 `json:"ClusterList,omitempty" name:"ClusterList"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Clusterhealth struct {
	// 产品名

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 错误信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 健康指标

	Metrics []*Metrics `json:"Metrics,omitempty" name:"Metrics"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群key

	ClusterKey *string `json:"ClusterKey,omitempty" name:"ClusterKey"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Node struct {
	// 内网IP

	PrivateIP *string `json:"PrivateIP,omitempty" name:"PrivateIP"`
	// 公网IP

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 机架

	Rack *string `json:"Rack,omitempty" name:"Rack"`
	// 可用区

	Az *string `json:"Az,omitempty" name:"Az"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 节点存活状态(正常=alive，其他=dead)

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DescribeCapacityStatisticsRequest struct {
	*tchttp.BaseRequest

	// 集群key

	ClusterKey *string `json:"ClusterKey,omitempty" name:"ClusterKey"`
}

func (r *DescribeCapacityStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCapacityStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductHealthStateRequest struct {
	*tchttp.BaseRequest

	// 产品名

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集群名（如果为空返回所有集群信息）

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *QueryProductHealthStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductHealthStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductFailureMigrateRequest struct {
	*tchttp.BaseRequest

	// 产品名

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 产品集群当前故障AZ

	CurrentAz *string `json:"CurrentAz,omitempty" name:"CurrentAz"`
	// 产品集群当前故障Region

	CurrentRegion *string `json:"CurrentRegion,omitempty" name:"CurrentRegion"`
	// 产品集群主节点目标迁移AZ

	TargetAz *string `json:"TargetAz,omitempty" name:"TargetAz"`
	// 产品集群主节点目标迁移Regio

	TargetRegion *string `json:"TargetRegion,omitempty" name:"TargetRegion"`
	// 故障场景

	FailureScenario *string `json:"FailureScenario,omitempty" name:"FailureScenario"`
	// 操作类型

	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
	// 调用方唯一凭据，避免重复调用

	CallerToken *string `json:"CallerToken,omitempty" name:"CallerToken"`
}

func (r *CreateProductFailureMigrateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductFailureMigrateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群设备集合

		DeviceSet []*DeviceSet `json:"DeviceSet,omitempty" name:"DeviceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductStateInfoRequest struct {
	*tchttp.BaseRequest

	// 产品名

	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *QueryProductStateInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductStateInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCreateProductFailureMigrateTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 切换任务UUID

	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
}

func (r *QueryCreateProductFailureMigrateTaskDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCreateProductFailureMigrateTaskDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest

	// 进群key

	ClusterKey *string `json:"ClusterKey,omitempty" name:"ClusterKey"`
}

func (r *DescribeDevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
