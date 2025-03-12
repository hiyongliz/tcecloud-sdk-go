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
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type AddConsoleMessageTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddConsoleMessageTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddConsoleMessageTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelThemeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelThemeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelThemeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Qywx struct {
	// 企业微信内容

	Message *QywxLanguage `json:"Message,omitempty" name:"Message"`
}

type QueryAuditListRequest struct {
	*tchttp.BaseRequest

	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryAuditListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryAuditListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DefineChannelElem struct {
	// 通道号

	SendChannel *string `json:"SendChannel,omitempty" name:"SendChannel"`
	// 内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

type ResMessageTypeList struct {
	// 二级消息分类id

	SecondType *uint64 `json:"SecondType,omitempty" name:"SecondType"`
	// 一级消息分类id

	FirstType *uint64 `json:"FirstType,omitempty" name:"FirstType"`
	// 类型名国际化

	Name *Languagefield `json:"Name,omitempty" name:"Name"`
	// 描述国际化

	Desc *Languagefield `json:"Desc,omitempty" name:"Desc"`
	// 创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 发送通道

	Channels *int64 `json:"Channels,omitempty" name:"Channels"`
	// 默认发送通道

	DefaultChannels *int64 `json:"DefaultChannels,omitempty" name:"DefaultChannels"`
	// 默认策略

	DefaultStrategy *string `json:"DefaultStrategy,omitempty" name:"DefaultStrategy"`
	// 是否订阅类型

	IsDisplay *uint64 `json:"IsDisplay,omitempty" name:"IsDisplay"`
	// 订阅权重

	DisplayWeight *uint64 `json:"DisplayWeight,omitempty" name:"DisplayWeight"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 站点

	Station *int64 `json:"Station,omitempty" name:"Station"`
	// 模式

	InvokePattern *int64 `json:"InvokePattern,omitempty" name:"InvokePattern"`
}

type Sms struct {
	// 短信发送类型

	SendType *string `json:"SendType,omitempty" name:"SendType"`
	// 短信类型

	SmsType *int64 `json:"SmsType,omitempty" name:"SmsType"`
	// 短信内容

	Message *SmsLanguage `json:"Message,omitempty" name:"Message"`
}

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

type Site struct {
	// 是否重要

	IsImportant *int64 `json:"IsImportant,omitempty" name:"IsImportant"`
	// 站内信

	Message *SiteLanguage `json:"Message,omitempty" name:"Message"`
}

type SmsLanguage struct {
	// 英文

	En *SmsElem `json:"En,omitempty" name:"En"`
	// 中文

	Zh *SmsElem `json:"Zh,omitempty" name:"Zh"`
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

type EmailElem struct {
	// email内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// mposssid

	MpbossId *int64 `json:"MpbossId,omitempty" name:"MpbossId"`
}

type Voice struct {
	// 语音内容

	Message *VoiceLanguage `json:"Message,omitempty" name:"Message"`
}

type AddMsgTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddMsgTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMsgTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Email struct {
	// 邮件内容

	Message *EmailLanguage `json:"Message,omitempty" name:"Message"`
	// 邮件发送类型

	SendType *string `json:"SendType,omitempty" name:"SendType"`
}

type FirstType struct {
	// 消息分类id

	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 展示权重

	DisplayWeight *uint64 `json:"DisplayWeight,omitempty" name:"DisplayWeight"`
	// 分类名

	CategoryName *Languagefield `json:"CategoryName,omitempty" name:"CategoryName"`
}

type ThemeDetail struct {
	// 负责人

	Admin *string `json:"Admin,omitempty" name:"Admin"`
	// 一级消息分类详情

	FirstType *FirstType `json:"FirstType,omitempty" name:"FirstType"`
	// 二级消息分类详情

	SecondType *MessageType `json:"SecondType,omitempty" name:"SecondType"`
	// 是否订阅

	IsDisplay *uint64 `json:"IsDisplay,omitempty" name:"IsDisplay"`
	// 创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 二级消息分类id

	MsgType *uint64 `json:"MsgType,omitempty" name:"MsgType"`
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

type SaveThemeRequest struct {
	*tchttp.BaseRequest

	// 管理人

	Admin *string `json:"Admin,omitempty" name:"Admin"`
	// 一级分类

	FirstType *uint64 `json:"FirstType,omitempty" name:"FirstType"`
	// 是否订阅

	IsDisplay *uint64 `json:"IsDisplay,omitempty" name:"IsDisplay"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 二级分类

	MsgType *uint64 `json:"MsgType,omitempty" name:"MsgType"`
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

type EmailLanguage struct {
	// 英文

	En *EmailElem `json:"En,omitempty" name:"En"`
	// 中午

	Zh *EmailElem `json:"Zh,omitempty" name:"Zh"`
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

type SubmitThemeRequest struct {
	*tchttp.BaseRequest

	// 管理人

	Admin *string `json:"Admin,omitempty" name:"Admin"`
	// 一级分类

	FirstType *uint64 `json:"FirstType,omitempty" name:"FirstType"`
	// 是否订阅

	IsDisplay *uint64 `json:"IsDisplay,omitempty" name:"IsDisplay"`
	// 修改人

	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`
	// 二级分类

	MsgType *uint64 `json:"MsgType,omitempty" name:"MsgType"`
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

type FilterString struct {
	// key过滤字段

	Name *string `json:"Name,omitempty" name:"Name"`
	// value过滤字段的值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type Product struct {
	// 产品简称

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 中文名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type SiteLanguage struct {
	// 英文

	En *SiteElem `json:"En,omitempty" name:"En"`
	// 中文

	Zh *SiteElem `json:"Zh,omitempty" name:"Zh"`
}

type QywxElem struct {
	// 内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

type UpdateConsoleMessageTemplateRequest struct {
	*tchttp.BaseRequest

	// 站内信标题

	TitleContent *string `json:"TitleContent,omitempty" name:"TitleContent"`
	// 站内信

	SiteContent *string `json:"SiteContent,omitempty" name:"SiteContent"`
	// 邮件

	EmailContent *string `json:"EmailContent,omitempty" name:"EmailContent"`
	// 短信

	SmsContent *string `json:"SmsContent,omitempty" name:"SmsContent"`
	// 第三方通道

	DefineChannel []*DefineChannelElem `json:"DefineChannel,omitempty" name:"DefineChannel"`
	// 消息分类

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
	// 发送通道

	SendChannel *int64 `json:"SendChannel,omitempty" name:"SendChannel"`
	// 接收人

	Receiver []*ReceiverElem `json:"Receiver,omitempty" name:"Receiver"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 模板id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 企业微信

	QywxContent *string `json:"QywxContent,omitempty" name:"QywxContent"`
	// 语音

	VoiceContent *string `json:"VoiceContent,omitempty" name:"VoiceContent"`
}

func (r *UpdateConsoleMessageTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateConsoleMessageTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QywxLanguage struct {
	// 英文内容

	En *QywxElem `json:"En,omitempty" name:"En"`
	// 中文内容

	Zh *QywxElem `json:"Zh,omitempty" name:"Zh"`
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

type ReceiverElem struct {
	// 主uin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 子uin

	UinList []*int64 `json:"UinList,omitempty" name:"UinList"`
}

type VoiceElem struct {
	// 内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

type QueryMsgTypeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表大小

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 消息类型数据

		Data []*ResMessageTypeList `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryMsgTypeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryMsgTypeListResponse) FromJsonString(s string) error {
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

type DeleteConsoleMessageTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConsoleMessageTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConsoleMessageTemplateResponse) FromJsonString(s string) error {
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

type AddConsoleMessageTemplateRequest struct {
	*tchttp.BaseRequest

	// 站内信标题

	TitleContent *string `json:"TitleContent,omitempty" name:"TitleContent"`
	// 站内信

	SiteContent *string `json:"SiteContent,omitempty" name:"SiteContent"`
	// 邮件

	EmailContent *string `json:"EmailContent,omitempty" name:"EmailContent"`
	// 短信

	SmsContent *string `json:"SmsContent,omitempty" name:"SmsContent"`
	// 第三方通道

	DefineChannel []*DefineChannelElem `json:"DefineChannel,omitempty" name:"DefineChannel"`
	// 消息分类

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
	// 发送通道

	SendChannel *int64 `json:"SendChannel,omitempty" name:"SendChannel"`
	// 接收人

	Receiver []*ReceiverElem `json:"Receiver,omitempty" name:"Receiver"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 企业微信

	QywxContent *string `json:"QywxContent,omitempty" name:"QywxContent"`
	// 语音

	VoiceContent *string `json:"VoiceContent,omitempty" name:"VoiceContent"`
}

func (r *AddConsoleMessageTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddConsoleMessageTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConsoleMessageTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板ID

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteConsoleMessageTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConsoleMessageTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditMsgTypeRequest struct {
	*tchttp.BaseRequest

	// 二级消息分类

	SecondType *uint64 `json:"SecondType,omitempty" name:"SecondType"`
	// 一级消息分类

	FirstType *uint64 `json:"FirstType,omitempty" name:"FirstType"`
	// 二级消息分类名

	Name *Languagefield `json:"Name,omitempty" name:"Name"`
	// 描述

	Desc *Languagefield `json:"Desc,omitempty" name:"Desc"`
	// 创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 发送通道

	Channels *int64 `json:"Channels,omitempty" name:"Channels"`
	// 默认发送通道

	DefaultChannels *int64 `json:"DefaultChannels,omitempty" name:"DefaultChannels"`
	// 角色

	DefaultStrategy *string `json:"DefaultStrategy,omitempty" name:"DefaultStrategy"`
	// 是否订阅

	IsDisplay *int64 `json:"IsDisplay,omitempty" name:"IsDisplay"`
	// 订阅权重

	DisplayWeight *int64 `json:"DisplayWeight,omitempty" name:"DisplayWeight"`
	// 模式

	InvokePattern *int64 `json:"InvokePattern,omitempty" name:"InvokePattern"`
	// 时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 站点

	Station *int64 `json:"Station,omitempty" name:"Station"`
}

func (r *EditMsgTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditMsgTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendConsoleUserMessageRequest struct {
	*tchttp.BaseRequest

	// 消息分类

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
	// 消息标题

	TitleContent *string `json:"TitleContent,omitempty" name:"TitleContent"`
	// 站内信内容

	SiteContent *string `json:"SiteContent,omitempty" name:"SiteContent"`
	// 邮件内容

	EmailContent *string `json:"EmailContent,omitempty" name:"EmailContent"`
	// 短信内容

	SmsContent *string `json:"SmsContent,omitempty" name:"SmsContent"`
	// 自定义渠道

	DefineChannel []*DefineChannelElem `json:"DefineChannel,omitempty" name:"DefineChannel"`
	// 发送通道

	SendChannel *string `json:"SendChannel,omitempty" name:"SendChannel"`
	// 接收人

	Receiver []*ReceiverElem `json:"Receiver,omitempty" name:"Receiver"`
	// 语言变量

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 企业微信

	QywxContent *string `json:"QywxContent,omitempty" name:"QywxContent"`
	// 语音

	VoiceContent *string `json:"VoiceContent,omitempty" name:"VoiceContent"`
}

func (r *SendConsoleUserMessageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendConsoleUserMessageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMsgCategorysRequest struct {
	*tchttp.BaseRequest
}

func (r *GetMsgCategorysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMsgCategorysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMsgCategorysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 一级消息分类

		Data []*FirstType `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMsgCategorysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMsgCategorysResponse) FromJsonString(s string) error {
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

type QueryConsoleMessageTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryConsoleMessageTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryConsoleMessageTemplateResponse) FromJsonString(s string) error {
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

type MessageType struct {
	// 消息类型id

	MsgType *int64 `json:"MsgType,omitempty" name:"MsgType"`
	// 一级消息分类ID

	FType *uint64 `json:"FType,omitempty" name:"FType"`
	// 消息类型名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 发送通道

	Channels *int64 `json:"Channels,omitempty" name:"Channels"`
	// 默认发送通道

	DefaultChannels *int64 `json:"DefaultChannels,omitempty" name:"DefaultChannels"`
	// 默认策略

	DefaultStrategy *string `json:"DefaultStrategy,omitempty" name:"DefaultStrategy"`
	// 是否订阅

	IsDisplay *uint64 `json:"IsDisplay,omitempty" name:"IsDisplay"`
	// 订阅权重

	DisplayWeight *uint64 `json:"DisplayWeight,omitempty" name:"DisplayWeight"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 类型名国家化

	NameW *string `json:"NameW,omitempty" name:"NameW"`
	// 类型描述国家化

	DescW *string `json:"DescW,omitempty" name:"DescW"`
	// 站点

	Station *int64 `json:"Station,omitempty" name:"Station"`
}

type QueryMsgTypeListRequest struct {
	*tchttp.BaseRequest

	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryMsgTypeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryMsgTypeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendTestMessageThemeRequest struct {
	*tchttp.BaseRequest

	// 模板id

	ThemeId *uint64 `json:"ThemeId,omitempty" name:"ThemeId"`
	// 主账号

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 主账号

	SubAccountList []*int64 `json:"SubAccountList,omitempty" name:"SubAccountList"`
	// 语言变量

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 发送通道

	SendChannel *int64 `json:"SendChannel,omitempty" name:"SendChannel"`
}

func (r *SendTestMessageThemeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendTestMessageThemeRequest) FromJsonString(s string) error {
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

type EditMsgTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditMsgTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditMsgTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SmsElem struct {
	// 短信内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

type CancelThemeRequest struct {
	*tchttp.BaseRequest

	// 模板id

	ThemeId *int64 `json:"ThemeId,omitempty" name:"ThemeId"`
}

func (r *CancelThemeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelThemeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendConsoleUserMessageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendConsoleUserMessageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendConsoleUserMessageResponse) FromJsonString(s string) error {
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
	// 一级二级消息分类详情

	MsgType *MsgType `json:"MsgType,omitempty" name:"MsgType"`
}

type AddMsgTypeRequest struct {
	*tchttp.BaseRequest

	// 二级消息分类

	SecondType *uint64 `json:"SecondType,omitempty" name:"SecondType"`
	// 一级消息分类

	FirstType *uint64 `json:"FirstType,omitempty" name:"FirstType"`
	// 二级消息分类名

	Name *Languagefield `json:"Name,omitempty" name:"Name"`
	// 描述

	Desc *Languagefield `json:"Desc,omitempty" name:"Desc"`
	// 创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 发送通道

	Channels *int64 `json:"Channels,omitempty" name:"Channels"`
	// 默认发送通道

	DefaultChannels *int64 `json:"DefaultChannels,omitempty" name:"DefaultChannels"`
	// 角色

	DefaultStrategy *string `json:"DefaultStrategy,omitempty" name:"DefaultStrategy"`
	// 是否订阅

	IsDisplay *int64 `json:"IsDisplay,omitempty" name:"IsDisplay"`
	// 订阅权重

	DisplayWeight *int64 `json:"DisplayWeight,omitempty" name:"DisplayWeight"`
	// 模式

	InvokePattern *int64 `json:"InvokePattern,omitempty" name:"InvokePattern"`
	// 时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 站点

	Station *int64 `json:"Station,omitempty" name:"Station"`
}

func (r *AddMsgTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMsgTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsgType struct {
	// 一级分类ID

	FirstType *uint64 `json:"FirstType,omitempty" name:"FirstType"`
	// 一级分类详情

	FirstTypeName *Languagefield `json:"FirstTypeName,omitempty" name:"FirstTypeName"`
	// 是否订阅类型

	IsDisplay *int64 `json:"IsDisplay,omitempty" name:"IsDisplay"`
	// 二级分类详情

	SecondTypeName *Languagefield `json:"SecondTypeName,omitempty" name:"SecondTypeName"`
	// 二级分类id

	SecondType *uint64 `json:"SecondType,omitempty" name:"SecondType"`
}

type SiteElem struct {
	// 内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 标题

	Title *string `json:"Title,omitempty" name:"Title"`
}

type QueryConsoleMessageTemplateRequest struct {
	*tchttp.BaseRequest

	// 分页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏离量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数

	Filters []*FilterString `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryConsoleMessageTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryConsoleMessageTemplateRequest) FromJsonString(s string) error {
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

type UpdateConsoleMessageTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateConsoleMessageTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateConsoleMessageTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Languagefield struct {
	// 中文

	Zh *string `json:"Zh,omitempty" name:"Zh"`
	// 英文

	En *string `json:"En,omitempty" name:"En"`
}

type VoiceLanguage struct {
	// 英文内容

	En *VoiceElem `json:"En,omitempty" name:"En"`
	// 中文内容

	Zh *VoiceElem `json:"Zh,omitempty" name:"Zh"`
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

type SendTestMessageThemeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendTestMessageThemeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendTestMessageThemeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitAuditRequest struct {
	*tchttp.BaseRequest

	// 模板id

	ThemeId *uint64 `json:"ThemeId,omitempty" name:"ThemeId"`
}

func (r *SubmitAuditRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubmitAuditRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
	// 一级消息分类

	FirstTypeName *Languagefield `json:"FirstTypeName,omitempty" name:"FirstTypeName"`
	// 是否订阅类型

	IsDisplay *uint64 `json:"IsDisplay,omitempty" name:"IsDisplay"`
	// 二级消息分类id

	SecondType *uint64 `json:"SecondType,omitempty" name:"SecondType"`
	// 二级消息分类消息

	SecondTypeName *Languagefield `json:"SecondTypeName,omitempty" name:"SecondTypeName"`
}
