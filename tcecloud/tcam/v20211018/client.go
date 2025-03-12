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

package v20211018

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2021-10-18"

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

func NewUpdatePresetStrategyStateRequest() (request *UpdatePresetStrategyStateRequest) {
	request = &UpdatePresetStrategyStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcam", APIVersion, "UpdatePresetStrategyState")
	return
}

func NewUpdatePresetStrategyStateResponse() (response *UpdatePresetStrategyStateResponse) {
	response = &UpdatePresetStrategyStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改预设策略状态
func (c *Client) UpdatePresetStrategyState(request *UpdatePresetStrategyStateRequest) (response *UpdatePresetStrategyStateResponse, err error) {
	if request == nil {
		request = NewUpdatePresetStrategyStateRequest()
	}
	response = NewUpdatePresetStrategyStateResponse()
	err = c.Send(request, response)
	return
}

func NewGetPresetStrategyListRequest() (request *GetPresetStrategyListRequest) {
	request = &GetPresetStrategyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcam", APIVersion, "GetPresetStrategyList")
	return
}

func NewGetPresetStrategyListResponse() (response *GetPresetStrategyListResponse) {
	response = &GetPresetStrategyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取预设策略列表
func (c *Client) GetPresetStrategyList(request *GetPresetStrategyListRequest) (response *GetPresetStrategyListResponse, err error) {
	if request == nil {
		request = NewGetPresetStrategyListRequest()
	}
	response = NewGetPresetStrategyListResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateServiceRoleInfoRequest() (request *UpdateServiceRoleInfoRequest) {
	request = &UpdateServiceRoleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcam", APIVersion, "UpdateServiceRoleInfo")
	return
}

func NewUpdateServiceRoleInfoResponse() (response *UpdateServiceRoleInfoResponse) {
	response = &UpdateServiceRoleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改/创建/删除预设服务角色
func (c *Client) UpdateServiceRoleInfo(request *UpdateServiceRoleInfoRequest) (response *UpdateServiceRoleInfoResponse, err error) {
	if request == nil {
		request = NewUpdateServiceRoleInfoRequest()
	}
	response = NewUpdateServiceRoleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateServiceRequest() (request *UpdateServiceRequest) {
	request = &UpdateServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcam", APIVersion, "UpdateService")
	return
}

func NewUpdateServiceResponse() (response *UpdateServiceResponse) {
	response = &UpdateServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改/创建/删除预设服务
func (c *Client) UpdateService(request *UpdateServiceRequest) (response *UpdateServiceResponse, err error) {
	if request == nil {
		request = NewUpdateServiceRequest()
	}
	response = NewUpdateServiceResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateServicePermRequest() (request *UpdateServicePermRequest) {
	request = &UpdateServicePermRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcam", APIVersion, "UpdateServicePerm")
	return
}

func NewUpdateServicePermResponse() (response *UpdateServicePermResponse) {
	response = &UpdateServicePermResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改/创建/删除预设服务
func (c *Client) UpdateServicePerm(request *UpdateServicePermRequest) (response *UpdateServicePermResponse, err error) {
	if request == nil {
		request = NewUpdateServicePermRequest()
	}
	response = NewUpdateServicePermResponse()
	err = c.Send(request, response)
	return
}

func NewGetServiceListRequest() (request *GetServiceListRequest) {
	request = &GetServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcam", APIVersion, "GetServiceList")
	return
}

func NewGetServiceListResponse() (response *GetServiceListResponse) {
	response = &GetServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetServiceList
func (c *Client) GetServiceList(request *GetServiceListRequest) (response *GetServiceListResponse, err error) {
	if request == nil {
		request = NewGetServiceListRequest()
	}
	response = NewGetServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewListPoliciesRequest() (request *ListPoliciesRequest) {
	request = &ListPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcam", APIVersion, "ListPolicies")
	return
}

func NewListPoliciesResponse() (response *ListPoliciesResponse) {
	response = &ListPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询策略列表
func (c *Client) ListPolicies(request *ListPoliciesRequest) (response *ListPoliciesResponse, err error) {
	if request == nil {
		request = NewListPoliciesRequest()
	}
	response = NewListPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePresetStrategyRequest() (request *UpdatePresetStrategyRequest) {
	request = &UpdatePresetStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcam", APIVersion, "UpdatePresetStrategy")
	return
}

func NewUpdatePresetStrategyResponse() (response *UpdatePresetStrategyResponse) {
	response = &UpdatePresetStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改/创建/删除预设策略
func (c *Client) UpdatePresetStrategy(request *UpdatePresetStrategyRequest) (response *UpdatePresetStrategyResponse, err error) {
	if request == nil {
		request = NewUpdatePresetStrategyRequest()
	}
	response = NewUpdatePresetStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewGetServiceRoleInfoListRequest() (request *GetServiceRoleInfoListRequest) {
	request = &GetServiceRoleInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcam", APIVersion, "GetServiceRoleInfoList")
	return
}

func NewGetServiceRoleInfoListResponse() (response *GetServiceRoleInfoListResponse) {
	response = &GetServiceRoleInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取服务角色列表
func (c *Client) GetServiceRoleInfoList(request *GetServiceRoleInfoListRequest) (response *GetServiceRoleInfoListResponse, err error) {
	if request == nil {
		request = NewGetServiceRoleInfoListRequest()
	}
	response = NewGetServiceRoleInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewGetStrategyDetailRequest() (request *GetStrategyDetailRequest) {
	request = &GetStrategyDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcam", APIVersion, "GetStrategyDetail")
	return
}

func NewGetStrategyDetailResponse() (response *GetStrategyDetailResponse) {
	response = &GetStrategyDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取策略详情
func (c *Client) GetStrategyDetail(request *GetStrategyDetailRequest) (response *GetStrategyDetailResponse, err error) {
	if request == nil {
		request = NewGetStrategyDetailRequest()
	}
	response = NewGetStrategyDetailResponse()
	err = c.Send(request, response)
	return
}

func NewGetServicePermListRequest() (request *GetServicePermListRequest) {
	request = &GetServicePermListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tcam", APIVersion, "GetServicePermList")
	return
}

func NewGetServicePermListResponse() (response *GetServicePermListResponse) {
	response = &GetServicePermListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取接口列表
func (c *Client) GetServicePermList(request *GetServicePermListRequest) (response *GetServicePermListResponse, err error) {
	if request == nil {
		request = NewGetServicePermListRequest()
	}
	response = NewGetServicePermListResponse()
	err = c.Send(request, response)
	return
}
