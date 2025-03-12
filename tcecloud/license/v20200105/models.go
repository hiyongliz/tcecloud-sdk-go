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

package v20200105

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type GetProductEditionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 交付中心产品名称

		ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
		// 产品功能版本

		ProductEdition *string `json:"ProductEdition,omitempty" name:"ProductEdition"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductEditionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductEditionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Product struct {
	// 交付中心产品名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品版本

	Edition *string `json:"Edition,omitempty" name:"Edition"`
	// 售卖单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 售卖数量

	Value *int64 `json:"Value,omitempty" name:"Value"`
	// //资源信息: 多个资源id以英文分号;隔开

	ResourceInfo *string `json:"ResourceInfo,omitempty" name:"ResourceInfo"`
	//  //有效期：具体可用截止日期或售卖时长

	ValidPeriod *string `json:"ValidPeriod,omitempty" name:"ValidPeriod"`
	// //有效期模式： 0 可用截止日期，1 售卖时长

	ValidPeriodType *string `json:"ValidPeriodType,omitempty" name:"ValidPeriodType"`
	// // 计费项

	BillingItem *string `json:"BillingItem,omitempty" name:"BillingItem"`
	// // 计费项ID

	BillingItemId *string `json:"BillingItemId,omitempty" name:"BillingItemId"`
	// // 本地计费项ID

	LicenseDefinitionUUID *string `json:"LicenseDefinitionUUID,omitempty" name:"LicenseDefinitionUUID"`
	// // 计费项虚拟ID

	LicenseID *string `json:"LicenseID,omitempty" name:"LicenseID"`
}

type Usage struct {
	// 交付中心产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 用量单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 用量信息

	Value *string `json:"Value,omitempty" name:"Value"`
}

type License struct {
	// 项目标识: LTC 项目编号

	ProjectID *string `json:"ProjectID,omitempty" name:"ProjectID"`
	// 授权类型: COMMERCIA 或 POC

	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`
	// 生效日期: 格式: 2006-01-02

	EffectiveDate *string `json:"EffectiveDate,omitempty" name:"EffectiveDate"`
	// 失效日期: 格式: 2006-01-02

	ExpireDate *string `json:"ExpireDate,omitempty" name:"ExpireDate"`
	// 交付中心: 交付编号

	CloudID *string `json:"CloudID,omitempty" name:"CloudID"`
	// 交付中心: 交付名称

	CloudName *string `json:"CloudName,omitempty" name:"CloudName"`
	// 交付中心: 产品规格列表

	Products []*Product `json:"Products,omitempty" name:"Products"`
	// 唯一标识: 授权编号

	LicenseID *string `json:"LicenseID,omitempty" name:"LicenseID"`
	// 客户信息: 客户名称

	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`
	// // 有效期模式： 0 可用截止日期，1 售卖时长

	ValidPeriod *string `json:"ValidPeriod,omitempty" name:"ValidPeriod"`
	// // 产品规范: 校验规则

	Rules []*Rule `json:"Rules,omitempty" name:"Rules"`
	// // 项目售卖模式: 0 长期， 1 订阅

	SellType *string `json:"SellType,omitempty" name:"SellType"`
	// // 类型：测试: TEST, 正式: OFFICIAL

	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`
	// // 专有云版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// // 有效期模式： 0 可用截止日期，1 售卖时长

	ValidPeriodType *string `json:"ValidPeriodType,omitempty" name:"ValidPeriodType"`
}

type GetRealTimeClockResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务端实时时钟

		RealTimeClock *string `json:"RealTimeClock,omitempty" name:"RealTimeClock"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRealTimeClockResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRealTimeClockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LicenseDetailRequest struct {
	*tchttp.BaseRequest
}

func (r *LicenseDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LicenseDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsageFilterEntity struct {
	// 状态列表

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 采集状态列表

	CollectStatus []*string `json:"CollectStatus,omitempty" name:"CollectStatus"`
	// 单位集合

	Unit []*string `json:"Unit,omitempty" name:"Unit"`
}

type DescribeLicenseStatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品总数量

		ProductCount *uint64 `json:"ProductCount,omitempty" name:"ProductCount"`
		// 超用总数量

		OveruseProductCount *uint64 `json:"OveruseProductCount,omitempty" name:"OveruseProductCount"`
		// 异常产品数量

		AbnormalProductCount *uint64 `json:"AbnormalProductCount,omitempty" name:"AbnormalProductCount"`
		// 采集产品总数

		CollectProductCount *uint64 `json:"CollectProductCount,omitempty" name:"CollectProductCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLicenseStatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRealTimeClockRequest struct {
	*tchttp.BaseRequest
}

