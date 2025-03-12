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

package v20200901

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-09-01"

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

func NewGetUserByLoginUinRequest() (request *GetUserByLoginUinRequest) {
	request = &GetUserByLoginUinRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsauth", APIVersion, "GetUserByLoginUin")
	return
}

func NewGetUserByLoginUinResponse() (response *GetUserByLoginUinResponse) {
	response = &GetUserByLoginUinResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询登录用户信息
func (c *Client) GetUserByLoginUin(request *GetUserByLoginUinRequest) (response *GetUserByLoginUinResponse, err error) {
	if request == nil {
		request = NewGetUserByLoginUinRequest()
	}
	response = NewGetUserByLoginUinResponse()
	err = c.Send(request, response)
	return
}

func NewListProjectsRequest() (request *ListProjectsRequest) {
	request = &ListProjectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcsauth", APIVersion, "ListProjects")
	return
}

func NewListProjectsResponse() (response *ListProjectsResponse) {
	response = &ListProjectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出项目
func (c *Client) ListProjects(request *ListProjectsRequest) (response *ListProjectsResponse, err error) {
	if request == nil {
		request = NewListProjectsRequest()
	}
	response = NewListProjectsResponse()
	err = c.Send(request, response)
	return
}
