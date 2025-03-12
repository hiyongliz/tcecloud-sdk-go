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

package v20210830

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type AgentPushTaskDetail struct {
	// 无

	Arg []*string `json:"Arg,omitempty" name:"Arg"`
	// 无

	Caller *string `json:"Caller,omitempty" name:"Caller"`
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	CustomArg *string `json:"CustomArg,omitempty" name:"CustomArg"`
	// 无

	ExecSucCnt *int64 `json:"ExecSucCnt,omitempty" name:"ExecSucCnt"`
	// 无

	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	MaxRetry *int64 `json:"MaxRetry,omitempty" name:"MaxRetry"`
	// 无

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 无

	OptionType *int64 `json:"OptionType,omitempty" name:"OptionType"`
	// 无

	PushSucCnt *int64 `json:"PushSucCnt,omitempty" name:"PushSucCnt"`
	// 无

	RetryInterval *int64 `json:"RetryInterval,omitempty" name:"RetryInterval"`
	// 无

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 无

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 无

	UuidCnt *int64 `json:"UuidCnt,omitempty" name:"UuidCnt"`
}

type CreateRemoteUpdateUserRequest struct {
	*tchttp.BaseRequest

	// 该用户当前的配置Id

	ConfigIds []*uint64 `json:"ConfigIds,omitempty" name:"ConfigIds"`
	// 用户信息
	//
	// *注* 创建时UserInfo中的Id字段后端会忽略掉

	UserInfo *RemoteUpdateUser `json:"UserInfo,omitempty" name:"UserInfo"`
}

