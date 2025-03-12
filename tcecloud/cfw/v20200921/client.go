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

package v20200921

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-09-21"

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

func NewDescribeConfigRequest() (request *DescribeConfigRequest) {
	request = &DescribeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeConfig")
	return
}

func NewDescribeConfigResponse() (response *DescribeConfigResponse) {
	response = &DescribeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取配置信息
func (c *Client) DescribeConfig(request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
	if request == nil {
		request = NewDescribeConfigRequest()
	}
	response = NewDescribeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewUpdataConfigRequest() (request *UpdataConfigRequest) {
	request = &UpdataConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "UpdataConfig")
	return
}

func NewUpdataConfigResponse() (response *UpdataConfigResponse) {
	response = &UpdataConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品设置-修改cvm log 配置
func (c *Client) UpdataConfig(request *UpdataConfigRequest) (response *UpdataConfigResponse, err error) {
	if request == nil {
		request = NewUpdataConfigRequest()
	}
	response = NewUpdataConfigResponse()
	err = c.Send(request, response)
	return
}

func NewAddKeepIdRequest() (request *AddKeepIdRequest) {
	request = &AddKeepIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "AddKeepId")
	return
}

func NewAddKeepIdResponse() (response *AddKeepIdResponse) {
	response = &AddKeepIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增VPC保留网段
func (c *Client) AddKeepId(request *AddKeepIdRequest) (response *AddKeepIdResponse, err error) {
	if request == nil {
		request = NewAddKeepIdRequest()
	}
	response = NewAddKeepIdResponse()
	err = c.Send(request, response)
	return
}

func NewAddUrlRequest() (request *AddUrlRequest) {
	request = &AddUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "AddUrl")
	return
}

func NewAddUrlResponse() (response *AddUrlResponse) {
	response = &AddUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增系统升级URL配置
func (c *Client) AddUrl(request *AddUrlRequest) (response *AddUrlResponse, err error) {
	if request == nil {
		request = NewAddUrlRequest()
	}
	response = NewAddUrlResponse()
	err = c.Send(request, response)
	return
}

func NewResetCvmConfigRequest() (request *ResetCvmConfigRequest) {
	request = &ResetCvmConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ResetCvmConfig")
	return
}

func NewResetCvmConfigResponse() (response *ResetCvmConfigResponse) {
	response = &ResetCvmConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回滚CVM模块化配置
func (c *Client) ResetCvmConfig(request *ResetCvmConfigRequest) (response *ResetCvmConfigResponse, err error) {
	if request == nil {
		request = NewResetCvmConfigRequest()
	}
	response = NewResetCvmConfigResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateModuleConfigRequest() (request *UpdateModuleConfigRequest) {
	request = &UpdateModuleConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "UpdateModuleConfig")
	return
}

func NewUpdateModuleConfigResponse() (response *UpdateModuleConfigResponse) {
	response = &UpdateModuleConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改引擎CVM时初始化模块配置
func (c *Client) UpdateModuleConfig(request *UpdateModuleConfigRequest) (response *UpdateModuleConfigResponse, err error) {
	if request == nil {
		request = NewUpdateModuleConfigRequest()
	}
	response = NewUpdateModuleConfigResponse()
	err = c.Send(request, response)
	return
}

func NewAddModuleConfigRequest() (request *AddModuleConfigRequest) {
	request = &AddModuleConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "AddModuleConfig")
	return
}

func NewAddModuleConfigResponse() (response *AddModuleConfigResponse) {
	response = &AddModuleConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增引擎CVM时初始化模块配置
func (c *Client) AddModuleConfig(request *AddModuleConfigRequest) (response *AddModuleConfigResponse, err error) {
	if request == nil {
		request = NewAddModuleConfigRequest()
	}
	response = NewAddModuleConfigResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateKeepIdRequest() (request *UpdateKeepIdRequest) {
	request = &UpdateKeepIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "UpdateKeepId")
	return
}

func NewUpdateKeepIdResponse() (response *UpdateKeepIdResponse) {
	response = &UpdateKeepIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置功能修改网段
func (c *Client) UpdateKeepId(request *UpdateKeepIdRequest) (response *UpdateKeepIdResponse, err error) {
	if request == nil {
		request = NewUpdateKeepIdRequest()
	}
	response = NewUpdateKeepIdResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRollbackRequest() (request *ModifyRollbackRequest) {
	request = &ModifyRollbackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyRollback")
	return
}

func NewModifyRollbackResponse() (response *ModifyRollbackResponse) {
	response = &ModifyRollbackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改产品模块配置回滚
func (c *Client) ModifyRollback(request *ModifyRollbackRequest) (response *ModifyRollbackResponse, err error) {
	if request == nil {
		request = NewModifyRollbackRequest()
	}
	response = NewModifyRollbackResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCvmConfigRequest() (request *DeleteCvmConfigRequest) {
	request = &DeleteCvmConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DeleteCvmConfig")
	return
}

func NewDeleteCvmConfigResponse() (response *DeleteCvmConfigResponse) {
	response = &DeleteCvmConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除CVM模块化配置
func (c *Client) DeleteCvmConfig(request *DeleteCvmConfigRequest) (response *DeleteCvmConfigResponse, err error) {
	if request == nil {
		request = NewDeleteCvmConfigRequest()
	}
	response = NewDeleteCvmConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiDispatchRequest() (request *DescribeApiDispatchRequest) {
	request = &DescribeApiDispatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "DescribeApiDispatch")
	return
}

func NewDescribeApiDispatchResponse() (response *DescribeApiDispatchResponse) {
	response = &DescribeApiDispatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口请求中转
func (c *Client) DescribeApiDispatch(request *DescribeApiDispatchRequest) (response *DescribeApiDispatchResponse, err error) {
	if request == nil {
		request = NewDescribeApiDispatchRequest()
	}
	response = NewDescribeApiDispatchResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReleaseVersionRequest() (request *ModifyReleaseVersionRequest) {
	request = &ModifyReleaseVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "ModifyReleaseVersion")
	return
}

func NewModifyReleaseVersionResponse() (response *ModifyReleaseVersionResponse) {
	response = &ModifyReleaseVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改产品模块配置
func (c *Client) ModifyReleaseVersion(request *ModifyReleaseVersionRequest) (response *ModifyReleaseVersionResponse, err error) {
	if request == nil {
		request = NewModifyReleaseVersionRequest()
	}
	response = NewModifyReleaseVersionResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateUrlRequest() (request *UpdateUrlRequest) {
	request = &UpdateUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cfw", APIVersion, "UpdateUrl")
	return
}

func NewUpdateUrlResponse() (response *UpdateUrlResponse) {
	response = &UpdateUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改系统升级URL配置
func (c *Client) UpdateUrl(request *UpdateUrlRequest) (response *UpdateUrlResponse, err error) {
	if request == nil {
		request = NewUpdateUrlRequest()
	}
	response = NewUpdateUrlResponse()
	err = c.Send(request, response)
	return
}
