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

package v20220616

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2022-06-16"

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

func NewDescribeJobRequest() (request *DescribeJobRequest) {
	request = &DescribeJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "DescribeJob")
	return
}

func NewDescribeJobResponse() (response *DescribeJobResponse) {
	response = &DescribeJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据请求中的jobUUID, 返回当前编排任务的执行总览
func (c *Client) DescribeJob(request *DescribeJobRequest) (response *DescribeJobResponse, err error) {
	if request == nil {
		request = NewDescribeJobRequest()
	}
	response = NewDescribeJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogsSummaryRequest() (request *DescribeLogsSummaryRequest) {
	request = &DescribeLogsSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "DescribeLogsSummary")
	return
}

func NewDescribeLogsSummaryResponse() (response *DescribeLogsSummaryResponse) {
	response = &DescribeLogsSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志汇总查询接口
func (c *Client) DescribeLogsSummary(request *DescribeLogsSummaryRequest) (response *DescribeLogsSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeLogsSummaryRequest()
	}
	response = NewDescribeLogsSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewResumeNodeRequest() (request *ResumeNodeRequest) {
	request = &ResumeNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "ResumeNode")
	return
}

func NewResumeNodeResponse() (response *ResumeNodeResponse) {
	response = &ResumeNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 对暂停的节点，进行恢复运行操作
func (c *Client) ResumeNode(request *ResumeNodeRequest) (response *ResumeNodeResponse, err error) {
	if request == nil {
		request = NewResumeNodeRequest()
	}
	response = NewResumeNodeResponse()
	err = c.Send(request, response)
	return
}

func NewPauseNodeRequest() (request *PauseNodeRequest) {
	request = &PauseNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "PauseNode")
	return
}

func NewPauseNodeResponse() (response *PauseNodeResponse) {
	response = &PauseNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 对运行中的节点暂停操作
func (c *Client) PauseNode(request *PauseNodeRequest) (response *PauseNodeResponse, err error) {
	if request == nil {
		request = NewPauseNodeRequest()
	}
	response = NewPauseNodeResponse()
	err = c.Send(request, response)
	return
}

func NewRetryNodeRequest() (request *RetryNodeRequest) {
	request = &RetryNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "RetryNode")
	return
}

func NewRetryNodeResponse() (response *RetryNodeResponse) {
	response = &RetryNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 对失败状态的节点进行重试操作
func (c *Client) RetryNode(request *RetryNodeRequest) (response *RetryNodeResponse, err error) {
	if request == nil {
		request = NewRetryNodeRequest()
	}
	response = NewRetryNodeResponse()
	err = c.Send(request, response)
	return
}

func NewTerminateJobRequest() (request *TerminateJobRequest) {
	request = &TerminateJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "TerminateJob")
	return
}

func NewTerminateJobResponse() (response *TerminateJobResponse) {
	response = &TerminateJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将运行态流程实例取消，取消后无法对流程实例和流程实例任何节点进行控制操作
func (c *Client) TerminateJob(request *TerminateJobRequest) (response *TerminateJobResponse, err error) {
	if request == nil {
		request = NewTerminateJobRequest()
	}
	response = NewTerminateJobResponse()
	err = c.Send(request, response)
	return
}

func NewRerunNodeRequest() (request *RerunNodeRequest) {
	request = &RerunNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "RerunNode")
	return
}

func NewRerunNodeResponse() (response *RerunNodeResponse) {
	response = &RerunNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 对失败状态的节点进行跳过操作，跳过的节点状态置为已跳过
func (c *Client) RerunNode(request *RerunNodeRequest) (response *RerunNodeResponse, err error) {
	if request == nil {
		request = NewRerunNodeRequest()
	}
	response = NewRerunNodeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJobOverviewRequest() (request *DescribeJobOverviewRequest) {
	request = &DescribeJobOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "DescribeJobOverview")
	return
}

func NewDescribeJobOverviewResponse() (response *DescribeJobOverviewResponse) {
	response = &DescribeJobOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据请求中的jobUUID, 返回当前编排任务的执行总览
func (c *Client) DescribeJobOverview(request *DescribeJobOverviewRequest) (response *DescribeJobOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeJobOverviewRequest()
	}
	response = NewDescribeJobOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewPauseJobRequest() (request *PauseJobRequest) {
	request = &PauseJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "PauseJob")
	return
}

func NewPauseJobResponse() (response *PauseJobResponse) {
	response = &PauseJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将运行态流程实例暂停，会将这个流程实例所有节点暂停
func (c *Client) PauseJob(request *PauseJobRequest) (response *PauseJobResponse, err error) {
	if request == nil {
		request = NewPauseJobRequest()
	}
	response = NewPauseJobResponse()
	err = c.Send(request, response)
	return
}

func NewRegisterPluginRequest() (request *RegisterPluginRequest) {
	request = &RegisterPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "RegisterPlugin")
	return
}

func NewRegisterPluginResponse() (response *RegisterPluginResponse) {
	response = &RegisterPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 注册插件
func (c *Client) RegisterPlugin(request *RegisterPluginRequest) (response *RegisterPluginResponse, err error) {
	if request == nil {
		request = NewRegisterPluginRequest()
	}
	response = NewRegisterPluginResponse()
	err = c.Send(request, response)
	return
}

func NewPutContextRequest() (request *PutContextRequest) {
	request = &PutContextRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "PutContext")
	return
}

func NewPutContextResponse() (response *PutContextResponse) {
	response = &PutContextResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 写上下文
func (c *Client) PutContext(request *PutContextRequest) (response *PutContextResponse, err error) {
	if request == nil {
		request = NewPutContextRequest()
	}
	response = NewPutContextResponse()
	err = c.Send(request, response)
	return
}

func NewSkipNodeRequest() (request *SkipNodeRequest) {
	request = &SkipNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "SkipNode")
	return
}

func NewSkipNodeResponse() (response *SkipNodeResponse) {
	response = &SkipNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 对失败状态的节点进行跳过操作，跳过的节点状态置为已跳过
func (c *Client) SkipNode(request *SkipNodeRequest) (response *SkipNodeResponse, err error) {
	if request == nil {
		request = NewSkipNodeRequest()
	}
	response = NewSkipNodeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJobDAGRequest() (request *DescribeJobDAGRequest) {
	request = &DescribeJobDAGRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "DescribeJobDAG")
	return
}

func NewDescribeJobDAGResponse() (response *DescribeJobDAGResponse) {
	response = &DescribeJobDAGResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询编排任务中任意节点状态和对应子节点的状态
func (c *Client) DescribeJobDAG(request *DescribeJobDAGRequest) (response *DescribeJobDAGResponse, err error) {
	if request == nil {
		request = NewDescribeJobDAGRequest()
	}
	response = NewDescribeJobDAGResponse()
	err = c.Send(request, response)
	return
}

func NewGetContextRequest() (request *GetContextRequest) {
	request = &GetContextRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "GetContext")
	return
}

func NewGetContextResponse() (response *GetContextResponse) {
	response = &GetContextResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 读取上下文
func (c *Client) GetContext(request *GetContextRequest) (response *GetContextResponse, err error) {
	if request == nil {
		request = NewGetContextRequest()
	}
	response = NewGetContextResponse()
	err = c.Send(request, response)
	return
}

func NewSetManualParamsRequest() (request *SetManualParamsRequest) {
	request = &SetManualParamsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "SetManualParams")
	return
}

func NewSetManualParamsResponse() (response *SetManualParamsResponse) {
	response = &SetManualParamsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置手动给定的修改参数
func (c *Client) SetManualParams(request *SetManualParamsRequest) (response *SetManualParamsResponse, err error) {
	if request == nil {
		request = NewSetManualParamsRequest()
	}
	response = NewSetManualParamsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogsDAGRequest() (request *DescribeLogsDAGRequest) {
	request = &DescribeLogsDAGRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "DescribeLogsDAG")
	return
}

func NewDescribeLogsDAGResponse() (response *DescribeLogsDAGResponse) {
	response = &DescribeLogsDAGResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 节点拓扑日志查询接口
func (c *Client) DescribeLogsDAG(request *DescribeLogsDAGRequest) (response *DescribeLogsDAGResponse, err error) {
	if request == nil {
		request = NewDescribeLogsDAGRequest()
	}
	response = NewDescribeLogsDAGResponse()
	err = c.Send(request, response)
	return
}

func NewResumeJobRequest() (request *ResumeJobRequest) {
	request = &ResumeJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "ResumeJob")
	return
}

func NewResumeJobResponse() (response *ResumeJobResponse) {
	response = &ResumeJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将运行态流程实例恢复，会将这个流程实例所有暂停的节点恢复运行
func (c *Client) ResumeJob(request *ResumeJobRequest) (response *ResumeJobResponse, err error) {
	if request == nil {
		request = NewResumeJobRequest()
	}
	response = NewResumeJobResponse()
	err = c.Send(request, response)
	return
}

func NewSetWaitingNodeStatusRequest() (request *SetWaitingNodeStatusRequest) {
	request = &SetWaitingNodeStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "SetWaitingNodeStatus")
	return
}

func NewSetWaitingNodeStatusResponse() (response *SetWaitingNodeStatusResponse) {
	response = &SetWaitingNodeStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将等待中的节点置为成功
func (c *Client) SetWaitingNodeStatus(request *SetWaitingNodeStatusRequest) (response *SetWaitingNodeStatusResponse, err error) {
	if request == nil {
		request = NewSetWaitingNodeStatusRequest()
	}
	response = NewSetWaitingNodeStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateJobRequest() (request *CreateJobRequest) {
	request = &CreateJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "CreateJob")
	return
}

func NewCreateJobResponse() (response *CreateJobResponse) {
	response = &CreateJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在创建解决方案版本时，获取对应解决方案接入模式列表
func (c *Client) CreateJob(request *CreateJobRequest) (response *CreateJobResponse, err error) {
	if request == nil {
		request = NewCreateJobRequest()
	}
	response = NewCreateJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeManualParamsRequest() (request *DescribeManualParamsRequest) {
	request = &DescribeManualParamsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "DescribeManualParams")
	return
}

func NewDescribeManualParamsResponse() (response *DescribeManualParamsResponse) {
	response = &DescribeManualParamsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询手动给定的修改参数
func (c *Client) DescribeManualParams(request *DescribeManualParamsRequest) (response *DescribeManualParamsResponse, err error) {
	if request == nil {
		request = NewDescribeManualParamsRequest()
	}
	response = NewDescribeManualParamsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNodeLogsRequest() (request *DescribeNodeLogsRequest) {
	request = &DescribeNodeLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnorch", APIVersion, "DescribeNodeLogs")
	return
}

func NewDescribeNodeLogsResponse() (response *DescribeNodeLogsResponse) {
	response = &DescribeNodeLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点的结构化执行日志
func (c *Client) DescribeNodeLogs(request *DescribeNodeLogsRequest) (response *DescribeNodeLogsResponse, err error) {
	if request == nil {
		request = NewDescribeNodeLogsRequest()
	}
	response = NewDescribeNodeLogsResponse()
	err = c.Send(request, response)
	return
}
