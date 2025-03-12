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

package v20200630

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-06-30"

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

func NewSearchMetaValuesRequest() (request *SearchMetaValuesRequest) {
	request = &SearchMetaValuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "SearchMetaValues")
	return
}

func NewSearchMetaValuesResponse() (response *SearchMetaValuesResponse) {
	response = &SearchMetaValuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 模糊查询MetaValue
func (c *Client) SearchMetaValues(request *SearchMetaValuesRequest) (response *SearchMetaValuesResponse, err error) {
	if request == nil {
		request = NewSearchMetaValuesRequest()
	}
	response = NewSearchMetaValuesResponse()
	err = c.Send(request, response)
	return
}

func NewSearchLogRuleGroupsRequest() (request *SearchLogRuleGroupsRequest) {
	request = &SearchLogRuleGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "SearchLogRuleGroups")
	return
}

func NewSearchLogRuleGroupsResponse() (response *SearchLogRuleGroupsResponse) {
	response = &SearchLogRuleGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件查询日志告警策略
func (c *Client) SearchLogRuleGroups(request *SearchLogRuleGroupsRequest) (response *SearchLogRuleGroupsResponse, err error) {
	if request == nil {
		request = NewSearchLogRuleGroupsRequest()
	}
	response = NewSearchLogRuleGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewGetAlertStatsRequest() (request *GetAlertStatsRequest) {
	request = &GetAlertStatsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetAlertStats")
	return
}

func NewGetAlertStatsResponse() (response *GetAlertStatsResponse) {
	response = &GetAlertStatsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取报警概述
func (c *Client) GetAlertStats(request *GetAlertStatsRequest) (response *GetAlertStatsResponse, err error) {
	if request == nil {
		request = NewGetAlertStatsRequest()
	}
	response = NewGetAlertStatsResponse()
	err = c.Send(request, response)
	return
}

func NewListLogDeliveryTargetRequest() (request *ListLogDeliveryTargetRequest) {
	request = &ListLogDeliveryTargetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ListLogDeliveryTarget")
	return
}

func NewListLogDeliveryTargetResponse() (response *ListLogDeliveryTargetResponse) {
	response = &ListLogDeliveryTargetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志投递下游源列表
func (c *Client) ListLogDeliveryTarget(request *ListLogDeliveryTargetRequest) (response *ListLogDeliveryTargetResponse, err error) {
	if request == nil {
		request = NewListLogDeliveryTargetRequest()
	}
	response = NewListLogDeliveryTargetResponse()
	err = c.Send(request, response)
	return
}

func NewAddMetaMetricRequest() (request *AddMetaMetricRequest) {
	request = &AddMetaMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "AddMetaMetric")
	return
}

func NewAddMetaMetricResponse() (response *AddMetaMetricResponse) {
	response = &AddMetaMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在资源对象类型下创建指标
func (c *Client) AddMetaMetric(request *AddMetaMetricRequest) (response *AddMetaMetricResponse, err error) {
	if request == nil {
		request = NewAddMetaMetricRequest()
	}
	response = NewAddMetaMetricResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogRuleGroupTemplesRequest() (request *DeleteLogRuleGroupTemplesRequest) {
	request = &DeleteLogRuleGroupTemplesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteLogRuleGroupTemples")
	return
}

func NewDeleteLogRuleGroupTemplesResponse() (response *DeleteLogRuleGroupTemplesResponse) {
	response = &DeleteLogRuleGroupTemplesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除日志告警策略模板
func (c *Client) DeleteLogRuleGroupTemples(request *DeleteLogRuleGroupTemplesRequest) (response *DeleteLogRuleGroupTemplesResponse, err error) {
	if request == nil {
		request = NewDeleteLogRuleGroupTemplesRequest()
	}
	response = NewDeleteLogRuleGroupTemplesResponse()
	err = c.Send(request, response)
	return
}

func NewListLogDeliveryTargetTypeRequest() (request *ListLogDeliveryTargetTypeRequest) {
	request = &ListLogDeliveryTargetTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ListLogDeliveryTargetType")
	return
}

func NewListLogDeliveryTargetTypeResponse() (response *ListLogDeliveryTargetTypeResponse) {
	response = &ListLogDeliveryTargetTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志投递下游源类型列表
func (c *Client) ListLogDeliveryTargetType(request *ListLogDeliveryTargetTypeRequest) (response *ListLogDeliveryTargetTypeResponse, err error) {
	if request == nil {
		request = NewListLogDeliveryTargetTypeRequest()
	}
	response = NewListLogDeliveryTargetTypeResponse()
	err = c.Send(request, response)
	return
}

func NewAddMetaResourceTypeRequest() (request *AddMetaResourceTypeRequest) {
	request = &AddMetaResourceTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "AddMetaResourceType")
	return
}

func NewAddMetaResourceTypeResponse() (response *AddMetaResourceTypeResponse) {
	response = &AddMetaResourceTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建资源对象类型
func (c *Client) AddMetaResourceType(request *AddMetaResourceTypeRequest) (response *AddMetaResourceTypeResponse, err error) {
	if request == nil {
		request = NewAddMetaResourceTypeRequest()
	}
	response = NewAddMetaResourceTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRuleGroupTemplesRequest() (request *DeleteRuleGroupTemplesRequest) {
	request = &DeleteRuleGroupTemplesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteRuleGroupTemples")
	return
}

func NewDeleteRuleGroupTemplesResponse() (response *DeleteRuleGroupTemplesResponse) {
	response = &DeleteRuleGroupTemplesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除策略模板
func (c *Client) DeleteRuleGroupTemples(request *DeleteRuleGroupTemplesRequest) (response *DeleteRuleGroupTemplesResponse, err error) {
	if request == nil {
		request = NewDeleteRuleGroupTemplesRequest()
	}
	response = NewDeleteRuleGroupTemplesResponse()
	err = c.Send(request, response)
	return
}

func NewGetNotificationRequest() (request *GetNotificationRequest) {
	request = &GetNotificationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetNotification")
	return
}

func NewGetNotificationResponse() (response *GetNotificationResponse) {
	response = &GetNotificationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定notification
func (c *Client) GetNotification(request *GetNotificationRequest) (response *GetNotificationResponse, err error) {
	if request == nil {
		request = NewGetNotificationRequest()
	}
	response = NewGetNotificationResponse()
	err = c.Send(request, response)
	return
}

func NewSearchRuleGroupTemplesRequest() (request *SearchRuleGroupTemplesRequest) {
	request = &SearchRuleGroupTemplesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "SearchRuleGroupTemples")
	return
}

func NewSearchRuleGroupTemplesResponse() (response *SearchRuleGroupTemplesResponse) {
	response = &SearchRuleGroupTemplesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件查询策略模板
func (c *Client) SearchRuleGroupTemples(request *SearchRuleGroupTemplesRequest) (response *SearchRuleGroupTemplesResponse, err error) {
	if request == nil {
		request = NewSearchRuleGroupTemplesRequest()
	}
	response = NewSearchRuleGroupTemplesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateLogConfigStorageStateRequest() (request *UpdateLogConfigStorageStateRequest) {
	request = &UpdateLogConfigStorageStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "UpdateLogConfigStorageState")
	return
}

func NewUpdateLogConfigStorageStateResponse() (response *UpdateLogConfigStorageStateResponse) {
	response = &UpdateLogConfigStorageStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改日志配置后端存储开关
func (c *Client) UpdateLogConfigStorageState(request *UpdateLogConfigStorageStateRequest) (response *UpdateLogConfigStorageStateResponse, err error) {
	if request == nil {
		request = NewUpdateLogConfigStorageStateRequest()
	}
	response = NewUpdateLogConfigStorageStateResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSilencesRequest() (request *DeleteSilencesRequest) {
	request = &DeleteSilencesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteSilences")
	return
}

func NewDeleteSilencesResponse() (response *DeleteSilencesResponse) {
	response = &DeleteSilencesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除静默规则
func (c *Client) DeleteSilences(request *DeleteSilencesRequest) (response *DeleteSilencesResponse, err error) {
	if request == nil {
		request = NewDeleteSilencesRequest()
	}
	response = NewDeleteSilencesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMetaResourceTypeRequest() (request *UpdateMetaResourceTypeRequest) {
	request = &UpdateMetaResourceTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "UpdateMetaResourceType")
	return
}

func NewUpdateMetaResourceTypeResponse() (response *UpdateMetaResourceTypeResponse) {
	response = &UpdateMetaResourceTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新资源对象类型信息
func (c *Client) UpdateMetaResourceType(request *UpdateMetaResourceTypeRequest) (response *UpdateMetaResourceTypeResponse, err error) {
	if request == nil {
		request = NewUpdateMetaResourceTypeRequest()
	}
	response = NewUpdateMetaResourceTypeResponse()
	err = c.Send(request, response)
	return
}

func NewSearchRoutesByLabelSetRequest() (request *SearchRoutesByLabelSetRequest) {
	request = &SearchRoutesByLabelSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "SearchRoutesByLabelSet")
	return
}

func NewSearchRoutesByLabelSetResponse() (response *SearchRoutesByLabelSetResponse) {
	response = &SearchRoutesByLabelSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件查询匹配的路由信息
func (c *Client) SearchRoutesByLabelSet(request *SearchRoutesByLabelSetRequest) (response *SearchRoutesByLabelSetResponse, err error) {
	if request == nil {
		request = NewSearchRoutesByLabelSetRequest()
	}
	response = NewSearchRoutesByLabelSetResponse()
	err = c.Send(request, response)
	return
}

func NewGetDashFolderByUidRequest() (request *GetDashFolderByUidRequest) {
	request = &GetDashFolderByUidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetDashFolderByUid")
	return
}

func NewGetDashFolderByUidResponse() (response *GetDashFolderByUidResponse) {
	response = &GetDashFolderByUidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过唯一ID获取详情
func (c *Client) GetDashFolderByUid(request *GetDashFolderByUidRequest) (response *GetDashFolderByUidResponse, err error) {
	if request == nil {
		request = NewGetDashFolderByUidRequest()
	}
	response = NewGetDashFolderByUidResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrUpdateDashboardRequest() (request *CreateOrUpdateDashboardRequest) {
	request = &CreateOrUpdateDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "CreateOrUpdateDashboard")
	return
}

func NewCreateOrUpdateDashboardResponse() (response *CreateOrUpdateDashboardResponse) {
	response = &CreateOrUpdateDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建或更新仪表盘
func (c *Client) CreateOrUpdateDashboard(request *CreateOrUpdateDashboardRequest) (response *CreateOrUpdateDashboardResponse, err error) {
	if request == nil {
		request = NewCreateOrUpdateDashboardRequest()
	}
	response = NewCreateOrUpdateDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogPanelsRequest() (request *DeleteLogPanelsRequest) {
	request = &DeleteLogPanelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteLogPanels")
	return
}

func NewDeleteLogPanelsResponse() (response *DeleteLogPanelsResponse) {
	response = &DeleteLogPanelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除日志仪表盘面板
func (c *Client) DeleteLogPanels(request *DeleteLogPanelsRequest) (response *DeleteLogPanelsResponse, err error) {
	if request == nil {
		request = NewDeleteLogPanelsRequest()
	}
	response = NewDeleteLogPanelsResponse()
	err = c.Send(request, response)
	return
}

func NewGetFiringRulesRequest() (request *GetFiringRulesRequest) {
	request = &GetFiringRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetFiringRules")
	return
}

func NewGetFiringRulesResponse() (response *GetFiringRulesResponse) {
	response = &GetFiringRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取有未恢复告警的规则
func (c *Client) GetFiringRules(request *GetFiringRulesRequest) (response *GetFiringRulesResponse, err error) {
	if request == nil {
		request = NewGetFiringRulesRequest()
	}
	response = NewGetFiringRulesResponse()
	err = c.Send(request, response)
	return
}

func NewApplyInfraTemplatesRequest() (request *ApplyInfraTemplatesRequest) {
	request = &ApplyInfraTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ApplyInfraTemplates")
	return
}

func NewApplyInfraTemplatesResponse() (response *ApplyInfraTemplatesResponse) {
	response = &ApplyInfraTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云哨模板批量应用接口
func (c *Client) ApplyInfraTemplates(request *ApplyInfraTemplatesRequest) (response *ApplyInfraTemplatesResponse, err error) {
	if request == nil {
		request = NewApplyInfraTemplatesRequest()
	}
	response = NewApplyInfraTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateQueryRecordRequest() (request *CreateQueryRecordRequest) {
	request = &CreateQueryRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "CreateQueryRecord")
	return
}

func NewCreateQueryRecordResponse() (response *CreateQueryRecordResponse) {
	response = &CreateQueryRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存日志查询记录
func (c *Client) CreateQueryRecord(request *CreateQueryRecordRequest) (response *CreateQueryRecordResponse, err error) {
	if request == nil {
		request = NewCreateQueryRecordRequest()
	}
	response = NewCreateQueryRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteResourceObjectRequest() (request *DeleteResourceObjectRequest) {
	request = &DeleteResourceObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteResourceObject")
	return
}

func NewDeleteResourceObjectResponse() (response *DeleteResourceObjectResponse) {
	response = &DeleteResourceObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DeleteResourceObject
func (c *Client) DeleteResourceObject(request *DeleteResourceObjectRequest) (response *DeleteResourceObjectResponse, err error) {
	if request == nil {
		request = NewDeleteResourceObjectRequest()
	}
	response = NewDeleteResourceObjectResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSilenceRequest() (request *UpdateSilenceRequest) {
	request = &UpdateSilenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "UpdateSilence")
	return
}

func NewUpdateSilenceResponse() (response *UpdateSilenceResponse) {
	response = &UpdateSilenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新屏蔽策略
func (c *Client) UpdateSilence(request *UpdateSilenceRequest) (response *UpdateSilenceResponse, err error) {
	if request == nil {
		request = NewUpdateSilenceRequest()
	}
	response = NewUpdateSilenceResponse()
	err = c.Send(request, response)
	return
}

func NewGetDashboardVersionContentByIdRequest() (request *GetDashboardVersionContentByIdRequest) {
	request = &GetDashboardVersionContentByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetDashboardVersionContentById")
	return
}

func NewGetDashboardVersionContentByIdResponse() (response *GetDashboardVersionContentByIdResponse) {
	response = &GetDashboardVersionContentByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看给定仪表盘历史版本的内容
func (c *Client) GetDashboardVersionContentById(request *GetDashboardVersionContentByIdRequest) (response *GetDashboardVersionContentByIdResponse, err error) {
	if request == nil {
		request = NewGetDashboardVersionContentByIdRequest()
	}
	response = NewGetDashboardVersionContentByIdResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrUpdateLogRuleGroupRequest() (request *CreateOrUpdateLogRuleGroupRequest) {
	request = &CreateOrUpdateLogRuleGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "CreateOrUpdateLogRuleGroup")
	return
}

func NewCreateOrUpdateLogRuleGroupResponse() (response *CreateOrUpdateLogRuleGroupResponse) {
	response = &CreateOrUpdateLogRuleGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建日志报警策略
func (c *Client) CreateOrUpdateLogRuleGroup(request *CreateOrUpdateLogRuleGroupRequest) (response *CreateOrUpdateLogRuleGroupResponse, err error) {
	if request == nil {
		request = NewCreateOrUpdateLogRuleGroupRequest()
	}
	response = NewCreateOrUpdateLogRuleGroupResponse()
	err = c.Send(request, response)
	return
}

func NewGetNotificationMsgRequest() (request *GetNotificationMsgRequest) {
	request = &GetNotificationMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetNotificationMsg")
	return
}

func NewGetNotificationMsgResponse() (response *GetNotificationMsgResponse) {
	response = &GetNotificationMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取通知消息
func (c *Client) GetNotificationMsg(request *GetNotificationMsgRequest) (response *GetNotificationMsgResponse, err error) {
	if request == nil {
		request = NewGetNotificationMsgRequest()
	}
	response = NewGetNotificationMsgResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteQueryRecordRequest() (request *DeleteQueryRecordRequest) {
	request = &DeleteQueryRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteQueryRecord")
	return
}

func NewDeleteQueryRecordResponse() (response *DeleteQueryRecordResponse) {
	response = &DeleteQueryRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除查询记录
func (c *Client) DeleteQueryRecord(request *DeleteQueryRecordRequest) (response *DeleteQueryRecordResponse, err error) {
	if request == nil {
		request = NewDeleteQueryRecordRequest()
	}
	response = NewDeleteQueryRecordResponse()
	err = c.Send(request, response)
	return
}

func NewProxyAllDataSourceRequestByNameRequest() (request *ProxyAllDataSourceRequestByNameRequest) {
	request = &ProxyAllDataSourceRequestByNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ProxyAllDataSourceRequestByName")
	return
}

func NewProxyAllDataSourceRequestByNameResponse() (response *ProxyAllDataSourceRequestByNameResponse) {
	response = &ProxyAllDataSourceRequestByNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据名称自动代理监控请求
func (c *Client) ProxyAllDataSourceRequestByName(request *ProxyAllDataSourceRequestByNameRequest) (response *ProxyAllDataSourceRequestByNameResponse, err error) {
	if request == nil {
		request = NewProxyAllDataSourceRequestByNameRequest()
	}
	response = NewProxyAllDataSourceRequestByNameResponse()
	err = c.Send(request, response)
	return
}

func NewGetAlertTrendRequest() (request *GetAlertTrendRequest) {
	request = &GetAlertTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetAlertTrend")
	return
}

func NewGetAlertTrendResponse() (response *GetAlertTrendResponse) {
	response = &GetAlertTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警趋势接口
func (c *Client) GetAlertTrend(request *GetAlertTrendRequest) (response *GetAlertTrendResponse, err error) {
	if request == nil {
		request = NewGetAlertTrendRequest()
	}
	response = NewGetAlertTrendResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogConfigRequest() (request *DeleteLogConfigRequest) {
	request = &DeleteLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteLogConfig")
	return
}

func NewDeleteLogConfigResponse() (response *DeleteLogConfigResponse) {
	response = &DeleteLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除日志接入配置
func (c *Client) DeleteLogConfig(request *DeleteLogConfigRequest) (response *DeleteLogConfigResponse, err error) {
	if request == nil {
		request = NewDeleteLogConfigRequest()
	}
	response = NewDeleteLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewSearchLogsRequest() (request *SearchLogsRequest) {
	request = &SearchLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "SearchLogs")
	return
}

func NewSearchLogsResponse() (response *SearchLogsResponse) {
	response = &SearchLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志查询
func (c *Client) SearchLogs(request *SearchLogsRequest) (response *SearchLogsResponse, err error) {
	if request == nil {
		request = NewSearchLogsRequest()
	}
	response = NewSearchLogsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrUpdateInfraTemplateRequest() (request *CreateOrUpdateInfraTemplateRequest) {
	request = &CreateOrUpdateInfraTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "CreateOrUpdateInfraTemplate")
	return
}

func NewCreateOrUpdateInfraTemplateResponse() (response *CreateOrUpdateInfraTemplateResponse) {
	response = &CreateOrUpdateInfraTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云哨模板创建或更新接口
func (c *Client) CreateOrUpdateInfraTemplate(request *CreateOrUpdateInfraTemplateRequest) (response *CreateOrUpdateInfraTemplateResponse, err error) {
	if request == nil {
		request = NewCreateOrUpdateInfraTemplateRequest()
	}
	response = NewCreateOrUpdateInfraTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewGetRelatedAlertNamesRequest() (request *GetRelatedAlertNamesRequest) {
	request = &GetRelatedAlertNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetRelatedAlertNames")
	return
}

func NewGetRelatedAlertNamesResponse() (response *GetRelatedAlertNamesResponse) {
	response = &GetRelatedAlertNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量通过事件名获取关联报警列表
func (c *Client) GetRelatedAlertNames(request *GetRelatedAlertNamesRequest) (response *GetRelatedAlertNamesResponse, err error) {
	if request == nil {
		request = NewGetRelatedAlertNamesRequest()
	}
	response = NewGetRelatedAlertNamesResponse()
	err = c.Send(request, response)
	return
}

func NewSearchAlertsRequest() (request *SearchAlertsRequest) {
	request = &SearchAlertsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "SearchAlerts")
	return
}

func NewSearchAlertsResponse() (response *SearchAlertsResponse) {
	response = &SearchAlertsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件搜索Alert
// 分页page、limit、asc/desc
func (c *Client) SearchAlerts(request *SearchAlertsRequest) (response *SearchAlertsResponse, err error) {
	if request == nil {
		request = NewSearchAlertsRequest()
	}
	response = NewSearchAlertsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSilenceRequest() (request *CreateSilenceRequest) {
	request = &CreateSilenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "CreateSilence")
	return
}

func NewCreateSilenceResponse() (response *CreateSilenceResponse) {
	response = &CreateSilenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建屏蔽策略
func (c *Client) CreateSilence(request *CreateSilenceRequest) (response *CreateSilenceResponse, err error) {
	if request == nil {
		request = NewCreateSilenceRequest()
	}
	response = NewCreateSilenceResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateResourceObjectRequest() (request *UpdateResourceObjectRequest) {
	request = &UpdateResourceObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "UpdateResourceObject")
	return
}

func NewUpdateResourceObjectResponse() (response *UpdateResourceObjectResponse) {
	response = &UpdateResourceObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新服务树节点资源对象
func (c *Client) UpdateResourceObject(request *UpdateResourceObjectRequest) (response *UpdateResourceObjectResponse, err error) {
	if request == nil {
		request = NewUpdateResourceObjectRequest()
	}
	response = NewUpdateResourceObjectResponse()
	err = c.Send(request, response)
	return
}

func NewLogMappingRequest() (request *LogMappingRequest) {
	request = &LogMappingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "LogMapping")
	return
}

func NewLogMappingResponse() (response *LogMappingResponse) {
	response = &LogMappingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志字段映射
func (c *Client) LogMapping(request *LogMappingRequest) (response *LogMappingResponse, err error) {
	if request == nil {
		request = NewLogMappingRequest()
	}
	response = NewLogMappingResponse()
	err = c.Send(request, response)
	return
}

func NewSearchNotificationMsgsRequest() (request *SearchNotificationMsgsRequest) {
	request = &SearchNotificationMsgsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "SearchNotificationMsgs")
	return
}

func NewSearchNotificationMsgsResponse() (response *SearchNotificationMsgsResponse) {
	response = &SearchNotificationMsgsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 消息查询接口
func (c *Client) SearchNotificationMsgs(request *SearchNotificationMsgsRequest) (response *SearchNotificationMsgsResponse, err error) {
	if request == nil {
		request = NewSearchNotificationMsgsRequest()
	}
	response = NewSearchNotificationMsgsResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceObjectRequest() (request *ListResourceObjectRequest) {
	request = &ListResourceObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ListResourceObject")
	return
}

func NewListResourceObjectResponse() (response *ListResourceObjectResponse) {
	response = &ListResourceObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务树资源对象
func (c *Client) ListResourceObject(request *ListResourceObjectRequest) (response *ListResourceObjectResponse, err error) {
	if request == nil {
		request = NewListResourceObjectRequest()
	}
	response = NewListResourceObjectResponse()
	err = c.Send(request, response)
	return
}

func NewProxyTypeInfoRequest() (request *ProxyTypeInfoRequest) {
	request = &ProxyTypeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ProxyTypeInfo")
	return
}

func NewProxyTypeInfoResponse() (response *ProxyTypeInfoResponse) {
	response = &ProxyTypeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 对象信息代理
func (c *Client) ProxyTypeInfo(request *ProxyTypeInfoRequest) (response *ProxyTypeInfoResponse, err error) {
	if request == nil {
		request = NewProxyTypeInfoRequest()
	}
	response = NewProxyTypeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrUpdateLogRuleGroupTempleRequest() (request *CreateOrUpdateLogRuleGroupTempleRequest) {
	request = &CreateOrUpdateLogRuleGroupTempleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "CreateOrUpdateLogRuleGroupTemple")
	return
}

func NewCreateOrUpdateLogRuleGroupTempleResponse() (response *CreateOrUpdateLogRuleGroupTempleResponse) {
	response = &CreateOrUpdateLogRuleGroupTempleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建日报报警策略模板
func (c *Client) CreateOrUpdateLogRuleGroupTemple(request *CreateOrUpdateLogRuleGroupTempleRequest) (response *CreateOrUpdateLogRuleGroupTempleResponse, err error) {
	if request == nil {
		request = NewCreateOrUpdateLogRuleGroupTempleRequest()
	}
	response = NewCreateOrUpdateLogRuleGroupTempleResponse()
	err = c.Send(request, response)
	return
}

func NewGenRegexRequest() (request *GenRegexRequest) {
	request = &GenRegexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GenRegex")
	return
}

func NewGenRegexResponse() (response *GenRegexResponse) {
	response = &GenRegexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 自动生成正则表达式
func (c *Client) GenRegex(request *GenRegexRequest) (response *GenRegexResponse, err error) {
	if request == nil {
		request = NewGenRegexRequest()
	}
	response = NewGenRegexResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogRuleGroupsRequest() (request *DeleteLogRuleGroupsRequest) {
	request = &DeleteLogRuleGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteLogRuleGroups")
	return
}

func NewDeleteLogRuleGroupsResponse() (response *DeleteLogRuleGroupsResponse) {
	response = &DeleteLogRuleGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除日志报警策略
func (c *Client) DeleteLogRuleGroups(request *DeleteLogRuleGroupsRequest) (response *DeleteLogRuleGroupsResponse, err error) {
	if request == nil {
		request = NewDeleteLogRuleGroupsRequest()
	}
	response = NewDeleteLogRuleGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewListMetaResourceTypeRequest() (request *ListMetaResourceTypeRequest) {
	request = &ListMetaResourceTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ListMetaResourceType")
	return
}

func NewListMetaResourceTypeResponse() (response *ListMetaResourceTypeResponse) {
	response = &ListMetaResourceTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资源对象类型列表
func (c *Client) ListMetaResourceType(request *ListMetaResourceTypeRequest) (response *ListMetaResourceTypeResponse, err error) {
	if request == nil {
		request = NewListMetaResourceTypeRequest()
	}
	response = NewListMetaResourceTypeResponse()
	err = c.Send(request, response)
	return
}

func NewSearchInfraTemplatesRefsRequest() (request *SearchInfraTemplatesRefsRequest) {
	request = &SearchInfraTemplatesRefsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "SearchInfraTemplatesRefs")
	return
}

func NewSearchInfraTemplatesRefsResponse() (response *SearchInfraTemplatesRefsResponse) {
	response = &SearchInfraTemplatesRefsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云哨模板关联实例查询接口
func (c *Client) SearchInfraTemplatesRefs(request *SearchInfraTemplatesRefsRequest) (response *SearchInfraTemplatesRefsResponse, err error) {
	if request == nil {
		request = NewSearchInfraTemplatesRefsRequest()
	}
	response = NewSearchInfraTemplatesRefsResponse()
	err = c.Send(request, response)
	return
}

func NewSearchNotificationsRequest() (request *SearchNotificationsRequest) {
	request = &SearchNotificationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "SearchNotifications")
	return
}

func NewSearchNotificationsResponse() (response *SearchNotificationsResponse) {
	response = &SearchNotificationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取聚合列表
func (c *Client) SearchNotifications(request *SearchNotificationsRequest) (response *SearchNotificationsResponse, err error) {
	if request == nil {
		request = NewSearchNotificationsRequest()
	}
	response = NewSearchNotificationsResponse()
	err = c.Send(request, response)
	return
}

func NewGetDashboardVersionsByIdRequest() (request *GetDashboardVersionsByIdRequest) {
	request = &GetDashboardVersionsByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetDashboardVersionsById")
	return
}

func NewGetDashboardVersionsByIdResponse() (response *GetDashboardVersionsByIdResponse) {
	response = &GetDashboardVersionsByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根椐仪表ID查询变更历史
func (c *Client) GetDashboardVersionsById(request *GetDashboardVersionsByIdRequest) (response *GetDashboardVersionsByIdResponse, err error) {
	if request == nil {
		request = NewGetDashboardVersionsByIdRequest()
	}
	response = NewGetDashboardVersionsByIdResponse()
	err = c.Send(request, response)
	return
}

func NewVerifyLogConfigRequest() (request *VerifyLogConfigRequest) {
	request = &VerifyLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "VerifyLogConfig")
	return
}

func NewVerifyLogConfigResponse() (response *VerifyLogConfigResponse) {
	response = &VerifyLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验日志接入配置
func (c *Client) VerifyLogConfig(request *VerifyLogConfigRequest) (response *VerifyLogConfigResponse, err error) {
	if request == nil {
		request = NewVerifyLogConfigRequest()
	}
	response = NewVerifyLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateLogTransformTaskRequest() (request *UpdateLogTransformTaskRequest) {
	request = &UpdateLogTransformTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "UpdateLogTransformTask")
	return
}

func NewUpdateLogTransformTaskResponse() (response *UpdateLogTransformTaskResponse) {
	response = &UpdateLogTransformTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新日志任务处理
func (c *Client) UpdateLogTransformTask(request *UpdateLogTransformTaskRequest) (response *UpdateLogTransformTaskResponse, err error) {
	if request == nil {
		request = NewUpdateLogTransformTaskRequest()
	}
	response = NewUpdateLogTransformTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDashboardByUidRequest() (request *DeleteDashboardByUidRequest) {
	request = &DeleteDashboardByUidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteDashboardByUid")
	return
}

func NewDeleteDashboardByUidResponse() (response *DeleteDashboardByUidResponse) {
	response = &DeleteDashboardByUidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根椐唯一ID删除仪表盘
func (c *Client) DeleteDashboardByUid(request *DeleteDashboardByUidRequest) (response *DeleteDashboardByUidResponse, err error) {
	if request == nil {
		request = NewDeleteDashboardByUidRequest()
	}
	response = NewDeleteDashboardByUidResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateLogConfigStateRequest() (request *UpdateLogConfigStateRequest) {
	request = &UpdateLogConfigStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "UpdateLogConfigState")
	return
}

func NewUpdateLogConfigStateResponse() (response *UpdateLogConfigStateResponse) {
	response = &UpdateLogConfigStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志配置状态变更接口 ，启用或禁用，支持批量
func (c *Client) UpdateLogConfigState(request *UpdateLogConfigStateRequest) (response *UpdateLogConfigStateResponse, err error) {
	if request == nil {
		request = NewUpdateLogConfigStateRequest()
	}
	response = NewUpdateLogConfigStateResponse()
	err = c.Send(request, response)
	return
}

func NewGetSilenceRequest() (request *GetSilenceRequest) {
	request = &GetSilenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetSilence")
	return
}

func NewGetSilenceResponse() (response *GetSilenceResponse) {
	response = &GetSilenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取屏蔽规则
func (c *Client) GetSilence(request *GetSilenceRequest) (response *GetSilenceResponse, err error) {
	if request == nil {
		request = NewGetSilenceRequest()
	}
	response = NewGetSilenceResponse()
	err = c.Send(request, response)
	return
}

func NewSearchRoutesRequest() (request *SearchRoutesRequest) {
	request = &SearchRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "SearchRoutes")
	return
}

func NewSearchRoutesResponse() (response *SearchRoutesResponse) {
	response = &SearchRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过名称搜索路由
func (c *Client) SearchRoutes(request *SearchRoutesRequest) (response *SearchRoutesResponse, err error) {
	if request == nil {
		request = NewSearchRoutesRequest()
	}
	response = NewSearchRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewListLogDeliveryTaskRequest() (request *ListLogDeliveryTaskRequest) {
	request = &ListLogDeliveryTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ListLogDeliveryTask")
	return
}

func NewListLogDeliveryTaskResponse() (response *ListLogDeliveryTaskResponse) {
	response = &ListLogDeliveryTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志投递任务列表
func (c *Client) ListLogDeliveryTask(request *ListLogDeliveryTaskRequest) (response *ListLogDeliveryTaskResponse, err error) {
	if request == nil {
		request = NewListLogDeliveryTaskRequest()
	}
	response = NewListLogDeliveryTaskResponse()
	err = c.Send(request, response)
	return
}

func NewListLogTransformTaskRequest() (request *ListLogTransformTaskRequest) {
	request = &ListLogTransformTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ListLogTransformTask")
	return
}

func NewListLogTransformTaskResponse() (response *ListLogTransformTaskResponse) {
	response = &ListLogTransformTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取日志处理任务列表
func (c *Client) ListLogTransformTask(request *ListLogTransformTaskRequest) (response *ListLogTransformTaskResponse, err error) {
	if request == nil {
		request = NewListLogTransformTaskRequest()
	}
	response = NewListLogTransformTaskResponse()
	err = c.Send(request, response)
	return
}

func NewVisualizeLogsRequest() (request *VisualizeLogsRequest) {
	request = &VisualizeLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "VisualizeLogs")
	return
}

func NewVisualizeLogsResponse() (response *VisualizeLogsResponse) {
	response = &VisualizeLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志分析
func (c *Client) VisualizeLogs(request *VisualizeLogsRequest) (response *VisualizeLogsResponse, err error) {
	if request == nil {
		request = NewVisualizeLogsRequest()
	}
	response = NewVisualizeLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogAlertRuleRequest() (request *DeleteLogAlertRuleRequest) {
	request = &DeleteLogAlertRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteLogAlertRule")
	return
}

func NewDeleteLogAlertRuleResponse() (response *DeleteLogAlertRuleResponse) {
	response = &DeleteLogAlertRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除日志报警
func (c *Client) DeleteLogAlertRule(request *DeleteLogAlertRuleRequest) (response *DeleteLogAlertRuleResponse, err error) {
	if request == nil {
		request = NewDeleteLogAlertRuleRequest()
	}
	response = NewDeleteLogAlertRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMetaResourceTypeRequest() (request *DeleteMetaResourceTypeRequest) {
	request = &DeleteMetaResourceTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteMetaResourceType")
	return
}

func NewDeleteMetaResourceTypeResponse() (response *DeleteMetaResourceTypeResponse) {
	response = &DeleteMetaResourceTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除资源对象类型对象
func (c *Client) DeleteMetaResourceType(request *DeleteMetaResourceTypeRequest) (response *DeleteMetaResourceTypeResponse, err error) {
	if request == nil {
		request = NewDeleteMetaResourceTypeRequest()
	}
	response = NewDeleteMetaResourceTypeResponse()
	err = c.Send(request, response)
	return
}

func NewGetRouteRequest() (request *GetRouteRequest) {
	request = &GetRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetRoute")
	return
}

func NewGetRouteResponse() (response *GetRouteResponse) {
	response = &GetRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定路由信息
func (c *Client) GetRoute(request *GetRouteRequest) (response *GetRouteResponse, err error) {
	if request == nil {
		request = NewGetRouteRequest()
	}
	response = NewGetRouteResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateLogAlertRuleRequest() (request *UpdateLogAlertRuleRequest) {
	request = &UpdateLogAlertRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "UpdateLogAlertRule")
	return
}

func NewUpdateLogAlertRuleResponse() (response *UpdateLogAlertRuleResponse) {
	response = &UpdateLogAlertRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新日志报警策略
func (c *Client) UpdateLogAlertRule(request *UpdateLogAlertRuleRequest) (response *UpdateLogAlertRuleResponse, err error) {
	if request == nil {
		request = NewUpdateLogAlertRuleRequest()
	}
	response = NewUpdateLogAlertRuleResponse()
	err = c.Send(request, response)
	return
}

func NewAnalyseLogsRequest() (request *AnalyseLogsRequest) {
	request = &AnalyseLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "AnalyseLogs")
	return
}

func NewAnalyseLogsResponse() (response *AnalyseLogsResponse) {
	response = &AnalyseLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志分析
func (c *Client) AnalyseLogs(request *AnalyseLogsRequest) (response *AnalyseLogsResponse, err error) {
	if request == nil {
		request = NewAnalyseLogsRequest()
	}
	response = NewAnalyseLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMetaMetricRequest() (request *DeleteMetaMetricRequest) {
	request = &DeleteMetaMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteMetaMetric")
	return
}

func NewDeleteMetaMetricResponse() (response *DeleteMetaMetricResponse) {
	response = &DeleteMetaMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除指标
func (c *Client) DeleteMetaMetric(request *DeleteMetaMetricRequest) (response *DeleteMetaMetricResponse, err error) {
	if request == nil {
		request = NewDeleteMetaMetricRequest()
	}
	response = NewDeleteMetaMetricResponse()
	err = c.Send(request, response)
	return
}

func NewSearchLogRuleGroupTemplesRequest() (request *SearchLogRuleGroupTemplesRequest) {
	request = &SearchLogRuleGroupTemplesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "SearchLogRuleGroupTemples")
	return
}

func NewSearchLogRuleGroupTemplesResponse() (response *SearchLogRuleGroupTemplesResponse) {
	response = &SearchLogRuleGroupTemplesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件查询日志告警策略模板
func (c *Client) SearchLogRuleGroupTemples(request *SearchLogRuleGroupTemplesRequest) (response *SearchLogRuleGroupTemplesResponse, err error) {
	if request == nil {
		request = NewSearchLogRuleGroupTemplesRequest()
	}
	response = NewSearchLogRuleGroupTemplesResponse()
	err = c.Send(request, response)
	return
}

func NewExportLogAlertRuleRequest() (request *ExportLogAlertRuleRequest) {
	request = &ExportLogAlertRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ExportLogAlertRule")
	return
}

func NewExportLogAlertRuleResponse() (response *ExportLogAlertRuleResponse) {
	response = &ExportLogAlertRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出日志告警策略
func (c *Client) ExportLogAlertRule(request *ExportLogAlertRuleRequest) (response *ExportLogAlertRuleResponse, err error) {
	if request == nil {
		request = NewExportLogAlertRuleRequest()
	}
	response = NewExportLogAlertRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRouteRequest() (request *CreateRouteRequest) {
	request = &CreateRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "CreateRoute")
	return
}

func NewCreateRouteResponse() (response *CreateRouteResponse) {
	response = &CreateRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建告警路由
func (c *Client) CreateRoute(request *CreateRouteRequest) (response *CreateRouteResponse, err error) {
	if request == nil {
		request = NewCreateRouteRequest()
	}
	response = NewCreateRouteResponse()
	err = c.Send(request, response)
	return
}

func NewCopyLogDashboardRequest() (request *CopyLogDashboardRequest) {
	request = &CopyLogDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "CopyLogDashboard")
	return
}

func NewCopyLogDashboardResponse() (response *CopyLogDashboardResponse) {
	response = &CopyLogDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 复制日志dashboard接口
func (c *Client) CopyLogDashboard(request *CopyLogDashboardRequest) (response *CopyLogDashboardResponse, err error) {
	if request == nil {
		request = NewCopyLogDashboardRequest()
	}
	response = NewCopyLogDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewExportLogConfigCRRequest() (request *ExportLogConfigCRRequest) {
	request = &ExportLogConfigCRRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ExportLogConfigCR")
	return
}

func NewExportLogConfigCRResponse() (response *ExportLogConfigCRResponse) {
	response = &ExportLogConfigCRResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志平台导出配置CRD
func (c *Client) ExportLogConfigCR(request *ExportLogConfigCRRequest) (response *ExportLogConfigCRResponse, err error) {
	if request == nil {
		request = NewExportLogConfigCRRequest()
	}
	response = NewExportLogConfigCRResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogPanelRequest() (request *DescribeLogPanelRequest) {
	request = &DescribeLogPanelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DescribeLogPanel")
	return
}

func NewDescribeLogPanelResponse() (response *DescribeLogPanelResponse) {
	response = &DescribeLogPanelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看日志面板详情接口
func (c *Client) DescribeLogPanel(request *DescribeLogPanelRequest) (response *DescribeLogPanelResponse, err error) {
	if request == nil {
		request = NewDescribeLogPanelRequest()
	}
	response = NewDescribeLogPanelResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogDashboardsRequest() (request *DeleteLogDashboardsRequest) {
	request = &DeleteLogDashboardsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteLogDashboards")
	return
}

func NewDeleteLogDashboardsResponse() (response *DeleteLogDashboardsResponse) {
	response = &DeleteLogDashboardsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除日志仪表盘
func (c *Client) DeleteLogDashboards(request *DeleteLogDashboardsRequest) (response *DeleteLogDashboardsResponse, err error) {
	if request == nil {
		request = NewDeleteLogDashboardsRequest()
	}
	response = NewDeleteLogDashboardsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteInfraTemplatesRequest() (request *DeleteInfraTemplatesRequest) {
	request = &DeleteInfraTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteInfraTemplates")
	return
}

func NewDeleteInfraTemplatesResponse() (response *DeleteInfraTemplatesResponse) {
	response = &DeleteInfraTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云哨模板批量删除接口
func (c *Client) DeleteInfraTemplates(request *DeleteInfraTemplatesRequest) (response *DeleteInfraTemplatesResponse, err error) {
	if request == nil {
		request = NewDeleteInfraTemplatesRequest()
	}
	response = NewDeleteInfraTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewGetDashboardTagsRequest() (request *GetDashboardTagsRequest) {
	request = &GetDashboardTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetDashboardTags")
	return
}

func NewGetDashboardTagsResponse() (response *GetDashboardTagsResponse) {
	response = &GetDashboardTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取仪表盘Tag
func (c *Client) GetDashboardTags(request *GetDashboardTagsRequest) (response *GetDashboardTagsResponse, err error) {
	if request == nil {
		request = NewGetDashboardTagsRequest()
	}
	response = NewGetDashboardTagsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDashFolderByUidRequest() (request *DeleteDashFolderByUidRequest) {
	request = &DeleteDashFolderByUidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteDashFolderByUid")
	return
}

func NewDeleteDashFolderByUidResponse() (response *DeleteDashFolderByUidResponse) {
	response = &DeleteDashFolderByUidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除仪表盘目录
func (c *Client) DeleteDashFolderByUid(request *DeleteDashFolderByUidRequest) (response *DeleteDashFolderByUidResponse, err error) {
	if request == nil {
		request = NewDeleteDashFolderByUidRequest()
	}
	response = NewDeleteDashFolderByUidResponse()
	err = c.Send(request, response)
	return
}

func NewSearchRetentionsRequest() (request *SearchRetentionsRequest) {
	request = &SearchRetentionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "SearchRetentions")
	return
}

func NewSearchRetentionsResponse() (response *SearchRetentionsResponse) {
	response = &SearchRetentionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索Retention
func (c *Client) SearchRetentions(request *SearchRetentionsRequest) (response *SearchRetentionsResponse, err error) {
	if request == nil {
		request = NewSearchRetentionsRequest()
	}
	response = NewSearchRetentionsResponse()
	err = c.Send(request, response)
	return
}

func NewAddLogAlertRuleRequest() (request *AddLogAlertRuleRequest) {
	request = &AddLogAlertRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "AddLogAlertRule")
	return
}

func NewAddLogAlertRuleResponse() (response *AddLogAlertRuleResponse) {
	response = &AddLogAlertRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加日志报警
func (c *Client) AddLogAlertRule(request *AddLogAlertRuleRequest) (response *AddLogAlertRuleResponse, err error) {
	if request == nil {
		request = NewAddLogAlertRuleRequest()
	}
	response = NewAddLogAlertRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrUpdateRuleGroupTempleRequest() (request *CreateOrUpdateRuleGroupTempleRequest) {
	request = &CreateOrUpdateRuleGroupTempleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "CreateOrUpdateRuleGroupTemple")
	return
}

func NewCreateOrUpdateRuleGroupTempleResponse() (response *CreateOrUpdateRuleGroupTempleResponse) {
	response = &CreateOrUpdateRuleGroupTempleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建或更新策略模板
func (c *Client) CreateOrUpdateRuleGroupTemple(request *CreateOrUpdateRuleGroupTempleRequest) (response *CreateOrUpdateRuleGroupTempleResponse, err error) {
	if request == nil {
		request = NewCreateOrUpdateRuleGroupTempleRequest()
	}
	response = NewCreateOrUpdateRuleGroupTempleResponse()
	err = c.Send(request, response)
	return
}

func NewGenTimeFormatRequest() (request *GenTimeFormatRequest) {
	request = &GenTimeFormatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GenTimeFormat")
	return
}

func NewGenTimeFormatResponse() (response *GenTimeFormatResponse) {
	response = &GenTimeFormatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 自动生成时间格式
func (c *Client) GenTimeFormat(request *GenTimeFormatRequest) (response *GenTimeFormatResponse, err error) {
	if request == nil {
		request = NewGenTimeFormatRequest()
	}
	response = NewGenTimeFormatResponse()
	err = c.Send(request, response)
	return
}

func NewPreviewLogTransformTaskRequest() (request *PreviewLogTransformTaskRequest) {
	request = &PreviewLogTransformTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "PreviewLogTransformTask")
	return
}

func NewPreviewLogTransformTaskResponse() (response *PreviewLogTransformTaskResponse) {
	response = &PreviewLogTransformTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预览日志任务处理
func (c *Client) PreviewLogTransformTask(request *PreviewLogTransformTaskRequest) (response *PreviewLogTransformTaskResponse, err error) {
	if request == nil {
		request = NewPreviewLogTransformTaskRequest()
	}
	response = NewPreviewLogTransformTaskResponse()
	err = c.Send(request, response)
	return
}

func NewSearchLogDashboardsRequest() (request *SearchLogDashboardsRequest) {
	request = &SearchLogDashboardsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "SearchLogDashboards")
	return
}

func NewSearchLogDashboardsResponse() (response *SearchLogDashboardsResponse) {
	response = &SearchLogDashboardsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询日志仪表盘
func (c *Client) SearchLogDashboards(request *SearchLogDashboardsRequest) (response *SearchLogDashboardsResponse, err error) {
	if request == nil {
		request = NewSearchLogDashboardsRequest()
	}
	response = NewSearchLogDashboardsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateLogTransformTaskEnabledRequest() (request *UpdateLogTransformTaskEnabledRequest) {
	request = &UpdateLogTransformTaskEnabledRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "UpdateLogTransformTaskEnabled")
	return
}

func NewUpdateLogTransformTaskEnabledResponse() (response *UpdateLogTransformTaskEnabledResponse) {
	response = &UpdateLogTransformTaskEnabledResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新任务的开启状态
func (c *Client) UpdateLogTransformTaskEnabled(request *UpdateLogTransformTaskEnabledRequest) (response *UpdateLogTransformTaskEnabledResponse, err error) {
	if request == nil {
		request = NewUpdateLogTransformTaskEnabledRequest()
	}
	response = NewUpdateLogTransformTaskEnabledResponse()
	err = c.Send(request, response)
	return
}

func NewAddLogTransformTaskRequest() (request *AddLogTransformTaskRequest) {
	request = &AddLogTransformTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "AddLogTransformTask")
	return
}

func NewAddLogTransformTaskResponse() (response *AddLogTransformTaskResponse) {
	response = &AddLogTransformTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加日志任务处理
func (c *Client) AddLogTransformTask(request *AddLogTransformTaskRequest) (response *AddLogTransformTaskResponse, err error) {
	if request == nil {
		request = NewAddLogTransformTaskRequest()
	}
	response = NewAddLogTransformTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLogConfigRequest() (request *CreateLogConfigRequest) {
	request = &CreateLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "CreateLogConfig")
	return
}

func NewCreateLogConfigResponse() (response *CreateLogConfigResponse) {
	response = &CreateLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建日志接入配置
func (c *Client) CreateLogConfig(request *CreateLogConfigRequest) (response *CreateLogConfigResponse, err error) {
	if request == nil {
		request = NewCreateLogConfigRequest()
	}
	response = NewCreateLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewExportLogsRequest() (request *ExportLogsRequest) {
	request = &ExportLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ExportLogs")
	return
}

func NewExportLogsResponse() (response *ExportLogsResponse) {
	response = &ExportLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志导出
func (c *Client) ExportLogs(request *ExportLogsRequest) (response *ExportLogsResponse, err error) {
	if request == nil {
		request = NewExportLogsRequest()
	}
	response = NewExportLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogDeliveryTargetRequest() (request *DeleteLogDeliveryTargetRequest) {
	request = &DeleteLogDeliveryTargetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteLogDeliveryTarget")
	return
}

func NewDeleteLogDeliveryTargetResponse() (response *DeleteLogDeliveryTargetResponse) {
	response = &DeleteLogDeliveryTargetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除日志投递下游源
func (c *Client) DeleteLogDeliveryTarget(request *DeleteLogDeliveryTargetRequest) (response *DeleteLogDeliveryTargetResponse, err error) {
	if request == nil {
		request = NewDeleteLogDeliveryTargetRequest()
	}
	response = NewDeleteLogDeliveryTargetResponse()
	err = c.Send(request, response)
	return
}

func NewGetAllUnitRequest() (request *GetAllUnitRequest) {
	request = &GetAllUnitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetAllUnit")
	return
}

func NewGetAllUnitResponse() (response *GetAllUnitResponse) {
	response = &GetAllUnitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有指标单位
func (c *Client) GetAllUnit(request *GetAllUnitRequest) (response *GetAllUnitResponse, err error) {
	if request == nil {
		request = NewGetAllUnitRequest()
	}
	response = NewGetAllUnitResponse()
	err = c.Send(request, response)
	return
}

func NewConfirmAlertRequest() (request *ConfirmAlertRequest) {
	request = &ConfirmAlertRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ConfirmAlert")
	return
}

func NewConfirmAlertResponse() (response *ConfirmAlertResponse) {
	response = &ConfirmAlertResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 确认指定Alert
func (c *Client) ConfirmAlert(request *ConfirmAlertRequest) (response *ConfirmAlertResponse, err error) {
	if request == nil {
		request = NewConfirmAlertRequest()
	}
	response = NewConfirmAlertResponse()
	err = c.Send(request, response)
	return
}

func NewSearchNotificationAlertsRequest() (request *SearchNotificationAlertsRequest) {
	request = &SearchNotificationAlertsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "SearchNotificationAlerts")
	return
}

func NewSearchNotificationAlertsResponse() (response *SearchNotificationAlertsResponse) {
	response = &SearchNotificationAlertsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警消息内查询alert
func (c *Client) SearchNotificationAlerts(request *SearchNotificationAlertsRequest) (response *SearchNotificationAlertsResponse, err error) {
	if request == nil {
		request = NewSearchNotificationAlertsRequest()
	}
	response = NewSearchNotificationAlertsResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceTypesRequest() (request *ListResourceTypesRequest) {
	request = &ListResourceTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ListResourceTypes")
	return
}

func NewListResourceTypesResponse() (response *ListResourceTypesResponse) {
	response = &ListResourceTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源对象列表
func (c *Client) ListResourceTypes(request *ListResourceTypesRequest) (response *ListResourceTypesResponse, err error) {
	if request == nil {
		request = NewListResourceTypesRequest()
	}
	response = NewListResourceTypesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDashFolderRequest() (request *CreateDashFolderRequest) {
	request = &CreateDashFolderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "CreateDashFolder")
	return
}

func NewCreateDashFolderResponse() (response *CreateDashFolderResponse) {
	response = &CreateDashFolderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建仪表盘目录
func (c *Client) CreateDashFolder(request *CreateDashFolderRequest) (response *CreateDashFolderResponse, err error) {
	if request == nil {
		request = NewCreateDashFolderRequest()
	}
	response = NewCreateDashFolderResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateLogDeliveryTaskRequest() (request *UpdateLogDeliveryTaskRequest) {
	request = &UpdateLogDeliveryTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "UpdateLogDeliveryTask")
	return
}

func NewUpdateLogDeliveryTaskResponse() (response *UpdateLogDeliveryTaskResponse) {
	response = &UpdateLogDeliveryTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新日志投递任务
func (c *Client) UpdateLogDeliveryTask(request *UpdateLogDeliveryTaskRequest) (response *UpdateLogDeliveryTaskResponse, err error) {
	if request == nil {
		request = NewUpdateLogDeliveryTaskRequest()
	}
	response = NewUpdateLogDeliveryTaskResponse()
	err = c.Send(request, response)
	return
}

func NewGetAlertRequest() (request *GetAlertRequest) {
	request = &GetAlertRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetAlert")
	return
}

func NewGetAlertResponse() (response *GetAlertResponse) {
	response = &GetAlertResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定Alert
func (c *Client) GetAlert(request *GetAlertRequest) (response *GetAlertResponse, err error) {
	if request == nil {
		request = NewGetAlertRequest()
	}
	response = NewGetAlertResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogTransformTaskRequest() (request *DeleteLogTransformTaskRequest) {
	request = &DeleteLogTransformTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteLogTransformTask")
	return
}

func NewDeleteLogTransformTaskResponse() (response *DeleteLogTransformTaskResponse) {
	response = &DeleteLogTransformTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除日志任务处理
func (c *Client) DeleteLogTransformTask(request *DeleteLogTransformTaskRequest) (response *DeleteLogTransformTaskResponse, err error) {
	if request == nil {
		request = NewDeleteLogTransformTaskRequest()
	}
	response = NewDeleteLogTransformTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRoutesRequest() (request *DeleteRoutesRequest) {
	request = &DeleteRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteRoutes")
	return
}

func NewDeleteRoutesResponse() (response *DeleteRoutesResponse) {
	response = &DeleteRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除指定路由信息
func (c *Client) DeleteRoutes(request *DeleteRoutesRequest) (response *DeleteRoutesResponse, err error) {
	if request == nil {
		request = NewDeleteRoutesRequest()
	}
	response = NewDeleteRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRuleGroupsRequest() (request *DeleteRuleGroupsRequest) {
	request = &DeleteRuleGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteRuleGroups")
	return
}

func NewDeleteRuleGroupsResponse() (response *DeleteRuleGroupsResponse) {
	response = &DeleteRuleGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除策略
func (c *Client) DeleteRuleGroups(request *DeleteRuleGroupsRequest) (response *DeleteRuleGroupsResponse, err error) {
	if request == nil {
		request = NewDeleteRuleGroupsRequest()
	}
	response = NewDeleteRuleGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogDashboardRequest() (request *DescribeLogDashboardRequest) {
	request = &DescribeLogDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DescribeLogDashboard")
	return
}

func NewDescribeLogDashboardResponse() (response *DescribeLogDashboardResponse) {
	response = &DescribeLogDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看日志仪表盘详情
func (c *Client) DescribeLogDashboard(request *DescribeLogDashboardRequest) (response *DescribeLogDashboardResponse, err error) {
	if request == nil {
		request = NewDescribeLogDashboardRequest()
	}
	response = NewDescribeLogDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogConfigsRequest() (request *DescribeLogConfigsRequest) {
	request = &DescribeLogConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DescribeLogConfigs")
	return
}

func NewDescribeLogConfigsResponse() (response *DescribeLogConfigsResponse) {
	response = &DescribeLogConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询接入配置
func (c *Client) DescribeLogConfigs(request *DescribeLogConfigsRequest) (response *DescribeLogConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeLogConfigsRequest()
	}
	response = NewDescribeLogConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewSearchSilencesRequest() (request *SearchSilencesRequest) {
	request = &SearchSilencesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "SearchSilences")
	return
}

func NewSearchSilencesResponse() (response *SearchSilencesResponse) {
	response = &SearchSilencesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件查询屏蔽策略
func (c *Client) SearchSilences(request *SearchSilencesRequest) (response *SearchSilencesResponse, err error) {
	if request == nil {
		request = NewSearchSilencesRequest()
	}
	response = NewSearchSilencesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateDashFolderByUidRequest() (request *UpdateDashFolderByUidRequest) {
	request = &UpdateDashFolderByUidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "UpdateDashFolderByUid")
	return
}

func NewUpdateDashFolderByUidResponse() (response *UpdateDashFolderByUidResponse) {
	response = &UpdateDashFolderByUidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新仪表盘目录
func (c *Client) UpdateDashFolderByUid(request *UpdateDashFolderByUidRequest) (response *UpdateDashFolderByUidResponse, err error) {
	if request == nil {
		request = NewUpdateDashFolderByUidRequest()
	}
	response = NewUpdateDashFolderByUidResponse()
	err = c.Send(request, response)
	return
}

func NewListInfraTemplatesRequest() (request *ListInfraTemplatesRequest) {
	request = &ListInfraTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ListInfraTemplates")
	return
}

func NewListInfraTemplatesResponse() (response *ListInfraTemplatesResponse) {
	response = &ListInfraTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云哨模板查询接口
func (c *Client) ListInfraTemplates(request *ListInfraTemplatesRequest) (response *ListInfraTemplatesResponse, err error) {
	if request == nil {
		request = NewListInfraTemplatesRequest()
	}
	response = NewListInfraTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQueryRecordsRequest() (request *DescribeQueryRecordsRequest) {
	request = &DescribeQueryRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DescribeQueryRecords")
	return
}

func NewDescribeQueryRecordsResponse() (response *DescribeQueryRecordsResponse) {
	response = &DescribeQueryRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询搜索记录
func (c *Client) DescribeQueryRecords(request *DescribeQueryRecordsRequest) (response *DescribeQueryRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeQueryRecordsRequest()
	}
	response = NewDescribeQueryRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogDeliveryTaskRequest() (request *DeleteLogDeliveryTaskRequest) {
	request = &DeleteLogDeliveryTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "DeleteLogDeliveryTask")
	return
}

func NewDeleteLogDeliveryTaskResponse() (response *DeleteLogDeliveryTaskResponse) {
	response = &DeleteLogDeliveryTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除日志投递任务
func (c *Client) DeleteLogDeliveryTask(request *DeleteLogDeliveryTaskRequest) (response *DeleteLogDeliveryTaskResponse, err error) {
	if request == nil {
		request = NewDeleteLogDeliveryTaskRequest()
	}
	response = NewDeleteLogDeliveryTaskResponse()
	err = c.Send(request, response)
	return
}

func NewGetDashFoldersRequest() (request *GetDashFoldersRequest) {
	request = &GetDashFoldersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetDashFolders")
	return
}

func NewGetDashFoldersResponse() (response *GetDashFoldersResponse) {
	response = &GetDashFoldersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过名称、ID列表等内容检索仪表盘目录
func (c *Client) GetDashFolders(request *GetDashFoldersRequest) (response *GetDashFoldersResponse, err error) {
	if request == nil {
		request = NewGetDashFoldersRequest()
	}
	response = NewGetDashFoldersResponse()
	err = c.Send(request, response)
	return
}

func NewGetDataBaradMetricRequest() (request *GetDataBaradMetricRequest) {
	request = &GetDataBaradMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetDataBaradMetric")
	return
}

func NewGetDataBaradMetricResponse() (response *GetDataBaradMetricResponse) {
	response = &GetDataBaradMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取Barad指标数据接口，为了兼容Barad接口格式而创建。
func (c *Client) GetDataBaradMetric(request *GetDataBaradMetricRequest) (response *GetDataBaradMetricResponse, err error) {
	if request == nil {
		request = NewGetDataBaradMetricRequest()
	}
	response = NewGetDataBaradMetricResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrUpdateRetentionRequest() (request *CreateOrUpdateRetentionRequest) {
	request = &CreateOrUpdateRetentionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "CreateOrUpdateRetention")
	return
}

func NewCreateOrUpdateRetentionResponse() (response *CreateOrUpdateRetentionResponse) {
	response = &CreateOrUpdateRetentionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建或更新Retention
func (c *Client) CreateOrUpdateRetention(request *CreateOrUpdateRetentionRequest) (response *CreateOrUpdateRetentionResponse, err error) {
	if request == nil {
		request = NewCreateOrUpdateRetentionRequest()
	}
	response = NewCreateOrUpdateRetentionResponse()
	err = c.Send(request, response)
	return
}

func NewListLogAlertRuleRequest() (request *ListLogAlertRuleRequest) {
	request = &ListLogAlertRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ListLogAlertRule")
	return
}

func NewListLogAlertRuleResponse() (response *ListLogAlertRuleResponse) {
	response = &ListLogAlertRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 展示日志报警
func (c *Client) ListLogAlertRule(request *ListLogAlertRuleRequest) (response *ListLogAlertRuleResponse, err error) {
	if request == nil {
		request = NewListLogAlertRuleRequest()
	}
	response = NewListLogAlertRuleResponse()
	err = c.Send(request, response)
	return
}

func NewAddLogDeliveryTargetRequest() (request *AddLogDeliveryTargetRequest) {
	request = &AddLogDeliveryTargetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "AddLogDeliveryTarget")
	return
}

func NewAddLogDeliveryTargetResponse() (response *AddLogDeliveryTargetResponse) {
	response = &AddLogDeliveryTargetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增日志投递下游源
func (c *Client) AddLogDeliveryTarget(request *AddLogDeliveryTargetRequest) (response *AddLogDeliveryTargetResponse, err error) {
	if request == nil {
		request = NewAddLogDeliveryTargetRequest()
	}
	response = NewAddLogDeliveryTargetResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMetaMetricRequest() (request *UpdateMetaMetricRequest) {
	request = &UpdateMetaMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "UpdateMetaMetric")
	return
}

func NewUpdateMetaMetricResponse() (response *UpdateMetaMetricResponse) {
	response = &UpdateMetaMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新指标信息
func (c *Client) UpdateMetaMetric(request *UpdateMetaMetricRequest) (response *UpdateMetaMetricResponse, err error) {
	if request == nil {
		request = NewUpdateMetaMetricRequest()
	}
	response = NewUpdateMetaMetricResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrUpdateLogDashboardRequest() (request *CreateOrUpdateLogDashboardRequest) {
	request = &CreateOrUpdateLogDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "CreateOrUpdateLogDashboard")
	return
}

func NewCreateOrUpdateLogDashboardResponse() (response *CreateOrUpdateLogDashboardResponse) {
	response = &CreateOrUpdateLogDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建日志dashboard接口
func (c *Client) CreateOrUpdateLogDashboard(request *CreateOrUpdateLogDashboardRequest) (response *CreateOrUpdateLogDashboardResponse, err error) {
	if request == nil {
		request = NewCreateOrUpdateLogDashboardRequest()
	}
	response = NewCreateOrUpdateLogDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewGetMetaRequest() (request *GetMetaRequest) {
	request = &GetMetaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetMeta")
	return
}

func NewGetMetaResponse() (response *GetMetaResponse) {
	response = &GetMetaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Meta树
func (c *Client) GetMeta(request *GetMetaRequest) (response *GetMetaResponse, err error) {
	if request == nil {
		request = NewGetMetaRequest()
	}
	response = NewGetMetaResponse()
	err = c.Send(request, response)
	return
}

func NewAddLogDeliveryTaskRequest() (request *AddLogDeliveryTaskRequest) {
	request = &AddLogDeliveryTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "AddLogDeliveryTask")
	return
}

func NewAddLogDeliveryTaskResponse() (response *AddLogDeliveryTaskResponse) {
	response = &AddLogDeliveryTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增日志投递任务
func (c *Client) AddLogDeliveryTask(request *AddLogDeliveryTaskRequest) (response *AddLogDeliveryTaskResponse, err error) {
	if request == nil {
		request = NewAddLogDeliveryTaskRequest()
	}
	response = NewAddLogDeliveryTaskResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateLogDeliveryTargetRequest() (request *UpdateLogDeliveryTargetRequest) {
	request = &UpdateLogDeliveryTargetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "UpdateLogDeliveryTarget")
	return
}

func NewUpdateLogDeliveryTargetResponse() (response *UpdateLogDeliveryTargetResponse) {
	response = &UpdateLogDeliveryTargetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新日志投递下游源
func (c *Client) UpdateLogDeliveryTarget(request *UpdateLogDeliveryTargetRequest) (response *UpdateLogDeliveryTargetResponse, err error) {
	if request == nil {
		request = NewUpdateLogDeliveryTargetRequest()
	}
	response = NewUpdateLogDeliveryTargetResponse()
	err = c.Send(request, response)
	return
}

func NewAddResourceObjectRequest() (request *AddResourceObjectRequest) {
	request = &AddResourceObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "AddResourceObject")
	return
}

func NewAddResourceObjectResponse() (response *AddResourceObjectResponse) {
	response = &AddResourceObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建服务树资源对象
func (c *Client) AddResourceObject(request *AddResourceObjectRequest) (response *AddResourceObjectResponse, err error) {
	if request == nil {
		request = NewAddResourceObjectRequest()
	}
	response = NewAddResourceObjectResponse()
	err = c.Send(request, response)
	return
}

func NewGetRuleGroupTempleRequest() (request *GetRuleGroupTempleRequest) {
	request = &GetRuleGroupTempleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetRuleGroupTemple")
	return
}

func NewGetRuleGroupTempleResponse() (response *GetRuleGroupTempleResponse) {
	response = &GetRuleGroupTempleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定策略模板
func (c *Client) GetRuleGroupTemple(request *GetRuleGroupTempleRequest) (response *GetRuleGroupTempleResponse, err error) {
	if request == nil {
		request = NewGetRuleGroupTempleRequest()
	}
	response = NewGetRuleGroupTempleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogConfigRequest() (request *ModifyLogConfigRequest) {
	request = &ModifyLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ModifyLogConfig")
	return
}

func NewModifyLogConfigResponse() (response *ModifyLogConfigResponse) {
	response = &ModifyLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 日志接入配置变更
func (c *Client) ModifyLogConfig(request *ModifyLogConfigRequest) (response *ModifyLogConfigResponse, err error) {
	if request == nil {
		request = NewModifyLogConfigRequest()
	}
	response = NewModifyLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewListResourceTypeEventRequest() (request *ListResourceTypeEventRequest) {
	request = &ListResourceTypeEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ListResourceTypeEvent")
	return
}

func NewListResourceTypeEventResponse() (response *ListResourceTypeEventResponse) {
	response = &ListResourceTypeEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源对象元数据-对象类型事件列表
func (c *Client) ListResourceTypeEvent(request *ListResourceTypeEventRequest) (response *ListResourceTypeEventResponse, err error) {
	if request == nil {
		request = NewListResourceTypeEventRequest()
	}
	response = NewListResourceTypeEventResponse()
	err = c.Send(request, response)
	return
}

func NewSearchRuleGroupsRequest() (request *SearchRuleGroupsRequest) {
	request = &SearchRuleGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "SearchRuleGroups")
	return
}

func NewSearchRuleGroupsResponse() (response *SearchRuleGroupsResponse) {
	response = &SearchRuleGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件查询策略
func (c *Client) SearchRuleGroups(request *SearchRuleGroupsRequest) (response *SearchRuleGroupsResponse, err error) {
	if request == nil {
		request = NewSearchRuleGroupsRequest()
	}
	response = NewSearchRuleGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrUpdateLogPanelRequest() (request *CreateOrUpdateLogPanelRequest) {
	request = &CreateOrUpdateLogPanelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "CreateOrUpdateLogPanel")
	return
}

func NewCreateOrUpdateLogPanelResponse() (response *CreateOrUpdateLogPanelResponse) {
	response = &CreateOrUpdateLogPanelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建或更新日志仪表盘面板
func (c *Client) CreateOrUpdateLogPanel(request *CreateOrUpdateLogPanelRequest) (response *CreateOrUpdateLogPanelResponse, err error) {
	if request == nil {
		request = NewCreateOrUpdateLogPanelRequest()
	}
	response = NewCreateOrUpdateLogPanelResponse()
	err = c.Send(request, response)
	return
}

func NewGenLineHeaderRegexRequest() (request *GenLineHeaderRegexRequest) {
	request = &GenLineHeaderRegexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GenLineHeaderRegex")
	return
}

func NewGenLineHeaderRegexResponse() (response *GenLineHeaderRegexResponse) {
	response = &GenLineHeaderRegexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 自动生成行首正则
func (c *Client) GenLineHeaderRegex(request *GenLineHeaderRegexRequest) (response *GenLineHeaderRegexResponse, err error) {
	if request == nil {
		request = NewGenLineHeaderRegexRequest()
	}
	response = NewGenLineHeaderRegexResponse()
	err = c.Send(request, response)
	return
}

func NewListMetaMetricRequest() (request *ListMetaMetricRequest) {
	request = &ListMetaMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "ListMetaMetric")
	return
}

func NewListMetaMetricResponse() (response *ListMetaMetricResponse) {
	response = &ListMetaMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源对象类型下的指标列表
func (c *Client) ListMetaMetric(request *ListMetaMetricRequest) (response *ListMetaMetricResponse, err error) {
	if request == nil {
		request = NewListMetaMetricRequest()
	}
	response = NewListMetaMetricResponse()
	err = c.Send(request, response)
	return
}

func NewSearchDashboardsRequest() (request *SearchDashboardsRequest) {
	request = &SearchDashboardsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "SearchDashboards")
	return
}

func NewSearchDashboardsResponse() (response *SearchDashboardsResponse) {
	response = &SearchDashboardsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询仪表盘
func (c *Client) SearchDashboards(request *SearchDashboardsRequest) (response *SearchDashboardsResponse, err error) {
	if request == nil {
		request = NewSearchDashboardsRequest()
	}
	response = NewSearchDashboardsResponse()
	err = c.Send(request, response)
	return
}

func NewGetDefaultGroupByRequest() (request *GetDefaultGroupByRequest) {
	request = &GetDefaultGroupByRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetDefaultGroupBy")
	return
}

func NewGetDefaultGroupByResponse() (response *GetDefaultGroupByResponse) {
	response = &GetDefaultGroupByResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取聚合维度
func (c *Client) GetDefaultGroupBy(request *GetDefaultGroupByRequest) (response *GetDefaultGroupByResponse, err error) {
	if request == nil {
		request = NewGetDefaultGroupByRequest()
	}
	response = NewGetDefaultGroupByResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrUpdateRuleGroupRequest() (request *CreateOrUpdateRuleGroupRequest) {
	request = &CreateOrUpdateRuleGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "CreateOrUpdateRuleGroup")
	return
}

func NewCreateOrUpdateRuleGroupResponse() (response *CreateOrUpdateRuleGroupResponse) {
	response = &CreateOrUpdateRuleGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建或更新策略
func (c *Client) CreateOrUpdateRuleGroup(request *CreateOrUpdateRuleGroupRequest) (response *CreateOrUpdateRuleGroupResponse, err error) {
	if request == nil {
		request = NewCreateOrUpdateRuleGroupRequest()
	}
	response = NewCreateOrUpdateRuleGroupResponse()
	err = c.Send(request, response)
	return
}

func NewGetDashboardContentByUidRequest() (request *GetDashboardContentByUidRequest) {
	request = &GetDashboardContentByUidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetDashboardContentByUid")
	return
}

func NewGetDashboardContentByUidResponse() (response *GetDashboardContentByUidResponse) {
	response = &GetDashboardContentByUidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取仪表盘详细内容
func (c *Client) GetDashboardContentByUid(request *GetDashboardContentByUidRequest) (response *GetDashboardContentByUidResponse, err error) {
	if request == nil {
		request = NewGetDashboardContentByUidRequest()
	}
	response = NewGetDashboardContentByUidResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateRouteRequest() (request *UpdateRouteRequest) {
	request = &UpdateRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "UpdateRoute")
	return
}

func NewUpdateRouteResponse() (response *UpdateRouteResponse) {
	response = &UpdateRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新指定路由信息
func (c *Client) UpdateRoute(request *UpdateRouteRequest) (response *UpdateRouteResponse, err error) {
	if request == nil {
		request = NewUpdateRouteRequest()
	}
	response = NewUpdateRouteResponse()
	err = c.Send(request, response)
	return
}

func NewGetRuleGroupRequest() (request *GetRuleGroupRequest) {
	request = &GetRuleGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastore", APIVersion, "GetRuleGroup")
	return
}

func NewGetRuleGroupResponse() (response *GetRuleGroupResponse) {
	response = &GetRuleGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定策略
func (c *Client) GetRuleGroup(request *GetRuleGroupRequest) (response *GetRuleGroupResponse, err error) {
	if request == nil {
		request = NewGetRuleGroupRequest()
	}
	response = NewGetRuleGroupResponse()
	err = c.Send(request, response)
	return
}
