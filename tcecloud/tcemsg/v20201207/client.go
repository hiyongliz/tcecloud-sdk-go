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

package v20201207

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-12-07"

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

func NewQueryDefineSmsChannelRequest() (request *QueryDefineSmsChannelRequest) {
	request = &QueryDefineSmsChannelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "QueryDefineSmsChannel")
	return
}

func NewQueryDefineSmsChannelResponse() (response *QueryDefineSmsChannelResponse) {
	response = &QueryDefineSmsChannelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列表产品自定义短信通道号
func (c *Client) QueryDefineSmsChannel(request *QueryDefineSmsChannelRequest) (response *QueryDefineSmsChannelResponse, err error) {
	if request == nil {
		request = NewQueryDefineSmsChannelRequest()
	}
	response = NewQueryDefineSmsChannelResponse()
	err = c.Send(request, response)
	return
}

func NewQuerySendChannelDetailRequest() (request *QuerySendChannelDetailRequest) {
	request = &QuerySendChannelDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "QuerySendChannelDetail")
	return
}

func NewQuerySendChannelDetailResponse() (response *QuerySendChannelDetailResponse) {
	response = &QuerySendChannelDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询发送渠道明细数据
func (c *Client) QuerySendChannelDetail(request *QuerySendChannelDetailRequest) (response *QuerySendChannelDetailResponse, err error) {
	if request == nil {
		request = NewQuerySendChannelDetailRequest()
	}
	response = NewQuerySendChannelDetailResponse()
	err = c.Send(request, response)
	return
}

func NewTestEmailRequest() (request *TestEmailRequest) {
	request = &TestEmailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "TestEmail")
	return
}

func NewTestEmailResponse() (response *TestEmailResponse) {
	response = &TestEmailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 测试邮件发送
func (c *Client) TestEmail(request *TestEmailRequest) (response *TestEmailResponse, err error) {
	if request == nil {
		request = NewTestEmailRequest()
	}
	response = NewTestEmailResponse()
	err = c.Send(request, response)
	return
}

func NewQueryMsgQuotaRequest() (request *QueryMsgQuotaRequest) {
	request = &QueryMsgQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "QueryMsgQuota")
	return
}

func NewQueryMsgQuotaResponse() (response *QueryMsgQuotaResponse) {
	response = &QueryMsgQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列表消息配额
func (c *Client) QueryMsgQuota(request *QueryMsgQuotaRequest) (response *QueryMsgQuotaResponse, err error) {
	if request == nil {
		request = NewQueryMsgQuotaRequest()
	}
	response = NewQueryMsgQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewTestSMSRequest() (request *TestSMSRequest) {
	request = &TestSMSRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "TestSMS")
	return
}

func NewTestSMSResponse() (response *TestSMSResponse) {
	response = &TestSMSResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 测试短信发送
func (c *Client) TestSMS(request *TestSMSRequest) (response *TestSMSResponse, err error) {
	if request == nil {
		request = NewTestSMSRequest()
	}
	response = NewTestSMSResponse()
	err = c.Send(request, response)
	return
}

func NewTestVoiceRequest() (request *TestVoiceRequest) {
	request = &TestVoiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "TestVoice")
	return
}

func NewTestVoiceResponse() (response *TestVoiceResponse) {
	response = &TestVoiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 测试语音发送
func (c *Client) TestVoice(request *TestVoiceRequest) (response *TestVoiceResponse, err error) {
	if request == nil {
		request = NewTestVoiceRequest()
	}
	response = NewTestVoiceResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateDefineSmsChannelRequest() (request *UpdateDefineSmsChannelRequest) {
	request = &UpdateDefineSmsChannelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "UpdateDefineSmsChannel")
	return
}

func NewUpdateDefineSmsChannelResponse() (response *UpdateDefineSmsChannelResponse) {
	response = &UpdateDefineSmsChannelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置自定义短信通道号
func (c *Client) UpdateDefineSmsChannel(request *UpdateDefineSmsChannelRequest) (response *UpdateDefineSmsChannelResponse, err error) {
	if request == nil {
		request = NewUpdateDefineSmsChannelRequest()
	}
	response = NewUpdateDefineSmsChannelResponse()
	err = c.Send(request, response)
	return
}

func NewSaveSendChannelRequest() (request *SaveSendChannelRequest) {
	request = &SaveSendChannelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "SaveSendChannel")
	return
}

func NewSaveSendChannelResponse() (response *SaveSendChannelResponse) {
	response = &SaveSendChannelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存消息发送 渠道配置
func (c *Client) SaveSendChannel(request *SaveSendChannelRequest) (response *SaveSendChannelResponse, err error) {
	if request == nil {
		request = NewSaveSendChannelRequest()
	}
	response = NewSaveSendChannelResponse()
	err = c.Send(request, response)
	return
}

func NewTestThirdPartyRequest() (request *TestThirdPartyRequest) {
	request = &TestThirdPartyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "TestThirdParty")
	return
}

func NewTestThirdPartyResponse() (response *TestThirdPartyResponse) {
	response = &TestThirdPartyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 第三方渠道发送测试
func (c *Client) TestThirdParty(request *TestThirdPartyRequest) (response *TestThirdPartyResponse, err error) {
	if request == nil {
		request = NewTestThirdPartyRequest()
	}
	response = NewTestThirdPartyResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMsgQuotaRequest() (request *UpdateMsgQuotaRequest) {
	request = &UpdateMsgQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "UpdateMsgQuota")
	return
}

func NewUpdateMsgQuotaResponse() (response *UpdateMsgQuotaResponse) {
	response = &UpdateMsgQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改消息配额限制
func (c *Client) UpdateMsgQuota(request *UpdateMsgQuotaRequest) (response *UpdateMsgQuotaResponse, err error) {
	if request == nil {
		request = NewUpdateMsgQuotaRequest()
	}
	response = NewUpdateMsgQuotaResponse()
	err = c.Send(request, response)
	return
}
