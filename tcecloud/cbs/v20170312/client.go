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

package v20170312

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2017-03-12"

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

func NewDescribeDiskDetailsRequest() (request *DescribeDiskDetailsRequest) {
	request = &DescribeDiskDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskDetails")
	return
}

func NewDescribeDiskDetailsResponse() (response *DescribeDiskDetailsResponse) {
	response = &DescribeDiskDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskDetails）用于查询云硬盘详情。
func (c *Client) DescribeDiskDetails(request *DescribeDiskDetailsRequest) (response *DescribeDiskDetailsResponse, err error) {
	if request == nil {
		request = NewDescribeDiskDetailsRequest()
	}
	response = NewDescribeDiskDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSnapshotBoxConfigsRequest() (request *DescribeSnapshotBoxConfigsRequest) {
	request = &DescribeSnapshotBoxConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshotBoxConfigs")
	return
}

func NewDescribeSnapshotBoxConfigsResponse() (response *DescribeSnapshotBoxConfigsResponse) {
	response = &DescribeSnapshotBoxConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询快照集群的配置。
func (c *Client) DescribeSnapshotBoxConfigs(request *DescribeSnapshotBoxConfigsRequest) (response *DescribeSnapshotBoxConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeSnapshotBoxConfigsRequest()
	}
	response = NewDescribeSnapshotBoxConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateDepotCellDeviceConfigRequest() (request *UpdateDepotCellDeviceConfigRequest) {
	request = &UpdateDepotCellDeviceConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "UpdateDepotCellDeviceConfig")
	return
}

func NewUpdateDepotCellDeviceConfigResponse() (response *UpdateDepotCellDeviceConfigResponse) {
	response = &UpdateDepotCellDeviceConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于更新存储池机型配置
func (c *Client) UpdateDepotCellDeviceConfig(request *UpdateDepotCellDeviceConfigRequest) (response *UpdateDepotCellDeviceConfigResponse, err error) {
	if request == nil {
		request = NewUpdateDepotCellDeviceConfigRequest()
	}
	response = NewUpdateDepotCellDeviceConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDepotConfigsRequest() (request *DescribeDepotConfigsRequest) {
	request = &DescribeDepotConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDepotConfigs")
	return
}

func NewDescribeDepotConfigsResponse() (response *DescribeDepotConfigsResponse) {
	response = &DescribeDepotConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDepotConfigs）用于描述存储池的属性值，包括master、cell、dbtrsf等组件的配置属性。
func (c *Client) DescribeDepotConfigs(request *DescribeDepotConfigsRequest) (response *DescribeDepotConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeDepotConfigsRequest()
	}
	response = NewDescribeDepotConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewReplaceDiskStorageDepotDiskRequest() (request *ReplaceDiskStorageDepotDiskRequest) {
	request = &ReplaceDiskStorageDepotDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ReplaceDiskStorageDepotDisk")
	return
}

func NewReplaceDiskStorageDepotDiskResponse() (response *ReplaceDiskStorageDepotDiskResponse) {
	response = &ReplaceDiskStorageDepotDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// * 本接口（ReplaceDiskStorageDepotDisk）用于更换故障的物理盘
func (c *Client) ReplaceDiskStorageDepotDisk(request *ReplaceDiskStorageDepotDiskRequest) (response *ReplaceDiskStorageDepotDiskResponse, err error) {
	if request == nil {
		request = NewReplaceDiskStorageDepotDiskRequest()
	}
	response = NewReplaceDiskStorageDepotDiskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBalanceDepotTasksRequest() (request *DescribeBalanceDepotTasksRequest) {
	request = &DescribeBalanceDepotTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeBalanceDepotTasks")
	return
}

func NewDescribeBalanceDepotTasksResponse() (response *DescribeBalanceDepotTasksResponse) {
	response = &DescribeBalanceDepotTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeBalanceDepotTasks）用于查询当前正在做小表自动均衡的存储集群进度。本接口会返回当前所有正在进行自动均衡的存储集群列表及其进度详情，
// 不在返回结果中的，说明当前没有小表均衡任务。
func (c *Client) DescribeBalanceDepotTasks(request *DescribeBalanceDepotTasksRequest) (response *DescribeBalanceDepotTasksResponse, err error) {
	if request == nil {
		request = NewDescribeBalanceDepotTasksRequest()
	}
	response = NewDescribeBalanceDepotTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskZKClustersRequest() (request *DescribeDiskZKClustersRequest) {
	request = &DescribeDiskZKClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskZKClusters")
	return
}

func NewDescribeDiskZKClustersResponse() (response *DescribeDiskZKClustersResponse) {
	response = &DescribeDiskZKClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskZKClusters）用于查询云硬盘后端ZooKeeper集群信息
func (c *Client) DescribeDiskZKClusters(request *DescribeDiskZKClustersRequest) (response *DescribeDiskZKClustersResponse, err error) {
	if request == nil {
		request = NewDescribeDiskZKClustersRequest()
	}
	response = NewDescribeDiskZKClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSetHostSideIOStatRequest() (request *DescribeSetHostSideIOStatRequest) {
	request = &DescribeSetHostSideIOStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeSetHostSideIOStat")
	return
}

func NewDescribeSetHostSideIOStatResponse() (response *DescribeSetHostSideIOStatResponse) {
	response = &DescribeSetHostSideIOStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询由云盘IO监控数据汇聚后的存储池IO监控数据
func (c *Client) DescribeSetHostSideIOStat(request *DescribeSetHostSideIOStatRequest) (response *DescribeSetHostSideIOStatResponse, err error) {
	if request == nil {
		request = NewDescribeSetHostSideIOStatRequest()
	}
	response = NewDescribeSetHostSideIOStatResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskStorageDeviceTypesRequest() (request *DescribeDiskStorageDeviceTypesRequest) {
	request = &DescribeDiskStorageDeviceTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskStorageDeviceTypes")
	return
}

func NewDescribeDiskStorageDeviceTypesResponse() (response *DescribeDiskStorageDeviceTypesResponse) {
	response = &DescribeDiskStorageDeviceTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskStorageDeviceTypes）用于查询云硬盘存储仓库支持的机型列表。
func (c *Client) DescribeDiskStorageDeviceTypes(request *DescribeDiskStorageDeviceTypesRequest) (response *DescribeDiskStorageDeviceTypesResponse, err error) {
	if request == nil {
		request = NewDescribeDiskStorageDeviceTypesRequest()
	}
	response = NewDescribeDiskStorageDeviceTypesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAutoSnapshotPoliciesRequest() (request *DescribeAutoSnapshotPoliciesRequest) {
	request = &DescribeAutoSnapshotPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeAutoSnapshotPolicies")
	return
}

func NewDescribeAutoSnapshotPoliciesResponse() (response *DescribeAutoSnapshotPoliciesResponse) {
	response = &DescribeAutoSnapshotPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeAutoSnapshotPolicies）用于查询定期快照策略。
//
// * 可以根据定期快照策略ID、名称或者状态等信息来查询定期快照策略的详细信息，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的定期快照策略表。
func (c *Client) DescribeAutoSnapshotPolicies(request *DescribeAutoSnapshotPoliciesRequest) (response *DescribeAutoSnapshotPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeAutoSnapshotPoliciesRequest()
	}
	response = NewDescribeAutoSnapshotPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewInstallDiskZKClusterRequest() (request *InstallDiskZKClusterRequest) {
	request = &InstallDiskZKClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "InstallDiskZKCluster")
	return
}

func NewInstallDiskZKClusterResponse() (response *InstallDiskZKClusterResponse) {
	response = &InstallDiskZKClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InstallDiskZkCluster）用于安装云硬盘ZooKeeper集群。
func (c *Client) InstallDiskZKCluster(request *InstallDiskZKClusterRequest) (response *InstallDiskZKClusterResponse, err error) {
	if request == nil {
		request = NewInstallDiskZKClusterRequest()
	}
	response = NewInstallDiskZKClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskMaintenanceTasksRequest() (request *DescribeDiskMaintenanceTasksRequest) {
	request = &DescribeDiskMaintenanceTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskMaintenanceTasks")
	return
}

func NewDescribeDiskMaintenanceTasksResponse() (response *DescribeDiskMaintenanceTasksResponse) {
	response = &DescribeDiskMaintenanceTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskMaintenanceTasks）用于查询云硬盘运维任务。
func (c *Client) DescribeDiskMaintenanceTasks(request *DescribeDiskMaintenanceTasksRequest) (response *DescribeDiskMaintenanceTasksResponse, err error) {
	if request == nil {
		request = NewDescribeDiskMaintenanceTasksRequest()
	}
	response = NewDescribeDiskMaintenanceTasksResponse()
	err = c.Send(request, response)
	return
}

func NewReplaceDiskStorageDepotMasterRequest() (request *ReplaceDiskStorageDepotMasterRequest) {
	request = &ReplaceDiskStorageDepotMasterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ReplaceDiskStorageDepotMaster")
	return
}

func NewReplaceDiskStorageDepotMasterResponse() (response *ReplaceDiskStorageDepotMasterResponse) {
	response = &ReplaceDiskStorageDepotMasterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ReplaceDiskStorageDepotMaster）用于存储集群master出现故障，替换整台master机器。
// * 只能替换slave节点，不可替换master节点，如需要替换master节点，请先进行主备替换
// * slave节点更换，要求当前master节点状态为正常
func (c *Client) ReplaceDiskStorageDepotMaster(request *ReplaceDiskStorageDepotMasterRequest) (response *ReplaceDiskStorageDepotMasterResponse, err error) {
	if request == nil {
		request = NewReplaceDiskStorageDepotMasterRequest()
	}
	response = NewReplaceDiskStorageDepotMasterResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDiskStoragePoolGroupRequest() (request *ModifyDiskStoragePoolGroupRequest) {
	request = &ModifyDiskStoragePoolGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifyDiskStoragePoolGroup")
	return
}

func NewModifyDiskStoragePoolGroupResponse() (response *ModifyDiskStoragePoolGroupResponse) {
	response = &ModifyDiskStoragePoolGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于修改云硬盘存储池组属性
func (c *Client) ModifyDiskStoragePoolGroup(request *ModifyDiskStoragePoolGroupRequest) (response *ModifyDiskStoragePoolGroupResponse, err error) {
	if request == nil {
		request = NewModifyDiskStoragePoolGroupRequest()
	}
	response = NewModifyDiskStoragePoolGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskResourceUsageRequest() (request *DescribeDiskResourceUsageRequest) {
	request = &DescribeDiskResourceUsageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskResourceUsage")
	return
}

func NewDescribeDiskResourceUsageResponse() (response *DescribeDiskResourceUsageResponse) {
	response = &DescribeDiskResourceUsageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于查询云硬盘资源全局使用统计
func (c *Client) DescribeDiskResourceUsage(request *DescribeDiskResourceUsageRequest) (response *DescribeDiskResourceUsageResponse, err error) {
	if request == nil {
		request = NewDescribeDiskResourceUsageRequest()
	}
	response = NewDescribeDiskResourceUsageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDiskStorageDepotAttributesRequest() (request *ModifyDiskStorageDepotAttributesRequest) {
	request = &ModifyDiskStorageDepotAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifyDiskStorageDepotAttributes")
	return
}

func NewModifyDiskStorageDepotAttributesResponse() (response *ModifyDiskStorageDepotAttributesResponse) {
	response = &ModifyDiskStorageDepotAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyDiskStorageDepotAttributes）用于修改云硬盘存储仓库配置。
func (c *Client) ModifyDiskStorageDepotAttributes(request *ModifyDiskStorageDepotAttributesRequest) (response *ModifyDiskStorageDepotAttributesResponse, err error) {
	if request == nil {
		request = NewModifyDiskStorageDepotAttributesRequest()
	}
	response = NewModifyDiskStorageDepotAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewBindUserToDiskStoragePoolGroupRequest() (request *BindUserToDiskStoragePoolGroupRequest) {
	request = &BindUserToDiskStoragePoolGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "BindUserToDiskStoragePoolGroup")
	return
}

func NewBindUserToDiskStoragePoolGroupResponse() (response *BindUserToDiskStoragePoolGroupResponse) {
	response = &BindUserToDiskStoragePoolGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于绑定用户到特定的云硬盘资源池
func (c *Client) BindUserToDiskStoragePoolGroup(request *BindUserToDiskStoragePoolGroupRequest) (response *BindUserToDiskStoragePoolGroupResponse, err error) {
	if request == nil {
		request = NewBindUserToDiskStoragePoolGroupRequest()
	}
	response = NewBindUserToDiskStoragePoolGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomerCapacityTopChangeStartDateRequest() (request *DescribeCustomerCapacityTopChangeStartDateRequest) {
	request = &DescribeCustomerCapacityTopChangeStartDateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeCustomerCapacityTopChangeStartDate")
	return
}

func NewDescribeCustomerCapacityTopChangeStartDateResponse() (response *DescribeCustomerCapacityTopChangeStartDateResponse) {
	response = &DescribeCustomerCapacityTopChangeStartDateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询top规模统计最小日期
func (c *Client) DescribeCustomerCapacityTopChangeStartDate(request *DescribeCustomerCapacityTopChangeStartDateRequest) (response *DescribeCustomerCapacityTopChangeStartDateResponse, err error) {
	if request == nil {
		request = NewDescribeCustomerCapacityTopChangeStartDateRequest()
	}
	response = NewDescribeCustomerCapacityTopChangeStartDateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReplaceDiskTaskRequest() (request *ModifyReplaceDiskTaskRequest) {
	request = &ModifyReplaceDiskTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifyReplaceDiskTask")
	return
}

func NewModifyReplaceDiskTaskResponse() (response *ModifyReplaceDiskTaskResponse) {
	response = &ModifyReplaceDiskTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// * 本接口（ModifyReplaceDiskTask）用于修改换盘任务的状态
// * 当换盘任务处于等待确认状态（JobSatus='wait_check'），并确认现场已换好新盘后，调该接口把任务修改为check_finish状态，换盘任务继续进行；
// * 如果需要取消换盘任务，调本接口将任务状态修改为stop；
func (c *Client) ModifyReplaceDiskTask(request *ModifyReplaceDiskTaskRequest) (response *ModifyReplaceDiskTaskResponse, err error) {
	if request == nil {
		request = NewModifyReplaceDiskTaskRequest()
	}
	response = NewModifyReplaceDiskTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSnapshotBoxUpgradeTasksRequest() (request *CreateSnapshotBoxUpgradeTasksRequest) {
	request = &CreateSnapshotBoxUpgradeTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "CreateSnapshotBoxUpgradeTasks")
	return
}

func NewCreateSnapshotBoxUpgradeTasksResponse() (response *CreateSnapshotBoxUpgradeTasksResponse) {
	response = &CreateSnapshotBoxUpgradeTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateSnapshotBoxUpgradeTasks）用于生成快照集群升级任务。
// * 选定快照集群的目标版本后，会自动生成该快照集群所有组件的升级任务；
func (c *Client) CreateSnapshotBoxUpgradeTasks(request *CreateSnapshotBoxUpgradeTasksRequest) (response *CreateSnapshotBoxUpgradeTasksResponse, err error) {
	if request == nil {
		request = NewCreateSnapshotBoxUpgradeTasksRequest()
	}
	response = NewCreateSnapshotBoxUpgradeTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAutoSnapshotPoliciesRequest() (request *DeleteAutoSnapshotPoliciesRequest) {
	request = &DeleteAutoSnapshotPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DeleteAutoSnapshotPolicies")
	return
}

func NewDeleteAutoSnapshotPoliciesResponse() (response *DeleteAutoSnapshotPoliciesResponse) {
	response = &DeleteAutoSnapshotPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteAutoSnapshotPolicies）用于删除定期快照策略。
//
// *  支持批量操作。如果多个定期快照策略存在无法删除的，则操作不执行，以特定错误码返回。
func (c *Client) DeleteAutoSnapshotPolicies(request *DeleteAutoSnapshotPoliciesRequest) (response *DeleteAutoSnapshotPoliciesResponse, err error) {
	if request == nil {
		request = NewDeleteAutoSnapshotPoliciesRequest()
	}
	response = NewDeleteAutoSnapshotPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewApplySnapshotRequest() (request *ApplySnapshotRequest) {
	request = &ApplySnapshotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ApplySnapshot")
	return
}

func NewApplySnapshotResponse() (response *ApplySnapshotResponse) {
	response = &ApplySnapshotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ApplySnapshot）用于回滚快照到原云硬盘。
//
// * 仅支持回滚到原云硬盘上。对于数据盘快照，如果您需要复制快照数据到其它云硬盘上，请使用[CreateDisks](/document/product/362/16312)接口创建新的弹性云盘，将快照数据复制到新购云盘上。
// * 用于回滚的快照必须处于NORMAL状态。快照状态可以通过[DescribeSnapshots](/document/product/362/15647)接口查询，见输出参数中SnapshotState字段解释。
// * 如果是弹性云盘，则云盘必须处于未挂载状态，云硬盘挂载状态可以通过[DescribeDisks](/document/product/362/16315)接口查询，见Attached字段解释；如果是随云主机一起购买的非弹性云盘，则云主机必须处于关机状态，云主机状态可以通过[DescribeInstancesStatus](/document/product/213/15738)接口查询。
func (c *Client) ApplySnapshot(request *ApplySnapshotRequest) (response *ApplySnapshotResponse, err error) {
	if request == nil {
		request = NewApplySnapshotRequest()
	}
	response = NewApplySnapshotResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskOperationsV2Request() (request *DescribeDiskOperationsV2Request) {
	request = &DescribeDiskOperationsV2Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskOperationsV2")
	return
}

func NewDescribeDiskOperationsV2Response() (response *DescribeDiskOperationsV2Response) {
	response = &DescribeDiskOperationsV2Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云硬盘及快照相关任务列表
func (c *Client) DescribeDiskOperationsV2(request *DescribeDiskOperationsV2Request) (response *DescribeDiskOperationsV2Response, err error) {
	if request == nil {
		request = NewDescribeDiskOperationsV2Request()
	}
	response = NewDescribeDiskOperationsV2Response()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskConfigQuotaRequest() (request *DescribeDiskConfigQuotaRequest) {
	request = &DescribeDiskConfigQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskConfigQuota")
	return
}

func NewDescribeDiskConfigQuotaResponse() (response *DescribeDiskConfigQuotaResponse) {
	response = &DescribeDiskConfigQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskConfigQuota）用于查询云硬盘配额。
func (c *Client) DescribeDiskConfigQuota(request *DescribeDiskConfigQuotaRequest) (response *DescribeDiskConfigQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeDiskConfigQuotaRequest()
	}
	response = NewDescribeDiskConfigQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskFlowControlRecordsRequest() (request *DescribeDiskFlowControlRecordsRequest) {
	request = &DescribeDiskFlowControlRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskFlowControlRecords")
	return
}

func NewDescribeDiskFlowControlRecordsResponse() (response *DescribeDiskFlowControlRecordsResponse) {
	response = &DescribeDiskFlowControlRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskFlowControlRecords）用于查询云硬盘流控历史。
func (c *Client) DescribeDiskFlowControlRecords(request *DescribeDiskFlowControlRecordsRequest) (response *DescribeDiskFlowControlRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeDiskFlowControlRecordsRequest()
	}
	response = NewDescribeDiskFlowControlRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSetBlockUsageDayStatisticsRequest() (request *DescribeSetBlockUsageDayStatisticsRequest) {
	request = &DescribeSetBlockUsageDayStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeSetBlockUsageDayStatistics")
	return
}

func NewDescribeSetBlockUsageDayStatisticsResponse() (response *DescribeSetBlockUsageDayStatisticsResponse) {
	response = &DescribeSetBlockUsageDayStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSetBlockUsageDayStatistics）用于查询快照的详细信息。\r\n\r\n* 根据仓库ID查询该仓库在过去一段时间的最大block利用率的值，查询结果是最长过去60天block利用率；查询的结果主要分为两部分，一部分是过去60天到昨天的block利用率，每天一个数据；第二部分是当天最近12小时的block利用率，每半小时一个数据。
func (c *Client) DescribeSetBlockUsageDayStatistics(request *DescribeSetBlockUsageDayStatisticsRequest) (response *DescribeSetBlockUsageDayStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeSetBlockUsageDayStatisticsRequest()
	}
	response = NewDescribeSetBlockUsageDayStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewBindAutoSnapshotPolicyRequest() (request *BindAutoSnapshotPolicyRequest) {
	request = &BindAutoSnapshotPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "BindAutoSnapshotPolicy")
	return
}

func NewBindAutoSnapshotPolicyResponse() (response *BindAutoSnapshotPolicyResponse) {
	response = &BindAutoSnapshotPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（BindAutoSnapshotPolicy）用于绑定云硬盘到指定的定期快照策略。
//
// * 每个地域最多可创建10个定期快照策略, 每个定期快照策略最多能绑定80个云硬盘。
// * 当已绑定定期快照策略的云硬盘处于未使用状态（即弹性云盘未挂载或非弹性云盘的主机处于关机状态）将不会创建定期快照。
func (c *Client) BindAutoSnapshotPolicy(request *BindAutoSnapshotPolicyRequest) (response *BindAutoSnapshotPolicyResponse, err error) {
	if request == nil {
		request = NewBindAutoSnapshotPolicyRequest()
	}
	response = NewBindAutoSnapshotPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewUninstallDiskStorageDepotCellRequest() (request *UninstallDiskStorageDepotCellRequest) {
	request = &UninstallDiskStorageDepotCellRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "UninstallDiskStorageDepotCell")
	return
}

func NewUninstallDiskStorageDepotCellResponse() (response *UninstallDiskStorageDepotCellResponse) {
	response = &UninstallDiskStorageDepotCellResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UninstallDiskStorageDepotCell）用于卸载云硬盘存储集群CELL组件。
func (c *Client) UninstallDiskStorageDepotCell(request *UninstallDiskStorageDepotCellRequest) (response *UninstallDiskStorageDepotCellResponse, err error) {
	if request == nil {
		request = NewUninstallDiskStorageDepotCellRequest()
	}
	response = NewUninstallDiskStorageDepotCellResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDepotDriverDetailsRequest() (request *DescribeDepotDriverDetailsRequest) {
	request = &DescribeDepotDriverDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDepotDriverDetails")
	return
}

func NewDescribeDepotDriverDetailsResponse() (response *DescribeDepotDriverDetailsResponse) {
	response = &DescribeDepotDriverDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDepotDriverDetails）用于查询当前存储池云盘挂载的母机列表详情。
// * 可以查询到每个母机上driver的版本号，状态；
// * 可以查询每个母机当前挂载了多少块云盘；
func (c *Client) DescribeDepotDriverDetails(request *DescribeDepotDriverDetailsRequest) (response *DescribeDepotDriverDetailsResponse, err error) {
	if request == nil {
		request = NewDescribeDepotDriverDetailsRequest()
	}
	response = NewDescribeDepotDriverDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewModifySnapshotAttributeRequest() (request *ModifySnapshotAttributeRequest) {
	request = &ModifySnapshotAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifySnapshotAttribute")
	return
}

func NewModifySnapshotAttributeResponse() (response *ModifySnapshotAttributeResponse) {
	response = &ModifySnapshotAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改运营端快照的属性。
// * 可将后台快照转为用户快照。
func (c *Client) ModifySnapshotAttribute(request *ModifySnapshotAttributeRequest) (response *ModifySnapshotAttributeResponse, err error) {
	if request == nil {
		request = NewModifySnapshotAttributeRequest()
	}
	response = NewModifySnapshotAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateDepotConfigRequest() (request *UpdateDepotConfigRequest) {
	request = &UpdateDepotConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "UpdateDepotConfig")
	return
}

func NewUpdateDepotConfigResponse() (response *UpdateDepotConfigResponse) {
	response = &UpdateDepotConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateDepotConfig）用于修改存储池的属性值，包括master、cell、dbtrsf等组件的配置属性。
func (c *Client) UpdateDepotConfig(request *UpdateDepotConfigRequest) (response *UpdateDepotConfigResponse, err error) {
	if request == nil {
		request = NewUpdateDepotConfigRequest()
	}
	response = NewUpdateDepotConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCancelClusterUpgradeTasksRequest() (request *CancelClusterUpgradeTasksRequest) {
	request = &CancelClusterUpgradeTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "CancelClusterUpgradeTasks")
	return
}

func NewCancelClusterUpgradeTasksResponse() (response *CancelClusterUpgradeTasksResponse) {
	response = &CancelClusterUpgradeTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CancelClusterUpgradeTasks）用于取消已发起的集群升级任务
// * 一次集群升级会包含若干个组件升级的子任务，只有子任务发起升级操作前才支持取消，如果有任何子任务已经在升级中，或升级完成，则本次集群升级任务不能取消；
// * 不支持取消单个组件的升级任务，只能批量取消一次集群升级的同一批升级子任务；
func (c *Client) CancelClusterUpgradeTasks(request *CancelClusterUpgradeTasksRequest) (response *CancelClusterUpgradeTasksResponse, err error) {
	if request == nil {
		request = NewCancelClusterUpgradeTasksRequest()
	}
	response = NewCancelClusterUpgradeTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDisksMonitorDataOverviewRequest() (request *DescribeDisksMonitorDataOverviewRequest) {
	request = &DescribeDisksMonitorDataOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDisksMonitorDataOverview")
	return
}

func NewDescribeDisksMonitorDataOverviewResponse() (response *DescribeDisksMonitorDataOverviewResponse) {
	response = &DescribeDisksMonitorDataOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDisksMonitorDataOverview）用于查询云硬盘相关资源预览监控数据。
func (c *Client) DescribeDisksMonitorDataOverview(request *DescribeDisksMonitorDataOverviewRequest) (response *DescribeDisksMonitorDataOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeDisksMonitorDataOverviewRequest()
	}
	response = NewDescribeDisksMonitorDataOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewReplaceDiskStorageZKRequest() (request *ReplaceDiskStorageZKRequest) {
	request = &ReplaceDiskStorageZKRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ReplaceDiskStorageZK")
	return
}

func NewReplaceDiskStorageZKResponse() (response *ReplaceDiskStorageZKResponse) {
	response = &ReplaceDiskStorageZKResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ReplaceDiskStorageZK）用于ZK节点出现整机故障，用新的机器将故障节点替换掉。
//
// * 仅flower的节点可以更换，且要求leader和另一个flower节点必须为正常状态
func (c *Client) ReplaceDiskStorageZK(request *ReplaceDiskStorageZKRequest) (response *ReplaceDiskStorageZKResponse, err error) {
	if request == nil {
		request = NewReplaceDiskStorageZKRequest()
	}
	response = NewReplaceDiskStorageZKResponse()
	err = c.Send(request, response)
	return
}

func NewInstallBoxDevicesRequest() (request *InstallBoxDevicesRequest) {
	request = &InstallBoxDevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "InstallBoxDevices")
	return
}

func NewInstallBoxDevicesResponse() (response *InstallBoxDevicesResponse) {
	response = &InstallBoxDevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// * 本接口（InstallBoxDevices）用于安装快照集群， 包括manager, scheduler, transfer
func (c *Client) InstallBoxDevices(request *InstallBoxDevicesRequest) (response *InstallBoxDevicesResponse, err error) {
	if request == nil {
		request = NewInstallBoxDevicesRequest()
	}
	response = NewInstallBoxDevicesResponse()
	err = c.Send(request, response)
	return
}

func NewInstallCapacityServerRequest() (request *InstallCapacityServerRequest) {
	request = &InstallCapacityServerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "InstallCapacityServer")
	return
}

func NewInstallCapacityServerResponse() (response *InstallCapacityServerResponse) {
	response = &InstallCapacityServerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用户安装容量server组件，为存储池提供容量管理功能。
func (c *Client) InstallCapacityServer(request *InstallCapacityServerRequest) (response *InstallCapacityServerResponse, err error) {
	if request == nil {
		request = NewInstallCapacityServerRequest()
	}
	response = NewInstallCapacityServerResponse()
	err = c.Send(request, response)
	return
}

func NewRejectDiskStorageDepotCellRequest() (request *RejectDiskStorageDepotCellRequest) {
	request = &RejectDiskStorageDepotCellRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "RejectDiskStorageDepotCell")
	return
}

func NewRejectDiskStorageDepotCellResponse() (response *RejectDiskStorageDepotCellResponse) {
	response = &RejectDiskStorageDepotCellResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（RejectDiskStorageDepotCell）用于剔除存储仓库cell节点或磁盘。
func (c *Client) RejectDiskStorageDepotCell(request *RejectDiskStorageDepotCellRequest) (response *RejectDiskStorageDepotCellResponse, err error) {
	if request == nil {
		request = NewRejectDiskStorageDepotCellRequest()
	}
	response = NewRejectDiskStorageDepotCellResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskResourceOverviewRequest() (request *DescribeDiskResourceOverviewRequest) {
	request = &DescribeDiskResourceOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskResourceOverview")
	return
}

func NewDescribeDiskResourceOverviewResponse() (response *DescribeDiskResourceOverviewResponse) {
	response = &DescribeDiskResourceOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskResourceOverview）用于查询用户在所有地域的云硬盘、快照及存储系统资源使用情况。
func (c *Client) DescribeDiskResourceOverview(request *DescribeDiskResourceOverviewRequest) (response *DescribeDiskResourceOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeDiskResourceOverviewRequest()
	}
	response = NewDescribeDiskResourceOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDiskStoragePoolGroupRequest() (request *CreateDiskStoragePoolGroupRequest) {
	request = &CreateDiskStoragePoolGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "CreateDiskStoragePoolGroup")
	return
}

func NewCreateDiskStoragePoolGroupResponse() (response *CreateDiskStoragePoolGroupResponse) {
	response = &CreateDiskStoragePoolGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于创建云硬盘存储池组。使用云硬盘存储池组，可以限制某一个用户或者让用户自行指定某些云硬盘使用特定类型的云硬盘存储池；可用于云硬盘资源隔离等场景。
func (c *Client) CreateDiskStoragePoolGroup(request *CreateDiskStoragePoolGroupRequest) (response *CreateDiskStoragePoolGroupResponse, err error) {
	if request == nil {
		request = NewCreateDiskStoragePoolGroupRequest()
	}
	response = NewCreateDiskStoragePoolGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDiskAttributesRequest() (request *ModifyDiskAttributesRequest) {
	request = &ModifyDiskAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifyDiskAttributes")
	return
}

func NewModifyDiskAttributesResponse() (response *ModifyDiskAttributesResponse) {
	response = &ModifyDiskAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyDiskAttributes）用于修改云硬盘属性。
func (c *Client) ModifyDiskAttributes(request *ModifyDiskAttributesRequest) (response *ModifyDiskAttributesResponse, err error) {
	if request == nil {
		request = NewModifyDiskAttributesRequest()
	}
	response = NewModifyDiskAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewBindDiskStorageDepotToDiskStoragePoolGroupRequest() (request *BindDiskStorageDepotToDiskStoragePoolGroupRequest) {
	request = &BindDiskStorageDepotToDiskStoragePoolGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "BindDiskStorageDepotToDiskStoragePoolGroup")
	return
}

func NewBindDiskStorageDepotToDiskStoragePoolGroupResponse() (response *BindDiskStorageDepotToDiskStoragePoolGroupResponse) {
	response = &BindDiskStorageDepotToDiskStoragePoolGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// BindDiskStorageDepotToDiskStoragePoolGroup
func (c *Client) BindDiskStorageDepotToDiskStoragePoolGroup(request *BindDiskStorageDepotToDiskStoragePoolGroupRequest) (response *BindDiskStorageDepotToDiskStoragePoolGroupResponse, err error) {
	if request == nil {
		request = NewBindDiskStorageDepotToDiskStoragePoolGroupRequest()
	}
	response = NewBindDiskStorageDepotToDiskStoragePoolGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDiskStoragePoolGroupRequest() (request *DeleteDiskStoragePoolGroupRequest) {
	request = &DeleteDiskStoragePoolGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DeleteDiskStoragePoolGroup")
	return
}

func NewDeleteDiskStoragePoolGroupResponse() (response *DeleteDiskStoragePoolGroupResponse) {
	response = &DeleteDiskStoragePoolGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于删除云硬盘存储资源池组；删除前需要解绑关联的云硬盘存储池
func (c *Client) DeleteDiskStoragePoolGroup(request *DeleteDiskStoragePoolGroupRequest) (response *DeleteDiskStoragePoolGroupResponse, err error) {
	if request == nil {
		request = NewDeleteDiskStoragePoolGroupRequest()
	}
	response = NewDeleteDiskStoragePoolGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSnapshotsRequest() (request *DeleteSnapshotsRequest) {
	request = &DeleteSnapshotsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DeleteSnapshots")
	return
}

func NewDeleteSnapshotsResponse() (response *DeleteSnapshotsResponse) {
	response = &DeleteSnapshotsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteSnapshots）用于删除快照。
//
// * 快照必须处于NORMAL状态，快照状态可以通过[DescribeSnapshots](/document/product/362/15647)接口查询，见输出参数中SnapshotState字段解释。
// * 支持批量操作。如果多个快照存在无法删除的快照，则操作不执行，以返回特定的错误码返回。
func (c *Client) DeleteSnapshots(request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
	if request == nil {
		request = NewDeleteSnapshotsRequest()
	}
	response = NewDeleteSnapshotsResponse()
	err = c.Send(request, response)
	return
}

func NewRetryDiskTaskRequest() (request *RetryDiskTaskRequest) {
	request = &RetryDiskTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "RetryDiskTask")
	return
}

func NewRetryDiskTaskResponse() (response *RetryDiskTaskResponse) {
	response = &RetryDiskTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（RetryDiskTask）用于重试云硬盘后台任务或者运维任务。
func (c *Client) RetryDiskTask(request *RetryDiskTaskRequest) (response *RetryDiskTaskResponse, err error) {
	if request == nil {
		request = NewRetryDiskTaskRequest()
	}
	response = NewRetryDiskTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDisksRequest() (request *DescribeDisksRequest) {
	request = &DescribeDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDisks")
	return
}

func NewDescribeDisksResponse() (response *DescribeDisksResponse) {
	response = &DescribeDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDisks）用于查询云硬盘列表。
//
// * 可以根据云硬盘ID、云硬盘类型或者云硬盘状态等信息来查询云硬盘的详细信息，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的云硬盘列表。
func (c *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
	if request == nil {
		request = NewDescribeDisksRequest()
	}
	response = NewDescribeDisksResponse()
	err = c.Send(request, response)
	return
}

func NewInstallTransferPairsRequest() (request *InstallTransferPairsRequest) {
	request = &InstallTransferPairsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "InstallTransferPairs")
	return
}

func NewInstallTransferPairsResponse() (response *InstallTransferPairsResponse) {
	response = &InstallTransferPairsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InstallTransferPairs）用于安装云硬盘Transfer Pairs，用于云盘双写迁移功能。
func (c *Client) InstallTransferPairs(request *InstallTransferPairsRequest) (response *InstallTransferPairsResponse, err error) {
	if request == nil {
		request = NewInstallTransferPairsRequest()
	}
	response = NewInstallTransferPairsResponse()
	err = c.Send(request, response)
	return
}

func NewUninstallTransferPairsRequest() (request *UninstallTransferPairsRequest) {
	request = &UninstallTransferPairsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "UninstallTransferPairs")
	return
}

func NewUninstallTransferPairsResponse() (response *UninstallTransferPairsResponse) {
	response = &UninstallTransferPairsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UninstallTransferPairs）用于卸载云硬盘Transfer Pair组件。
func (c *Client) UninstallTransferPairs(request *UninstallTransferPairsRequest) (response *UninstallTransferPairsResponse, err error) {
	if request == nil {
		request = NewUninstallTransferPairsRequest()
	}
	response = NewUninstallTransferPairsResponse()
	err = c.Send(request, response)
	return
}

func NewCancelDepotTransferTasksRequest() (request *CancelDepotTransferTasksRequest) {
	request = &CancelDepotTransferTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "CancelDepotTransferTasks")
	return
}

func NewCancelDepotTransferTasksResponse() (response *CancelDepotTransferTasksResponse) {
	response = &CancelDepotTransferTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于取消存储池小表迁移任务：
// * 只有出现异常导致长时间未结束的任务才能取消，正常迁移的任务不需要取消
func (c *Client) CancelDepotTransferTasks(request *CancelDepotTransferTasksRequest) (response *CancelDepotTransferTasksResponse, err error) {
	if request == nil {
		request = NewCancelDepotTransferTasksRequest()
	}
	response = NewCancelDepotTransferTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskConfigForSaleRequest() (request *DescribeDiskConfigForSaleRequest) {
	request = &DescribeDiskConfigForSaleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskConfigForSale")
	return
}

func NewDescribeDiskConfigForSaleResponse() (response *DescribeDiskConfigForSaleResponse) {
	response = &DescribeDiskConfigForSaleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskConfigForSale）用于查询云硬盘可售卖配置。
func (c *Client) DescribeDiskConfigForSale(request *DescribeDiskConfigForSaleRequest) (response *DescribeDiskConfigForSaleResponse, err error) {
	if request == nil {
		request = NewDescribeDiskConfigForSaleRequest()
	}
	response = NewDescribeDiskConfigForSaleResponse()
	err = c.Send(request, response)
	return
}

func NewUnbindDiskStorageDepotFromDiskStoragePoolGroupRequest() (request *UnbindDiskStorageDepotFromDiskStoragePoolGroupRequest) {
	request = &UnbindDiskStorageDepotFromDiskStoragePoolGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "UnbindDiskStorageDepotFromDiskStoragePoolGroup")
	return
}

func NewUnbindDiskStorageDepotFromDiskStoragePoolGroupResponse() (response *UnbindDiskStorageDepotFromDiskStoragePoolGroupResponse) {
	response = &UnbindDiskStorageDepotFromDiskStoragePoolGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从云硬盘存储仓库资源池中解绑存储仓库
func (c *Client) UnbindDiskStorageDepotFromDiskStoragePoolGroup(request *UnbindDiskStorageDepotFromDiskStoragePoolGroupRequest) (response *UnbindDiskStorageDepotFromDiskStoragePoolGroupResponse, err error) {
	if request == nil {
		request = NewUnbindDiskStorageDepotFromDiskStoragePoolGroupRequest()
	}
	response = NewUnbindDiskStorageDepotFromDiskStoragePoolGroupResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSnapshotBoxRequest() (request *CreateSnapshotBoxRequest) {
	request = &CreateSnapshotBoxRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "CreateSnapshotBox")
	return
}

func NewCreateSnapshotBoxResponse() (response *CreateSnapshotBoxResponse) {
	response = &CreateSnapshotBoxResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateSnapshotBox）创建云硬盘快照集群
func (c *Client) CreateSnapshotBox(request *CreateSnapshotBoxRequest) (response *CreateSnapshotBoxResponse, err error) {
	if request == nil {
		request = NewCreateSnapshotBoxRequest()
	}
	response = NewCreateSnapshotBoxResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskStorageDepotDetailsRequest() (request *DescribeDiskStorageDepotDetailsRequest) {
	request = &DescribeDiskStorageDepotDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskStorageDepotDetails")
	return
}

func NewDescribeDiskStorageDepotDetailsResponse() (response *DescribeDiskStorageDepotDetailsResponse) {
	response = &DescribeDiskStorageDepotDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskStorageDepotDetails）用于查询云硬盘存储仓库详情
func (c *Client) DescribeDiskStorageDepotDetails(request *DescribeDiskStorageDepotDetailsRequest) (response *DescribeDiskStorageDepotDetailsResponse, err error) {
	if request == nil {
		request = NewDescribeDiskStorageDepotDetailsRequest()
	}
	response = NewDescribeDiskStorageDepotDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBoxAttributesRequest() (request *ModifyBoxAttributesRequest) {
	request = &ModifyBoxAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifyBoxAttributes")
	return
}

func NewModifyBoxAttributesResponse() (response *ModifyBoxAttributesResponse) {
	response = &ModifyBoxAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyBoxAttributes）用于修改快照集群属性。
func (c *Client) ModifyBoxAttributes(request *ModifyBoxAttributesRequest) (response *ModifyBoxAttributesResponse, err error) {
	if request == nil {
		request = NewModifyBoxAttributesRequest()
	}
	response = NewModifyBoxAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAutoSnapshotPolicyAttributeRequest() (request *ModifyAutoSnapshotPolicyAttributeRequest) {
	request = &ModifyAutoSnapshotPolicyAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifyAutoSnapshotPolicyAttribute")
	return
}

func NewModifyAutoSnapshotPolicyAttributeResponse() (response *ModifyAutoSnapshotPolicyAttributeResponse) {
	response = &ModifyAutoSnapshotPolicyAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyAutoSnapshotPolicyAttribute）用于修改定期快照策略属性。
//
// * 可通过该接口修改定期快照策略的执行策略、名称、是否激活等属性。
// * 修改保留天数时必须保证不与是否永久保留属性冲突，否则整个操作失败，以特定的错误码返回。
func (c *Client) ModifyAutoSnapshotPolicyAttribute(request *ModifyAutoSnapshotPolicyAttributeRequest) (response *ModifyAutoSnapshotPolicyAttributeResponse, err error) {
	if request == nil {
		request = NewModifyAutoSnapshotPolicyAttributeRequest()
	}
	response = NewModifyAutoSnapshotPolicyAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewTransferDiskStorageDepotCellDiskPairRequest() (request *TransferDiskStorageDepotCellDiskPairRequest) {
	request = &TransferDiskStorageDepotCellDiskPairRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "TransferDiskStorageDepotCellDiskPair")
	return
}

func NewTransferDiskStorageDepotCellDiskPairResponse() (response *TransferDiskStorageDepotCellDiskPairResponse) {
	response = &TransferDiskStorageDepotCellDiskPairResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（TransferDiskStorageDepotCellDiskPair）用于对云硬盘存储集群节点或物理硬盘执行迁移操作
func (c *Client) TransferDiskStorageDepotCellDiskPair(request *TransferDiskStorageDepotCellDiskPairRequest) (response *TransferDiskStorageDepotCellDiskPairResponse, err error) {
	if request == nil {
		request = NewTransferDiskStorageDepotCellDiskPairRequest()
	}
	response = NewTransferDiskStorageDepotCellDiskPairResponse()
	err = c.Send(request, response)
	return
}

func NewInitializeDiskStorageDepotRequest() (request *InitializeDiskStorageDepotRequest) {
	request = &InitializeDiskStorageDepotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "InitializeDiskStorageDepot")
	return
}

func NewInitializeDiskStorageDepotResponse() (response *InitializeDiskStorageDepotResponse) {
	response = &InitializeDiskStorageDepotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InitializeDiskStorageDepot）用于初始化云盘存储集群。存储集群上架完成后需要先进行初始化才能打开售卖。
func (c *Client) InitializeDiskStorageDepot(request *InitializeDiskStorageDepotRequest) (response *InitializeDiskStorageDepotResponse, err error) {
	if request == nil {
		request = NewInitializeDiskStorageDepotRequest()
	}
	response = NewInitializeDiskStorageDepotResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCbsAlarmEventsRequest() (request *DescribeCbsAlarmEventsRequest) {
	request = &DescribeCbsAlarmEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeCbsAlarmEvents")
	return
}

func NewDescribeCbsAlarmEventsResponse() (response *DescribeCbsAlarmEventsResponse) {
	response = &DescribeCbsAlarmEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeCbsAlarmEvents）用于查询云硬盘告警事件。
func (c *Client) DescribeCbsAlarmEvents(request *DescribeCbsAlarmEventsRequest) (response *DescribeCbsAlarmEventsResponse, err error) {
	if request == nil {
		request = NewDescribeCbsAlarmEventsRequest()
	}
	response = NewDescribeCbsAlarmEventsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDiskStorageDepotRequest() (request *CreateDiskStorageDepotRequest) {
	request = &CreateDiskStorageDepotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "CreateDiskStorageDepot")
	return
}

func NewCreateDiskStorageDepotResponse() (response *CreateDiskStorageDepotResponse) {
	response = &CreateDiskStorageDepotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateDiskStorageDepot）用于创建云硬盘存储池。
func (c *Client) CreateDiskStorageDepot(request *CreateDiskStorageDepotRequest) (response *CreateDiskStorageDepotResponse, err error) {
	if request == nil {
		request = NewCreateDiskStorageDepotRequest()
	}
	response = NewCreateDiskStorageDepotResponse()
	err = c.Send(request, response)
	return
}

func NewAttachDisksRequest() (request *AttachDisksRequest) {
	request = &AttachDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "AttachDisks")
	return
}

func NewAttachDisksResponse() (response *AttachDisksResponse) {
	response = &AttachDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（AttachDisks）用于挂载云硬盘。
//
// * 支持批量操作，将多块云盘挂载到同一云主机。如果多个云盘存在不允许挂载的云盘，则操作不执行，以返回特定的错误码返回。
// * 本接口为异步接口，当挂载云盘的请求成功返回时，表示后台已发起挂载云盘的操作，可通过接口[DescribeDisks](/document/product/362/16315)来查询对应云盘的状态，如果云盘的状态由“ATTACHING”变为“ATTACHED”，则为挂载成功。
func (c *Client) AttachDisks(request *AttachDisksRequest) (response *AttachDisksResponse, err error) {
	if request == nil {
		request = NewAttachDisksRequest()
	}
	response = NewAttachDisksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomerCapacityTopChangesRequest() (request *DescribeCustomerCapacityTopChangesRequest) {
	request = &DescribeCustomerCapacityTopChangesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeCustomerCapacityTopChanges")
	return
}

func NewDescribeCustomerCapacityTopChangesResponse() (response *DescribeCustomerCapacityTopChangesResponse) {
	response = &DescribeCustomerCapacityTopChangesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户规模变化top统计
func (c *Client) DescribeCustomerCapacityTopChanges(request *DescribeCustomerCapacityTopChangesRequest) (response *DescribeCustomerCapacityTopChangesResponse, err error) {
	if request == nil {
		request = NewDescribeCustomerCapacityTopChangesRequest()
	}
	response = NewDescribeCustomerCapacityTopChangesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskStorageDepotAttributesRequest() (request *DescribeDiskStorageDepotAttributesRequest) {
	request = &DescribeDiskStorageDepotAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskStorageDepotAttributes")
	return
}

func NewDescribeDiskStorageDepotAttributesResponse() (response *DescribeDiskStorageDepotAttributesResponse) {
	response = &DescribeDiskStorageDepotAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskStorageDepotAttributes）用于查询云硬盘存储仓库可用属性配置信息
func (c *Client) DescribeDiskStorageDepotAttributes(request *DescribeDiskStorageDepotAttributesRequest) (response *DescribeDiskStorageDepotAttributesResponse, err error) {
	if request == nil {
		request = NewDescribeDiskStorageDepotAttributesRequest()
	}
	response = NewDescribeDiskStorageDepotAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskMigrateTasksRequest() (request *DescribeDiskMigrateTasksRequest) {
	request = &DescribeDiskMigrateTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskMigrateTasks")
	return
}

func NewDescribeDiskMigrateTasksResponse() (response *DescribeDiskMigrateTasksResponse) {
	response = &DescribeDiskMigrateTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskMigrateTasks）用于查询云硬盘迁移任务
func (c *Client) DescribeDiskMigrateTasks(request *DescribeDiskMigrateTasksRequest) (response *DescribeDiskMigrateTasksResponse, err error) {
	if request == nil {
		request = NewDescribeDiskMigrateTasksRequest()
	}
	response = NewDescribeDiskMigrateTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
	request = &DescribeSnapshotsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshots")
	return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
	response = &DescribeSnapshotsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSnapshots）用于查询快照的详细信息。
//
// * 根据快照ID、创建快照的云硬盘ID、创建快照的云硬盘类型等对结果进行过滤，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
// *  如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的快照列表。
func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
	if request == nil {
		request = NewDescribeSnapshotsRequest()
	}
	response = NewDescribeSnapshotsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSnapshotBoxRequest() (request *DeleteSnapshotBoxRequest) {
	request = &DeleteSnapshotBoxRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DeleteSnapshotBox")
	return
}

func NewDeleteSnapshotBoxResponse() (response *DeleteSnapshotBoxResponse) {
	response = &DeleteSnapshotBoxResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteSnapshotBox）删除云硬盘快照集群。
func (c *Client) DeleteSnapshotBox(request *DeleteSnapshotBoxRequest) (response *DeleteSnapshotBoxResponse, err error) {
	if request == nil {
		request = NewDeleteSnapshotBoxRequest()
	}
	response = NewDeleteSnapshotBoxResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskAssociatedAutoSnapshotPolicyRequest() (request *DescribeDiskAssociatedAutoSnapshotPolicyRequest) {
	request = &DescribeDiskAssociatedAutoSnapshotPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskAssociatedAutoSnapshotPolicy")
	return
}

func NewDescribeDiskAssociatedAutoSnapshotPolicyResponse() (response *DescribeDiskAssociatedAutoSnapshotPolicyResponse) {
	response = &DescribeDiskAssociatedAutoSnapshotPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskAssociatedAutoSnapshotPolicy）用于查询云盘绑定的定期快照策略。
func (c *Client) DescribeDiskAssociatedAutoSnapshotPolicy(request *DescribeDiskAssociatedAutoSnapshotPolicyRequest) (response *DescribeDiskAssociatedAutoSnapshotPolicyResponse, err error) {
	if request == nil {
		request = NewDescribeDiskAssociatedAutoSnapshotPolicyRequest()
	}
	response = NewDescribeDiskAssociatedAutoSnapshotPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDisksMonitorAlarmRulesRequest() (request *DescribeDisksMonitorAlarmRulesRequest) {
	request = &DescribeDisksMonitorAlarmRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDisksMonitorAlarmRules")
	return
}

func NewDescribeDisksMonitorAlarmRulesResponse() (response *DescribeDisksMonitorAlarmRulesResponse) {
	response = &DescribeDisksMonitorAlarmRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询判断云盘IO告警的默认告警策略，包括云盘iohang、svctm_high和存储池svctm_high告警
func (c *Client) DescribeDisksMonitorAlarmRules(request *DescribeDisksMonitorAlarmRulesRequest) (response *DescribeDisksMonitorAlarmRulesResponse, err error) {
	if request == nil {
		request = NewDescribeDisksMonitorAlarmRulesRequest()
	}
	response = NewDescribeDisksMonitorAlarmRulesResponse()
	err = c.Send(request, response)
	return
}

func NewUnbindAutoSnapshotPolicyRequest() (request *UnbindAutoSnapshotPolicyRequest) {
	request = &UnbindAutoSnapshotPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "UnbindAutoSnapshotPolicy")
	return
}

func NewUnbindAutoSnapshotPolicyResponse() (response *UnbindAutoSnapshotPolicyResponse) {
	response = &UnbindAutoSnapshotPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UnbindAutoSnapshotPolicy）用于解除云硬盘绑定的定期快照策略。
//
// * 支持批量操作，可一次解除多个云盘与同一定期快照策略的绑定。
// * 如果传入的云盘未绑定到当前定期快照策略，接口将自动跳过，仅解绑与当前定期快照策略绑定的云盘。
func (c *Client) UnbindAutoSnapshotPolicy(request *UnbindAutoSnapshotPolicyRequest) (response *UnbindAutoSnapshotPolicyResponse, err error) {
	if request == nil {
		request = NewUnbindAutoSnapshotPolicyRequest()
	}
	response = NewUnbindAutoSnapshotPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDepotFlowControlRecordsRequest() (request *DescribeDepotFlowControlRecordsRequest) {
	request = &DescribeDepotFlowControlRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDepotFlowControlRecords")
	return
}

func NewDescribeDepotFlowControlRecordsResponse() (response *DescribeDepotFlowControlRecordsResponse) {
	response = &DescribeDepotFlowControlRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDepotFlowControlRecords）用于查询云硬盘存储仓库流控历史。
func (c *Client) DescribeDepotFlowControlRecords(request *DescribeDepotFlowControlRecordsRequest) (response *DescribeDepotFlowControlRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeDepotFlowControlRecordsRequest()
	}
	response = NewDescribeDepotFlowControlRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskTasksRequest() (request *DescribeDiskTasksRequest) {
	request = &DescribeDiskTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskTasks")
	return
}

func NewDescribeDiskTasksResponse() (response *DescribeDiskTasksResponse) {
	response = &DescribeDiskTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskTasks）用于查询云硬盘操作记录。
func (c *Client) DescribeDiskTasks(request *DescribeDiskTasksRequest) (response *DescribeDiskTasksResponse, err error) {
	if request == nil {
		request = NewDescribeDiskTasksRequest()
	}
	response = NewDescribeDiskTasksResponse()
	err = c.Send(request, response)
	return
}

func NewInstallDiskStorageDepotCellRequest() (request *InstallDiskStorageDepotCellRequest) {
	request = &InstallDiskStorageDepotCellRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "InstallDiskStorageDepotCell")
	return
}

func NewInstallDiskStorageDepotCellResponse() (response *InstallDiskStorageDepotCellResponse) {
	response = &InstallDiskStorageDepotCellResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InstallDiskStorageDepotCell）用于安装云硬盘存储池Cell组件
func (c *Client) InstallDiskStorageDepotCell(request *InstallDiskStorageDepotCellRequest) (response *InstallDiskStorageDepotCellResponse, err error) {
	if request == nil {
		request = NewInstallDiskStorageDepotCellRequest()
	}
	response = NewInstallDiskStorageDepotCellResponse()
	err = c.Send(request, response)
	return
}

func NewUninstallDiskZKClusterRequest() (request *UninstallDiskZKClusterRequest) {
	request = &UninstallDiskZKClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "UninstallDiskZKCluster")
	return
}

func NewUninstallDiskZKClusterResponse() (response *UninstallDiskZKClusterResponse) {
	response = &UninstallDiskZKClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UninstallDiskZKCluster）用于卸载云硬盘ZooKeeper集群。
func (c *Client) UninstallDiskZKCluster(request *UninstallDiskZKClusterRequest) (response *UninstallDiskZKClusterResponse, err error) {
	if request == nil {
		request = NewUninstallDiskZKClusterRequest()
	}
	response = NewUninstallDiskZKClusterResponse()
	err = c.Send(request, response)
	return
}

func NewMigrateDisksRequest() (request *MigrateDisksRequest) {
	request = &MigrateDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "MigrateDisks")
	return
}

func NewMigrateDisksResponse() (response *MigrateDisksResponse) {
	response = &MigrateDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（MigrateDisks）用于创建云硬盘迁移任务。
func (c *Client) MigrateDisks(request *MigrateDisksRequest) (response *MigrateDisksResponse, err error) {
	if request == nil {
		request = NewMigrateDisksRequest()
	}
	response = NewMigrateDisksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSetCellSideIOStatRequest() (request *DescribeSetCellSideIOStatRequest) {
	request = &DescribeSetCellSideIOStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeSetCellSideIOStat")
	return
}

func NewDescribeSetCellSideIOStatResponse() (response *DescribeSetCellSideIOStatResponse) {
	response = &DescribeSetCellSideIOStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询由存储池内所有存储节点IO监控数据汇聚后的存储池IO监控数据
func (c *Client) DescribeSetCellSideIOStat(request *DescribeSetCellSideIOStatRequest) (response *DescribeSetCellSideIOStatResponse, err error) {
	if request == nil {
		request = NewDescribeSetCellSideIOStatRequest()
	}
	response = NewDescribeSetCellSideIOStatResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSnapshotBoxOverViewRequest() (request *DescribeSnapshotBoxOverViewRequest) {
	request = &DescribeSnapshotBoxOverViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshotBoxOverView")
	return
}

func NewDescribeSnapshotBoxOverViewResponse() (response *DescribeSnapshotBoxOverViewResponse) {
	response = &DescribeSnapshotBoxOverViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询快照集群各状态快照的大小和数量
func (c *Client) DescribeSnapshotBoxOverView(request *DescribeSnapshotBoxOverViewRequest) (response *DescribeSnapshotBoxOverViewResponse, err error) {
	if request == nil {
		request = NewDescribeSnapshotBoxOverViewRequest()
	}
	response = NewDescribeSnapshotBoxOverViewResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDepotCellDeviceConfigRequest() (request *CreateDepotCellDeviceConfigRequest) {
	request = &CreateDepotCellDeviceConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "CreateDepotCellDeviceConfig")
	return
}

func NewCreateDepotCellDeviceConfigResponse() (response *CreateDepotCellDeviceConfigResponse) {
	response = &CreateDepotCellDeviceConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于新增存储池机型配置
func (c *Client) CreateDepotCellDeviceConfig(request *CreateDepotCellDeviceConfigRequest) (response *CreateDepotCellDeviceConfigResponse, err error) {
	if request == nil {
		request = NewCreateDepotCellDeviceConfigRequest()
	}
	response = NewCreateDepotCellDeviceConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomerCapacityTopTrendsRequest() (request *DescribeCustomerCapacityTopTrendsRequest) {
	request = &DescribeCustomerCapacityTopTrendsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeCustomerCapacityTopTrends")
	return
}

func NewDescribeCustomerCapacityTopTrendsResponse() (response *DescribeCustomerCapacityTopTrendsResponse) {
	response = &DescribeCustomerCapacityTopTrendsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户维度的产品规模变化及增退变化趋势
func (c *Client) DescribeCustomerCapacityTopTrends(request *DescribeCustomerCapacityTopTrendsRequest) (response *DescribeCustomerCapacityTopTrendsResponse, err error) {
	if request == nil {
		request = NewDescribeCustomerCapacityTopTrendsRequest()
	}
	response = NewDescribeCustomerCapacityTopTrendsResponse()
	err = c.Send(request, response)
	return
}

func NewBalanceDiskStorageDepotRequest() (request *BalanceDiskStorageDepotRequest) {
	request = &BalanceDiskStorageDepotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "BalanceDiskStorageDepot")
	return
}

func NewBalanceDiskStorageDepotResponse() (response *BalanceDiskStorageDepotResponse) {
	response = &BalanceDiskStorageDepotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（BalanceDiskStorageDepot）用于对云硬盘存储池执行小表一键均衡操作
func (c *Client) BalanceDiskStorageDepot(request *BalanceDiskStorageDepotRequest) (response *BalanceDiskStorageDepotResponse, err error) {
	if request == nil {
		request = NewBalanceDiskStorageDepotRequest()
	}
	response = NewBalanceDiskStorageDepotResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSnapshotResourceUsageRequest() (request *DescribeSnapshotResourceUsageRequest) {
	request = &DescribeSnapshotResourceUsageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshotResourceUsage")
	return
}

func NewDescribeSnapshotResourceUsageResponse() (response *DescribeSnapshotResourceUsageResponse) {
	response = &DescribeSnapshotResourceUsageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询快照资源每日使用量统计
func (c *Client) DescribeSnapshotResourceUsage(request *DescribeSnapshotResourceUsageRequest) (response *DescribeSnapshotResourceUsageResponse, err error) {
	if request == nil {
		request = NewDescribeSnapshotResourceUsageRequest()
	}
	response = NewDescribeSnapshotResourceUsageResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDepotCellDeviceConfigRequest() (request *DeleteDepotCellDeviceConfigRequest) {
	request = &DeleteDepotCellDeviceConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DeleteDepotCellDeviceConfig")
	return
}

func NewDeleteDepotCellDeviceConfigResponse() (response *DeleteDepotCellDeviceConfigResponse) {
	response = &DeleteDepotCellDeviceConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用接口用于删除存储池机型配置
func (c *Client) DeleteDepotCellDeviceConfig(request *DeleteDepotCellDeviceConfigRequest) (response *DeleteDepotCellDeviceConfigResponse, err error) {
	if request == nil {
		request = NewDeleteDepotCellDeviceConfigRequest()
	}
	response = NewDeleteDepotCellDeviceConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSnapshotSharePermissionRequest() (request *DescribeSnapshotSharePermissionRequest) {
	request = &DescribeSnapshotSharePermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshotSharePermission")
	return
}

func NewDescribeSnapshotSharePermissionResponse() (response *DescribeSnapshotSharePermissionResponse) {
	response = &DescribeSnapshotSharePermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSnapshotSharePermission）用于查询快照的分享信息。
func (c *Client) DescribeSnapshotSharePermission(request *DescribeSnapshotSharePermissionRequest) (response *DescribeSnapshotSharePermissionResponse, err error) {
	if request == nil {
		request = NewDescribeSnapshotSharePermissionRequest()
	}
	response = NewDescribeSnapshotSharePermissionResponse()
	err = c.Send(request, response)
	return
}

func NewUninstallDiskStorageDepotMasterRequest() (request *UninstallDiskStorageDepotMasterRequest) {
	request = &UninstallDiskStorageDepotMasterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "UninstallDiskStorageDepotMaster")
	return
}

func NewUninstallDiskStorageDepotMasterResponse() (response *UninstallDiskStorageDepotMasterResponse) {
	response = &UninstallDiskStorageDepotMasterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UninstallDiskStorageDepotMaster）用于卸载云硬盘存储集群Master组件。
// * 卸载master前，需先卸载所有cell节点；
// * 卸载master时，要先卸载备节点，再卸载主节点。
func (c *Client) UninstallDiskStorageDepotMaster(request *UninstallDiskStorageDepotMasterRequest) (response *UninstallDiskStorageDepotMasterResponse, err error) {
	if request == nil {
		request = NewUninstallDiskStorageDepotMasterRequest()
	}
	response = NewUninstallDiskStorageDepotMasterResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAutoCloseAndOpenBlockRateRequest() (request *ModifyAutoCloseAndOpenBlockRateRequest) {
	request = &ModifyAutoCloseAndOpenBlockRateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifyAutoCloseAndOpenBlockRate")
	return
}

func NewModifyAutoCloseAndOpenBlockRateResponse() (response *ModifyAutoCloseAndOpenBlockRateResponse) {
	response = &ModifyAutoCloseAndOpenBlockRateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改自动打开和自动关闭仓库售卖的block利用率阈值：
// * 当传入ZoneId和DepotId会修改对应仓库的block利用率配置;
// * 当不传入任何参数时，会修改全局的block利用率配置;
func (c *Client) ModifyAutoCloseAndOpenBlockRate(request *ModifyAutoCloseAndOpenBlockRateRequest) (response *ModifyAutoCloseAndOpenBlockRateResponse, err error) {
	if request == nil {
		request = NewModifyAutoCloseAndOpenBlockRateRequest()
	}
	response = NewModifyAutoCloseAndOpenBlockRateResponse()
	err = c.Send(request, response)
	return
}

func NewRunClusterUpgradeTasksRequest() (request *RunClusterUpgradeTasksRequest) {
	request = &RunClusterUpgradeTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "RunClusterUpgradeTasks")
	return
}

func NewRunClusterUpgradeTasksResponse() (response *RunClusterUpgradeTasksResponse) {
	response = &RunClusterUpgradeTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（RunClusterUpgradeTasks）用于发起存储池或快照集群各组件的升级任务。
// * 本接口支持DryRun参数，传入为1时，表示只进行环境检查，不执行实际升级操作；检查成功后，任务状态会流转为CHECK_FINISH；
// * 集群各组件的升级是有顺序要求的，必须按接口DescribeClusterUpgradeTasks返回的顺序进行升级
// * 每个任务执行实际升级操作前，必须执行环境检查，待任务变成CHECK_FINISH后，才能发起实际升级操作
func (c *Client) RunClusterUpgradeTasks(request *RunClusterUpgradeTasksRequest) (response *RunClusterUpgradeTasksResponse, err error) {
	if request == nil {
		request = NewRunClusterUpgradeTasksRequest()
	}
	response = NewRunClusterUpgradeTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDepotTransferTasksRequest() (request *DescribeDepotTransferTasksRequest) {
	request = &DescribeDepotTransferTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDepotTransferTasks")
	return
}

func NewDescribeDepotTransferTasksResponse() (response *DescribeDepotTransferTasksResponse) {
	response = &DescribeDepotTransferTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDepotTransferTasks）用查询存储池小表迁移任务。
// * 目前只能查到正在迁移中的任务，已结束的历史任务无法查询。
func (c *Client) DescribeDepotTransferTasks(request *DescribeDepotTransferTasksRequest) (response *DescribeDepotTransferTasksResponse, err error) {
	if request == nil {
		request = NewDescribeDepotTransferTasksRequest()
	}
	response = NewDescribeDepotTransferTasksResponse()
	err = c.Send(request, response)
	return
}

func NewRecoverDiskRequest() (request *RecoverDiskRequest) {
	request = &RecoverDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "RecoverDisk")
	return
}

func NewRecoverDiskResponse() (response *RecoverDiskResponse) {
	response = &RecoverDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// * 本接口用于将已删除未收回的云硬盘再恢复成快照，在误删云盘的情况下，可以再找回数据；
// * 仅能恢复已删除未回收的云盘，如果盘在底层已回收，则无法恢复；
// * 恢复出的快照可以在快照列表页直接看到，且会正常计费；
func (c *Client) RecoverDisk(request *RecoverDiskRequest) (response *RecoverDiskResponse, err error) {
	if request == nil {
		request = NewRecoverDiskRequest()
	}
	response = NewRecoverDiskResponse()
	err = c.Send(request, response)
	return
}

func NewReplaceDiskStorageDepotCellRequest() (request *ReplaceDiskStorageDepotCellRequest) {
	request = &ReplaceDiskStorageDepotCellRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ReplaceDiskStorageDepotCell")
	return
}

func NewReplaceDiskStorageDepotCellResponse() (response *ReplaceDiskStorageDepotCellResponse) {
	response = &ReplaceDiskStorageDepotCellResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ReplaceDiskStorageDepotCell）用于CBS存储仓库cell节点出现故障时，替换整台cell机器。
func (c *Client) ReplaceDiskStorageDepotCell(request *ReplaceDiskStorageDepotCellRequest) (response *ReplaceDiskStorageDepotCellResponse, err error) {
	if request == nil {
		request = NewReplaceDiskStorageDepotCellRequest()
	}
	response = NewReplaceDiskStorageDepotCellResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDiskConfigForSaleRequest() (request *ModifyDiskConfigForSaleRequest) {
	request = &ModifyDiskConfigForSaleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifyDiskConfigForSale")
	return
}

func NewModifyDiskConfigForSaleResponse() (response *ModifyDiskConfigForSaleResponse) {
	response = &ModifyDiskConfigForSaleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyDiskConfigForSale）用于云硬盘可售卖配置。
func (c *Client) ModifyDiskConfigForSale(request *ModifyDiskConfigForSaleRequest) (response *ModifyDiskConfigForSaleResponse, err error) {
	if request == nil {
		request = NewModifyDiskConfigForSaleRequest()
	}
	response = NewModifyDiskConfigForSaleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDiskStorageDepotRequest() (request *DeleteDiskStorageDepotRequest) {
	request = &DeleteDiskStorageDepotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DeleteDiskStorageDepot")
	return
}

func NewDeleteDiskStorageDepotResponse() (response *DeleteDiskStorageDepotResponse) {
	response = &DeleteDiskStorageDepotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteDiskStorageDepot）用于删除云硬盘存储集群。
func (c *Client) DeleteDiskStorageDepot(request *DeleteDiskStorageDepotRequest) (response *DeleteDiskStorageDepotResponse, err error) {
	if request == nil {
		request = NewDeleteDiskStorageDepotRequest()
	}
	response = NewDeleteDiskStorageDepotResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSnapshotGroupsRequest() (request *DescribeSnapshotGroupsRequest) {
	request = &DescribeSnapshotGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshotGroups")
	return
}

func NewDescribeSnapshotGroupsResponse() (response *DescribeSnapshotGroupsResponse) {
	response = &DescribeSnapshotGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询快照组列表，返回状态正常、创建中、回滚中的快照组列表
func (c *Client) DescribeSnapshotGroups(request *DescribeSnapshotGroupsRequest) (response *DescribeSnapshotGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeSnapshotGroupsRequest()
	}
	response = NewDescribeSnapshotGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewUninstallCapacityServerRequest() (request *UninstallCapacityServerRequest) {
	request = &UninstallCapacityServerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "UninstallCapacityServer")
	return
}

func NewUninstallCapacityServerResponse() (response *UninstallCapacityServerResponse) {
	response = &UninstallCapacityServerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于卸载容量server组件
func (c *Client) UninstallCapacityServer(request *UninstallCapacityServerRequest) (response *UninstallCapacityServerResponse, err error) {
	if request == nil {
		request = NewUninstallCapacityServerRequest()
	}
	response = NewUninstallCapacityServerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSnapshotBoxsRequest() (request *DescribeSnapshotBoxsRequest) {
	request = &DescribeSnapshotBoxsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshotBoxs")
	return
}

func NewDescribeSnapshotBoxsResponse() (response *DescribeSnapshotBoxsResponse) {
	response = &DescribeSnapshotBoxsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSnapshotBoxs）用于查询快照集群的详细信息。
func (c *Client) DescribeSnapshotBoxs(request *DescribeSnapshotBoxsRequest) (response *DescribeSnapshotBoxsResponse, err error) {
	if request == nil {
		request = NewDescribeSnapshotBoxsRequest()
	}
	response = NewDescribeSnapshotBoxsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMonitorEsIndexStorageTimeRequest() (request *ModifyMonitorEsIndexStorageTimeRequest) {
	request = &ModifyMonitorEsIndexStorageTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifyMonitorEsIndexStorageTime")
	return
}

func NewModifyMonitorEsIndexStorageTimeResponse() (response *ModifyMonitorEsIndexStorageTimeResponse) {
	response = &ModifyMonitorEsIndexStorageTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改云盘和存储池的监控数据保留时间
func (c *Client) ModifyMonitorEsIndexStorageTime(request *ModifyMonitorEsIndexStorageTimeRequest) (response *ModifyMonitorEsIndexStorageTimeResponse, err error) {
	if request == nil {
		request = NewModifyMonitorEsIndexStorageTimeRequest()
	}
	response = NewModifyMonitorEsIndexStorageTimeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAutoCloseAndOpenBlockRateRequest() (request *DescribeAutoCloseAndOpenBlockRateRequest) {
	request = &DescribeAutoCloseAndOpenBlockRateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeAutoCloseAndOpenBlockRate")
	return
}

func NewDescribeAutoCloseAndOpenBlockRateResponse() (response *DescribeAutoCloseAndOpenBlockRateResponse) {
	response = &DescribeAutoCloseAndOpenBlockRateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询TCE环境自动打开和关闭仓库售卖的block利用率
// * 当传入ZoneId和DepotId会返回对应仓库的block利用率配置，若该仓库没有配置过，则返回全局的block利用率配置；
// * 当不传入任何参数时，会返回全局的block利用率配置
func (c *Client) DescribeAutoCloseAndOpenBlockRate(request *DescribeAutoCloseAndOpenBlockRateRequest) (response *DescribeAutoCloseAndOpenBlockRateResponse, err error) {
	if request == nil {
		request = NewDescribeAutoCloseAndOpenBlockRateRequest()
	}
	response = NewDescribeAutoCloseAndOpenBlockRateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserDiskUsageRequest() (request *DescribeUserDiskUsageRequest) {
	request = &DescribeUserDiskUsageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeUserDiskUsage")
	return
}

func NewDescribeUserDiskUsageResponse() (response *DescribeUserDiskUsageResponse) {
	response = &DescribeUserDiskUsageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于查询按用户维度的云硬盘相关资源的使用统计
func (c *Client) DescribeUserDiskUsage(request *DescribeUserDiskUsageRequest) (response *DescribeUserDiskUsageResponse, err error) {
	if request == nil {
		request = NewDescribeUserDiskUsageRequest()
	}
	response = NewDescribeUserDiskUsageResponse()
	err = c.Send(request, response)
	return
}

func NewManageRecycleRulesRequest() (request *ManageRecycleRulesRequest) {
	request = &ManageRecycleRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ManageRecycleRules")
	return
}

func NewManageRecycleRulesResponse() (response *ManageRecycleRulesResponse) {
	response = &ManageRecycleRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于增，删，改存储池或云盘的回收规则
func (c *Client) ManageRecycleRules(request *ManageRecycleRulesRequest) (response *ManageRecycleRulesResponse, err error) {
	if request == nil {
		request = NewManageRecycleRulesRequest()
	}
	response = NewManageRecycleRulesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAutoSnapshotPolicyRequest() (request *CreateAutoSnapshotPolicyRequest) {
	request = &CreateAutoSnapshotPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "CreateAutoSnapshotPolicy")
	return
}

func NewCreateAutoSnapshotPolicyResponse() (response *CreateAutoSnapshotPolicyResponse) {
	response = &CreateAutoSnapshotPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateAutoSnapshotPolicy）用于创建定期快照策略。
//
// * 每个地域最多创建10个定期快照策略。
// * 每个地域可创建的快照有数量和容量的限制，具体请见腾讯云控制台快照页面提示。
func (c *Client) CreateAutoSnapshotPolicy(request *CreateAutoSnapshotPolicyRequest) (response *CreateAutoSnapshotPolicyResponse, err error) {
	if request == nil {
		request = NewCreateAutoSnapshotPolicyRequest()
	}
	response = NewCreateAutoSnapshotPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterUpgradeTasksRequest() (request *DescribeClusterUpgradeTasksRequest) {
	request = &DescribeClusterUpgradeTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeClusterUpgradeTasks")
	return
}

func NewDescribeClusterUpgradeTasksResponse() (response *DescribeClusterUpgradeTasksResponse) {
	response = &DescribeClusterUpgradeTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeClusterUpgradeTasks）用于查询一次集群升级的所有组件升级任务
func (c *Client) DescribeClusterUpgradeTasks(request *DescribeClusterUpgradeTasksRequest) (response *DescribeClusterUpgradeTasksResponse, err error) {
	if request == nil {
		request = NewDescribeClusterUpgradeTasksRequest()
	}
	response = NewDescribeClusterUpgradeTasksResponse()
	err = c.Send(request, response)
	return
}

func NewRecycleDisksRequest() (request *RecycleDisksRequest) {
	request = &RecycleDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "RecycleDisks")
	return
}

func NewRecycleDisksResponse() (response *RecycleDisksResponse) {
	response = &RecycleDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// * 本接口用于把已删除未回收的云盘主动回收掉，释放盘在存储池占用的空间；
// * 调本接口把盘回收后，云硬盘数据将彻底释放，无法找回，调用前请确认清楚；
// * 回收是异步进行的，本接口调用成功后，表示添加回收任务成功，盘是否回收成功要根据盘的状态来判断
func (c *Client) RecycleDisks(request *RecycleDisksRequest) (response *RecycleDisksResponse, err error) {
	if request == nil {
		request = NewRecycleDisksRequest()
	}
	response = NewRecycleDisksResponse()
	err = c.Send(request, response)
	return
}

func NewReviveDiskStorageDepotCellRequest() (request *ReviveDiskStorageDepotCellRequest) {
	request = &ReviveDiskStorageDepotCellRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ReviveDiskStorageDepotCell")
	return
}

func NewReviveDiskStorageDepotCellResponse() (response *ReviveDiskStorageDepotCellResponse) {
	response = &ReviveDiskStorageDepotCellResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ReviveDiskStorageDepotCell）用于恢复cell节点或磁盘。
func (c *Client) ReviveDiskStorageDepotCell(request *ReviveDiskStorageDepotCellRequest) (response *ReviveDiskStorageDepotCellResponse, err error) {
	if request == nil {
		request = NewReviveDiskStorageDepotCellRequest()
	}
	response = NewReviveDiskStorageDepotCellResponse()
	err = c.Send(request, response)
	return
}

func NewUpgradeCommonComponentsRequest() (request *UpgradeCommonComponentsRequest) {
	request = &UpgradeCommonComponentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "UpgradeCommonComponents")
	return
}

func NewUpgradeCommonComponentsResponse() (response *UpgradeCommonComponentsResponse) {
	response = &UpgradeCommonComponentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于升级指定common组件到最新版本。
// * 如果传入的common组件未安装，则会自动新安装当前环境最新版本；
// * 如果common组件已完成，则会升级为当前环境的最新版本。
func (c *Client) UpgradeCommonComponents(request *UpgradeCommonComponentsRequest) (response *UpgradeCommonComponentsResponse, err error) {
	if request == nil {
		request = NewUpgradeCommonComponentsRequest()
	}
	response = NewUpgradeCommonComponentsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDepotCellDeviceConfigsRequest() (request *DescribeDepotCellDeviceConfigsRequest) {
	request = &DescribeDepotCellDeviceConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDepotCellDeviceConfigs")
	return
}

func NewDescribeDepotCellDeviceConfigsResponse() (response *DescribeDepotCellDeviceConfigsResponse) {
	response = &DescribeDepotCellDeviceConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询存储池存储机型的详细配置，包括主存盘数量及大小，缓存盘数量及大小
func (c *Client) DescribeDepotCellDeviceConfigs(request *DescribeDepotCellDeviceConfigsRequest) (response *DescribeDepotCellDeviceConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeDepotCellDeviceConfigsRequest()
	}
	response = NewDescribeDepotCellDeviceConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewUnbindUserFromDiskStoragePoolGroupRequest() (request *UnbindUserFromDiskStoragePoolGroupRequest) {
	request = &UnbindUserFromDiskStoragePoolGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "UnbindUserFromDiskStoragePoolGroup")
	return
}

func NewUnbindUserFromDiskStoragePoolGroupResponse() (response *UnbindUserFromDiskStoragePoolGroupResponse) {
	response = &UnbindUserFromDiskStoragePoolGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于将用户从特定的资源池解绑
func (c *Client) UnbindUserFromDiskStoragePoolGroup(request *UnbindUserFromDiskStoragePoolGroupRequest) (response *UnbindUserFromDiskStoragePoolGroupResponse, err error) {
	if request == nil {
		request = NewUnbindUserFromDiskStoragePoolGroupRequest()
	}
	response = NewUnbindUserFromDiskStoragePoolGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDepotRepotEventsRequest() (request *DescribeDepotRepotEventsRequest) {
	request = &DescribeDepotRepotEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDepotRepotEvents")
	return
}

func NewDescribeDepotRepotEventsResponse() (response *DescribeDepotRepotEventsResponse) {
	response = &DescribeDepotRepotEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDepotRepotEvents）用查询存储池事件。
func (c *Client) DescribeDepotRepotEvents(request *DescribeDepotRepotEventsRequest) (response *DescribeDepotRepotEventsResponse, err error) {
	if request == nil {
		request = NewDescribeDepotRepotEventsRequest()
	}
	response = NewDescribeDepotRepotEventsResponse()
	err = c.Send(request, response)
	return
}

func NewCancelMigrateDiskTasksRequest() (request *CancelMigrateDiskTasksRequest) {
	request = &CancelMigrateDiskTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "CancelMigrateDiskTasks")
	return
}

func NewCancelMigrateDiskTasksResponse() (response *CancelMigrateDiskTasksResponse) {
	response = &CancelMigrateDiskTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CancelMigrateDiskTasks）用于取消云盘迁移任务。
//
// *  支持批量操作；
// * 只有双写迁移的且未完成迁移的任务才支持取消；
// * 成功取消迁移任务后，数据全部在原盘上，不会丢失数据；
func (c *Client) CancelMigrateDiskTasks(request *CancelMigrateDiskTasksRequest) (response *CancelMigrateDiskTasksResponse, err error) {
	if request == nil {
		request = NewCancelMigrateDiskTasksRequest()
	}
	response = NewCancelMigrateDiskTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskIOStatRequest() (request *DescribeDiskIOStatRequest) {
	request = &DescribeDiskIOStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskIOStat")
	return
}

func NewDescribeDiskIOStatResponse() (response *DescribeDiskIOStatResponse) {
	response = &DescribeDiskIOStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询云盘秒级和分钟级的监控数据
func (c *Client) DescribeDiskIOStat(request *DescribeDiskIOStatRequest) (response *DescribeDiskIOStatResponse, err error) {
	if request == nil {
		request = NewDescribeDiskIOStatRequest()
	}
	response = NewDescribeDiskIOStatResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDepotMasterConfigRequest() (request *ModifyDepotMasterConfigRequest) {
	request = &ModifyDepotMasterConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifyDepotMasterConfig")
	return
}

func NewModifyDepotMasterConfigResponse() (response *ModifyDepotMasterConfigResponse) {
	response = &ModifyDepotMasterConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyDepotMasterConfig）用于设置存储池master的配置项
func (c *Client) ModifyDepotMasterConfig(request *ModifyDepotMasterConfigRequest) (response *ModifyDepotMasterConfigResponse, err error) {
	if request == nil {
		request = NewModifyDepotMasterConfigRequest()
	}
	response = NewModifyDepotMasterConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDisksRequest() (request *CreateDisksRequest) {
	request = &CreateDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "CreateDisks")
	return
}

func NewCreateDisksResponse() (response *CreateDisksResponse) {
	response = &CreateDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateDisks）用于创建云硬盘。
//
// * 预付费云盘的购买会预先扣除本次云盘购买所需金额，在调用本接口前请确保账户余额充足。
// * 本接口支持传入数据盘快照来创建云盘，实现将快照数据复制到新购云盘上。
// * 本接口为异步接口，当创建请求下发成功后会返回一个新建的云盘ID列表，此时云盘的创建并未立即完成。可以通过调用[DescribeDisks](/document/product/362/16315)接口根据DiskId查询对应云盘，如果能查到云盘，且状态为'UNATTACHED'或'ATTACHED'，则表示创建成功。
func (c *Client) CreateDisks(request *CreateDisksRequest) (response *CreateDisksResponse, err error) {
	if request == nil {
		request = NewCreateDisksRequest()
	}
	response = NewCreateDisksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMonitorEsIndexStorageTimeRequest() (request *DescribeMonitorEsIndexStorageTimeRequest) {
	request = &DescribeMonitorEsIndexStorageTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeMonitorEsIndexStorageTime")
	return
}

func NewDescribeMonitorEsIndexStorageTimeResponse() (response *DescribeMonitorEsIndexStorageTimeResponse) {
	response = &DescribeMonitorEsIndexStorageTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询云盘和存储池监控数据在ES的保留时间
func (c *Client) DescribeMonitorEsIndexStorageTime(request *DescribeMonitorEsIndexStorageTimeRequest) (response *DescribeMonitorEsIndexStorageTimeResponse, err error) {
	if request == nil {
		request = NewDescribeMonitorEsIndexStorageTimeRequest()
	}
	response = NewDescribeMonitorEsIndexStorageTimeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDepotTransferTaskOverviewRequest() (request *DescribeDepotTransferTaskOverviewRequest) {
	request = &DescribeDepotTransferTaskOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDepotTransferTaskOverview")
	return
}

func NewDescribeDepotTransferTaskOverviewResponse() (response *DescribeDepotTransferTaskOverviewResponse) {
	response = &DescribeDepotTransferTaskOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询存储池小表迁移的概况。
func (c *Client) DescribeDepotTransferTaskOverview(request *DescribeDepotTransferTaskOverviewRequest) (response *DescribeDepotTransferTaskOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeDepotTransferTaskOverviewRequest()
	}
	response = NewDescribeDepotTransferTaskOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDisksMonitorDataRequest() (request *DescribeDisksMonitorDataRequest) {
	request = &DescribeDisksMonitorDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDisksMonitorData")
	return
}

func NewDescribeDisksMonitorDataResponse() (response *DescribeDisksMonitorDataResponse) {
	response = &DescribeDisksMonitorDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDisksMonitorData）用于查询云硬盘、存储池等监控数据。为了便于用户使用，该接口入参出参参考[云监控接口-拉取指标监控数据](https://cloud.tencent.com/document/product/248/31014)设计。
func (c *Client) DescribeDisksMonitorData(request *DescribeDisksMonitorDataRequest) (response *DescribeDisksMonitorDataResponse, err error) {
	if request == nil {
		request = NewDescribeDisksMonitorDataRequest()
	}
	response = NewDescribeDisksMonitorDataResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDepotUpgradeTasksRequest() (request *CreateDepotUpgradeTasksRequest) {
	request = &CreateDepotUpgradeTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "CreateDepotUpgradeTasks")
	return
}

func NewCreateDepotUpgradeTasksResponse() (response *CreateDepotUpgradeTasksResponse) {
	response = &CreateDepotUpgradeTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateDepotUpgradeTasks）用于生成云硬盘存储池升级任务。
// * 选定存储池的目标版本后，会自动生成该存储池所有组件的升级任务；
func (c *Client) CreateDepotUpgradeTasks(request *CreateDepotUpgradeTasksRequest) (response *CreateDepotUpgradeTasksResponse, err error) {
	if request == nil {
		request = NewCreateDepotUpgradeTasksRequest()
	}
	response = NewCreateDepotUpgradeTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSnapshotBoxPerformanceDataRequest() (request *DescribeSnapshotBoxPerformanceDataRequest) {
	request = &DescribeSnapshotBoxPerformanceDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshotBoxPerformanceData")
	return
}

func NewDescribeSnapshotBoxPerformanceDataResponse() (response *DescribeSnapshotBoxPerformanceDataResponse) {
	response = &DescribeSnapshotBoxPerformanceDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询快照组件的性能监控数据
func (c *Client) DescribeSnapshotBoxPerformanceData(request *DescribeSnapshotBoxPerformanceDataRequest) (response *DescribeSnapshotBoxPerformanceDataResponse, err error) {
	if request == nil {
		request = NewDescribeSnapshotBoxPerformanceDataRequest()
	}
	response = NewDescribeSnapshotBoxPerformanceDataResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSnapshotBoxConfigRequest() (request *UpdateSnapshotBoxConfigRequest) {
	request = &UpdateSnapshotBoxConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "UpdateSnapshotBoxConfig")
	return
}

func NewUpdateSnapshotBoxConfigResponse() (response *UpdateSnapshotBoxConfigResponse) {
	response = &UpdateSnapshotBoxConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改快照集群配置。
// * 由于相同组件是多机部署的，执行更新前，接口会先统一检查，必须所有机器上都能更新成功，才能执行更新；只要有某台机器无法更新，所有机器上均不会执行更新。
func (c *Client) UpdateSnapshotBoxConfig(request *UpdateSnapshotBoxConfigRequest) (response *UpdateSnapshotBoxConfigResponse, err error) {
	if request == nil {
		request = NewUpdateSnapshotBoxConfigRequest()
	}
	response = NewUpdateSnapshotBoxConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskOperationsRequest() (request *DescribeDiskOperationsRequest) {
	request = &DescribeDiskOperationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskOperations")
	return
}

func NewDescribeDiskOperationsResponse() (response *DescribeDiskOperationsResponse) {
	response = &DescribeDiskOperationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskOperations）用于查询云盘、快照、定期快照策略等资源的操作流水。（当前仅支持查询定期快照操作流水）
//
// * 可以根据云硬盘ID、快照ID、定期快照策略ID、子机ID等信息来查询操作流水，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的定期快照策略表
func (c *Client) DescribeDiskOperations(request *DescribeDiskOperationsRequest) (response *DescribeDiskOperationsResponse, err error) {
	if request == nil {
		request = NewDescribeDiskOperationsRequest()
	}
	response = NewDescribeDiskOperationsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSnapshotRequest() (request *CreateSnapshotRequest) {
	request = &CreateSnapshotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "CreateSnapshot")
	return
}

func NewCreateSnapshotResponse() (response *CreateSnapshotResponse) {
	response = &CreateSnapshotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateSnapshot）用于对指定云盘创建快照。
//
// * 只有具有快照能力的云硬盘才能创建快照。云硬盘是否具有快照能力可由[DescribeDisks](/document/product/362/16315)接口查询，见SnapshotAbility字段。
// * 可创建快照数量限制见[产品使用限制](https://cloud.tencent.com/doc/product/362/5145)。
func (c *Client) CreateSnapshot(request *CreateSnapshotRequest) (response *CreateSnapshotResponse, err error) {
	if request == nil {
		request = NewCreateSnapshotRequest()
	}
	response = NewCreateSnapshotResponse()
	err = c.Send(request, response)
	return
}

func NewTerminateDisksRequest() (request *TerminateDisksRequest) {
	request = &TerminateDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "TerminateDisks")
	return
}

func NewTerminateDisksResponse() (response *TerminateDisksResponse) {
	response = &TerminateDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（TerminateDisks）用于退还云硬盘。
//
// * 当前仅支持退还包年包月云盘。
// * 支持批量操作，每次请求批量云硬盘的上限为50。如果批量云盘存在不允许操作的，请求会以特定错误码返回。
func (c *Client) TerminateDisks(request *TerminateDisksRequest) (response *TerminateDisksResponse, err error) {
	if request == nil {
		request = NewTerminateDisksRequest()
	}
	response = NewTerminateDisksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskStoragePoolGroupBoundRequest() (request *DescribeDiskStoragePoolGroupBoundRequest) {
	request = &DescribeDiskStoragePoolGroupBoundRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskStoragePoolGroupBound")
	return
}

func NewDescribeDiskStoragePoolGroupBoundResponse() (response *DescribeDiskStoragePoolGroupBoundResponse) {
	response = &DescribeDiskStoragePoolGroupBoundResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于查询云硬盘资源池组绑定情况
func (c *Client) DescribeDiskStoragePoolGroupBound(request *DescribeDiskStoragePoolGroupBoundRequest) (response *DescribeDiskStoragePoolGroupBoundResponse, err error) {
	if request == nil {
		request = NewDescribeDiskStoragePoolGroupBoundRequest()
	}
	response = NewDescribeDiskStoragePoolGroupBoundResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDisksMonitorAlarmRulesRequest() (request *ModifyDisksMonitorAlarmRulesRequest) {
	request = &ModifyDisksMonitorAlarmRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ModifyDisksMonitorAlarmRules")
	return
}

func NewModifyDisksMonitorAlarmRulesResponse() (response *ModifyDisksMonitorAlarmRulesResponse) {
	response = &ModifyDisksMonitorAlarmRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改判断云盘IO告警的默认告警策略，包括云盘iohang、svctm_high和存储池svctm_high告警
func (c *Client) ModifyDisksMonitorAlarmRules(request *ModifyDisksMonitorAlarmRulesRequest) (response *ModifyDisksMonitorAlarmRulesResponse, err error) {
	if request == nil {
		request = NewModifyDisksMonitorAlarmRulesRequest()
	}
	response = NewModifyDisksMonitorAlarmRulesResponse()
	err = c.Send(request, response)
	return
}

func NewUninstallBoxDeviceRequest() (request *UninstallBoxDeviceRequest) {
	request = &UninstallBoxDeviceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "UninstallBoxDevice")
	return
}

func NewUninstallBoxDeviceResponse() (response *UninstallBoxDeviceResponse) {
	response = &UninstallBoxDeviceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UninstallBoxDevice）卸载快照box设备， 包括manager, scheduler, transfer。
func (c *Client) UninstallBoxDevice(request *UninstallBoxDeviceRequest) (response *UninstallBoxDeviceResponse, err error) {
	if request == nil {
		request = NewUninstallBoxDeviceRequest()
	}
	response = NewUninstallBoxDeviceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDepotDiskDetailsRequest() (request *DescribeDepotDiskDetailsRequest) {
	request = &DescribeDepotDiskDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDepotDiskDetails")
	return
}

func NewDescribeDepotDiskDetailsResponse() (response *DescribeDepotDiskDetailsResponse) {
	response = &DescribeDepotDiskDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDepotDiskDetails）用于云硬盘存储仓库资源详情。
func (c *Client) DescribeDepotDiskDetails(request *DescribeDepotDiskDetailsRequest) (response *DescribeDepotDiskDetailsResponse, err error) {
	if request == nil {
		request = NewDescribeDepotDiskDetailsRequest()
	}
	response = NewDescribeDepotDiskDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceModifyDiskAttributesRequest() (request *InquiryPriceModifyDiskAttributesRequest) {
	request = &InquiryPriceModifyDiskAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "InquiryPriceModifyDiskAttributes")
	return
}

func NewInquiryPriceModifyDiskAttributesResponse() (response *InquiryPriceModifyDiskAttributesResponse) {
	response = &InquiryPriceModifyDiskAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InquiryPriceModifyDiskAttributes）用于修改云盘类型询价。
//
// * 当前仅支持弹性云盘修改类型（[DescribeDisks](/document/product/362/16315)接口的返回字段Portable为true表示弹性云盘）。
// * 当前仅支持云盘类型升级，不支持降级，具体如下:
//   - CLOUD_BASIC变更为CLOUD_PREMIUM；
//   - CLOUD_BASIC变更为CLOUD_SSD；
//   - CLOUD_PREMIUM变更为CLOUD_SSD。
func (c *Client) InquiryPriceModifyDiskAttributes(request *InquiryPriceModifyDiskAttributesRequest) (response *InquiryPriceModifyDiskAttributesResponse, err error) {
	if request == nil {
		request = NewInquiryPriceModifyDiskAttributesRequest()
	}
	response = NewInquiryPriceModifyDiskAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskStorageDepotsRequest() (request *DescribeDiskStorageDepotsRequest) {
	request = &DescribeDiskStorageDepotsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskStorageDepots")
	return
}

func NewDescribeDiskStorageDepotsResponse() (response *DescribeDiskStorageDepotsResponse) {
	response = &DescribeDiskStorageDepotsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDiskStorageDepots）用于查询云硬盘存储仓库信息。
func (c *Client) DescribeDiskStorageDepots(request *DescribeDiskStorageDepotsRequest) (response *DescribeDiskStorageDepotsResponse, err error) {
	if request == nil {
		request = NewDescribeDiskStorageDepotsRequest()
	}
	response = NewDescribeDiskStorageDepotsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRecycleRulesRequest() (request *DescribeRecycleRulesRequest) {
	request = &DescribeRecycleRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeRecycleRules")
	return
}

func NewDescribeRecycleRulesResponse() (response *DescribeRecycleRulesResponse) {
	response = &DescribeRecycleRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询云盘或存储池的回收规则
func (c *Client) DescribeRecycleRules(request *DescribeRecycleRulesRequest) (response *DescribeRecycleRulesResponse, err error) {
	if request == nil {
		request = NewDescribeRecycleRulesRequest()
	}
	response = NewDescribeRecycleRulesResponse()
	err = c.Send(request, response)
	return
}

func NewInstallDiskStorageDepotMasterRequest() (request *InstallDiskStorageDepotMasterRequest) {
	request = &InstallDiskStorageDepotMasterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "InstallDiskStorageDepotMaster")
	return
}

func NewInstallDiskStorageDepotMasterResponse() (response *InstallDiskStorageDepotMasterResponse) {
	response = &InstallDiskStorageDepotMasterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（InstallDiskStorageDepotMaster）用于安装云硬盘存储集群Master组件。
func (c *Client) InstallDiskStorageDepotMaster(request *InstallDiskStorageDepotMasterRequest) (response *InstallDiskStorageDepotMasterResponse, err error) {
	if request == nil {
		request = NewInstallDiskStorageDepotMasterRequest()
	}
	response = NewInstallDiskStorageDepotMasterResponse()
	err = c.Send(request, response)
	return
}

func NewDetachDisksRequest() (request *DetachDisksRequest) {
	request = &DetachDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DetachDisks")
	return
}

func NewDetachDisksResponse() (response *DetachDisksResponse) {
	response = &DetachDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DetachDisks）用于解挂云硬盘。
//
// * 支持批量操作，解挂卸载在同一主机上的多块云盘。如果多块云盘存在不允许卸载的云盘，则操作不执行，以返回特定的错误码返回。
// * 本接口为异步接口，当请求成功返回时，云盘并未立即从主机解挂载，可通过接口[DescribeDisks](/document/product/362/16315)来查询对应云盘的状态，如果云盘的状态由“ATTACHED”变为“UNATTACHED”，则为卸载成功。
func (c *Client) DetachDisks(request *DetachDisksRequest) (response *DetachDisksResponse, err error) {
	if request == nil {
		request = NewDetachDisksRequest()
	}
	response = NewDetachDisksResponse()
	err = c.Send(request, response)
	return
}

func NewResizeDiskRequest() (request *ResizeDiskRequest) {
	request = &ResizeDiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "ResizeDisk")
	return
}

func NewResizeDiskResponse() (response *ResizeDiskResponse) {
	response = &ResizeDiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ResizeDisk）用于扩容云硬盘。
//
// * 只支持扩容弹性云盘。云硬盘类型可以通过[DescribeDisks](/document/product/362/16315)接口查询，见输出参数中Portable字段解释。随云主机创建的云硬盘需通过[ResizeInstanceDisks](/document/product/213/15731)接口扩容。
// * 本接口为异步接口，接口成功返回时，云盘并未立即扩容到指定大小，可通过接口[DescribeDisks](/document/product/362/16315)来查询对应云盘的状态，如果云盘的状态为“EXPANDING”，表示正在扩容中，当状态变为“UNATTACHED”，表示扩容完成。
func (c *Client) ResizeDisk(request *ResizeDiskRequest) (response *ResizeDiskResponse, err error) {
	if request == nil {
		request = NewResizeDiskRequest()
	}
	response = NewResizeDiskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomerAccountsRequest() (request *DescribeCustomerAccountsRequest) {
	request = &DescribeCustomerAccountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeCustomerAccounts")
	return
}

func NewDescribeCustomerAccountsResponse() (response *DescribeCustomerAccountsResponse) {
	response = &DescribeCustomerAccountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户信息
func (c *Client) DescribeCustomerAccounts(request *DescribeCustomerAccountsRequest) (response *DescribeCustomerAccountsResponse, err error) {
	if request == nil {
		request = NewDescribeCustomerAccountsRequest()
	}
	response = NewDescribeCustomerAccountsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReplaceDiskTasksRequest() (request *DescribeReplaceDiskTasksRequest) {
	request = &DescribeReplaceDiskTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeReplaceDiskTasks")
	return
}

func NewDescribeReplaceDiskTasksResponse() (response *DescribeReplaceDiskTasksResponse) {
	response = &DescribeReplaceDiskTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// * 本接口（DescribeReplaceDiskTasks）用于查询换盘任务详情
func (c *Client) DescribeReplaceDiskTasks(request *DescribeReplaceDiskTasksRequest) (response *DescribeReplaceDiskTasksResponse, err error) {
	if request == nil {
		request = NewDescribeReplaceDiskTasksRequest()
	}
	response = NewDescribeReplaceDiskTasksResponse()
	err = c.Send(request, response)
	return
}

func NewStopDiskStorageDepotCellRequest() (request *StopDiskStorageDepotCellRequest) {
	request = &StopDiskStorageDepotCellRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "StopDiskStorageDepotCell")
	return
}

func NewStopDiskStorageDepotCellResponse() (response *StopDiskStorageDepotCellResponse) {
	response = &StopDiskStorageDepotCellResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（StopDiskStorageDepotCell）用于下架CBS存储集群CELL节点前，停止CBS存储集群CELL进程。
func (c *Client) StopDiskStorageDepotCell(request *StopDiskStorageDepotCellRequest) (response *StopDiskStorageDepotCellResponse, err error) {
	if request == nil {
		request = NewStopDiskStorageDepotCellRequest()
	}
	response = NewStopDiskStorageDepotCellResponse()
	err = c.Send(request, response)
	return
}
