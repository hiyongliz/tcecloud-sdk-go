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
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-09-10"

var _ = tchttp.POST

type Client struct {
	common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
	cpf := profile.NewClientProfile()
	client = &Client{}
	client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
	return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewDescribeApiServerMetricsRequest() (request *DescribeApiServerMetricsRequest) {
	request = &DescribeApiServerMetricsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeApiServerMetrics")
	return
}

func NewDescribeApiServerMetricsResponse() (response *DescribeApiServerMetricsResponse) {
	response = &DescribeApiServerMetricsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeApiServerMetrics
func (c *Client) DescribeApiServerMetrics(request *DescribeApiServerMetricsRequest) (response *DescribeApiServerMetricsResponse, err error) {
	if request == nil {
		request = NewDescribeApiServerMetricsRequest()
	}
	response = NewDescribeApiServerMetricsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHPARequest() (request *ModifyHPARequest) {
	request = &ModifyHPARequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyHPA")
	return
}

func NewModifyHPAResponse() (response *ModifyHPAResponse) {
	response = &ModifyHPAResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改HPA
func (c *Client) ModifyHPA(request *ModifyHPARequest) (response *ModifyHPAResponse, err error) {
	if request == nil {
		request = NewModifyHPARequest()
	}
	response = NewModifyHPAResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteStatefulSetRequest() (request *DeleteStatefulSetRequest) {
	request = &DeleteStatefulSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteStatefulSet")
	return
}

func NewDeleteStatefulSetResponse() (response *DeleteStatefulSetResponse) {
	response = &DeleteStatefulSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除有状态应用
func (c *Client) DeleteStatefulSet(request *DeleteStatefulSetRequest) (response *DeleteStatefulSetResponse, err error) {
	if request == nil {
		request = NewDeleteStatefulSetRequest()
	}
	response = NewDeleteStatefulSetResponse()
	err = c.Send(request, response)
	return
}

func NewInitMachineRequest() (request *InitMachineRequest) {
	request = &InitMachineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "InitMachine")
	return
}

func NewInitMachineResponse() (response *InitMachineResponse) {
	response = &InitMachineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 初始化服务器
func (c *Client) InitMachine(request *InitMachineRequest) (response *InitMachineResponse, err error) {
	if request == nil {
		request = NewInitMachineRequest()
	}
	response = NewInitMachineResponse()
	err = c.Send(request, response)
	return
}

func NewStartImageBuildingTaskRequest() (request *StartImageBuildingTaskRequest) {
	request = &StartImageBuildingTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "StartImageBuildingTask")
	return
}

func NewStartImageBuildingTaskResponse() (response *StartImageBuildingTaskResponse) {
	response = &StartImageBuildingTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动镜像构建任务
func (c *Client) StartImageBuildingTask(request *StartImageBuildingTaskRequest) (response *StartImageBuildingTaskResponse, err error) {
	if request == nil {
		request = NewStartImageBuildingTaskRequest()
	}
	response = NewStartImageBuildingTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeregisterTargetsRequest() (request *DeregisterTargetsRequest) {
	request = &DeregisterTargetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeregisterTargets")
	return
}

func NewDeregisterTargetsResponse() (response *DeregisterTargetsResponse) {
	response = &DeregisterTargetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从负载均衡监听器上解绑后端服务
func (c *Client) DeregisterTargets(request *DeregisterTargetsRequest) (response *DeregisterTargetsResponse, err error) {
	if request == nil {
		request = NewDeregisterTargetsRequest()
	}
	response = NewDeregisterTargetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLoadBalancersRequest() (request *DescribeLoadBalancersRequest) {
	request = &DescribeLoadBalancersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeLoadBalancers")
	return
}

func NewDescribeLoadBalancersResponse() (response *DescribeLoadBalancersResponse) {
	response = &DescribeLoadBalancersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询负载均衡实例列表
func (c *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
	if request == nil {
		request = NewDescribeLoadBalancersRequest()
	}
	response = NewDescribeLoadBalancersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceInstanceStatusOverviewRequest() (request *DescribeServiceInstanceStatusOverviewRequest) {
	request = &DescribeServiceInstanceStatusOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeServiceInstanceStatusOverview")
	return
}

func NewDescribeServiceInstanceStatusOverviewResponse() (response *DescribeServiceInstanceStatusOverviewResponse) {
	response = &DescribeServiceInstanceStatusOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 支持服务状态概览
func (c *Client) DescribeServiceInstanceStatusOverview(request *DescribeServiceInstanceStatusOverviewRequest) (response *DescribeServiceInstanceStatusOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeServiceInstanceStatusOverviewRequest()
	}
	response = NewDescribeServiceInstanceStatusOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRAIDRequest() (request *DeleteRAIDRequest) {
	request = &DeleteRAIDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteRAID")
	return
}

func NewDeleteRAIDResponse() (response *DeleteRAIDResponse) {
	response = &DeleteRAIDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除RAID
func (c *Client) DeleteRAID(request *DeleteRAIDRequest) (response *DeleteRAIDResponse, err error) {
	if request == nil {
		request = NewDeleteRAIDRequest()
	}
	response = NewDeleteRAIDResponse()
	err = c.Send(request, response)
	return
}

func NewListStorageClassesRequest() (request *ListStorageClassesRequest) {
	request = &ListStorageClassesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListStorageClasses")
	return
}

func NewListStorageClassesResponse() (response *ListStorageClassesResponse) {
	response = &ListStorageClassesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出StorageClass
func (c *Client) ListStorageClasses(request *ListStorageClassesRequest) (response *ListStorageClassesResponse, err error) {
	if request == nil {
		request = NewListStorageClassesRequest()
	}
	response = NewListStorageClassesResponse()
	err = c.Send(request, response)
	return
}

func NewSetVirtualMachineStatusRequest() (request *SetVirtualMachineStatusRequest) {
	request = &SetVirtualMachineStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SetVirtualMachineStatus")
	return
}

func NewSetVirtualMachineStatusResponse() (response *SetVirtualMachineStatusResponse) {
	response = &SetVirtualMachineStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置虚拟机状态
func (c *Client) SetVirtualMachineStatus(request *SetVirtualMachineStatusRequest) (response *SetVirtualMachineStatusResponse, err error) {
	if request == nil {
		request = NewSetVirtualMachineStatusRequest()
	}
	response = NewSetVirtualMachineStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDriveDetailByNodeMetricRequest() (request *DriveDetailByNodeMetricRequest) {
	request = &DriveDetailByNodeMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DriveDetailByNodeMetric")
	return
}

func NewDriveDetailByNodeMetricResponse() (response *DriveDetailByNodeMetricResponse) {
	response = &DriveDetailByNodeMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 节点下钻接口
func (c *Client) DriveDetailByNodeMetric(request *DriveDetailByNodeMetricRequest) (response *DriveDetailByNodeMetricResponse, err error) {
	if request == nil {
		request = NewDriveDetailByNodeMetricRequest()
	}
	response = NewDriveDetailByNodeMetricResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
	request = &DeleteClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteCluster")
	return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
	response = &DeleteClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除集群
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
	if request == nil {
		request = NewDeleteClusterRequest()
	}
	response = NewDeleteClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSilencesRequest() (request *DeleteSilencesRequest) {
	request = &DeleteSilencesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteSilences")
	return
}

func NewDeleteSilencesResponse() (response *DeleteSilencesResponse) {
	response = &DeleteSilencesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除静默规则
func (c *Client) DeleteSilences(request *DeleteSilencesRequest) (response *DeleteSilencesResponse, err error) {
	if request == nil {
		request = NewDeleteSilencesRequest()
	}
	response = NewDeleteSilencesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServerHarddisksRequest() (request *DescribeServerHarddisksRequest) {
	request = &DescribeServerHarddisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeServerHarddisks")
	return
}

func NewDescribeServerHarddisksResponse() (response *DescribeServerHarddisksResponse) {
	response = &DescribeServerHarddisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取物理机磁盘列表
func (c *Client) DescribeServerHarddisks(request *DescribeServerHarddisksRequest) (response *DescribeServerHarddisksResponse, err error) {
	if request == nil {
		request = NewDescribeServerHarddisksRequest()
	}
	response = NewDescribeServerHarddisksResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDashboardFolderRequest() (request *ModifyDashboardFolderRequest) {
	request = &ModifyDashboardFolderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyDashboardFolder")
	return
}

func NewModifyDashboardFolderResponse() (response *ModifyDashboardFolderResponse) {
	response = &ModifyDashboardFolderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改dashboard文件夹
func (c *Client) ModifyDashboardFolder(request *ModifyDashboardFolderRequest) (response *ModifyDashboardFolderResponse, err error) {
	if request == nil {
		request = NewModifyDashboardFolderRequest()
	}
	response = NewModifyDashboardFolderResponse()
	err = c.Send(request, response)
	return
}

func NewGetNotificationRuleRequest() (request *GetNotificationRuleRequest) {
	request = &GetNotificationRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetNotificationRule")
	return
}

func NewGetNotificationRuleResponse() (response *GetNotificationRuleResponse) {
	response = &GetNotificationRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群通知订阅信息
func (c *Client) GetNotificationRule(request *GetNotificationRuleRequest) (response *GetNotificationRuleResponse, err error) {
	if request == nil {
		request = NewGetNotificationRuleRequest()
	}
	response = NewGetNotificationRuleResponse()
	err = c.Send(request, response)
	return
}

func NewListIngressesRequest() (request *ListIngressesRequest) {
	request = &ListIngressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListIngresses")
	return
}

func NewListIngressesResponse() (response *ListIngressesResponse) {
	response = &ListIngressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Ingress列表
func (c *Client) ListIngresses(request *ListIngressesRequest) (response *ListIngressesResponse, err error) {
	if request == nil {
		request = NewListIngressesRequest()
	}
	response = NewListIngressesResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceTypeEventRequest() (request *ListResourceTypeEventRequest) {
	request = &ListResourceTypeEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListResourceTypeEvent")
	return
}

func NewListResourceTypeEventResponse() (response *ListResourceTypeEventResponse) {
	response = &ListResourceTypeEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源对象元数据-对象类型事件列表
func (c *Client) ListResourceTypeEvent(request *ListResourceTypeEventRequest) (response *ListResourceTypeEventResponse, err error) {
	if request == nil {
		request = NewListResourceTypeEventRequest()
	}
	response = NewListResourceTypeEventResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopKResourcesRequest() (request *DescribeTopKResourcesRequest) {
	request = &DescribeTopKResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeTopKResources")
	return
}

func NewDescribeTopKResourcesResponse() (response *DescribeTopKResourcesResponse) {
	response = &DescribeTopKResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询TOP K资源使用量的资源, 如集群/节点/项目/应用
func (c *Client) DescribeTopKResources(request *DescribeTopKResourcesRequest) (response *DescribeTopKResourcesResponse, err error) {
	if request == nil {
		request = NewDescribeTopKResourcesRequest()
	}
	response = NewDescribeTopKResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLabelsRequest() (request *ModifyLabelsRequest) {
	request = &ModifyLabelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyLabels")
	return
}

func NewModifyLabelsResponse() (response *ModifyLabelsResponse) {
	response = &ModifyLabelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改物理服务器标签
func (c *Client) ModifyLabels(request *ModifyLabelsRequest) (response *ModifyLabelsResponse, err error) {
	if request == nil {
		request = NewModifyLabelsRequest()
	}
	response = NewModifyLabelsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTcsComponentPodRequest() (request *DescribeTcsComponentPodRequest) {
	request = &DescribeTcsComponentPodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeTcsComponentPod")
	return
}

func NewDescribeTcsComponentPodResponse() (response *DescribeTcsComponentPodResponse) {
	response = &DescribeTcsComponentPodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeTcsComponentPod
func (c *Client) DescribeTcsComponentPod(request *DescribeTcsComponentPodRequest) (response *DescribeTcsComponentPodResponse, err error) {
	if request == nil {
		request = NewDescribeTcsComponentPodRequest()
	}
	response = NewDescribeTcsComponentPodResponse()
	err = c.Send(request, response)
	return
}

func NewListClusterRequest() (request *ListClusterRequest) {
	request = &ListClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListCluster")
	return
}

func NewListClusterResponse() (response *ListClusterResponse) {
	response = &ListClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群列表信息
func (c *Client) ListCluster(request *ListClusterRequest) (response *ListClusterResponse, err error) {
	if request == nil {
		request = NewListClusterRequest()
	}
	response = NewListClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRelocateServerHistoryRequest() (request *DescribeRelocateServerHistoryRequest) {
	request = &DescribeRelocateServerHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeRelocateServerHistory")
	return
}

func NewDescribeRelocateServerHistoryResponse() (response *DescribeRelocateServerHistoryResponse) {
	response = &DescribeRelocateServerHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询历史搬迁接口
func (c *Client) DescribeRelocateServerHistory(request *DescribeRelocateServerHistoryRequest) (response *DescribeRelocateServerHistoryResponse, err error) {
	if request == nil {
		request = NewDescribeRelocateServerHistoryRequest()
	}
	response = NewDescribeRelocateServerHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDaemonSetRequest() (request *CreateDaemonSetRequest) {
	request = &CreateDaemonSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateDaemonSet")
	return
}

func NewCreateDaemonSetResponse() (response *CreateDaemonSetResponse) {
	response = &CreateDaemonSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建DaemonSet
func (c *Client) CreateDaemonSet(request *CreateDaemonSetRequest) (response *CreateDaemonSetResponse, err error) {
	if request == nil {
		request = NewCreateDaemonSetRequest()
	}
	response = NewCreateDaemonSetResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDeploymentRequest() (request *CreateDeploymentRequest) {
	request = &CreateDeploymentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateDeployment")
	return
}

func NewCreateDeploymentResponse() (response *CreateDeploymentResponse) {
	response = &CreateDeploymentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建Deployment
func (c *Client) CreateDeployment(request *CreateDeploymentRequest) (response *CreateDeploymentResponse, err error) {
	if request == nil {
		request = NewCreateDeploymentRequest()
	}
	response = NewCreateDeploymentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCronJobRequest() (request *DescribeCronJobRequest) {
	request = &DescribeCronJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeCronJob")
	return
}

func NewDescribeCronJobResponse() (response *DescribeCronJobResponse) {
	response = &DescribeCronJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取CronJob详情
func (c *Client) DescribeCronJob(request *DescribeCronJobRequest) (response *DescribeCronJobResponse, err error) {
	if request == nil {
		request = NewDescribeCronJobRequest()
	}
	response = NewDescribeCronJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFieldsEnumExRequest() (request *DescribeFieldsEnumExRequest) {
	request = &DescribeFieldsEnumExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeFieldsEnumEx")
	return
}

func NewDescribeFieldsEnumExResponse() (response *DescribeFieldsEnumExResponse) {
	response = &DescribeFieldsEnumExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取cmdb各字段枚举值
func (c *Client) DescribeFieldsEnumEx(request *DescribeFieldsEnumExRequest) (response *DescribeFieldsEnumExResponse, err error) {
	if request == nil {
		request = NewDescribeFieldsEnumExRequest()
	}
	response = NewDescribeFieldsEnumExResponse()
	err = c.Send(request, response)
	return
}

func NewListServicesRequest() (request *ListServicesRequest) {
	request = &ListServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListServices")
	return
}

func NewListServicesResponse() (response *ListServicesResponse) {
	response = &ListServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群服务列表
func (c *Client) ListServices(request *ListServicesRequest) (response *ListServicesResponse, err error) {
	if request == nil {
		request = NewListServicesRequest()
	}
	response = NewListServicesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateJobRequest() (request *CreateJobRequest) {
	request = &CreateJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateJob")
	return
}

func NewCreateJobResponse() (response *CreateJobResponse) {
	response = &CreateJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建Job
func (c *Client) CreateJob(request *CreateJobRequest) (response *CreateJobResponse, err error) {
	if request == nil {
		request = NewCreateJobRequest()
	}
	response = NewCreateJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLoadBalancersDetailRequest() (request *DescribeLoadBalancersDetailRequest) {
	request = &DescribeLoadBalancersDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeLoadBalancersDetail")
	return
}

func NewDescribeLoadBalancersDetailResponse() (response *DescribeLoadBalancersDetailResponse) {
	response = &DescribeLoadBalancersDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询负载均衡详细信息
func (c *Client) DescribeLoadBalancersDetail(request *DescribeLoadBalancersDetailRequest) (response *DescribeLoadBalancersDetailResponse, err error) {
	if request == nil {
		request = NewDescribeLoadBalancersDetailRequest()
	}
	response = NewDescribeLoadBalancersDetailResponse()
	err = c.Send(request, response)
	return
}

func NewListCronJobsRequest() (request *ListCronJobsRequest) {
	request = &ListCronJobsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListCronJobs")
	return
}

func NewListCronJobsResponse() (response *ListCronJobsResponse) {
	response = &ListCronJobsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出CronJob列表
func (c *Client) ListCronJobs(request *ListCronJobsRequest) (response *ListCronJobsResponse, err error) {
	if request == nil {
		request = NewListCronJobsRequest()
	}
	response = NewListCronJobsResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceTypesRequest() (request *ListResourceTypesRequest) {
	request = &ListResourceTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListResourceTypes")
	return
}

func NewListResourceTypesResponse() (response *ListResourceTypesResponse) {
	response = &ListResourceTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源对象列表
func (c *Client) ListResourceTypes(request *ListResourceTypesRequest) (response *ListResourceTypesResponse, err error) {
	if request == nil {
		request = NewListResourceTypesRequest()
	}
	response = NewListResourceTypesResponse()
	err = c.Send(request, response)
	return
}

func NewSearchSilencesRequest() (request *SearchSilencesRequest) {
	request = &SearchSilencesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SearchSilences")
	return
}

func NewSearchSilencesResponse() (response *SearchSilencesResponse) {
	response = &SearchSilencesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件查询屏蔽策略
func (c *Client) SearchSilences(request *SearchSilencesRequest) (response *SearchSilencesResponse, err error) {
	if request == nil {
		request = NewSearchSilencesRequest()
	}
	response = NewSearchSilencesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMonitorTemplateRequest() (request *DescribeMonitorTemplateRequest) {
	request = &DescribeMonitorTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeMonitorTemplate")
	return
}

func NewDescribeMonitorTemplateResponse() (response *DescribeMonitorTemplateResponse) {
	response = &DescribeMonitorTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取特定的自定义监控模板
func (c *Client) DescribeMonitorTemplate(request *DescribeMonitorTemplateRequest) (response *DescribeMonitorTemplateResponse, err error) {
	if request == nil {
		request = NewDescribeMonitorTemplateRequest()
	}
	response = NewDescribeMonitorTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewGetNotificationMsgRequest() (request *GetNotificationMsgRequest) {
	request = &GetNotificationMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetNotificationMsg")
	return
}

func NewGetNotificationMsgResponse() (response *GetNotificationMsgResponse) {
	response = &GetNotificationMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取通知消息
func (c *Client) GetNotificationMsg(request *GetNotificationMsgRequest) (response *GetNotificationMsgResponse, err error) {
	if request == nil {
		request = NewGetNotificationMsgRequest()
	}
	response = NewGetNotificationMsgResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMiddlewareOverviewWithAlertsRequest() (request *DescribeMiddlewareOverviewWithAlertsRequest) {
	request = &DescribeMiddlewareOverviewWithAlertsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeMiddlewareOverviewWithAlerts")
	return
}

func NewDescribeMiddlewareOverviewWithAlertsResponse() (response *DescribeMiddlewareOverviewWithAlertsResponse) {
	response = &DescribeMiddlewareOverviewWithAlertsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云哨获取中间件信息
func (c *Client) DescribeMiddlewareOverviewWithAlerts(request *DescribeMiddlewareOverviewWithAlertsRequest) (response *DescribeMiddlewareOverviewWithAlertsResponse, err error) {
	if request == nil {
		request = NewDescribeMiddlewareOverviewWithAlertsRequest()
	}
	response = NewDescribeMiddlewareOverviewWithAlertsResponse()
	err = c.Send(request, response)
	return
}

func NewGetClustersOverviewRequest() (request *GetClustersOverviewRequest) {
	request = &GetClustersOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetClustersOverview")
	return
}

func NewGetClustersOverviewResponse() (response *GetClustersOverviewResponse) {
	response = &GetClustersOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取监控概览页集群概览
func (c *Client) GetClustersOverview(request *GetClustersOverviewRequest) (response *GetClustersOverviewResponse, err error) {
	if request == nil {
		request = NewGetClustersOverviewRequest()
	}
	response = NewGetClustersOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewCreateStorageClassRequest() (request *CreateStorageClassRequest) {
	request = &CreateStorageClassRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateStorageClass")
	return
}

func NewCreateStorageClassResponse() (response *CreateStorageClassResponse) {
	response = &CreateStorageClassResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建StorageClass
func (c *Client) CreateStorageClass(request *CreateStorageClassRequest) (response *CreateStorageClassResponse, err error) {
	if request == nil {
		request = NewCreateStorageClassRequest()
	}
	response = NewCreateStorageClassResponse()
	err = c.Send(request, response)
	return
}

func NewExportLogsRequest() (request *ExportLogsRequest) {
	request = &ExportLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ExportLogs")
	return
}

func NewExportLogsResponse() (response *ExportLogsResponse) {
	response = &ExportLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志导出
func (c *Client) ExportLogs(request *ExportLogsRequest) (response *ExportLogsResponse, err error) {
	if request == nil {
		request = NewExportLogsRequest()
	}
	response = NewExportLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCustomScriptSetsRequest() (request *DeleteCustomScriptSetsRequest) {
	request = &DeleteCustomScriptSetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteCustomScriptSets")
	return
}

func NewDeleteCustomScriptSetsResponse() (response *DeleteCustomScriptSetsResponse) {
	response = &DeleteCustomScriptSetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除自定义脚本集
func (c *Client) DeleteCustomScriptSets(request *DeleteCustomScriptSetsRequest) (response *DeleteCustomScriptSetsResponse, err error) {
	if request == nil {
		request = NewDeleteCustomScriptSetsRequest()
	}
	response = NewDeleteCustomScriptSetsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomScriptSetRequest() (request *ModifyCustomScriptSetRequest) {
	request = &ModifyCustomScriptSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyCustomScriptSet")
	return
}

func NewModifyCustomScriptSetResponse() (response *ModifyCustomScriptSetResponse) {
	response = &ModifyCustomScriptSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑自定义脚本集信息
func (c *Client) ModifyCustomScriptSet(request *ModifyCustomScriptSetRequest) (response *ModifyCustomScriptSetResponse, err error) {
	if request == nil {
		request = NewModifyCustomScriptSetRequest()
	}
	response = NewModifyCustomScriptSetResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVirtualMachineAttributesRequest() (request *ModifyVirtualMachineAttributesRequest) {
	request = &ModifyVirtualMachineAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyVirtualMachineAttributes")
	return
}

func NewModifyVirtualMachineAttributesResponse() (response *ModifyVirtualMachineAttributesResponse) {
	response = &ModifyVirtualMachineAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改虚拟机描述
func (c *Client) ModifyVirtualMachineAttributes(request *ModifyVirtualMachineAttributesRequest) (response *ModifyVirtualMachineAttributesResponse, err error) {
	if request == nil {
		request = NewModifyVirtualMachineAttributesRequest()
	}
	response = NewModifyVirtualMachineAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewServerRecycleLanIPExRequest() (request *ServerRecycleLanIPExRequest) {
	request = &ServerRecycleLanIPExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ServerRecycleLanIPEx")
	return
}

func NewServerRecycleLanIPExResponse() (response *ServerRecycleLanIPExResponse) {
	response = &ServerRecycleLanIPExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器回收内网IP接口
func (c *Client) ServerRecycleLanIPEx(request *ServerRecycleLanIPExRequest) (response *ServerRecycleLanIPExResponse, err error) {
	if request == nil {
		request = NewServerRecycleLanIPExRequest()
	}
	response = NewServerRecycleLanIPExResponse()
	err = c.Send(request, response)
	return
}

func NewGetNodeLabelsRequest() (request *GetNodeLabelsRequest) {
	request = &GetNodeLabelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetNodeLabels")
	return
}

func NewGetNodeLabelsResponse() (response *GetNodeLabelsResponse) {
	response = &GetNodeLabelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取某个节点的标签列表
func (c *Client) GetNodeLabels(request *GetNodeLabelsRequest) (response *GetNodeLabelsResponse, err error) {
	if request == nil {
		request = NewGetNodeLabelsRequest()
	}
	response = NewGetNodeLabelsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyImageRequest() (request *ModifyImageRequest) {
	request = &ModifyImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyImage")
	return
}

func NewModifyImageResponse() (response *ModifyImageResponse) {
	response = &ModifyImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改镜像
func (c *Client) ModifyImage(request *ModifyImageRequest) (response *ModifyImageResponse, err error) {
	if request == nil {
		request = NewModifyImageRequest()
	}
	response = NewModifyImageResponse()
	err = c.Send(request, response)
	return
}

func NewAddMetaResourceTypeRequest() (request *AddMetaResourceTypeRequest) {
	request = &AddMetaResourceTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "AddMetaResourceType")
	return
}

func NewAddMetaResourceTypeResponse() (response *AddMetaResourceTypeResponse) {
	response = &AddMetaResourceTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建资源对象类型
func (c *Client) AddMetaResourceType(request *AddMetaResourceTypeRequest) (response *AddMetaResourceTypeResponse, err error) {
	if request == nil {
		request = NewAddMetaResourceTypeRequest()
	}
	response = NewAddMetaResourceTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNodesOverviewRequest() (request *DescribeNodesOverviewRequest) {
	request = &DescribeNodesOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeNodesOverview")
	return
}

func NewDescribeNodesOverviewResponse() (response *DescribeNodesOverviewResponse) {
	response = &DescribeNodesOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点网格概览
func (c *Client) DescribeNodesOverview(request *DescribeNodesOverviewRequest) (response *DescribeNodesOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeNodesOverviewRequest()
	}
	response = NewDescribeNodesOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewListConfigMapsRequest() (request *ListConfigMapsRequest) {
	request = &ListConfigMapsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListConfigMaps")
	return
}

func NewListConfigMapsResponse() (response *ListConfigMapsResponse) {
	response = &ListConfigMapsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询ConfigMap列表
func (c *Client) ListConfigMaps(request *ListConfigMapsRequest) (response *ListConfigMapsResponse, err error) {
	if request == nil {
		request = NewListConfigMapsRequest()
	}
	response = NewListConfigMapsResponse()
	err = c.Send(request, response)
	return
}

func NewListPersistentVolumeClaimsRequest() (request *ListPersistentVolumeClaimsRequest) {
	request = &ListPersistentVolumeClaimsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListPersistentVolumeClaims")
	return
}

func NewListPersistentVolumeClaimsResponse() (response *ListPersistentVolumeClaimsResponse) {
	response = &ListPersistentVolumeClaimsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取PVC列表
func (c *Client) ListPersistentVolumeClaims(request *ListPersistentVolumeClaimsRequest) (response *ListPersistentVolumeClaimsResponse, err error) {
	if request == nil {
		request = NewListPersistentVolumeClaimsRequest()
	}
	response = NewListPersistentVolumeClaimsResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateServerSpecialWanIPListRequest() (request *AllocateServerSpecialWanIPListRequest) {
	request = &AllocateServerSpecialWanIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "AllocateServerSpecialWanIPList")
	return
}

func NewAllocateServerSpecialWanIPListResponse() (response *AllocateServerSpecialWanIPListResponse) {
	response = &AllocateServerSpecialWanIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分配服务器特殊外网IP
func (c *Client) AllocateServerSpecialWanIPList(request *AllocateServerSpecialWanIPListRequest) (response *AllocateServerSpecialWanIPListResponse, err error) {
	if request == nil {
		request = NewAllocateServerSpecialWanIPListRequest()
	}
	response = NewAllocateServerSpecialWanIPListResponse()
	err = c.Send(request, response)
	return
}

func NewQueryOutbandStrategyRequest() (request *QueryOutbandStrategyRequest) {
	request = &QueryOutbandStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "QueryOutbandStrategy")
	return
}

func NewQueryOutbandStrategyResponse() (response *QueryOutbandStrategyResponse) {
	response = &QueryOutbandStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询自定义带外密码策略
func (c *Client) QueryOutbandStrategy(request *QueryOutbandStrategyRequest) (response *QueryOutbandStrategyResponse, err error) {
	if request == nil {
		request = NewQueryOutbandStrategyRequest()
	}
	response = NewQueryOutbandStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateServerVirtualIPRequest() (request *AllocateServerVirtualIPRequest) {
	request = &AllocateServerVirtualIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "AllocateServerVirtualIP")
	return
}

func NewAllocateServerVirtualIPResponse() (response *AllocateServerVirtualIPResponse) {
	response = &AllocateServerVirtualIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分配物理服务器虚拟内外网ip
func (c *Client) AllocateServerVirtualIP(request *AllocateServerVirtualIPRequest) (response *AllocateServerVirtualIPResponse, err error) {
	if request == nil {
		request = NewAllocateServerVirtualIPRequest()
	}
	response = NewAllocateServerVirtualIPResponse()
	err = c.Send(request, response)
	return
}

func NewListStorageTypesRequest() (request *ListStorageTypesRequest) {
	request = &ListStorageTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListStorageTypes")
	return
}

func NewListStorageTypesResponse() (response *ListStorageTypesResponse) {
	response = &ListStorageTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取创建当前可用于创建SC的存储类型, 例如CSP/Local
func (c *Client) ListStorageTypes(request *ListStorageTypesRequest) (response *ListStorageTypesResponse, err error) {
	if request == nil {
		request = NewListStorageTypesRequest()
	}
	response = NewListStorageTypesResponse()
	err = c.Send(request, response)
	return
}

func NewSearchNotificationsRequest() (request *SearchNotificationsRequest) {
	request = &SearchNotificationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SearchNotifications")
	return
}

func NewSearchNotificationsResponse() (response *SearchNotificationsResponse) {
	response = &SearchNotificationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取告警聚合接口
func (c *Client) SearchNotifications(request *SearchNotificationsRequest) (response *SearchNotificationsResponse, err error) {
	if request == nil {
		request = NewSearchNotificationsRequest()
	}
	response = NewSearchNotificationsResponse()
	err = c.Send(request, response)
	return
}

func NewAddMetaMetricRequest() (request *AddMetaMetricRequest) {
	request = &AddMetaMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "AddMetaMetric")
	return
}

func NewAddMetaMetricResponse() (response *AddMetaMetricResponse) {
	response = &AddMetaMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在资源对象类型下创建指标
func (c *Client) AddMetaMetric(request *AddMetaMetricRequest) (response *AddMetaMetricResponse, err error) {
	if request == nil {
		request = NewAddMetaMetricRequest()
	}
	response = NewAddMetaMetricResponse()
	err = c.Send(request, response)
	return
}

func NewListDevOpsScenesRequest() (request *ListDevOpsScenesRequest) {
	request = &ListDevOpsScenesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListDevOpsScenes")
	return
}

func NewListDevOpsScenesResponse() (response *ListDevOpsScenesResponse) {
	response = &ListDevOpsScenesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 自愈场景列表
func (c *Client) ListDevOpsScenes(request *ListDevOpsScenesRequest) (response *ListDevOpsScenesResponse, err error) {
	if request == nil {
		request = NewListDevOpsScenesRequest()
	}
	response = NewListDevOpsScenesResponse()
	err = c.Send(request, response)
	return
}

func NewGetAlertRequest() (request *GetAlertRequest) {
	request = &GetAlertRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetAlert")
	return
}

func NewGetAlertResponse() (response *GetAlertResponse) {
	response = &GetAlertResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取一个告警
func (c *Client) GetAlert(request *GetAlertRequest) (response *GetAlertResponse, err error) {
	if request == nil {
		request = NewGetAlertRequest()
	}
	response = NewGetAlertResponse()
	err = c.Send(request, response)
	return
}

func NewCreateServiceRequest() (request *CreateServiceRequest) {
	request = &CreateServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateService")
	return
}

func NewCreateServiceResponse() (response *CreateServiceResponse) {
	response = &CreateServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建服务
func (c *Client) CreateService(request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
	if request == nil {
		request = NewCreateServiceRequest()
	}
	response = NewCreateServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNodeStatusRequest() (request *DescribeNodeStatusRequest) {
	request = &DescribeNodeStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeNodeStatus")
	return
}

func NewDescribeNodeStatusResponse() (response *DescribeNodeStatusResponse) {
	response = &DescribeNodeStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点状态
func (c *Client) DescribeNodeStatus(request *DescribeNodeStatusRequest) (response *DescribeNodeStatusResponse, err error) {
	if request == nil {
		request = NewDescribeNodeStatusRequest()
	}
	response = NewDescribeNodeStatusResponse()
	err = c.Send(request, response)
	return
}

func NewSearchNotificationAlertsRequest() (request *SearchNotificationAlertsRequest) {
	request = &SearchNotificationAlertsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SearchNotificationAlerts")
	return
}

func NewSearchNotificationAlertsResponse() (response *SearchNotificationAlertsResponse) {
	response = &SearchNotificationAlertsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警消息内查询alert
func (c *Client) SearchNotificationAlerts(request *SearchNotificationAlertsRequest) (response *SearchNotificationAlertsResponse, err error) {
	if request == nil {
		request = NewSearchNotificationAlertsRequest()
	}
	response = NewSearchNotificationAlertsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSvrNicAllocationRequest() (request *DeleteSvrNicAllocationRequest) {
	request = &DeleteSvrNicAllocationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteSvrNicAllocation")
	return
}

func NewDeleteSvrNicAllocationResponse() (response *DeleteSvrNicAllocationResponse) {
	response = &DeleteSvrNicAllocationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DeleteSvrNicAllocation 解绑物理服务器多网卡策略
func (c *Client) DeleteSvrNicAllocation(request *DeleteSvrNicAllocationRequest) (response *DeleteSvrNicAllocationResponse, err error) {
	if request == nil {
		request = NewDeleteSvrNicAllocationRequest()
	}
	response = NewDeleteSvrNicAllocationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTcsMonitorRequest() (request *DescribeTcsMonitorRequest) {
	request = &DescribeTcsMonitorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeTcsMonitor")
	return
}

func NewDescribeTcsMonitorResponse() (response *DescribeTcsMonitorResponse) {
	response = &DescribeTcsMonitorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeTcsMonitor
func (c *Client) DescribeTcsMonitor(request *DescribeTcsMonitorRequest) (response *DescribeTcsMonitorResponse, err error) {
	if request == nil {
		request = NewDescribeTcsMonitorRequest()
	}
	response = NewDescribeTcsMonitorResponse()
	err = c.Send(request, response)
	return
}

func NewPerformServerTaskExRequest() (request *PerformServerTaskExRequest) {
	request = &PerformServerTaskExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "PerformServerTaskEx")
	return
}

func NewPerformServerTaskExResponse() (response *PerformServerTaskExResponse) {
	response = &PerformServerTaskExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 带外管理 开机 关机 重启
func (c *Client) PerformServerTaskEx(request *PerformServerTaskExRequest) (response *PerformServerTaskExResponse, err error) {
	if request == nil {
		request = NewPerformServerTaskExRequest()
	}
	response = NewPerformServerTaskExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceYAMLRequest() (request *DescribeServiceYAMLRequest) {
	request = &DescribeServiceYAMLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeServiceYAML")
	return
}

func NewDescribeServiceYAMLResponse() (response *DescribeServiceYAMLResponse) {
	response = &DescribeServiceYAMLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询yaml格式的服务详情
func (c *Client) DescribeServiceYAML(request *DescribeServiceYAMLRequest) (response *DescribeServiceYAMLResponse, err error) {
	if request == nil {
		request = NewDescribeServiceYAMLRequest()
	}
	response = NewDescribeServiceYAMLResponse()
	err = c.Send(request, response)
	return
}

func NewGetRouteRequest() (request *GetRouteRequest) {
	request = &GetRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetRoute")
	return
}

func NewGetRouteResponse() (response *GetRouteResponse) {
	response = &GetRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定路由
func (c *Client) GetRoute(request *GetRouteRequest) (response *GetRouteResponse, err error) {
	if request == nil {
		request = NewGetRouteRequest()
	}
	response = NewGetRouteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServerRelatedNetdevicesRequest() (request *DescribeServerRelatedNetdevicesRequest) {
	request = &DescribeServerRelatedNetdevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeServerRelatedNetdevices")
	return
}

func NewDescribeServerRelatedNetdevicesResponse() (response *DescribeServerRelatedNetdevicesResponse) {
	response = &DescribeServerRelatedNetdevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务器关联网络设备
func (c *Client) DescribeServerRelatedNetdevices(request *DescribeServerRelatedNetdevicesRequest) (response *DescribeServerRelatedNetdevicesResponse, err error) {
	if request == nil {
		request = NewDescribeServerRelatedNetdevicesRequest()
	}
	response = NewDescribeServerRelatedNetdevicesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateDashFolderByUidRequest() (request *UpdateDashFolderByUidRequest) {
	request = &UpdateDashFolderByUidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "UpdateDashFolderByUid")
	return
}

func NewUpdateDashFolderByUidResponse() (response *UpdateDashFolderByUidResponse) {
	response = &UpdateDashFolderByUidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新仪表盘目录
func (c *Client) UpdateDashFolderByUid(request *UpdateDashFolderByUidRequest) (response *UpdateDashFolderByUidResponse, err error) {
	if request == nil {
		request = NewUpdateDashFolderByUidRequest()
	}
	response = NewUpdateDashFolderByUidResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateServiceYAMLRequest() (request *UpdateServiceYAMLRequest) {
	request = &UpdateServiceYAMLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "UpdateServiceYAML")
	return
}

func NewUpdateServiceYAMLResponse() (response *UpdateServiceYAMLResponse) {
	response = &UpdateServiceYAMLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新服务
func (c *Client) UpdateServiceYAML(request *UpdateServiceYAMLRequest) (response *UpdateServiceYAMLResponse, err error) {
	if request == nil {
		request = NewUpdateServiceYAMLRequest()
	}
	response = NewUpdateServiceYAMLResponse()
	err = c.Send(request, response)
	return
}

func NewListDeploymentsRequest() (request *ListDeploymentsRequest) {
	request = &ListDeploymentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListDeployments")
	return
}

func NewListDeploymentsResponse() (response *ListDeploymentsResponse) {
	response = &ListDeploymentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询部署Deployment列表
func (c *Client) ListDeployments(request *ListDeploymentsRequest) (response *ListDeploymentsResponse, err error) {
	if request == nil {
		request = NewListDeploymentsRequest()
	}
	response = NewListDeploymentsResponse()
	err = c.Send(request, response)
	return
}

func NewListVirtualMachinesRequest() (request *ListVirtualMachinesRequest) {
	request = &ListVirtualMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListVirtualMachines")
	return
}

func NewListVirtualMachinesResponse() (response *ListVirtualMachinesResponse) {
	response = &ListVirtualMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 虚拟机列表
func (c *Client) ListVirtualMachines(request *ListVirtualMachinesRequest) (response *ListVirtualMachinesResponse, err error) {
	if request == nil {
		request = NewListVirtualMachinesRequest()
	}
	response = NewListVirtualMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLoadBalancerRequest() (request *ModifyLoadBalancerRequest) {
	request = &ModifyLoadBalancerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyLoadBalancer")
	return
}

func NewModifyLoadBalancerResponse() (response *ModifyLoadBalancerResponse) {
	response = &ModifyLoadBalancerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改负载均衡实例的属性
func (c *Client) ModifyLoadBalancer(request *ModifyLoadBalancerRequest) (response *ModifyLoadBalancerResponse, err error) {
	if request == nil {
		request = NewModifyLoadBalancerRequest()
	}
	response = NewModifyLoadBalancerResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteServiceRequest() (request *DeleteServiceRequest) {
	request = &DeleteServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteService")
	return
}

func NewDeleteServiceResponse() (response *DeleteServiceResponse) {
	response = &DeleteServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除服务
func (c *Client) DeleteService(request *DeleteServiceRequest) (response *DeleteServiceResponse, err error) {
	if request == nil {
		request = NewDeleteServiceRequest()
	}
	response = NewDeleteServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLabelRequest() (request *DescribeLabelRequest) {
	request = &DescribeLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeLabel")
	return
}

func NewDescribeLabelResponse() (response *DescribeLabelResponse) {
	response = &DescribeLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 标签管理--查询标签
func (c *Client) DescribeLabel(request *DescribeLabelRequest) (response *DescribeLabelResponse, err error) {
	if request == nil {
		request = NewDescribeLabelRequest()
	}
	response = NewDescribeLabelResponse()
	err = c.Send(request, response)
	return
}

func NewGetDataBaradMetricRequest() (request *GetDataBaradMetricRequest) {
	request = &GetDataBaradMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetDataBaradMetric")
	return
}

func NewGetDataBaradMetricResponse() (response *GetDataBaradMetricResponse) {
	response = &GetDataBaradMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取Barad指标数据接口，为了兼容Barad接口格式而创建。
func (c *Client) GetDataBaradMetric(request *GetDataBaradMetricRequest) (response *GetDataBaradMetricResponse, err error) {
	if request == nil {
		request = NewGetDataBaradMetricRequest()
	}
	response = NewGetDataBaradMetricResponse()
	err = c.Send(request, response)
	return
}

func NewListMonitorMetricsRequest() (request *ListMonitorMetricsRequest) {
	request = &ListMonitorMetricsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListMonitorMetrics")
	return
}

func NewListMonitorMetricsResponse() (response *ListMonitorMetricsResponse) {
	response = &ListMonitorMetricsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有可自定义指标列表
func (c *Client) ListMonitorMetrics(request *ListMonitorMetricsRequest) (response *ListMonitorMetricsResponse, err error) {
	if request == nil {
		request = NewListMonitorMetricsRequest()
	}
	response = NewListMonitorMetricsResponse()
	err = c.Send(request, response)
	return
}

func NewGetClusterInstantMetricDataRequest() (request *GetClusterInstantMetricDataRequest) {
	request = &GetClusterInstantMetricDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetClusterInstantMetricData")
	return
}

func NewGetClusterInstantMetricDataResponse() (response *GetClusterInstantMetricDataResponse) {
	response = &GetClusterInstantMetricDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询cluster的某个时间点监控数据
func (c *Client) GetClusterInstantMetricData(request *GetClusterInstantMetricDataRequest) (response *GetClusterInstantMetricDataResponse, err error) {
	if request == nil {
		request = NewGetClusterInstantMetricDataRequest()
	}
	response = NewGetClusterInstantMetricDataResponse()
	err = c.Send(request, response)
	return
}

func NewListNodesRequest() (request *ListNodesRequest) {
	request = &ListNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListNodes")
	return
}

func NewListNodesResponse() (response *ListNodesResponse) {
	response = &ListNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点列表
func (c *Client) ListNodes(request *ListNodesRequest) (response *ListNodesResponse, err error) {
	if request == nil {
		request = NewListNodesRequest()
	}
	response = NewListNodesResponse()
	err = c.Send(request, response)
	return
}

func NewListVirtualMachineStoragesRequest() (request *ListVirtualMachineStoragesRequest) {
	request = &ListVirtualMachineStoragesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListVirtualMachineStorages")
	return
}

func NewListVirtualMachineStoragesResponse() (response *ListVirtualMachineStoragesResponse) {
	response = &ListVirtualMachineStoragesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 虚拟机存储列表
func (c *Client) ListVirtualMachineStorages(request *ListVirtualMachineStoragesRequest) (response *ListVirtualMachineStoragesResponse, err error) {
	if request == nil {
		request = NewListVirtualMachineStoragesRequest()
	}
	response = NewListVirtualMachineStoragesResponse()
	err = c.Send(request, response)
	return
}

func NewAddImageRequest() (request *AddImageRequest) {
	request = &AddImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "AddImage")
	return
}

func NewAddImageResponse() (response *AddImageResponse) {
	response = &AddImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增镜像
func (c *Client) AddImage(request *AddImageRequest) (response *AddImageResponse, err error) {
	if request == nil {
		request = NewAddImageRequest()
	}
	response = NewAddImageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDevOpsStatusRequest() (request *ModifyDevOpsStatusRequest) {
	request = &ModifyDevOpsStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyDevOpsStatus")
	return
}

func NewModifyDevOpsStatusResponse() (response *ModifyDevOpsStatusResponse) {
	response = &ModifyDevOpsStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启关闭集群的自愈场景
func (c *Client) ModifyDevOpsStatus(request *ModifyDevOpsStatusRequest) (response *ModifyDevOpsStatusResponse, err error) {
	if request == nil {
		request = NewModifyDevOpsStatusRequest()
	}
	response = NewModifyDevOpsStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePodRequest() (request *DeletePodRequest) {
	request = &DeletePodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeletePod")
	return
}

func NewDeletePodResponse() (response *DeletePodResponse) {
	response = &DeletePodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Pod
func (c *Client) DeletePod(request *DeletePodRequest) (response *DeletePodResponse, err error) {
	if request == nil {
		request = NewDeletePodRequest()
	}
	response = NewDeletePodResponse()
	err = c.Send(request, response)
	return
}

func NewGetCloudProductStatusRequest() (request *GetCloudProductStatusRequest) {
	request = &GetCloudProductStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetCloudProductStatus")
	return
}

func NewGetCloudProductStatusResponse() (response *GetCloudProductStatusResponse) {
	response = &GetCloudProductStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云产品状态
func (c *Client) GetCloudProductStatus(request *GetCloudProductStatusRequest) (response *GetCloudProductStatusResponse, err error) {
	if request == nil {
		request = NewGetCloudProductStatusRequest()
	}
	response = NewGetCloudProductStatusResponse()
	err = c.Send(request, response)
	return
}

func NewListAppInstancesRequest() (request *ListAppInstancesRequest) {
	request = &ListAppInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListAppInstances")
	return
}

func NewListAppInstancesResponse() (response *ListAppInstancesResponse) {
	response = &ListAppInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用列表
func (c *Client) ListAppInstances(request *ListAppInstancesRequest) (response *ListAppInstancesResponse, err error) {
	if request == nil {
		request = NewListAppInstancesRequest()
	}
	response = NewListAppInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterVersionRequest() (request *DescribeClusterVersionRequest) {
	request = &DescribeClusterVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeClusterVersion")
	return
}

func NewDescribeClusterVersionResponse() (response *DescribeClusterVersionResponse) {
	response = &DescribeClusterVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群版本信息
func (c *Client) DescribeClusterVersion(request *DescribeClusterVersionRequest) (response *DescribeClusterVersionResponse, err error) {
	if request == nil {
		request = NewDescribeClusterVersionRequest()
	}
	response = NewDescribeClusterVersionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCronJobRequest() (request *CreateCronJobRequest) {
	request = &CreateCronJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateCronJob")
	return
}

func NewCreateCronJobResponse() (response *CreateCronJobResponse) {
	response = &CreateCronJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建CronJob
func (c *Client) CreateCronJob(request *CreateCronJobRequest) (response *CreateCronJobResponse, err error) {
	if request == nil {
		request = NewCreateCronJobRequest()
	}
	response = NewCreateCronJobResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDaemonSetRequest() (request *ModifyDaemonSetRequest) {
	request = &ModifyDaemonSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyDaemonSet")
	return
}

func NewModifyDaemonSetResponse() (response *ModifyDaemonSetResponse) {
	response = &ModifyDaemonSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新DaemonSet
func (c *Client) ModifyDaemonSet(request *ModifyDaemonSetRequest) (response *ModifyDaemonSetResponse, err error) {
	if request == nil {
		request = NewModifyDaemonSetRequest()
	}
	response = NewModifyDaemonSetResponse()
	err = c.Send(request, response)
	return
}

func NewListNetworkPoliciesRequest() (request *ListNetworkPoliciesRequest) {
	request = &ListNetworkPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListNetworkPolicies")
	return
}

func NewListNetworkPoliciesResponse() (response *ListNetworkPoliciesResponse) {
	response = &ListNetworkPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络策略列表
func (c *Client) ListNetworkPolicies(request *ListNetworkPoliciesRequest) (response *ListNetworkPoliciesResponse, err error) {
	if request == nil {
		request = NewListNetworkPoliciesRequest()
	}
	response = NewListNetworkPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMultiNicDefinitionRequest() (request *CreateMultiNicDefinitionRequest) {
	request = &CreateMultiNicDefinitionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateMultiNicDefinition")
	return
}

func NewCreateMultiNicDefinitionResponse() (response *CreateMultiNicDefinitionResponse) {
	response = &CreateMultiNicDefinitionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateMultiNicDefinition 增加网卡信息
func (c *Client) CreateMultiNicDefinition(request *CreateMultiNicDefinitionRequest) (response *CreateMultiNicDefinitionResponse, err error) {
	if request == nil {
		request = NewCreateMultiNicDefinitionRequest()
	}
	response = NewCreateMultiNicDefinitionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteIngressRequest() (request *DeleteIngressRequest) {
	request = &DeleteIngressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteIngress")
	return
}

func NewDeleteIngressResponse() (response *DeleteIngressResponse) {
	response = &DeleteIngressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Ingress
func (c *Client) DeleteIngress(request *DeleteIngressRequest) (response *DeleteIngressResponse, err error) {
	if request == nil {
		request = NewDeleteIngressRequest()
	}
	response = NewDeleteIngressResponse()
	err = c.Send(request, response)
	return
}

func NewQueryServerStatusRequest() (request *QueryServerStatusRequest) {
	request = &QueryServerStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "QueryServerStatus")
	return
}

func NewQueryServerStatusResponse() (response *QueryServerStatusResponse) {
	response = &QueryServerStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询任务接口
func (c *Client) QueryServerStatus(request *QueryServerStatusRequest) (response *QueryServerStatusResponse, err error) {
	if request == nil {
		request = NewQueryServerStatusRequest()
	}
	response = NewQueryServerStatusResponse()
	err = c.Send(request, response)
	return
}

func NewServerSpecialAllocateLanIPExRequest() (request *ServerSpecialAllocateLanIPExRequest) {
	request = &ServerSpecialAllocateLanIPExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ServerSpecialAllocateLanIPEx")
	return
}

func NewServerSpecialAllocateLanIPExResponse() (response *ServerSpecialAllocateLanIPExResponse) {
	response = &ServerSpecialAllocateLanIPExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器批量分配相应虚拟内网段IP接口
func (c *Client) ServerSpecialAllocateLanIPEx(request *ServerSpecialAllocateLanIPExRequest) (response *ServerSpecialAllocateLanIPExResponse, err error) {
	if request == nil {
		request = NewServerSpecialAllocateLanIPExRequest()
	}
	response = NewServerSpecialAllocateLanIPExResponse()
	err = c.Send(request, response)
	return
}

func NewConfirmAlertRequest() (request *ConfirmAlertRequest) {
	request = &ConfirmAlertRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ConfirmAlert")
	return
}

func NewConfirmAlertResponse() (response *ConfirmAlertResponse) {
	response = &ConfirmAlertResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 确认指定Alert
func (c *Client) ConfirmAlert(request *ConfirmAlertRequest) (response *ConfirmAlertResponse, err error) {
	if request == nil {
		request = NewConfirmAlertRequest()
	}
	response = NewConfirmAlertResponse()
	err = c.Send(request, response)
	return
}

func NewCreateStatefulSetRequest() (request *CreateStatefulSetRequest) {
	request = &CreateStatefulSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateStatefulSet")
	return
}

func NewCreateStatefulSetResponse() (response *CreateStatefulSetResponse) {
	response = &CreateStatefulSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建有状态应用
func (c *Client) CreateStatefulSet(request *CreateStatefulSetRequest) (response *CreateStatefulSetResponse, err error) {
	if request == nil {
		request = NewCreateStatefulSetRequest()
	}
	response = NewCreateStatefulSetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterAttributesRequest() (request *DescribeClusterAttributesRequest) {
	request = &DescribeClusterAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeClusterAttributes")
	return
}

func NewDescribeClusterAttributesResponse() (response *DescribeClusterAttributesResponse) {
	response = &DescribeClusterAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群基本信息
func (c *Client) DescribeClusterAttributes(request *DescribeClusterAttributesRequest) (response *DescribeClusterAttributesResponse, err error) {
	if request == nil {
		request = NewDescribeClusterAttributesRequest()
	}
	response = NewDescribeClusterAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrUpdateRuleGroupRequest() (request *CreateOrUpdateRuleGroupRequest) {
	request = &CreateOrUpdateRuleGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateOrUpdateRuleGroup")
	return
}

func NewCreateOrUpdateRuleGroupResponse() (response *CreateOrUpdateRuleGroupResponse) {
	response = &CreateOrUpdateRuleGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建或更新策略
func (c *Client) CreateOrUpdateRuleGroup(request *CreateOrUpdateRuleGroupRequest) (response *CreateOrUpdateRuleGroupResponse, err error) {
	if request == nil {
		request = NewCreateOrUpdateRuleGroupRequest()
	}
	response = NewCreateOrUpdateRuleGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMultiNicDefinitionRequest() (request *DescribeMultiNicDefinitionRequest) {
	request = &DescribeMultiNicDefinitionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeMultiNicDefinition")
	return
}

func NewDescribeMultiNicDefinitionResponse() (response *DescribeMultiNicDefinitionResponse) {
	response = &DescribeMultiNicDefinitionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网卡信息 DescribeMultiNicDefinition
func (c *Client) DescribeMultiNicDefinition(request *DescribeMultiNicDefinitionRequest) (response *DescribeMultiNicDefinitionResponse, err error) {
	if request == nil {
		request = NewDescribeMultiNicDefinitionRequest()
	}
	response = NewDescribeMultiNicDefinitionResponse()
	err = c.Send(request, response)
	return
}

func NewGetDashboardVersionsByIdRequest() (request *GetDashboardVersionsByIdRequest) {
	request = &GetDashboardVersionsByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetDashboardVersionsById")
	return
}

func NewGetDashboardVersionsByIdResponse() (response *GetDashboardVersionsByIdResponse) {
	response = &GetDashboardVersionsByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根椐仪表ID查询变更历史
func (c *Client) GetDashboardVersionsById(request *GetDashboardVersionsByIdRequest) (response *GetDashboardVersionsByIdResponse, err error) {
	if request == nil {
		request = NewGetDashboardVersionsByIdRequest()
	}
	response = NewGetDashboardVersionsByIdResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePodsStatsRequest() (request *DescribePodsStatsRequest) {
	request = &DescribePodsStatsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribePodsStats")
	return
}

func NewDescribePodsStatsResponse() (response *DescribePodsStatsResponse) {
	response = &DescribePodsStatsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Pod统计信息
func (c *Client) DescribePodsStats(request *DescribePodsStatsRequest) (response *DescribePodsStatsResponse, err error) {
	if request == nil {
		request = NewDescribePodsStatsRequest()
	}
	response = NewDescribePodsStatsResponse()
	err = c.Send(request, response)
	return
}

func NewSearchRuleGroupTemplesRequest() (request *SearchRuleGroupTemplesRequest) {
	request = &SearchRuleGroupTemplesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SearchRuleGroupTemples")
	return
}

func NewSearchRuleGroupTemplesResponse() (response *SearchRuleGroupTemplesResponse) {
	response = &SearchRuleGroupTemplesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件查询策略模板
func (c *Client) SearchRuleGroupTemples(request *SearchRuleGroupTemplesRequest) (response *SearchRuleGroupTemplesResponse, err error) {
	if request == nil {
		request = NewSearchRuleGroupTemplesRequest()
	}
	response = NewSearchRuleGroupTemplesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRetentionsRequest() (request *DeleteRetentionsRequest) {
	request = &DeleteRetentionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteRetentions")
	return
}

func NewDeleteRetentionsResponse() (response *DeleteRetentionsResponse) {
	response = &DeleteRetentionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Retention
func (c *Client) DeleteRetentions(request *DeleteRetentionsRequest) (response *DeleteRetentionsResponse, err error) {
	if request == nil {
		request = NewDeleteRetentionsRequest()
	}
	response = NewDeleteRetentionsResponse()
	err = c.Send(request, response)
	return
}

func NewGetNodeHistoryMetricDataRequest() (request *GetNodeHistoryMetricDataRequest) {
	request = &GetNodeHistoryMetricDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetNodeHistoryMetricData")
	return
}

func NewGetNodeHistoryMetricDataResponse() (response *GetNodeHistoryMetricDataResponse) {
	response = &GetNodeHistoryMetricDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取节点历史性能数据
func (c *Client) GetNodeHistoryMetricData(request *GetNodeHistoryMetricDataRequest) (response *GetNodeHistoryMetricDataResponse, err error) {
	if request == nil {
		request = NewGetNodeHistoryMetricDataRequest()
	}
	response = NewGetNodeHistoryMetricDataResponse()
	err = c.Send(request, response)
	return
}

func NewGetPodInstantMetricDataRequest() (request *GetPodInstantMetricDataRequest) {
	request = &GetPodInstantMetricDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetPodInstantMetricData")
	return
}

func NewGetPodInstantMetricDataResponse() (response *GetPodInstantMetricDataResponse) {
	response = &GetPodInstantMetricDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询pod的某个时间点监控数据
func (c *Client) GetPodInstantMetricData(request *GetPodInstantMetricDataRequest) (response *GetPodInstantMetricDataResponse, err error) {
	if request == nil {
		request = NewGetPodInstantMetricDataRequest()
	}
	response = NewGetPodInstantMetricDataResponse()
	err = c.Send(request, response)
	return
}

func NewRecycleSvrNicIPRequest() (request *RecycleSvrNicIPRequest) {
	request = &RecycleSvrNicIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "RecycleSvrNicIP")
	return
}

func NewRecycleSvrNicIPResponse() (response *RecycleSvrNicIPResponse) {
	response = &RecycleSvrNicIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// RecycleSvrNicIP 服务器的多网卡ip回收
func (c *Client) RecycleSvrNicIP(request *RecycleSvrNicIPRequest) (response *RecycleSvrNicIPResponse, err error) {
	if request == nil {
		request = NewRecycleSvrNicIPRequest()
	}
	response = NewRecycleSvrNicIPResponse()
	err = c.Send(request, response)
	return
}

func NewAddServersRequest() (request *AddServersRequest) {
	request = &AddServersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "AddServers")
	return
}

func NewAddServersResponse() (response *AddServersResponse) {
	response = &AddServersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加服务器信息
func (c *Client) AddServers(request *AddServersRequest) (response *AddServersResponse, err error) {
	if request == nil {
		request = NewAddServersRequest()
	}
	response = NewAddServersResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDashboardByUidRequest() (request *DeleteDashboardByUidRequest) {
	request = &DeleteDashboardByUidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteDashboardByUid")
	return
}

func NewDeleteDashboardByUidResponse() (response *DeleteDashboardByUidResponse) {
	response = &DeleteDashboardByUidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根椐唯一ID删除仪表盘
func (c *Client) DeleteDashboardByUid(request *DeleteDashboardByUidRequest) (response *DeleteDashboardByUidResponse, err error) {
	if request == nil {
		request = NewDeleteDashboardByUidRequest()
	}
	response = NewDeleteDashboardByUidResponse()
	err = c.Send(request, response)
	return
}

func NewListMetaResourceTypeRequest() (request *ListMetaResourceTypeRequest) {
	request = &ListMetaResourceTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListMetaResourceType")
	return
}

func NewListMetaResourceTypeResponse() (response *ListMetaResourceTypeResponse) {
	response = &ListMetaResourceTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资源对象类型列表
func (c *Client) ListMetaResourceType(request *ListMetaResourceTypeRequest) (response *ListMetaResourceTypeResponse, err error) {
	if request == nil {
		request = NewListMetaResourceTypeRequest()
	}
	response = NewListMetaResourceTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteHPARequest() (request *DeleteHPARequest) {
	request = &DeleteHPARequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteHPA")
	return
}

func NewDeleteHPAResponse() (response *DeleteHPAResponse) {
	response = &DeleteHPAResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除HPA
func (c *Client) DeleteHPA(request *DeleteHPARequest) (response *DeleteHPAResponse, err error) {
	if request == nil {
		request = NewDeleteHPARequest()
	}
	response = NewDeleteHPAResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIngressRequest() (request *DescribeIngressRequest) {
	request = &DescribeIngressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeIngress")
	return
}

func NewDescribeIngressResponse() (response *DescribeIngressResponse) {
	response = &DescribeIngressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Ingress
func (c *Client) DescribeIngress(request *DescribeIngressRequest) (response *DescribeIngressResponse, err error) {
	if request == nil {
		request = NewDescribeIngressRequest()
	}
	response = NewDescribeIngressResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMetaResourceTypeRequest() (request *UpdateMetaResourceTypeRequest) {
	request = &UpdateMetaResourceTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "UpdateMetaResourceType")
	return
}

func NewUpdateMetaResourceTypeResponse() (response *UpdateMetaResourceTypeResponse) {
	response = &UpdateMetaResourceTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新资源对象类型信息
func (c *Client) UpdateMetaResourceType(request *UpdateMetaResourceTypeRequest) (response *UpdateMetaResourceTypeResponse, err error) {
	if request == nil {
		request = NewUpdateMetaResourceTypeRequest()
	}
	response = NewUpdateMetaResourceTypeResponse()
	err = c.Send(request, response)
	return
}

func NewAnalyseLogsRequest() (request *AnalyseLogsRequest) {
	request = &AnalyseLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "AnalyseLogs")
	return
}

func NewAnalyseLogsResponse() (response *AnalyseLogsResponse) {
	response = &AnalyseLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志分析
func (c *Client) AnalyseLogs(request *AnalyseLogsRequest) (response *AnalyseLogsResponse, err error) {
	if request == nil {
		request = NewAnalyseLogsRequest()
	}
	response = NewAnalyseLogsResponse()
	err = c.Send(request, response)
	return
}

func NewGetHealthStatusRequest() (request *GetHealthStatusRequest) {
	request = &GetHealthStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetHealthStatus")
	return
}

func NewGetHealthStatusResponse() (response *GetHealthStatusResponse) {
	response = &GetHealthStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取健康状态概览
func (c *Client) GetHealthStatus(request *GetHealthStatusRequest) (response *GetHealthStatusResponse, err error) {
	if request == nil {
		request = NewGetHealthStatusRequest()
	}
	response = NewGetHealthStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNodeAttributesRequest() (request *DescribeNodeAttributesRequest) {
	request = &DescribeNodeAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeNodeAttributes")
	return
}

func NewDescribeNodeAttributesResponse() (response *DescribeNodeAttributesResponse) {
	response = &DescribeNodeAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点基本信息
func (c *Client) DescribeNodeAttributes(request *DescribeNodeAttributesRequest) (response *DescribeNodeAttributesResponse, err error) {
	if request == nil {
		request = NewDescribeNodeAttributesRequest()
	}
	response = NewDescribeNodeAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeStatefulSetRequest() (request *DescribeStatefulSetRequest) {
	request = &DescribeStatefulSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeStatefulSet")
	return
}

func NewDescribeStatefulSetResponse() (response *DescribeStatefulSetResponse) {
	response = &DescribeStatefulSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询StatefulSet
func (c *Client) DescribeStatefulSet(request *DescribeStatefulSetRequest) (response *DescribeStatefulSetResponse, err error) {
	if request == nil {
		request = NewDescribeStatefulSetRequest()
	}
	response = NewDescribeStatefulSetResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePasswordRequest() (request *UpdatePasswordRequest) {
	request = &UpdatePasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "UpdatePassword")
	return
}

func NewUpdatePasswordResponse() (response *UpdatePasswordResponse) {
	response = &UpdatePasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新某个节点的密码元数据
func (c *Client) UpdatePassword(request *UpdatePasswordRequest) (response *UpdatePasswordResponse, err error) {
	if request == nil {
		request = NewUpdatePasswordRequest()
	}
	response = NewUpdatePasswordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRelocateInfoRequest() (request *DescribeRelocateInfoRequest) {
	request = &DescribeRelocateInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeRelocateInfo")
	return
}

func NewDescribeRelocateInfoResponse() (response *DescribeRelocateInfoResponse) {
	response = &DescribeRelocateInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询搬迁目的机架机位信息
func (c *Client) DescribeRelocateInfo(request *DescribeRelocateInfoRequest) (response *DescribeRelocateInfoResponse, err error) {
	if request == nil {
		request = NewDescribeRelocateInfoRequest()
	}
	response = NewDescribeRelocateInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServersRequest() (request *DescribeServersRequest) {
	request = &DescribeServersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeServers")
	return
}

func NewDescribeServersResponse() (response *DescribeServersResponse) {
	response = &DescribeServersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务器信息
func (c *Client) DescribeServers(request *DescribeServersRequest) (response *DescribeServersResponse, err error) {
	if request == nil {
		request = NewDescribeServersRequest()
	}
	response = NewDescribeServersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSystemComponentsRequest() (request *DescribeSystemComponentsRequest) {
	request = &DescribeSystemComponentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeSystemComponents")
	return
}

func NewDescribeSystemComponentsResponse() (response *DescribeSystemComponentsResponse) {
	response = &DescribeSystemComponentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群系统服务组件(etcd, api-server, controller-manager, scheduler)
func (c *Client) DescribeSystemComponents(request *DescribeSystemComponentsRequest) (response *DescribeSystemComponentsResponse, err error) {
	if request == nil {
		request = NewDescribeSystemComponentsRequest()
	}
	response = NewDescribeSystemComponentsResponse()
	err = c.Send(request, response)
	return
}

func NewGetMetricMatrixByDimensionRequest() (request *GetMetricMatrixByDimensionRequest) {
	request = &GetMetricMatrixByDimensionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetMetricMatrixByDimension")
	return
}

func NewGetMetricMatrixByDimensionResponse() (response *GetMetricMatrixByDimensionResponse) {
	response = &GetMetricMatrixByDimensionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取matrix格式的监控数据
func (c *Client) GetMetricMatrixByDimension(request *GetMetricMatrixByDimensionRequest) (response *GetMetricMatrixByDimensionResponse, err error) {
	if request == nil {
		request = NewGetMetricMatrixByDimensionRequest()
	}
	response = NewGetMetricMatrixByDimensionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLogConfigRequest() (request *CreateLogConfigRequest) {
	request = &CreateLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateLogConfig")
	return
}

func NewCreateLogConfigResponse() (response *CreateLogConfigResponse) {
	response = &CreateLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建日志接入配置
func (c *Client) CreateLogConfig(request *CreateLogConfigRequest) (response *CreateLogConfigResponse, err error) {
	if request == nil {
		request = NewCreateLogConfigRequest()
	}
	response = NewCreateLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewListDaemonSetsRequest() (request *ListDaemonSetsRequest) {
	request = &ListDaemonSetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListDaemonSets")
	return
}

func NewListDaemonSetsResponse() (response *ListDaemonSetsResponse) {
	response = &ListDaemonSetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询DaemonSet列表
func (c *Client) ListDaemonSets(request *ListDaemonSetsRequest) (response *ListDaemonSetsResponse, err error) {
	if request == nil {
		request = NewListDaemonSetsRequest()
	}
	response = NewListDaemonSetsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCronJobRequest() (request *ModifyCronJobRequest) {
	request = &ModifyCronJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyCronJob")
	return
}

func NewModifyCronJobResponse() (response *ModifyCronJobResponse) {
	response = &ModifyCronJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新CronJob
func (c *Client) ModifyCronJob(request *ModifyCronJobRequest) (response *ModifyCronJobResponse, err error) {
	if request == nil {
		request = NewModifyCronJobRequest()
	}
	response = NewModifyCronJobResponse()
	err = c.Send(request, response)
	return
}

func NewServerJoinToBanksRequest() (request *ServerJoinToBanksRequest) {
	request = &ServerJoinToBanksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ServerJoinToBanks")
	return
}

func NewServerJoinToBanksResponse() (response *ServerJoinToBanksResponse) {
	response = &ServerJoinToBanksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ServerJoinToBanks 服务器密码入库
func (c *Client) ServerJoinToBanks(request *ServerJoinToBanksRequest) (response *ServerJoinToBanksResponse, err error) {
	if request == nil {
		request = NewServerJoinToBanksRequest()
	}
	response = NewServerJoinToBanksResponse()
	err = c.Send(request, response)
	return
}

func NewGetPodHistoryMetricDataRequest() (request *GetPodHistoryMetricDataRequest) {
	request = &GetPodHistoryMetricDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetPodHistoryMetricData")
	return
}

func NewGetPodHistoryMetricDataResponse() (response *GetPodHistoryMetricDataResponse) {
	response = &GetPodHistoryMetricDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Pod历史性能数据
func (c *Client) GetPodHistoryMetricData(request *GetPodHistoryMetricDataRequest) (response *GetPodHistoryMetricDataResponse, err error) {
	if request == nil {
		request = NewGetPodHistoryMetricDataRequest()
	}
	response = NewGetPodHistoryMetricDataResponse()
	err = c.Send(request, response)
	return
}

func NewRecycleVMVirtualIPRequest() (request *RecycleVMVirtualIPRequest) {
	request = &RecycleVMVirtualIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "RecycleVMVirtualIP")
	return
}

func NewRecycleVMVirtualIPResponse() (response *RecycleVMVirtualIPResponse) {
	response = &RecycleVMVirtualIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回收VM虚拟内外网IP
func (c *Client) RecycleVMVirtualIP(request *RecycleVMVirtualIPRequest) (response *RecycleVMVirtualIPResponse, err error) {
	if request == nil {
		request = NewRecycleVMVirtualIPRequest()
	}
	response = NewRecycleVMVirtualIPResponse()
	err = c.Send(request, response)
	return
}

func NewAddOutbandStrategyRequest() (request *AddOutbandStrategyRequest) {
	request = &AddOutbandStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "AddOutbandStrategy")
	return
}

func NewAddOutbandStrategyResponse() (response *AddOutbandStrategyResponse) {
	response = &AddOutbandStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加自定义带外密码策略
func (c *Client) AddOutbandStrategy(request *AddOutbandStrategyRequest) (response *AddOutbandStrategyResponse, err error) {
	if request == nil {
		request = NewAddOutbandStrategyRequest()
	}
	response = NewAddOutbandStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewImportCmdbServersRequest() (request *ImportCmdbServersRequest) {
	request = &ImportCmdbServersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ImportCmdbServers")
	return
}

func NewImportCmdbServersResponse() (response *ImportCmdbServersResponse) {
	response = &ImportCmdbServersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入物理服务器
func (c *Client) ImportCmdbServers(request *ImportCmdbServersRequest) (response *ImportCmdbServersResponse, err error) {
	if request == nil {
		request = NewImportCmdbServersRequest()
	}
	response = NewImportCmdbServersResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVirtualMachineRequest() (request *DeleteVirtualMachineRequest) {
	request = &DeleteVirtualMachineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteVirtualMachine")
	return
}

func NewDeleteVirtualMachineResponse() (response *DeleteVirtualMachineResponse) {
	response = &DeleteVirtualMachineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除虚拟机
func (c *Client) DeleteVirtualMachine(request *DeleteVirtualMachineRequest) (response *DeleteVirtualMachineResponse, err error) {
	if request == nil {
		request = NewDeleteVirtualMachineRequest()
	}
	response = NewDeleteVirtualMachineResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDashFolderRequest() (request *CreateDashFolderRequest) {
	request = &CreateDashFolderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateDashFolder")
	return
}

func NewCreateDashFolderResponse() (response *CreateDashFolderResponse) {
	response = &CreateDashFolderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建仪表盘目录
func (c *Client) CreateDashFolder(request *CreateDashFolderRequest) (response *CreateDashFolderResponse, err error) {
	if request == nil {
		request = NewCreateDashFolderRequest()
	}
	response = NewCreateDashFolderResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSilenceRequest() (request *CreateSilenceRequest) {
	request = &CreateSilenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateSilence")
	return
}

func NewCreateSilenceResponse() (response *CreateSilenceResponse) {
	response = &CreateSilenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建屏蔽策略
func (c *Client) CreateSilence(request *CreateSilenceRequest) (response *CreateSilenceResponse, err error) {
	if request == nil {
		request = NewCreateSilenceRequest()
	}
	response = NewCreateSilenceResponse()
	err = c.Send(request, response)
	return
}

func NewListContainersRequest() (request *ListContainersRequest) {
	request = &ListContainersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListContainers")
	return
}

func NewListContainersResponse() (response *ListContainersResponse) {
	response = &ListContainersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出pod里的所有容器
func (c *Client) ListContainers(request *ListContainersRequest) (response *ListContainersResponse, err error) {
	if request == nil {
		request = NewListContainersRequest()
	}
	response = NewListContainersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateIngressClassRequest() (request *CreateIngressClassRequest) {
	request = &CreateIngressClassRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateIngressClass")
	return
}

func NewCreateIngressClassResponse() (response *CreateIngressClassResponse) {
	response = &CreateIngressClassResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建路由控制器
func (c *Client) CreateIngressClass(request *CreateIngressClassRequest) (response *CreateIngressClassResponse, err error) {
	if request == nil {
		request = NewCreateIngressClassRequest()
	}
	response = NewCreateIngressClassResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateVMLanIPListRequest() (request *AllocateVMLanIPListRequest) {
	request = &AllocateVMLanIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "AllocateVMLanIPList")
	return
}

func NewAllocateVMLanIPListResponse() (response *AllocateVMLanIPListResponse) {
	response = &AllocateVMLanIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分配VM内网IP
func (c *Client) AllocateVMLanIPList(request *AllocateVMLanIPListRequest) (response *AllocateVMLanIPListResponse, err error) {
	if request == nil {
		request = NewAllocateVMLanIPListRequest()
	}
	response = NewAllocateVMLanIPListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCronJobListRequest() (request *DescribeCronJobListRequest) {
	request = &DescribeCronJobListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeCronJobList")
	return
}

func NewDescribeCronJobListResponse() (response *DescribeCronJobListResponse) {
	response = &DescribeCronJobListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询CronJob列表
func (c *Client) DescribeCronJobList(request *DescribeCronJobListRequest) (response *DescribeCronJobListResponse, err error) {
	if request == nil {
		request = NewDescribeCronJobListRequest()
	}
	response = NewDescribeCronJobListResponse()
	err = c.Send(request, response)
	return
}

func NewGetClusterHistoryMetricDataRequest() (request *GetClusterHistoryMetricDataRequest) {
	request = &GetClusterHistoryMetricDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetClusterHistoryMetricData")
	return
}

func NewGetClusterHistoryMetricDataResponse() (response *GetClusterHistoryMetricDataResponse) {
	response = &GetClusterHistoryMetricDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群历史性能数据
func (c *Client) GetClusterHistoryMetricData(request *GetClusterHistoryMetricDataRequest) (response *GetClusterHistoryMetricDataResponse, err error) {
	if request == nil {
		request = NewGetClusterHistoryMetricDataRequest()
	}
	response = NewGetClusterHistoryMetricDataResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIngressRequest() (request *ModifyIngressRequest) {
	request = &ModifyIngressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyIngress")
	return
}

func NewModifyIngressResponse() (response *ModifyIngressResponse) {
	response = &ModifyIngressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改Ingress
func (c *Client) ModifyIngress(request *ModifyIngressRequest) (response *ModifyIngressResponse, err error) {
	if request == nil {
		request = NewModifyIngressRequest()
	}
	response = NewModifyIngressResponse()
	err = c.Send(request, response)
	return
}

func NewGetSilenceRequest() (request *GetSilenceRequest) {
	request = &GetSilenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetSilence")
	return
}

func NewGetSilenceResponse() (response *GetSilenceResponse) {
	response = &GetSilenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取屏蔽规则
func (c *Client) GetSilence(request *GetSilenceRequest) (response *GetSilenceResponse, err error) {
	if request == nil {
		request = NewGetSilenceRequest()
	}
	response = NewGetSilenceResponse()
	err = c.Send(request, response)
	return
}

func NewRebootNodeRequest() (request *RebootNodeRequest) {
	request = &RebootNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "RebootNode")
	return
}

func NewRebootNodeResponse() (response *RebootNodeResponse) {
	response = &RebootNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过SSH登录宿主机执行重启主机命令
func (c *Client) RebootNode(request *RebootNodeRequest) (response *RebootNodeResponse, err error) {
	if request == nil {
		request = NewRebootNodeRequest()
	}
	response = NewRebootNodeResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateIngressYAMLRequest() (request *UpdateIngressYAMLRequest) {
	request = &UpdateIngressYAMLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "UpdateIngressYAML")
	return
}

func NewUpdateIngressYAMLResponse() (response *UpdateIngressYAMLResponse) {
	response = &UpdateIngressYAMLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新Ingress YAML
func (c *Client) UpdateIngressYAML(request *UpdateIngressYAMLRequest) (response *UpdateIngressYAMLResponse, err error) {
	if request == nil {
		request = NewUpdateIngressYAMLRequest()
	}
	response = NewUpdateIngressYAMLResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirtualMachineAttributesRequest() (request *DescribeVirtualMachineAttributesRequest) {
	request = &DescribeVirtualMachineAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeVirtualMachineAttributes")
	return
}

func NewDescribeVirtualMachineAttributesResponse() (response *DescribeVirtualMachineAttributesResponse) {
	response = &DescribeVirtualMachineAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 虚拟机基本信息
func (c *Client) DescribeVirtualMachineAttributes(request *DescribeVirtualMachineAttributesRequest) (response *DescribeVirtualMachineAttributesResponse, err error) {
	if request == nil {
		request = NewDescribeVirtualMachineAttributesRequest()
	}
	response = NewDescribeVirtualMachineAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTLSSecretRequest() (request *CreateTLSSecretRequest) {
	request = &CreateTLSSecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateTLSSecret")
	return
}

func NewCreateTLSSecretResponse() (response *CreateTLSSecretResponse) {
	response = &CreateTLSSecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建TSL密钥
func (c *Client) CreateTLSSecret(request *CreateTLSSecretRequest) (response *CreateTLSSecretResponse, err error) {
	if request == nil {
		request = NewCreateTLSSecretRequest()
	}
	response = NewCreateTLSSecretResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProjectRequest() (request *DescribeProjectRequest) {
	request = &DescribeProjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeProject")
	return
}

func NewDescribeProjectResponse() (response *DescribeProjectResponse) {
	response = &DescribeProjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取项目健康状态
func (c *Client) DescribeProject(request *DescribeProjectRequest) (response *DescribeProjectResponse, err error) {
	if request == nil {
		request = NewDescribeProjectRequest()
	}
	response = NewDescribeProjectResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNetworkPolicyRequest() (request *ModifyNetworkPolicyRequest) {
	request = &ModifyNetworkPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyNetworkPolicy")
	return
}

func NewModifyNetworkPolicyResponse() (response *ModifyNetworkPolicyResponse) {
	response = &ModifyNetworkPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改网络策略
func (c *Client) ModifyNetworkPolicy(request *ModifyNetworkPolicyRequest) (response *ModifyNetworkPolicyResponse, err error) {
	if request == nil {
		request = NewModifyNetworkPolicyRequest()
	}
	response = NewModifyNetworkPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewRecycleServerSpecialWanIPListRequest() (request *RecycleServerSpecialWanIPListRequest) {
	request = &RecycleServerSpecialWanIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "RecycleServerSpecialWanIPList")
	return
}

func NewRecycleServerSpecialWanIPListResponse() (response *RecycleServerSpecialWanIPListResponse) {
	response = &RecycleServerSpecialWanIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回收服务器特殊外网IP
func (c *Client) RecycleServerSpecialWanIPList(request *RecycleServerSpecialWanIPListRequest) (response *RecycleServerSpecialWanIPListResponse, err error) {
	if request == nil {
		request = NewRecycleServerSpecialWanIPListRequest()
	}
	response = NewRecycleServerSpecialWanIPListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRAIDListRequest() (request *DescribeRAIDListRequest) {
	request = &DescribeRAIDListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeRAIDList")
	return
}

func NewDescribeRAIDListResponse() (response *DescribeRAIDListResponse) {
	response = &DescribeRAIDListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询raid
func (c *Client) DescribeRAIDList(request *DescribeRAIDListRequest) (response *DescribeRAIDListResponse, err error) {
	if request == nil {
		request = NewDescribeRAIDListRequest()
	}
	response = NewDescribeRAIDListResponse()
	err = c.Send(request, response)
	return
}

func NewSearchDashboardsRequest() (request *SearchDashboardsRequest) {
	request = &SearchDashboardsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SearchDashboards")
	return
}

func NewSearchDashboardsResponse() (response *SearchDashboardsResponse) {
	response = &SearchDashboardsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询仪表盘
func (c *Client) SearchDashboards(request *SearchDashboardsRequest) (response *SearchDashboardsResponse, err error) {
	if request == nil {
		request = NewSearchDashboardsRequest()
	}
	response = NewSearchDashboardsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomScriptSetsRequest() (request *DescribeCustomScriptSetsRequest) {
	request = &DescribeCustomScriptSetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeCustomScriptSets")
	return
}

func NewDescribeCustomScriptSetsResponse() (response *DescribeCustomScriptSetsResponse) {
	response = &DescribeCustomScriptSetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询自定义脚本集信息
func (c *Client) DescribeCustomScriptSets(request *DescribeCustomScriptSetsRequest) (response *DescribeCustomScriptSetsResponse, err error) {
	if request == nil {
		request = NewDescribeCustomScriptSetsRequest()
	}
	response = NewDescribeCustomScriptSetsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVirtualMachineRequest() (request *CreateVirtualMachineRequest) {
	request = &CreateVirtualMachineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateVirtualMachine")
	return
}

func NewCreateVirtualMachineResponse() (response *CreateVirtualMachineResponse) {
	response = &CreateVirtualMachineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建vm
func (c *Client) CreateVirtualMachine(request *CreateVirtualMachineRequest) (response *CreateVirtualMachineResponse, err error) {
	if request == nil {
		request = NewCreateVirtualMachineRequest()
	}
	response = NewCreateVirtualMachineResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDashFolderByUidRequest() (request *DeleteDashFolderByUidRequest) {
	request = &DeleteDashFolderByUidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteDashFolderByUid")
	return
}

func NewDeleteDashFolderByUidResponse() (response *DeleteDashFolderByUidResponse) {
	response = &DeleteDashFolderByUidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除仪表盘目录
func (c *Client) DeleteDashFolderByUid(request *DeleteDashFolderByUidRequest) (response *DeleteDashFolderByUidResponse, err error) {
	if request == nil {
		request = NewDeleteDashFolderByUidRequest()
	}
	response = NewDeleteDashFolderByUidResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteResourceObjectRequest() (request *DeleteResourceObjectRequest) {
	request = &DeleteResourceObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteResourceObject")
	return
}

func NewDeleteResourceObjectResponse() (response *DeleteResourceObjectResponse) {
	response = &DeleteResourceObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除资源对象
func (c *Client) DeleteResourceObject(request *DeleteResourceObjectRequest) (response *DeleteResourceObjectResponse, err error) {
	if request == nil {
		request = NewDeleteResourceObjectRequest()
	}
	response = NewDeleteResourceObjectResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBasicOverviewRequest() (request *DescribeBasicOverviewRequest) {
	request = &DescribeBasicOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeBasicOverview")
	return
}

func NewDescribeBasicOverviewResponse() (response *DescribeBasicOverviewResponse) {
	response = &DescribeBasicOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回账户数量历史趋势, 镜像数量历史趋势, 应用数量历史趋势, 项目数量历史趋势, 及当前集群/结点/Pod/支撑服务状态
func (c *Client) DescribeBasicOverview(request *DescribeBasicOverviewRequest) (response *DescribeBasicOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeBasicOverviewRequest()
	}
	response = NewDescribeBasicOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewGetDashFolderPermissionsByUidRequest() (request *GetDashFolderPermissionsByUidRequest) {
	request = &GetDashFolderPermissionsByUidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetDashFolderPermissionsByUid")
	return
}

func NewGetDashFolderPermissionsByUidResponse() (response *GetDashFolderPermissionsByUidResponse) {
	response = &GetDashFolderPermissionsByUidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取仪表盘目录权限
func (c *Client) GetDashFolderPermissionsByUid(request *GetDashFolderPermissionsByUidRequest) (response *GetDashFolderPermissionsByUidResponse, err error) {
	if request == nil {
		request = NewGetDashFolderPermissionsByUidRequest()
	}
	response = NewGetDashFolderPermissionsByUidResponse()
	err = c.Send(request, response)
	return
}

func NewWithdrawServerExRequest() (request *WithdrawServerExRequest) {
	request = &WithdrawServerExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "WithdrawServerEx")
	return
}

func NewWithdrawServerExResponse() (response *WithdrawServerExResponse) {
	response = &WithdrawServerExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下架物理服务器
func (c *Client) WithdrawServerEx(request *WithdrawServerExRequest) (response *WithdrawServerExResponse, err error) {
	if request == nil {
		request = NewWithdrawServerExRequest()
	}
	response = NewWithdrawServerExResponse()
	err = c.Send(request, response)
	return
}

func NewCreateListenerRequest() (request *CreateListenerRequest) {
	request = &CreateListenerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateListener")
	return
}

func NewCreateListenerResponse() (response *CreateListenerResponse) {
	response = &CreateListenerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建负载均衡监听器
func (c *Client) CreateListener(request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
	if request == nil {
		request = NewCreateListenerRequest()
	}
	response = NewCreateListenerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOsDeployDebugInfoExRequest() (request *DescribeOsDeployDebugInfoExRequest) {
	request = &DescribeOsDeployDebugInfoExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeOsDeployDebugInfoEx")
	return
}

func NewDescribeOsDeployDebugInfoExResponse() (response *DescribeOsDeployDebugInfoExResponse) {
	response = &DescribeOsDeployDebugInfoExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询os部署调试信息
func (c *Client) DescribeOsDeployDebugInfoEx(request *DescribeOsDeployDebugInfoExRequest) (response *DescribeOsDeployDebugInfoExResponse, err error) {
	if request == nil {
		request = NewDescribeOsDeployDebugInfoExRequest()
	}
	response = NewDescribeOsDeployDebugInfoExResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterAttributesRequest() (request *ModifyClusterAttributesRequest) {
	request = &ModifyClusterAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyClusterAttributes")
	return
}

func NewModifyClusterAttributesResponse() (response *ModifyClusterAttributesResponse) {
	response = &ModifyClusterAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集群信息
func (c *Client) ModifyClusterAttributes(request *ModifyClusterAttributesRequest) (response *ModifyClusterAttributesResponse, err error) {
	if request == nil {
		request = NewModifyClusterAttributesRequest()
	}
	response = NewModifyClusterAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewRecycleServerWanIPListRequest() (request *RecycleServerWanIPListRequest) {
	request = &RecycleServerWanIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "RecycleServerWanIPList")
	return
}

func NewRecycleServerWanIPListResponse() (response *RecycleServerWanIPListResponse) {
	response = &RecycleServerWanIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回收服务器外网IP
func (c *Client) RecycleServerWanIPList(request *RecycleServerWanIPListRequest) (response *RecycleServerWanIPListResponse, err error) {
	if request == nil {
		request = NewRecycleServerWanIPListRequest()
	}
	response = NewRecycleServerWanIPListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMultiNicDefinitionRequest() (request *DeleteMultiNicDefinitionRequest) {
	request = &DeleteMultiNicDefinitionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteMultiNicDefinition")
	return
}

func NewDeleteMultiNicDefinitionResponse() (response *DeleteMultiNicDefinitionResponse) {
	response = &DeleteMultiNicDefinitionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DeleteMultiNicDefinition 删除网卡信息
func (c *Client) DeleteMultiNicDefinition(request *DeleteMultiNicDefinitionRequest) (response *DeleteMultiNicDefinitionResponse, err error) {
	if request == nil {
		request = NewDeleteMultiNicDefinitionRequest()
	}
	response = NewDeleteMultiNicDefinitionResponse()
	err = c.Send(request, response)
	return
}

func NewListServiceBindingsRequest() (request *ListServiceBindingsRequest) {
	request = &ListServiceBindingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListServiceBindings")
	return
}

func NewListServiceBindingsResponse() (response *ListServiceBindingsResponse) {
	response = &ListServiceBindingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务绑定列表
func (c *Client) ListServiceBindings(request *ListServiceBindingsRequest) (response *ListServiceBindingsResponse, err error) {
	if request == nil {
		request = NewListServiceBindingsRequest()
	}
	response = NewListServiceBindingsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskFieldsEnumExRequest() (request *DescribeDiskFieldsEnumExRequest) {
	request = &DescribeDiskFieldsEnumExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeDiskFieldsEnumEx")
	return
}

func NewDescribeDiskFieldsEnumExResponse() (response *DescribeDiskFieldsEnumExResponse) {
	response = &DescribeDiskFieldsEnumExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取硬盘各字段枚举值
func (c *Client) DescribeDiskFieldsEnumEx(request *DescribeDiskFieldsEnumExRequest) (response *DescribeDiskFieldsEnumExResponse, err error) {
	if request == nil {
		request = NewDescribeDiskFieldsEnumExRequest()
	}
	response = NewDescribeDiskFieldsEnumExResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterNameRequest() (request *ModifyClusterNameRequest) {
	request = &ModifyClusterNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyClusterName")
	return
}

func NewModifyClusterNameResponse() (response *ModifyClusterNameResponse) {
	response = &ModifyClusterNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集群名
func (c *Client) ModifyClusterName(request *ModifyClusterNameRequest) (response *ModifyClusterNameResponse, err error) {
	if request == nil {
		request = NewModifyClusterNameRequest()
	}
	response = NewModifyClusterNameResponse()
	err = c.Send(request, response)
	return
}

func NewListBasePodsRequest() (request *ListBasePodsRequest) {
	request = &ListBasePodsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListBasePods")
	return
}

func NewListBasePodsResponse() (response *ListBasePodsResponse) {
	response = &ListBasePodsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点基本信息
func (c *Client) ListBasePods(request *ListBasePodsRequest) (response *ListBasePodsResponse, err error) {
	if request == nil {
		request = NewListBasePodsRequest()
	}
	response = NewListBasePodsResponse()
	err = c.Send(request, response)
	return
}

func NewVerifyLogConfigRequest() (request *VerifyLogConfigRequest) {
	request = &VerifyLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "VerifyLogConfig")
	return
}

func NewVerifyLogConfigResponse() (response *VerifyLogConfigResponse) {
	response = &VerifyLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验日志接入配置
func (c *Client) VerifyLogConfig(request *VerifyLogConfigRequest) (response *VerifyLogConfigResponse, err error) {
	if request == nil {
		request = NewVerifyLogConfigRequest()
	}
	response = NewVerifyLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteJobsRequest() (request *DeleteJobsRequest) {
	request = &DeleteJobsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteJobs")
	return
}

func NewDeleteJobsResponse() (response *DeleteJobsResponse) {
	response = &DeleteJobsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Job
func (c *Client) DeleteJobs(request *DeleteJobsRequest) (response *DeleteJobsResponse, err error) {
	if request == nil {
		request = NewDeleteJobsRequest()
	}
	response = NewDeleteJobsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogRuleGroupTemplesRequest() (request *DeleteLogRuleGroupTemplesRequest) {
	request = &DeleteLogRuleGroupTemplesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteLogRuleGroupTemples")
	return
}

func NewDeleteLogRuleGroupTemplesResponse() (response *DeleteLogRuleGroupTemplesResponse) {
	response = &DeleteLogRuleGroupTemplesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除日志告警策略模板
func (c *Client) DeleteLogRuleGroupTemples(request *DeleteLogRuleGroupTemplesRequest) (response *DeleteLogRuleGroupTemplesResponse, err error) {
	if request == nil {
		request = NewDeleteLogRuleGroupTemplesRequest()
	}
	response = NewDeleteLogRuleGroupTemplesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrUpdateRetentionRequest() (request *CreateOrUpdateRetentionRequest) {
	request = &CreateOrUpdateRetentionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateOrUpdateRetention")
	return
}

func NewCreateOrUpdateRetentionResponse() (response *CreateOrUpdateRetentionResponse) {
	response = &CreateOrUpdateRetentionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建或更新Retention
func (c *Client) CreateOrUpdateRetention(request *CreateOrUpdateRetentionRequest) (response *CreateOrUpdateRetentionResponse, err error) {
	if request == nil {
		request = NewCreateOrUpdateRetentionRequest()
	}
	response = NewCreateOrUpdateRetentionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeploymentsStatsRequest() (request *DescribeDeploymentsStatsRequest) {
	request = &DescribeDeploymentsStatsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeDeploymentsStats")
	return
}

func NewDescribeDeploymentsStatsResponse() (response *DescribeDeploymentsStatsResponse) {
	response = &DescribeDeploymentsStatsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Deployment统计
func (c *Client) DescribeDeploymentsStats(request *DescribeDeploymentsStatsRequest) (response *DescribeDeploymentsStatsResponse, err error) {
	if request == nil {
		request = NewDescribeDeploymentsStatsRequest()
	}
	response = NewDescribeDeploymentsStatsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteStorageClassRequest() (request *DeleteStorageClassRequest) {
	request = &DeleteStorageClassRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteStorageClass")
	return
}

func NewDeleteStorageClassResponse() (response *DeleteStorageClassResponse) {
	response = &DeleteStorageClassResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除StorageClass
func (c *Client) DeleteStorageClass(request *DeleteStorageClassRequest) (response *DeleteStorageClassResponse, err error) {
	if request == nil {
		request = NewDeleteStorageClassRequest()
	}
	response = NewDeleteStorageClassResponse()
	err = c.Send(request, response)
	return
}

func NewGetImageFieldsEnumRequest() (request *GetImageFieldsEnumRequest) {
	request = &GetImageFieldsEnumRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetImageFieldsEnum")
	return
}

func NewGetImageFieldsEnumResponse() (response *GetImageFieldsEnumResponse) {
	response = &GetImageFieldsEnumResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取镜像管理字段枚举值
func (c *Client) GetImageFieldsEnum(request *GetImageFieldsEnumRequest) (response *GetImageFieldsEnumResponse, err error) {
	if request == nil {
		request = NewGetImageFieldsEnumRequest()
	}
	response = NewGetImageFieldsEnumResponse()
	err = c.Send(request, response)
	return
}

func NewListHPAsRequest() (request *ListHPAsRequest) {
	request = &ListHPAsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListHPAs")
	return
}

func NewListHPAsResponse() (response *ListHPAsResponse) {
	response = &ListHPAsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 指定workload查询HPA列表
func (c *Client) ListHPAs(request *ListHPAsRequest) (response *ListHPAsResponse, err error) {
	if request == nil {
		request = NewListHPAsRequest()
	}
	response = NewListHPAsResponse()
	err = c.Send(request, response)
	return
}

func NewSearchLogRuleGroupsRequest() (request *SearchLogRuleGroupsRequest) {
	request = &SearchLogRuleGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SearchLogRuleGroups")
	return
}

func NewSearchLogRuleGroupsResponse() (response *SearchLogRuleGroupsResponse) {
	response = &SearchLogRuleGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件查询日志告警策略
func (c *Client) SearchLogRuleGroups(request *SearchLogRuleGroupsRequest) (response *SearchLogRuleGroupsResponse, err error) {
	if request == nil {
		request = NewSearchLogRuleGroupsRequest()
	}
	response = NewSearchLogRuleGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMetaMetricRequest() (request *DeleteMetaMetricRequest) {
	request = &DeleteMetaMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteMetaMetric")
	return
}

func NewDeleteMetaMetricResponse() (response *DeleteMetaMetricResponse) {
	response = &DeleteMetaMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除指标
func (c *Client) DeleteMetaMetric(request *DeleteMetaMetricRequest) (response *DeleteMetaMetricResponse, err error) {
	if request == nil {
		request = NewDeleteMetaMetricRequest()
	}
	response = NewDeleteMetaMetricResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePodsRequest() (request *DescribePodsRequest) {
	request = &DescribePodsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribePods")
	return
}

func NewDescribePodsResponse() (response *DescribePodsResponse) {
	response = &DescribePodsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云哨自监控pod异常列表和统计
func (c *Client) DescribePods(request *DescribePodsRequest) (response *DescribePodsResponse, err error) {
	if request == nil {
		request = NewDescribePodsRequest()
	}
	response = NewDescribePodsResponse()
	err = c.Send(request, response)
	return
}

func NewApplyWsTokenRequest() (request *ApplyWsTokenRequest) {
	request = &ApplyWsTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ApplyWsToken")
	return
}

func NewApplyWsTokenResponse() (response *ApplyWsTokenResponse) {
	response = &ApplyWsTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通用申请 Websocket Token
func (c *Client) ApplyWsToken(request *ApplyWsTokenRequest) (response *ApplyWsTokenResponse, err error) {
	if request == nil {
		request = NewApplyWsTokenRequest()
	}
	response = NewApplyWsTokenResponse()
	err = c.Send(request, response)
	return
}

func NewRollbackWorkloadRequest() (request *RollbackWorkloadRequest) {
	request = &RollbackWorkloadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "RollbackWorkload")
	return
}

func NewRollbackWorkloadResponse() (response *RollbackWorkloadResponse) {
	response = &RollbackWorkloadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回滚workload至某个历史版本
func (c *Client) RollbackWorkload(request *RollbackWorkloadRequest) (response *RollbackWorkloadResponse, err error) {
	if request == nil {
		request = NewRollbackWorkloadRequest()
	}
	response = NewRollbackWorkloadResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIgniterImagesRequest() (request *DescribeIgniterImagesRequest) {
	request = &DescribeIgniterImagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeIgniterImages")
	return
}

func NewDescribeIgniterImagesResponse() (response *DescribeIgniterImagesResponse) {
	response = &DescribeIgniterImagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Igniter装机镜像列表
func (c *Client) DescribeIgniterImages(request *DescribeIgniterImagesRequest) (response *DescribeIgniterImagesResponse, err error) {
	if request == nil {
		request = NewDescribeIgniterImagesRequest()
	}
	response = NewDescribeIgniterImagesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOutbandInfoExRequest() (request *DescribeOutbandInfoExRequest) {
	request = &DescribeOutbandInfoExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeOutbandInfoEx")
	return
}

func NewDescribeOutbandInfoExResponse() (response *DescribeOutbandInfoExResponse) {
	response = &DescribeOutbandInfoExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 带外管理,账号密码查询
func (c *Client) DescribeOutbandInfoEx(request *DescribeOutbandInfoExRequest) (response *DescribeOutbandInfoExResponse, err error) {
	if request == nil {
		request = NewDescribeOutbandInfoExRequest()
	}
	response = NewDescribeOutbandInfoExResponse()
	err = c.Send(request, response)
	return
}

func NewSearchMetaValuesRequest() (request *SearchMetaValuesRequest) {
	request = &SearchMetaValuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SearchMetaValues")
	return
}

func NewSearchMetaValuesResponse() (response *SearchMetaValuesResponse) {
	response = &SearchMetaValuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 模糊查询MetaValue
func (c *Client) SearchMetaValues(request *SearchMetaValuesRequest) (response *SearchMetaValuesResponse, err error) {
	if request == nil {
		request = NewSearchMetaValuesRequest()
	}
	response = NewSearchMetaValuesResponse()
	err = c.Send(request, response)
	return
}

func NewGetAlertStatsRequest() (request *GetAlertStatsRequest) {
	request = &GetAlertStatsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetAlertStats")
	return
}

func NewGetAlertStatsResponse() (response *GetAlertStatsResponse) {
	response = &GetAlertStatsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取报警概述
func (c *Client) GetAlertStats(request *GetAlertStatsRequest) (response *GetAlertStatsResponse, err error) {
	if request == nil {
		request = NewGetAlertStatsRequest()
	}
	response = NewGetAlertStatsResponse()
	err = c.Send(request, response)
	return
}

func NewListJobsRequest() (request *ListJobsRequest) {
	request = &ListJobsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListJobs")
	return
}

func NewListJobsResponse() (response *ListJobsResponse) {
	response = &ListJobsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Job列表
func (c *Client) ListJobs(request *ListJobsRequest) (response *ListJobsResponse, err error) {
	if request == nil {
		request = NewListJobsRequest()
	}
	response = NewListJobsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIngressStatRequest() (request *DescribeIngressStatRequest) {
	request = &DescribeIngressStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeIngressStat")
	return
}

func NewDescribeIngressStatResponse() (response *DescribeIngressStatResponse) {
	response = &DescribeIngressStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询路由控制器中路由统计信息
func (c *Client) DescribeIngressStat(request *DescribeIngressStatRequest) (response *DescribeIngressStatResponse, err error) {
	if request == nil {
		request = NewDescribeIngressStatRequest()
	}
	response = NewDescribeIngressStatResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePhysvrsListExRequest() (request *DescribePhysvrsListExRequest) {
	request = &DescribePhysvrsListExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribePhysvrsListEx")
	return
}

func NewDescribePhysvrsListExResponse() (response *DescribePhysvrsListExResponse) {
	response = &DescribePhysvrsListExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取物理服务器列表
func (c *Client) DescribePhysvrsListEx(request *DescribePhysvrsListExRequest) (response *DescribePhysvrsListExResponse, err error) {
	if request == nil {
		request = NewDescribePhysvrsListExRequest()
	}
	response = NewDescribePhysvrsListExResponse()
	err = c.Send(request, response)
	return
}

func NewListEventsRequest() (request *ListEventsRequest) {
	request = &ListEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListEvents")
	return
}

func NewListEventsResponse() (response *ListEventsResponse) {
	response = &ListEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据workload查询事件列表
func (c *Client) ListEvents(request *ListEventsRequest) (response *ListEventsResponse, err error) {
	if request == nil {
		request = NewListEventsRequest()
	}
	response = NewListEventsResponse()
	err = c.Send(request, response)
	return
}

func NewProxyAllDataSourceRequestByNameRequest() (request *ProxyAllDataSourceRequestByNameRequest) {
	request = &ProxyAllDataSourceRequestByNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ProxyAllDataSourceRequestByName")
	return
}

func NewProxyAllDataSourceRequestByNameResponse() (response *ProxyAllDataSourceRequestByNameResponse) {
	response = &ProxyAllDataSourceRequestByNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据名称自动代理监控请求
func (c *Client) ProxyAllDataSourceRequestByName(request *ProxyAllDataSourceRequestByNameRequest) (response *ProxyAllDataSourceRequestByNameResponse, err error) {
	if request == nil {
		request = NewProxyAllDataSourceRequestByNameRequest()
	}
	response = NewProxyAllDataSourceRequestByNameResponse()
	err = c.Send(request, response)
	return
}

func NewGetUpgradeProgressRequest() (request *GetUpgradeProgressRequest) {
	request = &GetUpgradeProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetUpgradeProgress")
	return
}

func NewGetUpgradeProgressResponse() (response *GetUpgradeProgressResponse) {
	response = &GetUpgradeProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取升级/回滚进度
func (c *Client) GetUpgradeProgress(request *GetUpgradeProgressRequest) (response *GetUpgradeProgressResponse, err error) {
	if request == nil {
		request = NewGetUpgradeProgressRequest()
	}
	response = NewGetUpgradeProgressResponse()
	err = c.Send(request, response)
	return
}

func NewListPodsRequest() (request *ListPodsRequest) {
	request = &ListPodsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListPods")
	return
}

func NewListPodsResponse() (response *ListPodsResponse) {
	response = &ListPodsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Pod信息
func (c *Client) ListPods(request *ListPodsRequest) (response *ListPodsResponse, err error) {
	if request == nil {
		request = NewListPodsRequest()
	}
	response = NewListPodsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyImageBuildingTaskRequest() (request *ModifyImageBuildingTaskRequest) {
	request = &ModifyImageBuildingTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyImageBuildingTask")
	return
}

func NewModifyImageBuildingTaskResponse() (response *ModifyImageBuildingTaskResponse) {
	response = &ModifyImageBuildingTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新镜像构建任务
func (c *Client) ModifyImageBuildingTask(request *ModifyImageBuildingTaskRequest) (response *ModifyImageBuildingTaskResponse, err error) {
	if request == nil {
		request = NewModifyImageBuildingTaskRequest()
	}
	response = NewModifyImageBuildingTaskResponse()
	err = c.Send(request, response)
	return
}

func NewGetUpgradeCheckResultRequest() (request *GetUpgradeCheckResultRequest) {
	request = &GetUpgradeCheckResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetUpgradeCheckResult")
	return
}

func NewGetUpgradeCheckResultResponse() (response *GetUpgradeCheckResultResponse) {
	response = &GetUpgradeCheckResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取升级前检查结果, 包含节点运行/系统资源使用率是否超过阈值/节点资源是否满足高可用
func (c *Client) GetUpgradeCheckResult(request *GetUpgradeCheckResultRequest) (response *GetUpgradeCheckResultResponse, err error) {
	if request == nil {
		request = NewGetUpgradeCheckResultRequest()
	}
	response = NewGetUpgradeCheckResultResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateNodeLabelsRequest() (request *UpdateNodeLabelsRequest) {
	request = &UpdateNodeLabelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "UpdateNodeLabels")
	return
}

func NewUpdateNodeLabelsResponse() (response *UpdateNodeLabelsResponse) {
	response = &UpdateNodeLabelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新某个结点的标签信息
func (c *Client) UpdateNodeLabels(request *UpdateNodeLabelsRequest) (response *UpdateNodeLabelsResponse, err error) {
	if request == nil {
		request = NewUpdateNodeLabelsRequest()
	}
	response = NewUpdateNodeLabelsResponse()
	err = c.Send(request, response)
	return
}

func NewGenTimeFormatRequest() (request *GenTimeFormatRequest) {
	request = &GenTimeFormatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GenTimeFormat")
	return
}

func NewGenTimeFormatResponse() (response *GenTimeFormatResponse) {
	response = &GenTimeFormatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 自动生成时间格式
func (c *Client) GenTimeFormat(request *GenTimeFormatRequest) (response *GenTimeFormatResponse, err error) {
	if request == nil {
		request = NewGenTimeFormatRequest()
	}
	response = NewGenTimeFormatResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetdeviceRelatedServersExRequest() (request *DescribeNetdeviceRelatedServersExRequest) {
	request = &DescribeNetdeviceRelatedServersExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeNetdeviceRelatedServersEx")
	return
}

func NewDescribeNetdeviceRelatedServersExResponse() (response *DescribeNetdeviceRelatedServersExResponse) {
	response = &DescribeNetdeviceRelatedServersExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取网络设备关联的服务器
func (c *Client) DescribeNetdeviceRelatedServersEx(request *DescribeNetdeviceRelatedServersExRequest) (response *DescribeNetdeviceRelatedServersExResponse, err error) {
	if request == nil {
		request = NewDescribeNetdeviceRelatedServersExRequest()
	}
	response = NewDescribeNetdeviceRelatedServersExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOSListRequest() (request *DescribeOSListRequest) {
	request = &DescribeOSListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeOSList")
	return
}

func NewDescribeOSListResponse() (response *DescribeOSListResponse) {
	response = &DescribeOSListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeOSList
func (c *Client) DescribeOSList(request *DescribeOSListRequest) (response *DescribeOSListResponse, err error) {
	if request == nil {
		request = NewDescribeOSListRequest()
	}
	response = NewDescribeOSListResponse()
	err = c.Send(request, response)
	return
}

func NewCheckMachineRequest() (request *CheckMachineRequest) {
	request = &CheckMachineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CheckMachine")
	return
}

func NewCheckMachineResponse() (response *CheckMachineResponse) {
	response = &CheckMachineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在线检查机器是否符合添加Node
func (c *Client) CheckMachine(request *CheckMachineRequest) (response *CheckMachineResponse, err error) {
	if request == nil {
		request = NewCheckMachineRequest()
	}
	response = NewCheckMachineResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrUpdateDashboardRequest() (request *CreateOrUpdateDashboardRequest) {
	request = &CreateOrUpdateDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateOrUpdateDashboard")
	return
}

func NewCreateOrUpdateDashboardResponse() (response *CreateOrUpdateDashboardResponse) {
	response = &CreateOrUpdateDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建或更新仪表盘
func (c *Client) CreateOrUpdateDashboard(request *CreateOrUpdateDashboardRequest) (response *CreateOrUpdateDashboardResponse, err error) {
	if request == nil {
		request = NewCreateOrUpdateDashboardRequest()
	}
	response = NewCreateOrUpdateDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogConfigRequest() (request *DeleteLogConfigRequest) {
	request = &DeleteLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteLogConfig")
	return
}

func NewDeleteLogConfigResponse() (response *DeleteLogConfigResponse) {
	response = &DeleteLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除日志接入配置
func (c *Client) DeleteLogConfig(request *DeleteLogConfigRequest) (response *DeleteLogConfigResponse, err error) {
	if request == nil {
		request = NewDeleteLogConfigRequest()
	}
	response = NewDeleteLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMetaResourceTypeRequest() (request *DeleteMetaResourceTypeRequest) {
	request = &DeleteMetaResourceTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteMetaResourceType")
	return
}

func NewDeleteMetaResourceTypeResponse() (response *DeleteMetaResourceTypeResponse) {
	response = &DeleteMetaResourceTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除资源对象类型对象
func (c *Client) DeleteMetaResourceType(request *DeleteMetaResourceTypeRequest) (response *DeleteMetaResourceTypeResponse, err error) {
	if request == nil {
		request = NewDeleteMetaResourceTypeRequest()
	}
	response = NewDeleteMetaResourceTypeResponse()
	err = c.Send(request, response)
	return
}

func NewRelocateServerCancelRequest() (request *RelocateServerCancelRequest) {
	request = &RelocateServerCancelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "RelocateServerCancel")
	return
}

func NewRelocateServerCancelResponse() (response *RelocateServerCancelResponse) {
	response = &RelocateServerCancelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器取消搬迁接口
func (c *Client) RelocateServerCancel(request *RelocateServerCancelRequest) (response *RelocateServerCancelResponse, err error) {
	if request == nil {
		request = NewRelocateServerCancelRequest()
	}
	response = NewRelocateServerCancelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInitMachineScriptRequest() (request *DescribeInitMachineScriptRequest) {
	request = &DescribeInitMachineScriptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeInitMachineScript")
	return
}

func NewDescribeInitMachineScriptResponse() (response *DescribeInitMachineScriptResponse) {
	response = &DescribeInitMachineScriptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询机器初始化脚本
func (c *Client) DescribeInitMachineScript(request *DescribeInitMachineScriptRequest) (response *DescribeInitMachineScriptResponse, err error) {
	if request == nil {
		request = NewDescribeInitMachineScriptRequest()
	}
	response = NewDescribeInitMachineScriptResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOutbandStrategyRequest() (request *DeleteOutbandStrategyRequest) {
	request = &DeleteOutbandStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteOutbandStrategy")
	return
}

func NewDeleteOutbandStrategyResponse() (response *DeleteOutbandStrategyResponse) {
	response = &DeleteOutbandStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除自定义带外密码策略
func (c *Client) DeleteOutbandStrategy(request *DeleteOutbandStrategyRequest) (response *DeleteOutbandStrategyResponse, err error) {
	if request == nil {
		request = NewDeleteOutbandStrategyRequest()
	}
	response = NewDeleteOutbandStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteQueryRecordRequest() (request *DeleteQueryRecordRequest) {
	request = &DeleteQueryRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteQueryRecord")
	return
}

func NewDeleteQueryRecordResponse() (response *DeleteQueryRecordResponse) {
	response = &DeleteQueryRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除查询记录
func (c *Client) DeleteQueryRecord(request *DeleteQueryRecordRequest) (response *DeleteQueryRecordResponse, err error) {
	if request == nil {
		request = NewDeleteQueryRecordRequest()
	}
	response = NewDeleteQueryRecordResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSilenceRequest() (request *UpdateSilenceRequest) {
	request = &UpdateSilenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "UpdateSilence")
	return
}

func NewUpdateSilenceResponse() (response *UpdateSilenceResponse) {
	response = &UpdateSilenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新屏蔽策略
func (c *Client) UpdateSilence(request *UpdateSilenceRequest) (response *UpdateSilenceResponse, err error) {
	if request == nil {
		request = NewUpdateSilenceRequest()
	}
	response = NewUpdateSilenceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteListenerRequest() (request *DeleteListenerRequest) {
	request = &DeleteListenerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteListener")
	return
}

func NewDeleteListenerResponse() (response *DeleteListenerResponse) {
	response = &DeleteListenerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除负载均衡监听器
func (c *Client) DeleteListener(request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
	if request == nil {
		request = NewDeleteListenerRequest()
	}
	response = NewDeleteListenerResponse()
	err = c.Send(request, response)
	return
}

func NewModifyServiceRequest() (request *ModifyServiceRequest) {
	request = &ModifyServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyService")
	return
}

func NewModifyServiceResponse() (response *ModifyServiceResponse) {
	response = &ModifyServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改服务
func (c *Client) ModifyService(request *ModifyServiceRequest) (response *ModifyServiceResponse, err error) {
	if request == nil {
		request = NewModifyServiceRequest()
	}
	response = NewModifyServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteIngressClassRequest() (request *DeleteIngressClassRequest) {
	request = &DeleteIngressClassRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteIngressClass")
	return
}

func NewDeleteIngressClassResponse() (response *DeleteIngressClassResponse) {
	response = &DeleteIngressClassResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除路由控制器
func (c *Client) DeleteIngressClass(request *DeleteIngressClassRequest) (response *DeleteIngressClassResponse, err error) {
	if request == nil {
		request = NewDeleteIngressClassRequest()
	}
	response = NewDeleteIngressClassResponse()
	err = c.Send(request, response)
	return
}

func NewRestartWorkloadRequest() (request *RestartWorkloadRequest) {
	request = &RestartWorkloadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "RestartWorkload")
	return
}

func NewRestartWorkloadResponse() (response *RestartWorkloadResponse) {
	response = &RestartWorkloadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重新部署工作负载，只支持Deployment, StatefulSet, DaemonSet
func (c *Client) RestartWorkload(request *RestartWorkloadRequest) (response *RestartWorkloadResponse, err error) {
	if request == nil {
		request = NewRestartWorkloadRequest()
	}
	response = NewRestartWorkloadResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNodeRequest() (request *CreateNodeRequest) {
	request = &CreateNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateNode")
	return
}

func NewCreateNodeResponse() (response *CreateNodeResponse) {
	response = &CreateNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 指定物理机IP、账号、密码等信息，将其添加到Kubernetes集群中
func (c *Client) CreateNode(request *CreateNodeRequest) (response *CreateNodeResponse, err error) {
	if request == nil {
		request = NewCreateNodeRequest()
	}
	response = NewCreateNodeResponse()
	err = c.Send(request, response)
	return
}

func NewGenRegexRequest() (request *GenRegexRequest) {
	request = &GenRegexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GenRegex")
	return
}

func NewGenRegexResponse() (response *GenRegexResponse) {
	response = &GenRegexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 自动生成正则表达式
func (c *Client) GenRegex(request *GenRegexRequest) (response *GenRegexResponse, err error) {
	if request == nil {
		request = NewGenRegexRequest()
	}
	response = NewGenRegexResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogConfigRequest() (request *ModifyLogConfigRequest) {
	request = &ModifyLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyLogConfig")
	return
}

func NewModifyLogConfigResponse() (response *ModifyLogConfigResponse) {
	response = &ModifyLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志接入配置变更
func (c *Client) ModifyLogConfig(request *ModifyLogConfigRequest) (response *ModifyLogConfigResponse, err error) {
	if request == nil {
		request = NewModifyLogConfigRequest()
	}
	response = NewModifyLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewRenameServersRequest() (request *RenameServersRequest) {
	request = &RenameServersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "RenameServers")
	return
}

func NewRenameServersResponse() (response *RenameServersResponse) {
	response = &RenameServersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重命名服务器
func (c *Client) RenameServers(request *RenameServersRequest) (response *RenameServersResponse, err error) {
	if request == nil {
		request = NewRenameServersRequest()
	}
	response = NewRenameServersResponse()
	err = c.Send(request, response)
	return
}

func NewStarredDashboardRequest() (request *StarredDashboardRequest) {
	request = &StarredDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "StarredDashboard")
	return
}

func NewStarredDashboardResponse() (response *StarredDashboardResponse) {
	response = &StarredDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 收藏dashboard
func (c *Client) StarredDashboard(request *StarredDashboardRequest) (response *StarredDashboardResponse, err error) {
	if request == nil {
		request = NewStarredDashboardRequest()
	}
	response = NewStarredDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNodeAlertsRequest() (request *DescribeNodeAlertsRequest) {
	request = &DescribeNodeAlertsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeNodeAlerts")
	return
}

func NewDescribeNodeAlertsResponse() (response *DescribeNodeAlertsResponse) {
	response = &DescribeNodeAlertsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点告警
func (c *Client) DescribeNodeAlerts(request *DescribeNodeAlertsRequest) (response *DescribeNodeAlertsResponse, err error) {
	if request == nil {
		request = NewDescribeNodeAlertsRequest()
	}
	response = NewDescribeNodeAlertsResponse()
	err = c.Send(request, response)
	return
}

func NewServerSpecialRecycleLanIPExRequest() (request *ServerSpecialRecycleLanIPExRequest) {
	request = &ServerSpecialRecycleLanIPExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ServerSpecialRecycleLanIPEx")
	return
}

func NewServerSpecialRecycleLanIPExResponse() (response *ServerSpecialRecycleLanIPExResponse) {
	response = &ServerSpecialRecycleLanIPExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器批量回收相应虚拟内网段IP接口
func (c *Client) ServerSpecialRecycleLanIPEx(request *ServerSpecialRecycleLanIPExRequest) (response *ServerSpecialRecycleLanIPExResponse, err error) {
	if request == nil {
		request = NewServerSpecialRecycleLanIPExRequest()
	}
	response = NewServerSpecialRecycleLanIPExResponse()
	err = c.Send(request, response)
	return
}

func NewListStatefulSetsRequest() (request *ListStatefulSetsRequest) {
	request = &ListStatefulSetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListStatefulSets")
	return
}

func NewListStatefulSetsResponse() (response *ListStatefulSetsResponse) {
	response = &ListStatefulSetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询有状态应用列表
func (c *Client) ListStatefulSets(request *ListStatefulSetsRequest) (response *ListStatefulSetsResponse, err error) {
	if request == nil {
		request = NewListStatefulSetsRequest()
	}
	response = NewListStatefulSetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterRequest() (request *DescribeClusterRequest) {
	request = &DescribeClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeCluster")
	return
}

func NewDescribeClusterResponse() (response *DescribeClusterResponse) {
	response = &DescribeClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群信息
func (c *Client) DescribeCluster(request *DescribeClusterRequest) (response *DescribeClusterResponse, err error) {
	if request == nil {
		request = NewDescribeClusterRequest()
	}
	response = NewDescribeClusterResponse()
	err = c.Send(request, response)
	return
}

func NewSearchLogsRequest() (request *SearchLogsRequest) {
	request = &SearchLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SearchLogs")
	return
}

func NewSearchLogsResponse() (response *SearchLogsResponse) {
	response = &SearchLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志查询接口
func (c *Client) SearchLogs(request *SearchLogsRequest) (response *SearchLogsResponse, err error) {
	if request == nil {
		request = NewSearchLogsRequest()
	}
	response = NewSearchLogsResponse()
	err = c.Send(request, response)
	return
}

func NewGetUpgradeHistoryDetailRequest() (request *GetUpgradeHistoryDetailRequest) {
	request = &GetUpgradeHistoryDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetUpgradeHistoryDetail")
	return
}

func NewGetUpgradeHistoryDetailResponse() (response *GetUpgradeHistoryDetailResponse) {
	response = &GetUpgradeHistoryDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取升级历史详情
func (c *Client) GetUpgradeHistoryDetail(request *GetUpgradeHistoryDetailRequest) (response *GetUpgradeHistoryDetailResponse, err error) {
	if request == nil {
		request = NewGetUpgradeHistoryDetailRequest()
	}
	response = NewGetUpgradeHistoryDetailResponse()
	err = c.Send(request, response)
	return
}

func NewListImageBuildingTasksRequest() (request *ListImageBuildingTasksRequest) {
	request = &ListImageBuildingTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListImageBuildingTasks")
	return
}

func NewListImageBuildingTasksResponse() (response *ListImageBuildingTasksResponse) {
	response = &ListImageBuildingTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取镜像构建任务列表
func (c *Client) ListImageBuildingTasks(request *ListImageBuildingTasksRequest) (response *ListImageBuildingTasksResponse, err error) {
	if request == nil {
		request = NewListImageBuildingTasksRequest()
	}
	response = NewListImageBuildingTasksResponse()
	err = c.Send(request, response)
	return
}

func NewSaveUpgradeConfigRequest() (request *SaveUpgradeConfigRequest) {
	request = &SaveUpgradeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SaveUpgradeConfig")
	return
}

func NewSaveUpgradeConfigResponse() (response *SaveUpgradeConfigResponse) {
	response = &SaveUpgradeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存集群升级配置信息
func (c *Client) SaveUpgradeConfig(request *SaveUpgradeConfigRequest) (response *SaveUpgradeConfigResponse, err error) {
	if request == nil {
		request = NewSaveUpgradeConfigRequest()
	}
	response = NewSaveUpgradeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMonitorTemplateRequest() (request *DeleteMonitorTemplateRequest) {
	request = &DeleteMonitorTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteMonitorTemplate")
	return
}

func NewDeleteMonitorTemplateResponse() (response *DeleteMonitorTemplateResponse) {
	response = &DeleteMonitorTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除监控自定义模板
func (c *Client) DeleteMonitorTemplate(request *DeleteMonitorTemplateRequest) (response *DeleteMonitorTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteMonitorTemplateRequest()
	}
	response = NewDeleteMonitorTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewSearchProductInfoRequest() (request *SearchProductInfoRequest) {
	request = &SearchProductInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SearchProductInfo")
	return
}

func NewSearchProductInfoResponse() (response *SearchProductInfoResponse) {
	response = &SearchProductInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品信息
func (c *Client) SearchProductInfo(request *SearchProductInfoRequest) (response *SearchProductInfoResponse, err error) {
	if request == nil {
		request = NewSearchProductInfoRequest()
	}
	response = NewSearchProductInfoResponse()
	err = c.Send(request, response)
	return
}

func NewRelocateServerFinishRequest() (request *RelocateServerFinishRequest) {
	request = &RelocateServerFinishRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "RelocateServerFinish")
	return
}

func NewRelocateServerFinishResponse() (response *RelocateServerFinishResponse) {
	response = &RelocateServerFinishResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器搬迁完成接口
func (c *Client) RelocateServerFinish(request *RelocateServerFinishRequest) (response *RelocateServerFinishResponse, err error) {
	if request == nil {
		request = NewRelocateServerFinishRequest()
	}
	response = NewRelocateServerFinishResponse()
	err = c.Send(request, response)
	return
}

func NewGetDashboardVersionContentByIdRequest() (request *GetDashboardVersionContentByIdRequest) {
	request = &GetDashboardVersionContentByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetDashboardVersionContentById")
	return
}

func NewGetDashboardVersionContentByIdResponse() (response *GetDashboardVersionContentByIdResponse) {
	response = &GetDashboardVersionContentByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看给定仪表盘历史版本的内容
func (c *Client) GetDashboardVersionContentById(request *GetDashboardVersionContentByIdRequest) (response *GetDashboardVersionContentByIdResponse, err error) {
	if request == nil {
		request = NewGetDashboardVersionContentByIdRequest()
	}
	response = NewGetDashboardVersionContentByIdResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTargetWeightRequest() (request *ModifyTargetWeightRequest) {
	request = &ModifyTargetWeightRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyTargetWeight")
	return
}

func NewModifyTargetWeightResponse() (response *ModifyTargetWeightResponse) {
	response = &ModifyTargetWeightResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改监听器绑定的后端机器的转发权重
func (c *Client) ModifyTargetWeight(request *ModifyTargetWeightRequest) (response *ModifyTargetWeightResponse, err error) {
	if request == nil {
		request = NewModifyTargetWeightRequest()
	}
	response = NewModifyTargetWeightResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDaemonSetRequest() (request *DescribeDaemonSetRequest) {
	request = &DescribeDaemonSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeDaemonSet")
	return
}

func NewDescribeDaemonSetResponse() (response *DescribeDaemonSetResponse) {
	response = &DescribeDaemonSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询DaemonSet
func (c *Client) DescribeDaemonSet(request *DescribeDaemonSetRequest) (response *DescribeDaemonSetResponse, err error) {
	if request == nil {
		request = NewDescribeDaemonSetRequest()
	}
	response = NewDescribeDaemonSetResponse()
	err = c.Send(request, response)
	return
}

func NewDriveDetailByClusterMetricRequest() (request *DriveDetailByClusterMetricRequest) {
	request = &DriveDetailByClusterMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DriveDetailByClusterMetric")
	return
}

func NewDriveDetailByClusterMetricResponse() (response *DriveDetailByClusterMetricResponse) {
	response = &DriveDetailByClusterMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群下钻接口
func (c *Client) DriveDetailByClusterMetric(request *DriveDetailByClusterMetricRequest) (response *DriveDetailByClusterMetricResponse, err error) {
	if request == nil {
		request = NewDriveDetailByClusterMetricRequest()
	}
	response = NewDriveDetailByClusterMetricResponse()
	err = c.Send(request, response)
	return
}

func NewGetUpgradeHistoryListRequest() (request *GetUpgradeHistoryListRequest) {
	request = &GetUpgradeHistoryListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetUpgradeHistoryList")
	return
}

func NewGetUpgradeHistoryListResponse() (response *GetUpgradeHistoryListResponse) {
	response = &GetUpgradeHistoryListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取升级历史列表
func (c *Client) GetUpgradeHistoryList(request *GetUpgradeHistoryListRequest) (response *GetUpgradeHistoryListResponse, err error) {
	if request == nil {
		request = NewGetUpgradeHistoryListRequest()
	}
	response = NewGetUpgradeHistoryListResponse()
	err = c.Send(request, response)
	return
}

func NewSearchRoutesRequest() (request *SearchRoutesRequest) {
	request = &SearchRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SearchRoutes")
	return
}

func NewSearchRoutesResponse() (response *SearchRoutesResponse) {
	response = &SearchRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件查询路由
func (c *Client) SearchRoutes(request *SearchRoutesRequest) (response *SearchRoutesResponse, err error) {
	if request == nil {
		request = NewSearchRoutesRequest()
	}
	response = NewSearchRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewSearchAlertsRequest() (request *SearchAlertsRequest) {
	request = &SearchAlertsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SearchAlerts")
	return
}

func NewSearchAlertsResponse() (response *SearchAlertsResponse) {
	response = &SearchAlertsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件搜索Alert
// 分页page、limit、asc/desc
func (c *Client) SearchAlerts(request *SearchAlertsRequest) (response *SearchAlertsResponse, err error) {
	if request == nil {
		request = NewSearchAlertsRequest()
	}
	response = NewSearchAlertsResponse()
	err = c.Send(request, response)
	return
}

func NewAddRAIDRequest() (request *AddRAIDRequest) {
	request = &AddRAIDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "AddRAID")
	return
}

func NewAddRAIDResponse() (response *AddRAIDResponse) {
	response = &AddRAIDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加RAID
func (c *Client) AddRAID(request *AddRAIDRequest) (response *AddRAIDResponse, err error) {
	if request == nil {
		request = NewAddRAIDRequest()
	}
	response = NewAddRAIDResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBatchDashboardByIdRequest() (request *DeleteBatchDashboardByIdRequest) {
	request = &DeleteBatchDashboardByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteBatchDashboardById")
	return
}

func NewDeleteBatchDashboardByIdResponse() (response *DeleteBatchDashboardByIdResponse) {
	response = &DeleteBatchDashboardByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除dashboard
func (c *Client) DeleteBatchDashboardById(request *DeleteBatchDashboardByIdRequest) (response *DeleteBatchDashboardByIdResponse, err error) {
	if request == nil {
		request = NewDeleteBatchDashboardByIdRequest()
	}
	response = NewDeleteBatchDashboardByIdResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDevOpsScenesByClusterRequest() (request *CreateDevOpsScenesByClusterRequest) {
	request = &CreateDevOpsScenesByClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateDevOpsScenesByCluster")
	return
}

func NewCreateDevOpsScenesByClusterResponse() (response *CreateDevOpsScenesByClusterResponse) {
	response = &CreateDevOpsScenesByClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理自愈场景
func (c *Client) CreateDevOpsScenesByCluster(request *CreateDevOpsScenesByClusterRequest) (response *CreateDevOpsScenesByClusterResponse, err error) {
	if request == nil {
		request = NewCreateDevOpsScenesByClusterRequest()
	}
	response = NewCreateDevOpsScenesByClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLoadBalancerRequest() (request *DeleteLoadBalancerRequest) {
	request = &DeleteLoadBalancerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteLoadBalancer")
	return
}

func NewDeleteLoadBalancerResponse() (response *DeleteLoadBalancerResponse) {
	response = &DeleteLoadBalancerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除负载均衡实例
func (c *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (response *DeleteLoadBalancerResponse, err error) {
	if request == nil {
		request = NewDeleteLoadBalancerRequest()
	}
	response = NewDeleteLoadBalancerResponse()
	err = c.Send(request, response)
	return
}

func NewSearchRoutesByLabelSetRequest() (request *SearchRoutesByLabelSetRequest) {
	request = &SearchRoutesByLabelSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SearchRoutesByLabelSet")
	return
}

func NewSearchRoutesByLabelSetResponse() (response *SearchRoutesByLabelSetResponse) {
	response = &SearchRoutesByLabelSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// SearchRoutesByLabelSet
func (c *Client) SearchRoutesByLabelSet(request *SearchRoutesByLabelSetRequest) (response *SearchRoutesByLabelSetResponse, err error) {
	if request == nil {
		request = NewSearchRoutesByLabelSetRequest()
	}
	response = NewSearchRoutesByLabelSetResponse()
	err = c.Send(request, response)
	return
}

func NewGetDashboardPermissionsByIdRequest() (request *GetDashboardPermissionsByIdRequest) {
	request = &GetDashboardPermissionsByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetDashboardPermissionsById")
	return
}

func NewGetDashboardPermissionsByIdResponse() (response *GetDashboardPermissionsByIdResponse) {
	response = &GetDashboardPermissionsByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取仪表盘权限
func (c *Client) GetDashboardPermissionsById(request *GetDashboardPermissionsByIdRequest) (response *GetDashboardPermissionsByIdResponse, err error) {
	if request == nil {
		request = NewGetDashboardPermissionsByIdRequest()
	}
	response = NewGetDashboardPermissionsByIdResponse()
	err = c.Send(request, response)
	return
}

func NewListMonitorTemplatesRequest() (request *ListMonitorTemplatesRequest) {
	request = &ListMonitorTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListMonitorTemplates")
	return
}

func NewListMonitorTemplatesResponse() (response *ListMonitorTemplatesResponse) {
	response = &ListMonitorTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取监控自定义模板列表
func (c *Client) ListMonitorTemplates(request *ListMonitorTemplatesRequest) (response *ListMonitorTemplatesResponse, err error) {
	if request == nil {
		request = NewListMonitorTemplatesRequest()
	}
	response = NewListMonitorTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDashFolderByIdRequest() (request *DeleteDashFolderByIdRequest) {
	request = &DeleteDashFolderByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteDashFolderById")
	return
}

func NewDeleteDashFolderByIdResponse() (response *DeleteDashFolderByIdResponse) {
	response = &DeleteDashFolderByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除仪表盘目录
func (c *Client) DeleteDashFolderById(request *DeleteDashFolderByIdRequest) (response *DeleteDashFolderByIdResponse, err error) {
	if request == nil {
		request = NewDeleteDashFolderByIdRequest()
	}
	response = NewDeleteDashFolderByIdResponse()
	err = c.Send(request, response)
	return
}

func NewGetMetricInstantByDimensionRequest() (request *GetMetricInstantByDimensionRequest) {
	request = &GetMetricInstantByDimensionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetMetricInstantByDimension")
	return
}

func NewGetMetricInstantByDimensionResponse() (response *GetMetricInstantByDimensionResponse) {
	response = &GetMetricInstantByDimensionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetMetricInstantByDimension
func (c *Client) GetMetricInstantByDimension(request *GetMetricInstantByDimensionRequest) (response *GetMetricInstantByDimensionResponse, err error) {
	if request == nil {
		request = NewGetMetricInstantByDimensionRequest()
	}
	response = NewGetMetricInstantByDimensionResponse()
	err = c.Send(request, response)
	return
}

func NewListBaseClustersRequest() (request *ListBaseClustersRequest) {
	request = &ListBaseClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListBaseClusters")
	return
}

func NewListBaseClustersResponse() (response *ListBaseClustersResponse) {
	response = &ListBaseClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群基本信息
func (c *Client) ListBaseClusters(request *ListBaseClustersRequest) (response *ListBaseClustersResponse, err error) {
	if request == nil {
		request = NewListBaseClustersRequest()
	}
	response = NewListBaseClustersResponse()
	err = c.Send(request, response)
	return
}

func NewListWorkloadPodsRequest() (request *ListWorkloadPodsRequest) {
	request = &ListWorkloadPodsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListWorkloadPods")
	return
}

func NewListWorkloadPodsResponse() (response *ListWorkloadPodsResponse) {
	response = &ListWorkloadPodsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询工作负载的Pod列表
func (c *Client) ListWorkloadPods(request *ListWorkloadPodsRequest) (response *ListWorkloadPodsResponse, err error) {
	if request == nil {
		request = NewListWorkloadPodsRequest()
	}
	response = NewListWorkloadPodsResponse()
	err = c.Send(request, response)
	return
}

func NewGetMetricHistoryByDimensionRequest() (request *GetMetricHistoryByDimensionRequest) {
	request = &GetMetricHistoryByDimensionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetMetricHistoryByDimension")
	return
}

func NewGetMetricHistoryByDimensionResponse() (response *GetMetricHistoryByDimensionResponse) {
	response = &GetMetricHistoryByDimensionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取各个维度列表展示的指标性能数据（和前端相关）
func (c *Client) GetMetricHistoryByDimension(request *GetMetricHistoryByDimensionRequest) (response *GetMetricHistoryByDimensionResponse, err error) {
	if request == nil {
		request = NewGetMetricHistoryByDimensionRequest()
	}
	response = NewGetMetricHistoryByDimensionResponse()
	err = c.Send(request, response)
	return
}

func NewRelocateServerStartRequest() (request *RelocateServerStartRequest) {
	request = &RelocateServerStartRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "RelocateServerStart")
	return
}

func NewRelocateServerStartResponse() (response *RelocateServerStartResponse) {
	response = &RelocateServerStartResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器搬迁发起接口
func (c *Client) RelocateServerStart(request *RelocateServerStartRequest) (response *RelocateServerStartResponse, err error) {
	if request == nil {
		request = NewRelocateServerStartRequest()
	}
	response = NewRelocateServerStartResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateDashFolderByIdRequest() (request *UpdateDashFolderByIdRequest) {
	request = &UpdateDashFolderByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "UpdateDashFolderById")
	return
}

func NewUpdateDashFolderByIdResponse() (response *UpdateDashFolderByIdResponse) {
	response = &UpdateDashFolderByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新文件夹
func (c *Client) UpdateDashFolderById(request *UpdateDashFolderByIdRequest) (response *UpdateDashFolderByIdResponse, err error) {
	if request == nil {
		request = NewUpdateDashFolderByIdRequest()
	}
	response = NewUpdateDashFolderByIdResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSvrNicAllocationRequest() (request *UpdateSvrNicAllocationRequest) {
	request = &UpdateSvrNicAllocationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "UpdateSvrNicAllocation")
	return
}

func NewUpdateSvrNicAllocationResponse() (response *UpdateSvrNicAllocationResponse) {
	response = &UpdateSvrNicAllocationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// UpdateSvrNicAllocation 绑定(更新)服务器多网卡策略
func (c *Client) UpdateSvrNicAllocation(request *UpdateSvrNicAllocationRequest) (response *UpdateSvrNicAllocationResponse, err error) {
	if request == nil {
		request = NewUpdateSvrNicAllocationRequest()
	}
	response = NewUpdateSvrNicAllocationResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrUpdateLogRuleGroupTempleRequest() (request *CreateOrUpdateLogRuleGroupTempleRequest) {
	request = &CreateOrUpdateLogRuleGroupTempleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateOrUpdateLogRuleGroupTemple")
	return
}

func NewCreateOrUpdateLogRuleGroupTempleResponse() (response *CreateOrUpdateLogRuleGroupTempleResponse) {
	response = &CreateOrUpdateLogRuleGroupTempleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建或更新日志报警策略模板
func (c *Client) CreateOrUpdateLogRuleGroupTemple(request *CreateOrUpdateLogRuleGroupTempleRequest) (response *CreateOrUpdateLogRuleGroupTempleResponse, err error) {
	if request == nil {
		request = NewCreateOrUpdateLogRuleGroupTempleRequest()
	}
	response = NewCreateOrUpdateLogRuleGroupTempleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJobRequest() (request *DescribeJobRequest) {
	request = &DescribeJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeJob")
	return
}

func NewDescribeJobResponse() (response *DescribeJobResponse) {
	response = &DescribeJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Job
func (c *Client) DescribeJob(request *DescribeJobRequest) (response *DescribeJobResponse, err error) {
	if request == nil {
		request = NewDescribeJobRequest()
	}
	response = NewDescribeJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductRequest() (request *DescribeProductRequest) {
	request = &DescribeProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeProduct")
	return
}

func NewDescribeProductResponse() (response *DescribeProductResponse) {
	response = &DescribeProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品健康状态
func (c *Client) DescribeProduct(request *DescribeProductRequest) (response *DescribeProductResponse, err error) {
	if request == nil {
		request = NewDescribeProductRequest()
	}
	response = NewDescribeProductResponse()
	err = c.Send(request, response)
	return
}

func NewListHistoryRevisionsRequest() (request *ListHistoryRevisionsRequest) {
	request = &ListHistoryRevisionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListHistoryRevisions")
	return
}

func NewListHistoryRevisionsResponse() (response *ListHistoryRevisionsResponse) {
	response = &ListHistoryRevisionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Workload历史版本
func (c *Client) ListHistoryRevisions(request *ListHistoryRevisionsRequest) (response *ListHistoryRevisionsResponse, err error) {
	if request == nil {
		request = NewListHistoryRevisionsRequest()
	}
	response = NewListHistoryRevisionsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRAIDRequest() (request *ModifyRAIDRequest) {
	request = &ModifyRAIDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyRAID")
	return
}

func NewModifyRAIDResponse() (response *ModifyRAIDResponse) {
	response = &ModifyRAIDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改RAID
func (c *Client) ModifyRAID(request *ModifyRAIDRequest) (response *ModifyRAIDResponse, err error) {
	if request == nil {
		request = NewModifyRAIDRequest()
	}
	response = NewModifyRAIDResponse()
	err = c.Send(request, response)
	return
}

func NewRegisterTargetsRequest() (request *RegisterTargetsRequest) {
	request = &RegisterTargetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "RegisterTargets")
	return
}

func NewRegisterTargetsResponse() (response *RegisterTargetsResponse) {
	response = &RegisterTargetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定后端机器到监听器上
func (c *Client) RegisterTargets(request *RegisterTargetsRequest) (response *RegisterTargetsResponse, err error) {
	if request == nil {
		request = NewRegisterTargetsRequest()
	}
	response = NewRegisterTargetsResponse()
	err = c.Send(request, response)
	return
}

func NewListClusterVersionsRequest() (request *ListClusterVersionsRequest) {
	request = &ListClusterVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListClusterVersions")
	return
}

func NewListClusterVersionsResponse() (response *ListClusterVersionsResponse) {
	response = &ListClusterVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群版本列表
func (c *Client) ListClusterVersions(request *ListClusterVersionsRequest) (response *ListClusterVersionsResponse, err error) {
	if request == nil {
		request = NewListClusterVersionsRequest()
	}
	response = NewListClusterVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewListDevOpsHistoryRequest() (request *ListDevOpsHistoryRequest) {
	request = &ListDevOpsHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListDevOpsHistory")
	return
}

func NewListDevOpsHistoryResponse() (response *ListDevOpsHistoryResponse) {
	response = &ListDevOpsHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取自愈历史
func (c *Client) ListDevOpsHistory(request *ListDevOpsHistoryRequest) (response *ListDevOpsHistoryResponse, err error) {
	if request == nil {
		request = NewListDevOpsHistoryRequest()
	}
	response = NewListDevOpsHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewListNodeRemediationTemplatesRequest() (request *ListNodeRemediationTemplatesRequest) {
	request = &ListNodeRemediationTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListNodeRemediationTemplates")
	return
}

func NewListNodeRemediationTemplatesResponse() (response *ListNodeRemediationTemplatesResponse) {
	response = &ListNodeRemediationTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 治愈场景模板列表，包括自愈动作
func (c *Client) ListNodeRemediationTemplates(request *ListNodeRemediationTemplatesRequest) (response *ListNodeRemediationTemplatesResponse, err error) {
	if request == nil {
		request = NewListNodeRemediationTemplatesRequest()
	}
	response = NewListNodeRemediationTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNodeRequest() (request *DeleteNodeRequest) {
	request = &DeleteNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteNode")
	return
}

func NewDeleteNodeResponse() (response *DeleteNodeResponse) {
	response = &DeleteNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从kubernetes集群中移除结点。注意结点上的数据、配置文件、已安装的包不会自动清理。
func (c *Client) DeleteNode(request *DeleteNodeRequest) (response *DeleteNodeResponse, err error) {
	if request == nil {
		request = NewDeleteNodeRequest()
	}
	response = NewDeleteNodeResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateRouteRequest() (request *UpdateRouteRequest) {
	request = &UpdateRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "UpdateRoute")
	return
}

func NewUpdateRouteResponse() (response *UpdateRouteResponse) {
	response = &UpdateRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新指定路由
func (c *Client) UpdateRoute(request *UpdateRouteRequest) (response *UpdateRouteResponse, err error) {
	if request == nil {
		request = NewUpdateRouteRequest()
	}
	response = NewUpdateRouteResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
	request = &CreateClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateCluster")
	return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
	response = &CreateClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集群
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
	if request == nil {
		request = NewCreateClusterRequest()
	}
	response = NewCreateClusterResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrUpdateRuleGroupTempleRequest() (request *CreateOrUpdateRuleGroupTempleRequest) {
	request = &CreateOrUpdateRuleGroupTempleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateOrUpdateRuleGroupTemple")
	return
}

func NewCreateOrUpdateRuleGroupTempleResponse() (response *CreateOrUpdateRuleGroupTempleResponse) {
	response = &CreateOrUpdateRuleGroupTempleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建或更新策略模板
func (c *Client) CreateOrUpdateRuleGroupTemple(request *CreateOrUpdateRuleGroupTempleRequest) (response *CreateOrUpdateRuleGroupTempleResponse, err error) {
	if request == nil {
		request = NewCreateOrUpdateRuleGroupTempleRequest()
	}
	response = NewCreateOrUpdateRuleGroupTempleResponse()
	err = c.Send(request, response)
	return
}

func NewGetNotificationRequest() (request *GetNotificationRequest) {
	request = &GetNotificationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetNotification")
	return
}

func NewGetNotificationResponse() (response *GetNotificationResponse) {
	response = &GetNotificationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定notification
func (c *Client) GetNotification(request *GetNotificationRequest) (response *GetNotificationResponse, err error) {
	if request == nil {
		request = NewGetNotificationRequest()
	}
	response = NewGetNotificationResponse()
	err = c.Send(request, response)
	return
}

func NewRecycleServerVirtualIPRequest() (request *RecycleServerVirtualIPRequest) {
	request = &RecycleServerVirtualIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "RecycleServerVirtualIP")
	return
}

func NewRecycleServerVirtualIPResponse() (response *RecycleServerVirtualIPResponse) {
	response = &RecycleServerVirtualIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回收物理服务器虚拟内外网ip
func (c *Client) RecycleServerVirtualIP(request *RecycleServerVirtualIPRequest) (response *RecycleServerVirtualIPResponse, err error) {
	if request == nil {
		request = NewRecycleServerVirtualIPRequest()
	}
	response = NewRecycleServerVirtualIPResponse()
	err = c.Send(request, response)
	return
}

func NewAddResourceObjectRequest() (request *AddResourceObjectRequest) {
	request = &AddResourceObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "AddResourceObject")
	return
}

func NewAddResourceObjectResponse() (response *AddResourceObjectResponse) {
	response = &AddResourceObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建服务树资源对象
func (c *Client) AddResourceObject(request *AddResourceObjectRequest) (response *AddResourceObjectResponse, err error) {
	if request == nil {
		request = NewAddResourceObjectRequest()
	}
	response = NewAddResourceObjectResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePodsOverviewWithAlertsRequest() (request *DescribePodsOverviewWithAlertsRequest) {
	request = &DescribePodsOverviewWithAlertsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribePodsOverviewWithAlerts")
	return
}

func NewDescribePodsOverviewWithAlertsResponse() (response *DescribePodsOverviewWithAlertsResponse) {
	response = &DescribePodsOverviewWithAlertsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云哨概览页获取pod信息
func (c *Client) DescribePodsOverviewWithAlerts(request *DescribePodsOverviewWithAlertsRequest) (response *DescribePodsOverviewWithAlertsResponse, err error) {
	if request == nil {
		request = NewDescribePodsOverviewWithAlertsRequest()
	}
	response = NewDescribePodsOverviewWithAlertsResponse()
	err = c.Send(request, response)
	return
}

func NewListWorkloadsRequest() (request *ListWorkloadsRequest) {
	request = &ListWorkloadsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListWorkloads")
	return
}

func NewListWorkloadsResponse() (response *ListWorkloadsResponse) {
	response = &ListWorkloadsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询工作负载列表
func (c *Client) ListWorkloads(request *ListWorkloadsRequest) (response *ListWorkloadsResponse, err error) {
	if request == nil {
		request = NewListWorkloadsRequest()
	}
	response = NewListWorkloadsResponse()
	err = c.Send(request, response)
	return
}

func NewAddVMListRequest() (request *AddVMListRequest) {
	request = &AddVMListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "AddVMList")
	return
}

func NewAddVMListResponse() (response *AddVMListResponse) {
	response = &AddVMListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加虚拟机
func (c *Client) AddVMList(request *AddVMListRequest) (response *AddVMListResponse, err error) {
	if request == nil {
		request = NewAddVMListRequest()
	}
	response = NewAddVMListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeListenersRequest() (request *DescribeListenersRequest) {
	request = &DescribeListenersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeListeners")
	return
}

func NewDescribeListenersResponse() (response *DescribeListenersResponse) {
	response = &DescribeListenersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询负载均衡的监听器列表
func (c *Client) DescribeListeners(request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
	if request == nil {
		request = NewDescribeListenersRequest()
	}
	response = NewDescribeListenersResponse()
	err = c.Send(request, response)
	return
}

func NewListCreateVolumeParametersRequest() (request *ListCreateVolumeParametersRequest) {
	request = &ListCreateVolumeParametersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListCreateVolumeParameters")
	return
}

func NewListCreateVolumeParametersResponse() (response *ListCreateVolumeParametersResponse) {
	response = &ListCreateVolumeParametersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 指定某个StorageClass, 列出创建Volume时所需的参数
func (c *Client) ListCreateVolumeParameters(request *ListCreateVolumeParametersRequest) (response *ListCreateVolumeParametersResponse, err error) {
	if request == nil {
		request = NewListCreateVolumeParametersRequest()
	}
	response = NewListCreateVolumeParametersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirtualMachineYamlRequest() (request *DescribeVirtualMachineYamlRequest) {
	request = &DescribeVirtualMachineYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeVirtualMachineYaml")
	return
}

func NewDescribeVirtualMachineYamlResponse() (response *DescribeVirtualMachineYamlResponse) {
	response = &DescribeVirtualMachineYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取虚拟机YAML
func (c *Client) DescribeVirtualMachineYaml(request *DescribeVirtualMachineYamlRequest) (response *DescribeVirtualMachineYamlResponse, err error) {
	if request == nil {
		request = NewDescribeVirtualMachineYamlRequest()
	}
	response = NewDescribeVirtualMachineYamlResponse()
	err = c.Send(request, response)
	return
}

func NewListVolumesRequest() (request *ListVolumesRequest) {
	request = &ListVolumesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListVolumes")
	return
}

func NewListVolumesResponse() (response *ListVolumesResponse) {
	response = &ListVolumesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出所有卷
func (c *Client) ListVolumes(request *ListVolumesRequest) (response *ListVolumesResponse, err error) {
	if request == nil {
		request = NewListVolumesRequest()
	}
	response = NewListVolumesResponse()
	err = c.Send(request, response)
	return
}

func NewGetMatrixRequest() (request *GetMatrixRequest) {
	request = &GetMatrixRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetMatrix")
	return
}

func NewGetMatrixResponse() (response *GetMatrixResponse) {
	response = &GetMatrixResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取dashboard图表数据
func (c *Client) GetMatrix(request *GetMatrixRequest) (response *GetMatrixResponse, err error) {
	if request == nil {
		request = NewGetMatrixRequest()
	}
	response = NewGetMatrixResponse()
	err = c.Send(request, response)
	return
}

func NewCreateIngressRequest() (request *CreateIngressRequest) {
	request = &CreateIngressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateIngress")
	return
}

func NewCreateIngressResponse() (response *CreateIngressResponse) {
	response = &CreateIngressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建Ingress
func (c *Client) CreateIngress(request *CreateIngressRequest) (response *CreateIngressResponse, err error) {
	if request == nil {
		request = NewCreateIngressRequest()
	}
	response = NewCreateIngressResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrUpdateLogRuleGroupRequest() (request *CreateOrUpdateLogRuleGroupRequest) {
	request = &CreateOrUpdateLogRuleGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateOrUpdateLogRuleGroup")
	return
}

func NewCreateOrUpdateLogRuleGroupResponse() (response *CreateOrUpdateLogRuleGroupResponse) {
	response = &CreateOrUpdateLogRuleGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建日志报警策略
func (c *Client) CreateOrUpdateLogRuleGroup(request *CreateOrUpdateLogRuleGroupRequest) (response *CreateOrUpdateLogRuleGroupResponse, err error) {
	if request == nil {
		request = NewCreateOrUpdateLogRuleGroupRequest()
	}
	response = NewCreateOrUpdateLogRuleGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeploymentRequest() (request *DescribeDeploymentRequest) {
	request = &DescribeDeploymentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeDeployment")
	return
}

func NewDescribeDeploymentResponse() (response *DescribeDeploymentResponse) {
	response = &DescribeDeploymentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询部署详情
func (c *Client) DescribeDeployment(request *DescribeDeploymentRequest) (response *DescribeDeploymentResponse, err error) {
	if request == nil {
		request = NewDescribeDeploymentRequest()
	}
	response = NewDescribeDeploymentResponse()
	err = c.Send(request, response)
	return
}

func NewCleanerScriptRequest() (request *CleanerScriptRequest) {
	request = &CleanerScriptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CleanerScript")
	return
}

func NewCleanerScriptResponse() (response *CleanerScriptResponse) {
	response = &CleanerScriptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除节点后，需要通过此接口获取手动清理宿主机上残留的数据、配置、安装包的脚本
func (c *Client) CleanerScript(request *CleanerScriptRequest) (response *CleanerScriptResponse, err error) {
	if request == nil {
		request = NewCleanerScriptRequest()
	}
	response = NewCleanerScriptResponse()
	err = c.Send(request, response)
	return
}

func NewGetSeriesRequest() (request *GetSeriesRequest) {
	request = &GetSeriesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetSeries")
	return
}

func NewGetSeriesResponse() (response *GetSeriesResponse) {
	response = &GetSeriesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取时间线
func (c *Client) GetSeries(request *GetSeriesRequest) (response *GetSeriesResponse, err error) {
	if request == nil {
		request = NewGetSeriesRequest()
	}
	response = NewGetSeriesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirtualMachineStatusRequest() (request *DescribeVirtualMachineStatusRequest) {
	request = &DescribeVirtualMachineStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeVirtualMachineStatus")
	return
}

func NewDescribeVirtualMachineStatusResponse() (response *DescribeVirtualMachineStatusResponse) {
	response = &DescribeVirtualMachineStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 虚拟机运行状态（资源使用信息）
func (c *Client) DescribeVirtualMachineStatus(request *DescribeVirtualMachineStatusRequest) (response *DescribeVirtualMachineStatusResponse, err error) {
	if request == nil {
		request = NewDescribeVirtualMachineStatusRequest()
	}
	response = NewDescribeVirtualMachineStatusResponse()
	err = c.Send(request, response)
	return
}

func NewRelocateServerQueryRequest() (request *RelocateServerQueryRequest) {
	request = &RelocateServerQueryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "RelocateServerQuery")
	return
}

func NewRelocateServerQueryResponse() (response *RelocateServerQueryResponse) {
	response = &RelocateServerQueryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器搬迁状态查询接口
func (c *Client) RelocateServerQuery(request *RelocateServerQueryRequest) (response *RelocateServerQueryResponse, err error) {
	if request == nil {
		request = NewRelocateServerQueryRequest()
	}
	response = NewRelocateServerQueryResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateResourceObjectRequest() (request *UpdateResourceObjectRequest) {
	request = &UpdateResourceObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "UpdateResourceObject")
	return
}

func NewUpdateResourceObjectResponse() (response *UpdateResourceObjectResponse) {
	response = &UpdateResourceObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新服务树节点资源对象
func (c *Client) UpdateResourceObject(request *UpdateResourceObjectRequest) (response *UpdateResourceObjectResponse, err error) {
	if request == nil {
		request = NewUpdateResourceObjectRequest()
	}
	response = NewUpdateResourceObjectResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNetworkPolicyRequest() (request *CreateNetworkPolicyRequest) {
	request = &CreateNetworkPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateNetworkPolicy")
	return
}

func NewCreateNetworkPolicyResponse() (response *CreateNetworkPolicyResponse) {
	response = &CreateNetworkPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建网络策略
func (c *Client) CreateNetworkPolicy(request *CreateNetworkPolicyRequest) (response *CreateNetworkPolicyResponse, err error) {
	if request == nil {
		request = NewCreateNetworkPolicyRequest()
	}
	response = NewCreateNetworkPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewGetMetaRequest() (request *GetMetaRequest) {
	request = &GetMetaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetMeta")
	return
}

func NewGetMetaResponse() (response *GetMetaResponse) {
	response = &GetMetaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Meta树
func (c *Client) GetMeta(request *GetMetaRequest) (response *GetMetaResponse, err error) {
	if request == nil {
		request = NewGetMetaRequest()
	}
	response = NewGetMetaResponse()
	err = c.Send(request, response)
	return
}

func NewGetPublicKeyRequest() (request *GetPublicKeyRequest) {
	request = &GetPublicKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetPublicKey")
	return
}

func NewGetPublicKeyResponse() (response *GetPublicKeyResponse) {
	response = &GetPublicKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取公钥, 用户加密密码传给后端
func (c *Client) GetPublicKey(request *GetPublicKeyRequest) (response *GetPublicKeyResponse, err error) {
	if request == nil {
		request = NewGetPublicKeyRequest()
	}
	response = NewGetPublicKeyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDeploymentRequest() (request *ModifyDeploymentRequest) {
	request = &ModifyDeploymentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyDeployment")
	return
}

func NewModifyDeploymentResponse() (response *ModifyDeploymentResponse) {
	response = &ModifyDeploymentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改部署Deployment
func (c *Client) ModifyDeployment(request *ModifyDeploymentRequest) (response *ModifyDeploymentResponse, err error) {
	if request == nil {
		request = NewModifyDeploymentRequest()
	}
	response = NewModifyDeploymentResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDaemonSetsRequest() (request *DeleteDaemonSetsRequest) {
	request = &DeleteDaemonSetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteDaemonSets")
	return
}

func NewDeleteDaemonSetsResponse() (response *DeleteDaemonSetsResponse) {
	response = &DeleteDaemonSetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除DaemonSet
func (c *Client) DeleteDaemonSets(request *DeleteDaemonSetsRequest) (response *DeleteDaemonSetsResponse, err error) {
	if request == nil {
		request = NewDeleteDaemonSetsRequest()
	}
	response = NewDeleteDaemonSetsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyServersRequest() (request *ModifyServersRequest) {
	request = &ModifyServersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyServers")
	return
}

func NewModifyServersResponse() (response *ModifyServersResponse) {
	response = &ModifyServersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新服务器信息请求信息
func (c *Client) ModifyServers(request *ModifyServersRequest) (response *ModifyServersResponse, err error) {
	if request == nil {
		request = NewModifyServersRequest()
	}
	response = NewModifyServersResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCronJobsRequest() (request *DeleteCronJobsRequest) {
	request = &DeleteCronJobsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteCronJobs")
	return
}

func NewDeleteCronJobsResponse() (response *DeleteCronJobsResponse) {
	response = &DeleteCronJobsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除CronJob
func (c *Client) DeleteCronJobs(request *DeleteCronJobsRequest) (response *DeleteCronJobsResponse, err error) {
	if request == nil {
		request = NewDeleteCronJobsRequest()
	}
	response = NewDeleteCronJobsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceManagementOverviewRequest() (request *DescribeServiceManagementOverviewRequest) {
	request = &DescribeServiceManagementOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeServiceManagementOverview")
	return
}

func NewDescribeServiceManagementOverviewResponse() (response *DescribeServiceManagementOverviewResponse) {
	response = &DescribeServiceManagementOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询支撑服务管控概览
func (c *Client) DescribeServiceManagementOverview(request *DescribeServiceManagementOverviewRequest) (response *DescribeServiceManagementOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeServiceManagementOverviewRequest()
	}
	response = NewDescribeServiceManagementOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCustomScriptSetRequest() (request *CreateCustomScriptSetRequest) {
	request = &CreateCustomScriptSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateCustomScriptSet")
	return
}

func NewCreateCustomScriptSetResponse() (response *CreateCustomScriptSetResponse) {
	response = &CreateCustomScriptSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建自定义脚本集
func (c *Client) CreateCustomScriptSet(request *CreateCustomScriptSetRequest) (response *CreateCustomScriptSetResponse, err error) {
	if request == nil {
		request = NewCreateCustomScriptSetRequest()
	}
	response = NewCreateCustomScriptSetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePhysvrsOverviewRequest() (request *DescribePhysvrsOverviewRequest) {
	request = &DescribePhysvrsOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribePhysvrsOverview")
	return
}

func NewDescribePhysvrsOverviewResponse() (response *DescribePhysvrsOverviewResponse) {
	response = &DescribePhysvrsOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器资源概览
func (c *Client) DescribePhysvrsOverview(request *DescribePhysvrsOverviewRequest) (response *DescribePhysvrsOverviewResponse, err error) {
	if request == nil {
		request = NewDescribePhysvrsOverviewRequest()
	}
	response = NewDescribePhysvrsOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewCheckClusterDeletableRequest() (request *CheckClusterDeletableRequest) {
	request = &CheckClusterDeletableRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CheckClusterDeletable")
	return
}

func NewCheckClusterDeletableResponse() (response *CheckClusterDeletableResponse) {
	response = &CheckClusterDeletableResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查是否业务集群，是否可删
func (c *Client) CheckClusterDeletable(request *CheckClusterDeletableRequest) (response *CheckClusterDeletableResponse, err error) {
	if request == nil {
		request = NewCheckClusterDeletableRequest()
	}
	response = NewCheckClusterDeletableResponse()
	err = c.Send(request, response)
	return
}

func NewImportClusterRequest() (request *ImportClusterRequest) {
	request = &ImportClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ImportCluster")
	return
}

func NewImportClusterResponse() (response *ImportClusterResponse) {
	response = &ImportClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入集群
func (c *Client) ImportCluster(request *ImportClusterRequest) (response *ImportClusterResponse, err error) {
	if request == nil {
		request = NewImportClusterRequest()
	}
	response = NewImportClusterResponse()
	err = c.Send(request, response)
	return
}

func NewGetDashboardContentByUidRequest() (request *GetDashboardContentByUidRequest) {
	request = &GetDashboardContentByUidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetDashboardContentByUid")
	return
}

func NewGetDashboardContentByUidResponse() (response *GetDashboardContentByUidResponse) {
	response = &GetDashboardContentByUidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取仪表盘详细内容
func (c *Client) GetDashboardContentByUid(request *GetDashboardContentByUidRequest) (response *GetDashboardContentByUidResponse, err error) {
	if request == nil {
		request = NewGetDashboardContentByUidRequest()
	}
	response = NewGetDashboardContentByUidResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceTopologyRequest() (request *DescribeServiceTopologyRequest) {
	request = &DescribeServiceTopologyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeServiceTopology")
	return
}

func NewDescribeServiceTopologyResponse() (response *DescribeServiceTopologyResponse) {
	response = &DescribeServiceTopologyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务拓扑
func (c *Client) DescribeServiceTopology(request *DescribeServiceTopologyRequest) (response *DescribeServiceTopologyResponse, err error) {
	if request == nil {
		request = NewDescribeServiceTopologyRequest()
	}
	response = NewDescribeServiceTopologyResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateServerWanIPListRequest() (request *AllocateServerWanIPListRequest) {
	request = &AllocateServerWanIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "AllocateServerWanIPList")
	return
}

func NewAllocateServerWanIPListResponse() (response *AllocateServerWanIPListResponse) {
	response = &AllocateServerWanIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分配服务器外网IP
func (c *Client) AllocateServerWanIPList(request *AllocateServerWanIPListRequest) (response *AllocateServerWanIPListResponse, err error) {
	if request == nil {
		request = NewAllocateServerWanIPListRequest()
	}
	response = NewAllocateServerWanIPListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRouteRequest() (request *CreateRouteRequest) {
	request = &CreateRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateRoute")
	return
}

func NewCreateRouteResponse() (response *CreateRouteResponse) {
	response = &CreateRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建告警路由
func (c *Client) CreateRoute(request *CreateRouteRequest) (response *CreateRouteResponse, err error) {
	if request == nil {
		request = NewCreateRouteRequest()
	}
	response = NewCreateRouteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopPodsByMetricRequest() (request *DescribeTopPodsByMetricRequest) {
	request = &DescribeTopPodsByMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeTopPodsByMetric")
	return
}

func NewDescribeTopPodsByMetricResponse() (response *DescribeTopPodsByMetricResponse) {
	response = &DescribeTopPodsByMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeTopPodsByMetric
func (c *Client) DescribeTopPodsByMetric(request *DescribeTopPodsByMetricRequest) (response *DescribeTopPodsByMetricResponse, err error) {
	if request == nil {
		request = NewDescribeTopPodsByMetricRequest()
	}
	response = NewDescribeTopPodsByMetricResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTopologyStateRequest() (request *ModifyTopologyStateRequest) {
	request = &ModifyTopologyStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyTopologyState")
	return
}

func NewModifyTopologyStateResponse() (response *ModifyTopologyStateResponse) {
	response = &ModifyTopologyStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改服务拓扑状态
func (c *Client) ModifyTopologyState(request *ModifyTopologyStateRequest) (response *ModifyTopologyStateResponse, err error) {
	if request == nil {
		request = NewModifyTopologyStateRequest()
	}
	response = NewModifyTopologyStateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCustomScriptRequest() (request *CreateCustomScriptRequest) {
	request = &CreateCustomScriptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateCustomScript")
	return
}

func NewCreateCustomScriptResponse() (response *CreateCustomScriptResponse) {
	response = &CreateCustomScriptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建用户自定义脚本，用于DCOS装机后置脚本的执行。
func (c *Client) CreateCustomScript(request *CreateCustomScriptRequest) (response *CreateCustomScriptResponse, err error) {
	if request == nil {
		request = NewCreateCustomScriptRequest()
	}
	response = NewCreateCustomScriptResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDeploymentRequest() (request *DeleteDeploymentRequest) {
	request = &DeleteDeploymentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteDeployment")
	return
}

func NewDeleteDeploymentResponse() (response *DeleteDeploymentResponse) {
	response = &DeleteDeploymentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除部署
func (c *Client) DeleteDeployment(request *DeleteDeploymentRequest) (response *DeleteDeploymentResponse, err error) {
	if request == nil {
		request = NewDeleteDeploymentRequest()
	}
	response = NewDeleteDeploymentResponse()
	err = c.Send(request, response)
	return
}

func NewListRegionsRequest() (request *ListRegionsRequest) {
	request = &ListRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListRegions")
	return
}

func NewListRegionsResponse() (response *ListRegionsResponse) {
	response = &ListRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出所有Region
func (c *Client) ListRegions(request *ListRegionsRequest) (response *ListRegionsResponse, err error) {
	if request == nil {
		request = NewListRegionsRequest()
	}
	response = NewListRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyStatefulSetRequest() (request *ModifyStatefulSetRequest) {
	request = &ModifyStatefulSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyStatefulSet")
	return
}

func NewModifyStatefulSetResponse() (response *ModifyStatefulSetResponse) {
	response = &ModifyStatefulSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改有状态应用
func (c *Client) ModifyStatefulSet(request *ModifyStatefulSetRequest) (response *ModifyStatefulSetResponse, err error) {
	if request == nil {
		request = NewModifyStatefulSetRequest()
	}
	response = NewModifyStatefulSetResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLoadBalancerListenersRequest() (request *DeleteLoadBalancerListenersRequest) {
	request = &DeleteLoadBalancerListenersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteLoadBalancerListeners")
	return
}

func NewDeleteLoadBalancerListenersResponse() (response *DeleteLoadBalancerListenersResponse) {
	response = &DeleteLoadBalancerListenersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除负载均衡多个监听器
func (c *Client) DeleteLoadBalancerListeners(request *DeleteLoadBalancerListenersRequest) (response *DeleteLoadBalancerListenersResponse, err error) {
	if request == nil {
		request = NewDeleteLoadBalancerListenersRequest()
	}
	response = NewDeleteLoadBalancerListenersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMonitorResourceRequest() (request *DescribeMonitorResourceRequest) {
	request = &DescribeMonitorResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeMonitorResource")
	return
}

func NewDescribeMonitorResourceResponse() (response *DescribeMonitorResourceResponse) {
	response = &DescribeMonitorResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云哨自监控统计
func (c *Client) DescribeMonitorResource(request *DescribeMonitorResourceRequest) (response *DescribeMonitorResourceResponse, err error) {
	if request == nil {
		request = NewDescribeMonitorResourceRequest()
	}
	response = NewDescribeMonitorResourceResponse()
	err = c.Send(request, response)
	return
}

func NewSearchRuleGroupsRequest() (request *SearchRuleGroupsRequest) {
	request = &SearchRuleGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SearchRuleGroups")
	return
}

func NewSearchRuleGroupsResponse() (response *SearchRuleGroupsResponse) {
	response = &SearchRuleGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件查询策略
func (c *Client) SearchRuleGroups(request *SearchRuleGroupsRequest) (response *SearchRuleGroupsResponse, err error) {
	if request == nil {
		request = NewSearchRuleGroupsRequest()
	}
	response = NewSearchRuleGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVolumeRequest() (request *DeleteVolumeRequest) {
	request = &DeleteVolumeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteVolume")
	return
}

func NewDeleteVolumeResponse() (response *DeleteVolumeResponse) {
	response = &DeleteVolumeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Volume
func (c *Client) DeleteVolume(request *DeleteVolumeRequest) (response *DeleteVolumeResponse, err error) {
	if request == nil {
		request = NewDeleteVolumeRequest()
	}
	response = NewDeleteVolumeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageListRequest() (request *DescribeImageListRequest) {
	request = &DescribeImageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeImageList")
	return
}

func NewDescribeImageListResponse() (response *DescribeImageListResponse) {
	response = &DescribeImageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像
func (c *Client) DescribeImageList(request *DescribeImageListRequest) (response *DescribeImageListResponse, err error) {
	if request == nil {
		request = NewDescribeImageListRequest()
	}
	response = NewDescribeImageListResponse()
	err = c.Send(request, response)
	return
}

func NewSearchDashFoldersRequest() (request *SearchDashFoldersRequest) {
	request = &SearchDashFoldersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SearchDashFolders")
	return
}

func NewSearchDashFoldersResponse() (response *SearchDashFoldersResponse) {
	response = &SearchDashFoldersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过名称、ID列表等内容检索仪表盘目录
func (c *Client) SearchDashFolders(request *SearchDashFoldersRequest) (response *SearchDashFoldersResponse, err error) {
	if request == nil {
		request = NewSearchDashFoldersRequest()
	}
	response = NewSearchDashFoldersResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVMListRequest() (request *DeleteVMListRequest) {
	request = &DeleteVMListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteVMList")
	return
}

func NewDeleteVMListResponse() (response *DeleteVMListResponse) {
	response = &DeleteVMListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除虚拟机
func (c *Client) DeleteVMList(request *DeleteVMListRequest) (response *DeleteVMListResponse, err error) {
	if request == nil {
		request = NewDeleteVMListRequest()
	}
	response = NewDeleteVMListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCustomScriptsRequest() (request *DeleteCustomScriptsRequest) {
	request = &DeleteCustomScriptsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteCustomScripts")
	return
}

func NewDeleteCustomScriptsResponse() (response *DeleteCustomScriptsResponse) {
	response = &DeleteCustomScriptsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户自定义脚本
func (c *Client) DeleteCustomScripts(request *DeleteCustomScriptsRequest) (response *DeleteCustomScriptsResponse, err error) {
	if request == nil {
		request = NewDeleteCustomScriptsRequest()
	}
	response = NewDeleteCustomScriptsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeploySubtaskStepsExRequest() (request *DescribeDeploySubtaskStepsExRequest) {
	request = &DescribeDeploySubtaskStepsExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeDeploySubtaskStepsEx")
	return
}

func NewDescribeDeploySubtaskStepsExResponse() (response *DescribeDeploySubtaskStepsExResponse) {
	response = &DescribeDeploySubtaskStepsExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看部署各节点状态
func (c *Client) DescribeDeploySubtaskStepsEx(request *DescribeDeploySubtaskStepsExRequest) (response *DescribeDeploySubtaskStepsExResponse, err error) {
	if request == nil {
		request = NewDescribeDeploySubtaskStepsExRequest()
	}
	response = NewDescribeDeploySubtaskStepsExResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRoutesRequest() (request *DeleteRoutesRequest) {
	request = &DeleteRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteRoutes")
	return
}

func NewDeleteRoutesResponse() (response *DeleteRoutesResponse) {
	response = &DeleteRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除路由
func (c *Client) DeleteRoutes(request *DeleteRoutesRequest) (response *DeleteRoutesResponse, err error) {
	if request == nil {
		request = NewDeleteRoutesRequest()
	}
	response = NewDeleteRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewListMachineVersionsRequest() (request *ListMachineVersionsRequest) {
	request = &ListMachineVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListMachineVersions")
	return
}

func NewListMachineVersionsResponse() (response *ListMachineVersionsResponse) {
	response = &ListMachineVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点版本信息
func (c *Client) ListMachineVersions(request *ListMachineVersionsRequest) (response *ListMachineVersionsResponse, err error) {
	if request == nil {
		request = NewListMachineVersionsRequest()
	}
	response = NewListMachineVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterStatusRequest() (request *DescribeClusterStatusRequest) {
	request = &DescribeClusterStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeClusterStatus")
	return
}

func NewDescribeClusterStatusResponse() (response *DescribeClusterStatusResponse) {
	response = &DescribeClusterStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群运行状态
func (c *Client) DescribeClusterStatus(request *DescribeClusterStatusRequest) (response *DescribeClusterStatusResponse, err error) {
	if request == nil {
		request = NewDescribeClusterStatusRequest()
	}
	response = NewDescribeClusterStatusResponse()
	err = c.Send(request, response)
	return
}

func NewGetDashBoardsRequest() (request *GetDashBoardsRequest) {
	request = &GetDashBoardsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetDashBoards")
	return
}

func NewGetDashBoardsResponse() (response *GetDashBoardsResponse) {
	response = &GetDashBoardsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetDashBoards
func (c *Client) GetDashBoards(request *GetDashBoardsRequest) (response *GetDashBoardsResponse, err error) {
	if request == nil {
		request = NewGetDashBoardsRequest()
	}
	response = NewGetDashBoardsResponse()
	err = c.Send(request, response)
	return
}

func NewGetDashboardTagsRequest() (request *GetDashboardTagsRequest) {
	request = &GetDashboardTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetDashboardTags")
	return
}

func NewGetDashboardTagsResponse() (response *GetDashboardTagsResponse) {
	response = &GetDashboardTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取仪表盘Tag
func (c *Client) GetDashboardTags(request *GetDashboardTagsRequest) (response *GetDashboardTagsResponse, err error) {
	if request == nil {
		request = NewGetDashboardTagsRequest()
	}
	response = NewGetDashboardTagsResponse()
	err = c.Send(request, response)
	return
}

func NewQueryTaskExRequest() (request *QueryTaskExRequest) {
	request = &QueryTaskExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "QueryTaskEx")
	return
}

func NewQueryTaskExResponse() (response *QueryTaskExResponse) {
	response = &QueryTaskExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询dcos装机结果
func (c *Client) QueryTaskEx(request *QueryTaskExRequest) (response *QueryTaskExResponse, err error) {
	if request == nil {
		request = NewQueryTaskExRequest()
	}
	response = NewQueryTaskExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecretRequest() (request *DescribeSecretRequest) {
	request = &DescribeSecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeSecret")
	return
}

func NewDescribeSecretResponse() (response *DescribeSecretResponse) {
	response = &DescribeSecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Secret详情
func (c *Client) DescribeSecret(request *DescribeSecretRequest) (response *DescribeSecretResponse, err error) {
	if request == nil {
		request = NewDescribeSecretRequest()
	}
	response = NewDescribeSecretResponse()
	err = c.Send(request, response)
	return
}

func NewGetRuleGroupTempleRequest() (request *GetRuleGroupTempleRequest) {
	request = &GetRuleGroupTempleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetRuleGroupTemple")
	return
}

func NewGetRuleGroupTempleResponse() (response *GetRuleGroupTempleResponse) {
	response = &GetRuleGroupTempleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定策略模板
func (c *Client) GetRuleGroupTemple(request *GetRuleGroupTempleRequest) (response *GetRuleGroupTempleResponse, err error) {
	if request == nil {
		request = NewGetRuleGroupTempleRequest()
	}
	response = NewGetRuleGroupTempleResponse()
	err = c.Send(request, response)
	return
}

func NewSearchLogRuleGroupTemplesRequest() (request *SearchLogRuleGroupTemplesRequest) {
	request = &SearchLogRuleGroupTemplesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SearchLogRuleGroupTemples")
	return
}

func NewSearchLogRuleGroupTemplesResponse() (response *SearchLogRuleGroupTemplesResponse) {
	response = &SearchLogRuleGroupTemplesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件查询日志告警策略模板
func (c *Client) SearchLogRuleGroupTemples(request *SearchLogRuleGroupTemplesRequest) (response *SearchLogRuleGroupTemplesResponse, err error) {
	if request == nil {
		request = NewSearchLogRuleGroupTemplesRequest()
	}
	response = NewSearchLogRuleGroupTemplesResponse()
	err = c.Send(request, response)
	return
}

func NewServerAllocateLanIPExRequest() (request *ServerAllocateLanIPExRequest) {
	request = &ServerAllocateLanIPExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ServerAllocateLanIPEx")
	return
}

func NewServerAllocateLanIPExResponse() (response *ServerAllocateLanIPExResponse) {
	response = &ServerAllocateLanIPExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器分配内网IP
func (c *Client) ServerAllocateLanIPEx(request *ServerAllocateLanIPExRequest) (response *ServerAllocateLanIPExResponse, err error) {
	if request == nil {
		request = NewServerAllocateLanIPExRequest()
	}
	response = NewServerAllocateLanIPExResponse()
	err = c.Send(request, response)
	return
}

func NewUnStarredDashboardRequest() (request *UnStarredDashboardRequest) {
	request = &UnStarredDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "UnStarredDashboard")
	return
}

func NewUnStarredDashboardResponse() (response *UnStarredDashboardResponse) {
	response = &UnStarredDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消dashboard 收藏
func (c *Client) UnStarredDashboard(request *UnStarredDashboardRequest) (response *UnStarredDashboardResponse, err error) {
	if request == nil {
		request = NewUnStarredDashboardRequest()
	}
	response = NewUnStarredDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewGetUpgradeConfigRequest() (request *GetUpgradeConfigRequest) {
	request = &GetUpgradeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetUpgradeConfig")
	return
}

func NewGetUpgradeConfigResponse() (response *GetUpgradeConfigResponse) {
	response = &GetUpgradeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群升级需要的版本信息
func (c *Client) GetUpgradeConfig(request *GetUpgradeConfigRequest) (response *GetUpgradeConfigResponse, err error) {
	if request == nil {
		request = NewGetUpgradeConfigRequest()
	}
	response = NewGetUpgradeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLabelRequest() (request *ModifyLabelRequest) {
	request = &ModifyLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyLabel")
	return
}

func NewModifyLabelResponse() (response *ModifyLabelResponse) {
	response = &ModifyLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 标签管理--修改标签
func (c *Client) ModifyLabel(request *ModifyLabelRequest) (response *ModifyLabelResponse, err error) {
	if request == nil {
		request = NewModifyLabelRequest()
	}
	response = NewModifyLabelResponse()
	err = c.Send(request, response)
	return
}

func NewGetDashFoldersRequest() (request *GetDashFoldersRequest) {
	request = &GetDashFoldersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetDashFolders")
	return
}

func NewGetDashFoldersResponse() (response *GetDashFoldersResponse) {
	response = &GetDashFoldersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过名称、ID列表等内容检索仪表盘目录
func (c *Client) GetDashFolders(request *GetDashFoldersRequest) (response *GetDashFoldersResponse, err error) {
	if request == nil {
		request = NewGetDashFoldersRequest()
	}
	response = NewGetDashFoldersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateQueryRecordRequest() (request *CreateQueryRecordRequest) {
	request = &CreateQueryRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateQueryRecord")
	return
}

func NewCreateQueryRecordResponse() (response *CreateQueryRecordResponse) {
	response = &CreateQueryRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存日志查询记录
func (c *Client) CreateQueryRecord(request *CreateQueryRecordRequest) (response *CreateQueryRecordResponse, err error) {
	if request == nil {
		request = NewCreateQueryRecordRequest()
	}
	response = NewCreateQueryRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTargetsRequest() (request *DescribeTargetsRequest) {
	request = &DescribeTargetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeTargets")
	return
}

func NewDescribeTargetsResponse() (response *DescribeTargetsResponse) {
	response = &DescribeTargetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询负载均衡绑定的后端服务列表
func (c *Client) DescribeTargets(request *DescribeTargetsRequest) (response *DescribeTargetsResponse, err error) {
	if request == nil {
		request = NewDescribeTargetsRequest()
	}
	response = NewDescribeTargetsResponse()
	err = c.Send(request, response)
	return
}

func NewGetCronJobRequest() (request *GetCronJobRequest) {
	request = &GetCronJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetCronJob")
	return
}

func NewGetCronJobResponse() (response *GetCronJobResponse) {
	response = &GetCronJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询CronJob详情
func (c *Client) GetCronJob(request *GetCronJobRequest) (response *GetCronJobResponse, err error) {
	if request == nil {
		request = NewGetCronJobRequest()
	}
	response = NewGetCronJobResponse()
	err = c.Send(request, response)
	return
}

func NewGetRuleGroupRequest() (request *GetRuleGroupRequest) {
	request = &GetRuleGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetRuleGroup")
	return
}

func NewGetRuleGroupResponse() (response *GetRuleGroupResponse) {
	response = &GetRuleGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定策略
func (c *Client) GetRuleGroup(request *GetRuleGroupRequest) (response *GetRuleGroupResponse, err error) {
	if request == nil {
		request = NewGetRuleGroupRequest()
	}
	response = NewGetRuleGroupResponse()
	err = c.Send(request, response)
	return
}

func NewLogMappingRequest() (request *LogMappingRequest) {
	request = &LogMappingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "LogMapping")
	return
}

func NewLogMappingResponse() (response *LogMappingResponse) {
	response = &LogMappingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志字段映射
func (c *Client) LogMapping(request *LogMappingRequest) (response *LogMappingResponse, err error) {
	if request == nil {
		request = NewLogMappingRequest()
	}
	response = NewLogMappingResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePersistentVolumeClaimRequest() (request *CreatePersistentVolumeClaimRequest) {
	request = &CreatePersistentVolumeClaimRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreatePersistentVolumeClaim")
	return
}

func NewCreatePersistentVolumeClaimResponse() (response *CreatePersistentVolumeClaimResponse) {
	response = &CreatePersistentVolumeClaimResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建PVC
func (c *Client) CreatePersistentVolumeClaim(request *CreatePersistentVolumeClaimRequest) (response *CreatePersistentVolumeClaimResponse, err error) {
	if request == nil {
		request = NewCreatePersistentVolumeClaimRequest()
	}
	response = NewCreatePersistentVolumeClaimResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVolumeRequest() (request *CreateVolumeRequest) {
	request = &CreateVolumeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateVolume")
	return
}

func NewCreateVolumeResponse() (response *CreateVolumeResponse) {
	response = &CreateVolumeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 指定一个StorageClass, 在其中创建卷
func (c *Client) CreateVolume(request *CreateVolumeRequest) (response *CreateVolumeResponse, err error) {
	if request == nil {
		request = NewCreateVolumeRequest()
	}
	response = NewCreateVolumeResponse()
	err = c.Send(request, response)
	return
}

func NewSearchNotificationMsgsRequest() (request *SearchNotificationMsgsRequest) {
	request = &SearchNotificationMsgsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SearchNotificationMsgs")
	return
}

func NewSearchNotificationMsgsResponse() (response *SearchNotificationMsgsResponse) {
	response = &SearchNotificationMsgsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 消息查询接口
func (c *Client) SearchNotificationMsgs(request *SearchNotificationMsgsRequest) (response *SearchNotificationMsgsResponse, err error) {
	if request == nil {
		request = NewSearchNotificationMsgsRequest()
	}
	response = NewSearchNotificationMsgsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomScriptsRequest() (request *DescribeCustomScriptsRequest) {
	request = &DescribeCustomScriptsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeCustomScripts")
	return
}

func NewDescribeCustomScriptsResponse() (response *DescribeCustomScriptsResponse) {
	response = &DescribeCustomScriptsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户自定义脚本
func (c *Client) DescribeCustomScripts(request *DescribeCustomScriptsRequest) (response *DescribeCustomScriptsResponse, err error) {
	if request == nil {
		request = NewDescribeCustomScriptsRequest()
	}
	response = NewDescribeCustomScriptsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNodesStatsRequest() (request *DescribeNodesStatsRequest) {
	request = &DescribeNodesStatsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeNodesStats")
	return
}

func NewDescribeNodesStatsResponse() (response *DescribeNodesStatsResponse) {
	response = &DescribeNodesStatsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点统计数据
func (c *Client) DescribeNodesStats(request *DescribeNodesStatsRequest) (response *DescribeNodesStatsResponse, err error) {
	if request == nil {
		request = NewDescribeNodesStatsRequest()
	}
	response = NewDescribeNodesStatsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNodeAttributesRequest() (request *ModifyNodeAttributesRequest) {
	request = &ModifyNodeAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyNodeAttributes")
	return
}

func NewModifyNodeAttributesResponse() (response *ModifyNodeAttributesResponse) {
	response = &ModifyNodeAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改Node信息
func (c *Client) ModifyNodeAttributes(request *ModifyNodeAttributesRequest) (response *ModifyNodeAttributesResponse, err error) {
	if request == nil {
		request = NewModifyNodeAttributesRequest()
	}
	response = NewModifyNodeAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMonitorTemplateRequest() (request *CreateMonitorTemplateRequest) {
	request = &CreateMonitorTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateMonitorTemplate")
	return
}

func NewCreateMonitorTemplateResponse() (response *CreateMonitorTemplateResponse) {
	response = &CreateMonitorTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建自定义监控模板
func (c *Client) CreateMonitorTemplate(request *CreateMonitorTemplateRequest) (response *CreateMonitorTemplateResponse, err error) {
	if request == nil {
		request = NewCreateMonitorTemplateRequest()
	}
	response = NewCreateMonitorTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewGetIngressYAMLRequest() (request *GetIngressYAMLRequest) {
	request = &GetIngressYAMLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetIngressYAML")
	return
}

func NewGetIngressYAMLResponse() (response *GetIngressYAMLResponse) {
	response = &GetIngressYAMLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Ingress的YAML描述
func (c *Client) GetIngressYAML(request *GetIngressYAMLRequest) (response *GetIngressYAMLResponse, err error) {
	if request == nil {
		request = NewGetIngressYAMLRequest()
	}
	response = NewGetIngressYAMLResponse()
	err = c.Send(request, response)
	return
}

func NewGetNodeInstantMetricDataRequest() (request *GetNodeInstantMetricDataRequest) {
	request = &GetNodeInstantMetricDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetNodeInstantMetricData")
	return
}

func NewGetNodeInstantMetricDataResponse() (response *GetNodeInstantMetricDataResponse) {
	response = &GetNodeInstantMetricDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询node的某个时间点监控数据
func (c *Client) GetNodeInstantMetricData(request *GetNodeInstantMetricDataRequest) (response *GetNodeInstantMetricDataResponse, err error) {
	if request == nil {
		request = NewGetNodeInstantMetricDataRequest()
	}
	response = NewGetNodeInstantMetricDataResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateSvrNicIPRequest() (request *AllocateSvrNicIPRequest) {
	request = &AllocateSvrNicIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "AllocateSvrNicIP")
	return
}

func NewAllocateSvrNicIPResponse() (response *AllocateSvrNicIPResponse) {
	response = &AllocateSvrNicIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// AllocateSvrNicIP 服务器的多网卡ip分配
func (c *Client) AllocateSvrNicIP(request *AllocateSvrNicIPRequest) (response *AllocateSvrNicIPResponse, err error) {
	if request == nil {
		request = NewAllocateSvrNicIPRequest()
	}
	response = NewAllocateSvrNicIPResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopKPodsRequest() (request *DescribeTopKPodsRequest) {
	request = &DescribeTopKPodsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeTopKPods")
	return
}

func NewDescribeTopKPodsResponse() (response *DescribeTopKPodsResponse) {
	response = &DescribeTopKPodsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询TOP K 云哨自监控pod使用
func (c *Client) DescribeTopKPods(request *DescribeTopKPodsRequest) (response *DescribeTopKPodsResponse, err error) {
	if request == nil {
		request = NewDescribeTopKPodsRequest()
	}
	response = NewDescribeTopKPodsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterByAdminRequest() (request *DeleteClusterByAdminRequest) {
	request = &DeleteClusterByAdminRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteClusterByAdmin")
	return
}

func NewDeleteClusterByAdminResponse() (response *DeleteClusterByAdminResponse) {
	response = &DeleteClusterByAdminResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 限管理员删除业务或纳管集群
func (c *Client) DeleteClusterByAdmin(request *DeleteClusterByAdminRequest) (response *DeleteClusterByAdminResponse, err error) {
	if request == nil {
		request = NewDeleteClusterByAdminRequest()
	}
	response = NewDeleteClusterByAdminResponse()
	err = c.Send(request, response)
	return
}

func NewListVolumesNodeViewRequest() (request *ListVolumesNodeViewRequest) {
	request = &ListVolumesNodeViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListVolumesNodeView")
	return
}

func NewListVolumesNodeViewResponse() (response *ListVolumesNodeViewResponse) {
	response = &ListVolumesNodeViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 结点视角列出所有卷
func (c *Client) ListVolumesNodeView(request *ListVolumesNodeViewRequest) (response *ListVolumesNodeViewResponse, err error) {
	if request == nil {
		request = NewListVolumesNodeViewRequest()
	}
	response = NewListVolumesNodeViewResponse()
	err = c.Send(request, response)
	return
}

func NewGetRollbackProgressRequest() (request *GetRollbackProgressRequest) {
	request = &GetRollbackProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetRollbackProgress")
	return
}

func NewGetRollbackProgressResponse() (response *GetRollbackProgressResponse) {
	response = &GetRollbackProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取回滚进度
func (c *Client) GetRollbackProgress(request *GetRollbackProgressRequest) (response *GetRollbackProgressResponse, err error) {
	if request == nil {
		request = NewGetRollbackProgressRequest()
	}
	response = NewGetRollbackProgressResponse()
	err = c.Send(request, response)
	return
}

func NewSetUnschedulableRequest() (request *SetUnschedulableRequest) {
	request = &SetUnschedulableRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SetUnschedulable")
	return
}

func NewSetUnschedulableResponse() (response *SetUnschedulableResponse) {
	response = &SetUnschedulableResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置结点是否可调度新的Pod，不影响已存在的Pod
func (c *Client) SetUnschedulable(request *SetUnschedulableRequest) (response *SetUnschedulableResponse, err error) {
	if request == nil {
		request = NewSetUnschedulableRequest()
	}
	response = NewSetUnschedulableResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRuleGroupsRequest() (request *DeleteRuleGroupsRequest) {
	request = &DeleteRuleGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteRuleGroups")
	return
}

func NewDeleteRuleGroupsResponse() (response *DeleteRuleGroupsResponse) {
	response = &DeleteRuleGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除策略
func (c *Client) DeleteRuleGroups(request *DeleteRuleGroupsRequest) (response *DeleteRuleGroupsResponse, err error) {
	if request == nil {
		request = NewDeleteRuleGroupsRequest()
	}
	response = NewDeleteRuleGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteServersRequest() (request *DeleteServersRequest) {
	request = &DeleteServersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteServers")
	return
}

func NewDeleteServersResponse() (response *DeleteServersResponse) {
	response = &DeleteServersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除服务器信息
func (c *Client) DeleteServers(request *DeleteServersRequest) (response *DeleteServersResponse, err error) {
	if request == nil {
		request = NewDeleteServersRequest()
	}
	response = NewDeleteServersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQueryRecordsRequest() (request *DescribeQueryRecordsRequest) {
	request = &DescribeQueryRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeQueryRecords")
	return
}

func NewDescribeQueryRecordsResponse() (response *DescribeQueryRecordsResponse) {
	response = &DescribeQueryRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询搜索记录
func (c *Client) DescribeQueryRecords(request *DescribeQueryRecordsRequest) (response *DescribeQueryRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeQueryRecordsRequest()
	}
	response = NewDescribeQueryRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewConfigServerDefRequest() (request *ConfigServerDefRequest) {
	request = &ConfigServerDefRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ConfigServerDef")
	return
}

func NewConfigServerDefResponse() (response *ConfigServerDefResponse) {
	response = &ConfigServerDefResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置初始化
func (c *Client) ConfigServerDef(request *ConfigServerDefRequest) (response *ConfigServerDefResponse, err error) {
	if request == nil {
		request = NewConfigServerDefRequest()
	}
	response = NewConfigServerDefResponse()
	err = c.Send(request, response)
	return
}

func NewRecycleVMLanIPListRequest() (request *RecycleVMLanIPListRequest) {
	request = &RecycleVMLanIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "RecycleVMLanIPList")
	return
}

func NewRecycleVMLanIPListResponse() (response *RecycleVMLanIPListResponse) {
	response = &RecycleVMLanIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回收VM内网IP
func (c *Client) RecycleVMLanIPList(request *RecycleVMLanIPListRequest) (response *RecycleVMLanIPListResponse, err error) {
	if request == nil {
		request = NewRecycleVMLanIPListRequest()
	}
	response = NewRecycleVMLanIPListResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateLogConfigStateRequest() (request *UpdateLogConfigStateRequest) {
	request = &UpdateLogConfigStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "UpdateLogConfigState")
	return
}

func NewUpdateLogConfigStateResponse() (response *UpdateLogConfigStateResponse) {
	response = &UpdateLogConfigStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志配置状态变更接口 ，启用或禁用，支持批量
func (c *Client) UpdateLogConfigState(request *UpdateLogConfigStateRequest) (response *UpdateLogConfigStateResponse, err error) {
	if request == nil {
		request = NewUpdateLogConfigStateRequest()
	}
	response = NewUpdateLogConfigStateResponse()
	err = c.Send(request, response)
	return
}

func NewListIngressClassesRequest() (request *ListIngressClassesRequest) {
	request = &ListIngressClassesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListIngressClasses")
	return
}

func NewListIngressClassesResponse() (response *ListIngressClassesResponse) {
	response = &ListIngressClassesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询路由控制器列表
func (c *Client) ListIngressClasses(request *ListIngressClassesRequest) (response *ListIngressClassesResponse, err error) {
	if request == nil {
		request = NewListIngressClassesRequest()
	}
	response = NewListIngressClassesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllVlanIdsExRequest() (request *DescribeAllVlanIdsExRequest) {
	request = &DescribeAllVlanIdsExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeAllVlanIdsEx")
	return
}

func NewDescribeAllVlanIdsExResponse() (response *DescribeAllVlanIdsExResponse) {
	response = &DescribeAllVlanIdsExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有vlanid接口
func (c *Client) DescribeAllVlanIdsEx(request *DescribeAllVlanIdsExRequest) (response *DescribeAllVlanIdsExResponse, err error) {
	if request == nil {
		request = NewDescribeAllVlanIdsExRequest()
	}
	response = NewDescribeAllVlanIdsExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOutBandIpResourceIpDetailExRequest() (request *DescribeOutBandIpResourceIpDetailExRequest) {
	request = &DescribeOutBandIpResourceIpDetailExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeOutBandIpResourceIpDetailEx")
	return
}

func NewDescribeOutBandIpResourceIpDetailExResponse() (response *DescribeOutBandIpResourceIpDetailExResponse) {
	response = &DescribeOutBandIpResourceIpDetailExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeOutBandIpResourceIpDetailEx 查询带外ip资源详细信息
func (c *Client) DescribeOutBandIpResourceIpDetailEx(request *DescribeOutBandIpResourceIpDetailExRequest) (response *DescribeOutBandIpResourceIpDetailExResponse, err error) {
	if request == nil {
		request = NewDescribeOutBandIpResourceIpDetailExRequest()
	}
	response = NewDescribeOutBandIpResourceIpDetailExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSvrNicAllocationRequest() (request *DescribeSvrNicAllocationRequest) {
	request = &DescribeSvrNicAllocationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeSvrNicAllocation")
	return
}

func NewDescribeSvrNicAllocationResponse() (response *DescribeSvrNicAllocationResponse) {
	response = &DescribeSvrNicAllocationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeSvrNicAllocation 查询已绑定多网卡策略的服务器信息
func (c *Client) DescribeSvrNicAllocation(request *DescribeSvrNicAllocationRequest) (response *DescribeSvrNicAllocationResponse, err error) {
	if request == nil {
		request = NewDescribeSvrNicAllocationRequest()
	}
	response = NewDescribeSvrNicAllocationResponse()
	err = c.Send(request, response)
	return
}

func NewEvictNodeRequest() (request *EvictNodeRequest) {
	request = &EvictNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "EvictNode")
	return
}

func NewEvictNodeResponse() (response *EvictNodeResponse) {
	response = &EvictNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 驱逐结点上所有pod
func (c *Client) EvictNode(request *EvictNodeRequest) (response *EvictNodeResponse, err error) {
	if request == nil {
		request = NewEvictNodeRequest()
	}
	response = NewEvictNodeResponse()
	err = c.Send(request, response)
	return
}

func NewReinstallOsExRequest() (request *ReinstallOsExRequest) {
	request = &ReinstallOsExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ReinstallOsEx")
	return
}

func NewReinstallOsExResponse() (response *ReinstallOsExResponse) {
	response = &ReinstallOsExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// OS部署 部署系统
func (c *Client) ReinstallOsEx(request *ReinstallOsExRequest) (response *ReinstallOsExResponse, err error) {
	if request == nil {
		request = NewReinstallOsExRequest()
	}
	response = NewReinstallOsExResponse()
	err = c.Send(request, response)
	return
}

func NewCreateImageBuildingTaskRequest() (request *CreateImageBuildingTaskRequest) {
	request = &CreateImageBuildingTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateImageBuildingTask")
	return
}

func NewCreateImageBuildingTaskResponse() (response *CreateImageBuildingTaskResponse) {
	response = &CreateImageBuildingTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建镜像构建任务
func (c *Client) CreateImageBuildingTask(request *CreateImageBuildingTaskRequest) (response *CreateImageBuildingTaskResponse, err error) {
	if request == nil {
		request = NewCreateImageBuildingTaskRequest()
	}
	response = NewCreateImageBuildingTaskResponse()
	err = c.Send(request, response)
	return
}

func NewGetDefaultGroupByRequest() (request *GetDefaultGroupByRequest) {
	request = &GetDefaultGroupByRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetDefaultGroupBy")
	return
}

func NewGetDefaultGroupByResponse() (response *GetDefaultGroupByResponse) {
	response = &GetDefaultGroupByResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取聚合维度
func (c *Client) GetDefaultGroupBy(request *GetDefaultGroupByRequest) (response *GetDefaultGroupByResponse, err error) {
	if request == nil {
		request = NewGetDefaultGroupByRequest()
	}
	response = NewGetDefaultGroupByResponse()
	err = c.Send(request, response)
	return
}

func NewSetDefaultNetworkPolicyRequest() (request *SetDefaultNetworkPolicyRequest) {
	request = &SetDefaultNetworkPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SetDefaultNetworkPolicy")
	return
}

func NewSetDefaultNetworkPolicyResponse() (response *SetDefaultNetworkPolicyResponse) {
	response = &SetDefaultNetworkPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置集群默认网络策略
func (c *Client) SetDefaultNetworkPolicy(request *SetDefaultNetworkPolicyRequest) (response *SetDefaultNetworkPolicyResponse, err error) {
	if request == nil {
		request = NewSetDefaultNetworkPolicyRequest()
	}
	response = NewSetDefaultNetworkPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLoadBalancerRequest() (request *CreateLoadBalancerRequest) {
	request = &CreateLoadBalancerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateLoadBalancer")
	return
}

func NewCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
	response = &CreateLoadBalancerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建负载均衡实例
func (c *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
	if request == nil {
		request = NewCreateLoadBalancerRequest()
	}
	response = NewCreateLoadBalancerResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePersistentVolumeClaimRequest() (request *DeletePersistentVolumeClaimRequest) {
	request = &DeletePersistentVolumeClaimRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeletePersistentVolumeClaim")
	return
}

func NewDeletePersistentVolumeClaimResponse() (response *DeletePersistentVolumeClaimResponse) {
	response = &DeletePersistentVolumeClaimResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除PVC
func (c *Client) DeletePersistentVolumeClaim(request *DeletePersistentVolumeClaimRequest) (response *DeletePersistentVolumeClaimResponse, err error) {
	if request == nil {
		request = NewDeletePersistentVolumeClaimRequest()
	}
	response = NewDeletePersistentVolumeClaimResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogConfigsRequest() (request *DescribeLogConfigsRequest) {
	request = &DescribeLogConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeLogConfigs")
	return
}

func NewDescribeLogConfigsResponse() (response *DescribeLogConfigsResponse) {
	response = &DescribeLogConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询接入配置
func (c *Client) DescribeLogConfigs(request *DescribeLogConfigsRequest) (response *DescribeLogConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeLogConfigsRequest()
	}
	response = NewDescribeLogConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateDashboardPermissionsByIdRequest() (request *UpdateDashboardPermissionsByIdRequest) {
	request = &UpdateDashboardPermissionsByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "UpdateDashboardPermissionsById")
	return
}

func NewUpdateDashboardPermissionsByIdResponse() (response *UpdateDashboardPermissionsByIdResponse) {
	response = &UpdateDashboardPermissionsByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新仪表盘权限
func (c *Client) UpdateDashboardPermissionsById(request *UpdateDashboardPermissionsByIdRequest) (response *UpdateDashboardPermissionsByIdResponse, err error) {
	if request == nil {
		request = NewUpdateDashboardPermissionsByIdRequest()
	}
	response = NewUpdateDashboardPermissionsByIdResponse()
	err = c.Send(request, response)
	return
}

func NewCreateHPARequest() (request *CreateHPARequest) {
	request = &CreateHPARequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "CreateHPA")
	return
}

func NewCreateHPAResponse() (response *CreateHPAResponse) {
	response = &CreateHPAResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建HPA
func (c *Client) CreateHPA(request *CreateHPARequest) (response *CreateHPAResponse, err error) {
	if request == nil {
		request = NewCreateHPARequest()
	}
	response = NewCreateHPAResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTargetPortRequest() (request *ModifyTargetPortRequest) {
	request = &ModifyTargetPortRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyTargetPort")
	return
}

func NewModifyTargetPortResponse() (response *ModifyTargetPortResponse) {
	response = &ModifyTargetPortResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改监听器绑定的后端机器的端口
func (c *Client) ModifyTargetPort(request *ModifyTargetPortRequest) (response *ModifyTargetPortResponse, err error) {
	if request == nil {
		request = NewModifyTargetPortRequest()
	}
	response = NewModifyTargetPortResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogRuleGroupsRequest() (request *DeleteLogRuleGroupsRequest) {
	request = &DeleteLogRuleGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteLogRuleGroups")
	return
}

func NewDeleteLogRuleGroupsResponse() (response *DeleteLogRuleGroupsResponse) {
	response = &DeleteLogRuleGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除日志报警策略
func (c *Client) DeleteLogRuleGroups(request *DeleteLogRuleGroupsRequest) (response *DeleteLogRuleGroupsResponse, err error) {
	if request == nil {
		request = NewDeleteLogRuleGroupsRequest()
	}
	response = NewDeleteLogRuleGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageBuildingTaskRequest() (request *DescribeImageBuildingTaskRequest) {
	request = &DescribeImageBuildingTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeImageBuildingTask")
	return
}

func NewDescribeImageBuildingTaskResponse() (response *DescribeImageBuildingTaskResponse) {
	response = &DescribeImageBuildingTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像构建任务
func (c *Client) DescribeImageBuildingTask(request *DescribeImageBuildingTaskRequest) (response *DescribeImageBuildingTaskResponse, err error) {
	if request == nil {
		request = NewDescribeImageBuildingTaskRequest()
	}
	response = NewDescribeImageBuildingTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkPolicyRequest() (request *DescribeNetworkPolicyRequest) {
	request = &DescribeNetworkPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeNetworkPolicy")
	return
}

func NewDescribeNetworkPolicyResponse() (response *DescribeNetworkPolicyResponse) {
	response = &DescribeNetworkPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络策略详情
func (c *Client) DescribeNetworkPolicy(request *DescribeNetworkPolicyRequest) (response *DescribeNetworkPolicyResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkPolicyRequest()
	}
	response = NewDescribeNetworkPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyJobRequest() (request *ModifyJobRequest) {
	request = &ModifyJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyJob")
	return
}

func NewModifyJobResponse() (response *ModifyJobResponse) {
	response = &ModifyJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新Job
func (c *Client) ModifyJob(request *ModifyJobRequest) (response *ModifyJobResponse, err error) {
	if request == nil {
		request = NewModifyJobRequest()
	}
	response = NewModifyJobResponse()
	err = c.Send(request, response)
	return
}

func NewGetAlertTrendRequest() (request *GetAlertTrendRequest) {
	request = &GetAlertTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetAlertTrend")
	return
}

func NewGetAlertTrendResponse() (response *GetAlertTrendResponse) {
	response = &GetAlertTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警趋势接口
func (c *Client) GetAlertTrend(request *GetAlertTrendRequest) (response *GetAlertTrendResponse, err error) {
	if request == nil {
		request = NewGetAlertTrendRequest()
	}
	response = NewGetAlertTrendResponse()
	err = c.Send(request, response)
	return
}

func NewDoUpgradeOperationRequest() (request *DoUpgradeOperationRequest) {
	request = &DoUpgradeOperationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DoUpgradeOperation")
	return
}

func NewDoUpgradeOperationResponse() (response *DoUpgradeOperationResponse) {
	response = &DoUpgradeOperationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 执行升级流程中的操作
func (c *Client) DoUpgradeOperation(request *DoUpgradeOperationRequest) (response *DoUpgradeOperationResponse, err error) {
	if request == nil {
		request = NewDoUpgradeOperationRequest()
	}
	response = NewDoUpgradeOperationResponse()
	err = c.Send(request, response)
	return
}

func NewSocPerformServerTaskExRequest() (request *SocPerformServerTaskExRequest) {
	request = &SocPerformServerTaskExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SocPerformServerTaskEx")
	return
}

func NewSocPerformServerTaskExResponse() (response *SocPerformServerTaskExResponse) {
	response = &SocPerformServerTaskExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 黑石服务器带外操作 开机 关机 重启
func (c *Client) SocPerformServerTaskEx(request *SocPerformServerTaskExRequest) (response *SocPerformServerTaskExResponse, err error) {
	if request == nil {
		request = NewSocPerformServerTaskExRequest()
	}
	response = NewSocPerformServerTaskExResponse()
	err = c.Send(request, response)
	return
}

func NewGetNodeDisruptionBudgetRequest() (request *GetNodeDisruptionBudgetRequest) {
	request = &GetNodeDisruptionBudgetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetNodeDisruptionBudget")
	return
}

func NewGetNodeDisruptionBudgetResponse() (response *GetNodeDisruptionBudgetResponse) {
	response = &GetNodeDisruptionBudgetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取自愈风险控制配置
func (c *Client) GetNodeDisruptionBudget(request *GetNodeDisruptionBudgetRequest) (response *GetNodeDisruptionBudgetResponse, err error) {
	if request == nil {
		request = NewGetNodeDisruptionBudgetRequest()
	}
	response = NewGetNodeDisruptionBudgetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTcsComponentRequest() (request *DescribeTcsComponentRequest) {
	request = &DescribeTcsComponentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeTcsComponent")
	return
}

func NewDescribeTcsComponentResponse() (response *DescribeTcsComponentResponse) {
	response = &DescribeTcsComponentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeTcsComponent
func (c *Client) DescribeTcsComponent(request *DescribeTcsComponentRequest) (response *DescribeTcsComponentResponse, err error) {
	if request == nil {
		request = NewDescribeTcsComponentRequest()
	}
	response = NewDescribeTcsComponentResponse()
	err = c.Send(request, response)
	return
}

func NewListNativeStorageClassesRequest() (request *ListNativeStorageClassesRequest) {
	request = &ListNativeStorageClassesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListNativeStorageClasses")
	return
}

func NewListNativeStorageClassesResponse() (response *ListNativeStorageClassesResponse) {
	response = &ListNativeStorageClassesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取原生StorageClass列表
func (c *Client) ListNativeStorageClasses(request *ListNativeStorageClassesRequest) (response *ListNativeStorageClassesResponse, err error) {
	if request == nil {
		request = NewListNativeStorageClassesRequest()
	}
	response = NewListNativeStorageClassesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNetworkPolicyRequest() (request *DeleteNetworkPolicyRequest) {
	request = &DeleteNetworkPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteNetworkPolicy")
	return
}

func NewDeleteNetworkPolicyResponse() (response *DeleteNetworkPolicyResponse) {
	response = &DeleteNetworkPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网络策略
func (c *Client) DeleteNetworkPolicy(request *DeleteNetworkPolicyRequest) (response *DeleteNetworkPolicyResponse, err error) {
	if request == nil {
		request = NewDeleteNetworkPolicyRequest()
	}
	response = NewDeleteNetworkPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductOverviewRequest() (request *DescribeProductOverviewRequest) {
	request = &DescribeProductOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeProductOverview")
	return
}

func NewDescribeProductOverviewResponse() (response *DescribeProductOverviewResponse) {
	response = &DescribeProductOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云产品状态接口
func (c *Client) DescribeProductOverview(request *DescribeProductOverviewRequest) (response *DescribeProductOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeProductOverviewRequest()
	}
	response = NewDescribeProductOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewQueryDcosFunctionServiceRelationRequest() (request *QueryDcosFunctionServiceRelationRequest) {
	request = &QueryDcosFunctionServiceRelationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "QueryDcosFunctionServiceRelation")
	return
}

func NewQueryDcosFunctionServiceRelationResponse() (response *QueryDcosFunctionServiceRelationResponse) {
	response = &QueryDcosFunctionServiceRelationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询dcos功能的后端服务ID
func (c *Client) QueryDcosFunctionServiceRelation(request *QueryDcosFunctionServiceRelationRequest) (response *QueryDcosFunctionServiceRelationResponse, err error) {
	if request == nil {
		request = NewQueryDcosFunctionServiceRelationRequest()
	}
	response = NewQueryDcosFunctionServiceRelationResponse()
	err = c.Send(request, response)
	return
}

func NewListBaseNodesRequest() (request *ListBaseNodesRequest) {
	request = &ListBaseNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListBaseNodes")
	return
}

func NewListBaseNodesResponse() (response *ListBaseNodesResponse) {
	response = &ListBaseNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点基本信息列表
func (c *Client) ListBaseNodes(request *ListBaseNodesRequest) (response *ListBaseNodesResponse, err error) {
	if request == nil {
		request = NewListBaseNodesRequest()
	}
	response = NewListBaseNodesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMetaMetricRequest() (request *UpdateMetaMetricRequest) {
	request = &UpdateMetaMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "UpdateMetaMetric")
	return
}

func NewUpdateMetaMetricResponse() (response *UpdateMetaMetricResponse) {
	response = &UpdateMetaMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新指标信息
func (c *Client) UpdateMetaMetric(request *UpdateMetaMetricRequest) (response *UpdateMetaMetricResponse, err error) {
	if request == nil {
		request = NewUpdateMetaMetricRequest()
	}
	response = NewUpdateMetaMetricResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateVMVirtualIPRequest() (request *AllocateVMVirtualIPRequest) {
	request = &AllocateVMVirtualIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "AllocateVMVirtualIP")
	return
}

func NewAllocateVMVirtualIPResponse() (response *AllocateVMVirtualIPResponse) {
	response = &AllocateVMVirtualIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分配VM虚拟内网IP
func (c *Client) AllocateVMVirtualIP(request *AllocateVMVirtualIPRequest) (response *AllocateVMVirtualIPResponse, err error) {
	if request == nil {
		request = NewAllocateVMVirtualIPRequest()
	}
	response = NewAllocateVMVirtualIPResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLabelRequest() (request *DeleteLabelRequest) {
	request = &DeleteLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteLabel")
	return
}

func NewDeleteLabelResponse() (response *DeleteLabelResponse) {
	response = &DeleteLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 标签管理--删除标签
func (c *Client) DeleteLabel(request *DeleteLabelRequest) (response *DeleteLabelResponse, err error) {
	if request == nil {
		request = NewDeleteLabelRequest()
	}
	response = NewDeleteLabelResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomScriptRequest() (request *ModifyCustomScriptRequest) {
	request = &ModifyCustomScriptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyCustomScript")
	return
}

func NewModifyCustomScriptResponse() (response *ModifyCustomScriptResponse) {
	response = &ModifyCustomScriptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改自定义脚本信息
func (c *Client) ModifyCustomScript(request *ModifyCustomScriptRequest) (response *ModifyCustomScriptResponse, err error) {
	if request == nil {
		request = NewModifyCustomScriptRequest()
	}
	response = NewModifyCustomScriptResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMultiNicDefinitionRequest() (request *ModifyMultiNicDefinitionRequest) {
	request = &ModifyMultiNicDefinitionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyMultiNicDefinition")
	return
}

func NewModifyMultiNicDefinitionResponse() (response *ModifyMultiNicDefinitionResponse) {
	response = &ModifyMultiNicDefinitionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyMultiNicDefinition 修改网卡信息
func (c *Client) ModifyMultiNicDefinition(request *ModifyMultiNicDefinitionRequest) (response *ModifyMultiNicDefinitionResponse, err error) {
	if request == nil {
		request = NewModifyMultiNicDefinitionRequest()
	}
	response = NewModifyMultiNicDefinitionResponse()
	err = c.Send(request, response)
	return
}

func NewListMetaMetricRequest() (request *ListMetaMetricRequest) {
	request = &ListMetaMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListMetaMetric")
	return
}

func NewListMetaMetricResponse() (response *ListMetaMetricResponse) {
	response = &ListMetaMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源对象类型下的指标列表
func (c *Client) ListMetaMetric(request *ListMetaMetricRequest) (response *ListMetaMetricResponse, err error) {
	if request == nil {
		request = NewListMetaMetricRequest()
	}
	response = NewListMetaMetricResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHistoryMonitorDataRequest() (request *DescribeHistoryMonitorDataRequest) {
	request = &DescribeHistoryMonitorDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeHistoryMonitorData")
	return
}

func NewDescribeHistoryMonitorDataResponse() (response *DescribeHistoryMonitorDataResponse) {
	response = &DescribeHistoryMonitorDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询历史监控数据
func (c *Client) DescribeHistoryMonitorData(request *DescribeHistoryMonitorDataRequest) (response *DescribeHistoryMonitorDataResponse, err error) {
	if request == nil {
		request = NewDescribeHistoryMonitorDataRequest()
	}
	response = NewDescribeHistoryMonitorDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDcosFunctionsRequest() (request *DescribeDcosFunctionsRequest) {
	request = &DescribeDcosFunctionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeDcosFunctions")
	return
}

func NewDescribeDcosFunctionsResponse() (response *DescribeDcosFunctionsResponse) {
	response = &DescribeDcosFunctionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// dcos支持功能列表
func (c *Client) DescribeDcosFunctions(request *DescribeDcosFunctionsRequest) (response *DescribeDcosFunctionsResponse, err error) {
	if request == nil {
		request = NewDescribeDcosFunctionsRequest()
	}
	response = NewDescribeDcosFunctionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImagePathRequest() (request *DescribeImagePathRequest) {
	request = &DescribeImagePathRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeImagePath")
	return
}

func NewDescribeImagePathResponse() (response *DescribeImagePathResponse) {
	response = &DescribeImagePathResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取镜像原始地址接口
func (c *Client) DescribeImagePath(request *DescribeImagePathRequest) (response *DescribeImagePathResponse, err error) {
	if request == nil {
		request = NewDescribeImagePathRequest()
	}
	response = NewDescribeImagePathResponse()
	err = c.Send(request, response)
	return
}

func NewGetFiringRulesRequest() (request *GetFiringRulesRequest) {
	request = &GetFiringRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetFiringRules")
	return
}

func NewGetFiringRulesResponse() (response *GetFiringRulesResponse) {
	response = &GetFiringRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取有未恢复告警的规则
func (c *Client) GetFiringRules(request *GetFiringRulesRequest) (response *GetFiringRulesResponse, err error) {
	if request == nil {
		request = NewGetFiringRulesRequest()
	}
	response = NewGetFiringRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMiddlewareComponentRequest() (request *DescribeMiddlewareComponentRequest) {
	request = &DescribeMiddlewareComponentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeMiddlewareComponent")
	return
}

func NewDescribeMiddlewareComponentResponse() (response *DescribeMiddlewareComponentResponse) {
	response = &DescribeMiddlewareComponentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeMiddlewareComponent
func (c *Client) DescribeMiddlewareComponent(request *DescribeMiddlewareComponentRequest) (response *DescribeMiddlewareComponentResponse, err error) {
	if request == nil {
		request = NewDescribeMiddlewareComponentRequest()
	}
	response = NewDescribeMiddlewareComponentResponse()
	err = c.Send(request, response)
	return
}

func NewApplyWebsocketTokenRequest() (request *ApplyWebsocketTokenRequest) {
	request = &ApplyWebsocketTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ApplyWebsocketToken")
	return
}

func NewApplyWebsocketTokenResponse() (response *ApplyWebsocketTokenResponse) {
	response = &ApplyWebsocketTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 申请连接Websocket的Token
func (c *Client) ApplyWebsocketToken(request *ApplyWebsocketTokenRequest) (response *ApplyWebsocketTokenResponse, err error) {
	if request == nil {
		request = NewApplyWebsocketTokenRequest()
	}
	response = NewApplyWebsocketTokenResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteImageRequest() (request *DeleteImageRequest) {
	request = &DeleteImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteImage")
	return
}

func NewDeleteImageResponse() (response *DeleteImageResponse) {
	response = &DeleteImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除镜像
func (c *Client) DeleteImage(request *DeleteImageRequest) (response *DeleteImageResponse, err error) {
	if request == nil {
		request = NewDeleteImageRequest()
	}
	response = NewDeleteImageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIgniterNodeExRequest() (request *DescribeIgniterNodeExRequest) {
	request = &DescribeIgniterNodeExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeIgniterNodeEx")
	return
}

func NewDescribeIgniterNodeExResponse() (response *DescribeIgniterNodeExResponse) {
	response = &DescribeIgniterNodeExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeIgniterNodeEx
func (c *Client) DescribeIgniterNodeEx(request *DescribeIgniterNodeExRequest) (response *DescribeIgniterNodeExResponse, err error) {
	if request == nil {
		request = NewDescribeIgniterNodeExRequest()
	}
	response = NewDescribeIgniterNodeExResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOutbandStrategyRequest() (request *ModifyOutbandStrategyRequest) {
	request = &ModifyOutbandStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ModifyOutbandStrategy")
	return
}

func NewModifyOutbandStrategyResponse() (response *ModifyOutbandStrategyResponse) {
	response = &ModifyOutbandStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改自定义带外密码策略
func (c *Client) ModifyOutbandStrategy(request *ModifyOutbandStrategyRequest) (response *ModifyOutbandStrategyResponse, err error) {
	if request == nil {
		request = NewModifyOutbandStrategyRequest()
	}
	response = NewModifyOutbandStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewReserveOutBandIPRequest() (request *ReserveOutBandIPRequest) {
	request = &ReserveOutBandIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ReserveOutBandIP")
	return
}

func NewReserveOutBandIPResponse() (response *ReserveOutBandIPResponse) {
	response = &ReserveOutBandIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ReserveOutBandIP 预留\取消带外Ip
func (c *Client) ReserveOutBandIP(request *ReserveOutBandIPRequest) (response *ReserveOutBandIPResponse, err error) {
	if request == nil {
		request = NewReserveOutBandIPRequest()
	}
	response = NewReserveOutBandIPResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNodesOverviewWithAlertsRequest() (request *DescribeNodesOverviewWithAlertsRequest) {
	request = &DescribeNodesOverviewWithAlertsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeNodesOverviewWithAlerts")
	return
}

func NewDescribeNodesOverviewWithAlertsResponse() (response *DescribeNodesOverviewWithAlertsResponse) {
	response = &DescribeNodesOverviewWithAlertsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云哨概览页获取节点信息
func (c *Client) DescribeNodesOverviewWithAlerts(request *DescribeNodesOverviewWithAlertsRequest) (response *DescribeNodesOverviewWithAlertsResponse, err error) {
	if request == nil {
		request = NewDescribeNodesOverviewWithAlertsRequest()
	}
	response = NewDescribeNodesOverviewWithAlertsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceRequest() (request *DescribeServiceRequest) {
	request = &DescribeServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeService")
	return
}

func NewDescribeServiceResponse() (response *DescribeServiceResponse) {
	response = &DescribeServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务
func (c *Client) DescribeService(request *DescribeServiceRequest) (response *DescribeServiceResponse, err error) {
	if request == nil {
		request = NewDescribeServiceRequest()
	}
	response = NewDescribeServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDefaultNetworkPolicyRequest() (request *DescribeDefaultNetworkPolicyRequest) {
	request = &DescribeDefaultNetworkPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeDefaultNetworkPolicy")
	return
}

func NewDescribeDefaultNetworkPolicyResponse() (response *DescribeDefaultNetworkPolicyResponse) {
	response = &DescribeDefaultNetworkPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询默认网络策略
func (c *Client) DescribeDefaultNetworkPolicy(request *DescribeDefaultNetworkPolicyRequest) (response *DescribeDefaultNetworkPolicyResponse, err error) {
	if request == nil {
		request = NewDescribeDefaultNetworkPolicyRequest()
	}
	response = NewDescribeDefaultNetworkPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeK8sComponentRequest() (request *DescribeK8sComponentRequest) {
	request = &DescribeK8sComponentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeK8sComponent")
	return
}

func NewDescribeK8sComponentResponse() (response *DescribeK8sComponentResponse) {
	response = &DescribeK8sComponentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeK8sComponent
func (c *Client) DescribeK8sComponent(request *DescribeK8sComponentRequest) (response *DescribeK8sComponentResponse, err error) {
	if request == nil {
		request = NewDescribeK8sComponentRequest()
	}
	response = NewDescribeK8sComponentResponse()
	err = c.Send(request, response)
	return
}

func NewGenLineHeaderRegexRequest() (request *GenLineHeaderRegexRequest) {
	request = &GenLineHeaderRegexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GenLineHeaderRegex")
	return
}

func NewGenLineHeaderRegexResponse() (response *GenLineHeaderRegexResponse) {
	response = &GenLineHeaderRegexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 自动生成行首正则
func (c *Client) GenLineHeaderRegex(request *GenLineHeaderRegexRequest) (response *GenLineHeaderRegexResponse, err error) {
	if request == nil {
		request = NewGenLineHeaderRegexRequest()
	}
	response = NewGenLineHeaderRegexResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRuleGroupTemplesRequest() (request *DeleteRuleGroupTemplesRequest) {
	request = &DeleteRuleGroupTemplesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteRuleGroupTemples")
	return
}

func NewDeleteRuleGroupTemplesResponse() (response *DeleteRuleGroupTemplesResponse) {
	response = &DeleteRuleGroupTemplesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除策略模板
func (c *Client) DeleteRuleGroupTemples(request *DeleteRuleGroupTemplesRequest) (response *DeleteRuleGroupTemplesResponse, err error) {
	if request == nil {
		request = NewDeleteRuleGroupTemplesRequest()
	}
	response = NewDeleteRuleGroupTemplesResponse()
	err = c.Send(request, response)
	return
}

func NewGetRelatedAlertNamesRequest() (request *GetRelatedAlertNamesRequest) {
	request = &GetRelatedAlertNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetRelatedAlertNames")
	return
}

func NewGetRelatedAlertNamesResponse() (response *GetRelatedAlertNamesResponse) {
	response = &GetRelatedAlertNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量通过事件名获取关联报警列表
func (c *Client) GetRelatedAlertNames(request *GetRelatedAlertNamesRequest) (response *GetRelatedAlertNamesResponse, err error) {
	if request == nil {
		request = NewGetRelatedAlertNamesRequest()
	}
	response = NewGetRelatedAlertNamesResponse()
	err = c.Send(request, response)
	return
}

func NewRecycleVMWanIPListRequest() (request *RecycleVMWanIPListRequest) {
	request = &RecycleVMWanIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "RecycleVMWanIPList")
	return
}

func NewRecycleVMWanIPListResponse() (response *RecycleVMWanIPListResponse) {
	response = &RecycleVMWanIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回收VM外网IP
func (c *Client) RecycleVMWanIPList(request *RecycleVMWanIPListRequest) (response *RecycleVMWanIPListResponse, err error) {
	if request == nil {
		request = NewRecycleVMWanIPListRequest()
	}
	response = NewRecycleVMWanIPListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteImageBuildingTaskRequest() (request *DeleteImageBuildingTaskRequest) {
	request = &DeleteImageBuildingTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteImageBuildingTask")
	return
}

func NewDeleteImageBuildingTaskResponse() (response *DeleteImageBuildingTaskResponse) {
	response = &DeleteImageBuildingTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除镜像构建任务
func (c *Client) DeleteImageBuildingTask(request *DeleteImageBuildingTaskRequest) (response *DeleteImageBuildingTaskResponse, err error) {
	if request == nil {
		request = NewDeleteImageBuildingTaskRequest()
	}
	response = NewDeleteImageBuildingTaskResponse()
	err = c.Send(request, response)
	return
}

func NewListSecretsRequest() (request *ListSecretsRequest) {
	request = &ListSecretsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListSecrets")
	return
}

func NewListSecretsResponse() (response *ListSecretsResponse) {
	response = &ListSecretsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Secret列表
func (c *Client) ListSecrets(request *ListSecretsRequest) (response *ListSecretsResponse, err error) {
	if request == nil {
		request = NewListSecretsRequest()
	}
	response = NewListSecretsResponse()
	err = c.Send(request, response)
	return
}

func NewAddLabelRequest() (request *AddLabelRequest) {
	request = &AddLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "AddLabel")
	return
}

func NewAddLabelResponse() (response *AddLabelResponse) {
	response = &AddLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 标签管理--新增标签
func (c *Client) AddLabel(request *AddLabelRequest) (response *AddLabelResponse, err error) {
	if request == nil {
		request = NewAddLabelRequest()
	}
	response = NewAddLabelResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateVMWanIPListRequest() (request *AllocateVMWanIPListRequest) {
	request = &AllocateVMWanIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "AllocateVMWanIPList")
	return
}

func NewAllocateVMWanIPListResponse() (response *AllocateVMWanIPListResponse) {
	response = &AllocateVMWanIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分配VM外网IP
func (c *Client) AllocateVMWanIPList(request *AllocateVMWanIPListRequest) (response *AllocateVMWanIPListResponse, err error) {
	if request == nil {
		request = NewAllocateVMWanIPListRequest()
	}
	response = NewAllocateVMWanIPListResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceObjectRequest() (request *ListResourceObjectRequest) {
	request = &ListResourceObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListResourceObject")
	return
}

func NewListResourceObjectResponse() (response *ListResourceObjectResponse) {
	response = &ListResourceObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务树资源对象
func (c *Client) ListResourceObject(request *ListResourceObjectRequest) (response *ListResourceObjectResponse, err error) {
	if request == nil {
		request = NewListResourceObjectRequest()
	}
	response = NewListResourceObjectResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDashboardByIdRequest() (request *DeleteDashboardByIdRequest) {
	request = &DeleteDashboardByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DeleteDashboardById")
	return
}

func NewDeleteDashboardByIdResponse() (response *DeleteDashboardByIdResponse) {
	response = &DeleteDashboardByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根椐唯一ID删除仪表盘
func (c *Client) DeleteDashboardById(request *DeleteDashboardByIdRequest) (response *DeleteDashboardByIdResponse, err error) {
	if request == nil {
		request = NewDeleteDashboardByIdRequest()
	}
	response = NewDeleteDashboardByIdResponse()
	err = c.Send(request, response)
	return
}

func NewListClusterStatusRequest() (request *ListClusterStatusRequest) {
	request = &ListClusterStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "ListClusterStatus")
	return
}

func NewListClusterStatusResponse() (response *ListClusterStatusResponse) {
	response = &ListClusterStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群健康情况
func (c *Client) ListClusterStatus(request *ListClusterStatusRequest) (response *ListClusterStatusResponse, err error) {
	if request == nil {
		request = NewListClusterStatusRequest()
	}
	response = NewListClusterStatusResponse()
	err = c.Send(request, response)
	return
}

func NewSearchRetentionsRequest() (request *SearchRetentionsRequest) {
	request = &SearchRetentionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "SearchRetentions")
	return
}

func NewSearchRetentionsResponse() (response *SearchRetentionsResponse) {
	response = &SearchRetentionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索Retention
func (c *Client) SearchRetentions(request *SearchRetentionsRequest) (response *SearchRetentionsResponse, err error) {
	if request == nil {
		request = NewSearchRetentionsRequest()
	}
	response = NewSearchRetentionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClustersOverviewRequest() (request *DescribeClustersOverviewRequest) {
	request = &DescribeClustersOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "DescribeClustersOverview")
	return
}

func NewDescribeClustersOverviewResponse() (response *DescribeClustersOverviewResponse) {
	response = &DescribeClustersOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群网格状态
func (c *Client) DescribeClustersOverview(request *DescribeClustersOverviewRequest) (response *DescribeClustersOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeClustersOverviewRequest()
	}
	response = NewDescribeClustersOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewGetDashFolderByUidRequest() (request *GetDashFolderByUidRequest) {
	request = &GetDashFolderByUidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsplatform", APIVersion, "GetDashFolderByUid")
	return
}

func NewGetDashFolderByUidResponse() (response *GetDashFolderByUidResponse) {
	response = &GetDashFolderByUidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过唯一ID获取详情
func (c *Client) GetDashFolderByUid(request *GetDashFolderByUidRequest) (response *GetDashFolderByUidResponse, err error) {
	if request == nil {
		request = NewGetDashFolderByUidRequest()
	}
	response = NewGetDashFolderByUidResponse()
	err = c.Send(request, response)
	return
}
