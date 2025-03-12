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

package v20200518

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-05-18"

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

func NewDeleteFlavorExRequest() (request *DeleteFlavorExRequest) {
	request = &DeleteFlavorExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DeleteFlavorEx")
	return
}

func NewDeleteFlavorExResponse() (response *DeleteFlavorExResponse) {
	response = &DeleteFlavorExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除机型信息
func (c *Client) DeleteFlavorEx(request *DeleteFlavorExRequest) (response *DeleteFlavorExResponse, err error) {
	if request == nil {
		request = NewDeleteFlavorExRequest()
	}
	response = NewDeleteFlavorExResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateBmsQuotaRequest() (request *UpdateBmsQuotaRequest) {
	request = &UpdateBmsQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "UpdateBmsQuota")
	return
}

func NewUpdateBmsQuotaResponse() (response *UpdateBmsQuotaResponse) {
	response = &UpdateBmsQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改租户的裸金属服务器配额信息
func (c *Client) UpdateBmsQuota(request *UpdateBmsQuotaRequest) (response *UpdateBmsQuotaResponse, err error) {
	if request == nil {
		request = NewUpdateBmsQuotaRequest()
	}
	response = NewUpdateBmsQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateFlavorExRequest() (request *UpdateFlavorExRequest) {
	request = &UpdateFlavorExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "UpdateFlavorEx")
	return
}

func NewUpdateFlavorExResponse() (response *UpdateFlavorExResponse) {
	response = &UpdateFlavorExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改机型信息
func (c *Client) UpdateFlavorEx(request *UpdateFlavorExRequest) (response *UpdateFlavorExResponse, err error) {
	if request == nil {
		request = NewUpdateFlavorExRequest()
	}
	response = NewUpdateFlavorExResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateDefaultQuotaRequest() (request *UpdateDefaultQuotaRequest) {
	request = &UpdateDefaultQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "UpdateDefaultQuota")
	return
}

func NewUpdateDefaultQuotaResponse() (response *UpdateDefaultQuotaResponse) {
	response = &UpdateDefaultQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改裸金属服务器的默认配额信息
func (c *Client) UpdateDefaultQuota(request *UpdateDefaultQuotaRequest) (response *UpdateDefaultQuotaResponse, err error) {
	if request == nil {
		request = NewUpdateDefaultQuotaRequest()
	}
	response = NewUpdateDefaultQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewAddBmsTypeBillExRequest() (request *AddBmsTypeBillExRequest) {
	request = &AddBmsTypeBillExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "AddBmsTypeBillEx")
	return
}

func NewAddBmsTypeBillExResponse() (response *AddBmsTypeBillExResponse) {
	response = &AddBmsTypeBillExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增机型四层定义上报
func (c *Client) AddBmsTypeBillEx(request *AddBmsTypeBillExRequest) (response *AddBmsTypeBillExResponse, err error) {
	if request == nil {
		request = NewAddBmsTypeBillExRequest()
	}
	response = NewAddBmsTypeBillExResponse()
	err = c.Send(request, response)
	return
}

func NewAddSwitchExRequest() (request *AddSwitchExRequest) {
	request = &AddSwitchExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "AddSwitchEx")
	return
}

func NewAddSwitchExResponse() (response *AddSwitchExResponse) {
	response = &AddSwitchExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加交换机信息
func (c *Client) AddSwitchEx(request *AddSwitchExRequest) (response *AddSwitchExResponse, err error) {
	if request == nil {
		request = NewAddSwitchExRequest()
	}
	response = NewAddSwitchExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSwitchtypeExRequest() (request *DescribeSwitchtypeExRequest) {
	request = &DescribeSwitchtypeExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeSwitchtypeEx")
	return
}

func NewDescribeSwitchtypeExResponse() (response *DescribeSwitchtypeExResponse) {
	response = &DescribeSwitchtypeExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询交换机厂商类型信息
func (c *Client) DescribeSwitchtypeEx(request *DescribeSwitchtypeExRequest) (response *DescribeSwitchtypeExResponse, err error) {
	if request == nil {
		request = NewDescribeSwitchtypeExRequest()
	}
	response = NewDescribeSwitchtypeExResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSwitchUserExRequest() (request *UpdateSwitchUserExRequest) {
	request = &UpdateSwitchUserExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "UpdateSwitchUserEx")
	return
}

func NewUpdateSwitchUserExResponse() (response *UpdateSwitchUserExResponse) {
	response = &UpdateSwitchUserExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改交换机用户信息
func (c *Client) UpdateSwitchUserEx(request *UpdateSwitchUserExRequest) (response *UpdateSwitchUserExResponse, err error) {
	if request == nil {
		request = NewUpdateSwitchUserExRequest()
	}
	response = NewUpdateSwitchUserExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNodeExRequest() (request *DescribeNodeExRequest) {
	request = &DescribeNodeExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeNodeEx")
	return
}

func NewDescribeNodeExResponse() (response *DescribeNodeExResponse) {
	response = &DescribeNodeExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取裸金属资源池资源信息
func (c *Client) DescribeNodeEx(request *DescribeNodeExRequest) (response *DescribeNodeExResponse, err error) {
	if request == nil {
		request = NewDescribeNodeExRequest()
	}
	response = NewDescribeNodeExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSellstateExRequest() (request *DescribeSellstateExRequest) {
	request = &DescribeSellstateExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeSellstateEx")
	return
}

func NewDescribeSellstateExResponse() (response *DescribeSellstateExResponse) {
	response = &DescribeSellstateExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询机型的售卖状态
func (c *Client) DescribeSellstateEx(request *DescribeSellstateExRequest) (response *DescribeSellstateExResponse, err error) {
	if request == nil {
		request = NewDescribeSellstateExRequest()
	}
	response = NewDescribeSellstateExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSwitchExRequest() (request *DescribeSwitchExRequest) {
	request = &DescribeSwitchExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeSwitchEx")
	return
}

func NewDescribeSwitchExResponse() (response *DescribeSwitchExResponse) {
	response = &DescribeSwitchExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询交换机信息
func (c *Client) DescribeSwitchEx(request *DescribeSwitchExRequest) (response *DescribeSwitchExResponse, err error) {
	if request == nil {
		request = NewDescribeSwitchExRequest()
	}
	response = NewDescribeSwitchExResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateNodeExRequest() (request *UpdateNodeExRequest) {
	request = &UpdateNodeExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "UpdateNodeEx")
	return
}

func NewUpdateNodeExResponse() (response *UpdateNodeExResponse) {
	response = &UpdateNodeExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改裸金属服务器信息
func (c *Client) UpdateNodeEx(request *UpdateNodeExRequest) (response *UpdateNodeExResponse, err error) {
	if request == nil {
		request = NewUpdateNodeExRequest()
	}
	response = NewUpdateNodeExResponse()
	err = c.Send(request, response)
	return
}

func NewAddNodeToVPCExRequest() (request *AddNodeToVPCExRequest) {
	request = &AddNodeToVPCExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "AddNodeToVPCEx")
	return
}

func NewAddNodeToVPCExResponse() (response *AddNodeToVPCExResponse) {
	response = &AddNodeToVPCExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// vpc添加资源
func (c *Client) AddNodeToVPCEx(request *AddNodeToVPCExRequest) (response *AddNodeToVPCExResponse, err error) {
	if request == nil {
		request = NewAddNodeToVPCExRequest()
	}
	response = NewAddNodeToVPCExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBmsTasksExRequest() (request *DescribeBmsTasksExRequest) {
	request = &DescribeBmsTasksExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeBmsTasksEx")
	return
}

func NewDescribeBmsTasksExResponse() (response *DescribeBmsTasksExResponse) {
	response = &DescribeBmsTasksExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 裸金属任务信息显示
func (c *Client) DescribeBmsTasksEx(request *DescribeBmsTasksExRequest) (response *DescribeBmsTasksExResponse, err error) {
	if request == nil {
		request = NewDescribeBmsTasksExRequest()
	}
	response = NewDescribeBmsTasksExResponse()
	err = c.Send(request, response)
	return
}

func NewAddSwitchUserExRequest() (request *AddSwitchUserExRequest) {
	request = &AddSwitchUserExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "AddSwitchUserEx")
	return
}

func NewAddSwitchUserExResponse() (response *AddSwitchUserExResponse) {
	response = &AddSwitchUserExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加交换机用户信息
func (c *Client) AddSwitchUserEx(request *AddSwitchUserExRequest) (response *AddSwitchUserExResponse, err error) {
	if request == nil {
		request = NewAddSwitchUserExRequest()
	}
	response = NewAddSwitchUserExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlavorExRequest() (request *DescribeFlavorExRequest) {
	request = &DescribeFlavorExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeFlavorEx")
	return
}

func NewDescribeFlavorExResponse() (response *DescribeFlavorExResponse) {
	response = &DescribeFlavorExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询机型信息
func (c *Client) DescribeFlavorEx(request *DescribeFlavorExRequest) (response *DescribeFlavorExResponse, err error) {
	if request == nil {
		request = NewDescribeFlavorExRequest()
	}
	response = NewDescribeFlavorExResponse()
	err = c.Send(request, response)
	return
}

func NewAddFlavorExRequest() (request *AddFlavorExRequest) {
	request = &AddFlavorExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "AddFlavorEx")
	return
}

func NewAddFlavorExResponse() (response *AddFlavorExResponse) {
	response = &AddFlavorExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加机型信息
func (c *Client) AddFlavorEx(request *AddFlavorExRequest) (response *AddFlavorExResponse, err error) {
	if request == nil {
		request = NewAddFlavorExRequest()
	}
	response = NewAddFlavorExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskExRequest() (request *DescribeTaskExRequest) {
	request = &DescribeTaskExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeTaskEx")
	return
}

func NewDescribeTaskExResponse() (response *DescribeTaskExResponse) {
	response = &DescribeTaskExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 任务查询
func (c *Client) DescribeTaskEx(request *DescribeTaskExRequest) (response *DescribeTaskExResponse, err error) {
	if request == nil {
		request = NewDescribeTaskExRequest()
	}
	response = NewDescribeTaskExResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSwitchExRequest() (request *UpdateSwitchExRequest) {
	request = &UpdateSwitchExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "UpdateSwitchEx")
	return
}

func NewUpdateSwitchExResponse() (response *UpdateSwitchExResponse) {
	response = &UpdateSwitchExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改交换机信息
func (c *Client) UpdateSwitchEx(request *UpdateSwitchExRequest) (response *UpdateSwitchExResponse, err error) {
	if request == nil {
		request = NewUpdateSwitchExRequest()
	}
	response = NewUpdateSwitchExResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNodeExRequest() (request *DeleteNodeExRequest) {
	request = &DeleteNodeExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DeleteNodeEx")
	return
}

func NewDeleteNodeExResponse() (response *DeleteNodeExResponse) {
	response = &DeleteNodeExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除资源
func (c *Client) DeleteNodeEx(request *DeleteNodeExRequest) (response *DeleteNodeExResponse, err error) {
	if request == nil {
		request = NewDeleteNodeExRequest()
	}
	response = NewDeleteNodeExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBmsDefaultQuotaRequest() (request *DescribeBmsDefaultQuotaRequest) {
	request = &DescribeBmsDefaultQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeBmsDefaultQuota")
	return
}

func NewDescribeBmsDefaultQuotaResponse() (response *DescribeBmsDefaultQuotaResponse) {
	response = &DescribeBmsDefaultQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取裸金属服务器的默认配额信息
func (c *Client) DescribeBmsDefaultQuota(request *DescribeBmsDefaultQuotaRequest) (response *DescribeBmsDefaultQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeBmsDefaultQuotaRequest()
	}
	response = NewDescribeBmsDefaultQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeObmsPoolNodesToAddExRequest() (request *DescribeObmsPoolNodesToAddExRequest) {
	request = &DescribeObmsPoolNodesToAddExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeObmsPoolNodesToAddEx")
	return
}

func NewDescribeObmsPoolNodesToAddExResponse() (response *DescribeObmsPoolNodesToAddExResponse) {
	response = &DescribeObmsPoolNodesToAddExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源池资源
func (c *Client) DescribeObmsPoolNodesToAddEx(request *DescribeObmsPoolNodesToAddExRequest) (response *DescribeObmsPoolNodesToAddExResponse, err error) {
	if request == nil {
		request = NewDescribeObmsPoolNodesToAddExRequest()
	}
	response = NewDescribeObmsPoolNodesToAddExResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSwitchExRequest() (request *DeleteSwitchExRequest) {
	request = &DeleteSwitchExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DeleteSwitchEx")
	return
}

func NewDeleteSwitchExResponse() (response *DeleteSwitchExResponse) {
	response = &DeleteSwitchExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除交换机记录
func (c *Client) DeleteSwitchEx(request *DeleteSwitchExRequest) (response *DeleteSwitchExResponse, err error) {
	if request == nil {
		request = NewDeleteSwitchExRequest()
	}
	response = NewDeleteSwitchExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSwitchUserExRequest() (request *DescribeSwitchUserExRequest) {
	request = &DescribeSwitchUserExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeSwitchUserEx")
	return
}

func NewDescribeSwitchUserExResponse() (response *DescribeSwitchUserExResponse) {
	response = &DescribeSwitchUserExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询交换机用户信息
func (c *Client) DescribeSwitchUserEx(request *DescribeSwitchUserExRequest) (response *DescribeSwitchUserExResponse, err error) {
	if request == nil {
		request = NewDescribeSwitchUserExRequest()
	}
	response = NewDescribeSwitchUserExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeObmsNodeFilterExRequest() (request *DescribeObmsNodeFilterExRequest) {
	request = &DescribeObmsNodeFilterExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeObmsNodeFilterEx")
	return
}

func NewDescribeObmsNodeFilterExResponse() (response *DescribeObmsNodeFilterExResponse) {
	response = &DescribeObmsNodeFilterExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 裸金属资源池资源资源信息
func (c *Client) DescribeObmsNodeFilterEx(request *DescribeObmsNodeFilterExRequest) (response *DescribeObmsNodeFilterExResponse, err error) {
	if request == nil {
		request = NewDescribeObmsNodeFilterExRequest()
	}
	response = NewDescribeObmsNodeFilterExResponse()
	err = c.Send(request, response)
	return
}

func NewAddNodeToFlavorExRequest() (request *AddNodeToFlavorExRequest) {
	request = &AddNodeToFlavorExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "AddNodeToFlavorEx")
	return
}

func NewAddNodeToFlavorExResponse() (response *AddNodeToFlavorExResponse) {
	response = &AddNodeToFlavorExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加资源到机型中
func (c *Client) AddNodeToFlavorEx(request *AddNodeToFlavorExRequest) (response *AddNodeToFlavorExResponse, err error) {
	if request == nil {
		request = NewAddNodeToFlavorExRequest()
	}
	response = NewAddNodeToFlavorExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBmsQuotaRequest() (request *DescribeBmsQuotaRequest) {
	request = &DescribeBmsQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeBmsQuota")
	return
}

func NewDescribeBmsQuotaResponse() (response *DescribeBmsQuotaResponse) {
	response = &DescribeBmsQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据输入的AppId列表获取租户裸金属配额
func (c *Client) DescribeBmsQuota(request *DescribeBmsQuotaRequest) (response *DescribeBmsQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeBmsQuotaRequest()
	}
	response = NewDescribeBmsQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewReturnNodeToObmsPoolExRequest() (request *ReturnNodeToObmsPoolExRequest) {
	request = &ReturnNodeToObmsPoolExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "ReturnNodeToObmsPoolEx")
	return
}

func NewReturnNodeToObmsPoolExResponse() (response *ReturnNodeToObmsPoolExResponse) {
	response = &ReturnNodeToObmsPoolExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 退还资源池机器到dcos
func (c *Client) ReturnNodeToObmsPoolEx(request *ReturnNodeToObmsPoolExRequest) (response *ReturnNodeToObmsPoolExResponse, err error) {
	if request == nil {
		request = NewReturnNodeToObmsPoolExRequest()
	}
	response = NewReturnNodeToObmsPoolExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBmsNodesExRequest() (request *DescribeBmsNodesExRequest) {
	request = &DescribeBmsNodesExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeBmsNodesEx")
	return
}

func NewDescribeBmsNodesExResponse() (response *DescribeBmsNodesExResponse) {
	response = &DescribeBmsNodesExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询裸金属资源
func (c *Client) DescribeBmsNodesEx(request *DescribeBmsNodesExRequest) (response *DescribeBmsNodesExResponse, err error) {
	if request == nil {
		request = NewDescribeBmsNodesExRequest()
	}
	response = NewDescribeBmsNodesExResponse()
	err = c.Send(request, response)
	return
}
