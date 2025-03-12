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
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2021-08-09"

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

func NewDescribeAvailableResourceRangeRequest() (request *DescribeAvailableResourceRangeRequest) {
	request = &DescribeAvailableResourceRangeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DescribeAvailableResourceRange")
	return
}

func NewDescribeAvailableResourceRangeResponse() (response *DescribeAvailableResourceRangeResponse) {
	response = &DescribeAvailableResourceRangeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调用该接口查询集群可配置的资源范围
func (c *Client) DescribeAvailableResourceRange(request *DescribeAvailableResourceRangeRequest) (response *DescribeAvailableResourceRangeResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableResourceRangeRequest()
	}
	response = NewDescribeAvailableResourceRangeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSaleZoneConfigsRequest() (request *DescribeSaleZoneConfigsRequest) {
	request = &DescribeSaleZoneConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DescribeSaleZoneConfigs")
	return
}

func NewDescribeSaleZoneConfigsResponse() (response *DescribeSaleZoneConfigsResponse) {
	response = &DescribeSaleZoneConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询售卖可用区信息
func (c *Client) DescribeSaleZoneConfigs(request *DescribeSaleZoneConfigsRequest) (response *DescribeSaleZoneConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeSaleZoneConfigsRequest()
	}
	response = NewDescribeSaleZoneConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyExclusiveGroupVisibilityRequest() (request *ModifyExclusiveGroupVisibilityRequest) {
	request = &ModifyExclusiveGroupVisibilityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "ModifyExclusiveGroupVisibility")
	return
}

func NewModifyExclusiveGroupVisibilityResponse() (response *ModifyExclusiveGroupVisibilityResponse) {
	response = &ModifyExclusiveGroupVisibilityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改资源池可见性
func (c *Client) ModifyExclusiveGroupVisibility(request *ModifyExclusiveGroupVisibilityRequest) (response *ModifyExclusiveGroupVisibilityResponse, err error) {
	if request == nil {
		request = NewModifyExclusiveGroupVisibilityRequest()
	}
	response = NewModifyExclusiveGroupVisibilityResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSaleZoneConfigRequest() (request *DeleteSaleZoneConfigRequest) {
	request = &DeleteSaleZoneConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DeleteSaleZoneConfig")
	return
}

func NewDeleteSaleZoneConfigResponse() (response *DeleteSaleZoneConfigResponse) {
	response = &DeleteSaleZoneConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除售卖可用区，注意只能删除跨可用区售卖的规则，同可用区的售卖规则只能设置为关闭而不可删除
func (c *Client) DeleteSaleZoneConfig(request *DeleteSaleZoneConfigRequest) (response *DeleteSaleZoneConfigResponse, err error) {
	if request == nil {
		request = NewDeleteSaleZoneConfigRequest()
	}
	response = NewDeleteSaleZoneConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBindExclusiveGroupsRequest() (request *DescribeBindExclusiveGroupsRequest) {
	request = &DescribeBindExclusiveGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DescribeBindExclusiveGroups")
	return
}

func NewDescribeBindExclusiveGroupsResponse() (response *DescribeBindExclusiveGroupsResponse) {
	response = &DescribeBindExclusiveGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资源池绑定租户信息
func (c *Client) DescribeBindExclusiveGroups(request *DescribeBindExclusiveGroupsRequest) (response *DescribeBindExclusiveGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeBindExclusiveGroupsRequest()
	}
	response = NewDescribeBindExclusiveGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewIsolateInstanceRequest() (request *IsolateInstanceRequest) {
	request = &IsolateInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "IsolateInstance")
	return
}

func NewIsolateInstanceResponse() (response *IsolateInstanceResponse) {
	response = &IsolateInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 隔离租户实例，一段时间后自动销毁
func (c *Client) IsolateInstance(request *IsolateInstanceRequest) (response *IsolateInstanceResponse, err error) {
	if request == nil {
		request = NewIsolateInstanceRequest()
	}
	response = NewIsolateInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceAvailableResourceRequest() (request *DescribeInstanceAvailableResourceRequest) {
	request = &DescribeInstanceAvailableResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DescribeInstanceAvailableResource")
	return
}

func NewDescribeInstanceAvailableResourceResponse() (response *DescribeInstanceAvailableResourceResponse) {
	response = &DescribeInstanceAvailableResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询热迁机器的可用资源列表
func (c *Client) DescribeInstanceAvailableResource(request *DescribeInstanceAvailableResourceRequest) (response *DescribeInstanceAvailableResourceResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceAvailableResourceRequest()
	}
	response = NewDescribeInstanceAvailableResourceResponse()
	err = c.Send(request, response)
	return
}

func NewAddSaleZoneConfigRequest() (request *AddSaleZoneConfigRequest) {
	request = &AddSaleZoneConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "AddSaleZoneConfig")
	return
}

func NewAddSaleZoneConfigResponse() (response *AddSaleZoneConfigResponse) {
	response = &AddSaleZoneConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加售卖可用区组合
func (c *Client) AddSaleZoneConfig(request *AddSaleZoneConfigRequest) (response *AddSaleZoneConfigResponse, err error) {
	if request == nil {
		request = NewAddSaleZoneConfigRequest()
	}
	response = NewAddSaleZoneConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDBEngineStatusRequest() (request *ModifyDBEngineStatusRequest) {
	request = &ModifyDBEngineStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "ModifyDBEngineStatus")
	return
}

func NewModifyDBEngineStatusResponse() (response *ModifyDBEngineStatusResponse) {
	response = &ModifyDBEngineStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改引擎售卖状态
func (c *Client) ModifyDBEngineStatus(request *ModifyDBEngineStatusRequest) (response *ModifyDBEngineStatusResponse, err error) {
	if request == nil {
		request = NewModifyDBEngineStatusRequest()
	}
	response = NewModifyDBEngineStatusResponse()
	err = c.Send(request, response)
	return
}

func NewMigrateInstanceRequest() (request *MigrateInstanceRequest) {
	request = &MigrateInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "MigrateInstance")
	return
}

func NewMigrateInstanceResponse() (response *MigrateInstanceResponse) {
	response = &MigrateInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将实例迁移到指定的机器上，当前支持通过IP列表指定目标设备
func (c *Client) MigrateInstance(request *MigrateInstanceRequest) (response *MigrateInstanceResponse, err error) {
	if request == nil {
		request = NewMigrateInstanceRequest()
	}
	response = NewMigrateInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewUnbindExclusiveGroupRequest() (request *UnbindExclusiveGroupRequest) {
	request = &UnbindExclusiveGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "UnbindExclusiveGroup")
	return
}

func NewUnbindExclusiveGroupResponse() (response *UnbindExclusiveGroupResponse) {
	response = &UnbindExclusiveGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解绑独显资源池与租户的绑定关系
func (c *Client) UnbindExclusiveGroup(request *UnbindExclusiveGroupRequest) (response *UnbindExclusiveGroupResponse, err error) {
	if request == nil {
		request = NewUnbindExclusiveGroupRequest()
	}
	response = NewUnbindExclusiveGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExclusiveGroupsRequest() (request *DescribeExclusiveGroupsRequest) {
	request = &DescribeExclusiveGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DescribeExclusiveGroups")
	return
}

func NewDescribeExclusiveGroupsResponse() (response *DescribeExclusiveGroupsResponse) {
	response = &DescribeExclusiveGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所有的独享资源池
func (c *Client) DescribeExclusiveGroups(request *DescribeExclusiveGroupsRequest) (response *DescribeExclusiveGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeExclusiveGroupsRequest()
	}
	response = NewDescribeExclusiveGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewModifySpecConfigRequest() (request *ModifySpecConfigRequest) {
	request = &ModifySpecConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "ModifySpecConfig")
	return
}

func NewModifySpecConfigResponse() (response *ModifySpecConfigResponse) {
	response = &ModifySpecConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改规格的属性，当前可修改 cpu 分配值、数据盘可选范围和售卖开关
func (c *Client) ModifySpecConfig(request *ModifySpecConfigRequest) (response *ModifySpecConfigResponse, err error) {
	if request == nil {
		request = NewModifySpecConfigRequest()
	}
	response = NewModifySpecConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTenantDBInstanceDetailRequest() (request *DescribeTenantDBInstanceDetailRequest) {
	request = &DescribeTenantDBInstanceDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DescribeTenantDBInstanceDetail")
	return
}

func NewDescribeTenantDBInstanceDetailResponse() (response *DescribeTenantDBInstanceDetailResponse) {
	response = &DescribeTenantDBInstanceDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据租户实例ID，查询实例详细信息
func (c *Client) DescribeTenantDBInstanceDetail(request *DescribeTenantDBInstanceDetailRequest) (response *DescribeTenantDBInstanceDetailResponse, err error) {
	if request == nil {
		request = NewDescribeTenantDBInstanceDetailRequest()
	}
	response = NewDescribeTenantDBInstanceDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyExclusiveGroupNameRequest() (request *ModifyExclusiveGroupNameRequest) {
	request = &ModifyExclusiveGroupNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "ModifyExclusiveGroupName")
	return
}

func NewModifyExclusiveGroupNameResponse() (response *ModifyExclusiveGroupNameResponse) {
	response = &ModifyExclusiveGroupNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改资源池名称
func (c *Client) ModifyExclusiveGroupName(request *ModifyExclusiveGroupNameRequest) (response *ModifyExclusiveGroupNameResponse, err error) {
	if request == nil {
		request = NewModifyExclusiveGroupNameRequest()
	}
	response = NewModifyExclusiveGroupNameResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceRecycleNetInfosRequest() (request *DescribeInstanceRecycleNetInfosRequest) {
	request = &DescribeInstanceRecycleNetInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DescribeInstanceRecycleNetInfos")
	return
}

func NewDescribeInstanceRecycleNetInfosResponse() (response *DescribeInstanceRecycleNetInfosResponse) {
	response = &DescribeInstanceRecycleNetInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询待回收资源
func (c *Client) DescribeInstanceRecycleNetInfos(request *DescribeInstanceRecycleNetInfosRequest) (response *DescribeInstanceRecycleNetInfosResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceRecycleNetInfosRequest()
	}
	response = NewDescribeInstanceRecycleNetInfosResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceRecycleNetRecycleTimeRequest() (request *ModifyInstanceRecycleNetRecycleTimeRequest) {
	request = &ModifyInstanceRecycleNetRecycleTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "ModifyInstanceRecycleNetRecycleTime")
	return
}

func NewModifyInstanceRecycleNetRecycleTimeResponse() (response *ModifyInstanceRecycleNetRecycleTimeResponse) {
	response = &ModifyInstanceRecycleNetRecycleTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改待回收网络资源的时间
func (c *Client) ModifyInstanceRecycleNetRecycleTime(request *ModifyInstanceRecycleNetRecycleTimeRequest) (response *ModifyInstanceRecycleNetRecycleTimeResponse, err error) {
	if request == nil {
		request = NewModifyInstanceRecycleNetRecycleTimeRequest()
	}
	response = NewModifyInstanceRecycleNetRecycleTimeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWorkFlowTaskInfoRequest() (request *DescribeWorkFlowTaskInfoRequest) {
	request = &DescribeWorkFlowTaskInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DescribeWorkFlowTaskInfo")
	return
}

func NewDescribeWorkFlowTaskInfoResponse() (response *DescribeWorkFlowTaskInfoResponse) {
	response = &DescribeWorkFlowTaskInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据流程信息查询流程任务详情
func (c *Client) DescribeWorkFlowTaskInfo(request *DescribeWorkFlowTaskInfoRequest) (response *DescribeWorkFlowTaskInfoResponse, err error) {
	if request == nil {
		request = NewDescribeWorkFlowTaskInfoRequest()
	}
	response = NewDescribeWorkFlowTaskInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWorkFlowsRequest() (request *DescribeWorkFlowsRequest) {
	request = &DescribeWorkFlowsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DescribeWorkFlows")
	return
}

func NewDescribeWorkFlowsResponse() (response *DescribeWorkFlowsResponse) {
	response = &DescribeWorkFlowsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户端流程信息
func (c *Client) DescribeWorkFlows(request *DescribeWorkFlowsRequest) (response *DescribeWorkFlowsResponse, err error) {
	if request == nil {
		request = NewDescribeWorkFlowsRequest()
	}
	response = NewDescribeWorkFlowsResponse()
	err = c.Send(request, response)
	return
}

func NewModifySaleZoneConfigStatusRequest() (request *ModifySaleZoneConfigStatusRequest) {
	request = &ModifySaleZoneConfigStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "ModifySaleZoneConfigStatus")
	return
}

func NewModifySaleZoneConfigStatusResponse() (response *ModifySaleZoneConfigStatusResponse) {
	response = &ModifySaleZoneConfigStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改可用区的售卖状态
func (c *Client) ModifySaleZoneConfigStatus(request *ModifySaleZoneConfigStatusRequest) (response *ModifySaleZoneConfigStatusResponse, err error) {
	if request == nil {
		request = NewModifySaleZoneConfigStatusRequest()
	}
	response = NewModifySaleZoneConfigStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTenantDBInstancesRequest() (request *DescribeTenantDBInstancesRequest) {
	request = &DescribeTenantDBInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DescribeTenantDBInstances")
	return
}

func NewDescribeTenantDBInstancesResponse() (response *DescribeTenantDBInstancesResponse) {
	response = &DescribeTenantDBInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户实例列表
func (c *Client) DescribeTenantDBInstances(request *DescribeTenantDBInstancesRequest) (response *DescribeTenantDBInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeTenantDBInstancesRequest()
	}
	response = NewDescribeTenantDBInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBEnginesRequest() (request *DescribeDBEnginesRequest) {
	request = &DescribeDBEnginesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DescribeDBEngines")
	return
}

func NewDescribeDBEnginesResponse() (response *DescribeDBEnginesResponse) {
	response = &DescribeDBEnginesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询引擎售卖配置
func (c *Client) DescribeDBEngines(request *DescribeDBEnginesRequest) (response *DescribeDBEnginesResponse, err error) {
	if request == nil {
		request = NewDescribeDBEnginesRequest()
	}
	response = NewDescribeDBEnginesResponse()
	err = c.Send(request, response)
	return
}

func NewAddSpecConfigRequest() (request *AddSpecConfigRequest) {
	request = &AddSpecConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "AddSpecConfig")
	return
}

func NewAddSpecConfigResponse() (response *AddSpecConfigResponse) {
	response = &AddSpecConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 为租户端添加一个售卖规格，注意当前暂不支持同一内存规格对应多个 cpu 规格
func (c *Client) AddSpecConfig(request *AddSpecConfigRequest) (response *AddSpecConfigResponse, err error) {
	if request == nil {
		request = NewAddSpecConfigRequest()
	}
	response = NewAddSpecConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCancelFlowRequest() (request *CancelFlowRequest) {
	request = &CancelFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "CancelFlow")
	return
}

func NewCancelFlowResponse() (response *CancelFlowResponse) {
	response = &CancelFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消租户端流程，仅`AllowedActions`包含`CancelFlow`的流程可以取消
func (c *Client) CancelFlow(request *CancelFlowRequest) (response *CancelFlowResponse, err error) {
	if request == nil {
		request = NewCancelFlowRequest()
	}
	response = NewCancelFlowResponse()
	err = c.Send(request, response)
	return
}

func NewRetryFlowRequest() (request *RetryFlowRequest) {
	request = &RetryFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "RetryFlow")
	return
}

func NewRetryFlowResponse() (response *RetryFlowResponse) {
	response = &RetryFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重试流程，仅`AllowedActions`包含`RetryFlow`的流程可以重试
func (c *Client) RetryFlow(request *RetryFlowRequest) (response *RetryFlowResponse, err error) {
	if request == nil {
		request = NewRetryFlowRequest()
	}
	response = NewRetryFlowResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSpecConfigRequest() (request *DeleteSpecConfigRequest) {
	request = &DeleteSpecConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DeleteSpecConfig")
	return
}

func NewDeleteSpecConfigResponse() (response *DeleteSpecConfigResponse) {
	response = &DeleteSpecConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除不需要再用的售卖规格，注意只能删除没有被引用的售卖规格
func (c *Client) DeleteSpecConfig(request *DeleteSpecConfigRequest) (response *DeleteSpecConfigResponse, err error) {
	if request == nil {
		request = NewDeleteSpecConfigRequest()
	}
	response = NewDeleteSpecConfigResponse()
	err = c.Send(request, response)
	return
}

func NewActivateInstanceRequest() (request *ActivateInstanceRequest) {
	request = &ActivateInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "ActivateInstance")
	return
}

func NewActivateInstanceResponse() (response *ActivateInstanceResponse) {
	response = &ActivateInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恢复已隔离的租户实例
func (c *Client) ActivateInstance(request *ActivateInstanceRequest) (response *ActivateInstanceResponse, err error) {
	if request == nil {
		request = NewActivateInstanceRequest()
	}
	response = NewActivateInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewBindExclusiveGroupRequest() (request *BindExclusiveGroupRequest) {
	request = &BindExclusiveGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "BindExclusiveGroup")
	return
}

func NewBindExclusiveGroupResponse() (response *BindExclusiveGroupResponse) {
	response = &BindExclusiveGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将独享资源池与租户进行绑定
func (c *Client) BindExclusiveGroup(request *BindExclusiveGroupRequest) (response *BindExclusiveGroupResponse, err error) {
	if request == nil {
		request = NewBindExclusiveGroupRequest()
	}
	response = NewBindExclusiveGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDestroyInstanceRequest() (request *DestroyInstanceRequest) {
	request = &DestroyInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DestroyInstance")
	return
}

func NewDestroyInstanceResponse() (response *DestroyInstanceResponse) {
	response = &DestroyInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下线已隔离的租户实例，将会立即删除实例数据，请谨慎操作
func (c *Client) DestroyInstance(request *DestroyInstanceRequest) (response *DestroyInstanceResponse, err error) {
	if request == nil {
		request = NewDestroyInstanceRequest()
	}
	response = NewDestroyInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceReleaseTimeRequest() (request *ModifyInstanceReleaseTimeRequest) {
	request = &ModifyInstanceReleaseTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "ModifyInstanceReleaseTime")
	return
}

func NewModifyInstanceReleaseTimeResponse() (response *ModifyInstanceReleaseTimeResponse) {
	response = &ModifyInstanceReleaseTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改临时实例的释放时间，目前仅限于回档产生的临时实例
func (c *Client) ModifyInstanceReleaseTime(request *ModifyInstanceReleaseTimeRequest) (response *ModifyInstanceReleaseTimeResponse, err error) {
	if request == nil {
		request = NewModifyInstanceReleaseTimeRequest()
	}
	response = NewModifyInstanceReleaseTimeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpecConfigsRequest() (request *DescribeSpecConfigsRequest) {
	request = &DescribeSpecConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DescribeSpecConfigs")
	return
}

func NewDescribeSpecConfigsResponse() (response *DescribeSpecConfigsResponse) {
	response = &DescribeSpecConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户端售卖规格列表，注意各地域分别维护各自售卖规格信息
func (c *Client) DescribeSpecConfigs(request *DescribeSpecConfigsRequest) (response *DescribeSpecConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeSpecConfigsRequest()
	}
	response = NewDescribeSpecConfigsResponse()
	err = c.Send(request, response)
	return
}
