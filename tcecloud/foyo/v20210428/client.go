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

package v20210428

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2021-04-28"

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

func NewQuerySwitchVersionsRequest() (request *QuerySwitchVersionsRequest) {
	request = &QuerySwitchVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("foyo", APIVersion, "QuerySwitchVersions")
	return
}

func NewQuerySwitchVersionsResponse() (response *QuerySwitchVersionsResponse) {
	response = &QuerySwitchVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询开关的历史版本接口
func (c *Client) QuerySwitchVersions(request *QuerySwitchVersionsRequest) (response *QuerySwitchVersionsResponse, err error) {
	if request == nil {
		request = NewQuerySwitchVersionsRequest()
	}
	response = NewQuerySwitchVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewRollbackToVersionRequest() (request *RollbackToVersionRequest) {
	request = &RollbackToVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("foyo", APIVersion, "RollbackToVersion")
	return
}

func NewRollbackToVersionResponse() (response *RollbackToVersionResponse) {
	response = &RollbackToVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回滚开关至某版本接口
func (c *Client) RollbackToVersion(request *RollbackToVersionRequest) (response *RollbackToVersionResponse, err error) {
	if request == nil {
		request = NewRollbackToVersionRequest()
	}
	response = NewRollbackToVersionResponse()
	err = c.Send(request, response)
	return
}

func NewSearchSwitchListRequest() (request *SearchSwitchListRequest) {
	request = &SearchSwitchListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("foyo", APIVersion, "SearchSwitchList")
	return
}

func NewSearchSwitchListResponse() (response *SearchSwitchListResponse) {
	response = &SearchSwitchListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过关键词查询某ns下的开关列表
func (c *Client) SearchSwitchList(request *SearchSwitchListRequest) (response *SearchSwitchListResponse, err error) {
	if request == nil {
		request = NewSearchSwitchListRequest()
	}
	response = NewSearchSwitchListResponse()
	err = c.Send(request, response)
	return
}

func NewQueryNamespacesRequest() (request *QueryNamespacesRequest) {
	request = &QueryNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("foyo", APIVersion, "QueryNamespaces")
	return
}

func NewQueryNamespacesResponse() (response *QueryNamespacesResponse) {
	response = &QueryNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所有namespace接口
func (c *Client) QueryNamespaces(request *QueryNamespacesRequest) (response *QueryNamespacesResponse, err error) {
	if request == nil {
		request = NewQueryNamespacesRequest()
	}
	response = NewQueryNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewDrawSwitchesRequest() (request *DrawSwitchesRequest) {
	request = &DrawSwitchesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("foyo", APIVersion, "DrawSwitches")
	return
}

func NewDrawSwitchesResponse() (response *DrawSwitchesResponse) {
	response = &DrawSwitchesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用来配置一系列开关的接口
func (c *Client) DrawSwitches(request *DrawSwitchesRequest) (response *DrawSwitchesResponse, err error) {
	if request == nil {
		request = NewDrawSwitchesRequest()
	}
	response = NewDrawSwitchesResponse()
	err = c.Send(request, response)
	return
}

func NewQuerySwitchRequest() (request *QuerySwitchRequest) {
	request = &QuerySwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("foyo", APIVersion, "QuerySwitch")
	return
}

func NewQuerySwitchResponse() (response *QuerySwitchResponse) {
	response = &QuerySwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询开关接口
func (c *Client) QuerySwitch(request *QuerySwitchRequest) (response *QuerySwitchResponse, err error) {
	if request == nil {
		request = NewQuerySwitchRequest()
	}
	response = NewQuerySwitchResponse()
	err = c.Send(request, response)
	return
}

func NewQuerySwitchListRequest() (request *QuerySwitchListRequest) {
	request = &QuerySwitchListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("foyo", APIVersion, "QuerySwitchList")
	return
}

func NewQuerySwitchListResponse() (response *QuerySwitchListResponse) {
	response = &QuerySwitchListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询开关列表接口
func (c *Client) QuerySwitchList(request *QuerySwitchListRequest) (response *QuerySwitchListResponse, err error) {
	if request == nil {
		request = NewQuerySwitchListRequest()
	}
	response = NewQuerySwitchListResponse()
	err = c.Send(request, response)
	return
}

func NewSearchNamespacesRequest() (request *SearchNamespacesRequest) {
	request = &SearchNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("foyo", APIVersion, "SearchNamespaces")
	return
}

func NewSearchNamespacesResponse() (response *SearchNamespacesResponse) {
	response = &SearchNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过关键词查询开关ns
func (c *Client) SearchNamespaces(request *SearchNamespacesRequest) (response *SearchNamespacesResponse, err error) {
	if request == nil {
		request = NewSearchNamespacesRequest()
	}
	response = NewSearchNamespacesResponse()
	err = c.Send(request, response)
	return
}
