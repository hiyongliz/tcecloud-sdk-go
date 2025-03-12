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

package v20200113

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type UsageDetailRequest struct {
	*tchttp.BaseRequest
}

func (r *UsageDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UsageDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Detail struct {
	// 地域编号

	RegionID *int64 `json:"RegionID,omitempty" name:"RegionID"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 可用区编号

	ZoneID *int64 `json:"ZoneID,omitempty" name:"ZoneID"`
	// 可用区名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 交付中心: 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 用量单位

	UsageUnit *string `json:"UsageUnit,omitempty" name:"UsageUnit"`
	// 用量

	UsageValue *int64 `json:"UsageValue,omitempty" name:"UsageValue"`
}

type UsageDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品用量信息列表

		Usages []*Usage `json:"Usages,omitempty" name:"Usages"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UsageDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UsageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Usage struct {
	// 交付中心: 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 用量单位

	UsageUnit *string `json:"UsageUnit,omitempty" name:"UsageUnit"`
	// 用量总和

	UsageValue *int64 `json:"UsageValue,omitempty" name:"UsageValue"`
	// Region/Zone 用量详情

	Details []*Detail `json:"Details,omitempty" name:"Details"`
}
