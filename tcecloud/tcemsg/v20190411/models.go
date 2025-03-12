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

package v20190411

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type QueryPostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公告内容

		List []*Post `json:"List,omitempty" name:"List"`
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

type QueryPostRequest struct {
	*tchttp.BaseRequest

	// 翻页参数: 公告开始的序号 (发布时间倒序)

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 翻页参数: 消息条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryPostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopicRes struct {
	// 数据库编号

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 编号

	TopicID *int64 `json:"TopicID,omitempty" name:"TopicID"`
	// 名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 消息类型

	MsgType *string `json:"MsgType,omitempty" name:"MsgType"`
	// 接收用户列表

	Users []*CamUserInfo `json:"Users,omitempty" name:"Users"`
	// 接受组列表

	Groups []*GroupIDName `json:"Groups,omitempty" name:"Groups"`
	// 描述信息

	TopicDesc *string `json:"TopicDesc,omitempty" name:"TopicDesc"`
	// 创建时间. 格式: 2006-01-02 15:04:05

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建时间. 格式: 2006-01-02 15:04:05

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type CamUserInfo struct {
	// 用户编号

	Uid *int64 `json:"Uid,omitempty" name:"Uid"`
	// Uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 用户名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 手机号码

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 国家编号

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
}

type Receiver struct {
	// OwnerUin

	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// UinList

	UinList []*uint64 `json:"UinList,omitempty" name:"UinList"`
}

type CmgtSiteMsg struct {
	// 数据库id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 消息编号

	MessageId *int64 `json:"MessageId,omitempty" name:"MessageId"`
	// 发送者

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 发送者名称

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 发送状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 消息内容

	Message *Message2 `json:"Message,omitempty" name:"Message"`
}

type Message struct {
	// 消息编号

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// 发送者UIN

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 发送日期

	DateTime *string `json:"DateTime,omitempty" name:"DateTime"`
	// 发送者

	Sender *string `json:"Sender,omitempty" name:"Sender"`
	// 接受者

	Receiver *string `json:"Receiver,omitempty" name:"Receiver"`
	// 消息标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 消息内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 消息类型: phone / email / web

	MsgType *string `json:"MsgType,omitempty" name:"MsgType"`
	// 消息发送状态: fail / ok

	Status *string `json:"Status,omitempty" name:"Status"`
}

type Post struct {
	// 公告编号

	PostId *int64 `json:"PostId,omitempty" name:"PostId"`
	// 公告内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 公告发布者

	Poster *string `json:"Poster,omitempty" name:"Poster"`
	// 公告发布者(等于Poster)

	Writer *string `json:"Writer,omitempty" name:"Writer"`
	// 公告发布时间, 格式满足: 2006-01-02 15:04:05

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
}

type DelUserCmgtSiteMsgRequest struct {
	*tchttp.BaseRequest

	// id数组

	IDs []*int64 `json:"IDs,omitempty" name:"IDs"`
}

func (r *DelUserCmgtSiteMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelUserCmgtSiteMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filters struct {
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Values

	Values []*string `json:"Values,omitempty" name:"Values"`
	// Op

	Op *string `json:"Op,omitempty" name:"Op"`
}

type AddPostRequest struct {
	*tchttp.BaseRequest

	// 公告内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 当前登陆用户名

	Poster *string `json:"Poster,omitempty" name:"Poster"`
}

func (r *AddPostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddPostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type TCloundAccount struct {
	// 主账号

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 账号状态

	Active *int64 `json:"Active,omitempty" name:"Active"`
	// AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 邮箱

	Mail *string `json:"Mail,omitempty" name:"Mail"`
	// 昵称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 电话号码

	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
	// 无

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 账号ID

	Uid *int64 `json:"Uid,omitempty" name:"Uid"`
	// Uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 账号类型

	UinType *string `json:"UinType,omitempty" name:"UinType"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 无

	Account *string `json:"Account,omitempty" name:"Account"`
}

type TemplateRes struct {
	// 模板ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 标题

	TitleContent *string `json:"TitleContent,omitempty" name:"TitleContent"`
	// 站内信模板内容

	SiteContent *string `json:"SiteContent,omitempty" name:"SiteContent"`
	// 邮件模板内容

	EmailContent *string `json:"EmailContent,omitempty" name:"EmailContent"`
	// 短信模板内容

	SmsContent *string `json:"SmsContent,omitempty" name:"SmsContent"`
	// 模板类型

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
	// 发送类型

	SendChannel *string `json:"SendChannel,omitempty" name:"SendChannel"`
	// 接收者

	Receiver *string `json:"Receiver,omitempty" name:"Receiver"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 无

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 创建时间. 格式: 2006-01-02 15:04:05

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间. 格式: 2006-01-02 15:04:05

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 接收者详情

	ReceiverObj []*ReceiverObj `json:"ReceiverObj,omitempty" name:"ReceiverObj"`
}

type ReceiverObj struct {
	// 主账号

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 账号状态

	Active *int64 `json:"Active,omitempty" name:"Active"`
	// AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 邮箱

	Mail *string `json:"Mail,omitempty" name:"Mail"`
	// 昵称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 电话号码

	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
	// 无

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// Uid

	Uid *int64 `json:"Uid,omitempty" name:"Uid"`
	// Uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 账号类型

	UinType *string `json:"UinType,omitempty" name:"UinType"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 无

	Account *string `json:"Account,omitempty" name:"Account"`
	// 子账号列表

	UinList []*TCloundAccount `json:"UinList,omitempty" name:"UinList"`
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

type GroupIDName struct {
	// 分组编号

	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`
	// 分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type Message2 struct {
	// 编号

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 发送状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 发送者

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 消息类型

	MsgType *string `json:"MsgType,omitempty" name:"MsgType"`
	// 创建时间. 格式: 2006-01-02 15:04:05

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 格式: 2006-01-02 15:04:05

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// TopicId

	MsgTopicId *int64 `json:"MsgTopicId,omitempty" name:"MsgTopicId"`
	// 来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 发送者

	Sender *string `json:"Sender,omitempty" name:"Sender"`
	// 接收者

	Receiver *string `json:"Receiver,omitempty" name:"Receiver"`
	// 抄送者

	CcReceiver *string `json:"CcReceiver,omitempty" name:"CcReceiver"`
	// 标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 内容

	Content *string `json:"Content,omitempty" name:"Content"`
}
