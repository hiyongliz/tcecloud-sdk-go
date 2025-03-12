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
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-06-25"

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

func NewDescribeZoneInstanceConfigInfosRequest() (request *DescribeZoneInstanceConfigInfosRequest) {
	request = &DescribeZoneInstanceConfigInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeZoneInstanceConfigInfos")
	return
}

func NewDescribeZoneInstanceConfigInfosResponse() (response *DescribeZoneInstanceConfigInfosResponse) {
	response = &DescribeZoneInstanceConfigInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeZoneInstanceConfigInfos) 获取可用区的机型信息
func (c *Client) DescribeZoneInstanceConfigInfos(request *DescribeZoneInstanceConfigInfosRequest) (response *DescribeZoneInstanceConfigInfosResponse, err error) {
	if request == nil {
		request = NewDescribeZoneInstanceConfigInfosRequest()
	}
	response = NewDescribeZoneInstanceConfigInfosResponse()
	err = c.Send(request, response)
	return
}

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
	request = &RebootInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "RebootInstances")
	return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
	response = &RebootInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (RebootInstances) 用于重启实例。
func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
	if request == nil {
		request = NewRebootInstancesRequest()
	}
	response = NewRebootInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTroubleshootingPanelRequest() (request *DescribeTroubleshootingPanelRequest) {
	request = &DescribeTroubleshootingPanelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeTroubleshootingPanel")
	return
}

