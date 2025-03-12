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

package v20200217

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-02-17"

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

func NewDeleteYumFileRequest() (request *DeleteYumFileRequest) {
	request = &DeleteYumFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("yumoss", APIVersion, "DeleteYumFile")
	return
}

func NewDeleteYumFileResponse() (response *DeleteYumFileResponse) {
	response = &DeleteYumFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除文件接口
func (c *Client) DeleteYumFile(request *DeleteYumFileRequest) (response *DeleteYumFileResponse, err error) {
	if request == nil {
		request = NewDeleteYumFileRequest()
	}
	response = NewDeleteYumFileResponse()
	err = c.Send(request, response)
	return
}

func NewSearchSyncHistoryRequest() (request *SearchSyncHistoryRequest) {
	request = &SearchSyncHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("yumoss", APIVersion, "SearchSyncHistory")
	return
}

func NewSearchSyncHistoryResponse() (response *SearchSyncHistoryResponse) {
	response = &SearchSyncHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询同步记录
func (c *Client) SearchSyncHistory(request *SearchSyncHistoryRequest) (response *SearchSyncHistoryResponse, err error) {
	if request == nil {
		request = NewSearchSyncHistoryRequest()
	}
	response = NewSearchSyncHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewCheckRouteEmptyRequest() (request *CheckRouteEmptyRequest) {
	request = &CheckRouteEmptyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("yumoss", APIVersion, "CheckRouteEmpty")
	return
}

func NewCheckRouteEmptyResponse() (response *CheckRouteEmptyResponse) {
	response = &CheckRouteEmptyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某仓库是否为空
func (c *Client) CheckRouteEmpty(request *CheckRouteEmptyRequest) (response *CheckRouteEmptyResponse, err error) {
	if request == nil {
		request = NewCheckRouteEmptyRequest()
	}
	response = NewCheckRouteEmptyResponse()
	err = c.Send(request, response)
	return
}

func NewSearchOperationHistoryRequest() (request *SearchOperationHistoryRequest) {
	request = &SearchOperationHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("yumoss", APIVersion, "SearchOperationHistory")
	return
}

func NewSearchOperationHistoryResponse() (response *SearchOperationHistoryResponse) {
	response = &SearchOperationHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询操作记录
func (c *Client) SearchOperationHistory(request *SearchOperationHistoryRequest) (response *SearchOperationHistoryResponse, err error) {
	if request == nil {
		request = NewSearchOperationHistoryRequest()
	}
	response = NewSearchOperationHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewInsertSyncConfigRequest() (request *InsertSyncConfigRequest) {
	request = &InsertSyncConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("yumoss", APIVersion, "InsertSyncConfig")
	return
}

func NewInsertSyncConfigResponse() (response *InsertSyncConfigResponse) {
	response = &InsertSyncConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建软件源同步配置
func (c *Client) InsertSyncConfig(request *InsertSyncConfigRequest) (response *InsertSyncConfigResponse, err error) {
	if request == nil {
		request = NewInsertSyncConfigRequest()
	}
	response = NewInsertSyncConfigResponse()
	err = c.Send(request, response)
	return
}

func NewSearchSyncConfigRequest() (request *SearchSyncConfigRequest) {
	request = &SearchSyncConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("yumoss", APIVersion, "SearchSyncConfig")
	return
}

func NewSearchSyncConfigResponse() (response *SearchSyncConfigResponse) {
	response = &SearchSyncConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询同步配置
func (c *Client) SearchSyncConfig(request *SearchSyncConfigRequest) (response *SearchSyncConfigResponse, err error) {
	if request == nil {
		request = NewSearchSyncConfigRequest()
	}
	response = NewSearchSyncConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSyncConfigRequest() (request *DeleteSyncConfigRequest) {
	request = &DeleteSyncConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("yumoss", APIVersion, "DeleteSyncConfig")
	return
}

func NewDeleteSyncConfigResponse() (response *DeleteSyncConfigResponse) {
	response = &DeleteSyncConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除软件源及其同步配置
func (c *Client) DeleteSyncConfig(request *DeleteSyncConfigRequest) (response *DeleteSyncConfigResponse, err error) {
	if request == nil {
		request = NewDeleteSyncConfigRequest()
	}
	response = NewDeleteSyncConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifySyncConfigRequest() (request *ModifySyncConfigRequest) {
	request = &ModifySyncConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("yumoss", APIVersion, "ModifySyncConfig")
	return
}

func NewModifySyncConfigResponse() (response *ModifySyncConfigResponse) {
	response = &ModifySyncConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更改同步配置
func (c *Client) ModifySyncConfig(request *ModifySyncConfigRequest) (response *ModifySyncConfigResponse, err error) {
	if request == nil {
		request = NewModifySyncConfigRequest()
	}
	response = NewModifySyncConfigResponse()
	err = c.Send(request, response)
	return
}

func NewSearchYumCatalogRequest() (request *SearchYumCatalogRequest) {
	request = &SearchYumCatalogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("yumoss", APIVersion, "SearchYumCatalog")
	return
}

func NewSearchYumCatalogResponse() (response *SearchYumCatalogResponse) {
	response = &SearchYumCatalogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询文件目录
func (c *Client) SearchYumCatalog(request *SearchYumCatalogRequest) (response *SearchYumCatalogResponse, err error) {
	if request == nil {
		request = NewSearchYumCatalogRequest()
	}
	response = NewSearchYumCatalogResponse()
	err = c.Send(request, response)
	return
}

func NewStartSyncRequest() (request *StartSyncRequest) {
	request = &StartSyncRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("yumoss", APIVersion, "StartSync")
	return
}

func NewStartSyncResponse() (response *StartSyncResponse) {
	response = &StartSyncResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用同步接口
func (c *Client) StartSync(request *StartSyncRequest) (response *StartSyncResponse, err error) {
	if request == nil {
		request = NewStartSyncRequest()
	}
	response = NewStartSyncResponse()
	err = c.Send(request, response)
	return
}

func NewUploadYumFileRequest() (request *UploadYumFileRequest) {
	request = &UploadYumFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("yumoss", APIVersion, "UploadYumFile")
	return
}

func NewUploadYumFileResponse() (response *UploadYumFileResponse) {
	response = &UploadYumFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 上传文件接口
func (c *Client) UploadYumFile(request *UploadYumFileRequest) (response *UploadYumFileResponse, err error) {
	if request == nil {
		request = NewUploadYumFileRequest()
	}
	response = NewUploadYumFileResponse()
	err = c.Send(request, response)
	return
}

func NewInsertNewFolderRequest() (request *InsertNewFolderRequest) {
	request = &InsertNewFolderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("yumoss", APIVersion, "InsertNewFolder")
	return
}

func NewInsertNewFolderResponse() (response *InsertNewFolderResponse) {
	response = &InsertNewFolderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建文件夹接口
func (c *Client) InsertNewFolder(request *InsertNewFolderRequest) (response *InsertNewFolderResponse, err error) {
	if request == nil {
		request = NewInsertNewFolderRequest()
	}
	response = NewInsertNewFolderResponse()
	err = c.Send(request, response)
	return
}

func NewRenameYumFileRequest() (request *RenameYumFileRequest) {
	request = &RenameYumFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("yumoss", APIVersion, "RenameYumFile")
	return
}

func NewRenameYumFileResponse() (response *RenameYumFileResponse) {
	response = &RenameYumFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重命名接口
func (c *Client) RenameYumFile(request *RenameYumFileRequest) (response *RenameYumFileResponse, err error) {
	if request == nil {
		request = NewRenameYumFileRequest()
	}
	response = NewRenameYumFileResponse()
	err = c.Send(request, response)
	return
}
