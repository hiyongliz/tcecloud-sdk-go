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

package v20181008

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2018-10-08"

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

func NewCreateFieldRequest() (request *CreateFieldRequest) {
	request = &CreateFieldRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "CreateField")
	return
}

func NewCreateFieldResponse() (response *CreateFieldResponse) {
	response = &CreateFieldResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建表单字段
func (c *Client) CreateField(request *CreateFieldRequest) (response *CreateFieldResponse, err error) {
	if request == nil {
		request = NewCreateFieldRequest()
	}
	response = NewCreateFieldResponse()
	err = c.Send(request, response)
	return
}

func NewStaffRespondRequest() (request *StaffRespondRequest) {
	request = &StaffRespondRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "StaffRespond")
	return
}

func NewStaffRespondResponse() (response *StaffRespondResponse) {
	response = &StaffRespondResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客服回复
func (c *Client) StaffRespond(request *StaffRespondRequest) (response *StaffRespondResponse, err error) {
	if request == nil {
		request = NewStaffRespondRequest()
	}
	response = NewStaffRespondResponse()
	err = c.Send(request, response)
	return
}

func NewStaffSignOutRequest() (request *StaffSignOutRequest) {
	request = &StaffSignOutRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "StaffSignOut")
	return
}

func NewStaffSignOutResponse() (response *StaffSignOutResponse) {
	response = &StaffSignOutResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// StaffSignOut
func (c *Client) StaffSignOut(request *StaffSignOutRequest) (response *StaffSignOutResponse, err error) {
	if request == nil {
		request = NewStaffSignOutRequest()
	}
	response = NewStaffSignOutResponse()
	err = c.Send(request, response)
	return
}

func NewGetSecretContentRequest() (request *GetSecretContentRequest) {
	request = &GetSecretContentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "GetSecretContent")
	return
}

func NewGetSecretContentResponse() (response *GetSecretContentResponse) {
	response = &GetSecretContentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取机密信息
func (c *Client) GetSecretContent(request *GetSecretContentRequest) (response *GetSecretContentResponse, err error) {
	if request == nil {
		request = NewGetSecretContentRequest()
	}
	response = NewGetSecretContentResponse()
	err = c.Send(request, response)
	return
}

func NewGetStaffsRequest() (request *GetStaffsRequest) {
	request = &GetStaffsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "GetStaffs")
	return
}

