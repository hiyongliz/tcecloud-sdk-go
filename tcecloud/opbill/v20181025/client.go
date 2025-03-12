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
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2018-10-25"

var _ = tchttp.POST

type Client struct {
	common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
	cpf := profile.NewClientProfile()
	client = &Client{}
	client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
	return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewDescribeTenantSubUinQuotasRequest() (request *DescribeTenantSubUinQuotasRequest) {
	request = &DescribeTenantSubUinQuotasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeTenantSubUinQuotas")
	return
}

func NewDescribeTenantSubUinQuotasResponse() (response *DescribeTenantSubUinQuotasResponse) {
	response = &DescribeTenantSubUinQuotasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端-子账号登录后，查询当前账户在主账户下的配额信息
func (c *Client) DescribeTenantSubUinQuotas(request *DescribeTenantSubUinQuotasRequest) (response *DescribeTenantSubUinQuotasResponse, err error) {
	if request == nil {
		request = NewDescribeTenantSubUinQuotasRequest()
	}
	response = NewDescribeTenantSubUinQuotasResponse()
	err = c.Send(request, response)
	return
}

func NewSetProductStatusGatewayRequest() (request *SetProductStatusGatewayRequest) {
	request = &SetProductStatusGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SetProductStatusGateway")
	return
}

func NewSetProductStatusGatewayResponse() (response *SetProductStatusGatewayResponse) {
	response = &SetProductStatusGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 变更产品定义状态
func (c *Client) SetProductStatusGateway(request *SetProductStatusGatewayRequest) (response *SetProductStatusGatewayResponse, err error) {
	if request == nil {
		request = NewSetProductStatusGatewayRequest()
	}
	response = NewSetProductStatusGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetComprehensiveSubProductStrategyRouteGatewayRequest() (request *GetComprehensiveSubProductStrategyRouteGatewayRequest) {
	request = &GetComprehensiveSubProductStrategyRouteGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetComprehensiveSubProductStrategyRouteGateway")
	return
}

func NewGetComprehensiveSubProductStrategyRouteGatewayResponse() (response *GetComprehensiveSubProductStrategyRouteGatewayResponse) {
	response = &GetComprehensiveSubProductStrategyRouteGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询子产品（二层）策略接口
func (c *Client) GetComprehensiveSubProductStrategyRouteGateway(request *GetComprehensiveSubProductStrategyRouteGatewayRequest) (response *GetComprehensiveSubProductStrategyRouteGatewayResponse, err error) {
	if request == nil {
		request = NewGetComprehensiveSubProductStrategyRouteGatewayRequest()
	}
	response = NewGetComprehensiveSubProductStrategyRouteGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillTrendByMonthGatewayRequest() (request *DescribeBillTrendByMonthGatewayRequest) {
	request = &DescribeBillTrendByMonthGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillTrendByMonthGateway")
	return
}

func NewDescribeBillTrendByMonthGatewayResponse() (response *DescribeBillTrendByMonthGatewayResponse) {
	response = &DescribeBillTrendByMonthGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账单按月消费趋势
func (c *Client) DescribeBillTrendByMonthGateway(request *DescribeBillTrendByMonthGatewayRequest) (response *DescribeBillTrendByMonthGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillTrendByMonthGatewayRequest()
	}
	response = NewDescribeBillTrendByMonthGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewAddProductRequest() (request *AddProductRequest) {
	request = &AddProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "AddProduct")
	return
}

func NewAddProductResponse() (response *AddProductResponse) {
	response = &AddProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加产品四层定义
func (c *Client) AddProduct(request *AddProductRequest) (response *AddProductResponse, err error) {
	if request == nil {
		request = NewAddProductRequest()
	}
	response = NewAddProductResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBasicAccountListRequest() (request *DescribeBasicAccountListRequest) {
	request = &DescribeBasicAccountListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBasicAccountList")
	return
}

func NewDescribeBasicAccountListResponse() (response *DescribeBasicAccountListResponse) {
	response = &DescribeBasicAccountListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据uin批量获取账户基本信息
func (c *Client) DescribeBasicAccountList(request *DescribeBasicAccountListRequest) (response *DescribeBasicAccountListResponse, err error) {
	if request == nil {
		request = NewDescribeBasicAccountListRequest()
	}
	response = NewDescribeBasicAccountListResponse()
	err = c.Send(request, response)
	return
}

func NewCouponInfoRequest() (request *CouponInfoRequest) {
	request = &CouponInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "CouponInfo")
	return
}

func NewCouponInfoResponse() (response *CouponInfoResponse) {
	response = &CouponInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 代金券查询接口
func (c *Client) CouponInfo(request *CouponInfoRequest) (response *CouponInfoResponse, err error) {
	if request == nil {
		request = NewCouponInfoRequest()
	}
	response = NewCouponInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyContractTypeGatewayRequest() (request *ModifyContractTypeGatewayRequest) {
	request = &ModifyContractTypeGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ModifyContractTypeGateway")
	return
}

func NewModifyContractTypeGatewayResponse() (response *ModifyContractTypeGatewayResponse) {
	response = &ModifyContractTypeGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改合同类型
func (c *Client) ModifyContractTypeGateway(request *ModifyContractTypeGatewayRequest) (response *ModifyContractTypeGatewayResponse, err error) {
	if request == nil {
		request = NewModifyContractTypeGatewayRequest()
	}
	response = NewModifyContractTypeGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCouponPresentRequest() (request *CouponPresentRequest) {
	request = &CouponPresentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "CouponPresent")
	return
}

func NewCouponPresentResponse() (response *CouponPresentResponse) {
	response = &CouponPresentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发放抵扣券
func (c *Client) CouponPresent(request *CouponPresentRequest) (response *CouponPresentResponse, err error) {
	if request == nil {
		request = NewCouponPresentRequest()
	}
	response = NewCouponPresentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadByResourceRequest() (request *DescribeBillDownloadByResourceRequest) {
	request = &DescribeBillDownloadByResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillDownloadByResource")
	return
}

func NewDescribeBillDownloadByResourceResponse() (response *DescribeBillDownloadByResourceResponse) {
	response = &DescribeBillDownloadByResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载定制化账单按资源汇总
func (c *Client) DescribeBillDownloadByResource(request *DescribeBillDownloadByResourceRequest) (response *DescribeBillDownloadByResourceResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadByResourceRequest()
	}
	response = NewDescribeBillDownloadByResourceResponse()
	err = c.Send(request, response)
	return
}

func NewGetComprehensiveSubProdctStrategyGatewayRequest() (request *GetComprehensiveSubProdctStrategyGatewayRequest) {
	request = &GetComprehensiveSubProdctStrategyGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetComprehensiveSubProdctStrategyGateway")
	return
}

func NewGetComprehensiveSubProdctStrategyGatewayResponse() (response *GetComprehensiveSubProdctStrategyGatewayResponse) {
	response = &GetComprehensiveSubProdctStrategyGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询子产品（二层）策略接口
func (c *Client) GetComprehensiveSubProdctStrategyGateway(request *GetComprehensiveSubProdctStrategyGatewayRequest) (response *GetComprehensiveSubProdctStrategyGatewayResponse, err error) {
	if request == nil {
		request = NewGetComprehensiveSubProdctStrategyGatewayRequest()
	}
	response = NewGetComprehensiveSubProdctStrategyGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetAllProductTreeRequest() (request *GetAllProductTreeRequest) {
	request = &GetAllProductTreeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetAllProductTree")
	return
}

func NewGetAllProductTreeResponse() (response *GetAllProductTreeResponse) {
	response = &GetAllProductTreeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过商品码，子商品码，计费项，计费细项来查询产品定义，参数都可选，如果不传则返回全部数据。不会隐藏产品
func (c *Client) GetAllProductTree(request *GetAllProductTreeRequest) (response *GetAllProductTreeResponse, err error) {
	if request == nil {
		request = NewGetAllProductTreeRequest()
	}
	response = NewGetAllProductTreeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByPayModeGatewayRequest() (request *DescribeBillSummaryByPayModeGatewayRequest) {
	request = &DescribeBillSummaryByPayModeGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillSummaryByPayModeGateway")
	return
}

func NewDescribeBillSummaryByPayModeGatewayResponse() (response *DescribeBillSummaryByPayModeGatewayResponse) {
	response = &DescribeBillSummaryByPayModeGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取计费模式汇总费用分布
func (c *Client) DescribeBillSummaryByPayModeGateway(request *DescribeBillSummaryByPayModeGatewayRequest) (response *DescribeBillSummaryByPayModeGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByPayModeGatewayRequest()
	}
	response = NewDescribeBillSummaryByPayModeGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByProductGatewayRequest() (request *DescribeBillSummaryByProductGatewayRequest) {
	request = &DescribeBillSummaryByProductGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillSummaryByProductGateway")
	return
}

func NewDescribeBillSummaryByProductGatewayResponse() (response *DescribeBillSummaryByProductGatewayResponse) {
	response = &DescribeBillSummaryByProductGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品汇总费用分布
func (c *Client) DescribeBillSummaryByProductGateway(request *DescribeBillSummaryByProductGatewayRequest) (response *DescribeBillSummaryByProductGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByProductGatewayRequest()
	}
	response = NewDescribeBillSummaryByProductGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDestroyGatewayRequest() (request *DestroyGatewayRequest) {
	request = &DestroyGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DestroyGateway")
	return
}

func NewDestroyGatewayResponse() (response *DestroyGatewayResponse) {
	response = &DestroyGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预付费资源销毁
func (c *Client) DestroyGateway(request *DestroyGatewayRequest) (response *DestroyGatewayResponse, err error) {
	if request == nil {
		request = NewDestroyGatewayRequest()
	}
	response = NewDestroyGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyComprehensiveStrategyRequest() (request *ModifyComprehensiveStrategyRequest) {
	request = &ModifyComprehensiveStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ModifyComprehensiveStrategy")
	return
}

func NewModifyComprehensiveStrategyResponse() (response *ModifyComprehensiveStrategyResponse) {
	response = &ModifyComprehensiveStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改策略包含第三方自定义和预设
func (c *Client) ModifyComprehensiveStrategy(request *ModifyComprehensiveStrategyRequest) (response *ModifyComprehensiveStrategyResponse, err error) {
	if request == nil {
		request = NewModifyComprehensiveStrategyRequest()
	}
	response = NewModifyComprehensiveStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContractTypeFlowListGatewayRequest() (request *DescribeContractTypeFlowListGatewayRequest) {
	request = &DescribeContractTypeFlowListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeContractTypeFlowListGateway")
	return
}

func NewDescribeContractTypeFlowListGatewayResponse() (response *DescribeContractTypeFlowListGatewayResponse) {
	response = &DescribeContractTypeFlowListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 合同类型变更流水列表
func (c *Client) DescribeContractTypeFlowListGateway(request *DescribeContractTypeFlowListGatewayRequest) (response *DescribeContractTypeFlowListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeContractTypeFlowListGatewayRequest()
	}
	response = NewDescribeContractTypeFlowListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetBillingProductCodeRequest() (request *GetBillingProductCodeRequest) {
	request = &GetBillingProductCodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetBillingProductCode")
	return
}

func NewGetBillingProductCodeResponse() (response *GetBillingProductCodeResponse) {
	response = &GetBillingProductCodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询由计量迁移至计费的产品码
func (c *Client) GetBillingProductCode(request *GetBillingProductCodeRequest) (response *GetBillingProductCodeResponse, err error) {
	if request == nil {
		request = NewGetBillingProductCodeRequest()
	}
	response = NewGetBillingProductCodeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInvoiceListRequest() (request *DescribeInvoiceListRequest) {
	request = &DescribeInvoiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeInvoiceList")
	return
}

func NewDescribeInvoiceListResponse() (response *DescribeInvoiceListResponse) {
	response = &DescribeInvoiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端、租户端-获取用户开票记录
func (c *Client) DescribeInvoiceList(request *DescribeInvoiceListRequest) (response *DescribeInvoiceListResponse, err error) {
	if request == nil {
		request = NewDescribeInvoiceListRequest()
	}
	response = NewDescribeInvoiceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceBillDownloadRequest() (request *DescribeResourceBillDownloadRequest) {
	request = &DescribeResourceBillDownloadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeResourceBillDownload")
	return
}

func NewDescribeResourceBillDownloadResponse() (response *DescribeResourceBillDownloadResponse) {
	response = &DescribeResourceBillDownloadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用量明细导出
func (c *Client) DescribeResourceBillDownload(request *DescribeResourceBillDownloadRequest) (response *DescribeResourceBillDownloadResponse, err error) {
	if request == nil {
		request = NewDescribeResourceBillDownloadRequest()
	}
	response = NewDescribeResourceBillDownloadResponse()
	err = c.Send(request, response)
	return
}

func NewGetComprehensiveStrategyWebConfigRequest() (request *GetComprehensiveStrategyWebConfigRequest) {
	request = &GetComprehensiveStrategyWebConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetComprehensiveStrategyWebConfig")
	return
}

func NewGetComprehensiveStrategyWebConfigResponse() (response *GetComprehensiveStrategyWebConfigResponse) {
	response = &GetComprehensiveStrategyWebConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有枚举配置
func (c *Client) GetComprehensiveStrategyWebConfig(request *GetComprehensiveStrategyWebConfigRequest) (response *GetComprehensiveStrategyWebConfigResponse, err error) {
	if request == nil {
		request = NewGetComprehensiveStrategyWebConfigRequest()
	}
	response = NewGetComprehensiveStrategyWebConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetAllProductRequest() (request *GetAllProductRequest) {
	request = &GetAllProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetAllProduct")
	return
}

func NewGetAllProductResponse() (response *GetAllProductResponse) {
	response = &GetAllProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取全量产品一层（产品模块）
func (c *Client) GetAllProduct(request *GetAllProductRequest) (response *GetAllProductResponse, err error) {
	if request == nil {
		request = NewGetAllProductRequest()
	}
	response = NewGetAllProductResponse()
	err = c.Send(request, response)
	return
}

func NewModifyContractStatusRequest() (request *ModifyContractStatusRequest) {
	request = &ModifyContractStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ModifyContractStatus")
	return
}

func NewModifyContractStatusResponse() (response *ModifyContractStatusResponse) {
	response = &ModifyContractStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改合同状态
func (c *Client) ModifyContractStatus(request *ModifyContractStatusRequest) (response *ModifyContractStatusResponse, err error) {
	if request == nil {
		request = NewModifyContractStatusRequest()
	}
	response = NewModifyContractStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadListRequest() (request *DescribeBillDownloadListRequest) {
	request = &DescribeBillDownloadListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillDownloadList")
	return
}

func NewDescribeBillDownloadListResponse() (response *DescribeBillDownloadListResponse) {
	response = &DescribeBillDownloadListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账单下载地址列表
func (c *Client) DescribeBillDownloadList(request *DescribeBillDownloadListRequest) (response *DescribeBillDownloadListResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadListRequest()
	}
	response = NewDescribeBillDownloadListResponse()
	err = c.Send(request, response)
	return
}

func NewSetProductPriceRequest() (request *SetProductPriceRequest) {
	request = &SetProductPriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SetProductPrice")
	return
}

func NewSetProductPriceResponse() (response *SetProductPriceResponse) {
	response = &SetProductPriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增、编辑价格策略
func (c *Client) SetProductPrice(request *SetProductPriceRequest) (response *SetProductPriceResponse, err error) {
	if request == nil {
		request = NewSetProductPriceRequest()
	}
	response = NewSetProductPriceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubProductsGatewayRequest() (request *DescribeSubProductsGatewayRequest) {
	request = &DescribeSubProductsGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeSubProductsGateway")
	return
}

func NewDescribeSubProductsGatewayResponse() (response *DescribeSubProductsGatewayResponse) {
	response = &DescribeSubProductsGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取子商品信息
func (c *Client) DescribeSubProductsGateway(request *DescribeSubProductsGatewayRequest) (response *DescribeSubProductsGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeSubProductsGatewayRequest()
	}
	response = NewDescribeSubProductsGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceBillDetailGatewayRequest() (request *DescribeResourceBillDetailGatewayRequest) {
	request = &DescribeResourceBillDetailGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeResourceBillDetailGateway")
	return
}

func NewDescribeResourceBillDetailGatewayResponse() (response *DescribeResourceBillDetailGatewayResponse) {
	response = &DescribeResourceBillDetailGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组件级别明细账单
func (c *Client) DescribeResourceBillDetailGateway(request *DescribeResourceBillDetailGatewayRequest) (response *DescribeResourceBillDetailGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeResourceBillDetailGatewayRequest()
	}
	response = NewDescribeResourceBillDetailGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadResourceDetailGatewayRequest() (request *DescribeBillDownloadResourceDetailGatewayRequest) {
	request = &DescribeBillDownloadResourceDetailGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillDownloadResourceDetailGateway")
	return
}

func NewDescribeBillDownloadResourceDetailGatewayResponse() (response *DescribeBillDownloadResourceDetailGatewayResponse) {
	response = &DescribeBillDownloadResourceDetailGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载定制化账单明细
func (c *Client) DescribeBillDownloadResourceDetailGateway(request *DescribeBillDownloadResourceDetailGatewayRequest) (response *DescribeBillDownloadResourceDetailGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadResourceDetailGatewayRequest()
	}
	response = NewDescribeBillDownloadResourceDetailGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContractListGatewayRequest() (request *DescribeContractListGatewayRequest) {
	request = &DescribeContractListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeContractListGateway")
	return
}

func NewDescribeContractListGatewayResponse() (response *DescribeContractListGatewayResponse) {
	response = &DescribeContractListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询合同列表接口
func (c *Client) DescribeContractListGateway(request *DescribeContractListGatewayRequest) (response *DescribeContractListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeContractListGatewayRequest()
	}
	response = NewDescribeContractListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDosagesGatewayRequest() (request *DescribeDosagesGatewayRequest) {
	request = &DescribeDosagesGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeDosagesGateway")
	return
}

func NewDescribeDosagesGatewayResponse() (response *DescribeDosagesGatewayResponse) {
	response = &DescribeDosagesGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 推量统计
func (c *Client) DescribeDosagesGateway(request *DescribeDosagesGatewayRequest) (response *DescribeDosagesGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeDosagesGatewayRequest()
	}
	response = NewDescribeDosagesGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCreateContractTypeRequest() (request *CreateContractTypeRequest) {
	request = &CreateContractTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "CreateContractType")
	return
}

func NewCreateContractTypeResponse() (response *CreateContractTypeResponse) {
	response = &CreateContractTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建合同类型
func (c *Client) CreateContractType(request *CreateContractTypeRequest) (response *CreateContractTypeResponse, err error) {
	if request == nil {
		request = NewCreateContractTypeRequest()
	}
	response = NewCreateContractTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadVoucherListFileGatewayRequest() (request *DescribeDownloadVoucherListFileGatewayRequest) {
	request = &DescribeDownloadVoucherListFileGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeDownloadVoucherListFileGateway")
	return
}

func NewDescribeDownloadVoucherListFileGatewayResponse() (response *DescribeDownloadVoucherListFileGatewayResponse) {
	response = &DescribeDownloadVoucherListFileGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载代金券信息excel
func (c *Client) DescribeDownloadVoucherListFileGateway(request *DescribeDownloadVoucherListFileGatewayRequest) (response *DescribeDownloadVoucherListFileGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadVoucherListFileGatewayRequest()
	}
	response = NewDescribeDownloadVoucherListFileGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetComprehensiveProductStrategyRequest() (request *GetComprehensiveProductStrategyRequest) {
	request = &GetComprehensiveProductStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetComprehensiveProductStrategy")
	return
}

func NewGetComprehensiveProductStrategyResponse() (response *GetComprehensiveProductStrategyResponse) {
	response = &GetComprehensiveProductStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品（一层）策略接口
func (c *Client) GetComprehensiveProductStrategy(request *GetComprehensiveProductStrategyRequest) (response *GetComprehensiveProductStrategyResponse, err error) {
	if request == nil {
		request = NewGetComprehensiveProductStrategyRequest()
	}
	response = NewGetComprehensiveProductStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewReversalResourceRequest() (request *ReversalResourceRequest) {
	request = &ReversalResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ReversalResource")
	return
}

func NewReversalResourceResponse() (response *ReversalResourceResponse) {
	response = &ReversalResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端-资源解隔离接口
func (c *Client) ReversalResource(request *ReversalResourceRequest) (response *ReversalResourceResponse, err error) {
	if request == nil {
		request = NewReversalResourceRequest()
	}
	response = NewReversalResourceResponse()
	err = c.Send(request, response)
	return
}

func NewSetProductInfoGatewayRequest() (request *SetProductInfoGatewayRequest) {
	request = &SetProductInfoGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SetProductInfoGateway")
	return
}

func NewSetProductInfoGatewayResponse() (response *SetProductInfoGatewayResponse) {
	response = &SetProductInfoGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品四层定义
func (c *Client) SetProductInfoGateway(request *SetProductInfoGatewayRequest) (response *SetProductInfoGatewayResponse, err error) {
	if request == nil {
		request = NewSetProductInfoGatewayRequest()
	}
	response = NewSetProductInfoGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadWithAuthGatewayRequest() (request *DescribeDownloadWithAuthGatewayRequest) {
	request = &DescribeDownloadWithAuthGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeDownloadWithAuthGateway")
	return
}

func NewDescribeDownloadWithAuthGatewayResponse() (response *DescribeDownloadWithAuthGatewayResponse) {
	response = &DescribeDownloadWithAuthGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据任务ID查询文件下载地址
func (c *Client) DescribeDownloadWithAuthGateway(request *DescribeDownloadWithAuthGatewayRequest) (response *DescribeDownloadWithAuthGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadWithAuthGatewayRequest()
	}
	response = NewDescribeDownloadWithAuthGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDiscountStatusRequest() (request *ModifyDiscountStatusRequest) {
	request = &ModifyDiscountStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ModifyDiscountStatus")
	return
}

func NewModifyDiscountStatusResponse() (response *ModifyDiscountStatusResponse) {
	response = &ModifyDiscountStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改折扣状态
func (c *Client) ModifyDiscountStatus(request *ModifyDiscountStatusRequest) (response *ModifyDiscountStatusResponse, err error) {
	if request == nil {
		request = NewModifyDiscountStatusRequest()
	}
	response = NewModifyDiscountStatusResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceSummaryDownloadGatewayRequest() (request *ListResourceSummaryDownloadGatewayRequest) {
	request = &ListResourceSummaryDownloadGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ListResourceSummaryDownloadGateway")
	return
}

func NewListResourceSummaryDownloadGatewayResponse() (response *ListResourceSummaryDownloadGatewayResponse) {
	response = &ListResourceSummaryDownloadGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资源详情
func (c *Client) ListResourceSummaryDownloadGateway(request *ListResourceSummaryDownloadGatewayRequest) (response *ListResourceSummaryDownloadGatewayResponse, err error) {
	if request == nil {
		request = NewListResourceSummaryDownloadGatewayRequest()
	}
	response = NewListResourceSummaryDownloadGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewIsolateResourceRequest() (request *IsolateResourceRequest) {
	request = &IsolateResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "IsolateResource")
	return
}

func NewIsolateResourceResponse() (response *IsolateResourceResponse) {
	response = &IsolateResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端资源隔离接口
func (c *Client) IsolateResource(request *IsolateResourceRequest) (response *IsolateResourceResponse, err error) {
	if request == nil {
		request = NewIsolateResourceRequest()
	}
	response = NewIsolateResourceResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceGatewayRequest() (request *ListResourceGatewayRequest) {
	request = &ListResourceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ListResourceGateway")
	return
}

func NewListResourceGatewayResponse() (response *ListResourceGatewayResponse) {
	response = &ListResourceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端根据Uin，ResourceId查询用户资源
func (c *Client) ListResourceGateway(request *ListResourceGatewayRequest) (response *ListResourceGatewayResponse, err error) {
	if request == nil {
		request = NewListResourceGatewayRequest()
	}
	response = NewListResourceGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceModifyLogRequest() (request *ListResourceModifyLogRequest) {
	request = &ListResourceModifyLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ListResourceModifyLog")
	return
}

func NewListResourceModifyLogResponse() (response *ListResourceModifyLogResponse) {
	response = &ListResourceModifyLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资源变配记录
func (c *Client) ListResourceModifyLog(request *ListResourceModifyLogRequest) (response *ListResourceModifyLogResponse, err error) {
	if request == nil {
		request = NewListResourceModifyLogRequest()
	}
	response = NewListResourceModifyLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQuotaTypeListRequest() (request *DescribeQuotaTypeListRequest) {
	request = &DescribeQuotaTypeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeQuotaTypeList")
	return
}

func NewDescribeQuotaTypeListResponse() (response *DescribeQuotaTypeListResponse) {
	response = &DescribeQuotaTypeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询配额类别：uin、sys
func (c *Client) DescribeQuotaTypeList(request *DescribeQuotaTypeListRequest) (response *DescribeQuotaTypeListResponse, err error) {
	if request == nil {
		request = NewDescribeQuotaTypeListRequest()
	}
	response = NewDescribeQuotaTypeListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemitRecordRequest() (request *DescribeRemitRecordRequest) {
	request = &DescribeRemitRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeRemitRecord")
	return
}

func NewDescribeRemitRecordResponse() (response *DescribeRemitRecordResponse) {
	response = &DescribeRemitRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端查询汇款记录列表
func (c *Client) DescribeRemitRecord(request *DescribeRemitRecordRequest) (response *DescribeRemitRecordResponse, err error) {
	if request == nil {
		request = NewDescribeRemitRecordRequest()
	}
	response = NewDescribeRemitRecordResponse()
	err = c.Send(request, response)
	return
}

func NewExportDealListOpGatewayRequest() (request *ExportDealListOpGatewayRequest) {
	request = &ExportDealListOpGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ExportDealListOpGateway")
	return
}

func NewExportDealListOpGatewayResponse() (response *ExportDealListOpGatewayResponse) {
	response = &ExportDealListOpGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端导入订单列表信息
func (c *Client) ExportDealListOpGateway(request *ExportDealListOpGatewayRequest) (response *ExportDealListOpGatewayResponse, err error) {
	if request == nil {
		request = NewExportDealListOpGatewayRequest()
	}
	response = NewExportDealListOpGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewSaveQuotaGatewayRequest() (request *SaveQuotaGatewayRequest) {
	request = &SaveQuotaGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SaveQuotaGateway")
	return
}

func NewSaveQuotaGatewayResponse() (response *SaveQuotaGatewayResponse) {
	response = &SaveQuotaGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品配额新增、编辑
func (c *Client) SaveQuotaGateway(request *SaveQuotaGatewayRequest) (response *SaveQuotaGatewayResponse, err error) {
	if request == nil {
		request = NewSaveQuotaGatewayRequest()
	}
	response = NewSaveQuotaGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDownloadRecordGatewayRequest() (request *CreateDownloadRecordGatewayRequest) {
	request = &CreateDownloadRecordGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "CreateDownloadRecordGateway")
	return
}

func NewCreateDownloadRecordGatewayResponse() (response *CreateDownloadRecordGatewayResponse) {
	response = &CreateDownloadRecordGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户点击下载时通知后台
func (c *Client) CreateDownloadRecordGateway(request *CreateDownloadRecordGatewayRequest) (response *CreateDownloadRecordGatewayResponse, err error) {
	if request == nil {
		request = NewCreateDownloadRecordGatewayRequest()
	}
	response = NewCreateDownloadRecordGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDealPriceRequest() (request *DescribeDealPriceRequest) {
	request = &DescribeDealPriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeDealPrice")
	return
}

func NewDescribeDealPriceResponse() (response *DescribeDealPriceResponse) {
	response = &DescribeDealPriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询订单价格
func (c *Client) DescribeDealPrice(request *DescribeDealPriceRequest) (response *DescribeDealPriceResponse, err error) {
	if request == nil {
		request = NewDescribeDealPriceRequest()
	}
	response = NewDescribeDealPriceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductRequest() (request *DeleteProductRequest) {
	request = &DeleteProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DeleteProduct")
	return
}

func NewDeleteProductResponse() (response *DeleteProductResponse) {
	response = &DeleteProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【计量】删除产品四层定义
func (c *Client) DeleteProduct(request *DeleteProductRequest) (response *DeleteProductResponse, err error) {
	if request == nil {
		request = NewDeleteProductRequest()
	}
	response = NewDeleteProductResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubBillingItemsRequest() (request *DescribeSubBillingItemsRequest) {
	request = &DescribeSubBillingItemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeSubBillingItems")
	return
}

func NewDescribeSubBillingItemsResponse() (response *DescribeSubBillingItemsResponse) {
	response = &DescribeSubBillingItemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取计费细项信息（第四层）
func (c *Client) DescribeSubBillingItems(request *DescribeSubBillingItemsRequest) (response *DescribeSubBillingItemsResponse, err error) {
	if request == nil {
		request = NewDescribeSubBillingItemsRequest()
	}
	response = NewDescribeSubBillingItemsResponse()
	err = c.Send(request, response)
	return
}

func NewGetTagsByMonthGatewayRequest() (request *GetTagsByMonthGatewayRequest) {
	request = &GetTagsByMonthGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetTagsByMonthGateway")
	return
}

func NewGetTagsByMonthGatewayResponse() (response *GetTagsByMonthGatewayResponse) {
	response = &GetTagsByMonthGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某月设置的标签键列表
func (c *Client) GetTagsByMonthGateway(request *GetTagsByMonthGatewayRequest) (response *GetTagsByMonthGatewayResponse, err error) {
	if request == nil {
		request = NewGetTagsByMonthGatewayRequest()
	}
	response = NewGetTagsByMonthGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomPropertiesRequest() (request *DescribeCustomPropertiesRequest) {
	request = &DescribeCustomPropertiesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeCustomProperties")
	return
}

func NewDescribeCustomPropertiesResponse() (response *DescribeCustomPropertiesResponse) {
	response = &DescribeCustomPropertiesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 自定义售卖属性列表查询
func (c *Client) DescribeCustomProperties(request *DescribeCustomPropertiesRequest) (response *DescribeCustomPropertiesResponse, err error) {
	if request == nil {
		request = NewDescribeCustomPropertiesRequest()
	}
	response = NewDescribeCustomPropertiesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQuotaDownloadRequest() (request *DescribeQuotaDownloadRequest) {
	request = &DescribeQuotaDownloadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeQuotaDownload")
	return
}

func NewDescribeQuotaDownloadResponse() (response *DescribeQuotaDownloadResponse) {
	response = &DescribeQuotaDownloadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配额导出
func (c *Client) DescribeQuotaDownload(request *DescribeQuotaDownloadRequest) (response *DescribeQuotaDownloadResponse, err error) {
	if request == nil {
		request = NewDescribeQuotaDownloadRequest()
	}
	response = NewDescribeQuotaDownloadResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeTrendRequest() (request *DescribeBillFeeTrendRequest) {
	request = &DescribeBillFeeTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillFeeTrend")
	return
}

func NewDescribeBillFeeTrendResponse() (response *DescribeBillFeeTrendResponse) {
	response = &DescribeBillFeeTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取各类汇总明细费用趋势
func (c *Client) DescribeBillFeeTrend(request *DescribeBillFeeTrendRequest) (response *DescribeBillFeeTrendResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeTrendRequest()
	}
	response = NewDescribeBillFeeTrendResponse()
	err = c.Send(request, response)
	return
}

func NewQueryBSPSubProductGatewayRequest() (request *QueryBSPSubProductGatewayRequest) {
	request = &QueryBSPSubProductGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "QueryBSPSubProductGateway")
	return
}

func NewQueryBSPSubProductGatewayResponse() (response *QueryBSPSubProductGatewayResponse) {
	response = &QueryBSPSubProductGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询二层产品定义及展示名称、产品标识
func (c *Client) QueryBSPSubProductGateway(request *QueryBSPSubProductGatewayRequest) (response *QueryBSPSubProductGatewayResponse, err error) {
	if request == nil {
		request = NewQueryBSPSubProductGatewayRequest()
	}
	response = NewQueryBSPSubProductGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLifecycleFlowListGatewayRequest() (request *DescribeLifecycleFlowListGatewayRequest) {
	request = &DescribeLifecycleFlowListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeLifecycleFlowListGateway")
	return
}

func NewDescribeLifecycleFlowListGatewayResponse() (response *DescribeLifecycleFlowListGatewayResponse) {
	response = &DescribeLifecycleFlowListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资源操作流水
func (c *Client) DescribeLifecycleFlowListGateway(request *DescribeLifecycleFlowListGatewayRequest) (response *DescribeLifecycleFlowListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeLifecycleFlowListGatewayRequest()
	}
	response = NewDescribeLifecycleFlowListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewExportChangeQuotaGatewayRequest() (request *ExportChangeQuotaGatewayRequest) {
	request = &ExportChangeQuotaGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ExportChangeQuotaGateway")
	return
}

func NewExportChangeQuotaGatewayResponse() (response *ExportChangeQuotaGatewayResponse) {
	response = &ExportChangeQuotaGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端下载配额变更记录
func (c *Client) ExportChangeQuotaGateway(request *ExportChangeQuotaGatewayRequest) (response *ExportChangeQuotaGatewayResponse, err error) {
	if request == nil {
		request = NewExportChangeQuotaGatewayRequest()
	}
	response = NewExportChangeQuotaGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetComprehensiveProductStrategyGatewayRequest() (request *GetComprehensiveProductStrategyGatewayRequest) {
	request = &GetComprehensiveProductStrategyGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetComprehensiveProductStrategyGateway")
	return
}

func NewGetComprehensiveProductStrategyGatewayResponse() (response *GetComprehensiveProductStrategyGatewayResponse) {
	response = &GetComprehensiveProductStrategyGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品（一层）策略接口
func (c *Client) GetComprehensiveProductStrategyGateway(request *GetComprehensiveProductStrategyGatewayRequest) (response *GetComprehensiveProductStrategyGatewayResponse, err error) {
	if request == nil {
		request = NewGetComprehensiveProductStrategyGatewayRequest()
	}
	response = NewGetComprehensiveProductStrategyGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCouponListRequest() (request *QueryCouponListRequest) {
	request = &QueryCouponListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "QueryCouponList")
	return
}

func NewQueryCouponListResponse() (response *QueryCouponListResponse) {
	response = &QueryCouponListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询券列表.
func (c *Client) QueryCouponList(request *QueryCouponListRequest) (response *QueryCouponListResponse, err error) {
	if request == nil {
		request = NewQueryCouponListRequest()
	}
	response = NewQueryCouponListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceBillDownloadGatewayRequest() (request *DescribeResourceBillDownloadGatewayRequest) {
	request = &DescribeResourceBillDownloadGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeResourceBillDownloadGateway")
	return
}

func NewDescribeResourceBillDownloadGatewayResponse() (response *DescribeResourceBillDownloadGatewayResponse) {
	response = &DescribeResourceBillDownloadGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用量明细导出
func (c *Client) DescribeResourceBillDownloadGateway(request *DescribeResourceBillDownloadGatewayRequest) (response *DescribeResourceBillDownloadGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeResourceBillDownloadGatewayRequest()
	}
	response = NewDescribeResourceBillDownloadGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewExportDiscountGatewayRequest() (request *ExportDiscountGatewayRequest) {
	request = &ExportDiscountGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ExportDiscountGateway")
	return
}

func NewExportDiscountGatewayResponse() (response *ExportDiscountGatewayResponse) {
	response = &ExportDiscountGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出折扣列表
func (c *Client) ExportDiscountGateway(request *ExportDiscountGatewayRequest) (response *ExportDiscountGatewayResponse, err error) {
	if request == nil {
		request = NewExportDiscountGatewayRequest()
	}
	response = NewExportDiscountGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewAddComprehensiveStrategyGatewayRequest() (request *AddComprehensiveStrategyGatewayRequest) {
	request = &AddComprehensiveStrategyGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "AddComprehensiveStrategyGateway")
	return
}

func NewAddComprehensiveStrategyGatewayResponse() (response *AddComprehensiveStrategyGatewayResponse) {
	response = &AddComprehensiveStrategyGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 可以新增第三方和预设的产品策略
func (c *Client) AddComprehensiveStrategyGateway(request *AddComprehensiveStrategyGatewayRequest) (response *AddComprehensiveStrategyGatewayResponse, err error) {
	if request == nil {
		request = NewAddComprehensiveStrategyGatewayRequest()
	}
	response = NewAddComprehensiveStrategyGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGoodsPriceRequest() (request *DescribeGoodsPriceRequest) {
	request = &DescribeGoodsPriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeGoodsPrice")
	return
}

func NewDescribeGoodsPriceResponse() (response *DescribeGoodsPriceResponse) {
	response = &DescribeGoodsPriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询商品价格
func (c *Client) DescribeGoodsPrice(request *DescribeGoodsPriceRequest) (response *DescribeGoodsPriceResponse, err error) {
	if request == nil {
		request = NewDescribeGoodsPriceRequest()
	}
	response = NewDescribeGoodsPriceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLifecycleFlowListRequest() (request *DescribeLifecycleFlowListRequest) {
	request = &DescribeLifecycleFlowListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeLifecycleFlowList")
	return
}

func NewDescribeLifecycleFlowListResponse() (response *DescribeLifecycleFlowListResponse) {
	response = &DescribeLifecycleFlowListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资源操作流水
func (c *Client) DescribeLifecycleFlowList(request *DescribeLifecycleFlowListRequest) (response *DescribeLifecycleFlowListResponse, err error) {
	if request == nil {
		request = NewDescribeLifecycleFlowListRequest()
	}
	response = NewDescribeLifecycleFlowListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTenantSubUinQuotasGatewayRequest() (request *DescribeTenantSubUinQuotasGatewayRequest) {
	request = &DescribeTenantSubUinQuotasGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeTenantSubUinQuotasGateway")
	return
}

func NewDescribeTenantSubUinQuotasGatewayResponse() (response *DescribeTenantSubUinQuotasGatewayResponse) {
	response = &DescribeTenantSubUinQuotasGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端-子账号登录后，查询当前账户在主账户下的配额信息
func (c *Client) DescribeTenantSubUinQuotasGateway(request *DescribeTenantSubUinQuotasGatewayRequest) (response *DescribeTenantSubUinQuotasGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeTenantSubUinQuotasGatewayRequest()
	}
	response = NewDescribeTenantSubUinQuotasGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsolidatedBillDownloadUrlByUinsGatewayRequest() (request *DescribeConsolidatedBillDownloadUrlByUinsGatewayRequest) {
	request = &DescribeConsolidatedBillDownloadUrlByUinsGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeConsolidatedBillDownloadUrlByUinsGateway")
	return
}

func NewDescribeConsolidatedBillDownloadUrlByUinsGatewayResponse() (response *DescribeConsolidatedBillDownloadUrlByUinsGatewayResponse) {
	response = &DescribeConsolidatedBillDownloadUrlByUinsGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账单包下载地址
func (c *Client) DescribeConsolidatedBillDownloadUrlByUinsGateway(request *DescribeConsolidatedBillDownloadUrlByUinsGatewayRequest) (response *DescribeConsolidatedBillDownloadUrlByUinsGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeConsolidatedBillDownloadUrlByUinsGatewayRequest()
	}
	response = NewDescribeConsolidatedBillDownloadUrlByUinsGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContractTypeFlowListRequest() (request *DescribeContractTypeFlowListRequest) {
	request = &DescribeContractTypeFlowListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeContractTypeFlowList")
	return
}

func NewDescribeContractTypeFlowListResponse() (response *DescribeContractTypeFlowListResponse) {
	response = &DescribeContractTypeFlowListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 合同类型变更流水列表
func (c *Client) DescribeContractTypeFlowList(request *DescribeContractTypeFlowListRequest) (response *DescribeContractTypeFlowListResponse, err error) {
	if request == nil {
		request = NewDescribeContractTypeFlowListRequest()
	}
	response = NewDescribeContractTypeFlowListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDealListGatewayRequest() (request *DescribeDealListGatewayRequest) {
	request = &DescribeDealListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeDealListGateway")
	return
}

func NewDescribeDealListGatewayResponse() (response *DescribeDealListGatewayResponse) {
	response = &DescribeDealListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端查询订单列表
func (c *Client) DescribeDealListGateway(request *DescribeDealListGatewayRequest) (response *DescribeDealListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeDealListGatewayRequest()
	}
	response = NewDescribeDealListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserDebtBillRequest() (request *DescribeUserDebtBillRequest) {
	request = &DescribeUserDebtBillRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeUserDebtBill")
	return
}

func NewDescribeUserDebtBillResponse() (response *DescribeUserDebtBillResponse) {
	response = &DescribeUserDebtBillResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询后付费账单
func (c *Client) DescribeUserDebtBill(request *DescribeUserDebtBillRequest) (response *DescribeUserDebtBillResponse, err error) {
	if request == nil {
		request = NewDescribeUserDebtBillRequest()
	}
	response = NewDescribeUserDebtBillResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReconciliationGatewayRequest() (request *ModifyReconciliationGatewayRequest) {
	request = &ModifyReconciliationGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ModifyReconciliationGateway")
	return
}

func NewModifyReconciliationGatewayResponse() (response *ModifyReconciliationGatewayResponse) {
	response = &ModifyReconciliationGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提交调账
func (c *Client) ModifyReconciliationGateway(request *ModifyReconciliationGatewayRequest) (response *ModifyReconciliationGatewayResponse, err error) {
	if request == nil {
		request = NewModifyReconciliationGatewayRequest()
	}
	response = NewModifyReconciliationGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductsGatewayRequest() (request *DescribeProductsGatewayRequest) {
	request = &DescribeProductsGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeProductsGateway")
	return
}

func NewDescribeProductsGatewayResponse() (response *DescribeProductsGatewayResponse) {
	response = &DescribeProductsGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取商品信息(第一层)
func (c *Client) DescribeProductsGateway(request *DescribeProductsGatewayRequest) (response *DescribeProductsGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeProductsGatewayRequest()
	}
	response = NewDescribeProductsGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadCommonSummaryRequest() (request *DescribeBillDownloadCommonSummaryRequest) {
	request = &DescribeBillDownloadCommonSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillDownloadCommonSummary")
	return
}

func NewDescribeBillDownloadCommonSummaryResponse() (response *DescribeBillDownloadCommonSummaryResponse) {
	response = &DescribeBillDownloadCommonSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载账单通用汇总（产品+地域）
func (c *Client) DescribeBillDownloadCommonSummary(request *DescribeBillDownloadCommonSummaryRequest) (response *DescribeBillDownloadCommonSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadCommonSummaryRequest()
	}
	response = NewDescribeBillDownloadCommonSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBspProductRelationsGatewayRequest() (request *DescribeBspProductRelationsGatewayRequest) {
	request = &DescribeBspProductRelationsGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBspProductRelationsGateway")
	return
}

func NewDescribeBspProductRelationsGatewayResponse() (response *DescribeBspProductRelationsGatewayResponse) {
	response = &DescribeBspProductRelationsGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品定义列表
func (c *Client) DescribeBspProductRelationsGateway(request *DescribeBspProductRelationsGatewayRequest) (response *DescribeBspProductRelationsGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBspProductRelationsGatewayRequest()
	}
	response = NewDescribeBspProductRelationsGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetTagsByMonthRequest() (request *GetTagsByMonthRequest) {
	request = &GetTagsByMonthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetTagsByMonth")
	return
}

func NewGetTagsByMonthResponse() (response *GetTagsByMonthResponse) {
	response = &GetTagsByMonthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某月设置的标签键列表
func (c *Client) GetTagsByMonth(request *GetTagsByMonthRequest) (response *GetTagsByMonthResponse, err error) {
	if request == nil {
		request = NewGetTagsByMonthRequest()
	}
	response = NewGetTagsByMonthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillingItemsRequest() (request *DescribeBillingItemsRequest) {
	request = &DescribeBillingItemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillingItems")
	return
}

func NewDescribeBillingItemsResponse() (response *DescribeBillingItemsResponse) {
	response = &DescribeBillingItemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取计费项信息（第三层）
func (c *Client) DescribeBillingItems(request *DescribeBillingItemsRequest) (response *DescribeBillingItemsResponse, err error) {
	if request == nil {
		request = NewDescribeBillingItemsRequest()
	}
	response = NewDescribeBillingItemsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductDownloadGatewayRequest() (request *DescribeProductDownloadGatewayRequest) {
	request = &DescribeProductDownloadGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeProductDownloadGateway")
	return
}

func NewDescribeProductDownloadGatewayResponse() (response *DescribeProductDownloadGatewayResponse) {
	response = &DescribeProductDownloadGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品定义导出
func (c *Client) DescribeProductDownloadGateway(request *DescribeProductDownloadGatewayRequest) (response *DescribeProductDownloadGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeProductDownloadGatewayRequest()
	}
	response = NewDescribeProductDownloadGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCreateContractTypeGatewayRequest() (request *CreateContractTypeGatewayRequest) {
	request = &CreateContractTypeGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "CreateContractTypeGateway")
	return
}

func NewCreateContractTypeGatewayResponse() (response *CreateContractTypeGatewayResponse) {
	response = &CreateContractTypeGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建合同类型
func (c *Client) CreateContractTypeGateway(request *CreateContractTypeGatewayRequest) (response *CreateContractTypeGatewayResponse, err error) {
	if request == nil {
		request = NewCreateContractTypeGatewayRequest()
	}
	response = NewCreateContractTypeGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInvoiceListGatewayRequest() (request *DescribeInvoiceListGatewayRequest) {
	request = &DescribeInvoiceListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeInvoiceListGateway")
	return
}

func NewDescribeInvoiceListGatewayResponse() (response *DescribeInvoiceListGatewayResponse) {
	response = &DescribeInvoiceListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端、租户端-获取用户开票记录
func (c *Client) DescribeInvoiceListGateway(request *DescribeInvoiceListGatewayRequest) (response *DescribeInvoiceListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeInvoiceListGatewayRequest()
	}
	response = NewDescribeInvoiceListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetQuotaGatewayRequest() (request *GetQuotaGatewayRequest) {
	request = &GetQuotaGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetQuotaGateway")
	return
}

func NewGetQuotaGatewayResponse() (response *GetQuotaGatewayResponse) {
	response = &GetQuotaGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据商品码和配额key查询配额
func (c *Client) GetQuotaGateway(request *GetQuotaGatewayRequest) (response *GetQuotaGatewayResponse, err error) {
	if request == nil {
		request = NewGetQuotaGatewayRequest()
	}
	response = NewGetQuotaGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyComprehensiveStrategyRouteGatewayRequest() (request *ModifyComprehensiveStrategyRouteGatewayRequest) {
	request = &ModifyComprehensiveStrategyRouteGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ModifyComprehensiveStrategyRouteGateway")
	return
}

func NewModifyComprehensiveStrategyRouteGatewayResponse() (response *ModifyComprehensiveStrategyRouteGatewayResponse) {
	response = &ModifyComprehensiveStrategyRouteGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改策略包含第三方自定义和预设
func (c *Client) ModifyComprehensiveStrategyRouteGateway(request *ModifyComprehensiveStrategyRouteGatewayRequest) (response *ModifyComprehensiveStrategyRouteGatewayResponse, err error) {
	if request == nil {
		request = NewModifyComprehensiveStrategyRouteGatewayRequest()
	}
	response = NewModifyComprehensiveStrategyRouteGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGeneralDataRequest() (request *DescribeGeneralDataRequest) {
	request = &DescribeGeneralDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeGeneralData")
	return
}

func NewDescribeGeneralDataResponse() (response *DescribeGeneralDataResponse) {
	response = &DescribeGeneralDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口用于通用查询
func (c *Client) DescribeGeneralData(request *DescribeGeneralDataRequest) (response *DescribeGeneralDataResponse, err error) {
	if request == nil {
		request = NewDescribeGeneralDataRequest()
	}
	response = NewDescribeGeneralDataResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceDetailGatewayRequest() (request *ListResourceDetailGatewayRequest) {
	request = &ListResourceDetailGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ListResourceDetailGateway")
	return
}

func NewListResourceDetailGatewayResponse() (response *ListResourceDetailGatewayResponse) {
	response = &ListResourceDetailGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资源详情
func (c *Client) ListResourceDetailGateway(request *ListResourceDetailGatewayRequest) (response *ListResourceDetailGatewayResponse, err error) {
	if request == nil {
		request = NewListResourceDetailGatewayRequest()
	}
	response = NewListResourceDetailGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeByTagRequest() (request *DescribeBillFeeByTagRequest) {
	request = &DescribeBillFeeByTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillFeeByTag")
	return
}

func NewDescribeBillFeeByTagResponse() (response *DescribeBillFeeByTagResponse) {
	response = &DescribeBillFeeByTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询按标签汇总明细数据
func (c *Client) DescribeBillFeeByTag(request *DescribeBillFeeByTagRequest) (response *DescribeBillFeeByTagResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeByTagRequest()
	}
	response = NewDescribeBillFeeByTagResponse()
	err = c.Send(request, response)
	return
}

func NewAuditInvoiceInfoRequest() (request *AuditInvoiceInfoRequest) {
	request = &AuditInvoiceInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "AuditInvoiceInfo")
	return
}

func NewAuditInvoiceInfoResponse() (response *AuditInvoiceInfoResponse) {
	response = &AuditInvoiceInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端-审核发票信息
func (c *Client) AuditInvoiceInfo(request *AuditInvoiceInfoRequest) (response *AuditInvoiceInfoResponse, err error) {
	if request == nil {
		request = NewAuditInvoiceInfoRequest()
	}
	response = NewAuditInvoiceInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeByPayModeGatewayRequest() (request *DescribeBillFeeByPayModeGatewayRequest) {
	request = &DescribeBillFeeByPayModeGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillFeeByPayModeGateway")
	return
}

func NewDescribeBillFeeByPayModeGatewayResponse() (response *DescribeBillFeeByPayModeGatewayResponse) {
	response = &DescribeBillFeeByPayModeGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按计费模式汇总明细费用
func (c *Client) DescribeBillFeeByPayModeGateway(request *DescribeBillFeeByPayModeGatewayRequest) (response *DescribeBillFeeByPayModeGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeByPayModeGatewayRequest()
	}
	response = NewDescribeBillFeeByPayModeGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteQuotaGatewayRequest() (request *DeleteQuotaGatewayRequest) {
	request = &DeleteQuotaGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DeleteQuotaGateway")
	return
}

func NewDeleteQuotaGatewayResponse() (response *DeleteQuotaGatewayResponse) {
	response = &DeleteQuotaGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口描述
func (c *Client) DeleteQuotaGateway(request *DeleteQuotaGatewayRequest) (response *DeleteQuotaGatewayResponse, err error) {
	if request == nil {
		request = NewDeleteQuotaGatewayRequest()
	}
	response = NewDeleteQuotaGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceDetailListRequest() (request *DescribeResourceDetailListRequest) {
	request = &DescribeResourceDetailListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeResourceDetailList")
	return
}

func NewDescribeResourceDetailListResponse() (response *DescribeResourceDetailListResponse) {
	response = &DescribeResourceDetailListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户续费参数
func (c *Client) DescribeResourceDetailList(request *DescribeResourceDetailListRequest) (response *DescribeResourceDetailListResponse, err error) {
	if request == nil {
		request = NewDescribeResourceDetailListRequest()
	}
	response = NewDescribeResourceDetailListResponse()
	err = c.Send(request, response)
	return
}

func NewSaveCouponConfigRequest() (request *SaveCouponConfigRequest) {
	request = &SaveCouponConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SaveCouponConfig")
	return
}

func NewSaveCouponConfigResponse() (response *SaveCouponConfigResponse) {
	response = &SaveCouponConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 抵扣券创建接口
func (c *Client) SaveCouponConfig(request *SaveCouponConfigRequest) (response *SaveCouponConfigResponse, err error) {
	if request == nil {
		request = NewSaveCouponConfigRequest()
	}
	response = NewSaveCouponConfigResponse()
	err = c.Send(request, response)
	return
}

func NewSubscribeDestroyRequest() (request *SubscribeDestroyRequest) {
	request = &SubscribeDestroyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SubscribeDestroy")
	return
}

func NewSubscribeDestroyResponse() (response *SubscribeDestroyResponse) {
	response = &SubscribeDestroyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 销毁资源。
func (c *Client) SubscribeDestroy(request *SubscribeDestroyRequest) (response *SubscribeDestroyResponse, err error) {
	if request == nil {
		request = NewSubscribeDestroyRequest()
	}
	response = NewSubscribeDestroyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadByResourceGatewayRequest() (request *DescribeBillDownloadByResourceGatewayRequest) {
	request = &DescribeBillDownloadByResourceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillDownloadByResourceGateway")
	return
}

func NewDescribeBillDownloadByResourceGatewayResponse() (response *DescribeBillDownloadByResourceGatewayResponse) {
	response = &DescribeBillDownloadByResourceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载定制化账单按资源汇总
func (c *Client) DescribeBillDownloadByResourceGateway(request *DescribeBillDownloadByResourceGatewayRequest) (response *DescribeBillDownloadByResourceGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadByResourceGatewayRequest()
	}
	response = NewDescribeBillDownloadByResourceGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeByRegionRequest() (request *DescribeBillFeeByRegionRequest) {
	request = &DescribeBillFeeByRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillFeeByRegion")
	return
}

func NewDescribeBillFeeByRegionResponse() (response *DescribeBillFeeByRegionResponse) {
	response = &DescribeBillFeeByRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按地域汇总明细费用
func (c *Client) DescribeBillFeeByRegion(request *DescribeBillFeeByRegionRequest) (response *DescribeBillFeeByRegionResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeByRegionRequest()
	}
	response = NewDescribeBillFeeByRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductDownloadRequest() (request *DescribeProductDownloadRequest) {
	request = &DescribeProductDownloadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeProductDownload")
	return
}

func NewDescribeProductDownloadResponse() (response *DescribeProductDownloadResponse) {
	response = &DescribeProductDownloadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品定义导出
func (c *Client) DescribeProductDownload(request *DescribeProductDownloadRequest) (response *DescribeProductDownloadResponse, err error) {
	if request == nil {
		request = NewDescribeProductDownloadRequest()
	}
	response = NewDescribeProductDownloadResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubBillingItemsGatewayRequest() (request *DescribeSubBillingItemsGatewayRequest) {
	request = &DescribeSubBillingItemsGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeSubBillingItemsGateway")
	return
}

func NewDescribeSubBillingItemsGatewayResponse() (response *DescribeSubBillingItemsGatewayResponse) {
	response = &DescribeSubBillingItemsGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取计费细项信息（第四层）
func (c *Client) DescribeSubBillingItemsGateway(request *DescribeSubBillingItemsGatewayRequest) (response *DescribeSubBillingItemsGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeSubBillingItemsGatewayRequest()
	}
	response = NewDescribeSubBillingItemsGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewQueryBSPBillingItemGatewayRequest() (request *QueryBSPBillingItemGatewayRequest) {
	request = &QueryBSPBillingItemGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "QueryBSPBillingItemGateway")
	return
}

func NewQueryBSPBillingItemGatewayResponse() (response *QueryBSPBillingItemGatewayResponse) {
	response = &QueryBSPBillingItemGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询三层产品定义及展示名称、产品标识
func (c *Client) QueryBSPBillingItemGateway(request *QueryBSPBillingItemGatewayRequest) (response *QueryBSPBillingItemGatewayResponse, err error) {
	if request == nil {
		request = NewQueryBSPBillingItemGatewayRequest()
	}
	response = NewQueryBSPBillingItemGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCancleContractTypeGatewayRequest() (request *CancleContractTypeGatewayRequest) {
	request = &CancleContractTypeGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "CancleContractTypeGateway")
	return
}

func NewCancleContractTypeGatewayResponse() (response *CancleContractTypeGatewayResponse) {
	response = &CancleContractTypeGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 作废合同类型
func (c *Client) CancleContractTypeGateway(request *CancleContractTypeGatewayRequest) (response *CancleContractTypeGatewayResponse, err error) {
	if request == nil {
		request = NewCancleContractTypeGatewayRequest()
	}
	response = NewCancleContractTypeGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByProductRequest() (request *DescribeBillSummaryByProductRequest) {
	request = &DescribeBillSummaryByProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillSummaryByProduct")
	return
}

func NewDescribeBillSummaryByProductResponse() (response *DescribeBillSummaryByProductResponse) {
	response = &DescribeBillSummaryByProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品汇总费用分布
func (c *Client) DescribeBillSummaryByProduct(request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByProductRequest()
	}
	response = NewDescribeBillSummaryByProductResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountListRequest() (request *DescribeAccountListRequest) {
	request = &DescribeAccountListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeAccountList")
	return
}

func NewDescribeAccountListResponse() (response *DescribeAccountListResponse) {
	response = &DescribeAccountListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端查询帐户列表和用户代金券总金额
func (c *Client) DescribeAccountList(request *DescribeAccountListRequest) (response *DescribeAccountListResponse, err error) {
	if request == nil {
		request = NewDescribeAccountListRequest()
	}
	response = NewDescribeAccountListResponse()
	err = c.Send(request, response)
	return
}

func NewSaveUserDiscountRequest() (request *SaveUserDiscountRequest) {
	request = &SaveUserDiscountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SaveUserDiscount")
	return
}

func NewSaveUserDiscountResponse() (response *SaveUserDiscountResponse) {
	response = &SaveUserDiscountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增用户折扣
func (c *Client) SaveUserDiscount(request *SaveUserDiscountRequest) (response *SaveUserDiscountResponse, err error) {
	if request == nil {
		request = NewSaveUserDiscountRequest()
	}
	response = NewSaveUserDiscountResponse()
	err = c.Send(request, response)
	return
}

func NewGetBillingProductCodeGatewayRequest() (request *GetBillingProductCodeGatewayRequest) {
	request = &GetBillingProductCodeGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetBillingProductCodeGateway")
	return
}

func NewGetBillingProductCodeGatewayResponse() (response *GetBillingProductCodeGatewayResponse) {
	response = &GetBillingProductCodeGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询由计量迁移至计费的产品码
func (c *Client) GetBillingProductCodeGateway(request *GetBillingProductCodeGatewayRequest) (response *GetBillingProductCodeGatewayResponse, err error) {
	if request == nil {
		request = NewGetBillingProductCodeGatewayRequest()
	}
	response = NewGetBillingProductCodeGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewSetBspProductInfoGatewayRequest() (request *SetBspProductInfoGatewayRequest) {
	request = &SetBspProductInfoGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SetBspProductInfoGateway")
	return
}

func NewSetBspProductInfoGatewayResponse() (response *SetBspProductInfoGatewayResponse) {
	response = &SetBspProductInfoGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品四层定义
func (c *Client) SetBspProductInfoGateway(request *SetBspProductInfoGatewayRequest) (response *SetBspProductInfoGatewayResponse, err error) {
	if request == nil {
		request = NewSetBspProductInfoGatewayRequest()
	}
	response = NewSetBspProductInfoGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewExportDealListRequest() (request *ExportDealListRequest) {
	request = &ExportDealListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ExportDealList")
	return
}

func NewExportDealListResponse() (response *ExportDealListResponse) {
	response = &ExportDealListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端导入订单列表信息
func (c *Client) ExportDealList(request *ExportDealListRequest) (response *ExportDealListResponse, err error) {
	if request == nil {
		request = NewExportDealListRequest()
	}
	response = NewExportDealListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQuotasGatewayRequest() (request *DescribeQuotasGatewayRequest) {
	request = &DescribeQuotasGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeQuotasGateway")
	return
}

func NewDescribeQuotasGatewayResponse() (response *DescribeQuotasGatewayResponse) {
	response = &DescribeQuotasGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品配额列表查询
func (c *Client) DescribeQuotasGateway(request *DescribeQuotasGatewayRequest) (response *DescribeQuotasGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeQuotasGatewayRequest()
	}
	response = NewDescribeQuotasGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceListGatewayRequest() (request *DescribeResourceListGatewayRequest) {
	request = &DescribeResourceListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeResourceListGateway")
	return
}

func NewDescribeResourceListGatewayResponse() (response *DescribeResourceListGatewayResponse) {
	response = &DescribeResourceListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端查询用户产品列表，只支持到一层产品
func (c *Client) DescribeResourceListGateway(request *DescribeResourceListGatewayRequest) (response *DescribeResourceListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeResourceListGatewayRequest()
	}
	response = NewDescribeResourceListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDiscountRequest() (request *DeleteDiscountRequest) {
	request = &DeleteDiscountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DeleteDiscount")
	return
}

func NewDeleteDiscountResponse() (response *DeleteDiscountResponse) {
	response = &DeleteDiscountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除折扣
func (c *Client) DeleteDiscount(request *DeleteDiscountRequest) (response *DeleteDiscountResponse, err error) {
	if request == nil {
		request = NewDeleteDiscountRequest()
	}
	response = NewDeleteDiscountResponse()
	err = c.Send(request, response)
	return
}

func NewAuditRemitRecordRequest() (request *AuditRemitRecordRequest) {
	request = &AuditRemitRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "AuditRemitRecord")
	return
}

func NewAuditRemitRecordResponse() (response *AuditRemitRecordResponse) {
	response = &AuditRemitRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 审核汇款记录
func (c *Client) AuditRemitRecord(request *AuditRemitRecordRequest) (response *AuditRemitRecordResponse, err error) {
	if request == nil {
		request = NewAuditRemitRecordRequest()
	}
	response = NewAuditRemitRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductQuotaGatewayRequest() (request *DeleteProductQuotaGatewayRequest) {
	request = &DeleteProductQuotaGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DeleteProductQuotaGateway")
	return
}

func NewDeleteProductQuotaGatewayResponse() (response *DeleteProductQuotaGatewayResponse) {
	response = &DeleteProductQuotaGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除产品配额
func (c *Client) DeleteProductQuotaGateway(request *DeleteProductQuotaGatewayRequest) (response *DeleteProductQuotaGatewayResponse, err error) {
	if request == nil {
		request = NewDeleteProductQuotaGatewayRequest()
	}
	response = NewDeleteProductQuotaGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillingSubUinQuotaGatewayRequest() (request *DescribeBillingSubUinQuotaGatewayRequest) {
	request = &DescribeBillingSubUinQuotaGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillingSubUinQuotaGateway")
	return
}

func NewDescribeBillingSubUinQuotaGatewayResponse() (response *DescribeBillingSubUinQuotaGatewayResponse) {
	response = &DescribeBillingSubUinQuotaGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端通过子账号uin查询配额全量
func (c *Client) DescribeBillingSubUinQuotaGateway(request *DescribeBillingSubUinQuotaGatewayRequest) (response *DescribeBillingSubUinQuotaGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillingSubUinQuotaGatewayRequest()
	}
	response = NewDescribeBillingSubUinQuotaGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInvoiceDetailGatewayRequest() (request *DescribeInvoiceDetailGatewayRequest) {
	request = &DescribeInvoiceDetailGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeInvoiceDetailGateway")
	return
}

func NewDescribeInvoiceDetailGatewayResponse() (response *DescribeInvoiceDetailGatewayResponse) {
	response = &DescribeInvoiceDetailGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端、租户端-获取发票的开票详情
func (c *Client) DescribeInvoiceDetailGateway(request *DescribeInvoiceDetailGatewayRequest) (response *DescribeInvoiceDetailGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeInvoiceDetailGatewayRequest()
	}
	response = NewDescribeInvoiceDetailGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewExportDealListOpRequest() (request *ExportDealListOpRequest) {
	request = &ExportDealListOpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ExportDealListOp")
	return
}

func NewExportDealListOpResponse() (response *ExportDealListOpResponse) {
	response = &ExportDealListOpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端导入订单列表信息
func (c *Client) ExportDealListOp(request *ExportDealListOpRequest) (response *ExportDealListOpResponse, err error) {
	if request == nil {
		request = NewExportDealListOpRequest()
	}
	response = NewExportDealListOpResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductGatewayRequest() (request *DeleteProductGatewayRequest) {
	request = &DeleteProductGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DeleteProductGateway")
	return
}

func NewDeleteProductGatewayResponse() (response *DeleteProductGatewayResponse) {
	response = &DeleteProductGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【计量】删除产品四层定义
func (c *Client) DeleteProductGateway(request *DeleteProductGatewayRequest) (response *DeleteProductGatewayResponse, err error) {
	if request == nil {
		request = NewDeleteProductGatewayRequest()
	}
	response = NewDeleteProductGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewSetHiddenProductRequest() (request *SetHiddenProductRequest) {
	request = &SetHiddenProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SetHiddenProduct")
	return
}

func NewSetHiddenProductResponse() (response *SetHiddenProductResponse) {
	response = &SetHiddenProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置隐藏冗余商品项接口，本接口接收前端请求存储和修改需隐藏的商品项。
func (c *Client) SetHiddenProduct(request *SetHiddenProductRequest) (response *SetHiddenProductResponse, err error) {
	if request == nil {
		request = NewSetHiddenProductRequest()
	}
	response = NewSetHiddenProductResponse()
	err = c.Send(request, response)
	return
}

func NewExportDiscountRequest() (request *ExportDiscountRequest) {
	request = &ExportDiscountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ExportDiscount")
	return
}

func NewExportDiscountResponse() (response *ExportDiscountResponse) {
	response = &ExportDiscountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出折扣列表
func (c *Client) ExportDiscount(request *ExportDiscountRequest) (response *ExportDiscountResponse, err error) {
	if request == nil {
		request = NewExportDiscountRequest()
	}
	response = NewExportDiscountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDebateListGatewayRequest() (request *DescribeDebateListGatewayRequest) {
	request = &DescribeDebateListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeDebateListGateway")
	return
}

func NewDescribeDebateListGatewayResponse() (response *DescribeDebateListGatewayResponse) {
	response = &DescribeDebateListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端查询用户账单是否欠费和欠费月份
func (c *Client) DescribeDebateListGateway(request *DescribeDebateListGatewayRequest) (response *DescribeDebateListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeDebateListGatewayRequest()
	}
	response = NewDescribeDebateListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBasicAccountListGatewayRequest() (request *DescribeBasicAccountListGatewayRequest) {
	request = &DescribeBasicAccountListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBasicAccountListGateway")
	return
}

func NewDescribeBasicAccountListGatewayResponse() (response *DescribeBasicAccountListGatewayResponse) {
	response = &DescribeBasicAccountListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据uin批量获取账户基本信息
func (c *Client) DescribeBasicAccountListGateway(request *DescribeBasicAccountListGatewayRequest) (response *DescribeBasicAccountListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBasicAccountListGatewayRequest()
	}
	response = NewDescribeBasicAccountListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContractTypeListGatewayRequest() (request *DescribeContractTypeListGatewayRequest) {
	request = &DescribeContractTypeListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeContractTypeListGateway")
	return
}

func NewDescribeContractTypeListGatewayResponse() (response *DescribeContractTypeListGatewayResponse) {
	response = &DescribeContractTypeListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取合同类型列表接口
func (c *Client) DescribeContractTypeListGateway(request *DescribeContractTypeListGatewayRequest) (response *DescribeContractTypeListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeContractTypeListGatewayRequest()
	}
	response = NewDescribeContractTypeListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDiscountGatewayRequest() (request *ModifyDiscountGatewayRequest) {
	request = &ModifyDiscountGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ModifyDiscountGateway")
	return
}

func NewModifyDiscountGatewayResponse() (response *ModifyDiscountGatewayResponse) {
	response = &ModifyDiscountGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改折扣
func (c *Client) ModifyDiscountGateway(request *ModifyDiscountGatewayRequest) (response *ModifyDiscountGatewayResponse, err error) {
	if request == nil {
		request = NewModifyDiscountGatewayRequest()
	}
	response = NewModifyDiscountGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewExportDealListGatewayRequest() (request *ExportDealListGatewayRequest) {
	request = &ExportDealListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ExportDealListGateway")
	return
}

func NewExportDealListGatewayResponse() (response *ExportDealListGatewayResponse) {
	response = &ExportDealListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端导入订单列表信息
func (c *Client) ExportDealListGateway(request *ExportDealListGatewayRequest) (response *ExportDealListGatewayResponse, err error) {
	if request == nil {
		request = NewExportDealListGatewayRequest()
	}
	response = NewExportDealListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContractListRequest() (request *DescribeContractListRequest) {
	request = &DescribeContractListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeContractList")
	return
}

func NewDescribeContractListResponse() (response *DescribeContractListResponse) {
	response = &DescribeContractListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询合同列表接口
func (c *Client) DescribeContractList(request *DescribeContractListRequest) (response *DescribeContractListResponse, err error) {
	if request == nil {
		request = NewDescribeContractListRequest()
	}
	response = NewDescribeContractListResponse()
	err = c.Send(request, response)
	return
}

func NewEditBreakdownModeRequest() (request *EditBreakdownModeRequest) {
	request = &EditBreakdownModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "EditBreakdownMode")
	return
}

func NewEditBreakdownModeResponse() (response *EditBreakdownModeResponse) {
	response = &EditBreakdownModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口编辑当前所处的模式，开关故障，编辑白名单，开关临时配额
func (c *Client) EditBreakdownMode(request *EditBreakdownModeRequest) (response *EditBreakdownModeResponse, err error) {
	if request == nil {
		request = NewEditBreakdownModeRequest()
	}
	response = NewEditBreakdownModeResponse()
	err = c.Send(request, response)
	return
}

func NewGetResourceBillRequest() (request *GetResourceBillRequest) {
	request = &GetResourceBillRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetResourceBill")
	return
}

func NewGetResourceBillResponse() (response *GetResourceBillResponse) {
	response = &GetResourceBillResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的账单信息
func (c *Client) GetResourceBill(request *GetResourceBillRequest) (response *GetResourceBillResponse, err error) {
	if request == nil {
		request = NewGetResourceBillRequest()
	}
	response = NewGetResourceBillResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductStrategyGatewayRequest() (request *GetProductStrategyGatewayRequest) {
	request = &GetProductStrategyGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetProductStrategyGateway")
	return
}

func NewGetProductStrategyGatewayResponse() (response *GetProductStrategyGatewayResponse) {
	response = &GetProductStrategyGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品定义列表-策略库
func (c *Client) GetProductStrategyGateway(request *GetProductStrategyGatewayRequest) (response *GetProductStrategyGatewayResponse, err error) {
	if request == nil {
		request = NewGetProductStrategyGatewayRequest()
	}
	response = NewGetProductStrategyGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewQueryHiddenProductRequest() (request *QueryHiddenProductRequest) {
	request = &QueryHiddenProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "QueryHiddenProduct")
	return
}

func NewQueryHiddenProductResponse() (response *QueryHiddenProductResponse) {
	response = &QueryHiddenProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询隐藏冗余商品项接口
func (c *Client) QueryHiddenProduct(request *QueryHiddenProductRequest) (response *QueryHiddenProductResponse, err error) {
	if request == nil {
		request = NewQueryHiddenProductRequest()
	}
	response = NewQueryHiddenProductResponse()
	err = c.Send(request, response)
	return
}

func NewQueryHiddenProductGatewayRequest() (request *QueryHiddenProductGatewayRequest) {
	request = &QueryHiddenProductGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "QueryHiddenProductGateway")
	return
}

func NewQueryHiddenProductGatewayResponse() (response *QueryHiddenProductGatewayResponse) {
	response = &QueryHiddenProductGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询隐藏冗余商品项接口
func (c *Client) QueryHiddenProductGateway(request *QueryHiddenProductGatewayRequest) (response *QueryHiddenProductGatewayResponse, err error) {
	if request == nil {
		request = NewQueryHiddenProductGatewayRequest()
	}
	response = NewQueryHiddenProductGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewSaveCommonDiscountGatewayRequest() (request *SaveCommonDiscountGatewayRequest) {
	request = &SaveCommonDiscountGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SaveCommonDiscountGateway")
	return
}

func NewSaveCommonDiscountGatewayResponse() (response *SaveCommonDiscountGatewayResponse) {
	response = &SaveCommonDiscountGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增官网折扣
func (c *Client) SaveCommonDiscountGateway(request *SaveCommonDiscountGatewayRequest) (response *SaveCommonDiscountGatewayResponse, err error) {
	if request == nil {
		request = NewSaveCommonDiscountGatewayRequest()
	}
	response = NewSaveCommonDiscountGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewExpireDiscountRequest() (request *ExpireDiscountRequest) {
	request = &ExpireDiscountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ExpireDiscount")
	return
}

func NewExpireDiscountResponse() (response *ExpireDiscountResponse) {
	response = &ExpireDiscountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 废弃已经生效的折扣
func (c *Client) ExpireDiscount(request *ExpireDiscountRequest) (response *ExpireDiscountResponse, err error) {
	if request == nil {
		request = NewExpireDiscountRequest()
	}
	response = NewExpireDiscountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadQuotasListFileRequest() (request *DescribeDownloadQuotasListFileRequest) {
	request = &DescribeDownloadQuotasListFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeDownloadQuotasListFile")
	return
}

func NewDescribeDownloadQuotasListFileResponse() (response *DescribeDownloadQuotasListFileResponse) {
	response = &DescribeDownloadQuotasListFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 计费配额导出接口
func (c *Client) DescribeDownloadQuotasListFile(request *DescribeDownloadQuotasListFileRequest) (response *DescribeDownloadQuotasListFileResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadQuotasListFileRequest()
	}
	response = NewDescribeDownloadQuotasListFileResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSignatureGatewayRequest() (request *DescribeSignatureGatewayRequest) {
	request = &DescribeSignatureGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeSignatureGateway")
	return
}

func NewDescribeSignatureGatewayResponse() (response *DescribeSignatureGatewayResponse) {
	response = &DescribeSignatureGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 前端拿到签名和上传地址后把文件和签名放在form表单中使用post方法上传到私有云。
// 具体流程参照https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-UsingHTTPPOST.html
// 注意：签名过期时间设置为2小时  文件file需放在表单最后 文件类型Content-Type 只支持application开头的
func (c *Client) DescribeSignatureGateway(request *DescribeSignatureGatewayRequest) (response *DescribeSignatureGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeSignatureGatewayRequest()
	}
	response = NewDescribeSignatureGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceModifyLogGatewayRequest() (request *ListResourceModifyLogGatewayRequest) {
	request = &ListResourceModifyLogGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ListResourceModifyLogGateway")
	return
}

func NewListResourceModifyLogGatewayResponse() (response *ListResourceModifyLogGatewayResponse) {
	response = &ListResourceModifyLogGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资源变配记录
func (c *Client) ListResourceModifyLogGateway(request *ListResourceModifyLogGatewayRequest) (response *ListResourceModifyLogGatewayResponse, err error) {
	if request == nil {
		request = NewListResourceModifyLogGatewayRequest()
	}
	response = NewListResourceModifyLogGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewSetAutoFlagRequest() (request *SetAutoFlagRequest) {
	request = &SetAutoFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SetAutoFlag")
	return
}

func NewSetAutoFlagResponse() (response *SetAutoFlagResponse) {
	response = &SetAutoFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// (暂时没用,废弃)设置续费类型
func (c *Client) SetAutoFlag(request *SetAutoFlagRequest) (response *SetAutoFlagResponse, err error) {
	if request == nil {
		request = NewSetAutoFlagRequest()
	}
	response = NewSetAutoFlagResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadRecordListRequest() (request *DescribeBillDownloadRecordListRequest) {
	request = &DescribeBillDownloadRecordListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillDownloadRecordList")
	return
}

func NewDescribeBillDownloadRecordListResponse() (response *DescribeBillDownloadRecordListResponse) {
	response = &DescribeBillDownloadRecordListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取下载记录列表
func (c *Client) DescribeBillDownloadRecordList(request *DescribeBillDownloadRecordListRequest) (response *DescribeBillDownloadRecordListResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadRecordListRequest()
	}
	response = NewDescribeBillDownloadRecordListResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadDosagesGatewayRequest() (request *DownloadDosagesGatewayRequest) {
	request = &DownloadDosagesGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DownloadDosagesGateway")
	return
}

func NewDownloadDosagesGatewayResponse() (response *DownloadDosagesGatewayResponse) {
	response = &DownloadDosagesGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用量统计下载
func (c *Client) DownloadDosagesGateway(request *DownloadDosagesGatewayRequest) (response *DownloadDosagesGatewayResponse, err error) {
	if request == nil {
		request = NewDownloadDosagesGatewayRequest()
	}
	response = NewDownloadDosagesGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewSetBspProductStatusGatewayRequest() (request *SetBspProductStatusGatewayRequest) {
	request = &SetBspProductStatusGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SetBspProductStatusGateway")
	return
}

func NewSetBspProductStatusGatewayResponse() (response *SetBspProductStatusGatewayResponse) {
	response = &SetBspProductStatusGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 变更产品定义状态
func (c *Client) SetBspProductStatusGateway(request *SetBspProductStatusGatewayRequest) (response *SetBspProductStatusGatewayResponse, err error) {
	if request == nil {
		request = NewSetBspProductStatusGatewayRequest()
	}
	response = NewSetBspProductStatusGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewQueryBSPSubBillingItemGatewayRequest() (request *QueryBSPSubBillingItemGatewayRequest) {
	request = &QueryBSPSubBillingItemGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "QueryBSPSubBillingItemGateway")
	return
}

func NewQueryBSPSubBillingItemGatewayResponse() (response *QueryBSPSubBillingItemGatewayResponse) {
	response = &QueryBSPSubBillingItemGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询四层产品定义及展示名称、产品标识
func (c *Client) QueryBSPSubBillingItemGateway(request *QueryBSPSubBillingItemGatewayRequest) (response *QueryBSPSubBillingItemGatewayResponse, err error) {
	if request == nil {
		request = NewQueryBSPSubBillingItemGatewayRequest()
	}
	response = NewQueryBSPSubBillingItemGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGetDownloadUrlWithAuthGatewayRequest() (request *DescribeGetDownloadUrlWithAuthGatewayRequest) {
	request = &DescribeGetDownloadUrlWithAuthGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeGetDownloadUrlWithAuthGateway")
	return
}

func NewDescribeGetDownloadUrlWithAuthGatewayResponse() (response *DescribeGetDownloadUrlWithAuthGatewayResponse) {
	response = &DescribeGetDownloadUrlWithAuthGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据任务ID查询文件下载地址
func (c *Client) DescribeGetDownloadUrlWithAuthGateway(request *DescribeGetDownloadUrlWithAuthGatewayRequest) (response *DescribeGetDownloadUrlWithAuthGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeGetDownloadUrlWithAuthGatewayRequest()
	}
	response = NewDescribeGetDownloadUrlWithAuthGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsolidatedBillDownloadUrlRequest() (request *DescribeConsolidatedBillDownloadUrlRequest) {
	request = &DescribeConsolidatedBillDownloadUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeConsolidatedBillDownloadUrl")
	return
}

func NewDescribeConsolidatedBillDownloadUrlResponse() (response *DescribeConsolidatedBillDownloadUrlResponse) {
	response = &DescribeConsolidatedBillDownloadUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账单包下载地址
func (c *Client) DescribeConsolidatedBillDownloadUrl(request *DescribeConsolidatedBillDownloadUrlRequest) (response *DescribeConsolidatedBillDownloadUrlResponse, err error) {
	if request == nil {
		request = NewDescribeConsolidatedBillDownloadUrlRequest()
	}
	response = NewDescribeConsolidatedBillDownloadUrlResponse()
	err = c.Send(request, response)
	return
}

func NewGetComprehensiveProductStrategyRouteGatewayRequest() (request *GetComprehensiveProductStrategyRouteGatewayRequest) {
	request = &GetComprehensiveProductStrategyRouteGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetComprehensiveProductStrategyRouteGateway")
	return
}

func NewGetComprehensiveProductStrategyRouteGatewayResponse() (response *GetComprehensiveProductStrategyRouteGatewayResponse) {
	response = &GetComprehensiveProductStrategyRouteGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品（一层）策略接口
func (c *Client) GetComprehensiveProductStrategyRouteGateway(request *GetComprehensiveProductStrategyRouteGatewayRequest) (response *GetComprehensiveProductStrategyRouteGatewayResponse, err error) {
	if request == nil {
		request = NewGetComprehensiveProductStrategyRouteGatewayRequest()
	}
	response = NewGetComprehensiveProductStrategyRouteGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewIsolateResourceGatewayRequest() (request *IsolateResourceGatewayRequest) {
	request = &IsolateResourceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "IsolateResourceGateway")
	return
}

func NewIsolateResourceGatewayResponse() (response *IsolateResourceGatewayResponse) {
	response = &IsolateResourceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端资源隔离接口
func (c *Client) IsolateResourceGateway(request *IsolateResourceGatewayRequest) (response *IsolateResourceGatewayResponse, err error) {
	if request == nil {
		request = NewIsolateResourceGatewayRequest()
	}
	response = NewIsolateResourceGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductStrategyRequest() (request *GetProductStrategyRequest) {
	request = &GetProductStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetProductStrategy")
	return
}

func NewGetProductStrategyResponse() (response *GetProductStrategyResponse) {
	response = &GetProductStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品定义列表-策略库
func (c *Client) GetProductStrategy(request *GetProductStrategyRequest) (response *GetProductStrategyResponse, err error) {
	if request == nil {
		request = NewGetProductStrategyRequest()
	}
	response = NewGetProductStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewSetPropertyRequest() (request *SetPropertyRequest) {
	request = &SetPropertyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SetProperty")
	return
}

func NewSetPropertyResponse() (response *SetPropertyResponse) {
	response = &SetPropertyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增、编辑售卖属性
func (c *Client) SetProperty(request *SetPropertyRequest) (response *SetPropertyResponse, err error) {
	if request == nil {
		request = NewSetPropertyRequest()
	}
	response = NewSetPropertyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadResourceDetailRequest() (request *DescribeBillDownloadResourceDetailRequest) {
	request = &DescribeBillDownloadResourceDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillDownloadResourceDetail")
	return
}

func NewDescribeBillDownloadResourceDetailResponse() (response *DescribeBillDownloadResourceDetailResponse) {
	response = &DescribeBillDownloadResourceDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载定制化账单明细
func (c *Client) DescribeBillDownloadResourceDetail(request *DescribeBillDownloadResourceDetailRequest) (response *DescribeBillDownloadResourceDetailResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadResourceDetailRequest()
	}
	response = NewDescribeBillDownloadResourceDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubProductsRequest() (request *DescribeSubProductsRequest) {
	request = &DescribeSubProductsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeSubProducts")
	return
}

func NewDescribeSubProductsResponse() (response *DescribeSubProductsResponse) {
	response = &DescribeSubProductsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取子商品信息
func (c *Client) DescribeSubProducts(request *DescribeSubProductsRequest) (response *DescribeSubProductsResponse, err error) {
	if request == nil {
		request = NewDescribeSubProductsRequest()
	}
	response = NewDescribeSubProductsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePropertiesRequest() (request *DescribePropertiesRequest) {
	request = &DescribePropertiesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeProperties")
	return
}

func NewDescribePropertiesResponse() (response *DescribePropertiesResponse) {
	response = &DescribePropertiesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询售卖属性列表
func (c *Client) DescribeProperties(request *DescribePropertiesRequest) (response *DescribePropertiesResponse, err error) {
	if request == nil {
		request = NewDescribePropertiesRequest()
	}
	response = NewDescribePropertiesResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceDetailRequest() (request *ListResourceDetailRequest) {
	request = &ListResourceDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ListResourceDetail")
	return
}

func NewListResourceDetailResponse() (response *ListResourceDetailResponse) {
	response = &ListResourceDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资源详情
func (c *Client) ListResourceDetail(request *ListResourceDetailRequest) (response *ListResourceDetailResponse, err error) {
	if request == nil {
		request = NewListResourceDetailRequest()
	}
	response = NewListResourceDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillingItemsGatewayRequest() (request *DescribeBillingItemsGatewayRequest) {
	request = &DescribeBillingItemsGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillingItemsGateway")
	return
}

func NewDescribeBillingItemsGatewayResponse() (response *DescribeBillingItemsGatewayResponse) {
	response = &DescribeBillingItemsGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取计费项信息（第三层）
func (c *Client) DescribeBillingItemsGateway(request *DescribeBillingItemsGatewayRequest) (response *DescribeBillingItemsGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillingItemsGatewayRequest()
	}
	response = NewDescribeBillingItemsGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewAddComprehensiveStrategyRouteGatewayRequest() (request *AddComprehensiveStrategyRouteGatewayRequest) {
	request = &AddComprehensiveStrategyRouteGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "AddComprehensiveStrategyRouteGateway")
	return
}

func NewAddComprehensiveStrategyRouteGatewayResponse() (response *AddComprehensiveStrategyRouteGatewayResponse) {
	response = &AddComprehensiveStrategyRouteGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 可以新增第三方和预设的产品策略
func (c *Client) AddComprehensiveStrategyRouteGateway(request *AddComprehensiveStrategyRouteGatewayRequest) (response *AddComprehensiveStrategyRouteGatewayResponse, err error) {
	if request == nil {
		request = NewAddComprehensiveStrategyRouteGatewayRequest()
	}
	response = NewAddComprehensiveStrategyRouteGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewAuditInvoiceInfoGatewayRequest() (request *AuditInvoiceInfoGatewayRequest) {
	request = &AuditInvoiceInfoGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "AuditInvoiceInfoGateway")
	return
}

func NewAuditInvoiceInfoGatewayResponse() (response *AuditInvoiceInfoGatewayResponse) {
	response = &AuditInvoiceInfoGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端-审核发票信息
func (c *Client) AuditInvoiceInfoGateway(request *AuditInvoiceInfoGatewayRequest) (response *AuditInvoiceInfoGatewayResponse, err error) {
	if request == nil {
		request = NewAuditInvoiceInfoGatewayRequest()
	}
	response = NewAuditInvoiceInfoGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDetailByResourceGatewayRequest() (request *DescribeBillDetailByResourceGatewayRequest) {
	request = &DescribeBillDetailByResourceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillDetailByResourceGateway")
	return
}

func NewDescribeBillDetailByResourceGatewayResponse() (response *DescribeBillDetailByResourceGatewayResponse) {
	response = &DescribeBillDetailByResourceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源花费详情
func (c *Client) DescribeBillDetailByResourceGateway(request *DescribeBillDetailByResourceGatewayRequest) (response *DescribeBillDetailByResourceGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillDetailByResourceGatewayRequest()
	}
	response = NewDescribeBillDetailByResourceGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyComprehensiveStrategyGatewayRequest() (request *ModifyComprehensiveStrategyGatewayRequest) {
	request = &ModifyComprehensiveStrategyGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ModifyComprehensiveStrategyGateway")
	return
}

func NewModifyComprehensiveStrategyGatewayResponse() (response *ModifyComprehensiveStrategyGatewayResponse) {
	response = &ModifyComprehensiveStrategyGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改策略包含第三方自定义和预设
func (c *Client) ModifyComprehensiveStrategyGateway(request *ModifyComprehensiveStrategyGatewayRequest) (response *ModifyComprehensiveStrategyGatewayResponse, err error) {
	if request == nil {
		request = NewModifyComprehensiveStrategyGatewayRequest()
	}
	response = NewModifyComprehensiveStrategyGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBspDownloadRecordListGatewayRequest() (request *DescribeBspDownloadRecordListGatewayRequest) {
	request = &DescribeBspDownloadRecordListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBspDownloadRecordListGateway")
	return
}

func NewDescribeBspDownloadRecordListGatewayResponse() (response *DescribeBspDownloadRecordListGatewayResponse) {
	response = &DescribeBspDownloadRecordListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取下载记录列表
func (c *Client) DescribeBspDownloadRecordListGateway(request *DescribeBspDownloadRecordListGatewayRequest) (response *DescribeBspDownloadRecordListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBspDownloadRecordListGatewayRequest()
	}
	response = NewDescribeBspDownloadRecordListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillingUinQuotaGatewayRequest() (request *DescribeBillingUinQuotaGatewayRequest) {
	request = &DescribeBillingUinQuotaGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillingUinQuotaGateway")
	return
}

func NewDescribeBillingUinQuotaGatewayResponse() (response *DescribeBillingUinQuotaGatewayResponse) {
	response = &DescribeBillingUinQuotaGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端通过appid查询主账号配额全量
func (c *Client) DescribeBillingUinQuotaGateway(request *DescribeBillingUinQuotaGatewayRequest) (response *DescribeBillingUinQuotaGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillingUinQuotaGatewayRequest()
	}
	response = NewDescribeBillingUinQuotaGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReconciliationListRequest() (request *DescribeReconciliationListRequest) {
	request = &DescribeReconciliationListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeReconciliationList")
	return
}

func NewDescribeReconciliationListResponse() (response *DescribeReconciliationListResponse) {
	response = &DescribeReconciliationListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调账记录查询调账列表
func (c *Client) DescribeReconciliationList(request *DescribeReconciliationListRequest) (response *DescribeReconciliationListResponse, err error) {
	if request == nil {
		request = NewDescribeReconciliationListRequest()
	}
	response = NewDescribeReconciliationListResponse()
	err = c.Send(request, response)
	return
}

func NewListLackingQuotaRecordGatewayRequest() (request *ListLackingQuotaRecordGatewayRequest) {
	request = &ListLackingQuotaRecordGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ListLackingQuotaRecordGateway")
	return
}

func NewListLackingQuotaRecordGatewayResponse() (response *ListLackingQuotaRecordGatewayResponse) {
	response = &ListLackingQuotaRecordGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取配额不足列表
func (c *Client) ListLackingQuotaRecordGateway(request *ListLackingQuotaRecordGatewayRequest) (response *ListLackingQuotaRecordGatewayResponse, err error) {
	if request == nil {
		request = NewListLackingQuotaRecordGatewayRequest()
	}
	response = NewListLackingQuotaRecordGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetComprehensiveProdctStrategyRequest() (request *GetComprehensiveProdctStrategyRequest) {
	request = &GetComprehensiveProdctStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetComprehensiveProdctStrategy")
	return
}

func NewGetComprehensiveProdctStrategyResponse() (response *GetComprehensiveProdctStrategyResponse) {
	response = &GetComprehensiveProdctStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品（一层）策略接口
func (c *Client) GetComprehensiveProdctStrategy(request *GetComprehensiveProdctStrategyRequest) (response *GetComprehensiveProdctStrategyResponse, err error) {
	if request == nil {
		request = NewGetComprehensiveProdctStrategyRequest()
	}
	response = NewGetComprehensiveProdctStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewAddProductGatewayRequest() (request *AddProductGatewayRequest) {
	request = &AddProductGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "AddProductGateway")
	return
}

func NewAddProductGatewayResponse() (response *AddProductGatewayResponse) {
	response = &AddProductGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加产品四层定义
func (c *Client) AddProductGateway(request *AddProductGatewayRequest) (response *AddProductGatewayResponse, err error) {
	if request == nil {
		request = NewAddProductGatewayRequest()
	}
	response = NewAddProductGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByTagGatewayRequest() (request *DescribeBillSummaryByTagGatewayRequest) {
	request = &DescribeBillSummaryByTagGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillSummaryByTagGateway")
	return
}

func NewDescribeBillSummaryByTagGatewayResponse() (response *DescribeBillSummaryByTagGatewayResponse) {
	response = &DescribeBillSummaryByTagGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取标签汇总费用分布
func (c *Client) DescribeBillSummaryByTagGateway(request *DescribeBillSummaryByTagGatewayRequest) (response *DescribeBillSummaryByTagGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByTagGatewayRequest()
	}
	response = NewDescribeBillSummaryByTagGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemitRecordOpRequest() (request *DescribeRemitRecordOpRequest) {
	request = &DescribeRemitRecordOpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeRemitRecordOp")
	return
}

func NewDescribeRemitRecordOpResponse() (response *DescribeRemitRecordOpResponse) {
	response = &DescribeRemitRecordOpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端查询汇款记录列表
func (c *Client) DescribeRemitRecordOp(request *DescribeRemitRecordOpRequest) (response *DescribeRemitRecordOpResponse, err error) {
	if request == nil {
		request = NewDescribeRemitRecordOpRequest()
	}
	response = NewDescribeRemitRecordOpResponse()
	err = c.Send(request, response)
	return
}

func NewGetDiscardProductRequest() (request *GetDiscardProductRequest) {
	request = &GetDiscardProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetDiscardProduct")
	return
}

func NewGetDiscardProductResponse() (response *GetDiscardProductResponse) {
	response = &GetDiscardProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取废弃的商品码
func (c *Client) GetDiscardProduct(request *GetDiscardProductRequest) (response *GetDiscardProductResponse, err error) {
	if request == nil {
		request = NewGetDiscardProductRequest()
	}
	response = NewGetDiscardProductResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemitRecordOpGatewayRequest() (request *DescribeRemitRecordOpGatewayRequest) {
	request = &DescribeRemitRecordOpGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeRemitRecordOpGateway")
	return
}

func NewDescribeRemitRecordOpGatewayResponse() (response *DescribeRemitRecordOpGatewayResponse) {
	response = &DescribeRemitRecordOpGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端查询汇款记录列表
func (c *Client) DescribeRemitRecordOpGateway(request *DescribeRemitRecordOpGatewayRequest) (response *DescribeRemitRecordOpGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeRemitRecordOpGatewayRequest()
	}
	response = NewDescribeRemitRecordOpGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDiscountRequest() (request *ModifyDiscountRequest) {
	request = &ModifyDiscountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ModifyDiscount")
	return
}

func NewModifyDiscountResponse() (response *ModifyDiscountResponse) {
	response = &ModifyDiscountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改折扣
func (c *Client) ModifyDiscount(request *ModifyDiscountRequest) (response *ModifyDiscountResponse, err error) {
	if request == nil {
		request = NewModifyDiscountRequest()
	}
	response = NewModifyDiscountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReconciliationRequest() (request *ModifyReconciliationRequest) {
	request = &ModifyReconciliationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ModifyReconciliation")
	return
}

func NewModifyReconciliationResponse() (response *ModifyReconciliationResponse) {
	response = &ModifyReconciliationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提交调账
func (c *Client) ModifyReconciliation(request *ModifyReconciliationRequest) (response *ModifyReconciliationResponse, err error) {
	if request == nil {
		request = NewModifyReconciliationRequest()
	}
	response = NewModifyReconciliationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadListGatewayRequest() (request *DescribeBillDownloadListGatewayRequest) {
	request = &DescribeBillDownloadListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillDownloadListGateway")
	return
}

func NewDescribeBillDownloadListGatewayResponse() (response *DescribeBillDownloadListGatewayResponse) {
	response = &DescribeBillDownloadListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账单下载地址列表
func (c *Client) DescribeBillDownloadListGateway(request *DescribeBillDownloadListGatewayRequest) (response *DescribeBillDownloadListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadListGatewayRequest()
	}
	response = NewDescribeBillDownloadListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccountInfoRequest() (request *ModifyAccountInfoRequest) {
	request = &ModifyAccountInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ModifyAccountInfo")
	return
}

func NewModifyAccountInfoResponse() (response *ModifyAccountInfoResponse) {
	response = &ModifyAccountInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改后付费账户的账期； 销户\冻结\解冻；
func (c *Client) ModifyAccountInfo(request *ModifyAccountInfoRequest) (response *ModifyAccountInfoResponse, err error) {
	if request == nil {
		request = NewModifyAccountInfoRequest()
	}
	response = NewModifyAccountInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountListGatewayRequest() (request *DescribeAccountListGatewayRequest) {
	request = &DescribeAccountListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeAccountListGateway")
	return
}

func NewDescribeAccountListGatewayResponse() (response *DescribeAccountListGatewayResponse) {
	response = &DescribeAccountListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端查询帐户列表和用户代金券总金额
func (c *Client) DescribeAccountListGateway(request *DescribeAccountListGatewayRequest) (response *DescribeAccountListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeAccountListGatewayRequest()
	}
	response = NewDescribeAccountListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProductGatewayRequest() (request *UpdateProductGatewayRequest) {
	request = &UpdateProductGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "UpdateProductGateway")
	return
}

func NewUpdateProductGatewayResponse() (response *UpdateProductGatewayResponse) {
	response = &UpdateProductGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【计量】修改产品四层定义
func (c *Client) UpdateProductGateway(request *UpdateProductGatewayRequest) (response *UpdateProductGatewayResponse, err error) {
	if request == nil {
		request = NewUpdateProductGatewayRequest()
	}
	response = NewUpdateProductGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDownloadRecordRequest() (request *CreateDownloadRecordRequest) {
	request = &CreateDownloadRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "CreateDownloadRecord")
	return
}

func NewCreateDownloadRecordResponse() (response *CreateDownloadRecordResponse) {
	response = &CreateDownloadRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户点击下载时通知后台
func (c *Client) CreateDownloadRecord(request *CreateDownloadRecordRequest) (response *CreateDownloadRecordResponse, err error) {
	if request == nil {
		request = NewCreateDownloadRecordRequest()
	}
	response = NewCreateDownloadRecordResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceDetailDownloadGatewayRequest() (request *ListResourceDetailDownloadGatewayRequest) {
	request = &ListResourceDetailDownloadGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ListResourceDetailDownloadGateway")
	return
}

func NewListResourceDetailDownloadGatewayResponse() (response *ListResourceDetailDownloadGatewayResponse) {
	response = &ListResourceDetailDownloadGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载资源详情
func (c *Client) ListResourceDetailDownloadGateway(request *ListResourceDetailDownloadGatewayRequest) (response *ListResourceDetailDownloadGatewayResponse, err error) {
	if request == nil {
		request = NewListResourceDetailDownloadGatewayRequest()
	}
	response = NewListResourceDetailDownloadGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewQueryBreakdownModeGatewayRequest() (request *QueryBreakdownModeGatewayRequest) {
	request = &QueryBreakdownModeGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "QueryBreakdownModeGateway")
	return
}

func NewQueryBreakdownModeGatewayResponse() (response *QueryBreakdownModeGatewayResponse) {
	response = &QueryBreakdownModeGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口返回当前所处模式，如果是故障中，则返回白名单，临时配额开关
func (c *Client) QueryBreakdownModeGateway(request *QueryBreakdownModeGatewayRequest) (response *QueryBreakdownModeGatewayResponse, err error) {
	if request == nil {
		request = NewQueryBreakdownModeGatewayRequest()
	}
	response = NewQueryBreakdownModeGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGeneralDataGatewayRequest() (request *DescribeGeneralDataGatewayRequest) {
	request = &DescribeGeneralDataGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeGeneralDataGateway")
	return
}

func NewDescribeGeneralDataGatewayResponse() (response *DescribeGeneralDataGatewayResponse) {
	response = &DescribeGeneralDataGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口用于通用查询
func (c *Client) DescribeGeneralDataGateway(request *DescribeGeneralDataGatewayRequest) (response *DescribeGeneralDataGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeGeneralDataGatewayRequest()
	}
	response = NewDescribeGeneralDataGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceListRequest() (request *DescribeResourceListRequest) {
	request = &DescribeResourceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeResourceList")
	return
}

func NewDescribeResourceListResponse() (response *DescribeResourceListResponse) {
	response = &DescribeResourceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端查询用户产品列表，只支持到一层产品
func (c *Client) DescribeResourceList(request *DescribeResourceListRequest) (response *DescribeResourceListResponse, err error) {
	if request == nil {
		request = NewDescribeResourceListRequest()
	}
	response = NewDescribeResourceListResponse()
	err = c.Send(request, response)
	return
}

func NewGetComprehensiveStrategyWebConfigGatewayRequest() (request *GetComprehensiveStrategyWebConfigGatewayRequest) {
	request = &GetComprehensiveStrategyWebConfigGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetComprehensiveStrategyWebConfigGateway")
	return
}

func NewGetComprehensiveStrategyWebConfigGatewayResponse() (response *GetComprehensiveStrategyWebConfigGatewayResponse) {
	response = &GetComprehensiveStrategyWebConfigGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有枚举配置
func (c *Client) GetComprehensiveStrategyWebConfigGateway(request *GetComprehensiveStrategyWebConfigGatewayRequest) (response *GetComprehensiveStrategyWebConfigGatewayResponse, err error) {
	if request == nil {
		request = NewGetComprehensiveStrategyWebConfigGatewayRequest()
	}
	response = NewGetComprehensiveStrategyWebConfigGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewQueryRecoverUinRequest() (request *QueryRecoverUinRequest) {
	request = &QueryRecoverUinRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "QueryRecoverUin")
	return
}

func NewQueryRecoverUinResponse() (response *QueryRecoverUinResponse) {
	response = &QueryRecoverUinResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口查询当前的恢复用户列表，支持分页查询，可以通过uin模糊匹配
func (c *Client) QueryRecoverUin(request *QueryRecoverUinRequest) (response *QueryRecoverUinResponse, err error) {
	if request == nil {
		request = NewQueryRecoverUinRequest()
	}
	response = NewQueryRecoverUinResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByPayModeRequest() (request *DescribeBillSummaryByPayModeRequest) {
	request = &DescribeBillSummaryByPayModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillSummaryByPayMode")
	return
}

func NewDescribeBillSummaryByPayModeResponse() (response *DescribeBillSummaryByPayModeResponse) {
	response = &DescribeBillSummaryByPayModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取计费模式汇总费用分布
func (c *Client) DescribeBillSummaryByPayMode(request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByPayModeRequest()
	}
	response = NewDescribeBillSummaryByPayModeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductRelationsRequest() (request *DescribeProductRelationsRequest) {
	request = &DescribeProductRelationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeProductRelations")
	return
}

func NewDescribeProductRelationsResponse() (response *DescribeProductRelationsResponse) {
	response = &DescribeProductRelationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品定义列表
func (c *Client) DescribeProductRelations(request *DescribeProductRelationsRequest) (response *DescribeProductRelationsResponse, err error) {
	if request == nil {
		request = NewDescribeProductRelationsRequest()
	}
	response = NewDescribeProductRelationsResponse()
	err = c.Send(request, response)
	return
}

func NewQueryPublicAccountRequest() (request *QueryPublicAccountRequest) {
	request = &QueryPublicAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "QueryPublicAccount")
	return
}

func NewQueryPublicAccountResponse() (response *QueryPublicAccountResponse) {
	response = &QueryPublicAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询公共帐户信息
func (c *Client) QueryPublicAccount(request *QueryPublicAccountRequest) (response *QueryPublicAccountResponse, err error) {
	if request == nil {
		request = NewQueryPublicAccountRequest()
	}
	response = NewQueryPublicAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadRecordListGatewayRequest() (request *DescribeBillDownloadRecordListGatewayRequest) {
	request = &DescribeBillDownloadRecordListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillDownloadRecordListGateway")
	return
}

func NewDescribeBillDownloadRecordListGatewayResponse() (response *DescribeBillDownloadRecordListGatewayResponse) {
	response = &DescribeBillDownloadRecordListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取下载记录列表
func (c *Client) DescribeBillDownloadRecordListGateway(request *DescribeBillDownloadRecordListGatewayRequest) (response *DescribeBillDownloadRecordListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadRecordListGatewayRequest()
	}
	response = NewDescribeBillDownloadRecordListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceDetailListGatewayRequest() (request *DescribeResourceDetailListGatewayRequest) {
	request = &DescribeResourceDetailListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeResourceDetailListGateway")
	return
}

func NewDescribeResourceDetailListGatewayResponse() (response *DescribeResourceDetailListGatewayResponse) {
	response = &DescribeResourceDetailListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户续费参数
func (c *Client) DescribeResourceDetailListGateway(request *DescribeResourceDetailListGatewayRequest) (response *DescribeResourceDetailListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeResourceDetailListGatewayRequest()
	}
	response = NewDescribeResourceDetailListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeByTagGatewayRequest() (request *DescribeBillFeeByTagGatewayRequest) {
	request = &DescribeBillFeeByTagGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillFeeByTagGateway")
	return
}

func NewDescribeBillFeeByTagGatewayResponse() (response *DescribeBillFeeByTagGatewayResponse) {
	response = &DescribeBillFeeByTagGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询按标签汇总明细数据
func (c *Client) DescribeBillFeeByTagGateway(request *DescribeBillFeeByTagGatewayRequest) (response *DescribeBillFeeByTagGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeByTagGatewayRequest()
	}
	response = NewDescribeBillFeeByTagGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductTreeGatewayRequest() (request *GetProductTreeGatewayRequest) {
	request = &GetProductTreeGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetProductTreeGateway")
	return
}

func NewGetProductTreeGatewayResponse() (response *GetProductTreeGatewayResponse) {
	response = &GetProductTreeGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过商品码，子商品码，计费项，计费细项来查询产品定义，参数都可选，如果不传则返回全部数据
func (c *Client) GetProductTreeGateway(request *GetProductTreeGatewayRequest) (response *GetProductTreeGatewayResponse, err error) {
	if request == nil {
		request = NewGetProductTreeGatewayRequest()
	}
	response = NewGetProductTreeGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewSetProductStatusRequest() (request *SetProductStatusRequest) {
	request = &SetProductStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SetProductStatus")
	return
}

func NewSetProductStatusResponse() (response *SetProductStatusResponse) {
	response = &SetProductStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 变更产品定义状态
func (c *Client) SetProductStatus(request *SetProductStatusRequest) (response *SetProductStatusResponse, err error) {
	if request == nil {
		request = NewSetProductStatusRequest()
	}
	response = NewSetProductStatusResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductListRequest() (request *GetProductListRequest) {
	request = &GetProductListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetProductList")
	return
}

func NewGetProductListResponse() (response *GetProductListResponse) {
	response = &GetProductListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过商品码，子商品码，计费项，计费细项来查询产品定义，参数都可选，如果不传则返回全部数据
func (c *Client) GetProductList(request *GetProductListRequest) (response *GetProductListResponse, err error) {
	if request == nil {
		request = NewGetProductListRequest()
	}
	response = NewGetProductListResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductDeployStatusGatewayRequest() (request *GetProductDeployStatusGatewayRequest) {
	request = &GetProductDeployStatusGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetProductDeployStatusGateway")
	return
}

func NewGetProductDeployStatusGatewayResponse() (response *GetProductDeployStatusGatewayResponse) {
	response = &GetProductDeployStatusGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品部署计费状态
func (c *Client) GetProductDeployStatusGateway(request *GetProductDeployStatusGatewayRequest) (response *GetProductDeployStatusGatewayResponse, err error) {
	if request == nil {
		request = NewGetProductDeployStatusGatewayRequest()
	}
	response = NewGetProductDeployStatusGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewSetProductInfoRequest() (request *SetProductInfoRequest) {
	request = &SetProductInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SetProductInfo")
	return
}

func NewSetProductInfoResponse() (response *SetProductInfoResponse) {
	response = &SetProductInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品四层定义
func (c *Client) SetProductInfo(request *SetProductInfoRequest) (response *SetProductInfoResponse, err error) {
	if request == nil {
		request = NewSetProductInfoRequest()
	}
	response = NewSetProductInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQuotasRequest() (request *DescribeQuotasRequest) {
	request = &DescribeQuotasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeQuotas")
	return
}

func NewDescribeQuotasResponse() (response *DescribeQuotasResponse) {
	response = &DescribeQuotasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品配额列表查询
func (c *Client) DescribeQuotas(request *DescribeQuotasRequest) (response *DescribeQuotasResponse, err error) {
	if request == nil {
		request = NewDescribeQuotasRequest()
	}
	response = NewDescribeQuotasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRegionRequest() (request *DescribeRegionRequest) {
	request = &DescribeRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeRegion")
	return
}

func NewDescribeRegionResponse() (response *DescribeRegionResponse) {
	response = &DescribeRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// (暂时没用，废弃)查询地域列表
func (c *Client) DescribeRegion(request *DescribeRegionRequest) (response *DescribeRegionResponse, err error) {
	if request == nil {
		request = NewDescribeRegionRequest()
	}
	response = NewDescribeRegionResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseResourceGatewayRequest() (request *ReleaseResourceGatewayRequest) {
	request = &ReleaseResourceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ReleaseResourceGateway")
	return
}

func NewReleaseResourceGatewayResponse() (response *ReleaseResourceGatewayResponse) {
	response = &ReleaseResourceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端资源销毁
func (c *Client) ReleaseResourceGateway(request *ReleaseResourceGatewayRequest) (response *ReleaseResourceGatewayResponse, err error) {
	if request == nil {
		request = NewReleaseResourceGatewayRequest()
	}
	response = NewReleaseResourceGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWaterInfoRequest() (request *DescribeWaterInfoRequest) {
	request = &DescribeWaterInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeWaterInfo")
	return
}

func NewDescribeWaterInfoResponse() (response *DescribeWaterInfoResponse) {
	response = &DescribeWaterInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资源流水信息
func (c *Client) DescribeWaterInfo(request *DescribeWaterInfoRequest) (response *DescribeWaterInfoResponse, err error) {
	if request == nil {
		request = NewDescribeWaterInfoRequest()
	}
	response = NewDescribeWaterInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCouponsWaterRequest() (request *CouponsWaterRequest) {
	request = &CouponsWaterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "CouponsWater")
	return
}

func NewCouponsWaterResponse() (response *CouponsWaterResponse) {
	response = &CouponsWaterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 代金券流水
func (c *Client) CouponsWater(request *CouponsWaterRequest) (response *CouponsWaterResponse, err error) {
	if request == nil {
		request = NewCouponsWaterRequest()
	}
	response = NewCouponsWaterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeByPayModeRequest() (request *DescribeBillFeeByPayModeRequest) {
	request = &DescribeBillFeeByPayModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillFeeByPayMode")
	return
}

func NewDescribeBillFeeByPayModeResponse() (response *DescribeBillFeeByPayModeResponse) {
	response = &DescribeBillFeeByPayModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按计费模式汇总明细费用
func (c *Client) DescribeBillFeeByPayMode(request *DescribeBillFeeByPayModeRequest) (response *DescribeBillFeeByPayModeResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeByPayModeRequest()
	}
	response = NewDescribeBillFeeByPayModeResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserStrategyGatewayRequest() (request *GetUserStrategyGatewayRequest) {
	request = &GetUserStrategyGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetUserStrategyGateway")
	return
}

func NewGetUserStrategyGatewayResponse() (response *GetUserStrategyGatewayResponse) {
	response = &GetUserStrategyGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【计量】查询账号订购策略列表
func (c *Client) GetUserStrategyGateway(request *GetUserStrategyGatewayRequest) (response *GetUserStrategyGatewayResponse, err error) {
	if request == nil {
		request = NewGetUserStrategyGatewayRequest()
	}
	response = NewGetUserStrategyGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDiscountStatusGatewayRequest() (request *ModifyDiscountStatusGatewayRequest) {
	request = &ModifyDiscountStatusGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ModifyDiscountStatusGateway")
	return
}

func NewModifyDiscountStatusGatewayResponse() (response *ModifyDiscountStatusGatewayResponse) {
	response = &ModifyDiscountStatusGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改折扣状态
func (c *Client) ModifyDiscountStatusGateway(request *ModifyDiscountStatusGatewayRequest) (response *ModifyDiscountStatusGatewayResponse, err error) {
	if request == nil {
		request = NewModifyDiscountStatusGatewayRequest()
	}
	response = NewModifyDiscountStatusGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewSetHiddenProductGatewayRequest() (request *SetHiddenProductGatewayRequest) {
	request = &SetHiddenProductGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SetHiddenProductGateway")
	return
}

func NewSetHiddenProductGatewayResponse() (response *SetHiddenProductGatewayResponse) {
	response = &SetHiddenProductGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置隐藏冗余商品项接口，本接口接收前端请求存储和修改需隐藏的商品项。
func (c *Client) SetHiddenProductGateway(request *SetHiddenProductGatewayRequest) (response *SetHiddenProductGatewayResponse, err error) {
	if request == nil {
		request = NewSetHiddenProductGatewayRequest()
	}
	response = NewSetHiddenProductGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRefundRequest() (request *DescribeRefundRequest) {
	request = &DescribeRefundRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeRefund")
	return
}

func NewDescribeRefundResponse() (response *DescribeRefundResponse) {
	response = &DescribeRefundResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询提现(退款)信息
func (c *Client) DescribeRefund(request *DescribeRefundRequest) (response *DescribeRefundResponse, err error) {
	if request == nil {
		request = NewDescribeRefundRequest()
	}
	response = NewDescribeRefundResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSignatureRequest() (request *DescribeSignatureRequest) {
	request = &DescribeSignatureRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeSignature")
	return
}

func NewDescribeSignatureResponse() (response *DescribeSignatureResponse) {
	response = &DescribeSignatureResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 前端拿到签名和上传地址后把文件和签名放在form表单中使用post方法上传到私有云。
// 具体流程参照https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-UsingHTTPPOST.html
// 注意：签名过期时间设置为2小时  文件file需放在表单最后 文件类型Content-Type 只支持application开头的
func (c *Client) DescribeSignature(request *DescribeSignatureRequest) (response *DescribeSignatureResponse, err error) {
	if request == nil {
		request = NewDescribeSignatureRequest()
	}
	response = NewDescribeSignatureResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillTrendByMonthRequest() (request *DescribeBillTrendByMonthRequest) {
	request = &DescribeBillTrendByMonthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillTrendByMonth")
	return
}

func NewDescribeBillTrendByMonthResponse() (response *DescribeBillTrendByMonthResponse) {
	response = &DescribeBillTrendByMonthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账单按月消费趋势
func (c *Client) DescribeBillTrendByMonth(request *DescribeBillTrendByMonthRequest) (response *DescribeBillTrendByMonthResponse, err error) {
	if request == nil {
		request = NewDescribeBillTrendByMonthRequest()
	}
	response = NewDescribeBillTrendByMonthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeByRegionGatewayRequest() (request *DescribeBillFeeByRegionGatewayRequest) {
	request = &DescribeBillFeeByRegionGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillFeeByRegionGateway")
	return
}

func NewDescribeBillFeeByRegionGatewayResponse() (response *DescribeBillFeeByRegionGatewayResponse) {
	response = &DescribeBillFeeByRegionGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按地域汇总明细费用
func (c *Client) DescribeBillFeeByRegionGateway(request *DescribeBillFeeByRegionGatewayRequest) (response *DescribeBillFeeByRegionGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeByRegionGatewayRequest()
	}
	response = NewDescribeBillFeeByRegionGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetComprehensiveSubProductStrategyRequest() (request *GetComprehensiveSubProductStrategyRequest) {
	request = &GetComprehensiveSubProductStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetComprehensiveSubProductStrategy")
	return
}

func NewGetComprehensiveSubProductStrategyResponse() (response *GetComprehensiveSubProductStrategyResponse) {
	response = &GetComprehensiveSubProductStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询子产品（二层）策略接口
func (c *Client) GetComprehensiveSubProductStrategy(request *GetComprehensiveSubProductStrategyRequest) (response *GetComprehensiveSubProductStrategyResponse, err error) {
	if request == nil {
		request = NewGetComprehensiveSubProductStrategyRequest()
	}
	response = NewGetComprehensiveSubProductStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByRegionGatewayRequest() (request *DescribeBillSummaryByRegionGatewayRequest) {
	request = &DescribeBillSummaryByRegionGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillSummaryByRegionGateway")
	return
}

func NewDescribeBillSummaryByRegionGatewayResponse() (response *DescribeBillSummaryByRegionGatewayResponse) {
	response = &DescribeBillSummaryByRegionGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按地域汇总费用分布
func (c *Client) DescribeBillSummaryByRegionGateway(request *DescribeBillSummaryByRegionGatewayRequest) (response *DescribeBillSummaryByRegionGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByRegionGatewayRequest()
	}
	response = NewDescribeBillSummaryByRegionGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDealPriceGatewayRequest() (request *DescribeDealPriceGatewayRequest) {
	request = &DescribeDealPriceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeDealPriceGateway")
	return
}

func NewDescribeDealPriceGatewayResponse() (response *DescribeDealPriceGatewayResponse) {
	response = &DescribeDealPriceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询订单价格
func (c *Client) DescribeDealPriceGateway(request *DescribeDealPriceGatewayRequest) (response *DescribeDealPriceGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeDealPriceGatewayRequest()
	}
	response = NewDescribeDealPriceGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQuotaDownloadGatewayRequest() (request *DescribeQuotaDownloadGatewayRequest) {
	request = &DescribeQuotaDownloadGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeQuotaDownloadGateway")
	return
}

func NewDescribeQuotaDownloadGatewayResponse() (response *DescribeQuotaDownloadGatewayResponse) {
	response = &DescribeQuotaDownloadGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配额导出
func (c *Client) DescribeQuotaDownloadGateway(request *DescribeQuotaDownloadGatewayRequest) (response *DescribeQuotaDownloadGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeQuotaDownloadGatewayRequest()
	}
	response = NewDescribeQuotaDownloadGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateBSPProductDisplayCodeGatewayRequest() (request *UpdateBSPProductDisplayCodeGatewayRequest) {
	request = &UpdateBSPProductDisplayCodeGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "UpdateBSPProductDisplayCodeGateway")
	return
}

func NewUpdateBSPProductDisplayCodeGatewayResponse() (response *UpdateBSPProductDisplayCodeGatewayResponse) {
	response = &UpdateBSPProductDisplayCodeGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新插入产品定义及展示名称、产品标识
func (c *Client) UpdateBSPProductDisplayCodeGateway(request *UpdateBSPProductDisplayCodeGatewayRequest) (response *UpdateBSPProductDisplayCodeGatewayResponse, err error) {
	if request == nil {
		request = NewUpdateBSPProductDisplayCodeGatewayRequest()
	}
	response = NewUpdateBSPProductDisplayCodeGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDownloadCommonSummaryGatewayRequest() (request *DescribeBillDownloadCommonSummaryGatewayRequest) {
	request = &DescribeBillDownloadCommonSummaryGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillDownloadCommonSummaryGateway")
	return
}

func NewDescribeBillDownloadCommonSummaryGatewayResponse() (response *DescribeBillDownloadCommonSummaryGatewayResponse) {
	response = &DescribeBillDownloadCommonSummaryGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载账单通用汇总（产品+地域）
func (c *Client) DescribeBillDownloadCommonSummaryGateway(request *DescribeBillDownloadCommonSummaryGatewayRequest) (response *DescribeBillDownloadCommonSummaryGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillDownloadCommonSummaryGatewayRequest()
	}
	response = NewDescribeBillDownloadCommonSummaryGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeByProductGatewayRequest() (request *DescribeBillFeeByProductGatewayRequest) {
	request = &DescribeBillFeeByProductGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillFeeByProductGateway")
	return
}

func NewDescribeBillFeeByProductGatewayResponse() (response *DescribeBillFeeByProductGatewayResponse) {
	response = &DescribeBillFeeByProductGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按产品汇总明细费用
func (c *Client) DescribeBillFeeByProductGateway(request *DescribeBillFeeByProductGatewayRequest) (response *DescribeBillFeeByProductGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeByProductGatewayRequest()
	}
	response = NewDescribeBillFeeByProductGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadQuotasListFileGatewayRequest() (request *DescribeDownloadQuotasListFileGatewayRequest) {
	request = &DescribeDownloadQuotasListFileGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeDownloadQuotasListFileGateway")
	return
}

func NewDescribeDownloadQuotasListFileGatewayResponse() (response *DescribeDownloadQuotasListFileGatewayResponse) {
	response = &DescribeDownloadQuotasListFileGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 计费配额导出接口
func (c *Client) DescribeDownloadQuotasListFileGateway(request *DescribeDownloadQuotasListFileGatewayRequest) (response *DescribeDownloadQuotasListFileGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadQuotasListFileGatewayRequest()
	}
	response = NewDescribeDownloadQuotasListFileGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQuotaAssignLogsListGatewayRequest() (request *DescribeQuotaAssignLogsListGatewayRequest) {
	request = &DescribeQuotaAssignLogsListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeQuotaAssignLogsListGateway")
	return
}

func NewDescribeQuotaAssignLogsListGatewayResponse() (response *DescribeQuotaAssignLogsListGatewayResponse) {
	response = &DescribeQuotaAssignLogsListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账户配额转移流水
func (c *Client) DescribeQuotaAssignLogsListGateway(request *DescribeQuotaAssignLogsListGatewayRequest) (response *DescribeQuotaAssignLogsListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeQuotaAssignLogsListGatewayRequest()
	}
	response = NewDescribeQuotaAssignLogsListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewQueryRecoverUinGatewayRequest() (request *QueryRecoverUinGatewayRequest) {
	request = &QueryRecoverUinGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "QueryRecoverUinGateway")
	return
}

func NewQueryRecoverUinGatewayResponse() (response *QueryRecoverUinGatewayResponse) {
	response = &QueryRecoverUinGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口查询当前的恢复用户列表，支持分页查询，可以通过uin模糊匹配
func (c *Client) QueryRecoverUinGateway(request *QueryRecoverUinGatewayRequest) (response *QueryRecoverUinGatewayResponse, err error) {
	if request == nil {
		request = NewQueryRecoverUinGatewayRequest()
	}
	response = NewQueryRecoverUinGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPublicAccountGatewayRequest() (request *ModifyPublicAccountGatewayRequest) {
	request = &ModifyPublicAccountGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ModifyPublicAccountGateway")
	return
}

func NewModifyPublicAccountGatewayResponse() (response *ModifyPublicAccountGatewayResponse) {
	response = &ModifyPublicAccountGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加和修改公共帐户信息
func (c *Client) ModifyPublicAccountGateway(request *ModifyPublicAccountGatewayRequest) (response *ModifyPublicAccountGatewayResponse, err error) {
	if request == nil {
		request = NewModifyPublicAccountGatewayRequest()
	}
	response = NewModifyPublicAccountGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateUserStrategyGatewayRequest() (request *UpdateUserStrategyGatewayRequest) {
	request = &UpdateUserStrategyGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "UpdateUserStrategyGateway")
	return
}

func NewUpdateUserStrategyGatewayResponse() (response *UpdateUserStrategyGatewayResponse) {
	response = &UpdateUserStrategyGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【计量】更新账号订购策略
func (c *Client) UpdateUserStrategyGateway(request *UpdateUserStrategyGatewayRequest) (response *UpdateUserStrategyGatewayResponse, err error) {
	if request == nil {
		request = NewUpdateUserStrategyGatewayRequest()
	}
	response = NewUpdateUserStrategyGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewAddQuotaRequest() (request *AddQuotaRequest) {
	request = &AddQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "AddQuota")
	return
}

func NewAddQuotaResponse() (response *AddQuotaResponse) {
	response = &AddQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增,修改quota配置  支持一次传入多个配额配置。
//
// quotaKey需符合四层定义
func (c *Client) AddQuota(request *AddQuotaRequest) (response *AddQuotaResponse, err error) {
	if request == nil {
		request = NewAddQuotaRequest()
	}
	response = NewAddQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserAccountGatewayRequest() (request *DescribeUserAccountGatewayRequest) {
	request = &DescribeUserAccountGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeUserAccountGateway")
	return
}

func NewDescribeUserAccountGatewayResponse() (response *DescribeUserAccountGatewayResponse) {
	response = &DescribeUserAccountGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询单uin后付费账户信息，SubAcctid=CREDIT_FIXED 时才返回新账期NextDueDelay等详情
func (c *Client) DescribeUserAccountGateway(request *DescribeUserAccountGatewayRequest) (response *DescribeUserAccountGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeUserAccountGatewayRequest()
	}
	response = NewDescribeUserAccountGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetDiscardProductGatewayRequest() (request *GetDiscardProductGatewayRequest) {
	request = &GetDiscardProductGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetDiscardProductGateway")
	return
}

func NewGetDiscardProductGatewayResponse() (response *GetDiscardProductGatewayResponse) {
	response = &GetDiscardProductGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取废弃的商品码
func (c *Client) GetDiscardProductGateway(request *GetDiscardProductGatewayRequest) (response *GetDiscardProductGatewayResponse, err error) {
	if request == nil {
		request = NewGetDiscardProductGatewayRequest()
	}
	response = NewGetDiscardProductGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGoodsPriceGatewayRequest() (request *DescribeGoodsPriceGatewayRequest) {
	request = &DescribeGoodsPriceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeGoodsPriceGateway")
	return
}

func NewDescribeGoodsPriceGatewayResponse() (response *DescribeGoodsPriceGatewayResponse) {
	response = &DescribeGoodsPriceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询商品价格
func (c *Client) DescribeGoodsPriceGateway(request *DescribeGoodsPriceGatewayRequest) (response *DescribeGoodsPriceGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeGoodsPriceGatewayRequest()
	}
	response = NewDescribeGoodsPriceGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQuotaTypeListGatewayRequest() (request *DescribeQuotaTypeListGatewayRequest) {
	request = &DescribeQuotaTypeListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeQuotaTypeListGateway")
	return
}

func NewDescribeQuotaTypeListGatewayResponse() (response *DescribeQuotaTypeListGatewayResponse) {
	response = &DescribeQuotaTypeListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询配额类别：uin、sys
func (c *Client) DescribeQuotaTypeListGateway(request *DescribeQuotaTypeListGatewayRequest) (response *DescribeQuotaTypeListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeQuotaTypeListGatewayRequest()
	}
	response = NewDescribeQuotaTypeListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCouponDestroyRequest() (request *CouponDestroyRequest) {
	request = &CouponDestroyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "CouponDestroy")
	return
}

func NewCouponDestroyResponse() (response *CouponDestroyResponse) {
	response = &CouponDestroyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 代金券作废
func (c *Client) CouponDestroy(request *CouponDestroyRequest) (response *CouponDestroyResponse, err error) {
	if request == nil {
		request = NewCouponDestroyRequest()
	}
	response = NewCouponDestroyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPublicAccountRequest() (request *ModifyPublicAccountRequest) {
	request = &ModifyPublicAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ModifyPublicAccount")
	return
}

func NewModifyPublicAccountResponse() (response *ModifyPublicAccountResponse) {
	response = &ModifyPublicAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加和修改公共帐户信息
func (c *Client) ModifyPublicAccount(request *ModifyPublicAccountRequest) (response *ModifyPublicAccountResponse, err error) {
	if request == nil {
		request = NewModifyPublicAccountRequest()
	}
	response = NewModifyPublicAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadVoucherListFileRequest() (request *DescribeDownloadVoucherListFileRequest) {
	request = &DescribeDownloadVoucherListFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeDownloadVoucherListFile")
	return
}

func NewDescribeDownloadVoucherListFileResponse() (response *DescribeDownloadVoucherListFileResponse) {
	response = &DescribeDownloadVoucherListFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载代金券信息excel
func (c *Client) DescribeDownloadVoucherListFile(request *DescribeDownloadVoucherListFileRequest) (response *DescribeDownloadVoucherListFileResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadVoucherListFileRequest()
	}
	response = NewDescribeDownloadVoucherListFileResponse()
	err = c.Send(request, response)
	return
}

func NewAuditRemitRecordGatewayRequest() (request *AuditRemitRecordGatewayRequest) {
	request = &AuditRemitRecordGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "AuditRemitRecordGateway")
	return
}

func NewAuditRemitRecordGatewayResponse() (response *AuditRemitRecordGatewayResponse) {
	response = &AuditRemitRecordGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 审核汇款记录
func (c *Client) AuditRemitRecordGateway(request *AuditRemitRecordGatewayRequest) (response *AuditRemitRecordGatewayResponse, err error) {
	if request == nil {
		request = NewAuditRemitRecordGatewayRequest()
	}
	response = NewAuditRemitRecordGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUserStrategyGatewayRequest() (request *DeleteUserStrategyGatewayRequest) {
	request = &DeleteUserStrategyGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DeleteUserStrategyGateway")
	return
}

func NewDeleteUserStrategyGatewayResponse() (response *DeleteUserStrategyGatewayResponse) {
	response = &DeleteUserStrategyGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【计量】删除账号订购策略
func (c *Client) DeleteUserStrategyGateway(request *DeleteUserStrategyGatewayRequest) (response *DeleteUserStrategyGatewayResponse, err error) {
	if request == nil {
		request = NewDeleteUserStrategyGatewayRequest()
	}
	response = NewDeleteUserStrategyGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemitRecordGatewayRequest() (request *DescribeRemitRecordGatewayRequest) {
	request = &DescribeRemitRecordGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeRemitRecordGateway")
	return
}

func NewDescribeRemitRecordGatewayResponse() (response *DescribeRemitRecordGatewayResponse) {
	response = &DescribeRemitRecordGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端查询汇款记录列表
func (c *Client) DescribeRemitRecordGateway(request *DescribeRemitRecordGatewayRequest) (response *DescribeRemitRecordGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeRemitRecordGatewayRequest()
	}
	response = NewDescribeRemitRecordGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetQuotaRequest() (request *GetQuotaRequest) {
	request = &GetQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetQuota")
	return
}

func NewGetQuotaResponse() (response *GetQuotaResponse) {
	response = &GetQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据商品码和配额key查询配额
func (c *Client) GetQuota(request *GetQuotaRequest) (response *GetQuotaResponse, err error) {
	if request == nil {
		request = NewGetQuotaRequest()
	}
	response = NewGetQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDealListRequest() (request *DescribeDealListRequest) {
	request = &DescribeDealListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeDealList")
	return
}

func NewDescribeDealListResponse() (response *DescribeDealListResponse) {
	response = &DescribeDealListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端查询订单列表
func (c *Client) DescribeDealList(request *DescribeDealListRequest) (response *DescribeDealListResponse, err error) {
	if request == nil {
		request = NewDescribeDealListRequest()
	}
	response = NewDescribeDealListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadWithAuthRequest() (request *DescribeDownloadWithAuthRequest) {
	request = &DescribeDownloadWithAuthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeDownloadWithAuth")
	return
}

func NewDescribeDownloadWithAuthResponse() (response *DescribeDownloadWithAuthResponse) {
	response = &DescribeDownloadWithAuthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据任务ID查询文件下载地址
func (c *Client) DescribeDownloadWithAuth(request *DescribeDownloadWithAuthRequest) (response *DescribeDownloadWithAuthResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadWithAuthRequest()
	}
	response = NewDescribeDownloadWithAuthResponse()
	err = c.Send(request, response)
	return
}

func NewGetQuotaLeftRequest() (request *GetQuotaLeftRequest) {
	request = &GetQuotaLeftRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetQuotaLeft")
	return
}

func NewGetQuotaLeftResponse() (response *GetQuotaLeftResponse) {
	response = &GetQuotaLeftResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取我的剩余配额
func (c *Client) GetQuotaLeft(request *GetQuotaLeftRequest) (response *GetQuotaLeftResponse, err error) {
	if request == nil {
		request = NewGetQuotaLeftRequest()
	}
	response = NewGetQuotaLeftResponse()
	err = c.Send(request, response)
	return
}

func NewQueryBreakdownModeRequest() (request *QueryBreakdownModeRequest) {
	request = &QueryBreakdownModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "QueryBreakdownMode")
	return
}

func NewQueryBreakdownModeResponse() (response *QueryBreakdownModeResponse) {
	response = &QueryBreakdownModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口返回当前所处模式，如果是故障中，则返回白名单，临时配额开关
func (c *Client) QueryBreakdownMode(request *QueryBreakdownModeRequest) (response *QueryBreakdownModeResponse, err error) {
	if request == nil {
		request = NewQueryBreakdownModeRequest()
	}
	response = NewQueryBreakdownModeResponse()
	err = c.Send(request, response)
	return
}

func NewAddComprehensiveStrategyRequest() (request *AddComprehensiveStrategyRequest) {
	request = &AddComprehensiveStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "AddComprehensiveStrategy")
	return
}

func NewAddComprehensiveStrategyResponse() (response *AddComprehensiveStrategyResponse) {
	response = &AddComprehensiveStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 可以新增第三方和预设的产品策略
func (c *Client) AddComprehensiveStrategy(request *AddComprehensiveStrategyRequest) (response *AddComprehensiveStrategyResponse, err error) {
	if request == nil {
		request = NewAddComprehensiveStrategyRequest()
	}
	response = NewAddComprehensiveStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewDestroyRequest() (request *DestroyRequest) {
	request = &DestroyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "Destroy")
	return
}

func NewDestroyResponse() (response *DestroyResponse) {
	response = &DestroyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预付费资源销毁
func (c *Client) Destroy(request *DestroyRequest) (response *DestroyResponse, err error) {
	if request == nil {
		request = NewDestroyRequest()
	}
	response = NewDestroyResponse()
	err = c.Send(request, response)
	return
}

func NewSearchDiscountGatewayRequest() (request *SearchDiscountGatewayRequest) {
	request = &SearchDiscountGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SearchDiscountGateway")
	return
}

func NewSearchDiscountGatewayResponse() (response *SearchDiscountGatewayResponse) {
	response = &SearchDiscountGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询折扣列表
func (c *Client) SearchDiscountGateway(request *SearchDiscountGatewayRequest) (response *SearchDiscountGatewayResponse, err error) {
	if request == nil {
		request = NewSearchDiscountGatewayRequest()
	}
	response = NewSearchDiscountGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductsRequest() (request *DescribeProductsRequest) {
	request = &DescribeProductsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeProducts")
	return
}

func NewDescribeProductsResponse() (response *DescribeProductsResponse) {
	response = &DescribeProductsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取商品信息(第一层)
func (c *Client) DescribeProducts(request *DescribeProductsRequest) (response *DescribeProductsResponse, err error) {
	if request == nil {
		request = NewDescribeProductsRequest()
	}
	response = NewDescribeProductsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAcctRequest() (request *ModifyAcctRequest) {
	request = &ModifyAcctRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ModifyAcct")
	return
}

func NewModifyAcctResponse() (response *ModifyAcctResponse) {
	response = &ModifyAcctResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调账
func (c *Client) ModifyAcct(request *ModifyAcctRequest) (response *ModifyAcctResponse, err error) {
	if request == nil {
		request = NewModifyAcctRequest()
	}
	response = NewModifyAcctResponse()
	err = c.Send(request, response)
	return
}

func NewModifyContractStatusGatewayRequest() (request *ModifyContractStatusGatewayRequest) {
	request = &ModifyContractStatusGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ModifyContractStatusGateway")
	return
}

func NewModifyContractStatusGatewayResponse() (response *ModifyContractStatusGatewayResponse) {
	response = &ModifyContractStatusGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改合同状态
func (c *Client) ModifyContractStatusGateway(request *ModifyContractStatusGatewayRequest) (response *ModifyContractStatusGatewayResponse, err error) {
	if request == nil {
		request = NewModifyContractStatusGatewayRequest()
	}
	response = NewModifyContractStatusGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCouponRequest() (request *QueryCouponRequest) {
	request = &QueryCouponRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "QueryCoupon")
	return
}

func NewQueryCouponResponse() (response *QueryCouponResponse) {
	response = &QueryCouponResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询单个券详情
func (c *Client) QueryCoupon(request *QueryCouponRequest) (response *QueryCouponResponse, err error) {
	if request == nil {
		request = NewQueryCouponRequest()
	}
	response = NewQueryCouponResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceDownloadRequest() (request *DescribeResourceDownloadRequest) {
	request = &DescribeResourceDownloadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeResourceDownload")
	return
}

func NewDescribeResourceDownloadResponse() (response *DescribeResourceDownloadResponse) {
	response = &DescribeResourceDownloadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 订购实例导出
func (c *Client) DescribeResourceDownload(request *DescribeResourceDownloadRequest) (response *DescribeResourceDownloadResponse, err error) {
	if request == nil {
		request = NewDescribeResourceDownloadRequest()
	}
	response = NewDescribeResourceDownloadResponse()
	err = c.Send(request, response)
	return
}

func NewReversalResourceGatewayRequest() (request *ReversalResourceGatewayRequest) {
	request = &ReversalResourceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ReversalResourceGateway")
	return
}

func NewReversalResourceGatewayResponse() (response *ReversalResourceGatewayResponse) {
	response = &ReversalResourceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端-资源解隔离接口
func (c *Client) ReversalResourceGateway(request *ReversalResourceGatewayRequest) (response *ReversalResourceGatewayResponse, err error) {
	if request == nil {
		request = NewReversalResourceGatewayRequest()
	}
	response = NewReversalResourceGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePayOrderRequest() (request *DescribePayOrderRequest) {
	request = &DescribePayOrderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribePayOrder")
	return
}

func NewDescribePayOrderResponse() (response *DescribePayOrderResponse) {
	response = &DescribePayOrderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询订单
func (c *Client) DescribePayOrder(request *DescribePayOrderRequest) (response *DescribePayOrderResponse, err error) {
	if request == nil {
		request = NewDescribePayOrderRequest()
	}
	response = NewDescribePayOrderResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDosagesRequest() (request *DescribeDosagesRequest) {
	request = &DescribeDosagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeDosages")
	return
}

func NewDescribeDosagesResponse() (response *DescribeDosagesResponse) {
	response = &DescribeDosagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 推量统计
func (c *Client) DescribeDosages(request *DescribeDosagesRequest) (response *DescribeDosagesResponse, err error) {
	if request == nil {
		request = NewDescribeDosagesRequest()
	}
	response = NewDescribeDosagesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContractTypeListRequest() (request *DescribeContractTypeListRequest) {
	request = &DescribeContractTypeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeContractTypeList")
	return
}

func NewDescribeContractTypeListResponse() (response *DescribeContractTypeListResponse) {
	response = &DescribeContractTypeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取合同类型列表接口
func (c *Client) DescribeContractTypeList(request *DescribeContractTypeListRequest) (response *DescribeContractTypeListResponse, err error) {
	if request == nil {
		request = NewDescribeContractTypeListRequest()
	}
	response = NewDescribeContractTypeListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReconciliationListGatewayRequest() (request *DescribeReconciliationListGatewayRequest) {
	request = &DescribeReconciliationListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeReconciliationListGateway")
	return
}

func NewDescribeReconciliationListGatewayResponse() (response *DescribeReconciliationListGatewayResponse) {
	response = &DescribeReconciliationListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调账记录查询调账列表
func (c *Client) DescribeReconciliationListGateway(request *DescribeReconciliationListGatewayRequest) (response *DescribeReconciliationListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeReconciliationListGatewayRequest()
	}
	response = NewDescribeReconciliationListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetComprehensiveStrategyWebConfigRouteGatewayRequest() (request *GetComprehensiveStrategyWebConfigRouteGatewayRequest) {
	request = &GetComprehensiveStrategyWebConfigRouteGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetComprehensiveStrategyWebConfigRouteGateway")
	return
}

func NewGetComprehensiveStrategyWebConfigRouteGatewayResponse() (response *GetComprehensiveStrategyWebConfigRouteGatewayResponse) {
	response = &GetComprehensiveStrategyWebConfigRouteGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有枚举配置
func (c *Client) GetComprehensiveStrategyWebConfigRouteGateway(request *GetComprehensiveStrategyWebConfigRouteGatewayRequest) (response *GetComprehensiveStrategyWebConfigRouteGatewayResponse, err error) {
	if request == nil {
		request = NewGetComprehensiveStrategyWebConfigRouteGatewayRequest()
	}
	response = NewGetComprehensiveStrategyWebConfigRouteGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetBspUinQuotaRequest() (request *GetBspUinQuotaRequest) {
	request = &GetBspUinQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetBspUinQuota")
	return
}

func NewGetBspUinQuotaResponse() (response *GetBspUinQuotaResponse) {
	response = &GetBspUinQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端通过appid查询主账号配额全量
func (c *Client) GetBspUinQuota(request *GetBspUinQuotaRequest) (response *GetBspUinQuotaResponse, err error) {
	if request == nil {
		request = NewGetBspUinQuotaRequest()
	}
	response = NewGetBspUinQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadWithOperatorRequest() (request *DescribeDownloadWithOperatorRequest) {
	request = &DescribeDownloadWithOperatorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeDownloadWithOperator")
	return
}

func NewDescribeDownloadWithOperatorResponse() (response *DescribeDownloadWithOperatorResponse) {
	response = &DescribeDownloadWithOperatorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据任务ID查询文件下载地址
func (c *Client) DescribeDownloadWithOperator(request *DescribeDownloadWithOperatorRequest) (response *DescribeDownloadWithOperatorResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadWithOperatorRequest()
	}
	response = NewDescribeDownloadWithOperatorResponse()
	err = c.Send(request, response)
	return
}

func NewModifyContractTypeRequest() (request *ModifyContractTypeRequest) {
	request = &ModifyContractTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ModifyContractType")
	return
}

func NewModifyContractTypeResponse() (response *ModifyContractTypeResponse) {
	response = &ModifyContractTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改合同类型
func (c *Client) ModifyContractType(request *ModifyContractTypeRequest) (response *ModifyContractTypeResponse, err error) {
	if request == nil {
		request = NewModifyContractTypeRequest()
	}
	response = NewModifyContractTypeResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProductRequest() (request *UpdateProductRequest) {
	request = &UpdateProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "UpdateProduct")
	return
}

func NewUpdateProductResponse() (response *UpdateProductResponse) {
	response = &UpdateProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【计量】修改产品四层定义
func (c *Client) UpdateProduct(request *UpdateProductRequest) (response *UpdateProductResponse, err error) {
	if request == nil {
		request = NewUpdateProductRequest()
	}
	response = NewUpdateProductResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillDetailByResourceRequest() (request *DescribeBillDetailByResourceRequest) {
	request = &DescribeBillDetailByResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillDetailByResource")
	return
}

func NewDescribeBillDetailByResourceResponse() (response *DescribeBillDetailByResourceResponse) {
	response = &DescribeBillDetailByResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源花费详情
func (c *Client) DescribeBillDetailByResource(request *DescribeBillDetailByResourceRequest) (response *DescribeBillDetailByResourceResponse, err error) {
	if request == nil {
		request = NewDescribeBillDetailByResourceRequest()
	}
	response = NewDescribeBillDetailByResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceBillDetailRequest() (request *DescribeResourceBillDetailRequest) {
	request = &DescribeResourceBillDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeResourceBillDetail")
	return
}

func NewDescribeResourceBillDetailResponse() (response *DescribeResourceBillDetailResponse) {
	response = &DescribeResourceBillDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组件级别明细账单
func (c *Client) DescribeResourceBillDetail(request *DescribeResourceBillDetailRequest) (response *DescribeResourceBillDetailResponse, err error) {
	if request == nil {
		request = NewDescribeResourceBillDetailRequest()
	}
	response = NewDescribeResourceBillDetailResponse()
	err = c.Send(request, response)
	return
}

func NewGetResourceBillGatewayRequest() (request *GetResourceBillGatewayRequest) {
	request = &GetResourceBillGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetResourceBillGateway")
	return
}

func NewGetResourceBillGatewayResponse() (response *GetResourceBillGatewayResponse) {
	response = &GetResourceBillGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的账单信息
func (c *Client) GetResourceBillGateway(request *GetResourceBillGatewayRequest) (response *GetResourceBillGatewayResponse, err error) {
	if request == nil {
		request = NewGetResourceBillGatewayRequest()
	}
	response = NewGetResourceBillGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCancleContractTypeRequest() (request *CancleContractTypeRequest) {
	request = &CancleContractTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "CancleContractType")
	return
}

func NewCancleContractTypeResponse() (response *CancleContractTypeResponse) {
	response = &CancleContractTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 作废合同类型
func (c *Client) CancleContractType(request *CancleContractTypeRequest) (response *CancleContractTypeResponse, err error) {
	if request == nil {
		request = NewCancleContractTypeRequest()
	}
	response = NewCancleContractTypeResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductTreeRequest() (request *GetProductTreeRequest) {
	request = &GetProductTreeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetProductTree")
	return
}

func NewGetProductTreeResponse() (response *GetProductTreeResponse) {
	response = &GetProductTreeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过商品码，子商品码，计费项，计费细项来查询产品定义，参数都可选，如果不传则返回全部数据
func (c *Client) GetProductTree(request *GetProductTreeRequest) (response *GetProductTreeResponse, err error) {
	if request == nil {
		request = NewGetProductTreeRequest()
	}
	response = NewGetProductTreeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByRegionRequest() (request *DescribeBillSummaryByRegionRequest) {
	request = &DescribeBillSummaryByRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillSummaryByRegion")
	return
}

func NewDescribeBillSummaryByRegionResponse() (response *DescribeBillSummaryByRegionResponse) {
	response = &DescribeBillSummaryByRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按地域汇总费用分布
func (c *Client) DescribeBillSummaryByRegion(request *DescribeBillSummaryByRegionRequest) (response *DescribeBillSummaryByRegionResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByRegionRequest()
	}
	response = NewDescribeBillSummaryByRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRenewInfoRequest() (request *DescribeRenewInfoRequest) {
	request = &DescribeRenewInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeRenewInfo")
	return
}

func NewDescribeRenewInfoResponse() (response *DescribeRenewInfoResponse) {
	response = &DescribeRenewInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取续费信息
func (c *Client) DescribeRenewInfo(request *DescribeRenewInfoRequest) (response *DescribeRenewInfoResponse, err error) {
	if request == nil {
		request = NewDescribeRenewInfoRequest()
	}
	response = NewDescribeRenewInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGetRuleActionsMappingGatewayRequest() (request *GetRuleActionsMappingGatewayRequest) {
	request = &GetRuleActionsMappingGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetRuleActionsMappingGateway")
	return
}

func NewGetRuleActionsMappingGatewayResponse() (response *GetRuleActionsMappingGatewayResponse) {
	response = &GetRuleActionsMappingGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询订购策略与接口的映射关系
func (c *Client) GetRuleActionsMappingGateway(request *GetRuleActionsMappingGatewayRequest) (response *GetRuleActionsMappingGatewayResponse, err error) {
	if request == nil {
		request = NewGetRuleActionsMappingGatewayRequest()
	}
	response = NewGetRuleActionsMappingGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByTagRequest() (request *DescribeBillSummaryByTagRequest) {
	request = &DescribeBillSummaryByTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillSummaryByTag")
	return
}

func NewDescribeBillSummaryByTagResponse() (response *DescribeBillSummaryByTagResponse) {
	response = &DescribeBillSummaryByTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取标签汇总费用分布
func (c *Client) DescribeBillSummaryByTag(request *DescribeBillSummaryByTagRequest) (response *DescribeBillSummaryByTagResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByTagRequest()
	}
	response = NewDescribeBillSummaryByTagResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserAccountRequest() (request *DescribeUserAccountRequest) {
	request = &DescribeUserAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeUserAccount")
	return
}

func NewDescribeUserAccountResponse() (response *DescribeUserAccountResponse) {
	response = &DescribeUserAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询单uin后付费账户信息，SubAcctid=CREDIT_FIXED 时才返回新账期NextDueDelay等详情
func (c *Client) DescribeUserAccount(request *DescribeUserAccountRequest) (response *DescribeUserAccountResponse, err error) {
	if request == nil {
		request = NewDescribeUserAccountRequest()
	}
	response = NewDescribeUserAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTempQuotaListRequest() (request *DescribeTempQuotaListRequest) {
	request = &DescribeTempQuotaListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeTempQuotaList")
	return
}

func NewDescribeTempQuotaListResponse() (response *DescribeTempQuotaListResponse) {
	response = &DescribeTempQuotaListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页查询临时配额
func (c *Client) DescribeTempQuotaList(request *DescribeTempQuotaListRequest) (response *DescribeTempQuotaListResponse, err error) {
	if request == nil {
		request = NewDescribeTempQuotaListRequest()
	}
	response = NewDescribeTempQuotaListResponse()
	err = c.Send(request, response)
	return
}

func NewExportChangeQuotaRequest() (request *ExportChangeQuotaRequest) {
	request = &ExportChangeQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ExportChangeQuota")
	return
}

func NewExportChangeQuotaResponse() (response *ExportChangeQuotaResponse) {
	response = &ExportChangeQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端下载配额变更记录
func (c *Client) ExportChangeQuota(request *ExportChangeQuotaRequest) (response *ExportChangeQuotaResponse, err error) {
	if request == nil {
		request = NewExportChangeQuotaRequest()
	}
	response = NewExportChangeQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQuotaAssignLogsListRequest() (request *DescribeQuotaAssignLogsListRequest) {
	request = &DescribeQuotaAssignLogsListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeQuotaAssignLogsList")
	return
}

func NewDescribeQuotaAssignLogsListResponse() (response *DescribeQuotaAssignLogsListResponse) {
	response = &DescribeQuotaAssignLogsListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账户配额转移流水
func (c *Client) DescribeQuotaAssignLogsList(request *DescribeQuotaAssignLogsListRequest) (response *DescribeQuotaAssignLogsListResponse, err error) {
	if request == nil {
		request = NewDescribeQuotaAssignLogsListRequest()
	}
	response = NewDescribeQuotaAssignLogsListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTempQuotaListGatewayRequest() (request *DescribeTempQuotaListGatewayRequest) {
	request = &DescribeTempQuotaListGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeTempQuotaListGateway")
	return
}

func NewDescribeTempQuotaListGatewayResponse() (response *DescribeTempQuotaListGatewayResponse) {
	response = &DescribeTempQuotaListGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页查询临时配额
func (c *Client) DescribeTempQuotaListGateway(request *DescribeTempQuotaListGatewayRequest) (response *DescribeTempQuotaListGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeTempQuotaListGatewayRequest()
	}
	response = NewDescribeTempQuotaListGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductRelationsGatewayRequest() (request *DescribeProductRelationsGatewayRequest) {
	request = &DescribeProductRelationsGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeProductRelationsGateway")
	return
}

func NewDescribeProductRelationsGatewayResponse() (response *DescribeProductRelationsGatewayResponse) {
	response = &DescribeProductRelationsGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品定义列表
func (c *Client) DescribeProductRelationsGateway(request *DescribeProductRelationsGatewayRequest) (response *DescribeProductRelationsGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeProductRelationsGatewayRequest()
	}
	response = NewDescribeProductRelationsGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewEditBreakdownModeGatewayRequest() (request *EditBreakdownModeGatewayRequest) {
	request = &EditBreakdownModeGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "EditBreakdownModeGateway")
	return
}

func NewEditBreakdownModeGatewayResponse() (response *EditBreakdownModeGatewayResponse) {
	response = &EditBreakdownModeGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口编辑当前所处的模式，开关故障，编辑白名单，开关临时配额
func (c *Client) EditBreakdownModeGateway(request *EditBreakdownModeGatewayRequest) (response *EditBreakdownModeGatewayResponse, err error) {
	if request == nil {
		request = NewEditBreakdownModeGatewayRequest()
	}
	response = NewEditBreakdownModeGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseResourceRequest() (request *ReleaseResourceRequest) {
	request = &ReleaseResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ReleaseResource")
	return
}

func NewReleaseResourceResponse() (response *ReleaseResourceResponse) {
	response = &ReleaseResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端资源销毁
func (c *Client) ReleaseResource(request *ReleaseResourceRequest) (response *ReleaseResourceResponse, err error) {
	if request == nil {
		request = NewReleaseResourceRequest()
	}
	response = NewReleaseResourceResponse()
	err = c.Send(request, response)
	return
}

func NewSaveUserDiscountGatewayRequest() (request *SaveUserDiscountGatewayRequest) {
	request = &SaveUserDiscountGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SaveUserDiscountGateway")
	return
}

func NewSaveUserDiscountGatewayResponse() (response *SaveUserDiscountGatewayResponse) {
	response = &SaveUserDiscountGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增用户折扣
func (c *Client) SaveUserDiscountGateway(request *SaveUserDiscountGatewayRequest) (response *SaveUserDiscountGatewayResponse, err error) {
	if request == nil {
		request = NewSaveUserDiscountGatewayRequest()
	}
	response = NewSaveUserDiscountGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeTrendGatewayRequest() (request *DescribeBillFeeTrendGatewayRequest) {
	request = &DescribeBillFeeTrendGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillFeeTrendGateway")
	return
}

func NewDescribeBillFeeTrendGatewayResponse() (response *DescribeBillFeeTrendGatewayResponse) {
	response = &DescribeBillFeeTrendGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取各类汇总明细费用趋势
func (c *Client) DescribeBillFeeTrendGateway(request *DescribeBillFeeTrendGatewayRequest) (response *DescribeBillFeeTrendGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeTrendGatewayRequest()
	}
	response = NewDescribeBillFeeTrendGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetAllProductTreeGatewayRequest() (request *GetAllProductTreeGatewayRequest) {
	request = &GetAllProductTreeGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetAllProductTreeGateway")
	return
}

func NewGetAllProductTreeGatewayResponse() (response *GetAllProductTreeGatewayResponse) {
	response = &GetAllProductTreeGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过商品码，子商品码，计费项，计费细项来查询产品定义，参数都可选，如果不传则返回全部数据。不会隐藏产品
func (c *Client) GetAllProductTreeGateway(request *GetAllProductTreeGatewayRequest) (response *GetAllProductTreeGatewayResponse, err error) {
	if request == nil {
		request = NewGetAllProductTreeGatewayRequest()
	}
	response = NewGetAllProductTreeGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountWaterRequest() (request *DescribeAccountWaterRequest) {
	request = &DescribeAccountWaterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeAccountWater")
	return
}

func NewDescribeAccountWaterResponse() (response *DescribeAccountWaterResponse) {
	response = &DescribeAccountWaterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端，查询账户的还款或销账流水
func (c *Client) DescribeAccountWater(request *DescribeAccountWaterRequest) (response *DescribeAccountWaterResponse, err error) {
	if request == nil {
		request = NewDescribeAccountWaterRequest()
	}
	response = NewDescribeAccountWaterResponse()
	err = c.Send(request, response)
	return
}

func NewSaveCommonDiscountRequest() (request *SaveCommonDiscountRequest) {
	request = &SaveCommonDiscountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SaveCommonDiscount")
	return
}

func NewSaveCommonDiscountResponse() (response *SaveCommonDiscountResponse) {
	response = &SaveCommonDiscountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增官网折扣
func (c *Client) SaveCommonDiscount(request *SaveCommonDiscountRequest) (response *SaveCommonDiscountResponse, err error) {
	if request == nil {
		request = NewSaveCommonDiscountRequest()
	}
	response = NewSaveCommonDiscountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInvoiceDetailRequest() (request *DescribeInvoiceDetailRequest) {
	request = &DescribeInvoiceDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeInvoiceDetail")
	return
}

func NewDescribeInvoiceDetailResponse() (response *DescribeInvoiceDetailResponse) {
	response = &DescribeInvoiceDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端、租户端-获取发票的开票详情
func (c *Client) DescribeInvoiceDetail(request *DescribeInvoiceDetailRequest) (response *DescribeInvoiceDetailResponse, err error) {
	if request == nil {
		request = NewDescribeInvoiceDetailRequest()
	}
	response = NewDescribeInvoiceDetailResponse()
	err = c.Send(request, response)
	return
}

func NewAddUserStrategyGatewayRequest() (request *AddUserStrategyGatewayRequest) {
	request = &AddUserStrategyGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "AddUserStrategyGateway")
	return
}

func NewAddUserStrategyGatewayResponse() (response *AddUserStrategyGatewayResponse) {
	response = &AddUserStrategyGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 【计量】创建账号订购策略
func (c *Client) AddUserStrategyGateway(request *AddUserStrategyGatewayRequest) (response *AddUserStrategyGatewayResponse, err error) {
	if request == nil {
		request = NewAddUserStrategyGatewayRequest()
	}
	response = NewAddUserStrategyGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewSearchDiscountRequest() (request *SearchDiscountRequest) {
	request = &SearchDiscountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SearchDiscount")
	return
}

func NewSearchDiscountResponse() (response *SearchDiscountResponse) {
	response = &SearchDiscountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询折扣列表
func (c *Client) SearchDiscount(request *SearchDiscountRequest) (response *SearchDiscountResponse, err error) {
	if request == nil {
		request = NewSearchDiscountRequest()
	}
	response = NewSearchDiscountResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductQuotaRequest() (request *DeleteProductQuotaRequest) {
	request = &DeleteProductQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DeleteProductQuota")
	return
}

func NewDeleteProductQuotaResponse() (response *DeleteProductQuotaResponse) {
	response = &DeleteProductQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除产品配额
func (c *Client) DeleteProductQuota(request *DeleteProductQuotaRequest) (response *DeleteProductQuotaResponse, err error) {
	if request == nil {
		request = NewDeleteProductQuotaRequest()
	}
	response = NewDeleteProductQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadWithOperatorGatewayRequest() (request *DescribeDownloadWithOperatorGatewayRequest) {
	request = &DescribeDownloadWithOperatorGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeDownloadWithOperatorGateway")
	return
}

func NewDescribeDownloadWithOperatorGatewayResponse() (response *DescribeDownloadWithOperatorGatewayResponse) {
	response = &DescribeDownloadWithOperatorGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据任务ID查询文件下载地址
func (c *Client) DescribeDownloadWithOperatorGateway(request *DescribeDownloadWithOperatorGatewayRequest) (response *DescribeDownloadWithOperatorGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadWithOperatorGatewayRequest()
	}
	response = NewDescribeDownloadWithOperatorGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetComprehensiveSubProductStrategyGatewayRequest() (request *GetComprehensiveSubProductStrategyGatewayRequest) {
	request = &GetComprehensiveSubProductStrategyGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetComprehensiveSubProductStrategyGateway")
	return
}

func NewGetComprehensiveSubProductStrategyGatewayResponse() (response *GetComprehensiveSubProductStrategyGatewayResponse) {
	response = &GetComprehensiveSubProductStrategyGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询子产品（二层）策略接口
func (c *Client) GetComprehensiveSubProductStrategyGateway(request *GetComprehensiveSubProductStrategyGatewayRequest) (response *GetComprehensiveSubProductStrategyGatewayResponse, err error) {
	if request == nil {
		request = NewGetComprehensiveSubProductStrategyGatewayRequest()
	}
	response = NewGetComprehensiveSubProductStrategyGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteQuotaRequest() (request *DeleteQuotaRequest) {
	request = &DeleteQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DeleteQuota")
	return
}

func NewDeleteQuotaResponse() (response *DeleteQuotaResponse) {
	response = &DeleteQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口描述
func (c *Client) DeleteQuota(request *DeleteQuotaRequest) (response *DeleteQuotaResponse, err error) {
	if request == nil {
		request = NewDeleteQuotaRequest()
	}
	response = NewDeleteQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewGetQuotaLeftGatewayRequest() (request *GetQuotaLeftGatewayRequest) {
	request = &GetQuotaLeftGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetQuotaLeftGateway")
	return
}

func NewGetQuotaLeftGatewayResponse() (response *GetQuotaLeftGatewayResponse) {
	response = &GetQuotaLeftGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取我的剩余配额
func (c *Client) GetQuotaLeftGateway(request *GetQuotaLeftGatewayRequest) (response *GetQuotaLeftGatewayResponse, err error) {
	if request == nil {
		request = NewGetQuotaLeftGatewayRequest()
	}
	response = NewGetQuotaLeftGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDiscountGatewayRequest() (request *DeleteDiscountGatewayRequest) {
	request = &DeleteDiscountGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DeleteDiscountGateway")
	return
}

func NewDeleteDiscountGatewayResponse() (response *DeleteDiscountGatewayResponse) {
	response = &DeleteDiscountGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除折扣
func (c *Client) DeleteDiscountGateway(request *DeleteDiscountGatewayRequest) (response *DeleteDiscountGatewayResponse, err error) {
	if request == nil {
		request = NewDeleteDiscountGatewayRequest()
	}
	response = NewDeleteDiscountGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByResourceGatewayRequest() (request *DescribeBillSummaryByResourceGatewayRequest) {
	request = &DescribeBillSummaryByResourceGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillSummaryByResourceGateway")
	return
}

func NewDescribeBillSummaryByResourceGatewayResponse() (response *DescribeBillSummaryByResourceGatewayResponse) {
	response = &DescribeBillSummaryByResourceGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按资源汇总数据
func (c *Client) DescribeBillSummaryByResourceGateway(request *DescribeBillSummaryByResourceGatewayRequest) (response *DescribeBillSummaryByResourceGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByResourceGatewayRequest()
	}
	response = NewDescribeBillSummaryByResourceGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewExpireDiscountGatewayRequest() (request *ExpireDiscountGatewayRequest) {
	request = &ExpireDiscountGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ExpireDiscountGateway")
	return
}

func NewExpireDiscountGatewayResponse() (response *ExpireDiscountGatewayResponse) {
	response = &ExpireDiscountGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 废弃已经生效的折扣
func (c *Client) ExpireDiscountGateway(request *ExpireDiscountGatewayRequest) (response *ExpireDiscountGatewayResponse, err error) {
	if request == nil {
		request = NewExpireDiscountGatewayRequest()
	}
	response = NewExpireDiscountGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDebateListRequest() (request *DescribeDebateListRequest) {
	request = &DescribeDebateListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeDebateList")
	return
}

func NewDescribeDebateListResponse() (response *DescribeDebateListResponse) {
	response = &DescribeDebateListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 管理端查询用户账单是否欠费和欠费月份
func (c *Client) DescribeDebateList(request *DescribeDebateListRequest) (response *DescribeDebateListResponse, err error) {
	if request == nil {
		request = NewDescribeDebateListRequest()
	}
	response = NewDescribeDebateListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductPricesRequest() (request *DescribeProductPricesRequest) {
	request = &DescribeProductPricesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeProductPrices")
	return
}

func NewDescribeProductPricesResponse() (response *DescribeProductPricesResponse) {
	response = &DescribeProductPricesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取价格策略列表
func (c *Client) DescribeProductPrices(request *DescribeProductPricesRequest) (response *DescribeProductPricesResponse, err error) {
	if request == nil {
		request = NewDescribeProductPricesRequest()
	}
	response = NewDescribeProductPricesResponse()
	err = c.Send(request, response)
	return
}

func NewGetAllProductGatewayRequest() (request *GetAllProductGatewayRequest) {
	request = &GetAllProductGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetAllProductGateway")
	return
}

func NewGetAllProductGatewayResponse() (response *GetAllProductGatewayResponse) {
	response = &GetAllProductGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取全量产品一层（产品模块）
func (c *Client) GetAllProductGateway(request *GetAllProductGatewayRequest) (response *GetAllProductGatewayResponse, err error) {
	if request == nil {
		request = NewGetAllProductGatewayRequest()
	}
	response = NewGetAllProductGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceRequest() (request *ListResourceRequest) {
	request = &ListResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ListResource")
	return
}

func NewListResourceResponse() (response *ListResourceResponse) {
	response = &ListResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端根据Uin，ResourceId查询用户资源
func (c *Client) ListResource(request *ListResourceRequest) (response *ListResourceResponse, err error) {
	if request == nil {
		request = NewListResourceRequest()
	}
	response = NewListResourceResponse()
	err = c.Send(request, response)
	return
}

func NewQueryBSPProductGatewayRequest() (request *QueryBSPProductGatewayRequest) {
	request = &QueryBSPProductGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "QueryBSPProductGateway")
	return
}

func NewQueryBSPProductGatewayResponse() (response *QueryBSPProductGatewayResponse) {
	response = &QueryBSPProductGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询一层产品定义及展示名称、产品标识
func (c *Client) QueryBSPProductGateway(request *QueryBSPProductGatewayRequest) (response *QueryBSPProductGatewayResponse, err error) {
	if request == nil {
		request = NewQueryBSPProductGatewayRequest()
	}
	response = NewQueryBSPProductGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillSummaryByResourceRequest() (request *DescribeBillSummaryByResourceRequest) {
	request = &DescribeBillSummaryByResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillSummaryByResource")
	return
}

func NewDescribeBillSummaryByResourceResponse() (response *DescribeBillSummaryByResourceResponse) {
	response = &DescribeBillSummaryByResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按资源汇总数据
func (c *Client) DescribeBillSummaryByResource(request *DescribeBillSummaryByResourceRequest) (response *DescribeBillSummaryByResourceResponse, err error) {
	if request == nil {
		request = NewDescribeBillSummaryByResourceRequest()
	}
	response = NewDescribeBillSummaryByResourceResponse()
	err = c.Send(request, response)
	return
}

func NewQueryPublicAccountGatewayRequest() (request *QueryPublicAccountGatewayRequest) {
	request = &QueryPublicAccountGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "QueryPublicAccountGateway")
	return
}

func NewQueryPublicAccountGatewayResponse() (response *QueryPublicAccountGatewayResponse) {
	response = &QueryPublicAccountGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询公共帐户信息
func (c *Client) QueryPublicAccountGateway(request *QueryPublicAccountGatewayRequest) (response *QueryPublicAccountGatewayResponse, err error) {
	if request == nil {
		request = NewQueryPublicAccountGatewayRequest()
	}
	response = NewQueryPublicAccountGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewSaveQuotaRequest() (request *SaveQuotaRequest) {
	request = &SaveQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "SaveQuota")
	return
}

func NewSaveQuotaResponse() (response *SaveQuotaResponse) {
	response = &SaveQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品配额新增、编辑
func (c *Client) SaveQuota(request *SaveQuotaRequest) (response *SaveQuotaResponse, err error) {
	if request == nil {
		request = NewSaveQuotaRequest()
	}
	response = NewSaveQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillFeeByProductRequest() (request *DescribeBillFeeByProductRequest) {
	request = &DescribeBillFeeByProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeBillFeeByProduct")
	return
}

func NewDescribeBillFeeByProductResponse() (response *DescribeBillFeeByProductResponse) {
	response = &DescribeBillFeeByProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取按产品汇总明细费用
func (c *Client) DescribeBillFeeByProduct(request *DescribeBillFeeByProductRequest) (response *DescribeBillFeeByProductResponse, err error) {
	if request == nil {
		request = NewDescribeBillFeeByProductRequest()
	}
	response = NewDescribeBillFeeByProductResponse()
	err = c.Send(request, response)
	return
}

func NewAddQuotaGatewayRequest() (request *AddQuotaGatewayRequest) {
	request = &AddQuotaGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "AddQuotaGateway")
	return
}

func NewAddQuotaGatewayResponse() (response *AddQuotaGatewayResponse) {
	response = &AddQuotaGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增,修改quota配置  支持一次传入多个配额配置。
//
// quotaKey需符合四层定义
func (c *Client) AddQuotaGateway(request *AddQuotaGatewayRequest) (response *AddQuotaGatewayResponse, err error) {
	if request == nil {
		request = NewAddQuotaGatewayRequest()
	}
	response = NewAddQuotaGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceDownloadGatewayRequest() (request *DescribeResourceDownloadGatewayRequest) {
	request = &DescribeResourceDownloadGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeResourceDownloadGateway")
	return
}

func NewDescribeResourceDownloadGatewayResponse() (response *DescribeResourceDownloadGatewayResponse) {
	response = &DescribeResourceDownloadGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 订购实例导出
func (c *Client) DescribeResourceDownloadGateway(request *DescribeResourceDownloadGatewayRequest) (response *DescribeResourceDownloadGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeResourceDownloadGatewayRequest()
	}
	response = NewDescribeResourceDownloadGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceSummaryGatewayRequest() (request *ListResourceSummaryGatewayRequest) {
	request = &ListResourceSummaryGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "ListResourceSummaryGateway")
	return
}

func NewListResourceSummaryGatewayResponse() (response *ListResourceSummaryGatewayResponse) {
	response = &ListResourceSummaryGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资源详情
func (c *Client) ListResourceSummaryGateway(request *ListResourceSummaryGatewayRequest) (response *ListResourceSummaryGatewayResponse, err error) {
	if request == nil {
		request = NewListResourceSummaryGatewayRequest()
	}
	response = NewListResourceSummaryGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewGetComprehensiveSubProdctStrategyRequest() (request *GetComprehensiveSubProdctStrategyRequest) {
	request = &GetComprehensiveSubProdctStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "GetComprehensiveSubProdctStrategy")
	return
}

func NewGetComprehensiveSubProdctStrategyResponse() (response *GetComprehensiveSubProdctStrategyResponse) {
	response = &GetComprehensiveSubProdctStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询子产品（二层）策略接口
func (c *Client) GetComprehensiveSubProdctStrategy(request *GetComprehensiveSubProdctStrategyRequest) (response *GetComprehensiveSubProdctStrategyResponse, err error) {
	if request == nil {
		request = NewGetComprehensiveSubProdctStrategyRequest()
	}
	response = NewGetComprehensiveSubProdctStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadDosagesRequest() (request *DownloadDosagesRequest) {
	request = &DownloadDosagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DownloadDosages")
	return
}

func NewDownloadDosagesResponse() (response *DownloadDosagesResponse) {
	response = &DownloadDosagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用量统计下载
func (c *Client) DownloadDosages(request *DownloadDosagesRequest) (response *DownloadDosagesResponse, err error) {
	if request == nil {
		request = NewDownloadDosagesRequest()
	}
	response = NewDownloadDosagesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsolidatedBillDownloadUrlGatewayRequest() (request *DescribeConsolidatedBillDownloadUrlGatewayRequest) {
	request = &DescribeConsolidatedBillDownloadUrlGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opbill", APIVersion, "DescribeConsolidatedBillDownloadUrlGateway")
	return
}

func NewDescribeConsolidatedBillDownloadUrlGatewayResponse() (response *DescribeConsolidatedBillDownloadUrlGatewayResponse) {
	response = &DescribeConsolidatedBillDownloadUrlGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账单包下载地址
func (c *Client) DescribeConsolidatedBillDownloadUrlGateway(request *DescribeConsolidatedBillDownloadUrlGatewayRequest) (response *DescribeConsolidatedBillDownloadUrlGatewayResponse, err error) {
	if request == nil {
		request = NewDescribeConsolidatedBillDownloadUrlGatewayRequest()
	}
	response = NewDescribeConsolidatedBillDownloadUrlGatewayResponse()
	err = c.Send(request, response)
	return
}
