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

package v20180409

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DescribeCatLogsRequest struct {
	*tchttp.BaseRequest

	// 从第Offset 条开始查询。缺省值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 本批次查询Limit 条记录。缺省值为20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 格式如：2017-05-09 00:00:00  缺省为1小时前

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 格式如：2017-05-10 00:00:00  缺省为当前时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 拨测任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 按时间升序或降序。默认降序。可选值： Desc, Asc

	SortType *string `json:"SortType,omitempty" name:"SortType"`
}

func (r *DescribeCatLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCatLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserLimitResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户可建立的最大任务数

		MaxTaskNum *uint64 `json:"MaxTaskNum,omitempty" name:"MaxTaskNum"`
		// 用户可用的最大拨测结点数

		MaxAgentNum *uint64 `json:"MaxAgentNum,omitempty" name:"MaxAgentNum"`
		// 用户可建立的最大拨测分组数

		MaxGroupNum *uint64 `json:"MaxGroupNum,omitempty" name:"MaxGroupNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPloicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPloicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmPloicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest

	// 拨测分组id，体现本拨测任务要采用那些运营商作为拨测源。一般可直接填写本用户的默认拨测分组。参见：DescribeAgentGroupList 接口，本参数使用返回结果里的groupId的值。

	AgentGroupId *uint64 `json:"AgentGroupId,omitempty" name:"AgentGroupId"`
	// http, https, ping, tcp 之一

	CatTypeName *string `json:"CatTypeName,omitempty" name:"CatTypeName"`
	// 拨测的url  例如：www.baidu.com (url域名解析需要能解析出具体的ip)

	Url *string `json:"Url,omitempty" name:"Url"`
	// 需要满足ip 的格式

	Host *string `json:"Host,omitempty" name:"Host"`
	// 服务端监听或接收数据的端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 是否为Header请求（非0 发起Header 请求。为0，且PostData 非空，发起POST请求。为0，PostData 为空，发起GET请求）

	IsHeader *uint64 `json:"IsHeader,omitempty" name:"IsHeader"`
	// url中含有https时有用。缺省为SSLv23。需要为 TLSv1_2, TLSv1_1, TLSv1, SSLv2, SSLv23, SSLv3 之一

	SslVer *string `json:"SslVer,omitempty" name:"SslVer"`
	// POST 请求数据。空字符串表示非POST请求

	PostData *string `json:"PostData,omitempty" name:"PostData"`
	// 用户agent 信息

	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`
	// 要在结果中进行匹配的字符串

	CheckStr *string `json:"CheckStr,omitempty" name:"CheckStr"`
	// 1 表示通过检查结果是否包含CheckStr 进行校验

	CheckType *uint64 `json:"CheckType,omitempty" name:"CheckType"`
	// 需要设置的cookie信息

	Cookie *string `json:"Cookie,omitempty" name:"Cookie"`
	// 拨测周期。取值可为1,5,15,30之一, 单位：分钟。精度不能低于用户等级规定的最小精度

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 拨测任务名称不能超过32个字符。同一个用户创建的任务名不可重复

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务号。用于验证且修改任务时传入原任务号

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *CreateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmTopicRequest struct {
	*tchttp.BaseRequest

	// 如果不存在拨测相关的主题，是否自动创建一个

	NeedAdd *uint64 `json:"NeedAdd,omitempty" name:"NeedAdd"`
}

func (r *DescribeAlarmTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmGroupInfo struct {
	// 告警接受组Id

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 告警接受组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 告警接受通道数(一个邮件或短信接收人为一个通道)

	Channel *uint64 `json:"Channel,omitempty" name:"Channel"`
	// 备注信息

	Remarks *string `json:"Remarks,omitempty" name:"Remarks"`
	// 接受组创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type CreateAgentGroupRequest struct {
	*tchttp.BaseRequest

	// 拨测分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 是否为默认分组

	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`
	// Province, Isp 需要成对地进行选择。参数对的取值范围。参见：DescribeAgentList 的返回结果。

	Agents []*CatAgent `json:"Agents,omitempty" name:"Agents"`
}

func (r *CreateAgentGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAgentGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTaskByWanIpRequest struct {
	*tchttp.BaseRequest

	// 从第几条开始拉取

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 一批拉取多少条

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetTaskByWanIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTaskByWanIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmsByTaskRequest struct {
	*tchttp.BaseRequest

	// 从第Offset 条开始查询。缺省值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 本批次查询Limit 条记录。缺省值为20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 0 全部, 1 已恢复, 2 未恢复  默认为0。其他值，视为0 查全部状态。

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 格式如：2017-05-09 00:00:00  缺省为7天前0点

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 格式如：2017-05-10 00:00:00  缺省为明天0点

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 拨测任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 排序字段，可为Time, ObjName, Duration, Status, Content 之一。缺省为Time。

	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`
	// 升序或降序。可为Desc, Asc之一。缺省为Desc。

	SortType *string `json:"SortType,omitempty" name:"SortType"`
}

func (r *DescribeAlarmsByTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmsByTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 本批告警信息列表

		AlarmInfos []*AlarmInfo `json:"AlarmInfos,omitempty" name:"AlarmInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户名下总的告警接收组数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 满足条件的告警接收组列表

		AlarmGroupInfos []*AlarmGroupInfo `json:"AlarmGroupInfos,omitempty" name:"AlarmGroupInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResultSummary struct {
	// 统计时间

	LogTime *string `json:"LogTime,omitempty" name:"LogTime"`
	// 任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 实时可用率

	AvailRatio *float64 `json:"AvailRatio,omitempty" name:"AvailRatio"`
	// 实时响应时间

	ReponseTime *float64 `json:"ReponseTime,omitempty" name:"ReponseTime"`
}

type BindAlarmPloicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindAlarmPloicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindAlarmPloicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAgentGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAgentGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAgentGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentGroupRequest struct {
	*tchttp.BaseRequest

	// 拨测分组id

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeAgentGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CatAgent struct {
	// 拨测结点所在的省份（拼音缩写）

	Province *string `json:"Province,omitempty" name:"Province"`
	// 拨测结点所在的运营商（英文缩写）

	Isp *string `json:"Isp,omitempty" name:"Isp"`
	// 拨测结点所在的省份（中文名称）

	ProvinceName *string `json:"ProvinceName,omitempty" name:"ProvinceName"`
	// 拨测结点所在的运营商（中文名称）

	IspName *string `json:"IspName,omitempty" name:"IspName"`
}

type CatTaskInfo struct {
	// 拨测任务类型

	CatTypeId *uint64 `json:"CatTypeId,omitempty" name:"CatTypeId"`
	// 拨测任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 拨测任务链接

	CgiUrl *string `json:"CgiUrl,omitempty" name:"CgiUrl"`
	// 拨测任务参数

	Parameters *string `json:"Parameters,omitempty" name:"Parameters"`
	// 拨测任务标志

	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
}

type DescribeTasksByTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的总任务数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 任务列表

		Tasks []*TaskAlarm `json:"Tasks,omitempty" name:"Tasks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksByTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTasksByTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IspDetail struct {
	// 运营商名称

	IspName *string `json:"IspName,omitempty" name:"IspName"`
	// 可用率

	AvailRatio *float64 `json:"AvailRatio,omitempty" name:"AvailRatio"`
}

type CreateTaskExRequest struct {
	*tchttp.BaseRequest

	// 拨测分组id，体现本拨测任务要采用那些运营商作为拨测源。一般可直接填写本用户的默认拨测分组。参见：DescribeAgentGroupList 接口，本参数使用返回结果里的groupId的值。

	AgentGroupId *uint64 `json:"AgentGroupId,omitempty" name:"AgentGroupId"`
	// http, https, ping, tcp, ftp, smtp, udp, dns 之一

	CatTypeName *string `json:"CatTypeName,omitempty" name:"CatTypeName"`
	// 拨测的url  例如：www.baidu.com (url域名解析需要能解析出具体的ip)

	Url *string `json:"Url,omitempty" name:"Url"`
	// 指定域名(如需要)

	Host *string `json:"Host,omitempty" name:"Host"`
	// 是否为Header请求（非0 发起Header 请求。为0，且PostData 非空，发起POST请求。为0，PostData 为空，发起GET请求）

	IsHeader *uint64 `json:"IsHeader,omitempty" name:"IsHeader"`
	// url中含有https时有用。缺省为SSLv23。需要为 TLSv1_2, TLSv1_1, TLSv1, SSLv2, SSLv23, SSLv3 之一

	SslVer *string `json:"SslVer,omitempty" name:"SslVer"`
	// POST 请求数据。空字符串表示非POST请求

	PostData *string `json:"PostData,omitempty" name:"PostData"`
	// 用户agent 信息

	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`
	// 要在结果中进行匹配的字符串

	CheckStr *string `json:"CheckStr,omitempty" name:"CheckStr"`
	// 1 表示通过检查结果是否包含CheckStr 进行校验

	CheckType *string `json:"CheckType,omitempty" name:"CheckType"`
	// 需要设置的cookie信息

	Cookie *string `json:"Cookie,omitempty" name:"Cookie"`
	// 拨测周期。取值可为1,5,15,30之一, 单位：分钟。精度不能低于用户等级规定的最小精度

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 拨测任务名称不能超过32个字符。同一个用户创建的任务名不可重复

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务号。用于验证且修改任务时传入原任务号

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 登陆服务器的账号。如果为空字符串，表示不用校验用户密码。只做简单连接服务器的拨测。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 登陆服务器的密码

	PassWord *string `json:"PassWord,omitempty" name:"PassWord"`
	// 缺省为0。0 表示请求为字符串类型。1表示为二进制类型

	ReqDataType *uint64 `json:"ReqDataType,omitempty" name:"ReqDataType"`
	// 发起tcp, udp请求的协议请求数据

	ReqData *string `json:"ReqData,omitempty" name:"ReqData"`
	// 缺省为0。0 表示请求为字符串类型。1表示为二进制类型

	RespDataType *uint64 `json:"RespDataType,omitempty" name:"RespDataType"`
	// 预期的udp请求的回应数据。字符串型，只需要返回的结果里包含本字符串算校验通过。二进制型，则需要严格等于才算通过

	RespData *string `json:"RespData,omitempty" name:"RespData"`
	// 目的dns服务器  可以为空字符串

	DnsSvr *string `json:"DnsSvr,omitempty" name:"DnsSvr"`
	// 需要检验是否在dns ip列表的ip。可以为空字符串，表示不校验

	DnsCheckIp *string `json:"DnsCheckIp,omitempty" name:"DnsCheckIp"`
	// 需要为下列值之一。缺省为A。A, MX, NS, CNAME, TXT, ANY

	DnsQueryType *string `json:"DnsQueryType,omitempty" name:"DnsQueryType"`
	// 是否使用安全链接ssl  0 不使用，1 使用

	UseSecConn *uint64 `json:"UseSecConn,omitempty" name:"UseSecConn"`
	// ftp登陆验证方式  0 不验证  1 匿名登陆  2 需要身份验证

	NeedAuth *uint64 `json:"NeedAuth,omitempty" name:"NeedAuth"`
	// 拨测目标的端口号

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// Type=0 默认 （站点监控）Type=2 可用率监控

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// IsVerify=0 非验证任务 IsVerify=1 验证任务，不传则默认为0

	IsVerify *uint64 `json:"IsVerify,omitempty" name:"IsVerify"`
}

func (r *CreateTaskExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTaskExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTaskByWanIpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 任务列表

		Tasks []*CatTaskInfo `json:"Tasks,omitempty" name:"Tasks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTaskByWanIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTaskByWanIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDailyAvailRatioRequest struct {
	*tchttp.BaseRequest

	// 拨测任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *GetDailyAvailRatioRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDailyAvailRatioRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CatReturnSummary struct {
	// 拨测失败个数

	ResultCount *uint64 `json:"ResultCount,omitempty" name:"ResultCount"`
	// 拨测失败返回码

	ResultCode *uint64 `json:"ResultCode,omitempty" name:"ResultCode"`
	// 拨测失败原因描述

	ErrorReason *string `json:"ErrorReason,omitempty" name:"ErrorReason"`
}

type CreateAlarmPloicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警策略Id

		PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmPloicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlarmPloicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户的拨测任务总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 满足条件的拨测任务列表

		Tasks []*CatTask `json:"Tasks,omitempty" name:"Tasks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAgentGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAgentGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAgentGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTaskExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 拨测任务id。验证通过后，创建任务时使用，传递给CreateTask 接口。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTaskExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTaskExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentGroup struct {
	// 拨测分组Id

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 拨测分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 是否默认拨测分组。1表示是，0表示否

	IsDefault *uint64 `json:"IsDefault,omitempty" name:"IsDefault"`
	// 使用本拨测分组的任务数

	TaskNum *uint64 `json:"TaskNum,omitempty" name:"TaskNum"`
	// 拨测结点列表

	Agents []*CatAgent `json:"Agents,omitempty" name:"Agents"`
	// 拨测结点列表

	GroupDetail []*CatAgent `json:"GroupDetail,omitempty" name:"GroupDetail"`
	// 最大拨测分组数

	MaxGroupNum *uint64 `json:"MaxGroupNum,omitempty" name:"MaxGroupNum"`
}

type GetRespTimeTrendExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据点集合，时延等走势数据

		DataPoints []*DataPoint `json:"DataPoints,omitempty" name:"DataPoints"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRespTimeTrendExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRespTimeTrendExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRespTimeTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据点集合，时延等走势数据

		DataPoints []*DataPoint `json:"DataPoints,omitempty" name:"DataPoints"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRespTimeTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRespTimeTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindAlarmPloicyRequest struct {
	*tchttp.BaseRequest

	// 拨测任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 告警策略组Id

	PolicyGroupId *uint64 `json:"PolicyGroupId,omitempty" name:"PolicyGroupId"`
	// 是否绑定操作。非0 为绑定， 0 为 解绑。缺省表示 绑定。

	IfBind *uint64 `json:"IfBind,omitempty" name:"IfBind"`
	// 告警主题Id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *BindAlarmPloicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindAlarmPloicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAgentGroupRequest struct {
	*tchttp.BaseRequest

	// 拨测分组id

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteAgentGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAgentGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmPloicyRequest struct {
	*tchttp.BaseRequest

	// 正整数。拨测任务id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 持续周期。值为任务的Period 乘以0、1、2、3、4。单位：分钟

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 目前取值仅支持 lt (小于)。

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// 门限百分比。比如：80，表示80%。成功率低于80%时告警。

	Threshold *uint64 `json:"Threshold,omitempty" name:"Threshold"`
	// 告警接收组的id。参见： DescribeAlarmGroups 接口。从返回结果里的GroupId 中选取一个。

	ReceiverGroupId *uint64 `json:"ReceiverGroupId,omitempty" name:"ReceiverGroupId"`
}

func (r *CreateAlarmPloicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlarmPloicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmGroupsRequest struct {
	*tchttp.BaseRequest

	// 满足条件的第几条开始

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每批多少条

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAlarmGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRealAvailRatioResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 整体平均可用率

		AvgAvailRatio *float64 `json:"AvgAvailRatio,omitempty" name:"AvgAvailRatio"`
		// 各省份最低可用率

		LowestAvailRatio *float64 `json:"LowestAvailRatio,omitempty" name:"LowestAvailRatio"`
		// 可用率最低的省份

		LowestProvince *string `json:"LowestProvince,omitempty" name:"LowestProvince"`
		// 可用率最低的运营商

		LowestIsp *string `json:"LowestIsp,omitempty" name:"LowestIsp"`
		// 分省份的可用率数据

		ProvinceData []*ProvinceDetail `json:"ProvinceData,omitempty" name:"ProvinceData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRealAvailRatioResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRealAvailRatioResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmsRequest struct {
	*tchttp.BaseRequest

	// 从第Offset 条开始查询。缺省值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 本批次查询Limit 条记录。缺省值为20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 0 全部, 1 已恢复, 2 未恢复  默认为0。其他值，视为0 查全部状态。

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 格式如：2017-05-09 00:00:00  缺省为7天前0点

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 格式如：2017-05-10 00:00:00  缺省为明天0点

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 告警任务名（模糊匹配）

	ObjName *string `json:"ObjName,omitempty" name:"ObjName"`
	// 排序字段，可为Time, ObjName, Duration, Status, Content 之一。缺省为Time。

	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`
	// 升序或降序。可为Desc, Asc之一。缺省为Desc。

	SortType *string `json:"SortType,omitempty" name:"SortType"`
}

func (r *DescribeAlarmsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTaskRequest struct {
	*tchttp.BaseRequest

	// 拨测分组id，体现本拨测任务要采用那些运营商作为拨测源。一般可直接填写本用户的默认拨测分组。参见：DescribeAgentGroupList 接口，本参数使用返回结果里的groupId的值。

	AgentGroupId *uint64 `json:"AgentGroupId,omitempty" name:"AgentGroupId"`
	// http, https, ping, tcp 之一

	CatTypeName *string `json:"CatTypeName,omitempty" name:"CatTypeName"`
	// 拨测的url  例如：www.baidu.com (url域名解析需要能解析出具体的ip)

	Url *string `json:"Url,omitempty" name:"Url"`
	// 需要满足ip 的格式

	Host *string `json:"Host,omitempty" name:"Host"`
	// 服务端监听或接收数据的端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 是否为Header请求（非0 发起Header 请求。为0，且PostData 非空，发起POST请求。为0，PostData 为空，发起GET请求）

	IsHeader *uint64 `json:"IsHeader,omitempty" name:"IsHeader"`
	// url中含有https时有用。缺省为SSLv23。需要为 TLSv1_2, TLSv1_1, TLSv1, SSLv2, SSLv23, SSLv3 之一

	SslVer *string `json:"SslVer,omitempty" name:"SslVer"`
	// POST 请求数据。空字符串表示非POST请求

	PostData *string `json:"PostData,omitempty" name:"PostData"`
	// 用户agent 信息

	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`
	// 要在结果中进行匹配的字符串

	CheckStr *string `json:"CheckStr,omitempty" name:"CheckStr"`
	// 1 表示通过检查结果是否包含checkStr 进行校验

	CheckType *uint64 `json:"CheckType,omitempty" name:"CheckType"`
	// 需要设置的cookie信息

	Cookie *string `json:"Cookie,omitempty" name:"Cookie"`
	// 拨测周期。取值可为1,5,15,30之一, 单位：分钟。精度不能低于用户等级规定的最小精度

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 拨测任务名称不能超过32个字符。同一个用户创建的任务名不可重复

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 验证成功的拨测任务id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ModifyTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmTopic struct {
	// 主题的Id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 主题的名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type CreateAgentGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 拨测分组Id

		GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAgentGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAgentGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 拨测结果查询id。接下来可以使用查询拨测是否能够成功，验证能否通过。

		ResultId *uint64 `json:"ResultId,omitempty" name:"ResultId"`
		// 拨测任务id。验证通过后，创建任务时使用，传递给CreateTask 接口。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentGroupsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAgentGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CatLog struct {
	// 拨测时间点

	Time *string `json:"Time,omitempty" name:"Time"`
	// 拨测类型

	CatTypeName *string `json:"CatTypeName,omitempty" name:"CatTypeName"`
	// 任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 拨测点所在城市

	City *string `json:"City,omitempty" name:"City"`
	// 拨测点所在运营商

	Isp *string `json:"Isp,omitempty" name:"Isp"`
	// 被拨测Server 的Ip

	ServerIp *string `json:"ServerIp,omitempty" name:"ServerIp"`
	// 被拨测Server 的域名

	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
	// 执行耗时，单位毫秒

	TotalTime *uint64 `json:"TotalTime,omitempty" name:"TotalTime"`
	// 成功失败(1 失败，0 成功）

	ResultType *uint64 `json:"ResultType,omitempty" name:"ResultType"`
	// 失败错误码

	ResultCode *uint64 `json:"ResultCode,omitempty" name:"ResultCode"`
	// 请求包大小

	ReqPkgSize *uint64 `json:"ReqPkgSize,omitempty" name:"ReqPkgSize"`
	// 回应包大小

	RspPkgSize *uint64 `json:"RspPkgSize,omitempty" name:"RspPkgSize"`
	// 拨测请求

	ReqMsg *string `json:"ReqMsg,omitempty" name:"ReqMsg"`
	// 拨测回应

	RespMsg *string `json:"RespMsg,omitempty" name:"RespMsg"`
}

type GetReturnCodeInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 拨测失败详情列表

		Details []*CatReturnDetail `json:"Details,omitempty" name:"Details"`
		// 拨测失败汇总列表

		Summary []*CatReturnSummary `json:"Summary,omitempty" name:"Summary"`
		// 开始时间

		BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
		// 截至时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetReturnCodeInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetReturnCodeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPloicyRequest struct {
	*tchttp.BaseRequest

	// 验证成功的拨测任务id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 持续周期。值为任务的Period 乘以0、1、2、3、4。单位：分钟

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 目前取值仅支持 lt (小于)

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// 门限百分比。比如：80，表示80%。成功率低于80%时告警

	Threshold *uint64 `json:"Threshold,omitempty" name:"Threshold"`
	// 拨测告警策略id

	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
	// 告警接收组的id。参见： DescribeAlarmGroups 接口。从返回结果里的GroupId 中选取一个

	ReceiverGroupId *uint64 `json:"ReceiverGroupId,omitempty" name:"ReceiverGroupId"`
}

func (r *ModifyAlarmPloicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmPloicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 拨测任务id 数组

	TaskIds []*uint64 `json:"TaskIds,omitempty" name:"TaskIds"`
}

func (r *DescribeTaskDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRespTimeTrendExRequest struct {
	*tchttp.BaseRequest

	// 验证成功的拨测任务id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 统计数据的发生日期。格式如：2017-05-09

	Date *string `json:"Date,omitempty" name:"Date"`
	// 可为 Isp, Province

	Dimentions []*string `json:"Dimentions,omitempty" name:"Dimentions"`
	// 可为  totalTime, parseTime, connectTime, sendTime, waitTime, receiveTime

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 数据的采集周期，单位分钟

	Period *uint64 `json:"Period,omitempty" name:"Period"`
}

func (r *GetRespTimeTrendExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRespTimeTrendExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主题个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 主题列表

		Topics []*AlarmTopic `json:"Topics,omitempty" name:"Topics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTasksRequest struct {
	*tchttp.BaseRequest

	// 拨测任务id

	TaskIds []*uint64 `json:"TaskIds,omitempty" name:"TaskIds"`
}

func (r *DeleteTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAvailRatioHistoryRequest struct {
	*tchttp.BaseRequest

	// 拨测任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 具体时间点

	TimeStamp *string `json:"TimeStamp,omitempty" name:"TimeStamp"`
}

func (r *GetAvailRatioHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAvailRatioHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAgentGroupRequest struct {
	*tchttp.BaseRequest

	// 拨测分组id

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
	// 拨测分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 是否为默认分组

	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`
	// Province, Isp 需要成对地进行选择。参数对的取值范围。参见：DescribeAgents 的返回结果。

	Agents []*CatAgent `json:"Agents,omitempty" name:"Agents"`
}

func (r *ModifyAgentGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAgentGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTaskExRequest struct {
	*tchttp.BaseRequest

	// 拨测分组id，体现本拨测任务要采用那些运营商作为拨测源。一般可直接填写本用户的默认拨测分组。参见：DescribeAgentGroupList 接口，本参数使用返回结果里的groupId的值。

	AgentGroupId *uint64 `json:"AgentGroupId,omitempty" name:"AgentGroupId"`
	// http, https, ping, tcp, ftp, smtp, udp, dns 之一

	CatTypeName *string `json:"CatTypeName,omitempty" name:"CatTypeName"`
	// 拨测的url  例如：www.baidu.com (url域名解析需要能解析出具体的ip)

	Url *string `json:"Url,omitempty" name:"Url"`
	// 指定域名(如需要)

	Host *string `json:"Host,omitempty" name:"Host"`
	// 服务端监听或接收数据的端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 是否为Header请求（非0 发起Header 请求。为0，且PostData 非空，发起POST请求。为0，PostData 为空，发起GET请求）

	IsHeader *uint64 `json:"IsHeader,omitempty" name:"IsHeader"`
	// url中含有https时有用。缺省为SSLv23。需要为 TLSv1_2, TLSv1_1, TLSv1, SSLv2, SSLv23, SSLv3 之一

	SslVer *string `json:"SslVer,omitempty" name:"SslVer"`
	// POST 请求数据。空字符串表示非POST请求

	PostData *string `json:"PostData,omitempty" name:"PostData"`
	// 用户agent 信息

	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`
	// 要在结果中进行匹配的字符串

	CheckStr *string `json:"CheckStr,omitempty" name:"CheckStr"`
	// 1 表示通过检查结果是否包含checkStr 进行校验

	CheckType *uint64 `json:"CheckType,omitempty" name:"CheckType"`
	// 需要设置的cookie信息

	Cookie *string `json:"Cookie,omitempty" name:"Cookie"`
	// 拨测周期。取值可为1,5,15,30之一, 单位：分钟。精度不能低于用户等级规定的最小精度

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 拨测任务名称不能超过32个字符。同一个用户创建的任务名不可重复

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 验证成功的拨测任务id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 登陆服务器的账号。如果为空字符串，表示不用校验用户密码。只做简单连接服务器的拨测。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 登陆服务器的密码

	PassWord *string `json:"PassWord,omitempty" name:"PassWord"`
	// 缺省为0。0 表示请求为字符串类型。1表示为二进制类型

	ReqDataType *uint64 `json:"ReqDataType,omitempty" name:"ReqDataType"`
	// 发起tcp, udp请求的协议请求数据

	ReqData *string `json:"ReqData,omitempty" name:"ReqData"`
	// 缺省为0。0 表示请求为字符串类型。1表示为二进制类型

	RespDataType *string `json:"RespDataType,omitempty" name:"RespDataType"`
	// 预期的udp请求的回应数据。字符串型，只需要返回的结果里包含本字符串算校验通过。二进制型，则需要严格等于才算通过

	RespData *string `json:"RespData,omitempty" name:"RespData"`
	// 目的dns服务器  可以为空字符串

	DnsSvr *string `json:"DnsSvr,omitempty" name:"DnsSvr"`
	// 需要检验是否在dns ip列表的ip。可以为空字符串，表示不校验

	DnsCheckIp *string `json:"DnsCheckIp,omitempty" name:"DnsCheckIp"`
	// 需要为下列值之一。缺省为A。A, MX, NS, CNAME, TXT, ANY

	DnsQueryType *string `json:"DnsQueryType,omitempty" name:"DnsQueryType"`
	// 是否使用安全链接ssl  0 不使用，1 使用

	UseSecConn *uint64 `json:"UseSecConn,omitempty" name:"UseSecConn"`
	// ftp登陆验证方式  0 不验证  1 匿名登陆  2 需要身份验证

	NeedAuth *uint64 `json:"NeedAuth,omitempty" name:"NeedAuth"`
	// Type=0 默认 （站点监控） Type=2 可用率监控

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *ModifyTaskExRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTaskExRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataPoint struct {
	// 数据点的时间

	LogTime *string `json:"LogTime,omitempty" name:"LogTime"`
	// 数据值

	MetricValue *float64 `json:"MetricValue,omitempty" name:"MetricValue"`
}

type DescribeAgentGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 拨测分组Id

		GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
		// 拨测分组名称

		GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
		// 是否为默认拨测分组

		IsDefault *uint64 `json:"IsDefault,omitempty" name:"IsDefault"`
		// 使用本拨测分组的任务数

		TaskNum *uint64 `json:"TaskNum,omitempty" name:"TaskNum"`
		// 拨测分组运营商列表

		Agents []*CatAgent `json:"Agents,omitempty" name:"Agents"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAgentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 拨测任务列表

		Tasks []*CatTask `json:"Tasks,omitempty" name:"Tasks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskAlarm struct {
	// 任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务周期，单位为分钟。目前支持1，5，15，30几种取值

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 拨测类型。http, https, ping, tcp, udp, smtp, pop3, dns 之一

	CatTypeName *string `json:"CatTypeName,omitempty" name:"CatTypeName"`
	// 任务状态。1表示暂停，2表示运行中，0为初始态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 拨测任务的Url

	CgiUrl *string `json:"CgiUrl,omitempty" name:"CgiUrl"`
	// 任务创建时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 告警状态。1 故障，0 正常

	AlarmStatus *uint64 `json:"AlarmStatus,omitempty" name:"AlarmStatus"`
	// 告警状态描述，统计信息

	StatusInfo *string `json:"StatusInfo,omitempty" name:"StatusInfo"`
	// 任务更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type VerifyResultRequest struct {
	*tchttp.BaseRequest

	// 要查询的拨测任务的结果id

	ResultId *uint64 `json:"ResultId,omitempty" name:"ResultId"`
}

func (r *VerifyResultRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyResultRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CatTaskDetail struct {
	// 任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务周期，单位为分钟。目前支持1，5，15，30几种取值

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 拨测类型。http, https, ping, tcp 之一

	CatTypeName *string `json:"CatTypeName,omitempty" name:"CatTypeName"`
	// 任务状态。1表示暂停，2表示运行中，0为初始态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 拨测任务的Url

	CgiUrl *string `json:"CgiUrl,omitempty" name:"CgiUrl"`
	// 任务创建时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 拨测分组id

	AgentGroupId *uint64 `json:"AgentGroupId,omitempty" name:"AgentGroupId"`
	// 告警策略组id

	PolicyGroupId *uint64 `json:"PolicyGroupId,omitempty" name:"PolicyGroupId"`
	// 任务类型。0 站点监控，2 可用性监控

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 绑定的统一告警主题Id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

type ProvinceDetail struct {
	// 可用率

	AvgAvailRatio *float64 `json:"AvgAvailRatio,omitempty" name:"AvgAvailRatio"`
	// 省份名称

	ProvinceName *string `json:"ProvinceName,omitempty" name:"ProvinceName"`
	// 省份英文名称

	Mapkey *string `json:"Mapkey,omitempty" name:"Mapkey"`
	// 统计时间点

	TimeStamp *string `json:"TimeStamp,omitempty" name:"TimeStamp"`
	// 分运营商可用率

	IspDetail []*IspDetail `json:"IspDetail,omitempty" name:"IspDetail"`
}

type DescribeTaskRequest struct {
	*tchttp.BaseRequest

	// 拨测任务id 数组

	TaskIds []*uint64 `json:"TaskIds,omitempty" name:"TaskIds"`
}

func (r *DescribeTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserLimitRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserLimitRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInfo struct {
	// 告警对象的名称

	ObjName *string `json:"ObjName,omitempty" name:"ObjName"`
	// 告警发生的时间

	FirstOccurTime *string `json:"FirstOccurTime,omitempty" name:"FirstOccurTime"`
	// 告警结束的时间

	LastOccurTime *string `json:"LastOccurTime,omitempty" name:"LastOccurTime"`
	// 告警状态。1 表示已恢复，0 表示未恢复，2表示数据不足

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 告警的内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

type DescribeAgentGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户所属的系统默认拨测分组

		SysDefaultGroup *AgentGroup `json:"SysDefaultGroup,omitempty" name:"SysDefaultGroup"`
		// 用户创建的拨测分组列表

		CustomGroups []*AgentGroup `json:"CustomGroups,omitempty" name:"CustomGroups"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunTaskRequest struct {
	*tchttp.BaseRequest

	// 任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *RunTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRealAvailRatioRequest struct {
	*tchttp.BaseRequest

	// 拨测任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *GetRealAvailRatioRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRealAvailRatioRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResultSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实时统计数据

		RealData []*ResultSummary `json:"RealData,omitempty" name:"RealData"`
		// 按天的统计数据

		DayData []*ResultSummary `json:"DayData,omitempty" name:"DayData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetResultSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResultSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetReturnCodeHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 拨测失败详情列表

		Details []*CatReturnDetail `json:"Details,omitempty" name:"Details"`
		// 拨测失败汇总列表

		Summarys []*CatReturnSummary `json:"Summarys,omitempty" name:"Summarys"`
		// 开始时间

		BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
		// 截至时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetReturnCodeHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetReturnCodeHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetReturnCodeInfoRequest struct {
	*tchttp.BaseRequest

	// 正整数。验证成功的拨测任务id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 开始时间点。格式如：2017-05-09 10:20:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间点。格式如：2017-05-09 10:25:00

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 省份名称的全拼

	Province *string `json:"Province,omitempty" name:"Province"`
}

func (r *GetReturnCodeInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetReturnCodeInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PauseTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PauseTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PauseTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCatLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 拨测记录列表

		CatLogs []*CatLog `json:"CatLogs,omitempty" name:"CatLogs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCatLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCatLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRespTimeTrendRequest struct {
	*tchttp.BaseRequest

	// 验证成功的拨测任务id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 统计数据的发生日期。格式如：2017-05-09

	Date *string `json:"Date,omitempty" name:"Date"`
	// 可为 Isp, Province

	Dimentions []*string `json:"Dimentions,omitempty" name:"Dimentions"`
	// 可为  totalTime, parseTime, connectTime, sendTime, waitTime, receiveTime

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 数据的采集周期，单位分钟

	Period *uint64 `json:"Period,omitempty" name:"Period"`
}

func (r *GetRespTimeTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRespTimeTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyResultResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 错误的原因

		ErrorReason []*string `json:"ErrorReason,omitempty" name:"ErrorReason"`
		// 错误号

		ResultCode *int64 `json:"ResultCode,omitempty" name:"ResultCode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyResultResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CatTask struct {
	// 任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务周期，单位为分钟。目前支持1，5，15，30几种取值

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 拨测类型。http, https, ping, tcp 之一

	CatTypeName *string `json:"CatTypeName,omitempty" name:"CatTypeName"`
	// 任务状态。1表示暂停，2表示运行中，0为初始态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 拨测任务的Url

	CgiUrl *string `json:"CgiUrl,omitempty" name:"CgiUrl"`
	// 任务创建时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 告警的可用率门限(0~100)

	AvailRatioThres *uint64 `json:"AvailRatioThres,omitempty" name:"AvailRatioThres"`
	// 告警的可用率持续时间。值为Period的0~4倍

	AvailRatioInterval *uint64 `json:"AvailRatioInterval,omitempty" name:"AvailRatioInterval"`
	// 告警接收组

	ReceiverGroupId *uint64 `json:"ReceiverGroupId,omitempty" name:"ReceiverGroupId"`
	// 拨测分组id

	AgentGroupId *uint64 `json:"AgentGroupId,omitempty" name:"AgentGroupId"`
}

type GetAvailRatioHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 整体平均可用率

		AvgAvailRatio *float64 `json:"AvgAvailRatio,omitempty" name:"AvgAvailRatio"`
		// 各省份最低可用率

		LowestAvailRatio *float64 `json:"LowestAvailRatio,omitempty" name:"LowestAvailRatio"`
		// 可用率最低的省份

		LowestProvince *string `json:"LowestProvince,omitempty" name:"LowestProvince"`
		// 可用率最低的运营商

		LowestIsp *string `json:"LowestIsp,omitempty" name:"LowestIsp"`
		// 分省份的可用率数据

		ProvinceData []*ProvinceDetail `json:"ProvinceData,omitempty" name:"ProvinceData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAvailRatioHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAvailRatioHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetReturnCodeHistoryRequest struct {
	*tchttp.BaseRequest

	// 正整数。验证成功的拨测任务id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 开始时间点。格式如：2017-05-09 10:20:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间点。格式如：2017-05-09 10:25:00

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 省份名称的全拼

	Province *string `json:"Province,omitempty" name:"Province"`
}

func (r *GetReturnCodeHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetReturnCodeHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 拨测任务列表

		Tasks []*CatTaskDetail `json:"Tasks,omitempty" name:"Tasks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmsByTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警信息列表

		AlarmInfos []*AlarmInfo `json:"AlarmInfos,omitempty" name:"AlarmInfos"`
		// 故障率

		FaultRatio *float64 `json:"FaultRatio,omitempty" name:"FaultRatio"`
		// 故障总时长

		FaultTimeSpec *string `json:"FaultTimeSpec,omitempty" name:"FaultTimeSpec"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmsByTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmsByTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// 从第Offset 条开始查询。缺省值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 本批次查询Limit 条记录。缺省值为20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 任务所使用的拨测分组Id

	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResultSummaryRequest struct {
	*tchttp.BaseRequest

	// 任务Id列表

	TaskIds []*uint64 `json:"TaskIds,omitempty" name:"TaskIds"`
}

func (r *GetResultSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResultSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PauseTaskRequest struct {
	*tchttp.BaseRequest

	// 拨测任务id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *PauseTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PauseTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTaskExResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 拨测结果查询id。接下来可以使用查询拨测是否能够成功，验证能否通过。

		ResultId *uint64 `json:"ResultId,omitempty" name:"ResultId"`
		// 拨测任务id。验证通过后，创建任务时使用，传递给CreateTask 接口。

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTaskExResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTaskExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksByTypeRequest struct {
	*tchttp.BaseRequest

	// 从第Offset 条开始查询。缺省值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 本批次查询Limit 条记录。缺省值为20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 拨测任务类型。0 站点监控，2 可用性监控

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeTasksByTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTasksByTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本用户可选的拨测点列表

		Agents []*CatAgent `json:"Agents,omitempty" name:"Agents"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDailyAvailRatioResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 整体平均可用率

		AvgAvailRatio *float64 `json:"AvgAvailRatio,omitempty" name:"AvgAvailRatio"`
		// 各省份最低可用率

		LowestAvailRatio *float64 `json:"LowestAvailRatio,omitempty" name:"LowestAvailRatio"`
		// 可用率最低的省份

		LowestProvince *string `json:"LowestProvince,omitempty" name:"LowestProvince"`
		// 分省份的可用率数据

		ProvinceData []*ProvinceDetail `json:"ProvinceData,omitempty" name:"ProvinceData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDailyAvailRatioResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDailyAvailRatioResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CatReturnDetail struct {
	// 运营商名称

	IspName *string `json:"IspName,omitempty" name:"IspName"`
	// 省份全拼

	Province *string `json:"Province,omitempty" name:"Province"`
	// 省份中文名称

	ProvinceName *string `json:"ProvinceName,omitempty" name:"ProvinceName"`
	// Map键值

	MapKey *string `json:"MapKey,omitempty" name:"MapKey"`
	// 拨测目标的ip

	ServerIp *string `json:"ServerIp,omitempty" name:"ServerIp"`
	// 拨测失败个数

	ResultCount *uint64 `json:"ResultCount,omitempty" name:"ResultCount"`
	// 拨测失败返回码

	ResultCode *uint64 `json:"ResultCode,omitempty" name:"ResultCode"`
	// 拨测失败原因描述

	ErrorReason *string `json:"ErrorReason,omitempty" name:"ErrorReason"`
}
