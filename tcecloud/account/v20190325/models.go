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

package v20190325

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type GetUserInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// UIN

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 用户名

		UserName *string `json:"UserName,omitempty" name:"UserName"`
		// 昵称

		Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
		// 主账号UIN

		OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
		// 是否为主账号

		IsOwner *int64 `json:"IsOwner,omitempty" name:"IsOwner"`
		// 资料是否审核通过

		CheckStatus *int64 `json:"CheckStatus,omitempty" name:"CheckStatus"`
		// 是否实名认证：0未认证，1已认证

		IsAuthenticate *int64 `json:"IsAuthenticate,omitempty" name:"IsAuthenticate"`
		// 邮箱是否审核通过

		MailStatus *int64 `json:"MailStatus,omitempty" name:"MailStatus"`
		// 邮箱

		Mail *string `json:"Mail,omitempty" name:"Mail"`
		// 手机号码

		PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
		// 用户指引标识位

		GuideBit *int64 `json:"GuideBit,omitempty" name:"GuideBit"`
		// 用户首次购买带外网IP的cvm设备的时间

		WanIpTime *string `json:"WanIpTime,omitempty" name:"WanIpTime"`
		// 标识外网是否受限

		WanRestrict *int64 `json:"WanRestrict,omitempty" name:"WanRestrict"`
		// 创建时间

		AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`
		// 修改时间

		ModTimestamp *string `json:"ModTimestamp,omitempty" name:"ModTimestamp"`
		// 线下审核状态

		OfflineStatus *int64 `json:"OfflineStatus,omitempty" name:"OfflineStatus"`
		// 业务信息

		BizInfo *string `json:"BizInfo,omitempty" name:"BizInfo"`
		// 来源平台

		SrcPlatform *string `json:"SrcPlatform,omitempty" name:"SrcPlatform"`
		// 是否测试用户

		IsTestUser *int64 `json:"IsTestUser,omitempty" name:"IsTestUser"`
		// 客户来源

		ClientFrom *string `json:"ClientFrom,omitempty" name:"ClientFrom"`
		// register refer

		Referer *string `json:"Referer,omitempty" name:"Referer"`
		// 是否注册成功

		IsRegSucc *bool `json:"IsRegSucc,omitempty" name:"IsRegSucc"`
		// 用户属性集合

		Attributes *int64 `json:"Attributes,omitempty" name:"Attributes"`
		// 否导入了即时通白名单

		Isprotect *int64 `json:"Isprotect,omitempty" name:"Isprotect"`
		// 用户指引

		IsSeeGuidelines *int64 `json:"IsSeeGuidelines,omitempty" name:"IsSeeGuidelines"`
		// 是否接收推广信息

		IsAcceptProMsg *int64 `json:"IsAcceptProMsg,omitempty" name:"IsAcceptProMsg"`
		// 部署模块

		DeployName *string `json:"DeployName,omitempty" name:"DeployName"`
		// 账号列表

		AccountList *string `json:"AccountList,omitempty" name:"AccountList"`
		// 账号类型

		AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`
		// 默认开发商

		DefaultOwner *int64 `json:"DefaultOwner,omitempty" name:"DefaultOwner"`
		// 国家代码

		CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
		// 邮箱验证

		MailVerify *int64 `json:"MailVerify,omitempty" name:"MailVerify"`
		// 接收信息语言

		MsgLang *string `json:"MsgLang,omitempty" name:"MsgLang"`
		// 地域

		Area *string `json:"Area,omitempty" name:"Area"`
		// 是否需要完善信息

		Needinfo *int64 `json:"Needinfo,omitempty" name:"Needinfo"`
		// 绑定账号

		Account *string `json:"Account,omitempty" name:"Account"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChangeMailPasswordRequest struct {
	*tchttp.BaseRequest

	// 旧密码

	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`
	// 新密码

	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 语言类型，zh或en

	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

func (r *ChangeMailPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeMailPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChangeSubAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// 旧密码

	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`
	// 新密码

	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`
	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
}

