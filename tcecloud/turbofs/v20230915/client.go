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

package v20230915

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2023-09-15"

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

func NewDescribeTurboVMConfigRequest() (request *DescribeTurboVMConfigRequest) {
	request = &DescribeTurboVMConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeTurboVMConfig")
	return
}

func NewDescribeTurboVMConfigResponse() (response *DescribeTurboVMConfigResponse) {
	response = &DescribeTurboVMConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询turbo虚拟机配置
func (c *Client) DescribeTurboVMConfig(request *DescribeTurboVMConfigRequest) (response *DescribeTurboVMConfigResponse, err error) {
	if request == nil {
		request = NewDescribeTurboVMConfigRequest()
	}
	response = NewDescribeTurboVMConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTurboTaskStateRequest() (request *ModifyTurboTaskStateRequest) {
	request = &ModifyTurboTaskStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "ModifyTurboTaskState")
	return
}

func NewModifyTurboTaskStateResponse() (response *ModifyTurboTaskStateResponse) {
	response = &ModifyTurboTaskStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 手动失败turbo任务
func (c *Client) ModifyTurboTaskState(request *ModifyTurboTaskStateRequest) (response *ModifyTurboTaskStateResponse, err error) {
	if request == nil {
		request = NewModifyTurboTaskStateRequest()
	}
	response = NewModifyTurboTaskStateResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCfsFileSystemRequest() (request *DeleteCfsFileSystemRequest) {
	request = &DeleteCfsFileSystemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DeleteCfsFileSystem")
	return
}

func NewDeleteCfsFileSystemResponse() (response *DeleteCfsFileSystemResponse) {
	response = &DeleteCfsFileSystemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除文件系统
func (c *Client) DeleteCfsFileSystem(request *DeleteCfsFileSystemRequest) (response *DeleteCfsFileSystemResponse, err error) {
	if request == nil {
		request = NewDeleteCfsFileSystemRequest()
	}
	response = NewDeleteCfsFileSystemResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTurboClusterStatusRequest() (request *DescribeTurboClusterStatusRequest) {
	request = &DescribeTurboClusterStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeTurboClusterStatus")
	return
}

func NewDescribeTurboClusterStatusResponse() (response *DescribeTurboClusterStatusResponse) {
	response = &DescribeTurboClusterStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询文件系统执行流程
func (c *Client) DescribeTurboClusterStatus(request *DescribeTurboClusterStatusRequest) (response *DescribeTurboClusterStatusResponse, err error) {
	if request == nil {
		request = NewDescribeTurboClusterStatusRequest()
	}
	response = NewDescribeTurboClusterStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTurboFsCvmInstancesRequest() (request *DescribeTurboFsCvmInstancesRequest) {
	request = &DescribeTurboFsCvmInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeTurboFsCvmInstances")
	return
}

func NewDescribeTurboFsCvmInstancesResponse() (response *DescribeTurboFsCvmInstancesResponse) {
	response = &DescribeTurboFsCvmInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询turbo文件系统的实例列表
func (c *Client) DescribeTurboFsCvmInstances(request *DescribeTurboFsCvmInstancesRequest) (response *DescribeTurboFsCvmInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeTurboFsCvmInstancesRequest()
	}
	response = NewDescribeTurboFsCvmInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTurboDiskConfigRequest() (request *ModifyTurboDiskConfigRequest) {
	request = &ModifyTurboDiskConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "ModifyTurboDiskConfig")
	return
}

func NewModifyTurboDiskConfigResponse() (response *ModifyTurboDiskConfigResponse) {
	response = &ModifyTurboDiskConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改turbo使用磁盘类型
func (c *Client) ModifyTurboDiskConfig(request *ModifyTurboDiskConfigRequest) (response *ModifyTurboDiskConfigResponse, err error) {
	if request == nil {
		request = NewModifyTurboDiskConfigRequest()
	}
	response = NewModifyTurboDiskConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTurboImageConfigRequest() (request *ModifyTurboImageConfigRequest) {
	request = &ModifyTurboImageConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "ModifyTurboImageConfig")
	return
}

func NewModifyTurboImageConfigResponse() (response *ModifyTurboImageConfigResponse) {
	response = &ModifyTurboImageConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改turbo镜像信息
func (c *Client) ModifyTurboImageConfig(request *ModifyTurboImageConfigRequest) (response *ModifyTurboImageConfigResponse, err error) {
	if request == nil {
		request = NewModifyTurboImageConfigRequest()
	}
	response = NewModifyTurboImageConfigResponse()
	err = c.Send(request, response)
	return
}

func NewScaleUpFileSystemRequest() (request *ScaleUpFileSystemRequest) {
	request = &ScaleUpFileSystemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "ScaleUpFileSystem")
	return
}

func NewScaleUpFileSystemResponse() (response *ScaleUpFileSystemResponse) {
	response = &ScaleUpFileSystemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 扩容文件系统
func (c *Client) ScaleUpFileSystem(request *ScaleUpFileSystemRequest) (response *ScaleUpFileSystemResponse, err error) {
	if request == nil {
		request = NewScaleUpFileSystemRequest()
	}
	response = NewScaleUpFileSystemResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateTurboCreateFailedRequest() (request *UpdateTurboCreateFailedRequest) {
	request = &UpdateTurboCreateFailedRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "UpdateTurboCreateFailed")
	return
}

func NewUpdateTurboCreateFailedResponse() (response *UpdateTurboCreateFailedResponse) {
	response = &UpdateTurboCreateFailedResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 手动失败文件系统状态
func (c *Client) UpdateTurboCreateFailed(request *UpdateTurboCreateFailedRequest) (response *UpdateTurboCreateFailedResponse, err error) {
	if request == nil {
		request = NewUpdateTurboCreateFailedRequest()
	}
	response = NewUpdateTurboCreateFailedResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTurboFsRequest() (request *DescribeTurboFsRequest) {
	request = &DescribeTurboFsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeTurboFs")
	return
}

func NewDescribeTurboFsResponse() (response *DescribeTurboFsResponse) {
	response = &DescribeTurboFsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询文件系统列表
func (c *Client) DescribeTurboFs(request *DescribeTurboFsRequest) (response *DescribeTurboFsResponse, err error) {
	if request == nil {
		request = NewDescribeTurboFsRequest()
	}
	response = NewDescribeTurboFsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTurboImageConfigRequest() (request *DescribeTurboImageConfigRequest) {
	request = &DescribeTurboImageConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeTurboImageConfig")
	return
}

func NewDescribeTurboImageConfigResponse() (response *DescribeTurboImageConfigResponse) {
	response = &DescribeTurboImageConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询turbo镜像配置
func (c *Client) DescribeTurboImageConfig(request *DescribeTurboImageConfigRequest) (response *DescribeTurboImageConfigResponse, err error) {
	if request == nil {
		request = NewDescribeTurboImageConfigRequest()
	}
	response = NewDescribeTurboImageConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTurboVMConfigRuleRequest() (request *DeleteTurboVMConfigRuleRequest) {
	request = &DeleteTurboVMConfigRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DeleteTurboVMConfigRule")
	return
}

func NewDeleteTurboVMConfigRuleResponse() (response *DeleteTurboVMConfigRuleResponse) {
	response = &DeleteTurboVMConfigRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除turbo虚拟机配置信息
func (c *Client) DeleteTurboVMConfigRule(request *DeleteTurboVMConfigRuleRequest) (response *DeleteTurboVMConfigRuleResponse, err error) {
	if request == nil {
		request = NewDeleteTurboVMConfigRuleRequest()
	}
	response = NewDeleteTurboVMConfigRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllTurboTasksRequest() (request *DescribeAllTurboTasksRequest) {
	request = &DescribeAllTurboTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeAllTurboTasks")
	return
}

func NewDescribeAllTurboTasksResponse() (response *DescribeAllTurboTasksResponse) {
	response = &DescribeAllTurboTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询turbo任务
func (c *Client) DescribeAllTurboTasks(request *DescribeAllTurboTasksRequest) (response *DescribeAllTurboTasksResponse, err error) {
	if request == nil {
		request = NewDescribeAllTurboTasksRequest()
	}
	response = NewDescribeAllTurboTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCfs3AmountsRequest() (request *DescribeCfs3AmountsRequest) {
	request = &DescribeCfs3AmountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeCfs3Amounts")
	return
}

func NewDescribeCfs3AmountsResponse() (response *DescribeCfs3AmountsResponse) {
	response = &DescribeCfs3AmountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询文件系统挂载点列表。
func (c *Client) DescribeCfs3Amounts(request *DescribeCfs3AmountsRequest) (response *DescribeCfs3AmountsResponse, err error) {
	if request == nil {
		request = NewDescribeCfs3AmountsRequest()
	}
	response = NewDescribeCfs3AmountsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTurboClientInfoRequest() (request *DescribeTurboClientInfoRequest) {
	request = &DescribeTurboClientInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeTurboClientInfo")
	return
}

func NewDescribeTurboClientInfoResponse() (response *DescribeTurboClientInfoResponse) {
	response = &DescribeTurboClientInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询文件系统客户端列表
func (c *Client) DescribeTurboClientInfo(request *DescribeTurboClientInfoRequest) (response *DescribeTurboClientInfoResponse, err error) {
	if request == nil {
		request = NewDescribeTurboClientInfoRequest()
	}
	response = NewDescribeTurboClientInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTurboVMConfigRuleRequest() (request *CreateTurboVMConfigRuleRequest) {
	request = &CreateTurboVMConfigRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "CreateTurboVMConfigRule")
	return
}

func NewCreateTurboVMConfigRuleResponse() (response *CreateTurboVMConfigRuleResponse) {
	response = &CreateTurboVMConfigRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加turbo虚拟机配置信息
func (c *Client) CreateTurboVMConfigRule(request *CreateTurboVMConfigRuleRequest) (response *CreateTurboVMConfigRuleResponse, err error) {
	if request == nil {
		request = NewCreateTurboVMConfigRuleRequest()
	}
	response = NewCreateTurboVMConfigRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTurboTaskByTaskIdRequest() (request *DescribeTurboTaskByTaskIdRequest) {
	request = &DescribeTurboTaskByTaskIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeTurboTaskByTaskId")
	return
}

func NewDescribeTurboTaskByTaskIdResponse() (response *DescribeTurboTaskByTaskIdResponse) {
	response = &DescribeTurboTaskByTaskIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据任务ID查询任务
func (c *Client) DescribeTurboTaskByTaskId(request *DescribeTurboTaskByTaskIdRequest) (response *DescribeTurboTaskByTaskIdResponse, err error) {
	if request == nil {
		request = NewDescribeTurboTaskByTaskIdRequest()
	}
	response = NewDescribeTurboTaskByTaskIdResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTurboTaskInfoRequest() (request *DescribeTurboTaskInfoRequest) {
	request = &DescribeTurboTaskInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeTurboTaskInfo")
	return
}

func NewDescribeTurboTaskInfoResponse() (response *DescribeTurboTaskInfoResponse) {
	response = &DescribeTurboTaskInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询turbo任务执行信息
func (c *Client) DescribeTurboTaskInfo(request *DescribeTurboTaskInfoRequest) (response *DescribeTurboTaskInfoResponse, err error) {
	if request == nil {
		request = NewDescribeTurboTaskInfoRequest()
	}
	response = NewDescribeTurboTaskInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTurboDiskConfigRequest() (request *DescribeTurboDiskConfigRequest) {
	request = &DescribeTurboDiskConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeTurboDiskConfig")
	return
}

func NewDescribeTurboDiskConfigResponse() (response *DescribeTurboDiskConfigResponse) {
	response = &DescribeTurboDiskConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询turbo使用磁盘信息
func (c *Client) DescribeTurboDiskConfig(request *DescribeTurboDiskConfigRequest) (response *DescribeTurboDiskConfigResponse, err error) {
	if request == nil {
		request = NewDescribeTurboDiskConfigRequest()
	}
	response = NewDescribeTurboDiskConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTurbofsAnalysisRequest() (request *DescribeTurbofsAnalysisRequest) {
	request = &DescribeTurbofsAnalysisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "DescribeTurbofsAnalysis")
	return
}

func NewDescribeTurbofsAnalysisResponse() (response *DescribeTurbofsAnalysisResponse) {
	response = &DescribeTurbofsAnalysisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取运营端概览信息
func (c *Client) DescribeTurbofsAnalysis(request *DescribeTurbofsAnalysisRequest) (response *DescribeTurbofsAnalysisResponse, err error) {
	if request == nil {
		request = NewDescribeTurbofsAnalysisRequest()
	}
	response = NewDescribeTurbofsAnalysisResponse()
	err = c.Send(request, response)
	return
}

func NewRetryTurboTaskRequest() (request *RetryTurboTaskRequest) {
	request = &RetryTurboTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("turbofs", APIVersion, "RetryTurboTask")
	return
}

func NewRetryTurboTaskResponse() (response *RetryTurboTaskResponse) {
	response = &RetryTurboTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重试turbo任务
func (c *Client) RetryTurboTask(request *RetryTurboTaskRequest) (response *RetryTurboTaskResponse, err error) {
	if request == nil {
		request = NewRetryTurboTaskRequest()
	}
	response = NewRetryTurboTaskResponse()
	err = c.Send(request, response)
	return
}
