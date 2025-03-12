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

package v20200119

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type RegionList struct {
	// 产品

	Products []*string `json:"Products,omitempty" name:"Products"`
	// 地域id

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域中文名称

	RegionNameCn *string `json:"RegionNameCn,omitempty" name:"RegionNameCn"`
	// 短地域名称

	RegionNameShort *string `json:"RegionNameShort,omitempty" name:"RegionNameShort"`
	// 地域状态

	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`
	// 可用区总数

	ZoneCount *int64 `json:"ZoneCount,omitempty" name:"ZoneCount"`
	// 可用区列表

	ZoneList []*ZoneList `json:"ZoneList,omitempty" name:"ZoneList"`
}

type DetailInfo struct {
	// 地域id

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 可用区列表

	Zones []*Zones `json:"Zones,omitempty" name:"Zones"`
}

type GetRegionZoneIdcInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *GetRegionZoneIdcInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRegionZoneIdcInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRegionZoneIdcInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		RegionCount *int64 `json:"RegionCount,omitempty" name:"RegionCount"`
		// 可用区列表信息

		RegionList []*RegionList `json:"RegionList,omitempty" name:"RegionList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRegionZoneIdcInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRegionZoneIdcInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetZonesInfoRequest struct {
	*tchttp.BaseRequest

	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

func (r *GetZonesInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetZonesInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetZonesInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可用区信息

		DetailInfo *DetailInfo `json:"DetailInfo,omitempty" name:"DetailInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetZonesInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetZonesInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IdcList struct {
	// IdcId

	IdcId *int64 `json:"IdcId,omitempty" name:"IdcId"`
	// IdcName

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
	// IdcNameCn

	IdcNameCn *string `json:"IdcNameCn,omitempty" name:"IdcNameCn"`
	// IdcState

	IdcState *string `json:"IdcState,omitempty" name:"IdcState"`
}

type ZoneList struct {
	// IDC 总数

	IdcCount *int64 `json:"IdcCount,omitempty" name:"IdcCount"`
	// IDC列表

	IdcList []*IdcList `json:"IdcList,omitempty" name:"IdcList"`
	// 可用区id

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 可用区中文名称

	ZoneNameCn *string `json:"ZoneNameCn,omitempty" name:"ZoneNameCn"`
	// 可用区状态

	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`
}

type Zones struct {
	// 可用区id

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 可用区中文名称

	ZoneNameCn *string `json:"ZoneNameCn,omitempty" name:"ZoneNameCn"`
	// 可用区状态

	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`
}
