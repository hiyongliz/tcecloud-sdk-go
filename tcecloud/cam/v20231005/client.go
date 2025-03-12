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

package v20231005

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2023-10-05"

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

func NewVerifyUserProductMenusRequest() (request *VerifyUserProductMenusRequest) {
	request = &VerifyUserProductMenusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "VerifyUserProductMenus")
	return
}

func NewVerifyUserProductMenusResponse() (response *VerifyUserProductMenusResponse) {
	response = &VerifyUserProductMenusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 验证用户能否访问产品菜单
func (c *Client) VerifyUserProductMenus(request *VerifyUserProductMenusRequest) (response *VerifyUserProductMenusResponse, err error) {
	if request == nil {
		request = NewVerifyUserProductMenusRequest()
	}
	response = NewVerifyUserProductMenusResponse()
	err = c.Send(request, response)
	return
}

func NewAddUserProductMenusRequest() (request *AddUserProductMenusRequest) {
	request = &AddUserProductMenusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "AddUserProductMenus")
	return
}

func NewAddUserProductMenusResponse() (response *AddUserProductMenusResponse) {
	response = &AddUserProductMenusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建用户产品菜单
func (c *Client) AddUserProductMenus(request *AddUserProductMenusRequest) (response *AddUserProductMenusResponse, err error) {
	if request == nil {
		request = NewAddUserProductMenusRequest()
	}
	response = NewAddUserProductMenusResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUserProductMenusRequest() (request *DeleteUserProductMenusRequest) {
	request = &DeleteUserProductMenusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DeleteUserProductMenus")
	return
}

func NewDeleteUserProductMenusResponse() (response *DeleteUserProductMenusResponse) {
	response = &DeleteUserProductMenusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户产品菜单
func (c *Client) DeleteUserProductMenus(request *DeleteUserProductMenusRequest) (response *DeleteUserProductMenusResponse, err error) {
	if request == nil {
		request = NewDeleteUserProductMenusRequest()
	}
	response = NewDeleteUserProductMenusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserProductMenusRequest() (request *DescribeUserProductMenusRequest) {
	request = &DescribeUserProductMenusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DescribeUserProductMenus")
	return
}

func NewDescribeUserProductMenusResponse() (response *DescribeUserProductMenusResponse) {
	response = &DescribeUserProductMenusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户产品菜单
func (c *Client) DescribeUserProductMenus(request *DescribeUserProductMenusRequest) (response *DescribeUserProductMenusResponse, err error) {
	if request == nil {
		request = NewDescribeUserProductMenusRequest()
	}
	response = NewDescribeUserProductMenusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserProductMenusRequest() (request *ModifyUserProductMenusRequest) {
	request = &ModifyUserProductMenusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ModifyUserProductMenus")
	return
}

func NewModifyUserProductMenusResponse() (response *ModifyUserProductMenusResponse) {
	response = &ModifyUserProductMenusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户产品菜单
func (c *Client) ModifyUserProductMenus(request *ModifyUserProductMenusRequest) (response *ModifyUserProductMenusResponse, err error) {
	if request == nil {
		request = NewModifyUserProductMenusRequest()
	}
	response = NewModifyUserProductMenusResponse()
	err = c.Send(request, response)
	return
}
