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
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-04-13"

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

func NewGetBucketFlowControlRequest() (request *GetBucketFlowControlRequest) {
	request = &GetBucketFlowControlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetBucketFlowControl")
	return
}

func NewGetBucketFlowControlResponse() (response *GetBucketFlowControlResponse) {
	response = &GetBucketFlowControlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取存储桶频控配置
func (c *Client) GetBucketFlowControl(request *GetBucketFlowControlRequest) (response *GetBucketFlowControlResponse, err error) {
	if request == nil {
		request = NewGetBucketFlowControlRequest()
	}
	response = NewGetBucketFlowControlResponse()
	err = c.Send(request, response)
	return
}

func NewGetObjectStorageBucketQuotaRequest() (request *GetObjectStorageBucketQuotaRequest) {
	request = &GetObjectStorageBucketQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetObjectStorageBucketQuota")
	return
}

func NewGetObjectStorageBucketQuotaResponse() (response *GetObjectStorageBucketQuotaResponse) {
	response = &GetObjectStorageBucketQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询存储桶配额
func (c *Client) GetObjectStorageBucketQuota(request *GetObjectStorageBucketQuotaRequest) (response *GetObjectStorageBucketQuotaResponse, err error) {
	if request == nil {
		request = NewGetObjectStorageBucketQuotaRequest()
	}
	response = NewGetObjectStorageBucketQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewCleanInspectionRequest() (request *CleanInspectionRequest) {
	request = &CleanInspectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "CleanInspection")
	return
}

func NewCleanInspectionResponse() (response *CleanInspectionResponse) {
	response = &CleanInspectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除巡检结果
func (c *Client) CleanInspection(request *CleanInspectionRequest) (response *CleanInspectionResponse, err error) {
	if request == nil {
		request = NewCleanInspectionRequest()
	}
	response = NewCleanInspectionResponse()
	err = c.Send(request, response)
	return
}

func NewDoEditHostRequest() (request *DoEditHostRequest) {
	request = &DoEditHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoEditHost")
	return
}

func NewDoEditHostResponse() (response *DoEditHostResponse) {
	response = &DoEditHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑节点（仅用于创建集群过程中）
func (c *Client) DoEditHost(request *DoEditHostRequest) (response *DoEditHostResponse, err error) {
	if request == nil {
		request = NewDoEditHostRequest()
	}
	response = NewDoEditHostResponse()
	err = c.Send(request, response)
	return
}

func NewDoRestartDiskRequest() (request *DoRestartDiskRequest) {
	request = &DoRestartDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoRestartDisk")
	return
}

func NewDoRestartDiskResponse() (response *DoRestartDiskResponse) {
	response = &DoRestartDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重启硬盘
func (c *Client) DoRestartDisk(request *DoRestartDiskRequest) (response *DoRestartDiskResponse, err error) {
	if request == nil {
		request = NewDoRestartDiskRequest()
	}
	response = NewDoRestartDiskResponse()
	err = c.Send(request, response)
	return
}

func NewGetPoolOverviewRequest() (request *GetPoolOverviewRequest) {
	request = &GetPoolOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetPoolOverview")
	return
}

func NewGetPoolOverviewResponse() (response *GetPoolOverviewResponse) {
	response = &GetPoolOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取存储池概览数据
func (c *Client) GetPoolOverview(request *GetPoolOverviewRequest) (response *GetPoolOverviewResponse, err error) {
	if request == nil {
		request = NewGetPoolOverviewRequest()
	}
	response = NewGetPoolOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewGetHostCpuMemoryDataRequest() (request *GetHostCpuMemoryDataRequest) {
	request = &GetHostCpuMemoryDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetHostCpuMemoryData")
	return
}

func NewGetHostCpuMemoryDataResponse() (response *GetHostCpuMemoryDataResponse) {
	response = &GetHostCpuMemoryDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载节点的CPU数据
func (c *Client) GetHostCpuMemoryData(request *GetHostCpuMemoryDataRequest) (response *GetHostCpuMemoryDataResponse, err error) {
	if request == nil {
		request = NewGetHostCpuMemoryDataRequest()
	}
	response = NewGetHostCpuMemoryDataResponse()
	err = c.Send(request, response)
	return
}

func NewGetInitClusterRequestIdRequest() (request *GetInitClusterRequestIdRequest) {
	request = &GetInitClusterRequestIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetInitClusterRequestId")
	return
}

func NewGetInitClusterRequestIdResponse() (response *GetInitClusterRequestIdResponse) {
	response = &GetInitClusterRequestIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取初始化集群的 requestId
func (c *Client) GetInitClusterRequestId(request *GetInitClusterRequestIdRequest) (response *GetInitClusterRequestIdResponse, err error) {
	if request == nil {
		request = NewGetInitClusterRequestIdRequest()
	}
	response = NewGetInitClusterRequestIdResponse()
	err = c.Send(request, response)
	return
}

func NewGetOSUsageOverviewGroupByUserIdRequest() (request *GetOSUsageOverviewGroupByUserIdRequest) {
	request = &GetOSUsageOverviewGroupByUserIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetOSUsageOverviewGroupByUserId")
	return
}

func NewGetOSUsageOverviewGroupByUserIdResponse() (response *GetOSUsageOverviewGroupByUserIdResponse) {
	response = &GetOSUsageOverviewGroupByUserIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按UserId分组获取存储桶用量总览
func (c *Client) GetOSUsageOverviewGroupByUserId(request *GetOSUsageOverviewGroupByUserIdRequest) (response *GetOSUsageOverviewGroupByUserIdResponse, err error) {
	if request == nil {
		request = NewGetOSUsageOverviewGroupByUserIdRequest()
	}
	response = NewGetOSUsageOverviewGroupByUserIdResponse()
	err = c.Send(request, response)
	return
}

func NewGetRecoveryControlConfigRequest() (request *GetRecoveryControlConfigRequest) {
	request = &GetRecoveryControlConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetRecoveryControlConfig")
	return
}

func NewGetRecoveryControlConfigResponse() (response *GetRecoveryControlConfigResponse) {
	response = &GetRecoveryControlConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群恢复控制设置
func (c *Client) GetRecoveryControlConfig(request *GetRecoveryControlConfigRequest) (response *GetRecoveryControlConfigResponse, err error) {
	if request == nil {
		request = NewGetRecoveryControlConfigRequest()
	}
	response = NewGetRecoveryControlConfigResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductHealthStateRequest() (request *QueryProductHealthStateRequest) {
	request = &QueryProductHealthStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "QueryProductHealthState")
	return
}

func NewQueryProductHealthStateResponse() (response *QueryProductHealthStateResponse) {
	response = &QueryProductHealthStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容灾系统产品健康状态检查接口
func (c *Client) QueryProductHealthState(request *QueryProductHealthStateRequest) (response *QueryProductHealthStateResponse, err error) {
	if request == nil {
		request = NewQueryProductHealthStateRequest()
	}
	response = NewQueryProductHealthStateResponse()
	err = c.Send(request, response)
	return
}

func NewGetSnmpConfigRequest() (request *GetSnmpConfigRequest) {
	request = &GetSnmpConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetSnmpConfig")
	return
}

func NewGetSnmpConfigResponse() (response *GetSnmpConfigResponse) {
	response = &GetSnmpConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群SNMP设置
func (c *Client) GetSnmpConfig(request *GetSnmpConfigRequest) (response *GetSnmpConfigResponse, err error) {
	if request == nil {
		request = NewGetSnmpConfigRequest()
	}
	response = NewGetSnmpConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDoAddComponentRequest() (request *DoAddComponentRequest) {
	request = &DoAddComponentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoAddComponent")
	return
}

func NewDoAddComponentResponse() (response *DoAddComponentResponse) {
	response = &DoAddComponentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 组件新增
func (c *Client) DoAddComponent(request *DoAddComponentRequest) (response *DoAddComponentResponse, err error) {
	if request == nil {
		request = NewDoAddComponentRequest()
	}
	response = NewDoAddComponentResponse()
	err = c.Send(request, response)
	return
}

func NewGetHostCountRequest() (request *GetHostCountRequest) {
	request = &GetHostCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetHostCount")
	return
}

func NewGetHostCountResponse() (response *GetHostCountResponse) {
	response = &GetHostCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群节点数量
func (c *Client) GetHostCount(request *GetHostCountRequest) (response *GetHostCountResponse, err error) {
	if request == nil {
		request = NewGetHostCountRequest()
	}
	response = NewGetHostCountResponse()
	err = c.Send(request, response)
	return
}

func NewSetMailAlertConfigRequest() (request *SetMailAlertConfigRequest) {
	request = &SetMailAlertConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "SetMailAlertConfig")
	return
}

func NewSetMailAlertConfigResponse() (response *SetMailAlertConfigResponse) {
	response = &SetMailAlertConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置集群邮件告警配置
func (c *Client) SetMailAlertConfig(request *SetMailAlertConfigRequest) (response *SetMailAlertConfigResponse, err error) {
	if request == nil {
		request = NewSetMailAlertConfigRequest()
	}
	response = NewSetMailAlertConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetAlertCountRequest() (request *GetAlertCountRequest) {
	request = &GetAlertCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetAlertCount")
	return
}

func NewGetAlertCountResponse() (response *GetAlertCountResponse) {
	response = &GetAlertCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取告警数量
func (c *Client) GetAlertCount(request *GetAlertCountRequest) (response *GetAlertCountResponse, err error) {
	if request == nil {
		request = NewGetAlertCountRequest()
	}
	response = NewGetAlertCountResponse()
	err = c.Send(request, response)
	return
}

func NewGetClusterOverviewRequest() (request *GetClusterOverviewRequest) {
	request = &GetClusterOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetClusterOverview")
	return
}

func NewGetClusterOverviewResponse() (response *GetClusterOverviewResponse) {
	response = &GetClusterOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群概览数据
func (c *Client) GetClusterOverview(request *GetClusterOverviewRequest) (response *GetClusterOverviewResponse, err error) {
	if request == nil {
		request = NewGetClusterOverviewRequest()
	}
	response = NewGetClusterOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewGetEvacuationProgressRequest() (request *GetEvacuationProgressRequest) {
	request = &GetEvacuationProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetEvacuationProgress")
	return
}

func NewGetEvacuationProgressResponse() (response *GetEvacuationProgressResponse) {
	response = &GetEvacuationProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取遣散节点进度
func (c *Client) GetEvacuationProgress(request *GetEvacuationProgressRequest) (response *GetEvacuationProgressResponse, err error) {
	if request == nil {
		request = NewGetEvacuationProgressRequest()
	}
	response = NewGetEvacuationProgressResponse()
	err = c.Send(request, response)
	return
}

func NewDoExpandPoolRequest() (request *DoExpandPoolRequest) {
	request = &DoExpandPoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoExpandPool")
	return
}

func NewDoExpandPoolResponse() (response *DoExpandPoolResponse) {
	response = &DoExpandPoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 扩容存储池
func (c *Client) DoExpandPool(request *DoExpandPoolRequest) (response *DoExpandPoolResponse, err error) {
	if request == nil {
		request = NewDoExpandPoolRequest()
	}
	response = NewDoExpandPoolResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductDataReplicationStateRequest() (request *QueryProductDataReplicationStateRequest) {
	request = &QueryProductDataReplicationStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "QueryProductDataReplicationState")
	return
}

func NewQueryProductDataReplicationStateResponse() (response *QueryProductDataReplicationStateResponse) {
	response = &QueryProductDataReplicationStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容灾系统产品数据同步状态查询接口
func (c *Client) QueryProductDataReplicationState(request *QueryProductDataReplicationStateRequest) (response *QueryProductDataReplicationStateResponse, err error) {
	if request == nil {
		request = NewQueryProductDataReplicationStateRequest()
	}
	response = NewQueryProductDataReplicationStateResponse()
	err = c.Send(request, response)
	return
}

func NewGetHostIfListRequest() (request *GetHostIfListRequest) {
	request = &GetHostIfListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetHostIfList")
	return
}

func NewGetHostIfListResponse() (response *GetHostIfListResponse) {
	response = &GetHostIfListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取节点网卡信息
func (c *Client) GetHostIfList(request *GetHostIfListRequest) (response *GetHostIfListResponse, err error) {
	if request == nil {
		request = NewGetHostIfListRequest()
	}
	response = NewGetHostIfListResponse()
	err = c.Send(request, response)
	return
}

func NewGetPoolHostsRequest() (request *GetPoolHostsRequest) {
	request = &GetPoolHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetPoolHosts")
	return
}

func NewGetPoolHostsResponse() (response *GetPoolHostsResponse) {
	response = &GetPoolHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取存储池节点列表
func (c *Client) GetPoolHosts(request *GetPoolHostsRequest) (response *GetPoolHostsResponse, err error) {
	if request == nil {
		request = NewGetPoolHostsRequest()
	}
	response = NewGetPoolHostsResponse()
	err = c.Send(request, response)
	return
}

func NewGetOperationLogsRequest() (request *GetOperationLogsRequest) {
	request = &GetOperationLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetOperationLogs")
	return
}

func NewGetOperationLogsResponse() (response *GetOperationLogsResponse) {
	response = &GetOperationLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群操作日志
func (c *Client) GetOperationLogs(request *GetOperationLogsRequest) (response *GetOperationLogsResponse, err error) {
	if request == nil {
		request = NewGetOperationLogsRequest()
	}
	response = NewGetOperationLogsResponse()
	err = c.Send(request, response)
	return
}

func NewSetBackupGlobalConfigRequest() (request *SetBackupGlobalConfigRequest) {
	request = &SetBackupGlobalConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "SetBackupGlobalConfig")
	return
}

func NewSetBackupGlobalConfigResponse() (response *SetBackupGlobalConfigResponse) {
	response = &SetBackupGlobalConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置存储桶备份全局配置
func (c *Client) SetBackupGlobalConfig(request *SetBackupGlobalConfigRequest) (response *SetBackupGlobalConfigResponse, err error) {
	if request == nil {
		request = NewSetBackupGlobalConfigRequest()
	}
	response = NewSetBackupGlobalConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetPoolLatestOverviewRequest() (request *GetPoolLatestOverviewRequest) {
	request = &GetPoolLatestOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetPoolLatestOverview")
	return
}

func NewGetPoolLatestOverviewResponse() (response *GetPoolLatestOverviewResponse) {
	response = &GetPoolLatestOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取存储池最新概览数据
func (c *Client) GetPoolLatestOverview(request *GetPoolLatestOverviewRequest) (response *GetPoolLatestOverviewResponse, err error) {
	if request == nil {
		request = NewGetPoolLatestOverviewRequest()
	}
	response = NewGetPoolLatestOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewGetDnsDomainSettingsRequest() (request *GetDnsDomainSettingsRequest) {
	request = &GetDnsDomainSettingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetDnsDomainSettings")
	return
}

func NewGetDnsDomainSettingsResponse() (response *GetDnsDomainSettingsResponse) {
	response = &GetDnsDomainSettingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取对象存储接入域名配置
func (c *Client) GetDnsDomainSettings(request *GetDnsDomainSettingsRequest) (response *GetDnsDomainSettingsResponse, err error) {
	if request == nil {
		request = NewGetDnsDomainSettingsRequest()
	}
	response = NewGetDnsDomainSettingsResponse()
	err = c.Send(request, response)
	return
}

func NewDoStopHostRequest() (request *DoStopHostRequest) {
	request = &DoStopHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoStopHost")
	return
}

func NewDoStopHostResponse() (response *DoStopHostResponse) {
	response = &DoStopHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 禁用节点
func (c *Client) DoStopHost(request *DoStopHostRequest) (response *DoStopHostResponse, err error) {
	if request == nil {
		request = NewDoStopHostRequest()
	}
	response = NewDoStopHostResponse()
	err = c.Send(request, response)
	return
}

func NewDoReInitializeDiskRequest() (request *DoReInitializeDiskRequest) {
	request = &DoReInitializeDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoReInitializeDisk")
	return
}

func NewDoReInitializeDiskResponse() (response *DoReInitializeDiskResponse) {
	response = &DoReInitializeDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 上线硬盘
func (c *Client) DoReInitializeDisk(request *DoReInitializeDiskRequest) (response *DoReInitializeDiskResponse, err error) {
	if request == nil {
		request = NewDoReInitializeDiskRequest()
	}
	response = NewDoReInitializeDiskResponse()
	err = c.Send(request, response)
	return
}

func NewGetRequestListRequest() (request *GetRequestListRequest) {
	request = &GetRequestListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetRequestList")
	return
}

func NewGetRequestListResponse() (response *GetRequestListResponse) {
	response = &GetRequestListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取长任务列表
func (c *Client) GetRequestList(request *GetRequestListRequest) (response *GetRequestListResponse, err error) {
	if request == nil {
		request = NewGetRequestListRequest()
	}
	response = NewGetRequestListResponse()
	err = c.Send(request, response)
	return
}

func NewSetAutoRecoveryRequest() (request *SetAutoRecoveryRequest) {
	request = &SetAutoRecoveryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "SetAutoRecovery")
	return
}

func NewSetAutoRecoveryResponse() (response *SetAutoRecoveryResponse) {
	response = &SetAutoRecoveryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置服务自启动状态
func (c *Client) SetAutoRecovery(request *SetAutoRecoveryRequest) (response *SetAutoRecoveryResponse, err error) {
	if request == nil {
		request = NewSetAutoRecoveryRequest()
	}
	response = NewSetAutoRecoveryResponse()
	err = c.Send(request, response)
	return
}

func NewGetBackupGlobalConfigRequest() (request *GetBackupGlobalConfigRequest) {
	request = &GetBackupGlobalConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetBackupGlobalConfig")
	return
}

func NewGetBackupGlobalConfigResponse() (response *GetBackupGlobalConfigResponse) {
	response = &GetBackupGlobalConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取存储桶备份全局配置
func (c *Client) GetBackupGlobalConfig(request *GetBackupGlobalConfigRequest) (response *GetBackupGlobalConfigResponse, err error) {
	if request == nil {
		request = NewGetBackupGlobalConfigRequest()
	}
	response = NewGetBackupGlobalConfigResponse()
	err = c.Send(request, response)
	return
}

func NewSetDnsDomainSettingsRequest() (request *SetDnsDomainSettingsRequest) {
	request = &SetDnsDomainSettingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "SetDnsDomainSettings")
	return
}

func NewSetDnsDomainSettingsResponse() (response *SetDnsDomainSettingsResponse) {
	response = &SetDnsDomainSettingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置对象存储接入域名配置
func (c *Client) SetDnsDomainSettings(request *SetDnsDomainSettingsRequest) (response *SetDnsDomainSettingsResponse, err error) {
	if request == nil {
		request = NewSetDnsDomainSettingsRequest()
	}
	response = NewSetDnsDomainSettingsResponse()
	err = c.Send(request, response)
	return
}

func NewDoRemoveComponentRequest() (request *DoRemoveComponentRequest) {
	request = &DoRemoveComponentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoRemoveComponent")
	return
}

func NewDoRemoveComponentResponse() (response *DoRemoveComponentResponse) {
	response = &DoRemoveComponentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 组件删除
func (c *Client) DoRemoveComponent(request *DoRemoveComponentRequest) (response *DoRemoveComponentResponse, err error) {
	if request == nil {
		request = NewDoRemoveComponentRequest()
	}
	response = NewDoRemoveComponentResponse()
	err = c.Send(request, response)
	return
}

func NewGetRgwLogLevelRequest() (request *GetRgwLogLevelRequest) {
	request = &GetRgwLogLevelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetRgwLogLevel")
	return
}

func NewGetRgwLogLevelResponse() (response *GetRgwLogLevelResponse) {
	response = &GetRgwLogLevelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取存储网关获取日志级别
func (c *Client) GetRgwLogLevel(request *GetRgwLogLevelRequest) (response *GetRgwLogLevelResponse, err error) {
	if request == nil {
		request = NewGetRgwLogLevelRequest()
	}
	response = NewGetRgwLogLevelResponse()
	err = c.Send(request, response)
	return
}

func NewGetRgwLogListRequest() (request *GetRgwLogListRequest) {
	request = &GetRgwLogListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetRgwLogList")
	return
}

func NewGetRgwLogListResponse() (response *GetRgwLogListResponse) {
	response = &GetRgwLogListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志列表
func (c *Client) GetRgwLogList(request *GetRgwLogListRequest) (response *GetRgwLogListResponse, err error) {
	if request == nil {
		request = NewGetRgwLogListRequest()
	}
	response = NewGetRgwLogListResponse()
	err = c.Send(request, response)
	return
}

func NewSetFlowControlConfigRequest() (request *SetFlowControlConfigRequest) {
	request = &SetFlowControlConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "SetFlowControlConfig")
	return
}

func NewSetFlowControlConfigResponse() (response *SetFlowControlConfigResponse) {
	response = &SetFlowControlConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置对象存储流控设置
func (c *Client) SetFlowControlConfig(request *SetFlowControlConfigRequest) (response *SetFlowControlConfigResponse, err error) {
	if request == nil {
		request = NewSetFlowControlConfigRequest()
	}
	response = NewSetFlowControlConfigResponse()
	err = c.Send(request, response)
	return
}

func NewSetObjectStorageBucketQuotaRequest() (request *SetObjectStorageBucketQuotaRequest) {
	request = &SetObjectStorageBucketQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "SetObjectStorageBucketQuota")
	return
}

func NewSetObjectStorageBucketQuotaResponse() (response *SetObjectStorageBucketQuotaResponse) {
	response = &SetObjectStorageBucketQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置存储桶配额
func (c *Client) SetObjectStorageBucketQuota(request *SetObjectStorageBucketQuotaRequest) (response *SetObjectStorageBucketQuotaResponse, err error) {
	if request == nil {
		request = NewSetObjectStorageBucketQuotaRequest()
	}
	response = NewSetObjectStorageBucketQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDoStopComponentRequest() (request *DoStopComponentRequest) {
	request = &DoStopComponentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoStopComponent")
	return
}

func NewDoStopComponentResponse() (response *DoStopComponentResponse) {
	response = &DoStopComponentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 组件禁用
func (c *Client) DoStopComponent(request *DoStopComponentRequest) (response *DoStopComponentResponse, err error) {
	if request == nil {
		request = NewDoStopComponentRequest()
	}
	response = NewDoStopComponentResponse()
	err = c.Send(request, response)
	return
}

func NewGetHostComponentStateListRequest() (request *GetHostComponentStateListRequest) {
	request = &GetHostComponentStateListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetHostComponentStateList")
	return
}

func NewGetHostComponentStateListResponse() (response *GetHostComponentStateListResponse) {
	response = &GetHostComponentStateListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取节点上组件状态
func (c *Client) GetHostComponentStateList(request *GetHostComponentStateListRequest) (response *GetHostComponentStateListResponse, err error) {
	if request == nil {
		request = NewGetHostComponentStateListRequest()
	}
	response = NewGetHostComponentStateListResponse()
	err = c.Send(request, response)
	return
}

func NewRetryRequestByIdRequest() (request *RetryRequestByIdRequest) {
	request = &RetryRequestByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "RetryRequestById")
	return
}

func NewRetryRequestByIdResponse() (response *RetryRequestByIdResponse) {
	response = &RetryRequestByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重试长异步请求
func (c *Client) RetryRequestById(request *RetryRequestByIdRequest) (response *RetryRequestByIdResponse, err error) {
	if request == nil {
		request = NewRetryRequestByIdRequest()
	}
	response = NewRetryRequestByIdResponse()
	err = c.Send(request, response)
	return
}

func NewDoDeleteHostRequest() (request *DoDeleteHostRequest) {
	request = &DoDeleteHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoDeleteHost")
	return
}

func NewDoDeleteHostResponse() (response *DoDeleteHostResponse) {
	response = &DoDeleteHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除节点
func (c *Client) DoDeleteHost(request *DoDeleteHostRequest) (response *DoDeleteHostResponse, err error) {
	if request == nil {
		request = NewDoDeleteHostRequest()
	}
	response = NewDoDeleteHostResponse()
	err = c.Send(request, response)
	return
}

func NewGetHostCandidateRequest() (request *GetHostCandidateRequest) {
	request = &GetHostCandidateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetHostCandidate")
	return
}

func NewGetHostCandidateResponse() (response *GetHostCandidateResponse) {
	response = &GetHostCandidateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取可用节点列表
func (c *Client) GetHostCandidate(request *GetHostCandidateRequest) (response *GetHostCandidateResponse, err error) {
	if request == nil {
		request = NewGetHostCandidateRequest()
	}
	response = NewGetHostCandidateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductFailureMigrateRequest() (request *CreateProductFailureMigrateRequest) {
	request = &CreateProductFailureMigrateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "CreateProductFailureMigrate")
	return
}

func NewCreateProductFailureMigrateResponse() (response *CreateProductFailureMigrateResponse) {
	response = &CreateProductFailureMigrateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容灾系统产品故障切换操作接口
func (c *Client) CreateProductFailureMigrate(request *CreateProductFailureMigrateRequest) (response *CreateProductFailureMigrateResponse, err error) {
	if request == nil {
		request = NewCreateProductFailureMigrateRequest()
	}
	response = NewCreateProductFailureMigrateResponse()
	err = c.Send(request, response)
	return
}

func NewGetBkcmdbAccessHostRequest() (request *GetBkcmdbAccessHostRequest) {
	request = &GetBkcmdbAccessHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetBkcmdbAccessHost")
	return
}

func NewGetBkcmdbAccessHostResponse() (response *GetBkcmdbAccessHostResponse) {
	response = &GetBkcmdbAccessHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [已废弃]获取bkcmdb中的存储网关节点列表
func (c *Client) GetBkcmdbAccessHost(request *GetBkcmdbAccessHostRequest) (response *GetBkcmdbAccessHostResponse, err error) {
	if request == nil {
		request = NewGetBkcmdbAccessHostRequest()
	}
	response = NewGetBkcmdbAccessHostResponse()
	err = c.Send(request, response)
	return
}

func NewGetRequestInfoByIdRequest() (request *GetRequestInfoByIdRequest) {
	request = &GetRequestInfoByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetRequestInfoById")
	return
}

func NewGetRequestInfoByIdResponse() (response *GetRequestInfoByIdResponse) {
	response = &GetRequestInfoByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取CSP任务的子任务列表详情
func (c *Client) GetRequestInfoById(request *GetRequestInfoByIdRequest) (response *GetRequestInfoByIdResponse, err error) {
	if request == nil {
		request = NewGetRequestInfoByIdRequest()
	}
	response = NewGetRequestInfoByIdResponse()
	err = c.Send(request, response)
	return
}

func NewDoEvacuateComponentRequest() (request *DoEvacuateComponentRequest) {
	request = &DoEvacuateComponentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoEvacuateComponent")
	return
}

func NewDoEvacuateComponentResponse() (response *DoEvacuateComponentResponse) {
	response = &DoEvacuateComponentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 节点遣散
func (c *Client) DoEvacuateComponent(request *DoEvacuateComponentRequest) (response *DoEvacuateComponentResponse, err error) {
	if request == nil {
		request = NewDoEvacuateComponentRequest()
	}
	response = NewDoEvacuateComponentResponse()
	err = c.Send(request, response)
	return
}

func NewGetHttpQpsRequest() (request *GetHttpQpsRequest) {
	request = &GetHttpQpsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetHttpQps")
	return
}

func NewGetHttpQpsResponse() (response *GetHttpQpsResponse) {
	response = &GetHttpQpsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群Http QPS 数据
func (c *Client) GetHttpQps(request *GetHttpQpsRequest) (response *GetHttpQpsResponse, err error) {
	if request == nil {
		request = NewGetHttpQpsRequest()
	}
	response = NewGetHttpQpsResponse()
	err = c.Send(request, response)
	return
}

func NewGetRecoveryExpansionRequest() (request *GetRecoveryExpansionRequest) {
	request = &GetRecoveryExpansionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetRecoveryExpansion")
	return
}

func NewGetRecoveryExpansionResponse() (response *GetRecoveryExpansionResponse) {
	response = &GetRecoveryExpansionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取存储池扩容、修复数据进度
func (c *Client) GetRecoveryExpansion(request *GetRecoveryExpansionRequest) (response *GetRecoveryExpansionResponse, err error) {
	if request == nil {
		request = NewGetRecoveryExpansionRequest()
	}
	response = NewGetRecoveryExpansionResponse()
	err = c.Send(request, response)
	return
}

func NewGetClusterLatestOverviewRequest() (request *GetClusterLatestOverviewRequest) {
	request = &GetClusterLatestOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetClusterLatestOverview")
	return
}

func NewGetClusterLatestOverviewResponse() (response *GetClusterLatestOverviewResponse) {
	response = &GetClusterLatestOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群最新概览数据
func (c *Client) GetClusterLatestOverview(request *GetClusterLatestOverviewRequest) (response *GetClusterLatestOverviewResponse, err error) {
	if request == nil {
		request = NewGetClusterLatestOverviewRequest()
	}
	response = NewGetClusterLatestOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
	request = &CreateClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "CreateCluster")
	return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
	response = &CreateClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 执行创建集群
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
	if request == nil {
		request = NewCreateClusterRequest()
	}
	response = NewCreateClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadHostDisksDataRequest() (request *DownloadHostDisksDataRequest) {
	request = &DownloadHostDisksDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DownloadHostDisksData")
	return
}

func NewDownloadHostDisksDataResponse() (response *DownloadHostDisksDataResponse) {
	response = &DownloadHostDisksDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载节点上磁盘读写带宽、读写iops
func (c *Client) DownloadHostDisksData(request *DownloadHostDisksDataRequest) (response *DownloadHostDisksDataResponse, err error) {
	if request == nil {
		request = NewDownloadHostDisksDataRequest()
	}
	response = NewDownloadHostDisksDataResponse()
	err = c.Send(request, response)
	return
}

func NewGetHostOverviewRequest() (request *GetHostOverviewRequest) {
	request = &GetHostOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetHostOverview")
	return
}

func NewGetHostOverviewResponse() (response *GetHostOverviewResponse) {
	response = &GetHostOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [已废弃]获取节点概览信息
func (c *Client) GetHostOverview(request *GetHostOverviewRequest) (response *GetHostOverviewResponse, err error) {
	if request == nil {
		request = NewGetHostOverviewRequest()
	}
	response = NewGetHostOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewGetRequestByIdRequest() (request *GetRequestByIdRequest) {
	request = &GetRequestByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetRequestById")
	return
}

func NewGetRequestByIdResponse() (response *GetRequestByIdResponse) {
	response = &GetRequestByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取CSP任务详情
func (c *Client) GetRequestById(request *GetRequestByIdRequest) (response *GetRequestByIdResponse, err error) {
	if request == nil {
		request = NewGetRequestByIdRequest()
	}
	response = NewGetRequestByIdResponse()
	err = c.Send(request, response)
	return
}

func NewQueryPrometheusRequest() (request *QueryPrometheusRequest) {
	request = &QueryPrometheusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "QueryPrometheus")
	return
}

func NewQueryPrometheusResponse() (response *QueryPrometheusResponse) {
	response = &QueryPrometheusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询普罗米修斯
func (c *Client) QueryPrometheus(request *QueryPrometheusRequest) (response *QueryPrometheusResponse, err error) {
	if request == nil {
		request = NewQueryPrometheusRequest()
	}
	response = NewQueryPrometheusResponse()
	err = c.Send(request, response)
	return
}

func NewDoMigrateComponentRequest() (request *DoMigrateComponentRequest) {
	request = &DoMigrateComponentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoMigrateComponent")
	return
}

func NewDoMigrateComponentResponse() (response *DoMigrateComponentResponse) {
	response = &DoMigrateComponentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 节点更换
func (c *Client) DoMigrateComponent(request *DoMigrateComponentRequest) (response *DoMigrateComponentResponse, err error) {
	if request == nil {
		request = NewDoMigrateComponentRequest()
	}
	response = NewDoMigrateComponentResponse()
	err = c.Send(request, response)
	return
}

func NewGetClusterConstStatusRequest() (request *GetClusterConstStatusRequest) {
	request = &GetClusterConstStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetClusterConstStatus")
	return
}

func NewGetClusterConstStatusResponse() (response *GetClusterConstStatusResponse) {
	response = &GetClusterConstStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集群获取常量信息
func (c *Client) GetClusterConstStatus(request *GetClusterConstStatusRequest) (response *GetClusterConstStatusResponse, err error) {
	if request == nil {
		request = NewGetClusterConstStatusRequest()
	}
	response = NewGetClusterConstStatusResponse()
	err = c.Send(request, response)
	return
}

func NewGetHttpCountRequest() (request *GetHttpCountRequest) {
	request = &GetHttpCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetHttpCount")
	return
}

func NewGetHttpCountResponse() (response *GetHttpCountResponse) {
	response = &GetHttpCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群Http统计数据
func (c *Client) GetHttpCount(request *GetHttpCountRequest) (response *GetHttpCountResponse, err error) {
	if request == nil {
		request = NewGetHttpCountRequest()
	}
	response = NewGetHttpCountResponse()
	err = c.Send(request, response)
	return
}

func NewSetRecoveryControlConfigRequest() (request *SetRecoveryControlConfigRequest) {
	request = &SetRecoveryControlConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "SetRecoveryControlConfig")
	return
}

func NewSetRecoveryControlConfigResponse() (response *SetRecoveryControlConfigResponse) {
	response = &SetRecoveryControlConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置集群恢复控制设置
func (c *Client) SetRecoveryControlConfig(request *SetRecoveryControlConfigRequest) (response *SetRecoveryControlConfigResponse, err error) {
	if request == nil {
		request = NewSetRecoveryControlConfigRequest()
	}
	response = NewSetRecoveryControlConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetFlowControlConfigRequest() (request *GetFlowControlConfigRequest) {
	request = &GetFlowControlConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetFlowControlConfig")
	return
}

func NewGetFlowControlConfigResponse() (response *GetFlowControlConfigResponse) {
	response = &GetFlowControlConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取对象存储频控设置
func (c *Client) GetFlowControlConfig(request *GetFlowControlConfigRequest) (response *GetFlowControlConfigResponse, err error) {
	if request == nil {
		request = NewGetFlowControlConfigRequest()
	}
	response = NewGetFlowControlConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetGlobalInfoRequest() (request *GetGlobalInfoRequest) {
	request = &GetGlobalInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetGlobalInfo")
	return
}

func NewGetGlobalInfoResponse() (response *GetGlobalInfoResponse) {
	response = &GetGlobalInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取全局参数
func (c *Client) GetGlobalInfo(request *GetGlobalInfoRequest) (response *GetGlobalInfoResponse, err error) {
	if request == nil {
		request = NewGetGlobalInfoRequest()
	}
	response = NewGetGlobalInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGetHostListRequest() (request *GetHostListRequest) {
	request = &GetHostListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetHostList")
	return
}

func NewGetHostListResponse() (response *GetHostListResponse) {
	response = &GetHostListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取节点列表
func (c *Client) GetHostList(request *GetHostListRequest) (response *GetHostListResponse, err error) {
	if request == nil {
		request = NewGetHostListRequest()
	}
	response = NewGetHostListResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductStateInfoRequest() (request *QueryProductStateInfoRequest) {
	request = &QueryProductStateInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "QueryProductStateInfo")
	return
}

func NewQueryProductStateInfoResponse() (response *QueryProductStateInfoResponse) {
	response = &QueryProductStateInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容灾系统产品状态信息查询接口
// https://iwiki.woa.com/pages/viewpage.action?pageId=377897543#id-接口设计-产品状态信息查询接口
func (c *Client) QueryProductStateInfo(request *QueryProductStateInfoRequest) (response *QueryProductStateInfoResponse, err error) {
	if request == nil {
		request = NewQueryProductStateInfoRequest()
	}
	response = NewQueryProductStateInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGetAutoRecoveryRequest() (request *GetAutoRecoveryRequest) {
	request = &GetAutoRecoveryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetAutoRecovery")
	return
}

func NewGetAutoRecoveryResponse() (response *GetAutoRecoveryResponse) {
	response = &GetAutoRecoveryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取服务自启动状态
func (c *Client) GetAutoRecovery(request *GetAutoRecoveryRequest) (response *GetAutoRecoveryResponse, err error) {
	if request == nil {
		request = NewGetAutoRecoveryRequest()
	}
	response = NewGetAutoRecoveryResponse()
	err = c.Send(request, response)
	return
}

func NewGetClusterStatusRequest() (request *GetClusterStatusRequest) {
	request = &GetClusterStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetClusterStatus")
	return
}

func NewGetClusterStatusResponse() (response *GetClusterStatusResponse) {
	response = &GetClusterStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集群获取信息
func (c *Client) GetClusterStatus(request *GetClusterStatusRequest) (response *GetClusterStatusResponse, err error) {
	if request == nil {
		request = NewGetClusterStatusRequest()
	}
	response = NewGetClusterStatusResponse()
	err = c.Send(request, response)
	return
}

func NewGetBkcmdbAllHostRequest() (request *GetBkcmdbAllHostRequest) {
	request = &GetBkcmdbAllHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetBkcmdbAllHost")
	return
}

func NewGetBkcmdbAllHostResponse() (response *GetBkcmdbAllHostResponse) {
	response = &GetBkcmdbAllHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bkcmdb中的节点列表
func (c *Client) GetBkcmdbAllHost(request *GetBkcmdbAllHostRequest) (response *GetBkcmdbAllHostResponse, err error) {
	if request == nil {
		request = NewGetBkcmdbAllHostRequest()
	}
	response = NewGetBkcmdbAllHostResponse()
	err = c.Send(request, response)
	return
}

func NewSetDefaultPlacementRequest() (request *SetDefaultPlacementRequest) {
	request = &SetDefaultPlacementRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "SetDefaultPlacement")
	return
}

func NewSetDefaultPlacementResponse() (response *SetDefaultPlacementResponse) {
	response = &SetDefaultPlacementResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置默认存储池
func (c *Client) SetDefaultPlacement(request *SetDefaultPlacementRequest) (response *SetDefaultPlacementResponse, err error) {
	if request == nil {
		request = NewSetDefaultPlacementRequest()
	}
	response = NewSetDefaultPlacementResponse()
	err = c.Send(request, response)
	return
}

func NewSetBucketFlowControlRequest() (request *SetBucketFlowControlRequest) {
	request = &SetBucketFlowControlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "SetBucketFlowControl")
	return
}

func NewSetBucketFlowControlResponse() (response *SetBucketFlowControlResponse) {
	response = &SetBucketFlowControlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置存储桶频控配置
func (c *Client) SetBucketFlowControl(request *SetBucketFlowControlRequest) (response *SetBucketFlowControlResponse, err error) {
	if request == nil {
		request = NewSetBucketFlowControlRequest()
	}
	response = NewSetBucketFlowControlResponse()
	err = c.Send(request, response)
	return
}

func NewGetBkcmdbMonitorHostRequest() (request *GetBkcmdbMonitorHostRequest) {
	request = &GetBkcmdbMonitorHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetBkcmdbMonitorHost")
	return
}

func NewGetBkcmdbMonitorHostResponse() (response *GetBkcmdbMonitorHostResponse) {
	response = &GetBkcmdbMonitorHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bkcmdb的监控节点列表
func (c *Client) GetBkcmdbMonitorHost(request *GetBkcmdbMonitorHostRequest) (response *GetBkcmdbMonitorHostResponse, err error) {
	if request == nil {
		request = NewGetBkcmdbMonitorHostRequest()
	}
	response = NewGetBkcmdbMonitorHostResponse()
	err = c.Send(request, response)
	return
}

func NewGetInspectionStatusRequest() (request *GetInspectionStatusRequest) {
	request = &GetInspectionStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetInspectionStatus")
	return
}

func NewGetInspectionStatusResponse() (response *GetInspectionStatusResponse) {
	response = &GetInspectionStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群巡检信息
func (c *Client) GetInspectionStatus(request *GetInspectionStatusRequest) (response *GetInspectionStatusResponse, err error) {
	if request == nil {
		request = NewGetInspectionStatusRequest()
	}
	response = NewGetInspectionStatusResponse()
	err = c.Send(request, response)
	return
}

func NewRetryInitClusterRequest() (request *RetryInitClusterRequest) {
	request = &RetryInitClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "RetryInitCluster")
	return
}

func NewRetryInitClusterResponse() (response *RetryInitClusterResponse) {
	response = &RetryInitClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重试初始化集群
func (c *Client) RetryInitCluster(request *RetryInitClusterRequest) (response *RetryInitClusterResponse, err error) {
	if request == nil {
		request = NewRetryInitClusterRequest()
	}
	response = NewRetryInitClusterResponse()
	err = c.Send(request, response)
	return
}

func NewStartInspectionRequest() (request *StartInspectionRequest) {
	request = &StartInspectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "StartInspection")
	return
}

func NewStartInspectionResponse() (response *StartInspectionResponse) {
	response = &StartInspectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动巡检
func (c *Client) StartInspection(request *StartInspectionRequest) (response *StartInspectionResponse, err error) {
	if request == nil {
		request = NewStartInspectionRequest()
	}
	response = NewStartInspectionResponse()
	err = c.Send(request, response)
	return
}

func NewDoAddHostsRequest() (request *DoAddHostsRequest) {
	request = &DoAddHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoAddHosts")
	return
}

func NewDoAddHostsResponse() (response *DoAddHostsResponse) {
	response = &DoAddHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增节点
func (c *Client) DoAddHosts(request *DoAddHostsRequest) (response *DoAddHostsResponse, err error) {
	if request == nil {
		request = NewDoAddHostsRequest()
	}
	response = NewDoAddHostsResponse()
	err = c.Send(request, response)
	return
}

func NewGetHostDiskOverviewRequest() (request *GetHostDiskOverviewRequest) {
	request = &GetHostDiskOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetHostDiskOverview")
	return
}

func NewGetHostDiskOverviewResponse() (response *GetHostDiskOverviewResponse) {
	response = &GetHostDiskOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// [已废弃]获取节点上磁盘数据
func (c *Client) GetHostDiskOverview(request *GetHostDiskOverviewRequest) (response *GetHostDiskOverviewResponse, err error) {
	if request == nil {
		request = NewGetHostDiskOverviewRequest()
	}
	response = NewGetHostDiskOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewGetObjectStoragePoolListRequest() (request *GetObjectStoragePoolListRequest) {
	request = &GetObjectStoragePoolListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetObjectStoragePoolList")
	return
}

func NewGetObjectStoragePoolListResponse() (response *GetObjectStoragePoolListResponse) {
	response = &GetObjectStoragePoolListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取对象存储池列表
func (c *Client) GetObjectStoragePoolList(request *GetObjectStoragePoolListRequest) (response *GetObjectStoragePoolListResponse, err error) {
	if request == nil {
		request = NewGetObjectStoragePoolListRequest()
	}
	response = NewGetObjectStoragePoolListResponse()
	err = c.Send(request, response)
	return
}

func NewSetClusterConfigRequest() (request *SetClusterConfigRequest) {
	request = &SetClusterConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "SetClusterConfig")
	return
}

func NewSetClusterConfigResponse() (response *SetClusterConfigResponse) {
	response = &SetClusterConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置集群配置
func (c *Client) SetClusterConfig(request *SetClusterConfigRequest) (response *SetClusterConfigResponse, err error) {
	if request == nil {
		request = NewSetClusterConfigRequest()
	}
	response = NewSetClusterConfigResponse()
	err = c.Send(request, response)
	return
}

func NewSetRgwLogLevelRequest() (request *SetRgwLogLevelRequest) {
	request = &SetRgwLogLevelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "SetRgwLogLevel")
	return
}

func NewSetRgwLogLevelResponse() (response *SetRgwLogLevelResponse) {
	response = &SetRgwLogLevelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置存储网关日志等级
func (c *Client) SetRgwLogLevel(request *SetRgwLogLevelRequest) (response *SetRgwLogLevelResponse, err error) {
	if request == nil {
		request = NewSetRgwLogLevelRequest()
	}
	response = NewSetRgwLogLevelResponse()
	err = c.Send(request, response)
	return
}

func NewDoLocateDiskRequest() (request *DoLocateDiskRequest) {
	request = &DoLocateDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoLocateDisk")
	return
}

func NewDoLocateDiskResponse() (response *DoLocateDiskResponse) {
	response = &DoLocateDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 硬盘点灯
func (c *Client) DoLocateDisk(request *DoLocateDiskRequest) (response *DoLocateDiskResponse, err error) {
	if request == nil {
		request = NewDoLocateDiskRequest()
	}
	response = NewDoLocateDiskResponse()
	err = c.Send(request, response)
	return
}

func NewSetSnmpConfigRequest() (request *SetSnmpConfigRequest) {
	request = &SetSnmpConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "SetSnmpConfig")
	return
}

func NewSetSnmpConfigResponse() (response *SetSnmpConfigResponse) {
	response = &SetSnmpConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置集群SNMP配置
func (c *Client) SetSnmpConfig(request *SetSnmpConfigRequest) (response *SetSnmpConfigResponse, err error) {
	if request == nil {
		request = NewSetSnmpConfigRequest()
	}
	response = NewSetSnmpConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetObjectStorageUserQuotaRequest() (request *GetObjectStorageUserQuotaRequest) {
	request = &GetObjectStorageUserQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetObjectStorageUserQuota")
	return
}

func NewGetObjectStorageUserQuotaResponse() (response *GetObjectStorageUserQuotaResponse) {
	response = &GetObjectStorageUserQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询对象存储用户配额
func (c *Client) GetObjectStorageUserQuota(request *GetObjectStorageUserQuotaRequest) (response *GetObjectStorageUserQuotaResponse, err error) {
	if request == nil {
		request = NewGetObjectStorageUserQuotaRequest()
	}
	response = NewGetObjectStorageUserQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewEnsureClusterNotInstalledRequest() (request *EnsureClusterNotInstalledRequest) {
	request = &EnsureClusterNotInstalledRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "EnsureClusterNotInstalled")
	return
}

func NewEnsureClusterNotInstalledResponse() (response *EnsureClusterNotInstalledResponse) {
	response = &EnsureClusterNotInstalledResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群是否未安装
func (c *Client) EnsureClusterNotInstalled(request *EnsureClusterNotInstalledRequest) (response *EnsureClusterNotInstalledResponse, err error) {
	if request == nil {
		request = NewEnsureClusterNotInstalledRequest()
	}
	response = NewEnsureClusterNotInstalledResponse()
	err = c.Send(request, response)
	return
}

func NewDoReleaseDiskRequest() (request *DoReleaseDiskRequest) {
	request = &DoReleaseDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoReleaseDisk")
	return
}

func NewDoReleaseDiskResponse() (response *DoReleaseDiskResponse) {
	response = &DoReleaseDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 卸载硬盘
func (c *Client) DoReleaseDisk(request *DoReleaseDiskRequest) (response *DoReleaseDiskResponse, err error) {
	if request == nil {
		request = NewDoReleaseDiskRequest()
	}
	response = NewDoReleaseDiskResponse()
	err = c.Send(request, response)
	return
}

func NewDoStartComponentRequest() (request *DoStartComponentRequest) {
	request = &DoStartComponentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoStartComponent")
	return
}

func NewDoStartComponentResponse() (response *DoStartComponentResponse) {
	response = &DoStartComponentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 组件启动
func (c *Client) DoStartComponent(request *DoStartComponentRequest) (response *DoStartComponentResponse, err error) {
	if request == nil {
		request = NewDoStartComponentRequest()
	}
	response = NewDoStartComponentResponse()
	err = c.Send(request, response)
	return
}

func NewGetClusterConfigRequest() (request *GetClusterConfigRequest) {
	request = &GetClusterConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetClusterConfig")
	return
}

func NewGetClusterConfigResponse() (response *GetClusterConfigResponse) {
	response = &GetClusterConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群配置
func (c *Client) GetClusterConfig(request *GetClusterConfigRequest) (response *GetClusterConfigResponse, err error) {
	if request == nil {
		request = NewGetClusterConfigRequest()
	}
	response = NewGetClusterConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDoRestartComponentRequest() (request *DoRestartComponentRequest) {
	request = &DoRestartComponentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoRestartComponent")
	return
}

func NewDoRestartComponentResponse() (response *DoRestartComponentResponse) {
	response = &DoRestartComponentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 组件重启
func (c *Client) DoRestartComponent(request *DoRestartComponentRequest) (response *DoRestartComponentResponse, err error) {
	if request == nil {
		request = NewDoRestartComponentRequest()
	}
	response = NewDoRestartComponentResponse()
	err = c.Send(request, response)
	return
}

func NewGetCurrentLicenseUsageRequest() (request *GetCurrentLicenseUsageRequest) {
	request = &GetCurrentLicenseUsageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetCurrentLicenseUsage")
	return
}

func NewGetCurrentLicenseUsageResponse() (response *GetCurrentLicenseUsageResponse) {
	response = &GetCurrentLicenseUsageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取当前License使用
func (c *Client) GetCurrentLicenseUsage(request *GetCurrentLicenseUsageRequest) (response *GetCurrentLicenseUsageResponse, err error) {
	if request == nil {
		request = NewGetCurrentLicenseUsageRequest()
	}
	response = NewGetCurrentLicenseUsageResponse()
	err = c.Send(request, response)
	return
}

func NewGetHostDetailRequest() (request *GetHostDetailRequest) {
	request = &GetHostDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetHostDetail")
	return
}

func NewGetHostDetailResponse() (response *GetHostDetailResponse) {
	response = &GetHostDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取节点详情信息
func (c *Client) GetHostDetail(request *GetHostDetailRequest) (response *GetHostDetailResponse, err error) {
	if request == nil {
		request = NewGetHostDetailRequest()
	}
	response = NewGetHostDetailResponse()
	err = c.Send(request, response)
	return
}

func NewGetBkcmdbStorageHostRequest() (request *GetBkcmdbStorageHostRequest) {
	request = &GetBkcmdbStorageHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetBkcmdbStorageHost")
	return
}

func NewGetBkcmdbStorageHostResponse() (response *GetBkcmdbStorageHostResponse) {
	response = &GetBkcmdbStorageHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取bkcmdb的集群节点列表
func (c *Client) GetBkcmdbStorageHost(request *GetBkcmdbStorageHostRequest) (response *GetBkcmdbStorageHostResponse, err error) {
	if request == nil {
		request = NewGetBkcmdbStorageHostRequest()
	}
	response = NewGetBkcmdbStorageHostResponse()
	err = c.Send(request, response)
	return
}

func NewGetPoolByNameRequest() (request *GetPoolByNameRequest) {
	request = &GetPoolByNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetPoolByName")
	return
}

func NewGetPoolByNameResponse() (response *GetPoolByNameResponse) {
	response = &GetPoolByNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取存储池详情
func (c *Client) GetPoolByName(request *GetPoolByNameRequest) (response *GetPoolByNameResponse, err error) {
	if request == nil {
		request = NewGetPoolByNameRequest()
	}
	response = NewGetPoolByNameResponse()
	err = c.Send(request, response)
	return
}

func NewGetOSUsageOverviewGroupByBucketRequest() (request *GetOSUsageOverviewGroupByBucketRequest) {
	request = &GetOSUsageOverviewGroupByBucketRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetOSUsageOverviewGroupByBucket")
	return
}

func NewGetOSUsageOverviewGroupByBucketResponse() (response *GetOSUsageOverviewGroupByBucketResponse) {
	response = &GetOSUsageOverviewGroupByBucketResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按存储桶名分组获取存储桶用量总览
func (c *Client) GetOSUsageOverviewGroupByBucket(request *GetOSUsageOverviewGroupByBucketRequest) (response *GetOSUsageOverviewGroupByBucketResponse, err error) {
	if request == nil {
		request = NewGetOSUsageOverviewGroupByBucketRequest()
	}
	response = NewGetOSUsageOverviewGroupByBucketResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadInspectionRequest() (request *DownloadInspectionRequest) {
	request = &DownloadInspectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DownloadInspection")
	return
}

func NewDownloadInspectionResponse() (response *DownloadInspectionResponse) {
	response = &DownloadInspectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载集群巡检信息
func (c *Client) DownloadInspection(request *DownloadInspectionRequest) (response *DownloadInspectionResponse, err error) {
	if request == nil {
		request = NewDownloadInspectionRequest()
	}
	response = NewDownloadInspectionResponse()
	err = c.Send(request, response)
	return
}

func NewGetHostRolesForDssidRequest() (request *GetHostRolesForDssidRequest) {
	request = &GetHostRolesForDssidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetHostRolesForDssid")
	return
}

func NewGetHostRolesForDssidResponse() (response *GetHostRolesForDssidResponse) {
	response = &GetHostRolesForDssidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取存储池状态相关的HostRoles
func (c *Client) GetHostRolesForDssid(request *GetHostRolesForDssidRequest) (response *GetHostRolesForDssidResponse, err error) {
	if request == nil {
		request = NewGetHostRolesForDssidRequest()
	}
	response = NewGetHostRolesForDssidResponse()
	err = c.Send(request, response)
	return
}

func NewInitClusterRequest() (request *InitClusterRequest) {
	request = &InitClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "InitCluster")
	return
}

func NewInitClusterResponse() (response *InitClusterResponse) {
	response = &InitClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 初始化集群
func (c *Client) InitCluster(request *InitClusterRequest) (response *InitClusterResponse, err error) {
	if request == nil {
		request = NewInitClusterRequest()
	}
	response = NewInitClusterResponse()
	err = c.Send(request, response)
	return
}

func NewGetMailAlertConfigRequest() (request *GetMailAlertConfigRequest) {
	request = &GetMailAlertConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetMailAlertConfig")
	return
}

func NewGetMailAlertConfigResponse() (response *GetMailAlertConfigResponse) {
	response = &GetMailAlertConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群邮件告警配置
func (c *Client) GetMailAlertConfig(request *GetMailAlertConfigRequest) (response *GetMailAlertConfigResponse, err error) {
	if request == nil {
		request = NewGetMailAlertConfigRequest()
	}
	response = NewGetMailAlertConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetOSUsageOverviewRequest() (request *GetOSUsageOverviewRequest) {
	request = &GetOSUsageOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetOSUsageOverview")
	return
}

func NewGetOSUsageOverviewResponse() (response *GetOSUsageOverviewResponse) {
	response = &GetOSUsageOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取存储桶用量总览
func (c *Client) GetOSUsageOverview(request *GetOSUsageOverviewRequest) (response *GetOSUsageOverviewResponse, err error) {
	if request == nil {
		request = NewGetOSUsageOverviewRequest()
	}
	response = NewGetOSUsageOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewGetCdcUserSuspendStatusRequest() (request *GetCdcUserSuspendStatusRequest) {
	request = &GetCdcUserSuspendStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetCdcUserSuspendStatus")
	return
}

func NewGetCdcUserSuspendStatusResponse() (response *GetCdcUserSuspendStatusResponse) {
	response = &GetCdcUserSuspendStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询CSP的CDC用户状态
func (c *Client) GetCdcUserSuspendStatus(request *GetCdcUserSuspendStatusRequest) (response *GetCdcUserSuspendStatusResponse, err error) {
	if request == nil {
		request = NewGetCdcUserSuspendStatusRequest()
	}
	response = NewGetCdcUserSuspendStatusResponse()
	err = c.Send(request, response)
	return
}

func NewGetHostComponentDssidListRequest() (request *GetHostComponentDssidListRequest) {
	request = &GetHostComponentDssidListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetHostComponentDssidList")
	return
}

func NewGetHostComponentDssidListResponse() (response *GetHostComponentDssidListResponse) {
	response = &GetHostComponentDssidListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取节点-组件存储池id对应关系列表
func (c *Client) GetHostComponentDssidList(request *GetHostComponentDssidListRequest) (response *GetHostComponentDssidListResponse, err error) {
	if request == nil {
		request = NewGetHostComponentDssidListRequest()
	}
	response = NewGetHostComponentDssidListResponse()
	err = c.Send(request, response)
	return
}

func NewSetObjectStorageUserQuotaRequest() (request *SetObjectStorageUserQuotaRequest) {
	request = &SetObjectStorageUserQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "SetObjectStorageUserQuota")
	return
}

func NewSetObjectStorageUserQuotaResponse() (response *SetObjectStorageUserQuotaResponse) {
	response = &SetObjectStorageUserQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置对象存储用户配额
func (c *Client) SetObjectStorageUserQuota(request *SetObjectStorageUserQuotaRequest) (response *SetObjectStorageUserQuotaResponse, err error) {
	if request == nil {
		request = NewSetObjectStorageUserQuotaRequest()
	}
	response = NewSetObjectStorageUserQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDoDeletePoolRequest() (request *DoDeletePoolRequest) {
	request = &DoDeletePoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoDeletePool")
	return
}

func NewDoDeletePoolResponse() (response *DoDeletePoolResponse) {
	response = &DoDeletePoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除存储池
func (c *Client) DoDeletePool(request *DoDeletePoolRequest) (response *DoDeletePoolResponse, err error) {
	if request == nil {
		request = NewDoDeletePoolRequest()
	}
	response = NewDoDeletePoolResponse()
	err = c.Send(request, response)
	return
}

func NewDoRetryPoolRequestRequest() (request *DoRetryPoolRequestRequest) {
	request = &DoRetryPoolRequestRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoRetryPoolRequest")
	return
}

func NewDoRetryPoolRequestResponse() (response *DoRetryPoolRequestResponse) {
	response = &DoRetryPoolRequestResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 再次创建存储池（出错时）
func (c *Client) DoRetryPoolRequest(request *DoRetryPoolRequestRequest) (response *DoRetryPoolRequestResponse, err error) {
	if request == nil {
		request = NewDoRetryPoolRequestRequest()
	}
	response = NewDoRetryPoolRequestResponse()
	err = c.Send(request, response)
	return
}

func NewGetAlertsRequest() (request *GetAlertsRequest) {
	request = &GetAlertsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetAlerts")
	return
}

func NewGetAlertsResponse() (response *GetAlertsResponse) {
	response = &GetAlertsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取告警列表
func (c *Client) GetAlerts(request *GetAlertsRequest) (response *GetAlertsResponse, err error) {
	if request == nil {
		request = NewGetAlertsRequest()
	}
	response = NewGetAlertsResponse()
	err = c.Send(request, response)
	return
}

func NewSetCdcUserSuspendStatusRequest() (request *SetCdcUserSuspendStatusRequest) {
	request = &SetCdcUserSuspendStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "SetCdcUserSuspendStatus")
	return
}

func NewSetCdcUserSuspendStatusResponse() (response *SetCdcUserSuspendStatusResponse) {
	response = &SetCdcUserSuspendStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置CSP的CDC用户状态
func (c *Client) SetCdcUserSuspendStatus(request *SetCdcUserSuspendStatusRequest) (response *SetCdcUserSuspendStatusResponse, err error) {
	if request == nil {
		request = NewSetCdcUserSuspendStatusRequest()
	}
	response = NewSetCdcUserSuspendStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDoStartHostRequest() (request *DoStartHostRequest) {
	request = &DoStartHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoStartHost")
	return
}

func NewDoStartHostResponse() (response *DoStartHostResponse) {
	response = &DoStartHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动节点
func (c *Client) DoStartHost(request *DoStartHostRequest) (response *DoStartHostResponse, err error) {
	if request == nil {
		request = NewDoStartHostRequest()
	}
	response = NewDoStartHostResponse()
	err = c.Send(request, response)
	return
}

func NewGetHostDisksRequest() (request *GetHostDisksRequest) {
	request = &GetHostDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetHostDisks")
	return
}

func NewGetHostDisksResponse() (response *GetHostDisksResponse) {
	response = &GetHostDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取节点的磁盘信息
func (c *Client) GetHostDisks(request *GetHostDisksRequest) (response *GetHostDisksResponse, err error) {
	if request == nil {
		request = NewGetHostDisksRequest()
	}
	response = NewGetHostDisksResponse()
	err = c.Send(request, response)
	return
}

func NewGetRequestDetailRequest() (request *GetRequestDetailRequest) {
	request = &GetRequestDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetRequestDetail")
	return
}

func NewGetRequestDetailResponse() (response *GetRequestDetailResponse) {
	response = &GetRequestDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取长任务详情
func (c *Client) GetRequestDetail(request *GetRequestDetailRequest) (response *GetRequestDetailResponse, err error) {
	if request == nil {
		request = NewGetRequestDetailRequest()
	}
	response = NewGetRequestDetailResponse()
	err = c.Send(request, response)
	return
}

func NewGetRgwLogDetailRequest() (request *GetRgwLogDetailRequest) {
	request = &GetRgwLogDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetRgwLogDetail")
	return
}

func NewGetRgwLogDetailResponse() (response *GetRgwLogDetailResponse) {
	response = &GetRgwLogDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取单个日志文件内容
func (c *Client) GetRgwLogDetail(request *GetRgwLogDetailRequest) (response *GetRgwLogDetailResponse, err error) {
	if request == nil {
		request = NewGetRgwLogDetailRequest()
	}
	response = NewGetRgwLogDetailResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductFailureMigrateTaskDetailRequest() (request *QueryProductFailureMigrateTaskDetailRequest) {
	request = &QueryProductFailureMigrateTaskDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "QueryProductFailureMigrateTaskDetail")
	return
}

func NewQueryProductFailureMigrateTaskDetailResponse() (response *QueryProductFailureMigrateTaskDetailResponse) {
	response = &QueryProductFailureMigrateTaskDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容灾系统产品故障切换结果查询接口
func (c *Client) QueryProductFailureMigrateTaskDetail(request *QueryProductFailureMigrateTaskDetailRequest) (response *QueryProductFailureMigrateTaskDetailResponse, err error) {
	if request == nil {
		request = NewQueryProductFailureMigrateTaskDetailRequest()
	}
	response = NewQueryProductFailureMigrateTaskDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDoCreatePoolRequest() (request *DoCreatePoolRequest) {
	request = &DoCreatePoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "DoCreatePool")
	return
}

func NewDoCreatePoolResponse() (response *DoCreatePoolResponse) {
	response = &DoCreatePoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建存储池
func (c *Client) DoCreatePool(request *DoCreatePoolRequest) (response *DoCreatePoolResponse, err error) {
	if request == nil {
		request = NewDoCreatePoolRequest()
	}
	response = NewDoCreatePoolResponse()
	err = c.Send(request, response)
	return
}

func NewGetBucketMonitoringRequest() (request *GetBucketMonitoringRequest) {
	request = &GetBucketMonitoringRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("csp", APIVersion, "GetBucketMonitoring")
	return
}

func NewGetBucketMonitoringResponse() (response *GetBucketMonitoringResponse) {
	response = &GetBucketMonitoringResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取储存桶监控数据
func (c *Client) GetBucketMonitoring(request *GetBucketMonitoringRequest) (response *GetBucketMonitoringResponse, err error) {
	if request == nil {
		request = NewGetBucketMonitoringRequest()
	}
	response = NewGetBucketMonitoringResponse()
	err = c.Send(request, response)
	return
}