func (r *GetRealTimeClockRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRealTimeClockRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LicenseUpdateRequest struct {
	*tchttp.BaseRequest

	// TCE 颁发的授权文件 ( base64编码 )

	License *string `json:"License,omitempty" name:"License"`
}

func (r *LicenseUpdateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LicenseUpdateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductBilling struct {
	// 计费项ID

	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
	// 计费项名称

	ProductIdDesc *string `json:"ProductIdDesc,omitempty" name:"ProductIdDesc"`
	// 已购买数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// license计费项导入后的生效时间

	EffectiveTime *string `json:"EffectiveTime,omitempty" name:"EffectiveTime"`
	// license计费项导入后的过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type GetProductEditionRequest struct {
	*tchttp.BaseRequest

	// 交付中心产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

func (r *GetProductEditionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductEditionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductUsagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 产品用量列表

		Products []*ProductUsage `json:"Products,omitempty" name:"Products"`
		// 筛选参数

		FilterEntity *UsageFilterEntity `json:"FilterEntity,omitempty" name:"FilterEntity"`
		// 是否在试用期内，true: 是，false: 否

		IsProbation *bool `json:"IsProbation,omitempty" name:"IsProbation"`
		// 试用期状态, Normal:正常, Overdue: 超期

		ProbationStatus *string `json:"ProbationStatus,omitempty" name:"ProbationStatus"`
		// license包级别有效期

		ValidityPeriod *string `json:"ValidityPeriod,omitempty" name:"ValidityPeriod"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductUsagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductUsagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Rule struct {
	// 校验项标识

	Key *string `json:"Key,omitempty" name:"Key"`
	// 校验表达式

	Expression *string `json:"Expression,omitempty" name:"Expression"`
	// 值类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 校验值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeLicenseStatRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLicenseStatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLicenseStatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductUsagesRequest struct {
	*tchttp.BaseRequest

	// 每页个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 抓取开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 抓取结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 排序字段

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 是否倒序

	Desc *bool `json:"Desc,omitempty" name:"Desc"`
	// 计费项id

	BillItemID *string `json:"BillItemID,omitempty" name:"BillItemID"`
}

func (r *ListProductUsagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductUsagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LicenseDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 授权文件内容

		License *License `json:"License,omitempty" name:"License"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LicenseDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LicenseDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LicenseUpdateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// License 授权文件的校验和

		Checksum *string `json:"Checksum,omitempty" name:"Checksum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LicenseUpdateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LicenseUpdateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 值列表

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 操作符，枚举值，equal:等于, like: 模糊查找

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type ProductUsage struct {
	// 产品名

	Product *string `json:"Product,omitempty" name:"Product"`
	// 计费项，多项以逗号分割

	BillingDesc *string `json:"BillingDesc,omitempty" name:"BillingDesc"`
	// 已购买数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 已使用数量

	UsedCount *string `json:"UsedCount,omitempty" name:"UsedCount"`
	// 使用率

	Usage *float64 `json:"Usage,omitempty" name:"Usage"`
	// 状态，Normal:正常，Overused: 超用

	Status *string `json:"Status,omitempty" name:"Status"`
	// 抓取状态，Normal:正常，Abnormal:异常 正常

	CollectStatus *string `json:"CollectStatus,omitempty" name:"CollectStatus"`
	// 正常抓取时间

	CollectTime *string `json:"CollectTime,omitempty" name:"CollectTime"`
	// 子计费项列表

	Billings []*ProductBilling `json:"Billings,omitempty" name:"Billings"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 所有异常状态列表

	AbnormalStatusList []*string `json:"AbnormalStatusList,omitempty" name:"AbnormalStatusList"`
	// 计费项英文名称，当前不支持

	BillItemEnName *string `json:"BillItemEnName,omitempty" name:"BillItemEnName"`
	// 过期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 是否为合并计费，true: 是, false:否

	IsCombinedBilling *bool `json:"IsCombinedBilling,omitempty" name:"IsCombinedBilling"`
	// 产品绑定资源信息

	ResourceInfo *string `json:"ResourceInfo,omitempty" name:"ResourceInfo"`
	// 生效时间

	EffectiveTime *string `json:"EffectiveTime,omitempty" name:"EffectiveTime"`
	// 计费项ID

	BillingItemId *string `json:"BillingItemId,omitempty" name:"BillingItemId"`
	// 首次超用时间

	FirstOverUsedTime *string `json:"FirstOverUsedTime,omitempty" name:"FirstOverUsedTime"`
}
