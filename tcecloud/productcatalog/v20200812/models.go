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

package v20200812

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type GetGrayscaleListByUuidRequest struct {
	*tchttp.BaseRequest

	// 产品uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *GetGrayscaleListByUuidRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGrayscaleListByUuidRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserProductCatalogsRequest struct {
	*tchttp.BaseRequest

	// 租户Appid

	Appid *int64 `json:"Appid,omitempty" name:"Appid"`
	// 产品目录

	ProductCatalogs []*ProductCatalogs `json:"ProductCatalogs,omitempty" name:"ProductCatalogs"`
}

func (r *ModifyUserProductCatalogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserProductCatalogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserProductCatalogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserProductCatalogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserProductCatalogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataList struct {
	// 租户账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 租户AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

type ProductCatalogInfo struct {
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 产品状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type AddUserProductCatalogsRequest struct {
	*tchttp.BaseRequest

	// 租户appId

	Appid *int64 `json:"Appid,omitempty" name:"Appid"`
	// 租户账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 产品目录

	ProductCatalogs []*ProductCatalogs `json:"ProductCatalogs,omitempty" name:"ProductCatalogs"`
}

func (r *AddUserProductCatalogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserProductCatalogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserProductCatalogsList struct {
	// 记录自增id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 租户appId

	Appid *int64 `json:"Appid,omitempty" name:"Appid"`
	// 租户账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 创建人uin

	CreaterUin *int64 `json:"CreaterUin,omitempty" name:"CreaterUin"`
	// 创建人名称

	CreaterName *string `json:"CreaterName,omitempty" name:"CreaterName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 产品目录列表

	ProductCatalogs []*string `json:"ProductCatalogs,omitempty" name:"ProductCatalogs"`
}

type AddUserProductCatalogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUserProductCatalogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserProductCatalogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserProductCatalogsRequest struct {
	*tchttp.BaseRequest

	// 租户Appid

	Appid *int64 `json:"Appid,omitempty" name:"Appid"`
}

func (r *DeleteUserProductCatalogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserProductCatalogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProductStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserProductCatalogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserProductCatalogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserProductCatalogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetGrayscaleListByUuidResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 租户信息列表

		DataList []*DataList `json:"DataList,omitempty" name:"DataList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGrayscaleListByUuidResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetGrayscaleListByUuidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUserProductCatalogsRequest struct {
	*tchttp.BaseRequest

	// 租户Appid

	Appid *int64 `json:"Appid,omitempty" name:"Appid"`
	// 租户账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 偏移量第一页为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页显示条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListUserProductCatalogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUserProductCatalogsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductCatalogInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 产品目录信息

		DataList *ProductCatalogInfo `json:"DataList,omitempty" name:"DataList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductCatalogInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductCatalogInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductCatalogInfoRequest struct {
	*tchttp.BaseRequest

	// "0": "已上架", "1": "灰度上架", "2":"未上架"

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 返回数据个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetProductCatalogInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductCatalogInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUserProductCatalogsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 定义了产品目录的租户列表

		DataList []*UserProductCatalogsList `json:"DataList,omitempty" name:"DataList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUserProductCatalogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListUserProductCatalogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductStatusRequest struct {
	*tchttp.BaseRequest

	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品Uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// "0": "已上架", "1": "灰度上架", "2":"未上架"

	Status *string `json:"Status,omitempty" name:"Status"`
	// 灰度名单 例:["1","2"]

	GrayscaleAppIdList []*string `json:"GrayscaleAppIdList,omitempty" name:"GrayscaleAppIdList"`
}

func (r *UpdateProductStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductCatalogs struct {
	// 产品uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}
