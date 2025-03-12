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

package v2018109

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2018-10-9"

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

func NewAddNodeRequest() (request *AddNodeRequest) {
	request = &AddNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "AddNode")
	return
}

func NewAddNodeResponse() (response *AddNodeResponse) {
	response = &AddNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增节点
func (c *Client) AddNode(request *AddNodeRequest) (response *AddNodeResponse, err error) {
	if request == nil {
		request = NewAddNodeRequest()
	}
	response = NewAddNodeResponse()
	err = c.Send(request, response)
	return
}

func NewDecideRequest() (request *DecideRequest) {
	request = &DecideRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "Decide")
	return
}

func NewDecideResponse() (response *DecideResponse) {
	response = &DecideResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Decide
func (c *Client) Decide(request *DecideRequest) (response *DecideResponse, err error) {
	if request == nil {
		request = NewDecideRequest()
	}
	response = NewDecideResponse()
	err = c.Send(request, response)
	return
}

func NewPathRequest() (request *PathRequest) {
	request = &PathRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "Path")
	return
}

func NewPathResponse() (response *PathResponse) {
	response = &PathResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Path
func (c *Client) Path(request *PathRequest) (response *PathResponse, err error) {
	if request == nil {
		request = NewPathRequest()
	}
	response = NewPathResponse()
	err = c.Send(request, response)
	return
}

func NewQueryNodePreviewRequest() (request *QueryNodePreviewRequest) {
	request = &QueryNodePreviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "QueryNodePreview")
	return
}

func NewQueryNodePreviewResponse() (response *QueryNodePreviewResponse) {
	response = &QueryNodePreviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询导入节点信息
func (c *Client) QueryNodePreview(request *QueryNodePreviewRequest) (response *QueryNodePreviewResponse, err error) {
	if request == nil {
		request = NewQueryNodePreviewRequest()
	}
	response = NewQueryNodePreviewResponse()
	err = c.Send(request, response)
	return
}

func NewReviewRequest() (request *ReviewRequest) {
	request = &ReviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "Review")
	return
}

func NewReviewResponse() (response *ReviewResponse) {
	response = &ReviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Review
func (c *Client) Review(request *ReviewRequest) (response *ReviewResponse, err error) {
	if request == nil {
		request = NewReviewRequest()
	}
	response = NewReviewResponse()
	err = c.Send(request, response)
	return
}

func NewQueryUpgradeLogRequest() (request *QueryUpgradeLogRequest) {
	request = &QueryUpgradeLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "QueryUpgradeLog")
	return
}

func NewQueryUpgradeLogResponse() (response *QueryUpgradeLogResponse) {
	response = &QueryUpgradeLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 文档系统查询升级记录
func (c *Client) QueryUpgradeLog(request *QueryUpgradeLogRequest) (response *QueryUpgradeLogResponse, err error) {
	if request == nil {
		request = NewQueryUpgradeLogRequest()
	}
	response = NewQueryUpgradeLogResponse()
	err = c.Send(request, response)
	return
}

func NewQueryDirRequest() (request *QueryDirRequest) {
	request = &QueryDirRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "QueryDir")
	return
}

func NewQueryDirResponse() (response *QueryDirResponse) {
	response = &QueryDirResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询类型信息
func (c *Client) QueryDir(request *QueryDirRequest) (response *QueryDirResponse, err error) {
	if request == nil {
		request = NewQueryDirRequest()
	}
	response = NewQueryDirResponse()
	err = c.Send(request, response)
	return
}

func NewConfirmUpgradeRequest() (request *ConfirmUpgradeRequest) {
	request = &ConfirmUpgradeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "ConfirmUpgrade")
	return
}

func NewConfirmUpgradeResponse() (response *ConfirmUpgradeResponse) {
	response = &ConfirmUpgradeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 文档系统确认升级
func (c *Client) ConfirmUpgrade(request *ConfirmUpgradeRequest) (response *ConfirmUpgradeResponse, err error) {
	if request == nil {
		request = NewConfirmUpgradeRequest()
	}
	response = NewConfirmUpgradeResponse()
	err = c.Send(request, response)
	return
}

func NewQueryNodeRequest() (request *QueryNodeRequest) {
	request = &QueryNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "QueryNode")
	return
}

