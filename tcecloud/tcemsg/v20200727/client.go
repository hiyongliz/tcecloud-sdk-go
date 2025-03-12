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

package v20200727

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-07-27"

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

func NewQueryUserCmgtSiteMsgRequest() (request *QueryUserCmgtSiteMsgRequest) {
	request = &QueryUserCmgtSiteMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "QueryUserCmgtSiteMsg")
	return
}

func NewQueryUserCmgtSiteMsgResponse() (response *QueryUserCmgtSiteMsgResponse) {
	response = &QueryUserCmgtSiteMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 只能查询自己的站内信
func (c *Client) QueryUserCmgtSiteMsg(request *QueryUserCmgtSiteMsgRequest) (response *QueryUserCmgtSiteMsgResponse, err error) {
	if request == nil {
		request = NewQueryUserCmgtSiteMsgRequest()
	}
	response = NewQueryUserCmgtSiteMsgResponse()
	err = c.Send(request, response)
	return
}

func NewDelUserCmgtSiteMsgRequest() (request *DelUserCmgtSiteMsgRequest) {
	request = &DelUserCmgtSiteMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "DelUserCmgtSiteMsg")
	return
}

func NewDelUserCmgtSiteMsgResponse() (response *DelUserCmgtSiteMsgResponse) {
	response = &DelUserCmgtSiteMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 只允许操作自己的站内信
func (c *Client) DelUserCmgtSiteMsg(request *DelUserCmgtSiteMsgRequest) (response *DelUserCmgtSiteMsgResponse, err error) {
	if request == nil {
		request = NewDelUserCmgtSiteMsgRequest()
	}
	response = NewDelUserCmgtSiteMsgResponse()
	err = c.Send(request, response)
	return
}

func NewGetReserveSendChannelNumberRequest() (request *GetReserveSendChannelNumberRequest) {
	request = &GetReserveSendChannelNumberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "GetReserveSendChannelNumber")
	return
}

func NewGetReserveSendChannelNumberResponse() (response *GetReserveSendChannelNumberResponse) {
	response = &GetReserveSendChannelNumberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发送通道查询可用第三方通道编号
func (c *Client) GetReserveSendChannelNumber(request *GetReserveSendChannelNumberRequest) (response *GetReserveSendChannelNumberResponse, err error) {
	if request == nil {
		request = NewGetReserveSendChannelNumberRequest()
	}
	response = NewGetReserveSendChannelNumberResponse()
	err = c.Send(request, response)
	return
}

func NewGetMsgRequest() (request *GetMsgRequest) {
	request = &GetMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "GetMsg")
	return
}

func NewGetMsgResponse() (response *GetMsgResponse) {
	response = &GetMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端消息流水详情
func (c *Client) GetMsg(request *GetMsgRequest) (response *GetMsgResponse, err error) {
	if request == nil {
		request = NewGetMsgRequest()
	}
	response = NewGetMsgResponse()
	err = c.Send(request, response)
	return
}

func NewQueryPostRequest() (request *QueryPostRequest) {
	request = &QueryPostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "QueryPost")
	return
}

func NewQueryPostResponse() (response *QueryPostResponse) {
	response = &QueryPostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端公告查询
func (c *Client) QueryPost(request *QueryPostRequest) (response *QueryPostResponse, err error) {
	if request == nil {
		request = NewQueryPostRequest()
	}
	response = NewQueryPostResponse()
	err = c.Send(request, response)
	return
}

func NewSubmitThemeRequest() (request *SubmitThemeRequest) {
	request = &SubmitThemeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "SubmitTheme")
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

func NewSaveThemeRequest() (request *SaveThemeRequest) {
	request = &SaveThemeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "SaveTheme")
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

func NewUpdatePostRequest() (request *UpdatePostRequest) {
	request = &UpdatePostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "UpdatePost")
	return
}

func NewUpdatePostResponse() (response *UpdatePostResponse) {
	response = &UpdatePostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端公告更新
func (c *Client) UpdatePost(request *UpdatePostRequest) (response *UpdatePostResponse, err error) {
	if request == nil {
		request = NewUpdatePostRequest()
	}
	response = NewUpdatePostResponse()
	err = c.Send(request, response)
	return
}

func NewQuerySendChannelRequest() (request *QuerySendChannelRequest) {
	request = &QuerySendChannelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "QuerySendChannel")
	return
}