func (r *ChangeSubAccountPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeSubAccountPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSubLoginUinListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSubLoginUinListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubLoginUinListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttributeUserInfo struct {
	// uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户名称

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 昵称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
}

type DeleteTokenRequest struct {
	*tchttp.BaseRequest

	// 待删除token

	DelToken *string `json:"DelToken,omitempty" name:"DelToken"`
}

func (r *DeleteTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAccountLoginStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可登陆状态，0-正常，1-临时锁定，2-运营端锁定

		LoginStatus *int64 `json:"LoginStatus,omitempty" name:"LoginStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAccountLoginStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAccountLoginStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttributeNameRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAttributeNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttributeNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSafeAuthConfigRequest struct {
	*tchttp.BaseRequest

	// 查询的用户uin

	UserUin *uint64 `json:"UserUin,omitempty" name:"UserUin"`
}

func (r *GetSafeAuthConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSafeAuthConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountBankInfo struct {
	// 账号id

	Accountid *string `json:"Accountid,omitempty" name:"Accountid"`
	// 账号名称

	Accountname *string `json:"Accountname,omitempty" name:"Accountname"`
	// 账号所属银行

	Accountbank *string `json:"Accountbank,omitempty" name:"Accountbank"`
	// 省

	Provincename *string `json:"Provincename,omitempty" name:"Provincename"`
	// 省ID

	Provinceid *string `json:"Provinceid,omitempty" name:"Provinceid"`
	// 城市

	Cityname *string `json:"Cityname,omitempty" name:"Cityname"`
	// 城市ID

	Cityid *string `json:"Cityid,omitempty" name:"Cityid"`
	// 银行名称

	Bankname *string `json:"Bankname,omitempty" name:"Bankname"`
	// 银行ID

	Bankid *string `json:"Bankid,omitempty" name:"Bankid"`
}

type SetSingleLoginFlagRequest struct {
	*tchttp.BaseRequest

	// 登陆临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *SetSingleLoginFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetSingleLoginFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTokenRequest struct {
	*tchttp.BaseRequest

	// 过期时间，单位s

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

func (r *GetTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRelatedUinSessionKeyRequest struct {
	*tchttp.BaseRequest

	// 当前账号登录态

	SessionKey *string `json:"SessionKey,omitempty" name:"SessionKey"`
}

func (r *DescribeRelatedUinSessionKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRelatedUinSessionKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMultiFactorParasRequest struct {
	*tchttp.BaseRequest
}

func (r *GetMultiFactorParasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMultiFactorParasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAppIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSubAccountInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 子账户uin

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 是否是控制台登陆

		ConsoleLogin *int64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`
		// 是否需要重置密码。需要-1

		NeedResetPassword *int64 `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSubAccountInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubAccountInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppIdByOwnerUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAppIdByOwnerUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppIdByOwnerUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTradeRequest struct {
	*tchttp.BaseRequest
}

func (r *GetTradeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTradeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCaptchaRequest struct {
	*tchttp.BaseRequest

	// 验证码关联token

	CaptchaToken *string `json:"CaptchaToken,omitempty" name:"CaptchaToken"`
	// 新设置的图形验证码

	Captcha *string `json:"Captcha,omitempty" name:"Captcha"`
}

func (r *SetCaptchaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCaptchaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChangeSubAccountPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ChangeSubAccountPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeSubAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuserGetUserInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuserGetUserInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuserGetUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SafeAuthOffsiteFlag struct {
	// 验证标识

	VerifyFlag *int64 `json:"VerifyFlag,omitempty" name:"VerifyFlag"`
	// 是否进行电话通知

	NotifyPhone *int64 `json:"NotifyPhone,omitempty" name:"NotifyPhone"`
	// 是否进行

	NotifyEmail *int64 `json:"NotifyEmail,omitempty" name:"NotifyEmail"`
}

type GetAttributeValuesRequest struct {
	*tchttp.BaseRequest

	// 被查询的账户uin

	TargetUin *int64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *GetAttributeValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttributeValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendVerifyCodeRequest struct {
	*tchttp.BaseRequest

	// 需要验证的操作

	VerifyAction *string `json:"VerifyAction,omitempty" name:"VerifyAction"`
	// 语言

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 验证码类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 手机号码

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 手机区号

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 邮箱号码

	Mail *string `json:"Mail,omitempty" name:"Mail"`
	// 修改来源

	From *string `json:"From,omitempty" name:"From"`
	// 目标账号uin

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *SendVerifyCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendVerifyCodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUinOwnerInOpenRequest struct {
	*tchttp.BaseRequest
}

func (r *GetUinOwnerInOpenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUinOwnerInOpenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserAreaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 区域，1-中国

		Area *int64 `json:"Area,omitempty" name:"Area"`
		// 国家名称，CN-中国

		CountryName *string `json:"CountryName,omitempty" name:"CountryName"`
		// 国家代码，86-中国

		CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserAreaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserAreaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendVerifyCodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendVerifyCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendVerifyCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckSubAccountUinRequest struct {
	*tchttp.BaseRequest

	// 子账户uin列表

	AccountUin []*uint64 `json:"AccountUin,omitempty" name:"AccountUin"`
}

func (r *CheckSubAccountUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSubAccountUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSoftTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// token序列号

		TokenSn *string `json:"TokenSn,omitempty" name:"TokenSn"`
		// 链接数据，用于生成google authentication客户端绑定二维码，base64编码

		QrcodeData *string `json:"QrcodeData,omitempty" name:"QrcodeData"`
		// 过期时间戳，秒

		ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
		// token标识

		Mark *string `json:"Mark,omitempty" name:"Mark"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSoftTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSoftTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInfoByFieldsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetInfoByFieldsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInfoByFieldsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserByAttributeValueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询到的用户总数

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 用户信息列表

		UserInfo []*AttributeUserInfo `json:"UserInfo,omitempty" name:"UserInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserByAttributeValueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserByAttributeValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInfoByFieldsRequest struct {
	*tchttp.BaseRequest

	// 账户信息

	Account *AccountRequest `json:"Account,omitempty" name:"Account"`
	// 查询属性列表

	Fields []*string `json:"Fields,omitempty" name:"Fields"`
}

func (r *GetInfoByFieldsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInfoByFieldsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserProductUsableInfo struct {
	// 单位

	Uint *string `json:"Uint,omitempty" name:"Uint"`
	// 产品id

	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
	// 是否打开

	IsOpen *int64 `json:"IsOpen,omitempty" name:"IsOpen"`
	// 总量

	Nums *int64 `json:"Nums,omitempty" name:"Nums"`
	// DeadNums

	DeadNums *int64 `json:"DeadNums,omitempty" name:"DeadNums"`
	// 前缀

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 区域id

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

type CheckTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAttributeNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 属性id列表

		AttributeId []*uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddAttributeNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAttributeNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSubAccountRequest struct {
	*tchttp.BaseRequest

	// 用户信息

	UserInfo []*UserInfo `json:"UserInfo,omitempty" name:"UserInfo"`
	// 语言

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 是否在白名单

	InWhiteList *bool `json:"InWhiteList,omitempty" name:"InWhiteList"`
	// 从Api创建

	FromAPI *uint64 `json:"FromAPI,omitempty" name:"FromAPI"`
}

func (r *AddSubAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSubAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppIdByOwnerUinRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAppIdByOwnerUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppIdByOwnerUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSubAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账号Uin

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 子用户uin列表

		List []*uint64 `json:"List,omitempty" name:"List"`
		// 子用户Uid

		Uids []*uint64 `json:"Uids,omitempty" name:"Uids"`
		// 子用户详情

		SubAccounts []*SubAccounts `json:"SubAccounts,omitempty" name:"SubAccounts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddSubAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSubAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TokenUnBindResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TokenUnBindResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TokenUnBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckMailPasswordRequest struct {
	*tchttp.BaseRequest

	// 用户名

	Account *string `json:"Account,omitempty" name:"Account"`
	// 密码，密文

	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *CheckMailPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckMailPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckMailPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否正确

		IsRight *bool `json:"IsRight,omitempty" name:"IsRight"`
		// 账户唯一id

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckMailPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckMailPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMasterListRequest struct {
	*tchttp.BaseRequest
}

func (r *GetMasterListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMasterListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserProductUsableResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品信息列表

		List []*UserProductUsableInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserProductUsableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserProductUsableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubAccounts struct {
	// 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// Token

	Token *string `json:"Token,omitempty" name:"Token"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 秘钥Id

	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
	// 秘钥Key

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
}

type GetUinOwnerInOpenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUinOwnerInOpenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUinOwnerInOpenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MFAStatusRequest struct {
	*tchttp.BaseRequest

	// 接口名

	Interface *string `json:"Interface,omitempty" name:"Interface"`
	// 客户端代理信息

	ClientUA *string `json:"ClientUA,omitempty" name:"ClientUA"`
	// 登陆临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *MFAStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MFAStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountAttributeAndValue struct {
	// 属性名称

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 属性

	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`
	// 属性id

	AttributeId *int64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 值id

	ValueId *int64 `json:"ValueId,omitempty" name:"ValueId"`
	// uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type AccountExistData struct {
	// 账户唯一id

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户名

	Account *string `json:"Account,omitempty" name:"Account"`
	// 昵称

	Nick *string `json:"Nick,omitempty" name:"Nick"`
	// 账户状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 账户绑定类型

	TypeBind *int64 `json:"TypeBind,omitempty" name:"TypeBind"`
	// 是否为注册账户

	IsRegAccount *int64 `json:"IsRegAccount,omitempty" name:"IsRegAccount"`
	// 最近一次更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ActionLoginFlag struct {
	// 电话

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 软Token

	Stoken *string `json:"Stoken,omitempty" name:"Stoken"`
	// 硬Token

	Token *string `json:"Token,omitempty" name:"Token"`
}

type DeleteTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMultiFactorParasRespData struct {
	// 多因子选项

	MultiFactorChoices *int64 `json:"MultiFactorChoices,omitempty" name:"MultiFactorChoices"`
	// 多因子是否可改变

	MultiFactorcChangable *int64 `json:"MultiFactorcChangable,omitempty" name:"MultiFactorcChangable"`
	// 多因子操作

	MultiActionChoices *int64 `json:"MultiActionChoices,omitempty" name:"MultiActionChoices"`
	// 多因子操作是否可改变

	MultiActionChangable *int64 `json:"MultiActionChangable,omitempty" name:"MultiActionChangable"`
	// 多因子硬件认证来源

	MfaHardTokenSource *string `json:"MfaHardTokenSource,omitempty" name:"MfaHardTokenSource"`
}

type GetMultiFactorParasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回体Data数据结构

		Data *GetMultiFactorParasRespData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMultiFactorParasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMultiFactorParasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserIdAttrRequest struct {
	*tchttp.BaseRequest

	// uin列表

	UinArr []*uint64 `json:"UinArr,omitempty" name:"UinArr"`
	// appid列表

	AppIdArr []*uint64 `json:"AppIdArr,omitempty" name:"AppIdArr"`
}

func (r *GetUserIdAttrRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserIdAttrRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCountryCodeRequest struct {
	*tchttp.BaseRequest
}

func (r *GetCountryCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCountryCodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExtAttr struct {
	// 需要重置mfa的token

	NeedResetToken *int64 `json:"NeedResetToken,omitempty" name:"NeedResetToken"`
	// 需要重置mfa的token

	NeedResetStoken *int64 `json:"NeedResetStoken,omitempty" name:"NeedResetStoken"`
}

type TokenBindResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TokenBindResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TokenBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAccountLoginStatusRequest struct {
	*tchttp.BaseRequest

	// 用户名

	Account *string `json:"Account,omitempty" name:"Account"`
}

func (r *GetAccountLoginStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAccountLoginStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMasterListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主账户信息列表

		OwnerInfo []*OwnerInfo `json:"OwnerInfo,omitempty" name:"OwnerInfo"`
		// 用户id

		LoginUid *uint64 `json:"LoginUid,omitempty" name:"LoginUid"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMasterListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMasterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTradeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTradeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserAreaByLoginUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地区枚举值，1-中国，2-其他

		Area *int64 `json:"Area,omitempty" name:"Area"`
		// 国家名称，中国-CN

		CountryName *string `json:"CountryName,omitempty" name:"CountryName"`
		// 国家代码，中国-86

		CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserAreaByLoginUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserAreaByLoginUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckCaptchaRequest struct {
	*tchttp.BaseRequest

	// 验证码

	Captcha *string `json:"Captcha,omitempty" name:"Captcha"`
	// 验证生命周期用的token

	CaptchaToken *string `json:"CaptchaToken,omitempty" name:"CaptchaToken"`
}

func (r *CheckCaptchaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckCaptchaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppIdRequest struct {
	*tchttp.BaseRequest

	// 业务传过来的uin

	ParaUin *int64 `json:"ParaUin,omitempty" name:"ParaUin"`
}

func (r *GetAppIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBindAccountByUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户唯一id

		Fakeuin *uint64 `json:"Fakeuin,omitempty" name:"Fakeuin"`
		// 用户名

		Account *string `json:"Account,omitempty" name:"Account"`
		// 账户昵称

		Nick *string `json:"Nick,omitempty" name:"Nick"`
		// 账户绑定类型

		TypeBind *int64 `json:"TypeBind,omitempty" name:"TypeBind"`
		// 是否为注册账户

		IsRegAccount *int64 `json:"IsRegAccount,omitempty" name:"IsRegAccount"`
		// 账户状态

		BindStatus *int64 `json:"BindStatus,omitempty" name:"BindStatus"`
		// 类型名称

		TypeName *string `json:"TypeName,omitempty" name:"TypeName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBindAccountByUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBindAccountByUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetSafeAuthFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetSafeAuthFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetSafeAuthFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNicknameRequest struct {
	*tchttp.BaseRequest

	// 临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *GetNicknameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNicknameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMasterListWithStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主账户信息列表

		OwnerInfo []*OwnerInfo `json:"OwnerInfo,omitempty" name:"OwnerInfo"`
		// uid

		LoginUid *uint64 `json:"LoginUid,omitempty" name:"LoginUid"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMasterListWithStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMasterListWithStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetMfaDeviceRequest struct {
	*tchttp.BaseRequest

	// 1-三方设备, 2-虚拟mfa设备

	TokenType *int64 `json:"TokenType,omitempty" name:"TokenType"`
	// 被分配客户uin

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *SetMfaDeviceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetMfaDeviceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginVerifyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 昵称

		Nick *string `json:"Nick,omitempty" name:"Nick"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LoginVerifyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LoginVerifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SafeAuthFlagInfo struct {
	// 登陆标识

	LoginFlag *SafeAuthFlag `json:"LoginFlag,omitempty" name:"LoginFlag"`
	// action标识

	ActionFlag *SafeAuthFlag `json:"ActionFlag,omitempty" name:"ActionFlag"`
	// 外围标识

	OffsiteFlag *SafeAuthOffsiteFlag `json:"OffsiteFlag,omitempty" name:"OffsiteFlag"`
}

type CheckSubAccountNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存在的子账户用户名列表

		Exist []*string `json:"Exist,omitempty" name:"Exist"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckSubAccountNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSubAccountNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendLoginVerifyCodeRequest struct {
	*tchttp.BaseRequest

	// 登陆临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *SendLoginVerifyCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendLoginVerifyCodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetManageAccountAttributesRequest struct {
	*tchttp.BaseRequest
}

func (r *GetManageAccountAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetManageAccountAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserProductUsableRequest struct {
	*tchttp.BaseRequest

	// 传"0"

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *GetUserProductUsableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserProductUsableRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckSubAccountNameRequest struct {
	*tchttp.BaseRequest

	// 用户名列表

	Names []*string `json:"Names,omitempty" name:"Names"`
}

func (r *CheckSubAccountNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSubAccountNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckVerifyCodeRequest struct {
	*tchttp.BaseRequest

	// 验证码类型，1-手机验证，2-邮箱验证

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 验证码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 手机号

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 地区码，默认86

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 邮箱

	Mail *string `json:"Mail,omitempty" name:"Mail"`
}

func (r *CheckVerifyCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckVerifyCodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSoftTokenRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateSoftTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSoftTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新创建的token

		Token *string `json:"Token,omitempty" name:"Token"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountAttributeValue struct {
	// 属性

	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`
	// 属性名称

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 属性值

	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
	// 属性id

	AttributeId *int64 `json:"AttributeId,omitempty" name:"AttributeId"`
}

type SafeAuthTokenInfo struct {
	// 状态，0-未分配，1-已分配，2-已绑定，3-已解绑

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// token序列号

	TokenSn *string `json:"TokenSn,omitempty" name:"TokenSn"`
	// token类型，1-hard token, 2-soft token

	TokenType *int64 `json:"TokenType,omitempty" name:"TokenType"`
}

type CheckCaptchaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckCaptchaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckCaptchaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MFAStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否需要认证，0-不需要，1-需要

		NeedAuth *int64 `json:"NeedAuth,omitempty" name:"NeedAuth"`
		// 过期时间

		ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
		// 认证类型

		AuthType *int64 `json:"AuthType,omitempty" name:"AuthType"`
		// 认证的手机号

		AuthPhone *string `json:"AuthPhone,omitempty" name:"AuthPhone"`
		// 认证的token

		AuthHardtoken *string `json:"AuthHardtoken,omitempty" name:"AuthHardtoken"`
		// 是否来自小程序

		IsWeapp *int64 `json:"IsWeapp,omitempty" name:"IsWeapp"`
		// 认证的email

		AuthEmail *string `json:"AuthEmail,omitempty" name:"AuthEmail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MFAStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MFAStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendLoginVerifyCodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendLoginVerifyCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendLoginVerifyCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountAttribute struct {
	// id

	AttributeId *int64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 属性名称

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 属性

	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`
}

type GetUserInfoRequest struct {
	*tchttp.BaseRequest

	// 是否为主账号

	IsOwner *uint64 `json:"IsOwner,omitempty" name:"IsOwner"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 审核状态

	CheckStatus *int64 `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 邮箱是否审核通过

	MailStatus *int64 `json:"MailStatus,omitempty" name:"MailStatus"`
	// 线下审核状态

	OfflineStatus *int64 `json:"OfflineStatus,omitempty" name:"OfflineStatus"`
	// 首次购买带外网IP的cvm设备的时间

	WanIpTime *string `json:"WanIpTime,omitempty" name:"WanIpTime"`
	// 外网是否受限

	WanRestrict *int64 `json:"WanRestrict,omitempty" name:"WanRestrict"`
	// 返回的字段

	Fields *string `json:"Fields,omitempty" name:"Fields"`
}

func (r *GetUserInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClearLoginFlagRequest struct {
	*tchttp.BaseRequest

	// 登陆临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *ClearLoginFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClearLoginFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserAreaByLoginUinRequest struct {
	*tchttp.BaseRequest
}

func (r *GetUserAreaByLoginUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserAreaByLoginUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCaptchaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCaptchaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCaptchaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppIdByLoginUinRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAppIdByLoginUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppIdByLoginUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOwnerUinByAppidRequest struct {
	*tchttp.BaseRequest

	// appid

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
}

func (r *GetOwnerUinByAppidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOwnerUinByAppidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSubAccountInfoRequest struct {
	*tchttp.BaseRequest

	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 主账户uin

	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

func (r *GetSubAccountInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubAccountInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SeedLoginTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 生成的token字符串

		Token *string `json:"Token,omitempty" name:"Token"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SeedLoginTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SeedLoginTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserByAttributeValueRequest struct {
	*tchttp.BaseRequest

	// 扩展属性

	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`
	// 扩展属性值

	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
	// 分页每页大小，默认20

	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
	// 分页值，从1开始

	Page *uint64 `json:"Page,omitempty" name:"Page"`
}

func (r *GetUserByAttributeValueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserByAttributeValueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMasterListWithStatusRequest struct {
	*tchttp.BaseRequest

	// 客户端用户代理

	ClientUA *string `json:"ClientUA,omitempty" name:"ClientUA"`
	// 临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *GetMasterListWithStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMasterListWithStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClearLoginFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearLoginFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClearLoginFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSubLoginUinListRequest struct {
	*tchttp.BaseRequest
}

func (r *GetSubLoginUinListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSubLoginUinListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserIdAttrResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserIdAttrResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserIdAttrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountCurInfo struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 区域

	Area *int64 `json:"Area,omitempty" name:"Area"`
	// 证件类型

	Idcardtype *int64 `json:"Idcardtype,omitempty" name:"Idcardtype"`
	// 证件名称

	Idcard *string `json:"Idcard,omitempty" name:"Idcard"`
	// 组织代码

	Organizationcode *string `json:"Organizationcode,omitempty" name:"Organizationcode"`
	// 认证类型

	AuthenticateType *string `json:"AuthenticateType,omitempty" name:"AuthenticateType"`
}

type LoginMfa struct {
	// 是否需要认证

	NeedAuth *int64 `json:"NeedAuth,omitempty" name:"NeedAuth"`
	// 认证类型

	AuthType *int64 `json:"AuthType,omitempty" name:"AuthType"`
	// 认证手机号

	AuthPhone *string `json:"AuthPhone,omitempty" name:"AuthPhone"`
}

type SetAttributeValuesRequest struct {
	*tchttp.BaseRequest

	// 被设置的目标uin，不填则设置登录用户

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 属性值列表

	Attributes *AccountAttributeValue `json:"Attributes,omitempty" name:"Attributes"`
}

func (r *SetAttributeValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAttributeValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SafeAuthFlag struct {
	// 是否进行手机号认证，0-否，1-是

	Phone *int64 `json:"Phone,omitempty" name:"Phone"`
	// 是否进行token认证，0-否，1-是

	Token *int64 `json:"Token,omitempty" name:"Token"`
	// 是否进行stoken认证，0-否，1-是

	Stoken *int64 `json:"Stoken,omitempty" name:"Stoken"`
}

type SetMfaDeviceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetMfaDeviceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetMfaDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttributeValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 认证类型

		IdentifyType *int64 `json:"IdentifyType,omitempty" name:"IdentifyType"`
		// 扩展属性值

		Items []*AccountAttributeAndValue `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAttributeValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttributeValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckSubAccountUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存在的子账户uin列表

		Exist []*uint64 `json:"Exist,omitempty" name:"Exist"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckSubAccountUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckSubAccountUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLastLoginInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLastLoginInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLastLoginInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserAreaRequest struct {
	*tchttp.BaseRequest
}

func (r *GetUserAreaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserAreaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginVerifyRequest struct {
	*tchttp.BaseRequest

	// 登陆临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
	// 域名id

	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
	// 应答是否需要昵称

	NeedNick *int64 `json:"NeedNick,omitempty" name:"NeedNick"`
	// 客户端代理

	ClientUA *string `json:"ClientUA,omitempty" name:"ClientUA"`
}

func (r *LoginVerifyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LoginVerifyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserInfoByLoginUinRequest struct {
	*tchttp.BaseRequest
}

func (r *GetUserInfoByLoginUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserInfoByLoginUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserInfo struct {
	// 子账号类型

	CanLogin *string `json:"CanLogin,omitempty" name:"CanLogin"`
	// 区号

	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
	// 详情

	Detail *AccountDetail `json:"Detail,omitempty" name:"Detail"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 电话号码

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 系统类型

	SyStemType *string `json:"SyStemType,omitempty" name:"SyStemType"`
}

type CheckTokenRequest struct {
	*tchttp.BaseRequest

	// 待校验的token

	CheckToken *string `json:"CheckToken,omitempty" name:"CheckToken"`
}

func (r *CheckTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLastLoginInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *GetLastLoginInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLastLoginInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSafeAuthConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// token信息

		TokenInfo *SafeAuthTokenInfo `json:"TokenInfo,omitempty" name:"TokenInfo"`
		// 标识信息

		Flag *SafeAuthFlagInfo `json:"Flag,omitempty" name:"Flag"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSafeAuthConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSafeAuthConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetSafeAuthFlagRequest struct {
	*tchttp.BaseRequest

	// 用户uin

	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`
	// 登陆标识

	LoginFlag *SafeAuthFlag `json:"LoginFlag,omitempty" name:"LoginFlag"`
	// action标识

	ActionFlag *SafeAuthFlag `json:"ActionFlag,omitempty" name:"ActionFlag"`
}

func (r *SetSafeAuthFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetSafeAuthFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountRequest struct {
	// 账户类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 账户id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 主账户id

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 子账户id列表

	SubAccountList *string `json:"SubAccountList,omitempty" name:"SubAccountList"`
	// uid列表

	UidList *string `json:"UidList,omitempty" name:"UidList"`
	// 角色列表

	RoleList *string `json:"RoleList,omitempty" name:"RoleList"`
	// 用户组列表

	GroupList *string `json:"GroupList,omitempty" name:"GroupList"`
}

type SetLoginFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetLoginFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLoginFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TokenBindRequest struct {
	*tchttp.BaseRequest

	// token类型，1-hard token; 2- soft token

	TokenType *int64 `json:"TokenType,omitempty" name:"TokenType"`
	// 被设置的用户uin

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// token序列号，tokenType = soft token时，必填

	TokenSn *string `json:"TokenSn,omitempty" name:"TokenSn"`
	// 验证码，动态口令

	Code *string `json:"Code,omitempty" name:"Code"`
}

func (r *TokenBindRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TokenBindRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAttributeValuesRequest struct {
	*tchttp.BaseRequest

	// 被删除的目标uin，不填则删除登录账户

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 属性列表

	Attributes []*string `json:"Attributes,omitempty" name:"Attributes"`
}

func (r *DeleteAttributeValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAttributeValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRelatedUinSessionKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 关联账号

		RelatedUin *uint64 `json:"RelatedUin,omitempty" name:"RelatedUin"`
		// 关联账号的主账号

		RelatedOwnerUin *uint64 `json:"RelatedOwnerUin,omitempty" name:"RelatedOwnerUin"`
		// 关联登录态

		RelatedSessionKey *string `json:"RelatedSessionKey,omitempty" name:"RelatedSessionKey"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRelatedUinSessionKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRelatedUinSessionKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetLoginFlagRequest struct {
	*tchttp.BaseRequest

	// 登陆临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *SetLoginFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetLoginFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetSingleLoginFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetSingleLoginFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetSingleLoginFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAttributeValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAttributeValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAttributeValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNicknameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 昵称

		Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
		// 显示的名称

		DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNicknameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNicknameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNicknameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNicknameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNicknameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckVerifyCodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// token

		Token *string `json:"Token,omitempty" name:"Token"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckVerifyCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckVerifyCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNicknameRequest struct {
	*tchttp.BaseRequest

	// 用户新昵称

	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
}

func (r *ModifyNicknameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNicknameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAccountExistRequest struct {
	*tchttp.BaseRequest

	// 用户名

	Account *string `json:"Account,omitempty" name:"Account"`
}

func (r *CheckAccountExistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAccountExistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAttributeNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 属性列表

		Attributes []*AccountAttribute `json:"Attributes,omitempty" name:"Attributes"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAttributeNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAttributeNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddAttributeNameRequest struct {
	*tchttp.BaseRequest

	// 属性信息

	Attributes []*AccountAttribute `json:"Attributes,omitempty" name:"Attributes"`
}

func (r *AddAttributeNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAttributeNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAccountExistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户唯一id

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// 用户名

		Account *string `json:"Account,omitempty" name:"Account"`
		// 昵称

		Nick *string `json:"Nick,omitempty" name:"Nick"`
		// 账户状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 账户绑定类型

		TypeBind *int64 `json:"TypeBind,omitempty" name:"TypeBind"`
		// 是否为注册账户

		IsRegAccount *int64 `json:"IsRegAccount,omitempty" name:"IsRegAccount"`
		// 最近一次更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAccountExistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAccountExistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOwnerUinByAppidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主账户id

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// appid

		AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
		// 账户创建时间

		AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOwnerUinByAppidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOwnerUinByAppidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserInfoByLoginUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户UIN

		Uin *string `json:"Uin,omitempty" name:"Uin"`
		// 主账号UIN

		Owneruin *int64 `json:"Owneruin,omitempty" name:"Owneruin"`
		// 登陆账号

		Account *string `json:"Account,omitempty" name:"Account"`
		// 账号类型名称

		Usertype *string `json:"Usertype,omitempty" name:"Usertype"`
		// 账号类型

		Type *int64 `json:"Type,omitempty" name:"Type"`
		// 区域编码

		Area *int64 `json:"Area,omitempty" name:"Area"`
		// 当前区域

		Curarea *int64 `json:"Curarea,omitempty" name:"Curarea"`
		// 邮件验证状态

		Mailverify *int64 `json:"Mailverify,omitempty" name:"Mailverify"`
		// 证件名称

		Idcard *string `json:"Idcard,omitempty" name:"Idcard"`
		// 证件类型

		Idcardtype *int64 `json:"Idcardtype,omitempty" name:"Idcardtype"`
		// 当前证件类型

		Curidcardtype *int64 `json:"Curidcardtype,omitempty" name:"Curidcardtype"`
		// 证件url

		Idcardurl *string `json:"Idcardurl,omitempty" name:"Idcardurl"`
		// 证件uuid

		Idcarduuid *string `json:"Idcarduuid,omitempty" name:"Idcarduuid"`
		// 证件缓存地址

		Idcardcacheurl *string `json:"Idcardcacheurl,omitempty" name:"Idcardcacheurl"`
		// 证件缓存uuid

		Idcardcacheuuid *string `json:"Idcardcacheuuid,omitempty" name:"Idcardcacheuuid"`
		// 证件新url

		Idcardurlnew *string `json:"Idcardurlnew,omitempty" name:"Idcardurlnew"`
		// 证件缓存新uuid

		Idcardcacheuuidnew *string `json:"Idcardcacheuuidnew,omitempty" name:"Idcardcacheuuidnew"`
		// 证件缓存新url

		Entrycardcacheurlnew *string `json:"Entrycardcacheurlnew,omitempty" name:"Entrycardcacheurlnew"`
		// 入境卡新url

		Entrycardurlnew *string `json:"Entrycardurlnew,omitempty" name:"Entrycardurlnew"`
		// 联系信息

		Contact *string `json:"Contact,omitempty" name:"Contact"`
		// 电话

		Tel *string `json:"Tel,omitempty" name:"Tel"`
		// 国家编码

		CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
		// 需求信息

		Needinfo *int64 `json:"Needinfo,omitempty" name:"Needinfo"`
		// 邮箱

		Mail *string `json:"Mail,omitempty" name:"Mail"`
		// 当前邮箱

		Curmail *string `json:"Curmail,omitempty" name:"Curmail"`
		// 邮箱验证状态

		Mailpass *int64 `json:"Mailpass,omitempty" name:"Mailpass"`
		// 当前邮箱验证状态

		Curmailpass *int64 `json:"Curmailpass,omitempty" name:"Curmailpass"`
		// 注册状态

		Registerstatus *string `json:"Registerstatus,omitempty" name:"Registerstatus"`
		// 当前开发检查状态

		Curdevcheckstate *int64 `json:"Curdevcheckstate,omitempty" name:"Curdevcheckstate"`
		// 开发检查状态

		Devcheckpass *int64 `json:"Devcheckpass,omitempty" name:"Devcheckpass"`
		// 银行账号

		Bankaccount *string `json:"Bankaccount,omitempty" name:"Bankaccount"`
		// 银行号码

		Banknumber *string `json:"Banknumber,omitempty" name:"Banknumber"`
		// 完成进度

		Perfection *int64 `json:"Perfection,omitempty" name:"Perfection"`
		// 协作者名称

		Collname *string `json:"Collname,omitempty" name:"Collname"`
		// 协作者手机号码

		Colltel *string `json:"Colltel,omitempty" name:"Colltel"`
		// 协作者邮箱

		Collmail *string `json:"Collmail,omitempty" name:"Collmail"`
		// wan ip时间

		WanIpTime *string `json:"WanIpTime,omitempty" name:"WanIpTime"`
		// wan限制

		WanRestrict *int64 `json:"WanRestrict,omitempty" name:"WanRestrict"`
		// 授权

		Accredit *string `json:"Accredit,omitempty" name:"Accredit"`
		// 当前授权

		Curaccredit *string `json:"Curaccredit,omitempty" name:"Curaccredit"`
		// 地址

		Addr *string `json:"Addr,omitempty" name:"Addr"`
		// 地址前缀

		Addrprefix *string `json:"Addrprefix,omitempty" name:"Addrprefix"`
		// 开发检查信息

		Devcheckmsg *string `json:"Devcheckmsg,omitempty" name:"Devcheckmsg"`
		// 当前开发检查信息

		Curdevcheckmsg *string `json:"Curdevcheckmsg,omitempty" name:"Curdevcheckmsg"`
		// 组织代码

		Organizationcode *string `json:"Organizationcode,omitempty" name:"Organizationcode"`
		// 当前组织代码

		Curorganizationcode *string `json:"Curorganizationcode,omitempty" name:"Curorganizationcode"`
		// 认证类型

		AuthenticateType *string `json:"AuthenticateType,omitempty" name:"AuthenticateType"`
		// 当前认证类型

		CurauthenticateType *string `json:"CurauthenticateType,omitempty" name:"CurauthenticateType"`
		// 是否修改

		IsModify *string `json:"IsModify,omitempty" name:"IsModify"`
		// 当前名称

		Curname *string `json:"Curname,omitempty" name:"Curname"`
		// 当前身份证件

		Curidcard *string `json:"Curidcard,omitempty" name:"Curidcard"`
		// 当前类型

		Curtype *int64 `json:"Curtype,omitempty" name:"Curtype"`
		// 账号来源

		SrcPlatform *string `json:"SrcPlatform,omitempty" name:"SrcPlatform"`
		// 业务信息

		BizInfo *string `json:"BizInfo,omitempty" name:"BizInfo"`
		// 开发检查时间

		Devchecktime *string `json:"Devchecktime,omitempty" name:"Devchecktime"`
		// 邮件检查时间

		Mailchecktime *string `json:"Mailchecktime,omitempty" name:"Mailchecktime"`
		// 开发首次检查时间

		Firstdevchecktime *string `json:"Firstdevchecktime,omitempty" name:"Firstdevchecktime"`
		// 接受开放协议

		Acceptopenprotocol *int64 `json:"Acceptopenprotocol,omitempty" name:"Acceptopenprotocol"`
		// 接受云协议

		Acceptyunprotocol *int64 `json:"Acceptyunprotocol,omitempty" name:"Acceptyunprotocol"`
		// 开放平台用户

		Openuser *int64 `json:"Openuser,omitempty" name:"Openuser"`
		// 接受开放平台

		Acceptopen *string `json:"Acceptopen,omitempty" name:"Acceptopen"`
		// 通过审核银行账号信息

		PassedBankInfo *AccountBankInfo `json:"PassedBankInfo,omitempty" name:"PassedBankInfo"`
		// 待审核银行账号信息

		NeedCheckBankInfo *AccountBankInfo `json:"NeedCheckBankInfo,omitempty" name:"NeedCheckBankInfo"`
		// 银行账号状态

		BankStatus *string `json:"BankStatus,omitempty" name:"BankStatus"`
		// 银行账号审核信息

		BankMsg *string `json:"BankMsg,omitempty" name:"BankMsg"`
		// 部署名称

		DeployName *string `json:"DeployName,omitempty" name:"DeployName"`
		// 存在子账号

		SubExist *int64 `json:"SubExist,omitempty" name:"SubExist"`
		// 昵称

		Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
		// 账号名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 消息语言

		MsgLang *string `json:"MsgLang,omitempty" name:"MsgLang"`
		// 账号当前身份信息

		Curinfo *AccountCurInfo `json:"Curinfo,omitempty" name:"Curinfo"`
		// 认证方式

		Authmethod *int64 `json:"Authmethod,omitempty" name:"Authmethod"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserInfoByLoginUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserInfoByLoginUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMasterListV2Request struct {
	*tchttp.BaseRequest

	// 是否获取账户属性

	WithAttr *int64 `json:"WithAttr,omitempty" name:"WithAttr"`
}

func (r *GetMasterListV2Request) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMasterListV2Request) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OwnerInfo struct {
	// 账户唯一id

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 显示名称

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 账户状态

	CheckStatus *int64 `json:"CheckStatus,omitempty" name:"CheckStatus"`
	// 用户名

	Account *string `json:"Account,omitempty" name:"Account"`
	// 是否为该子账户的默认主账户

	IsDefaultOwner *bool `json:"IsDefaultOwner,omitempty" name:"IsDefaultOwner"`
	// 手机号，打掩码处理

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 邮箱，掩码处理

	Email *string `json:"Email,omitempty" name:"Email"`
	// 扩展属性

	ExtAttr *ExtAttr `json:"ExtAttr,omitempty" name:"ExtAttr"`
	// 多因子认证

	LoginMfa *LoginMfa `json:"LoginMfa,omitempty" name:"LoginMfa"`
}

type SetAttributeValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetAttributeValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAttributeValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChangeMailPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ChangeMailPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ChangeMailPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppIdByLoginUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// appid

		AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAppIdByLoginUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppIdByLoginUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCountryCodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCountryCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCountryCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountDetail struct {
	// 敏感操作标识

	ActionFlag *ActionLoginFlag `json:"ActionFlag,omitempty" name:"ActionFlag"`
	// 是否允许控制台登录

	ConsoleLogin *string `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`
	// 登录保护

	LoginFlag *ActionLoginFlag `json:"LoginFlag,omitempty" name:"LoginFlag"`
	// 是否需要重置密码

	NeedResetPassword *string `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`
	// 用户密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 使用Api

	UseApi *string `json:"UseApi,omitempty" name:"UseApi"`
}

type GetMasterListV2Response struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主账户列表

		OwnerInfo []*OwnerInfo `json:"OwnerInfo,omitempty" name:"OwnerInfo"`
		// 账户uid

		LoginUid *uint64 `json:"LoginUid,omitempty" name:"LoginUid"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMasterListV2Response) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMasterListV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SeedLoginTokenRequest struct {
	*tchttp.BaseRequest

	// 登陆期临时密钥

	Skey *string `json:"Skey,omitempty" name:"Skey"`
	// 登陆临时id

	Tinyid *string `json:"Tinyid,omitempty" name:"Tinyid"`
}

func (r *SeedLoginTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SeedLoginTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TokenUnBindRequest struct {
	*tchttp.BaseRequest

	// 解绑目标账号

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *TokenUnBindRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TokenUnBindRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetManageAccountAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账号扩展属性列表

		Items []*AccountAttribute `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetManageAccountAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetManageAccountAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBindAccountByUinRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryBindAccountByUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBindAccountByUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuserGetUserInfoRequest struct {
	*tchttp.BaseRequest

	// 查询字段，字段名逗号分隔

	Fields *string `json:"Fields,omitempty" name:"Fields"`
}

func (r *QuserGetUserInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuserGetUserInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
