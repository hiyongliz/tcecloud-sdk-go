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

package v20201111

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-11-11"

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

func NewDeleteCategoryRequest() (request *DeleteCategoryRequest) {
	request = &DeleteCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "DeleteCategory")
	return
}

func NewDeleteCategoryResponse() (response *DeleteCategoryResponse) {
	response = &DeleteCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 目录分类删除，如果存在关联模板， 则不允许删除
func (c *Client) DeleteCategory(request *DeleteCategoryRequest) (response *DeleteCategoryResponse, err error) {
	if request == nil {
		request = NewDeleteCategoryRequest()
	}
	response = NewDeleteCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSubscribeRequest() (request *UpdateSubscribeRequest) {
	request = &UpdateSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "UpdateSubscribe")
	return
}

func NewUpdateSubscribeResponse() (response *UpdateSubscribeResponse) {
	response = &UpdateSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新订阅关系
func (c *Client) UpdateSubscribe(request *UpdateSubscribeRequest) (response *UpdateSubscribeResponse, err error) {
	if request == nil {
		request = NewUpdateSubscribeRequest()
	}
	response = NewUpdateSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSubscribeRequest() (request *CreateSubscribeRequest) {
	request = &CreateSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "CreateSubscribe")
	return
}

func NewCreateSubscribeResponse() (response *CreateSubscribeResponse) {
	response = &CreateSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建订阅关系
func (c *Client) CreateSubscribe(request *CreateSubscribeRequest) (response *CreateSubscribeResponse, err error) {
	if request == nil {
		request = NewCreateSubscribeRequest()
	}
	response = NewCreateSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewListMetricSourceRequest() (request *ListMetricSourceRequest) {
	request = &ListMetricSourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "ListMetricSource")
	return
}

func NewListMetricSourceResponse() (response *ListMetricSourceResponse) {
	response = &ListMetricSourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指标列表 & 标签
func (c *Client) ListMetricSource(request *ListMetricSourceRequest) (response *ListMetricSourceResponse, err error) {
	if request == nil {
		request = NewListMetricSourceRequest()
	}
	response = NewListMetricSourceResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateTemplateRequest() (request *UpdateTemplateRequest) {
	request = &UpdateTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "UpdateTemplate")
	return
}

func NewUpdateTemplateResponse() (response *UpdateTemplateResponse) {
	response = &UpdateTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 模版列表页(删除 + 新增)
func (c *Client) UpdateTemplate(request *UpdateTemplateRequest) (response *UpdateTemplateResponse, err error) {
	if request == nil {
		request = NewUpdateTemplateRequest()
	}
	response = NewUpdateTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTemplateRequest() (request *CreateTemplateRequest) {
	request = &CreateTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "CreateTemplate")
	return
}

func NewCreateTemplateResponse() (response *CreateTemplateResponse) {
	response = &CreateTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增模版
func (c *Client) CreateTemplate(request *CreateTemplateRequest) (response *CreateTemplateResponse, err error) {
	if request == nil {
		request = NewCreateTemplateRequest()
	}
	response = NewCreateTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewRenderTemplateRequest() (request *RenderTemplateRequest) {
	request = &RenderTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "RenderTemplate")
	return
}

func NewRenderTemplateResponse() (response *RenderTemplateResponse) {
	response = &RenderTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预览模版
func (c *Client) RenderTemplate(request *RenderTemplateRequest) (response *RenderTemplateResponse, err error) {
	if request == nil {
		request = NewRenderTemplateRequest()
	}
	response = NewRenderTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewListTemplateRequest() (request *ListTemplateRequest) {
	request = &ListTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "ListTemplate")
	return
}

func NewListTemplateResponse() (response *ListTemplateResponse) {
	response = &ListTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 模版列表页
func (c *Client) ListTemplate(request *ListTemplateRequest) (response *ListTemplateResponse, err error) {
	if request == nil {
		request = NewListTemplateRequest()
	}
	response = NewListTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewAddCatetoryRequest() (request *AddCatetoryRequest) {
	request = &AddCatetoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "AddCatetory")
	return
}

func NewAddCatetoryResponse() (response *AddCatetoryResponse) {
	response = &AddCatetoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增分类
func (c *Client) AddCatetory(request *AddCatetoryRequest) (response *AddCatetoryResponse, err error) {
	if request == nil {
		request = NewAddCatetoryRequest()
	}
	response = NewAddCatetoryResponse()
	err = c.Send(request, response)
	return
}

func NewGetReportChartUrlRequest() (request *GetReportChartUrlRequest) {
	request = &GetReportChartUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "GetReportChartUrl")
	return
}

func NewGetReportChartUrlResponse() (response *GetReportChartUrlResponse) {
	response = &GetReportChartUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载报表可视化pdf对应 url
func (c *Client) GetReportChartUrl(request *GetReportChartUrlRequest) (response *GetReportChartUrlResponse, err error) {
	if request == nil {
		request = NewGetReportChartUrlRequest()
	}
	response = NewGetReportChartUrlResponse()
	err = c.Send(request, response)
	return
}

func NewQueryTemplateRequest() (request *QueryTemplateRequest) {
	request = &QueryTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "QueryTemplate")
	return
}

func NewQueryTemplateResponse() (response *QueryTemplateResponse) {
	response = &QueryTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询模版详情
func (c *Client) QueryTemplate(request *QueryTemplateRequest) (response *QueryTemplateResponse, err error) {
	if request == nil {
		request = NewQueryTemplateRequest()
	}
	response = NewQueryTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewRenameCatetoryRequest() (request *RenameCatetoryRequest) {
	request = &RenameCatetoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "RenameCatetory")
	return
}

func NewRenameCatetoryResponse() (response *RenameCatetoryResponse) {
	response = &RenameCatetoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重命名分类
func (c *Client) RenameCatetory(request *RenameCatetoryRequest) (response *RenameCatetoryResponse, err error) {
	if request == nil {
		request = NewRenameCatetoryRequest()
	}
	response = NewRenameCatetoryResponse()
	err = c.Send(request, response)
	return
}

func NewListInstanceCategoryRequest() (request *ListInstanceCategoryRequest) {
	request = &ListInstanceCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "ListInstanceCategory")
	return
}

func NewListInstanceCategoryResponse() (response *ListInstanceCategoryResponse) {
	response = &ListInstanceCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取FT资源分类
func (c *Client) ListInstanceCategory(request *ListInstanceCategoryRequest) (response *ListInstanceCategoryResponse, err error) {
	if request == nil {
		request = NewListInstanceCategoryRequest()
	}
	response = NewListInstanceCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewListMetricCategoryRequest() (request *ListMetricCategoryRequest) {
	request = &ListMetricCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "ListMetricCategory")
	return
}

func NewListMetricCategoryResponse() (response *ListMetricCategoryResponse) {
	response = &ListMetricCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取监控分类
func (c *Client) ListMetricCategory(request *ListMetricCategoryRequest) (response *ListMetricCategoryResponse, err error) {
	if request == nil {
		request = NewListMetricCategoryRequest()
	}
	response = NewListMetricCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewQuerySubscribeRequest() (request *QuerySubscribeRequest) {
	request = &QuerySubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "QuerySubscribe")
	return
}

func NewQuerySubscribeResponse() (response *QuerySubscribeResponse) {
	response = &QuerySubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询订阅关系详情
func (c *Client) QuerySubscribe(request *QuerySubscribeRequest) (response *QuerySubscribeResponse, err error) {
	if request == nil {
		request = NewQuerySubscribeRequest()
	}
	response = NewQuerySubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewQueryReportRequest() (request *QueryReportRequest) {
	request = &QueryReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "QueryReport")
	return
}

func NewQueryReportResponse() (response *QueryReportResponse) {
	response = &QueryReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询详情页
func (c *Client) QueryReport(request *QueryReportRequest) (response *QueryReportResponse, err error) {
	if request == nil {
		request = NewQueryReportRequest()
	}
	response = NewQueryReportResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTemplateRequest() (request *DeleteTemplateRequest) {
	request = &DeleteTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "DeleteTemplate")
	return
}

func NewDeleteTemplateResponse() (response *DeleteTemplateResponse) {
	response = &DeleteTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除模版详情
func (c *Client) DeleteTemplate(request *DeleteTemplateRequest) (response *DeleteTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteTemplateRequest()
	}
	response = NewDeleteTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewListInstanceSourceRequest() (request *ListInstanceSourceRequest) {
	request = &ListInstanceSourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "ListInstanceSource")
	return
}

func NewListInstanceSourceResponse() (response *ListInstanceSourceResponse) {
	response = &ListInstanceSourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取FT资源字段元信息
func (c *Client) ListInstanceSource(request *ListInstanceSourceRequest) (response *ListInstanceSourceResponse, err error) {
	if request == nil {
		request = NewListInstanceSourceRequest()
	}
	response = NewListInstanceSourceResponse()
	err = c.Send(request, response)
	return
}

func NewListCatetoryRequest() (request *ListCatetoryRequest) {
	request = &ListCatetoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "ListCatetory")
	return
}

func NewListCatetoryResponse() (response *ListCatetoryResponse) {
	response = &ListCatetoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取分类列表
func (c *Client) ListCatetory(request *ListCatetoryRequest) (response *ListCatetoryResponse, err error) {
	if request == nil {
		request = NewListCatetoryRequest()
	}
	response = NewListCatetoryResponse()
	err = c.Send(request, response)
	return
}

func NewListReportRequest() (request *ListReportRequest) {
	request = &ListReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("report", APIVersion, "ListReport")
	return
}

func NewListReportResponse() (response *ListReportResponse) {
	response = &ListReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询列表页
func (c *Client) ListReport(request *ListReportRequest) (response *ListReportResponse, err error) {
	if request == nil {
		request = NewListReportRequest()
	}
	response = NewListReportResponse()
	err = c.Send(request, response)
	return
}
