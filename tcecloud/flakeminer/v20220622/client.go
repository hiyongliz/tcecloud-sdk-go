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

package v20220622

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2022-06-22"

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

func NewDeleteTaskRequest() (request *DeleteTaskRequest) {
	request = &DeleteTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("flakeminer", APIVersion, "DeleteTask")
	return
}

func NewDeleteTaskResponse() (response *DeleteTaskResponse) {
	response = &DeleteTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除采集任务
func (c *Client) DeleteTask(request *DeleteTaskRequest) (response *DeleteTaskResponse, err error) {
	if request == nil {
		request = NewDeleteTaskRequest()
	}
	response = NewDeleteTaskResponse()
	err = c.Send(request, response)
	return
}

func NewListStagesRequest() (request *ListStagesRequest) {
	request = &ListStagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("flakeminer", APIVersion, "ListStages")
	return
}

func NewListStagesResponse() (response *ListStagesResponse) {
	response = &ListStagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取stage列表
func (c *Client) ListStages(request *ListStagesRequest) (response *ListStagesResponse, err error) {
	if request == nil {
		request = NewListStagesRequest()
	}
	response = NewListStagesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskRunRequest() (request *DescribeTaskRunRequest) {
	request = &DescribeTaskRunRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("flakeminer", APIVersion, "DescribeTaskRun")
	return
}

func NewDescribeTaskRunResponse() (response *DescribeTaskRunResponse) {
	response = &DescribeTaskRunResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取采集任务实例详情
func (c *Client) DescribeTaskRun(request *DescribeTaskRunRequest) (response *DescribeTaskRunResponse, err error) {
	if request == nil {
		request = NewDescribeTaskRunRequest()
	}
	response = NewDescribeTaskRunResponse()
	err = c.Send(request, response)
	return
}

func NewListTaskRunsRequest() (request *ListTaskRunsRequest) {
	request = &ListTaskRunsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("flakeminer", APIVersion, "ListTaskRuns")
	return
}

func NewListTaskRunsResponse() (response *ListTaskRunsResponse) {
	response = &ListTaskRunsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取采集任务执行实例列表
func (c *Client) ListTaskRuns(request *ListTaskRunsRequest) (response *ListTaskRunsResponse, err error) {
	if request == nil {
		request = NewListTaskRunsRequest()
	}
	response = NewListTaskRunsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateTaskRequest() (request *UpdateTaskRequest) {
	request = &UpdateTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("flakeminer", APIVersion, "UpdateTask")
	return
}

func NewUpdateTaskResponse() (response *UpdateTaskResponse) {
	response = &UpdateTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新采集任务
func (c *Client) UpdateTask(request *UpdateTaskRequest) (response *UpdateTaskResponse, err error) {
	if request == nil {
		request = NewUpdateTaskRequest()
	}
	response = NewUpdateTaskResponse()
	err = c.Send(request, response)
	return
}

func NewExecuteTaskRequest() (request *ExecuteTaskRequest) {
	request = &ExecuteTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("flakeminer", APIVersion, "ExecuteTask")
	return
}

func NewExecuteTaskResponse() (response *ExecuteTaskResponse) {
	response = &ExecuteTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 执行采集任务(dryrun)
func (c *Client) ExecuteTask(request *ExecuteTaskRequest) (response *ExecuteTaskResponse, err error) {
	if request == nil {
		request = NewExecuteTaskRequest()
	}
	response = NewExecuteTaskResponse()
	err = c.Send(request, response)
	return
}

func NewListTasksRequest() (request *ListTasksRequest) {
	request = &ListTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("flakeminer", APIVersion, "ListTasks")
	return
}

func NewListTasksResponse() (response *ListTasksResponse) {
	response = &ListTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取采集任务列表
func (c *Client) ListTasks(request *ListTasksRequest) (response *ListTasksResponse, err error) {
	if request == nil {
		request = NewListTasksRequest()
	}
	response = NewListTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
	request = &DescribeTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("flakeminer", APIVersion, "DescribeTask")
	return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
	response = &DescribeTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取采集任务详情
func (c *Client) DescribeTask(request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
	if request == nil {
		request = NewDescribeTaskRequest()
	}
	response = NewDescribeTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTaskRequest() (request *CreateTaskRequest) {
	request = &CreateTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("flakeminer", APIVersion, "CreateTask")
	return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
	response = &CreateTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建采集任务
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
	if request == nil {
		request = NewCreateTaskRequest()
	}
	response = NewCreateTaskResponse()
	err = c.Send(request, response)
	return
}
