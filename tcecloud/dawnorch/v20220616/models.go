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

package v20220616

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DescribeLogsSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		NodeLogs []*NodeLog `json:"NodeLogs,omitempty" name:"NodeLogs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogsSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogsSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeManualParamsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Params []*ManualParams `json:"Params,omitempty" name:"Params"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeManualParamsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeManualParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeLogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		NodeLog *NodeLog `json:"NodeLog,omitempty" name:"NodeLog"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNodeLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterPluginResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RegisterPluginResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RerunNodeRequest struct {
	*tchttp.BaseRequest

	//

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	//

	DagNodeUUID *string `json:"DagNodeUUID,omitempty" name:"DagNodeUUID"`
	// Synchronous 要求API Server按照同步接口返回, 即等待操作完成后才返回

	Synchronous *bool `json:"Synchronous,omitempty" name:"Synchronous"`
}

func (r *RerunNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RerunNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetManualParamsRequest struct {
	*tchttp.BaseRequest

	// JobId is the id of job

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// Params is the params to be set

	Params *ManualParams `json:"Params,omitempty" name:"Params"`
}

func (r *SetManualParamsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetManualParamsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeLog struct {
	//

	NodeInfo *DagNode `json:"NodeInfo,omitempty" name:"NodeInfo"`
	//

	LogGroup []*LogItem `json:"LogGroup,omitempty" name:"LogGroup"`
}

type Record struct {
	//

	Message *string `json:"Message,omitempty" name:"Message"`
	//

	Time *string `json:"Time,omitempty" name:"Time"`
	//

	Level *string `json:"Level,omitempty" name:"Level"`
}

type PauseNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PauseNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PauseNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResumeNodeRequest struct {
	*tchttp.BaseRequest

	//

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	//

	DagNodeUUID *string `json:"DagNodeUUID,omitempty" name:"DagNodeUUID"`
	// Synchronous 要求API Server按照同步接口返回, 即等待操作完成后才返回

	Synchronous *bool `json:"Synchronous,omitempty" name:"Synchronous"`
}

func (r *ResumeNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResumeNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DagRunSpec struct {
	//

	Nodes []*DagNode `json:"Nodes,omitempty" name:"Nodes"`
	//

	Params []*Param `json:"Params,omitempty" name:"Params"`
	//

	Description *string `json:"Description,omitempty" name:"Description"`
	//

	JobUUID *string `json:"JobUUID,omitempty" name:"JobUUID"`
	// ContextScope defines the scope of context in a job

	ContextScope *string `json:"ContextScope,omitempty" name:"ContextScope"`
	// CollectOutputStrategy defines how to collect output, default to 'none' if not set

	CollectOutputStrategy *string `json:"CollectOutputStrategy,omitempty" name:"CollectOutputStrategy"`
	// Output specifies how to extract output of each node and combine to a unified dag level output

	Output []*KVPair `json:"Output,omitempty" name:"Output"`
	// MaxConcurrentNodes is the number of max nodes allowed to run concurrently
	// MaxConcurrentNodes can be a go-template format like <no value>, or just "3"

	MaxConcurrentNodes *string `json:"MaxConcurrentNodes,omitempty" name:"MaxConcurrentNodes"`
	//

	DagReference *string `json:"DagReference,omitempty" name:"DagReference"`
	//

	DagReferenceUUID *string `json:"DagReferenceUUID,omitempty" name:"DagReferenceUUID"`
	//

	TargetState *string `json:"TargetState,omitempty" name:"TargetState"`
	// NodeUUID indicates that this dagrun is also a node of another dagrun

	NodeUUID *string `json:"NodeUUID,omitempty" name:"NodeUUID"`
	//

	Input []*KVPair `json:"Input,omitempty" name:"Input"`
	//

	Depth *int64 `json:"Depth,omitempty" name:"Depth"`
	//

	NodePath *string `json:"NodePath,omitempty" name:"NodePath"`
}

type DescribeLogsDAGRequest struct {
	*tchttp.BaseRequest

	// 任务Id

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 搜索起始父节点的UUID

	DagNodeUUID *string `json:"DagNodeUUID,omitempty" name:"DagNodeUUID"`
	// 广度优先搜索深度

	Deep *int64 `json:"Deep,omitempty" name:"Deep"`
}

func (r *DescribeLogsDAGRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogsDAGRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterPluginRequest struct {
	*tchttp.BaseRequest

	// Kind is the kind that the plugin serves for

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// Version is the version of plugin

	Version *string `json:"Version,omitempty" name:"Version"`
	// Capabilities is the capabilities(name of implemented action) of the plugin.
	// e.g. ["RunNode", "QueryNodeStatus"]

	Capabilities []*string `json:"Capabilities,omitempty" name:"Capabilities"`
	// Endpoint the is endpoint of plugin, e.g. http://196.168.0.1:8080

	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`
}