func NewQuerySendChannelResponse() (response *QuerySendChannelResponse) {
	response = &QuerySendChannelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发送通道列表查询
func (c *Client) QuerySendChannel(request *QuerySendChannelRequest) (response *QuerySendChannelResponse, err error) {
	if request == nil {
		request = NewQuerySendChannelRequest()
	}
	response = NewQuerySendChannelResponse()
	err = c.Send(request, response)
	return
}

func NewResendMsgRequest() (request *ResendMsgRequest) {
	request = &ResendMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "ResendMsg")
	return
}

func NewResendMsgResponse() (response *ResendMsgResponse) {
	response = &ResendMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重发消息
func (c *Client) ResendMsg(request *ResendMsgRequest) (response *ResendMsgResponse, err error) {
	if request == nil {
		request = NewResendMsgRequest()
	}
	response = NewResendMsgResponse()
	err = c.Send(request, response)
	return
}

func NewQueryMsgRequest() (request *QueryMsgRequest) {
	request = &QueryMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "QueryMsg")
	return
}

func NewQueryMsgResponse() (response *QueryMsgResponse) {
	response = &QueryMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端消息流水查询
func (c *Client) QueryMsg(request *QueryMsgRequest) (response *QueryMsgResponse, err error) {
	if request == nil {
		request = NewQueryMsgRequest()
	}
	response = NewQueryMsgResponse()
	err = c.Send(request, response)
	return
}

func NewGetThemeDetailRequest() (request *GetThemeDetailRequest) {
	request = &GetThemeDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "GetThemeDetail")
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

func NewUpdateSendChannelRequest() (request *UpdateSendChannelRequest) {
	request = &UpdateSendChannelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "UpdateSendChannel")
	return
}

