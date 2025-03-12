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

package v20211025

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2021-10-25"

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

func NewQueryAuditRequest() (request *QueryAuditRequest) {
	request = &QueryAuditRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("agentmanager", APIVersion, "QueryAudit")
	return
}

func NewQueryAuditResponse() (response *QueryAuditResponse) {
	response = &QueryAuditResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询审计日志列表
func (c *Client) QueryAudit(request *QueryAuditRequest) (response *QueryAuditResponse, err error) {
	if request == nil {
		request = NewQueryAuditRequest()
	}
	response = NewQueryAuditResponse()
	err = c.Send(request, response)
	return
}

func NewQueryAgentVersionRequest() (request *QueryAgentVersionRequest) {
	request = &QueryAgentVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("agentmanager", APIVersion, "QueryAgentVersion")
	return
}

func NewQueryAgentVersionResponse() (response *QueryAgentVersionResponse) {
	response = &QueryAgentVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询 agent 版本列表
func (c *Client) QueryAgentVersion(request *QueryAgentVersionRequest) (response *QueryAgentVersionResponse, err error) {
	if request == nil {
		request = NewQueryAgentVersionRequest()
	}
	response = NewQueryAgentVersionResponse()
	err = c.Send(request, response)
	return
}

func NewUninstallRequest() (request *UninstallRequest) {
	request = &UninstallRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("agentmanager", APIVersion, "Uninstall")
	return
}

func NewUninstallResponse() (response *UninstallResponse) {
	response = &UninstallResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 卸载agent
func (c *Client) Uninstall(request *UninstallRequest) (response *UninstallResponse, err error) {
	if request == nil {
		request = NewUninstallRequest()
	}
	response = NewUninstallResponse()
	err = c.Send(request, response)
	return
}

func NewStartRequest() (request *StartRequest) {
	request = &StartRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("agentmanager", APIVersion, "Start")
	return
}

func NewStartResponse() (response *StartResponse) {
	response = &StartResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动agent，同样包括启动主控agent和被管理的agent
func (c *Client) Start(request *StartRequest) (response *StartResponse, err error) {
	if request == nil {
		request = NewStartRequest()
	}
	response = NewStartResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateRequest() (request *UpdateRequest) {
	request = &UpdateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("agentmanager", APIVersion, "Update")
	return
}

func NewUpdateResponse() (response *UpdateResponse) {
	response = &UpdateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新agent
func (c *Client) Update(request *UpdateRequest) (response *UpdateResponse, err error) {
	if request == nil {
		request = NewUpdateRequest()
	}
	response = NewUpdateResponse()
	err = c.Send(request, response)
	return
}

func NewInstallRequest() (request *InstallRequest) {
	request = &InstallRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("agentmanager", APIVersion, "Install")
	return
}

func NewInstallResponse() (response *InstallResponse) {
	response = &InstallResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安装agent（包括主控agent和被管理的agent）；
func (c *Client) Install(request *InstallRequest) (response *InstallResponse, err error) {
	if request == nil {
		request = NewInstallRequest()
	}
	response = NewInstallResponse()
	err = c.Send(request, response)
	return
}

func NewRestartRequest() (request *RestartRequest) {
	request = &RestartRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("agentmanager", APIVersion, "Restart")
	return
}

func NewRestartResponse() (response *RestartResponse) {
	response = &RestartResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重启agent
func (c *Client) Restart(request *RestartRequest) (response *RestartResponse, err error) {
	if request == nil {
		request = NewRestartRequest()
	}
	response = NewRestartResponse()
	err = c.Send(request, response)
	return
}

func NewQueryAgentConfigRequest() (request *QueryAgentConfigRequest) {
	request = &QueryAgentConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("agentmanager", APIVersion, "QueryAgentConfig")
	return
}

func NewQueryAgentConfigResponse() (response *QueryAgentConfigResponse) {
	response = &QueryAgentConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询 agent 配置列表
func (c *Client) QueryAgentConfig(request *QueryAgentConfigRequest) (response *QueryAgentConfigResponse, err error) {
	if request == nil {
		request = NewQueryAgentConfigRequest()
	}
	response = NewQueryAgentConfigResponse()
	err = c.Send(request, response)
	return
}

func NewStopRequest() (request *StopRequest) {
	request = &StopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("agentmanager", APIVersion, "Stop")
	return
}

func NewStopResponse() (response *StopResponse) {
	response = &StopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停止agent
func (c *Client) Stop(request *StopRequest) (response *StopResponse, err error) {
	if request == nil {
		request = NewStopRequest()
	}
	response = NewStopResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAgentConfigRequest() (request *UpdateAgentConfigRequest) {
	request = &UpdateAgentConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("agentmanager", APIVersion, "UpdateAgentConfig")
	return
}

func NewUpdateAgentConfigResponse() (response *UpdateAgentConfigResponse) {
	response = &UpdateAgentConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新agent配置，目前只能更新agent关联的业务模块
func (c *Client) UpdateAgentConfig(request *UpdateAgentConfigRequest) (response *UpdateAgentConfigResponse, err error) {
	if request == nil {
		request = NewUpdateAgentConfigRequest()
	}
	response = NewUpdateAgentConfigResponse()
	err = c.Send(request, response)
	return
}

func NewQueryAgentAliasRequest() (request *QueryAgentAliasRequest) {
	request = &QueryAgentAliasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("agentmanager", APIVersion, "QueryAgentAlias")
	return
}

func NewQueryAgentAliasResponse() (response *QueryAgentAliasResponse) {
	response = &QueryAgentAliasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所有的 Agent 的名称和别名映射
func (c *Client) QueryAgentAlias(request *QueryAgentAliasRequest) (response *QueryAgentAliasResponse, err error) {
	if request == nil {
		request = NewQueryAgentAliasRequest()
	}
	response = NewQueryAgentAliasResponse()
	err = c.Send(request, response)
	return
}

func NewQueryHostRequest() (request *QueryHostRequest) {
	request = &QueryHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("agentmanager", APIVersion, "QueryHost")
	return
}

func NewQueryHostResponse() (response *QueryHostResponse) {
	response = &QueryHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定 agent 关联的业务模块下包含的主机及其安装的 agent 实例
func (c *Client) QueryHost(request *QueryHostRequest) (response *QueryHostResponse, err error) {
	if request == nil {
		request = NewQueryHostRequest()
	}
	response = NewQueryHostResponse()
	err = c.Send(request, response)
	return
}
