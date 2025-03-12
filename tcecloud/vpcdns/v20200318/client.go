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

package v20200318

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-03-18"

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

func NewCreateForwardRecordRequest() (request *CreateForwardRecordRequest) {
	request = &CreateForwardRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "CreateForwardRecord")
	return
}

func NewCreateForwardRecordResponse() (response *CreateForwardRecordResponse) {
	response = &CreateForwardRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建forward Record记录
func (c *Client) CreateForwardRecord(request *CreateForwardRecordRequest) (response *CreateForwardRecordResponse, err error) {
	if request == nil {
		request = NewCreateForwardRecordRequest()
	}
	response = NewCreateForwardRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOcloudVpcDnsRecordListRequest() (request *DescribeOcloudVpcDnsRecordListRequest) {
	request = &DescribeOcloudVpcDnsRecordListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "DescribeOcloudVpcDnsRecordList")
	return
}

func NewDescribeOcloudVpcDnsRecordListResponse() (response *DescribeOcloudVpcDnsRecordListResponse) {
	response = &DescribeOcloudVpcDnsRecordListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取解析记录列表
func (c *Client) DescribeOcloudVpcDnsRecordList(request *DescribeOcloudVpcDnsRecordListRequest) (response *DescribeOcloudVpcDnsRecordListResponse, err error) {
	if request == nil {
		request = NewDescribeOcloudVpcDnsRecordListRequest()
	}
	response = NewDescribeOcloudVpcDnsRecordListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcDnsLogStatusRequest() (request *DescribeVpcDnsLogStatusRequest) {
	request = &DescribeVpcDnsLogStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "DescribeVpcDnsLogStatus")
	return
}

func NewDescribeVpcDnsLogStatusResponse() (response *DescribeVpcDnsLogStatusResponse) {
	response = &DescribeVpcDnsLogStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取VPCDNS解析日志状态
func (c *Client) DescribeVpcDnsLogStatus(request *DescribeVpcDnsLogStatusRequest) (response *DescribeVpcDnsLogStatusResponse, err error) {
	if request == nil {
		request = NewDescribeVpcDnsLogStatusRequest()
	}
	response = NewDescribeVpcDnsLogStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeForwardRecordsRequest() (request *DescribeForwardRecordsRequest) {
	request = &DescribeForwardRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "DescribeForwardRecords")
	return
}

func NewDescribeForwardRecordsResponse() (response *DescribeForwardRecordsResponse) {
	response = &DescribeForwardRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获得Forward Record记录列表
func (c *Client) DescribeForwardRecords(request *DescribeForwardRecordsRequest) (response *DescribeForwardRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeForwardRecordsRequest()
	}
	response = NewDescribeForwardRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcDnsLogStatusRequest() (request *ModifyVpcDnsLogStatusRequest) {
	request = &ModifyVpcDnsLogStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "ModifyVpcDnsLogStatus")
	return
}

func NewModifyVpcDnsLogStatusResponse() (response *ModifyVpcDnsLogStatusResponse) {
	response = &ModifyVpcDnsLogStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改VPCDNS解析日志状态
func (c *Client) ModifyVpcDnsLogStatus(request *ModifyVpcDnsLogStatusRequest) (response *ModifyVpcDnsLogStatusResponse, err error) {
	if request == nil {
		request = NewModifyVpcDnsLogStatusRequest()
	}
	response = NewModifyVpcDnsLogStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVpcDnsClusterRequest() (request *CreateVpcDnsClusterRequest) {
	request = &CreateVpcDnsClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "CreateVpcDnsCluster")
	return
}

func NewCreateVpcDnsClusterResponse() (response *CreateVpcDnsClusterResponse) {
	response = &CreateVpcDnsClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建VPCDNS集群
func (c *Client) CreateVpcDnsCluster(request *CreateVpcDnsClusterRequest) (response *CreateVpcDnsClusterResponse, err error) {
	if request == nil {
		request = NewCreateVpcDnsClusterRequest()
	}
	response = NewCreateVpcDnsClusterResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcDnsClusterRequest() (request *ModifyVpcDnsClusterRequest) {
	request = &ModifyVpcDnsClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "ModifyVpcDnsCluster")
	return
}

func NewModifyVpcDnsClusterResponse() (response *ModifyVpcDnsClusterResponse) {
	response = &ModifyVpcDnsClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改VPCDNS集群信息
func (c *Client) ModifyVpcDnsCluster(request *ModifyVpcDnsClusterRequest) (response *ModifyVpcDnsClusterResponse, err error) {
	if request == nil {
		request = NewModifyVpcDnsClusterRequest()
	}
	response = NewModifyVpcDnsClusterResponse()
	err = c.Send(request, response)
	return
}

func NewModifyForwardRecordRequest() (request *ModifyForwardRecordRequest) {
	request = &ModifyForwardRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "ModifyForwardRecord")
	return
}

func NewModifyForwardRecordResponse() (response *ModifyForwardRecordResponse) {
	response = &ModifyForwardRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改ForwardRecord
func (c *Client) ModifyForwardRecord(request *ModifyForwardRecordRequest) (response *ModifyForwardRecordResponse, err error) {
	if request == nil {
		request = NewModifyForwardRecordRequest()
	}
	response = NewModifyForwardRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpcDnsClusterRequest() (request *DeleteVpcDnsClusterRequest) {
	request = &DeleteVpcDnsClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "DeleteVpcDnsCluster")
	return
}

func NewDeleteVpcDnsClusterResponse() (response *DeleteVpcDnsClusterResponse) {
	response = &DeleteVpcDnsClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除VPCDNS集群
func (c *Client) DeleteVpcDnsCluster(request *DeleteVpcDnsClusterRequest) (response *DeleteVpcDnsClusterResponse, err error) {
	if request == nil {
		request = NewDeleteVpcDnsClusterRequest()
	}
	response = NewDeleteVpcDnsClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcDnsClusterListRequest() (request *DescribeVpcDnsClusterListRequest) {
	request = &DescribeVpcDnsClusterListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "DescribeVpcDnsClusterList")
	return
}

func NewDescribeVpcDnsClusterListResponse() (response *DescribeVpcDnsClusterListResponse) {
	response = &DescribeVpcDnsClusterListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取VPCDNS集群列表
func (c *Client) DescribeVpcDnsClusterList(request *DescribeVpcDnsClusterListRequest) (response *DescribeVpcDnsClusterListResponse, err error) {
	if request == nil {
		request = NewDescribeVpcDnsClusterListRequest()
	}
	response = NewDescribeVpcDnsClusterListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteForwardRecordRequest() (request *DeleteForwardRecordRequest) {
	request = &DeleteForwardRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "DeleteForwardRecord")
	return
}

func NewDeleteForwardRecordResponse() (response *DeleteForwardRecordResponse) {
	response = &DeleteForwardRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Forward Record
func (c *Client) DeleteForwardRecord(request *DeleteForwardRecordRequest) (response *DeleteForwardRecordResponse, err error) {
	if request == nil {
		request = NewDeleteForwardRecordRequest()
	}
	response = NewDeleteForwardRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOcloudVpcDnsDomainListRequest() (request *DescribeOcloudVpcDnsDomainListRequest) {
	request = &DescribeOcloudVpcDnsDomainListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vpcdns", APIVersion, "DescribeOcloudVpcDnsDomainList")
	return
}

func NewDescribeOcloudVpcDnsDomainListResponse() (response *DescribeOcloudVpcDnsDomainListResponse) {
	response = &DescribeOcloudVpcDnsDomainListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取解析域列表
func (c *Client) DescribeOcloudVpcDnsDomainList(request *DescribeOcloudVpcDnsDomainListRequest) (response *DescribeOcloudVpcDnsDomainListResponse, err error) {
	if request == nil {
		request = NewDescribeOcloudVpcDnsDomainListRequest()
	}
	response = NewDescribeOcloudVpcDnsDomainListResponse()
	err = c.Send(request, response)
	return
}
