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
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type UpdateApplicationInstanceDeployVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApplicationInstanceDeployVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationInstanceDeployVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateDagNodesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateDagNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateDagNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationRunningPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationRunningPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationRunningPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSolutionTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSolutionTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSolutionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationBranchRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateApplicationBranchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationBranchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCommonOperationSheetRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateCommonOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCommonOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductTemplatesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateProductTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialImportTaskRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateMaterialImportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialImportTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialMetadataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaterialMetadataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialMetadataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductVersionApplicationPackageRelsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProductVersionApplicationPackageRelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductVersionApplicationPackageRelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialImportTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMaterialImportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialImportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialInfosRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateMaterialInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSiteOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSiteOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSiteOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanServerInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPlanServerInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanServerInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSystemSettingRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetSystemSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSystemSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectTopologyDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProjectTopologyDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectTopologyDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectZonesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProjectZonesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectZonesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrchTemplateLabelRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateOrchTemplateLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOrchTemplateLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanApplicationInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePlanApplicationInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanApplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductVersionApplicationPackageRelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductVersionApplicationPackageRelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductVersionApplicationPackageRelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationSheetApprovalRecordRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateOperationSheetApprovalRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationSheetApprovalRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationBranchRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteApplicationBranchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationBranchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductOperationJobRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateProductOperationJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductOperationJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateAtomicsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateAtomicsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateAtomicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSolutionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteSolutionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSolutionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanTagRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdatePlanTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationRunningDataCollectionsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateApplicationRunningDataCollectionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationRunningDataCollectionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanApplicationInstancesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreatePlanApplicationInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanApplicationInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationCollectStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationCollectStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationCollectStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialCleanupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMaterialCleanupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialCleanupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanApplicationInstancesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListPlanApplicationInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanApplicationInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectTopologyDetailsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProjectTopologyDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectTopologyDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationSheetApprovalRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperationSheetApprovalRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationSheetApprovalRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCommonOperationSheetWithDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetCommonOperationSheetWithDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCommonOperationSheetWithDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFlywaySchemaHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFlywaySchemaHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFlywaySchemaHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectSiteInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProjectSiteInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectSiteInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialMetadataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMaterialMetadataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialMetadataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDagNodesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListDagNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDagNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationArtifactResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApplicationArtifactResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationArtifactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateSiteOperationSheetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateSiteOperationSheetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateSiteOperationSheetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationSheetApprovalRecordRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateOperationSheetApprovalRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationSheetApprovalRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductVersionArtifactTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductVersionArtifactTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductVersionArtifactTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductOperationJobsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductOperationJobsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductOperationJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanCustomConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPlanCustomConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanCustomConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationBranchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApplicationBranchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationBranchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationSheetAttachmentRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateOperationSheetAttachmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationSheetAttachmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialSitesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateMaterialSitesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialSitesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreatePlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectTopologyDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProjectTopologyDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectTopologyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanApplicationInstanceRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetPlanApplicationInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanApplicationInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductInstanceBackupRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProductInstanceBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductInstanceBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialImportTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateMaterialImportTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialImportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationSheetAttachmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperationSheetAttachmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationSheetAttachmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductInstanceOperationSheetRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteProductInstanceOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductInstanceOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationArtifactResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationArtifactResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationArtifactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationInstanceBackupRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateApplicationInstanceBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationInstanceBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOrchTemplateLabelRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetOrchTemplateLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOrchTemplateLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationArtifactChartHistoriesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationArtifactChartHistoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationArtifactChartHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationArtifactsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationArtifactsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationArtifactsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMaterialSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductInstanceOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProductInstanceOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductInstanceOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationPipelinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateOperationPipelinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationPipelinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanServerInfosRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreatePlanServerInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanServerInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDagNodeRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteDagNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDagNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialCustomizedTypeRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetMaterialCustomizedTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialCustomizedTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSystemSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSystemSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSystemSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTenantResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTenantResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTenantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationInstanceStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationInstanceStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationInstanceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectSiteInfoTopologyRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProjectSiteInfoTopologyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectSiteInfoTopologyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectTopologyDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProjectTopologyDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectTopologyDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSolutionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSolutionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSolutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSolutionVersionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateSolutionVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSolutionVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationInstanceBackupRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteApplicationInstanceBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationInstanceBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePlanSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProjectZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationInstanceBackupRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateApplicationInstanceBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationInstanceBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePlanTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanApplicationRuntimeNameRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetPlanApplicationRuntimeNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanApplicationRuntimeNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProjectRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPlanTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationCollectStatusesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateApplicationCollectStatusesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationCollectStatusesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateCommonOperationSheetsWithDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateCommonOperationSheetsWithDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateCommonOperationSheetsWithDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductVersionArtifactTagRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteProductVersionArtifactTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductVersionArtifactTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationBranchesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateApplicationBranchesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationBranchesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationSheetApprovalRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateOperationSheetApprovalRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationSheetApprovalRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateTenantsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateTenantsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateTenantsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationPackageCcDeclareRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetApplicationPackageCcDeclareRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationPackageCcDeclareRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanApplicationRuntimeNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePlanApplicationRuntimeNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanApplicationRuntimeNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialMachineSSHInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetMaterialMachineSSHInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialMachineSSHInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSiteOperationSheetRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetSiteOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSiteOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanApplicationInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePlanApplicationInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanApplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSystemSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSystemSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSystemSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanSitesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListPlanSitesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanSitesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSolutionVersionDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSolutionVersionDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSolutionVersionDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialSyncTaskRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateMaterialSyncTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialSyncTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanCustomConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePlanCustomConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanCustomConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationPipelineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOperationPipelineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationPipelineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductInstanceStateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateProductInstanceStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductInstanceStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationPackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialCommonImportDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateMaterialCommonImportDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialCommonImportDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCommonOperationSheetWithDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCommonOperationSheetWithDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCommonOperationSheetWithDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTenantsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListTenantsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTenantsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialSyncTasksRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListMaterialSyncTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialSyncTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationArtifactChartHistoriesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateApplicationArtifactChartHistoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationArtifactChartHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationSheetRelationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperationSheetRelationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationSheetRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePlanSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationArtifactResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationArtifactResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationArtifactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanProductInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPlanProductInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanProductInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductVersionArtifactTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductVersionArtifactTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductVersionArtifactTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTenantRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateTenantRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTenantRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateCommonOperationSheetsWithDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateCommonOperationSheetsWithDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateCommonOperationSheetsWithDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialMachineSSHInfosRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateMaterialMachineSSHInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialMachineSSHInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialCommonImportDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMaterialCommonImportDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialCommonImportDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationSheetTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOperationSheetTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationSheetTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTenantResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTenantResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTenantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductVersionsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProductVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanApplicationRuntimeNameRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdatePlanApplicationRuntimeNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanApplicationRuntimeNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanHistoryRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeletePlanHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectSiteInfoPlansResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProjectSiteInfoPlansResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectSiteInfoPlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationArtifactRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateApplicationArtifactRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationArtifactRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationInstanceStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApplicationInstanceStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationInstanceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteApplicationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCommonOperationSheetsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListCommonOperationSheetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCommonOperationSheetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationSheetLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOperationSheetLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationSheetLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationRunningDataCollectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationRunningDataCollectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationRunningDataCollectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationPackagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationPackagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationTemplatesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateOperationTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanSubsystemInstanceRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdatePlanSubsystemInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanSubsystemInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductOperationJobRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateProductOperationJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductOperationJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationArtifactChartHistoryRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteApplicationArtifactChartHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationArtifactChartHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductInstanceStateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteProductInstanceStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductInstanceStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateSiteOperationSheetsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateSiteOperationSheetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateSiteOperationSheetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCommonOperationSheetRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateCommonOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCommonOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetPlanDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectTopologyRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProjectTopologyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectTopologyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialSiteRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateMaterialSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductDictionaryRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteProductDictionaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductDictionaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationInstanceDeployVersionsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateApplicationInstanceDeployVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationInstanceDeployVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationRunningPlanRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteApplicationRunningPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationRunningPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSolutionTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSolutionTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSolutionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialSyncTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMaterialSyncTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialSyncTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanApplicationRuntimeNamesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreatePlanApplicationRuntimeNamesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanApplicationRuntimeNamesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductVersionApplicationPackageRelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductVersionApplicationPackageRelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductVersionApplicationPackageRelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSiteProductVersionRelRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetSiteProductVersionRelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSiteProductVersionRelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductOperationJobsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProductOperationJobsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductOperationJobsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOrchTemplateLabelsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateOrchTemplateLabelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOrchTemplateLabelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanGlobalConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPlanGlobalConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanGlobalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductInstanceStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProductInstanceStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductInstanceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationPackageCcDeclaresRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateApplicationPackageCcDeclaresRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationPackageCcDeclaresRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationSheetRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationSheetAttachmentRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetOperationSheetAttachmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationSheetAttachmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductDictionaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProductDictionaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductDictionaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProjectSiteInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProjectSiteInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProjectSiteInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProjectRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateProjectRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProjectRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanApplicationRuntimeNameRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreatePlanApplicationRuntimeNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanApplicationRuntimeNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialMachineSSHInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMaterialMachineSSHInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialMachineSSHInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationSheetTemplatesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListOperationSheetTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationSheetTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductTemplateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateProductTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationInstanceBackupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationInstanceBackupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationInstanceBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectTopologyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProjectTopologyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectTopologyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationPipelineRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateOperationPipelineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationPipelineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationInstanceStatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationInstanceStatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationInstanceStatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSiteDeployTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSiteDeployTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSiteDeployTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialCleanupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMaterialCleanupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialCleanupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationSheetTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperationSheetTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationSheetTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductInstanceBackupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductInstanceBackupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductInstanceBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductOperationJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductOperationJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductOperationJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationSheetApprovalRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOperationSheetApprovalRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationSheetApprovalRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanSubsystemInstanceRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetPlanSubsystemInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanSubsystemInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductVersionApplicationPackageRelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProductVersionApplicationPackageRelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductVersionApplicationPackageRelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialCleanupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateMaterialCleanupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialCleanupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCommonOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCommonOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCommonOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrchTemplateLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOrchTemplateLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrchTemplateLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOrchTemplateLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOrchTemplateLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOrchTemplateLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialSiteDetailsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListMaterialSiteDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialSiteDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanGlobalConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPlanGlobalConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanGlobalConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOtaCertificatesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateOtaCertificatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOtaCertificatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateSystemSettingsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateSystemSettingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateSystemSettingsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialImportTasksRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateMaterialImportTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialImportTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductVersionApplicationPackageRelRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateProductVersionApplicationPackageRelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductVersionApplicationPackageRelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTenantDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTenantDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTenantDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationArtifactRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteApplicationArtifactRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationArtifactRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialCleanupsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateMaterialCleanupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialCleanupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductInstanceOperationSheetsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateProductInstanceOperationSheetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductInstanceOperationSheetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProjectSiteInfosRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateProjectSiteInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProjectSiteInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTenantDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetTenantDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTenantDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationSheetLabelsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListOperationSheetLabelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationSheetLabelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanCustomConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreatePlanCustomConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanCustomConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationPackageCcDeclareRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteApplicationPackageCcDeclareRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationPackageCcDeclareRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductInstanceBackupsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateProductInstanceBackupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductInstanceBackupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProjectRegionsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateProjectRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProjectRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductInstanceOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductInstanceOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductInstanceOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSolutionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSolutionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSolutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationRunningDataCollectionsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListApplicationRunningDataCollectionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationRunningDataCollectionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSolutionVersionsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListSolutionVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSolutionVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductVersionApplicationPackageRelDetailsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProductVersionApplicationPackageRelDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductVersionApplicationPackageRelDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductVersionArtifactTagsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProductVersionArtifactTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductVersionArtifactTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateMaterialInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialApplicationImportDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateMaterialApplicationImportDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialApplicationImportDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductTemplateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateProductTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanApplicationInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePlanApplicationInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanApplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanCustomConfigRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetPlanCustomConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanCustomConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductVersionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProductVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialCustomizedTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaterialCustomizedTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialCustomizedTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialApplicationImportDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMaterialApplicationImportDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialApplicationImportDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationPackageRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateApplicationPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationPackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialMachineSSHInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateMaterialMachineSSHInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialMachineSSHInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationCollectStatusRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteApplicationCollectStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationCollectStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanSitesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPlanSitesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanSitesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialCommonImportDetailsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListMaterialCommonImportDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialCommonImportDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanSiteRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreatePlanSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationBuiltinDataSchemaRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetOperationBuiltinDataSchemaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationBuiltinDataSchemaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationCollectStatusesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationCollectStatusesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationCollectStatusesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationInstanceStatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateApplicationInstanceStatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationInstanceStatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaterialInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationSheetTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOperationSheetTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationSheetTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSolutionVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSolutionVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSolutionVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialMetadatasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialMetadatasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialMetadatasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialMachineSSHInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteMaterialMachineSSHInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialMachineSSHInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialCommonImportDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialCommonImportDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialCommonImportDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductOperationRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateProductOperationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductOperationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationTemplateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetOperationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationCollectStatusesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListApplicationCollectStatusesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationCollectStatusesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSiteProductVersionRelRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateSiteProductVersionRelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSiteProductVersionRelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanApplicationServerRelationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreatePlanApplicationServerRelationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanApplicationServerRelationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationBranchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationBranchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationBranchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationInstanceStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationInstanceStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationInstanceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductOperationJobsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateProductOperationJobsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductOperationJobsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateSolutionVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateSolutionVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateSolutionVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSolutionTemplateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateSolutionTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSolutionTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSolutionTemplateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetSolutionTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSolutionTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationPackagesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListApplicationPackagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationPackagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationInstanceDeployVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateApplicationInstanceDeployVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationInstanceDeployVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialImportTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialImportTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialImportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialCommonImportDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMaterialCommonImportDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialCommonImportDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanSubsystemInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPlanSubsystemInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanSubsystemInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSolutionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetSolutionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSolutionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationInstanceBackupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateApplicationInstanceBackupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationInstanceBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateApplicationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationSheetTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateOperationSheetTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationSheetTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectZoneRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteProjectZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectSiteInfoDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProjectSiteInfoDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectSiteInfoDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductVersionApplicationPackageRelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateProductVersionApplicationPackageRelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductVersionApplicationPackageRelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanSubsystemInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePlanSubsystemInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanSubsystemInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSiteOperationSheetsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListSiteOperationSheetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSiteOperationSheetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationSheetTemplatesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateOperationSheetTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationSheetTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductInstanceStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductInstanceStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductInstanceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSystemSettingRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateSystemSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSystemSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationSheetApprovalRecordRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteOperationSheetApprovalRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationSheetApprovalRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialCommonImportDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateMaterialCommonImportDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialCommonImportDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateSiteDeployTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateSiteDeployTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateSiteDeployTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationPipelineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperationPipelineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationPipelineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanTagRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreatePlanTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProjectRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationBuiltinDataSchemasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateOperationBuiltinDataSchemasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationBuiltinDataSchemasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectTopologyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProjectTopologyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectTopologyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateProductTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductDictionaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductDictionaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductDictionaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePlanHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanSubsystemInstanceRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreatePlanSubsystemInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanSubsystemInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationArtifactResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationArtifactResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationArtifactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationSheetLabelRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteOperationSheetLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationSheetLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationBranchesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListApplicationBranchesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationBranchesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductVersionArtifactTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProductVersionArtifactTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductVersionArtifactTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanSitesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreatePlanSitesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanSitesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationSheetRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanProductInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePlanProductInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanProductInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialMetadataDetailsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListMaterialMetadataDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialMetadataDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductInstanceStatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductInstanceStatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductInstanceStatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProductVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationSheetLabelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateOperationSheetLabelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationSheetLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreatePlanTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlansRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreatePlansRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlansRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDagNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDagNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDagNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialCleanupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaterialCleanupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialCleanupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanApplicationServerRelationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePlanApplicationServerRelationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanApplicationServerRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanServerInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreatePlanServerInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanServerInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductInstanceOperationSheetRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateProductInstanceOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductInstanceOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationBuiltinDataSchemasRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListOperationBuiltinDataSchemasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationBuiltinDataSchemasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationInstanceBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApplicationInstanceBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationInstanceBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSiteProductVersionRelRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateSiteProductVersionRelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSiteProductVersionRelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProjectRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProjectRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProjectRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialInfosRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListMaterialInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialCommonImportDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateMaterialCommonImportDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialCommonImportDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationPackageCcDeclareResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationPackageCcDeclareResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationPackageCcDeclareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSiteDeployTasksRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListSiteDeployTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSiteDeployTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSiteOperationSheetRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteSiteOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSiteOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductOperationRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProductOperationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductOperationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFlywaySchemaHistoryRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetFlywaySchemaHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFlywaySchemaHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSiteProductVersionRelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSiteProductVersionRelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSiteProductVersionRelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDagNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListDagNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListDagNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOperationTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectSiteInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProjectSiteInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectSiteInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductInstanceStatesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateProductInstanceStatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductInstanceStatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductVersionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateProductVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationSheetLabelRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetOperationSheetLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationSheetLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanProductInstanceRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetPlanProductInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanProductInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialMetadatasRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListMaterialMetadatasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialMetadatasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrchTemplateLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOrchTemplateLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOrchTemplateLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrchTemplateLabelRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateOrchTemplateLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrchTemplateLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationInstanceStateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetApplicationInstanceStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationInstanceStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationArtifactChartHistoriesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListApplicationArtifactChartHistoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationArtifactChartHistoriesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanCustomConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePlanCustomConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanCustomConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanApplicationServerRelationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPlanApplicationServerRelationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanApplicationServerRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationPackageCcDeclaresResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateApplicationPackageCcDeclaresResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationPackageCcDeclaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationSheetAttachmentsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateOperationSheetAttachmentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationSheetAttachmentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationCollectStatusRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateApplicationCollectStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationCollectStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanTagRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeletePlanTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDagNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDagNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDagNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationSheetLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperationSheetLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationSheetLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationInstanceStateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteApplicationInstanceStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationInstanceStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanApplicationServerRelationRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeletePlanApplicationServerRelationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanApplicationServerRelationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectSiteInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProjectSiteInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectSiteInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTenantRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetTenantRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTenantRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationInstanceStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationInstanceStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationInstanceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationPackageRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetApplicationPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationPackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductOperationRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateProductOperationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductOperationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationPackagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateApplicationPackagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductVersionArtifactTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductVersionArtifactTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductVersionArtifactTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationPackageDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationPackageDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationPackageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationSheetAttachmentRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateOperationSheetAttachmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationSheetAttachmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectSiteInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteProjectSiteInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectSiteInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductInstanceOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductInstanceOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductInstanceOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationRunningPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApplicationRunningPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationRunningPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateCommonOperationSheetsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateCommonOperationSheetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateCommonOperationSheetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationRunningDataCollectionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationRunningDataCollectionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationRunningDataCollectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPlanTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectZoneRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProjectZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationSheetApprovalRecordsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateOperationSheetApprovalRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationSheetApprovalRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCommonOperationSheetWithDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateCommonOperationSheetWithDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCommonOperationSheetWithDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCommonOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCommonOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCommonOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialSyncTaskRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteMaterialSyncTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialSyncTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialCommonImportDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMaterialCommonImportDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialCommonImportDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationPipelineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOperationPipelineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationPipelineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductDictionaryRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateProductDictionaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductDictionaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanSiteRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdatePlanSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationRunningDataCollectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationRunningDataCollectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationRunningDataCollectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialImportTaskRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteMaterialImportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialImportTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSiteProductVersionRelRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteSiteProductVersionRelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSiteProductVersionRelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationSheetRelationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateOperationSheetRelationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationSheetRelationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationSheetLabelRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateOperationSheetLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationSheetLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductInstanceOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductInstanceOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductInstanceOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductInstanceOperationSheetRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProductInstanceOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductInstanceOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTenantsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListTenantsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListTenantsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSolutionVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSolutionVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSolutionVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialSyncTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateMaterialSyncTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialSyncTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanProductInstanceRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeletePlanProductInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanProductInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectTopologyRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteProjectTopologyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectTopologyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationPackageCcDeclaresRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListApplicationPackageCcDeclaresRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationPackageCcDeclaresRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationPipelinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOperationPipelinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationPipelinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialImportTasksRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListMaterialImportTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialImportTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAtomicRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateAtomicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAtomicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialCleanupRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateMaterialCleanupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialCleanupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetApplicationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProjectRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialSyncTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMaterialSyncTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialSyncTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateFlywaySchemaHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateFlywaySchemaHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateFlywaySchemaHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSolutionTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSolutionTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSolutionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanServerInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreatePlanServerInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanServerInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMaterialInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationSheetTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperationSheetTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationSheetTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanGlobalConfigRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeletePlanGlobalConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanGlobalConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProjectZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateApplicationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductDictionaryRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateProductDictionaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductDictionaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSiteProductVersionRelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSiteProductVersionRelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSiteProductVersionRelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialCleanupRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetMaterialCleanupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialCleanupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationBuiltinDataSchemaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperationBuiltinDataSchemaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationBuiltinDataSchemaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationRunningPlansRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateApplicationRunningPlansRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationRunningPlansRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductOperationsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateProductOperationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductOperationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationTemplateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateOperationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectSiteInfoTopologyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProjectSiteInfoTopologyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectSiteInfoTopologyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationInstanceStateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateApplicationInstanceStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationInstanceStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationRunningDataCollectionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateApplicationRunningDataCollectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationRunningDataCollectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCommonOperationSheetRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteCommonOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCommonOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialSitesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialSitesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialSitesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationSheetRelationsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListOperationSheetRelationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationSheetRelationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialApplicationImportDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateMaterialApplicationImportDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialApplicationImportDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialApplicationImportDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaterialApplicationImportDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialApplicationImportDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProjectRegionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateProjectRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProjectRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSiteOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSiteOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSiteOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationRunningPlanRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetApplicationRunningPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationRunningPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationBuiltinDataSchemasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOperationBuiltinDataSchemasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationBuiltinDataSchemasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrchTemplateLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrchTemplateLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOrchTemplateLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSiteDeployTaskRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateSiteDeployTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSiteDeployTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanApplicationRuntimeNamesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPlanApplicationRuntimeNamesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanApplicationRuntimeNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanProductInstancesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListPlanProductInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanProductInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanCustomConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePlanCustomConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanCustomConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationSheetLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOperationSheetLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationSheetLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanApplicationRuntimeNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPlanApplicationRuntimeNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanApplicationRuntimeNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationPackageDetailsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListApplicationPackageDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationPackageDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListApplicationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSolutionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSolutionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSolutionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSystemSettingsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListSystemSettingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSystemSettingsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialMetadatasRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateMaterialMetadatasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialMetadatasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductInstanceBackupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateProductInstanceBackupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductInstanceBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanCustomConfigRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreatePlanCustomConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanCustomConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialCleanupRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteMaterialCleanupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialCleanupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteProductInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSiteDeployTaskRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateSiteDeployTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSiteDeployTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductInstanceBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProductInstanceBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductInstanceBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateSolutionTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateSolutionTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateSolutionTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectRegionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateProjectRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanProductInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPlanProductInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanProductInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialCustomizedTypeRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateMaterialCustomizedTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialCustomizedTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlansResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreatePlansResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationInstanceDeployVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationInstanceDeployVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationInstanceDeployVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOtaCertificateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteOtaCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOtaCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductInstanceOperationSheetRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateProductInstanceOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductInstanceOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanApplicationServerRelationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePlanApplicationServerRelationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanApplicationServerRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationSheetAttachmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOperationSheetAttachmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationSheetAttachmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateProductInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductInstanceStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductInstanceStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductInstanceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductOperationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductOperationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductOperationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationBranchRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetApplicationBranchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationBranchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectRegionDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProjectRegionDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectRegionDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListFlywaySchemaHistoriesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListFlywaySchemaHistoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListFlywaySchemaHistoriesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCommonOperationSheetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCommonOperationSheetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCommonOperationSheetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductVersionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateProductVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationInstanceBackupsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateApplicationInstanceBackupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationInstanceBackupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialCommonImportDetailsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateMaterialCommonImportDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialCommonImportDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialMachineSSHInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaterialMachineSSHInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialMachineSSHInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductDictionaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductDictionaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductDictionaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProjectSiteInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateProjectSiteInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProjectSiteInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectSiteInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateProjectSiteInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectSiteInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialCommonImportDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteMaterialCommonImportDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialCommonImportDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialApplicationImportDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetMaterialApplicationImportDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialApplicationImportDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanTagsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreatePlanTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanSubsystemInstanceRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeletePlanSubsystemInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanSubsystemInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialCustomizedTypesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListMaterialCustomizedTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialCustomizedTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationCollectStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationCollectStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationCollectStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOtaCertificatesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListOtaCertificatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOtaCertificatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanHistoriesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListPlanHistoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanHistoriesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProjectTopologyRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateProjectTopologyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProjectTopologyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectSiteInfoDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProjectSiteInfoDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectSiteInfoDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationInstanceStateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateApplicationInstanceStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationInstanceStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductInstanceStatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateProductInstanceStatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductInstanceStatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProjectTopologiesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateProjectTopologiesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProjectTopologiesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFlywaySchemaHistoryRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateFlywaySchemaHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFlywaySchemaHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationSheetRelationRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteOperationSheetRelationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationSheetRelationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProductInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationInstanceDeployVersionsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListApplicationInstanceDeployVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationInstanceDeployVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCommonOperationSheetsWithDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListCommonOperationSheetsWithDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCommonOperationSheetsWithDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateSiteProductVersionRelsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateSiteProductVersionRelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateSiteProductVersionRelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSolutionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSolutionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSolutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateApplicationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDagNodeRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateDagNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateDagNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProductInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationBranchesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateApplicationBranchesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationBranchesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialApplicationImportDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMaterialApplicationImportDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialApplicationImportDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialSyncTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaterialSyncTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialSyncTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationInstanceBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationInstanceBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationInstanceBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanApplicationServerRelationsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListPlanApplicationServerRelationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanApplicationServerRelationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductVersionArtifactTagRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateProductVersionArtifactTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductVersionArtifactTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectSiteInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProjectSiteInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectSiteInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOrchTemplateLabelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOrchTemplateLabelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOrchTemplateLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationPackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApplicationPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProjectTopologyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProjectTopologyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProjectTopologyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanProductInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePlanProductInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanProductInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectSiteInfoDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProjectSiteInfoDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectSiteInfoDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSystemSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSystemSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSystemSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationSheetAttachmentRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteOperationSheetAttachmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationSheetAttachmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationSheetRelationRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateOperationSheetRelationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationSheetRelationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanApplicationInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPlanApplicationInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanApplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialMachineSSHInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMaterialMachineSSHInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialMachineSSHInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSiteProductVersionRelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSiteProductVersionRelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSiteProductVersionRelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationArtifactRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateApplicationArtifactRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationArtifactRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialApplicationImportDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteMaterialApplicationImportDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialApplicationImportDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMaterialInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanSubsystemInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPlanSubsystemInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanSubsystemInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAtomicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAtomicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAtomicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSolutionVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSolutionVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSolutionVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationBuiltinDataSchemaRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteOperationBuiltinDataSchemaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationBuiltinDataSchemaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationPipelineRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetOperationPipelineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationPipelineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOtaCertificatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOtaCertificatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOtaCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectTopologiesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProjectTopologiesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectTopologiesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPlanDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTenantDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTenantDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTenantDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductVersionArtifactTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateProductVersionArtifactTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductVersionArtifactTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationInstanceBackupRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetApplicationInstanceBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationInstanceBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSystemSettingResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSystemSettingResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSystemSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationBuiltinDataSchemaRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateOperationBuiltinDataSchemaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationBuiltinDataSchemaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrchTemplateLabelRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteOrchTemplateLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOrchTemplateLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFlywaySchemaHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFlywaySchemaHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFlywaySchemaHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialCommonImportDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetMaterialCommonImportDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialCommonImportDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationBuiltinDataSchemaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOperationBuiltinDataSchemaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationBuiltinDataSchemaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationCollectStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApplicationCollectStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationCollectStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateTenantsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateTenantsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateTenantsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationPipelineRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteOperationPipelineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationPipelineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSiteDeployTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSiteDeployTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSiteDeployTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationSheetAttachmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperationSheetAttachmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationSheetAttachmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTenantDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateTenantDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTenantDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialSiteRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteMaterialSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductOperationJobRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProductOperationJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductOperationJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductTemplatesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProductTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationSheetAttachmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOperationSheetAttachmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationSheetAttachmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationBuiltinDataSchemasRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateOperationBuiltinDataSchemasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationBuiltinDataSchemasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPlanHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanServerInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePlanServerInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanServerInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProductTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationRunningDataCollectionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateApplicationRunningDataCollectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationRunningDataCollectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationArtifactsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateApplicationArtifactsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationArtifactsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanApplicationInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreatePlanApplicationInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanApplicationInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateProductInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationInstanceDeployVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationInstanceDeployVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationInstanceDeployVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationTemplatesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListOperationTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateMaterialInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListCommonOperationSheetsWithDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCommonOperationSheetsWithDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListCommonOperationSheetsWithDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationSheetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOperationSheetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationSheetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanGlobalConfigsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListPlanGlobalConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanGlobalConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationCollectStatusRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateApplicationCollectStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationCollectStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationSheetRelationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOperationSheetRelationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationSheetRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductVersionApplicationPackageRelDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductVersionApplicationPackageRelDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductVersionApplicationPackageRelDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductInstanceBackupRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateProductInstanceBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductInstanceBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductDictionariesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateProductDictionariesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductDictionariesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProjectRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectZoneRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateProjectZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSiteDeployTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSiteDeployTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSiteDeployTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationSheetLabelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperationSheetLabelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationSheetLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationRunningDataCollectionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateApplicationRunningDataCollectionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationRunningDataCollectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateSystemSettingsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateSystemSettingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateSystemSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialMachineSSHInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMaterialMachineSSHInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialMachineSSHInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanHistoriesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreatePlanHistoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationPackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanProductInstanceTreeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPlanProductInstanceTreeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanProductInstanceTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationInstanceBackupsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListApplicationInstanceBackupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationInstanceBackupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationSheetApprovalRecordsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListOperationSheetApprovalRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationSheetApprovalRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProjectZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialMetadataRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateMaterialMetadataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialMetadataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlansResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPlansResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanTagRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetPlanTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationPackageCcDeclaresResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationPackageCcDeclaresResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationPackageCcDeclaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialCleanupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialCleanupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialCleanupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialMetadataDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialMetadataDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialMetadataDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialImportTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMaterialImportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialImportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateFlywaySchemaHistoriesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateFlywaySchemaHistoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateFlywaySchemaHistoriesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectRegionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProjectRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSolutionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSolutionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSolutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationTemplateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteOperationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSiteOperationSheetRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateSiteOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSiteOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductInfosRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateProductInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductOperationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductOperationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductOperationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationInstanceDeployVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationInstanceDeployVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationInstanceDeployVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductInstanceBackupsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProductInstanceBackupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductInstanceBackupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationPackageDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationPackageDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationPackageDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductInstanceBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductInstanceBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductInstanceBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialSyncTaskRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateMaterialSyncTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialSyncTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductVersionArtifactTagsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateProductVersionArtifactTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductVersionArtifactTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanApplicationRuntimeNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePlanApplicationRuntimeNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanApplicationRuntimeNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductInstanceStateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateProductInstanceStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductInstanceStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductInstanceBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductInstanceBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductInstanceBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateFlywaySchemaHistoryRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateFlywaySchemaHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateFlywaySchemaHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialApplicationImportDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialApplicationImportDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialApplicationImportDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProjectZonesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateProjectZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProjectZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanGlobalConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePlanGlobalConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanGlobalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDagNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDagNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDagNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialCleanupsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListMaterialCleanupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialCleanupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectRegionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteProjectRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductVersionsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateProductVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProjectTopologiesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateProjectTopologiesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProjectTopologiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateSolutionsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateSolutionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateSolutionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanHistoryRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreatePlanHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListFlywaySchemaHistoriesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListFlywaySchemaHistoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListFlywaySchemaHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductOperationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductOperationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductOperationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSolutionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateSolutionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSolutionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationArtifactChartHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationArtifactChartHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationArtifactChartHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationInstanceDeployVersionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetApplicationInstanceDeployVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationInstanceDeployVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationRunningDataCollectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationRunningDataCollectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationRunningDataCollectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationPipelinesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListOperationPipelinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationPipelinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSiteProductVersionRelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSiteProductVersionRelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSiteProductVersionRelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationPackageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationPackageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSiteOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSiteOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSiteOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationInstanceBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationInstanceBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationInstanceBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialSiteDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialSiteDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialSiteDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialMachineSSHInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateMaterialMachineSSHInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialMachineSSHInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialSyncTasksRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateMaterialSyncTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialSyncTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanGlobalConfigsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreatePlanGlobalConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanGlobalConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationSheetRelationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOperationSheetRelationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationSheetRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationBranchesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationBranchesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationBranchesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAtomicsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListAtomicsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAtomicsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanServerInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdatePlanServerInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanServerInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAtomicRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetAtomicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAtomicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialSitesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListMaterialSitesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialSitesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanProductInstanceRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdatePlanProductInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanProductInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationArtifactsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateApplicationArtifactsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationArtifactsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOrchTemplateLabelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateOrchTemplateLabelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOrchTemplateLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdatePlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSolutionTemplateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateSolutionTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSolutionTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationArtifactRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetApplicationArtifactRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationArtifactRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanHistoryRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdatePlanHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanApplicationServerRelationRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetPlanApplicationServerRelationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanApplicationServerRelationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationRunningPlansResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationRunningPlansResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationRunningPlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectZonesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProjectZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationBuiltinDataSchemaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOperationBuiltinDataSchemaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationBuiltinDataSchemaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductInstanceBackupRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteProductInstanceBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductInstanceBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationArtifactChartHistoryRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetApplicationArtifactChartHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationArtifactChartHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationSheetLabelRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateOperationSheetLabelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationSheetLabelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAtomicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAtomicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAtomicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteMaterialInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationSheetRelationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOperationSheetRelationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationSheetRelationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCommonOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCommonOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateCommonOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductVersionArtifactTagRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateProductVersionArtifactTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductVersionArtifactTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanServerInfosRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListPlanServerInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanServerInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateSiteDeployTasksRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateSiteDeployTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateSiteDeployTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateApplicationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePlanTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectRegionDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProjectRegionDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectRegionDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationSheetAttachmentsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListOperationSheetAttachmentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationSheetAttachmentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationInstanceBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationInstanceBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationInstanceBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDagNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDagNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDagNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSolutionsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListSolutionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSolutionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationPackageCcDeclareRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateApplicationPackageCcDeclareRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationPackageCcDeclareRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialImportTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMaterialImportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialImportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanGlobalConfigRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetPlanGlobalConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanGlobalConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanProductInstanceTreeRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetPlanProductInstanceTreeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanProductInstanceTreeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialSyncTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialSyncTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialSyncTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSystemSettingRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateSystemSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSystemSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductOperationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProductOperationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductOperationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSiteOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSiteOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSiteOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationArtifactChartHistoryRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateApplicationArtifactChartHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationArtifactChartHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationPipelineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperationPipelineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationPipelineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialMachineSSHInfosRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListMaterialMachineSSHInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialMachineSSHInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationPackageCcDeclareResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApplicationPackageCcDeclareResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationPackageCcDeclareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSolutionVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSolutionVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSolutionVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialMachineSSHInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialMachineSSHInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialMachineSSHInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationSheetTemplateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateOperationSheetTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationSheetTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSolutionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateSolutionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSolutionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanApplicationServerRelationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePlanApplicationServerRelationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanApplicationServerRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMaterialInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTenantResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateTenantResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTenantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanSiteRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeletePlanSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOtaCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOtaCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOtaCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectRegionsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProjectRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialCommonImportDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaterialCommonImportDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialCommonImportDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCommonOperationSheetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCommonOperationSheetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCommonOperationSheetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductVersionArtifactTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductVersionArtifactTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductVersionArtifactTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationSheetsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListOperationSheetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationSheetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanTagsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListPlanTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationInstanceDeployVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationInstanceDeployVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationInstanceDeployVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateProductInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOtaCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOtaCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOtaCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSolutionTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSolutionTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSolutionTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOperationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOtaCertificateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateOtaCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOtaCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateAtomicsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateAtomicsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateAtomicsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationBranchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationBranchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationBranchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationSheetTemplateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateOperationSheetTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationSheetTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSystemSettingRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteSystemSettingRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSystemSettingRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationArtifactChartHistoryRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateApplicationArtifactChartHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationArtifactChartHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanApplicationServerRelationRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdatePlanApplicationServerRelationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanApplicationServerRelationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProjectZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProjectZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProjectZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanSitesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreatePlanSitesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanSitesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCommonOperationSheetRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetCommonOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCommonOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectSiteInfoDetailsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProjectSiteInfoDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectSiteInfoDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOtaCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOtaCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOtaCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationSheetLabelsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateOperationSheetLabelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationSheetLabelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAtomicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAtomicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAtomicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationBuiltinDataSchemaRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateOperationBuiltinDataSchemaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationBuiltinDataSchemaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationArtifactChartHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationArtifactChartHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationArtifactChartHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationRunningPlanRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateApplicationRunningPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationRunningPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDagNodeRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateDagNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDagNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanApplicationRuntimeNamesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListPlanApplicationRuntimeNamesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanApplicationRuntimeNamesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSolutionTemplatesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListSolutionTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSolutionTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSystemSettingsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSystemSettingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSystemSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanGlobalConfigRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdatePlanGlobalConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanGlobalConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialCustomizedTypeRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateMaterialCustomizedTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialCustomizedTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectTopologyDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateProjectTopologyDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectTopologyDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanCustomConfigRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeletePlanCustomConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanCustomConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanApplicationInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPlanApplicationInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanApplicationInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanGlobalConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePlanGlobalConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanGlobalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSiteOperationSheetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSiteOperationSheetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSiteOperationSheetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationCollectStatusesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateApplicationCollectStatusesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationCollectStatusesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanSubsystemInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreatePlanSubsystemInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanSubsystemInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductInstanceBackupRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateProductInstanceBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductInstanceBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialImportTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaterialImportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialImportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePlanHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanProductInstanceRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreatePlanProductInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanProductInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationRunningPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationRunningPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationRunningPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductDictionaryRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProductDictionaryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductDictionaryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductOperationJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductOperationJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductOperationJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialApplicationImportDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMaterialApplicationImportDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialApplicationImportDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanApplicationServerRelationsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreatePlanApplicationServerRelationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanApplicationServerRelationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaterialSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOrchTemplateLabelsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListOrchTemplateLabelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOrchTemplateLabelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialCleanupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMaterialCleanupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialCleanupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTenantResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTenantResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTenantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationPackageCcDeclareRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateApplicationPackageCcDeclareRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationPackageCcDeclareRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationSheetApprovalRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOperationSheetApprovalRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationSheetApprovalRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanSubsystemInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePlanSubsystemInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanSubsystemInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialApplicationImportDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateMaterialApplicationImportDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialApplicationImportDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialMetadatasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateMaterialMetadatasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialMetadatasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductVersionApplicationPackageRelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductVersionApplicationPackageRelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductVersionApplicationPackageRelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationInstanceDeployVersionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateApplicationInstanceDeployVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationInstanceDeployVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanHistoriesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreatePlanHistoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanHistoriesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductInfosRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProductInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationBuiltinDataSchemaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOperationBuiltinDataSchemaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationBuiltinDataSchemaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSolutionVersionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetSolutionVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSolutionVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSiteProductVersionRelsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListSiteProductVersionRelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSiteProductVersionRelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialApplicationImportDetailsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateMaterialApplicationImportDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialApplicationImportDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationPackageRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateApplicationPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationPackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectTopologiesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProjectTopologiesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectTopologiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationSheetApprovalRecordRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetOperationSheetApprovalRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationSheetApprovalRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationSheetRelationRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetOperationSheetRelationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationSheetRelationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPlanSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialMetadataRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateMaterialMetadataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialMetadataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialSyncTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMaterialSyncTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialSyncTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanHistoryRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetPlanHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductInstanceOperationSheetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateProductInstanceOperationSheetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductInstanceOperationSheetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialSyncTaskRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetMaterialSyncTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialSyncTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSolutionVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSolutionVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSolutionVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateFlywaySchemaHistoriesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateFlywaySchemaHistoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateFlywaySchemaHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationPipelinesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateOperationPipelinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationPipelinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialImportTaskRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetMaterialImportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialImportTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationSheetLabelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOperationSheetLabelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationSheetLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanProductInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePlanProductInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanProductInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePlanSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationRunningDataCollectionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetApplicationRunningDataCollectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationRunningDataCollectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialSiteRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateMaterialSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationSheetApprovalRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperationSheetApprovalRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationSheetApprovalRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDagNodeRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetDagNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDagNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSiteDeployTaskRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteSiteDeployTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSiteDeployTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialSiteRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetMaterialSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanCustomConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPlanCustomConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanCustomConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialCleanupRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateMaterialCleanupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialCleanupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanGlobalConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreatePlanGlobalConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanGlobalConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateSiteProductVersionRelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateSiteProductVersionRelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateSiteProductVersionRelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationSheetRelationsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateOperationSheetRelationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationSheetRelationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAtomicRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateAtomicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAtomicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateOperationTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateProductVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCommonOperationSheetWithDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCommonOperationSheetWithDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCommonOperationSheetWithDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductDictionaryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductDictionaryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductDictionaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAtomicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAtomicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAtomicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetMaterialInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanServerInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePlanServerInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanServerInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateDagNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateDagNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateDagNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductOperationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateProductOperationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductOperationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFlywaySchemaHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFlywaySchemaHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFlywaySchemaHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationSheetRelationRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateOperationSheetRelationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationSheetRelationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanServerInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePlanServerInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanServerInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationPackageCcDeclareResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationPackageCcDeclareResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationPackageCcDeclareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationSheetRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationBranchRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateApplicationBranchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationBranchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanApplicationInstanceRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdatePlanApplicationInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanApplicationInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectSiteInfosRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProjectSiteInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectSiteInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialImportTaskRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateMaterialImportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialImportTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationCollectStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationCollectStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationCollectStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanApplicationInstanceRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreatePlanApplicationInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanApplicationInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationPackageCcDeclareResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetApplicationPackageCcDeclareResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationPackageCcDeclareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanHistoriesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPlanHistoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProjectSiteInfoPlansRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProjectSiteInfoPlansRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProjectSiteInfoPlansRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductOperationJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProductOperationJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductOperationJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanSubsystemInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePlanSubsystemInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanSubsystemInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTenantRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateTenantRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTenantRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationSheetApprovalRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOperationSheetApprovalRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationSheetApprovalRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAtomicsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAtomicsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAtomicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationArtifactsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListApplicationArtifactsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationArtifactsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductDictionariesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateProductDictionariesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductDictionariesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductOperationJobsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateProductOperationJobsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductOperationJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanApplicationServerRelationRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreatePlanApplicationServerRelationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanApplicationServerRelationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductOperationRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteProductOperationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductOperationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMaterialSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductInstanceStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductInstanceStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductInstanceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialSitesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateMaterialSitesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialSitesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSolutionTemplateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteSolutionTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSolutionTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMaterialCustomizedTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateMaterialCustomizedTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMaterialCustomizedTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanApplicationRuntimeNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePlanApplicationRuntimeNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanApplicationRuntimeNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationSheetRelationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOperationSheetRelationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationSheetRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductVersionArtifactTagRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProductVersionArtifactTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductVersionArtifactTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProjectZonesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateProjectZonesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProjectZonesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFlywaySchemaHistoryRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteFlywaySchemaHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFlywaySchemaHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanServerInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetPlanServerInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanServerInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialMachineSSHInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateMaterialMachineSSHInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialMachineSSHInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePlanTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAtomicRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteAtomicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAtomicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanGlobalConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePlanGlobalConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanGlobalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationCollectStatusRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetApplicationCollectStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationCollectStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialCustomizedTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMaterialCustomizedTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialCustomizedTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProjectSiteInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateProjectSiteInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProjectSiteInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectTopologyDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProjectTopologyDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectTopologyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductInstanceStatesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProductInstanceStatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductInstanceStatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOtaCertificateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateOtaCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOtaCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductOperationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductOperationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductOperationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductVersionApplicationPackageRelRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProductVersionApplicationPackageRelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductVersionApplicationPackageRelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductInstanceOperationSheetsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProductInstanceOperationSheetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductInstanceOperationSheetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationArtifactChartHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApplicationArtifactChartHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationArtifactChartHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanProductInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreatePlanProductInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanProductInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOperationPipelineRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateOperationPipelineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationPipelineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanProductInstancesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreatePlanProductInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanProductInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialMetadataRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteMaterialMetadataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialMetadataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductOperationsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProductOperationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductOperationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanSubsystemInstancesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreatePlanSubsystemInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanSubsystemInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialSiteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMaterialSiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialSiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductOperationJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductOperationJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductOperationJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProductVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductTemplateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteProductTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateCommonOperationSheetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateCommonOperationSheetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateCommonOperationSheetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSiteOperationSheetRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateSiteOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSiteOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlansRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListPlansRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlansRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialCustomizedTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMaterialCustomizedTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialCustomizedTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialMetadataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMaterialMetadataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialMetadataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationRunningDataCollectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApplicationRunningDataCollectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationRunningDataCollectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialMetadataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMaterialMetadataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialMetadataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationSheetTemplateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetOperationSheetTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationSheetTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanApplicationServerRelationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPlanApplicationServerRelationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanApplicationServerRelationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductVersionApplicationPackageRelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductVersionApplicationPackageRelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductVersionApplicationPackageRelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOtaCertificatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateOtaCertificatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOtaCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductOperationJobRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteProductOperationJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductOperationJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOperationSheetAttachmentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListOperationSheetAttachmentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListOperationSheetAttachmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationSheetTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOperationSheetTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationSheetTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOperationSheetRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateOperationSheetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateOperationSheetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationInstanceDeployVersionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateApplicationInstanceDeployVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationInstanceDeployVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductInstanceBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductInstanceBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductInstanceBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialMetadataRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetMaterialMetadataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialMetadataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProjectTopologyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProjectTopologyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProjectTopologyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSolutionVersionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateSolutionVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSolutionVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationRunningPlansRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListApplicationRunningPlansRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationRunningPlansRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMaterialApplicationImportDetailsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListMaterialApplicationImportDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMaterialApplicationImportDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationArtifactChartHistoriesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateApplicationArtifactChartHistoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationArtifactChartHistoriesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateSolutionTemplatesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateSolutionTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateSolutionTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationBranchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationBranchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApplicationBranchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanGlobalConfigRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreatePlanGlobalConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanGlobalConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationInstanceDeployVersionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteApplicationInstanceDeployVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationInstanceDeployVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationSheetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateOperationSheetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationSheetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSolutionVersionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteSolutionVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSolutionVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTenantRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteTenantRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTenantRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanApplicationRuntimeNamesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreatePlanApplicationRuntimeNamesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanApplicationRuntimeNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanSubsystemInstancesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListPlanSubsystemInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanSubsystemInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectTopologyRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateProjectTopologyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateProjectTopologyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSiteDeployTaskRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetSiteDeployTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSiteDeployTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductVersionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteProductVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductInstanceStateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProductInstanceStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductInstanceStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataEngineData struct {
	// data engine request data

	InternalData *string `json:"InternalData,omitempty" name:"InternalData"`
}

type CreateOperationTemplateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateOperationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateOperationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreatePlanCustomConfigsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreatePlanCustomConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreatePlanCustomConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSiteDeployTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSiteDeployTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetSiteDeployTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductDictionariesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductDictionariesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductDictionariesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOtaCertificateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOtaCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOtaCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMaterialCustomizedTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMaterialCustomizedTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMaterialCustomizedTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationRunningPlanRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateApplicationRunningPlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationRunningPlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOperationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetOperationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOperationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanServerInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPlanServerInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanServerInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialCustomizedTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateMaterialCustomizedTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialCustomizedTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMaterialInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *CreateMaterialInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMaterialInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaterialCustomizedTypeRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteMaterialCustomizedTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMaterialCustomizedTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOperationSheetTemplateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteOperationSheetTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteOperationSheetTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanServerInfoRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeletePlanServerInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanServerInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApplicationPackageDetailRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetApplicationPackageDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApplicationPackageDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSolutionTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSolutionTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSolutionTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePlanCustomConfigRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdatePlanCustomConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePlanCustomConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductDictionariesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListProductDictionariesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductDictionariesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSolutionVersionDetailsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListSolutionVersionDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSolutionVersionDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProjectZoneRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateProjectZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProjectZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanApplicationInstanceRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeletePlanApplicationInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanApplicationInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductVersionApplicationPackageRelRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteProductVersionApplicationPackageRelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProductVersionApplicationPackageRelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetOtaCertificateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetOtaCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetOtaCertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationInstanceStatesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListApplicationInstanceStatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListApplicationInstanceStatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProductInstanceOperationSheetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListProductInstanceOperationSheetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListProductInstanceOperationSheetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationArtifactChartHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationArtifactChartHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationArtifactChartHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationRunningPlansResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateApplicationRunningPlansResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationRunningPlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationSheetAttachmentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateOperationSheetAttachmentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationSheetAttachmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateSolutionVersionsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateSolutionVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateSolutionVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSiteProductVersionRelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSiteProductVersionRelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSiteProductVersionRelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductVersionApplicationPackageRelRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UpdateProductVersionApplicationPackageRelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateProductVersionApplicationPackageRelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationPackagesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateApplicationPackagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationPackagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateProductVersionApplicationPackageRelsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateProductVersionApplicationPackageRelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateProductVersionApplicationPackageRelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSiteDeployTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSiteDeployTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateSiteDeployTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePlanHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePlanHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePlanHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationPackageRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteApplicationPackageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationPackageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationRunningDataCollectionRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeleteApplicationRunningDataCollectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationRunningDataCollectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateApplicationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApplicationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPlanCustomConfigsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *ListPlanCustomConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListPlanCustomConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateSolutionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BulkCreateSolutionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateSolutionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationRunningPlanResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationRunningPlanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteApplicationRunningPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectSiteInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProjectSiteInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteProjectSiteInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductTemplateRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetProductTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPlanSiteRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *GetPlanSiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPlanSiteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateApplicationInstanceStatesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateApplicationInstanceStatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateApplicationInstanceStatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateMaterialCustomizedTypesRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateMaterialCustomizedTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateMaterialCustomizedTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BulkCreateOperationSheetsRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BulkCreateOperationSheetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BulkCreateOperationSheetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeletePlanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePlanApplicationRuntimeNameRequest struct {
	*tchttp.BaseRequest

	// data engine request data

	Data *DataEngineData `json:"Data,omitempty" name:"Data"`
	// service type, fixed to 'DataEngine'

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DeletePlanApplicationRuntimeNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePlanApplicationRuntimeNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