func NewGetStaffsResponse() (response *GetStaffsResponse) {
	response = &GetStaffsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取客服列表
func (c *Client) GetStaffs(request *GetStaffsRequest) (response *GetStaffsResponse, err error) {
	if request == nil {
		request = NewGetStaffsRequest()
	}
	response = NewGetStaffsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMessagesRequest() (request *DescribeMessagesRequest) {
	request = &DescribeMessagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "DescribeMessages")
	return
}

func NewDescribeMessagesResponse() (response *DescribeMessagesResponse) {
	response = &DescribeMessagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询客服站内消息
func (c *Client) DescribeMessages(request *DescribeMessagesRequest) (response *DescribeMessagesResponse, err error) {
	if request == nil {
		request = NewDescribeMessagesRequest()
	}
	response = NewDescribeMessagesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateFieldRequest() (request *UpdateFieldRequest) {
	request = &UpdateFieldRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "UpdateField")
	return
}

func NewUpdateFieldResponse() (response *UpdateFieldResponse) {
	response = &UpdateFieldResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新表单字段
func (c *Client) UpdateField(request *UpdateFieldRequest) (response *UpdateFieldResponse, err error) {
	if request == nil {
		request = NewUpdateFieldRequest()
	}
	response = NewUpdateFieldResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCategoryRequest() (request *CreateCategoryRequest) {
	request = &CreateCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "CreateCategory")
	return
}

func NewCreateCategoryResponse() (response *CreateCategoryResponse) {
	response = &CreateCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建产品分类
func (c *Client) CreateCategory(request *CreateCategoryRequest) (response *CreateCategoryResponse, err error) {
	if request == nil {
		request = NewCreateCategoryRequest()
	}
	response = NewCreateCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewCreateStaffRequest() (request *CreateStaffRequest) {
	request = &CreateStaffRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "CreateStaff")
	return
}

func NewCreateStaffResponse() (response *CreateStaffResponse) {
	response = &CreateStaffResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建客服
func (c *Client) CreateStaff(request *CreateStaffRequest) (response *CreateStaffResponse, err error) {
	if request == nil {
		request = NewCreateStaffRequest()
	}
	response = NewCreateStaffResponse()
	err = c.Send(request, response)
	return
}

func NewGetGroupInfoRequest() (request *GetGroupInfoRequest) {
	request = &GetGroupInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "GetGroupInfo")
	return
}

func NewGetGroupInfoResponse() (response *GetGroupInfoResponse) {
	response = &GetGroupInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取客服组详情
func (c *Client) GetGroupInfo(request *GetGroupInfoRequest) (response *GetGroupInfoResponse, err error) {
	if request == nil {
		request = NewGetGroupInfoRequest()
	}
	response = NewGetGroupInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGetScheduleStaffRequest() (request *GetScheduleStaffRequest) {
	request = &GetScheduleStaffRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "GetScheduleStaff")
	return
}

func NewGetScheduleStaffResponse() (response *GetScheduleStaffResponse) {
	response = &GetScheduleStaffResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取排班客服
func (c *Client) GetScheduleStaff(request *GetScheduleStaffRequest) (response *GetScheduleStaffResponse, err error) {
	if request == nil {
		request = NewGetScheduleStaffRequest()
	}
	response = NewGetScheduleStaffResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCategoryRequest() (request *DeleteCategoryRequest) {
	request = &DeleteCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "DeleteCategory")
	return
}

func NewDeleteCategoryResponse() (response *DeleteCategoryResponse) {
	response = &DeleteCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除工单分类
func (c *Client) DeleteCategory(request *DeleteCategoryRequest) (response *DeleteCategoryResponse, err error) {
	if request == nil {
		request = NewDeleteCategoryRequest()
	}
	response = NewDeleteCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGroupStaffRequest() (request *DeleteGroupStaffRequest) {
	request = &DeleteGroupStaffRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "DeleteGroupStaff")
	return
}

func NewDeleteGroupStaffResponse() (response *DeleteGroupStaffResponse) {
	response = &DeleteGroupStaffResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除客服组成员
func (c *Client) DeleteGroupStaff(request *DeleteGroupStaffRequest) (response *DeleteGroupStaffResponse, err error) {
	if request == nil {
		request = NewDeleteGroupStaffRequest()
	}
	response = NewDeleteGroupStaffResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteStaffStatusRequest() (request *DeleteStaffStatusRequest) {
	request = &DeleteStaffStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "DeleteStaffStatus")
	return
}

func NewDeleteStaffStatusResponse() (response *DeleteStaffStatusResponse) {
	response = &DeleteStaffStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 强制客服离线
func (c *Client) DeleteStaffStatus(request *DeleteStaffStatusRequest) (response *DeleteStaffStatusResponse, err error) {
	if request == nil {
		request = NewDeleteStaffStatusRequest()
	}
	response = NewDeleteStaffStatusResponse()
	err = c.Send(request, response)
	return
}

func NewStaffEnquireCouldFinishTicketRequest() (request *StaffEnquireCouldFinishTicketRequest) {
	request = &StaffEnquireCouldFinishTicketRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "StaffEnquireCouldFinishTicket")
	return
}

func NewStaffEnquireCouldFinishTicketResponse() (response *StaffEnquireCouldFinishTicketResponse) {
	response = &StaffEnquireCouldFinishTicketResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 待客户结单
func (c *Client) StaffEnquireCouldFinishTicket(request *StaffEnquireCouldFinishTicketRequest) (response *StaffEnquireCouldFinishTicketResponse, err error) {
	if request == nil {
		request = NewStaffEnquireCouldFinishTicketRequest()
	}
	response = NewStaffEnquireCouldFinishTicketResponse()
	err = c.Send(request, response)
	return
}

func NewExportScheduleRequest() (request *ExportScheduleRequest) {
	request = &ExportScheduleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "ExportSchedule")
	return
}

func NewExportScheduleResponse() (response *ExportScheduleResponse) {
	response = &ExportScheduleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出排班记录
func (c *Client) ExportSchedule(request *ExportScheduleRequest) (response *ExportScheduleResponse, err error) {
	if request == nil {
		request = NewExportScheduleRequest()
	}
	response = NewExportScheduleResponse()
	err = c.Send(request, response)
	return
}

func NewStaffReadMessageRequest() (request *StaffReadMessageRequest) {
	request = &StaffReadMessageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "StaffReadMessage")
	return
}

func NewStaffReadMessageResponse() (response *StaffReadMessageResponse) {
	response = &StaffReadMessageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客服阅读站内消息
func (c *Client) StaffReadMessage(request *StaffReadMessageRequest) (response *StaffReadMessageResponse, err error) {
	if request == nil {
		request = NewStaffReadMessageRequest()
	}
	response = NewStaffReadMessageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTicketRequest() (request *DescribeTicketRequest) {
	request = &DescribeTicketRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "DescribeTicket")
	return
}

func NewDescribeTicketResponse() (response *DescribeTicketResponse) {
	response = &DescribeTicketResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 工单详情
func (c *Client) DescribeTicket(request *DescribeTicketRequest) (response *DescribeTicketResponse, err error) {
	if request == nil {
		request = NewDescribeTicketRequest()
	}
	response = NewDescribeTicketResponse()
	err = c.Send(request, response)
	return
}

func NewStaffSupplyRequest() (request *StaffSupplyRequest) {
	request = &StaffSupplyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "StaffSupply")
	return
}

func NewStaffSupplyResponse() (response *StaffSupplyResponse) {
	response = &StaffSupplyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客服内部补充（可转单）
func (c *Client) StaffSupply(request *StaffSupplyRequest) (response *StaffSupplyResponse, err error) {
	if request == nil {
		request = NewStaffSupplyRequest()
	}
	response = NewStaffSupplyResponse()
	err = c.Send(request, response)
	return
}

func NewStaffSignInRequest() (request *StaffSignInRequest) {
	request = &StaffSignInRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "StaffSignIn")
	return
}

func NewStaffSignInResponse() (response *StaffSignInResponse) {
	response = &StaffSignInResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// StaffSignIn
func (c *Client) StaffSignIn(request *StaffSignInRequest) (response *StaffSignInResponse, err error) {
	if request == nil {
		request = NewStaffSignInRequest()
	}
	response = NewStaffSignInResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
	request = &DeleteGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "DeleteGroup")
	return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
	response = &DeleteGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除客服组
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
	if request == nil {
		request = NewDeleteGroupRequest()
	}
	response = NewDeleteGroupResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGroupRequest() (request *UpdateGroupRequest) {
	request = &UpdateGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "UpdateGroup")
	return
}

func NewUpdateGroupResponse() (response *UpdateGroupResponse) {
	response = &UpdateGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新客服组
func (c *Client) UpdateGroup(request *UpdateGroupRequest) (response *UpdateGroupResponse, err error) {
	if request == nil {
		request = NewUpdateGroupRequest()
	}
	response = NewUpdateGroupResponse()
	err = c.Send(request, response)
	return
}

func NewGetStaffStatusRequest() (request *GetStaffStatusRequest) {
	request = &GetStaffStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "GetStaffStatus")
	return
}

func NewGetStaffStatusResponse() (response *GetStaffStatusResponse) {
	response = &GetStaffStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取客服在线状态
func (c *Client) GetStaffStatus(request *GetStaffStatusRequest) (response *GetStaffStatusResponse, err error) {
	if request == nil {
		request = NewGetStaffStatusRequest()
	}
	response = NewGetStaffStatusResponse()
	err = c.Send(request, response)
	return
}

func NewStaffPromptToSupplyRequest() (request *StaffPromptToSupplyRequest) {
	request = &StaffPromptToSupplyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "StaffPromptToSupply")
	return
}

func NewStaffPromptToSupplyResponse() (response *StaffPromptToSupplyResponse) {
	response = &StaffPromptToSupplyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 待客户补充
func (c *Client) StaffPromptToSupply(request *StaffPromptToSupplyRequest) (response *StaffPromptToSupplyResponse, err error) {
	if request == nil {
		request = NewStaffPromptToSupplyRequest()
	}
	response = NewStaffPromptToSupplyResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCategoryRequest() (request *UpdateCategoryRequest) {
	request = &UpdateCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "UpdateCategory")
	return
}

func NewUpdateCategoryResponse() (response *UpdateCategoryResponse) {
	response = &UpdateCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新产品分类
func (c *Client) UpdateCategory(request *UpdateCategoryRequest) (response *UpdateCategoryResponse, err error) {
	if request == nil {
		request = NewUpdateCategoryRequest()
	}
	response = NewUpdateCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTicketsRequest() (request *DescribeTicketsRequest) {
	request = &DescribeTicketsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "DescribeTickets")
	return
}

func NewDescribeTicketsResponse() (response *DescribeTicketsResponse) {
	response = &DescribeTicketsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 工单列表
func (c *Client) DescribeTickets(request *DescribeTicketsRequest) (response *DescribeTicketsResponse, err error) {
	if request == nil {
		request = NewDescribeTicketsRequest()
	}
	response = NewDescribeTicketsResponse()
	err = c.Send(request, response)
	return
}

func NewGetCosParamsByBucketRequest() (request *GetCosParamsByBucketRequest) {
	request = &GetCosParamsByBucketRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "GetCosParamsByBucket")
	return
}

func NewGetCosParamsByBucketResponse() (response *GetCosParamsByBucketResponse) {
	response = &GetCosParamsByBucketResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过bucket和区域获取COS参数
func (c *Client) GetCosParamsByBucket(request *GetCosParamsByBucketRequest) (response *GetCosParamsByBucketResponse, err error) {
	if request == nil {
		request = NewGetCosParamsByBucketRequest()
	}
	response = NewGetCosParamsByBucketResponse()
	err = c.Send(request, response)
	return
}

func NewGetGroupsRequest() (request *GetGroupsRequest) {
	request = &GetGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "GetGroups")
	return
}

func NewGetGroupsResponse() (response *GetGroupsResponse) {
	response = &GetGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取客服组列表
func (c *Client) GetGroups(request *GetGroupsRequest) (response *GetGroupsResponse, err error) {
	if request == nil {
		request = NewGetGroupsRequest()
	}
	response = NewGetGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewStaffUrgeRequest() (request *StaffUrgeRequest) {
	request = &StaffUrgeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "StaffUrge")
	return
}

func NewStaffUrgeResponse() (response *StaffUrgeResponse) {
	response = &StaffUrgeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客服催单
func (c *Client) StaffUrge(request *StaffUrgeRequest) (response *StaffUrgeResponse, err error) {
	if request == nil {
		request = NewStaffUrgeRequest()
	}
	response = NewStaffUrgeResponse()
	err = c.Send(request, response)
	return
}

func NewExportTicketsRequest() (request *ExportTicketsRequest) {
	request = &ExportTicketsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "ExportTickets")
	return
}

func NewExportTicketsResponse() (response *ExportTicketsResponse) {
	response = &ExportTicketsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出工单列表
func (c *Client) ExportTickets(request *ExportTicketsRequest) (response *ExportTicketsResponse, err error) {
	if request == nil {
		request = NewExportTicketsRequest()
	}
	response = NewExportTicketsResponse()
	err = c.Send(request, response)
	return
}

func NewStaffEnterRequest() (request *StaffEnterRequest) {
	request = &StaffEnterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "StaffEnter")
	return
}

func NewStaffEnterResponse() (response *StaffEnterResponse) {
	response = &StaffEnterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客服入职
func (c *Client) StaffEnter(request *StaffEnterRequest) (response *StaffEnterResponse, err error) {
	if request == nil {
		request = NewStaffEnterRequest()
	}
	response = NewStaffEnterResponse()
	err = c.Send(request, response)
	return
}

func NewStaffAssignTicketRequest() (request *StaffAssignTicketRequest) {
	request = &StaffAssignTicketRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "StaffAssignTicket")
	return
}

func NewStaffAssignTicketResponse() (response *StaffAssignTicketResponse) {
	response = &StaffAssignTicketResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客服转单
func (c *Client) StaffAssignTicket(request *StaffAssignTicketRequest) (response *StaffAssignTicketResponse, err error) {
	if request == nil {
		request = NewStaffAssignTicketRequest()
	}
	response = NewStaffAssignTicketResponse()
	err = c.Send(request, response)
	return
}

func NewStaffReceiveTicketRequest() (request *StaffReceiveTicketRequest) {
	request = &StaffReceiveTicketRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "StaffReceiveTicket")
	return
}

func NewStaffReceiveTicketResponse() (response *StaffReceiveTicketResponse) {
	response = &StaffReceiveTicketResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客服收到工单
func (c *Client) StaffReceiveTicket(request *StaffReceiveTicketRequest) (response *StaffReceiveTicketResponse, err error) {
	if request == nil {
		request = NewStaffReceiveTicketRequest()
	}
	response = NewStaffReceiveTicketResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateScheduleRequest() (request *UpdateScheduleRequest) {
	request = &UpdateScheduleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "UpdateSchedule")
	return
}

func NewUpdateScheduleResponse() (response *UpdateScheduleResponse) {
	response = &UpdateScheduleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新排班
func (c *Client) UpdateSchedule(request *UpdateScheduleRequest) (response *UpdateScheduleResponse, err error) {
	if request == nil {
		request = NewUpdateScheduleRequest()
	}
	response = NewUpdateScheduleResponse()
	err = c.Send(request, response)
	return
}

func NewGetCategoryRequest() (request *GetCategoryRequest) {
	request = &GetCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "GetCategory")
	return
}

func NewGetCategoryResponse() (response *GetCategoryResponse) {
	response = &GetCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品分类
func (c *Client) GetCategory(request *GetCategoryRequest) (response *GetCategoryResponse, err error) {
	if request == nil {
		request = NewGetCategoryRequest()
	}
	response = NewGetCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewGetStaffByUserIdRequest() (request *GetStaffByUserIdRequest) {
	request = &GetStaffByUserIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "GetStaffByUserId")
	return
}

func NewGetStaffByUserIdResponse() (response *GetStaffByUserIdResponse) {
	response = &GetStaffByUserIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 登录名UserId即StaffId，用于从全量用户中搜索出需要的用户，以添加为客服
func (c *Client) GetStaffByUserId(request *GetStaffByUserIdRequest) (response *GetStaffByUserIdResponse, err error) {
	if request == nil {
		request = NewGetStaffByUserIdRequest()
	}
	response = NewGetStaffByUserIdResponse()
	err = c.Send(request, response)
	return
}

func NewAddGroupStaffRequest() (request *AddGroupStaffRequest) {
	request = &AddGroupStaffRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "AddGroupStaff")
	return
}

func NewAddGroupStaffResponse() (response *AddGroupStaffResponse) {
	response = &AddGroupStaffResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加客服组成员
func (c *Client) AddGroupStaff(request *AddGroupStaffRequest) (response *AddGroupStaffResponse, err error) {
	if request == nil {
		request = NewAddGroupStaffRequest()
	}
	response = NewAddGroupStaffResponse()
	err = c.Send(request, response)
	return
}

func NewGetFieldRequest() (request *GetFieldRequest) {
	request = &GetFieldRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "GetField")
	return
}

func NewGetFieldResponse() (response *GetFieldResponse) {
	response = &GetFieldResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询表单字段
func (c *Client) GetField(request *GetFieldRequest) (response *GetFieldResponse, err error) {
	if request == nil {
		request = NewGetFieldRequest()
	}
	response = NewGetFieldResponse()
	err = c.Send(request, response)
	return
}

func NewGetTicketCategoryRequest() (request *GetTicketCategoryRequest) {
	request = &GetTicketCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "GetTicketCategory")
	return
}

func NewGetTicketCategoryResponse() (response *GetTicketCategoryResponse) {
	response = &GetTicketCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取工单分类
func (c *Client) GetTicketCategory(request *GetTicketCategoryRequest) (response *GetTicketCategoryResponse, err error) {
	if request == nil {
		request = NewGetTicketCategoryRequest()
	}
	response = NewGetTicketCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGroupRequest() (request *CreateGroupRequest) {
	request = &CreateGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "CreateGroup")
	return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
	response = &CreateGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建客服组
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
	if request == nil {
		request = NewCreateGroupRequest()
	}
	response = NewCreateGroupResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTicketRequest() (request *CreateTicketRequest) {
	request = &CreateTicketRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "CreateTicket")
	return
}

func NewCreateTicketResponse() (response *CreateTicketResponse) {
	response = &CreateTicketResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建工单
func (c *Client) CreateTicket(request *CreateTicketRequest) (response *CreateTicketResponse, err error) {
	if request == nil {
		request = NewCreateTicketRequest()
	}
	response = NewCreateTicketResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteFieldRequest() (request *DeleteFieldRequest) {
	request = &DeleteFieldRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "DeleteField")
	return
}

func NewDeleteFieldResponse() (response *DeleteFieldResponse) {
	response = &DeleteFieldResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除表单字段
func (c *Client) DeleteField(request *DeleteFieldRequest) (response *DeleteFieldResponse, err error) {
	if request == nil {
		request = NewDeleteFieldRequest()
	}
	response = NewDeleteFieldResponse()
	err = c.Send(request, response)
	return
}

func NewGetStaffSelfInfoRequest() (request *GetStaffSelfInfoRequest) {
	request = &GetStaffSelfInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "GetStaffSelfInfo")
	return
}

func NewGetStaffSelfInfoResponse() (response *GetStaffSelfInfoResponse) {
	response = &GetStaffSelfInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取客服个人信息
func (c *Client) GetStaffSelfInfo(request *GetStaffSelfInfoRequest) (response *GetStaffSelfInfoResponse, err error) {
	if request == nil {
		request = NewGetStaffSelfInfoRequest()
	}
	response = NewGetStaffSelfInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGetTicketCategoryByLevelRequest() (request *GetTicketCategoryByLevelRequest) {
	request = &GetTicketCategoryByLevelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "GetTicketCategoryByLevel")
	return
}

func NewGetTicketCategoryByLevelResponse() (response *GetTicketCategoryByLevelResponse) {
	response = &GetTicketCategoryByLevelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过分类Id获取三级分类和对应的表单字段
func (c *Client) GetTicketCategoryByLevel(request *GetTicketCategoryByLevelRequest) (response *GetTicketCategoryByLevelResponse, err error) {
	if request == nil {
		request = NewGetTicketCategoryByLevelRequest()
	}
	response = NewGetTicketCategoryByLevelResponse()
	err = c.Send(request, response)
	return
}

func NewCheckBeforeDeleteStaffStatusRequest() (request *CheckBeforeDeleteStaffStatusRequest) {
	request = &CheckBeforeDeleteStaffStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "CheckBeforeDeleteStaffStatus")
	return
}

func NewCheckBeforeDeleteStaffStatusResponse() (response *CheckBeforeDeleteStaffStatusResponse) {
	response = &CheckBeforeDeleteStaffStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 签退客服前检查客服未处理工单数量
func (c *Client) CheckBeforeDeleteStaffStatus(request *CheckBeforeDeleteStaffStatusRequest) (response *CheckBeforeDeleteStaffStatusResponse, err error) {
	if request == nil {
		request = NewCheckBeforeDeleteStaffStatusRequest()
	}
	response = NewCheckBeforeDeleteStaffStatusResponse()
	err = c.Send(request, response)
	return
}

func NewGetCosBucketRequest() (request *GetCosBucketRequest) {
	request = &GetCosBucketRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "GetCosBucket")
	return
}

func NewGetCosBucketResponse() (response *GetCosBucketResponse) {
	response = &GetCosBucketResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetCosBucket
func (c *Client) GetCosBucket(request *GetCosBucketRequest) (response *GetCosBucketResponse, err error) {
	if request == nil {
		request = NewGetCosBucketRequest()
	}
	response = NewGetCosBucketResponse()
	err = c.Send(request, response)
	return
}

func NewStaffDismissRequest() (request *StaffDismissRequest) {
	request = &StaffDismissRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "StaffDismiss")
	return
}

func NewStaffDismissResponse() (response *StaffDismissResponse) {
	response = &StaffDismissResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客服离职
func (c *Client) StaffDismiss(request *StaffDismissRequest) (response *StaffDismissResponse, err error) {
	if request == nil {
		request = NewStaffDismissRequest()
	}
	response = NewStaffDismissResponse()
	err = c.Send(request, response)
	return
}

func NewGetWorkbenchDataRequest() (request *GetWorkbenchDataRequest) {
	request = &GetWorkbenchDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "GetWorkbenchData")
	return
}

func NewGetWorkbenchDataResponse() (response *GetWorkbenchDataResponse) {
	response = &GetWorkbenchDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取运营端数据
func (c *Client) GetWorkbenchData(request *GetWorkbenchDataRequest) (response *GetWorkbenchDataResponse, err error) {
	if request == nil {
		request = NewGetWorkbenchDataRequest()
	}
	response = NewGetWorkbenchDataResponse()
	err = c.Send(request, response)
	return
}

func NewGetScheduleRequest() (request *GetScheduleRequest) {
	request = &GetScheduleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "GetSchedule")
	return
}

func NewGetScheduleResponse() (response *GetScheduleResponse) {
	response = &GetScheduleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取排班
func (c *Client) GetSchedule(request *GetScheduleRequest) (response *GetScheduleResponse, err error) {
	if request == nil {
		request = NewGetScheduleRequest()
	}
	response = NewGetScheduleResponse()
	err = c.Send(request, response)
	return
}

func NewGetServerTimeRequest() (request *GetServerTimeRequest) {
	request = &GetServerTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opticket", APIVersion, "GetServerTime")
	return
}

func NewGetServerTimeResponse() (response *GetServerTimeResponse) {
	response = &GetServerTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取服务器时间
func (c *Client) GetServerTime(request *GetServerTimeRequest) (response *GetServerTimeResponse, err error) {
	if request == nil {
		request = NewGetServerTimeRequest()
	}
	response = NewGetServerTimeResponse()
	err = c.Send(request, response)
	return
}
