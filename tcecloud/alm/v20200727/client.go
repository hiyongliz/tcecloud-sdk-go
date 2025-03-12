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

package v20200727

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-07-27"

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

func NewGetAppInstanceComponentDetailRequest() (request *GetAppInstanceComponentDetailRequest) {
	request = &GetAppInstanceComponentDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetAppInstanceComponentDetail")
	return
}

func NewGetAppInstanceComponentDetailResponse() (response *GetAppInstanceComponentDetailResponse) {
	response = &GetAppInstanceComponentDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取helm应用组件yaml详情
func (c *Client) GetAppInstanceComponentDetail(request *GetAppInstanceComponentDetailRequest) (response *GetAppInstanceComponentDetailResponse, err error) {
	if request == nil {
		request = NewGetAppInstanceComponentDetailRequest()
	}
	response = NewGetAppInstanceComponentDetailResponse()
	err = c.Send(request, response)
	return
}

func NewListBriefAppsRequest() (request *ListBriefAppsRequest) {
	request = &ListBriefAppsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListBriefApps")
	return
}

func NewListBriefAppsResponse() (response *ListBriefAppsResponse) {
	response = &ListBriefAppsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用市场App
func (c *Client) ListBriefApps(request *ListBriefAppsRequest) (response *ListBriefAppsResponse, err error) {
	if request == nil {
		request = NewListBriefAppsRequest()
	}
	response = NewListBriefAppsResponse()
	err = c.Send(request, response)
	return
}

func NewSaveConfigMapYamlRequest() (request *SaveConfigMapYamlRequest) {
	request = &SaveConfigMapYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "SaveConfigMapYaml")
	return
}

func NewSaveConfigMapYamlResponse() (response *SaveConfigMapYamlResponse) {
	response = &SaveConfigMapYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存ConfigMap Yaml
func (c *Client) SaveConfigMapYaml(request *SaveConfigMapYamlRequest) (response *SaveConfigMapYamlResponse, err error) {
	if request == nil {
		request = NewSaveConfigMapYamlRequest()
	}
	response = NewSaveConfigMapYamlResponse()
	err = c.Send(request, response)
	return
}

func NewGetCanaryPartitionRequest() (request *GetCanaryPartitionRequest) {
	request = &GetCanaryPartitionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetCanaryPartition")
	return
}

func NewGetCanaryPartitionResponse() (response *GetCanaryPartitionResponse) {
	response = &GetCanaryPartitionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取灰度发布保留旧版本pod的副本数
func (c *Client) GetCanaryPartition(request *GetCanaryPartitionRequest) (response *GetCanaryPartitionResponse, err error) {
	if request == nil {
		request = NewGetCanaryPartitionRequest()
	}
	response = NewGetCanaryPartitionResponse()
	err = c.Send(request, response)
	return
}

func NewSetCanaryPartitionRequest() (request *SetCanaryPartitionRequest) {
	request = &SetCanaryPartitionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "SetCanaryPartition")
	return
}

func NewSetCanaryPartitionResponse() (response *SetCanaryPartitionResponse) {
	response = &SetCanaryPartitionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置灰度发布保留旧版本的pod数量
func (c *Client) SetCanaryPartition(request *SetCanaryPartitionRequest) (response *SetCanaryPartitionResponse, err error) {
	if request == nil {
		request = NewSetCanaryPartitionRequest()
	}
	response = NewSetCanaryPartitionResponse()
	err = c.Send(request, response)
	return
}

func NewListWorkloadPodsRequest() (request *ListWorkloadPodsRequest) {
	request = &ListWorkloadPodsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListWorkloadPods")
	return
}

func NewListWorkloadPodsResponse() (response *ListWorkloadPodsResponse) {
	response = &ListWorkloadPodsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取helm应用workload的pod列表
func (c *Client) ListWorkloadPods(request *ListWorkloadPodsRequest) (response *ListWorkloadPodsResponse, err error) {
	if request == nil {
		request = NewListWorkloadPodsRequest()
	}
	response = NewListWorkloadPodsResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchScalerStatusRequest() (request *SwitchScalerStatusRequest) {
	request = &SwitchScalerStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "SwitchScalerStatus")
	return
}

func NewSwitchScalerStatusResponse() (response *SwitchScalerStatusResponse) {
	response = &SwitchScalerStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换策略状态，使策略开启或关闭
func (c *Client) SwitchScalerStatus(request *SwitchScalerStatusRequest) (response *SwitchScalerStatusResponse, err error) {
	if request == nil {
		request = NewSwitchScalerStatusRequest()
	}
	response = NewSwitchScalerStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAppPreCheckToYamlRequest() (request *CreateAppPreCheckToYamlRequest) {
	request = &CreateAppPreCheckToYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CreateAppPreCheckToYaml")
	return
}

func NewCreateAppPreCheckToYamlResponse() (response *CreateAppPreCheckToYamlResponse) {
	response = &CreateAppPreCheckToYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateAppPreCheckToYaml
func (c *Client) CreateAppPreCheckToYaml(request *CreateAppPreCheckToYamlRequest) (response *CreateAppPreCheckToYamlResponse, err error) {
	if request == nil {
		request = NewCreateAppPreCheckToYamlRequest()
	}
	response = NewCreateAppPreCheckToYamlResponse()
	err = c.Send(request, response)
	return
}

func NewGetAppInstanceIngressRequest() (request *GetAppInstanceIngressRequest) {
	request = &GetAppInstanceIngressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetAppInstanceIngress")
	return
}

func NewGetAppInstanceIngressResponse() (response *GetAppInstanceIngressResponse) {
	response = &GetAppInstanceIngressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用内的访问入口
func (c *Client) GetAppInstanceIngress(request *GetAppInstanceIngressRequest) (response *GetAppInstanceIngressResponse, err error) {
	if request == nil {
		request = NewGetAppInstanceIngressRequest()
	}
	response = NewGetAppInstanceIngressResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAppFromTadPkgRequest() (request *CreateAppFromTadPkgRequest) {
	request = &CreateAppFromTadPkgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CreateAppFromTadPkg")
	return
}

func NewCreateAppFromTadPkgResponse() (response *CreateAppFromTadPkgResponse) {
	response = &CreateAppFromTadPkgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 使用TAD应用包创建应用模版
func (c *Client) CreateAppFromTadPkg(request *CreateAppFromTadPkgRequest) (response *CreateAppFromTadPkgResponse, err error) {
	if request == nil {
		request = NewCreateAppFromTadPkgRequest()
	}
	response = NewCreateAppFromTadPkgResponse()
	err = c.Send(request, response)
	return
}

func NewGetNewestRevisionRequest() (request *GetNewestRevisionRequest) {
	request = &GetNewestRevisionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetNewestRevision")
	return
}

func NewGetNewestRevisionResponse() (response *GetNewestRevisionResponse) {
	response = &GetNewestRevisionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询下个build版本号
func (c *Client) GetNewestRevision(request *GetNewestRevisionRequest) (response *GetNewestRevisionResponse, err error) {
	if request == nil {
		request = NewGetNewestRevisionRequest()
	}
	response = NewGetNewestRevisionResponse()
	err = c.Send(request, response)
	return
}

func NewPromoteRolloutRequest() (request *PromoteRolloutRequest) {
	request = &PromoteRolloutRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "PromoteRollout")
	return
}

func NewPromoteRolloutResponse() (response *PromoteRolloutResponse) {
	response = &PromoteRolloutResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 完成当前更新
func (c *Client) PromoteRollout(request *PromoteRolloutRequest) (response *PromoteRolloutResponse, err error) {
	if request == nil {
		request = NewPromoteRolloutRequest()
	}
	response = NewPromoteRolloutResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAppInstanceIngressRequest() (request *CreateAppInstanceIngressRequest) {
	request = &CreateAppInstanceIngressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CreateAppInstanceIngress")
	return
}

func NewCreateAppInstanceIngressResponse() (response *CreateAppInstanceIngressResponse) {
	response = &CreateAppInstanceIngressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加应用内的访问入口信息
func (c *Client) CreateAppInstanceIngress(request *CreateAppInstanceIngressRequest) (response *CreateAppInstanceIngressResponse, err error) {
	if request == nil {
		request = NewCreateAppInstanceIngressRequest()
	}
	response = NewCreateAppInstanceIngressResponse()
	err = c.Send(request, response)
	return
}

func NewGetConfigMapYamlRequest() (request *GetConfigMapYamlRequest) {
	request = &GetConfigMapYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetConfigMapYaml")
	return
}

func NewGetConfigMapYamlResponse() (response *GetConfigMapYamlResponse) {
	response = &GetConfigMapYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询ConfigMap Yaml
func (c *Client) GetConfigMapYaml(request *GetConfigMapYamlRequest) (response *GetConfigMapYamlResponse, err error) {
	if request == nil {
		request = NewGetConfigMapYamlRequest()
	}
	response = NewGetConfigMapYamlResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConfigMapsRequest() (request *DeleteConfigMapsRequest) {
	request = &DeleteConfigMapsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "DeleteConfigMaps")
	return
}

func NewDeleteConfigMapsResponse() (response *DeleteConfigMapsResponse) {
	response = &DeleteConfigMapsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除ConfigMap
func (c *Client) DeleteConfigMaps(request *DeleteConfigMapsRequest) (response *DeleteConfigMapsResponse, err error) {
	if request == nil {
		request = NewDeleteConfigMapsRequest()
	}
	response = NewDeleteConfigMapsResponse()
	err = c.Send(request, response)
	return
}

func NewGetPaaSSummaryRequest() (request *GetPaaSSummaryRequest) {
	request = &GetPaaSSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetPaaSSummary")
	return
}

func NewGetPaaSSummaryResponse() (response *GetPaaSSummaryResponse) {
	response = &GetPaaSSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) GetPaaSSummary(request *GetPaaSSummaryRequest) (response *GetPaaSSummaryResponse, err error) {
	if request == nil {
		request = NewGetPaaSSummaryRequest()
	}
	response = NewGetPaaSSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewListAppInsProductsRequest() (request *ListAppInsProductsRequest) {
	request = &ListAppInsProductsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListAppInsProducts")
	return
}

func NewListAppInsProductsResponse() (response *ListAppInsProductsResponse) {
	response = &ListAppInsProductsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回产品列表
func (c *Client) ListAppInsProducts(request *ListAppInsProductsRequest) (response *ListAppInsProductsResponse, err error) {
	if request == nil {
		request = NewListAppInsProductsRequest()
	}
	response = NewListAppInsProductsResponse()
	err = c.Send(request, response)
	return
}

func NewListRevisionsRequest() (request *ListRevisionsRequest) {
	request = &ListRevisionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListRevisions")
	return
}

func NewListRevisionsResponse() (response *ListRevisionsResponse) {
	response = &ListRevisionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用的revision版本列表
func (c *Client) ListRevisions(request *ListRevisionsRequest) (response *ListRevisionsResponse, err error) {
	if request == nil {
		request = NewListRevisionsRequest()
	}
	response = NewListRevisionsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAppInstanceIngressRequest() (request *DeleteAppInstanceIngressRequest) {
	request = &DeleteAppInstanceIngressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "DeleteAppInstanceIngress")
	return
}

func NewDeleteAppInstanceIngressResponse() (response *DeleteAppInstanceIngressResponse) {
	response = &DeleteAppInstanceIngressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除应用内的访问入口信息
func (c *Client) DeleteAppInstanceIngress(request *DeleteAppInstanceIngressRequest) (response *DeleteAppInstanceIngressResponse, err error) {
	if request == nil {
		request = NewDeleteAppInstanceIngressRequest()
	}
	response = NewDeleteAppInstanceIngressResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationEventsRequest() (request *GetApplicationEventsRequest) {
	request = &GetApplicationEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetApplicationEvents")
	return
}

func NewGetApplicationEventsResponse() (response *GetApplicationEventsResponse) {
	response = &GetApplicationEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 应用事件查询接口
func (c *Client) GetApplicationEvents(request *GetApplicationEventsRequest) (response *GetApplicationEventsResponse, err error) {
	if request == nil {
		request = NewGetApplicationEventsRequest()
	}
	response = NewGetApplicationEventsResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductTgzRequest() (request *GetProductTgzRequest) {
	request = &GetProductTgzRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetProductTgz")
	return
}

func NewGetProductTgzResponse() (response *GetProductTgzResponse) {
	response = &GetProductTgzResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetProductTgz
func (c *Client) GetProductTgz(request *GetProductTgzRequest) (response *GetProductTgzResponse, err error) {
	if request == nil {
		request = NewGetProductTgzRequest()
	}
	response = NewGetProductTgzResponse()
	err = c.Send(request, response)
	return
}

func NewListScalersRequest() (request *ListScalersRequest) {
	request = &ListScalersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListScalers")
	return
}

func NewListScalersResponse() (response *ListScalersResponse) {
	response = &ListScalersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取弹性伸缩策略列表
func (c *Client) ListScalers(request *ListScalersRequest) (response *ListScalersResponse, err error) {
	if request == nil {
		request = NewListScalersRequest()
	}
	response = NewListScalersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAppInstanceFromYamlRequest() (request *CreateAppInstanceFromYamlRequest) {
	request = &CreateAppInstanceFromYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CreateAppInstanceFromYaml")
	return
}

func NewCreateAppInstanceFromYamlResponse() (response *CreateAppInstanceFromYamlResponse) {
	response = &CreateAppInstanceFromYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过yaml部署应用实例
func (c *Client) CreateAppInstanceFromYaml(request *CreateAppInstanceFromYamlRequest) (response *CreateAppInstanceFromYamlResponse, err error) {
	if request == nil {
		request = NewCreateAppInstanceFromYamlRequest()
	}
	response = NewCreateAppInstanceFromYamlResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateConfigMapsRequest() (request *UpdateConfigMapsRequest) {
	request = &UpdateConfigMapsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "UpdateConfigMaps")
	return
}

func NewUpdateConfigMapsResponse() (response *UpdateConfigMapsResponse) {
	response = &UpdateConfigMapsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新ConfigMap
func (c *Client) UpdateConfigMaps(request *UpdateConfigMapsRequest) (response *UpdateConfigMapsResponse, err error) {
	if request == nil {
		request = NewUpdateConfigMapsRequest()
	}
	response = NewUpdateConfigMapsResponse()
	err = c.Send(request, response)
	return
}

func NewGetRolloutSettingRequest() (request *GetRolloutSettingRequest) {
	request = &GetRolloutSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetRolloutSetting")
	return
}

func NewGetRolloutSettingResponse() (response *GetRolloutSettingResponse) {
	response = &GetRolloutSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取更新的设置
func (c *Client) GetRolloutSetting(request *GetRolloutSettingRequest) (response *GetRolloutSettingResponse, err error) {
	if request == nil {
		request = NewGetRolloutSettingRequest()
	}
	response = NewGetRolloutSettingResponse()
	err = c.Send(request, response)
	return
}

func NewGetRolloutStrategyRequest() (request *GetRolloutStrategyRequest) {
	request = &GetRolloutStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetRolloutStrategy")
	return
}

func NewGetRolloutStrategyResponse() (response *GetRolloutStrategyResponse) {
	response = &GetRolloutStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取更新策略
func (c *Client) GetRolloutStrategy(request *GetRolloutStrategyRequest) (response *GetRolloutStrategyResponse, err error) {
	if request == nil {
		request = NewGetRolloutStrategyRequest()
	}
	response = NewGetRolloutStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAppFromYamlRequest() (request *CreateAppFromYamlRequest) {
	request = &CreateAppFromYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CreateAppFromYaml")
	return
}

func NewCreateAppFromYamlResponse() (response *CreateAppFromYamlResponse) {
	response = &CreateAppFromYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从yaml创建app模板
func (c *Client) CreateAppFromYaml(request *CreateAppFromYamlRequest) (response *CreateAppFromYamlResponse, err error) {
	if request == nil {
		request = NewCreateAppFromYamlRequest()
	}
	response = NewCreateAppFromYamlResponse()
	err = c.Send(request, response)
	return
}

func NewSetRolloutStrategyRequest() (request *SetRolloutStrategyRequest) {
	request = &SetRolloutStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "SetRolloutStrategy")
	return
}

func NewSetRolloutStrategyResponse() (response *SetRolloutStrategyResponse) {
	response = &SetRolloutStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设定组件的更新策略和参数
func (c *Client) SetRolloutStrategy(request *SetRolloutStrategyRequest) (response *SetRolloutStrategyResponse, err error) {
	if request == nil {
		request = NewSetRolloutStrategyRequest()
	}
	response = NewSetRolloutStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAppRequest() (request *UpdateAppRequest) {
	request = &UpdateAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "UpdateApp")
	return
}

func NewUpdateAppResponse() (response *UpdateAppResponse) {
	response = &UpdateAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于修改未发布的模板，可以增删组件、trait，修改参数及其值
func (c *Client) UpdateApp(request *UpdateAppRequest) (response *UpdateAppResponse, err error) {
	if request == nil {
		request = NewUpdateAppRequest()
	}
	response = NewUpdateAppResponse()
	err = c.Send(request, response)
	return
}

func NewCreateScalerRequest() (request *CreateScalerRequest) {
	request = &CreateScalerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CreateScaler")
	return
}

func NewCreateScalerResponse() (response *CreateScalerResponse) {
	response = &CreateScalerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建弹性伸缩策略
func (c *Client) CreateScaler(request *CreateScalerRequest) (response *CreateScalerResponse, err error) {
	if request == nil {
		request = NewCreateScalerRequest()
	}
	response = NewCreateScalerResponse()
	err = c.Send(request, response)
	return
}

func NewListNamespacesRequest() (request *ListNamespacesRequest) {
	request = &ListNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListNamespaces")
	return
}

func NewListNamespacesResponse() (response *ListNamespacesResponse) {
	response = &ListNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询命名空间列表
func (c *Client) ListNamespaces(request *ListNamespacesRequest) (response *ListNamespacesResponse, err error) {
	if request == nil {
		request = NewListNamespacesRequest()
	}
	response = NewListNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewStartRolloutRequest() (request *StartRolloutRequest) {
	request = &StartRolloutRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "StartRollout")
	return
}

func NewStartRolloutResponse() (response *StartRolloutResponse) {
	response = &StartRolloutResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 触发更新
func (c *Client) StartRollout(request *StartRolloutRequest) (response *StartRolloutResponse, err error) {
	if request == nil {
		request = NewStartRolloutRequest()
	}
	response = NewStartRolloutResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpaScalerRequest() (request *DeleteVpaScalerRequest) {
	request = &DeleteVpaScalerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "DeleteVpaScaler")
	return
}

func NewDeleteVpaScalerResponse() (response *DeleteVpaScalerResponse) {
	response = &DeleteVpaScalerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DeleteVpaScaler
func (c *Client) DeleteVpaScaler(request *DeleteVpaScalerRequest) (response *DeleteVpaScalerResponse, err error) {
	if request == nil {
		request = NewDeleteVpaScalerRequest()
	}
	response = NewDeleteVpaScalerResponse()
	err = c.Send(request, response)
	return
}

func NewGetDimensionsRequest() (request *GetDimensionsRequest) {
	request = &GetDimensionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetDimensions")
	return
}

func NewGetDimensionsResponse() (response *GetDimensionsResponse) {
	response = &GetDimensionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetDimensions
func (c *Client) GetDimensions(request *GetDimensionsRequest) (response *GetDimensionsResponse, err error) {
	if request == nil {
		request = NewGetDimensionsRequest()
	}
	response = NewGetDimensionsResponse()
	err = c.Send(request, response)
	return
}

func NewGetSummaryRequest() (request *GetSummaryRequest) {
	request = &GetSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetSummary")
	return
}

func NewGetSummaryResponse() (response *GetSummaryResponse) {
	response = &GetSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用概要信息
func (c *Client) GetSummary(request *GetSummaryRequest) (response *GetSummaryResponse, err error) {
	if request == nil {
		request = NewGetSummaryRequest()
	}
	response = NewGetSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewGetSecretRequest() (request *GetSecretRequest) {
	request = &GetSecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetSecret")
	return
}

func NewGetSecretResponse() (response *GetSecretResponse) {
	response = &GetSecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取secret
func (c *Client) GetSecret(request *GetSecretRequest) (response *GetSecretResponse, err error) {
	if request == nil {
		request = NewGetSecretRequest()
	}
	response = NewGetSecretResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAppInsPreCheckToYamlRequest() (request *CreateAppInsPreCheckToYamlRequest) {
	request = &CreateAppInsPreCheckToYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CreateAppInsPreCheckToYaml")
	return
}

func NewCreateAppInsPreCheckToYamlResponse() (response *CreateAppInsPreCheckToYamlResponse) {
	response = &CreateAppInsPreCheckToYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateAppInsPreCheckToYaml
func (c *Client) CreateAppInsPreCheckToYaml(request *CreateAppInsPreCheckToYamlRequest) (response *CreateAppInsPreCheckToYamlResponse, err error) {
	if request == nil {
		request = NewCreateAppInsPreCheckToYamlRequest()
	}
	response = NewCreateAppInsPreCheckToYamlResponse()
	err = c.Send(request, response)
	return
}

func NewListTemplatesRequest() (request *ListTemplatesRequest) {
	request = &ListTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListTemplates")
	return
}

func NewListTemplatesResponse() (response *ListTemplatesResponse) {
	response = &ListTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询模板列表
func (c *Client) ListTemplates(request *ListTemplatesRequest) (response *ListTemplatesResponse, err error) {
	if request == nil {
		request = NewListTemplatesRequest()
	}
	response = NewListTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAppRequest() (request *DeleteAppRequest) {
	request = &DeleteAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "DeleteApp")
	return
}

func NewDeleteAppResponse() (response *DeleteAppResponse) {
	response = &DeleteAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除应用模板
func (c *Client) DeleteApp(request *DeleteAppRequest) (response *DeleteAppResponse, err error) {
	if request == nil {
		request = NewDeleteAppRequest()
	}
	response = NewDeleteAppResponse()
	err = c.Send(request, response)
	return
}

func NewGetAppFilesRequest() (request *GetAppFilesRequest) {
	request = &GetAppFilesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetAppFiles")
	return
}

func NewGetAppFilesResponse() (response *GetAppFilesResponse) {
	response = &GetAppFilesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取helm应用的文件
func (c *Client) GetAppFiles(request *GetAppFilesRequest) (response *GetAppFilesResponse, err error) {
	if request == nil {
		request = NewGetAppFilesRequest()
	}
	response = NewGetAppFilesResponse()
	err = c.Send(request, response)
	return
}

func NewGetRevisionYamlRequest() (request *GetRevisionYamlRequest) {
	request = &GetRevisionYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetRevisionYaml")
	return
}

func NewGetRevisionYamlResponse() (response *GetRevisionYamlResponse) {
	response = &GetRevisionYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取revision的yaml信息
func (c *Client) GetRevisionYaml(request *GetRevisionYamlRequest) (response *GetRevisionYamlResponse, err error) {
	if request == nil {
		request = NewGetRevisionYamlRequest()
	}
	response = NewGetRevisionYamlResponse()
	err = c.Send(request, response)
	return
}

func NewListPodsRequest() (request *ListPodsRequest) {
	request = &ListPodsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListPods")
	return
}

func NewListPodsResponse() (response *ListPodsResponse) {
	response = &ListPodsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Pod列表
func (c *Client) ListPods(request *ListPodsRequest) (response *ListPodsResponse, err error) {
	if request == nil {
		request = NewListPodsRequest()
	}
	response = NewListPodsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConfigMapFromYamlRequest() (request *CreateConfigMapFromYamlRequest) {
	request = &CreateConfigMapFromYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CreateConfigMapFromYaml")
	return
}

func NewCreateConfigMapFromYamlResponse() (response *CreateConfigMapFromYamlResponse) {
	response = &CreateConfigMapFromYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 使用yaml创建configmap
func (c *Client) CreateConfigMapFromYaml(request *CreateConfigMapFromYamlRequest) (response *CreateConfigMapFromYamlResponse, err error) {
	if request == nil {
		request = NewCreateConfigMapFromYamlRequest()
	}
	response = NewCreateConfigMapFromYamlResponse()
	err = c.Send(request, response)
	return
}

func NewGetRolloutStatusRequest() (request *GetRolloutStatusRequest) {
	request = &GetRolloutStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetRolloutStatus")
	return
}

func NewGetRolloutStatusResponse() (response *GetRolloutStatusResponse) {
	response = &GetRolloutStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取更新的状态
func (c *Client) GetRolloutStatus(request *GetRolloutStatusRequest) (response *GetRolloutStatusResponse, err error) {
	if request == nil {
		request = NewGetRolloutStatusRequest()
	}
	response = NewGetRolloutStatusResponse()
	err = c.Send(request, response)
	return
}

func NewGetAppRequest() (request *GetAppRequest) {
	request = &GetAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetApp")
	return
}

func NewGetAppResponse() (response *GetAppResponse) {
	response = &GetAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取App模板详情
func (c *Client) GetApp(request *GetAppRequest) (response *GetAppResponse, err error) {
	if request == nil {
		request = NewGetAppRequest()
	}
	response = NewGetAppResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateDescriptionRequest() (request *UpdateDescriptionRequest) {
	request = &UpdateDescriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "UpdateDescription")
	return
}

func NewUpdateDescriptionResponse() (response *UpdateDescriptionResponse) {
	response = &UpdateDescriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改实例描述信息
func (c *Client) UpdateDescription(request *UpdateDescriptionRequest) (response *UpdateDescriptionResponse, err error) {
	if request == nil {
		request = NewUpdateDescriptionRequest()
	}
	response = NewUpdateDescriptionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVpaScalerRequest() (request *CreateVpaScalerRequest) {
	request = &CreateVpaScalerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CreateVpaScaler")
	return
}

func NewCreateVpaScalerResponse() (response *CreateVpaScalerResponse) {
	response = &CreateVpaScalerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateVpaScaler
func (c *Client) CreateVpaScaler(request *CreateVpaScalerRequest) (response *CreateVpaScalerResponse, err error) {
	if request == nil {
		request = NewCreateVpaScalerRequest()
	}
	response = NewCreateVpaScalerResponse()
	err = c.Send(request, response)
	return
}

func NewListAppProductsRequest() (request *ListAppProductsRequest) {
	request = &ListAppProductsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListAppProducts")
	return
}

func NewListAppProductsResponse() (response *ListAppProductsResponse) {
	response = &ListAppProductsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ListAppProducts
func (c *Client) ListAppProducts(request *ListAppProductsRequest) (response *ListAppProductsResponse, err error) {
	if request == nil {
		request = NewListAppProductsRequest()
	}
	response = NewListAppProductsResponse()
	err = c.Send(request, response)
	return
}

func NewGetAvailableStrategyRequest() (request *GetAvailableStrategyRequest) {
	request = &GetAvailableStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetAvailableStrategy")
	return
}

func NewGetAvailableStrategyResponse() (response *GetAvailableStrategyResponse) {
	response = &GetAvailableStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Workload支持的升级策略
func (c *Client) GetAvailableStrategy(request *GetAvailableStrategyRequest) (response *GetAvailableStrategyResponse, err error) {
	if request == nil {
		request = NewGetAvailableStrategyRequest()
	}
	response = NewGetAvailableStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewGetBriefServiceInstanceRequest() (request *GetBriefServiceInstanceRequest) {
	request = &GetBriefServiceInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetBriefServiceInstance")
	return
}

func NewGetBriefServiceInstanceResponse() (response *GetBriefServiceInstanceResponse) {
	response = &GetBriefServiceInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取服务实例的简明信息
func (c *Client) GetBriefServiceInstance(request *GetBriefServiceInstanceRequest) (response *GetBriefServiceInstanceResponse, err error) {
	if request == nil {
		request = NewGetBriefServiceInstanceRequest()
	}
	response = NewGetBriefServiceInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewGetDetailVpaScalerByIdRequest() (request *GetDetailVpaScalerByIdRequest) {
	request = &GetDetailVpaScalerByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetDetailVpaScalerById")
	return
}

func NewGetDetailVpaScalerByIdResponse() (response *GetDetailVpaScalerByIdResponse) {
	response = &GetDetailVpaScalerByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetDetailVpaScalerById
func (c *Client) GetDetailVpaScalerById(request *GetDetailVpaScalerByIdRequest) (response *GetDetailVpaScalerByIdResponse, err error) {
	if request == nil {
		request = NewGetDetailVpaScalerByIdRequest()
	}
	response = NewGetDetailVpaScalerByIdResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateScalerRequest() (request *UpdateScalerRequest) {
	request = &UpdateScalerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "UpdateScaler")
	return
}

func NewUpdateScalerResponse() (response *UpdateScalerResponse) {
	response = &UpdateScalerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新弹性伸缩策略
func (c *Client) UpdateScaler(request *UpdateScalerRequest) (response *UpdateScalerResponse, err error) {
	if request == nil {
		request = NewUpdateScalerRequest()
	}
	response = NewUpdateScalerResponse()
	err = c.Send(request, response)
	return
}

func NewCheckAppInstanceExistRequest() (request *CheckAppInstanceExistRequest) {
	request = &CheckAppInstanceExistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CheckAppInstanceExist")
	return
}

func NewCheckAppInstanceExistResponse() (response *CheckAppInstanceExistResponse) {
	response = &CheckAppInstanceExistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查app实例是否已经存在
func (c *Client) CheckAppInstanceExist(request *CheckAppInstanceExistRequest) (response *CheckAppInstanceExistResponse, err error) {
	if request == nil {
		request = NewCheckAppInstanceExistRequest()
	}
	response = NewCheckAppInstanceExistResponse()
	err = c.Send(request, response)
	return
}

func NewGetAppConfigRequest() (request *GetAppConfigRequest) {
	request = &GetAppConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetAppConfig")
	return
}

func NewGetAppConfigResponse() (response *GetAppConfigResponse) {
	response = &GetAppConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取helm 应用的配置
func (c *Client) GetAppConfig(request *GetAppConfigRequest) (response *GetAppConfigResponse, err error) {
	if request == nil {
		request = NewGetAppConfigRequest()
	}
	response = NewGetAppConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAppInstanceByChartValuesRequest() (request *CreateAppInstanceByChartValuesRequest) {
	request = &CreateAppInstanceByChartValuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CreateAppInstanceByChartValues")
	return
}

func NewCreateAppInstanceByChartValuesResponse() (response *CreateAppInstanceByChartValuesResponse) {
	response = &CreateAppInstanceByChartValuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据模板及values创建实例
func (c *Client) CreateAppInstanceByChartValues(request *CreateAppInstanceByChartValuesRequest) (response *CreateAppInstanceByChartValuesResponse, err error) {
	if request == nil {
		request = NewCreateAppInstanceByChartValuesRequest()
	}
	response = NewCreateAppInstanceByChartValuesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSecretRequest() (request *DeleteSecretRequest) {
	request = &DeleteSecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "DeleteSecret")
	return
}

func NewDeleteSecretResponse() (response *DeleteSecretResponse) {
	response = &DeleteSecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除secret
func (c *Client) DeleteSecret(request *DeleteSecretRequest) (response *DeleteSecretResponse, err error) {
	if request == nil {
		request = NewDeleteSecretRequest()
	}
	response = NewDeleteSecretResponse()
	err = c.Send(request, response)
	return
}

func NewGetAppInstanceRequest() (request *GetAppInstanceRequest) {
	request = &GetAppInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetAppInstance")
	return
}

func NewGetAppInstanceResponse() (response *GetAppInstanceResponse) {
	response = &GetAppInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询app实例详情
func (c *Client) GetAppInstance(request *GetAppInstanceRequest) (response *GetAppInstanceResponse, err error) {
	if request == nil {
		request = NewGetAppInstanceRequest()
	}
	response = NewGetAppInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewGetConfigMapRequest() (request *GetConfigMapRequest) {
	request = &GetConfigMapRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetConfigMap")
	return
}

func NewGetConfigMapResponse() (response *GetConfigMapResponse) {
	response = &GetConfigMapResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询ConfigMap详情
func (c *Client) GetConfigMap(request *GetConfigMapRequest) (response *GetConfigMapResponse, err error) {
	if request == nil {
		request = NewGetConfigMapRequest()
	}
	response = NewGetConfigMapResponse()
	err = c.Send(request, response)
	return
}

func NewGetTadPkgRequest() (request *GetTadPkgRequest) {
	request = &GetTadPkgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetTadPkg")
	return
}

func NewGetTadPkgResponse() (response *GetTadPkgResponse) {
	response = &GetTadPkgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用包的二进制
func (c *Client) GetTadPkg(request *GetTadPkgRequest) (response *GetTadPkgResponse, err error) {
	if request == nil {
		request = NewGetTadPkgRequest()
	}
	response = NewGetTadPkgResponse()
	err = c.Send(request, response)
	return
}

func NewRollbackRequest() (request *RollbackRequest) {
	request = &RollbackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "Rollback")
	return
}

func NewRollbackResponse() (response *RollbackResponse) {
	response = &RollbackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回滚本次更新操作
func (c *Client) Rollback(request *RollbackRequest) (response *RollbackResponse, err error) {
	if request == nil {
		request = NewRollbackRequest()
	}
	response = NewRollbackResponse()
	err = c.Send(request, response)
	return
}

func NewGetAppInstanceYamlRequest() (request *GetAppInstanceYamlRequest) {
	request = &GetAppInstanceYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetAppInstanceYaml")
	return
}

func NewGetAppInstanceYamlResponse() (response *GetAppInstanceYamlResponse) {
	response = &GetAppInstanceYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口获取应用实例yaml信息
func (c *Client) GetAppInstanceYaml(request *GetAppInstanceYamlRequest) (response *GetAppInstanceYamlResponse, err error) {
	if request == nil {
		request = NewGetAppInstanceYamlRequest()
	}
	response = NewGetAppInstanceYamlResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationComponentsSummaryRequest() (request *GetApplicationComponentsSummaryRequest) {
	request = &GetApplicationComponentsSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetApplicationComponentsSummary")
	return
}

func NewGetApplicationComponentsSummaryResponse() (response *GetApplicationComponentsSummaryResponse) {
	response = &GetApplicationComponentsSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) GetApplicationComponentsSummary(request *GetApplicationComponentsSummaryRequest) (response *GetApplicationComponentsSummaryResponse, err error) {
	if request == nil {
		request = NewGetApplicationComponentsSummaryRequest()
	}
	response = NewGetApplicationComponentsSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewGetUpgradeConfigRequest() (request *GetUpgradeConfigRequest) {
	request = &GetUpgradeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetUpgradeConfig")
	return
}

func NewGetUpgradeConfigResponse() (response *GetUpgradeConfigResponse) {
	response = &GetUpgradeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 升级配置查询
func (c *Client) GetUpgradeConfig(request *GetUpgradeConfigRequest) (response *GetUpgradeConfigResponse, err error) {
	if request == nil {
		request = NewGetUpgradeConfigRequest()
	}
	response = NewGetUpgradeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewListVpaScalerRequest() (request *ListVpaScalerRequest) {
	request = &ListVpaScalerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListVpaScaler")
	return
}

func NewListVpaScalerResponse() (response *ListVpaScalerResponse) {
	response = &ListVpaScalerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ListVpaScaler
func (c *Client) ListVpaScaler(request *ListVpaScalerRequest) (response *ListVpaScalerResponse, err error) {
	if request == nil {
		request = NewListVpaScalerRequest()
	}
	response = NewListVpaScalerResponse()
	err = c.Send(request, response)
	return
}

func NewGetComponentInstanceRequest() (request *GetComponentInstanceRequest) {
	request = &GetComponentInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetComponentInstance")
	return
}

func NewGetComponentInstanceResponse() (response *GetComponentInstanceResponse) {
	response = &GetComponentInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用实例列表信息
func (c *Client) GetComponentInstance(request *GetComponentInstanceRequest) (response *GetComponentInstanceResponse, err error) {
	if request == nil {
		request = NewGetComponentInstanceRequest()
	}
	response = NewGetComponentInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewScaleInstanceReplicasRequest() (request *ScaleInstanceReplicasRequest) {
	request = &ScaleInstanceReplicasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ScaleInstanceReplicas")
	return
}

func NewScaleInstanceReplicasResponse() (response *ScaleInstanceReplicasResponse) {
	response = &ScaleInstanceReplicasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器组件扩缩容
func (c *Client) ScaleInstanceReplicas(request *ScaleInstanceReplicasRequest) (response *ScaleInstanceReplicasResponse, err error) {
	if request == nil {
		request = NewScaleInstanceReplicasRequest()
	}
	response = NewScaleInstanceReplicasResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSecretRequest() (request *CreateSecretRequest) {
	request = &CreateSecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CreateSecret")
	return
}

func NewCreateSecretResponse() (response *CreateSecretResponse) {
	response = &CreateSecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建secret
func (c *Client) CreateSecret(request *CreateSecretRequest) (response *CreateSecretResponse, err error) {
	if request == nil {
		request = NewCreateSecretRequest()
	}
	response = NewCreateSecretResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationSummaryRequest() (request *GetApplicationSummaryRequest) {
	request = &GetApplicationSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetApplicationSummary")
	return
}

func NewGetApplicationSummaryResponse() (response *GetApplicationSummaryResponse) {
	response = &GetApplicationSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) GetApplicationSummary(request *GetApplicationSummaryRequest) (response *GetApplicationSummaryResponse, err error) {
	if request == nil {
		request = NewGetApplicationSummaryRequest()
	}
	response = NewGetApplicationSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewGetCanaryTrafficRequest() (request *GetCanaryTrafficRequest) {
	request = &GetCanaryTrafficRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetCanaryTraffic")
	return
}

func NewGetCanaryTrafficResponse() (response *GetCanaryTrafficResponse) {
	response = &GetCanaryTrafficResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取灰度发布的流量设置
func (c *Client) GetCanaryTraffic(request *GetCanaryTrafficRequest) (response *GetCanaryTrafficResponse, err error) {
	if request == nil {
		request = NewGetCanaryTrafficRequest()
	}
	response = NewGetCanaryTrafficResponse()
	err = c.Send(request, response)
	return
}

func NewGetChartPkgRequest() (request *GetChartPkgRequest) {
	request = &GetChartPkgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetChartPkg")
	return
}

func NewGetChartPkgResponse() (response *GetChartPkgResponse) {
	response = &GetChartPkgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetChartPkg
func (c *Client) GetChartPkg(request *GetChartPkgRequest) (response *GetChartPkgResponse, err error) {
	if request == nil {
		request = NewGetChartPkgRequest()
	}
	response = NewGetChartPkgResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteScalerRequest() (request *DeleteScalerRequest) {
	request = &DeleteScalerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "DeleteScaler")
	return
}

func NewDeleteScalerResponse() (response *DeleteScalerResponse) {
	response = &DeleteScalerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除弹性伸缩策略
func (c *Client) DeleteScaler(request *DeleteScalerRequest) (response *DeleteScalerResponse, err error) {
	if request == nil {
		request = NewDeleteScalerRequest()
	}
	response = NewDeleteScalerResponse()
	err = c.Send(request, response)
	return
}

func NewCheckAppExistRequest() (request *CheckAppExistRequest) {
	request = &CheckAppExistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CheckAppExist")
	return
}

func NewCheckAppExistResponse() (response *CheckAppExistResponse) {
	response = &CheckAppExistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查app是否已经存在
func (c *Client) CheckAppExist(request *CheckAppExistRequest) (response *CheckAppExistResponse, err error) {
	if request == nil {
		request = NewCheckAppExistRequest()
	}
	response = NewCheckAppExistResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAppRequest() (request *CreateAppRequest) {
	request = &CreateAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CreateApp")
	return
}

func NewCreateAppResponse() (response *CreateAppResponse) {
	response = &CreateAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建应用模板
func (c *Client) CreateApp(request *CreateAppRequest) (response *CreateAppResponse, err error) {
	if request == nil {
		request = NewCreateAppRequest()
	}
	response = NewCreateAppResponse()
	err = c.Send(request, response)
	return
}

func NewListEventsRequest() (request *ListEventsRequest) {
	request = &ListEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListEvents")
	return
}

func NewListEventsResponse() (response *ListEventsResponse) {
	response = &ListEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询事件
func (c *Client) ListEvents(request *ListEventsRequest) (response *ListEventsResponse, err error) {
	if request == nil {
		request = NewListEventsRequest()
	}
	response = NewListEventsResponse()
	err = c.Send(request, response)
	return
}

func NewGetAppFileContentRequest() (request *GetAppFileContentRequest) {
	request = &GetAppFileContentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetAppFileContent")
	return
}

func NewGetAppFileContentResponse() (response *GetAppFileContentResponse) {
	response = &GetAppFileContentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用模板中文件的具体内容
func (c *Client) GetAppFileContent(request *GetAppFileContentRequest) (response *GetAppFileContentResponse, err error) {
	if request == nil {
		request = NewGetAppFileContentRequest()
	}
	response = NewGetAppFileContentResponse()
	err = c.Send(request, response)
	return
}

func NewSaveRevisionByComponentRequest() (request *SaveRevisionByComponentRequest) {
	request = &SaveRevisionByComponentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "SaveRevisionByComponent")
	return
}

func NewSaveRevisionByComponentResponse() (response *SaveRevisionByComponentResponse) {
	response = &SaveRevisionByComponentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 指定一个组件保存revision版本
func (c *Client) SaveRevisionByComponent(request *SaveRevisionByComponentRequest) (response *SaveRevisionByComponentResponse, err error) {
	if request == nil {
		request = NewSaveRevisionByComponentRequest()
	}
	response = NewSaveRevisionByComponentResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAppFromYamlRequest() (request *UpdateAppFromYamlRequest) {
	request = &UpdateAppFromYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "UpdateAppFromYaml")
	return
}

func NewUpdateAppFromYamlResponse() (response *UpdateAppFromYamlResponse) {
	response = &UpdateAppFromYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口通过yaml修改未发布的模板，可以增删组件、trait，修改参数及其值
func (c *Client) UpdateAppFromYaml(request *UpdateAppFromYamlRequest) (response *UpdateAppFromYamlResponse, err error) {
	if request == nil {
		request = NewUpdateAppFromYamlRequest()
	}
	response = NewUpdateAppFromYamlResponse()
	err = c.Send(request, response)
	return
}

func NewApplyRevisionRequest() (request *ApplyRevisionRequest) {
	request = &ApplyRevisionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ApplyRevision")
	return
}

func NewApplyRevisionResponse() (response *ApplyRevisionResponse) {
	response = &ApplyRevisionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 部署revision版本
func (c *Client) ApplyRevision(request *ApplyRevisionRequest) (response *ApplyRevisionResponse, err error) {
	if request == nil {
		request = NewApplyRevisionRequest()
	}
	response = NewApplyRevisionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAppInstanceRequest() (request *CreateAppInstanceRequest) {
	request = &CreateAppInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CreateAppInstance")
	return
}

func NewCreateAppInstanceResponse() (response *CreateAppInstanceResponse) {
	response = &CreateAppInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建应用实例
func (c *Client) CreateAppInstance(request *CreateAppInstanceRequest) (response *CreateAppInstanceResponse, err error) {
	if request == nil {
		request = NewCreateAppInstanceRequest()
	}
	response = NewCreateAppInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewSaveRevisionRequest() (request *SaveRevisionRequest) {
	request = &SaveRevisionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "SaveRevision")
	return
}

func NewSaveRevisionResponse() (response *SaveRevisionResponse) {
	response = &SaveRevisionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存revision版本
func (c *Client) SaveRevision(request *SaveRevisionRequest) (response *SaveRevisionResponse, err error) {
	if request == nil {
		request = NewSaveRevisionRequest()
	}
	response = NewSaveRevisionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAppInstanceRequest() (request *DeleteAppInstanceRequest) {
	request = &DeleteAppInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "DeleteAppInstance")
	return
}

func NewDeleteAppInstanceResponse() (response *DeleteAppInstanceResponse) {
	response = &DeleteAppInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除应用实例
func (c *Client) DeleteAppInstance(request *DeleteAppInstanceRequest) (response *DeleteAppInstanceResponse, err error) {
	if request == nil {
		request = NewDeleteAppInstanceRequest()
	}
	response = NewDeleteAppInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewListAppCategoriesRequest() (request *ListAppCategoriesRequest) {
	request = &ListAppCategoriesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListAppCategories")
	return
}

func NewListAppCategoriesResponse() (response *ListAppCategoriesResponse) {
	response = &ListAppCategoriesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用模板分类列表
func (c *Client) ListAppCategories(request *ListAppCategoriesRequest) (response *ListAppCategoriesResponse, err error) {
	if request == nil {
		request = NewListAppCategoriesRequest()
	}
	response = NewListAppCategoriesResponse()
	err = c.Send(request, response)
	return
}

func NewListScalerEventsRequest() (request *ListScalerEventsRequest) {
	request = &ListScalerEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListScalerEvents")
	return
}

func NewListScalerEventsResponse() (response *ListScalerEventsResponse) {
	response = &ListScalerEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取弹性伸缩历史事件
func (c *Client) ListScalerEvents(request *ListScalerEventsRequest) (response *ListScalerEventsResponse, err error) {
	if request == nil {
		request = NewListScalerEventsRequest()
	}
	response = NewListScalerEventsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSecretRequest() (request *UpdateSecretRequest) {
	request = &UpdateSecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "UpdateSecret")
	return
}

func NewUpdateSecretResponse() (response *UpdateSecretResponse) {
	response = &UpdateSecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新secret
func (c *Client) UpdateSecret(request *UpdateSecretRequest) (response *UpdateSecretResponse, err error) {
	if request == nil {
		request = NewUpdateSecretRequest()
	}
	response = NewUpdateSecretResponse()
	err = c.Send(request, response)
	return
}

func NewGetAppYamlRequest() (request *GetAppYamlRequest) {
	request = &GetAppYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetAppYaml")
	return
}

func NewGetAppYamlResponse() (response *GetAppYamlResponse) {
	response = &GetAppYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用模板的yaml
func (c *Client) GetAppYaml(request *GetAppYamlRequest) (response *GetAppYamlResponse, err error) {
	if request == nil {
		request = NewGetAppYamlRequest()
	}
	response = NewGetAppYamlResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAppsFromTadProductPkgRequest() (request *CreateAppsFromTadProductPkgRequest) {
	request = &CreateAppsFromTadProductPkgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CreateAppsFromTadProductPkg")
	return
}

func NewCreateAppsFromTadProductPkgResponse() (response *CreateAppsFromTadProductPkgResponse) {
	response = &CreateAppsFromTadProductPkgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateAppsFromTadProductPkg 批量上传模版/产品包
func (c *Client) CreateAppsFromTadProductPkg(request *CreateAppsFromTadProductPkgRequest) (response *CreateAppsFromTadProductPkgResponse, err error) {
	if request == nil {
		request = NewCreateAppsFromTadProductPkgRequest()
	}
	response = NewCreateAppsFromTadProductPkgResponse()
	err = c.Send(request, response)
	return
}

func NewGetScalerYamlRequest() (request *GetScalerYamlRequest) {
	request = &GetScalerYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetScalerYaml")
	return
}

func NewGetScalerYamlResponse() (response *GetScalerYamlResponse) {
	response = &GetScalerYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取弹性伸缩的yaml
func (c *Client) GetScalerYaml(request *GetScalerYamlRequest) (response *GetScalerYamlResponse, err error) {
	if request == nil {
		request = NewGetScalerYamlRequest()
	}
	response = NewGetScalerYamlResponse()
	err = c.Send(request, response)
	return
}

func NewEnableVpaScalerRequest() (request *EnableVpaScalerRequest) {
	request = &EnableVpaScalerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "EnableVpaScaler")
	return
}

func NewEnableVpaScalerResponse() (response *EnableVpaScalerResponse) {
	response = &EnableVpaScalerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// EnableVpaScaler
func (c *Client) EnableVpaScaler(request *EnableVpaScalerRequest) (response *EnableVpaScalerResponse, err error) {
	if request == nil {
		request = NewEnableVpaScalerRequest()
	}
	response = NewEnableVpaScalerResponse()
	err = c.Send(request, response)
	return
}

func NewGetEventVpaScalerByIdRequest() (request *GetEventVpaScalerByIdRequest) {
	request = &GetEventVpaScalerByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetEventVpaScalerById")
	return
}

func NewGetEventVpaScalerByIdResponse() (response *GetEventVpaScalerByIdResponse) {
	response = &GetEventVpaScalerByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetEventVpaScalerById
func (c *Client) GetEventVpaScalerById(request *GetEventVpaScalerByIdRequest) (response *GetEventVpaScalerByIdResponse, err error) {
	if request == nil {
		request = NewGetEventVpaScalerByIdRequest()
	}
	response = NewGetEventVpaScalerByIdResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAppPicFileRequest() (request *UpdateAppPicFileRequest) {
	request = &UpdateAppPicFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "UpdateAppPicFile")
	return
}

func NewUpdateAppPicFileResponse() (response *UpdateAppPicFileResponse) {
	response = &UpdateAppPicFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// UpdateAppPicFile
func (c *Client) UpdateAppPicFile(request *UpdateAppPicFileRequest) (response *UpdateAppPicFileResponse, err error) {
	if request == nil {
		request = NewUpdateAppPicFileRequest()
	}
	response = NewUpdateAppPicFileResponse()
	err = c.Send(request, response)
	return
}

func NewGetInstanceReplicasRequest() (request *GetInstanceReplicasRequest) {
	request = &GetInstanceReplicasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetInstanceReplicas")
	return
}

func NewGetInstanceReplicasResponse() (response *GetInstanceReplicasResponse) {
	response = &GetInstanceReplicasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器组件副本数查询
func (c *Client) GetInstanceReplicas(request *GetInstanceReplicasRequest) (response *GetInstanceReplicasResponse, err error) {
	if request == nil {
		request = NewGetInstanceReplicasRequest()
	}
	response = NewGetInstanceReplicasResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRevisionRequest() (request *DeleteRevisionRequest) {
	request = &DeleteRevisionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "DeleteRevision")
	return
}

func NewDeleteRevisionResponse() (response *DeleteRevisionResponse) {
	response = &DeleteRevisionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除revision版本
func (c *Client) DeleteRevision(request *DeleteRevisionRequest) (response *DeleteRevisionResponse, err error) {
	if request == nil {
		request = NewDeleteRevisionRequest()
	}
	response = NewDeleteRevisionResponse()
	err = c.Send(request, response)
	return
}

func NewGetClusterSummaryRequest() (request *GetClusterSummaryRequest) {
	request = &GetClusterSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetClusterSummary")
	return
}

func NewGetClusterSummaryResponse() (response *GetClusterSummaryResponse) {
	response = &GetClusterSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群概览数据
func (c *Client) GetClusterSummary(request *GetClusterSummaryRequest) (response *GetClusterSummaryResponse, err error) {
	if request == nil {
		request = NewGetClusterSummaryRequest()
	}
	response = NewGetClusterSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewListAppsRequest() (request *ListAppsRequest) {
	request = &ListAppsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListApps")
	return
}

func NewListAppsResponse() (response *ListAppsResponse) {
	response = &ListAppsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用模板列表
func (c *Client) ListApps(request *ListAppsRequest) (response *ListAppsResponse, err error) {
	if request == nil {
		request = NewListAppsRequest()
	}
	response = NewListAppsResponse()
	err = c.Send(request, response)
	return
}

func NewListConfigMapsRequest() (request *ListConfigMapsRequest) {
	request = &ListConfigMapsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListConfigMaps")
	return
}

func NewListConfigMapsResponse() (response *ListConfigMapsResponse) {
	response = &ListConfigMapsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询ConfigMap列表
func (c *Client) ListConfigMaps(request *ListConfigMapsRequest) (response *ListConfigMapsResponse, err error) {
	if request == nil {
		request = NewListConfigMapsRequest()
	}
	response = NewListConfigMapsResponse()
	err = c.Send(request, response)
	return
}

func NewListHostsRequest() (request *ListHostsRequest) {
	request = &ListHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListHosts")
	return
}

func NewListHostsResponse() (response *ListHostsResponse) {
	response = &ListHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主机列表
func (c *Client) ListHosts(request *ListHostsRequest) (response *ListHostsResponse, err error) {
	if request == nil {
		request = NewListHostsRequest()
	}
	response = NewListHostsResponse()
	err = c.Send(request, response)
	return
}

func NewEditVpaScalerRequest() (request *EditVpaScalerRequest) {
	request = &EditVpaScalerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "EditVpaScaler")
	return
}

func NewEditVpaScalerResponse() (response *EditVpaScalerResponse) {
	response = &EditVpaScalerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// EditVpaScaler
func (c *Client) EditVpaScaler(request *EditVpaScalerRequest) (response *EditVpaScalerResponse, err error) {
	if request == nil {
		request = NewEditVpaScalerRequest()
	}
	response = NewEditVpaScalerResponse()
	err = c.Send(request, response)
	return
}

func NewListAppInstanceNamespacesRequest() (request *ListAppInstanceNamespacesRequest) {
	request = &ListAppInstanceNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListAppInstanceNamespaces")
	return
}

func NewListAppInstanceNamespacesResponse() (response *ListAppInstanceNamespacesResponse) {
	response = &ListAppInstanceNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基于某模板部署的所有应用实例所在的命名空间
func (c *Client) ListAppInstanceNamespaces(request *ListAppInstanceNamespacesRequest) (response *ListAppInstanceNamespacesResponse, err error) {
	if request == nil {
		request = NewListAppInstanceNamespacesRequest()
	}
	response = NewListAppInstanceNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewListSecretsRequest() (request *ListSecretsRequest) {
	request = &ListSecretsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListSecrets")
	return
}

func NewListSecretsResponse() (response *ListSecretsResponse) {
	response = &ListSecretsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取secret列表
func (c *Client) ListSecrets(request *ListSecretsRequest) (response *ListSecretsResponse, err error) {
	if request == nil {
		request = NewListSecretsRequest()
	}
	response = NewListSecretsResponse()
	err = c.Send(request, response)
	return
}

func NewRedeployAppInstanceRequest() (request *RedeployAppInstanceRequest) {
	request = &RedeployAppInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "RedeployAppInstance")
	return
}

func NewRedeployAppInstanceResponse() (response *RedeployAppInstanceResponse) {
	response = &RedeployAppInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重新部署应用版本
func (c *Client) RedeployAppInstance(request *RedeployAppInstanceRequest) (response *RedeployAppInstanceResponse, err error) {
	if request == nil {
		request = NewRedeployAppInstanceRequest()
	}
	response = NewRedeployAppInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewListAppVersionsRequest() (request *ListAppVersionsRequest) {
	request = &ListAppVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListAppVersions")
	return
}

func NewListAppVersionsResponse() (response *ListAppVersionsResponse) {
	response = &ListAppVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定应用模板版本列表
func (c *Client) ListAppVersions(request *ListAppVersionsRequest) (response *ListAppVersionsResponse, err error) {
	if request == nil {
		request = NewListAppVersionsRequest()
	}
	response = NewListAppVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAppInstanceIngressRequest() (request *UpdateAppInstanceIngressRequest) {
	request = &UpdateAppInstanceIngressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "UpdateAppInstanceIngress")
	return
}

func NewUpdateAppInstanceIngressResponse() (response *UpdateAppInstanceIngressResponse) {
	response = &UpdateAppInstanceIngressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新应用内的访问入口信息
func (c *Client) UpdateAppInstanceIngress(request *UpdateAppInstanceIngressRequest) (response *UpdateAppInstanceIngressResponse, err error) {
	if request == nil {
		request = NewUpdateAppInstanceIngressRequest()
	}
	response = NewUpdateAppInstanceIngressResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCustomerAppInstanceRequest() (request *CreateCustomerAppInstanceRequest) {
	request = &CreateCustomerAppInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CreateCustomerAppInstance")
	return
}

func NewCreateCustomerAppInstanceResponse() (response *CreateCustomerAppInstanceResponse) {
	response = &CreateCustomerAppInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建自定义应用
func (c *Client) CreateCustomerAppInstance(request *CreateCustomerAppInstanceRequest) (response *CreateCustomerAppInstanceResponse, err error) {
	if request == nil {
		request = NewCreateCustomerAppInstanceRequest()
	}
	response = NewCreateCustomerAppInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewGetAppInstanceComponentsRequest() (request *GetAppInstanceComponentsRequest) {
	request = &GetAppInstanceComponentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetAppInstanceComponents")
	return
}

func NewGetAppInstanceComponentsResponse() (response *GetAppInstanceComponentsResponse) {
	response = &GetAppInstanceComponentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取AppInstance的组件列表
func (c *Client) GetAppInstanceComponents(request *GetAppInstanceComponentsRequest) (response *GetAppInstanceComponentsResponse, err error) {
	if request == nil {
		request = NewGetAppInstanceComponentsRequest()
	}
	response = NewGetAppInstanceComponentsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAppFromChartPkgRequest() (request *CreateAppFromChartPkgRequest) {
	request = &CreateAppFromChartPkgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CreateAppFromChartPkg")
	return
}

func NewCreateAppFromChartPkgResponse() (response *CreateAppFromChartPkgResponse) {
	response = &CreateAppFromChartPkgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateAppFromChartPkg
func (c *Client) CreateAppFromChartPkg(request *CreateAppFromChartPkgRequest) (response *CreateAppFromChartPkgResponse, err error) {
	if request == nil {
		request = NewCreateAppFromChartPkgRequest()
	}
	response = NewCreateAppFromChartPkgResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConfigMapsRequest() (request *CreateConfigMapsRequest) {
	request = &CreateConfigMapsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "CreateConfigMaps")
	return
}

func NewCreateConfigMapsResponse() (response *CreateConfigMapsResponse) {
	response = &CreateConfigMapsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建ConfigMap
func (c *Client) CreateConfigMaps(request *CreateConfigMapsRequest) (response *CreateConfigMapsResponse, err error) {
	if request == nil {
		request = NewCreateConfigMapsRequest()
	}
	response = NewCreateConfigMapsResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationResourceUsageSummaryRequest() (request *GetApplicationResourceUsageSummaryRequest) {
	request = &GetApplicationResourceUsageSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetApplicationResourceUsageSummary")
	return
}

func NewGetApplicationResourceUsageSummaryResponse() (response *GetApplicationResourceUsageSummaryResponse) {
	response = &GetApplicationResourceUsageSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) GetApplicationResourceUsageSummary(request *GetApplicationResourceUsageSummaryRequest) (response *GetApplicationResourceUsageSummaryResponse, err error) {
	if request == nil {
		request = NewGetApplicationResourceUsageSummaryRequest()
	}
	response = NewGetApplicationResourceUsageSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewListAppInstancesRequest() (request *ListAppInstancesRequest) {
	request = &ListAppInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "ListAppInstances")
	return
}

func NewListAppInstancesResponse() (response *ListAppInstancesResponse) {
	response = &ListAppInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询app实例列表
func (c *Client) ListAppInstances(request *ListAppInstancesRequest) (response *ListAppInstancesResponse, err error) {
	if request == nil {
		request = NewListAppInstancesRequest()
	}
	response = NewListAppInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewGetBriefComponentInstanceRequest() (request *GetBriefComponentInstanceRequest) {
	request = &GetBriefComponentInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "GetBriefComponentInstance")
	return
}

func NewGetBriefComponentInstanceResponse() (response *GetBriefComponentInstanceResponse) {
	response = &GetBriefComponentInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用实例列表信息
func (c *Client) GetBriefComponentInstance(request *GetBriefComponentInstanceRequest) (response *GetBriefComponentInstanceResponse, err error) {
	if request == nil {
		request = NewGetBriefComponentInstanceRequest()
	}
	response = NewGetBriefComponentInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewPublishAppRequest() (request *PublishAppRequest) {
	request = &PublishAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "PublishApp")
	return
}

func NewPublishAppResponse() (response *PublishAppResponse) {
	response = &PublishAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发布/下架应用模板
func (c *Client) PublishApp(request *PublishAppRequest) (response *PublishAppResponse, err error) {
	if request == nil {
		request = NewPublishAppRequest()
	}
	response = NewPublishAppResponse()
	err = c.Send(request, response)
	return
}

func NewSetCanaryTrafficRequest() (request *SetCanaryTrafficRequest) {
	request = &SetCanaryTrafficRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "SetCanaryTraffic")
	return
}

func NewSetCanaryTrafficResponse() (response *SetCanaryTrafficResponse) {
	response = &SetCanaryTrafficResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置灰度发布的流量
func (c *Client) SetCanaryTraffic(request *SetCanaryTrafficRequest) (response *SetCanaryTrafficResponse, err error) {
	if request == nil {
		request = NewSetCanaryTrafficRequest()
	}
	response = NewSetCanaryTrafficResponse()
	err = c.Send(request, response)
	return
}

func NewSaveRevisionFromYamlRequest() (request *SaveRevisionFromYamlRequest) {
	request = &SaveRevisionFromYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "SaveRevisionFromYaml")
	return
}

func NewSaveRevisionFromYamlResponse() (response *SaveRevisionFromYamlResponse) {
	response = &SaveRevisionFromYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过yaml保存revision版本
func (c *Client) SaveRevisionFromYaml(request *SaveRevisionFromYamlRequest) (response *SaveRevisionFromYamlResponse, err error) {
	if request == nil {
		request = NewSaveRevisionFromYamlRequest()
	}
	response = NewSaveRevisionFromYamlResponse()
	err = c.Send(request, response)
	return
}

func NewSetUpgradeConfigRequest() (request *SetUpgradeConfigRequest) {
	request = &SetUpgradeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("alm", APIVersion, "SetUpgradeConfig")
	return
}

func NewSetUpgradeConfigResponse() (response *SetUpgradeConfigResponse) {
	response = &SetUpgradeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置升级配置
func (c *Client) SetUpgradeConfig(request *SetUpgradeConfigRequest) (response *SetUpgradeConfigResponse, err error) {
	if request == nil {
		request = NewSetUpgradeConfigRequest()
	}
	response = NewSetUpgradeConfigResponse()
	err = c.Send(request, response)
	return
}
