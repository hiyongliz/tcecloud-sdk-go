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

package v20200910

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type PodMonitor struct {
	// cpu使用率

	CPURequestRatio *MetricData `json:"CPURequestRatio,omitempty" name:"CPURequestRatio"`
	// 内存使用量

	MemRequestRatio *MetricData `json:"MemRequestRatio,omitempty" name:"MemRequestRatio"`
	// 磁盘使用率

	DiskUsage *MetricData `json:"DiskUsage,omitempty" name:"DiskUsage"`
}

type OsParam struct {
	// 编号

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 系统名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeCustomScriptsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自定义脚本详细列表

		CustomScripts *CustomScripts `json:"CustomScripts,omitempty" name:"CustomScripts"`
		// 自定义脚本的总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomScriptsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomScriptsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeploymentsStatsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 可用的deployment数量, 可用的判断依据是deployment.status.condtions中类型为Available的condition是否为True

		AvailableCount *int64 `json:"AvailableCount,omitempty" name:"AvailableCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeploymentsStatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeploymentsStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchNotificationAlertsRequest struct {
	*tchttp.BaseRequest

	// offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 消息存储id

	MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
	// label matchers

	Query []*Matcher `json:"Query,omitempty" name:"Query"`
}

func (r *SearchNotificationAlertsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchNotificationAlertsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HTTPIngressPath struct {
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 路径类型，完全匹配Exact, 匹配前缀Prefix, 自定义ImplementationSpecific

	PathType *string `json:"PathType,omitempty" name:"PathType"`
	// 后端

	Backend *IngressBackend `json:"Backend,omitempty" name:"Backend"`
}

type CreateStatefulSetRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 有状态组件StatefulSet

	StatefulSet *string `json:"StatefulSet,omitempty" name:"StatefulSet"`
	// 服务列表

	Services []*ServiceRequest `json:"Services,omitempty" name:"Services"`
	// Ingress列表

	Ingresses []*IngressRequest `json:"Ingresses,omitempty" name:"Ingresses"`
}

func (r *CreateStatefulSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateStatefulSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancerListenersRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 指定删除的监听器ID数组

	ListenerIds *string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

func (r *DeleteLoadBalancerListenersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLoadBalancerListenersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监听器后端绑定的机器信息

		Listeners []*ListenerBackend `json:"Listeners,omitempty" name:"Listeners"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTargetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListHistoryRevisionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 历史版本的信息

		HistoryRevisions []*HistoryRevision `json:"HistoryRevisions,omitempty" name:"HistoryRevisions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListHistoryRevisionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListHistoryRevisionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecretRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeSecretRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecretRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceInstanceStatusOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各支撑服务实例状态分布

		ServiceInstanceStatus []*ServiceInstanceStatus `json:"ServiceInstanceStatus,omitempty" name:"ServiceInstanceStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceInstanceStatusOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceInstanceStatusOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOSListRequest struct {
	*tchttp.BaseRequest

	// "Op": "query_os"

	Op *string `json:"Op,omitempty" name:"Op"`
	// "Filters": [{"Name": "Arch", "Values": ["x86"]}]

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeOSListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOSListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirtualMachineAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟机详情

		VirtualMachine *VirtualMachine `json:"VirtualMachine,omitempty" name:"VirtualMachine"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirtualMachineAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirtualMachineAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSilenceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSilenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSilenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID, 如"global"

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段, 支持["TotalComponent", "PodCount"]

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否降序排序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 筛选条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 项目Id

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ListAppInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListContainersRequest struct {
	*tchttp.BaseRequest

	// 所属集群id

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// pod名称

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// pod所属命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *ListContainersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListContainersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMiddlewareOverviewWithAlertsRequest struct {
	*tchttp.BaseRequest

	// 是否展示中间件详情，默认false

	ShowDetail *bool `json:"ShowDetail,omitempty" name:"ShowDetail"`
}

func (r *DescribeMiddlewareOverviewWithAlertsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMiddlewareOverviewWithAlertsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRAIDParam struct {
	// RAID名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type HistoryRevision struct {
	// 版本号

	Revision *int64 `json:"Revision,omitempty" name:"Revision"`
	// 镜像名, 如有多个容器, 取第一个容器的镜像

	Image *string `json:"Image,omitempty" name:"Image"`
	// 镜像Tag, 如有多个容器, 取第一个容器的镜像

	Tag *string `json:"Tag,omitempty" name:"Tag"`
	// 更新时间

	UpdateTimestamp *string `json:"UpdateTimestamp,omitempty" name:"UpdateTimestamp"`
	// 该版本的workload的spec, 是一个YAML的字符串

	WorkloadYAML *string `json:"WorkloadYAML,omitempty" name:"WorkloadYAML"`
}

type DescribeCronJobListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CronJob总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// CronJob列表

		CronJobList *string `json:"CronJobList,omitempty" name:"CronJobList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCronJobListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCronJobListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群ID

		Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
		// 所属组件类型

		OwnerKind *string `json:"OwnerKind,omitempty" name:"OwnerKind"`
		// 所属组件名称

		OwnerName *string `json:"OwnerName,omitempty" name:"OwnerName"`
		// 服务详情

		Service *string `json:"Service,omitempty" name:"Service"`
		// 项目

		Project *string `json:"Project,omitempty" name:"Project"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterInstantMetricDataRequest struct {
	*tchttp.BaseRequest

	// 查询指标名称，多个用逗号分隔

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 查询时间点

	Start *string `json:"Start,omitempty" name:"Start"`
	// 用于过滤node的标签集合，key=value形式，多个用逗号分隔

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 精度，如1m

	Step *string `json:"Step,omitempty" name:"Step"`
}

func (r *GetClusterInstantMetricDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterInstantMetricDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustomScriptRequest struct {
	*tchttp.BaseRequest

	// 自定义脚本的名称

	CustomScriptName *string `json:"CustomScriptName,omitempty" name:"CustomScriptName"`
	// 自定义脚本的内容，base64编码

	CustomScriptContext *string `json:"CustomScriptContext,omitempty" name:"CustomScriptContext"`
}

func (r *CreateCustomScriptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomScriptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIngressClassResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIngressClassResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIngressClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterInstantMetricDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterInstantMetricDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterInstantMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleServerWanIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecycleServerWanIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleServerWanIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOSParam struct {
	// OS名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DiskFieldToEnum struct {
	// HdCapacity

	HdCapacity []*string `json:"HdCapacity,omitempty" name:"HdCapacity"`
	// HdInfType

	HdInfType []*string `json:"HdInfType,omitempty" name:"HdInfType"`
	// HdSlot

	HdSlot []*string `json:"HdSlot,omitempty" name:"HdSlot"`
	// IdcName

	IdcName []*string `json:"IdcName,omitempty" name:"IdcName"`
	// ZoneName

	ZoneName []*string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type ProductComponentStatus struct {
	// 健康的pod数量

	NormalCount *int64 `json:"NormalCount,omitempty" name:"NormalCount"`
	// 不健康的pod数量

	AbnormalCount *int64 `json:"AbnormalCount,omitempty" name:"AbnormalCount"`
	// 组件名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type RecycleServerWanIP struct {
	// 服务器资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器外网IP

	SvrWanIp *string `json:"SvrWanIp,omitempty" name:"SvrWanIp"`
}

type AllocateVMWanIPListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server_vm"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"allocate_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 回收VM外网IP信息

	AllocateVMWanIPSet []*AllocateVMWanIP `json:"AllocateVMWanIPSet,omitempty" name:"AllocateVMWanIPSet"`
}

func (r *AllocateVMWanIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateVMWanIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodesOverviewWithAlertsRequest struct {
	*tchttp.BaseRequest

	// 是否显示node详情，默认false

	ShowDetail *bool `json:"ShowDetail,omitempty" name:"ShowDetail"`
}

func (r *DescribeNodesOverviewWithAlertsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodesOverviewWithAlertsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerRecycleLanIPExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ServerRecycleLanIPInfo

		ServerRecycleLanIPInfo []*ServerRecycleLanIPInfo `json:"ServerRecycleLanIPInfo,omitempty" name:"ServerRecycleLanIPInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ServerRecycleLanIPExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerRecycleLanIPExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Netdevice struct {
	// 网络设备名称

	NetdevName *string `json:"NetdevName,omitempty" name:"NetdevName"`
	// 网络设备类型

	NetdevFun *string `json:"NetdevFun,omitempty" name:"NetdevFun"`
	// 网络设备资产号

	NetdevAssetId *string `json:"NetdevAssetId,omitempty" name:"NetdevAssetId"`
}

type NetworkPolicy struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 组件类型

	ComponentKind *string `json:"ComponentKind,omitempty" name:"ComponentKind"`
	// 组件名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 生效组件的标签

	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`
}

type CreateOrUpdateLogRuleGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrUpdateLogRuleGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateLogRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkPolicyRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 网络策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeNetworkPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListVolumesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合筛选条件的Volume总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Volume列表

		VolumeSet []*Volume `json:"VolumeSet,omitempty" name:"VolumeSet"`
		// StorageClass列表

		SCSet []*string `json:"SCSet,omitempty" name:"SCSet"`
		// 结点名列表

		NodeSet []*string `json:"NodeSet,omitempty" name:"NodeSet"`
		// 租户列表

		TenentSet []*string `json:"TenentSet,omitempty" name:"TenentSet"`
		// 项目列表

		ProjectSet []*string `json:"ProjectSet,omitempty" name:"ProjectSet"`
		// pod名列表

		PodSet []*string `json:"PodSet,omitempty" name:"PodSet"`
		// app名称列表

		AppSet []*string `json:"AppSet,omitempty" name:"AppSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListVolumesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListVolumesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VirtualMachinesStorage struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 读写模式

	ReadWriteMode *string `json:"ReadWriteMode,omitempty" name:"ReadWriteMode"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 策略

	ReclaimPolicy *string `json:"ReclaimPolicy,omitempty" name:"ReclaimPolicy"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 容量

	Capacity *string `json:"Capacity,omitempty" name:"Capacity"`
	// 是否挂载

	Mounted *bool `json:"Mounted,omitempty" name:"Mounted"`
}

type StartRelocateServerInfo struct {
	// 结果详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 结果错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 服务器资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type WorkerClusterUpgradeConfig struct {
	// 节点当前版本

	CurrentVersion *string `json:"CurrentVersion,omitempty" name:"CurrentVersion"`
	// 节点目标版本

	TargetVersion *string `json:"TargetVersion,omitempty" name:"TargetVersion"`
	// 首个节点升级后的观察时间(分钟)

	ObserveTime *int64 `json:"ObserveTime,omitempty" name:"ObserveTime"`
	// 并发节点数百分比上限

	ConcurrentPercentLimit *int64 `json:"ConcurrentPercentLimit,omitempty" name:"ConcurrentPercentLimit"`
	// 节点驱除,0-关,1-开

	ExcludeNodeEnable *int64 `json:"ExcludeNodeEnable,omitempty" name:"ExcludeNodeEnable"`
}

type ModifyLabelModifyInfo struct {
	// 修改后的标签键键

	LabelKey *string `json:"LabelKey,omitempty" name:"LabelKey"`
	// 修改后的标签键键

	LabelValue *string `json:"LabelValue,omitempty" name:"LabelValue"`
}

type GetDefaultGroupByRequest struct {
	*tchttp.BaseRequest
}

func (r *GetDefaultGroupByRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDefaultGroupByRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigInfo struct {
	// meta信息

	Metas []*Meta `json:"Metas,omitempty" name:"Metas"`
	// 日志用途标识

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志路径，支持通配

	Path *string `json:"Path,omitempty" name:"Path"`
	// 当path包含通配符时，该字段用于排除不必要的日志文件

	PathExcluded *string `json:"PathExcluded,omitempty" name:"PathExcluded"`
	// 是否启用多行匹配

	MultipleLine *bool `json:"MultipleLine,omitempty" name:"MultipleLine"`
	// 指定系统已有的parse

	ParserRef *string `json:"ParserRef,omitempty" name:"ParserRef"`
	// 用于手工指定配置，可用配置有json、regex等

	ParserFormat *string `json:"ParserFormat,omitempty" name:"ParserFormat"`
	// time字段名称

	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`
	// time格式

	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`
	// 提取正则

	Regex *string `json:"Regex,omitempty" name:"Regex"`
	// 行首正则

	LineHeaderRegex *string `json:"LineHeaderRegex,omitempty" name:"LineHeaderRegex"`
	// 当ParserFormat为LTSV时需要指定

	Type *string `json:"Type,omitempty" name:"Type"`
	// 用于手工指定配置，主要用于对特定定段再次进行解码

	Decoders []*string `json:"Decoders,omitempty" name:"Decoders"`
	// 配置状态，1 开启 0关闭

	State *uint64 `json:"State,omitempty" name:"State"`
	// 配置id

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 配置创建时间 秒级时间戳

	CreatedAt *uint64 `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 是否提取字段

	FieldExtraction *bool `json:"FieldExtraction,omitempty" name:"FieldExtraction"`
	// 是否使用系统时间

	UseSystemTime *bool `json:"UseSystemTime,omitempty" name:"UseSystemTime"`
	// 配置最近更新时间 秒级时间戳

	UpdatedAt *uint64 `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 日志样例

	Example *string `json:"Example,omitempty" name:"Example"`
	// 字段信息，json  样例:[{"Field":"cost","Type":"int"},{"Field":"log_level","Type":"string"},{"Field":"old_price","Type":"double"}]

	FieldMapping *string `json:"FieldMapping,omitempty" name:"FieldMapping"`
}

type Product struct {
	// 名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 状态

	ProductStatus *string `json:"ProductStatus,omitempty" name:"ProductStatus"`
	// Id

	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type GetIngressYAMLRequest struct {
	*tchttp.BaseRequest

	// 集群ID, 如"global"

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// ingress名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// Ingress namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *GetIngressYAMLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetIngressYAMLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProxyAllDataSourceRequestByNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProxyAllDataSourceRequestByNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProxyAllDataSourceRequestByNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchProductInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchProductInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchProductInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMultiNicDefinitionRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询条件，示例：{"Name":"DefinitionName","Values":[""]}

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeMultiNicDefinitionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMultiNicDefinitionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashboardTagsRequest struct {
	*tchttp.BaseRequest
}

func (r *GetDashboardTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleServerSpecialWanIP struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器IP

	SvrIp *string `json:"SvrIp,omitempty" name:"SvrIp"`
}

type InitMachineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 脚本运行输出

		Outputs []*RunScriptOutput `json:"Outputs,omitempty" name:"Outputs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitMachineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListServiceBindingsRequest struct {
	*tchttp.BaseRequest

	// 集群ID列表

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 搜索参数列表

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListServiceBindingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListServiceBindingsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnStarredDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 取消收藏的dashboard id

		DashId *string `json:"DashId,omitempty" name:"DashId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnStarredDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnStarredDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaMetricUpdateCommand struct {
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 说明

	Help *string `json:"Help,omitempty" name:"Help"`
}

type DeleteMultiNicDefinitionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除网卡信息出参

		MultiNicSet *DeleteMultiNicSet `json:"MultiNicSet,omitempty" name:"MultiNicSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMultiNicDefinitionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMultiNicDefinitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkPolicyRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 名称与命名空间列表

	NamespacedNames []*NamespacedName `json:"NamespacedNames,omitempty" name:"NamespacedNames"`
}

func (r *DeleteNetworkPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorResourceRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeMonitorResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Selector struct {
	// Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// value列表

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 描述key/values关系. 可选值: [In, NotIn, Exists, DoesNotExist, Gt, Lt]

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type ListConfigMapsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 符合条件的ConfigMap列表

		ConfigMapSet []*ConfigMap `json:"ConfigMapSet,omitempty" name:"ConfigMapSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListConfigMapsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListConfigMapsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDashboardPermissionsByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDashboardPermissionsByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDashboardPermissionsByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Meta struct {
	// meta key

	MetaKey *string `json:"MetaKey,omitempty" name:"MetaKey"`
	// meta value

	MetaValue *string `json:"MetaValue,omitempty" name:"MetaValue"`
	// meta name

	MetaName *string `json:"MetaName,omitempty" name:"MetaName"`
	// 子 meta数组

	SubMetas []*SubMeta `json:"SubMetas,omitempty" name:"SubMetas"`
	// config count

	ConfigCount *uint64 `json:"ConfigCount,omitempty" name:"ConfigCount"`
}

type AddVMListRequest struct {
	*tchttp.BaseRequest

	// server_vm

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// "SystemId":"sid_test"

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// "SystemKey":"sid_test"

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// add

	Op *string `json:"Op,omitempty" name:"Op"`
	// 增加VM信息

	AddVMSet []*VM `json:"AddVMSet,omitempty" name:"AddVMSet"`
}

func (r *AddVMListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVMListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDashboardByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAlertTrendRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartsAt *string `json:"StartsAt,omitempty" name:"StartsAt"`
	// 结束时间

	EndsAt *string `json:"EndsAt,omitempty" name:"EndsAt"`
}

func (r *GetAlertTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetVirtualMachineStatusRequest struct {
	*tchttp.BaseRequest

	// restart, start, stop, pause, restore

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 项目

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// vmid

	VmId *string `json:"VmId,omitempty" name:"VmId"`
}

func (r *SetVirtualMachineStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetVirtualMachineStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigServerDefData struct {
	// SvrCurrentStatus

	SvrCurrentStatus *string `json:"SvrCurrentStatus,omitempty" name:"SvrCurrentStatus"`
	// SvrIdcName

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// SvrSn

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// TaskId

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type MetaResourceTypeQueryParam struct {
	// 租户ID，对应AppId，通常不需要传，默认取当前登录用户的AppId。

	Organization *string `json:"Organization,omitempty" name:"Organization"`
	// 要查询的对象在服务树上路径，格式为{a="b",c="d"}

	ResourcePath *string `json:"ResourcePath,omitempty" name:"ResourcePath"`
	// 查询类型，可选值为on/under，分别代表当前节点上的数据和当前节点下的数据两种含义。

	ListType *string `json:"ListType,omitempty" name:"ListType"`
	// 过滤类型，可选类型为metric和event，可不填，代表取所有类型

	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`
}

type SearchResourceCommand struct {
	// 资源归属租户

	Organization *string `json:"Organization,omitempty" name:"Organization"`
	// 服务树节点选择器

	Selector *string `json:"Selector,omitempty" name:"Selector"`
}

type VMAllocLanIP struct {
	// 结果详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 服务器资产Id

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器序列号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
}

type DeleteDashboardByUidRequest struct {
	*tchttp.BaseRequest

	// 仪表盘唯一ID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
}

func (r *DeleteDashboardByUidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardByUidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIgniterNodeExRequest struct {
	*tchttp.BaseRequest

	// DescribeIgniterNodeEx入参

	Data *DescribeIgniterNodeExData `json:"Data,omitempty" name:"Data"`
}

func (r *DescribeIgniterNodeExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIgniterNodeExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleVMWanIPListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server_vm"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"recycle_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 回收VM外网IP信息

	RecycleVMWanIPSet []*RecycleVMWanIP `json:"RecycleVMWanIPSet,omitempty" name:"RecycleVMWanIPSet"`
}

func (r *RecycleVMWanIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleVMWanIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOsDeployDebugInfoExRequest struct {
	*tchttp.BaseRequest

	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// os部署调试信息

	OsDeployDebugInfo []*OsDeployDebugInfo `json:"OsDeployDebugInfo,omitempty" name:"OsDeployDebugInfo"`
}

func (r *DescribeOsDeployDebugInfoExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOsDeployDebugInfoExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageParams struct {
	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 自增ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type ListCreateVolumeParametersRequest struct {
	*tchttp.BaseRequest

	// kubernetes集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 待创建Volume的StorageClass名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ListCreateVolumeParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCreateVolumeParametersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopologyStateRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 状态，开启: active, 关闭: inactive

	State *string `json:"State,omitempty" name:"State"`
}

func (r *ModifyTopologyStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTopologyStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TreeNodeComponent struct {
	// 结点类型, Container/Support/Product, 分别对应容器类型, 支撑类型, 生成类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 如果类型为Container, 这个字段返回容器类型的组件列表

	ContainerComponets []*TreeNodePod `json:"ContainerComponets,omitempty" name:"ContainerComponets"`
	// 如果类型为Support, 这个字段返回支撑类型的组件列表

	SupportComponents []*TreeNodeMiddleware `json:"SupportComponents,omitempty" name:"SupportComponents"`
	// 如果类型为Product, 这个字段返回生产类型的组件列表

	ProductComponents []*string `json:"ProductComponents,omitempty" name:"ProductComponents"`
	// 组件名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DeletePersistentVolumeClaimResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePersistentVolumeClaimResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePersistentVolumeClaimResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRoutesRequest struct {
	*tchttp.BaseRequest

	// 批量删除，Ids  1,2,3

	Params *BatchDeleteIds `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAlertResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAlertResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDaemonSetsRequest struct {
	*tchttp.BaseRequest

	// 集群ID列表

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 数量限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 是否倒序排序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *ListDaemonSetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDaemonSetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StatusCount struct {
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 正常数

	NormalCount *int64 `json:"NormalCount,omitempty" name:"NormalCount"`
	// 异常数

	AbnormalCount *int64 `json:"AbnormalCount,omitempty" name:"AbnormalCount"`
}

type GetNodeHistoryMetricDataRequest struct {
	*tchttp.BaseRequest

	// 指标名，多个以逗号分隔

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 起始时间

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束时间

	End *string `json:"End,omitempty" name:"End"`
	// 精度，如10s, 1m

	Step *string `json:"Step,omitempty" name:"Step"`
	// 查询标签，用于过滤特定的node

	Labels *string `json:"Labels,omitempty" name:"Labels"`
}

func (r *GetNodeHistoryMetricDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNodeHistoryMetricDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRuleGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRuleGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IngressPath struct {
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 路径类型，完全匹配Exact, 匹配前缀Prefix, 自定义ImplementationSpecific

	PathType *string `json:"PathType,omitempty" name:"PathType"`
	// 服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务端口

	ServicePort *string `json:"ServicePort,omitempty" name:"ServicePort"`
}

type DescribeServiceYAMLRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 服务名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeServiceYAMLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceYAMLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListImageBuildingTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的总量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 构建任务数据集

		ImageBuildingTaskSet []*ImageBuildingTask `json:"ImageBuildingTaskSet,omitempty" name:"ImageBuildingTaskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListImageBuildingTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListImageBuildingTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMonitorMetricsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标总个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 符合条件的指标列表

		MonitorMetricSet []*MetricDefine `json:"MonitorMetricSet,omitempty" name:"MonitorMetricSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMonitorMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMonitorMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleVMVirtualIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecycleVMVirtualIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleVMVirtualIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSilenceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSilenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSilenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelocateCancelServer struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// IDC名称

	SvrRelocateIdcName *string `json:"SvrRelocateIdcName,omitempty" name:"SvrRelocateIdcName"`
	// 机架名称

	SvrRelocateRackName *string `json:"SvrRelocateRackName,omitempty" name:"SvrRelocateRackName"`
	// 机架位置

	SvrRelocatePosId *string `json:"SvrRelocatePosId,omitempty" name:"SvrRelocatePosId"`
}

type AllocateSvrNicIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务器的多网卡ip分配入参出参

		AllocateSvrNicIPSet []*AllocateSvrNicIPSet `json:"AllocateSvrNicIPSet,omitempty" name:"AllocateSvrNicIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateSvrNicIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateSvrNicIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNotificationMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNotificationMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNotificationMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeHistoryDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUpgradeHistoryDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeHistoryDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceType struct {
	// 类型名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 类型名称别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 父类型

	Parent *string `json:"Parent,omitempty" name:"Parent"`
}

type UnStarredDashboardRequest struct {
	*tchttp.BaseRequest

	// dashboard id

	DashId *string `json:"DashId,omitempty" name:"DashId"`
}

func (r *UnStarredDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnStarredDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateRetention struct {
	// 主键，根据Id是否大于0来区分create or update

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 名称，唯一

	Name *string `json:"Name,omitempty" name:"Name"`
	// 运行间隔，默认1d

	Interval *string `json:"Interval,omitempty" name:"Interval"`
	// 数据保留时长

	Retention *string `json:"Retention,omitempty" name:"Retention"`
	// 辅助信息，比如{"max":"365d", "min":"90d"}

	Annotation []*LabelPair `json:"Annotation,omitempty" name:"Annotation"`
	// 乐观锁

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 更新时间

	Updated *string `json:"Updated,omitempty" name:"Updated"`
}

type PromQLMetricOperation struct {
	// +-x/

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 运算值

	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type RAIDParams struct {
	// 编号

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// RAID名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type Overview struct {
	// 镜像类型信息

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 物理服务器状态

	SvrCurrentStatus *SvrCurrentStatus `json:"SvrCurrentStatus,omitempty" name:"SvrCurrentStatus"`
	// 物理服务器信息

	SvrType *string `json:"SvrType,omitempty" name:"SvrType"`
	// 黑石镜像信息

	SocOsType *string `json:"SocOsType,omitempty" name:"SocOsType"`
	// 黑石服务器信息

	SocSvrType *string `json:"SocSvrType,omitempty" name:"SocSvrType"`
}

type RecycleVMWanIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟机回收IP信息

		VmRecyWanIPSet []*VMRecyWanIP `json:"VmRecyWanIPSet,omitempty" name:"VmRecyWanIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecycleVMWanIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleVMWanIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUpgradeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFiringRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFiringRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFiringRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateVMLanIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟机分配内网IP信息

		VMAllocMLanIPSet []*VMAllocLanIP `json:"VMAllocMLanIPSet,omitempty" name:"VMAllocMLanIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateVMLanIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateVMLanIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Matcher struct {
	// 0,1,2,3(=, !=, !~, =~)

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DevopsJob struct {
	// 节点ip

	Node *string `json:"Node,omitempty" name:"Node"`
	// 集群名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 自愈模板名称

	RemediationName *string `json:"RemediationName,omitempty" name:"RemediationName"`
	// 自愈触发条件

	OperationCondition *string `json:"OperationCondition,omitempty" name:"OperationCondition"`
	// 自愈动作

	OperationAction *string `json:"OperationAction,omitempty" name:"OperationAction"`
	// 执行结果

	Result *string `json:"Result,omitempty" name:"Result"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 完成时间

	CompletionTime *string `json:"CompletionTime,omitempty" name:"CompletionTime"`
}

type HarddiskInfo struct {
	// 容量

	HdCapacity *string `json:"HdCapacity,omitempty" name:"HdCapacity"`
	// 固件版本

	HdFw *string `json:"HdFw,omitempty" name:"HdFw"`
	// 接口类型

	HdInfType *string `json:"HdInfType,omitempty" name:"HdInfType"`
	// 介质

	HdMedium *string `json:"HdMedium,omitempty" name:"HdMedium"`
	// HdRp

	HdRp *string `json:"HdRp,omitempty" name:"HdRp"`
	// 槽位

	HdSlot *string `json:"HdSlot,omitempty" name:"HdSlot"`
	// SN编号

	HdSn *string `json:"HdSn,omitempty" name:"HdSn"`
	// 类型

	HdType *string `json:"HdType,omitempty" name:"HdType"`
	// 厂家

	HdVendor *string `json:"HdVendor,omitempty" name:"HdVendor"`
	// 编号

	Id *string `json:"Id,omitempty" name:"Id"`
	// 关联服务器

	RelatedSvr *string `json:"RelatedSvr,omitempty" name:"RelatedSvr"`
	// 地域

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 可用区

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 数据中心名称

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// 数据中心ID

	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
	// 服务器资产编号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type ListVolumesNodeViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总结点数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 结点数据

		NodeSet []*VolumeNodeView `json:"NodeSet,omitempty" name:"NodeSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListVolumesNodeViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListVolumesNodeViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyStatefulSetRequest struct {
	*tchttp.BaseRequest

	// 集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 待修改的StatefulSet

	StatefulSet *string `json:"StatefulSet,omitempty" name:"StatefulSet"`
	// 描述, 不传表示不修改

	Description *string `json:"Description,omitempty" name:"Description"`
	// 标签, 不传表示不修改, 传空数组表示清空标签

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// 副本数量

	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *ModifyStatefulSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStatefulSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateSvrNicIPSet struct {
	// 返回结果值，0成功， 1失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 返回结果详细说明

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 分配的ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 网卡名称

	NicName *string `json:"NicName,omitempty" name:"NicName"`
	// 是否是默认；

	Default *bool `json:"Default,omitempty" name:"Default"`
}

type DescribeServersRequest struct {
	*tchttp.BaseRequest

	// "scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// "system_id":"sid_test"

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// "system_key":"sid_test"

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"query"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// //(根据需要填充需要返回的字段名，不同字段名使用","隔开)

	ResultItem *string `json:"ResultItem,omitempty" name:"ResultItem"`
	// //(根据需要填充需要模糊查询的字段名，不同字段名使用","隔开，对应的字段名一定要在Filters里面体现)

	RoughItem *string `json:"RoughItem,omitempty" name:"RoughItem"`
	// 过滤条件："Filters":[{"Name":"SvrLanIp","Value":"10.21.12.139"},{"Name":"SvrAssetId","Value":""},{"Name":"SvrIdcName","Value":""},{"Name":"SvrLogicArea","Value":""},{"Name":"SvrCurrentStatus","Value":""},{"Name":"SvrBussiness","Value":""}]}

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerAllocateLanIPExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 请求动作

	Op *string `json:"Op,omitempty" name:"Op"`
	// 物理服务器分配内网IP

	ServersAllocateLanIp []*ServersAllocateLanIp `json:"ServersAllocateLanIp,omitempty" name:"ServersAllocateLanIp"`
}

func (r *ServerAllocateLanIPExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerAllocateLanIPExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRetentionsRequest struct {
	*tchttp.BaseRequest

	// 搜索Retention

	Command *SearchRetention `json:"Command,omitempty" name:"Command"`
}

func (r *SearchRetentionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRetentionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomScriptsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCustomScriptsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomScriptsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeploySubtaskStepsExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeploySubtaskStepsExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeploySubtaskStepsExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePodsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异常pod列表

		PodSet []*DescribePod `json:"PodSet,omitempty" name:"PodSet"`
		// pod统计

		PodStatusMap *PodStatus `json:"PodStatusMap,omitempty" name:"PodStatusMap"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePodsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchNotificationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchNotificationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchNotificationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerJoinToBanksRequest struct {
	*tchttp.BaseRequest

	// 密码类型

	PasswordType *int64 `json:"PasswordType,omitempty" name:"PasswordType"`
	// 入库的服务器信息

	Servers []*JoinToBanksServers `json:"Servers,omitempty" name:"Servers"`
}

func (r *ServerJoinToBanksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerJoinToBanksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PromQLMetrics struct {
	// single|multiple

	Type *string `json:"Type,omitempty" name:"Type"`
	// 指标列表

	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`
	// 操作符列表

	Operator []*string `json:"Operator,omitempty" name:"Operator"`
}

type AllocateServerVirtualIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateServerVirtualIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateServerVirtualIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRAIDListRequest struct {
	*tchttp.BaseRequest

	// "Op": "query_raid"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 模糊匹配

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制返回个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRAIDListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRAIDListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecretResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群ID

		Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
		// 密钥数据

		Secret *SecretData `json:"Secret,omitempty" name:"Secret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecretResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOutbandInfoExRequest struct {
	*tchttp.BaseRequest

	// 可用区域ID

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// idcid

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// 请求动作

	Command *string `json:"Command,omitempty" name:"Command"`
	// 用户

	User *string `json:"User,omitempty" name:"User"`
	// 数据

	Data []*string `json:"Data,omitempty" name:"Data"`
}

func (r *DescribeOutbandInfoExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOutbandInfoExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EvictNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EvictNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EvictNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNotificationMsgRequest struct {
	*tchttp.BaseRequest

	// 消息id

	MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
}

func (r *GetNotificationMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNotificationMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogRuleGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchLogRuleGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogRuleGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomScriptSetsRequest struct {
	*tchttp.BaseRequest

	// 需要删除的自定义脚本集

	CustomScriptSetName []*string `json:"CustomScriptSetName,omitempty" name:"CustomScriptSetName"`
}

func (r *DeleteCustomScriptSetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomScriptSetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListClusterRequest struct {
	*tchttp.BaseRequest

	// 列表行数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 过滤参数列表

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMetaResourceTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 资源类型列表

		Data []*MetaResourceType `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMetaResourceTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMetaResourceTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReinstallOsExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ReinstallOsSet

		ReinstallOsSet []*ReinstallOsSet `json:"ReinstallOsSet,omitempty" name:"ReinstallOsSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReinstallOsExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReinstallOsExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatefulSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// StatefulSet的JSON字符串, 需要自行parse成JSON

		StatefulSet *string `json:"StatefulSet,omitempty" name:"StatefulSet"`
		// Service列表, 每个元素是一个JSON字符串, 需要自行parse

		Services []*string `json:"Services,omitempty" name:"Services"`
		// Ingress列表, 每个元素是一个JSON字符串, 需要自行parse

		Ingresses []*string `json:"Ingresses,omitempty" name:"Ingresses"`
		// StatefulSet的状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 上次更新时间

		LastUpdateTimestamp *string `json:"LastUpdateTimestamp,omitempty" name:"LastUpdateTimestamp"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStatefulSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStatefulSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNotificationRuleRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 通知类型，如npd类型

	Exporter *string `json:"Exporter,omitempty" name:"Exporter"`
}

func (r *GetNotificationRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNotificationRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductStatus struct {
	// 云产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 正常组件个数

	NormalCount *float64 `json:"NormalCount,omitempty" name:"NormalCount"`
	// 异常组件个数

	AbnormalCount *float64 `json:"AbnormalCount,omitempty" name:"AbnormalCount"`
	// 总组件个数

	TotalCount *float64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type StatefulSetUpdateStrategy struct {
	// 策略类型， RollingUpdate, OnDelete

	Type *string `json:"Type,omitempty" name:"Type"`
	// 滚动升级参数

	RollingUpdate *RollingUpdateStatefulSetStrategy `json:"RollingUpdate,omitempty" name:"RollingUpdate"`
}

type SearchMetaValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// metavalue

		MetaNames []*string `json:"MetaNames,omitempty" name:"MetaNames"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchMetaValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchMetaValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOSParam struct {
	// 序号

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type LogField struct {
	// 字段类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 字段名

	Field *string `json:"Field,omitempty" name:"Field"`
}

type GetSeriesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSeriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSeriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateResourceObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateResourceObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateResourceObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalyseLogsRequest struct {
	*tchttp.BaseRequest

	// LogQL查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 返回条数限制，默认10条

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分桶参数

	Buckets []*AnalyzeBucket `json:"Buckets,omitempty" name:"Buckets"`
	// 指标参数

	Metrics []*AnalyzeMetric `json:"Metrics,omitempty" name:"Metrics"`
	// 过滤参数

	Filters []*AnalyzeFilter `json:"Filters,omitempty" name:"Filters"`
	// 排序参数

	Sorts []*AnalyzeSort `json:"Sorts,omitempty" name:"Sorts"`
	// 起始时间 纳秒时间戳 或 RFC3339Nano

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 截止时间 纳秒时间戳 或 RFC3339Nano

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *AnalyseLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AnalyseLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SvrNicSet struct {
	// 服务器SN

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 自定义多网卡名称

	DefinitionName *string `json:"DefinitionName,omitempty" name:"DefinitionName"`
	// 已绑定多网卡的信息；示例:"Allocation":[{"DefaultGw":false,"Ip":"","NicName":"bond2","NicType":"lan"}]

	Allocation []*DescribeAllocation `json:"Allocation,omitempty" name:"Allocation"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ConfigServerDefSetLan struct {
	// 内网IP

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 掩码

	Netmask *string `json:"Netmask,omitempty" name:"Netmask"`
	// 网关

	Gateway *string `json:"Gateway,omitempty" name:"Gateway"`
}

type CustomScriptSets struct {
	// 自定义脚本集名

	CustomScriptSetName *string `json:"CustomScriptSetName,omitempty" name:"CustomScriptSetName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 加入自定义脚本集的脚本

	CustomScriptName []*string `json:"CustomScriptName,omitempty" name:"CustomScriptName"`
	// 脚本类型(0:初始化脚本,1:后置脚本)

	CustomScriptSetType *string `json:"CustomScriptSetType,omitempty" name:"CustomScriptSetType"`
}

type QuotaUsage struct {
	// 总量

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 已分配

	Allocated *int64 `json:"Allocated,omitempty" name:"Allocated"`
	// 使用率

	Usage *float64 `json:"Usage,omitempty" name:"Usage"`
}

type CreateDashFolderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建仪表盘目录结果内容

		Data *DashFolderContent `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDashFolderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDashFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVirtualMachineRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// json文本

	VmString *string `json:"VmString,omitempty" name:"VmString"`
	// 密码

	Pass *string `json:"Pass,omitempty" name:"Pass"`
}

func (r *CreateVirtualMachineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVirtualMachineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirtualMachineStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CPU使用

		CPU *QuotaUsage `json:"CPU,omitempty" name:"CPU"`
		// 内存使用

		Memory *QuotaUsage `json:"Memory,omitempty" name:"Memory"`
		// 磁盘使用

		Disk *QuotaUsage `json:"Disk,omitempty" name:"Disk"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirtualMachineStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirtualMachineStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleServerVirtualIPRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 虚拟内网："Op":"recycle_vitrual_lan_ip"     虚拟外网： "Op":"recycle_vitrual_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 回收物理服务器内网外网信息集合

	RecycleServerVirtualIPSet []*VirtualLanOrWanIP `json:"RecycleServerVirtualIPSet,omitempty" name:"RecycleServerVirtualIPSet"`
}

func (r *RecycleServerVirtualIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleServerVirtualIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PhysvrsResultInfo struct {
	// RegionName

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// SvrAdminSwitchAssetid

	SvrAdminSwitchAssetid *string `json:"SvrAdminSwitchAssetid,omitempty" name:"SvrAdminSwitchAssetid"`
	// SvrAgentVersion

	SvrAgentVersion *string `json:"SvrAgentVersion,omitempty" name:"SvrAgentVersion"`
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// SvrBelongingProducts

	SvrBelongingProducts *string `json:"SvrBelongingProducts,omitempty" name:"SvrBelongingProducts"`
	// SvrBussiness

	SvrBussiness *string `json:"SvrBussiness,omitempty" name:"SvrBussiness"`
	// SvrBussinessId

	SvrBussinessId *string `json:"SvrBussinessId,omitempty" name:"SvrBussinessId"`
	// SvrCurrentStatus

	SvrCurrentStatus *string `json:"SvrCurrentStatus,omitempty" name:"SvrCurrentStatus"`
	// SvrDeviceDescript

	SvrDeviceDescript *string `json:"SvrDeviceDescript,omitempty" name:"SvrDeviceDescript"`
	// SvrDeviceHeight

	SvrDeviceHeight *string `json:"SvrDeviceHeight,omitempty" name:"SvrDeviceHeight"`
	// SvrDeviceName

	SvrDeviceName *string `json:"SvrDeviceName,omitempty" name:"SvrDeviceName"`
	// SvrDeviceType

	SvrDeviceType *string `json:"SvrDeviceType,omitempty" name:"SvrDeviceType"`
	// SvrFirstUseTime

	SvrFirstUseTime *string `json:"SvrFirstUseTime,omitempty" name:"SvrFirstUseTime"`
	// SvrId

	SvrId *string `json:"SvrId,omitempty" name:"SvrId"`
	// SvrIdcId

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// SvrIdcName

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// SvrInsertTime

	SvrInsertTime *string `json:"SvrInsertTime,omitempty" name:"SvrInsertTime"`
	// SvrLanIp

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
	// SvrLanSwitchAssetid

	SvrLanSwitchAssetid *string `json:"SvrLanSwitchAssetid,omitempty" name:"SvrLanSwitchAssetid"`
	// SvrLogicArea

	SvrLogicArea *string `json:"SvrLogicArea,omitempty" name:"SvrLogicArea"`
	// SvrOsName

	SvrOsName *string `json:"SvrOsName,omitempty" name:"SvrOsName"`
	// SvrPosId

	SvrPosId *string `json:"SvrPosId,omitempty" name:"SvrPosId"`
	// SvrProducer

	SvrProducer *string `json:"SvrProducer,omitempty" name:"SvrProducer"`
	// SvrRackName

	SvrRackName *string `json:"SvrRackName,omitempty" name:"SvrRackName"`
	// SvrRaidType

	SvrRaidType *string `json:"SvrRaidType,omitempty" name:"SvrRaidType"`
	// SvrSn

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// SvrType

	SvrType *string `json:"SvrType,omitempty" name:"SvrType"`
	// SvrVmIndex

	SvrVmIndex *string `json:"SvrVmIndex,omitempty" name:"SvrVmIndex"`
	// SvrWanEip

	SvrWanEip *string `json:"SvrWanEip,omitempty" name:"SvrWanEip"`
	// SvrWanIp

	SvrWanIp *string `json:"SvrWanIp,omitempty" name:"SvrWanIp"`
	// SvrWanSwitchAssetid

	SvrWanSwitchAssetid *string `json:"SvrWanSwitchAssetid,omitempty" name:"SvrWanSwitchAssetid"`
	// ZoneName

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// TaskId

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// SvrLanGateway

	SvrLanGateway *string `json:"SvrLanGateway,omitempty" name:"SvrLanGateway"`
	// SvrLanMask

	SvrLanMask *string `json:"SvrLanMask,omitempty" name:"SvrLanMask"`
	// SvrMultiNicDefinitionName

	SvrMultiNicDefinitionName *string `json:"SvrMultiNicDefinitionName,omitempty" name:"SvrMultiNicDefinitionName"`
	// SvrVirtualLanIp

	SvrVirtualLanIp *string `json:"SvrVirtualLanIp,omitempty" name:"SvrVirtualLanIp"`
	// SvrArch

	SvrArch *string `json:"SvrArch,omitempty" name:"SvrArch"`
	// SvrLabels

	SvrLabels *string `json:"SvrLabels,omitempty" name:"SvrLabels"`
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageBuildingTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务

		ImageBuildingTask *ImageBuildingTask `json:"ImageBuildingTask,omitempty" name:"ImageBuildingTask"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageBuildingTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageBuildingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductRequest struct {
	*tchttp.BaseRequest

	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品uuid

	ProductUuid *string `json:"ProductUuid,omitempty" name:"ProductUuid"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Normal、Abnormal，状态列表

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 集群id列表

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

func (r *DescribeProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddImageParams struct {
	// 镜像路径

	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
	// 镜像名字

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 操作系统类型

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 系统架构

	SystemArch *string `json:"SystemArch,omitempty" name:"SystemArch"`
	// 系统平台

	SystemPlatform *string `json:"SystemPlatform,omitempty" name:"SystemPlatform"`
	// 系统版本

	SystemVersions *string `json:"SystemVersions,omitempty" name:"SystemVersions"`
	// 镜像描述

	ImageDescribe *string `json:"ImageDescribe,omitempty" name:"ImageDescribe"`
	// ZoneName

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 镜像格式(iso/sqfs)

	ImageFormat *string `json:"ImageFormat,omitempty" name:"ImageFormat"`
	// 镜像使用服务器设备类型(0:通用服务器;1:黑石服务器)

	AvailableModel *int64 `json:"AvailableModel,omitempty" name:"AvailableModel"`
	// IdcName

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
}

type CreateJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateDashboardRequest struct {
	*tchttp.BaseRequest

	// 创建或更新仪表盘内容，更新操作时dashboard字段里面需要包括仪表盘id，仪表盘版本version

	Command *CreateDashboardCommand `json:"Command,omitempty" name:"Command"`
}

func (r *CreateOrUpdateDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTopologyRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 过滤的命名空间列表

	Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces"`
	// 时间

	Time *string `json:"Time,omitempty" name:"Time"`
}

func (r *DescribeServiceTopologyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceTopologyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerAllocateLanIPInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type ModifyLabelsRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID，system_id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key，system_key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"modify"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 修改服务器参数

	LabelServerSet []*ModifyLabel `json:"LabelServerSet,omitempty" name:"LabelServerSet"`
}

func (r *ModifyLabelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLabelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNodeAttributesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyNodeAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNodeAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Application struct {
	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 应用名称

	AppName *string `json:"AppName,omitempty" name:"AppName"`
	// 应用健康状况， normal或者abNormal

	Status *string `json:"Status,omitempty" name:"Status"`
	// 组件总数

	ComTotalCount *int64 `json:"ComTotalCount,omitempty" name:"ComTotalCount"`
	// 异常组件数

	ComAbnormalCount *int64 `json:"ComAbnormalCount,omitempty" name:"ComAbnormalCount"`
}

type Arguments struct {
	// 重试次数

	RetryTimes *int64 `json:"RetryTimes,omitempty" name:"RetryTimes"`
	// 并发策略, Concurrent/Serial

	ConcurrencyPolicy *string `json:"ConcurrencyPolicy,omitempty" name:"ConcurrencyPolicy"`
	// 自定义参数

	CustomParameters []*CustomParameter `json:"CustomParameters,omitempty" name:"CustomParameters"`
}

type DeleteSilencesRequest struct {
	*tchttp.BaseRequest

	// 批量删除ids，"Ids":"1,2,3"

	Params *BatchDeleteIds `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteSilencesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSilencesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetWeightRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 要修改权重的后端服务列表

	Targets []*BackendTarget `json:"Targets,omitempty" name:"Targets"`
	// 后端服务新的转发权重，取值范围：0~100

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

func (r *ModifyTargetWeightRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTargetWeightRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQueryRecordsRequest struct {
	*tchttp.BaseRequest

	// 用户id

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 分页偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回记录条数，默认值：10，最大值：100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeQueryRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQueryRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHealthStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *GetHealthStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHealthStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRelatedAlertNamesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Result []*EventRelatedAlertNames `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRelatedAlertNamesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRelatedAlertNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListNodeRemediationTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自愈场景列表

		Items []*DevOpsScenes `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListNodeRemediationTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListNodeRemediationTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateRuleGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrUpdateRuleGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 维度，cluster,node,pod

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
	// labels,若是cluster维度则填name，若node维度则填ip，若pod维度则填podName

	Labels []*LabelPair `json:"Labels,omitempty" name:"Labels"`
}

func (r *DescribeMonitorTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleServerWanIPListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统Id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"recycle_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 回收服务器外网IP

	RecycleServerWanIPSet []*RecycleServerWanIP `json:"RecycleServerWanIPSet,omitempty" name:"RecycleServerWanIPSet"`
}

func (r *RecycleServerWanIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleServerWanIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SortInfo struct {
	// 待排序字段

	SortField *string `json:"SortField,omitempty" name:"SortField"`
	// 顺序，asc：正序，desc：倒序

	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeServerRelatedNetdevicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务器关联网络设备

		Netdeivces []*Netdevice `json:"Netdeivces,omitempty" name:"Netdeivces"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServerRelatedNetdevicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServerRelatedNetdevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 服务列表

		ServiceSet []*Service `json:"ServiceSet,omitempty" name:"ServiceSet"`
		// 筛选参数列表

		FilterEntity *ServiceFilterEntity `json:"FilterEntity,omitempty" name:"FilterEntity"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJobRequest struct {
	*tchttp.BaseRequest

	// 集群id

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 项目ID

	Project *string `json:"Project,omitempty" name:"Project"`
}

func (r *DescribeJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSilence struct {
	// 屏蔽单名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 级别

	Level *string `json:"Level,omitempty" name:"Level"`
	// 屏蔽开始时间

	StartsAt *string `json:"StartsAt,omitempty" name:"StartsAt"`
	// 屏蔽结束时间

	EndsAt *string `json:"EndsAt,omitempty" name:"EndsAt"`
	// 产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 对象

	Objects []*LabelPair `json:"Objects,omitempty" name:"Objects"`
	// 策略和规则信息

	Labels []*SilenceGroup `json:"Labels,omitempty" name:"Labels"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 乐观锁版本

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 自定义变更人员，默认是当前用户，多个用户间逗号分隔

	Updater *string `json:"Updater,omitempty" name:"Updater"`
}

type CreateCronJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCronJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCronJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListVirtualMachinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 虚拟机列表

		Data []*VirtualMachine `json:"Data,omitempty" name:"Data"`
		// 统计信息

		FilterEntity *VirtualMachineFilterEntity `json:"FilterEntity,omitempty" name:"FilterEntity"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListVirtualMachinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListVirtualMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCronJobsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCronJobsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCronJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenLineHeaderRegexRequest struct {
	*tchttp.BaseRequest

	// 日志样例

	Example *string `json:"Example,omitempty" name:"Example"`
}

func (r *GenLineHeaderRegexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenLineHeaderRegexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMetricMatrixByDimensionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMetricMatrixByDimensionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMetricMatrixByDimensionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 项目id

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Normal、Abnormal,状态列表

	AppStatus []*string `json:"AppStatus,omitempty" name:"AppStatus"`
}

func (r *DescribeProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRuleGroupTempleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRuleGroupTempleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRuleGroupTempleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRuleGroupTemplesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchRuleGroupTemplesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRuleGroupTemplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HTTPIngressRuleValue struct {
	// 路径列表

	Paths *HTTPIngressPath `json:"Paths,omitempty" name:"Paths"`
}

type DeleteImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// "ImageSet": [{"Detail": "ok","Result": 0,"ImageId": "ins-123456","ImageName": "test_os"}

		ImageSet []*string `json:"ImageSet,omitempty" name:"ImageSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePodResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePodResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OutbandInfo struct {
	// DhcpIp

	DhcpIp *string `json:"DhcpIp,omitempty" name:"DhcpIp"`
	// Password

	Password *string `json:"Password,omitempty" name:"Password"`
	// Sn

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// User

	User *string `json:"User,omitempty" name:"User"`
}

type QueryHistoryResult struct {
	// datas

	Data []*MetricData `json:"Data,omitempty" name:"Data"`
}

type SearchAlertData struct {
	// u544au8b66u603bu6570

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// u544au8b66u8be6u7ec6u4fe1u606f

	Alerts []*Alert `json:"Alerts,omitempty" name:"Alerts"`
}

type ServerRecycleLanIPInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// SvrLanIp

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
}

type ListNetworkPoliciesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 搜索参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListNetworkPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListNetworkPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelocateServerFinishRequest struct {
	*tchttp.BaseRequest

	// Scheme

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"relocate_finish"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 搬迁完成信息

	RelocateFinishServers []*RelocateFinishServer `json:"RelocateFinishServers,omitempty" name:"RelocateFinishServers"`
}

func (r *RelocateServerFinishRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelocateServerFinishRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaMetricQueryParam struct {
	// 租户ID，对应AppId，通常不需要传，默认取当前登录用户的AppId

	Organization *string `json:"Organization,omitempty" name:"Organization"`
}

type MetaResourceTypeCreateCommand struct {
	// 资源对象类型名称，在租户和资源组下唯一

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资源对象类型别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 生效资源对象，支持通配，如{tcs_product="*"}

	Selector *string `json:"Selector,omitempty" name:"Selector"`
}

type DeleteVolumeRequest struct {
	*tchttp.BaseRequest

	// kubernetes集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 待删除Volume名

	Names []*string `json:"Names,omitempty" name:"Names"`
}

func (r *DeleteVolumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVolumeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRelocateInfoRequest struct {
	*tchttp.BaseRequest

	// 查询搬迁信息

	DescribeRelocateData *DescribeRelocateData `json:"DescribeRelocateData,omitempty" name:"DescribeRelocateData"`
}

func (r *DescribeRelocateInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRelocateInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMetricInstantByDimensionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数值

		Data []*QueryHistoryKV `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMetricInstantByDimensionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMetricInstantByDimensionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SilenceGroup struct {
	// 策略ID和Name

	Group *LabelPair `json:"Group,omitempty" name:"Group"`
	// 规则Id和Name数组

	Rules []*LabelPair `json:"Rules,omitempty" name:"Rules"`
}

type QueryOutbandStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryOutbandStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryOutbandStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSvrNicSet struct {
	// 返回结果详细说明

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 返回结果值，0成功， 1失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 服务器SN

	Sn *string `json:"Sn,omitempty" name:"Sn"`
}

type DeleteDaemonSetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDaemonSetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDaemonSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 服务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 修改的Service JSON定义

	Service *string `json:"Service,omitempty" name:"Service"`
	// 修改所属组件类型

	OwnerKind *string `json:"OwnerKind,omitempty" name:"OwnerKind"`
	// 修改所属组件名称

	OwnerName *string `json:"OwnerName,omitempty" name:"OwnerName"`
}

func (r *ModifyServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PerformServerTaskExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// PerformServerTaskSet

		PerformServerTaskSet []*string `json:"PerformServerTaskSet,omitempty" name:"PerformServerTaskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PerformServerTaskExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PerformServerTaskExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMetaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// meta数组

		Meta []*Meta `json:"Meta,omitempty" name:"Meta"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMetaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMetaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPodInstantMetricDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPodInstantMetricDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPodInstantMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelocateServerQueryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前返回数

		ReturnNum *int64 `json:"ReturnNum,omitempty" name:"ReturnNum"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 搬迁服务器状态查询信息

		RelocateServerSet *string `json:"RelocateServerSet,omitempty" name:"RelocateServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RelocateServerQueryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelocateServerQueryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceUsage struct {
	// 资源名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 历史24小时的数据点

	Metrics []*Metric `json:"Metrics,omitempty" name:"Metrics"`
	// 单点值

	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type DeleteLogRuleGroupsRequest struct {
	*tchttp.BaseRequest

	// 批量删除ids，"Ids":"1,2,3"

	Params *BatchDeleteIds `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteLogRuleGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogRuleGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WanData struct {
	// 网卡

	Nic *string `json:"Nic,omitempty" name:"Nic"`
	// IP地址

	Ipaddress *string `json:"Ipaddress,omitempty" name:"Ipaddress"`
}

type RebootNodeRequest struct {
	*tchttp.BaseRequest

	// 集群名，不填默认使用global

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 待重启的node名

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
}

func (r *RebootNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleVMLanIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟机回收内网IP结果

		VmRecyLanIPSet []*VMRecyLanIP `json:"VmRecyLanIPSet,omitempty" name:"VmRecyLanIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecycleVMLanIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleVMLanIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetVirtualMachineStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟机id

		VmId *string `json:"VmId,omitempty" name:"VmId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetVirtualMachineStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetVirtualMachineStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashFolderContent struct {
	// 数据库ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// Unique Id

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 目录名称

	Title *string `json:"Title,omitempty" name:"Title"`
	// 创建用户

	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`
	// 最近更新用户

	UpdatedBy *string `json:"UpdatedBy,omitempty" name:"UpdatedBy"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 更新时间

	Updated *string `json:"Updated,omitempty" name:"Updated"`
	// 是否包含ACL

	HasAcl *bool `json:"HasAcl,omitempty" name:"HasAcl"`
	// 版本

	Version *uint64 `json:"Version,omitempty" name:"Version"`
	// 是否可保存

	CanSave *bool `json:"CanSave,omitempty" name:"CanSave"`
	// 是否可编辑

	CanEdit *bool `json:"CanEdit,omitempty" name:"CanEdit"`
	// 是否可管理

	CanAdmin *bool `json:"CanAdmin,omitempty" name:"CanAdmin"`
}

type DescribeSystemComponentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件列表

		ComponentSet []*SystemComponent `json:"ComponentSet,omitempty" name:"ComponentSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSystemComponentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkPolicyRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 网络策略

	NetworkPolicy *NetworkPolicyStruct `json:"NetworkPolicy,omitempty" name:"NetworkPolicy"`
}

func (r *ModifyNetworkPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetworkPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCluster struct {
	// 集群id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 集群名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 集群网络类型

	Network *string `json:"Network,omitempty" name:"Network"`
	// 容器网络范围

	PodCIDR *string `json:"PodCIDR,omitempty" name:"PodCIDR"`
	// 服务网络范围

	ServiceCIDR *string `json:"ServiceCIDR,omitempty" name:"ServiceCIDR"`
	// 集群地址列表

	Apiserver []*string `json:"Apiserver,omitempty" name:"Apiserver"`
	// token

	Token *string `json:"Token,omitempty" name:"Token"`
	// 集群ca证书，base64编码

	Cacert *string `json:"Cacert,omitempty" name:"Cacert"`
	// 集群状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 集群版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 集群创建时间

	CreateAt *string `json:"CreateAt,omitempty" name:"CreateAt"`
	// cpu总数

	TotalCpu *string `json:"TotalCpu,omitempty" name:"TotalCpu"`
	// 已分配cpu

	AllocatedCpu *string `json:"AllocatedCpu,omitempty" name:"AllocatedCpu"`
	// 内存总数

	TotalMemory *string `json:"TotalMemory,omitempty" name:"TotalMemory"`
	// 已分配内存

	AllocatedMemory *string `json:"AllocatedMemory,omitempty" name:"AllocatedMemory"`
	// limit cpu

	LimitCpu *string `json:"LimitCpu,omitempty" name:"LimitCpu"`
	// limit memory

	LimitMemory *string `json:"LimitMemory,omitempty" name:"LimitMemory"`
	// request cpu

	RequestCpu *string `json:"RequestCpu,omitempty" name:"RequestCpu"`
	// request memory

	RequestMemory *string `json:"RequestMemory,omitempty" name:"RequestMemory"`
}

type DevOpsAction struct {
	// 自愈action

	Action *string `json:"Action,omitempty" name:"Action"`
	// action中文解释

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否开启，on，off

	Switch *string `json:"Switch,omitempty" name:"Switch"`
	// action的参数，可能为空

	Args []*string `json:"Args,omitempty" name:"Args"`
}

type OsDeployDebugInfo struct {
	// IdcId

	IdcId *string `json:"IdcId,omitempty" name:"IdcId"`
	// 物理服务器Sn

	Sns []*string `json:"Sns,omitempty" name:"Sns"`
}

type ListBaseNodesRequest struct {
	*tchttp.BaseRequest

	// 列表行数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 查询条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 集群列表

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

func (r *ListBaseNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListBaseNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCronJobsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件, 支持字段: Name

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序, 支持字段: Name

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 降序排序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
}

func (r *ListCronJobsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCronJobsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterAttributesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 删除保护

	DeleteProtection *string `json:"DeleteProtection,omitempty" name:"DeleteProtection"`
	// 资源模式

	ResourceMode *string `json:"ResourceMode,omitempty" name:"ResourceMode"`
	// 租户列表

	Tenants []*string `json:"Tenants,omitempty" name:"Tenants"`
}

func (r *ModifyClusterAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchAlertsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 告警详细信息

		Alerts []*Alert `json:"Alerts,omitempty" name:"Alerts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchAlertsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchAlertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashFolderQueryResult struct {
	// 数据ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// Unique ID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 目录名称

	Title *string `json:"Title,omitempty" name:"Title"`
}

type RenameServer struct {
	// 修改内容

	RenameServerModify *RenameServerModify `json:"RenameServerModify,omitempty" name:"RenameServerModify"`
	// 修改条件

	RenameServerCondition *RenameServerCondition `json:"RenameServerCondition,omitempty" name:"RenameServerCondition"`
}

type AddOutbandStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddOutbandStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOutbandStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCronJobsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 待删除CronJob列表

	NamespacedNames []*NamespacedName `json:"NamespacedNames,omitempty" name:"NamespacedNames"`
}

func (r *DeleteCronJobsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCronJobsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMetaMetricRequest struct {
	*tchttp.BaseRequest

	// 资源对象类型归属信息

	Params *MetaMetricQueryParam `json:"Params,omitempty" name:"Params"`
	// 资源对象类型名称

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *ListMetaMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMetaMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceLink struct {
	// 标识

	Id *string `json:"Id,omitempty" name:"Id"`
	// 源端服务ID

	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`
	// 目的端服务ID

	DestinationId *string `json:"DestinationId,omitempty" name:"DestinationId"`
	// 目的端服务端口

	DestinationPort *uint64 `json:"DestinationPort,omitempty" name:"DestinationPort"`
	// 流出字节数

	OutBytes *uint64 `json:"OutBytes,omitempty" name:"OutBytes"`
	// 流入字节数

	InBytes *uint64 `json:"InBytes,omitempty" name:"InBytes"`
	// 方向

	Verdict *uint64 `json:"Verdict,omitempty" name:"Verdict"`
	// 协议

	IPProtocol *string `json:"IPProtocol,omitempty" name:"IPProtocol"`
}

type RecycleVMLanIPListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server_vm"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"recycle_lan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 回收VM内网IP

	RecycleVMLanIPSet []*RecycleVMLanIP `json:"RecycleVMLanIPSet,omitempty" name:"RecycleVMLanIPSet"`
}

func (r *RecycleVMLanIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleVMLanIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StarredDashboardRequest struct {
	*tchttp.BaseRequest

	// dashboard id

	DashId *string `json:"DashId,omitempty" name:"DashId"`
}

func (r *StarredDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StarredDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateQueryRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询记录

		QueryRecord *QueryRecord `json:"QueryRecord,omitempty" name:"QueryRecord"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateQueryRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateQueryRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTopologyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 时间

		Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
		// 服务列表

		Services []*TopoService `json:"Services,omitempty" name:"Services"`
		// 服务间连接列表

		Links []*ServiceLink `json:"Links,omitempty" name:"Links"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceTopologyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceTopologyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClustersOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群状态列表

		ClusterStatus []*ClusterStatus `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClustersOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClustersOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartImageBuildingTaskRequest struct {
	*tchttp.BaseRequest

	// 任务名, 需要设置Name

	ImageBuildingTask *ImageBuildingTask `json:"ImageBuildingTask,omitempty" name:"ImageBuildingTask"`
}

func (r *StartImageBuildingTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartImageBuildingTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateServerSpecialWanIP struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// Vlan

	Vlan *string `json:"Vlan,omitempty" name:"Vlan"`
	// IP数

	IpNum *string `json:"IpNum,omitempty" name:"IpNum"`
}

type ConfigMap struct {
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// configmap的key列表

	Keys []*string `json:"Keys,omitempty" name:"Keys"`
}

type ResourceObject struct {
	// 服务树节点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 父节点ID

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 类型标签

	LabelName *string `json:"LabelName,omitempty" name:"LabelName"`
	// 服务树节点选择标签，格式为{a="b",c="d"}

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 资源归属，取用户的AppId

	Organization *string `json:"Organization,omitempty" name:"Organization"`
}

type CreateDevOpsScenesByClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDevOpsScenesByClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDevOpsScenesByClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListVirtualMachineStoragesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 存储列表

		Data []*VirtualMachinesStorage `json:"Data,omitempty" name:"Data"`
		// 读写模式set

		ReadWriteModeSet []*string `json:"ReadWriteModeSet,omitempty" name:"ReadWriteModeSet"`
		// 回收策略set

		ReclaimPolicySet []*string `json:"ReclaimPolicySet,omitempty" name:"ReclaimPolicySet"`
		// 状态set

		StatusSet []*string `json:"StatusSet,omitempty" name:"StatusSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListVirtualMachineStoragesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListVirtualMachineStoragesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRoutesByLabelSetRequest struct {
	*tchttp.BaseRequest

	// Product

	Product *string `json:"Product,omitempty" name:"Product"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// 等级

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 策略

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
}

func (r *SearchRoutesByLabelSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRoutesByLabelSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMetaResourceTypeRequest struct {
	*tchttp.BaseRequest

	// 资源组归属信息，运营端使用

	Params *ResourceOwnerQueryParam `json:"Params,omitempty" name:"Params"`
	// 资源对象类型名称

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源对象类型信息，仅Alias和ResourcePath支持修改

	Command *MetaResourceTypeUpdateCommand `json:"Command,omitempty" name:"Command"`
}

func (r *UpdateMetaResourceTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMetaResourceTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WithdrawServerExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 请求动作

	Op *string `json:"Op,omitempty" name:"Op"`
	// 物理服务器信息

	WithdrawServers []*WithdrawServers `json:"WithdrawServers,omitempty" name:"WithdrawServers"`
}

func (r *WithdrawServerExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WithdrawServerExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListenerBackend struct {
	// 监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 协议，TCP, UDP

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 监听端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 监听器后端

	Targets []*BackendTarget `json:"Targets,omitempty" name:"Targets"`
}

type LoadBalancer struct {
	// LoadBalancer类型，tgw，keepalived

	Provisioner *string `json:"Provisioner,omitempty" name:"Provisioner"`
	// VIP列表

	VIPList []*string `json:"VIPList,omitempty" name:"VIPList"`
}

type CreateTLSSecretRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 证书

	CertKey *string `json:"CertKey,omitempty" name:"CertKey"`
	// 私钥

	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
}

func (r *CreateTLSSecretRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTLSSecretRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionArgs struct {
	// 动作

	Action *string `json:"Action,omitempty" name:"Action"`
	// 开、关

	Switch *string `json:"Switch,omitempty" name:"Switch"`
	// 参数列表

	Args []*string `json:"Args,omitempty" name:"Args"`
}

type DeleteRAIDRequest struct {
	*tchttp.BaseRequest

	// "Op": "del_raid"

	Op *string `json:"Op,omitempty" name:"Op"`
	// "DeleteRAIDParams": {"Id":3}

	DeleteRAIDParams *DeleteRAIDParam `json:"DeleteRAIDParams,omitempty" name:"DeleteRAIDParams"`
}

func (r *DeleteRAIDRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRAIDRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIngressRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeIngressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIngressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterHistoryMetricDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterHistoryMetricDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterHistoryMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryMonitorDataRequest struct {
	*tchttp.BaseRequest

	// 指标，多个指标以逗号隔开

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 查询标签

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 开始时间，时间戳

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束时间，时间戳

	End *string `json:"End,omitempty" name:"End"`
	// 精度，比如60s

	Step *string `json:"Step,omitempty" name:"Step"`
}

func (r *DescribeHistoryMonitorDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHistoryMonitorDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodesStatsRequest struct {
	*tchttp.BaseRequest

	// 集群ID, 如global

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 符合kubernetes格式的过滤器, 如app=niginx

	LabelSelector *string `json:"LabelSelector,omitempty" name:"LabelSelector"`
}

func (r *DescribeNodesStatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodesStatsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyServersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateVMWanIP struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 设备SN

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 设备外网IP

	SvrWanIp *string `json:"SvrWanIp,omitempty" name:"SvrWanIp"`
}

type AllocationParams struct {
	// 网卡名称

	NicName *string `json:"NicName,omitempty" name:"NicName"`
	// 网卡类型

	NicType *string `json:"NicType,omitempty" name:"NicType"`
	// 分配的ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type CreateVirtualMachineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVirtualMachineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVirtualMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDaemonSetsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 待删除的DaemonSet的Namespace及Name

	NamespacedNames []*NamespacedName `json:"NamespacedNames,omitempty" name:"NamespacedNames"`
}

func (r *DeleteDaemonSetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDaemonSetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePodRequest struct {
	*tchttp.BaseRequest

	// 待删除Pod所在集群名，不传默认global

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 待删除Pod名

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 待删除Pod所在命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DeletePodRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePodRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIgniterNodeExData struct {
	// 内层data数组，元素为服务器SN

	Data []*string `json:"Data,omitempty" name:"Data"`
}

type TemplateParameter struct {
	// 模板描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 模板名字

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 模板类别，cluster, node, pod

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
	// 模板定义

	Content []*MonitorTemplateContent `json:"Content,omitempty" name:"Content"`
}

type GetDashboardVersionsByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 仪表盘更新历史记录

		Data []*DashboardVersion `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDashboardVersionsByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardVersionsByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRoutes struct {
	// 路由名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品名称

	Product *string `json:"Product,omitempty" name:"Product"`
	// 创建人

	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`
	// 更新人

	UpdatedBy *string `json:"UpdatedBy,omitempty" name:"UpdatedBy"`
	// page offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// page limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 升序字段，filed1,filed2,...

	Asc *string `json:"Asc,omitempty" name:"Asc"`
	// 降序字段，field1,filed2,...

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 是否启用正则匹配，可选true、false，默认false

	IsRegex *string `json:"IsRegex,omitempty" name:"IsRegex"`
}

type DeleteLoadBalancerListenersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLoadBalancerListenersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLoadBalancerListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTargetsRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 监听器 ID 列表

	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
	// 监听器协议类型

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 监听器端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
}

func (r *DescribeTargetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTargetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OutbandStrategies struct {
	// 删除/修改条件

	Condition *OutbandStrategiesCondition `json:"Condition,omitempty" name:"Condition"`
	// 修改内容

	Modify *AddOutbandStrategies `json:"Modify,omitempty" name:"Modify"`
}

type GetNodeDisruptionBudgetRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetNodeDisruptionBudgetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNodeDisruptionBudgetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceInstanceStatus struct {
	// 支撑服务类型, 如CRedis/CKafka

	Service *string `json:"Service,omitempty" name:"Service"`
	// 状态正常的实例个数

	NormalCount *int64 `json:"NormalCount,omitempty" name:"NormalCount"`
	// 状态异常的实例个数

	AbnormalCount *int64 `json:"AbnormalCount,omitempty" name:"AbnormalCount"`
	// 致命错误的实例个数

	FatalCount *int64 `json:"FatalCount,omitempty" name:"FatalCount"`
	// 状态未知的实例个数

	UnknownCount *int64 `json:"UnknownCount,omitempty" name:"UnknownCount"`
}

type CheckClusterDeletableResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否可删

		CanDelete *uint64 `json:"CanDelete,omitempty" name:"CanDelete"`
		// worker节点列表

		WorkerNodeInfos []*WorkerNodeInfo `json:"WorkerNodeInfos,omitempty" name:"WorkerNodeInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckClusterDeletableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckClusterDeletableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogConfigsRequest struct {
	*tchttp.BaseRequest

	// 日志名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// meta信息

	Metas []*Meta `json:"Metas,omitempty" name:"Metas"`
	// 分页偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回记录条数，默认值：10，最大值：100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 配置id

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribeLogConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMultiNicSet struct {
	// 自定义网卡名称

	DefinitionName *string `json:"DefinitionName,omitempty" name:"DefinitionName"`
	// 返回结果详细说明

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 返回结果值，0成功， 1失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
}

type NodeStatus struct {
	// Ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 角色

	Role *string `json:"Role,omitempty" name:"Role"`
	// 状态（正常Normal，异常Abnormal）

	Status *string `json:"Status,omitempty" name:"Status"`
	// 时间戳

	CreationTimestamp *int64 `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
}

type DescribeDefaultNetworkPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 项目间默认互通: reachable, 项目间默认隔离:isolated

		Mode *string `json:"Mode,omitempty" name:"Mode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDefaultNetworkPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefaultNetworkPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOutbandStrategyRequest struct {
	*tchttp.BaseRequest

	// 请求串类型说明。
	// 固定值： "Scheme":"outband_passwd_strategy"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方systemid	此ID为请求方向cmdb注册的权限ID，由cmdb系统管理员提供和管理。
	// 默认值："SystemId":"sid_test"

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方systemkey	和上面systemid类似，由cmdb系统管理员提供和管理，cmdb侧会根据id和key以及调用方ip三者共同确定请求方的合法性。
	// 默认值："SystemKey":"sid_test"

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 修改自定义带外密码策略详细内容

	Data *OutbandStrategyData `json:"Data,omitempty" name:"Data"`
}

func (r *ModifyOutbandStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOutbandStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerRecycleLanIPExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 请求动作

	Op *string `json:"Op,omitempty" name:"Op"`
	// 用户

	User *string `json:"User,omitempty" name:"User"`
	// 物理服务器回收内网IP接口

	ServersRecycleLanIp []*ServersRecycleLanIp `json:"ServersRecycleLanIp,omitempty" name:"ServersRecycleLanIp"`
}

func (r *ServerRecycleLanIPExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerRecycleLanIPExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterStatus struct {
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 版本

	KubernetesVersion *string `json:"KubernetesVersion,omitempty" name:"KubernetesVersion"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
}

type ConfigServerDefRequest struct {
	*tchttp.BaseRequest

	// 请求动作。"Command": "config_server_def"

	Command *string `json:"Command,omitempty" name:"Command"`
	// 此ID为请求方向cmdb注册的权限ID，由cmdb系统管理员提供和管理

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 和上面systemid类似，由cmdb系统管理员提供和管理，cmdb侧会根据id和key以及调用方ip三者共同确定请求方的合法性

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 本次初始化涉及的物理机信息

	ConfigServerDefSet []*ConfigServerDefSet `json:"ConfigServerDefSet,omitempty" name:"ConfigServerDefSet"`
	// Scheme: "server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
}

func (r *ConfigServerDefRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfigServerDefRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddMultiNicSet struct {
	// 返回结果详细说明

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 返回结果值，0成功， 1失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
}

type ImageInfo struct {
	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 镜像状态

	ImageStatus *int64 `json:"ImageStatus,omitempty" name:"ImageStatus"`
	// 镜像大小

	ImageSize *string `json:"ImageSize,omitempty" name:"ImageSize"`
	// 系统类型

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 系统架构

	SystemArch *string `json:"SystemArch,omitempty" name:"SystemArch"`
	// 系统平台

	SystemPlatform *string `json:"SystemPlatform,omitempty" name:"SystemPlatform"`
	// 系统版本

	SystemVersions *int64 `json:"SystemVersions,omitempty" name:"SystemVersions"`
	// 镜像描述

	ImageDescribe *string `json:"ImageDescribe,omitempty" name:"ImageDescribe"`
	// 镜像类型

	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`
	// 镜像url

	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

type OperationSheetArgument struct {
	// 步骤名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数列表

	Parameters []*KVPair `json:"Parameters,omitempty" name:"Parameters"`
}

type OperationSheetDetail struct {
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 当前状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 变更单执行流程DAG图

	DAG *DAG `json:"DAG,omitempty" name:"DAG"`
	// 变更单YAML

	YAML *string `json:"YAML,omitempty" name:"YAML"`
	// 总步骤数量

	TotalSteps *int64 `json:"TotalSteps,omitempty" name:"TotalSteps"`
	// 已完成的步骤数量

	CompletedSteps *int64 `json:"CompletedSteps,omitempty" name:"CompletedSteps"`
	// 用于显示的名字, 不唯一

	ShowName *string `json:"ShowName,omitempty" name:"ShowName"`
	// 关联的TAPD单子

	TAPD *string `json:"TAPD,omitempty" name:"TAPD"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 产品名

	Product *string `json:"Product,omitempty" name:"Product"`
	// 部分产品有父级

	ProductParent *string `json:"ProductParent,omitempty" name:"ProductParent"`
}

type DeleteLogRuleGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogRuleGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogRuleGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceTypeEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 资源对象元数据-对象类型事件列表

		Data []*MetaResourceTypeEvent `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListResourceTypeEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceTypeEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskFieldsEnumExRequest struct {
	*tchttp.BaseRequest

	// 需要枚举的字段

	FieldFilters []*string `json:"FieldFilters,omitempty" name:"FieldFilters"`
}

func (r *DescribeDiskFieldsEnumExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskFieldsEnumExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置id

		ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigServerDefSet struct {
	// 服务器SN

	SN *string `json:"SN,omitempty" name:"SN"`
	// 服务器架构

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	// 内网IP信息

	Lan []*ConfigServerDefSetLan `json:"Lan,omitempty" name:"Lan"`
	// ZoneName

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// IdcName

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// 配置初始化后的脚本集名称

	PreInstallScript *string `json:"PreInstallScript,omitempty" name:"PreInstallScript"`
	// 是否强制初始化(true:强制初始化；默认false)

	Force *bool `json:"Force,omitempty" name:"Force"`
}

type ApplyWebsocketTokenRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// Pod名称

	Pod *string `json:"Pod,omitempty" name:"Pod"`
	// 容器名称

	Container *string `json:"Container,omitempty" name:"Container"`
}

func (r *ApplyWebsocketTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyWebsocketTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSecretsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Secret列表

		SecretSet []*Secret `json:"SecretSet,omitempty" name:"SecretSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSecretsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSecretsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardTag struct {
	// Tag名称

	Term *string `json:"Term,omitempty" name:"Term"`
	// 出现次数

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type DeploymentFilterEntity struct {
	// 状态列表

	PhaseSet []*string `json:"PhaseSet,omitempty" name:"PhaseSet"`
}

type IngressRequest struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// IngressController名称

	IngressClassName *string `json:"IngressClassName,omitempty" name:"IngressClassName"`
	// 网络类型，内网Inner，外网Outer

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// Ingress规则列表

	Rules *IngressRule `json:"Rules,omitempty" name:"Rules"`
	// 安全传输配置

	IngressTLS []*IngressTLS `json:"IngressTLS,omitempty" name:"IngressTLS"`
}

type DescribeHistoryMonitorDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监控数据

		Data *KVPair `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHistoryMonitorDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHistoryMonitorDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListenersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监听器列表

		Listeners []*LoadBalancerListener `json:"Listeners,omitempty" name:"Listeners"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeListenersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecretData struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 密钥类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 数据，Value值默认base64编码

	Data []*KVPair `json:"Data,omitempty" name:"Data"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
}

type ListEventsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资源类型, 如Deployment/StatefulSet/DaemonSet/CronJob/Job/HPA

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 同时列出多个资源的事件时, 设置这多个资源的名字

	Names []*string `json:"Names,omitempty" name:"Names"`
}

func (r *ListEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Type struct {
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 模块名称

	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
	// 模块uuid

	TypeUuid *string `json:"TypeUuid,omitempty" name:"TypeUuid"`
	// 集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 模块健康状况 normal或者abNormal

	Status *string `json:"Status,omitempty" name:"Status"`
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type CheckMachineRequest struct {
	*tchttp.BaseRequest

	// 机器

	Machine *Machine `json:"Machine,omitempty" name:"Machine"`
}

func (r *CheckMachineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckMachineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelocateServerQueryRequest struct {
	*tchttp.BaseRequest

	// Scheme

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// relocate_query

	Op *string `json:"Op,omitempty" name:"Op"`
	// svr_asset_id,svr_idc_name_relocate_before,svr_rack_name_relocate_before,svr_pos_id_relocate_before,svr_lan_ip_relocate_before,svr_idc_name_relocate_after,svr_rack_name_relocate_after,svr_pos_id_relocate_after,svr_logic_area_relocate_after,svr_lan_ip_relocate_after,svr_lan_gateway_relocate_after,svr_lan_mask_relocate_after

	ResultItem *string `json:"ResultItem,omitempty" name:"ResultItem"`
	// "Filters":[{"Name":"SvrIdcId","Values":["2001"]}

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *RelocateServerQueryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelocateServerQueryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Volume struct {
	// Volume名, 集群内唯一

	Name *string `json:"Name,omitempty" name:"Name"`
	// 所属StorageClass名

	StorageClass *string `json:"StorageClass,omitempty" name:"StorageClass"`
	// 文件系统类型

	FileSystem *string `json:"FileSystem,omitempty" name:"FileSystem"`
	// 回收策略

	ReclaimPolicy *string `json:"ReclaimPolicy,omitempty" name:"ReclaimPolicy"`
	// 卷大小

	Capacity *string `json:"Capacity,omitempty" name:"Capacity"`
	// 当前实际使用率

	Usage *float64 `json:"Usage,omitempty" name:"Usage"`
	// 是否已挂载

	Mounted *bool `json:"Mounted,omitempty" name:"Mounted"`
	// 读写模式

	ReadWriteMode *string `json:"ReadWriteMode,omitempty" name:"ReadWriteMode"`
	// 对应的PVC名

	PersistentVolumeClaim *string `json:"PersistentVolumeClaim,omitempty" name:"PersistentVolumeClaim"`
	// 所属组户

	Tenant *string `json:"Tenant,omitempty" name:"Tenant"`
	// 所属项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 所属应用

	App *string `json:"App,omitempty" name:"App"`
	// 所属组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 当前状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 挂载此卷的Pod名

	Pod *string `json:"Pod,omitempty" name:"Pod"`
}

type DeleteOutbandStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOutbandStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOutbandStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirtualMachineAttributesRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 项目

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// VMID

	VmId *string `json:"VmId,omitempty" name:"VmId"`
}

func (r *DescribeVirtualMachineAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirtualMachineAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DriveDetailByClusterMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 个数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 节点列表

		Nodes []*DriveNode `json:"Nodes,omitempty" name:"Nodes"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DriveDetailByClusterMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DriveDetailByClusterMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInitMachineScriptRequest struct {
	*tchttp.BaseRequest

	// 待初始化的机器

	Machine *Machine `json:"Machine,omitempty" name:"Machine"`
}

func (r *DescribeInitMachineScriptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInitMachineScriptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRAIDParam struct {
	// 编号

	Id *string `json:"Id,omitempty" name:"Id"`
}

type AllocateServerWanIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateServerWanIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateServerWanIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDashFolderByUidRequest struct {
	*tchttp.BaseRequest

	// 仪表盘目录唯一ID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
}

func (r *DeleteDashFolderByUidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashFolderByUidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群状态列表

		ClusterStatus []*ClusterStatus `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClustersOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClustersOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SvrCurrentStatus struct {
	// 总量

	Total *string `json:"Total,omitempty" name:"Total"`
	// 运行中

	Running *string `json:"Running,omitempty" name:"Running"`
	// 部署中

	OnInstall *string `json:"OnInstall,omitempty" name:"OnInstall"`
	// 待部署

	PreInstall *string `json:"PreInstall,omitempty" name:"PreInstall"`
	// 部署失败

	InstallFailed *string `json:"InstallFailed,omitempty" name:"InstallFailed"`
}

type Workload struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type ModifyServerCondition struct {
	// 设备固资编号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type SearchDashboardsRequest struct {
	*tchttp.BaseRequest

	// 仪表盘查询参数

	Params *DashboardQueryParam `json:"Params,omitempty" name:"Params"`
}

func (r *SearchDashboardsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchDashboardsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MdComponentStatistic struct {
	// 中间件列表

	Component []*MdComponent `json:"Component,omitempty" name:"Component"`
}

type CreateJobRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// k8s Job类型的JSON字符串

	JobJSON *string `json:"JobJSON,omitempty" name:"JobJSON"`
}

func (r *CreateJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancersDetailRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
}

func (r *DescribeLoadBalancersDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoadBalancersDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListHPAsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// workload名字

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
	// workload类型, Deployment/StatefulSet

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// workload的命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *ListHPAsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListHPAsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancersDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 负载均衡实例列表

		LoadBalancerDetailSet []*LoadBalancerInstance `json:"LoadBalancerDetailSet,omitempty" name:"LoadBalancerDetailSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoadBalancersDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoadBalancersDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePodsStatsRequest struct {
	*tchttp.BaseRequest

	// 集群ID, 如global

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 符合kubernetes格式的过滤器, 如app=niginx

	LabelSelector *string `json:"LabelSelector,omitempty" name:"LabelSelector"`
	// 如指定, 只统计改命名空间下的Pod, 否则统计整个集群的Pod

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribePodsStatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePodsStatsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCloudProductStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云产品状态列表

		Product *Product `json:"Product,omitempty" name:"Product"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCloudProductStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCloudProductStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchNotificationMsgsRequest struct {
	*tchttp.BaseRequest

	// offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 开始时间

	StartsAt *string `json:"StartsAt,omitempty" name:"StartsAt"`
	// 结束时间

	EndsAt *string `json:"EndsAt,omitempty" name:"EndsAt"`
	// label matcher

	Query []*Matcher `json:"Query,omitempty" name:"Query"`
}

func (r *SearchNotificationMsgsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchNotificationMsgsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Alert struct {
	// 告警名

	AlertName *string `json:"AlertName,omitempty" name:"AlertName"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 告警策略

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 归属产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 告警对象

	Object *string `json:"Object,omitempty" name:"Object"`
	// 区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 确认人

	ConfirmedBy *string `json:"ConfirmedBy,omitempty" name:"ConfirmedBy"`
	// 确认时间，为空表示未确认

	ConfirmedAt *string `json:"ConfirmedAt,omitempty" name:"ConfirmedAt"`
	// 恢复状态，firing未恢复，resolved已恢复

	Status *string `json:"Status,omitempty" name:"Status"`
	// 告警级别，1级严重，2级重要，3级次要，4级警告

	Severity *string `json:"Severity,omitempty" name:"Severity"`
}

type AddImageRequest struct {
	*tchttp.BaseRequest

	// "Op": "add_image"

	Op *string `json:"Op,omitempty" name:"Op"`
	// "AddImageParams"{"ImageUrl":os_url,"RegionName": "region_name","ImageName": "test_os","OsType": "Linux","SystemArch": "X86","SystemPlatform" : "CentOS",,"SystemVersions": 7,"ImageDescribe": "示例文字"}

	AddImageParams *AddImageParams `json:"AddImageParams,omitempty" name:"AddImageParams"`
}

func (r *AddImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServersRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server or server_vm"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 操作类型："Op":"dispose"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 删除服务器参数集合

	DeleteServerSet []*DeleteServer `json:"DeleteServerSet,omitempty" name:"DeleteServerSet"`
}

func (r *DeleteServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaResourceTypeEvent struct {
	// 事件名称

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 事件维度

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
	// 事件中文名称

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 事件hash

	Fingerprint *uint64 `json:"Fingerprint,omitempty" name:"Fingerprint"`
	// 资源对象类型ID

	ResourceTypeId *int64 `json:"ResourceTypeId,omitempty" name:"ResourceTypeId"`
}

type ListNetworkPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 策略列表

		NetworkPolicySet []*NetworkPolicy `json:"NetworkPolicySet,omitempty" name:"NetworkPolicySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListNetworkPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListNetworkPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceTypesRequest struct {
	*tchttp.BaseRequest
}

func (r *ListResourceTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTaskExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryTaskExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTaskExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogMappingRequest struct {
	*tchttp.BaseRequest

	// 查询LogQL(只支持日志查询LogQL，不能包含count_over_time等指标计算)

	Query *string `json:"Query,omitempty" name:"Query"`
	// 是否只展示原始字段

	OnlyOriginalField *bool `json:"OnlyOriginalField,omitempty" name:"OnlyOriginalField"`
}

func (r *LogMappingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LogMappingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceFilterEntity struct {
	// 项目

	ProjectSet []*string `json:"ProjectSet,omitempty" name:"ProjectSet"`
	// 应用

	AppSet []*string `json:"AppSet,omitempty" name:"AppSet"`
	// 组件

	ComponentSet []*string `json:"ComponentSet,omitempty" name:"ComponentSet"`
	// 服务类型

	TypeSet []*string `json:"TypeSet,omitempty" name:"TypeSet"`
	// 命名空间

	NamespaceSet []*string `json:"NamespaceSet,omitempty" name:"NamespaceSet"`
	// 网络类型

	NetworkTypeSet []*string `json:"NetworkTypeSet,omitempty" name:"NetworkTypeSet"`
	// 项目展示名称

	ProjectShowNameSet []*string `json:"ProjectShowNameSet,omitempty" name:"ProjectShowNameSet"`
}

type Wan struct {
	// OS部署

	Data []*WanInfo `json:"Data,omitempty" name:"Data"`
}

type GetUpgradeConfigRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetUpgradeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetdeviceRelatedServersExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 条数限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 网络设备资产号

	NetdevAssetId *string `json:"NetdevAssetId,omitempty" name:"NetdevAssetId"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNetdeviceRelatedServersExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetdeviceRelatedServersExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRAIDRequest struct {
	*tchttp.BaseRequest

	// "Op": "change_raid"

	Op *string `json:"Op,omitempty" name:"Op"`
	// "ModifyRAIDParams": {"Id":3,"Name":"Linux"}

	ModifyRAIDParams *ModifyRAIDParam `json:"ModifyRAIDParams,omitempty" name:"ModifyRAIDParams"`
}

func (r *ModifyRAIDRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRAIDRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IpsDetailSet struct {
	// 返回结果详细说明

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 返回结果值，0成功， 1失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// IPDetailResuIPDetailResultInfo 带外ip详细信息ltInfo

	IPDetailResultInfo *IPDetailResultInfo `json:"IPDetailResultInfo,omitempty" name:"IPDetailResultInfo"`
}

type AddServersRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 操作类型："Op":"add"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 增加服务器信息，具体参数在复杂类型中定义

	AddServerSet []*AddServer `json:"AddServerSet,omitempty" name:"AddServerSet"`
}

func (r *AddServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddServersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddVMListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddVMListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVMListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImagePathResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像地址列表

		ImagePathList []*string `json:"ImagePathList,omitempty" name:"ImagePathList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImagePathResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagePathResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Field2enum struct {
	// SvrBelongingProducts

	SvrBelongingProducts []*string `json:"SvrBelongingProducts,omitempty" name:"SvrBelongingProducts"`
	// SvrCurrentStatus

	SvrCurrentStatus []*string `json:"SvrCurrentStatus,omitempty" name:"SvrCurrentStatus"`
	// SvrDeviceType

	SvrDeviceType []*string `json:"SvrDeviceType,omitempty" name:"SvrDeviceType"`
	// SvrLogicArea

	SvrLogicArea []*string `json:"SvrLogicArea,omitempty" name:"SvrLogicArea"`
	// SvrRackName

	SvrRackName []*string `json:"SvrRackName,omitempty" name:"SvrRackName"`
	// CmdbStatusToModify

	CmdbStatusToModify []*string `json:"CmdbStatusToModify,omitempty" name:"CmdbStatusToModify"`
}

type DescribeDeploySubtaskStepsExRequest struct {
	*tchttp.BaseRequest

	// SN号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// idcid

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// TaskId

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeDeploySubtaskStepsExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeploySubtaskStepsExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopKResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各个资源的名字及历史24小时数据点

		ResourceUsage []*ResourceUsage `json:"ResourceUsage,omitempty" name:"ResourceUsage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopKResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopKResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashboardTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仪表盘Tag统计

		Data []*DashboardTag `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDashboardTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllocation struct {
	// 默认网关

	DefaultGw *bool `json:"DefaultGw,omitempty" name:"DefaultGw"`
	// 网卡名

	NicName *string `json:"NicName,omitempty" name:"NicName"`
	// 网卡分配的ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 网卡类型

	NicType *string `json:"NicType,omitempty" name:"NicType"`
}

type DescribeNodeAlertsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
}

func (r *DescribeNodeAlertsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeAlertsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DevOpsScenes struct {
	// 自愈模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 自愈场景说明

	TemplateComment *string `json:"TemplateComment,omitempty" name:"TemplateComment"`
	// 触发条件

	Condition *string `json:"Condition,omitempty" name:"Condition"`
	// 生效状态

	Healable *bool `json:"Healable,omitempty" name:"Healable"`
	// 自愈情况统计

	Statistics *Statistics `json:"Statistics,omitempty" name:"Statistics"`
	// 自愈动作

	ActionList []*DevOpsAction `json:"ActionList,omitempty" name:"ActionList"`
	// 动作模板

	OperationName *string `json:"OperationName,omitempty" name:"OperationName"`
	// 是否显示在管理自愈页面on/off

	Display *bool `json:"Display,omitempty" name:"Display"`
}

type DeleteNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本次查询返回的数据条数，对应下面data节点的数组个数

		ReturnNum *int64 `json:"ReturnNum,omitempty" name:"ReturnNum"`
		// 满足本次查询条件的数据总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回结果详细说明

		Detail *string `json:"Detail,omitempty" name:"Detail"`
		// 返回结果值，0成功， 1失败

		Result *int64 `json:"Result,omitempty" name:"Result"`
		// 返回串中的ImageSet节点，规范协议请求，方便后续节点增加。

		ImageSet *ImageInfo `json:"ImageSet,omitempty" name:"ImageSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenRegexRequest struct {
	*tchttp.BaseRequest

	// 划线值

	FieldContent *string `json:"FieldContent,omitempty" name:"FieldContent"`
	// 划线前缀

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
}

func (r *GenRegexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenRegexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// "ImageSet": {"task_id": "xxx","Result": 0}

		ImageSet []*ImageSet `json:"ImageSet,omitempty" name:"ImageSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHPARequest struct {
	*tchttp.BaseRequest

	// HPA

	HPA *HPA `json:"HPA,omitempty" name:"HPA"`
	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *CreateHPARequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHPARequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateRuleGroupTemple struct {
	// 规则名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 评估间隔，建议1m

	Interval *string `json:"Interval,omitempty" name:"Interval"`
	// 归属产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 规则数组

	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`
	// 对象类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 乐观锁版本

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 策略Id，根据id是否>0区分create还是update

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 更新时间

	Updated *string `json:"Updated,omitempty" name:"Updated"`
	// labels

	Labels []*LabelPair `json:"Labels,omitempty" name:"Labels"`
}

type DescribeRelocateData struct {
	// 服务器IDC

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
}

type UpdateRoute struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品名称，支持多个

	Products []*string `json:"Products,omitempty" name:"Products"`
	// 告警等级，支持多个和all

	Severities []*string `json:"Severities,omitempty" name:"Severities"`
	// 是否发送告警

	SendFiring *bool `json:"SendFiring,omitempty" name:"SendFiring"`
	// 是否发送恢复

	SendResolved *bool `json:"SendResolved,omitempty" name:"SendResolved"`
	// 发送通道，支持微信、短信、企业微信、邮件和callback

	Channels []*LabelPair `json:"Channels,omitempty" name:"Channels"`
	// 用户id

	Users []*LabelPair `json:"Users,omitempty" name:"Users"`
	// 用户组id

	Groups []*LabelPair `json:"Groups,omitempty" name:"Groups"`
	// 乐观锁

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 静默时间，也即重发间隔

	RepeatInterval *string `json:"RepeatInterval,omitempty" name:"RepeatInterval"`
	// 策略，支持多个和all

	Strategies []*string `json:"Strategies,omitempty" name:"Strategies"`
	// 0,1->product+severity

	Flag *int64 `json:"Flag,omitempty" name:"Flag"`
}

type CreateSilenceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSilenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSilenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListenersRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 要查询的负载均衡监听器 ID 数组

	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

func (r *DescribeListenersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeListenersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageBuildingTask struct {
	// 任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 输出的镜像仓库

	Repository *string `json:"Repository,omitempty" name:"Repository"`
	// 输出的镜像TAG

	Tag *string `json:"Tag,omitempty" name:"Tag"`
	// Dockerfile

	Dockerfile *string `json:"Dockerfile,omitempty" name:"Dockerfile"`
	// 启动次数

	Runtimes *int64 `json:"Runtimes,omitempty" name:"Runtimes"`
	// 当前状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 上次完成时间

	LastCompletionTimestamp *string `json:"LastCompletionTimestamp,omitempty" name:"LastCompletionTimestamp"`
	// 运行任务所在集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 运行任务所在Pod

	Pod *string `json:"Pod,omitempty" name:"Pod"`
	// 运行任务所在Container

	Container *string `json:"Container,omitempty" name:"Container"`
	// 运行任务所在名字空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 运行中的任务状态

	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
}

type QueryRecord struct {
	// 创建时间 s级时间戳

	CreatedAt *uint64 `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 记录id

	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`
	// 用户id

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 筛选过滤查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 查询名称

	QueryName *string `json:"QueryName,omitempty" name:"QueryName"`
}

type Target struct {
	// 查询字符串

	Query *string `json:"Query,omitempty" name:"Query"`
	// 最大点数

	MaxPoint *int64 `json:"MaxPoint,omitempty" name:"MaxPoint"`
	// 排序，升序asc，降序desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// promQL组成元素，取代Query字段

	Metrics *PromQLMetrics `json:"Metrics,omitempty" name:"Metrics"`
	// 聚合函数

	AggrOp *string `json:"AggrOp,omitempty" name:"AggrOp"`
	// 分组的label列表

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
	// asc/desc

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 最大返回记录数

	Max *int64 `json:"Max,omitempty" name:"Max"`
	// 开始时间

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束时间

	End *string `json:"End,omitempty" name:"End"`
	// label筛选层

	PromQL *string `json:"PromQL,omitempty" name:"PromQL"`
	// orderBy使用的指标

	RangeVector *string `json:"RangeVector,omitempty" name:"RangeVector"`
	// 对promQL结果做算术运算

	MetricOperation *PromQLMetricOperation `json:"MetricOperation,omitempty" name:"MetricOperation"`
}

type CreateMultiNicParams struct {
	// 自定义网卡名称

	DefinitionName *string `json:"DefinitionName,omitempty" name:"DefinitionName"`
	// 网卡配置；示例：NicConfig":[{"DefaulTgw":true,"NicName":"eth0","NicType":"lan"}]

	NicConfig []*NicConfig `json:"NicConfig,omitempty" name:"NicConfig"`
	// bond配置；示例："BondConfig":[{"DefaulTgw":true,"NicName":"bond1","NicType":"lan"}]

	BondConfig []*NicConfig `json:"BondConfig,omitempty" name:"BondConfig"`
	// 是否为默认网卡配置；true/false

	Default *bool `json:"Default,omitempty" name:"Default"`
}

type DeleteSilencesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSilencesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSilencesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoadBalancerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePod struct {
	// pod名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 主机名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 主机ip

	NodeIP *string `json:"NodeIP,omitempty" name:"NodeIP"`
	// pod ip

	IP *string `json:"IP,omitempty" name:"IP"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
}

type AddResourceObjectRequest struct {
	*tchttp.BaseRequest

	// 资源归属信息

	Params *ResourceOwnerQueryParam `json:"Params,omitempty" name:"Params"`
	// 资源对象信息

	Command *CreateResourceCommand `json:"Command,omitempty" name:"Command"`
}

func (r *AddResourceObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddResourceObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListClusterStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *ListClusterStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListClusterStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportCmdbServersRequest struct {
	*tchttp.BaseRequest

	// 前端解析导入文件后生成的列表

	ImportServerList []*ImportServerList `json:"ImportServerList,omitempty" name:"ImportServerList"`
}

func (r *ImportCmdbServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportCmdbServersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCronJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCronJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCronJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIngressRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// Ingress名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 修改的Ingress JSON内容

	Ingress *string `json:"Ingress,omitempty" name:"Ingress"`
	// 所属类型

	OwnerKind *string `json:"OwnerKind,omitempty" name:"OwnerKind"`
	// 所属名称

	OwnerName *string `json:"OwnerName,omitempty" name:"OwnerName"`
}

func (r *ModifyIngressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIngressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancerRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例名称

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 负载均衡网络类型

	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
}

func (r *ModifyLoadBalancerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoadBalancerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomScriptsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li> CustomScriptName - String List - 是否必填：否 -（过滤条件）按照脚本名过滤。</li>
	// <li> CustomScriptSetName - String List - 是否必填：否 - （过滤条件）通过脚本集名字过滤，如果需要没有加入脚本集的脚本，Values增加""，例如["xxx", ""]。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCustomScriptsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomScriptsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunScriptOutput struct {
	// IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 输出

	Output *string `json:"Output,omitempty" name:"Output"`
	// 错误信息

	Error *string `json:"Error,omitempty" name:"Error"`
}

type VMRecyWanIP struct {
	// 结果详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 服务器内网IP

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
	// 服务器编号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
}

type Aggs struct {
	// 自定义聚合命名

	AggName *string `json:"AggName,omitempty" name:"AggName"`
	// 聚合行为

	AggOperator *AggOperator `json:"AggOperator,omitempty" name:"AggOperator"`
}

type DriveNode struct {
	// 节点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 节点角色

	Role *string `json:"Role,omitempty" name:"Role"`
	// 节点ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 节点状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 节点对应指标

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 下钻时间

	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
	// 下钻值

	Latest *float64 `json:"Latest,omitempty" name:"Latest"`
}

type MetaResourceCatalog struct {
	// 资源目录ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 资源目录名称，英文

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资源目录别名，用于产品展示

	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type ModifyOutbandStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyOutbandStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyOutbandStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSvrNicAllocationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 偏移量

		SvrNicSet []*SvrNicSet `json:"SvrNicSet,omitempty" name:"SvrNicSet"`
		// 返回总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSvrNicAllocationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSvrNicAllocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PVC struct {
	// PVC的UUID, 集群内唯一

	UID *string `json:"UID,omitempty" name:"UID"`
	// PVC的Name, 集群内的Namespace下唯一

	Name *string `json:"Name,omitempty" name:"Name"`
	// 此PVC绑定的PV

	PersistentVolume *string `json:"PersistentVolume,omitempty" name:"PersistentVolume"`
	// PVC所在的命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// PVC当前状态, 值待定

	Status *string `json:"Status,omitempty" name:"Status"`
	// PVC所属项目, 如果不属于任何一个项目为空

	Project *string `json:"Project,omitempty" name:"Project"`
	// 存储大小, 如100Gi

	Capacity *string `json:"Capacity,omitempty" name:"Capacity"`
	// 创建时间, 格式如2006-01-02T15:04:05

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 读写模式, 可能有多个模式. 例如[ROX, RWO]

	AccessModes []*string `json:"AccessModes,omitempty" name:"AccessModes"`
	// PVC描述文本

	Description *string `json:"Description,omitempty" name:"Description"`
	// 挂载类型, Filesystem/Block分别对应文件系统/块设备

	VolumeMode *string `json:"VolumeMode,omitempty" name:"VolumeMode"`
	// PVC关联的StorageClass名

	StorageClass *string `json:"StorageClass,omitempty" name:"StorageClass"`
}

type DescribeTopKResourcesRequest struct {
	*tchttp.BaseRequest

	// 待查询的资源类型, cluster: 集群, node: 节点, project: 项目, application: 应用

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 排序类型: cpu: 按cpu使用率排序, memory: 按内存使用率排序, cpuquota: Kind必须为project, 按照项目的cpuquota排序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 取Top几

	TopK *int64 `json:"TopK,omitempty" name:"TopK"`
}

func (r *DescribeTopKResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopKResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashboardPermissionsByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 仪表盘权限查询

		Data []*DashAclContent `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDashboardPermissionsByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardPermissionsByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashFolderPermissionsByUidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 仪表盘目录权限查询

		Data []*DashAclContent `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDashFolderPermissionsByUidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashFolderPermissionsByUidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPublicKeyRequest struct {
	*tchttp.BaseRequest
}

func (r *GetPublicKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPublicKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBatchDashboardByIdRequest struct {
	*tchttp.BaseRequest

	// dashboard id数组

	DashIds []*string `json:"DashIds,omitempty" name:"DashIds"`
}

func (r *DeleteBatchDashboardByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBatchDashboardByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDefaultNetworkPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetDefaultNetworkPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDefaultNetworkPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyLogConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true 校验成功 false 校验失败

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyLogConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateLogRuleGroup struct {
	// 策略名称，同一个租户不可重复

	Name *string `json:"Name,omitempty" name:"Name"`
	// 评估间隔，建议1m

	Interval *string `json:"Interval,omitempty" name:"Interval"`
	// 产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 告警对象Labels

	Objects []*LabelPair `json:"Objects,omitempty" name:"Objects"`
	// 规则

	Rules []*LogRule `json:"Rules,omitempty" name:"Rules"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 是否启用，true启用，false暂不启用

	Enable *string `json:"Enable,omitempty" name:"Enable"`
	// 乐观锁版本, 创建传0

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 策略Id，根据id是否>0区分create还是update

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 更新时间

	Updated *string `json:"Updated,omitempty" name:"Updated"`
	// 策略类型，1为自定义事件生成策略，默认为0 ，告警策略

	GroupType *int64 `json:"GroupType,omitempty" name:"GroupType"`
	// 自定义事件名称，GroupType为1时必填

	EventName *string `json:"EventName,omitempty" name:"EventName"`
}

type CreateIngressRequest struct {
	*tchttp.BaseRequest

	// 集群ID, 如"global"

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// ingress名字, 需符合k8s命名规范

	Name *string `json:"Name,omitempty" name:"Name"`
	// 所属项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// Ingress JSON字符串

	Ingress *string `json:"Ingress,omitempty" name:"Ingress"`
	// 所属类型

	OwnerKind *string `json:"OwnerKind,omitempty" name:"OwnerKind"`
	// 所属名称

	OwnerName *string `json:"OwnerName,omitempty" name:"OwnerName"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *CreateIngressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIngressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeregisterTargetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeregisterTargetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeregisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryHistoryKV struct {
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标值数组

	Values *QueryHistoryResult `json:"Values,omitempty" name:"Values"`
}

type RAIDInfo struct {
	// 编号

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// RAID名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ServerSpecialAllocateLanIPInfo struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type DeleteRetentionsRequest struct {
	*tchttp.BaseRequest

	// 批量删除Retention规则

	Command *BatchDeleteIds `json:"Command,omitempty" name:"Command"`
}

func (r *DeleteRetentionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRetentionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOsDeployDebugInfoExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询os部署调试信息出参

		DebugInfo *string `json:"DebugInfo,omitempty" name:"DebugInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOsDeployDebugInfoExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOsDeployDebugInfoExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DAGStep struct {
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 状态, Running/Ready

	Status *string `json:"Status,omitempty" name:"Status"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 参数

	Arguments *Arguments `json:"Arguments,omitempty" name:"Arguments"`
	// 运行在哪些宿主机上

	RunningNodes []*string `json:"RunningNodes,omitempty" name:"RunningNodes"`
}

type GetDashFolderByUidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 仪表盘目录内容

		Data *DashFolderContent `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDashFolderByUidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashFolderByUidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StatefulSetRequest struct {
	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 副本数

	Replicas *uint64 `json:"Replicas,omitempty" name:"Replicas"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 滚动升级策略

	Strategy *StatefulSetUpdateStrategy `json:"Strategy,omitempty" name:"Strategy"`
	// Pod模板

	Template *PodTemplate `json:"Template,omitempty" name:"Template"`
}

type GenTimeFormatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志时间格式

		TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenTimeFormatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenTimeFormatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkloadPodsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 工作负载名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 工作负载类型，Deployment, DaemonSet, StatefulSet, Job

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 每页个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 按字段搜索

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
}

func (r *ListWorkloadPodsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListWorkloadPodsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRAIDListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// RAID查询结果

		RaidDictionarySet []*RAIDDictionaryInfo `json:"RaidDictionarySet,omitempty" name:"RaidDictionarySet"`
		// 返回总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRAIDListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRAIDListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorTemplateContent struct {
	// 监控组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 监控tab

	MetricSpec []*MetricSpec `json:"MetricSpec,omitempty" name:"MetricSpec"`
}

type DescribePhysvrsOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 物理服务器资源概览出参

		Overview *Overview `json:"Overview,omitempty" name:"Overview"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePhysvrsOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePhysvrsOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SocPerformServerTaskExRequest struct {
	*tchttp.BaseRequest

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// SvrIdcName

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 操作码：1（开机），2（关机），3（重启）

	OpCode *string `json:"OpCode,omitempty" name:"OpCode"`
	// 操作的服务器信息

	Servers []*SocPerformServerTaskEsServers `json:"Servers,omitempty" name:"Servers"`
	// SvrIdcId

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
}

func (r *SocPerformServerTaskExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SocPerformServerTaskExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IngressFilterEntity struct {
	// 网络类型，内外网

	NetworkTypeSet []*string `json:"NetworkTypeSet,omitempty" name:"NetworkTypeSet"`
	// VIP

	VIPSet []*string `json:"VIPSet,omitempty" name:"VIPSet"`
	// 应用

	AppSet []*string `json:"AppSet,omitempty" name:"AppSet"`
	// 项目

	ProjectSet []*string `json:"ProjectSet,omitempty" name:"ProjectSet"`
	// 项目显示名称

	ProjectShowNameSet []*string `json:"ProjectShowNameSet,omitempty" name:"ProjectShowNameSet"`
}

type CreateDeploymentRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 待创建的Deployment

	Deployment *string `json:"Deployment,omitempty" name:"Deployment"`
	// 待创建的Ingress数组

	Ingresses []*string `json:"Ingresses,omitempty" name:"Ingresses"`
	// 待创建的Service数组

	Services []*string `json:"Services,omitempty" name:"Services"`
}

func (r *CreateDeploymentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeploymentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMultiNicDefinitionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 增加网卡信息出参集合

		MultiNicSet *MultiNicSet `json:"MultiNicSet,omitempty" name:"MultiNicSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMultiNicDefinitionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMultiNicDefinitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNodeHistoryMetricDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNodeHistoryMetricDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNodeHistoryMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaEventQueryParam struct {
	// AppID

	Organization *string `json:"Organization,omitempty" name:"Organization"`
}

type DescribeK8sComponentRequest struct {
	*tchttp.BaseRequest

	// cluster

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *DescribeK8sComponentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sComponentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePodsStatsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 可用的pod数量, 可用的判断依据是pod.status.condtions中类型为Ready的condition是否为True

		AvailableCount *int64 `json:"AvailableCount,omitempty" name:"AvailableCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePodsStatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePodsStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNodeDisruptionBudgetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 风向控制配置数据

		Data *Ndb `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNodeDisruptionBudgetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNodeDisruptionBudgetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleGroupTemplesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRuleGroupTemplesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRuleGroupTemplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNotificationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNotificationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListContainersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// container列表

		ContainerSet []*ContainerSet `json:"ContainerSet,omitempty" name:"ContainerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListContainersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListContainersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Event struct {
	// 首次出现时间

	FirstTimestamp *string `json:"FirstTimestamp,omitempty" name:"FirstTimestamp"`
	// 最后出现时间

	LastTimestamp *string `json:"LastTimestamp,omitempty" name:"LastTimestamp"`
	// 级别, Normal/Warning

	Type *string `json:"Type,omitempty" name:"Type"`
	// 资源类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 资源名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 内容

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 详细描述

	Message *string `json:"Message,omitempty" name:"Message"`
	// 出现次数

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type CreateDeploymentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDeploymentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDeploymentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateListenerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监听器ID列表

		ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateListenerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelConditionInfo struct {
	// 要修改标签的id

	Id *string `json:"Id,omitempty" name:"Id"`
}

type DeleteRuleGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRuleGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRuleGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomScriptSetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自定义脚本集的详细信息列表

		CustomScriptSets []*CustomScriptSets `json:"CustomScriptSets,omitempty" name:"CustomScriptSets"`
		// 自定义脚本集的总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomScriptSetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomScriptSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceInstanceStatusOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeServiceInstanceStatusOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceInstanceStatusOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopKPodsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// top值

		ResourceUsage []*TopPodUsage `json:"ResourceUsage,omitempty" name:"ResourceUsage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopKPodsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopKPodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkPolicyRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 网络策略

	NetworkPolicy *NetworkPolicyStruct `json:"NetworkPolicy,omitempty" name:"NetworkPolicy"`
}

func (r *CreateNetworkPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateLogRuleGroupTempleRequest struct {
	*tchttp.BaseRequest

	// 策略模板定义

	Command *CreateOrUpdateRuleGroupTemple `json:"Command,omitempty" name:"Command"`
}

func (r *CreateOrUpdateLogRuleGroupTempleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateLogRuleGroupTempleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportCmdbServersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportCmdbServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportCmdbServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SocPerformServerTaskExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// SocPerformServerTaskSet

		SocPerformServerTaskSet []*string `json:"SocPerformServerTaskSet,omitempty" name:"SocPerformServerTaskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SocPerformServerTaskExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SocPerformServerTaskExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataQueryParams struct {
	// 查询表达式

	Query *string `json:"Query,omitempty" name:"Query"`
	// 时间，支持RFC3339和Unixj时间戳

	Time *string `json:"Time,omitempty" name:"Time"`
	// 超时时间

	Timeout *string `json:"Timeout,omitempty" name:"Timeout"`
	// 开始时间，支持RFC3339和Unixj时间戳

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束时间，支持RFC3339和Unixj时间戳

	End *string `json:"End,omitempty" name:"End"`
	// 查询配置条件

	Matches *string `json:"Matches,omitempty" name:"Matches"`
	// 步长

	Step *string `json:"Step,omitempty" name:"Step"`
	// 限制条数

	Limit *string `json:"Limit,omitempty" name:"Limit"`
	// BACKWARD倒序 FORWARD正序

	Direction *string `json:"Direction,omitempty" name:"Direction"`
}

type DescribePhysvrsListExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TotalCount

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// PhysvrsListSet

		PhysvrsListSet []*PhysvrsListSet `json:"PhysvrsListSet,omitempty" name:"PhysvrsListSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePhysvrsListExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePhysvrsListExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenLineHeaderRegexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 行首正则

		LineHeaderRegex *string `json:"LineHeaderRegex,omitempty" name:"LineHeaderRegex"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenLineHeaderRegexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenLineHeaderRegexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListNativeStorageClassesRequest struct {
	*tchttp.BaseRequest

	// kubernetes集群名, 如global

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 搜索条件

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

func (r *ListNativeStorageClassesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListNativeStorageClassesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddLabelInfo struct {
	// 标签建

	LabelKey *string `json:"LabelKey,omitempty" name:"LabelKey"`
	// 标签值

	LabelValue *string `json:"LabelValue,omitempty" name:"LabelValue"`
}

type SubSubMeta struct {
	// meta key

	MetaKey *string `json:"MetaKey,omitempty" name:"MetaKey"`
	// meta value

	MetaValue *string `json:"MetaValue,omitempty" name:"MetaValue"`
	// meta name

	MetaName *string `json:"MetaName,omitempty" name:"MetaName"`
	// name

	Names []*string `json:"Names,omitempty" name:"Names"`
	// 配置数量

	ConfigCount *uint64 `json:"ConfigCount,omitempty" name:"ConfigCount"`
}

type AllocateServerSpecialWanIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分配服务器特殊外网IP结果信息

		SvrSpecAllocWanIPSet []*SvrSpecAllocWanIP `json:"SvrSpecAllocWanIPSet,omitempty" name:"SvrSpecAllocWanIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateServerSpecialWanIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateServerSpecialWanIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Job的JSON字符串

		JobJSON *string `json:"JobJSON,omitempty" name:"JobJSON"`
		// 状态

		Phase *string `json:"Phase,omitempty" name:"Phase"`
		// 上次更新时间

		LastUpdateTimestamp *string `json:"LastUpdateTimestamp,omitempty" name:"LastUpdateTimestamp"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDataBaradMetricRequest struct {
	*tchttp.BaseRequest

	// 命名空间，每个云产品会有一个命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 监控统计周期。默认为取值为300，单位为s

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 起始时间，如 2018-01-01 00:00:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，默认为当前时间。 endTime不能小于startTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 实例对象的维度组合

	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
}

func (r *GetDataBaradMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDataBaradMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IngressRule struct {
	// 域名

	Host *string `json:"Host,omitempty" name:"Host"`
	// 路径规则

	Paths []*IngressPath `json:"Paths,omitempty" name:"Paths"`
	// 是否开启TLS

	TLS *bool `json:"TLS,omitempty" name:"TLS"`
	// Secret名称

	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

type NetworkPolicyPort struct {
	// 协议(TCP, UDP, SCTP), 默认TCP

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type DeleteSvrNicAllocationRequest struct {
	*tchttp.BaseRequest

	// 解绑物理服务器多网卡策略出参

	DeleteSvrNicParams *DeleteSvrNicAllocation `json:"DeleteSvrNicParams,omitempty" name:"DeleteSvrNicParams"`
}

func (r *DeleteSvrNicAllocationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSvrNicAllocationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVirtualMachineRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 虚拟机Id和namespace列表

	VMIdPair []*VirtualMachineNamespaceByName `json:"VMIdPair,omitempty" name:"VMIdPair"`
}

func (r *DeleteVirtualMachineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVirtualMachineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDcosFunctionServiceRelationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// dcos功能与服务ID映射关系集合

		FunctionServiceRelationItems []*DcosFunctionServiceItem `json:"FunctionServiceRelationItems,omitempty" name:"FunctionServiceRelationItems"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDcosFunctionServiceRelationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDcosFunctionServiceRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetImageFieldsEnumRequest struct {
	*tchttp.BaseRequest

	// "Op": "get_image_enum"

	Op *string `json:"Op,omitempty" name:"Op"`
}

func (r *GetImageFieldsEnumRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetImageFieldsEnumRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailReason struct {
	// IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 失败原因

	Error *string `json:"Error,omitempty" name:"Error"`
	// 失败模块

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type Region struct {
	// region的主域名，可以拼接出其他域名来

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// region的名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// region的角色，Master是主region，slave是子region

	Role *string `json:"Role,omitempty" name:"Role"`
}

type ResourceOwnerQueryParam struct {
	// 租户ID，对应AppId，运营端代理租户身份时使用

	Organization *string `json:"Organization,omitempty" name:"Organization"`
}

type ApplyWsTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// token字符串

		Token *string `json:"Token,omitempty" name:"Token"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyWsTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyWsTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIgniterImagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIgniterImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIgniterImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCloudProductStatusRequest struct {
	*tchttp.BaseRequest

	// 查询云产品名字，逗号分隔多个云产品名字

	Names *string `json:"Names,omitempty" name:"Names"`
}

func (r *GetCloudProductStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCloudProductStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServerHarddisksRequest struct {
	*tchttp.BaseRequest

	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeServerHarddisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServerHarddisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleGroupsRequest struct {
	*tchttp.BaseRequest

	// 批量删除ids，"Ids":"1,2,3"

	Params *BatchDeleteIds `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteRuleGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRuleGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRoute struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品名称，支持多个

	Products []*string `json:"Products,omitempty" name:"Products"`
	// 告警等级，支持多个和all

	Severities []*string `json:"Severities,omitempty" name:"Severities"`
	// 是否发送告警

	SendFiring *bool `json:"SendFiring,omitempty" name:"SendFiring"`
	// 是否发送恢复

	SendResolved *bool `json:"SendResolved,omitempty" name:"SendResolved"`
	// 发送通道，支持微信、短信、企业微信、邮件，callback

	Channels []*LabelPair `json:"Channels,omitempty" name:"Channels"`
	// 用户id

	Users []*LabelPair `json:"Users,omitempty" name:"Users"`
	// 用户组id

	Groups []*LabelPair `json:"Groups,omitempty" name:"Groups"`
	// 静默时间，也即重发间隔

	RepeatInterval *string `json:"RepeatInterval,omitempty" name:"RepeatInterval"`
	// 策略，支持多个和all

	Strategies []*string `json:"Strategies,omitempty" name:"Strategies"`
	// 0,1->product+severity

	Flag *int64 `json:"Flag,omitempty" name:"Flag"`
}

type ModuleScript struct {
	// 模块，NotFormatDisk（磁盘未格式化），TimeSyncError（时间同步问题），MachineNotInit（参数未初始化）

	Module *string `json:"Module,omitempty" name:"Module"`
	// 脚本

	Script *string `json:"Script,omitempty" name:"Script"`
}

type CreateOrUpdateLogRuleGroupTempleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrUpdateLogRuleGroupTempleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateLogRuleGroupTempleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDashAclCommand struct {
	// 所有要添加的权限列表

	Items []*DashAclItem `json:"Items,omitempty" name:"Items"`
}

type DescribeLabelRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"label"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 操作类型："Op":"query"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 返回数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回参数

	ResultItem *string `json:"ResultItem,omitempty" name:"ResultItem"`
}

func (r *DescribeLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashboardContentByUidRequest struct {
	*tchttp.BaseRequest

	// 仪表唯一ID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
}

func (r *GetDashboardContentByUidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardContentByUidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterMachine struct {
	// IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// ssh用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// ssh密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 数据盘挂载，比如 {"MountPoint": "/data", "Device": "/dev/vdb"}

	Disks []*MachineDisk `json:"Disks,omitempty" name:"Disks"`
	// 节点版本

	MachineVersion *string `json:"MachineVersion,omitempty" name:"MachineVersion"`
}

type DeleteIngressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIngressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchAlert struct {
	// 当前状态，firing/resolved, 分别表示未恢复/已恢复

	Status *string `json:"Status,omitempty" name:"Status"`
	// 产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 地区

	Region *string `json:"Region,omitempty" name:"Region"`
	// 告警级别， 1、2、3、4, 严重等级依次降低

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 确认状态，true/false

	Confirmed *string `json:"Confirmed,omitempty" name:"Confirmed"`
	// time range start

	StartsAt *string `json:"StartsAt,omitempty" name:"StartsAt"`
	// time range end

	EndsAt *string `json:"EndsAt,omitempty" name:"EndsAt"`
	// page offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// page limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 升序排列字段，逗号分隔多个字段

	Asc *string `json:"Asc,omitempty" name:"Asc"`
	// 降序排列字段，逗号分隔多个字段

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 对象查询字段，支持正则匹配

	Object *string `json:"Object,omitempty" name:"Object"`
	// 屏蔽规则id

	SilencedBy *string `json:"SilencedBy,omitempty" name:"SilencedBy"`
	// 告警策略名称

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 告警名称，支持正则匹配，替代Alertname

	AlertName *string `json:"AlertName,omitempty" name:"AlertName"`
	// 是否启用正则匹配，可选true、false，默认

	IsRegex *string `json:"IsRegex,omitempty" name:"IsRegex"`
	// 产品类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 废弃

	Alertname *string `json:"Alertname,omitempty" name:"Alertname"`
	// metric log

	Class *string `json:"Class,omitempty" name:"Class"`
	// notification id

	NotificationId *int64 `json:"NotificationId,omitempty" name:"NotificationId"`
}

type ListBaseNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 节点列表

		NodeSet []*BaseNode `json:"NodeSet,omitempty" name:"NodeSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListBaseNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListBaseNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNodeRequest struct {
	*tchttp.BaseRequest

	// 描述待添加主机的信息

	MachineSet []*Machine `json:"MachineSet,omitempty" name:"MachineSet"`
}

func (r *CreateNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CpuRequest

		CpuRequest *float64 `json:"CpuRequest,omitempty" name:"CpuRequest"`
		// CpuLimit

		CpuLimit *float64 `json:"CpuLimit,omitempty" name:"CpuLimit"`
		// MemoryRequest

		MemoryRequest *float64 `json:"MemoryRequest,omitempty" name:"MemoryRequest"`
		// MemoryLimit

		MemoryLimit *float64 `json:"MemoryLimit,omitempty" name:"MemoryLimit"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMonitorResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLogConfigStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateLogConfigStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogConfigStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMetaMetricRequest struct {
	*tchttp.BaseRequest

	// 资源归属信息

	Params *ResourceOwnerQueryParam `json:"Params,omitempty" name:"Params"`
	// 资源对象类型名称

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 指标指纹ID

	Fingerprint *string `json:"Fingerprint,omitempty" name:"Fingerprint"`
	// 指标更新信息，包含Alias，Type，Help，Unit

	Command *MetaMetricUpdateCommand `json:"Command,omitempty" name:"Command"`
}

func (r *UpdateMetaMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMetaMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateListenerRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 要将监听器创建到哪些端口，每个端口对应一个新的监听器

	Ports []*int64 `json:"Ports,omitempty" name:"Ports"`
	// 要创建的监听器名称列表，名称与Ports数组按序一一对应

	ListenerNames []*string `json:"ListenerNames,omitempty" name:"ListenerNames"`
	// 健康检查相关参数

	HealthCheck *ListenerHealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`
	// 协议，TCP或UDP

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *CreateListenerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateListenerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePhysvrsListExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求动作

	Op *string `json:"Op,omitempty" name:"Op"`
	// 查询条件

	FieldFilters *string `json:"FieldFilters,omitempty" name:"FieldFilters"`
	// RoughItem

	RoughItem *string `json:"RoughItem,omitempty" name:"RoughItem"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePhysvrsListExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePhysvrsListExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRetentionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchRetentionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRetentionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardFolderRequest struct {
	*tchttp.BaseRequest

	// dashboard id数组

	DashIds []*string `json:"DashIds,omitempty" name:"DashIds"`
	// 需要移动目标的文件夹id

	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
}

func (r *ModifyDashboardFolderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDashboardFolderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelPair struct {
	// Label Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Label Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeleteBatchDashboardByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功执行的dashboard id数组

		Succeeded []*string `json:"Succeeded,omitempty" name:"Succeeded"`
		// 失败执行的dashboard id数组

		Failed []*string `json:"Failed,omitempty" name:"Failed"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBatchDashboardByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBatchDashboardByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 满足过滤条件的负载均衡实例总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回的负载均衡实例数组

		LoadBalancerSet []*LoadBalancerInstance `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoadBalancersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateLogRuleGroupRequest struct {
	*tchttp.BaseRequest

	// 策略定义

	Command *CreateOrUpdateRuleGroup `json:"Command,omitempty" name:"Command"`
}

func (r *CreateOrUpdateLogRuleGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateLogRuleGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateRuleGroupTempleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrUpdateRuleGroupTempleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateRuleGroupTempleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVolumeRequest struct {
	*tchttp.BaseRequest

	// kubernetes集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 选中的StorageClass名

	StorageClass *string `json:"StorageClass,omitempty" name:"StorageClass"`
	// Volume名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数

	Parameters []*VolumeParameter `json:"Parameters,omitempty" name:"Parameters"`
	// 项目名

	Project *string `json:"Project,omitempty" name:"Project"`
}

func (r *CreateVolumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVolumeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRuleGroups struct {
	// 策略id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// true or false

	Enable *string `json:"Enable,omitempty" name:"Enable"`
	// page offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// page limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 升序字段: field1[,field2 ...]

	Asc *string `json:"Asc,omitempty" name:"Asc"`
	// 降序字段: field1[,field2 ...]

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 归属产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 创建人

	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`
	// 是否启用正则匹配，可选true、false，默认false

	IsRegex *string `json:"IsRegex,omitempty" name:"IsRegex"`
	// 策略类型 1 自定义事件策略 0 告警策略

	GroupType *int64 `json:"GroupType,omitempty" name:"GroupType"`
	// alert record

	Type *string `json:"Type,omitempty" name:"Type"`
	// tcs type

	TcsType *string `json:"TcsType,omitempty" name:"TcsType"`
}

type UpdateDashFolderCommand struct {
	// Unique Id

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 目录名称

	Title *string `json:"Title,omitempty" name:"Title"`
	// 目录版本

	Version *uint64 `json:"Version,omitempty" name:"Version"`
	// 是否覆写

	Overwrite *bool `json:"Overwrite,omitempty" name:"Overwrite"`
}

type AllocateVMLanIP struct {
	// 服务器资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器序列号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 服务器内网IP

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
}

type MetaResourceTypeMetric struct {
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标维度

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
	// 指标别名，用于界面展示

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 指标单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 指标Hash

	Fingerprint *uint64 `json:"Fingerprint,omitempty" name:"Fingerprint"`
	// 资源对象类型ID

	ResourceTypeId *int64 `json:"ResourceTypeId,omitempty" name:"ResourceTypeId"`
}

type ProxyAllDataSourceRequestByNameRequest struct {
	*tchttp.BaseRequest

	// 额外路径

	AdditionalPath *string `json:"AdditionalPath,omitempty" name:"AdditionalPath"`
	// 代理方法，当前只支持GET

	Method *string `json:"Method,omitempty" name:"Method"`
	// 数据类型，当前支持metrics和logs

	Name *string `json:"Name,omitempty" name:"Name"`
	// 数据查询参数

	Params *DataQueryParams `json:"Params,omitempty" name:"Params"`
}

func (r *ProxyAllDataSourceRequestByNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ProxyAllDataSourceRequestByNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDashFolderByIdRequest struct {
	*tchttp.BaseRequest

	// 仪表盘目录唯一ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 仪表盘目录更新内容

	Command *UpdateDashFolderCommand `json:"Command,omitempty" name:"Command"`
}

func (r *UpdateDashFolderByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDashFolderByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListClusterStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListClusterStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListClusterStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterTargetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RegisterTargetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StarredDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 收藏的dashboard id

		DashId *string `json:"DashId,omitempty" name:"DashId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StarredDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StarredDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNodeInstantMetricDataRequest struct {
	*tchttp.BaseRequest

	// 查询指标名称，多个用逗号分隔

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 用于过滤node的标签集合，key=value形式，多个用逗号分隔

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 查询时间点

	Start *string `json:"Start,omitempty" name:"Start"`
	// 精度，如1m

	Step *string `json:"Step,omitempty" name:"Step"`
}

func (r *GetNodeInstantMetricDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNodeInstantMetricDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRelatedAlertNamesRequest struct {
	*tchttp.BaseRequest

	// 参数

	Params *GetRelatedAlertNames `json:"Params,omitempty" name:"Params"`
}

func (r *GetRelatedAlertNamesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRelatedAlertNamesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetImageFieldsEnumResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ImageFieldsEnumSet

		ImageFieldsEnumSet *ImageFieldsEnumSet `json:"ImageFieldsEnumSet,omitempty" name:"ImageFieldsEnumSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetImageFieldsEnumResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetImageFieldsEnumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogMappingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志字段映射

		Mapping []*LogField `json:"Mapping,omitempty" name:"Mapping"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LogMappingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LogMappingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFilter struct {
	// 待搜索字段

	SearchField *string `json:"SearchField,omitempty" name:"SearchField"`
	// 待搜搜值

	SearchValues []*string `json:"SearchValues,omitempty" name:"SearchValues"`
}

type DescribeServerRelatedNetdevicesRequest struct {
	*tchttp.BaseRequest

	// 服务器机架名

	SvrRackName *string `json:"SvrRackName,omitempty" name:"SvrRackName"`
	// 服务器机架位置编号

	SvrPosId *string `json:"SvrPosId,omitempty" name:"SvrPosId"`
}

func (r *DescribeServerRelatedNetdevicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServerRelatedNetdevicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteQueryRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录id

		RecordId *string `json:"RecordId,omitempty" name:"RecordId"`
		// 删除是否成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteQueryRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteQueryRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IngressHost struct {
	// 域名

	Host *string `json:"Host,omitempty" name:"Host"`
	// 是否开启TLS

	TLS *bool `json:"TLS,omitempty" name:"TLS"`
	// Secret名称

	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
	// 路径列表

	Paths []*IngressPathPlus `json:"Paths,omitempty" name:"Paths"`
}

type AddLabelRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"label"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 操作类型："Op":"add"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 新增标签信息

	AddLabelSet []*AddLabelInfo `json:"AddLabelSet,omitempty" name:"AddLabelSet"`
}

func (r *AddLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateRetentionRequest struct {
	*tchttp.BaseRequest

	// Retention定义

	Command *CreateOrUpdateRetention `json:"Command,omitempty" name:"Command"`
}

func (r *CreateOrUpdateRetentionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateRetentionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNodeRequest struct {
	*tchttp.BaseRequest

	// 待删除结点的节点名

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 待删除结点的集群名，不传默认使用global

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *DeleteNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveUpgradeConfigRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 管控组件升级配置信息

	ControlConfig *ClusterUpgradeConfig `json:"ControlConfig,omitempty" name:"ControlConfig"`
	// Master节点升级配置信息

	MasterConfig *ClusterUpgradeConfig `json:"MasterConfig,omitempty" name:"MasterConfig"`
	// Worker节点升级配置信息

	WorkerConfig *WorkerClusterUpgradeConfig `json:"WorkerConfig,omitempty" name:"WorkerConfig"`
}

func (r *SaveUpgradeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveUpgradeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchNotificationsRequest struct {
	*tchttp.BaseRequest

	// 搜索条件

	Params *SearchNotifications `json:"Params,omitempty" name:"Params"`
}

func (r *SearchNotificationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchNotificationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Job struct {
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 并发数量

	Parallelism *int64 `json:"Parallelism,omitempty" name:"Parallelism"`
	// 重试次数

	BackoffLimit *int64 `json:"BackoffLimit,omitempty" name:"BackoffLimit"`
	// 状态

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 上次更新时间, 可能会为空

	LastUpdateTimestamp *string `json:"LastUpdateTimestamp,omitempty" name:"LastUpdateTimestamp"`
	// 项目显示名称

	ProjectShowName *string `json:"ProjectShowName,omitempty" name:"ProjectShowName"`
}

type MachineDisk struct {
	// 挂载点

	MountPoint *string `json:"MountPoint,omitempty" name:"MountPoint"`
	// 数据盘对应设备

	Device *string `json:"Device,omitempty" name:"Device"`
}

type DeleteJobsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 待删除的Job列表

	NamespacedNames []*NamespacedName `json:"NamespacedNames,omitempty" name:"NamespacedNames"`
}

func (r *DeleteJobsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteJobsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartWorkloadRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 工作负载类型

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
}

func (r *RestartWorkloadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartWorkloadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMonitorMetricsRequest struct {
	*tchttp.BaseRequest

	// 指标维度，cluster, node, pod

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
}

func (r *ListMonitorMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMonitorMetricsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopPodUsage struct {
	// pod名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// cluster

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 主机名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 主机ip

	NodeIP *string `json:"NodeIP,omitempty" name:"NodeIP"`
	// podip

	IP *string `json:"IP,omitempty" name:"IP"`
	// top值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeDeploymentRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// Deployment名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeDeploymentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeploymentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMiddlewareOverviewWithAlertsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMiddlewareOverviewWithAlertsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMiddlewareOverviewWithAlertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchNotificationAlertsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchNotificationAlertsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchNotificationAlertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerSpecialAllocateLanIPExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 请求动作

	Op *string `json:"Op,omitempty" name:"Op"`
	// 物理服务器批量分配相应虚拟内网段IP接口

	ServersSpecialAllocateLanIp []*ServersSpecialAllocateLanIp `json:"ServersSpecialAllocateLanIp,omitempty" name:"ServersSpecialAllocateLanIp"`
}

func (r *ServerSpecialAllocateLanIPExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerSpecialAllocateLanIPExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IngressTLS struct {
	// TLS认证的域名

	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`
	// 密钥名称

	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

type CreateQueryRecordRequest struct {
	*tchttp.BaseRequest

	// 用户id

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 查询名称

	QueryName *string `json:"QueryName,omitempty" name:"QueryName"`
}

func (r *CreateQueryRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateQueryRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Pod状态

		PodCount *StatusCount `json:"PodCount,omitempty" name:"PodCount"`
		// App状态

		AppCount *StatusCount `json:"AppCount,omitempty" name:"AppCount"`
		// 项目状态

		ProjectCount *StatusCount `json:"ProjectCount,omitempty" name:"ProjectCount"`
		// 租户状态

		TenantCount *StatusCount `json:"TenantCount,omitempty" name:"TenantCount"`
		// CPU使用

		CPUUsage *QuotaUsage `json:"CPUUsage,omitempty" name:"CPUUsage"`
		// 内存使用

		MemoryUsage *QuotaUsage `json:"MemoryUsage,omitempty" name:"MemoryUsage"`
		// 磁盘使用

		DiskUsage *QuotaUsage `json:"DiskUsage,omitempty" name:"DiskUsage"`
		// 容器组使用

		PodUsage *QuotaUsage `json:"PodUsage,omitempty" name:"PodUsage"`
		// 节点状态

		NodeReady *NodeCondition `json:"NodeReady,omitempty" name:"NodeReady"`
		// 内存压力

		MemoryPressure *NodeCondition `json:"MemoryPressure,omitempty" name:"MemoryPressure"`
		// 磁盘压力

		DiskPressure *NodeCondition `json:"DiskPressure,omitempty" name:"DiskPressure"`
		// 进程压力

		PIDPressure *NodeCondition `json:"PIDPressure,omitempty" name:"PIDPressure"`
		// 网络是否不可用

		NetworkUnavailable *NodeCondition `json:"NetworkUnavailable,omitempty" name:"NetworkUnavailable"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNodeStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRelocateServerHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前返回数

		ReturnNum *int64 `json:"ReturnNum,omitempty" name:"ReturnNum"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 历史搬迁服务器查询结果

		HistoryRelocateServerSet *string `json:"HistoryRelocateServerSet,omitempty" name:"HistoryRelocateServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRelocateServerHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRelocateServerHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHPARequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 新的HPA, 需要包含全量信息

	HPA *HPA `json:"HPA,omitempty" name:"HPA"`
}

func (r *ModifyHPARequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHPARequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterTargetsRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 待绑定的后端服务列表

	Targets []*BackendTarget `json:"Targets,omitempty" name:"Targets"`
}

func (r *RegisterTargetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterTargetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashFolderQueryParam struct {
	// 查询字符串

	Query []*LabelPair `json:"Query,omitempty" name:"Query"`
	// 页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 第几页

	Page *int64 `json:"Page,omitempty" name:"Page"`
}

type TreeNodeContainerComponent struct {
	// 该组件下的Pod列表

	Pods *TreeNodePod `json:"Pods,omitempty" name:"Pods"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeNodesOverviewWithAlertsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNodesOverviewWithAlertsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodesOverviewWithAlertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListIngressClassesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 搜索参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListIngressClassesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListIngressClassesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeBucket struct {
	// 分桶字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// 前端指标展示

	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type MetaResourceType struct {
	// 租户ID，对应AppId，当为*时，代表所有租户通用

	Organization *string `json:"Organization,omitempty" name:"Organization"`
	// 资源对象类型名称，在租户和资源组下唯一

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资源对象类型别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 生效资源对象，支持通配，如{tcs_product="*"}

	Selector *string `json:"Selector,omitempty" name:"Selector"`
}

type CleanerScriptRequest struct {
	*tchttp.BaseRequest
}

func (r *CleanerScriptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CleanerScriptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageBuildingTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageBuildingTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageBuildingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNodeInstantMetricDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNodeInstantMetricDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNodeInstantMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenameServersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenameServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServicePort struct {
	// 服务端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 协议，支持TCP, UDP, SCTP

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 目标容器端口

	TargetPort *string `json:"TargetPort,omitempty" name:"TargetPort"`
	// 节点端口

	NodePort *uint64 `json:"NodePort,omitempty" name:"NodePort"`
}

type DeregisterTargetsRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 待解绑的后端服务列表

	Targets []*BackendTarget `json:"Targets,omitempty" name:"Targets"`
}

func (r *DeregisterTargetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeregisterTargetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQueryRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 查询记录列表

		Records []*QueryRecord `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQueryRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQueryRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DriveDetailByClusterMetricRequest struct {
	*tchttp.BaseRequest

	// 集群id

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 下钻集群指标

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 下钻时间点，毫秒的时间戳

	DriveTimestamp *string `json:"DriveTimestamp,omitempty" name:"DriveTimestamp"`
	// 偏移量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 返回的符合条件的节点数量，默认10

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DriveDetailByClusterMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DriveDetailByClusterMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleServerSpecialWanIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 回收服务器特殊外网IP信息

		SvrSpecRecyWanIPSet []*SvrSpecRecyWanIP `json:"SvrSpecRecyWanIPSet,omitempty" name:"SvrSpecRecyWanIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecycleServerSpecialWanIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleServerSpecialWanIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JoinToBanksServers struct {
	// 内网ip

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
	// 服务器类型； "0"物理服务器，“3”黑石服务器

	SvrType *string `json:"SvrType,omitempty" name:"SvrType"`
}

type GetUpgradeCheckResultRequest struct {
	*tchttp.BaseRequest

	// 升级批次号

	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

func (r *GetUpgradeCheckResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeCheckResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListVolumesNodeViewRequest struct {
	*tchttp.BaseRequest

	// kubernetes集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 按某个字段排序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否降序排序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 搜索条件

	SearchFilters []*SearchFilter `json:"SearchFilters,omitempty" name:"SearchFilters"`
}

func (r *ListVolumesNodeViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListVolumesNodeViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollbackWorkloadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RollbackWorkloadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RollbackWorkloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DAG struct {
	// 步骤列表

	Steps []*DAGStep `json:"Steps,omitempty" name:"Steps"`
	// 边列表

	Edges []*DAGStep `json:"Edges,omitempty" name:"Edges"`
}

type ModifyLabel struct {
	// 修改条件

	Condition *LabelCondition `json:"Condition,omitempty" name:"Condition"`
	// 修改具体字段内容

	Modify *LabelModify `json:"Modify,omitempty" name:"Modify"`
}

type DescribeServersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前返回数

		ReturnNum *int64 `json:"ReturnNum,omitempty" name:"ReturnNum"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 查询服务器结果信息

		ServerSet []*DescribeServerInfo `json:"ServerSet,omitempty" name:"ServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeProgressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUpgradeProgressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListVirtualMachinesRequest struct {
	*tchttp.BaseRequest

	// 集群ids

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
	// az

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// orderby

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// desc

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
}

func (r *ListVirtualMachinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListVirtualMachinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListNativeStorageClassesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// StorageClass数组

		StorageClassSet []*NativeStorageClass `json:"StorageClassSet,omitempty" name:"StorageClassSet"`
		// 可搜索的CSI列表

		CSISet []*string `json:"CSISet,omitempty" name:"CSISet"`
		// 可搜索的存储类型列表

		StorageTypeSet []*string `json:"StorageTypeSet,omitempty" name:"StorageTypeSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListNativeStorageClassesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListNativeStorageClassesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HPAMetricSpec struct {
	// 指标名

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 支持多种格式, 例如"1Gi", "3", "0.01". 支持的单位有Ki/Mi/Gi/m/""/k/M/G等. 其中1000m=1, 1Ki=1024, 1K=1000

	AverageValue *string `json:"AverageValue,omitempty" name:"AverageValue"`
}

type DescribeAllVlanIdsExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 条数限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 请求动作

	Op *string `json:"Op,omitempty" name:"Op"`
	// 查询条件

	FieldFilters *string `json:"FieldFilters,omitempty" name:"FieldFilters"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAllVlanIdsExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllVlanIdsExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSilenceRequest struct {
	*tchttp.BaseRequest

	// 屏蔽策略 id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetSilenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSilenceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDashboardPermissionsByIdRequest struct {
	*tchttp.BaseRequest

	// 仪表盘ID

	DashId *string `json:"DashId,omitempty" name:"DashId"`
	// 仪表盘权限列表

	Command *UpdateDashAclCommand `json:"Command,omitempty" name:"Command"`
}

func (r *UpdateDashboardPermissionsByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDashboardPermissionsByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NotificationSpec struct {
	// 通知种类，如npd

	Exporter *string `json:"Exporter,omitempty" name:"Exporter"`
	// 通知开关，on/off

	Switch *bool `json:"Switch,omitempty" name:"Switch"`
	// 订阅用户列表

	Subscribers []*string `json:"Subscribers,omitempty" name:"Subscribers"`
	// 发送方式列表

	Channels []*string `json:"Channels,omitempty" name:"Channels"`
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeploymentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDeploymentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDeploymentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteJobsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteJobsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群信息

		Data *DescribeCluster `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashboardContentByUidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 仪表盘内容

		Data *DashboardContent `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDashboardContentByUidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardContentByUidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashboardVersionContentByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDashboardVersionContentByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardVersionContentByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeHistoryListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 升级历史列表

		Data []*UpgradeHistoryListRecord `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUpgradeHistoryListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeHistoryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListConfigMapsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 筛选条件, 支持查询的字段: [Name]

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段, 支持的字段: [Name]

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否降序排序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
}

func (r *ListConfigMapsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListConfigMapsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddResourceObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源对象

		Data *ResourceObject `json:"Data,omitempty" name:"Data"`
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddResourceObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddResourceObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDaemonSetRequest struct {
	*tchttp.BaseRequest

	// daemonset的json字符串

	DaemonSetJSON *string `json:"DaemonSetJSON,omitempty" name:"DaemonSetJSON"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *CreateDaemonSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDaemonSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServerInfo struct {
	// 详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 服务器查询信息

	ResultInfo *ServerResultInfo `json:"ResultInfo,omitempty" name:"ResultInfo"`
}

type MetaMetric struct {
	// 指标名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 指标Label列表

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 说明

	Help *string `json:"Help,omitempty" name:"Help"`
	// 指纹ID

	Fingerprint *uint64 `json:"Fingerprint,omitempty" name:"Fingerprint"`
}

type AddOutbandStrategyData struct {
	// 自定义带外策略具体内容

	Strategies []*AddOutbandStrategies `json:"Strategies,omitempty" name:"Strategies"`
}

type ClusterVersion struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 类型,  tcs or tsf or imported

	VersionType *string `json:"VersionType,omitempty" name:"VersionType"`
}

type CreateLoadBalancerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 负载均衡ID

		LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLoadBalancerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardMeta struct {
	// 是否可管理

	CanAdmin *bool `json:"CanAdmin,omitempty" name:"CanAdmin"`
	// 是否可编辑

	CanEdit *bool `json:"CanEdit,omitempty" name:"CanEdit"`
	// 是否可保存

	CanSave *bool `json:"CanSave,omitempty" name:"CanSave"`
	// 是否可加星

	CanStar *bool `json:"CanStar,omitempty" name:"CanStar"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 创建用户

	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`
	// 过期时间

	Expires *string `json:"Expires,omitempty" name:"Expires"`
	// 所在目录ID

	FolderId *int64 `json:"FolderId,omitempty" name:"FolderId"`
	// 所在目录名称

	FolderTitle *string `json:"FolderTitle,omitempty" name:"FolderTitle"`
	// 是否启用ACL

	HasAcl *bool `json:"HasAcl,omitempty" name:"HasAcl"`
	// 更新时间

	Updated *string `json:"Updated,omitempty" name:"Updated"`
	// 更新用户

	UpdatedBy *string `json:"UpdatedBy,omitempty" name:"UpdatedBy"`
	// 版本号

	Version *int64 `json:"Version,omitempty" name:"Version"`
}

type DashboardVersion struct {
	// 版本ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 仪表盘ID

	DashboardId *int64 `json:"DashboardId,omitempty" name:"DashboardId"`
	// 父版本号

	ParentVersion *int64 `json:"ParentVersion,omitempty" name:"ParentVersion"`
	// 恢复来源版本

	RestoredFrom *int64 `json:"RestoredFrom,omitempty" name:"RestoredFrom"`
	// 版本

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 创建人

	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`
	// 更新消息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type DeleteListenerRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 要删除的监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

func (r *DeleteListenerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteListenerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WithdrawServerExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// WithdrawServerSet

		WithdrawServerSet []*WithdrawServerSet `json:"WithdrawServerSet,omitempty" name:"WithdrawServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WithdrawServerExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *WithdrawServerExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMatrixResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMatrixResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMatrixResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeProgressRequest struct {
	*tchttp.BaseRequest

	// 升级批次号

	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

func (r *GetUpgradeProgressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeProgressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListIngressClassesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 路由控制器列表

		IngressClassSet []*IngressClass `json:"IngressClassSet,omitempty" name:"IngressClassSet"`
		// 过滤参数列表

		FilterEntity *IngressClassFilterEntity `json:"FilterEntity,omitempty" name:"FilterEntity"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListIngressClassesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListIngressClassesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMachineVersionsRequest struct {
	*tchttp.BaseRequest

	// 过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListMachineVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMachineVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetPortRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 要修改端口的后端服务列表

	Targets []*BackendTarget `json:"Targets,omitempty" name:"Targets"`
	// 后端服务绑定到监听器或转发规则的新端口

	NewPort *int64 `json:"NewPort,omitempty" name:"NewPort"`
}

func (r *ModifyTargetPortRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTargetPortRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApiServerMetricsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiServerMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiServerMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Ndb struct {
	// 最大cpu

	MaxCpuRatio *int64 `json:"MaxCpuRatio,omitempty" name:"MaxCpuRatio"`
	// 最大内存

	MaxMemRatio *int64 `json:"MaxMemRatio,omitempty" name:"MaxMemRatio"`
	// 节点在线率

	MinNodeDisable *int64 `json:"MinNodeDisable,omitempty" name:"MinNodeDisable"`
}

type GetMetaRequest struct {
	*tchttp.BaseRequest

	// Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 模糊查询MetaName

	MetaName *string `json:"MetaName,omitempty" name:"MetaName"`
	// 过滤类型，默认0 不过滤， 1 只展示已接入日志配置的meta

	FilterType *uint64 `json:"FilterType,omitempty" name:"FilterType"`
}

func (r *GetMetaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMetaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogConfigRequest struct {
	*tchttp.BaseRequest

	// 配置id

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 日志路径，支持通配

	Path *string `json:"Path,omitempty" name:"Path"`
	// 当path包含通配符时，该字段用于排除不必要的日志文件

	PathExcluded *string `json:"PathExcluded,omitempty" name:"PathExcluded"`
	// 是否启用多行匹配

	MultipleLine *bool `json:"MultipleLine,omitempty" name:"MultipleLine"`
	// 指定系统已有的parse

	ParserRef *string `json:"ParserRef,omitempty" name:"ParserRef"`
	// 用于手工指定配置，可用配置有json、regex等

	ParserFormat *string `json:"ParserFormat,omitempty" name:"ParserFormat"`
	// time字段名称

	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`
	// time格式

	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`
	// 提取正则

	Regex *string `json:"Regex,omitempty" name:"Regex"`
	// 行首正则

	LineHeaderRegex *string `json:"LineHeaderRegex,omitempty" name:"LineHeaderRegex"`
	// 当ParserFormat为LTSV时需要指定

	Type *string `json:"Type,omitempty" name:"Type"`
	// 用于手工指定配置，主要用于对特定定段再次进行解码

	Decoders []*string `json:"Decoders,omitempty" name:"Decoders"`
	// 是否提取字段

	FieldExtraction *bool `json:"FieldExtraction,omitempty" name:"FieldExtraction"`
	// 是否使用系统时间

	UseSystemTime *bool `json:"UseSystemTime,omitempty" name:"UseSystemTime"`
	// 日志名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志样例

	Example *string `json:"Example,omitempty" name:"Example"`
	// 字段信息，json  样例:[{"Field":"cost","Type":"int"},{"Field":"log_level","Type":"string"},{"Field":"old_price","Type":"double"}]

	FieldMapping *string `json:"FieldMapping,omitempty" name:"FieldMapping"`
}

func (r *ModifyLogConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleVMLanIP struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器序列号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 服务器内网IP

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
}

type TopoService struct {
	// 标识

	Id *string `json:"Id,omitempty" name:"Id"`
	// 服务名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 标签列表

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
	// DNS名称

	DNSNames []*string `json:"DNSNames,omitempty" name:"DNSNames"`
	// 策略状态

	VisibilityPolicyStatus []*string `json:"VisibilityPolicyStatus,omitempty" name:"VisibilityPolicyStatus"`
	// 创建时间

	CreationTimstamp *string `json:"CreationTimstamp,omitempty" name:"CreationTimstamp"`
}

type DescribeImagePathRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImagePathRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagePathRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeK8sComponentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeK8sComponentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeK8sComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTcsComponentRequest struct {
	*tchttp.BaseRequest

	// tcs_infra, tcs_alm, tcs_ssm, tcs_pd, tcs_devops

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
	// 集群id

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *DescribeTcsComponentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTcsComponentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNodeLabelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结点标签KV对

		LabelSet []*KVPair `json:"LabelSet,omitempty" name:"LabelSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNodeLabelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNodeLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PerformServerTaskExRequest struct {
	*tchttp.BaseRequest

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// SvrIdcId

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 用户

	User *string `json:"User,omitempty" name:"User"`
	// 操作类型

	OperateType *int64 `json:"OperateType,omitempty" name:"OperateType"`
	// 请求动作

	Command *string `json:"Command,omitempty" name:"Command"`
	// 服务方信息

	Servers []*PerformServerInfo `json:"Servers,omitempty" name:"Servers"`
}

func (r *PerformServerTaskExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PerformServerTaskExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyWsTokenRequest struct {
	*tchttp.BaseRequest

	// 参数列表

	Parameters []*KVPair `json:"Parameters,omitempty" name:"Parameters"`
	// 模块

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *ApplyWsTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyWsTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageRequest struct {
	*tchttp.BaseRequest

	// "Op": "delete_image"

	Op *string `json:"Op,omitempty" name:"Op"`
	// "DeleteImageParams":[{"ImageId": "ins-123456"},{"ImageId": "ins-56789"}

	DeleteImageParams *DeleteImageParams `json:"DeleteImageParams,omitempty" name:"DeleteImageParams"`
}

func (r *DeleteImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSilencesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchSilencesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSilencesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetricData struct {
	// 监控指标

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 监控值

	Value *float64 `json:"Value,omitempty" name:"Value"`
	// 时间截

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 标签

	Labels *string `json:"Labels,omitempty" name:"Labels"`
}

type CreateOrUpdateDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrUpdateDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVMListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server_vm"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"delete"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 删除VM信息

	DeleteVMSet []*VM `json:"DeleteVMSet,omitempty" name:"DeleteVMSet"`
}

func (r *DeleteVMListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVMListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIngressStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群ID

		Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
		// 路由控制器

		IngressClass *string `json:"IngressClass,omitempty" name:"IngressClass"`
		// 网络类型

		NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
		// 域名列表

		Hosts []*IngressHost `json:"Hosts,omitempty" name:"Hosts"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// VIP

		VIP *string `json:"VIP,omitempty" name:"VIP"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIngressStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIngressStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardContent struct {
	// 仪表盘元数据

	Meta *DashboardMeta `json:"Meta,omitempty" name:"Meta"`
	// 仪表盘内容

	Dashboard *DummyContent `json:"Dashboard,omitempty" name:"Dashboard"`
}

type ModifyLabelInfo struct {
	// 要修改的标签id

	Condition *LabelConditionInfo `json:"Condition,omitempty" name:"Condition"`
	// 修改后的标签内容

	Modify *ModifyLabelModifyInfo `json:"Modify,omitempty" name:"Modify"`
}

type CreateIngressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIngressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteStorageClassResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteStorageClassResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteStorageClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVMListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVMListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVMListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDevOpsHistoryRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 返回的符合条件的节点数量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 待搜索参数列表，支持RemediationName, Node, Result

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序字段，支持StartTime,CompletionTime,Node

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 格式：2006-01-02T15:04:05Z

	Start *string `json:"Start,omitempty" name:"Start"`
	// 格式：2006-01-02T15:04:05Z

	End *string `json:"End,omitempty" name:"End"`
	// 自愈动作或者自愈场景关键字

	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *ListDevOpsHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDevOpsHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMetaMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 指标列表

		Data []*MetaMetric `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMetaMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMetaMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeploymentRequest struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 标签

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// 副本数

	Replicas *uint64 `json:"Replicas,omitempty" name:"Replicas"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 部署策略

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// Pod模板

	Template *PodTemplate `json:"Template,omitempty" name:"Template"`
}

type DeleteResourceObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteResourceObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteResourceObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTcsComponentPodResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTcsComponentPodResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTcsComponentPodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeFilterEntity struct {
	// 角色

	Role []*string `json:"Role,omitempty" name:"Role"`
	// 状态

	Phase []*string `json:"Phase,omitempty" name:"Phase"`
}

type CreateLoadBalancerRequest struct {
	*tchttp.BaseRequest

	// 负载均衡实例ID，不填会自动生成

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡实例的名称

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 负载均衡网络类型，Inner内网，Outer外网

	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
	// Tunnel或FullNat

	LoadBalancerMode *string `json:"LoadBalancerMode,omitempty" name:"LoadBalancerMode"`
	// 用途

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// 指定VIP

	VIPs *string `json:"VIPs,omitempty" name:"VIPs"`
	// IP协议

	IPFamilies []*string `json:"IPFamilies,omitempty" name:"IPFamilies"`
}

func (r *CreateLoadBalancerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLoadBalancerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMonitorTemplatesRequest struct {
	*tchttp.BaseRequest

	// 维度取值：cluster, node, pod

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
}

func (r *ListMonitorTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMonitorTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMultiNicDefinitionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询网卡信息出参集合

		MultiNicSet []*MultiNicSet `json:"MultiNicSet,omitempty" name:"MultiNicSet"`
		// 返回数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMultiNicDefinitionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMultiNicDefinitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollingUpdateStatefulSetStrategy struct {
	// 分区

	Partition *uint64 `json:"Partition,omitempty" name:"Partition"`
}

type SearchMetaValuesRequest struct {
	*tchttp.BaseRequest

	// Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 模糊查询的MetaName

	MetaName *string `json:"MetaName,omitempty" name:"MetaName"`
}

func (r *SearchMetaValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchMetaValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenameServerCondition struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// Sn

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 内网IP

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
}

type GetMetricHistoryByDimensionRequest struct {
	*tchttp.BaseRequest

	// 维度，如cluster,node,pod

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
	// {name,value}

	Labels []*LabelPair `json:"Labels,omitempty" name:"Labels"`
	// 需要显示的列指标

	MetricList []*string `json:"MetricList,omitempty" name:"MetricList"`
	// 开始时间

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束时间

	End *string `json:"End,omitempty" name:"End"`
	// 时间间隔，10s,1m

	Step *string `json:"Step,omitempty" name:"Step"`
}

func (r *GetMetricHistoryByDimensionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMetricHistoryByDimensionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 资源类型列表

		Data []*ResourceType `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListResourceTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchAlertsRequest struct {
	*tchttp.BaseRequest

	// 告警搜索条件

	Params *SearchAlert `json:"Params,omitempty" name:"Params"`
}

func (r *SearchAlertsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchAlertsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AggResults struct {
	// 多桶型聚合结果

	Buckets []*Bucket `json:"Buckets,omitempty" name:"Buckets"`
}

type AppInstance struct {
	// 应用名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 应用所在namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 应用所属项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 用户名称

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 组件个数

	TotalComponent *int64 `json:"TotalComponent,omitempty" name:"TotalComponent"`
	// Ready状态的组件个数

	ReadyComponent *int64 `json:"ReadyComponent,omitempty" name:"ReadyComponent"`
	// 状态, Failed/Pending/Succeeded

	Status *string `json:"Status,omitempty" name:"Status"`
	// 部署方式, "直接部署"/"应用市场"

	DeployMethod *string `json:"DeployMethod,omitempty" name:"DeployMethod"`
	// 属于这个应用的Pod数量

	PodCount *int64 `json:"PodCount,omitempty" name:"PodCount"`
	// UUID, 用于跳转到对应应用的应用详情

	ID *string `json:"ID,omitempty" name:"ID"`
	// helm/oam

	Type *string `json:"Type,omitempty" name:"Type"`
}

type MetaResourceCatalogQueryParam struct {
	// 目录类型，支持product和resource

	Type *string `json:"Type,omitempty" name:"Type"`
}

type ModifyDaemonSetRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// DaemonSet的JSON字符串

	DaemonSet *string `json:"DaemonSet,omitempty" name:"DaemonSet"`
	// 描述, 不传表示不修改

	Description *string `json:"Description,omitempty" name:"Description"`
	// 标签, 不传表示不修改, 传空数组表示清空标签

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
}

func (r *ModifyDaemonSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDaemonSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleServerVirtualIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecycleServerVirtualIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleServerVirtualIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoadBalancersRequest struct {
	*tchttp.BaseRequest

	// 数据偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回的实例数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 1：倒序，0：顺序，默认按照创建时间倒序。

	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`
	// 负载均衡实例的 VIP 地址，支持多个

	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`
	// 负载均衡实例ID

	LoadBalancerIds *string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
	// 负载均衡网络类型

	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
	// 搜索字段

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeLoadBalancersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoadBalancersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListClusterVersionsRequest struct {
	*tchttp.BaseRequest

	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListClusterVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListClusterVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomScriptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomScriptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustomScriptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCustomScriptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIngressRequest struct {
	*tchttp.BaseRequest

	// 集群ID, 如"global"

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 待删除Ingress名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 待删除Ingress namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DeleteIngressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIngressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Bucket struct {
	// 分桶key的格式化结果

	KeyAsString *string `json:"KeyAsString,omitempty" name:"KeyAsString"`
	// 分桶key 纳秒时间戳

	Key *uint64 `json:"Key,omitempty" name:"Key"`
	// 桶内数据计数

	DocCount *uint64 `json:"DocCount,omitempty" name:"DocCount"`
}

type DAGEdge struct {
	// 起始节点名字

	From *string `json:"From,omitempty" name:"From"`
	// 目的节点名字

	To *string `json:"To,omitempty" name:"To"`
}

type ClusterUpgradeConfig struct {
	// 集群当前版本

	CurrentVersion *string `json:"CurrentVersion,omitempty" name:"CurrentVersion"`
	// 集群目标版本

	TargetVersion *string `json:"TargetVersion,omitempty" name:"TargetVersion"`
	// 升级结束的观察时间

	ObserveTime *int64 `json:"ObserveTime,omitempty" name:"ObserveTime"`
}

type Condition struct {
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 固资号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// Lanip

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
}

type ServiceBinding struct {
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 服务绑定名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 服务绑定对应实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 服务绑定对应实例的命名空间

	InstanceNamespace *string `json:"InstanceNamespace,omitempty" name:"InstanceNamespace"`
	// 外部访问ID

	ExternalID *string `json:"ExternalID,omitempty" name:"ExternalID"`
	// 密钥名称

	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
	// 所属产品

	Prodcut *string `json:"Prodcut,omitempty" name:"Prodcut"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DeleteHPAResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteHPAResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHPAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Expr struct {
	// metric name, label使用RuleGroup.Objects

	Name *string `json:"Name,omitempty" name:"Name"`
	// expr value

	Val *string `json:"Val,omitempty" name:"Val"`
	// == != > < >= <=  [ metric op val ]

	Op *string `json:"Op,omitempty" name:"Op"`
	// 聚合函数，count_over_time, rate

	AggrFunc *string `json:"AggrFunc,omitempty" name:"AggrFunc"`
	// sum,avg     [ func(metric) op val ]

	Func *string `json:"Func,omitempty" name:"Func"`
	// [1d,1w,1m]

	Range *string `json:"Range,omitempty" name:"Range"`
	// duraton 1d,1m,1w

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// and,or,*,/

	NextRel *string `json:"NextRel,omitempty" name:"NextRel"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// up,down

	Thresh *string `json:"Thresh,omitempty" name:"Thresh"`
	// 场景类型

	SceneType *string `json:"SceneType,omitempty" name:"SceneType"`
}

type ReinstallOsInfo struct {
	// SvrIdcName

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// 机器类型

	SvrDeviceType *string `json:"SvrDeviceType,omitempty" name:"SvrDeviceType"`
	// SN号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// Lan信息

	Lan *Lan `json:"Lan,omitempty" name:"Lan"`
	// Wan信息

	Wan *Wan `json:"Wan,omitempty" name:"Wan"`
	// 操作系统

	Os *string `json:"Os,omitempty" name:"Os"`
	// Raid信息

	Raid *string `json:"Raid,omitempty" name:"Raid"`
	// 固资编号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 1：Bonding 0：非Bonding

	BondingFlag *string `json:"BondingFlag,omitempty" name:"BondingFlag"`
	// UserShellPath

	UserShellPath *string `json:"UserShellPath,omitempty" name:"UserShellPath"`
	// SvrIdcId

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// 是否绑定raid，0:不绑定;1:绑定

	RaidFlag *string `json:"RaidFlag,omitempty" name:"RaidFlag"`
	// 硬件描述信息

	SvrDeviceDescript *string `json:"SvrDeviceDescript,omitempty" name:"SvrDeviceDescript"`
	// 架构

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	// 分组

	Group *string `json:"Group,omitempty" name:"Group"`
	// 是否支持ipv6 yes or no

	EnableIpv6 *string `json:"EnableIpv6,omitempty" name:"EnableIpv6"`
	// 设置后，只装机智能网卡，否则安装智能网卡+黑石主机。

	OnlySoc *bool `json:"OnlySoc,omitempty" name:"OnlySoc"`
	// 服务器类型（3 黑石主机 or 4 黑石智能网卡）

	SvrType *string `json:"SvrType,omitempty" name:"SvrType"`
	// 关联智能网卡和黑石主机。

	SvrRelatedSn *string `json:"SvrRelatedSn,omitempty" name:"SvrRelatedSn"`
	// 是否清理数据盘（0：不清理，8：清理）

	PartitionFlag *int64 `json:"PartitionFlag,omitempty" name:"PartitionFlag"`
}

type AddRAIDResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddRAIDResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddRAIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePersistentVolumeClaimRequest struct {
	*tchttp.BaseRequest

	// kubernetes集群名, 如global

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// StorageClass名字

	StorageClass *string `json:"StorageClass,omitempty" name:"StorageClass"`
	// 指定PV创建PVC, 如果不传则使用StorageClass动态创建PV

	PersistentVolume *string `json:"PersistentVolume,omitempty" name:"PersistentVolume"`
	// PVC所属项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 待创建PVC的名字, 名字需符合kubernetes规范: 长度<=253, [a-zA-Z0-9.-], 数字字母结尾及开头, 且.后面那一段也必须以数字字母结尾及开头, 例如a.a-a合法, 但是a.-a不合法.

	Name *string `json:"Name,omitempty" name:"Name"`
	// 需要申请的存储容量, 例如"100Gi", "500Mi"

	Capacity *string `json:"Capacity,omitempty" name:"Capacity"`
	// 读写模式, 例如["ReadWriteOnce"]

	AccessModes []*string `json:"AccessModes,omitempty" name:"AccessModes"`
	// 默认文件系统形式挂载, 如果要以块设备形式挂载, 设置成"Block"

	VolumeMode *string `json:"VolumeMode,omitempty" name:"VolumeMode"`
	// 标签列表

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// 注解列表

	Annotations []*KVPair `json:"Annotations,omitempty" name:"Annotations"`
	// 描述文本

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreatePersistentVolumeClaimRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePersistentVolumeClaimRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DriveDetailByNodeMetricRequest struct {
	*tchttp.BaseRequest

	// 集群id

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 节点ip

	Node *string `json:"Node,omitempty" name:"Node"`
	// 下钻节点指标

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 毫秒时间戳

	DriveTimestamp *string `json:"DriveTimestamp,omitempty" name:"DriveTimestamp"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制个数，默认10

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DriveDetailByNodeMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DriveDetailByNodeMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云产品状态列表

		ProductStatus []*ProductStatus `json:"ProductStatus,omitempty" name:"ProductStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoUpgradeOperationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoUpgradeOperationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoUpgradeOperationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteStatefulSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteStatefulSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteStatefulSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VolumeNodeView struct {
	// 结点名

	Node *string `json:"Node,omitempty" name:"Node"`
	// local Volume数量

	VolumeCount *int64 `json:"VolumeCount,omitempty" name:"VolumeCount"`
	// 已分配本地盘大小, 单位Byte

	Allocated *int64 `json:"Allocated,omitempty" name:"Allocated"`
	// 总的本地盘大小, 单位Byte

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 本地盘已使用大小, 单位Byte

	Used *int64 `json:"Used,omitempty" name:"Used"`
	// 分配率, 需要乘以100. 例如返回的是0.1, 表示10%

	AllocatedRatio *float64 `json:"AllocatedRatio,omitempty" name:"AllocatedRatio"`
	// 使用率, 需要乘以100. 例如返回的是0.1, 表示10%

	UsedRatio *float64 `json:"UsedRatio,omitempty" name:"UsedRatio"`
}

type ContainerSet struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 重启次数

	RestartCount *int64 `json:"RestartCount,omitempty" name:"RestartCount"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 是否ready

	Ready *bool `json:"Ready,omitempty" name:"Ready"`
	// 容器ID

	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`
	// 镜像

	Image *string `json:"Image,omitempty" name:"Image"`
	// 镜像ID

	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
	// 是否是初始化容器

	Init *bool `json:"Init,omitempty" name:"Init"`
	// 容器端口

	Ports []*ContainerPort `json:"Ports,omitempty" name:"Ports"`
}

type DeleteMultiNicParams struct {
	// 自定义网卡名称

	DefinitionName *string `json:"DefinitionName,omitempty" name:"DefinitionName"`
}

type ListBasePodsRequest struct {
	*tchttp.BaseRequest

	// 列表行数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 查询条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 节点名列表

	NodeNames []*string `json:"NodeNames,omitempty" name:"NodeNames"`
}

func (r *ListBasePodsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListBasePodsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCronJobsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CronJob列表

		CronJobSet []*CronJob `json:"CronJobSet,omitempty" name:"CronJobSet"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 状态列表

		PhaseSet []*string `json:"PhaseSet,omitempty" name:"PhaseSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCronJobsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCronJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkPolicyStruct struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 所属项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 网络策略生效的组件类型，可选 Deployment、StatefulSet、StatefulSetPlus、DaemonSet、Job 和 CronJob。如果为空则表示任何类型。

	ComponentKind *string `json:"ComponentKind,omitempty" name:"ComponentKind"`
	// 网络策略生效的组件名称或者 Any。当 ComponentKind 为空并且 ComponentName 为 Any 时表示对所有组件生效。

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 入方向是否开启

	IngressEnabled *bool `json:"IngressEnabled,omitempty" name:"IngressEnabled"`
	// 入方向规则列表

	IngressRules []*NetworkPolicyRule `json:"IngressRules,omitempty" name:"IngressRules"`
	// 出方向是否开启

	EgressEnabled *bool `json:"EgressEnabled,omitempty" name:"EgressEnabled"`
	// 出方向是否开启

	EgressRules []*NetworkPolicyRule `json:"EgressRules,omitempty" name:"EgressRules"`
	// 项目显示名称

	ProjectShowName *string `json:"ProjectShowName,omitempty" name:"ProjectShowName"`
	// 生效组件标签

	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`
}

type CreateMonitorTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMonitorTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMonitorTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirtualMachineStatusRequest struct {
	*tchttp.BaseRequest

	// 集群

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 虚拟机ID

	VmId *string `json:"VmId,omitempty" name:"VmId"`
}

func (r *DescribeVirtualMachineStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirtualMachineStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPersistentVolumeClaimsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合筛选条件的PVC总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 符合条件的PVC列表

		PersistentVolumeClaimSet []*PVC `json:"PersistentVolumeClaimSet,omitempty" name:"PersistentVolumeClaimSet"`
		// 用于搜索的AccessMode列表, 如["ROX", "RWX"]

		AccessModesSet []*string `json:"AccessModesSet,omitempty" name:"AccessModesSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPersistentVolumeClaimsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPersistentVolumeClaimsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomScriptSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomScriptSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomScriptSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PodTemplate struct {
	// 标签列表

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// 重启策略, Never/Always/OnFailure

	RestartPolicy *string `json:"RestartPolicy,omitempty" name:"RestartPolicy"`
	// 是否host网络模式

	HostNetwork *bool `json:"HostNetwork,omitempty" name:"HostNetwork"`
}

type GetMetricHistoryByDimensionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMetricHistoryByDimensionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMetricHistoryByDimensionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMonitorTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的模板总个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 符合条件的模板列表

		MonitorTemplateSet *ListTemplateRes `json:"MonitorTemplateSet,omitempty" name:"MonitorTemplateSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMonitorTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMonitorTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群信息

		Data *DescribeCluster `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// "ImageSet":[{"Detail": "ok","Result": 0}]

		ImageSet []*string `json:"ImageSet,omitempty" name:"ImageSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateResourceCommand struct {
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type ListImageBuildingTasksRequest struct {
	*tchttp.BaseRequest

	// 查询过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 倒序排序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 查询数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListImageBuildingTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListImageBuildingTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListJobsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// job数组

		JobSet []*Job `json:"JobSet,omitempty" name:"JobSet"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 状态列表

		PhaseSet []*string `json:"PhaseSet,omitempty" name:"PhaseSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListJobsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenameServersRequest struct {
	*tchttp.BaseRequest

	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// "Scheme": "server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 操作类型：modify

	Op *string `json:"Op,omitempty" name:"Op"`
	// 重命名服务器集合

	RenameServerSet []*RenameServer `json:"RenameServerSet,omitempty" name:"RenameServerSet"`
}

func (r *RenameServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenameServersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackendTarget struct {
	// IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 权重

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// 可用区ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 状态，Ready、NotReady

	Phase *string `json:"Phase,omitempty" name:"Phase"`
}

type DescribeTopKPodsRequest struct {
	*tchttp.BaseRequest

	// namespace_pod

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 排序类型: cpu: 按cpu使用率排序, memory: 按内存使用率排序；

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 取Top几

	TopK *int64 `json:"TopK,omitempty" name:"TopK"`
}

func (r *DescribeTopKPodsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopKPodsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetIngressYAMLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Ingress的YAML

		Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetIngressYAMLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetIngressYAMLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashAclItem struct {
	// 用户ID

	UserId *int64 `json:"UserId,omitempty" name:"UserId"`
	// 角色名称，可先值为Viewer、Editor、Admin

	Role *string `json:"Role,omitempty" name:"Role"`
	// 权限，可选值为1,2,4，分别代表View,Edit,Admin

	Permission *int64 `json:"Permission,omitempty" name:"Permission"`
}

type MetaMetricCreateCommand struct {
	// 指标名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 指标Label列表

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 说明

	Help *string `json:"Help,omitempty" name:"Help"`
	// 类型，如counter，gauge，summary，histogram

	Type *string `json:"Type,omitempty" name:"Type"`
}

type QueryServerStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// QueryServerStatusList

		QueryServerStatusList *string `json:"QueryServerStatusList,omitempty" name:"QueryServerStatusList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryServerStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryServerStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 参数名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤值

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 操作符, 不传默认like, 可选equal/gte/lte

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type ListResourceTypeEventRequest struct {
	*tchttp.BaseRequest

	// 资源对象元数据-对象类型名称

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 事件元数据查询参数

	Params *MetaEventQueryParam `json:"Params,omitempty" name:"Params"`
}

func (r *ListResourceTypeEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceTypeEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListVirtualMachineStoragesRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 项目

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 虚拟机id

	VmId *string `json:"VmId,omitempty" name:"VmId"`
	// offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// orderBy

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// Desc

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// filters

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListVirtualMachineStoragesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListVirtualMachineStoragesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportServerList struct {
	// 设备固资编号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器SN

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 设备型号

	SvrDeviceType *string `json:"SvrDeviceType,omitempty" name:"SvrDeviceType"`
	// IDC名称

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// 机架名称

	SvrRackName *string `json:"SvrRackName,omitempty" name:"SvrRackName"`
	// 存放机位ID

	SvrPosId *string `json:"SvrPosId,omitempty" name:"SvrPosId"`
	// 设备高度

	SvrDeviceHeight *string `json:"SvrDeviceHeight,omitempty" name:"SvrDeviceHeight"`
	// 服务器厂商

	SvrProducer *string `json:"SvrProducer,omitempty" name:"SvrProducer"`
	// 逻辑区域

	SvrLogicArea *string `json:"SvrLogicArea,omitempty" name:"SvrLogicArea"`
	// 服务器硬件描述

	SvrDeviceDescript *string `json:"SvrDeviceDescript,omitempty" name:"SvrDeviceDescript"`
	// region name

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// zone name

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// IDC id

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// 服务器架构

	SvrArch *string `json:"SvrArch,omitempty" name:"SvrArch"`
	// 服务器类型（0 物理机 or 3 黑石服务器 or 4 智能网卡）

	SvrType *string `json:"SvrType,omitempty" name:"SvrType"`
	// 关联SN(黑石服务器sn和智能网卡sn互相关联)

	SvrRelatedSn *string `json:"SvrRelatedSn,omitempty" name:"SvrRelatedSn"`
}

type VirtualLanOrWanIP struct {
	// 设备固资编号（实体服务器必须传）

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 虚拟机设备UUID

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 虚拟内网ip

	SvrVirtualLanIp *string `json:"SvrVirtualLanIp,omitempty" name:"SvrVirtualLanIp"`
	// 虚拟外网ip

	SvrVirtualWanIp *string `json:"SvrVirtualWanIp,omitempty" name:"SvrVirtualWanIp"`
}

type ModifyMultiNicParams struct {
	// 修改条件；示例："Condition":{"DefinitionName":"eth0"}

	Condition *ModifyMultiNicCondition `json:"Condition,omitempty" name:"Condition"`
	// 修改内容；注：NicConfig 和 BondConfig修改 需全量修改

	Modify *ModifyMultiNicModify `json:"Modify,omitempty" name:"Modify"`
}

type GetRollbackProgressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRollbackProgressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRollbackProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListBaseClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群列表

		ClusterSet []*BaseCluster `json:"ClusterSet,omitempty" name:"ClusterSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListBaseClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListBaseClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePodsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePodsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePodsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNotificationRequest struct {
	*tchttp.BaseRequest

	// notification id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetNotificationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNotificationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDeploymentRequest struct {
	*tchttp.BaseRequest

	// 集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 待修改的Deployment JSON

	Deployment *string `json:"Deployment,omitempty" name:"Deployment"`
	// 待修改的Service列表

	Services []*string `json:"Services,omitempty" name:"Services"`
	// 待修改的Ingress列表

	Ingresses []*string `json:"Ingresses,omitempty" name:"Ingresses"`
	// 副本数量

	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`
	// 描述, 不传表示不修改

	Description *string `json:"Description,omitempty" name:"Description"`
	// 标签, 不传表示不修改, 传空数组表示清空标签

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ModifyDeploymentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDeploymentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRuleGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchRuleGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRuleGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCronJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CronJob

		CronJob *string `json:"CronJob,omitempty" name:"CronJob"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCronJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCronJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeploymentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 部署详情

		Deployment *string `json:"Deployment,omitempty" name:"Deployment"`
		// Service列表

		Services []*string `json:"Services,omitempty" name:"Services"`
		// Ingress列表

		Ingresses *string `json:"Ingresses,omitempty" name:"Ingresses"`
		// 所属项目

		Project *string `json:"Project,omitempty" name:"Project"`
		// 集群ID

		Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
		// 状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 上次更新时间, 可能为空

		LastUpdateTimestamp *string `json:"LastUpdateTimestamp,omitempty" name:"LastUpdateTimestamp"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeploymentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeploymentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeStatusRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 节点名

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
}

func (r *DescribeNodeStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSvrNicAllocationRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询条件，示例：{"Name":"DefinitionName","Values":[""]}

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeSvrNicAllocationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSvrNicAllocationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomScriptRequest struct {
	*tchttp.BaseRequest

	// 自定义脚本名

	CustomScriptName *string `json:"CustomScriptName,omitempty" name:"CustomScriptName"`
	// 重命名的自定义脚本名

	CustomScriptNameNew *string `json:"CustomScriptNameNew,omitempty" name:"CustomScriptNameNew"`
	// 自定义脚本的内容，base64编码

	CustomScriptContext *string `json:"CustomScriptContext,omitempty" name:"CustomScriptContext"`
}

func (r *ModifyCustomScriptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomScriptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SocPerformServerTaskEsServers struct {
	// Sn

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// idcid

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// 服务器ip

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
	// idcname

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
}

type DeleteLogConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置id

		ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
		// 执行结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashBoardsRequest struct {
	*tchttp.BaseRequest

	// 查询参数

	Params *DashFolderQueryParam `json:"Params,omitempty" name:"Params"`
}

func (r *GetDashBoardsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashBoardsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirtualMachineYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟机YAML，json格式

		VmYaml *string `json:"VmYaml,omitempty" name:"VmYaml"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVirtualMachineYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirtualMachineYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListServicesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 过滤参数，支持Name, Namespace, Type, Selector, ClusterIP, OwnerType, OwnerName

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段，支持Name,Namespace,Type,ClusterIP,CreateTime

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 项目Id

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 集群ID列表

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

func (r *ListServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListStorageTypesRequest struct {
	*tchttp.BaseRequest

	// kubernetes集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *ListStorageTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListStorageTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelocateServerFinishResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RelocateServerFinishResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelocateServerFinishResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品健康列表

		Products []*Product `json:"Products,omitempty" name:"Products"`
		// 模块健康列表

		Type []*Type `json:"Type,omitempty" name:"Type"`
		// 模块总数

		TypeTotalCount *int64 `json:"TypeTotalCount,omitempty" name:"TypeTotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceYAMLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// yaml格式服务详情

		Service *string `json:"Service,omitempty" name:"Service"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceYAMLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceYAMLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlertStats struct {
	// 已确认告警数

	Confirmed *int64 `json:"Confirmed,omitempty" name:"Confirmed"`
	// "紧急"级别告警数

	Severity1 *int64 `json:"Severity1,omitempty" name:"Severity1"`
	// "严重"级别告警数

	Severity2 *int64 `json:"Severity2,omitempty" name:"Severity2"`
	// "一般"级别告警数

	Severity3 *int64 `json:"Severity3,omitempty" name:"Severity3"`
	// "警告"级别告警数

	Severity4 *int64 `json:"Severity4,omitempty" name:"Severity4"`
}

type Rule struct {
	// 策略中的规则名

	Name *string `json:"Name,omitempty" name:"Name"`
	// record or alert

	Type *string `json:"Type,omitempty" name:"Type"`
	// 持续时间

	For *string `json:"For,omitempty" name:"For"`
	// 规则表达式

	Expr *string `json:"Expr,omitempty" name:"Expr"`
	// 自定义label

	Labels []*LabelPair `json:"Labels,omitempty" name:"Labels"`
	// 自定义告警描述

	Annotation []*LabelPair `json:"Annotation,omitempty" name:"Annotation"`
	// 规则Id，前端忽略

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 规则描述，便于用户理解

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 策略id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 内部字段，前端忽略

	PromVal *string `json:"PromVal,omitempty" name:"PromVal"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 更新时间

	Updated *string `json:"Updated,omitempty" name:"Updated"`
	// 以产品场景进行分类，根据不同场景返回进行不同的解析

	ExprArray []*Expr `json:"ExprArray,omitempty" name:"ExprArray"`
	// > >= <= < = !=

	ExprAct *string `json:"ExprAct,omitempty" name:"ExprAct"`
	// 阈值，会有一个默认值，但允许用户改写

	ExprVal *string `json:"ExprVal,omitempty" name:"ExprVal"`
	// 等级

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 规则拆解后的值

	ExprArgs []*LabelPair `json:"ExprArgs,omitempty" name:"ExprArgs"`
}

type DescribeDiskFieldsEnumExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 获取硬盘各字段枚举值出参

		FieldToEnum *DiskFieldToEnum `json:"FieldToEnum,omitempty" name:"FieldToEnum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskFieldsEnumExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDiskFieldsEnumExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleVMWanIP struct {
	// 服务器资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器SN

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 服务器内网IP

	SvrWanIp *string `json:"SvrWanIp,omitempty" name:"SvrWanIp"`
}

type MonitorTemplateSpec struct {
	// 主账号

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 子账号

	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`
	// 模板名称

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 模板类型，cluster, node, pod

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
	// 模板描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 模板监控组具体定义

	Content []*MonitorTemplateContent `json:"Content,omitempty" name:"Content"`
}

type NetworkPolicyRule struct {
	// 端口，为空则是允许所有协议与端口

	Ports []*NetworkPolicyPort `json:"Ports,omitempty" name:"Ports"`
	// 网络策略对端，为空则允许所有来源或目标端

	Peers []*NetworkPolicyPeer `json:"Peers,omitempty" name:"Peers"`
	// 全部

	All *bool `json:"All,omitempty" name:"All"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest

	// 集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// k8s版本

	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`
	// 资源模式，共享Sharing，专属Exclusive

	ResourceMode *string `json:"ResourceMode,omitempty" name:"ResourceMode"`
	// 租户列表，专属集群需要设置

	Tenants *string `json:"Tenants,omitempty" name:"Tenants"`
	// 删除保护

	DeleteProtection *bool `json:"DeleteProtection,omitempty" name:"DeleteProtection"`
	// Cluster CIDR

	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
	// 最大Pod数

	MaxNodePodNum *uint64 `json:"MaxNodePodNum,omitempty" name:"MaxNodePodNum"`
	// 最大Service数

	MaxClusterServiceNum *int64 `json:"MaxClusterServiceNum,omitempty" name:"MaxClusterServiceNum"`
	// 主节点列表（建议节点数>=3）

	Machines []*ClusterMachine `json:"Machines,omitempty" name:"Machines"`
	// 负载均衡

	LoadBalancer *LoadBalancer `json:"LoadBalancer,omitempty" name:"LoadBalancer"`
	// 集群版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
}

func (r *CreateClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DriveDetailByNodeMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// pod列表

		Pods []*DrivePod `json:"Pods,omitempty" name:"Pods"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DriveDetailByNodeMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DriveDetailByNodeMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchNotificationMsgsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchNotificationMsgsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchNotificationMsgsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePodsOverviewWithAlertsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePodsOverviewWithAlertsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePodsOverviewWithAlertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageBuildingTaskRequest struct {
	*tchttp.BaseRequest

	// 镜像构建任务.

	ImageBuildingTask *ImageBuildingTask `json:"ImageBuildingTask,omitempty" name:"ImageBuildingTask"`
}

func (r *ModifyImageBuildingTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageBuildingTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDeploymentsRequest struct {
	*tchttp.BaseRequest

	// 集群ID列表

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索参数，名称、命名空间、项目

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
}

func (r *ListDeploymentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDeploymentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkloadsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 组件列表

		WorkloadSet []*Workload `json:"WorkloadSet,omitempty" name:"WorkloadSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListWorkloadsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListWorkloadsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOSParam struct {
	// 系统名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type Machine struct {
	// 待添加物理机SSH端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 指定添加到某个集群，不传默认使用global集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 验证方式，当前只支持密码方式，可以不传

	Auth *string `json:"Auth,omitempty" name:"Auth"`
	// ssh的用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// ssh密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 主机描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 标签数组

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// 待添加物理机的IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 是否在线检查，默认在线检查

	OnlineCheck *bool `json:"OnlineCheck,omitempty" name:"OnlineCheck"`
	// 节点版本

	MachineVersion *string `json:"MachineVersion,omitempty" name:"MachineVersion"`
	// 数据盘挂载，数据盘挂载，比如 {"MountPoint": "/data", "Device": "/dev/vdb"}

	Disks []*MachineDisk `json:"Disks,omitempty" name:"Disks"`
}

type DeleteMonitorTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板id

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteMonitorTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMonitorTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIngressStatRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 路由控制器名称，若为空，则查询默认路由控制器

	IngressClass *string `json:"IngressClass,omitempty" name:"IngressClass"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 搜索列表

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeIngressStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIngressStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EvictNodeRequest struct {
	*tchttp.BaseRequest

	// 待驱逐结点名

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 待驱逐结点所在集群名，不传默认使用global

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *EvictNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EvictNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMultiNicDefinitionRequest struct {
	*tchttp.BaseRequest

	// 修改网卡信息入参

	ModifyMultiNicParams *ModifyMultiNicParams `json:"ModifyMultiNicParams,omitempty" name:"ModifyMultiNicParams"`
}

func (r *ModifyMultiNicDefinitionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMultiNicDefinitionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOutBandIpResourceIpDetailExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询带外ip资源详细信息出参集合

		IpsDetailSet []*IpsDetailSet `json:"IpsDetailSet,omitempty" name:"IpsDetailSet"`
		// 返回总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOutBandIpResourceIpDetailExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOutBandIpResourceIpDetailExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyServer struct {
	// 修改条件

	Condition *ModifyServerCondition `json:"Condition,omitempty" name:"Condition"`
	// 修改具体字段内容

	Modify *ModifyServerPara `json:"Modify,omitempty" name:"Modify"`
}

type NodeInfo struct {
	// 内核版本

	KernelVersion *string `json:"KernelVersion,omitempty" name:"KernelVersion"`
	// 系统镜像

	OSImage *string `json:"OSImage,omitempty" name:"OSImage"`
	// 容器版本

	ContainerRuntimeVersion *string `json:"ContainerRuntimeVersion,omitempty" name:"ContainerRuntimeVersion"`
	// Kubelet版本

	KubeletVersion *string `json:"KubeletVersion,omitempty" name:"KubeletVersion"`
	// KubeProxy版本

	KubeProxyVersion *string `json:"KubeProxyVersion,omitempty" name:"KubeProxyVersion"`
	// 操作系统

	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`
	// 系统架构

	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
}

type DeleteDashFolderByIdRequest struct {
	*tchttp.BaseRequest

	// 仪表盘目录唯一ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteDashFolderByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashFolderByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListHistoryRevisionsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// workload类型, Deployment/StatefulSet/DaemonSet

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
}

func (r *ListHistoryRevisionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListHistoryRevisionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的节点数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 符合条件的节点列表

		NodeSet []*Node `json:"NodeSet,omitempty" name:"NodeSet"`
		// 筛选项

		FilterEntity *NodeFilterEntity `json:"FilterEntity,omitempty" name:"FilterEntity"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DummyContent struct {
	// 标题名称

	Title *string `json:"Title,omitempty" name:"Title"`
}

type GetAlertStatsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各级别告警条数

		Severities *AlertStats `json:"Severities,omitempty" name:"Severities"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAlertStatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelocateServerStartRequest struct {
	*tchttp.BaseRequest

	// Scheme

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"relocate_start"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 发起搬迁信息

	RelocateStartServers []*RelocateStartServer `json:"RelocateStartServers,omitempty" name:"RelocateStartServers"`
}

func (r *RelocateServerStartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelocateServerStartRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRouteRequest struct {
	*tchttp.BaseRequest

	// 路由信息struct

	Command *CreateRoute `json:"Command,omitempty" name:"Command"`
}

func (r *CreateRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopPodsByMetricRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// cpu或者mem

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// top几，最多1000

	TopK *int64 `json:"TopK,omitempty" name:"TopK"`
}

func (r *DescribeTopPodsByMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopPodsByMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTcsMonitorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTcsMonitorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTcsMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOutbandStrategyRequest struct {
	*tchttp.BaseRequest

	// 请求串类型说明。
	// 固定值： "Scheme":"outband_passwd_strategy"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方systemid	此ID为请求方向cmdb注册的权限ID，由cmdb系统管理员提供和管理。
	// 默认值："SystemId":"sid_test"

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方systemkey	和上面systemid类似，由cmdb系统管理员提供和管理，cmdb侧会根据id和key以及调用方ip三者共同确定请求方的合法性。
	// 默认值："SystemKey":"sid_test"

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 本次请求动作的起始索引地址	不填该字段，系统默认值为0，适合前台翻页查询等场景

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 本次请求动作的数据返回值数量。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询条件。
	// 默认："result_item":"id, strategy_name, have_number,have_lower_letter,have_upper_letter,special_characters,passwd_length,enable"

	FieldFilters *string `json:"FieldFilters,omitempty" name:"FieldFilters"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryOutbandStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryOutbandStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchSilencesRequest struct {
	*tchttp.BaseRequest

	// 搜索条件

	Params *SearchSilences `json:"Params,omitempty" name:"Params"`
}

func (r *SearchSilencesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchSilencesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetUnschedulableRequest struct {
	*tchttp.BaseRequest

	// 待封锁结点所在集群，不传默认使用global

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 待封锁节点名

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 是否封锁，true表示封锁结点，false表示取消封锁

	Unschedulable *bool `json:"Unschedulable,omitempty" name:"Unschedulable"`
	// 封锁或者取消封锁的原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 待封锁节点名列表

	NodeNames []*string `json:"NodeNames,omitempty" name:"NodeNames"`
}

func (r *SetUnschedulableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetUnschedulableRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNodeLabelsRequest struct {
	*tchttp.BaseRequest

	// 待更新标签的节点名字

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 待更新标签的节点所在集群，不传默认使用global

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 新标签全量，结点的原标签不会保留

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
}

func (r *UpdateNodeLabelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNodeLabelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHPAResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHPAResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHPAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFieldsEnumExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// FieldToEnum

		FieldToEnum *Field2enum `json:"FieldToEnum,omitempty" name:"FieldToEnum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFieldsEnumExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFieldsEnumExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlertsTrend struct {
	// 时间字符串，2020-6-1

	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
	// 已恢复个数

	ResolvedCount *int64 `json:"ResolvedCount,omitempty" name:"ResolvedCount"`
	// 未恢复个数

	FiringCount *int64 `json:"FiringCount,omitempty" name:"FiringCount"`
}

type DeleteServer struct {
	// 物理服务器固资号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type ImageSet struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// ImageInfo

	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`
}

type IngressClassFilterEntity struct {
	// 网络类型列表，inner内网，outer外网

	NetworkTypeSet []*string `json:"NetworkTypeSet,omitempty" name:"NetworkTypeSet"`
	// VIP列表

	VIPSet []*string `json:"VIPSet,omitempty" name:"VIPSet"`
}

type RelocateStartServer struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// IDC名

	SvrRelocateIdcName *string `json:"SvrRelocateIdcName,omitempty" name:"SvrRelocateIdcName"`
	// 机架名

	SvrRelocateRackName *string `json:"SvrRelocateRackName,omitempty" name:"SvrRelocateRackName"`
	// 机架位置

	SvrRelocatePosId *string `json:"SvrRelocatePosId,omitempty" name:"SvrRelocatePosId"`
	// 逻辑区

	SvrRelocateLogicArea *string `json:"SvrRelocateLogicArea,omitempty" name:"SvrRelocateLogicArea"`
}

type DescribeDcosFunctionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// dcos支持功能列表信息

		SupportedFunction *SupportedFunction `json:"SupportedFunction,omitempty" name:"SupportedFunction"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDcosFunctionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDcosFunctionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServersSpecialAllocateLanIp struct {
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// SvrIpWanted

	SvrIpWanted *string `json:"SvrIpWanted,omitempty" name:"SvrIpWanted"`
	// Vlan

	Vlan *string `json:"Vlan,omitempty" name:"Vlan"`
}

type DeleteStatefulSetRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间及名字的列表

	NamespacedNames []*NamespacedName `json:"NamespacedNames,omitempty" name:"NamespacedNames"`
}

func (r *DeleteStatefulSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteStatefulSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTcsComponentPodRequest struct {
	*tchttp.BaseRequest

	// 组件名字

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 组件类型

	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`
}

func (r *DescribeTcsComponentPodRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTcsComponentPodRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NamespacedName struct {
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 资源名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

type CustomParameter struct {
	// 参数名字

	Key *string `json:"Key,omitempty" name:"Key"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 参数是否是可选的

	Optional *bool `json:"Optional,omitempty" name:"Optional"`
	// 参数描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type PerformServerInfo struct {
	// 固资号

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// idcid

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// Lanip

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
	// idcName

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
}

type SvrReserveOutBandIPSet struct {
	// 返回结果详细说明

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 返回结果值，0成功， 1失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 预留\取消预留带外ip

	SvrOutBandIp *string `json:"SvrOutBandIp,omitempty" name:"SvrOutBandIp"`
	// 服务器SN

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 服务器固资号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type CreateDashFolderRequest struct {
	*tchttp.BaseRequest

	// 创建仪表盘目录内容

	Command *CreateDashFolderCommand `json:"Command,omitempty" name:"Command"`
}

func (r *CreateDashFolderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDashFolderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListServiceBindingsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 服务绑定列表

		ServiceBindings []*ServiceBinding `json:"ServiceBindings,omitempty" name:"ServiceBindings"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListServiceBindingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListServiceBindingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetricSpec struct {
	// tab名称

	MetricType *string `json:"MetricType,omitempty" name:"MetricType"`
	// Metrics

	Metrics []*MetricDefine `json:"Metrics,omitempty" name:"Metrics"`
}

type VMRecyLanIP struct {
	// 结果详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 服务器内网IP

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
	// 服务器序列号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
}

type AllocateServerWanIPListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"allocate_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 分配服务器外网IP

	AllocateServerWanIPSet []*AllocateServerWanIP `json:"AllocateServerWanIPSet,omitempty" name:"AllocateServerWanIPSet"`
}

func (r *AllocateServerWanIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateServerWanIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBasicOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBasicOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomScriptSetsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li> CustomScriptSetName - String List - 是否必填：否 - （过滤条件）通过脚本集名字过滤。</li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCustomScriptSetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomScriptSetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WithdrawServerSet struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *string `json:"Result,omitempty" name:"Result"`
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type ListResourceObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 服务树资源对象

		Data []*ResourceObject `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListResourceObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSvrNicAllocationRequest struct {
	*tchttp.BaseRequest

	// 绑定(更新)服务器多网卡策略 入参信息

	UpdateSvrNicParams *UpdateSvrNicParams `json:"UpdateSvrNicParams,omitempty" name:"UpdateSvrNicParams"`
}

func (r *UpdateSvrNicAllocationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSvrNicAllocationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CronJob struct {
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 标签

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// 调度策略

	Schedule *string `json:"Schedule,omitempty" name:"Schedule"`
	// 是否允许并发执行. Allow/Forbid/Replace

	ConcurrencyPolicy *string `json:"ConcurrencyPolicy,omitempty" name:"ConcurrencyPolicy"`
	// job历史最大数量限制

	SuccessfulJobsHistoryLimit *int64 `json:"SuccessfulJobsHistoryLimit,omitempty" name:"SuccessfulJobsHistoryLimit"`
	// 失败的job历史最大数量限制

	FailedJobsHistoryLimit *int64 `json:"FailedJobsHistoryLimit,omitempty" name:"FailedJobsHistoryLimit"`
	// 运行中job数量

	RunningJobs *int64 `json:"RunningJobs,omitempty" name:"RunningJobs"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 状态

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 并行pod数量

	Parallelism *int64 `json:"Parallelism,omitempty" name:"Parallelism"`
	// 上次更新时间

	LastUpdateTimestamp *string `json:"LastUpdateTimestamp,omitempty" name:"LastUpdateTimestamp"`
	// 项目显示名称

	ProjectShowName *string `json:"ProjectShowName,omitempty" name:"ProjectShowName"`
}

type DeleteDashFolderByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDashFolderByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashFolderByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TypedLocalObjectReference struct {
	// API组

	APIGroup *string `json:"APIGroup,omitempty" name:"APIGroup"`
	// 类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 资源名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type ListJobsRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段, 支持Name, BackoffLimit, CreationTimestamp

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// owner集群

	OwnerCluster *string `json:"OwnerCluster,omitempty" name:"OwnerCluster"`
	// owner命名空间

	OwnerNamespace *string `json:"OwnerNamespace,omitempty" name:"OwnerNamespace"`
	// owner名字

	OwnerName *string `json:"OwnerName,omitempty" name:"OwnerName"`
	// owner类型, 如CronJob

	OwnerKind *string `json:"OwnerKind,omitempty" name:"OwnerKind"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListJobsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListJobsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyLogConfigRequest struct {
	*tchttp.BaseRequest

	// 待校验字段(TimeFormat,LineHeaderRegex,Regex)

	Field *string `json:"Field,omitempty" name:"Field"`
	// 待校验字段值

	FieldValue *string `json:"FieldValue,omitempty" name:"FieldValue"`
	// 待校验字段规则

	FieldRule *string `json:"FieldRule,omitempty" name:"FieldRule"`
}

func (r *VerifyLogConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyLogConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDevOpsScenesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自愈场景列表

		Items []*DevOpsScenes `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListDevOpsScenesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDevOpsScenesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyJobRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 更新后Job的JSON字符串

	Job *string `json:"Job,omitempty" name:"Job"`
	// 描述, 不传表示不修改

	Description *string `json:"Description,omitempty" name:"Description"`
	// 标签, 不传表示不修改, 传空数组表示清空标签

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// 并行数

	Parallelism *int64 `json:"Parallelism,omitempty" name:"Parallelism"`
	// 重试次数

	BackoffLimit *int64 `json:"BackoffLimit,omitempty" name:"BackoffLimit"`
}

func (r *ModifyJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleVMVirtualIPRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server_vm"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 虚拟内网："Op":"recycle_vitrual_lan_ip"     虚拟外网： "Op":"recycle_vitrual_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 回收VM虚拟内网外网信息集合

	RecycleVMVirtualIPSet []*VirtualLanOrWanIP `json:"RecycleVMVirtualIPSet,omitempty" name:"RecycleVMVirtualIPSet"`
}

func (r *RecycleVMVirtualIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleVMVirtualIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchDashFoldersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仪表盘目录检索结果

		Data []*DashFolderQueryResult `json:"Data,omitempty" name:"Data"`
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchDashFoldersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchDashFoldersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Aggregations struct {
	// 聚合名称

	AggName *string `json:"AggName,omitempty" name:"AggName"`
	// 聚合结果

	AggResults *AggResults `json:"AggResults,omitempty" name:"AggResults"`
}

type GetRelatedAlertNames struct {
	// 待查询事件名列表

	EventNames []*string `json:"EventNames,omitempty" name:"EventNames"`
}

type IPDetailResultInfo struct {
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// UpdateTime

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// IpAddress

	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
	// 是否分配

	IsAllocated *bool `json:"IsAllocated,omitempty" name:"IsAllocated"`
	// 是否预留

	IsReserved *bool `json:"IsReserved,omitempty" name:"IsReserved"`
	// 服务器SN

	RelatedSN *string `json:"RelatedSN,omitempty" name:"RelatedSN"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 网段

	Segment *string `json:"Segment,omitempty" name:"Segment"`
}

type DeleteImageBuildingTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageBuildingTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageBuildingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHealthStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 项目总数

		ProjectTotalCount *int64 `json:"ProjectTotalCount,omitempty" name:"ProjectTotalCount"`
		// 项目异常总数

		ProjectAbnormalCount *int64 `json:"ProjectAbnormalCount,omitempty" name:"ProjectAbnormalCount"`
		// 应用总数

		AppTotalCount *int64 `json:"AppTotalCount,omitempty" name:"AppTotalCount"`
		// 应用异常总数

		AppAbnormalCount *int64 `json:"AppAbnormalCount,omitempty" name:"AppAbnormalCount"`
		// 组件总数

		ComTotalCount *int64 `json:"ComTotalCount,omitempty" name:"ComTotalCount"`
		// 组件异常总数

		ComAbnormalCount *int64 `json:"ComAbnormalCount,omitempty" name:"ComAbnormalCount"`
		// 产品总数

		ProductTotalCount *int64 `json:"ProductTotalCount,omitempty" name:"ProductTotalCount"`
		// 产品异常总数

		ProductAbnormalCount *int64 `json:"ProductAbnormalCount,omitempty" name:"ProductAbnormalCount"`
		// 集群总数

		ClusterTotalCount *int64 `json:"ClusterTotalCount,omitempty" name:"ClusterTotalCount"`
		// 集群异常总数

		ClusterAbnormalCount *int64 `json:"ClusterAbnormalCount,omitempty" name:"ClusterAbnormalCount"`
		// 模块总数

		TypeTotalCount *int64 `json:"TypeTotalCount,omitempty" name:"TypeTotalCount"`
		// 模块异常总数

		TypeAbnormalCount *int64 `json:"TypeAbnormalCount,omitempty" name:"TypeAbnormalCount"`
		// 节点总数

		NodeTotalCount *int64 `json:"NodeTotalCount,omitempty" name:"NodeTotalCount"`
		// 节点异常总数

		NodeAbnormalCount *int64 `json:"NodeAbnormalCount,omitempty" name:"NodeAbnormalCount"`
		// pod总数

		PodTotalCount *int64 `json:"PodTotalCount,omitempty" name:"PodTotalCount"`
		// pod异常总数

		PodAbnormalCount *int64 `json:"PodAbnormalCount,omitempty" name:"PodAbnormalCount"`
		// k8s组件总数

		K8sComTotalCount *int64 `json:"K8sComTotalCount,omitempty" name:"K8sComTotalCount"`
		// k8s组件异常总数

		K8sComAbnormalCount *int64 `json:"K8sComAbnormalCount,omitempty" name:"K8sComAbnormalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHealthStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHealthStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LanData struct {
	// 网卡

	Nic *string `json:"Nic,omitempty" name:"Nic"`
	// 子网掩码

	Netmask *string `json:"Netmask,omitempty" name:"Netmask"`
	// IP地址

	Ipaddress *string `json:"Ipaddress,omitempty" name:"Ipaddress"`
	// 网关

	Gateway *string `json:"Gateway,omitempty" name:"Gateway"`
}

type MultiNicSet struct {
	// 自定义网卡名称

	DefinitionName *string `json:"DefinitionName,omitempty" name:"DefinitionName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// bond配置；示例："BondConfig":[{"DefaulTgw":true,"NicName":"bond1","NicType":"lan"}]

	BondConfig []*NicConfig `json:"BondConfig,omitempty" name:"BondConfig"`
	// 网卡配置；示例：NicConfig":[{"DefaulTgw":true,"NicName":"eth0","NicType":"lan"}]

	NicConfig []*NicConfig `json:"NicConfig,omitempty" name:"NicConfig"`
}

type DescribeNetdeviceRelatedServersExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// RelatedServersSet

		RelatedServersSet []*Netdevice `json:"RelatedServersSet,omitempty" name:"RelatedServersSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetdeviceRelatedServersExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetdeviceRelatedServersExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterHistoryMetricDataRequest struct {
	*tchttp.BaseRequest

	// 指标名，多个以逗号分隔

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 起始时间

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束时间

	End *string `json:"End,omitempty" name:"End"`
	// 精度，如10s, 1m

	Step *string `json:"Step,omitempty" name:"Step"`
	// 查询标签，用于过滤特定的cluster

	Labels *string `json:"Labels,omitempty" name:"Labels"`
}

func (r *GetClusterHistoryMetricDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterHistoryMetricDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDataBaradMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监控指标

		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
		// 数据点起始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 数据点结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 数据统计周期

		Period *uint64 `json:"Period,omitempty" name:"Period"`
		// 监控数据列表

		DataPionts []*PointsObject `json:"DataPionts,omitempty" name:"DataPionts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDataBaradMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDataBaradMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDevOpsScenesByClusterRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点列表

	NodeList []*string `json:"NodeList,omitempty" name:"NodeList"`
	// 场景列表

	Scenes []*Scenes `json:"Scenes,omitempty" name:"Scenes"`
	// 风险控制

	Ndb *Ndb `json:"Ndb,omitempty" name:"Ndb"`
	// 通知详情

	Notification *DevopsNotification `json:"Notification,omitempty" name:"Notification"`
}

func (r *CreateDevOpsScenesByClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDevOpsScenesByClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRetentionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRetentionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRetentionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PerformServerTaskSet struct {
	// SvrCurrentStatus

	SvrCurrentStatus *string `json:"SvrCurrentStatus,omitempty" name:"SvrCurrentStatus"`
	// SvrIdcId

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// SvrSn

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// TaskId

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

type ListStorageClassesRequest struct {
	*tchttp.BaseRequest

	// kubernetes集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 按某个字段排序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否降序排序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 搜索条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListStorageClassesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListStorageClassesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleSvrNicIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务器的多网卡ip回收出参

		RecycleSvrNicSet []*RecycleSvrNicSet `json:"RecycleSvrNicSet,omitempty" name:"RecycleSvrNicSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecycleSvrNicIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleSvrNicIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AggOperator struct {
	// 日期直方聚合

	DateHistogram *DateHistogram `json:"DateHistogram,omitempty" name:"DateHistogram"`
}

type WanInfo struct {
	// Nic

	Nic *string `json:"Nic,omitempty" name:"Nic"`
	// Ipaddress

	Ipaddress *string `json:"Ipaddress,omitempty" name:"Ipaddress"`
}

type AddServersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTcsComponentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTcsComponentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTcsComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFieldsEnumExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 显示条目数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统键值

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 执行动作

	Op *string `json:"Op,omitempty" name:"Op"`
	// 过滤条件

	FieldFilters *string `json:"FieldFilters,omitempty" name:"FieldFilters"`
	// 用于条件过滤查询。例如过滤ID、名称、状态等

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
}

func (r *DescribeFieldsEnumExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFieldsEnumExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDevOpsStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回信息

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDevOpsStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDevOpsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IngressPathPlus struct {
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务端口

	ServicePort *string `json:"ServicePort,omitempty" name:"ServicePort"`
	// 失效

	Invalid *bool `json:"Invalid,omitempty" name:"Invalid"`
	// 路由名称

	IngressName *string `json:"IngressName,omitempty" name:"IngressName"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 灰度

	Gray *IngressGray `json:"Gray,omitempty" name:"Gray"`
}

type DeleteIngressClassResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIngressClassResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIngressClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceManagementOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 支撑服务管控组件状态

		ServiceManagementStatus []*ServiceInstanceStatus `json:"ServiceManagementStatus,omitempty" name:"ServiceManagementStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceManagementOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceManagementOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReserveOutBandIPRequest struct {
	*tchttp.BaseRequest

	// 请求串类型；示例：OutBandIP

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方id；默认：dcos

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方key；默认：dcos

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 请求动作；默认reserve

	Op *string `json:"Op,omitempty" name:"Op"`
	// 是否取消预留

	UnReserve *bool `json:"UnReserve,omitempty" name:"UnReserve"`
	// 预留、取消预留信息集合ReserveOutBandIPSet

	ReserveOutBandIPSet []*ReserveOutBandIPSet `json:"ReserveOutBandIPSet,omitempty" name:"ReserveOutBandIPSet"`
}

func (r *ReserveOutBandIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReserveOutBandIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddRAIDParam struct {
	// RAID名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type AddServer struct {
	// '设备固资编号'

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器SN(虚拟机此处为UUID,该值由云平提供和更新)

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// '设备型号',

	SvrDeviceType *string `json:"SvrDeviceType,omitempty" name:"SvrDeviceType"`
	// 'IDC名称',

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// '机架名称',

	SvrRackName *string `json:"SvrRackName,omitempty" name:"SvrRackName"`
	// '存放机位ID',

	SvrPosId *string `json:"SvrPosId,omitempty" name:"SvrPosId"`
	// '服务器厂商',

	SvrProducer *string `json:"SvrProducer,omitempty" name:"SvrProducer"`
	// 'IDC id',

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
}

type DeleteMonitorTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMonitorTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMonitorTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCronJobRequest struct {
	*tchttp.BaseRequest

	// 字符串ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 名字空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// CronJob的名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeCronJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCronJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NicConfig struct {
	// 默认网关；示例：true/false

	DefaultGw *bool `json:"DefaultGw,omitempty" name:"DefaultGw"`
	// 网卡名称；示例：eth0、eth1

	NicName *string `json:"NicName,omitempty" name:"NicName"`
	// 网卡类型；示例：lan/wan

	NicType *string `json:"NicType,omitempty" name:"NicType"`
}

type DeleteDashboardByUidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDashboardByUidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardByUidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashFolderByUidRequest struct {
	*tchttp.BaseRequest

	// 仪表盘目录唯一ID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
}

func (r *GetDashFolderByUidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashFolderByUidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashboardVersionContentByIdRequest struct {
	*tchttp.BaseRequest

	// 仪表盘ID

	DashId *string `json:"DashId,omitempty" name:"DashId"`
	// 版本号

	DashVersion *string `json:"DashVersion,omitempty" name:"DashVersion"`
}

func (r *GetDashboardVersionContentByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardVersionContentByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Laninfo struct {
	// Nic

	Nic *string `json:"Nic,omitempty" name:"Nic"`
	// Ipaddress

	Ipaddress *string `json:"Ipaddress,omitempty" name:"Ipaddress"`
	// Netmask

	Netmask *string `json:"Netmask,omitempty" name:"Netmask"`
	// Gateway

	Gateway *string `json:"Gateway,omitempty" name:"Gateway"`
}

type ServerHardDisk struct {
	// 返回详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 返回错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 硬盘信息

	HarddiskInfo *HarddiskInfo `json:"HarddiskInfo,omitempty" name:"HarddiskInfo"`
}

type UpgradeHistoryListRecord struct {
	// 升级批次号

	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
	// Upgrade/Rollback

	Type *string `json:"Type,omitempty" name:"Type"`
	// 老版本

	OldVersion *string `json:"OldVersion,omitempty" name:"OldVersion"`
	// 新版本

	NewVersion *string `json:"NewVersion,omitempty" name:"NewVersion"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type CreateLogConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志配置

		ConfigInfo *ConfigInfo `json:"ConfigInfo,omitempty" name:"ConfigInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLogConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVolumeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVolumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVolumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Statistics struct {
	// 状态，Running（执行中），NotTrigger（未触发），NotExecute（已触发未执行）

	Status *string `json:"Status,omitempty" name:"Status"`
	// 成功个数

	Succeeded *int64 `json:"Succeeded,omitempty" name:"Succeeded"`
	// 失败个数

	Failed *int64 `json:"Failed,omitempty" name:"Failed"`
}

type CreateCustomScriptSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCustomScriptSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomScriptSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateIngressYAMLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateIngressYAMLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateIngressYAMLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSilenceRequest struct {
	*tchttp.BaseRequest

	// 策略 id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 策略定义

	Command *UpdateSilence `json:"Command,omitempty" name:"Command"`
}

func (r *UpdateSilenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSilenceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopPodsByMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopPodsByMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopPodsByMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeCheckResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUpgradeCheckResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeCheckResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddMetaResourceTypeRequest struct {
	*tchttp.BaseRequest

	// 资源组对象

	Command *MetaResourceTypeCreateCommand `json:"Command,omitempty" name:"Command"`
	// 资源归属信息，运营端使用

	Param *ResourceOwnerQueryParam `json:"Param,omitempty" name:"Param"`
}

func (r *AddMetaResourceTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMetaResourceTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePhysvrsOverviewRequest struct {
	*tchttp.BaseRequest

	// 筛选AZ

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePhysvrsOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePhysvrsOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRAIDResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRAIDResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRAIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeAlertsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 告警列表

		AlertSet []*Alert `json:"AlertSet,omitempty" name:"AlertSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNodeAlertsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeAlertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeCondition struct {
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 状态，"True"、"False"、"Unknown"三个值

	Status *string `json:"Status,omitempty" name:"Status"`
	// 原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 相关信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 上次更新时间

	LastHeartbeatTime *string `json:"LastHeartbeatTime,omitempty" name:"LastHeartbeatTime"`
}

type ListClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群信息

		Clusters []*ClusterInfo `json:"Clusters,omitempty" name:"Clusters"`
		// 过滤参数

		FilterEntity *ClusterFilterEntity `json:"FilterEntity,omitempty" name:"FilterEntity"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServersAllocateLanIp struct {
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// SvrIpWanted

	SvrIpWanted *string `json:"SvrIpWanted,omitempty" name:"SvrIpWanted"`
	// Vlan

	Vlan *string `json:"Vlan,omitempty" name:"Vlan"`
}

type ModifyMultiNicModify struct {
	// 自定义网卡名称

	DefinitionName *string `json:"DefinitionName,omitempty" name:"DefinitionName"`
	// 网卡配置；示例：NicConfig":[{"DefaulTgw":true,"NicName":"eth0","NicType":"lan"}]

	NicConfig []*NicConfig `json:"NicConfig,omitempty" name:"NicConfig"`
	// bond配置；示例："BondConfig":[{"DefaulTgw":true,"NicName":"bond1","NicType":"lan"}]

	BondConfig []*NicConfig `json:"BondConfig,omitempty" name:"BondConfig"`
}

type CreateOrUpdateRuleGroupTempleRequest struct {
	*tchttp.BaseRequest

	// 策略模板定义

	Command *CreateOrUpdateRuleGroupTemple `json:"Command,omitempty" name:"Command"`
}

func (r *CreateOrUpdateRuleGroupTempleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateRuleGroupTempleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRelocateServerHistoryRequest struct {
	*tchttp.BaseRequest

	// Scheme

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"relocate_history_query"

	Op *string `json:"Op,omitempty" name:"Op"`
	// "ResultItem":"svr_asset_id"

	ResultItem *string `json:"ResultItem,omitempty" name:"ResultItem"`
	// "Filters":[{"Name":"SvrSn","Value":""},{"Name":"SvrAssetId","Value":"TYSV190521X1"}]}

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRelocateServerHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRelocateServerHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaseCluster struct {
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type BaseNode struct {
	// 节点名

	Name *string `json:"Name,omitempty" name:"Name"`
	// IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

type Service struct {
	// 服务名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 服务类型，ClusterIP，NodePort, LoadBalancer, ExternalName

	Type *string `json:"Type,omitempty" name:"Type"`
	// 选择器，过滤服务连接的Pod

	Selector *string `json:"Selector,omitempty" name:"Selector"`
	// 集群IP

	ClusterIP *string `json:"ClusterIP,omitempty" name:"ClusterIP"`
	// 服务端口

	PortMapping *ServicePort `json:"PortMapping,omitempty" name:"PortMapping"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 所属项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 所属应用

	App *string `json:"App,omitempty" name:"App"`
	// 所属组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// owner名字, 可以为空

	OwnerName *string `json:"OwnerName,omitempty" name:"OwnerName"`
	// owner类型, 例如Deployment. 可以为空

	OwnerKind *string `json:"OwnerKind,omitempty" name:"OwnerKind"`
	// 外部流量策略

	ExternalTrafficPolicy *string `json:"ExternalTrafficPolicy,omitempty" name:"ExternalTrafficPolicy"`
	// 负载均衡IP

	LoadBalancerIP *string `json:"LoadBalancerIP,omitempty" name:"LoadBalancerIP"`
	// 协议列表

	Protocols []*string `json:"Protocols,omitempty" name:"Protocols"`
	// 服务端口列表

	Ports []*string `json:"Ports,omitempty" name:"Ports"`
	// 容器端口列表

	TargetPorts []*string `json:"TargetPorts,omitempty" name:"TargetPorts"`
	// 节点端口列表

	NodePorts []*string `json:"NodePorts,omitempty" name:"NodePorts"`
	// 网络类型，内网:inner, 外网:outer

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// 所属集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 项目展示名称

	ProjectShowName *string `json:"ProjectShowName,omitempty" name:"ProjectShowName"`
	// Service JSON字符串

	JSON *string `json:"JSON,omitempty" name:"JSON"`
}

type DeleteClusterByAdminResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterByAdminResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterByAdminResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetPortResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTargetPortResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTargetPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeMetric struct {
	// 指标类型。count: 全部查询总数；field_count: 包含字段计数；cardinality:字段值去重计数；sum: 字段值加和；avg: 字段值平均值；max:字段值最大值；min：字段值最小值

	Type *string `json:"Type,omitempty" name:"Type"`
	// 字段名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 指标前端展示名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type ReinstallOsSet struct {
	// SvrSn

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// TaskId

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// SvrIdcName

	SvrIdcName *string `json:"SvrIdcName,omitempty" name:"SvrIdcName"`
	// SvrIdcId

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
}

type CreateServiceRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 服务

	Service *string `json:"Service,omitempty" name:"Service"`
	// 所属Owner类型，如Deployment, StatefulSet

	OwnerKind *string `json:"OwnerKind,omitempty" name:"OwnerKind"`
	// 所属Owner名称

	OwnerName *string `json:"OwnerName,omitempty" name:"OwnerName"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指向的组件类型

	ComponentKind *string `json:"ComponentKind,omitempty" name:"ComponentKind"`
	// 指向的组件类型

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
}

func (r *CreateServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelocateServerStartResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务器搬迁结果信息

		StartRelocateServerSet []*StartRelocateServerInfo `json:"StartRelocateServerSet,omitempty" name:"StartRelocateServerSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RelocateServerStartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelocateServerStartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IngressGray struct {
	// 是否开启

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 权重

	Weight *string `json:"Weight,omitempty" name:"Weight"`
	// 请求头名称

	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`
	// 请求头值

	HeaderValue *string `json:"HeaderValue,omitempty" name:"HeaderValue"`
	// 请求匹配正则

	HeaderPattern *string `json:"HeaderPattern,omitempty" name:"HeaderPattern"`
	// Cookie名称

	CookieName *string `json:"CookieName,omitempty" name:"CookieName"`
}

type Metric struct {
	// Unix时间戳

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 数值

	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type VirtualMachineFilterEntity struct {
	// 项目集合

	ProjectSet []*string `json:"ProjectSet,omitempty" name:"ProjectSet"`
	// 状态集合

	StatusSet []*string `json:"StatusSet,omitempty" name:"StatusSet"`
}

type CreateStatefulSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateStatefulSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateStatefulSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterVersionRequest struct {
	*tchttp.BaseRequest

	// 集群ID, 如global

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 如要查询容器版本信息, 需要指定节点, 不指定的话仅返回kubernetes版本信息

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
}

func (r *DescribeClusterVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HPA struct {
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 标签

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// workload类型, Deployment/StatefulSet

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// workload名字

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
	// 最大副本数量

	MaxReplicas *int64 `json:"MaxReplicas,omitempty" name:"MaxReplicas"`
	// 最小副本数量, 允许为null

	MinReplicas *int64 `json:"MinReplicas,omitempty" name:"MinReplicas"`
	// 指标

	Metrics []*HPAMetricSpec `json:"Metrics,omitempty" name:"Metrics"`
	// 缩容冷却时间, 单位为秒

	ScaleDownCooling *int64 `json:"ScaleDownCooling,omitempty" name:"ScaleDownCooling"`
	// 触发方式, 目前固定为Metric

	Method *string `json:"Method,omitempty" name:"Method"`
}

type LabelCondition struct {
	// 物理服务器固资编号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 物理服务器SN

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
}

type LogRule struct {
	// 策略中的规则名

	Name *string `json:"Name,omitempty" name:"Name"`
	// record or alert

	Type *string `json:"Type,omitempty" name:"Type"`
	// 持续时间

	For *string `json:"For,omitempty" name:"For"`
	// 规则表达式

	Expr *string `json:"Expr,omitempty" name:"Expr"`
	// 阈值，会有一个默认值，但允许用户改写

	ExprVal *string `json:"ExprVal,omitempty" name:"ExprVal"`
	// 级别

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 自定义label

	Labels []*LabelPair `json:"Labels,omitempty" name:"Labels"`
	// 自定义告警描述

	Annotation []*LabelPair `json:"Annotation,omitempty" name:"Annotation"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// > >= <= < = !=

	ExprAct *string `json:"ExprAct,omitempty" name:"ExprAct"`
	// 规则Id，前端忽略

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 规则描述，便于用户理解

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 策略id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 内部字段，前端忽略

	PromVal *string `json:"PromVal,omitempty" name:"PromVal"`
	// 规则拆解后的值

	ExprArgs []*LabelPair `json:"ExprArgs,omitempty" name:"ExprArgs"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 更新时间

	Updated *string `json:"Updated,omitempty" name:"Updated"`
}

type DescribeCronJobListRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间, 如为空字符串返回所有命名空间下的CronJob

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 根据label筛选符合条件的CronJob, 格式如: keyX in (foo,,baz),keyY,KeyZ notin ()

	LabelSelector *string `json:"LabelSelector,omitempty" name:"LabelSelector"`
}

func (r *DescribeCronJobListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCronJobListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelocateServerCancelRequest struct {
	*tchttp.BaseRequest

	// Scheme

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"relocate_cancel"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 取消搬迁信息

	RelocateCancelServers []*RelocateCancelServer `json:"RelocateCancelServers,omitempty" name:"RelocateCancelServers"`
}

func (r *RelocateServerCancelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelocateServerCancelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddServerLabelInfo struct {
	// '设备固资编号'

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// '服务器SN(虚拟机此处为UUID,该值由云平提供和更新)'

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// '新增服务器标签'

	SvrLabel *string `json:"SvrLabel,omitempty" name:"SvrLabel"`
}

type GetFiringRulesRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 数量上限

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 升序

	Asc *string `json:"Asc,omitempty" name:"Asc"`
	// 规则名称

	AlertName *string `json:"AlertName,omitempty" name:"AlertName"`
	// 降序

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 策略名称

	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`
	// 产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 模块

	Type *string `json:"Type,omitempty" name:"Type"`
	// 等级

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 类别

	Class *string `json:"Class,omitempty" name:"Class"`
}

func (r *GetFiringRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFiringRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Quota struct {
	// 总量

	Total *float64 `json:"Total,omitempty" name:"Total"`
	// 已分配量

	Allocated *float64 `json:"Allocated,omitempty" name:"Allocated"`
	// 最小需求量

	Request *float64 `json:"Request,omitempty" name:"Request"`
	// 最大限制量

	Limit *float64 `json:"Limit,omitempty" name:"Limit"`
	// 分配率

	Usage *float64 `json:"Usage,omitempty" name:"Usage"`
}

type DeleteOutbandStrategyRequest struct {
	*tchttp.BaseRequest

	// 请求串类型说明。
	// 固定值： "Scheme":"outband_passwd_strategy"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方systemid	此ID为请求方向cmdb注册的权限ID，由cmdb系统管理员提供和管理。
	// 默认值："SystemId":"sid_test"

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方systemkey	和上面systemid类似，由cmdb系统管理员提供和管理，cmdb侧会根据id和key以及调用方ip三者共同确定请求方的合法性。
	// 默认值："SystemKey":"sid_test"

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 删除自定义带外密码策略详细内容

	Data *OutbandStrategyData `json:"Data,omitempty" name:"Data"`
}

func (r *DeleteOutbandStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOutbandStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetworkPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetworkPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleSvrNicIPCondition struct {
	// 服务器SN

	Sn *string `json:"Sn,omitempty" name:"Sn"`
}

type StorageParameters struct {
	// 存储类型, 当前可能的值: CSP/CBS/local/ThirdParty

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 可选参数, 对于CSP/CBS/local等内置类型, 返回的是所有可选参数, 运营端管理员无法自行增加参数(可以删减). 对于ThirdParty类型的SC, 管理员可以自行增加选择型参数

	Parameters []*VolumeParameter `json:"Parameters,omitempty" name:"Parameters"`
	// 后端对应的真正StorageClass, 值由后端返回, 创建SC时需要带上.

	Backend *string `json:"Backend,omitempty" name:"Backend"`
	// 参数是否可编辑. ThirdParty类型可编辑, 其他内置类型不可编辑

	Editable *bool `json:"Editable,omitempty" name:"Editable"`
}

type AddOutbandStrategyRequest struct {
	*tchttp.BaseRequest

	// 请求方systemid	此ID为请求方向cmdb注册的权限ID，由cmdb系统管理员提供和管理。
	// 默认值："SystemId":"sid_test"

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方systemkey	和上面systemid类似，由cmdb系统管理员提供和管理，cmdb侧会根据id和key以及调用方ip三者共同确定请求方的合法性。
	// 默认值："SystemKey":"sid_test"

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 新增自定义带外密码策略详细内容

	Data *AddOutbandStrategyData `json:"Data,omitempty" name:"Data"`
	// 请求串类型说明。
	// 固定值： "Scheme":"outband_passwd_strategy"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
}

func (r *AddOutbandStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddOutbandStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMetaMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMetaMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMetaMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRouteRequest struct {
	*tchttp.BaseRequest

	// 路由ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReinstallOsExRequest struct {
	*tchttp.BaseRequest

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// idcid

	SvrIdcId *string `json:"SvrIdcId,omitempty" name:"SvrIdcId"`
	// 请求动作

	Command *string `json:"Command,omitempty" name:"Command"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 用户

	User *string `json:"User,omitempty" name:"User"`
	// 操作类型

	OperateType *int64 `json:"OperateType,omitempty" name:"OperateType"`
	// OS部署信息

	Data *ReinstallOsServers `json:"Data,omitempty" name:"Data"`
	// 选择装机工具：dcos或igniter

	DeployTool *string `json:"DeployTool,omitempty" name:"DeployTool"`
	// true:只装机智能网卡，false:安装智能网卡+黑石主机

	OnlySoc *bool `json:"OnlySoc,omitempty" name:"OnlySoc"`
	// 黑石主机和网卡关联的SN

	SvrRelatedSn *string `json:"SvrRelatedSn,omitempty" name:"SvrRelatedSn"`
	// 3：黑石主机，4：智能网卡

	SvrType *int64 `json:"SvrType,omitempty" name:"SvrType"`
}

func (r *ReinstallOsExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReinstallOsExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogRuleGroupTemplesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchLogRuleGroupTemplesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogRuleGroupTemplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageFieldsEnumSet struct {
	// 操作系统枚举值

	OsType []*string `json:"OsType,omitempty" name:"OsType"`
	// 系统架构枚举值

	SystemArch []*string `json:"SystemArch,omitempty" name:"SystemArch"`
	// 系统平台枚举值

	SystemPlatform []*string `json:"SystemPlatform,omitempty" name:"SystemPlatform"`
	// 镜像格式枚举值

	ImageType []*string `json:"ImageType,omitempty" name:"ImageType"`
	// 可用区枚举值

	ZoneName []*string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type ServerResultInfo struct {
	// 服务器资产ID

	SvrId *string `json:"SvrId,omitempty" name:"SvrId"`
}

type DescribeNodesOverviewRequest struct {
	*tchttp.BaseRequest

	// 查询ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *DescribeNodesOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodesOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerAllocateLanIPExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ServerAllocateLanIPInfo

		ServerAllocateLanIPInfo []*ServerAllocateLanIPInfo `json:"ServerAllocateLanIPInfo,omitempty" name:"ServerAllocateLanIPInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ServerAllocateLanIPExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerAllocateLanIPExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDashFolderByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仪表盘目录更新结果

		Data *DashFolderContent `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDashFolderByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDashFolderByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Container struct {
	// 容器名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 镜像名

	Image *string `json:"Image,omitempty" name:"Image"`
	// 镜像拉取策略, Always/Never/IfNotPresent

	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" name:"ImagePullPolicy"`
	// 是否是初始化容器

	Init *bool `json:"Init,omitempty" name:"Init"`
	// 是否是特权容器

	Privileged *bool `json:"Privileged,omitempty" name:"Privileged"`
	// 启动命令

	Command []*string `json:"Command,omitempty" name:"Command"`
	// 启动命令参数

	Args *string `json:"Args,omitempty" name:"Args"`
}

type CreateLogConfigRequest struct {
	*tchttp.BaseRequest

	// meta信息

	Metas []*Meta `json:"Metas,omitempty" name:"Metas"`
	// 日志用途标识

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志路径，支持统配

	Path *string `json:"Path,omitempty" name:"Path"`
	// 当path包含通配符时，该字段用于排除不必要的日志文件

	PathExcluded *string `json:"PathExcluded,omitempty" name:"PathExcluded"`
	// 是否启用多行匹配 默认不启用

	MultipleLine *bool `json:"MultipleLine,omitempty" name:"MultipleLine"`
	// 指定系统已有的parse

	ParserRef *string `json:"ParserRef,omitempty" name:"ParserRef"`
	// 用于手工指定配置，可用配置有json、regex等

	ParserFormat *string `json:"ParserFormat,omitempty" name:"ParserFormat"`
	// time字段名称

	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`
	// time格式

	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`
	// 提取正则

	Regex *string `json:"Regex,omitempty" name:"Regex"`
	// 行首正则

	LineHeaderRegex *string `json:"LineHeaderRegex,omitempty" name:"LineHeaderRegex"`
	// 当ParserFormat为LTSV时需要指定

	Type *string `json:"Type,omitempty" name:"Type"`
	// 用于手工指定配置，主要用于对特定定段再次进行解码

	Decoders []*string `json:"Decoders,omitempty" name:"Decoders"`
	// 是否提取字段

	FieldExtraction *bool `json:"FieldExtraction,omitempty" name:"FieldExtraction"`
	// 是否使用系统时间

	UseSystemTime *bool `json:"UseSystemTime,omitempty" name:"UseSystemTime"`
	// 日志样例

	Example *string `json:"Example,omitempty" name:"Example"`
	// 字段信息，json  样例:[{"Field":"cost","Type":"int"},{"Field":"log_level","Type":"string"},{"Field":"old_price","Type":"double"}]

	FieldMapping *string `json:"FieldMapping,omitempty" name:"FieldMapping"`
}

func (r *CreateLogConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogRuleGroupTemplesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogRuleGroupTemplesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogRuleGroupTemplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTargetWeightResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTargetWeightResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTargetWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IngressBackend struct {
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务端口

	ServicePort *string `json:"ServicePort,omitempty" name:"ServicePort"`
}

type StatefulSet struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 状态

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 正常副本数

	ReadyReplicas *int64 `json:"ReadyReplicas,omitempty" name:"ReadyReplicas"`
	// 全部副本数

	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 访问地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// 上次更新时间

	LastUpdateTimestamp *string `json:"LastUpdateTimestamp,omitempty" name:"LastUpdateTimestamp"`
	// 项目显示名称

	ProjectShowName *string `json:"ProjectShowName,omitempty" name:"ProjectShowName"`
}

type DeleteMetaResourceTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMetaResourceTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMetaResourceTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOSListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// {                "Id":270,                "Name":"Linux",                "CreateTime":"2020-05-08 16:30:39"            }

		OsDictionarySet []*OsParam `json:"OsDictionarySet,omitempty" name:"OsDictionarySet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOSListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOSListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageBuildingTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageBuildingTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageBuildingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPodHistoryMetricDataRequest struct {
	*tchttp.BaseRequest

	// 指标名，多个以逗号分隔

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 起始时间

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束时间

	End *string `json:"End,omitempty" name:"End"`
	// 查询标签，用于过滤特定的node

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 精度，如10s, 1m

	Step *string `json:"Step,omitempty" name:"Step"`
}

func (r *GetPodHistoryMetricDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPodHistoryMetricDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSeriesRequest struct {
	*tchttp.BaseRequest

	// match表达式

	Match *string `json:"Match,omitempty" name:"Match"`
	// 开始时间，s

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束时间，s

	End *string `json:"End,omitempty" name:"End"`
	// 指标

	Metrics *PromQLMetrics `json:"Metrics,omitempty" name:"Metrics"`
	// label过滤列表

	PromQL *string `json:"PromQL,omitempty" name:"PromQL"`
}

func (r *GetSeriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSeriesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleSvrNicIPRequest struct {
	*tchttp.BaseRequest

	// 服务器的多网卡ip回收入参

	RecycleSvrNicIPParams *RecycleSvrNicIPParams `json:"RecycleSvrNicIPParams,omitempty" name:"RecycleSvrNicIPParams"`
}

func (r *RecycleSvrNicIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleSvrNicIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceYAMLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateServiceYAMLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateServiceYAMLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenRegexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自动生成的正则表达式

		Regex *string `json:"Regex,omitempty" name:"Regex"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenRegexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenRegexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDefaultGroupByResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDefaultGroupByResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDefaultGroupByResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Secret struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 类型, TLS为kubernetes.io/tls

	Type *string `json:"Type,omitempty" name:"Type"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
}

type ServiceRequest struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 服务类型，ClusterIP, NodePort, LoadBalancer, ExternalName

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 选择器

	Selector []*KVPair `json:"Selector,omitempty" name:"Selector"`
	// 所属项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 所属应用

	App *string `json:"App,omitempty" name:"App"`
	// 端口映射列表

	Ports []*ServicePort `json:"Ports,omitempty" name:"Ports"`
	// 网络类型，Inner 内网，Outer 外网

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
}

type SvrSpecAllocWanIP struct {
	// 结果详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 资产Id

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type LabelModify struct {
	// 物理服务器标签

	SvrLabels *string `json:"SvrLabels,omitempty" name:"SvrLabels"`
}

type ListenerHealthCheck struct {
	// 是否开启健康检查：1（开启）、0（关闭）

	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`
	// 健康检查端口

	CheckPort *int64 `json:"CheckPort,omitempty" name:"CheckPort"`
	// 健康检查的响应超时时间

	TimeOut *int64 `json:"TimeOut,omitempty" name:"TimeOut"`
	// 健康检查探测间隔时间

	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`
	// 健康阈值，默认值：3，表示当连续探测三次健康则表示该转发正常

	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`
	// 不健康阈值，默认值：3，表示当连续探测三次不健康则表示该转发异常

	UnHealthNum *int64 `json:"UnHealthNum,omitempty" name:"UnHealthNum"`
}

type PV struct {
	// PV的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// PV的uid

	UID *string `json:"UID,omitempty" name:"UID"`
	// PV关联的StorageClass名称。存储类型

	StorageClass *string `json:"StorageClass,omitempty" name:"StorageClass"`
	// PV的回收策略. [Recycle, Delete, Retaim]

	ReclaimPolicy *string `json:"ReclaimPolicy,omitempty" name:"ReclaimPolicy"`
	// PV的存储空间大小，如100Gi

	Capacity *string `json:"Capacity,omitempty" name:"Capacity"`
	// PV所在的节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// PV挂载的参数

	MountOptions []*string `json:"MountOptions,omitempty" name:"MountOptions"`
	// 读写模式, 可能有多个模式. 例如[ROX, RWO]

	AccessModes []*string `json:"AccessModes,omitempty" name:"AccessModes"`
	// 关联的PVC名称

	PersistentVolumeClaim *string `json:"PersistentVolumeClaim,omitempty" name:"PersistentVolumeClaim"`
	// 关联的PVC命名空间

	PersistentVolumeClaimNamespace *string `json:"PersistentVolumeClaimNamespace,omitempty" name:"PersistentVolumeClaimNamespace"`
	// PVC所属项目, 如果不属于任何一个项目为空

	Project *string `json:"Project,omitempty" name:"Project"`
	// PV卷的模式，如 Block, Filesystem

	VolumeMode *string `json:"VolumeMode,omitempty" name:"VolumeMode"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 创建时间, 格式如2006-01-02T15:04:05

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// kubernetes对象 annotations

	Annotations []*Annotation `json:"Annotations,omitempty" name:"Annotations"`
}

type ModifyLabelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLabelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVirtualMachineAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVirtualMachineAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirtualMachineAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventRelatedAlertNames struct {
	// 事件唯一标识，事件名

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 关联该事件的告警名

	RelatedAlertNames []*string `json:"RelatedAlertNames,omitempty" name:"RelatedAlertNames"`
}

type SvrSpecRecyWanIP struct {
	// 结果详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 结果错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器IP

	SvrIp *string `json:"SvrIp,omitempty" name:"SvrIp"`
}

type TreeNodePod struct {
	// Pod名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Pod下的容器列表

	Containers []*string `json:"Containers,omitempty" name:"Containers"`
	// pod所在集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// Pod所在namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type CheckClusterDeletableRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CheckClusterDeletableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckClusterDeletableRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDaemonSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDaemonSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDaemonSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIngressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIngressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMultiNicDefinitionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改网卡信息出参

		MultiNicSet *ModifyMultiNicSet `json:"MultiNicSet,omitempty" name:"MultiNicSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMultiNicDefinitionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMultiNicDefinitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubMeta struct {
	// meta key

	MetaKey *string `json:"MetaKey,omitempty" name:"MetaKey"`
	// meta value

	MetaValue *string `json:"MetaValue,omitempty" name:"MetaValue"`
	// meta name

	MetaName *string `json:"MetaName,omitempty" name:"MetaName"`
	// 子meta

	SubMetas []*SubSubMeta `json:"SubMetas,omitempty" name:"SubMetas"`
	// config count

	ConfigCount *uint64 `json:"ConfigCount,omitempty" name:"ConfigCount"`
}

type VMAllocWanIP struct {
	// 结果详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 错误码

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 服务器资产Id

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器序列号

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
}

type AllocateServerSpecialWanIPListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"special_allocate_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 分配服务器特殊外网IP

	AllocateServerSpecialWanIPSet []*AllocateServerSpecialWanIP `json:"AllocateServerSpecialWanIPSet,omitempty" name:"AllocateServerSpecialWanIPSet"`
}

func (r *AllocateServerSpecialWanIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateServerSpecialWanIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDashFolderCommand struct {
	// Unique Id

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 仪表盘目录名称

	Title *string `json:"Title,omitempty" name:"Title"`
}

type DescribeClusterStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点状态

		NodeStatus *StatusCount `json:"NodeStatus,omitempty" name:"NodeStatus"`
		// 容器组状态

		PodStatus *StatusCount `json:"PodStatus,omitempty" name:"PodStatus"`
		// 项目状态

		ProjectStatus *StatusCount `json:"ProjectStatus,omitempty" name:"ProjectStatus"`
		// 应用状态

		AppStatus *StatusCount `json:"AppStatus,omitempty" name:"AppStatus"`
		// CPU配额

		CPUQuota *Quota `json:"CPUQuota,omitempty" name:"CPUQuota"`
		// 内存配额

		MemoryQuota *Quota `json:"MemoryQuota,omitempty" name:"MemoryQuota"`
		// 磁盘配额

		DiskQuota *Quota `json:"DiskQuota,omitempty" name:"DiskQuota"`
		// 容器组配额

		PodQuota *Quota `json:"PodQuota,omitempty" name:"PodQuota"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListNodesRequest struct {
	*tchttp.BaseRequest

	// 集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 偏移量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 返回的符合条件的节点数量，默认20

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 搜索字段名，支持Name，Role，IP，Description，已废弃

	SearchField *string `json:"SearchField,omitempty" name:"SearchField"`
	// 待搜索的值，已废弃

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 待搜索参数列表，支持Name, Role, IP, Description,Phase

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序字段，支持Name,PodCount,PodCIDR

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 集群ID列表

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

func (r *ListNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartWorkloadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartWorkloadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartWorkloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddMetaMetricRequest struct {
	*tchttp.BaseRequest

	// 资源归属信息，运营端使用

	Params *ResourceOwnerQueryParam `json:"Params,omitempty" name:"Params"`
	// 指标信息

	Command *MetaMetricCreateCommand `json:"Command,omitempty" name:"Command"`
}

func (r *AddMetaMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMetaMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeAttribute struct {
	// 节点名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 角色

	Role *string `json:"Role,omitempty" name:"Role"`
	// 节点IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 生命周期状态

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 容器组数

	PodCount *int64 `json:"PodCount,omitempty" name:"PodCount"`
	// Pod CIDR

	PodCIDR *string `json:"PodCIDR,omitempty" name:"PodCIDR"`
	// 所属集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 状态（Normal、Abnormal）

	Status *string `json:"Status,omitempty" name:"Status"`
	// 节点系统信息

	NodeInfo *NodeInfo `json:"NodeInfo,omitempty" name:"NodeInfo"`
	// 是否被封锁

	Unschedulable *bool `json:"Unschedulable,omitempty" name:"Unschedulable"`
	// 是否可封锁

	Sealable *bool `json:"Sealable,omitempty" name:"Sealable"`
	// 是否可解除封锁

	Unsealable *bool `json:"Unsealable,omitempty" name:"Unsealable"`
	// 是否可修改密码

	PasswordChangeable *bool `json:"PasswordChangeable,omitempty" name:"PasswordChangeable"`
	// 是否可驱逐

	Evictable *bool `json:"Evictable,omitempty" name:"Evictable"`
	// 是否可列出Pod

	PodsListable *bool `json:"PodsListable,omitempty" name:"PodsListable"`
	// 是否可以删除

	Deletable *bool `json:"Deletable,omitempty" name:"Deletable"`
	// 是否可重启

	Rebootable *bool `json:"Rebootable,omitempty" name:"Rebootable"`
}

type ModifyImageCondition struct {
	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

type DescribeNodesStatsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 可用的node数量, 可用的判断依据是node.status.condtions中类型为Ready的condition是否为True

		AvailableCount *int64 `json:"AvailableCount,omitempty" name:"AvailableCount"`
		// 不可用的node数量

		UnavailableCount *int64 `json:"UnavailableCount,omitempty" name:"UnavailableCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNodesStatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodesStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNodeLabelsRequest struct {
	*tchttp.BaseRequest

	// 待获取标签的节点名字

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 待获取标签的节点所在集群，不传默认使用global

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *GetNodeLabelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNodeLabelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSvrNicAllocationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定(更新)服务器多网卡策略出参

		UpdateSvrNicSet []*UpdateSvrNicSet `json:"UpdateSvrNicSet,omitempty" name:"UpdateSvrNicSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSvrNicAllocationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSvrNicAllocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalyseLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AnalyseLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AnalyseLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePersistentVolumeClaimResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePersistentVolumeClaimResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePersistentVolumeClaimResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalyzeSort struct {
	// 排序字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// 默认desc 倒序，asc正序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 字段值是否数值型

	IsNumber *bool `json:"IsNumber,omitempty" name:"IsNumber"`
}

type ListIngressesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合筛选条件的总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Ingress数组

		IngressSet []*Ingress `json:"IngressSet,omitempty" name:"IngressSet"`
		// 搜索参数列表

		FilterEntity *IngressFilterEntity `json:"FilterEntity,omitempty" name:"FilterEntity"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListIngressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListIngressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Annotation struct {
	// annotation的键

	Key *string `json:"Key,omitempty" name:"Key"`
	// annotation的值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ImportClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCreateVolumeParametersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 参数列表

		Parameters []*VolumeParameter `json:"Parameters,omitempty" name:"Parameters"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCreateVolumeParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCreateVolumeParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTcsMonitorRequest struct {
	*tchttp.BaseRequest

	// clusterid

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// start

	Start *string `json:"Start,omitempty" name:"Start"`
	// end

	End *string `json:"End,omitempty" name:"End"`
	// step

	Step *string `json:"Step,omitempty" name:"Step"`
}

func (r *DescribeTcsMonitorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTcsMonitorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyServersRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID，system_id 和 system_key为双方约定，主要是进行请求方鉴权

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key，system_id 和 system_key为双方约定，主要是进行请求方鉴权

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"modify"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 修改服务器参数

	ModifyServerSet []*ModifyServer `json:"ModifyServerSet,omitempty" name:"ModifyServerSet"`
}

func (r *ModifyServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyServersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContainerPort struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 主机端口

	HostPort *uint64 `json:"HostPort,omitempty" name:"HostPort"`
	// 主机IP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 容器端口

	ContainerPort *uint64 `json:"ContainerPort,omitempty" name:"ContainerPort"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DeleteVolumeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVolumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVolumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersOverviewRequest struct {
	*tchttp.BaseRequest

	// 集群名称，用于搜索

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeClustersOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClustersOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DevopsNotification struct {
	// 是否开启

	Switch *string `json:"Switch,omitempty" name:"Switch"`
	// 通知场景

	Type []*string `json:"Type,omitempty" name:"Type"`
	// 通知方式

	Method []*string `json:"Method,omitempty" name:"Method"`
	// 用户uid列表

	Users []*string `json:"Users,omitempty" name:"Users"`
}

type DeleteCustomScriptsRequest struct {
	*tchttp.BaseRequest

	// 自定义脚本的名称

	CustomScriptName []*string `json:"CustomScriptName,omitempty" name:"CustomScriptName"`
}

func (r *DeleteCustomScriptsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomScriptsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterNameRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群新名称

	ClusterDisplayName *string `json:"ClusterDisplayName,omitempty" name:"ClusterDisplayName"`
}

func (r *ModifyClusterNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Lan struct {
	// OS部署

	Data []*Laninfo `json:"Data,omitempty" name:"Data"`
}

type Project struct {
	// 项目健康状态 normal 或者abnormal

	Status *string `json:"Status,omitempty" name:"Status"`
	// 项目名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 项目id

	Id *string `json:"Id,omitempty" name:"Id"`
}

type SearchSilences struct {
	// 屏蔽规则id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 屏蔽规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 级别

	Level *string `json:"Level,omitempty" name:"Level"`
	// 对象类型、产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 创建者

	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`
	// 更新者

	UpdatedBy *string `json:"UpdatedBy,omitempty" name:"UpdatedBy"`
	// page offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// page limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 升序字段，field1,field2...

	Asc *string `json:"Asc,omitempty" name:"Asc"`
	// 降序字段，field1,field2...

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 屏蔽状态，active/pending/expired

	Effective *string `json:"Effective,omitempty" name:"Effective"`
	// 时间范围，开始时间

	StartsAt *string `json:"StartsAt,omitempty" name:"StartsAt"`
	// 时间范围，结束时间

	EndsAt *string `json:"EndsAt,omitempty" name:"EndsAt"`
	// 自定义变更人员，默认是当前用户，多个用户间逗号分隔

	Updater *string `json:"Updater,omitempty" name:"Updater"`
	// 是否启用正则匹配，可选true、false，默认false

	IsRegex *string `json:"IsRegex,omitempty" name:"IsRegex"`
}

type SystemComponent struct {
	// 组件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 正常数量

	NormalCount *int64 `json:"NormalCount,omitempty" name:"NormalCount"`
	// 异常数量

	AbnormalCount *int64 `json:"AbnormalCount,omitempty" name:"AbnormalCount"`
}

type ListBaseClustersRequest struct {
	*tchttp.BaseRequest

	// 列表行数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 查询条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListBaseClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListBaseClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReserveOutBandIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ReserveOutBandIP出参集合

		SvrReserveOutBandIPSet []*SvrReserveOutBandIPSet `json:"SvrReserveOutBandIPSet,omitempty" name:"SvrReserveOutBandIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReserveOutBandIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReserveOutBandIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleServerSpecialWanIPListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"special_recycle_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 回收服务器特殊外网IP

	RecycleServerSpecialWanIPSet []*RecycleServerSpecialWanIP `json:"RecycleServerSpecialWanIPSet,omitempty" name:"RecycleServerSpecialWanIPSet"`
}

func (r *RecycleServerSpecialWanIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecycleServerSpecialWanIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDevOpsScenesRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ListDevOpsScenesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDevOpsScenesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListStorageClassesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// StorageClass数组

		StorageClassSet []*StorageClass `json:"StorageClassSet,omitempty" name:"StorageClassSet"`
		// 可搜索的CSI列表

		CSISet []*string `json:"CSISet,omitempty" name:"CSISet"`
		// 可搜索的存储类型列表

		StorageTypeSet []*string `json:"StorageTypeSet,omitempty" name:"StorageTypeSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListStorageClassesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListStorageClassesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigServerDefResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回串中的data节点。"data":{"task_id":xxx}

		Data []*ConfigServerDefData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConfigServerDefResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfigServerDefResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListBasePodsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Pod列表

		PodSet []*BasePod `json:"PodSet,omitempty" name:"PodSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListBasePodsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListBasePodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IngressClass struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 所属项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 网络类型，内网:inner, 外网:outer

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// 资源规格, 大型: large, 中型: medium，小型:small

	ResourceSpec *string `json:"ResourceSpec,omitempty" name:"ResourceSpec"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 负载均衡VIP

	VIP *string `json:"VIP,omitempty" name:"VIP"`
	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 状态，active 运行中，pending 部署中，deleting 删除中，fail 部署失败，init 初始化中，updating 更新中

	AppStatus *string `json:"AppStatus,omitempty" name:"AppStatus"`
	// 是否默认

	Default *bool `json:"Default,omitempty" name:"Default"`
	// 项目展示名称

	ProjectShowName *string `json:"ProjectShowName,omitempty" name:"ProjectShowName"`
}

type LoadBalancerListener struct {
	// 监听器ID

	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
	// 监听器协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 监听器端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 监听器的健康检查信息

	HealthCheck *ListenerHealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`
	// 后端信息

	Targets []*BackendTarget `json:"Targets,omitempty" name:"Targets"`
	// 状态

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type RecycleSvrNicSet struct {
	// 返回结果详细说明

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 返回结果值，0成功， 1失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// 回收的I

	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type DescribeServiceManagementOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeServiceManagementOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceManagementOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatefulSetRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeStatefulSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStatefulSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListStatefulSetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 有状态应用列表

		StatefulSets []*StatefulSet `json:"StatefulSets,omitempty" name:"StatefulSets"`
		// 状态列表

		FilterEntity *StatefulSetFilterEntity `json:"FilterEntity,omitempty" name:"FilterEntity"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListStatefulSetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListStatefulSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogRuleGroupTemplesRequest struct {
	*tchttp.BaseRequest

	// 搜索条件

	Params *SearchRuleGroupTemples `json:"Params,omitempty" name:"Params"`
}

func (r *SearchLogRuleGroupTemplesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogRuleGroupTemplesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddMetaResourceTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 资源对象类型对象

		Data *MetaResourceType `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddMetaResourceTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMetaResourceTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMultiNicDefinitionRequest struct {
	*tchttp.BaseRequest

	// 删除网卡信息入参

	DeleteMultiNicParams *DeleteMultiNicParams `json:"DeleteMultiNicParams,omitempty" name:"DeleteMultiNicParams"`
}

func (r *DeleteMultiNicDefinitionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMultiNicDefinitionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomScriptSetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCustomScriptSetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomScriptSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardByIdRequest struct {
	*tchttp.BaseRequest

	// 仪表盘唯一ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteDashboardByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashFoldersRequest struct {
	*tchttp.BaseRequest

	// 仪表盘目录查询参数

	Params *DashFolderQueryParam `json:"Params,omitempty" name:"Params"`
}

func (r *GetDashFoldersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashFoldersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelocateServerCancelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RelocateServerCancelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RelocateServerCancelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyWebsocketTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// token

		Token *string `json:"Token,omitempty" name:"Token"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyWebsocketTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyWebsocketTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerSpecialAllocateLanIPExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ServerSpecialAllocateLanIPInfo

		ServerSpecialAllocateLanIPInfo []*ServerSpecialAllocateLanIPInfo `json:"ServerSpecialAllocateLanIPInfo,omitempty" name:"ServerSpecialAllocateLanIPInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ServerSpecialAllocateLanIPExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerSpecialAllocateLanIPExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMetaResourceTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMetaResourceTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMetaResourceTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewTrend struct {
	// 当前值

	Current *int64 `json:"Current,omitempty" name:"Current"`
	// 历史数据点

	Trend []*Metric `json:"Trend,omitempty" name:"Trend"`
	// 相比上周增长率

	GrowthRate *string `json:"GrowthRate,omitempty" name:"GrowthRate"`
}

type CreateNetworkPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetworkPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllVlanIdsExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllVlanIdsExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllVlanIdsExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListStorageTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可选存储类型及对应的参数

		StorageTypes []*StorageParameters `json:"StorageTypes,omitempty" name:"StorageTypes"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListStorageTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListStorageTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogRuleGroupsRequest struct {
	*tchttp.BaseRequest

	// 搜索条件

	Params *SearchRuleGroups `json:"Params,omitempty" name:"Params"`
}

func (r *SearchLogRuleGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogRuleGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群实例状态分布

		ClustersStatus *Status `json:"ClustersStatus,omitempty" name:"ClustersStatus"`
		// 结点状态分布

		NodesStatus *Status `json:"NodesStatus,omitempty" name:"NodesStatus"`
		// Pod状态分布

		PodsStatus *Status `json:"PodsStatus,omitempty" name:"PodsStatus"`
		// 支撑服务状态分布

		ServicesStatus *Status `json:"ServicesStatus,omitempty" name:"ServicesStatus"`
		// 账号数量历史趋势

		Accounts *OverviewTrend `json:"Accounts,omitempty" name:"Accounts"`
		// 项目数量历史趋势

		Projects *OverviewTrend `json:"Projects,omitempty" name:"Projects"`
		// 镜像数量历史趋势

		Images *OverviewTrend `json:"Images,omitempty" name:"Images"`
		// 应用数量历史趋势

		Applications []*OverviewTrend `json:"Applications,omitempty" name:"Applications"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBasicOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBasicOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitMachineRequest struct {
	*tchttp.BaseRequest

	// 待初始化的服务器

	Machine *Machine `json:"Machine,omitempty" name:"Machine"`
	// 脚本列表

	Scripts []*string `json:"Scripts,omitempty" name:"Scripts"`
}

func (r *InitMachineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitMachineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePersistentVolumeClaimRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 待删除的PVC列表, 只需传Namespace+Name

	PersistentVolumeClaims []*PVC `json:"PersistentVolumeClaims,omitempty" name:"PersistentVolumeClaims"`
}

func (r *DeletePersistentVolumeClaimRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePersistentVolumeClaimRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMetaResourceTypeRequest struct {
	*tchttp.BaseRequest

	// 资源组查询参数

	Params *MetaResourceTypeQueryParam `json:"Params,omitempty" name:"Params"`
}

func (r *ListMetaResourceTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMetaResourceTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRuleGroupsRequest struct {
	*tchttp.BaseRequest

	// 搜索条件

	Params *SearchRuleGroups `json:"Params,omitempty" name:"Params"`
}

func (r *SearchRuleGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRuleGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPodsRequest struct {
	*tchttp.BaseRequest

	// 若指定，则只列出指定节点上的Pod

	NodeNames []*string `json:"NodeNames,omitempty" name:"NodeNames"`
	// 待查询节点所在集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回的Pod数量，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段, 可选值: ["RestartTimes", "Name"], 默认按照Name排序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 降序排序, 默认false, 即增序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 描述多个查询条件, 模糊查询字段包括["Name", "Namespace", "IP", "WorkloadName"]

	SearchFilters []*SearchFilter `json:"SearchFilters,omitempty" name:"SearchFilters"`
	// 描述多个查询条件, 模糊查询字段包括["Name", "Namespace", "IP", "WorkloadName"]

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 项目Id

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ListPodsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPodsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VolumeParameter struct {
	// 参数唯一id

	Key *string `json:"Key,omitempty" name:"Key"`
	// 用于前端显示的中文名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数的详细描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// Select/Rang: 分别表示选择型参数和范围型参数

	ParameterType *string `json:"ParameterType,omitempty" name:"ParameterType"`
	// 租户端创建Volume时选择的确定性参数值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 运营端创建SC时, 管理员确定的的租户端可选参数列表

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 运营端创建SC时,  管理员确定的的租户端可选参数值的可选值

	AllowedValues []*string `json:"AllowedValues,omitempty" name:"AllowedValues"`
	// 范围型参数的最小值

	Min *int64 `json:"Min,omitempty" name:"Min"`
	// 范围型参数的最大值

	Max *int64 `json:"Max,omitempty" name:"Max"`
	// 租户端选择的范围型参数值

	ValueInt *int64 `json:"ValueInt,omitempty" name:"ValueInt"`
}

type Modify struct {
	// 服务器当前状态

	SvrCurrentStatus *string `json:"SvrCurrentStatus,omitempty" name:"SvrCurrentStatus"`
}

type DescribeMonitorTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监控模板

		MonitorTemplate *MonitorTemplateSpec `json:"MonitorTemplate,omitempty" name:"MonitorTemplate"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMonitorTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRollbackProgressRequest struct {
	*tchttp.BaseRequest

	// 回滚批次号

	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

func (r *GetRollbackProgressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRollbackProgressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoUpgradeOperationRequest struct {
	*tchttp.BaseRequest

	// 升级批次号

	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
	// 操作类型:Upgrade/Rollback

	Type *string `json:"Type,omitempty" name:"Type"`
	// 阶段

	Period *string `json:"Period,omitempty" name:"Period"`
	// 操作指令

	ActionCmd *string `json:"ActionCmd,omitempty" name:"ActionCmd"`
	// 升级前检查阶段重试用到

	PrecheckRetryItem *int64 `json:"PrecheckRetryItem,omitempty" name:"PrecheckRetryItem"`
}

func (r *DoUpgradeOperationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoUpgradeOperationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHPAResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHPAResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHPAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterAttributesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMiddlewareComponentRequest struct {
	*tchttp.BaseRequest

	// 集群id

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *DescribeMiddlewareComponentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMiddlewareComponentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRouteRequest struct {
	*tchttp.BaseRequest

	// 路由ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 路由信息struct

	Command *UpdateRoute `json:"Command,omitempty" name:"Command"`
}

func (r *UpdateRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkPolicyPeer struct {
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 组件类型

	ComponentKind *string `json:"ComponentKind,omitempty" name:"ComponentKind"`
	// 组件名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// IP段，"192.168.1.1/24","2001:db9::/64"

	CIDR *string `json:"CIDR,omitempty" name:"CIDR"`
	// 生效命名空间的标签

	NamespaceSelector *string `json:"NamespaceSelector,omitempty" name:"NamespaceSelector"`
	// 生效Pod的标签

	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`
	// 项目名称

	Project *string `json:"Project,omitempty" name:"Project"`
	// 项目显示名称

	ProjectShowName *string `json:"ProjectShowName,omitempty" name:"ProjectShowName"`
}

type ModifyCronJobRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// CronJob的JSON字符串

	CronJob *string `json:"CronJob,omitempty" name:"CronJob"`
	// 调度规则

	Schedule *string `json:"Schedule,omitempty" name:"Schedule"`
	// 并发策略, Allow/Forbid/Replace

	ConcurrencyPolicy *string `json:"ConcurrencyPolicy,omitempty" name:"ConcurrencyPolicy"`
	// 成功的历史任务数量限制

	SuccessfulJobsHistoryLimit *int64 `json:"SuccessfulJobsHistoryLimit,omitempty" name:"SuccessfulJobsHistoryLimit"`
	// 失败的历史任务数量限制

	FailedJobsHistoryLimit *int64 `json:"FailedJobsHistoryLimit,omitempty" name:"FailedJobsHistoryLimit"`
	// 描述, 不传表示不修改

	Description *string `json:"Description,omitempty" name:"Description"`
	// 标签, 不传表示不修改, 传空数组表示清空标签

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// 并行pod数量, 不传表示不修改

	Parallelism *int64 `json:"Parallelism,omitempty" name:"Parallelism"`
	// 修改为停止/启动, 不传表示不修改

	Suspend *bool `json:"Suspend,omitempty" name:"Suspend"`
}

func (r *ModifyCronJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCronJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateResourceCommand struct {
	// 资源对象名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 资源归属，取租户AppId

	Organization *string `json:"Organization,omitempty" name:"Organization"`
	// 类型标签

	LabelName *string `json:"LabelName,omitempty" name:"LabelName"`
	// 服务树节点选择标签，格式为{a="b",c="d"}

	Labels *string `json:"Labels,omitempty" name:"Labels"`
}

type LoadBalancerInstance struct {
	// 负载均衡实例 ID

	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
	// 负载均衡实例的名称

	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`
	// 负载均衡实例网络类型，内网Inner, 外网Outer

	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`
	// 模式，Tunnel或者FullNat

	Mode *string `json:"Mode,omitempty" name:"Mode"`
	// 用途

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// 负载均衡实例的 VIP 列表

	LoadBalancerVips *string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`
	// VIP

	VIP *string `json:"VIP,omitempty" name:"VIP"`
	// IP协议，IPv4，IPv6

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type AllocateVMVirtualIPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateVMVirtualIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateVMVirtualIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeHistoryDetailRequest struct {
	*tchttp.BaseRequest

	// 升级批次号

	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

func (r *GetUpgradeHistoryDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeHistoryDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlertStatData struct {
	// u5404u7ea7u522bu544au8b66u6761u6570

	Severities *AlertStats `json:"Severities,omitempty" name:"Severities"`
}

type RAIDDictionaryInfo struct {
	// 序号

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeClusterAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群ID

		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
		// 集群名称

		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
		// 描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// 集群类型，Global管控集群，TCS业务集群，TSF TSF集群, Imported纳管集群

		ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
		// 状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// k8s版本

		K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`
		// 是否专属

		IsExclusive *bool `json:"IsExclusive,omitempty" name:"IsExclusive"`
		// 资源模式，共享或专属

		ResourceMode *string `json:"ResourceMode,omitempty" name:"ResourceMode"`
		// 租户列表

		Tenants []*string `json:"Tenants,omitempty" name:"Tenants"`
		// 是否删除保护

		DeleteProtection *bool `json:"DeleteProtection,omitempty" name:"DeleteProtection"`
		// 创建时间

		CreateAt *string `json:"CreateAt,omitempty" name:"CreateAt"`
		// ApiServer列表

		ApiServers []*string `json:"ApiServers,omitempty" name:"ApiServers"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAlertTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// data

		Data []*AlertsTrend `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAlertTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterInfo struct {
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群类型，TCS(业务集群)、Global（管控集群）、TSF（TSF集群）、Imported（纳管集群）

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 集群cpu总数

	TotalCpu *string `json:"TotalCpu,omitempty" name:"TotalCpu"`
	// 集群已分配cpu总数

	AllocatedCpu *string `json:"AllocatedCpu,omitempty" name:"AllocatedCpu"`
	// 集群内存总数

	TotalMemory *string `json:"TotalMemory,omitempty" name:"TotalMemory"`
	// 集群已分配内存数

	AllocatedMemory *string `json:"AllocatedMemory,omitempty" name:"AllocatedMemory"`
	// 集群版本信息

	KubernetesVersion *string `json:"KubernetesVersion,omitempty" name:"KubernetesVersion"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 状态

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 节点数

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 异常节点数

	AbnormalNodeCount *int64 `json:"AbnormalNodeCount,omitempty" name:"AbnormalNodeCount"`
	// 容器组数

	PodCount *int64 `json:"PodCount,omitempty" name:"PodCount"`
	// 异常容器组数

	AbnormalPodCount *int64 `json:"AbnormalPodCount,omitempty" name:"AbnormalPodCount"`
	// 存储卷数

	VolumeCount *int64 `json:"VolumeCount,omitempty" name:"VolumeCount"`
	// 异常存储卷数

	AbnormalVolumeCount *int64 `json:"AbnormalVolumeCount,omitempty" name:"AbnormalVolumeCount"`
	// 告警数

	AlertCount *int64 `json:"AlertCount,omitempty" name:"AlertCount"`
	// master可升级版本

	AvailableMasterVersion *string `json:"AvailableMasterVersion,omitempty" name:"AvailableMasterVersion"`
	// worker可升级版本

	AvailableWorkerVersion *string `json:"AvailableWorkerVersion,omitempty" name:"AvailableWorkerVersion"`
	// worker当前版本

	WorkerVersion *string `json:"WorkerVersion,omitempty" name:"WorkerVersion"`
}

type ReserveOutBandIPSet struct {
	// 服务器SN

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 服务器带外ip

	SvrOutBandIp *string `json:"SvrOutBandIp,omitempty" name:"SvrOutBandIp"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 服务器固资号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
}

type SearchNotifications struct {
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// Fingerprint

	Fingerprint *string `json:"Fingerprint,omitempty" name:"Fingerprint"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Status

	Status *string `json:"Status,omitempty" name:"Status"`
	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Asc

	Asc *string `json:"Asc,omitempty" name:"Asc"`
	// Desc

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// IsRegex

	IsRegex *bool `json:"IsRegex,omitempty" name:"IsRegex"`
	// StartsAt

	StartsAt *string `json:"StartsAt,omitempty" name:"StartsAt"`
	// EndsAt

	EndsAt *string `json:"EndsAt,omitempty" name:"EndsAt"`
}

type SearchRetention struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type GetAlertRequest struct {
	*tchttp.BaseRequest

	// 告警id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetAlertRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryServerStatusRequest struct {
	*tchttp.BaseRequest

	// Zone

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// IdcName

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// SystemId

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// SystemKey

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// TaskId

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *QueryServerStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryServerStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreDashboardCommand struct {
	// 仪表盘版本

	Version *int64 `json:"Version,omitempty" name:"Version"`
}

type DeleteMetaMetricRequest struct {
	*tchttp.BaseRequest

	// 资源归属信息

	Params *ResourceOwnerQueryParam `json:"Params,omitempty" name:"Params"`
	// 资源对象类型名称

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 指标指纹ID

	Fingerprint *string `json:"Fingerprint,omitempty" name:"Fingerprint"`
}

func (r *DeleteMetaMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMetaMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Node信息

		Node *NodeAttribute `json:"Node,omitempty" name:"Node"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNodeAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DateHistogram struct {
	// 聚合日期字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// 时间间隔  s:秒，m:分，h：小时，d:天，w:周，M:月，y:年

	FixedInterval *string `json:"FixedInterval,omitempty" name:"FixedInterval"`
	// 时区，默认 Asia/Shanghai

	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
	// 最小匹配数，小于最小匹配数的日期聚合数据不展示

	MinDocCount *uint64 `json:"MinDocCount,omitempty" name:"MinDocCount"`
}

type GetDashboardVersionsByIdRequest struct {
	*tchttp.BaseRequest

	// 仪表盘ID

	DashId *string `json:"DashId,omitempty" name:"DashId"`
}

func (r *GetDashboardVersionsByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardVersionsByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListHPAsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// HPA列表

		HPAs []*HPA `json:"HPAs,omitempty" name:"HPAs"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListHPAsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListHPAsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DaemonSet struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 标签列表

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// 注解列表

	Annotations []*KVPair `json:"Annotations,omitempty" name:"Annotations"`
	// pod模板

	Template *PodTemplate `json:"Template,omitempty" name:"Template"`
	// 目标pod数量

	DesiredNumberScheduled *int64 `json:"DesiredNumberScheduled,omitempty" name:"DesiredNumberScheduled"`
	// 实际可用状态的pod数量

	NumberAvailable *int64 `json:"NumberAvailable,omitempty" name:"NumberAvailable"`
	// 项目名, 如果不存在对应的项目, 返回空字符串

	Project *string `json:"Project,omitempty" name:"Project"`
	// 状态

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 上次更新时间, 可能为空

	LastUpdateTimestamp *string `json:"LastUpdateTimestamp,omitempty" name:"LastUpdateTimestamp"`
	// 项目显示名称

	ProjectShowName *string `json:"ProjectShowName,omitempty" name:"ProjectShowName"`
}

type DescribeDefaultNetworkPolicyRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *DescribeDefaultNetworkPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefaultNetworkPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTaskExRequest struct {
	*tchttp.BaseRequest

	// 服务器SN

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *QueryTaskExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTaskExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CleanerScriptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 脚本内容

		Data *float64 `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CleanerScriptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CleanerScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageBuildingTaskRequest struct {
	*tchttp.BaseRequest

	// 待查询的任务, 需指定Name

	ImageBuildingTask *ImageBuildingTask `json:"ImageBuildingTask,omitempty" name:"ImageBuildingTask"`
}

func (r *DescribeImageBuildingTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageBuildingTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateServerWanIP struct {
	// 服务器资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 服务器需要分配IP

	SvrIpWanted *string `json:"SvrIpWanted,omitempty" name:"SvrIpWanted"`
}

type DescribeProjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 项目健康状况列表

		Projects []*Project `json:"Projects,omitempty" name:"Projects"`
		// 应用健康状况列表

		Application []*Application `json:"Application,omitempty" name:"Application"`
		// 应用总数

		AppTotalCount *int64 `json:"AppTotalCount,omitempty" name:"AppTotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkloadsRequest struct {
	*tchttp.BaseRequest

	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 搜索条件，Name,  Kind

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *ListWorkloadsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListWorkloadsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSvrNicAllocation struct {
	// 服务器Sn

	Sn *string `json:"Sn,omitempty" name:"Sn"`
}

type Node struct {
	// 节点名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 角色

	Role *string `json:"Role,omitempty" name:"Role"`
	// 节点IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 生命周期

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 已被分配的CPU核数

	AllocatedCpu *float64 `json:"AllocatedCpu,omitempty" name:"AllocatedCpu"`
	// 已被分配的内存，单位字节

	AllocatedMemory *int64 `json:"AllocatedMemory,omitempty" name:"AllocatedMemory"`
	// 总CPU核数

	TotalCpu *float64 `json:"TotalCpu,omitempty" name:"TotalCpu"`
	// 总内存，单位字节

	TotalMemory *int64 `json:"TotalMemory,omitempty" name:"TotalMemory"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// Pod CIDR

	PodCIDR *string `json:"PodCIDR,omitempty" name:"PodCIDR"`
	// 是否被封锁

	Unschedulable *bool `json:"Unschedulable,omitempty" name:"Unschedulable"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 是否可封锁

	Sealable *bool `json:"Sealable,omitempty" name:"Sealable"`
	// 是否可解除封锁

	Unsealable *bool `json:"Unsealable,omitempty" name:"Unsealable"`
	// 是否可修改密码

	PasswordChangeable *bool `json:"PasswordChangeable,omitempty" name:"PasswordChangeable"`
	// 是否可驱逐

	Evictable *bool `json:"Evictable,omitempty" name:"Evictable"`
	// 是否可列出Pod

	PodsListable *bool `json:"PodsListable,omitempty" name:"PodsListable"`
	// 是否可以删除

	Deletable *bool `json:"Deletable,omitempty" name:"Deletable"`
	// 所属k8s集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 是否可重启

	Rebootable *bool `json:"Rebootable,omitempty" name:"Rebootable"`
	// 容器组数

	PodCount *int64 `json:"PodCount,omitempty" name:"PodCount"`
	// 异常Pod数

	AbnormalPodCount *int64 `json:"AbnormalPodCount,omitempty" name:"AbnormalPodCount"`
	// 总磁盘，单位字节

	TotalDisk *int64 `json:"TotalDisk,omitempty" name:"TotalDisk"`
	// 已分配磁盘

	AllocatedDisk *float64 `json:"AllocatedDisk,omitempty" name:"AllocatedDisk"`
	// 系统信息

	NodeInfo *NodeInfo `json:"NodeInfo,omitempty" name:"NodeInfo"`
	// CPU使用率

	CPUUsage *float64 `json:"CPUUsage,omitempty" name:"CPUUsage"`
	// 内存使用率

	MemoryUsage *float64 `json:"MemoryUsage,omitempty" name:"MemoryUsage"`
	// 磁盘使用率

	DiskUsage *float64 `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// 正常异常状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 开启自愈true，关闭自愈false

	Healable *bool `json:"Healable,omitempty" name:"Healable"`
}

type DeleteLogRuleGroupTemplesRequest struct {
	*tchttp.BaseRequest

	// 批量删除ids，"Ids":"1,2,3"

	Params *BatchDeleteIds `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteLogRuleGroupTemplesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogRuleGroupTemplesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNetworkPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveUpgradeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveUpgradeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveUpgradeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDefaultNetworkPolicyRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 默认模式：项目间互通:reachable, 项目间隔离：isolated

	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

func (r *SetDefaultNetworkPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDefaultNetworkPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddOutbandStrategies struct {
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 是否包含数字（默认值1）

	HaveNumber *uint64 `json:"HaveNumber,omitempty" name:"HaveNumber"`
	// 是否包含小写字母（默认值1）

	HaveLowerLetter *uint64 `json:"HaveLowerLetter,omitempty" name:"HaveLowerLetter"`
	// 是否包含大写字母（默认值1）

	HaveUpperLetter *uint64 `json:"HaveUpperLetter,omitempty" name:"HaveUpperLetter"`
	// 特殊字符

	SpecialCharacters *string `json:"SpecialCharacters,omitempty" name:"SpecialCharacters"`
	// 密码长度，默认15

	PasswdLength *uint64 `json:"PasswdLength,omitempty" name:"PasswdLength"`
	// 是否启用该策略 （默认禁用，即0）

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
}

type ConfirmAlertRequest struct {
	*tchttp.BaseRequest

	// alert id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *ConfirmAlertRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfirmAlertRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群

		Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
		// 网络策略

		NetworkPolicy *NetworkPolicyStruct `json:"NetworkPolicy,omitempty" name:"NetworkPolicy"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenTimeFormatRequest struct {
	*tchttp.BaseRequest

	// 日志时间值

	LogTime *string `json:"LogTime,omitempty" name:"LogTime"`
}

func (r *GenTimeFormatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenTimeFormatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMetricMatrixByDimensionRequest struct {
	*tchttp.BaseRequest

	// 维度，如cluster,node,pod

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
	// {name,value}

	Labels []*LabelPair `json:"Labels,omitempty" name:"Labels"`
	// 需要显示的列指标

	MetricList []*string `json:"MetricList,omitempty" name:"MetricList"`
	// 开始时间

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束时间

	End *string `json:"End,omitempty" name:"End"`
	// 时间间隔，10s,1m

	Step *string `json:"Step,omitempty" name:"Step"`
}

func (r *GetMetricMatrixByDimensionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMetricMatrixByDimensionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListStatefulSetsRequest struct {
	*tchttp.BaseRequest

	// 集群ID列表

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
	// 每页条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 搜索参数，名称、命名空间、项目

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListStatefulSetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListStatefulSetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllocateServerVirtualIPRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 虚拟内网："Op":"allocate_vitrual_lan_ip"     虚拟外网： "Op":"allocate_vitrual_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 分配物理服务器虚拟内网外网信息集合

	AllocateServerVirtualIPSet []*VirtualLanOrWanIP `json:"AllocateServerVirtualIPSet,omitempty" name:"AllocateServerVirtualIPSet"`
}

func (r *AllocateServerVirtualIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateServerVirtualIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImageListRequest struct {
	*tchttp.BaseRequest

	// "Op": "query_image"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 过滤条件信息

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制返回个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeImageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePasswordRequest struct {
	*tchttp.BaseRequest

	// 待更新的节点名字

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 待更新的节点所在集群，不传默认使用global

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 新用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 新密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *UpdatePasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageParams struct {
	// 要修改目标

	Condition *ModifyImageCondition `json:"Condition,omitempty" name:"Condition"`
	// 修改的内容

	Modify *ModifyImageModify `json:"Modify,omitempty" name:"Modify"`
}

type AllocateSvrNicIPParams struct {
	// 服务器SN

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// 分配ip的信息；示例："Allocation":[{"NicName":"bond22222","NicType":"lan","Ip":"127.0.0.1"}]

	Allocation []*AllocationParams `json:"Allocation,omitempty" name:"Allocation"`
}

type AnalyzeFilter struct {
	// > < >= <= = !=

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 过滤指标key

	FilterKey *string `json:"FilterKey,omitempty" name:"FilterKey"`
	// 过滤指标值

	FilterValue *uint64 `json:"FilterValue,omitempty" name:"FilterValue"`
}

type DeleteDeploymentRequest struct {
	*tchttp.BaseRequest

	// 集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间以及名字的列表

	NamespacedNames []*NamespacedName `json:"NamespacedNames,omitempty" name:"NamespacedNames"`
}

func (r *DeleteDeploymentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDeploymentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollbackWorkloadRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// workload类型, Deployment/StatefulSet/DaemonSet

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 要回滚到的版本号

	Revision *int64 `json:"Revision,omitempty" name:"Revision"`
}

func (r *RollbackWorkloadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RollbackWorkloadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Hit struct {
	// 标签信息

	Labels []*Label `json:"Labels,omitempty" name:"Labels"`
	// 日志信息

	LogInfos []*LogInfo `json:"LogInfos,omitempty" name:"LogInfos"`
}

type BasePod struct {
	// 集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 节点名

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 容器组名

	Name *string `json:"Name,omitempty" name:"Name"`
	// IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type DescribeRelocateInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRelocateInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRelocateInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMatrixRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	Start *string `json:"Start,omitempty" name:"Start"`
	// 步长

	Step *string `json:"Step,omitempty" name:"Step"`
	// 结束时间

	End *string `json:"End,omitempty" name:"End"`
	// 查询query数组

	Targets []*Target `json:"Targets,omitempty" name:"Targets"`
}

func (r *GetMatrixRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMatrixRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDashFolderByUidRequest struct {
	*tchttp.BaseRequest

	// 仪表盘目录唯一ID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 仪表盘目录更新内容

	Command *UpdateDashFolderCommand `json:"Command,omitempty" name:"Command"`
}

func (r *UpdateDashFolderByUidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDashFolderByUidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServersRecycleLanIp struct {
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// SvrLanIp

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
}

type DeleteResourceObjectRequest struct {
	*tchttp.BaseRequest

	// 资源归属

	Params *ResourceOwnerQueryParam `json:"Params,omitempty" name:"Params"`
	// 服务树节点标签

	Labels *string `json:"Labels,omitempty" name:"Labels"`
}

func (r *DeleteResourceObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteResourceObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDcosFunctionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDcosFunctionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDcosFunctionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerJoinToBanksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		LogId *string `json:"LogId,omitempty" name:"LogId"`
		// 返回信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ServerJoinToBanksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerJoinToBanksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMetricInstantByDimensionRequest struct {
	*tchttp.BaseRequest

	// 维度，如cluster,node,pod

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
	// {name,value}

	Labels []*LabelPair `json:"Labels,omitempty" name:"Labels"`
	// 需要显示的列指标

	MetricList []*string `json:"MetricList,omitempty" name:"MetricList"`
	// 开始时间

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束时间

	End *string `json:"End,omitempty" name:"End"`
	// 时间间隔，10s,1m

	Step *string `json:"Step,omitempty" name:"Step"`
}

func (r *GetMetricInstantByDimensionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMetricInstantByDimensionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRoutesByLabelSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchRoutesByLabelSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRoutesByLabelSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogConfigRequest struct {
	*tchttp.BaseRequest

	// 配置id

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeleteLogConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleGroupTemplesRequest struct {
	*tchttp.BaseRequest

	// 批量删除ids，"Ids":"1,2,3"

	Params *BatchDeleteIds `json:"Params,omitempty" name:"Params"`
}

func (r *DeleteRuleGroupTemplesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRuleGroupTemplesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkloadPodsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Pod列表

		PodSet []*Pod `json:"PodSet,omitempty" name:"PodSet"`
		// 筛选列表

		FilterEntity *PodFilterEntity `json:"FilterEntity,omitempty" name:"FilterEntity"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListWorkloadPodsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListWorkloadPodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashboardPermissionsByIdRequest struct {
	*tchttp.BaseRequest

	// 仪表盘ID

	DashId *string `json:"DashId,omitempty" name:"DashId"`
}

func (r *GetDashboardPermissionsByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardPermissionsByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImageRequest struct {
	*tchttp.BaseRequest

	// "Op": "modify_image"

	Op *string `json:"Op,omitempty" name:"Op"`
	// "ModifyImageParams":[{"Condition":{"ImageId": "ins-123456"},"Modify":{"ImageName": "image_modify_test","ImageDescribe": "镜像修改","SystemArch":"arm64","OsType":"windows"}}]

	ModifyImageParams *ModifyImageParams `json:"ModifyImageParams,omitempty" name:"ModifyImageParams"`
}

func (r *ModifyImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchDeleteIds struct {
	// 逗号分隔，"Ids": "3,4,5"

	Ids *string `json:"Ids,omitempty" name:"Ids"`
	// 逗号分隔，”names“:"z,z"

	Names *string `json:"Names,omitempty" name:"Names"`
}

type AddRAIDRequest struct {
	*tchttp.BaseRequest

	// "Op": "add_raid"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 参考"AddRAIDParams": {"Name":"Linux"}

	AddRAIDParams *AddRAIDParam `json:"AddRAIDParams,omitempty" name:"AddRAIDParams"`
}

func (r *AddRAIDRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddRAIDRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 服务名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DeleteServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIgniterImagesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeIgniterImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIgniterImagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReinstallOsServers struct {
	// 服务器信息

	Servers []*ReinstallOsInfo `json:"Servers,omitempty" name:"Servers"`
}

type IgniterNodesSet struct {
	// 详细信息

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 查询结果成功/失败(0/1)

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// NetBoot

	NetBoot *string `json:"NetBoot,omitempty" name:"NetBoot"`
	// IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// Arch

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	// SerialPort

	SerialPort *string `json:"SerialPort,omitempty" name:"SerialPort"`
	// Mgt

	Mgt *string `json:"Mgt,omitempty" name:"Mgt"`
	// BMC

	BMC *string `json:"BMC,omitempty" name:"BMC"`
	// Groups

	Groups *string `json:"Groups,omitempty" name:"Groups"`
	// InstallNic

	InstallNic *string `json:"InstallNic,omitempty" name:"InstallNic"`
	// BmcPassword

	BmcPassword *string `json:"BmcPassword,omitempty" name:"BmcPassword"`
	// BmcUsername

	BmcUsername *string `json:"BmcUsername,omitempty" name:"BmcUsername"`
	// Serial

	Serial *string `json:"Serial,omitempty" name:"Serial"`
	// PostBootScripts

	PostBootScripts *string `json:"PostBootScripts,omitempty" name:"PostBootScripts"`
	// Cons

	Cons *string `json:"Cons,omitempty" name:"Cons"`
	// Mtm

	Mtm *string `json:"Mtm,omitempty" name:"Mtm"`
	// AddKCmdline

	AddKCmdline *string `json:"AddKCmdline,omitempty" name:"AddKCmdline"`
	// Status

	Status *string `json:"Status,omitempty" name:"Status"`
	// Mac

	Mac *string `json:"Mac,omitempty" name:"Mac"`
	// CpuCount

	CpuCount *string `json:"CpuCount,omitempty" name:"CpuCount"`
	// CpuType

	CpuType *string `json:"CpuType,omitempty" name:"CpuType"`
	// Memory

	Memory *string `json:"Memory,omitempty" name:"Memory"`
	// DiskSize

	DiskSize *string `json:"DiskSize,omitempty" name:"DiskSize"`
	// OS

	OS *string `json:"OS,omitempty" name:"OS"`
	// TftpServer

	TftpServer *string `json:"TftpServer,omitempty" name:"TftpServer"`
	// TftpDir

	TftpDir *string `json:"TftpDir,omitempty" name:"TftpDir"`
	// PeerDevice

	PeerDevice *string `json:"PeerDevice,omitempty" name:"PeerDevice"`
	// NodeType

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// SerialFlow

	SerialFlow *string `json:"SerialFlow,omitempty" name:"SerialFlow"`
	// StatusTime

	StatusTime *string `json:"StatusTime,omitempty" name:"StatusTime"`
	// SupportedArchs

	SupportedArchs *string `json:"SupportedArchs,omitempty" name:"SupportedArchs"`
	// IgniterMaster

	IgniterMaster *string `json:"IgniterMaster,omitempty" name:"IgniterMaster"`
	// SN

	SN *string `json:"SN,omitempty" name:"SN"`
	// InstallResult

	InstallResult *string `json:"InstallResult,omitempty" name:"InstallResult"`
	// PostScripts

	PostScripts *string `json:"PostScripts,omitempty" name:"PostScripts"`
}

type NativeStorageClass struct {
	// sc名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 标签列表

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// 注解列表

	Annotations []*KVPair `json:"Annotations,omitempty" name:"Annotations"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 存储类型, Local/CSP/ThirdParty

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 后端CSI插件名字

	CSI *string `json:"CSI,omitempty" name:"CSI"`
	// 回收策略, Delete/Retain/Recycle

	ReclaimPolicy *string `json:"ReclaimPolicy,omitempty" name:"ReclaimPolicy"`
	// mount参数

	MountOptions []*string `json:"MountOptions,omitempty" name:"MountOptions"`
	// 文件系统类型, 如ext4, 目前只对于csp/local适用

	FSType *string `json:"FSType,omitempty" name:"FSType"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 该SC支持的挂载模式列表, Filesystem/Block

	SupportedVolumeModes []*string `json:"SupportedVolumeModes,omitempty" name:"SupportedVolumeModes"`
	// 该SC支持的挂载读写模式, ReadWriteOnce/ReadWriteMany/ReadOnlyMany

	SupportedAccessModes []*string `json:"SupportedAccessModes,omitempty" name:"SupportedAccessModes"`
}

type QueryMonitorResult struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 指标性能数据列表

	Data []*MetricData `json:"Data,omitempty" name:"Data"`
}

type DescribeApiServerMetricsRequest struct {
	*tchttp.BaseRequest

	// cluster

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// metric

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// start

	Start *string `json:"Start,omitempty" name:"Start"`
	// end

	End *string `json:"End,omitempty" name:"End"`
	// step

	Step *string `json:"Step,omitempty" name:"Step"`
}

func (r *DescribeApiServerMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApiServerMetricsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRuleGroupTemplesRequest struct {
	*tchttp.BaseRequest

	// 搜索条件

	Params *SearchRuleGroupTemples `json:"Params,omitempty" name:"Params"`
}

func (r *SearchRuleGroupTemplesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRuleGroupTemplesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddMetaMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 指标信息

		Data *MetaMetric `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddMetaMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMetaMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 服务名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListClusterVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 版本列表

		ClusterVersions []*ClusterVersion `json:"ClusterVersions,omitempty" name:"ClusterVersions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListClusterVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListClusterVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDevOpsHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自愈历史

		Items []*DevopsJob `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListDevOpsHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDevOpsHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardQueryParam struct {
	// 查询字符串

	Query *string `json:"Query,omitempty" name:"Query"`
	// 页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Page

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 是否加星

	Starred *string `json:"Starred,omitempty" name:"Starred"`
	// 标签

	Tag []*string `json:"Tag,omitempty" name:"Tag"`
	// 类型，可以是dashboard或dash-folder，默认两者都返回

	Type *string `json:"Type,omitempty" name:"Type"`
}

type ModifyMultiNicCondition struct {
	// 自定义网卡名称

	DefinitionName *string `json:"DefinitionName,omitempty" name:"DefinitionName"`
}

type ConfirmAlertResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConfirmAlertResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfirmAlertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancerRequest struct {
	*tchttp.BaseRequest

	// 待删除的负载均衡实例ID列表

	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

func (r *DeleteLoadBalancerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLoadBalancerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterStatusRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAlertStatsRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAlertStatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertStatsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPodHistoryMetricDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPodHistoryMetricDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPodHistoryMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardFolderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功id数组

		Succeeded []*string `json:"Succeeded,omitempty" name:"Succeeded"`
		// 失败id数组

		Failed []*string `json:"Failed,omitempty" name:"Failed"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDashboardFolderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDashboardFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDashFolderByUidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 仪表盘目录更新结果

		Data *DashFolderContent `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDashFolderByUidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDashFolderByUidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateStorageClassResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateStorageClassResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateStorageClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePodsOverviewWithAlertsRequest struct {
	*tchttp.BaseRequest

	// 是否展示pod详情，默认false

	ShowDetail *bool `json:"ShowDetail,omitempty" name:"ShowDetail"`
}

func (r *DescribePodsOverviewWithAlertsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePodsOverviewWithAlertsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Label struct {
	// 标签名

	LabelName *string `json:"LabelName,omitempty" name:"LabelName"`
	// 标签值

	LabelValue *string `json:"LabelValue,omitempty" name:"LabelValue"`
}

type Pod struct {
	// Pod名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 生命周期

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 正常异常状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 所在主机名

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 所在主机IP

	NodeIp *string `json:"NodeIp,omitempty" name:"NodeIp"`
	// Pod IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 请求的CPU量，单位为核

	CpuRequest *float64 `json:"CpuRequest,omitempty" name:"CpuRequest"`
	// 请求的内存量，单位为字节数

	MemoryRequest *int64 `json:"MemoryRequest,omitempty" name:"MemoryRequest"`
	// 工作负载类型，如Deployment、Statefulset

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 工作负载名字

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
	// 重启次数

	RestartTimes *int64 `json:"RestartTimes,omitempty" name:"RestartTimes"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 所属应用

	Application *string `json:"Application,omitempty" name:"Application"`
	// pod指标map

	Metrics *PodMonitor `json:"Metrics,omitempty" name:"Metrics"`
	// User

	User *string `json:"User,omitempty" name:"User"`
	// 运行时间

	Age *float64 `json:"Age,omitempty" name:"Age"`
}

type ModifyMultiNicSet struct {
	// 返回结果详细说明

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// 返回结果详细说明

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 返回结果值，0成功， 1失败

	Result *int64 `json:"Result,omitempty" name:"Result"`
}

type GetCronJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CronJob的JSON字符串

		CronJobJSON *string `json:"CronJobJSON,omitempty" name:"CronJobJSON"`
		// 项目ID

		Project *string `json:"Project,omitempty" name:"Project"`
		// 项目名字

		ProjectShowName *string `json:"ProjectShowName,omitempty" name:"ProjectShowName"`
		// 状态

		Phase *string `json:"Phase,omitempty" name:"Phase"`
		// 上次更新时间

		LastUpdateTimestamp *string `json:"LastUpdateTimestamp,omitempty" name:"LastUpdateTimestamp"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCronJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCronJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNotificationRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 通知规则列表

		Data []*NotificationSpec `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNotificationRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNotificationRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PhysvrsListSet struct {
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Result

	Result *int64 `json:"Result,omitempty" name:"Result"`
	// PhysvrsResultInfo

	PhysvrsResultInfo *PhysvrsResultInfo `json:"PhysvrsResultInfo,omitempty" name:"PhysvrsResultInfo"`
}

type AllocateVMWanIPListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟机分配外网IP信息

		VMAllocWanIPSet []*VMAllocWanIP `json:"VMAllocWanIPSet,omitempty" name:"VMAllocWanIPSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateVMWanIPListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateVMWanIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateRuleGroupRequest struct {
	*tchttp.BaseRequest

	// 策略定义

	Command *CreateOrUpdateRuleGroup `json:"Command,omitempty" name:"Command"`
}

func (r *CreateOrUpdateRuleGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateRuleGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListClusterData struct {
	// 集群总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 集群列表信息

	Clusters []*ClusterInfo `json:"Clusters,omitempty" name:"Clusters"`
}

type QueryServerStatus struct {
	// state

	State *string `json:"State,omitempty" name:"State"`
	// Status

	Status *string `json:"Status,omitempty" name:"Status"`
	// Detail

	Detail *string `json:"Detail,omitempty" name:"Detail"`
	// Error

	Error *string `json:"Error,omitempty" name:"Error"`
	// TaskType

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
}

type AllocateVMLanIPListRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server_vm"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// "Op":"allocate_lan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 分配VM内网信息集合

	AllocateVMLanIPSet []*AllocateVMLanIP `json:"AllocateVMLanIPSet,omitempty" name:"AllocateVMLanIPSet"`
}

func (r *AllocateVMLanIPListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateVMLanIPListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDashFolderByUidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDashFolderByUidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashFolderByUidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashBoardsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDashBoardsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashBoardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRuleGroupRequest struct {
	*tchttp.BaseRequest

	// 策略 id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetRuleGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRuleGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNodeAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNodeAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNodeAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaResourceTypeUpdateCommand struct {
	// 资源对象类型名称，在租户和资源组下唯一

	Name *string `json:"Name,omitempty" name:"Name"`
	// 资源对象类型别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 生效资源对象，支持通配，如{tcs_product="*"}

	Selector *string `json:"Selector,omitempty" name:"Selector"`
}

type AllocateSvrNicIPRequest struct {
	*tchttp.BaseRequest

	// 服务器的多网卡ip分配入参

	AllocateSvrNicIPParams *AllocateSvrNicIPParams `json:"AllocateSvrNicIPParams,omitempty" name:"AllocateSvrNicIPParams"`
}

func (r *AllocateSvrNicIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateSvrNicIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TreeNode struct {
	// 结点名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否含有子结点

	HasChildren *bool `json:"HasChildren,omitempty" name:"HasChildren"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 状态, Running/NotReady

	Status *string `json:"Status,omitempty" name:"Status"`
	// 用于显示的结点名

	ShowName *string `json:"ShowName,omitempty" name:"ShowName"`
}

type UpdateLogConfigStateRequest struct {
	*tchttp.BaseRequest

	// 待修改配置ids

	ConfigIds []*string `json:"ConfigIds,omitempty" name:"ConfigIds"`
	// 状态，0 禁用 1启用

	State *uint64 `json:"State,omitempty" name:"State"`
}

func (r *UpdateLogConfigStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogConfigStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateRuleGroup struct {
	// 策略名称，同一个租户不可重复

	Name *string `json:"Name,omitempty" name:"Name"`
	// 评估间隔，建议1m

	Interval *string `json:"Interval,omitempty" name:"Interval"`
	// 类型、产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 告警对象Labels

	Objects []*Matcher `json:"Objects,omitempty" name:"Objects"`
	// 规则数组

	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 是否启用，true启用，false暂不启用

	Enable *string `json:"Enable,omitempty" name:"Enable"`
	// 乐观锁版本, 创建传0

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 策略Id，根据id是否>0区分create还是update

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 更新时间

	Updated *string `json:"Updated,omitempty" name:"Updated"`
	// 策略类型，1为自定义事件生成策略，默认为0 ，告警策略

	GroupType *int64 `json:"GroupType,omitempty" name:"GroupType"`
	// 自定义事件名称，GroupType为1时必填

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// alert or record

	Type *string `json:"Type,omitempty" name:"Type"`
	// labels

	Labels []*LabelPair `json:"Labels,omitempty" name:"Labels"`
}

type WorkerNodeInfo struct {
	// 节点地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// 节点类型：worker

	Role *string `json:"Role,omitempty" name:"Role"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// pod数目

	PodNum *int64 `json:"PodNum,omitempty" name:"PodNum"`
}

type CreateSilenceRequest struct {
	*tchttp.BaseRequest

	// 屏蔽struct

	Command *CreateSilence `json:"Command,omitempty" name:"Command"`
}

func (r *CreateSilenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSilenceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeploymentsStatsRequest struct {
	*tchttp.BaseRequest

	// 集群ID, 如global

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 符合kubernetes格式的过滤器, 如app=niginx

	LabelSelector *string `json:"LabelSelector,omitempty" name:"LabelSelector"`
	// 如指定, 只统计改命名空间下的Deployment, 否则统计整个集群的Deployment

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeDeploymentsStatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeploymentsStatsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVirtualMachineAttributesRequest struct {
	*tchttp.BaseRequest

	// 集群

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 虚拟ID

	VmId *string `json:"VmId,omitempty" name:"VmId"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyVirtualMachineAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVirtualMachineAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRuleGroupTemples struct {
	// 策略模板id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 策略模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 归属产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// page offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// page limit, default 1000

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 升序字段，file1,field2...

	Asc *string `json:"Asc,omitempty" name:"Asc"`
	// 降序字段，file1,field2...

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 对象类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 是否启用正则匹配，可选true、false，默认false

	IsRegex *string `json:"IsRegex,omitempty" name:"IsRegex"`
}

type ServerStatus struct {
	// SN

	Sn *string `json:"Sn,omitempty" name:"Sn"`
	// mac地址

	Mac *string `json:"Mac,omitempty" name:"Mac"`
	// 机器状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type OutbandStrategiesCondition struct {
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 策略ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

type CreateTLSSecretResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTLSSecretResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTLSSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLabelRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"label"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 操作类型："Op":"delete"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 修改的标签内容

	DeleteLabelSet []*LabelConditionInfo `json:"DeleteLabelSet,omitempty" name:"DeleteLabelSet"`
}

func (r *DeleteLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashFoldersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仪表盘目录检索结果

		Data []*DashFolderQueryResult `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDashFoldersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashFoldersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeHistoryListRequest struct {
	*tchttp.BaseRequest

	// 升级历史记录列表过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 记录数起始偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回最大记录数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetUpgradeHistoryListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeHistoryListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListNodeRemediationTemplatesRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ListNodeRemediationTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListNodeRemediationTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecycleSvrNicIPParams struct {
	// 选择ip回收的条件；"Condition":{"Sn":"TEST123"}示例：

	Condition *RecycleSvrNicIPCondition `json:"Condition,omitempty" name:"Condition"`
	// 回收的ip数组；示例:["127.0.0.1"]

	Ips []*string `json:"Ips,omitempty" name:"Ips"`
}

type CreateStorageClassRequest struct {
	*tchttp.BaseRequest

	// kubernetes集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// SC名称, 全局唯一

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 存储类型, 如local/CSP/CBS/ThirdParty

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 后端对应的k8s StorageClass

	Backend *string `json:"Backend,omitempty" name:"Backend"`
	// 租户端创建Volume时的参数

	Parameters []*VolumeParameter `json:"Parameters,omitempty" name:"Parameters"`
}

func (r *CreateStorageClassRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateStorageClassRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMetaResourceTypeRequest struct {
	*tchttp.BaseRequest

	// 资源对象归属信息，运营端使用

	Params *ResourceOwnerQueryParam `json:"Params,omitempty" name:"Params"`
	// 资源对象类型名称

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *DeleteMetaResourceTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMetaResourceTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebootNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RebootNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateIngressYAMLRequest struct {
	*tchttp.BaseRequest

	// 集群ID, 如"global"

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 新的YAML

	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
	// 待更新的Ingress名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 待更新的Ingress namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *UpdateIngressYAMLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateIngressYAMLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudProduct struct {
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 告警是否已经配置

	AlertConfiged *bool `json:"AlertConfiged,omitempty" name:"AlertConfiged"`
	// 首次部署时间

	FirstDeployTimestamp *string `json:"FirstDeployTimestamp,omitempty" name:"FirstDeployTimestamp"`
	// 最后变更时间, 可能为空

	LastUpdateTimestamp *string `json:"LastUpdateTimestamp,omitempty" name:"LastUpdateTimestamp"`
}

type DescribeMiddlewareComponentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 正常个数

		NormalCount *int64 `json:"NormalCount,omitempty" name:"NormalCount"`
		// 异常个数

		AbnormalCount *int64 `json:"AbnormalCount,omitempty" name:"AbnormalCount"`
		// 中间件列表

		MdComponent *MdComponentStatistic `json:"MdComponent,omitempty" name:"MdComponent"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMiddlewareComponentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMiddlewareComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportLogsRequest struct {
	*tchttp.BaseRequest

	// Meta信息

	Metas []*Meta `json:"Metas,omitempty" name:"Metas"`
	// 日志名，可以选多个

	Name []*string `json:"Name,omitempty" name:"Name"`
	// 时间范围起始，纳秒时间戳，默认当前时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 时间范围结束，纳秒时间戳，默认一小时前

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 返回记录条数，默认值：500，最大值：1000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// LogQL

	Query *string `json:"Query,omitempty" name:"Query"`
	// 排序

	Sort []*SortInfo `json:"Sort,omitempty" name:"Sort"`
	// 查询方式，0: LogQL

	QueryType *uint64 `json:"QueryType,omitempty" name:"QueryType"`
	// 数据类型 默认 0：日志 1 事件

	DataType *uint64 `json:"DataType,omitempty" name:"DataType"`
}

func (r *ExportLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMachineVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 版本列表

		MachineVersions []*MachineVersion `json:"MachineVersions,omitempty" name:"MachineVersions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMachineVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMachineVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashFolderPermissionsByUidRequest struct {
	*tchttp.BaseRequest

	// 仪表盘目录唯一ID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
}

func (r *GetDashFolderPermissionsByUidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashFolderPermissionsByUidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetUnschedulableResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetUnschedulableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetUnschedulableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteListenerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteListenerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDaemonSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// kubernetes类型的DaemonSet描述JSON字符串

		DaemonSetJSON *string `json:"DaemonSetJSON,omitempty" name:"DaemonSetJSON"`
		// 状态

		Phase *string `json:"Phase,omitempty" name:"Phase"`
		// 上次更新时间

		LastUpdateTimestamp *string `json:"LastUpdateTimestamp,omitempty" name:"LastUpdateTimestamp"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDaemonSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDaemonSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDaemonSetRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 项目ID

	Project *string `json:"Project,omitempty" name:"Project"`
}

func (r *DescribeDaemonSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDaemonSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDashboardCommand struct {
	// 仪表盘内容，JSON字符串

	Dashboard *string `json:"Dashboard,omitempty" name:"Dashboard"`
	// 仪表盘变更消息说明

	Message *string `json:"Message,omitempty" name:"Message"`
	// 目录ID

	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
}

type VirtualMachineNamespaceByName struct {
	// vmid

	Name *string `json:"Name,omitempty" name:"Name"`
	// vm namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type ProductTopologyNode struct {
	// 名称, 如某个region

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件列表

	Components []*TreeNode `json:"Components,omitempty" name:"Components"`
}

type WithdrawServers struct {
	// 服务器信息

	Condition *Condition `json:"Condition,omitempty" name:"Condition"`
	// 服务器当前状态

	Modify *Modify `json:"Modify,omitempty" name:"Modify"`
}

type AllocateVMVirtualIPRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"server_vm"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 虚拟内网："Op":"allocate_vitrual_lan_ip"     虚拟外网： "Op":"allocate_vitrual_wan_ip"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 分配VM虚拟内网外网信息集合

	AllocateVMVirtualIPSet []*VirtualLanOrWanIP `json:"AllocateVMVirtualIPSet,omitempty" name:"AllocateVMVirtualIPSet"`
}

func (r *AllocateVMVirtualIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllocateVMVirtualIPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInitMachineScriptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 脚本列表

		Scripts []*ModuleScript `json:"Scripts,omitempty" name:"Scripts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInitMachineScriptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInitMachineScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPersistentVolumeClaimsRequest struct {
	*tchttp.BaseRequest

	// kubernetes集群名, 如global

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 按某个字段排序, [Name, ID, Namespace, PersistentVolume, CreationTimestamp]

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否降序排序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 搜索条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 若指定, 只返回指定namespace下的PVC

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *ListPersistentVolumeClaimsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPersistentVolumeClaimsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSecretsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 每页条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 搜索条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListSecretsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSecretsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchDashFoldersRequest struct {
	*tchttp.BaseRequest

	// 仪表盘目录查询参数

	Params *DashFolderQueryParam `json:"Params,omitempty" name:"Params"`
}

func (r *SearchDashFoldersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchDashFoldersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PodFilterEntity struct {
	// 状态列表

	PhaseSet []*string `json:"PhaseSet,omitempty" name:"PhaseSet"`
	// 正常、异常状态列表

	StatusSet []*string `json:"StatusSet,omitempty" name:"StatusSet"`
	// 应该列表

	AppSet []*string `json:"AppSet,omitempty" name:"AppSet"`
	// 用户列表

	UinSet *string `json:"UinSet,omitempty" name:"UinSet"`
	// Node IP列表

	NodeIpSet []*string `json:"NodeIpSet,omitempty" name:"NodeIpSet"`
	// 项目列表

	ProjectSet []*string `json:"ProjectSet,omitempty" name:"ProjectSet"`
}

type PointsObject struct {
	// 监控实例的维度组合

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 监控数据点数组，每个点的时间跨度为一个Period值

	Points []*float64 `json:"Points,omitempty" name:"Points"`
}

type ListIngressesRequest struct {
	*tchttp.BaseRequest

	// 集群ID, 如"global"

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段, 支持["Name"]

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否降序排序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 筛选条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 项目Id，非必填

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 集群ID列表

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

func (r *ListIngressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListIngressesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceObjectRequest struct {
	*tchttp.BaseRequest

	// 查询资源服务对象参数

	Params *SearchResourceCommand `json:"Params,omitempty" name:"Params"`
}

func (r *ListResourceObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServersSpecialRecycleLanIp struct {
	// SvrAssetId

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// SvrLanIp

	SvrLanIp *string `json:"SvrLanIp,omitempty" name:"SvrLanIp"`
}

type OutbandStrategyData struct {
	// 删除\修改自定义带外策略Strategies

	Strategies []*OutbandStrategies `json:"Strategies,omitempty" name:"Strategies"`
}

type ModifyImageModify struct {
	// 修改后的镜像名

	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
	// 修改后的系统架构

	SystemArch *string `json:"SystemArch,omitempty" name:"SystemArch"`
	// 修改后的操作系统类型

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 修改后的镜像描述

	ImageDescribe *string `json:"ImageDescribe,omitempty" name:"ImageDescribe"`
	// 修改后的系统平台

	SystemPlatform *string `json:"SystemPlatform,omitempty" name:"SystemPlatform"`
	// 修改后的系统版本

	SystemVersions *string `json:"SystemVersions,omitempty" name:"SystemVersions"`
	// 镜像使用服务器设备类型(0:通用服务器;1:黑石服务器)

	AvailableModel *uint64 `json:"AvailableModel,omitempty" name:"AvailableModel"`
}

type DcosFunctionServiceItem struct {
	// dcos功能名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 服务ID

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
}

type Deployment struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 副本数

	Replicas *uint64 `json:"Replicas,omitempty" name:"Replicas"`
	// 正常副本数

	ReadyReplicas *uint64 `json:"ReadyReplicas,omitempty" name:"ReadyReplicas"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 访问地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 状态

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 上次更新时间, 可能为空

	LastUpdateTimestamp *string `json:"LastUpdateTimestamp,omitempty" name:"LastUpdateTimestamp"`
	// 项目显示名称

	ProjectShowName *string `json:"ProjectShowName,omitempty" name:"ProjectShowName"`
}

type CreateIngressClassRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 网络类型，内网:inner，外网:outer

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// 资源规格, 小型 small，中型 medium 大型 large

	ResourceSpec *string `json:"ResourceSpec,omitempty" name:"ResourceSpec"`
}

func (r *CreateIngressClassRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIngressClassRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoadBalancerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLoadBalancerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDeploymentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDeploymentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDeploymentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPodsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 满足条件的Pod列表

		PodSet []*Pod `json:"PodSet,omitempty" name:"PodSet"`
		// 满足条件的Pod数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 满足条件的Pod正常异常状态列表

		StatusSet []*string `json:"StatusSet,omitempty" name:"StatusSet"`
		// 满足条件的Pod app列表

		AppSet []*string `json:"AppSet,omitempty" name:"AppSet"`
		// 满足条件的Pod 用户列表

		UinSet []*string `json:"UinSet,omitempty" name:"UinSet"`
		// 满足条件的Pod node列表

		NodeIpSet []*string `json:"NodeIpSet,omitempty" name:"NodeIpSet"`
		// 满足条件的Pod 项目列表

		ProjectSet []*string `json:"ProjectSet,omitempty" name:"ProjectSet"`
		// 满足条件的Pod状态列表

		PhaseSet []*string `json:"PhaseSet,omitempty" name:"PhaseSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPodsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomScriptSetRequest struct {
	*tchttp.BaseRequest

	// 自定义脚本集名

	CustomScriptSetName *string `json:"CustomScriptSetName,omitempty" name:"CustomScriptSetName"`
	// 重新命名的自定义脚本集名

	CustomScriptSetNameNew *string `json:"CustomScriptSetNameNew,omitempty" name:"CustomScriptSetNameNew"`
	// 加入脚本集中的自定义脚本名列表

	CustomScriptName []*string `json:"CustomScriptName,omitempty" name:"CustomScriptName"`
	// 脚本集类型(0:初始化脚本,1:后置脚本)

	CustomScriptSetType *string `json:"CustomScriptSetType,omitempty" name:"CustomScriptSetType"`
}

func (r *ModifyCustomScriptSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomScriptSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCronJobRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// CronJob的JSON字符串

	CronJob *string `json:"CronJob,omitempty" name:"CronJob"`
}

func (r *CreateCronJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCronJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMonitorTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板参数

	TemplateParameter *TemplateParameter `json:"TemplateParameter,omitempty" name:"TemplateParameter"`
}

func (r *CreateMonitorTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMonitorTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClustersOverviewRequest struct {
	*tchttp.BaseRequest

	// 集群id，多个集群用逗号分隔

	Ids *string `json:"Ids,omitempty" name:"Ids"`
}

func (r *GetClustersOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClustersOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合筛选条件的总应用个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 符合条件的应用信息

		ApplicationSet []*AppInstance `json:"ApplicationSet,omitempty" name:"ApplicationSet"`
		// 用于筛选的项目名列表

		ProjectSet []*string `json:"ProjectSet,omitempty" name:"ProjectSet"`
		// 用于筛选的用户名列表

		UserNameSet []*string `json:"UserNameSet,omitempty" name:"UserNameSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAppInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PodStatus struct {
	// 正常pod个数

	NormalCount *int64 `json:"NormalCount,omitempty" name:"NormalCount"`
	// 异常pod个数

	AbnormalCount *int64 `json:"AbnormalCount,omitempty" name:"AbnormalCount"`
}

type CreateDaemonSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDaemonSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDaemonSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProductOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenameServerModify struct {
	// 设备名称

	SvrDeviceName *string `json:"SvrDeviceName,omitempty" name:"SvrDeviceName"`
}

type StorageClass struct {
	// 名称, 唯一ID

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 存储类型, 如local/CSP/CBS/ThirdParty

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 后端CSI

	CSI *string `json:"CSI,omitempty" name:"CSI"`
	// 读写模式

	ReadWriteMode *string `json:"ReadWriteMode,omitempty" name:"ReadWriteMode"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 后端对应的k8s StorageClass

	Backend *string `json:"Backend,omitempty" name:"Backend"`
	// 租户端创建Volume时的参数

	Parameters []*VolumeParameter `json:"Parameters,omitempty" name:"Parameters"`
	// 可创建磁盘的容量下限

	LowerBound *string `json:"LowerBound,omitempty" name:"LowerBound"`
	// 可创建磁盘的容量上限

	UpperBound *string `json:"UpperBound,omitempty" name:"UpperBound"`
	// 回收策略

	ReclaimPolicy *string `json:"ReclaimPolicy,omitempty" name:"ReclaimPolicy"`
	// 支持的文件系统

	FSType []*string `json:"FSType,omitempty" name:"FSType"`
}

type ModifyServerPara struct {
	// '服务器所属业务'

	SvrBussiness *string `json:"SvrBussiness,omitempty" name:"SvrBussiness"`
	// 服务器SN

	SvrSn *string `json:"SvrSn,omitempty" name:"SvrSn"`
	// 服务器当前状态(0未开电 1已开电 2运营中 3已关机 4已下架 5待部署 6搬迁中 20运行中(VM) 21已销毁(VM))

	SvrCurrentStatus *string `json:"SvrCurrentStatus,omitempty" name:"SvrCurrentStatus"`
	// 服务器厂商

	SvrProducer *string `json:"SvrProducer,omitempty" name:"SvrProducer"`
	// 设备固资编号

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// 机架名称

	SvrRackName *string `json:"SvrRackName,omitempty" name:"SvrRackName"`
	// 设备型号

	SvrDeviceType *string `json:"SvrDeviceType,omitempty" name:"SvrDeviceType"`
	// 服务器硬件描述

	SvrDeviceDescript *string `json:"SvrDeviceDescript,omitempty" name:"SvrDeviceDescript"`
	// 存放机位ID

	SvrPosId *string `json:"SvrPosId,omitempty" name:"SvrPosId"`
	// 服务器GPU信息

	SvrGpu *string `json:"SvrGpu,omitempty" name:"SvrGpu"`
}

type ListDaemonSetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// daemonset列表

		DaemonSets []*DaemonSet `json:"DaemonSets,omitempty" name:"DaemonSets"`
		// 状态列表

		PhaseSet []*string `json:"PhaseSet,omitempty" name:"PhaseSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListDaemonSetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDaemonSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRAIDResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRAIDResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRAIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DrivePod struct {
	// pod名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// pod状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 查询指标

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 下钻时间

	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
	// 应用

	App *string `json:"App,omitempty" name:"App"`
	// 值

	Latest *float64 `json:"Latest,omitempty" name:"Latest"`
}

type CreateMultiNicDefinitionRequest struct {
	*tchttp.BaseRequest

	// 新增网卡信息集合

	CreateMultiNicParams *CreateMultiNicParams `json:"CreateMultiNicParams,omitempty" name:"CreateMultiNicParams"`
}

func (r *CreateMultiNicDefinitionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMultiNicDefinitionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeAffinity struct {
	// 强制亲和参数

	Required []*Selector `json:"Required,omitempty" name:"Required"`
	// 尽量亲和权重, 1-100

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// 尽量亲和参数

	Preferred []*Selector `json:"Preferred,omitempty" name:"Preferred"`
}

type SearchProductInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *SearchProductInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchProductInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerSpecialRecycleLanIPExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ServerSpecialRecycleLanIPInfo

		ServerSpecialRecycleLanIPInfo []*ServerSpecialAllocateLanIPInfo `json:"ServerSpecialRecycleLanIPInfo,omitempty" name:"ServerSpecialRecycleLanIPInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ServerSpecialRecycleLanIPExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerSpecialRecycleLanIPExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Ingress struct {
	// ingress唯一ID

	Name *string `json:"Name,omitempty" name:"Name"`
	// ingress所属namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 固定为"负载均衡"

	Type *string `json:"Type,omitempty" name:"Type"`
	// VIP

	VIP *string `json:"VIP,omitempty" name:"VIP"`
	// 后端service, 格式为service:port，多个以逗号隔开

	Backend *string `json:"Backend,omitempty" name:"Backend"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// owner名字, 可以为空

	OwnerName *string `json:"OwnerName,omitempty" name:"OwnerName"`
	// owner类型, 可以为空

	OwnerKind *string `json:"OwnerKind,omitempty" name:"OwnerKind"`
	// 后端服务名称列表

	Services []*string `json:"Services,omitempty" name:"Services"`
	// 后端服务端口列表，与服务名称一一对应

	ServicePorts []*string `json:"ServicePorts,omitempty" name:"ServicePorts"`
	// 路由控制器

	IngressClass *string `json:"IngressClass,omitempty" name:"IngressClass"`
	// 网络类型

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// 协议列表

	Protocols []*string `json:"Protocols,omitempty" name:"Protocols"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 规则列表

	Rules []*IngressRule `json:"Rules,omitempty" name:"Rules"`
	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 应用

	Application *string `json:"Application,omitempty" name:"Application"`
	// 项目显示名称

	ProjectShowName *string `json:"ProjectShowName,omitempty" name:"ProjectShowName"`
	// Ingress JSON字符串

	JSON *string `json:"JSON,omitempty" name:"JSON"`
	// 是否是默认路由控制器

	Default *bool `json:"Default,omitempty" name:"Default"`
}

type Status struct {
	// 状态正常的实例个数

	NormalCount *int64 `json:"NormalCount,omitempty" name:"NormalCount"`
	// 状态异常的实例个数

	AbnormalCount *int64 `json:"AbnormalCount,omitempty" name:"AbnormalCount"`
	// 致命错误的实例个数

	FatalCount *int64 `json:"FatalCount,omitempty" name:"FatalCount"`
	// 状态未知的实例个数

	UnknownCount *int64 `json:"UnknownCount,omitempty" name:"UnknownCount"`
}

type DescribeSystemComponentsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeSystemComponentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSystemComponentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogInfo struct {
	// 纳秒级时间戳

	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
	// 原始日志

	OriginalLog *string `json:"OriginalLog,omitempty" name:"OriginalLog"`
}

type VirtualMachine struct {
	// vmid

	VmId *string `json:"VmId,omitempty" name:"VmId"`
	// 虚拟机名称

	VmName *string `json:"VmName,omitempty" name:"VmName"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 项目

	Project *string `json:"Project,omitempty" name:"Project"`
	// 主机ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// cpu

	Cpu *float64 `json:"Cpu,omitempty" name:"Cpu"`
	// 内存

	Memory *float64 `json:"Memory,omitempty" name:"Memory"`
	// 镜像

	Image *string `json:"Image,omitempty" name:"Image"`
	// ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 创建时间

	CreateTimestamp *string `json:"CreateTimestamp,omitempty" name:"CreateTimestamp"`
	// 系统盘

	SystemDisk *string `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 健康状态

	Health *bool `json:"Health,omitempty" name:"Health"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 操作系统

	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`
}

type DeleteSvrNicAllocationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSvrNicAllocationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSvrNicAllocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件列表

		Events []*Event `json:"Events,omitempty" name:"Events"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *ListRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchDashboardsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchDashboardsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchDashboardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命中条件总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 查询耗时 单位 ms

		Took *float64 `json:"Took,omitempty" name:"Took"`
		// 命中数据信息

		Hits []*Hit `json:"Hits,omitempty" name:"Hits"`
		// 聚合信息

		Aggregations *Aggregations `json:"Aggregations,omitempty" name:"Aggregations"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOutbandInfoExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// OutbandInfo

		OutbandInfo []*OutbandInfo `json:"OutbandInfo,omitempty" name:"OutbandInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOutbandInfoExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOutbandInfoExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVirtualMachineYamlRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 虚拟机ID

	VmId *string `json:"VmId,omitempty" name:"VmId"`
}

func (r *DescribeVirtualMachineYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVirtualMachineYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListVolumesRequest struct {
	*tchttp.BaseRequest

	// kubernetes集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每次查询数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 按某个字段排序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否降序排序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 搜索条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 项目Id

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ListVolumesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListVolumesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationSheet struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建时间

	CreationTimestamp *string `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
	// 启动时间

	StartTimestamp *string `json:"StartTimestamp,omitempty" name:"StartTimestamp"`
	// 结束时间

	FinishTimestamp *string `json:"FinishTimestamp,omitempty" name:"FinishTimestamp"`
	// 当前状态, Ready/Running/Paused/Completed/Acked/Failed/Aborted/Rollbacked, 分别对应待运行/运行中/暂停/完成未确认/完成已确认/失败终止/人工终止/已回滚

	Status *string `json:"Status,omitempty" name:"Status"`
	// 上次操作人员uin

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 回滚时间

	RollbackTimestamp *string `json:"RollbackTimestamp,omitempty" name:"RollbackTimestamp"`
	// 变更单内容

	YAML *string `json:"YAML,omitempty" name:"YAML"`
}

type ModifyRAIDParam struct {
	// 编号

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateCustomScriptSetRequest struct {
	*tchttp.BaseRequest

	// 需要创建的自定义脚本集名

	CustomScriptSetName *string `json:"CustomScriptSetName,omitempty" name:"CustomScriptSetName"`
	// 需要加入的自定义脚本名列表

	CustomScriptName []*string `json:"CustomScriptName,omitempty" name:"CustomScriptName"`
	// 脚本集类型(0:初始化脚本,1:后置脚本)

	CustomScriptSetType *string `json:"CustomScriptSetType,omitempty" name:"CustomScriptSetType"`
}

func (r *CreateCustomScriptSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomScriptSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHPARequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 名字空间及名字的组合列表

	NamespacedNames []*NamespacedName `json:"NamespacedNames,omitempty" name:"NamespacedNames"`
}

func (r *DeleteHPARequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHPARequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDeploymentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 部署列表

		DeploymentSet []*Deployment `json:"DeploymentSet,omitempty" name:"DeploymentSet"`
		// 状态列表

		FilterEntity *DeploymentFilterEntity `json:"FilterEntity,omitempty" name:"FilterEntity"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListDeploymentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDeploymentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// region信息

		Regions []*Region `json:"Regions,omitempty" name:"Regions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNodeLabelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateNodeLabelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateNodeLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSvrNicParams struct {
	// 服务器SN

	Sns []*string `json:"Sns,omitempty" name:"Sns"`
	// 服务器

	NewDefinitionName *string `json:"NewDefinitionName,omitempty" name:"NewDefinitionName"`
}

type DescribeClusterVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// kubernetes版本

		KubernetesVersion *string `json:"KubernetesVersion,omitempty" name:"KubernetesVersion"`
		// 容器运行时名字及版本, 例如docker://18.0.0

		ContainerRuntimeVersion *string `json:"ContainerRuntimeVersion,omitempty" name:"ContainerRuntimeVersion"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIngressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群ID

		Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
		// 所属项目

		Project *string `json:"Project,omitempty" name:"Project"`
		// Ingress JSON字符串

		Ingress *string `json:"Ingress,omitempty" name:"Ingress"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIngressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOutBandIpResourceIpDetailExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件；仅支持ip和网段查询；示例："Filters":[{"Name":"Ip","Values":["127.0.0.1"]}]或"Filters":[{"Name":"Segment","Values":["127.0.0.1/25"]}]

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 检索类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeOutBandIpResourceIpDetailExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOutBandIpResourceIpDetailExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportClusterRequest struct {
	*tchttp.BaseRequest

	// 集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// ApiServer

	ApiServer *string `json:"ApiServer,omitempty" name:"ApiServer"`
	// CertFile

	CertFile *string `json:"CertFile,omitempty" name:"CertFile"`
	// Token，

	AuthToken *string `json:"AuthToken,omitempty" name:"AuthToken"`
	// KubeConfig

	KubeConfig *string `json:"KubeConfig,omitempty" name:"KubeConfig"`
}

func (r *ImportClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTemplateRes struct {
	// 主账号

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 子账号

	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 模板类别

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
	// 模板名称

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 模板id

	Name *string `json:"Name,omitempty" name:"Name"`
}

type MachineVersion struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 类型，master or worker

	VersionType *string `json:"VersionType,omitempty" name:"VersionType"`
}

type TreeNodeMiddleware struct {
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 服务类型

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type DescribeLogConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志配置列表

		Configs []*ConfigInfo `json:"Configs,omitempty" name:"Configs"`
		// 日志总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodesOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// node 网格列表

		NodeStatus []*NodeStatus `json:"NodeStatus,omitempty" name:"NodeStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNodesOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodesOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VM struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// Vlan

	Vlan *string `json:"Vlan,omitempty" name:"Vlan"`
	// 默认内网IP

	DefaultLanIP *string `json:"DefaultLanIP,omitempty" name:"DefaultLanIP"`
}

type DescribeServerHarddisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数量

		ReturnNum *int64 `json:"ReturnNum,omitempty" name:"ReturnNum"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回硬盘列表

		ServerHardDiskSet []*ServerHardDisk `json:"ServerHardDiskSet,omitempty" name:"ServerHardDiskSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServerHarddisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServerHarddisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLabelRequest struct {
	*tchttp.BaseRequest

	// "Scheme":"label"

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 系统ID

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 系统Key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 操作类型："Op":"modify"

	Op *string `json:"Op,omitempty" name:"Op"`
	// 修改的标签内容

	ModifyLibelSet *ModifyLabelInfo `json:"ModifyLibelSet,omitempty" name:"ModifyLibelSet"`
}

func (r *ModifyLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateResourceObjectRequest struct {
	*tchttp.BaseRequest

	// 资源归属信息

	Params *ResourceOwnerQueryParam `json:"Params,omitempty" name:"Params"`
	// 服务树节点标签

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 更新信息

	Command *UpdateResourceCommand `json:"Command,omitempty" name:"Command"`
}

func (r *UpdateResourceObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateResourceObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageBuildingTaskRequest struct {
	*tchttp.BaseRequest

	// 待删除的任务, 需要设置Name

	ImageBuildingTask *ImageBuildingTask `json:"ImageBuildingTask,omitempty" name:"ImageBuildingTask"`
}

func (r *DeleteImageBuildingTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImageBuildingTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVirtualMachineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVirtualMachineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVirtualMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartImageBuildingTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartImageBuildingTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartImageBuildingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteStorageClassRequest struct {
	*tchttp.BaseRequest

	// kubernetes集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 待删除StorageClass名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteStorageClassRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteStorageClassRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterFilterEntity struct {
	// k8s版本

	KubernetsVersion []*string `json:"KubernetsVersion,omitempty" name:"KubernetsVersion"`
	// 状态，Initializing、Running、Failed、Upgrading、Terminating

	Phase []*string `json:"Phase,omitempty" name:"Phase"`
	// 集群类型，TCS(业务集群)、Global（管控集群）、TSF（TSF集群）、Imported（纳管集群）

	ClusterType []*string `json:"ClusterType,omitempty" name:"ClusterType"`
}

type CreateSilence struct {
	// 规则名，工单名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 屏蔽级别

	Level *string `json:"Level,omitempty" name:"Level"`
	// 屏蔽开始时间

	StartsAt *string `json:"StartsAt,omitempty" name:"StartsAt"`
	// 屏蔽结束时间

	EndsAt *string `json:"EndsAt,omitempty" name:"EndsAt"`
	// 产品

	Product *string `json:"Product,omitempty" name:"Product"`
	// 对象信息

	Objects []*LabelPair `json:"Objects,omitempty" name:"Objects"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 自定义变更人员，默认是当前用户，多个用户间逗号分隔

	Updater *string `json:"Updater,omitempty" name:"Updater"`
	// 策略和规则信息

	Labels []*SilenceGroup `json:"Labels,omitempty" name:"Labels"`
}

type StatefulSetFilterEntity struct {
	// 状态列表

	PhaseSet []*string `json:"PhaseSet,omitempty" name:"PhaseSet"`
}

type CheckMachineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败原因

		FailReason *FailReason `json:"FailReason,omitempty" name:"FailReason"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckMachineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRuleGroupTempleRequest struct {
	*tchttp.BaseRequest

	// 策略模板 id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetRuleGroupTempleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRuleGroupTempleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogsRequest struct {
	*tchttp.BaseRequest

	// Meta信息

	Metas []*Meta `json:"Metas,omitempty" name:"Metas"`
	// 日志名，可以选多个

	Name []*string `json:"Name,omitempty" name:"Name"`
	// 时间范围起始，纳秒时间戳，默认当前时间(总数精确到秒级，如果需要准确查询时间范围内总数，时间起止改为 秒级时间戳*10^9)

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 时间范围结束，纳秒时间戳，默认一小时前总数精确到秒级，如果需要准确查询时间范围内总数，时间起止改为 秒级时间戳*10^9)

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 返回记录条数，默认值：500，最大值：1000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 聚集查询

	Aggs *Aggs `json:"Aggs,omitempty" name:"Aggs"`
	// 排序

	Sort []*SortInfo `json:"Sort,omitempty" name:"Sort"`
	// 查询方式，0: LogQL

	QueryType *uint64 `json:"QueryType,omitempty" name:"QueryType"`
	// 数据类型  默认 0：日志  1：事件

	DataType *uint64 `json:"DataType,omitempty" name:"DataType"`
	// 分步查询类型，为空时，所有数据都查。值为log_detail时，只查日志详情，值为date_histogram时，只查直方图和总数

	StepQuery *string `json:"StepQuery,omitempty" name:"StepQuery"`
}

func (r *SearchLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageBuildingTaskRequest struct {
	*tchttp.BaseRequest

	// 构建任务

	ImageBuildingTask *ImageBuildingTask `json:"ImageBuildingTask,omitempty" name:"ImageBuildingTask"`
}

func (r *CreateImageBuildingTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateImageBuildingTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrUpdateRetentionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrUpdateRetentionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrUpdateRetentionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIgniterNodeExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点详细信息

		IgniterNodesSet []*IgniterNodesSet `json:"IgniterNodesSet,omitempty" name:"IgniterNodesSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIgniterNodeExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIgniterNodeExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomScripts struct {
	// 自定义脚本名

	CustomScriptName *string `json:"CustomScriptName,omitempty" name:"CustomScriptName"`
	// 自定义脚本加入的脚本集

	CustomScriptSetName *string `json:"CustomScriptSetName,omitempty" name:"CustomScriptSetName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type MdComponent struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无特殊含义

	Type *string `json:"Type,omitempty" name:"Type"`
	// 状态Normal，Abnormal，NoData，Warning

	Status *string `json:"Status,omitempty" name:"Status"`
	// 运行阶段

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 类别

	Class *string `json:"Class,omitempty" name:"Class"`
	// 异常原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type RelocateFinishServer struct {
	// 资产ID

	SvrAssetId *string `json:"SvrAssetId,omitempty" name:"SvrAssetId"`
	// IDC名称

	SvrRelocateIdcName *string `json:"SvrRelocateIdcName,omitempty" name:"SvrRelocateIdcName"`
	// 机架名称

	SvrRelocateRackName *string `json:"SvrRelocateRackName,omitempty" name:"SvrRelocateRackName"`
	// 机架位置

	SvrRelocatePosId *string `json:"SvrRelocatePosId,omitempty" name:"SvrRelocatePosId"`
}

type DeleteQueryRecordRequest struct {
	*tchttp.BaseRequest

	// 用户id

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 记录id

	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`
}

func (r *DeleteQueryRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteQueryRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDcosFunctionServiceRelationRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryDcosFunctionServiceRelationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDcosFunctionServiceRelationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDevOpsStatusRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 场景id

	RemediationName *string `json:"RemediationName,omitempty" name:"RemediationName"`
	// 开启：Enable，关闭：Disabled

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDevOpsStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDevOpsStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRoutesRequest struct {
	*tchttp.BaseRequest

	// 路由搜索条件

	Params *SearchRoutes `json:"Params,omitempty" name:"Params"`
}

func (r *SearchRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMetaMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态码

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 指标信息

		Data *MetaMetric `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMetaMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMetaMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashAclContent struct {
	// 仪表盘ID

	DashboardId *int64 `json:"DashboardId,omitempty" name:"DashboardId"`
	// 仪表盘目录ID

	FolderId *int64 `json:"FolderId,omitempty" name:"FolderId"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// 更新时间

	Updated *string `json:"Updated,omitempty" name:"Updated"`
	// 用户ID

	UserId *int64 `json:"UserId,omitempty" name:"UserId"`
	// 用户名称

	UserLogin *string `json:"UserLogin,omitempty" name:"UserLogin"`
	// 用户邮件地址

	UserEmail *string `json:"UserEmail,omitempty" name:"UserEmail"`
	// 角色

	Role *string `json:"Role,omitempty" name:"Role"`
	// 权限

	Permission *string `json:"Permission,omitempty" name:"Permission"`
	// 权限名称

	PermissionName *string `json:"PermissionName,omitempty" name:"PermissionName"`
	// 仪表盘或仪表盘目录Unique ID

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 仪表盘或仪表盘目录名称

	Title *string `json:"Title,omitempty" name:"Title"`
	// 是否继续来的权限

	Inherited *bool `json:"Inherited,omitempty" name:"Inherited"`
}

type Scenes struct {
	// 场景名称

	RemediationName *string `json:"RemediationName,omitempty" name:"RemediationName"`
	// 动作名称

	OperationName *string `json:"OperationName,omitempty" name:"OperationName"`
	// 动作列表

	ActionArgs []*ActionArgs `json:"ActionArgs,omitempty" name:"ActionArgs"`
}

type DescribeNodeAttributesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
}

func (r *DescribeNodeAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SupportedFunction struct {
	// IP模块支持

	IPModule *bool `json:"IPModule,omitempty" name:"IPModule"`
	// 带外模块支持

	OutbandModule *bool `json:"OutbandModule,omitempty" name:"OutbandModule"`
	// 其他模块支持

	OtherModule *bool `json:"OtherModule,omitempty" name:"OtherModule"`
}

type TreeSubNode struct {
	// 产品名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type ImageFilters struct {
	// 过滤键名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ModifyTopologyStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopologyStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTopologyStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyStatefulSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyStatefulSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStatefulSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterByAdminRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteClusterByAdminRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterByAdminRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIngressClassRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DeleteIngressClassRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIngressClassRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPublicKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公钥, PEM格式

		PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPublicKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPublicKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerSpecialRecycleLanIPExRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 请求串类型

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 请求方id

	SystemId *string `json:"SystemId,omitempty" name:"SystemId"`
	// 请求方key

	SystemKey *string `json:"SystemKey,omitempty" name:"SystemKey"`
	// 请求动作

	Op *string `json:"Op,omitempty" name:"Op"`
	// 用户

	User *string `json:"User,omitempty" name:"User"`
	// 物理服务器批量回收相应虚拟内网段IP接口

	ServersSpecialRecycleLanIp []*ServersSpecialRecycleLanIp `json:"ServersSpecialRecycleLanIp,omitempty" name:"ServersSpecialRecycleLanIp"`
}

func (r *ServerSpecialRecycleLanIPExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ServerSpecialRecycleLanIPExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetricDefine struct {
	// core和all取值

	TemplateMetricType *string `json:"TemplateMetricType,omitempty" name:"TemplateMetricType"`
	// 指标名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// k8s里面的name

	K8sName *string `json:"K8sName,omitempty" name:"K8sName"`
	// 选择器

	Selector []*string `json:"Selector,omitempty" name:"Selector"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 指标类别，如cpu，mem，disk

	MetricType *string `json:"MetricType,omitempty" name:"MetricType"`
	// 指标维度，如cluster, node, pod,

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
	// 指标单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 指标中文名

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
}

type GetCronJobRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *GetCronJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCronJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPodInstantMetricDataRequest struct {
	*tchttp.BaseRequest

	// 查询指标名称，多个用逗号分隔

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 用于过滤node的标签集合，key=value形式，多个用逗号分隔

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 查询时间点

	Start *string `json:"Start,omitempty" name:"Start"`
	// 精度，如1m

	Step *string `json:"Step,omitempty" name:"Step"`
}

func (r *GetPodInstantMetricDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPodInstantMetricDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceYAMLRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 服务名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 更新的yaml

	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
}

func (r *UpdateServiceYAMLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateServiceYAMLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KVPair struct {
	// 监控指标

	Key *string `json:"Key,omitempty" name:"Key"`
	// 监控数据

	Value *string `json:"Value,omitempty" name:"Value"`
}
