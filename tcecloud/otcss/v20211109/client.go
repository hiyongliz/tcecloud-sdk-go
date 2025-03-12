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

package v20211109

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2021-11-09"

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

func NewDescribeVulSummaryRequest() (request *DescribeVulSummaryRequest) {
	request = &DescribeVulSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVulSummary")
	return
}

func NewDescribeVulSummaryResponse() (response *DescribeVulSummaryResponse) {
	response = &DescribeVulSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞风险统计概览
func (c *Client) DescribeVulSummary(request *DescribeVulSummaryRequest) (response *DescribeVulSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeVulSummaryRequest()
	}
	response = NewDescribeVulSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeK8sApiAbnormalEventInfoRequest() (request *DescribeK8sApiAbnormalEventInfoRequest) {
	request = &DescribeK8sApiAbnormalEventInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeK8sApiAbnormalEventInfo")
	return
}

func NewDescribeK8sApiAbnormalEventInfoResponse() (response *DescribeK8sApiAbnormalEventInfoResponse) {
	response = &DescribeK8sApiAbnormalEventInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询k8s api 异常事件详情
func (c *Client) DescribeK8sApiAbnormalEventInfo(request *DescribeK8sApiAbnormalEventInfoRequest) (response *DescribeK8sApiAbnormalEventInfoResponse, err error) {
	if request == nil {
		request = NewDescribeK8sApiAbnormalEventInfoRequest()
	}
	response = NewDescribeK8sApiAbnormalEventInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellWhiteListsRequest() (request *DescribeReverseShellWhiteListsRequest) {
	request = &DescribeReverseShellWhiteListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeReverseShellWhiteLists")
	return
}

func NewDescribeReverseShellWhiteListsResponse() (response *DescribeReverseShellWhiteListsResponse) {
	response = &DescribeReverseShellWhiteListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时反弹shell白名单列表
func (c *Client) DescribeReverseShellWhiteLists(request *DescribeReverseShellWhiteListsRequest) (response *DescribeReverseShellWhiteListsResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellWhiteListsRequest()
	}
	response = NewDescribeReverseShellWhiteListsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessControlEventsExportRequest() (request *DescribeAccessControlEventsExportRequest) {
	request = &DescribeAccessControlEventsExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAccessControlEventsExport")
	return
}

func NewDescribeAccessControlEventsExportResponse() (response *DescribeAccessControlEventsExportResponse) {
	response = &DescribeAccessControlEventsExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时访问控制事件列表导出
func (c *Client) DescribeAccessControlEventsExport(request *DescribeAccessControlEventsExportRequest) (response *DescribeAccessControlEventsExportResponse, err error) {
	if request == nil {
		request = NewDescribeAccessControlEventsExportRequest()
	}
	response = NewDescribeAccessControlEventsExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageRegistryNamespaceListRequest() (request *DescribeImageRegistryNamespaceListRequest) {
	request = &DescribeImageRegistryNamespaceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeImageRegistryNamespaceList")
	return
}

func NewDescribeImageRegistryNamespaceListResponse() (response *DescribeImageRegistryNamespaceListResponse) {
	response = &DescribeImageRegistryNamespaceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户镜像仓库下的项目名称列表
func (c *Client) DescribeImageRegistryNamespaceList(request *DescribeImageRegistryNamespaceListRequest) (response *DescribeImageRegistryNamespaceListResponse, err error) {
	if request == nil {
		request = NewDescribeImageRegistryNamespaceListRequest()
	}
	response = NewDescribeImageRegistryNamespaceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessControlEventsRequest() (request *DescribeAccessControlEventsRequest) {
	request = &DescribeAccessControlEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAccessControlEvents")
	return
}

func NewDescribeAccessControlEventsResponse() (response *DescribeAccessControlEventsResponse) {
	response = &DescribeAccessControlEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时访问控制事件列表
func (c *Client) DescribeAccessControlEvents(request *DescribeAccessControlEventsRequest) (response *DescribeAccessControlEventsResponse, err error) {
	if request == nil {
		request = NewDescribeAccessControlEventsRequest()
	}
	response = NewDescribeAccessControlEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterDetailRequest() (request *DescribeClusterDetailRequest) {
	request = &DescribeClusterDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeClusterDetail")
	return
}

func NewDescribeClusterDetailResponse() (response *DescribeClusterDetailResponse) {
	response = &DescribeClusterDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询单个集群的详细信息
func (c *Client) DescribeClusterDetail(request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
	if request == nil {
		request = NewDescribeClusterDetailRequest()
	}
	response = NewDescribeClusterDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComplianceAssetPolicyItemListRequest() (request *DescribeComplianceAssetPolicyItemListRequest) {
	request = &DescribeComplianceAssetPolicyItemListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeComplianceAssetPolicyItemList")
	return
}

func NewDescribeComplianceAssetPolicyItemListResponse() (response *DescribeComplianceAssetPolicyItemListResponse) {
	response = &DescribeComplianceAssetPolicyItemListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某资产下的检测项列表
func (c *Client) DescribeComplianceAssetPolicyItemList(request *DescribeComplianceAssetPolicyItemListRequest) (response *DescribeComplianceAssetPolicyItemListResponse, err error) {
	if request == nil {
		request = NewDescribeComplianceAssetPolicyItemListRequest()
	}
	response = NewDescribeComplianceAssetPolicyItemListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskSyscallWhiteListDetailRequest() (request *DescribeRiskSyscallWhiteListDetailRequest) {
	request = &DescribeRiskSyscallWhiteListDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeRiskSyscallWhiteListDetail")
	return
}

func NewDescribeRiskSyscallWhiteListDetailResponse() (response *DescribeRiskSyscallWhiteListDetailResponse) {
	response = &DescribeRiskSyscallWhiteListDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时高危系统调用白名单详细信息（暂未开发）
func (c *Client) DescribeRiskSyscallWhiteListDetail(request *DescribeRiskSyscallWhiteListDetailRequest) (response *DescribeRiskSyscallWhiteListDetailResponse, err error) {
	if request == nil {
		request = NewDescribeRiskSyscallWhiteListDetailRequest()
	}
	response = NewDescribeRiskSyscallWhiteListDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterDetailExportJobRequest() (request *CreateClusterDetailExportJobRequest) {
	request = &CreateClusterDetailExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateClusterDetailExportJob")
	return
}

func NewCreateClusterDetailExportJobResponse() (response *CreateClusterDetailExportJobResponse) {
	response = &CreateClusterDetailExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集群详情导出任务,只支持单个集群的详情导出
func (c *Client) CreateClusterDetailExportJob(request *CreateClusterDetailExportJobRequest) (response *CreateClusterDetailExportJobResponse, err error) {
	if request == nil {
		request = NewCreateClusterDetailExportJobRequest()
	}
	response = NewCreateClusterDetailExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessControlDetailRequest() (request *DescribeAccessControlDetailRequest) {
	request = &DescribeAccessControlDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAccessControlDetail")
	return
}

func NewDescribeAccessControlDetailResponse() (response *DescribeAccessControlDetailResponse) {
	response = &DescribeAccessControlDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时访问控制事件详细信息
func (c *Client) DescribeAccessControlDetail(request *DescribeAccessControlDetailRequest) (response *DescribeAccessControlDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAccessControlDetailRequest()
	}
	response = NewDescribeAccessControlDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSearchTemplatesRequest() (request *DescribeSearchTemplatesRequest) {
	request = &DescribeSearchTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeSearchTemplates")
	return
}

func NewDescribeSearchTemplatesResponse() (response *DescribeSearchTemplatesResponse) {
	response = &DescribeSearchTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取快速检索列表
func (c *Client) DescribeSearchTemplates(request *DescribeSearchTemplatesRequest) (response *DescribeSearchTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeSearchTemplatesRequest()
	}
	response = NewDescribeSearchTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetComponentListRequest() (request *DescribeAssetComponentListRequest) {
	request = &DescribeAssetComponentListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetComponentList")
	return
}

func NewDescribeAssetComponentListResponse() (response *DescribeAssetComponentListResponse) {
	response = &DescribeAssetComponentListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器组件列表
func (c *Client) DescribeAssetComponentList(request *DescribeAssetComponentListRequest) (response *DescribeAssetComponentListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetComponentListRequest()
	}
	response = NewDescribeAssetComponentListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEscapeWhiteListExportJobRequest() (request *CreateEscapeWhiteListExportJobRequest) {
	request = &CreateEscapeWhiteListExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateEscapeWhiteListExportJob")
	return
}

func NewCreateEscapeWhiteListExportJobResponse() (response *CreateEscapeWhiteListExportJobResponse) {
	response = &CreateEscapeWhiteListExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建逃逸白名单导出任务
func (c *Client) CreateEscapeWhiteListExportJob(request *CreateEscapeWhiteListExportJobRequest) (response *CreateEscapeWhiteListExportJobResponse, err error) {
	if request == nil {
		request = NewCreateEscapeWhiteListExportJobRequest()
	}
	response = NewCreateEscapeWhiteListExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskResultSummaryRequest() (request *DescribeTaskResultSummaryRequest) {
	request = &DescribeTaskResultSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeTaskResultSummary")
	return
}

func NewDescribeTaskResultSummaryResponse() (response *DescribeTaskResultSummaryResponse) {
	response = &DescribeTaskResultSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询检查结果总览，返回受影响的节点数量，返回7天的数据，总共7个
func (c *Client) DescribeTaskResultSummary(request *DescribeTaskResultSummaryRequest) (response *DescribeTaskResultSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeTaskResultSummaryRequest()
	}
	response = NewDescribeTaskResultSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeESHitsRequest() (request *DescribeESHitsRequest) {
	request = &DescribeESHitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeESHits")
	return
}

func NewDescribeESHitsResponse() (response *DescribeESHitsResponse) {
	response = &DescribeESHitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ES查询文档列表
func (c *Client) DescribeESHits(request *DescribeESHitsRequest) (response *DescribeESHitsResponse, err error) {
	if request == nil {
		request = NewDescribeESHitsRequest()
	}
	response = NewDescribeESHitsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusAutoIsolateSampleDownloadURLRequest() (request *DescribeVirusAutoIsolateSampleDownloadURLRequest) {
	request = &DescribeVirusAutoIsolateSampleDownloadURLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVirusAutoIsolateSampleDownloadURL")
	return
}

func NewDescribeVirusAutoIsolateSampleDownloadURLResponse() (response *DescribeVirusAutoIsolateSampleDownloadURLResponse) {
	response = &DescribeVirusAutoIsolateSampleDownloadURLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询木马自动隔离样本下载链接
func (c *Client) DescribeVirusAutoIsolateSampleDownloadURL(request *DescribeVirusAutoIsolateSampleDownloadURLRequest) (response *DescribeVirusAutoIsolateSampleDownloadURLResponse, err error) {
	if request == nil {
		request = NewDescribeVirusAutoIsolateSampleDownloadURLRequest()
	}
	response = NewDescribeVirusAutoIsolateSampleDownloadURLResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNamespaceListRequest() (request *DescribeNamespaceListRequest) {
	request = &DescribeNamespaceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeNamespaceList")
	return
}

func NewDescribeNamespaceListResponse() (response *DescribeNamespaceListResponse) {
	response = &DescribeNamespaceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取命名空间列表
func (c *Client) DescribeNamespaceList(request *DescribeNamespaceListRequest) (response *DescribeNamespaceListResponse, err error) {
	if request == nil {
		request = NewDescribeNamespaceListRequest()
	}
	response = NewDescribeNamespaceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageRiskSummaryRequest() (request *DescribeImageRiskSummaryRequest) {
	request = &DescribeImageRiskSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeImageRiskSummary")
	return
}

func NewDescribeImageRiskSummaryResponse() (response *DescribeImageRiskSummaryResponse) {
	response = &DescribeImageRiskSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询本地镜像风险概览
func (c *Client) DescribeImageRiskSummary(request *DescribeImageRiskSummaryRequest) (response *DescribeImageRiskSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeImageRiskSummaryRequest()
	}
	response = NewDescribeImageRiskSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVulRegistryImageExportJobRequest() (request *CreateVulRegistryImageExportJobRequest) {
	request = &CreateVulRegistryImageExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateVulRegistryImageExportJob")
	return
}

func NewCreateVulRegistryImageExportJobResponse() (response *CreateVulRegistryImageExportJobResponse) {
	response = &CreateVulRegistryImageExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建受漏洞影响的仓库镜像导出任务
func (c *Client) CreateVulRegistryImageExportJob(request *CreateVulRegistryImageExportJobRequest) (response *CreateVulRegistryImageExportJobResponse, err error) {
	if request == nil {
		request = NewCreateVulRegistryImageExportJobRequest()
	}
	response = NewCreateVulRegistryImageExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageComponentListRequest() (request *DescribeImageComponentListRequest) {
	request = &DescribeImageComponentListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeImageComponentList")
	return
}

func NewDescribeImageComponentListResponse() (response *DescribeImageComponentListResponse) {
	response = &DescribeImageComponentListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询本地镜像组件列表
func (c *Client) DescribeImageComponentList(request *DescribeImageComponentListRequest) (response *DescribeImageComponentListResponse, err error) {
	if request == nil {
		request = NewDescribeImageComponentListRequest()
	}
	response = NewDescribeImageComponentListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusScanTaskStatusRequest() (request *DescribeVirusScanTaskStatusRequest) {
	request = &DescribeVirusScanTaskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVirusScanTaskStatus")
	return
}

func NewDescribeVirusScanTaskStatusResponse() (response *DescribeVirusScanTaskStatusResponse) {
	response = &DescribeVirusScanTaskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时查询文件查杀任务状态
func (c *Client) DescribeVirusScanTaskStatus(request *DescribeVirusScanTaskStatusRequest) (response *DescribeVirusScanTaskStatusResponse, err error) {
	if request == nil {
		request = NewDescribeVirusScanTaskStatusRequest()
	}
	response = NewDescribeVirusScanTaskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskSyscallNamesRequest() (request *DescribeRiskSyscallNamesRequest) {
	request = &DescribeRiskSyscallNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeRiskSyscallNames")
	return
}

func NewDescribeRiskSyscallNamesResponse() (response *DescribeRiskSyscallNamesResponse) {
	response = &DescribeRiskSyscallNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时高危系统调用系统名称列表
func (c *Client) DescribeRiskSyscallNames(request *DescribeRiskSyscallNamesRequest) (response *DescribeRiskSyscallNamesResponse, err error) {
	if request == nil {
		request = NewDescribeRiskSyscallNamesRequest()
	}
	response = NewDescribeRiskSyscallNamesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEmergencyVulListRequest() (request *DescribeEmergencyVulListRequest) {
	request = &DescribeEmergencyVulListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeEmergencyVulList")
	return
}

func NewDescribeEmergencyVulListResponse() (response *DescribeEmergencyVulListResponse) {
	response = &DescribeEmergencyVulListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应急漏洞列表
func (c *Client) DescribeEmergencyVulList(request *DescribeEmergencyVulListRequest) (response *DescribeEmergencyVulListResponse, err error) {
	if request == nil {
		request = NewDescribeEmergencyVulListRequest()
	}
	response = NewDescribeEmergencyVulListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSystemVulListRequest() (request *DescribeSystemVulListRequest) {
	request = &DescribeSystemVulListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeSystemVulList")
	return
}

func NewDescribeSystemVulListResponse() (response *DescribeSystemVulListResponse) {
	response = &DescribeSystemVulListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询系统漏洞列表
func (c *Client) DescribeSystemVulList(request *DescribeSystemVulListRequest) (response *DescribeSystemVulListResponse, err error) {
	if request == nil {
		request = NewDescribeSystemVulListRequest()
	}
	response = NewDescribeSystemVulListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateExportComplianceStatusListJobRequest() (request *CreateExportComplianceStatusListJobRequest) {
	request = &CreateExportComplianceStatusListJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateExportComplianceStatusListJob")
	return
}

func NewCreateExportComplianceStatusListJobResponse() (response *CreateExportComplianceStatusListJobResponse) {
	response = &CreateExportComplianceStatusListJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建一个导出安全合规信息的任务
func (c *Client) CreateExportComplianceStatusListJob(request *CreateExportComplianceStatusListJobRequest) (response *CreateExportComplianceStatusListJobResponse, err error) {
	if request == nil {
		request = NewCreateExportComplianceStatusListJobRequest()
	}
	response = NewCreateExportComplianceStatusListJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageVirusListRequest() (request *DescribeAssetImageVirusListRequest) {
	request = &DescribeAssetImageVirusListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageVirusList")
	return
}

func NewDescribeAssetImageVirusListResponse() (response *DescribeAssetImageVirusListResponse) {
	response = &DescribeAssetImageVirusListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像病毒列表
func (c *Client) DescribeAssetImageVirusList(request *DescribeAssetImageVirusListRequest) (response *DescribeAssetImageVirusListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageVirusListRequest()
	}
	response = NewDescribeAssetImageVirusListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSearchLogsRequest() (request *DescribeSearchLogsRequest) {
	request = &DescribeSearchLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeSearchLogs")
	return
}

func NewDescribeSearchLogsResponse() (response *DescribeSearchLogsResponse) {
	response = &DescribeSearchLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取历史搜索记录
func (c *Client) DescribeSearchLogs(request *DescribeSearchLogsRequest) (response *DescribeSearchLogsResponse, err error) {
	if request == nil {
		request = NewDescribeSearchLogsRequest()
	}
	response = NewDescribeSearchLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSupportRegionsRequest() (request *DescribeSupportRegionsRequest) {
	request = &DescribeSupportRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeSupportRegions")
	return
}

func NewDescribeSupportRegionsResponse() (response *DescribeSupportRegionsResponse) {
	response = &DescribeSupportRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取支持的地域
func (c *Client) DescribeSupportRegions(request *DescribeSupportRegionsRequest) (response *DescribeSupportRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeSupportRegionsRequest()
	}
	response = NewDescribeSupportRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusSummaryRequest() (request *DescribeVirusSummaryRequest) {
	request = &DescribeVirusSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVirusSummary")
	return
}

func NewDescribeVirusSummaryResponse() (response *DescribeVirusSummaryResponse) {
	response = &DescribeVirusSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时查询木马概览信息
func (c *Client) DescribeVirusSummary(request *DescribeVirusSummaryRequest) (response *DescribeVirusSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeVirusSummaryRequest()
	}
	response = NewDescribeVirusSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewCreateImageDenyEventExportJobRequest() (request *CreateImageDenyEventExportJobRequest) {
	request = &CreateImageDenyEventExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateImageDenyEventExportJob")
	return
}

func NewCreateImageDenyEventExportJobResponse() (response *CreateImageDenyEventExportJobResponse) {
	response = &CreateImageDenyEventExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建镜像拦截事件导出任务
func (c *Client) CreateImageDenyEventExportJob(request *CreateImageDenyEventExportJobRequest) (response *CreateImageDenyEventExportJobResponse, err error) {
	if request == nil {
		request = NewCreateImageDenyEventExportJobRequest()
	}
	response = NewCreateImageDenyEventExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewModifySystemPolicyRequest() (request *ModifySystemPolicyRequest) {
	request = &ModifySystemPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "ModifySystemPolicy")
	return
}

func NewModifySystemPolicyResponse() (response *ModifySystemPolicyResponse) {
	response = &ModifySystemPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑运行时系统策略列表
func (c *Client) ModifySystemPolicy(request *ModifySystemPolicyRequest) (response *ModifySystemPolicyResponse, err error) {
	if request == nil {
		request = NewModifySystemPolicyRequest()
	}
	response = NewModifySystemPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComplianceAssetDetailInfoRequest() (request *DescribeComplianceAssetDetailInfoRequest) {
	request = &DescribeComplianceAssetDetailInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeComplianceAssetDetailInfo")
	return
}

func NewDescribeComplianceAssetDetailInfoResponse() (response *DescribeComplianceAssetDetailInfoResponse) {
	response = &DescribeComplianceAssetDetailInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某个资产的详情
func (c *Client) DescribeComplianceAssetDetailInfo(request *DescribeComplianceAssetDetailInfoRequest) (response *DescribeComplianceAssetDetailInfoResponse, err error) {
	if request == nil {
		request = NewDescribeComplianceAssetDetailInfoRequest()
	}
	response = NewDescribeComplianceAssetDetailInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbnormalProcessEventTendencyRequest() (request *DescribeAbnormalProcessEventTendencyRequest) {
	request = &DescribeAbnormalProcessEventTendencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAbnormalProcessEventTendency")
	return
}

func NewDescribeAbnormalProcessEventTendencyResponse() (response *DescribeAbnormalProcessEventTendencyResponse) {
	response = &DescribeAbnormalProcessEventTendencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询待处理异常进程事件趋势
func (c *Client) DescribeAbnormalProcessEventTendency(request *DescribeAbnormalProcessEventTendencyRequest) (response *DescribeAbnormalProcessEventTendencyResponse, err error) {
	if request == nil {
		request = NewDescribeAbnormalProcessEventTendencyRequest()
	}
	response = NewDescribeAbnormalProcessEventTendencyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCheckItemListRequest() (request *DescribeCheckItemListRequest) {
	request = &DescribeCheckItemListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeCheckItemList")
	return
}

func NewDescribeCheckItemListResponse() (response *DescribeCheckItemListResponse) {
	response = &DescribeCheckItemListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所有检查项接口，返回总数和检查项列表
func (c *Client) DescribeCheckItemList(request *DescribeCheckItemListRequest) (response *DescribeCheckItemListResponse, err error) {
	if request == nil {
		request = NewDescribeCheckItemListRequest()
	}
	response = NewDescribeCheckItemListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryVulListRequest() (request *DescribeAssetImageRegistryVulListRequest) {
	request = &DescribeAssetImageRegistryVulListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageRegistryVulList")
	return
}

func NewDescribeAssetImageRegistryVulListResponse() (response *DescribeAssetImageRegistryVulListResponse) {
	response = &DescribeAssetImageRegistryVulListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库查询镜像漏洞列表
func (c *Client) DescribeAssetImageRegistryVulList(request *DescribeAssetImageRegistryVulListRequest) (response *DescribeAssetImageRegistryVulListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryVulListRequest()
	}
	response = NewDescribeAssetImageRegistryVulListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbnormalProcessRulesExportRequest() (request *DescribeAbnormalProcessRulesExportRequest) {
	request = &DescribeAbnormalProcessRulesExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAbnormalProcessRulesExport")
	return
}

func NewDescribeAbnormalProcessRulesExportResponse() (response *DescribeAbnormalProcessRulesExportResponse) {
	response = &DescribeAbnormalProcessRulesExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时异常进程策略列表导出
func (c *Client) DescribeAbnormalProcessRulesExport(request *DescribeAbnormalProcessRulesExportRequest) (response *DescribeAbnormalProcessRulesExportResponse, err error) {
	if request == nil {
		request = NewDescribeAbnormalProcessRulesExportRequest()
	}
	response = NewDescribeAbnormalProcessRulesExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRiskListExportRequest() (request *DescribeAssetImageRiskListExportRequest) {
	request = &DescribeAssetImageRiskListExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageRiskListExport")
	return
}

func NewDescribeAssetImageRiskListExportResponse() (response *DescribeAssetImageRiskListExportResponse) {
	response = &DescribeAssetImageRiskListExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像风险列表导出
func (c *Client) DescribeAssetImageRiskListExport(request *DescribeAssetImageRiskListExportRequest) (response *DescribeAssetImageRiskListExportResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRiskListExportRequest()
	}
	response = NewDescribeAssetImageRiskListExportResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVulImageExportJobRequest() (request *CreateVulImageExportJobRequest) {
	request = &CreateVulImageExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateVulImageExportJob")
	return
}

func NewCreateVulImageExportJobResponse() (response *CreateVulImageExportJobResponse) {
	response = &CreateVulImageExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建受漏洞影响的镜像导出任务
func (c *Client) CreateVulImageExportJob(request *CreateVulImageExportJobRequest) (response *CreateVulImageExportJobResponse, err error) {
	if request == nil {
		request = NewCreateVulImageExportJobRequest()
	}
	response = NewCreateVulImageExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserPurchaseInfoExportRequest() (request *DescribeUserPurchaseInfoExportRequest) {
	request = &DescribeUserPurchaseInfoExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeUserPurchaseInfoExport")
	return
}

func NewDescribeUserPurchaseInfoExportResponse() (response *DescribeUserPurchaseInfoExportResponse) {
	response = &DescribeUserPurchaseInfoExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户购买容器安全服务列表导出
func (c *Client) DescribeUserPurchaseInfoExport(request *DescribeUserPurchaseInfoExportRequest) (response *DescribeUserPurchaseInfoExportResponse, err error) {
	if request == nil {
		request = NewDescribeUserPurchaseInfoExportRequest()
	}
	response = NewDescribeUserPurchaseInfoExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogStorageStatisticRequest() (request *DescribeLogStorageStatisticRequest) {
	request = &DescribeLogStorageStatisticRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeLogStorageStatistic")
	return
}

func NewDescribeLogStorageStatisticResponse() (response *DescribeLogStorageStatisticResponse) {
	response = &DescribeLogStorageStatisticResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志检索容量使用统计
func (c *Client) DescribeLogStorageStatistic(request *DescribeLogStorageStatisticRequest) (response *DescribeLogStorageStatisticResponse, err error) {
	if request == nil {
		request = NewDescribeLogStorageStatisticRequest()
	}
	response = NewDescribeLogStorageStatisticResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeK8sApiAbnormalRuleScopeListRequest() (request *DescribeK8sApiAbnormalRuleScopeListRequest) {
	request = &DescribeK8sApiAbnormalRuleScopeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeK8sApiAbnormalRuleScopeList")
	return
}

func NewDescribeK8sApiAbnormalRuleScopeListResponse() (response *DescribeK8sApiAbnormalRuleScopeListResponse) {
	response = &DescribeK8sApiAbnormalRuleScopeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询k8s api 异常规则中范围列表
func (c *Client) DescribeK8sApiAbnormalRuleScopeList(request *DescribeK8sApiAbnormalRuleScopeListRequest) (response *DescribeK8sApiAbnormalRuleScopeListResponse, err error) {
	if request == nil {
		request = NewDescribeK8sApiAbnormalRuleScopeListRequest()
	}
	response = NewDescribeK8sApiAbnormalRuleScopeListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVulContainerExportJobRequest() (request *CreateVulContainerExportJobRequest) {
	request = &CreateVulContainerExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateVulContainerExportJob")
	return
}

func NewCreateVulContainerExportJobResponse() (response *CreateVulContainerExportJobResponse) {
	response = &CreateVulContainerExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建受漏洞影响的容器导出任务
func (c *Client) CreateVulContainerExportJob(request *CreateVulContainerExportJobRequest) (response *CreateVulContainerExportJobResponse, err error) {
	if request == nil {
		request = NewCreateVulContainerExportJobRequest()
	}
	response = NewCreateVulContainerExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusEventTendencyRequest() (request *DescribeVirusEventTendencyRequest) {
	request = &DescribeVirusEventTendencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVirusEventTendency")
	return
}

func NewDescribeVirusEventTendencyResponse() (response *DescribeVirusEventTendencyResponse) {
	response = &DescribeVirusEventTendencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询木马事件趋势
func (c *Client) DescribeVirusEventTendency(request *DescribeVirusEventTendencyRequest) (response *DescribeVirusEventTendencyResponse, err error) {
	if request == nil {
		request = NewDescribeVirusEventTendencyRequest()
	}
	response = NewDescribeVirusEventTendencyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusAutoIsolateSettingListRequest() (request *DescribeVirusAutoIsolateSettingListRequest) {
	request = &DescribeVirusAutoIsolateSettingListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVirusAutoIsolateSettingList")
	return
}

func NewDescribeVirusAutoIsolateSettingListResponse() (response *DescribeVirusAutoIsolateSettingListResponse) {
	response = &DescribeVirusAutoIsolateSettingListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询木马自动隔离设置列表
func (c *Client) DescribeVirusAutoIsolateSettingList(request *DescribeVirusAutoIsolateSettingListRequest) (response *DescribeVirusAutoIsolateSettingListResponse, err error) {
	if request == nil {
		request = NewDescribeVirusAutoIsolateSettingListRequest()
	}
	response = NewDescribeVirusAutoIsolateSettingListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEventEscapeImageDetailRequest() (request *DescribeEventEscapeImageDetailRequest) {
	request = &DescribeEventEscapeImageDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeEventEscapeImageDetail")
	return
}

func NewDescribeEventEscapeImageDetailResponse() (response *DescribeEventEscapeImageDetailResponse) {
	response = &DescribeEventEscapeImageDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询风险容器镜像详情
func (c *Client) DescribeEventEscapeImageDetail(request *DescribeEventEscapeImageDetailRequest) (response *DescribeEventEscapeImageDetailResponse, err error) {
	if request == nil {
		request = NewDescribeEventEscapeImageDetailRequest()
	}
	response = NewDescribeEventEscapeImageDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskSyscallWhiteListsExportRequest() (request *DescribeRiskSyscallWhiteListsExportRequest) {
	request = &DescribeRiskSyscallWhiteListsExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeRiskSyscallWhiteListsExport")
	return
}

func NewDescribeRiskSyscallWhiteListsExportResponse() (response *DescribeRiskSyscallWhiteListsExportResponse) {
	response = &DescribeRiskSyscallWhiteListsExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时高危系统调用白名单列表导出
func (c *Client) DescribeRiskSyscallWhiteListsExport(request *DescribeRiskSyscallWhiteListsExportRequest) (response *DescribeRiskSyscallWhiteListsExportResponse, err error) {
	if request == nil {
		request = NewDescribeRiskSyscallWhiteListsExportRequest()
	}
	response = NewDescribeRiskSyscallWhiteListsExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEscapeWhiteListRequest() (request *DescribeEscapeWhiteListRequest) {
	request = &DescribeEscapeWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeEscapeWhiteList")
	return
}

func NewDescribeEscapeWhiteListResponse() (response *DescribeEscapeWhiteListResponse) {
	response = &DescribeEscapeWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询逃逸白名单
func (c *Client) DescribeEscapeWhiteList(request *DescribeEscapeWhiteListRequest) (response *DescribeEscapeWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeEscapeWhiteListRequest()
	}
	response = NewDescribeEscapeWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageListExportRequest() (request *DescribeAssetImageListExportRequest) {
	request = &DescribeAssetImageListExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageListExport")
	return
}

func NewDescribeAssetImageListExportResponse() (response *DescribeAssetImageListExportResponse) {
	response = &DescribeAssetImageListExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像列表导出
func (c *Client) DescribeAssetImageListExport(request *DescribeAssetImageListExportRequest) (response *DescribeAssetImageListExportResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageListExportRequest()
	}
	response = NewDescribeAssetImageListExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterNodesRequest() (request *DescribeClusterNodesRequest) {
	request = &DescribeClusterNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeClusterNodes")
	return
}

func NewDescribeClusterNodesResponse() (response *DescribeClusterNodesResponse) {
	response = &DescribeClusterNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群节点信息
func (c *Client) DescribeClusterNodes(request *DescribeClusterNodesRequest) (response *DescribeClusterNodesResponse, err error) {
	if request == nil {
		request = NewDescribeClusterNodesRequest()
	}
	response = NewDescribeClusterNodesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageDenyEventListRequest() (request *DescribeImageDenyEventListRequest) {
	request = &DescribeImageDenyEventListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeImageDenyEventList")
	return
}

func NewDescribeImageDenyEventListResponse() (response *DescribeImageDenyEventListResponse) {
	response = &DescribeImageDenyEventListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像拦截事件列表
func (c *Client) DescribeImageDenyEventList(request *DescribeImageDenyEventListRequest) (response *DescribeImageDenyEventListResponse, err error) {
	if request == nil {
		request = NewDescribeImageDenyEventListRequest()
	}
	response = NewDescribeImageDenyEventListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePodContainersRequest() (request *DescribePodContainersRequest) {
	request = &DescribePodContainersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribePodContainers")
	return
}

func NewDescribePodContainersResponse() (response *DescribePodContainersResponse) {
	response = &DescribePodContainersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询pod对应的容器信息
func (c *Client) DescribePodContainers(request *DescribePodContainersRequest) (response *DescribePodContainersResponse, err error) {
	if request == nil {
		request = NewDescribePodContainersRequest()
	}
	response = NewDescribePodContainersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbnormalProcessEventsRequest() (request *DescribeAbnormalProcessEventsRequest) {
	request = &DescribeAbnormalProcessEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAbnormalProcessEvents")
	return
}

func NewDescribeAbnormalProcessEventsResponse() (response *DescribeAbnormalProcessEventsResponse) {
	response = &DescribeAbnormalProcessEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时异常进程列表
func (c *Client) DescribeAbnormalProcessEvents(request *DescribeAbnormalProcessEventsRequest) (response *DescribeAbnormalProcessEventsResponse, err error) {
	if request == nil {
		request = NewDescribeAbnormalProcessEventsRequest()
	}
	response = NewDescribeAbnormalProcessEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterIngressYamlRequest() (request *DescribeClusterIngressYamlRequest) {
	request = &DescribeClusterIngressYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeClusterIngressYaml")
	return
}

func NewDescribeClusterIngressYamlResponse() (response *DescribeClusterIngressYamlResponse) {
	response = &DescribeClusterIngressYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群Ingress的yaml文件
func (c *Client) DescribeClusterIngressYaml(request *DescribeClusterIngressYamlRequest) (response *DescribeClusterIngressYamlResponse, err error) {
	if request == nil {
		request = NewDescribeClusterIngressYamlRequest()
	}
	response = NewDescribeClusterIngressYamlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskDnsEventDetailRequest() (request *DescribeRiskDnsEventDetailRequest) {
	request = &DescribeRiskDnsEventDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeRiskDnsEventDetail")
	return
}

func NewDescribeRiskDnsEventDetailResponse() (response *DescribeRiskDnsEventDetailResponse) {
	response = &DescribeRiskDnsEventDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询恶意请求事件详情
func (c *Client) DescribeRiskDnsEventDetail(request *DescribeRiskDnsEventDetailRequest) (response *DescribeRiskDnsEventDetailResponse, err error) {
	if request == nil {
		request = NewDescribeRiskDnsEventDetailRequest()
	}
	response = NewDescribeRiskDnsEventDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusListRequest() (request *DescribeVirusListRequest) {
	request = &DescribeVirusListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVirusList")
	return
}

func NewDescribeVirusListResponse() (response *DescribeVirusListResponse) {
	response = &DescribeVirusListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时文件查杀事件列表
func (c *Client) DescribeVirusList(request *DescribeVirusListRequest) (response *DescribeVirusListResponse, err error) {
	if request == nil {
		request = NewDescribeVirusListRequest()
	}
	response = NewDescribeVirusListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaliciousConnectionRuleSummaryRequest() (request *DescribeMaliciousConnectionRuleSummaryRequest) {
	request = &DescribeMaliciousConnectionRuleSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeMaliciousConnectionRuleSummary")
	return
}

func NewDescribeMaliciousConnectionRuleSummaryResponse() (response *DescribeMaliciousConnectionRuleSummaryResponse) {
	response = &DescribeMaliciousConnectionRuleSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恶意外连黑白名单统计，主要展示黑白名单的数据统计
func (c *Client) DescribeMaliciousConnectionRuleSummary(request *DescribeMaliciousConnectionRuleSummaryRequest) (response *DescribeMaliciousConnectionRuleSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeMaliciousConnectionRuleSummaryRequest()
	}
	response = NewDescribeMaliciousConnectionRuleSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEscapeEventsExportRequest() (request *DescribeEscapeEventsExportRequest) {
	request = &DescribeEscapeEventsExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeEscapeEventsExport")
	return
}

func NewDescribeEscapeEventsExportResponse() (response *DescribeEscapeEventsExportResponse) {
	response = &DescribeEscapeEventsExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器逃逸事件列表导出
func (c *Client) DescribeEscapeEventsExport(request *DescribeEscapeEventsExportRequest) (response *DescribeEscapeEventsExportResponse, err error) {
	if request == nil {
		request = NewDescribeEscapeEventsExportRequest()
	}
	response = NewDescribeEscapeEventsExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeESAggregationsRequest() (request *DescribeESAggregationsRequest) {
	request = &DescribeESAggregationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeESAggregations")
	return
}

func NewDescribeESAggregationsResponse() (response *DescribeESAggregationsResponse) {
	response = &DescribeESAggregationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ES字段聚合结果
func (c *Client) DescribeESAggregations(request *DescribeESAggregationsRequest) (response *DescribeESAggregationsResponse, err error) {
	if request == nil {
		request = NewDescribeESAggregationsRequest()
	}
	response = NewDescribeESAggregationsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEscapeEventsExportJobRequest() (request *CreateEscapeEventsExportJobRequest) {
	request = &CreateEscapeEventsExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateEscapeEventsExportJob")
	return
}

func NewCreateEscapeEventsExportJobResponse() (response *CreateEscapeEventsExportJobResponse) {
	response = &CreateEscapeEventsExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建逃逸事件导出异步任务
func (c *Client) CreateEscapeEventsExportJob(request *CreateEscapeEventsExportJobRequest) (response *CreateEscapeEventsExportJobResponse, err error) {
	if request == nil {
		request = NewCreateEscapeEventsExportJobRequest()
	}
	response = NewCreateEscapeEventsExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRefreshTaskRequest() (request *DescribeRefreshTaskRequest) {
	request = &DescribeRefreshTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeRefreshTask")
	return
}

func NewDescribeRefreshTaskResponse() (response *DescribeRefreshTaskResponse) {
	response = &DescribeRefreshTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询刷新任务
func (c *Client) DescribeRefreshTask(request *DescribeRefreshTaskRequest) (response *DescribeRefreshTaskResponse, err error) {
	if request == nil {
		request = NewDescribeRefreshTaskRequest()
	}
	response = NewDescribeRefreshTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSystemVulExportJobRequest() (request *CreateSystemVulExportJobRequest) {
	request = &CreateSystemVulExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateSystemVulExportJob")
	return
}

func NewCreateSystemVulExportJobResponse() (response *CreateSystemVulExportJobResponse) {
	response = &CreateSystemVulExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建系统漏洞导出任务
func (c *Client) CreateSystemVulExportJob(request *CreateSystemVulExportJobRequest) (response *CreateSystemVulExportJobResponse, err error) {
	if request == nil {
		request = NewCreateSystemVulExportJobRequest()
	}
	response = NewCreateSystemVulExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRiskListExportJobRequest() (request *CreateRiskListExportJobRequest) {
	request = &CreateRiskListExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateRiskListExportJob")
	return
}

func NewCreateRiskListExportJobResponse() (response *CreateRiskListExportJobResponse) {
	response = &CreateRiskListExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建风险项导出任务，可指定1个或多个风险导出
func (c *Client) CreateRiskListExportJob(request *CreateRiskListExportJobRequest) (response *CreateRiskListExportJobResponse, err error) {
	if request == nil {
		request = NewCreateRiskListExportJobRequest()
	}
	response = NewCreateRiskListExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageVulListRequest() (request *DescribeAssetImageVulListRequest) {
	request = &DescribeAssetImageVulListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageVulList")
	return
}

func NewDescribeAssetImageVulListResponse() (response *DescribeAssetImageVulListResponse) {
	response = &DescribeAssetImageVulListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像漏洞列表
func (c *Client) DescribeAssetImageVulList(request *DescribeAssetImageVulListRequest) (response *DescribeAssetImageVulListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageVulListRequest()
	}
	response = NewDescribeAssetImageVulListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEscapeRuleInfoRequest() (request *DescribeEscapeRuleInfoRequest) {
	request = &DescribeEscapeRuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeEscapeRuleInfo")
	return
}

func NewDescribeEscapeRuleInfoResponse() (response *DescribeEscapeRuleInfoResponse) {
	response = &DescribeEscapeRuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeEscapeRuleInfo 查询容器逃逸扫描规则信息
func (c *Client) DescribeEscapeRuleInfo(request *DescribeEscapeRuleInfoRequest) (response *DescribeEscapeRuleInfoResponse, err error) {
	if request == nil {
		request = NewDescribeEscapeRuleInfoRequest()
	}
	response = NewDescribeEscapeRuleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClustersNamespacesRequest() (request *DescribeClustersNamespacesRequest) {
	request = &DescribeClustersNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeClustersNamespaces")
	return
}

func NewDescribeClustersNamespacesResponse() (response *DescribeClustersNamespacesResponse) {
	response = &DescribeClustersNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群的命名空间
func (c *Client) DescribeClustersNamespaces(request *DescribeClustersNamespacesRequest) (response *DescribeClustersNamespacesResponse, err error) {
	if request == nil {
		request = NewDescribeClustersNamespacesRequest()
	}
	response = NewDescribeClustersNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetPortListRequest() (request *DescribeAssetPortListRequest) {
	request = &DescribeAssetPortListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetPortList")
	return
}

func NewDescribeAssetPortListResponse() (response *DescribeAssetPortListResponse) {
	response = &DescribeAssetPortListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询端口占用列表
func (c *Client) DescribeAssetPortList(request *DescribeAssetPortListRequest) (response *DescribeAssetPortListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetPortListRequest()
	}
	response = NewDescribeAssetPortListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellWhiteListDetailRequest() (request *DescribeReverseShellWhiteListDetailRequest) {
	request = &DescribeReverseShellWhiteListDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeReverseShellWhiteListDetail")
	return
}

func NewDescribeReverseShellWhiteListDetailResponse() (response *DescribeReverseShellWhiteListDetailResponse) {
	response = &DescribeReverseShellWhiteListDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时反弹shell白名单详细信息
func (c *Client) DescribeReverseShellWhiteListDetail(request *DescribeReverseShellWhiteListDetailRequest) (response *DescribeReverseShellWhiteListDetailResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellWhiteListDetailRequest()
	}
	response = NewDescribeReverseShellWhiteListDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnfinishRefreshTaskRequest() (request *DescribeUnfinishRefreshTaskRequest) {
	request = &DescribeUnfinishRefreshTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeUnfinishRefreshTask")
	return
}

func NewDescribeUnfinishRefreshTaskResponse() (response *DescribeUnfinishRefreshTaskResponse) {
	response = &DescribeUnfinishRefreshTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询未完成的刷新资产任务信息
func (c *Client) DescribeUnfinishRefreshTask(request *DescribeUnfinishRefreshTaskRequest) (response *DescribeUnfinishRefreshTaskResponse, err error) {
	if request == nil {
		request = NewDescribeUnfinishRefreshTaskRequest()
	}
	response = NewDescribeUnfinishRefreshTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetAppServiceListRequest() (request *DescribeAssetAppServiceListRequest) {
	request = &DescribeAssetAppServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetAppServiceList")
	return
}

func NewDescribeAssetAppServiceListResponse() (response *DescribeAssetAppServiceListResponse) {
	response = &DescribeAssetAppServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询app服务列表
func (c *Client) DescribeAssetAppServiceList(request *DescribeAssetAppServiceListRequest) (response *DescribeAssetAppServiceListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetAppServiceListRequest()
	}
	response = NewDescribeAssetAppServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterServiceYamlRequest() (request *DescribeClusterServiceYamlRequest) {
	request = &DescribeClusterServiceYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeClusterServiceYaml")
	return
}

func NewDescribeClusterServiceYamlResponse() (response *DescribeClusterServiceYamlResponse) {
	response = &DescribeClusterServiceYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群service的yaml文件
func (c *Client) DescribeClusterServiceYaml(request *DescribeClusterServiceYamlRequest) (response *DescribeClusterServiceYamlResponse, err error) {
	if request == nil {
		request = NewDescribeClusterServiceYamlRequest()
	}
	response = NewDescribeClusterServiceYamlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusSampleDownloadUrlRequest() (request *DescribeVirusSampleDownloadUrlRequest) {
	request = &DescribeVirusSampleDownloadUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVirusSampleDownloadUrl")
	return
}

func NewDescribeVirusSampleDownloadUrlResponse() (response *DescribeVirusSampleDownloadUrlResponse) {
	response = &DescribeVirusSampleDownloadUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询木马样本下载url
func (c *Client) DescribeVirusSampleDownloadUrl(request *DescribeVirusSampleDownloadUrlRequest) (response *DescribeVirusSampleDownloadUrlResponse, err error) {
	if request == nil {
		request = NewDescribeVirusSampleDownloadUrlRequest()
	}
	response = NewDescribeVirusSampleDownloadUrlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCompliancePolicyItemAffectedAssetListRequest() (request *DescribeCompliancePolicyItemAffectedAssetListRequest) {
	request = &DescribeCompliancePolicyItemAffectedAssetListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeCompliancePolicyItemAffectedAssetList")
	return
}

func NewDescribeCompliancePolicyItemAffectedAssetListResponse() (response *DescribeCompliancePolicyItemAffectedAssetListResponse) {
	response = &DescribeCompliancePolicyItemAffectedAssetListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按照 检测项 → 资产 的两级层次展开的第二层级：资产层级。
func (c *Client) DescribeCompliancePolicyItemAffectedAssetList(request *DescribeCompliancePolicyItemAffectedAssetListRequest) (response *DescribeCompliancePolicyItemAffectedAssetListResponse, err error) {
	if request == nil {
		request = NewDescribeCompliancePolicyItemAffectedAssetListRequest()
	}
	response = NewDescribeCompliancePolicyItemAffectedAssetListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessControlRulesExportRequest() (request *DescribeAccessControlRulesExportRequest) {
	request = &DescribeAccessControlRulesExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAccessControlRulesExport")
	return
}

func NewDescribeAccessControlRulesExportResponse() (response *DescribeAccessControlRulesExportResponse) {
	response = &DescribeAccessControlRulesExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时访问控制策略列表导出
func (c *Client) DescribeAccessControlRulesExport(request *DescribeAccessControlRulesExportRequest) (response *DescribeAccessControlRulesExportResponse, err error) {
	if request == nil {
		request = NewDescribeAccessControlRulesExportRequest()
	}
	response = NewDescribeAccessControlRulesExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistrySummaryRequest() (request *DescribeAssetImageRegistrySummaryRequest) {
	request = &DescribeAssetImageRegistrySummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageRegistrySummary")
	return
}

func NewDescribeAssetImageRegistrySummaryResponse() (response *DescribeAssetImageRegistrySummaryResponse) {
	response = &DescribeAssetImageRegistrySummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库查询镜像统计信息
func (c *Client) DescribeAssetImageRegistrySummary(request *DescribeAssetImageRegistrySummaryRequest) (response *DescribeAssetImageRegistrySummaryResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistrySummaryRequest()
	}
	response = NewDescribeAssetImageRegistrySummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusManualScanEstimateTimeoutRequest() (request *DescribeVirusManualScanEstimateTimeoutRequest) {
	request = &DescribeVirusManualScanEstimateTimeoutRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVirusManualScanEstimateTimeout")
	return
}

func NewDescribeVirusManualScanEstimateTimeoutResponse() (response *DescribeVirusManualScanEstimateTimeoutResponse) {
	response = &DescribeVirusManualScanEstimateTimeoutResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询木马一键检测预估超时时间
func (c *Client) DescribeVirusManualScanEstimateTimeout(request *DescribeVirusManualScanEstimateTimeoutRequest) (response *DescribeVirusManualScanEstimateTimeoutResponse, err error) {
	if request == nil {
		request = NewDescribeVirusManualScanEstimateTimeoutRequest()
	}
	response = NewDescribeVirusManualScanEstimateTimeoutResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageDenyRuleListRequest() (request *DescribeImageDenyRuleListRequest) {
	request = &DescribeImageDenyRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeImageDenyRuleList")
	return
}

func NewDescribeImageDenyRuleListResponse() (response *DescribeImageDenyRuleListResponse) {
	response = &DescribeImageDenyRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像拦截规则列表
func (c *Client) DescribeImageDenyRuleList(request *DescribeImageDenyRuleListRequest) (response *DescribeImageDenyRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeImageDenyRuleListRequest()
	}
	response = NewDescribeImageDenyRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaliciousConnectionWhiteListRequest() (request *DescribeMaliciousConnectionWhiteListRequest) {
	request = &DescribeMaliciousConnectionWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeMaliciousConnectionWhiteList")
	return
}

func NewDescribeMaliciousConnectionWhiteListResponse() (response *DescribeMaliciousConnectionWhiteListResponse) {
	response = &DescribeMaliciousConnectionWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询恶意外连白名单
func (c *Client) DescribeMaliciousConnectionWhiteList(request *DescribeMaliciousConnectionWhiteListRequest) (response *DescribeMaliciousConnectionWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeMaliciousConnectionWhiteListRequest()
	}
	response = NewDescribeMaliciousConnectionWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryRiskListExportRequest() (request *DescribeAssetImageRegistryRiskListExportRequest) {
	request = &DescribeAssetImageRegistryRiskListExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageRegistryRiskListExport")
	return
}

func NewDescribeAssetImageRegistryRiskListExportResponse() (response *DescribeAssetImageRegistryRiskListExportResponse) {
	response = &DescribeAssetImageRegistryRiskListExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库敏感信息列表导出
func (c *Client) DescribeAssetImageRegistryRiskListExport(request *DescribeAssetImageRegistryRiskListExportRequest) (response *DescribeAssetImageRegistryRiskListExportResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryRiskListExportRequest()
	}
	response = NewDescribeAssetImageRegistryRiskListExportResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVulExportJobRequest() (request *CreateVulExportJobRequest) {
	request = &CreateVulExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateVulExportJob")
	return
}

func NewCreateVulExportJobResponse() (response *CreateVulExportJobResponse) {
	response = &CreateVulExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询本地镜像漏洞列表导出
func (c *Client) CreateVulExportJob(request *CreateVulExportJobRequest) (response *CreateVulExportJobResponse, err error) {
	if request == nil {
		request = NewCreateVulExportJobRequest()
	}
	response = NewCreateVulExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAutoStartSecurityServiceStateRequest() (request *ModifyAutoStartSecurityServiceStateRequest) {
	request = &ModifyAutoStartSecurityServiceStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "ModifyAutoStartSecurityServiceState")
	return
}

func NewModifyAutoStartSecurityServiceStateResponse() (response *ModifyAutoStartSecurityServiceStateResponse) {
	response = &ModifyAutoStartSecurityServiceStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改新增机器是否自动开通容器安全服务
func (c *Client) ModifyAutoStartSecurityServiceState(request *ModifyAutoStartSecurityServiceStateRequest) (response *ModifyAutoStartSecurityServiceStateResponse, err error) {
	if request == nil {
		request = NewModifyAutoStartSecurityServiceStateRequest()
	}
	response = NewModifyAutoStartSecurityServiceStateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEscapeEventTypeSummaryRequest() (request *DescribeEscapeEventTypeSummaryRequest) {
	request = &DescribeEscapeEventTypeSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeEscapeEventTypeSummary")
	return
}

func NewDescribeEscapeEventTypeSummaryResponse() (response *DescribeEscapeEventTypeSummaryResponse) {
	response = &DescribeEscapeEventTypeSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 统计容器逃逸各事件类型和待处理事件数
func (c *Client) DescribeEscapeEventTypeSummary(request *DescribeEscapeEventTypeSummaryRequest) (response *DescribeEscapeEventTypeSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeEscapeEventTypeSummaryRequest()
	}
	response = NewDescribeEscapeEventTypeSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskDnsSummaryRequest() (request *DescribeRiskDnsSummaryRequest) {
	request = &DescribeRiskDnsSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeRiskDnsSummary")
	return
}

func NewDescribeRiskDnsSummaryResponse() (response *DescribeRiskDnsSummaryResponse) {
	response = &DescribeRiskDnsSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询恶意请求事件统计
func (c *Client) DescribeRiskDnsSummary(request *DescribeRiskDnsSummaryRequest) (response *DescribeRiskDnsSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeRiskDnsSummaryRequest()
	}
	response = NewDescribeRiskDnsSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryRiskInfoListRequest() (request *DescribeAssetImageRegistryRiskInfoListRequest) {
	request = &DescribeAssetImageRegistryRiskInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageRegistryRiskInfoList")
	return
}

func NewDescribeAssetImageRegistryRiskInfoListResponse() (response *DescribeAssetImageRegistryRiskInfoListResponse) {
	response = &DescribeAssetImageRegistryRiskInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库查询镜像高危行为列表
func (c *Client) DescribeAssetImageRegistryRiskInfoList(request *DescribeAssetImageRegistryRiskInfoListRequest) (response *DescribeAssetImageRegistryRiskInfoListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryRiskInfoListRequest()
	}
	response = NewDescribeAssetImageRegistryRiskInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAffectedNodeListRequest() (request *DescribeAffectedNodeListRequest) {
	request = &DescribeAffectedNodeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAffectedNodeList")
	return
}

func NewDescribeAffectedNodeListResponse() (response *DescribeAffectedNodeListResponse) {
	response = &DescribeAffectedNodeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点类型的影响范围，返回节点列表
func (c *Client) DescribeAffectedNodeList(request *DescribeAffectedNodeListRequest) (response *DescribeAffectedNodeListResponse, err error) {
	if request == nil {
		request = NewDescribeAffectedNodeListRequest()
	}
	response = NewDescribeAffectedNodeListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusAutoIsolateSettingRequest() (request *DescribeVirusAutoIsolateSettingRequest) {
	request = &DescribeVirusAutoIsolateSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVirusAutoIsolateSetting")
	return
}

func NewDescribeVirusAutoIsolateSettingResponse() (response *DescribeVirusAutoIsolateSettingResponse) {
	response = &DescribeVirusAutoIsolateSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询木马自动隔离设置
func (c *Client) DescribeVirusAutoIsolateSetting(request *DescribeVirusAutoIsolateSettingRequest) (response *DescribeVirusAutoIsolateSettingResponse, err error) {
	if request == nil {
		request = NewDescribeVirusAutoIsolateSettingRequest()
	}
	response = NewDescribeVirusAutoIsolateSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserWorkloadKindsRequest() (request *DescribeUserWorkloadKindsRequest) {
	request = &DescribeUserWorkloadKindsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeUserWorkloadKinds")
	return
}

func NewDescribeUserWorkloadKindsResponse() (response *DescribeUserWorkloadKindsResponse) {
	response = &DescribeUserWorkloadKindsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 枚举用户所有的工作负载类型
func (c *Client) DescribeUserWorkloadKinds(request *DescribeUserWorkloadKindsRequest) (response *DescribeUserWorkloadKindsResponse, err error) {
	if request == nil {
		request = NewDescribeUserWorkloadKindsRequest()
	}
	response = NewDescribeUserWorkloadKindsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskTargetCountRequest() (request *DescribeRiskTargetCountRequest) {
	request = &DescribeRiskTargetCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeRiskTargetCount")
	return
}

func NewDescribeRiskTargetCountResponse() (response *DescribeRiskTargetCountResponse) {
	response = &DescribeRiskTargetCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按检查对象展示风险数量
func (c *Client) DescribeRiskTargetCount(request *DescribeRiskTargetCountRequest) (response *DescribeRiskTargetCountResponse, err error) {
	if request == nil {
		request = NewDescribeRiskTargetCountRequest()
	}
	response = NewDescribeRiskTargetCountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEventEscapeImageStatusRequest() (request *ModifyEventEscapeImageStatusRequest) {
	request = &ModifyEventEscapeImageStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "ModifyEventEscapeImageStatus")
	return
}

func NewModifyEventEscapeImageStatusResponse() (response *ModifyEventEscapeImageStatusResponse) {
	response = &ModifyEventEscapeImageStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改风险容器镜像状态
func (c *Client) ModifyEventEscapeImageStatus(request *ModifyEventEscapeImageStatusRequest) (response *ModifyEventEscapeImageStatusResponse, err error) {
	if request == nil {
		request = NewModifyEventEscapeImageStatusRequest()
	}
	response = NewModifyEventEscapeImageStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetHostDetailRequest() (request *DescribeAssetHostDetailRequest) {
	request = &DescribeAssetHostDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetHostDetail")
	return
}

func NewDescribeAssetHostDetailResponse() (response *DescribeAssetHostDetailResponse) {
	response = &DescribeAssetHostDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主机详细信息
func (c *Client) DescribeAssetHostDetail(request *DescribeAssetHostDetailRequest) (response *DescribeAssetHostDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAssetHostDetailRequest()
	}
	response = NewDescribeAssetHostDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetWebServiceListRequest() (request *DescribeAssetWebServiceListRequest) {
	request = &DescribeAssetWebServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetWebServiceList")
	return
}

func NewDescribeAssetWebServiceListResponse() (response *DescribeAssetWebServiceListResponse) {
	response = &DescribeAssetWebServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询web服务列表
func (c *Client) DescribeAssetWebServiceList(request *DescribeAssetWebServiceListRequest) (response *DescribeAssetWebServiceListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetWebServiceListRequest()
	}
	response = NewDescribeAssetWebServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaliciousConnectionBlackListRequest() (request *DescribeMaliciousConnectionBlackListRequest) {
	request = &DescribeMaliciousConnectionBlackListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeMaliciousConnectionBlackList")
	return
}

func NewDescribeMaliciousConnectionBlackListResponse() (response *DescribeMaliciousConnectionBlackListResponse) {
	response = &DescribeMaliciousConnectionBlackListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询恶意外连黑名单
func (c *Client) DescribeMaliciousConnectionBlackList(request *DescribeMaliciousConnectionBlackListRequest) (response *DescribeMaliciousConnectionBlackListResponse, err error) {
	if request == nil {
		request = NewDescribeMaliciousConnectionBlackListRequest()
	}
	response = NewDescribeMaliciousConnectionBlackListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIndexListRequest() (request *DescribeIndexListRequest) {
	request = &DescribeIndexListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeIndexList")
	return
}

func NewDescribeIndexListResponse() (response *DescribeIndexListResponse) {
	response = &DescribeIndexListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取索引列表
func (c *Client) DescribeIndexList(request *DescribeIndexListRequest) (response *DescribeIndexListResponse, err error) {
	if request == nil {
		request = NewDescribeIndexListRequest()
	}
	response = NewDescribeIndexListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserPodListRequest() (request *DescribeUserPodListRequest) {
	request = &DescribeUserPodListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeUserPodList")
	return
}

func NewDescribeUserPodListResponse() (response *DescribeUserPodListResponse) {
	response = &DescribeUserPodListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的pod列表
func (c *Client) DescribeUserPodList(request *DescribeUserPodListRequest) (response *DescribeUserPodListResponse, err error) {
	if request == nil {
		request = NewDescribeUserPodListRequest()
	}
	response = NewDescribeUserPodListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellEventsExportRequest() (request *DescribeReverseShellEventsExportRequest) {
	request = &DescribeReverseShellEventsExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeReverseShellEventsExport")
	return
}

func NewDescribeReverseShellEventsExportResponse() (response *DescribeReverseShellEventsExportResponse) {
	response = &DescribeReverseShellEventsExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时反弹shell事件列表信息导出
func (c *Client) DescribeReverseShellEventsExport(request *DescribeReverseShellEventsExportRequest) (response *DescribeReverseShellEventsExportResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellEventsExportRequest()
	}
	response = NewDescribeReverseShellEventsExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetDBServiceListRequest() (request *DescribeAssetDBServiceListRequest) {
	request = &DescribeAssetDBServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetDBServiceList")
	return
}

func NewDescribeAssetDBServiceListResponse() (response *DescribeAssetDBServiceListResponse) {
	response = &DescribeAssetDBServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询db服务列表
func (c *Client) DescribeAssetDBServiceList(request *DescribeAssetDBServiceListRequest) (response *DescribeAssetDBServiceListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetDBServiceListRequest()
	}
	response = NewDescribeAssetDBServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessControlRuleDetailRequest() (request *DescribeAccessControlRuleDetailRequest) {
	request = &DescribeAccessControlRuleDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAccessControlRuleDetail")
	return
}

func NewDescribeAccessControlRuleDetailResponse() (response *DescribeAccessControlRuleDetailResponse) {
	response = &DescribeAccessControlRuleDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时访问控制策略详细信息
func (c *Client) DescribeAccessControlRuleDetail(request *DescribeAccessControlRuleDetailRequest) (response *DescribeAccessControlRuleDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAccessControlRuleDetailRequest()
	}
	response = NewDescribeAccessControlRuleDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageVirusListExportRequest() (request *DescribeAssetImageVirusListExportRequest) {
	request = &DescribeAssetImageVirusListExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageVirusListExport")
	return
}

func NewDescribeAssetImageVirusListExportResponse() (response *DescribeAssetImageVirusListExportResponse) {
	response = &DescribeAssetImageVirusListExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像木马列表导出
func (c *Client) DescribeAssetImageVirusListExport(request *DescribeAssetImageVirusListExportRequest) (response *DescribeAssetImageVirusListExportResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageVirusListExportRequest()
	}
	response = NewDescribeAssetImageVirusListExportResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVirusExportJobRequest() (request *CreateVirusExportJobRequest) {
	request = &CreateVirusExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateVirusExportJob")
	return
}

func NewCreateVirusExportJobResponse() (response *CreateVirusExportJobResponse) {
	response = &CreateVirusExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时文件查杀事件列表导出
func (c *Client) CreateVirusExportJob(request *CreateVirusExportJobRequest) (response *CreateVirusExportJobResponse, err error) {
	if request == nil {
		request = NewCreateVirusExportJobRequest()
	}
	response = NewCreateVirusExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryVulListExportRequest() (request *DescribeAssetImageRegistryVulListExportRequest) {
	request = &DescribeAssetImageRegistryVulListExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageRegistryVulListExport")
	return
}

func NewDescribeAssetImageRegistryVulListExportResponse() (response *DescribeAssetImageRegistryVulListExportResponse) {
	response = &DescribeAssetImageRegistryVulListExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库漏洞列表导出
func (c *Client) DescribeAssetImageRegistryVulListExport(request *DescribeAssetImageRegistryVulListExportRequest) (response *DescribeAssetImageRegistryVulListExportResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryVulListExportRequest()
	}
	response = NewDescribeAssetImageRegistryVulListExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbnormalProcessDetailRequest() (request *DescribeAbnormalProcessDetailRequest) {
	request = &DescribeAbnormalProcessDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAbnormalProcessDetail")
	return
}

func NewDescribeAbnormalProcessDetailResponse() (response *DescribeAbnormalProcessDetailResponse) {
	response = &DescribeAbnormalProcessDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时异常进程事件详细信息
func (c *Client) DescribeAbnormalProcessDetail(request *DescribeAbnormalProcessDetailRequest) (response *DescribeAbnormalProcessDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAbnormalProcessDetailRequest()
	}
	response = NewDescribeAbnormalProcessDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryVirusListExportRequest() (request *DescribeAssetImageRegistryVirusListExportRequest) {
	request = &DescribeAssetImageRegistryVirusListExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageRegistryVirusListExport")
	return
}

func NewDescribeAssetImageRegistryVirusListExportResponse() (response *DescribeAssetImageRegistryVirusListExportResponse) {
	response = &DescribeAssetImageRegistryVirusListExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库木马信息列表导出
func (c *Client) DescribeAssetImageRegistryVirusListExport(request *DescribeAssetImageRegistryVirusListExportRequest) (response *DescribeAssetImageRegistryVirusListExportResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryVirusListExportRequest()
	}
	response = NewDescribeAssetImageRegistryVirusListExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserPurchaseInfoRequest() (request *DescribeUserPurchaseInfoRequest) {
	request = &DescribeUserPurchaseInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeUserPurchaseInfo")
	return
}

func NewDescribeUserPurchaseInfoResponse() (response *DescribeUserPurchaseInfoResponse) {
	response = &DescribeUserPurchaseInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户购买容器安全服务列表
func (c *Client) DescribeUserPurchaseInfo(request *DescribeUserPurchaseInfoRequest) (response *DescribeUserPurchaseInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUserPurchaseInfoRequest()
	}
	response = NewDescribeUserPurchaseInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterServiceListRequest() (request *DescribeClusterServiceListRequest) {
	request = &DescribeClusterServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeClusterServiceList")
	return
}

func NewDescribeClusterServiceListResponse() (response *DescribeClusterServiceListResponse) {
	response = &DescribeClusterServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群service列表
func (c *Client) DescribeClusterServiceList(request *DescribeClusterServiceListRequest) (response *DescribeClusterServiceListResponse, err error) {
	if request == nil {
		request = NewDescribeClusterServiceListRequest()
	}
	response = NewDescribeClusterServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebVulListRequest() (request *DescribeWebVulListRequest) {
	request = &DescribeWebVulListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeWebVulList")
	return
}

func NewDescribeWebVulListResponse() (response *DescribeWebVulListResponse) {
	response = &DescribeWebVulListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询web应用漏洞列表
func (c *Client) DescribeWebVulList(request *DescribeWebVulListRequest) (response *DescribeWebVulListResponse, err error) {
	if request == nil {
		request = NewDescribeWebVulListRequest()
	}
	response = NewDescribeWebVulListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageDenyRuleDetailRequest() (request *DescribeImageDenyRuleDetailRequest) {
	request = &DescribeImageDenyRuleDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeImageDenyRuleDetail")
	return
}

func NewDescribeImageDenyRuleDetailResponse() (response *DescribeImageDenyRuleDetailResponse) {
	response = &DescribeImageDenyRuleDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像拦截规则详情
func (c *Client) DescribeImageDenyRuleDetail(request *DescribeImageDenyRuleDetailRequest) (response *DescribeImageDenyRuleDetailResponse, err error) {
	if request == nil {
		request = NewDescribeImageDenyRuleDetailRequest()
	}
	response = NewDescribeImageDenyRuleDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageBindRuleInfoRequest() (request *DescribeAssetImageBindRuleInfoRequest) {
	request = &DescribeAssetImageBindRuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageBindRuleInfo")
	return
}

func NewDescribeAssetImageBindRuleInfoResponse() (response *DescribeAssetImageBindRuleInfoResponse) {
	response = &DescribeAssetImageBindRuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像绑定规则列表信息，包含运行时访问控制和异常进程公用
func (c *Client) DescribeAssetImageBindRuleInfo(request *DescribeAssetImageBindRuleInfoRequest) (response *DescribeAssetImageBindRuleInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageBindRuleInfoRequest()
	}
	response = NewDescribeAssetImageBindRuleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSystemPolicyRequest() (request *DescribeSystemPolicyRequest) {
	request = &DescribeSystemPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeSystemPolicy")
	return
}

func NewDescribeSystemPolicyResponse() (response *DescribeSystemPolicyResponse) {
	response = &DescribeSystemPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时系统策略列表
func (c *Client) DescribeSystemPolicy(request *DescribeSystemPolicyRequest) (response *DescribeSystemPolicyResponse, err error) {
	if request == nil {
		request = NewDescribeSystemPolicyRequest()
	}
	response = NewDescribeSystemPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerSecEventSummaryRequest() (request *DescribeContainerSecEventSummaryRequest) {
	request = &DescribeContainerSecEventSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeContainerSecEventSummary")
	return
}

func NewDescribeContainerSecEventSummaryResponse() (response *DescribeContainerSecEventSummaryResponse) {
	response = &DescribeContainerSecEventSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器安全未处理事件概览
func (c *Client) DescribeContainerSecEventSummary(request *DescribeContainerSecEventSummaryRequest) (response *DescribeContainerSecEventSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeContainerSecEventSummaryRequest()
	}
	response = NewDescribeContainerSecEventSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryListExportRequest() (request *DescribeAssetImageRegistryListExportRequest) {
	request = &DescribeAssetImageRegistryListExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageRegistryListExport")
	return
}

func NewDescribeAssetImageRegistryListExportResponse() (response *DescribeAssetImageRegistryListExportResponse) {
	response = &DescribeAssetImageRegistryListExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库镜像列表导出
func (c *Client) DescribeAssetImageRegistryListExport(request *DescribeAssetImageRegistryListExportRequest) (response *DescribeAssetImageRegistryListExportResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryListExportRequest()
	}
	response = NewDescribeAssetImageRegistryListExportResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRiskDnsEventExportJobRequest() (request *CreateRiskDnsEventExportJobRequest) {
	request = &CreateRiskDnsEventExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateRiskDnsEventExportJob")
	return
}

func NewCreateRiskDnsEventExportJobResponse() (response *CreateRiskDnsEventExportJobResponse) {
	response = &CreateRiskDnsEventExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建恶意请求事件导出任务
func (c *Client) CreateRiskDnsEventExportJob(request *CreateRiskDnsEventExportJobRequest) (response *CreateRiskDnsEventExportJobResponse, err error) {
	if request == nil {
		request = NewCreateRiskDnsEventExportJobRequest()
	}
	response = NewCreateRiskDnsEventExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComplianceTaskAssetSummaryRequest() (request *DescribeComplianceTaskAssetSummaryRequest) {
	request = &DescribeComplianceTaskAssetSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeComplianceTaskAssetSummary")
	return
}

func NewDescribeComplianceTaskAssetSummaryResponse() (response *DescribeComplianceTaskAssetSummaryResponse) {
	response = &DescribeComplianceTaskAssetSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询上次任务的资产通过率汇总信息
func (c *Client) DescribeComplianceTaskAssetSummary(request *DescribeComplianceTaskAssetSummaryRequest) (response *DescribeComplianceTaskAssetSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeComplianceTaskAssetSummaryRequest()
	}
	response = NewDescribeComplianceTaskAssetSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusAutoIsolateSampleDetailRequest() (request *DescribeVirusAutoIsolateSampleDetailRequest) {
	request = &DescribeVirusAutoIsolateSampleDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVirusAutoIsolateSampleDetail")
	return
}

func NewDescribeVirusAutoIsolateSampleDetailResponse() (response *DescribeVirusAutoIsolateSampleDetailResponse) {
	response = &DescribeVirusAutoIsolateSampleDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询木马自动隔离样本详情
func (c *Client) DescribeVirusAutoIsolateSampleDetail(request *DescribeVirusAutoIsolateSampleDetailRequest) (response *DescribeVirusAutoIsolateSampleDetailResponse, err error) {
	if request == nil {
		request = NewDescribeVirusAutoIsolateSampleDetailRequest()
	}
	response = NewDescribeVirusAutoIsolateSampleDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryListRequest() (request *DescribeAssetImageRegistryListRequest) {
	request = &DescribeAssetImageRegistryListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageRegistryList")
	return
}

func NewDescribeAssetImageRegistryListResponse() (response *DescribeAssetImageRegistryListResponse) {
	response = &DescribeAssetImageRegistryListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库查询镜像仓库列表
func (c *Client) DescribeAssetImageRegistryList(request *DescribeAssetImageRegistryListRequest) (response *DescribeAssetImageRegistryListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryListRequest()
	}
	response = NewDescribeAssetImageRegistryListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExportJobResultRequest() (request *DescribeExportJobResultRequest) {
	request = &DescribeExportJobResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeExportJobResult")
	return
}

func NewDescribeExportJobResultResponse() (response *DescribeExportJobResultResponse) {
	response = &DescribeExportJobResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询导出任务的结果
func (c *Client) DescribeExportJobResult(request *DescribeExportJobResultRequest) (response *DescribeExportJobResultResponse, err error) {
	if request == nil {
		request = NewDescribeExportJobResultRequest()
	}
	response = NewDescribeExportJobResultResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskListRequest() (request *DescribeRiskListRequest) {
	request = &DescribeRiskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeRiskList")
	return
}

func NewDescribeRiskListResponse() (response *DescribeRiskListResponse) {
	response = &DescribeRiskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询最近一次任务发现的风险项的信息列表，支持根据特殊字段进行过滤
func (c *Client) DescribeRiskList(request *DescribeRiskListRequest) (response *DescribeRiskListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskListRequest()
	}
	response = NewDescribeRiskListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateK8sApiAbnormalEventExportJobRequest() (request *CreateK8sApiAbnormalEventExportJobRequest) {
	request = &CreateK8sApiAbnormalEventExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateK8sApiAbnormalEventExportJob")
	return
}

func NewCreateK8sApiAbnormalEventExportJobResponse() (response *CreateK8sApiAbnormalEventExportJobResponse) {
	response = &CreateK8sApiAbnormalEventExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建k8s api异常事件导出任务
func (c *Client) CreateK8sApiAbnormalEventExportJob(request *CreateK8sApiAbnormalEventExportJobRequest) (response *CreateK8sApiAbnormalEventExportJobResponse, err error) {
	if request == nil {
		request = NewCreateK8sApiAbnormalEventExportJobRequest()
	}
	response = NewCreateK8sApiAbnormalEventExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterRiskRequest() (request *DescribeClusterRiskRequest) {
	request = &DescribeClusterRiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeClusterRisk")
	return
}

func NewDescribeClusterRiskResponse() (response *DescribeClusterRiskResponse) {
	response = &DescribeClusterRiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按集群类别查询风险数量
func (c *Client) DescribeClusterRisk(request *DescribeClusterRiskRequest) (response *DescribeClusterRiskResponse, err error) {
	if request == nil {
		request = NewDescribeClusterRiskRequest()
	}
	response = NewDescribeClusterRiskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulIgnoreLocalImageListRequest() (request *DescribeVulIgnoreLocalImageListRequest) {
	request = &DescribeVulIgnoreLocalImageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVulIgnoreLocalImageList")
	return
}

func NewDescribeVulIgnoreLocalImageListResponse() (response *DescribeVulIgnoreLocalImageListResponse) {
	response = &DescribeVulIgnoreLocalImageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞扫描忽略的本地镜像列表
func (c *Client) DescribeVulIgnoreLocalImageList(request *DescribeVulIgnoreLocalImageListRequest) (response *DescribeVulIgnoreLocalImageListResponse, err error) {
	if request == nil {
		request = NewDescribeVulIgnoreLocalImageListRequest()
	}
	response = NewDescribeVulIgnoreLocalImageListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulIgnoreRegistryImageListRequest() (request *DescribeVulIgnoreRegistryImageListRequest) {
	request = &DescribeVulIgnoreRegistryImageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVulIgnoreRegistryImageList")
	return
}

func NewDescribeVulIgnoreRegistryImageListResponse() (response *DescribeVulIgnoreRegistryImageListResponse) {
	response = &DescribeVulIgnoreRegistryImageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞扫描忽略的仓库镜像列表
func (c *Client) DescribeVulIgnoreRegistryImageList(request *DescribeVulIgnoreRegistryImageListRequest) (response *DescribeVulIgnoreRegistryImageListResponse, err error) {
	if request == nil {
		request = NewDescribeVulIgnoreRegistryImageListRequest()
	}
	response = NewDescribeVulIgnoreRegistryImageListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskSyscallDetailRequest() (request *DescribeRiskSyscallDetailRequest) {
	request = &DescribeRiskSyscallDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeRiskSyscallDetail")
	return
}

func NewDescribeRiskSyscallDetailResponse() (response *DescribeRiskSyscallDetailResponse) {
	response = &DescribeRiskSyscallDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时高危系统调用详细信息
func (c *Client) DescribeRiskSyscallDetail(request *DescribeRiskSyscallDetailRequest) (response *DescribeRiskSyscallDetailResponse, err error) {
	if request == nil {
		request = NewDescribeRiskSyscallDetailRequest()
	}
	response = NewDescribeRiskSyscallDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetContainerDetailRequest() (request *DescribeAssetContainerDetailRequest) {
	request = &DescribeAssetContainerDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetContainerDetail")
	return
}

func NewDescribeAssetContainerDetailResponse() (response *DescribeAssetContainerDetailResponse) {
	response = &DescribeAssetContainerDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器详细信息
func (c *Client) DescribeAssetContainerDetail(request *DescribeAssetContainerDetailRequest) (response *DescribeAssetContainerDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAssetContainerDetailRequest()
	}
	response = NewDescribeAssetContainerDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskSyscallEventsExportRequest() (request *DescribeRiskSyscallEventsExportRequest) {
	request = &DescribeRiskSyscallEventsExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeRiskSyscallEventsExport")
	return
}

func NewDescribeRiskSyscallEventsExportResponse() (response *DescribeRiskSyscallEventsExportResponse) {
	response = &DescribeRiskSyscallEventsExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时高危系统调用列表导出
func (c *Client) DescribeRiskSyscallEventsExport(request *DescribeRiskSyscallEventsExportRequest) (response *DescribeRiskSyscallEventsExportResponse, err error) {
	if request == nil {
		request = NewDescribeRiskSyscallEventsExportRequest()
	}
	response = NewDescribeRiskSyscallEventsExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComplianceAssetListRequest() (request *DescribeComplianceAssetListRequest) {
	request = &DescribeComplianceAssetListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeComplianceAssetList")
	return
}

func NewDescribeComplianceAssetListResponse() (response *DescribeComplianceAssetListResponse) {
	response = &DescribeComplianceAssetListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某类资产的列表
func (c *Client) DescribeComplianceAssetList(request *DescribeComplianceAssetListRequest) (response *DescribeComplianceAssetListResponse, err error) {
	if request == nil {
		request = NewDescribeComplianceAssetListRequest()
	}
	response = NewDescribeComplianceAssetListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusTaskListRequest() (request *DescribeVirusTaskListRequest) {
	request = &DescribeVirusTaskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVirusTaskList")
	return
}

func NewDescribeVirusTaskListResponse() (response *DescribeVirusTaskListResponse) {
	response = &DescribeVirusTaskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时查询文件查杀任务列表
func (c *Client) DescribeVirusTaskList(request *DescribeVirusTaskListRequest) (response *DescribeVirusTaskListResponse, err error) {
	if request == nil {
		request = NewDescribeVirusTaskListRequest()
	}
	response = NewDescribeVirusTaskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusMonitorSettingRequest() (request *DescribeVirusMonitorSettingRequest) {
	request = &DescribeVirusMonitorSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVirusMonitorSetting")
	return
}

func NewDescribeVirusMonitorSettingResponse() (response *DescribeVirusMonitorSettingResponse) {
	response = &DescribeVirusMonitorSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时查询文件查杀实时监控设置
func (c *Client) DescribeVirusMonitorSetting(request *DescribeVirusMonitorSettingRequest) (response *DescribeVirusMonitorSettingResponse, err error) {
	if request == nil {
		request = NewDescribeVirusMonitorSettingRequest()
	}
	response = NewDescribeVirusMonitorSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageRiskTendencyRequest() (request *DescribeImageRiskTendencyRequest) {
	request = &DescribeImageRiskTendencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeImageRiskTendency")
	return
}

func NewDescribeImageRiskTendencyResponse() (response *DescribeImageRiskTendencyResponse) {
	response = &DescribeImageRiskTendencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器安全本地镜像风险趋势
func (c *Client) DescribeImageRiskTendency(request *DescribeImageRiskTendencyRequest) (response *DescribeImageRiskTendencyResponse, err error) {
	if request == nil {
		request = NewDescribeImageRiskTendencyRequest()
	}
	response = NewDescribeImageRiskTendencyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryVirusListRequest() (request *DescribeAssetImageRegistryVirusListRequest) {
	request = &DescribeAssetImageRegistryVirusListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageRegistryVirusList")
	return
}

func NewDescribeAssetImageRegistryVirusListResponse() (response *DescribeAssetImageRegistryVirusListResponse) {
	response = &DescribeAssetImageRegistryVirusListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库查询木马病毒列表
func (c *Client) DescribeAssetImageRegistryVirusList(request *DescribeAssetImageRegistryVirusListRequest) (response *DescribeAssetImageRegistryVirusListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryVirusListRequest()
	}
	response = NewDescribeAssetImageRegistryVirusListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRegistryDetailRequest() (request *DescribeAssetImageRegistryDetailRequest) {
	request = &DescribeAssetImageRegistryDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageRegistryDetail")
	return
}

func NewDescribeAssetImageRegistryDetailResponse() (response *DescribeAssetImageRegistryDetailResponse) {
	response = &DescribeAssetImageRegistryDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库查询镜像仓库详情
func (c *Client) DescribeAssetImageRegistryDetail(request *DescribeAssetImageRegistryDetailRequest) (response *DescribeAssetImageRegistryDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRegistryDetailRequest()
	}
	response = NewDescribeAssetImageRegistryDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecLogCleanSettingInfoRequest() (request *DescribeSecLogCleanSettingInfoRequest) {
	request = &DescribeSecLogCleanSettingInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeSecLogCleanSettingInfo")
	return
}

func NewDescribeSecLogCleanSettingInfoResponse() (response *DescribeSecLogCleanSettingInfoResponse) {
	response = &DescribeSecLogCleanSettingInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全日志清理设置详情
func (c *Client) DescribeSecLogCleanSettingInfo(request *DescribeSecLogCleanSettingInfoRequest) (response *DescribeSecLogCleanSettingInfoResponse, err error) {
	if request == nil {
		request = NewDescribeSecLogCleanSettingInfoRequest()
	}
	response = NewDescribeSecLogCleanSettingInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageDetailRequest() (request *DescribeAssetImageDetailRequest) {
	request = &DescribeAssetImageDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageDetail")
	return
}

func NewDescribeAssetImageDetailResponse() (response *DescribeAssetImageDetailResponse) {
	response = &DescribeAssetImageDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像详细信息
func (c *Client) DescribeAssetImageDetail(request *DescribeAssetImageDetailRequest) (response *DescribeAssetImageDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageDetailRequest()
	}
	response = NewDescribeAssetImageDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRiskSyscallWhiteListsExportJobRequest() (request *CreateRiskSyscallWhiteListsExportJobRequest) {
	request = &CreateRiskSyscallWhiteListsExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateRiskSyscallWhiteListsExportJob")
	return
}

func NewCreateRiskSyscallWhiteListsExportJobResponse() (response *CreateRiskSyscallWhiteListsExportJobResponse) {
	response = &CreateRiskSyscallWhiteListsExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建运行时高危系统调用白名单列表导出任务
func (c *Client) CreateRiskSyscallWhiteListsExportJob(request *CreateRiskSyscallWhiteListsExportJobRequest) (response *CreateRiskSyscallWhiteListsExportJobResponse, err error) {
	if request == nil {
		request = NewCreateRiskSyscallWhiteListsExportJobRequest()
	}
	response = NewCreateRiskSyscallWhiteListsExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetContainerListRequest() (request *DescribeAssetContainerListRequest) {
	request = &DescribeAssetContainerListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetContainerList")
	return
}

func NewDescribeAssetContainerListResponse() (response *DescribeAssetContainerListResponse) {
	response = &DescribeAssetContainerListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器列表
func (c *Client) DescribeAssetContainerList(request *DescribeAssetContainerListRequest) (response *DescribeAssetContainerListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetContainerListRequest()
	}
	response = NewDescribeAssetContainerListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageListRequest() (request *DescribeAssetImageListRequest) {
	request = &DescribeAssetImageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageList")
	return
}

func NewDescribeAssetImageListResponse() (response *DescribeAssetImageListResponse) {
	response = &DescribeAssetImageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像列表
func (c *Client) DescribeAssetImageList(request *DescribeAssetImageListRequest) (response *DescribeAssetImageListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageListRequest()
	}
	response = NewDescribeAssetImageListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateImageVulListExportJobRequest() (request *CreateImageVulListExportJobRequest) {
	request = &CreateImageVulListExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateImageVulListExportJob")
	return
}

func NewCreateImageVulListExportJobResponse() (response *CreateImageVulListExportJobResponse) {
	response = &CreateImageVulListExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建镜像漏洞列表导出任务
func (c *Client) CreateImageVulListExportJob(request *CreateImageVulListExportJobRequest) (response *CreateImageVulListExportJobResponse, err error) {
	if request == nil {
		request = NewCreateImageVulListExportJobRequest()
	}
	response = NewCreateImageVulListExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageDenyRuleSummaryRequest() (request *DescribeImageDenyRuleSummaryRequest) {
	request = &DescribeImageDenyRuleSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeImageDenyRuleSummary")
	return
}

func NewDescribeImageDenyRuleSummaryResponse() (response *DescribeImageDenyRuleSummaryResponse) {
	response = &DescribeImageDenyRuleSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像拦截规则统计
func (c *Client) DescribeImageDenyRuleSummary(request *DescribeImageDenyRuleSummaryRequest) (response *DescribeImageDenyRuleSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeImageDenyRuleSummaryRequest()
	}
	response = NewDescribeImageDenyRuleSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUserClusterExportJobRequest() (request *CreateUserClusterExportJobRequest) {
	request = &CreateUserClusterExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateUserClusterExportJob")
	return
}

func NewCreateUserClusterExportJobResponse() (response *CreateUserClusterExportJobResponse) {
	response = &CreateUserClusterExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集群信息导出任务，可以指定一个或者多个导出
func (c *Client) CreateUserClusterExportJob(request *CreateUserClusterExportJobRequest) (response *CreateUserClusterExportJobResponse, err error) {
	if request == nil {
		request = NewCreateUserClusterExportJobRequest()
	}
	response = NewCreateUserClusterExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCompliancePolicyItemAffectedSummaryRequest() (request *DescribeCompliancePolicyItemAffectedSummaryRequest) {
	request = &DescribeCompliancePolicyItemAffectedSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeCompliancePolicyItemAffectedSummary")
	return
}

func NewDescribeCompliancePolicyItemAffectedSummaryResponse() (response *DescribeCompliancePolicyItemAffectedSummaryResponse) {
	response = &DescribeCompliancePolicyItemAffectedSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按照 检测项 → 资产 的两级层次展开的第一层级：检测项层级。
func (c *Client) DescribeCompliancePolicyItemAffectedSummary(request *DescribeCompliancePolicyItemAffectedSummaryRequest) (response *DescribeCompliancePolicyItemAffectedSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeCompliancePolicyItemAffectedSummaryRequest()
	}
	response = NewDescribeCompliancePolicyItemAffectedSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageHostListRequest() (request *DescribeAssetImageHostListRequest) {
	request = &DescribeAssetImageHostListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageHostList")
	return
}

func NewDescribeAssetImageHostListResponse() (response *DescribeAssetImageHostListResponse) {
	response = &DescribeAssetImageHostListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像关联主机
func (c *Client) DescribeAssetImageHostList(request *DescribeAssetImageHostListRequest) (response *DescribeAssetImageHostListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageHostListRequest()
	}
	response = NewDescribeAssetImageHostListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateComponentExportJobRequest() (request *CreateComponentExportJobRequest) {
	request = &CreateComponentExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateComponentExportJob")
	return
}

func NewCreateComponentExportJobResponse() (response *CreateComponentExportJobResponse) {
	response = &CreateComponentExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询本地镜像组件列表导出
func (c *Client) CreateComponentExportJob(request *CreateComponentExportJobRequest) (response *CreateComponentExportJobResponse, err error) {
	if request == nil {
		request = NewCreateComponentExportJobRequest()
	}
	response = NewCreateComponentExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskDnsListRequest() (request *DescribeRiskDnsListRequest) {
	request = &DescribeRiskDnsListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeRiskDnsList")
	return
}

func NewDescribeRiskDnsListResponse() (response *DescribeRiskDnsListResponse) {
	response = &DescribeRiskDnsListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询恶意请求事件列表
func (c *Client) DescribeRiskDnsList(request *DescribeRiskDnsListRequest) (response *DescribeRiskDnsListResponse, err error) {
	if request == nil {
		request = NewDescribeRiskDnsListRequest()
	}
	response = NewDescribeRiskDnsListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeK8sApiAbnormalSummaryRequest() (request *DescribeK8sApiAbnormalSummaryRequest) {
	request = &DescribeK8sApiAbnormalSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeK8sApiAbnormalSummary")
	return
}

func NewDescribeK8sApiAbnormalSummaryResponse() (response *DescribeK8sApiAbnormalSummaryResponse) {
	response = &DescribeK8sApiAbnormalSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询k8sapi异常事件统计
func (c *Client) DescribeK8sApiAbnormalSummary(request *DescribeK8sApiAbnormalSummaryRequest) (response *DescribeK8sApiAbnormalSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeK8sApiAbnormalSummaryRequest()
	}
	response = NewDescribeK8sApiAbnormalSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbnormalProcessLevelSummaryRequest() (request *DescribeAbnormalProcessLevelSummaryRequest) {
	request = &DescribeAbnormalProcessLevelSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAbnormalProcessLevelSummary")
	return
}

func NewDescribeAbnormalProcessLevelSummaryResponse() (response *DescribeAbnormalProcessLevelSummaryResponse) {
	response = &DescribeAbnormalProcessLevelSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 统计异常进程各威胁等级待处理事件数
func (c *Client) DescribeAbnormalProcessLevelSummary(request *DescribeAbnormalProcessLevelSummaryRequest) (response *DescribeAbnormalProcessLevelSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeAbnormalProcessLevelSummaryRequest()
	}
	response = NewDescribeAbnormalProcessLevelSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAbnormalProcessRulesExportJobRequest() (request *CreateAbnormalProcessRulesExportJobRequest) {
	request = &CreateAbnormalProcessRulesExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateAbnormalProcessRulesExportJob")
	return
}

func NewCreateAbnormalProcessRulesExportJobResponse() (response *CreateAbnormalProcessRulesExportJobResponse) {
	response = &CreateAbnormalProcessRulesExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建异常进程规则导出任务
func (c *Client) CreateAbnormalProcessRulesExportJob(request *CreateAbnormalProcessRulesExportJobRequest) (response *CreateAbnormalProcessRulesExportJobResponse, err error) {
	if request == nil {
		request = NewCreateAbnormalProcessRulesExportJobRequest()
	}
	response = NewCreateAbnormalProcessRulesExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAssetImageRiskExportJobRequest() (request *CreateAssetImageRiskExportJobRequest) {
	request = &CreateAssetImageRiskExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateAssetImageRiskExportJob")
	return
}

func NewCreateAssetImageRiskExportJobResponse() (response *CreateAssetImageRiskExportJobResponse) {
	response = &CreateAssetImageRiskExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建本地镜像敏感信息列表导出任务
func (c *Client) CreateAssetImageRiskExportJob(request *CreateAssetImageRiskExportJobRequest) (response *CreateAssetImageRiskExportJobResponse, err error) {
	if request == nil {
		request = NewCreateAssetImageRiskExportJobRequest()
	}
	response = NewCreateAssetImageRiskExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerAssetSummaryRequest() (request *DescribeContainerAssetSummaryRequest) {
	request = &DescribeContainerAssetSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeContainerAssetSummary")
	return
}

func NewDescribeContainerAssetSummaryResponse() (response *DescribeContainerAssetSummaryResponse) {
	response = &DescribeContainerAssetSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器安全资产概览
func (c *Client) DescribeContainerAssetSummary(request *DescribeContainerAssetSummaryRequest) (response *DescribeContainerAssetSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeContainerAssetSummaryRequest()
	}
	response = NewDescribeContainerAssetSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageRiskListRequest() (request *DescribeAssetImageRiskListRequest) {
	request = &DescribeAssetImageRiskListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageRiskList")
	return
}

func NewDescribeAssetImageRiskListResponse() (response *DescribeAssetImageRiskListResponse) {
	response = &DescribeAssetImageRiskListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像风险列表
func (c *Client) DescribeAssetImageRiskList(request *DescribeAssetImageRiskListRequest) (response *DescribeAssetImageRiskListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageRiskListRequest()
	}
	response = NewDescribeAssetImageRiskListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulRegistryImageListRequest() (request *DescribeVulRegistryImageListRequest) {
	request = &DescribeVulRegistryImageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVulRegistryImageList")
	return
}

func NewDescribeVulRegistryImageListResponse() (response *DescribeVulRegistryImageListResponse) {
	response = &DescribeVulRegistryImageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞影响的仓库镜像列表
func (c *Client) DescribeVulRegistryImageList(request *DescribeVulRegistryImageListRequest) (response *DescribeVulRegistryImageListResponse, err error) {
	if request == nil {
		request = NewDescribeVulRegistryImageListRequest()
	}
	response = NewDescribeVulRegistryImageListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusScanTimeoutSettingRequest() (request *DescribeVirusScanTimeoutSettingRequest) {
	request = &DescribeVirusScanTimeoutSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVirusScanTimeoutSetting")
	return
}

func NewDescribeVirusScanTimeoutSettingResponse() (response *DescribeVirusScanTimeoutSettingResponse) {
	response = &DescribeVirusScanTimeoutSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时文件扫描超时设置查询
func (c *Client) DescribeVirusScanTimeoutSetting(request *DescribeVirusScanTimeoutSettingRequest) (response *DescribeVirusScanTimeoutSettingResponse, err error) {
	if request == nil {
		request = NewDescribeVirusScanTimeoutSettingRequest()
	}
	response = NewDescribeVirusScanTimeoutSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbnormalProcessEventsExportRequest() (request *DescribeAbnormalProcessEventsExportRequest) {
	request = &DescribeAbnormalProcessEventsExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAbnormalProcessEventsExport")
	return
}

func NewDescribeAbnormalProcessEventsExportResponse() (response *DescribeAbnormalProcessEventsExportResponse) {
	response = &DescribeAbnormalProcessEventsExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时异常进程列表导出
func (c *Client) DescribeAbnormalProcessEventsExport(request *DescribeAbnormalProcessEventsExportRequest) (response *DescribeAbnormalProcessEventsExportResponse, err error) {
	if request == nil {
		request = NewDescribeAbnormalProcessEventsExportRequest()
	}
	response = NewDescribeAbnormalProcessEventsExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterAssetSummaryRequest() (request *DescribeClusterAssetSummaryRequest) {
	request = &DescribeClusterAssetSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeClusterAssetSummary")
	return
}

func NewDescribeClusterAssetSummaryResponse() (response *DescribeClusterAssetSummaryResponse) {
	response = &DescribeClusterAssetSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户集群资产的总览信息
func (c *Client) DescribeClusterAssetSummary(request *DescribeClusterAssetSummaryRequest) (response *DescribeClusterAssetSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeClusterAssetSummaryRequest()
	}
	response = NewDescribeClusterAssetSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellEventsRequest() (request *DescribeReverseShellEventsRequest) {
	request = &DescribeReverseShellEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeReverseShellEvents")
	return
}

func NewDescribeReverseShellEventsResponse() (response *DescribeReverseShellEventsResponse) {
	response = &DescribeReverseShellEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时反弹shell列表
func (c *Client) DescribeReverseShellEvents(request *DescribeReverseShellEventsRequest) (response *DescribeReverseShellEventsResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellEventsRequest()
	}
	response = NewDescribeReverseShellEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageSimpleListRequest() (request *DescribeAssetImageSimpleListRequest) {
	request = &DescribeAssetImageSimpleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageSimpleList")
	return
}

func NewDescribeAssetImageSimpleListResponse() (response *DescribeAssetImageSimpleListResponse) {
	response = &DescribeAssetImageSimpleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像简略信息列表
func (c *Client) DescribeAssetImageSimpleList(request *DescribeAssetImageSimpleListRequest) (response *DescribeAssetImageSimpleListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageSimpleListRequest()
	}
	response = NewDescribeAssetImageSimpleListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostSecurityServiceStateRequest() (request *ModifyHostSecurityServiceStateRequest) {
	request = &ModifyHostSecurityServiceStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "ModifyHostSecurityServiceState")
	return
}

func NewModifyHostSecurityServiceStateResponse() (response *ModifyHostSecurityServiceStateResponse) {
	response = &ModifyHostSecurityServiceStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑主机容器安全服务开启或停止
func (c *Client) ModifyHostSecurityServiceState(request *ModifyHostSecurityServiceStateRequest) (response *ModifyHostSecurityServiceStateResponse, err error) {
	if request == nil {
		request = NewModifyHostSecurityServiceStateRequest()
	}
	response = NewModifyHostSecurityServiceStateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyFlexBillingStateRequest() (request *ModifyFlexBillingStateRequest) {
	request = &ModifyFlexBillingStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "ModifyFlexBillingState")
	return
}

func NewModifyFlexBillingStateResponse() (response *ModifyFlexBillingStateResponse) {
	response = &ModifyFlexBillingStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改弹性计费状态
func (c *Client) ModifyFlexBillingState(request *ModifyFlexBillingStateRequest) (response *ModifyFlexBillingStateResponse, err error) {
	if request == nil {
		request = NewModifyFlexBillingStateRequest()
	}
	response = NewModifyFlexBillingStateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulImageListRequest() (request *DescribeVulImageListRequest) {
	request = &DescribeVulImageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVulImageList")
	return
}

func NewDescribeVulImageListResponse() (response *DescribeVulImageListResponse) {
	response = &DescribeVulImageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞影响的镜像列表
func (c *Client) DescribeVulImageList(request *DescribeVulImageListRequest) (response *DescribeVulImageListResponse, err error) {
	if request == nil {
		request = NewDescribeVulImageListRequest()
	}
	response = NewDescribeVulImageListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeK8sApiAbnormalEventListRequest() (request *DescribeK8sApiAbnormalEventListRequest) {
	request = &DescribeK8sApiAbnormalEventListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeK8sApiAbnormalEventList")
	return
}

func NewDescribeK8sApiAbnormalEventListResponse() (response *DescribeK8sApiAbnormalEventListResponse) {
	response = &DescribeK8sApiAbnormalEventListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询k8s api异常事件列表
func (c *Client) DescribeK8sApiAbnormalEventList(request *DescribeK8sApiAbnormalEventListRequest) (response *DescribeK8sApiAbnormalEventListResponse, err error) {
	if request == nil {
		request = NewDescribeK8sApiAbnormalEventListRequest()
	}
	response = NewDescribeK8sApiAbnormalEventListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusAutoIsolateSampleListRequest() (request *DescribeVirusAutoIsolateSampleListRequest) {
	request = &DescribeVirusAutoIsolateSampleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVirusAutoIsolateSampleList")
	return
}

func NewDescribeVirusAutoIsolateSampleListResponse() (response *DescribeVirusAutoIsolateSampleListResponse) {
	response = &DescribeVirusAutoIsolateSampleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询木马自动隔离样本列表
func (c *Client) DescribeVirusAutoIsolateSampleList(request *DescribeVirusAutoIsolateSampleListRequest) (response *DescribeVirusAutoIsolateSampleListResponse, err error) {
	if request == nil {
		request = NewDescribeVirusAutoIsolateSampleListRequest()
	}
	response = NewDescribeVirusAutoIsolateSampleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageDenyEventTendencyRequest() (request *DescribeImageDenyEventTendencyRequest) {
	request = &DescribeImageDenyEventTendencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeImageDenyEventTendency")
	return
}

func NewDescribeImageDenyEventTendencyResponse() (response *DescribeImageDenyEventTendencyResponse) {
	response = &DescribeImageDenyEventTendencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像拦截事件趋势
func (c *Client) DescribeImageDenyEventTendency(request *DescribeImageDenyEventTendencyRequest) (response *DescribeImageDenyEventTendencyResponse, err error) {
	if request == nil {
		request = NewDescribeImageDenyEventTendencyRequest()
	}
	response = NewDescribeImageDenyEventTendencyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecEventsTendencyRequest() (request *DescribeSecEventsTendencyRequest) {
	request = &DescribeSecEventsTendencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeSecEventsTendency")
	return
}

func NewDescribeSecEventsTendencyResponse() (response *DescribeSecEventsTendencyResponse) {
	response = &DescribeSecEventsTendencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器运行时安全时间趋势
func (c *Client) DescribeSecEventsTendency(request *DescribeSecEventsTendencyRequest) (response *DescribeSecEventsTendencyResponse, err error) {
	if request == nil {
		request = NewDescribeSecEventsTendencyRequest()
	}
	response = NewDescribeSecEventsTendencyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetImageVulListExportRequest() (request *DescribeAssetImageVulListExportRequest) {
	request = &DescribeAssetImageVulListExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetImageVulListExport")
	return
}

func NewDescribeAssetImageVulListExportResponse() (response *DescribeAssetImageVulListExportResponse) {
	response = &DescribeAssetImageVulListExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像漏洞列表导出
func (c *Client) DescribeAssetImageVulListExport(request *DescribeAssetImageVulListExportRequest) (response *DescribeAssetImageVulListExportResponse, err error) {
	if request == nil {
		request = NewDescribeAssetImageVulListExportRequest()
	}
	response = NewDescribeAssetImageVulListExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeK8sApiAbnormalRuleListRequest() (request *DescribeK8sApiAbnormalRuleListRequest) {
	request = &DescribeK8sApiAbnormalRuleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeK8sApiAbnormalRuleList")
	return
}

func NewDescribeK8sApiAbnormalRuleListResponse() (response *DescribeK8sApiAbnormalRuleListResponse) {
	response = &DescribeK8sApiAbnormalRuleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询k8sapi异常请求规则列表
func (c *Client) DescribeK8sApiAbnormalRuleList(request *DescribeK8sApiAbnormalRuleListRequest) (response *DescribeK8sApiAbnormalRuleListResponse, err error) {
	if request == nil {
		request = NewDescribeK8sApiAbnormalRuleListRequest()
	}
	response = NewDescribeK8sApiAbnormalRuleListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComplianceScanFailedAssetListRequest() (request *DescribeComplianceScanFailedAssetListRequest) {
	request = &DescribeComplianceScanFailedAssetListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeComplianceScanFailedAssetList")
	return
}

func NewDescribeComplianceScanFailedAssetListResponse() (response *DescribeComplianceScanFailedAssetListResponse) {
	response = &DescribeComplianceScanFailedAssetListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按照 资产 → 检测项 二层结构展示的信息。这里查询第一层 资产的通过率汇总信息。
func (c *Client) DescribeComplianceScanFailedAssetList(request *DescribeComplianceScanFailedAssetListRequest) (response *DescribeComplianceScanFailedAssetListResponse, err error) {
	if request == nil {
		request = NewDescribeComplianceScanFailedAssetListRequest()
	}
	response = NewDescribeComplianceScanFailedAssetListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComplianceTaskPolicyItemSummaryListRequest() (request *DescribeComplianceTaskPolicyItemSummaryListRequest) {
	request = &DescribeComplianceTaskPolicyItemSummaryListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeComplianceTaskPolicyItemSummaryList")
	return
}

func NewDescribeComplianceTaskPolicyItemSummaryListResponse() (response *DescribeComplianceTaskPolicyItemSummaryListResponse) {
	response = &DescribeComplianceTaskPolicyItemSummaryListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询最近一次任务发现的检测项的汇总信息列表，按照 检测项 → 资产 的两级层次展开。
func (c *Client) DescribeComplianceTaskPolicyItemSummaryList(request *DescribeComplianceTaskPolicyItemSummaryListRequest) (response *DescribeComplianceTaskPolicyItemSummaryListResponse, err error) {
	if request == nil {
		request = NewDescribeComplianceTaskPolicyItemSummaryListRequest()
	}
	response = NewDescribeComplianceTaskPolicyItemSummaryListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterSummaryRequest() (request *DescribeClusterSummaryRequest) {
	request = &DescribeClusterSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeClusterSummary")
	return
}

func NewDescribeClusterSummaryResponse() (response *DescribeClusterSummaryResponse) {
	response = &DescribeClusterSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户集群资产总览
func (c *Client) DescribeClusterSummary(request *DescribeClusterSummaryRequest) (response *DescribeClusterSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeClusterSummaryRequest()
	}
	response = NewDescribeClusterSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetClusterListRequest() (request *DescribeAssetClusterListRequest) {
	request = &DescribeAssetClusterListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetClusterList")
	return
}

func NewDescribeAssetClusterListResponse() (response *DescribeAssetClusterListResponse) {
	response = &DescribeAssetClusterListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群列表
func (c *Client) DescribeAssetClusterList(request *DescribeAssetClusterListRequest) (response *DescribeAssetClusterListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetClusterListRequest()
	}
	response = NewDescribeAssetClusterListResponse()
	err = c.Send(request, response)
	return
}

func NewOpenTcssTrialRequest() (request *OpenTcssTrialRequest) {
	request = &OpenTcssTrialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "OpenTcssTrial")
	return
}

func NewOpenTcssTrialResponse() (response *OpenTcssTrialResponse) {
	response = &OpenTcssTrialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 申请使用容器安全服务
func (c *Client) OpenTcssTrial(request *OpenTcssTrialRequest) (response *OpenTcssTrialResponse, err error) {
	if request == nil {
		request = NewOpenTcssTrialRequest()
	}
	response = NewOpenTcssTrialResponse()
	err = c.Send(request, response)
	return
}

func NewCreateHostExportJobRequest() (request *CreateHostExportJobRequest) {
	request = &CreateHostExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateHostExportJob")
	return
}

func NewCreateHostExportJobResponse() (response *CreateHostExportJobResponse) {
	response = &CreateHostExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建主机列表导出任务
func (c *Client) CreateHostExportJob(request *CreateHostExportJobRequest) (response *CreateHostExportJobResponse, err error) {
	if request == nil {
		request = NewCreateHostExportJobRequest()
	}
	response = NewCreateHostExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEscapeEventTendencyRequest() (request *DescribeEscapeEventTendencyRequest) {
	request = &DescribeEscapeEventTendencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeEscapeEventTendency")
	return
}

func NewDescribeEscapeEventTendencyResponse() (response *DescribeEscapeEventTendencyResponse) {
	response = &DescribeEscapeEventTendencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询待处理逃逸事件趋势
func (c *Client) DescribeEscapeEventTendency(request *DescribeEscapeEventTendencyRequest) (response *DescribeEscapeEventTendencyResponse, err error) {
	if request == nil {
		request = NewDescribeEscapeEventTendencyRequest()
	}
	response = NewDescribeEscapeEventTendencyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateWebVulExportJobRequest() (request *CreateWebVulExportJobRequest) {
	request = &CreateWebVulExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateWebVulExportJob")
	return
}

func NewCreateWebVulExportJobResponse() (response *CreateWebVulExportJobResponse) {
	response = &CreateWebVulExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建web漏洞导出任务
func (c *Client) CreateWebVulExportJob(request *CreateWebVulExportJobRequest) (response *CreateWebVulExportJobResponse, err error) {
	if request == nil {
		request = NewCreateWebVulExportJobRequest()
	}
	response = NewCreateWebVulExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbnormalProcessRuleDetailRequest() (request *DescribeAbnormalProcessRuleDetailRequest) {
	request = &DescribeAbnormalProcessRuleDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAbnormalProcessRuleDetail")
	return
}

func NewDescribeAbnormalProcessRuleDetailResponse() (response *DescribeAbnormalProcessRuleDetailResponse) {
	response = &DescribeAbnormalProcessRuleDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行时异常进程策略详细信息
func (c *Client) DescribeAbnormalProcessRuleDetail(request *DescribeAbnormalProcessRuleDetailRequest) (response *DescribeAbnormalProcessRuleDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAbnormalProcessRuleDetailRequest()
	}
	response = NewDescribeAbnormalProcessRuleDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateImageDenyRuleExportJobRequest() (request *CreateImageDenyRuleExportJobRequest) {
	request = &CreateImageDenyRuleExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateImageDenyRuleExportJob")
	return
}

func NewCreateImageDenyRuleExportJobResponse() (response *CreateImageDenyRuleExportJobResponse) {
	response = &CreateImageDenyRuleExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建镜像拦截规则导出任务
func (c *Client) CreateImageDenyRuleExportJob(request *CreateImageDenyRuleExportJobRequest) (response *CreateImageDenyRuleExportJobResponse, err error) {
	if request == nil {
		request = NewCreateImageDenyRuleExportJobRequest()
	}
	response = NewCreateImageDenyRuleExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellDetailRequest() (request *DescribeReverseShellDetailRequest) {
	request = &DescribeReverseShellDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeReverseShellDetail")
	return
}

func NewDescribeReverseShellDetailResponse() (response *DescribeReverseShellDetailResponse) {
	response = &DescribeReverseShellDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时反弹shell事件详细信息
func (c *Client) DescribeReverseShellDetail(request *DescribeReverseShellDetailRequest) (response *DescribeReverseShellDetailResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellDetailRequest()
	}
	response = NewDescribeReverseShellDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateK8sApiAbnormalRuleExportJobRequest() (request *CreateK8sApiAbnormalRuleExportJobRequest) {
	request = &CreateK8sApiAbnormalRuleExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateK8sApiAbnormalRuleExportJob")
	return
}

func NewCreateK8sApiAbnormalRuleExportJobResponse() (response *CreateK8sApiAbnormalRuleExportJobResponse) {
	response = &CreateK8sApiAbnormalRuleExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建k8sApi异常规则导出任务
func (c *Client) CreateK8sApiAbnormalRuleExportJob(request *CreateK8sApiAbnormalRuleExportJobRequest) (response *CreateK8sApiAbnormalRuleExportJobResponse, err error) {
	if request == nil {
		request = NewCreateK8sApiAbnormalRuleExportJobRequest()
	}
	response = NewCreateK8sApiAbnormalRuleExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProcessEventsExportJobRequest() (request *CreateProcessEventsExportJobRequest) {
	request = &CreateProcessEventsExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateProcessEventsExportJob")
	return
}

func NewCreateProcessEventsExportJobResponse() (response *CreateProcessEventsExportJobResponse) {
	response = &CreateProcessEventsExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建异常进程事件导出异步任务
func (c *Client) CreateProcessEventsExportJob(request *CreateProcessEventsExportJobRequest) (response *CreateProcessEventsExportJobResponse, err error) {
	if request == nil {
		request = NewCreateProcessEventsExportJobRequest()
	}
	response = NewCreateProcessEventsExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAccessControlsRuleExportJobRequest() (request *CreateAccessControlsRuleExportJobRequest) {
	request = &CreateAccessControlsRuleExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateAccessControlsRuleExportJob")
	return
}

func NewCreateAccessControlsRuleExportJobResponse() (response *CreateAccessControlsRuleExportJobResponse) {
	response = &CreateAccessControlsRuleExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建文件篡改规则导出任务
func (c *Client) CreateAccessControlsRuleExportJob(request *CreateAccessControlsRuleExportJobRequest) (response *CreateAccessControlsRuleExportJobResponse, err error) {
	if request == nil {
		request = NewCreateAccessControlsRuleExportJobRequest()
	}
	response = NewCreateAccessControlsRuleExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEventEscapeImageExportJobRequest() (request *CreateEventEscapeImageExportJobRequest) {
	request = &CreateEventEscapeImageExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateEventEscapeImageExportJob")
	return
}

func NewCreateEventEscapeImageExportJobResponse() (response *CreateEventEscapeImageExportJobResponse) {
	response = &CreateEventEscapeImageExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建风险容器镜像导出任务
func (c *Client) CreateEventEscapeImageExportJob(request *CreateEventEscapeImageExportJobRequest) (response *CreateEventEscapeImageExportJobResponse, err error) {
	if request == nil {
		request = NewCreateEventEscapeImageExportJobRequest()
	}
	response = NewCreateEventEscapeImageExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbnormalProcessRulesRequest() (request *DescribeAbnormalProcessRulesRequest) {
	request = &DescribeAbnormalProcessRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAbnormalProcessRules")
	return
}

func NewDescribeAbnormalProcessRulesResponse() (response *DescribeAbnormalProcessRulesResponse) {
	response = &DescribeAbnormalProcessRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时异常进程策略列表
func (c *Client) DescribeAbnormalProcessRules(request *DescribeAbnormalProcessRulesRequest) (response *DescribeAbnormalProcessRulesResponse, err error) {
	if request == nil {
		request = NewDescribeAbnormalProcessRulesRequest()
	}
	response = NewDescribeAbnormalProcessRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserClusterRequest() (request *DescribeUserClusterRequest) {
	request = &DescribeUserClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeUserCluster")
	return
}

func NewDescribeUserClusterResponse() (response *DescribeUserClusterResponse) {
	response = &DescribeUserClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 安全概览和集群安全页进入调用该接口，查询用户集群相关信息。
func (c *Client) DescribeUserCluster(request *DescribeUserClusterRequest) (response *DescribeUserClusterResponse, err error) {
	if request == nil {
		request = NewDescribeUserClusterRequest()
	}
	response = NewDescribeUserClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessControlRulesRequest() (request *DescribeAccessControlRulesRequest) {
	request = &DescribeAccessControlRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAccessControlRules")
	return
}

func NewDescribeAccessControlRulesResponse() (response *DescribeAccessControlRulesResponse) {
	response = &DescribeAccessControlRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时访问控制策略列表
func (c *Client) DescribeAccessControlRules(request *DescribeAccessControlRulesRequest) (response *DescribeAccessControlRulesResponse, err error) {
	if request == nil {
		request = NewDescribeAccessControlRulesRequest()
	}
	response = NewDescribeAccessControlRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterNameListRequest() (request *DescribeClusterNameListRequest) {
	request = &DescribeClusterNameListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeClusterNameList")
	return
}

func NewDescribeClusterNameListResponse() (response *DescribeClusterNameListResponse) {
	response = &DescribeClusterNameListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户所有集群名字
func (c *Client) DescribeClusterNameList(request *DescribeClusterNameListRequest) (response *DescribeClusterNameListResponse, err error) {
	if request == nil {
		request = NewDescribeClusterNameListRequest()
	}
	response = NewDescribeClusterNameListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeK8sApiAbnormalRuleInfoRequest() (request *DescribeK8sApiAbnormalRuleInfoRequest) {
	request = &DescribeK8sApiAbnormalRuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeK8sApiAbnormalRuleInfo")
	return
}

func NewDescribeK8sApiAbnormalRuleInfoResponse() (response *DescribeK8sApiAbnormalRuleInfoResponse) {
	response = &DescribeK8sApiAbnormalRuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询k8sapi异常请求规则详情
func (c *Client) DescribeK8sApiAbnormalRuleInfo(request *DescribeK8sApiAbnormalRuleInfoRequest) (response *DescribeK8sApiAbnormalRuleInfoResponse, err error) {
	if request == nil {
		request = NewDescribeK8sApiAbnormalRuleInfoRequest()
	}
	response = NewDescribeK8sApiAbnormalRuleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterNodeWhiteListRequest() (request *DescribeClusterNodeWhiteListRequest) {
	request = &DescribeClusterNodeWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeClusterNodeWhiteList")
	return
}

func NewDescribeClusterNodeWhiteListResponse() (response *DescribeClusterNodeWhiteListResponse) {
	response = &DescribeClusterNodeWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取忽略的节点信息白名单
func (c *Client) DescribeClusterNodeWhiteList(request *DescribeClusterNodeWhiteListRequest) (response *DescribeClusterNodeWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeClusterNodeWhiteListRequest()
	}
	response = NewDescribeClusterNodeWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReverseShellWhiteListsExportRequest() (request *DescribeReverseShellWhiteListsExportRequest) {
	request = &DescribeReverseShellWhiteListsExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeReverseShellWhiteListsExport")
	return
}

func NewDescribeReverseShellWhiteListsExportResponse() (response *DescribeReverseShellWhiteListsExportResponse) {
	response = &DescribeReverseShellWhiteListsExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 反弹shell白名单列表导出
func (c *Client) DescribeReverseShellWhiteListsExport(request *DescribeReverseShellWhiteListsExportRequest) (response *DescribeReverseShellWhiteListsExportResponse, err error) {
	if request == nil {
		request = NewDescribeReverseShellWhiteListsExportRequest()
	}
	response = NewDescribeReverseShellWhiteListsExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterWorkloadsRequest() (request *DescribeClusterWorkloadsRequest) {
	request = &DescribeClusterWorkloadsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeClusterWorkloads")
	return
}

func NewDescribeClusterWorkloadsResponse() (response *DescribeClusterWorkloadsResponse) {
	response = &DescribeClusterWorkloadsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群工作负载信息
func (c *Client) DescribeClusterWorkloads(request *DescribeClusterWorkloadsRequest) (response *DescribeClusterWorkloadsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterWorkloadsRequest()
	}
	response = NewDescribeClusterWorkloadsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeK8sApiAbnormalTendencyRequest() (request *DescribeK8sApiAbnormalTendencyRequest) {
	request = &DescribeK8sApiAbnormalTendencyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeK8sApiAbnormalTendency")
	return
}

func NewDescribeK8sApiAbnormalTendencyResponse() (response *DescribeK8sApiAbnormalTendencyResponse) {
	response = &DescribeK8sApiAbnormalTendencyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询k8sapi异常事件趋势
func (c *Client) DescribeK8sApiAbnormalTendency(request *DescribeK8sApiAbnormalTendencyRequest) (response *DescribeK8sApiAbnormalTendencyResponse, err error) {
	if request == nil {
		request = NewDescribeK8sApiAbnormalTendencyRequest()
	}
	response = NewDescribeK8sApiAbnormalTendencyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAffectedWorkloadListRequest() (request *DescribeAffectedWorkloadListRequest) {
	request = &DescribeAffectedWorkloadListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAffectedWorkloadList")
	return
}

func NewDescribeAffectedWorkloadListResponse() (response *DescribeAffectedWorkloadListResponse) {
	response = &DescribeAffectedWorkloadListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询workload类型的影响范围，返回workload列表
func (c *Client) DescribeAffectedWorkloadList(request *DescribeAffectedWorkloadListRequest) (response *DescribeAffectedWorkloadListResponse, err error) {
	if request == nil {
		request = NewDescribeAffectedWorkloadListRequest()
	}
	response = NewDescribeAffectedWorkloadListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetHostListRequest() (request *DescribeAssetHostListRequest) {
	request = &DescribeAssetHostListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetHostList")
	return
}

func NewDescribeAssetHostListResponse() (response *DescribeAssetHostListResponse) {
	response = &DescribeAssetHostListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主机列表
func (c *Client) DescribeAssetHostList(request *DescribeAssetHostListRequest) (response *DescribeAssetHostListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetHostListRequest()
	}
	response = NewDescribeAssetHostListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostStartTheServiceRequest() (request *ModifyHostStartTheServiceRequest) {
	request = &ModifyHostStartTheServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "ModifyHostStartTheService")
	return
}

func NewModifyHostStartTheServiceResponse() (response *ModifyHostStartTheServiceResponse) {
	response = &ModifyHostStartTheServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 主机开启或停止容器服务
func (c *Client) ModifyHostStartTheService(request *ModifyHostStartTheServiceRequest) (response *ModifyHostStartTheServiceResponse, err error) {
	if request == nil {
		request = NewModifyHostStartTheServiceRequest()
	}
	response = NewModifyHostStartTheServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulContainerListRequest() (request *DescribeVulContainerListRequest) {
	request = &DescribeVulContainerListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVulContainerList")
	return
}

func NewDescribeVulContainerListResponse() (response *DescribeVulContainerListResponse) {
	response = &DescribeVulContainerListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询受漏洞的容器列表
func (c *Client) DescribeVulContainerList(request *DescribeVulContainerListRequest) (response *DescribeVulContainerListResponse, err error) {
	if request == nil {
		request = NewDescribeVulContainerListRequest()
	}
	response = NewDescribeVulContainerListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusScanSettingRequest() (request *DescribeVirusScanSettingRequest) {
	request = &DescribeVirusScanSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVirusScanSetting")
	return
}

func NewDescribeVirusScanSettingResponse() (response *DescribeVirusScanSettingResponse) {
	response = &DescribeVirusScanSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时查询文件查杀设置
func (c *Client) DescribeVirusScanSetting(request *DescribeVirusScanSettingRequest) (response *DescribeVirusScanSettingResponse, err error) {
	if request == nil {
		request = NewDescribeVirusScanSettingRequest()
	}
	response = NewDescribeVirusScanSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAffectedClusterCountRequest() (request *DescribeAffectedClusterCountRequest) {
	request = &DescribeAffectedClusterCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAffectedClusterCount")
	return
}

func NewDescribeAffectedClusterCountResponse() (response *DescribeAffectedClusterCountResponse) {
	response = &DescribeAffectedClusterCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取受影响的集群数量，返回数量
func (c *Client) DescribeAffectedClusterCount(request *DescribeAffectedClusterCountRequest) (response *DescribeAffectedClusterCountResponse, err error) {
	if request == nil {
		request = NewDescribeAffectedClusterCountRequest()
	}
	response = NewDescribeAffectedClusterCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEscapeEventInfoRequest() (request *DescribeEscapeEventInfoRequest) {
	request = &DescribeEscapeEventInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeEscapeEventInfo")
	return
}

func NewDescribeEscapeEventInfoResponse() (response *DescribeEscapeEventInfoResponse) {
	response = &DescribeEscapeEventInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器逃逸事件列表
func (c *Client) DescribeEscapeEventInfo(request *DescribeEscapeEventInfoRequest) (response *DescribeEscapeEventInfoResponse, err error) {
	if request == nil {
		request = NewDescribeEscapeEventInfoRequest()
	}
	response = NewDescribeEscapeEventInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskSyscallWhiteListsRequest() (request *DescribeRiskSyscallWhiteListsRequest) {
	request = &DescribeRiskSyscallWhiteListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeRiskSyscallWhiteLists")
	return
}

func NewDescribeRiskSyscallWhiteListsResponse() (response *DescribeRiskSyscallWhiteListsResponse) {
	response = &DescribeRiskSyscallWhiteListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时高危系统调用白名单列表
func (c *Client) DescribeRiskSyscallWhiteLists(request *DescribeRiskSyscallWhiteListsRequest) (response *DescribeRiskSyscallWhiteListsResponse, err error) {
	if request == nil {
		request = NewDescribeRiskSyscallWhiteListsRequest()
	}
	response = NewDescribeRiskSyscallWhiteListsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVirusDetailRequest() (request *DescribeVirusDetailRequest) {
	request = &DescribeVirusDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVirusDetail")
	return
}

func NewDescribeVirusDetailResponse() (response *DescribeVirusDetailResponse) {
	response = &DescribeVirusDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时查询木马文件信息
func (c *Client) DescribeVirusDetail(request *DescribeVirusDetailRequest) (response *DescribeVirusDetailResponse, err error) {
	if request == nil {
		request = NewDescribeVirusDetailRequest()
	}
	response = NewDescribeVirusDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRiskSyscallEventsRequest() (request *DescribeRiskSyscallEventsRequest) {
	request = &DescribeRiskSyscallEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeRiskSyscallEvents")
	return
}

func NewDescribeRiskSyscallEventsResponse() (response *DescribeRiskSyscallEventsResponse) {
	response = &DescribeRiskSyscallEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运行时高危系统调用列表
func (c *Client) DescribeRiskSyscallEvents(request *DescribeRiskSyscallEventsRequest) (response *DescribeRiskSyscallEventsResponse, err error) {
	if request == nil {
		request = NewDescribeRiskSyscallEventsRequest()
	}
	response = NewDescribeRiskSyscallEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterWorkloadWhiteListRequest() (request *DescribeClusterWorkloadWhiteListRequest) {
	request = &DescribeClusterWorkloadWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeClusterWorkloadWhiteList")
	return
}

func NewDescribeClusterWorkloadWhiteListResponse() (response *DescribeClusterWorkloadWhiteListResponse) {
	response = &DescribeClusterWorkloadWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群工作负载白名单
func (c *Client) DescribeClusterWorkloadWhiteList(request *DescribeClusterWorkloadWhiteListRequest) (response *DescribeClusterWorkloadWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeClusterWorkloadWhiteListRequest()
	}
	response = NewDescribeClusterWorkloadWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEventEscapeImageListRequest() (request *DescribeEventEscapeImageListRequest) {
	request = &DescribeEventEscapeImageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeEventEscapeImageList")
	return
}

func NewDescribeEventEscapeImageListResponse() (response *DescribeEventEscapeImageListResponse) {
	response = &DescribeEventEscapeImageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeEventEscapeImageList查询风险容器镜像列表
func (c *Client) DescribeEventEscapeImageList(request *DescribeEventEscapeImageListRequest) (response *DescribeEventEscapeImageListResponse, err error) {
	if request == nil {
		request = NewDescribeEventEscapeImageListRequest()
	}
	response = NewDescribeEventEscapeImageListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterIngressListRequest() (request *DescribeClusterIngressListRequest) {
	request = &DescribeClusterIngressListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeClusterIngressList")
	return
}

func NewDescribeClusterIngressListResponse() (response *DescribeClusterIngressListResponse) {
	response = &DescribeClusterIngressListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群的Ingress列表
func (c *Client) DescribeClusterIngressList(request *DescribeClusterIngressListRequest) (response *DescribeClusterIngressListResponse, err error) {
	if request == nil {
		request = NewDescribeClusterIngressListRequest()
	}
	response = NewDescribeClusterIngressListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateReverseShellWhiteListsExportJobRequest() (request *CreateReverseShellWhiteListsExportJobRequest) {
	request = &CreateReverseShellWhiteListsExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateReverseShellWhiteListsExportJob")
	return
}

func NewCreateReverseShellWhiteListsExportJobResponse() (response *CreateReverseShellWhiteListsExportJobResponse) {
	response = &CreateReverseShellWhiteListsExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建反弹shell白名单列表导出任务
func (c *Client) CreateReverseShellWhiteListsExportJob(request *CreateReverseShellWhiteListsExportJobRequest) (response *CreateReverseShellWhiteListsExportJobResponse, err error) {
	if request == nil {
		request = NewCreateReverseShellWhiteListsExportJobRequest()
	}
	response = NewCreateReverseShellWhiteListsExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVulDetailRequest() (request *DescribeVulDetailRequest) {
	request = &DescribeVulDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeVulDetail")
	return
}

func NewDescribeVulDetailResponse() (response *DescribeVulDetailResponse) {
	response = &DescribeVulDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询漏洞详情
func (c *Client) DescribeVulDetail(request *DescribeVulDetailRequest) (response *DescribeVulDetailResponse, err error) {
	if request == nil {
		request = NewDescribeVulDetailRequest()
	}
	response = NewDescribeVulDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageDenyEventDetailRequest() (request *DescribeImageDenyEventDetailRequest) {
	request = &DescribeImageDenyEventDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeImageDenyEventDetail")
	return
}

func NewDescribeImageDenyEventDetailResponse() (response *DescribeImageDenyEventDetailResponse) {
	response = &DescribeImageDenyEventDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像拦截事件详情
func (c *Client) DescribeImageDenyEventDetail(request *DescribeImageDenyEventDetailRequest) (response *DescribeImageDenyEventDetailResponse, err error) {
	if request == nil {
		request = NewDescribeImageDenyEventDetailRequest()
	}
	response = NewDescribeImageDenyEventDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScanIgnoreVulListRequest() (request *DescribeScanIgnoreVulListRequest) {
	request = &DescribeScanIgnoreVulListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeScanIgnoreVulList")
	return
}

func NewDescribeScanIgnoreVulListResponse() (response *DescribeScanIgnoreVulListResponse) {
	response = &DescribeScanIgnoreVulListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询扫描忽略的漏洞列表
func (c *Client) DescribeScanIgnoreVulList(request *DescribeScanIgnoreVulListRequest) (response *DescribeScanIgnoreVulListResponse, err error) {
	if request == nil {
		request = NewDescribeScanIgnoreVulListRequest()
	}
	response = NewDescribeScanIgnoreVulListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEmergencyVulExportJobRequest() (request *CreateEmergencyVulExportJobRequest) {
	request = &CreateEmergencyVulExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateEmergencyVulExportJob")
	return
}

func NewCreateEmergencyVulExportJobResponse() (response *CreateEmergencyVulExportJobResponse) {
	response = &CreateEmergencyVulExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建应急漏洞导出任务
func (c *Client) CreateEmergencyVulExportJob(request *CreateEmergencyVulExportJobRequest) (response *CreateEmergencyVulExportJobResponse, err error) {
	if request == nil {
		request = NewCreateEmergencyVulExportJobRequest()
	}
	response = NewCreateEmergencyVulExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAssetImageVirusExportJobRequest() (request *CreateAssetImageVirusExportJobRequest) {
	request = &CreateAssetImageVirusExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateAssetImageVirusExportJob")
	return
}

func NewCreateAssetImageVirusExportJobResponse() (response *CreateAssetImageVirusExportJobResponse) {
	response = &CreateAssetImageVirusExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建镜像木马列表导出任务
func (c *Client) CreateAssetImageVirusExportJob(request *CreateAssetImageVirusExportJobRequest) (response *CreateAssetImageVirusExportJobResponse, err error) {
	if request == nil {
		request = NewCreateAssetImageVirusExportJobRequest()
	}
	response = NewCreateAssetImageVirusExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewCreateImageExportJobRequest() (request *CreateImageExportJobRequest) {
	request = &CreateImageExportJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "CreateImageExportJob")
	return
}

func NewCreateImageExportJobResponse() (response *CreateImageExportJobResponse) {
	response = &CreateImageExportJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建镜像导出任务
func (c *Client) CreateImageExportJob(request *CreateImageExportJobRequest) (response *CreateImageExportJobResponse, err error) {
	if request == nil {
		request = NewCreateImageExportJobRequest()
	}
	response = NewCreateImageExportJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIngressForwardConfigRequest() (request *DescribeIngressForwardConfigRequest) {
	request = &DescribeIngressForwardConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeIngressForwardConfig")
	return
}

func NewDescribeIngressForwardConfigResponse() (response *DescribeIngressForwardConfigResponse) {
	response = &DescribeIngressForwardConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Ingress的转发配置信息列表
func (c *Client) DescribeIngressForwardConfig(request *DescribeIngressForwardConfigRequest) (response *DescribeIngressForwardConfigResponse, err error) {
	if request == nil {
		request = NewDescribeIngressForwardConfigRequest()
	}
	response = NewDescribeIngressForwardConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetProcessListRequest() (request *DescribeAssetProcessListRequest) {
	request = &DescribeAssetProcessListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAssetProcessList")
	return
}

func NewDescribeAssetProcessListResponse() (response *DescribeAssetProcessListResponse) {
	response = &DescribeAssetProcessListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询进程列表
func (c *Client) DescribeAssetProcessList(request *DescribeAssetProcessListRequest) (response *DescribeAssetProcessListResponse, err error) {
	if request == nil {
		request = NewDescribeAssetProcessListRequest()
	}
	response = NewDescribeAssetProcessListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEscapeEventDetailRequest() (request *DescribeEscapeEventDetailRequest) {
	request = &DescribeEscapeEventDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeEscapeEventDetail")
	return
}

func NewDescribeEscapeEventDetailResponse() (response *DescribeEscapeEventDetailResponse) {
	response = &DescribeEscapeEventDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器逃逸事件详情
func (c *Client) DescribeEscapeEventDetail(request *DescribeEscapeEventDetailRequest) (response *DescribeEscapeEventDetailResponse, err error) {
	if request == nil {
		request = NewDescribeEscapeEventDetailRequest()
	}
	response = NewDescribeEscapeEventDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddImageDenyRuleInfoRequest() (request *DescribeAddImageDenyRuleInfoRequest) {
	request = &DescribeAddImageDenyRuleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("otcss", APIVersion, "DescribeAddImageDenyRuleInfo")
	return
}

func NewDescribeAddImageDenyRuleInfoResponse() (response *DescribeAddImageDenyRuleInfoResponse) {
	response = &DescribeAddImageDenyRuleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户新增镜像拦截规则信息
func (c *Client) DescribeAddImageDenyRuleInfo(request *DescribeAddImageDenyRuleInfoRequest) (response *DescribeAddImageDenyRuleInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAddImageDenyRuleInfoRequest()
	}
	response = NewDescribeAddImageDenyRuleInfoResponse()
	err = c.Send(request, response)
	return
}
