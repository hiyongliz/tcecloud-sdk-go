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

package v20210824

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2021-08-24"

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

func NewDescribeToolRequest() (request *DescribeToolRequest) {
	request = &DescribeToolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeTool")
	return
}

func NewDescribeToolResponse() (response *DescribeToolResponse) {
	response = &DescribeToolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询工具详情
func (c *Client) DescribeTool(request *DescribeToolRequest) (response *DescribeToolResponse, err error) {
	if request == nil {
		request = NewDescribeToolRequest()
	}
	response = NewDescribeToolResponse()
	err = c.Send(request, response)
	return
}

func NewGetOperationTaskLogRequest() (request *GetOperationTaskLogRequest) {
	request = &GetOperationTaskLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "GetOperationTaskLog")
	return
}

func NewGetOperationTaskLogResponse() (response *GetOperationTaskLogResponse) {
	response = &GetOperationTaskLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询变更任务日志
func (c *Client) GetOperationTaskLog(request *GetOperationTaskLogRequest) (response *GetOperationTaskLogResponse, err error) {
	if request == nil {
		request = NewGetOperationTaskLogRequest()
	}
	response = NewGetOperationTaskLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComponentInstanceYAMLRequest() (request *DescribeComponentInstanceYAMLRequest) {
	request = &DescribeComponentInstanceYAMLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeComponentInstanceYAML")
	return
}

func NewDescribeComponentInstanceYAMLResponse() (response *DescribeComponentInstanceYAMLResponse) {
	response = &DescribeComponentInstanceYAMLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询组件实例的YAML
func (c *Client) DescribeComponentInstanceYAML(request *DescribeComponentInstanceYAMLRequest) (response *DescribeComponentInstanceYAMLResponse, err error) {
	if request == nil {
		request = NewDescribeComponentInstanceYAMLRequest()
	}
	response = NewDescribeComponentInstanceYAMLResponse()
	err = c.Send(request, response)
	return
}

func NewListExecutionsRequest() (request *ListExecutionsRequest) {
	request = &ListExecutionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ListExecutions")
	return
}

func NewListExecutionsResponse() (response *ListExecutionsResponse) {
	response = &ListExecutionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出运维工具执行历史
func (c *Client) ListExecutions(request *ListExecutionsRequest) (response *ListExecutionsResponse, err error) {
	if request == nil {
		request = NewListExecutionsRequest()
	}
	response = NewListExecutionsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteFlowRequest() (request *DeleteFlowRequest) {
	request = &DeleteFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DeleteFlow")
	return
}

func NewDeleteFlowResponse() (response *DeleteFlowResponse) {
	response = &DeleteFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Flow
func (c *Client) DeleteFlow(request *DeleteFlowRequest) (response *DeleteFlowResponse, err error) {
	if request == nil {
		request = NewDeleteFlowRequest()
	}
	response = NewDeleteFlowResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProjectRequest() (request *UpdateProjectRequest) {
	request = &UpdateProjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "UpdateProject")
	return
}

func NewUpdateProjectResponse() (response *UpdateProjectResponse) {
	response = &UpdateProjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新项目
func (c *Client) UpdateProject(request *UpdateProjectRequest) (response *UpdateProjectResponse, err error) {
	if request == nil {
		request = NewUpdateProjectRequest()
	}
	response = NewUpdateProjectResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOperationInstructionRequest() (request *CreateOperationInstructionRequest) {
	request = &CreateOperationInstructionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "CreateOperationInstruction")
	return
}

func NewCreateOperationInstructionResponse() (response *CreateOperationInstructionResponse) {
	response = &CreateOperationInstructionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建变更单
func (c *Client) CreateOperationInstruction(request *CreateOperationInstructionRequest) (response *CreateOperationInstructionResponse, err error) {
	if request == nil {
		request = NewCreateOperationInstructionRequest()
	}
	response = NewCreateOperationInstructionResponse()
	err = c.Send(request, response)
	return
}

func NewRollbackToolRequest() (request *RollbackToolRequest) {
	request = &RollbackToolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "RollbackTool")
	return
}

func NewRollbackToolResponse() (response *RollbackToolResponse) {
	response = &RollbackToolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回滚工具版本
func (c *Client) RollbackTool(request *RollbackToolRequest) (response *RollbackToolResponse, err error) {
	if request == nil {
		request = NewRollbackToolRequest()
	}
	response = NewRollbackToolResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExecutionRequest() (request *DescribeExecutionRequest) {
	request = &DescribeExecutionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeExecution")
	return
}

func NewDescribeExecutionResponse() (response *DescribeExecutionResponse) {
	response = &DescribeExecutionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询任务详情
func (c *Client) DescribeExecution(request *DescribeExecutionRequest) (response *DescribeExecutionResponse, err error) {
	if request == nil {
		request = NewDescribeExecutionRequest()
	}
	response = NewDescribeExecutionResponse()
	err = c.Send(request, response)
	return
}

func NewReadyToRollbackPlanRequest() (request *ReadyToRollbackPlanRequest) {
	request = &ReadyToRollbackPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ReadyToRollbackPlan")
	return
}

func NewReadyToRollbackPlanResponse() (response *ReadyToRollbackPlanResponse) {
	response = &ReadyToRollbackPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询是否规划变更可回滚
func (c *Client) ReadyToRollbackPlan(request *ReadyToRollbackPlanRequest) (response *ReadyToRollbackPlanResponse, err error) {
	if request == nil {
		request = NewReadyToRollbackPlanRequest()
	}
	response = NewReadyToRollbackPlanResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductComponentStatusRequest() (request *DescribeProductComponentStatusRequest) {
	request = &DescribeProductComponentStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeProductComponentStatus")
	return
}

func NewDescribeProductComponentStatusResponse() (response *DescribeProductComponentStatusResponse) {
	response = &DescribeProductComponentStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云产品组件健康状态
func (c *Client) DescribeProductComponentStatus(request *DescribeProductComponentStatusRequest) (response *DescribeProductComponentStatusResponse, err error) {
	if request == nil {
		request = NewDescribeProductComponentStatusRequest()
	}
	response = NewDescribeProductComponentStatusResponse()
	err = c.Send(request, response)
	return
}

func NewBatchImportSheetsRequest() (request *BatchImportSheetsRequest) {
	request = &BatchImportSheetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "BatchImportSheets")
	return
}

func NewBatchImportSheetsResponse() (response *BatchImportSheetsResponse) {
	response = &BatchImportSheetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量导入变更单
func (c *Client) BatchImportSheets(request *BatchImportSheetsRequest) (response *BatchImportSheetsResponse, err error) {
	if request == nil {
		request = NewBatchImportSheetsRequest()
	}
	response = NewBatchImportSheetsResponse()
	err = c.Send(request, response)
	return
}

func NewMapInfrastoreProductRequest() (request *MapInfrastoreProductRequest) {
	request = &MapInfrastoreProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "MapInfrastoreProduct")
	return
}

func NewMapInfrastoreProductResponse() (response *MapInfrastoreProductResponse) {
	response = &MapInfrastoreProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 映射产品名到云稍产品名
func (c *Client) MapInfrastoreProduct(request *MapInfrastoreProductRequest) (response *MapInfrastoreProductResponse, err error) {
	if request == nil {
		request = NewMapInfrastoreProductRequest()
	}
	response = NewMapInfrastoreProductResponse()
	err = c.Send(request, response)
	return
}

func NewAdaptInstructionWithPlanRequest() (request *AdaptInstructionWithPlanRequest) {
	request = &AdaptInstructionWithPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "AdaptInstructionWithPlan")
	return
}

func NewAdaptInstructionWithPlanResponse() (response *AdaptInstructionWithPlanResponse) {
	response = &AdaptInstructionWithPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 变更单适配环境规划信息
func (c *Client) AdaptInstructionWithPlan(request *AdaptInstructionWithPlanRequest) (response *AdaptInstructionWithPlanResponse, err error) {
	if request == nil {
		request = NewAdaptInstructionWithPlanRequest()
	}
	response = NewAdaptInstructionWithPlanResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProjectRequest() (request *DescribeProjectRequest) {
	request = &DescribeProjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeProject")
	return
}

func NewDescribeProjectResponse() (response *DescribeProjectResponse) {
	response = &DescribeProjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询项目详情
func (c *Client) DescribeProject(request *DescribeProjectRequest) (response *DescribeProjectResponse, err error) {
	if request == nil {
		request = NewDescribeProjectRequest()
	}
	response = NewDescribeProjectResponse()
	err = c.Send(request, response)
	return
}

func NewListProductComponentsRequest() (request *ListProductComponentsRequest) {
	request = &ListProductComponentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ListProductComponents")
	return
}

func NewListProductComponentsResponse() (response *ListProductComponentsResponse) {
	response = &ListProductComponentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品及组件名字
func (c *Client) ListProductComponents(request *ListProductComponentsRequest) (response *ListProductComponentsResponse, err error) {
	if request == nil {
		request = NewListProductComponentsRequest()
	}
	response = NewListProductComponentsResponse()
	err = c.Send(request, response)
	return
}

func NewSetFlowStatusRequest() (request *SetFlowStatusRequest) {
	request = &SetFlowStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "SetFlowStatus")
	return
}

func NewSetFlowStatusResponse() (response *SetFlowStatusResponse) {
	response = &SetFlowStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置Flow状态
func (c *Client) SetFlowStatus(request *SetFlowStatusRequest) (response *SetFlowStatusResponse, err error) {
	if request == nil {
		request = NewSetFlowStatusRequest()
	}
	response = NewSetFlowStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeToolDownloadURLRequest() (request *DescribeToolDownloadURLRequest) {
	request = &DescribeToolDownloadURLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeToolDownloadURL")
	return
}

func NewDescribeToolDownloadURLResponse() (response *DescribeToolDownloadURLResponse) {
	response = &DescribeToolDownloadURLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取工具临时下载链接
func (c *Client) DescribeToolDownloadURL(request *DescribeToolDownloadURLRequest) (response *DescribeToolDownloadURLResponse, err error) {
	if request == nil {
		request = NewDescribeToolDownloadURLRequest()
	}
	response = NewDescribeToolDownloadURLResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMaterialNodeListRequest() (request *DescribeMaterialNodeListRequest) {
	request = &DescribeMaterialNodeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeMaterialNodeList")
	return
}

func NewDescribeMaterialNodeListResponse() (response *DescribeMaterialNodeListResponse) {
	response = &DescribeMaterialNodeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询物料机列表
func (c *Client) DescribeMaterialNodeList(request *DescribeMaterialNodeListRequest) (response *DescribeMaterialNodeListResponse, err error) {
	if request == nil {
		request = NewDescribeMaterialNodeListRequest()
	}
	response = NewDescribeMaterialNodeListResponse()
	err = c.Send(request, response)
	return
}

func NewListToolVersionsRequest() (request *ListToolVersionsRequest) {
	request = &ListToolVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ListToolVersions")
	return
}

func NewListToolVersionsResponse() (response *ListToolVersionsResponse) {
	response = &ListToolVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询工具历史版本
func (c *Client) ListToolVersions(request *ListToolVersionsRequest) (response *ListToolVersionsResponse, err error) {
	if request == nil {
		request = NewListToolVersionsRequest()
	}
	response = NewListToolVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewParseTarFileDirTreeRequest() (request *ParseTarFileDirTreeRequest) {
	request = &ParseTarFileDirTreeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ParseTarFileDirTree")
	return
}

func NewParseTarFileDirTreeResponse() (response *ParseTarFileDirTreeResponse) {
	response = &ParseTarFileDirTreeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解析压缩包文件目录
func (c *Client) ParseTarFileDirTree(request *ParseTarFileDirTreeRequest) (response *ParseTarFileDirTreeResponse, err error) {
	if request == nil {
		request = NewParseTarFileDirTreeRequest()
	}
	response = NewParseTarFileDirTreeResponse()
	err = c.Send(request, response)
	return
}

func NewListFlowsRequest() (request *ListFlowsRequest) {
	request = &ListFlowsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ListFlows")
	return
}

func NewListFlowsResponse() (response *ListFlowsResponse) {
	response = &ListFlowsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Flow列表
func (c *Client) ListFlows(request *ListFlowsRequest) (response *ListFlowsResponse, err error) {
	if request == nil {
		request = NewListFlowsRequest()
	}
	response = NewListFlowsResponse()
	err = c.Send(request, response)
	return
}

func NewListZonesRequest() (request *ListZonesRequest) {
	request = &ListZonesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ListZones")
	return
}

func NewListZonesResponse() (response *ListZonesResponse) {
	response = &ListZonesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出环境AZ
func (c *Client) ListZones(request *ListZonesRequest) (response *ListZonesResponse, err error) {
	if request == nil {
		request = NewListZonesRequest()
	}
	response = NewListZonesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationInstructionComponentFlowRequest() (request *DescribeOperationInstructionComponentFlowRequest) {
	request = &DescribeOperationInstructionComponentFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeOperationInstructionComponentFlow")
	return
}

func NewDescribeOperationInstructionComponentFlowResponse() (response *DescribeOperationInstructionComponentFlowResponse) {
	response = &DescribeOperationInstructionComponentFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询变更单组件流程
func (c *Client) DescribeOperationInstructionComponentFlow(request *DescribeOperationInstructionComponentFlowRequest) (response *DescribeOperationInstructionComponentFlowResponse, err error) {
	if request == nil {
		request = NewDescribeOperationInstructionComponentFlowRequest()
	}
	response = NewDescribeOperationInstructionComponentFlowResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUpdatePlanTaskRequest() (request *CreateUpdatePlanTaskRequest) {
	request = &CreateUpdatePlanTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "CreateUpdatePlanTask")
	return
}

func NewCreateUpdatePlanTaskResponse() (response *CreateUpdatePlanTaskResponse) {
	response = &CreateUpdatePlanTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建上传规划任务
func (c *Client) CreateUpdatePlanTask(request *CreateUpdatePlanTaskRequest) (response *CreateUpdatePlanTaskResponse, err error) {
	if request == nil {
		request = NewCreateUpdatePlanTaskRequest()
	}
	response = NewCreateUpdatePlanTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteToolRequest() (request *DeleteToolRequest) {
	request = &DeleteToolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DeleteTool")
	return
}

func NewDeleteToolResponse() (response *DeleteToolResponse) {
	response = &DeleteToolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除工具
func (c *Client) DeleteTool(request *DeleteToolRequest) (response *DeleteToolResponse, err error) {
	if request == nil {
		request = NewDeleteToolRequest()
	}
	response = NewDeleteToolResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComponentCurrentVersionRequest() (request *DescribeComponentCurrentVersionRequest) {
	request = &DescribeComponentCurrentVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeComponentCurrentVersion")
	return
}

func NewDescribeComponentCurrentVersionResponse() (response *DescribeComponentCurrentVersionResponse) {
	response = &DescribeComponentCurrentVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询变更单中组件当前版本
func (c *Client) DescribeComponentCurrentVersion(request *DescribeComponentCurrentVersionRequest) (response *DescribeComponentCurrentVersionResponse, err error) {
	if request == nil {
		request = NewDescribeComponentCurrentVersionRequest()
	}
	response = NewDescribeComponentCurrentVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUploadPackageTaskRequest() (request *DescribeUploadPackageTaskRequest) {
	request = &DescribeUploadPackageTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeUploadPackageTask")
	return
}

func NewDescribeUploadPackageTaskResponse() (response *DescribeUploadPackageTaskResponse) {
	response = &DescribeUploadPackageTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询上传任务进度
func (c *Client) DescribeUploadPackageTask(request *DescribeUploadPackageTaskRequest) (response *DescribeUploadPackageTaskResponse, err error) {
	if request == nil {
		request = NewDescribeUploadPackageTaskRequest()
	}
	response = NewDescribeUploadPackageTaskResponse()
	err = c.Send(request, response)
	return
}

func NewImportToolRequest() (request *ImportToolRequest) {
	request = &ImportToolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ImportTool")
	return
}

func NewImportToolResponse() (response *ImportToolResponse) {
	response = &ImportToolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入工具
func (c *Client) ImportTool(request *ImportToolRequest) (response *ImportToolResponse, err error) {
	if request == nil {
		request = NewImportToolRequest()
	}
	response = NewImportToolResponse()
	err = c.Send(request, response)
	return
}

func NewCreateToolParseURLRequest() (request *CreateToolParseURLRequest) {
	request = &CreateToolParseURLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "CreateToolParseURL")
	return
}

func NewCreateToolParseURLResponse() (response *CreateToolParseURLResponse) {
	response = &CreateToolParseURLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取解析压缩包的上传链接
func (c *Client) CreateToolParseURL(request *CreateToolParseURLRequest) (response *CreateToolParseURLResponse, err error) {
	if request == nil {
		request = NewCreateToolParseURLRequest()
	}
	response = NewCreateToolParseURLResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUploadPackagesTaskRequest() (request *CreateUploadPackagesTaskRequest) {
	request = &CreateUploadPackagesTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "CreateUploadPackagesTask")
	return
}

func NewCreateUploadPackagesTaskResponse() (response *CreateUploadPackagesTaskResponse) {
	response = &CreateUploadPackagesTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建上传物料任务
func (c *Client) CreateUploadPackagesTask(request *CreateUploadPackagesTaskRequest) (response *CreateUploadPackagesTaskResponse, err error) {
	if request == nil {
		request = NewCreateUploadPackagesTaskRequest()
	}
	response = NewCreateUploadPackagesTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationInstructionDAGRequest() (request *DescribeOperationInstructionDAGRequest) {
	request = &DescribeOperationInstructionDAGRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeOperationInstructionDAG")
	return
}

func NewDescribeOperationInstructionDAGResponse() (response *DescribeOperationInstructionDAGResponse) {
	response = &DescribeOperationInstructionDAGResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询变更单DAG
func (c *Client) DescribeOperationInstructionDAG(request *DescribeOperationInstructionDAGRequest) (response *DescribeOperationInstructionDAGResponse, err error) {
	if request == nil {
		request = NewDescribeOperationInstructionDAGRequest()
	}
	response = NewDescribeOperationInstructionDAGResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOperationInstructionRequest() (request *UpdateOperationInstructionRequest) {
	request = &UpdateOperationInstructionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "UpdateOperationInstruction")
	return
}

func NewUpdateOperationInstructionResponse() (response *UpdateOperationInstructionResponse) {
	response = &UpdateOperationInstructionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改变更单
func (c *Client) UpdateOperationInstruction(request *UpdateOperationInstructionRequest) (response *UpdateOperationInstructionResponse, err error) {
	if request == nil {
		request = NewUpdateOperationInstructionRequest()
	}
	response = NewUpdateOperationInstructionResponse()
	err = c.Send(request, response)
	return
}

func NewListToolsRequest() (request *ListToolsRequest) {
	request = &ListToolsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ListTools")
	return
}

func NewListToolsResponse() (response *ListToolsResponse) {
	response = &ListToolsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询工具列表
func (c *Client) ListTools(request *ListToolsRequest) (response *ListToolsResponse, err error) {
	if request == nil {
		request = NewListToolsRequest()
	}
	response = NewListToolsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyToolRequest() (request *ModifyToolRequest) {
	request = &ModifyToolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ModifyTool")
	return
}

func NewModifyToolResponse() (response *ModifyToolResponse) {
	response = &ModifyToolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改工具
func (c *Client) ModifyTool(request *ModifyToolRequest) (response *ModifyToolResponse, err error) {
	if request == nil {
		request = NewModifyToolRequest()
	}
	response = NewModifyToolResponse()
	err = c.Send(request, response)
	return
}

func NewListCloudProductsRequest() (request *ListCloudProductsRequest) {
	request = &ListCloudProductsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ListCloudProducts")
	return
}

func NewListCloudProductsResponse() (response *ListCloudProductsResponse) {
	response = &ListCloudProductsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云产品列表
func (c *Client) ListCloudProducts(request *ListCloudProductsRequest) (response *ListCloudProductsResponse, err error) {
	if request == nil {
		request = NewListCloudProductsRequest()
	}
	response = NewListCloudProductsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOperationInstructionRequest() (request *DeleteOperationInstructionRequest) {
	request = &DeleteOperationInstructionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DeleteOperationInstruction")
	return
}

func NewDeleteOperationInstructionResponse() (response *DeleteOperationInstructionResponse) {
	response = &DeleteOperationInstructionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除变更单
func (c *Client) DeleteOperationInstruction(request *DeleteOperationInstructionRequest) (response *DeleteOperationInstructionResponse, err error) {
	if request == nil {
		request = NewDeleteOperationInstructionRequest()
	}
	response = NewDeleteOperationInstructionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowRequest() (request *DescribeFlowRequest) {
	request = &DescribeFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeFlow")
	return
}

func NewDescribeFlowResponse() (response *DescribeFlowResponse) {
	response = &DescribeFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Flow
func (c *Client) DescribeFlow(request *DescribeFlowRequest) (response *DescribeFlowResponse, err error) {
	if request == nil {
		request = NewDescribeFlowRequest()
	}
	response = NewDescribeFlowResponse()
	err = c.Send(request, response)
	return
}

func NewCreateToolImportURLRequest() (request *CreateToolImportURLRequest) {
	request = &CreateToolImportURLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "CreateToolImportURL")
	return
}

func NewCreateToolImportURLResponse() (response *CreateToolImportURLResponse) {
	response = &CreateToolImportURLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取导入工具的上传链接
func (c *Client) CreateToolImportURL(request *CreateToolImportURLRequest) (response *CreateToolImportURLResponse, err error) {
	if request == nil {
		request = NewCreateToolImportURLRequest()
	}
	response = NewCreateToolImportURLResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBizTreeHostsRequest() (request *DescribeBizTreeHostsRequest) {
	request = &DescribeBizTreeHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeBizTreeHosts")
	return
}

func NewDescribeBizTreeHostsResponse() (response *DescribeBizTreeHostsResponse) {
	response = &DescribeBizTreeHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询业务树结点对应主机
func (c *Client) DescribeBizTreeHosts(request *DescribeBizTreeHostsRequest) (response *DescribeBizTreeHostsResponse, err error) {
	if request == nil {
		request = NewDescribeBizTreeHostsRequest()
	}
	response = NewDescribeBizTreeHostsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductComponentResourceRequest() (request *DescribeProductComponentResourceRequest) {
	request = &DescribeProductComponentResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeProductComponentResource")
	return
}

func NewDescribeProductComponentResourceResponse() (response *DescribeProductComponentResourceResponse) {
	response = &DescribeProductComponentResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询组件使用的pod及主机资源
func (c *Client) DescribeProductComponentResource(request *DescribeProductComponentResourceRequest) (response *DescribeProductComponentResourceResponse, err error) {
	if request == nil {
		request = NewDescribeProductComponentResourceRequest()
	}
	response = NewDescribeProductComponentResourceResponse()
	err = c.Send(request, response)
	return
}

func NewListProductAlertsRequest() (request *ListProductAlertsRequest) {
	request = &ListProductAlertsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ListProductAlerts")
	return
}

func NewListProductAlertsResponse() (response *ListProductAlertsResponse) {
	response = &ListProductAlertsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云产品告警列表
func (c *Client) ListProductAlerts(request *ListProductAlertsRequest) (response *ListProductAlertsResponse, err error) {
	if request == nil {
		request = NewListProductAlertsRequest()
	}
	response = NewListProductAlertsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationInstructionLogRequest() (request *DescribeOperationInstructionLogRequest) {
	request = &DescribeOperationInstructionLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeOperationInstructionLog")
	return
}

func NewDescribeOperationInstructionLogResponse() (response *DescribeOperationInstructionLogResponse) {
	response = &DescribeOperationInstructionLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询变更单任务日志
func (c *Client) DescribeOperationInstructionLog(request *DescribeOperationInstructionLogRequest) (response *DescribeOperationInstructionLogResponse, err error) {
	if request == nil {
		request = NewDescribeOperationInstructionLogRequest()
	}
	response = NewDescribeOperationInstructionLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeComponentPlanRequest() (request *DescribeComponentPlanRequest) {
	request = &DescribeComponentPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeComponentPlan")
	return
}

func NewDescribeComponentPlanResponse() (response *DescribeComponentPlanResponse) {
	response = &DescribeComponentPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询组件规划详情
func (c *Client) DescribeComponentPlan(request *DescribeComponentPlanRequest) (response *DescribeComponentPlanResponse, err error) {
	if request == nil {
		request = NewDescribeComponentPlanRequest()
	}
	response = NewDescribeComponentPlanResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePackageUploadURLRequest() (request *CreatePackageUploadURLRequest) {
	request = &CreatePackageUploadURLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "CreatePackageUploadURL")
	return
}

func NewCreatePackageUploadURLResponse() (response *CreatePackageUploadURLResponse) {
	response = &CreatePackageUploadURLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建组件包上传链接
func (c *Client) CreatePackageUploadURL(request *CreatePackageUploadURLRequest) (response *CreatePackageUploadURLResponse, err error) {
	if request == nil {
		request = NewCreatePackageUploadURLRequest()
	}
	response = NewCreatePackageUploadURLResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudProductRequest() (request *DescribeCloudProductRequest) {
	request = &DescribeCloudProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeCloudProduct")
	return
}

func NewDescribeCloudProductResponse() (response *DescribeCloudProductResponse) {
	response = &DescribeCloudProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云产品详情
func (c *Client) DescribeCloudProduct(request *DescribeCloudProductRequest) (response *DescribeCloudProductResponse, err error) {
	if request == nil {
		request = NewDescribeCloudProductRequest()
	}
	response = NewDescribeCloudProductResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductComponentsRequest() (request *DescribeProductComponentsRequest) {
	request = &DescribeProductComponentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeProductComponents")
	return
}

func NewDescribeProductComponentsResponse() (response *DescribeProductComponentsResponse) {
	response = &DescribeProductComponentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出产品下的组件
func (c *Client) DescribeProductComponents(request *DescribeProductComponentsRequest) (response *DescribeProductComponentsResponse, err error) {
	if request == nil {
		request = NewDescribeProductComponentsRequest()
	}
	response = NewDescribeProductComponentsResponse()
	err = c.Send(request, response)
	return
}

func NewListToolLabelsRequest() (request *ListToolLabelsRequest) {
	request = &ListToolLabelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ListToolLabels")
	return
}

func NewListToolLabelsResponse() (response *ListToolLabelsResponse) {
	response = &ListToolLabelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询工具标签
func (c *Client) ListToolLabels(request *ListToolLabelsRequest) (response *ListToolLabelsResponse, err error) {
	if request == nil {
		request = NewListToolLabelsRequest()
	}
	response = NewListToolLabelsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUploadPackageTaskRequest() (request *CreateUploadPackageTaskRequest) {
	request = &CreateUploadPackageTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "CreateUploadPackageTask")
	return
}

func NewCreateUploadPackageTaskResponse() (response *CreateUploadPackageTaskResponse) {
	response = &CreateUploadPackageTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开始上传组件
func (c *Client) CreateUploadPackageTask(request *CreateUploadPackageTaskRequest) (response *CreateUploadPackageTaskResponse, err error) {
	if request == nil {
		request = NewCreateUploadPackageTaskRequest()
	}
	response = NewCreateUploadPackageTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCurrentPlanVersionRequest() (request *DescribeCurrentPlanVersionRequest) {
	request = &DescribeCurrentPlanVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeCurrentPlanVersion")
	return
}

func NewDescribeCurrentPlanVersionResponse() (response *DescribeCurrentPlanVersionResponse) {
	response = &DescribeCurrentPlanVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当前规划包版本
func (c *Client) DescribeCurrentPlanVersion(request *DescribeCurrentPlanVersionRequest) (response *DescribeCurrentPlanVersionResponse, err error) {
	if request == nil {
		request = NewDescribeCurrentPlanVersionRequest()
	}
	response = NewDescribeCurrentPlanVersionResponse()
	err = c.Send(request, response)
	return
}

func NewSetOperationInstructionStatusRequest() (request *SetOperationInstructionStatusRequest) {
	request = &SetOperationInstructionStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "SetOperationInstructionStatus")
	return
}

func NewSetOperationInstructionStatusResponse() (response *SetOperationInstructionStatusResponse) {
	response = &SetOperationInstructionStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置变更单状态
func (c *Client) SetOperationInstructionStatus(request *SetOperationInstructionStatusRequest) (response *SetOperationInstructionStatusResponse, err error) {
	if request == nil {
		request = NewSetOperationInstructionStatusRequest()
	}
	response = NewSetOperationInstructionStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationInstructionRequest() (request *DescribeOperationInstructionRequest) {
	request = &DescribeOperationInstructionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeOperationInstruction")
	return
}

func NewDescribeOperationInstructionResponse() (response *DescribeOperationInstructionResponse) {
	response = &DescribeOperationInstructionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询变更单
func (c *Client) DescribeOperationInstruction(request *DescribeOperationInstructionRequest) (response *DescribeOperationInstructionResponse, err error) {
	if request == nil {
		request = NewDescribeOperationInstructionRequest()
	}
	response = NewDescribeOperationInstructionResponse()
	err = c.Send(request, response)
	return
}

func NewListWorkflowCommandsRequest() (request *ListWorkflowCommandsRequest) {
	request = &ListWorkflowCommandsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ListWorkflowCommands")
	return
}

func NewListWorkflowCommandsResponse() (response *ListWorkflowCommandsResponse) {
	response = &ListWorkflowCommandsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出原子命令列表
func (c *Client) ListWorkflowCommands(request *ListWorkflowCommandsRequest) (response *ListWorkflowCommandsResponse, err error) {
	if request == nil {
		request = NewListWorkflowCommandsRequest()
	}
	response = NewListWorkflowCommandsResponse()
	err = c.Send(request, response)
	return
}

func NewOperationInstructionReadyToStartRequest() (request *OperationInstructionReadyToStartRequest) {
	request = &OperationInstructionReadyToStartRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "OperationInstructionReadyToStart")
	return
}

func NewOperationInstructionReadyToStartResponse() (response *OperationInstructionReadyToStartResponse) {
	response = &OperationInstructionReadyToStartResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询变更单是否可启动
func (c *Client) OperationInstructionReadyToStart(request *OperationInstructionReadyToStartRequest) (response *OperationInstructionReadyToStartResponse, err error) {
	if request == nil {
		request = NewOperationInstructionReadyToStartRequest()
	}
	response = NewOperationInstructionReadyToStartResponse()
	err = c.Send(request, response)
	return
}

func NewParseInstructionComponentsRequest() (request *ParseInstructionComponentsRequest) {
	request = &ParseInstructionComponentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ParseInstructionComponents")
	return
}

func NewParseInstructionComponentsResponse() (response *ParseInstructionComponentsResponse) {
	response = &ParseInstructionComponentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解析变跟单YAML获取待变更组件列表
func (c *Client) ParseInstructionComponents(request *ParseInstructionComponentsRequest) (response *ParseInstructionComponentsResponse, err error) {
	if request == nil {
		request = NewParseInstructionComponentsRequest()
	}
	response = NewParseInstructionComponentsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductBizTreeRequest() (request *DescribeProductBizTreeRequest) {
	request = &DescribeProductBizTreeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeProductBizTree")
	return
}

func NewDescribeProductBizTreeResponse() (response *DescribeProductBizTreeResponse) {
	response = &DescribeProductBizTreeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品运维中心某个云产品的业务树
func (c *Client) DescribeProductBizTree(request *DescribeProductBizTreeRequest) (response *DescribeProductBizTreeResponse, err error) {
	if request == nil {
		request = NewDescribeProductBizTreeRequest()
	}
	response = NewDescribeProductBizTreeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationStatsRequest() (request *DescribeOperationStatsRequest) {
	request = &DescribeOperationStatsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeOperationStats")
	return
}

func NewDescribeOperationStatsResponse() (response *DescribeOperationStatsResponse) {
	response = &DescribeOperationStatsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询变更记录统计
func (c *Client) DescribeOperationStats(request *DescribeOperationStatsRequest) (response *DescribeOperationStatsResponse, err error) {
	if request == nil {
		request = NewDescribeOperationStatsRequest()
	}
	response = NewDescribeOperationStatsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProjectPlanRequest() (request *DescribeProjectPlanRequest) {
	request = &DescribeProjectPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeProjectPlan")
	return
}

func NewDescribeProjectPlanResponse() (response *DescribeProjectPlanResponse) {
	response = &DescribeProjectPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询项目下各个变更单规划数据
func (c *Client) DescribeProjectPlan(request *DescribeProjectPlanRequest) (response *DescribeProjectPlanResponse, err error) {
	if request == nil {
		request = NewDescribeProjectPlanRequest()
	}
	response = NewDescribeProjectPlanResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBizTreeRequest() (request *DescribeBizTreeRequest) {
	request = &DescribeBizTreeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeBizTree")
	return
}

func NewDescribeBizTreeResponse() (response *DescribeBizTreeResponse) {
	response = &DescribeBizTreeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询CMDB业务树
func (c *Client) DescribeBizTree(request *DescribeBizTreeRequest) (response *DescribeBizTreeResponse, err error) {
	if request == nil {
		request = NewDescribeBizTreeRequest()
	}
	response = NewDescribeBizTreeResponse()
	err = c.Send(request, response)
	return
}

func NewExecuteToolRequest() (request *ExecuteToolRequest) {
	request = &ExecuteToolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ExecuteTool")
	return
}

func NewExecuteToolResponse() (response *ExecuteToolResponse) {
	response = &ExecuteToolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 执行工具
func (c *Client) ExecuteTool(request *ExecuteToolRequest) (response *ExecuteToolResponse, err error) {
	if request == nil {
		request = NewExecuteToolRequest()
	}
	response = NewExecuteToolResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductTopologyRequest() (request *DescribeProductTopologyRequest) {
	request = &DescribeProductTopologyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeProductTopology")
	return
}

func NewDescribeProductTopologyResponse() (response *DescribeProductTopologyResponse) {
	response = &DescribeProductTopologyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云产品运行拓扑
func (c *Client) DescribeProductTopology(request *DescribeProductTopologyRequest) (response *DescribeProductTopologyResponse, err error) {
	if request == nil {
		request = NewDescribeProductTopologyRequest()
	}
	response = NewDescribeProductTopologyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeToolStatsRequest() (request *DescribeToolStatsRequest) {
	request = &DescribeToolStatsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeToolStats")
	return
}

func NewDescribeToolStatsResponse() (response *DescribeToolStatsResponse) {
	response = &DescribeToolStatsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运维工具统计数据
func (c *Client) DescribeToolStats(request *DescribeToolStatsRequest) (response *DescribeToolStatsResponse, err error) {
	if request == nil {
		request = NewDescribeToolStatsRequest()
	}
	response = NewDescribeToolStatsResponse()
	err = c.Send(request, response)
	return
}

func NewListOperationInstructionsRequest() (request *ListOperationInstructionsRequest) {
	request = &ListOperationInstructionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ListOperationInstructions")
	return
}

func NewListOperationInstructionsResponse() (response *ListOperationInstructionsResponse) {
	response = &ListOperationInstructionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出变更控制表
func (c *Client) ListOperationInstructions(request *ListOperationInstructionsRequest) (response *ListOperationInstructionsResponse, err error) {
	if request == nil {
		request = NewListOperationInstructionsRequest()
	}
	response = NewListOperationInstructionsResponse()
	err = c.Send(request, response)
	return
}

func NewListProjectsRequest() (request *ListProjectsRequest) {
	request = &ListProjectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ListProjects")
	return
}

func NewListProjectsResponse() (response *ListProjectsResponse) {
	response = &ListProjectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询项目列表
func (c *Client) ListProjects(request *ListProjectsRequest) (response *ListProjectsResponse, err error) {
	if request == nil {
		request = NewListProjectsRequest()
	}
	response = NewListProjectsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateToolUploadURLRequest() (request *CreateToolUploadURLRequest) {
	request = &CreateToolUploadURLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "CreateToolUploadURL")
	return
}

func NewCreateToolUploadURLResponse() (response *CreateToolUploadURLResponse) {
	response = &CreateToolUploadURLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于创建、更新工具时获取文件上传链接
func (c *Client) CreateToolUploadURL(request *CreateToolUploadURLRequest) (response *CreateToolUploadURLResponse, err error) {
	if request == nil {
		request = NewCreateToolUploadURLRequest()
	}
	response = NewCreateToolUploadURLResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationInstructionPlanRequest() (request *DescribeOperationInstructionPlanRequest) {
	request = &DescribeOperationInstructionPlanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeOperationInstructionPlan")
	return
}

func NewDescribeOperationInstructionPlanResponse() (response *DescribeOperationInstructionPlanResponse) {
	response = &DescribeOperationInstructionPlanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询变更单组件规划
func (c *Client) DescribeOperationInstructionPlan(request *DescribeOperationInstructionPlanRequest) (response *DescribeOperationInstructionPlanResponse, err error) {
	if request == nil {
		request = NewDescribeOperationInstructionPlanRequest()
	}
	response = NewDescribeOperationInstructionPlanResponse()
	err = c.Send(request, response)
	return
}

func NewCreateToolRequest() (request *CreateToolRequest) {
	request = &CreateToolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "CreateTool")
	return
}

func NewCreateToolResponse() (response *CreateToolResponse) {
	response = &CreateToolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建工具
func (c *Client) CreateTool(request *CreateToolRequest) (response *CreateToolResponse, err error) {
	if request == nil {
		request = NewCreateToolRequest()
	}
	response = NewCreateToolResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductAlertStatsRequest() (request *DescribeProductAlertStatsRequest) {
	request = &DescribeProductAlertStatsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeProductAlertStats")
	return
}

func NewDescribeProductAlertStatsResponse() (response *DescribeProductAlertStatsResponse) {
	response = &DescribeProductAlertStatsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询云产品告警统计数据
func (c *Client) DescribeProductAlertStats(request *DescribeProductAlertStatsRequest) (response *DescribeProductAlertStatsResponse, err error) {
	if request == nil {
		request = NewDescribeProductAlertStatsRequest()
	}
	response = NewDescribeProductAlertStatsResponse()
	err = c.Send(request, response)
	return
}

func NewSetTaskInstanceStatusRequest() (request *SetTaskInstanceStatusRequest) {
	request = &SetTaskInstanceStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "SetTaskInstanceStatus")
	return
}

func NewSetTaskInstanceStatusResponse() (response *SetTaskInstanceStatusResponse) {
	response = &SetTaskInstanceStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重试/跳过某个任务步骤
func (c *Client) SetTaskInstanceStatus(request *SetTaskInstanceStatusRequest) (response *SetTaskInstanceStatusResponse, err error) {
	if request == nil {
		request = NewSetTaskInstanceStatusRequest()
	}
	response = NewSetTaskInstanceStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFindPackagePathTaskRequest() (request *DescribeFindPackagePathTaskRequest) {
	request = &DescribeFindPackagePathTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "DescribeFindPackagePathTask")
	return
}

func NewDescribeFindPackagePathTaskResponse() (response *DescribeFindPackagePathTaskResponse) {
	response = &DescribeFindPackagePathTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询查找物料路径任务详情
func (c *Client) DescribeFindPackagePathTask(request *DescribeFindPackagePathTaskRequest) (response *DescribeFindPackagePathTaskResponse, err error) {
	if request == nil {
		request = NewDescribeFindPackagePathTaskRequest()
	}
	response = NewDescribeFindPackagePathTaskResponse()
	err = c.Send(request, response)
	return
}

func NewImportChangePackageRequest() (request *ImportChangePackageRequest) {
	request = &ImportChangePackageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tops", APIVersion, "ImportChangePackage")
	return
}

func NewImportChangePackageResponse() (response *ImportChangePackageResponse) {
	response = &ImportChangePackageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入单个变更单
func (c *Client) ImportChangePackage(request *ImportChangePackageRequest) (response *ImportChangePackageResponse, err error) {
	if request == nil {
		request = NewImportChangePackageRequest()
	}
	response = NewImportChangePackageResponse()
	err = c.Send(request, response)
	return
}
