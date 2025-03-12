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

package v20211025

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type Host struct {
	// 该主机所在的业务模块

	BizSetModule []*string `json:"BizSetModule,omitempty" name:"BizSetModule"`
	// 母机IP

	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 主机状态

	HostStatus *string `json:"HostStatus,omitempty" name:"HostStatus"`
	// 主机 UUID

	HostUUID *string `json:"HostUUID,omitempty" name:"HostUUID"`
	// 内网IP

	InnerIP *string `json:"InnerIP,omitempty" name:"InnerIP"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 主机 SN

	SN *string `json:"SN,omitempty" name:"SN"`
	// 可用区名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 安装在该主机上的 agent 实例列表

	AgentList []*AgentInstance `json:"AgentList,omitempty" name:"AgentList"`
}

type QueryAgentConfigResponseData struct {
	// agent 配置列表

	AgentConfigList []*AgentConfig `json:"AgentConfigList,omitempty" name:"AgentConfigList"`
}

type InstallResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安装agent的响应体

		Data *OperateResponseData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InstallResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Audit struct {
	// agent 名称

	Agent *string `json:"Agent,omitempty" name:"Agent"`
	// 日志 UUID

	AuditUUID *string `json:"AuditUUID,omitempty" name:"AuditUUID"`
	// 操作的 agent 关联的业务模块

	BizSetModule []*string `json:"BizSetModule,omitempty" name:"BizSetModule"`
	// 主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 内网 IP

	InnerIP *string `json:"InnerIP,omitempty" name:"InnerIP"`
	// 审计日志详情

	Log *string `json:"Log,omitempty" name:"Log"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 操作的 agent 所在的主机的地域

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 操作的 agent 所在的主机的可用区

	Status *string `json:"Status,omitempty" name:"Status"`
	// 该操作所对应的异步任务 UUID

	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`
	// 操作类型/名称

	Type *string `json:"Type,omitempty" name:"Type"`
	// 日志更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 操作的 agent 的版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 可用区名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 日志创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
}

type UpdateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 更新agent响应体

		Data *OperateResponseData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询 agent 关联的业务模块下包含的主机列表响应体

		Data *QueryHostResponseData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 重启agent的响应体

		Data *OperateResponseData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartRequest struct {
	*tchttp.BaseRequest

	// 客户端组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// 目标启动的主机数组

	IPList []*string `json:"IPList,omitempty" name:"IPList"`
	// 是否为主控

	Master *bool `json:"Master,omitempty" name:"Master"`
	// 启动的agent版本

	ExpectedVersion *string `json:"ExpectedVersion,omitempty" name:"ExpectedVersion"`
	// 启动的agent名称

	Agent *string `json:"Agent,omitempty" name:"Agent"`
}

func (r *StartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentConfig struct {
	// agent 名称

	Agent *string `json:"Agent,omitempty" name:"Agent"`
	// 关联的业务模块 id 格式为("bizID>setID>moduleID")

	BizSetModule []*string `json:"BizSetModule,omitempty" name:"BizSetModule"`
	// 关联的业务模块 name 格式为("bizName>setName>moduleName")

	BizSetModuleName []*string `json:"BizSetModuleName,omitempty" name:"BizSetModuleName"`
	// 配置UUID

	ConfigUUID *string `json:"ConfigUUID,omitempty" name:"ConfigUUID"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 是否为主控

	Master *bool `json:"Master,omitempty" name:"Master"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// agent 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type UpdateAgentConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 更新agent配置响应体

		Data *UpdateAgentConfigResponseData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAgentConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAgentConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 索引字段的名称（大驼峰）

	Name *string `json:"Name,omitempty" name:"Name"`
	// 查询的类型（精确匹配 Exact，模糊匹配 Fuzzy，In 查询，范围查询 Range）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 字段的值（若为 In 或者 Range，则使用 "," 分隔）

	Value *string `json:"Value,omitempty" name:"Value"`
}

type StopResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 停止agent的响应体

		Data *OperateResponseData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentInstance struct {
	// agent名称

	Agent *string `json:"Agent,omitempty" name:"Agent"`
	// 是否健康

	Healthy *bool `json:"Healthy,omitempty" name:"Healthy"`
	// 心跳时间

	Heartbeat *string `json:"Heartbeat,omitempty" name:"Heartbeat"`
	// agent 实例 ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 是否为主控

	Master *bool `json:"Master,omitempty" name:"Master"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 状态描述日志

	StatusDescription *string `json:"StatusDescription,omitempty" name:"StatusDescription"`
	// agent 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 该机器上的主控状态

	MasterStatus *string `json:"MasterStatus,omitempty" name:"MasterStatus"`
	// 该机器上的主控版本

	MasterVersion *string `json:"MasterVersion,omitempty" name:"MasterVersion"`
}

