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

package v20190606

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-06-06"

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

func NewAddViewRequest() (request *AddViewRequest) {
	request = &AddViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "AddView")
	return
}

func NewAddViewResponse() (response *AddViewResponse) {
	response = &AddViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus新增视图
func (c *Client) AddView(request *AddViewRequest) (response *AddViewResponse, err error) {
	if request == nil {
		request = NewAddViewRequest()
	}
	response = NewAddViewResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateViewRequest() (request *UpdateViewRequest) {
	request = &UpdateViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "UpdateView")
	return
}

func NewUpdateViewResponse() (response *UpdateViewResponse) {
	response = &UpdateViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus更新视图
func (c *Client) UpdateView(request *UpdateViewRequest) (response *UpdateViewResponse, err error) {
	if request == nil {
		request = NewUpdateViewRequest()
	}
	response = NewUpdateViewResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteFavoriteRequest() (request *DeleteFavoriteRequest) {
	request = &DeleteFavoriteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "DeleteFavorite")
	return
}

func NewDeleteFavoriteResponse() (response *DeleteFavoriteResponse) {
	response = &DeleteFavoriteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus删除收藏夹
func (c *Client) DeleteFavorite(request *DeleteFavoriteRequest) (response *DeleteFavoriteResponse, err error) {
	if request == nil {
		request = NewDeleteFavoriteRequest()
	}
	response = NewDeleteFavoriteResponse()
	err = c.Send(request, response)
	return
}

func NewGetPoliciesRequest() (request *GetPoliciesRequest) {
	request = &GetPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "GetPolicies")
	return
}

func NewGetPoliciesResponse() (response *GetPoliciesResponse) {
	response = &GetPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus查询告警策略
func (c *Client) GetPolicies(request *GetPoliciesRequest) (response *GetPoliciesResponse, err error) {
	if request == nil {
		request = NewGetPoliciesRequest()
	}
	response = NewGetPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNamespaceRequest() (request *DeleteNamespaceRequest) {
	request = &DeleteNamespaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "DeleteNamespace")
	return
}

func NewDeleteNamespaceResponse() (response *DeleteNamespaceResponse) {
	response = &DeleteNamespaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus删除namespace
func (c *Client) DeleteNamespace(request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
	if request == nil {
		request = NewDeleteNamespaceRequest()
	}
	response = NewDeleteNamespaceResponse()
	err = c.Send(request, response)
	return
}

func NewGetDataRequest() (request *GetDataRequest) {
	request = &GetDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "GetData")
	return
}

func NewGetDataResponse() (response *GetDataResponse) {
	response = &GetDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus查询数据
func (c *Client) GetData(request *GetDataRequest) (response *GetDataResponse, err error) {
	if request == nil {
		request = NewGetDataRequest()
	}
	response = NewGetDataResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteViewRequest() (request *DeleteViewRequest) {
	request = &DeleteViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "DeleteView")
	return
}

func NewDeleteViewResponse() (response *DeleteViewResponse) {
	response = &DeleteViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus删除视图
func (c *Client) DeleteView(request *DeleteViewRequest) (response *DeleteViewResponse, err error) {
	if request == nil {
		request = NewDeleteViewRequest()
	}
	response = NewDeleteViewResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateFavoriteRequest() (request *UpdateFavoriteRequest) {
	request = &UpdateFavoriteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "UpdateFavorite")
	return
}

func NewUpdateFavoriteResponse() (response *UpdateFavoriteResponse) {
	response = &UpdateFavoriteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus更新收藏夹
func (c *Client) UpdateFavorite(request *UpdateFavoriteRequest) (response *UpdateFavoriteResponse, err error) {
	if request == nil {
		request = NewUpdateFavoriteRequest()
	}
	response = NewUpdateFavoriteResponse()
	err = c.Send(request, response)
	return
}

func NewGetAlertHistoriesRequest() (request *GetAlertHistoriesRequest) {
	request = &GetAlertHistoriesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "GetAlertHistories")
	return
}

func NewGetAlertHistoriesResponse() (response *GetAlertHistoriesResponse) {
	response = &GetAlertHistoriesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus查询告警历史
func (c *Client) GetAlertHistories(request *GetAlertHistoriesRequest) (response *GetAlertHistoriesResponse, err error) {
	if request == nil {
		request = NewGetAlertHistoriesRequest()
	}
	response = NewGetAlertHistoriesResponse()
	err = c.Send(request, response)
	return
}

func NewGetViewsRequest() (request *GetViewsRequest) {
	request = &GetViewsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "GetViews")
	return
}

func NewGetViewsResponse() (response *GetViewsResponse) {
	response = &GetViewsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus查询视图
func (c *Client) GetViews(request *GetViewsRequest) (response *GetViewsResponse, err error) {
	if request == nil {
		request = NewGetViewsRequest()
	}
	response = NewGetViewsResponse()
	err = c.Send(request, response)
	return
}

func NewAddFavoriteRequest() (request *AddFavoriteRequest) {
	request = &AddFavoriteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "AddFavorite")
	return
}

func NewAddFavoriteResponse() (response *AddFavoriteResponse) {
	response = &AddFavoriteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus新增收藏夹
func (c *Client) AddFavorite(request *AddFavoriteRequest) (response *AddFavoriteResponse, err error) {
	if request == nil {
		request = NewAddFavoriteRequest()
	}
	response = NewAddFavoriteResponse()
	err = c.Send(request, response)
	return
}

func NewGetDimensionsRequest() (request *GetDimensionsRequest) {
	request = &GetDimensionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "GetDimensions")
	return
}

func NewGetDimensionsResponse() (response *GetDimensionsResponse) {
	response = &GetDimensionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus获取维度下拉框数据
func (c *Client) GetDimensions(request *GetDimensionsRequest) (response *GetDimensionsResponse, err error) {
	if request == nil {
		request = NewGetDimensionsRequest()
	}
	response = NewGetDimensionsResponse()
	err = c.Send(request, response)
	return
}

func NewArgusReportRequest() (request *ArgusReportRequest) {
	request = &ArgusReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "ArgusReport")
	return
}

func NewArgusReportResponse() (response *ArgusReportResponse) {
	response = &ArgusReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// （1）提供Line Protocol和Json两种数据格式上报
// （2）单用户上报请求数限制为100次/s，单个请求最大包体为256K，超过大小返回错误，最多上报150条，超过限制的数据自动丢弃。
// （3）支持单个请求同时上报多个命名空间，上报的错误命名空间数据会进行丢弃，返回结果中显示上报成功数和失败数。
// （4）对上报的不存在的维度和指标进行忽略，未上报的维度以空字符串进行存储
// （5）提供多种语言的SDK上报，同时API 3.0 Explorer提供自动生成代码和测试上报功能
func (c *Client) ArgusReport(request *ArgusReportRequest) (response *ArgusReportResponse, err error) {
	if request == nil {
		request = NewArgusReportRequest()
	}
	response = NewArgusReportResponse()
	err = c.Send(request, response)
	return
}

func NewGetNamespacesRequest() (request *GetNamespacesRequest) {
	request = &GetNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "GetNamespaces")
	return
}

func NewGetNamespacesResponse() (response *GetNamespacesResponse) {
	response = &GetNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus查询namespace
func (c *Client) GetNamespaces(request *GetNamespacesRequest) (response *GetNamespacesResponse, err error) {
	if request == nil {
		request = NewGetNamespacesRequest()
	}
	response = NewGetNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewDescPolicyRequest() (request *DescPolicyRequest) {
	request = &DescPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "DescPolicy")
	return
}

func NewDescPolicyResponse() (response *DescPolicyResponse) {
	response = &DescPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus获取告警策略详情
func (c *Client) DescPolicy(request *DescPolicyRequest) (response *DescPolicyResponse, err error) {
	if request == nil {
		request = NewDescPolicyRequest()
	}
	response = NewDescPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescViewRequest() (request *DescViewRequest) {
	request = &DescViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "DescView")
	return
}

func NewDescViewResponse() (response *DescViewResponse) {
	response = &DescViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus获取视图详情
func (c *Client) DescView(request *DescViewRequest) (response *DescViewResponse, err error) {
	if request == nil {
		request = NewDescViewRequest()
	}
	response = NewDescViewResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateNamespaceRequest() (request *UpdateNamespaceRequest) {
	request = &UpdateNamespaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "UpdateNamespace")
	return
}

func NewUpdateNamespaceResponse() (response *UpdateNamespaceResponse) {
	response = &UpdateNamespaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus更新namespace
func (c *Client) UpdateNamespace(request *UpdateNamespaceRequest) (response *UpdateNamespaceResponse, err error) {
	if request == nil {
		request = NewUpdateNamespaceRequest()
	}
	response = NewUpdateNamespaceResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePolicyRequest() (request *DeletePolicyRequest) {
	request = &DeletePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "DeletePolicy")
	return
}

func NewDeletePolicyResponse() (response *DeletePolicyResponse) {
	response = &DeletePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus删除告警策略
func (c *Client) DeletePolicy(request *DeletePolicyRequest) (response *DeletePolicyResponse, err error) {
	if request == nil {
		request = NewDeletePolicyRequest()
	}
	response = NewDeletePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewGetFavoritesRequest() (request *GetFavoritesRequest) {
	request = &GetFavoritesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "GetFavorites")
	return
}

func NewGetFavoritesResponse() (response *GetFavoritesResponse) {
	response = &GetFavoritesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus查询收藏夹
func (c *Client) GetFavorites(request *GetFavoritesRequest) (response *GetFavoritesResponse, err error) {
	if request == nil {
		request = NewGetFavoritesRequest()
	}
	response = NewGetFavoritesResponse()
	err = c.Send(request, response)
	return
}

func NewDescFavoriteRequest() (request *DescFavoriteRequest) {
	request = &DescFavoriteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "DescFavorite")
	return
}

func NewDescFavoriteResponse() (response *DescFavoriteResponse) {
	response = &DescFavoriteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus获取收藏夹详情
func (c *Client) DescFavorite(request *DescFavoriteRequest) (response *DescFavoriteResponse, err error) {
	if request == nil {
		request = NewDescFavoriteRequest()
	}
	response = NewDescFavoriteResponse()
	err = c.Send(request, response)
	return
}

func NewGetParamsRequest() (request *GetParamsRequest) {
	request = &GetParamsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "GetParams")
	return
}

func NewGetParamsResponse() (response *GetParamsResponse) {
	response = &GetParamsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus获取下拉框参数
func (c *Client) GetParams(request *GetParamsRequest) (response *GetParamsResponse, err error) {
	if request == nil {
		request = NewGetParamsRequest()
	}
	response = NewGetParamsResponse()
	err = c.Send(request, response)
	return
}

func NewAddNamespaceRequest() (request *AddNamespaceRequest) {
	request = &AddNamespaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "AddNamespace")
	return
}

func NewAddNamespaceResponse() (response *AddNamespaceResponse) {
	response = &AddNamespaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus新增namespace
func (c *Client) AddNamespace(request *AddNamespaceRequest) (response *AddNamespaceResponse, err error) {
	if request == nil {
		request = NewAddNamespaceRequest()
	}
	response = NewAddNamespaceResponse()
	err = c.Send(request, response)
	return
}

func NewDescNamespaceRequest() (request *DescNamespaceRequest) {
	request = &DescNamespaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "DescNamespace")
	return
}

func NewDescNamespaceResponse() (response *DescNamespaceResponse) {
	response = &DescNamespaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus获取namespace详情
func (c *Client) DescNamespace(request *DescNamespaceRequest) (response *DescNamespaceResponse, err error) {
	if request == nil {
		request = NewDescNamespaceRequest()
	}
	response = NewDescNamespaceResponse()
	err = c.Send(request, response)
	return
}

func NewAddPolicyRequest() (request *AddPolicyRequest) {
	request = &AddPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "AddPolicy")
	return
}

func NewAddPolicyResponse() (response *AddPolicyResponse) {
	response = &AddPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus新增告警策略
func (c *Client) AddPolicy(request *AddPolicyRequest) (response *AddPolicyResponse, err error) {
	if request == nil {
		request = NewAddPolicyRequest()
	}
	response = NewAddPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewGetOverviewRequest() (request *GetOverviewRequest) {
	request = &GetOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "GetOverview")
	return
}

func NewGetOverviewResponse() (response *GetOverviewResponse) {
	response = &GetOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus查询概览
func (c *Client) GetOverview(request *GetOverviewRequest) (response *GetOverviewResponse, err error) {
	if request == nil {
		request = NewGetOverviewRequest()
	}
	response = NewGetOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicyRequest() (request *UpdatePolicyRequest) {
	request = &UpdatePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opargus", APIVersion, "UpdatePolicy")
	return
}

func NewUpdatePolicyResponse() (response *UpdatePolicyResponse) {
	response = &UpdatePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Argus更新告警策略
func (c *Client) UpdatePolicy(request *UpdatePolicyRequest) (response *UpdatePolicyResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyRequest()
	}
	response = NewUpdatePolicyResponse()
	err = c.Send(request, response)
	return
}
