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

package v20220518

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2022-05-18"

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

func NewCreateFlowRequest() (request *CreateFlowRequest) {
	request = &CreateFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oapproval", APIVersion, "CreateFlow")
	return
}

func NewCreateFlowResponse() (response *CreateFlowResponse) {
	response = &CreateFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建运营端审批流
func (c *Client) CreateFlow(request *CreateFlowRequest) (response *CreateFlowResponse, err error) {
	if request == nil {
		request = NewCreateFlowRequest()
	}
	response = NewCreateFlowResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCurrApprovalDocRequest() (request *QueryCurrApprovalDocRequest) {
	request = &QueryCurrApprovalDocRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oapproval", APIVersion, "QueryCurrApprovalDoc")
	return
}

func NewQueryCurrApprovalDocResponse() (response *QueryCurrApprovalDocResponse) {
	response = &QueryCurrApprovalDocResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营端当前待审批列表
func (c *Client) QueryCurrApprovalDoc(request *QueryCurrApprovalDocRequest) (response *QueryCurrApprovalDocResponse, err error) {
	if request == nil {
		request = NewQueryCurrApprovalDocRequest()
	}
	response = NewQueryCurrApprovalDocResponse()
	err = c.Send(request, response)
	return
}

func NewQueryApprovalDocDetailRequest() (request *QueryApprovalDocDetailRequest) {
	request = &QueryApprovalDocDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oapproval", APIVersion, "QueryApprovalDocDetail")
	return
}

func NewQueryApprovalDocDetailResponse() (response *QueryApprovalDocDetailResponse) {
	response = &QueryApprovalDocDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取运营端审批单详情
func (c *Client) QueryApprovalDocDetail(request *QueryApprovalDocDetailRequest) (response *QueryApprovalDocDetailResponse, err error) {
	if request == nil {
		request = NewQueryApprovalDocDetailRequest()
	}
	response = NewQueryApprovalDocDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApprovalFlowRequest() (request *DeleteApprovalFlowRequest) {
	request = &DeleteApprovalFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oapproval", APIVersion, "DeleteApprovalFlow")
	return
}

func NewDeleteApprovalFlowResponse() (response *DeleteApprovalFlowResponse) {
	response = &DeleteApprovalFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除审批流
func (c *Client) DeleteApprovalFlow(request *DeleteApprovalFlowRequest) (response *DeleteApprovalFlowResponse, err error) {
	if request == nil {
		request = NewDeleteApprovalFlowRequest()
	}
	response = NewDeleteApprovalFlowResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApprovalFlowRequest() (request *ModifyApprovalFlowRequest) {
	request = &ModifyApprovalFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oapproval", APIVersion, "ModifyApprovalFlow")
	return
}

func NewModifyApprovalFlowResponse() (response *ModifyApprovalFlowResponse) {
	response = &ModifyApprovalFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改审批流
func (c *Client) ModifyApprovalFlow(request *ModifyApprovalFlowRequest) (response *ModifyApprovalFlowResponse, err error) {
	if request == nil {
		request = NewModifyApprovalFlowRequest()
	}
	response = NewModifyApprovalFlowResponse()
	err = c.Send(request, response)
	return
}

func NewPerformApprovalRequest() (request *PerformApprovalRequest) {
	request = &PerformApprovalRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oapproval", APIVersion, "PerformApproval")
	return
}

func NewPerformApprovalResponse() (response *PerformApprovalResponse) {
	response = &PerformApprovalResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端用户执行审批操作
func (c *Client) PerformApproval(request *PerformApprovalRequest) (response *PerformApprovalResponse, err error) {
	if request == nil {
		request = NewPerformApprovalRequest()
	}
	response = NewPerformApprovalResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApprovalActionRequest() (request *CreateApprovalActionRequest) {
	request = &CreateApprovalActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oapproval", APIVersion, "CreateApprovalAction")
	return
}

func NewCreateApprovalActionResponse() (response *CreateApprovalActionResponse) {
	response = &CreateApprovalActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建审批业务
func (c *Client) CreateApprovalAction(request *CreateApprovalActionRequest) (response *CreateApprovalActionResponse, err error) {
	if request == nil {
		request = NewCreateApprovalActionRequest()
	}
	response = NewCreateApprovalActionResponse()
	err = c.Send(request, response)
	return
}

func NewOperateFlowStatusRequest() (request *OperateFlowStatusRequest) {
	request = &OperateFlowStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oapproval", APIVersion, "OperateFlowStatus")
	return
}

func NewOperateFlowStatusResponse() (response *OperateFlowStatusResponse) {
	response = &OperateFlowStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 操作审批流启用状态
func (c *Client) OperateFlowStatus(request *OperateFlowStatusRequest) (response *OperateFlowStatusResponse, err error) {
	if request == nil {
		request = NewOperateFlowStatusRequest()
	}
	response = NewOperateFlowStatusResponse()
	err = c.Send(request, response)
	return
}

func NewGetFlowDetailRequest() (request *GetFlowDetailRequest) {
	request = &GetFlowDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oapproval", APIVersion, "GetFlowDetail")
	return
}

func NewGetFlowDetailResponse() (response *GetFlowDetailResponse) {
	response = &GetFlowDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取运营端审批流详情
func (c *Client) GetFlowDetail(request *GetFlowDetailRequest) (response *GetFlowDetailResponse, err error) {
	if request == nil {
		request = NewGetFlowDetailRequest()
	}
	response = NewGetFlowDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApprovalActionRequest() (request *DeleteApprovalActionRequest) {
	request = &DeleteApprovalActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oapproval", APIVersion, "DeleteApprovalAction")
	return
}

func NewDeleteApprovalActionResponse() (response *DeleteApprovalActionResponse) {
	response = &DeleteApprovalActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除审批业务， 预设业务，已关联审批流的业务不允许删除
func (c *Client) DeleteApprovalAction(request *DeleteApprovalActionRequest) (response *DeleteApprovalActionResponse, err error) {
	if request == nil {
		request = NewDeleteApprovalActionRequest()
	}
	response = NewDeleteApprovalActionResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApprovalActionRequest() (request *UpdateApprovalActionRequest) {
	request = &UpdateApprovalActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oapproval", APIVersion, "UpdateApprovalAction")
	return
}

func NewUpdateApprovalActionResponse() (response *UpdateApprovalActionResponse) {
	response = &UpdateApprovalActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新审批业务
func (c *Client) UpdateApprovalAction(request *UpdateApprovalActionRequest) (response *UpdateApprovalActionResponse, err error) {
	if request == nil {
		request = NewUpdateApprovalActionRequest()
	}
	response = NewUpdateApprovalActionResponse()
	err = c.Send(request, response)
	return
}

func NewQueryActionSetRequest() (request *QueryActionSetRequest) {
	request = &QueryActionSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oapproval", APIVersion, "QueryActionSet")
	return
}

func NewQueryActionSetResponse() (response *QueryActionSetResponse) {
	response = &QueryActionSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询接入审批的业务接口列表
func (c *Client) QueryActionSet(request *QueryActionSetRequest) (response *QueryActionSetResponse, err error) {
	if request == nil {
		request = NewQueryActionSetRequest()
	}
	response = NewQueryActionSetResponse()
	err = c.Send(request, response)
	return
}

func NewGetApprovedSetRequest() (request *GetApprovedSetRequest) {
	request = &GetApprovedSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oapproval", APIVersion, "GetApprovedSet")
	return
}

func NewGetApprovedSetResponse() (response *GetApprovedSetResponse) {
	response = &GetApprovedSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营端当前已审批列表
func (c *Client) GetApprovedSet(request *GetApprovedSetRequest) (response *GetApprovedSetResponse, err error) {
	if request == nil {
		request = NewGetApprovedSetRequest()
	}
	response = NewGetApprovedSetResponse()
	err = c.Send(request, response)
	return
}

func NewQueryUsersUnderActionRequest() (request *QueryUsersUnderActionRequest) {
	request = &QueryUsersUnderActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oapproval", APIVersion, "QueryUsersUnderAction")
	return
}

func NewQueryUsersUnderActionResponse() (response *QueryUsersUnderActionResponse) {
	response = &QueryUsersUnderActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取运营端审批接口已关联的用户
func (c *Client) QueryUsersUnderAction(request *QueryUsersUnderActionRequest) (response *QueryUsersUnderActionResponse, err error) {
	if request == nil {
		request = NewQueryUsersUnderActionRequest()
	}
	response = NewQueryUsersUnderActionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowSetRequest() (request *DescribeFlowSetRequest) {
	request = &DescribeFlowSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oapproval", APIVersion, "DescribeFlowSet")
	return
}

func NewDescribeFlowSetResponse() (response *DescribeFlowSetResponse) {
	response = &DescribeFlowSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取运营端审批流列表
func (c *Client) DescribeFlowSet(request *DescribeFlowSetRequest) (response *DescribeFlowSetResponse, err error) {
	if request == nil {
		request = NewDescribeFlowSetRequest()
	}
	response = NewDescribeFlowSetResponse()
	err = c.Send(request, response)
	return
}

func NewQueryApprovalDocRequest() (request *QueryApprovalDocRequest) {
	request = &QueryApprovalDocRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oapproval", APIVersion, "QueryApprovalDoc")
	return
}

func NewQueryApprovalDocResponse() (response *QueryApprovalDocResponse) {
	response = &QueryApprovalDocResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营端审批单列表
func (c *Client) QueryApprovalDoc(request *QueryApprovalDocRequest) (response *QueryApprovalDocResponse, err error) {
	if request == nil {
		request = NewQueryApprovalDocRequest()
	}
	response = NewQueryApprovalDocResponse()
	err = c.Send(request, response)
	return
}
