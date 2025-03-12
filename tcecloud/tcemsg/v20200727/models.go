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
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type PublishThemeRequest struct {
	*tchttp.BaseRequest

	// 主键id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *PublishThemeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishThemeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QywxLanguage struct {
	// 英文内容

	En *QywxElem `json:"En,omitempty" name:"En"`
	// 中文内容

	Zh *QywxElem `json:"Zh,omitempty" name:"Zh"`
}

type GetReserveSendChannelNumberRequest struct {
	*tchttp.BaseRequest
}

func (r *GetReserveSendChannelNumberRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetReserveSendChannelNumberRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveThemeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveThemeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveThemeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitAuditResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitAuditResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubmitAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CamGroup struct {
	// Cam 数据库主键

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 用户组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type ResMessageThemeList struct {
	// 模板id

	ThemeId *uint64 `json:"ThemeId,omitempty" name:"ThemeId"`
	// 模板主题

	Title *Languagefield `json:"Title,omitempty" name:"Title"`
	// 模板状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 负责人

	Admin *string `json:"Admin,omitempty" name:"Admin"`
	// 创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 产品

	Product *Product `json:"Product,omitempty" name:"Product"`
}

type ReadUserCmgtSiteMsgRequest struct {
	*tchttp.BaseRequest

	// 数据库主键列表(通过查询接口获取)

	IDs []*int64 `json:"IDs,omitempty" name:"IDs"`
	// 是否“全部标记为已读”

	ReadAll *bool `json:"ReadAll,omitempty" name:"ReadAll"`
}

func (r *ReadUserCmgtSiteMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReadUserCmgtSiteMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateThemeAdminRequest struct {
	*tchttp.BaseRequest

	// 模版ID

	ThemeId *uint64 `json:"ThemeId,omitempty" name:"ThemeId"`
	// 负责人

	Admin *string `json:"Admin,omitempty" name:"Admin"`
}

func (r *UpdateThemeAdminRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateThemeAdminRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryThemeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 消息模板数据

		Data []*ResMessageThemeList `json:"Data,omitempty" name:"Data"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryThemeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryThemeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySendChannelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 发送通道总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 发送通道列表

		SendChannels []*TSendChannel `json:"SendChannels,omitempty" name:"SendChannels"`
		// 是否第三方渠道(yes-是，no-否)

		IsThirdparty *string `json:"IsThirdparty,omitempty" name:"IsThirdparty"`
		// 渠道名称

		ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySendChannelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySendChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelUserCmgtSiteMsgRequest struct {
	*tchttp.BaseRequest

	// 数据库主键列表(通过查询接口获取)

	IDs []*int64 `json:"IDs,omitempty" name:"IDs"`
}

func (r *DelUserCmgtSiteMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelUserCmgtSiteMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchSendChannelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchSendChannelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchSendChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefuseApplyRequest struct {
	*tchttp.BaseRequest

	// 主键id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 驳回原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

func (r *RefuseApplyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefuseApplyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteThemeRequest struct {
	*tchttp.BaseRequest

	// 模版ID

	ThemeId *uint64 `json:"ThemeId,omitempty" name:"ThemeId"`
}

func (r *DeleteThemeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteThemeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VoiceElem struct {
	// 内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

type GetUserCmgtSiteMsgRequest struct {
	*tchttp.BaseRequest

	// 数据库主键(通过查询接口获取)

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *GetUserCmgtSiteMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserCmgtSiteMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddPostRequest struct {
	*tchttp.BaseRequest

	// 公告内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 发表者Userid

	Poster *string `json:"Poster,omitempty" name:"Poster"`
}

func (r *AddPostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddPostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveThemeRequest struct {
	*tchttp.BaseRequest

	// 管理人

	Admin *string `json:"Admin,omitempty" name:"Admin"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 发送通道

	SendChannel *int64 `json:"SendChannel,omitempty" name:"SendChannel"`
	// 模版id

	ThemeId *uint64 `json:"ThemeId,omitempty" name:"ThemeId"`
	// 模版主题

	Title *Languagefield `json:"Title,omitempty" name:"Title"`
	// 产品

	Product *Product `json:"Product,omitempty" name:"Product"`
	// 邮件模版

	Email *Email `json:"Email,omitempty" name:"Email"`
	// 站内信模版

	Site *Site `json:"Site,omitempty" name:"Site"`
	// 短信模版

	Sms *Sms `json:"Sms,omitempty" name:"Sms"`
	// 创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 企业微信

	Qywx *Qywx `json:"Qywx,omitempty" name:"Qywx"`
	// 语音

	Voice *Voice `json:"Voice,omitempty" name:"Voice"`
}

func (r *SaveThemeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveThemeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResMessageAuditList struct {
	// 主键ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 模板ID

	ThemeId *uint64 `json:"ThemeId,omitempty" name:"ThemeId"`
	// 模板主题

	Title *Languagefield `json:"Title,omitempty" name:"Title"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 产品

	Product *Product `json:"Product,omitempty" name:"Product"`
	// 申请人

	Applicant *string `json:"Applicant,omitempty" name:"Applicant"`
	// 原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 提交时间

	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`
	// 完成时间

	FinishedTime *string `json:"FinishedTime,omitempty" name:"FinishedTime"`
	// 当前处理人

	CurHandlerPerson *string `json:"CurHandlerPerson,omitempty" name:"CurHandlerPerson"`
}

type SmsElem struct {
	// 短信内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

type TEmailMessageFlow struct {
	// 数据库主键

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 运营端消息ID(语意: 哪次运营端消息产生了这条站内信)

	MessageID *int64 `json:"MessageID,omitempty" name:"MessageID"`
	// 邮件标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 邮件内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 邮件接收邮箱

	Receiver *string `json:"Receiver,omitempty" name:"Receiver"`
	// 邮件抄送列表

	CcReceiver *string `json:"CcReceiver,omitempty" name:"CcReceiver"`
	// 邮件密送列表

	BccReceiver *string `json:"BccReceiver,omitempty" name:"BccReceiver"`
	// 邮件发送状态(pending:待发送; processing:发送中; ok: 发送成功; fail:  发送失败;)

	Status *string `json:"Status,omitempty" name:"Status"`
	// 邮件创建时间

	Createtime *string `json:"Createtime,omitempty" name:"Createtime"`
	// 消息源

	Sender *string `json:"Sender,omitempty" name:"Sender"`
	// 接收者Uin

	ReceiverUin *string `json:"ReceiverUin,omitempty" name:"ReceiverUin"`
}

type DeleteSendChannelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSendChannelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSendChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddTopicRequest struct {
	*tchttp.BaseRequest

	// 订阅名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 订阅描述

	TopicDesc *string `json:"TopicDesc,omitempty" name:"TopicDesc"`
	// 订阅使用的发送通道列表(or)

	SendChannel *int64 `json:"SendChannel,omitempty" name:"SendChannel"`
	// 订阅用户(UIN列表)

	Users []*int64 `json:"Users,omitempty" name:"Users"`
	// 订阅用户组(Group列表)

	Groups []*int64 `json:"Groups,omitempty" name:"Groups"`
	// 内部使用(预设数据)

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 内部使用(预设数据)

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

func (r *AddTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelPostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DelPostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelPostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAuditListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回大小

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 消息审核

		Data []*ResMessageAuditList `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAuditListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryAuditListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResendMsgRequest struct {
	*tchttp.BaseRequest

	// 主键id

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
	// 消息渠道

	Channel *int64 `json:"Channel,omitempty" name:"Channel"`
}

func (r *ResendMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResendMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 字段名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段值列表

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DelUserCmgtSiteMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DelUserCmgtSiteMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelUserCmgtSiteMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitThemeRequest struct {
	*tchttp.BaseRequest

	// 管理人

	Admin *string `json:"Admin,omitempty" name:"Admin"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 发送通道

	SendChannel *int64 `json:"SendChannel,omitempty" name:"SendChannel"`
	// 模版id

	ThemeId *uint64 `json:"ThemeId,omitempty" name:"ThemeId"`
	// 模版主题

	Title *Languagefield `json:"Title,omitempty" name:"Title"`
	// 产品

	Product *Product `json:"Product,omitempty" name:"Product"`
	// 邮件模版

	Email *Email `json:"Email,omitempty" name:"Email"`
	// 站内信模版

	Site *Site `json:"Site,omitempty" name:"Site"`
	// 短信模版

	Sms *Sms `json:"Sms,omitempty" name:"Sms"`
	// 创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 企业微信

	Qywx *Qywx `json:"Qywx,omitempty" name:"Qywx"`
	// 语音

	Voice *Voice `json:"Voice,omitempty" name:"Voice"`
}

func (r *SubmitThemeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubmitThemeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TQywxMessageFlow struct {
	// 数据库主键

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 运营端消息ID(语意: 哪次运营端消息产生了这条站内信)

	MessageID *int64 `json:"MessageID,omitempty" name:"MessageID"`
	// 企业微信内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 企业微信号

	Receiver *string `json:"Receiver,omitempty" name:"Receiver"`
	// 发送状态(pending:待发送; processing:发送中; ok: 发送成功; fail:  发送失败;)

	Status *string `json:"Status,omitempty" name:"Status"`
	// 创建时间

	Createtime *string `json:"Createtime,omitempty" name:"Createtime"`
	// 接收者Uin

	ReceiverUin *string `json:"ReceiverUin,omitempty" name:"ReceiverUin"`
	// 消息源

	Sender *string `json:"Sender,omitempty" name:"Sender"`
	// 消息模板

	ThemeID *int64 `json:"ThemeID,omitempty" name:"ThemeID"`
	// 主账号

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

type QuerySendChannelRequest struct {
	*tchttp.BaseRequest

	// 记录偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 记录分页数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 是否第三方渠道

	IsThirdparty *string `json:"IsThirdparty,omitempty" name:"IsThirdparty"`
	// 渠道名称

	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`
}

func (r *QuerySendChannelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySendChannelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 发送通道

		SendChannel *int64 `json:"SendChannel,omitempty" name:"SendChannel"`
		// 站内信列表

		SiteMsgs []*TSiteMessageFlow `json:"SiteMsgs,omitempty" name:"SiteMsgs"`
		// 邮件列表

		EmailMsgs []*TEmailMessageFlow `json:"EmailMsgs,omitempty" name:"EmailMsgs"`
		// 短信列表

		PhoneMsgs []*TSmsMessageFlow `json:"PhoneMsgs,omitempty" name:"PhoneMsgs"`
		// 第三方消息列表

		ReserveMsgs []*TThirdpartyMessageFlow `json:"ReserveMsgs,omitempty" name:"ReserveMsgs"`
		// 消息总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 企业微信列表

		QywxMsgs []*TQywxMessageFlow `json:"QywxMsgs,omitempty" name:"QywxMsgs"`
		// 渠道描述，用于错误信息展示

		ChannelDesc *string `json:"ChannelDesc,omitempty" name:"ChannelDesc"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryThemeListRequest struct {
	*tchttp.BaseRequest

	// 页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 数据偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤查询字段

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryThemeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryThemeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetReserveSendChannelNumberResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 发送通道编号列表(已过滤:当前可用)

		ChannelNumbers []*int64 `json:"ChannelNumbers,omitempty" name:"ChannelNumbers"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetReserveSendChannelNumberResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetReserveSendChannelNumberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Qywx struct {
	// 企业微信内容

	Message *QywxLanguage `json:"Message,omitempty" name:"Message"`
}

type RefuseApplyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RefuseApplyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefuseApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SiteLanguage struct {
	// 英文

	En *SiteElem `json:"En,omitempty" name:"En"`
	// 中文

	Zh *SiteElem `json:"Zh,omitempty" name:"Zh"`
}

type SmsLanguage struct {
	// 英文

	En *SmsElem `json:"En,omitempty" name:"En"`
	// 中文

	Zh *SmsElem `json:"Zh,omitempty" name:"Zh"`
}

type DeleteSendChannelRequest struct {
	*tchttp.BaseRequest

	// 数据库主键(通过查询接口获取)

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 发送通道编号

	ChannelNumber *int64 `json:"ChannelNumber,omitempty" name:"ChannelNumber"`
}

func (r *DeleteSendChannelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSendChannelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetThemeDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板详情

		Data *ThemeDetail `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetThemeDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetThemeDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EmailLanguage struct {
	// 英文

	En *EmailElem `json:"En,omitempty" name:"En"`
	// 中文

	Zh *EmailElem `json:"Zh,omitempty" name:"Zh"`
}

type SubmitAuditRequest struct {
	*tchttp.BaseRequest

	// 模版ID

	ThemeId *uint64 `json:"ThemeId,omitempty" name:"ThemeId"`
}

func (r *SubmitAuditRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubmitAuditRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Email struct {
	// 邮件内容

	Message *EmailLanguage `json:"Message,omitempty" name:"Message"`
	// 邮件发送类型

	SendType *string `json:"SendType,omitempty" name:"SendType"`
}

type SwitchSendChannelRequest struct {
	*tchttp.BaseRequest

	// 数据库主键(通过查询接口获得)

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 发送通道编号

	ChannelNumber *int64 `json:"ChannelNumber,omitempty" name:"ChannelNumber"`
	// 开关(enable:开; disable:关)

	Enable *string `json:"Enable,omitempty" name:"Enable"`
}

func (r *SwitchSendChannelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchSendChannelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetThemeDetailRequest struct {
	*tchttp.BaseRequest

	// 消息模板ID

	ThemeId *uint64 `json:"ThemeId,omitempty" name:"ThemeId"`
}

func (r *GetThemeDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetThemeDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryUserCmgtSiteMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 站内信列表

		SiteMsgs []*TSiteMessageFlow `json:"SiteMsgs,omitempty" name:"SiteMsgs"`
		// 站内信总数(已过滤:Status)

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryUserCmgtSiteMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryUserCmgtSiteMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TAnnouncement struct {
	// Primary key in a database.

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// Announcement Content

	Content *string `json:"Content,omitempty" name:"Content"`
	// Announcement creator.

	Poster *string `json:"Poster,omitempty" name:"Poster"`
	// Announcement editor who does the final touch.

	Writer *string `json:"Writer,omitempty" name:"Writer"`
	// Announcement creation time.

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
}

type DelTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DelTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 发送通道

		SendChannel *int64 `json:"SendChannel,omitempty" name:"SendChannel"`
		// 站内信

		SiteMsg *TSiteMessageFlow `json:"SiteMsg,omitempty" name:"SiteMsg"`
		// 邮件

		EmailMsg *TEmailMessageFlow `json:"EmailMsg,omitempty" name:"EmailMsg"`
		// 短信

		PhoneMsg *TSmsMessageFlow `json:"PhoneMsg,omitempty" name:"PhoneMsg"`
		// 第三方消息

		ReserveMsg *TThirdpartyMessageFlow `json:"ReserveMsg,omitempty" name:"ReserveMsg"`
		// 企业微信

		QywxMsg *TQywxMessageFlow `json:"QywxMsg,omitempty" name:"QywxMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSendChannelRequest struct {
	*tchttp.BaseRequest

	// 消息通道名称

	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`
	// 消息插件地址

	PluginUrl *string `json:"PluginUrl,omitempty" name:"PluginUrl"`
	// 租户端账号扩展Tag

	ConsoleReceiverTag *string `json:"ConsoleReceiverTag,omitempty" name:"ConsoleReceiverTag"`
	// 运营端账号扩展Tag

	CmgtReceiverTag *string `json:"CmgtReceiverTag,omitempty" name:"CmgtReceiverTag"`
}

func (r *AddSendChannelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSendChannelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAuditListRequest struct {
	*tchttp.BaseRequest

	// 页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 数据偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤查询字段

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryAuditListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryAuditListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Voice struct {
	// 语音内容

	Message *VoiceLanguage `json:"Message,omitempty" name:"Message"`
}

type GetMsgRequest struct {
	*tchttp.BaseRequest

	// 数据库主键(通过查询接口获得)

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 发送通道

	SendChannel *int64 `json:"SendChannel,omitempty" name:"SendChannel"`
	// 租户端或运营端

	Source *string `json:"Source,omitempty" name:"Source"`
}

func (r *GetMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReadUserCmgtSiteMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReadUserCmgtSiteMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReadUserCmgtSiteMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ThemeDetail struct {
	// 负责人

	Admin *string `json:"Admin,omitempty" name:"Admin"`
	// 创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 发送通道

	SendChannel *int64 `json:"SendChannel,omitempty" name:"SendChannel"`
	// 模板id

	ThemeId *uint64 `json:"ThemeId,omitempty" name:"ThemeId"`
	// 模板主题

	Title *Languagefield `json:"Title,omitempty" name:"Title"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 产品ID

	ProductId *int64 `json:"ProductId,omitempty" name:"ProductId"`
	// 产品

	Product *Product `json:"Product,omitempty" name:"Product"`
	// 邮件

	Email *Email `json:"Email,omitempty" name:"Email"`
	// 站内信

	Site *Site `json:"Site,omitempty" name:"Site"`
	// 短信

	Sms *Sms `json:"Sms,omitempty" name:"Sms"`
	// 企业微信

	Qywx *Qywx `json:"Qywx,omitempty" name:"Qywx"`
	// 语音

	Voice *Voice `json:"Voice,omitempty" name:"Voice"`
}

type GetTopicRequest struct {
	*tchttp.BaseRequest

	// 记录偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 记录分页数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserCmgtSiteMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 站内信详情

		SiteMsg *TSiteMessageFlow `json:"SiteMsg,omitempty" name:"SiteMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserCmgtSiteMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserCmgtSiteMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterString struct {
	// key过滤字段

	Name *string `json:"Name,omitempty" name:"Name"`
	// value过滤字段的值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type QueryPostRequest struct {
	*tchttp.BaseRequest

	// 记录偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 记录分页数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryPostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Product struct {
	// 产品简称

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 中文名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type AddSendChannelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新增通道信息

		SendChannel *TSendChannel `json:"SendChannel,omitempty" name:"SendChannel"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSendChannelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSendChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteThemeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteThemeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteThemeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddPostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddPostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddPostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TThirdpartyMessageFlow struct {
	// 数据库主键

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 运营端消息ID(语意: 哪次运营端消息产生了这条消息)

	MessageID *uint64 `json:"MessageID,omitempty" name:"MessageID"`
	// 消息标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 消息内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 消息接收者(第三方账号)

	Receiver *string `json:"Receiver,omitempty" name:"Receiver"`
	// 消息发送通道

	SendChannel *int64 `json:"SendChannel,omitempty" name:"SendChannel"`
	// 消息发送状态(pending:待发送; processing:发送中; ok: 发送成功; fail:  发送失败;)

	Status *string `json:"Status,omitempty" name:"Status"`
	// 消息创建时间

	Createtime *string `json:"Createtime,omitempty" name:"Createtime"`
	// 消息源

	Sender *string `json:"Sender,omitempty" name:"Sender"`
	// 接收者Uin

	ReceiverUin *string `json:"ReceiverUin,omitempty" name:"ReceiverUin"`
}

type QueryPostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公告列表

		List []*TAnnouncement `json:"List,omitempty" name:"List"`
		// 公告总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryPostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QywxElem struct {
	// 内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

type DelTopicRequest struct {
	*tchttp.BaseRequest

	// 数据库主键(通过查询接口获取)

	TopicID *int64 `json:"TopicID,omitempty" name:"TopicID"`
}

func (r *DelTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSendChannelRequest struct {
	*tchttp.BaseRequest

	// 数据库主键(通过查询接口获得)

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 发送通道编号[58~62]

	ChannelNumber *int64 `json:"ChannelNumber,omitempty" name:"ChannelNumber"`
	// 发送通道名称

	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`
	// 发送插件地址

	PluginURL *string `json:"PluginURL,omitempty" name:"PluginURL"`
	// 租户端账号扩展Tag

	ConsoleReceiverTag *string `json:"ConsoleReceiverTag,omitempty" name:"ConsoleReceiverTag"`
	// 运营端账号扩展Tag

	CmgtReceiverTag *string `json:"CmgtReceiverTag,omitempty" name:"CmgtReceiverTag"`
}

func (r *UpdateSendChannelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSendChannelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TSendChannel struct {
	// 数据库主键(通过查询接口获取)

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 发送通道编号(前端展示使用)

	ChannelNumber *string `json:"ChannelNumber,omitempty" name:"ChannelNumber"`
	// 发送通道名称

	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`
	// 发送通道(2^x)

	SendChannel *int64 `json:"SendChannel,omitempty" name:"SendChannel"`
	// 是否开启(enable:开启; disable:关闭)

	Enable *string `json:"Enable,omitempty" name:"Enable"`
	// 是否为第三方通道(yes: 是; no:否)

	IsThirdparty *string `json:"IsThirdparty,omitempty" name:"IsThirdparty"`
	// 插件地址

	PluginUrl *string `json:"PluginUrl,omitempty" name:"PluginUrl"`
	// 插件调用超时(ms)

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	// 租户端账号扩展Tag

	ConsoleReceiverTag *string `json:"ConsoleReceiverTag,omitempty" name:"ConsoleReceiverTag"`
	// 运营端账号扩展Tag

	CmgtReceiverTag *string `json:"CmgtReceiverTag,omitempty" name:"CmgtReceiverTag"`
	// 创建时间

	Createtime *string `json:"Createtime,omitempty" name:"Createtime"`
	// 更新时间

	Updatetime *string `json:"Updatetime,omitempty" name:"Updatetime"`
}

type PublishThemeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishThemeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishThemeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateThemeAdminResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateThemeAdminResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateThemeAdminResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTopicRequest struct {
	*tchttp.BaseRequest

	// 数据库主键(通过查询接口获取)

	TopicID *int64 `json:"TopicID,omitempty" name:"TopicID"`
	// 订阅名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 订阅描述

	TopicDesc *string `json:"TopicDesc,omitempty" name:"TopicDesc"`
	// 订阅使用的发送通道列表(or)

	SendChannel *int64 `json:"SendChannel,omitempty" name:"SendChannel"`
	// 订阅用户(UIN列表)

	Users []*int64 `json:"Users,omitempty" name:"Users"`
	// 订阅用户组(GroupID列表)

	Groups []*int64 `json:"Groups,omitempty" name:"Groups"`
	// 内部使用(预设数据)

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 内部使用(预设数据)

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

func (r *UpdateTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePostRequest struct {
	*tchttp.BaseRequest

	// 数据库主键(通过查询接口获取)

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 编辑者(当前登陆用户Userid)

	Writer *string `json:"Writer,omitempty" name:"Writer"`
	// 公告内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *UpdatePostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitThemeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitThemeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubmitThemeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 订阅记录列表

		Topics []*TopicExt `json:"Topics,omitempty" name:"Topics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryUserCmgtSiteMsgRequest struct {
	*tchttp.BaseRequest

	// 记录偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 记录分页数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤器

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryUserCmgtSiteMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryUserCmgtSiteMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Languagefield struct {
	// 中文

	Zh *string `json:"Zh,omitempty" name:"Zh"`
	// 英文

	En *string `json:"En,omitempty" name:"En"`
}

type TSmsMessageFlow struct {
	// 数据库主键

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 运营端消息ID(语意: 哪次运营端消息产生了这条站内信)

	MessageID *int64 `json:"MessageID,omitempty" name:"MessageID"`
	// 短信内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 短信接收手机号

	Receiver *string `json:"Receiver,omitempty" name:"Receiver"`
	// 短信发送状态(pending:待发送; processing:发送中; ok: 发送成功; fail:  发送失败;)

	Status *string `json:"Status,omitempty" name:"Status"`
	// 短信创建时间

	Createtime *string `json:"Createtime,omitempty" name:"Createtime"`
	// 接收者Uin

	ReceiverUin *string `json:"ReceiverUin,omitempty" name:"ReceiverUin"`
	// 电话区号

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 消息源

	Sender *string `json:"Sender,omitempty" name:"Sender"`
}

type UpdateSendChannelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSendChannelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSendChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Site struct {
	// 是否重要

	IsImportant *int64 `json:"IsImportant,omitempty" name:"IsImportant"`
	// 站内信

	Message *SiteLanguage `json:"Message,omitempty" name:"Message"`
}

type AddTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CamUser struct {
	// cam 数据库主键

	Uid *int64 `json:"Uid,omitempty" name:"Uid"`
	// Uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 登陆账号名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 手机号

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 邮箱地址

	Email *string `json:"Email,omitempty" name:"Email"`
	// 手机号国家编码

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
}

type ResendMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResendMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResendMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelPostRequest struct {
	*tchttp.BaseRequest

	// 数据库主键(通过查询接口获取)

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DelPostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelPostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMsgRequest struct {
	*tchttp.BaseRequest

	// 发送通道

	SendChannel *int64 `json:"SendChannel,omitempty" name:"SendChannel"`
	// 过滤器

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 记录偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 记录分页数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Sms struct {
	// 短信发送类型

	SendType *string `json:"SendType,omitempty" name:"SendType"`
	// 短信类型

	SmsType *int64 `json:"SmsType,omitempty" name:"SmsType"`
	// 短信内容

	Message *SiteLanguage `json:"Message,omitempty" name:"Message"`
}

type SiteElem struct {
	// 内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 标题

	Title *string `json:"Title,omitempty" name:"Title"`
}

type TopicExt struct {
	// 数据库主键(发送接口使用)

	TopicID *int64 `json:"TopicID,omitempty" name:"TopicID"`
	// 订阅模版的名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 订阅模版的描述

	TopicDesc *string `json:"TopicDesc,omitempty" name:"TopicDesc"`
	// 订阅模板使用的发送通道列表(or)

	SendChannel *string `json:"SendChannel,omitempty" name:"SendChannel"`
	// 订阅账号列表(UIN)

	Users []*CamUser `json:"Users,omitempty" name:"Users"`
	// 订阅用户组列表(Group)

	Groups []*CamGroup `json:"Groups,omitempty" name:"Groups"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最后更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 产品名称(内部使用:预设数据)

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称(内部使用:预设数据)

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type EmailElem struct {
	// email内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// email标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// mposssid

	MpbossId *int64 `json:"MpbossId,omitempty" name:"MpbossId"`
}

type TSiteMessageFlow struct {
	// 数据库主键

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 站内信接收者(运营端子账号)

	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
	// 空(前端看下是否使用?)

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 是否已读(pending:未读取, 其它:已读取)

	Status *string `json:"Status,omitempty" name:"Status"`
	// 运营端消息ID(语意: 哪次运营端消息产生了这条站内信)

	MessageID *int64 `json:"MessageID,omitempty" name:"MessageID"`
	// 站内信标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 站内信内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 站内信创建时间

	Createtime *string `json:"Createtime,omitempty" name:"Createtime"`
	// 消息源

	Sender *string `json:"Sender,omitempty" name:"Sender"`
}

type VoiceLanguage struct {
	// 英文内容

	En *VoiceElem `json:"En,omitempty" name:"En"`
	// 中文内容

	Zh *VoiceElem `json:"Zh,omitempty" name:"Zh"`
}
