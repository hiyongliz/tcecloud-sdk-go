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

func NewGetApplicationPackageCcDeclareRequest() (request *GetApplicationPackageCcDeclareRequest) {
	request = &GetApplicationPackageCcDeclareRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetApplicationPackageCcDeclare")
	return
}

func NewGetApplicationPackageCcDeclareResponse() (response *GetApplicationPackageCcDeclareResponse) {
	response = &GetApplicationPackageCcDeclareResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetApplicationPackageCcDeclare(request *GetApplicationPackageCcDeclareRequest) (response *GetApplicationPackageCcDeclareResponse, err error) {
	if request == nil {
		request = NewGetApplicationPackageCcDeclareRequest()
	}
	response = NewGetApplicationPackageCcDeclareResponse()
	err = c.Send(request, response)
	return
}

func NewGetAtomicRequest() (request *GetAtomicRequest) {
	request = &GetAtomicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetAtomic")
	return
}

func NewGetAtomicResponse() (response *GetAtomicResponse) {
	response = &GetAtomicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetAtomic(request *GetAtomicRequest) (response *GetAtomicResponse, err error) {
	if request == nil {
		request = NewGetAtomicRequest()
	}
	response = NewGetAtomicResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSiteDeployTaskRequest() (request *CreateSiteDeployTaskRequest) {
	request = &CreateSiteDeployTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateSiteDeployTask")
	return
}

func NewCreateSiteDeployTaskResponse() (response *CreateSiteDeployTaskResponse) {
	response = &CreateSiteDeployTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateSiteDeployTask(request *CreateSiteDeployTaskRequest) (response *CreateSiteDeployTaskResponse, err error) {
	if request == nil {
		request = NewCreateSiteDeployTaskRequest()
	}
	response = NewCreateSiteDeployTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSiteOperationSheetRequest() (request *DeleteSiteOperationSheetRequest) {
	request = &DeleteSiteOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteSiteOperationSheet")
	return
}

func NewDeleteSiteOperationSheetResponse() (response *DeleteSiteOperationSheetResponse) {
	response = &DeleteSiteOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteSiteOperationSheet(request *DeleteSiteOperationSheetRequest) (response *DeleteSiteOperationSheetResponse, err error) {
	if request == nil {
		request = NewDeleteSiteOperationSheetRequest()
	}
	response = NewDeleteSiteOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApplicationInstanceStateRequest() (request *DeleteApplicationInstanceStateRequest) {
	request = &DeleteApplicationInstanceStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteApplicationInstanceState")
	return
}

func NewDeleteApplicationInstanceStateResponse() (response *DeleteApplicationInstanceStateResponse) {
	response = &DeleteApplicationInstanceStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteApplicationInstanceState(request *DeleteApplicationInstanceStateRequest) (response *DeleteApplicationInstanceStateResponse, err error) {
	if request == nil {
		request = NewDeleteApplicationInstanceStateRequest()
	}
	response = NewDeleteApplicationInstanceStateResponse()
	err = c.Send(request, response)
	return
}

func NewListApplicationRunningDataCollectionsRequest() (request *ListApplicationRunningDataCollectionsRequest) {
	request = &ListApplicationRunningDataCollectionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListApplicationRunningDataCollections")
	return
}

func NewListApplicationRunningDataCollectionsResponse() (response *ListApplicationRunningDataCollectionsResponse) {
	response = &ListApplicationRunningDataCollectionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListApplicationRunningDataCollections(request *ListApplicationRunningDataCollectionsRequest) (response *ListApplicationRunningDataCollectionsResponse, err error) {
	if request == nil {
		request = NewListApplicationRunningDataCollectionsRequest()
	}
	response = NewListApplicationRunningDataCollectionsResponse()
	err = c.Send(request, response)
	return
}

func NewGetPlanRequest() (request *GetPlanRequest) {
	request = &GetPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetPlan")
	return
}

func NewGetPlanResponse() (response *GetPlanResponse) {
	response = &GetPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetPlan(request *GetPlanRequest) (response *GetPlanResponse, err error) {
	if request == nil {
		request = NewGetPlanRequest()
	}
	response = NewGetPlanResponse()
	err = c.Send(request, response)
	return
}

func NewGetSolutionVersionRequest() (request *GetSolutionVersionRequest) {
	request = &GetSolutionVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetSolutionVersion")
	return
}

func NewGetSolutionVersionResponse() (response *GetSolutionVersionResponse) {
	response = &GetSolutionVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetSolutionVersion(request *GetSolutionVersionRequest) (response *GetSolutionVersionResponse, err error) {
	if request == nil {
		request = NewGetSolutionVersionRequest()
	}
	response = NewGetSolutionVersionResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePlanGlobalConfigRequest() (request *UpdatePlanGlobalConfigRequest) {
	request = &UpdatePlanGlobalConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdatePlanGlobalConfig")
	return
}

func NewUpdatePlanGlobalConfigResponse() (response *UpdatePlanGlobalConfigResponse) {
	response = &UpdatePlanGlobalConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdatePlanGlobalConfig(request *UpdatePlanGlobalConfigRequest) (response *UpdatePlanGlobalConfigResponse, err error) {
	if request == nil {
		request = NewUpdatePlanGlobalConfigRequest()
	}
	response = NewUpdatePlanGlobalConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApplicationArtifactChartHistoryRequest() (request *CreateApplicationArtifactChartHistoryRequest) {
	request = &CreateApplicationArtifactChartHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateApplicationArtifactChartHistory")
	return
}

func NewCreateApplicationArtifactChartHistoryResponse() (response *CreateApplicationArtifactChartHistoryResponse) {
	response = &CreateApplicationArtifactChartHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateApplicationArtifactChartHistory(request *CreateApplicationArtifactChartHistoryRequest) (response *CreateApplicationArtifactChartHistoryResponse, err error) {
	if request == nil {
		request = NewCreateApplicationArtifactChartHistoryRequest()
	}
	response = NewCreateApplicationArtifactChartHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaterialCustomizedTypeRequest() (request *DeleteMaterialCustomizedTypeRequest) {
	request = &DeleteMaterialCustomizedTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteMaterialCustomizedType")
	return
}

func NewDeleteMaterialCustomizedTypeResponse() (response *DeleteMaterialCustomizedTypeResponse) {
	response = &DeleteMaterialCustomizedTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteMaterialCustomizedType(request *DeleteMaterialCustomizedTypeRequest) (response *DeleteMaterialCustomizedTypeResponse, err error) {
	if request == nil {
		request = NewDeleteMaterialCustomizedTypeRequest()
	}
	response = NewDeleteMaterialCustomizedTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaterialCleanupRequest() (request *DeleteMaterialCleanupRequest) {
	request = &DeleteMaterialCleanupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteMaterialCleanup")
	return
}

func NewDeleteMaterialCleanupResponse() (response *DeleteMaterialCleanupResponse) {
	response = &DeleteMaterialCleanupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteMaterialCleanup(request *DeleteMaterialCleanupRequest) (response *DeleteMaterialCleanupResponse, err error) {
	if request == nil {
		request = NewDeleteMaterialCleanupRequest()
	}
	response = NewDeleteMaterialCleanupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProjectTopologyRequest() (request *DeleteProjectTopologyRequest) {
	request = &DeleteProjectTopologyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteProjectTopology")
	return
}

func NewDeleteProjectTopologyResponse() (response *DeleteProjectTopologyResponse) {
	response = &DeleteProjectTopologyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteProjectTopology(request *DeleteProjectTopologyRequest) (response *DeleteProjectTopologyResponse, err error) {
	if request == nil {
		request = NewDeleteProjectTopologyRequest()
	}
	response = NewDeleteProjectTopologyResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateOperationSheetTemplatesRequest() (request *BulkCreateOperationSheetTemplatesRequest) {
	request = &BulkCreateOperationSheetTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateOperationSheetTemplates")
	return
}

func NewBulkCreateOperationSheetTemplatesResponse() (response *BulkCreateOperationSheetTemplatesResponse) {
	response = &BulkCreateOperationSheetTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateOperationSheetTemplates(request *BulkCreateOperationSheetTemplatesRequest) (response *BulkCreateOperationSheetTemplatesResponse, err error) {
	if request == nil {
		request = NewBulkCreateOperationSheetTemplatesRequest()
	}
	response = NewBulkCreateOperationSheetTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateSolutionsRequest() (request *BulkCreateSolutionsRequest) {
	request = &BulkCreateSolutionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateSolutions")
	return
}

func NewBulkCreateSolutionsResponse() (response *BulkCreateSolutionsResponse) {
	response = &BulkCreateSolutionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateSolutions(request *BulkCreateSolutionsRequest) (response *BulkCreateSolutionsResponse, err error) {
	if request == nil {
		request = NewBulkCreateSolutionsRequest()
	}
	response = NewBulkCreateSolutionsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperationSheetLabelRequest() (request *DeleteOperationSheetLabelRequest) {
	request = &DeleteOperationSheetLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteOperationSheetLabel")
	return
}

func NewDeleteOperationSheetLabelResponse() (response *DeleteOperationSheetLabelResponse) {
	response = &DeleteOperationSheetLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteOperationSheetLabel(request *DeleteOperationSheetLabelRequest) (response *DeleteOperationSheetLabelResponse, err error) {
	if request == nil {
		request = NewDeleteOperationSheetLabelRequest()
	}
	response = NewDeleteOperationSheetLabelResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductInfoRequest() (request *GetProductInfoRequest) {
	request = &GetProductInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProductInfo")
	return
}

func NewGetProductInfoResponse() (response *GetProductInfoResponse) {
	response = &GetProductInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProductInfo(request *GetProductInfoRequest) (response *GetProductInfoResponse, err error) {
	if request == nil {
		request = NewGetProductInfoRequest()
	}
	response = NewGetProductInfoResponse()
	err = c.Send(request, response)
	return
}

func NewListProductOperationsRequest() (request *ListProductOperationsRequest) {
	request = &ListProductOperationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProductOperations")
	return
}

func NewListProductOperationsResponse() (response *ListProductOperationsResponse) {
	response = &ListProductOperationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProductOperations(request *ListProductOperationsRequest) (response *ListProductOperationsResponse, err error) {
	if request == nil {
		request = NewListProductOperationsRequest()
	}
	response = NewListProductOperationsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreatePlanCustomConfigsRequest() (request *BulkCreatePlanCustomConfigsRequest) {
	request = &BulkCreatePlanCustomConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreatePlanCustomConfigs")
	return
}

func NewBulkCreatePlanCustomConfigsResponse() (response *BulkCreatePlanCustomConfigsResponse) {
	response = &BulkCreatePlanCustomConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreatePlanCustomConfigs(request *BulkCreatePlanCustomConfigsRequest) (response *BulkCreatePlanCustomConfigsResponse, err error) {
	if request == nil {
		request = NewBulkCreatePlanCustomConfigsRequest()
	}
	response = NewBulkCreatePlanCustomConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewListApplicationRunningPlansRequest() (request *ListApplicationRunningPlansRequest) {
	request = &ListApplicationRunningPlansRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListApplicationRunningPlans")
	return
}

func NewListApplicationRunningPlansResponse() (response *ListApplicationRunningPlansResponse) {
	response = &ListApplicationRunningPlansResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListApplicationRunningPlans(request *ListApplicationRunningPlansRequest) (response *ListApplicationRunningPlansResponse, err error) {
	if request == nil {
		request = NewListApplicationRunningPlansRequest()
	}
	response = NewListApplicationRunningPlansResponse()
	err = c.Send(request, response)
	return
}

func NewListProjectSiteInfoPlansRequest() (request *ListProjectSiteInfoPlansRequest) {
	request = &ListProjectSiteInfoPlansRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProjectSiteInfoPlans")
	return
}

func NewListProjectSiteInfoPlansResponse() (response *ListProjectSiteInfoPlansResponse) {
	response = &ListProjectSiteInfoPlansResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProjectSiteInfoPlans(request *ListProjectSiteInfoPlansRequest) (response *ListProjectSiteInfoPlansResponse, err error) {
	if request == nil {
		request = NewListProjectSiteInfoPlansRequest()
	}
	response = NewListProjectSiteInfoPlansResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAtomicRequest() (request *UpdateAtomicRequest) {
	request = &UpdateAtomicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateAtomic")
	return
}

func NewUpdateAtomicResponse() (response *UpdateAtomicResponse) {
	response = &UpdateAtomicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateAtomic(request *UpdateAtomicRequest) (response *UpdateAtomicResponse, err error) {
	if request == nil {
		request = NewUpdateAtomicRequest()
	}
	response = NewUpdateAtomicResponse()
	err = c.Send(request, response)
	return
}

func NewGetPlanCustomConfigRequest() (request *GetPlanCustomConfigRequest) {
	request = &GetPlanCustomConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetPlanCustomConfig")
	return
}

func NewGetPlanCustomConfigResponse() (response *GetPlanCustomConfigResponse) {
	response = &GetPlanCustomConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetPlanCustomConfig(request *GetPlanCustomConfigRequest) (response *GetPlanCustomConfigResponse, err error) {
	if request == nil {
		request = NewGetPlanCustomConfigRequest()
	}
	response = NewGetPlanCustomConfigResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateMaterialCleanupsRequest() (request *BulkCreateMaterialCleanupsRequest) {
	request = &BulkCreateMaterialCleanupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateMaterialCleanups")
	return
}

func NewBulkCreateMaterialCleanupsResponse() (response *BulkCreateMaterialCleanupsResponse) {
	response = &BulkCreateMaterialCleanupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateMaterialCleanups(request *BulkCreateMaterialCleanupsRequest) (response *BulkCreateMaterialCleanupsResponse, err error) {
	if request == nil {
		request = NewBulkCreateMaterialCleanupsRequest()
	}
	response = NewBulkCreateMaterialCleanupsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSiteDeployTaskRequest() (request *UpdateSiteDeployTaskRequest) {
	request = &UpdateSiteDeployTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateSiteDeployTask")
	return
}

func NewUpdateSiteDeployTaskResponse() (response *UpdateSiteDeployTaskResponse) {
	response = &UpdateSiteDeployTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateSiteDeployTask(request *UpdateSiteDeployTaskRequest) (response *UpdateSiteDeployTaskResponse, err error) {
	if request == nil {
		request = NewUpdateSiteDeployTaskRequest()
	}
	response = NewUpdateSiteDeployTaskResponse()
	err = c.Send(request, response)
	return
}

func NewListProjectTopologiesRequest() (request *ListProjectTopologiesRequest) {
	request = &ListProjectTopologiesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProjectTopologies")
	return
}

func NewListProjectTopologiesResponse() (response *ListProjectTopologiesResponse) {
	response = &ListProjectTopologiesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProjectTopologies(request *ListProjectTopologiesRequest) (response *ListProjectTopologiesResponse, err error) {
	if request == nil {
		request = NewListProjectTopologiesRequest()
	}
	response = NewListProjectTopologiesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApplicationPackageCcDeclareRequest() (request *UpdateApplicationPackageCcDeclareRequest) {
	request = &UpdateApplicationPackageCcDeclareRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateApplicationPackageCcDeclare")
	return
}

func NewUpdateApplicationPackageCcDeclareResponse() (response *UpdateApplicationPackageCcDeclareResponse) {
	response = &UpdateApplicationPackageCcDeclareResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateApplicationPackageCcDeclare(request *UpdateApplicationPackageCcDeclareRequest) (response *UpdateApplicationPackageCcDeclareResponse, err error) {
	if request == nil {
		request = NewUpdateApplicationPackageCcDeclareRequest()
	}
	response = NewUpdateApplicationPackageCcDeclareResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperationPipelineRequest() (request *CreateOperationPipelineRequest) {
	request = &CreateOperationPipelineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateOperationPipeline")
	return
}

func NewCreateOperationPipelineResponse() (response *CreateOperationPipelineResponse) {
	response = &CreateOperationPipelineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateOperationPipeline(request *CreateOperationPipelineRequest) (response *CreateOperationPipelineResponse, err error) {
	if request == nil {
		request = NewCreateOperationPipelineRequest()
	}
	response = NewCreateOperationPipelineResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApplicationInstanceBackupRequest() (request *UpdateApplicationInstanceBackupRequest) {
	request = &UpdateApplicationInstanceBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateApplicationInstanceBackup")
	return
}

func NewUpdateApplicationInstanceBackupResponse() (response *UpdateApplicationInstanceBackupResponse) {
	response = &UpdateApplicationInstanceBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateApplicationInstanceBackup(request *UpdateApplicationInstanceBackupRequest) (response *UpdateApplicationInstanceBackupResponse, err error) {
	if request == nil {
		request = NewUpdateApplicationInstanceBackupRequest()
	}
	response = NewUpdateApplicationInstanceBackupResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateApplicationArtifactsRequest() (request *BulkCreateApplicationArtifactsRequest) {
	request = &BulkCreateApplicationArtifactsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateApplicationArtifacts")
	return
}

func NewBulkCreateApplicationArtifactsResponse() (response *BulkCreateApplicationArtifactsResponse) {
	response = &BulkCreateApplicationArtifactsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateApplicationArtifacts(request *BulkCreateApplicationArtifactsRequest) (response *BulkCreateApplicationArtifactsResponse, err error) {
	if request == nil {
		request = NewBulkCreateApplicationArtifactsRequest()
	}
	response = NewBulkCreateApplicationArtifactsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSystemSettingRequest() (request *UpdateSystemSettingRequest) {
	request = &UpdateSystemSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateSystemSetting")
	return
}

func NewUpdateSystemSettingResponse() (response *UpdateSystemSettingResponse) {
	response = &UpdateSystemSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateSystemSetting(request *UpdateSystemSettingRequest) (response *UpdateSystemSettingResponse, err error) {
	if request == nil {
		request = NewUpdateSystemSettingRequest()
	}
	response = NewUpdateSystemSettingResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateMaterialInfosRequest() (request *BulkCreateMaterialInfosRequest) {
	request = &BulkCreateMaterialInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateMaterialInfos")
	return
}

func NewBulkCreateMaterialInfosResponse() (response *BulkCreateMaterialInfosResponse) {
	response = &BulkCreateMaterialInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateMaterialInfos(request *BulkCreateMaterialInfosRequest) (response *BulkCreateMaterialInfosResponse, err error) {
	if request == nil {
		request = NewBulkCreateMaterialInfosRequest()
	}
	response = NewBulkCreateMaterialInfosResponse()
	err = c.Send(request, response)
	return
}

func NewListPlanApplicationInstancesRequest() (request *ListPlanApplicationInstancesRequest) {
	request = &ListPlanApplicationInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListPlanApplicationInstances")
	return
}

func NewListPlanApplicationInstancesResponse() (response *ListPlanApplicationInstancesResponse) {
	response = &ListPlanApplicationInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListPlanApplicationInstances(request *ListPlanApplicationInstancesRequest) (response *ListPlanApplicationInstancesResponse, err error) {
	if request == nil {
		request = NewListPlanApplicationInstancesRequest()
	}
	response = NewListPlanApplicationInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewListPlanGlobalConfigsRequest() (request *ListPlanGlobalConfigsRequest) {
	request = &ListPlanGlobalConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListPlanGlobalConfigs")
	return
}

func NewListPlanGlobalConfigsResponse() (response *ListPlanGlobalConfigsResponse) {
	response = &ListPlanGlobalConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListPlanGlobalConfigs(request *ListPlanGlobalConfigsRequest) (response *ListPlanGlobalConfigsResponse, err error) {
	if request == nil {
		request = NewListPlanGlobalConfigsRequest()
	}
	response = NewListPlanGlobalConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreatePlanGlobalConfigsRequest() (request *BulkCreatePlanGlobalConfigsRequest) {
	request = &BulkCreatePlanGlobalConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreatePlanGlobalConfigs")
	return
}

func NewBulkCreatePlanGlobalConfigsResponse() (response *BulkCreatePlanGlobalConfigsResponse) {
	response = &BulkCreatePlanGlobalConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreatePlanGlobalConfigs(request *BulkCreatePlanGlobalConfigsRequest) (response *BulkCreatePlanGlobalConfigsResponse, err error) {
	if request == nil {
		request = NewBulkCreatePlanGlobalConfigsRequest()
	}
	response = NewBulkCreatePlanGlobalConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewListPlansRequest() (request *ListPlansRequest) {
	request = &ListPlansRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListPlans")
	return
}

func NewListPlansResponse() (response *ListPlansResponse) {
	response = &ListPlansResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListPlans(request *ListPlansRequest) (response *ListPlansResponse, err error) {
	if request == nil {
		request = NewListPlansRequest()
	}
	response = NewListPlansResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProjectRegionRequest() (request *UpdateProjectRegionRequest) {
	request = &UpdateProjectRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateProjectRegion")
	return
}

func NewUpdateProjectRegionResponse() (response *UpdateProjectRegionResponse) {
	response = &UpdateProjectRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateProjectRegion(request *UpdateProjectRegionRequest) (response *UpdateProjectRegionResponse, err error) {
	if request == nil {
		request = NewUpdateProjectRegionRequest()
	}
	response = NewUpdateProjectRegionResponse()
	err = c.Send(request, response)
	return
}

func NewGetPlanTagRequest() (request *GetPlanTagRequest) {
	request = &GetPlanTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetPlanTag")
	return
}

func NewGetPlanTagResponse() (response *GetPlanTagResponse) {
	response = &GetPlanTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetPlanTag(request *GetPlanTagRequest) (response *GetPlanTagResponse, err error) {
	if request == nil {
		request = NewGetPlanTagRequest()
	}
	response = NewGetPlanTagResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMaterialMachineSSHInfoRequest() (request *CreateMaterialMachineSSHInfoRequest) {
	request = &CreateMaterialMachineSSHInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateMaterialMachineSSHInfo")
	return
}

func NewCreateMaterialMachineSSHInfoResponse() (response *CreateMaterialMachineSSHInfoResponse) {
	response = &CreateMaterialMachineSSHInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateMaterialMachineSSHInfo(request *CreateMaterialMachineSSHInfoRequest) (response *CreateMaterialMachineSSHInfoResponse, err error) {
	if request == nil {
		request = NewCreateMaterialMachineSSHInfoRequest()
	}
	response = NewCreateMaterialMachineSSHInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePlanSubsystemInstanceRequest() (request *DeletePlanSubsystemInstanceRequest) {
	request = &DeletePlanSubsystemInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeletePlanSubsystemInstance")
	return
}

func NewDeletePlanSubsystemInstanceResponse() (response *DeletePlanSubsystemInstanceResponse) {
	response = &DeletePlanSubsystemInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeletePlanSubsystemInstance(request *DeletePlanSubsystemInstanceRequest) (response *DeletePlanSubsystemInstanceResponse, err error) {
	if request == nil {
		request = NewDeletePlanSubsystemInstanceRequest()
	}
	response = NewDeletePlanSubsystemInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateProductInstanceStatesRequest() (request *BulkCreateProductInstanceStatesRequest) {
	request = &BulkCreateProductInstanceStatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateProductInstanceStates")
	return
}

func NewBulkCreateProductInstanceStatesResponse() (response *BulkCreateProductInstanceStatesResponse) {
	response = &BulkCreateProductInstanceStatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateProductInstanceStates(request *BulkCreateProductInstanceStatesRequest) (response *BulkCreateProductInstanceStatesResponse, err error) {
	if request == nil {
		request = NewBulkCreateProductInstanceStatesRequest()
	}
	response = NewBulkCreateProductInstanceStatesResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePlanServerInfoRequest() (request *CreatePlanServerInfoRequest) {
	request = &CreatePlanServerInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreatePlanServerInfo")
	return
}

func NewCreatePlanServerInfoResponse() (response *CreatePlanServerInfoResponse) {
	response = &CreatePlanServerInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreatePlanServerInfo(request *CreatePlanServerInfoRequest) (response *CreatePlanServerInfoResponse, err error) {
	if request == nil {
		request = NewCreatePlanServerInfoRequest()
	}
	response = NewCreatePlanServerInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProjectTopologyDetailRequest() (request *CreateProjectTopologyDetailRequest) {
	request = &CreateProjectTopologyDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateProjectTopologyDetail")
	return
}

func NewCreateProjectTopologyDetailResponse() (response *CreateProjectTopologyDetailResponse) {
	response = &CreateProjectTopologyDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateProjectTopologyDetail(request *CreateProjectTopologyDetailRequest) (response *CreateProjectTopologyDetailResponse, err error) {
	if request == nil {
		request = NewCreateProjectTopologyDetailRequest()
	}
	response = NewCreateProjectTopologyDetailResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreatePlanApplicationServerRelationsRequest() (request *BulkCreatePlanApplicationServerRelationsRequest) {
	request = &BulkCreatePlanApplicationServerRelationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreatePlanApplicationServerRelations")
	return
}

func NewBulkCreatePlanApplicationServerRelationsResponse() (response *BulkCreatePlanApplicationServerRelationsResponse) {
	response = &BulkCreatePlanApplicationServerRelationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreatePlanApplicationServerRelations(request *BulkCreatePlanApplicationServerRelationsRequest) (response *BulkCreatePlanApplicationServerRelationsResponse, err error) {
	if request == nil {
		request = NewBulkCreatePlanApplicationServerRelationsRequest()
	}
	response = NewBulkCreatePlanApplicationServerRelationsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePlanApplicationServerRelationRequest() (request *UpdatePlanApplicationServerRelationRequest) {
	request = &UpdatePlanApplicationServerRelationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdatePlanApplicationServerRelation")
	return
}

func NewUpdatePlanApplicationServerRelationResponse() (response *UpdatePlanApplicationServerRelationResponse) {
	response = &UpdatePlanApplicationServerRelationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdatePlanApplicationServerRelation(request *UpdatePlanApplicationServerRelationRequest) (response *UpdatePlanApplicationServerRelationResponse, err error) {
	if request == nil {
		request = NewUpdatePlanApplicationServerRelationRequest()
	}
	response = NewUpdatePlanApplicationServerRelationResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePlanCustomConfigRequest() (request *UpdatePlanCustomConfigRequest) {
	request = &UpdatePlanCustomConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdatePlanCustomConfig")
	return
}

func NewUpdatePlanCustomConfigResponse() (response *UpdatePlanCustomConfigResponse) {
	response = &UpdatePlanCustomConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdatePlanCustomConfig(request *UpdatePlanCustomConfigRequest) (response *UpdatePlanCustomConfigResponse, err error) {
	if request == nil {
		request = NewUpdatePlanCustomConfigRequest()
	}
	response = NewUpdatePlanCustomConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProjectSiteInfoRequest() (request *CreateProjectSiteInfoRequest) {
	request = &CreateProjectSiteInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateProjectSiteInfo")
	return
}

func NewCreateProjectSiteInfoResponse() (response *CreateProjectSiteInfoResponse) {
	response = &CreateProjectSiteInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateProjectSiteInfo(request *CreateProjectSiteInfoRequest) (response *CreateProjectSiteInfoResponse, err error) {
	if request == nil {
		request = NewCreateProjectSiteInfoRequest()
	}
	response = NewCreateProjectSiteInfoResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOperationPipelineRequest() (request *UpdateOperationPipelineRequest) {
	request = &UpdateOperationPipelineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateOperationPipeline")
	return
}

func NewUpdateOperationPipelineResponse() (response *UpdateOperationPipelineResponse) {
	response = &UpdateOperationPipelineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateOperationPipeline(request *UpdateOperationPipelineRequest) (response *UpdateOperationPipelineResponse, err error) {
	if request == nil {
		request = NewUpdateOperationPipelineRequest()
	}
	response = NewUpdateOperationPipelineResponse()
	err = c.Send(request, response)
	return
}

func NewListProjectSiteInfoDetailsRequest() (request *ListProjectSiteInfoDetailsRequest) {
	request = &ListProjectSiteInfoDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProjectSiteInfoDetails")
	return
}

func NewListProjectSiteInfoDetailsResponse() (response *ListProjectSiteInfoDetailsResponse) {
	response = &ListProjectSiteInfoDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProjectSiteInfoDetails(request *ListProjectSiteInfoDetailsRequest) (response *ListProjectSiteInfoDetailsResponse, err error) {
	if request == nil {
		request = NewListProjectSiteInfoDetailsRequest()
	}
	response = NewListProjectSiteInfoDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMaterialSiteRequest() (request *UpdateMaterialSiteRequest) {
	request = &UpdateMaterialSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateMaterialSite")
	return
}

func NewUpdateMaterialSiteResponse() (response *UpdateMaterialSiteResponse) {
	response = &UpdateMaterialSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateMaterialSite(request *UpdateMaterialSiteRequest) (response *UpdateMaterialSiteResponse, err error) {
	if request == nil {
		request = NewUpdateMaterialSiteRequest()
	}
	response = NewUpdateMaterialSiteResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationArtifactChartHistoryRequest() (request *GetApplicationArtifactChartHistoryRequest) {
	request = &GetApplicationArtifactChartHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetApplicationArtifactChartHistory")
	return
}

func NewGetApplicationArtifactChartHistoryResponse() (response *GetApplicationArtifactChartHistoryResponse) {
	response = &GetApplicationArtifactChartHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetApplicationArtifactChartHistory(request *GetApplicationArtifactChartHistoryRequest) (response *GetApplicationArtifactChartHistoryResponse, err error) {
	if request == nil {
		request = NewGetApplicationArtifactChartHistoryRequest()
	}
	response = NewGetApplicationArtifactChartHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialMetadatasRequest() (request *ListMaterialMetadatasRequest) {
	request = &ListMaterialMetadatasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListMaterialMetadatas")
	return
}

func NewListMaterialMetadatasResponse() (response *ListMaterialMetadatasResponse) {
	response = &ListMaterialMetadatasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListMaterialMetadatas(request *ListMaterialMetadatasRequest) (response *ListMaterialMetadatasResponse, err error) {
	if request == nil {
		request = NewListMaterialMetadatasRequest()
	}
	response = NewListMaterialMetadatasResponse()
	err = c.Send(request, response)
	return
}

func NewGetMaterialMetadataRequest() (request *GetMaterialMetadataRequest) {
	request = &GetMaterialMetadataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetMaterialMetadata")
	return
}

func NewGetMaterialMetadataResponse() (response *GetMaterialMetadataResponse) {
	response = &GetMaterialMetadataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetMaterialMetadata(request *GetMaterialMetadataRequest) (response *GetMaterialMetadataResponse, err error) {
	if request == nil {
		request = NewGetMaterialMetadataRequest()
	}
	response = NewGetMaterialMetadataResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateApplicationRunningDataCollectionsRequest() (request *BulkCreateApplicationRunningDataCollectionsRequest) {
	request = &BulkCreateApplicationRunningDataCollectionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateApplicationRunningDataCollections")
	return
}

func NewBulkCreateApplicationRunningDataCollectionsResponse() (response *BulkCreateApplicationRunningDataCollectionsResponse) {
	response = &BulkCreateApplicationRunningDataCollectionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateApplicationRunningDataCollections(request *BulkCreateApplicationRunningDataCollectionsRequest) (response *BulkCreateApplicationRunningDataCollectionsResponse, err error) {
	if request == nil {
		request = NewBulkCreateApplicationRunningDataCollectionsRequest()
	}
	response = NewBulkCreateApplicationRunningDataCollectionsResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationCollectStatusRequest() (request *GetApplicationCollectStatusRequest) {
	request = &GetApplicationCollectStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetApplicationCollectStatus")
	return
}

func NewGetApplicationCollectStatusResponse() (response *GetApplicationCollectStatusResponse) {
	response = &GetApplicationCollectStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetApplicationCollectStatus(request *GetApplicationCollectStatusRequest) (response *GetApplicationCollectStatusResponse, err error) {
	if request == nil {
		request = NewGetApplicationCollectStatusRequest()
	}
	response = NewGetApplicationCollectStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperationSheetAttachmentRequest() (request *CreateOperationSheetAttachmentRequest) {
	request = &CreateOperationSheetAttachmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateOperationSheetAttachment")
	return
}

func NewCreateOperationSheetAttachmentResponse() (response *CreateOperationSheetAttachmentResponse) {
	response = &CreateOperationSheetAttachmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateOperationSheetAttachment(request *CreateOperationSheetAttachmentRequest) (response *CreateOperationSheetAttachmentResponse, err error) {
	if request == nil {
		request = NewCreateOperationSheetAttachmentRequest()
	}
	response = NewCreateOperationSheetAttachmentResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaterialApplicationImportDetailRequest() (request *DeleteMaterialApplicationImportDetailRequest) {
	request = &DeleteMaterialApplicationImportDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteMaterialApplicationImportDetail")
	return
}

func NewDeleteMaterialApplicationImportDetailResponse() (response *DeleteMaterialApplicationImportDetailResponse) {
	response = &DeleteMaterialApplicationImportDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteMaterialApplicationImportDetail(request *DeleteMaterialApplicationImportDetailRequest) (response *DeleteMaterialApplicationImportDetailResponse, err error) {
	if request == nil {
		request = NewDeleteMaterialApplicationImportDetailRequest()
	}
	response = NewDeleteMaterialApplicationImportDetailResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateSolutionTemplatesRequest() (request *BulkCreateSolutionTemplatesRequest) {
	request = &BulkCreateSolutionTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateSolutionTemplates")
	return
}

func NewBulkCreateSolutionTemplatesResponse() (response *BulkCreateSolutionTemplatesResponse) {
	response = &BulkCreateSolutionTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateSolutionTemplates(request *BulkCreateSolutionTemplatesRequest) (response *BulkCreateSolutionTemplatesResponse, err error) {
	if request == nil {
		request = NewBulkCreateSolutionTemplatesRequest()
	}
	response = NewBulkCreateSolutionTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewListProductVersionArtifactTagsRequest() (request *ListProductVersionArtifactTagsRequest) {
	request = &ListProductVersionArtifactTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProductVersionArtifactTags")
	return
}

func NewListProductVersionArtifactTagsResponse() (response *ListProductVersionArtifactTagsResponse) {
	response = &ListProductVersionArtifactTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProductVersionArtifactTags(request *ListProductVersionArtifactTagsRequest) (response *ListProductVersionArtifactTagsResponse, err error) {
	if request == nil {
		request = NewListProductVersionArtifactTagsRequest()
	}
	response = NewListProductVersionArtifactTagsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePlanServerInfoRequest() (request *UpdatePlanServerInfoRequest) {
	request = &UpdatePlanServerInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdatePlanServerInfo")
	return
}

func NewUpdatePlanServerInfoResponse() (response *UpdatePlanServerInfoResponse) {
	response = &UpdatePlanServerInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdatePlanServerInfo(request *UpdatePlanServerInfoRequest) (response *UpdatePlanServerInfoResponse, err error) {
	if request == nil {
		request = NewUpdatePlanServerInfoRequest()
	}
	response = NewUpdatePlanServerInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGetProjectRegionRequest() (request *GetProjectRegionRequest) {
	request = &GetProjectRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProjectRegion")
	return
}

func NewGetProjectRegionResponse() (response *GetProjectRegionResponse) {
	response = &GetProjectRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProjectRegion(request *GetProjectRegionRequest) (response *GetProjectRegionResponse, err error) {
	if request == nil {
		request = NewGetProjectRegionRequest()
	}
	response = NewGetProjectRegionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperationSheetApprovalRecordRequest() (request *CreateOperationSheetApprovalRecordRequest) {
	request = &CreateOperationSheetApprovalRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateOperationSheetApprovalRecord")
	return
}

func NewCreateOperationSheetApprovalRecordResponse() (response *CreateOperationSheetApprovalRecordResponse) {
	response = &CreateOperationSheetApprovalRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateOperationSheetApprovalRecord(request *CreateOperationSheetApprovalRecordRequest) (response *CreateOperationSheetApprovalRecordResponse, err error) {
	if request == nil {
		request = NewCreateOperationSheetApprovalRecordRequest()
	}
	response = NewCreateOperationSheetApprovalRecordResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProductVersionApplicationPackageRelRequest() (request *UpdateProductVersionApplicationPackageRelRequest) {
	request = &UpdateProductVersionApplicationPackageRelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateProductVersionApplicationPackageRel")
	return
}

func NewUpdateProductVersionApplicationPackageRelResponse() (response *UpdateProductVersionApplicationPackageRelResponse) {
	response = &UpdateProductVersionApplicationPackageRelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateProductVersionApplicationPackageRel(request *UpdateProductVersionApplicationPackageRelRequest) (response *UpdateProductVersionApplicationPackageRelResponse, err error) {
	if request == nil {
		request = NewUpdateProductVersionApplicationPackageRelRequest()
	}
	response = NewUpdateProductVersionApplicationPackageRelResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOtaCertificateRequest() (request *DeleteOtaCertificateRequest) {
	request = &DeleteOtaCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteOtaCertificate")
	return
}

func NewDeleteOtaCertificateResponse() (response *DeleteOtaCertificateResponse) {
	response = &DeleteOtaCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteOtaCertificate(request *DeleteOtaCertificateRequest) (response *DeleteOtaCertificateResponse, err error) {
	if request == nil {
		request = NewDeleteOtaCertificateRequest()
	}
	response = NewDeleteOtaCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePlanTagRequest() (request *DeletePlanTagRequest) {
	request = &DeletePlanTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeletePlanTag")
	return
}

func NewDeletePlanTagResponse() (response *DeletePlanTagResponse) {
	response = &DeletePlanTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeletePlanTag(request *DeletePlanTagRequest) (response *DeletePlanTagResponse, err error) {
	if request == nil {
		request = NewDeletePlanTagRequest()
	}
	response = NewDeletePlanTagResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMaterialInfoRequest() (request *UpdateMaterialInfoRequest) {
	request = &UpdateMaterialInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateMaterialInfo")
	return
}

func NewUpdateMaterialInfoResponse() (response *UpdateMaterialInfoResponse) {
	response = &UpdateMaterialInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateMaterialInfo(request *UpdateMaterialInfoRequest) (response *UpdateMaterialInfoResponse, err error) {
	if request == nil {
		request = NewUpdateMaterialInfoRequest()
	}
	response = NewUpdateMaterialInfoResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateOperationBuiltinDataSchemasRequest() (request *BulkCreateOperationBuiltinDataSchemasRequest) {
	request = &BulkCreateOperationBuiltinDataSchemasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateOperationBuiltinDataSchemas")
	return
}

func NewBulkCreateOperationBuiltinDataSchemasResponse() (response *BulkCreateOperationBuiltinDataSchemasResponse) {
	response = &BulkCreateOperationBuiltinDataSchemasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateOperationBuiltinDataSchemas(request *BulkCreateOperationBuiltinDataSchemasRequest) (response *BulkCreateOperationBuiltinDataSchemasResponse, err error) {
	if request == nil {
		request = NewBulkCreateOperationBuiltinDataSchemasRequest()
	}
	response = NewBulkCreateOperationBuiltinDataSchemasResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationInstanceBackupRequest() (request *GetApplicationInstanceBackupRequest) {
	request = &GetApplicationInstanceBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetApplicationInstanceBackup")
	return
}

func NewGetApplicationInstanceBackupResponse() (response *GetApplicationInstanceBackupResponse) {
	response = &GetApplicationInstanceBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetApplicationInstanceBackup(request *GetApplicationInstanceBackupRequest) (response *GetApplicationInstanceBackupResponse, err error) {
	if request == nil {
		request = NewGetApplicationInstanceBackupRequest()
	}
	response = NewGetApplicationInstanceBackupResponse()
	err = c.Send(request, response)
	return
}

func NewGetFlywaySchemaHistoryRequest() (request *GetFlywaySchemaHistoryRequest) {
	request = &GetFlywaySchemaHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetFlywaySchemaHistory")
	return
}

func NewGetFlywaySchemaHistoryResponse() (response *GetFlywaySchemaHistoryResponse) {
	response = &GetFlywaySchemaHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetFlywaySchemaHistory(request *GetFlywaySchemaHistoryRequest) (response *GetFlywaySchemaHistoryResponse, err error) {
	if request == nil {
		request = NewGetFlywaySchemaHistoryRequest()
	}
	response = NewGetFlywaySchemaHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateProductDictionariesRequest() (request *BulkCreateProductDictionariesRequest) {
	request = &BulkCreateProductDictionariesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateProductDictionaries")
	return
}

func NewBulkCreateProductDictionariesResponse() (response *BulkCreateProductDictionariesResponse) {
	response = &BulkCreateProductDictionariesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateProductDictionaries(request *BulkCreateProductDictionariesRequest) (response *BulkCreateProductDictionariesResponse, err error) {
	if request == nil {
		request = NewBulkCreateProductDictionariesRequest()
	}
	response = NewBulkCreateProductDictionariesResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePlanApplicationRuntimeNameRequest() (request *CreatePlanApplicationRuntimeNameRequest) {
	request = &CreatePlanApplicationRuntimeNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreatePlanApplicationRuntimeName")
	return
}

func NewCreatePlanApplicationRuntimeNameResponse() (response *CreatePlanApplicationRuntimeNameResponse) {
	response = &CreatePlanApplicationRuntimeNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreatePlanApplicationRuntimeName(request *CreatePlanApplicationRuntimeNameRequest) (response *CreatePlanApplicationRuntimeNameResponse, err error) {
	if request == nil {
		request = NewCreatePlanApplicationRuntimeNameRequest()
	}
	response = NewCreatePlanApplicationRuntimeNameResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApplicationRunningDataCollectionRequest() (request *UpdateApplicationRunningDataCollectionRequest) {
	request = &UpdateApplicationRunningDataCollectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateApplicationRunningDataCollection")
	return
}

func NewUpdateApplicationRunningDataCollectionResponse() (response *UpdateApplicationRunningDataCollectionResponse) {
	response = &UpdateApplicationRunningDataCollectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateApplicationRunningDataCollection(request *UpdateApplicationRunningDataCollectionRequest) (response *UpdateApplicationRunningDataCollectionResponse, err error) {
	if request == nil {
		request = NewUpdateApplicationRunningDataCollectionRequest()
	}
	response = NewUpdateApplicationRunningDataCollectionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMaterialInfoRequest() (request *CreateMaterialInfoRequest) {
	request = &CreateMaterialInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateMaterialInfo")
	return
}

func NewCreateMaterialInfoResponse() (response *CreateMaterialInfoResponse) {
	response = &CreateMaterialInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateMaterialInfo(request *CreateMaterialInfoRequest) (response *CreateMaterialInfoResponse, err error) {
	if request == nil {
		request = NewCreateMaterialInfoRequest()
	}
	response = NewCreateMaterialInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductOperationJobRequest() (request *DeleteProductOperationJobRequest) {
	request = &DeleteProductOperationJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteProductOperationJob")
	return
}

func NewDeleteProductOperationJobResponse() (response *DeleteProductOperationJobResponse) {
	response = &DeleteProductOperationJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteProductOperationJob(request *DeleteProductOperationJobRequest) (response *DeleteProductOperationJobResponse, err error) {
	if request == nil {
		request = NewDeleteProductOperationJobRequest()
	}
	response = NewDeleteProductOperationJobResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationPackageRequest() (request *GetApplicationPackageRequest) {
	request = &GetApplicationPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetApplicationPackage")
	return
}

func NewGetApplicationPackageResponse() (response *GetApplicationPackageResponse) {
	response = &GetApplicationPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetApplicationPackage(request *GetApplicationPackageRequest) (response *GetApplicationPackageResponse, err error) {
	if request == nil {
		request = NewGetApplicationPackageRequest()
	}
	response = NewGetApplicationPackageResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePlanRequest() (request *CreatePlanRequest) {
	request = &CreatePlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreatePlan")
	return
}

func NewCreatePlanResponse() (response *CreatePlanResponse) {
	response = &CreatePlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreatePlan(request *CreatePlanRequest) (response *CreatePlanResponse, err error) {
	if request == nil {
		request = NewCreatePlanRequest()
	}
	response = NewCreatePlanResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProductTemplateRequest() (request *UpdateProductTemplateRequest) {
	request = &UpdateProductTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateProductTemplate")
	return
}

func NewUpdateProductTemplateResponse() (response *UpdateProductTemplateResponse) {
	response = &UpdateProductTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateProductTemplate(request *UpdateProductTemplateRequest) (response *UpdateProductTemplateResponse, err error) {
	if request == nil {
		request = NewUpdateProductTemplateRequest()
	}
	response = NewUpdateProductTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProjectTopologyRequest() (request *UpdateProjectTopologyRequest) {
	request = &UpdateProjectTopologyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateProjectTopology")
	return
}

func NewUpdateProjectTopologyResponse() (response *UpdateProjectTopologyResponse) {
	response = &UpdateProjectTopologyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateProjectTopology(request *UpdateProjectTopologyRequest) (response *UpdateProjectTopologyResponse, err error) {
	if request == nil {
		request = NewUpdateProjectTopologyRequest()
	}
	response = NewUpdateProjectTopologyResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOperationSheetAttachmentRequest() (request *UpdateOperationSheetAttachmentRequest) {
	request = &UpdateOperationSheetAttachmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateOperationSheetAttachment")
	return
}

func NewUpdateOperationSheetAttachmentResponse() (response *UpdateOperationSheetAttachmentResponse) {
	response = &UpdateOperationSheetAttachmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateOperationSheetAttachment(request *UpdateOperationSheetAttachmentRequest) (response *UpdateOperationSheetAttachmentResponse, err error) {
	if request == nil {
		request = NewUpdateOperationSheetAttachmentRequest()
	}
	response = NewUpdateOperationSheetAttachmentResponse()
	err = c.Send(request, response)
	return
}

func NewListOperationTemplatesRequest() (request *ListOperationTemplatesRequest) {
	request = &ListOperationTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListOperationTemplates")
	return
}

func NewListOperationTemplatesResponse() (response *ListOperationTemplatesResponse) {
	response = &ListOperationTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListOperationTemplates(request *ListOperationTemplatesRequest) (response *ListOperationTemplatesResponse, err error) {
	if request == nil {
		request = NewListOperationTemplatesRequest()
	}
	response = NewListOperationTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialMachineSSHInfosRequest() (request *ListMaterialMachineSSHInfosRequest) {
	request = &ListMaterialMachineSSHInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListMaterialMachineSSHInfos")
	return
}

func NewListMaterialMachineSSHInfosResponse() (response *ListMaterialMachineSSHInfosResponse) {
	response = &ListMaterialMachineSSHInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListMaterialMachineSSHInfos(request *ListMaterialMachineSSHInfosRequest) (response *ListMaterialMachineSSHInfosResponse, err error) {
	if request == nil {
		request = NewListMaterialMachineSSHInfosRequest()
	}
	response = NewListMaterialMachineSSHInfosResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePlanHistoryRequest() (request *CreatePlanHistoryRequest) {
	request = &CreatePlanHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreatePlanHistory")
	return
}

func NewCreatePlanHistoryResponse() (response *CreatePlanHistoryResponse) {
	response = &CreatePlanHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreatePlanHistory(request *CreatePlanHistoryRequest) (response *CreatePlanHistoryResponse, err error) {
	if request == nil {
		request = NewCreatePlanHistoryRequest()
	}
	response = NewCreatePlanHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialInfosRequest() (request *ListMaterialInfosRequest) {
	request = &ListMaterialInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListMaterialInfos")
	return
}

func NewListMaterialInfosResponse() (response *ListMaterialInfosResponse) {
	response = &ListMaterialInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListMaterialInfos(request *ListMaterialInfosRequest) (response *ListMaterialInfosResponse, err error) {
	if request == nil {
		request = NewListMaterialInfosRequest()
	}
	response = NewListMaterialInfosResponse()
	err = c.Send(request, response)
	return
}

func NewGetOrchTemplateLabelRequest() (request *GetOrchTemplateLabelRequest) {
	request = &GetOrchTemplateLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetOrchTemplateLabel")
	return
}

func NewGetOrchTemplateLabelResponse() (response *GetOrchTemplateLabelResponse) {
	response = &GetOrchTemplateLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetOrchTemplateLabel(request *GetOrchTemplateLabelRequest) (response *GetOrchTemplateLabelResponse, err error) {
	if request == nil {
		request = NewGetOrchTemplateLabelRequest()
	}
	response = NewGetOrchTemplateLabelResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateProductInstanceOperationSheetsRequest() (request *BulkCreateProductInstanceOperationSheetsRequest) {
	request = &BulkCreateProductInstanceOperationSheetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateProductInstanceOperationSheets")
	return
}

func NewBulkCreateProductInstanceOperationSheetsResponse() (response *BulkCreateProductInstanceOperationSheetsResponse) {
	response = &BulkCreateProductInstanceOperationSheetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateProductInstanceOperationSheets(request *BulkCreateProductInstanceOperationSheetsRequest) (response *BulkCreateProductInstanceOperationSheetsResponse, err error) {
	if request == nil {
		request = NewBulkCreateProductInstanceOperationSheetsRequest()
	}
	response = NewBulkCreateProductInstanceOperationSheetsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateSiteDeployTasksRequest() (request *BulkCreateSiteDeployTasksRequest) {
	request = &BulkCreateSiteDeployTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateSiteDeployTasks")
	return
}

func NewBulkCreateSiteDeployTasksResponse() (response *BulkCreateSiteDeployTasksResponse) {
	response = &BulkCreateSiteDeployTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateSiteDeployTasks(request *BulkCreateSiteDeployTasksRequest) (response *BulkCreateSiteDeployTasksResponse, err error) {
	if request == nil {
		request = NewBulkCreateSiteDeployTasksRequest()
	}
	response = NewBulkCreateSiteDeployTasksResponse()
	err = c.Send(request, response)
	return
}

func NewGetOtaCertificateRequest() (request *GetOtaCertificateRequest) {
	request = &GetOtaCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetOtaCertificate")
	return
}

func NewGetOtaCertificateResponse() (response *GetOtaCertificateResponse) {
	response = &GetOtaCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetOtaCertificate(request *GetOtaCertificateRequest) (response *GetOtaCertificateResponse, err error) {
	if request == nil {
		request = NewGetOtaCertificateRequest()
	}
	response = NewGetOtaCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewGetTenantRequest() (request *GetTenantRequest) {
	request = &GetTenantRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetTenant")
	return
}

func NewGetTenantResponse() (response *GetTenantResponse) {
	response = &GetTenantResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetTenant(request *GetTenantRequest) (response *GetTenantResponse, err error) {
	if request == nil {
		request = NewGetTenantRequest()
	}
	response = NewGetTenantResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOperationSheetRequest() (request *UpdateOperationSheetRequest) {
	request = &UpdateOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateOperationSheet")
	return
}

func NewUpdateOperationSheetResponse() (response *UpdateOperationSheetResponse) {
	response = &UpdateOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateOperationSheet(request *UpdateOperationSheetRequest) (response *UpdateOperationSheetResponse, err error) {
	if request == nil {
		request = NewUpdateOperationSheetRequest()
	}
	response = NewUpdateOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePlanSiteRequest() (request *UpdatePlanSiteRequest) {
	request = &UpdatePlanSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdatePlanSite")
	return
}

func NewUpdatePlanSiteResponse() (response *UpdatePlanSiteResponse) {
	response = &UpdatePlanSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdatePlanSite(request *UpdatePlanSiteRequest) (response *UpdatePlanSiteResponse, err error) {
	if request == nil {
		request = NewUpdatePlanSiteRequest()
	}
	response = NewUpdatePlanSiteResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProductInstanceOperationSheetRequest() (request *UpdateProductInstanceOperationSheetRequest) {
	request = &UpdateProductInstanceOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateProductInstanceOperationSheet")
	return
}

func NewUpdateProductInstanceOperationSheetResponse() (response *UpdateProductInstanceOperationSheetResponse) {
	response = &UpdateProductInstanceOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateProductInstanceOperationSheet(request *UpdateProductInstanceOperationSheetRequest) (response *UpdateProductInstanceOperationSheetResponse, err error) {
	if request == nil {
		request = NewUpdateProductInstanceOperationSheetRequest()
	}
	response = NewUpdateProductInstanceOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApplicationPackageRequest() (request *UpdateApplicationPackageRequest) {
	request = &UpdateApplicationPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateApplicationPackage")
	return
}

func NewUpdateApplicationPackageResponse() (response *UpdateApplicationPackageResponse) {
	response = &UpdateApplicationPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateApplicationPackage(request *UpdateApplicationPackageRequest) (response *UpdateApplicationPackageResponse, err error) {
	if request == nil {
		request = NewUpdateApplicationPackageRequest()
	}
	response = NewUpdateApplicationPackageResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOperationSheetTemplateRequest() (request *UpdateOperationSheetTemplateRequest) {
	request = &UpdateOperationSheetTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateOperationSheetTemplate")
	return
}

func NewUpdateOperationSheetTemplateResponse() (response *UpdateOperationSheetTemplateResponse) {
	response = &UpdateOperationSheetTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateOperationSheetTemplate(request *UpdateOperationSheetTemplateRequest) (response *UpdateOperationSheetTemplateResponse, err error) {
	if request == nil {
		request = NewUpdateOperationSheetTemplateRequest()
	}
	response = NewUpdateOperationSheetTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductDictionaryRequest() (request *DeleteProductDictionaryRequest) {
	request = &DeleteProductDictionaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteProductDictionary")
	return
}

func NewDeleteProductDictionaryResponse() (response *DeleteProductDictionaryResponse) {
	response = &DeleteProductDictionaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteProductDictionary(request *DeleteProductDictionaryRequest) (response *DeleteProductDictionaryResponse, err error) {
	if request == nil {
		request = NewDeleteProductDictionaryRequest()
	}
	response = NewDeleteProductDictionaryResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTenantDetailRequest() (request *CreateTenantDetailRequest) {
	request = &CreateTenantDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateTenantDetail")
	return
}

func NewCreateTenantDetailResponse() (response *CreateTenantDetailResponse) {
	response = &CreateTenantDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateTenantDetail(request *CreateTenantDetailRequest) (response *CreateTenantDetailResponse, err error) {
	if request == nil {
		request = NewCreateTenantDetailRequest()
	}
	response = NewCreateTenantDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApplicationArtifactRequest() (request *DeleteApplicationArtifactRequest) {
	request = &DeleteApplicationArtifactRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteApplicationArtifact")
	return
}

func NewDeleteApplicationArtifactResponse() (response *DeleteApplicationArtifactResponse) {
	response = &DeleteApplicationArtifactResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteApplicationArtifact(request *DeleteApplicationArtifactRequest) (response *DeleteApplicationArtifactResponse, err error) {
	if request == nil {
		request = NewDeleteApplicationArtifactRequest()
	}
	response = NewDeleteApplicationArtifactResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePlanApplicationServerRelationRequest() (request *DeletePlanApplicationServerRelationRequest) {
	request = &DeletePlanApplicationServerRelationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeletePlanApplicationServerRelation")
	return
}

func NewDeletePlanApplicationServerRelationResponse() (response *DeletePlanApplicationServerRelationResponse) {
	response = &DeletePlanApplicationServerRelationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeletePlanApplicationServerRelation(request *DeletePlanApplicationServerRelationRequest) (response *DeletePlanApplicationServerRelationResponse, err error) {
	if request == nil {
		request = NewDeletePlanApplicationServerRelationRequest()
	}
	response = NewDeletePlanApplicationServerRelationResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateMaterialCustomizedTypesRequest() (request *BulkCreateMaterialCustomizedTypesRequest) {
	request = &BulkCreateMaterialCustomizedTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateMaterialCustomizedTypes")
	return
}

func NewBulkCreateMaterialCustomizedTypesResponse() (response *BulkCreateMaterialCustomizedTypesResponse) {
	response = &BulkCreateMaterialCustomizedTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateMaterialCustomizedTypes(request *BulkCreateMaterialCustomizedTypesRequest) (response *BulkCreateMaterialCustomizedTypesResponse, err error) {
	if request == nil {
		request = NewBulkCreateMaterialCustomizedTypesRequest()
	}
	response = NewBulkCreateMaterialCustomizedTypesResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductInstanceBackupRequest() (request *GetProductInstanceBackupRequest) {
	request = &GetProductInstanceBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProductInstanceBackup")
	return
}

func NewGetProductInstanceBackupResponse() (response *GetProductInstanceBackupResponse) {
	response = &GetProductInstanceBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProductInstanceBackup(request *GetProductInstanceBackupRequest) (response *GetProductInstanceBackupResponse, err error) {
	if request == nil {
		request = NewGetProductInstanceBackupRequest()
	}
	response = NewGetProductInstanceBackupResponse()
	err = c.Send(request, response)
	return
}

func NewListOperationSheetApprovalRecordsRequest() (request *ListOperationSheetApprovalRecordsRequest) {
	request = &ListOperationSheetApprovalRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListOperationSheetApprovalRecords")
	return
}

func NewListOperationSheetApprovalRecordsResponse() (response *ListOperationSheetApprovalRecordsResponse) {
	response = &ListOperationSheetApprovalRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListOperationSheetApprovalRecords(request *ListOperationSheetApprovalRecordsRequest) (response *ListOperationSheetApprovalRecordsResponse, err error) {
	if request == nil {
		request = NewListOperationSheetApprovalRecordsRequest()
	}
	response = NewListOperationSheetApprovalRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateSystemSettingsRequest() (request *BulkCreateSystemSettingsRequest) {
	request = &BulkCreateSystemSettingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateSystemSettings")
	return
}

func NewBulkCreateSystemSettingsResponse() (response *BulkCreateSystemSettingsResponse) {
	response = &BulkCreateSystemSettingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateSystemSettings(request *BulkCreateSystemSettingsRequest) (response *BulkCreateSystemSettingsResponse, err error) {
	if request == nil {
		request = NewBulkCreateSystemSettingsRequest()
	}
	response = NewBulkCreateSystemSettingsResponse()
	err = c.Send(request, response)
	return
}

func NewListCommonOperationSheetsRequest() (request *ListCommonOperationSheetsRequest) {
	request = &ListCommonOperationSheetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListCommonOperationSheets")
	return
}

func NewListCommonOperationSheetsResponse() (response *ListCommonOperationSheetsResponse) {
	response = &ListCommonOperationSheetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListCommonOperationSheets(request *ListCommonOperationSheetsRequest) (response *ListCommonOperationSheetsResponse, err error) {
	if request == nil {
		request = NewListCommonOperationSheetsRequest()
	}
	response = NewListCommonOperationSheetsResponse()
	err = c.Send(request, response)
	return
}

func NewListProductVersionsRequest() (request *ListProductVersionsRequest) {
	request = &ListProductVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProductVersions")
	return
}

func NewListProductVersionsResponse() (response *ListProductVersionsResponse) {
	response = &ListProductVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProductVersions(request *ListProductVersionsRequest) (response *ListProductVersionsResponse, err error) {
	if request == nil {
		request = NewListProductVersionsRequest()
	}
	response = NewListProductVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewListPlanSitesRequest() (request *ListPlanSitesRequest) {
	request = &ListPlanSitesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListPlanSites")
	return
}

func NewListPlanSitesResponse() (response *ListPlanSitesResponse) {
	response = &ListPlanSitesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListPlanSites(request *ListPlanSitesRequest) (response *ListPlanSitesResponse, err error) {
	if request == nil {
		request = NewListPlanSitesRequest()
	}
	response = NewListPlanSitesResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateMaterialMachineSSHInfosRequest() (request *BulkCreateMaterialMachineSSHInfosRequest) {
	request = &BulkCreateMaterialMachineSSHInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateMaterialMachineSSHInfos")
	return
}

func NewBulkCreateMaterialMachineSSHInfosResponse() (response *BulkCreateMaterialMachineSSHInfosResponse) {
	response = &BulkCreateMaterialMachineSSHInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateMaterialMachineSSHInfos(request *BulkCreateMaterialMachineSSHInfosRequest) (response *BulkCreateMaterialMachineSSHInfosResponse, err error) {
	if request == nil {
		request = NewBulkCreateMaterialMachineSSHInfosRequest()
	}
	response = NewBulkCreateMaterialMachineSSHInfosResponse()
	err = c.Send(request, response)
	return
}

func NewGetProjectSiteInfoRequest() (request *GetProjectSiteInfoRequest) {
	request = &GetProjectSiteInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProjectSiteInfo")
	return
}

func NewGetProjectSiteInfoResponse() (response *GetProjectSiteInfoResponse) {
	response = &GetProjectSiteInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProjectSiteInfo(request *GetProjectSiteInfoRequest) (response *GetProjectSiteInfoResponse, err error) {
	if request == nil {
		request = NewGetProjectSiteInfoRequest()
	}
	response = NewGetProjectSiteInfoResponse()
	err = c.Send(request, response)
	return
}

func NewListTenantsRequest() (request *ListTenantsRequest) {
	request = &ListTenantsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListTenants")
	return
}

func NewListTenantsResponse() (response *ListTenantsResponse) {
	response = &ListTenantsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListTenants(request *ListTenantsRequest) (response *ListTenantsResponse, err error) {
	if request == nil {
		request = NewListTenantsRequest()
	}
	response = NewListTenantsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDagNodeRequest() (request *DeleteDagNodeRequest) {
	request = &DeleteDagNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteDagNode")
	return
}

func NewDeleteDagNodeResponse() (response *DeleteDagNodeResponse) {
	response = &DeleteDagNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteDagNode(request *DeleteDagNodeRequest) (response *DeleteDagNodeResponse, err error) {
	if request == nil {
		request = NewDeleteDagNodeRequest()
	}
	response = NewDeleteDagNodeResponse()
	err = c.Send(request, response)
	return
}

func NewListOperationSheetAttachmentsRequest() (request *ListOperationSheetAttachmentsRequest) {
	request = &ListOperationSheetAttachmentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListOperationSheetAttachments")
	return
}

func NewListOperationSheetAttachmentsResponse() (response *ListOperationSheetAttachmentsResponse) {
	response = &ListOperationSheetAttachmentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListOperationSheetAttachments(request *ListOperationSheetAttachmentsRequest) (response *ListOperationSheetAttachmentsResponse, err error) {
	if request == nil {
		request = NewListOperationSheetAttachmentsRequest()
	}
	response = NewListOperationSheetAttachmentsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateApplicationInstanceDeployVersionsRequest() (request *BulkCreateApplicationInstanceDeployVersionsRequest) {
	request = &BulkCreateApplicationInstanceDeployVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateApplicationInstanceDeployVersions")
	return
}

func NewBulkCreateApplicationInstanceDeployVersionsResponse() (response *BulkCreateApplicationInstanceDeployVersionsResponse) {
	response = &BulkCreateApplicationInstanceDeployVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateApplicationInstanceDeployVersions(request *BulkCreateApplicationInstanceDeployVersionsRequest) (response *BulkCreateApplicationInstanceDeployVersionsResponse, err error) {
	if request == nil {
		request = NewBulkCreateApplicationInstanceDeployVersionsRequest()
	}
	response = NewBulkCreateApplicationInstanceDeployVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductVersionRequest() (request *DeleteProductVersionRequest) {
	request = &DeleteProductVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteProductVersion")
	return
}

func NewDeleteProductVersionResponse() (response *DeleteProductVersionResponse) {
	response = &DeleteProductVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteProductVersion(request *DeleteProductVersionRequest) (response *DeleteProductVersionResponse, err error) {
	if request == nil {
		request = NewDeleteProductVersionRequest()
	}
	response = NewDeleteProductVersionResponse()
	err = c.Send(request, response)
	return
}

func NewGetMaterialSiteRequest() (request *GetMaterialSiteRequest) {
	request = &GetMaterialSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetMaterialSite")
	return
}

func NewGetMaterialSiteResponse() (response *GetMaterialSiteResponse) {
	response = &GetMaterialSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetMaterialSite(request *GetMaterialSiteRequest) (response *GetMaterialSiteResponse, err error) {
	if request == nil {
		request = NewGetMaterialSiteRequest()
	}
	response = NewGetMaterialSiteResponse()
	err = c.Send(request, response)
	return
}

func NewGetSolutionRequest() (request *GetSolutionRequest) {
	request = &GetSolutionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetSolution")
	return
}

func NewGetSolutionResponse() (response *GetSolutionResponse) {
	response = &GetSolutionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetSolution(request *GetSolutionRequest) (response *GetSolutionResponse, err error) {
	if request == nil {
		request = NewGetSolutionRequest()
	}
	response = NewGetSolutionResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationArtifactRequest() (request *GetApplicationArtifactRequest) {
	request = &GetApplicationArtifactRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetApplicationArtifact")
	return
}

func NewGetApplicationArtifactResponse() (response *GetApplicationArtifactResponse) {
	response = &GetApplicationArtifactResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetApplicationArtifact(request *GetApplicationArtifactRequest) (response *GetApplicationArtifactResponse, err error) {
	if request == nil {
		request = NewGetApplicationArtifactRequest()
	}
	response = NewGetApplicationArtifactResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOtaCertificateRequest() (request *UpdateOtaCertificateRequest) {
	request = &UpdateOtaCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateOtaCertificate")
	return
}

func NewUpdateOtaCertificateResponse() (response *UpdateOtaCertificateResponse) {
	response = &UpdateOtaCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateOtaCertificate(request *UpdateOtaCertificateRequest) (response *UpdateOtaCertificateResponse, err error) {
	if request == nil {
		request = NewUpdateOtaCertificateRequest()
	}
	response = NewUpdateOtaCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperationTemplateRequest() (request *CreateOperationTemplateRequest) {
	request = &CreateOperationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateOperationTemplate")
	return
}

func NewCreateOperationTemplateResponse() (response *CreateOperationTemplateResponse) {
	response = &CreateOperationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateOperationTemplate(request *CreateOperationTemplateRequest) (response *CreateOperationTemplateResponse, err error) {
	if request == nil {
		request = NewCreateOperationTemplateRequest()
	}
	response = NewCreateOperationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApplicationRequest() (request *UpdateApplicationRequest) {
	request = &UpdateApplicationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateApplication")
	return
}

func NewUpdateApplicationResponse() (response *UpdateApplicationResponse) {
	response = &UpdateApplicationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateApplication(request *UpdateApplicationRequest) (response *UpdateApplicationResponse, err error) {
	if request == nil {
		request = NewUpdateApplicationRequest()
	}
	response = NewUpdateApplicationResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProductVersionArtifactTagRequest() (request *UpdateProductVersionArtifactTagRequest) {
	request = &UpdateProductVersionArtifactTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateProductVersionArtifactTag")
	return
}

func NewUpdateProductVersionArtifactTagResponse() (response *UpdateProductVersionArtifactTagResponse) {
	response = &UpdateProductVersionArtifactTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateProductVersionArtifactTag(request *UpdateProductVersionArtifactTagRequest) (response *UpdateProductVersionArtifactTagResponse, err error) {
	if request == nil {
		request = NewUpdateProductVersionArtifactTagRequest()
	}
	response = NewUpdateProductVersionArtifactTagResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateApplicationCollectStatusesRequest() (request *BulkCreateApplicationCollectStatusesRequest) {
	request = &BulkCreateApplicationCollectStatusesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateApplicationCollectStatuses")
	return
}

func NewBulkCreateApplicationCollectStatusesResponse() (response *BulkCreateApplicationCollectStatusesResponse) {
	response = &BulkCreateApplicationCollectStatusesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateApplicationCollectStatuses(request *BulkCreateApplicationCollectStatusesRequest) (response *BulkCreateApplicationCollectStatusesResponse, err error) {
	if request == nil {
		request = NewBulkCreateApplicationCollectStatusesRequest()
	}
	response = NewBulkCreateApplicationCollectStatusesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductInstanceBackupRequest() (request *CreateProductInstanceBackupRequest) {
	request = &CreateProductInstanceBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateProductInstanceBackup")
	return
}

func NewCreateProductInstanceBackupResponse() (response *CreateProductInstanceBackupResponse) {
	response = &CreateProductInstanceBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateProductInstanceBackup(request *CreateProductInstanceBackupRequest) (response *CreateProductInstanceBackupResponse, err error) {
	if request == nil {
		request = NewCreateProductInstanceBackupRequest()
	}
	response = NewCreateProductInstanceBackupResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePlanApplicationInstanceRequest() (request *DeletePlanApplicationInstanceRequest) {
	request = &DeletePlanApplicationInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeletePlanApplicationInstance")
	return
}

func NewDeletePlanApplicationInstanceResponse() (response *DeletePlanApplicationInstanceResponse) {
	response = &DeletePlanApplicationInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeletePlanApplicationInstance(request *DeletePlanApplicationInstanceRequest) (response *DeletePlanApplicationInstanceResponse, err error) {
	if request == nil {
		request = NewDeletePlanApplicationInstanceRequest()
	}
	response = NewDeletePlanApplicationInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateApplicationsRequest() (request *BulkCreateApplicationsRequest) {
	request = &BulkCreateApplicationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateApplications")
	return
}

func NewBulkCreateApplicationsResponse() (response *BulkCreateApplicationsResponse) {
	response = &BulkCreateApplicationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateApplications(request *BulkCreateApplicationsRequest) (response *BulkCreateApplicationsResponse, err error) {
	if request == nil {
		request = NewBulkCreateApplicationsRequest()
	}
	response = NewBulkCreateApplicationsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMaterialSiteRequest() (request *CreateMaterialSiteRequest) {
	request = &CreateMaterialSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateMaterialSite")
	return
}

func NewCreateMaterialSiteResponse() (response *CreateMaterialSiteResponse) {
	response = &CreateMaterialSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateMaterialSite(request *CreateMaterialSiteRequest) (response *CreateMaterialSiteResponse, err error) {
	if request == nil {
		request = NewCreateMaterialSiteRequest()
	}
	response = NewCreateMaterialSiteResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApplicationInstanceDeployVersionRequest() (request *CreateApplicationInstanceDeployVersionRequest) {
	request = &CreateApplicationInstanceDeployVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateApplicationInstanceDeployVersion")
	return
}

func NewCreateApplicationInstanceDeployVersionResponse() (response *CreateApplicationInstanceDeployVersionResponse) {
	response = &CreateApplicationInstanceDeployVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateApplicationInstanceDeployVersion(request *CreateApplicationInstanceDeployVersionRequest) (response *CreateApplicationInstanceDeployVersionResponse, err error) {
	if request == nil {
		request = NewCreateApplicationInstanceDeployVersionRequest()
	}
	response = NewCreateApplicationInstanceDeployVersionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProjectZoneRequest() (request *CreateProjectZoneRequest) {
	request = &CreateProjectZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateProjectZone")
	return
}

func NewCreateProjectZoneResponse() (response *CreateProjectZoneResponse) {
	response = &CreateProjectZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateProjectZone(request *CreateProjectZoneRequest) (response *CreateProjectZoneResponse, err error) {
	if request == nil {
		request = NewCreateProjectZoneRequest()
	}
	response = NewCreateProjectZoneResponse()
	err = c.Send(request, response)
	return
}

func NewGetTenantDetailRequest() (request *GetTenantDetailRequest) {
	request = &GetTenantDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetTenantDetail")
	return
}

func NewGetTenantDetailResponse() (response *GetTenantDetailResponse) {
	response = &GetTenantDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetTenantDetail(request *GetTenantDetailRequest) (response *GetTenantDetailResponse, err error) {
	if request == nil {
		request = NewGetTenantDetailRequest()
	}
	response = NewGetTenantDetailResponse()
	err = c.Send(request, response)
	return
}

func NewListOperationBuiltinDataSchemasRequest() (request *ListOperationBuiltinDataSchemasRequest) {
	request = &ListOperationBuiltinDataSchemasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListOperationBuiltinDataSchemas")
	return
}

func NewListOperationBuiltinDataSchemasResponse() (response *ListOperationBuiltinDataSchemasResponse) {
	response = &ListOperationBuiltinDataSchemasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListOperationBuiltinDataSchemas(request *ListOperationBuiltinDataSchemasRequest) (response *ListOperationBuiltinDataSchemasResponse, err error) {
	if request == nil {
		request = NewListOperationBuiltinDataSchemasRequest()
	}
	response = NewListOperationBuiltinDataSchemasResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreatePlanSubsystemInstancesRequest() (request *BulkCreatePlanSubsystemInstancesRequest) {
	request = &BulkCreatePlanSubsystemInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreatePlanSubsystemInstances")
	return
}

func NewBulkCreatePlanSubsystemInstancesResponse() (response *BulkCreatePlanSubsystemInstancesResponse) {
	response = &BulkCreatePlanSubsystemInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreatePlanSubsystemInstances(request *BulkCreatePlanSubsystemInstancesRequest) (response *BulkCreatePlanSubsystemInstancesResponse, err error) {
	if request == nil {
		request = NewBulkCreatePlanSubsystemInstancesRequest()
	}
	response = NewBulkCreatePlanSubsystemInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePlanHistoryRequest() (request *UpdatePlanHistoryRequest) {
	request = &UpdatePlanHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdatePlanHistory")
	return
}

func NewUpdatePlanHistoryResponse() (response *UpdatePlanHistoryResponse) {
	response = &UpdatePlanHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdatePlanHistory(request *UpdatePlanHistoryRequest) (response *UpdatePlanHistoryResponse, err error) {
	if request == nil {
		request = NewUpdatePlanHistoryRequest()
	}
	response = NewUpdatePlanHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewListOrchTemplateLabelsRequest() (request *ListOrchTemplateLabelsRequest) {
	request = &ListOrchTemplateLabelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListOrchTemplateLabels")
	return
}

func NewListOrchTemplateLabelsResponse() (response *ListOrchTemplateLabelsResponse) {
	response = &ListOrchTemplateLabelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListOrchTemplateLabels(request *ListOrchTemplateLabelsRequest) (response *ListOrchTemplateLabelsResponse, err error) {
	if request == nil {
		request = NewListOrchTemplateLabelsRequest()
	}
	response = NewListOrchTemplateLabelsResponse()
	err = c.Send(request, response)
	return
}

func NewListOperationSheetLabelsRequest() (request *ListOperationSheetLabelsRequest) {
	request = &ListOperationSheetLabelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListOperationSheetLabels")
	return
}

func NewListOperationSheetLabelsResponse() (response *ListOperationSheetLabelsResponse) {
	response = &ListOperationSheetLabelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListOperationSheetLabels(request *ListOperationSheetLabelsRequest) (response *ListOperationSheetLabelsResponse, err error) {
	if request == nil {
		request = NewListOperationSheetLabelsRequest()
	}
	response = NewListOperationSheetLabelsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreatePlanSitesRequest() (request *BulkCreatePlanSitesRequest) {
	request = &BulkCreatePlanSitesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreatePlanSites")
	return
}

func NewBulkCreatePlanSitesResponse() (response *BulkCreatePlanSitesResponse) {
	response = &BulkCreatePlanSitesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreatePlanSites(request *BulkCreatePlanSitesRequest) (response *BulkCreatePlanSitesResponse, err error) {
	if request == nil {
		request = NewBulkCreatePlanSitesRequest()
	}
	response = NewBulkCreatePlanSitesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSolutionVersionRequest() (request *CreateSolutionVersionRequest) {
	request = &CreateSolutionVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateSolutionVersion")
	return
}

func NewCreateSolutionVersionResponse() (response *CreateSolutionVersionResponse) {
	response = &CreateSolutionVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateSolutionVersion(request *CreateSolutionVersionRequest) (response *CreateSolutionVersionResponse, err error) {
	if request == nil {
		request = NewCreateSolutionVersionRequest()
	}
	response = NewCreateSolutionVersionResponse()
	err = c.Send(request, response)
	return
}

func NewListApplicationInstanceBackupsRequest() (request *ListApplicationInstanceBackupsRequest) {
	request = &ListApplicationInstanceBackupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListApplicationInstanceBackups")
	return
}

func NewListApplicationInstanceBackupsResponse() (response *ListApplicationInstanceBackupsResponse) {
	response = &ListApplicationInstanceBackupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListApplicationInstanceBackups(request *ListApplicationInstanceBackupsRequest) (response *ListApplicationInstanceBackupsResponse, err error) {
	if request == nil {
		request = NewListApplicationInstanceBackupsRequest()
	}
	response = NewListApplicationInstanceBackupsResponse()
	err = c.Send(request, response)
	return
}

func NewListSiteProductVersionRelsRequest() (request *ListSiteProductVersionRelsRequest) {
	request = &ListSiteProductVersionRelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListSiteProductVersionRels")
	return
}

func NewListSiteProductVersionRelsResponse() (response *ListSiteProductVersionRelsResponse) {
	response = &ListSiteProductVersionRelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListSiteProductVersionRels(request *ListSiteProductVersionRelsRequest) (response *ListSiteProductVersionRelsResponse, err error) {
	if request == nil {
		request = NewListSiteProductVersionRelsRequest()
	}
	response = NewListSiteProductVersionRelsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductTemplateRequest() (request *CreateProductTemplateRequest) {
	request = &CreateProductTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateProductTemplate")
	return
}

func NewCreateProductTemplateResponse() (response *CreateProductTemplateResponse) {
	response = &CreateProductTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateProductTemplate(request *CreateProductTemplateRequest) (response *CreateProductTemplateResponse, err error) {
	if request == nil {
		request = NewCreateProductTemplateRequest()
	}
	response = NewCreateProductTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApplicationPackageRequest() (request *DeleteApplicationPackageRequest) {
	request = &DeleteApplicationPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteApplicationPackage")
	return
}

func NewDeleteApplicationPackageResponse() (response *DeleteApplicationPackageResponse) {
	response = &DeleteApplicationPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteApplicationPackage(request *DeleteApplicationPackageRequest) (response *DeleteApplicationPackageResponse, err error) {
	if request == nil {
		request = NewDeleteApplicationPackageRequest()
	}
	response = NewDeleteApplicationPackageResponse()
	err = c.Send(request, response)
	return
}

func NewGetMaterialMachineSSHInfoRequest() (request *GetMaterialMachineSSHInfoRequest) {
	request = &GetMaterialMachineSSHInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetMaterialMachineSSHInfo")
	return
}

func NewGetMaterialMachineSSHInfoResponse() (response *GetMaterialMachineSSHInfoResponse) {
	response = &GetMaterialMachineSSHInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetMaterialMachineSSHInfo(request *GetMaterialMachineSSHInfoRequest) (response *GetMaterialMachineSSHInfoResponse, err error) {
	if request == nil {
		request = NewGetMaterialMachineSSHInfoRequest()
	}
	response = NewGetMaterialMachineSSHInfoResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateProductTemplatesRequest() (request *BulkCreateProductTemplatesRequest) {
	request = &BulkCreateProductTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateProductTemplates")
	return
}

func NewBulkCreateProductTemplatesResponse() (response *BulkCreateProductTemplatesResponse) {
	response = &BulkCreateProductTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateProductTemplates(request *BulkCreateProductTemplatesRequest) (response *BulkCreateProductTemplatesResponse, err error) {
	if request == nil {
		request = NewBulkCreateProductTemplatesRequest()
	}
	response = NewBulkCreateProductTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDagNodeRequest() (request *CreateDagNodeRequest) {
	request = &CreateDagNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateDagNode")
	return
}

func NewCreateDagNodeResponse() (response *CreateDagNodeResponse) {
	response = &CreateDagNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateDagNode(request *CreateDagNodeRequest) (response *CreateDagNodeResponse, err error) {
	if request == nil {
		request = NewCreateDagNodeRequest()
	}
	response = NewCreateDagNodeResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProductOperationRequest() (request *UpdateProductOperationRequest) {
	request = &UpdateProductOperationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateProductOperation")
	return
}

func NewUpdateProductOperationResponse() (response *UpdateProductOperationResponse) {
	response = &UpdateProductOperationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateProductOperation(request *UpdateProductOperationRequest) (response *UpdateProductOperationResponse, err error) {
	if request == nil {
		request = NewUpdateProductOperationRequest()
	}
	response = NewUpdateProductOperationResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSolutionRequest() (request *DeleteSolutionRequest) {
	request = &DeleteSolutionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteSolution")
	return
}

func NewDeleteSolutionResponse() (response *DeleteSolutionResponse) {
	response = &DeleteSolutionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteSolution(request *DeleteSolutionRequest) (response *DeleteSolutionResponse, err error) {
	if request == nil {
		request = NewDeleteSolutionRequest()
	}
	response = NewDeleteSolutionResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOrchTemplateLabelRequest() (request *UpdateOrchTemplateLabelRequest) {
	request = &UpdateOrchTemplateLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateOrchTemplateLabel")
	return
}

func NewUpdateOrchTemplateLabelResponse() (response *UpdateOrchTemplateLabelResponse) {
	response = &UpdateOrchTemplateLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateOrchTemplateLabel(request *UpdateOrchTemplateLabelRequest) (response *UpdateOrchTemplateLabelResponse, err error) {
	if request == nil {
		request = NewUpdateOrchTemplateLabelRequest()
	}
	response = NewUpdateOrchTemplateLabelResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationRunningDataCollectionRequest() (request *GetApplicationRunningDataCollectionRequest) {
	request = &GetApplicationRunningDataCollectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetApplicationRunningDataCollection")
	return
}

func NewGetApplicationRunningDataCollectionResponse() (response *GetApplicationRunningDataCollectionResponse) {
	response = &GetApplicationRunningDataCollectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetApplicationRunningDataCollection(request *GetApplicationRunningDataCollectionRequest) (response *GetApplicationRunningDataCollectionResponse, err error) {
	if request == nil {
		request = NewGetApplicationRunningDataCollectionRequest()
	}
	response = NewGetApplicationRunningDataCollectionResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductVersionArtifactTagRequest() (request *GetProductVersionArtifactTagRequest) {
	request = &GetProductVersionArtifactTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProductVersionArtifactTag")
	return
}

func NewGetProductVersionArtifactTagResponse() (response *GetProductVersionArtifactTagResponse) {
	response = &GetProductVersionArtifactTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProductVersionArtifactTag(request *GetProductVersionArtifactTagRequest) (response *GetProductVersionArtifactTagResponse, err error) {
	if request == nil {
		request = NewGetProductVersionArtifactTagRequest()
	}
	response = NewGetProductVersionArtifactTagResponse()
	err = c.Send(request, response)
	return
}

func NewListPlanApplicationServerRelationsRequest() (request *ListPlanApplicationServerRelationsRequest) {
	request = &ListPlanApplicationServerRelationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListPlanApplicationServerRelations")
	return
}

func NewListPlanApplicationServerRelationsResponse() (response *ListPlanApplicationServerRelationsResponse) {
	response = &ListPlanApplicationServerRelationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListPlanApplicationServerRelations(request *ListPlanApplicationServerRelationsRequest) (response *ListPlanApplicationServerRelationsResponse, err error) {
	if request == nil {
		request = NewListPlanApplicationServerRelationsRequest()
	}
	response = NewListPlanApplicationServerRelationsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProductOperationJobRequest() (request *UpdateProductOperationJobRequest) {
	request = &UpdateProductOperationJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateProductOperationJob")
	return
}

func NewUpdateProductOperationJobResponse() (response *UpdateProductOperationJobResponse) {
	response = &UpdateProductOperationJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateProductOperationJob(request *UpdateProductOperationJobRequest) (response *UpdateProductOperationJobResponse, err error) {
	if request == nil {
		request = NewUpdateProductOperationJobRequest()
	}
	response = NewUpdateProductOperationJobResponse()
	err = c.Send(request, response)
	return
}

func NewListProductInstanceOperationSheetsRequest() (request *ListProductInstanceOperationSheetsRequest) {
	request = &ListProductInstanceOperationSheetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProductInstanceOperationSheets")
	return
}

func NewListProductInstanceOperationSheetsResponse() (response *ListProductInstanceOperationSheetsResponse) {
	response = &ListProductInstanceOperationSheetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProductInstanceOperationSheets(request *ListProductInstanceOperationSheetsRequest) (response *ListProductInstanceOperationSheetsResponse, err error) {
	if request == nil {
		request = NewListProductInstanceOperationSheetsRequest()
	}
	response = NewListProductInstanceOperationSheetsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProjectRegionRequest() (request *CreateProjectRegionRequest) {
	request = &CreateProjectRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateProjectRegion")
	return
}

func NewCreateProjectRegionResponse() (response *CreateProjectRegionResponse) {
	response = &CreateProjectRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateProjectRegion(request *CreateProjectRegionRequest) (response *CreateProjectRegionResponse, err error) {
	if request == nil {
		request = NewCreateProjectRegionRequest()
	}
	response = NewCreateProjectRegionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSiteProductVersionRelRequest() (request *CreateSiteProductVersionRelRequest) {
	request = &CreateSiteProductVersionRelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateSiteProductVersionRel")
	return
}

func NewCreateSiteProductVersionRelResponse() (response *CreateSiteProductVersionRelResponse) {
	response = &CreateSiteProductVersionRelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateSiteProductVersionRel(request *CreateSiteProductVersionRelRequest) (response *CreateSiteProductVersionRelResponse, err error) {
	if request == nil {
		request = NewCreateSiteProductVersionRelRequest()
	}
	response = NewCreateSiteProductVersionRelResponse()
	err = c.Send(request, response)
	return
}

func NewGetSiteOperationSheetRequest() (request *GetSiteOperationSheetRequest) {
	request = &GetSiteOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetSiteOperationSheet")
	return
}

func NewGetSiteOperationSheetResponse() (response *GetSiteOperationSheetResponse) {
	response = &GetSiteOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetSiteOperationSheet(request *GetSiteOperationSheetRequest) (response *GetSiteOperationSheetResponse, err error) {
	if request == nil {
		request = NewGetSiteOperationSheetRequest()
	}
	response = NewGetSiteOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductInstanceOperationSheetRequest() (request *DeleteProductInstanceOperationSheetRequest) {
	request = &DeleteProductInstanceOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteProductInstanceOperationSheet")
	return
}

func NewDeleteProductInstanceOperationSheetResponse() (response *DeleteProductInstanceOperationSheetResponse) {
	response = &DeleteProductInstanceOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteProductInstanceOperationSheet(request *DeleteProductInstanceOperationSheetRequest) (response *DeleteProductInstanceOperationSheetResponse, err error) {
	if request == nil {
		request = NewDeleteProductInstanceOperationSheetRequest()
	}
	response = NewDeleteProductInstanceOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMaterialCommonImportDetailRequest() (request *UpdateMaterialCommonImportDetailRequest) {
	request = &UpdateMaterialCommonImportDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateMaterialCommonImportDetail")
	return
}

func NewUpdateMaterialCommonImportDetailResponse() (response *UpdateMaterialCommonImportDetailResponse) {
	response = &UpdateMaterialCommonImportDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateMaterialCommonImportDetail(request *UpdateMaterialCommonImportDetailRequest) (response *UpdateMaterialCommonImportDetailResponse, err error) {
	if request == nil {
		request = NewUpdateMaterialCommonImportDetailRequest()
	}
	response = NewUpdateMaterialCommonImportDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApplicationCollectStatusRequest() (request *CreateApplicationCollectStatusRequest) {
	request = &CreateApplicationCollectStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateApplicationCollectStatus")
	return
}

func NewCreateApplicationCollectStatusResponse() (response *CreateApplicationCollectStatusResponse) {
	response = &CreateApplicationCollectStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateApplicationCollectStatus(request *CreateApplicationCollectStatusRequest) (response *CreateApplicationCollectStatusResponse, err error) {
	if request == nil {
		request = NewCreateApplicationCollectStatusRequest()
	}
	response = NewCreateApplicationCollectStatusResponse()
	err = c.Send(request, response)
	return
}

func NewGetSystemSettingRequest() (request *GetSystemSettingRequest) {
	request = &GetSystemSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetSystemSetting")
	return
}

func NewGetSystemSettingResponse() (response *GetSystemSettingResponse) {
	response = &GetSystemSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetSystemSetting(request *GetSystemSettingRequest) (response *GetSystemSettingResponse, err error) {
	if request == nil {
		request = NewGetSystemSettingRequest()
	}
	response = NewGetSystemSettingResponse()
	err = c.Send(request, response)
	return
}

func NewListApplicationBranchesRequest() (request *ListApplicationBranchesRequest) {
	request = &ListApplicationBranchesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListApplicationBranches")
	return
}

func NewListApplicationBranchesResponse() (response *ListApplicationBranchesResponse) {
	response = &ListApplicationBranchesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListApplicationBranches(request *ListApplicationBranchesRequest) (response *ListApplicationBranchesResponse, err error) {
	if request == nil {
		request = NewListApplicationBranchesRequest()
	}
	response = NewListApplicationBranchesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductDictionaryRequest() (request *CreateProductDictionaryRequest) {
	request = &CreateProductDictionaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateProductDictionary")
	return
}

func NewCreateProductDictionaryResponse() (response *CreateProductDictionaryResponse) {
	response = &CreateProductDictionaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateProductDictionary(request *CreateProductDictionaryRequest) (response *CreateProductDictionaryResponse, err error) {
	if request == nil {
		request = NewCreateProductDictionaryRequest()
	}
	response = NewCreateProductDictionaryResponse()
	err = c.Send(request, response)
	return
}

func NewListApplicationArtifactsRequest() (request *ListApplicationArtifactsRequest) {
	request = &ListApplicationArtifactsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListApplicationArtifacts")
	return
}

func NewListApplicationArtifactsResponse() (response *ListApplicationArtifactsResponse) {
	response = &ListApplicationArtifactsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListApplicationArtifacts(request *ListApplicationArtifactsRequest) (response *ListApplicationArtifactsResponse, err error) {
	if request == nil {
		request = NewListApplicationArtifactsRequest()
	}
	response = NewListApplicationArtifactsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMaterialCustomizedTypeRequest() (request *UpdateMaterialCustomizedTypeRequest) {
	request = &UpdateMaterialCustomizedTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateMaterialCustomizedType")
	return
}

func NewUpdateMaterialCustomizedTypeResponse() (response *UpdateMaterialCustomizedTypeResponse) {
	response = &UpdateMaterialCustomizedTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateMaterialCustomizedType(request *UpdateMaterialCustomizedTypeRequest) (response *UpdateMaterialCustomizedTypeResponse, err error) {
	if request == nil {
		request = NewUpdateMaterialCustomizedTypeRequest()
	}
	response = NewUpdateMaterialCustomizedTypeResponse()
	err = c.Send(request, response)
	return
}

func NewListPlanHistoriesRequest() (request *ListPlanHistoriesRequest) {
	request = &ListPlanHistoriesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListPlanHistories")
	return
}

func NewListPlanHistoriesResponse() (response *ListPlanHistoriesResponse) {
	response = &ListPlanHistoriesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListPlanHistories(request *ListPlanHistoriesRequest) (response *ListPlanHistoriesResponse, err error) {
	if request == nil {
		request = NewListPlanHistoriesRequest()
	}
	response = NewListPlanHistoriesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApplicationInstanceDeployVersionRequest() (request *UpdateApplicationInstanceDeployVersionRequest) {
	request = &UpdateApplicationInstanceDeployVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateApplicationInstanceDeployVersion")
	return
}

func NewUpdateApplicationInstanceDeployVersionResponse() (response *UpdateApplicationInstanceDeployVersionResponse) {
	response = &UpdateApplicationInstanceDeployVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateApplicationInstanceDeployVersion(request *UpdateApplicationInstanceDeployVersionRequest) (response *UpdateApplicationInstanceDeployVersionResponse, err error) {
	if request == nil {
		request = NewUpdateApplicationInstanceDeployVersionRequest()
	}
	response = NewUpdateApplicationInstanceDeployVersionResponse()
	err = c.Send(request, response)
	return
}

func NewListPlanTagsRequest() (request *ListPlanTagsRequest) {
	request = &ListPlanTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListPlanTags")
	return
}

func NewListPlanTagsResponse() (response *ListPlanTagsResponse) {
	response = &ListPlanTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListPlanTags(request *ListPlanTagsRequest) (response *ListPlanTagsResponse, err error) {
	if request == nil {
		request = NewListPlanTagsRequest()
	}
	response = NewListPlanTagsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOperationSheetRelationRequest() (request *UpdateOperationSheetRelationRequest) {
	request = &UpdateOperationSheetRelationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateOperationSheetRelation")
	return
}

func NewUpdateOperationSheetRelationResponse() (response *UpdateOperationSheetRelationResponse) {
	response = &UpdateOperationSheetRelationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateOperationSheetRelation(request *UpdateOperationSheetRelationRequest) (response *UpdateOperationSheetRelationResponse, err error) {
	if request == nil {
		request = NewUpdateOperationSheetRelationRequest()
	}
	response = NewUpdateOperationSheetRelationResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateProjectRegionsRequest() (request *BulkCreateProjectRegionsRequest) {
	request = &BulkCreateProjectRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateProjectRegions")
	return
}

func NewBulkCreateProjectRegionsResponse() (response *BulkCreateProjectRegionsResponse) {
	response = &BulkCreateProjectRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateProjectRegions(request *BulkCreateProjectRegionsRequest) (response *BulkCreateProjectRegionsResponse, err error) {
	if request == nil {
		request = NewBulkCreateProjectRegionsRequest()
	}
	response = NewBulkCreateProjectRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApplicationInstanceStateRequest() (request *CreateApplicationInstanceStateRequest) {
	request = &CreateApplicationInstanceStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateApplicationInstanceState")
	return
}

func NewCreateApplicationInstanceStateResponse() (response *CreateApplicationInstanceStateResponse) {
	response = &CreateApplicationInstanceStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateApplicationInstanceState(request *CreateApplicationInstanceStateRequest) (response *CreateApplicationInstanceStateResponse, err error) {
	if request == nil {
		request = NewCreateApplicationInstanceStateRequest()
	}
	response = NewCreateApplicationInstanceStateResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCommonOperationSheetRequest() (request *DeleteCommonOperationSheetRequest) {
	request = &DeleteCommonOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteCommonOperationSheet")
	return
}

func NewDeleteCommonOperationSheetResponse() (response *DeleteCommonOperationSheetResponse) {
	response = &DeleteCommonOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteCommonOperationSheet(request *DeleteCommonOperationSheetRequest) (response *DeleteCommonOperationSheetResponse, err error) {
	if request == nil {
		request = NewDeleteCommonOperationSheetRequest()
	}
	response = NewDeleteCommonOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApplicationRequest() (request *CreateApplicationRequest) {
	request = &CreateApplicationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateApplication")
	return
}

func NewCreateApplicationResponse() (response *CreateApplicationResponse) {
	response = &CreateApplicationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateApplication(request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
	if request == nil {
		request = NewCreateApplicationRequest()
	}
	response = NewCreateApplicationResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCommonOperationSheetRequest() (request *CreateCommonOperationSheetRequest) {
	request = &CreateCommonOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateCommonOperationSheet")
	return
}

func NewCreateCommonOperationSheetResponse() (response *CreateCommonOperationSheetResponse) {
	response = &CreateCommonOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateCommonOperationSheet(request *CreateCommonOperationSheetRequest) (response *CreateCommonOperationSheetResponse, err error) {
	if request == nil {
		request = NewCreateCommonOperationSheetRequest()
	}
	response = NewCreateCommonOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewListSiteDeployTasksRequest() (request *ListSiteDeployTasksRequest) {
	request = &ListSiteDeployTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListSiteDeployTasks")
	return
}

func NewListSiteDeployTasksResponse() (response *ListSiteDeployTasksResponse) {
	response = &ListSiteDeployTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListSiteDeployTasks(request *ListSiteDeployTasksRequest) (response *ListSiteDeployTasksResponse, err error) {
	if request == nil {
		request = NewListSiteDeployTasksRequest()
	}
	response = NewListSiteDeployTasksResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOperationSheetApprovalRecordRequest() (request *UpdateOperationSheetApprovalRecordRequest) {
	request = &UpdateOperationSheetApprovalRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateOperationSheetApprovalRecord")
	return
}

func NewUpdateOperationSheetApprovalRecordResponse() (response *UpdateOperationSheetApprovalRecordResponse) {
	response = &UpdateOperationSheetApprovalRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateOperationSheetApprovalRecord(request *UpdateOperationSheetApprovalRecordRequest) (response *UpdateOperationSheetApprovalRecordResponse, err error) {
	if request == nil {
		request = NewUpdateOperationSheetApprovalRecordRequest()
	}
	response = NewUpdateOperationSheetApprovalRecordResponse()
	err = c.Send(request, response)
	return
}

func NewGetMaterialSyncTaskRequest() (request *GetMaterialSyncTaskRequest) {
	request = &GetMaterialSyncTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetMaterialSyncTask")
	return
}

func NewGetMaterialSyncTaskResponse() (response *GetMaterialSyncTaskResponse) {
	response = &GetMaterialSyncTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetMaterialSyncTask(request *GetMaterialSyncTaskRequest) (response *GetMaterialSyncTaskResponse, err error) {
	if request == nil {
		request = NewGetMaterialSyncTaskRequest()
	}
	response = NewGetMaterialSyncTaskResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialImportTasksRequest() (request *ListMaterialImportTasksRequest) {
	request = &ListMaterialImportTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListMaterialImportTasks")
	return
}

func NewListMaterialImportTasksResponse() (response *ListMaterialImportTasksResponse) {
	response = &ListMaterialImportTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListMaterialImportTasks(request *ListMaterialImportTasksRequest) (response *ListMaterialImportTasksResponse, err error) {
	if request == nil {
		request = NewListMaterialImportTasksRequest()
	}
	response = NewListMaterialImportTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProjectSiteInfoRequest() (request *DeleteProjectSiteInfoRequest) {
	request = &DeleteProjectSiteInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteProjectSiteInfo")
	return
}

func NewDeleteProjectSiteInfoResponse() (response *DeleteProjectSiteInfoResponse) {
	response = &DeleteProjectSiteInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteProjectSiteInfo(request *DeleteProjectSiteInfoRequest) (response *DeleteProjectSiteInfoResponse, err error) {
	if request == nil {
		request = NewDeleteProjectSiteInfoRequest()
	}
	response = NewDeleteProjectSiteInfoResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialCommonImportDetailsRequest() (request *ListMaterialCommonImportDetailsRequest) {
	request = &ListMaterialCommonImportDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListMaterialCommonImportDetails")
	return
}

func NewListMaterialCommonImportDetailsResponse() (response *ListMaterialCommonImportDetailsResponse) {
	response = &ListMaterialCommonImportDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListMaterialCommonImportDetails(request *ListMaterialCommonImportDetailsRequest) (response *ListMaterialCommonImportDetailsResponse, err error) {
	if request == nil {
		request = NewListMaterialCommonImportDetailsRequest()
	}
	response = NewListMaterialCommonImportDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewListSolutionsRequest() (request *ListSolutionsRequest) {
	request = &ListSolutionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListSolutions")
	return
}

func NewListSolutionsResponse() (response *ListSolutionsResponse) {
	response = &ListSolutionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListSolutions(request *ListSolutionsRequest) (response *ListSolutionsResponse, err error) {
	if request == nil {
		request = NewListSolutionsRequest()
	}
	response = NewListSolutionsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOtaCertificateRequest() (request *CreateOtaCertificateRequest) {
	request = &CreateOtaCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateOtaCertificate")
	return
}

func NewCreateOtaCertificateResponse() (response *CreateOtaCertificateResponse) {
	response = &CreateOtaCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateOtaCertificate(request *CreateOtaCertificateRequest) (response *CreateOtaCertificateResponse, err error) {
	if request == nil {
		request = NewCreateOtaCertificateRequest()
	}
	response = NewCreateOtaCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperationTemplateRequest() (request *DeleteOperationTemplateRequest) {
	request = &DeleteOperationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteOperationTemplate")
	return
}

func NewDeleteOperationTemplateResponse() (response *DeleteOperationTemplateResponse) {
	response = &DeleteOperationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteOperationTemplate(request *DeleteOperationTemplateRequest) (response *DeleteOperationTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteOperationTemplateRequest()
	}
	response = NewDeleteOperationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateMaterialCommonImportDetailsRequest() (request *BulkCreateMaterialCommonImportDetailsRequest) {
	request = &BulkCreateMaterialCommonImportDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateMaterialCommonImportDetails")
	return
}

func NewBulkCreateMaterialCommonImportDetailsResponse() (response *BulkCreateMaterialCommonImportDetailsResponse) {
	response = &BulkCreateMaterialCommonImportDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateMaterialCommonImportDetails(request *BulkCreateMaterialCommonImportDetailsRequest) (response *BulkCreateMaterialCommonImportDetailsResponse, err error) {
	if request == nil {
		request = NewBulkCreateMaterialCommonImportDetailsRequest()
	}
	response = NewBulkCreateMaterialCommonImportDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateTenantsRequest() (request *BulkCreateTenantsRequest) {
	request = &BulkCreateTenantsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateTenants")
	return
}

func NewBulkCreateTenantsResponse() (response *BulkCreateTenantsResponse) {
	response = &BulkCreateTenantsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateTenants(request *BulkCreateTenantsRequest) (response *BulkCreateTenantsResponse, err error) {
	if request == nil {
		request = NewBulkCreateTenantsRequest()
	}
	response = NewBulkCreateTenantsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateFlywaySchemaHistoryRequest() (request *CreateFlywaySchemaHistoryRequest) {
	request = &CreateFlywaySchemaHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateFlywaySchemaHistory")
	return
}

func NewCreateFlywaySchemaHistoryResponse() (response *CreateFlywaySchemaHistoryResponse) {
	response = &CreateFlywaySchemaHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateFlywaySchemaHistory(request *CreateFlywaySchemaHistoryRequest) (response *CreateFlywaySchemaHistoryResponse, err error) {
	if request == nil {
		request = NewCreateFlywaySchemaHistoryRequest()
	}
	response = NewCreateFlywaySchemaHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperationSheetTemplateRequest() (request *DeleteOperationSheetTemplateRequest) {
	request = &DeleteOperationSheetTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteOperationSheetTemplate")
	return
}

func NewDeleteOperationSheetTemplateResponse() (response *DeleteOperationSheetTemplateResponse) {
	response = &DeleteOperationSheetTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteOperationSheetTemplate(request *DeleteOperationSheetTemplateRequest) (response *DeleteOperationSheetTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteOperationSheetTemplateRequest()
	}
	response = NewDeleteOperationSheetTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialMetadataDetailsRequest() (request *ListMaterialMetadataDetailsRequest) {
	request = &ListMaterialMetadataDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListMaterialMetadataDetails")
	return
}

func NewListMaterialMetadataDetailsResponse() (response *ListMaterialMetadataDetailsResponse) {
	response = &ListMaterialMetadataDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListMaterialMetadataDetails(request *ListMaterialMetadataDetailsRequest) (response *ListMaterialMetadataDetailsResponse, err error) {
	if request == nil {
		request = NewListMaterialMetadataDetailsRequest()
	}
	response = NewListMaterialMetadataDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateDagNodeRequest() (request *UpdateDagNodeRequest) {
	request = &UpdateDagNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateDagNode")
	return
}

func NewUpdateDagNodeResponse() (response *UpdateDagNodeResponse) {
	response = &UpdateDagNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateDagNode(request *UpdateDagNodeRequest) (response *UpdateDagNodeResponse, err error) {
	if request == nil {
		request = NewUpdateDagNodeRequest()
	}
	response = NewUpdateDagNodeResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialCleanupsRequest() (request *ListMaterialCleanupsRequest) {
	request = &ListMaterialCleanupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListMaterialCleanups")
	return
}

func NewListMaterialCleanupsResponse() (response *ListMaterialCleanupsResponse) {
	response = &ListMaterialCleanupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListMaterialCleanups(request *ListMaterialCleanupsRequest) (response *ListMaterialCleanupsResponse, err error) {
	if request == nil {
		request = NewListMaterialCleanupsRequest()
	}
	response = NewListMaterialCleanupsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreatePlanApplicationInstancesRequest() (request *BulkCreatePlanApplicationInstancesRequest) {
	request = &BulkCreatePlanApplicationInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreatePlanApplicationInstances")
	return
}

func NewBulkCreatePlanApplicationInstancesResponse() (response *BulkCreatePlanApplicationInstancesResponse) {
	response = &BulkCreatePlanApplicationInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreatePlanApplicationInstances(request *BulkCreatePlanApplicationInstancesRequest) (response *BulkCreatePlanApplicationInstancesResponse, err error) {
	if request == nil {
		request = NewBulkCreatePlanApplicationInstancesRequest()
	}
	response = NewBulkCreatePlanApplicationInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperationSheetRequest() (request *CreateOperationSheetRequest) {
	request = &CreateOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateOperationSheet")
	return
}

func NewCreateOperationSheetResponse() (response *CreateOperationSheetResponse) {
	response = &CreateOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateOperationSheet(request *CreateOperationSheetRequest) (response *CreateOperationSheetResponse, err error) {
	if request == nil {
		request = NewCreateOperationSheetRequest()
	}
	response = NewCreateOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaterialMachineSSHInfoRequest() (request *DeleteMaterialMachineSSHInfoRequest) {
	request = &DeleteMaterialMachineSSHInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteMaterialMachineSSHInfo")
	return
}

func NewDeleteMaterialMachineSSHInfoResponse() (response *DeleteMaterialMachineSSHInfoResponse) {
	response = &DeleteMaterialMachineSSHInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteMaterialMachineSSHInfo(request *DeleteMaterialMachineSSHInfoRequest) (response *DeleteMaterialMachineSSHInfoResponse, err error) {
	if request == nil {
		request = NewDeleteMaterialMachineSSHInfoRequest()
	}
	response = NewDeleteMaterialMachineSSHInfoResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateMaterialImportTasksRequest() (request *BulkCreateMaterialImportTasksRequest) {
	request = &BulkCreateMaterialImportTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateMaterialImportTasks")
	return
}

func NewBulkCreateMaterialImportTasksResponse() (response *BulkCreateMaterialImportTasksResponse) {
	response = &BulkCreateMaterialImportTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateMaterialImportTasks(request *BulkCreateMaterialImportTasksRequest) (response *BulkCreateMaterialImportTasksResponse, err error) {
	if request == nil {
		request = NewBulkCreateMaterialImportTasksRequest()
	}
	response = NewBulkCreateMaterialImportTasksResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateOtaCertificatesRequest() (request *BulkCreateOtaCertificatesRequest) {
	request = &BulkCreateOtaCertificatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateOtaCertificates")
	return
}

func NewBulkCreateOtaCertificatesResponse() (response *BulkCreateOtaCertificatesResponse) {
	response = &BulkCreateOtaCertificatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateOtaCertificates(request *BulkCreateOtaCertificatesRequest) (response *BulkCreateOtaCertificatesResponse, err error) {
	if request == nil {
		request = NewBulkCreateOtaCertificatesRequest()
	}
	response = NewBulkCreateOtaCertificatesResponse()
	err = c.Send(request, response)
	return
}

func NewGetDagNodeRequest() (request *GetDagNodeRequest) {
	request = &GetDagNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetDagNode")
	return
}

func NewGetDagNodeResponse() (response *GetDagNodeResponse) {
	response = &GetDagNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetDagNode(request *GetDagNodeRequest) (response *GetDagNodeResponse, err error) {
	if request == nil {
		request = NewGetDagNodeRequest()
	}
	response = NewGetDagNodeResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateMaterialMetadatasRequest() (request *BulkCreateMaterialMetadatasRequest) {
	request = &BulkCreateMaterialMetadatasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateMaterialMetadatas")
	return
}

func NewBulkCreateMaterialMetadatasResponse() (response *BulkCreateMaterialMetadatasResponse) {
	response = &BulkCreateMaterialMetadatasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateMaterialMetadatas(request *BulkCreateMaterialMetadatasRequest) (response *BulkCreateMaterialMetadatasResponse, err error) {
	if request == nil {
		request = NewBulkCreateMaterialMetadatasRequest()
	}
	response = NewBulkCreateMaterialMetadatasResponse()
	err = c.Send(request, response)
	return
}

func NewGetCommonOperationSheetRequest() (request *GetCommonOperationSheetRequest) {
	request = &GetCommonOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetCommonOperationSheet")
	return
}

func NewGetCommonOperationSheetResponse() (response *GetCommonOperationSheetResponse) {
	response = &GetCommonOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetCommonOperationSheet(request *GetCommonOperationSheetRequest) (response *GetCommonOperationSheetResponse, err error) {
	if request == nil {
		request = NewGetCommonOperationSheetRequest()
	}
	response = NewGetCommonOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewListProductDictionariesRequest() (request *ListProductDictionariesRequest) {
	request = &ListProductDictionariesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProductDictionaries")
	return
}

func NewListProductDictionariesResponse() (response *ListProductDictionariesResponse) {
	response = &ListProductDictionariesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProductDictionaries(request *ListProductDictionariesRequest) (response *ListProductDictionariesResponse, err error) {
	if request == nil {
		request = NewListProductDictionariesRequest()
	}
	response = NewListProductDictionariesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMaterialImportTaskRequest() (request *CreateMaterialImportTaskRequest) {
	request = &CreateMaterialImportTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateMaterialImportTask")
	return
}

func NewCreateMaterialImportTaskResponse() (response *CreateMaterialImportTaskResponse) {
	response = &CreateMaterialImportTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateMaterialImportTask(request *CreateMaterialImportTaskRequest) (response *CreateMaterialImportTaskResponse, err error) {
	if request == nil {
		request = NewCreateMaterialImportTaskRequest()
	}
	response = NewCreateMaterialImportTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApplicationBranchRequest() (request *DeleteApplicationBranchRequest) {
	request = &DeleteApplicationBranchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteApplicationBranch")
	return
}

func NewDeleteApplicationBranchResponse() (response *DeleteApplicationBranchResponse) {
	response = &DeleteApplicationBranchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteApplicationBranch(request *DeleteApplicationBranchRequest) (response *DeleteApplicationBranchResponse, err error) {
	if request == nil {
		request = NewDeleteApplicationBranchRequest()
	}
	response = NewDeleteApplicationBranchResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMaterialCleanupRequest() (request *UpdateMaterialCleanupRequest) {
	request = &UpdateMaterialCleanupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateMaterialCleanup")
	return
}

func NewUpdateMaterialCleanupResponse() (response *UpdateMaterialCleanupResponse) {
	response = &UpdateMaterialCleanupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateMaterialCleanup(request *UpdateMaterialCleanupRequest) (response *UpdateMaterialCleanupResponse, err error) {
	if request == nil {
		request = NewUpdateMaterialCleanupRequest()
	}
	response = NewUpdateMaterialCleanupResponse()
	err = c.Send(request, response)
	return
}

func NewGetProjectSiteInfoDetailRequest() (request *GetProjectSiteInfoDetailRequest) {
	request = &GetProjectSiteInfoDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProjectSiteInfoDetail")
	return
}

func NewGetProjectSiteInfoDetailResponse() (response *GetProjectSiteInfoDetailResponse) {
	response = &GetProjectSiteInfoDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProjectSiteInfoDetail(request *GetProjectSiteInfoDetailRequest) (response *GetProjectSiteInfoDetailResponse, err error) {
	if request == nil {
		request = NewGetProjectSiteInfoDetailRequest()
	}
	response = NewGetProjectSiteInfoDetailResponse()
	err = c.Send(request, response)
	return
}

func NewListCommonOperationSheetsWithDetailRequest() (request *ListCommonOperationSheetsWithDetailRequest) {
	request = &ListCommonOperationSheetsWithDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListCommonOperationSheetsWithDetail")
	return
}

func NewListCommonOperationSheetsWithDetailResponse() (response *ListCommonOperationSheetsWithDetailResponse) {
	response = &ListCommonOperationSheetsWithDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListCommonOperationSheetsWithDetail(request *ListCommonOperationSheetsWithDetailRequest) (response *ListCommonOperationSheetsWithDetailResponse, err error) {
	if request == nil {
		request = NewListCommonOperationSheetsWithDetailRequest()
	}
	response = NewListCommonOperationSheetsWithDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductVersionArtifactTagRequest() (request *CreateProductVersionArtifactTagRequest) {
	request = &CreateProductVersionArtifactTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateProductVersionArtifactTag")
	return
}

func NewCreateProductVersionArtifactTagResponse() (response *CreateProductVersionArtifactTagResponse) {
	response = &CreateProductVersionArtifactTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateProductVersionArtifactTag(request *CreateProductVersionArtifactTagRequest) (response *CreateProductVersionArtifactTagResponse, err error) {
	if request == nil {
		request = NewCreateProductVersionArtifactTagRequest()
	}
	response = NewCreateProductVersionArtifactTagResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApplicationInstanceBackupRequest() (request *DeleteApplicationInstanceBackupRequest) {
	request = &DeleteApplicationInstanceBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteApplicationInstanceBackup")
	return
}

func NewDeleteApplicationInstanceBackupResponse() (response *DeleteApplicationInstanceBackupResponse) {
	response = &DeleteApplicationInstanceBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteApplicationInstanceBackup(request *DeleteApplicationInstanceBackupRequest) (response *DeleteApplicationInstanceBackupResponse, err error) {
	if request == nil {
		request = NewDeleteApplicationInstanceBackupRequest()
	}
	response = NewDeleteApplicationInstanceBackupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaterialImportTaskRequest() (request *DeleteMaterialImportTaskRequest) {
	request = &DeleteMaterialImportTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteMaterialImportTask")
	return
}

func NewDeleteMaterialImportTaskResponse() (response *DeleteMaterialImportTaskResponse) {
	response = &DeleteMaterialImportTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteMaterialImportTask(request *DeleteMaterialImportTaskRequest) (response *DeleteMaterialImportTaskResponse, err error) {
	if request == nil {
		request = NewDeleteMaterialImportTaskRequest()
	}
	response = NewDeleteMaterialImportTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperationPipelineRequest() (request *DeleteOperationPipelineRequest) {
	request = &DeleteOperationPipelineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteOperationPipeline")
	return
}

func NewDeleteOperationPipelineResponse() (response *DeleteOperationPipelineResponse) {
	response = &DeleteOperationPipelineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteOperationPipeline(request *DeleteOperationPipelineRequest) (response *DeleteOperationPipelineResponse, err error) {
	if request == nil {
		request = NewDeleteOperationPipelineRequest()
	}
	response = NewDeleteOperationPipelineResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateProductVersionsRequest() (request *BulkCreateProductVersionsRequest) {
	request = &BulkCreateProductVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateProductVersions")
	return
}

func NewBulkCreateProductVersionsResponse() (response *BulkCreateProductVersionsResponse) {
	response = &BulkCreateProductVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateProductVersions(request *BulkCreateProductVersionsRequest) (response *BulkCreateProductVersionsResponse, err error) {
	if request == nil {
		request = NewBulkCreateProductVersionsRequest()
	}
	response = NewBulkCreateProductVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewGetProjectSiteInfoTopologyRequest() (request *GetProjectSiteInfoTopologyRequest) {
	request = &GetProjectSiteInfoTopologyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProjectSiteInfoTopology")
	return
}

func NewGetProjectSiteInfoTopologyResponse() (response *GetProjectSiteInfoTopologyResponse) {
	response = &GetProjectSiteInfoTopologyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProjectSiteInfoTopology(request *GetProjectSiteInfoTopologyRequest) (response *GetProjectSiteInfoTopologyResponse, err error) {
	if request == nil {
		request = NewGetProjectSiteInfoTopologyRequest()
	}
	response = NewGetProjectSiteInfoTopologyResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateApplicationBranchesRequest() (request *BulkCreateApplicationBranchesRequest) {
	request = &BulkCreateApplicationBranchesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateApplicationBranches")
	return
}

func NewBulkCreateApplicationBranchesResponse() (response *BulkCreateApplicationBranchesResponse) {
	response = &BulkCreateApplicationBranchesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateApplicationBranches(request *BulkCreateApplicationBranchesRequest) (response *BulkCreateApplicationBranchesResponse, err error) {
	if request == nil {
		request = NewBulkCreateApplicationBranchesRequest()
	}
	response = NewBulkCreateApplicationBranchesResponse()
	err = c.Send(request, response)
	return
}

func NewListSolutionVersionsRequest() (request *ListSolutionVersionsRequest) {
	request = &ListSolutionVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListSolutionVersions")
	return
}

func NewListSolutionVersionsResponse() (response *ListSolutionVersionsResponse) {
	response = &ListSolutionVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListSolutionVersions(request *ListSolutionVersionsRequest) (response *ListSolutionVersionsResponse, err error) {
	if request == nil {
		request = NewListSolutionVersionsRequest()
	}
	response = NewListSolutionVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApplicationCollectStatusRequest() (request *UpdateApplicationCollectStatusRequest) {
	request = &UpdateApplicationCollectStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateApplicationCollectStatus")
	return
}

func NewUpdateApplicationCollectStatusResponse() (response *UpdateApplicationCollectStatusResponse) {
	response = &UpdateApplicationCollectStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateApplicationCollectStatus(request *UpdateApplicationCollectStatusRequest) (response *UpdateApplicationCollectStatusResponse, err error) {
	if request == nil {
		request = NewUpdateApplicationCollectStatusRequest()
	}
	response = NewUpdateApplicationCollectStatusResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePlanApplicationInstanceRequest() (request *UpdatePlanApplicationInstanceRequest) {
	request = &UpdatePlanApplicationInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdatePlanApplicationInstance")
	return
}

func NewUpdatePlanApplicationInstanceResponse() (response *UpdatePlanApplicationInstanceResponse) {
	response = &UpdatePlanApplicationInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdatePlanApplicationInstance(request *UpdatePlanApplicationInstanceRequest) (response *UpdatePlanApplicationInstanceResponse, err error) {
	if request == nil {
		request = NewUpdatePlanApplicationInstanceRequest()
	}
	response = NewUpdatePlanApplicationInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProjectRegionRequest() (request *DeleteProjectRegionRequest) {
	request = &DeleteProjectRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteProjectRegion")
	return
}

func NewDeleteProjectRegionResponse() (response *DeleteProjectRegionResponse) {
	response = &DeleteProjectRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteProjectRegion(request *DeleteProjectRegionRequest) (response *DeleteProjectRegionResponse, err error) {
	if request == nil {
		request = NewDeleteProjectRegionRequest()
	}
	response = NewDeleteProjectRegionResponse()
	err = c.Send(request, response)
	return
}

func NewGetMaterialCommonImportDetailRequest() (request *GetMaterialCommonImportDetailRequest) {
	request = &GetMaterialCommonImportDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetMaterialCommonImportDetail")
	return
}

func NewGetMaterialCommonImportDetailResponse() (response *GetMaterialCommonImportDetailResponse) {
	response = &GetMaterialCommonImportDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetMaterialCommonImportDetail(request *GetMaterialCommonImportDetailRequest) (response *GetMaterialCommonImportDetailResponse, err error) {
	if request == nil {
		request = NewGetMaterialCommonImportDetailRequest()
	}
	response = NewGetMaterialCommonImportDetailResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApplicationBranchRequest() (request *UpdateApplicationBranchRequest) {
	request = &UpdateApplicationBranchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateApplicationBranch")
	return
}

func NewUpdateApplicationBranchResponse() (response *UpdateApplicationBranchResponse) {
	response = &UpdateApplicationBranchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateApplicationBranch(request *UpdateApplicationBranchRequest) (response *UpdateApplicationBranchResponse, err error) {
	if request == nil {
		request = NewUpdateApplicationBranchRequest()
	}
	response = NewUpdateApplicationBranchResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePlanServerInfoRequest() (request *DeletePlanServerInfoRequest) {
	request = &DeletePlanServerInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeletePlanServerInfo")
	return
}

func NewDeletePlanServerInfoResponse() (response *DeletePlanServerInfoResponse) {
	response = &DeletePlanServerInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeletePlanServerInfo(request *DeletePlanServerInfoRequest) (response *DeletePlanServerInfoResponse, err error) {
	if request == nil {
		request = NewDeletePlanServerInfoRequest()
	}
	response = NewDeletePlanServerInfoResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateTenantRequest() (request *UpdateTenantRequest) {
	request = &UpdateTenantRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateTenant")
	return
}

func NewUpdateTenantResponse() (response *UpdateTenantResponse) {
	response = &UpdateTenantResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateTenant(request *UpdateTenantRequest) (response *UpdateTenantResponse, err error) {
	if request == nil {
		request = NewUpdateTenantRequest()
	}
	response = NewUpdateTenantResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationInstanceDeployVersionRequest() (request *GetApplicationInstanceDeployVersionRequest) {
	request = &GetApplicationInstanceDeployVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetApplicationInstanceDeployVersion")
	return
}

func NewGetApplicationInstanceDeployVersionResponse() (response *GetApplicationInstanceDeployVersionResponse) {
	response = &GetApplicationInstanceDeployVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetApplicationInstanceDeployVersion(request *GetApplicationInstanceDeployVersionRequest) (response *GetApplicationInstanceDeployVersionResponse, err error) {
	if request == nil {
		request = NewGetApplicationInstanceDeployVersionRequest()
	}
	response = NewGetApplicationInstanceDeployVersionResponse()
	err = c.Send(request, response)
	return
}

func NewListApplicationsRequest() (request *ListApplicationsRequest) {
	request = &ListApplicationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListApplications")
	return
}

func NewListApplicationsResponse() (response *ListApplicationsResponse) {
	response = &ListApplicationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListApplications(request *ListApplicationsRequest) (response *ListApplicationsResponse, err error) {
	if request == nil {
		request = NewListApplicationsRequest()
	}
	response = NewListApplicationsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApplicationRunningDataCollectionRequest() (request *DeleteApplicationRunningDataCollectionRequest) {
	request = &DeleteApplicationRunningDataCollectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteApplicationRunningDataCollection")
	return
}

func NewDeleteApplicationRunningDataCollectionResponse() (response *DeleteApplicationRunningDataCollectionResponse) {
	response = &DeleteApplicationRunningDataCollectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteApplicationRunningDataCollection(request *DeleteApplicationRunningDataCollectionRequest) (response *DeleteApplicationRunningDataCollectionResponse, err error) {
	if request == nil {
		request = NewDeleteApplicationRunningDataCollectionRequest()
	}
	response = NewDeleteApplicationRunningDataCollectionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductVersionRequest() (request *CreateProductVersionRequest) {
	request = &CreateProductVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateProductVersion")
	return
}

func NewCreateProductVersionResponse() (response *CreateProductVersionResponse) {
	response = &CreateProductVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateProductVersion(request *CreateProductVersionRequest) (response *CreateProductVersionResponse, err error) {
	if request == nil {
		request = NewCreateProductVersionRequest()
	}
	response = NewCreateProductVersionResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialApplicationImportDetailsRequest() (request *ListMaterialApplicationImportDetailsRequest) {
	request = &ListMaterialApplicationImportDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListMaterialApplicationImportDetails")
	return
}

func NewListMaterialApplicationImportDetailsResponse() (response *ListMaterialApplicationImportDetailsResponse) {
	response = &ListMaterialApplicationImportDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListMaterialApplicationImportDetails(request *ListMaterialApplicationImportDetailsRequest) (response *ListMaterialApplicationImportDetailsResponse, err error) {
	if request == nil {
		request = NewListMaterialApplicationImportDetailsRequest()
	}
	response = NewListMaterialApplicationImportDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialSyncTasksRequest() (request *ListMaterialSyncTasksRequest) {
	request = &ListMaterialSyncTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListMaterialSyncTasks")
	return
}

func NewListMaterialSyncTasksResponse() (response *ListMaterialSyncTasksResponse) {
	response = &ListMaterialSyncTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListMaterialSyncTasks(request *ListMaterialSyncTasksRequest) (response *ListMaterialSyncTasksResponse, err error) {
	if request == nil {
		request = NewListMaterialSyncTasksRequest()
	}
	response = NewListMaterialSyncTasksResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductVersionApplicationPackageRelRequest() (request *CreateProductVersionApplicationPackageRelRequest) {
	request = &CreateProductVersionApplicationPackageRelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateProductVersionApplicationPackageRel")
	return
}

func NewCreateProductVersionApplicationPackageRelResponse() (response *CreateProductVersionApplicationPackageRelResponse) {
	response = &CreateProductVersionApplicationPackageRelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateProductVersionApplicationPackageRel(request *CreateProductVersionApplicationPackageRelRequest) (response *CreateProductVersionApplicationPackageRelResponse, err error) {
	if request == nil {
		request = NewCreateProductVersionApplicationPackageRelRequest()
	}
	response = NewCreateProductVersionApplicationPackageRelResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOrchTemplateLabelRequest() (request *DeleteOrchTemplateLabelRequest) {
	request = &DeleteOrchTemplateLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteOrchTemplateLabel")
	return
}

func NewDeleteOrchTemplateLabelResponse() (response *DeleteOrchTemplateLabelResponse) {
	response = &DeleteOrchTemplateLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteOrchTemplateLabel(request *DeleteOrchTemplateLabelRequest) (response *DeleteOrchTemplateLabelResponse, err error) {
	if request == nil {
		request = NewDeleteOrchTemplateLabelRequest()
	}
	response = NewDeleteOrchTemplateLabelResponse()
	err = c.Send(request, response)
	return
}

func NewListPlanCustomConfigsRequest() (request *ListPlanCustomConfigsRequest) {
	request = &ListPlanCustomConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListPlanCustomConfigs")
	return
}

func NewListPlanCustomConfigsResponse() (response *ListPlanCustomConfigsResponse) {
	response = &ListPlanCustomConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListPlanCustomConfigs(request *ListPlanCustomConfigsRequest) (response *ListPlanCustomConfigsResponse, err error) {
	if request == nil {
		request = NewListPlanCustomConfigsRequest()
	}
	response = NewListPlanCustomConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductTemplateRequest() (request *GetProductTemplateRequest) {
	request = &GetProductTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProductTemplate")
	return
}

func NewGetProductTemplateResponse() (response *GetProductTemplateResponse) {
	response = &GetProductTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProductTemplate(request *GetProductTemplateRequest) (response *GetProductTemplateResponse, err error) {
	if request == nil {
		request = NewGetProductTemplateRequest()
	}
	response = NewGetProductTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateApplicationInstanceBackupsRequest() (request *BulkCreateApplicationInstanceBackupsRequest) {
	request = &BulkCreateApplicationInstanceBackupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateApplicationInstanceBackups")
	return
}

func NewBulkCreateApplicationInstanceBackupsResponse() (response *BulkCreateApplicationInstanceBackupsResponse) {
	response = &BulkCreateApplicationInstanceBackupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateApplicationInstanceBackups(request *BulkCreateApplicationInstanceBackupsRequest) (response *BulkCreateApplicationInstanceBackupsResponse, err error) {
	if request == nil {
		request = NewBulkCreateApplicationInstanceBackupsRequest()
	}
	response = NewBulkCreateApplicationInstanceBackupsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateProductInstanceBackupsRequest() (request *BulkCreateProductInstanceBackupsRequest) {
	request = &BulkCreateProductInstanceBackupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateProductInstanceBackups")
	return
}

func NewBulkCreateProductInstanceBackupsResponse() (response *BulkCreateProductInstanceBackupsResponse) {
	response = &BulkCreateProductInstanceBackupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateProductInstanceBackups(request *BulkCreateProductInstanceBackupsRequest) (response *BulkCreateProductInstanceBackupsResponse, err error) {
	if request == nil {
		request = NewBulkCreateProductInstanceBackupsRequest()
	}
	response = NewBulkCreateProductInstanceBackupsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperationSheetRelationRequest() (request *CreateOperationSheetRelationRequest) {
	request = &CreateOperationSheetRelationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateOperationSheetRelation")
	return
}

func NewCreateOperationSheetRelationResponse() (response *CreateOperationSheetRelationResponse) {
	response = &CreateOperationSheetRelationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateOperationSheetRelation(request *CreateOperationSheetRelationRequest) (response *CreateOperationSheetRelationResponse, err error) {
	if request == nil {
		request = NewCreateOperationSheetRelationRequest()
	}
	response = NewCreateOperationSheetRelationResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApplicationInstanceDeployVersionRequest() (request *DeleteApplicationInstanceDeployVersionRequest) {
	request = &DeleteApplicationInstanceDeployVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteApplicationInstanceDeployVersion")
	return
}

func NewDeleteApplicationInstanceDeployVersionResponse() (response *DeleteApplicationInstanceDeployVersionResponse) {
	response = &DeleteApplicationInstanceDeployVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteApplicationInstanceDeployVersion(request *DeleteApplicationInstanceDeployVersionRequest) (response *DeleteApplicationInstanceDeployVersionResponse, err error) {
	if request == nil {
		request = NewDeleteApplicationInstanceDeployVersionRequest()
	}
	response = NewDeleteApplicationInstanceDeployVersionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperationBuiltinDataSchemaRequest() (request *CreateOperationBuiltinDataSchemaRequest) {
	request = &CreateOperationBuiltinDataSchemaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateOperationBuiltinDataSchema")
	return
}

func NewCreateOperationBuiltinDataSchemaResponse() (response *CreateOperationBuiltinDataSchemaResponse) {
	response = &CreateOperationBuiltinDataSchemaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateOperationBuiltinDataSchema(request *CreateOperationBuiltinDataSchemaRequest) (response *CreateOperationBuiltinDataSchemaResponse, err error) {
	if request == nil {
		request = NewCreateOperationBuiltinDataSchemaRequest()
	}
	response = NewCreateOperationBuiltinDataSchemaResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePlanProductInstanceRequest() (request *CreatePlanProductInstanceRequest) {
	request = &CreatePlanProductInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreatePlanProductInstance")
	return
}

func NewCreatePlanProductInstanceResponse() (response *CreatePlanProductInstanceResponse) {
	response = &CreatePlanProductInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreatePlanProductInstance(request *CreatePlanProductInstanceRequest) (response *CreatePlanProductInstanceResponse, err error) {
	if request == nil {
		request = NewCreatePlanProductInstanceRequest()
	}
	response = NewCreatePlanProductInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateProductOperationJobsRequest() (request *BulkCreateProductOperationJobsRequest) {
	request = &BulkCreateProductOperationJobsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateProductOperationJobs")
	return
}

func NewBulkCreateProductOperationJobsResponse() (response *BulkCreateProductOperationJobsResponse) {
	response = &BulkCreateProductOperationJobsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateProductOperationJobs(request *BulkCreateProductOperationJobsRequest) (response *BulkCreateProductOperationJobsResponse, err error) {
	if request == nil {
		request = NewBulkCreateProductOperationJobsRequest()
	}
	response = NewBulkCreateProductOperationJobsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreatePlanTagsRequest() (request *BulkCreatePlanTagsRequest) {
	request = &BulkCreatePlanTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreatePlanTags")
	return
}

func NewBulkCreatePlanTagsResponse() (response *BulkCreatePlanTagsResponse) {
	response = &BulkCreatePlanTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreatePlanTags(request *BulkCreatePlanTagsRequest) (response *BulkCreatePlanTagsResponse, err error) {
	if request == nil {
		request = NewBulkCreatePlanTagsRequest()
	}
	response = NewBulkCreatePlanTagsResponse()
	err = c.Send(request, response)
	return
}

func NewGetPlanSubsystemInstanceRequest() (request *GetPlanSubsystemInstanceRequest) {
	request = &GetPlanSubsystemInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetPlanSubsystemInstance")
	return
}

func NewGetPlanSubsystemInstanceResponse() (response *GetPlanSubsystemInstanceResponse) {
	response = &GetPlanSubsystemInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetPlanSubsystemInstance(request *GetPlanSubsystemInstanceRequest) (response *GetPlanSubsystemInstanceResponse, err error) {
	if request == nil {
		request = NewGetPlanSubsystemInstanceRequest()
	}
	response = NewGetPlanSubsystemInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialCustomizedTypesRequest() (request *ListMaterialCustomizedTypesRequest) {
	request = &ListMaterialCustomizedTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListMaterialCustomizedTypes")
	return
}

func NewListMaterialCustomizedTypesResponse() (response *ListMaterialCustomizedTypesResponse) {
	response = &ListMaterialCustomizedTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListMaterialCustomizedTypes(request *ListMaterialCustomizedTypesRequest) (response *ListMaterialCustomizedTypesResponse, err error) {
	if request == nil {
		request = NewListMaterialCustomizedTypesRequest()
	}
	response = NewListMaterialCustomizedTypesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMaterialImportTaskRequest() (request *UpdateMaterialImportTaskRequest) {
	request = &UpdateMaterialImportTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateMaterialImportTask")
	return
}

func NewUpdateMaterialImportTaskResponse() (response *UpdateMaterialImportTaskResponse) {
	response = &UpdateMaterialImportTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateMaterialImportTask(request *UpdateMaterialImportTaskRequest) (response *UpdateMaterialImportTaskResponse, err error) {
	if request == nil {
		request = NewUpdateMaterialImportTaskRequest()
	}
	response = NewUpdateMaterialImportTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrchTemplateLabelRequest() (request *CreateOrchTemplateLabelRequest) {
	request = &CreateOrchTemplateLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateOrchTemplateLabel")
	return
}

func NewCreateOrchTemplateLabelResponse() (response *CreateOrchTemplateLabelResponse) {
	response = &CreateOrchTemplateLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateOrchTemplateLabel(request *CreateOrchTemplateLabelRequest) (response *CreateOrchTemplateLabelResponse, err error) {
	if request == nil {
		request = NewCreateOrchTemplateLabelRequest()
	}
	response = NewCreateOrchTemplateLabelResponse()
	err = c.Send(request, response)
	return
}

func NewGetPlanServerInfoRequest() (request *GetPlanServerInfoRequest) {
	request = &GetPlanServerInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetPlanServerInfo")
	return
}

func NewGetPlanServerInfoResponse() (response *GetPlanServerInfoResponse) {
	response = &GetPlanServerInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetPlanServerInfo(request *GetPlanServerInfoRequest) (response *GetPlanServerInfoResponse, err error) {
	if request == nil {
		request = NewGetPlanServerInfoRequest()
	}
	response = NewGetPlanServerInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGetOperationTemplateRequest() (request *GetOperationTemplateRequest) {
	request = &GetOperationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetOperationTemplate")
	return
}

func NewGetOperationTemplateResponse() (response *GetOperationTemplateResponse) {
	response = &GetOperationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetOperationTemplate(request *GetOperationTemplateRequest) (response *GetOperationTemplateResponse, err error) {
	if request == nil {
		request = NewGetOperationTemplateRequest()
	}
	response = NewGetOperationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewGetPlanSiteRequest() (request *GetPlanSiteRequest) {
	request = &GetPlanSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetPlanSite")
	return
}

func NewGetPlanSiteResponse() (response *GetPlanSiteResponse) {
	response = &GetPlanSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetPlanSite(request *GetPlanSiteRequest) (response *GetPlanSiteResponse, err error) {
	if request == nil {
		request = NewGetPlanSiteRequest()
	}
	response = NewGetPlanSiteResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateSiteProductVersionRelsRequest() (request *BulkCreateSiteProductVersionRelsRequest) {
	request = &BulkCreateSiteProductVersionRelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateSiteProductVersionRels")
	return
}

func NewBulkCreateSiteProductVersionRelsResponse() (response *BulkCreateSiteProductVersionRelsResponse) {
	response = &BulkCreateSiteProductVersionRelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateSiteProductVersionRels(request *BulkCreateSiteProductVersionRelsRequest) (response *BulkCreateSiteProductVersionRelsResponse, err error) {
	if request == nil {
		request = NewBulkCreateSiteProductVersionRelsRequest()
	}
	response = NewBulkCreateSiteProductVersionRelsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteFlywaySchemaHistoryRequest() (request *DeleteFlywaySchemaHistoryRequest) {
	request = &DeleteFlywaySchemaHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteFlywaySchemaHistory")
	return
}

func NewDeleteFlywaySchemaHistoryResponse() (response *DeleteFlywaySchemaHistoryResponse) {
	response = &DeleteFlywaySchemaHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteFlywaySchemaHistory(request *DeleteFlywaySchemaHistoryRequest) (response *DeleteFlywaySchemaHistoryResponse, err error) {
	if request == nil {
		request = NewDeleteFlywaySchemaHistoryRequest()
	}
	response = NewDeleteFlywaySchemaHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewGetPlanApplicationRuntimeNameRequest() (request *GetPlanApplicationRuntimeNameRequest) {
	request = &GetPlanApplicationRuntimeNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetPlanApplicationRuntimeName")
	return
}

func NewGetPlanApplicationRuntimeNameResponse() (response *GetPlanApplicationRuntimeNameResponse) {
	response = &GetPlanApplicationRuntimeNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetPlanApplicationRuntimeName(request *GetPlanApplicationRuntimeNameRequest) (response *GetPlanApplicationRuntimeNameResponse, err error) {
	if request == nil {
		request = NewGetPlanApplicationRuntimeNameRequest()
	}
	response = NewGetPlanApplicationRuntimeNameResponse()
	err = c.Send(request, response)
	return
}

func NewListDagNodesRequest() (request *ListDagNodesRequest) {
	request = &ListDagNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListDagNodes")
	return
}

func NewListDagNodesResponse() (response *ListDagNodesResponse) {
	response = &ListDagNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListDagNodes(request *ListDagNodesRequest) (response *ListDagNodesResponse, err error) {
	if request == nil {
		request = NewListDagNodesRequest()
	}
	response = NewListDagNodesResponse()
	err = c.Send(request, response)
	return
}

func NewListOtaCertificatesRequest() (request *ListOtaCertificatesRequest) {
	request = &ListOtaCertificatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListOtaCertificates")
	return
}

func NewListOtaCertificatesResponse() (response *ListOtaCertificatesResponse) {
	response = &ListOtaCertificatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListOtaCertificates(request *ListOtaCertificatesRequest) (response *ListOtaCertificatesResponse, err error) {
	if request == nil {
		request = NewListOtaCertificatesRequest()
	}
	response = NewListOtaCertificatesResponse()
	err = c.Send(request, response)
	return
}

func NewListPlanSubsystemInstancesRequest() (request *ListPlanSubsystemInstancesRequest) {
	request = &ListPlanSubsystemInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListPlanSubsystemInstances")
	return
}

func NewListPlanSubsystemInstancesResponse() (response *ListPlanSubsystemInstancesResponse) {
	response = &ListPlanSubsystemInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListPlanSubsystemInstances(request *ListPlanSubsystemInstancesRequest) (response *ListPlanSubsystemInstancesResponse, err error) {
	if request == nil {
		request = NewListPlanSubsystemInstancesRequest()
	}
	response = NewListPlanSubsystemInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSolutionRequest() (request *UpdateSolutionRequest) {
	request = &UpdateSolutionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateSolution")
	return
}

func NewUpdateSolutionResponse() (response *UpdateSolutionResponse) {
	response = &UpdateSolutionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateSolution(request *UpdateSolutionRequest) (response *UpdateSolutionResponse, err error) {
	if request == nil {
		request = NewUpdateSolutionRequest()
	}
	response = NewUpdateSolutionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaterialSyncTaskRequest() (request *DeleteMaterialSyncTaskRequest) {
	request = &DeleteMaterialSyncTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteMaterialSyncTask")
	return
}

func NewDeleteMaterialSyncTaskResponse() (response *DeleteMaterialSyncTaskResponse) {
	response = &DeleteMaterialSyncTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteMaterialSyncTask(request *DeleteMaterialSyncTaskRequest) (response *DeleteMaterialSyncTaskResponse, err error) {
	if request == nil {
		request = NewDeleteMaterialSyncTaskRequest()
	}
	response = NewDeleteMaterialSyncTaskResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateOperationPipelinesRequest() (request *BulkCreateOperationPipelinesRequest) {
	request = &BulkCreateOperationPipelinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateOperationPipelines")
	return
}

func NewBulkCreateOperationPipelinesResponse() (response *BulkCreateOperationPipelinesResponse) {
	response = &BulkCreateOperationPipelinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateOperationPipelines(request *BulkCreateOperationPipelinesRequest) (response *BulkCreateOperationPipelinesResponse, err error) {
	if request == nil {
		request = NewBulkCreateOperationPipelinesRequest()
	}
	response = NewBulkCreateOperationPipelinesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaterialSiteRequest() (request *DeleteMaterialSiteRequest) {
	request = &DeleteMaterialSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteMaterialSite")
	return
}

func NewDeleteMaterialSiteResponse() (response *DeleteMaterialSiteResponse) {
	response = &DeleteMaterialSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteMaterialSite(request *DeleteMaterialSiteRequest) (response *DeleteMaterialSiteResponse, err error) {
	if request == nil {
		request = NewDeleteMaterialSiteRequest()
	}
	response = NewDeleteMaterialSiteResponse()
	err = c.Send(request, response)
	return
}

func NewGetProjectRegionDetailRequest() (request *GetProjectRegionDetailRequest) {
	request = &GetProjectRegionDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProjectRegionDetail")
	return
}

func NewGetProjectRegionDetailResponse() (response *GetProjectRegionDetailResponse) {
	response = &GetProjectRegionDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProjectRegionDetail(request *GetProjectRegionDetailRequest) (response *GetProjectRegionDetailResponse, err error) {
	if request == nil {
		request = NewGetProjectRegionDetailRequest()
	}
	response = NewGetProjectRegionDetailResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationRunningPlanRequest() (request *GetApplicationRunningPlanRequest) {
	request = &GetApplicationRunningPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetApplicationRunningPlan")
	return
}

func NewGetApplicationRunningPlanResponse() (response *GetApplicationRunningPlanResponse) {
	response = &GetApplicationRunningPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetApplicationRunningPlan(request *GetApplicationRunningPlanRequest) (response *GetApplicationRunningPlanResponse, err error) {
	if request == nil {
		request = NewGetApplicationRunningPlanRequest()
	}
	response = NewGetApplicationRunningPlanResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProductInstanceBackupRequest() (request *UpdateProductInstanceBackupRequest) {
	request = &UpdateProductInstanceBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateProductInstanceBackup")
	return
}

func NewUpdateProductInstanceBackupResponse() (response *UpdateProductInstanceBackupResponse) {
	response = &UpdateProductInstanceBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateProductInstanceBackup(request *UpdateProductInstanceBackupRequest) (response *UpdateProductInstanceBackupResponse, err error) {
	if request == nil {
		request = NewUpdateProductInstanceBackupRequest()
	}
	response = NewUpdateProductInstanceBackupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApplicationCollectStatusRequest() (request *DeleteApplicationCollectStatusRequest) {
	request = &DeleteApplicationCollectStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteApplicationCollectStatus")
	return
}

func NewDeleteApplicationCollectStatusResponse() (response *DeleteApplicationCollectStatusResponse) {
	response = &DeleteApplicationCollectStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteApplicationCollectStatus(request *DeleteApplicationCollectStatusRequest) (response *DeleteApplicationCollectStatusResponse, err error) {
	if request == nil {
		request = NewDeleteApplicationCollectStatusRequest()
	}
	response = NewDeleteApplicationCollectStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePlanTagRequest() (request *CreatePlanTagRequest) {
	request = &CreatePlanTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreatePlanTag")
	return
}

func NewCreatePlanTagResponse() (response *CreatePlanTagResponse) {
	response = &CreatePlanTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreatePlanTag(request *CreatePlanTagRequest) (response *CreatePlanTagResponse, err error) {
	if request == nil {
		request = NewCreatePlanTagRequest()
	}
	response = NewCreatePlanTagResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductInfoRequest() (request *CreateProductInfoRequest) {
	request = &CreateProductInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateProductInfo")
	return
}

func NewCreateProductInfoResponse() (response *CreateProductInfoResponse) {
	response = &CreateProductInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateProductInfo(request *CreateProductInfoRequest) (response *CreateProductInfoResponse, err error) {
	if request == nil {
		request = NewCreateProductInfoRequest()
	}
	response = NewCreateProductInfoResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOperationSheetLabelRequest() (request *UpdateOperationSheetLabelRequest) {
	request = &UpdateOperationSheetLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateOperationSheetLabel")
	return
}

func NewUpdateOperationSheetLabelResponse() (response *UpdateOperationSheetLabelResponse) {
	response = &UpdateOperationSheetLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateOperationSheetLabel(request *UpdateOperationSheetLabelRequest) (response *UpdateOperationSheetLabelResponse, err error) {
	if request == nil {
		request = NewUpdateOperationSheetLabelRequest()
	}
	response = NewUpdateOperationSheetLabelResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductTemplateRequest() (request *DeleteProductTemplateRequest) {
	request = &DeleteProductTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteProductTemplate")
	return
}

func NewDeleteProductTemplateResponse() (response *DeleteProductTemplateResponse) {
	response = &DeleteProductTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteProductTemplate(request *DeleteProductTemplateRequest) (response *DeleteProductTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteProductTemplateRequest()
	}
	response = NewDeleteProductTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePlanCustomConfigRequest() (request *DeletePlanCustomConfigRequest) {
	request = &DeletePlanCustomConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeletePlanCustomConfig")
	return
}

func NewDeletePlanCustomConfigResponse() (response *DeletePlanCustomConfigResponse) {
	response = &DeletePlanCustomConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeletePlanCustomConfig(request *DeletePlanCustomConfigRequest) (response *DeletePlanCustomConfigResponse, err error) {
	if request == nil {
		request = NewDeletePlanCustomConfigRequest()
	}
	response = NewDeletePlanCustomConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePlanGlobalConfigRequest() (request *CreatePlanGlobalConfigRequest) {
	request = &CreatePlanGlobalConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreatePlanGlobalConfig")
	return
}

func NewCreatePlanGlobalConfigResponse() (response *CreatePlanGlobalConfigResponse) {
	response = &CreatePlanGlobalConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreatePlanGlobalConfig(request *CreatePlanGlobalConfigRequest) (response *CreatePlanGlobalConfigResponse, err error) {
	if request == nil {
		request = NewCreatePlanGlobalConfigRequest()
	}
	response = NewCreatePlanGlobalConfigResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApplicationArtifactChartHistoryRequest() (request *UpdateApplicationArtifactChartHistoryRequest) {
	request = &UpdateApplicationArtifactChartHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateApplicationArtifactChartHistory")
	return
}

func NewUpdateApplicationArtifactChartHistoryResponse() (response *UpdateApplicationArtifactChartHistoryResponse) {
	response = &UpdateApplicationArtifactChartHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateApplicationArtifactChartHistory(request *UpdateApplicationArtifactChartHistoryRequest) (response *UpdateApplicationArtifactChartHistoryResponse, err error) {
	if request == nil {
		request = NewUpdateApplicationArtifactChartHistoryRequest()
	}
	response = NewUpdateApplicationArtifactChartHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePlanApplicationRuntimeNameRequest() (request *UpdatePlanApplicationRuntimeNameRequest) {
	request = &UpdatePlanApplicationRuntimeNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdatePlanApplicationRuntimeName")
	return
}

func NewUpdatePlanApplicationRuntimeNameResponse() (response *UpdatePlanApplicationRuntimeNameResponse) {
	response = &UpdatePlanApplicationRuntimeNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdatePlanApplicationRuntimeName(request *UpdatePlanApplicationRuntimeNameRequest) (response *UpdatePlanApplicationRuntimeNameResponse, err error) {
	if request == nil {
		request = NewUpdatePlanApplicationRuntimeNameRequest()
	}
	response = NewUpdatePlanApplicationRuntimeNameResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAtomicRequest() (request *CreateAtomicRequest) {
	request = &CreateAtomicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateAtomic")
	return
}

func NewCreateAtomicResponse() (response *CreateAtomicResponse) {
	response = &CreateAtomicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateAtomic(request *CreateAtomicRequest) (response *CreateAtomicResponse, err error) {
	if request == nil {
		request = NewCreateAtomicRequest()
	}
	response = NewCreateAtomicResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApplicationRunningDataCollectionRequest() (request *CreateApplicationRunningDataCollectionRequest) {
	request = &CreateApplicationRunningDataCollectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateApplicationRunningDataCollection")
	return
}

func NewCreateApplicationRunningDataCollectionResponse() (response *CreateApplicationRunningDataCollectionResponse) {
	response = &CreateApplicationRunningDataCollectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateApplicationRunningDataCollection(request *CreateApplicationRunningDataCollectionRequest) (response *CreateApplicationRunningDataCollectionResponse, err error) {
	if request == nil {
		request = NewCreateApplicationRunningDataCollectionRequest()
	}
	response = NewCreateApplicationRunningDataCollectionResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreatePlanProductInstancesRequest() (request *BulkCreatePlanProductInstancesRequest) {
	request = &BulkCreatePlanProductInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreatePlanProductInstances")
	return
}

func NewBulkCreatePlanProductInstancesResponse() (response *BulkCreatePlanProductInstancesResponse) {
	response = &BulkCreatePlanProductInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreatePlanProductInstances(request *BulkCreatePlanProductInstancesRequest) (response *BulkCreatePlanProductInstancesResponse, err error) {
	if request == nil {
		request = NewBulkCreatePlanProductInstancesRequest()
	}
	response = NewBulkCreatePlanProductInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewGetProjectTopologyRequest() (request *GetProjectTopologyRequest) {
	request = &GetProjectTopologyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProjectTopology")
	return
}

func NewGetProjectTopologyResponse() (response *GetProjectTopologyResponse) {
	response = &GetProjectTopologyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProjectTopology(request *GetProjectTopologyRequest) (response *GetProjectTopologyResponse, err error) {
	if request == nil {
		request = NewGetProjectTopologyRequest()
	}
	response = NewGetProjectTopologyResponse()
	err = c.Send(request, response)
	return
}

func NewGetProjectTopologyDetailRequest() (request *GetProjectTopologyDetailRequest) {
	request = &GetProjectTopologyDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProjectTopologyDetail")
	return
}

func NewGetProjectTopologyDetailResponse() (response *GetProjectTopologyDetailResponse) {
	response = &GetProjectTopologyDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProjectTopologyDetail(request *GetProjectTopologyDetailRequest) (response *GetProjectTopologyDetailResponse, err error) {
	if request == nil {
		request = NewGetProjectTopologyDetailRequest()
	}
	response = NewGetProjectTopologyDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaterialInfoRequest() (request *DeleteMaterialInfoRequest) {
	request = &DeleteMaterialInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteMaterialInfo")
	return
}

func NewDeleteMaterialInfoResponse() (response *DeleteMaterialInfoResponse) {
	response = &DeleteMaterialInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteMaterialInfo(request *DeleteMaterialInfoRequest) (response *DeleteMaterialInfoResponse, err error) {
	if request == nil {
		request = NewDeleteMaterialInfoRequest()
	}
	response = NewDeleteMaterialInfoResponse()
	err = c.Send(request, response)
	return
}

func NewListApplicationPackageCcDeclaresRequest() (request *ListApplicationPackageCcDeclaresRequest) {
	request = &ListApplicationPackageCcDeclaresRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListApplicationPackageCcDeclares")
	return
}

func NewListApplicationPackageCcDeclaresResponse() (response *ListApplicationPackageCcDeclaresResponse) {
	response = &ListApplicationPackageCcDeclaresResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListApplicationPackageCcDeclares(request *ListApplicationPackageCcDeclaresRequest) (response *ListApplicationPackageCcDeclaresResponse, err error) {
	if request == nil {
		request = NewListApplicationPackageCcDeclaresRequest()
	}
	response = NewListApplicationPackageCcDeclaresResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductOperationRequest() (request *GetProductOperationRequest) {
	request = &GetProductOperationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProductOperation")
	return
}

func NewGetProductOperationResponse() (response *GetProductOperationResponse) {
	response = &GetProductOperationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProductOperation(request *GetProductOperationRequest) (response *GetProductOperationResponse, err error) {
	if request == nil {
		request = NewGetProductOperationRequest()
	}
	response = NewGetProductOperationResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateCommonOperationSheetsRequest() (request *BulkCreateCommonOperationSheetsRequest) {
	request = &BulkCreateCommonOperationSheetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateCommonOperationSheets")
	return
}

func NewBulkCreateCommonOperationSheetsResponse() (response *BulkCreateCommonOperationSheetsResponse) {
	response = &BulkCreateCommonOperationSheetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateCommonOperationSheets(request *BulkCreateCommonOperationSheetsRequest) (response *BulkCreateCommonOperationSheetsResponse, err error) {
	if request == nil {
		request = NewBulkCreateCommonOperationSheetsRequest()
	}
	response = NewBulkCreateCommonOperationSheetsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateProductOperationsRequest() (request *BulkCreateProductOperationsRequest) {
	request = &BulkCreateProductOperationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateProductOperations")
	return
}

func NewBulkCreateProductOperationsResponse() (response *BulkCreateProductOperationsResponse) {
	response = &BulkCreateProductOperationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateProductOperations(request *BulkCreateProductOperationsRequest) (response *BulkCreateProductOperationsResponse, err error) {
	if request == nil {
		request = NewBulkCreateProductOperationsRequest()
	}
	response = NewBulkCreateProductOperationsResponse()
	err = c.Send(request, response)
	return
}

func NewListProjectRegionsRequest() (request *ListProjectRegionsRequest) {
	request = &ListProjectRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProjectRegions")
	return
}

func NewListProjectRegionsResponse() (response *ListProjectRegionsResponse) {
	response = &ListProjectRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProjectRegions(request *ListProjectRegionsRequest) (response *ListProjectRegionsResponse, err error) {
	if request == nil {
		request = NewListProjectRegionsRequest()
	}
	response = NewListProjectRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateProjectZonesRequest() (request *BulkCreateProjectZonesRequest) {
	request = &BulkCreateProjectZonesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateProjectZones")
	return
}

func NewBulkCreateProjectZonesResponse() (response *BulkCreateProjectZonesResponse) {
	response = &BulkCreateProjectZonesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateProjectZones(request *BulkCreateProjectZonesRequest) (response *BulkCreateProjectZonesResponse, err error) {
	if request == nil {
		request = NewBulkCreateProjectZonesRequest()
	}
	response = NewBulkCreateProjectZonesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductOperationRequest() (request *DeleteProductOperationRequest) {
	request = &DeleteProductOperationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteProductOperation")
	return
}

func NewDeleteProductOperationResponse() (response *DeleteProductOperationResponse) {
	response = &DeleteProductOperationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteProductOperation(request *DeleteProductOperationRequest) (response *DeleteProductOperationResponse, err error) {
	if request == nil {
		request = NewDeleteProductOperationRequest()
	}
	response = NewDeleteProductOperationResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSiteProductVersionRelRequest() (request *DeleteSiteProductVersionRelRequest) {
	request = &DeleteSiteProductVersionRelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteSiteProductVersionRel")
	return
}

func NewDeleteSiteProductVersionRelResponse() (response *DeleteSiteProductVersionRelResponse) {
	response = &DeleteSiteProductVersionRelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteSiteProductVersionRel(request *DeleteSiteProductVersionRelRequest) (response *DeleteSiteProductVersionRelResponse, err error) {
	if request == nil {
		request = NewDeleteSiteProductVersionRelRequest()
	}
	response = NewDeleteSiteProductVersionRelResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSolutionVersionRequest() (request *DeleteSolutionVersionRequest) {
	request = &DeleteSolutionVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteSolutionVersion")
	return
}

func NewDeleteSolutionVersionResponse() (response *DeleteSolutionVersionResponse) {
	response = &DeleteSolutionVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteSolutionVersion(request *DeleteSolutionVersionRequest) (response *DeleteSolutionVersionResponse, err error) {
	if request == nil {
		request = NewDeleteSolutionVersionRequest()
	}
	response = NewDeleteSolutionVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePlanRequest() (request *DeletePlanRequest) {
	request = &DeletePlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeletePlan")
	return
}

func NewDeletePlanResponse() (response *DeletePlanResponse) {
	response = &DeletePlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeletePlan(request *DeletePlanRequest) (response *DeletePlanResponse, err error) {
	if request == nil {
		request = NewDeletePlanRequest()
	}
	response = NewDeletePlanResponse()
	err = c.Send(request, response)
	return
}

func NewListPlanApplicationRuntimeNamesRequest() (request *ListPlanApplicationRuntimeNamesRequest) {
	request = &ListPlanApplicationRuntimeNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListPlanApplicationRuntimeNames")
	return
}

func NewListPlanApplicationRuntimeNamesResponse() (response *ListPlanApplicationRuntimeNamesResponse) {
	response = &ListPlanApplicationRuntimeNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListPlanApplicationRuntimeNames(request *ListPlanApplicationRuntimeNamesRequest) (response *ListPlanApplicationRuntimeNamesResponse, err error) {
	if request == nil {
		request = NewListPlanApplicationRuntimeNamesRequest()
	}
	response = NewListPlanApplicationRuntimeNamesResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePlanApplicationInstanceRequest() (request *CreatePlanApplicationInstanceRequest) {
	request = &CreatePlanApplicationInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreatePlanApplicationInstance")
	return
}

func NewCreatePlanApplicationInstanceResponse() (response *CreatePlanApplicationInstanceResponse) {
	response = &CreatePlanApplicationInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreatePlanApplicationInstance(request *CreatePlanApplicationInstanceRequest) (response *CreatePlanApplicationInstanceResponse, err error) {
	if request == nil {
		request = NewCreatePlanApplicationInstanceRequest()
	}
	response = NewCreatePlanApplicationInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationInstanceStateRequest() (request *GetApplicationInstanceStateRequest) {
	request = &GetApplicationInstanceStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetApplicationInstanceState")
	return
}

func NewGetApplicationInstanceStateResponse() (response *GetApplicationInstanceStateResponse) {
	response = &GetApplicationInstanceStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetApplicationInstanceState(request *GetApplicationInstanceStateRequest) (response *GetApplicationInstanceStateResponse, err error) {
	if request == nil {
		request = NewGetApplicationInstanceStateRequest()
	}
	response = NewGetApplicationInstanceStateResponse()
	err = c.Send(request, response)
	return
}

func NewGetOperationSheetTemplateRequest() (request *GetOperationSheetTemplateRequest) {
	request = &GetOperationSheetTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetOperationSheetTemplate")
	return
}

func NewGetOperationSheetTemplateResponse() (response *GetOperationSheetTemplateResponse) {
	response = &GetOperationSheetTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetOperationSheetTemplate(request *GetOperationSheetTemplateRequest) (response *GetOperationSheetTemplateResponse, err error) {
	if request == nil {
		request = NewGetOperationSheetTemplateRequest()
	}
	response = NewGetOperationSheetTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductInstanceOperationSheetRequest() (request *GetProductInstanceOperationSheetRequest) {
	request = &GetProductInstanceOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProductInstanceOperationSheet")
	return
}

func NewGetProductInstanceOperationSheetResponse() (response *GetProductInstanceOperationSheetResponse) {
	response = &GetProductInstanceOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProductInstanceOperationSheet(request *GetProductInstanceOperationSheetRequest) (response *GetProductInstanceOperationSheetResponse, err error) {
	if request == nil {
		request = NewGetProductInstanceOperationSheetRequest()
	}
	response = NewGetProductInstanceOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMaterialSyncTaskRequest() (request *UpdateMaterialSyncTaskRequest) {
	request = &UpdateMaterialSyncTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateMaterialSyncTask")
	return
}

func NewUpdateMaterialSyncTaskResponse() (response *UpdateMaterialSyncTaskResponse) {
	response = &UpdateMaterialSyncTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateMaterialSyncTask(request *UpdateMaterialSyncTaskRequest) (response *UpdateMaterialSyncTaskResponse, err error) {
	if request == nil {
		request = NewUpdateMaterialSyncTaskRequest()
	}
	response = NewUpdateMaterialSyncTaskResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePlanTagRequest() (request *UpdatePlanTagRequest) {
	request = &UpdatePlanTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdatePlanTag")
	return
}

func NewUpdatePlanTagResponse() (response *UpdatePlanTagResponse) {
	response = &UpdatePlanTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdatePlanTag(request *UpdatePlanTagRequest) (response *UpdatePlanTagResponse, err error) {
	if request == nil {
		request = NewUpdatePlanTagRequest()
	}
	response = NewUpdatePlanTagResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductOperationRequest() (request *CreateProductOperationRequest) {
	request = &CreateProductOperationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateProductOperation")
	return
}

func NewCreateProductOperationResponse() (response *CreateProductOperationResponse) {
	response = &CreateProductOperationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateProductOperation(request *CreateProductOperationRequest) (response *CreateProductOperationResponse, err error) {
	if request == nil {
		request = NewCreateProductOperationRequest()
	}
	response = NewCreateProductOperationResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApplicationPackageRequest() (request *CreateApplicationPackageRequest) {
	request = &CreateApplicationPackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateApplicationPackage")
	return
}

func NewCreateApplicationPackageResponse() (response *CreateApplicationPackageResponse) {
	response = &CreateApplicationPackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateApplicationPackage(request *CreateApplicationPackageRequest) (response *CreateApplicationPackageResponse, err error) {
	if request == nil {
		request = NewCreateApplicationPackageRequest()
	}
	response = NewCreateApplicationPackageResponse()
	err = c.Send(request, response)
	return
}

func NewGetPlanApplicationInstanceRequest() (request *GetPlanApplicationInstanceRequest) {
	request = &GetPlanApplicationInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetPlanApplicationInstance")
	return
}

func NewGetPlanApplicationInstanceResponse() (response *GetPlanApplicationInstanceResponse) {
	response = &GetPlanApplicationInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetPlanApplicationInstance(request *GetPlanApplicationInstanceRequest) (response *GetPlanApplicationInstanceResponse, err error) {
	if request == nil {
		request = NewGetPlanApplicationInstanceRequest()
	}
	response = NewGetPlanApplicationInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProductDictionaryRequest() (request *UpdateProductDictionaryRequest) {
	request = &UpdateProductDictionaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateProductDictionary")
	return
}

func NewUpdateProductDictionaryResponse() (response *UpdateProductDictionaryResponse) {
	response = &UpdateProductDictionaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateProductDictionary(request *UpdateProductDictionaryRequest) (response *UpdateProductDictionaryResponse, err error) {
	if request == nil {
		request = NewUpdateProductDictionaryRequest()
	}
	response = NewUpdateProductDictionaryResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperationSheetRelationRequest() (request *DeleteOperationSheetRelationRequest) {
	request = &DeleteOperationSheetRelationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteOperationSheetRelation")
	return
}

func NewDeleteOperationSheetRelationResponse() (response *DeleteOperationSheetRelationResponse) {
	response = &DeleteOperationSheetRelationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteOperationSheetRelation(request *DeleteOperationSheetRelationRequest) (response *DeleteOperationSheetRelationResponse, err error) {
	if request == nil {
		request = NewDeleteOperationSheetRelationRequest()
	}
	response = NewDeleteOperationSheetRelationResponse()
	err = c.Send(request, response)
	return
}

func NewListProductInstanceStatesRequest() (request *ListProductInstanceStatesRequest) {
	request = &ListProductInstanceStatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProductInstanceStates")
	return
}

func NewListProductInstanceStatesResponse() (response *ListProductInstanceStatesResponse) {
	response = &ListProductInstanceStatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProductInstanceStates(request *ListProductInstanceStatesRequest) (response *ListProductInstanceStatesResponse, err error) {
	if request == nil {
		request = NewListProductInstanceStatesRequest()
	}
	response = NewListProductInstanceStatesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSiteProductVersionRelRequest() (request *UpdateSiteProductVersionRelRequest) {
	request = &UpdateSiteProductVersionRelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateSiteProductVersionRel")
	return
}

func NewUpdateSiteProductVersionRelResponse() (response *UpdateSiteProductVersionRelResponse) {
	response = &UpdateSiteProductVersionRelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateSiteProductVersionRel(request *UpdateSiteProductVersionRelRequest) (response *UpdateSiteProductVersionRelResponse, err error) {
	if request == nil {
		request = NewUpdateSiteProductVersionRelRequest()
	}
	response = NewUpdateSiteProductVersionRelResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePlanSiteRequest() (request *DeletePlanSiteRequest) {
	request = &DeletePlanSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeletePlanSite")
	return
}

func NewDeletePlanSiteResponse() (response *DeletePlanSiteResponse) {
	response = &DeletePlanSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeletePlanSite(request *DeletePlanSiteRequest) (response *DeletePlanSiteResponse, err error) {
	if request == nil {
		request = NewDeletePlanSiteRequest()
	}
	response = NewDeletePlanSiteResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductDictionaryRequest() (request *GetProductDictionaryRequest) {
	request = &GetProductDictionaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProductDictionary")
	return
}

func NewGetProductDictionaryResponse() (response *GetProductDictionaryResponse) {
	response = &GetProductDictionaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProductDictionary(request *GetProductDictionaryRequest) (response *GetProductDictionaryResponse, err error) {
	if request == nil {
		request = NewGetProductDictionaryRequest()
	}
	response = NewGetProductDictionaryResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateApplicationRunningPlansRequest() (request *BulkCreateApplicationRunningPlansRequest) {
	request = &BulkCreateApplicationRunningPlansRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateApplicationRunningPlans")
	return
}

func NewBulkCreateApplicationRunningPlansResponse() (response *BulkCreateApplicationRunningPlansResponse) {
	response = &BulkCreateApplicationRunningPlansResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateApplicationRunningPlans(request *BulkCreateApplicationRunningPlansRequest) (response *BulkCreateApplicationRunningPlansResponse, err error) {
	if request == nil {
		request = NewBulkCreateApplicationRunningPlansRequest()
	}
	response = NewBulkCreateApplicationRunningPlansResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMaterialMetadataRequest() (request *UpdateMaterialMetadataRequest) {
	request = &UpdateMaterialMetadataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateMaterialMetadata")
	return
}

func NewUpdateMaterialMetadataResponse() (response *UpdateMaterialMetadataResponse) {
	response = &UpdateMaterialMetadataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateMaterialMetadata(request *UpdateMaterialMetadataRequest) (response *UpdateMaterialMetadataResponse, err error) {
	if request == nil {
		request = NewUpdateMaterialMetadataRequest()
	}
	response = NewUpdateMaterialMetadataResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMaterialCustomizedTypeRequest() (request *CreateMaterialCustomizedTypeRequest) {
	request = &CreateMaterialCustomizedTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateMaterialCustomizedType")
	return
}

func NewCreateMaterialCustomizedTypeResponse() (response *CreateMaterialCustomizedTypeResponse) {
	response = &CreateMaterialCustomizedTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateMaterialCustomizedType(request *CreateMaterialCustomizedTypeRequest) (response *CreateMaterialCustomizedTypeResponse, err error) {
	if request == nil {
		request = NewCreateMaterialCustomizedTypeRequest()
	}
	response = NewCreateMaterialCustomizedTypeResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateSolutionVersionsRequest() (request *BulkCreateSolutionVersionsRequest) {
	request = &BulkCreateSolutionVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateSolutionVersions")
	return
}

func NewBulkCreateSolutionVersionsResponse() (response *BulkCreateSolutionVersionsResponse) {
	response = &BulkCreateSolutionVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateSolutionVersions(request *BulkCreateSolutionVersionsRequest) (response *BulkCreateSolutionVersionsResponse, err error) {
	if request == nil {
		request = NewBulkCreateSolutionVersionsRequest()
	}
	response = NewBulkCreateSolutionVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateFlywaySchemaHistoriesRequest() (request *BulkCreateFlywaySchemaHistoriesRequest) {
	request = &BulkCreateFlywaySchemaHistoriesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateFlywaySchemaHistories")
	return
}

func NewBulkCreateFlywaySchemaHistoriesResponse() (response *BulkCreateFlywaySchemaHistoriesResponse) {
	response = &BulkCreateFlywaySchemaHistoriesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateFlywaySchemaHistories(request *BulkCreateFlywaySchemaHistoriesRequest) (response *BulkCreateFlywaySchemaHistoriesResponse, err error) {
	if request == nil {
		request = NewBulkCreateFlywaySchemaHistoriesRequest()
	}
	response = NewBulkCreateFlywaySchemaHistoriesResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateOperationSheetApprovalRecordsRequest() (request *BulkCreateOperationSheetApprovalRecordsRequest) {
	request = &BulkCreateOperationSheetApprovalRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateOperationSheetApprovalRecords")
	return
}

func NewBulkCreateOperationSheetApprovalRecordsResponse() (response *BulkCreateOperationSheetApprovalRecordsResponse) {
	response = &BulkCreateOperationSheetApprovalRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateOperationSheetApprovalRecords(request *BulkCreateOperationSheetApprovalRecordsRequest) (response *BulkCreateOperationSheetApprovalRecordsResponse, err error) {
	if request == nil {
		request = NewBulkCreateOperationSheetApprovalRecordsRequest()
	}
	response = NewBulkCreateOperationSheetApprovalRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePlanCustomConfigRequest() (request *CreatePlanCustomConfigRequest) {
	request = &CreatePlanCustomConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreatePlanCustomConfig")
	return
}

func NewCreatePlanCustomConfigResponse() (response *CreatePlanCustomConfigResponse) {
	response = &CreatePlanCustomConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreatePlanCustomConfig(request *CreatePlanCustomConfigRequest) (response *CreatePlanCustomConfigResponse, err error) {
	if request == nil {
		request = NewCreatePlanCustomConfigRequest()
	}
	response = NewCreatePlanCustomConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductVersionApplicationPackageRelRequest() (request *GetProductVersionApplicationPackageRelRequest) {
	request = &GetProductVersionApplicationPackageRelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProductVersionApplicationPackageRel")
	return
}

func NewGetProductVersionApplicationPackageRelResponse() (response *GetProductVersionApplicationPackageRelResponse) {
	response = &GetProductVersionApplicationPackageRelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProductVersionApplicationPackageRel(request *GetProductVersionApplicationPackageRelRequest) (response *GetProductVersionApplicationPackageRelResponse, err error) {
	if request == nil {
		request = NewGetProductVersionApplicationPackageRelRequest()
	}
	response = NewGetProductVersionApplicationPackageRelResponse()
	err = c.Send(request, response)
	return
}

func NewGetOperationSheetRelationRequest() (request *GetOperationSheetRelationRequest) {
	request = &GetOperationSheetRelationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetOperationSheetRelation")
	return
}

func NewGetOperationSheetRelationResponse() (response *GetOperationSheetRelationResponse) {
	response = &GetOperationSheetRelationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetOperationSheetRelation(request *GetOperationSheetRelationRequest) (response *GetOperationSheetRelationResponse, err error) {
	if request == nil {
		request = NewGetOperationSheetRelationRequest()
	}
	response = NewGetOperationSheetRelationResponse()
	err = c.Send(request, response)
	return
}

func NewListApplicationArtifactChartHistoriesRequest() (request *ListApplicationArtifactChartHistoriesRequest) {
	request = &ListApplicationArtifactChartHistoriesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListApplicationArtifactChartHistories")
	return
}

func NewListApplicationArtifactChartHistoriesResponse() (response *ListApplicationArtifactChartHistoriesResponse) {
	response = &ListApplicationArtifactChartHistoriesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListApplicationArtifactChartHistories(request *ListApplicationArtifactChartHistoriesRequest) (response *ListApplicationArtifactChartHistoriesResponse, err error) {
	if request == nil {
		request = NewListApplicationArtifactChartHistoriesRequest()
	}
	response = NewListApplicationArtifactChartHistoriesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTenantRequest() (request *DeleteTenantRequest) {
	request = &DeleteTenantRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteTenant")
	return
}

func NewDeleteTenantResponse() (response *DeleteTenantResponse) {
	response = &DeleteTenantResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteTenant(request *DeleteTenantRequest) (response *DeleteTenantResponse, err error) {
	if request == nil {
		request = NewDeleteTenantRequest()
	}
	response = NewDeleteTenantResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProjectSiteInfoRequest() (request *UpdateProjectSiteInfoRequest) {
	request = &UpdateProjectSiteInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateProjectSiteInfo")
	return
}

func NewUpdateProjectSiteInfoResponse() (response *UpdateProjectSiteInfoResponse) {
	response = &UpdateProjectSiteInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateProjectSiteInfo(request *UpdateProjectSiteInfoRequest) (response *UpdateProjectSiteInfoResponse, err error) {
	if request == nil {
		request = NewUpdateProjectSiteInfoRequest()
	}
	response = NewUpdateProjectSiteInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMaterialSyncTaskRequest() (request *CreateMaterialSyncTaskRequest) {
	request = &CreateMaterialSyncTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateMaterialSyncTask")
	return
}

func NewCreateMaterialSyncTaskResponse() (response *CreateMaterialSyncTaskResponse) {
	response = &CreateMaterialSyncTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateMaterialSyncTask(request *CreateMaterialSyncTaskRequest) (response *CreateMaterialSyncTaskResponse, err error) {
	if request == nil {
		request = NewCreateMaterialSyncTaskRequest()
	}
	response = NewCreateMaterialSyncTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAtomicRequest() (request *DeleteAtomicRequest) {
	request = &DeleteAtomicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteAtomic")
	return
}

func NewDeleteAtomicResponse() (response *DeleteAtomicResponse) {
	response = &DeleteAtomicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteAtomic(request *DeleteAtomicRequest) (response *DeleteAtomicResponse, err error) {
	if request == nil {
		request = NewDeleteAtomicRequest()
	}
	response = NewDeleteAtomicResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSolutionTemplateRequest() (request *DeleteSolutionTemplateRequest) {
	request = &DeleteSolutionTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteSolutionTemplate")
	return
}

func NewDeleteSolutionTemplateResponse() (response *DeleteSolutionTemplateResponse) {
	response = &DeleteSolutionTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteSolutionTemplate(request *DeleteSolutionTemplateRequest) (response *DeleteSolutionTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteSolutionTemplateRequest()
	}
	response = NewDeleteSolutionTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperationSheetTemplateRequest() (request *CreateOperationSheetTemplateRequest) {
	request = &CreateOperationSheetTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateOperationSheetTemplate")
	return
}

func NewCreateOperationSheetTemplateResponse() (response *CreateOperationSheetTemplateResponse) {
	response = &CreateOperationSheetTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateOperationSheetTemplate(request *CreateOperationSheetTemplateRequest) (response *CreateOperationSheetTemplateResponse, err error) {
	if request == nil {
		request = NewCreateOperationSheetTemplateRequest()
	}
	response = NewCreateOperationSheetTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductVersionApplicationPackageRelRequest() (request *DeleteProductVersionApplicationPackageRelRequest) {
	request = &DeleteProductVersionApplicationPackageRelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteProductVersionApplicationPackageRel")
	return
}

func NewDeleteProductVersionApplicationPackageRelResponse() (response *DeleteProductVersionApplicationPackageRelResponse) {
	response = &DeleteProductVersionApplicationPackageRelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteProductVersionApplicationPackageRel(request *DeleteProductVersionApplicationPackageRelRequest) (response *DeleteProductVersionApplicationPackageRelResponse, err error) {
	if request == nil {
		request = NewDeleteProductVersionApplicationPackageRelRequest()
	}
	response = NewDeleteProductVersionApplicationPackageRelResponse()
	err = c.Send(request, response)
	return
}

func NewGetOperationSheetRequest() (request *GetOperationSheetRequest) {
	request = &GetOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetOperationSheet")
	return
}

func NewGetOperationSheetResponse() (response *GetOperationSheetResponse) {
	response = &GetOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetOperationSheet(request *GetOperationSheetRequest) (response *GetOperationSheetResponse, err error) {
	if request == nil {
		request = NewGetOperationSheetRequest()
	}
	response = NewGetOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewListProductTemplatesRequest() (request *ListProductTemplatesRequest) {
	request = &ListProductTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProductTemplates")
	return
}

func NewListProductTemplatesResponse() (response *ListProductTemplatesResponse) {
	response = &ListProductTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProductTemplates(request *ListProductTemplatesRequest) (response *ListProductTemplatesResponse, err error) {
	if request == nil {
		request = NewListProductTemplatesRequest()
	}
	response = NewListProductTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateOrchTemplateLabelsRequest() (request *BulkCreateOrchTemplateLabelsRequest) {
	request = &BulkCreateOrchTemplateLabelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateOrchTemplateLabels")
	return
}

func NewBulkCreateOrchTemplateLabelsResponse() (response *BulkCreateOrchTemplateLabelsResponse) {
	response = &BulkCreateOrchTemplateLabelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateOrchTemplateLabels(request *BulkCreateOrchTemplateLabelsRequest) (response *BulkCreateOrchTemplateLabelsResponse, err error) {
	if request == nil {
		request = NewBulkCreateOrchTemplateLabelsRequest()
	}
	response = NewBulkCreateOrchTemplateLabelsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateSiteOperationSheetsRequest() (request *BulkCreateSiteOperationSheetsRequest) {
	request = &BulkCreateSiteOperationSheetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateSiteOperationSheets")
	return
}

func NewBulkCreateSiteOperationSheetsResponse() (response *BulkCreateSiteOperationSheetsResponse) {
	response = &BulkCreateSiteOperationSheetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateSiteOperationSheets(request *BulkCreateSiteOperationSheetsRequest) (response *BulkCreateSiteOperationSheetsResponse, err error) {
	if request == nil {
		request = NewBulkCreateSiteOperationSheetsRequest()
	}
	response = NewBulkCreateSiteOperationSheetsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApplicationPackageCcDeclareRequest() (request *DeleteApplicationPackageCcDeclareRequest) {
	request = &DeleteApplicationPackageCcDeclareRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteApplicationPackageCcDeclare")
	return
}

func NewDeleteApplicationPackageCcDeclareResponse() (response *DeleteApplicationPackageCcDeclareResponse) {
	response = &DeleteApplicationPackageCcDeclareResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteApplicationPackageCcDeclare(request *DeleteApplicationPackageCcDeclareRequest) (response *DeleteApplicationPackageCcDeclareResponse, err error) {
	if request == nil {
		request = NewDeleteApplicationPackageCcDeclareRequest()
	}
	response = NewDeleteApplicationPackageCcDeclareResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProjectZoneRequest() (request *UpdateProjectZoneRequest) {
	request = &UpdateProjectZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateProjectZone")
	return
}

func NewUpdateProjectZoneResponse() (response *UpdateProjectZoneResponse) {
	response = &UpdateProjectZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateProjectZone(request *UpdateProjectZoneRequest) (response *UpdateProjectZoneResponse, err error) {
	if request == nil {
		request = NewUpdateProjectZoneRequest()
	}
	response = NewUpdateProjectZoneResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOperationBuiltinDataSchemaRequest() (request *UpdateOperationBuiltinDataSchemaRequest) {
	request = &UpdateOperationBuiltinDataSchemaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateOperationBuiltinDataSchema")
	return
}

func NewUpdateOperationBuiltinDataSchemaResponse() (response *UpdateOperationBuiltinDataSchemaResponse) {
	response = &UpdateOperationBuiltinDataSchemaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateOperationBuiltinDataSchema(request *UpdateOperationBuiltinDataSchemaRequest) (response *UpdateOperationBuiltinDataSchemaResponse, err error) {
	if request == nil {
		request = NewUpdateOperationBuiltinDataSchemaRequest()
	}
	response = NewUpdateOperationBuiltinDataSchemaResponse()
	err = c.Send(request, response)
	return
}

func NewGetOperationSheetApprovalRecordRequest() (request *GetOperationSheetApprovalRecordRequest) {
	request = &GetOperationSheetApprovalRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetOperationSheetApprovalRecord")
	return
}

func NewGetOperationSheetApprovalRecordResponse() (response *GetOperationSheetApprovalRecordResponse) {
	response = &GetOperationSheetApprovalRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetOperationSheetApprovalRecord(request *GetOperationSheetApprovalRecordRequest) (response *GetOperationSheetApprovalRecordResponse, err error) {
	if request == nil {
		request = NewGetOperationSheetApprovalRecordRequest()
	}
	response = NewGetOperationSheetApprovalRecordResponse()
	err = c.Send(request, response)
	return
}

func NewListProductVersionApplicationPackageRelsRequest() (request *ListProductVersionApplicationPackageRelsRequest) {
	request = &ListProductVersionApplicationPackageRelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProductVersionApplicationPackageRels")
	return
}

func NewListProductVersionApplicationPackageRelsResponse() (response *ListProductVersionApplicationPackageRelsResponse) {
	response = &ListProductVersionApplicationPackageRelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProductVersionApplicationPackageRels(request *ListProductVersionApplicationPackageRelsRequest) (response *ListProductVersionApplicationPackageRelsResponse, err error) {
	if request == nil {
		request = NewListProductVersionApplicationPackageRelsRequest()
	}
	response = NewListProductVersionApplicationPackageRelsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductInstanceOperationSheetRequest() (request *CreateProductInstanceOperationSheetRequest) {
	request = &CreateProductInstanceOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateProductInstanceOperationSheet")
	return
}

func NewCreateProductInstanceOperationSheetResponse() (response *CreateProductInstanceOperationSheetResponse) {
	response = &CreateProductInstanceOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateProductInstanceOperationSheet(request *CreateProductInstanceOperationSheetRequest) (response *CreateProductInstanceOperationSheetResponse, err error) {
	if request == nil {
		request = NewCreateProductInstanceOperationSheetRequest()
	}
	response = NewCreateProductInstanceOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewGetOperationPipelineRequest() (request *GetOperationPipelineRequest) {
	request = &GetOperationPipelineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetOperationPipeline")
	return
}

func NewGetOperationPipelineResponse() (response *GetOperationPipelineResponse) {
	response = &GetOperationPipelineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetOperationPipeline(request *GetOperationPipelineRequest) (response *GetOperationPipelineResponse, err error) {
	if request == nil {
		request = NewGetOperationPipelineRequest()
	}
	response = NewGetOperationPipelineResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperationSheetAttachmentRequest() (request *DeleteOperationSheetAttachmentRequest) {
	request = &DeleteOperationSheetAttachmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteOperationSheetAttachment")
	return
}

func NewDeleteOperationSheetAttachmentResponse() (response *DeleteOperationSheetAttachmentResponse) {
	response = &DeleteOperationSheetAttachmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteOperationSheetAttachment(request *DeleteOperationSheetAttachmentRequest) (response *DeleteOperationSheetAttachmentResponse, err error) {
	if request == nil {
		request = NewDeleteOperationSheetAttachmentRequest()
	}
	response = NewDeleteOperationSheetAttachmentResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMaterialApplicationImportDetailRequest() (request *UpdateMaterialApplicationImportDetailRequest) {
	request = &UpdateMaterialApplicationImportDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateMaterialApplicationImportDetail")
	return
}

func NewUpdateMaterialApplicationImportDetailResponse() (response *UpdateMaterialApplicationImportDetailResponse) {
	response = &UpdateMaterialApplicationImportDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateMaterialApplicationImportDetail(request *UpdateMaterialApplicationImportDetailRequest) (response *UpdateMaterialApplicationImportDetailResponse, err error) {
	if request == nil {
		request = NewUpdateMaterialApplicationImportDetailRequest()
	}
	response = NewUpdateMaterialApplicationImportDetailResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateCommonOperationSheetsWithDetailRequest() (request *BulkCreateCommonOperationSheetsWithDetailRequest) {
	request = &BulkCreateCommonOperationSheetsWithDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateCommonOperationSheetsWithDetail")
	return
}

func NewBulkCreateCommonOperationSheetsWithDetailResponse() (response *BulkCreateCommonOperationSheetsWithDetailResponse) {
	response = &BulkCreateCommonOperationSheetsWithDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateCommonOperationSheetsWithDetail(request *BulkCreateCommonOperationSheetsWithDetailRequest) (response *BulkCreateCommonOperationSheetsWithDetailResponse, err error) {
	if request == nil {
		request = NewBulkCreateCommonOperationSheetsWithDetailRequest()
	}
	response = NewBulkCreateCommonOperationSheetsWithDetailResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateMaterialSitesRequest() (request *BulkCreateMaterialSitesRequest) {
	request = &BulkCreateMaterialSitesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateMaterialSites")
	return
}

func NewBulkCreateMaterialSitesResponse() (response *BulkCreateMaterialSitesResponse) {
	response = &BulkCreateMaterialSitesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateMaterialSites(request *BulkCreateMaterialSitesRequest) (response *BulkCreateMaterialSitesResponse, err error) {
	if request == nil {
		request = NewBulkCreateMaterialSitesRequest()
	}
	response = NewBulkCreateMaterialSitesResponse()
	err = c.Send(request, response)
	return
}

func NewGetOperationSheetAttachmentRequest() (request *GetOperationSheetAttachmentRequest) {
	request = &GetOperationSheetAttachmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetOperationSheetAttachment")
	return
}

func NewGetOperationSheetAttachmentResponse() (response *GetOperationSheetAttachmentResponse) {
	response = &GetOperationSheetAttachmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetOperationSheetAttachment(request *GetOperationSheetAttachmentRequest) (response *GetOperationSheetAttachmentResponse, err error) {
	if request == nil {
		request = NewGetOperationSheetAttachmentRequest()
	}
	response = NewGetOperationSheetAttachmentResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApplicationBranchRequest() (request *CreateApplicationBranchRequest) {
	request = &CreateApplicationBranchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateApplicationBranch")
	return
}

func NewCreateApplicationBranchResponse() (response *CreateApplicationBranchResponse) {
	response = &CreateApplicationBranchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateApplicationBranch(request *CreateApplicationBranchRequest) (response *CreateApplicationBranchResponse, err error) {
	if request == nil {
		request = NewCreateApplicationBranchRequest()
	}
	response = NewCreateApplicationBranchResponse()
	err = c.Send(request, response)
	return
}

func NewListApplicationCollectStatusesRequest() (request *ListApplicationCollectStatusesRequest) {
	request = &ListApplicationCollectStatusesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListApplicationCollectStatuses")
	return
}

func NewListApplicationCollectStatusesResponse() (response *ListApplicationCollectStatusesResponse) {
	response = &ListApplicationCollectStatusesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListApplicationCollectStatuses(request *ListApplicationCollectStatusesRequest) (response *ListApplicationCollectStatusesResponse, err error) {
	if request == nil {
		request = NewListApplicationCollectStatusesRequest()
	}
	response = NewListApplicationCollectStatusesResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePlanHistoryRequest() (request *DeletePlanHistoryRequest) {
	request = &DeletePlanHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeletePlanHistory")
	return
}

func NewDeletePlanHistoryResponse() (response *DeletePlanHistoryResponse) {
	response = &DeletePlanHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeletePlanHistory(request *DeletePlanHistoryRequest) (response *DeletePlanHistoryResponse, err error) {
	if request == nil {
		request = NewDeletePlanHistoryRequest()
	}
	response = NewDeletePlanHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTenantRequest() (request *CreateTenantRequest) {
	request = &CreateTenantRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateTenant")
	return
}

func NewCreateTenantResponse() (response *CreateTenantResponse) {
	response = &CreateTenantResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateTenant(request *CreateTenantRequest) (response *CreateTenantResponse, err error) {
	if request == nil {
		request = NewCreateTenantRequest()
	}
	response = NewCreateTenantResponse()
	err = c.Send(request, response)
	return
}

func NewListProjectZonesRequest() (request *ListProjectZonesRequest) {
	request = &ListProjectZonesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProjectZones")
	return
}

func NewListProjectZonesResponse() (response *ListProjectZonesResponse) {
	response = &ListProjectZonesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProjectZones(request *ListProjectZonesRequest) (response *ListProjectZonesResponse, err error) {
	if request == nil {
		request = NewListProjectZonesRequest()
	}
	response = NewListProjectZonesResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductInstanceStateRequest() (request *GetProductInstanceStateRequest) {
	request = &GetProductInstanceStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProductInstanceState")
	return
}

func NewGetProductInstanceStateResponse() (response *GetProductInstanceStateResponse) {
	response = &GetProductInstanceStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProductInstanceState(request *GetProductInstanceStateRequest) (response *GetProductInstanceStateResponse, err error) {
	if request == nil {
		request = NewGetProductInstanceStateRequest()
	}
	response = NewGetProductInstanceStateResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateApplicationPackageCcDeclaresRequest() (request *BulkCreateApplicationPackageCcDeclaresRequest) {
	request = &BulkCreateApplicationPackageCcDeclaresRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateApplicationPackageCcDeclares")
	return
}

func NewBulkCreateApplicationPackageCcDeclaresResponse() (response *BulkCreateApplicationPackageCcDeclaresResponse) {
	response = &BulkCreateApplicationPackageCcDeclaresResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateApplicationPackageCcDeclares(request *BulkCreateApplicationPackageCcDeclaresRequest) (response *BulkCreateApplicationPackageCcDeclaresResponse, err error) {
	if request == nil {
		request = NewBulkCreateApplicationPackageCcDeclaresRequest()
	}
	response = NewBulkCreateApplicationPackageCcDeclaresResponse()
	err = c.Send(request, response)
	return
}

func NewListApplicationInstanceDeployVersionsRequest() (request *ListApplicationInstanceDeployVersionsRequest) {
	request = &ListApplicationInstanceDeployVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListApplicationInstanceDeployVersions")
	return
}

func NewListApplicationInstanceDeployVersionsResponse() (response *ListApplicationInstanceDeployVersionsResponse) {
	response = &ListApplicationInstanceDeployVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListApplicationInstanceDeployVersions(request *ListApplicationInstanceDeployVersionsRequest) (response *ListApplicationInstanceDeployVersionsResponse, err error) {
	if request == nil {
		request = NewListApplicationInstanceDeployVersionsRequest()
	}
	response = NewListApplicationInstanceDeployVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApplicationInstanceStateRequest() (request *UpdateApplicationInstanceStateRequest) {
	request = &UpdateApplicationInstanceStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateApplicationInstanceState")
	return
}

func NewUpdateApplicationInstanceStateResponse() (response *UpdateApplicationInstanceStateResponse) {
	response = &UpdateApplicationInstanceStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateApplicationInstanceState(request *UpdateApplicationInstanceStateRequest) (response *UpdateApplicationInstanceStateResponse, err error) {
	if request == nil {
		request = NewUpdateApplicationInstanceStateRequest()
	}
	response = NewUpdateApplicationInstanceStateResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateProductVersionApplicationPackageRelsRequest() (request *BulkCreateProductVersionApplicationPackageRelsRequest) {
	request = &BulkCreateProductVersionApplicationPackageRelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateProductVersionApplicationPackageRels")
	return
}

func NewBulkCreateProductVersionApplicationPackageRelsResponse() (response *BulkCreateProductVersionApplicationPackageRelsResponse) {
	response = &BulkCreateProductVersionApplicationPackageRelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateProductVersionApplicationPackageRels(request *BulkCreateProductVersionApplicationPackageRelsRequest) (response *BulkCreateProductVersionApplicationPackageRelsResponse, err error) {
	if request == nil {
		request = NewBulkCreateProductVersionApplicationPackageRelsRequest()
	}
	response = NewBulkCreateProductVersionApplicationPackageRelsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateOperationSheetsRequest() (request *BulkCreateOperationSheetsRequest) {
	request = &BulkCreateOperationSheetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateOperationSheets")
	return
}

func NewBulkCreateOperationSheetsResponse() (response *BulkCreateOperationSheetsResponse) {
	response = &BulkCreateOperationSheetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateOperationSheets(request *BulkCreateOperationSheetsRequest) (response *BulkCreateOperationSheetsResponse, err error) {
	if request == nil {
		request = NewBulkCreateOperationSheetsRequest()
	}
	response = NewBulkCreateOperationSheetsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSystemSettingRequest() (request *CreateSystemSettingRequest) {
	request = &CreateSystemSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateSystemSetting")
	return
}

func NewCreateSystemSettingResponse() (response *CreateSystemSettingResponse) {
	response = &CreateSystemSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateSystemSetting(request *CreateSystemSettingRequest) (response *CreateSystemSettingResponse, err error) {
	if request == nil {
		request = NewCreateSystemSettingRequest()
	}
	response = NewCreateSystemSettingResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMaterialApplicationImportDetailRequest() (request *CreateMaterialApplicationImportDetailRequest) {
	request = &CreateMaterialApplicationImportDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateMaterialApplicationImportDetail")
	return
}

func NewCreateMaterialApplicationImportDetailResponse() (response *CreateMaterialApplicationImportDetailResponse) {
	response = &CreateMaterialApplicationImportDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateMaterialApplicationImportDetail(request *CreateMaterialApplicationImportDetailRequest) (response *CreateMaterialApplicationImportDetailResponse, err error) {
	if request == nil {
		request = NewCreateMaterialApplicationImportDetailRequest()
	}
	response = NewCreateMaterialApplicationImportDetailResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationPackageDetailRequest() (request *GetApplicationPackageDetailRequest) {
	request = &GetApplicationPackageDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetApplicationPackageDetail")
	return
}

func NewGetApplicationPackageDetailResponse() (response *GetApplicationPackageDetailResponse) {
	response = &GetApplicationPackageDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetApplicationPackageDetail(request *GetApplicationPackageDetailRequest) (response *GetApplicationPackageDetailResponse, err error) {
	if request == nil {
		request = NewGetApplicationPackageDetailRequest()
	}
	response = NewGetApplicationPackageDetailResponse()
	err = c.Send(request, response)
	return
}

func NewGetPlanGlobalConfigRequest() (request *GetPlanGlobalConfigRequest) {
	request = &GetPlanGlobalConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetPlanGlobalConfig")
	return
}

func NewGetPlanGlobalConfigResponse() (response *GetPlanGlobalConfigResponse) {
	response = &GetPlanGlobalConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetPlanGlobalConfig(request *GetPlanGlobalConfigRequest) (response *GetPlanGlobalConfigResponse, err error) {
	if request == nil {
		request = NewGetPlanGlobalConfigRequest()
	}
	response = NewGetPlanGlobalConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductVersionRequest() (request *GetProductVersionRequest) {
	request = &GetProductVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProductVersion")
	return
}

func NewGetProductVersionResponse() (response *GetProductVersionResponse) {
	response = &GetProductVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProductVersion(request *GetProductVersionRequest) (response *GetProductVersionResponse, err error) {
	if request == nil {
		request = NewGetProductVersionRequest()
	}
	response = NewGetProductVersionResponse()
	err = c.Send(request, response)
	return
}

func NewListPlanServerInfosRequest() (request *ListPlanServerInfosRequest) {
	request = &ListPlanServerInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListPlanServerInfos")
	return
}

func NewListPlanServerInfosResponse() (response *ListPlanServerInfosResponse) {
	response = &ListPlanServerInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListPlanServerInfos(request *ListPlanServerInfosRequest) (response *ListPlanServerInfosResponse, err error) {
	if request == nil {
		request = NewListPlanServerInfosRequest()
	}
	response = NewListPlanServerInfosResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSiteOperationSheetRequest() (request *CreateSiteOperationSheetRequest) {
	request = &CreateSiteOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateSiteOperationSheet")
	return
}

func NewCreateSiteOperationSheetResponse() (response *CreateSiteOperationSheetResponse) {
	response = &CreateSiteOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateSiteOperationSheet(request *CreateSiteOperationSheetRequest) (response *CreateSiteOperationSheetResponse, err error) {
	if request == nil {
		request = NewCreateSiteOperationSheetRequest()
	}
	response = NewCreateSiteOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewGetCommonOperationSheetWithDetailRequest() (request *GetCommonOperationSheetWithDetailRequest) {
	request = &GetCommonOperationSheetWithDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetCommonOperationSheetWithDetail")
	return
}

func NewGetCommonOperationSheetWithDetailResponse() (response *GetCommonOperationSheetWithDetailResponse) {
	response = &GetCommonOperationSheetWithDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetCommonOperationSheetWithDetail(request *GetCommonOperationSheetWithDetailRequest) (response *GetCommonOperationSheetWithDetailResponse, err error) {
	if request == nil {
		request = NewGetCommonOperationSheetWithDetailRequest()
	}
	response = NewGetCommonOperationSheetWithDetailResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateAtomicsRequest() (request *BulkCreateAtomicsRequest) {
	request = &BulkCreateAtomicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateAtomics")
	return
}

func NewBulkCreateAtomicsResponse() (response *BulkCreateAtomicsResponse) {
	response = &BulkCreateAtomicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateAtomics(request *BulkCreateAtomicsRequest) (response *BulkCreateAtomicsResponse, err error) {
	if request == nil {
		request = NewBulkCreateAtomicsRequest()
	}
	response = NewBulkCreateAtomicsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePlanRequest() (request *UpdatePlanRequest) {
	request = &UpdatePlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdatePlan")
	return
}

func NewUpdatePlanResponse() (response *UpdatePlanResponse) {
	response = &UpdatePlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdatePlan(request *UpdatePlanRequest) (response *UpdatePlanResponse, err error) {
	if request == nil {
		request = NewUpdatePlanRequest()
	}
	response = NewUpdatePlanResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSolutionRequest() (request *CreateSolutionRequest) {
	request = &CreateSolutionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateSolution")
	return
}

func NewCreateSolutionResponse() (response *CreateSolutionResponse) {
	response = &CreateSolutionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateSolution(request *CreateSolutionRequest) (response *CreateSolutionResponse, err error) {
	if request == nil {
		request = NewCreateSolutionRequest()
	}
	response = NewCreateSolutionResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePlanApplicationRuntimeNameRequest() (request *DeletePlanApplicationRuntimeNameRequest) {
	request = &DeletePlanApplicationRuntimeNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeletePlanApplicationRuntimeName")
	return
}

func NewDeletePlanApplicationRuntimeNameResponse() (response *DeletePlanApplicationRuntimeNameResponse) {
	response = &DeletePlanApplicationRuntimeNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeletePlanApplicationRuntimeName(request *DeletePlanApplicationRuntimeNameRequest) (response *DeletePlanApplicationRuntimeNameResponse, err error) {
	if request == nil {
		request = NewDeletePlanApplicationRuntimeNameRequest()
	}
	response = NewDeletePlanApplicationRuntimeNameResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductInstanceStateRequest() (request *CreateProductInstanceStateRequest) {
	request = &CreateProductInstanceStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateProductInstanceState")
	return
}

func NewCreateProductInstanceStateResponse() (response *CreateProductInstanceStateResponse) {
	response = &CreateProductInstanceStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateProductInstanceState(request *CreateProductInstanceStateRequest) (response *CreateProductInstanceStateResponse, err error) {
	if request == nil {
		request = NewCreateProductInstanceStateRequest()
	}
	response = NewCreateProductInstanceStateResponse()
	err = c.Send(request, response)
	return
}

func NewListSystemSettingsRequest() (request *ListSystemSettingsRequest) {
	request = &ListSystemSettingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListSystemSettings")
	return
}

func NewListSystemSettingsResponse() (response *ListSystemSettingsResponse) {
	response = &ListSystemSettingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListSystemSettings(request *ListSystemSettingsRequest) (response *ListSystemSettingsResponse, err error) {
	if request == nil {
		request = NewListSystemSettingsRequest()
	}
	response = NewListSystemSettingsResponse()
	err = c.Send(request, response)
	return
}

func NewGetPlanProductInstanceTreeRequest() (request *GetPlanProductInstanceTreeRequest) {
	request = &GetPlanProductInstanceTreeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetPlanProductInstanceTree")
	return
}

func NewGetPlanProductInstanceTreeResponse() (response *GetPlanProductInstanceTreeResponse) {
	response = &GetPlanProductInstanceTreeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetPlanProductInstanceTree(request *GetPlanProductInstanceTreeRequest) (response *GetPlanProductInstanceTreeResponse, err error) {
	if request == nil {
		request = NewGetPlanProductInstanceTreeRequest()
	}
	response = NewGetPlanProductInstanceTreeResponse()
	err = c.Send(request, response)
	return
}

func NewListSolutionTemplatesRequest() (request *ListSolutionTemplatesRequest) {
	request = &ListSolutionTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListSolutionTemplates")
	return
}

func NewListSolutionTemplatesResponse() (response *ListSolutionTemplatesResponse) {
	response = &ListSolutionTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListSolutionTemplates(request *ListSolutionTemplatesRequest) (response *ListSolutionTemplatesResponse, err error) {
	if request == nil {
		request = NewListSolutionTemplatesRequest()
	}
	response = NewListSolutionTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCommonOperationSheetRequest() (request *UpdateCommonOperationSheetRequest) {
	request = &UpdateCommonOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateCommonOperationSheet")
	return
}

func NewUpdateCommonOperationSheetResponse() (response *UpdateCommonOperationSheetResponse) {
	response = &UpdateCommonOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateCommonOperationSheet(request *UpdateCommonOperationSheetRequest) (response *UpdateCommonOperationSheetResponse, err error) {
	if request == nil {
		request = NewUpdateCommonOperationSheetRequest()
	}
	response = NewUpdateCommonOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePlanSubsystemInstanceRequest() (request *CreatePlanSubsystemInstanceRequest) {
	request = &CreatePlanSubsystemInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreatePlanSubsystemInstance")
	return
}

func NewCreatePlanSubsystemInstanceResponse() (response *CreatePlanSubsystemInstanceResponse) {
	response = &CreatePlanSubsystemInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreatePlanSubsystemInstance(request *CreatePlanSubsystemInstanceRequest) (response *CreatePlanSubsystemInstanceResponse, err error) {
	if request == nil {
		request = NewCreatePlanSubsystemInstanceRequest()
	}
	response = NewCreatePlanSubsystemInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewListOperationSheetRelationsRequest() (request *ListOperationSheetRelationsRequest) {
	request = &ListOperationSheetRelationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListOperationSheetRelations")
	return
}

func NewListOperationSheetRelationsResponse() (response *ListOperationSheetRelationsResponse) {
	response = &ListOperationSheetRelationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListOperationSheetRelations(request *ListOperationSheetRelationsRequest) (response *ListOperationSheetRelationsResponse, err error) {
	if request == nil {
		request = NewListOperationSheetRelationsRequest()
	}
	response = NewListOperationSheetRelationsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProjectZoneRequest() (request *DeleteProjectZoneRequest) {
	request = &DeleteProjectZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteProjectZone")
	return
}

func NewDeleteProjectZoneResponse() (response *DeleteProjectZoneResponse) {
	response = &DeleteProjectZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteProjectZone(request *DeleteProjectZoneRequest) (response *DeleteProjectZoneResponse, err error) {
	if request == nil {
		request = NewDeleteProjectZoneRequest()
	}
	response = NewDeleteProjectZoneResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateMaterialSyncTasksRequest() (request *BulkCreateMaterialSyncTasksRequest) {
	request = &BulkCreateMaterialSyncTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateMaterialSyncTasks")
	return
}

func NewBulkCreateMaterialSyncTasksResponse() (response *BulkCreateMaterialSyncTasksResponse) {
	response = &BulkCreateMaterialSyncTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateMaterialSyncTasks(request *BulkCreateMaterialSyncTasksRequest) (response *BulkCreateMaterialSyncTasksResponse, err error) {
	if request == nil {
		request = NewBulkCreateMaterialSyncTasksRequest()
	}
	response = NewBulkCreateMaterialSyncTasksResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateOperationTemplatesRequest() (request *BulkCreateOperationTemplatesRequest) {
	request = &BulkCreateOperationTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateOperationTemplates")
	return
}

func NewBulkCreateOperationTemplatesResponse() (response *BulkCreateOperationTemplatesResponse) {
	response = &BulkCreateOperationTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateOperationTemplates(request *BulkCreateOperationTemplatesRequest) (response *BulkCreateOperationTemplatesResponse, err error) {
	if request == nil {
		request = NewBulkCreateOperationTemplatesRequest()
	}
	response = NewBulkCreateOperationTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewGetPlanHistoryRequest() (request *GetPlanHistoryRequest) {
	request = &GetPlanHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetPlanHistory")
	return
}

func NewGetPlanHistoryResponse() (response *GetPlanHistoryResponse) {
	response = &GetPlanHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetPlanHistory(request *GetPlanHistoryRequest) (response *GetPlanHistoryResponse, err error) {
	if request == nil {
		request = NewGetPlanHistoryRequest()
	}
	response = NewGetPlanHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSiteDeployTaskRequest() (request *DeleteSiteDeployTaskRequest) {
	request = &DeleteSiteDeployTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteSiteDeployTask")
	return
}

func NewDeleteSiteDeployTaskResponse() (response *DeleteSiteDeployTaskResponse) {
	response = &DeleteSiteDeployTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteSiteDeployTask(request *DeleteSiteDeployTaskRequest) (response *DeleteSiteDeployTaskResponse, err error) {
	if request == nil {
		request = NewDeleteSiteDeployTaskRequest()
	}
	response = NewDeleteSiteDeployTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApplicationInstanceBackupRequest() (request *CreateApplicationInstanceBackupRequest) {
	request = &CreateApplicationInstanceBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateApplicationInstanceBackup")
	return
}

func NewCreateApplicationInstanceBackupResponse() (response *CreateApplicationInstanceBackupResponse) {
	response = &CreateApplicationInstanceBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateApplicationInstanceBackup(request *CreateApplicationInstanceBackupRequest) (response *CreateApplicationInstanceBackupResponse, err error) {
	if request == nil {
		request = NewCreateApplicationInstanceBackupRequest()
	}
	response = NewCreateApplicationInstanceBackupResponse()
	err = c.Send(request, response)
	return
}

func NewListOperationSheetsRequest() (request *ListOperationSheetsRequest) {
	request = &ListOperationSheetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListOperationSheets")
	return
}

func NewListOperationSheetsResponse() (response *ListOperationSheetsResponse) {
	response = &ListOperationSheetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListOperationSheets(request *ListOperationSheetsRequest) (response *ListOperationSheetsResponse, err error) {
	if request == nil {
		request = NewListOperationSheetsRequest()
	}
	response = NewListOperationSheetsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMaterialCleanupRequest() (request *CreateMaterialCleanupRequest) {
	request = &CreateMaterialCleanupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateMaterialCleanup")
	return
}

func NewCreateMaterialCleanupResponse() (response *CreateMaterialCleanupResponse) {
	response = &CreateMaterialCleanupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateMaterialCleanup(request *CreateMaterialCleanupRequest) (response *CreateMaterialCleanupResponse, err error) {
	if request == nil {
		request = NewCreateMaterialCleanupRequest()
	}
	response = NewCreateMaterialCleanupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductInstanceStateRequest() (request *DeleteProductInstanceStateRequest) {
	request = &DeleteProductInstanceStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteProductInstanceState")
	return
}

func NewDeleteProductInstanceStateResponse() (response *DeleteProductInstanceStateResponse) {
	response = &DeleteProductInstanceStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteProductInstanceState(request *DeleteProductInstanceStateRequest) (response *DeleteProductInstanceStateResponse, err error) {
	if request == nil {
		request = NewDeleteProductInstanceStateRequest()
	}
	response = NewDeleteProductInstanceStateResponse()
	err = c.Send(request, response)
	return
}

func NewGetOperationBuiltinDataSchemaRequest() (request *GetOperationBuiltinDataSchemaRequest) {
	request = &GetOperationBuiltinDataSchemaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetOperationBuiltinDataSchema")
	return
}

func NewGetOperationBuiltinDataSchemaResponse() (response *GetOperationBuiltinDataSchemaResponse) {
	response = &GetOperationBuiltinDataSchemaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetOperationBuiltinDataSchema(request *GetOperationBuiltinDataSchemaRequest) (response *GetOperationBuiltinDataSchemaResponse, err error) {
	if request == nil {
		request = NewGetOperationBuiltinDataSchemaRequest()
	}
	response = NewGetOperationBuiltinDataSchemaResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreatePlanServerInfosRequest() (request *BulkCreatePlanServerInfosRequest) {
	request = &BulkCreatePlanServerInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreatePlanServerInfos")
	return
}

func NewBulkCreatePlanServerInfosResponse() (response *BulkCreatePlanServerInfosResponse) {
	response = &BulkCreatePlanServerInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreatePlanServerInfos(request *BulkCreatePlanServerInfosRequest) (response *BulkCreatePlanServerInfosResponse, err error) {
	if request == nil {
		request = NewBulkCreatePlanServerInfosRequest()
	}
	response = NewBulkCreatePlanServerInfosResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePlanApplicationServerRelationRequest() (request *CreatePlanApplicationServerRelationRequest) {
	request = &CreatePlanApplicationServerRelationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreatePlanApplicationServerRelation")
	return
}

func NewCreatePlanApplicationServerRelationResponse() (response *CreatePlanApplicationServerRelationResponse) {
	response = &CreatePlanApplicationServerRelationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreatePlanApplicationServerRelation(request *CreatePlanApplicationServerRelationRequest) (response *CreatePlanApplicationServerRelationResponse, err error) {
	if request == nil {
		request = NewCreatePlanApplicationServerRelationRequest()
	}
	response = NewCreatePlanApplicationServerRelationResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductVersionArtifactTagRequest() (request *DeleteProductVersionArtifactTagRequest) {
	request = &DeleteProductVersionArtifactTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteProductVersionArtifactTag")
	return
}

func NewDeleteProductVersionArtifactTagResponse() (response *DeleteProductVersionArtifactTagResponse) {
	response = &DeleteProductVersionArtifactTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteProductVersionArtifactTag(request *DeleteProductVersionArtifactTagRequest) (response *DeleteProductVersionArtifactTagResponse, err error) {
	if request == nil {
		request = NewDeleteProductVersionArtifactTagRequest()
	}
	response = NewDeleteProductVersionArtifactTagResponse()
	err = c.Send(request, response)
	return
}

func NewListSiteOperationSheetsRequest() (request *ListSiteOperationSheetsRequest) {
	request = &ListSiteOperationSheetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListSiteOperationSheets")
	return
}

func NewListSiteOperationSheetsResponse() (response *ListSiteOperationSheetsResponse) {
	response = &ListSiteOperationSheetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListSiteOperationSheets(request *ListSiteOperationSheetsRequest) (response *ListSiteOperationSheetsResponse, err error) {
	if request == nil {
		request = NewListSiteOperationSheetsRequest()
	}
	response = NewListSiteOperationSheetsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSolutionVersionRequest() (request *UpdateSolutionVersionRequest) {
	request = &UpdateSolutionVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateSolutionVersion")
	return
}

func NewUpdateSolutionVersionResponse() (response *UpdateSolutionVersionResponse) {
	response = &UpdateSolutionVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateSolutionVersion(request *UpdateSolutionVersionRequest) (response *UpdateSolutionVersionResponse, err error) {
	if request == nil {
		request = NewUpdateSolutionVersionRequest()
	}
	response = NewUpdateSolutionVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductInfoRequest() (request *DeleteProductInfoRequest) {
	request = &DeleteProductInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteProductInfo")
	return
}

func NewDeleteProductInfoResponse() (response *DeleteProductInfoResponse) {
	response = &DeleteProductInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteProductInfo(request *DeleteProductInfoRequest) (response *DeleteProductInfoResponse, err error) {
	if request == nil {
		request = NewDeleteProductInfoRequest()
	}
	response = NewDeleteProductInfoResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateMaterialApplicationImportDetailsRequest() (request *BulkCreateMaterialApplicationImportDetailsRequest) {
	request = &BulkCreateMaterialApplicationImportDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateMaterialApplicationImportDetails")
	return
}

func NewBulkCreateMaterialApplicationImportDetailsResponse() (response *BulkCreateMaterialApplicationImportDetailsResponse) {
	response = &BulkCreateMaterialApplicationImportDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateMaterialApplicationImportDetails(request *BulkCreateMaterialApplicationImportDetailsRequest) (response *BulkCreateMaterialApplicationImportDetailsResponse, err error) {
	if request == nil {
		request = NewBulkCreateMaterialApplicationImportDetailsRequest()
	}
	response = NewBulkCreateMaterialApplicationImportDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateProductInfosRequest() (request *BulkCreateProductInfosRequest) {
	request = &BulkCreateProductInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateProductInfos")
	return
}

func NewBulkCreateProductInfosResponse() (response *BulkCreateProductInfosResponse) {
	response = &BulkCreateProductInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateProductInfos(request *BulkCreateProductInfosRequest) (response *BulkCreateProductInfosResponse, err error) {
	if request == nil {
		request = NewBulkCreateProductInfosRequest()
	}
	response = NewBulkCreateProductInfosResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApplicationPackageCcDeclareRequest() (request *CreateApplicationPackageCcDeclareRequest) {
	request = &CreateApplicationPackageCcDeclareRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateApplicationPackageCcDeclare")
	return
}

func NewCreateApplicationPackageCcDeclareResponse() (response *CreateApplicationPackageCcDeclareResponse) {
	response = &CreateApplicationPackageCcDeclareResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateApplicationPackageCcDeclare(request *CreateApplicationPackageCcDeclareRequest) (response *CreateApplicationPackageCcDeclareResponse, err error) {
	if request == nil {
		request = NewCreateApplicationPackageCcDeclareRequest()
	}
	response = NewCreateApplicationPackageCcDeclareResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApplicationRequest() (request *DeleteApplicationRequest) {
	request = &DeleteApplicationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteApplication")
	return
}

func NewDeleteApplicationResponse() (response *DeleteApplicationResponse) {
	response = &DeleteApplicationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteApplication(request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
	if request == nil {
		request = NewDeleteApplicationRequest()
	}
	response = NewDeleteApplicationResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperationSheetApprovalRecordRequest() (request *DeleteOperationSheetApprovalRecordRequest) {
	request = &DeleteOperationSheetApprovalRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteOperationSheetApprovalRecord")
	return
}

func NewDeleteOperationSheetApprovalRecordResponse() (response *DeleteOperationSheetApprovalRecordResponse) {
	response = &DeleteOperationSheetApprovalRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteOperationSheetApprovalRecord(request *DeleteOperationSheetApprovalRecordRequest) (response *DeleteOperationSheetApprovalRecordResponse, err error) {
	if request == nil {
		request = NewDeleteOperationSheetApprovalRecordRequest()
	}
	response = NewDeleteOperationSheetApprovalRecordResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationBranchRequest() (request *GetApplicationBranchRequest) {
	request = &GetApplicationBranchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetApplicationBranch")
	return
}

func NewGetApplicationBranchResponse() (response *GetApplicationBranchResponse) {
	response = &GetApplicationBranchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetApplicationBranch(request *GetApplicationBranchRequest) (response *GetApplicationBranchResponse, err error) {
	if request == nil {
		request = NewGetApplicationBranchRequest()
	}
	response = NewGetApplicationBranchResponse()
	err = c.Send(request, response)
	return
}

func NewGetMaterialInfoRequest() (request *GetMaterialInfoRequest) {
	request = &GetMaterialInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetMaterialInfo")
	return
}

func NewGetMaterialInfoResponse() (response *GetMaterialInfoResponse) {
	response = &GetMaterialInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetMaterialInfo(request *GetMaterialInfoRequest) (response *GetMaterialInfoResponse, err error) {
	if request == nil {
		request = NewGetMaterialInfoRequest()
	}
	response = NewGetMaterialInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApplicationRunningPlanRequest() (request *CreateApplicationRunningPlanRequest) {
	request = &CreateApplicationRunningPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateApplicationRunningPlan")
	return
}

func NewCreateApplicationRunningPlanResponse() (response *CreateApplicationRunningPlanResponse) {
	response = &CreateApplicationRunningPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateApplicationRunningPlan(request *CreateApplicationRunningPlanRequest) (response *CreateApplicationRunningPlanResponse, err error) {
	if request == nil {
		request = NewCreateApplicationRunningPlanRequest()
	}
	response = NewCreateApplicationRunningPlanResponse()
	err = c.Send(request, response)
	return
}

func NewGetMaterialImportTaskRequest() (request *GetMaterialImportTaskRequest) {
	request = &GetMaterialImportTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetMaterialImportTask")
	return
}

func NewGetMaterialImportTaskResponse() (response *GetMaterialImportTaskResponse) {
	response = &GetMaterialImportTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetMaterialImportTask(request *GetMaterialImportTaskRequest) (response *GetMaterialImportTaskResponse, err error) {
	if request == nil {
		request = NewGetMaterialImportTaskRequest()
	}
	response = NewGetMaterialImportTaskResponse()
	err = c.Send(request, response)
	return
}

func NewListProductInstanceBackupsRequest() (request *ListProductInstanceBackupsRequest) {
	request = &ListProductInstanceBackupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProductInstanceBackups")
	return
}

func NewListProductInstanceBackupsResponse() (response *ListProductInstanceBackupsResponse) {
	response = &ListProductInstanceBackupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProductInstanceBackups(request *ListProductInstanceBackupsRequest) (response *ListProductInstanceBackupsResponse, err error) {
	if request == nil {
		request = NewListProductInstanceBackupsRequest()
	}
	response = NewListProductInstanceBackupsResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductOperationJobRequest() (request *GetProductOperationJobRequest) {
	request = &GetProductOperationJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProductOperationJob")
	return
}

func NewGetProductOperationJobResponse() (response *GetProductOperationJobResponse) {
	response = &GetProductOperationJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProductOperationJob(request *GetProductOperationJobRequest) (response *GetProductOperationJobResponse, err error) {
	if request == nil {
		request = NewGetProductOperationJobRequest()
	}
	response = NewGetProductOperationJobResponse()
	err = c.Send(request, response)
	return
}

func NewGetSolutionTemplateRequest() (request *GetSolutionTemplateRequest) {
	request = &GetSolutionTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetSolutionTemplate")
	return
}

func NewGetSolutionTemplateResponse() (response *GetSolutionTemplateResponse) {
	response = &GetSolutionTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetSolutionTemplate(request *GetSolutionTemplateRequest) (response *GetSolutionTemplateResponse, err error) {
	if request == nil {
		request = NewGetSolutionTemplateRequest()
	}
	response = NewGetSolutionTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateApplicationInstanceStatesRequest() (request *BulkCreateApplicationInstanceStatesRequest) {
	request = &BulkCreateApplicationInstanceStatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateApplicationInstanceStates")
	return
}

func NewBulkCreateApplicationInstanceStatesResponse() (response *BulkCreateApplicationInstanceStatesResponse) {
	response = &BulkCreateApplicationInstanceStatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateApplicationInstanceStates(request *BulkCreateApplicationInstanceStatesRequest) (response *BulkCreateApplicationInstanceStatesResponse, err error) {
	if request == nil {
		request = NewBulkCreateApplicationInstanceStatesRequest()
	}
	response = NewBulkCreateApplicationInstanceStatesResponse()
	err = c.Send(request, response)
	return
}

func NewGetMaterialCleanupRequest() (request *GetMaterialCleanupRequest) {
	request = &GetMaterialCleanupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetMaterialCleanup")
	return
}

func NewGetMaterialCleanupResponse() (response *GetMaterialCleanupResponse) {
	response = &GetMaterialCleanupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetMaterialCleanup(request *GetMaterialCleanupRequest) (response *GetMaterialCleanupResponse, err error) {
	if request == nil {
		request = NewGetMaterialCleanupRequest()
	}
	response = NewGetMaterialCleanupResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateOperationSheetLabelsRequest() (request *BulkCreateOperationSheetLabelsRequest) {
	request = &BulkCreateOperationSheetLabelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateOperationSheetLabels")
	return
}

func NewBulkCreateOperationSheetLabelsResponse() (response *BulkCreateOperationSheetLabelsResponse) {
	response = &BulkCreateOperationSheetLabelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateOperationSheetLabels(request *BulkCreateOperationSheetLabelsRequest) (response *BulkCreateOperationSheetLabelsResponse, err error) {
	if request == nil {
		request = NewBulkCreateOperationSheetLabelsRequest()
	}
	response = NewBulkCreateOperationSheetLabelsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateProjectTopologiesRequest() (request *BulkCreateProjectTopologiesRequest) {
	request = &BulkCreateProjectTopologiesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateProjectTopologies")
	return
}

func NewBulkCreateProjectTopologiesResponse() (response *BulkCreateProjectTopologiesResponse) {
	response = &BulkCreateProjectTopologiesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateProjectTopologies(request *BulkCreateProjectTopologiesRequest) (response *BulkCreateProjectTopologiesResponse, err error) {
	if request == nil {
		request = NewBulkCreateProjectTopologiesRequest()
	}
	response = NewBulkCreateProjectTopologiesResponse()
	err = c.Send(request, response)
	return
}

func NewListOperationSheetTemplatesRequest() (request *ListOperationSheetTemplatesRequest) {
	request = &ListOperationSheetTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListOperationSheetTemplates")
	return
}

func NewListOperationSheetTemplatesResponse() (response *ListOperationSheetTemplatesResponse) {
	response = &ListOperationSheetTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListOperationSheetTemplates(request *ListOperationSheetTemplatesRequest) (response *ListOperationSheetTemplatesResponse, err error) {
	if request == nil {
		request = NewListOperationSheetTemplatesRequest()
	}
	response = NewListOperationSheetTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewListProductInfosRequest() (request *ListProductInfosRequest) {
	request = &ListProductInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProductInfos")
	return
}

func NewListProductInfosResponse() (response *ListProductInfosResponse) {
	response = &ListProductInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProductInfos(request *ListProductInfosRequest) (response *ListProductInfosResponse, err error) {
	if request == nil {
		request = NewListProductInfosRequest()
	}
	response = NewListProductInfosResponse()
	err = c.Send(request, response)
	return
}

func NewListProductOperationJobsRequest() (request *ListProductOperationJobsRequest) {
	request = &ListProductOperationJobsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProductOperationJobs")
	return
}

func NewListProductOperationJobsResponse() (response *ListProductOperationJobsResponse) {
	response = &ListProductOperationJobsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProductOperationJobs(request *ListProductOperationJobsRequest) (response *ListProductOperationJobsResponse, err error) {
	if request == nil {
		request = NewListProductOperationJobsRequest()
	}
	response = NewListProductOperationJobsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperationBuiltinDataSchemaRequest() (request *DeleteOperationBuiltinDataSchemaRequest) {
	request = &DeleteOperationBuiltinDataSchemaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteOperationBuiltinDataSchema")
	return
}

func NewDeleteOperationBuiltinDataSchemaResponse() (response *DeleteOperationBuiltinDataSchemaResponse) {
	response = &DeleteOperationBuiltinDataSchemaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteOperationBuiltinDataSchema(request *DeleteOperationBuiltinDataSchemaRequest) (response *DeleteOperationBuiltinDataSchemaResponse, err error) {
	if request == nil {
		request = NewDeleteOperationBuiltinDataSchemaRequest()
	}
	response = NewDeleteOperationBuiltinDataSchemaResponse()
	err = c.Send(request, response)
	return
}

func NewGetPlanDetailRequest() (request *GetPlanDetailRequest) {
	request = &GetPlanDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetPlanDetail")
	return
}

func NewGetPlanDetailResponse() (response *GetPlanDetailResponse) {
	response = &GetPlanDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetPlanDetail(request *GetPlanDetailRequest) (response *GetPlanDetailResponse, err error) {
	if request == nil {
		request = NewGetPlanDetailRequest()
	}
	response = NewGetPlanDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApplicationRunningPlanRequest() (request *DeleteApplicationRunningPlanRequest) {
	request = &DeleteApplicationRunningPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteApplicationRunningPlan")
	return
}

func NewDeleteApplicationRunningPlanResponse() (response *DeleteApplicationRunningPlanResponse) {
	response = &DeleteApplicationRunningPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteApplicationRunningPlan(request *DeleteApplicationRunningPlanRequest) (response *DeleteApplicationRunningPlanResponse, err error) {
	if request == nil {
		request = NewDeleteApplicationRunningPlanRequest()
	}
	response = NewDeleteApplicationRunningPlanResponse()
	err = c.Send(request, response)
	return
}

func NewGetOperationSheetLabelRequest() (request *GetOperationSheetLabelRequest) {
	request = &GetOperationSheetLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetOperationSheetLabel")
	return
}

func NewGetOperationSheetLabelResponse() (response *GetOperationSheetLabelResponse) {
	response = &GetOperationSheetLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetOperationSheetLabel(request *GetOperationSheetLabelRequest) (response *GetOperationSheetLabelResponse, err error) {
	if request == nil {
		request = NewGetOperationSheetLabelRequest()
	}
	response = NewGetOperationSheetLabelResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMaterialMetadataRequest() (request *CreateMaterialMetadataRequest) {
	request = &CreateMaterialMetadataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateMaterialMetadata")
	return
}

func NewCreateMaterialMetadataResponse() (response *CreateMaterialMetadataResponse) {
	response = &CreateMaterialMetadataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateMaterialMetadata(request *CreateMaterialMetadataRequest) (response *CreateMaterialMetadataResponse, err error) {
	if request == nil {
		request = NewCreateMaterialMetadataRequest()
	}
	response = NewCreateMaterialMetadataResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSiteOperationSheetRequest() (request *UpdateSiteOperationSheetRequest) {
	request = &UpdateSiteOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateSiteOperationSheet")
	return
}

func NewUpdateSiteOperationSheetResponse() (response *UpdateSiteOperationSheetResponse) {
	response = &UpdateSiteOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateSiteOperationSheet(request *UpdateSiteOperationSheetRequest) (response *UpdateSiteOperationSheetResponse, err error) {
	if request == nil {
		request = NewUpdateSiteOperationSheetRequest()
	}
	response = NewUpdateSiteOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateFlywaySchemaHistoryRequest() (request *UpdateFlywaySchemaHistoryRequest) {
	request = &UpdateFlywaySchemaHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateFlywaySchemaHistory")
	return
}

func NewUpdateFlywaySchemaHistoryResponse() (response *UpdateFlywaySchemaHistoryResponse) {
	response = &UpdateFlywaySchemaHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateFlywaySchemaHistory(request *UpdateFlywaySchemaHistoryRequest) (response *UpdateFlywaySchemaHistoryResponse, err error) {
	if request == nil {
		request = NewUpdateFlywaySchemaHistoryRequest()
	}
	response = NewUpdateFlywaySchemaHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOperationTemplateRequest() (request *UpdateOperationTemplateRequest) {
	request = &UpdateOperationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateOperationTemplate")
	return
}

func NewUpdateOperationTemplateResponse() (response *UpdateOperationTemplateResponse) {
	response = &UpdateOperationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateOperationTemplate(request *UpdateOperationTemplateRequest) (response *UpdateOperationTemplateResponse, err error) {
	if request == nil {
		request = NewUpdateOperationTemplateRequest()
	}
	response = NewUpdateOperationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateDagNodesRequest() (request *BulkCreateDagNodesRequest) {
	request = &BulkCreateDagNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateDagNodes")
	return
}

func NewBulkCreateDagNodesResponse() (response *BulkCreateDagNodesResponse) {
	response = &BulkCreateDagNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateDagNodes(request *BulkCreateDagNodesRequest) (response *BulkCreateDagNodesResponse, err error) {
	if request == nil {
		request = NewBulkCreateDagNodesRequest()
	}
	response = NewBulkCreateDagNodesResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateOperationSheetRelationsRequest() (request *BulkCreateOperationSheetRelationsRequest) {
	request = &BulkCreateOperationSheetRelationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateOperationSheetRelations")
	return
}

func NewBulkCreateOperationSheetRelationsResponse() (response *BulkCreateOperationSheetRelationsResponse) {
	response = &BulkCreateOperationSheetRelationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateOperationSheetRelations(request *BulkCreateOperationSheetRelationsRequest) (response *BulkCreateOperationSheetRelationsResponse, err error) {
	if request == nil {
		request = NewBulkCreateOperationSheetRelationsRequest()
	}
	response = NewBulkCreateOperationSheetRelationsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductOperationJobRequest() (request *CreateProductOperationJobRequest) {
	request = &CreateProductOperationJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateProductOperationJob")
	return
}

func NewCreateProductOperationJobResponse() (response *CreateProductOperationJobResponse) {
	response = &CreateProductOperationJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateProductOperationJob(request *CreateProductOperationJobRequest) (response *CreateProductOperationJobResponse, err error) {
	if request == nil {
		request = NewCreateProductOperationJobRequest()
	}
	response = NewCreateProductOperationJobResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApplicationArtifactChartHistoryRequest() (request *DeleteApplicationArtifactChartHistoryRequest) {
	request = &DeleteApplicationArtifactChartHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteApplicationArtifactChartHistory")
	return
}

func NewDeleteApplicationArtifactChartHistoryResponse() (response *DeleteApplicationArtifactChartHistoryResponse) {
	response = &DeleteApplicationArtifactChartHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteApplicationArtifactChartHistory(request *DeleteApplicationArtifactChartHistoryRequest) (response *DeleteApplicationArtifactChartHistoryResponse, err error) {
	if request == nil {
		request = NewDeleteApplicationArtifactChartHistoryRequest()
	}
	response = NewDeleteApplicationArtifactChartHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewListProjectTopologyDetailsRequest() (request *ListProjectTopologyDetailsRequest) {
	request = &ListProjectTopologyDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProjectTopologyDetails")
	return
}

func NewListProjectTopologyDetailsResponse() (response *ListProjectTopologyDetailsResponse) {
	response = &ListProjectTopologyDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProjectTopologyDetails(request *ListProjectTopologyDetailsRequest) (response *ListProjectTopologyDetailsResponse, err error) {
	if request == nil {
		request = NewListProjectTopologyDetailsRequest()
	}
	response = NewListProjectTopologyDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMaterialMachineSSHInfoRequest() (request *UpdateMaterialMachineSSHInfoRequest) {
	request = &UpdateMaterialMachineSSHInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateMaterialMachineSSHInfo")
	return
}

func NewUpdateMaterialMachineSSHInfoResponse() (response *UpdateMaterialMachineSSHInfoResponse) {
	response = &UpdateMaterialMachineSSHInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateMaterialMachineSSHInfo(request *UpdateMaterialMachineSSHInfoRequest) (response *UpdateMaterialMachineSSHInfoResponse, err error) {
	if request == nil {
		request = NewUpdateMaterialMachineSSHInfoRequest()
	}
	response = NewUpdateMaterialMachineSSHInfoResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePlanSubsystemInstanceRequest() (request *UpdatePlanSubsystemInstanceRequest) {
	request = &UpdatePlanSubsystemInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdatePlanSubsystemInstance")
	return
}

func NewUpdatePlanSubsystemInstanceResponse() (response *UpdatePlanSubsystemInstanceResponse) {
	response = &UpdatePlanSubsystemInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdatePlanSubsystemInstance(request *UpdatePlanSubsystemInstanceRequest) (response *UpdatePlanSubsystemInstanceResponse, err error) {
	if request == nil {
		request = NewUpdatePlanSubsystemInstanceRequest()
	}
	response = NewUpdatePlanSubsystemInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewListApplicationPackageDetailsRequest() (request *ListApplicationPackageDetailsRequest) {
	request = &ListApplicationPackageDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListApplicationPackageDetails")
	return
}

func NewListApplicationPackageDetailsResponse() (response *ListApplicationPackageDetailsResponse) {
	response = &ListApplicationPackageDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListApplicationPackageDetails(request *ListApplicationPackageDetailsRequest) (response *ListApplicationPackageDetailsResponse, err error) {
	if request == nil {
		request = NewListApplicationPackageDetailsRequest()
	}
	response = NewListApplicationPackageDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSystemSettingRequest() (request *DeleteSystemSettingRequest) {
	request = &DeleteSystemSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteSystemSetting")
	return
}

func NewDeleteSystemSettingResponse() (response *DeleteSystemSettingResponse) {
	response = &DeleteSystemSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteSystemSetting(request *DeleteSystemSettingRequest) (response *DeleteSystemSettingResponse, err error) {
	if request == nil {
		request = NewDeleteSystemSettingRequest()
	}
	response = NewDeleteSystemSettingResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateApplicationArtifactChartHistoriesRequest() (request *BulkCreateApplicationArtifactChartHistoriesRequest) {
	request = &BulkCreateApplicationArtifactChartHistoriesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateApplicationArtifactChartHistories")
	return
}

func NewBulkCreateApplicationArtifactChartHistoriesResponse() (response *BulkCreateApplicationArtifactChartHistoriesResponse) {
	response = &BulkCreateApplicationArtifactChartHistoriesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateApplicationArtifactChartHistories(request *BulkCreateApplicationArtifactChartHistoriesRequest) (response *BulkCreateApplicationArtifactChartHistoriesResponse, err error) {
	if request == nil {
		request = NewBulkCreateApplicationArtifactChartHistoriesRequest()
	}
	response = NewBulkCreateApplicationArtifactChartHistoriesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProjectTopologyRequest() (request *CreateProjectTopologyRequest) {
	request = &CreateProjectTopologyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateProjectTopology")
	return
}

func NewCreateProjectTopologyResponse() (response *CreateProjectTopologyResponse) {
	response = &CreateProjectTopologyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateProjectTopology(request *CreateProjectTopologyRequest) (response *CreateProjectTopologyResponse, err error) {
	if request == nil {
		request = NewCreateProjectTopologyRequest()
	}
	response = NewCreateProjectTopologyResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateProductVersionArtifactTagsRequest() (request *BulkCreateProductVersionArtifactTagsRequest) {
	request = &BulkCreateProductVersionArtifactTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateProductVersionArtifactTags")
	return
}

func NewBulkCreateProductVersionArtifactTagsResponse() (response *BulkCreateProductVersionArtifactTagsResponse) {
	response = &BulkCreateProductVersionArtifactTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateProductVersionArtifactTags(request *BulkCreateProductVersionArtifactTagsRequest) (response *BulkCreateProductVersionArtifactTagsResponse, err error) {
	if request == nil {
		request = NewBulkCreateProductVersionArtifactTagsRequest()
	}
	response = NewBulkCreateProductVersionArtifactTagsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSolutionTemplateRequest() (request *CreateSolutionTemplateRequest) {
	request = &CreateSolutionTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateSolutionTemplate")
	return
}

func NewCreateSolutionTemplateResponse() (response *CreateSolutionTemplateResponse) {
	response = &CreateSolutionTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateSolutionTemplate(request *CreateSolutionTemplateRequest) (response *CreateSolutionTemplateResponse, err error) {
	if request == nil {
		request = NewCreateSolutionTemplateRequest()
	}
	response = NewCreateSolutionTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaterialMetadataRequest() (request *DeleteMaterialMetadataRequest) {
	request = &DeleteMaterialMetadataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteMaterialMetadata")
	return
}

func NewDeleteMaterialMetadataResponse() (response *DeleteMaterialMetadataResponse) {
	response = &DeleteMaterialMetadataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteMaterialMetadata(request *DeleteMaterialMetadataRequest) (response *DeleteMaterialMetadataResponse, err error) {
	if request == nil {
		request = NewDeleteMaterialMetadataRequest()
	}
	response = NewDeleteMaterialMetadataResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreatePlansRequest() (request *BulkCreatePlansRequest) {
	request = &BulkCreatePlansRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreatePlans")
	return
}

func NewBulkCreatePlansResponse() (response *BulkCreatePlansResponse) {
	response = &BulkCreatePlansResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreatePlans(request *BulkCreatePlansRequest) (response *BulkCreatePlansResponse, err error) {
	if request == nil {
		request = NewBulkCreatePlansRequest()
	}
	response = NewBulkCreatePlansResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialSiteDetailsRequest() (request *ListMaterialSiteDetailsRequest) {
	request = &ListMaterialSiteDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListMaterialSiteDetails")
	return
}

func NewListMaterialSiteDetailsResponse() (response *ListMaterialSiteDetailsResponse) {
	response = &ListMaterialSiteDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListMaterialSiteDetails(request *ListMaterialSiteDetailsRequest) (response *ListMaterialSiteDetailsResponse, err error) {
	if request == nil {
		request = NewListMaterialSiteDetailsRequest()
	}
	response = NewListMaterialSiteDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperationSheetRequest() (request *DeleteOperationSheetRequest) {
	request = &DeleteOperationSheetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteOperationSheet")
	return
}

func NewDeleteOperationSheetResponse() (response *DeleteOperationSheetResponse) {
	response = &DeleteOperationSheetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteOperationSheet(request *DeleteOperationSheetRequest) (response *DeleteOperationSheetResponse, err error) {
	if request == nil {
		request = NewDeleteOperationSheetRequest()
	}
	response = NewDeleteOperationSheetResponse()
	err = c.Send(request, response)
	return
}

func NewGetPlanApplicationServerRelationRequest() (request *GetPlanApplicationServerRelationRequest) {
	request = &GetPlanApplicationServerRelationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetPlanApplicationServerRelation")
	return
}

func NewGetPlanApplicationServerRelationResponse() (response *GetPlanApplicationServerRelationResponse) {
	response = &GetPlanApplicationServerRelationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetPlanApplicationServerRelation(request *GetPlanApplicationServerRelationRequest) (response *GetPlanApplicationServerRelationResponse, err error) {
	if request == nil {
		request = NewGetPlanApplicationServerRelationRequest()
	}
	response = NewGetPlanApplicationServerRelationResponse()
	err = c.Send(request, response)
	return
}

func NewListAtomicsRequest() (request *ListAtomicsRequest) {
	request = &ListAtomicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListAtomics")
	return
}

func NewListAtomicsResponse() (response *ListAtomicsResponse) {
	response = &ListAtomicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListAtomics(request *ListAtomicsRequest) (response *ListAtomicsResponse, err error) {
	if request == nil {
		request = NewListAtomicsRequest()
	}
	response = NewListAtomicsResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateOperationSheetAttachmentsRequest() (request *BulkCreateOperationSheetAttachmentsRequest) {
	request = &BulkCreateOperationSheetAttachmentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateOperationSheetAttachments")
	return
}

func NewBulkCreateOperationSheetAttachmentsResponse() (response *BulkCreateOperationSheetAttachmentsResponse) {
	response = &BulkCreateOperationSheetAttachmentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateOperationSheetAttachments(request *BulkCreateOperationSheetAttachmentsRequest) (response *BulkCreateOperationSheetAttachmentsResponse, err error) {
	if request == nil {
		request = NewBulkCreateOperationSheetAttachmentsRequest()
	}
	response = NewBulkCreateOperationSheetAttachmentsResponse()
	err = c.Send(request, response)
	return
}

func NewListProjectSiteInfosRequest() (request *ListProjectSiteInfosRequest) {
	request = &ListProjectSiteInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProjectSiteInfos")
	return
}

func NewListProjectSiteInfosResponse() (response *ListProjectSiteInfosResponse) {
	response = &ListProjectSiteInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProjectSiteInfos(request *ListProjectSiteInfosRequest) (response *ListProjectSiteInfosResponse, err error) {
	if request == nil {
		request = NewListProjectSiteInfosRequest()
	}
	response = NewListProjectSiteInfosResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreatePlanHistoriesRequest() (request *BulkCreatePlanHistoriesRequest) {
	request = &BulkCreatePlanHistoriesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreatePlanHistories")
	return
}

func NewBulkCreatePlanHistoriesResponse() (response *BulkCreatePlanHistoriesResponse) {
	response = &BulkCreatePlanHistoriesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreatePlanHistories(request *BulkCreatePlanHistoriesRequest) (response *BulkCreatePlanHistoriesResponse, err error) {
	if request == nil {
		request = NewBulkCreatePlanHistoriesRequest()
	}
	response = NewBulkCreatePlanHistoriesResponse()
	err = c.Send(request, response)
	return
}

func NewListMaterialSitesRequest() (request *ListMaterialSitesRequest) {
	request = &ListMaterialSitesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListMaterialSites")
	return
}

func NewListMaterialSitesResponse() (response *ListMaterialSitesResponse) {
	response = &ListMaterialSitesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListMaterialSites(request *ListMaterialSitesRequest) (response *ListMaterialSitesResponse, err error) {
	if request == nil {
		request = NewListMaterialSitesRequest()
	}
	response = NewListMaterialSitesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApplicationArtifactRequest() (request *UpdateApplicationArtifactRequest) {
	request = &UpdateApplicationArtifactRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateApplicationArtifact")
	return
}

func NewUpdateApplicationArtifactResponse() (response *UpdateApplicationArtifactResponse) {
	response = &UpdateApplicationArtifactResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateApplicationArtifact(request *UpdateApplicationArtifactRequest) (response *UpdateApplicationArtifactResponse, err error) {
	if request == nil {
		request = NewUpdateApplicationArtifactRequest()
	}
	response = NewUpdateApplicationArtifactResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMaterialCommonImportDetailRequest() (request *CreateMaterialCommonImportDetailRequest) {
	request = &CreateMaterialCommonImportDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateMaterialCommonImportDetail")
	return
}

func NewCreateMaterialCommonImportDetailResponse() (response *CreateMaterialCommonImportDetailResponse) {
	response = &CreateMaterialCommonImportDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateMaterialCommonImportDetail(request *CreateMaterialCommonImportDetailRequest) (response *CreateMaterialCommonImportDetailResponse, err error) {
	if request == nil {
		request = NewCreateMaterialCommonImportDetailRequest()
	}
	response = NewCreateMaterialCommonImportDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperationSheetLabelRequest() (request *CreateOperationSheetLabelRequest) {
	request = &CreateOperationSheetLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateOperationSheetLabel")
	return
}

func NewCreateOperationSheetLabelResponse() (response *CreateOperationSheetLabelResponse) {
	response = &CreateOperationSheetLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateOperationSheetLabel(request *CreateOperationSheetLabelRequest) (response *CreateOperationSheetLabelResponse, err error) {
	if request == nil {
		request = NewCreateOperationSheetLabelRequest()
	}
	response = NewCreateOperationSheetLabelResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePlanProductInstanceRequest() (request *UpdatePlanProductInstanceRequest) {
	request = &UpdatePlanProductInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdatePlanProductInstance")
	return
}

func NewUpdatePlanProductInstanceResponse() (response *UpdatePlanProductInstanceResponse) {
	response = &UpdatePlanProductInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdatePlanProductInstance(request *UpdatePlanProductInstanceRequest) (response *UpdatePlanProductInstanceResponse, err error) {
	if request == nil {
		request = NewUpdatePlanProductInstanceRequest()
	}
	response = NewUpdatePlanProductInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewGetSiteDeployTaskRequest() (request *GetSiteDeployTaskRequest) {
	request = &GetSiteDeployTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetSiteDeployTask")
	return
}

func NewGetSiteDeployTaskResponse() (response *GetSiteDeployTaskResponse) {
	response = &GetSiteDeployTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetSiteDeployTask(request *GetSiteDeployTaskRequest) (response *GetSiteDeployTaskResponse, err error) {
	if request == nil {
		request = NewGetSiteDeployTaskRequest()
	}
	response = NewGetSiteDeployTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMaterialCommonImportDetailRequest() (request *DeleteMaterialCommonImportDetailRequest) {
	request = &DeleteMaterialCommonImportDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteMaterialCommonImportDetail")
	return
}

func NewDeleteMaterialCommonImportDetailResponse() (response *DeleteMaterialCommonImportDetailResponse) {
	response = &DeleteMaterialCommonImportDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteMaterialCommonImportDetail(request *DeleteMaterialCommonImportDetailRequest) (response *DeleteMaterialCommonImportDetailResponse, err error) {
	if request == nil {
		request = NewDeleteMaterialCommonImportDetailRequest()
	}
	response = NewDeleteMaterialCommonImportDetailResponse()
	err = c.Send(request, response)
	return
}

func NewGetMaterialApplicationImportDetailRequest() (request *GetMaterialApplicationImportDetailRequest) {
	request = &GetMaterialApplicationImportDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetMaterialApplicationImportDetail")
	return
}

func NewGetMaterialApplicationImportDetailResponse() (response *GetMaterialApplicationImportDetailResponse) {
	response = &GetMaterialApplicationImportDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetMaterialApplicationImportDetail(request *GetMaterialApplicationImportDetailRequest) (response *GetMaterialApplicationImportDetailResponse, err error) {
	if request == nil {
		request = NewGetMaterialApplicationImportDetailRequest()
	}
	response = NewGetMaterialApplicationImportDetailResponse()
	err = c.Send(request, response)
	return
}

func NewGetSiteProductVersionRelRequest() (request *GetSiteProductVersionRelRequest) {
	request = &GetSiteProductVersionRelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetSiteProductVersionRel")
	return
}

func NewGetSiteProductVersionRelResponse() (response *GetSiteProductVersionRelResponse) {
	response = &GetSiteProductVersionRelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetSiteProductVersionRel(request *GetSiteProductVersionRelRequest) (response *GetSiteProductVersionRelResponse, err error) {
	if request == nil {
		request = NewGetSiteProductVersionRelRequest()
	}
	response = NewGetSiteProductVersionRelResponse()
	err = c.Send(request, response)
	return
}

func NewGetMaterialCustomizedTypeRequest() (request *GetMaterialCustomizedTypeRequest) {
	request = &GetMaterialCustomizedTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetMaterialCustomizedType")
	return
}

func NewGetMaterialCustomizedTypeResponse() (response *GetMaterialCustomizedTypeResponse) {
	response = &GetMaterialCustomizedTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetMaterialCustomizedType(request *GetMaterialCustomizedTypeRequest) (response *GetMaterialCustomizedTypeResponse, err error) {
	if request == nil {
		request = NewGetMaterialCustomizedTypeRequest()
	}
	response = NewGetMaterialCustomizedTypeResponse()
	err = c.Send(request, response)
	return
}

func NewListOperationPipelinesRequest() (request *ListOperationPipelinesRequest) {
	request = &ListOperationPipelinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListOperationPipelines")
	return
}

func NewListOperationPipelinesResponse() (response *ListOperationPipelinesResponse) {
	response = &ListOperationPipelinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListOperationPipelines(request *ListOperationPipelinesRequest) (response *ListOperationPipelinesResponse, err error) {
	if request == nil {
		request = NewListOperationPipelinesRequest()
	}
	response = NewListOperationPipelinesResponse()
	err = c.Send(request, response)
	return
}

func NewGetProjectZoneRequest() (request *GetProjectZoneRequest) {
	request = &GetProjectZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetProjectZone")
	return
}

func NewGetProjectZoneResponse() (response *GetProjectZoneResponse) {
	response = &GetProjectZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetProjectZone(request *GetProjectZoneRequest) (response *GetProjectZoneResponse, err error) {
	if request == nil {
		request = NewGetProjectZoneRequest()
	}
	response = NewGetProjectZoneResponse()
	err = c.Send(request, response)
	return
}

func NewListFlywaySchemaHistoriesRequest() (request *ListFlywaySchemaHistoriesRequest) {
	request = &ListFlywaySchemaHistoriesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListFlywaySchemaHistories")
	return
}

func NewListFlywaySchemaHistoriesResponse() (response *ListFlywaySchemaHistoriesResponse) {
	response = &ListFlywaySchemaHistoriesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListFlywaySchemaHistories(request *ListFlywaySchemaHistoriesRequest) (response *ListFlywaySchemaHistoriesResponse, err error) {
	if request == nil {
		request = NewListFlywaySchemaHistoriesRequest()
	}
	response = NewListFlywaySchemaHistoriesResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePlanSiteRequest() (request *CreatePlanSiteRequest) {
	request = &CreatePlanSiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreatePlanSite")
	return
}

func NewCreatePlanSiteResponse() (response *CreatePlanSiteResponse) {
	response = &CreatePlanSiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreatePlanSite(request *CreatePlanSiteRequest) (response *CreatePlanSiteResponse, err error) {
	if request == nil {
		request = NewCreatePlanSiteRequest()
	}
	response = NewCreatePlanSiteResponse()
	err = c.Send(request, response)
	return
}

func NewGetPlanProductInstanceRequest() (request *GetPlanProductInstanceRequest) {
	request = &GetPlanProductInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetPlanProductInstance")
	return
}

func NewGetPlanProductInstanceResponse() (response *GetPlanProductInstanceResponse) {
	response = &GetPlanProductInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetPlanProductInstance(request *GetPlanProductInstanceRequest) (response *GetPlanProductInstanceResponse, err error) {
	if request == nil {
		request = NewGetPlanProductInstanceRequest()
	}
	response = NewGetPlanProductInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewListPlanProductInstancesRequest() (request *ListPlanProductInstancesRequest) {
	request = &ListPlanProductInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListPlanProductInstances")
	return
}

func NewListPlanProductInstancesResponse() (response *ListPlanProductInstancesResponse) {
	response = &ListPlanProductInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListPlanProductInstances(request *ListPlanProductInstancesRequest) (response *ListPlanProductInstancesResponse, err error) {
	if request == nil {
		request = NewListPlanProductInstancesRequest()
	}
	response = NewListPlanProductInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePlanGlobalConfigRequest() (request *DeletePlanGlobalConfigRequest) {
	request = &DeletePlanGlobalConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeletePlanGlobalConfig")
	return
}

func NewDeletePlanGlobalConfigResponse() (response *DeletePlanGlobalConfigResponse) {
	response = &DeletePlanGlobalConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeletePlanGlobalConfig(request *DeletePlanGlobalConfigRequest) (response *DeletePlanGlobalConfigResponse, err error) {
	if request == nil {
		request = NewDeletePlanGlobalConfigRequest()
	}
	response = NewDeletePlanGlobalConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductInstanceBackupRequest() (request *DeleteProductInstanceBackupRequest) {
	request = &DeleteProductInstanceBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeleteProductInstanceBackup")
	return
}

func NewDeleteProductInstanceBackupResponse() (response *DeleteProductInstanceBackupResponse) {
	response = &DeleteProductInstanceBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeleteProductInstanceBackup(request *DeleteProductInstanceBackupRequest) (response *DeleteProductInstanceBackupResponse, err error) {
	if request == nil {
		request = NewDeleteProductInstanceBackupRequest()
	}
	response = NewDeleteProductInstanceBackupResponse()
	err = c.Send(request, response)
	return
}

func NewListProductVersionApplicationPackageRelDetailsRequest() (request *ListProductVersionApplicationPackageRelDetailsRequest) {
	request = &ListProductVersionApplicationPackageRelDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListProductVersionApplicationPackageRelDetails")
	return
}

func NewListProductVersionApplicationPackageRelDetailsResponse() (response *ListProductVersionApplicationPackageRelDetailsResponse) {
	response = &ListProductVersionApplicationPackageRelDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListProductVersionApplicationPackageRelDetails(request *ListProductVersionApplicationPackageRelDetailsRequest) (response *ListProductVersionApplicationPackageRelDetailsResponse, err error) {
	if request == nil {
		request = NewListProductVersionApplicationPackageRelDetailsRequest()
	}
	response = NewListProductVersionApplicationPackageRelDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApplicationArtifactRequest() (request *CreateApplicationArtifactRequest) {
	request = &CreateApplicationArtifactRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateApplicationArtifact")
	return
}

func NewCreateApplicationArtifactResponse() (response *CreateApplicationArtifactResponse) {
	response = &CreateApplicationArtifactResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateApplicationArtifact(request *CreateApplicationArtifactRequest) (response *CreateApplicationArtifactResponse, err error) {
	if request == nil {
		request = NewCreateApplicationArtifactRequest()
	}
	response = NewCreateApplicationArtifactResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApplicationRunningPlanRequest() (request *UpdateApplicationRunningPlanRequest) {
	request = &UpdateApplicationRunningPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateApplicationRunningPlan")
	return
}

func NewUpdateApplicationRunningPlanResponse() (response *UpdateApplicationRunningPlanResponse) {
	response = &UpdateApplicationRunningPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateApplicationRunningPlan(request *UpdateApplicationRunningPlanRequest) (response *UpdateApplicationRunningPlanResponse, err error) {
	if request == nil {
		request = NewUpdateApplicationRunningPlanRequest()
	}
	response = NewUpdateApplicationRunningPlanResponse()
	err = c.Send(request, response)
	return
}

func NewGetApplicationRequest() (request *GetApplicationRequest) {
	request = &GetApplicationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "GetApplication")
	return
}

func NewGetApplicationResponse() (response *GetApplicationResponse) {
	response = &GetApplicationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) GetApplication(request *GetApplicationRequest) (response *GetApplicationResponse, err error) {
	if request == nil {
		request = NewGetApplicationRequest()
	}
	response = NewGetApplicationResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreatePlanApplicationRuntimeNamesRequest() (request *BulkCreatePlanApplicationRuntimeNamesRequest) {
	request = &BulkCreatePlanApplicationRuntimeNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreatePlanApplicationRuntimeNames")
	return
}

func NewBulkCreatePlanApplicationRuntimeNamesResponse() (response *BulkCreatePlanApplicationRuntimeNamesResponse) {
	response = &BulkCreatePlanApplicationRuntimeNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreatePlanApplicationRuntimeNames(request *BulkCreatePlanApplicationRuntimeNamesRequest) (response *BulkCreatePlanApplicationRuntimeNamesResponse, err error) {
	if request == nil {
		request = NewBulkCreatePlanApplicationRuntimeNamesRequest()
	}
	response = NewBulkCreatePlanApplicationRuntimeNamesResponse()
	err = c.Send(request, response)
	return
}

func NewListApplicationInstanceStatesRequest() (request *ListApplicationInstanceStatesRequest) {
	request = &ListApplicationInstanceStatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListApplicationInstanceStates")
	return
}

func NewListApplicationInstanceStatesResponse() (response *ListApplicationInstanceStatesResponse) {
	response = &ListApplicationInstanceStatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListApplicationInstanceStates(request *ListApplicationInstanceStatesRequest) (response *ListApplicationInstanceStatesResponse, err error) {
	if request == nil {
		request = NewListApplicationInstanceStatesRequest()
	}
	response = NewListApplicationInstanceStatesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProductInfoRequest() (request *UpdateProductInfoRequest) {
	request = &UpdateProductInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateProductInfo")
	return
}

func NewUpdateProductInfoResponse() (response *UpdateProductInfoResponse) {
	response = &UpdateProductInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateProductInfo(request *UpdateProductInfoRequest) (response *UpdateProductInfoResponse, err error) {
	if request == nil {
		request = NewUpdateProductInfoRequest()
	}
	response = NewUpdateProductInfoResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateProjectSiteInfosRequest() (request *BulkCreateProjectSiteInfosRequest) {
	request = &BulkCreateProjectSiteInfosRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateProjectSiteInfos")
	return
}

func NewBulkCreateProjectSiteInfosResponse() (response *BulkCreateProjectSiteInfosResponse) {
	response = &BulkCreateProjectSiteInfosResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateProjectSiteInfos(request *BulkCreateProjectSiteInfosRequest) (response *BulkCreateProjectSiteInfosResponse, err error) {
	if request == nil {
		request = NewBulkCreateProjectSiteInfosRequest()
	}
	response = NewBulkCreateProjectSiteInfosResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePlanProductInstanceRequest() (request *DeletePlanProductInstanceRequest) {
	request = &DeletePlanProductInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "DeletePlanProductInstance")
	return
}

func NewDeletePlanProductInstanceResponse() (response *DeletePlanProductInstanceResponse) {
	response = &DeletePlanProductInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) DeletePlanProductInstance(request *DeletePlanProductInstanceRequest) (response *DeletePlanProductInstanceResponse, err error) {
	if request == nil {
		request = NewDeletePlanProductInstanceRequest()
	}
	response = NewDeletePlanProductInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSolutionTemplateRequest() (request *UpdateSolutionTemplateRequest) {
	request = &UpdateSolutionTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateSolutionTemplate")
	return
}

func NewUpdateSolutionTemplateResponse() (response *UpdateSolutionTemplateResponse) {
	response = &UpdateSolutionTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateSolutionTemplate(request *UpdateSolutionTemplateRequest) (response *UpdateSolutionTemplateResponse, err error) {
	if request == nil {
		request = NewUpdateSolutionTemplateRequest()
	}
	response = NewUpdateSolutionTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewBulkCreateApplicationPackagesRequest() (request *BulkCreateApplicationPackagesRequest) {
	request = &BulkCreateApplicationPackagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "BulkCreateApplicationPackages")
	return
}

func NewBulkCreateApplicationPackagesResponse() (response *BulkCreateApplicationPackagesResponse) {
	response = &BulkCreateApplicationPackagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) BulkCreateApplicationPackages(request *BulkCreateApplicationPackagesRequest) (response *BulkCreateApplicationPackagesResponse, err error) {
	if request == nil {
		request = NewBulkCreateApplicationPackagesRequest()
	}
	response = NewBulkCreateApplicationPackagesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCommonOperationSheetWithDetailRequest() (request *CreateCommonOperationSheetWithDetailRequest) {
	request = &CreateCommonOperationSheetWithDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "CreateCommonOperationSheetWithDetail")
	return
}

func NewCreateCommonOperationSheetWithDetailResponse() (response *CreateCommonOperationSheetWithDetailResponse) {
	response = &CreateCommonOperationSheetWithDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) CreateCommonOperationSheetWithDetail(request *CreateCommonOperationSheetWithDetailRequest) (response *CreateCommonOperationSheetWithDetailResponse, err error) {
	if request == nil {
		request = NewCreateCommonOperationSheetWithDetailRequest()
	}
	response = NewCreateCommonOperationSheetWithDetailResponse()
	err = c.Send(request, response)
	return
}

func NewListApplicationPackagesRequest() (request *ListApplicationPackagesRequest) {
	request = &ListApplicationPackagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListApplicationPackages")
	return
}

func NewListApplicationPackagesResponse() (response *ListApplicationPackagesResponse) {
	response = &ListApplicationPackagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListApplicationPackages(request *ListApplicationPackagesRequest) (response *ListApplicationPackagesResponse, err error) {
	if request == nil {
		request = NewListApplicationPackagesRequest()
	}
	response = NewListApplicationPackagesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProductVersionRequest() (request *UpdateProductVersionRequest) {
	request = &UpdateProductVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateProductVersion")
	return
}

func NewUpdateProductVersionResponse() (response *UpdateProductVersionResponse) {
	response = &UpdateProductVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateProductVersion(request *UpdateProductVersionRequest) (response *UpdateProductVersionResponse, err error) {
	if request == nil {
		request = NewUpdateProductVersionRequest()
	}
	response = NewUpdateProductVersionResponse()
	err = c.Send(request, response)
	return
}

func NewListSolutionVersionDetailsRequest() (request *ListSolutionVersionDetailsRequest) {
	request = &ListSolutionVersionDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "ListSolutionVersionDetails")
	return
}

func NewListSolutionVersionDetailsResponse() (response *ListSolutionVersionDetailsResponse) {
	response = &ListSolutionVersionDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) ListSolutionVersionDetails(request *ListSolutionVersionDetailsRequest) (response *ListSolutionVersionDetailsResponse, err error) {
	if request == nil {
		request = NewListSolutionVersionDetailsRequest()
	}
	response = NewListSolutionVersionDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProductInstanceStateRequest() (request *UpdateProductInstanceStateRequest) {
	request = &UpdateProductInstanceStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dataengine", APIVersion, "UpdateProductInstanceState")
	return
}

func NewUpdateProductInstanceStateResponse() (response *UpdateProductInstanceStateResponse) {
	response = &UpdateProductInstanceStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DataEngine请求
func (c *Client) UpdateProductInstanceState(request *UpdateProductInstanceStateRequest) (response *UpdateProductInstanceStateResponse, err error) {
	if request == nil {
		request = NewUpdateProductInstanceStateRequest()
	}
	response = NewUpdateProductInstanceStateResponse()
	err = c.Send(request, response)
	return
}
