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

package v20210801

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DeleteDataSourceRequest struct {
	*tchttp.BaseRequest

	// 数据源uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
}

func (r *DeleteDataSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDataSourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Project struct {
	// 项目唯一ID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 项目名称

	Label *string `json:"Label,omitempty" name:"Label"`
	// 项目中文名称

	LabelCn *string `json:"LabelCn,omitempty" name:"LabelCn"`
	// 项目英文名称

	LabelEn *string `json:"LabelEn,omitempty" name:"LabelEn"`
	// 项目描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 修改者

	Updater *string `json:"Updater,omitempty" name:"Updater"`
	// 同项目页面

	Pages []*Page `json:"Pages,omitempty" name:"Pages"`
	// 项目免密信息

	AuthList []*Auth `json:"AuthList,omitempty" name:"AuthList"`
}

type ModifyDataSourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDataSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Count *int64 `json:"Count,omitempty" name:"Count"`
		// 页面列表

		List []*Page `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPageRequest struct {
	*tchttp.BaseRequest

	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 页面uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 页面配置json

	Conf *string `json:"Conf,omitempty" name:"Conf"`
	// 页面中文名

	LabelCn *string `json:"LabelCn,omitempty" name:"LabelCn"`
	// 页面英文名

	LabelEn *string `json:"LabelEn,omitempty" name:"LabelEn"`
	// 页面所属产品Code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 页面描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 项目id

	ProjectUuid *string `json:"ProjectUuid,omitempty" name:"ProjectUuid"`
	// 最后更新Token

	Lct *string `json:"Lct,omitempty" name:"Lct"`
}

func (r *ModifyPageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryComponentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件列表

		List []*Component `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryComponentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDataSourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDataSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMiddlePlatformRequest struct {
	*tchttp.BaseRequest

	// 查询语句

	Fql *string `json:"Fql,omitempty" name:"Fql"`
	// 查询语句

	LastFql *string `json:"LastFql,omitempty" name:"LastFql"`
}

func (r *QueryMiddlePlatformRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryMiddlePlatformRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductCodeInPagesRequest struct {
	*tchttp.BaseRequest

	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 搜索参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryProductCodeInPagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductCodeInPagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filters struct {
	// 搜索key

	Name *string `json:"Name,omitempty" name:"Name"`
	// 搜索值

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type CopyPageRequest struct {
	*tchttp.BaseRequest

	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 所属产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 页面英文名称

	LabelEn *string `json:"LabelEn,omitempty" name:"LabelEn"`
	// 页面唯一id

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 页面名称

	LabelCn *string `json:"LabelCn,omitempty" name:"LabelCn"`
	// 页面描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 来源id

	FromUuid *string `json:"FromUuid,omitempty" name:"FromUuid"`
}

func (r *CopyPageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyPageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishPageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishPageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDataSourcesRequest struct {
	*tchttp.BaseRequest

	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 是否包含配置文件

	InCludeConf *bool `json:"InCludeConf,omitempty" name:"InCludeConf"`
	// 搜索参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryDataSourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDataSourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetConfPageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetConfPageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetConfPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDataSourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDataSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest

	// 项目uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDataSourceRequest struct {
	*tchttp.BaseRequest

	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 数据源类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 数据源英文名称

	LabelEn *string `json:"LabelEn,omitempty" name:"LabelEn"`
	// 数据源唯一id

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 数据源名称

	LabelCn *string `json:"LabelCn,omitempty" name:"LabelCn"`
	// 数据源描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 数据源配置

	DataConf *string `json:"DataConf,omitempty" name:"DataConf"`
	// 数据源返回值

	RespConf *string `json:"RespConf,omitempty" name:"RespConf"`
	// 数据源标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

func (r *ModifyDataSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDataSourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDataSourceRequest struct {
	*tchttp.BaseRequest

	// 数据源uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
}

func (r *QueryDataSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDataSourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductCodeInPagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 页面产品列表

		List []*PageProduct `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryProductCodeInPagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductCodeInPagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePageRequest struct {
	*tchttp.BaseRequest

	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 所属产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 页面英文名称

	LabelEn *string `json:"LabelEn,omitempty" name:"LabelEn"`
	// 页面唯一id

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 页面名称

	LabelCn *string `json:"LabelCn,omitempty" name:"LabelCn"`
	// 页面描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 项目唯一id

	ProjectUuid *string `json:"ProjectUuid,omitempty" name:"ProjectUuid"`
	// 页面类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 组件名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 组件版本

	ComponentVersion *int64 `json:"ComponentVersion,omitempty" name:"ComponentVersion"`
	// 主题类型

	Theme *string `json:"Theme,omitempty" name:"Theme"`
	// 页面初始配置

	BaseConf *string `json:"BaseConf,omitempty" name:"BaseConf"`
}

func (r *CreatePageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDataSourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Count *int64 `json:"Count,omitempty" name:"Count"`
		// 页面列表

		List []*Page `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDataSourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDataSourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefreshAuthRequest struct {
	*tchttp.BaseRequest
}

func (r *RefreshAuthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefreshAuthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest

	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 项目英文名称

	LabelEn *string `json:"LabelEn,omitempty" name:"LabelEn"`
	// 项目唯一id

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 项目名称

	LabelCn *string `json:"LabelCn,omitempty" name:"LabelCn"`
	// 项目描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAuthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 免密加密后token

		Token *string `json:"Token,omitempty" name:"Token"`
		// 免密ID

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetAuthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Auth struct {
	// 页面ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 免密标题

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过期时间

	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// 免密口令

	Psw *string `json:"Psw,omitempty" name:"Psw"`
	// 免密类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 关联项目

	PageProject *string `json:"PageProject,omitempty" name:"PageProject"`
	// 免密token

	Token *string `json:"Token,omitempty" name:"Token"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 更新时间

	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`
	// 更新者

	Updater *string `json:"Updater,omitempty" name:"Updater"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
}

type QueryDataSouceDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDataSouceDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDataSouceDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAuthRequest struct {
	*tchttp.BaseRequest

	// 配置ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 免密标题

	Name *string `json:"Name,omitempty" name:"Name"`
	// 免密项目uuid

	Project *string `json:"Project,omitempty" name:"Project"`
	// 免密页面uuid

	Pages []*string `json:"Pages,omitempty" name:"Pages"`
	// 过期时间戳

	ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// 免密口令

	Psw *string `json:"Psw,omitempty" name:"Psw"`
	// 描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *SetAuthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAuthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Component struct {
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 修改者

	Updater *string `json:"Updater,omitempty" name:"Updater"`
	// 是否删除

	Del *bool `json:"Del,omitempty" name:"Del"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 组件名称

	Label *string `json:"Label,omitempty" name:"Label"`
	// 组件code

	Name *string `json:"Name,omitempty" name:"Name"`
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 组件类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 组件次类型

	SubType *string `json:"SubType,omitempty" name:"SubType"`
	// 组件版本

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// 组件小版本

	BuildNum *int64 `json:"BuildNum,omitempty" name:"BuildNum"`
	// 组件icon 地址

	Icon *string `json:"Icon,omitempty" name:"Icon"`
	// 组件预览图地址

	PreviewImg *string `json:"PreviewImg,omitempty" name:"PreviewImg"`
	// 组件配置json string

	Conf *string `json:"Conf,omitempty" name:"Conf"`
}

type CreatePageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPagesRequest struct {
	*tchttp.BaseRequest

	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 是否包含配置文件

	InCludeConf *bool `json:"InCludeConf,omitempty" name:"InCludeConf"`
	// 搜索参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 支持国际化

	I18n *bool `json:"I18n,omitempty" name:"I18n"`
}

func (r *QueryPagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProjectsRequest struct {
	*tchttp.BaseRequest

	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 搜索参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
	// 分页偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页页数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *QueryProjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProjectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type YsInput struct {
	// 起始时间

	Start *string `json:"Start,omitempty" name:"Start"`
	// 结束时间

	End *string `json:"End,omitempty" name:"End"`
	// 步长

	Step *string `json:"Step,omitempty" name:"Step"`
	// 查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
}

type ApiProxyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApiProxyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApiProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishPageRequest struct {
	*tchttp.BaseRequest

	// 页面uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 最后更新Token

	Lct *string `json:"Lct,omitempty" name:"Lct"`
}

func (r *PublishPageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishPageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PageSort struct {
	// 页面UUID

	PcUuid *string `json:"PcUuid,omitempty" name:"PcUuid"`
	// 序号

	Sort *int64 `json:"Sort,omitempty" name:"Sort"`
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest

	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 项目英文名称

	LabelEn *string `json:"LabelEn,omitempty" name:"LabelEn"`
	// 项目唯一id

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 项目名称

	LabelCn *string `json:"LabelCn,omitempty" name:"LabelCn"`
	// 项目描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePageRequest struct {
	*tchttp.BaseRequest

	// 页面uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 最后更新Token

	Lct *string `json:"Lct,omitempty" name:"Lct"`
}

func (r *DeletePageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MpInput struct {
	// 查询sql

	Fql *string `json:"Fql,omitempty" name:"Fql"`
}

type QueryTagsRequest struct {
	*tchttp.BaseRequest

	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页每页数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Page struct {
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 修改者

	Updater *string `json:"Updater,omitempty" name:"Updater"`
	// 唯一code

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 名称

	Label *string `json:"Label,omitempty" name:"Label"`
	// 描述详情

	Description *string `json:"Description,omitempty" name:"Description"`
	// 页面类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 页面json

	ComponentConf *string `json:"ComponentConf,omitempty" name:"ComponentConf"`
	// 页面中文名称

	LabelCn *string `json:"LabelCn,omitempty" name:"LabelCn"`
	// 页面英文名称

	LabelEn *string `json:"LabelEn,omitempty" name:"LabelEn"`
	// 所属项目

	Project *Project `json:"Project,omitempty" name:"Project"`
	// 发布时间

	ReleasedTime *string `json:"ReleasedTime,omitempty" name:"ReleasedTime"`
	// 页面基本配置

	BaseConf *string `json:"BaseConf,omitempty" name:"BaseConf"`
	// 页面顶层组件

	ComponetName *string `json:"ComponetName,omitempty" name:"ComponetName"`
	// 页面顶层组件版本

	CompnentVersion *uint64 `json:"CompnentVersion,omitempty" name:"CompnentVersion"`
	// 页面发布json

	ComponentRealsedConf *string `json:"ComponentRealsedConf,omitempty" name:"ComponentRealsedConf"`
	// 所属项目uuid

	ProjectUuid *string `json:"ProjectUuid,omitempty" name:"ProjectUuid"`
	// 所属产品Code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 页面序号

	Sort *uint64 `json:"Sort,omitempty" name:"Sort"`
	// 同项目导航可见

	ShowInProjectNav *bool `json:"ShowInProjectNav,omitempty" name:"ShowInProjectNav"`
	// 免密信息

	AuthList []*Auth `json:"AuthList,omitempty" name:"AuthList"`
}

type PageProduct struct {
	// 产品 code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 页面数

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type Visible struct {
	// 页面uuid

	PcUuid *string `json:"PcUuid,omitempty" name:"PcUuid"`
	// 是否可见

	Visible *bool `json:"Visible,omitempty" name:"Visible"`
}

type BindProjectWithPageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindProjectWithPageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindProjectWithPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDataSourceRequest struct {
	*tchttp.BaseRequest

	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 数据源类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 数据源英文名称

	LabelEn *string `json:"LabelEn,omitempty" name:"LabelEn"`
	// 数据源唯一id

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 数据源名称

	LabelCn *string `json:"LabelCn,omitempty" name:"LabelCn"`
	// 数据源描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 数据源配置

	DataConf *string `json:"DataConf,omitempty" name:"DataConf"`
	// 数据源返回值

	RespConf *string `json:"RespConf,omitempty" name:"RespConf"`
	// 数据源标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateDataSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDataSourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryComponentsRequest struct {
	*tchttp.BaseRequest

	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 搜索参数

	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *QueryComponentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryComponentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDataSouceDataRequest struct {
	*tchttp.BaseRequest

	// 独立数据源Id

	DsId *string `json:"DsId,omitempty" name:"DsId"`
	// 数据源类型，云api: yunApi, 云哨：yunShao ; 中台： middlePlatform

	DsType *string `json:"DsType,omitempty" name:"DsType"`
	// 数据源所引用变量json字符串数据

	VariableList *string `json:"VariableList,omitempty" name:"VariableList"`
}

func (r *QueryDataSouceDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDataSouceDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigProjectRequest struct {
	*tchttp.BaseRequest

	// 项目UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 页面排序信息

	Sort []*PageSort `json:"Sort,omitempty" name:"Sort"`
	// 页面可见信息

	Visible []*Visible `json:"Visible,omitempty" name:"Visible"`
}

func (r *ConfigProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfigProjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyPageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyPageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefreshAuthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RefreshAuthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefreshAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetConfPageRequest struct {
	*tchttp.BaseRequest

	// 页面UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 页面基本配置

	Conf *string `json:"Conf,omitempty" name:"Conf"`
}

func (r *SetConfPageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetConfPageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type YapiInput struct {
	// 版本

	Verion *string `json:"Verion,omitempty" name:"Verion"`
}

type ApiProxyRequest struct {
	*tchttp.BaseRequest

	// 云哨数据源配置的入参

	YsInput *YsInput `json:"YsInput,omitempty" name:"YsInput"`
	// 云API数据源配置的入参

	YapiInput *YapiInput `json:"YapiInput,omitempty" name:"YapiInput"`
	// 中台数据源配置的入参

	MpInput *MpInput `json:"MpInput,omitempty" name:"MpInput"`
}

func (r *ApiProxyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApiProxyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMiddlePlatformResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryMiddlePlatformResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryMiddlePlatformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Count *int64 `json:"Count,omitempty" name:"Count"`
		// 页面列表

		List []*Page `json:"List,omitempty" name:"List"`
		// 环境信息

		TceEnv *bool `json:"TceEnv,omitempty" name:"TceEnv"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryPagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProjectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Count *int64 `json:"Count,omitempty" name:"Count"`
		// 项目列表

		List []*Project `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryProjectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindProjectWithPageRequest struct {
	*tchttp.BaseRequest

	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 页面UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 项目UUID

	ProjectUuid *string `json:"ProjectUuid,omitempty" name:"ProjectUuid"`
}

func (r *BindProjectWithPageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindProjectWithPageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigProjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConfigProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfigProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDataSourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDataSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
