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

package v20200413

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DoStopComponentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoStopComponentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoStopComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostDetailRsp struct {
	// cpu_count

	// data_center

	// host_state

	// os_arch

	// os_type

	// rack

	// serial_number

	// total_mem
}

type GetBkcmdbAllHostRequest struct {
	*tchttp.BaseRequest
}

func (r *GetBkcmdbAllHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBkcmdbAllHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadInspectionRspResultCase struct {
	// cases

	// status
}

type QueryProductHealthStateRsp struct {
	// Product

	Product *string `json:"Product,omitempty" name:"Product"`
	// Cluster

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// Status

	Status *string `json:"Status,omitempty" name:"Status"`
	// Message

	Message *string `json:"Message,omitempty" name:"Message"`
	// Metrics

	Metrics []*QueryProductHealthStateRspMetrics `json:"Metrics,omitempty" name:"Metrics"`
}

type QueryProductFailureMigrateTaskDetailRequest struct {
	*tchttp.BaseRequest

	// TaskUUID

	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
}

func (r *QueryProductFailureMigrateTaskDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductFailureMigrateTaskDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPrometheusRequest struct {
	*tchttp.BaseRequest

	// Query https://prometheus.io/docs/prometheus/2.19/querying/api/

	Query *string `json:"Query,omitempty" name:"Query"`
	// 时间戳, 毫秒

	Start *int64 `json:"Start,omitempty" name:"Start"`
	// 时间戳, 毫秒

	End *int64 `json:"End,omitempty" name:"End"`
	// Step, 秒

	Step *int64 `json:"Step,omitempty" name:"Step"`
}

func (r *QueryPrometheusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPrometheusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetClusterConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCurrentLicenseUsageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 裸容量, 单位TB

		RawCapacity *int64 `json:"RawCapacity,omitempty" name:"RawCapacity"`
		// 节点数量

		StorageHostCount *int64 `json:"StorageHostCount,omitempty" name:"StorageHostCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCurrentLicenseUsageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCurrentLicenseUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetMailAlertConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetMailAlertConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetMailAlertConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoRemoveComponentRequest struct {
	*tchttp.BaseRequest

	// 节点数组

	HostNameList []*string `json:"HostNameList,omitempty" name:"HostNameList"`
	// 组件名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
}

func (r *DoRemoveComponentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoRemoveComponentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCurrentLicenseUsageRequest struct {
	*tchttp.BaseRequest
}

func (r *GetCurrentLicenseUsageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCurrentLicenseUsageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetEvacuationProgressRequest struct {
	*tchttp.BaseRequest

	// 存储池 Id

	Dssid *int64 `json:"Dssid,omitempty" name:"Dssid"`
}

func (r *GetEvacuationProgressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetEvacuationProgressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetObjectStorageBucketQuotaRequest struct {
	*tchttp.BaseRequest

	// 存储桶名称

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
}

func (r *GetObjectStorageBucketQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetObjectStorageBucketQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRequestListRequest struct {
	*tchttp.BaseRequest

	// 分页起始值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页长度, 默认20

	Length *int64 `json:"Length,omitempty" name:"Length"`
}

func (r *GetRequestListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRequestListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostIfListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 网卡列表

		IfList []*GetHostIfListResp `json:"IfList,omitempty" name:"IfList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHostIfListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostIfListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRgwLogDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRgwLogDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRgwLogDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoRestartDiskRequest struct {
	*tchttp.BaseRequest

	// 节点IP

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 硬盘名称

	Device *string `json:"Device,omitempty" name:"Device"`
}

func (r *DoRestartDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoRestartDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAlertCountRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAlertCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostComponentDssidListRequest struct {
	*tchttp.BaseRequest

	// 节点列表

	HostNames []*string `json:"HostNames,omitempty" name:"HostNames"`
}

func (r *GetHostComponentDssidListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostComponentDssidListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductDataReplicationStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品数据同步状态查询接口返回值, CSP不适用该接口, 返回固定内容

		Data *QueryProductDataReplicationStateRsp `json:"Data,omitempty" name:"Data"`
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

type DoCreatePoolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoCreatePoolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoCreatePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetObjectStoragePoolListRequest struct {
	*tchttp.BaseRequest
}

func (r *GetObjectStoragePoolListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetObjectStoragePoolListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductHealthStateRspMetrics struct {
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Status

	Status *string `json:"Status,omitempty" name:"Status"`
	// Message

	Message *string `json:"Message,omitempty" name:"Message"`
}

type GetBkcmdbStorageHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点列表

		Hosts []*GetBkcmdbHostResp `json:"Hosts,omitempty" name:"Hosts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBkcmdbStorageHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBkcmdbStorageHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostCandidateRequest struct {
	*tchttp.BaseRequest

	// 组件名称数组

	ComponentNameList []*string `json:"ComponentNameList,omitempty" name:"ComponentNameList"`
}

func (r *GetHostCandidateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostCandidateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHttpQpsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Overview Qps概览数据

		Data *QpsData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHttpQpsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHttpQpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRgwLogDetailRequest struct {
	*tchttp.BaseRequest

	// 节点 IP

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 日志文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 是否base64编码

	Base64 *bool `json:"Base64,omitempty" name:"Base64"`
}

func (r *GetRgwLogDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRgwLogDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterConstStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterConstStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterConstStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryInitClusterRequest struct {
	*tchttp.BaseRequest
}

func (r *RetryInitClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryInitClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadInspectionRspResultCaseHost struct {
	// host

	// status

	// stderr

	// stdout
}

type FlowControlConfig struct {
	// 是否开启

	// 写iops阈值

	// 读iops阈值

	// 写带宽阈值

	// 读带宽阈值

	// post iops阈值

	// put iops阈值

	// delete iops阈值

	// get iops阈值

	// head iops阈值

	// options iops阈值
}

type CleanInspectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CleanInspectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CleanInspectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoRestartDiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoRestartDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoRestartDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostComponentStateListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHostComponentStateListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostComponentStateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Host struct {
	// 节点名称

	// 数据中心

	// 机架位置

	// 序列号
}

type DoStopComponentRequest struct {
	*tchttp.BaseRequest

	// 节点数组

	HostNameList []*string `json:"HostNameList,omitempty" name:"HostNameList"`
	// 组件名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
}

func (r *DoStopComponentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoStopComponentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterConstStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *GetClusterConstStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterConstStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHostOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInitClusterRequestIdRequest struct {
	*tchttp.BaseRequest
}

func (r *GetInitClusterRequestIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInitClusterRequestIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterComponents struct {
	// 监控节点A

	CLUSTERMON_A *string `json:"CLUSTERMON_A,omitempty" name:"CLUSTERMON_A"`
	// 监控节点B

	CLUSTERMON_B *string `json:"CLUSTERMON_B,omitempty" name:"CLUSTERMON_B"`
	// 监控节点C

	CLUSTERMON_C *string `json:"CLUSTERMON_C,omitempty" name:"CLUSTERMON_C"`
	// 指标收集节点，这里与监控节点C相同

	METRICS_COLLECTOR *string `json:"METRICS_COLLECTOR,omitempty" name:"METRICS_COLLECTOR"`
}

type CreateClusterConfig struct {
	// NTP服务器

	// 地域名称

	// 地域Id

	// 可用区

	// public网段

	// cluster网段

	// CAM 身份认证与鉴权 URI

	// CAM 账号查询 URI

	// CAM 策略查询 URI

	// CAM 策略修改 URI

	// 计费上报 URI

	// 对象存储开户 URI

	// 存储主域名

	// 对象存储备域名

	// tagGetUri

	// tagSetUri

	// cosBaradNwsUrl

	// mgmtStatisticsPushgatewayUrl

	// cryptTencentKmsRegion

	// cryptTencentKmsEndpoint

	// cryptTencentStsEndpoint

	// cryptTencentOcloudKmsEndpoint
}

type CleanInspectionRequest struct {
	*tchttp.BaseRequest
}

func (r *CleanInspectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CleanInspectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoExpandPoolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoExpandPoolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoExpandPoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoStopHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoStopHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoStopHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostDisksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Disks

		Disks *GetHostDisksRsp `json:"Disks,omitempty" name:"Disks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHostDisksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoRestartComponentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoRestartComponentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoRestartComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCdcUserSuspendStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户禁用状态。有效值：0，1，2，4分别表示enable，normal suspend，deactivate suspend，isolate suspend

		Suspend *int64 `json:"Suspend,omitempty" name:"Suspend"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCdcUserSuspendStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCdcUserSuspendStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryRequestByIdRequest struct {
	*tchttp.BaseRequest

	// request id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *RetryRequestByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryRequestByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBackupGlobalConfigRequest struct {
	*tchttp.BaseRequest

	// 备份临时数据存储桶名称, 不填时传入""

	AsynctaskBucket *string `json:"AsynctaskBucket,omitempty" name:"AsynctaskBucket"`
	// 格式如0-5,22-24，表示0～5点，22～24点执行备份. 传入""代表全天,

	Timerange *string `json:"Timerange,omitempty" name:"Timerange"`
	// DNS服务器地址, 可传"".

	Resolver *string `json:"Resolver,omitempty" name:"Resolver"`
	// 备份策略模式, 可传null.  STORAGE_FIRST 业务优先, STANDARD 均衡, BACKUP_FIRST 备份优先

	BackupStrategy *string `json:"BackupStrategy,omitempty" name:"BackupStrategy"`
	// 备份最大并发线程, 可传null. 可选值20～500

	BackupConcurrency *int64 `json:"BackupConcurrency,omitempty" name:"BackupConcurrency"`
	// 每秒最大备份数, 可传null

	BackupTps *int64 `json:"BackupTps,omitempty" name:"BackupTps"`
}

func (r *SetBackupGlobalConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBackupGlobalConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitClusterRequest struct {
	*tchttp.BaseRequest

	// 初始化集群组件对象

	Components *ClusterComponents `json:"Components,omitempty" name:"Components"`
}

func (r *InitClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoEvacuateComponentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoEvacuateComponentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoEvacuateComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPoolLatestOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 所有存储池的概览信息

		Data []*PoolOverviewDetail `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPoolLatestOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPoolLatestOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRecoveryControlConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetRecoveryControlConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRecoveryControlConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPoolHostsRequest struct {
	*tchttp.BaseRequest

	// 存储池 dssid

	Dssid *string `json:"Dssid,omitempty" name:"Dssid"`
}

func (r *GetPoolHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPoolHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoRetryPoolRequestRequest struct {
	*tchttp.BaseRequest

	// 创建存储池出错时的 request id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DoRetryPoolRequestRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoRetryPoolRequestRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDnsDomainSettingsRequest struct {
	*tchttp.BaseRequest
}

func (r *GetDnsDomainSettingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDnsDomainSettingsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOSUsageOverviewGroupByUserIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// -

		Data []*GetOSUsageOverviewGroupByUserIdItem `json:"Data,omitempty" name:"Data"`
		// -

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOSUsageOverviewGroupByUserIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOSUsageOverviewGroupByUserIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostDisksRspDisk struct {
	// device

	// rotational

	// usage

	// size

	// role

	// status

	// componentName

	// relatedComponents

	// lighting

	// smart

	// 挂载点
}

type DoDeletePoolRequest struct {
	*tchttp.BaseRequest

	// 存储池名称

	PlacementName *string `json:"PlacementName,omitempty" name:"PlacementName"`
}

func (r *DoDeletePoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoDeletePoolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoReInitializeDiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoReInitializeDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoReInitializeDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostComponentStateListRequest struct {
	*tchttp.BaseRequest

	// 节点数组

	HostNames []*string `json:"HostNames,omitempty" name:"HostNames"`
}

func (r *GetHostComponentStateListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostComponentStateListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRequestInfoByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRequestInfoByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRequestInfoByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAutoRecoveryRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAutoRecoveryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAutoRecoveryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostCountRequest struct {
	*tchttp.BaseRequest
}

func (r *GetHostCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群节点数量

		HostNum *int64 `json:"HostNum,omitempty" name:"HostNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHostCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoAddHostsRequest struct {
	*tchttp.BaseRequest

	// 节点信息列表

	Hosts []*Host `json:"Hosts,omitempty" name:"Hosts"`
	// root账户密码

	RootPassword *string `json:"RootPassword,omitempty" name:"RootPassword"`
	// 端口

	SshPort *int64 `json:"SshPort,omitempty" name:"SshPort"`
}

func (r *DoAddHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoAddHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoDeleteHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoDeleteHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoDeleteHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetRgwLogLevelRequest struct {
	*tchttp.BaseRequest

	// 节点IP

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 日志等级

	LogRate *int64 `json:"LogRate,omitempty" name:"LogRate"`
}

func (r *SetRgwLogLevelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetRgwLogLevelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DnsDomainSettings struct {
	// 域名列表

	// 静态网站域名前缀
}

type MailAlertConfig struct {
	// SMTP 认证是否开启

	// SMTP 帐号

	// SMTP 密码

	// 邮件收件人

	// 邮件发件人

	// SMTP 服务器

	// SMTP 端口

	// SslTrust

	// StartTLS 加密是否开启
}

type DoCreatePoolRequest struct {
	*tchttp.BaseRequest

	// 创建存储池所需配置项

	CreatePoolConfig *CreatePoolConfig `json:"CreatePoolConfig,omitempty" name:"CreatePoolConfig"`
}

func (r *DoCreatePoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoCreatePoolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBkcmdbMonitorHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点列表

		Hosts []*GetBkcmdbHostResp `json:"Hosts,omitempty" name:"Hosts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBkcmdbMonitorHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBkcmdbMonitorHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostRolesForDssidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Hosts

		Hosts []*GetHostRolesForDssidRsp `json:"Hosts,omitempty" name:"Hosts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHostRolesForDssidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostRolesForDssidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAutoRecoveryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetAutoRecoveryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAutoRecoveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecoveryControlConfig struct {
	// 是否开启恢复控制

	// 恢复控制模式

	// 恢复控制高峰时间段
}

type DoMigrateComponentRequest struct {
	*tchttp.BaseRequest

	// 被更换的节点

	FromHostName *string `json:"FromHostName,omitempty" name:"FromHostName"`
	// 新节点

	ToHostName *string `json:"ToHostName,omitempty" name:"ToHostName"`
	// 组件名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
}

func (r *DoMigrateComponentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoMigrateComponentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCdcUserSuspendStatusRequest struct {
	*tchttp.BaseRequest

	// CdcUserUin

	CdcUserUin *int64 `json:"CdcUserUin,omitempty" name:"CdcUserUin"`
}

func (r *GetCdcUserSuspendStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCdcUserSuspendStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDefaultPlacementResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetDefaultPlacementResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDefaultPlacementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetObjectStorageBucketQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetObjectStorageBucketQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetObjectStorageBucketQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductFailureMigrateRsp struct {
	// TaskUUID

	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
}

type SetBucketFlowControlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CspRequestId

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetBucketFlowControlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBucketFlowControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetClusterConfigRequest struct {
	*tchttp.BaseRequest

	// 集群配置

	Config *ClusterConfig `json:"Config,omitempty" name:"Config"`
}

func (r *SetClusterConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetClusterConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartInspectionRequest struct {
	*tchttp.BaseRequest
}

func (r *StartInspectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartInspectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAlertsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Alerts

		Alerts []*GetAlertsRsp `json:"Alerts,omitempty" name:"Alerts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAlertsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetEvacuationProgressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetEvacuationProgressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetEvacuationProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBucketFlowControlRequest struct {
	*tchttp.BaseRequest

	// 存储桶名称

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 是否启用频控

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 写带宽阈值, 单位MB/s

	BrandI *int64 `json:"BrandI,omitempty" name:"BrandI"`
	// 读带宽阈值, 单位MB/s

	BrandO *int64 `json:"BrandO,omitempty" name:"BrandO"`
	// post iops阈值

	IopsPost *int64 `json:"IopsPost,omitempty" name:"IopsPost"`
	// put iops阈值

	IopsPut *int64 `json:"IopsPut,omitempty" name:"IopsPut"`
	// delete iops阈值

	IopsDelete *int64 `json:"IopsDelete,omitempty" name:"IopsDelete"`
	// get iops阈值

	IopsGet *int64 `json:"IopsGet,omitempty" name:"IopsGet"`
	// head iops阈值

	IopsHead *int64 `json:"IopsHead,omitempty" name:"IopsHead"`
	// options iops阈值

	IopsOptions *int64 `json:"IopsOptions,omitempty" name:"IopsOptions"`
}

func (r *SetBucketFlowControlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBucketFlowControlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRgwLogLevelResp struct {
	// 日志级别
}

type GetObjectStorageUserQuotaRequest struct {
	*tchttp.BaseRequest

	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *GetObjectStorageUserQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetObjectStorageUserQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductFailureMigrateTaskDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// response

		Data *QueryProductFailureMigrateTaskDetailRsp `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryProductFailureMigrateTaskDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductFailureMigrateTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDnsDomainSettingsRequest struct {
	*tchttp.BaseRequest

	// 接入域名配置

	Config *DnsDomainSettings `json:"Config,omitempty" name:"Config"`
}

func (r *SetDnsDomainSettingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDnsDomainSettingsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *GetClusterStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInitClusterRequestIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建集群request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetInitClusterRequestIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInitClusterRequestIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSnmpConfigResp struct {
	// 端口

	// snmp server地址
}

type GetDnsDomainSettingsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接入域名配置

		Config *DnsDomainSettings `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDnsDomainSettingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDnsDomainSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostDiskOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHostDiskOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostDiskOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOSUsageOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// -

		Data []*GetOSUsageOverviewItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOSUsageOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOSUsageOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostDisksRsp struct {
	// disks

	// unplugged
}

type SetCdcUserSuspendStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCdcUserSuspendStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCdcUserSuspendStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetFlowControlConfigRequest struct {
	*tchttp.BaseRequest

	// 流控设置

	Config *FlowControlConfig `json:"Config,omitempty" name:"Config"`
}

func (r *SetFlowControlConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetFlowControlConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRecoveryControlConfigRsp struct {
	// enabled

	// mode

	// 时间段, [8,10]代表8~10点
}

type GetHostCandidateResp struct {
	// 候选安装组件列表

	// 节点IP

	// CPU数量

	// 内存大小

	// 系统类型

	// 系统架构
}

type GetRecoveryExpansionRequest struct {
	*tchttp.BaseRequest

	// 存储池 id 列表

	DssidList []*int64 `json:"DssidList,omitempty" name:"DssidList"`
}

func (r *GetRecoveryExpansionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRecoveryExpansionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRgwLogLevelRequest struct {
	*tchttp.BaseRequest

	// 节点IP

	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

func (r *GetRgwLogLevelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRgwLogLevelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetObjectStorageUserQuotaRequest struct {
	*tchttp.BaseRequest

	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 存储桶数量

	Buckets *int64 `json:"Buckets,omitempty" name:"Buckets"`
	// 存储对象数量

	Objects *int64 `json:"Objects,omitempty" name:"Objects"`
	// 存储容量

	Capacity *int64 `json:"Capacity,omitempty" name:"Capacity"`
}

func (r *SetObjectStorageUserQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetObjectStorageUserQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartInspectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartInspectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartInspectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type DoRemoveComponentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoRemoveComponentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoRemoveComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBackupGlobalConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 备份临时数据存储桶名称

		AsynctaskBucket *string `json:"AsynctaskBucket,omitempty" name:"AsynctaskBucket"`
		// 格式如0-5,22-24，表示0～5点，22～24点执行备份. 空代表全天

		Timerange *string `json:"Timerange,omitempty" name:"Timerange"`
		// DNS服务器地址

		Resolver *string `json:"Resolver,omitempty" name:"Resolver"`
		// 备份策略模式, 可传null.  STORAGE_FIRST 业务优先, STANDARD 均衡, BACKUP_FIRST 备份优先

		BackupStrategy *string `json:"BackupStrategy,omitempty" name:"BackupStrategy"`
		// 备份最大并发线程

		BackupConcurrency *int64 `json:"BackupConcurrency,omitempty" name:"BackupConcurrency"`
		// 每秒最大备份数

		BackupTps *int64 `json:"BackupTps,omitempty" name:"BackupTps"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBackupGlobalConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBackupGlobalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOSUsageOverviewGroupByUserIdItem struct {
	// -

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// -

	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
	// -

	InBytesSum *string `json:"InBytesSum,omitempty" name:"InBytesSum"`
	// -

	OutBytesSum *string `json:"OutBytesSum,omitempty" name:"OutBytesSum"`
	// -

	ReadCountSum *string `json:"ReadCountSum,omitempty" name:"ReadCountSum"`
	// -

	WriteCountSum *string `json:"WriteCountSum,omitempty" name:"WriteCountSum"`
	// -

	SizeByteSum *string `json:"SizeByteSum,omitempty" name:"SizeByteSum"`
	// -

	NumObjectsSum *string `json:"NumObjectsSum,omitempty" name:"NumObjectsSum"`
	// -

	BucketNameCount *string `json:"BucketNameCount,omitempty" name:"BucketNameCount"`
}

type DoStartHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoStartHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoStartHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBkcmdbAccessHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBkcmdbAccessHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBkcmdbAccessHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGlobalInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *GetGlobalInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGlobalInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostIfListRequest struct {
	*tchttp.BaseRequest

	// 节点IP

	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

func (r *GetHostIfListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostIfListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryRequestByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryRequestByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryRequestByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAutoRecoveryRequest struct {
	*tchttp.BaseRequest

	// 是否启用

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

func (r *SetAutoRecoveryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAutoRecoveryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HttpCodeMapping struct {
	// 100个数

	// 101个数

	// 103个数

	// 200个数

	// 201个数

	// 202个数

	// 203个数

	// 204个数

	// 205个数

	// 206个数

	// 300个数

	// 301个数

	// 302个数

	// 303个数

	// 304个数

	// 305个数

	// 306个数

	// 307个数

	// 308个数

	// 400个数

	// 401个数

	// 402个数

	// 403个数

	// 404个数

	// 405个数

	// 406个数

	// 407个数

	// 408个数

	// 409个数

	// 410个数

	// 411个数

	// 412个数

	// 413个数

	// 414个数

	// 415个数

	// 416个数

	// 417个数

	// 418个数

	// 422个数

	// 425个数

	// 426个数

	// 428个数

	// 429个数

	// 431个数

	// 451个数

	// 500个数

	// 501个数

	// 502个数

	// 503个数

	// 504个数

	// 505个数

	// 506个数

	// 507个数

	// 508个数

	// 510个数

	// 511个数
}

type SetSnmpConfigRequest struct {
	*tchttp.BaseRequest

	// SNMP配置

	Config *SnmpConfig `json:"Config,omitempty" name:"Config"`
}

func (r *SetSnmpConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetSnmpConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBkcmdbAllHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBkcmdbAllHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBkcmdbAllHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostCandidateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点列表

		Hosts []*GetHostCandidateResp `json:"Hosts,omitempty" name:"Hosts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHostCandidateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostCandidateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetObjectStoragePoolListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 对象存储存储池列表

		Pools []*PoolDetail `json:"Pools,omitempty" name:"Pools"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetObjectStoragePoolListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetObjectStoragePoolListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductDataReplicationStateRequest struct {
	*tchttp.BaseRequest

	// Product

	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *QueryProductDataReplicationStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductDataReplicationStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOSUsageOverviewGroupByBucketItem struct {
	// -

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// -

	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
	// -

	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
	// -

	InBytesSum *string `json:"InBytesSum,omitempty" name:"InBytesSum"`
	// -

	OutBytesSum *string `json:"OutBytesSum,omitempty" name:"OutBytesSum"`
	// -

	ReadCountSum *string `json:"ReadCountSum,omitempty" name:"ReadCountSum"`
	// -

	WriteCountSum *string `json:"WriteCountSum,omitempty" name:"WriteCountSum"`
	// -

	SizeByteSum *string `json:"SizeByteSum,omitempty" name:"SizeByteSum"`
	// -

	NumObjectsSum *string `json:"NumObjectsSum,omitempty" name:"NumObjectsSum"`
}

type GetClusterLatestOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *GetClusterLatestOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterLatestOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterLatestOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterLatestOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterLatestOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSnmpConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetSnmpConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSnmpConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBucketFlowControlRequest struct {
	*tchttp.BaseRequest

	// Bucket名称

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
}

func (r *GetBucketFlowControlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBucketFlowControlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterOverviewRequest struct {
	*tchttp.BaseRequest

	// 开始时间戳

	From *int64 `json:"From,omitempty" name:"From"`
	// 结束时间戳

	To *int64 `json:"To,omitempty" name:"To"`
}

func (r *GetClusterOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInspectionStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *GetInspectionStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInspectionStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetMailAlertConfigRequest struct {
	*tchttp.BaseRequest

	// 邮件告警配置

	Config *MailAlertConfig `json:"Config,omitempty" name:"Config"`
}

func (r *SetMailAlertConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetMailAlertConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetSnmpConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetSnmpConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetSnmpConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoEvacuateComponentRequest struct {
	*tchttp.BaseRequest

	// 节点数组

	HostNameList []*string `json:"HostNameList,omitempty" name:"HostNameList"`
	// 组件名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
}

func (r *DoEvacuateComponentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoEvacuateComponentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoReleaseDiskRequest struct {
	*tchttp.BaseRequest

	// 节点IP

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 硬盘名称

	Device *string `json:"Device,omitempty" name:"Device"`
}

func (r *DoReleaseDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoReleaseDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadHostDisksDataRequest struct {
	*tchttp.BaseRequest

	// 节点名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 磁盘名称，多个磁盘中间以 逗号 相连

	Disks *string `json:"Disks,omitempty" name:"Disks"`
}

func (r *DownloadHostDisksDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadHostDisksDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostDisksRequest struct {
	*tchttp.BaseRequest

	// 节点IP

	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

func (r *GetHostDisksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostDisksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPoolByNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Pool

		Pool *GetPoolByNameRsp `json:"Pool,omitempty" name:"Pool"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPoolByNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPoolByNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostListRsp struct {
	// host_name

	// data_center

	// rack

	// component_state

	// host_state

	// serial_number

	// components
}

type QueryProductDataReplicationStateRsp struct {
	// Product

	Product *string `json:"Product,omitempty" name:"Product"`
	// 数据复制成功个数

	DataReplicationSuccess *int64 `json:"DataReplicationSuccess,omitempty" name:"DataReplicationSuccess"`
	// 数据复制失败个数

	DataReplicationFailed *int64 `json:"DataReplicationFailed,omitempty" name:"DataReplicationFailed"`
	// 实体列表，扩展字段

	EntityList []*string `json:"EntityList,omitempty" name:"EntityList"`
	// 状态最后上报时间（更新时间）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type SnmpConfig struct {
	// snmp server地址

	// snmp端口号
}

type GetOSUsageOverviewRequest struct {
	*tchttp.BaseRequest

	// 时间戳, 秒

	Start *int64 `json:"Start,omitempty" name:"Start"`
	// 时间戳, 秒

	End *int64 `json:"End,omitempty" name:"End"`
}

func (r *GetOSUsageOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOSUsageOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCdcUserSuspendStatusRequest struct {
	*tchttp.BaseRequest

	// CdcUserUin

	CdcUserUin *string `json:"CdcUserUin,omitempty" name:"CdcUserUin"`
	// 用户禁用状态。有效值：0，1，2，4分别表示enable，normal suspend，deactivate suspend，isolate suspend，cdc禁用用户时传递2即可

	Suspend *int64 `json:"Suspend,omitempty" name:"Suspend"`
}

func (r *SetCdcUserSuspendStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCdcUserSuspendStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoDeleteHostRequest struct {
	*tchttp.BaseRequest

	// 节点名

	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

func (r *DoDeleteHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoDeleteHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoStopHostRequest struct {
	*tchttp.BaseRequest

	// 节点名

	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

func (r *DoStopHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoStopHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHttpQpsRequest struct {
	*tchttp.BaseRequest

	// 开始时间 YYYY-MM-DD

	From *string `json:"From,omitempty" name:"From"`
	// 结束时间 YYYY-MM-DD

	To *string `json:"To,omitempty" name:"To"`
}

func (r *GetHttpQpsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHttpQpsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMailAlertConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetMailAlertConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMailAlertConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostForPool struct {
	// 节点名称

	// 盘符信息数组
}

type GetRecoveryExpansionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRecoveryExpansionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRecoveryExpansionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductStateInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 容灾系统产品状态信息查询返回值

		Data *ProductStateInfo `json:"Data,omitempty" name:"Data"`
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

type GetPoolByNameRsp struct {
	// cache_mode

	// create_time

	// dssid

	// ec_k

	// ec_m

	// failure_domain

	// 文件系统ID

	// founder

	// is_default

	// placement_name

	// policy

	// replicas

	// request_id

	// type
}

type GetHostDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// HostDetail

		HostDetail *GetHostDetailRsp `json:"HostDetail,omitempty" name:"HostDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHostDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostListRequest struct {
	*tchttp.BaseRequest

	// 筛选节点组件名

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
}

func (r *GetHostListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostRolesForDssidRequest struct {
	*tchttp.BaseRequest
}

func (r *GetHostRolesForDssidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostRolesForDssidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMailAlertConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMailAlertConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMailAlertConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HttpMethodMapping struct {
	// delete方法请求返回码统计

	DELETE *HttpCodeMapping `json:"DELETE,omitempty" name:"DELETE"`
	// get方法请求结果统计

	GET *HttpCodeMapping `json:"GET,omitempty" name:"GET"`
	// head方法请求结果统计

	HEAD *HttpCodeMapping `json:"HEAD,omitempty" name:"HEAD"`
	// options方法请求结果统计

	OPTIONS *HttpCodeMapping `json:"OPTIONS,omitempty" name:"OPTIONS"`
	// post方法请求结果统计

	POST *HttpCodeMapping `json:"POST,omitempty" name:"POST"`
	// put方法请求结果统计

	PUT *HttpCodeMapping `json:"PUT,omitempty" name:"PUT"`
}

type Disk struct {
	// 盘符名称

	// 盘符类型
}

type GetHostIfListResp struct {
	// 网卡名称

	// group

	// macaddr

	// master

	// mtu

	// ipv4addr

	// state

	// qlen

	// qdisc
}

type GetObjectStorageUserQuotaRsp struct {
	// buckets

	// capacity

	// objects
}

type DoAddComponentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoAddComponentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoAddComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoEditHostRequest struct {
	*tchttp.BaseRequest

	// 节点信息（IP、数据中心、机架、序列号）

	Host *Host `json:"Host,omitempty" name:"Host"`
}

func (r *DoEditHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoEditHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostDiskOverviewRequest struct {
	*tchttp.BaseRequest

	// 节点名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 磁盘名称

	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`
	// 开始时间

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束时间

	End *string `json:"End,omitempty" name:"End"`
}

func (r *GetHostDiskOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostDiskOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductFailureMigrateTaskDetailRspDetail struct {
	// StepName

	StepName *string `json:"StepName,omitempty" name:"StepName"`
	// Status

	Status *string `json:"Status,omitempty" name:"Status"`
	// Message

	Message *string `json:"Message,omitempty" name:"Message"`
	// StartAt

	StartAt *string `json:"StartAt,omitempty" name:"StartAt"`
	// EndAt

	EndAt *string `json:"EndAt,omitempty" name:"EndAt"`
}

type GetHostRolesForDssidRspHostComponentHostRoles struct {
	// component_name

	// dss_id
}

type SetClusterConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CspRequestId

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetClusterConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetClusterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetRgwLogLevelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetRgwLogLevelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetRgwLogLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoAddHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CspRequestId

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoAddHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoAddHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInspectionStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Status

		Status *GetInspectionStatusRsp `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetInspectionStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInspectionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostDisksRspUnplugged struct {
	// device

	// rotational

	// size

	// role

	// status

	// componentName

	// relatedComponents

	// lighting
}

type QueryProductFailureMigrateTaskDetailRsp struct {
	// 切换任务UUID

	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 进度

	Progress *int64 `json:"Progress,omitempty" name:"Progress"`
	// 执行信息/错误信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 详情

	Detail *QueryProductFailureMigrateTaskDetailRspDetail `json:"Detail,omitempty" name:"Detail"`
}

type GetOperationLogsRsp struct {
	// complete_time

	// context

	// data

	// id

	// operate_time

	// operator

	// state

	// type
}

type DoDeletePoolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoDeletePoolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoDeletePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoLocateDiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoLocateDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoLocateDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPoolHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPoolHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPoolHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBackupGlobalConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CspRequestId

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetBackupGlobalConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBackupGlobalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostListRspComponents struct {
	// component_name

	// maintenance_state

	// state

	// dss_id
}

type DoAddComponentRequest struct {
	*tchttp.BaseRequest

	// 节点列表

	HostNameList []*string `json:"HostNameList,omitempty" name:"HostNameList"`
	// 组件名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
}

func (r *DoAddComponentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoAddComponentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHttpCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求返回码统计

		Data *HttpCountObject `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHttpCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHttpCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationLogsRequest struct {
	*tchttp.BaseRequest

	// 开始时间戳

	Start *int64 `json:"Start,omitempty" name:"Start"`
	// 结束时间戳

	End *int64 `json:"End,omitempty" name:"End"`
}

func (r *GetOperationLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRequestDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRequestDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRequestDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRgwLogListRequest struct {
	*tchttp.BaseRequest

	// 节点IP

	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

func (r *GetRgwLogListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRgwLogListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAlertsRsp struct {
	// id

	// host

	// level

	// message

	// state

	// timestamp

	// timestamp
}

type GetAlertsRequest struct {
	*tchttp.BaseRequest

	// 告警状态

	State *string `json:"State,omitempty" name:"State"`
	// 告警级别

	Level *string `json:"Level,omitempty" name:"Level"`
	// 开始时间戳

	Start *int64 `json:"Start,omitempty" name:"Start"`
	// 结束时间戳

	End *int64 `json:"End,omitempty" name:"End"`
}

func (r *GetAlertsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostOverviewRequest struct {
	*tchttp.BaseRequest

	// 节点IP

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 开始时间戳

	Start *int64 `json:"Start,omitempty" name:"Start"`
	// 结束时间戳

	End *int64 `json:"End,omitempty" name:"End"`
}

func (r *GetHostOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetObjectStorageBucketQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetObjectStorageBucketQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetObjectStorageBucketQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetObjectStorageUserQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Quota

		Quota *GetObjectStorageUserQuotaRsp `json:"Quota,omitempty" name:"Quota"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetObjectStorageUserQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetObjectStorageUserQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRequestByIdRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
}

func (r *GetRequestByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRequestByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Hosts

		Hosts []*GetHostListRsp `json:"Hosts,omitempty" name:"Hosts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHostListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostDetailRequest struct {
	*tchttp.BaseRequest

	// 节点IP

	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

func (r *GetHostDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductFailureMigrateRequest struct {
	*tchttp.BaseRequest

	// 产品集群当前故障AZ

	CurrentAz *string `json:"CurrentAz,omitempty" name:"CurrentAz"`
	// 产品集群当前故障Region

	CurrentRegion *string `json:"CurrentRegion,omitempty" name:"CurrentRegion"`
	// 产品集群主节点目标迁移AZ

	TargetAz *string `json:"TargetAz,omitempty" name:"TargetAz"`
	// 产品集群主节点目标迁移Region

	TargetRegion *string `json:"TargetRegion,omitempty" name:"TargetRegion"`
	// 故障场景

	FailureScenario *string `json:"FailureScenario,omitempty" name:"FailureScenario"`
	// 操作类型，两种类型：1、切走；2、切回；

	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
	// 产品，对于像CLB拥有子产品的类型使用子产品名

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
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

type DoEditHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoEditHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoEditHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoReleaseDiskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoReleaseDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoReleaseDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBucketMonitoringResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBucketMonitoringResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBucketMonitoringResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDefaultPlacementRequest struct {
	*tchttp.BaseRequest

	// 存储池名称

	PlacementName *string `json:"PlacementName,omitempty" name:"PlacementName"`
}

func (r *SetDefaultPlacementRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDefaultPlacementRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInspectionStatusRsp struct {
	// state

	// requestId
}

type GetBackupGlobalConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetBackupGlobalConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBackupGlobalConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPoolLatestOverviewRequest struct {
	*tchttp.BaseRequest

	// 存储池 dssid 列表

	Dssids []*string `json:"Dssids,omitempty" name:"Dssids"`
}

func (r *GetPoolLatestOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPoolLatestOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOSUsageOverviewItem struct {
	// -

	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
	// -

	UserIdCount *string `json:"UserIdCount,omitempty" name:"UserIdCount"`
	// -

	BucketNameCount *string `json:"BucketNameCount,omitempty" name:"BucketNameCount"`
	// -

	InBytesSum *string `json:"InBytesSum,omitempty" name:"InBytesSum"`
	// -

	OutBytesSum *string `json:"OutBytesSum,omitempty" name:"OutBytesSum"`
	// -

	ReadCountSum *string `json:"ReadCountSum,omitempty" name:"ReadCountSum"`
	// -

	WriteCountSum *string `json:"WriteCountSum,omitempty" name:"WriteCountSum"`
	// -

	SizeByteSum *string `json:"SizeByteSum,omitempty" name:"SizeByteSum"`
	// -

	NumObjectsSum *string `json:"NumObjectsSum,omitempty" name:"NumObjectsSum"`
}

type QueryPrometheusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryPrometheusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPrometheusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PoolOverviewDetail struct {
	// 对应存储池id

	// 已使用容量，单位Byte

	// 总容量，单位Byte

	// 读iops

	// 写iops

	// 读带宽

	// 写带宽

	// 请求时间戳
}

type DoStartComponentRequest struct {
	*tchttp.BaseRequest

	// 节点数组

	HostNameList []*string `json:"HostNameList,omitempty" name:"HostNameList"`
	// 组件名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
}

func (r *DoStartComponentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoStartComponentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFlowControlConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 频控设置

		Config *FlowControlConfig `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFlowControlConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFlowControlConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFlowControlConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetFlowControlConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFlowControlConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductFailureMigrateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CreateProductFailureMigrateRsp

		Data []*CreateProductFailureMigrateRsp `json:"Data,omitempty" name:"Data"`
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

type GetBucketFlowControlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否启用频控

		Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
		// 写带宽阈值, 单位MB/s

		BrandI *int64 `json:"BrandI,omitempty" name:"BrandI"`
		// 读带宽阈值, 单位MB/s

		BrandO *int64 `json:"BrandO,omitempty" name:"BrandO"`
		// post iops阈值

		IopsPost *int64 `json:"IopsPost,omitempty" name:"IopsPost"`
		// put iops阈值

		IopsPut *int64 `json:"IopsPut,omitempty" name:"IopsPut"`
		// delete iops阈值

		IopsDelete *int64 `json:"IopsDelete,omitempty" name:"IopsDelete"`
		// get iops阈值

		IopsGet *int64 `json:"IopsGet,omitempty" name:"IopsGet"`
		// head iops阈值

		IopsHead *int64 `json:"IopsHead,omitempty" name:"IopsHead"`
		// options iops阈值

		IopsOptions *int64 `json:"IopsOptions,omitempty" name:"IopsOptions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBucketFlowControlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBucketFlowControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadInspectionRsp struct {
	// last_modified_time

	// start_time

	// status

	// result
}

type DoMigrateComponentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoMigrateComponentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoMigrateComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostCpuMemoryDataRequest struct {
	*tchttp.BaseRequest

	// 节点名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

func (r *GetHostCpuMemoryDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostCpuMemoryDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostRolesForDssidRspHostComponent struct {
	// HostRoles

	HostRoles []*GetHostRolesForDssidRspHostComponentHostRoles `json:"HostRoles,omitempty" name:"HostRoles"`
}

type HostsConfig struct {
	// 节点数组

	Hosts []*Host `json:"Hosts,omitempty" name:"Hosts"`
	// root账户密码

	RootPassword *string `json:"RootPassword,omitempty" name:"RootPassword"`
	// 端口号

	SshPort *string `json:"SshPort,omitempty" name:"SshPort"`
}

type DownloadInspectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Result

		Result *DownloadInspectionRsp `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadInspectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadInspectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHttpCountRequest struct {
	*tchttp.BaseRequest

	// 开始时间 YYYY-MM-DD

	From *string `json:"From,omitempty" name:"From"`
	// 结束时间 YYYY-MM-DD

	To *string `json:"To,omitempty" name:"To"`
}

func (r *GetHttpCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHttpCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePoolConfig struct {
	// 存储池名称

	// 冗余策略

	// 安全数据级别

	// 存储池 LocalStore

	// 是否预分裂

	// 分配模式

	// 副本数

	// 纠删码 K

	// 纠删码 M

	// 节点信息数组
}

type GetClusterConfigRsp struct {
	// publicNetworkSegment

	// clusterNetworkSegment

	// billCreateResourceUri

	// billUri

	// camAccountUri

	// camGetPolicyUri

	// camSetPolicyUri

	// camUri

	// clusterName

	// ntpServer

	// region

	// regionName

	// tagGetUri

	// tagSetUri

	// cosBaradNwsUrl

	// mgmtStatisticsPushgatewayUrl

	// zone

	// cryptTencentKmsRegion

	// cryptTencentKmsEndpoint

	// cryptTencentStsEndpoint

	// cryptTencentOcloudKmsEndpoint
}

type GetRequestDetailRequest struct {
	*tchttp.BaseRequest

	// request id

	CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
	// 是否输出任务的详细后台执行信息

	NeedStdio *bool `json:"NeedStdio,omitempty" name:"NeedStdio"`
	// 子任务id数组

	NeedStdioTaskIdList []*int64 `json:"NeedStdioTaskIdList,omitempty" name:"NeedStdioTaskIdList"`
}

func (r *GetRequestDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRequestDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSnmpConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// GetSnmpConfigResp

		Config *GetSnmpConfigResp `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSnmpConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSnmpConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoExpandPoolRequest struct {
	*tchttp.BaseRequest

	// 存储池名称

	PlacementName *string `json:"PlacementName,omitempty" name:"PlacementName"`
	// 扩容存储池所需的节点及磁盘信息

	HostsForPool []*HostForPool `json:"HostsForPool,omitempty" name:"HostsForPool"`
}

func (r *DoExpandPoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoExpandPoolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadInspectionRspResult struct {
	// result

	// sub_request
}

type ClusterConfig struct {
	// 集群名称

	// ntp地址

	// 地域名称

	// 地域Id

	// 区域名称

	// camUri

	// camAccountUri

	// camGetPolicyUri

	// camSetPolicyUri

	// billUri

	// billUri

	// 公共网段

	// 集群网段

	// tagGetUri

	// tagSetUri

	// cosBaradNwsUrl

	// mgmtStatisticsPushgatewayUrl

	// cryptTencentKmsEndpoint

	// cryptTencentStsEndpoint

	// cryptTencentOcloudKmsEndpoint
}

type DoLocateDiskRequest struct {
	*tchttp.BaseRequest

	// 节点IP

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 硬盘名称

	Device *string `json:"Device,omitempty" name:"Device"`
	// 点灯状态，'on' 表示点灯，'off' 表示关灯

	State *string `json:"State,omitempty" name:"State"`
}

func (r *DoLocateDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoLocateDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoRestartComponentRequest struct {
	*tchttp.BaseRequest

	// 节点数组

	HostNameList []*string `json:"HostNameList,omitempty" name:"HostNameList"`
	// 组件名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
}

func (r *DoRestartComponentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoRestartComponentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGlobalInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGlobalInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGlobalInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadInspectionRequest struct {
	*tchttp.BaseRequest
}

func (r *DownloadInspectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadInspectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostDisksRspDiskSmart struct {
	// status

	// spare_pct

	// retcode

	// stat_error

	// life_pct

	// online_hours

	// write_rounds.
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群配置项

	ClusterConfiguration *CreateClusterConfig `json:"ClusterConfiguration,omitempty" name:"ClusterConfiguration"`
}

func (r *CreateClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnsureClusterNotInstalledRequest struct {
	*tchttp.BaseRequest
}

func (r *EnsureClusterNotInstalledRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnsureClusterNotInstalledRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBkcmdbMonitorHostRequest struct {
	*tchttp.BaseRequest
}

func (r *GetBkcmdbMonitorHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBkcmdbMonitorHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetFlowControlConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetFlowControlConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetFlowControlConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetObjectStorageBucketQuotaRequest struct {
	*tchttp.BaseRequest

	// 存储桶名称

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 是否启用配额

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 对象数量配额, -1表示不限制

	ObjectQuota *int64 `json:"ObjectQuota,omitempty" name:"ObjectQuota"`
	// 容量配额, 单位byte, -1表示不限制

	CapacityQuota *int64 `json:"CapacityQuota,omitempty" name:"CapacityQuota"`
}

func (r *SetObjectStorageBucketQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetObjectStorageBucketQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoStartComponentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoStartComponentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoStartComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBkcmdbStorageHostRequest struct {
	*tchttp.BaseRequest
}

func (r *GetBkcmdbStorageHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBkcmdbStorageHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOSUsageOverviewGroupByUserIdRequest struct {
	*tchttp.BaseRequest

	// 时间戳, 秒

	Start *int64 `json:"Start,omitempty" name:"Start"`
	// 时间戳, 秒

	End *int64 `json:"End,omitempty" name:"End"`
	// 分页长度

	Length *int64 `json:"Length,omitempty" name:"Length"`
	// 分页起始值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// UserId

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// OrderBy字段, 可选: UserIdCount, BucketNameCount, InBytesSum, OutBytesSum, ReadCountSum, WriteCountSum, SizeByteSum, NumObjectsSum

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序类型, 可选: ASC, DESC

	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`
}

func (r *GetOSUsageOverviewGroupByUserIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOSUsageOverviewGroupByUserIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductStateInfo struct {
	// 产品名称标识

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集群列表

	ClusterList []*ClusterMemberNode `json:"ClusterList,omitempty" name:"ClusterList"`
	// 集群信息更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type DoReInitializeDiskRequest struct {
	*tchttp.BaseRequest

	// 节点IP

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 组件名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 硬盘名称

	Device *string `json:"Device,omitempty" name:"Device"`
}

func (r *DoReInitializeDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoReInitializeDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBkcmdbAccessHostRequest struct {
	*tchttp.BaseRequest
}

func (r *GetBkcmdbAccessHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBkcmdbAccessHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPoolByNameRequest struct {
	*tchttp.BaseRequest

	// 存储池名称

	PlacementName *string `json:"PlacementName,omitempty" name:"PlacementName"`
}

func (r *GetPoolByNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPoolByNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDnsDomainSettingsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetDnsDomainSettingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDnsDomainSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnsureClusterNotInstalledResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true为未安装，false未安装

		NotInstall *bool `json:"NotInstall,omitempty" name:"NotInstall"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnsureClusterNotInstalledResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnsureClusterNotInstalledResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostComponentDssidListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHostComponentDssidListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostComponentDssidListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRequestListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRequestListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRequestListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductHealthStateRequest struct {
	*tchttp.BaseRequest

	// 产品名

	Product *string `json:"Product,omitempty" name:"Product"`
	// 集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *QueryProductHealthStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductHealthStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPoolOverviewRequest struct {
	*tchttp.BaseRequest

	// 存储池 Dssid 列表

	Dssids []*string `json:"Dssids,omitempty" name:"Dssids"`
	// 开始时间戳

	From *int64 `json:"From,omitempty" name:"From"`
	// 结束时间戳

	To *int64 `json:"To,omitempty" name:"To"`
}

func (r *GetPoolOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPoolOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRequestByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRequestByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRequestByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOSUsageOverviewGroupByBucketResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// -

		Data []*GetOSUsageOverviewGroupByBucketItem `json:"Data,omitempty" name:"Data"`
		// -

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOSUsageOverviewGroupByBucketResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOSUsageOverviewGroupByBucketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPoolOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 所有存储池的概览信息

		Data []*PoolOverviewDetail `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPoolOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPoolOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRecoveryControlConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Config

		Config *GetRecoveryControlConfigRsp `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRecoveryControlConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRecoveryControlConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DoRetryPoolRequestResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新的request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DoRetryPoolRequestResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoRetryPoolRequestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Config

		Config *GetClusterConfigRsp `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostRolesForDssidRspHost struct {
	// host_name
}

type HttpCountObject struct {
	// 请求方法统计mapping

	// 全部请求总数
}

type GetRgwLogLevelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存储网关日志级别配置

		Config *GetRgwLogLevelResp `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRgwLogLevelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRgwLogLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRgwLogListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRgwLogListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRgwLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetRecoveryControlConfigRequest struct {
	*tchttp.BaseRequest

	// 恢复控制设置

	Config *RecoveryControlConfig `json:"Config,omitempty" name:"Config"`
}

func (r *SetRecoveryControlConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetRecoveryControlConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBkcmdbHostResp struct {
	// 数据中心

	// 节点IP

	// 机架

	// 序列号
}

type GetAutoRecoveryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 未启用的服务数量

		DisabledCount *int64 `json:"DisabledCount,omitempty" name:"DisabledCount"`
		// 已启用的服务数量

		EnabledCount *int64 `json:"EnabledCount,omitempty" name:"EnabledCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAutoRecoveryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAutoRecoveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBucketMonitoringRequest struct {
	*tchttp.BaseRequest

	// 储存桶名称

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 开始时间

	From *string `json:"From,omitempty" name:"From"`
	// 结束时间

	To *string `json:"To,omitempty" name:"To"`
	// 获取的字段, 默认为*, 代表获取所有字段

	Fields *string `json:"Fields,omitempty" name:"Fields"`
}

func (r *GetBucketMonitoringRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBucketMonitoringRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// OperationLogs

		OperationLogs []*GetOperationLogsRsp `json:"OperationLogs,omitempty" name:"OperationLogs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOperationLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOSUsageOverviewGroupByBucketRequest struct {
	*tchttp.BaseRequest

	// 时间戳, 秒

	Start *int64 `json:"Start,omitempty" name:"Start"`
	// 时间戳, 秒

	End *int64 `json:"End,omitempty" name:"End"`
	// 分页长度

	Length *int64 `json:"Length,omitempty" name:"Length"`
	// 分页起始值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// BucketName

	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
	// OrderBy字段, 可选: UserIdCount, BucketNameCount, InBytesSum, OutBytesSum, ReadCountSum, WriteCountSum, SizeByteSum, NumObjectsSum

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序类型, 可选: ASC, DESC

	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`
	// UserId筛选

	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *GetOSUsageOverviewGroupByBucketRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOSUsageOverviewGroupByBucketRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRequestInfoByIdRequest struct {
	*tchttp.BaseRequest

	// 任务 ID

	CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
}

func (r *GetRequestInfoByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRequestInfoByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostRolesForDssidRsp struct {
	// Hosts

	Hosts *GetHostRolesForDssidRspHost `json:"Hosts,omitempty" name:"Hosts"`
	// host_components
}

type QpsData struct {
	// qps值
}

type SetObjectStorageUserQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetObjectStorageUserQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetObjectStorageUserQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetRecoveryControlConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CspRequestId

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetRecoveryControlConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetRecoveryControlConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadHostDisksDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadHostDisksDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadHostDisksDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductHealthStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// response

		Data []*QueryProductHealthStateRsp `json:"Data,omitempty" name:"Data"`
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

type QueryProductStateInfoRequest struct {
	*tchttp.BaseRequest

	// 产品名

	Product *string `json:"Product,omitempty" name:"Product"`
	// 特殊参数

	Params *string `json:"Params,omitempty" name:"Params"`
}

func (r *QueryProductStateInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductStateInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterMemberNode struct {
	// 内网IP

	PrivateIP *string `json:"PrivateIP,omitempty" name:"PrivateIP"`
	// 公网IP

	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
	// 机器名称

	Host *string `json:"Host,omitempty" name:"Host"`
	// 机架名称

	Rack *string `json:"Rack,omitempty" name:"Rack"`
	// 设备所在可用区名称

	Az *string `json:"Az,omitempty" name:"Az"`
	// 设备所在地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 设备状态：alive, dead

	Status *string `json:"Status,omitempty" name:"Status"`
}

type PoolDetail struct {
	// 分配模式

	// 创建时间

	// 存储池id

	// 纠删码 k

	// 纠删码 m

	// 安全数据级别

	// 创建者

	// 是否为默认存储池，1代表是，0代表否

	// 存储池名称

	// 冗余策略

	// 副本数

	// 正在创建的存储池的requestId

	// 存储池状态，COMPLETED为存储池已安装完成

	// 存储池类型，os为对象存储

	// 当前存储池操作，例如删除、创建

	// 扩容进度

	// 恢复进度

	// 当前告警数量

	// 文件系统ID
}

type DoStartHostRequest struct {
	*tchttp.BaseRequest

	// 节点名

	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

func (r *DoStartHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DoStartHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAlertCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 未修复的重要告警数量

		WarningNum *int64 `json:"WarningNum,omitempty" name:"WarningNum"`
		// 未修复的紧急告警数量

		CriticalNum *int64 `json:"CriticalNum,omitempty" name:"CriticalNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAlertCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlertCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostCpuMemoryDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHostCpuMemoryDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostCpuMemoryDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryInitClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// request id

		CspRequestId *int64 `json:"CspRequestId,omitempty" name:"CspRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryInitClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryInitClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
