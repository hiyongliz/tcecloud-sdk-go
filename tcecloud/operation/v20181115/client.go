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

package v20181115

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2018-11-15"

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

func NewGetAnnounceClassRequest() (request *GetAnnounceClassRequest) {
	request = &GetAnnounceClassRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "GetAnnounceClass")
	return
}

func NewGetAnnounceClassResponse() (response *GetAnnounceClassResponse) {
	response = &GetAnnounceClassResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询公告分类信息
func (c *Client) GetAnnounceClass(request *GetAnnounceClassRequest) (response *GetAnnounceClassResponse, err error) {
	if request == nil {
		request = NewGetAnnounceClassRequest()
	}
	response = NewGetAnnounceClassResponse()
	err = c.Send(request, response)
	return
}

func NewGetAnnounceContentRequest() (request *GetAnnounceContentRequest) {
	request = &GetAnnounceContentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "GetAnnounceContent")
	return
}

func NewGetAnnounceContentResponse() (response *GetAnnounceContentResponse) {
	response = &GetAnnounceContentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询公告内容
func (c *Client) GetAnnounceContent(request *GetAnnounceContentRequest) (response *GetAnnounceContentResponse, err error) {
	if request == nil {
		request = NewGetAnnounceContentRequest()
	}
	response = NewGetAnnounceContentResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAnnounceClassRequest() (request *ModifyAnnounceClassRequest) {
	request = &ModifyAnnounceClassRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "ModifyAnnounceClass")
	return
}

func NewModifyAnnounceClassResponse() (response *ModifyAnnounceClassResponse) {
	response = &ModifyAnnounceClassResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于修改公告分类
func (c *Client) ModifyAnnounceClass(request *ModifyAnnounceClassRequest) (response *ModifyAnnounceClassResponse, err error) {
	if request == nil {
		request = NewModifyAnnounceClassRequest()
	}
	response = NewModifyAnnounceClassResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteHostAreaRequest() (request *DeleteHostAreaRequest) {
	request = &DeleteHostAreaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "DeleteHostArea")
	return
}

func NewDeleteHostAreaResponse() (response *DeleteHostAreaResponse) {
	response = &DeleteHostAreaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于删除机房地域信息
func (c *Client) DeleteHostArea(request *DeleteHostAreaRequest) (response *DeleteHostAreaResponse, err error) {
	if request == nil {
		request = NewDeleteHostAreaRequest()
	}
	response = NewDeleteHostAreaResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAnnounceLinkRequest() (request *ModifyAnnounceLinkRequest) {
	request = &ModifyAnnounceLinkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "ModifyAnnounceLink")
	return
}

func NewModifyAnnounceLinkResponse() (response *ModifyAnnounceLinkResponse) {
	response = &ModifyAnnounceLinkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改公告链接信息
func (c *Client) ModifyAnnounceLink(request *ModifyAnnounceLinkRequest) (response *ModifyAnnounceLinkResponse, err error) {
	if request == nil {
		request = NewModifyAnnounceLinkRequest()
	}
	response = NewModifyAnnounceLinkResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAnnounceLinkRequest() (request *CreateAnnounceLinkRequest) {
	request = &CreateAnnounceLinkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "CreateAnnounceLink")
	return
}

func NewCreateAnnounceLinkResponse() (response *CreateAnnounceLinkResponse) {
	response = &CreateAnnounceLinkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建公告链接
func (c *Client) CreateAnnounceLink(request *CreateAnnounceLinkRequest) (response *CreateAnnounceLinkResponse, err error) {
	if request == nil {
		request = NewCreateAnnounceLinkRequest()
	}
	response = NewCreateAnnounceLinkResponse()
	err = c.Send(request, response)
	return
}

func NewGetAnnounceLinkRequest() (request *GetAnnounceLinkRequest) {
	request = &GetAnnounceLinkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "GetAnnounceLink")
	return
}

func NewGetAnnounceLinkResponse() (response *GetAnnounceLinkResponse) {
	response = &GetAnnounceLinkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取公告链接信息
func (c *Client) GetAnnounceLink(request *GetAnnounceLinkRequest) (response *GetAnnounceLinkResponse, err error) {
	if request == nil {
		request = NewGetAnnounceLinkRequest()
	}
	response = NewGetAnnounceLinkResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAnnounceRequest() (request *DeleteAnnounceRequest) {
	request = &DeleteAnnounceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "DeleteAnnounce")
	return
}

func NewDeleteAnnounceResponse() (response *DeleteAnnounceResponse) {
	response = &DeleteAnnounceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除公告信息
func (c *Client) DeleteAnnounce(request *DeleteAnnounceRequest) (response *DeleteAnnounceResponse, err error) {
	if request == nil {
		request = NewDeleteAnnounceRequest()
	}
	response = NewDeleteAnnounceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostAreaRequest() (request *ModifyHostAreaRequest) {
	request = &ModifyHostAreaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "ModifyHostArea")
	return
}

func NewModifyHostAreaResponse() (response *ModifyHostAreaResponse) {
	response = &ModifyHostAreaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改公告地域信息
func (c *Client) ModifyHostArea(request *ModifyHostAreaRequest) (response *ModifyHostAreaResponse, err error) {
	if request == nil {
		request = NewModifyHostAreaRequest()
	}
	response = NewModifyHostAreaResponse()
	err = c.Send(request, response)
	return
}

func NewCreateHostAreaRequest() (request *CreateHostAreaRequest) {
	request = &CreateHostAreaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "CreateHostArea")
	return
}

func NewCreateHostAreaResponse() (response *CreateHostAreaResponse) {
	response = &CreateHostAreaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于创建机房地域
func (c *Client) CreateHostArea(request *CreateHostAreaRequest) (response *CreateHostAreaResponse, err error) {
	if request == nil {
		request = NewCreateHostAreaRequest()
	}
	response = NewCreateHostAreaResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAnnounceContentRequest() (request *CreateAnnounceContentRequest) {
	request = &CreateAnnounceContentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "CreateAnnounceContent")
	return
}

func NewCreateAnnounceContentResponse() (response *CreateAnnounceContentResponse) {
	response = &CreateAnnounceContentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于创建公告内容
func (c *Client) CreateAnnounceContent(request *CreateAnnounceContentRequest) (response *CreateAnnounceContentResponse, err error) {
	if request == nil {
		request = NewCreateAnnounceContentRequest()
	}
	response = NewCreateAnnounceContentResponse()
	err = c.Send(request, response)
	return
}

func NewGetAnnounceListRequest() (request *GetAnnounceListRequest) {
	request = &GetAnnounceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "GetAnnounceList")
	return
}

func NewGetAnnounceListResponse() (response *GetAnnounceListResponse) {
	response = &GetAnnounceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询公告列表
func (c *Client) GetAnnounceList(request *GetAnnounceListRequest) (response *GetAnnounceListResponse, err error) {
	if request == nil {
		request = NewGetAnnounceListRequest()
	}
	response = NewGetAnnounceListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAnnounceClassRequest() (request *CreateAnnounceClassRequest) {
	request = &CreateAnnounceClassRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "CreateAnnounceClass")
	return
}

func NewCreateAnnounceClassResponse() (response *CreateAnnounceClassResponse) {
	response = &CreateAnnounceClassResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于创建公告分类
func (c *Client) CreateAnnounceClass(request *CreateAnnounceClassRequest) (response *CreateAnnounceClassResponse, err error) {
	if request == nil {
		request = NewCreateAnnounceClassRequest()
	}
	response = NewCreateAnnounceClassResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAnnounceClassRequest() (request *DeleteAnnounceClassRequest) {
	request = &DeleteAnnounceClassRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "DeleteAnnounceClass")
	return
}

func NewDeleteAnnounceClassResponse() (response *DeleteAnnounceClassResponse) {
	response = &DeleteAnnounceClassResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于删除公告分类
func (c *Client) DeleteAnnounceClass(request *DeleteAnnounceClassRequest) (response *DeleteAnnounceClassResponse, err error) {
	if request == nil {
		request = NewDeleteAnnounceClassRequest()
	}
	response = NewDeleteAnnounceClassResponse()
	err = c.Send(request, response)
	return
}

func NewGetHostAreaRequest() (request *GetHostAreaRequest) {
	request = &GetHostAreaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "GetHostArea")
	return
}

func NewGetHostAreaResponse() (response *GetHostAreaResponse) {
	response = &GetHostAreaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询机房地域
func (c *Client) GetHostArea(request *GetHostAreaRequest) (response *GetHostAreaResponse, err error) {
	if request == nil {
		request = NewGetHostAreaRequest()
	}
	response = NewGetHostAreaResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAnnounceRequest() (request *ModifyAnnounceRequest) {
	request = &ModifyAnnounceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("operation", APIVersion, "ModifyAnnounce")
	return
}

func NewModifyAnnounceResponse() (response *ModifyAnnounceResponse) {
	response = &ModifyAnnounceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改公告信息
func (c *Client) ModifyAnnounce(request *ModifyAnnounceRequest) (response *ModifyAnnounceResponse, err error) {
	if request == nil {
		request = NewModifyAnnounceRequest()
	}
	response = NewModifyAnnounceResponse()
	err = c.Send(request, response)
	return
}
