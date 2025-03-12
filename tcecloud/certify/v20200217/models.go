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

package v20200217

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type AllowCertificateApplyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllowCertificateApplyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllowCertificateApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCertificateApplyListParams struct {
	// 租户端APPID

	Appid *string `json:"Appid,omitempty" name:"Appid"`
	// 租户端UIN

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 租户端账号

	Account *string `json:"Account,omitempty" name:"Account"`
	// 租户端账号手机号码

	UinMobile *string `json:"UinMobile,omitempty" name:"UinMobile"`
	// 租户端账号客户名称

	UinNickname *string `json:"UinNickname,omitempty" name:"UinNickname"`
	// 联系人名称

	ContactName *string `json:"ContactName,omitempty" name:"ContactName"`
	// 手机号码

	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`
	// Pageid

	Pageid *uint64 `json:"Pageid,omitempty" name:"Pageid"`
	// Pagesize

	Pagesize *uint64 `json:"Pagesize,omitempty" name:"Pagesize"`
}

type UpdateCertificateStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCertificateStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCertificateStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCertificateStatusParams struct {
	// 取值"True/False"

	Status *string `json:"Status,omitempty" name:"Status"`
}

type AllowCertificateApplyParams struct {
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type RefuseCertificateApplyParams struct {
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 审核拒绝原因

	StatusReason *string `json:"StatusReason,omitempty" name:"StatusReason"`
}

type QueryCertificateDetailRequest struct {
	*tchttp.BaseRequest

	// 入参

	Params *QueryCertificateDetailParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryCertificateDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCertificateDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCertificateStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryCertificateStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCertificateStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCertificateStatusRequest struct {
	*tchttp.BaseRequest

	// input

	Params []*UpdateCertificateStatusParams `json:"Params,omitempty" name:"Params"`
}

func (r *UpdateCertificateStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCertificateStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCertificateApplyListRequest struct {
	*tchttp.BaseRequest

	// 无

	Params *QueryCertificateApplyListParams `json:"Params,omitempty" name:"Params"`
}

func (r *QueryCertificateApplyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCertificateApplyListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCertificateDetailParams struct {
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type AllowCertificateApplyRequest struct {
	*tchttp.BaseRequest

	// Input

	Params *AllowCertificateApplyParams `json:"Params,omitempty" name:"Params"`
}

func (r *AllowCertificateApplyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AllowCertificateApplyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCertificateApplyListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCertificateApplyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCertificateApplyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProvinceAndCityInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryProvinceAndCityInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProvinceAndCityInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefuseCertificateApplyRequest struct {
	*tchttp.BaseRequest

	// Input

	Params *RefuseCertificateApplyParams `json:"Params,omitempty" name:"Params"`
}

func (r *RefuseCertificateApplyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefuseCertificateApplyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProvinceAndCityInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryProvinceAndCityInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProvinceAndCityInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCertificateDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCertificateDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCertificateDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCertificateStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCertificateStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCertificateStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefuseCertificateApplyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RefuseCertificateApplyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RefuseCertificateApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
