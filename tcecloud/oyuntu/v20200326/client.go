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

package v20200326

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-03-26"

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

func NewGetModuleUrlRelateApiRequest() (request *GetModuleUrlRelateApiRequest) {
	request = &GetModuleUrlRelateApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetModuleUrlRelateApi")
	return
}

func NewGetModuleUrlRelateApiResponse() (response *GetModuleUrlRelateApiResponse) {
	response = &GetModuleUrlRelateApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取后端地址关联的API
func (c *Client) GetModuleUrlRelateApi(request *GetModuleUrlRelateApiRequest) (response *GetModuleUrlRelateApiResponse, err error) {
	if request == nil {
		request = NewGetModuleUrlRelateApiRequest()
	}
	response = NewGetModuleUrlRelateApiResponse()
	err = c.Send(request, response)
	return
}

func NewGetCategoryByIDRequest() (request *GetCategoryByIDRequest) {
	request = &GetCategoryByIDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetCategoryByID")
	return
}

func NewGetCategoryByIDResponse() (response *GetCategoryByIDResponse) {
	response = &GetCategoryByIDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按id查询官方文档
func (c *Client) GetCategoryByID(request *GetCategoryByIDRequest) (response *GetCategoryByIDResponse, err error) {
	if request == nil {
		request = NewGetCategoryByIDRequest()
	}
	response = NewGetCategoryByIDResponse()
	err = c.Send(request, response)
	return
}

func NewQueryApiInfoFilterRequest() (request *QueryApiInfoFilterRequest) {
	request = &QueryApiInfoFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "QueryApiInfoFilter")
	return
}

func NewQueryApiInfoFilterResponse() (response *QueryApiInfoFilterResponse) {
	response = &QueryApiInfoFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页、搜索列表API接口
func (c *Client) QueryApiInfoFilter(request *QueryApiInfoFilterRequest) (response *QueryApiInfoFilterResponse, err error) {
	if request == nil {
		request = NewQueryApiInfoFilterRequest()
	}
	response = NewQueryApiInfoFilterResponse()
	err = c.Send(request, response)
	return
}

func NewGetApiInfoRequest() (request *GetApiInfoRequest) {
	request = &GetApiInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetApiInfo")
	return
}

func NewGetApiInfoResponse() (response *GetApiInfoResponse) {
	response = &GetApiInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取API详情
func (c *Client) GetApiInfo(request *GetApiInfoRequest) (response *GetApiInfoResponse, err error) {
	if request == nil {
		request = NewGetApiInfoRequest()
	}
	response = NewGetApiInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGetSystemCommonParamsRequest() (request *GetSystemCommonParamsRequest) {
	request = &GetSystemCommonParamsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetSystemCommonParams")
	return
}

func NewGetSystemCommonParamsResponse() (response *GetSystemCommonParamsResponse) {
	response = &GetSystemCommonParamsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取系统公共参数
func (c *Client) GetSystemCommonParams(request *GetSystemCommonParamsRequest) (response *GetSystemCommonParamsResponse, err error) {
	if request == nil {
		request = NewGetSystemCommonParamsRequest()
	}
	response = NewGetSystemCommonParamsResponse()
	err = c.Send(request, response)
	return
}

func NewAddApiInfoRequest() (request *AddApiInfoRequest) {
	request = &AddApiInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "AddApiInfo")
	return
}

func NewAddApiInfoResponse() (response *AddApiInfoResponse) {
	response = &AddApiInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增接口
func (c *Client) AddApiInfo(request *AddApiInfoRequest) (response *AddApiInfoResponse, err error) {
	if request == nil {
		request = NewAddApiInfoRequest()
	}
	response = NewAddApiInfoResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateModuleRequest() (request *UpdateModuleRequest) {
	request = &UpdateModuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "UpdateModule")
	return
}

func NewUpdateModuleResponse() (response *UpdateModuleResponse) {
	response = &UpdateModuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新产品模块描述或产品负责人
func (c *Client) UpdateModule(request *UpdateModuleRequest) (response *UpdateModuleResponse, err error) {
	if request == nil {
		request = NewUpdateModuleRequest()
	}
	response = NewUpdateModuleResponse()
	err = c.Send(request, response)
	return
}

func NewGetModuleErrorCodeRelateApiRequest() (request *GetModuleErrorCodeRelateApiRequest) {
	request = &GetModuleErrorCodeRelateApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetModuleErrorCodeRelateApi")
	return
}

func NewGetModuleErrorCodeRelateApiResponse() (response *GetModuleErrorCodeRelateApiResponse) {
	response = &GetModuleErrorCodeRelateApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询错误码关联的接口
func (c *Client) GetModuleErrorCodeRelateApi(request *GetModuleErrorCodeRelateApiRequest) (response *GetModuleErrorCodeRelateApiResponse, err error) {
	if request == nil {
		request = NewGetModuleErrorCodeRelateApiRequest()
	}
	response = NewGetModuleErrorCodeRelateApiResponse()
	err = c.Send(request, response)
	return
}

func NewGetModuleUrlRequest() (request *GetModuleUrlRequest) {
	request = &GetModuleUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetModuleUrl")
	return
}

func NewGetModuleUrlResponse() (response *GetModuleUrlResponse) {
	response = &GetModuleUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取后端地址
func (c *Client) GetModuleUrl(request *GetModuleUrlRequest) (response *GetModuleUrlResponse, err error) {
	if request == nil {
		request = NewGetModuleUrlRequest()
	}
	response = NewGetModuleUrlResponse()
	err = c.Send(request, response)
	return
}

func NewQueryModuleUrlFilterRequest() (request *QueryModuleUrlFilterRequest) {
	request = &QueryModuleUrlFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "QueryModuleUrlFilter")
	return
}

func NewQueryModuleUrlFilterResponse() (response *QueryModuleUrlFilterResponse) {
	response = &QueryModuleUrlFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页、搜索、列表后端地址
func (c *Client) QueryModuleUrlFilter(request *QueryModuleUrlFilterRequest) (response *QueryModuleUrlFilterResponse, err error) {
	if request == nil {
		request = NewQueryModuleUrlFilterRequest()
	}
	response = NewQueryModuleUrlFilterResponse()
	err = c.Send(request, response)
	return
}

func NewRunActionRequest() (request *RunActionRequest) {
	request = &RunActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "RunAction")
	return
}

func NewRunActionResponse() (response *RunActionResponse) {
	response = &RunActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// api explorer 入口
func (c *Client) RunAction(request *RunActionRequest) (response *RunActionResponse, err error) {
	if request == nil {
		request = NewRunActionRequest()
	}
	response = NewRunActionResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateComplexObjectRequest() (request *UpdateComplexObjectRequest) {
	request = &UpdateComplexObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "UpdateComplexObject")
	return
}

func NewUpdateComplexObjectResponse() (response *UpdateComplexObjectResponse) {
	response = &UpdateComplexObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新复杂对象
func (c *Client) UpdateComplexObject(request *UpdateComplexObjectRequest) (response *UpdateComplexObjectResponse, err error) {
	if request == nil {
		request = NewUpdateComplexObjectRequest()
	}
	response = NewUpdateComplexObjectResponse()
	err = c.Send(request, response)
	return
}

func NewQueryVersionFilterRequest() (request *QueryVersionFilterRequest) {
	request = &QueryVersionFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "QueryVersionFilter")
	return
}

func NewQueryVersionFilterResponse() (response *QueryVersionFilterResponse) {
	response = &QueryVersionFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 过滤查询模块下的接口版本
func (c *Client) QueryVersionFilter(request *QueryVersionFilterRequest) (response *QueryVersionFilterResponse, err error) {
	if request == nil {
		request = NewQueryVersionFilterRequest()
	}
	response = NewQueryVersionFilterResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCategoryRequest() (request *DeleteCategoryRequest) {
	request = &DeleteCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "DeleteCategory")
	return
}

func NewDeleteCategoryResponse() (response *DeleteCategoryResponse) {
	response = &DeleteCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除官方文档分类
func (c *Client) DeleteCategory(request *DeleteCategoryRequest) (response *DeleteCategoryResponse, err error) {
	if request == nil {
		request = NewDeleteCategoryRequest()
	}
	response = NewDeleteCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewGetCategorysRequest() (request *GetCategorysRequest) {
	request = &GetCategorysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetCategorys")
	return
}

func NewGetCategorysResponse() (response *GetCategorysResponse) {
	response = &GetCategorysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取接口文档分类
func (c *Client) GetCategorys(request *GetCategorysRequest) (response *GetCategorysResponse, err error) {
	if request == nil {
		request = NewGetCategorysRequest()
	}
	response = NewGetCategorysResponse()
	err = c.Send(request, response)
	return
}

func NewQueryModuleErrorCodeFilterRequest() (request *QueryModuleErrorCodeFilterRequest) {
	request = &QueryModuleErrorCodeFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "QueryModuleErrorCodeFilter")
	return
}

func NewQueryModuleErrorCodeFilterResponse() (response *QueryModuleErrorCodeFilterResponse) {
	response = &QueryModuleErrorCodeFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页、搜索、列表错误码
func (c *Client) QueryModuleErrorCodeFilter(request *QueryModuleErrorCodeFilterRequest) (response *QueryModuleErrorCodeFilterResponse, err error) {
	if request == nil {
		request = NewQueryModuleErrorCodeFilterRequest()
	}
	response = NewQueryModuleErrorCodeFilterResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCategoryFilterRequest() (request *QueryCategoryFilterRequest) {
	request = &QueryCategoryFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "QueryCategoryFilter")
	return
}

func NewQueryCategoryFilterResponse() (response *QueryCategoryFilterResponse) {
	response = &QueryCategoryFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页搜索列表官文文档分类
func (c *Client) QueryCategoryFilter(request *QueryCategoryFilterRequest) (response *QueryCategoryFilterResponse, err error) {
	if request == nil {
		request = NewQueryCategoryFilterRequest()
	}
	response = NewQueryCategoryFilterResponse()
	err = c.Send(request, response)
	return
}

func NewQueryObjectHistoryFilterRequest() (request *QueryObjectHistoryFilterRequest) {
	request = &QueryObjectHistoryFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "QueryObjectHistoryFilter")
	return
}

func NewQueryObjectHistoryFilterResponse() (response *QueryObjectHistoryFilterResponse) {
	response = &QueryObjectHistoryFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询变更历史记录
func (c *Client) QueryObjectHistoryFilter(request *QueryObjectHistoryFilterRequest) (response *QueryObjectHistoryFilterResponse, err error) {
	if request == nil {
		request = NewQueryObjectHistoryFilterRequest()
	}
	response = NewQueryObjectHistoryFilterResponse()
	err = c.Send(request, response)
	return
}

func NewGetComplexObjectRequest() (request *GetComplexObjectRequest) {
	request = &GetComplexObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetComplexObject")
	return
}

func NewGetComplexObjectResponse() (response *GetComplexObjectResponse) {
	response = &GetComplexObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取单个复杂对象的详情
func (c *Client) GetComplexObject(request *GetComplexObjectRequest) (response *GetComplexObjectResponse, err error) {
	if request == nil {
		request = NewGetComplexObjectRequest()
	}
	response = NewGetComplexObjectResponse()
	err = c.Send(request, response)
	return
}

func NewGetModuleUrlsRequest() (request *GetModuleUrlsRequest) {
	request = &GetModuleUrlsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetModuleUrls")
	return
}

func NewGetModuleUrlsResponse() (response *GetModuleUrlsResponse) {
	response = &GetModuleUrlsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取模块下所有的后端地址
func (c *Client) GetModuleUrls(request *GetModuleUrlsRequest) (response *GetModuleUrlsResponse, err error) {
	if request == nil {
		request = NewGetModuleUrlsRequest()
	}
	response = NewGetModuleUrlsResponse()
	err = c.Send(request, response)
	return
}

func NewGetFirstLevelErrorCodeRequest() (request *GetFirstLevelErrorCodeRequest) {
	request = &GetFirstLevelErrorCodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetFirstLevelErrorCode")
	return
}

func NewGetFirstLevelErrorCodeResponse() (response *GetFirstLevelErrorCodeResponse) {
	response = &GetFirstLevelErrorCodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有的一级错误码
func (c *Client) GetFirstLevelErrorCode(request *GetFirstLevelErrorCodeRequest) (response *GetFirstLevelErrorCodeResponse, err error) {
	if request == nil {
		request = NewGetFirstLevelErrorCodeRequest()
	}
	response = NewGetFirstLevelErrorCodeResponse()
	err = c.Send(request, response)
	return
}

func NewQueryModuleFilterRequest() (request *QueryModuleFilterRequest) {
	request = &QueryModuleFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "QueryModuleFilter")
	return
}

func NewQueryModuleFilterResponse() (response *QueryModuleFilterResponse) {
	response = &QueryModuleFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页、搜索、列表模块
func (c *Client) QueryModuleFilter(request *QueryModuleFilterRequest) (response *QueryModuleFilterResponse, err error) {
	if request == nil {
		request = NewQueryModuleFilterRequest()
	}
	response = NewQueryModuleFilterResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVersionRequest() (request *DeleteVersionRequest) {
	request = &DeleteVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "DeleteVersion")
	return
}

func NewDeleteVersionResponse() (response *DeleteVersionResponse) {
	response = &DeleteVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除模块下某一接口版本
func (c *Client) DeleteVersion(request *DeleteVersionRequest) (response *DeleteVersionResponse, err error) {
	if request == nil {
		request = NewDeleteVersionRequest()
	}
	response = NewDeleteVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApiInfoRequest() (request *DeleteApiInfoRequest) {
	request = &DeleteApiInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "DeleteApiInfo")
	return
}

func NewDeleteApiInfoResponse() (response *DeleteApiInfoResponse) {
	response = &DeleteApiInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除接口信息
func (c *Client) DeleteApiInfo(request *DeleteApiInfoRequest) (response *DeleteApiInfoResponse, err error) {
	if request == nil {
		request = NewDeleteApiInfoRequest()
	}
	response = NewDeleteApiInfoResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateModuleUrlRequest() (request *UpdateModuleUrlRequest) {
	request = &UpdateModuleUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "UpdateModuleUrl")
	return
}

func NewUpdateModuleUrlResponse() (response *UpdateModuleUrlResponse) {
	response = &UpdateModuleUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新后端地址
func (c *Client) UpdateModuleUrl(request *UpdateModuleUrlRequest) (response *UpdateModuleUrlResponse, err error) {
	if request == nil {
		request = NewUpdateModuleUrlRequest()
	}
	response = NewUpdateModuleUrlResponse()
	err = c.Send(request, response)
	return
}

func NewQueryComplexObjectFilterRequest() (request *QueryComplexObjectFilterRequest) {
	request = &QueryComplexObjectFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "QueryComplexObjectFilter")
	return
}

func NewQueryComplexObjectFilterResponse() (response *QueryComplexObjectFilterResponse) {
	response = &QueryComplexObjectFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页、搜索、列表复杂对象信息
func (c *Client) QueryComplexObjectFilter(request *QueryComplexObjectFilterRequest) (response *QueryComplexObjectFilterResponse, err error) {
	if request == nil {
		request = NewQueryComplexObjectFilterRequest()
	}
	response = NewQueryComplexObjectFilterResponse()
	err = c.Send(request, response)
	return
}

func NewAddModuleRequest() (request *AddModuleRequest) {
	request = &AddModuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "AddModule")
	return
}

func NewAddModuleResponse() (response *AddModuleResponse) {
	response = &AddModuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增产品模块名
func (c *Client) AddModule(request *AddModuleRequest) (response *AddModuleResponse, err error) {
	if request == nil {
		request = NewAddModuleRequest()
	}
	response = NewAddModuleResponse()
	err = c.Send(request, response)
	return
}

func NewGetVersionsRequest() (request *GetVersionsRequest) {
	request = &GetVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetVersions")
	return
}

func NewGetVersionsResponse() (response *GetVersionsResponse) {
	response = &GetVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询模块下的接口版本
func (c *Client) GetVersions(request *GetVersionsRequest) (response *GetVersionsResponse, err error) {
	if request == nil {
		request = NewGetVersionsRequest()
	}
	response = NewGetVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewSyncDocumentRequest() (request *SyncDocumentRequest) {
	request = &SyncDocumentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "SyncDocument")
	return
}

func NewSyncDocumentResponse() (response *SyncDocumentResponse) {
	response = &SyncDocumentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 生成API文档并同步到文档中心
func (c *Client) SyncDocument(request *SyncDocumentRequest) (response *SyncDocumentResponse, err error) {
	if request == nil {
		request = NewSyncDocumentRequest()
	}
	response = NewSyncDocumentResponse()
	err = c.Send(request, response)
	return
}

func NewAddCategoryRequest() (request *AddCategoryRequest) {
	request = &AddCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "AddCategory")
	return
}

func NewAddCategoryResponse() (response *AddCategoryResponse) {
	response = &AddCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增官方文档分类
func (c *Client) AddCategory(request *AddCategoryRequest) (response *AddCategoryResponse, err error) {
	if request == nil {
		request = NewAddCategoryRequest()
	}
	response = NewAddCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewAddModuleUrlRequest() (request *AddModuleUrlRequest) {
	request = &AddModuleUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "AddModuleUrl")
	return
}

func NewAddModuleUrlResponse() (response *AddModuleUrlResponse) {
	response = &AddModuleUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增后端地址
func (c *Client) AddModuleUrl(request *AddModuleUrlRequest) (response *AddModuleUrlResponse, err error) {
	if request == nil {
		request = NewAddModuleUrlRequest()
	}
	response = NewAddModuleUrlResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteModuleErrorCodeRequest() (request *DeleteModuleErrorCodeRequest) {
	request = &DeleteModuleErrorCodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "DeleteModuleErrorCode")
	return
}

func NewDeleteModuleErrorCodeResponse() (response *DeleteModuleErrorCodeResponse) {
	response = &DeleteModuleErrorCodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除错误码
func (c *Client) DeleteModuleErrorCode(request *DeleteModuleErrorCodeRequest) (response *DeleteModuleErrorCodeResponse, err error) {
	if request == nil {
		request = NewDeleteModuleErrorCodeRequest()
	}
	response = NewDeleteModuleErrorCodeResponse()
	err = c.Send(request, response)
	return
}

func NewGetModuleErrorCodesRequest() (request *GetModuleErrorCodesRequest) {
	request = &GetModuleErrorCodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetModuleErrorCodes")
	return
}

func NewGetModuleErrorCodesResponse() (response *GetModuleErrorCodesResponse) {
	response = &GetModuleErrorCodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取模块下的错误码code和信息
func (c *Client) GetModuleErrorCodes(request *GetModuleErrorCodesRequest) (response *GetModuleErrorCodesResponse, err error) {
	if request == nil {
		request = NewGetModuleErrorCodesRequest()
	}
	response = NewGetModuleErrorCodesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateModuleUrlToActionRequest() (request *UpdateModuleUrlToActionRequest) {
	request = &UpdateModuleUrlToActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "UpdateModuleUrlToAction")
	return
}

func NewUpdateModuleUrlToActionResponse() (response *UpdateModuleUrlToActionResponse) {
	response = &UpdateModuleUrlToActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定后端地址到API接口
func (c *Client) UpdateModuleUrlToAction(request *UpdateModuleUrlToActionRequest) (response *UpdateModuleUrlToActionResponse, err error) {
	if request == nil {
		request = NewUpdateModuleUrlToActionRequest()
	}
	response = NewUpdateModuleUrlToActionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteModuleRequest() (request *DeleteModuleRequest) {
	request = &DeleteModuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "DeleteModule")
	return
}

func NewDeleteModuleResponse() (response *DeleteModuleResponse) {
	response = &DeleteModuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除产品模块
func (c *Client) DeleteModule(request *DeleteModuleRequest) (response *DeleteModuleResponse, err error) {
	if request == nil {
		request = NewDeleteModuleRequest()
	}
	response = NewDeleteModuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteModuleUrlRequest() (request *DeleteModuleUrlRequest) {
	request = &DeleteModuleUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "DeleteModuleUrl")
	return
}

func NewDeleteModuleUrlResponse() (response *DeleteModuleUrlResponse) {
	response = &DeleteModuleUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除后端地址
func (c *Client) DeleteModuleUrl(request *DeleteModuleUrlRequest) (response *DeleteModuleUrlResponse, err error) {
	if request == nil {
		request = NewDeleteModuleUrlRequest()
	}
	response = NewDeleteModuleUrlResponse()
	err = c.Send(request, response)
	return
}

func NewGetApiInfoHistoryRequest() (request *GetApiInfoHistoryRequest) {
	request = &GetApiInfoHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetApiInfoHistory")
	return
}

func NewGetApiInfoHistoryResponse() (response *GetApiInfoHistoryResponse) {
	response = &GetApiInfoHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取接口的变更记录
func (c *Client) GetApiInfoHistory(request *GetApiInfoHistoryRequest) (response *GetApiInfoHistoryResponse, err error) {
	if request == nil {
		request = NewGetApiInfoHistoryRequest()
	}
	response = NewGetApiInfoHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewAddComplexObjectRequest() (request *AddComplexObjectRequest) {
	request = &AddComplexObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "AddComplexObject")
	return
}

func NewAddComplexObjectResponse() (response *AddComplexObjectResponse) {
	response = &AddComplexObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增复杂对象
func (c *Client) AddComplexObject(request *AddComplexObjectRequest) (response *AddComplexObjectResponse, err error) {
	if request == nil {
		request = NewAddComplexObjectRequest()
	}
	response = NewAddComplexObjectResponse()
	err = c.Send(request, response)
	return
}

func NewGetEnvironmentRequest() (request *GetEnvironmentRequest) {
	request = &GetEnvironmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetEnvironment")
	return
}

func NewGetEnvironmentResponse() (response *GetEnvironmentResponse) {
	response = &GetEnvironmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当前是客户环境还是内部环境
func (c *Client) GetEnvironment(request *GetEnvironmentRequest) (response *GetEnvironmentResponse, err error) {
	if request == nil {
		request = NewGetEnvironmentRequest()
	}
	response = NewGetEnvironmentResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteComplexObjectRequest() (request *DeleteComplexObjectRequest) {
	request = &DeleteComplexObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "DeleteComplexObject")
	return
}

func NewDeleteComplexObjectResponse() (response *DeleteComplexObjectResponse) {
	response = &DeleteComplexObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除复杂对象
func (c *Client) DeleteComplexObject(request *DeleteComplexObjectRequest) (response *DeleteComplexObjectResponse, err error) {
	if request == nil {
		request = NewDeleteComplexObjectRequest()
	}
	response = NewDeleteComplexObjectResponse()
	err = c.Send(request, response)
	return
}

func NewGetDataTypesRequest() (request *GetDataTypesRequest) {
	request = &GetDataTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetDataTypes")
	return
}

func NewGetDataTypesResponse() (response *GetDataTypesResponse) {
	response = &GetDataTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询数据类型
func (c *Client) GetDataTypes(request *GetDataTypesRequest) (response *GetDataTypesResponse, err error) {
	if request == nil {
		request = NewGetDataTypesRequest()
	}
	response = NewGetDataTypesResponse()
	err = c.Send(request, response)
	return
}

func NewGetCategoryRequest() (request *GetCategoryRequest) {
	request = &GetCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetCategory")
	return
}

func NewGetCategoryResponse() (response *GetCategoryResponse) {
	response = &GetCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取模块和版本下的文档分类
func (c *Client) GetCategory(request *GetCategoryRequest) (response *GetCategoryResponse, err error) {
	if request == nil {
		request = NewGetCategoryRequest()
	}
	response = NewGetCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewGetComplexObjectRelateApiRequest() (request *GetComplexObjectRelateApiRequest) {
	request = &GetComplexObjectRelateApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetComplexObjectRelateApi")
	return
}

func NewGetComplexObjectRelateApiResponse() (response *GetComplexObjectRelateApiResponse) {
	response = &GetComplexObjectRelateApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取复杂对象分别作为入参、出参时关联的接口
func (c *Client) GetComplexObjectRelateApi(request *GetComplexObjectRelateApiRequest) (response *GetComplexObjectRelateApiResponse, err error) {
	if request == nil {
		request = NewGetComplexObjectRelateApiRequest()
	}
	response = NewGetComplexObjectRelateApiResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateModuleErrorCodeRequest() (request *UpdateModuleErrorCodeRequest) {
	request = &UpdateModuleErrorCodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "UpdateModuleErrorCode")
	return
}

func NewUpdateModuleErrorCodeResponse() (response *UpdateModuleErrorCodeResponse) {
	response = &UpdateModuleErrorCodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新错误码
func (c *Client) UpdateModuleErrorCode(request *UpdateModuleErrorCodeRequest) (response *UpdateModuleErrorCodeResponse, err error) {
	if request == nil {
		request = NewUpdateModuleErrorCodeRequest()
	}
	response = NewUpdateModuleErrorCodeResponse()
	err = c.Send(request, response)
	return
}

func NewGetObjectHistoryIdsRequest() (request *GetObjectHistoryIdsRequest) {
	request = &GetObjectHistoryIdsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetObjectHistoryIds")
	return
}

func NewGetObjectHistoryIdsResponse() (response *GetObjectHistoryIdsResponse) {
	response = &GetObjectHistoryIdsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取各类型变更记录的ID
func (c *Client) GetObjectHistoryIds(request *GetObjectHistoryIdsRequest) (response *GetObjectHistoryIdsResponse, err error) {
	if request == nil {
		request = NewGetObjectHistoryIdsRequest()
	}
	response = NewGetObjectHistoryIdsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSdkDemosRequest() (request *DescribeSdkDemosRequest) {
	request = &DescribeSdkDemosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "DescribeSdkDemos")
	return
}

func NewDescribeSdkDemosResponse() (response *DescribeSdkDemosResponse) {
	response = &DescribeSdkDemosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Action的代码模板
func (c *Client) DescribeSdkDemos(request *DescribeSdkDemosRequest) (response *DescribeSdkDemosResponse, err error) {
	if request == nil {
		request = NewDescribeSdkDemosRequest()
	}
	response = NewDescribeSdkDemosResponse()
	err = c.Send(request, response)
	return
}

func NewGetModuleNamesRequest() (request *GetModuleNamesRequest) {
	request = &GetModuleNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetModuleNames")
	return
}

func NewGetModuleNamesResponse() (response *GetModuleNamesResponse) {
	response = &GetModuleNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回所有的模块名
func (c *Client) GetModuleNames(request *GetModuleNamesRequest) (response *GetModuleNamesResponse, err error) {
	if request == nil {
		request = NewGetModuleNamesRequest()
	}
	response = NewGetModuleNamesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApiInfoRequest() (request *UpdateApiInfoRequest) {
	request = &UpdateApiInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "UpdateApiInfo")
	return
}

func NewUpdateApiInfoResponse() (response *UpdateApiInfoResponse) {
	response = &UpdateApiInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新接口信息
func (c *Client) UpdateApiInfo(request *UpdateApiInfoRequest) (response *UpdateApiInfoResponse, err error) {
	if request == nil {
		request = NewUpdateApiInfoRequest()
	}
	response = NewUpdateApiInfoResponse()
	err = c.Send(request, response)
	return
}

func NewAddModuleErrorCodeRequest() (request *AddModuleErrorCodeRequest) {
	request = &AddModuleErrorCodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "AddModuleErrorCode")
	return
}

func NewAddModuleErrorCodeResponse() (response *AddModuleErrorCodeResponse) {
	response = &AddModuleErrorCodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增错误码
func (c *Client) AddModuleErrorCode(request *AddModuleErrorCodeRequest) (response *AddModuleErrorCodeResponse, err error) {
	if request == nil {
		request = NewAddModuleErrorCodeRequest()
	}
	response = NewAddModuleErrorCodeResponse()
	err = c.Send(request, response)
	return
}

func NewGetComplexObjectByObjIDRequest() (request *GetComplexObjectByObjIDRequest) {
	request = &GetComplexObjectByObjIDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetComplexObjectByObjID")
	return
}

func NewGetComplexObjectByObjIDResponse() (response *GetComplexObjectByObjIDResponse) {
	response = &GetComplexObjectByObjIDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过复杂对象ID获取详情
func (c *Client) GetComplexObjectByObjID(request *GetComplexObjectByObjIDRequest) (response *GetComplexObjectByObjIDResponse, err error) {
	if request == nil {
		request = NewGetComplexObjectByObjIDRequest()
	}
	response = NewGetComplexObjectByObjIDResponse()
	err = c.Send(request, response)
	return
}

func NewGetRegionsRequest() (request *GetRegionsRequest) {
	request = &GetRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "GetRegions")
	return
}

func NewGetRegionsResponse() (response *GetRegionsResponse) {
	response = &GetRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetRegions
func (c *Client) GetRegions(request *GetRegionsRequest) (response *GetRegionsResponse, err error) {
	if request == nil {
		request = NewGetRegionsRequest()
	}
	response = NewGetRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewAddVersionRequest() (request *AddVersionRequest) {
	request = &AddVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "AddVersion")
	return
}

func NewAddVersionResponse() (response *AddVersionResponse) {
	response = &AddVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增接口版本
func (c *Client) AddVersion(request *AddVersionRequest) (response *AddVersionResponse, err error) {
	if request == nil {
		request = NewAddVersionRequest()
	}
	response = NewAddVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeActionRequest() (request *DescribeActionRequest) {
	request = &DescribeActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "DescribeAction")
	return
}

func NewDescribeActionResponse() (response *DescribeActionResponse) {
	response = &DescribeActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询接口信息
func (c *Client) DescribeAction(request *DescribeActionRequest) (response *DescribeActionResponse, err error) {
	if request == nil {
		request = NewDescribeActionRequest()
	}
	response = NewDescribeActionResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCategoryRequest() (request *UpdateCategoryRequest) {
	request = &UpdateCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "UpdateCategory")
	return
}

func NewUpdateCategoryResponse() (response *UpdateCategoryResponse) {
	response = &UpdateCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新官方文档分类名
func (c *Client) UpdateCategory(request *UpdateCategoryRequest) (response *UpdateCategoryResponse, err error) {
	if request == nil {
		request = NewUpdateCategoryRequest()
	}
	response = NewUpdateCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadSdkRequest() (request *DownloadSdkRequest) {
	request = &DownloadSdkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("oyuntu", APIVersion, "DownloadSdk")
	return
}

func NewDownloadSdkResponse() (response *DownloadSdkResponse) {
	response = &DownloadSdkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载API SDK
func (c *Client) DownloadSdk(request *DownloadSdkRequest) (response *DownloadSdkResponse, err error) {
	if request == nil {
		request = NewDownloadSdkRequest()
	}
	response = NewDownloadSdkResponse()
	err = c.Send(request, response)
	return
}
