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

package v20200501

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type FilterUnit struct {
	// 字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// 操作符

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type InfoUnit struct {
	// 命名空间名称

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 指标英文名称

	Fields []*string `json:"Fields,omitempty" name:"Fields"`
}

type GetAggreDataRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 截止时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 请求业务监控用的module

	Module *string `json:"Module,omitempty" name:"Module"`
	// 目标ip

	IpList []*string `json:"IpList,omitempty" name:"IpList"`
	// 筛选条件

	Filter []*FilterUnit `json:"Filter,omitempty" name:"Filter"`
	// 分页

	Page *Page `json:"Page,omitempty" name:"Page"`
	// 监控命名空间指标信息

	Info []*InfoUnit `json:"Info,omitempty" name:"Info"`
}

func (r *GetAggreDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAggreDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAggreDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAggreDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAggreDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Page struct {
	// 起始值

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}
