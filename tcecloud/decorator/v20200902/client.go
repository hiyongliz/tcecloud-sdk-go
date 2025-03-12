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

package v20200902

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-09-02"

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

func NewQueryPreviewRequest() (request *QueryPreviewRequest) {
	request = &QueryPreviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("decorator", APIVersion, "QueryPreview")
	return
}

func NewQueryPreviewResponse() (response *QueryPreviewResponse) {
	response = &QueryPreviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询预览页面模板
func (c *Client) QueryPreview(request *QueryPreviewRequest) (response *QueryPreviewResponse, err error) {
	if request == nil {
		request = NewQueryPreviewRequest()
	}
	response = NewQueryPreviewResponse()
	err = c.Send(request, response)
	return
}

func NewDelPageRequest() (request *DelPageRequest) {
	request = &DelPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("decorator", APIVersion, "DelPage")
	return
}

func NewDelPageResponse() (response *DelPageResponse) {
	response = &DelPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除页面模板
func (c *Client) DelPage(request *DelPageRequest) (response *DelPageResponse, err error) {
	if request == nil {
		request = NewDelPageRequest()
	}
	response = NewDelPageResponse()
	err = c.Send(request, response)
	return
}

func NewCommitPreviewRequest() (request *CommitPreviewRequest) {
	request = &CommitPreviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("decorator", APIVersion, "CommitPreview")
	return
}

func NewCommitPreviewResponse() (response *CommitPreviewResponse) {
	response = &CommitPreviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提交预览页面模板
func (c *Client) CommitPreview(request *CommitPreviewRequest) (response *CommitPreviewResponse, err error) {
	if request == nil {
		request = NewCommitPreviewRequest()
	}
	response = NewCommitPreviewResponse()
	err = c.Send(request, response)
	return
}

func NewPutPreviewRequest() (request *PutPreviewRequest) {
	request = &PutPreviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("decorator", APIVersion, "PutPreview")
	return
}

func NewPutPreviewResponse() (response *PutPreviewResponse) {
	response = &PutPreviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑预览页面模板
func (c *Client) PutPreview(request *PutPreviewRequest) (response *PutPreviewResponse, err error) {
	if request == nil {
		request = NewPutPreviewRequest()
	}
	response = NewPutPreviewResponse()
	err = c.Send(request, response)
	return
}

func NewPutComponentRequest() (request *PutComponentRequest) {
	request = &PutComponentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("decorator", APIVersion, "PutComponent")
	return
}

func NewPutComponentResponse() (response *PutComponentResponse) {
	response = &PutComponentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑组件模板
func (c *Client) PutComponent(request *PutComponentRequest) (response *PutComponentResponse, err error) {
	if request == nil {
		request = NewPutComponentRequest()
	}
	response = NewPutComponentResponse()
	err = c.Send(request, response)
	return
}

func NewQueryComponentsRequest() (request *QueryComponentsRequest) {
	request = &QueryComponentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("decorator", APIVersion, "QueryComponents")
	return
}

func NewQueryComponentsResponse() (response *QueryComponentsResponse) {
	response = &QueryComponentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询组件模板
func (c *Client) QueryComponents(request *QueryComponentsRequest) (response *QueryComponentsResponse, err error) {
	if request == nil {
		request = NewQueryComponentsRequest()
	}
	response = NewQueryComponentsResponse()
	err = c.Send(request, response)
	return
}

func NewDelComponentRequest() (request *DelComponentRequest) {
	request = &DelComponentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("decorator", APIVersion, "DelComponent")
	return
}

func NewDelComponentResponse() (response *DelComponentResponse) {
	response = &DelComponentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除组件模板
func (c *Client) DelComponent(request *DelComponentRequest) (response *DelComponentResponse, err error) {
	if request == nil {
		request = NewDelComponentRequest()
	}
	response = NewDelComponentResponse()
	err = c.Send(request, response)
	return
}

func NewDelPreviewRequest() (request *DelPreviewRequest) {
	request = &DelPreviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("decorator", APIVersion, "DelPreview")
	return
}

func NewDelPreviewResponse() (response *DelPreviewResponse) {
	response = &DelPreviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除页面预览
func (c *Client) DelPreview(request *DelPreviewRequest) (response *DelPreviewResponse, err error) {
	if request == nil {
		request = NewDelPreviewRequest()
	}
	response = NewDelPreviewResponse()
	err = c.Send(request, response)
	return
}

func NewQueryPagesRequest() (request *QueryPagesRequest) {
	request = &QueryPagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("decorator", APIVersion, "QueryPages")
	return
}

func NewQueryPagesResponse() (response *QueryPagesResponse) {
	response = &QueryPagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询页面模板
func (c *Client) QueryPages(request *QueryPagesRequest) (response *QueryPagesResponse, err error) {
	if request == nil {
		request = NewQueryPagesRequest()
	}
	response = NewQueryPagesResponse()
	err = c.Send(request, response)
	return
}
