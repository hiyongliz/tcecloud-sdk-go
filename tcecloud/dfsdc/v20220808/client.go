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

package v20220808

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2022-08-08"

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

func NewDescribeAccountingResultsOfLastPeriodRequest() (request *DescribeAccountingResultsOfLastPeriodRequest) {
	request = &DescribeAccountingResultsOfLastPeriodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dfsdc", APIVersion, "DescribeAccountingResultsOfLastPeriod")
	return
}

func NewDescribeAccountingResultsOfLastPeriodResponse() (response *DescribeAccountingResultsOfLastPeriodResponse) {
	response = &DescribeAccountingResultsOfLastPeriodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeAccountingResultsOfLastPeriod) 用于查询综合对账任务上一对账周期的结果，只返回对账结果有异常的任务。
func (c *Client) DescribeAccountingResultsOfLastPeriod(request *DescribeAccountingResultsOfLastPeriodRequest) (response *DescribeAccountingResultsOfLastPeriodResponse, err error) {
	if request == nil {
		request = NewDescribeAccountingResultsOfLastPeriodRequest()
	}
	response = NewDescribeAccountingResultsOfLastPeriodResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountingBusinessesRequest() (request *DescribeAccountingBusinessesRequest) {
	request = &DescribeAccountingBusinessesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dfsdc", APIVersion, "DescribeAccountingBusinesses")
	return
}

func NewDescribeAccountingBusinessesResponse() (response *DescribeAccountingBusinessesResponse) {
	response = &DescribeAccountingBusinessesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeAccountingBusinesses) 用于查询支持综合对账的业务类型，当前只支持专线（DC）业务。
func (c *Client) DescribeAccountingBusinesses(request *DescribeAccountingBusinessesRequest) (response *DescribeAccountingBusinessesResponse, err error) {
	if request == nil {
		request = NewDescribeAccountingBusinessesRequest()
	}
	response = NewDescribeAccountingBusinessesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountingResultsByTaskRequest() (request *DescribeAccountingResultsByTaskRequest) {
	request = &DescribeAccountingResultsByTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dfsdc", APIVersion, "DescribeAccountingResultsByTask")
	return
}

func NewDescribeAccountingResultsByTaskResponse() (response *DescribeAccountingResultsByTaskResponse) {
	response = &DescribeAccountingResultsByTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeAccountingResultsByTask) 用于根据任务信息查询综合对账任务结果。
func (c *Client) DescribeAccountingResultsByTask(request *DescribeAccountingResultsByTaskRequest) (response *DescribeAccountingResultsByTaskResponse, err error) {
	if request == nil {
		request = NewDescribeAccountingResultsByTaskRequest()
	}
	response = NewDescribeAccountingResultsByTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAccountingTasksRequest() (request *CreateAccountingTasksRequest) {
	request = &CreateAccountingTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dfsdc", APIVersion, "CreateAccountingTasks")
	return
}

func NewCreateAccountingTasksResponse() (response *CreateAccountingTasksResponse) {
	response = &CreateAccountingTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateAccountingTasks）创建专线对账任务。
func (c *Client) CreateAccountingTasks(request *CreateAccountingTasksRequest) (response *CreateAccountingTasksResponse, err error) {
	if request == nil {
		request = NewCreateAccountingTasksRequest()
	}
	response = NewCreateAccountingTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountingResultsRequest() (request *DescribeAccountingResultsRequest) {
	request = &DescribeAccountingResultsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dfsdc", APIVersion, "DescribeAccountingResults")
	return
}

func NewDescribeAccountingResultsResponse() (response *DescribeAccountingResultsResponse) {
	response = &DescribeAccountingResultsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeAccountingResults) 用于查询综合对账某一时间段内任务结果，只返回对账结果有异常的任务。
func (c *Client) DescribeAccountingResults(request *DescribeAccountingResultsRequest) (response *DescribeAccountingResultsResponse, err error) {
	if request == nil {
		request = NewDescribeAccountingResultsRequest()
	}
	response = NewDescribeAccountingResultsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAccountingTasksRequest() (request *UpdateAccountingTasksRequest) {
	request = &UpdateAccountingTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dfsdc", APIVersion, "UpdateAccountingTasks")
	return
}

func NewUpdateAccountingTasksResponse() (response *UpdateAccountingTasksResponse) {
	response = &UpdateAccountingTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateAccountingTasks）更新综合对账任务，主要是用来触发修复。
func (c *Client) UpdateAccountingTasks(request *UpdateAccountingTasksRequest) (response *UpdateAccountingTasksResponse, err error) {
	if request == nil {
		request = NewUpdateAccountingTasksRequest()
	}
	response = NewUpdateAccountingTasksResponse()
	err = c.Send(request, response)
	return
}
