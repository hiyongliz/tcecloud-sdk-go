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

package v20200726

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-07-26"

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

func NewUpdateConsoleMessageTemplateRequest() (request *UpdateConsoleMessageTemplateRequest) {
	request = &UpdateConsoleMessageTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "UpdateConsoleMessageTemplate")
	return
}

func NewUpdateConsoleMessageTemplateResponse() (response *UpdateConsoleMessageTemplateResponse) {
	response = &UpdateConsoleMessageTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新实时消息模板
func (c *Client) UpdateConsoleMessageTemplate(request *UpdateConsoleMessageTemplateRequest) (response *UpdateConsoleMessageTemplateResponse, err error) {
	if request == nil {
		request = NewUpdateConsoleMessageTemplateRequest()
	}
	response = NewUpdateConsoleMessageTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewCancelThemeRequest() (request *CancelThemeRequest) {
	request = &CancelThemeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "CancelTheme")
	return
}

func NewCancelThemeResponse() (response *CancelThemeResponse) {
	response = &CancelThemeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 消息模板撤销审核
func (c *Client) CancelTheme(request *CancelThemeRequest) (response *CancelThemeResponse, err error) {
	if request == nil {
		request = NewCancelThemeRequest()
	}
	response = NewCancelThemeResponse()
	err = c.Send(request, response)
	return
}

func NewSubmitThemeRequest() (request *SubmitThemeRequest) {
	request = &SubmitThemeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "SubmitTheme")
	return
}

func NewSubmitThemeResponse() (response *SubmitThemeResponse) {
	response = &SubmitThemeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增消息模板
func (c *Client) SubmitTheme(request *SubmitThemeRequest) (response *SubmitThemeResponse, err error) {
	if request == nil {
		request = NewSubmitThemeRequest()
	}
	response = NewSubmitThemeResponse()
	err = c.Send(request, response)
	return
}

func NewGetThemeDetailRequest() (request *GetThemeDetailRequest) {
	request = &GetThemeDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "GetThemeDetail")
	return
}

func NewGetThemeDetailResponse() (response *GetThemeDetailResponse) {
	response = &GetThemeDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询消息模板详情
func (c *Client) GetThemeDetail(request *GetThemeDetailRequest) (response *GetThemeDetailResponse, err error) {
	if request == nil {
		request = NewGetThemeDetailRequest()
	}
	response = NewGetThemeDetailResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateThemeAdminRequest() (request *UpdateThemeAdminRequest) {
	request = &UpdateThemeAdminRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "UpdateThemeAdmin")
	return
}

func NewUpdateThemeAdminResponse() (response *UpdateThemeAdminResponse) {
	response = &UpdateThemeAdminResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新消息模板负责人
func (c *Client) UpdateThemeAdmin(request *UpdateThemeAdminRequest) (response *UpdateThemeAdminResponse, err error) {
	if request == nil {
		request = NewUpdateThemeAdminRequest()
	}
	response = NewUpdateThemeAdminResponse()
	err = c.Send(request, response)
	return
}

func NewAddMsgTypeRequest() (request *AddMsgTypeRequest) {
	request = &AddMsgTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "AddMsgType")
	return
}

func NewAddMsgTypeResponse() (response *AddMsgTypeResponse) {
	response = &AddMsgTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增二级消息分类
func (c *Client) AddMsgType(request *AddMsgTypeRequest) (response *AddMsgTypeResponse, err error) {
	if request == nil {
		request = NewAddMsgTypeRequest()
	}
	response = NewAddMsgTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteThemeRequest() (request *DeleteThemeRequest) {
	request = &DeleteThemeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "DeleteTheme")
	return
}

func NewDeleteThemeResponse() (response *DeleteThemeResponse) {
	response = &DeleteThemeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除消息模板
func (c *Client) DeleteTheme(request *DeleteThemeRequest) (response *DeleteThemeResponse, err error) {
	if request == nil {
		request = NewDeleteThemeRequest()
	}
	response = NewDeleteThemeResponse()
	err = c.Send(request, response)
	return
}

func NewQueryMsgTypeListRequest() (request *QueryMsgTypeListRequest) {
	request = &QueryMsgTypeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "QueryMsgTypeList")
	return
}

func NewQueryMsgTypeListResponse() (response *QueryMsgTypeListResponse) {
	response = &QueryMsgTypeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列表查询消息类型
func (c *Client) QueryMsgTypeList(request *QueryMsgTypeListRequest) (response *QueryMsgTypeListResponse, err error) {
	if request == nil {
		request = NewQueryMsgTypeListRequest()
	}
	response = NewQueryMsgTypeListResponse()
	err = c.Send(request, response)
	return
}

func NewSendTestMessageThemeRequest() (request *SendTestMessageThemeRequest) {
	request = &SendTestMessageThemeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "SendTestMessageTheme")
	return
}

func NewSendTestMessageThemeResponse() (response *SendTestMessageThemeResponse) {
	response = &SendTestMessageThemeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将模板直接试发送
func (c *Client) SendTestMessageTheme(request *SendTestMessageThemeRequest) (response *SendTestMessageThemeResponse, err error) {
	if request == nil {
		request = NewSendTestMessageThemeRequest()
	}
	response = NewSendTestMessageThemeResponse()
	err = c.Send(request, response)
	return
}

func NewSubmitAuditRequest() (request *SubmitAuditRequest) {
	request = &SubmitAuditRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "SubmitAudit")
	return
}

func NewSubmitAuditResponse() (response *SubmitAuditResponse) {
	response = &SubmitAuditResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提交消息模板审核
func (c *Client) SubmitAudit(request *SubmitAuditRequest) (response *SubmitAuditResponse, err error) {
	if request == nil {
		request = NewSubmitAuditRequest()
	}
	response = NewSubmitAuditResponse()
	err = c.Send(request, response)
	return
}

func NewAddConsoleMessageTemplateRequest() (request *AddConsoleMessageTemplateRequest) {
	request = &AddConsoleMessageTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "AddConsoleMessageTemplate")
	return
}

func NewAddConsoleMessageTemplateResponse() (response *AddConsoleMessageTemplateResponse) {
	response = &AddConsoleMessageTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加消息实时推送模板
func (c *Client) AddConsoleMessageTemplate(request *AddConsoleMessageTemplateRequest) (response *AddConsoleMessageTemplateResponse, err error) {
	if request == nil {
		request = NewAddConsoleMessageTemplateRequest()
	}
	response = NewAddConsoleMessageTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewPublishThemeRequest() (request *PublishThemeRequest) {
	request = &PublishThemeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "PublishTheme")
	return
}

func NewPublishThemeResponse() (response *PublishThemeResponse) {
	response = &PublishThemeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 消息模板通过审核
func (c *Client) PublishTheme(request *PublishThemeRequest) (response *PublishThemeResponse, err error) {
	if request == nil {
		request = NewPublishThemeRequest()
	}
	response = NewPublishThemeResponse()
	err = c.Send(request, response)
	return
}

func NewQueryThemeListRequest() (request *QueryThemeListRequest) {
	request = &QueryThemeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "QueryThemeList")
	return
}

func NewQueryThemeListResponse() (response *QueryThemeListResponse) {
	response = &QueryThemeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列表分页查询过滤消息模板
func (c *Client) QueryThemeList(request *QueryThemeListRequest) (response *QueryThemeListResponse, err error) {
	if request == nil {
		request = NewQueryThemeListRequest()
	}
	response = NewQueryThemeListResponse()
	err = c.Send(request, response)
	return
}

func NewSendConsoleUserMessageRequest() (request *SendConsoleUserMessageRequest) {
	request = &SendConsoleUserMessageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "SendConsoleUserMessage")
	return
}

func NewSendConsoleUserMessageResponse() (response *SendConsoleUserMessageResponse) {
	response = &SendConsoleUserMessageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 消息实时推送接口
func (c *Client) SendConsoleUserMessage(request *SendConsoleUserMessageRequest) (response *SendConsoleUserMessageResponse, err error) {
	if request == nil {
		request = NewSendConsoleUserMessageRequest()
	}
	response = NewSendConsoleUserMessageResponse()
	err = c.Send(request, response)
	return
}

func NewGetMsgCategorysRequest() (request *GetMsgCategorysRequest) {
	request = &GetMsgCategorysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "GetMsgCategorys")
	return
}

func NewGetMsgCategorysResponse() (response *GetMsgCategorysResponse) {
	response = &GetMsgCategorysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有的一级消息分类
func (c *Client) GetMsgCategorys(request *GetMsgCategorysRequest) (response *GetMsgCategorysResponse, err error) {
	if request == nil {
		request = NewGetMsgCategorysRequest()
	}
	response = NewGetMsgCategorysResponse()
	err = c.Send(request, response)
	return
}

func NewQueryConsoleMessageTemplateRequest() (request *QueryConsoleMessageTemplateRequest) {
	request = &QueryConsoleMessageTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "QueryConsoleMessageTemplate")
	return
}

func NewQueryConsoleMessageTemplateResponse() (response *QueryConsoleMessageTemplateResponse) {
	response = &QueryConsoleMessageTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列表查询消息实时发送模板
func (c *Client) QueryConsoleMessageTemplate(request *QueryConsoleMessageTemplateRequest) (response *QueryConsoleMessageTemplateResponse, err error) {
	if request == nil {
		request = NewQueryConsoleMessageTemplateRequest()
	}
	response = NewQueryConsoleMessageTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConsoleMessageTemplateRequest() (request *DeleteConsoleMessageTemplateRequest) {
	request = &DeleteConsoleMessageTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "DeleteConsoleMessageTemplate")
	return
}

func NewDeleteConsoleMessageTemplateResponse() (response *DeleteConsoleMessageTemplateResponse) {
	response = &DeleteConsoleMessageTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除消息实时模板
func (c *Client) DeleteConsoleMessageTemplate(request *DeleteConsoleMessageTemplateRequest) (response *DeleteConsoleMessageTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteConsoleMessageTemplateRequest()
	}
	response = NewDeleteConsoleMessageTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewEditMsgTypeRequest() (request *EditMsgTypeRequest) {
	request = &EditMsgTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "EditMsgType")
	return
}

func NewEditMsgTypeResponse() (response *EditMsgTypeResponse) {
	response = &EditMsgTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新消息类型
func (c *Client) EditMsgType(request *EditMsgTypeRequest) (response *EditMsgTypeResponse, err error) {
	if request == nil {
		request = NewEditMsgTypeRequest()
	}
	response = NewEditMsgTypeResponse()
	err = c.Send(request, response)
	return
}

func NewRefuseApplyRequest() (request *RefuseApplyRequest) {
	request = &RefuseApplyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "RefuseApply")
	return
}

func NewRefuseApplyResponse() (response *RefuseApplyResponse) {
	response = &RefuseApplyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 驳回审核
func (c *Client) RefuseApply(request *RefuseApplyRequest) (response *RefuseApplyResponse, err error) {
	if request == nil {
		request = NewRefuseApplyRequest()
	}
	response = NewRefuseApplyResponse()
	err = c.Send(request, response)
	return
}

func NewQueryAuditListRequest() (request *QueryAuditListRequest) {
	request = &QueryAuditListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "QueryAuditList")
	return
}

func NewQueryAuditListResponse() (response *QueryAuditListResponse) {
	response = &QueryAuditListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 消息审核列表
func (c *Client) QueryAuditList(request *QueryAuditListRequest) (response *QueryAuditListResponse, err error) {
	if request == nil {
		request = NewQueryAuditListRequest()
	}
	response = NewQueryAuditListResponse()
	err = c.Send(request, response)
	return
}

func NewSaveThemeRequest() (request *SaveThemeRequest) {
	request = &SaveThemeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcloudmsg", APIVersion, "SaveTheme")
	return
}

func NewSaveThemeResponse() (response *SaveThemeResponse) {
	response = &SaveThemeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑消息模板
func (c *Client) SaveTheme(request *SaveThemeRequest) (response *SaveThemeResponse, err error) {
	if request == nil {
		request = NewSaveThemeRequest()
	}
	response = NewSaveThemeResponse()
	err = c.Send(request, response)
	return
}