func (r *RegisterPluginRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RegisterPluginRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Param struct {
	//

	Name *string `json:"Name,omitempty" name:"Name"`
	//

	Description *string `json:"Description,omitempty" name:"Description"`
	//

	Type *string `json:"Type,omitempty" name:"Type"`
	//

	Default *string `json:"Default,omitempty" name:"Default"`
	//

	Value *string `json:"Value,omitempty" name:"Value"`
}

type PauseNodeRequest struct {
	*tchttp.BaseRequest

	//

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	//

	DagNodeUUID *string `json:"DagNodeUUID,omitempty" name:"DagNodeUUID"`
	// Synchronous 要求API Server按照同步接口返回, 即等待操作完成后才返回

	Synchronous *bool `json:"Synchronous,omitempty" name:"Synchronous"`
}

func (r *PauseNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PauseNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJobOverviewRequest struct {
	*tchttp.BaseRequest

	//

	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeJobOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJobOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DagNodeStatus struct {
	//

	Name *string `json:"Name,omitempty" name:"Name"`
	//

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	//

	CurrentState *string `json:"CurrentState,omitempty" name:"CurrentState"`
	//

	TargetState *string `json:"TargetState,omitempty" name:"TargetState"`
	//

	DisplayStatus *string `json:"DisplayStatus,omitempty" name:"DisplayStatus"`
	//

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
	//

	Output []*KVPair `json:"Output,omitempty" name:"Output"`
	// auroraservice:override:type=string

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// auroraservice:override:type=string

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// auroraservice:override:type=string

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	//

	SubDagRunUUID *string `json:"SubDagRunUUID,omitempty" name:"SubDagRunUUID"`
}

type KVPair struct {
	//

	Key *string `json:"Key,omitempty" name:"Key"`
	//

	Value *string `json:"Value,omitempty" name:"Value"`
	//

	Options *Options `json:"Options,omitempty" name:"Options"`
}

type ManualParams struct {
	//

	NodePath *string `json:"NodePath,omitempty" name:"NodePath"`
	//

	NodeUUID *string `json:"NodeUUID,omitempty" name:"NodeUUID"`
	//

	Input []*KVPair `json:"Input,omitempty" name:"Input"`
	//

	Condition *string `json:"Condition,omitempty" name:"Condition"`
}

type PauseJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PauseJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PauseJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DagRun struct {
	// Name is the name of this object, unique in a scope

	Name *string `json:"Name,omitempty" name:"Name"`
	// UUID is the unique id of this object

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// Labels is a list of KV pairt

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// Annotations is a list of KV pairt

	Annotations []*KVPair `json:"Annotations,omitempty" name:"Annotations"`
	// Version is the object's version

	Version *string `json:"Version,omitempty" name:"Version"`
	// Kind is the object's kind

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// CreatedAt is the creation timestamp

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// UpdatedAt is the last update timestamp

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	//

	Spec *DagRunSpec `json:"Spec,omitempty" name:"Spec"`
	//

	Status *DagRunStatus `json:"Status,omitempty" name:"Status"`
}

type Options struct {
	//

	ShouldNotRender *bool `json:"ShouldNotRender,omitempty" name:"ShouldNotRender"`
	//

	BehaviorOnRenderError *string `json:"BehaviorOnRenderError,omitempty" name:"BehaviorOnRenderError"`
	//

	DefaultValueOnRenderError *string `json:"DefaultValueOnRenderError,omitempty" name:"DefaultValueOnRenderError"`
	// Format indicates the value's format, like json/yaml/text

	Format *string `json:"Format,omitempty" name:"Format"`
}

type DescribeJobDAGResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		DagNode *DagNode `json:"DagNode,omitempty" name:"DagNode"`
		//

		DagNodeList []*DagNode `json:"DagNodeList,omitempty" name:"DagNodeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJobDAGResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJobDAGResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetContextResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Context []*KVPair `json:"Context,omitempty" name:"Context"`
		// `text` or `binary`

		Type *string `json:"Type,omitempty" name:"Type"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetContextResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetContextResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetWaitingNodeStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetWaitingNodeStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetWaitingNodeStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJobDAGRequest struct {
	*tchttp.BaseRequest

	//

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	//

	DagNodeUUID *string `json:"DagNodeUUID,omitempty" name:"DagNodeUUID"`
	//

	Deep *int64 `json:"Deep,omitempty" name:"Deep"`
	//

	NodePath *string `json:"NodePath,omitempty" name:"NodePath"`
	//

	NodeRunColumns []*string `json:"NodeRunColumns,omitempty" name:"NodeRunColumns"`
}

func (r *DescribeJobDAGRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJobDAGRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetContextRequest struct {
	*tchttp.BaseRequest

	//

	JobUUID *string `json:"JobUUID,omitempty" name:"JobUUID"`
	//

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// Key is optional

	Key *string `json:"Key,omitempty" name:"Key"`
}

func (r *GetContextRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetContextRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PauseJobRequest struct {
	*tchttp.BaseRequest

	//

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// Synchronous 要求API Server按照同步接口返回, 即等待操作完成后才返回

	Synchronous *bool `json:"Synchronous,omitempty" name:"Synchronous"`
}

func (r *PauseJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PauseJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateJobRequest struct {
	*tchttp.BaseRequest

	//

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// Synchronous 要求API Server按照同步接口返回, 即等待操作完成后才返回

	Synchronous *bool `json:"Synchronous,omitempty" name:"Synchronous"`
}

func (r *TerminateJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateJobRequest struct {
	*tchttp.BaseRequest

	// Kind 编排对象类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// Version 编排对象版本，默认最新版本

	Version *string `json:"Version,omitempty" name:"Version"`
	//

	Business *string `json:"Business,omitempty" name:"Business"`
	// DagRun 可以直接指定DagRun和DagList创建任务, 不依赖插件对Business转换

	DagRun *DagRun `json:"DagRun,omitempty" name:"DagRun"`
	//

	DagList []*Dag `json:"DagList,omitempty" name:"DagList"`
}

func (r *CreateJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogsDAGResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 父节点

		DagNode *DagNode `json:"DagNode,omitempty" name:"DagNode"`
		// 节点日志列表

		NodeRunLog []*NodeLog `json:"NodeRunLog,omitempty" name:"NodeRunLog"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogsDAGResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogsDAGResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeLogsRequest struct {
	*tchttp.BaseRequest

	//

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	//

	DagNodeUUID *string `json:"DagNodeUUID,omitempty" name:"DagNodeUUID"`
}

func (r *DescribeNodeLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeLogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Dag struct {
	// Name is the name of this object, unique in a scope

	Name *string `json:"Name,omitempty" name:"Name"`
	// UUID is the unique id of this object

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// Labels is a list of KV pairt

	Labels []*KVPair `json:"Labels,omitempty" name:"Labels"`
	// Annotations is a list of KV pairt

	Annotations []*KVPair `json:"Annotations,omitempty" name:"Annotations"`
	// Version is the object's version

	Version *string `json:"Version,omitempty" name:"Version"`
	// Kind is the object's kind

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// CreatedAt is the creation timestamp

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// UpdatedAt is the last update timestamp

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	//

	Spec *DagSpec `json:"Spec,omitempty" name:"Spec"`
}

type DagSpec struct {
	//

	Nodes []*DagNode `json:"Nodes,omitempty" name:"Nodes"`
	//

	Params []*Param `json:"Params,omitempty" name:"Params"`
	//

	Description *string `json:"Description,omitempty" name:"Description"`
	//

	JobUUID *string `json:"JobUUID,omitempty" name:"JobUUID"`
	// ContextScope defines the scope of context in a job

	ContextScope *string `json:"ContextScope,omitempty" name:"ContextScope"`
	// CollectOutputStrategy defines how to collect output, default to 'none' if not set

	CollectOutputStrategy *string `json:"CollectOutputStrategy,omitempty" name:"CollectOutputStrategy"`
	// Output specifies how to extract output of each node and combine to a unified dag level output

	Output []*KVPair `json:"Output,omitempty" name:"Output"`
	// MaxConcurrentNodes is the number of max nodes allowed to run concurrently
	// MaxConcurrentNodes can be a go-template format like <no value>, or just "3"

	MaxConcurrentNodes *string `json:"MaxConcurrentNodes,omitempty" name:"MaxConcurrentNodes"`
}

type DescribeJobRequest struct {
	*tchttp.BaseRequest

	//

	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResumeJobRequest struct {
	*tchttp.BaseRequest

	//

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// Synchronous 要求API Server按照同步接口返回, 即等待操作完成后才返回

	Synchronous *bool `json:"Synchronous,omitempty" name:"Synchronous"`
}

func (r *ResumeJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResumeJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogsSummaryRequest struct {
	*tchttp.BaseRequest

	// 任务Id

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// 搜索起始父节点UUID

	DagNodeUUID *string `json:"DagNodeUUID,omitempty" name:"DagNodeUUID"`
	// 筛选节点类型

	Kinds []*string `json:"Kinds,omitempty" name:"Kinds"`
	// 搜索深度

	Deep *int64 `json:"Deep,omitempty" name:"Deep"`
}

func (r *DescribeLogsSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogsSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResumeJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResumeJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryNodeRequest struct {
	*tchttp.BaseRequest

	//

	DagNodeUUID *string `json:"DagNodeUUID,omitempty" name:"DagNodeUUID"`
	// RetryPolicy 重试策略，是重建当前节点（RebuildCurrentNode），还是重建当前节点下的所有失败状态的叶子节点（RebuildAllFailedLeafNode）

	RetryPolicy *string `json:"RetryPolicy,omitempty" name:"RetryPolicy"`
}

func (r *RetryNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		Name *string `json:"Name,omitempty" name:"Name"`
		//

		Status *string `json:"Status,omitempty" name:"Status"`
		//

		Description *string `json:"Description,omitempty" name:"Description"`
		//

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		//

		FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
		//

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		//

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		//

		Business *string `json:"Business,omitempty" name:"Business"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RerunNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RerunNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RerunNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetWaitingNodeStatusRequest struct {
	*tchttp.BaseRequest

	// DagNodeUUID 表示结点UUID

	DagNodeUUID *string `json:"DagNodeUUID,omitempty" name:"DagNodeUUID"`
	// Status 表示需要将waiting结点设置成继续运行还是失败. 可选值: success/failed

	Status *string `json:"Status,omitempty" name:"Status"`
	// Synchronous 要求API Server按照同步接口返回, 即等待操作完成后才返回

	Synchronous *bool `json:"Synchronous,omitempty" name:"Synchronous"`
}

func (r *SetWaitingNodeStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetWaitingNodeStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeJobOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		//

		JobId *string `json:"JobId,omitempty" name:"JobId"`
		//

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJobOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeJobOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SkipNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SkipNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SkipNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DagRunStatus struct {
	// auroraservice:override:type=string

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// auroraservice:override:type=string

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// auroraservice:override:type=string

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	//

	CurrentState *string `json:"CurrentState,omitempty" name:"CurrentState"`
	//

	DisplayStatus *string `json:"DisplayStatus,omitempty" name:"DisplayStatus"`
	//

	Nodes []*DagNodeStatus `json:"Nodes,omitempty" name:"Nodes"`
	//

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
	//

	Output []*KVPair `json:"Output,omitempty" name:"Output"`
}

type LogItem struct {
	//

	Name *string `json:"Name,omitempty" name:"Name"`
	//

	Description *string `json:"Description,omitempty" name:"Description"`
	//

	Structured *bool `json:"Structured,omitempty" name:"Structured"`
	//

	Type *int64 `json:"Type,omitempty" name:"Type"`
	//

	Log []*Record `json:"Log,omitempty" name:"Log"`
}

type PutContextResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutContextResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PutContextResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResumeNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResumeNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetManualParamsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetManualParamsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetManualParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SkipNodeRequest struct {
	*tchttp.BaseRequest

	//

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	//

	DagNodeUUID *string `json:"DagNodeUUID,omitempty" name:"DagNodeUUID"`
	// Synchronous 要求API Server按照同步接口返回, 即等待操作完成后才返回

	Synchronous *bool `json:"Synchronous,omitempty" name:"Synchronous"`
}

func (r *SkipNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SkipNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeManualParamsRequest struct {
	*tchttp.BaseRequest

	// JobId is the id of job

	JobId *string `json:"JobId,omitempty" name:"JobId"`
	// NodePaths is the node paths to be described

	NodePaths []*string `json:"NodePaths,omitempty" name:"NodePaths"`
}

func (r *DescribeManualParamsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeManualParamsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PutContextRequest struct {
	*tchttp.BaseRequest

	//

	JobUUID *string `json:"JobUUID,omitempty" name:"JobUUID"`
	//

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// context to be put

	Context []*KVPair `json:"Context,omitempty" name:"Context"`
	// `text` or `binary`

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *PutContextRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PutContextRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DagNode struct {
	//

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	//

	NodeUUID *string `json:"NodeUUID,omitempty" name:"NodeUUID"`
	//

	Status *string `json:"Status,omitempty" name:"Status"`
	//

	Name *string `json:"Name,omitempty" name:"Name"`
	//

	RunType *string `json:"RunType,omitempty" name:"RunType"`
	//

	Parallel *bool `json:"Parallel,omitempty" name:"Parallel"`
	//

	RetryLimit *int64 `json:"RetryLimit,omitempty" name:"RetryLimit"`
	//

	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
	//

	Condition *string `json:"Condition,omitempty" name:"Condition"`
	//

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	//

	Inputs *string `json:"Inputs,omitempty" name:"Inputs"`
	//

	Outputs *string `json:"Outputs,omitempty" name:"Outputs"`
	//

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
	//

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	//

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	//

	SubDagNodes []*string `json:"SubDagNodes,omitempty" name:"SubDagNodes"`
	//

	Description *string `json:"Description,omitempty" name:"Description"`
	//

	DescriptionExtended *string `json:"DescriptionExtended,omitempty" name:"DescriptionExtended"`
	//

	NodePath *string `json:"NodePath,omitempty" name:"NodePath"`
}
