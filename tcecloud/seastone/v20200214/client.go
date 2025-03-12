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

package v20200214

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-02-14"

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

func NewDescribeControllerSetsRequest() (request *DescribeControllerSetsRequest) {
	request = &DescribeControllerSetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "DescribeControllerSets")
	return
}

func NewDescribeControllerSetsResponse() (response *DescribeControllerSetsResponse) {
	response = &DescribeControllerSetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取控制器集群列表
func (c *Client) DescribeControllerSets(request *DescribeControllerSetsRequest) (response *DescribeControllerSetsResponse, err error) {
	if request == nil {
		request = NewDescribeControllerSetsRequest()
	}
	response = NewDescribeControllerSetsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateWhiteListRequest() (request *CreateWhiteListRequest) {
	request = &CreateWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "CreateWhiteList")
	return
}

func NewCreateWhiteListResponse() (response *CreateWhiteListResponse) {
	response = &CreateWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于创建一个白名单信息，用于配置控制器集群适用的范围
func (c *Client) CreateWhiteList(request *CreateWhiteListRequest) (response *CreateWhiteListResponse, err error) {
	if request == nil {
		request = NewCreateWhiteListRequest()
	}
	response = NewCreateWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteControllerSetsRequest() (request *DeleteControllerSetsRequest) {
	request = &DeleteControllerSetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "DeleteControllerSets")
	return
}

func NewDeleteControllerSetsResponse() (response *DeleteControllerSetsResponse) {
	response = &DeleteControllerSetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除控制器集群
func (c *Client) DeleteControllerSets(request *DeleteControllerSetsRequest) (response *DeleteControllerSetsResponse, err error) {
	if request == nil {
		request = NewDeleteControllerSetsRequest()
	}
	response = NewDeleteControllerSetsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateLimitRuleRequest() (request *UpdateLimitRuleRequest) {
	request = &UpdateLimitRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "UpdateLimitRule")
	return
}

func NewUpdateLimitRuleResponse() (response *UpdateLimitRuleResponse) {
	response = &UpdateLimitRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于更新一条限速规则带宽等信息
func (c *Client) UpdateLimitRule(request *UpdateLimitRuleRequest) (response *UpdateLimitRuleResponse, err error) {
	if request == nil {
		request = NewUpdateLimitRuleRequest()
	}
	response = NewUpdateLimitRuleResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGatewaySetsRequest() (request *UpdateGatewaySetsRequest) {
	request = &UpdateGatewaySetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "UpdateGatewaySets")
	return
}

func NewUpdateGatewaySetsResponse() (response *UpdateGatewaySetsResponse) {
	response = &UpdateGatewaySetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于批量更新gateway集群属性
func (c *Client) UpdateGatewaySets(request *UpdateGatewaySetsRequest) (response *UpdateGatewaySetsResponse, err error) {
	if request == nil {
		request = NewUpdateGatewaySetsRequest()
	}
	response = NewUpdateGatewaySetsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLimitRuleRequest() (request *CreateLimitRuleRequest) {
	request = &CreateLimitRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "CreateLimitRule")
	return
}

func NewCreateLimitRuleResponse() (response *CreateLimitRuleResponse) {
	response = &CreateLimitRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于新增一条限速规则
func (c *Client) CreateLimitRule(request *CreateLimitRuleRequest) (response *CreateLimitRuleResponse, err error) {
	if request == nil {
		request = NewCreateLimitRuleRequest()
	}
	response = NewCreateLimitRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLimitRulesRequest() (request *DescribeLimitRulesRequest) {
	request = &DescribeLimitRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "DescribeLimitRules")
	return
}

func NewDescribeLimitRulesResponse() (response *DescribeLimitRulesResponse) {
	response = &DescribeLimitRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取限速规则列表
func (c *Client) DescribeLimitRules(request *DescribeLimitRulesRequest) (response *DescribeLimitRulesResponse, err error) {
	if request == nil {
		request = NewDescribeLimitRulesRequest()
	}
	response = NewDescribeLimitRulesResponse()
	err = c.Send(request, response)
	return
}

func NewZkRecoveryRequest() (request *ZkRecoveryRequest) {
	request = &ZkRecoveryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "ZkRecovery")
	return
}

func NewZkRecoveryResponse() (response *ZkRecoveryResponse) {
	response = &ZkRecoveryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ZK对账及恢复能力
func (c *Client) ZkRecovery(request *ZkRecoveryRequest) (response *ZkRecoveryResponse, err error) {
	if request == nil {
		request = NewZkRecoveryRequest()
	}
	response = NewZkRecoveryResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGatewaySetRequest() (request *CreateGatewaySetRequest) {
	request = &CreateGatewaySetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "CreateGatewaySet")
	return
}

func NewCreateGatewaySetResponse() (response *CreateGatewaySetResponse) {
	response = &CreateGatewaySetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于创建一个gateway集群
func (c *Client) CreateGatewaySet(request *CreateGatewaySetRequest) (response *CreateGatewaySetResponse, err error) {
	if request == nil {
		request = NewCreateGatewaySetRequest()
	}
	response = NewCreateGatewaySetResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDevicesRequest() (request *DeleteDevicesRequest) {
	request = &DeleteDevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "DeleteDevices")
	return
}

func NewDeleteDevicesResponse() (response *DeleteDevicesResponse) {
	response = &DeleteDevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除网关节点
func (c *Client) DeleteDevices(request *DeleteDevicesRequest) (response *DeleteDevicesResponse, err error) {
	if request == nil {
		request = NewDeleteDevicesRequest()
	}
	response = NewDeleteDevicesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlgorithmsRequest() (request *DescribeAlgorithmsRequest) {
	request = &DescribeAlgorithmsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "DescribeAlgorithms")
	return
}

func NewDescribeAlgorithmsResponse() (response *DescribeAlgorithmsResponse) {
	response = &DescribeAlgorithmsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取支持算法列表
func (c *Client) DescribeAlgorithms(request *DescribeAlgorithmsRequest) (response *DescribeAlgorithmsResponse, err error) {
	if request == nil {
		request = NewDescribeAlgorithmsRequest()
	}
	response = NewDescribeAlgorithmsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeControllersRequest() (request *DescribeControllersRequest) {
	request = &DescribeControllersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "DescribeControllers")
	return
}

func NewDescribeControllersResponse() (response *DescribeControllersResponse) {
	response = &DescribeControllersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取控制器节点列表
func (c *Client) DescribeControllers(request *DescribeControllersRequest) (response *DescribeControllersResponse, err error) {
	if request == nil {
		request = NewDescribeControllersRequest()
	}
	response = NewDescribeControllersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDeviceRequest() (request *CreateDeviceRequest) {
	request = &CreateDeviceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "CreateDevice")
	return
}

func NewCreateDeviceResponse() (response *CreateDeviceResponse) {
	response = &CreateDeviceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于新增一台网关节点
func (c *Client) CreateDevice(request *CreateDeviceRequest) (response *CreateDeviceResponse, err error) {
	if request == nil {
		request = NewCreateDeviceRequest()
	}
	response = NewCreateDeviceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLimitRulesRequest() (request *DeleteLimitRulesRequest) {
	request = &DeleteLimitRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "DeleteLimitRules")
	return
}

func NewDeleteLimitRulesResponse() (response *DeleteLimitRulesResponse) {
	response = &DeleteLimitRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除限速规则
func (c *Client) DeleteLimitRules(request *DeleteLimitRulesRequest) (response *DeleteLimitRulesResponse, err error) {
	if request == nil {
		request = NewDeleteLimitRulesRequest()
	}
	response = NewDeleteLimitRulesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateControllerSetsRequest() (request *UpdateControllerSetsRequest) {
	request = &UpdateControllerSetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "UpdateControllerSets")
	return
}

func NewUpdateControllerSetsResponse() (response *UpdateControllerSetsResponse) {
	response = &UpdateControllerSetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于批量更新controller集群属性
func (c *Client) UpdateControllerSets(request *UpdateControllerSetsRequest) (response *UpdateControllerSetsResponse, err error) {
	if request == nil {
		request = NewUpdateControllerSetsRequest()
	}
	response = NewUpdateControllerSetsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateControllerRequest() (request *CreateControllerRequest) {
	request = &CreateControllerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "CreateController")
	return
}

func NewCreateControllerResponse() (response *CreateControllerResponse) {
	response = &CreateControllerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于新增一个controller节点
func (c *Client) CreateController(request *CreateControllerRequest) (response *CreateControllerResponse, err error) {
	if request == nil {
		request = NewCreateControllerRequest()
	}
	response = NewCreateControllerResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateDevicesRequest() (request *UpdateDevicesRequest) {
	request = &UpdateDevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "UpdateDevices")
	return
}

func NewUpdateDevicesResponse() (response *UpdateDevicesResponse) {
	response = &UpdateDevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于批量更新网关节点属性
func (c *Client) UpdateDevices(request *UpdateDevicesRequest) (response *UpdateDevicesResponse, err error) {
	if request == nil {
		request = NewUpdateDevicesRequest()
	}
	response = NewUpdateDevicesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewaySetsRequest() (request *DescribeGatewaySetsRequest) {
	request = &DescribeGatewaySetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "DescribeGatewaySets")
	return
}

func NewDescribeGatewaySetsResponse() (response *DescribeGatewaySetsResponse) {
	response = &DescribeGatewaySetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取网关集群列表
func (c *Client) DescribeGatewaySets(request *DescribeGatewaySetsRequest) (response *DescribeGatewaySetsResponse, err error) {
	if request == nil {
		request = NewDescribeGatewaySetsRequest()
	}
	response = NewDescribeGatewaySetsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateControllersRequest() (request *UpdateControllersRequest) {
	request = &UpdateControllersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "UpdateControllers")
	return
}

func NewUpdateControllersResponse() (response *UpdateControllersResponse) {
	response = &UpdateControllersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于批量更新controller属性
func (c *Client) UpdateControllers(request *UpdateControllersRequest) (response *UpdateControllersResponse, err error) {
	if request == nil {
		request = NewUpdateControllersRequest()
	}
	response = NewUpdateControllersResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGatewaySetsRequest() (request *DeleteGatewaySetsRequest) {
	request = &DeleteGatewaySetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "DeleteGatewaySets")
	return
}

func NewDeleteGatewaySetsResponse() (response *DeleteGatewaySetsResponse) {
	response = &DeleteGatewaySetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于批量删除网关集群
func (c *Client) DeleteGatewaySets(request *DeleteGatewaySetsRequest) (response *DeleteGatewaySetsResponse, err error) {
	if request == nil {
		request = NewDeleteGatewaySetsRequest()
	}
	response = NewDeleteGatewaySetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWhiteListsRequest() (request *DescribeWhiteListsRequest) {
	request = &DescribeWhiteListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "DescribeWhiteLists")
	return
}

func NewDescribeWhiteListsResponse() (response *DescribeWhiteListsResponse) {
	response = &DescribeWhiteListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取白名单列表
func (c *Client) DescribeWhiteLists(request *DescribeWhiteListsRequest) (response *DescribeWhiteListsResponse, err error) {
	if request == nil {
		request = NewDescribeWhiteListsRequest()
	}
	response = NewDescribeWhiteListsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateWhiteListRequest() (request *UpdateWhiteListRequest) {
	request = &UpdateWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "UpdateWhiteList")
	return
}

func NewUpdateWhiteListResponse() (response *UpdateWhiteListResponse) {
	response = &UpdateWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于更新一个集群关联的策略信息
func (c *Client) UpdateWhiteList(request *UpdateWhiteListRequest) (response *UpdateWhiteListResponse, err error) {
	if request == nil {
		request = NewUpdateWhiteListRequest()
	}
	response = NewUpdateWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateControllerSetRequest() (request *CreateControllerSetRequest) {
	request = &CreateControllerSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "CreateControllerSet")
	return
}

func NewCreateControllerSetResponse() (response *CreateControllerSetResponse) {
	response = &CreateControllerSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于创建一个controller集群
func (c *Client) CreateControllerSet(request *CreateControllerSetRequest) (response *CreateControllerSetResponse, err error) {
	if request == nil {
		request = NewCreateControllerSetRequest()
	}
	response = NewCreateControllerSetResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteControllersRequest() (request *DeleteControllersRequest) {
	request = &DeleteControllersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "DeleteControllers")
	return
}

func NewDeleteControllersResponse() (response *DeleteControllersResponse) {
	response = &DeleteControllersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于批量删除controller设备
func (c *Client) DeleteControllers(request *DeleteControllersRequest) (response *DeleteControllersResponse, err error) {
	if request == nil {
		request = NewDeleteControllersRequest()
	}
	response = NewDeleteControllersResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteWhiteListsRequest() (request *DeleteWhiteListsRequest) {
	request = &DeleteWhiteListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "DeleteWhiteLists")
	return
}

func NewDeleteWhiteListsResponse() (response *DeleteWhiteListsResponse) {
	response = &DeleteWhiteListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除集群关联策略
func (c *Client) DeleteWhiteLists(request *DeleteWhiteListsRequest) (response *DeleteWhiteListsResponse, err error) {
	if request == nil {
		request = NewDeleteWhiteListsRequest()
	}
	response = NewDeleteWhiteListsResponse()
	err = c.Send(request, response)
	return
}

func NewMigrateLimitRuleRequest() (request *MigrateLimitRuleRequest) {
	request = &MigrateLimitRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "MigrateLimitRule")
	return
}

func NewMigrateLimitRuleResponse() (response *MigrateLimitRuleResponse) {
	response = &MigrateLimitRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量迁移限速规则
func (c *Client) MigrateLimitRule(request *MigrateLimitRuleRequest) (response *MigrateLimitRuleResponse, err error) {
	if request == nil {
		request = NewMigrateLimitRuleRequest()
	}
	response = NewMigrateLimitRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDevicesRequest() (request *DescribeDevicesRequest) {
	request = &DescribeDevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("seastone", APIVersion, "DescribeDevices")
	return
}

func NewDescribeDevicesResponse() (response *DescribeDevicesResponse) {
	response = &DescribeDevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群网关节点列表
func (c *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
	if request == nil {
		request = NewDescribeDevicesRequest()
	}
	response = NewDescribeDevicesResponse()
	err = c.Send(request, response)
	return
}
