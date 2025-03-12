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

package v20200217

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type OperateNtpRequest struct {
	*tchttp.BaseRequest

	// Filters

	Filters *Filters `json:"Filters,omitempty" name:"Filters"`
	// 操作列表，字符串0表示修改配置，1表示重启服务，传入其他参数会被忽略

	Operate []*string `json:"Operate,omitempty" name:"Operate"`
	// 配置

	Value *ConfValue `json:"Value,omitempty" name:"Value"`
	// 备注信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

func (r *OperateNtpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OperateNtpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOpRecordsRequest struct {
	*tchttp.BaseRequest

	// Filters

	Filters *Filters `json:"Filters,omitempty" name:"Filters"`
	// Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetOpRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOpRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNtpConfRequest struct {
	*tchttp.BaseRequest

	// Filters

	Filters *Filters `json:"Filters,omitempty" name:"Filters"`
	// Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetNtpConfRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNtpConfRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadClientTimeStampRequest struct {
	*tchttp.BaseRequest

	// 毫秒级时间戳

	ClientTimeStamp *string `json:"ClientTimeStamp,omitempty" name:"ClientTimeStamp"`
}

func (r *UploadClientTimeStampRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadClientTimeStampRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfValue struct {
	// Upstream配置，可为空

	Upstream []*Upstream `json:"Upstream,omitempty" name:"Upstream"`
}

type Upstream struct {
	// Ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type GetMonitorDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可用区个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 可用区监控信息

		AvailableZones []*MonitorAz `json:"AvailableZones,omitempty" name:"AvailableZones"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMonitorDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMonitorDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNtpConfResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可用区的ntp配置信息

		AvailableZones []*AzConf `json:"AvailableZones,omitempty" name:"AvailableZones"`
		// 可用区个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNtpConfResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNtpConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadClientTimeStampResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadClientTimeStampResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadClientTimeStampResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerConf struct {
	// ntp配置

	Conf *ConfValue `json:"Conf,omitempty" name:"Conf"`
	// ntp服务器ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// ntp服务最新的重启时间

	RebootTime *uint64 `json:"RebootTime,omitempty" name:"RebootTime"`
	// ntp服务器状态：1，2，3，4，5，6分别代表：配置修改中，配置修改失败，配置修改成功，服务重启中，服务重启失败，服务重启成功

	State *int64 `json:"State,omitempty" name:"State"`
	// ntp配置最新的修改时间

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type GetMonitorDataRequest struct {
	*tchttp.BaseRequest

	// 筛选

	Filters *Filters `json:"Filters,omitempty" name:"Filters"`
	// Offset

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// Limit

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetMonitorDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMonitorDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filters struct {
	// 服务器内网IP，和可用区参数只能二选一。支持此参数的接口：GetMonitorData，GetOpRecords，OperateNtp

	Ip []*string `json:"Ip,omitempty" name:"Ip"`
	// 可用区，和内网IP参数只能二选一。支持此参数的接口：GetNtpConf，GetMonitorData，GetOpRecords，OperateNtp

	AvailableZone []*string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	// 操作人。支持此参数的接口：GetOpRecords

	OpUser []*string `json:"OpUser,omitempty" name:"OpUser"`
	// 操作类型。支持此参数的接口：GetOpRecords

	OpType []*string `json:"OpType,omitempty" name:"OpType"`
	// 操作记录的状态。支持此参数的接口：GetOpRecords。此字段已废弃，暂时不会在任务中返回

	OpState []*string `json:"OpState,omitempty" name:"OpState"`
	// 操作时间，只能填开始和截止时间的毫秒级时间戳字符串。支持此参数的接口：GetOpRecords

	OpTime []*string `json:"OpTime,omitempty" name:"OpTime"`
}

type MonitorUpStream struct {
	// 对应ntpq的网络延时

	Delay *float64 `json:"Delay,omitempty" name:"Delay"`
	// 上级源IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 对应ntpq的抖动值

	Jitter *float64 `json:"Jitter,omitempty" name:"Jitter"`
	// 最近一次同步上级源的毫秒时间戳，可能为""，说明过于久远没同步成功过了

	LastSync *string `json:"LastSync,omitempty" name:"LastSync"`
	// 对应ntpq的偏移值

	Offset *float64 `json:"Offset,omitempty" name:"Offset"`
	// 对应ntpq的时钟源等级

	Stratum *int64 `json:"Stratum,omitempty" name:"Stratum"`
	// 对应ntpq的同步间隔，秒

	SyncInterval *float64 `json:"SyncInterval,omitempty" name:"SyncInterval"`
	// 对应ntpq的同步状态：空格 = 0；x = 1；- = 2；. = 3；# = 4； + = 5； * = 6； o = 7；

	SyncState *int64 `json:"SyncState,omitempty" name:"SyncState"`
}

type OpRecord struct {
	// 操作目标ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 操作的目标可用区

	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	// 操作者

	OpUser *string `json:"OpUser,omitempty" name:"OpUser"`
	// 操作类型:"0","1"分别代表“修改配置”和“重启服务”

	OpType *string `json:"OpType,omitempty" name:"OpType"`
	// 操作时间（毫秒级时间戳）

	OpTime *string `json:"OpTime,omitempty" name:"OpTime"`
	// 0，1，2分别表示“失败”，“成功”，“进行中”

	OpState *string `json:"OpState,omitempty" name:"OpState"`
	// 操作信息

	OpMessage *string `json:"OpMessage,omitempty" name:"OpMessage"`
}

type GetOpInfoListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可筛选的可用区列表

		AvailableZone *ListValue `json:"AvailableZone,omitempty" name:"AvailableZone"`
		// 可筛选的Ip列表

		Ip *ListValue `json:"Ip,omitempty" name:"Ip"`
		// 可筛选的状态列表

		OpState *ListValue `json:"OpState,omitempty" name:"OpState"`
		// 可筛选的操作类型列表

		OpType *ListValue `json:"OpType,omitempty" name:"OpType"`
		// 可筛选的操作者列表

		OpUser *ListValue `json:"OpUser,omitempty" name:"OpUser"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOpInfoListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOpInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOpRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作记录数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 操作记录

		OpRecords []*OpRecord `json:"OpRecords,omitempty" name:"OpRecords"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOpRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOpRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorAz struct {
	// 可用区名称

	AzName *string `json:"AzName,omitempty" name:"AzName"`
	// 该可用区下的ntp服务器总数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 该可用区下有异常服务器为0，全部正常为1

	AzState *string `json:"AzState,omitempty" name:"AzState"`
	// 可用区下的服务器监控数据

	Servers []*MonitorServer `json:"Servers,omitempty" name:"Servers"`
}

type GetOpInfoListRequest struct {
	*tchttp.BaseRequest
}

func (r *GetOpInfoListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOpInfoListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AzConf struct {
	// 可用区名称

	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	// 可用区下的服务器的ntp配置

	Servers []*ServerConf `json:"Servers,omitempty" name:"Servers"`
}

type MonitorServer struct {
	// ntp配置的上游监控信息

	UpStream []*MonitorUpStream `json:"UpStream,omitempty" name:"UpStream"`
	// 服务器的本地时间

	ServerTime *string `json:"ServerTime,omitempty" name:"ServerTime"`
	// 0，1分别表示异常和正常

	ServerState *string `json:"ServerState,omitempty" name:"ServerState"`
	// 可用区名称

	AzName *string `json:"AzName,omitempty" name:"AzName"`
	// ntp服务器ip

	ServerIp *string `json:"ServerIp,omitempty" name:"ServerIp"`
	// 请求信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type OperateNtpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OperateNtpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OperateNtpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListValue struct {
	// 值列表

	List []*string `json:"List,omitempty" name:"List"`
	// List的元素个数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}
