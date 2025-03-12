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

package v20210830

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2021-08-30"

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

func NewDeleteRemoteUpdateTagRequest() (request *DeleteRemoteUpdateTagRequest) {
	request = &DeleteRemoteUpdateTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteRemoteUpdateTag")
	return
}

func NewDeleteRemoteUpdateTagResponse() (response *DeleteRemoteUpdateTagResponse) {
	response = &DeleteRemoteUpdateTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除远程更新标签
func (c *Client) DeleteRemoteUpdateTag(request *DeleteRemoteUpdateTagRequest) (response *DeleteRemoteUpdateTagResponse, err error) {
	if request == nil {
		request = NewDeleteRemoteUpdateTagRequest()
	}
	response = NewDeleteRemoteUpdateTagResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemoteUpdateUserListRequest() (request *DescribeRemoteUpdateUserListRequest) {
	request = &DescribeRemoteUpdateUserListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeRemoteUpdateUserList")
	return
}

func NewDescribeRemoteUpdateUserListResponse() (response *DescribeRemoteUpdateUserListResponse) {
	response = &DescribeRemoteUpdateUserListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取远程更新用户列表
func (c *Client) DescribeRemoteUpdateUserList(request *DescribeRemoteUpdateUserListRequest) (response *DescribeRemoteUpdateUserListResponse, err error) {
	if request == nil {
		request = NewDescribeRemoteUpdateUserListRequest()
	}
	response = NewDescribeRemoteUpdateUserListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemoteUpdatePublishLogsRequest() (request *DescribeRemoteUpdatePublishLogsRequest) {
	request = &DescribeRemoteUpdatePublishLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeRemoteUpdatePublishLogs")
	return
}

func NewDescribeRemoteUpdatePublishLogsResponse() (response *DescribeRemoteUpdatePublishLogsResponse) {
	response = &DescribeRemoteUpdatePublishLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取远程更新发布历史记录
func (c *Client) DescribeRemoteUpdatePublishLogs(request *DescribeRemoteUpdatePublishLogsRequest) (response *DescribeRemoteUpdatePublishLogsResponse, err error) {
	if request == nil {
		request = NewDescribeRemoteUpdatePublishLogsRequest()
	}
	response = NewDescribeRemoteUpdatePublishLogsResponse()
	err = c.Send(request, response)
	return
}

func NewExportVulEventsRequest() (request *ExportVulEventsRequest) {
	request = &ExportVulEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportVulEvents")
	return
}

func NewExportVulEventsResponse() (response *ExportVulEventsResponse) {
	response = &ExportVulEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞&基线管理 - 安全漏洞 -列表 -导出
func (c *Client) ExportVulEvents(request *ExportVulEventsRequest) (response *ExportVulEventsResponse, err error) {
	if request == nil {
		request = NewExportVulEventsRequest()
	}
	response = NewExportVulEventsResponse()
	err = c.Send(request, response)
	return
}

func NewExportBashEventsRequest() (request *ExportBashEventsRequest) {
	request = &ExportBashEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportBashEvents")
	return
}

func NewExportBashEventsResponse() (response *ExportBashEventsResponse) {
	response = &ExportBashEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出高危命令列表
func (c *Client) ExportBashEvents(request *ExportBashEventsRequest) (response *ExportBashEventsResponse, err error) {
	if request == nil {
		request = NewExportBashEventsRequest()
	}
	response = NewExportBashEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentPushTaskResultRequest() (request *DescribeAgentPushTaskResultRequest) {
	request = &DescribeAgentPushTaskResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAgentPushTaskResult")
	return
}

func NewDescribeAgentPushTaskResultResponse() (response *DescribeAgentPushTaskResultResponse) {
	response = &DescribeAgentPushTaskResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 命令推送 - 列表 - 结果
func (c *Client) DescribeAgentPushTaskResult(request *DescribeAgentPushTaskResultRequest) (response *DescribeAgentPushTaskResultResponse, err error) {
	if request == nil {
		request = NewDescribeAgentPushTaskResultRequest()
	}
	response = NewDescribeAgentPushTaskResultResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetMachineListRequest() (request *DescribeAssetMachineListRequest) {
	request = &DescribeAssetMachineListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetMachineList")
	return
}

func NewDescribeAssetMachineListResponse() (response *DescribeAssetMachineListResponse) {
	response = &DescribeAssetMachineListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源监控列表
func (c *Client) DescribeAssetMachineList(request *DescribeAssetMachineListRequest) (response *DescribeAssetMachineListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetMachineListRequest()
	}
	response = NewDescribeAssetMachineListResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetProcessInfoListRequest() (request *ExportAssetProcessInfoListRequest) {
	request = &ExportAssetProcessInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAssetProcessInfoList")
	return
}

func NewExportAssetProcessInfoListResponse() (response *ExportAssetProcessInfoListResponse) {
	response = &ExportAssetProcessInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理进程列表
func (c *Client) ExportAssetProcessInfoList(request *ExportAssetProcessInfoListRequest) (response *ExportAssetProcessInfoListResponse, err error) {
	if request == nil {
		request = NewExportAssetProcessInfoListRequest()
	}
	response = NewExportAssetProcessInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebLocationPathListRequest() (request *DescribeAssetWebLocationPathListRequest) {
	request = &DescribeAssetWebLocationPathListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetWebLocationPathList")
	return
}

func NewDescribeAssetWebLocationPathListResponse() (response *DescribeAssetWebLocationPathListResponse) {
	response = &DescribeAssetWebLocationPathListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Web站点虚拟目录列表
func (c *Client) DescribeAssetWebLocationPathList(request *DescribeAssetWebLocationPathListRequest) (response *DescribeAssetWebLocationPathListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebLocationPathListRequest()
	}
	response = NewDescribeAssetWebLocationPathListResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadPushResultRequest() (request *DownloadPushResultRequest) {
	request = &DownloadPushResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DownloadPushResult")
	return
}

func NewDownloadPushResultResponse() (response *DownloadPushResultResponse) {
	response = &DownloadPushResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载推送结果
func (c *Client) DownloadPushResult(request *DownloadPushResultRequest) (response *DownloadPushResultResponse, err error) {
	if request == nil {
		request = NewDownloadPushResultRequest()
	}
	response = NewDownloadPushResultResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteInstallPackageRequest() (request *DeleteInstallPackageRequest) {
	request = &DeleteInstallPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteInstallPackage")
	return
}

func NewDeleteInstallPackageResponse() (response *DeleteInstallPackageResponse) {
	response = &DeleteInstallPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除安装包记录
func (c *Client) DeleteInstallPackage(request *DeleteInstallPackageRequest) (response *DeleteInstallPackageResponse, err error) {
	if request == nil {
		request = NewDeleteInstallPackageRequest()
	}
	response = NewDeleteInstallPackageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJavaMemShellInfoRequest() (request *DescribeJavaMemShellInfoRequest) {
	request = &DescribeJavaMemShellInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeJavaMemShellInfo")
	return
}

func NewDescribeJavaMemShellInfoResponse() (response *DescribeJavaMemShellInfoResponse) {
	response = &DescribeJavaMemShellInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取内存马详情
func (c *Client) DescribeJavaMemShellInfo(request *DescribeJavaMemShellInfoRequest) (response *DescribeJavaMemShellInfoResponse, err error) {
	if request == nil {
		request = NewDescribeJavaMemShellInfoRequest()
	}
	response = NewDescribeJavaMemShellInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVersionStatisticsRequest() (request *DescribeVersionStatisticsRequest) {
	request = &DescribeVersionStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeVersionStatistics")
	return
}

func NewDescribeVersionStatisticsResponse() (response *DescribeVersionStatisticsResponse) {
	response = &DescribeVersionStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeVersionStatistics) 用于统计专业版和基础版机器数
func (c *Client) DescribeVersionStatistics(request *DescribeVersionStatisticsRequest) (response *DescribeVersionStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeVersionStatisticsRequest()
	}
	response = NewDescribeVersionStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewExportAgentsListRequest() (request *ExportAgentsListRequest) {
	request = &ExportAgentsListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAgentsList")
	return
}

func NewExportAgentsListResponse() (response *ExportAgentsListResponse) {
	response = &ExportAgentsListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 主机管理 - ip to uuid
func (c *Client) ExportAgentsList(request *ExportAgentsListRequest) (response *ExportAgentsListResponse, err error) {
	if request == nil {
		request = NewExportAgentsListRequest()
	}
	response = NewExportAgentsListResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetEnvListRequest() (request *ExportAssetEnvListRequest) {
	request = &ExportAssetEnvListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAssetEnvList")
	return
}

func NewExportAssetEnvListResponse() (response *ExportAssetEnvListResponse) {
	response = &ExportAssetEnvListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理环境变量列表
func (c *Client) ExportAssetEnvList(request *ExportAssetEnvListRequest) (response *ExportAssetEnvListResponse, err error) {
	if request == nil {
		request = NewExportAssetEnvListRequest()
	}
	response = NewExportAssetEnvListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentDetailRequest() (request *DescribeAgentDetailRequest) {
	request = &DescribeAgentDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAgentDetail")
	return
}

func NewDescribeAgentDetailResponse() (response *DescribeAgentDetailResponse) {
	response = &DescribeAgentDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主机基本信息
func (c *Client) DescribeAgentDetail(request *DescribeAgentDetailRequest) (response *DescribeAgentDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAgentDetailRequest()
	}
	response = NewDescribeAgentDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebAppPluginListRequest() (request *DescribeAssetWebAppPluginListRequest) {
	request = &DescribeAssetWebAppPluginListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetWebAppPluginList")
	return
}

func NewDescribeAssetWebAppPluginListResponse() (response *DescribeAssetWebAppPluginListResponse) {
	response = &DescribeAssetWebAppPluginListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Web应用插件列表
func (c *Client) DescribeAssetWebAppPluginList(request *DescribeAssetWebAppPluginListRequest) (response *DescribeAssetWebAppPluginListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebAppPluginListRequest()
	}
	response = NewDescribeAssetWebAppPluginListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebLocationListRequest() (request *DescribeAssetWebLocationListRequest) {
	request = &DescribeAssetWebLocationListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetWebLocationList")
	return
}

func NewDescribeAssetWebLocationListResponse() (response *DescribeAssetWebLocationListResponse) {
	response = &DescribeAssetWebLocationListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Web站点列表
func (c *Client) DescribeAssetWebLocationList(request *DescribeAssetWebLocationListRequest) (response *DescribeAssetWebLocationListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebLocationListRequest()
	}
	response = NewDescribeAssetWebLocationListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemoteUpdateDownloadTokenRequest() (request *DescribeRemoteUpdateDownloadTokenRequest) {
	request = &DescribeRemoteUpdateDownloadTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeRemoteUpdateDownloadToken")
	return
}

func NewDescribeRemoteUpdateDownloadTokenResponse() (response *DescribeRemoteUpdateDownloadTokenResponse) {
	response = &DescribeRemoteUpdateDownloadTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取下载远程更新文件Token
func (c *Client) DescribeRemoteUpdateDownloadToken(request *DescribeRemoteUpdateDownloadTokenRequest) (response *DescribeRemoteUpdateDownloadTokenResponse, err error) {
	if request == nil {
		request = NewDescribeRemoteUpdateDownloadTokenRequest()
	}
	response = NewDescribeRemoteUpdateDownloadTokenResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMalwareEvilRequest() (request *ModifyMalwareEvilRequest) {
	request = &ModifyMalwareEvilRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyMalwareEvil")
	return
}

func NewModifyMalwareEvilResponse() (response *ModifyMalwareEvilResponse) {
	response = &ModifyMalwareEvilResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改恶意文件成功确认
func (c *Client) ModifyMalwareEvil(request *ModifyMalwareEvilRequest) (response *ModifyMalwareEvilResponse, err error) {
	if request == nil {
		request = NewModifyMalwareEvilRequest()
	}
	response = NewModifyMalwareEvilResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetWebServiceInfoListRequest() (request *ExportAssetWebServiceInfoListRequest) {
	request = &ExportAssetWebServiceInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAssetWebServiceInfoList")
	return
}

func NewExportAssetWebServiceInfoListResponse() (response *ExportAssetWebServiceInfoListResponse) {
	response = &ExportAssetWebServiceInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理Web服务列表
func (c *Client) ExportAssetWebServiceInfoList(request *ExportAssetWebServiceInfoListRequest) (response *ExportAssetWebServiceInfoListResponse, err error) {
	if request == nil {
		request = NewExportAssetWebServiceInfoListRequest()
	}
	response = NewExportAssetWebServiceInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProLicenseCountRequest() (request *ModifyProLicenseCountRequest) {
	request = &ModifyProLicenseCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyProLicenseCount")
	return
}

func NewModifyProLicenseCountResponse() (response *ModifyProLicenseCountResponse) {
	response = &ModifyProLicenseCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新授权
func (c *Client) ModifyProLicenseCount(request *ModifyProLicenseCountRequest) (response *ModifyProLicenseCountResponse, err error) {
	if request == nil {
		request = NewModifyProLicenseCountRequest()
	}
	response = NewModifyProLicenseCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJavaMemShellPluginListRequest() (request *DescribeJavaMemShellPluginListRequest) {
	request = &DescribeJavaMemShellPluginListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeJavaMemShellPluginList")
	return
}

func NewDescribeJavaMemShellPluginListResponse() (response *DescribeJavaMemShellPluginListResponse) {
	response = &DescribeJavaMemShellPluginListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询License列表
func (c *Client) DescribeJavaMemShellPluginList(request *DescribeJavaMemShellPluginListRequest) (response *DescribeJavaMemShellPluginListResponse, err error) {
	if request == nil {
		request = NewDescribeJavaMemShellPluginListRequest()
	}
	response = NewDescribeJavaMemShellPluginListResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetPortInfoListRequest() (request *ExportAssetPortInfoListRequest) {
	request = &ExportAssetPortInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAssetPortInfoList")
	return
}

func NewExportAssetPortInfoListResponse() (response *ExportAssetPortInfoListResponse) {
	response = &ExportAssetPortInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理端口列表
func (c *Client) ExportAssetPortInfoList(request *ExportAssetPortInfoListRequest) (response *ExportAssetPortInfoListResponse, err error) {
	if request == nil {
		request = NewExportAssetPortInfoListRequest()
	}
	response = NewExportAssetPortInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBashEventsStatusRequest() (request *ModifyBashEventsStatusRequest) {
	request = &ModifyBashEventsStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyBashEventsStatus")
	return
}

func NewModifyBashEventsStatusResponse() (response *ModifyBashEventsStatusResponse) {
	response = &ModifyBashEventsStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 高危命令 - 标记
func (c *Client) ModifyBashEventsStatus(request *ModifyBashEventsStatusRequest) (response *ModifyBashEventsStatusResponse, err error) {
	if request == nil {
		request = NewModifyBashEventsStatusRequest()
	}
	response = NewModifyBashEventsStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTrojanSettingRequest() (request *ModifyTrojanSettingRequest) {
	request = &ModifyTrojanSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyTrojanSetting")
	return
}

func NewModifyTrojanSettingResponse() (response *ModifyTrojanSettingResponse) {
	response = &ModifyTrojanSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存木马隔离
func (c *Client) ModifyTrojanSetting(request *ModifyTrojanSettingRequest) (response *ModifyTrojanSettingResponse, err error) {
	if request == nil {
		request = NewModifyTrojanSettingRequest()
	}
	response = NewModifyTrojanSettingResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetUserListRequest() (request *ExportAssetUserListRequest) {
	request = &ExportAssetUserListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAssetUserList")
	return
}

func NewExportAssetUserListResponse() (response *ExportAssetUserListResponse) {
	response = &ExportAssetUserListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出账号列表
func (c *Client) ExportAssetUserList(request *ExportAssetUserListRequest) (response *ExportAssetUserListResponse, err error) {
	if request == nil {
		request = NewExportAssetUserListRequest()
	}
	response = NewExportAssetUserListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetPlanTaskListRequest() (request *DescribeAssetPlanTaskListRequest) {
	request = &DescribeAssetPlanTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetPlanTaskList")
	return
}

func NewDescribeAssetPlanTaskListResponse() (response *DescribeAssetPlanTaskListResponse) {
	response = &DescribeAssetPlanTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取计划任务列表
func (c *Client) DescribeAssetPlanTaskList(request *DescribeAssetPlanTaskListRequest) (response *DescribeAssetPlanTaskListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetPlanTaskListRequest()
	}
	response = NewDescribeAssetPlanTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseHistoryReleasedRequest() (request *DescribeLicenseHistoryReleasedRequest) {
	request = &DescribeLicenseHistoryReleasedRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeLicenseHistoryReleased")
	return
}

func NewDescribeLicenseHistoryReleasedResponse() (response *DescribeLicenseHistoryReleasedResponse) {
	response = &DescribeLicenseHistoryReleasedResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询License历史签发记录
func (c *Client) DescribeLicenseHistoryReleased(request *DescribeLicenseHistoryReleasedRequest) (response *DescribeLicenseHistoryReleasedResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseHistoryReleasedRequest()
	}
	response = NewDescribeLicenseHistoryReleasedResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulTypeListRequest() (request *DescribeVulTypeListRequest) {
	request = &DescribeVulTypeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeVulTypeList")
	return
}

func NewDescribeVulTypeListResponse() (response *DescribeVulTypeListResponse) {
	response = &DescribeVulTypeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞类型列表
func (c *Client) DescribeVulTypeList(request *DescribeVulTypeListRequest) (response *DescribeVulTypeListResponse, err error) {
	if request == nil {
		request = NewDescribeVulTypeListRequest()
	}
	response = NewDescribeVulTypeListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulScanConfigListRequest() (request *DescribeVulScanConfigListRequest) {
	request = &DescribeVulScanConfigListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeVulScanConfigList")
	return
}

func NewDescribeVulScanConfigListResponse() (response *DescribeVulScanConfigListResponse) {
	response = &DescribeVulScanConfigListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置文件管理列表
func (c *Client) DescribeVulScanConfigList(request *DescribeVulScanConfigListRequest) (response *DescribeVulScanConfigListResponse, err error) {
	if request == nil {
		request = NewDescribeVulScanConfigListRequest()
	}
	response = NewDescribeVulScanConfigListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetAppListRequest() (request *DescribeAssetAppListRequest) {
	request = &DescribeAssetAppListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetAppList")
	return
}

func NewDescribeAssetAppListResponse() (response *DescribeAssetAppListResponse) {
	response = &DescribeAssetAppListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用列表
func (c *Client) DescribeAssetAppList(request *DescribeAssetAppListRequest) (response *DescribeAssetAppListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetAppListRequest()
	}
	response = NewDescribeAssetAppListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetSystemPackageListRequest() (request *DescribeAssetSystemPackageListRequest) {
	request = &DescribeAssetSystemPackageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetSystemPackageList")
	return
}

func NewDescribeAssetSystemPackageListResponse() (response *DescribeAssetSystemPackageListResponse) {
	response = &DescribeAssetSystemPackageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取系统安装包列表
func (c *Client) DescribeAssetSystemPackageList(request *DescribeAssetSystemPackageListRequest) (response *DescribeAssetSystemPackageListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetSystemPackageListRequest()
	}
	response = NewDescribeAssetSystemPackageListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineItemInfoRequest() (request *DescribeBaselineItemInfoRequest) {
	request = &DescribeBaselineItemInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeBaselineItemInfo")
	return
}

func NewDescribeBaselineItemInfoResponse() (response *DescribeBaselineItemInfoResponse) {
	response = &DescribeBaselineItemInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线检测项信息
func (c *Client) DescribeBaselineItemInfo(request *DescribeBaselineItemInfoRequest) (response *DescribeBaselineItemInfoResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineItemInfoRequest()
	}
	response = NewDescribeBaselineItemInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSocTamperProofDetailRequest() (request *DescribeSocTamperProofDetailRequest) {
	request = &DescribeSocTamperProofDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeSocTamperProofDetail")
	return
}

func NewDescribeSocTamperProofDetailResponse() (response *DescribeSocTamperProofDetailResponse) {
	response = &DescribeSocTamperProofDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 网页防篡改详情
func (c *Client) DescribeSocTamperProofDetail(request *DescribeSocTamperProofDetailRequest) (response *DescribeSocTamperProofDetailResponse, err error) {
	if request == nil {
		request = NewDescribeSocTamperProofDetailRequest()
	}
	response = NewDescribeSocTamperProofDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDefenceEventDetailRequest() (request *DescribeDefenceEventDetailRequest) {
	request = &DescribeDefenceEventDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeDefenceEventDetail")
	return
}

func NewDescribeDefenceEventDetailResponse() (response *DescribeDefenceEventDetailResponse) {
	response = &DescribeDefenceEventDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取漏洞防御事件详情
func (c *Client) DescribeDefenceEventDetail(request *DescribeDefenceEventDetailRequest) (response *DescribeDefenceEventDetailResponse, err error) {
	if request == nil {
		request = NewDescribeDefenceEventDetailRequest()
	}
	response = NewDescribeDefenceEventDetailResponse()
	err = c.Send(request, response)
	return
}

func NewPublishInstallPackageRequest() (request *PublishInstallPackageRequest) {
	request = &PublishInstallPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "PublishInstallPackage")
	return
}

func NewPublishInstallPackageResponse() (response *PublishInstallPackageResponse) {
	response = &PublishInstallPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 版本管理 - 安装包管理 - 发布安装包
func (c *Client) PublishInstallPackage(request *PublishInstallPackageRequest) (response *PublishInstallPackageResponse, err error) {
	if request == nil {
		request = NewPublishInstallPackageRequest()
	}
	response = NewPublishInstallPackageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyFlagshipLicenseRequest() (request *ModifyFlagshipLicenseRequest) {
	request = &ModifyFlagshipLicenseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyFlagshipLicense")
	return
}

func NewModifyFlagshipLicenseResponse() (response *ModifyFlagshipLicenseResponse) {
	response = &ModifyFlagshipLicenseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新旗舰版License/升级旗舰版License
func (c *Client) ModifyFlagshipLicense(request *ModifyFlagshipLicenseRequest) (response *ModifyFlagshipLicenseResponse, err error) {
	if request == nil {
		request = NewModifyFlagshipLicenseRequest()
	}
	response = NewModifyFlagshipLicenseResponse()
	err = c.Send(request, response)
	return
}

func NewExportHostTrendRequest() (request *ExportHostTrendRequest) {
	request = &ExportHostTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportHostTrend")
	return
}

func NewExportHostTrendResponse() (response *ExportHostTrendResponse) {
	response = &ExportHostTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出服务器在线趋势
func (c *Client) ExportHostTrend(request *ExportHostTrendRequest) (response *ExportHostTrendResponse, err error) {
	if request == nil {
		request = NewExportHostTrendRequest()
	}
	response = NewExportHostTrendResponse()
	err = c.Send(request, response)
	return
}

func NewExportTaskComponentsRequest() (request *ExportTaskComponentsRequest) {
	request = &ExportTaskComponentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportTaskComponents")
	return
}

func NewExportTaskComponentsResponse() (response *ExportTaskComponentsResponse) {
	response = &ExportTaskComponentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出组件管理列表
func (c *Client) ExportTaskComponents(request *ExportTaskComponentsRequest) (response *ExportTaskComponentsResponse, err error) {
	if request == nil {
		request = NewExportTaskComponentsRequest()
	}
	response = NewExportTaskComponentsResponse()
	err = c.Send(request, response)
	return
}

func NewExportRemoteUpdateConfigRequest() (request *ExportRemoteUpdateConfigRequest) {
	request = &ExportRemoteUpdateConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportRemoteUpdateConfig")
	return
}

func NewExportRemoteUpdateConfigResponse() (response *ExportRemoteUpdateConfigResponse) {
	response = &ExportRemoteUpdateConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出远程更新升级配置
func (c *Client) ExportRemoteUpdateConfig(request *ExportRemoteUpdateConfigRequest) (response *ExportRemoteUpdateConfigResponse, err error) {
	if request == nil {
		request = NewExportRemoteUpdateConfigRequest()
	}
	response = NewExportRemoteUpdateConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVulEventsStatusRequest() (request *ModifyVulEventsStatusRequest) {
	request = &ModifyVulEventsStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyVulEventsStatus")
	return
}

func NewModifyVulEventsStatusResponse() (response *ModifyVulEventsStatusResponse) {
	response = &ModifyVulEventsStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全漏洞 - 操作 - 忽略/重新检测/批量确认漏洞
func (c *Client) ModifyVulEventsStatus(request *ModifyVulEventsStatusRequest) (response *ModifyVulEventsStatusResponse, err error) {
	if request == nil {
		request = NewModifyVulEventsStatusRequest()
	}
	response = NewModifyVulEventsStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstallPackagePublishListRequest() (request *DescribeInstallPackagePublishListRequest) {
	request = &DescribeInstallPackagePublishListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeInstallPackagePublishList")
	return
}

func NewDescribeInstallPackagePublishListResponse() (response *DescribeInstallPackagePublishListResponse) {
	response = &DescribeInstallPackagePublishListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 版本管理 - 安装包管理 - 获取发布安装包列表
func (c *Client) DescribeInstallPackagePublishList(request *DescribeInstallPackagePublishListRequest) (response *DescribeInstallPackagePublishListResponse, err error) {
	if request == nil {
		request = NewDescribeInstallPackagePublishListRequest()
	}
	response = NewDescribeInstallPackagePublishListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyZKConfigInfoRequest() (request *ModifyZKConfigInfoRequest) {
	request = &ModifyZKConfigInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyZKConfigInfo")
	return
}

func NewModifyZKConfigInfoResponse() (response *ModifyZKConfigInfoResponse) {
	response = &ModifyZKConfigInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 系统管理 - Zookeeper配置 - 新建/修改配置
func (c *Client) ModifyZKConfigInfo(request *ModifyZKConfigInfoRequest) (response *ModifyZKConfigInfoResponse, err error) {
	if request == nil {
		request = NewModifyZKConfigInfoRequest()
	}
	response = NewModifyZKConfigInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteZKGroupInfoRequest() (request *DeleteZKGroupInfoRequest) {
	request = &DeleteZKGroupInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteZKGroupInfo")
	return
}

func NewDeleteZKGroupInfoResponse() (response *DeleteZKGroupInfoResponse) {
	response = &DeleteZKGroupInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 系统管理 - Zookeeper配置 - 删除分组
func (c *Client) DeleteZKGroupInfo(request *DeleteZKGroupInfoRequest) (response *DeleteZKGroupInfoResponse, err error) {
	if request == nil {
		request = NewDeleteZKGroupInfoRequest()
	}
	response = NewDeleteZKGroupInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpecimenFileRequest() (request *DescribeSpecimenFileRequest) {
	request = &DescribeSpecimenFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeSpecimenFile")
	return
}

func NewDescribeSpecimenFileResponse() (response *DescribeSpecimenFileResponse) {
	response = &DescribeSpecimenFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 文件查杀 - 列表 - 样本详情
func (c *Client) DescribeSpecimenFile(request *DescribeSpecimenFileRequest) (response *DescribeSpecimenFileResponse, err error) {
	if request == nil {
		request = NewDescribeSpecimenFileRequest()
	}
	response = NewDescribeSpecimenFileResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineRuleCategoryListRequest() (request *DescribeBaselineRuleCategoryListRequest) {
	request = &DescribeBaselineRuleCategoryListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeBaselineRuleCategoryList")
	return
}

func NewDescribeBaselineRuleCategoryListResponse() (response *DescribeBaselineRuleCategoryListResponse) {
	response = &DescribeBaselineRuleCategoryListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线分类列表
func (c *Client) DescribeBaselineRuleCategoryList(request *DescribeBaselineRuleCategoryListRequest) (response *DescribeBaselineRuleCategoryListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineRuleCategoryListRequest()
	}
	response = NewDescribeBaselineRuleCategoryListResponse()
	err = c.Send(request, response)
	return
}

func NewExportJavaMemShellPluginsRequest() (request *ExportJavaMemShellPluginsRequest) {
	request = &ExportJavaMemShellPluginsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportJavaMemShellPlugins")
	return
}

func NewExportJavaMemShellPluginsResponse() (response *ExportJavaMemShellPluginsResponse) {
	response = &ExportJavaMemShellPluginsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 内存马列表 - 列表 - 导出
func (c *Client) ExportJavaMemShellPlugins(request *ExportJavaMemShellPluginsRequest) (response *ExportJavaMemShellPluginsResponse, err error) {
	if request == nil {
		request = NewExportJavaMemShellPluginsRequest()
	}
	response = NewExportJavaMemShellPluginsResponse()
	err = c.Send(request, response)
	return
}

func NewImportKnowledgeRequest() (request *ImportKnowledgeRequest) {
	request = &ImportKnowledgeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ImportKnowledge")
	return
}

func NewImportKnowledgeResponse() (response *ImportKnowledgeResponse) {
	response = &ImportKnowledgeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 巡检任务-知识库管理-导入知识库
func (c *Client) ImportKnowledge(request *ImportKnowledgeRequest) (response *ImportKnowledgeResponse, err error) {
	if request == nil {
		request = NewImportKnowledgeRequest()
	}
	response = NewImportKnowledgeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentsListRequest() (request *DescribeAgentsListRequest) {
	request = &DescribeAgentsListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAgentsList")
	return
}

func NewDescribeAgentsListResponse() (response *DescribeAgentsListResponse) {
	response = &DescribeAgentsListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 主机管理 - 列表
func (c *Client) DescribeAgentsList(request *DescribeAgentsListRequest) (response *DescribeAgentsListResponse, err error) {
	if request == nil {
		request = NewDescribeAgentsListRequest()
	}
	response = NewDescribeAgentsListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulPackageListRequest() (request *DescribeVulPackageListRequest) {
	request = &DescribeVulPackageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeVulPackageList")
	return
}

func NewDescribeVulPackageListResponse() (response *DescribeVulPackageListResponse) {
	response = &DescribeVulPackageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 任务包管理列表
func (c *Client) DescribeVulPackageList(request *DescribeVulPackageListRequest) (response *DescribeVulPackageListResponse, err error) {
	if request == nil {
		request = NewDescribeVulPackageListRequest()
	}
	response = NewDescribeVulPackageListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetProcessCountRequest() (request *DescribeAssetProcessCountRequest) {
	request = &DescribeAssetProcessCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetProcessCount")
	return
}

func NewDescribeAssetProcessCountResponse() (response *DescribeAssetProcessCountResponse) {
	response = &DescribeAssetProcessCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有进程数量
func (c *Client) DescribeAssetProcessCount(request *DescribeAssetProcessCountRequest) (response *DescribeAssetProcessCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetProcessCountRequest()
	}
	response = NewDescribeAssetProcessCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBashEventsRequest() (request *DescribeBashEventsRequest) {
	request = &DescribeBashEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeBashEvents")
	return
}

func NewDescribeBashEventsResponse() (response *DescribeBashEventsResponse) {
	response = &DescribeBashEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 高危命令 - 列表
func (c *Client) DescribeBashEvents(request *DescribeBashEventsRequest) (response *DescribeBashEventsResponse, err error) {
	if request == nil {
		request = NewDescribeBashEventsRequest()
	}
	response = NewDescribeBashEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBruteForceEventsRequest() (request *DescribeBruteForceEventsRequest) {
	request = &DescribeBruteForceEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeBruteForceEvents")
	return
}

func NewDescribeBruteForceEventsResponse() (response *DescribeBruteForceEventsResponse) {
	response = &DescribeBruteForceEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 密码破解 - 列表
func (c *Client) DescribeBruteForceEvents(request *DescribeBruteForceEventsRequest) (response *DescribeBruteForceEventsResponse, err error) {
	if request == nil {
		request = NewDescribeBruteForceEventsRequest()
	}
	response = NewDescribeBruteForceEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUpdateGrayRequest() (request *DeleteUpdateGrayRequest) {
	request = &DeleteUpdateGrayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteUpdateGray")
	return
}

func NewDeleteUpdateGrayResponse() (response *DeleteUpdateGrayResponse) {
	response = &DeleteUpdateGrayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 删除版本任务
func (c *Client) DeleteUpdateGray(request *DeleteUpdateGrayRequest) (response *DeleteUpdateGrayResponse, err error) {
	if request == nil {
		request = NewDeleteUpdateGrayRequest()
	}
	response = NewDeleteUpdateGrayResponse()
	err = c.Send(request, response)
	return
}

func NewExportFileSpecimenListRequest() (request *ExportFileSpecimenListRequest) {
	request = &ExportFileSpecimenListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportFileSpecimenList")
	return
}

func NewExportFileSpecimenListResponse() (response *ExportFileSpecimenListResponse) {
	response = &ExportFileSpecimenListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全策略 - 样本管理 - 列表 - 导出
func (c *Client) ExportFileSpecimenList(request *ExportFileSpecimenListRequest) (response *ExportFileSpecimenListResponse, err error) {
	if request == nil {
		request = NewExportFileSpecimenListRequest()
	}
	response = NewExportFileSpecimenListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulPackageFullListRequest() (request *DescribeVulPackageFullListRequest) {
	request = &DescribeVulPackageFullListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeVulPackageFullList")
	return
}

func NewDescribeVulPackageFullListResponse() (response *DescribeVulPackageFullListResponse) {
	response = &DescribeVulPackageFullListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 任务管理 - 列表
func (c *Client) DescribeVulPackageFullList(request *DescribeVulPackageFullListRequest) (response *DescribeVulPackageFullListResponse, err error) {
	if request == nil {
		request = NewDescribeVulPackageFullListRequest()
	}
	response = NewDescribeVulPackageFullListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetCoreModuleListRequest() (request *DescribeAssetCoreModuleListRequest) {
	request = &DescribeAssetCoreModuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetCoreModuleList")
	return
}

func NewDescribeAssetCoreModuleListResponse() (response *DescribeAssetCoreModuleListResponse) {
	response = &DescribeAssetCoreModuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取内核模块列表
func (c *Client) DescribeAssetCoreModuleList(request *DescribeAssetCoreModuleListRequest) (response *DescribeAssetCoreModuleListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetCoreModuleListRequest()
	}
	response = NewDescribeAssetCoreModuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefencePluginListRequest() (request *DescribeVulDefencePluginListRequest) {
	request = &DescribeVulDefencePluginListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeVulDefencePluginList")
	return
}

func NewDescribeVulDefencePluginListResponse() (response *DescribeVulDefencePluginListResponse) {
	response = &DescribeVulDefencePluginListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 插件详情列表
func (c *Client) DescribeVulDefencePluginList(request *DescribeVulDefencePluginListRequest) (response *DescribeVulDefencePluginListResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefencePluginListRequest()
	}
	response = NewDescribeVulDefencePluginListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyZKGroupInfoRequest() (request *ModifyZKGroupInfoRequest) {
	request = &ModifyZKGroupInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyZKGroupInfo")
	return
}

func NewModifyZKGroupInfoResponse() (response *ModifyZKGroupInfoResponse) {
	response = &ModifyZKGroupInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑Zookeeper分组配置
func (c *Client) ModifyZKGroupInfo(request *ModifyZKGroupInfoRequest) (response *ModifyZKGroupInfoResponse, err error) {
	if request == nil {
		request = NewModifyZKGroupInfoRequest()
	}
	response = NewModifyZKGroupInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVulScanTaskNewRequest() (request *CreateVulScanTaskNewRequest) {
	request = &CreateVulScanTaskNewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "CreateVulScanTaskNew")
	return
}

func NewCreateVulScanTaskNewResponse() (response *CreateVulScanTaskNewResponse) {
	response = &CreateVulScanTaskNewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 任务管理 - 添加任务-同380一致
func (c *Client) CreateVulScanTaskNew(request *CreateVulScanTaskNewRequest) (response *CreateVulScanTaskNewResponse, err error) {
	if request == nil {
		request = NewCreateVulScanTaskNewRequest()
	}
	response = NewCreateVulScanTaskNewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefenceCountRequest() (request *DescribeVulDefenceCountRequest) {
	request = &DescribeVulDefenceCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeVulDefenceCount")
	return
}

func NewDescribeVulDefenceCountResponse() (response *DescribeVulDefenceCountResponse) {
	response = &DescribeVulDefenceCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取漏洞防御主机总数
func (c *Client) DescribeVulDefenceCount(request *DescribeVulDefenceCountRequest) (response *DescribeVulDefenceCountResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefenceCountRequest()
	}
	response = NewDescribeVulDefenceCountResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetWebAppListRequest() (request *ExportAssetWebAppListRequest) {
	request = &ExportAssetWebAppListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAssetWebAppList")
	return
}

func NewExportAssetWebAppListResponse() (response *ExportAssetWebAppListResponse) {
	response = &ExportAssetWebAppListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理Web应用列表
func (c *Client) ExportAssetWebAppList(request *ExportAssetWebAppListRequest) (response *ExportAssetWebAppListResponse, err error) {
	if request == nil {
		request = NewExportAssetWebAppListRequest()
	}
	response = NewExportAssetWebAppListResponse()
	err = c.Send(request, response)
	return
}

func NewExportKnowledgeRequest() (request *ExportKnowledgeRequest) {
	request = &ExportKnowledgeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportKnowledge")
	return
}

func NewExportKnowledgeResponse() (response *ExportKnowledgeResponse) {
	response = &ExportKnowledgeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出知识库
func (c *Client) ExportKnowledge(request *ExportKnowledgeRequest) (response *ExportKnowledgeResponse, err error) {
	if request == nil {
		request = NewExportKnowledgeRequest()
	}
	response = NewExportKnowledgeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetDatabaseListRequest() (request *DescribeAssetDatabaseListRequest) {
	request = &DescribeAssetDatabaseListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetDatabaseList")
	return
}

func NewDescribeAssetDatabaseListResponse() (response *DescribeAssetDatabaseListResponse) {
	response = &DescribeAssetDatabaseListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取db列表
func (c *Client) DescribeAssetDatabaseList(request *DescribeAssetDatabaseListRequest) (response *DescribeAssetDatabaseListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetDatabaseListRequest()
	}
	response = NewDescribeAssetDatabaseListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteJavaMemShellsRequest() (request *DeleteJavaMemShellsRequest) {
	request = &DeleteJavaMemShellsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteJavaMemShells")
	return
}

func NewDeleteJavaMemShellsResponse() (response *DeleteJavaMemShellsResponse) {
	response = &DeleteJavaMemShellsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除内存马事件
func (c *Client) DeleteJavaMemShells(request *DeleteJavaMemShellsRequest) (response *DeleteJavaMemShellsResponse, err error) {
	if request == nil {
		request = NewDeleteJavaMemShellsRequest()
	}
	response = NewDeleteJavaMemShellsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemoteUpdateUserNameUniqueRequest() (request *DescribeRemoteUpdateUserNameUniqueRequest) {
	request = &DescribeRemoteUpdateUserNameUniqueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeRemoteUpdateUserNameUnique")
	return
}

func NewDescribeRemoteUpdateUserNameUniqueResponse() (response *DescribeRemoteUpdateUserNameUniqueResponse) {
	response = &DescribeRemoteUpdateUserNameUniqueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询远程更新用户名称唯一性
func (c *Client) DescribeRemoteUpdateUserNameUnique(request *DescribeRemoteUpdateUserNameUniqueRequest) (response *DescribeRemoteUpdateUserNameUniqueResponse, err error) {
	if request == nil {
		request = NewDescribeRemoteUpdateUserNameUniqueRequest()
	}
	response = NewDescribeRemoteUpdateUserNameUniqueResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDetailRequest() (request *DescribeVulDetailRequest) {
	request = &DescribeVulDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeVulDetail")
	return
}

func NewDescribeVulDetailResponse() (response *DescribeVulDetailResponse) {
	response = &DescribeVulDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全漏洞 -列表 -详情
func (c *Client) DescribeVulDetail(request *DescribeVulDetailRequest) (response *DescribeVulDetailResponse, err error) {
	if request == nil {
		request = NewDescribeVulDetailRequest()
	}
	response = NewDescribeVulDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetUserListRequest() (request *DescribeAssetUserListRequest) {
	request = &DescribeAssetUserListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetUserList")
	return
}

func NewDescribeAssetUserListResponse() (response *DescribeAssetUserListResponse) {
	response = &DescribeAssetUserListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账号列表
func (c *Client) DescribeAssetUserList(request *DescribeAssetUserListRequest) (response *DescribeAssetUserListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetUserListRequest()
	}
	response = NewDescribeAssetUserListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDNSKnowledgeStatusRequest() (request *ModifyDNSKnowledgeStatusRequest) {
	request = &ModifyDNSKnowledgeStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyDNSKnowledgeStatus")
	return
}

func NewModifyDNSKnowledgeStatusResponse() (response *ModifyDNSKnowledgeStatusResponse) {
	response = &ModifyDNSKnowledgeStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全策略 - DNS知识库 - 开关
func (c *Client) ModifyDNSKnowledgeStatus(request *ModifyDNSKnowledgeStatusRequest) (response *ModifyDNSKnowledgeStatusResponse, err error) {
	if request == nil {
		request = NewModifyDNSKnowledgeStatusRequest()
	}
	response = NewModifyDNSKnowledgeStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSocAuthListRequest() (request *DescribeSocAuthListRequest) {
	request = &DescribeSocAuthListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeSocAuthList")
	return
}

func NewDescribeSocAuthListResponse() (response *DescribeSocAuthListResponse) {
	response = &DescribeSocAuthListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 网页防篡授权详情
func (c *Client) DescribeSocAuthList(request *DescribeSocAuthListRequest) (response *DescribeSocAuthListResponse, err error) {
	if request == nil {
		request = NewDescribeSocAuthListRequest()
	}
	response = NewDescribeSocAuthListResponse()
	err = c.Send(request, response)
	return
}

func NewExportDNSEventsRequest() (request *ExportDNSEventsRequest) {
	request = &ExportDNSEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportDNSEvents")
	return
}

func NewExportDNSEventsResponse() (response *ExportDNSEventsResponse) {
	response = &ExportDNSEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出恶意请求列表
func (c *Client) ExportDNSEvents(request *ExportDNSEventsRequest) (response *ExportDNSEventsResponse, err error) {
	if request == nil {
		request = NewExportDNSEventsRequest()
	}
	response = NewExportDNSEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetDatabaseInfoRequest() (request *DescribeAssetDatabaseInfoRequest) {
	request = &DescribeAssetDatabaseInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetDatabaseInfo")
	return
}

func NewDescribeAssetDatabaseInfoResponse() (response *DescribeAssetDatabaseInfoResponse) {
	response = &DescribeAssetDatabaseInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取DB详情
func (c *Client) DescribeAssetDatabaseInfo(request *DescribeAssetDatabaseInfoRequest) (response *DescribeAssetDatabaseInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetDatabaseInfoRequest()
	}
	response = NewDescribeAssetDatabaseInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFileTamperEventsRequest() (request *DescribeFileTamperEventsRequest) {
	request = &DescribeFileTamperEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeFileTamperEvents")
	return
}

func NewDescribeFileTamperEventsResponse() (response *DescribeFileTamperEventsResponse) {
	response = &DescribeFileTamperEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 核心文件事件列表
func (c *Client) DescribeFileTamperEvents(request *DescribeFileTamperEventsRequest) (response *DescribeFileTamperEventsResponse, err error) {
	if request == nil {
		request = NewDescribeFileTamperEventsRequest()
	}
	response = NewDescribeFileTamperEventsResponse()
	err = c.Send(request, response)
	return
}

func NewFindAgentUuidsByPlatformRequest() (request *FindAgentUuidsByPlatformRequest) {
	request = &FindAgentUuidsByPlatformRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "FindAgentUuidsByPlatform")
	return
}

func NewFindAgentUuidsByPlatformResponse() (response *FindAgentUuidsByPlatformResponse) {
	response = &FindAgentUuidsByPlatformResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 任务管理 - 添加任务- 根据机器类型获取所有uuid
func (c *Client) FindAgentUuidsByPlatform(request *FindAgentUuidsByPlatformRequest) (response *FindAgentUuidsByPlatformResponse, err error) {
	if request == nil {
		request = NewFindAgentUuidsByPlatformRequest()
	}
	response = NewFindAgentUuidsByPlatformResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetAppListRequest() (request *ExportAssetAppListRequest) {
	request = &ExportAssetAppListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAssetAppList")
	return
}

func NewExportAssetAppListResponse() (response *ExportAssetAppListResponse) {
	response = &ExportAssetAppListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理应用列表
func (c *Client) ExportAssetAppList(request *ExportAssetAppListRequest) (response *ExportAssetAppListResponse, err error) {
	if request == nil {
		request = NewExportAssetAppListRequest()
	}
	response = NewExportAssetAppListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVulPackageInfoRequest() (request *ModifyVulPackageInfoRequest) {
	request = &ModifyVulPackageInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyVulPackageInfo")
	return
}

func NewModifyVulPackageInfoResponse() (response *ModifyVulPackageInfoResponse) {
	response = &ModifyVulPackageInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存或编辑任务包文件
func (c *Client) ModifyVulPackageInfo(request *ModifyVulPackageInfoRequest) (response *ModifyVulPackageInfoResponse, err error) {
	if request == nil {
		request = NewModifyVulPackageInfoRequest()
	}
	response = NewModifyVulPackageInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRemoteUpdatePublishItemRequest() (request *CreateRemoteUpdatePublishItemRequest) {
	request = &CreateRemoteUpdatePublishItemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "CreateRemoteUpdatePublishItem")
	return
}

func NewCreateRemoteUpdatePublishItemResponse() (response *CreateRemoteUpdatePublishItemResponse) {
	response = &CreateRemoteUpdatePublishItemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建单个远程更新发布项
func (c *Client) CreateRemoteUpdatePublishItem(request *CreateRemoteUpdatePublishItemRequest) (response *CreateRemoteUpdatePublishItemResponse, err error) {
	if request == nil {
		request = NewCreateRemoteUpdatePublishItemRequest()
	}
	response = NewCreateRemoteUpdatePublishItemResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDNSEventRequest() (request *DeleteDNSEventRequest) {
	request = &DeleteDNSEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteDNSEvent")
	return
}

func NewDeleteDNSEventResponse() (response *DeleteDNSEventResponse) {
	response = &DeleteDNSEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除恶意请求
func (c *Client) DeleteDNSEvent(request *DeleteDNSEventRequest) (response *DeleteDNSEventResponse, err error) {
	if request == nil {
		request = NewDeleteDNSEventRequest()
	}
	response = NewDeleteDNSEventResponse()
	err = c.Send(request, response)
	return
}

func NewPublishKnowledgeConfigRequest() (request *PublishKnowledgeConfigRequest) {
	request = &PublishKnowledgeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "PublishKnowledgeConfig")
	return
}

func NewPublishKnowledgeConfigResponse() (response *PublishKnowledgeConfigResponse) {
	response = &PublishKnowledgeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新发布知识库
func (c *Client) PublishKnowledgeConfig(request *PublishKnowledgeConfigRequest) (response *PublishKnowledgeConfigResponse, err error) {
	if request == nil {
		request = NewPublishKnowledgeConfigRequest()
	}
	response = NewPublishKnowledgeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCwpProductUsagesRequest() (request *DescribeCwpProductUsagesRequest) {
	request = &DescribeCwpProductUsagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeCwpProductUsages")
	return
}

func NewDescribeCwpProductUsagesResponse() (response *DescribeCwpProductUsagesResponse) {
	response = &DescribeCwpProductUsagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用量采集接口
func (c *Client) DescribeCwpProductUsages(request *DescribeCwpProductUsagesRequest) (response *DescribeCwpProductUsagesResponse, err error) {
	if request == nil {
		request = NewDescribeCwpProductUsagesRequest()
	}
	response = NewDescribeCwpProductUsagesResponse()
	err = c.Send(request, response)
	return
}

func NewCancelPublishInstallPackageRequest() (request *CancelPublishInstallPackageRequest) {
	request = &CancelPublishInstallPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "CancelPublishInstallPackage")
	return
}

func NewCancelPublishInstallPackageResponse() (response *CancelPublishInstallPackageResponse) {
	response = &CancelPublishInstallPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 版本管理 - 安装包管理 - 删除安装包记录
func (c *Client) CancelPublishInstallPackage(request *CancelPublishInstallPackageRequest) (response *CancelPublishInstallPackageResponse, err error) {
	if request == nil {
		request = NewCancelPublishInstallPackageRequest()
	}
	response = NewCancelPublishInstallPackageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetDatabaseCountRequest() (request *DescribeAssetDatabaseCountRequest) {
	request = &DescribeAssetDatabaseCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetDatabaseCount")
	return
}

func NewDescribeAssetDatabaseCountResponse() (response *DescribeAssetDatabaseCountResponse) {
	response = &DescribeAssetDatabaseCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有db数量
func (c *Client) DescribeAssetDatabaseCount(request *DescribeAssetDatabaseCountRequest) (response *DescribeAssetDatabaseCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetDatabaseCountRequest()
	}
	response = NewDescribeAssetDatabaseCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHostRiskTrendRequest() (request *DescribeHostRiskTrendRequest) {
	request = &DescribeHostRiskTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeHostRiskTrend")
	return
}

func NewDescribeHostRiskTrendResponse() (response *DescribeHostRiskTrendResponse) {
	response = &DescribeHostRiskTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 风险主机事件趋势
func (c *Client) DescribeHostRiskTrend(request *DescribeHostRiskTrendRequest) (response *DescribeHostRiskTrendResponse, err error) {
	if request == nil {
		request = NewDescribeHostRiskTrendRequest()
	}
	response = NewDescribeHostRiskTrendResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBruteEventsRequest() (request *ModifyBruteEventsRequest) {
	request = &ModifyBruteEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyBruteEvents")
	return
}

func NewModifyBruteEventsResponse() (response *ModifyBruteEventsResponse) {
	response = &ModifyBruteEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 密码破解 -列表 - 忽略/误报/删除
func (c *Client) ModifyBruteEvents(request *ModifyBruteEventsRequest) (response *ModifyBruteEventsResponse, err error) {
	if request == nil {
		request = NewModifyBruteEventsRequest()
	}
	response = NewModifyBruteEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentPushTaskListRequest() (request *DescribeAgentPushTaskListRequest) {
	request = &DescribeAgentPushTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAgentPushTaskList")
	return
}

func NewDescribeAgentPushTaskListResponse() (response *DescribeAgentPushTaskListResponse) {
	response = &DescribeAgentPushTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询客户端命令推送任务列表
func (c *Client) DescribeAgentPushTaskList(request *DescribeAgentPushTaskListRequest) (response *DescribeAgentPushTaskListResponse, err error) {
	if request == nil {
		request = NewDescribeAgentPushTaskListRequest()
	}
	response = NewDescribeAgentPushTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetCoreModuleInfoRequest() (request *DescribeAssetCoreModuleInfoRequest) {
	request = &DescribeAssetCoreModuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetCoreModuleInfo")
	return
}

func NewDescribeAssetCoreModuleInfoResponse() (response *DescribeAssetCoreModuleInfoResponse) {
	response = &DescribeAssetCoreModuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取内核模块详情
func (c *Client) DescribeAssetCoreModuleInfo(request *DescribeAssetCoreModuleInfoRequest) (response *DescribeAssetCoreModuleInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetCoreModuleInfoRequest()
	}
	response = NewDescribeAssetCoreModuleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineItemListRequest() (request *DescribeBaselineItemListRequest) {
	request = &DescribeBaselineItemListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeBaselineItemList")
	return
}

func NewDescribeBaselineItemListResponse() (response *DescribeBaselineItemListResponse) {
	response = &DescribeBaselineItemListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线项检测结果列表
func (c *Client) DescribeBaselineItemList(request *DescribeBaselineItemListRequest) (response *DescribeBaselineItemListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineItemListRequest()
	}
	response = NewDescribeBaselineItemListResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetWebFrameListRequest() (request *ExportAssetWebFrameListRequest) {
	request = &ExportAssetWebFrameListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAssetWebFrameList")
	return
}

func NewExportAssetWebFrameListResponse() (response *ExportAssetWebFrameListResponse) {
	response = &ExportAssetWebFrameListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理Web框架列表
func (c *Client) ExportAssetWebFrameList(request *ExportAssetWebFrameListRequest) (response *ExportAssetWebFrameListResponse, err error) {
	if request == nil {
		request = NewExportAssetWebFrameListRequest()
	}
	response = NewExportAssetWebFrameListResponse()
	err = c.Send(request, response)
	return
}

func NewExportPrivilegeEventsRequest() (request *ExportPrivilegeEventsRequest) {
	request = &ExportPrivilegeEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportPrivilegeEvents")
	return
}

func NewExportPrivilegeEventsResponse() (response *ExportPrivilegeEventsResponse) {
	response = &ExportPrivilegeEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出本地提权列表
func (c *Client) ExportPrivilegeEvents(request *ExportPrivilegeEventsRequest) (response *ExportPrivilegeEventsResponse, err error) {
	if request == nil {
		request = NewExportPrivilegeEventsRequest()
	}
	response = NewExportPrivilegeEventsResponse()
	err = c.Send(request, response)
	return
}

func NewPublishPocRequest() (request *PublishPocRequest) {
	request = &PublishPocRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "PublishPoc")
	return
}

func NewPublishPocResponse() (response *PublishPocResponse) {
	response = &PublishPocResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 异步生成PoC包
func (c *Client) PublishPoc(request *PublishPocRequest) (response *PublishPocResponse, err error) {
	if request == nil {
		request = NewPublishPocRequest()
	}
	response = NewPublishPocResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRemoteUpdateUsersRequest() (request *DeleteRemoteUpdateUsersRequest) {
	request = &DeleteRemoteUpdateUsersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteRemoteUpdateUsers")
	return
}

func NewDeleteRemoteUpdateUsersResponse() (response *DeleteRemoteUpdateUsersResponse) {
	response = &DeleteRemoteUpdateUsersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除远程更新用户
func (c *Client) DeleteRemoteUpdateUsers(request *DeleteRemoteUpdateUsersRequest) (response *DeleteRemoteUpdateUsersResponse, err error) {
	if request == nil {
		request = NewDeleteRemoteUpdateUsersRequest()
	}
	response = NewDeleteRemoteUpdateUsersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebServiceProcessListRequest() (request *DescribeAssetWebServiceProcessListRequest) {
	request = &DescribeAssetWebServiceProcessListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetWebServiceProcessList")
	return
}

func NewDescribeAssetWebServiceProcessListResponse() (response *DescribeAssetWebServiceProcessListResponse) {
	response = &DescribeAssetWebServiceProcessListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Web服务关联进程列表
func (c *Client) DescribeAssetWebServiceProcessList(request *DescribeAssetWebServiceProcessListRequest) (response *DescribeAssetWebServiceProcessListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebServiceProcessListRequest()
	}
	response = NewDescribeAssetWebServiceProcessListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefenceMachineListRequest() (request *DescribeVulDefenceMachineListRequest) {
	request = &DescribeVulDefenceMachineListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeVulDefenceMachineList")
	return
}

func NewDescribeVulDefenceMachineListResponse() (response *DescribeVulDefenceMachineListResponse) {
	response = &DescribeVulDefenceMachineListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防御主机列表
func (c *Client) DescribeVulDefenceMachineList(request *DescribeVulDefenceMachineListRequest) (response *DescribeVulDefenceMachineListResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefenceMachineListRequest()
	}
	response = NewDescribeVulDefenceMachineListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulTopRequest() (request *DescribeVulTopRequest) {
	request = &DescribeVulTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeVulTop")
	return
}

func NewDescribeVulTopResponse() (response *DescribeVulTopResponse) {
	response = &DescribeVulTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取漏洞top
func (c *Client) DescribeVulTop(request *DescribeVulTopRequest) (response *DescribeVulTopResponse, err error) {
	if request == nil {
		request = NewDescribeVulTopRequest()
	}
	response = NewDescribeVulTopResponse()
	err = c.Send(request, response)
	return
}

func NewExportHostSafetyStatusRequest() (request *ExportHostSafetyStatusRequest) {
	request = &ExportHostSafetyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportHostSafetyStatus")
	return
}

func NewExportHostSafetyStatusResponse() (response *ExportHostSafetyStatusResponse) {
	response = &ExportHostSafetyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全状态导出
func (c *Client) ExportHostSafetyStatus(request *ExportHostSafetyStatusRequest) (response *ExportHostSafetyStatusResponse, err error) {
	if request == nil {
		request = NewExportHostSafetyStatusRequest()
	}
	response = NewExportHostSafetyStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemoteUpdateTagListRequest() (request *DescribeRemoteUpdateTagListRequest) {
	request = &DescribeRemoteUpdateTagListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeRemoteUpdateTagList")
	return
}

func NewDescribeRemoteUpdateTagListResponse() (response *DescribeRemoteUpdateTagListResponse) {
	response = &DescribeRemoteUpdateTagListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取远程更新标签列表
func (c *Client) DescribeRemoteUpdateTagList(request *DescribeRemoteUpdateTagListRequest) (response *DescribeRemoteUpdateTagListResponse, err error) {
	if request == nil {
		request = NewDescribeRemoteUpdateTagListRequest()
	}
	response = NewDescribeRemoteUpdateTagListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDNSKnowledgeInfoRequest() (request *ModifyDNSKnowledgeInfoRequest) {
	request = &ModifyDNSKnowledgeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyDNSKnowledgeInfo")
	return
}

func NewModifyDNSKnowledgeInfoResponse() (response *ModifyDNSKnowledgeInfoResponse) {
	response = &ModifyDNSKnowledgeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全策略 - DNS知识库 - 编辑/添加
func (c *Client) ModifyDNSKnowledgeInfo(request *ModifyDNSKnowledgeInfoRequest) (response *ModifyDNSKnowledgeInfoResponse, err error) {
	if request == nil {
		request = NewModifyDNSKnowledgeInfoRequest()
	}
	response = NewModifyDNSKnowledgeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBruteAttacksRequest() (request *DeleteBruteAttacksRequest) {
	request = &DeleteBruteAttacksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteBruteAttacks")
	return
}

func NewDeleteBruteAttacksResponse() (response *DeleteBruteAttacksResponse) {
	response = &DeleteBruteAttacksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除密码破解
func (c *Client) DeleteBruteAttacks(request *DeleteBruteAttacksRequest) (response *DeleteBruteAttacksResponse, err error) {
	if request == nil {
		request = NewDeleteBruteAttacksRequest()
	}
	response = NewDeleteBruteAttacksResponse()
	err = c.Send(request, response)
	return
}

func NewModifyJavaMemShellPluginSwitchRequest() (request *ModifyJavaMemShellPluginSwitchRequest) {
	request = &ModifyJavaMemShellPluginSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyJavaMemShellPluginSwitch")
	return
}

func NewModifyJavaMemShellPluginSwitchResponse() (response *ModifyJavaMemShellPluginSwitchResponse) {
	response = &ModifyJavaMemShellPluginSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改内存马插件状态
func (c *Client) ModifyJavaMemShellPluginSwitch(request *ModifyJavaMemShellPluginSwitchRequest) (response *ModifyJavaMemShellPluginSwitchResponse, err error) {
	if request == nil {
		request = NewModifyJavaMemShellPluginSwitchRequest()
	}
	response = NewModifyJavaMemShellPluginSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebAppCountRequest() (request *DescribeAssetWebAppCountRequest) {
	request = &DescribeAssetWebAppCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetWebAppCount")
	return
}

func NewDescribeAssetWebAppCountResponse() (response *DescribeAssetWebAppCountResponse) {
	response = &DescribeAssetWebAppCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有Web应用数量
func (c *Client) DescribeAssetWebAppCount(request *DescribeAssetWebAppCountRequest) (response *DescribeAssetWebAppCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebAppCountRequest()
	}
	response = NewDescribeAssetWebAppCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemoteUpdateConfigMainInfoRequest() (request *DescribeRemoteUpdateConfigMainInfoRequest) {
	request = &DescribeRemoteUpdateConfigMainInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeRemoteUpdateConfigMainInfo")
	return
}

func NewDescribeRemoteUpdateConfigMainInfoResponse() (response *DescribeRemoteUpdateConfigMainInfoResponse) {
	response = &DescribeRemoteUpdateConfigMainInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询远程更新配置主要信息
func (c *Client) DescribeRemoteUpdateConfigMainInfo(request *DescribeRemoteUpdateConfigMainInfoRequest) (response *DescribeRemoteUpdateConfigMainInfoResponse, err error) {
	if request == nil {
		request = NewDescribeRemoteUpdateConfigMainInfoRequest()
	}
	response = NewDescribeRemoteUpdateConfigMainInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetProcessInfoListRequest() (request *DescribeAssetProcessInfoListRequest) {
	request = &DescribeAssetProcessInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetProcessInfoList")
	return
}

func NewDescribeAssetProcessInfoListResponse() (response *DescribeAssetProcessInfoListResponse) {
	response = &DescribeAssetProcessInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取进程列表
func (c *Client) DescribeAssetProcessInfoList(request *DescribeAssetProcessInfoListRequest) (response *DescribeAssetProcessInfoListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetProcessInfoListRequest()
	}
	response = NewDescribeAssetProcessInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulEventsRequest() (request *DescribeVulEventsRequest) {
	request = &DescribeVulEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeVulEvents")
	return
}

func NewDescribeVulEventsResponse() (response *DescribeVulEventsResponse) {
	response = &DescribeVulEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞&基线管理 - 安全漏洞 -列表
func (c *Client) DescribeVulEvents(request *DescribeVulEventsRequest) (response *DescribeVulEventsResponse, err error) {
	if request == nil {
		request = NewDescribeVulEventsRequest()
	}
	response = NewDescribeVulEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZKGroupFullListRequest() (request *DescribeZKGroupFullListRequest) {
	request = &DescribeZKGroupFullListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeZKGroupFullList")
	return
}

func NewDescribeZKGroupFullListResponse() (response *DescribeZKGroupFullListResponse) {
	response = &DescribeZKGroupFullListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 系统管理 - Zookeeper配置 - 查看 - 查看配置
func (c *Client) DescribeZKGroupFullList(request *DescribeZKGroupFullListRequest) (response *DescribeZKGroupFullListResponse, err error) {
	if request == nil {
		request = NewDescribeZKGroupFullListRequest()
	}
	response = NewDescribeZKGroupFullListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDNSEventsRequest() (request *ModifyDNSEventsRequest) {
	request = &ModifyDNSEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyDNSEvents")
	return
}

func NewModifyDNSEventsResponse() (response *ModifyDNSEventsResponse) {
	response = &ModifyDNSEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改恶意请求状态
func (c *Client) ModifyDNSEvents(request *ModifyDNSEventsRequest) (response *ModifyDNSEventsResponse, err error) {
	if request == nil {
		request = NewModifyDNSEventsRequest()
	}
	response = NewModifyDNSEventsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTaskComponentsSwitchRequest() (request *ModifyTaskComponentsSwitchRequest) {
	request = &ModifyTaskComponentsSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyTaskComponentsSwitch")
	return
}

func NewModifyTaskComponentsSwitchResponse() (response *ModifyTaskComponentsSwitchResponse) {
	response = &ModifyTaskComponentsSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新检测开关
func (c *Client) ModifyTaskComponentsSwitch(request *ModifyTaskComponentsSwitchRequest) (response *ModifyTaskComponentsSwitchResponse, err error) {
	if request == nil {
		request = NewModifyTaskComponentsSwitchRequest()
	}
	response = NewModifyTaskComponentsSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVulTypeInfoRequest() (request *ModifyVulTypeInfoRequest) {
	request = &ModifyVulTypeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyVulTypeInfo")
	return
}

func NewModifyVulTypeInfoResponse() (response *ModifyVulTypeInfoResponse) {
	response = &ModifyVulTypeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新/添加漏洞类型
func (c *Client) ModifyVulTypeInfo(request *ModifyVulTypeInfoRequest) (response *ModifyVulTypeInfoResponse, err error) {
	if request == nil {
		request = NewModifyVulTypeInfoRequest()
	}
	response = NewModifyVulTypeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTaskComponentsInfoRequest() (request *ModifyTaskComponentsInfoRequest) {
	request = &ModifyTaskComponentsInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyTaskComponentsInfo")
	return
}

func NewModifyTaskComponentsInfoResponse() (response *ModifyTaskComponentsInfoResponse) {
	response = &ModifyTaskComponentsInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新漏洞组件
func (c *Client) ModifyTaskComponentsInfo(request *ModifyTaskComponentsInfoRequest) (response *ModifyTaskComponentsInfoResponse, err error) {
	if request == nil {
		request = NewModifyTaskComponentsInfoRequest()
	}
	response = NewModifyTaskComponentsInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyFileTamperEventsRequest() (request *ModifyFileTamperEventsRequest) {
	request = &ModifyFileTamperEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyFileTamperEvents")
	return
}

func NewModifyFileTamperEventsResponse() (response *ModifyFileTamperEventsResponse) {
	response = &ModifyFileTamperEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新核心文件事件
func (c *Client) ModifyFileTamperEvents(request *ModifyFileTamperEventsRequest) (response *ModifyFileTamperEventsResponse, err error) {
	if request == nil {
		request = NewModifyFileTamperEventsRequest()
	}
	response = NewModifyFileTamperEventsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRemoteUpdateConfigAutoUpdateRequest() (request *ModifyRemoteUpdateConfigAutoUpdateRequest) {
	request = &ModifyRemoteUpdateConfigAutoUpdateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyRemoteUpdateConfigAutoUpdate")
	return
}

func NewModifyRemoteUpdateConfigAutoUpdateResponse() (response *ModifyRemoteUpdateConfigAutoUpdateResponse) {
	response = &ModifyRemoteUpdateConfigAutoUpdateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改远程更新配置的自动升级
func (c *Client) ModifyRemoteUpdateConfigAutoUpdate(request *ModifyRemoteUpdateConfigAutoUpdateRequest) (response *ModifyRemoteUpdateConfigAutoUpdateResponse, err error) {
	if request == nil {
		request = NewModifyRemoteUpdateConfigAutoUpdateRequest()
	}
	response = NewModifyRemoteUpdateConfigAutoUpdateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUpdateConfigRequest() (request *ModifyUpdateConfigRequest) {
	request = &ModifyUpdateConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyUpdateConfig")
	return
}

func NewModifyUpdateConfigResponse() (response *ModifyUpdateConfigResponse) {
	response = &ModifyUpdateConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 更新版本配置
func (c *Client) ModifyUpdateConfig(request *ModifyUpdateConfigRequest) (response *ModifyUpdateConfigResponse, err error) {
	if request == nil {
		request = NewModifyUpdateConfigRequest()
	}
	response = NewModifyUpdateConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineFileTamperRulesRequest() (request *DescribeMachineFileTamperRulesRequest) {
	request = &DescribeMachineFileTamperRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeMachineFileTamperRules")
	return
}

func NewDescribeMachineFileTamperRulesResponse() (response *DescribeMachineFileTamperRulesResponse) {
	response = &DescribeMachineFileTamperRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主机相关核心文件监控规则列 表
func (c *Client) DescribeMachineFileTamperRules(request *DescribeMachineFileTamperRulesRequest) (response *DescribeMachineFileTamperRulesResponse, err error) {
	if request == nil {
		request = NewDescribeMachineFileTamperRulesRequest()
	}
	response = NewDescribeMachineFileTamperRulesResponse()
	err = c.Send(request, response)
	return
}

func NewPublishVulScanConfigRequest() (request *PublishVulScanConfigRequest) {
	request = &PublishVulScanConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "PublishVulScanConfig")
	return
}

func NewPublishVulScanConfigResponse() (response *PublishVulScanConfigResponse) {
	response = &PublishVulScanConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置文件管理 - 发布
func (c *Client) PublishVulScanConfig(request *PublishVulScanConfigRequest) (response *PublishVulScanConfigResponse, err error) {
	if request == nil {
		request = NewPublishVulScanConfigRequest()
	}
	response = NewPublishVulScanConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUpdateTaskRequest() (request *DescribeUpdateTaskRequest) {
	request = &DescribeUpdateTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeUpdateTask")
	return
}

func NewDescribeUpdateTaskResponse() (response *DescribeUpdateTaskResponse) {
	response = &DescribeUpdateTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 列表
func (c *Client) DescribeUpdateTask(request *DescribeUpdateTaskRequest) (response *DescribeUpdateTaskResponse, err error) {
	if request == nil {
		request = NewDescribeUpdateTaskRequest()
	}
	response = NewDescribeUpdateTaskResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRemoteUpdateTagRequest() (request *ModifyRemoteUpdateTagRequest) {
	request = &ModifyRemoteUpdateTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyRemoteUpdateTag")
	return
}

func NewModifyRemoteUpdateTagResponse() (response *ModifyRemoteUpdateTagResponse) {
	response = &ModifyRemoteUpdateTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改远程更新标签
func (c *Client) ModifyRemoteUpdateTag(request *ModifyRemoteUpdateTagRequest) (response *ModifyRemoteUpdateTagResponse, err error) {
	if request == nil {
		request = NewModifyRemoteUpdateTagRequest()
	}
	response = NewModifyRemoteUpdateTagResponse()
	err = c.Send(request, response)
	return
}

func NewSyncAssetScanRequest() (request *SyncAssetScanRequest) {
	request = &SyncAssetScanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "SyncAssetScan")
	return
}

func NewSyncAssetScanResponse() (response *SyncAssetScanResponse) {
	response = &SyncAssetScanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步资产扫描信息
func (c *Client) SyncAssetScan(request *SyncAssetScanRequest) (response *SyncAssetScanResponse, err error) {
	if request == nil {
		request = NewSyncAssetScanRequest()
	}
	response = NewSyncAssetScanResponse()
	err = c.Send(request, response)
	return
}

func NewStartUpdateTaskRequest() (request *StartUpdateTaskRequest) {
	request = &StartUpdateTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "StartUpdateTask")
	return
}

func NewStartUpdateTaskResponse() (response *StartUpdateTaskResponse) {
	response = &StartUpdateTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 启动升级任务
func (c *Client) StartUpdateTask(request *StartUpdateTaskRequest) (response *StartUpdateTaskResponse, err error) {
	if request == nil {
		request = NewStartUpdateTaskRequest()
	}
	response = NewStartUpdateTaskResponse()
	err = c.Send(request, response)
	return
}

func NewAddPushTaskRequest() (request *AddPushTaskRequest) {
	request = &AddPushTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "AddPushTask")
	return
}

func NewAddPushTaskResponse() (response *AddPushTaskResponse) {
	response = &AddPushTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 资产管理 - 命令推送 - 新建命令/复制 - 保存
func (c *Client) AddPushTask(request *AddPushTaskRequest) (response *AddPushTaskResponse, err error) {
	if request == nil {
		request = NewAddPushTaskRequest()
	}
	response = NewAddPushTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUpdateTaskStatusRequest() (request *DescribeUpdateTaskStatusRequest) {
	request = &DescribeUpdateTaskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeUpdateTaskStatus")
	return
}

func NewDescribeUpdateTaskStatusResponse() (response *DescribeUpdateTaskStatusResponse) {
	response = &DescribeUpdateTaskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 获取版本任务执行状态列表
func (c *Client) DescribeUpdateTaskStatus(request *DescribeUpdateTaskStatusRequest) (response *DescribeUpdateTaskStatusResponse, err error) {
	if request == nil {
		request = NewDescribeUpdateTaskStatusRequest()
	}
	response = NewDescribeUpdateTaskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUpdateTaskGrayListRequest() (request *DescribeUpdateTaskGrayListRequest) {
	request = &DescribeUpdateTaskGrayListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeUpdateTaskGrayList")
	return
}

func NewDescribeUpdateTaskGrayListResponse() (response *DescribeUpdateTaskGrayListResponse) {
	response = &DescribeUpdateTaskGrayListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 获取版本灰度列表(升级目标)
func (c *Client) DescribeUpdateTaskGrayList(request *DescribeUpdateTaskGrayListRequest) (response *DescribeUpdateTaskGrayListResponse, err error) {
	if request == nil {
		request = NewDescribeUpdateTaskGrayListRequest()
	}
	response = NewDescribeUpdateTaskGrayListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTaskComponentsRequest() (request *DeleteTaskComponentsRequest) {
	request = &DeleteTaskComponentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteTaskComponents")
	return
}

func NewDeleteTaskComponentsResponse() (response *DeleteTaskComponentsResponse) {
	response = &DeleteTaskComponentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除漏洞组件
func (c *Client) DeleteTaskComponents(request *DeleteTaskComponentsRequest) (response *DeleteTaskComponentsResponse, err error) {
	if request == nil {
		request = NewDeleteTaskComponentsRequest()
	}
	response = NewDeleteTaskComponentsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetMachineDetailRequest() (request *DescribeAssetMachineDetailRequest) {
	request = &DescribeAssetMachineDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetMachineDetail")
	return
}

func NewDescribeAssetMachineDetailResponse() (response *DescribeAssetMachineDetailResponse) {
	response = &DescribeAssetMachineDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产管理主机资源详细信息
func (c *Client) DescribeAssetMachineDetail(request *DescribeAssetMachineDetailRequest) (response *DescribeAssetMachineDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAssetMachineDetailRequest()
	}
	response = NewDescribeAssetMachineDetailResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateEvilMalwareRequest() (request *UpdateEvilMalwareRequest) {
	request = &UpdateEvilMalwareRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "UpdateEvilMalware")
	return
}

func NewUpdateEvilMalwareResponse() (response *UpdateEvilMalwareResponse) {
	response = &UpdateEvilMalwareResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 入侵检测确认恶意文件成功
func (c *Client) UpdateEvilMalware(request *UpdateEvilMalwareRequest) (response *UpdateEvilMalwareResponse, err error) {
	if request == nil {
		request = NewUpdateEvilMalwareRequest()
	}
	response = NewUpdateEvilMalwareResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemoteUpdateConfigVersionListRequest() (request *DescribeRemoteUpdateConfigVersionListRequest) {
	request = &DescribeRemoteUpdateConfigVersionListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeRemoteUpdateConfigVersionList")
	return
}

func NewDescribeRemoteUpdateConfigVersionListResponse() (response *DescribeRemoteUpdateConfigVersionListResponse) {
	response = &DescribeRemoteUpdateConfigVersionListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询远程更新配置版本列表
func (c *Client) DescribeRemoteUpdateConfigVersionList(request *DescribeRemoteUpdateConfigVersionListRequest) (response *DescribeRemoteUpdateConfigVersionListResponse, err error) {
	if request == nil {
		request = NewDescribeRemoteUpdateConfigVersionListRequest()
	}
	response = NewDescribeRemoteUpdateConfigVersionListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDNSEventsRequest() (request *DescribeDNSEventsRequest) {
	request = &DescribeDNSEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeDNSEvents")
	return
}

func NewDescribeDNSEventsResponse() (response *DescribeDNSEventsResponse) {
	response = &DescribeDNSEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恶意请求 - 列表
func (c *Client) DescribeDNSEvents(request *DescribeDNSEventsRequest) (response *DescribeDNSEventsResponse, err error) {
	if request == nil {
		request = NewDescribeDNSEventsRequest()
	}
	response = NewDescribeDNSEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJavaMemShellListRequest() (request *DescribeJavaMemShellListRequest) {
	request = &DescribeJavaMemShellListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeJavaMemShellList")
	return
}

func NewDescribeJavaMemShellListResponse() (response *DescribeJavaMemShellListResponse) {
	response = &DescribeJavaMemShellListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取java内存马列表
func (c *Client) DescribeJavaMemShellList(request *DescribeJavaMemShellListRequest) (response *DescribeJavaMemShellListResponse, err error) {
	if request == nil {
		request = NewDescribeJavaMemShellListRequest()
	}
	response = NewDescribeJavaMemShellListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetInitServiceListRequest() (request *DescribeAssetInitServiceListRequest) {
	request = &DescribeAssetInitServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetInitServiceList")
	return
}

func NewDescribeAssetInitServiceListResponse() (response *DescribeAssetInitServiceListResponse) {
	response = &DescribeAssetInitServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产管理启动服务列表
func (c *Client) DescribeAssetInitServiceList(request *DescribeAssetInitServiceListRequest) (response *DescribeAssetInitServiceListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetInitServiceListRequest()
	}
	response = NewDescribeAssetInitServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachinesSimpleRequest() (request *DescribeMachinesSimpleRequest) {
	request = &DescribeMachinesSimpleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeMachinesSimple")
	return
}

func NewDescribeMachinesSimpleResponse() (response *DescribeMachinesSimpleResponse) {
	response = &DescribeMachinesSimpleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeMachinesSimple) 用于获取主机列表。
func (c *Client) DescribeMachinesSimple(request *DescribeMachinesSimpleRequest) (response *DescribeMachinesSimpleResponse, err error) {
	if request == nil {
		request = NewDescribeMachinesSimpleRequest()
	}
	response = NewDescribeMachinesSimpleResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetMachineListRequest() (request *ExportAssetMachineListRequest) {
	request = &ExportAssetMachineListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAssetMachineList")
	return
}

func NewExportAssetMachineListResponse() (response *ExportAssetMachineListResponse) {
	response = &ExportAssetMachineListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资源监控列表
func (c *Client) ExportAssetMachineList(request *ExportAssetMachineListRequest) (response *ExportAssetMachineListResponse, err error) {
	if request == nil {
		request = NewExportAssetMachineListRequest()
	}
	response = NewExportAssetMachineListResponse()
	err = c.Send(request, response)
	return
}

func NewImportSpecimenFileRequest() (request *ImportSpecimenFileRequest) {
	request = &ImportSpecimenFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ImportSpecimenFile")
	return
}

func NewImportSpecimenFileResponse() (response *ImportSpecimenFileResponse) {
	response = &ImportSpecimenFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全策略 - 样本管理 - 导入样本 -保存
func (c *Client) ImportSpecimenFile(request *ImportSpecimenFileRequest) (response *ImportSpecimenFileResponse, err error) {
	if request == nil {
		request = NewImportSpecimenFileRequest()
	}
	response = NewImportSpecimenFileResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLoginEventsRequest() (request *ModifyLoginEventsRequest) {
	request = &ModifyLoginEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyLoginEvents")
	return
}

func NewModifyLoginEventsResponse() (response *ModifyLoginEventsResponse) {
	response = &ModifyLoginEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 异常登录 -列表 - 忽略/误报/删除
func (c *Client) ModifyLoginEvents(request *ModifyLoginEventsRequest) (response *ModifyLoginEventsResponse, err error) {
	if request == nil {
		request = NewModifyLoginEventsRequest()
	}
	response = NewModifyLoginEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetDiskListRequest() (request *DescribeAssetDiskListRequest) {
	request = &DescribeAssetDiskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetDiskList")
	return
}

func NewDescribeAssetDiskListResponse() (response *DescribeAssetDiskListResponse) {
	response = &DescribeAssetDiskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主机磁盘分区列表
func (c *Client) DescribeAssetDiskList(request *DescribeAssetDiskListRequest) (response *DescribeAssetDiskListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetDiskListRequest()
	}
	response = NewDescribeAssetDiskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebLocationCountRequest() (request *DescribeAssetWebLocationCountRequest) {
	request = &DescribeAssetWebLocationCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetWebLocationCount")
	return
}

func NewDescribeAssetWebLocationCountResponse() (response *DescribeAssetWebLocationCountResponse) {
	response = &DescribeAssetWebLocationCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有Web站点数量
func (c *Client) DescribeAssetWebLocationCount(request *DescribeAssetWebLocationCountRequest) (response *DescribeAssetWebLocationCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebLocationCountRequest()
	}
	response = NewDescribeAssetWebLocationCountResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUpdateTaskRequest() (request *DeleteUpdateTaskRequest) {
	request = &DeleteUpdateTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteUpdateTask")
	return
}

func NewDeleteUpdateTaskResponse() (response *DeleteUpdateTaskResponse) {
	response = &DeleteUpdateTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 删除升级任务记录
func (c *Client) DeleteUpdateTask(request *DeleteUpdateTaskRequest) (response *DeleteUpdateTaskResponse, err error) {
	if request == nil {
		request = NewDeleteUpdateTaskRequest()
	}
	response = NewDeleteUpdateTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFileTamperEventRuleInfoRequest() (request *DescribeFileTamperEventRuleInfoRequest) {
	request = &DescribeFileTamperEventRuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeFileTamperEventRuleInfo")
	return
}

func NewDescribeFileTamperEventRuleInfoResponse() (response *DescribeFileTamperEventRuleInfoResponse) {
	response = &DescribeFileTamperEventRuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看产生事件时规则详情接口
func (c *Client) DescribeFileTamperEventRuleInfo(request *DescribeFileTamperEventRuleInfoRequest) (response *DescribeFileTamperEventRuleInfoResponse, err error) {
	if request == nil {
		request = NewDescribeFileTamperEventRuleInfoRequest()
	}
	response = NewDescribeFileTamperEventRuleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetAppProcessListRequest() (request *DescribeAssetAppProcessListRequest) {
	request = &DescribeAssetAppProcessListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetAppProcessList")
	return
}

func NewDescribeAssetAppProcessListResponse() (response *DescribeAssetAppProcessListResponse) {
	response = &DescribeAssetAppProcessListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取软件关联进程列表
func (c *Client) DescribeAssetAppProcessList(request *DescribeAssetAppProcessListRequest) (response *DescribeAssetAppProcessListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetAppProcessListRequest()
	}
	response = NewDescribeAssetAppProcessListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetHostTotalCountRequest() (request *DescribeAssetHostTotalCountRequest) {
	request = &DescribeAssetHostTotalCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetHostTotalCount")
	return
}

func NewDescribeAssetHostTotalCountResponse() (response *DescribeAssetHostTotalCountResponse) {
	response = &DescribeAssetHostTotalCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主机各种类型资源数量
func (c *Client) DescribeAssetHostTotalCount(request *DescribeAssetHostTotalCountRequest) (response *DescribeAssetHostTotalCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetHostTotalCountRequest()
	}
	response = NewDescribeAssetHostTotalCountResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetInitServiceListRequest() (request *ExportAssetInitServiceListRequest) {
	request = &ExportAssetInitServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAssetInitServiceList")
	return
}

func NewExportAssetInitServiceListResponse() (response *ExportAssetInitServiceListResponse) {
	response = &ExportAssetInitServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理启动服务列表
func (c *Client) ExportAssetInitServiceList(request *ExportAssetInitServiceListRequest) (response *ExportAssetInitServiceListResponse, err error) {
	if request == nil {
		request = NewExportAssetInitServiceListRequest()
	}
	response = NewExportAssetInitServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWeeksRequest() (request *DescribeWeeksRequest) {
	request = &DescribeWeeksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeWeeks")
	return
}

func NewDescribeWeeksResponse() (response *DescribeWeeksResponse) {
	response = &DescribeWeeksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查找当前周列表
func (c *Client) DescribeWeeks(request *DescribeWeeksRequest) (response *DescribeWeeksResponse, err error) {
	if request == nil {
		request = NewDescribeWeeksRequest()
	}
	response = NewDescribeWeeksResponse()
	err = c.Send(request, response)
	return
}

func NewExportCustomerLicenseListRequest() (request *ExportCustomerLicenseListRequest) {
	request = &ExportCustomerLicenseListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportCustomerLicenseList")
	return
}

func NewExportCustomerLicenseListResponse() (response *ExportCustomerLicenseListResponse) {
	response = &ExportCustomerLicenseListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// License列表 - 列表 - 导出
func (c *Client) ExportCustomerLicenseList(request *ExportCustomerLicenseListRequest) (response *ExportCustomerLicenseListResponse, err error) {
	if request == nil {
		request = NewExportCustomerLicenseListRequest()
	}
	response = NewExportCustomerLicenseListResponse()
	err = c.Send(request, response)
	return
}

func NewModifySpecimenStatusRequest() (request *ModifySpecimenStatusRequest) {
	request = &ModifySpecimenStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifySpecimenStatus")
	return
}

func NewModifySpecimenStatusResponse() (response *ModifySpecimenStatusResponse) {
	response = &ModifySpecimenStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 文件查杀 - 列表 - 样本详情 - 设置属性
func (c *Client) ModifySpecimenStatus(request *ModifySpecimenStatusRequest) (response *ModifySpecimenStatusResponse, err error) {
	if request == nil {
		request = NewModifySpecimenStatusRequest()
	}
	response = NewModifySpecimenStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetPortInfoListRequest() (request *DescribeAssetPortInfoListRequest) {
	request = &DescribeAssetPortInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetPortInfoList")
	return
}

func NewDescribeAssetPortInfoListResponse() (response *DescribeAssetPortInfoListResponse) {
	response = &DescribeAssetPortInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取端口列表
func (c *Client) DescribeAssetPortInfoList(request *DescribeAssetPortInfoListRequest) (response *DescribeAssetPortInfoListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetPortInfoListRequest()
	}
	response = NewDescribeAssetPortInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewExportMalwareEventsRequest() (request *ExportMalwareEventsRequest) {
	request = &ExportMalwareEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportMalwareEvents")
	return
}

func NewExportMalwareEventsResponse() (response *ExportMalwareEventsResponse) {
	response = &ExportMalwareEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 文件查杀 - 列表 - 导出
func (c *Client) ExportMalwareEvents(request *ExportMalwareEventsRequest) (response *ExportMalwareEventsResponse, err error) {
	if request == nil {
		request = NewExportMalwareEventsRequest()
	}
	response = NewExportMalwareEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebServiceCountRequest() (request *DescribeAssetWebServiceCountRequest) {
	request = &DescribeAssetWebServiceCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetWebServiceCount")
	return
}

func NewDescribeAssetWebServiceCountResponse() (response *DescribeAssetWebServiceCountResponse) {
	response = &DescribeAssetWebServiceCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有Web服务数
func (c *Client) DescribeAssetWebServiceCount(request *DescribeAssetWebServiceCountRequest) (response *DescribeAssetWebServiceCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebServiceCountRequest()
	}
	response = NewDescribeAssetWebServiceCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFileTamperRuleCountRequest() (request *DescribeFileTamperRuleCountRequest) {
	request = &DescribeFileTamperRuleCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeFileTamperRuleCount")
	return
}

func NewDescribeFileTamperRuleCountResponse() (response *DescribeFileTamperRuleCountResponse) {
	response = &DescribeFileTamperRuleCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主机关联文件监控规则数量
func (c *Client) DescribeFileTamperRuleCount(request *DescribeFileTamperRuleCountRequest) (response *DescribeFileTamperRuleCountResponse, err error) {
	if request == nil {
		request = NewDescribeFileTamperRuleCountRequest()
	}
	response = NewDescribeFileTamperRuleCountResponse()
	err = c.Send(request, response)
	return
}

func NewExportDefenseVulInfoRequest() (request *ExportDefenseVulInfoRequest) {
	request = &ExportDefenseVulInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportDefenseVulInfo")
	return
}

func NewExportDefenseVulInfoResponse() (response *ExportDefenseVulInfoResponse) {
	response = &ExportDefenseVulInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞管理-导出防御漏洞详情
func (c *Client) ExportDefenseVulInfo(request *ExportDefenseVulInfoRequest) (response *ExportDefenseVulInfoResponse, err error) {
	if request == nil {
		request = NewExportDefenseVulInfoRequest()
	}
	response = NewExportDefenseVulInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetUserCountRequest() (request *DescribeAssetUserCountRequest) {
	request = &DescribeAssetUserCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetUserCount")
	return
}

func NewDescribeAssetUserCountResponse() (response *DescribeAssetUserCountResponse) {
	response = &DescribeAssetUserCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有账号数量
func (c *Client) DescribeAssetUserCount(request *DescribeAssetUserCountRequest) (response *DescribeAssetUserCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetUserCountRequest()
	}
	response = NewDescribeAssetUserCountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRemoteUpdateUserRequest() (request *ModifyRemoteUpdateUserRequest) {
	request = &ModifyRemoteUpdateUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyRemoteUpdateUser")
	return
}

func NewModifyRemoteUpdateUserResponse() (response *ModifyRemoteUpdateUserResponse) {
	response = &ModifyRemoteUpdateUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改远程更新用户信息
func (c *Client) ModifyRemoteUpdateUser(request *ModifyRemoteUpdateUserRequest) (response *ModifyRemoteUpdateUserResponse, err error) {
	if request == nil {
		request = NewModifyRemoteUpdateUserRequest()
	}
	response = NewModifyRemoteUpdateUserResponse()
	err = c.Send(request, response)
	return
}

func NewExportBaselineItemListRequest() (request *ExportBaselineItemListRequest) {
	request = &ExportBaselineItemListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportBaselineItemList")
	return
}

func NewExportBaselineItemListResponse() (response *ExportBaselineItemListResponse) {
	response = &ExportBaselineItemListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出检测项结果列表
func (c *Client) ExportBaselineItemList(request *ExportBaselineItemListRequest) (response *ExportBaselineItemListResponse, err error) {
	if request == nil {
		request = NewExportBaselineItemListRequest()
	}
	response = NewExportBaselineItemListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRemoteUpdateUserRequest() (request *CreateRemoteUpdateUserRequest) {
	request = &CreateRemoteUpdateUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "CreateRemoteUpdateUser")
	return
}

func NewCreateRemoteUpdateUserResponse() (response *CreateRemoteUpdateUserResponse) {
	response = &CreateRemoteUpdateUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加远程更新用户
func (c *Client) CreateRemoteUpdateUser(request *CreateRemoteUpdateUserRequest) (response *CreateRemoteUpdateUserResponse, err error) {
	if request == nil {
		request = NewCreateRemoteUpdateUserRequest()
	}
	response = NewCreateRemoteUpdateUserResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentAlarmRequest() (request *DescribeAgentAlarmRequest) {
	request = &DescribeAgentAlarmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAgentAlarm")
	return
}

func NewDescribeAgentAlarmResponse() (response *DescribeAgentAlarmResponse) {
	response = &DescribeAgentAlarmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查找告警信息
func (c *Client) DescribeAgentAlarm(request *DescribeAgentAlarmRequest) (response *DescribeAgentAlarmResponse, err error) {
	if request == nil {
		request = NewDescribeAgentAlarmRequest()
	}
	response = NewDescribeAgentAlarmResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOverviewStatRequest() (request *DescribeOverviewStatRequest) {
	request = &DescribeOverviewStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeOverviewStat")
	return
}

func NewDescribeOverviewStatResponse() (response *DescribeOverviewStatResponse) {
	response = &DescribeOverviewStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全总览
func (c *Client) DescribeOverviewStat(request *DescribeOverviewStatRequest) (response *DescribeOverviewStatResponse, err error) {
	if request == nil {
		request = NewDescribeOverviewStatRequest()
	}
	response = NewDescribeOverviewStatResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRemoteUpdateConfigsRequest() (request *DeleteRemoteUpdateConfigsRequest) {
	request = &DeleteRemoteUpdateConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteRemoteUpdateConfigs")
	return
}

func NewDeleteRemoteUpdateConfigsResponse() (response *DeleteRemoteUpdateConfigsResponse) {
	response = &DeleteRemoteUpdateConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除远程更新配置
func (c *Client) DeleteRemoteUpdateConfigs(request *DeleteRemoteUpdateConfigsRequest) (response *DeleteRemoteUpdateConfigsResponse, err error) {
	if request == nil {
		request = NewDeleteRemoteUpdateConfigsRequest()
	}
	response = NewDeleteRemoteUpdateConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDefenseVulListRequest() (request *DescribeDefenseVulListRequest) {
	request = &DescribeDefenseVulListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeDefenseVulList")
	return
}

func NewDescribeDefenseVulListResponse() (response *DescribeDefenseVulListResponse) {
	response = &DescribeDefenseVulListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞管理-漏洞列表
func (c *Client) DescribeDefenseVulList(request *DescribeDefenseVulListRequest) (response *DescribeDefenseVulListResponse, err error) {
	if request == nil {
		request = NewDescribeDefenseVulListRequest()
	}
	response = NewDescribeDefenseVulListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDNSEventDetailRequest() (request *DescribeDNSEventDetailRequest) {
	request = &DescribeDNSEventDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeDNSEventDetail")
	return
}

func NewDescribeDNSEventDetailResponse() (response *DescribeDNSEventDetailResponse) {
	response = &DescribeDNSEventDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恶意请求详情
func (c *Client) DescribeDNSEventDetail(request *DescribeDNSEventDetailRequest) (response *DescribeDNSEventDetailResponse, err error) {
	if request == nil {
		request = NewDescribeDNSEventDetailRequest()
	}
	response = NewDescribeDNSEventDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUpdateConfigRequest() (request *DescribeUpdateConfigRequest) {
	request = &DescribeUpdateConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeUpdateConfig")
	return
}

func NewDescribeUpdateConfigResponse() (response *DescribeUpdateConfigResponse) {
	response = &DescribeUpdateConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 获取版本配置列表
func (c *Client) DescribeUpdateConfig(request *DescribeUpdateConfigRequest) (response *DescribeUpdateConfigResponse, err error) {
	if request == nil {
		request = NewDescribeUpdateConfigRequest()
	}
	response = NewDescribeUpdateConfigResponse()
	err = c.Send(request, response)
	return
}

func NewExportReverseShellEventsRequest() (request *ExportReverseShellEventsRequest) {
	request = &ExportReverseShellEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportReverseShellEvents")
	return
}

func NewExportReverseShellEventsResponse() (response *ExportReverseShellEventsResponse) {
	response = &ExportReverseShellEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出反弹shell列表
func (c *Client) ExportReverseShellEvents(request *ExportReverseShellEventsRequest) (response *ExportReverseShellEventsResponse, err error) {
	if request == nil {
		request = NewExportReverseShellEventsRequest()
	}
	response = NewExportReverseShellEventsResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetDatabaseListRequest() (request *ExportAssetDatabaseListRequest) {
	request = &ExportAssetDatabaseListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAssetDatabaseList")
	return
}

func NewExportAssetDatabaseListResponse() (response *ExportAssetDatabaseListResponse) {
	response = &ExportAssetDatabaseListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理数据库列表
func (c *Client) ExportAssetDatabaseList(request *ExportAssetDatabaseListRequest) (response *ExportAssetDatabaseListResponse, err error) {
	if request == nil {
		request = NewExportAssetDatabaseListRequest()
	}
	response = NewExportAssetDatabaseListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVulDefenceSettingRequest() (request *ModifyVulDefenceSettingRequest) {
	request = &ModifyVulDefenceSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyVulDefenceSetting")
	return
}

func NewModifyVulDefenceSettingResponse() (response *ModifyVulDefenceSettingResponse) {
	response = &ModifyVulDefenceSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改漏洞防御开关
func (c *Client) ModifyVulDefenceSetting(request *ModifyVulDefenceSettingRequest) (response *ModifyVulDefenceSettingResponse, err error) {
	if request == nil {
		request = NewModifyVulDefenceSettingRequest()
	}
	response = NewModifyVulDefenceSettingResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSocTamperProofInfoRequest() (request *UpdateSocTamperProofInfoRequest) {
	request = &UpdateSocTamperProofInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "UpdateSocTamperProofInfo")
	return
}

func NewUpdateSocTamperProofInfoResponse() (response *UpdateSocTamperProofInfoResponse) {
	response = &UpdateSocTamperProofInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 网页防篡改更新状态
func (c *Client) UpdateSocTamperProofInfo(request *UpdateSocTamperProofInfoRequest) (response *UpdateSocTamperProofInfoResponse, err error) {
	if request == nil {
		request = NewUpdateSocTamperProofInfoRequest()
	}
	response = NewUpdateSocTamperProofInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUpdateGrayDetailRequest() (request *DescribeUpdateGrayDetailRequest) {
	request = &DescribeUpdateGrayDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeUpdateGrayDetail")
	return
}

func NewDescribeUpdateGrayDetailResponse() (response *DescribeUpdateGrayDetailResponse) {
	response = &DescribeUpdateGrayDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 获取版本灰度详情
func (c *Client) DescribeUpdateGrayDetail(request *DescribeUpdateGrayDetailRequest) (response *DescribeUpdateGrayDetailResponse, err error) {
	if request == nil {
		request = NewDescribeUpdateGrayDetailRequest()
	}
	response = NewDescribeUpdateGrayDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemoteUpdateUserAccountUniqueRequest() (request *DescribeRemoteUpdateUserAccountUniqueRequest) {
	request = &DescribeRemoteUpdateUserAccountUniqueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeRemoteUpdateUserAccountUnique")
	return
}

func NewDescribeRemoteUpdateUserAccountUniqueResponse() (response *DescribeRemoteUpdateUserAccountUniqueResponse) {
	response = &DescribeRemoteUpdateUserAccountUniqueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询远程更新用户账号唯一性
func (c *Client) DescribeRemoteUpdateUserAccountUnique(request *DescribeRemoteUpdateUserAccountUniqueRequest) (response *DescribeRemoteUpdateUserAccountUniqueResponse, err error) {
	if request == nil {
		request = NewDescribeRemoteUpdateUserAccountUniqueRequest()
	}
	response = NewDescribeRemoteUpdateUserAccountUniqueResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetEnvListRequest() (request *DescribeAssetEnvListRequest) {
	request = &DescribeAssetEnvListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetEnvList")
	return
}

func NewDescribeAssetEnvListResponse() (response *DescribeAssetEnvListResponse) {
	response = &DescribeAssetEnvListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取环境变量列表
func (c *Client) DescribeAssetEnvList(request *DescribeAssetEnvListRequest) (response *DescribeAssetEnvListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetEnvListRequest()
	}
	response = NewDescribeAssetEnvListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCosTempCredentialRequest() (request *DescribeCosTempCredentialRequest) {
	request = &DescribeCosTempCredentialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeCosTempCredential")
	return
}

func NewDescribeCosTempCredentialResponse() (response *DescribeCosTempCredentialResponse) {
	response = &DescribeCosTempCredentialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取COS临时密钥
func (c *Client) DescribeCosTempCredential(request *DescribeCosTempCredentialRequest) (response *DescribeCosTempCredentialResponse, err error) {
	if request == nil {
		request = NewDescribeCosTempCredentialRequest()
	}
	response = NewDescribeCosTempCredentialResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemoteUpdateUploadTempTokenRequest() (request *DescribeRemoteUpdateUploadTempTokenRequest) {
	request = &DescribeRemoteUpdateUploadTempTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeRemoteUpdateUploadTempToken")
	return
}

func NewDescribeRemoteUpdateUploadTempTokenResponse() (response *DescribeRemoteUpdateUploadTempTokenResponse) {
	response = &DescribeRemoteUpdateUploadTempTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取上传远程更新包的临时Token
func (c *Client) DescribeRemoteUpdateUploadTempToken(request *DescribeRemoteUpdateUploadTempTokenRequest) (response *DescribeRemoteUpdateUploadTempTokenResponse, err error) {
	if request == nil {
		request = NewDescribeRemoteUpdateUploadTempTokenRequest()
	}
	response = NewDescribeRemoteUpdateUploadTempTokenResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUpdateTaskLogRequest() (request *DescribeUpdateTaskLogRequest) {
	request = &DescribeUpdateTaskLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeUpdateTaskLog")
	return
}

func NewDescribeUpdateTaskLogResponse() (response *DescribeUpdateTaskLogResponse) {
	response = &DescribeUpdateTaskLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 获取版本任务日志列表
func (c *Client) DescribeUpdateTaskLog(request *DescribeUpdateTaskLogRequest) (response *DescribeUpdateTaskLogResponse, err error) {
	if request == nil {
		request = NewDescribeUpdateTaskLogRequest()
	}
	response = NewDescribeUpdateTaskLogResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRemoteUpdateTagRequest() (request *CreateRemoteUpdateTagRequest) {
	request = &CreateRemoteUpdateTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "CreateRemoteUpdateTag")
	return
}

func NewCreateRemoteUpdateTagResponse() (response *CreateRemoteUpdateTagResponse) {
	response = &CreateRemoteUpdateTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加远程更新标签
func (c *Client) CreateRemoteUpdateTag(request *CreateRemoteUpdateTagRequest) (response *CreateRemoteUpdateTagResponse, err error) {
	if request == nil {
		request = NewCreateRemoteUpdateTagRequest()
	}
	response = NewCreateRemoteUpdateTagResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVulVulsRequest() (request *DeleteVulVulsRequest) {
	request = &DeleteVulVulsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteVulVuls")
	return
}

func NewDeleteVulVulsResponse() (response *DeleteVulVulsResponse) {
	response = &DeleteVulVulsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞管理-删除漏洞
func (c *Client) DeleteVulVuls(request *DeleteVulVulsRequest) (response *DeleteVulVulsResponse, err error) {
	if request == nil {
		request = NewDeleteVulVulsRequest()
	}
	response = NewDeleteVulVulsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemoteUpdateConfigNameListRequest() (request *DescribeRemoteUpdateConfigNameListRequest) {
	request = &DescribeRemoteUpdateConfigNameListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeRemoteUpdateConfigNameList")
	return
}

func NewDescribeRemoteUpdateConfigNameListResponse() (response *DescribeRemoteUpdateConfigNameListResponse) {
	response = &DescribeRemoteUpdateConfigNameListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询远程更新配置名称列表
func (c *Client) DescribeRemoteUpdateConfigNameList(request *DescribeRemoteUpdateConfigNameListRequest) (response *DescribeRemoteUpdateConfigNameListResponse, err error) {
	if request == nil {
		request = NewDescribeRemoteUpdateConfigNameListRequest()
	}
	response = NewDescribeRemoteUpdateConfigNameListResponse()
	err = c.Send(request, response)
	return
}

func NewStopUpdateTaskRequest() (request *StopUpdateTaskRequest) {
	request = &StopUpdateTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "StopUpdateTask")
	return
}

func NewStopUpdateTaskResponse() (response *StopUpdateTaskResponse) {
	response = &StopUpdateTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 停止升级任务
func (c *Client) StopUpdateTask(request *StopUpdateTaskRequest) (response *StopUpdateTaskResponse, err error) {
	if request == nil {
		request = NewStopUpdateTaskRequest()
	}
	response = NewStopUpdateTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetTotalCountRequest() (request *DescribeAssetTotalCountRequest) {
	request = &DescribeAssetTotalCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetTotalCount")
	return
}

func NewDescribeAssetTotalCountResponse() (response *DescribeAssetTotalCountResponse) {
	response = &DescribeAssetTotalCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主机各类型资源数量
func (c *Client) DescribeAssetTotalCount(request *DescribeAssetTotalCountRequest) (response *DescribeAssetTotalCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetTotalCountRequest()
	}
	response = NewDescribeAssetTotalCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomerLicenseDetailsRequest() (request *DescribeCustomerLicenseDetailsRequest) {
	request = &DescribeCustomerLicenseDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeCustomerLicenseDetails")
	return
}

func NewDescribeCustomerLicenseDetailsResponse() (response *DescribeCustomerLicenseDetailsResponse) {
	response = &DescribeCustomerLicenseDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询License详情
func (c *Client) DescribeCustomerLicenseDetails(request *DescribeCustomerLicenseDetailsRequest) (response *DescribeCustomerLicenseDetailsResponse, err error) {
	if request == nil {
		request = NewDescribeCustomerLicenseDetailsRequest()
	}
	response = NewDescribeCustomerLicenseDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebFrameCountRequest() (request *DescribeAssetWebFrameCountRequest) {
	request = &DescribeAssetWebFrameCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetWebFrameCount")
	return
}

func NewDescribeAssetWebFrameCountResponse() (response *DescribeAssetWebFrameCountResponse) {
	response = &DescribeAssetWebFrameCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有Web框架数量
func (c *Client) DescribeAssetWebFrameCount(request *DescribeAssetWebFrameCountRequest) (response *DescribeAssetWebFrameCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebFrameCountRequest()
	}
	response = NewDescribeAssetWebFrameCountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVulDefenceEventStatusRequest() (request *ModifyVulDefenceEventStatusRequest) {
	request = &ModifyVulDefenceEventStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyVulDefenceEventStatus")
	return
}

func NewModifyVulDefenceEventStatusResponse() (response *ModifyVulDefenceEventStatusResponse) {
	response = &ModifyVulDefenceEventStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改漏洞防御事件状态
func (c *Client) ModifyVulDefenceEventStatus(request *ModifyVulDefenceEventStatusRequest) (response *ModifyVulDefenceEventStatusResponse, err error) {
	if request == nil {
		request = NewModifyVulDefenceEventStatusRequest()
	}
	response = NewModifyVulDefenceEventStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpecimenDetailRequest() (request *DescribeSpecimenDetailRequest) {
	request = &DescribeSpecimenDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeSpecimenDetail")
	return
}

func NewDescribeSpecimenDetailResponse() (response *DescribeSpecimenDetailResponse) {
	response = &DescribeSpecimenDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 文件查杀 - 列表 - 样本详情
func (c *Client) DescribeSpecimenDetail(request *DescribeSpecimenDetailRequest) (response *DescribeSpecimenDetailResponse, err error) {
	if request == nil {
		request = NewDescribeSpecimenDetailRequest()
	}
	response = NewDescribeSpecimenDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulVulsListRequest() (request *DescribeVulVulsListRequest) {
	request = &DescribeVulVulsListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeVulVulsList")
	return
}

func NewDescribeVulVulsListResponse() (response *DescribeVulVulsListResponse) {
	response = &DescribeVulVulsListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞管理-漏洞列表
func (c *Client) DescribeVulVulsList(request *DescribeVulVulsListRequest) (response *DescribeVulVulsListResponse, err error) {
	if request == nil {
		request = NewDescribeVulVulsListRequest()
	}
	response = NewDescribeVulVulsListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyStatisticRequest() (request *ModifyStatisticRequest) {
	request = &ModifyStatisticRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyStatistic")
	return
}

func NewModifyStatisticResponse() (response *ModifyStatisticResponse) {
	response = &ModifyStatisticResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新数据统计
func (c *Client) ModifyStatistic(request *ModifyStatisticRequest) (response *ModifyStatisticResponse, err error) {
	if request == nil {
		request = NewModifyStatisticRequest()
	}
	response = NewModifyStatisticResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteZKConfigInfoRequest() (request *DeleteZKConfigInfoRequest) {
	request = &DeleteZKConfigInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteZKConfigInfo")
	return
}

func NewDeleteZKConfigInfoResponse() (response *DeleteZKConfigInfoResponse) {
	response = &DeleteZKConfigInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 系统管理 - Zookeeper配置 - 查看 - 删除配置
func (c *Client) DeleteZKConfigInfo(request *DeleteZKConfigInfoRequest) (response *DeleteZKConfigInfoResponse, err error) {
	if request == nil {
		request = NewDeleteZKConfigInfoRequest()
	}
	response = NewDeleteZKConfigInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyJavaMemShellsStatusRequest() (request *ModifyJavaMemShellsStatusRequest) {
	request = &ModifyJavaMemShellsStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyJavaMemShellsStatus")
	return
}

func NewModifyJavaMemShellsStatusResponse() (response *ModifyJavaMemShellsStatusResponse) {
	response = &ModifyJavaMemShellsStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改内存马事件状态
func (c *Client) ModifyJavaMemShellsStatus(request *ModifyJavaMemShellsStatusRequest) (response *ModifyJavaMemShellsStatusResponse, err error) {
	if request == nil {
		request = NewModifyJavaMemShellsStatusRequest()
	}
	response = NewModifyJavaMemShellsStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSocTamperProofListRequest() (request *DescribeSocTamperProofListRequest) {
	request = &DescribeSocTamperProofListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeSocTamperProofList")
	return
}

func NewDescribeSocTamperProofListResponse() (response *DescribeSocTamperProofListResponse) {
	response = &DescribeSocTamperProofListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 网页防篡改列表
func (c *Client) DescribeSocTamperProofList(request *DescribeSocTamperProofListRequest) (response *DescribeSocTamperProofListResponse, err error) {
	if request == nil {
		request = NewDescribeSocTamperProofListRequest()
	}
	response = NewDescribeSocTamperProofListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUpdateTaskDetailRequest() (request *DescribeUpdateTaskDetailRequest) {
	request = &DescribeUpdateTaskDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeUpdateTaskDetail")
	return
}

func NewDescribeUpdateTaskDetailResponse() (response *DescribeUpdateTaskDetailResponse) {
	response = &DescribeUpdateTaskDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 获取版本任务详情
func (c *Client) DescribeUpdateTaskDetail(request *DescribeUpdateTaskDetailRequest) (response *DescribeUpdateTaskDetailResponse, err error) {
	if request == nil {
		request = NewDescribeUpdateTaskDetailRequest()
	}
	response = NewDescribeUpdateTaskDetailResponse()
	err = c.Send(request, response)
	return
}

func NewExportVulDefencePluginListRequest() (request *ExportVulDefencePluginListRequest) {
	request = &ExportVulDefencePluginListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportVulDefencePluginList")
	return
}

func NewExportVulDefencePluginListResponse() (response *ExportVulDefencePluginListResponse) {
	response = &ExportVulDefencePluginListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出插件详情列表
func (c *Client) ExportVulDefencePluginList(request *ExportVulDefencePluginListRequest) (response *ExportVulDefencePluginListResponse, err error) {
	if request == nil {
		request = NewExportVulDefencePluginListRequest()
	}
	response = NewExportVulDefencePluginListResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadMalwareFileRequest() (request *DownloadMalwareFileRequest) {
	request = &DownloadMalwareFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DownloadMalwareFile")
	return
}

func NewDownloadMalwareFileResponse() (response *DownloadMalwareFileResponse) {
	response = &DownloadMalwareFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下载木马文件
func (c *Client) DownloadMalwareFile(request *DownloadMalwareFileRequest) (response *DownloadMalwareFileResponse, err error) {
	if request == nil {
		request = NewDownloadMalwareFileRequest()
	}
	response = NewDownloadMalwareFileResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMalwareEventsRequest() (request *DescribeMalwareEventsRequest) {
	request = &DescribeMalwareEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeMalwareEvents")
	return
}

func NewDescribeMalwareEventsResponse() (response *DescribeMalwareEventsResponse) {
	response = &DescribeMalwareEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 文件查杀 - 列表
func (c *Client) DescribeMalwareEvents(request *DescribeMalwareEventsRequest) (response *DescribeMalwareEventsResponse, err error) {
	if request == nil {
		request = NewDescribeMalwareEventsRequest()
	}
	response = NewDescribeMalwareEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskComponentsRequest() (request *DescribeTaskComponentsRequest) {
	request = &DescribeTaskComponentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeTaskComponents")
	return
}

func NewDescribeTaskComponentsResponse() (response *DescribeTaskComponentsResponse) {
	response = &DescribeTaskComponentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 组件管理列表
func (c *Client) DescribeTaskComponents(request *DescribeTaskComponentsRequest) (response *DescribeTaskComponentsResponse, err error) {
	if request == nil {
		request = NewDescribeTaskComponentsRequest()
	}
	response = NewDescribeTaskComponentsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetJarInfoRequest() (request *DescribeAssetJarInfoRequest) {
	request = &DescribeAssetJarInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetJarInfo")
	return
}

func NewDescribeAssetJarInfoResponse() (response *DescribeAssetJarInfoResponse) {
	response = &DescribeAssetJarInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Jar详情
func (c *Client) DescribeAssetJarInfo(request *DescribeAssetJarInfoRequest) (response *DescribeAssetJarInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetJarInfoRequest()
	}
	response = NewDescribeAssetJarInfoResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetSystemPackageListRequest() (request *ExportAssetSystemPackageListRequest) {
	request = &ExportAssetSystemPackageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAssetSystemPackageList")
	return
}

func NewExportAssetSystemPackageListResponse() (response *ExportAssetSystemPackageListResponse) {
	response = &ExportAssetSystemPackageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// yunapi response
func (c *Client) ExportAssetSystemPackageList(request *ExportAssetSystemPackageListRequest) (response *ExportAssetSystemPackageListResponse, err error) {
	if request == nil {
		request = NewExportAssetSystemPackageListRequest()
	}
	response = NewExportAssetSystemPackageListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFileSpecimenListRequest() (request *DescribeFileSpecimenListRequest) {
	request = &DescribeFileSpecimenListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeFileSpecimenList")
	return
}

func NewDescribeFileSpecimenListResponse() (response *DescribeFileSpecimenListResponse) {
	response = &DescribeFileSpecimenListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全策略 - 样本管理 - 列表
func (c *Client) DescribeFileSpecimenList(request *DescribeFileSpecimenListRequest) (response *DescribeFileSpecimenListResponse, err error) {
	if request == nil {
		request = NewDescribeFileSpecimenListRequest()
	}
	response = NewDescribeFileSpecimenListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemoteUpdateUserRequest() (request *DescribeRemoteUpdateUserRequest) {
	request = &DescribeRemoteUpdateUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeRemoteUpdateUser")
	return
}

func NewDescribeRemoteUpdateUserResponse() (response *DescribeRemoteUpdateUserResponse) {
	response = &DescribeRemoteUpdateUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取远程更新用户信息
func (c *Client) DescribeRemoteUpdateUser(request *DescribeRemoteUpdateUserRequest) (response *DescribeRemoteUpdateUserResponse, err error) {
	if request == nil {
		request = NewDescribeRemoteUpdateUserRequest()
	}
	response = NewDescribeRemoteUpdateUserResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUpdateTaskConfigListRequest() (request *DescribeUpdateTaskConfigListRequest) {
	request = &DescribeUpdateTaskConfigListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeUpdateTaskConfigList")
	return
}

func NewDescribeUpdateTaskConfigListResponse() (response *DescribeUpdateTaskConfigListResponse) {
	response = &DescribeUpdateTaskConfigListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 获取版本配置列表
func (c *Client) DescribeUpdateTaskConfigList(request *DescribeUpdateTaskConfigListRequest) (response *DescribeUpdateTaskConfigListResponse, err error) {
	if request == nil {
		request = NewDescribeUpdateTaskConfigListRequest()
	}
	response = NewDescribeUpdateTaskConfigListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebAppListRequest() (request *DescribeAssetWebAppListRequest) {
	request = &DescribeAssetWebAppListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetWebAppList")
	return
}

func NewDescribeAssetWebAppListResponse() (response *DescribeAssetWebAppListResponse) {
	response = &DescribeAssetWebAppListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Web应用列表
func (c *Client) DescribeAssetWebAppList(request *DescribeAssetWebAppListRequest) (response *DescribeAssetWebAppListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebAppListRequest()
	}
	response = NewDescribeAssetWebAppListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseRequest() (request *DescribeLicenseRequest) {
	request = &DescribeLicenseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeLicense")
	return
}

func NewDescribeLicenseResponse() (response *DescribeLicenseResponse) {
	response = &DescribeLicenseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取当前license信息
func (c *Client) DescribeLicense(request *DescribeLicenseRequest) (response *DescribeLicenseResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseRequest()
	}
	response = NewDescribeLicenseResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulScanTaskListRequest() (request *DescribeVulScanTaskListRequest) {
	request = &DescribeVulScanTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeVulScanTaskList")
	return
}

func NewDescribeVulScanTaskListResponse() (response *DescribeVulScanTaskListResponse) {
	response = &DescribeVulScanTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 任务配置 - 任务管理 -列表
func (c *Client) DescribeVulScanTaskList(request *DescribeVulScanTaskListRequest) (response *DescribeVulScanTaskListResponse, err error) {
	if request == nil {
		request = NewDescribeVulScanTaskListRequest()
	}
	response = NewDescribeVulScanTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewExportJavaMemShellPluginsInfoRequest() (request *ExportJavaMemShellPluginsInfoRequest) {
	request = &ExportJavaMemShellPluginsInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportJavaMemShellPluginsInfo")
	return
}

func NewExportJavaMemShellPluginsInfoResponse() (response *ExportJavaMemShellPluginsInfoResponse) {
	response = &ExportJavaMemShellPluginsInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 内存马列表 - 详情 - 导出
func (c *Client) ExportJavaMemShellPluginsInfo(request *ExportJavaMemShellPluginsInfoRequest) (response *ExportJavaMemShellPluginsInfoResponse, err error) {
	if request == nil {
		request = NewExportJavaMemShellPluginsInfoRequest()
	}
	response = NewExportJavaMemShellPluginsInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineRuleListRequest() (request *DescribeBaselineRuleListRequest) {
	request = &DescribeBaselineRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeBaselineRuleList")
	return
}

func NewDescribeBaselineRuleListResponse() (response *DescribeBaselineRuleListResponse) {
	response = &DescribeBaselineRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线规则列表
func (c *Client) DescribeBaselineRuleList(request *DescribeBaselineRuleListRequest) (response *DescribeBaselineRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineRuleListRequest()
	}
	response = NewDescribeBaselineRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteReverseShellEventRequest() (request *DeleteReverseShellEventRequest) {
	request = &DeleteReverseShellEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteReverseShellEvent")
	return
}

func NewDeleteReverseShellEventResponse() (response *DeleteReverseShellEventResponse) {
	response = &DeleteReverseShellEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除反弹shell
func (c *Client) DeleteReverseShellEvent(request *DeleteReverseShellEventRequest) (response *DeleteReverseShellEventResponse, err error) {
	if request == nil {
		request = NewDeleteReverseShellEventRequest()
	}
	response = NewDeleteReverseShellEventResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVulDefenceEventRequest() (request *DeleteVulDefenceEventRequest) {
	request = &DeleteVulDefenceEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteVulDefenceEvent")
	return
}

func NewDeleteVulDefenceEventResponse() (response *DeleteVulDefenceEventResponse) {
	response = &DeleteVulDefenceEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除漏洞防御事件
func (c *Client) DeleteVulDefenceEvent(request *DeleteVulDefenceEventRequest) (response *DeleteVulDefenceEventResponse, err error) {
	if request == nil {
		request = NewDeleteVulDefenceEventRequest()
	}
	response = NewDeleteVulDefenceEventResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebFrameListRequest() (request *DescribeAssetWebFrameListRequest) {
	request = &DescribeAssetWebFrameListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetWebFrameList")
	return
}

func NewDescribeAssetWebFrameListResponse() (response *DescribeAssetWebFrameListResponse) {
	response = &DescribeAssetWebFrameListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Web框架列表
func (c *Client) DescribeAssetWebFrameList(request *DescribeAssetWebFrameListRequest) (response *DescribeAssetWebFrameListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebFrameListRequest()
	}
	response = NewDescribeAssetWebFrameListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTrojanSettingRequest() (request *DescribeTrojanSettingRequest) {
	request = &DescribeTrojanSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeTrojanSetting")
	return
}

func NewDescribeTrojanSettingResponse() (response *DescribeTrojanSettingResponse) {
	response = &DescribeTrojanSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全策略 - 木马隔离 - 详情
func (c *Client) DescribeTrojanSetting(request *DescribeTrojanSettingRequest) (response *DescribeTrojanSettingResponse, err error) {
	if request == nil {
		request = NewDescribeTrojanSettingRequest()
	}
	response = NewDescribeTrojanSettingResponse()
	err = c.Send(request, response)
	return
}

func NewExportJavaMemShellsRequest() (request *ExportJavaMemShellsRequest) {
	request = &ExportJavaMemShellsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportJavaMemShells")
	return
}

func NewExportJavaMemShellsResponse() (response *ExportJavaMemShellsResponse) {
	response = &ExportJavaMemShellsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出java内存马列表
func (c *Client) ExportJavaMemShells(request *ExportJavaMemShellsRequest) (response *ExportJavaMemShellsResponse, err error) {
	if request == nil {
		request = NewExportJavaMemShellsRequest()
	}
	response = NewExportJavaMemShellsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePrivilegeEventsRequest() (request *DescribePrivilegeEventsRequest) {
	request = &DescribePrivilegeEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribePrivilegeEvents")
	return
}

func NewDescribePrivilegeEventsResponse() (response *DescribePrivilegeEventsResponse) {
	response = &DescribePrivilegeEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本地提权列表
func (c *Client) DescribePrivilegeEvents(request *DescribePrivilegeEventsRequest) (response *DescribePrivilegeEventsResponse, err error) {
	if request == nil {
		request = NewDescribePrivilegeEventsRequest()
	}
	response = NewDescribePrivilegeEventsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVulVulsInfoRequest() (request *ModifyVulVulsInfoRequest) {
	request = &ModifyVulVulsInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyVulVulsInfo")
	return
}

func NewModifyVulVulsInfoResponse() (response *ModifyVulVulsInfoResponse) {
	response = &ModifyVulVulsInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞管理-修改或添加漏洞
func (c *Client) ModifyVulVulsInfo(request *ModifyVulVulsInfoRequest) (response *ModifyVulVulsInfoResponse, err error) {
	if request == nil {
		request = NewModifyVulVulsInfoRequest()
	}
	response = NewModifyVulVulsInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRemoteUpdateUserTagRequest() (request *ModifyRemoteUpdateUserTagRequest) {
	request = &ModifyRemoteUpdateUserTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyRemoteUpdateUserTag")
	return
}

func NewModifyRemoteUpdateUserTagResponse() (response *ModifyRemoteUpdateUserTagResponse) {
	response = &ModifyRemoteUpdateUserTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改远程更新用户的标签
func (c *Client) ModifyRemoteUpdateUserTag(request *ModifyRemoteUpdateUserTagRequest) (response *ModifyRemoteUpdateUserTagResponse, err error) {
	if request == nil {
		request = NewModifyRemoteUpdateUserTagRequest()
	}
	response = NewModifyRemoteUpdateUserTagResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHostTrendRequest() (request *DescribeHostTrendRequest) {
	request = &DescribeHostTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeHostTrend")
	return
}

func NewDescribeHostTrendResponse() (response *DescribeHostTrendResponse) {
	response = &DescribeHostTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查找服务器在线趋势
func (c *Client) DescribeHostTrend(request *DescribeHostTrendRequest) (response *DescribeHostTrendResponse, err error) {
	if request == nil {
		request = NewDescribeHostTrendRequest()
	}
	response = NewDescribeHostTrendResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellEventsRequest() (request *DescribeReverseShellEventsRequest) {
	request = &DescribeReverseShellEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeReverseShellEvents")
	return
}

func NewDescribeReverseShellEventsResponse() (response *DescribeReverseShellEventsResponse) {
	response = &DescribeReverseShellEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 反弹shell - 列表
func (c *Client) DescribeReverseShellEvents(request *DescribeReverseShellEventsRequest) (response *DescribeReverseShellEventsResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellEventsRequest()
	}
	response = NewDescribeReverseShellEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUpdateGrayRequest() (request *DescribeUpdateGrayRequest) {
	request = &DescribeUpdateGrayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeUpdateGray")
	return
}

func NewDescribeUpdateGrayResponse() (response *DescribeUpdateGrayResponse) {
	response = &DescribeUpdateGrayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 获取发布目标列表
func (c *Client) DescribeUpdateGray(request *DescribeUpdateGrayRequest) (response *DescribeUpdateGrayResponse, err error) {
	if request == nil {
		request = NewDescribeUpdateGrayRequest()
	}
	response = NewDescribeUpdateGrayResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePrivilegeEventRequest() (request *DeletePrivilegeEventRequest) {
	request = &DeletePrivilegeEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeletePrivilegeEvent")
	return
}

func NewDeletePrivilegeEventResponse() (response *DeletePrivilegeEventResponse) {
	response = &DeletePrivilegeEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除本地提权
func (c *Client) DeletePrivilegeEvent(request *DeletePrivilegeEventRequest) (response *DeletePrivilegeEventResponse, err error) {
	if request == nil {
		request = NewDeletePrivilegeEventRequest()
	}
	response = NewDeletePrivilegeEventResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUpdateConfigRequest() (request *DeleteUpdateConfigRequest) {
	request = &DeleteUpdateConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteUpdateConfig")
	return
}

func NewDeleteUpdateConfigResponse() (response *DeleteUpdateConfigResponse) {
	response = &DeleteUpdateConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 删除配置
func (c *Client) DeleteUpdateConfig(request *DeleteUpdateConfigRequest) (response *DeleteUpdateConfigResponse, err error) {
	if request == nil {
		request = NewDeleteUpdateConfigRequest()
	}
	response = NewDeleteUpdateConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVulTypeRequest() (request *DeleteVulTypeRequest) {
	request = &DeleteVulTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteVulType")
	return
}

func NewDeleteVulTypeResponse() (response *DeleteVulTypeResponse) {
	response = &DeleteVulTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除漏洞类型
func (c *Client) DeleteVulType(request *DeleteVulTypeRequest) (response *DeleteVulTypeResponse, err error) {
	if request == nil {
		request = NewDeleteVulTypeRequest()
	}
	response = NewDeleteVulTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetPortCountRequest() (request *DescribeAssetPortCountRequest) {
	request = &DescribeAssetPortCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetPortCount")
	return
}

func NewDescribeAssetPortCountResponse() (response *DescribeAssetPortCountResponse) {
	response = &DescribeAssetPortCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有端口数量
func (c *Client) DescribeAssetPortCount(request *DescribeAssetPortCountRequest) (response *DescribeAssetPortCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetPortCountRequest()
	}
	response = NewDescribeAssetPortCountResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRemoteUpdateConfigRequest() (request *CreateRemoteUpdateConfigRequest) {
	request = &CreateRemoteUpdateConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "CreateRemoteUpdateConfig")
	return
}

func NewCreateRemoteUpdateConfigResponse() (response *CreateRemoteUpdateConfigResponse) {
	response = &CreateRemoteUpdateConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建远程更新配置
func (c *Client) CreateRemoteUpdateConfig(request *CreateRemoteUpdateConfigRequest) (response *CreateRemoteUpdateConfigResponse, err error) {
	if request == nil {
		request = NewCreateRemoteUpdateConfigRequest()
	}
	response = NewCreateRemoteUpdateConfigResponse()
	err = c.Send(request, response)
	return
}

func NewExportAgentPushTaskResultRequest() (request *ExportAgentPushTaskResultRequest) {
	request = &ExportAgentPushTaskResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAgentPushTaskResult")
	return
}

func NewExportAgentPushTaskResultResponse() (response *ExportAgentPushTaskResultResponse) {
	response = &ExportAgentPushTaskResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 命令推送 - 列表 - 详情 - 导出
func (c *Client) ExportAgentPushTaskResult(request *ExportAgentPushTaskResultRequest) (response *ExportAgentPushTaskResultResponse, err error) {
	if request == nil {
		request = NewExportAgentPushTaskResultRequest()
	}
	response = NewExportAgentPushTaskResultResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemoteUpdateConfigAutoUpdateRequest() (request *DescribeRemoteUpdateConfigAutoUpdateRequest) {
	request = &DescribeRemoteUpdateConfigAutoUpdateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeRemoteUpdateConfigAutoUpdate")
	return
}

func NewDescribeRemoteUpdateConfigAutoUpdateResponse() (response *DescribeRemoteUpdateConfigAutoUpdateResponse) {
	response = &DescribeRemoteUpdateConfigAutoUpdateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某一个用户的某一类配置自动升级开关的状态
func (c *Client) DescribeRemoteUpdateConfigAutoUpdate(request *DescribeRemoteUpdateConfigAutoUpdateRequest) (response *DescribeRemoteUpdateConfigAutoUpdateResponse, err error) {
	if request == nil {
		request = NewDescribeRemoteUpdateConfigAutoUpdateRequest()
	}
	response = NewDescribeRemoteUpdateConfigAutoUpdateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseAgentListRequest() (request *DescribeLicenseAgentListRequest) {
	request = &DescribeLicenseAgentListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeLicenseAgentList")
	return
}

func NewDescribeLicenseAgentListResponse() (response *DescribeLicenseAgentListResponse) {
	response = &DescribeLicenseAgentListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查找agent列表
func (c *Client) DescribeLicenseAgentList(request *DescribeLicenseAgentListRequest) (response *DescribeLicenseAgentListResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseAgentListRequest()
	}
	response = NewDescribeLicenseAgentListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLicenseInfoRequest() (request *CreateLicenseInfoRequest) {
	request = &CreateLicenseInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "CreateLicenseInfo")
	return
}

func NewCreateLicenseInfoResponse() (response *CreateLicenseInfoResponse) {
	response = &CreateLicenseInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增License
func (c *Client) CreateLicenseInfo(request *CreateLicenseInfoRequest) (response *CreateLicenseInfoResponse, err error) {
	if request == nil {
		request = NewCreateLicenseInfoRequest()
	}
	response = NewCreateLicenseInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUpdateTaskListRequest() (request *DescribeUpdateTaskListRequest) {
	request = &DescribeUpdateTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeUpdateTaskList")
	return
}

func NewDescribeUpdateTaskListResponse() (response *DescribeUpdateTaskListResponse) {
	response = &DescribeUpdateTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理- 列表
func (c *Client) DescribeUpdateTaskList(request *DescribeUpdateTaskListRequest) (response *DescribeUpdateTaskListResponse, err error) {
	if request == nil {
		request = NewDescribeUpdateTaskListRequest()
	}
	response = NewDescribeUpdateTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJavaMemShellPluginInfoRequest() (request *DescribeJavaMemShellPluginInfoRequest) {
	request = &DescribeJavaMemShellPluginInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeJavaMemShellPluginInfo")
	return
}

func NewDescribeJavaMemShellPluginInfoResponse() (response *DescribeJavaMemShellPluginInfoResponse) {
	response = &DescribeJavaMemShellPluginInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询给定主机java内存马插件信息
func (c *Client) DescribeJavaMemShellPluginInfo(request *DescribeJavaMemShellPluginInfoRequest) (response *DescribeJavaMemShellPluginInfoResponse, err error) {
	if request == nil {
		request = NewDescribeJavaMemShellPluginInfoRequest()
	}
	response = NewDescribeJavaMemShellPluginInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVulScanConfigRequest() (request *DeleteVulScanConfigRequest) {
	request = &DeleteVulScanConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteVulScanConfig")
	return
}

func NewDeleteVulScanConfigResponse() (response *DeleteVulScanConfigResponse) {
	response = &DeleteVulScanConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置文件管理 - 删除
func (c *Client) DeleteVulScanConfig(request *DeleteVulScanConfigRequest) (response *DeleteVulScanConfigResponse, err error) {
	if request == nil {
		request = NewDeleteVulScanConfigRequest()
	}
	response = NewDeleteVulScanConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZKConfigListRequest() (request *DescribeZKConfigListRequest) {
	request = &DescribeZKConfigListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeZKConfigList")
	return
}

func NewDescribeZKConfigListResponse() (response *DescribeZKConfigListResponse) {
	response = &DescribeZKConfigListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看Zookeeper配置l列表
func (c *Client) DescribeZKConfigList(request *DescribeZKConfigListRequest) (response *DescribeZKConfigListResponse, err error) {
	if request == nil {
		request = NewDescribeZKConfigListRequest()
	}
	response = NewDescribeZKConfigListResponse()
	err = c.Send(request, response)
	return
}

func NewExportBruteAttackTopRequest() (request *ExportBruteAttackTopRequest) {
	request = &ExportBruteAttackTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportBruteAttackTop")
	return
}

func NewExportBruteAttackTopResponse() (response *ExportBruteAttackTopResponse) {
	response = &ExportBruteAttackTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出获取暴力破解top
func (c *Client) ExportBruteAttackTop(request *ExportBruteAttackTopRequest) (response *ExportBruteAttackTopResponse, err error) {
	if request == nil {
		request = NewExportBruteAttackTopRequest()
	}
	response = NewExportBruteAttackTopResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentPushTaskDetailRequest() (request *DescribeAgentPushTaskDetailRequest) {
	request = &DescribeAgentPushTaskDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAgentPushTaskDetail")
	return
}

func NewDescribeAgentPushTaskDetailResponse() (response *DescribeAgentPushTaskDetailResponse) {
	response = &DescribeAgentPushTaskDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 命令推送 - 列表 - 详情
func (c *Client) DescribeAgentPushTaskDetail(request *DescribeAgentPushTaskDetailRequest) (response *DescribeAgentPushTaskDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAgentPushTaskDetailRequest()
	}
	response = NewDescribeAgentPushTaskDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePushCommandListRequest() (request *DescribePushCommandListRequest) {
	request = &DescribePushCommandListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribePushCommandList")
	return
}

func NewDescribePushCommandListResponse() (response *DescribePushCommandListResponse) {
	response = &DescribePushCommandListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询命令列表
func (c *Client) DescribePushCommandList(request *DescribePushCommandListRequest) (response *DescribePushCommandListResponse, err error) {
	if request == nil {
		request = NewDescribePushCommandListRequest()
	}
	response = NewDescribePushCommandListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskComponentsDetailRequest() (request *DescribeTaskComponentsDetailRequest) {
	request = &DescribeTaskComponentsDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeTaskComponentsDetail")
	return
}

func NewDescribeTaskComponentsDetailResponse() (response *DescribeTaskComponentsDetailResponse) {
	response = &DescribeTaskComponentsDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组件详情
func (c *Client) DescribeTaskComponentsDetail(request *DescribeTaskComponentsDetailRequest) (response *DescribeTaskComponentsDetailResponse, err error) {
	if request == nil {
		request = NewDescribeTaskComponentsDetailRequest()
	}
	response = NewDescribeTaskComponentsDetailResponse()
	err = c.Send(request, response)
	return
}

func NewExportSecurityDynamicsRequest() (request *ExportSecurityDynamicsRequest) {
	request = &ExportSecurityDynamicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportSecurityDynamics")
	return
}

func NewExportSecurityDynamicsResponse() (response *ExportSecurityDynamicsResponse) {
	response = &ExportSecurityDynamicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出安全事件动态消息
func (c *Client) ExportSecurityDynamics(request *ExportSecurityDynamicsRequest) (response *ExportSecurityDynamicsResponse, err error) {
	if request == nil {
		request = NewExportSecurityDynamicsRequest()
	}
	response = NewExportSecurityDynamicsResponse()
	err = c.Send(request, response)
	return
}

func NewExportVulAllEventsRequest() (request *ExportVulAllEventsRequest) {
	request = &ExportVulAllEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportVulAllEvents")
	return
}

func NewExportVulAllEventsResponse() (response *ExportVulAllEventsResponse) {
	response = &ExportVulAllEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出安全漏洞(全量) 列表
func (c *Client) ExportVulAllEvents(request *ExportVulAllEventsRequest) (response *ExportVulAllEventsResponse, err error) {
	if request == nil {
		request = NewExportVulAllEventsRequest()
	}
	response = NewExportVulAllEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseListRequest() (request *DescribeLicenseListRequest) {
	request = &DescribeLicenseListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeLicenseList")
	return
}

func NewDescribeLicenseListResponse() (response *DescribeLicenseListResponse) {
	response = &DescribeLicenseListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查找列表license
func (c *Client) DescribeLicenseList(request *DescribeLicenseListRequest) (response *DescribeLicenseListResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseListRequest()
	}
	response = NewDescribeLicenseListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUpdateTaskRequest() (request *ModifyUpdateTaskRequest) {
	request = &ModifyUpdateTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyUpdateTask")
	return
}

func NewModifyUpdateTaskResponse() (response *ModifyUpdateTaskResponse) {
	response = &ModifyUpdateTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 修改或增加任务
func (c *Client) ModifyUpdateTask(request *ModifyUpdateTaskRequest) (response *ModifyUpdateTaskResponse, err error) {
	if request == nil {
		request = NewModifyUpdateTaskRequest()
	}
	response = NewModifyUpdateTaskResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstallPackageRequest() (request *ModifyInstallPackageRequest) {
	request = &ModifyInstallPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyInstallPackage")
	return
}

func NewModifyInstallPackageResponse() (response *ModifyInstallPackageResponse) {
	response = &ModifyInstallPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改安装包记录信息
func (c *Client) ModifyInstallPackage(request *ModifyInstallPackageRequest) (response *ModifyInstallPackageResponse, err error) {
	if request == nil {
		request = NewModifyInstallPackageRequest()
	}
	response = NewModifyInstallPackageResponse()
	err = c.Send(request, response)
	return
}

func NewExportFileTamperEventsRequest() (request *ExportFileTamperEventsRequest) {
	request = &ExportFileTamperEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportFileTamperEvents")
	return
}

func NewExportFileTamperEventsResponse() (response *ExportFileTamperEventsResponse) {
	response = &ExportFileTamperEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出核心文件事件列表
func (c *Client) ExportFileTamperEvents(request *ExportFileTamperEventsRequest) (response *ExportFileTamperEventsResponse, err error) {
	if request == nil {
		request = NewExportFileTamperEventsRequest()
	}
	response = NewExportFileTamperEventsResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetWebLocationListRequest() (request *ExportAssetWebLocationListRequest) {
	request = &ExportAssetWebLocationListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAssetWebLocationList")
	return
}

func NewExportAssetWebLocationListResponse() (response *ExportAssetWebLocationListResponse) {
	response = &ExportAssetWebLocationListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出Web站点列表
func (c *Client) ExportAssetWebLocationList(request *ExportAssetWebLocationListRequest) (response *ExportAssetWebLocationListResponse, err error) {
	if request == nil {
		request = NewExportAssetWebLocationListRequest()
	}
	response = NewExportAssetWebLocationListResponse()
	err = c.Send(request, response)
	return
}

func NewImportKnowledgePackageRequest() (request *ImportKnowledgePackageRequest) {
	request = &ImportKnowledgePackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ImportKnowledgePackage")
	return
}

func NewImportKnowledgePackageResponse() (response *ImportKnowledgePackageResponse) {
	response = &ImportKnowledgePackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 巡检任务-知识库管理-导入知识库包文件
func (c *Client) ImportKnowledgePackage(request *ImportKnowledgePackageRequest) (response *ImportKnowledgePackageResponse, err error) {
	if request == nil {
		request = NewImportKnowledgePackageRequest()
	}
	response = NewImportKnowledgePackageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUpdateGrayRequest() (request *ModifyUpdateGrayRequest) {
	request = &ModifyUpdateGrayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyUpdateGray")
	return
}

func NewModifyUpdateGrayResponse() (response *ModifyUpdateGrayResponse) {
	response = &ModifyUpdateGrayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 更新版本任务
func (c *Client) ModifyUpdateGray(request *ModifyUpdateGrayRequest) (response *ModifyUpdateGrayResponse, err error) {
	if request == nil {
		request = NewModifyUpdateGrayRequest()
	}
	response = NewModifyUpdateGrayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDefenseVulInfoRequest() (request *DescribeDefenseVulInfoRequest) {
	request = &DescribeDefenseVulInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeDefenseVulInfo")
	return
}

func NewDescribeDefenseVulInfoResponse() (response *DescribeDefenseVulInfoResponse) {
	response = &DescribeDefenseVulInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞管理-防御漏洞详情
func (c *Client) DescribeDefenseVulInfo(request *DescribeDefenseVulInfoRequest) (response *DescribeDefenseVulInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDefenseVulInfoRequest()
	}
	response = NewDescribeDefenseVulInfoResponse()
	err = c.Send(request, response)
	return
}

func NewImportRemoteUpdateConfigRequest() (request *ImportRemoteUpdateConfigRequest) {
	request = &ImportRemoteUpdateConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ImportRemoteUpdateConfig")
	return
}

func NewImportRemoteUpdateConfigResponse() (response *ImportRemoteUpdateConfigResponse) {
	response = &ImportRemoteUpdateConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入远程更新配置
func (c *Client) ImportRemoteUpdateConfig(request *ImportRemoteUpdateConfigRequest) (response *ImportRemoteUpdateConfigResponse, err error) {
	if request == nil {
		request = NewImportRemoteUpdateConfigRequest()
	}
	response = NewImportRemoteUpdateConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseFunctionsRequest() (request *DescribeLicenseFunctionsRequest) {
	request = &DescribeLicenseFunctionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeLicenseFunctions")
	return
}

func NewDescribeLicenseFunctionsResponse() (response *DescribeLicenseFunctionsResponse) {
	response = &DescribeLicenseFunctionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询授权功能
func (c *Client) DescribeLicenseFunctions(request *DescribeLicenseFunctionsRequest) (response *DescribeLicenseFunctionsResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseFunctionsRequest()
	}
	response = NewDescribeLicenseFunctionsResponse()
	err = c.Send(request, response)
	return
}

func NewExportBaselineHostDetectListRequest() (request *ExportBaselineHostDetectListRequest) {
	request = &ExportBaselineHostDetectListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportBaselineHostDetectList")
	return
}

func NewExportBaselineHostDetectListResponse() (response *ExportBaselineHostDetectListResponse) {
	response = &ExportBaselineHostDetectListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出基线主机检测
func (c *Client) ExportBaselineHostDetectList(request *ExportBaselineHostDetectListRequest) (response *ExportBaselineHostDetectListResponse, err error) {
	if request == nil {
		request = NewExportBaselineHostDetectListRequest()
	}
	response = NewExportBaselineHostDetectListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLoginEventsRequest() (request *DescribeLoginEventsRequest) {
	request = &DescribeLoginEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeLoginEvents")
	return
}

func NewDescribeLoginEventsResponse() (response *DescribeLoginEventsResponse) {
	response = &DescribeLoginEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 异常登录 - 列表
func (c *Client) DescribeLoginEvents(request *DescribeLoginEventsRequest) (response *DescribeLoginEventsResponse, err error) {
	if request == nil {
		request = NewDescribeLoginEventsRequest()
	}
	response = NewDescribeLoginEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetJarListRequest() (request *DescribeAssetJarListRequest) {
	request = &DescribeAssetJarListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetJarList")
	return
}

func NewDescribeAssetJarListResponse() (response *DescribeAssetJarListResponse) {
	response = &DescribeAssetJarListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Jar列表
func (c *Client) DescribeAssetJarList(request *DescribeAssetJarListRequest) (response *DescribeAssetJarListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetJarListRequest()
	}
	response = NewDescribeAssetJarListResponse()
	err = c.Send(request, response)
	return
}

func NewExportAgentsInfoListRequest() (request *ExportAgentsInfoListRequest) {
	request = &ExportAgentsInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAgentsInfoList")
	return
}

func NewExportAgentsInfoListResponse() (response *ExportAgentsInfoListResponse) {
	response = &ExportAgentsInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 资产管理 - 主机管理 - 导出列表
func (c *Client) ExportAgentsInfoList(request *ExportAgentsInfoListRequest) (response *ExportAgentsInfoListResponse, err error) {
	if request == nil {
		request = NewExportAgentsInfoListRequest()
	}
	response = NewExportAgentsInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRemoteUpdatePublishLogRequest() (request *ModifyRemoteUpdatePublishLogRequest) {
	request = &ModifyRemoteUpdatePublishLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyRemoteUpdatePublishLog")
	return
}

func NewModifyRemoteUpdatePublishLogResponse() (response *ModifyRemoteUpdatePublishLogResponse) {
	response = &ModifyRemoteUpdatePublishLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改远程更新发布历史记录
func (c *Client) ModifyRemoteUpdatePublishLog(request *ModifyRemoteUpdatePublishLogRequest) (response *ModifyRemoteUpdatePublishLogResponse, err error) {
	if request == nil {
		request = NewModifyRemoteUpdatePublishLogRequest()
	}
	response = NewModifyRemoteUpdatePublishLogResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVulVulsSwitchRequest() (request *ModifyVulVulsSwitchRequest) {
	request = &ModifyVulVulsSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyVulVulsSwitch")
	return
}

func NewModifyVulVulsSwitchResponse() (response *ModifyVulVulsSwitchResponse) {
	response = &ModifyVulVulsSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞管理-漏洞开关切换
func (c *Client) ModifyVulVulsSwitch(request *ModifyVulVulsSwitchRequest) (response *ModifyVulVulsSwitchResponse, err error) {
	if request == nil {
		request = NewModifyVulVulsSwitchRequest()
	}
	response = NewModifyVulVulsSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetJarListRequest() (request *ExportAssetJarListRequest) {
	request = &ExportAssetJarListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAssetJarList")
	return
}

func NewExportAssetJarListResponse() (response *ExportAssetJarListResponse) {
	response = &ExportAssetJarListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出Jar包列表
func (c *Client) ExportAssetJarList(request *ExportAssetJarListRequest) (response *ExportAssetJarListResponse, err error) {
	if request == nil {
		request = NewExportAssetJarListRequest()
	}
	response = NewExportAssetJarListResponse()
	err = c.Send(request, response)
	return
}

func NewExportVulDefenceMachineListRequest() (request *ExportVulDefenceMachineListRequest) {
	request = &ExportVulDefenceMachineListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportVulDefenceMachineList")
	return
}

func NewExportVulDefenceMachineListResponse() (response *ExportVulDefenceMachineListResponse) {
	response = &ExportVulDefenceMachineListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出防御主机列表
func (c *Client) ExportVulDefenceMachineList(request *ExportVulDefenceMachineListRequest) (response *ExportVulDefenceMachineListResponse, err error) {
	if request == nil {
		request = NewExportVulDefenceMachineListRequest()
	}
	response = NewExportVulDefenceMachineListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSocTamperProofInfoRequest() (request *DeleteSocTamperProofInfoRequest) {
	request = &DeleteSocTamperProofInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteSocTamperProofInfo")
	return
}

func NewDeleteSocTamperProofInfoResponse() (response *DeleteSocTamperProofInfoResponse) {
	response = &DeleteSocTamperProofInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 网页防篡改事件删除
func (c *Client) DeleteSocTamperProofInfo(request *DeleteSocTamperProofInfoRequest) (response *DeleteSocTamperProofInfoResponse, err error) {
	if request == nil {
		request = NewDeleteSocTamperProofInfoRequest()
	}
	response = NewDeleteSocTamperProofInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVulPackageRequest() (request *DeleteVulPackageRequest) {
	request = &DeleteVulPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteVulPackage")
	return
}

func NewDeleteVulPackageResponse() (response *DeleteVulPackageResponse) {
	response = &DeleteVulPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 任务包管理删除包文件
func (c *Client) DeleteVulPackage(request *DeleteVulPackageRequest) (response *DeleteVulPackageResponse, err error) {
	if request == nil {
		request = NewDeleteVulPackageRequest()
	}
	response = NewDeleteVulPackageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebLocationInfoRequest() (request *DescribeAssetWebLocationInfoRequest) {
	request = &DescribeAssetWebLocationInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetWebLocationInfo")
	return
}

func NewDescribeAssetWebLocationInfoResponse() (response *DescribeAssetWebLocationInfoResponse) {
	response = &DescribeAssetWebLocationInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Web站点详情
func (c *Client) DescribeAssetWebLocationInfo(request *DescribeAssetWebLocationInfoRequest) (response *DescribeAssetWebLocationInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebLocationInfoRequest()
	}
	response = NewDescribeAssetWebLocationInfoResponse()
	err = c.Send(request, response)
	return
}

func NewExportTasksRequest() (request *ExportTasksRequest) {
	request = &ExportTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportTasks")
	return
}

func NewExportTasksResponse() (response *ExportTasksResponse) {
	response = &ExportTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于异步导出数据量大的日志文件
func (c *Client) ExportTasks(request *ExportTasksRequest) (response *ExportTasksResponse, err error) {
	if request == nil {
		request = NewExportTasksRequest()
	}
	response = NewExportTasksResponse()
	err = c.Send(request, response)
	return
}

func NewCheckFileStatusRequest() (request *CheckFileStatusRequest) {
	request = &CheckFileStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "CheckFileStatus")
	return
}

func NewCheckFileStatusResponse() (response *CheckFileStatusResponse) {
	response = &CheckFileStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验文件状态
func (c *Client) CheckFileStatus(request *CheckFileStatusRequest) (response *CheckFileStatusResponse, err error) {
	if request == nil {
		request = NewCheckFileStatusRequest()
	}
	response = NewCheckFileStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVulScanTaskRequest() (request *CreateVulScanTaskRequest) {
	request = &CreateVulScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "CreateVulScanTask")
	return
}

func NewCreateVulScanTaskResponse() (response *CreateVulScanTaskResponse) {
	response = &CreateVulScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 任务管理 - 添加任务
func (c *Client) CreateVulScanTask(request *CreateVulScanTaskRequest) (response *CreateVulScanTaskResponse, err error) {
	if request == nil {
		request = NewCreateVulScanTaskRequest()
	}
	response = NewCreateVulScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewSeparateMalwareRequest() (request *SeparateMalwareRequest) {
	request = &SeparateMalwareRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "SeparateMalware")
	return
}

func NewSeparateMalwareResponse() (response *SeparateMalwareResponse) {
	response = &SeparateMalwareResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 隔离 木马文件
func (c *Client) SeparateMalware(request *SeparateMalwareRequest) (response *SeparateMalwareResponse, err error) {
	if request == nil {
		request = NewSeparateMalwareRequest()
	}
	response = NewSeparateMalwareResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityDynamicsRequest() (request *DescribeSecurityDynamicsRequest) {
	request = &DescribeSecurityDynamicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeSecurityDynamics")
	return
}

func NewDescribeSecurityDynamicsResponse() (response *DescribeSecurityDynamicsResponse) {
	response = &DescribeSecurityDynamicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取安全事件动态消息
func (c *Client) DescribeSecurityDynamics(request *DescribeSecurityDynamicsRequest) (response *DescribeSecurityDynamicsResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityDynamicsRequest()
	}
	response = NewDescribeSecurityDynamicsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetUserInfoRequest() (request *DescribeAssetUserInfoRequest) {
	request = &DescribeAssetUserInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetUserInfo")
	return
}

func NewDescribeAssetUserInfoResponse() (response *DescribeAssetUserInfoResponse) {
	response = &DescribeAssetUserInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账号详情
func (c *Client) DescribeAssetUserInfo(request *DescribeAssetUserInfoRequest) (response *DescribeAssetUserInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetUserInfoRequest()
	}
	response = NewDescribeAssetUserInfoResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetPlanTaskListRequest() (request *ExportAssetPlanTaskListRequest) {
	request = &ExportAssetPlanTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAssetPlanTaskList")
	return
}

func NewExportAssetPlanTaskListResponse() (response *ExportAssetPlanTaskListResponse) {
	response = &ExportAssetPlanTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理计划任务列表
func (c *Client) ExportAssetPlanTaskList(request *ExportAssetPlanTaskListRequest) (response *ExportAssetPlanTaskListResponse, err error) {
	if request == nil {
		request = NewExportAssetPlanTaskListRequest()
	}
	response = NewExportAssetPlanTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewExportAssetCoreModuleListRequest() (request *ExportAssetCoreModuleListRequest) {
	request = &ExportAssetCoreModuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportAssetCoreModuleList")
	return
}

func NewExportAssetCoreModuleListResponse() (response *ExportAssetCoreModuleListResponse) {
	response = &ExportAssetCoreModuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出资产管理内核模块列表
func (c *Client) ExportAssetCoreModuleList(request *ExportAssetCoreModuleListRequest) (response *ExportAssetCoreModuleListResponse, err error) {
	if request == nil {
		request = NewExportAssetCoreModuleListRequest()
	}
	response = NewExportAssetCoreModuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBaselineHostDetectListRequest() (request *DescribeBaselineHostDetectListRequest) {
	request = &DescribeBaselineHostDetectListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeBaselineHostDetectList")
	return
}

func NewDescribeBaselineHostDetectListResponse() (response *DescribeBaselineHostDetectListResponse) {
	response = &DescribeBaselineHostDetectListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基线检测主机列表
func (c *Client) DescribeBaselineHostDetectList(request *DescribeBaselineHostDetectListRequest) (response *DescribeBaselineHostDetectListResponse, err error) {
	if request == nil {
		request = NewDescribeBaselineHostDetectListRequest()
	}
	response = NewDescribeBaselineHostDetectListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBashEventRequest() (request *DeleteBashEventRequest) {
	request = &DeleteBashEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteBashEvent")
	return
}

func NewDeleteBashEventResponse() (response *DeleteBashEventResponse) {
	response = &DeleteBashEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除高危命令
func (c *Client) DeleteBashEvent(request *DeleteBashEventRequest) (response *DeleteBashEventResponse, err error) {
	if request == nil {
		request = NewDeleteBashEventRequest()
	}
	response = NewDeleteBashEventResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetAppCountRequest() (request *DescribeAssetAppCountRequest) {
	request = &DescribeAssetAppCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetAppCount")
	return
}

func NewDescribeAssetAppCountResponse() (response *DescribeAssetAppCountResponse) {
	response = &DescribeAssetAppCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有软件应用数量
func (c *Client) DescribeAssetAppCount(request *DescribeAssetAppCountRequest) (response *DescribeAssetAppCountResponse, err error) {
	if request == nil {
		request = NewDescribeAssetAppCountRequest()
	}
	response = NewDescribeAssetAppCountResponse()
	err = c.Send(request, response)
	return
}

func NewExportLoginEventsRequest() (request *ExportLoginEventsRequest) {
	request = &ExportLoginEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportLoginEvents")
	return
}

func NewExportLoginEventsResponse() (response *ExportLoginEventsResponse) {
	response = &ExportLoginEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 异常登录 - 列表 - 导出
func (c *Client) ExportLoginEvents(request *ExportLoginEventsRequest) (response *ExportLoginEventsResponse, err error) {
	if request == nil {
		request = NewExportLoginEventsRequest()
	}
	response = NewExportLoginEventsResponse()
	err = c.Send(request, response)
	return
}

func NewExportVulDefenceEventRequest() (request *ExportVulDefenceEventRequest) {
	request = &ExportVulDefenceEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportVulDefenceEvent")
	return
}

func NewExportVulDefenceEventResponse() (response *ExportVulDefenceEventResponse) {
	response = &ExportVulDefenceEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出漏洞防御事件
func (c *Client) ExportVulDefenceEvent(request *ExportVulDefenceEventRequest) (response *ExportVulDefenceEventResponse, err error) {
	if request == nil {
		request = NewExportVulDefenceEventRequest()
	}
	response = NewExportVulDefenceEventResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomerLicenseListRequest() (request *DescribeCustomerLicenseListRequest) {
	request = &DescribeCustomerLicenseListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeCustomerLicenseList")
	return
}

func NewDescribeCustomerLicenseListResponse() (response *DescribeCustomerLicenseListResponse) {
	response = &DescribeCustomerLicenseListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询License列表
func (c *Client) DescribeCustomerLicenseList(request *DescribeCustomerLicenseListRequest) (response *DescribeCustomerLicenseListResponse, err error) {
	if request == nil {
		request = NewDescribeCustomerLicenseListRequest()
	}
	response = NewDescribeCustomerLicenseListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemoteUpdateUserCurrentConfigsRequest() (request *DescribeRemoteUpdateUserCurrentConfigsRequest) {
	request = &DescribeRemoteUpdateUserCurrentConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeRemoteUpdateUserCurrentConfigs")
	return
}

func NewDescribeRemoteUpdateUserCurrentConfigsResponse() (response *DescribeRemoteUpdateUserCurrentConfigsResponse) {
	response = &DescribeRemoteUpdateUserCurrentConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询远程更新用户当前更新配置
func (c *Client) DescribeRemoteUpdateUserCurrentConfigs(request *DescribeRemoteUpdateUserCurrentConfigsRequest) (response *DescribeRemoteUpdateUserCurrentConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeRemoteUpdateUserCurrentConfigsRequest()
	}
	response = NewDescribeRemoteUpdateUserCurrentConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVulScanConfigInfoRequest() (request *CreateVulScanConfigInfoRequest) {
	request = &CreateVulScanConfigInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "CreateVulScanConfigInfo")
	return
}

func NewCreateVulScanConfigInfoResponse() (response *CreateVulScanConfigInfoResponse) {
	response = &CreateVulScanConfigInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置文件管理 - 新建配置文件
func (c *Client) CreateVulScanConfigInfo(request *CreateVulScanConfigInfoRequest) (response *CreateVulScanConfigInfoResponse, err error) {
	if request == nil {
		request = NewCreateVulScanConfigInfoRequest()
	}
	response = NewCreateVulScanConfigInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDNSKnowledgeListRequest() (request *DescribeDNSKnowledgeListRequest) {
	request = &DescribeDNSKnowledgeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeDNSKnowledgeList")
	return
}

func NewDescribeDNSKnowledgeListResponse() (response *DescribeDNSKnowledgeListResponse) {
	response = &DescribeDNSKnowledgeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全策略 - DNS知识库 - 列表
func (c *Client) DescribeDNSKnowledgeList(request *DescribeDNSKnowledgeListRequest) (response *DescribeDNSKnowledgeListResponse, err error) {
	if request == nil {
		request = NewDescribeDNSKnowledgeListRequest()
	}
	response = NewDescribeDNSKnowledgeListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulScanConfigFullListRequest() (request *DescribeVulScanConfigFullListRequest) {
	request = &DescribeVulScanConfigFullListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeVulScanConfigFullList")
	return
}

func NewDescribeVulScanConfigFullListResponse() (response *DescribeVulScanConfigFullListResponse) {
	response = &DescribeVulScanConfigFullListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞扫描设置列表
func (c *Client) DescribeVulScanConfigFullList(request *DescribeVulScanConfigFullListRequest) (response *DescribeVulScanConfigFullListResponse, err error) {
	if request == nil {
		request = NewDescribeVulScanConfigFullListRequest()
	}
	response = NewDescribeVulScanConfigFullListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUpdateTaskResultRequest() (request *DescribeUpdateTaskResultRequest) {
	request = &DescribeUpdateTaskResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeUpdateTaskResult")
	return
}

func NewDescribeUpdateTaskResultResponse() (response *DescribeUpdateTaskResultResponse) {
	response = &DescribeUpdateTaskResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询升级任务结果详情
func (c *Client) DescribeUpdateTaskResult(request *DescribeUpdateTaskResultRequest) (response *DescribeUpdateTaskResultResponse, err error) {
	if request == nil {
		request = NewDescribeUpdateTaskResultRequest()
	}
	response = NewDescribeUpdateTaskResultResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDefenceEventRequest() (request *DescribeVulDefenceEventRequest) {
	request = &DescribeVulDefenceEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeVulDefenceEvent")
	return
}

func NewDescribeVulDefenceEventResponse() (response *DescribeVulDefenceEventResponse) {
	response = &DescribeVulDefenceEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取漏洞防御事件列表
func (c *Client) DescribeVulDefenceEvent(request *DescribeVulDefenceEventRequest) (response *DescribeVulDefenceEventResponse, err error) {
	if request == nil {
		request = NewDescribeVulDefenceEventRequest()
	}
	response = NewDescribeVulDefenceEventResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteFileSpecimenRequest() (request *DeleteFileSpecimenRequest) {
	request = &DeleteFileSpecimenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteFileSpecimen")
	return
}

func NewDeleteFileSpecimenResponse() (response *DeleteFileSpecimenResponse) {
	response = &DeleteFileSpecimenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全策略 - 样本管理 - 删除
func (c *Client) DeleteFileSpecimen(request *DeleteFileSpecimenRequest) (response *DeleteFileSpecimenResponse, err error) {
	if request == nil {
		request = NewDeleteFileSpecimenRequest()
	}
	response = NewDeleteFileSpecimenResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRemoteUpdatePublishItemsRequest() (request *DeleteRemoteUpdatePublishItemsRequest) {
	request = &DeleteRemoteUpdatePublishItemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteRemoteUpdatePublishItems")
	return
}

func NewDeleteRemoteUpdatePublishItemsResponse() (response *DeleteRemoteUpdatePublishItemsResponse) {
	response = &DeleteRemoteUpdatePublishItemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除远程更新发布项
func (c *Client) DeleteRemoteUpdatePublishItems(request *DeleteRemoteUpdatePublishItemsRequest) (response *DeleteRemoteUpdatePublishItemsResponse, err error) {
	if request == nil {
		request = NewDeleteRemoteUpdatePublishItemsRequest()
	}
	response = NewDeleteRemoteUpdatePublishItemsResponse()
	err = c.Send(request, response)
	return
}

func NewPublishAgentRequest() (request *PublishAgentRequest) {
	request = &PublishAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "PublishAgent")
	return
}

func NewPublishAgentResponse() (response *PublishAgentResponse) {
	response = &PublishAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级管理 - 任务管理 - 发布配置
func (c *Client) PublishAgent(request *PublishAgentRequest) (response *PublishAgentResponse, err error) {
	if request == nil {
		request = NewPublishAgentRequest()
	}
	response = NewPublishAgentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBruteAttackTopRequest() (request *DescribeBruteAttackTopRequest) {
	request = &DescribeBruteAttackTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeBruteAttackTop")
	return
}

func NewDescribeBruteAttackTopResponse() (response *DescribeBruteAttackTopResponse) {
	response = &DescribeBruteAttackTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取暴力破解top
func (c *Client) DescribeBruteAttackTop(request *DescribeBruteAttackTopRequest) (response *DescribeBruteAttackTopResponse, err error) {
	if request == nil {
		request = NewDescribeBruteAttackTopRequest()
	}
	response = NewDescribeBruteAttackTopResponse()
	err = c.Send(request, response)
	return
}

func NewExportVulTopRequest() (request *ExportVulTopRequest) {
	request = &ExportVulTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportVulTop")
	return
}

func NewExportVulTopResponse() (response *ExportVulTopResponse) {
	response = &ExportVulTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出获取漏洞top
func (c *Client) ExportVulTop(request *ExportVulTopRequest) (response *ExportVulTopResponse, err error) {
	if request == nil {
		request = NewExportVulTopRequest()
	}
	response = NewExportVulTopResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMalwareRequest() (request *UpdateMalwareRequest) {
	request = &UpdateMalwareRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "UpdateMalware")
	return
}

func NewUpdateMalwareResponse() (response *UpdateMalwareResponse) {
	response = &UpdateMalwareResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 信任/取消信任/删除 木马文件
func (c *Client) UpdateMalware(request *UpdateMalwareRequest) (response *UpdateMalwareResponse, err error) {
	if request == nil {
		request = NewUpdateMalwareRequest()
	}
	response = NewUpdateMalwareResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVulScanTaskRequest() (request *DeleteVulScanTaskRequest) {
	request = &DeleteVulScanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DeleteVulScanTask")
	return
}

func NewDeleteVulScanTaskResponse() (response *DeleteVulScanTaskResponse) {
	response = &DeleteVulScanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 任务管理 - 删除任务
func (c *Client) DeleteVulScanTask(request *DeleteVulScanTaskRequest) (response *DeleteVulScanTaskResponse, err error) {
	if request == nil {
		request = NewDeleteVulScanTaskRequest()
	}
	response = NewDeleteVulScanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewExportVulVulsListRequest() (request *ExportVulVulsListRequest) {
	request = &ExportVulVulsListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportVulVulsList")
	return
}

func NewExportVulVulsListResponse() (response *ExportVulVulsListResponse) {
	response = &ExportVulVulsListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞管理-漏洞列表-导出
func (c *Client) ExportVulVulsList(request *ExportVulVulsListRequest) (response *ExportVulVulsListResponse, err error) {
	if request == nil {
		request = NewExportVulVulsListRequest()
	}
	response = NewExportVulVulsListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVulDefenceMainSettingRequest() (request *ModifyVulDefenceMainSettingRequest) {
	request = &ModifyVulDefenceMainSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyVulDefenceMainSetting")
	return
}

func NewModifyVulDefenceMainSettingResponse() (response *ModifyVulDefenceMainSettingResponse) {
	response = &ModifyVulDefenceMainSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改漏洞防御总开关
func (c *Client) ModifyVulDefenceMainSetting(request *ModifyVulDefenceMainSettingRequest) (response *ModifyVulDefenceMainSettingResponse, err error) {
	if request == nil {
		request = NewModifyVulDefenceMainSettingRequest()
	}
	response = NewModifyVulDefenceMainSettingResponse()
	err = c.Send(request, response)
	return
}

func NewExportBruteForceEventsRequest() (request *ExportBruteForceEventsRequest) {
	request = &ExportBruteForceEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportBruteForceEvents")
	return
}

func NewExportBruteForceEventsResponse() (response *ExportBruteForceEventsResponse) {
	response = &ExportBruteForceEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出密码破解列表
func (c *Client) ExportBruteForceEvents(request *ExportBruteForceEventsRequest) (response *ExportBruteForceEventsResponse, err error) {
	if request == nil {
		request = NewExportBruteForceEventsRequest()
	}
	response = NewExportBruteForceEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebServiceInfoListRequest() (request *DescribeAssetWebServiceInfoListRequest) {
	request = &DescribeAssetWebServiceInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetWebServiceInfoList")
	return
}

func NewDescribeAssetWebServiceInfoListResponse() (response *DescribeAssetWebServiceInfoListResponse) {
	response = &DescribeAssetWebServiceInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产管理Web服务列表
func (c *Client) DescribeAssetWebServiceInfoList(request *DescribeAssetWebServiceInfoListRequest) (response *DescribeAssetWebServiceInfoListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebServiceInfoListRequest()
	}
	response = NewDescribeAssetWebServiceInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewExportDefenseVulVulsListRequest() (request *ExportDefenseVulVulsListRequest) {
	request = &ExportDefenseVulVulsListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportDefenseVulVulsList")
	return
}

func NewExportDefenseVulVulsListResponse() (response *ExportDefenseVulVulsListResponse) {
	response = &ExportDefenseVulVulsListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 漏洞管理-导出漏洞列表
func (c *Client) ExportDefenseVulVulsList(request *ExportDefenseVulVulsListRequest) (response *ExportDefenseVulVulsListResponse, err error) {
	if request == nil {
		request = NewExportDefenseVulVulsListRequest()
	}
	response = NewExportDefenseVulVulsListResponse()
	err = c.Send(request, response)
	return
}

func NewExportSocTamperProofInfoRequest() (request *ExportSocTamperProofInfoRequest) {
	request = &ExportSocTamperProofInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportSocTamperProofInfo")
	return
}

func NewExportSocTamperProofInfoResponse() (response *ExportSocTamperProofInfoResponse) {
	response = &ExportSocTamperProofInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出网页防篡改列表
func (c *Client) ExportSocTamperProofInfo(request *ExportSocTamperProofInfoRequest) (response *ExportSocTamperProofInfoResponse, err error) {
	if request == nil {
		request = NewExportSocTamperProofInfoRequest()
	}
	response = NewExportSocTamperProofInfoResponse()
	err = c.Send(request, response)
	return
}

func NewExportHostRiskTrendRequest() (request *ExportHostRiskTrendRequest) {
	request = &ExportHostRiskTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ExportHostRiskTrend")
	return
}

func NewExportHostRiskTrendResponse() (response *ExportHostRiskTrendResponse) {
	response = &ExportHostRiskTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出风险主机事件趋势
func (c *Client) ExportHostRiskTrend(request *ExportHostRiskTrendRequest) (response *ExportHostRiskTrendResponse, err error) {
	if request == nil {
		request = NewExportHostRiskTrendRequest()
	}
	response = NewExportHostRiskTrendResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRemoteUpdatePublishItemsRequest() (request *CreateRemoteUpdatePublishItemsRequest) {
	request = &CreateRemoteUpdatePublishItemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "CreateRemoteUpdatePublishItems")
	return
}

func NewCreateRemoteUpdatePublishItemsResponse() (response *CreateRemoteUpdatePublishItemsResponse) {
	response = &CreateRemoteUpdatePublishItemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量创建远程更新发布项
func (c *Client) CreateRemoteUpdatePublishItems(request *CreateRemoteUpdatePublishItemsRequest) (response *CreateRemoteUpdatePublishItemsResponse, err error) {
	if request == nil {
		request = NewCreateRemoteUpdatePublishItemsRequest()
	}
	response = NewCreateRemoteUpdatePublishItemsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulAllEventsRequest() (request *DescribeVulAllEventsRequest) {
	request = &DescribeVulAllEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeVulAllEvents")
	return
}

func NewDescribeVulAllEventsResponse() (response *DescribeVulAllEventsResponse) {
	response = &DescribeVulAllEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全漏洞(全量) -列表
func (c *Client) DescribeVulAllEvents(request *DescribeVulAllEventsRequest) (response *DescribeVulAllEventsResponse, err error) {
	if request == nil {
		request = NewDescribeVulAllEventsRequest()
	}
	response = NewDescribeVulAllEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZKGroupListRequest() (request *DescribeZKGroupListRequest) {
	request = &DescribeZKGroupListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeZKGroupList")
	return
}

func NewDescribeZKGroupListResponse() (response *DescribeZKGroupListResponse) {
	response = &DescribeZKGroupListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 系统管理 - Zookeeper配置 - 列表
func (c *Client) DescribeZKGroupList(request *DescribeZKGroupListRequest) (response *DescribeZKGroupListResponse, err error) {
	if request == nil {
		request = NewDescribeZKGroupListRequest()
	}
	response = NewDescribeZKGroupListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetTypesRequest() (request *DescribeAssetTypesRequest) {
	request = &DescribeAssetTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeAssetTypes")
	return
}

func NewDescribeAssetTypesResponse() (response *DescribeAssetTypesResponse) {
	response = &DescribeAssetTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资产指纹类型列表
func (c *Client) DescribeAssetTypes(request *DescribeAssetTypesRequest) (response *DescribeAssetTypesResponse, err error) {
	if request == nil {
		request = NewDescribeAssetTypesRequest()
	}
	response = NewDescribeAssetTypesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemoteUpdateConfigListRequest() (request *DescribeRemoteUpdateConfigListRequest) {
	request = &DescribeRemoteUpdateConfigListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeRemoteUpdateConfigList")
	return
}

func NewDescribeRemoteUpdateConfigListResponse() (response *DescribeRemoteUpdateConfigListResponse) {
	response = &DescribeRemoteUpdateConfigListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取远程更新配置列表
func (c *Client) DescribeRemoteUpdateConfigList(request *DescribeRemoteUpdateConfigListRequest) (response *DescribeRemoteUpdateConfigListResponse, err error) {
	if request == nil {
		request = NewDescribeRemoteUpdateConfigListRequest()
	}
	response = NewDescribeRemoteUpdateConfigListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeStatisticTimeRequest() (request *DescribeStatisticTimeRequest) {
	request = &DescribeStatisticTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeStatisticTime")
	return
}

func NewDescribeStatisticTimeResponse() (response *DescribeStatisticTimeResponse) {
	response = &DescribeStatisticTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 数据包括各状态主机数量及其它列表数据总数
func (c *Client) DescribeStatisticTime(request *DescribeStatisticTimeRequest) (response *DescribeStatisticTimeResponse, err error) {
	if request == nil {
		request = NewDescribeStatisticTimeRequest()
	}
	response = NewDescribeStatisticTimeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLicenseInfoLicenseInfoRequest() (request *ModifyLicenseInfoLicenseInfoRequest) {
	request = &ModifyLicenseInfoLicenseInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "ModifyLicenseInfoLicenseInfo")
	return
}

func NewModifyLicenseInfoLicenseInfoResponse() (response *ModifyLicenseInfoLicenseInfoResponse) {
	response = &ModifyLicenseInfoLicenseInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑License  只能修改客户名和局点名
func (c *Client) ModifyLicenseInfoLicenseInfo(request *ModifyLicenseInfoLicenseInfoRequest) (response *ModifyLicenseInfoLicenseInfoResponse, err error) {
	if request == nil {
		request = NewModifyLicenseInfoLicenseInfoRequest()
	}
	response = NewModifyLicenseInfoLicenseInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstallPackageListRequest() (request *DescribeInstallPackageListRequest) {
	request = &DescribeInstallPackageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeInstallPackageList")
	return
}

func NewDescribeInstallPackageListResponse() (response *DescribeInstallPackageListResponse) {
	response = &DescribeInstallPackageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 版本管理 - 安装包管理 - 获取安装包列表
func (c *Client) DescribeInstallPackageList(request *DescribeInstallPackageListRequest) (response *DescribeInstallPackageListResponse, err error) {
	if request == nil {
		request = NewDescribeInstallPackageListRequest()
	}
	response = NewDescribeInstallPackageListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHostSafetyStatusRequest() (request *DescribeHostSafetyStatusRequest) {
	request = &DescribeHostSafetyStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ocwp", APIVersion, "DescribeHostSafetyStatus")
	return
}

func NewDescribeHostSafetyStatusResponse() (response *DescribeHostSafetyStatusResponse) {
	response = &DescribeHostSafetyStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全状态
func (c *Client) DescribeHostSafetyStatus(request *DescribeHostSafetyStatusRequest) (response *DescribeHostSafetyStatusResponse, err error) {
	if request == nil {
		request = NewDescribeHostSafetyStatusRequest()
	}
	response = NewDescribeHostSafetyStatusResponse()
	err = c.Send(request, response)
	return
}