func (r *CreateRemoteUpdateUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRemoteUpdateUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RspSocTamperProofList struct {
	// 运营防篡改事件目录

	EventPath *string `json:"EventPath,omitempty" name:"EventPath"`
	// 事件状态

	EventStatus *uint64 `json:"EventStatus,omitempty" name:"EventStatus"`
	// 篡改事件时间

	EventTamperTime *string `json:"EventTamperTime,omitempty" name:"EventTamperTime"`
	// 事件类型

	EventType *uint64 `json:"EventType,omitempty" name:"EventType"`
	// 运营防篡改ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 操作状态

	OptStatus *uint64 `json:"OptStatus,omitempty" name:"OptStatus"`
	// 防护开启时间

	ProtectOpenTime *string `json:"ProtectOpenTime,omitempty" name:"ProtectOpenTime"`
	// 防护状态

	ProtectStatus *uint64 `json:"ProtectStatus,omitempty" name:"ProtectStatus"`
	// 恢复开关

	RecoverySwitch *uint64 `json:"RecoverySwitch,omitempty" name:"RecoverySwitch"`
	// 运营防篡改服务器名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 运营防篡改ID

	SocTamperProofId *uint64 `json:"SocTamperProofId,omitempty" name:"SocTamperProofId"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ExportFileTamperEventsRequest struct {
	*tchttp.BaseRequest

	// By

	By *string `json:"By,omitempty" name:"By"`
	// 支持条件：AppId、Level、Status、Quuids、Uuids、FileAction、CreateBeginTime、CreateEndTime、ModifyBeginTime、ModifyEndTime、RuleName、HostIp、RuleCategory

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportFileTamperEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportFileTamperEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBashEventsStatusRequest struct {
	*tchttp.BaseRequest

	// 无

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// MarkStatus：1 2

	MarkStatus *int64 `json:"MarkStatus,omitempty" name:"MarkStatus"`
}

func (r *ModifyBashEventsStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBashEventsStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebServiceProcessListRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetWebServiceProcessListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebServiceProcessListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwareEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*EventsMalware `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMalwareEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwareEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpdateGrayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// id、name

		List []*UpdateGray `json:"List,omitempty" name:"List"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpdateGrayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateGrayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportRemoteUpdateConfigRequest struct {
	*tchttp.BaseRequest

	// 导入的配置文件的COS对象路径

	CosObject *string `json:"CosObject,omitempty" name:"CosObject"`
}

func (r *ImportRemoteUpdateConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportRemoteUpdateConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckFileStatusRequest struct {
	*tchttp.BaseRequest

	// 文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

func (r *CheckFileStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckFileStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVulTypeRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteVulTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVulTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebFrameListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 无

		WebFrames []*AssetWebFrameBaseInfo `json:"WebFrames,omitempty" name:"WebFrames"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebFrameListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebFrameListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineHostDetect struct {
	// 0:未通过 1:忽略 3:通过 5:检测中

	DetectStatus *int64 `json:"DetectStatus,omitempty" name:"DetectStatus"`
	// 首次检测时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 主机Id

	HostId *string `json:"HostId,omitempty" name:"HostId"`
	// 内网Ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 关联检测项数

	ItemCount *int64 `json:"ItemCount,omitempty" name:"ItemCount"`
	// 最后检测时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 检测未通过数

	NotPassedItemCount *int64 `json:"NotPassedItemCount,omitempty" name:"NotPassedItemCount"`
	// 检测通过数

	PassedItemCount *int64 `json:"PassedItemCount,omitempty" name:"PassedItemCount"`
	// 主机安全UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 外网Ip

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
}

type DescribeUpdateTaskLogRequest struct {
	*tchttp.BaseRequest

	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeUpdateTaskLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateTaskLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportSocTamperProofInfoRequest struct {
	*tchttp.BaseRequest

	// 排序字段方式 EventTamperTime

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportSocTamperProofInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportSocTamperProofInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentsListRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAgentsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentsListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetDatabaseInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Database *AssetDatabaseDetail `json:"Database,omitempty" name:"Database"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetDatabaseInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDatabaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*EventsHostLogin `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoginEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结束时间

		FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
		// 开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStatisticTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStatisticTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZKGroupListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*ZkGroup `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZKGroupListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZKGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLicenseInfoLicenseInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件下载链接

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 文件名称

		FileName *string `json:"FileName,omitempty" name:"FileName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLicenseInfoLicenseInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLicenseInfoLicenseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type License struct {
	// 租户数

	AppIdNum *int64 `json:"AppIdNum,omitempty" name:"AppIdNum"`
	// 授权到期时间

	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`
	// 支持功能

	Functions []*string `json:"Functions,omitempty" name:"Functions"`
	// 支持功能英文翻译

	FunctionsEn []*string `json:"FunctionsEn,omitempty" name:"FunctionsEn"`
	// 已售出授权数

	InquireNum *int64 `json:"InquireNum,omitempty" name:"InquireNum"`
	// 授权总数

	PurchaseNum *int64 `json:"PurchaseNum,omitempty" name:"PurchaseNum"`
	// 刷新时间

	RefreshTime *string `json:"RefreshTime,omitempty" name:"RefreshTime"`
	// 剩余可售授权

	SurplusInquireNum *int64 `json:"SurplusInquireNum,omitempty" name:"SurplusInquireNum"`
	// 类型 专业版：Pro 旗舰版：Flagship

	Type *string `json:"Type,omitempty" name:"Type"`
	// 已绑定agent总数

	UsedNum *int64 `json:"UsedNum,omitempty" name:"UsedNum"`
}

type DescribeAssetJarInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetJarInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetJarInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineRuleListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [LastTime|ItemCount|PassedItemCount|NotPassedItemCount|FirstTime]

	By *string `json:"By,omitempty" name:"By"`
	// <li>AppUserId - int64 - 是否必填：否 - 租户appId</li><li>HostName - string - 是否必填：否 - 主机名称</i><li>HostIp - string - 是否必填：否 - 主机Ip</i><li>DetectStatus - int - 是否必填：否 - 检测状态</li><li>StartTime - string - 是否必填：否 - 开始时间</li><li>EndTime - string - 是否必填：否 - 结束时间</li>

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBaselineRuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineRuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpecimenDetail struct {
	// 0未备份 1已经备份

	Backup *int64 `json:"Backup,omitempty" name:"Backup"`
	// 黑白灰属性

	BlackWhiteType *int64 `json:"BlackWhiteType,omitempty" name:"BlackWhiteType"`
	// 无

	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	DangerLevel *int64 `json:"DangerLevel,omitempty" name:"DangerLevel"`
	// 0是正常 1是会被删除记录且删除redis

	DelRecord *int64 `json:"DelRecord,omitempty" name:"DelRecord"`
	// 原始文件路径

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// sample filesize

	Filesize *int64 `json:"Filesize,omitempty" name:"Filesize"`
	// 样本来源

	From *string `json:"From,omitempty" name:"From"`
	// 0=位置是否能下在，1=ftp正常 2=ftp失败或者文件md5不一样

	FtpValid *int64 `json:"FtpValid,omitempty" name:"FtpValid"`
	// 分布式存储路径

	HdfsPath *string `json:"HdfsPath,omitempty" name:"HdfsPath"`
	// 文件被发现次数

	Hits *int64 `json:"Hits,omitempty" name:"Hits"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	IsDiff *bool `json:"IsDiff,omitempty" name:"IsDiff"`
	// 无

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 无

	NoFile *int64 `json:"NoFile,omitempty" name:"NoFile"`
	// 该记录操作者

	OperateUser *string `json:"OperateUser,omitempty" name:"OperateUser"`
	// 报毒数量

	Ratio *int64 `json:"Ratio,omitempty" name:"Ratio"`
	// 回扫计数器

	ReScanCount *int64 `json:"ReScanCount,omitempty" name:"ReScanCount"`
	// 监控到的原因

	Reason *int64 `json:"Reason,omitempty" name:"Reason"`
	// 无

	RedisBlackWhiteType *string `json:"RedisBlackWhiteType,omitempty" name:"RedisBlackWhiteType"`
	// 无

	ScantTimes *int64 `json:"ScantTimes,omitempty" name:"ScantTimes"`
	// 用于页面展示的黑白灰属性

	ShowBlackWhiteType *int64 `json:"ShowBlackWhiteType,omitempty" name:"ShowBlackWhiteType"`
	// 实时status=1未检测 status=2已经检测完毕

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 保存样本所在的ip

	StoreHost *string `json:"StoreHost,omitempty" name:"StoreHost"`
	// 无

	StorePath *string `json:"StorePath,omitempty" name:"StorePath"`
	// 无

	Tags *int64 `json:"Tags,omitempty" name:"Tags"`
	// 文件类型  webshell=0  binary=1

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 病毒名称

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
}

type DescribeHostRiskTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Data *SecurityDynamicsStatistic `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostRiskTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostRiskTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpecimenFile struct {
	// 无

	Content *string `json:"Content,omitempty" name:"Content"`
}

type DescribeTaskComponentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*VulComponent `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskComponentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FindAgentUuidsByPlatformRequest struct {
	*tchttp.BaseRequest

	// 平台：win linux

	Platform *string `json:"Platform,omitempty" name:"Platform"`
}

func (r *FindAgentUuidsByPlatformRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FindAgentUuidsByPlatformRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRemoteUpdateUserTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRemoteUpdateUserTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRemoteUpdateUserTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetMachineDetail struct {
	// 无

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// 无

	BootTime *string `json:"BootTime,omitempty" name:"BootTime"`
	// 无

	BuyTime *string `json:"BuyTime,omitempty" name:"BuyTime"`
	// 无

	CoreVersion *string `json:"CoreVersion,omitempty" name:"CoreVersion"`
	// 无

	Cpu *string `json:"Cpu,omitempty" name:"Cpu"`
	// 无

	CpuLoad *string `json:"CpuLoad,omitempty" name:"CpuLoad"`
	// 无

	CpuLoadVul *string `json:"CpuLoadVul,omitempty" name:"CpuLoadVul"`
	// 无

	CpuSize *uint64 `json:"CpuSize,omitempty" name:"CpuSize"`
	// 无

	DeviceVersion *string `json:"DeviceVersion,omitempty" name:"DeviceVersion"`
	// 无

	DiskLoad *string `json:"DiskLoad,omitempty" name:"DiskLoad"`
	// 无

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 无

	Disks []*DiskInfo `json:"Disks,omitempty" name:"Disks"`
	// 无

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 无

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 无

	InstallTime *string `json:"InstallTime,omitempty" name:"InstallTime"`
	// 无

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 无

	LastLiveTime *string `json:"LastLiveTime,omitempty" name:"LastLiveTime"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	MemLoad *string `json:"MemLoad,omitempty" name:"MemLoad"`
	// 无

	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`
	// 无

	NetCards []*NetCardsArray `json:"NetCards,omitempty" name:"NetCards"`
	// 无

	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	OsType *string `json:"OsType,omitempty" name:"OsType"`
	// 无

	PartitionCount *uint64 `json:"PartitionCount,omitempty" name:"PartitionCount"`
	// 无

	Producer *string `json:"Producer,omitempty" name:"Producer"`
	// 无

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 无

	ProtectDays *uint64 `json:"ProtectDays,omitempty" name:"ProtectDays"`
	// 无

	ProtectLevel *uint64 `json:"ProtectLevel,omitempty" name:"ProtectLevel"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	RiskStatus *string `json:"RiskStatus,omitempty" name:"RiskStatus"`
	// 无

	SerialNumber *string `json:"SerialNumber,omitempty" name:"SerialNumber"`
	// 无

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAssetDiskListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 主机Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 主机Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetDiskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDiskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetPortCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Ports []*AssetKeyVal `json:"Ports,omitempty" name:"Ports"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetPortCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPortCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesSimpleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机列表

		Machines []*MachineSimple `json:"Machines,omitempty" name:"Machines"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachinesSimpleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachinesSimpleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishInstallPackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishInstallPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishInstallPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAgentsListRequest struct {
	*tchttp.BaseRequest

	// uuid appid quuid hostip

	Filed *string `json:"Filed,omitempty" name:"Filed"`
	// 目标

	Targets *string `json:"Targets,omitempty" name:"Targets"`
}

func (r *ExportAgentsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAgentsListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBruteAttackTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBruteAttackTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBruteAttackTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDefencePluginListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVulDefencePluginListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulDefencePluginListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetInitServiceBaseInfo struct {
	// 无

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 无

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	Path *string `json:"Path,omitempty" name:"Path"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 无

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	User *string `json:"User,omitempty" name:"User"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeVulAllEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*EventsVul `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulAllEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulAllEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateUserCurrentConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的配置个数

		ConfigTotalCount *uint64 `json:"ConfigTotalCount,omitempty" name:"ConfigTotalCount"`
		// 用户某一类型的当前配置

		CurrentConfigs []*RemoteUpdateCurrentConfigInfo `json:"CurrentConfigs,omitempty" name:"CurrentConfigs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemoteUpdateUserCurrentConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateUserCurrentConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoteUpdateComponent struct {
	// 组件支持的架构系统，取值如下
	//
	// 1：x86
	//
	// 2：x86_64
	//
	// 3：arm
	//
	// 4：arm64

	Architecture *uint64 `json:"Architecture,omitempty" name:"Architecture"`
	// 组件支持的系统类型，取值如下
	//
	// 1：windows
	//
	// 2：linux

	OS *uint64 `json:"OS,omitempty" name:"OS"`
	// 组件类型，取值如下
	//
	// 1：agent
	//
	// 2：malware的TAV引擎
	//
	// 3：漏洞库

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type DeleteBashEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBashEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBashEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetInitServiceListRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetInitServiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetInitServiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineRuleCategoryListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [LastTime|ItemCount|PassedItemCount|NotPassedItemCount|FirstTime]

	By *string `json:"By,omitempty" name:"By"`
	// <li>AppUserId - int64 - 是否必填：否 - 租户appId</li><li>HostName - string - 是否必填：否 - 主机名称</i><li>HostIp - string - 是否必填：否 - 主机Ip</i><li>DetectStatus - int - 是否必填：否 - 检测状态</li><li>StartTime - string - 是否必填：否 - 开始时间</li><li>EndTime - string - 是否必填：否 - 结束时间</li>

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBaselineRuleCategoryListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineRuleCategoryListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticTimeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeStatisticTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStatisticTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTaskLog struct {
	// 无

	Content *int64 `json:"Content,omitempty" name:"Content"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	LogTime *string `json:"LogTime,omitempty" name:"LogTime"`
	// 无

	OperateUser *string `json:"OperateUser,omitempty" name:"OperateUser"`
	// 无

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DeletePrivilegeEventRequest struct {
	*tchttp.BaseRequest

	// 无

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeletePrivilegeEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivilegeEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRemoteUpdateUsersRequest struct {
	*tchttp.BaseRequest

	// 要删除的用户ID列表

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteRemoteUpdateUsersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRemoteUpdateUsersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVulScanConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVulScanConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVulScanConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDefenceEventRequest struct {
	*tchttp.BaseRequest

	// By

	By *string `json:"By,omitempty" name:"By"`
	// Name="AppId", "IP", "SourceIp", "VulName", "DefenseStatus", "EventType", "Status", "CreateTime","MergeTime"

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportVulDefenceEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulDefenceEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLicenseInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`
	// 无

	Ddl *string `json:"Ddl,omitempty" name:"Ddl"`
	// 无

	DomainMain *string `json:"DomainMain,omitempty" name:"DomainMain"`
	// 无

	Functions []*string `json:"Functions,omitempty" name:"Functions"`
	// 无

	Number *int64 `json:"Number,omitempty" name:"Number"`
	// 无

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 无

	PositionName *string `json:"PositionName,omitempty" name:"PositionName"`
	// 无

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 无

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *CreateLicenseInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLicenseInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefencePluginListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*VulDefencePluginDetail `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDefencePluginListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefencePluginListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProLicenseCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProLicenseCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProLicenseCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSocTamperProofInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSocTamperProofInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSocTamperProofInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JavaMemShellEvent struct {
	// 主机名称

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 租户appid

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 首次发现时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 说明

	Description *string `json:"Description,omitempty" name:"Description"`
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 实例id

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 内网ip

	PrivateIP *string `json:"PrivateIP,omitempty" name:"PrivateIP"`
	// quuid

	QUUID *string `json:"QUUID,omitempty" name:"QUUID"`
	// 最近扫描时间

	RecentFoundTime *string `json:"RecentFoundTime,omitempty" name:"RecentFoundTime"`
	// 事件状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// Javan内存马类型

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// uuid

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 外网ip

	WanIP *string `json:"WanIP,omitempty" name:"WanIP"`
}

type ModifyDNSEventsRequest struct {
	*tchttp.BaseRequest

	// 无

	Ids []*DNSIdInfo `json:"Ids,omitempty" name:"Ids"`
	// 是否是物理删除：true\false

	IsHardDelete *bool `json:"IsHardDelete,omitempty" name:"IsHardDelete"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDNSEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDNSEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoginEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoginEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTaskComponentsRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteTaskComponentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTaskComponentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetUserListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAssetUserListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetUserListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartUpdateTaskRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *StartUpdateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartUpdateTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetAppListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAssetAppListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetAppListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileTamperRuleCountRequest struct {
	*tchttp.BaseRequest

	// 租户APPID

	UserAppId *int64 `json:"UserAppId,omitempty" name:"UserAppId"`
	// 查询的主机uuids 一次性最多只能查100个

	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
}

func (r *DescribeFileTamperRuleCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileTamperRuleCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开关:0关，1开

		Enable *int64 `json:"Enable,omitempty" name:"Enable"`
		// HostCount

		HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDefenceCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞详情

		VulDetail *VulVuls `json:"VulDetail,omitempty" name:"VulDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefenseVulInfoRequest struct {
	*tchttp.BaseRequest

	// 漏洞Id

	VulId *int64 `json:"VulId,omitempty" name:"VulId"`
}

func (r *DescribeDefenseVulInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefenseVulInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportDefenseVulVulsListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportDefenseVulVulsListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDefenseVulVulsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventsVul struct {
	// 无

	Appid *int64 `json:"Appid,omitempty" name:"Appid"`
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	CvmName *string `json:"CvmName,omitempty" name:"CvmName"`
	// 无

	Descript *string `json:"Descript,omitempty" name:"Descript"`
	// 无

	FixStatus *int64 `json:"FixStatus,omitempty" name:"FixStatus"`
	// 无

	FixStatusMsg *string `json:"FixStatusMsg,omitempty" name:"FixStatusMsg"`
	// 无

	Guid *string `json:"Guid,omitempty" name:"Guid"`
	// 无

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	LatestFixTime *string `json:"LatestFixTime,omitempty" name:"LatestFixTime"`
	// 无

	LatestScanTime *string `json:"LatestScanTime,omitempty" name:"LatestScanTime"`
	// 无

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 无

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 无

	VulCategory *int64 `json:"VulCategory,omitempty" name:"VulCategory"`
	// 无

	VulId *int64 `json:"VulId,omitempty" name:"VulId"`
	// 无

	VulName *string `json:"VulName,omitempty" name:"VulName"`
}

type DescribeVulScanConfigListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*VulScanConfig `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulScanConfigListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulScanConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUpdateGrayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// id、name

		Config []*SimpleSelect `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUpdateGrayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUpdateGrayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSocTamperProofInfoRequest struct {
	*tchttp.BaseRequest

	// 操作状态

	OptStatus *uint64 `json:"OptStatus,omitempty" name:"OptStatus"`
	// 网页防篡改ID

	SocTamperProofId []*uint64 `json:"SocTamperProofId,omitempty" name:"SocTamperProofId"`
}

func (r *UpdateSocTamperProofInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSocTamperProofInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVulDefenceEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVulDefenceEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVulDefenceEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*EventsReverseShell `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MachineExtraInfo struct {
	// 主机名

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 网络名，vpc网络情况下会返回vpc_id

	NetworkName *string `json:"NetworkName,omitempty" name:"NetworkName"`
	// 网络类型，1:vpc网络 2:基础网络 3:非腾讯云网络

	NetworkType *int64 `json:"NetworkType,omitempty" name:"NetworkType"`
	// 内网IP

	PrivateIP *string `json:"PrivateIP,omitempty" name:"PrivateIP"`
	// 公网IP

	WanIP *string `json:"WanIP,omitempty" name:"WanIP"`
}

type DescribeCwpProductUsagesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCwpProductUsagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCwpProductUsagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstallPackageListRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeInstallPackageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstallPackageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetEnvListRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters *FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAssetEnvListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetEnvListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterParam struct {
	// 模糊搜索

	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 返回数量，默认为10，最大值为100。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type UpdateGray struct {
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Num *int64 `json:"Num,omitempty" name:"Num"`
	// 无

	PublishName *string `json:"PublishName,omitempty" name:"PublishName"`
	// 无

	Reverse *string `json:"Reverse,omitempty" name:"Reverse"`
	// 无

	Static *int64 `json:"Static,omitempty" name:"Static"`
	// 无

	Time *string `json:"Time,omitempty" name:"Time"`
	// 无

	User *string `json:"User,omitempty" name:"User"`
	// 无

	Uuids *string `json:"Uuids,omitempty" name:"Uuids"`
}

type DeleteUpdateGrayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// id、name

		Config []*SimpleSelect `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUpdateGrayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUpdateGrayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetJarInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Jar *AssetJarDetail `json:"Jar,omitempty" name:"Jar"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetJarInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetJarInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileSpecimenListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*FileSpecimenAll `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileSpecimenListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileSpecimenListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetCoreModuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Modules []*AssetCoreModuleBaseInfo `json:"Modules,omitempty" name:"Modules"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetCoreModuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetCoreModuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpdateGrayDetailRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeUpdateGrayDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateGrayDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetDatabaseBaseInfo struct {
	// 无

	BinPath *string `json:"BinPath,omitempty" name:"BinPath"`
	// 无

	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`
	// 无

	DataPath *string `json:"DataPath,omitempty" name:"DataPath"`
	// 无

	ErrorLogPath *string `json:"ErrorLogPath,omitempty" name:"ErrorLogPath"`
	// 无

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 无

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 无

	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	Param *string `json:"Param,omitempty" name:"Param"`
	// 无

	Permission *string `json:"Permission,omitempty" name:"Permission"`
	// 无

	PlugInPath *string `json:"PlugInPath,omitempty" name:"PlugInPath"`
	// 无

	Port *string `json:"Port,omitempty" name:"Port"`
	// 无

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 无

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	User *string `json:"User,omitempty" name:"User"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DeleteVulScanConfigRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteVulScanConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVulScanConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineHostDetectListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [LastTime|ItemCount|PassedItemCount|NotPassedItemCount|FirstTime]

	By *string `json:"By,omitempty" name:"By"`
	// <li>AppUserId - int64 - 是否必填：否 - 租户appId</li><li>HostName - string - 是否必填：否 - 主机名称</i><li>HostIp - string - 是否必填：否 - 主机Ip</i><li>DetectStatus - int - 是否必填：否 - 检测状态</li><li>StartTime - string - 是否必填：否 - 开始时间</li><li>EndTime - string - 是否必填：否 - 结束时间</li>

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBaselineHostDetectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineHostDetectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportJavaMemShellPluginsInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportJavaMemShellPluginsInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportJavaMemShellPluginsInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMalwareEvilResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMalwareEvilResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMalwareEvilResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJavaMemShellInfoRequest struct {
	*tchttp.BaseRequest

	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeJavaMemShellInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJavaMemShellInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MachineInfo struct {
	// 机器区域

	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
	// 机器类型

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

type DeleteSocTamperProofInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	SocTamperProofId []*uint64 `json:"SocTamperProofId,omitempty" name:"SocTamperProofId"`
}

func (r *DeleteSocTamperProofInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSocTamperProofInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefenceEventDetailRequest struct {
	*tchttp.BaseRequest

	// 漏洞事件id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeDefenceEventDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefenceEventDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetUserBaseInfo struct {
	// 无

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 无

	Gid *string `json:"Gid,omitempty" name:"Gid"`
	// 无

	HomePath *string `json:"HomePath,omitempty" name:"HomePath"`
	// 无

	IsDomain *uint64 `json:"IsDomain,omitempty" name:"IsDomain"`
	// 无

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 无

	IsRoot *uint64 `json:"IsRoot,omitempty" name:"IsRoot"`
	// 无

	IsSshLogin *uint64 `json:"IsSshLogin,omitempty" name:"IsSshLogin"`
	// 无

	IsSudo *uint64 `json:"IsSudo,omitempty" name:"IsSudo"`
	// 无

	LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`
	// 无

	LoginType *uint64 `json:"LoginType,omitempty" name:"LoginType"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	PasswordChangeTime *string `json:"PasswordChangeTime,omitempty" name:"PasswordChangeTime"`
	// 无

	PasswordDueTime *string `json:"PasswordDueTime,omitempty" name:"PasswordDueTime"`
	// 无

	PasswordLockDays *int64 `json:"PasswordLockDays,omitempty" name:"PasswordLockDays"`
	// 无

	PasswordStatus *uint64 `json:"PasswordStatus,omitempty" name:"PasswordStatus"`
	// 无

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Shell *string `json:"Shell,omitempty" name:"Shell"`
	// 无

	ShellLoginStatus *uint64 `json:"ShellLoginStatus,omitempty" name:"ShellLoginStatus"`
	// 无

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 无

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	UserType *uint64 `json:"UserType,omitempty" name:"UserType"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ModifyVulVulsInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVulVulsInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulVulsInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebAppPluginListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Plugins []*PluginInfo `json:"Plugins,omitempty" name:"Plugins"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebAppPluginListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebAppPluginListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZKGroupFullListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*ZkGroup `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZKGroupFullListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZKGroupFullListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetDatabaseListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAssetDatabaseListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetDatabaseListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetProcessInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetProcessInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetProcessInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportSocTamperProofInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportSocTamperProofInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportSocTamperProofInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetPlanTaskListRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters *FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *bool `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *ExportAssetPlanTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetPlanTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUpdateConfigRequest struct {
	*tchttp.BaseRequest

	// Md5

	MdSum *string `json:"MdSum,omitempty" name:"MdSum"`
	// 文件名ydeyes_linux64_3.5.0.106.tar.gz

	UpdatePackagePath *string `json:"UpdatePackagePath,omitempty" name:"UpdatePackagePath"`
}

func (r *ModifyUpdateConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUpdateConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZkGroup struct {
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeAssetWebLocationInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetWebLocationInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebLocationInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDNSKnowledgeStatusRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// Status=0 1

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDNSKnowledgeStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDNSKnowledgeStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportKnowledgeRequest struct {
	*tchttp.BaseRequest
}

func (r *ExportKnowledgeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportKnowledgeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebLocationInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		WebServices *AssetWebLocationDetail `json:"WebServices,omitempty" name:"WebServices"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebLocationInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebLocationInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportHostSafetyStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportHostSafetyStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportHostSafetyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetEnvListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetEnvListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetEnvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDNSKnowledgeInfoRequest struct {
	*tchttp.BaseRequest

	// Description=uri

	Description *string `json:"Description,omitempty" name:"Description"`
	// Domain=fqdn

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Reference *string `json:"Reference,omitempty" name:"Reference"`
}

func (r *ModifyDNSKnowledgeInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDNSKnowledgeInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateUserAccountUniqueRequest struct {
	*tchttp.BaseRequest

	// 用户账号

	UserAccount *string `json:"UserAccount,omitempty" name:"UserAccount"`
}

func (r *DescribeRemoteUpdateUserAccountUniqueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateUserAccountUniqueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulScanConfigFullListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVulScanConfigFullListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulScanConfigFullListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulTopInfo struct {
	// 无

	VulCategory *int64 `json:"VulCategory,omitempty" name:"VulCategory"`
	// 无

	VulCount *int64 `json:"VulCount,omitempty" name:"VulCount"`
	// 无

	VulEventId *int64 `json:"VulEventId,omitempty" name:"VulEventId"`
	// 无

	VulId *int64 `json:"VulId,omitempty" name:"VulId"`
	// 无

	VulLevel *int64 `json:"VulLevel,omitempty" name:"VulLevel"`
	// 无

	VulName *string `json:"VulName,omitempty" name:"VulName"`
}

type ModifyDNSKnowledgeInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDNSKnowledgeInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDNSKnowledgeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebLocationPathListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Paths []*PathInfo `json:"Paths,omitempty" name:"Paths"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebLocationPathListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebLocationPathListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBruteAttackTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 暴力破解ip数组

		BruteAttackList []*BruteAttackTopInfo `json:"BruteAttackList,omitempty" name:"BruteAttackList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBruteAttackTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBruteAttackTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyJavaMemShellsStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyJavaMemShellsStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyJavaMemShellsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZKGroupInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyZKGroupInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyZKGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulDefenceSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVulDefenceSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulDefenceSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDNSKnowledgeListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Status=0 1

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeDNSKnowledgeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDNSKnowledgeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncAssetScanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 最新结束同步时间

		LatestEndTime *string `json:"LatestEndTime,omitempty" name:"LatestEndTime"`
		// 最新开始同步时间

		LatestStartTime *string `json:"LatestStartTime,omitempty" name:"LatestStartTime"`
		// 枚举值有(大写)：NOTASK（没有同步任务），SYNCING（同步中），FINISHED（同步完成）

		State *string `json:"State,omitempty" name:"State"`
		// 无

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncAssetScanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncAssetScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpdateTaskStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// UpdateTask

		List []*UpdateTaskStatus `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpdateTaskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVersionStatisticsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 基础版数量

		BasicVersionNum *uint64 `json:"BasicVersionNum,omitempty" name:"BasicVersionNum"`
		// 专业版数量

		ProVersionNum *uint64 `json:"ProVersionNum,omitempty" name:"ProVersionNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVersionStatisticsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVersionStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportKnowledgeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportKnowledgeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportKnowledgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulScanTaskListRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeVulScanTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulScanTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRemoteUpdateConfigRequest struct {
	*tchttp.BaseRequest

	// 要导出配置的Id

	ConfigIds []*uint64 `json:"ConfigIds,omitempty" name:"ConfigIds"`
}

func (r *ExportRemoteUpdateConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRemoteUpdateConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDefenceMachineListRequest struct {
	*tchttp.BaseRequest

	// By

	By *string `json:"By,omitempty" name:"By"`
	// Name="AppId", "IP", "InstanceName", "InstanceId", "Enable", "Status", "CreateTime", "ModifyTime"

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportVulDefenceMachineListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulDefenceMachineListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBruteForceEventsRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportBruteForceEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBruteForceEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetCoreModuleBaseInfo struct {
	// 无

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 无

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	ModuleCount *uint64 `json:"ModuleCount,omitempty" name:"ModuleCount"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	Path *string `json:"Path,omitempty" name:"Path"`
	// 无

	ProcessCount *uint64 `json:"ProcessCount,omitempty" name:"ProcessCount"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DnsKnowledge struct {
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	Description *string `json:"Description,omitempty" name:"Description"`
	// 无

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 无

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 无

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 无

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 无

	Reference *string `json:"Reference,omitempty" name:"Reference"`
	// 无

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type FileSpecimenAll struct {
	// 0未备份 1已经备份

	Backup *int64 `json:"Backup,omitempty" name:"Backup"`
	// 黑白灰属性

	BlackWhiteType *int64 `json:"BlackWhiteType,omitempty" name:"BlackWhiteType"`
	// 无

	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	DangerLevel *int64 `json:"DangerLevel,omitempty" name:"DangerLevel"`
	// 0是正常 1是会被删除记录且删除redis

	DelRecord *int64 `json:"DelRecord,omitempty" name:"DelRecord"`
	// 无

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 原始文件路径

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// sample filesize

	Filesize *int64 `json:"Filesize,omitempty" name:"Filesize"`
	// 样本来源

	From *string `json:"From,omitempty" name:"From"`
	// 0=位置是否能下在，1=ftp正常 2=ftp失败或者文件md5不一样

	FtpValid *int64 `json:"FtpValid,omitempty" name:"FtpValid"`
	// 分布式存储路径

	HdfsPath *string `json:"HdfsPath,omitempty" name:"HdfsPath"`
	// 文件被发现次数

	Hits *int64 `json:"Hits,omitempty" name:"Hits"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 样本md5

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 无

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 无

	NoFile *int64 `json:"NoFile,omitempty" name:"NoFile"`
	// 该记录操作者 如果是程序插入的，写程序名称，如果是运营平台插入的，写操作者名字

	OperateUser *string `json:"OperateUser,omitempty" name:"OperateUser"`
	// 报毒数量

	Ratio *int64 `json:"Ratio,omitempty" name:"Ratio"`
	// 回扫计数器

	ReScanCount *int64 `json:"ReScanCount,omitempty" name:"ReScanCount"`
	// 监控到的原因

	Reason *int64 `json:"Reason,omitempty" name:"Reason"`
	// 无

	ScanTimes *int64 `json:"ScanTimes,omitempty" name:"ScanTimes"`
	// 实时status=1未检测 status=2已经检测完毕

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 保存样本所在的ip

	StoreHost *string `json:"StoreHost,omitempty" name:"StoreHost"`
	// 保存路径

	StorePath *string `json:"StorePath,omitempty" name:"StorePath"`
	// 无

	Tags *int64 `json:"Tags,omitempty" name:"Tags"`
	// 文件类型  webshell=0  binary=1

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 病毒名称

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
}

type DeleteFileSpecimenRequest struct {
	*tchttp.BaseRequest

	// 无

	Md5 []*string `json:"Md5,omitempty" name:"Md5"`
}

func (r *DeleteFileSpecimenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFileSpecimenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVersionStatisticsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVersionStatisticsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVersionStatisticsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetMachineDetailRequest struct {
	*tchttp.BaseRequest

	// 无

	NeedDisk *bool `json:"NeedDisk,omitempty" name:"NeedDisk"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetMachineDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetMachineDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRemoteUpdateUserRequest struct {
	*tchttp.BaseRequest

	// 用户当前更新配置的Id

	ConfigIds []*uint64 `json:"ConfigIds,omitempty" name:"ConfigIds"`
	// 用户信息

	UserInfo *RemoteUpdateUser `json:"UserInfo,omitempty" name:"UserInfo"`
}

func (r *ModifyRemoteUpdateUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRemoteUpdateUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetWebAppListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetWebAppListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetWebAppListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportKnowledgeRequest struct {
	*tchttp.BaseRequest

	// 文件名字

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件MD5

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

func (r *ImportKnowledgeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportKnowledgeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulVulsInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	Author *string `json:"Author,omitempty" name:"Author"`
	// 无

	CheckRule *string `json:"CheckRule,omitempty" name:"CheckRule"`
	// 无

	ComponentId *int64 `json:"ComponentId,omitempty" name:"ComponentId"`
	// 无

	CveId *string `json:"CveId,omitempty" name:"CveId"`
	// 无

	Cvss *string `json:"Cvss,omitempty" name:"Cvss"`
	// 无

	CvssScore *float64 `json:"CvssScore,omitempty" name:"CvssScore"`
	// 无

	Describe *string `json:"Describe,omitempty" name:"Describe"`
	// 英文描述

	DescribeEn *string `json:"DescribeEn,omitempty" name:"DescribeEn"`
	// 无

	Emergency *int64 `json:"Emergency,omitempty" name:"Emergency"`
	// 无

	Fix *string `json:"Fix,omitempty" name:"Fix"`
	// 英文修复建议

	FixEn *string `json:"FixEn,omitempty" name:"FixEn"`
	// 漏洞id，修改时必填

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 英文标签

	LableEn *string `json:"LableEn,omitempty" name:"LableEn"`
	// 无

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 漏洞名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 英文名称

	NameEn *string `json:"NameEn,omitempty" name:"NameEn"`
	// 无

	PocFile *string `json:"PocFile,omitempty" name:"PocFile"`
	// 无

	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
	// 参考地址

	Reference *string `json:"Reference,omitempty" name:"Reference"`
	// 0关、1开、2测试

	Switch *int64 `json:"Switch,omitempty" name:"Switch"`
	// 无

	VerifyRule *string `json:"VerifyRule,omitempty" name:"VerifyRule"`
	// 无

	VulCategory *int64 `json:"VulCategory,omitempty" name:"VulCategory"`
	// 无

	VulTypeId *int64 `json:"VulTypeId,omitempty" name:"VulTypeId"`
	// Version

	VulVersion *string `json:"VulVersion,omitempty" name:"VulVersion"`
}

func (r *ModifyVulVulsInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulVulsInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBashEventRequest struct {
	*tchttp.BaseRequest

	// 无

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteBashEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBashEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteZKGroupInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteZKGroupInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteZKGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetAppListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAssetAppListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetAppListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstallPackageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// UpdateInstallPackage

		List []*UpdateInstallPackage `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstallPackageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstallPackageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 漏洞top信息

		VulTopList []*VulTopInfo `json:"VulTopList,omitempty" name:"VulTopList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulDefenceEventStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVulDefenceEventStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulDefenceEventStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishInstallPackageRequest struct {
	*tchttp.BaseRequest

	// 无

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *PublishInstallPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishInstallPackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRemoteUpdateUserTagRequest struct {
	*tchttp.BaseRequest

	// 无

	TagId *uint64 `json:"TagId,omitempty" name:"TagId"`
	// 无

	UserId *uint64 `json:"UserId,omitempty" name:"UserId"`
}

func (r *ModifyRemoteUpdateUserTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRemoteUpdateUserTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetDatabaseDetail struct {
	// 无

	BinPath *string `json:"BinPath,omitempty" name:"BinPath"`
	// 无

	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`
	// 无

	DataPath *string `json:"DataPath,omitempty" name:"DataPath"`
	// 无

	ErrorLogPath *string `json:"ErrorLogPath,omitempty" name:"ErrorLogPath"`
	// 无

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 无

	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	Param *string `json:"Param,omitempty" name:"Param"`
	// 无

	Permission *string `json:"Permission,omitempty" name:"Permission"`
	// 无

	PlugInPath *string `json:"PlugInPath,omitempty" name:"PlugInPath"`
	// 无

	Port *string `json:"Port,omitempty" name:"Port"`
	// 无

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	User *string `json:"User,omitempty" name:"User"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeAssetSystemPackageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Packages []*AssetSystemPackageBaseInfo `json:"Packages,omitempty" name:"Packages"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetSystemPackageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSystemPackageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebFrameCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		WebApps []*AssetKeyVal `json:"WebApps,omitempty" name:"WebApps"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebFrameCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebFrameCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSocTamperProofListRequest struct {
	*tchttp.BaseRequest

	// 排序字段方式 EventTamperTime

	By *string `json:"By,omitempty" name:"By"`
	// 条件过滤, ExactMatch参数无效

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeSocTamperProofListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSocTamperProofListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetPortInfoListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAssetPortInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetPortInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpecimenFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 样本文件

		SpecimenFile *SpecimenFile `json:"SpecimenFile,omitempty" name:"SpecimenFile"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSpecimenFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpecimenFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZKGroupListRequest struct {
	*tchttp.BaseRequest

	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeZKGroupListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZKGroupListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDNSEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDNSEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDNSEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpecimenDetailRequest struct {
	*tchttp.BaseRequest

	// 无

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

func (r *DescribeSpecimenDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpecimenDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulEventsStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVulEventsStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulEventsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulTypeInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	Describe *string `json:"Describe,omitempty" name:"Describe"`
	// 无

	DescribeEn *string `json:"DescribeEn,omitempty" name:"DescribeEn"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	NameEn *string `json:"NameEn,omitempty" name:"NameEn"`
}

func (r *ModifyVulTypeInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulTypeInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulTopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVulTopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineHostDetectListRequest struct {
	*tchttp.BaseRequest

	// 0:过滤的结果导出；1:全部导出

	ExportAll *int64 `json:"ExportAll,omitempty" name:"ExportAll"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportBaselineHostDetectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineHostDetectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebLocationPathListRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetWebLocationPathListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebLocationPathListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoteUpdatePublishFailedConfig struct {
	// 失败原因

	Cause *string `json:"Cause,omitempty" name:"Cause"`
	// 配置Id

	ConfigId *uint64 `json:"ConfigId,omitempty" name:"ConfigId"`
}

type RemoteUpdateUser struct {
	// 账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 允许访问的IP。格式支持单个IP，范围（用中横杠分隔），和网段多种形式

	AllowAccessIPs []*string `json:"AllowAccessIPs,omitempty" name:"AllowAccessIPs"`
	// 是否授权可远程升级，取值如下
	//
	// 0：否
	//
	// 1：是
	//
	// *注意* 业务后台在检查时间是否超出合同期后，更新该值

	Authorized *uint64 `json:"Authorized,omitempty" name:"Authorized"`
	// 后台架构类型，取值如下
	//
	// 1：x86
	//
	// 2：x86_64
	//
	// 3：arm
	//
	// 4：arm64

	BackstageArchitecture *uint64 `json:"BackstageArchitecture,omitempty" name:"BackstageArchitecture"`
	// 后台系统类型，取值如下
	//
	// 1：windows
	//
	// 2：linux

	BackstageOS *uint64 `json:"BackstageOS,omitempty" name:"BackstageOS"`
	// 合同结束时间

	ContractEndTime *string `json:"ContractEndTime,omitempty" name:"ContractEndTime"`
	// 合同开始时间

	ContractStartTime *string `json:"ContractStartTime,omitempty" name:"ContractStartTime"`
	// 租户授权数

	CustomerCount *uint64 `json:"CustomerCount,omitempty" name:"CustomerCount"`
	// 环境类型，取值如下
	//
	// 1：TCE
	//
	// 2：私有化
	//
	// 3：云环境

	EnvironmentType *uint64 `json:"EnvironmentType,omitempty" name:"EnvironmentType"`
	// 框架版本

	FrameVersion *string `json:"FrameVersion,omitempty" name:"FrameVersion"`
	// 前场区技姓名

	FrontOperatorName *string `json:"FrontOperatorName,omitempty" name:"FrontOperatorName"`
	// 用户ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 用户名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Overlay设备agent数量

	OverAgentCount *uint64 `json:"OverAgentCount,omitempty" name:"OverAgentCount"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 密钥

	SecurityKey *string `json:"SecurityKey,omitempty" name:"SecurityKey"`
	// 该用户的标签Id

	TagId *uint64 `json:"TagId,omitempty" name:"TagId"`
	// Underlay设备agent数量

	UnderAgentCount *uint64 `json:"UnderAgentCount,omitempty" name:"UnderAgentCount"`
}

type DescribeAssetDatabaseListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAssetDatabaseListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDatabaseListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulScanTaskNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVulScanTaskNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulScanTaskNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyJavaMemShellPluginSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyJavaMemShellPluginSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyJavaMemShellPluginSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetCoreModuleDetail struct {
	// 无

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 无

	Modules *string `json:"Modules,omitempty" name:"Modules"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Params []*ParamInfo `json:"Params,omitempty" name:"Params"`
	// 无

	Path *string `json:"Path,omitempty" name:"Path"`
	// 无

	Processes *string `json:"Processes,omitempty" name:"Processes"`
	// 无

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DeleteRemoteUpdatePublishItemsRequest struct {
	*tchttp.BaseRequest

	// 要删除的发布项的Id

	PublishItemIds []*uint64 `json:"PublishItemIds,omitempty" name:"PublishItemIds"`
}

func (r *DeleteRemoteUpdatePublishItemsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRemoteUpdatePublishItemsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetMachineListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetMachineListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetMachineListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulScanTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*VulScanTask `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulScanTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulScanTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetWebServiceInfoListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *ExportAssetWebServiceInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetWebServiceInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportKnowledgePackageRequest struct {
	*tchttp.BaseRequest

	// 文件名字

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件MD5

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

func (r *ImportKnowledgePackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportKnowledgePackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetCoreModuleListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetCoreModuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetCoreModuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellEventsRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeReverseShellEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReverseShellEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineItemListRequest struct {
	*tchttp.BaseRequest

	// 0:过滤的结果导出；1:全部导出

	ExportAll *int64 `json:"ExportAll,omitempty" name:"ExportAll"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportBaselineItemListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineItemListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelPublishInstallPackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelPublishInstallPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelPublishInstallPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetDiskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 磁盘分区列表

		Disks []*AssetDiskPartitionInfo `json:"Disks,omitempty" name:"Disks"`
		// 分区总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetDiskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetWebAppListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *ExportAssetWebAppListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetWebAppListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetWebLocationListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetWebLocationListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetWebLocationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PluginInfo struct {
	// 无

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 无

	Link *string `json:"Link,omitempty" name:"Link"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
}

type SecurityDynamicsStatistic struct {
	// 无

	DateTime *string `json:"DateTime,omitempty" name:"DateTime"`
	// 无

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 无

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type DescribeRemoteUpdateConfigNameListRequest struct {
	*tchttp.BaseRequest

	// 配置（所属组件的）类型

	ComponentType *uint64 `json:"ComponentType,omitempty" name:"ComponentType"`
	// 配置版本

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
}

func (r *DescribeRemoteUpdateConfigNameListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateConfigNameListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpdateConfigRequest struct {
	*tchttp.BaseRequest

	// "UpdateName"

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUpdateConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDefenceEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVulDefenceEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulDefenceEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAgentPushTaskResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAgentPushTaskResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAgentPushTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventsMalware struct {
	// 发送告警次数

	AlarmSendTimes *int64 `json:"AlarmSendTimes,omitempty" name:"AlarmSendTimes"`
	// 云用户唯一标识

	Appid *int64 `json:"Appid,omitempty" name:"Appid"`
	// 无用

	Area *string `json:"Area,omitempty" name:"Area"`
	// 黑属性，20-疑似黑，21~29-黑

	BlackWhiteType *int64 `json:"BlackWhiteType,omitempty" name:"BlackWhiteType"`
	// 记录创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无用

	DataFrom *int64 `json:"DataFrom,omitempty" name:"DataFrom"`
	// 无

	Feature *string `json:"Feature,omitempty" name:"Feature"`
	// accesstime

	FileAccessTime *int64 `json:"FileAccessTime,omitempty" name:"FileAccessTime"`
	// createtime

	FileCreateTime *int64 `json:"FileCreateTime,omitempty" name:"FileCreateTime"`
	// modifiertime

	FileModifierTime *int64 `json:"FileModifierTime,omitempty" name:"FileModifierTime"`
	// 文件路径

	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
	// 文件大小

	Filesize *int64 `json:"Filesize,omitempty" name:"Filesize"`
	// 预留未用，主机唯一标识

	Guid *int64 `json:"Guid,omitempty" name:"Guid"`
	// 主机ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// is_check

	IsCheck *bool `json:"IsCheck,omitempty" name:"IsCheck"`
	// 最近扫描时间

	LatestScanTime *string `json:"LatestScanTime,omitempty" name:"LatestScanTime"`
	// 无

	MatchResult *string `json:"MatchResult,omitempty" name:"MatchResult"`
	// 文件MD5

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 无用

	MgtIp *string `json:"MgtIp,omitempty" name:"MgtIp"`
	// 记录最近修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 操作文件的返回值

	OperateRet *int64 `json:"OperateRet,omitempty" name:"OperateRet"`
	// 最近操作该记录的人或后台服务

	OperateUser *string `json:"OperateUser,omitempty" name:"OperateUser"`
	// 文件创建者

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 文件路径的md5

	PathMd5 *string `json:"PathMd5,omitempty" name:"PathMd5"`
	// ntfs监控时的reason，位标识了create、write、overwrite等状态

	Reason *int64 `json:"Reason,omitempty" name:"Reason"`
	// 无用

	Region *int64 `json:"Region,omitempty" name:"Region"`
	// 0为未恢复，1为已恢复

	Restored *int64 `json:"Restored,omitempty" name:"Restored"`
	// 状态， 0~2-待审核，4-待处理，5-已信任，6-已隔离， 8-文件已删除， 10-隔离中， 12-已删除

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 文件类型，1-二进制，0-webshell

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 主机唯一标识

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 病毒名

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
	// 无用

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
}

type MachineFileTamperRule struct {
	// 唯一id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规则

	Rule []*FileTamperRule `json:"Rule,omitempty" name:"Rule"`
	// 规则类型 0 ：系统规则  1：用户规则

	RuleCategory *uint64 `json:"RuleCategory,omitempty" name:"RuleCategory"`
}

type DescribeAssetUserInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		User *AssetUserDetail `json:"User,omitempty" name:"User"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetUserInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineHostDetectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselineHostDetect `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineHostDetectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineHostDetectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostSafetyStatusRequest struct {
	*tchttp.BaseRequest

	// StartTime

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 无

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeHostSafetyStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostSafetyStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVulEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpdateTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 无

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeUpdateTaskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateTaskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分周列表

		List []*string `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWeeksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZKConfigListRequest struct {
	*tchttp.BaseRequest

	// 无

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeZKConfigListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZKConfigListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadPushResultRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DownloadPushResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadPushResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetInitServiceListRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters *FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *ExportAssetInitServiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetInitServiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFileSpecimenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFileSpecimenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFileSpecimenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetProcessCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Process []*AssetKeyVal `json:"Process,omitempty" name:"Process"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetProcessCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetProcessCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebServiceCountRequest struct {
	*tchttp.BaseRequest

	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAssetWebServiceCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebServiceCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulVulsListRequest struct {
	*tchttp.BaseRequest

	// Filters："ListType", "Id", "Name", "Level", "VulCategory", "Switch"

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ExportVulVulsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulVulsListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostSafetyStatusStatistic struct {
	// 主机总数

	HostTotalCount *int64 `json:"HostTotalCount,omitempty" name:"HostTotalCount"`
	// 无

	NotRiskHostCount *int64 `json:"NotRiskHostCount,omitempty" name:"NotRiskHostCount"`
	// 风险主机数

	RiskHostCount *int64 `json:"RiskHostCount,omitempty" name:"RiskHostCount"`
	// 风险事件数组

	SecurityDynamicsStatistic []*SecurityDynamicsStatistic `json:"SecurityDynamicsStatistic,omitempty" name:"SecurityDynamicsStatistic"`
}

type ExportCustomerLicenseListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportCustomerLicenseListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportCustomerLicenseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDefencePluginListRequest struct {
	*tchttp.BaseRequest

	// By

	By *string `json:"By,omitempty" name:"By"`
	// Name="Quuid", "Pid", "MainClass", "Status", "ErrorLog"

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportVulDefencePluginListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulDefencePluginListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetSystemPackageBaseInfo struct {
	// 无

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 无

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 无

	InstallTime *string `json:"InstallTime,omitempty" name:"InstallTime"`
	// 无

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	Type *string `json:"Type,omitempty" name:"Type"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
}

type AssetProcessBaseInfo struct {
	// 无

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 无

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 无

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 无

	HasSign *uint64 `json:"HasSign,omitempty" name:"HasSign"`
	// 无

	InstallByPackage *uint64 `json:"InstallByPackage,omitempty" name:"InstallByPackage"`
	// 无

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`
	// 无

	Param *string `json:"Param,omitempty" name:"Param"`
	// 无

	ParentProcessName *string `json:"ParentProcessName,omitempty" name:"ParentProcessName"`
	// 无

	Path *string `json:"Path,omitempty" name:"Path"`
	// 无

	Pid *string `json:"Pid,omitempty" name:"Pid"`
	// 无

	Ppid *string `json:"Ppid,omitempty" name:"Ppid"`
	// 无

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 无

	Status *string `json:"Status,omitempty" name:"Status"`
	// 无

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 无

	Tty *string `json:"Tty,omitempty" name:"Tty"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	User *string `json:"User,omitempty" name:"User"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
}

type Filter struct {
	// 模糊搜索

	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
	// 过滤键的名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一个或者多个过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type ExportAssetDatabaseListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetDatabaseListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetDatabaseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Types []*AssetType `json:"Types,omitempty" name:"Types"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMalwareRequest struct {
	*tchttp.BaseRequest

	// trust removeTrust del

	CMD *string `json:"CMD,omitempty" name:"CMD"`
	// 无

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *UpdateMalwareRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMalwareRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentPushTaskListRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAgentPushTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentPushTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetProcessInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Process []*AssetProcessBaseInfo `json:"Process,omitempty" name:"Process"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetProcessInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetProcessInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateUserRequest struct {
	*tchttp.BaseRequest

	// 用户ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeRemoteUpdateUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetUserListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetUserListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportPrivilegeEventsRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportPrivilegeEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportPrivilegeEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineRule struct {
	// 无

	AssetType *int64 `json:"AssetType,omitempty" name:"AssetType"`
	// 规则分类

	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 主机数

	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`
	// 无

	HostIDs []*string `json:"HostIDs,omitempty" name:"HostIDs"`
	// 无

	HostIPs []*string `json:"HostIPs,omitempty" name:"HostIPs"`
	// 适配项ID列表

	Items []*Item `json:"Items,omitempty" name:"Items"`
	// 规则描述

	RuleDesc *string `json:"RuleDesc,omitempty" name:"RuleDesc"`
	// 规则Id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称,长度不超过128英文字符

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 规则类型

	RuleType *int64 `json:"RuleType,omitempty" name:"RuleType"`
}

type DescribeDNSEventDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// dns事件详情

		DNSEventDetail *DNSEventDetail `json:"DNSEventDetail,omitempty" name:"DNSEventDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDNSEventDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDNSEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebAppPluginListRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetWebAppPluginListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebAppPluginListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*LicenseUpgrade `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulDefenceEvent struct {
	// AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 地址

	City *string `json:"City,omitempty" name:"City"`
	// 攻击次数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 首次攻击时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// CveId

	CveId *string `json:"CveId,omitempty" name:"CveId"`
	// 攻击状态（全部状态、尝试攻击、攻击成功、防御成功）

	EventType *uint64 `json:"EventType,omitempty" name:"EventType"`
	// HostName

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// InstanceID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 最近更新时间

	MergeTime *string `json:"MergeTime,omitempty" name:"MergeTime"`
	// PrivateIp

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// PublicIp

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 攻击来源IP

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 攻击源端口

	SourcePort *string `json:"SourcePort,omitempty" name:"SourcePort"`
	// 处理状态  （全部状态、待处理、已处理、已忽略、租户已删除）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 漏洞分类

	VulCategory *int64 `json:"VulCategory,omitempty" name:"VulCategory"`
	// VulId

	VulId *int64 `json:"VulId,omitempty" name:"VulId"`
	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
}

type ExportJavaMemShellPluginsInfoRequest struct {
	*tchttp.BaseRequest

	// FilterParams

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportJavaMemShellPluginsInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportJavaMemShellPluginsInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetAppProcessListRequest struct {
	*tchttp.BaseRequest

	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetAppProcessListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetAppProcessListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateUserAccountUniqueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账号是否已存在

		Existed *bool `json:"Existed,omitempty" name:"Existed"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemoteUpdateUserAccountUniqueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateUserAccountUniqueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRemoteUpdatePublishItemResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 发布项的Id

		PublishItemId *uint64 `json:"PublishItemId,omitempty" name:"PublishItemId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRemoteUpdatePublishItemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRemoteUpdatePublishItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileTamperEvent struct {
	// APPID

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 发生时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 危害描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 事件产生次数

	EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`
	// 进程名

	ExeMd5 *string `json:"ExeMd5,omitempty" name:"ExeMd5"`
	// 进程名称

	ExeName *string `json:"ExeName,omitempty" name:"ExeName"`
	// 进程权限

	ExePermission *string `json:"ExePermission,omitempty" name:"ExePermission"`
	// 进程pid

	ExePid *uint64 `json:"ExePid,omitempty" name:"ExePid"`
	// 进程文件大小

	ExeSize *uint64 `json:"ExeSize,omitempty" name:"ExeSize"`
	// 进程执行时长

	ExeTime *uint64 `json:"ExeTime,omitempty" name:"ExeTime"`
	// 机器IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 机器名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 事件id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 风险等级 0：无， 1: 高危， 2:中危， 3: 低危

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	//  主机额外信息

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// 主机在线信息 ONLINE、OFFLINE

	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`
	// 最近发生时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 进程参数

	ProcessArg *string `json:"ProcessArg,omitempty" name:"ProcessArg"`
	// 进程路径

	ProcessExe *string `json:"ProcessExe,omitempty" name:"ProcessExe"`
	// 事件详情: json格式

	PsTree *string `json:"PsTree,omitempty" name:"PsTree"`
	// cvm id

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 参考链接

	Reference *string `json:"Reference,omitempty" name:"Reference"`
	// 规则类型 0系统规则 1自定义规则

	RuleCategory *uint64 `json:"RuleCategory,omitempty" name:"RuleCategory"`
	// 规则id

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 规则名称

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 处理状态  0 -- 待处理 1 -- 已加白 2 -- 已删除 3 - 已忽略 4-已手动处理

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 修护建议

	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`
	// 目标文件路径

	Target *string `json:"Target,omitempty" name:"Target"`
	// 目标文件创建时间

	TargetCreatTime *string `json:"TargetCreatTime,omitempty" name:"TargetCreatTime"`
	// 目标文件更新时间

	TargetModifyTime *string `json:"TargetModifyTime,omitempty" name:"TargetModifyTime"`
	// 文件名称

	TargetName *string `json:"TargetName,omitempty" name:"TargetName"`
	// 目标文件权限

	TargetPermission *string `json:"TargetPermission,omitempty" name:"TargetPermission"`
	// 目标文件大小

	TargetSize *uint64 `json:"TargetSize,omitempty" name:"TargetSize"`
	// 事件类型/动作  0 -- 告警

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 用户组

	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ExportAssetJarListRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters *FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *ExportAssetJarListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetJarListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskComponentsRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTaskComponentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskComponentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZKConfigInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Path *string `json:"Path,omitempty" name:"Path"`
	// 无

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 无

	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *ModifyZKConfigInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyZKConfigInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateUserListRequest struct {
	*tchttp.BaseRequest

	// 是否降序，0：否，1：是

	Descending *uint64 `json:"Descending,omitempty" name:"Descending"`
	// 过滤，支持条件如下：
	// UserName  - 是否必填: 否 - 按用户名称过滤
	// TagId - 是否必填: 否 - 按标签过滤
	// FrameVersion - 是否必填: 否 - 按框架版本过滤
	// UserId - 是否必填：否 - 按UserId过滤
	// Account - 是否必填：否 - 按账号过滤
	// BackstageArchitecture - 是否必填：否 - 按后台架构过滤
	// BackstageOS - 是否必填 - 否 - 按后台系统过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序，支持如下传入：
	//
	// Name：按名称排序（默认）
	//
	// CreateTime：按创建时间排序
	//
	// UpdateTime：按更新时间排序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *DescribeRemoteUpdateUserListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateUserListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportTaskComponentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportTaskComponentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportTaskComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFileTamperEventsRequest struct {
	*tchttp.BaseRequest

	// 事件ID

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
	// 是否物理删除：true、false

	IsHardDelete *bool `json:"IsHardDelete,omitempty" name:"IsHardDelete"`
	// 状态：0待处理、1已加白、2逻辑删除、3忽略、4已处理

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyFileTamperEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFileTamperEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefenseVulListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*VulVuls `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDefenseVulListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefenseVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateDownloadTokenRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRemoteUpdateDownloadTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateDownloadTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LicenseFunctions struct {
	// 授权功能描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 授权功能名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type AgentListItem struct {
	// 云用户appid

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	Guid *string `json:"Guid,omitempty" name:"Guid"`
	// 无

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否安装云镜，1：已安装 0：未安装

	InstalledCwp *int64 `json:"InstalledCwp,omitempty" name:"InstalledCwp"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例状态

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 实例的IPv6地址

	Ipv6Addresses *string `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`
	// 无

	IsPro *bool `json:"IsPro,omitempty" name:"IsPro"`
	// 最近同步时间

	LatestSyncTime *string `json:"LatestSyncTime,omitempty" name:"LatestSyncTime"`
	// 无

	LiveVersion *string `json:"LiveVersion,omitempty" name:"LiveVersion"`
	// cvm/bm/ecm

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 无

	Online *int64 `json:"Online,omitempty" name:"Online"`
	// 操作系统名称

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 操作系统类型

	OsType *int64 `json:"OsType,omitempty" name:"OsType"`
	// 实例主网卡的内网IP列表，多个IP通过“;”分隔

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 项目ID

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 实例主网卡的公网IP列表，多个IP通过“;”分隔

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// eg:ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例业务状态

	RestrictState *string `json:"RestrictState,omitempty" name:"RestrictState"`
	// 云用户UIN

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 私有网络ID

	VpcID *string `json:"VpcID,omitempty" name:"VpcID"`
}

type DescribeRemoteUpdateUserListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户列表

		UserList []*RemoteUpdateUser `json:"UserList,omitempty" name:"UserList"`
		// 符合条件的用户总数

		UserTotalCount *uint64 `json:"UserTotalCount,omitempty" name:"UserTotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemoteUpdateUserListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DNSEventDetail struct {
	// 无

	AccessCount *int64 `json:"AccessCount,omitempty" name:"AccessCount"`
	// 无

	Appid *int64 `json:"Appid,omitempty" name:"Appid"`
	// 无

	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 无

	GlobalRuleId *int64 `json:"GlobalRuleId,omitempty" name:"GlobalRuleId"`
	// 无

	Guid *string `json:"Guid,omitempty" name:"Guid"`
	// 主机IP，alphabet型，v4或v6

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	IsAlarmSend *int64 `json:"IsAlarmSend,omitempty" name:"IsAlarmSend"`
	// 无

	IsProVersion *int64 `json:"IsProVersion,omitempty" name:"IsProVersion"`
	// 无

	MergeTime *string `json:"MergeTime,omitempty" name:"MergeTime"`
	// 无

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 无

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 无

	ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`
	// 无

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Reference *string `json:"Reference,omitempty" name:"Reference"`
	// 无

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 无

	RuleSource *int64 `json:"RuleSource,omitempty" name:"RuleSource"`
	// 无

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 无

	UniqId *string `json:"UniqId,omitempty" name:"UniqId"`
	// 无

	Url *string `json:"Url,omitempty" name:"Url"`
	// 无

	UserRuleId *int64 `json:"UserRuleId,omitempty" name:"UserRuleId"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ExportVulAllEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVulAllEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulAllEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUpdateTaskRequest struct {
	*tchttp.BaseRequest

	// 目标配置ID

	PublishLinkId *int64 `json:"PublishLinkId,omitempty" name:"PublishLinkId"`
	// 备注

	Reverse *string `json:"Reverse,omitempty" name:"Reverse"`
	// 版本配置ID

	UpdateLinkId *int64 `json:"UpdateLinkId,omitempty" name:"UpdateLinkId"`
}

func (r *ModifyUpdateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUpdateTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulVulsSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVulVulsSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulVulsSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportJavaMemShellsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportJavaMemShellsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportJavaMemShellsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JavaMemShellPluginInfo struct {
	// Alias

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 错误日志

	ErrorLog *string `json:"ErrorLog,omitempty" name:"ErrorLog"`
	// HostIp

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 注入进程主类

	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`
	// 注入进程pid

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// Quuid

	QUUID *string `json:"QUUID,omitempty" name:"QUUID"`
	// 注入状态：0: 注入中, 1: 注入成功, 2: 插件超时, 3: 插件退出, 4: 注入失败 5: 软删除

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// WanIp

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
}

type DeleteBruteAttacksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBruteAttacksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetDatabaseCountRequest struct {
	*tchttp.BaseRequest

	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAssetDatabaseCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDatabaseCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVulDefenceEventRequest struct {
	*tchttp.BaseRequest

	// Ids

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 是否是物理删除

	IsHardDelete *bool `json:"IsHardDelete,omitempty" name:"IsHardDelete"`
}

func (r *DeleteVulDefenceEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVulDefenceEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetAppProcessListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Process []*AssetProcessListInfo `json:"Process,omitempty" name:"Process"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetAppProcessListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetAppProcessListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebFrameListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAssetWebFrameListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebFrameListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivilegeEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePrivilegeEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivilegeEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskInfo struct {
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Path *string `json:"Path,omitempty" name:"Path"`
	// 无

	Percent *float64 `json:"Percent,omitempty" name:"Percent"`
	// 无

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 无

	Type *string `json:"Type,omitempty" name:"Type"`
	// 无

	Used *uint64 `json:"Used,omitempty" name:"Used"`
}

type RemoteUpdateCurrentConfigInfo struct {
	// 该类型配置的自动更新开关，0：关，1：开

	AutoUpdateSwitch *int64 `json:"AutoUpdateSwitch,omitempty" name:"AutoUpdateSwitch"`
	// 配置信息

	ConfigInfo *RemoteUpdateConfigInfo `json:"ConfigInfo,omitempty" name:"ConfigInfo"`
	// 当该配置作为某一个用户的当前升级配置时，该字段表示配置针对该用户的状态，取值如下：  1：已最新  2：待升级  3：下发失败

	CurrentStatus *int64 `json:"CurrentStatus,omitempty" name:"CurrentStatus"`
}

type DescribeRemoteUpdateConfigNameListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 名称列表，按最近一次操作时间排序

		NameList []*string `json:"NameList,omitempty" name:"NameList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemoteUpdateConfigNameListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateConfigNameListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityDynamicsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全事件消息数组

		SecurityDynamics []*SecurityDynamic `json:"SecurityDynamics,omitempty" name:"SecurityDynamics"`
		// 无

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityDynamicsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityDynamicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportHostTrendRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 截至时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ExportHostTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportHostTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySpecimenStatusRequest struct {
	*tchttp.BaseRequest

	// Bwtype 21 11 30

	BlackWhiteType *int64 `json:"BlackWhiteType,omitempty" name:"BlackWhiteType"`
	// 样本

	Specimen []*Specimen `json:"Specimen,omitempty" name:"Specimen"`
	// Type 0 1

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *ModifySpecimenStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySpecimenStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetWebLocationBaseInfo struct {
	// 无

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	MainPath *string `json:"MainPath,omitempty" name:"MainPath"`
	// 无

	MainPathOwner *string `json:"MainPathOwner,omitempty" name:"MainPathOwner"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	PathCount *uint64 `json:"PathCount,omitempty" name:"PathCount"`
	// 无

	Permission *string `json:"Permission,omitempty" name:"Permission"`
	// 无

	Port *string `json:"Port,omitempty" name:"Port"`
	// 无

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 无

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 无

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	User *string `json:"User,omitempty" name:"User"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAssetTotalCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Types []*AssetKeyVal `json:"Types,omitempty" name:"Types"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetTotalCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetTotalCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetWebFrameListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAssetWebFrameListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetWebFrameListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomerLicense struct {
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`
	// 无

	DDL *string `json:"DDL,omitempty" name:"DDL"`
	// 无

	DomainMain *string `json:"DomainMain,omitempty" name:"DomainMain"`
	// 无

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
	// 无

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 无

	Functions []*string `json:"Functions,omitempty" name:"Functions"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 无

	Number *int64 `json:"Number,omitempty" name:"Number"`
	// 无

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 无

	PositionName *string `json:"PositionName,omitempty" name:"PositionName"`
	// 无

	ReleasedAt *string `json:"ReleasedAt,omitempty" name:"ReleasedAt"`
	// 无

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 无

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type DescribeUpdateTaskDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		PublishCount *int64 `json:"PublishCount,omitempty" name:"PublishCount"`
		// 无

		PublishRatio *float64 `json:"PublishRatio,omitempty" name:"PublishRatio"`
		// UpdateTask

		RestTaskDetail []*UpdateTaskDetailRsp `json:"RestTaskDetail,omitempty" name:"RestTaskDetail"`
		// 无

		SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
		// 无

		SuccessRatio *float64 `json:"SuccessRatio,omitempty" name:"SuccessRatio"`
		// 无

		UuidsCount *int64 `json:"UuidsCount,omitempty" name:"UuidsCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpdateTaskDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRemoteUpdateConfigAutoUpdateRequest struct {
	*tchttp.BaseRequest

	// 架构

	Architecture *uint64 `json:"Architecture,omitempty" name:"Architecture"`
	// 自动升级开关

	AutoUpdateFlag *uint64 `json:"AutoUpdateFlag,omitempty" name:"AutoUpdateFlag"`
	// 组件类型

	ComponentType *uint64 `json:"ComponentType,omitempty" name:"ComponentType"`
	// 系统

	OS *uint64 `json:"OS,omitempty" name:"OS"`
	// 升级类型

	UpdateType *uint64 `json:"UpdateType,omitempty" name:"UpdateType"`
	// 用户Id

	UserId *uint64 `json:"UserId,omitempty" name:"UserId"`
}

func (r *ModifyRemoteUpdateConfigAutoUpdateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRemoteUpdateConfigAutoUpdateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulComponent struct {
	// 作者

	Author *string `json:"Author,omitempty" name:"Author"`
	// 校验规则

	CheckRule *string `json:"CheckRule,omitempty" name:"CheckRule"`
	// 组件描述

	Descript *string `json:"Descript,omitempty" name:"Descript"`
	// 无

	DescriptEn *string `json:"DescriptEn,omitempty" name:"DescriptEn"`
	// 组件官网

	Homepage *string `json:"Homepage,omitempty" name:"Homepage"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 组件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	NameEn *string `json:"NameEn,omitempty" name:"NameEn"`
	// 无

	Num *int64 `json:"Num,omitempty" name:"Num"`
	// 操作者

	OperateUser *string `json:"OperateUser,omitempty" name:"OperateUser"`
	// 检测脚本

	PocFile *string `json:"PocFile,omitempty" name:"PocFile"`
	// 检测开关 1-开启 0-关闭

	Switch *int64 `json:"Switch,omitempty" name:"Switch"`
	// 组件类型 1-系统组件 2-Web组件

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 校验规则

	VerifyRule *string `json:"VerifyRule,omitempty" name:"VerifyRule"`
}

type DescribeAgentsListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*AgentListItem `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentsListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportFileSpecimenListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportFileSpecimenListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportFileSpecimenListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportLoginEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportLoginEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportLoginEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportKnowledgePackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportKnowledgePackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportKnowledgePackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBashEventsStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBashEventsStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBashEventsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTaskComponentsSwitchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTaskComponentsSwitchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTaskComponentsSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJavaMemShellListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*JavaMemShellEvent `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJavaMemShellListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJavaMemShellListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDetailRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeVulDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReverseShellEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteReverseShellEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReverseShellEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateConfigVersionListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 版本列表，从高到低降序排列

		VersionList []*string `json:"VersionList,omitempty" name:"VersionList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemoteUpdateConfigVersionListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateConfigVersionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateEvilMalwareRequest struct {
	*tchttp.BaseRequest

	// 无

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *UpdateEvilMalwareRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateEvilMalwareRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDNSEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDNSEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDNSEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulEventsStatusRequest struct {
	*tchttp.BaseRequest

	// 无

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
	// Status-1 0 1 2 3 4 5 10 11

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyVulEventsStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulEventsStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulTypeInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVulTypeInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulTypeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentAlarmRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAgentAlarmRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentAlarmRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceCountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVulDefenceCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAgentsInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAgentsInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAgentsInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUpdateConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// id、name

		Config []*SimpleSelect `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUpdateConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUpdateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulDefenceDetail struct {
	// 主机名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 攻击源ip地址所在城市

	City *string `json:"City,omitempty" name:"City"`
	// 事件发生次数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 创建事件时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// cve编号

	CveId *string `json:"CveId,omitempty" name:"CveId"`
	// 漏洞描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 漏洞ID相关的事件详情(json array格式 rasp特有)

	EventDetail *string `json:"EventDetail,omitempty" name:"EventDetail"`
	// 0: 尝试攻击(WeDetect) 1:尝试攻击成功(WeDetect) 2:rasp防御事件

	EventType *uint64 `json:"EventType,omitempty" name:"EventType"`
	// 主机失陷事件进程树(json格式 WeDetect特有)

	ExceptionPsTree *string `json:"ExceptionPsTree,omitempty" name:"ExceptionPsTree"`
	// 修复建议

	Fix *string `json:"Fix,omitempty" name:"Fix"`
	// 漏洞事件id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// MachineExtraInfo

	MachineExtraInfo *MachineExtraInfo `json:"MachineExtraInfo,omitempty" name:"MachineExtraInfo"`
	// ONLINE OFFLINE

	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`
	// 关联进程主类名

	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`
	// 更新事件时间

	MergeTime *string `json:"MergeTime,omitempty" name:"MergeTime"`
	// 攻击payload

	NetworkPayload *string `json:"NetworkPayload,omitempty" name:"NetworkPayload"`
	// 关联进程pid

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 内网ip

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 公网ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 主机quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 攻击源ip

	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
	// 攻击源端口

	SourcePort []*uint64 `json:"SourcePort,omitempty" name:"SourcePort"`
	// 堆栈信息(rasp特有)

	StackTrace *string `json:"StackTrace,omitempty" name:"StackTrace"`
	// 状态 0: 待处理 1:已防御 2:已处理 3: 已忽略 4: 已删除

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 漏洞名称

	VulName *string `json:"VulName,omitempty" name:"VulName"`
}

type DescribeAssetWebLocationListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetWebLocationListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebLocationListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTaskComponentsInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTaskComponentsInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTaskComponentsInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulPackage struct {
	// 无

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OperateUser *string `json:"OperateUser,omitempty" name:"OperateUser"`
	// 无

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	UpdateUrl *string `json:"UpdateUrl,omitempty" name:"UpdateUrl"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeUpdateTaskResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List *UpdateTaskResult `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpdateTaskResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetHostTotalCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Types []*AssetKeyVal `json:"Types,omitempty" name:"Types"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetHostTotalCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetHostTotalCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostSafetyStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全状态出参

		Data *DescribeHostSafetyStatusStatistic `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostSafetyStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostSafetyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetWebFrameBaseInfo struct {
	// 无

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 无

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 无

	Lang *string `json:"Lang,omitempty" name:"Lang"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	Path *string `json:"Path,omitempty" name:"Path"`
	// 无

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 无

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
}

type ExportAssetPlanTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetPlanTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetPlanTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportHostRiskTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportHostRiskTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportHostRiskTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TrojanSetting struct {
	// 是否启用自动隔离：0-未启用；1-启用

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// MD5黑名单列表，精确，以;分隔

	Md5Black *string `json:"Md5Black,omitempty" name:"Md5Black"`
}

type ExportAssetInitServiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetInitServiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetInitServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportReverseShellEventsRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportReverseShellEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportReverseShellEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineCategory struct {
	// 分类Id

	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 分类名称

	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`
	// 父分类ID,如果为0则没有父分类

	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" name:"ParentCategoryId"`
}

type DescribeAssetWebAppCountRequest struct {
	*tchttp.BaseRequest

	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAssetWebAppCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebAppCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebFrameCountRequest struct {
	*tchttp.BaseRequest

	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAssetWebFrameCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebFrameCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZKGroupInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyZKGroupInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyZKGroupInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PushCommandDesc struct {
	// 无

	Args []*ArgsDetail `json:"Args,omitempty" name:"Args"`
	// 无

	ArgsNum *int64 `json:"ArgsNum,omitempty" name:"ArgsNum"`
	// 无

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 无

	Disabled *int64 `json:"Disabled,omitempty" name:"Disabled"`
	// 无

	Group *string `json:"Group,omitempty" name:"Group"`
	// 无

	IsHighRisk *int64 `json:"IsHighRisk,omitempty" name:"IsHighRisk"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OptionType *int64 `json:"OptionType,omitempty" name:"OptionType"`
}

type ExportDNSEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出数据

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 导出任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportDNSEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDNSEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportJavaMemShellPluginsRequest struct {
	*tchttp.BaseRequest

	// FilterParams

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportJavaMemShellPluginsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportJavaMemShellPluginsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncAssetScanRequest struct {
	*tchttp.BaseRequest

	// 是否同步

	Sync *bool `json:"Sync,omitempty" name:"Sync"`
}

func (r *SyncAssetScanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SyncAssetScanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Specimen struct {
	// 无

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 无

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
}

type CreateRemoteUpdatePublishItemsRequest struct {
	*tchttp.BaseRequest

	// 配置（所属组件的）类型

	ComponentType *uint64 `json:"ComponentType,omitempty" name:"ComponentType"`
	// 配置名称

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
	// 配置版本

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
	// 操作人姓名

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 用户Id

	UserIds []*uint64 `json:"UserIds,omitempty" name:"UserIds"`
}

func (r *CreateRemoteUpdatePublishItemsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRemoteUpdatePublishItemsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentDetailRequest struct {
	*tchttp.BaseRequest

	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAgentDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopUpdateTaskRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *StopUpdateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopUpdateTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateConfigListRequest struct {
	*tchttp.BaseRequest

	// 是否降序，0：否，1：是

	Descending *uint64 `json:"Descending,omitempty" name:"Descending"`
	// 过滤，支持条件如下：
	// Version - 是否必填：否 - 版本
	// OS - 是否必填：否 - 系统
	// Type - 是否必填：是 - 类型
	// Architecture - 是否必填：否 - 架构
	// UserId - 是否必填：否 - 用户ID
	// ConfigName - 是否必填：否 - 配置名称

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序，支持如下传入：
	//
	// Name：按配置名称排序
	//
	// CreateTime：按配置创建时间排序
	//
	// UpdateTime：按配置更新时间排序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *DescribeRemoteUpdateConfigListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateConfigListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileTamperEventRuleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		FileTamperRuleDetail *FileTamperRuleDetail `json:"FileTamperRuleDetail,omitempty" name:"FileTamperRuleDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileTamperEventRuleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileTamperEventRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateUploadTempTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Token允许的路径

		AllowPath *string `json:"AllowPath,omitempty" name:"AllowPath"`
		// COS的AppID

		AppID *string `json:"AppID,omitempty" name:"AppID"`
		// 桶名称

		Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
		// 过期时间

		ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
		// 支撑组件标志0:支撑csp  1:cos

		IsCSPSupportSTS *int64 `json:"IsCSPSupportSTS,omitempty" name:"IsCSPSupportSTS"`
		// 地域

		Region *string `json:"Region,omitempty" name:"Region"`
		// 临时SecrectId

		SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
		// 临时SecrectKey

		SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
		// 临时密钥

		SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
		// 开始时间

		StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemoteUpdateUploadTempTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateUploadTempTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineCustomRuleIdName struct {
	// 自定义规则ID

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 自定义规则名字

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type BaselineItemInfo struct {
	// 检测项的修复方法

	FixMethod *string `json:"FixMethod,omitempty" name:"FixMethod"`
	// 检测项描述

	ItemDesc *string `json:"ItemDesc,omitempty" name:"ItemDesc"`
	// 基线检测项ID

	ItemId *int64 `json:"ItemId,omitempty" name:"ItemId"`
	// 检测项名字

	ItemName *string `json:"ItemName,omitempty" name:"ItemName"`
	// 危险等级

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 被引自定义规则信息

	RelatedCustomRuleInfo []*BaselineCustomRuleIdName `json:"RelatedCustomRuleInfo,omitempty" name:"RelatedCustomRuleInfo"`
	// 检测项所属规则的ID

	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`
	// 检测项所属规则名字

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 系统规则ID

	SysRuleId *int64 `json:"SysRuleId,omitempty" name:"SysRuleId"`
}

type Credentials struct {
	// 无

	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
	// 无

	TmpSecretID *string `json:"TmpSecretID,omitempty" name:"TmpSecretID"`
	// 无

	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
}

type DescribeAssetDatabaseListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Databases []*AssetDatabaseBaseInfo `json:"Databases,omitempty" name:"Databases"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetDatabaseListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDatabaseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFileTamperEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFileTamperEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFileTamperEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamInfo struct {
	// 无

	Data *string `json:"Data,omitempty" name:"Data"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
}

type ExportAssetWebServiceInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetWebServiceInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetWebServiceInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventTamperInfo struct {
	// 原文件下载

	DownloadOriginalFile *string `json:"DownloadOriginalFile,omitempty" name:"DownloadOriginalFile"`
	// 事件目录文件

	EventPath *string `json:"EventPath,omitempty" name:"EventPath"`
	// 事件状态

	EventStatus *uint64 `json:"EventStatus,omitempty" name:"EventStatus"`
	// 篡改事件时间

	EventTamperTime *string `json:"EventTamperTime,omitempty" name:"EventTamperTime"`
	// 事件类型

	EventType *uint64 `json:"EventType,omitempty" name:"EventType"`
	// 服务器ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 服务器

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 防护文件类型

	ProtectFileType *string `json:"ProtectFileType,omitempty" name:"ProtectFileType"`
	// 防护开启时间

	ProtectOpenTime *string `json:"ProtectOpenTime,omitempty" name:"ProtectOpenTime"`
	// 防护状态

	ProtectStatus *uint64 `json:"ProtectStatus,omitempty" name:"ProtectStatus"`
	// 自动恢复开关

	RecoverySwitch *uint64 `json:"RecoverySwitch,omitempty" name:"RecoverySwitch"`
}

type DescribeAgentPushTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 主键

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAgentPushTaskDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentPushTaskDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetMachineDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		MachineDetail *AssetMachineDetail `json:"MachineDetail,omitempty" name:"MachineDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetMachineDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetMachineDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetAppListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetAppListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetAppListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetProcessCountRequest struct {
	*tchttp.BaseRequest

	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAssetProcessCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetProcessCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineItemListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselineItem `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineItemListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineItemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetPlanTaskListRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetPlanTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPlanTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetPlanTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Tasks []*AssetPlanTaskBaseInfo `json:"Tasks,omitempty" name:"Tasks"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetPlanTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPlanTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRemoteUpdatePublishItemsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRemoteUpdatePublishItemsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRemoteUpdatePublishItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetMachineListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetMachineListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetMachineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRemoteUpdatePublishItemsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRemoteUpdatePublishItemsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRemoteUpdatePublishItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineItemInfoRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [LastTime|ItemCount|PassedItemCount|NotPassedItemCount|FirstTime]

	By *string `json:"By,omitempty" name:"By"`
	// <li>AppUserId - int64 - 是否必填：否 - 租户appId</li><li>HostName - string - 是否必填：否 - 主机名称</i><li>HostIp - string - 是否必填：否 - 主机Ip</i><li>DetectStatus - int - 是否必填：否 - 检测状态</li><li>StartTime - string - 是否必填：否 - 开始时间</li><li>EndTime - string - 是否必填：否 - 结束时间</li>

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBaselineItemInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineItemInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateConfigAutoUpdateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自动更新开关，取值：0：关，1：开

		AutoUpdateFlag *uint64 `json:"AutoUpdateFlag,omitempty" name:"AutoUpdateFlag"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemoteUpdateConfigAutoUpdateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateConfigAutoUpdateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRemoteUpdateTagRequest struct {
	*tchttp.BaseRequest

	// 标签信息

	TagInfo *RemoteUpdateTag `json:"TagInfo,omitempty" name:"TagInfo"`
}

func (r *CreateRemoteUpdateTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRemoteUpdateTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileTamperRuleCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机关联核心文件规则数量信息

		List []*FileTamperRuleCount `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileTamperRuleCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileTamperRuleCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulScanTaskRequest struct {
	*tchttp.BaseRequest

	// 无

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	TaskConfigId *int64 `json:"TaskConfigId,omitempty" name:"TaskConfigId"`
	// 无

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *CreateVulScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVulVulsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVulVulsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVulVulsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReverseShellEventRequest struct {
	*tchttp.BaseRequest

	// 无

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteReverseShellEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteReverseShellEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulVulsSwitchRequest struct {
	*tchttp.BaseRequest

	// 漏洞id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 0关，1开

	Switch *int64 `json:"Switch,omitempty" name:"Switch"`
}

func (r *ModifyVulVulsSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulVulsSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRemoteUpdateTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRemoteUpdateTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRemoteUpdateTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBruteForceEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBruteForceEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBruteForceEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerLicenseListRequest struct {
	*tchttp.BaseRequest

	// By

	By *string `json:"By,omitempty" name:"By"`
	// <li>AppUserId - int64 - 是否必填：否 - 租户appId</li><li>HostName - string - 是否必填：否 - 主机名称</i><li>HostIp - string - 是否必填：否 - 主机Ip</i><li>DetectStatus - int - 是否必填：否 - 检测状态</li><li>StartTime - string - 是否必填：否 - 开始时间</li><li>EndTime - string - 是否必填：否 - 结束时间</li>

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeCustomerLicenseListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerLicenseListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCwpProductUsagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// agent使用总数

		AgentUsedNum *int64 `json:"AgentUsedNum,omitempty" name:"AgentUsedNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCwpProductUsagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCwpProductUsagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJavaMemShellInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Info

		Info *JavaMemShellInfo `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJavaMemShellInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJavaMemShellInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLicenseInfoLicenseInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	PositionName *string `json:"PositionName,omitempty" name:"PositionName"`
	// 无

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyLicenseInfoLicenseInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLicenseInfoLicenseInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// 地域标志，如 ap-guangzhou，ap-shanghai，ap-beijing

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域代码，如 gz，sh，bj

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域中文名，如华南地区（广州），华东地区（上海金融），华北地区（北京）

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域英文名

	RegionNameEn *string `json:"RegionNameEn,omitempty" name:"RegionNameEn"`
}

type AssetWebLocationDetail struct {
	// 无

	Command *string `json:"Command,omitempty" name:"Command"`
	// 无

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 无

	MainPath *string `json:"MainPath,omitempty" name:"MainPath"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Port *string `json:"Port,omitempty" name:"Port"`
	// 无

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// 无

	SafeStatus *uint64 `json:"SafeStatus,omitempty" name:"SafeStatus"`
	// 无

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	User *string `json:"User,omitempty" name:"User"`
}

type DescribeAssetWebLocationCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		WebLocations []*AssetKeyVal `json:"WebLocations,omitempty" name:"WebLocations"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebLocationCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebLocationCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJavaMemShellPluginInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// java内存马插件列表

		List []*JavaMemShellPluginInfo `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJavaMemShellPluginInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJavaMemShellPluginInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventsBash struct {
	// 无

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
	// 无

	BashCmd *string `json:"BashCmd,omitempty" name:"BashCmd"`
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	DetectBy *int64 `json:"DetectBy,omitempty" name:"DetectBy"`
	// 无

	Exe *string `json:"Exe,omitempty" name:"Exe"`
	// 无

	ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
	// 主机IP，alphabet型，v4或v6

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 无

	Hostname *string `json:"Hostname,omitempty" name:"Hostname"`
	// 无

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 无

	IsProVersion *uint64 `json:"IsProVersion,omitempty" name:"IsProVersion"`
	// 无

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 无

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 无

	Platform *int64 `json:"Platform,omitempty" name:"Platform"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 无

	RuleLevel *int64 `json:"RuleLevel,omitempty" name:"RuleLevel"`
	// 无

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 无

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 无

	User *string `json:"User,omitempty" name:"User"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribePushCommandListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePushCommandListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePushCommandListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetPortCountRequest struct {
	*tchttp.BaseRequest

	// 无

	Port *uint64 `json:"Port,omitempty" name:"Port"`
}

func (r *DescribeAssetPortCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPortCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportDefenseVulInfoRequest struct {
	*tchttp.BaseRequest

	// 漏洞Id

	VulId *int64 `json:"VulId,omitempty" name:"VulId"`
}

func (r *ExportDefenseVulInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDefenseVulInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据

		Data []*HostTrendData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewStat struct {
	// 无

	AccountAllNum *int64 `json:"AccountAllNum,omitempty" name:"AccountAllNum"`
	// 无

	AgentsBasic *int64 `json:"AgentsBasic,omitempty" name:"AgentsBasic"`
	// 无

	AgentsFlagShip *int64 `json:"AgentsFlagShip,omitempty" name:"AgentsFlagShip"`
	// 无

	AgentsOffline *int64 `json:"AgentsOffline,omitempty" name:"AgentsOffline"`
	// 无

	AgentsOnline *int64 `json:"AgentsOnline,omitempty" name:"AgentsOnline"`
	// 无

	AgentsPro *int64 `json:"AgentsPro,omitempty" name:"AgentsPro"`
	// 无

	MachinesAll *int64 `json:"MachinesAll,omitempty" name:"MachinesAll"`
	// 无

	MachinesUninstalled *int64 `json:"MachinesUninstalled,omitempty" name:"MachinesUninstalled"`
}

type DescribeAssetWebServiceCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		WebServices []*AssetKeyVal `json:"WebServices,omitempty" name:"WebServices"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebServiceCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebServiceCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportSpecimenFileRequest struct {
	*tchttp.BaseRequest

	// Bwtype=21 22 23 11

	BlackWhiteType *int64 `json:"BlackWhiteType,omitempty" name:"BlackWhiteType"`
	// md5文件

	Md5File []*Md5File `json:"Md5File,omitempty" name:"Md5File"`
	// Type=0 1

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 无

	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`
}

func (r *ImportSpecimenFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportSpecimenFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginEventsRequest struct {
	*tchttp.BaseRequest

	// del ignore mis

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// 无

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
	// 指定周

	Week *string `json:"Week,omitempty" name:"Week"`
}

func (r *ModifyLoginEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLoginEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdatePublishLogsRequest struct {
	*tchttp.BaseRequest

	// 组件类型，取值如下
	//
	// 1：agent
	//
	// 2：malware的TAV引擎
	//
	// 3：漏洞库

	ComponentType *int64 `json:"ComponentType,omitempty" name:"ComponentType"`
	// 获取多少条目数的日志

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 用户ID

	UserId *uint64 `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeRemoteUpdatePublishLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdatePublishLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSocAuthListRequest struct {
	*tchttp.BaseRequest

	// 运营端租户appId，为了与租户端区分

	AppUserId *uint64 `json:"AppUserId,omitempty" name:"AppUserId"`
	// 排序字段方式 CreateTime

	By *string `json:"By,omitempty" name:"By"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeSocAuthListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSocAuthListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceMachineListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*VulDefenceMachine `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDefenceMachineListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceMachineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulDefenceSettingMachine struct {
	// AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAgentPushTaskResultRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 主键

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAgentPushTaskResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentPushTaskResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateUserNameUniqueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 名称是否已存在

		Existed *bool `json:"Existed,omitempty" name:"Existed"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemoteUpdateUserNameUniqueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateUserNameUniqueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDNSEventRequest struct {
	*tchttp.BaseRequest

	// 无

	Ids []*DNSIdInfo `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteDNSEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDNSEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportTasksRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ExportTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LicenseUpgrade struct {
	// 无

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 已售出授权数

	InquireNum *int64 `json:"InquireNum,omitempty" name:"InquireNum"`
	// 已绑定授权数总数

	UsedNum *int64 `json:"UsedNum,omitempty" name:"UsedNum"`
}

type DescribeBashEventsRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBashEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBashEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceEventRequest struct {
	*tchttp.BaseRequest

	// By

	By *string `json:"By,omitempty" name:"By"`
	// Name="AppId", "IP", "SourceIp", "VulName", "DefenseStatus", "EventType", "Status", "CreateTime","MergeTime"

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeVulDefenceEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportHostTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId []*string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportHostTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportHostTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartUpdateTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartUpdateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartUpdateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVulScanTaskRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteVulScanTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVulScanTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateConfigAutoUpdateRequest struct {
	*tchttp.BaseRequest

	// 架构

	Architecture *uint64 `json:"Architecture,omitempty" name:"Architecture"`
	// 配置所属组件类型类型

	ComponentType *uint64 `json:"ComponentType,omitempty" name:"ComponentType"`
	// 系统

	OS *uint64 `json:"OS,omitempty" name:"OS"`
	// 用户ID

	UserId *uint64 `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeRemoteUpdateConfigAutoUpdateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateConfigAutoUpdateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetType struct {
	// 无

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
}

type PrivilegeEscalationProcess struct {
	// 无

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
	// 无

	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	EffectiveGid *uint64 `json:"EffectiveGid,omitempty" name:"EffectiveGid"`
	// 无

	EffectiveUid *uint64 `json:"EffectiveUid,omitempty" name:"EffectiveUid"`
	// 无

	FindTime *string `json:"FindTime,omitempty" name:"FindTime"`
	// 无

	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`
	// 无

	Gid *uint64 `json:"Gid,omitempty" name:"Gid"`
	// 主机IP，alphabet型，v4或v6

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	IsProVersion *uint64 `json:"IsProVersion,omitempty" name:"IsProVersion"`
	// 无

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 无

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 无

	ParentProcGroup *string `json:"ParentProcGroup,omitempty" name:"ParentProcGroup"`
	// 无

	ParentProcName *string `json:"ParentProcName,omitempty" name:"ParentProcName"`
	// 无

	ParentProcPath *string `json:"ParentProcPath,omitempty" name:"ParentProcPath"`
	// 无

	ParentProcUser *string `json:"ParentProcUser,omitempty" name:"ParentProcUser"`
	// 无

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 无

	Ppid *uint64 `json:"Ppid,omitempty" name:"Ppid"`
	// 无

	ProcFilePrivilege *string `json:"ProcFilePrivilege,omitempty" name:"ProcFilePrivilege"`
	// 无

	ProcTree *string `json:"ProcTree,omitempty" name:"ProcTree"`
	// 无

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Sid *uint64 `json:"Sid,omitempty" name:"Sid"`
	// 无

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 无

	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`
	// 无

	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`
	// 无

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ExportAssetPortInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetPortInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetPortInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tags struct {
	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type DescribeCustomerLicenseListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*CustomerLicense `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomerLicenseListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerLicenseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpdateTaskConfigListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUpdateTaskConfigListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateTaskConfigListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpdateTaskListRequest struct {
	*tchttp.BaseRequest

	// "Name"

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUpdateTaskListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateTaskListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventsHostLogin struct {
	// 无

	Appid *int64 `json:"Appid,omitempty" name:"Appid"`
	// 无

	Area *string `json:"Area,omitempty" name:"Area"`
	// 无

	City *int64 `json:"City,omitempty" name:"City"`
	// 无

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 无

	Country *int64 `json:"Country,omitempty" name:"Country"`
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 0=云镜 1=洋葱

	DataFrom *int64 `json:"DataFrom,omitempty" name:"DataFrom"`
	// 攻击目标port

	DstPort *int64 `json:"DstPort,omitempty" name:"DstPort"`
	// 无

	EventType *int64 `json:"EventType,omitempty" name:"EventType"`
	// 预留未用，主机唯一标识

	Guid *string `json:"Guid,omitempty" name:"Guid"`
	// 主机ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 登录地是否异常

	IsRiskArea *int64 `json:"IsRiskArea,omitempty" name:"IsRiskArea"`
	// 登录IP是否异常

	IsRiskSrcIp *int64 `json:"IsRiskSrcIp,omitempty" name:"IsRiskSrcIp"`
	// 登录事件是否异常

	IsRiskTime *int64 `json:"IsRiskTime,omitempty" name:"IsRiskTime"`
	// 登录用户是否异常

	IsRiskUser *int64 `json:"IsRiskUser,omitempty" name:"IsRiskUser"`
	// 无

	LoginTime *string `json:"LoginTime,omitempty" name:"LoginTime"`
	// 无

	MgtIp *string `json:"MgtIp,omitempty" name:"MgtIp"`
	// 无

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 无

	OriginStatus *int64 `json:"OriginStatus,omitempty" name:"OriginStatus"`
	// 无

	Province *int64 `json:"Province,omitempty" name:"Province"`
	// 0默认未知   公有云-1，黑石物理机-2，黑石虚拟机-3， 云外-4

	Region *int64 `json:"Region,omitempty" name:"Region"`
	// 危险等级：0高危，1可疑，

	RiskLevel *int64 `json:"RiskLevel,omitempty" name:"RiskLevel"`
	// 登录异常分数，暂定大于7500分的为异常记录，

	Score *int64 `json:"Score,omitempty" name:"Score"`
	// 源ip

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 无

	SrcMachineName *string `json:"SrcMachineName,omitempty" name:"SrcMachineName"`
	// 0=刚入库未处理 1-正常登录，2-异常

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 主机唯一标识

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 无

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeRemoteUpdateTagListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRemoteUpdateTagListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateTagListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstallPackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstallPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstallPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFlagshipLicenseResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFlagshipLicenseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFlagshipLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulDefenceEventStatusRequest struct {
	*tchttp.BaseRequest

	// Ids

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
	// 状态 0: 待处理 1:已防御 2:已处理 3: 已忽略

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyVulDefenceEventStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulDefenceEventStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MachineSimple struct {
	// APPID

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 云标签信息

	CloudTags []*Tags `json:"CloudTags,omitempty" name:"CloudTags"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主机名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 是否在线：离OFFLINE、在线ONLINE

	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`
	// 机器所属专区类型 CVM 云服务器, BM 黑石, ECM 边缘计算, LH 轻量应用服务器 ,Other 混合云专区

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 主机系统。

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 主机IP。

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 防护版本 BASIC_VERSION 基础版， PRO_VERSION 专业版，Flagship 旗舰版，GENERAL_DISCOUNT 普惠版.

	ProtectType *string `json:"ProtectType,omitempty" name:"ProtectType"`
	// 主机外网IP。

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// CVM或BM机器唯一Uuid。

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 标签信息

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 云镜客户端唯一Uuid，若客户端长时间不在线将返回空字符。

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ModifyVulPackageInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVulPackageInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulPackageInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JavaMemShellPluginSetting struct {
	// 服务器名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 插件是否存在异常 0: 正常 1: 异常

	Exception *uint64 `json:"Exception,omitempty" name:"Exception"`
	// 服务器ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 服务器实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// javashell插件开关 0: 关闭 1: 开启

	JavaShellStatus *uint64 `json:"JavaShellStatus,omitempty" name:"JavaShellStatus"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// quuid

	QUUID *string `json:"QUUID,omitempty" name:"QUUID"`
	// uuid

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 服务器公网ip

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
}

type ModifyProLicenseCountRequest struct {
	*tchttp.BaseRequest
}

func (r *ModifyProLicenseCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProLicenseCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpecimenDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 样本信息

		SpecimenDetail *SpecimenDetail `json:"SpecimenDetail,omitempty" name:"SpecimenDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSpecimenDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpecimenDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentsPushResult struct {
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	ExeOut *string `json:"ExeOut,omitempty" name:"ExeOut"`
	// 无

	FinishCode *int64 `json:"FinishCode,omitempty" name:"FinishCode"`
	// 主键

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	IsBigTable *bool `json:"IsBigTable,omitempty" name:"IsBigTable"`
	// 无

	Message *string `json:"Message,omitempty" name:"Message"`
	// 无

	Result *string `json:"Result,omitempty" name:"Result"`
	// 无

	ResultType *int64 `json:"ResultType,omitempty" name:"ResultType"`
	// 无

	Retcode *int64 `json:"Retcode,omitempty" name:"Retcode"`
	// 无

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 任务id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 无

	TaskUuid *string `json:"TaskUuid,omitempty" name:"TaskUuid"`
	// 无

	TryTimes *int64 `json:"TryTimes,omitempty" name:"TryTimes"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 云镜uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAssetWebServiceInfoListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// Uuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAssetWebServiceInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebServiceInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDNSKnowledgeStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDNSKnowledgeStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDNSKnowledgeStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTrojanSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTrojanSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTrojanSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetDiskPartitionInfo struct {
	// 分区名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 挂载目录

	Path *string `json:"Path,omitempty" name:"Path"`
	// 分区使用率

	Percent *float64 `json:"Percent,omitempty" name:"Percent"`
	// 分区大小：单位G

	Size *uint64 `json:"Size,omitempty" name:"Size"`
	// 文件系统类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 已使用空间：单位G

	Used *uint64 `json:"Used,omitempty" name:"Used"`
}

type Md5File struct {
	// 无

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 无

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

type DescribeJavaMemShellPluginListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// java内存马插件列表

		List []*JavaMemShellPluginSetting `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJavaMemShellPluginListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJavaMemShellPluginListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetSystemPackageListRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters *FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *ExportAssetSystemPackageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetSystemPackageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishKnowledgeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishKnowledgeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishKnowledgeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoteUpdatePublishLog struct {
	// 架构

	Architecture *uint64 `json:"Architecture,omitempty" name:"Architecture"`
	// 配置名称

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
	// 配置版本

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
	// 日志ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 说明

	Instruction *string `json:"Instruction,omitempty" name:"Instruction"`
	// 系统

	OS *int64 `json:"OS,omitempty" name:"OS"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 发布时间

	PublishTime *int64 `json:"PublishTime,omitempty" name:"PublishTime"`
}

type DescribeAssetWebServiceProcessListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Process []*AssetProcessListInfo `json:"Process,omitempty" name:"Process"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebServiceProcessListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebServiceProcessListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportDNSEventsRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportDNSEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDNSEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteZKConfigInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteZKConfigInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteZKConfigInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetAppCountRequest struct {
	*tchttp.BaseRequest

	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAssetAppCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetAppCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetMachineListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportAssetMachineListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetMachineListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVulScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVulScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVulScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseAgentListRequest struct {
	*tchttp.BaseRequest

	// 无

	AppUserId *uint64 `json:"AppUserId,omitempty" name:"AppUserId"`
	// 升序：asc 降序：desc

	By *string `json:"By,omitempty" name:"By"`
	// "MachineName", "IP", "Uuid"

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段：create_time

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeLicenseAgentListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseAgentListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileTamperEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*FileTamperEvent `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileTamperEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileTamperEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskComponentsDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		ComponentDetail *VulComponent `json:"ComponentDetail,omitempty" name:"ComponentDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskComponentsDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskComponentsDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportDefenseVulVulsListRequest struct {
	*tchttp.BaseRequest

	// Filters："Name", "Level", "CvssScore", "CveId", "StartTime", "EndTime"

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ExportDefenseVulVulsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDefenseVulVulsListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulDefencePluginDetail struct {
	// 错误详情

	ErrorLog *string `json:"ErrorLog,omitempty" name:"ErrorLog"`
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// java主类

	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`
	// java进程pid

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 0: 注入中, 1: 注入成功, 2: 插件超时, 3: 插件退出, 4: 注入失败 5: 软删除

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ExportMalwareEventsRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 是否是全量

	IsAll *bool `json:"IsAll,omitempty" name:"IsAll"`
}

func (r *ExportMalwareEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportMalwareEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetCardsArray struct {
	// 无

	DnsServer *string `json:"DnsServer,omitempty" name:"DnsServer"`
	// 无

	GateWay *string `json:"GateWay,omitempty" name:"GateWay"`
	// 无

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 无

	Ipv6 *string `json:"Ipv6,omitempty" name:"Ipv6"`
	// 无

	Mac *string `json:"Mac,omitempty" name:"Mac"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateVulScanConfigInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVulScanConfigInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulScanConfigInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostRiskTrendRequest struct {
	*tchttp.BaseRequest

	// StartTime

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 无

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeHostRiskTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostRiskTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpdateGrayDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		UpdateGrayDetail *UpdateGray `json:"UpdateGrayDetail,omitempty" name:"UpdateGrayDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpdateGrayDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateGrayDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLicenseInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件下载链接

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 文件名称

		FileName *string `json:"FileName,omitempty" name:"FileName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLicenseInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLicenseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineRuleCategoryListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselineCategory `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineRuleCategoryListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineRuleCategoryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpdateTaskGrayListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// id、name

		Config []*SimpleSelect `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpdateTaskGrayListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateTaskGrayListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulAllEventsRequest struct {
	*tchttp.BaseRequest

	// Name=AppUserId、Status、Hostip、Name、Uuid

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 1:Web安全漏洞 2:系统组件漏洞 3:安全基线

	VulCategory *int64 `json:"VulCategory,omitempty" name:"VulCategory"`
}

func (r *ExportVulAllEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulAllEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateTagListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 标签列表

		TagList []*RemoteUpdateTag `json:"TagList,omitempty" name:"TagList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemoteUpdateTagListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateTagListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyJavaMemShellsStatusRequest struct {
	*tchttp.BaseRequest

	// 事件Id数组

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
	// 目标处理状态： 0 - 待处理 1 - 已加白 2 - 已删除 3 - 已忽略 4 - 已手动处理

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyJavaMemShellsStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyJavaMemShellsStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineItemListRequest struct {
	*tchttp.BaseRequest

	// 可选排序列: [LastTime|ItemCount|PassedItemCount|NotPassedItemCount|FirstTime]

	By *string `json:"By,omitempty" name:"By"`
	// <li>AppUserId - int64 - 是否必填：否 - 租户appId</li><li>HostName - string - 是否必填：否 - 主机名称</i><li>HostIp - string - 是否必填：否 - 主机Ip</i><li>DetectStatus - int - 是否必填：否 - 检测状态</li><li>StartTime - string - 是否必填：否 - 开始时间</li><li>EndTime - string - 是否必填：否 - 结束时间</li>

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式: [ASC:升序|DESC:降序]

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeBaselineItemListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineItemListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstallPackageRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Md5Path *string `json:"Md5Path,omitempty" name:"Md5Path"`
	// 无

	Md5PathId *string `json:"Md5PathId,omitempty" name:"Md5PathId"`
	// 无

	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`
	// 无

	PackagePath *string `json:"PackagePath,omitempty" name:"PackagePath"`
	// 无

	ReleaseId *string `json:"ReleaseId,omitempty" name:"ReleaseId"`
	// 无

	Reverse *string `json:"Reverse,omitempty" name:"Reverse"`
	// 无

	ScriptPath *string `json:"ScriptPath,omitempty" name:"ScriptPath"`
	// 无

	ScriptPathId *string `json:"ScriptPathId,omitempty" name:"ScriptPathId"`
}

func (r *ModifyInstallPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstallPackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishPocRequest struct {
	*tchttp.BaseRequest
}

func (r *PublishPocRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishPocRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetUserListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAssetUserListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetUserListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSocTamperProofDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件信息

		EventTamperInfo *EventTamperInfo `json:"EventTamperInfo,omitempty" name:"EventTamperInfo"`
		// 列表

		List []*RspSocTamperProofList `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 租户信息

		UserInfo *TamperProtectUserInfo `json:"UserInfo,omitempty" name:"UserInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSocTamperProofDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSocTamperProofDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstallPackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstallPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstallPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileTamperRuleCount struct {
	// 关联规则的数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 关联规则的名称（仅展示其中一条）

	Name *string `json:"Name,omitempty" name:"Name"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type AssetEnvBaseInfo struct {
	// 无

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 无

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	User *string `json:"User,omitempty" name:"User"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 无

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeJavaMemShellListRequest struct {
	*tchttp.BaseRequest

	// 升序:asc 降序:desc

	By *string `json:"By,omitempty" name:"By"`
	// "AppId", "Type", "Status", "IP", "CreateBeginTime", "CreateEndTime","RecentFoundBeginTime","RecentFoundEndTime","Uuid"支持搜索字段

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段 首次发现时间:create_time 最近扫描时间:recent_found_time

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeJavaMemShellListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJavaMemShellListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLicenseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateConfigMainInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置主要信息（没有文件信息）

		ConfigInfo *RemoteUpdateConfigInfo `json:"ConfigInfo,omitempty" name:"ConfigInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemoteUpdateConfigMainInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateConfigMainInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportMalwareEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportMalwareEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportMalwareEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRemoteUpdateConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置导出文件

		ConfigExportFiles []*RemoteUpdateConfigExportFile `json:"ConfigExportFiles,omitempty" name:"ConfigExportFiles"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportRemoteUpdateConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRemoteUpdateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTaskComponentsSwitchRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 检测开关 1-开启 0-关闭

	Switch *int64 `json:"Switch,omitempty" name:"Switch"`
}

func (r *ModifyTaskComponentsSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTaskComponentsSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishVulScanConfigRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *PublishVulScanConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishVulScanConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCosTempCredentialRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCosTempCredentialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCosTempCredentialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateConfigListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合入参条件并切片后的配置信息

		ConfigList []*RemoteUpdateConfigInfo `json:"ConfigList,omitempty" name:"ConfigList"`
		// 符合入参条件的配置总数（不是是ConfigInfoList数组的长度）

		ConfigTotalCount *uint64 `json:"ConfigTotalCount,omitempty" name:"ConfigTotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemoteUpdateConfigListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulTypeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*VulType `json:"List,omitempty" name:"List"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulTypeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstallPackagePublishListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstallPackagePublishListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstallPackagePublishListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportPrivilegeEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportPrivilegeEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportPrivilegeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTaskDetailRsp struct {
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	MaxVersion *string `json:"MaxVersion,omitempty" name:"MaxVersion"`
	// 无

	MinVersion *string `json:"MinVersion,omitempty" name:"MinVersion"`
	// 升级目标

	OperateStatus *int64 `json:"OperateStatus,omitempty" name:"OperateStatus"`
	// 关联的任务id（界面不展示）

	PublishLinkId *int64 `json:"PublishLinkId,omitempty" name:"PublishLinkId"`
	// 灰度数

	PublishQuantity *int64 `json:"PublishQuantity,omitempty" name:"PublishQuantity"`
	// 无

	PushTaskId *int64 `json:"PushTaskId,omitempty" name:"PushTaskId"`
	// 时间

	Reverse *string `json:"Reverse,omitempty" name:"Reverse"`
	// 执行时间

	RunTime *string `json:"RunTime,omitempty" name:"RunTime"`
	// 成功数

	SuccessQuantity *int64 `json:"SuccessQuantity,omitempty" name:"SuccessQuantity"`
	// 无

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 无

	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`
	// 成功数

	Time *string `json:"Time,omitempty" name:"Time"`
	// 无

	UpdateConfigCfgPath *string `json:"UpdateConfigCfgPath,omitempty" name:"UpdateConfigCfgPath"`
	// 无

	UpdateConfigName *string `json:"UpdateConfigName,omitempty" name:"UpdateConfigName"`
	// 无

	UpdateConfigPlatform *string `json:"UpdateConfigPlatform,omitempty" name:"UpdateConfigPlatform"`
	// 无

	UpdateConfigUpdateType *int64 `json:"UpdateConfigUpdateType,omitempty" name:"UpdateConfigUpdateType"`
	// 无

	UpdateConfigUpdateVersion *string `json:"UpdateConfigUpdateVersion,omitempty" name:"UpdateConfigUpdateVersion"`
	// 无

	UpdateGrayName *string `json:"UpdateGrayName,omitempty" name:"UpdateGrayName"`
	// 无

	UpdateGrayNum *int64 `json:"UpdateGrayNum,omitempty" name:"UpdateGrayNum"`
	// 无

	UpdateGrayStatic *int64 `json:"UpdateGrayStatic,omitempty" name:"UpdateGrayStatic"`
	// 关联的升级目标id

	UpdateLinkId *int64 `json:"UpdateLinkId,omitempty" name:"UpdateLinkId"`
	// 总数

	User *string `json:"User,omitempty" name:"User"`
	// 关联的升级包id

	VerifyStatus *int64 `json:"VerifyStatus,omitempty" name:"VerifyStatus"`
	// 成功数

	WorkLog *string `json:"WorkLog,omitempty" name:"WorkLog"`
	// 任务名称

	WorkStatus *int64 `json:"WorkStatus,omitempty" name:"WorkStatus"`
}

type ExportAssetSystemPackageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetSystemPackageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetSystemPackageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySpecimenStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySpecimenStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySpecimenStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetCoreModuleInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Module *AssetCoreModuleDetail `json:"Module,omitempty" name:"Module"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetCoreModuleInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetCoreModuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerLicenseDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		CustomerLicense []*CustomerLicense `json:"CustomerLicense,omitempty" name:"CustomerLicense"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomerLicenseDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerLicenseDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulVulsListRequest struct {
	*tchttp.BaseRequest

	// Filters："ListType", "Id", "Name", "Level", "VulCategory", "Switch"

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeVulVulsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulVulsListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateConfigVersionListRequest struct {
	*tchttp.BaseRequest

	// 组件类型，取值如下
	//
	// 1：agent
	//
	// 2：malware的TAV引擎
	//
	// 3：漏洞库

	ComponentType *uint64 `json:"ComponentType,omitempty" name:"ComponentType"`
}

func (r *DescribeRemoteUpdateConfigVersionListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateConfigVersionListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrojanSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 木马隔离详情

		TrojanSettingDetail *TrojanSetting `json:"TrojanSettingDetail,omitempty" name:"TrojanSettingDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrojanSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrojanSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTrojanSettingRequest struct {
	*tchttp.BaseRequest

	// 是否开启自动隔离1是0否

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 无

	Md5Black *string `json:"Md5Black,omitempty" name:"Md5Black"`
}

func (r *ModifyTrojanSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTrojanSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetTypesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineItemListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBaselineItemListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineItemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportJavaMemShellsRequest struct {
	*tchttp.BaseRequest

	// 升序asc 降序desc

	By *string `json:"By,omitempty" name:"By"`
	// "AppId", "Type", "Status", "IP", "CreateBeginTime", "CreateEndTime","RecentFoundBeginTime","RecentFoundEndTime"支持搜索条件

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportJavaMemShellsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportJavaMemShellsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRemoteUpdateTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新创建标签的Id

		TagId *uint64 `json:"TagId,omitempty" name:"TagId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRemoteUpdateTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRemoteUpdateTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateUploadTempTokenRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRemoteUpdateUploadTempTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateUploadTempTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefencePluginListRequest struct {
	*tchttp.BaseRequest

	// By

	By *string `json:"By,omitempty" name:"By"`
	// Name="Quuid", "Pid", "MainClass", "Status", "ErrorLog"

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeVulDefencePluginListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefencePluginListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetAppBaseInfo struct {
	// 无

	BinPath *string `json:"BinPath,omitempty" name:"BinPath"`
	// 无

	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`
	// 无

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 无

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 无

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	ProcessCount *uint64 `json:"ProcessCount,omitempty" name:"ProcessCount"`
	// 无

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 无

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
}

type UpgradeAgents struct {
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	//  实例主网卡的内网IP列表，多个IP通过“;”分隔

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// 实例主网卡的公网IP列表，多个IP通过“;”分隔

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ModifyVulDefenceSettingRequest struct {
	*tchttp.BaseRequest

	// Enable

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// Machines

	Machines []*VulDefenceSettingMachine `json:"Machines,omitempty" name:"Machines"`
}

func (r *ModifyVulDefenceSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulDefenceSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVulPackageRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteVulPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVulPackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetCoreModuleInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetCoreModuleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetCoreModuleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetSystemPackageListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetSystemPackageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetSystemPackageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentModel struct {
	// 无

	AgentId *int64 `json:"AgentId,omitempty" name:"AgentId"`
	// agent版本号

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// 主机别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 应用程序的配置文件内容

	AppConfig *string `json:"AppConfig,omitempty" name:"AppConfig"`
	// 租户appid

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
	// 地域

	Area *string `json:"Area,omitempty" name:"Area"`
	// BIOS名

	BiosInfo *string `json:"BiosInfo,omitempty" name:"BiosInfo"`
	// 服务器所见的客户端IP（可能是网关）

	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`
	// 采集过滤的配置

	CollectionConfig *string `json:"CollectionConfig,omitempty" name:"CollectionConfig"`
	// 保留字段，客户端可以提交任意字符标签保存到该字段

	Cookie *string `json:"Cookie,omitempty" name:"Cookie"`
	// cup名字+型号

	CpuId *string `json:"CpuId,omitempty" name:"CpuId"`
	// 该记录创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 0是云镜 1是洋葱

	DataFrom *int64 `json:"DataFrom,omitempty" name:"DataFrom"`
	// 主机分组（保留字段）

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 云镜guid

	Guid *string `json:"Guid,omitempty" name:"Guid"`
	// 硬盘名

	HdInfo *string `json:"HdInfo,omitempty" name:"HdInfo"`
	// 主机IP，alphabet型，v4或v6

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 主机名

	Hostname *string `json:"Hostname,omitempty" name:"Hostname"`
	// 无

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 安装key，已无用

	InstallKey *string `json:"InstallKey,omitempty" name:"InstallKey"`
	// 主机instance_id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// IP列表，逗号分割

	IpList *string `json:"IpList,omitempty" name:"IpList"`
	// 最后一次连接时间

	LastLogin *string `json:"LastLogin,omitempty" name:"LastLogin"`
	// YDLive最后活跃时间

	LiveLastTime *string `json:"LiveLastTime,omitempty" name:"LiveLastTime"`
	// YDLive版本

	LiveVersion *string `json:"LiveVersion,omitempty" name:"LiveVersion"`
	// mac地址，逗号分割

	MacList *string `json:"MacList,omitempty" name:"MacList"`
	// 主板名

	MbInfo *string `json:"MbInfo,omitempty" name:"MbInfo"`
	// 黑石ip

	MgtIp *string `json:"MgtIp,omitempty" name:"MgtIp"`
	// 该记录最后一次修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 1在线  0不在线

	Online *int64 `json:"Online,omitempty" name:"Online"`
	// 操作系统全称

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 平台信息(win32/win64/linux64)

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 登陆后，发送业务信息前的操作

	PreOperate *string `json:"PreOperate,omitempty" name:"PreOperate"`
	// 实例主网卡的内网IP列表，多个IP通过“;”分隔

	PrivateIpAddresses *string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 无

	PublicIpAddresses *string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 0默认未知   公有云-1，黑石物理机-2，黑石虚拟机-3， 云外-4

	Region *int64 `json:"Region,omitempty" name:"Region"`
	// 规则的版本（保留字段）

	RuleVersion *string `json:"RuleVersion,omitempty" name:"RuleVersion"`
	// 安全策略配置

	SecurityConfig *string `json:"SecurityConfig,omitempty" name:"SecurityConfig"`
	// 状态信息1有效，2主机已经释放

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 升级程序最后一次活跃时间

	UpdateLastTime *string `json:"UpdateLastTime,omitempty" name:"UpdateLastTime"`
	// 升级程序版本号

	UpdateVersion *string `json:"UpdateVersion,omitempty" name:"UpdateVersion"`
	// 用户id

	Userid *string `json:"Userid,omitempty" name:"Userid"`
	// 云镜uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 无

	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeUpdateConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// UpdateTask

		List []*UpdateConfig `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpdateConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUpdateConfigRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteUpdateConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUpdateConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHostTrendRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 截至时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeHostTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHostTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZKGroupFullListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeZKGroupFullListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZKGroupFullListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulEventsRequest struct {
	*tchttp.BaseRequest

	// Name=AppUserId、Status、Hostip、Name、Uuid

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 1:Web安全漏洞 2:系统组件漏洞 3:安全基线

	VulCategory *int64 `json:"VulCategory,omitempty" name:"VulCategory"`
}

func (r *ExportVulEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRemoteUpdateTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRemoteUpdateTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRemoteUpdateTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetHostTotalCountRequest struct {
	*tchttp.BaseRequest

	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetHostTotalCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetHostTotalCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetJarListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Jars []*AssetJarBaseInfo `json:"Jars,omitempty" name:"Jars"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetJarListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetJarListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetWebLocationListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *ExportAssetWebLocationListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetWebLocationListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SeparateMalwareResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SeparateMalwareResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SeparateMalwareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PathInfo struct {
	// 无

	Group *string `json:"Group,omitempty" name:"Group"`
	// 无

	Permission *string `json:"Permission,omitempty" name:"Permission"`
	// 无

	RealPath *string `json:"RealPath,omitempty" name:"RealPath"`
	// 无

	User *string `json:"User,omitempty" name:"User"`
	// 无

	VirtualPath *string `json:"VirtualPath,omitempty" name:"VirtualPath"`
}

type DeleteUpdateTaskRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteUpdateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUpdateTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefenseVulListRequest struct {
	*tchttp.BaseRequest

	// Filters："Name", "Level", "CvssScore", "CveId", "StartTime", "EndTime"

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDefenseVulListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefenseVulListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventsRiskDns struct {
	// 无

	AccessCount *int64 `json:"AccessCount,omitempty" name:"AccessCount"`
	// 云用户唯一标识

	Appid *int64 `json:"Appid,omitempty" name:"Appid"`
	// 无

	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	Description *string `json:"Description,omitempty" name:"Description"`
	// 无

	GlobalRuleId *int64 `json:"GlobalRuleId,omitempty" name:"GlobalRuleId"`
	// 无

	Guid *string `json:"Guid,omitempty" name:"Guid"`
	// 主机IP，alphabet型，v4或v6

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 无

	IsAlarmSend *int64 `json:"IsAlarmSend,omitempty" name:"IsAlarmSend"`
	// 无

	IsProVersion *int64 `json:"IsProVersion,omitempty" name:"IsProVersion"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// OFFLINE、ONLINE

	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`
	// 外网ip

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	MergeTime *string `json:"MergeTime,omitempty" name:"MergeTime"`
	// 无

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 无

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 无

	ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`
	// 无

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 参考地址

	Reference *string `json:"Reference,omitempty" name:"Reference"`
	// 参考地址数组

	References []*string `json:"References,omitempty" name:"References"`
	// 无

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 建议方案

	SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`
	// 标签特性

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 无

	UniqId *string `json:"UniqId,omitempty" name:"UniqId"`
	// 无

	Url *string `json:"Url,omitempty" name:"Url"`
	// 无

	UserRuleId *int64 `json:"UserRuleId,omitempty" name:"UserRuleId"`
	// 主机唯一标识

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAssetDatabaseCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Databases []*AssetKeyVal `json:"Databases,omitempty" name:"Databases"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetDatabaseCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDatabaseCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaselineItemInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselineItemInfo `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineItemInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineItemInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAgentsListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出数据

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAgentsListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAgentsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetJarListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetJarListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetJarListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishPocResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务描述

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishPocResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishPocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulScanTaskNewRequest struct {
	*tchttp.BaseRequest

	// windows：1 linux：2

	ConfigType *int64 `json:"ConfigType,omitempty" name:"ConfigType"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *CreateVulScanTaskNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulScanTaskNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpdateTaskRequest struct {
	*tchttp.BaseRequest

	// "Name"

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUpdateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportLoginEventsRequest struct {
	*tchttp.BaseRequest

	// FilterParams

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportLoginEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportLoginEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportFileTamperEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出任务ID 可通过ExportTasks接口下载

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportFileTamperEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportFileTamperEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetPortInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Ports []*AssetPortBaseInfo `json:"Ports,omitempty" name:"Ports"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetPortInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPortInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// License授权信息

		License *License `json:"License,omitempty" name:"License"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAgentPushTaskResultRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *ExportAgentPushTaskResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAgentPushTaskResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Item struct {
	// Id

	ItemId *int64 `json:"ItemId,omitempty" name:"ItemId"`
	// 名称

	ItemName *string `json:"ItemName,omitempty" name:"ItemName"`
}

type DescribePrivilegeEventsRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribePrivilegeEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivilegeEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBaselineHostDetectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBaselineHostDetectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBaselineHostDetectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefenceEventDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *VulDefenceDetail `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDefenceEventDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefenceEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateUserCurrentConfigsRequest struct {
	*tchttp.BaseRequest

	// 配置（对应组件的）类型

	ComponentType *uint64 `json:"ComponentType,omitempty" name:"ComponentType"`
	// 是否降序排列，升序：0，降序：1

	Descending *int64 `json:"Descending,omitempty" name:"Descending"`
	// 是否获取所有类型的配置，1：是，0：否

	GetAll *int64 `json:"GetAll,omitempty" name:"GetAll"`
	// 限制数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 按配置的哪一个字段排序，目前支持：
	// LatestPullTime - string - 按最近一次拉取时间排序

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 用户ID

	UserId *uint64 `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeRemoteUpdateUserCurrentConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateUserCurrentConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulTypeListRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeVulTypeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulTypeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulScanConfig struct {
	// 无

	ArmFrameworkPackageId *int64 `json:"ArmFrameworkPackageId,omitempty" name:"ArmFrameworkPackageId"`
	// 无

	ArmPythonPackageId *int64 `json:"ArmPythonPackageId,omitempty" name:"ArmPythonPackageId"`
	// 无

	ConfigType *int64 `json:"ConfigType,omitempty" name:"ConfigType"`
	// 无

	FrameworkPackageId *int64 `json:"FrameworkPackageId,omitempty" name:"FrameworkPackageId"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OperateUser *string `json:"OperateUser,omitempty" name:"OperateUser"`
	// 无

	PocPackageId *int64 `json:"PocPackageId,omitempty" name:"PocPackageId"`
	// 无

	Publish *int64 `json:"Publish,omitempty" name:"Publish"`
	// 无

	PythonPackageId *int64 `json:"PythonPackageId,omitempty" name:"PythonPackageId"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	UpdateUrl *string `json:"UpdateUrl,omitempty" name:"UpdateUrl"`
	// 无

	XmlPath *string `json:"XmlPath,omitempty" name:"XmlPath"`
}

type AddPushTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddPushTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddPushTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVulTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVulTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVulTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBruteAttackTopRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 获取top几

	Top *uint64 `json:"Top,omitempty" name:"Top"`
}

func (r *DescribeBruteAttackTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBruteAttackTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelatedMachineTamperList struct {
	// 防护文件类型

	ProtectFileType *string `json:"ProtectFileType,omitempty" name:"ProtectFileType"`
	// 防护开启时间

	ProtectOpenTime *string `json:"ProtectOpenTime,omitempty" name:"ProtectOpenTime"`
	// 防护目录地址

	ProtectPathAddr *string `json:"ProtectPathAddr,omitempty" name:"ProtectPathAddr"`
	// 防护目录名称

	ProtectPathName *string `json:"ProtectPathName,omitempty" name:"ProtectPathName"`
	// 防护状态

	ProtectStatus *uint64 `json:"ProtectStatus,omitempty" name:"ProtectStatus"`
	// 恢复开关

	RecoverySwitch *uint64 `json:"RecoverySwitch,omitempty" name:"RecoverySwitch"`
	// 防护ID

	SocTamperProtectId *uint64 `json:"SocTamperProtectId,omitempty" name:"SocTamperProtectId"`
}

type DescribeAgentAlarmResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警信息：OverUsed SellOut AboutToSellOut

		Alarm *string `json:"Alarm,omitempty" name:"Alarm"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentAlarmResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebAppCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		WebApps []*AssetKeyVal `json:"WebApps,omitempty" name:"WebApps"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebAppCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebAppCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUpdateTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// id、name

		Config []*SimpleSelect `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUpdateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUpdateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetWebAppBaseInfo struct {
	// 无

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 无

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 无

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	PluginCount *uint64 `json:"PluginCount,omitempty" name:"PluginCount"`
	// 无

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	RootPath *string `json:"RootPath,omitempty" name:"RootPath"`
	// 无

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 无

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
	// 无

	VirtualPath *string `json:"VirtualPath,omitempty" name:"VirtualPath"`
}

type DescribeLicenseHistoryReleasedResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*CustomerLicense `json:"List,omitempty" name:"List"`
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseHistoryReleasedResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseHistoryReleasedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAgentsInfoListRequest struct {
	*tchttp.BaseRequest

	// Filters："Online"

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportAgentsInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAgentsInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulScanTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVulScanTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentPushTaskDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命令推送信息

		AgentPushTaskDetail *AgentsPushResult `json:"AgentPushTaskDetail,omitempty" name:"AgentPushTaskDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentPushTaskDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentPushTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetTotalCountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetTotalCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetTotalCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Data *OverviewStat `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOverviewStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOverviewStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportKnowledgeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportKnowledgeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportKnowledgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulPackageInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 无

	Types *int64 `json:"Types,omitempty" name:"Types"`
}

func (r *ModifyVulPackageInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulPackageInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBruteAttackTopRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 获取top几

	Top *uint64 `json:"Top,omitempty" name:"Top"`
}

func (r *ExportBruteAttackTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBruteAttackTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyJavaMemShellPluginSwitchRequest struct {
	*tchttp.BaseRequest

	// 无

	JavaShellStatus *int64 `json:"JavaShellStatus,omitempty" name:"JavaShellStatus"`
	// 无

	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

func (r *ModifyJavaMemShellPluginSwitchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyJavaMemShellPluginSwitchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ArgsDetail struct {
	// 无

	Default *string `json:"Default,omitempty" name:"Default"`
	// 无

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type DeleteJavaMemShellsRequest struct {
	*tchttp.BaseRequest

	// id

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
	// 是否是物理删除（运营端删除）

	IsHardDelete *bool `json:"IsHardDelete,omitempty" name:"IsHardDelete"`
}

func (r *DeleteJavaMemShellsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteJavaMemShellsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseListRequest struct {
	*tchttp.BaseRequest

	// 升序asc 降序 desc

	By *string `json:"By,omitempty" name:"By"`
	// "AppUserId"

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序字段：used_num inquire_num

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeLicenseListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRemoteUpdateConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建成功的配置Id

		ConfigId *uint64 `json:"ConfigId,omitempty" name:"ConfigId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRemoteUpdateConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRemoteUpdateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetEnvListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Envs []*AssetEnvBaseInfo `json:"Envs,omitempty" name:"Envs"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetEnvListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetEnvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetKeyVal struct {
	// 无

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 无

	Key *string `json:"Key,omitempty" name:"Key"`
	// 无

	NewCount *int64 `json:"NewCount,omitempty" name:"NewCount"`
	// 无

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type FileTamperRule struct {
	// 执行动作 跳过：skip，告警：alert

	Action *string `json:"Action,omitempty" name:"Action"`
	// 文件监控内容 write read read-write

	FileAction *string `json:"FileAction,omitempty" name:"FileAction"`
	// 进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 被访问文件路径

	Target *string `json:"Target,omitempty" name:"Target"`
}

type DeleteVulVulsRequest struct {
	*tchttp.BaseRequest

	// 漏洞id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteVulVulsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVulVulsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetUserInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetUserInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetUserInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerLicenseDetailsRequest struct {
	*tchttp.BaseRequest

	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeCustomerLicenseDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerLicenseDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulVulsListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVulVulsListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulVulsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDNSEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*EventsRiskDns `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDNSEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDNSEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpdateGrayRequest struct {
	*tchttp.BaseRequest

	// PublishName

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUpdateGrayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateGrayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportSecurityDynamicsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportSecurityDynamicsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportSecurityDynamicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetUserListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 无

		Users []*AssetUserBaseInfo `json:"Users,omitempty" name:"Users"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetUserListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulPackageFullListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*VulPackage `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulPackageFullListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulPackageFullListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVulScanConfigInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	FrameworkPackageId *int64 `json:"FrameworkPackageId,omitempty" name:"FrameworkPackageId"`
	// 无

	PocPackageId *int64 `json:"PocPackageId,omitempty" name:"PocPackageId"`
	// Publish：0 1

	Publish *int64 `json:"Publish,omitempty" name:"Publish"`
}

func (r *CreateVulScanConfigInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVulScanConfigInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentInfoRsp struct {
	// 无

	Agent *AgentModel `json:"Agent,omitempty" name:"Agent"`
	// 无

	AgentStatus *AgentStatus `json:"AgentStatus,omitempty" name:"AgentStatus"`
	// 无

	CVMDetailInfo *CvmDetailInfo `json:"CVMDetailInfo,omitempty" name:"CVMDetailInfo"`
	// 无

	HostComponents []*string `json:"HostComponents,omitempty" name:"HostComponents"`
	// 无

	IsProVersion *int64 `json:"IsProVersion,omitempty" name:"IsProVersion"`
	// agent最近活跃时间

	LastActiveTime *string `json:"LastActiveTime,omitempty" name:"LastActiveTime"`
	// 无

	MachineInfo *MachineInfo `json:"MachineInfo,omitempty" name:"MachineInfo"`
	// 无

	WebComponents []*string `json:"WebComponents,omitempty" name:"WebComponents"`
}

type RemoteUpdateConfigInfo struct {
	// 配置适用的框架版本范围

	ApplyFrameVersion *string `json:"ApplyFrameVersion,omitempty" name:"ApplyFrameVersion"`
	// 配置应用的目标架构

	Architecture *uint64 `json:"Architecture,omitempty" name:"Architecture"`
	// 配置（所属组件的）类型

	ComponentType *uint64 `json:"ComponentType,omitempty" name:"ComponentType"`
	// 创建时间

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建人姓名

	CreatorName *string `json:"CreatorName,omitempty" name:"CreatorName"`
	// 环境类型

	EnvironmentType *uint64 `json:"EnvironmentType,omitempty" name:"EnvironmentType"`
	// 配置涉及的文件列表

	FileList []*RemoteUpdateFile `json:"FileList,omitempty" name:"FileList"`
	// 配置Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 最近一次操作时间

	LatestOperateTime *uint64 `json:"LatestOperateTime,omitempty" name:"LatestOperateTime"`
	// 最近一次操作人姓名

	LatestOperatorName *string `json:"LatestOperatorName,omitempty" name:"LatestOperatorName"`
	// 当该配置作为某一个用户的当前升级配置时，该字段表示最新拉取时间

	LatestPullTime *uint64 `json:"LatestPullTime,omitempty" name:"LatestPullTime"`
	// 配置名称（在某个配置版本下，该名称唯一）

	Name *string `json:"Name,omitempty" name:"Name"`
	// 配置应用的目标系统

	OS *uint64 `json:"OS,omitempty" name:"OS"`
	// 当配置针对某一个用户处于发布中时，该字段是发布项的Id。否则为0

	PublishItemId *uint64 `json:"PublishItemId,omitempty" name:"PublishItemId"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 升级类型，大部分为1即默认类型。某些组件要区分升级类型，例如agent

	UpdateType *uint64 `json:"UpdateType,omitempty" name:"UpdateType"`
	// 该配置是否可以被自动更新使用。0：否，1：是

	UsedByAutoUpdate *uint64 `json:"UsedByAutoUpdate,omitempty" name:"UsedByAutoUpdate"`
	// 配置版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeAssetAppCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Apps []*AssetKeyVal `json:"Apps,omitempty" name:"Apps"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetAppCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetAppCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetUserCountRequest struct {
	*tchttp.BaseRequest

	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAssetUserCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetUserCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpdateTaskConfigListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// id、name

		Config []*SimpleSelect `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpdateTaskConfigListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateTaskConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CredentialError struct {
	// 无

	Code *string `json:"Code,omitempty" name:"Code"`
	// 无

	Message *string `json:"Message,omitempty" name:"Message"`
}

type DeleteZKConfigInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteZKConfigInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteZKConfigInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishVulScanConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishVulScanConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishVulScanConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetJarBaseInfo struct {
	// 无

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	Path *string `json:"Path,omitempty" name:"Path"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 无

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
}

type UpdateTasksLinkGray struct {
	// 无

	AllUpdate *int64 `json:"AllUpdate,omitempty" name:"AllUpdate"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OperateUser *string `json:"OperateUser,omitempty" name:"OperateUser"`
	// 无

	PackageId *int64 `json:"PackageId,omitempty" name:"PackageId"`
	// 无

	PublishLinkId *int64 `json:"PublishLinkId,omitempty" name:"PublishLinkId"`
	// 无

	PushTaskId *int64 `json:"PushTaskId,omitempty" name:"PushTaskId"`
	// 无

	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
	// 无

	TaskCount *int64 `json:"TaskCount,omitempty" name:"TaskCount"`
	// 无

	UpdateGrayName *string `json:"UpdateGrayName,omitempty" name:"UpdateGrayName"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeLicenseFunctionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 授权功能

		LicenseFunctions []*LicenseFunctions `json:"LicenseFunctions,omitempty" name:"LicenseFunctions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseFunctionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseFunctionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteZKGroupInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteZKGroupInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteZKGroupInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulPackageFullListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVulPackageFullListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulPackageFullListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoteUpdateFile struct {
	// 文件在COS上的对象

	CosObject *string `json:"CosObject,omitempty" name:"CosObject"`
	// 文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件类型，不同的组件用到的文件类型不同

	FileType *int64 `json:"FileType,omitempty" name:"FileType"`
	// 文件Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 文件MD5值

	MD5 *string `json:"MD5,omitempty" name:"MD5"`
	// 文件大小

	Size *uint64 `json:"Size,omitempty" name:"Size"`
}

type ExportAssetWebFrameListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetWebFrameListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetWebFrameListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBashEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBashEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBashEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportRemoteUpdateConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导入成功后的配置ID

		ConfigId *uint64 `json:"ConfigId,omitempty" name:"ConfigId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportRemoteUpdateConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportRemoteUpdateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulDefenceMainSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVulDefenceMainSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulDefenceMainSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBashEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*EventsBash `json:"List,omitempty" name:"List"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBashEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBashEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadPushResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadPushResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadPushResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateInstallPackage struct {
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Md5Path *string `json:"Md5Path,omitempty" name:"Md5Path"`
	// 无

	PackagePath *string `json:"PackagePath,omitempty" name:"PackagePath"`
	// 无

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 无

	Publish *int64 `json:"Publish,omitempty" name:"Publish"`
	// 无

	ReleaseId *string `json:"ReleaseId,omitempty" name:"ReleaseId"`
	// 无

	Reverse *string `json:"Reverse,omitempty" name:"Reverse"`
	// 无

	ScriptPath *string `json:"ScriptPath,omitempty" name:"ScriptPath"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	UpdateUrl *string `json:"UpdateUrl,omitempty" name:"UpdateUrl"`
	// 无

	UrlPath *string `json:"UrlPath,omitempty" name:"UrlPath"`
	// 无

	User *string `json:"User,omitempty" name:"User"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
}

type ModifyStatisticRequest struct {
	*tchttp.BaseRequest
}

func (r *ModifyStatisticRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStatisticRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostTrendData struct {
	// 分类

	Category *string `json:"Category,omitempty" name:"Category"`
	// 时间

	Date *string `json:"Date,omitempty" name:"Date"`
	// 内容

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type DescribeDNSEventsRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDNSEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDNSEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewStatRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOverviewStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOverviewStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulScanTask struct {
	// 无

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OperateUser *string `json:"OperateUser,omitempty" name:"OperateUser"`
	// 无

	PushTaskId *int64 `json:"PushTaskId,omitempty" name:"PushTaskId"`
	// 无

	TaskConfigId *int64 `json:"TaskConfigId,omitempty" name:"TaskConfigId"`
	// 无

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	UpdateUrl *string `json:"UpdateUrl,omitempty" name:"UpdateUrl"`
	// 无

	VulScanConfigType *int64 `json:"VulScanConfigType,omitempty" name:"VulScanConfigType"`
}

type RemoteUpdateTag struct {
	// 标签Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 标签名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 使用该标签的用户的ID列表

	UserIds []*uint64 `json:"UserIds,omitempty" name:"UserIds"`
}

type StopUpdateTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopUpdateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopUpdateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetPlanTaskBaseInfo struct {
	// 无

	Command *string `json:"Command,omitempty" name:"Command"`
	// 无

	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`
	// 无

	Cycle *string `json:"Cycle,omitempty" name:"Cycle"`
	// 无

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 无

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	User *string `json:"User,omitempty" name:"User"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ExportVulTopRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 获取top几

	Top *uint64 `json:"Top,omitempty" name:"Top"`
}

func (r *ExportVulTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRemoteUpdateConfigAutoUpdateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRemoteUpdateConfigAutoUpdateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRemoteUpdateConfigAutoUpdateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVulDefenceMainSettingRequest struct {
	*tchttp.BaseRequest

	// Enable

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

func (r *ModifyVulDefenceMainSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVulDefenceMainSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JavaMemShellInfo struct {
	// 注释

	Annotations *string `json:"Annotations,omitempty" name:"Annotations"`
	// java进程命令行参数

	Args *string `json:"Args,omitempty" name:"Args"`
	// java内存马二进制代码(base64)

	ClassContent *string `json:"ClassContent,omitempty" name:"ClassContent"`
	// java内存马反编译代码

	ClassContentPretty *string `json:"ClassContentPretty,omitempty" name:"ClassContentPretty"`
	// java加载器类名

	ClassLoaderName *string `json:"ClassLoaderName,omitempty" name:"ClassLoaderName"`
	// 类名

	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`
	// 首次发现时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 说明

	Description *string `json:"Description,omitempty" name:"Description"`
	// 事件描述

	EventDescription *string `json:"EventDescription,omitempty" name:"EventDescription"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 容器名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 容器状态：RUNNING,STOPPED,SHUTDOWN...

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 继承的接口

	Interfaces *string `json:"Interfaces,omitempty" name:"Interfaces"`
	// 类文件MD5

	MD5 *string `json:"MD5,omitempty" name:"MD5"`
	// 进程pid

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 内网IP

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// java进程路径

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 公共ip

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// 最近检测时间

	RecentFoundTime *string `json:"RecentFoundTime,omitempty" name:"RecentFoundTime"`
	// 安全建议

	SecurityAdvice *string `json:"SecurityAdvice,omitempty" name:"SecurityAdvice"`
	// 处理状态  0 -- 待处理 1 -- 已加白 2 -- 已删除 3 - 已忽略  4 - 已手动处理

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 父类名

	SuperClassName *string `json:"SuperClassName,omitempty" name:"SuperClassName"`
	// 内存马类型  0:Filter型 1:Listener型 2:Servlet型 3:Interceptors型 4:Agent型 5:其他

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type CreateRemoteUpdateConfigRequest struct {
	*tchttp.BaseRequest

	// 配置信息

	ConfigInfo *RemoteUpdateConfigInfo `json:"ConfigInfo,omitempty" name:"ConfigInfo"`
}

func (r *CreateRemoteUpdateConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRemoteUpdateConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileTamperEventsRequest struct {
	*tchttp.BaseRequest

	// By

	By *string `json:"By,omitempty" name:"By"`
	// 支持条件：AppId、Level、Status、Quuids、Uuids、FileAction、CreateBeginTime、CreateEndTime、ModifyBeginTime、ModifyEndTime、RuleName、HostIp、RuleCategory

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeFileTamperEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileTamperEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 导出数据内容，用于前端生成文件下载

		Content *string `json:"Content,omitempty" name:"Content"`
		// 下载链接

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// PENDING：正在生成下载链接，FINISHED：下载链接已生成，ERROR：网络异常等异常情况

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MachineTag struct {
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Rid *int64 `json:"Rid,omitempty" name:"Rid"`
	// 无

	TagId *uint64 `json:"TagId,omitempty" name:"TagId"`
}

type UpdateEvilMalwareResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateEvilMalwareResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateEvilMalwareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TamperProtectUserInfo struct {
	// 云用户AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 网页防篡改已购授权数

	AuthNum *uint64 `json:"AuthNum,omitempty" name:"AuthNum"`
	// 基础版数

	BasicNum *uint64 `json:"BasicNum,omitempty" name:"BasicNum"`
	// 用户名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 专业版数

	ProfessionalNum *uint64 `json:"ProfessionalNum,omitempty" name:"ProfessionalNum"`
	// 服务器总数

	ServerCnt *uint64 `json:"ServerCnt,omitempty" name:"ServerCnt"`
	// 网页防篡改剩余数

	Surplus *uint64 `json:"Surplus,omitempty" name:"Surplus"`
	// 未安装数

	UninstallNum *uint64 `json:"UninstallNum,omitempty" name:"UninstallNum"`
	// 网页防篡改已使用数

	Used *uint64 `json:"Used,omitempty" name:"Used"`
}

type SimpleSelect struct {
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DeleteVulPackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVulPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVulPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		AgentInfo *AgentInfoRsp `json:"AgentInfo,omitempty" name:"AgentInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventsReverseShell struct {
	// 无

	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`
	// 无

	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	DetectBy *int64 `json:"DetectBy,omitempty" name:"DetectBy"`
	// 无

	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`
	// 无

	DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`
	// 无

	FindTime *string `json:"FindTime,omitempty" name:"FindTime"`
	// 无

	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`
	// 主机IP，alphabet型，v4或v6

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 无

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 无

	IsProVersion *uint64 `json:"IsProVersion,omitempty" name:"IsProVersion"`
	// 无

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 无

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 无

	ParentProcGroup *string `json:"ParentProcGroup,omitempty" name:"ParentProcGroup"`
	// 无

	ParentProcName *string `json:"ParentProcName,omitempty" name:"ParentProcName"`
	// 无

	ParentProcPath *string `json:"ParentProcPath,omitempty" name:"ParentProcPath"`
	// 无

	ParentProcUser *string `json:"ParentProcUser,omitempty" name:"ParentProcUser"`
	// 无

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 无

	Ppid *uint64 `json:"Ppid,omitempty" name:"Ppid"`
	// 无

	ProcTree *string `json:"ProcTree,omitempty" name:"ProcTree"`
	// 无

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 无

	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`
	// 无

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeAssetMachineListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Machines []*AssetMachineBaseInfo `json:"Machines,omitempty" name:"Machines"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetMachineListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetMachineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulPackageListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*VulPackage `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulPackageListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulPackageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpdateTaskGrayListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUpdateTaskGrayListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateTaskGrayListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventsBruteAttack struct {
	// 无

	Appid *int64 `json:"Appid,omitempty" name:"Appid"`
	// 无

	Area *string `json:"Area,omitempty" name:"Area"`
	// 无

	Banned *uint64 `json:"Banned,omitempty" name:"Banned"`
	// 无

	City *int64 `json:"City,omitempty" name:"City"`
	// 无

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 无

	Country *int64 `json:"Country,omitempty" name:"Country"`
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	DataFrom *int64 `json:"DataFrom,omitempty" name:"DataFrom"`
	// 无

	DstPort *int64 `json:"DstPort,omitempty" name:"DstPort"`
	// 无

	EventType *int64 `json:"EventType,omitempty" name:"EventType"`
	// 预留未用，主机唯一标识

	Guid *string `json:"Guid,omitempty" name:"Guid"`
	// 主机ip

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	MgtIp *string `json:"MgtIp,omitempty" name:"MgtIp"`
	// 无

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 无

	Province *int64 `json:"Province,omitempty" name:"Province"`
	// 无

	Region *int64 `json:"Region,omitempty" name:"Region"`
	// 源ip

	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`
	// 无

	SrcMachineName *string `json:"SrcMachineName,omitempty" name:"SrcMachineName"`
	// 无

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 无

	Username *string `json:"Username,omitempty" name:"Username"`
	// 主机唯一标识

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 无

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
}

type FindAgentUuidsByPlatformResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// uuid数组

		Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FindAgentUuidsByPlatformResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FindAgentUuidsByPlatformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTaskComponentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTaskComponentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTaskComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetDatabaseInfoRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetDatabaseInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetDatabaseInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadMalwareFileRequest struct {
	*tchttp.BaseRequest

	// 无

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

func (r *DownloadMalwareFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadMalwareFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUpdateGrayRequest struct {
	*tchttp.BaseRequest

	// 不为空则修改

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	PublishName *string `json:"PublishName,omitempty" name:"PublishName"`
	// 备注

	Reverse *string `json:"Reverse,omitempty" name:"Reverse"`
	// 无

	Uuids *string `json:"Uuids,omitempty" name:"Uuids"`
}

func (r *ModifyUpdateGrayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUpdateGrayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulType struct {
	// 无

	Descript *string `json:"Descript,omitempty" name:"Descript"`
	// 无

	DescriptEn *string `json:"DescriptEn,omitempty" name:"DescriptEn"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	NameEn *string `json:"NameEn,omitempty" name:"NameEn"`
	// 无

	OperateUser *string `json:"OperateUser,omitempty" name:"OperateUser"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeAgentPushTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*AgentPushTaskDetail `json:"List,omitempty" name:"List"`
		// 列表

		RiskCmdList []*int64 `json:"RiskCmdList,omitempty" name:"RiskCmdList"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentPushTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentPushTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCosTempCredentialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		AppID *string `json:"AppID,omitempty" name:"AppID"`
		// 无

		Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
		// Token生效的结束时间

		ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
		// 支撑组件标志0:支撑csp  1:cos

		IsCSPSupportSTS *int64 `json:"IsCSPSupportSTS,omitempty" name:"IsCSPSupportSTS"`
		// 无

		Region *string `json:"Region,omitempty" name:"Region"`
		// 临时secretId

		SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
		// 临时Key

		SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
		// 临时下载Token

		SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
		// Token生效的开始时间

		StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
		// 无

		UploadUrl *string `json:"UploadUrl,omitempty" name:"UploadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCosTempCredentialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCosTempCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulVulsListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*VulVuls `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulVulsListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulVulsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetProcessInfoListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAssetProcessInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetProcessInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBruteEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBruteEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBruteEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUpdateGrayRequest struct {
	*tchttp.BaseRequest

	// 不为空则修改

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteUpdateGrayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUpdateGrayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebLocationCountRequest struct {
	*tchttp.BaseRequest

	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeAssetWebLocationCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebLocationCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrojanSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTrojanSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTrojanSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户信息

		UserInfo *RemoteUpdateUser `json:"UserInfo,omitempty" name:"UserInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemoteUpdateUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRemoteUpdateConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRemoteUpdateConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRemoteUpdateConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSocTamperProofInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSocTamperProofInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSocTamperProofInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulAllEventsRequest struct {
	*tchttp.BaseRequest

	// Name=AppUserId、Status、Hostip、Name、Uuid

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 1:Web安全漏洞 2:系统组件漏洞 3:安全基线

	VulCategory *int64 `json:"VulCategory,omitempty" name:"VulCategory"`
}

func (r *DescribeVulAllEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulAllEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetUserDetail struct {
	// 无

	DisableTime *string `json:"DisableTime,omitempty" name:"DisableTime"`
	// 无

	Gid *string `json:"Gid,omitempty" name:"Gid"`
	// 无

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 无

	HomePath *string `json:"HomePath,omitempty" name:"HomePath"`
	// 无

	IsDomain *uint64 `json:"IsDomain,omitempty" name:"IsDomain"`
	// 无

	IsRoot *uint64 `json:"IsRoot,omitempty" name:"IsRoot"`
	// 无

	IsSshLogin *uint64 `json:"IsSshLogin,omitempty" name:"IsSshLogin"`
	// 无

	Keys []*UserKeyInfo `json:"Keys,omitempty" name:"Keys"`
	// 无

	LastLoginIp *string `json:"LastLoginIp,omitempty" name:"LastLoginIp"`
	// 无

	LastLoginLoc *string `json:"LastLoginLoc,omitempty" name:"LastLoginLoc"`
	// 无

	LastLoginTerminal *string `json:"LastLoginTerminal,omitempty" name:"LastLoginTerminal"`
	// 无

	LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	PasswordChangeTime *string `json:"PasswordChangeTime,omitempty" name:"PasswordChangeTime"`
	// 无

	PasswordChangeType *uint64 `json:"PasswordChangeType,omitempty" name:"PasswordChangeType"`
	// 无

	PasswordDueTime *string `json:"PasswordDueTime,omitempty" name:"PasswordDueTime"`
	// 无

	PasswordLockDays *int64 `json:"PasswordLockDays,omitempty" name:"PasswordLockDays"`
	// 无

	PasswordWarnDays *uint64 `json:"PasswordWarnDays,omitempty" name:"PasswordWarnDays"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 无

	Shell *string `json:"Shell,omitempty" name:"Shell"`
	// 无

	ShellLoginStatus *uint64 `json:"ShellLoginStatus,omitempty" name:"ShellLoginStatus"`
	// 无

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 无

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	UserType *uint64 `json:"UserType,omitempty" name:"UserType"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type CancelPublishInstallPackageRequest struct {
	*tchttp.BaseRequest

	// 无

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *CancelPublishInstallPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelPublishInstallPackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityDynamicsRequest struct {
	*tchttp.BaseRequest

	// 无

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSecurityDynamicsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityDynamicsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRemoteUpdateTagRequest struct {
	*tchttp.BaseRequest

	// 标签信息

	TagInfo *RemoteUpdateTag `json:"TagInfo,omitempty" name:"TagInfo"`
}

func (r *ModifyRemoteUpdateTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRemoteUpdateTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulPackageListRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeVulPackageListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulPackageListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRemoteUpdateUsersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRemoteUpdateUsersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRemoteUpdateUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZKConfigListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*ZkGroup `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZKConfigListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZKConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportHostSafetyStatusRequest struct {
	*tchttp.BaseRequest

	// StartTime

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 无

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ExportHostSafetyStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportHostSafetyStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetUserCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Users []*AssetKeyVal `json:"Users,omitempty" name:"Users"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetUserCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetUserCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDefenseVulInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CVSS信息

		CVSS *string `json:"CVSS,omitempty" name:"CVSS"`
		// 漏洞CVEID

		CveId *string `json:"CveId,omitempty" name:"CveId"`
		// cvss详情

		CveInfo *string `json:"CveInfo,omitempty" name:"CveInfo"`
		// Cvss分数

		CvssScore *uint64 `json:"CvssScore,omitempty" name:"CvssScore"`
		// cvss 分数 浮点型

		CvssScoreFloat *float64 `json:"CvssScoreFloat,omitempty" name:"CvssScoreFloat"`
		// 已防御的攻击次数

		DefenseAttackCount *uint64 `json:"DefenseAttackCount,omitempty" name:"DefenseAttackCount"`
		// 漏洞描述信息

		Description *string `json:"Description,omitempty" name:"Description"`
		// 修复是否支持：0-windows/linux均不支持修复 ;1-windows/linux 均支持修复 ;2-仅linux支持修复;3-仅windows支持修复

		FixSwitch *int64 `json:"FixSwitch,omitempty" name:"FixSwitch"`
		// 漏洞标签 多个逗号分割

		Labels *string `json:"Labels,omitempty" name:"Labels"`
		// 发布时间

		PublicDate *string `json:"PublicDate,omitempty" name:"PublicDate"`
		// 参考链接

		Reference *string `json:"Reference,omitempty" name:"Reference"`
		// 修复方案

		RepairPlan *string `json:"RepairPlan,omitempty" name:"RepairPlan"`
		// 全网修复成功次数, 不支持自动修复的漏洞默认返回0

		SuccessFixCount *uint64 `json:"SuccessFixCount,omitempty" name:"SuccessFixCount"`
		// 漏洞id

		VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
		// 危害等级：1-低危；2-中危；3-高危；4-严重

		VulLevel *uint64 `json:"VulLevel,omitempty" name:"VulLevel"`
		// 漏洞名称

		VulName *string `json:"VulName,omitempty" name:"VulName"`
		// 漏洞分类 1: web-cms漏洞 2:应用漏洞  4: Linux软件漏洞 5: Windows系统漏洞

		VulType *uint64 `json:"VulType,omitempty" name:"VulType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDefenseVulInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDefenseVulInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadMalwareFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Content

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadMalwareFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadMalwareFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetJarListRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetJarListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetJarListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishKnowledgeConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *PublishKnowledgeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishKnowledgeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CvmDetailInfo struct {
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最后一次同步时间

	LatestSyncTime *string `json:"LatestSyncTime,omitempty" name:"LatestSyncTime"`
	// 公网ip地址

	PublicIpAddresses *string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`
	// vpc id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type ModifyRemoteUpdatePublishLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRemoteUpdatePublishLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRemoteUpdatePublishLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstallPackageRequest struct {
	*tchttp.BaseRequest

	// 无

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteInstallPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstallPackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRemoteUpdateConfigsRequest struct {
	*tchttp.BaseRequest

	// 要删除的更新配置的ID列表

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteRemoteUpdateConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRemoteUpdateConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpdateTaskResultRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 主键

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUpdateTaskResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateTaskResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetCoreModuleListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters *FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *ExportAssetCoreModuleListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetCoreModuleListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulDefenceMachine struct {
	// AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// CreateTime

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 漏洞防御开关（全部、开启、关闭）Enable 0关闭 1开启

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// InstanceId

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// InstanceName

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// ModifyTime

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// PrivateIp

	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
	// PublicIp

	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
	// Quuid

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 防御状态（全部、全部正常、存在异常、暂未开启）Status 插件是否存在异常 0: 正常 1: 异常 2:暂未开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ExportJavaMemShellPluginsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportJavaMemShellPluginsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportJavaMemShellPluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetJarDetail struct {
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	Path *string `json:"Path,omitempty" name:"Path"`
	// 无

	Process []*AssetProcessListInfo `json:"Process,omitempty" name:"Process"`
	// 无

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 无

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeRemoteUpdateUserNameUniqueRequest struct {
	*tchttp.BaseRequest

	// 用户名称

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *DescribeRemoteUpdateUserNameUniqueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateUserNameUniqueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*VulDefenceEvent `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulDefenceEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRemoteUpdateUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新创建的用户的Id

		UserId *uint64 `json:"UserId,omitempty" name:"UserId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRemoteUpdateUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRemoteUpdateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseAgentListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*UpgradeAgents `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseAgentListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseAgentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeksRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeWeeksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWeeksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTaskComponentsInfoRequest struct {
	*tchttp.BaseRequest

	// 作者

	Author *string `json:"Author,omitempty" name:"Author"`
	// 校验规则

	CheckRule *string `json:"CheckRule,omitempty" name:"CheckRule"`
	// 组件描述

	Describe *string `json:"Describe,omitempty" name:"Describe"`
	// 组件英文描述

	DescribeEn *string `json:"DescribeEn,omitempty" name:"DescribeEn"`
	// 组件官网

	Homepage *string `json:"Homepage,omitempty" name:"Homepage"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 组件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件英文名称

	NameEn *string `json:"NameEn,omitempty" name:"NameEn"`
	// 检测脚本

	PocFile *string `json:"PocFile,omitempty" name:"PocFile"`
	// 检测开关 1-开启 0-关闭

	Switch *int64 `json:"Switch,omitempty" name:"Switch"`
	// 组件类型 1-系统组件 2-Web组件

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 验证规则

	VerifyRule *string `json:"VerifyRule,omitempty" name:"VerifyRule"`
}

func (r *ModifyTaskComponentsInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTaskComponentsInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishAgentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishAgentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddPushTaskRequest struct {
	*tchttp.BaseRequest

	// 命令参数集合

	Args []*string `json:"Args,omitempty" name:"Args"`
	// 无

	OptionType *int64 `json:"OptionType,omitempty" name:"OptionType"`
	// 无

	Uuids *string `json:"Uuids,omitempty" name:"Uuids"`
}

func (r *AddPushTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddPushTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpdateTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// UpdateTask

		List []*UpdateTaskDetailRsp `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpdateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulScanConfigListRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Publish：0 1

	Publish *int64 `json:"Publish,omitempty" name:"Publish"`
}

func (r *DescribeVulScanConfigListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulScanConfigListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportTaskComponentsRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportTaskComponentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportTaskComponentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulEventsRequest struct {
	*tchttp.BaseRequest

	// Name=AppUserId、Status、Hostip、Name、Uuid

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 1:Web安全漏洞 2:系统组件漏洞 3:安全基线

	VulCategory *int64 `json:"VulCategory,omitempty" name:"VulCategory"`
}

func (r *DescribeVulEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFlagshipLicenseRequest struct {
	*tchttp.BaseRequest

	// 文件内容

	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`
}

func (r *ModifyFlagshipLicenseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFlagshipLicenseRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoteUpdateConfigExportFile struct {
	// 配置Id

	ConfigId *uint64 `json:"ConfigId,omitempty" name:"ConfigId"`
	// 下载链接。有效时间5分钟

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
	// 文件状态。ok：表示可用，generating：表示正在生成中

	FileStatus *string `json:"FileStatus,omitempty" name:"FileStatus"`
}

type AgentsPushTaskDetail struct {
	// 无

	Arg *string `json:"Arg,omitempty" name:"Arg"`
	// 无

	Caller *string `json:"Caller,omitempty" name:"Caller"`
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 无

	CustomArg *string `json:"CustomArg,omitempty" name:"CustomArg"`
	// 无

	ExecSucCnt *int64 `json:"ExecSucCnt,omitempty" name:"ExecSucCnt"`
	// 无

	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
	// 主键

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	MaxRetry *int64 `json:"MaxRetry,omitempty" name:"MaxRetry"`
	// 无

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 无

	OptionType *int64 `json:"OptionType,omitempty" name:"OptionType"`
	// 无

	PushSucCnt *int64 `json:"PushSucCnt,omitempty" name:"PushSucCnt"`
	// 无

	RetryInterval *int64 `json:"RetryInterval,omitempty" name:"RetryInterval"`
	// 无

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 无

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	UuidCnt *int64 `json:"UuidCnt,omitempty" name:"UuidCnt"`
}

type DescribeAssetProcessInfoListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAssetProcessInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetProcessInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBruteForceEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*EventsBruteAttack `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBruteForceEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBruteForceEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileSpecimenListRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeFileSpecimenListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileSpecimenListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpdateTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeUpdateTaskDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateTaskDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckFileStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载url

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckFileStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckFileStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebAppListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetWebAppListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebAppListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetMachineBaseInfo struct {
	// 无

	Cpu *string `json:"Cpu,omitempty" name:"Cpu"`
	// 无

	CpuLoad *string `json:"CpuLoad,omitempty" name:"CpuLoad"`
	// 无

	CpuSize *uint64 `json:"CpuSize,omitempty" name:"CpuSize"`
	// 无

	DiskLoad *string `json:"DiskLoad,omitempty" name:"DiskLoad"`
	// 无

	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 无

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 无

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	MemLoad *string `json:"MemLoad,omitempty" name:"MemLoad"`
	// 无

	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	PartitionCount *uint64 `json:"PartitionCount,omitempty" name:"PartitionCount"`
	// 无

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ModifyStatisticResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyStatisticResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseFunctionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLicenseFunctionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseFunctionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesSimpleRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版 | Flagship : 旗舰版 | ProtectedMachines: 专业版+旗舰版 | UnFlagship : 非旗舰版 | PRO_POST_PAY：专业版按量计费 | PRO_PRE_PAY：专业版包年包月）</li>
	// <li>TagId - String - 是否必填：否 - 标签ID </li>

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeMachinesSimpleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachinesSimpleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulTopRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 获取top几

	Top *uint64 `json:"Top,omitempty" name:"Top"`
}

func (r *DescribeVulTopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulTopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportSecurityDynamicsRequest struct {
	*tchttp.BaseRequest

	// 无

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ExportSecurityDynamicsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportSecurityDynamicsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBruteForceEventsRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBruteForceEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBruteForceEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportSpecimenFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportSpecimenFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportSpecimenFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetWebServiceBaseInfo struct {
	// 无

	BinPath *string `json:"BinPath,omitempty" name:"BinPath"`
	// 无

	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`
	// 无

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 无

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	InstallPath *string `json:"InstallPath,omitempty" name:"InstallPath"`
	// 无

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	ProcessCount *uint64 `json:"ProcessCount,omitempty" name:"ProcessCount"`
	// 无

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	User *string `json:"User,omitempty" name:"User"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeBaselineRuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		List []*BaselineRule `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaselineRuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaselineRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineFileTamperRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询主机相关核心文件监控规则详情

		List []*MachineFileTamperRule `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineFileTamperRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineFileTamperRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTaskStatus struct {
	// 无

	Descript *string `json:"Descript,omitempty" name:"Descript"`
	// Id     int

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 无

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeDNSEventDetailRequest struct {
	*tchttp.BaseRequest

	// 无

	AppUserId *uint64 `json:"AppUserId,omitempty" name:"AppUserId"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeDNSEventDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDNSEventDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteUpdateDownloadTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 允许下载的文件路径

		AllowPath *string `json:"AllowPath,omitempty" name:"AllowPath"`
		// 应用Id

		AppID *string `json:"AppID,omitempty" name:"AppID"`
		// COS桶

		Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
		// Token生效的结束时间

		ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
		// 支撑组件标志0:支撑csp  1:cos

		IsCSPSupportSTS *int64 `json:"IsCSPSupportSTS,omitempty" name:"IsCSPSupportSTS"`
		// COS地域

		Region *string `json:"Region,omitempty" name:"Region"`
		// 临时secretId

		SecretId *string `json:"SecretId,omitempty" name:"SecretId"`
		// 临时Key

		SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
		// 临时下载Token

		SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`
		// Token生效的开始时间

		StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemoteUpdateDownloadTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateDownloadTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityDynamic struct {
	// 无

	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`
	// 无

	EventType *string `json:"EventType,omitempty" name:"EventType"`
	// 无

	Message *string `json:"Message,omitempty" name:"Message"`
	// 无

	SecurityLevel *string `json:"SecurityLevel,omitempty" name:"SecurityLevel"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeFileTamperEventRuleInfoRequest struct {
	*tchttp.BaseRequest

	// 事件ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 租户APPID

	UserAppId *int64 `json:"UserAppId,omitempty" name:"UserAppId"`
}

func (r *DescribeFileTamperEventRuleInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFileTamperEventRuleInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginEventsRequest struct {
	*tchttp.BaseRequest

	// FilterParams

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeLoginEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLoginEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulVuls struct {
	// 告警开关: 0-windows/linux均关闭, 1-windows开启, 2-linux开启, 3-windows/linux都开启

	Alarm *int64 `json:"Alarm,omitempty" name:"Alarm"`
	// 作者

	Author *string `json:"Author,omitempty" name:"Author"`
	// 检测规则

	CheckRule *string `json:"CheckRule,omitempty" name:"CheckRule"`
	// 组件id

	ComponentId *int64 `json:"ComponentId,omitempty" name:"ComponentId"`
	// 无

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// cve id

	CveId *string `json:"CveId,omitempty" name:"CveId"`
	// cvss 详情

	Cvss *string `json:"Cvss,omitempty" name:"Cvss"`
	// cvss 分数

	CvssScore *float64 `json:"CvssScore,omitempty" name:"CvssScore"`
	// 无

	DescTemplateId *int64 `json:"DescTemplateId,omitempty" name:"DescTemplateId"`
	// 描述

	Descript *string `json:"Descript,omitempty" name:"Descript"`
	// 无

	DescriptEn *string `json:"DescriptEn,omitempty" name:"DescriptEn"`
	// 是否是应急漏洞 0 否 1 是

	Emergency *int64 `json:"Emergency,omitempty" name:"Emergency"`
	// 修复建议

	Fix *string `json:"Fix,omitempty" name:"Fix"`
	// 无

	FixEn *string `json:"FixEn,omitempty" name:"FixEn"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Label *string `json:"Label,omitempty" name:"Label"`
	// 无

	LabelEn *string `json:"LabelEn,omitempty" name:"LabelEn"`
	// 等级

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 漏洞名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	NameEn *string `json:"NameEn,omitempty" name:"NameEn"`
	// 无

	Num *int64 `json:"Num,omitempty" name:"Num"`
	// 操作人

	OperateUser *string `json:"OperateUser,omitempty" name:"OperateUser"`
	// poc文件

	PocFile *string `json:"PocFile,omitempty" name:"PocFile"`
	// 发布时间

	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
	// 参考地址

	Reference *string `json:"Reference,omitempty" name:"Reference"`
	// 开关 1 开 0 关

	Switch *int64 `json:"Switch,omitempty" name:"Switch"`
	// 无

	Test *int64 `json:"Test,omitempty" name:"Test"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 校验规则

	VerifyRule *string `json:"VerifyRule,omitempty" name:"VerifyRule"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 漏洞分类id 1 Web应用漏洞 2 系统组件漏洞 3 安全基线

	VulCategory *int64 `json:"VulCategory,omitempty" name:"VulCategory"`
	// 无

	VulComponentName *string `json:"VulComponentName,omitempty" name:"VulComponentName"`
	// 漏洞类型ID

	VulTypeId *int64 `json:"VulTypeId,omitempty" name:"VulTypeId"`
	// 无

	VulTypeName *string `json:"VulTypeName,omitempty" name:"VulTypeName"`
}

type DescribeRemoteUpdateConfigMainInfoRequest struct {
	*tchttp.BaseRequest

	// 配置（所属组件的）类型

	ComponentType *uint64 `json:"ComponentType,omitempty" name:"ComponentType"`
	// 配置名称

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
	// 配置版本

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
}

func (r *DescribeRemoteUpdateConfigMainInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdateConfigMainInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportReverseShellEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportReverseShellEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportReverseShellEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyZKConfigInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// zk配置信息

		ZKConfigDetail *ZKConfigInfo `json:"ZKConfigDetail,omitempty" name:"ZKConfigDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyZKConfigInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyZKConfigInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZKConfigInfo struct {
	// 无

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Path *string `json:"Path,omitempty" name:"Path"`
	// 无

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 无

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeAssetWebAppListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 无

		WebApps []*AssetWebAppBaseInfo `json:"WebApps,omitempty" name:"WebApps"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebAppListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebAppListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLicenseHistoryReleasedRequest struct {
	*tchttp.BaseRequest

	// Id

	DoMain *string `json:"DoMain,omitempty" name:"DoMain"`
	// 限制条数,默认10,最大100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeLicenseHistoryReleasedRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseHistoryReleasedRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpecimenFileRequest struct {
	*tchttp.BaseRequest

	// 无

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

func (r *DescribeSpecimenFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpecimenFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportCustomerLicenseListRequest struct {
	*tchttp.BaseRequest

	// By

	By *string `json:"By,omitempty" name:"By"`
	// FilterParams

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *ExportCustomerLicenseListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportCustomerLicenseListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportDefenseVulInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportDefenseVulInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDefenseVulInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentPushTaskResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*AgentsPushResult `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentPushTaskResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentPushTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetProcessListInfo struct {
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Path *string `json:"Path,omitempty" name:"Path"`
	// 无

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 无

	Status *string `json:"Status,omitempty" name:"Status"`
	// 无

	User *string `json:"User,omitempty" name:"User"`
	// 无

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeAssetPortInfoListRequest struct {
	*tchttp.BaseRequest

	// 无

	By *string `json:"By,omitempty" name:"By"`
	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Order *string `json:"Order,omitempty" name:"Order"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAssetPortInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetPortInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJavaMemShellPluginInfoRequest struct {
	*tchttp.BaseRequest

	// By

	By *string `json:"By,omitempty" name:"By"`
	// 过滤条件：Pid精确匹配，MainClass模糊匹配

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeJavaMemShellPluginInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJavaMemShellPluginInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSocTamperProofDetailRequest struct {
	*tchttp.BaseRequest

	// 排序字段方式 EventTamperTime

	By *string `json:"By,omitempty" name:"By"`
	// 条件过滤, ExactMatch参数无效

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 返回数量，默认为10，最大值为100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式 asc,desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 网页防篡改ID

	SocTamperProofId *uint64 `json:"SocTamperProofId,omitempty" name:"SocTamperProofId"`
}

func (r *DescribeSocTamperProofDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSocTamperProofDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteJavaMemShellsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteJavaMemShellsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteJavaMemShellsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetEnvListRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 类型：0

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 服务器Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeAssetEnvListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetEnvListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineFileTamperRulesRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 租户APPID

	UserAppId *int64 `json:"UserAppId,omitempty" name:"UserAppId"`
	// 主机uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeMachineFileTamperRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineFileTamperRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetInitServiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Services []*AssetInitServiceBaseInfo `json:"Services,omitempty" name:"Services"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetInitServiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetInitServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwareEventsRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 是否是全量

	IsAll *bool `json:"IsAll,omitempty" name:"IsAll"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeMalwareEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMalwareEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMalwareEvilRequest struct {
	*tchttp.BaseRequest

	// Id列表

	Ids *int64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *ModifyMalwareEvilRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMalwareEvilRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentStatus struct {
	// 客户端版本

	ClientVersion *string `json:"ClientVersion,omitempty" name:"ClientVersion"`
	// flo w

	Flow *string `json:"Flow,omitempty" name:"Flow"`
	// 网关ip

	GatewayIp *string `json:"GatewayIp,omitempty" name:"GatewayIp"`
	// 网关端口

	GatewayPort *string `json:"GatewayPort,omitempty" name:"GatewayPort"`
	// 机器唯一标识

	Guid *string `json:"Guid,omitempty" name:"Guid"`
	// 本地ip

	LocalIp *string `json:"LocalIp,omitempty" name:"LocalIp"`
	// 本地端口

	LocalPort *string `json:"LocalPort,omitempty" name:"LocalPort"`
	// 时间戳

	LoginTimestamp *string `json:"LoginTimestamp,omitempty" name:"LoginTimestamp"`
	// 机器名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// proxy ip

	ProxyIp *string `json:"ProxyIp,omitempty" name:"ProxyIp"`
	// 代理端口

	ProxyPort *string `json:"ProxyPort,omitempty" name:"ProxyPort"`
	// 推送标志

	PushFlag *string `json:"PushFlag,omitempty" name:"PushFlag"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 时间戳

	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
}

type DescribeRemoteUpdatePublishLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查到的该用户的针对某一组件的发布日志

		Logs []*RemoteUpdatePublishLog `json:"Logs,omitempty" name:"Logs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemoteUpdatePublishLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteUpdatePublishLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetPortBaseInfo struct {
	// 无

	BindIp *string `json:"BindIp,omitempty" name:"BindIp"`
	// 无

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 无

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 无

	IsNew *int64 `json:"IsNew,omitempty" name:"IsNew"`
	// 无

	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`
	// 无

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 无

	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`
	// 无

	Md5 *string `json:"Md5,omitempty" name:"Md5"`
	// 无

	OsInfo *string `json:"OsInfo,omitempty" name:"OsInfo"`
	// 无

	Param *string `json:"Param,omitempty" name:"Param"`
	// 无

	ParentProcessName *string `json:"ParentProcessName,omitempty" name:"ParentProcessName"`
	// 无

	Pid *string `json:"Pid,omitempty" name:"Pid"`
	// 无

	Port *string `json:"Port,omitempty" name:"Port"`
	// 无

	Ppid *string `json:"Ppid,omitempty" name:"Ppid"`
	// 无

	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`
	// 无

	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`
	// 无

	ProcessVersion *string `json:"ProcessVersion,omitempty" name:"ProcessVersion"`
	// 无

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 无

	Proto *string `json:"Proto,omitempty" name:"Proto"`
	// 无

	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
	// 无

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 无

	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`
	// 无

	Teletype *string `json:"Teletype,omitempty" name:"Teletype"`
	// 无

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	User *string `json:"User,omitempty" name:"User"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribeVulScanConfigFullListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*VulScanConfig `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulScanConfigFullListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulScanConfigFullListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBashEventsRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportBashEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportBashEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMalwareResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMalwareResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMalwareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskComponentsDetailRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeTaskComponentsDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskComponentsDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SeparateMalwareRequest struct {
	*tchttp.BaseRequest

	// isolate recovery

	CMD *string `json:"CMD,omitempty" name:"CMD"`
	// 无

	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *SeparateMalwareRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SeparateMalwareRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUpdateTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUpdateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUpdateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTaskResult struct {
	// 无

	Descript *string `json:"Descript,omitempty" name:"Descript"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 无

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 无

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
	// 无

	UpdateTime []*string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DescribePrivilegeEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List *PrivilegeEscalationProcess `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivilegeEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePrivilegeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpdateTaskListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// UpdateTask

		List []*UpdateTasksLinkGray `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpdateTaskListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulDefenceMachineListRequest struct {
	*tchttp.BaseRequest

	// By

	By *string `json:"By,omitempty" name:"By"`
	// Name="AppId", "IP", "InstanceName", "InstanceId", "Enable", "Status", "CreateTime", "ModifyTime"

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Order

	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeVulDefenceMachineListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulDefenceMachineListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishAgentRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *PublishAgentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishAgentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RspSocAuthDetailList struct {
	// 授权服务器IP

	AuthServerIp *string `json:"AuthServerIp,omitempty" name:"AuthServerIp"`
	// 授权服务器名

	AuthServerName *string `json:"AuthServerName,omitempty" name:"AuthServerName"`
	// 到期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 开通时间

	OpenTime *string `json:"OpenTime,omitempty" name:"OpenTime"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type DescribeJavaMemShellPluginListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：Keywords: ip或者主机名模糊查询, JavaShellStatus，Exception精确匹配

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
	// 需要返回的数量，默认为10，最大值为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeJavaMemShellPluginListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJavaMemShellPluginListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBruteEventsRequest struct {
	*tchttp.BaseRequest

	// Cmd：del ignore mis

	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`
	// 无

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
	// 无

	Week *int64 `json:"Week,omitempty" name:"Week"`
}

func (r *ModifyBruteEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBruteEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetAppListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Apps []*AssetAppBaseInfo `json:"Apps,omitempty" name:"Apps"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetAppListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetAppListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRemoteUpdatePublishLogRequest struct {
	*tchttp.BaseRequest

	// 发布日志Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 说明

	Instruction *string `json:"Instruction,omitempty" name:"Instruction"`
}

func (r *ModifyRemoteUpdatePublishLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRemoteUpdatePublishLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DNSIdInfo struct {
	// 无

	AppUserId *int64 `json:"AppUserId,omitempty" name:"AppUserId"`
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type DescribePushCommandListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*PushCommandDesc `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePushCommandListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePushCommandListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDefenceMachineListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVulDefenceMachineListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportVulDefenceMachineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CredentialResult struct {
	// 无

	Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`
	// 无

	Error *CredentialError `json:"Error,omitempty" name:"Error"`
	// 无

	Expiration *string `json:"Expiration,omitempty" name:"Expiration"`
	// 无

	ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// 无

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
}

type UpdateConfig struct {
	// 无

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 无

	IsPublish *int64 `json:"IsPublish,omitempty" name:"IsPublish"`
	// Md5

	MdSum *string `json:"MdSum,omitempty" name:"MdSum"`
	// 配置名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 升级类型

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 无

	UrlPath *string `json:"UrlPath,omitempty" name:"UrlPath"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeInstallPackagePublishListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 发布的安装包列表

		List []*UpdateInstallPackage `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstallPackagePublishListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstallPackagePublishListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LicenseOrder struct {
	// 授权ID

	LicenseId *uint64 `json:"LicenseId,omitempty" name:"LicenseId"`
	// 授权类型

	LicenseType *uint64 `json:"LicenseType,omitempty" name:"LicenseType"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 订单类型

	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`
	// 授权订单资源状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type UserKeyInfo struct {
	// 无

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 无

	EncryptType *string `json:"EncryptType,omitempty" name:"EncryptType"`
	// 无

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeleteUpdateConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUpdateConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUpdateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetWebServiceInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 无

		WebServices []*AssetWebServiceBaseInfo `json:"WebServices,omitempty" name:"WebServices"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebServiceInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebServiceInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDNSKnowledgeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*DnsKnowledge `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDNSKnowledgeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDNSKnowledgeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportHostRiskTrendRequest struct {
	*tchttp.BaseRequest

	// StartTime

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 无

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ExportHostRiskTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportHostRiskTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*EventsVul `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVulEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRemoteUpdateTagRequest struct {
	*tchttp.BaseRequest

	// 要删除的标签的ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteRemoteUpdateTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRemoteUpdateTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSocTamperProofListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*RspSocTamperProofList `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSocTamperProofListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSocTamperProofListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BruteAttackTopInfo struct {
	// 爆破总数

	BruteAttackCount *int64 `json:"BruteAttackCount,omitempty" name:"BruteAttackCount"`
	// 爆破失败数

	FailCount *int64 `json:"FailCount,omitempty" name:"FailCount"`
	// 无

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 爆破成功数

	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
}

type DescribeUpdateTaskLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Limit

		List []*UpdateTaskLog `json:"List,omitempty" name:"List"`
		// Offset

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpdateTaskLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpdateTaskLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetCoreModuleListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetCoreModuleListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportAssetCoreModuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaselineItem struct {
	// 是否可以修复

	CanBeFixed *int64 `json:"CanBeFixed,omitempty" name:"CanBeFixed"`
	// 检测项分类

	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`
	// 检测结果描述

	DetectResultDesc *string `json:"DetectResultDesc,omitempty" name:"DetectResultDesc"`
	// 检测状态：0 未通过，1：忽略，3：通过，5：检测中

	DetectStatus *int64 `json:"DetectStatus,omitempty" name:"DetectStatus"`
	// 第一次出现时间

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 修复方法

	FixMethod *string `json:"FixMethod,omitempty" name:"FixMethod"`
	// 主机ID

	HostId *string `json:"HostId,omitempty" name:"HostId"`
	// 主机IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 主机名

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 项描述

	ItemDesc *string `json:"ItemDesc,omitempty" name:"ItemDesc"`
	// 项Id

	ItemId *int64 `json:"ItemId,omitempty" name:"ItemId"`
	// 项名称

	ItemName *string `json:"ItemName,omitempty" name:"ItemName"`
	// 最近出现时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 危险等级

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 所属规则

	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
	// 主机安全uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 外网IP

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
}

type ExportFileSpecimenListRequest struct {
	*tchttp.BaseRequest

	// 无

	Filters []*FilterParam `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportFileSpecimenListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportFileSpecimenListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileTamperRuleDetail struct {
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 规则id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 是否全局规则(默认否) 0：否 ，1：是

	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
	// 风险等级 0：无， 1: 高危， 2:中危， 3: 低危

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 更新时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规则

	Rule []*FileTamperRule `json:"Rule,omitempty" name:"Rule"`
	// 状态 0: 启用 1: 已关闭

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 生效主机的总数

	UuidTotalCount *uint64 `json:"UuidTotalCount,omitempty" name:"UuidTotalCount"`
	// 生效主机uuid,空表示全部主机，通过参数可控制返回的条数

	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`
}

type DescribeAssetWebLocationListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Locations []*AssetWebLocationBaseInfo `json:"Locations,omitempty" name:"Locations"`
		// 无

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetWebLocationListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAssetWebLocationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSocAuthListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表

		List []*RspSocAuthDetailList `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSocAuthListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSocAuthListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRemoteUpdateUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRemoteUpdateUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRemoteUpdateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRemoteUpdatePublishItemRequest struct {
	*tchttp.BaseRequest

	// 配置Id

	ConfigId *uint64 `json:"ConfigId,omitempty" name:"ConfigId"`
	// 操作人名称

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 用户Id

	UserId *uint64 `json:"UserId,omitempty" name:"UserId"`
}

func (r *CreateRemoteUpdatePublishItemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRemoteUpdatePublishItemRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBruteAttacksRequest struct {
	*tchttp.BaseRequest

	// 无

	Ids []*int64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteBruteAttacksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBruteAttacksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
