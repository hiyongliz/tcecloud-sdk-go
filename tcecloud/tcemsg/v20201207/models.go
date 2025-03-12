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
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type TestVoiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求数据

		Request *string `json:"Request,omitempty" name:"Request"`
		// 返回数据

		Response *string `json:"Response,omitempty" name:"Response"`
		// 错误信息

		ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
		// 状态(1-成功,2-失败)

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TestVoiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestVoiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveSendChannelRequest struct {
	*tchttp.BaseRequest

	// id，新增时可为空，修改时必填

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 发送类型 1-标准形式，2-自定义网关，3-plugin模式

	SendType *int64 `json:"SendType,omitempty" name:"SendType"`
	// 渠道编号（0-站内信，1-邮件，2-短信，4-语音，5-企业微信，其余为第三方渠道）

	ChannelNumber *int64 `json:"ChannelNumber,omitempty" name:"ChannelNumber"`
	// 渠道名称

	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`
	// 标准SMTP配置参数

	StandardSMTP *StandardSMTP `json:"StandardSMTP,omitempty" name:"StandardSMTP"`
	// 标准短信配置参数

	StandardSMS *StandardSMS `json:"StandardSMS,omitempty" name:"StandardSMS"`
	// 标准语音配置参数

	StandardVoice *StandardVoice `json:"StandardVoice,omitempty" name:"StandardVoice"`
	// 自定义http配置参数

	CustomHttp *CustomHttp `json:"CustomHttp,omitempty" name:"CustomHttp"`
	// 插件配置参数

	Plugin *Plugin `json:"Plugin,omitempty" name:"Plugin"`
}

func (r *SaveSendChannelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveSendChannelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TDefineSmsChannel struct {
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// Updated By

	UpdateName *string `json:"UpdateName,omitempty" name:"UpdateName"`
	// Creation time.

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// Update time.

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// Whether to enable channel number.

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// Channel number.

	Channel *string `json:"Channel,omitempty" name:"Channel"`
	// Product name.

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// The Chinese name of the product.

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type TestEmailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求数据

		Request *string `json:"Request,omitempty" name:"Request"`
		// 返回数据

		Response *string `json:"Response,omitempty" name:"Response"`
		// 错误信息

		ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
		// 状态(1-成功,2-失败)

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TestEmailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 字段名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段值列表

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type TestSMSResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 发送请求参数

		Request *string `json:"Request,omitempty" name:"Request"`
		// 发送返回参数

		Response *string `json:"Response,omitempty" name:"Response"`
		// 发送错误信息

		ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
		// 状态(1-成功,2-失败)

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TestSMSResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestSMSResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TestSMSRequest struct {
	*tchttp.BaseRequest

	// 手机号

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// id，新增时可为空，修改时必填

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 发送类型 1-标准形式，2-自定义网关，3-plugin模式

	SendType *int64 `json:"SendType,omitempty" name:"SendType"`
	// 渠道编号（0-站内信，1-邮件，2-短信，4-语音，5-企业微信，其余为第三方渠道）

	ChannelNumber *int64 `json:"ChannelNumber,omitempty" name:"ChannelNumber"`
	// 渠道名称

	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`
	// 标准SMTP配置参数

	StandardSMTP *StandardSMTP `json:"StandardSMTP,omitempty" name:"StandardSMTP"`
	// 标准短信配置参数

	StandardSMS *StandardSMS `json:"StandardSMS,omitempty" name:"StandardSMS"`
	// 标准语音配置参数

	StandardVoice *StandardVoice `json:"StandardVoice,omitempty" name:"StandardVoice"`
	// 自定义http配置参数

	CustomHttp *CustomHttp `json:"CustomHttp,omitempty" name:"CustomHttp"`
	// 插件配置参数

	Plugin *Plugin `json:"Plugin,omitempty" name:"Plugin"`
}

