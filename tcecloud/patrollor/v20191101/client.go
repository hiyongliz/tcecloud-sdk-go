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

package v20191101

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-11-01"

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

func NewDescribeToolDetailRequest() (request *DescribeToolDetailRequest) {
	request = &DescribeToolDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("patrollor", APIVersion, "DescribeToolDetail")
	return
}

func NewDescribeToolDetailResponse() (response *DescribeToolDetailResponse) {
	response = &DescribeToolDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取作业工具中某作业的详情
func (c *Client) DescribeToolDetail(request *DescribeToolDetailRequest) (response *DescribeToolDetailResponse, err error) {
	if request == nil {
		request = NewDescribeToolDetailRequest()
	}
	response = NewDescribeToolDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePatrolCasesRequest() (request *DescribePatrolCasesRequest) {
	request = &DescribePatrolCasesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("patrollor", APIVersion, "DescribePatrolCases")
	return
}

func NewDescribePatrolCasesResponse() (response *DescribePatrolCasesResponse) {
	response = &DescribePatrolCasesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按条件列举巡检用例
func (c *Client) DescribePatrolCases(request *DescribePatrolCasesRequest) (response *DescribePatrolCasesResponse, err error) {
	if request == nil {
		request = NewDescribePatrolCasesRequest()
	}
	response = NewDescribePatrolCasesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowDetailRequest() (request *DescribeFlowDetailRequest) {
	request = &DescribeFlowDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("patrollor", APIVersion, "DescribeFlowDetail")
	return
}

func NewDescribeFlowDetailResponse() (response *DescribeFlowDetailResponse) {
	response = &DescribeFlowDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取作业编排详情
func (c *Client) DescribeFlowDetail(request *DescribeFlowDetailRequest) (response *DescribeFlowDetailResponse, err error) {
	if request == nil {
		request = NewDescribeFlowDetailRequest()
	}
	response = NewDescribeFlowDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeModuleHostsRequest() (request *DescribeModuleHostsRequest) {
	request = &DescribeModuleHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("patrollor", APIVersion, "DescribeModuleHosts")
	return
}

func NewDescribeModuleHostsResponse() (response *DescribeModuleHostsResponse) {
	response = &DescribeModuleHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 给定模块的 bk_tx_unique_code，返回 Host 列表
func (c *Client) DescribeModuleHosts(request *DescribeModuleHostsRequest) (response *DescribeModuleHostsResponse, err error) {
	if request == nil {
		request = NewDescribeModuleHostsRequest()
	}
	response = NewDescribeModuleHostsResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePatrolCaseRequest() (request *DeletePatrolCaseRequest) {
	request = &DeletePatrolCaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("patrollor", APIVersion, "DeletePatrolCase")
	return
}

func NewDeletePatrolCaseResponse() (response *DeletePatrolCaseResponse) {
	response = &DeletePatrolCaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除指定的巡检用例
func (c *Client) DeletePatrolCase(request *DeletePatrolCaseRequest) (response *DeletePatrolCaseResponse, err error) {
	if request == nil {
		request = NewDeletePatrolCaseRequest()
	}
	response = NewDeletePatrolCaseResponse()
	err = c.Send(request, response)
	return
}

func NewGetSuggestionsRequest() (request *GetSuggestionsRequest) {
	request = &GetSuggestionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("patrollor", APIVersion, "GetSuggestions")
	return
}

func NewGetSuggestionsResponse() (response *GetSuggestionsResponse) {
	response = &GetSuggestionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取各种输入推荐，如巡检用例名
func (c *Client) GetSuggestions(request *GetSuggestionsRequest) (response *GetSuggestionsResponse, err error) {
	if request == nil {
		request = NewGetSuggestionsRequest()
	}
	response = NewGetSuggestionsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPatrolCaseRequest() (request *ModifyPatrolCaseRequest) {
	request = &ModifyPatrolCaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("patrollor", APIVersion, "ModifyPatrolCase")
	return
}

func NewModifyPatrolCaseResponse() (response *ModifyPatrolCaseResponse) {
	response = &ModifyPatrolCaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新巡检用例
func (c *Client) ModifyPatrolCase(request *ModifyPatrolCaseRequest) (response *ModifyPatrolCaseResponse, err error) {
	if request == nil {
		request = NewModifyPatrolCaseRequest()
	}
	response = NewModifyPatrolCaseResponse()
	err = c.Send(request, response)
	return
}

func NewRunPatrolCaseRequest() (request *RunPatrolCaseRequest) {
	request = &RunPatrolCaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("patrollor", APIVersion, "RunPatrolCase")
	return
}

func NewRunPatrolCaseResponse() (response *RunPatrolCaseResponse) {
	response = &RunPatrolCaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 执行指定的巡检用例
func (c *Client) RunPatrolCase(request *RunPatrolCaseRequest) (response *RunPatrolCaseResponse, err error) {
	if request == nil {
		request = NewRunPatrolCaseRequest()
	}
	response = NewRunPatrolCaseResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationsRequest() (request *DescribeOperationsRequest) {
	request = &DescribeOperationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("patrollor", APIVersion, "DescribeOperations")
	return
}

func NewDescribeOperationsResponse() (response *DescribeOperationsResponse) {
	response = &DescribeOperationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取巡检平台上的操作记录
func (c *Client) DescribeOperations(request *DescribeOperationsRequest) (response *DescribeOperationsResponse, err error) {
	if request == nil {
		request = NewDescribeOperationsRequest()
	}
	response = NewDescribeOperationsResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePatrolCaseRequest() (request *CreatePatrolCaseRequest) {
	request = &CreatePatrolCaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("patrollor", APIVersion, "CreatePatrolCase")
	return
}

func NewCreatePatrolCaseResponse() (response *CreatePatrolCaseResponse) {
	response = &CreatePatrolCaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建巡检用例
func (c *Client) CreatePatrolCase(request *CreatePatrolCaseRequest) (response *CreatePatrolCaseResponse, err error) {
	if request == nil {
		request = NewCreatePatrolCaseRequest()
	}
	response = NewCreatePatrolCaseResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePatrolRecordsRequest() (request *DescribePatrolRecordsRequest) {
	request = &DescribePatrolRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("patrollor", APIVersion, "DescribePatrolRecords")
	return
}

func NewDescribePatrolRecordsResponse() (response *DescribePatrolRecordsResponse) {
	response = &DescribePatrolRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按指定条件获取巡检记录
func (c *Client) DescribePatrolRecords(request *DescribePatrolRecordsRequest) (response *DescribePatrolRecordsResponse, err error) {
	if request == nil {
		request = NewDescribePatrolRecordsRequest()
	}
	response = NewDescribePatrolRecordsResponse()
	err = c.Send(request, response)
	return
}
