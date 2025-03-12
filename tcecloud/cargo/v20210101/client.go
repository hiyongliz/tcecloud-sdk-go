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

package v20210101

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2021-01-01"

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

func NewCleanImagesRequest() (request *CleanImagesRequest) {
	request = &CleanImagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "CleanImages")
	return
}

func NewCleanImagesResponse() (response *CleanImagesResponse) {
	response = &CleanImagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像清理
func (c *Client) CleanImages(request *CleanImagesRequest) (response *CleanImagesResponse, err error) {
	if request == nil {
		request = NewCleanImagesRequest()
	}
	response = NewCleanImagesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTagsRequest() (request *DescribeTagsRequest) {
	request = &DescribeTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "DescribeTags")
	return
}

func NewDescribeTagsResponse() (response *DescribeTagsResponse) {
	response = &DescribeTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定镜像的 Tag
func (c *Client) DescribeTags(request *DescribeTagsRequest) (response *DescribeTagsResponse, err error) {
	if request == nil {
		request = NewDescribeTagsRequest()
	}
	response = NewDescribeTagsResponse()
	err = c.Send(request, response)
	return
}

func NewListProjectMembersRequest() (request *ListProjectMembersRequest) {
	request = &ListProjectMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "ListProjectMembers")
	return
}

func NewListProjectMembersResponse() (response *ListProjectMembersResponse) {
	response = &ListProjectMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ListProjectMembers
func (c *Client) ListProjectMembers(request *ListProjectMembersRequest) (response *ListProjectMembersResponse, err error) {
	if request == nil {
		request = NewListProjectMembersRequest()
	}
	response = NewListProjectMembersResponse()
	err = c.Send(request, response)
	return
}

func NewModifyImageBuildingTaskRequest() (request *ModifyImageBuildingTaskRequest) {
	request = &ModifyImageBuildingTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "ModifyImageBuildingTask")
	return
}

func NewModifyImageBuildingTaskResponse() (response *ModifyImageBuildingTaskResponse) {
	response = &ModifyImageBuildingTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新镜像构建任务
func (c *Client) ModifyImageBuildingTask(request *ModifyImageBuildingTaskRequest) (response *ModifyImageBuildingTaskResponse, err error) {
	if request == nil {
		request = NewModifyImageBuildingTaskRequest()
	}
	response = NewModifyImageBuildingTaskResponse()
	err = c.Send(request, response)
	return
}

func NewGetGCStatusRequest() (request *GetGCStatusRequest) {
	request = &GetGCStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "GetGCStatus")
	return
}

func NewGetGCStatusResponse() (response *GetGCStatusResponse) {
	response = &GetGCStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetGCStatus
func (c *Client) GetGCStatus(request *GetGCStatusRequest) (response *GetGCStatusResponse, err error) {
	if request == nil {
		request = NewGetGCStatusRequest()
	}
	response = NewGetGCStatusResponse()
	err = c.Send(request, response)
	return
}

func NewImageScanningStatusRequest() (request *ImageScanningStatusRequest) {
	request = &ImageScanningStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "ImageScanningStatus")
	return
}

func NewImageScanningStatusResponse() (response *ImageScanningStatusResponse) {
	response = &ImageScanningStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取某一镜像的漏洞扫描状态
func (c *Client) ImageScanningStatus(request *ImageScanningStatusRequest) (response *ImageScanningStatusResponse, err error) {
	if request == nil {
		request = NewImageScanningStatusRequest()
	}
	response = NewImageScanningStatusResponse()
	err = c.Send(request, response)
	return
}

func NewResetTokenRequest() (request *ResetTokenRequest) {
	request = &ResetTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "ResetToken")
	return
}

func NewResetTokenResponse() (response *ResetTokenResponse) {
	response = &ResetTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置仓库的Token
func (c *Client) ResetToken(request *ResetTokenRequest) (response *ResetTokenResponse, err error) {
	if request == nil {
		request = NewResetTokenRequest()
	}
	response = NewResetTokenResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRetentionRequest() (request *CreateRetentionRequest) {
	request = &CreateRetentionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "CreateRetention")
	return
}

func NewCreateRetentionResponse() (response *CreateRetentionResponse) {
	response = &CreateRetentionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建镜像清理策略
func (c *Client) CreateRetention(request *CreateRetentionRequest) (response *CreateRetentionResponse, err error) {
	if request == nil {
		request = NewCreateRetentionRequest()
	}
	response = NewCreateRetentionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProjectsRequest() (request *DescribeProjectsRequest) {
	request = &DescribeProjectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "DescribeProjects")
	return
}

func NewDescribeProjectsResponse() (response *DescribeProjectsResponse) {
	response = &DescribeProjectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取项目列表
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
	if request == nil {
		request = NewDescribeProjectsRequest()
	}
	response = NewDescribeProjectsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRetentionRequest() (request *DescribeRetentionRequest) {
	request = &DescribeRetentionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "DescribeRetention")
	return
}

func NewDescribeRetentionResponse() (response *DescribeRetentionResponse) {
	response = &DescribeRetentionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取镜像清理策略
func (c *Client) DescribeRetention(request *DescribeRetentionRequest) (response *DescribeRetentionResponse, err error) {
	if request == nil {
		request = NewDescribeRetentionRequest()
	}
	response = NewDescribeRetentionResponse()
	err = c.Send(request, response)
	return
}

func NewListProjectsRequest() (request *ListProjectsRequest) {
	request = &ListProjectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "ListProjects")
	return
}

func NewListProjectsResponse() (response *ListProjectsResponse) {
	response = &ListProjectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据过滤条件列举仓库中的项目
func (c *Client) ListProjects(request *ListProjectsRequest) (response *ListProjectsResponse, err error) {
	if request == nil {
		request = NewListProjectsRequest()
	}
	response = NewListProjectsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGCHistoriesRequest() (request *DescribeGCHistoriesRequest) {
	request = &DescribeGCHistoriesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "DescribeGCHistories")
	return
}

func NewDescribeGCHistoriesResponse() (response *DescribeGCHistoriesResponse) {
	response = &DescribeGCHistoriesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取仓库容量回收历史记录
func (c *Client) DescribeGCHistories(request *DescribeGCHistoriesRequest) (response *DescribeGCHistoriesResponse, err error) {
	if request == nil {
		request = NewDescribeGCHistoriesRequest()
	}
	response = NewDescribeGCHistoriesResponse()
	err = c.Send(request, response)
	return
}

func NewGetAccountRequest() (request *GetAccountRequest) {
	request = &GetAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "GetAccount")
	return
}

func NewGetAccountResponse() (response *GetAccountResponse) {
	response = &GetAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 给指定用户获取仓库账号
func (c *Client) GetAccount(request *GetAccountRequest) (response *GetAccountResponse, err error) {
	if request == nil {
		request = NewGetAccountRequest()
	}
	response = NewGetAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGCHistoryLogRequest() (request *DescribeGCHistoryLogRequest) {
	request = &DescribeGCHistoryLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "DescribeGCHistoryLog")
	return
}

func NewDescribeGCHistoryLogResponse() (response *DescribeGCHistoryLogResponse) {
	response = &DescribeGCHistoryLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取仓库容量回收日志
func (c *Client) DescribeGCHistoryLog(request *DescribeGCHistoryLogRequest) (response *DescribeGCHistoryLogResponse, err error) {
	if request == nil {
		request = NewDescribeGCHistoryLogRequest()
	}
	response = NewDescribeGCHistoryLogResponse()
	err = c.Send(request, response)
	return
}

func NewGetGCScheduleRequest() (request *GetGCScheduleRequest) {
	request = &GetGCScheduleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "GetGCSchedule")
	return
}

func NewGetGCScheduleResponse() (response *GetGCScheduleResponse) {
	response = &GetGCScheduleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取仓库容量回收配置
func (c *Client) GetGCSchedule(request *GetGCScheduleRequest) (response *GetGCScheduleResponse, err error) {
	if request == nil {
		request = NewGetGCScheduleRequest()
	}
	response = NewGetGCScheduleResponse()
	err = c.Send(request, response)
	return
}

func NewBlockProjectRequest() (request *BlockProjectRequest) {
	request = &BlockProjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "BlockProject")
	return
}

func NewBlockProjectResponse() (response *BlockProjectResponse) {
	response = &BlockProjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像锁定/解锁
func (c *Client) BlockProject(request *BlockProjectRequest) (response *BlockProjectResponse, err error) {
	if request == nil {
		request = NewBlockProjectRequest()
	}
	response = NewBlockProjectResponse()
	err = c.Send(request, response)
	return
}

func NewStartImageBuildingTaskRequest() (request *StartImageBuildingTaskRequest) {
	request = &StartImageBuildingTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "StartImageBuildingTask")
	return
}

func NewStartImageBuildingTaskResponse() (response *StartImageBuildingTaskResponse) {
	response = &StartImageBuildingTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动镜像构建任务
func (c *Client) StartImageBuildingTask(request *StartImageBuildingTaskRequest) (response *StartImageBuildingTaskResponse, err error) {
	if request == nil {
		request = NewStartImageBuildingTaskRequest()
	}
	response = NewStartImageBuildingTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProjectRequest() (request *DeleteProjectRequest) {
	request = &DeleteProjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "DeleteProject")
	return
}

func NewDeleteProjectResponse() (response *DeleteProjectResponse) {
	response = &DeleteProjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除仓库中的项目
func (c *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
	if request == nil {
		request = NewDeleteProjectRequest()
	}
	response = NewDeleteProjectResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRegistryRequest() (request *DescribeRegistryRequest) {
	request = &DescribeRegistryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "DescribeRegistry")
	return
}

func NewDescribeRegistryResponse() (response *DescribeRegistryResponse) {
	response = &DescribeRegistryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取仓库信息
func (c *Client) DescribeRegistry(request *DescribeRegistryRequest) (response *DescribeRegistryResponse, err error) {
	if request == nil {
		request = NewDescribeRegistryRequest()
	}
	response = NewDescribeRegistryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulnerabilitiesRequest() (request *DescribeVulnerabilitiesRequest) {
	request = &DescribeVulnerabilitiesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "DescribeVulnerabilities")
	return
}

func NewDescribeVulnerabilitiesResponse() (response *DescribeVulnerabilitiesResponse) {
	response = &DescribeVulnerabilitiesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看镜像的漏洞信息
func (c *Client) DescribeVulnerabilities(request *DescribeVulnerabilitiesRequest) (response *DescribeVulnerabilitiesResponse, err error) {
	if request == nil {
		request = NewDescribeVulnerabilitiesRequest()
	}
	response = NewDescribeVulnerabilitiesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGCScheduleRequest() (request *CreateGCScheduleRequest) {
	request = &CreateGCScheduleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "CreateGCSchedule")
	return
}

func NewCreateGCScheduleResponse() (response *CreateGCScheduleResponse) {
	response = &CreateGCScheduleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建仓库回收配置
func (c *Client) CreateGCSchedule(request *CreateGCScheduleRequest) (response *CreateGCScheduleResponse, err error) {
	if request == nil {
		request = NewCreateGCScheduleRequest()
	}
	response = NewCreateGCScheduleResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAllProjectMembersRequest() (request *UpdateAllProjectMembersRequest) {
	request = &UpdateAllProjectMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "UpdateAllProjectMembers")
	return
}

func NewUpdateAllProjectMembersResponse() (response *UpdateAllProjectMembersResponse) {
	response = &UpdateAllProjectMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新所有项目成员
func (c *Client) UpdateAllProjectMembers(request *UpdateAllProjectMembersRequest) (response *UpdateAllProjectMembersResponse, err error) {
	if request == nil {
		request = NewUpdateAllProjectMembersRequest()
	}
	response = NewUpdateAllProjectMembersResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGCScheduleRequest() (request *UpdateGCScheduleRequest) {
	request = &UpdateGCScheduleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "UpdateGCSchedule")
	return
}

func NewUpdateGCScheduleResponse() (response *UpdateGCScheduleResponse) {
	response = &UpdateGCScheduleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新仓库容量回收配置
func (c *Client) UpdateGCSchedule(request *UpdateGCScheduleRequest) (response *UpdateGCScheduleResponse, err error) {
	if request == nil {
		request = NewUpdateGCScheduleRequest()
	}
	response = NewUpdateGCScheduleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTagRequest() (request *DeleteTagRequest) {
	request = &DeleteTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "DeleteTag")
	return
}

func NewDeleteTagResponse() (response *DeleteTagResponse) {
	response = &DeleteTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除 Tag
func (c *Client) DeleteTag(request *DeleteTagRequest) (response *DeleteTagResponse, err error) {
	if request == nil {
		request = NewDeleteTagRequest()
	}
	response = NewDeleteTagResponse()
	err = c.Send(request, response)
	return
}

func NewGetGCHistoryRequest() (request *GetGCHistoryRequest) {
	request = &GetGCHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "GetGCHistory")
	return
}

func NewGetGCHistoryResponse() (response *GetGCHistoryResponse) {
	response = &GetGCHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取仓库容量回收单个历史记录
func (c *Client) GetGCHistory(request *GetGCHistoryRequest) (response *GetGCHistoryResponse, err error) {
	if request == nil {
		request = NewGetGCHistoryRequest()
	}
	response = NewGetGCHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateRetentionRequest() (request *UpdateRetentionRequest) {
	request = &UpdateRetentionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "UpdateRetention")
	return
}

func NewUpdateRetentionResponse() (response *UpdateRetentionResponse) {
	response = &UpdateRetentionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新镜像清理策略
func (c *Client) UpdateRetention(request *UpdateRetentionRequest) (response *UpdateRetentionResponse, err error) {
	if request == nil {
		request = NewUpdateRetentionRequest()
	}
	response = NewUpdateRetentionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
	request = &DescribeImagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "DescribeImages")
	return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
	response = &DescribeImagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取某一项目下的镜像
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
	if request == nil {
		request = NewDescribeImagesRequest()
	}
	response = NewDescribeImagesResponse()
	err = c.Send(request, response)
	return
}

func NewSetGCStatusRequest() (request *SetGCStatusRequest) {
	request = &SetGCStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "SetGCStatus")
	return
}

func NewSetGCStatusResponse() (response *SetGCStatusResponse) {
	response = &SetGCStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置GC开关
func (c *Client) SetGCStatus(request *SetGCStatusRequest) (response *SetGCStatusResponse, err error) {
	if request == nil {
		request = NewSetGCStatusRequest()
	}
	response = NewSetGCStatusResponse()
	err = c.Send(request, response)
	return
}

func NewListImageBuildingTasksRequest() (request *ListImageBuildingTasksRequest) {
	request = &ListImageBuildingTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "ListImageBuildingTasks")
	return
}

func NewListImageBuildingTasksResponse() (response *ListImageBuildingTasksResponse) {
	response = &ListImageBuildingTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取镜像构建任务列表
func (c *Client) ListImageBuildingTasks(request *ListImageBuildingTasksRequest) (response *ListImageBuildingTasksResponse, err error) {
	if request == nil {
		request = NewListImageBuildingTasksRequest()
	}
	response = NewListImageBuildingTasksResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProjectRequest() (request *ModifyProjectRequest) {
	request = &ModifyProjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "ModifyProject")
	return
}

func NewModifyProjectResponse() (response *ModifyProjectResponse) {
	response = &ModifyProjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新仓库中项目的属性
func (c *Client) ModifyProject(request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
	if request == nil {
		request = NewModifyProjectRequest()
	}
	response = NewModifyProjectResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteImageRequest() (request *DeleteImageRequest) {
	request = &DeleteImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "DeleteImage")
	return
}

func NewDeleteImageResponse() (response *DeleteImageResponse) {
	response = &DeleteImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除镜像
func (c *Client) DeleteImage(request *DeleteImageRequest) (response *DeleteImageResponse, err error) {
	if request == nil {
		request = NewDeleteImageRequest()
	}
	response = NewDeleteImageResponse()
	err = c.Send(request, response)
	return
}

func NewCopyImagesRequest() (request *CopyImagesRequest) {
	request = &CopyImagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "CopyImages")
	return
}

func NewCopyImagesResponse() (response *CopyImagesResponse) {
	response = &CopyImagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从一个项目复制镜像到另一个项目
func (c *Client) CopyImages(request *CopyImagesRequest) (response *CopyImagesResponse, err error) {
	if request == nil {
		request = NewCopyImagesRequest()
	}
	response = NewCopyImagesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRetentionExecutionsRequest() (request *DescribeRetentionExecutionsRequest) {
	request = &DescribeRetentionExecutionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "DescribeRetentionExecutions")
	return
}

func NewDescribeRetentionExecutionsResponse() (response *DescribeRetentionExecutionsResponse) {
	response = &DescribeRetentionExecutionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取镜像清理记录
func (c *Client) DescribeRetentionExecutions(request *DescribeRetentionExecutionsRequest) (response *DescribeRetentionExecutionsResponse, err error) {
	if request == nil {
		request = NewDescribeRetentionExecutionsRequest()
	}
	response = NewDescribeRetentionExecutionsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateImageBuildingTaskRequest() (request *CreateImageBuildingTaskRequest) {
	request = &CreateImageBuildingTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "CreateImageBuildingTask")
	return
}

func NewCreateImageBuildingTaskResponse() (response *CreateImageBuildingTaskResponse) {
	response = &CreateImageBuildingTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建镜像构建任务
func (c *Client) CreateImageBuildingTask(request *CreateImageBuildingTaskRequest) (response *CreateImageBuildingTaskResponse, err error) {
	if request == nil {
		request = NewCreateImageBuildingTaskRequest()
	}
	response = NewCreateImageBuildingTaskResponse()
	err = c.Send(request, response)
	return
}

func NewGetRetentionExecutionLogRequest() (request *GetRetentionExecutionLogRequest) {
	request = &GetRetentionExecutionLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "GetRetentionExecutionLog")
	return
}

func NewGetRetentionExecutionLogResponse() (response *GetRetentionExecutionLogResponse) {
	response = &GetRetentionExecutionLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像清理历史记录日志信息
func (c *Client) GetRetentionExecutionLog(request *GetRetentionExecutionLogRequest) (response *GetRetentionExecutionLogResponse, err error) {
	if request == nil {
		request = NewGetRetentionExecutionLogRequest()
	}
	response = NewGetRetentionExecutionLogResponse()
	err = c.Send(request, response)
	return
}

func NewGetRetentionExecutionRequest() (request *GetRetentionExecutionRequest) {
	request = &GetRetentionExecutionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "GetRetentionExecution")
	return
}

func NewGetRetentionExecutionResponse() (response *GetRetentionExecutionResponse) {
	response = &GetRetentionExecutionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取单个镜像清理记录
func (c *Client) GetRetentionExecution(request *GetRetentionExecutionRequest) (response *GetRetentionExecutionResponse, err error) {
	if request == nil {
		request = NewGetRetentionExecutionRequest()
	}
	response = NewGetRetentionExecutionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProjectRequest() (request *CreateProjectRequest) {
	request = &CreateProjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "CreateProject")
	return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
	response = &CreateProjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在仓库中创建项目
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
	if request == nil {
		request = NewCreateProjectRequest()
	}
	response = NewCreateProjectResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteImageBuildingTaskRequest() (request *DeleteImageBuildingTaskRequest) {
	request = &DeleteImageBuildingTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "DeleteImageBuildingTask")
	return
}

func NewDeleteImageBuildingTaskResponse() (response *DeleteImageBuildingTaskResponse) {
	response = &DeleteImageBuildingTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除镜像构建任务
func (c *Client) DeleteImageBuildingTask(request *DeleteImageBuildingTaskRequest) (response *DeleteImageBuildingTaskResponse, err error) {
	if request == nil {
		request = NewDeleteImageBuildingTaskRequest()
	}
	response = NewDeleteImageBuildingTaskResponse()
	err = c.Send(request, response)
	return
}

func NewSearchImagesRequest() (request *SearchImagesRequest) {
	request = &SearchImagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "SearchImages")
	return
}

func NewSearchImagesResponse() (response *SearchImagesResponse) {
	response = &SearchImagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 全局查询镜像
func (c *Client) SearchImages(request *SearchImagesRequest) (response *SearchImagesResponse, err error) {
	if request == nil {
		request = NewSearchImagesRequest()
	}
	response = NewSearchImagesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageBuildingTaskRequest() (request *DescribeImageBuildingTaskRequest) {
	request = &DescribeImageBuildingTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "DescribeImageBuildingTask")
	return
}

func NewDescribeImageBuildingTaskResponse() (response *DescribeImageBuildingTaskResponse) {
	response = &DescribeImageBuildingTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像构建任务
func (c *Client) DescribeImageBuildingTask(request *DescribeImageBuildingTaskRequest) (response *DescribeImageBuildingTaskResponse, err error) {
	if request == nil {
		request = NewDescribeImageBuildingTaskRequest()
	}
	response = NewDescribeImageBuildingTaskResponse()
	err = c.Send(request, response)
	return
}

func NewScanImageRequest() (request *ScanImageRequest) {
	request = &ScanImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cargo", APIVersion, "ScanImage")
	return
}

func NewScanImageResponse() (response *ScanImageResponse) {
	response = &ScanImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 扫描指定镜像的安全漏洞
func (c *Client) ScanImage(request *ScanImageRequest) (response *ScanImageResponse, err error) {
	if request == nil {
		request = NewScanImageRequest()
	}
	response = NewScanImageResponse()
	err = c.Send(request, response)
	return
}
