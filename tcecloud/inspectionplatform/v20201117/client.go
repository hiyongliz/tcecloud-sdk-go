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

package v20201117

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-11-17"

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

func NewGetReportHistoryRequest() (request *GetReportHistoryRequest) {
	request = &GetReportHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "GetReportHistory")
	return
}

func NewGetReportHistoryResponse() (response *GetReportHistoryResponse) {
	response = &GetReportHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取巡检报告历史
func (c *Client) GetReportHistory(request *GetReportHistoryRequest) (response *GetReportHistoryResponse, err error) {
	if request == nil {
		request = NewGetReportHistoryRequest()
	}
	response = NewGetReportHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewRunCronJobRequest() (request *RunCronJobRequest) {
	request = &RunCronJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "RunCronJob")
	return
}

func NewRunCronJobResponse() (response *RunCronJobResponse) {
	response = &RunCronJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 立即执行巡检任务
func (c *Client) RunCronJob(request *RunCronJobRequest) (response *RunCronJobResponse, err error) {
	if request == nil {
		request = NewRunCronJobRequest()
	}
	response = NewRunCronJobResponse()
	err = c.Send(request, response)
	return
}

func NewListCMDBProductRequest() (request *ListCMDBProductRequest) {
	request = &ListCMDBProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "ListCMDBProduct")
	return
}

func NewListCMDBProductResponse() (response *ListCMDBProductResponse) {
	response = &ListCMDBProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列表CMDB产品名
func (c *Client) ListCMDBProduct(request *ListCMDBProductRequest) (response *ListCMDBProductResponse, err error) {
	if request == nil {
		request = NewListCMDBProductRequest()
	}
	response = NewListCMDBProductResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCronJobRequest() (request *CreateCronJobRequest) {
	request = &CreateCronJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "CreateCronJob")
	return
}

func NewCreateCronJobResponse() (response *CreateCronJobResponse) {
	response = &CreateCronJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建巡检计划
func (c *Client) CreateCronJob(request *CreateCronJobRequest) (response *CreateCronJobResponse, err error) {
	if request == nil {
		request = NewCreateCronJobRequest()
	}
	response = NewCreateCronJobResponse()
	err = c.Send(request, response)
	return
}

func NewStartCronJobRequest() (request *StartCronJobRequest) {
	request = &StartCronJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "StartCronJob")
	return
}

func NewStartCronJobResponse() (response *StartCronJobResponse) {
	response = &StartCronJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用巡检计划
func (c *Client) StartCronJob(request *StartCronJobRequest) (response *StartCronJobResponse, err error) {
	if request == nil {
		request = NewStartCronJobRequest()
	}
	response = NewStartCronJobResponse()
	err = c.Send(request, response)
	return
}

func NewCheckImageURLRequest() (request *CheckImageURLRequest) {
	request = &CheckImageURLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "CheckImageURL")
	return
}

func NewCheckImageURLResponse() (response *CheckImageURLResponse) {
	response = &CheckImageURLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查 URL 是否可达
func (c *Client) CheckImageURL(request *CheckImageURLRequest) (response *CheckImageURLResponse, err error) {
	if request == nil {
		request = NewCheckImageURLRequest()
	}
	response = NewCheckImageURLResponse()
	err = c.Send(request, response)
	return
}

func NewStopJobRequest() (request *StopJobRequest) {
	request = &StopJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "StopJob")
	return
}

func NewStopJobResponse() (response *StopJobResponse) {
	response = &StopJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停止巡检任务
func (c *Client) StopJob(request *StopJobRequest) (response *StopJobResponse, err error) {
	if request == nil {
		request = NewStopJobRequest()
	}
	response = NewStopJobResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteInspectionItemRequest() (request *DeleteInspectionItemRequest) {
	request = &DeleteInspectionItemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "DeleteInspectionItem")
	return
}

func NewDeleteInspectionItemResponse() (response *DeleteInspectionItemResponse) {
	response = &DeleteInspectionItemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除巡检项
func (c *Client) DeleteInspectionItem(request *DeleteInspectionItemRequest) (response *DeleteInspectionItemResponse, err error) {
	if request == nil {
		request = NewDeleteInspectionItemRequest()
	}
	response = NewDeleteInspectionItemResponse()
	err = c.Send(request, response)
	return
}

func NewListJobRequest() (request *ListJobRequest) {
	request = &ListJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "ListJob")
	return
}

func NewListJobResponse() (response *ListJobResponse) {
	response = &ListJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列表巡检任务
func (c *Client) ListJob(request *ListJobRequest) (response *ListJobResponse, err error) {
	if request == nil {
		request = NewListJobRequest()
	}
	response = NewListJobResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCronJobRequest() (request *DeleteCronJobRequest) {
	request = &DeleteCronJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "DeleteCronJob")
	return
}

func NewDeleteCronJobResponse() (response *DeleteCronJobResponse) {
	response = &DeleteCronJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除巡检任务
func (c *Client) DeleteCronJob(request *DeleteCronJobRequest) (response *DeleteCronJobResponse, err error) {
	if request == nil {
		request = NewDeleteCronJobRequest()
	}
	response = NewDeleteCronJobResponse()
	err = c.Send(request, response)
	return
}

func NewListCronJobRequest() (request *ListCronJobRequest) {
	request = &ListCronJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "ListCronJob")
	return
}

func NewListCronJobResponse() (response *ListCronJobResponse) {
	response = &ListCronJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列表巡检计划
func (c *Client) ListCronJob(request *ListCronJobRequest) (response *ListCronJobResponse, err error) {
	if request == nil {
		request = NewListCronJobRequest()
	}
	response = NewListCronJobResponse()
	err = c.Send(request, response)
	return
}

func NewListReportRequest() (request *ListReportRequest) {
	request = &ListReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "ListReport")
	return
}

func NewListReportResponse() (response *ListReportResponse) {
	response = &ListReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列表巡检报告
func (c *Client) ListReport(request *ListReportRequest) (response *ListReportResponse, err error) {
	if request == nil {
		request = NewListReportRequest()
	}
	response = NewListReportResponse()
	err = c.Send(request, response)
	return
}

func NewListCustomProductRequest() (request *ListCustomProductRequest) {
	request = &ListCustomProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "ListCustomProduct")
	return
}

func NewListCustomProductResponse() (response *ListCustomProductResponse) {
	response = &ListCustomProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取自定义产品列表
func (c *Client) ListCustomProduct(request *ListCustomProductRequest) (response *ListCustomProductResponse, err error) {
	if request == nil {
		request = NewListCustomProductRequest()
	}
	response = NewListCustomProductResponse()
	err = c.Send(request, response)
	return
}

func NewListCronJobHistoryRequest() (request *ListCronJobHistoryRequest) {
	request = &ListCronJobHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "ListCronJobHistory")
	return
}

func NewListCronJobHistoryResponse() (response *ListCronJobHistoryResponse) {
	response = &ListCronJobHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列表巡检计划历史
func (c *Client) ListCronJobHistory(request *ListCronJobHistoryRequest) (response *ListCronJobHistoryResponse, err error) {
	if request == nil {
		request = NewListCronJobHistoryRequest()
	}
	response = NewListCronJobHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewGetInspectionItemRequest() (request *GetInspectionItemRequest) {
	request = &GetInspectionItemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "GetInspectionItem")
	return
}

func NewGetInspectionItemResponse() (response *GetInspectionItemResponse) {
	response = &GetInspectionItemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取巡检项
func (c *Client) GetInspectionItem(request *GetInspectionItemRequest) (response *GetInspectionItemResponse, err error) {
	if request == nil {
		request = NewGetInspectionItemRequest()
	}
	response = NewGetInspectionItemResponse()
	err = c.Send(request, response)
	return
}

func NewGetReportRequest() (request *GetReportRequest) {
	request = &GetReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "GetReport")
	return
}

func NewGetReportResponse() (response *GetReportResponse) {
	response = &GetReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取巡检报告
func (c *Client) GetReport(request *GetReportRequest) (response *GetReportResponse, err error) {
	if request == nil {
		request = NewGetReportRequest()
	}
	response = NewGetReportResponse()
	err = c.Send(request, response)
	return
}

func NewGetKeepTimeRequest() (request *GetKeepTimeRequest) {
	request = &GetKeepTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "GetKeepTime")
	return
}

func NewGetKeepTimeResponse() (response *GetKeepTimeResponse) {
	response = &GetKeepTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取巡检报告、巡检任务保留时间
func (c *Client) GetKeepTime(request *GetKeepTimeRequest) (response *GetKeepTimeResponse, err error) {
	if request == nil {
		request = NewGetKeepTimeRequest()
	}
	response = NewGetKeepTimeResponse()
	err = c.Send(request, response)
	return
}

func NewListJobHistoryRequest() (request *ListJobHistoryRequest) {
	request = &ListJobHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "ListJobHistory")
	return
}

func NewListJobHistoryResponse() (response *ListJobHistoryResponse) {
	response = &ListJobHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列表巡检任务历史
func (c *Client) ListJobHistory(request *ListJobHistoryRequest) (response *ListJobHistoryResponse, err error) {
	if request == nil {
		request = NewListJobHistoryRequest()
	}
	response = NewListJobHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewStopCronJobRequest() (request *StopCronJobRequest) {
	request = &StopCronJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "StopCronJob")
	return
}

func NewStopCronJobResponse() (response *StopCronJobResponse) {
	response = &StopCronJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停用巡检任务
func (c *Client) StopCronJob(request *StopCronJobRequest) (response *StopCronJobResponse, err error) {
	if request == nil {
		request = NewStopCronJobRequest()
	}
	response = NewStopCronJobResponse()
	err = c.Send(request, response)
	return
}

func NewGetCronJobRequest() (request *GetCronJobRequest) {
	request = &GetCronJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "GetCronJob")
	return
}

func NewGetCronJobResponse() (response *GetCronJobResponse) {
	response = &GetCronJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取巡检任务
func (c *Client) GetCronJob(request *GetCronJobRequest) (response *GetCronJobResponse, err error) {
	if request == nil {
		request = NewGetCronJobRequest()
	}
	response = NewGetCronJobResponse()
	err = c.Send(request, response)
	return
}

func NewListOperationRecordRequest() (request *ListOperationRecordRequest) {
	request = &ListOperationRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "ListOperationRecord")
	return
}

func NewListOperationRecordResponse() (response *ListOperationRecordResponse) {
	response = &ListOperationRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列表操作记录
func (c *Client) ListOperationRecord(request *ListOperationRecordRequest) (response *ListOperationRecordResponse, err error) {
	if request == nil {
		request = NewListOperationRecordRequest()
	}
	response = NewListOperationRecordResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateInspectionItemRequest() (request *UpdateInspectionItemRequest) {
	request = &UpdateInspectionItemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "UpdateInspectionItem")
	return
}

func NewUpdateInspectionItemResponse() (response *UpdateInspectionItemResponse) {
	response = &UpdateInspectionItemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新巡检项
func (c *Client) UpdateInspectionItem(request *UpdateInspectionItemRequest) (response *UpdateInspectionItemResponse, err error) {
	if request == nil {
		request = NewUpdateInspectionItemRequest()
	}
	response = NewUpdateInspectionItemResponse()
	err = c.Send(request, response)
	return
}

func NewCreateInspectionItemRequest() (request *CreateInspectionItemRequest) {
	request = &CreateInspectionItemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "CreateInspectionItem")
	return
}

func NewCreateInspectionItemResponse() (response *CreateInspectionItemResponse) {
	response = &CreateInspectionItemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建巡检项
func (c *Client) CreateInspectionItem(request *CreateInspectionItemRequest) (response *CreateInspectionItemResponse, err error) {
	if request == nil {
		request = NewCreateInspectionItemRequest()
	}
	response = NewCreateInspectionItemResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCronJobRequest() (request *UpdateCronJobRequest) {
	request = &UpdateCronJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "UpdateCronJob")
	return
}

func NewUpdateCronJobResponse() (response *UpdateCronJobResponse) {
	response = &UpdateCronJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新巡检计划
func (c *Client) UpdateCronJob(request *UpdateCronJobRequest) (response *UpdateCronJobResponse, err error) {
	if request == nil {
		request = NewUpdateCronJobRequest()
	}
	response = NewUpdateCronJobResponse()
	err = c.Send(request, response)
	return
}

func NewConfigureKeepTimeRequest() (request *ConfigureKeepTimeRequest) {
	request = &ConfigureKeepTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "ConfigureKeepTime")
	return
}

func NewConfigureKeepTimeResponse() (response *ConfigureKeepTimeResponse) {
	response = &ConfigureKeepTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置巡检报告、巡检任务保留时间
func (c *Client) ConfigureKeepTime(request *ConfigureKeepTimeRequest) (response *ConfigureKeepTimeResponse, err error) {
	if request == nil {
		request = NewConfigureKeepTimeRequest()
	}
	response = NewConfigureKeepTimeResponse()
	err = c.Send(request, response)
	return
}

func NewCheckRequest() (request *CheckRequest) {
	request = &CheckRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "Check")
	return
}

func NewCheckResponse() (response *CheckResponse) {
	response = &CheckResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查目标是否可达
func (c *Client) Check(request *CheckRequest) (response *CheckResponse, err error) {
	if request == nil {
		request = NewCheckRequest()
	}
	response = NewCheckResponse()
	err = c.Send(request, response)
	return
}

func NewListInspectionItemRequest() (request *ListInspectionItemRequest) {
	request = &ListInspectionItemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "ListInspectionItem")
	return
}

func NewListInspectionItemResponse() (response *ListInspectionItemResponse) {
	response = &ListInspectionItemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列表巡检项
func (c *Client) ListInspectionItem(request *ListInspectionItemRequest) (response *ListInspectionItemResponse, err error) {
	if request == nil {
		request = NewListInspectionItemRequest()
	}
	response = NewListInspectionItemResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCSPDownloadURLRequest() (request *CreateCSPDownloadURLRequest) {
	request = &CreateCSPDownloadURLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("inspectionplatform", APIVersion, "CreateCSPDownloadURL")
	return
}

func NewCreateCSPDownloadURLResponse() (response *CreateCSPDownloadURLResponse) {
	response = &CreateCSPDownloadURLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取CSP下载URL
func (c *Client) CreateCSPDownloadURL(request *CreateCSPDownloadURLRequest) (response *CreateCSPDownloadURLResponse, err error) {
	if request == nil {
		request = NewCreateCSPDownloadURLRequest()
	}
	response = NewCreateCSPDownloadURLResponse()
	err = c.Send(request, response)
	return
}
