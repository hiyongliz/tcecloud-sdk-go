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

type CreateConfigMapsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConfigMapsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigMapsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 命名空间

		Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
		// 用户uin

		UserUin *string `json:"UserUin,omitempty" name:"UserUin"`
		// 所属集群标识

		Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
		// 产品名称

		ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
		// 状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// ready的组件数量

		ReadyComponent *int64 `json:"ReadyComponent,omitempty" name:"ReadyComponent"`
		// 组件总数

		TotalComponent *int64 `json:"TotalComponent,omitempty" name:"TotalComponent"`
		// 版本

		AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
		// 描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// 创建时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 应用模板Id

		TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
		// 组件实例详情

		ComponentInstances []*ComponentInstance `json:"ComponentInstances,omitempty" name:"ComponentInstances"`
		// 用户名称

		UserName *string `json:"UserName,omitempty" name:"UserName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAppInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSummaryRequest struct {
	*tchttp.BaseRequest

	// 概览种类，取值：APP_INSTANCE、COMPONENT_INSTANCE、APP

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 项目对应的命名空间

	Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces"`
}

func (r *GetSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetRolloutStrategyRequest struct {
	*tchttp.BaseRequest

	// 组件实例的ID号

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
	// 更新策略

	Strategy *Strategy `json:"Strategy,omitempty" name:"Strategy"`
}

func (r *SetRolloutStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetRolloutStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListScalerEventsRequest struct {
	*tchttp.BaseRequest

	// 弹性伸缩策略所归属的应用实例的ID

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
}

func (r *ListScalerEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListScalerEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAppExistRequest struct {
	*tchttp.BaseRequest

	// app 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// app 版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// 修改对象时的Id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *CheckAppExistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAppExistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppInstanceFromYamlRequest struct {
	*tchttp.BaseRequest

	// 用户uin

	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 集群标识

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// App版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// 应用模板Id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// yaml内容

	YamlList []*string `json:"YamlList,omitempty" name:"YamlList"`
}

func (r *CreateAppInstanceFromYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppInstanceFromYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppFilesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// helm chart package中的文件内容

		Files []*string `json:"Files,omitempty" name:"Files"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAppFilesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppYamlRequest struct {
	*tchttp.BaseRequest

	// 应用模板Id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetAppYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Yaml列表

		YamlList []*string `json:"YamlList,omitempty" name:"YamlList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAppYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCanaryPartitionRequest struct {
	*tchttp.BaseRequest

	// 组件实例的ID号

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
}

func (r *GetCanaryPartitionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCanaryPartitionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 主机列表

		Hosts []*Host `json:"Hosts,omitempty" name:"Hosts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Component struct {
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 负载类型

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 应用Id

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// trait列表

	Traits []*Trait `json:"Traits,omitempty" name:"Traits"`
	// 组件参数

	Parameters []*ComponentParameter `json:"Parameters,omitempty" name:"Parameters"`
	// 组件类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 配置

	Configs []*Config `json:"Configs,omitempty" name:"Configs"`
	// 输出

	Outputs []*Output `json:"Outputs,omitempty" name:"Outputs"`
	// 输入

	Inputs []*Input `json:"Inputs,omitempty" name:"Inputs"`
}

type Inputs struct {
	// 输入实例

	InputInstanceId *string `json:"InputInstanceId,omitempty" name:"InputInstanceId"`
	// 输出名称

	OutputName *string `json:"OutputName,omitempty" name:"OutputName"`
	// 文件路径

	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
	// 必须否

	Required *bool `json:"Required,omitempty" name:"Required"`
	// 实例id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 无

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 实例名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type Strategy struct {
	// Partition

	Partition *int64 `json:"Partition,omitempty" name:"Partition"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 最大pod数

	MaxSurge *string `json:"MaxSurge,omitempty" name:"MaxSurge"`
	// 最大不可用pod数

	MaxUnavailable *string `json:"MaxUnavailable,omitempty" name:"MaxUnavailable"`
	// 灰度和蓝绿策略时访问旧版本(也是业务暴露)的Service名字

	ActiveService *string `json:"ActiveService,omitempty" name:"ActiveService"`
	// 灰度和蓝绿策略时新版本暴露的Service名字

	PreviewService *string `json:"PreviewService,omitempty" name:"PreviewService"`
}

type GetAppInstanceComponentDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件的yaml内容

		Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAppInstanceComponentDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppInstanceComponentDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPaaSSummaryRequest struct {
	*tchttp.BaseRequest

	// 无

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 无

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetPaaSSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPaaSSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishAppRequest struct {
	*tchttp.BaseRequest

	// 模板Id列表

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
	// 是否发布上线

	Online *bool `json:"Online,omitempty" name:"Online"`
}

func (r *PublishAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSecretRequest struct {
	*tchttp.BaseRequest

	// 待更新的secret id

	Id *string `json:"Id,omitempty" name:"Id"`
	// secret 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// project 名字

	Project *string `json:"Project,omitempty" name:"Project"`
	// secret 所在 Namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// secret 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// secret内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *UpdateSecretRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSecretRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigMapItem struct {
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type RedeployAppInstanceRequest struct {
	*tchttp.BaseRequest

	// UserUin

	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 模版id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *RedeployAppInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RedeployAppInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例Id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppInstanceByChartValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例ID

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppInstanceByChartValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppInstanceByChartValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppInstanceComponentsRequest struct {
	*tchttp.BaseRequest

	// AppInstance的id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetAppInstanceComponentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppInstanceComponentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCanaryTrafficRequest struct {
	*tchttp.BaseRequest

	// 组件实例的ID号

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
}

func (r *GetCanaryTrafficRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCanaryTrafficRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetConfigMapYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Yaml信息

		Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConfigMapYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetConfigMapYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PromoteRolloutResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PromoteRolloutResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PromoteRolloutResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAppPicFileRequest struct {
	*tchttp.BaseRequest

	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// PicName

	PicName *string `json:"PicName,omitempty" name:"PicName"`
	// PicContent的base64字符串

	PicContent *string `json:"PicContent,omitempty" name:"PicContent"`
}

func (r *UpdateAppPicFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAppPicFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateScalerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// scaler id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateScalerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetConfigMapYamlRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetConfigMapYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetConfigMapYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAppExistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否存在

		Exist *bool `json:"Exist,omitempty" name:"Exist"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAppExistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAppExistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppInstanceIngressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// trait 实例的id

		TraitInstanceId *string `json:"TraitInstanceId,omitempty" name:"TraitInstanceId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppInstanceIngressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppInstanceIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppPreCheckToYamlRequest struct {
	*tchttp.BaseRequest

	// 应用模版id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 应用名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 无

	Description *string `json:"Description,omitempty" name:"Description"`
	// 无

	Components []*Component `json:"Components,omitempty" name:"Components"`
	// 无

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
}

func (r *CreateAppPreCheckToYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppPreCheckToYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// app实例列表

		AppInstances []*AppInstance `json:"AppInstances,omitempty" name:"AppInstances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAppInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpaScalerRequest struct {
	*tchttp.BaseRequest

	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteVpaScalerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpaScalerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDetailVpaScalerByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Data

		Data *VpaScaler `json:"Data,omitempty" name:"Data"`
		// 预估资源在此对象

		VpaK8sJson *string `json:"VpaK8sJson,omitempty" name:"VpaK8sJson"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDetailVpaScalerByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDetailVpaScalerByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTadPkgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 包名

		PkgName *string `json:"PkgName,omitempty" name:"PkgName"`
		// 包内容的二进制的base64字符串

		PkgContent *string `json:"PkgContent,omitempty" name:"PkgContent"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTadPkgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTadPkgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Input struct {
	// 输出参数名

	OutputName *string `json:"OutputName,omitempty" name:"OutputName"`
	// 路径

	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
	// 服务Id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务查询范围

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 是否必须

	Required *bool `json:"Required,omitempty" name:"Required"`
	// ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 组件ID

	ComponentId *string `json:"ComponentId,omitempty" name:"ComponentId"`
	// trait ID

	TraitId *string `json:"TraitId,omitempty" name:"TraitId"`
}

type GetBriefServiceInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// component instance 名字

		Name *string `json:"Name,omitempty" name:"Name"`
		// Component instance类型

		Type *string `json:"Type,omitempty" name:"Type"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBriefServiceInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBriefServiceInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRevisionsRequest struct {
	*tchttp.BaseRequest

	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 无

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 无

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 无

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 无

	Sort []*Sort `json:"Sort,omitempty" name:"Sort"`
	// 无

	IsRunning *bool `json:"IsRunning,omitempty" name:"IsRunning"`
}

func (r *ListRevisionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRevisionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAppPicFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAppPicFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAppPicFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationEventsRequest struct {
	*tchttp.BaseRequest

	// 要查询的命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 要查询的命名空间所在的集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetApplicationEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDimensionsRequest struct {
	*tchttp.BaseRequest
}

func (r *GetDimensionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDimensionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListHostsRequest struct {
	*tchttp.BaseRequest

	// 最大值

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 组件实例Id

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
}

func (r *ListHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartRolloutRequest struct {
	*tchttp.BaseRequest

	// 组件实例的ID

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
	// 需要更新的容器和image的列表

	UpdateImages []*UpdateImage `json:"UpdateImages,omitempty" name:"UpdateImages"`
}

func (r *StartRolloutRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartRolloutRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpaScaler struct {
	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	Name *string `json:"Name,omitempty" name:"Name"`
	// 无

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
	// 无

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
	// 无

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 无

	ResourcePolicy *string `json:"ResourcePolicy,omitempty" name:"ResourcePolicy"`
	// 无

	UpdatePolicy *string `json:"UpdatePolicy,omitempty" name:"UpdatePolicy"`
	// 无

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

type CreateAppInstanceFromYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例Id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppInstanceFromYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppInstanceFromYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAvailableStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAvailableStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAvailableStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBriefComponentInstanceRequest struct {
	*tchttp.BaseRequest

	// tad应用实例ID列表

	TadIds []*string `json:"TadIds,omitempty" name:"TadIds"`
	// helm应用实例ID列表

	HelmIds []*string `json:"HelmIds,omitempty" name:"HelmIds"`
}

func (r *GetBriefComponentInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBriefComponentInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetScalerYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 弹性伸缩策略的yaml字符串

		Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetScalerYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetScalerYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppPreCheckToYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Yaml列表

		YamlList []*string `json:"YamlList,omitempty" name:"YamlList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppPreCheckToYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppPreCheckToYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustomerAppInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例Id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCustomerAppInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomerAppInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteScalerRequest struct {
	*tchttp.BaseRequest

	// 弹性伸缩策略id

	ScalerId *string `json:"ScalerId,omitempty" name:"ScalerId"`
}

func (r *DeleteScalerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteScalerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppFileContentRequest struct {
	*tchttp.BaseRequest

	// 应用模板名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 应用模板版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// 应用模板中的具体文件名字

	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

func (r *GetAppFileContentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppFileContentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetConfigMapResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConfigMapResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetConfigMapResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RedeployAppInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RedeployAppInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RedeployAppInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveConfigMapYamlRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
}

func (r *SaveConfigMapYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveConfigMapYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchScalerStatusRequest struct {
	*tchttp.BaseRequest

	// 弹性伸缩策略id

	ScalerId *string `json:"ScalerId,omitempty" name:"ScalerId"`
}

func (r *SwitchScalerStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchScalerStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Scaler struct {
	// 弹性伸缩策略id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 弹性伸缩策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 弹性伸缩策略关联应用实例的ID

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
	// 弹性伸缩策略关联组件的ID

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
	// 弹性伸缩策略关联组件的名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 触发方式

	Type *string `json:"Type,omitempty" name:"Type"`
	// 触发策略

	TriggerStrategy *string `json:"TriggerStrategy,omitempty" name:"TriggerStrategy"`
	// 冷却时间

	CoolDown *int64 `json:"CoolDown,omitempty" name:"CoolDown"`
	// 是否启用

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

type CreateAppFromChartPkgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模版id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppFromChartPkgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppFromChartPkgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAppRequest struct {
	*tchttp.BaseRequest

	// 待删除应用模板Id列表

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppInsProductsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品列表

		Products []*string `json:"Products,omitempty" name:"Products"`
		// 产品列表数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAppInsProductsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppInsProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListBriefAppsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// App列表

		Apps []*App `json:"Apps,omitempty" name:"Apps"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListBriefAppsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListBriefAppsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScaleInstanceReplicasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScaleInstanceReplicasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScaleInstanceReplicasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Versions struct {
	// app版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// 组件数量

	ComponentNum *int64 `json:"ComponentNum,omitempty" name:"ComponentNum"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// ProductName

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 发布时间

	PublishedAt *string `json:"PublishedAt,omitempty" name:"PublishedAt"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type CreateAppRequest struct {
	*tchttp.BaseRequest

	// 应用名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 应用组件信息

	Components []*ComponentInput `json:"Components,omitempty" name:"Components"`
	// 应用的关键字

	KW *string `json:"KW,omitempty" name:"KW"`
	// helm chart的仓库地址

	Repo *string `json:"Repo,omitempty" name:"Repo"`
	// Config helm chart values.yaml

	Config *string `json:"Config,omitempty" name:"Config"`
	// helm chart logo base64 content

	PicFile *string `json:"PicFile,omitempty" name:"PicFile"`
	// helm chart package base64 content

	AppFile *string `json:"AppFile,omitempty" name:"AppFile"`
	// helm chart logo name

	PicName *string `json:"PicName,omitempty" name:"PicName"`
	// 应用模板类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 应用类别

	Category *string `json:"Category,omitempty" name:"Category"`
}

func (r *CreateAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAppInstanceRequest struct {
	*tchttp.BaseRequest

	// app实例id列表

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteAppInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAppInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAppInstanceIngressRequest struct {
	*tchttp.BaseRequest

	// ingress trait实例ID

	TraitInstanceId *string `json:"TraitInstanceId,omitempty" name:"TraitInstanceId"`
}

func (r *DeleteAppInstanceIngressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAppInstanceIngressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 名称

		Name *string `json:"Name,omitempty" name:"Name"`
		// 产品名称

		ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
		// 状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 版本

		AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
		// 描述信息

		Description *string `json:"Description,omitempty" name:"Description"`
		// 组件数

		ComponentNum *int64 `json:"ComponentNum,omitempty" name:"ComponentNum"`
		// 创建时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 发布时间

		PublishedAt *string `json:"PublishedAt,omitempty" name:"PublishedAt"`
		// 组件详情

		Components []*Component `json:"Components,omitempty" name:"Components"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppInstanceNamespacesRequest struct {
	*tchttp.BaseRequest

	// 模板ID列表

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *ListAppInstanceNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppInstanceNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Host struct {
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 资产Id

	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
	// SN

	SN *string `json:"SN,omitempty" name:"SN"`
	// 设备类型

	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
	// 主机类型:0 physical， 1 vm

	HostType *int64 `json:"HostType,omitempty" name:"HostType"`
	// 内部Ip

	InternalIp *string `json:"InternalIp,omitempty" name:"InternalIp"`
	// host IP

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 上联设备

	UplinkDevice *string `json:"UplinkDevice,omitempty" name:"UplinkDevice"`
	// 机架名

	RackName *string `json:"RackName,omitempty" name:"RackName"`
	// AZ

	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
}

type RevisionInput struct {
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	OutputName *string `json:"OutputName,omitempty" name:"OutputName"`
	// 路径

	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 范围

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 必要性

	Required *bool `json:"Required,omitempty" name:"Required"`
	// 组件id

	RevisionComponentId *string `json:"RevisionComponentId,omitempty" name:"RevisionComponentId"`
	// trait id

	RevisionTraitId *string `json:"RevisionTraitId,omitempty" name:"RevisionTraitId"`
	// input实例id

	InputInstanceId *string `json:"InputInstanceId,omitempty" name:"InputInstanceId"`
}

type RevisionConfig struct {
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件id

	RevisionComponentId *string `json:"RevisionComponentId,omitempty" name:"RevisionComponentId"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 内容

	RawContent *string `json:"RawContent,omitempty" name:"RawContent"`
	// 配置实例id

	ConfigInstanceId *string `json:"ConfigInstanceId,omitempty" name:"ConfigInstanceId"`
	// revision 配置列表

	Inputs []*RevisionInput `json:"Inputs,omitempty" name:"Inputs"`
	// 配置KV

	KVs []*ConfigKV `json:"KVs,omitempty" name:"KVs"`
}

type ListConfigMapsRequest struct {
	*tchttp.BaseRequest

	// 一页最大数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移的记录条数

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *ListConfigMapsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListConfigMapsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetRolloutStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetRolloutStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetRolloutStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateScalerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 已更新的弹性伸缩策略的id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateScalerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateScalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComponentInstanceInput struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 关联的组件Id

	ComponentId *string `json:"ComponentId,omitempty" name:"ComponentId"`
	// trait实例信息

	TraitInstances []*TraitInstance `json:"TraitInstances,omitempty" name:"TraitInstances"`
	// 组件参数输入信息

	Parameters []*ParameterInput `json:"Parameters,omitempty" name:"Parameters"`
	// 配置

	Configs []*ConfigInstance `json:"Configs,omitempty" name:"Configs"`
	// 输出

	Outputs []*OutputInstance `json:"Outputs,omitempty" name:"Outputs"`
	// 输入

	Inputs []*InputInstance `json:"Inputs,omitempty" name:"Inputs"`
}

type Containers struct {
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 镜像

	Image *string `json:"Image,omitempty" name:"Image"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type Event struct {
	// 更新时间

	LastTimestamp *string `json:"LastTimestamp,omitempty" name:"LastTimestamp"`
	// 时间原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 时间信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type CreateAppInstanceByChartValuesRequest struct {
	*tchttp.BaseRequest

	// 集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 模板版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// 模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 模板id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// yaml形式的values字符串

	Values *string `json:"Values,omitempty" name:"Values"`
}

func (r *CreateAppInstanceByChartValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppInstanceByChartValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDimensionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 打散维度

		Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDimensionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDimensionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRolloutStatusRequest struct {
	*tchttp.BaseRequest

	// 组件实例的ID

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
}

func (r *GetRolloutStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRolloutStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Revisions

		Revisions []*RevisionVersion `json:"Revisions,omitempty" name:"Revisions"`
		// Total

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAppVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveRevisionRequest struct {
	*tchttp.BaseRequest

	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// uin

	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 产品

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 实例id

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
	// revsion

	Revision *string `json:"Revision,omitempty" name:"Revision"`
	// 组件列表

	RevisionComponents []*RevisionComponent `json:"RevisionComponents,omitempty" name:"RevisionComponents"`
}

func (r *SaveRevisionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveRevisionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCanaryTrafficRequest struct {
	*tchttp.BaseRequest

	// 组件实例的ID号

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
	// 设定的流量参数

	TrafficInfo *TrafficInfo `json:"TrafficInfo,omitempty" name:"TrafficInfo"`
}

func (r *SetCanaryTrafficRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCanaryTrafficRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppsFromTadProductPkgRequest struct {
	*tchttp.BaseRequest

	// 文件名

	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
	// 文件名 版本

	TemplateVersion *string `json:"TemplateVersion,omitempty" name:"TemplateVersion"`
	// 默认为 tad

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
	// 产品包为product

	PkgType *string `json:"PkgType,omitempty" name:"PkgType"`
	// 包的base64

	PkgContent *string `json:"PkgContent,omitempty" name:"PkgContent"`
}

func (r *CreateAppsFromTadProductPkgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppsFromTadProductPkgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// app列表

		Apps []*App `json:"Apps,omitempty" name:"Apps"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAppsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCanaryPartitionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCanaryPartitionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCanaryPartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Pod struct {
	// Age

	Age *string `json:"Age,omitempty" name:"Age"`
	// Containers

	Containers []*Containers `json:"Containers,omitempty" name:"Containers"`
	// HostIp

	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// PodIp

	PodIp *string `json:"PodIp,omitempty" name:"PodIp"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// Pod所属命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type TrafficInfo struct {
	// Nginx配置

	NginxInfo *NginxInfo `json:"NginxInfo,omitempty" name:"NginxInfo"`
}

type GetPaaSSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPaaSSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPaaSSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppProductsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品列表数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 产品列表

		Products []*string `json:"Products,omitempty" name:"Products"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAppProductsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComponentInstanceParameter struct {
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 对应模板路径

	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 是否必填

	Required *bool `json:"Required,omitempty" name:"Required"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 组件实例Id

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
}

type ComponentParameter struct {
	// 参数Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数对应模板路径

	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 是否必填

	Required *bool `json:"Required,omitempty" name:"Required"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 组件Id

	ComponentId *string `json:"ComponentId,omitempty" name:"ComponentId"`
}

type CheckAppInstanceExistRequest struct {
	*tchttp.BaseRequest

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群id

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 修改对象时的Id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *CheckAppInstanceExistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAppInstanceExistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppInsPreCheckToYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// yaml列表

		YamlList []*string `json:"YamlList,omitempty" name:"YamlList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppInsPreCheckToYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppInsPreCheckToYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppInstanceComponentDetailRequest struct {
	*tchttp.BaseRequest

	// 组件的名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件的命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 组件的类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 应用实例组件所属集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 组件的gv

	GroupVersion *string `json:"GroupVersion,omitempty" name:"GroupVersion"`
}

func (r *GetAppInstanceComponentDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppInstanceComponentDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetChartPkgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件名（包名）

		PkgName *string `json:"PkgName,omitempty" name:"PkgName"`
		// 文件的base64字符串

		PkgContent *string `json:"PkgContent,omitempty" name:"PkgContent"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetChartPkgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetChartPkgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDetailVpaScalerByIdRequest struct {
	*tchttp.BaseRequest

	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetDetailVpaScalerByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDetailVpaScalerByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetUpgradeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetUpgradeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetUpgradeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAppRequest struct {
	*tchttp.BaseRequest

	// 应用模板Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 应用名称，模糊过滤

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 无

	Description *string `json:"Description,omitempty" name:"Description"`
	// 无

	Components []*Component `json:"Components,omitempty" name:"Components"`
	// 无

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
}

func (r *UpdateAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppResp struct {
	// 模版id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 错误信息

	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
	// 模版名

	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

type CreateAppFromYamlRequest struct {
	*tchttp.BaseRequest

	// yaml列表

	YamlList []*string `json:"YamlList,omitempty" name:"YamlList"`
	// 分类

	Category *string `json:"Category,omitempty" name:"Category"`
}

func (r *CreateAppFromYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppFromYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppInstanceYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// yaml信息列表

		YamlList []*string `json:"YamlList,omitempty" name:"YamlList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAppInstanceYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppInstanceYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAvailableStrategyRequest struct {
	*tchttp.BaseRequest

	// workload的版本号

	ApiVersion *string `json:"ApiVersion,omitempty" name:"ApiVersion"`
	// workload的类型名称

	Kind *string `json:"Kind,omitempty" name:"Kind"`
}

func (r *GetAvailableStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAvailableStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInstanceReplicasRequest struct {
	*tchttp.BaseRequest

	// ComponentInstanceId

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
}

func (r *GetInstanceReplicasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInstanceReplicasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用模板Id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAppInstanceIngressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAppInstanceIngressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAppInstanceIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRevisionRequest struct {
	*tchttp.BaseRequest

	// 版本id

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteRevisionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRevisionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRolloutStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRolloutStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRolloutStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppInstancesRequest struct {
	*tchttp.BaseRequest

	// 最大查询数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// app名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 排序规则

	Sort []*Sort `json:"Sort,omitempty" name:"Sort"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 应用状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ListAppInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NginxInfo struct {
	// 稳定版本暴露服务的Service Name

	ServiceStable *string `json:"ServiceStable,omitempty" name:"ServiceStable"`
	// 新版本灰度暴露服务的Service Name

	ServicePreview *string `json:"ServicePreview,omitempty" name:"ServicePreview"`
	// 稳定版本暴露服务的Ingress Name

	IngressStable *string `json:"IngressStable,omitempty" name:"IngressStable"`
	// 新版本灰度暴露服务的Ingress Name

	IngressPreview *string `json:"IngressPreview,omitempty" name:"IngressPreview"`
	// 灰度版本的权值0-100

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// 通过cookie进行A/B testing，http请求带上CanaryCookie: always的Cookie，流量即将流入灰度版本

	CanaryCookie *string `json:"CanaryCookie,omitempty" name:"CanaryCookie"`
	// 通过header进行A/B testing，header的key

	CanaryHeader *string `json:"CanaryHeader,omitempty" name:"CanaryHeader"`
	// Header的value满足正则表达式即流量导入到灰度版本

	HeaderPattern *string `json:"HeaderPattern,omitempty" name:"HeaderPattern"`
	// Header的value等于ByValue即流量流向灰度版本

	HeaderValue *string `json:"HeaderValue,omitempty" name:"HeaderValue"`
}

type TraitInput struct {
	// trait名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// trait参数值列表

	Parameters []*ParameterInput `json:"Parameters,omitempty" name:"Parameters"`
	// 输出参数列表

	Outputs []*OutputReq `json:"Outputs,omitempty" name:"Outputs"`
	// 输入参数列表

	Inputs []*Input `json:"Inputs,omitempty" name:"Inputs"`
}

type GetTadPkgRequest struct {
	*tchttp.BaseRequest

	// 模版id

	AppID *string `json:"AppID,omitempty" name:"AppID"`
}

func (r *GetTadPkgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTadPkgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppsRequest struct {
	*tchttp.BaseRequest

	// 应用名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 最大记录数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序规则

	Sort []*Sort `json:"Sort,omitempty" name:"Sort"`
	// tad or helm

	Type *string `json:"Type,omitempty" name:"Type"`
	// 应用类别

	Category *string `json:"Category,omitempty" name:"Category"`
}

func (r *ListAppsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppInstance struct {
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 用户uin

	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`
	// 集群标识

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// ready的组件数目

	ReadyComponent *int64 `json:"ReadyComponent,omitempty" name:"ReadyComponent"`
	// 组件总数

	TotalComponent *int64 `json:"TotalComponent,omitempty" name:"TotalComponent"`
	// 版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// app模板id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 用户名称

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 应用部署空间

	Scope *string `json:"Scope,omitempty" name:"Scope"`
}

type CreateSecretResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// secret主键

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecretResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetConfigMapRequest struct {
	*tchttp.BaseRequest

	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetConfigMapRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetConfigMapRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Sort struct {
	// 排序字段

	SortField *string `json:"SortField,omitempty" name:"SortField"`
	// 升序：asc  降序：desc

	Order *string `json:"Order,omitempty" name:"Order"`
}

type CreateVpaScalerRequest struct {
	*tchttp.BaseRequest

	// 无

	Data *VpaScaler `json:"Data,omitempty" name:"Data"`
}

func (r *CreateVpaScalerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpaScalerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppConfigRequest struct {
	*tchttp.BaseRequest

	// 应用模板的名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 应用模板的版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
}

func (r *GetAppConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件列表

		EventList []*Event `json:"EventList,omitempty" name:"EventList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkloadPodsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Pod列表的详细内容

		Pods []*Pod `json:"Pods,omitempty" name:"Pods"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListWorkloadPodsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListWorkloadPodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComponentInstance struct {
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 负载类型

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// app实例Id

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
	// 关联的组件id

	ComponentId *string `json:"ComponentId,omitempty" name:"ComponentId"`
	// trait实例信息

	TraitInstances []*TraitInstance `json:"TraitInstances,omitempty" name:"TraitInstances"`
	// 组件实例参数

	Parameters []*ComponentInstanceParameter `json:"Parameters,omitempty" name:"Parameters"`
	// 组件配置实例

	Configs []*ConfigInstance `json:"Configs,omitempty" name:"Configs"`
	// 组件标记

	Mark *string `json:"Mark,omitempty" name:"Mark"`
	// input参数列表

	Inputs []*InputInstance `json:"Inputs,omitempty" name:"Inputs"`
	// output参数列表

	Outputs []*OutputInstance `json:"Outputs,omitempty" name:"Outputs"`
	// 组件类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
}

type TraitInstanceInput struct {
	// trait实例名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 关联的trait id

	TraitId *string `json:"TraitId,omitempty" name:"TraitId"`
	// 参数值

	Parameters []*ParameterInput `json:"Parameters,omitempty" name:"Parameters"`
}

type CreateConfigMapsRequest struct {
	*tchttp.BaseRequest

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// KV 数据

	KVs []*ConfigKV `json:"KVs,omitempty" name:"KVs"`
	// 集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *CreateConfigMapsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigMapsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRevisionYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// yaml列表

		YamlList []*string `json:"YamlList,omitempty" name:"YamlList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRevisionYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRevisionYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPodsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Pods

		Pods []*Pod `json:"Pods,omitempty" name:"Pods"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPodsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListVpaScalerRequest struct {
	*tchttp.BaseRequest

	// AppInstanceId

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// Sort

	Sort []*Sort `json:"Sort,omitempty" name:"Sort"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ListVpaScalerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListVpaScalerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCanaryTrafficResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCanaryTrafficResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCanaryTrafficResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OutputInstance struct {
	// ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// output Id

	OutputId *string `json:"OutputId,omitempty" name:"OutputId"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 目标值路径

	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
	// 条件

	Conditions *string `json:"Conditions,omitempty" name:"Conditions"`
	// 服务Id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 作用范围

	Level *string `json:"Level,omitempty" name:"Level"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 主机

	Host *string `json:"Host,omitempty" name:"Host"`
	// uri path

	Path *string `json:"Path,omitempty" name:"Path"`
	// AZ

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 是否underlay网络

	Underlay *bool `json:"Underlay,omitempty" name:"Underlay"`
	// 所属组件实例id

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
	// 所属trait实例id，仅在trait下存在

	TraitInstanceId *string `json:"TraitInstanceId,omitempty" name:"TraitInstanceId"`
	// 无

	Portal *bool `json:"Portal,omitempty" name:"Portal"`
	// 无

	Redirect *bool `json:"Redirect,omitempty" name:"Redirect"`
}

type UpdateSecretResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// secret id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSecretResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditVpaScalerRequest struct {
	*tchttp.BaseRequest

	// EditData

	EditData *VpaScaler `json:"EditData,omitempty" name:"EditData"`
}

func (r *EditVpaScalerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditVpaScalerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNewestRevisionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新版本

		NewestRevision *string `json:"NewestRevision,omitempty" name:"NewestRevision"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetNewestRevisionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNewestRevisionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPodsRequest struct {
	*tchttp.BaseRequest

	// 组件实例id

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
}

func (r *ListPodsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPodsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Config struct {
	// 配置名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 内容

	RawContent *string `json:"RawContent,omitempty" name:"RawContent"`
	// ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 组件Id

	ComponentId *string `json:"ComponentId,omitempty" name:"ComponentId"`
	// 依赖输入

	Inputs []*Input `json:"Inputs,omitempty" name:"Inputs"`
	// 配置的KV

	KVs []*ConfigKV `json:"KVs,omitempty" name:"KVs"`
}

type GetApplicationComponentsSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationComponentsSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationComponentsSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationSummaryRequest struct {
	*tchttp.BaseRequest

	// 项目名（k8s里的namespace）

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetApplicationSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListNamespacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询命名空间列表

		Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RevisionOutput struct {
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 路径

	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
	// 条件

	Conditions *string `json:"Conditions,omitempty" name:"Conditions"`
	// 服务id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 范围

	Level *string `json:"Level,omitempty" name:"Level"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 主机

	Host *string `json:"Host,omitempty" name:"Host"`
	// uri路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// AZ

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 是否underlay

	Underlay *bool `json:"Underlay,omitempty" name:"Underlay"`
	// 组件id

	RevisionComponentId *string `json:"RevisionComponentId,omitempty" name:"RevisionComponentId"`
	// trait id

	RevisionTraitId *string `json:"RevisionTraitId,omitempty" name:"RevisionTraitId"`
	// 输出实例id

	OutputInstanceId *string `json:"OutputInstanceId,omitempty" name:"OutputInstanceId"`
	// 无

	Portal *bool `json:"Portal,omitempty" name:"Portal"`
	// 无

	Redirect *bool `json:"Redirect,omitempty" name:"Redirect"`
}

type ApplyRevisionRequest struct {
	*tchttp.BaseRequest

	// UserUin

	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`
	// UserName

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// AppInstanceId

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
	// RevisionId

	RevisionId *string `json:"RevisionId,omitempty" name:"RevisionId"`
}

func (r *ApplyRevisionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyRevisionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCanaryTrafficResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCanaryTrafficResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCanaryTrafficResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductTgzRequest struct {
	*tchttp.BaseRequest

	// 产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 应用ID列表

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *GetProductTgzRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductTgzRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCanaryPartitionRequest struct {
	*tchttp.BaseRequest

	// 组件实例的ID号

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
	// 保留旧版本pod的数量

	Partition *int64 `json:"Partition,omitempty" name:"Partition"`
}

func (r *SetCanaryPartitionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCanaryPartitionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartRolloutResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartRolloutResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartRolloutResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigMapsRequest struct {
	*tchttp.BaseRequest

	// configmap 的Id列表

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteConfigMapsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigMapsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditVpaScalerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		ActionResult *string `json:"ActionResult,omitempty" name:"ActionResult"`
		// 无

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditVpaScalerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditVpaScalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigMapsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigMapsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigMapsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用模板的配置信息

		Config *string `json:"Config,omitempty" name:"Config"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAppConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComponentInstanceRequest struct {
	*tchttp.BaseRequest

	// 组件实例ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetComponentInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComponentInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PromoteRolloutRequest struct {
	*tchttp.BaseRequest

	// 组件实例的ID号

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
}

func (r *PromoteRolloutRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PromoteRolloutRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TraitInstanceParameter struct {
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 对应模板路径

	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 是否必填

	Required *string `json:"Required,omitempty" name:"Required"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
	// trait实例Id

	TraitInstanceId *string `json:"TraitInstanceId,omitempty" name:"TraitInstanceId"`
}

type CreateAppInstanceRequest struct {
	*tchttp.BaseRequest

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 用户uin

	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`
	// 集群标志

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 关联的App模板Id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 组件实例信息

	ComponentInstances []*ComponentInstanceInput `json:"ComponentInstances,omitempty" name:"ComponentInstances"`
	// 用户名称

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// helm应用实例的配置信息

	Content *string `json:"Content,omitempty" name:"Content"`
	// helm应用实例复用某版本

	Reused *bool `json:"Reused,omitempty" name:"Reused"`
	// 应用类别

	Type *string `json:"Type,omitempty" name:"Type"`
	// 应用实例id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 默认都是ns级别部署的helm应用，只有传入Scope=cluster时应用为集群级别，可跨多个ns。

	Scope *string `json:"Scope,omitempty" name:"Scope"`
}

func (r *CreateAppInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppInstanceIngressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ingress trait实例的列表

		IngressList []*TraitInstance `json:"IngressList,omitempty" name:"IngressList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAppInstanceIngressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppInstanceIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KVMap struct {
	// 无

	Key *string `json:"Key,omitempty" name:"Key"`
	// 无

	Value *string `json:"Value,omitempty" name:"Value"`
}

type Output struct {
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 文件路径

	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
	// 查询条件

	Conditions *string `json:"Conditions,omitempty" name:"Conditions"`
	// 主机

	Host *string `json:"Host,omitempty" name:"Host"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 无

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 层级

	Level *string `json:"Level,omitempty" name:"Level"`
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 无

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 无

	Underlay *bool `json:"Underlay,omitempty" name:"Underlay"`
	// 组件id

	ComponentId *string `json:"ComponentId,omitempty" name:"ComponentId"`
	// trait id

	TraitId *string `json:"TraitId,omitempty" name:"TraitId"`
	// 无

	Redirect *bool `json:"Redirect,omitempty" name:"Redirect"`
	// 无

	Portal *bool `json:"Portal,omitempty" name:"Portal"`
}

type GetAppFileContentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Yaml文件的内容

		Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAppFileContentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppFileContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductTgzResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件名

		PkgName *string `json:"PkgName,omitempty" name:"PkgName"`
		// 文件base64

		PkgContent *string `json:"PkgContent,omitempty" name:"PkgContent"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductTgzResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductTgzResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRevisionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// build版本列表

		BuildVersions []*BuildVersion `json:"BuildVersions,omitempty" name:"BuildVersions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListRevisionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListRevisionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StateSummary struct {
	// 状态

	State *string `json:"State,omitempty" name:"State"`
	// 数量

	Num *int64 `json:"Num,omitempty" name:"Num"`
}

type PublishAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAppFromYamlRequest struct {
	*tchttp.BaseRequest

	// 应用模板Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// yaml信息列表

	YamlList []*string `json:"YamlList,omitempty" name:"YamlList"`
}

func (r *UpdateAppFromYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAppFromYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppFilesRequest struct {
	*tchttp.BaseRequest

	// 应用模板名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 应用模板的版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
}

func (r *GetAppFilesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppFilesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListScalersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 当前页策略列表

		Scalers []*Scaler `json:"Scalers,omitempty" name:"Scalers"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListScalersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListScalersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDescriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDescriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustomerAppInstanceRequest struct {
	*tchttp.BaseRequest

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 组件信息

	ComponentInstances []*ComponentInstance `json:"ComponentInstances,omitempty" name:"ComponentInstances"`
	// tad or helm

	Type *string `json:"Type,omitempty" name:"Type"`
	// 是否另存为模板

	IsSaved *bool `json:"IsSaved,omitempty" name:"IsSaved"`
}

func (r *CreateCustomerAppInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomerAppInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpaScalerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Id *string `json:"Id,omitempty" name:"Id"`
		// 无

		ActionResult *string `json:"ActionResult,omitempty" name:"ActionResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpaScalerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpaScalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListScalersRequest struct {
	*tchttp.BaseRequest

	// 应用实例的id

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
	// 上限

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序

	Sort []*Sort `json:"Sort,omitempty" name:"Sort"`
}

func (r *ListScalersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListScalersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Configs struct {
	// 配置 id

	ConfigInstanceId *string `json:"ConfigInstanceId,omitempty" name:"ConfigInstanceId"`
	// 内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

type OutputReq struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数路径

	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
	// 条件

	Conditions *string `json:"Conditions,omitempty" name:"Conditions"`
	// Service唯一标志

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 服务级别

	Level *string `json:"Level,omitempty" name:"Level"`
	// 端口

	Port *int64 `json:"Port,omitempty" name:"Port"`
	// 协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 主机名

	Host *string `json:"Host,omitempty" name:"Host"`
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// Zone标识

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 是否underlay

	Underlay *bool `json:"Underlay,omitempty" name:"Underlay"`
	// 无

	Portal *bool `json:"Portal,omitempty" name:"Portal"`
	// 无

	Redirect *bool `json:"Redirect,omitempty" name:"Redirect"`
}

type RevisionApp struct {
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 用户uin

	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 产品

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 组件数

	TotalComponent *int64 `json:"TotalComponent,omitempty" name:"TotalComponent"`
	// app 版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// revision

	Revision *string `json:"Revision,omitempty" name:"Revision"`
	// 是否运行

	IsRunning *bool `json:"IsRunning,omitempty" name:"IsRunning"`
	// app实例id

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
	// 组件列表

	RevisionComponents []*RevisionComponent `json:"RevisionComponents,omitempty" name:"RevisionComponents"`
}

type DeleteSecretRequest struct {
	*tchttp.BaseRequest

	// secret id

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteSecretRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecretRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRevisionYamlRequest struct {
	*tchttp.BaseRequest

	// revision的ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetRevisionYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRevisionYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListBriefAppsRequest struct {
	*tchttp.BaseRequest

	// 应用名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 最大记录数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序规则

	Sort *Sort `json:"Sort,omitempty" name:"Sort"`
	// 模板类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 按照分类进行查询，例如中间件、数据库等

	Category *string `json:"Category,omitempty" name:"Category"`
}

func (r *ListBriefAppsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListBriefAppsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollbackRequest struct {
	*tchttp.BaseRequest

	// 组件实例的ID号

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
}

func (r *RollbackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RollbackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveRevisionByComponentRequest struct {
	*tchttp.BaseRequest

	// 应用实例ID

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
	// 要保存的revision版本

	Revision *string `json:"Revision,omitempty" name:"Revision"`
	// revision组件信息

	RevisionComponent *RevisionComponent `json:"RevisionComponent,omitempty" name:"RevisionComponent"`
}

func (r *SaveRevisionByComponentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveRevisionByComponentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Trait struct {
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 组件Id

	ComponentId *string `json:"ComponentId,omitempty" name:"ComponentId"`
	// trait参数

	Parameters []*TraitParameter `json:"Parameters,omitempty" name:"Parameters"`
	// 输出

	Outputs []*Output `json:"Outputs,omitempty" name:"Outputs"`
	// 输入

	Inputs []*Input `json:"Inputs,omitempty" name:"Inputs"`
}

type CreateAppInsPreCheckToYamlRequest struct {
	*tchttp.BaseRequest

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 用户uin

	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`
	// 集群名

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 关联的App模板Id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 组件实例信息

	ComponentInstances []*ComponentInstanceInput `json:"ComponentInstances,omitempty" name:"ComponentInstances"`
	// 用户名称

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// helm应用实例的配置信息

	Content *string `json:"Content,omitempty" name:"Content"`
	// helm应用实例复用某版本

	Reused *bool `json:"Reused,omitempty" name:"Reused"`
	// 应用类别

	Type *string `json:"Type,omitempty" name:"Type"`
	// 应用实例id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *CreateAppInsPreCheckToYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppInsPreCheckToYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppCategoriesRequest struct {
	*tchttp.BaseRequest
}

func (r *ListAppCategoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppCategoriesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigKV struct {
	// 无

	Key *string `json:"Key,omitempty" name:"Key"`
	// 无

	Value *string `json:"Value,omitempty" name:"Value"`
}

type RevisionVersion struct {
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// revision

	Revision *string `json:"Revision,omitempty" name:"Revision"`
	// 生成者

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 注释

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 是否运行

	IsRunning *string `json:"IsRunning,omitempty" name:"IsRunning"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
}

type UpdateScalerRequest struct {
	*tchttp.BaseRequest

	// 待更新的弹性伸缩策略id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 弹性伸缩策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 弹性伸缩策略关联应用实例的ID

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
	// 弹性伸缩策略关联组件的ID

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
	// 弹性伸缩策略关联组件的名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 触发方式

	Type *string `json:"Type,omitempty" name:"Type"`
	// 触发策略（含实例范围)

	TriggerStrategy *string `json:"TriggerStrategy,omitempty" name:"TriggerStrategy"`
	// 冷却时间（分钟）

	CoolDown *int64 `json:"CoolDown,omitempty" name:"CoolDown"`
	// 是否启用，1为启用，0为禁用

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

func (r *UpdateScalerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateScalerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAppInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAppInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAppInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBriefServiceInstanceRequest struct {
	*tchttp.BaseRequest

	// Component Instance的id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetBriefServiceInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBriefServiceInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAppInstanceIngressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAppInstanceIngressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAppInstanceIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveRevisionFromYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveRevisionFromYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveRevisionFromYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigMapFromYamlRequest struct {
	*tchttp.BaseRequest

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群标识

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// Yaml格式的configmap

	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
}

func (r *CreateConfigMapFromYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigMapFromYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppInsProductsRequest struct {
	*tchttp.BaseRequest

	// Namespaces数组

	Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces"`
}

func (r *ListAppInsProductsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppInsProductsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListEventsRequest struct {
	*tchttp.BaseRequest

	// 应用实例id

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
	// 组件实例id

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
}

func (r *ListEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板列表

		Templates []*Template `json:"Templates,omitempty" name:"Templates"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetEventVpaScalerByIdRequest struct {
	*tchttp.BaseRequest

	// 无

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetEventVpaScalerByIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetEventVpaScalerByIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyRevisionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyRevisionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyRevisionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationComponentsSummaryRequest struct {
	*tchttp.BaseRequest

	// 无

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetApplicationComponentsSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationComponentsSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetChartPkgRequest struct {
	*tchttp.BaseRequest

	// 模版id

	AppID *string `json:"AppID,omitempty" name:"AppID"`
}

func (r *GetChartPkgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetChartPkgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// State概要列表

		States []*StateSummary `json:"States,omitempty" name:"States"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSecretsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总页数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// secret 列表

		Secrets []*Secret `json:"Secrets,omitempty" name:"Secrets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSecretsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSecretsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type App struct {
	// App的唯一标识

	Id *string `json:"Id,omitempty" name:"Id"`
	// 应用名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 模板状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 版本号

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 包含的组件总数

	ComponentNum *int64 `json:"ComponentNum,omitempty" name:"ComponentNum"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 发布时间

	PublishedAt *string `json:"PublishedAt,omitempty" name:"PublishedAt"`
	// 模板所属的分类，例如中间件、数据库等

	Category *string `json:"Category,omitempty" name:"Category"`
}

type RevisionTraitParameter struct {
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 路径

	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 是否必要

	Required *bool `json:"Required,omitempty" name:"Required"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 关联的trait Id

	RevisionTraitId *string `json:"RevisionTraitId,omitempty" name:"RevisionTraitId"`
	// 参数实例Id

	ParameterInstanceId *string `json:"ParameterInstanceId,omitempty" name:"ParameterInstanceId"`
}

type GetUpgradeConfigRequest struct {
	*tchttp.BaseRequest

	// 组件实例id

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
}

func (r *GetUpgradeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppInstanceIngressRequest struct {
	*tchttp.BaseRequest

	// 应用实例ID

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
	// ingress trait的信息

	Ingress *TraitInstance `json:"Ingress,omitempty" name:"Ingress"`
}

func (r *CreateAppInstanceIngressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppInstanceIngressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteScalerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 弹性伸缩策略id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteScalerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteScalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateConfigMapsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateConfigMapsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateConfigMapsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Template struct {
	// 模板Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 模板种类

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 模板类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 模板版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 参数列表

	Parameters []*Parameter `json:"Parameters,omitempty" name:"Parameters"`
}

type DeleteSecretResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// secret id

		Ids []*string `json:"Ids,omitempty" name:"Ids"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecretResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBriefComponentInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBriefComponentInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBriefComponentInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScaleInstanceReplicasRequest struct {
	*tchttp.BaseRequest

	// 组件实例id

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
	// 扩缩容的副本数

	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`
}

func (r *ScaleInstanceReplicasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScaleInstanceReplicasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDescriptionRequest struct {
	*tchttp.BaseRequest

	// 类型，

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 对象ID

	ObjectId *string `json:"ObjectId,omitempty" name:"ObjectId"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateDescriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDescriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetAppInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppInstanceIngressRequest struct {
	*tchttp.BaseRequest

	// 应用实例ID

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
}

func (r *GetAppInstanceIngressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppInstanceIngressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetEventVpaScalerByIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Total

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Events

		Events []*Event `json:"Events,omitempty" name:"Events"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetEventVpaScalerByIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetEventVpaScalerByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveRevisionByComponentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// revision版本id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveRevisionByComponentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveRevisionByComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RevisionComponent struct {
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 负载

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// app id

	RevisionAppId *string `json:"RevisionAppId,omitempty" name:"RevisionAppId"`
	// 组件实例id

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
	// trait列表

	RevisionTraits []*RevisionTrait `json:"RevisionTraits,omitempty" name:"RevisionTraits"`
	// 参数列表

	Parameters []*RevisionComponentParameter `json:"Parameters,omitempty" name:"Parameters"`
	// config列表

	Configs []*RevisionConfig `json:"Configs,omitempty" name:"Configs"`
	// input列表

	Inputs []*RevisionInput `json:"Inputs,omitempty" name:"Inputs"`
	// output列表

	Outputs []*RevisionOutput `json:"Outputs,omitempty" name:"Outputs"`
}

type Secret struct {
	// secret Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// secret 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// secret 所在project

	Project *string `json:"Project,omitempty" name:"Project"`
	// secret 所在ns

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// secret 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// secret 具体内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type CreateAppFromTadPkgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用模版id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppFromTadPkgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppFromTadPkgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterSummaryRequest struct {
	*tchttp.BaseRequest
}

func (r *GetClusterSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Replicas

		Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`
		// Replicas

		Strategy *Strategy `json:"Strategy,omitempty" name:"Strategy"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUpgradeConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAppInstanceIngressRequest struct {
	*tchttp.BaseRequest

	// appinstanceid

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
	// trait instance

	Ingress *TraitInstance `json:"Ingress,omitempty" name:"Ingress"`
}

func (r *UpdateAppInstanceIngressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAppInstanceIngressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppInstanceNamespacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 基于某模板部署的所有应用实例所在的命名空间列表

		Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAppInstanceNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppInstanceNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAppInstanceExistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否存在

		Exist *bool `json:"Exist,omitempty" name:"Exist"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAppInstanceExistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAppInstanceExistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRevisionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRevisionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRevisionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpaScalerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ActionResult

		ActionResult *string `json:"ActionResult,omitempty" name:"ActionResult"`
		// Id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpaScalerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpaScalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationResourceUsageSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationResourceUsageSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationResourceUsageSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchScalerStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 弹性伸缩策略id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchScalerStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchScalerStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateConfigMapsRequest struct {
	*tchttp.BaseRequest

	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// KV数组

	KVs []*ConfigKV `json:"KVs,omitempty" name:"KVs"`
}

func (r *UpdateConfigMapsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateConfigMapsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RevisionComponentParameter struct {
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 路径

	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 是否必要

	Required *bool `json:"Required,omitempty" name:"Required"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 组件id

	RevisionComponentId *string `json:"RevisionComponentId,omitempty" name:"RevisionComponentId"`
	// 参数实例id

	ParameterInstanceId *string `json:"ParameterInstanceId,omitempty" name:"ParameterInstanceId"`
}

type CreateAppFromChartPkgRequest struct {
	*tchttp.BaseRequest

	// 文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 文件内容的base64字符串

	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`
	// 图片名

	PicName *string `json:"PicName,omitempty" name:"PicName"`
	// 图片内容的base64字符串

	PicContent *string `json:"PicContent,omitempty" name:"PicContent"`
}

func (r *CreateAppFromChartPkgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppFromChartPkgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateScalerRequest struct {
	*tchttp.BaseRequest

	// 弹性伸缩策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 弹性伸缩策略关联应用实例的ID

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
	// 弹性伸缩策略关联组件的ID

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
	// 弹性伸缩策略关联组件的名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 触发方式

	Type *string `json:"Type,omitempty" name:"Type"`
	// 触发策略（含实例范围)

	TriggerStrategy *string `json:"TriggerStrategy,omitempty" name:"TriggerStrategy"`
	// 冷却时间

	CoolDown *int64 `json:"CoolDown,omitempty" name:"CoolDown"`
	// 默认是否开启

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 弹性伸缩策略的id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *CreateScalerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScalerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationResourceUsageSummaryRequest struct {
	*tchttp.BaseRequest

	// 无

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 无

	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`
	// 无

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 无

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetApplicationResourceUsageSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationResourceUsageSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComponentInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComponentInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComponentInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRolloutSettingRequest struct {
	*tchttp.BaseRequest

	// 组件实例的ID号

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
}

func (r *GetRolloutSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRolloutSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateImage struct {
	// 容器名称

	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
	// 更新的镜像名称

	Image *string `json:"Image,omitempty" name:"Image"`
}

type GetAppInstanceYamlRequest struct {
	*tchttp.BaseRequest

	// 应用实例Id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetAppInstanceYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppInstanceYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRolloutStrategyRequest struct {
	*tchttp.BaseRequest

	// 组件实例的ID

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
}

func (r *GetRolloutStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRolloutStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveRevisionFromYamlRequest struct {
	*tchttp.BaseRequest

	// revision版本

	Revision *string `json:"Revision,omitempty" name:"Revision"`
	// yaml内容

	YamlList []*string `json:"YamlList,omitempty" name:"YamlList"`
	// App实例id

	AppInstanceId *string `json:"AppInstanceId,omitempty" name:"AppInstanceId"`
	// 操作的者用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 操作的者用户ID

	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`
}

func (r *SaveRevisionFromYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveRevisionFromYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Parameter struct {
	// 参数Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 参数名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数对应模板的路径

	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 参数类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 是否必要参数

	Required *bool `json:"Required,omitempty" name:"Required"`
	// 参数所属模板Id

	SpecTemplateId *string `json:"SpecTemplateId,omitempty" name:"SpecTemplateId"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type Parameters struct {
	// 参数实例

	ParameterInstanceId *string `json:"ParameterInstanceId,omitempty" name:"ParameterInstanceId"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateAppFromTadPkgRequest struct {
	*tchttp.BaseRequest

	// 包名字

	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
	// 包版本

	TemplateVersion *string `json:"TemplateVersion,omitempty" name:"TemplateVersion"`
	// 默认为tad

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
	// yaml或者pkgfile

	PkgType *string `json:"PkgType,omitempty" name:"PkgType"`
	// 包内容的二进制的base64字符串

	PkgContent *string `json:"PkgContent,omitempty" name:"PkgContent"`
}

func (r *CreateAppFromTadPkgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppFromTadPkgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppsFromTadProductPkgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建批量模版返回结果

		CreateAppResp []*CreateAppResp `json:"CreateAppResp,omitempty" name:"CreateAppResp"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppsFromTadProductPkgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppsFromTadProductPkgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigMapFromYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConfigMapFromYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigMapFromYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecretRequest struct {
	*tchttp.BaseRequest

	// secret 名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// secret 所在 project

	Project *string `json:"Project,omitempty" name:"Project"`
	// secret所在namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// secret类别

	Type *string `json:"Type,omitempty" name:"Type"`
	// secret 内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 创建者用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *CreateSecretRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecretRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableVpaScalerRequest struct {
	*tchttp.BaseRequest

	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 1 或者0

	Enabel *int64 `json:"Enabel,omitempty" name:"Enabel"`
}

func (r *EnableVpaScalerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableVpaScalerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetNewestRevisionRequest struct {
	*tchttp.BaseRequest

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 无

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *GetNewestRevisionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetNewestRevisionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableVpaScalerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Id

		Id *string `json:"Id,omitempty" name:"Id"`
		// ActionResult

		ActionResult *string `json:"ActionResult,omitempty" name:"ActionResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableVpaScalerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableVpaScalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCanaryPartitionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCanaryPartitionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCanaryPartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetScalerYamlRequest struct {
	*tchttp.BaseRequest

	// 弹性伸缩的id

	ScalerId *string `json:"ScalerId,omitempty" name:"ScalerId"`
}

func (r *GetScalerYamlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetScalerYamlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTemplatesRequest struct {
	*tchttp.BaseRequest

	// 模板种类

	Kind *string `json:"Kind,omitempty" name:"Kind"`
}

func (r *ListTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppFromYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用模板Id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppFromYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAppFromYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollbackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RollbackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveRevisionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// id

		Id *string `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveRevisionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveRevisionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAppInstanceComponentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Component数组

		Components []*Component `json:"Components,omitempty" name:"Components"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAppInstanceComponentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppInstanceComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppProductsRequest struct {
	*tchttp.BaseRequest

	// 命名空间

	Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces"`
	// 产品名

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ListAppProductsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppProductsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSecretsRequest struct {
	*tchttp.BaseRequest

	// 上限

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 起始页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序

	Sort []*Sort `json:"Sort,omitempty" name:"Sort"`
	// TLS等类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// secret 名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ListSecretsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSecretsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParameterInput struct {
	// 参数名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数对应模板路径

	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type TraitInstance struct {
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 所属组件实例Id

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
	// 关联的traitId

	TraitId *string `json:"TraitId,omitempty" name:"TraitId"`
	// 参数列表

	Parameters []*TraitInstanceParameter `json:"Parameters,omitempty" name:"Parameters"`
	// 输出实例列表

	Outputs []*OutputInstance `json:"Outputs,omitempty" name:"Outputs"`
}

type GetInstanceReplicasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ComponentInstanceId

		ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
		// Replicas

		Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetInstanceReplicasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetInstanceReplicasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppCategoriesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用模板类别列表

		Categories []*string `json:"Categories,omitempty" name:"Categories"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAppCategoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppCategoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListNamespacesRequest struct {
	*tchttp.BaseRequest
}

func (r *ListNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BuildVersion struct {
	// 无

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// 无

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 无

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 无

	Id *string `json:"Id,omitempty" name:"Id"`
	// 无

	IsRunning *string `json:"IsRunning,omitempty" name:"IsRunning"`
	// 无

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 无

	Revision *string `json:"Revision,omitempty" name:"Revision"`
}

type GetSecretRequest struct {
	*tchttp.BaseRequest

	// secret id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetSecretRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSecretRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSecretResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// secret内容

		Secret *Secret `json:"Secret,omitempty" name:"Secret"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSecretResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListConfigMapsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListConfigMapsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListConfigMapsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TraitParameter struct {
	// 参数Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 是否必要参数

	Required *bool `json:"Required,omitempty" name:"Required"`
	// 参数值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 参数所属trait

	TraitId *string `json:"TraitId,omitempty" name:"TraitId"`
	// 对应模板的路径

	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
}

type GetRolloutStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRolloutStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRolloutStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetUpgradeConfigRequest struct {
	*tchttp.BaseRequest

	// 组件实例id

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
	// Strategy

	Strategy *Strategy `json:"Strategy,omitempty" name:"Strategy"`
}

func (r *SetUpgradeConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetUpgradeConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAppResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAppResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppVersionsRequest struct {
	*tchttp.BaseRequest

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 集群

	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
	// 最大值

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 排序

	Sort []*Sort `json:"Sort,omitempty" name:"Sort"`
}

func (r *ListAppVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAppVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RevisionTrait struct {
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 组件id

	RevisionComponentId *string `json:"RevisionComponentId,omitempty" name:"RevisionComponentId"`
	// trait id

	TraitInstanceId *string `json:"TraitInstanceId,omitempty" name:"TraitInstanceId"`
	// 参数

	Parameters []*RevisionTraitParameter `json:"Parameters,omitempty" name:"Parameters"`
	// input

	Inputs []*RevisionInput `json:"Inputs,omitempty" name:"Inputs"`
	// output

	Outputs []*RevisionOutput `json:"Outputs,omitempty" name:"Outputs"`
}

type ListScalerEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 弹性伸缩策略总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 弹性伸缩历史事件列表

		Scalers []*Event `json:"Scalers,omitempty" name:"Scalers"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListScalerEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListScalerEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAppFromYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAppFromYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAppFromYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InputInstance struct {
	// id

	Id *string `json:"Id,omitempty" name:"Id"`
	// input id

	InputId *string `json:"InputId,omitempty" name:"InputId"`
	// 名称

	OutputName *string `json:"OutputName,omitempty" name:"OutputName"`
	// 目标路径

	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 服务Id

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 范围

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 是否必须

	Required *bool `json:"Required,omitempty" name:"Required"`
	// 所属组件实例id

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
	// 所属trait实例id

	TraitInstanceId *string `json:"TraitInstanceId,omitempty" name:"TraitInstanceId"`
}

type GetAppRequest struct {
	*tchttp.BaseRequest

	// App模板Id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *GetAppRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAppRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkloadPodsRequest struct {
	*tchttp.BaseRequest

	// 目标应用实例的名字

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 目标应用实例的命名空间

	InstanceNamespace *string `json:"InstanceNamespace,omitempty" name:"InstanceNamespace"`
	// 目标workload的名字

	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
	// 目标workload的类型

	WorkloadKind *string `json:"WorkloadKind,omitempty" name:"WorkloadKind"`
	// 目标workload的命名空间

	WorkloadNamespace *string `json:"WorkloadNamespace,omitempty" name:"WorkloadNamespace"`
}

func (r *ListWorkloadPodsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListWorkloadPodsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComponentInput struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 负载类型

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 组件参数

	Parameters []*ParameterInput `json:"Parameters,omitempty" name:"Parameters"`
	// 关联的trait

	Traits []*TraitInput `json:"Traits,omitempty" name:"Traits"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 配置数据

	Configs []*Config `json:"Configs,omitempty" name:"Configs"`
	// 输出参数依赖

	Outputs []*OutputReq `json:"Outputs,omitempty" name:"Outputs"`
	// 输入参数依赖

	Inputs []*Input `json:"Inputs,omitempty" name:"Inputs"`
}

type ConfigInstance struct {
	// 配置id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 配置名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 关联的实例Id

	ComponentInstanceId *string `json:"ComponentInstanceId,omitempty" name:"ComponentInstanceId"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 内容

	RawContent *string `json:"RawContent,omitempty" name:"RawContent"`
	// 依赖输入

	Inputs []*InputInstance `json:"Inputs,omitempty" name:"Inputs"`
	// 无

	KVs []*ConfigKV `json:"KVs,omitempty" name:"KVs"`
}

type GetRolloutSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRolloutSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRolloutSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListVpaScalerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Total

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// VpaScalers

		VpaScalers []*VpaScaler `json:"VpaScalers,omitempty" name:"VpaScalers"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListVpaScalerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListVpaScalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveConfigMapYamlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveConfigMapYamlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveConfigMapYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
