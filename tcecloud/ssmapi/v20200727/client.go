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

package v20200727

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-07-27"

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

func NewDeletePlanRequest() (request *DeletePlanRequest) {
	request = &DeletePlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ssmapi", APIVersion, "DeletePlan")
	return
}

func NewDeletePlanResponse() (response *DeletePlanResponse) {
	response = &DeletePlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除指定的自定义plan
func (c *Client) DeletePlan(request *DeletePlanRequest) (response *DeletePlanResponse, err error) {
	if request == nil {
		request = NewDeletePlanRequest()
	}
	response = NewDeletePlanResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBindingRequest() (request *DeleteBindingRequest) {
	request = &DeleteBindingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ssmapi", APIVersion, "DeleteBinding")
	return
}

func NewDeleteBindingResponse() (response *DeleteBindingResponse) {
	response = &DeleteBindingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除绑定
func (c *Client) DeleteBinding(request *DeleteBindingRequest) (response *DeleteBindingResponse, err error) {
	if request == nil {
		request = NewDeleteBindingRequest()
	}
	response = NewDeleteBindingResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePlanRequest() (request *UpdatePlanRequest) {
	request = &UpdatePlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ssmapi", APIVersion, "UpdatePlan")
	return
}

func NewUpdatePlanResponse() (response *UpdatePlanResponse) {
	response = &UpdatePlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新自定义规格
func (c *Client) UpdatePlan(request *UpdatePlanRequest) (response *UpdatePlanResponse, err error) {
	if request == nil {
		request = NewUpdatePlanRequest()
	}
	response = NewUpdatePlanResponse()
	err = c.Send(request, response)
	return
}

func NewGetInstanceRequest() (request *GetInstanceRequest) {
	request = &GetInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ssmapi", APIVersion, "GetInstance")
	return
}

func NewGetInstanceResponse() (response *GetInstanceResponse) {
	response = &GetInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取实例详情
func (c *Client) GetInstance(request *GetInstanceRequest) (response *GetInstanceResponse, err error) {
	if request == nil {
		request = NewGetInstanceRequest()
	}
	response = NewGetInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBackupRecordRequest() (request *DeleteBackupRecordRequest) {
	request = &DeleteBackupRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ssmapi", APIVersion, "DeleteBackupRecord")
	return
}

func NewDeleteBackupRecordResponse() (response *DeleteBackupRecordResponse) {
	response = &DeleteBackupRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除备份记录
func (c *Client) DeleteBackupRecord(request *DeleteBackupRecordRequest) (response *DeleteBackupRecordResponse, err error) {
	if request == nil {
		request = NewDeleteBackupRecordRequest()
	}
	response = NewDeleteBackupRecordResponse()
	err = c.Send(request, response)
	return
}

func NewListBindingsRequest() (request *ListBindingsRequest) {
	request = &ListBindingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ssmapi", APIVersion, "ListBindings")
	return
}

func NewListBindingsResponse() (response *ListBindingsResponse) {
	response = &ListBindingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出所有绑定信息
func (c *Client) ListBindings(request *ListBindingsRequest) (response *ListBindingsResponse, err error) {
	if request == nil {
		request = NewListBindingsRequest()
	}
	response = NewListBindingsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
	request = &CreateInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ssmapi", APIVersion, "CreateInstance")
	return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
	response = &CreateInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建实例
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
	if request == nil {
		request = NewCreateInstanceRequest()
	}
	response = NewCreateInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
	request = &DeleteInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ssmapi", APIVersion, "DeleteInstance")
	return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
	response = &DeleteInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除实例
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
	if request == nil {
		request = NewDeleteInstanceRequest()
	}
	response = NewDeleteInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewListBackupRecordsRequest() (request *ListBackupRecordsRequest) {
	request = &ListBackupRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ssmapi", APIVersion, "ListBackupRecords")
	return
}

func NewListBackupRecordsResponse() (response *ListBackupRecordsResponse) {
	response = &ListBackupRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出所有备份条目
func (c *Client) ListBackupRecords(request *ListBackupRecordsRequest) (response *ListBackupRecordsResponse, err error) {
	if request == nil {
		request = NewListBackupRecordsRequest()
	}
	response = NewListBackupRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBindingRequest() (request *CreateBindingRequest) {
	request = &CreateBindingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ssmapi", APIVersion, "CreateBinding")
	return
}

func NewCreateBindingResponse() (response *CreateBindingResponse) {
	response = &CreateBindingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建绑定
func (c *Client) CreateBinding(request *CreateBindingRequest) (response *CreateBindingResponse, err error) {
	if request == nil {
		request = NewCreateBindingRequest()
	}
	response = NewCreateBindingResponse()
	err = c.Send(request, response)
	return
}

func NewListInstancesRequest() (request *ListInstancesRequest) {
	request = &ListInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ssmapi", APIVersion, "ListInstances")
	return
}

func NewListInstancesResponse() (response *ListInstancesResponse) {
	response = &ListInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出服务下所有相关实例
func (c *Client) ListInstances(request *ListInstancesRequest) (response *ListInstancesResponse, err error) {
	if request == nil {
		request = NewListInstancesRequest()
	}
	response = NewListInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateBackupRequest() (request *UpdateBackupRequest) {
	request = &UpdateBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ssmapi", APIVersion, "UpdateBackup")
	return
}

func NewUpdateBackupResponse() (response *UpdateBackupResponse) {
	response = &UpdateBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新备份计划
func (c *Client) UpdateBackup(request *UpdateBackupRequest) (response *UpdateBackupResponse, err error) {
	if request == nil {
		request = NewUpdateBackupRequest()
	}
	response = NewUpdateBackupResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePlanRequest() (request *CreatePlanRequest) {
	request = &CreatePlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ssmapi", APIVersion, "CreatePlan")
	return
}

func NewCreatePlanResponse() (response *CreatePlanResponse) {
	response = &CreatePlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建特定服务的自定义规格
func (c *Client) CreatePlan(request *CreatePlanRequest) (response *CreatePlanResponse, err error) {
	if request == nil {
		request = NewCreatePlanRequest()
	}
	response = NewCreatePlanResponse()
	err = c.Send(request, response)
	return
}

func NewListAllServicesRequest() (request *ListAllServicesRequest) {
	request = &ListAllServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ssmapi", APIVersion, "ListAllServices")
	return
}

func NewListAllServicesResponse() (response *ListAllServicesResponse) {
	response = &ListAllServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出所有支撑服务信息
func (c *Client) ListAllServices(request *ListAllServicesRequest) (response *ListAllServicesResponse, err error) {
	if request == nil {
		request = NewListAllServicesRequest()
	}
	response = NewListAllServicesResponse()
	err = c.Send(request, response)
	return
}

func NewListPlansRequest() (request *ListPlansRequest) {
	request = &ListPlansRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ssmapi", APIVersion, "ListPlans")
	return
}

func NewListPlansResponse() (response *ListPlansResponse) {
	response = &ListPlansResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出所有服务规格
func (c *Client) ListPlans(request *ListPlansRequest) (response *ListPlansResponse, err error) {
	if request == nil {
		request = NewListPlansRequest()
	}
	response = NewListPlansResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBackupRequest() (request *CreateBackupRequest) {
	request = &CreateBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ssmapi", APIVersion, "CreateBackup")
	return
}

func NewCreateBackupResponse() (response *CreateBackupResponse) {
	response = &CreateBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建实例备份
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
	if request == nil {
		request = NewCreateBackupRequest()
	}
	response = NewCreateBackupResponse()
	err = c.Send(request, response)
	return
}