func (r *TestSMSRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestSMSRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDefineSmsChannelRequest struct {
	*tchttp.BaseRequest

	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤字段

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryDefineSmsChannelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDefineSmsChannelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomHttp struct {
	// 地址

	Url *string `json:"Url,omitempty" name:"Url"`
	// 是否对url进行编码

	UrlEncode *bool `json:"UrlEncode,omitempty" name:"UrlEncode"`
	// http 请求头信息

	Headers []*string `json:"Headers,omitempty" name:"Headers"`
	// 请求内容编码格式(utf8、gbk、gb2312))

	Encoding *string `json:"Encoding,omitempty" name:"Encoding"`
	// 签名算法 (none-无需签名，EXPR-表达式进行签名，MD5，SHA1，SHA256，SHA512，HMACMD5，HMACSHA1，HMACSHA256，HMACSHA512)

	SignatureMethod *string `json:"SignatureMethod,omitempty" name:"SignatureMethod"`
	// 签名编码（hex-16进制，base64-base64编码）

	SignatureEncoding *string `json:"SignatureEncoding,omitempty" name:"SignatureEncoding"`
	// 请求参数（get-与url进行拼接，post-直接发送）

	Params *string `json:"Params,omitempty" name:"Params"`
	// 签名字符串（可为表达式，也可为直接拼接数据）

	SignatureStr *string `json:"SignatureStr,omitempty" name:"SignatureStr"`
	// 超时时间，单位s

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	// 成功后返回数据包含字符串

	SuccessContains *string `json:"SuccessContains,omitempty" name:"SuccessContains"`
	// get、post

	Method *string `json:"Method,omitempty" name:"Method"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 租户端用户标识

	ConsoleReceiverTag *string `json:"ConsoleReceiverTag,omitempty" name:"ConsoleReceiverTag"`
	// 运营端用户标识

	CmgtReceiverTag *string `json:"CmgtReceiverTag,omitempty" name:"CmgtReceiverTag"`
}

type SaveSendChannelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// id

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveSendChannelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveSendChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDefineSmsChannelRequest struct {
	*tchttp.BaseRequest

	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 通道号

	Channel *string `json:"Channel,omitempty" name:"Channel"`
	// 启用状态 0默认、1自定义

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *UpdateDefineSmsChannelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDefineSmsChannelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDefineSmsChannelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDefineSmsChannelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDefineSmsChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMsgQuotaRequest struct {
	*tchttp.BaseRequest

	// 过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryMsgQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryMsgQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TestThirdPartyRequest struct {
	*tchttp.BaseRequest

	// 接收人

	Receiver *string `json:"Receiver,omitempty" name:"Receiver"`
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 发送类型 1-标准形式，2-自定义网关，3-plugin模式

	SendType *int64 `json:"SendType,omitempty" name:"SendType"`
	// 渠道编号（0-站内信，1-邮件，2-短信，4-语音，5-企业微信，其余为第三方渠道）

	ChannelNumber *int64 `json:"ChannelNumber,omitempty" name:"ChannelNumber"`
	// 渠道名称

	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`
	// 标准SMTP配置参数

	StandardSMTP *StandardSMTP `json:"StandardSMTP,omitempty" name:"StandardSMTP"`
	// 标准短信配置参数

	StandardSMS *StandardSMS `json:"StandardSMS,omitempty" name:"StandardSMS"`
	// 标准语音配置参数

	StandardVoice *StandardVoice `json:"StandardVoice,omitempty" name:"StandardVoice"`
	// 自定义http配置参数

	CustomHttp *CustomHttp `json:"CustomHttp,omitempty" name:"CustomHttp"`
	// 插件配置参数

	Plugin *Plugin `json:"Plugin,omitempty" name:"Plugin"`
}

func (r *TestThirdPartyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestThirdPartyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StandardSMTP struct {
	// host

	Host *string `json:"Host,omitempty" name:"Host"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 发件人邮箱地址

	Sender *string `json:"Sender,omitempty" name:"Sender"`
	// 是否使用SSL认证

	IsSSL *bool `json:"IsSSL,omitempty" name:"IsSSL"`
	// Desc

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 发件人名称

	SenderName *string `json:"SenderName,omitempty" name:"SenderName"`
}

type UpdateMsgQuotaRequest struct {
	*tchttp.BaseRequest

	// ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 配额大小

	Quota *int64 `json:"Quota,omitempty" name:"Quota"`
	// 限制状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *UpdateMsgQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMsgQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySendChannelDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Id

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 发送类型 1-标准形式，2-自定义网关，3-plugin模式

		SendType *int64 `json:"SendType,omitempty" name:"SendType"`
		// 渠道编号（0-站内信，1-邮件，2-短信，4-语音，5-企业微信，其他为第三方渠道）

		ChannelNumber *int64 `json:"ChannelNumber,omitempty" name:"ChannelNumber"`
		// 渠道名称

		ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`
		// 标准邮件配置

		StandardSMTP *StandardSMTP `json:"StandardSMTP,omitempty" name:"StandardSMTP"`
		// 标准短信配置

		StandardSMS *StandardSMS `json:"StandardSMS,omitempty" name:"StandardSMS"`
		// 标准语音配置

		StandardVoice *StandardVoice `json:"StandardVoice,omitempty" name:"StandardVoice"`
		// 自定义http网关配置

		CustomHttp *CustomHttp `json:"CustomHttp,omitempty" name:"CustomHttp"`
		// 插件配置

		Plugin *Plugin `json:"Plugin,omitempty" name:"Plugin"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySendChannelDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySendChannelDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMsgQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryMsgQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryMsgQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySendChannelDetailRequest struct {
	*tchttp.BaseRequest

	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *QuerySendChannelDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySendChannelDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDefineSmsChannelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据

		Data []*TDefineSmsChannel `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDefineSmsChannelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDefineSmsChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StandardVoice struct {
	// SecretId

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// SecretKey

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
	// sdk appid

	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`
	// TemplateId

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 请就近填写，默认ap-guangzhou，具体枚举请参考https://cloud.tencent.com/document/api/382/52071#.E4.BD.BF.E7.94.A8.E7.AD.BE.E5.90.8D.E6.96.B9.E6.B3.95-v3-.E7.9A.84.E5.85.AC.E5.85.B1.E5.8F.82.E6.95.B0

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// region

	Region *string `json:"Region,omitempty" name:"Region"`
}

type TestThirdPartyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求参数

		Request *string `json:"Request,omitempty" name:"Request"`
		// 返回参数

		Response *string `json:"Response,omitempty" name:"Response"`
		// 错误信息

		ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
		// 状态(1-成功,2-失败)

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TestThirdPartyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestThirdPartyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMsgQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMsgQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMsgQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StandardSMS struct {
	// SecretId

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// SecretKey

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
	// SdkAppId

	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`
	// 模版ID

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 短信签名

	Sign *string `json:"Sign,omitempty" name:"Sign"`
	// 请就近填写，默认ap-guangzhou，具体枚举请参考https://cloud.tencent.com/document/api/382/52071#.E4.BD.BF.E7.94.A8.E7.AD.BE.E5.90.8D.E6.96.B9.E6.B3.95-v3-.E7.9A.84.E5.85.AC.E5.85.B1.E5.8F.82.E6.95.B0

	Region *string `json:"Region,omitempty" name:"Region"`
}

type TestEmailRequest struct {
	*tchttp.BaseRequest

	// 邮件地址

	Email *string `json:"Email,omitempty" name:"Email"`
	// id，新增时可为空，修改时必填

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 发送类型 1-标准形式，2-自定义网关，3-plugin模式

	SendType *int64 `json:"SendType,omitempty" name:"SendType"`
	// 渠道编号（0-站内信，1-邮件，2-短信，4-语音，5-企业微信，其余为第三方渠道）

	ChannelNumber *int64 `json:"ChannelNumber,omitempty" name:"ChannelNumber"`
	// 渠道名称

	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`
	// 标准SMTP配置参数

	StandardSMTP *StandardSMTP `json:"StandardSMTP,omitempty" name:"StandardSMTP"`
	// 标准短信配置参数

	StandardSMS *StandardSMS `json:"StandardSMS,omitempty" name:"StandardSMS"`
	// 标准语音配置参数

	StandardVoice *StandardVoice `json:"StandardVoice,omitempty" name:"StandardVoice"`
	// 自定义http配置参数

	CustomHttp *CustomHttp `json:"CustomHttp,omitempty" name:"CustomHttp"`
	// 插件配置参数

	Plugin *Plugin `json:"Plugin,omitempty" name:"Plugin"`
}

func (r *TestEmailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestEmailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TestVoiceRequest struct {
	*tchttp.BaseRequest

	// 手机号

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// id，新增时可为空，修改时必填

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 发送类型 1-标准形式，2-自定义网关，3-plugin模式

	SendType *int64 `json:"SendType,omitempty" name:"SendType"`
	// 渠道编号（0-站内信，1-邮件，2-短信，4-语音，5-企业微信，其余为第三方渠道）

	ChannelNumber *int64 `json:"ChannelNumber,omitempty" name:"ChannelNumber"`
	// 渠道名称

	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`
	// 标准SMTP配置参数

	StandardSMTP *StandardSMTP `json:"StandardSMTP,omitempty" name:"StandardSMTP"`
	// 标准短信配置参数

	StandardSMS *StandardSMS `json:"StandardSMS,omitempty" name:"StandardSMS"`
	// 标准语音配置参数

	StandardVoice *StandardVoice `json:"StandardVoice,omitempty" name:"StandardVoice"`
	// 自定义http配置参数

	CustomHttp *CustomHttp `json:"CustomHttp,omitempty" name:"CustomHttp"`
	// 插件配置参数

	Plugin *Plugin `json:"Plugin,omitempty" name:"Plugin"`
}

func (r *TestVoiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TestVoiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Plugin struct {
	// 插件地址

	PluginUrl *string `json:"PluginUrl,omitempty" name:"PluginUrl"`
	// 租户端用户标识

	ConsoleReceiverTag *string `json:"ConsoleReceiverTag,omitempty" name:"ConsoleReceiverTag"`
	// 运营端用户标识

	CmgtReceiverTag *string `json:"CmgtReceiverTag,omitempty" name:"CmgtReceiverTag"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}
