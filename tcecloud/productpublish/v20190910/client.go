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

package v20190910

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-09-10"

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

func NewInsertOperateHistoryRequest() (request *InsertOperateHistoryRequest) {
	request = &InsertOperateHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("productpublish", APIVersion, "InsertOperateHistory")
	return
}

func NewInsertOperateHistoryResponse() (response *InsertOperateHistoryResponse) {
	response = &InsertOperateHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客制化产品化--发布管理--新增操作历史
func (c *Client) InsertOperateHistory(request *InsertOperateHistoryRequest) (response *InsertOperateHistoryResponse, err error) {
	if request == nil {
		request = NewInsertOperateHistoryRequest()
	}
	response = NewInsertOperateHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewSearchReleaseHistoryRequest() (request *SearchReleaseHistoryRequest) {
	request = &SearchReleaseHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("productpublish", APIVersion, "SearchReleaseHistory")
	return
}

func NewSearchReleaseHistoryResponse() (response *SearchReleaseHistoryResponse) {
	response = &SearchReleaseHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客制化产品化--发布管理--搜索发布历史
func (c *Client) SearchReleaseHistory(request *SearchReleaseHistoryRequest) (response *SearchReleaseHistoryResponse, err error) {
	if request == nil {
		request = NewSearchReleaseHistoryRequest()
	}
	response = NewSearchReleaseHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseRequest() (request *ReleaseRequest) {
	request = &ReleaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("productpublish", APIVersion, "Release")
	return
}

func NewReleaseResponse() (response *ReleaseResponse) {
	response = &ReleaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客制化产品化--发布管理--发布
func (c *Client) Release(request *ReleaseRequest) (response *ReleaseResponse, err error) {
	if request == nil {
		request = NewReleaseRequest()
	}
	response = NewReleaseResponse()
	err = c.Send(request, response)
	return
}

func NewSearchOperateHistoryRequest() (request *SearchOperateHistoryRequest) {
	request = &SearchOperateHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("productpublish", APIVersion, "SearchOperateHistory")
	return
}

func NewSearchOperateHistoryResponse() (response *SearchOperateHistoryResponse) {
	response = &SearchOperateHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客制化产品化--发布管理--搜索操作历史
func (c *Client) SearchOperateHistory(request *SearchOperateHistoryRequest) (response *SearchOperateHistoryResponse, err error) {
	if request == nil {
		request = NewSearchOperateHistoryRequest()
	}
	response = NewSearchOperateHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewOuterCustomizeProductHandlerRequest() (request *OuterCustomizeProductHandlerRequest) {
	request = &OuterCustomizeProductHandlerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("productpublish", APIVersion, "OuterCustomizeProductHandler")
	return
}

func NewOuterCustomizeProductHandlerResponse() (response *OuterCustomizeProductHandlerResponse) {
	response = &OuterCustomizeProductHandlerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客制化产品化操作产品菜单添加接口
func (c *Client) OuterCustomizeProductHandler(request *OuterCustomizeProductHandlerRequest) (response *OuterCustomizeProductHandlerResponse, err error) {
	if request == nil {
		request = NewOuterCustomizeProductHandlerRequest()
	}
	response = NewOuterCustomizeProductHandlerResponse()
	err = c.Send(request, response)
	return
}

func NewCustomDeployRequest() (request *CustomDeployRequest) {
	request = &CustomDeployRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("productpublish", APIVersion, "CustomDeploy")
	return
}

func NewCustomDeployResponse() (response *CustomDeployResponse) {
	response = &CustomDeployResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客制化产品化--发布管理--部署
func (c *Client) CustomDeploy(request *CustomDeployRequest) (response *CustomDeployResponse, err error) {
	if request == nil {
		request = NewCustomDeployRequest()
	}
	response = NewCustomDeployResponse()
	err = c.Send(request, response)
	return
}

func NewQueryReleaseStatusRequest() (request *QueryReleaseStatusRequest) {
	request = &QueryReleaseStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("productpublish", APIVersion, "QueryReleaseStatus")
	return
}

func NewQueryReleaseStatusResponse() (response *QueryReleaseStatusResponse) {
	response = &QueryReleaseStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询发布状态
func (c *Client) QueryReleaseStatus(request *QueryReleaseStatusRequest) (response *QueryReleaseStatusResponse, err error) {
	if request == nil {
		request = NewQueryReleaseStatusRequest()
	}
	response = NewQueryReleaseStatusResponse()
	err = c.Send(request, response)
	return
}

func NewSearchReleaseDetailRequest() (request *SearchReleaseDetailRequest) {
	request = &SearchReleaseDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("productpublish", APIVersion, "SearchReleaseDetail")
	return
}

func NewSearchReleaseDetailResponse() (response *SearchReleaseDetailResponse) {
	response = &SearchReleaseDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客制化产品化--发布管理--查询发布明细
func (c *Client) SearchReleaseDetail(request *SearchReleaseDetailRequest) (response *SearchReleaseDetailResponse, err error) {
	if request == nil {
		request = NewSearchReleaseDetailRequest()
	}
	response = NewSearchReleaseDetailResponse()
	err = c.Send(request, response)
	return
}
