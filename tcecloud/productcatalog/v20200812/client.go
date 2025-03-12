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

package v20200812

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-08-12"

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

func NewDeleteUserProductCatalogsRequest() (request *DeleteUserProductCatalogsRequest) {
	request = &DeleteUserProductCatalogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("productcatalog", APIVersion, "DeleteUserProductCatalogs")
	return
}

func NewDeleteUserProductCatalogsResponse() (response *DeleteUserProductCatalogsResponse) {
	response = &DeleteUserProductCatalogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除租户产品目录记录
func (c *Client) DeleteUserProductCatalogs(request *DeleteUserProductCatalogsRequest) (response *DeleteUserProductCatalogsResponse, err error) {
	if request == nil {
		request = NewDeleteUserProductCatalogsRequest()
	}
	response = NewDeleteUserProductCatalogsResponse()
	err = c.Send(request, response)
	return
}

func NewListUserProductCatalogsRequest() (request *ListUserProductCatalogsRequest) {
	request = &ListUserProductCatalogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("productcatalog", APIVersion, "ListUserProductCatalogs")
	return
}

func NewListUserProductCatalogsResponse() (response *ListUserProductCatalogsResponse) {
	response = &ListUserProductCatalogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取定义了产品目录的租户列表
func (c *Client) ListUserProductCatalogs(request *ListUserProductCatalogsRequest) (response *ListUserProductCatalogsResponse, err error) {
	if request == nil {
		request = NewListUserProductCatalogsRequest()
	}
	response = NewListUserProductCatalogsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProductStatusRequest() (request *UpdateProductStatusRequest) {
	request = &UpdateProductStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("productcatalog", APIVersion, "UpdateProductStatus")
	return
}

func NewUpdateProductStatusResponse() (response *UpdateProductStatusResponse) {
	response = &UpdateProductStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改产品状态
func (c *Client) UpdateProductStatus(request *UpdateProductStatusRequest) (response *UpdateProductStatusResponse, err error) {
	if request == nil {
		request = NewUpdateProductStatusRequest()
	}
	response = NewUpdateProductStatusResponse()
	err = c.Send(request, response)
	return
}

func NewAddUserProductCatalogsRequest() (request *AddUserProductCatalogsRequest) {
	request = &AddUserProductCatalogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("productcatalog", APIVersion, "AddUserProductCatalogs")
	return
}

func NewAddUserProductCatalogsResponse() (response *AddUserProductCatalogsResponse) {
	response = &AddUserProductCatalogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增租户产品目录
func (c *Client) AddUserProductCatalogs(request *AddUserProductCatalogsRequest) (response *AddUserProductCatalogsResponse, err error) {
	if request == nil {
		request = NewAddUserProductCatalogsRequest()
	}
	response = NewAddUserProductCatalogsResponse()
	err = c.Send(request, response)
	return
}

func NewGetGrayscaleListByUuidRequest() (request *GetGrayscaleListByUuidRequest) {
	request = &GetGrayscaleListByUuidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("productcatalog", APIVersion, "GetGrayscaleListByUuid")
	return
}

func NewGetGrayscaleListByUuidResponse() (response *GetGrayscaleListByUuidResponse) {
	response = &GetGrayscaleListByUuidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过产品uuid获取灰度名单信息
func (c *Client) GetGrayscaleListByUuid(request *GetGrayscaleListByUuidRequest) (response *GetGrayscaleListByUuidResponse, err error) {
	if request == nil {
		request = NewGetGrayscaleListByUuidRequest()
	}
	response = NewGetGrayscaleListByUuidResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserProductCatalogsRequest() (request *ModifyUserProductCatalogsRequest) {
	request = &ModifyUserProductCatalogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("productcatalog", APIVersion, "ModifyUserProductCatalogs")
	return
}

func NewModifyUserProductCatalogsResponse() (response *ModifyUserProductCatalogsResponse) {
	response = &ModifyUserProductCatalogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改指定租户的产品目录
func (c *Client) ModifyUserProductCatalogs(request *ModifyUserProductCatalogsRequest) (response *ModifyUserProductCatalogsResponse, err error) {
	if request == nil {
		request = NewModifyUserProductCatalogsRequest()
	}
	response = NewModifyUserProductCatalogsResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductCatalogInfoRequest() (request *GetProductCatalogInfoRequest) {
	request = &GetProductCatalogInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("productcatalog", APIVersion, "GetProductCatalogInfo")
	return
}

func NewGetProductCatalogInfoResponse() (response *GetProductCatalogInfoResponse) {
	response = &GetProductCatalogInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品目录列表
func (c *Client) GetProductCatalogInfo(request *GetProductCatalogInfoRequest) (response *GetProductCatalogInfoResponse, err error) {
	if request == nil {
		request = NewGetProductCatalogInfoRequest()
	}
	response = NewGetProductCatalogInfoResponse()
	err = c.Send(request, response)
	return
}
