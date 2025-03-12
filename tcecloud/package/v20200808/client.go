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

package v20200808

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-08-08"

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

func NewCreatePackageVersionRequest() (request *CreatePackageVersionRequest) {
	request = &CreatePackageVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("package", APIVersion, "CreatePackageVersion")
	return
}

func NewCreatePackageVersionResponse() (response *CreatePackageVersionResponse) {
	response = &CreatePackageVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 软件包新增版本
func (c *Client) CreatePackageVersion(request *CreatePackageVersionRequest) (response *CreatePackageVersionResponse, err error) {
	if request == nil {
		request = NewCreatePackageVersionRequest()
	}
	response = NewCreatePackageVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePackagesRequest() (request *DescribePackagesRequest) {
	request = &DescribePackagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("package", APIVersion, "DescribePackages")
	return
}

func NewDescribePackagesResponse() (response *DescribePackagesResponse) {
	response = &DescribePackagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看软件包列表
func (c *Client) DescribePackages(request *DescribePackagesRequest) (response *DescribePackagesResponse, err error) {
	if request == nil {
		request = NewDescribePackagesRequest()
	}
	response = NewDescribePackagesResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePackageVersionsRequest() (request *DeletePackageVersionsRequest) {
	request = &DeletePackageVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("package", APIVersion, "DeletePackageVersions")
	return
}

func NewDeletePackageVersionsResponse() (response *DeletePackageVersionsResponse) {
	response = &DeletePackageVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除软件包版本，支持一次删除多个版本，但只支持PkgUUID形式指定包
func (c *Client) DeletePackageVersions(request *DeletePackageVersionsRequest) (response *DeletePackageVersionsResponse, err error) {
	if request == nil {
		request = NewDeletePackageVersionsRequest()
	}
	response = NewDeletePackageVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePackageRequest() (request *CreatePackageRequest) {
	request = &CreatePackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("package", APIVersion, "CreatePackage")
	return
}

func NewCreatePackageResponse() (response *CreatePackageResponse) {
	response = &CreatePackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建软件包，不涉及具体的数据上传等信息
func (c *Client) CreatePackage(request *CreatePackageRequest) (response *CreatePackageResponse, err error) {
	if request == nil {
		request = NewCreatePackageRequest()
	}
	response = NewCreatePackageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEventsRequest() (request *DescribeEventsRequest) {
	request = &DescribeEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("package", APIVersion, "DescribeEvents")
	return
}

func NewDescribeEventsResponse() (response *DescribeEventsResponse) {
	response = &DescribeEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看包变更事件列表
func (c *Client) DescribeEvents(request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
	if request == nil {
		request = NewDescribeEventsRequest()
	}
	response = NewDescribeEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePackageInstancesRequest() (request *DescribePackageInstancesRequest) {
	request = &DescribePackageInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("package", APIVersion, "DescribePackageInstances")
	return
}

func NewDescribePackageInstancesResponse() (response *DescribePackageInstancesResponse) {
	response = &DescribePackageInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 支持通过Name、PkgUUID、Name+PkgVersion、PkgUUID+PkgVersion来查询其已经部署的实例列表信息
func (c *Client) DescribePackageInstances(request *DescribePackageInstancesRequest) (response *DescribePackageInstancesResponse, err error) {
	if request == nil {
		request = NewDescribePackageInstancesRequest()
	}
	response = NewDescribePackageInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePackageVersionsRequest() (request *DescribePackageVersionsRequest) {
	request = &DescribePackageVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("package", APIVersion, "DescribePackageVersions")
	return
}

func NewDescribePackageVersionsResponse() (response *DescribePackageVersionsResponse) {
	response = &DescribePackageVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 当Filters输入对应版本信息时返回单个版本详情，否则返回同个包的版本详情
func (c *Client) DescribePackageVersions(request *DescribePackageVersionsRequest) (response *DescribePackageVersionsResponse, err error) {
	if request == nil {
		request = NewDescribePackageVersionsRequest()
	}
	response = NewDescribePackageVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePackageRequest() (request *DeletePackageRequest) {
	request = &DeletePackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("package", APIVersion, "DeletePackage")
	return
}

func NewDeletePackageResponse() (response *DeletePackageResponse) {
	response = &DeletePackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除软件包，注意：当此软件包有未删除的版本时调用此接口将会报错
func (c *Client) DeletePackage(request *DeletePackageRequest) (response *DeletePackageResponse, err error) {
	if request == nil {
		request = NewDeletePackageRequest()
	}
	response = NewDeletePackageResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePackageInstancesRequest() (request *DeletePackageInstancesRequest) {
	request = &DeletePackageInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("package", APIVersion, "DeletePackageInstances")
	return
}

func NewDeletePackageInstancesResponse() (response *DeletePackageInstancesResponse) {
	response = &DeletePackageInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 支持通过InstanceID、PkgUUID+PkgVersion+SvrIP删除，且支持一次删除多个包的多个实例
func (c *Client) DeletePackageInstances(request *DeletePackageInstancesRequest) (response *DeletePackageInstancesResponse, err error) {
	if request == nil {
		request = NewDeletePackageInstancesRequest()
	}
	response = NewDeletePackageInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePackageRequest() (request *UpdatePackageRequest) {
	request = &UpdatePackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("package", APIVersion, "UpdatePackage")
	return
}

func NewUpdatePackageResponse() (response *UpdatePackageResponse) {
	response = &UpdatePackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 支持修改负责人、对应产品信息、描述三个字段
func (c *Client) UpdatePackage(request *UpdatePackageRequest) (response *UpdatePackageResponse, err error) {
	if request == nil {
		request = NewUpdatePackageRequest()
	}
	response = NewUpdatePackageResponse()
	err = c.Send(request, response)
	return
}

func NewRegisterPackageInstanceRequest() (request *RegisterPackageInstanceRequest) {
	request = &RegisterPackageInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("package", APIVersion, "RegisterPackageInstance")
	return
}

func NewRegisterPackageInstanceResponse() (response *RegisterPackageInstanceResponse) {
	response = &RegisterPackageInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 当流程引擎或者其他系统对包进行了发布安装，可以通过此接口进行包实例的注册
func (c *Client) RegisterPackageInstance(request *RegisterPackageInstanceRequest) (response *RegisterPackageInstanceResponse, err error) {
	if request == nil {
		request = NewRegisterPackageInstanceRequest()
	}
	response = NewRegisterPackageInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCSPDownloadURLRequest() (request *CreateCSPDownloadURLRequest) {
	request = &CreateCSPDownloadURLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("package", APIVersion, "CreateCSPDownloadURL")
	return
}

func NewCreateCSPDownloadURLResponse() (response *CreateCSPDownloadURLResponse) {
	response = &CreateCSPDownloadURLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建新版本之前通过此接口获得文件下载URL
func (c *Client) CreateCSPDownloadURL(request *CreateCSPDownloadURLRequest) (response *CreateCSPDownloadURLResponse, err error) {
	if request == nil {
		request = NewCreateCSPDownloadURLRequest()
	}
	response = NewCreateCSPDownloadURLResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCSPUploadURLRequest() (request *CreateCSPUploadURLRequest) {
	request = &CreateCSPUploadURLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("package", APIVersion, "CreateCSPUploadURL")
	return
}

func NewCreateCSPUploadURLResponse() (response *CreateCSPUploadURLResponse) {
	response = &CreateCSPUploadURLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建新版本之前通过此接口获得上传URL
func (c *Client) CreateCSPUploadURL(request *CreateCSPUploadURLRequest) (response *CreateCSPUploadURLResponse, err error) {
	if request == nil {
		request = NewCreateCSPUploadURLRequest()
	}
	response = NewCreateCSPUploadURLResponse()
	err = c.Send(request, response)
	return
}
