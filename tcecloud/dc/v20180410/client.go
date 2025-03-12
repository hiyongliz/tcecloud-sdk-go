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

package v20180410

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2018-04-10"

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

func NewModifyAccessPointRequest() (request *ModifyAccessPointRequest) {
	request = &ModifyAccessPointRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "ModifyAccessPoint")
	return
}

func NewModifyAccessPointResponse() (response *ModifyAccessPointResponse) {
	response = &ModifyAccessPointResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改专线接入点信息
func (c *Client) ModifyAccessPoint(request *ModifyAccessPointRequest) (response *ModifyAccessPointResponse, err error) {
	if request == nil {
		request = NewModifyAccessPointRequest()
	}
	response = NewModifyAccessPointResponse()
	err = c.Send(request, response)
	return
}

func NewAcceptDirectConnectRequest() (request *AcceptDirectConnectRequest) {
	request = &AcceptDirectConnectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "AcceptDirectConnect")
	return
}

func NewAcceptDirectConnectResponse() (response *AcceptDirectConnectResponse) {
	response = &AcceptDirectConnectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 审批通过专线接入申请，调用时传入PortId，则使用指定的端口创建专线，否则自动分配交换机端口
func (c *Client) AcceptDirectConnect(request *AcceptDirectConnectRequest) (response *AcceptDirectConnectResponse, err error) {
	if request == nil {
		request = NewAcceptDirectConnectRequest()
	}
	response = NewAcceptDirectConnectResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessPointsRequest() (request *DescribeAccessPointsRequest) {
	request = &DescribeAccessPointsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "DescribeAccessPoints")
	return
}

func NewDescribeAccessPointsResponse() (response *DescribeAccessPointsResponse) {
	response = &DescribeAccessPointsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询物理专线接入点
func (c *Client) DescribeAccessPoints(request *DescribeAccessPointsRequest) (response *DescribeAccessPointsResponse, err error) {
	if request == nil {
		request = NewDescribeAccessPointsRequest()
	}
	response = NewDescribeAccessPointsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDirectConnectTunnelsRequest() (request *DescribeDirectConnectTunnelsRequest) {
	request = &DescribeDirectConnectTunnelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "DescribeDirectConnectTunnels")
	return
}

func NewDescribeDirectConnectTunnelsResponse() (response *DescribeDirectConnectTunnelsResponse) {
	response = &DescribeDirectConnectTunnelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询专用通道列表。
func (c *Client) DescribeDirectConnectTunnels(request *DescribeDirectConnectTunnelsRequest) (response *DescribeDirectConnectTunnelsResponse, err error) {
	if request == nil {
		request = NewDescribeDirectConnectTunnelsRequest()
	}
	response = NewDescribeDirectConnectTunnelsResponse()
	err = c.Send(request, response)
	return
}

func NewRejectDirectConnectRequest() (request *RejectDirectConnectRequest) {
	request = &RejectDirectConnectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "RejectDirectConnect")
	return
}

func NewRejectDirectConnectResponse() (response *RejectDirectConnectResponse) {
	response = &RejectDirectConnectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端拒绝专线申请
func (c *Client) RejectDirectConnect(request *RejectDirectConnectRequest) (response *RejectDirectConnectResponse, err error) {
	if request == nil {
		request = NewRejectDirectConnectRequest()
	}
	response = NewRejectDirectConnectResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAccessPointRequest() (request *CreateAccessPointRequest) {
	request = &CreateAccessPointRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "CreateAccessPoint")
	return
}

func NewCreateAccessPointResponse() (response *CreateAccessPointResponse) {
	response = &CreateAccessPointResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建专线接入点
func (c *Client) CreateAccessPoint(request *CreateAccessPointRequest) (response *CreateAccessPointResponse, err error) {
	if request == nil {
		request = NewCreateAccessPointRequest()
	}
	response = NewCreateAccessPointResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSwitchRequest() (request *DeleteSwitchRequest) {
	request = &DeleteSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "DeleteSwitch")
	return
}

func NewDeleteSwitchResponse() (response *DeleteSwitchResponse) {
	response = &DeleteSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除专线接入交换机，仅在交换机上没有正在运行的专线时才能删除该交换机
func (c *Client) DeleteSwitch(request *DeleteSwitchRequest) (response *DeleteSwitchResponse, err error) {
	if request == nil {
		request = NewDeleteSwitchRequest()
	}
	response = NewDeleteSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSwitchVendorsRequest() (request *DescribeSwitchVendorsRequest) {
	request = &DescribeSwitchVendorsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "DescribeSwitchVendors")
	return
}

func NewDescribeSwitchVendorsResponse() (response *DescribeSwitchVendorsResponse) {
	response = &DescribeSwitchVendorsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询交换机厂商列表
func (c *Client) DescribeSwitchVendors(request *DescribeSwitchVendorsRequest) (response *DescribeSwitchVendorsResponse, err error) {
	if request == nil {
		request = NewDescribeSwitchVendorsRequest()
	}
	response = NewDescribeSwitchVendorsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductFailureMigrateRequest() (request *CreateProductFailureMigrateRequest) {
	request = &CreateProductFailureMigrateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "CreateProductFailureMigrate")
	return
}

func NewCreateProductFailureMigrateResponse() (response *CreateProductFailureMigrateResponse) {
	response = &CreateProductFailureMigrateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通知可用区故障，返回故障迁移任务ID
func (c *Client) CreateProductFailureMigrate(request *CreateProductFailureMigrateRequest) (response *CreateProductFailureMigrateResponse, err error) {
	if request == nil {
		request = NewCreateProductFailureMigrateRequest()
	}
	response = NewCreateProductFailureMigrateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSwitchesRequest() (request *DescribeSwitchesRequest) {
	request = &DescribeSwitchesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "DescribeSwitches")
	return
}

func NewDescribeSwitchesResponse() (response *DescribeSwitchesResponse) {
	response = &DescribeSwitchesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询专线接入交换机
func (c *Client) DescribeSwitches(request *DescribeSwitchesRequest) (response *DescribeSwitchesResponse, err error) {
	if request == nil {
		request = NewDescribeSwitchesRequest()
	}
	response = NewDescribeSwitchesResponse()
	err = c.Send(request, response)
	return
}

func NewImplementDirectConnectRequest() (request *ImplementDirectConnectRequest) {
	request = &ImplementDirectConnectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "ImplementDirectConnect")
	return
}

func NewImplementDirectConnectResponse() (response *ImplementDirectConnectResponse) {
	response = &ImplementDirectConnectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 专线人工确认实施完成，填写线路编号，保障联系人，保障联系电话
func (c *Client) ImplementDirectConnect(request *ImplementDirectConnectRequest) (response *ImplementDirectConnectResponse, err error) {
	if request == nil {
		request = NewImplementDirectConnectRequest()
	}
	response = NewImplementDirectConnectResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDirectConnectsRequest() (request *DescribeDirectConnectsRequest) {
	request = &DescribeDirectConnectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "DescribeDirectConnects")
	return
}

func NewDescribeDirectConnectsResponse() (response *DescribeDirectConnectsResponse) {
	response = &DescribeDirectConnectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询物理专线列表。
func (c *Client) DescribeDirectConnects(request *DescribeDirectConnectsRequest) (response *DescribeDirectConnectsResponse, err error) {
	if request == nil {
		request = NewDescribeDirectConnectsRequest()
	}
	response = NewDescribeDirectConnectsResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductStateInfoRequest() (request *QueryProductStateInfoRequest) {
	request = &QueryProductStateInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "QueryProductStateInfo")
	return
}

func NewQueryProductStateInfoResponse() (response *QueryProductStateInfoResponse) {
	response = &QueryProductStateInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品集群列表和集群内节点列表
func (c *Client) QueryProductStateInfo(request *QueryProductStateInfoRequest) (response *QueryProductStateInfoResponse, err error) {
	if request == nil {
		request = NewQueryProductStateInfoRequest()
	}
	response = NewQueryProductStateInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDirectConnectAttributeRequest() (request *ModifyDirectConnectAttributeRequest) {
	request = &ModifyDirectConnectAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "ModifyDirectConnectAttribute")
	return
}

func NewModifyDirectConnectAttributeResponse() (response *ModifyDirectConnectAttributeResponse) {
	response = &ModifyDirectConnectAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改物理专线的属性。
// 如果调用此接口时传入了TrunkName，MemberPorts和SwitchId需要同时提供。
func (c *Client) ModifyDirectConnectAttribute(request *ModifyDirectConnectAttributeRequest) (response *ModifyDirectConnectAttributeResponse, err error) {
	if request == nil {
		request = NewModifyDirectConnectAttributeRequest()
	}
	response = NewModifyDirectConnectAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSwitchRequest() (request *CreateSwitchRequest) {
	request = &CreateSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "CreateSwitch")
	return
}

func NewCreateSwitchResponse() (response *CreateSwitchResponse) {
	response = &CreateSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建专线接入交换机
func (c *Client) CreateSwitch(request *CreateSwitchRequest) (response *CreateSwitchResponse, err error) {
	if request == nil {
		request = NewCreateSwitchRequest()
	}
	response = NewCreateSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductHealthStateRequest() (request *QueryProductHealthStateRequest) {
	request = &QueryProductHealthStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "QueryProductHealthState")
	return
}

func NewQueryProductHealthStateResponse() (response *QueryProductHealthStateResponse) {
	response = &QueryProductHealthStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回集群健康状态，以及集群健康指标数组
func (c *Client) QueryProductHealthState(request *QueryProductHealthStateRequest) (response *QueryProductHealthStateResponse, err error) {
	if request == nil {
		request = NewQueryProductHealthStateRequest()
	}
	response = NewQueryProductHealthStateResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductFailureMigrateTaskDetailRequest() (request *QueryProductFailureMigrateTaskDetailRequest) {
	request = &QueryProductFailureMigrateTaskDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "QueryProductFailureMigrateTaskDetail")
	return
}

func NewQueryProductFailureMigrateTaskDetailResponse() (response *QueryProductFailureMigrateTaskDetailResponse) {
	response = &QueryProductFailureMigrateTaskDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据任务ID查询故障迁移结果
func (c *Client) QueryProductFailureMigrateTaskDetail(request *QueryProductFailureMigrateTaskDetailRequest) (response *QueryProductFailureMigrateTaskDetailResponse, err error) {
	if request == nil {
		request = NewQueryProductFailureMigrateTaskDetailRequest()
	}
	response = NewQueryProductFailureMigrateTaskDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSwitchPortsRequest() (request *DescribeSwitchPortsRequest) {
	request = &DescribeSwitchPortsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "DescribeSwitchPorts")
	return
}

func NewDescribeSwitchPortsResponse() (response *DescribeSwitchPortsResponse) {
	response = &DescribeSwitchPortsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询交换机可用端口
func (c *Client) DescribeSwitchPorts(request *DescribeSwitchPortsRequest) (response *DescribeSwitchPortsResponse, err error) {
	if request == nil {
		request = NewDescribeSwitchPortsRequest()
	}
	response = NewDescribeSwitchPortsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableSwitchesRequest() (request *DescribeAvailableSwitchesRequest) {
	request = &DescribeAvailableSwitchesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "DescribeAvailableSwitches")
	return
}

func NewDescribeAvailableSwitchesResponse() (response *DescribeAvailableSwitchesResponse) {
	response = &DescribeAvailableSwitchesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从DCOS获取专线接入交换机
func (c *Client) DescribeAvailableSwitches(request *DescribeAvailableSwitchesRequest) (response *DescribeAvailableSwitchesResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableSwitchesRequest()
	}
	response = NewDescribeAvailableSwitchesResponse()
	err = c.Send(request, response)
	return
}

func NewModifySwitchRequest() (request *ModifySwitchRequest) {
	request = &ModifySwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dc", APIVersion, "ModifySwitch")
	return
}

func NewModifySwitchResponse() (response *ModifySwitchResponse) {
	response = &ModifySwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改专线接入交换机信息
func (c *Client) ModifySwitch(request *ModifySwitchRequest) (response *ModifySwitchResponse, err error) {
	if request == nil {
		request = NewModifySwitchRequest()
	}
	response = NewModifySwitchResponse()
	err = c.Send(request, response)
	return
}
