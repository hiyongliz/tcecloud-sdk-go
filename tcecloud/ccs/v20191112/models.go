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

package v20191112

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type GetDailyStatByAppIdRequest struct {
	*tchttp.BaseRequest

	// UserAppId

	UserAppId *uint64 `json:"UserAppId,omitempty" name:"UserAppId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 截止时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *GetDailyStatByAppIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDailyStatByAppIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetK8sVersionDistributionRequest struct {
	*tchttp.BaseRequest
}

func (r *GetK8sVersionDistributionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetK8sVersionDistributionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTopKUserDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopKUserDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopKUserDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTopKGrowUserDataRequest struct {
	*tchttp.BaseRequest

	// TopK

	TopK *int64 `json:"TopK,omitempty" name:"TopK"`
}

func (r *GetTopKGrowUserDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopKGrowUserDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsageStat struct {
	// Date

	Date *string `json:"Date,omitempty" name:"Date"`
	// AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// Clusters

	Clusters *int64 `json:"Clusters,omitempty" name:"Clusters"`
	// Nodes

	Nodes *int64 `json:"Nodes,omitempty" name:"Nodes"`
	// Pods

	Pods *int64 `json:"Pods,omitempty" name:"Pods"`
	// Mems

	Mems *int64 `json:"Mems,omitempty" name:"Mems"`
	// Cores

	Cores *int64 `json:"Cores,omitempty" name:"Cores"`
	// CvmCores

	CvmCores *int64 `json:"CvmCores,omitempty" name:"CvmCores"`
	// CloudProviderCores

	CloudProviderCores *int64 `json:"CloudProviderCores,omitempty" name:"CloudProviderCores"`
	// BlackstoneCores

	BlackstoneCores *int64 `json:"BlackstoneCores,omitempty" name:"BlackstoneCores"`
}

type DescribeClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群总个数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群信息列表。

		Clusters []*string `json:"Clusters,omitempty" name:"Clusters"`
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

type DescribeClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID。

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 用户appId。

	UserAppId *uint64 `json:"UserAppId,omitempty" name:"UserAppId"`
	// 集群状态。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 偏移量。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 窗口大小。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// UserAppId

	UserAppId *string `json:"UserAppId,omitempty" name:"UserAppId"`
}

func (r *DescribeInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetQuotaRequest struct {
	*tchttp.BaseRequest

	// UserUin

	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`
}

func (r *GetQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterQuotaRequest struct {
	*tchttp.BaseRequest

	// 单集群最大节点数量

	MaxNodeNum *uint64 `json:"MaxNodeNum,omitempty" name:"MaxNodeNum"`
	// 集群最大个数

	MaxClusterNum *uint64 `json:"MaxClusterNum,omitempty" name:"MaxClusterNum"`
	// UserAppId

	UserAppId *uint64 `json:"UserAppId,omitempty" name:"UserAppId"`
}

func (r *ModifyClusterQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Header struct {
	// http请求headers的key

	Key *string `json:"Key,omitempty" name:"Key"`
	// http请求headers的value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type GetClusterQuotaRequest struct {
	*tchttp.BaseRequest

	// UserAppId

	UserAppId *uint64 `json:"UserAppId,omitempty" name:"UserAppId"`
}

func (r *GetClusterQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDailyStatByAppIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// Data

		Data []*UsageStat `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDailyStatByAppIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDailyStatByAppIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTcrPersonalQuotaRespData struct {
	// TcrLimit

	LimitInfo []*TcrLimit `json:"LimitInfo,omitempty" name:"LimitInfo"`
}

type DescribeInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// TotalNum

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// NodeList

		NodeList []*VmInstance `json:"NodeList,omitempty" name:"NodeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TkeApiResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TkeApiResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TkeApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTopKGrowUserDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopKGrowUserDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopKGrowUserDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyQuotaRequest struct {
	*tchttp.BaseRequest

	// 额度

	Quota *uint64 `json:"Quota,omitempty" name:"Quota"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// UserUin

	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`
}

func (r *ModifyQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TcrLimit struct {
	// Username

	Username *string `json:"Username,omitempty" name:"Username"`
	// Type

	Type *string `json:"Type,omitempty" name:"Type"`
	// Value

	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type VmInstance struct {
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// ClusterInstanceId

	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" name:"ClusterInstanceId"`
	// UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// InstanceId

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// LanIp

	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`
	// Status

	Status *string `json:"Status,omitempty" name:"Status"`
	// IsReady

	IsReady *int64 `json:"IsReady,omitempty" name:"IsReady"`
	// NodeCondition

	NodeCondition *string `json:"NodeCondition,omitempty" name:"NodeCondition"`
	// Cpu

	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
	// Mem

	Mem *uint64 `json:"Mem,omitempty" name:"Mem"`
	// CvmPayMode

	CvmPayMode *uint64 `json:"CvmPayMode,omitempty" name:"CvmPayMode"`
	// NodeRole

	NodeRole *uint64 `json:"NodeRole,omitempty" name:"NodeRole"`
	// Version

	Version *uint64 `json:"Version,omitempty" name:"Version"`
	// Os

	Os *string `json:"Os,omitempty" name:"Os"`
	// Context

	Context *string `json:"Context,omitempty" name:"Context"`
	// Note

	Note *string `json:"Note,omitempty" name:"Note"`
	// MountTarget

	MountTarget *string `json:"MountTarget,omitempty" name:"MountTarget"`
	// DockerGraphPath

	DockerGraphPath *string `json:"DockerGraphPath,omitempty" name:"DockerGraphPath"`
	// DrainStatus

	DrainStatus *string `json:"DrainStatus,omitempty" name:"DrainStatus"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// UpdatedAt

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// K8sVersion

	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`
}

type GetQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// Data

		Data *GetTcrPersonalQuotaRespData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStatisticsDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStatisticsDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStatisticsDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStatisticsDataRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *GetStatisticsDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetStatisticsDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTopKReturnUserDataRequest struct {
	*tchttp.BaseRequest

	// TopK

	TopK *int64 `json:"TopK,omitempty" name:"TopK"`
}

func (r *GetTopKReturnUserDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopKReturnUserDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *Quota `json:"Data,omitempty" name:"Data"`
		// code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Quota struct {
	// AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// MaxClusterNum

	MaxClusterNum *uint64 `json:"MaxClusterNum,omitempty" name:"MaxClusterNum"`
	// MaxNodeNum

	MaxNodeNum *uint64 `json:"MaxNodeNum,omitempty" name:"MaxNodeNum"`
	// UpdateTime

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type GetK8sVersionDistributionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Code

		Code *int64 `json:"Code,omitempty" name:"Code"`
		// Message

		Message *string `json:"Message,omitempty" name:"Message"`
		// Data

		Data []*UsageStat `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetK8sVersionDistributionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetK8sVersionDistributionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTopKUserDataRequest struct {
	*tchttp.BaseRequest

	// OrderField

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *GetTopKUserDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopKUserDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTopKReturnUserDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopKReturnUserDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopKReturnUserDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TkeApiRequest struct {
	*tchttp.BaseRequest

	// 请求消息头

	Headers []*Header `json:"Headers,omitempty" name:"Headers"`
	// 请求方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// 请求path

	InterfacePath *string `json:"InterfacePath,omitempty" name:"InterfacePath"`
}

func (r *TkeApiRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TkeApiRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Headers struct{}
