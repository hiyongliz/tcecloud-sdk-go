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

package v20181115

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type CreateAnnounceLinkRequest struct {
	*tchttp.BaseRequest

	// 公告链接url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 公告链接标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 公告链接状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 有效起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 有效结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *CreateAnnounceLinkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAnnounceLinkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAnnounceClassRequest struct {
	*tchttp.BaseRequest

	// 公告分类id，创建或查询时返回的ClassId

	ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`
	// 公告分类中文名称

	NameCn *string `json:"NameCn,omitempty" name:"NameCn"`
	// 公告分类级别

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAnnounceClassRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAnnounceClassRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAnnounceContentRequest struct {
	*tchttp.BaseRequest

	// 公告标题，建议用base64编码

	Toptic *string `json:"Toptic,omitempty" name:"Toptic"`
	// 公告内容，建议用base64编码

	Content *string `json:"Content,omitempty" name:"Content"`
	// 机房地域id列表，多个用;号隔开

	CityIds *string `json:"CityIds,omitempty" name:"CityIds"`
	// 公告分类id列表，多个用;号隔开

	ClassIds *string `json:"ClassIds,omitempty" name:"ClassIds"`
	// 公告状态信息

	Status *string `json:"Status,omitempty" name:"Status"`
	// 公告生效起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 公告生效结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 发布时间，默认为数据入库时间

	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
	// 图片名称标识

	ImgMark *string `json:"ImgMark,omitempty" name:"ImgMark"`
}

func (r *CreateAnnounceContentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAnnounceContentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAnnounceContentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公告唯一ID

		AnnId *uint64 `json:"AnnId,omitempty" name:"AnnId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAnnounceContentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAnnounceContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAnnounceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态，修改成功为true

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAnnounceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAnnounceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAnnounceLinkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公告链接数组

		AnnounceLinkSet []*AnnounceLink `json:"AnnounceLinkSet,omitempty" name:"AnnounceLinkSet"`
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAnnounceLinkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAnnounceLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAnnounceListRequest struct {
	*tchttp.BaseRequest

	// 公告分类id列表，多个用;号隔开

	ClassIds *string `json:"ClassIds,omitempty" name:"ClassIds"`
	// 地域id列表，多个用;号隔开

	CityIds *string `json:"CityIds,omitempty" name:"CityIds"`
	// 公告状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 生效起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 生效结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 发布起始时间

	StartPublishTime *string `json:"StartPublishTime,omitempty" name:"StartPublishTime"`
	// 发布结束时间

	EndPublishTime *string `json:"EndPublishTime,omitempty" name:"EndPublishTime"`
}

func (r *GetAnnounceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAnnounceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAnnounceLinkRequest struct {
	*tchttp.BaseRequest

	// 公告链接id

	LinkId *uint64 `json:"LinkId,omitempty" name:"LinkId"`
	// 公告链接url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 公告链接标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 有效起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 有效结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ModifyAnnounceLinkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAnnounceLinkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Annclass struct {
	// 公告分类中文名称

	NameCn *string `json:"NameCn,omitempty" name:"NameCn"`
	// 公告分类级别

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 公告分类唯一id

	ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DeleteAnnounceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 受影响的行

		AffectedRows *uint64 `json:"AffectedRows,omitempty" name:"AffectedRows"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAnnounceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAnnounceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAnnounceLinkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公告链接唯一ID

		LinkId *uint64 `json:"LinkId,omitempty" name:"LinkId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAnnounceLinkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAnnounceLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAnnounceClassResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公告分类级别，从0递增,可用于查询时排序

		ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAnnounceClassResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAnnounceClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAnnounceClassRequest struct {
	*tchttp.BaseRequest

	// 公告分类唯一id

	ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`
	// 公告分类级别，用于查询某一级别的告警分类

	Level *uint64 `json:"Level,omitempty" name:"Level"`
	// 以level的升序还是降序的方式返回结果，可选值 ASC 和 DESC

	Order *string `json:"Order,omitempty" name:"Order"`
	// 偏移量，默认为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认返回所有

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *GetAnnounceClassRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAnnounceClassRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAnnounceClassRequest struct {
	*tchttp.BaseRequest

	// 公告分类中文名称如: 运维

	NameCn *string `json:"NameCn,omitempty" name:"NameCn"`
	// 公告分类级别，从0递增,可用于查询时排序

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *CreateAnnounceClassRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAnnounceClassRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHostAreaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 受影响的行

		AffectedRows *uint64 `json:"AffectedRows,omitempty" name:"AffectedRows"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteHostAreaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHostAreaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHostAreaRequest struct {
	*tchttp.BaseRequest

	// 机房地域唯一id

	CityId *uint64 `json:"CityId,omitempty" name:"CityId"`
}

func (r *DeleteHostAreaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHostAreaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAnnounceClassResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公告分类列表数组

		AnnclassSet []*Annclass `json:"AnnclassSet,omitempty" name:"AnnclassSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAnnounceClassResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAnnounceClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAnnounceContentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公告标题

		Toptic *string `json:"Toptic,omitempty" name:"Toptic"`
		// 地域id列表，多个用;号隔开

		CityIds *string `json:"CityIds,omitempty" name:"CityIds"`
		// 公告分类id

		ClassId *string `json:"ClassId,omitempty" name:"ClassId"`
		// 公告状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 生效起始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 生效结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 发布时间

		PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
		// 公告内容

		Content *string `json:"Content,omitempty" name:"Content"`
		// 图片文件标识

		ImgMark *string `json:"ImgMark,omitempty" name:"ImgMark"`
		// 公告ID

		AnnId *uint64 `json:"AnnId,omitempty" name:"AnnId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAnnounceContentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAnnounceContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnnounceLink struct {
	// 公告链接ID

	LinkId *uint64 `json:"LinkId,omitempty" name:"LinkId"`
	// 公告链接标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 链接Url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 有效起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 有效结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 创建时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
}

type CreateHostAreaRequest struct {
	*tchttp.BaseRequest

	// 地域中文名称

	CityCn *string `json:"CityCn,omitempty" name:"CityCn"`
	// 国家英文名称

	Country *string `json:"Country,omitempty" name:"Country"`
	// 区域英文名称

	CityRegion *string `json:"CityRegion,omitempty" name:"CityRegion"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *CreateHostAreaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHostAreaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostAreaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域信息列表数组

		HostAreaSet []*HostArea `json:"HostAreaSet,omitempty" name:"HostAreaSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetHostAreaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostAreaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostAreaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态标识

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostAreaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostAreaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHostAreaRequest struct {
	*tchttp.BaseRequest

	// 地域唯一id

	CityId *uint64 `json:"CityId,omitempty" name:"CityId"`
	// 按国家查询

	Country *string `json:"Country,omitempty" name:"Country"`
	// 按区域查询

	CityRegion *string `json:"CityRegion,omitempty" name:"CityRegion"`
	// 按Country排序，asc升序，desc降序

	OrderCountry *string `json:"OrderCountry,omitempty" name:"OrderCountry"`
	// 按Region排序，asc升序，desc降序

	OrderRegion *string `json:"OrderRegion,omitempty" name:"OrderRegion"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认返回所有

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *GetHostAreaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHostAreaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAnnounceClassRequest struct {
	*tchttp.BaseRequest

	// 公告分类id

	ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`
}

func (r *DeleteAnnounceClassRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAnnounceClassRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAnnounceClassResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 受影响的行

		AffectedRows *uint64 `json:"AffectedRows,omitempty" name:"AffectedRows"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAnnounceClassResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAnnounceClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAnnounceContentRequest struct {
	*tchttp.BaseRequest

	// 公告唯一ID

	AnnId *uint64 `json:"AnnId,omitempty" name:"AnnId"`
}

func (r *GetAnnounceContentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAnnounceContentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostArea struct {
	// 地域唯一id

	CityId *uint64 `json:"CityId,omitempty" name:"CityId"`
	// 地域中文名称

	CityCn *string `json:"CityCn,omitempty" name:"CityCn"`
	// 所在国家

	Country *string `json:"Country,omitempty" name:"Country"`
	// 所有区域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyAnnounceClassResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态，修改成功为true

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAnnounceClassResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAnnounceClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAnnounceLinkRequest struct {
	*tchttp.BaseRequest

	// 公告链接ID

	LinkId *uint64 `json:"LinkId,omitempty" name:"LinkId"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 公告链接Url

	Url *string `json:"Url,omitempty" name:"Url"`
	// 有效起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 有效结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 数据限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetAnnounceLinkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAnnounceLinkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAnnounceRequest struct {
	*tchttp.BaseRequest

	// 公告唯一ID

	AnnId *uint64 `json:"AnnId,omitempty" name:"AnnId"`
	// 公告标题

	Toptic *string `json:"Toptic,omitempty" name:"Toptic"`
	// 公告内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 地域id列表，多个用;号隔开

	CityIds *string `json:"CityIds,omitempty" name:"CityIds"`
	// 公告分类id列表，多个用;号隔开

	ClassIds *string `json:"ClassIds,omitempty" name:"ClassIds"`
	// 公告状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 生效起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 生效结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 发布时间

	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
	// 图片文件标识

	ImgMark *string `json:"ImgMark,omitempty" name:"ImgMark"`
}

func (r *ModifyAnnounceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAnnounceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Announce struct {
	// 公告唯一id

	AnnId *uint64 `json:"AnnId,omitempty" name:"AnnId"`
	// 公告标题，建议用base64

	Toptic *string `json:"Toptic,omitempty" name:"Toptic"`
	// 地域id列表，多个用;号隔开

	CityIds *string `json:"CityIds,omitempty" name:"CityIds"`
	// 公告分类id

	ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`
	// 公告状态信息

	Status *string `json:"Status,omitempty" name:"Status"`
	// 生效起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 生效结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 发布时间

	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
	// 图片文件标识

	ImgMark *string `json:"ImgMark,omitempty" name:"ImgMark"`
}

type GetAnnounceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公告列表，默认按照发布时间倒序排序

		AnnounceSet []*Announce `json:"AnnounceSet,omitempty" name:"AnnounceSet"`
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAnnounceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAnnounceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHostAreaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域信息唯一ID

		CityId *uint64 `json:"CityId,omitempty" name:"CityId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHostAreaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHostAreaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHostAreaRequest struct {
	*tchttp.BaseRequest

	// 地域信息唯一ID

	CityId *uint64 `json:"CityId,omitempty" name:"CityId"`
	// 地域名称

	CityCn *string `json:"CityCn,omitempty" name:"CityCn"`
	// 国家名称

	Country *string `json:"Country,omitempty" name:"Country"`
	// 区域信息

	CityRegion *string `json:"CityRegion,omitempty" name:"CityRegion"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyHostAreaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHostAreaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAnnounceRequest struct {
	*tchttp.BaseRequest

	// 公告唯一id

	AnnId *uint64 `json:"AnnId,omitempty" name:"AnnId"`
}

func (r *DeleteAnnounceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAnnounceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAnnounceLinkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 状态，修改成功为0

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAnnounceLinkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAnnounceLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
