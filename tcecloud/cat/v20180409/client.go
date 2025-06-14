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

package v20180409

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2018-04-09"

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

func NewGetRespTimeTrendExRequest() (request *GetRespTimeTrendExRequest) {
	request = &GetRespTimeTrendExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "GetRespTimeTrendEx")
	return
}

func NewGetRespTimeTrendExResponse() (response *GetRespTimeTrendExResponse) {
	response = &GetRespTimeTrendExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询拨测任务的走势数据
func (c *Client) GetRespTimeTrendEx(request *GetRespTimeTrendExRequest) (response *GetRespTimeTrendExResponse, err error) {
	if request == nil {
		request = NewGetRespTimeTrendExRequest()
	}
	response = NewGetRespTimeTrendExResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTaskExRequest() (request *CreateTaskExRequest) {
	request = &CreateTaskExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "CreateTaskEx")
	return
}

func NewCreateTaskExResponse() (response *CreateTaskExResponse) {
	response = &CreateTaskExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建拨测任务(扩充)
func (c *Client) CreateTaskEx(request *CreateTaskExRequest) (response *CreateTaskExResponse, err error) {
	if request == nil {
		request = NewCreateTaskExRequest()
	}
	response = NewCreateTaskExResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAlarmPloicyRequest() (request *CreateAlarmPloicyRequest) {
	request = &CreateAlarmPloicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "CreateAlarmPloicy")
	return
}

func NewCreateAlarmPloicyResponse() (response *CreateAlarmPloicyResponse) {
	response = &CreateAlarmPloicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 为拨测任务创建告警策略
func (c *Client) CreateAlarmPloicy(request *CreateAlarmPloicyRequest) (response *CreateAlarmPloicyResponse, err error) {
	if request == nil {
		request = NewCreateAlarmPloicyRequest()
	}
	response = NewCreateAlarmPloicyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAgentGroupRequest() (request *ModifyAgentGroupRequest) {
	request = &ModifyAgentGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "ModifyAgentGroup")
	return
}

func NewModifyAgentGroupResponse() (response *ModifyAgentGroupResponse) {
	response = &ModifyAgentGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改拨测分组
func (c *Client) ModifyAgentGroup(request *ModifyAgentGroupRequest) (response *ModifyAgentGroupResponse, err error) {
	if request == nil {
		request = NewModifyAgentGroupRequest()
	}
	response = NewModifyAgentGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTaskExRequest() (request *ModifyTaskExRequest) {
	request = &ModifyTaskExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "ModifyTaskEx")
	return
}

func NewModifyTaskExResponse() (response *ModifyTaskExResponse) {
	response = &ModifyTaskExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改拨测任务(扩展)
func (c *Client) ModifyTaskEx(request *ModifyTaskExRequest) (response *ModifyTaskExResponse, err error) {
	if request == nil {
		request = NewModifyTaskExRequest()
	}
	response = NewModifyTaskExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentGroupsRequest() (request *DescribeAgentGroupsRequest) {
	request = &DescribeAgentGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "DescribeAgentGroups")
	return
}

func NewDescribeAgentGroupsResponse() (response *DescribeAgentGroupsResponse) {
	response = &DescribeAgentGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询拨测分组列表
func (c *Client) DescribeAgentGroups(request *DescribeAgentGroupsRequest) (response *DescribeAgentGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeAgentGroupsRequest()
	}
	response = NewDescribeAgentGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmsByTaskRequest() (request *DescribeAlarmsByTaskRequest) {
	request = &DescribeAlarmsByTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "DescribeAlarmsByTask")
	return
}

func NewDescribeAlarmsByTaskResponse() (response *DescribeAlarmsByTaskResponse) {
	response = &DescribeAlarmsByTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按任务查询拨测告警列表
func (c *Client) DescribeAlarmsByTask(request *DescribeAlarmsByTaskRequest) (response *DescribeAlarmsByTaskResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmsByTaskRequest()
	}
	response = NewDescribeAlarmsByTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
	request = &DescribeTaskDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "DescribeTaskDetail")
	return
}

func NewDescribeTaskDetailResponse() (response *DescribeTaskDetailResponse) {
	response = &DescribeTaskDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询拨测任务信息
func (c *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
	if request == nil {
		request = NewDescribeTaskDetailRequest()
	}
	response = NewDescribeTaskDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTaskRequest() (request *ModifyTaskRequest) {
	request = &ModifyTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "ModifyTask")
	return
}

func NewModifyTaskResponse() (response *ModifyTaskResponse) {
	response = &ModifyTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改 拨测任务。
// 如果验证未成功，请先验证成功。避免修改为失败率100%的拨测任务。
func (c *Client) ModifyTask(request *ModifyTaskRequest) (response *ModifyTaskResponse, err error) {
	if request == nil {
		request = NewModifyTaskRequest()
	}
	response = NewModifyTaskResponse()
	err = c.Send(request, response)
	return
}

func NewGetRespTimeTrendRequest() (request *GetRespTimeTrendRequest) {
	request = &GetRespTimeTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "GetRespTimeTrend")
	return
}

func NewGetRespTimeTrendResponse() (response *GetRespTimeTrendResponse) {
	response = &GetRespTimeTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询拨测任务的统计数据
func (c *Client) GetRespTimeTrend(request *GetRespTimeTrendRequest) (response *GetRespTimeTrendResponse, err error) {
	if request == nil {
		request = NewGetRespTimeTrendRequest()
	}
	response = NewGetRespTimeTrendResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTaskRequest() (request *CreateTaskRequest) {
	request = &CreateTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "CreateTask")
	return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
	response = &CreateTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建拨测任务（创建任务并发起验证）
//
// 操作提示：
// 下一步，请通过VerifyCatResult 接口，验证一下拨测验证是否成功。如果成功，则可通过RunTask 接口运行该任务。
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
	if request == nil {
		request = NewCreateTaskRequest()
	}
	response = NewCreateTaskResponse()
	err = c.Send(request, response)
	return
}

func NewGetResultSummaryRequest() (request *GetResultSummaryRequest) {
	request = &GetResultSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "GetResultSummary")
	return
}

func NewGetResultSummaryResponse() (response *GetResultSummaryResponse) {
	response = &GetResultSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取任务列表的实时数据
func (c *Client) GetResultSummary(request *GetResultSummaryRequest) (response *GetResultSummaryResponse, err error) {
	if request == nil {
		request = NewGetResultSummaryRequest()
	}
	response = NewGetResultSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewGetReturnCodeHistoryRequest() (request *GetReturnCodeHistoryRequest) {
	request = &GetReturnCodeHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "GetReturnCodeHistory")
	return
}

func NewGetReturnCodeHistoryResponse() (response *GetReturnCodeHistoryResponse) {
	response = &GetReturnCodeHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询拨测任务的历史返回码信息
func (c *Client) GetReturnCodeHistory(request *GetReturnCodeHistoryRequest) (response *GetReturnCodeHistoryResponse, err error) {
	if request == nil {
		request = NewGetReturnCodeHistoryRequest()
	}
	response = NewGetReturnCodeHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewPauseTaskRequest() (request *PauseTaskRequest) {
	request = &PauseTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "PauseTask")
	return
}

func NewPauseTaskResponse() (response *PauseTaskResponse) {
	response = &PauseTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 暂停拨测任务
func (c *Client) PauseTask(request *PauseTaskRequest) (response *PauseTaskResponse, err error) {
	if request == nil {
		request = NewPauseTaskRequest()
	}
	response = NewPauseTaskResponse()
	err = c.Send(request, response)
	return
}

func NewGetDailyAvailRatioRequest() (request *GetDailyAvailRatioRequest) {
	request = &GetDailyAvailRatioRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "GetDailyAvailRatio")
	return
}

func NewGetDailyAvailRatioResponse() (response *GetDailyAvailRatioResponse) {
	response = &GetDailyAvailRatioResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取一天的整体可用率信息
func (c *Client) GetDailyAvailRatio(request *GetDailyAvailRatioRequest) (response *GetDailyAvailRatioResponse, err error) {
	if request == nil {
		request = NewGetDailyAvailRatioRequest()
	}
	response = NewGetDailyAvailRatioResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAgentGroupRequest() (request *CreateAgentGroupRequest) {
	request = &CreateAgentGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "CreateAgentGroup")
	return
}

func NewCreateAgentGroupResponse() (response *CreateAgentGroupResponse) {
	response = &CreateAgentGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加拨测分组
func (c *Client) CreateAgentGroup(request *CreateAgentGroupRequest) (response *CreateAgentGroupResponse, err error) {
	if request == nil {
		request = NewCreateAgentGroupRequest()
	}
	response = NewCreateAgentGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmTopicRequest() (request *DescribeAlarmTopicRequest) {
	request = &DescribeAlarmTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "DescribeAlarmTopic")
	return
}

func NewDescribeAlarmTopicResponse() (response *DescribeAlarmTopicResponse) {
	response = &DescribeAlarmTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户的告警主题列表
func (c *Client) DescribeAlarmTopic(request *DescribeAlarmTopicRequest) (response *DescribeAlarmTopicResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmTopicRequest()
	}
	response = NewDescribeAlarmTopicResponse()
	err = c.Send(request, response)
	return
}

func NewBindAlarmPloicyRequest() (request *BindAlarmPloicyRequest) {
	request = &BindAlarmPloicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "BindAlarmPloicy")
	return
}

func NewBindAlarmPloicyResponse() (response *BindAlarmPloicyResponse) {
	response = &BindAlarmPloicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定任务和告警策略组
func (c *Client) BindAlarmPloicy(request *BindAlarmPloicyRequest) (response *BindAlarmPloicyResponse, err error) {
	if request == nil {
		request = NewBindAlarmPloicyRequest()
	}
	response = NewBindAlarmPloicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTasksByTypeRequest() (request *DescribeTasksByTypeRequest) {
	request = &DescribeTasksByTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "DescribeTasksByType")
	return
}

func NewDescribeTasksByTypeResponse() (response *DescribeTasksByTypeResponse) {
	response = &DescribeTasksByTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按类型查询拨测任务列表
func (c *Client) DescribeTasksByType(request *DescribeTasksByTypeRequest) (response *DescribeTasksByTypeResponse, err error) {
	if request == nil {
		request = NewDescribeTasksByTypeRequest()
	}
	response = NewDescribeTasksByTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentGroupRequest() (request *DescribeAgentGroupRequest) {
	request = &DescribeAgentGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "DescribeAgentGroup")
	return
}

func NewDescribeAgentGroupResponse() (response *DescribeAgentGroupResponse) {
	response = &DescribeAgentGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询拨测分组详情
func (c *Client) DescribeAgentGroup(request *DescribeAgentGroupRequest) (response *DescribeAgentGroupResponse, err error) {
	if request == nil {
		request = NewDescribeAgentGroupRequest()
	}
	response = NewDescribeAgentGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAgentGroupRequest() (request *DeleteAgentGroupRequest) {
	request = &DeleteAgentGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "DeleteAgentGroup")
	return
}

func NewDeleteAgentGroupResponse() (response *DeleteAgentGroupResponse) {
	response = &DeleteAgentGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除拨测分组
func (c *Client) DeleteAgentGroup(request *DeleteAgentGroupRequest) (response *DeleteAgentGroupResponse, err error) {
	if request == nil {
		request = NewDeleteAgentGroupRequest()
	}
	response = NewDeleteAgentGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserLimitRequest() (request *DescribeUserLimitRequest) {
	request = &DescribeUserLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "DescribeUserLimit")
	return
}

func NewDescribeUserLimitResponse() (response *DescribeUserLimitResponse) {
	response = &DescribeUserLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户可用资源限制
func (c *Client) DescribeUserLimit(request *DescribeUserLimitRequest) (response *DescribeUserLimitResponse, err error) {
	if request == nil {
		request = NewDescribeUserLimitRequest()
	}
	response = NewDescribeUserLimitResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmGroupsRequest() (request *DescribeAlarmGroupsRequest) {
	request = &DescribeAlarmGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "DescribeAlarmGroups")
	return
}

func NewDescribeAlarmGroupsResponse() (response *DescribeAlarmGroupsResponse) {
	response = &DescribeAlarmGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户的告警接收组列表
func (c *Client) DescribeAlarmGroups(request *DescribeAlarmGroupsRequest) (response *DescribeAlarmGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmGroupsRequest()
	}
	response = NewDescribeAlarmGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAlarmPloicyRequest() (request *ModifyAlarmPloicyRequest) {
	request = &ModifyAlarmPloicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "ModifyAlarmPloicy")
	return
}

func NewModifyAlarmPloicyResponse() (response *ModifyAlarmPloicyResponse) {
	response = &ModifyAlarmPloicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 为拨测任务修改告警策略
func (c *Client) ModifyAlarmPloicy(request *ModifyAlarmPloicyRequest) (response *ModifyAlarmPloicyResponse, err error) {
	if request == nil {
		request = NewModifyAlarmPloicyRequest()
	}
	response = NewModifyAlarmPloicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmsRequest() (request *DescribeAlarmsRequest) {
	request = &DescribeAlarmsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "DescribeAlarms")
	return
}

func NewDescribeAlarmsResponse() (response *DescribeAlarmsResponse) {
	response = &DescribeAlarmsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询拨测告警列表
func (c *Client) DescribeAlarms(request *DescribeAlarmsRequest) (response *DescribeAlarmsResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmsRequest()
	}
	response = NewDescribeAlarmsResponse()
	err = c.Send(request, response)
	return
}

func NewVerifyResultRequest() (request *VerifyResultRequest) {
	request = &VerifyResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "VerifyResult")
	return
}

func NewVerifyResultResponse() (response *VerifyResultResponse) {
	response = &VerifyResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 验证拨测任务，结果验证查询（验证成功的，才建议创建拨测任务）
func (c *Client) VerifyResult(request *VerifyResultRequest) (response *VerifyResultResponse, err error) {
	if request == nil {
		request = NewVerifyResultRequest()
	}
	response = NewVerifyResultResponse()
	err = c.Send(request, response)
	return
}

func NewRunTaskRequest() (request *RunTaskRequest) {
	request = &RunTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "RunTask")
	return
}

func NewRunTaskResponse() (response *RunTaskResponse) {
	response = &RunTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行拨测任务
func (c *Client) RunTask(request *RunTaskRequest) (response *RunTaskResponse, err error) {
	if request == nil {
		request = NewRunTaskRequest()
	}
	response = NewRunTaskResponse()
	err = c.Send(request, response)
	return
}

func NewGetAvailRatioHistoryRequest() (request *GetAvailRatioHistoryRequest) {
	request = &GetAvailRatioHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "GetAvailRatioHistory")
	return
}

func NewGetAvailRatioHistoryResponse() (response *GetAvailRatioHistoryResponse) {
	response = &GetAvailRatioHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定时刻的可用率地图信息
func (c *Client) GetAvailRatioHistory(request *GetAvailRatioHistoryRequest) (response *GetAvailRatioHistoryResponse, err error) {
	if request == nil {
		request = NewGetAvailRatioHistoryRequest()
	}
	response = NewGetAvailRatioHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewGetTaskByWanIpRequest() (request *GetTaskByWanIpRequest) {
	request = &GetTaskByWanIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "GetTaskByWanIp")
	return
}

func NewGetTaskByWanIpResponse() (response *GetTaskByWanIpResponse) {
	response = &GetTaskByWanIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取本拨测点的任务列表
func (c *Client) GetTaskByWanIp(request *GetTaskByWanIpRequest) (response *GetTaskByWanIpResponse, err error) {
	if request == nil {
		request = NewGetTaskByWanIpRequest()
	}
	response = NewGetTaskByWanIpResponse()
	err = c.Send(request, response)
	return
}

func NewGetRealAvailRatioRequest() (request *GetRealAvailRatioRequest) {
	request = &GetRealAvailRatioRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "GetRealAvailRatio")
	return
}

func NewGetRealAvailRatioResponse() (response *GetRealAvailRatioResponse) {
	response = &GetRealAvailRatioResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取实时可用率信息
func (c *Client) GetRealAvailRatio(request *GetRealAvailRatioRequest) (response *GetRealAvailRatioResponse, err error) {
	if request == nil {
		request = NewGetRealAvailRatioRequest()
	}
	response = NewGetRealAvailRatioResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTasksRequest() (request *DeleteTasksRequest) {
	request = &DeleteTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "DeleteTasks")
	return
}

func NewDeleteTasksResponse() (response *DeleteTasksResponse) {
	response = &DeleteTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除多个拨测任务
func (c *Client) DeleteTasks(request *DeleteTasksRequest) (response *DeleteTasksResponse, err error) {
	if request == nil {
		request = NewDeleteTasksRequest()
	}
	response = NewDeleteTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentsRequest() (request *DescribeAgentsRequest) {
	request = &DescribeAgentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "DescribeAgents")
	return
}

func NewDescribeAgentsResponse() (response *DescribeAgentsResponse) {
	response = &DescribeAgentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询本用户可选的拨测点列表
func (c *Client) DescribeAgents(request *DescribeAgentsRequest) (response *DescribeAgentsResponse, err error) {
	if request == nil {
		request = NewDescribeAgentsRequest()
	}
	response = NewDescribeAgentsResponse()
	err = c.Send(request, response)
	return
}

func NewGetReturnCodeInfoRequest() (request *GetReturnCodeInfoRequest) {
	request = &GetReturnCodeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "GetReturnCodeInfo")
	return
}

func NewGetReturnCodeInfoResponse() (response *GetReturnCodeInfoResponse) {
	response = &GetReturnCodeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询拨测任务的返回码统计信息
func (c *Client) GetReturnCodeInfo(request *GetReturnCodeInfoRequest) (response *GetReturnCodeInfoResponse, err error) {
	if request == nil {
		request = NewGetReturnCodeInfoRequest()
	}
	response = NewGetReturnCodeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
	request = &DescribeTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "DescribeTask")
	return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
	response = &DescribeTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询拨测任务详情
func (c *Client) DescribeTask(request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
	if request == nil {
		request = NewDescribeTaskRequest()
	}
	response = NewDescribeTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCatLogsRequest() (request *DescribeCatLogsRequest) {
	request = &DescribeCatLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "DescribeCatLogs")
	return
}

func NewDescribeCatLogsResponse() (response *DescribeCatLogsResponse) {
	response = &DescribeCatLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询拨测流水
func (c *Client) DescribeCatLogs(request *DescribeCatLogsRequest) (response *DescribeCatLogsResponse, err error) {
	if request == nil {
		request = NewDescribeCatLogsRequest()
	}
	response = NewDescribeCatLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
	request = &DescribeTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cat", APIVersion, "DescribeTasks")
	return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
	response = &DescribeTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询拨测任务列表
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
	if request == nil {
		request = NewDescribeTasksRequest()
	}
	response = NewDescribeTasksResponse()
	err = c.Send(request, response)
	return
}
