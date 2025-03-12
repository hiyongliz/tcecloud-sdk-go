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

package v2018109

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type EditDirResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditDirResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditDirResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryNodeUpgradeRequest struct {
	*tchttp.BaseRequest

	// input

	Params *QueryNodeUpgradeParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryNodeUpgradeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryNodeUpgradeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryNodeUpgradeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryNodeUpgradeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryNodeUpgradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCategoriesParams struct {
	// 平台

	Platform *string `json:"Platform,omitempty" name:"Platform"`
}

type QueryNodePreviewParams struct {
	// 文档树id

	TdocTreeId *uint64 `json:"TdocTreeId,omitempty" name:"TdocTreeId"`
	// 来源,值有tdoc_tree和tdoc_tree_upgrade

	Origin *string `json:"Origin,omitempty" name:"Origin"`
}

type QueryCategoriesRequest struct {
	*tchttp.BaseRequest

	// 参数

	Params *QueryCategoriesParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryCategoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCategoriesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTreePreviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryTreePreviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTreePreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PathRequest struct {
	*tchttp.BaseRequest

	// 无

	Params *PathParams `json:"Params,omitempty" name:"Params"`
}

func (r *PathRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PathRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelDirParams struct {
	// TdocTreeId

	TdocTreeId *int64 `json:"TdocTreeId,omitempty" name:"TdocTreeId"`
	// Action

	Action *string `json:"Action,omitempty" name:"Action"`
	// Editor

	Editor *string `json:"Editor,omitempty" name:"Editor"`
}

type EditNodeParams struct {
	// TdocTreeId

	TdocTreeId *int64 `json:"TdocTreeId,omitempty" name:"TdocTreeId"`
	// parentId

	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`
	// attr

	Attr *string `json:"Attr,omitempty" name:"Attr"`
	// alias

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// title

	Title *string `json:"Title,omitempty" name:"Title"`
	// title

	Body *string `json:"Body,omitempty" name:"Body"`
	// weight

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// editor

	Editor *string `json:"Editor,omitempty" name:"Editor"`
	// Deleted

	Deleted *int64 `json:"Deleted,omitempty" name:"Deleted"`
	// TdocReviewId

	TdocReviewId *int64 `json:"TdocReviewId,omitempty" name:"TdocReviewId"`
}

type DecideRequest struct {
	*tchttp.BaseRequest

	// wu

	Params *DecideParams `json:"Params,omitempty" name:"Params"`
}

func (r *DecideRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DecideRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryExportLogRequest struct {
	*tchttp.BaseRequest

	// 参数

	Params *QueryExportLogParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryExportLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryExportLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryNodePreviewRequest struct {
	*tchttp.BaseRequest

	// 参数

	Params *QueryNodePreviewParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryNodePreviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryNodePreviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDocProductCodesParams struct {
	// 平台，ocloud/tcloud

	Platform *string `json:"Platform,omitempty" name:"Platform"`
}

type DelDirRequest struct {
	*tchttp.BaseRequest

	// 无

	Params *DelDirParams `json:"Params,omitempty" name:"Params"`
}

func (r *DelDirRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelDirRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTreeParams struct {
	// Editor

	Editor *string `json:"Editor,omitempty" name:"Editor"`
	// TdocTreeId

	TdocTreeId *int64 `json:"TdocTreeId,omitempty" name:"TdocTreeId"`
	// Depth

	Depth *uint64 `json:"Depth,omitempty" name:"Depth"`
}

type QueryDocProductCodesRequest struct {
	*tchttp.BaseRequest

	// 参数

	Params *QueryDocProductCodesParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryDocProductCodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDocProductCodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDocProductCodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回数据

		Data *QueryDocProductCodesData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDocProductCodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDocProductCodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InsertOrUpdateNodeParams struct {
	// 节点列表

	NodeList []*Node `json:"NodeList,omitempty" name:"NodeList"`
}

type QueryCategoriesData struct {
	// 分类列表

	CategoryList []*CategoryData `json:"CategoryList,omitempty" name:"CategoryList"`
}

type AddDirResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddDirResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddDirResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeableResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateUpgradeLogParams struct {
	// 状态，当前允许值7/8

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 编辑人

	Editor *string `json:"Editor,omitempty" name:"Editor"`
}

type AddNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DecideResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DecideResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DecideResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GiveupUpgradeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GiveupUpgradeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GiveupUpgradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCategoriesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接口返回数据

		Data *QueryCategoriesData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCategoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCategoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Node struct {
	// 能所有到节点的url

	TreePath *string `json:"TreePath,omitempty" name:"TreePath"`
	// 标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 内容

	Body *string `json:"Body,omitempty" name:"Body"`
	// 是否对外可见，0 可见 3 不可见

	Visibility *uint64 `json:"Visibility,omitempty" name:"Visibility"`
	// 属性 ，dir 为目录 doc 为文档

	Attr *string `json:"Attr,omitempty" name:"Attr"`
	// Alias

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 权重，默认为0，值越大越靠前

	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
}

type QueryProductTreePathParams struct {
	// 平台 ocloud/tcloud

	Platform *string `json:"Platform,omitempty" name:"Platform"`
	// 产品 code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

type QueryDirResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDirResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDirResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReviewRequest struct {
	*tchttp.BaseRequest

	// wu

	Params *ReviewParams `json:"Params,omitempty" name:"Params"`
}

func (r *ReviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DecideParams struct {
	// tdocReviewIds

	TdocReviewIds []*int64 `json:"TdocReviewIds,omitempty" name:"TdocReviewIds"`
	// reviewer

	Reviewer *string `json:"Reviewer,omitempty" name:"Reviewer"`
	// action

	Action *string `json:"Action,omitempty" name:"Action"`
	// reason

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type QueryExportLogParams struct {
	// 导出记录id

	ExportLogId *uint64 `json:"ExportLogId,omitempty" name:"ExportLogId"`
	// 导出人

	Editor *string `json:"Editor,omitempty" name:"Editor"`
	// 状态

	Status []*uint64 `json:"Status,omitempty" name:"Status"`
	// 页数

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 每页条数

	Size *uint64 `json:"Size,omitempty" name:"Size"`
}

type DelNodeParams struct {
	// 能索引到节点的url，例如：tcloud/Compute/CVM/APIs，通过斜杆"/"分割大于3

	TreePath *string `json:"TreePath,omitempty" name:"TreePath"`
}

type QueryNodeUpgradeParams struct {
	// id

	TdocTreeId *uint64 `json:"TdocTreeId,omitempty" name:"TdocTreeId"`
}

type QueryProductTreePathData struct {
	// 树路径，例如 Compute/CVM

	TreePath *string `json:"TreePath,omitempty" name:"TreePath"`
}

type QueryTreeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryTreeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateUpgradeLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateUpgradeLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateUpgradeLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportDocParams struct {
	// 导出人

	Editor *string `json:"Editor,omitempty" name:"Editor"`
	// 导出节点信息

	Nodes *string `json:"Nodes,omitempty" name:"Nodes"`
}

type PathParams struct {
	// TdocReviewId

	TdocReviewId *int64 `json:"TdocReviewId,omitempty" name:"TdocReviewId"`
}

type QueryTreeUpgradeParams struct {
	// id

	TdocTreeId *uint64 `json:"TdocTreeId,omitempty" name:"TdocTreeId"`
	// 递归深度

	Depth *uint64 `json:"Depth,omitempty" name:"Depth"`
}

type EditDirRequest struct {
	*tchttp.BaseRequest

	// 无

	Params *EditDirParams `json:"Params,omitempty" name:"Params"`
}

func (r *EditDirRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditDirRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryNodeRequest struct {
	*tchttp.BaseRequest

	// 无

	Params *QueryNodeParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTreeUpgradeRequest struct {
	*tchttp.BaseRequest

	// input

	Params *QueryTreeUpgradeParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryTreeUpgradeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTreeUpgradeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddNodeRequest struct {
	*tchttp.BaseRequest

	// 无

	Params *AddNodeParams `json:"Params,omitempty" name:"Params"`
}

func (r *AddNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeableRequest struct {
	*tchttp.BaseRequest
}

func (r *UpgradeableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeableRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfirmUpgradeParams struct {
	// 操作人

	Editor *string `json:"Editor,omitempty" name:"Editor"`
	// 导入类型,值有:incremental or overwrite

	ImportType *string `json:"ImportType,omitempty" name:"ImportType"`
}

type QueryTreePreviewParams struct {
	// 文档树id

	TdocTreeId *uint64 `json:"TdocTreeId,omitempty" name:"TdocTreeId"`
	// 来源,值有tdoc_tree和tdoc_tree_upgrade

	Origin *string `json:"Origin,omitempty" name:"Origin"`
	// 深度

	Depth *uint64 `json:"Depth,omitempty" name:"Depth"`
	// 预览类型，值:incremental or overwrite

	PreviewType *string `json:"PreviewType,omitempty" name:"PreviewType"`
}

type ConfirmUpgradeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConfirmUpgradeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfirmUpgradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportDocResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportDocResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDirRequest struct {
	*tchttp.BaseRequest

	// 查询文档类型

	Params *QueryDirParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryDirRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryDirRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryNodePreviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryNodePreviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryNodePreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTreePreviewRequest struct {
	*tchttp.BaseRequest

	// 参数

	Params *QueryTreePreviewParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryTreePreviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTreePreviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateUpgradeLogRequest struct {
	*tchttp.BaseRequest

	// 参数

	Params *UpdateUpgradeLogParams `json:"Params,omitempty" name:"Params"`
}

func (r *UpdateUpgradeLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateUpgradeLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditNodeRequest struct {
	*tchttp.BaseRequest

	// 无

	Params *EditNodeParams `json:"Params,omitempty" name:"Params"`
}

func (r *EditNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddDirParams struct {
	// Action

	Action *string `json:"Action,omitempty" name:"Action"`
	// Alias

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// Editor

	Editor *string `json:"Editor,omitempty" name:"Editor"`
	// Image

	Image *string `json:"Image,omitempty" name:"Image"`
	// ParentId

	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`
	// Title

	Title *string `json:"Title,omitempty" name:"Title"`
	// Weight

	Weight *string `json:"Weight,omitempty" name:"Weight"`
	// Protocol

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type InsertOrUpdateNodeData struct {
	// 失败列表

	FailList []*Node `json:"FailList,omitempty" name:"FailList"`
	// 成功列表

	SuccessList []*Node `json:"SuccessList,omitempty" name:"SuccessList"`
}

type QueryExportLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryExportLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryExportLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTreeUpgradeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryTreeUpgradeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTreeUpgradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GiveupUpgradeParams struct {
	// 操作人

	Editor *string `json:"Editor,omitempty" name:"Editor"`
}

type PathResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PathResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PathResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDocProductCodesData struct {
	// 分类列表

	CategoryList []*CategoryData `json:"CategoryList,omitempty" name:"CategoryList"`
}

type QueryTreeRequest struct {
	*tchttp.BaseRequest

	// wu

	Params *QueryTreeParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryTreeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryTreeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelNodeData struct {
	// 文档树id，表示删除了这个树下的所有子孙节点

	TdocTreeId *uint64 `json:"TdocTreeId,omitempty" name:"TdocTreeId"`
}

type EditNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportDocRequest struct {
	*tchttp.BaseRequest

	// 参数

	Params *ExportDocParams `json:"Params,omitempty" name:"Params"`
}

func (r *ExportDocRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDocRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryUpgradeLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryUpgradeLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryUpgradeLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryNodeParams struct {
	// TdocTreeId

	TdocTreeId *int64 `json:"TdocTreeId,omitempty" name:"TdocTreeId"`
	// TdocReviewId

	TdocReviewId *int64 `json:"TdocReviewId,omitempty" name:"TdocReviewId"`
}

type CategoryData struct {
	// 类别名称(英文)

	CategoryEnName *string `json:"CategoryEnName,omitempty" name:"CategoryEnName"`
	// 类别名称

	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`
	// 产品列表

	ProductList []*Product `json:"ProductList,omitempty" name:"ProductList"`
}

type GiveupUpgradeRequest struct {
	*tchttp.BaseRequest

	// input

	Params *GiveupUpgradeParams `json:"Params,omitempty" name:"Params"`
}

func (r *GiveupUpgradeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GiveupUpgradeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Product struct {
	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品英文名称

	ProductEnName *string `json:"ProductEnName,omitempty" name:"ProductEnName"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type QueryDirParams struct {
	// Depth

	Depth *int64 `json:"Depth,omitempty" name:"Depth"`
	// Path

	Path *string `json:"Path,omitempty" name:"Path"`
	// TdocTreeId

	TdocTreeId *int64 `json:"TdocTreeId,omitempty" name:"TdocTreeId"`
}

type DelDirResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DelDirResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DelDirResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryUpgradeLogRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryUpgradeLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryUpgradeLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddDirRequest struct {
	*tchttp.BaseRequest

	// Params

	Params *AddDirParams `json:"Params,omitempty" name:"Params"`
}

func (r *AddDirRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddDirRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfirmUpgradeRequest struct {
	*tchttp.BaseRequest

	// input

	Params *ConfirmUpgradeParams `json:"Params,omitempty" name:"Params"`
}

func (r *ConfirmUpgradeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConfirmUpgradeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddNodeParams struct {
	// ParentId

	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`
	// Attr

	Attr *string `json:"Attr,omitempty" name:"Attr"`
	// Alias

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// Title

	Title *string `json:"Title,omitempty" name:"Title"`
	// Body

	Body *string `json:"Body,omitempty" name:"Body"`
	// Weight

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
	// Editor

	Editor *string `json:"Editor,omitempty" name:"Editor"`
}

type EditDirParams struct {
	// Action

	Action *string `json:"Action,omitempty" name:"Action"`
	// Alias

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// Editor

	Editor *string `json:"Editor,omitempty" name:"Editor"`
	// Image

	Image *string `json:"Image,omitempty" name:"Image"`
	// ParentId

	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`
	// TdocTreeId

	TdocTreeId *int64 `json:"TdocTreeId,omitempty" name:"TdocTreeId"`
	// Title

	Title *string `json:"Title,omitempty" name:"Title"`
	// Weight

	Weight *string `json:"Weight,omitempty" name:"Weight"`
	// Protocol

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// Deleted

	Deleted *int64 `json:"Deleted,omitempty" name:"Deleted"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

type ReviewParams struct {
	// a

	Title *string `json:"Title,omitempty" name:"Title"`
	// a

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// a

	Size *int64 `json:"Size,omitempty" name:"Size"`
	// a

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// a

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// aa

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// aa

	Editor *string `json:"Editor,omitempty" name:"Editor"`
}
