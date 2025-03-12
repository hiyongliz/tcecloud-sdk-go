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

package v20200501

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-05-01"

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

func NewCreatePeriodicTaskRequest() (request *CreatePeriodicTaskRequest) {
	request = &CreatePeriodicTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "CreatePeriodicTask")
	return
}

func NewCreatePeriodicTaskResponse() (response *CreatePeriodicTaskResponse) {
	response = &CreatePeriodicTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建定时任务。
func (c *Client) CreatePeriodicTask(request *CreatePeriodicTaskRequest) (response *CreatePeriodicTaskResponse, err error) {
	if request == nil {
		request = NewCreatePeriodicTaskRequest()
	}
	response = NewCreatePeriodicTaskResponse()
	err = c.Send(request, response)
	return
}

func NewQueryWorkflowInstanceRequest() (request *QueryWorkflowInstanceRequest) {
	request = &QueryWorkflowInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "QueryWorkflowInstance")
	return
}

func NewQueryWorkflowInstanceResponse() (response *QueryWorkflowInstanceResponse) {
	response = &QueryWorkflowInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据通用过滤条件查询单个流程实例，或者流程实例列表。只包含流程基本信息，不包含流程下的节点的信息。
//
// 如根据InstanceUUID查询，则Filter.Name为InstanceUUID，而 Filter.Value，则为此目标InstanceUUID数组的序列化JSON字符串。
func (c *Client) QueryWorkflowInstance(request *QueryWorkflowInstanceRequest) (response *QueryWorkflowInstanceResponse, err error) {
	if request == nil {
		request = NewQueryWorkflowInstanceRequest()
	}
	response = NewQueryWorkflowInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePeriodicTaskRequest() (request *UpdatePeriodicTaskRequest) {
	request = &UpdatePeriodicTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "UpdatePeriodicTask")
	return
}

func NewUpdatePeriodicTaskResponse() (response *UpdatePeriodicTaskResponse) {
	response = &UpdatePeriodicTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新定时任务。
func (c *Client) UpdatePeriodicTask(request *UpdatePeriodicTaskRequest) (response *UpdatePeriodicTaskResponse, err error) {
	if request == nil {
		request = NewUpdatePeriodicTaskRequest()
	}
	response = NewUpdatePeriodicTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateWorkflowInstanceRequest() (request *CreateWorkflowInstanceRequest) {
	request = &CreateWorkflowInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "CreateWorkflowInstance")
	return
}

func NewCreateWorkflowInstanceResponse() (response *CreateWorkflowInstanceResponse) {
	response = &CreateWorkflowInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动指定的流程实例。即执行作业工具或作业编排。
func (c *Client) CreateWorkflowInstance(request *CreateWorkflowInstanceRequest) (response *CreateWorkflowInstanceResponse, err error) {
	if request == nil {
		request = NewCreateWorkflowInstanceRequest()
	}
	response = NewCreateWorkflowInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePeriodicTaskRequest() (request *DeletePeriodicTaskRequest) {
	request = &DeletePeriodicTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "DeletePeriodicTask")
	return
}

func NewDeletePeriodicTaskResponse() (response *DeletePeriodicTaskResponse) {
	response = &DeletePeriodicTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除定时任务。
func (c *Client) DeletePeriodicTask(request *DeletePeriodicTaskRequest) (response *DeletePeriodicTaskResponse, err error) {
	if request == nil {
		request = NewDeletePeriodicTaskRequest()
	}
	response = NewDeletePeriodicTaskResponse()
	err = c.Send(request, response)
	return
}

func NewGetAllPluginWithCategoryRequest() (request *GetAllPluginWithCategoryRequest) {
	request = &GetAllPluginWithCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "GetAllPluginWithCategory")
	return
}

func NewGetAllPluginWithCategoryResponse() (response *GetAllPluginWithCategoryResponse) {
	response = &GetAllPluginWithCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取插件分类和列表
func (c *Client) GetAllPluginWithCategory(request *GetAllPluginWithCategoryRequest) (response *GetAllPluginWithCategoryResponse, err error) {
	if request == nil {
		request = NewGetAllPluginWithCategoryRequest()
	}
	response = NewGetAllPluginWithCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewCheckWorkflowTemplateRequest() (request *CheckWorkflowTemplateRequest) {
	request = &CheckWorkflowTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "CheckWorkflowTemplate")
	return
}

func NewCheckWorkflowTemplateResponse() (response *CheckWorkflowTemplateResponse) {
	response = &CheckWorkflowTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过给定作业工具或者编排uuid校验其是否存在
func (c *Client) CheckWorkflowTemplate(request *CheckWorkflowTemplateRequest) (response *CheckWorkflowTemplateResponse, err error) {
	if request == nil {
		request = NewCheckWorkflowTemplateRequest()
	}
	response = NewCheckWorkflowTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewQueryBizTreeRequest() (request *QueryBizTreeRequest) {
	request = &QueryBizTreeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "QueryBizTree")
	return
}

func NewQueryBizTreeResponse() (response *QueryBizTreeResponse) {
	response = &QueryBizTreeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定业务的完整的业务模块树。
func (c *Client) QueryBizTree(request *QueryBizTreeRequest) (response *QueryBizTreeResponse, err error) {
	if request == nil {
		request = NewQueryBizTreeRequest()
	}
	response = NewQueryBizTreeResponse()
	err = c.Send(request, response)
	return
}

func NewQueryWorkflowInstanceDetailRequest() (request *QueryWorkflowInstanceDetailRequest) {
	request = &QueryWorkflowInstanceDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "QueryWorkflowInstanceDetail")
	return
}

func NewQueryWorkflowInstanceDetailResponse() (response *QueryWorkflowInstanceDetailResponse) {
	response = &QueryWorkflowInstanceDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于查询作业工具或者作业编排的执行记录，且不包含详细的节点执行信息。只包含该流程的信息以及其包含的节点基本信息。
func (c *Client) QueryWorkflowInstanceDetail(request *QueryWorkflowInstanceDetailRequest) (response *QueryWorkflowInstanceDetailResponse, err error) {
	if request == nil {
		request = NewQueryWorkflowInstanceDetailRequest()
	}
	response = NewQueryWorkflowInstanceDetailResponse()
	err = c.Send(request, response)
	return
}

func NewQueryWorkflowNodeDetailRequest() (request *QueryWorkflowNodeDetailRequest) {
	request = &QueryWorkflowNodeDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "QueryWorkflowNodeDetail")
	return
}

func NewQueryWorkflowNodeDetailResponse() (response *QueryWorkflowNodeDetailResponse) {
	response = &QueryWorkflowNodeDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于查询作业工具或者作业编排的执行记录，且只包含指定节点指定执行次数的执行信息，以及节点基本信息。不包含该流程的信息。
func (c *Client) QueryWorkflowNodeDetail(request *QueryWorkflowNodeDetailRequest) (response *QueryWorkflowNodeDetailResponse, err error) {
	if request == nil {
		request = NewQueryWorkflowNodeDetailRequest()
	}
	response = NewQueryWorkflowNodeDetailResponse()
	err = c.Send(request, response)
	return
}

func NewQueryBizListRequest() (request *QueryBizListRequest) {
	request = &QueryBizListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "QueryBizList")
	return
}

func NewQueryBizListResponse() (response *QueryBizListResponse) {
	response = &QueryBizListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询业务列表，不包含完整的业务树信息，只包含业务本身信息的列表
func (c *Client) QueryBizList(request *QueryBizListRequest) (response *QueryBizListResponse, err error) {
	if request == nil {
		request = NewQueryBizListRequest()
	}
	response = NewQueryBizListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateWorkflowTemplatesRequest() (request *CreateWorkflowTemplatesRequest) {
	request = &CreateWorkflowTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "CreateWorkflowTemplates")
	return
}

func NewCreateWorkflowTemplatesResponse() (response *CreateWorkflowTemplatesResponse) {
	response = &CreateWorkflowTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量创建作业工具或者作业编排
func (c *Client) CreateWorkflowTemplates(request *CreateWorkflowTemplatesRequest) (response *CreateWorkflowTemplatesResponse, err error) {
	if request == nil {
		request = NewCreateWorkflowTemplatesRequest()
	}
	response = NewCreateWorkflowTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewStopPeriodicTaskRequest() (request *StopPeriodicTaskRequest) {
	request = &StopPeriodicTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "StopPeriodicTask")
	return
}

func NewStopPeriodicTaskResponse() (response *StopPeriodicTaskResponse) {
	response = &StopPeriodicTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据通用过滤器停止指定定时任务。
func (c *Client) StopPeriodicTask(request *StopPeriodicTaskRequest) (response *StopPeriodicTaskResponse, err error) {
	if request == nil {
		request = NewStopPeriodicTaskRequest()
	}
	response = NewStopPeriodicTaskResponse()
	err = c.Send(request, response)
	return
}

func NewQueryPluginOutputSpecRequest() (request *QueryPluginOutputSpecRequest) {
	request = &QueryPluginOutputSpecRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "QueryPluginOutputSpec")
	return
}

func NewQueryPluginOutputSpecResponse() (response *QueryPluginOutputSpecResponse) {
	response = &QueryPluginOutputSpecResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定插件的输出参数规范列表
func (c *Client) QueryPluginOutputSpec(request *QueryPluginOutputSpecRequest) (response *QueryPluginOutputSpecResponse, err error) {
	if request == nil {
		request = NewQueryPluginOutputSpecRequest()
	}
	response = NewQueryPluginOutputSpecResponse()
	err = c.Send(request, response)
	return
}

func NewStartPeriodicTaskRequest() (request *StartPeriodicTaskRequest) {
	request = &StartPeriodicTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "StartPeriodicTask")
	return
}

func NewStartPeriodicTaskResponse() (response *StartPeriodicTaskResponse) {
	response = &StartPeriodicTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动定时任务。
func (c *Client) StartPeriodicTask(request *StartPeriodicTaskRequest) (response *StartPeriodicTaskResponse, err error) {
	if request == nil {
		request = NewStartPeriodicTaskRequest()
	}
	response = NewStartPeriodicTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMultiTemplateRequest() (request *CreateMultiTemplateRequest) {
	request = &CreateMultiTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "CreateMultiTemplate")
	return
}

func NewCreateMultiTemplateResponse() (response *CreateMultiTemplateResponse) {
	response = &CreateMultiTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量创建作业工具或者作业编排，用于作业工具和作业编排的导入功能。
func (c *Client) CreateMultiTemplate(request *CreateMultiTemplateRequest) (response *CreateMultiTemplateResponse, err error) {
	if request == nil {
		request = NewCreateMultiTemplateRequest()
	}
	response = NewCreateMultiTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewQueryAgentListRequest() (request *QueryAgentListRequest) {
	request = &QueryAgentListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "QueryAgentList")
	return
}

func NewQueryAgentListResponse() (response *QueryAgentListResponse) {
	response = &QueryAgentListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定业务/集群/模块下的主机列表，同时可通过主机名及ip或sn进行模糊搜索
func (c *Client) QueryAgentList(request *QueryAgentListRequest) (response *QueryAgentListResponse, err error) {
	if request == nil {
		request = NewQueryAgentListRequest()
	}
	response = NewQueryAgentListResponse()
	err = c.Send(request, response)
	return
}

func NewQueryWorkflowTemplateRequest() (request *QueryWorkflowTemplateRequest) {
	request = &QueryWorkflowTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "QueryWorkflowTemplate")
	return
}

func NewQueryWorkflowTemplateResponse() (response *QueryWorkflowTemplateResponse) {
	response = &QueryWorkflowTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据指定条件查询作业工具或模板列表
func (c *Client) QueryWorkflowTemplate(request *QueryWorkflowTemplateRequest) (response *QueryWorkflowTemplateResponse, err error) {
	if request == nil {
		request = NewQueryWorkflowTemplateRequest()
	}
	response = NewQueryWorkflowTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewRetryRunWorkflowNodeRequest() (request *RetryRunWorkflowNodeRequest) {
	request = &RetryRunWorkflowNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "RetryRunWorkflowNode")
	return
}

func NewRetryRunWorkflowNodeResponse() (response *RetryRunWorkflowNodeResponse) {
	response = &RetryRunWorkflowNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 当用户设置了允许手动重试失败节点时。
// 在活动节点中，如果插件调用失败，则节点会进入失败状态，可以通过这个接口重新运行节点。
func (c *Client) RetryRunWorkflowNode(request *RetryRunWorkflowNodeRequest) (response *RetryRunWorkflowNodeResponse, err error) {
	if request == nil {
		request = NewRetryRunWorkflowNodeRequest()
	}
	response = NewRetryRunWorkflowNodeResponse()
	err = c.Send(request, response)
	return
}

func NewSkipWorkflowNodeRequest() (request *SkipWorkflowNodeRequest) {
	request = &SkipWorkflowNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "SkipWorkflowNode")
	return
}

func NewSkipWorkflowNodeResponse() (response *SkipWorkflowNodeResponse) {
	response = &SkipWorkflowNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 当用户设置了允许手动跳过失败节点时，该接口用于跳过操作。
// 在活动节点中，如果插件调用失败，则节点会进入失败状态，可以通过这个接口直接跳过失败的节点，流程继续执行下一个节点。
func (c *Client) SkipWorkflowNode(request *SkipWorkflowNodeRequest) (response *SkipWorkflowNodeResponse, err error) {
	if request == nil {
		request = NewSkipWorkflowNodeRequest()
	}
	response = NewSkipWorkflowNodeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteWorkflowTemplateRequest() (request *DeleteWorkflowTemplateRequest) {
	request = &DeleteWorkflowTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "DeleteWorkflowTemplate")
	return
}

func NewDeleteWorkflowTemplateResponse() (response *DeleteWorkflowTemplateResponse) {
	response = &DeleteWorkflowTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据指定过滤条件来删除指定作业工具或者作业编排。需鉴权。
//
// 如根据TemplateUUID删除，则Filter.Name为TemplateUUID，而 Filter.Value，则为此目标TemplateUUID数组的序列化JSON字符串。
func (c *Client) DeleteWorkflowTemplate(request *DeleteWorkflowTemplateRequest) (response *DeleteWorkflowTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteWorkflowTemplateRequest()
	}
	response = NewDeleteWorkflowTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewQueryPluginInputSpecRequest() (request *QueryPluginInputSpecRequest) {
	request = &QueryPluginInputSpecRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "QueryPluginInputSpec")
	return
}

func NewQueryPluginInputSpecResponse() (response *QueryPluginInputSpecResponse) {
	response = &QueryPluginInputSpecResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定插件的输入参数规范列表
func (c *Client) QueryPluginInputSpec(request *QueryPluginInputSpecRequest) (response *QueryPluginInputSpecResponse, err error) {
	if request == nil {
		request = NewQueryPluginInputSpecRequest()
	}
	response = NewQueryPluginInputSpecResponse()
	err = c.Send(request, response)
	return
}

func NewStopWorkflowInstanceRequest() (request *StopWorkflowInstanceRequest) {
	request = &StopWorkflowInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "StopWorkflowInstance")
	return
}

func NewStopWorkflowInstanceResponse() (response *StopWorkflowInstanceResponse) {
	response = &StopWorkflowInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据通用过滤器停止流程实例的执行。
func (c *Client) StopWorkflowInstance(request *StopWorkflowInstanceRequest) (response *StopWorkflowInstanceResponse, err error) {
	if request == nil {
		request = NewStopWorkflowInstanceRequest()
	}
	response = NewStopWorkflowInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePeriodicTasksRequest() (request *CreatePeriodicTasksRequest) {
	request = &CreatePeriodicTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "CreatePeriodicTasks")
	return
}

func NewCreatePeriodicTasksResponse() (response *CreatePeriodicTasksResponse) {
	response = &CreatePeriodicTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量创定时任务用于导入操作。
func (c *Client) CreatePeriodicTasks(request *CreatePeriodicTasksRequest) (response *CreatePeriodicTasksResponse, err error) {
	if request == nil {
		request = NewCreatePeriodicTasksRequest()
	}
	response = NewCreatePeriodicTasksResponse()
	err = c.Send(request, response)
	return
}

func NewVerifyAgentListRequest() (request *VerifyAgentListRequest) {
	request = &VerifyAgentListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "VerifyAgentList")
	return
}

func NewVerifyAgentListResponse() (response *VerifyAgentListResponse) {
	response = &VerifyAgentListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验用户输入的主机IP列表。返回合法的host结构数组和非法的ip列表。
func (c *Client) VerifyAgentList(request *VerifyAgentListRequest) (response *VerifyAgentListResponse, err error) {
	if request == nil {
		request = NewVerifyAgentListRequest()
	}
	response = NewVerifyAgentListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateWorkflowTemplateRequest() (request *CreateWorkflowTemplateRequest) {
	request = &CreateWorkflowTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "CreateWorkflowTemplate")
	return
}

func NewCreateWorkflowTemplateResponse() (response *CreateWorkflowTemplateResponse) {
	response = &CreateWorkflowTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建作业工具或者作业编排的目标
func (c *Client) CreateWorkflowTemplate(request *CreateWorkflowTemplateRequest) (response *CreateWorkflowTemplateResponse, err error) {
	if request == nil {
		request = NewCreateWorkflowTemplateRequest()
	}
	response = NewCreateWorkflowTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewQueryPeriodicTaskRequest() (request *QueryPeriodicTaskRequest) {
	request = &QueryPeriodicTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "QueryPeriodicTask")
	return
}

func NewQueryPeriodicTaskResponse() (response *QueryPeriodicTaskResponse) {
	response = &QueryPeriodicTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询定时任务。
func (c *Client) QueryPeriodicTask(request *QueryPeriodicTaskRequest) (response *QueryPeriodicTaskResponse, err error) {
	if request == nil {
		request = NewQueryPeriodicTaskRequest()
	}
	response = NewQueryPeriodicTaskResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateWorkflowTemplateRequest() (request *UpdateWorkflowTemplateRequest) {
	request = &UpdateWorkflowTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("workflowtool", APIVersion, "UpdateWorkflowTemplate")
	return
}

func NewUpdateWorkflowTemplateResponse() (response *UpdateWorkflowTemplateResponse) {
	response = &UpdateWorkflowTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据模板UUID更新作业工具或者作业编排。需鉴权。
func (c *Client) UpdateWorkflowTemplate(request *UpdateWorkflowTemplateRequest) (response *UpdateWorkflowTemplateResponse, err error) {
	if request == nil {
		request = NewUpdateWorkflowTemplateRequest()
	}
	response = NewUpdateWorkflowTemplateResponse()
	err = c.Send(request, response)
	return
}