func NewDescribeTroubleshootingPanelResponse() (response *DescribeTroubleshootingPanelResponse) {
	response = &DescribeTroubleshootingPanelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 故障看板查询，返回回滚失败，任务超时，消息丢失，基本完成等任务列表
func (c *Client) DescribeTroubleshootingPanel(request *DescribeTroubleshootingPanelRequest) (response *DescribeTroubleshootingPanelResponse, err error) {
	if request == nil {
		request = NewDescribeTroubleshootingPanelRequest()
	}
	response = NewDescribeTroubleshootingPanelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePreSchedulerRequest() (request *DescribePreSchedulerRequest) {
	request = &DescribePreSchedulerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribePreScheduler")
	return
}

func NewDescribePreSchedulerResponse() (response *DescribePreSchedulerResponse) {
	response = &DescribePreSchedulerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过查询预装箱失败来查询实例售罄的原因
func (c *Client) DescribePreScheduler(request *DescribePreSchedulerRequest) (response *DescribePreSchedulerResponse, err error) {
	if request == nil {
		request = NewDescribePreSchedulerRequest()
	}
	response = NewDescribePreSchedulerResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostsTagRequest() (request *ModifyHostsTagRequest) {
	request = &ModifyHostsTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ModifyHostsTag")
	return
}

func NewModifyHostsTagResponse() (response *ModifyHostsTagResponse) {
	response = &ModifyHostsTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加宿主机标签
func (c *Client) ModifyHostsTag(request *ModifyHostsTagRequest) (response *ModifyHostsTagResponse, err error) {
	if request == nil {
		request = NewModifyHostsTagRequest()
	}
	response = NewModifyHostsTagResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
	request = &DescribeTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeTasks")
	return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
	response = &DescribeTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询不同任务的状态
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
	if request == nil {
		request = NewDescribeTasksRequest()
	}
	response = NewDescribeTasksResponse()
	err = c.Send(request, response)
	return
}

func NewStopInstancesRequest() (request *StopInstancesRequest) {
	request = &StopInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "StopInstances")
	return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
	response = &StopInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (StopInstances) 用于关闭一个或多个实例。
func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
	if request == nil {
		request = NewStopInstancesRequest()
	}
	response = NewStopInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserCpuTypeRequest() (request *ModifyUserCpuTypeRequest) {
	request = &ModifyUserCpuTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ModifyUserCpuType")
	return
}

func NewModifyUserCpuTypeResponse() (response *ModifyUserCpuTypeResponse) {
	response = &ModifyUserCpuTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户CPU绑定类型
func (c *Client) ModifyUserCpuType(request *ModifyUserCpuTypeRequest) (response *ModifyUserCpuTypeResponse, err error) {
	if request == nil {
		request = NewModifyUserCpuTypeRequest()
	}
	response = NewModifyUserCpuTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComputeRequest() (request *DescribeComputeRequest) {
	request = &DescribeComputeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeCompute")
	return
}

func NewDescribeComputeResponse() (response *DescribeComputeResponse) {
	response = &DescribeComputeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 链路日志中查询母机日志
func (c *Client) DescribeCompute(request *DescribeComputeRequest) (response *DescribeComputeResponse, err error) {
	if request == nil {
		request = NewDescribeComputeRequest()
	}
	response = NewDescribeComputeResponse()
	err = c.Send(request, response)
	return
}

func NewStartInstancesRequest() (request *StartInstancesRequest) {
	request = &StartInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "StartInstances")
	return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
	response = &StartInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (StartInstances) 用于启动一个或多个实例
func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
	if request == nil {
		request = NewStartInstancesRequest()
	}
	response = NewStartInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceTypeConfigsRequest() (request *DescribeInstanceTypeConfigsRequest) {
	request = &DescribeInstanceTypeConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeInstanceTypeConfigs")
	return
}

func NewDescribeInstanceTypeConfigsResponse() (response *DescribeInstanceTypeConfigsResponse) {
	response = &DescribeInstanceTypeConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeInstanceTypeConfigs) 用于查询实例机型配置。
func (c *Client) DescribeInstanceTypeConfigs(request *DescribeInstanceTypeConfigsRequest) (response *DescribeInstanceTypeConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceTypeConfigsRequest()
	}
	response = NewDescribeInstanceTypeConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewDetachUsbRequest() (request *DetachUsbRequest) {
	request = &DetachUsbRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DetachUsb")
	return
}

func NewDetachUsbResponse() (response *DetachUsbResponse) {
	response = &DetachUsbResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 实例弹出USB设备
func (c *Client) DetachUsb(request *DetachUsbRequest) (response *DetachUsbResponse, err error) {
	if request == nil {
		request = NewDetachUsbRequest()
	}
	response = NewDetachUsbResponse()
	err = c.Send(request, response)
	return
}

func NewLiveMigrateInstancesRequest() (request *LiveMigrateInstancesRequest) {
	request = &LiveMigrateInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "LiveMigrateInstances")
	return
}

func NewLiveMigrateInstancesResponse() (response *LiveMigrateInstancesResponse) {
	response = &LiveMigrateInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 不关机热迁移实例
func (c *Client) LiveMigrateInstances(request *LiveMigrateInstancesRequest) (response *LiveMigrateInstancesResponse, err error) {
	if request == nil {
		request = NewLiveMigrateInstancesRequest()
	}
	response = NewLiveMigrateInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLatestOperationStateRequest() (request *ModifyLatestOperationStateRequest) {
	request = &ModifyLatestOperationStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ModifyLatestOperationState")
	return
}

func NewModifyLatestOperationStateResponse() (response *ModifyLatestOperationStateResponse) {
	response = &ModifyLatestOperationStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 刷新长期operating状态的虚拟机
func (c *Client) ModifyLatestOperationState(request *ModifyLatestOperationStateRequest) (response *ModifyLatestOperationStateResponse, err error) {
	if request == nil {
		request = NewModifyLatestOperationStateRequest()
	}
	response = NewModifyLatestOperationStateResponse()
	err = c.Send(request, response)
	return
}

func NewQueryTaskReportRequest() (request *QueryTaskReportRequest) {
	request = &QueryTaskReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "QueryTaskReport")
	return
}

func NewQueryTaskReportResponse() (response *QueryTaskReportResponse) {
	response = &QueryTaskReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询 VStation 任务系统报表
func (c *Client) QueryTaskReport(request *QueryTaskReportRequest) (response *QueryTaskReportResponse, err error) {
	if request == nil {
		request = NewQueryTaskReportRequest()
	}
	response = NewQueryTaskReportResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceTypeConfigRequest() (request *ModifyInstanceTypeConfigRequest) {
	request = &ModifyInstanceTypeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ModifyInstanceTypeConfig")
	return
}

func NewModifyInstanceTypeConfigResponse() (response *ModifyInstanceTypeConfigResponse) {
	response = &ModifyInstanceTypeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改全局配置的机器
func (c *Client) ModifyInstanceTypeConfig(request *ModifyInstanceTypeConfigRequest) (response *ModifyInstanceTypeConfigResponse, err error) {
	if request == nil {
		request = NewModifyInstanceTypeConfigRequest()
	}
	response = NewModifyInstanceTypeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateHostTypesRequest() (request *CreateHostTypesRequest) {
	request = &CreateHostTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "CreateHostTypes")
	return
}

func NewCreateHostTypesResponse() (response *CreateHostTypesResponse) {
	response = &CreateHostTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于创建新的宿主机售卖类型
func (c *Client) CreateHostTypes(request *CreateHostTypesRequest) (response *CreateHostTypesResponse, err error) {
	if request == nil {
		request = NewCreateHostTypesRequest()
	}
	response = NewCreateHostTypesResponse()
	err = c.Send(request, response)
	return
}

func NewSyncImagesRequest() (request *SyncImagesRequest) {
	request = &SyncImagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "SyncImages")
	return
}

func NewSyncImagesResponse() (response *SyncImagesResponse) {
	response = &SyncImagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（SyncImages）用于将自定义镜像同步到其它地区。
//
// * 该接口每次调用只支持同步一个镜像。
// * 该接口支持多个同步地域。
// * 单个帐号在每个地域最多支持存在10个自定义镜像。
func (c *Client) SyncImages(request *SyncImagesRequest) (response *SyncImagesResponse, err error) {
	if request == nil {
		request = NewSyncImagesRequest()
	}
	response = NewSyncImagesResponse()
	err = c.Send(request, response)
	return
}

func NewSyncInstanceFamilyConfigsRequest() (request *SyncInstanceFamilyConfigsRequest) {
	request = &SyncInstanceFamilyConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "SyncInstanceFamilyConfigs")
	return
}

func NewSyncInstanceFamilyConfigsResponse() (response *SyncInstanceFamilyConfigsResponse) {
	response = &SyncInstanceFamilyConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将实例族信息同步到指定的可用区
func (c *Client) SyncInstanceFamilyConfigs(request *SyncInstanceFamilyConfigsRequest) (response *SyncInstanceFamilyConfigsResponse, err error) {
	if request == nil {
		request = NewSyncInstanceFamilyConfigsRequest()
	}
	response = NewSyncInstanceFamilyConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCvmTypeMapRequest() (request *CreateCvmTypeMapRequest) {
	request = &CreateCvmTypeMapRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "CreateCvmTypeMap")
	return
}

func NewCreateCvmTypeMapResponse() (response *CreateCvmTypeMapResponse) {
	response = &CreateCvmTypeMapResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增用户自定义机型定义
func (c *Client) CreateCvmTypeMap(request *CreateCvmTypeMapRequest) (response *CreateCvmTypeMapResponse, err error) {
	if request == nil {
		request = NewCreateCvmTypeMapRequest()
	}
	response = NewCreateCvmTypeMapResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePublicImageStatusRequest() (request *UpdatePublicImageStatusRequest) {
	request = &UpdatePublicImageStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "UpdatePublicImageStatus")
	return
}

func NewUpdatePublicImageStatusResponse() (response *UpdatePublicImageStatusResponse) {
	response = &UpdatePublicImageStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改公共镜像状态
func (c *Client) UpdatePublicImageStatus(request *UpdatePublicImageStatusRequest) (response *UpdatePublicImageStatusResponse, err error) {
	if request == nil {
		request = NewUpdatePublicImageStatusRequest()
	}
	response = NewUpdatePublicImageStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSoldPoolRequest() (request *DeleteSoldPoolRequest) {
	request = &DeleteSoldPoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DeleteSoldPool")
	return
}

func NewDeleteSoldPoolResponse() (response *DeleteSoldPoolResponse) {
	response = &DeleteSoldPoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除指定售卖池。
// 注：只有当售卖池内没有任何宿主机的情况下，才允许删除。
func (c *Client) DeleteSoldPool(request *DeleteSoldPoolRequest) (response *DeleteSoldPoolResponse, err error) {
	if request == nil {
		request = NewDeleteSoldPoolRequest()
	}
	response = NewDeleteSoldPoolResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceUsbInfoRequest() (request *DescribeInstanceUsbInfoRequest) {
	request = &DescribeInstanceUsbInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeInstanceUsbInfo")
	return
}

func NewDescribeInstanceUsbInfoResponse() (response *DescribeInstanceUsbInfoResponse) {
	response = &DescribeInstanceUsbInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询实例USB信息
func (c *Client) DescribeInstanceUsbInfo(request *DescribeInstanceUsbInfoRequest) (response *DescribeInstanceUsbInfoResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceUsbInfoRequest()
	}
	response = NewDescribeInstanceUsbInfoResponse()
	err = c.Send(request, response)
	return
}

func NewQueryTaskReportErrorRequest() (request *QueryTaskReportErrorRequest) {
	request = &QueryTaskReportErrorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "QueryTaskReportError")
	return
}

func NewQueryTaskReportErrorResponse() (response *QueryTaskReportErrorResponse) {
	response = &QueryTaskReportErrorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询 VStation 任务系统失败报表
func (c *Client) QueryTaskReportError(request *QueryTaskReportErrorRequest) (response *QueryTaskReportErrorResponse, err error) {
	if request == nil {
		request = NewQueryTaskReportErrorRequest()
	}
	response = NewQueryTaskReportErrorResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCvmApiLogStatisticsRequest() (request *DescribeCvmApiLogStatisticsRequest) {
	request = &DescribeCvmApiLogStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeCvmApiLogStatistics")
	return
}

func NewDescribeCvmApiLogStatisticsResponse() (response *DescribeCvmApiLogStatisticsResponse) {
	response = &DescribeCvmApiLogStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 统计CVM-API日志中不同错误码
func (c *Client) DescribeCvmApiLogStatistics(request *DescribeCvmApiLogStatisticsRequest) (response *DescribeCvmApiLogStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeCvmApiLogStatisticsRequest()
	}
	response = NewDescribeCvmApiLogStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSoldPoolsOwnerRequest() (request *DescribeSoldPoolsOwnerRequest) {
	request = &DescribeSoldPoolsOwnerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeSoldPoolsOwner")
	return
}

func NewDescribeSoldPoolsOwnerResponse() (response *DescribeSoldPoolsOwnerResponse) {
	response = &DescribeSoldPoolsOwnerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询售卖池相应的用户
func (c *Client) DescribeSoldPoolsOwner(request *DescribeSoldPoolsOwnerRequest) (response *DescribeSoldPoolsOwnerResponse, err error) {
	if request == nil {
		request = NewDescribeSoldPoolsOwnerRequest()
	}
	response = NewDescribeSoldPoolsOwnerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowLogsRequest() (request *DescribeFlowLogsRequest) {
	request = &DescribeFlowLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeFlowLogs")
	return
}

func NewDescribeFlowLogsResponse() (response *DescribeFlowLogsResponse) {
	response = &DescribeFlowLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeFlowLogs) 用于查询 CVM 后端业务模块的日志流水。
//
// * 支持根据 AppId、Uin、SubAccountUin 查询该用户在指定业务模块下操作记录。
// * 支持根据 RequestId 查询该请求在 CVM_API 这个模块下的操作日志信息。
// * 支持根据 RequestId 查询该请求调用 CVM_DES 这个模块返回的的任务 ID。
// * 如果参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的日志信息。
func (c *Client) DescribeFlowLogs(request *DescribeFlowLogsRequest) (response *DescribeFlowLogsResponse, err error) {
	if request == nil {
		request = NewDescribeFlowLogsRequest()
	}
	response = NewDescribeFlowLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKeyPairsRequest() (request *DescribeKeyPairsRequest) {
	request = &DescribeKeyPairsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeKeyPairs")
	return
}

func NewDescribeKeyPairsResponse() (response *DescribeKeyPairsResponse) {
	response = &DescribeKeyPairsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询密钥对
func (c *Client) DescribeKeyPairs(request *DescribeKeyPairsRequest) (response *DescribeKeyPairsResponse, err error) {
	if request == nil {
		request = NewDescribeKeyPairsRequest()
	}
	response = NewDescribeKeyPairsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceFamilyConfigsRequest() (request *DescribeInstanceFamilyConfigsRequest) {
	request = &DescribeInstanceFamilyConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeInstanceFamilyConfigs")
	return
}

func NewDescribeInstanceFamilyConfigsResponse() (response *DescribeInstanceFamilyConfigsResponse) {
	response = &DescribeInstanceFamilyConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeInstanceFamilyConfigs）查询当前用户和地域所支持的机型族列表信息
func (c *Client) DescribeInstanceFamilyConfigs(request *DescribeInstanceFamilyConfigsRequest) (response *DescribeInstanceFamilyConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceFamilyConfigsRequest()
	}
	response = NewDescribeInstanceFamilyConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFaultRecordRequest() (request *DescribeFaultRecordRequest) {
	request = &DescribeFaultRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeFaultRecord")
	return
}

func NewDescribeFaultRecordResponse() (response *DescribeFaultRecordResponse) {
	response = &DescribeFaultRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询故障迁移记录
func (c *Client) DescribeFaultRecord(request *DescribeFaultRecordRequest) (response *DescribeFaultRecordResponse, err error) {
	if request == nil {
		request = NewDescribeFaultRecordRequest()
	}
	response = NewDescribeFaultRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCpuModelsRequest() (request *DescribeCpuModelsRequest) {
	request = &DescribeCpuModelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeCpuModels")
	return
}

func NewDescribeCpuModelsResponse() (response *DescribeCpuModelsResponse) {
	response = &DescribeCpuModelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询支持的CPU架构列表
func (c *Client) DescribeCpuModels(request *DescribeCpuModelsRequest) (response *DescribeCpuModelsResponse, err error) {
	if request == nil {
		request = NewDescribeCpuModelsRequest()
	}
	response = NewDescribeCpuModelsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteHostsTagRequest() (request *DeleteHostsTagRequest) {
	request = &DeleteHostsTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DeleteHostsTag")
	return
}

func NewDeleteHostsTagResponse() (response *DeleteHostsTagResponse) {
	response = &DeleteHostsTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除宿主机指定的标签
// 只有宿主机IP + tag + reason 与已有的信息一致时，才能删除成功，否则无法删除。
func (c *Client) DeleteHostsTag(request *DeleteHostsTagRequest) (response *DeleteHostsTagResponse, err error) {
	if request == nil {
		request = NewDeleteHostsTagRequest()
	}
	response = NewDeleteHostsTagResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseAgentRequest() (request *ReleaseAgentRequest) {
	request = &ReleaseAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ReleaseAgent")
	return
}

func NewReleaseAgentResponse() (response *ReleaseAgentResponse) {
	response = &ReleaseAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 允许使用压缩包，发布Agent
func (c *Client) ReleaseAgent(request *ReleaseAgentRequest) (response *ReleaseAgentResponse, err error) {
	if request == nil {
		request = NewReleaseAgentRequest()
	}
	response = NewReleaseAgentResponse()
	err = c.Send(request, response)
	return
}

func NewFailTaskRequest() (request *FailTaskRequest) {
	request = &FailTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "FailTask")
	return
}

func NewFailTaskResponse() (response *FailTaskResponse) {
	response = &FailTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置一个vs任务为失败
func (c *Client) FailTask(request *FailTaskRequest) (response *FailTaskResponse, err error) {
	if request == nil {
		request = NewFailTaskRequest()
	}
	response = NewFailTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCDHExternalMachineTypeRequest() (request *CreateCDHExternalMachineTypeRequest) {
	request = &CreateCDHExternalMachineTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "CreateCDHExternalMachineType")
	return
}

func NewCreateCDHExternalMachineTypeResponse() (response *CreateCDHExternalMachineTypeResponse) {
	response = &CreateCDHExternalMachineTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateCDHExternalMachineType）用于创建CDH机型。
func (c *Client) CreateCDHExternalMachineType(request *CreateCDHExternalMachineTypeRequest) (response *CreateCDHExternalMachineTypeResponse, err error) {
	if request == nil {
		request = NewCreateCDHExternalMachineTypeRequest()
	}
	response = NewCreateCDHExternalMachineTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSchedulerRequest() (request *DescribeSchedulerRequest) {
	request = &DescribeSchedulerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeScheduler")
	return
}

func NewDescribeSchedulerResponse() (response *DescribeSchedulerResponse) {
	response = &DescribeSchedulerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 主要查询创建虚拟机任务装箱失败的原因
func (c *Client) DescribeScheduler(request *DescribeSchedulerRequest) (response *DescribeSchedulerResponse, err error) {
	if request == nil {
		request = NewDescribeSchedulerRequest()
	}
	response = NewDescribeSchedulerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTenantInstancesRequest() (request *DescribeTenantInstancesRequest) {
	request = &DescribeTenantInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeTenantInstances")
	return
}

func NewDescribeTenantInstancesResponse() (response *DescribeTenantInstancesResponse) {
	response = &DescribeTenantInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端调用
func (c *Client) DescribeTenantInstances(request *DescribeTenantInstancesRequest) (response *DescribeTenantInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeTenantInstancesRequest()
	}
	response = NewDescribeTenantInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewGetRecycleInfoRequest() (request *GetRecycleInfoRequest) {
	request = &GetRecycleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "GetRecycleInfo")
	return
}

func NewGetRecycleInfoResponse() (response *GetRecycleInfoResponse) {
	response = &GetRecycleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetRecycleInfo
func (c *Client) GetRecycleInfo(request *GetRecycleInfoRequest) (response *GetRecycleInfoResponse, err error) {
	if request == nil {
		request = NewGetRecycleInfoRequest()
	}
	response = NewGetRecycleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCvmTypeMapRequest() (request *ModifyCvmTypeMapRequest) {
	request = &ModifyCvmTypeMapRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ModifyCvmTypeMap")
	return
}

func NewModifyCvmTypeMapResponse() (response *ModifyCvmTypeMapResponse) {
	response = &ModifyCvmTypeMapResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 机型绑定CPU架构
func (c *Client) ModifyCvmTypeMap(request *ModifyCvmTypeMapRequest) (response *ModifyCvmTypeMapResponse, err error) {
	if request == nil {
		request = NewModifyCvmTypeMapRequest()
	}
	response = NewModifyCvmTypeMapResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogStatisticsRequest() (request *DescribeLogStatisticsRequest) {
	request = &DescribeLogStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeLogStatistics")
	return
}

func NewDescribeLogStatisticsResponse() (response *DescribeLogStatisticsResponse) {
	response = &DescribeLogStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 统计vstation日志，Allinone日志的不同状态码的结果
func (c *Client) DescribeLogStatistics(request *DescribeLogStatisticsRequest) (response *DescribeLogStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeLogStatisticsRequest()
	}
	response = NewDescribeLogStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMQQueuesRequest() (request *DescribeMQQueuesRequest) {
	request = &DescribeMQQueuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeMQQueues")
	return
}

func NewDescribeMQQueuesResponse() (response *DescribeMQQueuesResponse) {
	response = &DescribeMQQueuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看MQ队列
func (c *Client) DescribeMQQueues(request *DescribeMQQueuesRequest) (response *DescribeMQQueuesResponse, err error) {
	if request == nil {
		request = NewDescribeMQQueuesRequest()
	}
	response = NewDescribeMQQueuesResponse()
	err = c.Send(request, response)
	return
}

func NewRetryTaskRequest() (request *RetryTaskRequest) {
	request = &RetryTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "RetryTask")
	return
}

func NewRetryTaskResponse() (response *RetryTaskResponse) {
	response = &RetryTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Vstation任务重试
func (c *Client) RetryTask(request *RetryTaskRequest) (response *RetryTaskResponse, err error) {
	if request == nil {
		request = NewRetryTaskRequest()
	}
	response = NewRetryTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHostUsbRequest() (request *DescribeHostUsbRequest) {
	request = &DescribeHostUsbRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeHostUsb")
	return
}

func NewDescribeHostUsbResponse() (response *DescribeHostUsbResponse) {
	response = &DescribeHostUsbResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询宿主机USB信息
func (c *Client) DescribeHostUsb(request *DescribeHostUsbRequest) (response *DescribeHostUsbResponse, err error) {
	if request == nil {
		request = NewDescribeHostUsbRequest()
	}
	response = NewDescribeHostUsbResponse()
	err = c.Send(request, response)
	return
}

func NewRetryFaultMigrateRequest() (request *RetryFaultMigrateRequest) {
	request = &RetryFaultMigrateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "RetryFaultMigrate")
	return
}

func NewRetryFaultMigrateResponse() (response *RetryFaultMigrateResponse) {
	response = &RetryFaultMigrateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 故障迁移重试
func (c *Client) RetryFaultMigrate(request *RetryFaultMigrateRequest) (response *RetryFaultMigrateResponse, err error) {
	if request == nil {
		request = NewRetryFaultMigrateRequest()
	}
	response = NewRetryFaultMigrateResponse()
	err = c.Send(request, response)
	return
}

func NewResetInstanceRequest() (request *ResetInstanceRequest) {
	request = &ResetInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ResetInstance")
	return
}

func NewResetInstanceResponse() (response *ResetInstanceResponse) {
	response = &ResetInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ResetInstance) 用于重装指定实例上的操作系统。
func (c *Client) ResetInstance(request *ResetInstanceRequest) (response *ResetInstanceResponse, err error) {
	if request == nil {
		request = NewResetInstanceRequest()
	}
	response = NewResetInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePublicImageRequest() (request *CreatePublicImageRequest) {
	request = &CreatePublicImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "CreatePublicImage")
	return
}

func NewCreatePublicImageResponse() (response *CreatePublicImageResponse) {
	response = &CreatePublicImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreatePublicImage将自定义镜像转为公共镜像
func (c *Client) CreatePublicImage(request *CreatePublicImageRequest) (response *CreatePublicImageResponse, err error) {
	if request == nil {
		request = NewCreatePublicImageRequest()
	}
	response = NewCreatePublicImageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHostTypesRequest() (request *DescribeHostTypesRequest) {
	request = &DescribeHostTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeHostTypes")
	return
}

func NewDescribeHostTypesResponse() (response *DescribeHostTypesResponse) {
	response = &DescribeHostTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当前系统支持的宿主机类型列表
func (c *Client) DescribeHostTypes(request *DescribeHostTypesRequest) (response *DescribeHostTypesResponse, err error) {
	if request == nil {
		request = NewDescribeHostTypesRequest()
	}
	response = NewDescribeHostTypesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFaultEventRequest() (request *DescribeFaultEventRequest) {
	request = &DescribeFaultEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeFaultEvent")
	return
}

func NewDescribeFaultEventResponse() (response *DescribeFaultEventResponse) {
	response = &DescribeFaultEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询故障迁移事件
func (c *Client) DescribeFaultEvent(request *DescribeFaultEventRequest) (response *DescribeFaultEventResponse, err error) {
	if request == nil {
		request = NewDescribeFaultEventRequest()
	}
	response = NewDescribeFaultEventResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePrivateImageStatusRequest() (request *UpdatePrivateImageStatusRequest) {
	request = &UpdatePrivateImageStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "UpdatePrivateImageStatus")
	return
}

func NewUpdatePrivateImageStatusResponse() (response *UpdatePrivateImageStatusResponse) {
	response = &UpdatePrivateImageStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改私有镜像状态
func (c *Client) UpdatePrivateImageStatus(request *UpdatePrivateImageStatusRequest) (response *UpdatePrivateImageStatusResponse, err error) {
	if request == nil {
		request = NewUpdatePrivateImageStatusRequest()
	}
	response = NewUpdatePrivateImageStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateInstanceTypeConfigRequest() (request *CreateInstanceTypeConfigRequest) {
	request = &CreateInstanceTypeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "CreateInstanceTypeConfig")
	return
}

func NewCreateInstanceTypeConfigResponse() (response *CreateInstanceTypeConfigResponse) {
	response = &CreateInstanceTypeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供运营端调用，创建全部规格配置
func (c *Client) CreateInstanceTypeConfig(request *CreateInstanceTypeConfigRequest) (response *CreateInstanceTypeConfigResponse, err error) {
	if request == nil {
		request = NewCreateInstanceTypeConfigRequest()
	}
	response = NewCreateInstanceTypeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewSyncHostTypesRequest() (request *SyncHostTypesRequest) {
	request = &SyncHostTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "SyncHostTypes")
	return
}

func NewSyncHostTypesResponse() (response *SyncHostTypesResponse) {
	response = &SyncHostTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将指定地域的宿主机类型配置同步到当前地域
// 1、已存在的宿主机类型配置不会被覆盖
// 2、可装箱的虚拟机类型如果在当前地域不支持，则同步过来的宿主机类型也不支持该虚拟机类型
func (c *Client) SyncHostTypes(request *SyncHostTypesRequest) (response *SyncHostTypesResponse, err error) {
	if request == nil {
		request = NewSyncHostTypesRequest()
	}
	response = NewSyncHostTypesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateImageRequest() (request *CreateImageRequest) {
	request = &CreateImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "CreateImage")
	return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
	response = &CreateImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建自定义镜像接口
func (c *Client) CreateImage(request *CreateImageRequest) (response *CreateImageResponse, err error) {
	if request == nil {
		request = NewCreateImageRequest()
	}
	response = NewCreateImageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllinoneTasksRequest() (request *DescribeAllinoneTasksRequest) {
	request = &DescribeAllinoneTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeAllinoneTasks")
	return
}

func NewDescribeAllinoneTasksResponse() (response *DescribeAllinoneTasksResponse) {
	response = &DescribeAllinoneTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Vstation allinone日志详情
func (c *Client) DescribeAllinoneTasks(request *DescribeAllinoneTasksRequest) (response *DescribeAllinoneTasksResponse, err error) {
	if request == nil {
		request = NewDescribeAllinoneTasksRequest()
	}
	response = NewDescribeAllinoneTasksResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostsResourceExclusiveHostFlagRequest() (request *ModifyHostsResourceExclusiveHostFlagRequest) {
	request = &ModifyHostsResourceExclusiveHostFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ModifyHostsResourceExclusiveHostFlag")
	return
}

func NewModifyHostsResourceExclusiveHostFlagResponse() (response *ModifyHostsResourceExclusiveHostFlagResponse) {
	response = &ModifyHostsResourceExclusiveHostFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改母机独立宿主机标识
func (c *Client) ModifyHostsResourceExclusiveHostFlag(request *ModifyHostsResourceExclusiveHostFlagRequest) (response *ModifyHostsResourceExclusiveHostFlagResponse, err error) {
	if request == nil {
		request = NewModifyHostsResourceExclusiveHostFlagRequest()
	}
	response = NewModifyHostsResourceExclusiveHostFlagResponse()
	err = c.Send(request, response)
	return
}

func NewUpgradeInstanceRequest() (request *UpgradeInstanceRequest) {
	request = &UpgradeInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "UpgradeInstance")
	return
}

func NewUpgradeInstanceResponse() (response *UpgradeInstanceResponse) {
	response = &UpgradeInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 自由调整实例配置
func (c *Client) UpgradeInstance(request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
	if request == nil {
		request = NewUpgradeInstanceRequest()
	}
	response = NewUpgradeInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSoldPoolsRequest() (request *DescribeSoldPoolsRequest) {
	request = &DescribeSoldPoolsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeSoldPools")
	return
}

func NewDescribeSoldPoolsResponse() (response *DescribeSoldPoolsResponse) {
	response = &DescribeSoldPoolsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询售卖池的基本信息
func (c *Client) DescribeSoldPools(request *DescribeSoldPoolsRequest) (response *DescribeSoldPoolsResponse, err error) {
	if request == nil {
		request = NewDescribeSoldPoolsRequest()
	}
	response = NewDescribeSoldPoolsResponse()
	err = c.Send(request, response)
	return
}

func NewCheckCreatePublicImageNameRequest() (request *CheckCreatePublicImageNameRequest) {
	request = &CheckCreatePublicImageNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "CheckCreatePublicImageName")
	return
}

func NewCheckCreatePublicImageNameResponse() (response *CheckCreatePublicImageNameResponse) {
	response = &CheckCreatePublicImageNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建公共镜像时，查验名称是否已经存在
func (c *Client) CheckCreatePublicImageName(request *CheckCreatePublicImageNameRequest) (response *CheckCreatePublicImageNameResponse, err error) {
	if request == nil {
		request = NewCheckCreatePublicImageNameRequest()
	}
	response = NewCheckCreatePublicImageNameResponse()
	err = c.Send(request, response)
	return
}

func NewColdMigrateTenantInstancesRequest() (request *ColdMigrateTenantInstancesRequest) {
	request = &ColdMigrateTenantInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ColdMigrateTenantInstances")
	return
}

func NewColdMigrateTenantInstancesResponse() (response *ColdMigrateTenantInstancesResponse) {
	response = &ColdMigrateTenantInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端调用，去冷迁移租户端的子机
func (c *Client) ColdMigrateTenantInstances(request *ColdMigrateTenantInstancesRequest) (response *ColdMigrateTenantInstancesResponse, err error) {
	if request == nil {
		request = NewColdMigrateTenantInstancesRequest()
	}
	response = NewColdMigrateTenantInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDeviceStatusRequest() (request *ModifyDeviceStatusRequest) {
	request = &ModifyDeviceStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ModifyDeviceStatus")
	return
}

func NewModifyDeviceStatusResponse() (response *ModifyDeviceStatusResponse) {
	response = &ModifyDeviceStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 当已退款的虚拟机在CVM这边遇到异常导致前端显示一直创建中时需要手动刷新数据库
func (c *Client) ModifyDeviceStatus(request *ModifyDeviceStatusRequest) (response *ModifyDeviceStatusResponse, err error) {
	if request == nil {
		request = NewModifyDeviceStatusRequest()
	}
	response = NewModifyDeviceStatusResponse()
	err = c.Send(request, response)
	return
}

func NewAttachUsbRequest() (request *AttachUsbRequest) {
	request = &AttachUsbRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "AttachUsb")
	return
}

func NewAttachUsbResponse() (response *AttachUsbResponse) {
	response = &AttachUsbResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 实例绑定USB设备
func (c *Client) AttachUsb(request *AttachUsbRequest) (response *AttachUsbResponse, err error) {
	if request == nil {
		request = NewAttachUsbRequest()
	}
	response = NewAttachUsbResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeModifyRecycleTimeLogsRequest() (request *DescribeModifyRecycleTimeLogsRequest) {
	request = &DescribeModifyRecycleTimeLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeModifyRecycleTimeLogs")
	return
}

func NewDescribeModifyRecycleTimeLogsResponse() (response *DescribeModifyRecycleTimeLogsResponse) {
	response = &DescribeModifyRecycleTimeLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用来查询历史的回收站停留时长记录
func (c *Client) DescribeModifyRecycleTimeLogs(request *DescribeModifyRecycleTimeLogsRequest) (response *DescribeModifyRecycleTimeLogsResponse, err error) {
	if request == nil {
		request = NewDescribeModifyRecycleTimeLogsRequest()
	}
	response = NewDescribeModifyRecycleTimeLogsResponse()
	err = c.Send(request, response)
	return
}

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
	request = &TerminateInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "TerminateInstances")
	return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
	response = &TerminateInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (TerminateInstances) 用于主动退还实例。
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
	if request == nil {
		request = NewTerminateInstancesRequest()
	}
	response = NewTerminateInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewBindCpuCoresRequest() (request *BindCpuCoresRequest) {
	request = &BindCpuCoresRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "BindCpuCores")
	return
}

func NewBindCpuCoresResponse() (response *BindCpuCoresResponse) {
	response = &BindCpuCoresResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定指定的cpu set
func (c *Client) BindCpuCores(request *BindCpuCoresRequest) (response *BindCpuCoresResponse, err error) {
	if request == nil {
		request = NewBindCpuCoresRequest()
	}
	response = NewBindCpuCoresResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceConfigInfosRequest() (request *DescribeInstanceConfigInfosRequest) {
	request = &DescribeInstanceConfigInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeInstanceConfigInfos")
	return
}

func NewDescribeInstanceConfigInfosResponse() (response *DescribeInstanceConfigInfosResponse) {
	response = &DescribeInstanceConfigInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询全局机器信息
func (c *Client) DescribeInstanceConfigInfos(request *DescribeInstanceConfigInfosRequest) (response *DescribeInstanceConfigInfosResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceConfigInfosRequest()
	}
	response = NewDescribeInstanceConfigInfosResponse()
	err = c.Send(request, response)
	return
}

func NewModifyZoneInstanceTypeConfigRequest() (request *ModifyZoneInstanceTypeConfigRequest) {
	request = &ModifyZoneInstanceTypeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ModifyZoneInstanceTypeConfig")
	return
}

func NewModifyZoneInstanceTypeConfigResponse() (response *ModifyZoneInstanceTypeConfigResponse) {
	response = &ModifyZoneInstanceTypeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改可用区内的实例机型
func (c *Client) ModifyZoneInstanceTypeConfig(request *ModifyZoneInstanceTypeConfigRequest) (response *ModifyZoneInstanceTypeConfigResponse, err error) {
	if request == nil {
		request = NewModifyZoneInstanceTypeConfigRequest()
	}
	response = NewModifyZoneInstanceTypeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewQueryHostInstanceMapRequest() (request *QueryHostInstanceMapRequest) {
	request = &QueryHostInstanceMapRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "QueryHostInstanceMap")
	return
}

func NewQueryHostInstanceMapResponse() (response *QueryHostInstanceMapResponse) {
	response = &QueryHostInstanceMapResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 机型和实例交叉索引
func (c *Client) QueryHostInstanceMap(request *QueryHostInstanceMapRequest) (response *QueryHostInstanceMapResponse, err error) {
	if request == nil {
		request = NewQueryHostInstanceMapRequest()
	}
	response = NewQueryHostInstanceMapResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteImagesRequest() (request *DeleteImagesRequest) {
	request = &DeleteImagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DeleteImages")
	return
}

func NewDeleteImagesResponse() (response *DeleteImagesResponse) {
	response = &DeleteImagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteImages）用于删除一个或多个镜像。
func (c *Client) DeleteImages(request *DeleteImagesRequest) (response *DeleteImagesResponse, err error) {
	if request == nil {
		request = NewDeleteImagesRequest()
	}
	response = NewDeleteImagesResponse()
	err = c.Send(request, response)
	return
}

func NewBindSoldPoolOwnersRequest() (request *BindSoldPoolOwnersRequest) {
	request = &BindSoldPoolOwnersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "BindSoldPoolOwners")
	return
}

func NewBindSoldPoolOwnersResponse() (response *BindSoldPoolOwnersResponse) {
	response = &BindSoldPoolOwnersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将指定账号列表绑定到售卖池内
func (c *Client) BindSoldPoolOwners(request *BindSoldPoolOwnersRequest) (response *BindSoldPoolOwnersResponse, err error) {
	if request == nil {
		request = NewBindSoldPoolOwnersRequest()
	}
	response = NewBindSoldPoolOwnersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceClassConfigsRequest() (request *DescribeInstanceClassConfigsRequest) {
	request = &DescribeInstanceClassConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeInstanceClassConfigs")
	return
}

func NewDescribeInstanceClassConfigsResponse() (response *DescribeInstanceClassConfigsResponse) {
	response = &DescribeInstanceClassConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当前支持的实例系列的相信信息，比如“S”
func (c *Client) DescribeInstanceClassConfigs(request *DescribeInstanceClassConfigsRequest) (response *DescribeInstanceClassConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceClassConfigsRequest()
	}
	response = NewDescribeInstanceClassConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewGetHostsCountRequest() (request *GetHostsCountRequest) {
	request = &GetHostsCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "GetHostsCount")
	return
}

func NewGetHostsCountResponse() (response *GetHostsCountResponse) {
	response = &GetHostsCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取宿主机数量
func (c *Client) GetHostsCount(request *GetHostsCountRequest) (response *GetHostsCountResponse, err error) {
	if request == nil {
		request = NewGetHostsCountRequest()
	}
	response = NewGetHostsCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
	request = &DescribeImagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeImages")
	return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
	response = &DescribeImagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 统一查询镜像的接口
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
	if request == nil {
		request = NewDescribeImagesRequest()
	}
	response = NewDescribeImagesResponse()
	err = c.Send(request, response)
	return
}

func NewUnbindSoldPoolOwnersRequest() (request *UnbindSoldPoolOwnersRequest) {
	request = &UnbindSoldPoolOwnersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "UnbindSoldPoolOwners")
	return
}

func NewUnbindSoldPoolOwnersResponse() (response *UnbindSoldPoolOwnersResponse) {
	response = &UnbindSoldPoolOwnersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 售卖池中解绑账户，账户相应无法在使用该售卖池中的资源
func (c *Client) UnbindSoldPoolOwners(request *UnbindSoldPoolOwnersRequest) (response *UnbindSoldPoolOwnersResponse, err error) {
	if request == nil {
		request = NewUnbindSoldPoolOwnersRequest()
	}
	response = NewUnbindSoldPoolOwnersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCvmTypeBillRequest() (request *CreateCvmTypeBillRequest) {
	request = &CreateCvmTypeBillRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "CreateCvmTypeBill")
	return
}

func NewCreateCvmTypeBillResponse() (response *CreateCvmTypeBillResponse) {
	response = &CreateCvmTypeBillResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增机型四层定义上报
func (c *Client) CreateCvmTypeBill(request *CreateCvmTypeBillRequest) (response *CreateCvmTypeBillResponse, err error) {
	if request == nil {
		request = NewCreateCvmTypeBillRequest()
	}
	response = NewCreateCvmTypeBillResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVncRequest() (request *DescribeVncRequest) {
	request = &DescribeVncRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeVnc")
	return
}

func NewDescribeVncResponse() (response *DescribeVncResponse) {
	response = &DescribeVncResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看VNC信息
func (c *Client) DescribeVnc(request *DescribeVncRequest) (response *DescribeVncResponse, err error) {
	if request == nil {
		request = NewDescribeVncRequest()
	}
	response = NewDescribeVncResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVsChildTaskIdByParentRequest() (request *DescribeVsChildTaskIdByParentRequest) {
	request = &DescribeVsChildTaskIdByParentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeVsChildTaskIdByParent")
	return
}

func NewDescribeVsChildTaskIdByParentResponse() (response *DescribeVsChildTaskIdByParentResponse) {
	response = &DescribeVsChildTaskIdByParentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeVsChildTaskIdByParent) 用于根据父 CVM_VS 任务ID查询子任务。
func (c *Client) DescribeVsChildTaskIdByParent(request *DescribeVsChildTaskIdByParentRequest) (response *DescribeVsChildTaskIdByParentResponse, err error) {
	if request == nil {
		request = NewDescribeVsChildTaskIdByParentRequest()
	}
	response = NewDescribeVsChildTaskIdByParentResponse()
	err = c.Send(request, response)
	return
}

func NewFailoverMigrateInstancesRequest() (request *FailoverMigrateInstancesRequest) {
	request = &FailoverMigrateInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "FailoverMigrateInstances")
	return
}

func NewFailoverMigrateInstancesResponse() (response *FailoverMigrateInstancesResponse) {
	response = &FailoverMigrateInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 当母机产生故障的时候迁移虚拟机
func (c *Client) FailoverMigrateInstances(request *FailoverMigrateInstancesRequest) (response *FailoverMigrateInstancesResponse, err error) {
	if request == nil {
		request = NewFailoverMigrateInstancesRequest()
	}
	response = NewFailoverMigrateInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAllHostTypesRequest() (request *CreateAllHostTypesRequest) {
	request = &CreateAllHostTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "CreateAllHostTypes")
	return
}

func NewCreateAllHostTypesResponse() (response *CreateAllHostTypesResponse) {
	response = &CreateAllHostTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 宿主机可售卖资源上报
func (c *Client) CreateAllHostTypes(request *CreateAllHostTypesRequest) (response *CreateAllHostTypesResponse, err error) {
	if request == nil {
		request = NewCreateAllHostTypesRequest()
	}
	response = NewCreateAllHostTypesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceFamilyRequest() (request *ModifyInstanceFamilyRequest) {
	request = &ModifyInstanceFamilyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ModifyInstanceFamily")
	return
}

func NewModifyInstanceFamilyResponse() (response *ModifyInstanceFamilyResponse) {
	response = &ModifyInstanceFamilyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改机型实例族信息，目前仅支持修改白名单
func (c *Client) ModifyInstanceFamily(request *ModifyInstanceFamilyRequest) (response *ModifyInstanceFamilyResponse, err error) {
	if request == nil {
		request = NewModifyInstanceFamilyRequest()
	}
	response = NewModifyInstanceFamilyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVsTaskInfoRequest() (request *DescribeVsTaskInfoRequest) {
	request = &DescribeVsTaskInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeVsTaskInfo")
	return
}

func NewDescribeVsTaskInfoResponse() (response *DescribeVsTaskInfoResponse) {
	response = &DescribeVsTaskInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询vstation链路日志时每个task id的任务日志
func (c *Client) DescribeVsTaskInfo(request *DescribeVsTaskInfoRequest) (response *DescribeVsTaskInfoResponse, err error) {
	if request == nil {
		request = NewDescribeVsTaskInfoRequest()
	}
	response = NewDescribeVsTaskInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstancesAttributeRequest() (request *ModifyInstancesAttributeRequest) {
	request = &ModifyInstancesAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ModifyInstancesAttribute")
	return
}

func NewModifyInstancesAttributeResponse() (response *ModifyInstancesAttributeResponse) {
	response = &ModifyInstancesAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyInstancesAttribute) 用于修改实例的属性（目前只支持修改实例的名称）。
func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
	if request == nil {
		request = NewModifyInstancesAttributeRequest()
	}
	response = NewModifyInstancesAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceUsbControllerRequest() (request *DescribeInstanceUsbControllerRequest) {
	request = &DescribeInstanceUsbControllerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeInstanceUsbController")
	return
}

func NewDescribeInstanceUsbControllerResponse() (response *DescribeInstanceUsbControllerResponse) {
	response = &DescribeInstanceUsbControllerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询实例USB控制器信息
func (c *Client) DescribeInstanceUsbController(request *DescribeInstanceUsbControllerRequest) (response *DescribeInstanceUsbControllerResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceUsbControllerRequest()
	}
	response = NewDescribeInstanceUsbControllerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZoneTenantInstanceConfigInfosRequest() (request *DescribeZoneTenantInstanceConfigInfosRequest) {
	request = &DescribeZoneTenantInstanceConfigInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeZoneTenantInstanceConfigInfos")
	return
}

func NewDescribeZoneTenantInstanceConfigInfosResponse() (response *DescribeZoneTenantInstanceConfigInfosResponse) {
	response = &DescribeZoneTenantInstanceConfigInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取指定租户的可用区的机型信息
func (c *Client) DescribeZoneTenantInstanceConfigInfos(request *DescribeZoneTenantInstanceConfigInfosRequest) (response *DescribeZoneTenantInstanceConfigInfosResponse, err error) {
	if request == nil {
		request = NewDescribeZoneTenantInstanceConfigInfosRequest()
	}
	response = NewDescribeZoneTenantInstanceConfigInfosResponse()
	err = c.Send(request, response)
	return
}

func NewRunInstancesRequest() (request *RunInstancesRequest) {
	request = &RunInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "RunInstances")
	return
}

func NewRunInstancesResponse() (response *RunInstancesResponse) {
	response = &RunInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端创建实例，以运营端的身份创建实例，和租户资源分开
func (c *Client) RunInstances(request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
	if request == nil {
		request = NewRunInstancesRequest()
	}
	response = NewRunInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateImageVisibleRangeRequest() (request *UpdateImageVisibleRangeRequest) {
	request = &UpdateImageVisibleRangeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "UpdateImageVisibleRange")
	return
}

func NewUpdateImageVisibleRangeResponse() (response *UpdateImageVisibleRangeResponse) {
	response = &UpdateImageVisibleRangeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改公共镜像的生效范围。
// 镜像生效范围只能增加,不能取消,镜像生效范围是租户端和运营端,不再支持修改镜像生效范围
func (c *Client) UpdateImageVisibleRange(request *UpdateImageVisibleRangeRequest) (response *UpdateImageVisibleRangeResponse, err error) {
	if request == nil {
		request = NewUpdateImageVisibleRangeRequest()
	}
	response = NewUpdateImageVisibleRangeResponse()
	err = c.Send(request, response)
	return
}

func NewModifySoldPoolRequest() (request *ModifySoldPoolRequest) {
	request = &ModifySoldPoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ModifySoldPool")
	return
}

func NewModifySoldPoolResponse() (response *ModifySoldPoolResponse) {
	response = &ModifySoldPoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改售卖池的信息，支持修改描述信息、调高售卖比(谨慎调整)、空售卖池支持调低售卖比、售卖比支持小数(需要客制化开启)
func (c *Client) ModifySoldPool(request *ModifySoldPoolRequest) (response *ModifySoldPoolResponse, err error) {
	if request == nil {
		request = NewModifySoldPoolRequest()
	}
	response = NewModifySoldPoolResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageRecordRequest() (request *DescribeImageRecordRequest) {
	request = &DescribeImageRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeImageRecord")
	return
}

func NewDescribeImageRecordResponse() (response *DescribeImageRecordResponse) {
	response = &DescribeImageRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 包括自定义镜像转公共镜像及公共镜像同步到多region的记录
func (c *Client) DescribeImageRecord(request *DescribeImageRecordRequest) (response *DescribeImageRecordResponse, err error) {
	if request == nil {
		request = NewDescribeImageRecordRequest()
	}
	response = NewDescribeImageRecordResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSoldPoolRequest() (request *CreateSoldPoolRequest) {
	request = &CreateSoldPoolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "CreateSoldPool")
	return
}

func NewCreateSoldPoolResponse() (response *CreateSoldPoolResponse) {
	response = &CreateSoldPoolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建售卖池，请谨慎选择售卖池名称和售卖比例
// 注意：可以支持售卖比例的修改，不能使用系统默认的三个池子："plain"、"oversold"、"small"。售卖比默认为整数，支持小数需要开启客制化(谨慎开启)
func (c *Client) CreateSoldPool(request *CreateSoldPoolRequest) (response *CreateSoldPoolResponse, err error) {
	if request == nil {
		request = NewCreateSoldPoolRequest()
	}
	response = NewCreateSoldPoolResponse()
	err = c.Send(request, response)
	return
}

func NewMigrantLiveMigrateRequest() (request *MigrantLiveMigrateRequest) {
	request = &MigrantLiveMigrateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "MigrantLiveMigrate")
	return
}

func NewMigrantLiveMigrateResponse() (response *MigrantLiveMigrateResponse) {
	response = &MigrantLiveMigrateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量热迁移接口
func (c *Client) MigrantLiveMigrate(request *MigrantLiveMigrateRequest) (response *MigrantLiveMigrateResponse, err error) {
	if request == nil {
		request = NewMigrantLiveMigrateRequest()
	}
	response = NewMigrantLiveMigrateResponse()
	err = c.Send(request, response)
	return
}

func NewLiveMigrateTenantInstancesRequest() (request *LiveMigrateTenantInstancesRequest) {
	request = &LiveMigrateTenantInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "LiveMigrateTenantInstances")
	return
}

func NewLiveMigrateTenantInstancesResponse() (response *LiveMigrateTenantInstancesResponse) {
	response = &LiveMigrateTenantInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 热迁移租户端实例
func (c *Client) LiveMigrateTenantInstances(request *LiveMigrateTenantInstancesRequest) (response *LiveMigrateTenantInstancesResponse, err error) {
	if request == nil {
		request = NewLiveMigrateTenantInstancesRequest()
	}
	response = NewLiveMigrateTenantInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCvmTypeMapWithHostInfosRequest() (request *DescribeCvmTypeMapWithHostInfosRequest) {
	request = &DescribeCvmTypeMapWithHostInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeCvmTypeMapWithHostInfos")
	return
}

func NewDescribeCvmTypeMapWithHostInfosResponse() (response *DescribeCvmTypeMapWithHostInfosResponse) {
	response = &DescribeCvmTypeMapWithHostInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询机型实例族和对应的宿主机信息
func (c *Client) DescribeCvmTypeMapWithHostInfos(request *DescribeCvmTypeMapWithHostInfosRequest) (response *DescribeCvmTypeMapWithHostInfosResponse, err error) {
	if request == nil {
		request = NewDescribeCvmTypeMapWithHostInfosRequest()
	}
	response = NewDescribeCvmTypeMapWithHostInfosResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCDHExternalMachineTypeRequest() (request *DeleteCDHExternalMachineTypeRequest) {
	request = &DeleteCDHExternalMachineTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DeleteCDHExternalMachineType")
	return
}

func NewDeleteCDHExternalMachineTypeResponse() (response *DeleteCDHExternalMachineTypeResponse) {
	response = &DeleteCDHExternalMachineTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于，删除CDH机型信息
func (c *Client) DeleteCDHExternalMachineType(request *DeleteCDHExternalMachineTypeRequest) (response *DeleteCDHExternalMachineTypeResponse, err error) {
	if request == nil {
		request = NewDeleteCDHExternalMachineTypeRequest()
	}
	response = NewDeleteCDHExternalMachineTypeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceTypeDisasterGroupConfigRequest() (request *ModifyInstanceTypeDisasterGroupConfigRequest) {
	request = &ModifyInstanceTypeDisasterGroupConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ModifyInstanceTypeDisasterGroupConfig")
	return
}

func NewModifyInstanceTypeDisasterGroupConfigResponse() (response *ModifyInstanceTypeDisasterGroupConfigResponse) {
	response = &ModifyInstanceTypeDisasterGroupConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理机型置放群组黑名单
func (c *Client) ModifyInstanceTypeDisasterGroupConfig(request *ModifyInstanceTypeDisasterGroupConfigRequest) (response *ModifyInstanceTypeDisasterGroupConfigResponse, err error) {
	if request == nil {
		request = NewModifyInstanceTypeDisasterGroupConfigRequest()
	}
	response = NewModifyInstanceTypeDisasterGroupConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCDHExternalMachineTypeRequest() (request *DescribeCDHExternalMachineTypeRequest) {
	request = &DescribeCDHExternalMachineTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeCDHExternalMachineType")
	return
}

func NewDescribeCDHExternalMachineTypeResponse() (response *DescribeCDHExternalMachineTypeResponse) {
	response = &DescribeCDHExternalMachineTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口用于查询CDH机型信息
func (c *Client) DescribeCDHExternalMachineType(request *DescribeCDHExternalMachineTypeRequest) (response *DescribeCDHExternalMachineTypeResponse, err error) {
	if request == nil {
		request = NewDescribeCDHExternalMachineTypeRequest()
	}
	response = NewDescribeCDHExternalMachineTypeResponse()
	err = c.Send(request, response)
	return
}

func NewColdMigrateInstancesRequest() (request *ColdMigrateInstancesRequest) {
	request = &ColdMigrateInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ColdMigrateInstances")
	return
}

func NewColdMigrateInstancesResponse() (response *ColdMigrateInstancesResponse) {
	response = &ColdMigrateInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将子机迁移到指定的宿主机上面
func (c *Client) ColdMigrateInstances(request *ColdMigrateInstancesRequest) (response *ColdMigrateInstancesResponse, err error) {
	if request == nil {
		request = NewColdMigrateInstancesRequest()
	}
	response = NewColdMigrateInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostsResourceRequest() (request *ModifyHostsResourceRequest) {
	request = &ModifyHostsResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ModifyHostsResource")
	return
}

func NewModifyHostsResourceResponse() (response *ModifyHostsResourceResponse) {
	response = &ModifyHostsResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改宿主机资源
func (c *Client) ModifyHostsResource(request *ModifyHostsResourceRequest) (response *ModifyHostsResourceResponse, err error) {
	if request == nil {
		request = NewModifyHostsResourceRequest()
	}
	response = NewModifyHostsResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHostsRequest() (request *DescribeHostsRequest) {
	request = &DescribeHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeHosts")
	return
}

func NewDescribeHostsResponse() (response *DescribeHostsResponse) {
	response = &DescribeHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询宿主机信息
func (c *Client) DescribeHosts(request *DescribeHostsRequest) (response *DescribeHostsResponse, err error) {
	if request == nil {
		request = NewDescribeHostsRequest()
	}
	response = NewDescribeHostsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeInstances")
	return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeInstances) 用于查询一个或多个实例的详细信息。
// * 可以根据实例`ID`、实例名称或者实例计费模式等信息来查询实例的详细信息。过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的实例。
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesRequest()
	}
	response = NewDescribeInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRecycleTimeRequest() (request *ModifyRecycleTimeRequest) {
	request = &ModifyRecycleTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ModifyRecycleTime")
	return
}

func NewModifyRecycleTimeResponse() (response *ModifyRecycleTimeResponse) {
	response = &ModifyRecycleTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyRecycleTime）用于自定义配置云服务器回收时间接口。
func (c *Client) ModifyRecycleTime(request *ModifyRecycleTimeRequest) (response *ModifyRecycleTimeResponse, err error) {
	if request == nil {
		request = NewModifyRecycleTimeRequest()
	}
	response = NewModifyRecycleTimeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCDHHostTypesRequest() (request *DescribeCDHHostTypesRequest) {
	request = &DescribeCDHHostTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeCDHHostTypes")
	return
}

func NewDescribeCDHHostTypesResponse() (response *DescribeCDHHostTypesResponse) {
	response = &DescribeCDHHostTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口用于查询CDH宿主机机型信息
func (c *Client) DescribeCDHHostTypes(request *DescribeCDHHostTypesRequest) (response *DescribeCDHHostTypesResponse, err error) {
	if request == nil {
		request = NewDescribeCDHHostTypesRequest()
	}
	response = NewDescribeCDHHostTypesResponse()
	err = c.Send(request, response)
	return
}

func NewResetInstancesPasswordRequest() (request *ResetInstancesPasswordRequest) {
	request = &ResetInstancesPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "ResetInstancesPassword")
	return
}

func NewResetInstancesPasswordResponse() (response *ResetInstancesPasswordResponse) {
	response = &ResetInstancesPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ResetInstancesPassword) 用于将实例操作系统的密码重置为用户指定的密码。
func (c *Client) ResetInstancesPassword(request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
	if request == nil {
		request = NewResetInstancesPasswordRequest()
	}
	response = NewResetInstancesPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZoneImportInstanceTypeConfigsRequest() (request *DescribeZoneImportInstanceTypeConfigsRequest) {
	request = &DescribeZoneImportInstanceTypeConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeZoneImportInstanceTypeConfigs")
	return
}

func NewDescribeZoneImportInstanceTypeConfigsResponse() (response *DescribeZoneImportInstanceTypeConfigsResponse) {
	response = &DescribeZoneImportInstanceTypeConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看指定可用区可导入的机型配置
func (c *Client) DescribeZoneImportInstanceTypeConfigs(request *DescribeZoneImportInstanceTypeConfigsRequest) (response *DescribeZoneImportInstanceTypeConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeZoneImportInstanceTypeConfigsRequest()
	}
	response = NewDescribeZoneImportInstanceTypeConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentLogsRequest() (request *DescribeAgentLogsRequest) {
	request = &DescribeAgentLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeAgentLogs")
	return
}

func NewDescribeAgentLogsResponse() (response *DescribeAgentLogsResponse) {
	response = &DescribeAgentLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 允许查看Agent发布日志
func (c *Client) DescribeAgentLogs(request *DescribeAgentLogsRequest) (response *DescribeAgentLogsResponse, err error) {
	if request == nil {
		request = NewDescribeAgentLogsRequest()
	}
	response = NewDescribeAgentLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVsChildTasksRequest() (request *DescribeVsChildTasksRequest) {
	request = &DescribeVsChildTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeVsChildTasks")
	return
}

func NewDescribeVsChildTasksResponse() (response *DescribeVsChildTasksResponse) {
	response = &DescribeVsChildTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeVsChildTasks) 用于查询 CVM_VS 子任务详情。
func (c *Client) DescribeVsChildTasks(request *DescribeVsChildTasksRequest) (response *DescribeVsChildTasksResponse, err error) {
	if request == nil {
		request = NewDescribeVsChildTasksRequest()
	}
	response = NewDescribeVsChildTasksResponse()
	err = c.Send(request, response)
	return
}

func NewQueryTaskTranslogRequest() (request *QueryTaskTranslogRequest) {
	request = &QueryTaskTranslogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "QueryTaskTranslog")
	return
}

func NewQueryTaskTranslogResponse() (response *QueryTaskTranslogResponse) {
	response = &QueryTaskTranslogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询任务详细信息
func (c *Client) QueryTaskTranslog(request *QueryTaskTranslogRequest) (response *QueryTaskTranslogResponse, err error) {
	if request == nil {
		request = NewQueryTaskTranslogRequest()
	}
	response = NewQueryTaskTranslogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentsRequest() (request *DescribeAgentsRequest) {
	request = &DescribeAgentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "DescribeAgents")
	return
}

func NewDescribeAgentsResponse() (response *DescribeAgentsResponse) {
	response = &DescribeAgentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询系统中当前存在的agent
func (c *Client) DescribeAgents(request *DescribeAgentsRequest) (response *DescribeAgentsResponse, err error) {
	if request == nil {
		request = NewDescribeAgentsRequest()
	}
	response = NewDescribeAgentsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateZoneInstanceTypeConfigRequest() (request *CreateZoneInstanceTypeConfigRequest) {
	request = &CreateZoneInstanceTypeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opcvm", APIVersion, "CreateZoneInstanceTypeConfig")
	return
}

func NewCreateZoneInstanceTypeConfigResponse() (response *CreateZoneInstanceTypeConfigResponse) {
	response = &CreateZoneInstanceTypeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 基于全局机型配置创建可用区机型
func (c *Client) CreateZoneInstanceTypeConfig(request *CreateZoneInstanceTypeConfigRequest) (response *CreateZoneInstanceTypeConfigResponse, err error) {
	if request == nil {
		request = NewCreateZoneInstanceTypeConfigRequest()
	}
	response = NewCreateZoneInstanceTypeConfigResponse()
	err = c.Send(request, response)
	return
}