type QueryHostResponseData struct {
	// 主机列表

	HostList []*Host `json:"HostList,omitempty" name:"HostList"`
	// 总数目

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type QueryAgentAliasRequest struct {
	*tchttp.BaseRequest

	// 客户端组件名称

	Component *string `json:"Component,omitempty" name:"Component"`
}

func (r *QueryAgentAliasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryAgentAliasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAgentVersionRequest struct {
	*tchttp.BaseRequest

	// 客户端组件名称

	Component *string `json:"Component,omitempty" name:"Component"`
	// Agent名称

	Agent *string `json:"Agent,omitempty" name:"Agent"`
}

func (r *QueryAgentVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryAgentVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAuditResponseData struct {
	// 审计日志列表

	AuditList []*Audit `json:"AuditList,omitempty" name:"AuditList"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type InstallRequest struct {
	*tchttp.BaseRequest

	// 客户端组件名称

	Component *string `json:"Component,omitempty" name:"Component"`
	// 目标安装的主机数组

	IPList []*string `json:"IPList,omitempty" name:"IPList"`
	// 安装的agent是否为主控

	Master *bool `json:"Master,omitempty" name:"Master"`
	// 安装的agent的版本

	ExpectedVersion *string `json:"ExpectedVersion,omitempty" name:"ExpectedVersion"`
	// 安装的agent的名称（唯一）

	Agent *string `json:"Agent,omitempty" name:"Agent"`
}

func (r *InstallRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 卸载agent的响应体

		Data *OperateResponseData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UninstallResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAuditRequest struct {
	*tchttp.BaseRequest

	// 客户端组件名称

	Component *string `json:"Component,omitempty" name:"Component"`
	// 过滤参数数组

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序参数数组

	Sorters []*Sorter `json:"Sorters,omitempty" name:"Sorters"`
	// 分页起始行

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryAuditRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryAuditRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopRequest struct {
	*tchttp.BaseRequest

	// 客户端组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// 目标停止的主机数组

	IPList []*string `json:"IPList,omitempty" name:"IPList"`
	// 是否为主控

	Master *bool `json:"Master,omitempty" name:"Master"`
	// agent版本

	ExpectedVersion *string `json:"ExpectedVersion,omitempty" name:"ExpectedVersion"`
	// agent名称

	Agent *string `json:"Agent,omitempty" name:"Agent"`
}

func (r *StopRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallRequest struct {
	*tchttp.BaseRequest

	// 客户端组件名

	Component *string `json:"Component,omitempty" name:"Component"`
	// 目标卸载的主机列表

	IPList []*string `json:"IPList,omitempty" name:"IPList"`
	// 是否为主控

	Master *bool `json:"Master,omitempty" name:"Master"`
	// agent版本

	ExpectedVersion *string `json:"ExpectedVersion,omitempty" name:"ExpectedVersion"`
	// agent名称

	Agent *string `json:"Agent,omitempty" name:"Agent"`
}

func (r *UninstallRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentAlias struct {
	// Agent名称

	Agent *string `json:"Agent,omitempty" name:"Agent"`
	// Agent别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type QueryAgentAliasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Agent 别名数组

		AgentAliasList []*AgentAlias `json:"AgentAliasList,omitempty" name:"AgentAliasList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAgentAliasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryAgentAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAgentVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 版本列表

		VersionList []*string `json:"VersionList,omitempty" name:"VersionList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAgentVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryAgentVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRequest struct {
	*tchttp.BaseRequest

	// 客户端组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// 目标更新的主机列表

	IPList []*string `json:"IPList,omitempty" name:"IPList"`
	// 是否为主控

	Master *bool `json:"Master,omitempty" name:"Master"`
	// agent版本

	ExpectedVersion *string `json:"ExpectedVersion,omitempty" name:"ExpectedVersion"`
	// agent名称

	Agent *string `json:"Agent,omitempty" name:"Agent"`
}

func (r *UpdateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAgentConfigResponseData struct {
	// 是否更新成功

	Success *bool `json:"Success,omitempty" name:"Success"`
}

type StartResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 启动agent响应体

		Data *OperateResponseData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAgentConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询agent列表的响应体

		Data *QueryAgentConfigResponseData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAgentConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryAgentConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAuditResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询审计日志列表响应体

		Data *QueryAuditResponseData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAuditResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateResponseData struct {
	// 操作的每一台主机所对应的审计日志 UUID

	AuditUUIDs []*string `json:"AuditUUIDs,omitempty" name:"AuditUUIDs"`
}

type Sorter struct {
	// 排序字段名称（大驼峰）

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否为降序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
}

type QueryHostRequest struct {
	*tchttp.BaseRequest

	// 客户端组件名称

	Component *string `json:"Component,omitempty" name:"Component"`
	// 分页的页大小

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页的起始行

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询过滤字段

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 查询的排序字段

	Sorter []*Sorter `json:"Sorter,omitempty" name:"Sorter"`
}

func (r *QueryHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartRequest struct {
	*tchttp.BaseRequest

	// 客户端组件

	Component *string `json:"Component,omitempty" name:"Component"`
	// 目标重启的主机数组

	IPList []*string `json:"IPList,omitempty" name:"IPList"`
	// 是否为主控

	Master *bool `json:"Master,omitempty" name:"Master"`
	// agent版本

	ExpectedVersion *string `json:"ExpectedVersion,omitempty" name:"ExpectedVersion"`
	// agent名称

	Agent *string `json:"Agent,omitempty" name:"Agent"`
}

func (r *RestartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAgentConfigRequest struct {
	*tchttp.BaseRequest

	// 客户端组件名称（前端可硬编码）

	Component *string `json:"Component,omitempty" name:"Component"`
	// 被更新的 agent 配置

	AgentConfig *AgentConfig `json:"AgentConfig,omitempty" name:"AgentConfig"`
}

func (r *UpdateAgentConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAgentConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAgentConfigRequest struct {
	*tchttp.BaseRequest

	// 请求组件名称（前端可硬编码）

	Component *string `json:"Component,omitempty" name:"Component"`
}

func (r *QueryAgentConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryAgentConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
