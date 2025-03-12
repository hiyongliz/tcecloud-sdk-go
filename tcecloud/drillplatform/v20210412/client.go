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

package v20210412

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2021-04-12"

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

func NewGetOverviewInfoRequest() (request *GetOverviewInfoRequest) {
	request = &GetOverviewInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "GetOverviewInfo")
	return
}

func NewGetOverviewInfoResponse() (response *GetOverviewInfoResponse) {
	response = &GetOverviewInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取概览信息
func (c *Client) GetOverviewInfo(request *GetOverviewInfoRequest) (response *GetOverviewInfoResponse, err error) {
	if request == nil {
		request = NewGetOverviewInfoRequest()
	}
	response = NewGetOverviewInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGetSceneDrillLibraryBizNamesRequest() (request *GetSceneDrillLibraryBizNamesRequest) {
	request = &GetSceneDrillLibraryBizNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "GetSceneDrillLibraryBizNames")
	return
}

func NewGetSceneDrillLibraryBizNamesResponse() (response *GetSceneDrillLibraryBizNamesResponse) {
	response = &GetSceneDrillLibraryBizNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取场景演练库的biz列表
func (c *Client) GetSceneDrillLibraryBizNames(request *GetSceneDrillLibraryBizNamesRequest) (response *GetSceneDrillLibraryBizNamesResponse, err error) {
	if request == nil {
		request = NewGetSceneDrillLibraryBizNamesRequest()
	}
	response = NewGetSceneDrillLibraryBizNamesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteComponentDrillLibraryRequest() (request *DeleteComponentDrillLibraryRequest) {
	request = &DeleteComponentDrillLibraryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "DeleteComponentDrillLibrary")
	return
}

func NewDeleteComponentDrillLibraryResponse() (response *DeleteComponentDrillLibraryResponse) {
	response = &DeleteComponentDrillLibraryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除组件演练库
func (c *Client) DeleteComponentDrillLibrary(request *DeleteComponentDrillLibraryRequest) (response *DeleteComponentDrillLibraryResponse, err error) {
	if request == nil {
		request = NewDeleteComponentDrillLibraryRequest()
	}
	response = NewDeleteComponentDrillLibraryResponse()
	err = c.Send(request, response)
	return
}

func NewReserveSceneDrillRequest() (request *ReserveSceneDrillRequest) {
	request = &ReserveSceneDrillRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "ReserveSceneDrill")
	return
}

func NewReserveSceneDrillResponse() (response *ReserveSceneDrillResponse) {
	response = &ReserveSceneDrillResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预约场景演练库演练
func (c *Client) ReserveSceneDrill(request *ReserveSceneDrillRequest) (response *ReserveSceneDrillResponse, err error) {
	if request == nil {
		request = NewReserveSceneDrillRequest()
	}
	response = NewReserveSceneDrillResponse()
	err = c.Send(request, response)
	return
}

func NewEditComponentDrillLibraryRequest() (request *EditComponentDrillLibraryRequest) {
	request = &EditComponentDrillLibraryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "EditComponentDrillLibrary")
	return
}

func NewEditComponentDrillLibraryResponse() (response *EditComponentDrillLibraryResponse) {
	response = &EditComponentDrillLibraryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑组件演练库
func (c *Client) EditComponentDrillLibrary(request *EditComponentDrillLibraryRequest) (response *EditComponentDrillLibraryResponse, err error) {
	if request == nil {
		request = NewEditComponentDrillLibraryRequest()
	}
	response = NewEditComponentDrillLibraryResponse()
	err = c.Send(request, response)
	return
}

func NewExecuteSceneDrillLibraryRequest() (request *ExecuteSceneDrillLibraryRequest) {
	request = &ExecuteSceneDrillLibraryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "ExecuteSceneDrillLibrary")
	return
}

func NewExecuteSceneDrillLibraryResponse() (response *ExecuteSceneDrillLibraryResponse) {
	response = &ExecuteSceneDrillLibraryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 执行场景演练库演练
func (c *Client) ExecuteSceneDrillLibrary(request *ExecuteSceneDrillLibraryRequest) (response *ExecuteSceneDrillLibraryResponse, err error) {
	if request == nil {
		request = NewExecuteSceneDrillLibraryRequest()
	}
	response = NewExecuteSceneDrillLibraryResponse()
	err = c.Send(request, response)
	return
}

func NewCancelReserveSceneDrillRequest() (request *CancelReserveSceneDrillRequest) {
	request = &CancelReserveSceneDrillRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "CancelReserveSceneDrill")
	return
}

func NewCancelReserveSceneDrillResponse() (response *CancelReserveSceneDrillResponse) {
	response = &CancelReserveSceneDrillResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消预约场景演练库演练
func (c *Client) CancelReserveSceneDrill(request *CancelReserveSceneDrillRequest) (response *CancelReserveSceneDrillResponse, err error) {
	if request == nil {
		request = NewCancelReserveSceneDrillRequest()
	}
	response = NewCancelReserveSceneDrillResponse()
	err = c.Send(request, response)
	return
}

func NewGetSubTypeListRequest() (request *GetSubTypeListRequest) {
	request = &GetSubTypeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "GetSubTypeList")
	return
}

func NewGetSubTypeListResponse() (response *GetSubTypeListResponse) {
	response = &GetSubTypeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取子类型列表
func (c *Client) GetSubTypeList(request *GetSubTypeListRequest) (response *GetSubTypeListResponse, err error) {
	if request == nil {
		request = NewGetSubTypeListRequest()
	}
	response = NewGetSubTypeListResponse()
	err = c.Send(request, response)
	return
}

func NewGetTaskListRequest() (request *GetTaskListRequest) {
	request = &GetTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "GetTaskList")
	return
}

func NewGetTaskListResponse() (response *GetTaskListResponse) {
	response = &GetTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询任务列表
func (c *Client) GetTaskList(request *GetTaskListRequest) (response *GetTaskListResponse, err error) {
	if request == nil {
		request = NewGetTaskListRequest()
	}
	response = NewGetTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSceneDrillLibraryRequest() (request *DeleteSceneDrillLibraryRequest) {
	request = &DeleteSceneDrillLibraryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "DeleteSceneDrillLibrary")
	return
}

func NewDeleteSceneDrillLibraryResponse() (response *DeleteSceneDrillLibraryResponse) {
	response = &DeleteSceneDrillLibraryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除场景演练库
func (c *Client) DeleteSceneDrillLibrary(request *DeleteSceneDrillLibraryRequest) (response *DeleteSceneDrillLibraryResponse, err error) {
	if request == nil {
		request = NewDeleteSceneDrillLibraryRequest()
	}
	response = NewDeleteSceneDrillLibraryResponse()
	err = c.Send(request, response)
	return
}

func NewReExecuteTaskRequest() (request *ReExecuteTaskRequest) {
	request = &ReExecuteTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "ReExecuteTask")
	return
}

func NewReExecuteTaskResponse() (response *ReExecuteTaskResponse) {
	response = &ReExecuteTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 再次执行任务
func (c *Client) ReExecuteTask(request *ReExecuteTaskRequest) (response *ReExecuteTaskResponse, err error) {
	if request == nil {
		request = NewReExecuteTaskRequest()
	}
	response = NewReExecuteTaskResponse()
	err = c.Send(request, response)
	return
}

func NewGetMainTypeListRequest() (request *GetMainTypeListRequest) {
	request = &GetMainTypeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "GetMainTypeList")
	return
}

func NewGetMainTypeListResponse() (response *GetMainTypeListResponse) {
	response = &GetMainTypeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主类型列表
func (c *Client) GetMainTypeList(request *GetMainTypeListRequest) (response *GetMainTypeListResponse, err error) {
	if request == nil {
		request = NewGetMainTypeListRequest()
	}
	response = NewGetMainTypeListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSceneDrillLibraryRequest() (request *CreateSceneDrillLibraryRequest) {
	request = &CreateSceneDrillLibraryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "CreateSceneDrillLibrary")
	return
}

func NewCreateSceneDrillLibraryResponse() (response *CreateSceneDrillLibraryResponse) {
	response = &CreateSceneDrillLibraryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建场景演练库
func (c *Client) CreateSceneDrillLibrary(request *CreateSceneDrillLibraryRequest) (response *CreateSceneDrillLibraryResponse, err error) {
	if request == nil {
		request = NewCreateSceneDrillLibraryRequest()
	}
	response = NewCreateSceneDrillLibraryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComponentDrillLibraryRequest() (request *DescribeComponentDrillLibraryRequest) {
	request = &DescribeComponentDrillLibraryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "DescribeComponentDrillLibrary")
	return
}

func NewDescribeComponentDrillLibraryResponse() (response *DescribeComponentDrillLibraryResponse) {
	response = &DescribeComponentDrillLibraryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询组件演练库详情
func (c *Client) DescribeComponentDrillLibrary(request *DescribeComponentDrillLibraryRequest) (response *DescribeComponentDrillLibraryResponse, err error) {
	if request == nil {
		request = NewDescribeComponentDrillLibraryRequest()
	}
	response = NewDescribeComponentDrillLibraryResponse()
	err = c.Send(request, response)
	return
}

func NewGetComponentDrillLibraryBizNamesRequest() (request *GetComponentDrillLibraryBizNamesRequest) {
	request = &GetComponentDrillLibraryBizNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "GetComponentDrillLibraryBizNames")
	return
}

func NewGetComponentDrillLibraryBizNamesResponse() (response *GetComponentDrillLibraryBizNamesResponse) {
	response = &GetComponentDrillLibraryBizNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组件演练库的biz列表
func (c *Client) GetComponentDrillLibraryBizNames(request *GetComponentDrillLibraryBizNamesRequest) (response *GetComponentDrillLibraryBizNamesResponse, err error) {
	if request == nil {
		request = NewGetComponentDrillLibraryBizNamesRequest()
	}
	response = NewGetComponentDrillLibraryBizNamesResponse()
	err = c.Send(request, response)
	return
}

func NewTerminateTaskRequest() (request *TerminateTaskRequest) {
	request = &TerminateTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "TerminateTask")
	return
}

func NewTerminateTaskResponse() (response *TerminateTaskResponse) {
	response = &TerminateTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 终止任务
func (c *Client) TerminateTask(request *TerminateTaskRequest) (response *TerminateTaskResponse, err error) {
	if request == nil {
		request = NewTerminateTaskRequest()
	}
	response = NewTerminateTaskResponse()
	err = c.Send(request, response)
	return
}

func NewEditSceneDrillLibraryRequest() (request *EditSceneDrillLibraryRequest) {
	request = &EditSceneDrillLibraryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "EditSceneDrillLibrary")
	return
}

func NewEditSceneDrillLibraryResponse() (response *EditSceneDrillLibraryResponse) {
	response = &EditSceneDrillLibraryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑场景演练库
func (c *Client) EditSceneDrillLibrary(request *EditSceneDrillLibraryRequest) (response *EditSceneDrillLibraryResponse, err error) {
	if request == nil {
		request = NewEditSceneDrillLibraryRequest()
	}
	response = NewEditSceneDrillLibraryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExperienceLibraryRequest() (request *DescribeExperienceLibraryRequest) {
	request = &DescribeExperienceLibraryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "DescribeExperienceLibrary")
	return
}

func NewDescribeExperienceLibraryResponse() (response *DescribeExperienceLibraryResponse) {
	response = &DescribeExperienceLibraryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询经验库详情
func (c *Client) DescribeExperienceLibrary(request *DescribeExperienceLibraryRequest) (response *DescribeExperienceLibraryResponse, err error) {
	if request == nil {
		request = NewDescribeExperienceLibraryRequest()
	}
	response = NewDescribeExperienceLibraryResponse()
	err = c.Send(request, response)
	return
}

func NewEditExperienceLibraryRequest() (request *EditExperienceLibraryRequest) {
	request = &EditExperienceLibraryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "EditExperienceLibrary")
	return
}

func NewEditExperienceLibraryResponse() (response *EditExperienceLibraryResponse) {
	response = &EditExperienceLibraryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑经验库
func (c *Client) EditExperienceLibrary(request *EditExperienceLibraryRequest) (response *EditExperienceLibraryResponse, err error) {
	if request == nil {
		request = NewEditExperienceLibraryRequest()
	}
	response = NewEditExperienceLibraryResponse()
	err = c.Send(request, response)
	return
}

func NewGetTaskInfoRequest() (request *GetTaskInfoRequest) {
	request = &GetTaskInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "GetTaskInfo")
	return
}

func NewGetTaskInfoResponse() (response *GetTaskInfoResponse) {
	response = &GetTaskInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看任务详情
func (c *Client) GetTaskInfo(request *GetTaskInfoRequest) (response *GetTaskInfoResponse, err error) {
	if request == nil {
		request = NewGetTaskInfoRequest()
	}
	response = NewGetTaskInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateComponentDrillLibraryRequest() (request *CreateComponentDrillLibraryRequest) {
	request = &CreateComponentDrillLibraryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "CreateComponentDrillLibrary")
	return
}

func NewCreateComponentDrillLibraryResponse() (response *CreateComponentDrillLibraryResponse) {
	response = &CreateComponentDrillLibraryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建组件演练库
func (c *Client) CreateComponentDrillLibrary(request *CreateComponentDrillLibraryRequest) (response *CreateComponentDrillLibraryResponse, err error) {
	if request == nil {
		request = NewCreateComponentDrillLibraryRequest()
	}
	response = NewCreateComponentDrillLibraryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSceneDrillLibraryRequest() (request *DescribeSceneDrillLibraryRequest) {
	request = &DescribeSceneDrillLibraryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "DescribeSceneDrillLibrary")
	return
}

func NewDescribeSceneDrillLibraryResponse() (response *DescribeSceneDrillLibraryResponse) {
	response = &DescribeSceneDrillLibraryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询场景演练库详情
func (c *Client) DescribeSceneDrillLibrary(request *DescribeSceneDrillLibraryRequest) (response *DescribeSceneDrillLibraryResponse, err error) {
	if request == nil {
		request = NewDescribeSceneDrillLibraryRequest()
	}
	response = NewDescribeSceneDrillLibraryResponse()
	err = c.Send(request, response)
	return
}

func NewGetExperienceLibraryListRequest() (request *GetExperienceLibraryListRequest) {
	request = &GetExperienceLibraryListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "GetExperienceLibraryList")
	return
}

func NewGetExperienceLibraryListResponse() (response *GetExperienceLibraryListResponse) {
	response = &GetExperienceLibraryListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询经验库列表
func (c *Client) GetExperienceLibraryList(request *GetExperienceLibraryListRequest) (response *GetExperienceLibraryListResponse, err error) {
	if request == nil {
		request = NewGetExperienceLibraryListRequest()
	}
	response = NewGetExperienceLibraryListResponse()
	err = c.Send(request, response)
	return
}

func NewReserveComponentDrillRequest() (request *ReserveComponentDrillRequest) {
	request = &ReserveComponentDrillRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "ReserveComponentDrill")
	return
}

func NewReserveComponentDrillResponse() (response *ReserveComponentDrillResponse) {
	response = &ReserveComponentDrillResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预约组件演练库演练
func (c *Client) ReserveComponentDrill(request *ReserveComponentDrillRequest) (response *ReserveComponentDrillResponse, err error) {
	if request == nil {
		request = NewReserveComponentDrillRequest()
	}
	response = NewReserveComponentDrillResponse()
	err = c.Send(request, response)
	return
}

func NewGetComponentDrillLibraryListRequest() (request *GetComponentDrillLibraryListRequest) {
	request = &GetComponentDrillLibraryListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "GetComponentDrillLibraryList")
	return
}

func NewGetComponentDrillLibraryListResponse() (response *GetComponentDrillLibraryListResponse) {
	response = &GetComponentDrillLibraryListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询组件演练库列表
func (c *Client) GetComponentDrillLibraryList(request *GetComponentDrillLibraryListRequest) (response *GetComponentDrillLibraryListResponse, err error) {
	if request == nil {
		request = NewGetComponentDrillLibraryListRequest()
	}
	response = NewGetComponentDrillLibraryListResponse()
	err = c.Send(request, response)
	return
}

func NewGetExperienceLibraryCountRequest() (request *GetExperienceLibraryCountRequest) {
	request = &GetExperienceLibraryCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "GetExperienceLibraryCount")
	return
}

func NewGetExperienceLibraryCountResponse() (response *GetExperienceLibraryCountResponse) {
	response = &GetExperienceLibraryCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取各类型经验库数量
func (c *Client) GetExperienceLibraryCount(request *GetExperienceLibraryCountRequest) (response *GetExperienceLibraryCountResponse, err error) {
	if request == nil {
		request = NewGetExperienceLibraryCountRequest()
	}
	response = NewGetExperienceLibraryCountResponse()
	err = c.Send(request, response)
	return
}

func NewGetSceneDrillLibraryListRequest() (request *GetSceneDrillLibraryListRequest) {
	request = &GetSceneDrillLibraryListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "GetSceneDrillLibraryList")
	return
}

func NewGetSceneDrillLibraryListResponse() (response *GetSceneDrillLibraryListResponse) {
	response = &GetSceneDrillLibraryListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询场景演练库列表
func (c *Client) GetSceneDrillLibraryList(request *GetSceneDrillLibraryListRequest) (response *GetSceneDrillLibraryListResponse, err error) {
	if request == nil {
		request = NewGetSceneDrillLibraryListRequest()
	}
	response = NewGetSceneDrillLibraryListResponse()
	err = c.Send(request, response)
	return
}

func NewGetOpRecordListRequest() (request *GetOpRecordListRequest) {
	request = &GetOpRecordListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "GetOpRecordList")
	return
}

func NewGetOpRecordListResponse() (response *GetOpRecordListResponse) {
	response = &GetOpRecordListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取操作记录列表
func (c *Client) GetOpRecordList(request *GetOpRecordListRequest) (response *GetOpRecordListResponse, err error) {
	if request == nil {
		request = NewGetOpRecordListRequest()
	}
	response = NewGetOpRecordListResponse()
	err = c.Send(request, response)
	return
}

func NewGetBaseInfoRequest() (request *GetBaseInfoRequest) {
	request = &GetBaseInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "GetBaseInfo")
	return
}

func NewGetBaseInfoResponse() (response *GetBaseInfoResponse) {
	response = &GetBaseInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基础信息列表
func (c *Client) GetBaseInfo(request *GetBaseInfoRequest) (response *GetBaseInfoResponse, err error) {
	if request == nil {
		request = NewGetBaseInfoRequest()
	}
	response = NewGetBaseInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGetTargetTypeListRequest() (request *GetTargetTypeListRequest) {
	request = &GetTargetTypeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "GetTargetTypeList")
	return
}

func NewGetTargetTypeListResponse() (response *GetTargetTypeListResponse) {
	response = &GetTargetTypeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取支持的目标类型
func (c *Client) GetTargetTypeList(request *GetTargetTypeListRequest) (response *GetTargetTypeListResponse, err error) {
	if request == nil {
		request = NewGetTargetTypeListRequest()
	}
	response = NewGetTargetTypeListResponse()
	err = c.Send(request, response)
	return
}

func NewGetSubTaskListRequest() (request *GetSubTaskListRequest) {
	request = &GetSubTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "GetSubTaskList")
	return
}

func NewGetSubTaskListResponse() (response *GetSubTaskListResponse) {
	response = &GetSubTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取子任务列表
func (c *Client) GetSubTaskList(request *GetSubTaskListRequest) (response *GetSubTaskListResponse, err error) {
	if request == nil {
		request = NewGetSubTaskListRequest()
	}
	response = NewGetSubTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewExecuteComponentDrillLibraryRequest() (request *ExecuteComponentDrillLibraryRequest) {
	request = &ExecuteComponentDrillLibraryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("drillplatform", APIVersion, "ExecuteComponentDrillLibrary")
	return
}

func NewExecuteComponentDrillLibraryResponse() (response *ExecuteComponentDrillLibraryResponse) {
	response = &ExecuteComponentDrillLibraryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 执行组件演练库演练
func (c *Client) ExecuteComponentDrillLibrary(request *ExecuteComponentDrillLibraryRequest) (response *ExecuteComponentDrillLibraryResponse, err error) {
	if request == nil {
		request = NewExecuteComponentDrillLibraryRequest()
	}
	response = NewExecuteComponentDrillLibraryResponse()
	err = c.Send(request, response)
	return
}