func NewQueryNodeResponse() (response *QueryNodeResponse) {
	response = &QueryNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点
func (c *Client) QueryNode(request *QueryNodeRequest) (response *QueryNodeResponse, err error) {
	if request == nil {
		request = NewQueryNodeRequest()
	}
	response = NewQueryNodeResponse()
	err = c.Send(request, response)
	return
}

func NewQueryTreeUpgradeRequest() (request *QueryTreeUpgradeRequest) {
	request = &QueryTreeUpgradeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "QueryTreeUpgrade")
	return
}

func NewQueryTreeUpgradeResponse() (response *QueryTreeUpgradeResponse) {
	response = &QueryTreeUpgradeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 文档系统查询升级表树型结构
func (c *Client) QueryTreeUpgrade(request *QueryTreeUpgradeRequest) (response *QueryTreeUpgradeResponse, err error) {
	if request == nil {
		request = NewQueryTreeUpgradeRequest()
	}
	response = NewQueryTreeUpgradeResponse()
	err = c.Send(request, response)
	return
}

func NewGiveupUpgradeRequest() (request *GiveupUpgradeRequest) {
	request = &GiveupUpgradeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "GiveupUpgrade")
	return
}

func NewGiveupUpgradeResponse() (response *GiveupUpgradeResponse) {
	response = &GiveupUpgradeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 文档系统放弃升级
func (c *Client) GiveupUpgrade(request *GiveupUpgradeRequest) (response *GiveupUpgradeResponse, err error) {
	if request == nil {
		request = NewGiveupUpgradeRequest()
	}
	response = NewGiveupUpgradeResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCategoriesRequest() (request *QueryCategoriesRequest) {
	request = &QueryCategoriesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "QueryCategories")
	return
}

func NewQueryCategoriesResponse() (response *QueryCategoriesResponse) {
	response = &QueryCategoriesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询类别，返回类别和类别下的产品
func (c *Client) QueryCategories(request *QueryCategoriesRequest) (response *QueryCategoriesResponse, err error) {
	if request == nil {
		request = NewQueryCategoriesRequest()
	}
	response = NewQueryCategoriesResponse()
	err = c.Send(request, response)
	return
}

func NewQueryTreePreviewRequest() (request *QueryTreePreviewRequest) {
	request = &QueryTreePreviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "QueryTreePreview")
	return
}

func NewQueryTreePreviewResponse() (response *QueryTreePreviewResponse) {
	response = &QueryTreePreviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询导入树型结构
func (c *Client) QueryTreePreview(request *QueryTreePreviewRequest) (response *QueryTreePreviewResponse, err error) {
	if request == nil {
		request = NewQueryTreePreviewRequest()
	}
	response = NewQueryTreePreviewResponse()
	err = c.Send(request, response)
	return
}

func NewQueryDocProductCodesRequest() (request *QueryDocProductCodesRequest) {
	request = &QueryDocProductCodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "QueryDocProductCodes")
	return
}

func NewQueryDocProductCodesResponse() (response *QueryDocProductCodesResponse) {
	response = &QueryDocProductCodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询文档产品code
func (c *Client) QueryDocProductCodes(request *QueryDocProductCodesRequest) (response *QueryDocProductCodesResponse, err error) {
	if request == nil {
		request = NewQueryDocProductCodesRequest()
	}
	response = NewQueryDocProductCodesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateUpgradeLogRequest() (request *UpdateUpgradeLogRequest) {
	request = &UpdateUpgradeLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "UpdateUpgradeLog")
	return
}

func NewUpdateUpgradeLogResponse() (response *UpdateUpgradeLogResponse) {
	response = &UpdateUpgradeLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新升级日志记录
func (c *Client) UpdateUpgradeLog(request *UpdateUpgradeLogRequest) (response *UpdateUpgradeLogResponse, err error) {
	if request == nil {
		request = NewUpdateUpgradeLogRequest()
	}
	response = NewUpdateUpgradeLogResponse()
	err = c.Send(request, response)
	return
}

func NewExportDocRequest() (request *ExportDocRequest) {
	request = &ExportDocRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "ExportDoc")
	return
}

func NewExportDocResponse() (response *ExportDocResponse) {
	response = &ExportDocResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出文档
func (c *Client) ExportDoc(request *ExportDocRequest) (response *ExportDocResponse, err error) {
	if request == nil {
		request = NewExportDocRequest()
	}
	response = NewExportDocResponse()
	err = c.Send(request, response)
	return
}

func NewQueryExportLogRequest() (request *QueryExportLogRequest) {
	request = &QueryExportLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "QueryExportLog")
	return
}

func NewQueryExportLogResponse() (response *QueryExportLogResponse) {
	response = &QueryExportLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询导出记录
func (c *Client) QueryExportLog(request *QueryExportLogRequest) (response *QueryExportLogResponse, err error) {
	if request == nil {
		request = NewQueryExportLogRequest()
	}
	response = NewQueryExportLogResponse()
	err = c.Send(request, response)
	return
}

func NewQueryTreeRequest() (request *QueryTreeRequest) {
	request = &QueryTreeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "QueryTree")
	return
}

func NewQueryTreeResponse() (response *QueryTreeResponse) {
	response = &QueryTreeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询树
func (c *Client) QueryTree(request *QueryTreeRequest) (response *QueryTreeResponse, err error) {
	if request == nil {
		request = NewQueryTreeRequest()
	}
	response = NewQueryTreeResponse()
	err = c.Send(request, response)
	return
}

func NewAddDirRequest() (request *AddDirRequest) {
	request = &AddDirRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "AddDir")
	return
}

func NewAddDirResponse() (response *AddDirResponse) {
	response = &AddDirResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增类别
func (c *Client) AddDir(request *AddDirRequest) (response *AddDirResponse, err error) {
	if request == nil {
		request = NewAddDirRequest()
	}
	response = NewAddDirResponse()
	err = c.Send(request, response)
	return
}

func NewDelDirRequest() (request *DelDirRequest) {
	request = &DelDirRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "DelDir")
	return
}

func NewDelDirResponse() (response *DelDirResponse) {
	response = &DelDirResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除类别
func (c *Client) DelDir(request *DelDirRequest) (response *DelDirResponse, err error) {
	if request == nil {
		request = NewDelDirRequest()
	}
	response = NewDelDirResponse()
	err = c.Send(request, response)
	return
}

func NewEditNodeRequest() (request *EditNodeRequest) {
	request = &EditNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "EditNode")
	return
}

func NewEditNodeResponse() (response *EditNodeResponse) {
	response = &EditNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑节点
func (c *Client) EditNode(request *EditNodeRequest) (response *EditNodeResponse, err error) {
	if request == nil {
		request = NewEditNodeRequest()
	}
	response = NewEditNodeResponse()
	err = c.Send(request, response)
	return
}

func NewQueryNodeUpgradeRequest() (request *QueryNodeUpgradeRequest) {
	request = &QueryNodeUpgradeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "QueryNodeUpgrade")
	return
}

func NewQueryNodeUpgradeResponse() (response *QueryNodeUpgradeResponse) {
	response = &QueryNodeUpgradeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 文档系统查询升级表节点信息
func (c *Client) QueryNodeUpgrade(request *QueryNodeUpgradeRequest) (response *QueryNodeUpgradeResponse, err error) {
	if request == nil {
		request = NewQueryNodeUpgradeRequest()
	}
	response = NewQueryNodeUpgradeResponse()
	err = c.Send(request, response)
	return
}

func NewUpgradeableRequest() (request *UpgradeableRequest) {
	request = &UpgradeableRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "Upgradeable")
	return
}

func NewUpgradeableResponse() (response *UpgradeableResponse) {
	response = &UpgradeableResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 文档系统查询是否可以升级
func (c *Client) Upgradeable(request *UpgradeableRequest) (response *UpgradeableResponse, err error) {
	if request == nil {
		request = NewUpgradeableRequest()
	}
	response = NewUpgradeableResponse()
	err = c.Send(request, response)
	return
}

func NewEditDirRequest() (request *EditDirRequest) {
	request = &EditDirRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("document", APIVersion, "EditDir")
	return
}

func NewEditDirResponse() (response *EditDirResponse) {
	response = &EditDirResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑目录
func (c *Client) EditDir(request *EditDirRequest) (response *EditDirResponse, err error) {
	if request == nil {
		request = NewEditDirRequest()
	}
	response = NewEditDirResponse()
	err = c.Send(request, response)
	return
}
