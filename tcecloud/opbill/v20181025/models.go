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

package v20181025

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type AddComprehensiveStrategyRouteGatewayRequest struct {
	*tchttp.BaseRequest

	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品code

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 新购是否同步发货 默认0

	PurchaseSyncDelivery *uint64 `json:"PurchaseSyncDelivery,omitempty" name:"PurchaseSyncDelivery"`
	// 变配是否同步发货 默认0

	ModifySyncDelivery *uint64 `json:"ModifySyncDelivery,omitempty" name:"ModifySyncDelivery"`
	// 接口名

	InterfaceNames *InterfaceNames `json:"InterfaceNames,omitempty" name:"InterfaceNames"`
	// url名

	Region2interface *string `json:"Region2interface,omitempty" name:"Region2interface"`
	// 添加的是计量产品还是计费产品，取值范围计量（bsp）计费（billing），不填写则默认bsp

	Source *string `json:"Source,omitempty" name:"Source"`
}

func (r *AddComprehensiveStrategyRouteGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddComprehensiveStrategyRouteGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReconciliationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyReconciliationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReconciliationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubProductsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 子产品信息

		List []*SubProductDefineList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubProductsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLifecycleFlowList struct {
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 更改时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 请求

	Request *string `json:"Request,omitempty" name:"Request"`
	// 账号UIN

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 产品名称

	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
	// 操作流程名称

	OperateActionName *string `json:"OperateActionName,omitempty" name:"OperateActionName"`
	// 操作者UIN

	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 操作流程

	OperateAction *string `json:"OperateAction,omitempty" name:"OperateAction"`
}

type CreateDownloadRecordRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 文件ID

	FileId *string `json:"FileId,omitempty" name:"FileId"`
}

func (r *CreateDownloadRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDownloadRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByProductRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillSummaryByProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInvoiceDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户开票流水号

		OpenId *int64 `json:"OpenId,omitempty" name:"OpenId"`
		// 发票金额

		Amount *int64 `json:"Amount,omitempty" name:"Amount"`
		// 申请开票时间

		ApplicationTime *string `json:"ApplicationTime,omitempty" name:"ApplicationTime"`
		// 发票类型,取值为(  2:公司普通发票  3:公司增值税发票 )

		UserType *int64 `json:"UserType,omitempty" name:"UserType"`
		// 发票抬头

		InvoiceHead *string `json:"InvoiceHead,omitempty" name:"InvoiceHead"`
		// 税务登记号

		TaxNumber *string `json:"TaxNumber,omitempty" name:"TaxNumber"`
		// 开户行

		BankUserName *string `json:"BankUserName,omitempty" name:"BankUserName"`
		// 银行账号

		AccountNumber *string `json:"AccountNumber,omitempty" name:"AccountNumber"`
		// 注册全地

		RegisterFullAddr *string `json:"RegisterFullAddr,omitempty" name:"RegisterFullAddr"`
		// 注册电话

		RegisterPhone *string `json:"RegisterPhone,omitempty" name:"RegisterPhone"`
		// 地址

		Addr *string `json:"Addr,omitempty" name:"Addr"`
		// 联系人

		Contact *string `json:"Contact,omitempty" name:"Contact"`
		// 手机号码

		Cellphone *string `json:"Cellphone,omitempty" name:"Cellphone"`
		// 发票状态,1:处理中,6:已取消 30:审核通过  31:审核不通过

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 发票备注

		InvoiceRemark *string `json:"InvoiceRemark,omitempty" name:"InvoiceRemark"`
		// 城市

		City *string `json:"City,omitempty" name:"City"`
		// 省份

		Province *string `json:"Province,omitempty" name:"Province"`
		// 账期

		AccountPeriods *string `json:"AccountPeriods,omitempty" name:"AccountPeriods"`
		// 审核操作人

		Auditor *string `json:"Auditor,omitempty" name:"Auditor"`
		// 审核人拒绝原因

		Reason *string `json:"Reason,omitempty" name:"Reason"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInvoiceDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInvoiceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseResourceGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseResourceGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseResourceGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductGatewayRequest struct {
	*tchttp.BaseRequest

	// 产品大类中文名

	ProductGroupName *string `json:"ProductGroupName,omitempty" name:"ProductGroupName"`
	// 产品大类英文名

	ProductGroupEngName *string `json:"ProductGroupEngName,omitempty" name:"ProductGroupEngName"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品英文名称

	ProductEngName *string `json:"ProductEngName,omitempty" name:"ProductEngName"`
	// 子产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品名称

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 子产品英文名称

	SubProductEngName *string `json:"SubProductEngName,omitempty" name:"SubProductEngName"`
	// 计费项code

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费项名称

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 计费项英文名称

	BillingItemEngName *string `json:"BillingItemEngName,omitempty" name:"BillingItemEngName"`
	// 计费细项代码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 计费细项名称

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
	// 计费细项英文名

	SubBillingItemEngName *string `json:"SubBillingItemEngName,omitempty" name:"SubBillingItemEngName"`
	// 单位中文名

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 单位英文名

	UnitEng *string `json:"UnitEng,omitempty" name:"UnitEng"`
	// 计费周期

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 结算周期

	CalcUnit *string `json:"CalcUnit,omitempty" name:"CalcUnit"`
	// 主键ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *UpdateProductGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddQuotaPara struct {
	// 用户uin。此参数不传时设置默认配额。

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 配额key

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
	// 配额值

	QuotaValue *string `json:"QuotaValue,omitempty" name:"QuotaValue"`
}

type GetDownloadListRecordItem struct {
	// 文件ID

	FileId *string `json:"FileId,omitempty" name:"FileId"`
	// 文件类型

	FileType *string `json:"FileType,omitempty" name:"FileType"`
	// 文件类型名称

	FileTypeName *string `json:"FileTypeName,omitempty" name:"FileTypeName"`
	// 文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 生成时间

	GenerateTime *string `json:"GenerateTime,omitempty" name:"GenerateTime"`
	// 下载时间

	DownloadTime *string `json:"DownloadTime,omitempty" name:"DownloadTime"`
	// 下载地址

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
}

type DescribeBillTrendByMonthGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 每月花费详情

		Detail []*TrendByMonthDetail `json:"Detail,omitempty" name:"Detail"`
		// 平均花费详情

		Stat []*TrendByMonthStat `json:"Stat,omitempty" name:"Stat"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillTrendByMonthGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillTrendByMonthGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContractListRequest struct {
	*tchttp.BaseRequest

	// 管理端查询用户UIN，默认缺省查全部

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 合同编号

	ContractId *uint64 `json:"ContractId,omitempty" name:"ContractId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 合同类型ID号

	ContractTypeId *uint64 `json:"ContractTypeId,omitempty" name:"ContractTypeId"`
	// 创建时间排序正序asc 倒序desc

	CreateTimeOrder *string `json:"CreateTimeOrder,omitempty" name:"CreateTimeOrder"`
	// 合同状态0审核中1已邮寄2审核拒绝3作废4取消

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeContractListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContractListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteQuotaGatewayRequest struct {
	*tchttp.BaseRequest

	// 商品码，商品码和QuotaKey不能同时为空

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 配额key

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
	// uin

	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`
}

func (r *DeleteQuotaGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteQuotaGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadVoucherListFileGatewayRequest struct {
	*tchttp.BaseRequest

	// 代金券筛选条件对象

	VoucherConditions *VoucherConditions `json:"VoucherConditions,omitempty" name:"VoucherConditions"`
}

func (r *DescribeDownloadVoucherListFileGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadVoucherListFileGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyComprehensiveStrategyGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyComprehensiveStrategyGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyComprehensiveStrategyGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByProductGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 总花费详情

		SummaryTotal *BusinessSummaryTotal `json:"SummaryTotal,omitempty" name:"SummaryTotal"`
		// 各产品花费分布

		SummaryOverview []*BusinessSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByProductGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByProductGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchDiscountRequest struct {
	*tchttp.BaseRequest

	// 折扣类型  0用户折扣 1官网折扣

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 折扣ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 开始生效时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 失效时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 折扣创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 折扣状态 0:未生效 1:生效；2:已失效

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 用户UIN

	TargetUserId *uint64 `json:"TargetUserId,omitempty" name:"TargetUserId"`
	// 页码 从1开始

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 分页大小

	Rows *uint64 `json:"Rows,omitempty" name:"Rows"`
	// 开始时间大于该值

	PreFBeginTime *string `json:"PreFBeginTime,omitempty" name:"PreFBeginTime"`
	// 开始时间小于该值

	TailFBeginTime *string `json:"TailFBeginTime,omitempty" name:"TailFBeginTime"`
	// 结束时间大于该值

	PreFEndTime *string `json:"PreFEndTime,omitempty" name:"PreFEndTime"`
	// 结束时间小于该值

	TailFEndTime *string `json:"TailFEndTime,omitempty" name:"TailFEndTime"`
	// 创建时间大于该值

	PreFCreateTime *string `json:"PreFCreateTime,omitempty" name:"PreFCreateTime"`
	// 创建时间小于该值

	TailFCreateTime *string `json:"TailFCreateTime,omitempty" name:"TailFCreateTime"`
	// 更新时间大于该值（用户折扣preFUpdateTime表示优先级）

	PreFUpdateTime *string `json:"PreFUpdateTime,omitempty" name:"PreFUpdateTime"`
	// 更新时间小于该值（用户折扣preFUpdateTime表示优先级）

	TailFUpdateTime *string `json:"TailFUpdateTime,omitempty" name:"TailFUpdateTime"`
	// 用户更新时间大于该值（只对用户折扣有效）

	PreFAutoUpdateTime *string `json:"PreFAutoUpdateTime,omitempty" name:"PreFAutoUpdateTime"`
	// 用户更新时间小于该值（只对用户折扣有效）

	TailFAutoUpdateTime *string `json:"TailFAutoUpdateTime,omitempty" name:"TailFAutoUpdateTime"`
	// 付费模式

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
}

func (r *SearchDiscountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchDiscountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 地域花费详情

		SummaryDetail []*RegionSummaryDetailItem `json:"SummaryDetail,omitempty" name:"SummaryDetail"`
		// 记录数，不传NeedRecordNum或者传的值为0，该字段值为null

		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillFeeByRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeByRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAllProductTreeGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品四层定义列表

		List []*ProductRelation `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAllProductTreeGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAllProductTreeGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductStrategy struct {
	// ProductCode

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// SubProductCode

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// PayMode

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// TimeUnit

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// CalcUnit

	CalcUnit *string `json:"CalcUnit,omitempty" name:"CalcUnit"`
	// ActionType

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
	// BillingType

	BillingType *string `json:"BillingType,omitempty" name:"BillingType"`
	// PurchaseSyncDelivery

	PurchaseSyncDelivery *string `json:"PurchaseSyncDelivery,omitempty" name:"PurchaseSyncDelivery"`
	// ModifySyncDelivery

	ModifySyncDelivery *string `json:"ModifySyncDelivery,omitempty" name:"ModifySyncDelivery"`
	// ResourceRecycleType

	ResourceRecycleType *string `json:"ResourceRecycleType,omitempty" name:"ResourceRecycleType"`
	// InterfaceNames

	InterfaceNames *InterfaceNames `json:"InterfaceNames,omitempty" name:"InterfaceNames"`
	// Region2interface

	Region2interface *Region2interface `json:"Region2interface,omitempty" name:"Region2interface"`
}

type AddUserStrategyGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUserStrategyGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserStrategyGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInvoiceDetailGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户开票流水号

		OpenId *int64 `json:"OpenId,omitempty" name:"OpenId"`
		// 发票金额

		Amount *int64 `json:"Amount,omitempty" name:"Amount"`
		// 申请开票时间

		ApplicationTime *string `json:"ApplicationTime,omitempty" name:"ApplicationTime"`
		// 发票类型,取值为(  2:公司普通发票  3:公司增值税发票 )

		UserType *int64 `json:"UserType,omitempty" name:"UserType"`
		// 发票抬头

		InvoiceHead *string `json:"InvoiceHead,omitempty" name:"InvoiceHead"`
		// 税务登记号

		TaxNumber *string `json:"TaxNumber,omitempty" name:"TaxNumber"`
		// 开户行

		BankUserName *string `json:"BankUserName,omitempty" name:"BankUserName"`
		// 银行账号

		AccountNumber *string `json:"AccountNumber,omitempty" name:"AccountNumber"`
		// 注册全地

		RegisterFullAddr *string `json:"RegisterFullAddr,omitempty" name:"RegisterFullAddr"`
		// 注册电话

		RegisterPhone *string `json:"RegisterPhone,omitempty" name:"RegisterPhone"`
		// 地址

		Addr *string `json:"Addr,omitempty" name:"Addr"`
		// 联系人

		Contact *string `json:"Contact,omitempty" name:"Contact"`
		// 手机号码

		Cellphone *string `json:"Cellphone,omitempty" name:"Cellphone"`
		// 发票状态,1:处理中,6:已取消 30:审核通过  31:审核不通过

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 发票备注

		InvoiceRemark *string `json:"InvoiceRemark,omitempty" name:"InvoiceRemark"`
		// 城市

		City *string `json:"City,omitempty" name:"City"`
		// 省份

		Province *string `json:"Province,omitempty" name:"Province"`
		// 账期

		AccountPeriods *string `json:"AccountPeriods,omitempty" name:"AccountPeriods"`
		// 审核操作人

		Auditor *string `json:"Auditor,omitempty" name:"Auditor"`
		// 审核人拒绝原因

		Reason *string `json:"Reason,omitempty" name:"Reason"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInvoiceDetailGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInvoiceDetailGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSignatureGatewayRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSignatureGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSignatureGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResourceBillGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 计量推量明细数组

		ResourceBill []*BspResourceBill `json:"ResourceBill,omitempty" name:"ResourceBill"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetResourceBillGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResourceBillGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadWithOperatorRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeDownloadWithOperatorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadWithOperatorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountListGatewayRequest struct {
	*tchttp.BaseRequest

	// 管理端查询用户UIN，默认缺省查全部

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAccountListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDosagesGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 如果没有用量，值为0

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 如果没有用量，则空数组

		Dosages *DosageList `json:"Dosages,omitempty" name:"Dosages"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDosagesGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDosagesGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInvoiceDetailGatewayRequest struct {
	*tchttp.BaseRequest

	// 用户开票流水号

	OpenId *int64 `json:"OpenId,omitempty" name:"OpenId"`
}

func (r *DescribeInvoiceDetailGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInvoiceDetailGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceBillDownloadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载链接

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceBillDownloadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceBillDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 资源状态信息结果集

		List []*ResourceInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListResourceGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Region2interface struct {
	// url

	Url []*string `json:"Url,omitempty" name:"Url"`
	// host

	Host *string `json:"Host,omitempty" name:"Host"`
}

type DescribeBillDownloadByResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务是否已完成（0未完成/1已完成）

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 文件名称，当任务已完成时返回该字段

		Name *string `json:"Name,omitempty" name:"Name"`
		// 下载链接，当任务已完成时返回该字段

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDownloadByResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadByResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByPayModeGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillSummaryByPayModeGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByPayModeGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BspSubProductList struct {
	// 子产品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品名称

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 产品标识

	DisplayCode *string `json:"DisplayCode,omitempty" name:"DisplayCode"`
	// 产品标识

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 产品细项数量

	SubBillingItemCount *string `json:"SubBillingItemCount,omitempty" name:"SubBillingItemCount"`
}

type DescribeQuotasGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配额数据列表

		List []*QuotaList `json:"List,omitempty" name:"List"`
		// 数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQuotasGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotasGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveStrategyWebConfigGatewayRequest struct {
	*tchttp.BaseRequest
}

func (r *GetComprehensiveStrategyWebConfigGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveStrategyWebConfigGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiscountRequest struct {
	*tchttp.BaseRequest

	// 折扣类型 0用户折扣 1官网折扣

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 折扣Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 开始生效时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 失效时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 折扣值

	Discount *string `json:"Discount,omitempty" name:"Discount"`
	// 折扣条件 json字符串

	ConditionList *string `json:"ConditionList,omitempty" name:"ConditionList"`
}

func (r *ModifyDiscountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiscountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReversalResourceGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReversalResourceGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReversalResourceGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConditionProject struct{}

type DescribeQuotaAssignLogsListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 账户配额转移流水列表

		QuotaAssignList []*QuotaAssignList `json:"QuotaAssignList,omitempty" name:"QuotaAssignList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQuotaAssignLogsListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotaAssignLogsListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserLimit struct {
	// 天限

	DayLimit *int64 `json:"DayLimit,omitempty" name:"DayLimit"`
	// 周限

	WeekLimit *int64 `json:"WeekLimit,omitempty" name:"WeekLimit"`
	// 月限

	MonthLimit *int64 `json:"MonthLimit,omitempty" name:"MonthLimit"`
	// 总限

	AllLimit *int64 `json:"AllLimit,omitempty" name:"AllLimit"`
}

type GetComprehensiveStrategyWebConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计费模式

		PayMode []*string `json:"PayMode,omitempty" name:"PayMode"`
		// timeUnit

		TimeUnit []*string `json:"TimeUnit,omitempty" name:"TimeUnit"`
		// calcUnit

		CalcUnit []*string `json:"CalcUnit,omitempty" name:"CalcUnit"`
		// actionType

		ActionType []*string `json:"ActionType,omitempty" name:"ActionType"`
		// billingType

		BillingType []*string `json:"BillingType,omitempty" name:"BillingType"`
		// priceType

		PriceType []*string `json:"PriceType,omitempty" name:"PriceType"`
		// purchaseSyncDelivery

		PurchaseSyncDelivery []*string `json:"PurchaseSyncDelivery,omitempty" name:"PurchaseSyncDelivery"`
		// modifySyncDelivery

		ModifySyncDelivery []*string `json:"ModifySyncDelivery,omitempty" name:"ModifySyncDelivery"`
		// status

		Status []*string `json:"Status,omitempty" name:"Status"`
		// ResourceRecycleType

		ResourceRecycleType []*string `json:"ResourceRecycleType,omitempty" name:"ResourceRecycleType"`
		// region

		Region []*string `json:"Region,omitempty" name:"Region"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComprehensiveStrategyWebConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveStrategyWebConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductPropertyValues struct {
	// kv数组

	DataArray *PropertyFieldsValuesDataArray `json:"DataArray,omitempty" name:"DataArray"`
	// 该value状态： 1-正常 2-失效

	Status *string `json:"Status,omitempty" name:"Status"`
}

type SetHiddenProductRequest struct {
	*tchttp.BaseRequest

	// 隐藏商品项信息

	ProductInfo []*HiddenProduct `json:"ProductInfo,omitempty" name:"ProductInfo"`
}

func (r *SetHiddenProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetHiddenProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryRecoverUinGatewayRequest struct {
	*tchttp.BaseRequest

	// 跳过条数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 本页返回条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 本次需要查询的uin，支持模糊查询

	SearchUin *string `json:"SearchUin,omitempty" name:"SearchUin"`
}

func (r *QueryRecoverUinGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryRecoverUinGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLackingQuotaRecordGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据总条数

		Total *string `json:"Total,omitempty" name:"Total"`
		// 分页列表

		LackingQuotaRecordList []*LackingQuotaRecordList `json:"LackingQuotaRecordList,omitempty" name:"LackingQuotaRecordList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListLackingQuotaRecordGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLackingQuotaRecordGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 各地域花费分布详情

		SummaryOverview []*RegionSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductPropertyList struct {
	// key中文名

	KeyNameCn *string `json:"KeyNameCn,omitempty" name:"KeyNameCn"`
	// key英文名

	KeyNameEn *string `json:"KeyNameEn,omitempty" name:"KeyNameEn"`
	// key码，后端定义标识规则

	KeyCode *string `json:"KeyCode,omitempty" name:"KeyCode"`
	// key的字段列表

	Fields []*ProductPropertyFields `json:"Fields,omitempty" name:"Fields"`
	// JSON String 属性值列表  jsond ecode后是数组，数组元素的key不定  例如：[{

	Values *string `json:"Values,omitempty" name:"Values"`
	// 是否可被编辑。0是不可以；1是可以

	Modifiable *int64 `json:"Modifiable,omitempty" name:"Modifiable"`
}

type DescribeBillDownloadCommonSummaryGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务是否已完成（0未完成/1已完成）

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 文件名称，当任务已完成时返回该字段

		Name *string `json:"Name,omitempty" name:"Name"`
		// 下载链接，当任务已完成时返回该字段

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDownloadCommonSummaryGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadCommonSummaryGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceDetailListRequest struct {
	*tchttp.BaseRequest

	// 资源信息

	ResourcePara *string `json:"ResourcePara,omitempty" name:"ResourcePara"`
	// 用户Appid

	UserAppid *int64 `json:"UserAppid,omitempty" name:"UserAppid"`
	// 用户Uin

	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`
}

func (r *DescribeResourceDetailListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceDetailListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetQuotaLeftGatewayRequest struct {
	*tchttp.BaseRequest

	// 一层商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 账户uin

	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`
	// 商品配额key，由四层定义某几层以及"#"号拼接而成，如p_cvm#sp_cvm_vself2#v_cvm_cpu#sv_cvm_cpu_vself2

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
}

func (r *GetQuotaLeftGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetQuotaLeftGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillTrendByMonthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 每月花费详情

		Detail []*TrendByMonthDetail `json:"Detail,omitempty" name:"Detail"`
		// 平均花费详情

		Stat []*TrendByMonthStat `json:"Stat,omitempty" name:"Stat"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillTrendByMonthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillTrendByMonthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConditionProduct struct {
	// 产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 产品名称

	SubProductCodeName *string `json:"SubProductCodeName,omitempty" name:"SubProductCodeName"`
}

type CancleContractTypeGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancleContractTypeGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancleContractTypeGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BasicAccount struct {
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 用户appId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 账户是否合法，1:合法，0：非法

	IsExist *int64 `json:"IsExist,omitempty" name:"IsExist"`
}

type SaveUserDiscountGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ID

		// 条数

		// 是否成功

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveUserDiscountGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveUserDiscountGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductDefineList struct {
	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品中文名称

	ProductChnName *string `json:"ProductChnName,omitempty" name:"ProductChnName"`
	// 产品英文名称

	ProductEngName *string `json:"ProductEngName,omitempty" name:"ProductEngName"`
	// 子产品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品中文名称

	SubProductChnName *string `json:"SubProductChnName,omitempty" name:"SubProductChnName"`
	// 子产品英文名称

	SubProductEngName *string `json:"SubProductEngName,omitempty" name:"SubProductEngName"`
	// 计费项代码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费项中文名称

	BillingItemChnName *string `json:"BillingItemChnName,omitempty" name:"BillingItemChnName"`
	// 计费项英文名称

	BillingItemEngName *string `json:"BillingItemEngName,omitempty" name:"BillingItemEngName"`
	// 计费细项代码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 计费细项中文名称

	SubBillingItemChnName *string `json:"SubBillingItemChnName,omitempty" name:"SubBillingItemChnName"`
	// 计费细项英文名称

	SubBillingItemEngName *string `json:"SubBillingItemEngName,omitempty" name:"SubBillingItemEngName"`
	// 计费细项单位

	SubBillingItemsGoodNumUnit *string `json:"SubBillingItemsGoodNumUnit,omitempty" name:"SubBillingItemsGoodNumUnit"`
	// 计费细项单位英文

	SubBillingItemsGoodNumUnitEn *string `json:"SubBillingItemsGoodNumUnitEn,omitempty" name:"SubBillingItemsGoodNumUnitEn"`
	// 四层定义状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 最近操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type SearchDiscountList struct {
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 失效时间

	FEndTime *string `json:"FEndTime,omitempty" name:"FEndTime"`
	// 计费细项

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 计费项

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 子商品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 用户UIN（当Type为user时有效
	// ）

	FUserId *string `json:"FUserId,omitempty" name:"FUserId"`
	// 是否和官网折扣冲突（当Type为user时有效
	// ）

	FConflict *uint64 `json:"FConflict,omitempty" name:"FConflict"`
	// 是否内部用户折扣，0否，1是

	FType *uint64 `json:"FType,omitempty" name:"FType"`
	// 折扣创建者

	FCreator *string `json:"FCreator,omitempty" name:"FCreator"`
	// 优惠状态，0初始化，1生效，2废弃，3未开始，4已结束

	FStatus *uint64 `json:"FStatus,omitempty" name:"FStatus"`
	// 创建时间

	FCreateTime *string `json:"FCreateTime,omitempty" name:"FCreateTime"`
	// 折扣Id（不传折扣Id为新增，传Id为修改
	// ）

	FId *uint64 `json:"FId,omitempty" name:"FId"`
	// 开始生效时间

	FBeginTime *string `json:"FBeginTime,omitempty" name:"FBeginTime"`
	// 折扣备注

	FRemark *string `json:"FRemark,omitempty" name:"FRemark"`
	// 修饰符

	FModifier *string `json:"FModifier,omitempty" name:"FModifier"`
	// 折扣更新时间

	FUpdateTime *string `json:"FUpdateTime,omitempty" name:"FUpdateTime"`
	// 折扣百分值

	FDiscount *string `json:"FDiscount,omitempty" name:"FDiscount"`
	// 折扣条件（具体参数条件对象请求参数
	// ）

	FConditions *string `json:"FConditions,omitempty" name:"FConditions"`
}

type DescribeRefundRequest struct {
	*tchttp.BaseRequest

	// 米大师appid，由计平分配

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 查询的起始时间，unix时间戳（格林威治时间）,精确到秒

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询的结束时间，unix时间戳（格林威治时间）,精确到秒

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 每页返回的记录数。根据用户号码查询订单列表时需要传。用于分页展示

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 记录数偏移量，默认从0开始。根据用户号码查询订单列表时需要传。用于分页展示

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 用户Id

	UserNum *string `json:"UserNum,omitempty" name:"UserNum"`
	// 指定退款id查询。

	RefundId *string `json:"RefundId,omitempty" name:"RefundId"`
	// 本场景固定传: unionpay

	Channel *string `json:"Channel,omitempty" name:"Channel"`
	// type=by_refund_id，根据订单号查订单. type=by_user，根据用户Id查询订单.

	Type *string `json:"Type,omitempty" name:"Type"`
	// 按状态过滤.提现状态.1-提现中;2-提现成功;3-提现失败.
	// 不传默认查询全部

	State []*string `json:"State,omitempty" name:"State"`
}

func (r *DescribeRefundRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRefundRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadCommonSummaryGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillDownloadCommonSummaryGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadCommonSummaryGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLifecycleFlowListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 流水列表

		List []*DescribeLifecycleFlowList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLifecycleFlowListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLifecycleFlowListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceList struct {
	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 地域id

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 资源id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type DescribeContractTypeListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 合同类型列表

		List []*ContractTypeListData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContractTypeListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContractTypeListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAutoFlagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetAutoFlagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAutoFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaTypeListGatewayRequest struct {
	*tchttp.BaseRequest

	// midas appid

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
}

func (r *DescribeQuotaTypeListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotaTypeListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveSubProdctStrategyGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 跳过条数

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 本次返回条数

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据

		Info []*ProductStrategy `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComprehensiveSubProdctStrategyGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveSubProdctStrategyGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditRemitRecordGatewayRequest struct {
	*tchttp.BaseRequest

	// 目标uin

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 0:处理中、1:充值成功、2:充值失败、3:拒绝审核

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 审核数据id

	TargetId *uint64 `json:"TargetId,omitempty" name:"TargetId"`
	// 拒绝理由

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

func (r *AuditRemitRecordGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuditRemitRecordGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByPayModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 计费模式花费详情

		SummaryDetail []*PayModeSummaryDetailItem `json:"SummaryDetail,omitempty" name:"SummaryDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillFeeByPayModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeByPayModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportChangeQuotaGatewayRequest struct {
	*tchttp.BaseRequest

	// 变更时间-开始

	CreateStartTime *string `json:"CreateStartTime,omitempty" name:"CreateStartTime"`
	// 变更时间-结束

	CreateEndTime *string `json:"CreateEndTime,omitempty" name:"CreateEndTime"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 产品细项码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 子产品细项码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 变更者UIN

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 租户UIN

	SearchUin *string `json:"SearchUin,omitempty" name:"SearchUin"`
	// 配额类型

	QuotaType *string `json:"QuotaType,omitempty" name:"QuotaType"`
}

func (r *ExportChangeQuotaGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportChangeQuotaGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateResourceRequest struct {
	*tchttp.BaseRequest

	// 目标用户

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

func (r *IsolateResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolateResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAcctResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAcctResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAcctResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefundDetail struct {
	// 提现id

	RefundId *string `json:"RefundId,omitempty" name:"RefundId"`
	// 外部订单号

	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`
	// 用户id

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 提现金额,单位：分

	RefundAmt *string `json:"RefundAmt,omitempty" name:"RefundAmt"`
	// 原订单支付金额,单位分

	OrderPayAmt *string `json:"OrderPayAmt,omitempty" name:"OrderPayAmt"`
	// 支付渠道

	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`
	// 提现时间

	RefundTime *string `json:"RefundTime,omitempty" name:"RefundTime"`
	// 提现状态.1-提现中;2-提现成功;3-提现失败

	State *string `json:"State,omitempty" name:"State"`
}

type GetAllProductRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAllProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAllProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditInvoiceInfoRequest struct {
	*tchttp.BaseRequest

	// 发票流水id

	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`
	// 审核状态 取值为(2:审核通过 3:审核驳回)

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 审核驳回原因,如果审核通过可以为空

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

func (r *AuditInvoiceInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuditInvoiceInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContractTypeListGatewayRequest struct {
	*tchttp.BaseRequest

	// 类型状态 1生效中 0作废

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeContractTypeListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContractTypeListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotasGatewayRequest struct {
	*tchttp.BaseRequest

	// 传空字符串返回system默认配额，
	// 不传就返回全量用户配额

	QuotaUin *string `json:"QuotaUin,omitempty" name:"QuotaUin"`
	// 产品code。不传或者空就是查询全部

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
}

func (r *DescribeQuotasGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotasGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByRegionGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 每次获取数据量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 是否返回最大页码，建议首次或变更PageSize后才需要，用于提升接口性

	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
}

func (r *DescribeBillFeeByRegionGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeByRegionGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConditionBillingItem struct {
	// 组件码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 组件名称

	BillingItemCodeName *string `json:"BillingItemCodeName,omitempty" name:"BillingItemCodeName"`
}

type SetHiddenProductGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetHiddenProductGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetHiddenProductGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByPayModeRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillSummaryByPayModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByPayModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportChangeQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportChangeQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportChangeQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBreakdownModeRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryBreakdownModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBreakdownModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Conditions struct {
	// 只支持6和12两个值

	TimeRange *uint64 `json:"TimeRange,omitempty" name:"TimeRange"`
	// 地域ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 付费模式，可选prePay和postPay

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 资源关键字

	ResourceKeyword *string `json:"ResourceKeyword,omitempty" name:"ResourceKeyword"`
	// 资源所属项目id

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目目录id列表

	OrgIds []*string `json:"OrgIds,omitempty" name:"OrgIds"`
	// 项目及目录列表

	ProjectSet []*ProjectSet `json:"ProjectSet,omitempty" name:"ProjectSet"`
	// 产品编码

	ProductCodes []*string `json:"ProductCodes,omitempty" name:"ProductCodes"`
	// 子产品编码

	SubProductCodes []*string `json:"SubProductCodes,omitempty" name:"SubProductCodes"`
	// 地域ID

	RegionIds []*int64 `json:"RegionIds,omitempty" name:"RegionIds"`
	// 付费模式，可选prePay和postPay

	PayModes []*string `json:"PayModes,omitempty" name:"PayModes"`
	// 交易类型

	ActionTypes []*string `json:"ActionTypes,omitempty" name:"ActionTypes"`
	// 是否隐藏0元流水

	HideFreeCost *int64 `json:"HideFreeCost,omitempty" name:"HideFreeCost"`
	// 排序规则，可选desc和asc

	OrderByCost *string `json:"OrderByCost,omitempty" name:"OrderByCost"`
	// 交易ID

	BillIds []*string `json:"BillIds,omitempty" name:"BillIds"`
	// 组件编码

	BillingItemCodes []*string `json:"BillingItemCodes,omitempty" name:"BillingItemCodes"`
	// 文件ID

	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`
	// 文件类型

	FileTypes []*string `json:"FileTypes,omitempty" name:"FileTypes"`
	// 状态

	Status []*uint64 `json:"Status,omitempty" name:"Status"`
	// 导出字段

	ExportFields []*string `json:"ExportFields,omitempty" name:"ExportFields"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 标签数组

	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`
	// 导出标签

	ExportTags []*string `json:"ExportTags,omitempty" name:"ExportTags"`
	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type VoucherConditions struct {
	// 导出字段

	ExportFields []*string `json:"ExportFields,omitempty" name:"ExportFields"`
	// 用户ID

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 指定查询状态

	SpStatus *string `json:"SpStatus,omitempty" name:"SpStatus"`
	// 指定代金券名称

	SpName *string `json:"SpName,omitempty" name:"SpName"`
	// 指定代金券编号

	SpId *string `json:"SpId,omitempty" name:"SpId"`
	// 指定代金券批次

	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
	// 有效开始时间

	ValidBeginTime *string `json:"ValidBeginTime,omitempty" name:"ValidBeginTime"`
	// 有效结束时间

	ValidEndTime *string `json:"ValidEndTime,omitempty" name:"ValidEndTime"`
	// 发放开始时间

	PresentBeginTime *string `json:"PresentBeginTime,omitempty" name:"PresentBeginTime"`
	// 发放结束时间

	PresentEndTime *string `json:"PresentEndTime,omitempty" name:"PresentEndTime"`
	// 指定排序字段：begin_time、end_time。

	SortField *string `json:"SortField,omitempty" name:"SortField"`
	// 指定升序降序：desc、asc。

	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

type ExportDealListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportDealListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDealListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReconciliationConditions struct {
	// 支付者uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 账单月份

	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`
	// 调账类型 add-补偿/minus-扣费

	AdjustType *string `json:"AdjustType,omitempty" name:"AdjustType"`
	// 操作起始时间，当月第一秒，格式为

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 操作结束时间，当月最后一秒，格式为

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 状态，1为调账成功，0为调账失败的

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 调账申请人

	ApplyUin *string `json:"ApplyUin,omitempty" name:"ApplyUin"`
}

type DescribeBspProductRelationsGatewayRequest struct {
	*tchttp.BaseRequest

	// 产品code或名称

	SearchProduct *string `json:"SearchProduct,omitempty" name:"SearchProduct"`
	// 子产品code或名称

	SearchSubProduct *string `json:"SearchSubProduct,omitempty" name:"SearchSubProduct"`
	// 计费项code或名称

	SearchBillingItem *string `json:"SearchBillingItem,omitempty" name:"SearchBillingItem"`
	// 计费细项code或名称

	SearchSubBillingItem *string `json:"SearchSubBillingItem,omitempty" name:"SearchSubBillingItem"`
	// 四层定义状态，1-生效 0-失效

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 页码

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 产品code数组

	SearchProductCodes []*string `json:"SearchProductCodes,omitempty" name:"SearchProductCodes"`
	// 子商品代码数组

	SearchSubProductCodes []*string `json:"SearchSubProductCodes,omitempty" name:"SearchSubProductCodes"`
	// 计费项代码数组

	SearchBillingItemCodes []*string `json:"SearchBillingItemCodes,omitempty" name:"SearchBillingItemCodes"`
	// 计费细项代码数组

	SearchSubBillingItemCodes []*string `json:"SearchSubBillingItemCodes,omitempty" name:"SearchSubBillingItemCodes"`
	// 数据来源类型。local（本系统。默认），third（第三方接入）。不传默认是local

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 添加的是计量产品还是计费产品，取值范围计量（bsp）计费（billing），不填写则默认bsp

	Source *string `json:"Source,omitempty" name:"Source"`
}

func (r *DescribeBspProductRelationsGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBspProductRelationsGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePayOrderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单数量

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 订单详情

		OrderList []*OrderInfo `json:"OrderList,omitempty" name:"OrderList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePayOrderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePayOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubBillingItemsGatewayRequest struct {
	*tchttp.BaseRequest

	// midas appid

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 计费细项code列表

	SubBillingItemCodes []*string `json:"SubBillingItemCodes,omitempty" name:"SubBillingItemCodes"`
}

func (r *DescribeSubBillingItemsGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubBillingItemsGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductTreeGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品四层定义列表

		List []*ProductRelation `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductTreeGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductTreeGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReversalResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReversalResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReversalResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuotaLeft struct {
	// 配额key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 剩余配额

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type QueryBSPBillingItemGatewayRequest struct {
	*tchttp.BaseRequest

	// 商品代码或名称

	SearchProduct *string `json:"SearchProduct,omitempty" name:"SearchProduct"`
	// 子商品代码或名称

	SearchSubProduct *string `json:"SearchSubProduct,omitempty" name:"SearchSubProduct"`
	// 计费项代码或名称

	SearchBillingItem *string `json:"SearchBillingItem,omitempty" name:"SearchBillingItem"`
	// 计费细项代码或名称

	SearchSubBillingItem *string `json:"SearchSubBillingItem,omitempty" name:"SearchSubBillingItem"`
	// 四层定义状态，1-生效 0-失效 （不传默认生效

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 商品代码

	SearchProductCode *string `json:"SearchProductCode,omitempty" name:"SearchProductCode"`
	// 子商品代码

	SearchSubProductCode *string `json:"SearchSubProductCode,omitempty" name:"SearchSubProductCode"`
	// 计费项代码

	SearchBillingItemCode *string `json:"SearchBillingItemCode,omitempty" name:"SearchBillingItemCode"`
	// 计费细项代码

	SearchSubBillingItemCode *string `json:"SearchSubBillingItemCode,omitempty" name:"SearchSubBillingItemCode"`
	// 运营模式：bsp-计量、billing 计费（不传默认计费

	Source *string `json:"Source,omitempty" name:"Source"`
}

func (r *QueryBSPBillingItemGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBSPBillingItemGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDiscountGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1

		// 1

		// 1

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDiscountGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDiscountGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteQuotaGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteQuotaGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteQuotaGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductStrategyGatewayRequest struct {
	*tchttp.BaseRequest

	// 产品编号

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品子编号

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 地域信息编号

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 偏移量 默认值为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 单次查询限制条数 默认值为10

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询商品状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *GetProductStrategyGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductStrategyGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductPropertyListFields struct {
	// 字段名

	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`
	// 字段码

	FieldCode *string `json:"FieldCode,omitempty" name:"FieldCode"`
	// 字段标志： 0-普通 1-主标记

	FieldFlag *string `json:"FieldFlag,omitempty" name:"FieldFlag"`
}

type DescribeGoodsPriceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGoodsPriceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGoodsPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveStrategyWebConfigGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计费模式

		PayMode []*string `json:"PayMode,omitempty" name:"PayMode"`
		// timeUnit

		TimeUnit []*string `json:"TimeUnit,omitempty" name:"TimeUnit"`
		// calcUnit

		CalcUnit []*string `json:"CalcUnit,omitempty" name:"CalcUnit"`
		// actionType

		ActionType []*string `json:"ActionType,omitempty" name:"ActionType"`
		// billingType

		BillingType []*string `json:"BillingType,omitempty" name:"BillingType"`
		// priceType

		PriceType []*string `json:"PriceType,omitempty" name:"PriceType"`
		// purchaseSyncDelivery

		PurchaseSyncDelivery []*string `json:"PurchaseSyncDelivery,omitempty" name:"PurchaseSyncDelivery"`
		// modifySyncDelivery

		ModifySyncDelivery []*string `json:"ModifySyncDelivery,omitempty" name:"ModifySyncDelivery"`
		// status

		Status []*string `json:"Status,omitempty" name:"Status"`
		// ResourceRecycleType

		ResourceRecycleType []*string `json:"ResourceRecycleType,omitempty" name:"ResourceRecycleType"`
		// region

		Region []*string `json:"Region,omitempty" name:"Region"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComprehensiveStrategyWebConfigGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveStrategyWebConfigGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductStrategyRequest struct {
	*tchttp.BaseRequest

	// 产品编号

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品子编号

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 地域信息编号

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 偏移量 默认值为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 单次查询限制条数 默认值为10

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询商品状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *GetProductStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadRecordListRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 每次获取数据量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 是否需要记录条数，0需要，1不需要，默认不需要

	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
	// 排序规则

	Sort *GetDownloadListSort `json:"Sort,omitempty" name:"Sort"`
	// 只支持FIleIds，FileTypes，Status三种过滤条件

	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
	// 是否需要条件取值（0不需要/1需要，默认为0）

	NeedConditionValue *string `json:"NeedConditionValue,omitempty" name:"NeedConditionValue"`
}

func (r *DescribeBillDownloadRecordListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadRecordListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTagsByMonthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// tagkey的数组，如果没有记录则为空数组

		TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTagsByMonthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTagsByMonthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveCommonDiscountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 折扣id

		// 条数

		// 是否成功

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveCommonDiscountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveCommonDiscountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditInvoiceInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AuditInvoiceInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuditInvoiceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAllProductTreeRequest struct {
	*tchttp.BaseRequest

	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 页索引

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 是否生效.0未生效 1生效

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *GetAllProductTreeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAllProductTreeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyContractTypeGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyContractTypeGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyContractTypeGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 记录数，不传NeedRecordNum或者传的值为0，该字段值为null

		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`
		// 产品花费详情

		SummaryDetail []*BusinessSummaryDetailItem `json:"SummaryDetail,omitempty" name:"SummaryDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillFeeByProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeByProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceDownloadGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载链接

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceDownloadGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceDownloadGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePropertiesRequest struct {
	*tchttp.BaseRequest

	// key码，空时返回所有

	KeyCode *string `json:"KeyCode,omitempty" name:"KeyCode"`
	// 页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 类型： 0（或者不传或者传空）是普通属性（默认）；1是自定义属性

	PropertyType *int64 `json:"PropertyType,omitempty" name:"PropertyType"`
}

func (r *DescribePropertiesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePropertiesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTagsByMonthGatewayRequest struct {
	*tchttp.BaseRequest

	// uin

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// appId

	PayerAppId *string `json:"PayerAppId,omitempty" name:"PayerAppId"`
	// 月份

	Month *string `json:"Month,omitempty" name:"Month"`
	// 查询全部，传1

	State *uint64 `json:"State,omitempty" name:"State"`
}

func (r *GetTagsByMonthGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTagsByMonthGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillingUinQuotaGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集合

		List *UinQuota `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillingUinQuotaGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillingUinQuotaGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBspUinQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集合

		List *UinQuota `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBspUinQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBspUinQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByProductGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 每次获取数据量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 是否返回最大页码，建议首次或变更PageSize后才需要，用于提升接口性

	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
}

func (r *DescribeBillFeeByProductGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeByProductGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuotaAssignList struct {
	// 计费项变更前数据

	BillingItemAfterData *string `json:"BillingItemAfterData,omitempty" name:"BillingItemAfterData"`
	// 计费项变更后数据

	BillingItemBeforeData *string `json:"BillingItemBeforeData,omitempty" name:"BillingItemBeforeData"`
	// 计费项code

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费项中文名

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 变更内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 变更时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 操作人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 产品变更前数据

	ProductAfterData *string `json:"ProductAfterData,omitempty" name:"ProductAfterData"`
	// 产品变更后数据

	ProductBeforeData *string `json:"ProductBeforeData,omitempty" name:"ProductBeforeData"`
	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品中文名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 配额类型

	QuotaType *string `json:"QuotaType,omitempty" name:"QuotaType"`
	// 计费子项变更前数据

	SubBillingItemAfterData *string `json:"SubBillingItemAfterData,omitempty" name:"SubBillingItemAfterData"`
	// 计费子项变更后数据

	SubBillingItemBeforeData *string `json:"SubBillingItemBeforeData,omitempty" name:"SubBillingItemBeforeData"`
	// 计费子项code

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 计费子项中文名

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
	// 子产品变更前数据

	SubProductAfterData *string `json:"SubProductAfterData,omitempty" name:"SubProductAfterData"`
	// 子产品变更后数据

	SubProductBeforeData *string `json:"SubProductBeforeData,omitempty" name:"SubProductBeforeData"`
	// 子产品code

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品中文名

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 被修改配额的用户Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 产品英文名

	ProductEngName *string `json:"ProductEngName,omitempty" name:"ProductEngName"`
	// 子产品英文名

	SubProductEngName *string `json:"SubProductEngName,omitempty" name:"SubProductEngName"`
	// 产品项英文名

	BillingItemEngName *string `json:"BillingItemEngName,omitempty" name:"BillingItemEngName"`
	// 产品细项英文名

	SubBillingItemEngName *string `json:"SubBillingItemEngName,omitempty" name:"SubBillingItemEngName"`
	// 英文内容

	EngContent *string `json:"EngContent,omitempty" name:"EngContent"`
}

type DescribeResourceDownloadRequest struct {
	*tchttp.BaseRequest

	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 资源id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 时间范围-开始时间

	StartAddTime *string `json:"StartAddTime,omitempty" name:"StartAddTime"`
	// 时间范围-结束时间

	EndAddTime *string `json:"EndAddTime,omitempty" name:"EndAddTime"`
}

func (r *DescribeResourceDownloadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceDownloadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveSubProductStrategyRequest struct {
	*tchttp.BaseRequest

	// 产品代码，如果该值不传则默认查询所有产品

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品代码，如果该值不传则默认查询该产品下所有子产品

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 是否只查询第三方自定义产品策略,如果该值不传或者为false则查询全部产品策略

	IfCustom *string `json:"IfCustom,omitempty" name:"IfCustom"`
	// 状态，如果该值不传则默认查询所有状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 本次查询条数，如果该值不传则默认100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 跳过条数如果该值不传则默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetComprehensiveSubProductStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveSubProductStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceDetailDownloadGatewayRequest struct {
	*tchttp.BaseRequest

	// 数据所属用户的UIN

	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 产品编码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

func (r *ListResourceDetailDownloadGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceDetailDownloadGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGetDownloadUrlWithAuthGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGetDownloadUrlWithAuthGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGetDownloadUrlWithAuthGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPkgDownloadListItem struct {
	// 账单月份

	MonthName *string `json:"MonthName,omitempty" name:"MonthName"`
	// 当月开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 当月结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 状态(0未出账/1已出账但未生成账单/2已出账且账单生成中/4已出账且已生成账单)

	Status *string `json:"Status,omitempty" name:"Status"`
	// 应付金额

	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`
	// 代金券支付总价

	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
	// 下载链接，Status等于4该值才有效

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
	// 当月周期类型

	PeriodType *string `json:"PeriodType,omitempty" name:"PeriodType"`
	// 出账日期

	BillDate *string `json:"BillDate,omitempty" name:"BillDate"`
}

type InvoiceAmount struct {
	// 可开票金额,单位分

	AvailableOpenAmount *int64 `json:"AvailableOpenAmount,omitempty" name:"AvailableOpenAmount"`
	// 消耗金额,单位分

	Consume *int64 `json:"Consume,omitempty" name:"Consume"`
	// 已开票金额,单位分

	OpenedAmount *int64 `json:"OpenedAmount,omitempty" name:"OpenedAmount"`
}

type LackingQuotaRecordList struct {
	// 主键

	Id *string `json:"Id,omitempty" name:"Id"`
	// 配额不足类型；1：项目级配额不足  2：主账号级(租户级)配额不足 3：子账号级配额不足

	LackingType *string `json:"LackingType,omitempty" name:"LackingType"`
	// 租户AppId

	MasterAccountAppid *string `json:"MasterAccountAppid,omitempty" name:"MasterAccountAppid"`
	// 子账号uin

	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 配额项名称

	QuotaName *string `json:"QuotaName,omitempty" name:"QuotaName"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 设置的配额值

	QuotaValue *string `json:"QuotaValue,omitempty" name:"QuotaValue"`
	// 已经使用的配额值

	UsedQuota *string `json:"UsedQuota,omitempty" name:"UsedQuota"`
	// 计划使用配额

	PlanQuota *string `json:"PlanQuota,omitempty" name:"PlanQuota"`
	// 配额单位；只有当前配额项为4层产品单位才有值

	QuotaUnit *string `json:"QuotaUnit,omitempty" name:"QuotaUnit"`
	// 记录时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 操作账号uin

	OperatorUin *string `json:"OperatorUin,omitempty" name:"OperatorUin"`
}

type DescribeQuotasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配额数据列表

		List []*QuotaList `json:"List,omitempty" name:"List"`
		// 数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQuotasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyContractStatusGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyContractStatusGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyContractStatusGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBspProductInfoGatewayRequest struct {
	*tchttp.BaseRequest

	// 产品定义

	Product *Product `json:"Product,omitempty" name:"Product"`
	// 子产品定义

	SubProduct *SubProduct `json:"SubProduct,omitempty" name:"SubProduct"`
	// 计费项定义

	BillingItem *BillingItem `json:"BillingItem,omitempty" name:"BillingItem"`
	// 计费细项定义

	SubBillingItem *SubBillingItem `json:"SubBillingItem,omitempty" name:"SubBillingItem"`
	// 别名

	AliasCode *string `json:"AliasCode,omitempty" name:"AliasCode"`
	// 计量周期

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 推量周期

	CalcUnit *string `json:"CalcUnit,omitempty" name:"CalcUnit"`
	// 数据来源类型。local（本系统。默认），third（第三方接入）。不传默认是local

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 添加的是计量产品还是计费产品，取值范围计量（bsp）计费（billing），不填写则默认bsp

	Source *string `json:"Source,omitempty" name:"Source"`
	// 需要修改的策略id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *SetBspProductInfoGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBspProductInfoGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDebateListGatewayRequest struct {
	*tchttp.BaseRequest

	// 管理端查询用户UIN，默认缺省查全部

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 欠费月份，0就是查欠账月份大于等于0

	UnpayDebtNum *uint64 `json:"UnpayDebtNum,omitempty" name:"UnpayDebtNum"`
	// 针对DebtNum的比较条件， E: 等于， G: 大于，GE: 大于等于

	NumSort *string `json:"NumSort,omitempty" name:"NumSort"`
}

func (r *DescribeDebateListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDebateListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceBillDownloadGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载链接

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceBillDownloadGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceBillDownloadGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ObjectArray struct {
	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 配额key

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
	// 配额值

	QuotaValue *uint64 `json:"QuotaValue,omitempty" name:"QuotaValue"`
}

type QueryRecoverUinResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本页返回条数

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 跳过条数

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 恢复uin列表

		Info []*string `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryRecoverUinResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryRecoverUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveCommonDiscountGatewayRequest struct {
	*tchttp.BaseRequest

	// 折扣参数

	DiscountPara []*DiscountPara `json:"DiscountPara,omitempty" name:"DiscountPara"`
}

func (r *SaveCommonDiscountGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveCommonDiscountGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGetDownloadUrlWithAuthGatewayRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
}

func (r *DescribeGetDownloadUrlWithAuthGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGetDownloadUrlWithAuthGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByResourceRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 每次获取数据量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 是否需要记录总数，0不需要，1需要

	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
	// 是否需要过滤条件，0需要，1不需要

	NeedConditionValue *uint64 `json:"NeedConditionValue,omitempty" name:"NeedConditionValue"`
	// 只支持BusinessCodes，ProductCodes，PayModes, ProjectIds, RegionIds, PayModes, ActionTypes, HideFreeCost, ResourceKeyword, OrderByCost

	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeBillSummaryByResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContractTypeListRequest struct {
	*tchttp.BaseRequest

	// 类型状态 1生效中 0作废

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeContractTypeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContractTypeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemitRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 列表

		List []*RemitRecordData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemitRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemitRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryHiddenProductGatewayRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryHiddenProductGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryHiddenProductGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceListData struct {
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 商品码名称

	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
	// 资源数量

	Num *string `json:"Num,omitempty" name:"Num"`
	// 0正常1隔离中2隔离3恢复中4销毁中5销毁

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 隔离时间

	IsolateNum *string `json:"IsolateNum,omitempty" name:"IsolateNum"`
	// 销毁时间

	DestroyNum *string `json:"DestroyNum,omitempty" name:"DestroyNum"`
	// 正常使用天数

	NormalNum *string `json:"NormalNum,omitempty" name:"NormalNum"`
}

type DescribeAccountWaterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 所有流水总数

		AllNum *string `json:"AllNum,omitempty" name:"AllNum"`
		// 此处返回的流水条数

		RetNum *string `json:"RetNum,omitempty" name:"RetNum"`
		// cursor，查询流水的当前游标值

		Offset *string `json:"Offset,omitempty" name:"Offset"`
		// 流水结果列表

		WaterSet []*WaterListData `json:"WaterSet,omitempty" name:"WaterSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountWaterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountWaterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillingItemsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计费项信息列表

		List []*ProductBillingItemInfoList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillingItemsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillingItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillingSubUinQuotaGatewayRequest struct {
	*tchttp.BaseRequest

	// limit分页

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子账号

	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`
}

func (r *DescribeBillingSubUinQuotaGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillingSubUinQuotaGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadVoucherListFileGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务是否完成，0未完成，1完成

		TaskFinished *uint64 `json:"TaskFinished,omitempty" name:"TaskFinished"`
		// 文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 文件名称

		FileName *string `json:"FileName,omitempty" name:"FileName"`
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDownloadVoucherListFileGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadVoucherListFileGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CouponWater struct {
	// 抵扣订单号

	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`
	// 抵扣产品

	ProductInfo *string `json:"ProductInfo,omitempty" name:"ProductInfo"`
	// 抵扣金额

	Amount *uint64 `json:"Amount,omitempty" name:"Amount"`
	// 抵扣时间

	TranTime *uint64 `json:"TranTime,omitempty" name:"TranTime"`
}

type CreateDownloadRecordGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDownloadRecordGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDownloadRecordGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByTagGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 每次获取数据量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 是否返回最大页码，建议首次或变更PageSize后才需要，用于提升接口性

	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
}

func (r *DescribeBillFeeByTagGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeByTagGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDiscountRequest struct {
	*tchttp.BaseRequest

	// 折扣类型 0用户折扣 1官网折扣

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 折扣id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteDiscountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDiscountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveSubProductStrategyRouteGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 跳过条数

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 本次返回条数

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据

		Info []*ProductStrategy `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComprehensiveSubProductStrategyRouteGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveSubProductStrategyRouteGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TotalLimit struct {
	// 天限

	DayLimit *int64 `json:"DayLimit,omitempty" name:"DayLimit"`
	// 周限

	WeekLimit *int64 `json:"WeekLimit,omitempty" name:"WeekLimit"`
	// 月限

	MonthLimit *int64 `json:"MonthLimit,omitempty" name:"MonthLimit"`
	// 总限

	AllLimit *int64 `json:"AllLimit,omitempty" name:"AllLimit"`
}

type DescribeBillDownloadListRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 开始月份

	BeginMonth *string `json:"BeginMonth,omitempty" name:"BeginMonth"`
	// 结束月份

	EndMonth *string `json:"EndMonth,omitempty" name:"EndMonth"`
}

func (r *DescribeBillDownloadListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemitRecordOpRequest struct {
	*tchttp.BaseRequest

	// 目标用户uin

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 汇款户名

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 汇款账号

	AccountNum *uint64 `json:"AccountNum,omitempty" name:"AccountNum"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRemitRecordOpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemitRecordOpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveSubProductStrategyGatewayRequest struct {
	*tchttp.BaseRequest

	// 产品代码，如果该值不传则默认查询所有产品

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品代码，如果该值不传则默认查询该产品下所有子产品

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 是否只查询第三方自定义产品策略,如果该值不传或者为false则查询全部产品策略

	IfCustom *string `json:"IfCustom,omitempty" name:"IfCustom"`
	// 状态，如果该值不传则默认查询所有状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 本次查询条数，如果该值不传则默认100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 跳过条数如果该值不传则默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetComprehensiveSubProductStrategyGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveSubProductStrategyGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceDetailRequest struct {
	*tchttp.BaseRequest

	// 数据所属用户的UIN

	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 产品编码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 页码

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 当前页条数

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListResourceDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CouponsWaterRequest struct {
	*tchttp.BaseRequest

	// 米大师appid

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 代金券批次id

	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
	// 代金券CodeId

	CodeId *string `json:"CodeId,omitempty" name:"CodeId"`
	// 代金券id

	SpId *string `json:"SpId,omitempty" name:"SpId"`
	// 用户id

	UserNum *string `json:"UserNum,omitempty" name:"UserNum"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 分页游标

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页展示记录数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *CouponsWaterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CouponsWaterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`
		// 资源列表

		List []*ResourceListData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditInvoiceInfoGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AuditInvoiceInfoGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuditInvoiceInfoGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpireDiscountRequest struct {
	*tchttp.BaseRequest

	// 折扣类型 0用户折扣 1官网折扣

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 折扣id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *ExpireDiscountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExpireDiscountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActionDetailItem struct {
	// 操作类型代码

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
	// 操作类型名称

	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`
	// 实际花费

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 费用所占百分比，两位小数

	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`
}

type DescribeDebateListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`
		// 账单欠费月份信息

		List []*DebateListData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDebateListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDebateListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditBreakdownModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditBreakdownModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditBreakdownModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceModifyLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 变配记录

		List *BspResourceBill `json:"List,omitempty" name:"List"`
		// 数据条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListResourceModifyLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceModifyLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MpInfo struct {
	// 代金券信息

	SuperCoupon []*SuperCoupon `json:"SuperCoupon,omitempty" name:"SuperCoupon"`
}

type BusinessSummaryDetailItem struct {
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
	// 实际花费

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 应付金额

	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`
	// 代金券花费

	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
	// 对比上月趋势

	Trend *SummaryTrend `json:"Trend,omitempty" name:"Trend"`
	// 标签值

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
}

type DescribeAccountListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`
		// 帐户信息列表

		List []*AccountListData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByTagGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 记录数，不传NeedRecordNum或者传的值为0，该字段值为null

		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`
		// 产品花费详情

		SummaryDetail *BusinessSummaryDetailItem `json:"SummaryDetail,omitempty" name:"SummaryDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillFeeByTagGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeByTagGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 各产品花费分布

		SummaryTotal *BusinessSummaryTotal `json:"SummaryTotal,omitempty" name:"SummaryTotal"`
		// 总花费详情

		SummaryOverview *BusinessSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPublicAccountRequest struct {
	*tchttp.BaseRequest

	// 公共账号编号，默认1（只支持填1）

	AccountId *uint64 `json:"AccountId,omitempty" name:"AccountId"`
	// 收款人户名

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 收款卡号

	AccountNum *uint64 `json:"AccountNum,omitempty" name:"AccountNum"`
	// 收款银行

	Bank *string `json:"Bank,omitempty" name:"Bank"`
	// 备注

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *ModifyPublicAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPublicAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceGatewayRequest struct {
	*tchttp.BaseRequest

	// uin

	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`
	// 资源id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 页序号，从0开始

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 业务侧的appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 录入开始时间

	StartAddTime *string `json:"StartAddTime,omitempty" name:"StartAddTime"`
	// 录入结束时间

	EndAddTime *string `json:"EndAddTime,omitempty" name:"EndAddTime"`
}

func (r *ListResourceGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsumptionResourceSummaryDetailComponentDetailItem struct {
	// 组件名称

	ComponentCodeName *string `json:"ComponentCodeName,omitempty" name:"ComponentCodeName"`
	// 组件用量

	UsedAmount *string `json:"UsedAmount,omitempty" name:"UsedAmount"`
	// 组件用量单位

	UsedAmountUnit *string `json:"UsedAmountUnit,omitempty" name:"UsedAmountUnit"`
	// 组件折后总价

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 组件折后总价占资源花费比例

	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`
}

type CouponDestroyRequest struct {
	*tchttp.BaseRequest

	// 米大师appid

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 指定代金券编号

	SpId *string `json:"SpId,omitempty" name:"SpId"`
	// codeid

	CodeId *string `json:"CodeId,omitempty" name:"CodeId"`
	// 指定代金券批次

	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
	// 用户id

	UserNum *string `json:"UserNum,omitempty" name:"UserNum"`
}

func (r *CouponDestroyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CouponDestroyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaAssignLogsListGatewayRequest struct {
	*tchttp.BaseRequest

	// 计费MidasAppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 例如：“2019-10-10 10:05:01”

	CreateStartTime *string `json:"CreateStartTime,omitempty" name:"CreateStartTime"`
	// 例如：“2019-10-12 10:05:02”

	CreateEndTime *string `json:"CreateEndTime,omitempty" name:"CreateEndTime"`
	// uin

	SearchUin *string `json:"SearchUin,omitempty" name:"SearchUin"`
	// 产品标签

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品标签

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 项目标签

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 子项目标签

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 第几页，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 一页行数，默认30

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 配额类型：sys或uin

	QuotaType *string `json:"QuotaType,omitempty" name:"QuotaType"`
}

func (r *DescribeQuotaAssignLogsListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotaAssignLogsListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetProductInfoGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetProductInfoGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetProductInfoGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolateResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GeneralData struct {
	// 通用对象数组

	Data []*GeneralObject `json:"Data,omitempty" name:"Data"`
}

type ProjectSet struct {
	// 项目目录ID

	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`
	// 项目目录名称

	OrgName *string `json:"OrgName,omitempty" name:"OrgName"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 创建者uin

	CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`
	// 创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeBspDownloadRecordListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数量，NeedRecordNum为0时返回null

		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`
		// 下载记录列表

		Records *GetDownloadListRecordItem `json:"Records,omitempty" name:"Records"`
		// 过滤条件，NeedConditionValue为0是返回null

		ConditionValue *DetailConditionValue `json:"ConditionValue,omitempty" name:"ConditionValue"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBspDownloadRecordListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBspDownloadRecordListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserAccountGatewayRequest struct {
	*tchttp.BaseRequest

	// 接入后付费账户的唯一标识

	AcctId *string `json:"AcctId,omitempty" name:"AcctId"`
	// 用户ID

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 传后付费固定额度子账户标识: CREDIT_FIXED

	SubAcctid *string `json:"SubAcctid,omitempty" name:"SubAcctid"`
	// 请求方设备，0表示PC，1表示手机

	Device *int64 `json:"Device,omitempty" name:"Device"`
	// UNIX时间戳（从格林威治时间1970年01月01日00时00分00秒起至现在的总秒数）

	Ts *int64 `json:"Ts,omitempty" name:"Ts"`
}

func (r *DescribeUserAccountGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserAccountGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBspProductStatusGatewayRequest struct {
	*tchttp.BaseRequest

	// 商品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项代码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项代码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 状态，1-生效 0-失效

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *SetBspProductStatusGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBspProductStatusGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsumptionResourceSummaryDetailTotal struct {
	// 折后总价

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 应付金额

	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`
	// 代金券支付

	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
}

type DescribeBillDownloadRecordListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数量，NeedRecordNum为0时返回null

		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`
		// 下载记录列表

		Records *GetDownloadListRecordItem `json:"Records,omitempty" name:"Records"`
		// 过滤条件，NeedConditionValue为0是返回null

		ConditionValue *DetailConditionValue `json:"ConditionValue,omitempty" name:"ConditionValue"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDownloadRecordListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadRecordListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceBillDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否展示数据，0不展示，1展示

		DisplayData *uint64 `json:"DisplayData,omitempty" name:"DisplayData"`
		// 记录数量，NeedRecordNum为0是返回null

		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`
		// 过滤条件，NeedConditionValue为0是返回null

		ConditionValue *DetailConditionValue `json:"ConditionValue,omitempty" name:"ConditionValue"`
		// 总花费详情

		Total *DetailTotal `json:"Total,omitempty" name:"Total"`
		// 组件花费详情

		Detail []*DetailDetailItem `json:"Detail,omitempty" name:"Detail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceBillDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceBillDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UinQuota struct {
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 配额键

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
	// 配额值

	QuotaValue *string `json:"QuotaValue,omitempty" name:"QuotaValue"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 计量还是计费

	Source *string `json:"Source,omitempty" name:"Source"`
	// 门阀

	ThresHold *string `json:"ThresHold,omitempty" name:"ThresHold"`
	// 子产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 产品细项码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 子产品细项码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 剩余配额值

	QuotaLeft *int64 `json:"QuotaLeft,omitempty" name:"QuotaLeft"`
	// 已使用的配额值

	QuotaUsed *int64 `json:"QuotaUsed,omitempty" name:"QuotaUsed"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品组名称

	ProductGroupName *string `json:"ProductGroupName,omitempty" name:"ProductGroupName"`
	// 产品英文名称

	ProductEngName *string `json:"ProductEngName,omitempty" name:"ProductEngName"`
	// 产品组英文名称

	ProductGroupEngName *string `json:"ProductGroupEngName,omitempty" name:"ProductGroupEngName"`
	// 子产品名称

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 子产品英文名称

	SubProductEngName *string `json:"SubProductEngName,omitempty" name:"SubProductEngName"`
	// 产品项名称

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 产品项英文名称

	BillingItemEngName *string `json:"BillingItemEngName,omitempty" name:"BillingItemEngName"`
	// 子产品项名称

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
	// 子产品项英文名称

	SubBillingItemEngName *string `json:"SubBillingItemEngName,omitempty" name:"SubBillingItemEngName"`
	// 单位名称

	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`
}

type CouponDestroyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CouponDestroyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CouponDestroyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CouponPresentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 发劵返回值

		CouponsInfoList []*CouponsInfoList `json:"CouponsInfoList,omitempty" name:"CouponsInfoList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CouponPresentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CouponPresentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDosagesRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 1 最外层(按产品和地域)汇总， 2: 按天汇总， 3 按小时汇总

	Type *string `json:"Type,omitempty" name:"Type"`
	// 地域id

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 当前页

	PageNo *uint64 `json:"PageNo,omitempty" name:"PageNo"`
	// 当前页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeDosagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDosagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByPayModeGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillFeeByPayModeGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeByPayModeGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportChangeQuotaGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportChangeQuotaGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportChangeQuotaGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuoteTypeList struct {
	// select的key

	QuotaTypeKey *string `json:"QuotaTypeKey,omitempty" name:"QuotaTypeKey"`
	// select的显示名称

	QuotaTypeValue *string `json:"QuotaTypeValue,omitempty" name:"QuotaTypeValue"`
	// select的显示英文名称

	QuotaTypeValueEng *string `json:"QuotaTypeValueEng,omitempty" name:"QuotaTypeValueEng"`
}

type SubUinQuota struct {
	// 设置的配额值

	QuotaValue *uint64 `json:"QuotaValue,omitempty" name:"QuotaValue"`
	// 剩余的配额

	QuotaLeft *int64 `json:"QuotaLeft,omitempty" name:"QuotaLeft"`
	// 已使用的配额

	QuotaUsed *int64 `json:"QuotaUsed,omitempty" name:"QuotaUsed"`
	// 子账户用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 主账户uin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// appId

	AppID *int64 `json:"AppID,omitempty" name:"AppID"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 产品细项

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 子产品细项

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品组名称

	ProductGroupName *string `json:"ProductGroupName,omitempty" name:"ProductGroupName"`
	// 产品英文名称

	ProductEngName *string `json:"ProductEngName,omitempty" name:"ProductEngName"`
	// 产品组英文名称

	ProductGroupEngName *string `json:"ProductGroupEngName,omitempty" name:"ProductGroupEngName"`
	// 子产品名称

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 子产品英文名称

	SubProductEngName *string `json:"SubProductEngName,omitempty" name:"SubProductEngName"`
	// 产品细项名称

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 产品细项英文名称

	BillingItemEngName *string `json:"BillingItemEngName,omitempty" name:"BillingItemEngName"`
	// 子产品细项名称

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
	// 子产品细项英文名称

	SubBillingItemEngName *string `json:"SubBillingItemEngName,omitempty" name:"SubBillingItemEngName"`
	// 单位名称

	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`
}

type DescribeBasicAccountListGatewayRequest struct {
	*tchttp.BaseRequest

	// 用户uin数组

	Uins []*string `json:"Uins,omitempty" name:"Uins"`
}

func (r *DescribeBasicAccountListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBasicAccountListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveSubProdctStrategyRequest struct {
	*tchttp.BaseRequest

	// 产品代码，如果该值不传则默认查询所有产品

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品代码，如果该值不传则默认查询该产品下所有子产品

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 是否只查询第三方自定义产品策略,如果该值不传或者为false则查询全部产品策略

	IfCustom *string `json:"IfCustom,omitempty" name:"IfCustom"`
	// 状态，如果该值不传则默认查询所有状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 本次查询条数，如果该值不传则默认100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 跳过条数如果该值不传则默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetComprehensiveSubProdctStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveSubProdctStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveSubProductStrategyGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 跳过条数

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 本次返回条数

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据

		Info []*ProductStrategy `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComprehensiveSubProductStrategyGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveSubProductStrategyGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BspProductDefineList struct {
	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品中文名称

	ProductChnName *string `json:"ProductChnName,omitempty" name:"ProductChnName"`
	// 产品英文名称

	ProductEngName *string `json:"ProductEngName,omitempty" name:"ProductEngName"`
	// 子产品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品中文名称

	SubProductChnName *string `json:"SubProductChnName,omitempty" name:"SubProductChnName"`
	// 子产品英文名称

	SubProductEngName *string `json:"SubProductEngName,omitempty" name:"SubProductEngName"`
	// 计费项代码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费项中文名称

	BillingItemChnName *string `json:"BillingItemChnName,omitempty" name:"BillingItemChnName"`
	// 计费项英文名称

	BillingItemEngName *string `json:"BillingItemEngName,omitempty" name:"BillingItemEngName"`
	// 计费细项代码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 计费细项中文名称

	SubBillingItemChnName *string `json:"SubBillingItemChnName,omitempty" name:"SubBillingItemChnName"`
	// 计费细项英文名称

	SubBillingItemEngName *string `json:"SubBillingItemEngName,omitempty" name:"SubBillingItemEngName"`
	// 计费细项单位

	SubBillingItemsGoodNumUnit *string `json:"SubBillingItemsGoodNumUnit,omitempty" name:"SubBillingItemsGoodNumUnit"`
	// 计费细项单位英文

	SubBillingItemsGoodNumUnitEn *string `json:"SubBillingItemsGoodNumUnitEn,omitempty" name:"SubBillingItemsGoodNumUnitEn"`
	// 四层定义状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 计量周期

	TimeUnit *uint64 `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 推量周期

	CalcUnit *uint64 `json:"CalcUnit,omitempty" name:"CalcUnit"`
	// 最近操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type ConditionFileStatus struct {
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 状态名称

	StatusName *uint64 `json:"StatusName,omitempty" name:"StatusName"`
}

type DeleteProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDealPriceRequest struct {
	*tchttp.BaseRequest

	// 商品询价参数

	ResInfo *string `json:"ResInfo,omitempty" name:"ResInfo"`
}

func (r *DescribeDealPriceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDealPriceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionSummaryDetailItem struct {
	// 该项目下产品消费明细

	Product []*BusinessSummaryDetailItem `json:"Product,omitempty" name:"Product"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 实际花费

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 应付金额

	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`
	// 代金券花费

	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
	// 对比上月趋势

	Trend *SummaryTrend `json:"Trend,omitempty" name:"Trend"`
}

type UpdateUserStrategyGatewayRequest struct {
	*tchttp.BaseRequest

	// 计量id标识

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 策略提示信息

	Prompt *string `json:"Prompt,omitempty" name:"Prompt"`
	// 限制操作，例如["create","modify"]

	ValidRules []*string `json:"ValidRules,omitempty" name:"ValidRules"`
}

func (r *UpdateUserStrategyGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateUserStrategyGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountListData struct {
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 账户名称

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 现金余额

	CashBalance *int64 `json:"CashBalance,omitempty" name:"CashBalance"`
	// 固定额度

	Fixedlimit *string `json:"Fixedlimit,omitempty" name:"Fixedlimit"`
	// 未支付金额

	UpdAmount *int64 `json:"UpdAmount,omitempty" name:"UpdAmount"`
	// 可用余额

	UsableLimit *int64 `json:"UsableLimit,omitempty" name:"UsableLimit"`
	// 已用额度

	UsedLimit *int64 `json:"UsedLimit,omitempty" name:"UsedLimit"`
}

type ProductCodeInfo struct {
	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
}

type ProductInfoList struct {
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeContractTypeFlowListGatewayRequest struct {
	*tchttp.BaseRequest

	// 合同类型ID号

	ContractTypeId *uint64 `json:"ContractTypeId,omitempty" name:"ContractTypeId"`
	// 合同状态0审核中1已邮寄2审核拒绝3作废4取消

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 管理端查询用户UIN，默认缺省查全部

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeContractTypeFlowListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContractTypeFlowListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiscountGatewayRequest struct {
	*tchttp.BaseRequest

	// 折扣类型 0用户折扣 1官网折扣

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 折扣Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 开始生效时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 失效时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 折扣值

	Discount *string `json:"Discount,omitempty" name:"Discount"`
	// 折扣条件 json字符串

	ConditionList *string `json:"ConditionList,omitempty" name:"ConditionList"`
}

func (r *ModifyDiscountGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiscountGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContractTypeFlowListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 流水列表

		List []*ContractTypeFlowData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContractTypeFlowListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContractTypeFlowListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBreakdownModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1 正常 2恢复中 3故障中

		Status *string `json:"Status,omitempty" name:"Status"`
		// 白名单列表，只有故障中才会返回

		WhiteList []*string `json:"WhiteList,omitempty" name:"WhiteList"`
		// 临时配额开关，1 关闭 2开启

		TempQuto *string `json:"TempQuto,omitempty" name:"TempQuto"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBreakdownModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBreakdownModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeTrendResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 每月花费详情

		Detail *MonthCostDetailItem `json:"Detail,omitempty" name:"Detail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillFeeTrendResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContractListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 合同列表

		List []*ContractListData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContractListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContractListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAllProductTreeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品四层定义列表

		List []*ProductRelation `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAllProductTreeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAllProductTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveProdctStrategyRequest struct {
	*tchttp.BaseRequest

	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 是否只查询第三方自定义产品策略,如果该值不传或者为false则查询全部产品策略

	IfCustom *bool `json:"IfCustom,omitempty" name:"IfCustom"`
	// 状态，如果该值不传则默认查询所有状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 本次查询条数，如果该值不传则默认100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 跳过条数如果该值不传则默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetComprehensiveProdctStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveProdctStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddComprehensiveStrategyRouteGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddComprehensiveStrategyRouteGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddComprehensiveStrategyRouteGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInvoiceListRequest struct {
	*tchttp.BaseRequest

	// 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 开始时间 形如'Y-m-d H:m:s'

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间 形如'Y-m-d H:m:s'

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 根据申请时间排序,0:降序,1:升序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 搜索的状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 搜索的uin信息,如果是租户端,SearchUin=Uin

	SearchUin *int64 `json:"SearchUin,omitempty" name:"SearchUin"`
}

func (r *DescribeInvoiceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInvoiceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInvoiceListGatewayRequest struct {
	*tchttp.BaseRequest

	// 分页偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页限制数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 开始时间 形如'Y-m-d H:m:s'

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间 形如'Y-m-d H:m:s'

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 根据申请时间排序,0:降序,1:升序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 搜索的状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 搜索的uin信息,如果是租户端,SearchUin=Uin

	SearchUin *int64 `json:"SearchUin,omitempty" name:"SearchUin"`
}

func (r *DescribeInvoiceListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInvoiceListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBspDownloadRecordListGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 每次获取数据量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 是否需要记录条数，0需要，1不需要，默认不需要

	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
	// 排序规则

	Sort *GetDownloadListSort `json:"Sort,omitempty" name:"Sort"`
	// 只支持FIleIds，FileTypes，Status三种过滤条件

	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
	// 是否需要条件取值（0不需要/1需要，默认为0）

	NeedConditionValue *string `json:"NeedConditionValue,omitempty" name:"NeedConditionValue"`
}

func (r *DescribeBspDownloadRecordListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBspDownloadRecordListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDownloadListSort struct {
	// 可选desc或者asc

	DownloadTime *string `json:"DownloadTime,omitempty" name:"DownloadTime"`
	// 可选desc或者asc

	GenerateTime *string `json:"GenerateTime,omitempty" name:"GenerateTime"`
}

type DescribeGoodsPriceRequest struct {
	*tchttp.BaseRequest

	// 商品询价参数

	ResInfo []*Goods `json:"ResInfo,omitempty" name:"ResInfo"`
}

func (r *DescribeGoodsPriceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGoodsPriceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillingUinQuotaGatewayRequest struct {
	*tchttp.BaseRequest

	// appid

	QuotaAppId *uint64 `json:"QuotaAppId,omitempty" name:"QuotaAppId"`
	// 分页limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页码

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 产品一层

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

func (r *DescribeBillingUinQuotaGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillingUinQuotaGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddComprehensiveStrategyGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddComprehensiveStrategyGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddComprehensiveStrategyGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDiscountGatewayRequest struct {
	*tchttp.BaseRequest

	// 折扣类型 0用户折扣 1官网折扣

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 折扣id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteDiscountGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDiscountGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTenantSubUinQuotasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总行数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 列表数据

		List *SubUinQuotaList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTenantSubUinQuotasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTenantSubUinQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDebateListRequest struct {
	*tchttp.BaseRequest

	// 管理端查询用户UIN，默认缺省查全部

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 欠费月份，0就是查欠账月份大于等于0

	UnpayDebtNum *uint64 `json:"UnpayDebtNum,omitempty" name:"UnpayDebtNum"`
	// 针对DebtNum的比较条件， E: 等于， G: 大于，GE: 大于等于

	NumSort *string `json:"NumSort,omitempty" name:"NumSort"`
}

func (r *DescribeDebateListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDebateListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HiddenProduct struct {
	// 一层商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 二层子商品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 三层计费项

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 四层计费细项

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 0隐藏，1显示

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 0来自计量 1来自计费

	IfBilling *uint64 `json:"IfBilling,omitempty" name:"IfBilling"`
}

type VoucherParams struct {
	// 用户ID

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 指定查询状态

	SpStatus *string `json:"SpStatus,omitempty" name:"SpStatus"`
	// 指定代金券名称

	SpName *string `json:"SpName,omitempty" name:"SpName"`
	// 指定代金券编号

	SpId *string `json:"SpId,omitempty" name:"SpId"`
	// 指定代金券批次

	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
	// 有效开始时间

	ValidBeginTime *string `json:"ValidBeginTime,omitempty" name:"ValidBeginTime"`
	// 有效结束时间

	ValidEndTime *string `json:"ValidEndTime,omitempty" name:"ValidEndTime"`
	// 发放开始时间

	PresentBeginTime *string `json:"PresentBeginTime,omitempty" name:"PresentBeginTime"`
	// 发放结束时间

	PresentEndTime *string `json:"PresentEndTime,omitempty" name:"PresentEndTime"`
	// 指定排序字段：begin_time、end_time。

	SortField *string `json:"SortField,omitempty" name:"SortField"`
	// 指定升序降序：desc、asc。

	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

type DescribeDownloadWithAuthRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
}

func (r *DescribeDownloadWithAuthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadWithAuthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveQuotaRequest struct {
	*tchttp.BaseRequest

	// 用户uin。传空字符串就新增或编辑system默认配额

	QuotaUin *string `json:"QuotaUin,omitempty" name:"QuotaUin"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 要新增或更新的数据数组

	SaveDataArray []*SaveQuotaDataArray `json:"SaveDataArray,omitempty" name:"SaveDataArray"`
}

func (r *SaveQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BspResourceBill struct {
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 用户appId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 资源id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 计费单位

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 计费用时

	TimeSpan *float64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 数量

	UsedAmount *float64 `json:"UsedAmount,omitempty" name:"UsedAmount"`
	// 单位

	UsedAmountUnit *string `json:"UsedAmountUnit,omitempty" name:"UsedAmountUnit"`
	// 产品大类

	ProductClass *string `json:"ProductClass,omitempty" name:"ProductClass"`
	// 产品编码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 子产品编码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品名称

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 地域id

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 地域

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 可用区id

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 计费项编码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费项名称

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 子计费项编码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 子计费项名称

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
	// 资源的规格配置

	ResourceDetail *string `json:"ResourceDetail,omitempty" name:"ResourceDetail"`
}

type DescribeRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域明细

		RegionDetail *RegionDetail `json:"RegionDetail,omitempty" name:"RegionDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemitRecordRequest struct {
	*tchttp.BaseRequest

	// 目标用户uin

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 汇款户名

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 汇款账号

	AccountNum *uint64 `json:"AccountNum,omitempty" name:"AccountNum"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 状态 0：处理中 1：成功 2：失败

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeRemitRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemitRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWaterInfoRequest struct {
	*tchttp.BaseRequest

	// 由计平统一分配

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 用户id

	UserNum *string `json:"UserNum,omitempty" name:"UserNum"`
	// 资源id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 地域id

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 0:后付费；1:预付费

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 游标

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 记录数目,最小值为1,最大值为10.超过范围则以最值代替

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 固定传: query_water_info

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeWaterInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWaterInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载地址列表

		DownloadList *GetPkgDownloadListItem `json:"DownloadList,omitempty" name:"DownloadList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDownloadListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByPayModeGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 付费模式花费分布

		SummaryOverview []*PayModeSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByPayModeGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByPayModeGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTempQuotaListRequest struct {
	*tchttp.BaseRequest

	// 页数

	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`
	// 一页数目

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 产品一层定义

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品二层定义

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 产品三层定义

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 产品四层定义

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// uin

	QueryUin *string `json:"QueryUin,omitempty" name:"QueryUin"`
}

func (r *DescribeTempQuotaListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTempQuotaListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByPayModeGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 计费模式花费详情

		SummaryDetail []*PayModeSummaryDetailItem `json:"SummaryDetail,omitempty" name:"SummaryDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillFeeByPayModeGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeByPayModeGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConditionBusiness struct {
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
}

type DescribeResourceListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`
		// 资源列表

		List []*ResourceListData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyComprehensiveStrategyRouteGatewayRequest struct {
	*tchttp.BaseRequest

	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品code

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 新购是否同步发货 默认0

	PurchaseSyncDelivery *uint64 `json:"PurchaseSyncDelivery,omitempty" name:"PurchaseSyncDelivery"`
	// 变配是否同步发货 默认0

	ModifySyncDelivery *uint64 `json:"ModifySyncDelivery,omitempty" name:"ModifySyncDelivery"`
	// 接口名

	InterfaceNames *InterfaceNames `json:"InterfaceNames,omitempty" name:"InterfaceNames"`
	// url名

	Region2interface *string `json:"Region2interface,omitempty" name:"Region2interface"`
	// 状态 1生效 2不生效

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 修改的是计量产品还是计费产品，取值范围计量（bsp）计费（billing），不填写则默认bsp

	Source *string `json:"Source,omitempty" name:"Source"`
}

func (r *ModifyComprehensiveStrategyRouteGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyComprehensiveStrategyRouteGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBillingProductCodeRequest struct {
	*tchttp.BaseRequest
}

func (r *GetBillingProductCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBillingProductCodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductListRequest struct {
	*tchttp.BaseRequest

	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 页索引

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 是否生效.0未生效 1生效

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *GetProductListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DealListData struct {
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 商品码名称

	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
	// 子商品码名称

	SubProductCodeName *string `json:"SubProductCodeName,omitempty" name:"SubProductCodeName"`
	// 类型，新购续费变配等

	DealAction *string `json:"DealAction,omitempty" name:"DealAction"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 提单人

	Payer *string `json:"Payer,omitempty" name:"Payer"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 大订单号

	BigDealId *uint64 `json:"BigDealId,omitempty" name:"BigDealId"`
	// 订单详情，jason字符串

	GoodsDetail *string `json:"GoodsDetail,omitempty" name:"GoodsDetail"`
	// 付费模式

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 订单数量

	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 业务返回资源ID详情，jason字符串

	TaskDetail *string `json:"TaskDetail,omitempty" name:"TaskDetail"`
	// 类型中文名称

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 状态名称

	StatusName *string `json:"StatusName,omitempty" name:"StatusName"`
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 升配公式

	Formula *string `json:"Formula,omitempty" name:"Formula"`
	// 区域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 折扣，范围：0~1

	Discount *float64 `json:"Discount,omitempty" name:"Discount"`
	// 折扣金额（因折扣减少的金额）

	DiscountCost *uint64 `json:"DiscountCost,omitempty" name:"DiscountCost"`
	// 订单价格详情（包括折扣详情），json字符串

	GoodsPrice *string `json:"GoodsPrice,omitempty" name:"GoodsPrice"`
	// 订单号

	DealName *uint64 `json:"DealName,omitempty" name:"DealName"`
	// 代金券ID

	VoucherId *string `json:"VoucherId,omitempty" name:"VoucherId"`
	// 单价

	Price *uint64 `json:"Price,omitempty" name:"Price"`
	// 金额总计

	TotalCost *uint64 `json:"TotalCost,omitempty" name:"TotalCost"`
	// 付款人

	ProviderOwnerUin *string `json:"ProviderOwnerUin,omitempty" name:"ProviderOwnerUin"`
	// 创建人

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 实际金额

	RealTotalCost *uint64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 代金券金额

	VoucherDecline *uint64 `json:"VoucherDecline,omitempty" name:"VoucherDecline"`
	// 代金券交易ID

	VoucherTradeId *string `json:"VoucherTradeId,omitempty" name:"VoucherTradeId"`
}

type GetComprehensiveProductStrategyRouteGatewayRequest struct {
	*tchttp.BaseRequest

	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 是否只查询第三方自定义产品策略,如果该值不传或者为false则查询全部产品策略

	IfCustom *bool `json:"IfCustom,omitempty" name:"IfCustom"`
	// 状态，如果该值不传则默认查询所有状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 本次查询条数，如果该值不传则默认100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 跳过条数如果该值不传则默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetComprehensiveProductStrategyRouteGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveProductStrategyRouteGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDealPriceGatewayRequest struct {
	*tchttp.BaseRequest

	// 商品询价参数

	ResInfo *string `json:"ResInfo,omitempty" name:"ResInfo"`
}

func (r *DescribeDealPriceGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDealPriceGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGeneralDataRequest struct {
	*tchttp.BaseRequest

	// 查询类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 查询条件

	Conditions []*GeneralObject `json:"Conditions,omitempty" name:"Conditions"`
	// 返回的字段

	Filter []*string `json:"Filter,omitempty" name:"Filter"`
	// 排序字段

	SortField *string `json:"SortField,omitempty" name:"SortField"`
	// 排序方式 “asc”或“desc”

	SortType *string `json:"SortType,omitempty" name:"SortType"`
	// 分页，起始值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页，数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGeneralDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsolidatedBillDownloadUrlGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 子uin列表

	ChildUins []*string `json:"ChildUins,omitempty" name:"ChildUins"`
	// 内部操作者

	InnerOperator *string `json:"InnerOperator,omitempty" name:"InnerOperator"`
}

func (r *DescribeConsolidatedBillDownloadUrlGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsolidatedBillDownloadUrlGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetQuotaGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户额度详细

		List []*QuotaDetailList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetQuotaGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetQuotaGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveUserDiscountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ID

		// 条数

		// 是否成功

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveUserDiscountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveUserDiscountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddProductGatewayRequest struct {
	*tchttp.BaseRequest

	// 产品大类中文名

	ProductGroupName *string `json:"ProductGroupName,omitempty" name:"ProductGroupName"`
	// 产品大类英文名

	ProductGroupEngName *string `json:"ProductGroupEngName,omitempty" name:"ProductGroupEngName"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品中文名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品英文名

	ProductEngName *string `json:"ProductEngName,omitempty" name:"ProductEngName"`
	// 子产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品中文名

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 子产品英文名

	SubProductEngName *string `json:"SubProductEngName,omitempty" name:"SubProductEngName"`
	// 计费项代码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费项中文名

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 计费项英文名

	BillingItemEngName *string `json:"BillingItemEngName,omitempty" name:"BillingItemEngName"`
	// 计费细项代码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 计费细项中文名

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
	// 计费细项英文名

	SubBillingItemEngName *string `json:"SubBillingItemEngName,omitempty" name:"SubBillingItemEngName"`
	// 单位中文名

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 单位英文名

	UnitEng *string `json:"UnitEng,omitempty" name:"UnitEng"`
	// 计费周期

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 结算周期

	CalcUnit *string `json:"CalcUnit,omitempty" name:"CalcUnit"`
}

func (r *AddProductGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddProductGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemitRecordOpGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 列表

		List *RemitRecordData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemitRecordOpGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemitRecordOpGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductStrategyGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductStrategyGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductStrategyGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryRecoverUinRequest struct {
	*tchttp.BaseRequest

	// 跳过条数

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 本页返回条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 本次需要查询的uin，支持模糊查询

	SearchUin *string `json:"SearchUin,omitempty" name:"SearchUin"`
}

func (r *QueryRecoverUinRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryRecoverUinRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BusinessSummaryTotal struct {
	// 总花费

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
}

type ConditionPayMode struct {
	// 付费模式

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 付费模式名称

	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`
}

type AddProductRequest struct {
	*tchttp.BaseRequest

	// 产品大类中文名

	ProductGroupName *string `json:"ProductGroupName,omitempty" name:"ProductGroupName"`
	// 产品大类英文名

	ProductGroupEngName *string `json:"ProductGroupEngName,omitempty" name:"ProductGroupEngName"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品中文名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品英文名

	ProductEngName *string `json:"ProductEngName,omitempty" name:"ProductEngName"`
	// 子产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品中文名

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 子产品英文名

	SubProductEngName *string `json:"SubProductEngName,omitempty" name:"SubProductEngName"`
	// 计费项代码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费项中文名

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 计费项英文名

	BillingItemEngName *string `json:"BillingItemEngName,omitempty" name:"BillingItemEngName"`
	// 计费细项代码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 计费细项中文名

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
	// 计费细项英文名

	SubBillingItemEngName *string `json:"SubBillingItemEngName,omitempty" name:"SubBillingItemEngName"`
	// 单位中文名

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 单位英文名

	UnitEng *string `json:"UnitEng,omitempty" name:"UnitEng"`
	// 计费周期

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 结算周期

	CalcUnit *string `json:"CalcUnit,omitempty" name:"CalcUnit"`
}

func (r *AddProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceRequest struct {
	*tchttp.BaseRequest

	// uin

	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`
	// 资源id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 页序号，从0开始

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 业务侧的appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 录入开始时间

	StartAddTime *string `json:"StartAddTime,omitempty" name:"StartAddTime"`
	// 录入结束时间

	EndAddTime *string `json:"EndAddTime,omitempty" name:"EndAddTime"`
	// 资源配置变更查询时间范围，开始时间

	StartModifyTime *string `json:"StartModifyTime,omitempty" name:"StartModifyTime"`
	// 资源配置变更查询时间范围，截止时间

	EndModifyTime *string `json:"EndModifyTime,omitempty" name:"EndModifyTime"`
	// 资源销毁查询时间范围，开始时间

	StartIsolatedTime *string `json:"StartIsolatedTime,omitempty" name:"StartIsolatedTime"`
	// 资源销毁查询时间范围，截止时间

	EndIsolatedTime *string `json:"EndIsolatedTime,omitempty" name:"EndIsolatedTime"`
}

func (r *ListResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionSummaryOverviewItem struct {
	// 地域ID

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 实际花费

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 费用所占百分比，两位小数

	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`
}

type AddQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubProductDefineList struct {
	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品中文名称

	ProductChnName *string `json:"ProductChnName,omitempty" name:"ProductChnName"`
	// 产品英文名称

	ProductEngName *string `json:"ProductEngName,omitempty" name:"ProductEngName"`
	// 子产品英文代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品中文名称

	ChnName *string `json:"ChnName,omitempty" name:"ChnName"`
	// 子产品英文名称

	EngName *string `json:"EngName,omitempty" name:"EngName"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type EditBreakdownModeRequest struct {
	*tchttp.BaseRequest

	// 故障模式操作，2开启恢复模式 3开启故障模式

	Status *string `json:"Status,omitempty" name:"Status"`
	// 白名单列表，如果不传则不修改，否则删除原有数据，并插入所传的列表

	WhiteList []*string `json:"WhiteList,omitempty" name:"WhiteList"`
	// 临时配额开关，1关闭 2开启

	TempQuto *string `json:"TempQuto,omitempty" name:"TempQuto"`
}

func (r *EditBreakdownModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditBreakdownModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsolidatedBillDownloadUrlRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 子uin列表

	ChildUins []*string `json:"ChildUins,omitempty" name:"ChildUins"`
	// 内部操作者

	InnerOperator *string `json:"InnerOperator,omitempty" name:"InnerOperator"`
}

func (r *DescribeConsolidatedBillDownloadUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsolidatedBillDownloadUrlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductsGatewayRequest struct {
	*tchttp.BaseRequest

	// 多个产品代码,例如['p_cvm','p_cvm2'...]

	ProductCodes []*string `json:"ProductCodes,omitempty" name:"ProductCodes"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 数据是否过滤隐藏数据,true表示需要

	ExceptHidden *bool `json:"ExceptHidden,omitempty" name:"ExceptHidden"`
	// 过滤隐藏计量或计费数据标志1计费0计量,不传为全部隐藏

	IfBilling *int64 `json:"IfBilling,omitempty" name:"IfBilling"`
}

func (r *DescribeProductsGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductsGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiscountPara struct {
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 开始生效时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 失效时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 折扣备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 折扣条件

	ConditionList *string `json:"ConditionList,omitempty" name:"ConditionList"`
	// 用户uin

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 折扣百分值

	Discount *float64 `json:"Discount,omitempty" name:"Discount"`
	// 折扣状态 0未生效 1已生效 3废弃

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type GetQuotaRequest struct {
	*tchttp.BaseRequest

	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 商品配额key，由四层定义某几层以及"#"号拼接而成，如p_cvm#sp_cvm_vself2#v_cvm_cpu#sv_cvm_cpu_vself2

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
	// AccountUin

	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`
}

func (r *GetQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceModifyLogGatewayRequest struct {
	*tchttp.BaseRequest

	// 数据所属用户的UIN

	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 产品编码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品细项编码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 页码

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 当前页条数

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListResourceModifyLogGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceModifyLogGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 资源汇总详情

		Data []*SummaryByResourceDataItem `json:"Data,omitempty" name:"Data"`
		// 资源汇总花费详情

		Total *SummaryByResourceTotal `json:"Total,omitempty" name:"Total"`
		// 记录数量，NeedRecordNum为0时该值为null

		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`
		// 过滤条件，NeedConditionValue为0时该值为null

		ConditionValue *SummaryByResourceConditionValue `json:"ConditionValue,omitempty" name:"ConditionValue"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillTrendByMonthRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *uint64 `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 只支持TimeRange

	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeBillTrendByMonthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillTrendByMonthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDealPriceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDealPriceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDealPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddComprehensiveStrategyRequest struct {
	*tchttp.BaseRequest

	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品code

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 付费模式

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 时间周期 如 1,2,4,6

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 结算周期 如 3,4

	CalcUnit *string `json:"CalcUnit,omitempty" name:"CalcUnit"`
	// 交易类型 如 1,2,3,5

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
	// 计费模式 如1,2,3,5

	BillingType *string `json:"BillingType,omitempty" name:"BillingType"`
	// 新购是否同步发货 默认0

	PurchaseSyncDelivery *uint64 `json:"PurchaseSyncDelivery,omitempty" name:"PurchaseSyncDelivery"`
	// 变配是否同步发货 默认0

	ModifySyncDelivery *uint64 `json:"ModifySyncDelivery,omitempty" name:"ModifySyncDelivery"`
	// 资源回收类型 默认1

	ResourceRecycleType *uint64 `json:"ResourceRecycleType,omitempty" name:"ResourceRecycleType"`
	// 接口名

	InterfaceNames *InterfaceNames `json:"InterfaceNames,omitempty" name:"InterfaceNames"`
	// url名

	Region2interface *string `json:"Region2interface,omitempty" name:"Region2interface"`
	// 定价类型,2:线性价格

	PriceType *uint64 `json:"PriceType,omitempty" name:"PriceType"`
}

func (r *AddComprehensiveStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddComprehensiveStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CouponInfoRequest struct {
	*tchttp.BaseRequest

	// 米大师appid

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 指定查询状态

	SpStatus *string `json:"SpStatus,omitempty" name:"SpStatus"`
	// 指定代金券名称

	SpName *string `json:"SpName,omitempty" name:"SpName"`
	// 指定代金券编号

	SpId *string `json:"SpId,omitempty" name:"SpId"`
	// 指定代金券批次

	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
	// 有效开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 有效结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 发放开始时间

	CreateBeginTime *string `json:"CreateBeginTime,omitempty" name:"CreateBeginTime"`
	// 发放结束时间

	CreateEndTime *string `json:"CreateEndTime,omitempty" name:"CreateEndTime"`
	// 指定排序字段

	SortField *string `json:"SortField,omitempty" name:"SortField"`
	// 指定升序降序：desc、asc

	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
	// 指定查询的每页最大记录数，默认10个

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 指定当前查询第几页，默认第1页

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 用户id

	UserNum *string `json:"UserNum,omitempty" name:"UserNum"`
	// 最大过期时间

	LargestEndTime *string `json:"LargestEndTime,omitempty" name:"LargestEndTime"`
}

func (r *CouponInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CouponInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBspUinQuotaRequest struct {
	*tchttp.BaseRequest

	// appid

	QuotaAppId *uint64 `json:"QuotaAppId,omitempty" name:"QuotaAppId"`
	// 分页limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页码

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 产品一层

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

func (r *GetBspUinQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBspUinQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceSummaryDetailTotal struct {
	// 折后总价

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 应付金额

	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`
	// 代金券支付总价

	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
}

type DescribeBillDownloadCommonSummaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务是否已完成（0未完成/1已完成）

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 文件名称，当任务已完成时返回该字段

		Name *string `json:"Name,omitempty" name:"Name"`
		// 下载链接，当任务已完成时返回该字段

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDownloadCommonSummaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadCommonSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPublicAccountGatewayRequest struct {
	*tchttp.BaseRequest

	// 公共帐户ID默认写死1

	AccountId *uint64 `json:"AccountId,omitempty" name:"AccountId"`
}

func (r *QueryPublicAccountGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPublicAccountGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditBreakdownModeGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditBreakdownModeGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditBreakdownModeGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBspProductStatusGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetBspProductStatusGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBspProductStatusGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BigDealListData struct {
	// 大订单号

	BigDealId *int64 `json:"BigDealId,omitempty" name:"BigDealId"`
	// 用户唯一标识

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 产品名称

	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
	// 子产品名称

	SubProductCodeName *string `json:"SubProductCodeName,omitempty" name:"SubProductCodeName"`
	// 交易动作

	DealAction *string `json:"DealAction,omitempty" name:"DealAction"`
	// 动作名称

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 订单过期时间

	OverdueTime *string `json:"OverdueTime,omitempty" name:"OverdueTime"`
	// 付款时间

	PayEndTime *string `json:"PayEndTime,omitempty" name:"PayEndTime"`
	// 创建人uin

	Payer *string `json:"Payer,omitempty" name:"Payer"`
	// 状态 1待支付 2已支付 3发货中 4发货成功 5发货失败 6已退款 7取消 100删除

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 金额总计

	TotalCost *uint64 `json:"TotalCost,omitempty" name:"TotalCost"`
	// 代金券金额

	VoucherDecline *uint64 `json:"VoucherDecline,omitempty" name:"VoucherDecline"`
	// 实付金额总计

	RealTotalCost *uint64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 付款人

	ProviderOwnerUin *string `json:"ProviderOwnerUin,omitempty" name:"ProviderOwnerUin"`
	// 状态名称

	StatusName *string `json:"StatusName,omitempty" name:"StatusName"`
	// 小订单列表

	DealList []*DealListData `json:"DealList,omitempty" name:"DealList"`
	// 折扣金额（因折扣而减少的金额）

	DiscountCost *uint64 `json:"DiscountCost,omitempty" name:"DiscountCost"`
	// 创建者

	Creater *string `json:"Creater,omitempty" name:"Creater"`
}

type DiscardProductCode struct {
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项代码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项代码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 废弃层级

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type DescribeDownloadWithOperatorGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDownloadWithOperatorGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadWithOperatorGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadDosagesGatewayRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 地域id

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *DownloadDosagesGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadDosagesGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BusinessSummaryOverviewItem struct {
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
	// 实际花费

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 费用所占百分比，两位小数

	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type ModifyComprehensiveStrategyGatewayRequest struct {
	*tchttp.BaseRequest

	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品code

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 付费模式

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 时间周期 如 1,2,4,6

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 结算周期 如 3,4

	CalcUnit *string `json:"CalcUnit,omitempty" name:"CalcUnit"`
	// 交易类型 如 1,2,3,5

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
	// 计费模式 如1,2,3,5

	BillingType *string `json:"BillingType,omitempty" name:"BillingType"`
	// 新购是否同步发货 默认0

	PurchaseSyncDelivery *uint64 `json:"PurchaseSyncDelivery,omitempty" name:"PurchaseSyncDelivery"`
	// 变配是否同步发货 默认0

	ModifySyncDelivery *uint64 `json:"ModifySyncDelivery,omitempty" name:"ModifySyncDelivery"`
	// 定价类型,2:线性价格

	PriceType *uint64 `json:"PriceType,omitempty" name:"PriceType"`
	// 资源回收类型 默认1

	ResourceRecycleType *uint64 `json:"ResourceRecycleType,omitempty" name:"ResourceRecycleType"`
	// 接口名

	InterfaceNames *InterfaceNames `json:"InterfaceNames,omitempty" name:"InterfaceNames"`
	// url名

	Region2interface *string `json:"Region2interface,omitempty" name:"Region2interface"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyComprehensiveStrategyGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyComprehensiveStrategyGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseResourceRequest struct {
	*tchttp.BaseRequest

	// 目标用户

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

func (r *ReleaseResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PropertyFieldsValuesDataArray struct {
	// 对应fieldCode的key

	FieldCode *string `json:"FieldCode,omitempty" name:"FieldCode"`
	// 对应fieldCode的值

	FieldValue *string `json:"FieldValue,omitempty" name:"FieldValue"`
}

type ExportDealListRequest struct {
	*tchttp.BaseRequest

	// 订单号

	DealId *uint64 `json:"DealId,omitempty" name:"DealId"`
	// 大订单号

	BigDealId *uint64 `json:"BigDealId,omitempty" name:"BigDealId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 提单人

	Payer *uint64 `json:"Payer,omitempty" name:"Payer"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 订单类型

	DealAction *string `json:"DealAction,omitempty" name:"DealAction"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 排序asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 大订单或小订单

	DealOrBigDealId *string `json:"DealOrBigDealId,omitempty" name:"DealOrBigDealId"`
	// 目标uin

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 付费模式，0后付费/1预付费

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// 产品名称

	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
	// 动作名称 隔离 销毁 返还 新购 续费 变配

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
}

func (r *ExportDealListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDealListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyContractTypeGatewayRequest struct {
	*tchttp.BaseRequest

	// 合同类型唯一ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 类型名称

	ContractTypeName *string `json:"ContractTypeName,omitempty" name:"ContractTypeName"`
	// 下载链接

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
}

func (r *ModifyContractTypeGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyContractTypeGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceSummaryDetailRealTotalCostTrend struct {
	// 组件历史月份花费详情

	Detail []*ResourceSummaryDetailRealTotalCostTrendDetailItem `json:"Detail,omitempty" name:"Detail"`
	// 历史月份平均值

	Average *string `json:"Average,omitempty" name:"Average"`
}

type ExportDealListGatewayRequest struct {
	*tchttp.BaseRequest

	// 订单号

	DealId *uint64 `json:"DealId,omitempty" name:"DealId"`
	// 大订单号

	BigDealId *uint64 `json:"BigDealId,omitempty" name:"BigDealId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 提单人

	Payer *uint64 `json:"Payer,omitempty" name:"Payer"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 订单类型

	DealAction *string `json:"DealAction,omitempty" name:"DealAction"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 排序asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 大订单或小订单

	DealOrBigDealId *string `json:"DealOrBigDealId,omitempty" name:"DealOrBigDealId"`
	// 目标uin

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 付费模式，0后付费/1预付费

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// 产品名称

	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
	// 动作名称 隔离 销毁 返还 新购 续费 变配

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
}

func (r *ExportDealListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDealListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceSummaryDownloadGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件下载地址

		DownloadUrl *BspResourceBill `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListResourceSummaryDownloadGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceSummaryDownloadGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditRemitRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AuditRemitRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuditRemitRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SummaryByResourceTotal struct {
	// 应付金额

	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`
	// 代金券支付总价

	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
	// 实际花费

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
}

type QueryPublicAccountRequest struct {
	*tchttp.BaseRequest

	// 公共帐户ID默认写死1

	AccountId *uint64 `json:"AccountId,omitempty" name:"AccountId"`
}

func (r *QueryPublicAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPublicAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConditionFileType struct {
	// 文件类型

	FileType *string `json:"FileType,omitempty" name:"FileType"`
	// 文件类型名称

	FileTypeName *string `json:"FileTypeName,omitempty" name:"FileTypeName"`
}

type ContractTypeListData struct {
	// 合同类型名称

	ContractTypeName *string `json:"ContractTypeName,omitempty" name:"ContractTypeName"`
	// 状态0废弃1正常

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 下载地址

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
	// 合同类型ID

	ContractTypeId *uint64 `json:"ContractTypeId,omitempty" name:"ContractTypeId"`
	// 备注信息

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type CreateDownloadRecordGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 文件ID

	FileId *string `json:"FileId,omitempty" name:"FileId"`
}

func (r *CreateDownloadRecordGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDownloadRecordGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubProductsRequest struct {
	*tchttp.BaseRequest

	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 子产品代码,例如['sp_cvm','sp_cvm_2']

	SubProductCodes []*string `json:"SubProductCodes,omitempty" name:"SubProductCodes"`
	// 产品代码或名称

	SearchProduct *string `json:"SearchProduct,omitempty" name:"SearchProduct"`
	// 页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 状态，1-生效 0-失效

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeSubProductsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubProductsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceSummaryDetailRealTotalCostTrendDetailItem struct {
	// 月份，如2018-08

	Month *string `json:"Month,omitempty" name:"Month"`
	// 资源折后总价

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
}

type DescribeProductsRequest struct {
	*tchttp.BaseRequest

	// 多个产品代码,例如['p_cvm','p_cvm2'...]

	ProductCodes []*string `json:"ProductCodes,omitempty" name:"ProductCodes"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
}

func (r *DescribeProductsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaTypeListRequest struct {
	*tchttp.BaseRequest

	// midas appid

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
}

func (r *DescribeQuotaTypeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotaTypeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReconciliationGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 调账原因：计费错误-billing_error 产品故障-business_error 内部核销-internal_write_off 其它原因-other

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 调账类型 add-补偿/minus-扣费

	AdjustType *string `json:"AdjustType,omitempty" name:"AdjustType"`
	// 调账月份： "2019-05"

	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`
	// 调账金额：元

	Amount *string `json:"Amount,omitempty" name:"Amount"`
	// 调账订单号

	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`
}

func (r *ModifyReconciliationGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReconciliationGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductCodeAndName struct {
	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type SetProductStatusRequest struct {
	*tchttp.BaseRequest

	// 商品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项代码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项代码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 状态，1-生效 0-失效

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
}

func (r *SetProductStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetProductStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PayModeSummaryOverviewItem struct {
	// 付费模式代码

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 付费模式名称

	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`
	// 实际花费

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 费用所占百分比，两位小数

	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`
	// 操作详情

	Detail []*ActionDetailItem `json:"Detail,omitempty" name:"Detail"`
}

type DeleteProductRequest struct {
	*tchttp.BaseRequest

	// 主键ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDiscardProductGatewayRequest struct {
	*tchttp.BaseRequest

	// 废弃层级：1-产品码

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *GetDiscardProductGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDiscardProductGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReconciliationDataItem struct {
	// 调账用户UIN

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 调账类型 add-补偿/minus-扣费

	AdjustType *string `json:"AdjustType,omitempty" name:"AdjustType"`
	// 调账总金额。元

	Amount *string `json:"Amount,omitempty" name:"Amount"`
	// 调账原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 调账月份

	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`
	// 调账单号

	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`
	// 调账申请时间

	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`
	// 调账申请人

	ApplyUin *string `json:"ApplyUin,omitempty" name:"ApplyUin"`
	// 调账商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 调账子商品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 调账计费项

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 调账计费细项

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 调账商品码名称

	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
	// 调账子商品码名称

	SubProductCodeName *string `json:"SubProductCodeName,omitempty" name:"SubProductCodeName"`
	// 调账计费项名称

	BillingItemCodeName *string `json:"BillingItemCodeName,omitempty" name:"BillingItemCodeName"`
	// 调账计费细项名称

	SubBillingItemCodeName *string `json:"SubBillingItemCodeName,omitempty" name:"SubBillingItemCodeName"`
	// 调账状态，1为成功，0为失败

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DescribeConsolidatedBillDownloadUrlGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务是否完成，0未完成，1完成

		TaskFinished *uint64 `json:"TaskFinished,omitempty" name:"TaskFinished"`
		// 文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 文件名称

		FileName *string `json:"FileName,omitempty" name:"FileName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsolidatedBillDownloadUrlGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsolidatedBillDownloadUrlGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLifecycleFlowListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 流水列表

		List []*DescribeLifecycleFlowList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLifecycleFlowListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLifecycleFlowListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductRelationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 产品定义列表

		List []*ProductDefineList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductRelationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductRelationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDiscardProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 废弃的产品码List

		List []*DiscardProductCode `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDiscardProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDiscardProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveCommonDiscountGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 折扣id

		// 条数

		// 是否成功

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveCommonDiscountGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveCommonDiscountGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserStrategyGatewayRequest struct {
	*tchttp.BaseRequest

	// 删除计量ID列表

	IdList []*int64 `json:"IdList,omitempty" name:"IdList"`
}

func (r *DeleteUserStrategyGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserStrategyGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLifecycleFlowListGatewayRequest struct {
	*tchttp.BaseRequest

	// 要查询的UIN

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 操作者UIN

	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`
	// 偏移量

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *string `json:"Limit,omitempty" name:"Limit"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 操作

	OperateAction *string `json:"OperateAction,omitempty" name:"OperateAction"`
}

func (r *DescribeLifecycleFlowListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLifecycleFlowListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserAccountRequest struct {
	*tchttp.BaseRequest

	// 接入后付费账户的唯一标识

	AcctId *string `json:"AcctId,omitempty" name:"AcctId"`
	// 用户ID

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 传后付费固定额度子账户标识: CREDIT_FIXED

	SubAcctid *string `json:"SubAcctid,omitempty" name:"SubAcctid"`
	// 请求方设备，0表示PC，1表示手机

	Device *int64 `json:"Device,omitempty" name:"Device"`
	// UNIX时间戳（从格林威治时间1970年01月01日00时00分00秒起至现在的总秒数）

	Ts *int64 `json:"Ts,omitempty" name:"Ts"`
}

func (r *DescribeUserAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceDetailGatewayRequest struct {
	*tchttp.BaseRequest

	// 数据所属用户的UIN

	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 产品编码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 页码

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 当前页条数

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListResourceDetailGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceDetailGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadResourceDetailGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务是否已完成（0未完成/1已完成）

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 文件名称，当任务已完成时返回该字段

		Name *string `json:"Name,omitempty" name:"Name"`
		// 下载链接，当任务已完成时返回该字段

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDownloadResourceDetailGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadResourceDetailGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetPropertyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetPropertyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetPropertyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SummaryByResourceDataItem struct {
	// 记录ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 支付者uin

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
	// 子产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品名称

	SubProductCodeName *string `json:"SubProductCodeName,omitempty" name:"SubProductCodeName"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 现金支付总价

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 应付金额

	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`
	// 代金券支付总价

	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 付费模式名称

	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`
	// 付费模式

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 地域ID

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 流水号

	BillId *string `json:"BillId,omitempty" name:"BillId"`
	// 资源所有者uin

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 操作者uin

	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`
	// 交易类型名称

	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`
	// 可用区名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 支付时间

	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`
	// 开始使用时间

	FeeBeginTime *string `json:"FeeBeginTime,omitempty" name:"FeeBeginTime"`
	// 结束使用时间

	FeeEndTime *string `json:"FeeEndTime,omitempty" name:"FeeEndTime"`
	// 组件配置

	ComponentConfig *string `json:"ComponentConfig,omitempty" name:"ComponentConfig"`
	// 订单号

	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
	// 账户名称

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 原价

	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
	// 折扣率

	TotalDiscount *string `json:"TotalDiscount,omitempty" name:"TotalDiscount"`
}

type GetComprehensiveProductStrategyGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本次查询条数

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 跳过条数

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据

		Info []*ProductStrategy `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComprehensiveProductStrategyGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveProductStrategyGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyContractStatusGatewayRequest struct {
	*tchttp.BaseRequest

	// 管理端查询用户UIN，默认缺省查全部

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 合同编号

	ContractId *uint64 `json:"ContractId,omitempty" name:"ContractId"`
	// 0审核中1已邮寄2审核拒绝3作废4取消

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 操作描述拒绝或通过原因

	OperateDesc *string `json:"OperateDesc,omitempty" name:"OperateDesc"`
}

func (r *ModifyContractStatusGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyContractStatusGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddUserStrategyGatewayRequest struct {
	*tchttp.BaseRequest

	// 账号名称

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 账户uin

	ControlUin *string `json:"ControlUin,omitempty" name:"ControlUin"`
	// 页码，从0开始

	Prompt *string `json:"Prompt,omitempty" name:"Prompt"`
	// 固定“bsp”

	Scene *string `json:"Scene,omitempty" name:"Scene"`
	// 固定“admin”

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 限制操作，例如["create","modify"]

	ValidRules []*string `json:"ValidRules,omitempty" name:"ValidRules"`
}

func (r *AddUserStrategyGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddUserStrategyGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDetailByResourceGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 资源花费详情

		Total *ResourceSummaryDetailTotal `json:"Total,omitempty" name:"Total"`
		// 组件花费详情

		BillingItemDetail []*ResourceSummaryDetailComponentDetailItem `json:"BillingItemDetail,omitempty" name:"BillingItemDetail"`
		// 资源历史月份花费趋势

		RealTotalCostTrend *ResourceSummaryDetailRealTotalCostTrend `json:"RealTotalCostTrend,omitempty" name:"RealTotalCostTrend"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDetailByResourceGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDetailByResourceGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContractTypeFlowListRequest struct {
	*tchttp.BaseRequest

	// 合同类型ID号

	ContractTypeId *uint64 `json:"ContractTypeId,omitempty" name:"ContractTypeId"`
	// 合同状态0审核中1已邮寄2审核拒绝3作废4取消

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 管理端查询用户UIN，默认缺省查全部

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeContractTypeFlowListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContractTypeFlowListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPublicAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 条数

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 收款信息列表

		List *QueryPublicAccountList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryPublicAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPublicAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByRegionGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 各地域花费分布详情

		SummaryOverview []*RegionSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByRegionGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByRegionGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBspProductRelationsGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 产品定义列表

		List []*BspProductDefineList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBspProductRelationsGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBspProductRelationsGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductRelationsGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 产品定义列表

		List []*ProductDefineList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductRelationsGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductRelationsGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePropertiesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总行数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 售卖属性列表

		List []*ProductPropertyList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePropertiesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePropertiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddComprehensiveStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddComprehensiveStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddComprehensiveStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicAccountListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 基本账户信息List

		List []*BasicAccount `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBasicAccountListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBasicAccountListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuotaDetailList struct {
	// 账户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品配额标识

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
	// 产品配额值

	QuotaValue *string `json:"QuotaValue,omitempty" name:"QuotaValue"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 子产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 子计费码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品类名

	ProductGroupName *string `json:"ProductGroupName,omitempty" name:"ProductGroupName"`
	// 子产品名

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 计费名

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 子计费名

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
	// 单位名

	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`
}

type DescribeBillSummaryByTagGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 各产品花费分布

		SummaryTotal *BusinessSummaryTotal `json:"SummaryTotal,omitempty" name:"SummaryTotal"`
		// 总花费详情

		SummaryOverview *BusinessSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByTagGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByTagGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateBSPProductDisplayCodeGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateBSPProductDisplayCodeGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateBSPProductDisplayCodeGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductRelationsRequest struct {
	*tchttp.BaseRequest

	// 商品代码或名称

	SearchProduct *string `json:"SearchProduct,omitempty" name:"SearchProduct"`
	// 子商品代码或名称

	SearchSubProduct *string `json:"SearchSubProduct,omitempty" name:"SearchSubProduct"`
	// 计费项代码或名称

	SearchBillingItem *string `json:"SearchBillingItem,omitempty" name:"SearchBillingItem"`
	// 计费细项代码或名称

	SearchSubBillingItem *string `json:"SearchSubBillingItem,omitempty" name:"SearchSubBillingItem"`
	// 四层定义状态，1-生效 0-失效

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 商品代码数组

	SearchProductCodes []*string `json:"SearchProductCodes,omitempty" name:"SearchProductCodes"`
	// 子商品代码数组

	SearchSubProductCodes []*string `json:"SearchSubProductCodes,omitempty" name:"SearchSubProductCodes"`
	// 计费项代码数组

	SearchBillingItemCodes []*string `json:"SearchBillingItemCodes,omitempty" name:"SearchBillingItemCodes"`
	// 计费细项代码数组

	SearchSubBillingItemCodes []*string `json:"SearchSubBillingItemCodes,omitempty" name:"SearchSubBillingItemCodes"`
	// 数据来源类型。local（本系统。默认），third（第三方接入）。不传默认是local

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
}

func (r *DescribeProductRelationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductRelationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTempQuotaListGatewayRequest struct {
	*tchttp.BaseRequest

	// 页数

	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`
	// 一页数目

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 产品一层定义

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品二层定义

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 产品三层定义

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 产品四层定义

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// uin

	QueryUin *string `json:"QueryUin,omitempty" name:"QueryUin"`
}

func (r *DescribeTempQuotaListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTempQuotaListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRenewInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 续费信息

		RenewDetail *RenewDetail `json:"RenewDetail,omitempty" name:"RenewDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRenewInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRenewInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WaterList struct {
	// 订单号

	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 操作时间

	OperateTime *string `json:"OperateTime,omitempty" name:"OperateTime"`
	// 操作类型.取值:open-新购;autopay-自动续费;manualpay-手动续费;upgrade_change-预付费升配;downgrade_change-预付费降配;isolate-隔离;destroy-销毁;modify-后付费变配;return-返还;set_auto_flag-设置续费类型

	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`
	// 操作人

	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`
	// Paras

	Paras []*Paras `json:"Paras,omitempty" name:"Paras"`
	// 0:后付费；1-预付费

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
}

type GetComprehensiveSubProductStrategyRouteGatewayRequest struct {
	*tchttp.BaseRequest

	// 产品代码，如果该值不传则默认查询所有产品

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品代码，如果该值不传则默认查询该产品下所有子产品

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 是否只查询第三方自定义产品策略,如果该值不传或者为false则查询全部产品策略

	IfCustom *string `json:"IfCustom,omitempty" name:"IfCustom"`
	// 状态，如果该值不传则默认查询所有状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 本次查询条数，如果该值不传则默认100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 跳过条数如果该值不传则默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetComprehensiveSubProductStrategyRouteGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveSubProductStrategyRouteGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CouponsInfoList struct {
	// 代金券id

	CodeId *string `json:"CodeId,omitempty" name:"CodeId"`
	// 代金券批次

	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
	// 用户uin

	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DescribeRegionRequest struct {
	*tchttp.BaseRequest

	// 由计平统一分配

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 用户id

	UserNum *string `json:"UserNum,omitempty" name:"UserNum"`
	// 费类型.0:后付费,1:预付费

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 类型.固定传:query_region_by_uin

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryHiddenProductGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 隐藏商品项

		ProductInfo []*HiddenProduct `json:"ProductInfo,omitempty" name:"ProductInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryHiddenProductGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryHiddenProductGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReversalResourceRequest struct {
	*tchttp.BaseRequest

	// 目标用户

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

func (r *ReversalResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReversalResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSignatureRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSignatureRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSignatureRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductDownloadGatewayRequest struct {
	*tchttp.BaseRequest

	// 产品代码

	SearchProductCode *string `json:"SearchProductCode,omitempty" name:"SearchProductCode"`
}

func (r *DescribeProductDownloadGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductDownloadGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductPricesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 价格策略列表

		List []*ProductPriceList `json:"List,omitempty" name:"List"`
		// JSON string；是自定义属性keyCode的数组，包括了中英文名称
		// 例如：{"system":{"chnName":"操作系统","engName":"systemEnglishName"},"brand":{"chnName":"品牌","engName":"brandEnglishName"}}

		CustomPropertyFields *string `json:"CustomPropertyFields,omitempty" name:"CustomPropertyFields"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductPricesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductPricesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetQuotaLeftGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 剩余配额数组

		List []*QuotaLeft `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetQuotaLeftGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetQuotaLeftGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchDiscountGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 折扣列表详情信息

		List []*SearchDiscountList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchDiscountGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchDiscountGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContractTypeFlowData struct {
	// 合同类型名称

	ContractTypeName *string `json:"ContractTypeName,omitempty" name:"ContractTypeName"`
	// 合同状态0审核中1已邮寄2审核拒绝3作废4取消

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 操作描述拒绝或通过原因

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 操作

	Action *string `json:"Action,omitempty" name:"Action"`
	// 模板链接

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
}

type CouponPresentRequest struct {
	*tchttp.BaseRequest

	// 米大师appid

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 米大师发放代金券活动的id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 固定传：supercoupon

	PresentType *string `json:"PresentType,omitempty" name:"PresentType"`
	// 用户ID列表，最多20个。

	UserIdList []*PresentCouponUserInfo `json:"UserIdList,omitempty" name:"UserIdList"`
}

func (r *CouponPresentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CouponPresentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemitRecordData struct {
	// 用户uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 汇款银行名称

	Bank *string `json:"Bank,omitempty" name:"Bank"`
	// 汇款银行账户

	AccountNum *uint64 `json:"AccountNum,omitempty" name:"AccountNum"`
	// 汇款户名

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 汇款时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 联系人手机

	Tel *string `json:"Tel,omitempty" name:"Tel"`
	// 状态 0:处理中 1:成功 2:失败

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 修改时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 汇款人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 汇款金额

	Amount *uint64 `json:"Amount,omitempty" name:"Amount"`
	// 汇款时间

	RemitDate *string `json:"RemitDate,omitempty" name:"RemitDate"`
}

type TempQuotaList struct {
	// uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 产品一层Code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品二层Code

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 产品三层Code

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 产品四层层Code

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 产品一层name

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品二层name

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 产品三层name

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 产品四层name

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
	// 配额key

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
	// 配额值

	QuotaValue *string `json:"QuotaValue,omitempty" name:"QuotaValue"`
}

type DescribeContractTypeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 合同类型列表

		List []*ContractTypeListData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContractTypeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContractTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GeneralObject struct {
	// 字段键值

	Column *string `json:"Column,omitempty" name:"Column"`
	// 字段内容

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ModifyAccountInfoRequest struct {
	*tchttp.BaseRequest

	// 接入后付费账户的唯一标识

	AcctId *string `json:"AcctId,omitempty" name:"AcctId"`
	// 用户ID

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 传后付费固定额度子账户标识: CREDIT_FIXED

	SubAcctid *string `json:"SubAcctid,omitempty" name:"SubAcctid"`
	// 请求方设备，0表示PC，1表示手机

	Device *int64 `json:"Device,omitempty" name:"Device"`
	// UNIX时间戳（从格林威治时间1970年01月01日00时00分00秒起至现在的总秒数）

	Ts *int64 `json:"Ts,omitempty" name:"Ts"`
	// 新的账期，调整后在下个月的出账日生效

	DueDelay *int64 `json:"DueDelay,omitempty" name:"DueDelay"`
	// 枚举值。1：销户； 2：冻结；  3：解冻。 其他情况传空

	Subcmd *int64 `json:"Subcmd,omitempty" name:"Subcmd"`
	// 二级额度（TCE的实际额度/固定额度），注意单位与后付费账户单位保持一致

	SubDplimit *string `json:"SubDplimit,omitempty" name:"SubDplimit"`
}

func (r *ModifyAccountInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccountInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiscountStatusRequest struct {
	*tchttp.BaseRequest

	// 折扣类型 0用户折扣 1官网折扣

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 折扣id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 新的折扣状态值 0未生效 1生效 2失效

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDiscountStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiscountStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubscribeDestroyRequest struct {
	*tchttp.BaseRequest

	// 由计平统一分配

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 用户id

	UserNum *string `json:"UserNum,omitempty" name:"UserNum"`
	// 用户AppId,跟UserId对应

	UserAppId *string `json:"UserAppId,omitempty" name:"UserAppId"`
	// 支付订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合

	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`
	// 操作类型.填destroy.

	Type *string `json:"Type,omitempty" name:"Type"`
	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 资源Id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 地域id

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 付费类型。0:后付费,1:预付费

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 操作人

	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *SubscribeDestroyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubscribeDestroyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDosagesGatewayRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 1 最外层(按产品和地域)汇总， 2: 按天汇总， 3 按小时汇总

	Type *string `json:"Type,omitempty" name:"Type"`
	// 地域id

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 当前页

	PageNo *string `json:"PageNo,omitempty" name:"PageNo"`
	// 当前页大小

	PageSize *string `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeDosagesGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDosagesGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRuleActionsMappingGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRuleActionsMappingGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRuleActionsMappingGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAllProductGatewayRequest struct {
	*tchttp.BaseRequest
}

func (r *GetAllProductGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAllProductGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDiscardProductGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 废弃的产品码List

		List []*DiscardProductCode `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDiscardProductGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDiscardProductGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyComprehensiveStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyComprehensiveStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyComprehensiveStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductRelation struct {
	// 主键ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 产品大类中文名

	ProductGroupName *string `json:"ProductGroupName,omitempty" name:"ProductGroupName"`
	// 产口码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品中文名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 子产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品中文名

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 计费项code

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费项中文名

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 计费细项code

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 计费项英文名

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 计费周期

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 结算周期

	CalcUnit *string `json:"CalcUnit,omitempty" name:"CalcUnit"`
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 单位英文名

	UnitEng *string `json:"UnitEng,omitempty" name:"UnitEng"`
	// 计费细项英文名

	SubBillingItemEngName *string `json:"SubBillingItemEngName,omitempty" name:"SubBillingItemEngName"`
	// 计费项英文名

	BillingItemEngName *string `json:"BillingItemEngName,omitempty" name:"BillingItemEngName"`
	// 子产品英文名

	SubProductEngName *string `json:"SubProductEngName,omitempty" name:"SubProductEngName"`
	// 产品英文名

	ProductEngName *string `json:"ProductEngName,omitempty" name:"ProductEngName"`
	// 产品大类英文名

	ProductGroupEngName *string `json:"ProductGroupEngName,omitempty" name:"ProductGroupEngName"`
	// 1来自计费 0来自计量

	IfBilling *uint64 `json:"IfBilling,omitempty" name:"IfBilling"`
}

type QueryBSPBillingItemGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品三层定义列表

		List []*BspBillingItemList `json:"List,omitempty" name:"List"`
		// 产品三层定义总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBSPBillingItemGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBSPBillingItemGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户状态，0为未开户或销户，1为正常，其他异常
		// （可根据账户状态来判断是否开户）

		Status *string `json:"Status,omitempty" name:"Status"`
		// 后付费余额（包含了cash）

		Balance *string `json:"Balance,omitempty" name:"Balance"`
		// 累计金额，后付费账户表示累积还款额（当前接口无需关心）

		AllIn *string `json:"AllIn,omitempty" name:"AllIn"`
		// 历史累计支出金额，后付费账户表示累积消耗额（当前接口无需关心）

		AllOut *string `json:"AllOut,omitempty" name:"AllOut"`
		// 当天累计消耗

		DayOut *string `json:"DayOut,omitempty" name:"DayOut"`
		// 过期时间，非过期类返回0，否则返回时间UNIX时间戳（当前接口无需关心）

		Exptime *string `json:"Exptime,omitempty" name:"Exptime"`
		// 过期订单信息（当前接口无需关心）

		ExpireinfoExtend *string `json:"ExpireinfoExtend,omitempty" name:"ExpireinfoExtend"`
		// 后付费额度上限

		Dplimit *string `json:"Dplimit,omitempty" name:"Dplimit"`
		// 后付费剩余可用额度（包含了cash）

		Dpbalance *string `json:"Dpbalance,omitempty" name:"Dpbalance"`
		// 当前额度使用量（包含转出和预授权等）

		Debt *string `json:"Debt,omitempty" name:"Debt"`
		// 后付费账户，存储现金/溢出金额，下次出账时抵消欠款

		Cash *string `json:"Cash,omitempty" name:"Cash"`
		// 出账日，1-28的数字，每个月第几天出账单

		BillingDate *string `json:"BillingDate,omitempty" name:"BillingDate"`
		// 还款日，1-28的数字，每个月第几天还款

		DueDate *string `json:"DueDate,omitempty" name:"DueDate"`
		// 商户类型，1：代理，2：子客，3：直客

		MerchantType *string `json:"MerchantType,omitempty" name:"MerchantType"`
		// 关联商户信息，当商户类型为代理或直客时，该字段为空；当为子客时，标识代理账户信息

		RelatedMerchant *string `json:"RelatedMerchant,omitempty" name:"RelatedMerchant"`
		// 账期类型，1：按月，2：按天

		DueDelayType *string `json:"DueDelayType,omitempty" name:"DueDelayType"`
		// 账期，还款允许延期时间，当账期类型按月时，表示月份数，按天时，表示天数

		DueDelay *string `json:"DueDelay,omitempty" name:"DueDelay"`
		// 下一个出账日将生效的新账期（无新账期则返回空）

		NextDueDelay *string `json:"NextDueDelay,omitempty" name:"NextDueDelay"`
		// 还款时账单未还清是否立即恢复额度
		// 0 (默认)不立即恢复 1 立即恢复

		RecoverImmediately *string `json:"RecoverImmediately,omitempty" name:"RecoverImmediately"`
		// 是否校验账单还款顺序
		// 0 (默认)按顺序  1 不按顺序

		RefundSequence *string `json:"RefundSequence,omitempty" name:"RefundSequence"`
		// 提前还款是否立即恢复额度
		// 0 (默认)不立即恢复，出账时抵消账单才恢复；  1 立即恢复

		CashRecoverImmediately *string `json:"CashRecoverImmediately,omitempty" name:"CashRecoverImmediately"`
		// 出账周期类型，1：按月。其他暂不支持

		BillingCycleType *string `json:"BillingCycleType,omitempty" name:"BillingCycleType"`
		// 出账周期，若按月则表示月份数，按天则表示天数

		BillingCycle *string `json:"BillingCycle,omitempty" name:"BillingCycle"`
		// 二级额度（TCE的实际额度/固定额度），注意单位与后付费账户单位保持一致

		SubDplimit *string `json:"SubDplimit,omitempty" name:"SubDplimit"`
		// 已用额度

		UsedLimit *uint64 `json:"UsedLimit,omitempty" name:"UsedLimit"`
		// 现金余额

		CashBalance *uint64 `json:"CashBalance,omitempty" name:"CashBalance"`
		// 固定额度

		Fixedlimit *string `json:"Fixedlimit,omitempty" name:"Fixedlimit"`
		// 可用额度

		UsableLimit *uint64 `json:"UsableLimit,omitempty" name:"UsableLimit"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductQuotaGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductQuotaGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductQuotaGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCouponResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计费MidasAppId

		MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
		// 发券开始时间

		BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
		// 发券结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 券类型，card: 卡券，balance: 余额券

		CouponType *string `json:"CouponType,omitempty" name:"CouponType"`
		// 活动号

		ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`
		// 状态  1为正常，<=0为下线状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 活动名称

		ActivityName *string `json:"ActivityName,omitempty" name:"ActivityName"`
		// 有效期类型

		ExpiryTimeType *string `json:"ExpiryTimeType,omitempty" name:"ExpiryTimeType"`
		// 有效时间值

		TimeValue *int64 `json:"TimeValue,omitempty" name:"TimeValue"`
		// 有效时间单位，D天，H小时

		TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
		// 固定有效时间开始

		ExpiryTimeBegin *string `json:"ExpiryTimeBegin,omitempty" name:"ExpiryTimeBegin"`
		// 固定有效时间结束

		ExpiryTimeEnd *string `json:"ExpiryTimeEnd,omitempty" name:"ExpiryTimeEnd"`
		// 创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 最后更新时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 券门槛（分）

		CouponThreshold *int64 `json:"CouponThreshold,omitempty" name:"CouponThreshold"`
		// 券面值（分）

		CouponValue *int64 `json:"CouponValue,omitempty" name:"CouponValue"`
		// 付费模式 0-预付费，1-后付费，0,1-全部

		PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
		// 四层产品定义

		ProductInfo *string `json:"ProductInfo,omitempty" name:"ProductInfo"`
		// 创建用户

		CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
		// 用户领取限制

		UserLimit *UserLimit `json:"UserLimit,omitempty" name:"UserLimit"`
		// 发放总限制

		ActivityLimit *ActivityLimit `json:"ActivityLimit,omitempty" name:"ActivityLimit"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCouponResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCouponResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByRegionRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 每次获取数据量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 是否返回最大页码，建议首次或变更PageSize后才需要，用于提升接口性

	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
}

func (r *DescribeBillFeeByRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeByRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDetailByResourceGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 付费模式，只能是prePay或者postPay

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 预付费才需要传，为返回的BillId

	BillId *string `json:"BillId,omitempty" name:"BillId"`
	// 预付费才需要传，为返回的Id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeBillDetailByResourceGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDetailByResourceGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByResourceGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 资源汇总详情

		Data []*SummaryByResourceDataItem `json:"Data,omitempty" name:"Data"`
		// 资源汇总花费详情

		Total *SummaryByResourceTotal `json:"Total,omitempty" name:"Total"`
		// 记录数量，NeedRecordNum为0时该值为null

		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`
		// 过滤条件，NeedConditionValue为0时该值为null

		ConditionValue *SummaryByResourceConditionValue `json:"ConditionValue,omitempty" name:"ConditionValue"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByResourceGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByResourceGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LifecycleFlowListData struct {
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 操作

	OperateAction *string `json:"OperateAction,omitempty" name:"OperateAction"`
	// 请求包

	Request *string `json:"Request,omitempty" name:"Request"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 操作者UIN

	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`
	// 商品码名称

	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
	// 序列号

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 操作中文名称

	OperateActionName *string `json:"OperateActionName,omitempty" name:"OperateActionName"`
}

type WaterListData struct {
	// 流水类型

	WaterType *string `json:"WaterType,omitempty" name:"WaterType"`
	// 存取标志，1：支出；2：收入

	IoFlag *string `json:"IoFlag,omitempty" name:"IoFlag"`
	// 给外部提供的接口所对应的交易类型：  2-支付  10-转账  19-还款 20-销帐 21-预授权 22-预授权取消（恢复） 23-预授权确认（消耗）24-退款 25-产生账单 26-修改账单 27-临时账户新增扩展额度 28-撤销还款 29-预授权确认回退 30-预授权改单 31-预授权回退

	ExternTranType *string `json:"ExternTranType,omitempty" name:"ExternTranType"`
	// 对应到内部server接口类型： 2-支付  10-开户（开户接口触发） 14-转账扣款   16-转账充值  17-转账取消   42-注销   52-后付费上调额度   53-后付费下调额度   54-后付费更改账户信息（除调额外） 55-还款 56-销账 57-提前还款   58-抵消账单   59-预授权   60-预授权取消（恢复） 61-预授权工具取消（恢复） 62-预授权确认（消耗） 63-预授权订单方面 64-预授权取消订单方面 65-退款（给账户充） 66-应恢复额度流水 67-补扣/补量68-增加账单金额（出账工具）69-减少账单金额（出账工具）70-还原cash金额（出账工具）71-修改账单导致余额增加（出账工具）72-修改账单导致余额减少（出账工具）73-新增临时账户扩展额度   74-撤销还款   75-撤销销帐   76-撤销提前还款   82-预授权改单上调，账户减少   83-预授权改单下调，账户增加   84-预授权改单上调，预授权单增加   85-预授权改单下调，预授权单减少   86-预授权退款   87-预授权补扣

	BaseTranType *string `json:"BaseTranType,omitempty" name:"BaseTranType"`
	// 订单号

	BillNo *string `json:"BillNo,omitempty" name:"BillNo"`
	// 交易金额

	TranAmt *string `json:"TranAmt,omitempty" name:"TranAmt"`
	// 余额，后付费账户balance根据账户属性决定等于dpbalance还是dpbalance+cash

	Balance *string `json:"Balance,omitempty" name:"Balance"`
	// 流水产生时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 渠道来源

	Source *string `json:"Source,omitempty" name:"Source"`
	// 客户端设备，0：PC；1：手机

	Device *string `json:"Device,omitempty" name:"Device"`
	// 用户IP

	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`
	// 交易描述

	TranInfo *string `json:"TranInfo,omitempty" name:"TranInfo"`
	// 交易备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 保留字段1

	Reserve1 *string `json:"Reserve1,omitempty" name:"Reserve1"`
	// 保留字段2

	Reserve2 *string `json:"Reserve2,omitempty" name:"Reserve2"`
	// 保留字段3 （记录流水插入DB时间）

	Reserve3 *string `json:"Reserve3,omitempty" name:"Reserve3"`
	// 保留字段4（1. 更改账户的标记；2.还款流水记录销账的账单号；3.出账时，完全抵消账单金额标记AllOffSet）

	Reserve4 *string `json:"Reserve4,omitempty" name:"Reserve4"`
	// 保留字段5（交易时间微秒us）

	Reserve5 *string `json:"Reserve5,omitempty" name:"Reserve5"`
	// 保留字段6 （销帐记录：TotalRefund,1000;CurDebt,1110;DebtBillDate,201706; PayChannel,bank;）

	Reserve6 *string `json:"Reserve6,omitempty" name:"Reserve6"`
	// 保留字段7（销帐记录账单恢复金额：  Type,还款类型;Recover,恢复额度;ChgTimeUs,HOLD中的交易微妙; TranSeq,保存在账户中的交易序列号;）

	Reserve7 *string `json:"Reserve7,omitempty" name:"Reserve7"`
	// 保留字段8

	Reserve8 *string `json:"Reserve8,omitempty" name:"Reserve8"`
	// 由外部输入的扩展性保留字段1

	Extreserve1 *string `json:"Extreserve1,omitempty" name:"Extreserve1"`
	// 由外部输入的扩展性保留字段2

	Extreserve2 *string `json:"Extreserve2,omitempty" name:"Extreserve2"`
	// 关联id，如销账的账单号

	RelationId *string `json:"RelationId,omitempty" name:"RelationId"`
	// 后付费额度上限

	Dplimit *string `json:"Dplimit,omitempty" name:"Dplimit"`
	// 后付费剩余可用额度（根据账户属性决定是否包含cash）

	Dpbalance *string `json:"Dpbalance,omitempty" name:"Dpbalance"`
	// 后付费账户，固定账户存储现金余额

	Cash *string `json:"Cash,omitempty" name:"Cash"`
	// 商户类型，1：代理，2：子客，3：直客

	MerchantType *string `json:"MerchantType,omitempty" name:"MerchantType"`
	// 关联商户信息，当商户类型为代理或直客时，该字段为空；当为子客时，标识代理账户信息

	RelatedMerchant *string `json:"RelatedMerchant,omitempty" name:"RelatedMerchant"`
}

type DebtBillListData struct {
	// 账单号

	DebtBillNo *string `json:"DebtBillNo,omitempty" name:"DebtBillNo"`
	// 账单状态，数字 1-已还清 2-未还清

	DebtBillStatus *string `json:"DebtBillStatus,omitempty" name:"DebtBillStatus"`
	// 账单月份yyyymm

	BillDate *string `json:"BillDate,omitempty" name:"BillDate"`
	// 账单结束日期yyyymmdd（灵活出账周期模式）

	BillDateEnd *string `json:"BillDateEnd,omitempty" name:"BillDateEnd"`
	// 最后还款时间，UNIX时间戳

	DueDate *string `json:"DueDate,omitempty" name:"DueDate"`
	// 逾期利息

	DebtInterest *string `json:"DebtInterest,omitempty" name:"DebtInterest"`
	// 欠款总额

	TotalDebt *string `json:"TotalDebt,omitempty" name:"TotalDebt"`
	// 当前仍欠金额

	CurrentDebt *string `json:"CurrentDebt,omitempty" name:"CurrentDebt"`
	// 已还金额

	TotalRefund *string `json:"TotalRefund,omitempty" name:"TotalRefund"`
	// 还款记录，格式（time为UNIX时戳）： time1,refundbillno1,amt1;time2,refundbillno2,amt2;….

	RefundRecord *string `json:"RefundRecord,omitempty" name:"RefundRecord"`
	// 账期类型，1：按月，2：按天

	DuedelayType *string `json:"DuedelayType,omitempty" name:"DuedelayType"`
	// 账期，还款允许延期时间，当账期类型按月时，表示月份数，按天时，表示天数

	DueDelay *string `json:"DueDelay,omitempty" name:"DueDelay"`
	// 账单生成时间，UNIX时间戳

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 账单最后更新时间，UNIX时间戳

	LastChgTimes *string `json:"LastChgTimes,omitempty" name:"LastChgTimes"`
	// 账单备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// Kv拼接的扩展字段，采用逗号和分号分割，示例如下：
	// Campid,xxxxx;Subclientid,xxxx;

	Reserve2 *string `json:"Reserve2,omitempty" name:"Reserve2"`
}

type ExportChangeQuotaRequest struct {
	*tchttp.BaseRequest

	// 变更时间-开始

	CreateStartTime *string `json:"CreateStartTime,omitempty" name:"CreateStartTime"`
	// 变更时间-结束

	CreateEndTime *string `json:"CreateEndTime,omitempty" name:"CreateEndTime"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 产品细项码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 子产品细项码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 变更者UIN

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 租户UIN

	SearchUin *string `json:"SearchUin,omitempty" name:"SearchUin"`
	// 配额类型

	QuotaType *string `json:"QuotaType,omitempty" name:"QuotaType"`
}

func (r *ExportChangeQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportChangeQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCouponRequest struct {
	*tchttp.BaseRequest

	// 计费MidasAppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 发券活动ID/groupid

	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`
}

func (r *QueryCouponRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCouponRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateUserStrategyGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateUserStrategyGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateUserStrategyGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpireDiscountGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExpireDiscountGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExpireDiscountGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportDealListOpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportDealListOpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDealListOpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDiscardProductRequest struct {
	*tchttp.BaseRequest

	// 废弃层级：1-产品码

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *GetDiscardProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDiscardProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInvoiceListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总和

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 当前页数

		Cur *int64 `json:"Cur,omitempty" name:"Cur"`
		// 总页数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 开票记录列表

		List []*DescribeInvoiceList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInvoiceListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInvoiceListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotasRequest struct {
	*tchttp.BaseRequest

	// 传空字符串返回system默认配额，
	// 不传就返回全量用户配额

	QuotaUin *string `json:"QuotaUin,omitempty" name:"QuotaUin"`
	// 产品code。不传或者空就是查询全部

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
}

func (r *DescribeQuotasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGeneralDataGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 查询条件

	Conditions []*GeneralObject `json:"Conditions,omitempty" name:"Conditions"`
	// 返回的字段

	Filter []*string `json:"Filter,omitempty" name:"Filter"`
	// 排序字段

	SortField *string `json:"SortField,omitempty" name:"SortField"`
	// 排序方式 “asc”或“desc”

	SortType *string `json:"SortType,omitempty" name:"SortType"`
	// 分页，起始值

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页，数目

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGeneralDataGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralDataGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceListGatewayRequest struct {
	*tchttp.BaseRequest

	// 管理端查询用户，只支持单个查询

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 支付模式

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
}

func (r *DescribeResourceListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveSubProdctStrategyGatewayRequest struct {
	*tchttp.BaseRequest

	// 产品代码，如果该值不传则默认查询所有产品

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品代码，如果该值不传则默认查询该产品下所有子产品

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 是否只查询第三方自定义产品策略,如果该值不传或者为false则查询全部产品策略

	IfCustom *string `json:"IfCustom,omitempty" name:"IfCustom"`
	// 状态，如果该值不传则默认查询所有状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 本次查询条数，如果该值不传则默认100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 跳过条数如果该值不传则默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetComprehensiveSubProdctStrategyGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveSubProdctStrategyGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateResourceGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateResourceGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolateResourceGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceDetailDownloadGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListResourceDetailDownloadGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceDetailDownloadGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetPropertyRequest struct {
	*tchttp.BaseRequest

	// key中文名

	KeyNameCn *string `json:"KeyNameCn,omitempty" name:"KeyNameCn"`
	// key英文名

	KeyNameEn *string `json:"KeyNameEn,omitempty" name:"KeyNameEn"`
	// key码

	KeyCode *string `json:"KeyCode,omitempty" name:"KeyCode"`
	// key的字段列表

	Fields []*ProductPropertyFields `json:"Fields,omitempty" name:"Fields"`
	// JSON String 属性值列表
	// jsond ecode后是数组，数组元素的key不定
	// 例如：[{"ChnName":"深圳","EngName":"Shenzhen","AreaCode":"szx","ID":"37","Status":"1","Desc":"test"}]

	Values *string `json:"Values,omitempty" name:"Values"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 类型： 0（或者不传或者传空）是普通属性；1是自定义属性
	// 自定义属性写入后，一定是可编辑的。Modifiable=1

	PropertyType *int64 `json:"PropertyType,omitempty" name:"PropertyType"`
	// 自定义属性关联的四层定义。
	// 字符串数组，例如： ["p_cvm##v_cvm_mem#sv_cvm_mem_s1","p_cvm##v_cvm_mem#sv_cvm_mem_s2"]

	CustomPropertyRelateProducts []*string `json:"CustomPropertyRelateProducts,omitempty" name:"CustomPropertyRelateProducts"`
	// 自定义属性关联的模块。
	// 字符串数组，例如： ["price"]

	CustomPropertyRelateModels []*string `json:"CustomPropertyRelateModels,omitempty" name:"CustomPropertyRelateModels"`
}

func (r *SetPropertyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetPropertyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsolidatedBillDownloadUrlByUinsGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务是否准备就绪

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsolidatedBillDownloadUrlByUinsGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsolidatedBillDownloadUrlByUinsGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadWithAuthGatewayRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
}

func (r *DescribeDownloadWithAuthGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadWithAuthGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InterfaceNames struct {
	// 查询资源

	QueryAllResource *string `json:"QueryAllResource,omitempty" name:"QueryAllResource"`
	// 对账接口

	QueryResource *string `json:"QueryResource,omitempty" name:"QueryResource"`
	// 隔离资源

	IsolateResource *string `json:"IsolateResource,omitempty" name:"IsolateResource"`
	// 销毁资源

	DestroyResource *string `json:"DestroyResource,omitempty" name:"DestroyResource"`
	// 检查发货参数

	CheckCreate *string `json:"CheckCreate,omitempty" name:"CheckCreate"`
	// 发货

	CreateResource *string `json:"CreateResource,omitempty" name:"CreateResource"`
	// 设置自动续费标记

	SetRenewFlag *string `json:"SetRenewFlag,omitempty" name:"SetRenewFlag"`
	// 检查续费参数（预付费）

	CheckRenew *string `json:"CheckRenew,omitempty" name:"CheckRenew"`
	// 续费（预付费）

	RenewResource *string `json:"RenewResource,omitempty" name:"RenewResource"`
	// 检查变配参数

	CheckModify *string `json:"CheckModify,omitempty" name:"CheckModify"`
	// 变配

	ModifyResource *string `json:"ModifyResource,omitempty" name:"ModifyResource"`
	// 检查资源状态（针对异步发货产品）

	QueryFlow *string `json:"QueryFlow,omitempty" name:"QueryFlow"`
	// 隔离资源

	PostPayIsolateResource *string `json:"PostPayIsolateResource,omitempty" name:"PostPayIsolateResource"`
	// 后付费释放资源

	PostPayReleaseResource *string `json:"PostPayReleaseResource,omitempty" name:"PostPayReleaseResource"`
	// 冲正解封

	PostPayReversalResource *string `json:"PostPayReversalResource,omitempty" name:"PostPayReversalResource"`
}

type DescribeBillDetailByResourceRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 付费模式，只能是prePay或者postPay

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 预付费才需要传，为返回的BillId

	BillId *string `json:"BillId,omitempty" name:"BillId"`
	// 预付费才需要传，为返回的Id

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeBillDetailByResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDetailByResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetProductInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetProductInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetProductInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PresentCouponUserInfo struct {
	// 用户id

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 发代金券订单号

	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`
}

type ModifyReconciliationRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 调账原因：计费错误-billing_error 产品故障-business_error 内部核销-internal_write_off 其它原因-other

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 调账类型 add-补偿/minus-扣费

	AdjustType *string `json:"AdjustType,omitempty" name:"AdjustType"`
	// 调账月份： "2019-05"

	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`
	// 调账金额：元

	Amount *string `json:"Amount,omitempty" name:"Amount"`
	// 调账订单号

	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`
}

func (r *ModifyReconciliationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReconciliationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActivityLimit struct {
	// 天限

	DayLimit *int64 `json:"DayLimit,omitempty" name:"DayLimit"`
	// 周限

	WeekLimit *int64 `json:"WeekLimit,omitempty" name:"WeekLimit"`
	// 月限

	MonthLimit *int64 `json:"MonthLimit,omitempty" name:"MonthLimit"`
	// 总限

	AllLimit *int64 `json:"AllLimit,omitempty" name:"AllLimit"`
}

type DescribeSignatureGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求状态标识

		SuccessActionStatus *int64 `json:"SuccessActionStatus,omitempty" name:"SuccessActionStatus"`
		// 凭据

		XAmzCredential *string `json:"XAmzCredential,omitempty" name:"XAmzCredential"`
		// 算法

		XAmzAlgorithm *string `json:"XAmzAlgorithm,omitempty" name:"XAmzAlgorithm"`
		// 日期

		XAmzDate *string `json:"XAmzDate,omitempty" name:"XAmzDate"`
		// 密文

		Policy *string `json:"Policy,omitempty" name:"Policy"`
		// 签名

		XAmzSignature *string `json:"XAmzSignature,omitempty" name:"XAmzSignature"`
		// 地址

		Url *string `json:"Url,omitempty" name:"Url"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSignatureGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSignatureGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetAutoFlagRequest struct {
	*tchttp.BaseRequest

	// 由计平统一分配

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 用户id

	UserNum *string `json:"UserNum,omitempty" name:"UserNum"`
	// 操作人

	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`
	// 资源id列表，json格式见下.最多10个

	ResourceList []*ResourceList `json:"ResourceList,omitempty" name:"ResourceList"`
	// 自动续费模式：
	// 1：自动续费;
	// 0:手动续费,到期后一定间隔时间后回收;
	// 2:到期关闭

	AutopayType *uint64 `json:"AutopayType,omitempty" name:"AutopayType"`
	// 防重订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合。

	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`
	// 付费模式,1:预付费。 目前只支持预付费

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 类型.固定传: set_auto_flag

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *SetAutoFlagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetAutoFlagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetProductStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetProductStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetProductStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadWithOperatorGatewayRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeDownloadWithOperatorGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadWithOperatorGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceBillDetailGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 每次获取数据量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 是否需要记录总数，0不需要，1需要

	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
	// 是否需要过滤条件，0需要，1不需要

	NeedConditionValue *uint64 `json:"NeedConditionValue,omitempty" name:"NeedConditionValue"`
	// 只支持BusinessCodes，ProductCodes，ComponentCodes，PayMode, ProjectIds, RegionIds, PayModes, ActionTypes, BillIds, HideFreeCost, ResourceKeyword

	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
}

func (r *DescribeResourceBillDetailGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceBillDetailGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResourceBillResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 计量推量明细数组

		ResourceBill []*BspResourceBill `json:"ResourceBill,omitempty" name:"ResourceBill"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetResourceBillResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResourceBillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsolidatedBillDownloadUrlByUinsGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUins []*string `json:"PayerUins,omitempty" name:"PayerUins"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 内部操作者

	InnerOperator *string `json:"InnerOperator,omitempty" name:"InnerOperator"`
}

func (r *DescribeConsolidatedBillDownloadUrlByUinsGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsolidatedBillDownloadUrlByUinsGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonthCostDetailItem struct {
	// 月份中文名，如2018年12月

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实际花费

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 应收花费

	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`
	// 月份，如2018-12

	Month *string `json:"Month,omitempty" name:"Month"`
}

type ExportDiscountGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportDiscountGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDiscountGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseResourceGatewayRequest struct {
	*tchttp.BaseRequest

	// 目标用户

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

func (r *ReleaseResourceGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseResourceGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContractListData struct {
	// 用户uin

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 状态0审核中1已邮寄2审核拒绝3作废4取消

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 合同编号

	ContractId *uint64 `json:"ContractId,omitempty" name:"ContractId"`
	// 申请时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 合同类型ID号

	ContractTypeId *uint64 `json:"ContractTypeId,omitempty" name:"ContractTypeId"`
	// 合同类型名称

	ContractTypeName *string `json:"ContractTypeName,omitempty" name:"ContractTypeName"`
	// 甲方名称

	PartAName *string `json:"PartAName,omitempty" name:"PartAName"`
	// 甲方地址

	PartAAddress *string `json:"PartAAddress,omitempty" name:"PartAAddress"`
	// 甲方联系人

	ConnectUser *string `json:"ConnectUser,omitempty" name:"ConnectUser"`
	// 甲方电话

	Phone *string `json:"Phone,omitempty" name:"Phone"`
	// 甲方邮箱

	Email *string `json:"Email,omitempty" name:"Email"`
	// 收件人

	AdresseeName *string `json:"AdresseeName,omitempty" name:"AdresseeName"`
	// 收件人电话

	AdresseePhone *string `json:"AdresseePhone,omitempty" name:"AdresseePhone"`
	// 收件人地址

	AdresseeDress *string `json:"AdresseeDress,omitempty" name:"AdresseeDress"`
	// 操作描述拒绝或通过原因

	OperateDesc *string `json:"OperateDesc,omitempty" name:"OperateDesc"`
	// 备注描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 下载链接

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
}

type DescribeTenantSubUinQuotasGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总行数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 列表数据

		List *SubUinQuotaList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTenantSubUinQuotasGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTenantSubUinQuotasGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveSubProdctStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 跳过条数

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 本次返回条数

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据

		Info []*ProductStrategy `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComprehensiveSubProdctStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveSubProdctStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionDetail struct {
	// 数量

	TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`
	// 地域列表

	RegionList []*string `json:"RegionList,omitempty" name:"RegionList"`
}

type ExportDiscountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportDiscountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDiscountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPublicAccountGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 条数

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 收款信息列表

		List *QueryPublicAccountList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryPublicAccountGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPublicAccountGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadCommonSummaryRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillDownloadCommonSummaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadCommonSummaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadRecordListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数量，NeedRecordNum为0时返回null

		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`
		// 下载记录列表

		Records *GetDownloadListRecordItem `json:"Records,omitempty" name:"Records"`
		// 过滤条件，NeedConditionValue为0是返回null

		ConditionValue *DetailConditionValue `json:"ConditionValue,omitempty" name:"ConditionValue"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDownloadRecordListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadRecordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByProductRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 每次获取数据量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 是否返回最大页码，建议首次或变更PageSize后才需要，用于提升接口性

	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
}

func (r *DescribeBillFeeByProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeByProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BillingItem struct {
	// 计费项代码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费项中文名

	ChnName *string `json:"ChnName,omitempty" name:"ChnName"`
	// 计费项英文名

	EngName *string `json:"EngName,omitempty" name:"EngName"`
	// 计费细项状

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeBillDownloadByResourceRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 需要导出相关字段

	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeBillDownloadByResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadByResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadDosagesGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadDosagesGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadDosagesGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WaterDetail struct {
	// 数量

	TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`
	// 流水

	WaterList []*WaterList `json:"WaterList,omitempty" name:"WaterList"`
}

type DescribeResourceBillDownloadRequest struct {
	*tchttp.BaseRequest

	// UIN

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 资源Id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 标签组合字符串

	Tags *string `json:"Tags,omitempty" name:"Tags"`
}

func (r *DescribeResourceBillDownloadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceBillDownloadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductSubBillingInfoList struct {
	// 产品code（第一层）

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品计费项code（第三层）

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 产品计费细项code（第四层）

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 中文名

	ChnName *string `json:"ChnName,omitempty" name:"ChnName"`
	// 英文名

	EngName *string `json:"EngName,omitempty" name:"EngName"`
	// 单位

	GoodNumUnit *string `json:"GoodNumUnit,omitempty" name:"GoodNumUnit"`
	// 单位英文名

	GoodNumUnitEn *string `json:"GoodNumUnitEn,omitempty" name:"GoodNumUnitEn"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ListResourceDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源详情信息

		List *BspResourceBill `json:"List,omitempty" name:"List"`
		// 数据条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListResourceDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTenantSubUinQuotasRequest struct {
	*tchttp.BaseRequest

	// 计费MidasAppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 第一层产品定义的code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

func (r *DescribeTenantSubUinQuotasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTenantSubUinQuotasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadQuotasListFileRequest struct {
	*tchttp.BaseRequest

	// 用户uin

	QuotaUin *string `json:"QuotaUin,omitempty" name:"QuotaUin"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 添加任务者，用于内部TCB邮件通知

	InnerOperator *string `json:"InnerOperator,omitempty" name:"InnerOperator"`
}

func (r *DescribeDownloadQuotasListFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadQuotasListFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGoodsPriceGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGoodsPriceGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGoodsPriceGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceBillDetailGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否展示数据，0不展示，1展示

		DisplayData *uint64 `json:"DisplayData,omitempty" name:"DisplayData"`
		// 记录数量，NeedRecordNum为0是返回null

		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`
		// 过滤条件，NeedConditionValue为0是返回null

		ConditionValue *DetailConditionValue `json:"ConditionValue,omitempty" name:"ConditionValue"`
		// 总花费详情

		Total *DetailTotal `json:"Total,omitempty" name:"Total"`
		// 组件花费详情

		Detail []*DetailDetailItem `json:"Detail,omitempty" name:"Detail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceBillDetailGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceBillDetailGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBreakdownModeGatewayRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryBreakdownModeGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBreakdownModeGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteQuotaRequest struct {
	*tchttp.BaseRequest

	// 商品码，商品码和QuotaKey不能同时为空

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 配额key

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
	// uin

	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`
}

func (r *DeleteQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByResourceGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 每次获取数据量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 是否需要记录总数，0不需要，1需要

	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
	// 是否需要过滤条件，0需要，1不需要

	NeedConditionValue *uint64 `json:"NeedConditionValue,omitempty" name:"NeedConditionValue"`
	// 只支持BusinessCodes，ProductCodes，PayModes, ProjectIds, RegionIds, PayModes, ActionTypes, HideFreeCost, ResourceKeyword, OrderByCost

	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
}

func (r *DescribeBillSummaryByResourceGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByResourceGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContractListGatewayRequest struct {
	*tchttp.BaseRequest

	// 管理端查询用户UIN，默认缺省查全部

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 合同编号

	ContractId *uint64 `json:"ContractId,omitempty" name:"ContractId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 合同类型ID号

	ContractTypeId *uint64 `json:"ContractTypeId,omitempty" name:"ContractTypeId"`
	// 创建时间排序正序asc 倒序desc

	CreateTimeOrder *string `json:"CreateTimeOrder,omitempty" name:"CreateTimeOrder"`
	// 合同状态0审核中1已邮寄2审核拒绝3作废4取消

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeContractListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContractListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancleContractTypeRequest struct {
	*tchttp.BaseRequest

	// 合同类型唯一ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 状态1正常2作废

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *CancleContractTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancleContractTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyContractStatusRequest struct {
	*tchttp.BaseRequest

	// 管理端查询用户UIN，默认缺省查全部

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 合同编号

	ContractId *uint64 `json:"ContractId,omitempty" name:"ContractId"`
	// 0审核中1已邮寄2审核拒绝3作废4取消

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 操作描述拒绝或通过原因

	OperateDesc *string `json:"OperateDesc,omitempty" name:"OperateDesc"`
}

func (r *ModifyContractStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyContractStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConditionSubProduct struct {
	// 子产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品名称

	SubProductCodeName *string `json:"SubProductCodeName,omitempty" name:"SubProductCodeName"`
}

type ExpireDiscountGatewayRequest struct {
	*tchttp.BaseRequest

	// 折扣类型 0用户折扣 1官网折扣

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 折扣id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *ExpireDiscountGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExpireDiscountGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTagsByMonthRequest struct {
	*tchttp.BaseRequest

	// uin

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 月份

	Month *string `json:"Month,omitempty" name:"Month"`
	// 查询全部，传1

	State *uint64 `json:"State,omitempty" name:"State"`
}

func (r *GetTagsByMonthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTagsByMonthRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadByResourceGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 需要导出相关字段

	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeBillDownloadByResourceGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadByResourceGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaTypeListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配额类别列表

		QuoteTypeList []*QuoteTypeList `json:"QuoteTypeList,omitempty" name:"QuoteTypeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQuotaTypeListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotaTypeListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByTagRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
}

func (r *DescribeBillSummaryByTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户额度详细

		List []*QuotaDetailList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BspSubBillingItemList struct {
	// 子产品代码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 子产品名称

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
	// 子产品名称

	DisplayCode *string `json:"DisplayCode,omitempty" name:"DisplayCode"`
	// 产品展示名称

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 产品展示名称

	TimeUint *uint64 `json:"TimeUint,omitempty" name:"TimeUint"`
	// 产品展示名称

	CalcUnit *uint64 `json:"CalcUnit,omitempty" name:"CalcUnit"`
	// 计量单位

	Unit *uint64 `json:"Unit,omitempty" name:"Unit"`
}

type DescribeBillFeeByTagRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 每次获取数据量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 是否返回最大页码，建议首次或变更PageSize后才需要，用于提升接口性

	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
}

func (r *DescribeBillFeeByTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeByTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemitRecordGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 列表

		List []*RemitRecordData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemitRecordGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemitRecordGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaTypeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配额类别列表

		QuoteTypeList []*QuoteTypeList `json:"QuoteTypeList,omitempty" name:"QuoteTypeList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQuotaTypeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotaTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductTreeGatewayRequest struct {
	*tchttp.BaseRequest

	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 页索引

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 是否生效.0未生效 1生效

	Status *string `json:"Status,omitempty" name:"Status"`
	// 数据是否过滤隐藏数据,true表示需要

	ExceptHidden *bool `json:"ExceptHidden,omitempty" name:"ExceptHidden"`
	// 过滤隐藏计量或计费数据标志1计费0计量,不传为全部隐藏

	IfBilling *int64 `json:"IfBilling,omitempty" name:"IfBilling"`
}

func (r *GetProductTreeGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductTreeGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiscountStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 折扣id

		// 条数

		// 是否成功

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDiscountStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiscountStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBreakdownModeGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1 正常 2恢复中 3故障中

		Status *string `json:"Status,omitempty" name:"Status"`
		// 白名单列表，只有故障中才会返回

		WhiteList []*string `json:"WhiteList,omitempty" name:"WhiteList"`
		// 临时配额开关，1 关闭 2开启

		TempQuto *string `json:"TempQuto,omitempty" name:"TempQuto"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBreakdownModeGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBreakdownModeGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryRecoverUinGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本页返回条数

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 跳过条数

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 恢复uin列表

		Info []*string `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryRecoverUinGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryRecoverUinGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadListGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 开始月份

	BeginMonth *string `json:"BeginMonth,omitempty" name:"BeginMonth"`
	// 结束月份

	EndMonth *string `json:"EndMonth,omitempty" name:"EndMonth"`
}

func (r *DescribeBillDownloadListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTempQuotaListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// list

		List *TempQuotaList `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTempQuotaListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTempQuotaListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDebtBillResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账单总数

		AllNum *string `json:"AllNum,omitempty" name:"AllNum"`
		// 返回的账单数量

		RetNum *string `json:"RetNum,omitempty" name:"RetNum"`
		// 当前游标

		Offset *string `json:"Offset,omitempty" name:"Offset"`
		// 账单详情列表

		DebtBillSet []*DebtBillListData `json:"DebtBillSet,omitempty" name:"DebtBillSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserDebtBillResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserDebtBillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportDiscountRequest struct {
	*tchttp.BaseRequest

	// 折扣类型 0用户折扣 1官网折扣

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 开始生效时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 失效时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 折扣创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 折扣状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 开始时间大于该值

	PreFBeginTime *string `json:"PreFBeginTime,omitempty" name:"PreFBeginTime"`
	// 开始时间小于该值

	TailFBeginTime *string `json:"TailFBeginTime,omitempty" name:"TailFBeginTime"`
	// 结束时间大于该值

	PreFEndTime *string `json:"PreFEndTime,omitempty" name:"PreFEndTime"`
	// 结束时间小于该值

	TailFEndTime *string `json:"TailFEndTime,omitempty" name:"TailFEndTime"`
	// 创建时间大于该值

	PreFCreateTime *string `json:"PreFCreateTime,omitempty" name:"PreFCreateTime"`
	// 创建时间小于该值

	TailFCreateTime *string `json:"TailFCreateTime,omitempty" name:"TailFCreateTime"`
	// 更新时间大于该值

	PreFUpdateTime *string `json:"PreFUpdateTime,omitempty" name:"PreFUpdateTime"`
	// 更新时间小于该值

	TailFUpdateTime *string `json:"TailFUpdateTime,omitempty" name:"TailFUpdateTime"`
	// 用户Uin

	TargetUserId *uint64 `json:"TargetUserId,omitempty" name:"TargetUserId"`
	// 导出Excel需要的表头

	Fields []*string `json:"Fields,omitempty" name:"Fields"`
}

func (r *ExportDiscountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDiscountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancleContractTypeGatewayRequest struct {
	*tchttp.BaseRequest

	// 合同类型唯一ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 状态1正常2作废

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *CancleContractTypeGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancleContractTypeGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByRegionGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillSummaryByRegionGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByRegionGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInvoiceDetailRequest struct {
	*tchttp.BaseRequest

	// 用户开票流水号

	OpenId *int64 `json:"OpenId,omitempty" name:"OpenId"`
}

func (r *DescribeInvoiceDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInvoiceDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSignatureResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求状态标识

		SuccessActionStatus *int64 `json:"SuccessActionStatus,omitempty" name:"SuccessActionStatus"`
		// 凭据

		XAmzCredential *string `json:"XAmzCredential,omitempty" name:"XAmzCredential"`
		// 算法

		XAmzAlgorithm *string `json:"XAmzAlgorithm,omitempty" name:"XAmzAlgorithm"`
		// 日期

		XAmzDate *string `json:"XAmzDate,omitempty" name:"XAmzDate"`
		// 密文

		Policy *string `json:"Policy,omitempty" name:"Policy"`
		// 签名

		XAmzSignature *string `json:"XAmzSignature,omitempty" name:"XAmzSignature"`
		// 地址

		Url *string `json:"Url,omitempty" name:"Url"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSignatureResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSignatureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyComprehensiveStrategyRequest struct {
	*tchttp.BaseRequest

	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品code

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 付费模式

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 时间周期 如 1,2,4,6

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 结算周期 如 3,4

	CalcUnit *string `json:"CalcUnit,omitempty" name:"CalcUnit"`
	// 交易类型 如 1,2,3,5

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
	// 计费模式 如1,2,3,5

	BillingType *string `json:"BillingType,omitempty" name:"BillingType"`
	// 新购是否同步发货 默认0

	PurchaseSyncDelivery *uint64 `json:"PurchaseSyncDelivery,omitempty" name:"PurchaseSyncDelivery"`
	// 变配是否同步发货 默认0

	ModifySyncDelivery *uint64 `json:"ModifySyncDelivery,omitempty" name:"ModifySyncDelivery"`
	// 定价类型,2:线性价格

	PriceType *uint64 `json:"PriceType,omitempty" name:"PriceType"`
	// 资源回收类型 默认1

	ResourceRecycleType *uint64 `json:"ResourceRecycleType,omitempty" name:"ResourceRecycleType"`
	// 接口名

	InterfaceNames *InterfaceNames `json:"InterfaceNames,omitempty" name:"InterfaceNames"`
	// url名

	Region2interface *string `json:"Region2interface,omitempty" name:"Region2interface"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyComprehensiveStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyComprehensiveStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductPropertyListValues struct {
	// 对应field_code的取值

	FieldCode *string `json:"FieldCode,omitempty" name:"FieldCode"`
	// 该value状态： 1-正常 2-失效

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DescribeDosagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 如果没有用量，值为0

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 如果没有用量，则空数组

		Dosages *DosageList `json:"Dosages,omitempty" name:"Dosages"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDosagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDosagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 记录数，不传NeedRecordNum或者传的值为0，该字段值为null

		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`
		// 产品花费详情

		SummaryDetail *BusinessSummaryDetailItem `json:"SummaryDetail,omitempty" name:"SummaryDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillFeeByTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeByTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductDownloadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载链接

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductDownloadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BspProductList struct {
	// 产品Code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品展示标识

	DisplayCode *string `json:"DisplayCode,omitempty" name:"DisplayCode"`
	// 产品展示名称

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 子产品数据量

	SubProductNum *string `json:"SubProductNum,omitempty" name:"SubProductNum"`
	// 产品项数量

	BillingItemNum *string `json:"BillingItemNum,omitempty" name:"BillingItemNum"`
}

type DescribeBillFeeByRegionGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 地域花费详情

		SummaryDetail []*RegionSummaryDetailItem `json:"SummaryDetail,omitempty" name:"SummaryDetail"`
		// 记录数，不传NeedRecordNum或者传的值为0，该字段值为null

		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillFeeByRegionGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeByRegionGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportFieldsItem struct {
	// 支付者UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 商品码名称

	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
	// 子商品码名称

	SubProductCodeName *string `json:"SubProductCodeName,omitempty" name:"SubProductCodeName"`
	// 计费模式

	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`
	// 地域

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 可用区

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 实例名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 交易类型

	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`
	// 订单ID

	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
	// 交易ID

	BillId *string `json:"BillId,omitempty" name:"BillId"`
	// 扣费时间

	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`
	// 开始使用时间

	FeeBeginTime *string `json:"FeeBeginTime,omitempty" name:"FeeBeginTime"`
	// 结束使用时间

	FeeEndTime *string `json:"FeeEndTime,omitempty" name:"FeeEndTime"`
	// 组件名称

	BillingItemCodeName *string `json:"BillingItemCodeName,omitempty" name:"BillingItemCodeName"`
	// 组件类型名称

	SubBillingItemCodeName *string `json:"SubBillingItemCodeName,omitempty" name:"SubBillingItemCodeName"`
	// 组件刊例价

	SinglePrice *string `json:"SinglePrice,omitempty" name:"SinglePrice"`
	// 组件价格单位

	SinglePriceUnit *string `json:"SinglePriceUnit,omitempty" name:"SinglePriceUnit"`
	// 组件用量

	UsedAmount *string `json:"UsedAmount,omitempty" name:"UsedAmount"`
	// 组件用量单位

	UsedAmountUnit *string `json:"UsedAmountUnit,omitempty" name:"UsedAmountUnit"`
	// 使用时长

	TimeSpan *string `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 时长单位

	TimeUnitName *string `json:"TimeUnitName,omitempty" name:"TimeUnitName"`
	// 组件原件

	Cost *string `json:"Cost,omitempty" name:"Cost"`
	// 优惠后总价

	RealCost *string `json:"RealCost,omitempty" name:"RealCost"`
	// 赠送账户支出(元)

	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
	// 应付金额

	PayAbleAmount *string `json:"PayAbleAmount,omitempty" name:"PayAbleAmount"`
	// 国内国际

	RegionTypeName *string `json:"RegionTypeName,omitempty" name:"RegionTypeName"`
	// 使用者UIN

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 操作者UIN

	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`
}

type Object struct {
	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 配额key

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
	// 配额值

	QuotaValue *uint64 `json:"QuotaValue,omitempty" name:"QuotaValue"`
}

type ModifyPublicAccountGatewayRequest struct {
	*tchttp.BaseRequest

	// 公共账号编号，默认1（只支持填1）

	AccountId *uint64 `json:"AccountId,omitempty" name:"AccountId"`
	// 收款人户名

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 收款卡号

	AccountNum *uint64 `json:"AccountNum,omitempty" name:"AccountNum"`
	// 收款银行

	Bank *string `json:"Bank,omitempty" name:"Bank"`
	// 备注

	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *ModifyPublicAccountGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPublicAccountGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveQuotaDataArray struct {
	// 如果传了就是需要新增或编辑

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// ProductCode的数据项 如果传了就是需要新增或编辑

	ProductQuotaValue *int64 `json:"ProductQuotaValue,omitempty" name:"ProductQuotaValue"`
	// 如果传了就是需要新增或编辑

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// SubProductCode 的数据项 如果传了就是需要新增或编辑

	SubProductQuotaValue *int64 `json:"SubProductQuotaValue,omitempty" name:"SubProductQuotaValue"`
	// 如果传了就是需要新增或编辑

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// BillingItemCode 的数据项 如果传了就是需要新增或编辑

	BillingItemQuotaValue *int64 `json:"BillingItemQuotaValue,omitempty" name:"BillingItemQuotaValue"`
	// 如果传了就是需要新增或编辑

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// SubBillingItemCode 的数据项 如果传了就是需要新增或编辑

	SubBillingItemQuotaValue *int64 `json:"SubBillingItemQuotaValue,omitempty" name:"SubBillingItemQuotaValue"`
}

type DescribeBillDownloadRecordListGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 每次获取数据量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 是否需要记录条数，0需要，1不需要，默认不需要

	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
	// 排序规则

	Sort *GetDownloadListSort `json:"Sort,omitempty" name:"Sort"`
	// 只支持FIleIds，FileTypes，Status三种过滤条件

	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
	// 是否需要条件取值（0不需要/1需要，默认为0）

	NeedConditionValue *string `json:"NeedConditionValue,omitempty" name:"NeedConditionValue"`
}

func (r *DescribeBillDownloadRecordListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadRecordListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDealPriceGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDealPriceGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDealPriceGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportDealListOpGatewayRequest struct {
	*tchttp.BaseRequest

	// 操作uin，查全量缺省

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 大订单或小订单

	DealOrBigDealId *string `json:"DealOrBigDealId,omitempty" name:"DealOrBigDealId"`
	// 订单号

	DealId *string `json:"DealId,omitempty" name:"DealId"`
	// 大订单号

	BigDealId *string `json:"BigDealId,omitempty" name:"BigDealId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 提单人

	Payer *uint64 `json:"Payer,omitempty" name:"Payer"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 订单类型

	DealAction *string `json:"DealAction,omitempty" name:"DealAction"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 排序asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *ExportDealListOpGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDealListOpGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReconciliationGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyReconciliationGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyReconciliationGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrderInfo struct {
	// 订单号

	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`
	// 米大师订单号

	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`
	// 用户id

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 支付渠道

	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`
	// ISO 货币代码，CNY

	CurrencyType *string `json:"CurrencyType,omitempty" name:"CurrencyType"`
	// 支付金额，单位：分。

	Amount *uint64 `json:"Amount,omitempty" name:"Amount"`
	// 当前订单的订单状态  0： 初始状态，获取Midas交易订单成功；   1： 拉起Midas支付页面成功，用户未支付；   2：用户支付成功，正在发货；   3：用户支付成功，发货失败；   4：用户支付成功，发货成功；   5：Midas支付页面正在失效中；   6：Midas支付页面已经失效；

	OrderState *string `json:"OrderState,omitempty" name:"OrderState"`
	// 1:已退款

	RefundFlag *string `json:"RefundFlag,omitempty" name:"RefundFlag"`
	// 支付时间

	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`
}

type BspBillingItemList struct {
	// DisplayName

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 产品项名称

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 产品项名称

	DisplayCode *string `json:"DisplayCode,omitempty" name:"DisplayCode"`
	// 产品展示名称

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
}

type DescribeQuotaDownloadGatewayRequest struct {
	*tchttp.BaseRequest

	// 搜索用户uin

	QuotaUin *string `json:"QuotaUin,omitempty" name:"QuotaUin"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

func (r *DescribeQuotaDownloadGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotaDownloadGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetProductStatusGatewayRequest struct {
	*tchttp.BaseRequest

	// 商品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项代码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项代码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 状态，1-生效 0-失效

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
}

func (r *SetProductStatusGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetProductStatusGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountWaterRequest struct {
	*tchttp.BaseRequest

	// 接入后付费账户的唯一标识

	AcctId *string `json:"AcctId,omitempty" name:"AcctId"`
	// 用户ID

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 传后付费固定额度子账户标识: CREDIT_FIXED

	SubAcctid *string `json:"SubAcctid,omitempty" name:"SubAcctid"`
	// 请求方设备，0表示PC，1表示手机

	Device *int64 `json:"Device,omitempty" name:"Device"`
	// UNIX时间戳（从格林威治时间1970年01月01日00时00分00秒起至现在的总秒数）

	Ts *int64 `json:"Ts,omitempty" name:"Ts"`
	// 查询流水开始时间，数字，年月日20150617或年月日时分秒20150617111855

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询流水结束时间，数字，年月日20150617或年月日时分秒20150617111855

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// num，查询流水条数（最大支持20）

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// cursor，查询流水的游标

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 若传了， 则根据流水类型查询。目前支持：
	// RefundDebt ： 还款流水
	// WriteOff：销账流水

	WaterFilter *string `json:"WaterFilter,omitempty" name:"WaterFilter"`
	// 需要查询销账记录时，可传入对应还款订单号

	BillNo *string `json:"BillNo,omitempty" name:"BillNo"`
}

func (r *DescribeAccountWaterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountWaterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubBillingItemsGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计费细项信息列表

		List []*ProductSubBillingInfoList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubBillingItemsGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubBillingItemsGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportDealListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportDealListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDealListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductGatewayRequest struct {
	*tchttp.BaseRequest

	// 主键ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteProductGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadQuotasListFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务是否已完成（0未完成/1已完成）

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 文件名称，当任务已完成时返回该字段

		Name *string `json:"Name,omitempty" name:"Name"`
		// 下载链接，当任务已完成时返回该字段

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDownloadQuotasListFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadQuotasListFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveProductStrategyRouteGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本次查询条数

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 跳过条数

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 产品策略数据

		Info []*ProductStrategy `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComprehensiveProductStrategyRouteGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveProductStrategyRouteGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLackingQuotaRecordGatewayRequest struct {
	*tchttp.BaseRequest

	// 配额不足类型；1：项目级配额不足  2：主账号级(租户级)配额不足 3：子账号级配额不足

	LackingType *string `json:"LackingType,omitempty" name:"LackingType"`
	// 关键字类型；ProjectId：关键字为项目id；AppId：关键字为租户appId；AccountUin：关键字为账号uin

	KeywordType *string `json:"KeywordType,omitempty" name:"KeywordType"`
	// 关键字

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 页码；从1开始

	PageNo *uint64 `json:"PageNo,omitempty" name:"PageNo"`
	// 单页数据

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 开始时间；例如：2021-09-06 00:00:01

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间；例如：2021-09-06 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 一层产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

func (r *ListLackingQuotaRecordGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLackingQuotaRecordGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiscountStatusGatewayRequest struct {
	*tchttp.BaseRequest

	// 折扣类型 0用户折扣 1官网折扣

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 折扣id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 新的折扣状态值 0未生效 1生效 2失效

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDiscountStatusGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiscountStatusGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddQuotaRequest struct {
	*tchttp.BaseRequest

	// 配额复杂类型

	SpreadPara []*AddQuotaPara `json:"SpreadPara,omitempty" name:"SpreadPara"`
}

func (r *AddQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateContractTypeGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateContractTypeGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateContractTypeGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContractListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 合同列表

		List []*ContractListData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContractListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContractListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceSummaryDetailComponentDetailItem struct {
	// 组件用量

	UsedAmount *string `json:"UsedAmount,omitempty" name:"UsedAmount"`
	// 组件码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 组件名称

	BillingItemCodeName *string `json:"BillingItemCodeName,omitempty" name:"BillingItemCodeName"`
	// 组件折后总价

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 组件折后总价占资源折后总价比率

	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`
}

type GetUserStrategyGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserStrategyGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserStrategyGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchDiscountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 折扣列表详情信息

		List []*SearchDiscountList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchDiscountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchDiscountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CouponsWaterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 代金券信息

		CouponsList []*CouponWater `json:"CouponsList,omitempty" name:"CouponsList"`
		// 记录总数

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CouponsWaterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CouponsWaterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubBillingItemsRequest struct {
	*tchttp.BaseRequest

	// midas appid

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 计费细项code列表

	SubBillingItemCodes []*string `json:"SubBillingItemCodes,omitempty" name:"SubBillingItemCodes"`
}

func (r *DescribeSubBillingItemsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubBillingItemsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetProductPriceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetProductPriceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetProductPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubscribeDestroyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubscribeDestroyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubscribeDestroyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConditionComponent struct {
	// 组件码

	ComponentCode *string `json:"ComponentCode,omitempty" name:"ComponentCode"`
	// 组件名称

	ComponentCodeName *string `json:"ComponentCodeName,omitempty" name:"ComponentCodeName"`
}

type EasyProduct struct {
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 商品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type DescribeQuotaAssignLogsListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 账户配额转移流水列表

		QuotaAssignList []*QuotaAssignList `json:"QuotaAssignList,omitempty" name:"QuotaAssignList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQuotaAssignLogsListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotaAssignLogsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceSummaryGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源信息

		List *BspResourceBill `json:"List,omitempty" name:"List"`
		// 总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListResourceSummaryGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceSummaryGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByProductGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 记录数，不传NeedRecordNum或者传的值为0，该字段值为null

		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`
		// 产品花费详情

		SummaryDetail []*BusinessSummaryDetailItem `json:"SummaryDetail,omitempty" name:"SummaryDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillFeeByProductGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeByProductGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubProduct struct {
	// 子产品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品中文名称

	ChnName *string `json:"ChnName,omitempty" name:"ChnName"`
	// 子产品英文名称

	EngName *string `json:"EngName,omitempty" name:"EngName"`
	// 子产品状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type GetProductTreeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品四层定义列表

		List []*ProductRelation `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductTreeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetHiddenProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetHiddenProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetHiddenProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceInfo struct {
	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// app ID

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 创建时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 资源状态

	BizStatus *string `json:"BizStatus,omitempty" name:"BizStatus"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 子产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品名称

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 地区ID

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 地区名

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 区域ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 销毁时间(新增字段，342前已销毁资源没有该数据）

	BizIsolatedTime *string `json:"BizIsolatedTime,omitempty" name:"BizIsolatedTime"`
	// 修改时间(也记录了状态变销毁的时间)

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 区域名

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 计费码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费名称

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 子计费码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 子计费名称

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
	// 资源数量

	Num *string `json:"Num,omitempty" name:"Num"`
	// 单位名称

	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`
}

type DescribeTenantSubUinQuotasGatewayRequest struct {
	*tchttp.BaseRequest

	// 计费MidasAppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 第一层产品定义的code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

func (r *DescribeTenantSubUinQuotasGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTenantSubUinQuotasGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpireDiscountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExpireDiscountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExpireDiscountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByPayModeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 付费模式花费分布

		SummaryOverview []*PayModeSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByPayModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByPayModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGeneralDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询总数

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 目标数据list

		List []*GeneralData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGeneralDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 资源状态信息结果集

		List []*ResourceInfo `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductDownloadGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载链接

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductDownloadGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductDownloadGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDiscountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1

		// 1

		// 1

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDiscountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDiscountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeByPayModeRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillFeeByPayModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeByPayModeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadVoucherListFileRequest struct {
	*tchttp.BaseRequest

	// 代金券筛选条件对象

	VoucherConditions *VoucherConditions `json:"VoucherConditions,omitempty" name:"VoucherConditions"`
}

func (r *DescribeDownloadVoucherListFileRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadVoucherListFileRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemitRecordGatewayRequest struct {
	*tchttp.BaseRequest

	// 目标用户uin

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 汇款户名

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 汇款账号

	AccountNum *uint64 `json:"AccountNum,omitempty" name:"AccountNum"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 状态 0：处理中 1：成功 2：失败

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeRemitRecordGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemitRecordGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveProductStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本次查询条数

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 跳过条数

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据

		Info []*ProductStrategy `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComprehensiveProductStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveProductStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewDetail struct {
	// 数量

	TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`
	// 续费列表

	RenewList []*RenewList `json:"RenewList,omitempty" name:"RenewList"`
}

type DescribeSubBillingItemsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计费细项信息列表

		List []*ProductSubBillingInfoList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubBillingItemsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubBillingItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyComprehensiveStrategyRouteGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyComprehensiveStrategyRouteGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyComprehensiveStrategyRouteGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyGatewayRequest struct {
	*tchttp.BaseRequest

	// 目标用户uin

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 资源销毁参数信息

	Resources *string `json:"Resources,omitempty" name:"Resources"`
}

func (r *DestroyGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CouponList struct {
	// 计费MidasAppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 发券开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 发券结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 券类型，card: 卡券，balance: 余额券

	CouponType *string `json:"CouponType,omitempty" name:"CouponType"`
	// 活动号

	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`
	// 状态  1为正常，<=0为下线状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 活动名称

	ActivityName *string `json:"ActivityName,omitempty" name:"ActivityName"`
	// 有效期类型

	ExpiryTimeType *string `json:"ExpiryTimeType,omitempty" name:"ExpiryTimeType"`
	// 有效时间值

	TimeValue *int64 `json:"TimeValue,omitempty" name:"TimeValue"`
	// 有效时间单位，D天，H小时

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 固定有效时间开始

	ExpiryTimeBegin *string `json:"ExpiryTimeBegin,omitempty" name:"ExpiryTimeBegin"`
	// 固定有效时间结束

	ExpiryTimeEnd *string `json:"ExpiryTimeEnd,omitempty" name:"ExpiryTimeEnd"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最后更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 券门槛（分）

	CouponThreshold *int64 `json:"CouponThreshold,omitempty" name:"CouponThreshold"`
	// 券面值（分）

	CouponValue *int64 `json:"CouponValue,omitempty" name:"CouponValue"`
	// 付费模式 0-预付费，1-后付费，0,1-全部

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 四层产品定义

	ProductInfo *string `json:"ProductInfo,omitempty" name:"ProductInfo"`
	// 创建用户

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// 用户领取限制

	UserLimit *UserLimit `json:"UserLimit,omitempty" name:"UserLimit"`
	// 发放总限制

	ActivityLimit *ActivityLimit `json:"ActivityLimit,omitempty" name:"ActivityLimit"`
}

type AddQuotaGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddQuotaGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddQuotaGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaDownloadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载链接

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQuotaDownloadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotaDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveUserDiscountRequest struct {
	*tchttp.BaseRequest

	// 折扣参数

	DiscountPara []*DiscountPara `json:"DiscountPara,omitempty" name:"DiscountPara"`
}

func (r *SaveUserDiscountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveUserDiscountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubUinQuotaList struct {
	// 主用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 当前登录的子账户Uin

	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品中文名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品英文名称

	ProductEngName *string `json:"ProductEngName,omitempty" name:"ProductEngName"`
	// 子产品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品中文名称

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 子产品英文名称

	SubProductEngName *string `json:"SubProductEngName,omitempty" name:"SubProductEngName"`
	// 计费项代码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费项中文名称

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 计费项英文名称

	BillingItemEngName *string `json:"BillingItemEngName,omitempty" name:"BillingItemEngName"`
	// 计费细项代码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 计费细项中文名称

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
	// 计费细项英文名称

	SubBillingItemEngName *string `json:"SubBillingItemEngName,omitempty" name:"SubBillingItemEngName"`
	// 配额值

	QuotaValue *int64 `json:"QuotaValue,omitempty" name:"QuotaValue"`
}

type CreateContractTypeRequest struct {
	*tchttp.BaseRequest

	// 类型名称

	ContractTypeName *string `json:"ContractTypeName,omitempty" name:"ContractTypeName"`
	// 下载地址

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
}

func (r *CreateContractTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateContractTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceModifyLogGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 变配记录

		List *BspResourceBill `json:"List,omitempty" name:"List"`
		// 数据条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListResourceModifyLogGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceModifyLogGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDealListRequest struct {
	*tchttp.BaseRequest

	// 操作uin，查全量缺省

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 大订单或小订单

	DealOrBigDealId *string `json:"DealOrBigDealId,omitempty" name:"DealOrBigDealId"`
	// 订单号

	DealId *string `json:"DealId,omitempty" name:"DealId"`
	// 大订单号

	BigDealId *string `json:"BigDealId,omitempty" name:"BigDealId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 订单支付人uin

	Payer *uint64 `json:"Payer,omitempty" name:"Payer"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 订单类型

	DealAction *string `json:"DealAction,omitempty" name:"DealAction"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 排序asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 付费模式，0后付费/1预付费

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 产品名称

	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
	// 动作名称 隔离 销毁 返还 新购 续费 升配 降配

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 订单创建人uin

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 子产品名称

	SubProductCodeName *string `json:"SubProductCodeName,omitempty" name:"SubProductCodeName"`
}

func (r *DescribeDealListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDealListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetProductInfoGatewayRequest struct {
	*tchttp.BaseRequest

	// 产品定义

	Product *Product `json:"Product,omitempty" name:"Product"`
	// 子产品定义

	SubProduct *SubProduct `json:"SubProduct,omitempty" name:"SubProduct"`
	// 计费项定义

	BillingItem *BillingItem `json:"BillingItem,omitempty" name:"BillingItem"`
	// 计费细项定义

	SubBillingItem *SubBillingItem `json:"SubBillingItem,omitempty" name:"SubBillingItem"`
	// 别名

	AliasCode *string `json:"AliasCode,omitempty" name:"AliasCode"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 数据来源类型。local（本系统。默认），third（第三方接入）。不传默认是local

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
}

func (r *SetProductInfoGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetProductInfoGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Goods struct {
	// 订单类型ID，接入计费时由计费后台分配

	GoodsCategoryId *int64 `json:"GoodsCategoryId,omitempty" name:"GoodsCategoryId"`
	// 区域ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 可用区ID，没有可用区概念则传0

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 新购时表示商品数量，续费与配置变更时固定传1

	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 项目ID，没有项目概念则传0

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 付费模式，0后付费/1预付费

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// 平台，0开放平台/1云平台

	Platform *int64 `json:"Platform,omitempty" name:"Platform"`
	// 商品详情，timeSpan表示购买时间长度，timeUnit表示购买时间单位，pid表示定价公式ID

	GoodsDetail *string `json:"GoodsDetail,omitempty" name:"GoodsDetail"`
	// 预付费操作方式，隔离isolate,销毁release,返还return,新购purchase,续费renew,变配降配downgrade，变配升配upgrade

	Action *string `json:"Action,omitempty" name:"Action"`
}

type CreateContractTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateContractTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateContractTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDetailByResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 资源花费详情

		Total *ResourceSummaryDetailTotal `json:"Total,omitempty" name:"Total"`
		// 组件花费详情

		BillingItemDetail []*ResourceSummaryDetailComponentDetailItem `json:"BillingItemDetail,omitempty" name:"BillingItemDetail"`
		// 资源历史月份花费趋势

		RealTotalCostTrend *ResourceSummaryDetailRealTotalCostTrend `json:"RealTotalCostTrend,omitempty" name:"RealTotalCostTrend"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDetailByResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDetailByResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaDownloadGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载链接

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQuotaDownloadGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotaDownloadGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveStrategyWebConfigRouteGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计费模式

		PayMode []*string `json:"PayMode,omitempty" name:"PayMode"`
		// timeUnit

		TimeUnit []*string `json:"TimeUnit,omitempty" name:"TimeUnit"`
		// calcUnit

		CalcUnit []*string `json:"CalcUnit,omitempty" name:"CalcUnit"`
		// actionType

		ActionType []*string `json:"ActionType,omitempty" name:"ActionType"`
		// billingType

		BillingType []*string `json:"BillingType,omitempty" name:"BillingType"`
		// priceType

		PriceType []*string `json:"PriceType,omitempty" name:"PriceType"`
		// purchaseSyncDelivery

		PurchaseSyncDelivery []*string `json:"PurchaseSyncDelivery,omitempty" name:"PurchaseSyncDelivery"`
		// modifySyncDelivery

		ModifySyncDelivery []*string `json:"ModifySyncDelivery,omitempty" name:"ModifySyncDelivery"`
		// status

		Status []*string `json:"Status,omitempty" name:"Status"`
		// ResourceRecycleType

		ResourceRecycleType []*string `json:"ResourceRecycleType,omitempty" name:"ResourceRecycleType"`
		// region

		Region []*string `json:"Region,omitempty" name:"Region"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComprehensiveStrategyWebConfigRouteGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveStrategyWebConfigRouteGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomPropertiesRequest struct {
	*tchttp.BaseRequest

	// 计费MidasAppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项代码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项代码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 自定义售卖属性应用模块

	CustomPropertyModel *string `json:"CustomPropertyModel,omitempty" name:"CustomPropertyModel"`
}

func (r *DescribeCustomPropertiesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomPropertiesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadWithOperatorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDownloadWithOperatorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadWithOperatorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadResourceDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务是否已完成（0未完成/1已完成）

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 文件名称，当任务已完成时返回该字段

		Name *string `json:"Name,omitempty" name:"Name"`
		// 下载链接，当任务已完成时返回该字段

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDownloadResourceDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadResourceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveProductStrategyRequest struct {
	*tchttp.BaseRequest

	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 是否只查询第三方自定义产品策略,如果该值不传或者为false则查询全部产品策略

	IfCustom *bool `json:"IfCustom,omitempty" name:"IfCustom"`
	// 状态，如果该值不传则默认查询所有状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 本次查询条数，如果该值不传则默认100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 跳过条数如果该值不传则默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetComprehensiveProductStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveProductStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TrendByMonthDetail struct {
	// 当前月份统计

	Current *MonthCostDetailItem `json:"Current,omitempty" name:"Current"`
	// 历史月份统计

	Previous []*MonthCostDetailItem `json:"Previous,omitempty" name:"Previous"`
	// 未来月份统计

	Next *MonthCostDetailItem `json:"Next,omitempty" name:"Next"`
}

type GetProductDeployStatusGatewayRequest struct {
	*tchttp.BaseRequest

	// 产品码

	ProductCodes []*string `json:"ProductCodes,omitempty" name:"ProductCodes"`
}

func (r *GetProductDeployStatusGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductDeployStatusGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiscountGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDiscountGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiscountGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDealListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 大订单列表

		List []*BigDealListData `json:"List,omitempty" name:"List"`
		// 服务器当前时间戳

		Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDealListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDealListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillingSubUinQuotaGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集合

		List *SubUinQuota `json:"List,omitempty" name:"List"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillingSubUinQuotaGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillingSubUinQuotaGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiscountStatusGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 折扣id

		// 条数

		// 是否成功

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDiscountStatusGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiscountStatusGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillingItemsGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计费项信息列表

		List []*ProductBillingItemInfoList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillingItemsGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillingItemsGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceSummaryGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询的appId

	AccountAppId *uint64 `json:"AccountAppId,omitempty" name:"AccountAppId"`
	// 查询的UIN

	AccountUin *uint64 `json:"AccountUin,omitempty" name:"AccountUin"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 产品编码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 资源状态

	BizStatus *string `json:"BizStatus,omitempty" name:"BizStatus"`
	// 页序号，从0开始

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 每页数量

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 产品子编码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 录入开始时间

	StartAddTime *string `json:"StartAddTime,omitempty" name:"StartAddTime"`
	// 录入结束时间

	EndAddTime *string `json:"EndAddTime,omitempty" name:"EndAddTime"`
	// 资源配置变更查询时间范围，开始时间

	StartModifyTime *string `json:"StartModifyTime,omitempty" name:"StartModifyTime"`
	// 资源配置变更查询时间范围，截止时间

	EndModifyTime *string `json:"EndModifyTime,omitempty" name:"EndModifyTime"`
	// 资源销毁查询时间范围，开始时间

	StartIsolatedTime *string `json:"StartIsolatedTime,omitempty" name:"StartIsolatedTime"`
	// 资源销毁查询时间范围，截止时间

	EndIsolatedTime *string `json:"EndIsolatedTime,omitempty" name:"EndIsolatedTime"`
	// 是否包含详情

	IsDetail *bool `json:"IsDetail,omitempty" name:"IsDetail"`
	// 任务标志

	TaskFlag *bool `json:"TaskFlag,omitempty" name:"TaskFlag"`
}

func (r *ListResourceSummaryGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceSummaryGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDebateListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`
		// 账单欠费月份信息

		List []*DebateListData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDebateListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDebateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductDownloadRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProductDownloadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductDownloadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductRelationsGatewayRequest struct {
	*tchttp.BaseRequest

	// 商品代码或名称

	SearchProduct *string `json:"SearchProduct,omitempty" name:"SearchProduct"`
	// 子商品代码或名称

	SearchSubProduct *string `json:"SearchSubProduct,omitempty" name:"SearchSubProduct"`
	// 计费项代码或名称

	SearchBillingItem *string `json:"SearchBillingItem,omitempty" name:"SearchBillingItem"`
	// 计费细项代码或名称

	SearchSubBillingItem *string `json:"SearchSubBillingItem,omitempty" name:"SearchSubBillingItem"`
	// 四层定义状态，1-生效 0-失效

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 商品代码数组

	SearchProductCodes []*string `json:"SearchProductCodes,omitempty" name:"SearchProductCodes"`
	// 子商品代码数组

	SearchSubProductCodes []*string `json:"SearchSubProductCodes,omitempty" name:"SearchSubProductCodes"`
	// 计费项代码数组

	SearchBillingItemCodes []*string `json:"SearchBillingItemCodes,omitempty" name:"SearchBillingItemCodes"`
	// 计费细项代码数组

	SearchSubBillingItemCodes []*string `json:"SearchSubBillingItemCodes,omitempty" name:"SearchSubBillingItemCodes"`
	// 数据来源类型。local（本系统。默认），third（第三方接入）。不传默认是local

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
	// 数据是否过滤隐藏数据,true表示需要

	ExceptHidden *bool `json:"ExceptHidden,omitempty" name:"ExceptHidden"`
	// 过滤隐藏计量或计费数据标志1计费0计量,不传为全部隐藏

	IfBilling *int64 `json:"IfBilling,omitempty" name:"IfBilling"`
}

func (r *DescribeProductRelationsGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductRelationsGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportDealListOpGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportDealListOpGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDealListOpGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByProductGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillSummaryByProductGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByProductGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByTagGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
}

func (r *DescribeBillSummaryByTagGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByTagGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceDetailListDate struct {
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 资源到期时间

	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 资源自动续费标识

	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 地域

	Region *uint64 `json:"Region,omitempty" name:"Region"`
	// 账号ID

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 产品详情信息

	GoodsDetail *GoodsDetailList `json:"GoodsDetail,omitempty" name:"GoodsDetail"`
	// 隔离时间

	IsolatedTimestamp *string `json:"IsolatedTimestamp,omitempty" name:"IsolatedTimestamp"`
	// 账号UIN

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 计费模式

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 项目ID

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 区

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type GetRuleActionsMappingGatewayRequest struct {
	*tchttp.BaseRequest
}

func (r *GetRuleActionsMappingGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRuleActionsMappingGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConditionRegion struct {
	// 地域ID

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

type MidasQueryCouponDo struct {
	// AppId

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// ActivityId

	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`
}

type SummaryByResourceConditionValue struct {
	// 产品列表

	Product []*ConditionBusiness `json:"Product,omitempty" name:"Product"`
	// 子产品列表

	SubProduct []*ConditionProduct `json:"SubProduct,omitempty" name:"SubProduct"`
	// 地域列表

	Region []*ConditionRegion `json:"Region,omitempty" name:"Region"`
	// 付费模式列表

	PayMode []*ConditionPayMode `json:"PayMode,omitempty" name:"PayMode"`
	// 交易类型列表

	ActionType []*ConditionActionType `json:"ActionType,omitempty" name:"ActionType"`
}

type DescribeSubProductsGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 子产品信息

		List []*SubProductDefineList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubProductsGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubProductsGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiscountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDiscountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDiscountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailConditionValue struct {
	// 产品列表

	Product []*ConditionBusiness `json:"Product,omitempty" name:"Product"`
	// 子产品列表

	SubProduct []*ConditionProduct `json:"SubProduct,omitempty" name:"SubProduct"`
	// 组件列表

	BillingItem []*ConditionComponent `json:"BillingItem,omitempty" name:"BillingItem"`
	// 地域列表

	Region []*ConditionRegion `json:"Region,omitempty" name:"Region"`
	// 付费模式列表

	PayMode []*ConditionPayMode `json:"PayMode,omitempty" name:"PayMode"`
	// 交易类型列表

	ActionType []*ConditionActionType `json:"ActionType,omitempty" name:"ActionType"`
}

type CancleContractTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancleContractTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancleContractTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportDealListOpRequest struct {
	*tchttp.BaseRequest

	// 操作uin，查全量缺省

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 大订单或小订单

	DealOrBigDealId *string `json:"DealOrBigDealId,omitempty" name:"DealOrBigDealId"`
	// 订单号

	DealId *string `json:"DealId,omitempty" name:"DealId"`
	// 大订单号

	BigDealId *string `json:"BigDealId,omitempty" name:"BigDealId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 提单人

	Payer *uint64 `json:"Payer,omitempty" name:"Payer"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 订单类型

	DealAction *string `json:"DealAction,omitempty" name:"DealAction"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 排序asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *ExportDealListOpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDealListOpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductTreeRequest struct {
	*tchttp.BaseRequest

	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 页索引

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 是否生效.0未生效 1生效

	Status *string `json:"Status,omitempty" name:"Status"`
	// 数据是否过滤隐藏数据,true表示需要

	ExceptHidden *bool `json:"ExceptHidden,omitempty" name:"ExceptHidden"`
	// 过滤隐藏计量或计费数据标志1计费0计量,不传为全部隐藏

	IfBilling *int64 `json:"IfBilling,omitempty" name:"IfBilling"`
}

func (r *GetProductTreeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductTreeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveSubProductStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 跳过条数

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 本次返回条数

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据

		Info []*ProductStrategy `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComprehensiveSubProductStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveSubProductStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyContractStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyContractStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyContractStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditRemitRecordRequest struct {
	*tchttp.BaseRequest

	// 目标uin

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 0:处理中、1:充值成功、2:充值失败、3:拒绝审核

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 审核数据id

	TargetId *uint64 `json:"TargetId,omitempty" name:"TargetId"`
	// 拒绝理由

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

func (r *AuditRemitRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuditRemitRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTempQuotaListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// list

		List *TempQuotaList `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTempQuotaListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTempQuotaListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductBillingItemInfoList struct {
	// 产品code（第一层）

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品计费项code（第三层）

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 中文名

	ChnName *string `json:"ChnName,omitempty" name:"ChnName"`
	// 英文名

	EngName *string `json:"EngName,omitempty" name:"EngName"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type RenewList struct {
	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 资源id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 地域id

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 区域id

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 时长

	TimeSpan *string `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 时长单位：1:年,2:月,3:日,4:小时,5:分钟,6:秒

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 付费类型。0:后付费,1:预付费

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 用户在业务侧的承载号码

	AppIdUserNum *string `json:"AppIdUserNum,omitempty" name:"AppIdUserNum"`
	// Paras

	Paras []*Paras `json:"Paras,omitempty" name:"Paras"`
	// 自动续费模式： 0:手动续费,到期后一定间隔时间后回收;1：自动续费; 2:到期关闭

	AutopayType *string `json:"AutopayType,omitempty" name:"AutopayType"`
	// 预计隔离时间,格式为datetime

	ExpectIsolateTime *string `json:"ExpectIsolateTime,omitempty" name:"ExpectIsolateTime"`
}

type DeleteProductQuotaRequest struct {
	*tchttp.BaseRequest

	// 用户uin，传空字符串就删除system默认配额

	QuotaUin *string `json:"QuotaUin,omitempty" name:"QuotaUin"`
	// 产品标签

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品标签

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项标签

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项标签

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
}

func (r *DeleteProductQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveCommonDiscountRequest struct {
	*tchttp.BaseRequest

	// 折扣参数

	DiscountPara []*DiscountPara `json:"DiscountPara,omitempty" name:"DiscountPara"`
}

func (r *SaveCommonDiscountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveCommonDiscountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResourceBillRequest struct {
	*tchttp.BaseRequest

	// 资源id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 分页大小(1~100)

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 页码(从0开始)

	PageNo *uint64 `json:"PageNo,omitempty" name:"PageNo"`
	// 业务侧的appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

func (r *GetResourceBillRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResourceBillRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDownloadRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDownloadRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDownloadRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDebtBillRequest struct {
	*tchttp.BaseRequest

	// 接入后付费账户的唯一标识

	AcctId *string `json:"AcctId,omitempty" name:"AcctId"`
	// 用户ID

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 传后付费固定额度子账户标识: CREDIT_FIXED

	SubAcctid *string `json:"SubAcctid,omitempty" name:"SubAcctid"`
	// 请求方设备，0表示PC，1表示手机

	Device *int64 `json:"Device,omitempty" name:"Device"`
	// UNIX时间戳（从格林威治时间1970年01月01日00时00分00秒起至现在的总秒数）

	Ts *int64 `json:"Ts,omitempty" name:"Ts"`
	// 账单起始年月yyyymm

	BeginMonth *string `json:"BeginMonth,omitempty" name:"BeginMonth"`
	// 账单结束年月yyyymm

	EndMonth *string `json:"EndMonth,omitempty" name:"EndMonth"`
	// 查询的账单还款状态，数字 1-已还清 2-未还清

	DebtbillStatus *int64 `json:"DebtbillStatus,omitempty" name:"DebtbillStatus"`
	// 查询的账单是否逾期，数字 1-未逾期 2-逾期

	OverDue *int64 `json:"OverDue,omitempty" name:"OverDue"`
	// num，查询账单条数（最大支持20）

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// cursor，查询流水的游标

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUserDebtBillRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserDebtBillRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveProdctStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本次查询条数

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 跳过条数

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 数据

		Info []*ProductStrategy `json:"Info,omitempty" name:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComprehensiveProdctStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveProdctStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveStrategyWebConfigRouteGatewayRequest struct {
	*tchttp.BaseRequest
}

func (r *GetComprehensiveStrategyWebConfigRouteGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveStrategyWebConfigRouteGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetProductInfoRequest struct {
	*tchttp.BaseRequest

	// 产品定义

	Product *Product `json:"Product,omitempty" name:"Product"`
	// 子产品定义

	SubProduct *SubProduct `json:"SubProduct,omitempty" name:"SubProduct"`
	// 计费项定义

	BillingItem *BillingItem `json:"BillingItem,omitempty" name:"BillingItem"`
	// 计费细项定义

	SubBillingItem *SubBillingItem `json:"SubBillingItem,omitempty" name:"SubBillingItem"`
	// 别名

	AliasCode *string `json:"AliasCode,omitempty" name:"AliasCode"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 数据来源类型。local（本系统。默认），third（第三方接入）。不传默认是local

	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`
}

func (r *SetProductInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetProductInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 总花费详情

		SummaryTotal *BusinessSummaryTotal `json:"SummaryTotal,omitempty" name:"SummaryTotal"`
		// 各产品花费分布

		SummaryOverview []*BusinessSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaAssignLogsListRequest struct {
	*tchttp.BaseRequest

	// 计费MidasAppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 例如：“2019-10-10 10:05:01”

	CreateStartTime *string `json:"CreateStartTime,omitempty" name:"CreateStartTime"`
	// 例如：“2019-10-12 10:05:02”

	CreateEndTime *string `json:"CreateEndTime,omitempty" name:"CreateEndTime"`
	// uin

	SearchUin *string `json:"SearchUin,omitempty" name:"SearchUin"`
	// 产品标签

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品标签

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 项目标签

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 子项目标签

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 第几页，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 一页行数，默认30

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 配额类型：sys或uin

	QuotaType *string `json:"QuotaType,omitempty" name:"QuotaType"`
}

func (r *DescribeQuotaAssignLogsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotaAssignLogsListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuotaList struct {
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品配额值

	ProductQuotaValue *uint64 `json:"ProductQuotaValue,omitempty" name:"ProductQuotaValue"`
	// 子产品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品名称

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 子产品配额值

	SubProductQuotaValue *uint64 `json:"SubProductQuotaValue,omitempty" name:"SubProductQuotaValue"`
	// 计费项代码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费项名称

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 计费项配额值

	BillingItemQuotaValue *uint64 `json:"BillingItemQuotaValue,omitempty" name:"BillingItemQuotaValue"`
	// 计费细项代码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 计费细项名称

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
	// 计费细项配额值

	SubBillingItemQuotaValue *uint64 `json:"SubBillingItemQuotaValue,omitempty" name:"SubBillingItemQuotaValue"`
	// 产品英文名称

	ProductEngName *string `json:"ProductEngName,omitempty" name:"ProductEngName"`
	// 子产品英文名称

	SubProductEngName *string `json:"SubProductEngName,omitempty" name:"SubProductEngName"`
	// 产品项英文名称

	BillingItemEngName *string `json:"BillingItemEngName,omitempty" name:"BillingItemEngName"`
	// 产品细项英文名称

	SubBillingItemEngName *string `json:"SubBillingItemEngName,omitempty" name:"SubBillingItemEngName"`
}

type TrendByMonthStat struct {
	// 平均花费详情

	Average *TrendByMonthStatItem `json:"Average,omitempty" name:"Average"`
}

type DescribeCustomPropertiesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 列表数据

		CustomPropertyList *CustomPropertyList `json:"CustomPropertyList,omitempty" name:"CustomPropertyList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomPropertiesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomPropertiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReconciliationListRequest struct {
	*tchttp.BaseRequest

	// 每次获取数据量

	Limit *string `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 调账记录查询调账列表部分参数

	Conditions *ReconciliationConditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeReconciliationListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReconciliationListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBSPSubProductGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品二层定义列表

		List []*BspSubProductList `json:"List,omitempty" name:"List"`
		// 产品二层定总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBSPSubProductGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBSPSubProductGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadWithAuthGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDownloadWithAuthGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadWithAuthGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditBreakdownModeGatewayRequest struct {
	*tchttp.BaseRequest

	// 故障模式操作，2开启恢复模式 3开启故障模式

	Status *string `json:"Status,omitempty" name:"Status"`
	// 白名单列表，如果不传则不修改，否则删除原有数据，并插入所传的列表

	WhiteList []*string `json:"WhiteList,omitempty" name:"WhiteList"`
	// 临时配额开关，1关闭 2开启

	TempQuto *string `json:"TempQuto,omitempty" name:"TempQuto"`
}

func (r *EditBreakdownModeGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EditBreakdownModeGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditRemitRecordGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AuditRemitRecordGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuditRemitRecordGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductQuotaGatewayRequest struct {
	*tchttp.BaseRequest

	// 用户uin，传空字符串就删除system默认配额

	QuotaUin *string `json:"QuotaUin,omitempty" name:"QuotaUin"`
	// 产品标签

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品标签

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项标签

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项标签

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
}

func (r *DeleteProductQuotaGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductQuotaGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadResourceDetailRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 需要导出相关字段

	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeBillDownloadResourceDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadResourceDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCouponListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分页

		Pagination *Pagination `json:"Pagination,omitempty" name:"Pagination"`
		// 数据

		ActivityList []*CouponDetail `json:"ActivityList,omitempty" name:"ActivityList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCouponListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCouponListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBillingProductCodeGatewayRequest struct {
	*tchttp.BaseRequest
}

func (r *GetBillingProductCodeGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBillingProductCodeGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillingItemsGatewayRequest struct {
	*tchttp.BaseRequest

	// midas appid

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 计费项code列表

	BillingItemCodes []*string `json:"BillingItemCodes,omitempty" name:"BillingItemCodes"`
}

func (r *DescribeBillingItemsGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillingItemsGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemitRecordOpResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 列表

		List *RemitRecordData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemitRecordOpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemitRecordOpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveQuotaGatewayRequest struct {
	*tchttp.BaseRequest

	// 用户uin。传空字符串就新增或编辑system默认配额

	QuotaUin *string `json:"QuotaUin,omitempty" name:"QuotaUin"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 要新增或更新的数据数组

	SaveDataArray []*SaveQuotaDataArray `json:"SaveDataArray,omitempty" name:"SaveDataArray"`
}

func (r *SaveQuotaGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveQuotaGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBspProductInfoGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetBspProductInfoGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBspProductInfoGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBSPSubBillingItemGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品四层定义列表

		List []*BspSubBillingItemList `json:"List,omitempty" name:"List"`
		// 产品四层定义总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBSPSubBillingItemGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBSPSubBillingItemGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductsGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品列表

		List []*Product `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductsGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductsGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReconciliationListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 每次获取数据量，与请求包一致

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 偏移量，与请求包一致

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 记录数，若调用时传的needRecordNum=0，则不返回该字段

		RecordNum *string `json:"RecordNum,omitempty" name:"RecordNum"`
		// 调账列表详情

		Data []*ReconciliationDataItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReconciliationListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReconciliationListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DosageList struct {
	// 产品Id

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 地域Id

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 推量周期

	CycleTimes *string `json:"CycleTimes,omitempty" name:"CycleTimes"`
	// 推量数

	CycleNum *uint64 `json:"CycleNum,omitempty" name:"CycleNum"`
	// 最后推量时间

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
}

type DescribeBillSummaryByRegionRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillSummaryByRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillSummaryByRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductPricesRequest struct {
	*tchttp.BaseRequest

	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
}

func (r *DescribeProductPricesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductPricesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceDownloadResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载链接

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceDownloadResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveStrategyWebConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *GetComprehensiveStrategyWebConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveStrategyWebConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Pagination struct {
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 页码

	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`
	// 每页数量

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeBillDownloadByResourceGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务是否已完成（0未完成/1已完成）

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 文件名称，当任务已完成时返回该字段

		Name *string `json:"Name,omitempty" name:"Name"`
		// 下载链接，当任务已完成时返回该字段

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDownloadByResourceGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadByResourceGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeTrendRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 只支持ProductCodes，RegionId，PayMode

	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeBillFeeTrendRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeTrendRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetHiddenProductGatewayRequest struct {
	*tchttp.BaseRequest

	// 隐藏商品项信息

	ProductInfo []*HiddenProduct `json:"ProductInfo,omitempty" name:"ProductInfo"`
}

func (r *SetHiddenProductGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetHiddenProductGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProductGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillFeeTrendGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 只支持ProductCodes，RegionId，PayMode

	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeBillFeeTrendGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeTrendGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAllProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品一层名称和code

		List []*ProductCodeAndName `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAllProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAllProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductDeployStatusGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 出参需要使用字符串包装来满足apiv3规范

		StatusJsonStr *string `json:"StatusJsonStr,omitempty" name:"StatusJsonStr"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductDeployStatusGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductDeployStatusGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetResourceBillGatewayRequest struct {
	*tchttp.BaseRequest

	// 资源id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 分页大小(1~100)

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 页码(从0开始)

	PageNo *uint64 `json:"PageNo,omitempty" name:"PageNo"`
	// 业务侧的appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

func (r *GetResourceBillGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetResourceBillGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyContractTypeRequest struct {
	*tchttp.BaseRequest

	// 合同类型唯一ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 类型名称

	ContractTypeName *string `json:"ContractTypeName,omitempty" name:"ContractTypeName"`
	// 下载链接

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
}

func (r *ModifyContractTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyContractTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountListRequest struct {
	*tchttp.BaseRequest

	// 管理端查询用户UIN，默认缺省查全部

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAccountListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPublicAccountGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPublicAccountGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPublicAccountGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadResourceDetailGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 需要导出相关字段

	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeBillDownloadResourceDetailGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadResourceDetailGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBSPProductGatewayRequest struct {
	*tchttp.BaseRequest

	// 商品代码或名称

	SearchPro0duct *string `json:"SearchPro0duct,omitempty" name:"SearchPro0duct"`
	// 子商品代码或名称

	SearchSubProduct *string `json:"SearchSubProduct,omitempty" name:"SearchSubProduct"`
	// 计费项代码或名称

	SearchBillingItem *string `json:"SearchBillingItem,omitempty" name:"SearchBillingItem"`
	// 计费细项代码或名称

	SearchSubBillingItem *string `json:"SearchSubBillingItem,omitempty" name:"SearchSubBillingItem"`
	// 四层定义状态，1-生效 0-失效 （不传默认生效

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 商品代码

	SearchProductCode *string `json:"SearchProductCode,omitempty" name:"SearchProductCode"`
	// 子商品代码

	SearchSubProductCode *string `json:"SearchSubProductCode,omitempty" name:"SearchSubProductCode"`
	// 计费项代码

	SearchBillingItemCode *string `json:"SearchBillingItemCode,omitempty" name:"SearchBillingItemCode"`
	// 计费细项代码

	SearchSubBillingItemCode *string `json:"SearchSubBillingItemCode,omitempty" name:"SearchSubBillingItemCode"`
	// 运营模式：bsp-计量、billing 计费（不传默认计费

	Source *string `json:"Source,omitempty" name:"Source"`
}

func (r *QueryBSPProductGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBSPProductGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CouponInfo struct {
	// 持有者

	UserNum *string `json:"UserNum,omitempty" name:"UserNum"`
	// 代金券名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 代金券codeid

	CodeId *string `json:"CodeId,omitempty" name:"CodeId"`
	// 批次id

	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
	// 面值

	Value *uint64 `json:"Value,omitempty" name:"Value"`
	// 余额

	Balance *uint64 `json:"Balance,omitempty" name:"Balance"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 到期时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 付费类型

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 发放时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 适用产品

	ProductInfo *string `json:"ProductInfo,omitempty" name:"ProductInfo"`
	// 代金券id

	SpId *string `json:"SpId,omitempty" name:"SpId"`
	// 代金券批次创建时间

	BatchCreateTime *string `json:"BatchCreateTime,omitempty" name:"BatchCreateTime"`
	// 发放者

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
}

type AddProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicAccountListRequest struct {
	*tchttp.BaseRequest

	// 用户uin数组

	Uins []*string `json:"Uins,omitempty" name:"Uins"`
}

func (r *DescribeBasicAccountListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBasicAccountListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceDownloadGatewayRequest struct {
	*tchttp.BaseRequest

	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 资源id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 时间范围-开始时间

	StartAddTime *string `json:"StartAddTime,omitempty" name:"StartAddTime"`
	// 时间范围-结束时间

	EndAddTime *string `json:"EndAddTime,omitempty" name:"EndAddTime"`
}

func (r *DescribeResourceDownloadGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceDownloadGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAcctRequest struct {
	*tchttp.BaseRequest

	// 计平分配的appid

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 用户id

	UserId *string `json:"UserId,omitempty" name:"UserId"`
	// 支付订单号，需要确保唯一

	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`
	// 4层定义信息

	ProductInfo *ProductCodeInfo `json:"ProductInfo,omitempty" name:"ProductInfo"`
	// 计平分配

	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
	// 调账原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 调整金额，单位：分

	Amount *uint64 `json:"Amount,omitempty" name:"Amount"`
	// 调账类型
	// add：补偿
	// minus：扣费

	AdjustType *string `json:"AdjustType,omitempty" name:"AdjustType"`
	// 固定传：dpacct

	Channel *string `json:"Channel,omitempty" name:"Channel"`
	// 支付发生的实际时间（UNIX时间戳），出账时将算入该时间所在月份的账单，而不是交易发生的当前月。

	ActualTime *int64 `json:"ActualTime,omitempty" name:"ActualTime"`
}

func (r *ModifyAcctRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAcctRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductRequest struct {
	*tchttp.BaseRequest

	// 产品大类中文名

	ProductGroupName *string `json:"ProductGroupName,omitempty" name:"ProductGroupName"`
	// 产品大类英文名

	ProductGroupEngName *string `json:"ProductGroupEngName,omitempty" name:"ProductGroupEngName"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品英文名称

	ProductEngName *string `json:"ProductEngName,omitempty" name:"ProductEngName"`
	// 子产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品名称

	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
	// 子产品英文名称

	SubProductEngName *string `json:"SubProductEngName,omitempty" name:"SubProductEngName"`
	// 计费项code

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费项名称

	BillingItemName *string `json:"BillingItemName,omitempty" name:"BillingItemName"`
	// 计费项英文名称

	BillingItemEngName *string `json:"BillingItemEngName,omitempty" name:"BillingItemEngName"`
	// 计费细项代码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 计费细项名称

	SubBillingItemName *string `json:"SubBillingItemName,omitempty" name:"SubBillingItemName"`
	// 计费细项英文名

	SubBillingItemEngName *string `json:"SubBillingItemEngName,omitempty" name:"SubBillingItemEngName"`
	// 单位中文名

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 单位英文名

	UnitEng *string `json:"UnitEng,omitempty" name:"UnitEng"`
	// 计费周期

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 结算周期

	CalcUnit *string `json:"CalcUnit,omitempty" name:"CalcUnit"`
	// 主键ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *UpdateProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRefundResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 提现(退款)明细

		RefundDetail []*RefundDetail `json:"RefundDetail,omitempty" name:"RefundDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRefundResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRefundResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceBillDownloadGatewayRequest struct {
	*tchttp.BaseRequest

	// UIN

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// appid

	AppId *string `json:"AppId,omitempty" name:"AppId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 资源Id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 标签组合字符串

	Tags *string `json:"Tags,omitempty" name:"Tags"`
}

func (r *DescribeResourceBillDownloadGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceBillDownloadGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchDiscountGatewayRequest struct {
	*tchttp.BaseRequest

	// 折扣类型  0用户折扣 1官网折扣

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 折扣ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 开始生效时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 失效时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 折扣创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 折扣状态 0:未生效 1:生效；2:已失效

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 用户UIN

	TargetUserId *uint64 `json:"TargetUserId,omitempty" name:"TargetUserId"`
	// 页码 从1开始

	Page *uint64 `json:"Page,omitempty" name:"Page"`
	// 分页大小

	Rows *uint64 `json:"Rows,omitempty" name:"Rows"`
	// 开始时间大于该值

	PreFBeginTime *string `json:"PreFBeginTime,omitempty" name:"PreFBeginTime"`
	// 开始时间小于该值

	TailFBeginTime *string `json:"TailFBeginTime,omitempty" name:"TailFBeginTime"`
	// 结束时间大于该值

	PreFEndTime *string `json:"PreFEndTime,omitempty" name:"PreFEndTime"`
	// 结束时间小于该值

	TailFEndTime *string `json:"TailFEndTime,omitempty" name:"TailFEndTime"`
	// 创建时间大于该值

	PreFCreateTime *string `json:"PreFCreateTime,omitempty" name:"PreFCreateTime"`
	// 创建时间小于该值

	TailFCreateTime *string `json:"TailFCreateTime,omitempty" name:"TailFCreateTime"`
	// 更新时间大于该值（用户折扣preFUpdateTime表示优先级）

	PreFUpdateTime *string `json:"PreFUpdateTime,omitempty" name:"PreFUpdateTime"`
	// 更新时间小于该值（用户折扣preFUpdateTime表示优先级）

	TailFUpdateTime *string `json:"TailFUpdateTime,omitempty" name:"TailFUpdateTime"`
	// 用户更新时间大于该值（只对用户折扣有效）

	PreFAutoUpdateTime *string `json:"PreFAutoUpdateTime,omitempty" name:"PreFAutoUpdateTime"`
	// 用户更新时间小于该值（只对用户折扣有效）

	TailFAutoUpdateTime *string `json:"TailFAutoUpdateTime,omitempty" name:"TailFAutoUpdateTime"`
	// 付费模式

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
}

func (r *SearchDiscountGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchDiscountGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemitRecordOpGatewayRequest struct {
	*tchttp.BaseRequest

	// 目标用户uin

	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
	// 汇款户名

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 汇款账号

	AccountNum *uint64 `json:"AccountNum,omitempty" name:"AccountNum"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRemitRecordOpGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemitRecordOpGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBSPProductGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品一层定义列表

		List []*BspProductList `json:"List,omitempty" name:"List"`
		// 数据总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBSPProductGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBSPProductGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceSummaryDownloadGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询的appId

	AccountAppId *uint64 `json:"AccountAppId,omitempty" name:"AccountAppId"`
	// 查询的UIN

	AccountUin *uint64 `json:"AccountUin,omitempty" name:"AccountUin"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 产品编码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 资源状态

	BizStatus *string `json:"BizStatus,omitempty" name:"BizStatus"`
	// 可用区ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 产品子编码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
}

func (r *ListResourceSummaryDownloadGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceSummaryDownloadGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReversalResourceGatewayRequest struct {
	*tchttp.BaseRequest

	// 目标用户

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

func (r *ReversalResourceGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReversalResourceGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DebateListData struct {
	// 用户uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 欠费月份数量

	DebtNum *string `json:"DebtNum,omitempty" name:"DebtNum"`
	// 账户登录名

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
}

type DeleteUserStrategyGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserStrategyGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserStrategyGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDealListGatewayRequest struct {
	*tchttp.BaseRequest

	// 操作uin，查全量缺省

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 大订单或小订单

	DealOrBigDealId *string `json:"DealOrBigDealId,omitempty" name:"DealOrBigDealId"`
	// 订单号

	DealId *string `json:"DealId,omitempty" name:"DealId"`
	// 大订单号

	BigDealId *string `json:"BigDealId,omitempty" name:"BigDealId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 订单支付人uin

	Payer *uint64 `json:"Payer,omitempty" name:"Payer"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 订单类型

	DealAction *string `json:"DealAction,omitempty" name:"DealAction"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 排序asc desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 付费模式，0后付费/1预付费

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 产品名称

	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
	// 动作名称 隔离 销毁 返还 新购 续费 升配 降配

	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`
	// 订单创建人uin

	Creater *string `json:"Creater,omitempty" name:"Creater"`
	// 子产品名称

	SubProductCodeName *string `json:"SubProductCodeName,omitempty" name:"SubProductCodeName"`
}

func (r *DescribeDealListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDealListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceDetailListGatewayRequest struct {
	*tchttp.BaseRequest

	// 资源信息

	ResourcePara *string `json:"ResourcePara,omitempty" name:"ResourcePara"`
	// 用户Appid

	UserAppid *int64 `json:"UserAppid,omitempty" name:"UserAppid"`
	// 用户Uin

	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`
}

func (r *DescribeResourceDetailListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceDetailListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicAccountListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 基本账户信息List

		List []*BasicAccount `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBasicAccountListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBasicAccountListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadVoucherListFileResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务是否完成，0未完成，1完成

		TaskFinished *uint64 `json:"TaskFinished,omitempty" name:"TaskFinished"`
		// 文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 文件名称

		FileName *string `json:"FileName,omitempty" name:"FileName"`
		// 任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDownloadVoucherListFileResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadVoucherListFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaDownloadRequest struct {
	*tchttp.BaseRequest

	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

func (r *DescribeQuotaDownloadRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotaDownloadRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditInvoiceInfoGatewayRequest struct {
	*tchttp.BaseRequest

	// 发票流水id

	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`
	// 审核状态 取值为(2:审核通过 3:审核驳回)

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 审核驳回原因,如果审核通过可以为空

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

func (r *AuditInvoiceInfoGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuditInvoiceInfoGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`
		// 帐户信息列表

		List []*AccountListData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryHiddenProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 隐藏商品项

		ProductInfo []*HiddenProduct `json:"ProductInfo,omitempty" name:"ProductInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryHiddenProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryHiddenProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillingItemsRequest struct {
	*tchttp.BaseRequest

	// midas appid

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 计费项code列表

	BillingItemCodes []*string `json:"BillingItemCodes,omitempty" name:"BillingItemCodes"`
}

func (r *DescribeBillingItemsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillingItemsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsolidatedBillDownloadUrlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务是否完成，0未完成，1完成

		TaskFinished *uint64 `json:"TaskFinished,omitempty" name:"TaskFinished"`
		// 文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 文件名称

		FileName *string `json:"FileName,omitempty" name:"FileName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsolidatedBillDownloadUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsolidatedBillDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tags struct {
	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue []*string `json:"TagValue,omitempty" name:"TagValue"`
}

type AddQuotaGatewayRequest struct {
	*tchttp.BaseRequest

	// 配额复杂类型

	SpreadPara []*AddQuotaPara `json:"SpreadPara,omitempty" name:"SpreadPara"`
}

func (r *AddQuotaGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddQuotaGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGeneralDataGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询总数

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 目标数据list

		List []*GeneralData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGeneralDataGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralDataGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserAccountGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户状态，0为未开户或销户，1为正常，其他异常
		// （可根据账户状态来判断是否开户）

		Status *string `json:"Status,omitempty" name:"Status"`
		// 后付费余额（包含了cash）

		Balance *string `json:"Balance,omitempty" name:"Balance"`
		// 累计金额，后付费账户表示累积还款额（当前接口无需关心）

		AllIn *string `json:"AllIn,omitempty" name:"AllIn"`
		// 历史累计支出金额，后付费账户表示累积消耗额（当前接口无需关心）

		AllOut *string `json:"AllOut,omitempty" name:"AllOut"`
		// 当天累计消耗

		DayOut *string `json:"DayOut,omitempty" name:"DayOut"`
		// 过期时间，非过期类返回0，否则返回时间UNIX时间戳（当前接口无需关心）

		Exptime *string `json:"Exptime,omitempty" name:"Exptime"`
		// 过期订单信息（当前接口无需关心）

		ExpireinfoExtend *string `json:"ExpireinfoExtend,omitempty" name:"ExpireinfoExtend"`
		// 后付费额度上限

		Dplimit *string `json:"Dplimit,omitempty" name:"Dplimit"`
		// 后付费剩余可用额度（包含了cash）

		Dpbalance *string `json:"Dpbalance,omitempty" name:"Dpbalance"`
		// 当前额度使用量（包含转出和预授权等）

		Debt *string `json:"Debt,omitempty" name:"Debt"`
		// 后付费账户，存储现金/溢出金额，下次出账时抵消欠款

		Cash *string `json:"Cash,omitempty" name:"Cash"`
		// 出账日，1-28的数字，每个月第几天出账单

		BillingDate *string `json:"BillingDate,omitempty" name:"BillingDate"`
		// 还款日，1-28的数字，每个月第几天还款

		DueDate *string `json:"DueDate,omitempty" name:"DueDate"`
		// 商户类型，1：代理，2：子客，3：直客

		MerchantType *string `json:"MerchantType,omitempty" name:"MerchantType"`
		// 关联商户信息，当商户类型为代理或直客时，该字段为空；当为子客时，标识代理账户信息

		RelatedMerchant *string `json:"RelatedMerchant,omitempty" name:"RelatedMerchant"`
		// 账期类型，1：按月，2：按天

		DueDelayType *string `json:"DueDelayType,omitempty" name:"DueDelayType"`
		// 账期，还款允许延期时间，当账期类型按月时，表示月份数，按天时，表示天数

		DueDelay *string `json:"DueDelay,omitempty" name:"DueDelay"`
		// 下一个出账日将生效的新账期（无新账期则返回空）

		NextDueDelay *string `json:"NextDueDelay,omitempty" name:"NextDueDelay"`
		// 还款时账单未还清是否立即恢复额度
		// 0 (默认)不立即恢复 1 立即恢复

		RecoverImmediately *string `json:"RecoverImmediately,omitempty" name:"RecoverImmediately"`
		// 是否校验账单还款顺序
		// 0 (默认)按顺序  1 不按顺序

		RefundSequence *string `json:"RefundSequence,omitempty" name:"RefundSequence"`
		// 提前还款是否立即恢复额度
		// 0 (默认)不立即恢复，出账时抵消账单才恢复；  1 立即恢复

		CashRecoverImmediately *string `json:"CashRecoverImmediately,omitempty" name:"CashRecoverImmediately"`
		// 出账周期类型，1：按月。其他暂不支持

		BillingCycleType *string `json:"BillingCycleType,omitempty" name:"BillingCycleType"`
		// 出账周期，若按月则表示月份数，按天则表示天数

		BillingCycle *string `json:"BillingCycle,omitempty" name:"BillingCycle"`
		// 二级额度（TCE的实际额度/固定额度），注意单位与后付费账户单位保持一致

		SubDplimit *string `json:"SubDplimit,omitempty" name:"SubDplimit"`
		// 已用额度

		UsedLimit *uint64 `json:"UsedLimit,omitempty" name:"UsedLimit"`
		// 现金余额

		CashBalance *uint64 `json:"CashBalance,omitempty" name:"CashBalance"`
		// 固定额度

		Fixedlimit *string `json:"Fixedlimit,omitempty" name:"Fixedlimit"`
		// 可用额度

		UsableLimit *uint64 `json:"UsableLimit,omitempty" name:"UsableLimit"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserAccountGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserAccountGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品四层定义列表

		List []*ProductRelation `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTagsByMonthGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// tagkey的数组，如果没有记录则为空数组

		TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTagsByMonthGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTagsByMonthGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateContractTypeGatewayRequest struct {
	*tchttp.BaseRequest

	// 类型名称

	ContractTypeName *string `json:"ContractTypeName,omitempty" name:"ContractTypeName"`
	// 下载地址

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
}

func (r *CreateContractTypeGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateContractTypeGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubProductsGatewayRequest struct {
	*tchttp.BaseRequest

	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 子产品代码,例如['sp_cvm','sp_cvm_2']

	SubProductCodes []*string `json:"SubProductCodes,omitempty" name:"SubProductCodes"`
	// 产品代码或名称

	SearchProduct *string `json:"SearchProduct,omitempty" name:"SearchProduct"`
	// 页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 状态，1-生效 0-失效

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeSubProductsGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubProductsGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInvoiceList struct {
	// 用户开票流水号

	OpenId *int64 `json:"OpenId,omitempty" name:"OpenId"`
	// 账号UIN

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 金额

	Amount *float64 `json:"Amount,omitempty" name:"Amount"`
	// 申请时间时间戳

	ApplicationTime *uint64 `json:"ApplicationTime,omitempty" name:"ApplicationTime"`
	// 发票类型

	UserType *uint64 `json:"UserType,omitempty" name:"UserType"`
	// 抬头

	InvoiceHead *string `json:"InvoiceHead,omitempty" name:"InvoiceHead"`
	// 税号

	TaxNumber *string `json:"TaxNumber,omitempty" name:"TaxNumber"`
	// 银行用户名

	BankUserName *string `json:"BankUserName,omitempty" name:"BankUserName"`
	// 账号

	AccountNumber *string `json:"AccountNumber,omitempty" name:"AccountNumber"`
	// 注册地址

	RegisterFullAddr *string `json:"RegisterFullAddr,omitempty" name:"RegisterFullAddr"`
	// 注册电话

	RegisterPhone *string `json:"RegisterPhone,omitempty" name:"RegisterPhone"`
	// 发票ID

	ArrId *uint64 `json:"ArrId,omitempty" name:"ArrId"`
	// 省

	Province *string `json:"Province,omitempty" name:"Province"`
	// 市/区

	City *string `json:"City,omitempty" name:"City"`
	// 详细地址

	Addr *string `json:"Addr,omitempty" name:"Addr"`
	// 联系方式

	Contact *string `json:"Contact,omitempty" name:"Contact"`
	// 电话号码

	Cellphone *string `json:"Cellphone,omitempty" name:"Cellphone"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 备注

	InvoiceRemark *string `json:"InvoiceRemark,omitempty" name:"InvoiceRemark"`
	// 账户时间范围

	AccountPeriods *string `json:"AccountPeriods,omitempty" name:"AccountPeriods"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 审核人

	Auditor *string `json:"Auditor,omitempty" name:"Auditor"`
	// 原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type CustomPropertyList struct {
	// key中文名

	KeyNameCn *string `json:"KeyNameCn,omitempty" name:"KeyNameCn"`
	// key英文名

	KeyNameEn *string `json:"KeyNameEn,omitempty" name:"KeyNameEn"`
	// key码，后端定义标识规则

	KeyCode *string `json:"KeyCode,omitempty" name:"KeyCode"`
	// key的字段列表

	Fields *ProductPropertyFields `json:"Fields,omitempty" name:"Fields"`
	// JSON string 属性值列表
	// json decode后是数组，数组元素的key不定
	// 例如：[{"ChnName":"深圳","EngName":"Shenzhen","AreaCode":"szx","ID":"37","Status":"1","Desc":"test"}]

	Values *string `json:"Values,omitempty" name:"Values"`
	// 是否可被编辑。0是不可以；1是可以

	Modifiable *int64 `json:"Modifiable,omitempty" name:"Modifiable"`
}

type DescribeDownloadQuotasListFileGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务是否已完成（0未完成/1已完成）

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 文件名称，当任务已完成时返回该字段

		Name *string `json:"Name,omitempty" name:"Name"`
		// 下载链接，当任务已完成时返回该字段

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDownloadQuotasListFileGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadQuotasListFileGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLifecycleFlowListRequest struct {
	*tchttp.BaseRequest

	// 要查询的UIN

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 操作者UIN

	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`
	// 偏移量

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *string `json:"Limit,omitempty" name:"Limit"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 操作

	OperateAction *string `json:"OperateAction,omitempty" name:"OperateAction"`
}

func (r *DescribeLifecycleFlowListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLifecycleFlowListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceDetailListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 续费参数信息

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceDetailListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceDetailListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCouponListRequest struct {
	*tchttp.BaseRequest

	// 计费MidasAppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 页码，默认1

	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`
	// 每页数量 ，默认20，不得超过50

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 排序字段 ，createtime：创建时间，updatetime：更新时间，默认createtime

	SortField *string `json:"SortField,omitempty" name:"SortField"`
	// 排序类型 ，asc：顺序，desc：倒序，默认desc

	SortType *string `json:"SortType,omitempty" name:"SortType"`
	// 券名，模糊匹配，空或不传表示不过滤

	ActivityName *string `json:"ActivityName,omitempty" name:"ActivityName"`
	// 发券活动ID/groupid，空或不传表示不过滤

	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`
	// 创建用户，空或不传表示不过滤

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
}

func (r *QueryCouponListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryCouponListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TrendByMonthStatItem struct {
	// 开始月份，如2018-05

	BeginMonth *string `json:"BeginMonth,omitempty" name:"BeginMonth"`
	// 结束月份，如2018-11

	EndMonth *string `json:"EndMonth,omitempty" name:"EndMonth"`
	// 应付金额

	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`
	// 代金券花费

	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
}

type DeleteProductGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDownloadListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载地址列表

		DownloadList *GetPkgDownloadListItem `json:"DownloadList,omitempty" name:"DownloadList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDownloadListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillDownloadListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInvoiceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总和

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 当前页数

		Cur *int64 `json:"Cur,omitempty" name:"Cur"`
		// 总页数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 开票记录列表

		List []*DescribeInvoiceList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInvoiceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInvoiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConditionActionType struct {
	// 交易类型

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
	// 交易类型名称

	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`
}

type ProductPriceList struct {
	// 产品标签

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品标签

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项标签

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费项名称

	BillingItemChnName *string `json:"BillingItemChnName,omitempty" name:"BillingItemChnName"`
	// 计费细项标签

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 计费细项名称

	SubBillingItemChnName *string `json:"SubBillingItemChnName,omitempty" name:"SubBillingItemChnName"`
	// priceID,价格策略id

	PriceID *string `json:"PriceID,omitempty" name:"PriceID"`
	// 付费类型

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 计费周期

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 计费模式

	BillingType *string `json:"BillingType,omitempty" name:"BillingType"`
	// 定价类型

	PriceType *string `json:"PriceType,omitempty" name:"PriceType"`
	// 地域ID

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	// 地域代码

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
	// 价格,单位元

	Price *string `json:"Price,omitempty" name:"Price"`
	// 最近操作人

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// JSON string；是自定义属性keyCode和value组成。
	// 例如：{"system":"windows","brand":"tencent"}

	CustomPropertyValues *string `json:"CustomPropertyValues,omitempty" name:"CustomPropertyValues"`
	// 产品项英文名称

	BillingItemEngName *string `json:"BillingItemEngName,omitempty" name:"BillingItemEngName"`
	// 产品细项英文名称

	SubBillingItemEngName *string `json:"SubBillingItemEngName,omitempty" name:"SubBillingItemEngName"`
}

type DescribeBillFeeTrendGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据是否准备好，0未准备好，1准备好

		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`
		// 每月花费详情

		Detail *MonthCostDetailItem `json:"Detail,omitempty" name:"Detail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillFeeTrendGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillFeeTrendGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveCouponConfigRequest struct {
	*tchttp.BaseRequest

	// 计费MidasAppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 活动名称

	ActivityName *string `json:"ActivityName,omitempty" name:"ActivityName"`
	// 发券开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 发券结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 券类型，card: 卡券，balance: 余额券【固定传balance】

	CouponType *string `json:"CouponType,omitempty" name:"CouponType"`
	// 有效期类型，float: 浮动有效期【传TimeValue+TimeUnit】，fixed:固定有效期【传ExpiryTimeBegin+ExpiryTimeEnd】

	ExpiryTimeType *string `json:"ExpiryTimeType,omitempty" name:"ExpiryTimeType"`
	// 有效期时间值，跟ExpiryFloatUnit组合为有效期

	TimeValue *int64 `json:"TimeValue,omitempty" name:"TimeValue"`
	// 有效期时间单位，D:天，H:小时

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 固定有效期开始时间

	ExpiryTimeBegin *string `json:"ExpiryTimeBegin,omitempty" name:"ExpiryTimeBegin"`
	// 固定有效期结束时间）

	ExpiryTimeEnd *string `json:"ExpiryTimeEnd,omitempty" name:"ExpiryTimeEnd"`
	// 门槛（分）【固定传0】

	CouponThreshold *int64 `json:"CouponThreshold,omitempty" name:"CouponThreshold"`
	// 券面值（分）

	CouponValue *int64 `json:"CouponValue,omitempty" name:"CouponValue"`
	// 付费模式 0-预付费，1-后付费，0,1-全部

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 单用户领取次数限制，-1表示不限，{'DayLimit' => -1, 'WeekLimit' => -1, 'MonthLimit' => -1, 'AllLimit' => -1}

	UserLimit *UserLimit `json:"UserLimit,omitempty" name:"UserLimit"`
	// 券总发放次数限制，-1表示不限，{'DayLimit' => -1, 'WeekLimit' => -1, 'MonthLimit' => -1, 'AllLimit' => -1}

	ActivityLimit *ActivityLimit `json:"ActivityLimit,omitempty" name:"ActivityLimit"`
	// json格式四层定义

	ProductInfo *string `json:"ProductInfo,omitempty" name:"ProductInfo"`
	// 发布环境，sandbox:沙箱，actual:现网，空则不发布【固定传actual】

	PublishEnv *string `json:"PublishEnv,omitempty" name:"PublishEnv"`
}

func (r *SaveCouponConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveCouponConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品列表

		List []*Product `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateResourceGatewayRequest struct {
	*tchttp.BaseRequest

	// 目标用户

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

func (r *IsolateResourceGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolateResourceGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWaterInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 流水明细

		WaterDetail *WaterDetail `json:"WaterDetail,omitempty" name:"WaterDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWaterInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWaterInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CouponDetail struct {
	// 计费MidasAppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 发券开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 发券结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 券类型，card: 卡券，balance: 余额券

	CouponType *string `json:"CouponType,omitempty" name:"CouponType"`
	// 活动号

	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`
	// 状态  1为正常，<=0为下线状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 活动名称

	ActivityName *string `json:"ActivityName,omitempty" name:"ActivityName"`
	// 有效期类型

	ExpiryTimeType *string `json:"ExpiryTimeType,omitempty" name:"ExpiryTimeType"`
	// 有效时间值

	TimeValue *int64 `json:"TimeValue,omitempty" name:"TimeValue"`
	// 有效时间单位，D天，H小时

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 固定有效时间开始

	ExpiryTimeBegin *string `json:"ExpiryTimeBegin,omitempty" name:"ExpiryTimeBegin"`
	// 固定有效时间结束

	ExpiryTimeEnd *string `json:"ExpiryTimeEnd,omitempty" name:"ExpiryTimeEnd"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最后更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 券门槛（分）

	CouponThreshold *int64 `json:"CouponThreshold,omitempty" name:"CouponThreshold"`
	// 券面值（分）

	CouponValue *int64 `json:"CouponValue,omitempty" name:"CouponValue"`
	// 付费模式 0-预付费，1-后付费，0,1-全部

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 四层产品定义

	ProductInfo *string `json:"ProductInfo,omitempty" name:"ProductInfo"`
	// 创建用户

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// 用户领取限制

	UserLimit *UserLimit `json:"UserLimit,omitempty" name:"UserLimit"`
	// 发放总限制

	ActivityLimit *ActivityLimit `json:"ActivityLimit,omitempty" name:"ActivityLimit"`
}

type DetailDetailItem struct {
	// 资源所有者uin

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 支付者uin

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 操作者uin

	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`
	// 付费模式

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 付费模式名称

	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 资源名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品名称

	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
	// 子产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 子产品名称

	SubProductCodeName *string `json:"SubProductCodeName,omitempty" name:"SubProductCodeName"`
	// 组件类型

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 组件类型名称

	BillingItemCodeName *string `json:"BillingItemCodeName,omitempty" name:"BillingItemCodeName"`
	// 组件码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 组件名称

	SubBillingItemCodeName *string `json:"SubBillingItemCodeName,omitempty" name:"SubBillingItemCodeName"`
	// 交易类型

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
	// 交易类型名称

	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`
	// 地域ID

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域类型

	RegionType *string `json:"RegionType,omitempty" name:"RegionType"`
	// 地域类型名称

	RegionTypeName *string `json:"RegionTypeName,omitempty" name:"RegionTypeName"`
	// 可用区ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可用区名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 支付时间

	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`
	// 开始使用时间

	FeeBeginTime *string `json:"FeeBeginTime,omitempty" name:"FeeBeginTime"`
	// 结束使用时间

	FeeEndTime *string `json:"FeeEndTime,omitempty" name:"FeeEndTime"`
	// 时间长度

	TimeSpan *string `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 时间单位

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 时间单位名称

	TimeUnitName *string `json:"TimeUnitName,omitempty" name:"TimeUnitName"`
	// 订单号

	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
	// 交易流水号

	BillId *string `json:"BillId,omitempty" name:"BillId"`
	// 组件单价

	SinglePrice *string `json:"SinglePrice,omitempty" name:"SinglePrice"`
	// 组件价格单位

	SinglePriceUnit *string `json:"SinglePriceUnit,omitempty" name:"SinglePriceUnit"`
	// 组件用量

	UsedAmount *string `json:"UsedAmount,omitempty" name:"UsedAmount"`
	// 组件用量单位

	UsedAmountUnit *string `json:"UsedAmountUnit,omitempty" name:"UsedAmountUnit"`
	// 组件折扣

	Discount *string `json:"Discount,omitempty" name:"Discount"`
	// 资源折前总价

	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
	// 组件应付金额

	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`
	// 组件代金券支付金额

	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
	// 账户名称

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 资源总价（折后）

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 资源总折扣

	TotalDiscount *string `json:"TotalDiscount,omitempty" name:"TotalDiscount"`
	// 账号ID

	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

type Product struct {
	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品中文名称

	ChnName *string `json:"ChnName,omitempty" name:"ChnName"`
	// 产品英文名称

	EngName *string `json:"EngName,omitempty" name:"EngName"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type QueryBSPSubProductGatewayRequest struct {
	*tchttp.BaseRequest

	// 商品代码或名称

	SearchProduct *string `json:"SearchProduct,omitempty" name:"SearchProduct"`
	// 子商品代码或名称

	SearchSubProduct *string `json:"SearchSubProduct,omitempty" name:"SearchSubProduct"`
	// 计费项代码或名称

	SearchBillingItem *string `json:"SearchBillingItem,omitempty" name:"SearchBillingItem"`
	// 计费细项代码或名称

	SearchSubBillingItem *string `json:"SearchSubBillingItem,omitempty" name:"SearchSubBillingItem"`
	// 四层定义状态，1-生效 0-失效 （不传默认生效

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 商品代码

	SearchProductCode *string `json:"SearchProductCode,omitempty" name:"SearchProductCode"`
	// 子商品代码

	SearchSubProductCode *string `json:"SearchSubProductCode,omitempty" name:"SearchSubProductCode"`
	// 计费项代码

	SearchBillingItemCode *string `json:"SearchBillingItemCode,omitempty" name:"SearchBillingItemCode"`
	// 计费细项代码

	SearchSubBillingItemCode *string `json:"SearchSubBillingItemCode,omitempty" name:"SearchSubBillingItemCode"`
	// 运营模式：bsp-计量、billing 计费（不传默认计费

	Source *string `json:"Source,omitempty" name:"Source"`
}

func (r *QueryBSPSubProductGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBSPSubProductGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBSPSubBillingItemGatewayRequest struct {
	*tchttp.BaseRequest

	// 商品代码或名称

	SearchProduct *string `json:"SearchProduct,omitempty" name:"SearchProduct"`
	// 子商品代码或名称

	SearchSubProduct *string `json:"SearchSubProduct,omitempty" name:"SearchSubProduct"`
	// 计费项代码或名称

	SearchBillingItem *string `json:"SearchBillingItem,omitempty" name:"SearchBillingItem"`
	// 计费细项代码或名称

	SearchSubBillingItem *string `json:"SearchSubBillingItem,omitempty" name:"SearchSubBillingItem"`
	// 四层定义状态，1-生效 0-失效 （不传默认生效

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 商品代码

	SearchProductCode *string `json:"SearchProductCode,omitempty" name:"SearchProductCode"`
	// 子商品代码

	SearchSubProductCode *string `json:"SearchSubProductCode,omitempty" name:"SearchSubProductCode"`
	// 计费项代码

	SearchBillingItemCode *string `json:"SearchBillingItemCode,omitempty" name:"SearchBillingItemCode"`
	// 计费细项代码

	SearchSubBillingItemCode *string `json:"SearchSubBillingItemCode,omitempty" name:"SearchSubBillingItemCode"`
	// 运营模式：bsp-计量、billing 计费（不传默认计费

	Source *string `json:"Source,omitempty" name:"Source"`
}

func (r *QueryBSPSubBillingItemGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryBSPSubBillingItemGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceListRequest struct {
	*tchttp.BaseRequest

	// 管理端查询用户，只支持单个查询

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 支付模式

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
}

func (r *DescribeResourceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAllProductTreeGatewayRequest struct {
	*tchttp.BaseRequest

	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 页索引

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 页大小

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 是否生效.0未生效 1生效

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *GetAllProductTreeGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAllProductTreeGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyContractTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyContractTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyContractTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDealListGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 大订单列表

		List []*BigDealListData `json:"List,omitempty" name:"List"`
		// 服务器当前时间戳

		Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDealListGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDealListGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePayOrderRequest struct {
	*tchttp.BaseRequest

	// 米大师appid，由计平分配

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// type=by_order，根据订单号查订单
	// type=by_user，根据用户id查订单

	Type *string `json:"Type,omitempty" name:"Type"`
	// 查询的起始时间，unix时间戳,精确到秒。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询的结束时间，unix时间戳,精确到秒。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 记录数偏移量，默认从0开始。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 用户id

	UserNum *string `json:"UserNum,omitempty" name:"UserNum"`
	// 支付订单号

	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`
	// 支付渠道

	Channel *string `json:"Channel,omitempty" name:"Channel"`
	// 提现标记.0-可提现, 1-已提现

	RefundFlag *string `json:"RefundFlag,omitempty" name:"RefundFlag"`
	// 订单状态，用于过滤。本场景固定传["2","3","4"]

	OrderStateList []*string `json:"OrderStateList,omitempty" name:"OrderStateList"`
	// 每页返回的记录数。根据用户号码查询订单列表时需要传。用于分页展示。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePayOrderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePayOrderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDownloadListConditionValue struct {
	// 文件类型过滤条件

	FileType *ConditionFileType `json:"FileType,omitempty" name:"FileType"`
	// 文件状态过滤条件

	Status *ConditionFileStatus `json:"Status,omitempty" name:"Status"`
}

type SubBillingItem struct {
	// 计费细项代码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 计费细项中文名

	ChnName *string `json:"ChnName,omitempty" name:"ChnName"`
	// 计费细项英文名

	EngName *string `json:"EngName,omitempty" name:"EngName"`
	// 计费细项单位

	GoodNumUnit *string `json:"GoodNumUnit,omitempty" name:"GoodNumUnit"`
	// 计费细项英文单位

	GoodNumUnitEn *string `json:"GoodNumUnitEn,omitempty" name:"GoodNumUnitEn"`
	// 计费细项状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeBillTrendByMonthGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *uint64 `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 只支持TimeRange

	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeBillTrendByMonthGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillTrendByMonthGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveQuotaGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveQuotaGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveQuotaGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyRequest struct {
	*tchttp.BaseRequest

	// 目标用户uin

	TargetUin *string `json:"TargetUin,omitempty" name:"TargetUin"`
	// 资源销毁参数信息

	Resources *string `json:"Resources,omitempty" name:"Resources"`
}

func (r *DestroyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DestroyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProductResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CouponInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 代金券信息

		CouponsList []*CouponInfo `json:"CouponsList,omitempty" name:"CouponsList"`
		// 记录总数

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CouponInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CouponInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PayModeSummaryDetailItem struct {
	// 计费模式

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 付费模式名称

	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`
	// 实际花费

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 应付金额

	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`
	// 代金券花费

	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
	// 对比上月趋势

	Trend *SummaryTrend `json:"Trend,omitempty" name:"Trend"`
	// 商品详情

	Product []*BusinessSummaryDetailItem `json:"Product,omitempty" name:"Product"`
}

type ModifyPublicAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPublicAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPublicAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetProductPriceRequest struct {
	*tchttp.BaseRequest

	// 产品标签

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品标签

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项标签

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项标签

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// priceID,价格策略id,新增的时候不传,其他时候用服务器返回值

	PriceID *string `json:"PriceID,omitempty" name:"PriceID"`
	// 付费类型,0:后付费,1:预付费,2:预留实例

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 计费周期,1:年,2:月,3:日,4:小时,5:分钟,6:秒,7:无

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 计费模式,3:按量计费 (后付费才存在)

	BillingType *string `json:"BillingType,omitempty" name:"BillingType"`
	// 定价类型,2:线性价格

	PriceType *string `json:"PriceType,omitempty" name:"PriceType"`
	// 地域ID,例如1

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	// 地域代码,例如gz

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
	// 价格，单位元

	Price *string `json:"Price,omitempty" name:"Price"`
	// 米大师AppId

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// JSON string；是自定义属性key和value组成。
	// 例如：{"system":"windows","brand":"tencent"}

	CustomPropertyData *string `json:"CustomPropertyData,omitempty" name:"CustomPropertyData"`
}

func (r *SetProductPriceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetProductPriceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailTotal struct {
	// 实际花费

	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 应付金额

	PayableAmount *string `json:"PayableAmount,omitempty" name:"PayableAmount"`
	// 代金券花费

	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
	// 折扣金额

	TaxAmount *string `json:"TaxAmount,omitempty" name:"TaxAmount"`
}

type ProductPropertyFields struct {
	// 字段名

	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`
	// 字段码

	FieldCode *string `json:"FieldCode,omitempty" name:"FieldCode"`
	// 字段标志： 0-普通 1-主标记

	FieldFlag *string `json:"FieldFlag,omitempty" name:"FieldFlag"`
}

type DescribeRenewInfoRequest struct {
	*tchttp.BaseRequest

	// 由计平统一分配

	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`
	// 用户id

	UserNum *string `json:"UserNum,omitempty" name:"UserNum"`
	// 0:后付费,1:预付费。不传查询全部

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 自动续费模式：
	// 1：自动续费;
	// 0:手动续费,到期后一定间隔时间后回收;
	// 2:到期关闭

	AutopayType *uint64 `json:"AutopayType,omitempty" name:"AutopayType"`
	// 状态。0：正常；1：销毁；2：隔离;不传查询全部

	Status *string `json:"Status,omitempty" name:"Status"`
	// 地域

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 资源id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 指定排序字段.取值:begin_time、end_time

	SortField *string `json:"SortField,omitempty" name:"SortField"`
	// 指定升序降序：desc、asc

	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
	// 游标

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 记录数目,最小值为1,最大值为10.超过范围则以最值代替

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 固定填:query_renew_info

	Type *string `json:"Type,omitempty" name:"Type"`
	// 开始时间

	EndTimeBegin *string `json:"EndTimeBegin,omitempty" name:"EndTimeBegin"`
	// 结束时间

	EndTimeEnd *string `json:"EndTimeEnd,omitempty" name:"EndTimeEnd"`
}

func (r *DescribeRenewInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRenewInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadDosagesRequest struct {
	*tchttp.BaseRequest

	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 地域id

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *DownloadDosagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadDosagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserStrategyGatewayRequest struct {
	*tchttp.BaseRequest

	// 查询指定租户策略

	ControlUin *string `json:"ControlUin,omitempty" name:"ControlUin"`
	// 固定“bsp”

	Scene *string `json:"Scene,omitempty" name:"Scene"`
	// 页码，从0开始

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数据量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 需要的过滤scene

	Filter *string `json:"Filter,omitempty" name:"Filter"`
}

func (r *GetUserStrategyGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserStrategyGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetComprehensiveProductStrategyGatewayRequest struct {
	*tchttp.BaseRequest

	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 是否只查询第三方自定义产品策略,如果该值不传或者为false则查询全部产品策略

	IfCustom *bool `json:"IfCustom,omitempty" name:"IfCustom"`
	// 状态，如果该值不传则默认查询所有状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 本次查询条数，如果该值不传则默认100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 跳过条数如果该值不传则默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetComprehensiveProductStrategyGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetComprehensiveProductStrategyGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContractTypeFlowListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 流水列表

		List []*ContractTypeFlowData `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContractTypeFlowListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContractTypeFlowListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGoodsPriceGatewayRequest struct {
	*tchttp.BaseRequest

	// 商品询价参数

	ResInfo []*Goods `json:"ResInfo,omitempty" name:"ResInfo"`
}

func (r *DescribeGoodsPriceGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGoodsPriceGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceBillDetailRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 每次获取数据量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 是否需要记录总数，0不需要，1需要

	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
	// 是否需要过滤条件，0需要，1不需要

	NeedConditionValue *uint64 `json:"NeedConditionValue,omitempty" name:"NeedConditionValue"`
	// 只支持BusinessCodes，ProductCodes，ComponentCodes，PayMode, ProjectIds, RegionIds, PayModes, ActionTypes, BillIds, HideFreeCost, ResourceKeyword

	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeResourceBillDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceBillDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetQuotaLeftResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 剩余配额数组

		List []*QuotaLeft `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetQuotaLeftResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetQuotaLeftResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryHiddenProductRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryHiddenProductRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryHiddenProductRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportDiscountGatewayRequest struct {
	*tchttp.BaseRequest

	// 折扣类型 0用户折扣 1官网折扣

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 开始生效时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 失效时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 折扣创建者

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 折扣状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子商品码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 开始时间大于该值

	PreFBeginTime *string `json:"PreFBeginTime,omitempty" name:"PreFBeginTime"`
	// 开始时间小于该值

	TailFBeginTime *string `json:"TailFBeginTime,omitempty" name:"TailFBeginTime"`
	// 结束时间大于该值

	PreFEndTime *string `json:"PreFEndTime,omitempty" name:"PreFEndTime"`
	// 结束时间小于该值

	TailFEndTime *string `json:"TailFEndTime,omitempty" name:"TailFEndTime"`
	// 创建时间大于该值

	PreFCreateTime *string `json:"PreFCreateTime,omitempty" name:"PreFCreateTime"`
	// 创建时间小于该值

	TailFCreateTime *string `json:"TailFCreateTime,omitempty" name:"TailFCreateTime"`
	// 更新时间大于该值

	PreFUpdateTime *string `json:"PreFUpdateTime,omitempty" name:"PreFUpdateTime"`
	// 更新时间小于该值

	TailFUpdateTime *string `json:"TailFUpdateTime,omitempty" name:"TailFUpdateTime"`
	// 用户Uin

	TargetUserId *uint64 `json:"TargetUserId,omitempty" name:"TargetUserId"`
	// 导出Excel需要的表头

	Fields []*string `json:"Fields,omitempty" name:"Fields"`
}

func (r *ExportDiscountGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportDiscountGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetQuotaGatewayRequest struct {
	*tchttp.BaseRequest

	// 商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 商品配额key，由四层定义某几层以及"#"号拼接而成，如p_cvm#sp_cvm_vself2#v_cvm_cpu#sv_cvm_cpu_vself2

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
	// AccountUin

	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`
}

func (r *GetQuotaGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetQuotaGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetQuotaLeftRequest struct {
	*tchttp.BaseRequest

	// 一层商品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 账户uin

	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`
	// 商品配额key，由四层定义某几层以及"#"号拼接而成，如p_cvm#sp_cvm_vself2#v_cvm_cpu#sv_cvm_cpu_vself2

	QuotaKey *string `json:"QuotaKey,omitempty" name:"QuotaKey"`
}

func (r *GetQuotaLeftRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetQuotaLeftRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GoodsDetailList struct {
	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品code

	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`
	// 必传，购买该资源时询价所用的父pid

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 截止时间

	CurDeadline *string `json:"CurDeadline,omitempty" name:"CurDeadline"`
	// 产品详情页面展示参数

	ProductInfo []*ProductInfoList `json:"ProductInfo,omitempty" name:"ProductInfo"`
	// 产品类型

	MediumType *string `json:"MediumType,omitempty" name:"MediumType"`
	// 资源名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// cbs用量大小

	CbsSize *uint64 `json:"CbsSize,omitempty" name:"CbsSize"`
}

type Paras struct {
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 别名对应的值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 三层定义

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 四层定义

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
}

type GetBillingProductCodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品列表

		ProductList *EasyProduct `json:"ProductList,omitempty" name:"ProductList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBillingProductCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBillingProductCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveCouponConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		// 代金券ID

		ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveCouponConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveCouponConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAllProductGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 一层

		List []*ProductCodeAndName `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAllProductGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAllProductGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SummaryTrend struct {
	// upward上升/downward下降/none无

	Type *string `json:"Type,omitempty" name:"Type"`
	// 趋势白分值，保留两位小数

	Value *string `json:"Value,omitempty" name:"Value"`
	// 上月消耗

	PreMonthRealTotalCost *string `json:"PreMonthRealTotalCost,omitempty" name:"PreMonthRealTotalCost"`
}

type DescribeDownloadWithAuthResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDownloadWithAuthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadWithAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBillingProductCodeGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品列表

		ProductList *EasyProduct `json:"ProductList,omitempty" name:"ProductList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBillingProductCodeGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBillingProductCodeGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetProductStatusGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetProductStatusGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetProductStatusGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddProductGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddProductGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddProductGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateBSPProductDisplayCodeGatewayRequest struct {
	*tchttp.BaseRequest

	// 当前页面产品定义代码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 当前页面产品定义名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 产品展示编码

	DisplayCode *string `json:"DisplayCode,omitempty" name:"DisplayCode"`
	// 计费细项代码产品标识名称名称

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 产品定义层级（第几层产品定义

	CodeLevel *int64 `json:"CodeLevel,omitempty" name:"CodeLevel"`
}

func (r *UpdateBSPProductDisplayCodeGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateBSPProductDisplayCodeGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户状态，0为开户或销户，1为正常，其他异常
		// （可根据账户状态来判断是否开户）

		Status *string `json:"Status,omitempty" name:"Status"`
		// 余额，后付费账户，包含现金余额和后付费余额

		Balance *string `json:"Balance,omitempty" name:"Balance"`
		// 后付费额度上限

		Dplimit *string `json:"Dplimit,omitempty" name:"Dplimit"`
		// 后付费剩余可用额度

		Dpbalance *string `json:"Dpbalance,omitempty" name:"Dpbalance"`
		// 后付费账户，现金/溢出余额

		Cash *string `json:"Cash,omitempty" name:"Cash"`
		// 出账日，1-28的数字，每个月第几天出账单

		BillingDate *string `json:"BillingDate,omitempty" name:"BillingDate"`
		// 还款日，1-28的数字，每个月第几天还款

		DueDate *string `json:"DueDate,omitempty" name:"DueDate"`
		// 商户类型，1：代理，2：子客，3：直客

		MerchantType *string `json:"MerchantType,omitempty" name:"MerchantType"`
		// 关联商户信息，当商户类型为代理或直客时，该字段为空；当为子客时，标识代理账户信息

		RelatedMerchant *string `json:"RelatedMerchant,omitempty" name:"RelatedMerchant"`
		// 账期类型，1：按月，2：按天

		DueDelayType *string `json:"DueDelayType,omitempty" name:"DueDelayType"`
		// 账期，还款允许延期时间，当账期类型按月时，表示月份数，按天时，表示天数

		DueDelay *string `json:"DueDelay,omitempty" name:"DueDelay"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAccountInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubQuotaChangeList struct {
	// 主账号uin

	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 子账号uin

	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`
	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 计费项代码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 计费细项代码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 变更前数据值

	SubQuotaChangeAfterData *string `json:"SubQuotaChangeAfterData,omitempty" name:"SubQuotaChangeAfterData"`
	// 变更后数据值

	SubQuotaChangeBeforeData *string `json:"SubQuotaChangeBeforeData,omitempty" name:"SubQuotaChangeBeforeData"`
	// 内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 创建人

	Creator *string `json:"Creator,omitempty" name:"Creator"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeResourceDetailListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录总数

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 续费参数信息

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceDetailListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceDetailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadDosagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载地址

		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadDosagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadDosagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadQuotasListFileGatewayRequest struct {
	*tchttp.BaseRequest

	// 用户uin

	QuotaUin *string `json:"QuotaUin,omitempty" name:"QuotaUin"`
	// 产品码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 添加任务者，用于内部TCB邮件通知

	InnerOperator *string `json:"InnerOperator,omitempty" name:"InnerOperator"`
}

func (r *DescribeDownloadQuotasListFileGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDownloadQuotasListFileGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPublicAccountList struct {
	// 和入参ID保持一致（默认只有1）

	AccountId *uint64 `json:"AccountId,omitempty" name:"AccountId"`
	// 收款人户名

	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
	// 收款卡号

	AccountNum *uint64 `json:"AccountNum,omitempty" name:"AccountNum"`
	// 银行

	Bank *string `json:"Bank,omitempty" name:"Bank"`
}

type SaveUserDiscountGatewayRequest struct {
	*tchttp.BaseRequest

	// 折扣参数

	DiscountPara []*DiscountPara `json:"DiscountPara,omitempty" name:"DiscountPara"`
}

func (r *SaveUserDiscountGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SaveUserDiscountGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReconciliationListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 每次获取数据量，与请求包一致

		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
		// 偏移量，与请求包一致

		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
		// 记录数，若调用时传的needRecordNum=0，则不返回该字段

		RecordNum *string `json:"RecordNum,omitempty" name:"RecordNum"`
		// 调账列表详情

		Data []*ReconciliationDataItem `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReconciliationListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReconciliationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SuperCoupon struct {
	// 券信息id

	CodeId *string `json:"CodeId,omitempty" name:"CodeId"`
	// 批次id

	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
	// 券有效开始时间，unix时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
	// 券有效截至时间，unix时间

	EndTime []*string `json:"EndTime,omitempty" name:"EndTime"`
	// 发券结果

	ResultCode *string `json:"ResultCode,omitempty" name:"ResultCode"`
	// 发券结果信息

	ResultInfo *string `json:"ResultInfo,omitempty" name:"ResultInfo"`
}

type ListResourceDetailGatewayResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源详情信息

		List *BspResourceBill `json:"List,omitempty" name:"List"`
		// 数据条数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListResourceDetailGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceDetailGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddComprehensiveStrategyGatewayRequest struct {
	*tchttp.BaseRequest

	// 产品code

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品code

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 付费模式

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 时间周期 如 1,2,4,6

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 结算周期 如 3,4

	CalcUnit *string `json:"CalcUnit,omitempty" name:"CalcUnit"`
	// 交易类型 如 1,2,3,5

	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
	// 计费模式 如1,2,3,5

	BillingType *string `json:"BillingType,omitempty" name:"BillingType"`
	// 新购是否同步发货 默认0

	PurchaseSyncDelivery *uint64 `json:"PurchaseSyncDelivery,omitempty" name:"PurchaseSyncDelivery"`
	// 变配是否同步发货 默认0

	ModifySyncDelivery *uint64 `json:"ModifySyncDelivery,omitempty" name:"ModifySyncDelivery"`
	// 资源回收类型 默认1

	ResourceRecycleType *uint64 `json:"ResourceRecycleType,omitempty" name:"ResourceRecycleType"`
	// 接口名

	InterfaceNames *InterfaceNames `json:"InterfaceNames,omitempty" name:"InterfaceNames"`
	// url名

	Region2interface *string `json:"Region2interface,omitempty" name:"Region2interface"`
	// 定价类型,2:线性价格

	PriceType *uint64 `json:"PriceType,omitempty" name:"PriceType"`
}

func (r *AddComprehensiveStrategyGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddComprehensiveStrategyGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReconciliationListGatewayRequest struct {
	*tchttp.BaseRequest

	// 每次获取数据量

	Limit *string `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *string `json:"Offset,omitempty" name:"Offset"`
	// 调账记录查询调账列表部分参数

	Conditions *ReconciliationConditions `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *DescribeReconciliationListGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReconciliationListGatewayRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListResourceModifyLogRequest struct {
	*tchttp.BaseRequest

	// 数据所属用户的UIN

	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 产品编码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 产品细项编码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 页码

	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 当前页条数

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListResourceModifyLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResourceModifyLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
