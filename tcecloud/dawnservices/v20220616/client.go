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

func NewExportOperationTemplateRequest() (request *ExportOperationTemplateRequest) {
	request = &ExportOperationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ExportOperationTemplate")
	return
}

func NewExportOperationTemplateResponse() (response *ExportOperationTemplateResponse) {
	response = &ExportOperationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出运维模板物料包
func (c *Client) ExportOperationTemplate(request *ExportOperationTemplateRequest) (response *ExportOperationTemplateResponse, err error) {
	if request == nil {
		request = NewExportOperationTemplateRequest()
	}
	response = NewExportOperationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewImportPlanRequest() (request *ImportPlanRequest) {
	request = &ImportPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ImportPlan")
	return
}

func NewImportPlanResponse() (response *ImportPlanResponse) {
	response = &ImportPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入极光规划包, 同时对规划进行初始化
func (c *Client) ImportPlan(request *ImportPlanRequest) (response *ImportPlanResponse, err error) {
	if request == nil {
		request = NewImportPlanRequest()
	}
	response = NewImportPlanResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUploadMaterialTaskRequest() (request *CreateUploadMaterialTaskRequest) {
	request = &CreateUploadMaterialTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CreateUploadMaterialTask")
	return
}

func NewCreateUploadMaterialTaskResponse() (response *CreateUploadMaterialTaskResponse) {
	response = &CreateUploadMaterialTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建上传物料任务
func (c *Client) CreateUploadMaterialTask(request *CreateUploadMaterialTaskRequest) (response *CreateUploadMaterialTaskResponse, err error) {
	if request == nil {
		request = NewCreateUploadMaterialTaskRequest()
	}
	response = NewCreateUploadMaterialTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationTemplateRequest() (request *DescribeOperationTemplateRequest) {
	request = &DescribeOperationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeOperationTemplate")
	return
}

func NewDescribeOperationTemplateResponse() (response *DescribeOperationTemplateResponse) {
	response = &DescribeOperationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取运维模版详细信息
func (c *Client) DescribeOperationTemplate(request *DescribeOperationTemplateRequest) (response *DescribeOperationTemplateResponse, err error) {
	if request == nil {
		request = NewDescribeOperationTemplateRequest()
	}
	response = NewDescribeOperationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationArtifactChartHistoryRequest() (request *GetApplicationArtifactChartHistoryRequest) {
	request = &GetApplicationArtifactChartHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetApplicationArtifactChartHistory")
	return
}

func NewGetApplicationArtifactChartHistoryResponse() (response *GetApplicationArtifactChartHistoryResponse) {
	response = &GetApplicationArtifactChartHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用chart包修改历史
func (c *Client) GetApplicationArtifactChartHistory(request *GetApplicationArtifactChartHistoryRequest) (response *GetApplicationArtifactChartHistoryResponse, err error) {
	if request == nil {
		request = NewGetApplicationArtifactChartHistoryRequest()
	}
	response = NewGetApplicationArtifactChartHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewGetLatestMaterialPathRequest() (request *GetLatestMaterialPathRequest) {
	request = &GetLatestMaterialPathRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetLatestMaterialPath")
	return
}

func NewGetLatestMaterialPathResponse() (response *GetLatestMaterialPathResponse) {
	response = &GetLatestMaterialPathResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取服务器上最新物料路径
func (c *Client) GetLatestMaterialPath(request *GetLatestMaterialPathRequest) (response *GetLatestMaterialPathResponse, err error) {
	if request == nil {
		request = NewGetLatestMaterialPathRequest()
	}
	response = NewGetLatestMaterialPathResponse()
	err = c.Send(request, response)
	return
}

func NewStartDevelopmentPhaseApplicationDeployRequest() (request *StartDevelopmentPhaseApplicationDeployRequest) {
	request = &StartDevelopmentPhaseApplicationDeployRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "StartDevelopmentPhaseApplicationDeploy")
	return
}

func NewStartDevelopmentPhaseApplicationDeployResponse() (response *StartDevelopmentPhaseApplicationDeployResponse) {
	response = &StartDevelopmentPhaseApplicationDeployResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 部署研发阶段的应用, 环境中包含当前应用的规划和物料
func (c *Client) StartDevelopmentPhaseApplicationDeploy(request *StartDevelopmentPhaseApplicationDeployRequest) (response *StartDevelopmentPhaseApplicationDeployResponse, err error) {
	if request == nil {
		request = NewStartDevelopmentPhaseApplicationDeployRequest()
	}
	response = NewStartDevelopmentPhaseApplicationDeployResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMaterialSyncTaskRequest() (request *CreateMaterialSyncTaskRequest) {
	request = &CreateMaterialSyncTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CreateMaterialSyncTask")
	return
}

func NewCreateMaterialSyncTaskResponse() (response *CreateMaterialSyncTaskResponse) {
	response = &CreateMaterialSyncTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建物料同步任务
func (c *Client) CreateMaterialSyncTask(request *CreateMaterialSyncTaskRequest) (response *CreateMaterialSyncTaskResponse, err error) {
	if request == nil {
		request = NewCreateMaterialSyncTaskRequest()
	}
	response = NewCreateMaterialSyncTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeployEnvRequest() (request *DeployEnvRequest) {
	request = &DeployEnvRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DeployEnv")
	return
}

func NewDeployEnvResponse() (response *DeployEnvResponse) {
	response = &DeployEnvResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 部署环境底座
func (c *Client) DeployEnv(request *DeployEnvRequest) (response *DeployEnvResponse, err error) {
	if request == nil {
		request = NewDeployEnvRequest()
	}
	response = NewDeployEnvResponse()
	err = c.Send(request, response)
	return
}

func NewModifySystemConfigRequest() (request *ModifySystemConfigRequest) {
	request = &ModifySystemConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ModifySystemConfig")
	return
}

func NewModifySystemConfigResponse() (response *ModifySystemConfigResponse) {
	response = &ModifySystemConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改系统配置
func (c *Client) ModifySystemConfig(request *ModifySystemConfigRequest) (response *ModifySystemConfigResponse, err error) {
	if request == nil {
		request = NewModifySystemConfigRequest()
	}
	response = NewModifySystemConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductInstanceOperationSheetRequest() (request *CreateProductInstanceOperationSheetRequest) {
	request = &CreateProductInstanceOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CreateProductInstanceOperationSheet")
	return
}

func NewCreateProductInstanceOperationSheetResponse() (response *CreateProductInstanceOperationSheetResponse) {
	response = &CreateProductInstanceOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建产品实例运维单
func (c *Client) CreateProductInstanceOperationSheet(request *CreateProductInstanceOperationSheetRequest) (response *CreateProductInstanceOperationSheetResponse, err error) {
	if request == nil {
		request = NewCreateProductInstanceOperationSheetRequest()
	}
	response = NewCreateProductInstanceOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAppRuntimeNameRequest() (request *ModifyAppRuntimeNameRequest) {
	request = &ModifyAppRuntimeNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ModifyAppRuntimeName")
	return
}

func NewModifyAppRuntimeNameResponse() (response *ModifyAppRuntimeNameResponse) {
	response = &ModifyAppRuntimeNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改应用运行时名称
func (c *Client) ModifyAppRuntimeName(request *ModifyAppRuntimeNameRequest) (response *ModifyAppRuntimeNameResponse, err error) {
	if request == nil {
		request = NewModifyAppRuntimeNameRequest()
	}
	response = NewModifyAppRuntimeNameResponse()
	err = c.Send(request, response)
	return
}

func NewApplySiteMaterialRequest() (request *ApplySiteMaterialRequest) {
	request = &ApplySiteMaterialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ApplySiteMaterial")
	return
}

func NewApplySiteMaterialResponse() (response *ApplySiteMaterialResponse) {
	response = &ApplySiteMaterialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 为局点关联标准物料
func (c *Client) ApplySiteMaterial(request *ApplySiteMaterialRequest) (response *ApplySiteMaterialResponse, err error) {
	if request == nil {
		request = NewApplySiteMaterialRequest()
	}
	response = NewApplySiteMaterialResponse()
	err = c.Send(request, response)
	return
}

func NewGenerateApplicationPipelineRequest() (request *GenerateApplicationPipelineRequest) {
	request = &GenerateApplicationPipelineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GenerateApplicationPipeline")
	return
}

func NewGenerateApplicationPipelineResponse() (response *GenerateApplicationPipelineResponse) {
	response = &GenerateApplicationPipelineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 生成应用运维任务描述
func (c *Client) GenerateApplicationPipeline(request *GenerateApplicationPipelineRequest) (response *GenerateApplicationPipelineResponse, err error) {
	if request == nil {
		request = NewGenerateApplicationPipelineRequest()
	}
	response = NewGenerateApplicationPipelineResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationPlanUUIDRequest() (request *GetApplicationPlanUUIDRequest) {
	request = &GetApplicationPlanUUIDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetApplicationPlanUUID")
	return
}

func NewGetApplicationPlanUUIDResponse() (response *GetApplicationPlanUUIDResponse) {
	response = &GetApplicationPlanUUIDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用规划UUID
func (c *Client) GetApplicationPlanUUID(request *GetApplicationPlanUUIDRequest) (response *GetApplicationPlanUUIDResponse, err error) {
	if request == nil {
		request = NewGetApplicationPlanUUIDRequest()
	}
	response = NewGetApplicationPlanUUIDResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperationSheetRequest() (request *CreateOperationSheetRequest) {
	request = &CreateOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CreateOperationSheet")
	return
}

func NewCreateOperationSheetResponse() (response *CreateOperationSheetResponse) {
	response = &CreateOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建运维单
func (c *Client) CreateOperationSheet(request *CreateOperationSheetRequest) (response *CreateOperationSheetResponse, err error) {
	if request == nil {
		request = NewCreateOperationSheetRequest()
	}
	response = NewCreateOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewListProductCommonEnumRequest() (request *ListProductCommonEnumRequest) {
	request = &ListProductCommonEnumRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListProductCommonEnum")
	return
}

func NewListProductCommonEnumResponse() (response *ListProductCommonEnumResponse) {
	response = &ListProductCommonEnumResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品公共枚举类型
func (c *Client) ListProductCommonEnum(request *ListProductCommonEnumRequest) (response *ListProductCommonEnumResponse, err error) {
	if request == nil {
		request = NewListProductCommonEnumRequest()
	}
	response = NewListProductCommonEnumResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductInfoRequest() (request *GetProductInfoRequest) {
	request = &GetProductInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetProductInfo")
	return
}

func NewGetProductInfoResponse() (response *GetProductInfoResponse) {
	response = &GetProductInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品详情
func (c *Client) GetProductInfo(request *GetProductInfoRequest) (response *GetProductInfoResponse, err error) {
	if request == nil {
		request = NewGetProductInfoRequest()
	}
	response = NewGetProductInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApplicationYamlRequest() (request *DescribeApplicationYamlRequest) {
	request = &DescribeApplicationYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeApplicationYaml")
	return
}

func NewDescribeApplicationYamlResponse() (response *DescribeApplicationYamlResponse) {
	response = &DescribeApplicationYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看应用values.yaml文件
func (c *Client) DescribeApplicationYaml(request *DescribeApplicationYamlRequest) (response *DescribeApplicationYamlResponse, err error) {
	if request == nil {
		request = NewDescribeApplicationYamlRequest()
	}
	response = NewDescribeApplicationYamlResponse()
	err = c.Send(request, response)
	return
}

func NewListClusterRequest() (request *ListClusterRequest) {
	request = &ListClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListCluster")
	return
}

func NewListClusterResponse() (response *ListClusterResponse) {
	response = &ListClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群列表
func (c *Client) ListCluster(request *ListClusterRequest) (response *ListClusterResponse, err error) {
	if request == nil {
		request = NewListClusterRequest()
	}
	response = NewListClusterResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCommonOperationSheetTemplateRequest() (request *UpdateCommonOperationSheetTemplateRequest) {
	request = &UpdateCommonOperationSheetTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "UpdateCommonOperationSheetTemplate")
	return
}

func NewUpdateCommonOperationSheetTemplateResponse() (response *UpdateCommonOperationSheetTemplateResponse) {
	response = &UpdateCommonOperationSheetTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新通用变更单编排内容
func (c *Client) UpdateCommonOperationSheetTemplate(request *UpdateCommonOperationSheetTemplateRequest) (response *UpdateCommonOperationSheetTemplateResponse, err error) {
	if request == nil {
		request = NewUpdateCommonOperationSheetTemplateRequest()
	}
	response = NewUpdateCommonOperationSheetTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewImportOtaCertificateRequest() (request *ImportOtaCertificateRequest) {
	request = &ImportOtaCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ImportOtaCertificate")
	return
}

func NewImportOtaCertificateResponse() (response *ImportOtaCertificateResponse) {
	response = &ImportOtaCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入局点OTA证书
func (c *Client) ImportOtaCertificate(request *ImportOtaCertificateRequest) (response *ImportOtaCertificateResponse, err error) {
	if request == nil {
		request = NewImportOtaCertificateRequest()
	}
	response = NewImportOtaCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewListProductsRequest() (request *ListProductsRequest) {
	request = &ListProductsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListProducts")
	return
}

func NewListProductsResponse() (response *ListProductsResponse) {
	response = &ListProductsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品列表
func (c *Client) ListProducts(request *ListProductsRequest) (response *ListProductsResponse, err error) {
	if request == nil {
		request = NewListProductsRequest()
	}
	response = NewListProductsResponse()
	err = c.Send(request, response)
	return
}

func NewListProductMaterialsRequest() (request *ListProductMaterialsRequest) {
	request = &ListProductMaterialsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListProductMaterials")
	return
}

func NewListProductMaterialsResponse() (response *ListProductMaterialsResponse) {
	response = &ListProductMaterialsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品物料信息
func (c *Client) ListProductMaterials(request *ListProductMaterialsRequest) (response *ListProductMaterialsResponse, err error) {
	if request == nil {
		request = NewListProductMaterialsRequest()
	}
	response = NewListProductMaterialsResponse()
	err = c.Send(request, response)
	return
}

func NewListAtomicToPipelineRequest() (request *ListAtomicToPipelineRequest) {
	request = &ListAtomicToPipelineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListAtomicToPipeline")
	return
}

func NewListAtomicToPipelineResponse() (response *ListAtomicToPipelineResponse) {
	response = &ListAtomicToPipelineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 请求原子操作的名称版本列表，导出原子操作pipeline yaml文件列表
func (c *Client) ListAtomicToPipeline(request *ListAtomicToPipelineRequest) (response *ListAtomicToPipelineResponse, err error) {
	if request == nil {
		request = NewListAtomicToPipelineRequest()
	}
	response = NewListAtomicToPipelineResponse()
	err = c.Send(request, response)
	return
}

func NewListProductPlanRequest() (request *ListProductPlanRequest) {
	request = &ListProductPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListProductPlan")
	return
}

func NewListProductPlanResponse() (response *ListProductPlanResponse) {
	response = &ListProductPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据请求参数, 查询指定产品或产品实例的规划信息
func (c *Client) ListProductPlan(request *ListProductPlanRequest) (response *ListProductPlanResponse, err error) {
	if request == nil {
		request = NewListProductPlanRequest()
	}
	response = NewListProductPlanResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSheetTemplateRequest() (request *DescribeSheetTemplateRequest) {
	request = &DescribeSheetTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeSheetTemplate")
	return
}

func NewDescribeSheetTemplateResponse() (response *DescribeSheetTemplateResponse) {
	response = &DescribeSheetTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询实例运维单编排模板
func (c *Client) DescribeSheetTemplate(request *DescribeSheetTemplateRequest) (response *DescribeSheetTemplateResponse, err error) {
	if request == nil {
		request = NewDescribeSheetTemplateRequest()
	}
	response = NewDescribeSheetTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePlanTagRequest() (request *CreatePlanTagRequest) {
	request = &CreatePlanTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CreatePlanTag")
	return
}

func NewCreatePlanTagResponse() (response *CreatePlanTagResponse) {
	response = &CreatePlanTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 给规划打tag,关联当前时刻之前的所有修改记录
func (c *Client) CreatePlanTag(request *CreatePlanTagRequest) (response *CreatePlanTagResponse, err error) {
	if request == nil {
		request = NewCreatePlanTagRequest()
	}
	response = NewCreatePlanTagResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePlanHistoryRequest() (request *DescribePlanHistoryRequest) {
	request = &DescribePlanHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribePlanHistory")
	return
}

func NewDescribePlanHistoryResponse() (response *DescribePlanHistoryResponse) {
	response = &DescribePlanHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某次规划修改详情
func (c *Client) DescribePlanHistory(request *DescribePlanHistoryRequest) (response *DescribePlanHistoryResponse, err error) {
	if request == nil {
		request = NewDescribePlanHistoryRequest()
	}
	response = NewDescribePlanHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewListProductInstanceOperationSheetsRequest() (request *ListProductInstanceOperationSheetsRequest) {
	request = &ListProductInstanceOperationSheetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListProductInstanceOperationSheets")
	return
}

func NewListProductInstanceOperationSheetsResponse() (response *ListProductInstanceOperationSheetsResponse) {
	response = &ListProductInstanceOperationSheetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品实例运维单列表
func (c *Client) ListProductInstanceOperationSheets(request *ListProductInstanceOperationSheetsRequest) (response *ListProductInstanceOperationSheetsResponse, err error) {
	if request == nil {
		request = NewListProductInstanceOperationSheetsRequest()
	}
	response = NewListProductInstanceOperationSheetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApprovalHistoryRequest() (request *DescribeApprovalHistoryRequest) {
	request = &DescribeApprovalHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeApprovalHistory")
	return
}

func NewDescribeApprovalHistoryResponse() (response *DescribeApprovalHistoryResponse) {
	response = &DescribeApprovalHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询变更单审批历史
func (c *Client) DescribeApprovalHistory(request *DescribeApprovalHistoryRequest) (response *DescribeApprovalHistoryResponse, err error) {
	if request == nil {
		request = NewDescribeApprovalHistoryRequest()
	}
	response = NewDescribeApprovalHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSystemConfigRequest() (request *DescribeSystemConfigRequest) {
	request = &DescribeSystemConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeSystemConfig")
	return
}

func NewDescribeSystemConfigResponse() (response *DescribeSystemConfigResponse) {
	response = &DescribeSystemConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询系统配置
func (c *Client) DescribeSystemConfig(request *DescribeSystemConfigRequest) (response *DescribeSystemConfigResponse, err error) {
	if request == nil {
		request = NewDescribeSystemConfigRequest()
	}
	response = NewDescribeSystemConfigResponse()
	err = c.Send(request, response)
	return
}

func NewListSolutionTemplatesRequest() (request *ListSolutionTemplatesRequest) {
	request = &ListSolutionTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListSolutionTemplates")
	return
}

func NewListSolutionTemplatesResponse() (response *ListSolutionTemplatesResponse) {
	response = &ListSolutionTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询解决方案运维模版
func (c *Client) ListSolutionTemplates(request *ListSolutionTemplatesRequest) (response *ListSolutionTemplatesResponse, err error) {
	if request == nil {
		request = NewListSolutionTemplatesRequest()
	}
	response = NewListSolutionTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewCopyProductInstanceOperationSheetRequest() (request *CopyProductInstanceOperationSheetRequest) {
	request = &CopyProductInstanceOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CopyProductInstanceOperationSheet")
	return
}

func NewCopyProductInstanceOperationSheetResponse() (response *CopyProductInstanceOperationSheetResponse) {
	response = &CopyProductInstanceOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 复制产品实例运维单
func (c *Client) CopyProductInstanceOperationSheet(request *CopyProductInstanceOperationSheetRequest) (response *CopyProductInstanceOperationSheetResponse, err error) {
	if request == nil {
		request = NewCopyProductInstanceOperationSheetRequest()
	}
	response = NewCopyProductInstanceOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperationTemplateRequest() (request *CreateOperationTemplateRequest) {
	request = &CreateOperationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CreateOperationTemplate")
	return
}

func NewCreateOperationTemplateResponse() (response *CreateOperationTemplateResponse) {
	response = &CreateOperationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取运维模版列表
func (c *Client) CreateOperationTemplate(request *CreateOperationTemplateRequest) (response *CreateOperationTemplateResponse, err error) {
	if request == nil {
		request = NewCreateOperationTemplateRequest()
	}
	response = NewCreateOperationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewImportProductDataRequest() (request *ImportProductDataRequest) {
	request = &ImportProductDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ImportProductData")
	return
}

func NewImportProductDataResponse() (response *ImportProductDataResponse) {
	response = &ImportProductDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品数据导入
func (c *Client) ImportProductData(request *ImportProductDataRequest) (response *ImportProductDataResponse, err error) {
	if request == nil {
		request = NewImportProductDataRequest()
	}
	response = NewImportProductDataResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSiteOperationSheetRequest() (request *CreateSiteOperationSheetRequest) {
	request = &CreateSiteOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CreateSiteOperationSheet")
	return
}

func NewCreateSiteOperationSheetResponse() (response *CreateSiteOperationSheetResponse) {
	response = &CreateSiteOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建局点变更单
func (c *Client) CreateSiteOperationSheet(request *CreateSiteOperationSheetRequest) (response *CreateSiteOperationSheetResponse, err error) {
	if request == nil {
		request = NewCreateSiteOperationSheetRequest()
	}
	response = NewCreateSiteOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewImportApplicationPlanRequest() (request *ImportApplicationPlanRequest) {
	request = &ImportApplicationPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ImportApplicationPlan")
	return
}

func NewImportApplicationPlanResponse() (response *ImportApplicationPlanResponse) {
	response = &ImportApplicationPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入单个应用的规划
func (c *Client) ImportApplicationPlan(request *ImportApplicationPlanRequest) (response *ImportApplicationPlanResponse, err error) {
	if request == nil {
		request = NewImportApplicationPlanRequest()
	}
	response = NewImportApplicationPlanResponse()
	err = c.Send(request, response)
	return
}

func NewListOperationSheetInstancesRequest() (request *ListOperationSheetInstancesRequest) {
	request = &ListOperationSheetInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListOperationSheetInstances")
	return
}

func NewListOperationSheetInstancesResponse() (response *ListOperationSheetInstancesResponse) {
	response = &ListOperationSheetInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询变更单实例列表
func (c *Client) ListOperationSheetInstances(request *ListOperationSheetInstancesRequest) (response *ListOperationSheetInstancesResponse, err error) {
	if request == nil {
		request = NewListOperationSheetInstancesRequest()
	}
	response = NewListOperationSheetInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSheetTemplateRequest() (request *UpdateSheetTemplateRequest) {
	request = &UpdateSheetTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "UpdateSheetTemplate")
	return
}

func NewUpdateSheetTemplateResponse() (response *UpdateSheetTemplateResponse) {
	response = &UpdateSheetTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新实例运维单编排模板
func (c *Client) UpdateSheetTemplate(request *UpdateSheetTemplateRequest) (response *UpdateSheetTemplateResponse, err error) {
	if request == nil {
		request = NewUpdateSheetTemplateRequest()
	}
	response = NewUpdateSheetTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewImportProductInstanceOperationSheetRequest() (request *ImportProductInstanceOperationSheetRequest) {
	request = &ImportProductInstanceOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ImportProductInstanceOperationSheet")
	return
}

func NewImportProductInstanceOperationSheetResponse() (response *ImportProductInstanceOperationSheetResponse) {
	response = &ImportProductInstanceOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入产品实例运维单
func (c *Client) ImportProductInstanceOperationSheet(request *ImportProductInstanceOperationSheetRequest) (response *ImportProductInstanceOperationSheetResponse, err error) {
	if request == nil {
		request = NewImportProductInstanceOperationSheetRequest()
	}
	response = NewImportProductInstanceOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewListLatestApplicationPackagesRequest() (request *ListLatestApplicationPackagesRequest) {
	request = &ListLatestApplicationPackagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListLatestApplicationPackages")
	return
}

func NewListLatestApplicationPackagesResponse() (response *ListLatestApplicationPackagesResponse) {
	response = &ListLatestApplicationPackagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出最新的应用包信息，用于部署场景
func (c *Client) ListLatestApplicationPackages(request *ListLatestApplicationPackagesRequest) (response *ListLatestApplicationPackagesResponse, err error) {
	if request == nil {
		request = NewListLatestApplicationPackagesRequest()
	}
	response = NewListLatestApplicationPackagesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOtaCheckConfigRequest() (request *UpdateOtaCheckConfigRequest) {
	request = &UpdateOtaCheckConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "UpdateOtaCheckConfig")
	return
}

func NewUpdateOtaCheckConfigResponse() (response *UpdateOtaCheckConfigResponse) {
	response = &UpdateOtaCheckConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物料自动同步设置
func (c *Client) UpdateOtaCheckConfig(request *UpdateOtaCheckConfigRequest) (response *UpdateOtaCheckConfigResponse, err error) {
	if request == nil {
		request = NewUpdateOtaCheckConfigRequest()
	}
	response = NewUpdateOtaCheckConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSheetAttachmentsRequest() (request *DescribeSheetAttachmentsRequest) {
	request = &DescribeSheetAttachmentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeSheetAttachments")
	return
}

func NewDescribeSheetAttachmentsResponse() (response *DescribeSheetAttachmentsResponse) {
	response = &DescribeSheetAttachmentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询变更单的关联附件列表
func (c *Client) DescribeSheetAttachments(request *DescribeSheetAttachmentsRequest) (response *DescribeSheetAttachmentsResponse, err error) {
	if request == nil {
		request = NewDescribeSheetAttachmentsRequest()
	}
	response = NewDescribeSheetAttachmentsResponse()
	err = c.Send(request, response)
	return
}

func NewListTaskRequest() (request *ListTaskRequest) {
	request = &ListTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListTask")
	return
}

func NewListTaskResponse() (response *ListTaskResponse) {
	response = &ListTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回创建原子操作可用的task类型和每个类型对应的入参结构
func (c *Client) ListTask(request *ListTaskRequest) (response *ListTaskResponse, err error) {
	if request == nil {
		request = NewListTaskRequest()
	}
	response = NewListTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppRuntimeConfigRequest() (request *DescribeAppRuntimeConfigRequest) {
	request = &DescribeAppRuntimeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeAppRuntimeConfig")
	return
}

func NewDescribeAppRuntimeConfigResponse() (response *DescribeAppRuntimeConfigResponse) {
	response = &DescribeAppRuntimeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用运行态的应用配置
func (c *Client) DescribeAppRuntimeConfig(request *DescribeAppRuntimeConfigRequest) (response *DescribeAppRuntimeConfigResponse, err error) {
	if request == nil {
		request = NewDescribeAppRuntimeConfigRequest()
	}
	response = NewDescribeAppRuntimeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAtomicRequest() (request *UpdateAtomicRequest) {
	request = &UpdateAtomicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "UpdateAtomic")
	return
}

func NewUpdateAtomicResponse() (response *UpdateAtomicResponse) {
	response = &UpdateAtomicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 请求原子操作名称&Version，传递原子操作的内容定义，修改原子操作定义
func (c *Client) UpdateAtomic(request *UpdateAtomicRequest) (response *UpdateAtomicResponse, err error) {
	if request == nil {
		request = NewUpdateAtomicRequest()
	}
	response = NewUpdateAtomicResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductRuntimeConfigRequest() (request *DescribeProductRuntimeConfigRequest) {
	request = &DescribeProductRuntimeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeProductRuntimeConfig")
	return
}

func NewDescribeProductRuntimeConfigResponse() (response *DescribeProductRuntimeConfigResponse) {
	response = &DescribeProductRuntimeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品实例运行态的全局配置
func (c *Client) DescribeProductRuntimeConfig(request *DescribeProductRuntimeConfigRequest) (response *DescribeProductRuntimeConfigResponse, err error) {
	if request == nil {
		request = NewDescribeProductRuntimeConfigRequest()
	}
	response = NewDescribeProductRuntimeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOperationSheetRequest() (request *UpdateOperationSheetRequest) {
	request = &UpdateOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "UpdateOperationSheet")
	return
}

func NewUpdateOperationSheetResponse() (response *UpdateOperationSheetResponse) {
	response = &UpdateOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改运维单
func (c *Client) UpdateOperationSheet(request *UpdateOperationSheetRequest) (response *UpdateOperationSheetResponse, err error) {
	if request == nil {
		request = NewUpdateOperationSheetRequest()
	}
	response = NewUpdateOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewPauseNodeRequest() (request *PauseNodeRequest) {
	request = &PauseNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "PauseNode")
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

func NewSubmitOperationSheetApprovalRequest() (request *SubmitOperationSheetApprovalRequest) {
	request = &SubmitOperationSheetApprovalRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "SubmitOperationSheetApproval")
	return
}

func NewSubmitOperationSheetApprovalResponse() (response *SubmitOperationSheetApprovalResponse) {
	response = &SubmitOperationSheetApprovalResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提交变更单审批
func (c *Client) SubmitOperationSheetApproval(request *SubmitOperationSheetApprovalRequest) (response *SubmitOperationSheetApprovalResponse, err error) {
	if request == nil {
		request = NewSubmitOperationSheetApprovalRequest()
	}
	response = NewSubmitOperationSheetApprovalResponse()
	err = c.Send(request, response)
	return
}

func NewParseMaterialRequest() (request *ParseMaterialRequest) {
	request = &ParseMaterialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ParseMaterial")
	return
}

func NewParseMaterialResponse() (response *ParseMaterialResponse) {
	response = &ParseMaterialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物料解析
func (c *Client) ParseMaterial(request *ParseMaterialRequest) (response *ParseMaterialResponse, err error) {
	if request == nil {
		request = NewParseMaterialRequest()
	}
	response = NewParseMaterialResponse()
	err = c.Send(request, response)
	return
}

func NewMatchOperationTemplateByApplicationRequest() (request *MatchOperationTemplateByApplicationRequest) {
	request = &MatchOperationTemplateByApplicationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "MatchOperationTemplateByApplication")
	return
}

func NewMatchOperationTemplateByApplicationResponse() (response *MatchOperationTemplateByApplicationResponse) {
	response = &MatchOperationTemplateByApplicationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 匹配应用模板
func (c *Client) MatchOperationTemplateByApplication(request *MatchOperationTemplateByApplicationRequest) (response *MatchOperationTemplateByApplicationResponse, err error) {
	if request == nil {
		request = NewMatchOperationTemplateByApplicationRequest()
	}
	response = NewMatchOperationTemplateByApplicationResponse()
	err = c.Send(request, response)
	return
}

func NewRerunNodeRequest() (request *RerunNodeRequest) {
	request = &RerunNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "RerunNode")
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

func NewGetApplicationValuesRequest() (request *GetApplicationValuesRequest) {
	request = &GetApplicationValuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetApplicationValues")
	return
}

func NewGetApplicationValuesResponse() (response *GetApplicationValuesResponse) {
	response = &GetApplicationValuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看应用版本values
func (c *Client) GetApplicationValues(request *GetApplicationValuesRequest) (response *GetApplicationValuesResponse, err error) {
	if request == nil {
		request = NewGetApplicationValuesRequest()
	}
	response = NewGetApplicationValuesResponse()
	err = c.Send(request, response)
	return
}

func NewListSiteMaterialsRequest() (request *ListSiteMaterialsRequest) {
	request = &ListSiteMaterialsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListSiteMaterials")
	return
}

func NewListSiteMaterialsResponse() (response *ListSiteMaterialsResponse) {
	response = &ListSiteMaterialsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据SiteUUID获取所有导入任务信息，和导入物料信息
func (c *Client) ListSiteMaterials(request *ListSiteMaterialsRequest) (response *ListSiteMaterialsResponse, err error) {
	if request == nil {
		request = NewListSiteMaterialsRequest()
	}
	response = NewListSiteMaterialsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOtaDownloadConfigRequest() (request *UpdateOtaDownloadConfigRequest) {
	request = &UpdateOtaDownloadConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "UpdateOtaDownloadConfig")
	return
}

func NewUpdateOtaDownloadConfigResponse() (response *UpdateOtaDownloadConfigResponse) {
	response = &UpdateOtaDownloadConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物料自动下载设置
func (c *Client) UpdateOtaDownloadConfig(request *UpdateOtaDownloadConfigRequest) (response *UpdateOtaDownloadConfigResponse, err error) {
	if request == nil {
		request = NewUpdateOtaDownloadConfigRequest()
	}
	response = NewUpdateOtaDownloadConfigResponse()
	err = c.Send(request, response)
	return
}

func NewImportCommonOperationSheetRequest() (request *ImportCommonOperationSheetRequest) {
	request = &ImportCommonOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ImportCommonOperationSheet")
	return
}

func NewImportCommonOperationSheetResponse() (response *ImportCommonOperationSheetResponse) {
	response = &ImportCommonOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入通用变更单
func (c *Client) ImportCommonOperationSheet(request *ImportCommonOperationSheetRequest) (response *ImportCommonOperationSheetResponse, err error) {
	if request == nil {
		request = NewImportCommonOperationSheetRequest()
	}
	response = NewImportCommonOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProductTemplateRequest() (request *ModifyProductTemplateRequest) {
	request = &ModifyProductTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ModifyProductTemplate")
	return
}

func NewModifyProductTemplateResponse() (response *ModifyProductTemplateResponse) {
	response = &ModifyProductTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改产品模板
func (c *Client) ModifyProductTemplate(request *ModifyProductTemplateRequest) (response *ModifyProductTemplateResponse, err error) {
	if request == nil {
		request = NewModifyProductTemplateRequest()
	}
	response = NewModifyProductTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEnvRequest() (request *CreateEnvRequest) {
	request = &CreateEnvRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CreateEnv")
	return
}

func NewCreateEnvResponse() (response *CreateEnvResponse) {
	response = &CreateEnvResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建环境
func (c *Client) CreateEnv(request *CreateEnvRequest) (response *CreateEnvResponse, err error) {
	if request == nil {
		request = NewCreateEnvRequest()
	}
	response = NewCreateEnvResponse()
	err = c.Send(request, response)
	return
}

func NewPublishProductInstanceOperationSheetRequest() (request *PublishProductInstanceOperationSheetRequest) {
	request = &PublishProductInstanceOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "PublishProductInstanceOperationSheet")
	return
}

func NewPublishProductInstanceOperationSheetResponse() (response *PublishProductInstanceOperationSheetResponse) {
	response = &PublishProductInstanceOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发布产品实例运维单
func (c *Client) PublishProductInstanceOperationSheet(request *PublishProductInstanceOperationSheetRequest) (response *PublishProductInstanceOperationSheetResponse, err error) {
	if request == nil {
		request = NewPublishProductInstanceOperationSheetRequest()
	}
	response = NewPublishProductInstanceOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewStartApplicationDeployRequest() (request *StartApplicationDeployRequest) {
	request = &StartApplicationDeployRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "StartApplicationDeploy")
	return
}

func NewStartApplicationDeployResponse() (response *StartApplicationDeployResponse) {
	response = &StartApplicationDeployResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 部署研发阶段的应用
func (c *Client) StartApplicationDeploy(request *StartApplicationDeployRequest) (response *StartApplicationDeployResponse, err error) {
	if request == nil {
		request = NewStartApplicationDeployRequest()
	}
	response = NewStartApplicationDeployResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCommonOperationSheetRequest() (request *DescribeCommonOperationSheetRequest) {
	request = &DescribeCommonOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeCommonOperationSheet")
	return
}

func NewDescribeCommonOperationSheetResponse() (response *DescribeCommonOperationSheetResponse) {
	response = &DescribeCommonOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询通用变更单详情
func (c *Client) DescribeCommonOperationSheet(request *DescribeCommonOperationSheetRequest) (response *DescribeCommonOperationSheetResponse, err error) {
	if request == nil {
		request = NewDescribeCommonOperationSheetRequest()
	}
	response = NewDescribeCommonOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewListSiteOperationSheetsRequest() (request *ListSiteOperationSheetsRequest) {
	request = &ListSiteOperationSheetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListSiteOperationSheets")
	return
}

func NewListSiteOperationSheetsResponse() (response *ListSiteOperationSheetsResponse) {
	response = &ListSiteOperationSheetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询局点变更单列表
func (c *Client) ListSiteOperationSheets(request *ListSiteOperationSheetsRequest) (response *ListSiteOperationSheetsResponse, err error) {
	if request == nil {
		request = NewListSiteOperationSheetsRequest()
	}
	response = NewListSiteOperationSheetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJobDAGRequest() (request *DescribeJobDAGRequest) {
	request = &DescribeJobDAGRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeJobDAG")
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

func NewDeployPlanProductRequest() (request *DeployPlanProductRequest) {
	request = &DeployPlanProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DeployPlanProduct")
	return
}

func NewDeployPlanProductResponse() (response *DeployPlanProductResponse) {
	response = &DeployPlanProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 部署规划包中的产品
func (c *Client) DeployPlanProduct(request *DeployPlanProductRequest) (response *DeployPlanProductResponse, err error) {
	if request == nil {
		request = NewDeployPlanProductRequest()
	}
	response = NewDeployPlanProductResponse()
	err = c.Send(request, response)
	return
}

func NewResumeJobRequest() (request *ResumeJobRequest) {
	request = &ResumeJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ResumeJob")
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

func NewSkipNodeRequest() (request *SkipNodeRequest) {
	request = &SkipNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "SkipNode")
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

func NewCreateCommonOperationSheetRequest() (request *CreateCommonOperationSheetRequest) {
	request = &CreateCommonOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CreateCommonOperationSheet")
	return
}

func NewCreateCommonOperationSheetResponse() (response *CreateCommonOperationSheetResponse) {
	response = &CreateCommonOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建通用变更单详情
func (c *Client) CreateCommonOperationSheet(request *CreateCommonOperationSheetRequest) (response *CreateCommonOperationSheetResponse, err error) {
	if request == nil {
		request = NewCreateCommonOperationSheetRequest()
	}
	response = NewCreateCommonOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewRestartOPMControllerRequest() (request *RestartOPMControllerRequest) {
	request = &RestartOPMControllerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "RestartOPMController")
	return
}

func NewRestartOPMControllerResponse() (response *RestartOPMControllerResponse) {
	response = &RestartOPMControllerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重启OPM Controller
func (c *Client) RestartOPMController(request *RestartOPMControllerRequest) (response *RestartOPMControllerResponse, err error) {
	if request == nil {
		request = NewRestartOPMControllerRequest()
	}
	response = NewRestartOPMControllerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePlanRequest() (request *DescribePlanRequest) {
	request = &DescribePlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribePlan")
	return
}

func NewDescribePlanResponse() (response *DescribePlanResponse) {
	response = &DescribePlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定规划版本信息
func (c *Client) DescribePlan(request *DescribePlanRequest) (response *DescribePlanResponse, err error) {
	if request == nil {
		request = NewDescribePlanRequest()
	}
	response = NewDescribePlanResponse()
	err = c.Send(request, response)
	return
}

func NewCopyOperationTemplateRequest() (request *CopyOperationTemplateRequest) {
	request = &CopyOperationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CopyOperationTemplate")
	return
}

func NewCopyOperationTemplateResponse() (response *CopyOperationTemplateResponse) {
	response = &CopyOperationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 复制运维模板
func (c *Client) CopyOperationTemplate(request *CopyOperationTemplateRequest) (response *CopyOperationTemplateResponse, err error) {
	if request == nil {
		request = NewCopyOperationTemplateRequest()
	}
	response = NewCopyOperationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductInstanceInfoRequest() (request *DescribeProductInstanceInfoRequest) {
	request = &DescribeProductInstanceInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeProductInstanceInfo")
	return
}

func NewDescribeProductInstanceInfoResponse() (response *DescribeProductInstanceInfoResponse) {
	response = &DescribeProductInstanceInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品实例基本信息
func (c *Client) DescribeProductInstanceInfo(request *DescribeProductInstanceInfoRequest) (response *DescribeProductInstanceInfoResponse, err error) {
	if request == nil {
		request = NewDescribeProductInstanceInfoRequest()
	}
	response = NewDescribeProductInstanceInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGenerateOperationSheetInstanceRequest() (request *GenerateOperationSheetInstanceRequest) {
	request = &GenerateOperationSheetInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GenerateOperationSheetInstance")
	return
}

func NewGenerateOperationSheetInstanceResponse() (response *GenerateOperationSheetInstanceResponse) {
	response = &GenerateOperationSheetInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 生成变更单实例
func (c *Client) GenerateOperationSheetInstance(request *GenerateOperationSheetInstanceRequest) (response *GenerateOperationSheetInstanceResponse, err error) {
	if request == nil {
		request = NewGenerateOperationSheetInstanceRequest()
	}
	response = NewGenerateOperationSheetInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialProductSubsystemInstanceRequest() (request *ListMaterialProductSubsystemInstanceRequest) {
	request = &ListMaterialProductSubsystemInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListMaterialProductSubsystemInstance")
	return
}

func NewListMaterialProductSubsystemInstanceResponse() (response *ListMaterialProductSubsystemInstanceResponse) {
	response = &ListMaterialProductSubsystemInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取物料交付方案子系统实例列表
func (c *Client) ListMaterialProductSubsystemInstance(request *ListMaterialProductSubsystemInstanceRequest) (response *ListMaterialProductSubsystemInstanceResponse, err error) {
	if request == nil {
		request = NewListMaterialProductSubsystemInstanceRequest()
	}
	response = NewListMaterialProductSubsystemInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewSetInstantiateAnnotationForOperationTemplateRequest() (request *SetInstantiateAnnotationForOperationTemplateRequest) {
	request = &SetInstantiateAnnotationForOperationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "SetInstantiateAnnotationForOperationTemplate")
	return
}

func NewSetInstantiateAnnotationForOperationTemplateResponse() (response *SetInstantiateAnnotationForOperationTemplateResponse) {
	response = &SetInstantiateAnnotationForOperationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 给变更单编排设置实例化Annotation
func (c *Client) SetInstantiateAnnotationForOperationTemplate(request *SetInstantiateAnnotationForOperationTemplateRequest) (response *SetInstantiateAnnotationForOperationTemplateResponse, err error) {
	if request == nil {
		request = NewSetInstantiateAnnotationForOperationTemplateRequest()
	}
	response = NewSetInstantiateAnnotationForOperationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSolutionTemplateRequest() (request *UpdateSolutionTemplateRequest) {
	request = &UpdateSolutionTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "UpdateSolutionTemplate")
	return
}

func NewUpdateSolutionTemplateResponse() (response *UpdateSolutionTemplateResponse) {
	response = &UpdateSolutionTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新解决方案运维模版
func (c *Client) UpdateSolutionTemplate(request *UpdateSolutionTemplateRequest) (response *UpdateSolutionTemplateResponse, err error) {
	if request == nil {
		request = NewUpdateSolutionTemplateRequest()
	}
	response = NewUpdateSolutionTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductOperationTaskRequest() (request *CreateProductOperationTaskRequest) {
	request = &CreateProductOperationTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CreateProductOperationTask")
	return
}

func NewCreateProductOperationTaskResponse() (response *CreateProductOperationTaskResponse) {
	response = &CreateProductOperationTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建产品运维任务
func (c *Client) CreateProductOperationTask(request *CreateProductOperationTaskRequest) (response *CreateProductOperationTaskResponse, err error) {
	if request == nil {
		request = NewCreateProductOperationTaskRequest()
	}
	response = NewCreateProductOperationTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCheckSiteOperationSheetPlanExistenceRequest() (request *CheckSiteOperationSheetPlanExistenceRequest) {
	request = &CheckSiteOperationSheetPlanExistenceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CheckSiteOperationSheetPlanExistence")
	return
}

func NewCheckSiteOperationSheetPlanExistenceResponse() (response *CheckSiteOperationSheetPlanExistenceResponse) {
	response = &CheckSiteOperationSheetPlanExistenceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查Site变更单中应用的规划是否存在
func (c *Client) CheckSiteOperationSheetPlanExistence(request *CheckSiteOperationSheetPlanExistenceRequest) (response *CheckSiteOperationSheetPlanExistenceResponse, err error) {
	if request == nil {
		request = NewCheckSiteOperationSheetPlanExistenceRequest()
	}
	response = NewCheckSiteOperationSheetPlanExistenceResponse()
	err = c.Send(request, response)
	return
}

func NewExportProductInstanceOperationSheetRequest() (request *ExportProductInstanceOperationSheetRequest) {
	request = &ExportProductInstanceOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ExportProductInstanceOperationSheet")
	return
}

func NewExportProductInstanceOperationSheetResponse() (response *ExportProductInstanceOperationSheetResponse) {
	response = &ExportProductInstanceOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出产品实例运维单
func (c *Client) ExportProductInstanceOperationSheet(request *ExportProductInstanceOperationSheetRequest) (response *ExportProductInstanceOperationSheetResponse, err error) {
	if request == nil {
		request = NewExportProductInstanceOperationSheetRequest()
	}
	response = NewExportProductInstanceOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGlobalConfigRequest() (request *DescribeGlobalConfigRequest) {
	request = &DescribeGlobalConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeGlobalConfig")
	return
}

func NewDescribeGlobalConfigResponse() (response *DescribeGlobalConfigResponse) {
	response = &DescribeGlobalConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据请求参数, 查询一朵云的详细信息
func (c *Client) DescribeGlobalConfig(request *DescribeGlobalConfigRequest) (response *DescribeGlobalConfigResponse, err error) {
	if request == nil {
		request = NewDescribeGlobalConfigRequest()
	}
	response = NewDescribeGlobalConfigResponse()
	err = c.Send(request, response)
	return
}

func NewListPlanServerInfosRequest() (request *ListPlanServerInfosRequest) {
	request = &ListPlanServerInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListPlanServerInfos")
	return
}

func NewListPlanServerInfosResponse() (response *ListPlanServerInfosResponse) {
	response = &ListPlanServerInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务器信息
func (c *Client) ListPlanServerInfos(request *ListPlanServerInfosRequest) (response *ListPlanServerInfosResponse, err error) {
	if request == nil {
		request = NewListPlanServerInfosRequest()
	}
	response = NewListPlanServerInfosResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProjectSiteInfoDetailRequest() (request *DescribeProjectSiteInfoDetailRequest) {
	request = &DescribeProjectSiteInfoDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeProjectSiteInfoDetail")
	return
}

func NewDescribeProjectSiteInfoDetailResponse() (response *DescribeProjectSiteInfoDetailResponse) {
	response = &DescribeProjectSiteInfoDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询局点的详细信息
func (c *Client) DescribeProjectSiteInfoDetail(request *DescribeProjectSiteInfoDetailRequest) (response *DescribeProjectSiteInfoDetailResponse, err error) {
	if request == nil {
		request = NewDescribeProjectSiteInfoDetailRequest()
	}
	response = NewDescribeProjectSiteInfoDetailResponse()
	err = c.Send(request, response)
	return
}

func NewGetOperationDataSchemaRequest() (request *GetOperationDataSchemaRequest) {
	request = &GetOperationDataSchemaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetOperationDataSchema")
	return
}

func NewGetOperationDataSchemaResponse() (response *GetOperationDataSchemaResponse) {
	response = &GetOperationDataSchemaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取运维数据格式列表
func (c *Client) GetOperationDataSchema(request *GetOperationDataSchemaRequest) (response *GetOperationDataSchemaResponse, err error) {
	if request == nil {
		request = NewGetOperationDataSchemaRequest()
	}
	response = NewGetOperationDataSchemaResponse()
	err = c.Send(request, response)
	return
}

func NewGetProjectSiteInfoDetailRequest() (request *GetProjectSiteInfoDetailRequest) {
	request = &GetProjectSiteInfoDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetProjectSiteInfoDetail")
	return
}

func NewGetProjectSiteInfoDetailResponse() (response *GetProjectSiteInfoDetailResponse) {
	response = &GetProjectSiteInfoDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取局点详情
func (c *Client) GetProjectSiteInfoDetail(request *GetProjectSiteInfoDetailRequest) (response *GetProjectSiteInfoDetailResponse, err error) {
	if request == nil {
		request = NewGetProjectSiteInfoDetailRequest()
	}
	response = NewGetProjectSiteInfoDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperationTemplateRequest() (request *DeleteOperationTemplateRequest) {
	request = &DeleteOperationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DeleteOperationTemplate")
	return
}

func NewDeleteOperationTemplateResponse() (response *DeleteOperationTemplateResponse) {
	response = &DeleteOperationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除运维模板
func (c *Client) DeleteOperationTemplate(request *DeleteOperationTemplateRequest) (response *DeleteOperationTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteOperationTemplateRequest()
	}
	response = NewDeleteOperationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSheetAppInstancesRequest() (request *UpdateSheetAppInstancesRequest) {
	request = &UpdateSheetAppInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "UpdateSheetAppInstances")
	return
}

func NewUpdateSheetAppInstancesResponse() (response *UpdateSheetAppInstancesResponse) {
	response = &UpdateSheetAppInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新运维单关联应用实例列表
func (c *Client) UpdateSheetAppInstances(request *UpdateSheetAppInstancesRequest) (response *UpdateSheetAppInstancesResponse, err error) {
	if request == nil {
		request = NewUpdateSheetAppInstancesRequest()
	}
	response = NewUpdateSheetAppInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewQueryUploadMaterialTaskRequest() (request *QueryUploadMaterialTaskRequest) {
	request = &QueryUploadMaterialTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "QueryUploadMaterialTask")
	return
}

func NewQueryUploadMaterialTaskResponse() (response *QueryUploadMaterialTaskResponse) {
	response = &QueryUploadMaterialTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询物料上传任务
func (c *Client) QueryUploadMaterialTask(request *QueryUploadMaterialTaskRequest) (response *QueryUploadMaterialTaskResponse, err error) {
	if request == nil {
		request = NewQueryUploadMaterialTaskRequest()
	}
	response = NewQueryUploadMaterialTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCopyCommonOperationSheetRequest() (request *CopyCommonOperationSheetRequest) {
	request = &CopyCommonOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CopyCommonOperationSheet")
	return
}

func NewCopyCommonOperationSheetResponse() (response *CopyCommonOperationSheetResponse) {
	response = &CopyCommonOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 复制通用变更单
func (c *Client) CopyCommonOperationSheet(request *CopyCommonOperationSheetRequest) (response *CopyCommonOperationSheetResponse, err error) {
	if request == nil {
		request = NewCopyCommonOperationSheetRequest()
	}
	response = NewCopyCommonOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewListSiteApplicationsRequest() (request *ListSiteApplicationsRequest) {
	request = &ListSiteApplicationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListSiteApplications")
	return
}

func NewListSiteApplicationsResponse() (response *ListSiteApplicationsResponse) {
	response = &ListSiteApplicationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据SiteUUID获取获取应用列表
func (c *Client) ListSiteApplications(request *ListSiteApplicationsRequest) (response *ListSiteApplicationsResponse, err error) {
	if request == nil {
		request = NewListSiteApplicationsRequest()
	}
	response = NewListSiteApplicationsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGlobalConfigRequest() (request *ModifyGlobalConfigRequest) {
	request = &ModifyGlobalConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ModifyGlobalConfig")
	return
}

func NewModifyGlobalConfigResponse() (response *ModifyGlobalConfigResponse) {
	response = &ModifyGlobalConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改全局配置信息
func (c *Client) ModifyGlobalConfig(request *ModifyGlobalConfigRequest) (response *ModifyGlobalConfigResponse, err error) {
	if request == nil {
		request = NewModifyGlobalConfigRequest()
	}
	response = NewModifyGlobalConfigResponse()
	err = c.Send(request, response)
	return
}

func NewImportStandardMaterialRequest() (request *ImportStandardMaterialRequest) {
	request = &ImportStandardMaterialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ImportStandardMaterial")
	return
}

func NewImportStandardMaterialResponse() (response *ImportStandardMaterialResponse) {
	response = &ImportStandardMaterialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入极光标准物料
func (c *Client) ImportStandardMaterial(request *ImportStandardMaterialRequest) (response *ImportStandardMaterialResponse, err error) {
	if request == nil {
		request = NewImportStandardMaterialRequest()
	}
	response = NewImportStandardMaterialResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSiteRequest() (request *CreateSiteRequest) {
	request = &CreateSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CreateSite")
	return
}

func NewCreateSiteResponse() (response *CreateSiteResponse) {
	response = &CreateSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建局点
func (c *Client) CreateSite(request *CreateSiteRequest) (response *CreateSiteResponse, err error) {
	if request == nil {
		request = NewCreateSiteRequest()
	}
	response = NewCreateSiteResponse()
	err = c.Send(request, response)
	return
}

func NewListCommonOperationSheetAppInstancesRequest() (request *ListCommonOperationSheetAppInstancesRequest) {
	request = &ListCommonOperationSheetAppInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListCommonOperationSheetAppInstances")
	return
}

func NewListCommonOperationSheetAppInstancesResponse() (response *ListCommonOperationSheetAppInstancesResponse) {
	response = &ListCommonOperationSheetAppInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询变更单（非变更项）的产品、子系统、应用列表
func (c *Client) ListCommonOperationSheetAppInstances(request *ListCommonOperationSheetAppInstancesRequest) (response *ListCommonOperationSheetAppInstancesResponse, err error) {
	if request == nil {
		request = NewListCommonOperationSheetAppInstancesRequest()
	}
	response = NewListCommonOperationSheetAppInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppRuntimeInfoRequest() (request *DescribeAppRuntimeInfoRequest) {
	request = &DescribeAppRuntimeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeAppRuntimeInfo")
	return
}

func NewDescribeAppRuntimeInfoResponse() (response *DescribeAppRuntimeInfoResponse) {
	response = &DescribeAppRuntimeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用运行时信息
func (c *Client) DescribeAppRuntimeInfo(request *DescribeAppRuntimeInfoRequest) (response *DescribeAppRuntimeInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAppRuntimeInfoRequest()
	}
	response = NewDescribeAppRuntimeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePlanApplicationTreeRequest() (request *DescribePlanApplicationTreeRequest) {
	request = &DescribePlanApplicationTreeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribePlanApplicationTree")
	return
}

func NewDescribePlanApplicationTreeResponse() (response *DescribePlanApplicationTreeResponse) {
	response = &DescribePlanApplicationTreeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定版本的规划应用树
func (c *Client) DescribePlanApplicationTree(request *DescribePlanApplicationTreeRequest) (response *DescribePlanApplicationTreeResponse, err error) {
	if request == nil {
		request = NewDescribePlanApplicationTreeRequest()
	}
	response = NewDescribePlanApplicationTreeResponse()
	err = c.Send(request, response)
	return
}

func NewListCloudClustersRequest() (request *ListCloudClustersRequest) {
	request = &ListCloudClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListCloudClusters")
	return
}

func NewListCloudClustersResponse() (response *ListCloudClustersResponse) {
	response = &ListCloudClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询K8S集群列表，集群列表中有当前集群部署产品的个数
func (c *Client) ListCloudClusters(request *ListCloudClustersRequest) (response *ListCloudClustersResponse, err error) {
	if request == nil {
		request = NewListCloudClustersRequest()
	}
	response = NewListCloudClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDiffApplicationInstanceConfigRequest() (request *DiffApplicationInstanceConfigRequest) {
	request = &DiffApplicationInstanceConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DiffApplicationInstanceConfig")
	return
}

func NewDiffApplicationInstanceConfigResponse() (response *DiffApplicationInstanceConfigResponse) {
	response = &DiffApplicationInstanceConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查应用规划是否有变化
func (c *Client) DiffApplicationInstanceConfig(request *DiffApplicationInstanceConfigRequest) (response *DiffApplicationInstanceConfigResponse, err error) {
	if request == nil {
		request = NewDiffApplicationInstanceConfigRequest()
	}
	response = NewDiffApplicationInstanceConfigResponse()
	err = c.Send(request, response)
	return
}

func NewExportCommonOperationSheetRequest() (request *ExportCommonOperationSheetRequest) {
	request = &ExportCommonOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ExportCommonOperationSheet")
	return
}

func NewExportCommonOperationSheetResponse() (response *ExportCommonOperationSheetResponse) {
	response = &ExportCommonOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出通用变更单
func (c *Client) ExportCommonOperationSheet(request *ExportCommonOperationSheetRequest) (response *ExportCommonOperationSheetResponse, err error) {
	if request == nil {
		request = NewExportCommonOperationSheetRequest()
	}
	response = NewExportCommonOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewSetApplicationArtifactChartEffectRequest() (request *SetApplicationArtifactChartEffectRequest) {
	request = &SetApplicationArtifactChartEffectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "SetApplicationArtifactChartEffect")
	return
}

func NewSetApplicationArtifactChartEffectResponse() (response *SetApplicationArtifactChartEffectResponse) {
	response = &SetApplicationArtifactChartEffectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置应用历史chart包生效
func (c *Client) SetApplicationArtifactChartEffect(request *SetApplicationArtifactChartEffectRequest) (response *SetApplicationArtifactChartEffectResponse, err error) {
	if request == nil {
		request = NewSetApplicationArtifactChartEffectRequest()
	}
	response = NewSetApplicationArtifactChartEffectResponse()
	err = c.Send(request, response)
	return
}

func NewListAtomicRequest() (request *ListAtomicRequest) {
	request = &ListAtomicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListAtomic")
	return
}

func NewListAtomicResponse() (response *ListAtomicResponse) {
	response = &ListAtomicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据label等标签，返回全量原子操作的内容定义
func (c *Client) ListAtomic(request *ListAtomicRequest) (response *ListAtomicResponse, err error) {
	if request == nil {
		request = NewListAtomicRequest()
	}
	response = NewListAtomicResponse()
	err = c.Send(request, response)
	return
}

func NewListProductVersionsRequest() (request *ListProductVersionsRequest) {
	request = &ListProductVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListProductVersions")
	return
}

func NewListProductVersionsResponse() (response *ListProductVersionsResponse) {
	response = &ListProductVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品版本列表
func (c *Client) ListProductVersions(request *ListProductVersionsRequest) (response *ListProductVersionsResponse, err error) {
	if request == nil {
		request = NewListProductVersionsRequest()
	}
	response = NewListProductVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewTerminateJobRequest() (request *TerminateJobRequest) {
	request = &TerminateJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "TerminateJob")
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

func NewCheckOperationTemplateRequest() (request *CheckOperationTemplateRequest) {
	request = &CheckOperationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CheckOperationTemplate")
	return
}

func NewCheckOperationTemplateResponse() (response *CheckOperationTemplateResponse) {
	response = &CheckOperationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运维模版校验
func (c *Client) CheckOperationTemplate(request *CheckOperationTemplateRequest) (response *CheckOperationTemplateResponse, err error) {
	if request == nil {
		request = NewCheckOperationTemplateRequest()
	}
	response = NewCheckOperationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialApplicationRequest() (request *ListMaterialApplicationRequest) {
	request = &ListMaterialApplicationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListMaterialApplication")
	return
}

func NewListMaterialApplicationResponse() (response *ListMaterialApplicationResponse) {
	response = &ListMaterialApplicationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取物料应用列表
func (c *Client) ListMaterialApplication(request *ListMaterialApplicationRequest) (response *ListMaterialApplicationResponse, err error) {
	if request == nil {
		request = NewListMaterialApplicationRequest()
	}
	response = NewListMaterialApplicationResponse()
	err = c.Send(request, response)
	return
}

func NewListPublicMaterialSolutionVersionsRequest() (request *ListPublicMaterialSolutionVersionsRequest) {
	request = &ListPublicMaterialSolutionVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListPublicMaterialSolutionVersions")
	return
}

func NewListPublicMaterialSolutionVersionsResponse() (response *ListPublicMaterialSolutionVersionsResponse) {
	response = &ListPublicMaterialSolutionVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取公共市场解决方案版本列表，可用于多云中选择物料
func (c *Client) ListPublicMaterialSolutionVersions(request *ListPublicMaterialSolutionVersionsRequest) (response *ListPublicMaterialSolutionVersionsResponse, err error) {
	if request == nil {
		request = NewListPublicMaterialSolutionVersionsRequest()
	}
	response = NewListPublicMaterialSolutionVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewListChildOperationSheetsRequest() (request *ListChildOperationSheetsRequest) {
	request = &ListChildOperationSheetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListChildOperationSheets")
	return
}

func NewListChildOperationSheetsResponse() (response *ListChildOperationSheetsResponse) {
	response = &ListChildOperationSheetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询子变更单列表
func (c *Client) ListChildOperationSheets(request *ListChildOperationSheetsRequest) (response *ListChildOperationSheetsResponse, err error) {
	if request == nil {
		request = NewListChildOperationSheetsRequest()
	}
	response = NewListChildOperationSheetsResponse()
	err = c.Send(request, response)
	return
}

func NewGetPackageManagerHostRequest() (request *GetPackageManagerHostRequest) {
	request = &GetPackageManagerHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetPackageManagerHost")
	return
}

func NewGetPackageManagerHostResponse() (response *GetPackageManagerHostResponse) {
	response = &GetPackageManagerHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取包管理服务前端使用的域名
func (c *Client) GetPackageManagerHost(request *GetPackageManagerHostRequest) (response *GetPackageManagerHostResponse, err error) {
	if request == nil {
		request = NewGetPackageManagerHostRequest()
	}
	response = NewGetPackageManagerHostResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationTemplateIDRequest() (request *GetApplicationTemplateIDRequest) {
	request = &GetApplicationTemplateIDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetApplicationTemplateID")
	return
}

func NewGetApplicationTemplateIDResponse() (response *GetApplicationTemplateIDResponse) {
	response = &GetApplicationTemplateIDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用部署模版
func (c *Client) GetApplicationTemplateID(request *GetApplicationTemplateIDRequest) (response *GetApplicationTemplateIDResponse, err error) {
	if request == nil {
		request = NewGetApplicationTemplateIDRequest()
	}
	response = NewGetApplicationTemplateIDResponse()
	err = c.Send(request, response)
	return
}

func NewListProductInstanceAppsRequest() (request *ListProductInstanceAppsRequest) {
	request = &ListProductInstanceAppsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListProductInstanceApps")
	return
}

func NewListProductInstanceAppsResponse() (response *ListProductInstanceAppsResponse) {
	response = &ListProductInstanceAppsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品实例应用列表
func (c *Client) ListProductInstanceApps(request *ListProductInstanceAppsRequest) (response *ListProductInstanceAppsResponse, err error) {
	if request == nil {
		request = NewListProductInstanceAppsRequest()
	}
	response = NewListProductInstanceAppsResponse()
	err = c.Send(request, response)
	return
}

func NewListApplicationInstancesByApplicationNameRequest() (request *ListApplicationInstancesByApplicationNameRequest) {
	request = &ListApplicationInstancesByApplicationNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListApplicationInstancesByApplicationName")
	return
}

func NewListApplicationInstancesByApplicationNameResponse() (response *ListApplicationInstancesByApplicationNameResponse) {
	response = &ListApplicationInstancesByApplicationNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据应用名称获取到当前生效局点下的所有应用实例
func (c *Client) ListApplicationInstancesByApplicationName(request *ListApplicationInstancesByApplicationNameRequest) (response *ListApplicationInstancesByApplicationNameResponse, err error) {
	if request == nil {
		request = NewListApplicationInstancesByApplicationNameRequest()
	}
	response = NewListApplicationInstancesByApplicationNameResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProjectSiteRequest() (request *DeleteProjectSiteRequest) {
	request = &DeleteProjectSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DeleteProjectSite")
	return
}

func NewDeleteProjectSiteResponse() (response *DeleteProjectSiteResponse) {
	response = &DeleteProjectSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除局点
func (c *Client) DeleteProjectSite(request *DeleteProjectSiteRequest) (response *DeleteProjectSiteResponse, err error) {
	if request == nil {
		request = NewDeleteProjectSiteRequest()
	}
	response = NewDeleteProjectSiteResponse()
	err = c.Send(request, response)
	return
}

func NewStartOperationSheetRequest() (request *StartOperationSheetRequest) {
	request = &StartOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "StartOperationSheet")
	return
}

func NewStartOperationSheetResponse() (response *StartOperationSheetResponse) {
	response = &StartOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动运维单
func (c *Client) StartOperationSheet(request *StartOperationSheetRequest) (response *StartOperationSheetResponse, err error) {
	if request == nil {
		request = NewStartOperationSheetRequest()
	}
	response = NewStartOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewCheckMaterialRequest() (request *CheckMaterialRequest) {
	request = &CheckMaterialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CheckMaterial")
	return
}

func NewCheckMaterialResponse() (response *CheckMaterialResponse) {
	response = &CheckMaterialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物料检查更新
func (c *Client) CheckMaterial(request *CheckMaterialRequest) (response *CheckMaterialResponse, err error) {
	if request == nil {
		request = NewCheckMaterialRequest()
	}
	response = NewCheckMaterialResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialOperationSheetRequest() (request *ListMaterialOperationSheetRequest) {
	request = &ListMaterialOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListMaterialOperationSheet")
	return
}

func NewListMaterialOperationSheetResponse() (response *ListMaterialOperationSheetResponse) {
	response = &ListMaterialOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取物料包中的变更单信息
func (c *Client) ListMaterialOperationSheet(request *ListMaterialOperationSheetRequest) (response *ListMaterialOperationSheetResponse, err error) {
	if request == nil {
		request = NewListMaterialOperationSheetRequest()
	}
	response = NewListMaterialOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewListProductApplicationsRequest() (request *ListProductApplicationsRequest) {
	request = &ListProductApplicationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListProductApplications")
	return
}

func NewListProductApplicationsResponse() (response *ListProductApplicationsResponse) {
	response = &ListProductApplicationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品版本下应用列表
func (c *Client) ListProductApplications(request *ListProductApplicationsRequest) (response *ListProductApplicationsResponse, err error) {
	if request == nil {
		request = NewListProductApplicationsRequest()
	}
	response = NewListProductApplicationsResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationArtifactChartFileRequest() (request *GetApplicationArtifactChartFileRequest) {
	request = &GetApplicationArtifactChartFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetApplicationArtifactChartFile")
	return
}

func NewGetApplicationArtifactChartFileResponse() (response *GetApplicationArtifactChartFileResponse) {
	response = &GetApplicationArtifactChartFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用chart包文件
func (c *Client) GetApplicationArtifactChartFile(request *GetApplicationArtifactChartFileRequest) (response *GetApplicationArtifactChartFileResponse, err error) {
	if request == nil {
		request = NewGetApplicationArtifactChartFileRequest()
	}
	response = NewGetApplicationArtifactChartFileResponse()
	err = c.Send(request, response)
	return
}

func NewListSubOperationSheetsRequest() (request *ListSubOperationSheetsRequest) {
	request = &ListSubOperationSheetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListSubOperationSheets")
	return
}

func NewListSubOperationSheetsResponse() (response *ListSubOperationSheetsResponse) {
	response = &ListSubOperationSheetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询局点变更单的子变更单
func (c *Client) ListSubOperationSheets(request *ListSubOperationSheetsRequest) (response *ListSubOperationSheetsResponse, err error) {
	if request == nil {
		request = NewListSubOperationSheetsRequest()
	}
	response = NewListSubOperationSheetsResponse()
	err = c.Send(request, response)
	return
}

func NewGetCompareChartFileListRequest() (request *GetCompareChartFileListRequest) {
	request = &GetCompareChartFileListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetCompareChartFileList")
	return
}

func NewGetCompareChartFileListResponse() (response *GetCompareChartFileListResponse) {
	response = &GetCompareChartFileListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取对比chart包所有文件
func (c *Client) GetCompareChartFileList(request *GetCompareChartFileListRequest) (response *GetCompareChartFileListResponse, err error) {
	if request == nil {
		request = NewGetCompareChartFileListRequest()
	}
	response = NewGetCompareChartFileListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductPlanRequest() (request *DescribeProductPlanRequest) {
	request = &DescribeProductPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeProductPlan")
	return
}

func NewDescribeProductPlanResponse() (response *DescribeProductPlanResponse) {
	response = &DescribeProductPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据请求参数, 查询产品实例规划信息
func (c *Client) DescribeProductPlan(request *DescribeProductPlanRequest) (response *DescribeProductPlanResponse, err error) {
	if request == nil {
		request = NewDescribeProductPlanRequest()
	}
	response = NewDescribeProductPlanResponse()
	err = c.Send(request, response)
	return
}

func NewGetRunningApplicationNameRequest() (request *GetRunningApplicationNameRequest) {
	request = &GetRunningApplicationNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetRunningApplicationName")
	return
}

func NewGetRunningApplicationNameResponse() (response *GetRunningApplicationNameResponse) {
	response = &GetRunningApplicationNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 部署应用时，可以使用指定应用名称部署应用
func (c *Client) GetRunningApplicationName(request *GetRunningApplicationNameRequest) (response *GetRunningApplicationNameResponse, err error) {
	if request == nil {
		request = NewGetRunningApplicationNameRequest()
	}
	response = NewGetRunningApplicationNameResponse()
	err = c.Send(request, response)
	return
}

func NewGetLatestPackageUrlRequest() (request *GetLatestPackageUrlRequest) {
	request = &GetLatestPackageUrlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetLatestPackageUrl")
	return
}

func NewGetLatestPackageUrlResponse() (response *GetLatestPackageUrlResponse) {
	response = &GetLatestPackageUrlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取最新导入的packageUrl
func (c *Client) GetLatestPackageUrl(request *GetLatestPackageUrlRequest) (response *GetLatestPackageUrlResponse, err error) {
	if request == nil {
		request = NewGetLatestPackageUrlRequest()
	}
	response = NewGetLatestPackageUrlResponse()
	err = c.Send(request, response)
	return
}

func NewListSiteProductInstancesRequest() (request *ListSiteProductInstancesRequest) {
	request = &ListSiteProductInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListSiteProductInstances")
	return
}

func NewListSiteProductInstancesResponse() (response *ListSiteProductInstancesResponse) {
	response = &ListSiteProductInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询局点产品实例列表
func (c *Client) ListSiteProductInstances(request *ListSiteProductInstancesRequest) (response *ListSiteProductInstancesResponse, err error) {
	if request == nil {
		request = NewListSiteProductInstancesRequest()
	}
	response = NewListSiteProductInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewImportOperationTemplateRequest() (request *ImportOperationTemplateRequest) {
	request = &ImportOperationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ImportOperationTemplate")
	return
}

func NewImportOperationTemplateResponse() (response *ImportOperationTemplateResponse) {
	response = &ImportOperationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入运维模板物料包
func (c *Client) ImportOperationTemplate(request *ImportOperationTemplateRequest) (response *ImportOperationTemplateResponse, err error) {
	if request == nil {
		request = NewImportOperationTemplateRequest()
	}
	response = NewImportOperationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewListOperationSheetApprovalRequest() (request *ListOperationSheetApprovalRequest) {
	request = &ListOperationSheetApprovalRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListOperationSheetApproval")
	return
}

func NewListOperationSheetApprovalResponse() (response *ListOperationSheetApprovalResponse) {
	response = &ListOperationSheetApprovalResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询变更审批列表，包括待审批和已审批的变更单
func (c *Client) ListOperationSheetApproval(request *ListOperationSheetApprovalRequest) (response *ListOperationSheetApprovalResponse, err error) {
	if request == nil {
		request = NewListOperationSheetApprovalRequest()
	}
	response = NewListOperationSheetApprovalResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSolutionTemplateRequest() (request *DeleteSolutionTemplateRequest) {
	request = &DeleteSolutionTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DeleteSolutionTemplate")
	return
}

func NewDeleteSolutionTemplateResponse() (response *DeleteSolutionTemplateResponse) {
	response = &DeleteSolutionTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除解决方案运维模版
func (c *Client) DeleteSolutionTemplate(request *DeleteSolutionTemplateRequest) (response *DeleteSolutionTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteSolutionTemplateRequest()
	}
	response = NewDeleteSolutionTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewGetMachineSSHConfigRequest() (request *GetMachineSSHConfigRequest) {
	request = &GetMachineSSHConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetMachineSSHConfig")
	return
}

func NewGetMachineSSHConfigResponse() (response *GetMachineSSHConfigResponse) {
	response = &GetMachineSSHConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新ssh配置
func (c *Client) GetMachineSSHConfig(request *GetMachineSSHConfigRequest) (response *GetMachineSSHConfigResponse, err error) {
	if request == nil {
		request = NewGetMachineSSHConfigRequest()
	}
	response = NewGetMachineSSHConfigResponse()
	err = c.Send(request, response)
	return
}

func NewListCommonOperationSheetsRequest() (request *ListCommonOperationSheetsRequest) {
	request = &ListCommonOperationSheetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListCommonOperationSheets")
	return
}

func NewListCommonOperationSheetsResponse() (response *ListCommonOperationSheetsResponse) {
	response = &ListCommonOperationSheetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询通用变更单列表
func (c *Client) ListCommonOperationSheets(request *ListCommonOperationSheetsRequest) (response *ListCommonOperationSheetsResponse, err error) {
	if request == nil {
		request = NewListCommonOperationSheetsRequest()
	}
	response = NewListCommonOperationSheetsResponse()
	err = c.Send(request, response)
	return
}

func NewListProjectSiteInfoDetailsRequest() (request *ListProjectSiteInfoDetailsRequest) {
	request = &ListProjectSiteInfoDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListProjectSiteInfoDetails")
	return
}

func NewListProjectSiteInfoDetailsResponse() (response *ListProjectSiteInfoDetailsResponse) {
	response = &ListProjectSiteInfoDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看云列表
func (c *Client) ListProjectSiteInfoDetails(request *ListProjectSiteInfoDetailsRequest) (response *ListProjectSiteInfoDetailsResponse, err error) {
	if request == nil {
		request = NewListProjectSiteInfoDetailsRequest()
	}
	response = NewListProjectSiteInfoDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppComponentTreeRequest() (request *DescribeAppComponentTreeRequest) {
	request = &DescribeAppComponentTreeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeAppComponentTree")
	return
}

func NewDescribeAppComponentTreeResponse() (response *DescribeAppComponentTreeResponse) {
	response = &DescribeAppComponentTreeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用模块树
func (c *Client) DescribeAppComponentTree(request *DescribeAppComponentTreeRequest) (response *DescribeAppComponentTreeResponse, err error) {
	if request == nil {
		request = NewDescribeAppComponentTreeRequest()
	}
	response = NewDescribeAppComponentTreeResponse()
	err = c.Send(request, response)
	return
}

func NewExportAtomicRequest() (request *ExportAtomicRequest) {
	request = &ExportAtomicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ExportAtomic")
	return
}

func NewExportAtomicResponse() (response *ExportAtomicResponse) {
	response = &ExportAtomicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 请求原子操作名称和tag，导出原子操作包
func (c *Client) ExportAtomic(request *ExportAtomicRequest) (response *ExportAtomicResponse, err error) {
	if request == nil {
		request = NewExportAtomicRequest()
	}
	response = NewExportAtomicResponse()
	err = c.Send(request, response)
	return
}

func NewPauseJobRequest() (request *PauseJobRequest) {
	request = &PauseJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "PauseJob")
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

func NewGetMaterialCleanupSetRequest() (request *GetMaterialCleanupSetRequest) {
	request = &GetMaterialCleanupSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetMaterialCleanupSet")
	return
}

func NewGetMaterialCleanupSetResponse() (response *GetMaterialCleanupSetResponse) {
	response = &GetMaterialCleanupSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取物料清理设置
func (c *Client) GetMaterialCleanupSet(request *GetMaterialCleanupSetRequest) (response *GetMaterialCleanupSetResponse, err error) {
	if request == nil {
		request = NewGetMaterialCleanupSetRequest()
	}
	response = NewGetMaterialCleanupSetResponse()
	err = c.Send(request, response)
	return
}

func NewListSiteProductMaterialRequest() (request *ListSiteProductMaterialRequest) {
	request = &ListSiteProductMaterialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListSiteProductMaterial")
	return
}

func NewListSiteProductMaterialResponse() (response *ListSiteProductMaterialResponse) {
	response = &ListSiteProductMaterialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取局点产品物料列表
func (c *Client) ListSiteProductMaterial(request *ListSiteProductMaterialRequest) (response *ListSiteProductMaterialResponse, err error) {
	if request == nil {
		request = NewListSiteProductMaterialRequest()
	}
	response = NewListSiteProductMaterialResponse()
	err = c.Send(request, response)
	return
}

func NewStartProductInstanceOperationSheetRequest() (request *StartProductInstanceOperationSheetRequest) {
	request = &StartProductInstanceOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "StartProductInstanceOperationSheet")
	return
}

func NewStartProductInstanceOperationSheetResponse() (response *StartProductInstanceOperationSheetResponse) {
	response = &StartProductInstanceOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动产品实例运维单
func (c *Client) StartProductInstanceOperationSheet(request *StartProductInstanceOperationSheetRequest) (response *StartProductInstanceOperationSheetResponse, err error) {
	if request == nil {
		request = NewStartProductInstanceOperationSheetRequest()
	}
	response = NewStartProductInstanceOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewGetInitDeploySheetOperationTemplateRequest() (request *GetInitDeploySheetOperationTemplateRequest) {
	request = &GetInitDeploySheetOperationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetInitDeploySheetOperationTemplate")
	return
}

func NewGetInitDeploySheetOperationTemplateResponse() (response *GetInitDeploySheetOperationTemplateResponse) {
	response = &GetInitDeploySheetOperationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取初始化部署变更单编排模版
func (c *Client) GetInitDeploySheetOperationTemplate(request *GetInitDeploySheetOperationTemplateRequest) (response *GetInitDeploySheetOperationTemplateResponse, err error) {
	if request == nil {
		request = NewGetInitDeploySheetOperationTemplateRequest()
	}
	response = NewGetInitDeploySheetOperationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewGeneratePipelineRunRequest() (request *GeneratePipelineRunRequest) {
	request = &GeneratePipelineRunRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GeneratePipelineRun")
	return
}

func NewGeneratePipelineRunResponse() (response *GeneratePipelineRunResponse) {
	response = &GeneratePipelineRunResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 生成运维任务描述
func (c *Client) GeneratePipelineRun(request *GeneratePipelineRunRequest) (response *GeneratePipelineRunResponse, err error) {
	if request == nil {
		request = NewGeneratePipelineRunRequest()
	}
	response = NewGeneratePipelineRunResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductOverviewRequest() (request *DescribeProductOverviewRequest) {
	request = &DescribeProductOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeProductOverview")
	return
}

func NewDescribeProductOverviewResponse() (response *DescribeProductOverviewResponse) {
	response = &DescribeProductOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品概览，返回产品下的应用、工作负载、POD等信息
func (c *Client) DescribeProductOverview(request *DescribeProductOverviewRequest) (response *DescribeProductOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeProductOverviewRequest()
	}
	response = NewDescribeProductOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewConfirmCommonOperationSheetRequest() (request *ConfirmCommonOperationSheetRequest) {
	request = &ConfirmCommonOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ConfirmCommonOperationSheet")
	return
}

func NewConfirmCommonOperationSheetResponse() (response *ConfirmCommonOperationSheetResponse) {
	response = &ConfirmCommonOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 确认变更单变更完成
func (c *Client) ConfirmCommonOperationSheet(request *ConfirmCommonOperationSheetRequest) (response *ConfirmCommonOperationSheetResponse, err error) {
	if request == nil {
		request = NewConfirmCommonOperationSheetRequest()
	}
	response = NewConfirmCommonOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSolutionTemplateRequest() (request *CreateSolutionTemplateRequest) {
	request = &CreateSolutionTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CreateSolutionTemplate")
	return
}

func NewCreateSolutionTemplateResponse() (response *CreateSolutionTemplateResponse) {
	response = &CreateSolutionTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建解决方案运维模版
func (c *Client) CreateSolutionTemplate(request *CreateSolutionTemplateRequest) (response *CreateSolutionTemplateResponse, err error) {
	if request == nil {
		request = NewCreateSolutionTemplateRequest()
	}
	response = NewCreateSolutionTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAppPlanRequest() (request *ModifyAppPlanRequest) {
	request = &ModifyAppPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ModifyAppPlan")
	return
}

func NewModifyAppPlanResponse() (response *ModifyAppPlanResponse) {
	response = &ModifyAppPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据请求参数, 修改应用实例规划信息
func (c *Client) ModifyAppPlan(request *ModifyAppPlanRequest) (response *ModifyAppPlanResponse, err error) {
	if request == nil {
		request = NewModifyAppPlanRequest()
	}
	response = NewModifyAppPlanResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAtomicRenderDataRequest() (request *ModifyAtomicRenderDataRequest) {
	request = &ModifyAtomicRenderDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ModifyAtomicRenderData")
	return
}

func NewModifyAtomicRenderDataResponse() (response *ModifyAtomicRenderDataResponse) {
	response = &ModifyAtomicRenderDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改原子操作渲染数据
func (c *Client) ModifyAtomicRenderData(request *ModifyAtomicRenderDataRequest) (response *ModifyAtomicRenderDataResponse, err error) {
	if request == nil {
		request = NewModifyAtomicRenderDataRequest()
	}
	response = NewModifyAtomicRenderDataResponse()
	err = c.Send(request, response)
	return
}

func NewCreateImportMaterialTaskRequest() (request *CreateImportMaterialTaskRequest) {
	request = &CreateImportMaterialTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CreateImportMaterialTask")
	return
}

func NewCreateImportMaterialTaskResponse() (response *CreateImportMaterialTaskResponse) {
	response = &CreateImportMaterialTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建导入物料任务(包含下载物料)
func (c *Client) CreateImportMaterialTask(request *CreateImportMaterialTaskRequest) (response *CreateImportMaterialTaskResponse, err error) {
	if request == nil {
		request = NewCreateImportMaterialTaskRequest()
	}
	response = NewCreateImportMaterialTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAtomicRequest() (request *DescribeAtomicRequest) {
	request = &DescribeAtomicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeAtomic")
	return
}

func NewDescribeAtomicResponse() (response *DescribeAtomicResponse) {
	response = &DescribeAtomicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 请求原子操作名称&Version，返回原子操作的内容定义
func (c *Client) DescribeAtomic(request *DescribeAtomicRequest) (response *DescribeAtomicResponse, err error) {
	if request == nil {
		request = NewDescribeAtomicRequest()
	}
	response = NewDescribeAtomicResponse()
	err = c.Send(request, response)
	return
}

func NewListDeployedApplicationDetailsRequest() (request *ListDeployedApplicationDetailsRequest) {
	request = &ListDeployedApplicationDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListDeployedApplicationDetails")
	return
}

func NewListDeployedApplicationDetailsResponse() (response *ListDeployedApplicationDetailsResponse) {
	response = &ListDeployedApplicationDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品基本运维信息
func (c *Client) ListDeployedApplicationDetails(request *ListDeployedApplicationDetailsRequest) (response *ListDeployedApplicationDetailsResponse, err error) {
	if request == nil {
		request = NewListDeployedApplicationDetailsRequest()
	}
	response = NewListDeployedApplicationDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewListPlanMiddlewareInfosRequest() (request *ListPlanMiddlewareInfosRequest) {
	request = &ListPlanMiddlewareInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListPlanMiddlewareInfos")
	return
}

func NewListPlanMiddlewareInfosResponse() (response *ListPlanMiddlewareInfosResponse) {
	response = &ListPlanMiddlewareInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用模块树
func (c *Client) ListPlanMiddlewareInfos(request *ListPlanMiddlewareInfosRequest) (response *ListPlanMiddlewareInfosResponse, err error) {
	if request == nil {
		request = NewListPlanMiddlewareInfosRequest()
	}
	response = NewListPlanMiddlewareInfosResponse()
	err = c.Send(request, response)
	return
}

func NewListOperationTemplateRequest() (request *ListOperationTemplateRequest) {
	request = &ListOperationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListOperationTemplate")
	return
}

func NewListOperationTemplateResponse() (response *ListOperationTemplateResponse) {
	response = &ListOperationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取运维模版列表
func (c *Client) ListOperationTemplate(request *ListOperationTemplateRequest) (response *ListOperationTemplateResponse, err error) {
	if request == nil {
		request = NewListOperationTemplateRequest()
	}
	response = NewListOperationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewGetLocalTCSCoreIPListRequest() (request *GetLocalTCSCoreIPListRequest) {
	request = &GetLocalTCSCoreIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetLocalTCSCoreIPList")
	return
}

func NewGetLocalTCSCoreIPListResponse() (response *GetLocalTCSCoreIPListResponse) {
	response = &GetLocalTCSCoreIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取本地k8s底座的ip信息
func (c *Client) GetLocalTCSCoreIPList(request *GetLocalTCSCoreIPListRequest) (response *GetLocalTCSCoreIPListResponse, err error) {
	if request == nil {
		request = NewGetLocalTCSCoreIPListRequest()
	}
	response = NewGetLocalTCSCoreIPListResponse()
	err = c.Send(request, response)
	return
}

func NewExportServerResourcesRequest() (request *ExportServerResourcesRequest) {
	request = &ExportServerResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ExportServerResources")
	return
}

func NewExportServerResourcesResponse() (response *ExportServerResourcesResponse) {
	response = &ExportServerResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出规划包
func (c *Client) ExportServerResources(request *ExportServerResourcesRequest) (response *ExportServerResourcesResponse, err error) {
	if request == nil {
		request = NewExportServerResourcesRequest()
	}
	response = NewExportServerResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewAtomicJsonToYamlRequest() (request *AtomicJsonToYamlRequest) {
	request = &AtomicJsonToYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "AtomicJsonToYaml")
	return
}

func NewAtomicJsonToYamlResponse() (response *AtomicJsonToYamlResponse) {
	response = &AtomicJsonToYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 请求原子操作json，输出原子操作yaml结构
func (c *Client) AtomicJsonToYaml(request *AtomicJsonToYamlRequest) (response *AtomicJsonToYamlResponse, err error) {
	if request == nil {
		request = NewAtomicJsonToYamlRequest()
	}
	response = NewAtomicJsonToYamlResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAtomicRequest() (request *DeleteAtomicRequest) {
	request = &DeleteAtomicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DeleteAtomic")
	return
}

func NewDeleteAtomicResponse() (response *DeleteAtomicResponse) {
	response = &DeleteAtomicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 请求原子操作名称&Version，清理原子操作，如果当前原子操作被其他任意原子操作引用，就无法删除
func (c *Client) DeleteAtomic(request *DeleteAtomicRequest) (response *DeleteAtomicResponse, err error) {
	if request == nil {
		request = NewDeleteAtomicRequest()
	}
	response = NewDeleteAtomicResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductInstanceOperationSheetRequest() (request *DescribeProductInstanceOperationSheetRequest) {
	request = &DescribeProductInstanceOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeProductInstanceOperationSheet")
	return
}

func NewDescribeProductInstanceOperationSheetResponse() (response *DescribeProductInstanceOperationSheetResponse) {
	response = &DescribeProductInstanceOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品实例运维单
func (c *Client) DescribeProductInstanceOperationSheet(request *DescribeProductInstanceOperationSheetRequest) (response *DescribeProductInstanceOperationSheetResponse, err error) {
	if request == nil {
		request = NewDescribeProductInstanceOperationSheetRequest()
	}
	response = NewDescribeProductInstanceOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewCheckProductInstanceOperationSheetTemplateAndAtomicRequest() (request *CheckProductInstanceOperationSheetTemplateAndAtomicRequest) {
	request = &CheckProductInstanceOperationSheetTemplateAndAtomicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CheckProductInstanceOperationSheetTemplateAndAtomic")
	return
}

func NewCheckProductInstanceOperationSheetTemplateAndAtomicResponse() (response *CheckProductInstanceOperationSheetTemplateAndAtomicResponse) {
	response = &CheckProductInstanceOperationSheetTemplateAndAtomicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验变更项物料是否存在，包括原子操作和运维模板
func (c *Client) CheckProductInstanceOperationSheetTemplateAndAtomic(request *CheckProductInstanceOperationSheetTemplateAndAtomicRequest) (response *CheckProductInstanceOperationSheetTemplateAndAtomicResponse, err error) {
	if request == nil {
		request = NewCheckProductInstanceOperationSheetTemplateAndAtomicRequest()
	}
	response = NewCheckProductInstanceOperationSheetTemplateAndAtomicResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductBaseInfoRequest() (request *DescribeProductBaseInfoRequest) {
	request = &DescribeProductBaseInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeProductBaseInfo")
	return
}

func NewDescribeProductBaseInfoResponse() (response *DescribeProductBaseInfoResponse) {
	response = &DescribeProductBaseInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品基本运维信息
func (c *Client) DescribeProductBaseInfo(request *DescribeProductBaseInfoRequest) (response *DescribeProductBaseInfoResponse, err error) {
	if request == nil {
		request = NewDescribeProductBaseInfoRequest()
	}
	response = NewDescribeProductBaseInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPlanServerInfoRequest() (request *ModifyPlanServerInfoRequest) {
	request = &ModifyPlanServerInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ModifyPlanServerInfo")
	return
}

func NewModifyPlanServerInfoResponse() (response *ModifyPlanServerInfoResponse) {
	response = &ModifyPlanServerInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改规划的服务器信息
func (c *Client) ModifyPlanServerInfo(request *ModifyPlanServerInfoRequest) (response *ModifyPlanServerInfoResponse, err error) {
	if request == nil {
		request = NewModifyPlanServerInfoRequest()
	}
	response = NewModifyPlanServerInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductTemplateRequest() (request *DescribeProductTemplateRequest) {
	request = &DescribeProductTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeProductTemplate")
	return
}

func NewDescribeProductTemplateResponse() (response *DescribeProductTemplateResponse) {
	response = &DescribeProductTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品模板
func (c *Client) DescribeProductTemplate(request *DescribeProductTemplateRequest) (response *DescribeProductTemplateResponse, err error) {
	if request == nil {
		request = NewDescribeProductTemplateRequest()
	}
	response = NewDescribeProductTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProductPlanRequest() (request *ModifyProductPlanRequest) {
	request = &ModifyProductPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ModifyProductPlan")
	return
}

func NewModifyProductPlanResponse() (response *ModifyProductPlanResponse) {
	response = &ModifyProductPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据请求参数, 修改产品实例规划信息
func (c *Client) ModifyProductPlan(request *ModifyProductPlanRequest) (response *ModifyProductPlanResponse, err error) {
	if request == nil {
		request = NewModifyProductPlanRequest()
	}
	response = NewModifyProductPlanResponse()
	err = c.Send(request, response)
	return
}

func NewGetSiteMaterialInfoRequest() (request *GetSiteMaterialInfoRequest) {
	request = &GetSiteMaterialInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetSiteMaterialInfo")
	return
}

func NewGetSiteMaterialInfoResponse() (response *GetSiteMaterialInfoResponse) {
	response = &GetSiteMaterialInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取局点物料详情
func (c *Client) GetSiteMaterialInfo(request *GetSiteMaterialInfoRequest) (response *GetSiteMaterialInfoResponse, err error) {
	if request == nil {
		request = NewGetSiteMaterialInfoRequest()
	}
	response = NewGetSiteMaterialInfoResponse()
	err = c.Send(request, response)
	return
}

func NewImportDeploymentPhaseApplicationArtifactRequest() (request *ImportDeploymentPhaseApplicationArtifactRequest) {
	request = &ImportDeploymentPhaseApplicationArtifactRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ImportDeploymentPhaseApplicationArtifact")
	return
}

func NewImportDeploymentPhaseApplicationArtifactResponse() (response *ImportDeploymentPhaseApplicationArtifactResponse) {
	response = &ImportDeploymentPhaseApplicationArtifactResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 研发阶段应用物料导入
func (c *Client) ImportDeploymentPhaseApplicationArtifact(request *ImportDeploymentPhaseApplicationArtifactRequest) (response *ImportDeploymentPhaseApplicationArtifactResponse, err error) {
	if request == nil {
		request = NewImportDeploymentPhaseApplicationArtifactRequest()
	}
	response = NewImportDeploymentPhaseApplicationArtifactResponse()
	err = c.Send(request, response)
	return
}

func NewRetryNodeRequest() (request *RetryNodeRequest) {
	request = &RetryNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "RetryNode")
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

func NewRetryImportMaterialRequest() (request *RetryImportMaterialRequest) {
	request = &RetryImportMaterialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "RetryImportMaterial")
	return
}

func NewRetryImportMaterialResponse() (response *RetryImportMaterialResponse) {
	response = &RetryImportMaterialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物料导入重试
func (c *Client) RetryImportMaterial(request *RetryImportMaterialRequest) (response *RetryImportMaterialResponse, err error) {
	if request == nil {
		request = NewRetryImportMaterialRequest()
	}
	response = NewRetryImportMaterialResponse()
	err = c.Send(request, response)
	return
}

func NewGeneratePipelineRequest() (request *GeneratePipelineRequest) {
	request = &GeneratePipelineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GeneratePipeline")
	return
}

func NewGeneratePipelineResponse() (response *GeneratePipelineResponse) {
	response = &GeneratePipelineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 生成编排Pipeline
func (c *Client) GeneratePipeline(request *GeneratePipelineRequest) (response *GeneratePipelineResponse, err error) {
	if request == nil {
		request = NewGeneratePipelineRequest()
	}
	response = NewGeneratePipelineResponse()
	err = c.Send(request, response)
	return
}

func NewListProductInstancePlansRequest() (request *ListProductInstancePlansRequest) {
	request = &ListProductInstancePlansRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListProductInstancePlans")
	return
}

func NewListProductInstancePlansResponse() (response *ListProductInstancePlansResponse) {
	response = &ListProductInstancePlansResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回产品实例规划信息列表
func (c *Client) ListProductInstancePlans(request *ListProductInstancePlansRequest) (response *ListProductInstancePlansResponse, err error) {
	if request == nil {
		request = NewListProductInstancePlansRequest()
	}
	response = NewListProductInstancePlansResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppRollbackConfigRequest() (request *DescribeAppRollbackConfigRequest) {
	request = &DescribeAppRollbackConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeAppRollbackConfig")
	return
}

func NewDescribeAppRollbackConfigResponse() (response *DescribeAppRollbackConfigResponse) {
	response = &DescribeAppRollbackConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用回滚配置
func (c *Client) DescribeAppRollbackConfig(request *DescribeAppRollbackConfigRequest) (response *DescribeAppRollbackConfigResponse, err error) {
	if request == nil {
		request = NewDescribeAppRollbackConfigRequest()
	}
	response = NewDescribeAppRollbackConfigResponse()
	err = c.Send(request, response)
	return
}

func NewListAtomicLabelRequest() (request *ListAtomicLabelRequest) {
	request = &ListAtomicLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListAtomicLabel")
	return
}

func NewListAtomicLabelResponse() (response *ListAtomicLabelResponse) {
	response = &ListAtomicLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 输出原子操作标签列表
func (c *Client) ListAtomicLabel(request *ListAtomicLabelRequest) (response *ListAtomicLabelResponse, err error) {
	if request == nil {
		request = NewListAtomicLabelRequest()
	}
	response = NewListAtomicLabelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJobOverviewRequest() (request *DescribeJobOverviewRequest) {
	request = &DescribeJobOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeJobOverview")
	return
}

func NewDescribeJobOverviewResponse() (response *DescribeJobOverviewResponse) {
	response = &DescribeJobOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询编排任务状态
func (c *Client) DescribeJobOverview(request *DescribeJobOverviewRequest) (response *DescribeJobOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeJobOverviewRequest()
	}
	response = NewDescribeJobOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOperationTemplateRequest() (request *ModifyOperationTemplateRequest) {
	request = &ModifyOperationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ModifyOperationTemplate")
	return
}

func NewModifyOperationTemplateResponse() (response *ModifyOperationTemplateResponse) {
	response = &ModifyOperationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新运维模版
func (c *Client) ModifyOperationTemplate(request *ModifyOperationTemplateRequest) (response *ModifyOperationTemplateResponse, err error) {
	if request == nil {
		request = NewModifyOperationTemplateRequest()
	}
	response = NewModifyOperationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewGetRuntimeSecretRequest() (request *GetRuntimeSecretRequest) {
	request = &GetRuntimeSecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetRuntimeSecret")
	return
}

func NewGetRuntimeSecretResponse() (response *GetRuntimeSecretResponse) {
	response = &GetRuntimeSecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 使用Binding获取运行时Secret信息
func (c *Client) GetRuntimeSecret(request *GetRuntimeSecretRequest) (response *GetRuntimeSecretResponse, err error) {
	if request == nil {
		request = NewGetRuntimeSecretRequest()
	}
	response = NewGetRuntimeSecretResponse()
	err = c.Send(request, response)
	return
}

func NewListPlanHistoriesRequest() (request *ListPlanHistoriesRequest) {
	request = &ListPlanHistoriesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListPlanHistories")
	return
}

func NewListPlanHistoriesResponse() (response *ListPlanHistoriesResponse) {
	response = &ListPlanHistoriesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询规划修改历史记录
func (c *Client) ListPlanHistories(request *ListPlanHistoriesRequest) (response *ListPlanHistoriesResponse, err error) {
	if request == nil {
		request = NewListPlanHistoriesRequest()
	}
	response = NewListPlanHistoriesResponse()
	err = c.Send(request, response)
	return
}

func NewResumeNodeRequest() (request *ResumeNodeRequest) {
	request = &ResumeNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ResumeNode")
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

func NewDescribeCloudTopologyRequest() (request *DescribeCloudTopologyRequest) {
	request = &DescribeCloudTopologyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeCloudTopology")
	return
}

func NewDescribeCloudTopologyResponse() (response *DescribeCloudTopologyResponse) {
	response = &DescribeCloudTopologyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取局点（Cloud）下的Region、Zone信息
func (c *Client) DescribeCloudTopology(request *DescribeCloudTopologyRequest) (response *DescribeCloudTopologyResponse, err error) {
	if request == nil {
		request = NewDescribeCloudTopologyRequest()
	}
	response = NewDescribeCloudTopologyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductRollbackConfigRequest() (request *DescribeProductRollbackConfigRequest) {
	request = &DescribeProductRollbackConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeProductRollbackConfig")
	return
}

func NewDescribeProductRollbackConfigResponse() (response *DescribeProductRollbackConfigResponse) {
	response = &DescribeProductRollbackConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品实例回滚配置
func (c *Client) DescribeProductRollbackConfig(request *DescribeProductRollbackConfigRequest) (response *DescribeProductRollbackConfigResponse, err error) {
	if request == nil {
		request = NewDescribeProductRollbackConfigRequest()
	}
	response = NewDescribeProductRollbackConfigResponse()
	err = c.Send(request, response)
	return
}

func NewListApplicationPackagesRequest() (request *ListApplicationPackagesRequest) {
	request = &ListApplicationPackagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListApplicationPackages")
	return
}

func NewListApplicationPackagesResponse() (response *ListApplicationPackagesResponse) {
	response = &ListApplicationPackagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用版本
func (c *Client) ListApplicationPackages(request *ListApplicationPackagesRequest) (response *ListApplicationPackagesResponse, err error) {
	if request == nil {
		request = NewListApplicationPackagesRequest()
	}
	response = NewListApplicationPackagesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppRuntimeYAMLRequest() (request *DescribeAppRuntimeYAMLRequest) {
	request = &DescribeAppRuntimeYAMLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeAppRuntimeYAML")
	return
}

func NewDescribeAppRuntimeYAMLResponse() (response *DescribeAppRuntimeYAMLResponse) {
	response = &DescribeAppRuntimeYAMLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用运行时YAML信息
func (c *Client) DescribeAppRuntimeYAML(request *DescribeAppRuntimeYAMLRequest) (response *DescribeAppRuntimeYAMLResponse, err error) {
	if request == nil {
		request = NewDescribeAppRuntimeYAMLRequest()
	}
	response = NewDescribeAppRuntimeYAMLResponse()
	err = c.Send(request, response)
	return
}

func NewImportSolutionTemplateRequest() (request *ImportSolutionTemplateRequest) {
	request = &ImportSolutionTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ImportSolutionTemplate")
	return
}

func NewImportSolutionTemplateResponse() (response *ImportSolutionTemplateResponse) {
	response = &ImportSolutionTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入解决方案运维模板物料包
func (c *Client) ImportSolutionTemplate(request *ImportSolutionTemplateRequest) (response *ImportSolutionTemplateResponse, err error) {
	if request == nil {
		request = NewImportSolutionTemplateRequest()
	}
	response = NewImportSolutionTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyChartFileRequest() (request *ModifyChartFileRequest) {
	request = &ModifyChartFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ModifyChartFile")
	return
}

func NewModifyChartFileResponse() (response *ModifyChartFileResponse) {
	response = &ModifyChartFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改chart文件内容
func (c *Client) ModifyChartFile(request *ModifyChartFileRequest) (response *ModifyChartFileResponse, err error) {
	if request == nil {
		request = NewModifyChartFileRequest()
	}
	response = NewModifyChartFileResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialProductApplicationInstanceRequest() (request *ListMaterialProductApplicationInstanceRequest) {
	request = &ListMaterialProductApplicationInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListMaterialProductApplicationInstance")
	return
}

func NewListMaterialProductApplicationInstanceResponse() (response *ListMaterialProductApplicationInstanceResponse) {
	response = &ListMaterialProductApplicationInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取物料交付方案应用实例列表
func (c *Client) ListMaterialProductApplicationInstance(request *ListMaterialProductApplicationInstanceRequest) (response *ListMaterialProductApplicationInstanceResponse, err error) {
	if request == nil {
		request = NewListMaterialProductApplicationInstanceRequest()
	}
	response = NewListMaterialProductApplicationInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCommonOperationSheetRequest() (request *UpdateCommonOperationSheetRequest) {
	request = &UpdateCommonOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "UpdateCommonOperationSheet")
	return
}

func NewUpdateCommonOperationSheetResponse() (response *UpdateCommonOperationSheetResponse) {
	response = &UpdateCommonOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新通用变更单
func (c *Client) UpdateCommonOperationSheet(request *UpdateCommonOperationSheetRequest) (response *UpdateCommonOperationSheetResponse, err error) {
	if request == nil {
		request = NewUpdateCommonOperationSheetRequest()
	}
	response = NewUpdateCommonOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationCcDeclareRequest() (request *GetApplicationCcDeclareRequest) {
	request = &GetApplicationCcDeclareRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetApplicationCcDeclare")
	return
}

func NewGetApplicationCcDeclareResponse() (response *GetApplicationCcDeclareResponse) {
	response = &GetApplicationCcDeclareResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用包CcDeclare
func (c *Client) GetApplicationCcDeclare(request *GetApplicationCcDeclareRequest) (response *GetApplicationCcDeclareResponse, err error) {
	if request == nil {
		request = NewGetApplicationCcDeclareRequest()
	}
	response = NewGetApplicationCcDeclareResponse()
	err = c.Send(request, response)
	return
}

func NewListSheetAppInstancesRequest() (request *ListSheetAppInstancesRequest) {
	request = &ListSheetAppInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListSheetAppInstances")
	return
}

func NewListSheetAppInstancesResponse() (response *ListSheetAppInstancesResponse) {
	response = &ListSheetAppInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运维单关联应用实例列表
func (c *Client) ListSheetAppInstances(request *ListSheetAppInstancesRequest) (response *ListSheetAppInstancesResponse, err error) {
	if request == nil {
		request = NewListSheetAppInstancesRequest()
	}
	response = NewListSheetAppInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProductInstanceOperationSheetRequest() (request *ModifyProductInstanceOperationSheetRequest) {
	request = &ModifyProductInstanceOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ModifyProductInstanceOperationSheet")
	return
}

func NewModifyProductInstanceOperationSheetResponse() (response *ModifyProductInstanceOperationSheetResponse) {
	response = &ModifyProductInstanceOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改产品实例运维单
func (c *Client) ModifyProductInstanceOperationSheet(request *ModifyProductInstanceOperationSheetRequest) (response *ModifyProductInstanceOperationSheetResponse, err error) {
	if request == nil {
		request = NewModifyProductInstanceOperationSheetRequest()
	}
	response = NewModifyProductInstanceOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewCheckProductInstanceOperationSheetMaterialRequest() (request *CheckProductInstanceOperationSheetMaterialRequest) {
	request = &CheckProductInstanceOperationSheetMaterialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CheckProductInstanceOperationSheetMaterial")
	return
}

func NewCheckProductInstanceOperationSheetMaterialResponse() (response *CheckProductInstanceOperationSheetMaterialResponse) {
	response = &CheckProductInstanceOperationSheetMaterialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验应用制品是否存在
func (c *Client) CheckProductInstanceOperationSheetMaterial(request *CheckProductInstanceOperationSheetMaterialRequest) (response *CheckProductInstanceOperationSheetMaterialResponse, err error) {
	if request == nil {
		request = NewCheckProductInstanceOperationSheetMaterialRequest()
	}
	response = NewCheckProductInstanceOperationSheetMaterialResponse()
	err = c.Send(request, response)
	return
}

func NewConvertPipelineFormatRequest() (request *ConvertPipelineFormatRequest) {
	request = &ConvertPipelineFormatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ConvertPipelineFormat")
	return
}

func NewConvertPipelineFormatResponse() (response *ConvertPipelineFormatResponse) {
	response = &ConvertPipelineFormatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编排文件格式转换
func (c *Client) ConvertPipelineFormat(request *ConvertPipelineFormatRequest) (response *ConvertPipelineFormatResponse, err error) {
	if request == nil {
		request = NewConvertPipelineFormatRequest()
	}
	response = NewConvertPipelineFormatResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperationSheetRequest() (request *DeleteOperationSheetRequest) {
	request = &DeleteOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DeleteOperationSheet")
	return
}

func NewDeleteOperationSheetResponse() (response *DeleteOperationSheetResponse) {
	response = &DeleteOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除运维历史
func (c *Client) DeleteOperationSheet(request *DeleteOperationSheetRequest) (response *DeleteOperationSheetResponse, err error) {
	if request == nil {
		request = NewDeleteOperationSheetRequest()
	}
	response = NewDeleteOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewListPlansRequest() (request *ListPlansRequest) {
	request = &ListPlansRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListPlans")
	return
}

func NewListPlansResponse() (response *ListPlansResponse) {
	response = &ListPlansResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询规划版本列表
func (c *Client) ListPlans(request *ListPlansRequest) (response *ListPlansResponse, err error) {
	if request == nil {
		request = NewListPlansRequest()
	}
	response = NewListPlansResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaterialRequest() (request *DeleteMaterialRequest) {
	request = &DeleteMaterialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DeleteMaterial")
	return
}

func NewDeleteMaterialResponse() (response *DeleteMaterialResponse) {
	response = &DeleteMaterialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除物料的导入目录并标识删除人和物料原始物料已删除
func (c *Client) DeleteMaterial(request *DeleteMaterialRequest) (response *DeleteMaterialResponse, err error) {
	if request == nil {
		request = NewDeleteMaterialRequest()
	}
	response = NewDeleteMaterialResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppPlanRequest() (request *DescribeAppPlanRequest) {
	request = &DescribeAppPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeAppPlan")
	return
}

func NewDescribeAppPlanResponse() (response *DescribeAppPlanResponse) {
	response = &DescribeAppPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据请求参数, 查询应用实例规划信息
func (c *Client) DescribeAppPlan(request *DescribeAppPlanRequest) (response *DescribeAppPlanResponse, err error) {
	if request == nil {
		request = NewDescribeAppPlanRequest()
	}
	response = NewDescribeAppPlanResponse()
	err = c.Send(request, response)
	return
}

func NewDiffPlanPackageRequest() (request *DiffPlanPackageRequest) {
	request = &DiffPlanPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DiffPlanPackage")
	return
}

func NewDiffPlanPackageResponse() (response *DiffPlanPackageResponse) {
	response = &DiffPlanPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回两个规划包有差异的资源列表
func (c *Client) DiffPlanPackage(request *DiffPlanPackageRequest) (response *DiffPlanPackageResponse, err error) {
	if request == nil {
		request = NewDiffPlanPackageRequest()
	}
	response = NewDiffPlanPackageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOtaCertificateRequest() (request *DescribeOtaCertificateRequest) {
	request = &DescribeOtaCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeOtaCertificate")
	return
}

func NewDescribeOtaCertificateResponse() (response *DescribeOtaCertificateResponse) {
	response = &DescribeOtaCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取局点OTA证书信息
func (c *Client) DescribeOtaCertificate(request *DescribeOtaCertificateRequest) (response *DescribeOtaCertificateResponse, err error) {
	if request == nil {
		request = NewDescribeOtaCertificateRequest()
	}
	response = NewDescribeOtaCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialOperationTemplateRequest() (request *ListMaterialOperationTemplateRequest) {
	request = &ListMaterialOperationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListMaterialOperationTemplate")
	return
}

func NewListMaterialOperationTemplateResponse() (response *ListMaterialOperationTemplateResponse) {
	response = &ListMaterialOperationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取物料运维模板
func (c *Client) ListMaterialOperationTemplate(request *ListMaterialOperationTemplateRequest) (response *ListMaterialOperationTemplateResponse, err error) {
	if request == nil {
		request = NewListMaterialOperationTemplateRequest()
	}
	response = NewListMaterialOperationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewImportAtomicRequest() (request *ImportAtomicRequest) {
	request = &ImportAtomicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ImportAtomic")
	return
}

func NewImportAtomicResponse() (response *ImportAtomicResponse) {
	response = &ImportAtomicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据原子操作物料包，导入原子操作
func (c *Client) ImportAtomic(request *ImportAtomicRequest) (response *ImportAtomicResponse, err error) {
	if request == nil {
		request = NewImportAtomicRequest()
	}
	response = NewImportAtomicResponse()
	err = c.Send(request, response)
	return
}

func NewMaterialCleanupSetRequest() (request *MaterialCleanupSetRequest) {
	request = &MaterialCleanupSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "MaterialCleanupSet")
	return
}

func NewMaterialCleanupSetResponse() (response *MaterialCleanupSetResponse) {
	response = &MaterialCleanupSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物料清理设置
func (c *Client) MaterialCleanupSet(request *MaterialCleanupSetRequest) (response *MaterialCleanupSetResponse, err error) {
	if request == nil {
		request = NewMaterialCleanupSetRequest()
	}
	response = NewMaterialCleanupSetResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialProductDetailsRequest() (request *ListMaterialProductDetailsRequest) {
	request = &ListMaterialProductDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListMaterialProductDetails")
	return
}

func NewListMaterialProductDetailsResponse() (response *ListMaterialProductDetailsResponse) {
	response = &ListMaterialProductDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解析物料中的产品信息
func (c *Client) ListMaterialProductDetails(request *ListMaterialProductDetailsRequest) (response *ListMaterialProductDetailsResponse, err error) {
	if request == nil {
		request = NewListMaterialProductDetailsRequest()
	}
	response = NewListMaterialProductDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewParseProductTemplateRequest() (request *ParseProductTemplateRequest) {
	request = &ParseProductTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ParseProductTemplate")
	return
}

func NewParseProductTemplateResponse() (response *ParseProductTemplateResponse) {
	response = &ParseProductTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解析产品模板
func (c *Client) ParseProductTemplate(request *ParseProductTemplateRequest) (response *ParseProductTemplateResponse, err error) {
	if request == nil {
		request = NewParseProductTemplateRequest()
	}
	response = NewParseProductTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewStartSiteOperationSheetRequest() (request *StartSiteOperationSheetRequest) {
	request = &StartSiteOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "StartSiteOperationSheet")
	return
}

func NewStartSiteOperationSheetResponse() (response *StartSiteOperationSheetResponse) {
	response = &StartSiteOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动局点变更单
func (c *Client) StartSiteOperationSheet(request *StartSiteOperationSheetRequest) (response *StartSiteOperationSheetResponse, err error) {
	if request == nil {
		request = NewStartSiteOperationSheetRequest()
	}
	response = NewStartSiteOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialAtomicRequest() (request *ListMaterialAtomicRequest) {
	request = &ListMaterialAtomicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListMaterialAtomic")
	return
}

func NewListMaterialAtomicResponse() (response *ListMaterialAtomicResponse) {
	response = &ListMaterialAtomicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取物料原子操作
func (c *Client) ListMaterialAtomic(request *ListMaterialAtomicRequest) (response *ListMaterialAtomicResponse, err error) {
	if request == nil {
		request = NewListMaterialAtomicRequest()
	}
	response = NewListMaterialAtomicResponse()
	err = c.Send(request, response)
	return
}

func NewAtomicYamlToJsonRequest() (request *AtomicYamlToJsonRequest) {
	request = &AtomicYamlToJsonRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "AtomicYamlToJson")
	return
}

func NewAtomicYamlToJsonResponse() (response *AtomicYamlToJsonResponse) {
	response = &AtomicYamlToJsonResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 请求原子操作yaml结构，输出原子操作json结构
func (c *Client) AtomicYamlToJson(request *AtomicYamlToJsonRequest) (response *AtomicYamlToJsonResponse, err error) {
	if request == nil {
		request = NewAtomicYamlToJsonRequest()
	}
	response = NewAtomicYamlToJsonResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCommonOperationSheetRequest() (request *DeleteCommonOperationSheetRequest) {
	request = &DeleteCommonOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DeleteCommonOperationSheet")
	return
}

func NewDeleteCommonOperationSheetResponse() (response *DeleteCommonOperationSheetResponse) {
	response = &DeleteCommonOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除通用变更单
func (c *Client) DeleteCommonOperationSheet(request *DeleteCommonOperationSheetRequest) (response *DeleteCommonOperationSheetResponse, err error) {
	if request == nil {
		request = NewDeleteCommonOperationSheetRequest()
	}
	response = NewDeleteCommonOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewListProductInstanceSubsystemsRequest() (request *ListProductInstanceSubsystemsRequest) {
	request = &ListProductInstanceSubsystemsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListProductInstanceSubsystems")
	return
}

func NewListProductInstanceSubsystemsResponse() (response *ListProductInstanceSubsystemsResponse) {
	response = &ListProductInstanceSubsystemsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取运维中心的产品实例子系统列表
func (c *Client) ListProductInstanceSubsystems(request *ListProductInstanceSubsystemsRequest) (response *ListProductInstanceSubsystemsResponse, err error) {
	if request == nil {
		request = NewListProductInstanceSubsystemsRequest()
	}
	response = NewListProductInstanceSubsystemsResponse()
	err = c.Send(request, response)
	return
}

func NewSetActivePlanRequest() (request *SetActivePlanRequest) {
	request = &SetActivePlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "SetActivePlan")
	return
}

func NewSetActivePlanResponse() (response *SetActivePlanResponse) {
	response = &SetActivePlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 为局点设置生效的规划版本
func (c *Client) SetActivePlan(request *SetActivePlanRequest) (response *SetActivePlanResponse, err error) {
	if request == nil {
		request = NewSetActivePlanRequest()
	}
	response = NewSetActivePlanResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductTemplateIDRequest() (request *GetProductTemplateIDRequest) {
	request = &GetProductTemplateIDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetProductTemplateID")
	return
}

func NewGetProductTemplateIDResponse() (response *GetProductTemplateIDResponse) {
	response = &GetProductTemplateIDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品部署模版
func (c *Client) GetProductTemplateID(request *GetProductTemplateIDRequest) (response *GetProductTemplateIDResponse, err error) {
	if request == nil {
		request = NewGetProductTemplateIDRequest()
	}
	response = NewGetProductTemplateIDResponse()
	err = c.Send(request, response)
	return
}

func NewParseProductDataRequest() (request *ParseProductDataRequest) {
	request = &ParseProductDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ParseProductData")
	return
}

func NewParseProductDataResponse() (response *ParseProductDataResponse) {
	response = &ParseProductDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品数据解析
func (c *Client) ParseProductData(request *ParseProductDataRequest) (response *ParseProductDataResponse, err error) {
	if request == nil {
		request = NewParseProductDataRequest()
	}
	response = NewParseProductDataResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAtomicRequest() (request *CreateAtomicRequest) {
	request = &CreateAtomicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CreateAtomic")
	return
}

func NewCreateAtomicResponse() (response *CreateAtomicResponse) {
	response = &CreateAtomicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 请求原子操作定义，创建原子操作申明
func (c *Client) CreateAtomic(request *CreateAtomicRequest) (response *CreateAtomicResponse, err error) {
	if request == nil {
		request = NewCreateAtomicRequest()
	}
	response = NewCreateAtomicResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSiteOperationSheetRequest() (request *DescribeSiteOperationSheetRequest) {
	request = &DescribeSiteOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeSiteOperationSheet")
	return
}

func NewDescribeSiteOperationSheetResponse() (response *DescribeSiteOperationSheetResponse) {
	response = &DescribeSiteOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询局点变更单详情
func (c *Client) DescribeSiteOperationSheet(request *DescribeSiteOperationSheetRequest) (response *DescribeSiteOperationSheetResponse, err error) {
	if request == nil {
		request = NewDescribeSiteOperationSheetRequest()
	}
	response = NewDescribeSiteOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationSheetRequest() (request *DescribeOperationSheetRequest) {
	request = &DescribeOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeOperationSheet")
	return
}

func NewDescribeOperationSheetResponse() (response *DescribeOperationSheetResponse) {
	response = &DescribeOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运维单
func (c *Client) DescribeOperationSheet(request *DescribeOperationSheetRequest) (response *DescribeOperationSheetResponse, err error) {
	if request == nil {
		request = NewDescribeOperationSheetRequest()
	}
	response = NewDescribeOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewGetRuntimePlanRequest() (request *GetRuntimePlanRequest) {
	request = &GetRuntimePlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetRuntimePlan")
	return
}

func NewGetRuntimePlanResponse() (response *GetRuntimePlanResponse) {
	response = &GetRuntimePlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用运行时规划
func (c *Client) GetRuntimePlan(request *GetRuntimePlanRequest) (response *GetRuntimePlanResponse, err error) {
	if request == nil {
		request = NewGetRuntimePlanRequest()
	}
	response = NewGetRuntimePlanResponse()
	err = c.Send(request, response)
	return
}

func NewListPublicMaterialInfosRequest() (request *ListPublicMaterialInfosRequest) {
	request = &ListPublicMaterialInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListPublicMaterialInfos")
	return
}

func NewListPublicMaterialInfosResponse() (response *ListPublicMaterialInfosResponse) {
	response = &ListPublicMaterialInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取公共市场物料，可用于多云中选择部署的物料信息
func (c *Client) ListPublicMaterialInfos(request *ListPublicMaterialInfosRequest) (response *ListPublicMaterialInfosResponse, err error) {
	if request == nil {
		request = NewListPublicMaterialInfosRequest()
	}
	response = NewListPublicMaterialInfosResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductInstanceTreeRequest() (request *DescribeProductInstanceTreeRequest) {
	request = &DescribeProductInstanceTreeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeProductInstanceTree")
	return
}

func NewDescribeProductInstanceTreeResponse() (response *DescribeProductInstanceTreeResponse) {
	response = &DescribeProductInstanceTreeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取产品实例树，包括：产品实例，子系统实例，应用实例
func (c *Client) DescribeProductInstanceTree(request *DescribeProductInstanceTreeRequest) (response *DescribeProductInstanceTreeResponse, err error) {
	if request == nil {
		request = NewDescribeProductInstanceTreeRequest()
	}
	response = NewDescribeProductInstanceTreeResponse()
	err = c.Send(request, response)
	return
}

func NewGetProgressImportTaskRequest() (request *GetProgressImportTaskRequest) {
	request = &GetProgressImportTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetProgressImportTask")
	return
}

func NewGetProgressImportTaskResponse() (response *GetProgressImportTaskResponse) {
	response = &GetProgressImportTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取正在进行中的导入任务
func (c *Client) GetProgressImportTask(request *GetProgressImportTaskRequest) (response *GetProgressImportTaskResponse, err error) {
	if request == nil {
		request = NewGetProgressImportTaskRequest()
	}
	response = NewGetProgressImportTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCheckActiveSiteRequest() (request *CheckActiveSiteRequest) {
	request = &CheckActiveSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CheckActiveSite")
	return
}

func NewCheckActiveSiteResponse() (response *CheckActiveSiteResponse) {
	response = &CheckActiveSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查指定局点是否是当前环境生效局点
func (c *Client) CheckActiveSite(request *CheckActiveSiteRequest) (response *CheckActiveSiteResponse, err error) {
	if request == nil {
		request = NewCheckActiveSiteRequest()
	}
	response = NewCheckActiveSiteResponse()
	err = c.Send(request, response)
	return
}

func NewCheckSheetTemplateRequest() (request *CheckSheetTemplateRequest) {
	request = &CheckSheetTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CheckSheetTemplate")
	return
}

func NewCheckSheetTemplateResponse() (response *CheckSheetTemplateResponse) {
	response = &CheckSheetTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查实例运维单编排模板
func (c *Client) CheckSheetTemplate(request *CheckSheetTemplateRequest) (response *CheckSheetTemplateResponse, err error) {
	if request == nil {
		request = NewCheckSheetTemplateRequest()
	}
	response = NewCheckSheetTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewCheckSiteOperationSheetMaterialRequest() (request *CheckSiteOperationSheetMaterialRequest) {
	request = &CheckSiteOperationSheetMaterialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CheckSiteOperationSheetMaterial")
	return
}

func NewCheckSiteOperationSheetMaterialResponse() (response *CheckSiteOperationSheetMaterialResponse) {
	response = &CheckSiteOperationSheetMaterialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查Site变更单中应用的物料是否存在
func (c *Client) CheckSiteOperationSheetMaterial(request *CheckSiteOperationSheetMaterialRequest) (response *CheckSiteOperationSheetMaterialResponse, err error) {
	if request == nil {
		request = NewCheckSiteOperationSheetMaterialRequest()
	}
	response = NewCheckSiteOperationSheetMaterialResponse()
	err = c.Send(request, response)
	return
}

func NewPublishCommonOperationSheetRequest() (request *PublishCommonOperationSheetRequest) {
	request = &PublishCommonOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "PublishCommonOperationSheet")
	return
}

func NewPublishCommonOperationSheetResponse() (response *PublishCommonOperationSheetResponse) {
	response = &PublishCommonOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关联通用变更单父子关系
func (c *Client) PublishCommonOperationSheet(request *PublishCommonOperationSheetRequest) (response *PublishCommonOperationSheetResponse, err error) {
	if request == nil {
		request = NewPublishCommonOperationSheetRequest()
	}
	response = NewPublishCommonOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewCopySiteRequest() (request *CopySiteRequest) {
	request = &CopySiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "CopySite")
	return
}

func NewCopySiteResponse() (response *CopySiteResponse) {
	response = &CopySiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在多云管理中用于复制云
func (c *Client) CopySite(request *CopySiteRequest) (response *CopySiteResponse, err error) {
	if request == nil {
		request = NewCopySiteRequest()
	}
	response = NewCopySiteResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationPackageDetailRequest() (request *GetApplicationPackageDetailRequest) {
	request = &GetApplicationPackageDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "GetApplicationPackageDetail")
	return
}

func NewGetApplicationPackageDetailResponse() (response *GetApplicationPackageDetailResponse) {
	response = &GetApplicationPackageDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用包详情
func (c *Client) GetApplicationPackageDetail(request *GetApplicationPackageDetailRequest) (response *GetApplicationPackageDetailResponse, err error) {
	if request == nil {
		request = NewGetApplicationPackageDetailRequest()
	}
	response = NewGetApplicationPackageDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAtomicRenderDataRequest() (request *DescribeAtomicRenderDataRequest) {
	request = &DescribeAtomicRenderDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DescribeAtomicRenderData")
	return
}

func NewDescribeAtomicRenderDataResponse() (response *DescribeAtomicRenderDataResponse) {
	response = &DescribeAtomicRenderDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询原子操作渲染数据
func (c *Client) DescribeAtomicRenderData(request *DescribeAtomicRenderDataRequest) (response *DescribeAtomicRenderDataResponse, err error) {
	if request == nil {
		request = NewDescribeAtomicRenderDataRequest()
	}
	response = NewDescribeAtomicRenderDataResponse()
	err = c.Send(request, response)
	return
}

func NewInstantiateCommonOperationSheetRequest() (request *InstantiateCommonOperationSheetRequest) {
	request = &InstantiateCommonOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "InstantiateCommonOperationSheet")
	return
}

func NewInstantiateCommonOperationSheetResponse() (response *InstantiateCommonOperationSheetResponse) {
	response = &InstantiateCommonOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 实例化通用变更单
func (c *Client) InstantiateCommonOperationSheet(request *InstantiateCommonOperationSheetRequest) (response *InstantiateCommonOperationSheetResponse, err error) {
	if request == nil {
		request = NewInstantiateCommonOperationSheetRequest()
	}
	response = NewInstantiateCommonOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewExportSolutionTemplateRequest() (request *ExportSolutionTemplateRequest) {
	request = &ExportSolutionTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ExportSolutionTemplate")
	return
}

func NewExportSolutionTemplateResponse() (response *ExportSolutionTemplateResponse) {
	response = &ExportSolutionTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出解决方案运维模板
func (c *Client) ExportSolutionTemplate(request *ExportSolutionTemplateRequest) (response *ExportSolutionTemplateResponse, err error) {
	if request == nil {
		request = NewExportSolutionTemplateRequest()
	}
	response = NewExportSolutionTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDeploySingleAppRequest() (request *DeploySingleAppRequest) {
	request = &DeploySingleAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "DeploySingleApp")
	return
}

func NewDeploySingleAppResponse() (response *DeploySingleAppResponse) {
	response = &DeploySingleAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从应用列表部署单个app、tool、atomic
func (c *Client) DeploySingleApp(request *DeploySingleAppRequest) (response *DeploySingleAppResponse, err error) {
	if request == nil {
		request = NewDeploySingleAppRequest()
	}
	response = NewDeploySingleAppResponse()
	err = c.Send(request, response)
	return
}

func NewListAppRuntimeNamesRequest() (request *ListAppRuntimeNamesRequest) {
	request = &ListAppRuntimeNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "ListAppRuntimeNames")
	return
}

func NewListAppRuntimeNamesResponse() (response *ListAppRuntimeNamesResponse) {
	response = &ListAppRuntimeNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用运行时名称
func (c *Client) ListAppRuntimeNames(request *ListAppRuntimeNamesRequest) (response *ListAppRuntimeNamesResponse, err error) {
	if request == nil {
		request = NewListAppRuntimeNamesRequest()
	}
	response = NewListAppRuntimeNamesResponse()
	err = c.Send(request, response)
	return
}

func NewRelateCommonOperationSheetRequest() (request *RelateCommonOperationSheetRequest) {
	request = &RelateCommonOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dawnservices", APIVersion, "RelateCommonOperationSheet")
	return
}

func NewRelateCommonOperationSheetResponse() (response *RelateCommonOperationSheetResponse) {
	response = &RelateCommonOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关联通用变更单父子关系
func (c *Client) RelateCommonOperationSheet(request *RelateCommonOperationSheetRequest) (response *RelateCommonOperationSheetResponse, err error) {
	if request == nil {
		request = NewRelateCommonOperationSheetRequest()
	}
	response = NewRelateCommonOperationSheetResponse()
	err = c.Send(request, response)
	return
}
