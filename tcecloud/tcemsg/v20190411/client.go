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

package v20190411

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-04-11"

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

func NewQueryPostRequest() (request *QueryPostRequest) {
	request = &QueryPostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "QueryPost")
	return
}

func NewQueryPostResponse() (response *QueryPostResponse) {
	response = &QueryPostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询首页公告列表
func (c *Client) QueryPost(request *QueryPostRequest) (response *QueryPostResponse, err error) {
	if request == nil {
		request = NewQueryPostRequest()
	}
	response = NewQueryPostResponse()
	err = c.Send(request, response)
	return
}

func NewDelUserCmgtSiteMsgRequest() (request *DelUserCmgtSiteMsgRequest) {
	request = &DelUserCmgtSiteMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "DelUserCmgtSiteMsg")
	return
}

func NewDelUserCmgtSiteMsgResponse() (response *DelUserCmgtSiteMsgResponse) {
	response = &DelUserCmgtSiteMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除运营端用户站内信
func (c *Client) DelUserCmgtSiteMsg(request *DelUserCmgtSiteMsgRequest) (response *DelUserCmgtSiteMsgResponse, err error) {
	if request == nil {
		request = NewDelUserCmgtSiteMsgRequest()
	}
	response = NewDelUserCmgtSiteMsgResponse()
	err = c.Send(request, response)
	return
}

func NewAddPostRequest() (request *AddPostRequest) {
	request = &AddPostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcemsg", APIVersion, "AddPost")
	return
}

func NewAddPostResponse() (response *AddPostResponse) {
	response = &AddPostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发布首页公告
func (c *Client) AddPost(request *AddPostRequest) (response *AddPostResponse, err error) {
	if request == nil {
		request = NewAddPostRequest()
	}
	response = NewAddPostResponse()
	err = c.Send(request, response)
	return
}
