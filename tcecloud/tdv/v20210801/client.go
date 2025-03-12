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

package v20210801

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2021-08-01"

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

func NewPublishPageRequest() (request *PublishPageRequest) {
	request = &PublishPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "PublishPage")
	return
}

func NewPublishPageResponse() (response *PublishPageResponse) {
	response = &PublishPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发布页面
func (c *Client) PublishPage(request *PublishPageRequest) (response *PublishPageResponse, err error) {
	if request == nil {
		request = NewPublishPageRequest()
	}
	response = NewPublishPageResponse()
	err = c.Send(request, response)
	return
}

func NewQueryComponentsRequest() (request *QueryComponentsRequest) {
	request = &QueryComponentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "QueryComponents")
	return
}

func NewQueryComponentsResponse() (response *QueryComponentsResponse) {
	response = &QueryComponentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组件列表
func (c *Client) QueryComponents(request *QueryComponentsRequest) (response *QueryComponentsResponse, err error) {
	if request == nil {
		request = NewQueryComponentsRequest()
	}
	response = NewQueryComponentsResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductCodeInPagesRequest() (request *QueryProductCodeInPagesRequest) {
	request = &QueryProductCodeInPagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "QueryProductCodeInPages")
	return
}

func NewQueryProductCodeInPagesResponse() (response *QueryProductCodeInPagesResponse) {
	response = &QueryProductCodeInPagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取页面中productCode
func (c *Client) QueryProductCodeInPages(request *QueryProductCodeInPagesRequest) (response *QueryProductCodeInPagesResponse, err error) {
	if request == nil {
		request = NewQueryProductCodeInPagesRequest()
	}
	response = NewQueryProductCodeInPagesResponse()
	err = c.Send(request, response)
	return
}

func NewSetAuthRequest() (request *SetAuthRequest) {
	request = &SetAuthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "SetAuth")
	return
}

func NewSetAuthResponse() (response *SetAuthResponse) {
	response = &SetAuthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置免密访问
func (c *Client) SetAuth(request *SetAuthRequest) (response *SetAuthResponse, err error) {
	if request == nil {
		request = NewSetAuthRequest()
	}
	response = NewSetAuthResponse()
	err = c.Send(request, response)
	return
}

func NewCopyPageRequest() (request *CopyPageRequest) {
	request = &CopyPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "CopyPage")
	return
}

func NewCopyPageResponse() (response *CopyPageResponse) {
	response = &CopyPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 复制页面
func (c *Client) CopyPage(request *CopyPageRequest) (response *CopyPageResponse, err error) {
	if request == nil {
		request = NewCopyPageRequest()
	}
	response = NewCopyPageResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePageRequest() (request *DeletePageRequest) {
	request = &DeletePageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "DeletePage")
	return
}

func NewDeletePageResponse() (response *DeletePageResponse) {
	response = &DeletePageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除页面
func (c *Client) DeletePage(request *DeletePageRequest) (response *DeletePageResponse, err error) {
	if request == nil {
		request = NewDeletePageRequest()
	}
	response = NewDeletePageResponse()
	err = c.Send(request, response)
	return
}

func NewQueryPagesRequest() (request *QueryPagesRequest) {
	request = &QueryPagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "QueryPages")
	return
}

func NewQueryPagesResponse() (response *QueryPagesResponse) {
	response = &QueryPagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取页面列表
func (c *Client) QueryPages(request *QueryPagesRequest) (response *QueryPagesResponse, err error) {
	if request == nil {
		request = NewQueryPagesRequest()
	}
	response = NewQueryPagesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProjectRequest() (request *DeleteProjectRequest) {
	request = &DeleteProjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "DeleteProject")
	return
}

func NewDeleteProjectResponse() (response *DeleteProjectResponse) {
	response = &DeleteProjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除项目
func (c *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
	if request == nil {
		request = NewDeleteProjectRequest()
	}
	response = NewDeleteProjectResponse()
	err = c.Send(request, response)
	return
}

func NewRefreshAuthRequest() (request *RefreshAuthRequest) {
	request = &RefreshAuthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "RefreshAuth")
	return
}

func NewRefreshAuthResponse() (response *RefreshAuthResponse) {
	response = &RefreshAuthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 刷新项目接口免密策略
func (c *Client) RefreshAuth(request *RefreshAuthRequest) (response *RefreshAuthResponse, err error) {
	if request == nil {
		request = NewRefreshAuthRequest()
	}
	response = NewRefreshAuthResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDataSourceRequest() (request *ModifyDataSourceRequest) {
	request = &ModifyDataSourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "ModifyDataSource")
	return
}

func NewModifyDataSourceResponse() (response *ModifyDataSourceResponse) {
	response = &ModifyDataSourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑数据源
func (c *Client) ModifyDataSource(request *ModifyDataSourceRequest) (response *ModifyDataSourceResponse, err error) {
	if request == nil {
		request = NewModifyDataSourceRequest()
	}
	response = NewModifyDataSourceResponse()
	err = c.Send(request, response)
	return
}

func NewBindProjectWithPageRequest() (request *BindProjectWithPageRequest) {
	request = &BindProjectWithPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "BindProjectWithPage")
	return
}

func NewBindProjectWithPageResponse() (response *BindProjectWithPageResponse) {
	response = &BindProjectWithPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 项目绑定页面
func (c *Client) BindProjectWithPage(request *BindProjectWithPageRequest) (response *BindProjectWithPageResponse, err error) {
	if request == nil {
		request = NewBindProjectWithPageRequest()
	}
	response = NewBindProjectWithPageResponse()
	err = c.Send(request, response)
	return
}

func NewQueryDataSouceDataRequest() (request *QueryDataSouceDataRequest) {
	request = &QueryDataSouceDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "QueryDataSouceData")
	return
}

func NewQueryDataSouceDataResponse() (response *QueryDataSouceDataResponse) {
	response = &QueryDataSouceDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取独立数据源接口数据
func (c *Client) QueryDataSouceData(request *QueryDataSouceDataRequest) (response *QueryDataSouceDataResponse, err error) {
	if request == nil {
		request = NewQueryDataSouceDataRequest()
	}
	response = NewQueryDataSouceDataResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProjectRequest() (request *CreateProjectRequest) {
	request = &CreateProjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "CreateProject")
	return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
	response = &CreateProjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建项目
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
	if request == nil {
		request = NewCreateProjectRequest()
	}
	response = NewCreateProjectResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDataSourceRequest() (request *DeleteDataSourceRequest) {
	request = &DeleteDataSourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "DeleteDataSource")
	return
}

func NewDeleteDataSourceResponse() (response *DeleteDataSourceResponse) {
	response = &DeleteDataSourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除数据源
func (c *Client) DeleteDataSource(request *DeleteDataSourceRequest) (response *DeleteDataSourceResponse, err error) {
	if request == nil {
		request = NewDeleteDataSourceRequest()
	}
	response = NewDeleteDataSourceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPageRequest() (request *ModifyPageRequest) {
	request = &ModifyPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "ModifyPage")
	return
}

func NewModifyPageResponse() (response *ModifyPageResponse) {
	response = &ModifyPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改页面信息
func (c *Client) ModifyPage(request *ModifyPageRequest) (response *ModifyPageResponse, err error) {
	if request == nil {
		request = NewModifyPageRequest()
	}
	response = NewModifyPageResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDataSourceRequest() (request *CreateDataSourceRequest) {
	request = &CreateDataSourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "CreateDataSource")
	return
}

func NewCreateDataSourceResponse() (response *CreateDataSourceResponse) {
	response = &CreateDataSourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建数据源
func (c *Client) CreateDataSource(request *CreateDataSourceRequest) (response *CreateDataSourceResponse, err error) {
	if request == nil {
		request = NewCreateDataSourceRequest()
	}
	response = NewCreateDataSourceResponse()
	err = c.Send(request, response)
	return
}

func NewApiProxyRequest() (request *ApiProxyRequest) {
	request = &ApiProxyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "ApiProxy")
	return
}

func NewApiProxyResponse() (response *ApiProxyResponse) {
	response = &ApiProxyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取数据源接口数据
func (c *Client) ApiProxy(request *ApiProxyRequest) (response *ApiProxyResponse, err error) {
	if request == nil {
		request = NewApiProxyRequest()
	}
	response = NewApiProxyResponse()
	err = c.Send(request, response)
	return
}

func NewQueryDataSourceRequest() (request *QueryDataSourceRequest) {
	request = &QueryDataSourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "QueryDataSource")
	return
}

func NewQueryDataSourceResponse() (response *QueryDataSourceResponse) {
	response = &QueryDataSourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询数据源
func (c *Client) QueryDataSource(request *QueryDataSourceRequest) (response *QueryDataSourceResponse, err error) {
	if request == nil {
		request = NewQueryDataSourceRequest()
	}
	response = NewQueryDataSourceResponse()
	err = c.Send(request, response)
	return
}

func NewConfigProjectRequest() (request *ConfigProjectRequest) {
	request = &ConfigProjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "ConfigProject")
	return
}

func NewConfigProjectResponse() (response *ConfigProjectResponse) {
	response = &ConfigProjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置项目
func (c *Client) ConfigProject(request *ConfigProjectRequest) (response *ConfigProjectResponse, err error) {
	if request == nil {
		request = NewConfigProjectRequest()
	}
	response = NewConfigProjectResponse()
	err = c.Send(request, response)
	return
}

func NewSetConfPageRequest() (request *SetConfPageRequest) {
	request = &SetConfPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "SetConfPage")
	return
}

func NewSetConfPageResponse() (response *SetConfPageResponse) {
	response = &SetConfPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置页面基本配置
func (c *Client) SetConfPage(request *SetConfPageRequest) (response *SetConfPageResponse, err error) {
	if request == nil {
		request = NewSetConfPageRequest()
	}
	response = NewSetConfPageResponse()
	err = c.Send(request, response)
	return
}

func NewQueryMiddlePlatformRequest() (request *QueryMiddlePlatformRequest) {
	request = &QueryMiddlePlatformRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "QueryMiddlePlatform")
	return
}

func NewQueryMiddlePlatformResponse() (response *QueryMiddlePlatformResponse) {
	response = &QueryMiddlePlatformResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取中台数据源数据
func (c *Client) QueryMiddlePlatform(request *QueryMiddlePlatformRequest) (response *QueryMiddlePlatformResponse, err error) {
	if request == nil {
		request = NewQueryMiddlePlatformRequest()
	}
	response = NewQueryMiddlePlatformResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePageRequest() (request *CreatePageRequest) {
	request = &CreatePageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "CreatePage")
	return
}

func NewCreatePageResponse() (response *CreatePageResponse) {
	response = &CreatePageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建页面
func (c *Client) CreatePage(request *CreatePageRequest) (response *CreatePageResponse, err error) {
	if request == nil {
		request = NewCreatePageRequest()
	}
	response = NewCreatePageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProjectRequest() (request *ModifyProjectRequest) {
	request = &ModifyProjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "ModifyProject")
	return
}

func NewModifyProjectResponse() (response *ModifyProjectResponse) {
	response = &ModifyProjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑项目
func (c *Client) ModifyProject(request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
	if request == nil {
		request = NewModifyProjectRequest()
	}
	response = NewModifyProjectResponse()
	err = c.Send(request, response)
	return
}

func NewQueryTagsRequest() (request *QueryTagsRequest) {
	request = &QueryTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "QueryTags")
	return
}

func NewQueryTagsResponse() (response *QueryTagsResponse) {
	response = &QueryTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取标签列表
func (c *Client) QueryTags(request *QueryTagsRequest) (response *QueryTagsResponse, err error) {
	if request == nil {
		request = NewQueryTagsRequest()
	}
	response = NewQueryTagsResponse()
	err = c.Send(request, response)
	return
}

func NewQueryDataSourcesRequest() (request *QueryDataSourcesRequest) {
	request = &QueryDataSourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "QueryDataSources")
	return
}

func NewQueryDataSourcesResponse() (response *QueryDataSourcesResponse) {
	response = &QueryDataSourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取数据源列表
func (c *Client) QueryDataSources(request *QueryDataSourcesRequest) (response *QueryDataSourcesResponse, err error) {
	if request == nil {
		request = NewQueryDataSourcesRequest()
	}
	response = NewQueryDataSourcesResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProjectsRequest() (request *QueryProjectsRequest) {
	request = &QueryProjectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdv", APIVersion, "QueryProjects")
	return
}

func NewQueryProjectsResponse() (response *QueryProjectsResponse) {
	response = &QueryProjectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取项目列表
func (c *Client) QueryProjects(request *QueryProjectsRequest) (response *QueryProjectsResponse, err error) {
	if request == nil {
		request = NewQueryProjectsRequest()
	}
	response = NewQueryProjectsResponse()
	err = c.Send(request, response)
	return
}
