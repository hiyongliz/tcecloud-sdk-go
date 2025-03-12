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

package v20231005

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type DeleteUserProductMenusRequest struct {
	*tchttp.BaseRequest

	// 目标用户uin

	TargetUserUin *uint64 `json:"TargetUserUin,omitempty" name:"TargetUserUin"`
	// 目标用户名

	TargetUserName *string `json:"TargetUserName,omitempty" name:"TargetUserName"`
}

func (r *DeleteUserProductMenusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserProductMenusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserProductMenusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserProductMenusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserProductMenusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserProductMenusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 菜单列表

		DataList []*string `json:"DataList,omitempty" name:"DataList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserProductMenusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserProductMenusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyUserProductMenusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyUserProductMenusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyUserProductMenusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserProductMenusRequest struct {
	*tchttp.BaseRequest

	// 用户 uin

	TargetUserUin *uint64 `json:"TargetUserUin,omitempty" name:"TargetUserUin"`
}

func (r *DescribeUserProductMenusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserProductMenusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductMenuItem struct {
	// 产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 菜单uuid

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type AddUserProductMenusRequest struct {
	*tchttp.BaseRequest

	// 目标用户uin

	TargetUserUin *uint64 `json:"TargetUserUin,omitempty" name:"TargetUserUin"`
	// 目标用户名

	TargetUserName *string `json:"TargetUserName,omitempty" name:"TargetUserName"`
	// 产品菜单

	ProductMenus []*ProductMenuItem `json:"ProductMenus,omitempty" name:"ProductMenus"`
}

func (r *AddUserProductMenusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserProductMenusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddUserProductMenusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUserProductMenusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserProductMenusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserProductMenusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserProductMenusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserProductMenusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserProductMenusRequest struct {
	*tchttp.BaseRequest

	// 目标用户uin

	TargetUserUin *uint64 `json:"TargetUserUin,omitempty" name:"TargetUserUin"`
	// 目标用户名

	TargetUserName *string `json:"TargetUserName,omitempty" name:"TargetUserName"`
	// 产品菜单

	ProductMenus []*ProductMenuItem `json:"ProductMenus,omitempty" name:"ProductMenus"`
}

func (r *ModifyUserProductMenusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserProductMenusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyUserProductMenusRequest struct {
	*tchttp.BaseRequest

	// 产品菜单uuid

	ProductMenuUuid *string `json:"ProductMenuUuid,omitempty" name:"ProductMenuUuid"`
}

func (r *VerifyUserProductMenusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyUserProductMenusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
