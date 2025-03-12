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

type GetServiceApiListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// api列表

		ApiList []*TagApiList `json:"ApiList,omitempty" name:"ApiList"`
		// ArnDocument

		ArnDocument *string `json:"ArnDocument,omitempty" name:"ArnDocument"`
		// ConditionKeyList

		ConditionKeyList []*string `json:"ConditionKeyList,omitempty" name:"ConditionKeyList"`
		// 英文名

		EnName *string `json:"EnName,omitempty" name:"EnName"`
		// 中文名

		Name *string `json:"Name,omitempty" name:"Name"`
		// 服务类型

		ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServiceApiListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceApiListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigArray struct {
	// 版本

	DataVersion *string `json:"DataVersion,omitempty" name:"DataVersion"`
	// 是否显示

	Display *string `json:"Display,omitempty" name:"Display"`
	// 白名单

	WhiteList []*string `json:"WhiteList,omitempty" name:"WhiteList"`
	// 资源类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 中文名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 英文名

	EnName *string `json:"EnName,omitempty" name:"EnName"`
	// 子项

	Child []*ChildArray `json:"Child,omitempty" name:"Child"`
}

type CheckTagResourceInfoRequest struct {
	*tchttp.BaseRequest

	// 服务ID

	AccessServiceId *int64 `json:"AccessServiceId,omitempty" name:"AccessServiceId"`
	// 子节点中文名

	ChildEnname *string `json:"ChildEnname,omitempty" name:"ChildEnname"`
	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 子节点

	ChildName *string `json:"ChildName,omitempty" name:"ChildName"`
	// 资源前缀

	ChildResourcePrefix *string `json:"ChildResourcePrefix,omitempty" name:"ChildResourcePrefix"`
}

func (r *CheckTagResourceInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckTagResourceInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckTagResourceInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckTagResourceInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckTagResourceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagResourceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// List

		List []*ObjectArrayTag `json:"List,omitempty" name:"List"`
		// 数量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagResourceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagResourceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagServiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务列表

		List []*TagService `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagServiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagService struct {
	// ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 服务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 服务英文名

	EnName *string `json:"EnName,omitempty" name:"EnName"`
	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 是否隐藏

	Display *string `json:"Display,omitempty" name:"Display"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// WhiteList

	WhiteList []*string `json:"WhiteList,omitempty" name:"WhiteList"`
}