func NewUpdateSendChannelResponse() (response *UpdateSendChannelResponse) {
	response = &UpdateSendChannelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新发送通道配置
func (c *Client) UpdateSendChannel(request *UpdateSendChannelRequest) (response *UpdateSendChannelResponse, err error) {
	if request == nil {
		request = NewUpdateSendChannelRequest()
	}
	response = NewUpdateSendChannelResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateThemeAdminRequest() (request *UpdateThemeAdminRequest) {
	request = &UpdateThemeAdminRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "UpdateThemeAdmin")
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

func NewPublishThemeRequest() (request *PublishThemeRequest) {
	request = &PublishThemeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "PublishTheme")
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

func NewDeleteSendChannelRequest() (request *DeleteSendChannelRequest) {
	request = &DeleteSendChannelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "DeleteSendChannel")
	return
}

func NewDeleteSendChannelResponse() (response *DeleteSendChannelResponse) {
	response = &DeleteSendChannelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 仅支持删除第三方消息通道
func (c *Client) DeleteSendChannel(request *DeleteSendChannelRequest) (response *DeleteSendChannelResponse, err error) {
	if request == nil {
		request = NewDeleteSendChannelRequest()
	}
	response = NewDeleteSendChannelResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserCmgtSiteMsgRequest() (request *GetUserCmgtSiteMsgRequest) {
	request = &GetUserCmgtSiteMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "GetUserCmgtSiteMsg")
	return
}

func NewGetUserCmgtSiteMsgResponse() (response *GetUserCmgtSiteMsgResponse) {
	response = &GetUserCmgtSiteMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 仅允许查看自己的站内信
func (c *Client) GetUserCmgtSiteMsg(request *GetUserCmgtSiteMsgRequest) (response *GetUserCmgtSiteMsgResponse, err error) {
	if request == nil {
		request = NewGetUserCmgtSiteMsgRequest()
	}
	response = NewGetUserCmgtSiteMsgResponse()
	err = c.Send(request, response)
	return
}

func NewAddTopicRequest() (request *AddTopicRequest) {
	request = &AddTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "AddTopic")
	return
}

func NewAddTopicResponse() (response *AddTopicResponse) {
	response = &AddTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加订阅模版
func (c *Client) AddTopic(request *AddTopicRequest) (response *AddTopicResponse, err error) {
	if request == nil {
		request = NewAddTopicRequest()
	}
	response = NewAddTopicResponse()
	err = c.Send(request, response)
	return
}

func NewReadUserCmgtSiteMsgRequest() (request *ReadUserCmgtSiteMsgRequest) {
	request = &ReadUserCmgtSiteMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "ReadUserCmgtSiteMsg")
	return
}

func NewReadUserCmgtSiteMsgResponse() (response *ReadUserCmgtSiteMsgResponse) {
	response = &ReadUserCmgtSiteMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 只允许操作自己的站内信
func (c *Client) ReadUserCmgtSiteMsg(request *ReadUserCmgtSiteMsgRequest) (response *ReadUserCmgtSiteMsgResponse, err error) {
	if request == nil {
		request = NewReadUserCmgtSiteMsgRequest()
	}
	response = NewReadUserCmgtSiteMsgResponse()
	err = c.Send(request, response)
	return
}

func NewDelPostRequest() (request *DelPostRequest) {
	request = &DelPostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "DelPost")
	return
}

func NewDelPostResponse() (response *DelPostResponse) {
	response = &DelPostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端公告删除
func (c *Client) DelPost(request *DelPostRequest) (response *DelPostResponse, err error) {
	if request == nil {
		request = NewDelPostRequest()
	}
	response = NewDelPostResponse()
	err = c.Send(request, response)
	return
}

func NewQueryAuditListRequest() (request *QueryAuditListRequest) {
	request = &QueryAuditListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "QueryAuditList")
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

func NewAddPostRequest() (request *AddPostRequest) {
	request = &AddPostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "AddPost")
	return
}

func NewAddPostResponse() (response *AddPostResponse) {
	response = &AddPostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端公告发表
func (c *Client) AddPost(request *AddPostRequest) (response *AddPostResponse, err error) {
	if request == nil {
		request = NewAddPostRequest()
	}
	response = NewAddPostResponse()
	err = c.Send(request, response)
	return
}

func NewGetTopicRequest() (request *GetTopicRequest) {
	request = &GetTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "GetTopic")
	return
}

func NewGetTopicResponse() (response *GetTopicResponse) {
	response = &GetTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 订阅模版详情
func (c *Client) GetTopic(request *GetTopicRequest) (response *GetTopicResponse, err error) {
	if request == nil {
		request = NewGetTopicRequest()
	}
	response = NewGetTopicResponse()
	err = c.Send(request, response)
	return
}

func NewRefuseApplyRequest() (request *RefuseApplyRequest) {
	request = &RefuseApplyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "RefuseApply")
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

func NewUpdateTopicRequest() (request *UpdateTopicRequest) {
	request = &UpdateTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "UpdateTopic")
	return
}

func NewUpdateTopicResponse() (response *UpdateTopicResponse) {
	response = &UpdateTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 订阅模版更新
func (c *Client) UpdateTopic(request *UpdateTopicRequest) (response *UpdateTopicResponse, err error) {
	if request == nil {
		request = NewUpdateTopicRequest()
	}
	response = NewUpdateTopicResponse()
	err = c.Send(request, response)
	return
}

func NewAddSendChannelRequest() (request *AddSendChannelRequest) {
	request = &AddSendChannelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "AddSendChannel")
	return
}

func NewAddSendChannelResponse() (response *AddSendChannelResponse) {
	response = &AddSendChannelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加第三方消息通道配置
func (c *Client) AddSendChannel(request *AddSendChannelRequest) (response *AddSendChannelResponse, err error) {
	if request == nil {
		request = NewAddSendChannelRequest()
	}
	response = NewAddSendChannelResponse()
	err = c.Send(request, response)
	return
}

func NewQueryThemeListRequest() (request *QueryThemeListRequest) {
	request = &QueryThemeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "QueryThemeList")
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

func NewDeleteThemeRequest() (request *DeleteThemeRequest) {
	request = &DeleteThemeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "DeleteTheme")
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

func NewDelTopicRequest() (request *DelTopicRequest) {
	request = &DelTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "DelTopic")
	return
}

func NewDelTopicResponse() (response *DelTopicResponse) {
	response = &DelTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 订阅模版删除
func (c *Client) DelTopic(request *DelTopicRequest) (response *DelTopicResponse, err error) {
	if request == nil {
		request = NewDelTopicRequest()
	}
	response = NewDelTopicResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchSendChannelRequest() (request *SwitchSendChannelRequest) {
	request = &SwitchSendChannelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "SwitchSendChannel")
	return
}

func NewSwitchSendChannelResponse() (response *SwitchSendChannelResponse) {
	response = &SwitchSendChannelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开关消息通道
func (c *Client) SwitchSendChannel(request *SwitchSendChannelRequest) (response *SwitchSendChannelResponse, err error) {
	if request == nil {
		request = NewSwitchSendChannelRequest()
	}
	response = NewSwitchSendChannelResponse()
	err = c.Send(request, response)
	return
}

func NewSubmitAuditRequest() (request *SubmitAuditRequest) {
	request = &SubmitAuditRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "SubmitAudit")
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