type DescribeTagServiceListRequest struct {
	*tchttp.BaseRequest

	// 参数固定为last-production

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeTagServiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagServiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTagStatusConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateTagStatusConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTagStatusConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagResourceListRequest struct {
	*tchttp.BaseRequest

	// 3：代表发布，1：默认，2、预发布

	Status []*int64 `json:"Status,omitempty" name:"Status"`
	// 资源类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 资源前缀

	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
	// 资源中文名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 一页多少条

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 第几页

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 搜索：英文名

	Enname *string `json:"Enname,omitempty" name:"Enname"`
	// 搜索：是否可见

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeTagResourceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagResourceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTagStatusConfigRequest struct {
	*tchttp.BaseRequest

	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 0：设置不可见 1：设置可见

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *UpdateTagStatusConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTagStatusConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ObjectArrayTag struct {
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// AccessServiceId

	AccessServiceId *uint64 `json:"AccessServiceId,omitempty" name:"AccessServiceId"`
	// 服务类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// Index

	Index *uint64 `json:"Index,omitempty" name:"Index"`
	// 中文名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 英文名

	EnName *string `json:"EnName,omitempty" name:"EnName"`
	// 资源前缀

	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
	// 是否隐藏

	Display *string `json:"Display,omitempty" name:"Display"`
	// CapiServiceType

	CapiServiceType *string `json:"CapiServiceType,omitempty" name:"CapiServiceType"`
	// CapiAction

	CapiAction *string `json:"CapiAction,omitempty" name:"CapiAction"`
	// CapiVersion

	CapiVersion *string `json:"CapiVersion,omitempty" name:"CapiVersion"`
	// CapiTransformInFunc

	CapiTransformInFunc *string `json:"CapiTransformInFunc,omitempty" name:"CapiTransformInFunc"`
	// CapiTransformOutFunc

	CapiTransformOutFunc *string `json:"CapiTransformOutFunc,omitempty" name:"CapiTransformOutFunc"`
	// 擦咯做人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// Status

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 服务类型

	RealServiceType *string `json:"RealServiceType,omitempty" name:"RealServiceType"`
	// InstanceSearch

	InstanceSearch *InstanceSearch `json:"InstanceSearch,omitempty" name:"InstanceSearch"`
	// 英文名

	RealEnName *string `json:"RealEnName,omitempty" name:"RealEnName"`
	// 是否预设数据

	PresetData *uint64 `json:"PresetData,omitempty" name:"PresetData"`
	// 资源名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 真实中文名

	RealName *string `json:"RealName,omitempty" name:"RealName"`
}

type GetServiceApiListRequest struct {
	*tchttp.BaseRequest

	// 按服务Id匹配，如cvm

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 默认1

	IsReturnApi *int64 `json:"IsReturnApi,omitempty" name:"IsReturnApi"`
}

func (r *GetServiceApiListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceApiListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveTagServiceConfigRequest struct {
	*tchttp.BaseRequest

	// 标签配置

	Config []*ConfigArray `json:"Config,omitempty" name:"Config"`
	// 资源类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 操作者

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *SaveTagServiceConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveTagServiceConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateServiceConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateServiceConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagApiList struct {
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// IsNeedObject

	IsNeedObject *string `json:"IsNeedObject,omitempty" name:"IsNeedObject"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type SaveTagServiceConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveTagServiceConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveTagServiceConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceConfigRequest struct {
	*tchttp.BaseRequest

	// 服务

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 服务英文名

	ServiceEnName *string `json:"ServiceEnName,omitempty" name:"ServiceEnName"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 配置

	ConfigList []*ConfigList `json:"ConfigList,omitempty" name:"ConfigList"`
	// 是否删除服务

	DeleteService *bool `json:"DeleteService,omitempty" name:"DeleteService"`
}

func (r *UpdateServiceConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateServiceConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChildArray struct {
	// 是否隐藏

	Display *string `json:"Display,omitempty" name:"Display"`
	// 白名单列表

	WhiteList []*string `json:"WhiteList,omitempty" name:"WhiteList"`
	// 资源类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 资源前缀

	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
	// index

	Index *int64 `json:"Index,omitempty" name:"Index"`
	// 中文名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 英文名

	EnName *string `json:"EnName,omitempty" name:"EnName"`
	// serviceType

	CapiServiceType *string `json:"CapiServiceType,omitempty" name:"CapiServiceType"`
	// 接口名

	CapiAction *string `json:"CapiAction,omitempty" name:"CapiAction"`
	// 版本号

	CapiVersion *string `json:"CapiVersion,omitempty" name:"CapiVersion"`
	// capiTransformInFunc

	CapiTransformInFunc *string `json:"CapiTransformInFunc,omitempty" name:"CapiTransformInFunc"`
	// capiTransformOutFunc

	CapiTransformOutFunc *string `json:"CapiTransformOutFunc,omitempty" name:"CapiTransformOutFunc"`
	// 资源ID检索开关

	InstanceSearch *InstanceSearch `json:"InstanceSearch,omitempty" name:"InstanceSearch"`
	// child

	Child []*string `json:"Child,omitempty" name:"Child"`
	// 新资源类型

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

type ConfigList struct {
	// 新资源名称

	NewResourceName *string `json:"NewResourceName,omitempty" name:"NewResourceName"`
	// 新资源英文名

	NewResourceEnName *string `json:"NewResourceEnName,omitempty" name:"NewResourceEnName"`
	// 新资源前缀

	NewResourcePrefix *string `json:"NewResourcePrefix,omitempty" name:"NewResourcePrefix"`
	// 旧资源名称

	OldResourceName *string `json:"OldResourceName,omitempty" name:"OldResourceName"`
	// 旧资源英文名

	OldResourceEnName *string `json:"OldResourceEnName,omitempty" name:"OldResourceEnName"`
	// 旧资源前缀

	OldResourcePrefix *string `json:"OldResourcePrefix,omitempty" name:"OldResourcePrefix"`
}

type InstanceSearch struct {
	// ParamKey

	ParamKey *string `json:"ParamKey,omitempty" name:"ParamKey"`
	// Enable

	Enable *string `json:"Enable,omitempty" name:"Enable"`
}
